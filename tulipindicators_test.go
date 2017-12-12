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
	source := []float64{1.1, 2.2, 3.3, 4.4}
	newCPtr := castToCDoubleArray(source)

	for i, v := range source {
		//doing naughty things with go because cgo can't be used in tests.
		ptrIndex := uintptr(unsafe.Pointer(newCPtr)) + uintptr(sizeOfCDouble*i)
		check := (*float64)((unsafe.Pointer(ptrIndex)))

		if *check != v {
			t.Errorf("Expected value %d at index %d to equal %d", *check, i, v)
		}
	}
}

func TestCastToC2dDoubleArray(t *testing.T) {
	source := [][]float64{
		{1.1, 2.2, 3.3, 4.4},
		{5.5, 6.6},
		{7.7, 8.8, 9.9},
	}
	newCPtr := castToC2dDoubleArray(source)

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

func TestGet(t *testing.T) {
	for _, val := range indicatorList {
		info := Get(val)

		hasDoIndicator := info.DoIndicator != nil
		hasStartIndicator := info.StartIndicator != nil

		if !hasDoIndicator {
			t.Errorf("Expected do indicator for %s", val)
		}

		if !hasStartIndicator {
			t.Errorf("Expected start indicator for %s", val)
		}
	}
}

func TestStart(t *testing.T) {
	for _, val := range indicatorList {
		if startResult, startErr := Start(val, optionsd); startErr != nil {
			t.Errorf("Got an error from start function %s: %s", val, startErr.Error())
		}
	}
}

func TestDo(t *testing.T) {

	size := len(dummyIn)

	for _, val := range indicatorList {
		info := Get(val)

		inputs := make([][]float64, 0)
		inputs0 := make([][]float64, 0)
		outputs := make([][]float64, 0)

		for i := 0; i < 10; i++ {
			if i < info.inputs {
				inputs = append(inputs, dummyIn)
				inputs0 = append(inputs0, dummyIn0)
			}

			if i < info.inputs {
				outputs = append(outputs, dummyOt)
			}
		}

		doResult0, doErr := Do(val, 0, inputs, optionsd, outputs)
		doResult1, doErr := Do(val, 1, inputs, optionsd, outputs)
		doResult2, doErr := Do(val, 2, inputs, optionsd, outputs)
		doResult3, doErr := Do(val, 3, inputs, optionsd, outputs)
		doResultN, doErr := Do(val, size, inputs, optionsd, outputs)
		doResultZeros, doErr := Do(val, size, inputs0, optionsd, outputs)
	}

	/* for _, val := range indicatorList {


	} */
}
