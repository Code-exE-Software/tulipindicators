package tulipindicators

// #cgo LDFLAGS: -L./external -lindicators
// #include <external/indicators.h>
import (
	"C"
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
	startIndicators = map[string](func([]float64) (int, error)){
		"abs": AbsStart,
	}
	doIndicators = map[string](func(int, [][]float64, []float64, [][]float64) (int, error)){
		"abs": Abs,
	}
	tiTypes = map[int]string{
		1: "OVERLAY",     /* These have roughly the same range as the input data. */
		2: "INDICATOR",   /* Everything else (e.g. oscillators). */
		3: "MATH",        /* These aren't so good for plotting, but are useful with formulas. */
		4: "SIMPLE",      /* These apply a simple operator (e.g. addition, sin, sqrt). */
		5: "COMPARITIVE", /* These are designed to take inputs from different securities. i.e. compare stock A to stock B.*/
	}
)

func castDoParams(size int, inputs [][]float64, options []float64, outputs [][]float64) (C.int, **C.TI_REAL, *C.TI_REAL, **C.TI_REAL) {
	return C.int(size),
		castToC2dDoubleArray(inputs),
		castToCDoubleArray(options),
		castToC2dDoubleArray(outputs)
}

func castToCDoubleArray(source []float64) *C.TI_REAL {
	/* bar := (C.malloc(C.size_t(C.TI_REAL) * len(source)))

	cast := (*C.TI_REAL)(bar)

	for val, index := range source {
		cast + (index * C.size_t(C.TI_REAL)) = val
	}

	return cast */

	cast := make([]C.TI_REAL, len(source))

	for index, val := range source {
		cast[index] = C.double(val)
	}

	//return cast
	return &cast[0]
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
	start                                (func([]float64) (int, error))
	indicator                            (func(int, [][]float64, []float64, [][]float64) (int, error))
}

// Get ...
func Get(indicatorName string) IndicatorInfo {
	cIndicatorInfo := C.ti_find_indicator(C.CString(indicatorName))

	return IndicatorInfo{
		C.GoString(cIndicatorInfo.name),
		C.GoString(cIndicatorInfo.full_name),
		tiTypes[int(cIndicatorInfo._type)],
		int(cIndicatorInfo.inputs),
		int(cIndicatorInfo.options),
		int(cIndicatorInfo.outputs),
		getNames(cIndicatorInfo.input_names),
		getNames(cIndicatorInfo.option_names),
		getNames(cIndicatorInfo.output_names),
		startIndicators[indicatorName],
		doIndicators[indicatorName],
	}

}

// Start ...
func Start(indicatorName string, options []float64) (int, error) {
	return startIndicators[indicatorName](options)
}

// Do ...
func Do(indicatorName string, size int, inputs [][]float64, options []float64, outputs [][]float64) (int, error) {
	return doIndicators[indicatorName](size, inputs, options, outputs)
}

// AbsStart ...
func AbsStart(options []float64) (int, error) {
	castOptions := castToCDoubleArray(options)

	startResponse, startErr := C.ti_abs_start(castOptions)

	return int(startResponse), startErr
}

// Abs ...
func Abs(size int, inputs [][]float64, options []float64, outputs [][]float64) (int, error) {
	castSize, castInputs, castOptions, castOutputs := castDoParams(size, inputs, options, outputs)

	doResponse, doError := C.ti_abs(castSize, castInputs, castOptions, castOutputs)

	return int(doResponse), doError
}
