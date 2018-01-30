package tulipindicators

/*
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <stdlib.h>
#include "indicators.h"
*/
import (
	"C"
)
import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

const (
	maxIndicators = 10
)

type indicatorFunc = func([][]float64, []float64) ([][]float64, error)

//IndicatorInfo Go implementation of C.struct_ti_indicator_info
type IndicatorInfo struct {
	Name                                 string
	FullName                             string
	IndicatorType                        string
	Inputs, Options, Outputs             int
	InputNames, OptionNames, OutputNames []string
	Indicator                            indicatorFunc
}

var (
	// IndicatorInfos map of structs showing info and requirements for a given indicator.
	IndicatorInfos = map[string]IndicatorInfo{}

	// Indicators map of functions allowing consumers of this lib to call the indicator functions directly.
	Indicators = map[string]indicatorFunc{}

	tiTypes = map[int]string{
		1: "OVERLAY",     /* These have roughly the same range as the input data. */
		2: "INDICATOR",   /* Everything else (e.g. oscillators). */
		3: "MATH",        /* These aren't so good for plotting, but are useful with formulas. */
		4: "SIMPLE",      /* These apply a simple operator (e.g. addition, sin, sqrt). */
		5: "COMPARITIVE", /* These are designed to take inputs from different securities. i.e. compare stock A to stock B.*/
	}
)

func castToCDoubleArray(source []float64) *C.TI_REAL {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, source)
	return (*C.TI_REAL)(C.CBytes(buf.Bytes()))
}

func castToC2dDoubleArray(source [][]float64) (**C.TI_REAL, [][]float64) {
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
	return &cast[0], validSource
}

func extractOutputs(cOutputs **C.double, goOutputs [][]float64) [][]float64 {
	for outerIndex, outerVal := range goOutputs {
		ptrOuter := uintptr(unsafe.Pointer(cOutputs)) + uintptr(C.sizeof_TI_REAL*outerIndex)

		for innerIndex, _ := range outerVal {
			ptrInner := (*unsafe.Pointer)(unsafe.Pointer(ptrOuter))
			ptrInnerIndex := uintptr(*ptrInner) + uintptr(C.sizeof_TI_REAL*innerIndex)
			val := (*float64)((unsafe.Pointer(ptrInnerIndex)))
			goOutputs[outerIndex][innerIndex] = *val
		}
	}

	return goOutputs
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

// Init Initialize the library
func init() {
	var cIndicatorInfo *C.ti_indicator_info

	for _, name := range indicatorNames {
		cIndicatorInfo = C.ti_find_indicator(C.CString(name))
		IndicatorInfos[name] = IndicatorInfo{
			C.GoString(cIndicatorInfo.name),
			C.GoString(cIndicatorInfo.full_name),
			tiTypes[int(cIndicatorInfo._type)],
			int(cIndicatorInfo.inputs),
			int(cIndicatorInfo.options),
			int(cIndicatorInfo.outputs),
			getNames(cIndicatorInfo.input_names),
			getNames(cIndicatorInfo.option_names),
			getNames(cIndicatorInfo.output_names),
			func(closureIndicatorInfo *C.ti_indicator_info) indicatorFunc {
				return func(inputs [][]float64, options []float64) ([][]float64, error) {
					return indicator(
						int(closureIndicatorInfo.outputs),
						closureIndicatorInfo.start,
						closureIndicatorInfo.indicator,
						inputs,
						options,
					)
				}
			}(cIndicatorInfo),
		}

		Indicators[name] = IndicatorInfos[name].Indicator
	}
}
