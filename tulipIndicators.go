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

var (
	startIndicators = map[string](func([]float64) int){
		"absStart": AbsStart,
	}
)

// Start ...
func Start(indicatorName string, options []float64) int {
	return startIndicators[indicatorName](options)
}

// Do ...
type Do func(size int, inputs [][]float64, options []float64, outputs [][]float64) int

// AbsStart ...
func AbsStart(options []float64) int {
	castOptions := castToDoubleArray(options)

	startResponse := C.ti_abs_start(&castOptions[0])

	return int(startResponse)
}

func castToDoubleArray(source []float64) []C.double {
	cast := make([]C.double, 0)

	for v := range source {
		cast = append(cast, C.double(v))
	}

	return cast
}
