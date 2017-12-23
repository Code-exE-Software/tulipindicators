package tulipindicators

import (
	"testing"
	"unsafe"
)

const (
	sizeOfCDouble = 8
)

var (
	indicatorList = []string{
		"abs",
	}
	//copied from the fuzzer.c file
	optionsd = []float64{-20, -1, 0, .1, .5, .7, 1, 2, 2.5, 3, 4, 5, 6, 7, 8, 9, 10, 11, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 100}
	dummyIn  = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	dummyIn0 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	dummyOt  = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

func TestCastToCDoubleArray(t *testing.T) {
	source := optionsd //[]float64{1.1, 2.2, 3.3, 4.4}
	newCPtr := castToCDoubleArray(source)

	for i, v := range source {
		//doing naughty things with go because cgo can't be used in tests.
		ptrIndex := uintptr(unsafe.Pointer(newCPtr)) + uintptr(sizeOfCDouble*i)
		check := (*float64)((unsafe.Pointer(ptrIndex)))

		if *check != v {
			t.Errorf("Expected value %v at index %d to equal %v", *check, i, v)
		}
	}
}

func TestCastToC2dDoubleArray(t *testing.T) {
	source := [][]float64{
		{1.1, 2.2, 3.3, 4.4},
		{5.5, 6.6},
		{7.7, 8.8, 9.9},
		{1.1, 2.2, 3.3, 4.4},
		{5.5, 6.6},
		{7.7, 8.8, 9.9},
		{1.1, 2.2, 3.3, 4.4},
		{5.5, 6.6},
		{7.7, 8.8, 9.9},
		{1.1, 2.2, 3.3, 4.4},
		{5.5, 6.6},
		{7.7, 8.8, 9.9},
	}
	newCPtr, source := castToC2dDoubleArray(source)

	//doing naughty things with go because cgo can't be used in tests.
	for outerIndex, outerVal := range source {

		//@todo I expect this unit test will only work when testing the function on 64 bit systems.
		ptrOuter := uintptr(unsafe.Pointer(newCPtr)) + uintptr(sizeOfCDouble*outerIndex)

		for innerIndex, innerVal := range outerVal {
			ptrInner := (*unsafe.Pointer)(unsafe.Pointer(ptrOuter))
			ptrInnerIndex := uintptr(*ptrInner) + uintptr(sizeOfCDouble*innerIndex)
			check := (*float64)((unsafe.Pointer(ptrInnerIndex)))

			if *check != innerVal {
				t.Errorf("Expected value %v at index %d to equal %v", *check, innerIndex, innerVal)
			}
		}
	}
}

func TestFreeCDoubleArray(t *testing.T) {
	source := []float64{1.1, 2.2, 3.3, 4.4}
	newCPtr := castToCDoubleArray(source)

	freeCDoubleArray(newCPtr)
}

func TestFreeC2dDoubleArray(t *testing.T) {
	source := [][]float64{
		{1.1, 2.2, 3.3, 4.4},
		{5.5, 6.6},
		{7.7, 8.8, 9.9},
	}
	newCPtr, source := castToC2dDoubleArray(source)

	freeC2dDoubleArray(newCPtr, len(source))
}

func TestIndicator(t *testing.T) {

	size := len(dummyIn)

	for _, val := range indicatorNames {
		var info IndicatorInfo
		var ok bool

		if info, ok = IndicatorInfos[val]; !ok {
			t.Errorf("Unable to get info for indicator %s.  Did the TestGet function fail here too?", val)
		}

		inputs := make([][]float64, 0)
		inputs0 := make([][]float64, 0)
		outputs := make([][]float64, 0)

		for i := 0; i < 10; i++ {
			if i < info.Inputs {
				inputs = append(inputs, dummyIn)
				inputs0 = append(inputs0, dummyIn0)
			}

			if i < info.Outputs {
				outputs = append(outputs, dummyOt)
			}
		}

		args := [](struct {
			size   int
			inputs [][]float64
		}){
			{0, inputs},
			{1, inputs},
			{2, inputs},
			{3, inputs},
			{size, inputs},
			{size, inputs0},
		}

		for _, argsVal := range args {
			var doErr error
			var doOutputs [][]float64

			if doOutputs, doErr = Indicators[val](argsVal.inputs, optionsd); doErr != nil {
				t.Errorf("Error thrown from indicator function %s: %s", val, doErr.Error())
				break
			}

			t.Logf("Do function for %s returned value %v", val, doOutputs)
		}
	}
}

func TestIndicatorNoOptions(t *testing.T) {

	numEntries := 10 /* * 24 */

	inputs := append(make([][]float64, 0), make([]float64, numEntries))

	for i := 0; i < numEntries; i++ {
		inputs[0][i] = float64(i * -1)
	}

	var outputs [][]float64
	var err error

	if outputs, err = Indicators["abs"](inputs, nil); err != nil {
		t.Errorf("ERRRRRR %v \n", err)
		return
	}

	for i := 0; i < numEntries; i++ {
		if outputs[0][i] != float64(i) {
			t.Errorf("Bad output %d expected %d", outputs[0][i], float64(i))
		}
	}
}
