package tulipindicators

/*
 #cgo LDFLAGS: -L./external -lindicators
 #include <external/indicators.h>

 int bridgeStartFunction(ti_indicator_start_function f, TI_REAL const *options) {
    return f(options);
 }

 int bridgeIndicatorFunction(ti_indicator_function f,
    int size,
    TI_REAL const *const *inputs,
    TI_REAL const *options,
    TI_REAL *const *outputs) {
    return f(size, inputs, options, outputs);
 }
*/
import (
	"C"
)

func indicator(
	numOutputs int,
	startFunc /* unsafe.Pointer, */ C.ti_indicator_start_function,
	indicatorFunc /* unsafe.Pointer, */ C.ti_indicator_function,
	size int,
	inputs [][]float64,
	options []float64,
) (int, [][]float64, error) {

	castSize := C.int(size)

	castOptions := castToCDoubleArray(options)
	defer freeCDoubleArray(castOptions)

	castInputs, inputs := castToC2dDoubleArray(inputs)
	defer freeC2dDoubleArray(castInputs, len(inputs))

	outputSizeDiff := C.bridgeStartFunction(C.ti_indicator_start_function(startFunc), castOptions)

	outputs := make([][]float64, numOutputs)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs, outputs := castToC2dDoubleArray(outputs)
	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.bridgeIndicatorFunction(
		C.ti_indicator_function(indicatorFunc),
		castSize,
		castInputs,
		castOptions,
		castOutputs,
	)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}
