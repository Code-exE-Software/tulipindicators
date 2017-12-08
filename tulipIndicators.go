package tulipIndicators

// #cgo LDFLAGS: -L./external -lindicators
// #include <external/indicators.h>
import (
	"C"
)

/* type Start func(options []float64) int
type Do func(size int, inputs []float64, options []float64, outputs []float64) int */

/* Vector Absolute Value */
/* Type: simple */
/* Input arrays: 1    Options: 0    Output arrays: 1 */
/* Inputs: real */
/* Options: none */
/* Outputs: abs */
/* int ti_abs_start(TI_REAL const *options);
int ti_abs(int size,
      TI_REAL const *const *inputs,
      TI_REAL const *options,
	  TI_REAL *const *outputs); */

//type indicatorInfo C.struct_ti_indicator_info

func foobar() {

	info := C.ti_find_indicator(C.CString("abs"))
}
