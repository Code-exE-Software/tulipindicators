package tulipindicators

// #cgo LDFLAGS: -L./external -lindicators
// #include <external/indicators.h>
// #include <stdlib.h>
import (
	"C"
)
import (
	"fmt"
	"unsafe"
)

/*int ti_abs_start(TI_REAL const *options);
int ti_abs(int size,
      TI_REAL const *const *inputs,
      TI_REAL const *options,
      TI_REAL *const *outputs);
*/
const (
	maxIndicators = 10
)

var (
	doIndicatorFuncs = map[string](func(int, [][]float64, []float64) (int, [][]float64, error)){
		"abs": abs,
	}
	tiTypes = map[int]string{
		1: "OVERLAY",     /* These have roughly the same range as the input data. */
		2: "INDICATOR",   /* Everything else (e.g. oscillators). */
		3: "MATH",        /* These aren't so good for plotting, but are useful with formulas. */
		4: "SIMPLE",      /* These apply a simple operator (e.g. addition, sin, sqrt). */
		5: "COMPARITIVE", /* These are designed to take inputs from different securities. i.e. compare stock A to stock B.*/
	}
	memoizedIndicatorInfo = map[string]IndicatorInfo{}
)

func castDoParams(size int, inputs [][]float64, options []float64) (C.int, **C.TI_REAL, *C.TI_REAL) {
	return C.int(size),
		castToC2dDoubleArray(inputs),
		castToCDoubleArray(options)
}

func castToCDoubleArray(source []float64) *C.TI_REAL {
	mallocBytes := C.sizeof_TI_REAL * len(source)
	cast := (*C.TI_REAL)(C.malloc(C.size_t(mallocBytes)))

	for index, val := range source {
		offset := index * C.sizeof_TI_REAL
		ptrIndex := uintptr(unsafe.Pointer(cast)) + uintptr(offset)
		castAtIndex := (*C.double)(unsafe.Pointer(ptrIndex))
		*castAtIndex = ((C.double)(val))
	}
	return cast
}

func castToC2dDoubleArray(source [][]float64) **C.TI_REAL {
	// per tulipindicators docs, there should never be more than 10 sets of indicator numbers here.
	validSource := source
	if len(validSource) > 10 {
		validSource = validSource[:10]
	}

	cast := make([]*C.TI_REAL, len(validSource))

	for outerIndex, row := range validSource {
		inner := castToCDoubleArray(row)

		cast[outerIndex] = inner
	}

	//return cast
	return &cast[0]
}

func extractOutputs(cOutputs **C.double, goOutputs *([][]float64)) {
	//doing naughty things with go because cgo can't be used in tests.
	for outerIndex, outerVal := range *goOutputs {

		//@todo I expect this unit test will only work when testing the function on 64 bit systems.
		ptrOuter := uintptr(unsafe.Pointer(cOutputs)) + uintptr(C.sizeof_TI_REAL*outerIndex)

		for innerIndex := range outerVal {
			ptrInner := (*unsafe.Pointer)(unsafe.Pointer(ptrOuter))
			ptrInnerIndex := uintptr(*ptrInner) + uintptr(C.sizeof_TI_REAL*innerIndex)
			val := (*float64)((unsafe.Pointer(ptrInnerIndex)))

			(*goOutputs)[outerIndex][innerIndex] = *val
		}
	}
}

func freeCDoubleArray(source *C.TI_REAL) {
	C.free(unsafe.Pointer(source))
}

func freeC2dDoubleArray(source **C.TI_REAL, length int) {
	for i := 0; i < length; i++ {
		ptrOuterAddress := uintptr(unsafe.Pointer(source)) + uintptr(C.sizeof_TI_REAL*i)
		ptrOuter := (*unsafe.Pointer)(unsafe.Pointer(ptrOuterAddress))
		ptrInner := (*C.double)(unsafe.Pointer(uintptr(*ptrOuter)))

		freeCDoubleArray(ptrInner)
	}
}

func getNames(source [10]*C.char) []string {
	result := make([]string, 10)

	for index := 0; index < 10; index++ {
		result[index] = C.GoString(source[index])
	}

	return result
}

//IndicatorInfo ..
type IndicatorInfo struct {
	name                                 string
	fullName                             string
	indicatorType                        string
	inputs, options, outputs             int
	inputNames, optionNames, outputNames []string
	indicator                            (func(int, [][]float64, []float64) (int, [][]float64, error))
}

// Get ...
func Get(indicatorName string) (IndicatorInfo, error) {
	if memoizedInfo, ok := memoizedIndicatorInfo[indicatorName]; ok {
		return memoizedInfo, nil
	}

	var cIndicatorInfo *C.ti_indicator_info
	var findErr error

	if cIndicatorInfo, findErr = C.ti_find_indicator(C.CString(indicatorName)); findErr != nil {
		return IndicatorInfo{}, findErr
	}

	memoizedIndicatorInfo[indicatorName] = IndicatorInfo{
		C.GoString(cIndicatorInfo.name),
		C.GoString(cIndicatorInfo.full_name),
		tiTypes[int(cIndicatorInfo._type)],
		int(cIndicatorInfo.inputs),
		int(cIndicatorInfo.options),
		int(cIndicatorInfo.outputs),
		getNames(cIndicatorInfo.input_names),
		getNames(cIndicatorInfo.option_names),
		getNames(cIndicatorInfo.output_names),
		doIndicatorFuncs[indicatorName],
	}

	return memoizedIndicatorInfo[indicatorName], nil
}

// Indicator ...
func Indicator(indicatorName string, size int, inputs [][]float64, options []float64) (int, [][]float64, error) {
	var info IndicatorInfo
	var getErr error
	if info, getErr = Get(indicatorName); getErr != nil {
		return 0, [][]float64{}, getErr
	}

	return info.indicator(size, inputs, options)
}

// Init Initialize the library
func Init() error {
	for name := range doIndicatorFuncs {
		if _, getErr := Get(name); getErr != nil {
			return getErr
		}
	}

	return nil
}

// Abs ...
func abs(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["abs"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_abs_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_abs(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}
