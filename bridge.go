package tulipindicators

/*
 #cgo LDFLAGS: -L./external -lindicators
 #include <external/indicators.h>
 #include <stdio.h>

 int bridgeStartFunction(ti_indicator_start_function f, TI_REAL const *options) {

	return f(options);
 }

 int bridgeIndicatorFunction(ti_indicator_function f,
    int size,
    TI_REAL const *const *inputs,
    TI_REAL const *options,
    TI_REAL *const *outputs) {
		printf("why?");
		int result = f(size, inputs, options, outputs);

		char outputstr[25];
		sprintf(outputstr, "%d", result);
		printf(outputstr);
		printf("\n");
		//return 5;
		return result;
 }
*/
import (
	"C"
)
import (
	"fmt"
)

func indicator(
	numOutputs int,
	startFunc /* unsafe.Pointer, */ C.ti_indicator_start_function,
	indicatorFunc /* unsafe.Pointer, */ C.ti_indicator_function,
	inputs [][]float64,
	options []float64,
) ([][]float64, error) {

	castOptions := castToCDoubleArray(options)
	defer freeCDoubleArray(castOptions)

	castInputs, inputs := castToC2dDoubleArray(inputs)
	defer freeC2dDoubleArray(castInputs, len(inputs))

	outputSizeDiff := C.bridgeStartFunction(startFunc, castOptions)

	outputSize := len(inputs[0]) - int(outputSizeDiff)

	if outputSize < 1 {
		return nil, fmt.Errorf("insufficient inputs")
	}

	outputs := make([][]float64, numOutputs)

	for i := range outputs {
		outputs[i] = make([]float64, outputSize)
	}

	castOutputs, outputs := castToC2dDoubleArray(outputs)
	defer freeC2dDoubleArray(castOutputs, len(outputs))

	castOutputSize := C.int(outputSize)

	doResponse, doError := C.bridgeIndicatorFunction(
		indicatorFunc,
		castOutputSize,
		castInputs,
		castOptions,
		castOutputs,
	)

	if doError != nil {
		fmt.Printf("Windows error generated here:   \n%v\n%v\n", indicatorFunc, doResponse)
		//return nil, doError //skipping error because the output *is* actually valid.  SCARY
	}

	if doResponse == C.TI_INVALID_OPTION {
		return nil, fmt.Errorf("invalid Option for TulipIndicator")
	}

	outputs = extractOutputs(castOutputs, outputs)

	return outputs, nil
}
