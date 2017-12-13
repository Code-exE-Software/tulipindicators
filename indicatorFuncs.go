package tulipindicators

// #cgo LDFLAGS: -L./external -lindicators
// #include <external/indicators.h>
import (
	"C"
)
import "fmt"

/*

   {"aroonosc", "Aroon Oscillator", ti_aroonosc_start, ti_aroonosc, TI_TYPE_INDICATOR, 2, 1, 1, {"high","low",0}, {"period",0}, {"aroonosc",0}},
   {"asin", "Vector Arcsine", ti_asin_start, ti_asin, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"asin",0}},
   {"atan", "Vector Arctangent", ti_atan_start, ti_atan, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"atan",0}},
   {"atr", "Average True Range", ti_atr_start, ti_atr, TI_TYPE_INDICATOR, 3, 1, 1, {"high","low","close",0}, {"period",0}, {"atr",0}},
   {"avgprice", "Average Price", ti_avgprice_start, ti_avgprice, TI_TYPE_OVERLAY, 4, 0, 1, {"open","high","low","close",0}, {"",0}, {"avgprice",0}},
   {"bbands", "Bollinger Bands", ti_bbands_start, ti_bbands, TI_TYPE_OVERLAY, 1, 2, 3, {"real",0}, {"period","stddev",0}, {"bbands_lower","bbands_middle","bbands_upper",0}},
   {"bop", "Balance of Power", ti_bop_start, ti_bop, TI_TYPE_INDICATOR, 4, 0, 1, {"open","high","low","close",0}, {"",0}, {"bop",0}},
   {"cci", "Commodity Channel Index", ti_cci_start, ti_cci, TI_TYPE_INDICATOR, 3, 1, 1, {"high","low","close",0}, {"period",0}, {"cci",0}},
   {"ceil", "Vector Ceiling", ti_ceil_start, ti_ceil, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"ceil",0}},
   {"cmo", "Chande Momentum Oscillator", ti_cmo_start, ti_cmo, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"cmo",0}},
   {"cos", "Vector Cosine", ti_cos_start, ti_cos, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"cos",0}},
   {"cosh", "Vector Hyperbolic Cosine", ti_cosh_start, ti_cosh, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"cosh",0}},
   {"crossany", "Crossany", ti_crossany_start, ti_crossany, TI_TYPE_MATH, 2, 0, 1, {"real","real",0}, {"",0}, {"crossany",0}},
   {"crossover", "Crossover", ti_crossover_start, ti_crossover, TI_TYPE_MATH, 2, 0, 1, {"real","real",0}, {"",0}, {"crossover",0}},
   {"cvi", "Chaikins Volatility", ti_cvi_start, ti_cvi, TI_TYPE_INDICATOR, 2, 1, 1, {"high","low",0}, {"period",0}, {"cvi",0}},
   {"decay", "Linear Decay", ti_decay_start, ti_decay, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"decay",0}},
   {"dema", "Double Exponential Moving Average", ti_dema_start, ti_dema, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"dema",0}},
   {"di", "Directional Indicator", ti_di_start, ti_di, TI_TYPE_INDICATOR, 3, 1, 2, {"high","low","close",0}, {"period",0}, {"plus_di","minus_di",0}},
   {"div", "Vector Division", ti_div_start, ti_div, TI_TYPE_SIMPLE, 2, 0, 1, {"real","real",0}, {"",0}, {"div",0}},
   {"dm", "Directional Movement", ti_dm_start, ti_dm, TI_TYPE_INDICATOR, 2, 1, 2, {"high","low",0}, {"period",0}, {"plus_dm","minus_dm",0}},
   {"dpo", "Detrended Price Oscillator", ti_dpo_start, ti_dpo, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"dpo",0}},
   {"dx", "Directional Movement Index", ti_dx_start, ti_dx, TI_TYPE_INDICATOR, 3, 1, 1, {"high","low","close",0}, {"period",0}, {"dx",0}},
   {"edecay", "Exponential Decay", ti_edecay_start, ti_edecay, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"edecay",0}},
   {"ema", "Exponential Moving Average", ti_ema_start, ti_ema, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"ema",0}},
   {"emv", "Ease of Movement", ti_emv_start, ti_emv, TI_TYPE_INDICATOR, 3, 0, 1, {"high","low","volume",0}, {"",0}, {"emv",0}},
   {"exp", "Vector Exponential", ti_exp_start, ti_exp, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"exp",0}},
   {"fisher", "Fisher Transform", ti_fisher_start, ti_fisher, TI_TYPE_INDICATOR, 2, 1, 2, {"high","low",0}, {"period",0}, {"fisher","fisher_signal",0}},
   {"floor", "Vector Floor", ti_floor_start, ti_floor, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"floor",0}},
   {"fosc", "Forecast Oscillator", ti_fosc_start, ti_fosc, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"fosc",0}},
   {"hma", "Hull Moving Average", ti_hma_start, ti_hma, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"hma",0}},
   {"kama", "Kaufman Adaptive Moving Average", ti_kama_start, ti_kama, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"kama",0}},
   {"kvo", "Klinger Volume Oscillator", ti_kvo_start, ti_kvo, TI_TYPE_INDICATOR, 4, 2, 1, {"high","low","close","volume",0}, {"short period","long period",0}, {"kvo",0}},
   {"lag", "Lag", ti_lag_start, ti_lag, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"lag",0}},
   {"linreg", "Linear Regression", ti_linreg_start, ti_linreg, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"linreg",0}},
   {"linregintercept", "Linear Regression Intercept", ti_linregintercept_start, ti_linregintercept, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"linregintercept",0}},
   {"linregslope", "Linear Regression Slope", ti_linregslope_start, ti_linregslope, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"linregslope",0}},
   {"ln", "Vector Natural Log", ti_ln_start, ti_ln, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"ln",0}},
   {"log10", "Vector Base-10 Log", ti_log10_start, ti_log10, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"log10",0}},
   {"macd", "Moving Average Convergence/Divergence", ti_macd_start, ti_macd, TI_TYPE_INDICATOR, 1, 3, 3, {"real",0}, {"short period","long period","signal period",0}, {"macd","macd_signal","macd_histogram",0}},
   {"marketfi", "Market Facilitation Index", ti_marketfi_start, ti_marketfi, TI_TYPE_INDICATOR, 3, 0, 1, {"high","low","volume",0}, {"",0}, {"marketfi",0}},
   {"mass", "Mass Index", ti_mass_start, ti_mass, TI_TYPE_INDICATOR, 2, 1, 1, {"high","low",0}, {"period",0}, {"mass",0}},
   {"max", "Maximum In Period", ti_max_start, ti_max, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"max",0}},
   {"md", "Mean Deviation Over Period", ti_md_start, ti_md, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"md",0}},
   {"medprice", "Median Price", ti_medprice_start, ti_medprice, TI_TYPE_OVERLAY, 2, 0, 1, {"high","low",0}, {"",0}, {"medprice",0}},
   {"mfi", "Money Flow Index", ti_mfi_start, ti_mfi, TI_TYPE_INDICATOR, 4, 1, 1, {"high","low","close","volume",0}, {"period",0}, {"mfi",0}},
   {"min", "Minimum In Period", ti_min_start, ti_min, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"min",0}},
   {"mom", "Momentum", ti_mom_start, ti_mom, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"mom",0}},
   {"msw", "Mesa Sine Wave", ti_msw_start, ti_msw, TI_TYPE_INDICATOR, 1, 1, 2, {"real",0}, {"period",0}, {"msw_sine","msw_lead",0}},
   {"mul", "Vector Multiplication", ti_mul_start, ti_mul, TI_TYPE_SIMPLE, 2, 0, 1, {"real","real",0}, {"",0}, {"mul",0}},
   {"natr", "Normalized Average True Range", ti_natr_start, ti_natr, TI_TYPE_INDICATOR, 3, 1, 1, {"high","low","close",0}, {"period",0}, {"natr",0}},
   {"nvi", "Negative Volume Index", ti_nvi_start, ti_nvi, TI_TYPE_INDICATOR, 2, 0, 1, {"close","volume",0}, {"",0}, {"nvi",0}},
   {"obv", "On Balance Volume", ti_obv_start, ti_obv, TI_TYPE_INDICATOR, 2, 0, 1, {"close","volume",0}, {"",0}, {"obv",0}},
   {"ppo", "Percentage Price Oscillator", ti_ppo_start, ti_ppo, TI_TYPE_INDICATOR, 1, 2, 1, {"real",0}, {"short period","long period",0}, {"ppo",0}},
   {"psar", "Parabolic SAR", ti_psar_start, ti_psar, TI_TYPE_OVERLAY, 2, 2, 1, {"high","low",0}, {"acceleration factor step","acceleration factor maximum",0}, {"psar",0}},
   {"pvi", "Positive Volume Index", ti_pvi_start, ti_pvi, TI_TYPE_INDICATOR, 2, 0, 1, {"close","volume",0}, {"",0}, {"pvi",0}},
   {"qstick", "Qstick", ti_qstick_start, ti_qstick, TI_TYPE_INDICATOR, 2, 1, 1, {"open","close",0}, {"period",0}, {"qstick",0}},
   {"roc", "Rate of Change", ti_roc_start, ti_roc, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"roc",0}},
   {"rocr", "Rate of Change Ratio", ti_rocr_start, ti_rocr, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"rocr",0}},
   {"round", "Vector Round", ti_round_start, ti_round, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"round",0}},
   {"rsi", "Relative Strength Index", ti_rsi_start, ti_rsi, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"rsi",0}},
   {"sin", "Vector Sine", ti_sin_start, ti_sin, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"sin",0}},
   {"sinh", "Vector Hyperbolic Sine", ti_sinh_start, ti_sinh, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"sinh",0}},
   {"sma", "Simple Moving Average", ti_sma_start, ti_sma, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"sma",0}},
   {"sqrt", "Vector Square Root", ti_sqrt_start, ti_sqrt, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"sqrt",0}},
   {"stddev", "Standard Deviation Over Period", ti_stddev_start, ti_stddev, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"stddev",0}},
   {"stderr", "Standard Error Over Period", ti_stderr_start, ti_stderr, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"stderr",0}},
   {"stoch", "Stochastic Oscillator", ti_stoch_start, ti_stoch, TI_TYPE_INDICATOR, 3, 3, 2, {"high","low","close",0}, {"%k period","%k slowing period","%d period",0}, {"stoch_k","stoch_d",0}},
   {"sub", "Vector Subtraction", ti_sub_start, ti_sub, TI_TYPE_SIMPLE, 2, 0, 1, {"real","real",0}, {"",0}, {"sub",0}},
   {"sum", "Sum Over Period", ti_sum_start, ti_sum, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"sum",0}},
   {"tan", "Vector Tangent", ti_tan_start, ti_tan, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"tan",0}},
   {"tanh", "Vector Hyperbolic Tangent", ti_tanh_start, ti_tanh, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"tanh",0}},
   {"tema", "Triple Exponential Moving Average", ti_tema_start, ti_tema, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"tema",0}},
   {"todeg", "Vector Degree Conversion", ti_todeg_start, ti_todeg, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"degrees",0}},
   {"torad", "Vector Radian Conversion", ti_torad_start, ti_torad, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"radians",0}},
   {"tr", "True Range", ti_tr_start, ti_tr, TI_TYPE_INDICATOR, 3, 0, 1, {"high","low","close",0}, {"",0}, {"tr",0}},
   {"trima", "Triangular Moving Average", ti_trima_start, ti_trima, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"trima",0}},
   {"trix", "Trix", ti_trix_start, ti_trix, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"trix",0}},
   {"trunc", "Vector Truncate", ti_trunc_start, ti_trunc, TI_TYPE_SIMPLE, 1, 0, 1, {"real",0}, {"",0}, {"trunc",0}},
   {"tsf", "Time Series Forecast", ti_tsf_start, ti_tsf, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"tsf",0}},
   {"typprice", "Typical Price", ti_typprice_start, ti_typprice, TI_TYPE_OVERLAY, 3, 0, 1, {"high","low","close",0}, {"",0}, {"typprice",0}},
   {"ultosc", "Ultimate Oscillator", ti_ultosc_start, ti_ultosc, TI_TYPE_INDICATOR, 3, 3, 1, {"high","low","close",0}, {"short period","medium period","long period",0}, {"ultosc",0}},
   {"var", "Variance Over Period", ti_var_start, ti_var, TI_TYPE_MATH, 1, 1, 1, {"real",0}, {"period",0}, {"var",0}},
   {"vhf", "Vertical Horizontal Filter", ti_vhf_start, ti_vhf, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"vhf",0}},
   {"vidya", "Variable Index Dynamic Average", ti_vidya_start, ti_vidya, TI_TYPE_OVERLAY, 1, 3, 1, {"real",0}, {"short period","long period","alpha",0}, {"vidya",0}},
   {"volatility", "Annualized Historical Volatility", ti_volatility_start, ti_volatility, TI_TYPE_INDICATOR, 1, 1, 1, {"real",0}, {"period",0}, {"volatility",0}},
   {"vosc", "Volume Oscillator", ti_vosc_start, ti_vosc, TI_TYPE_INDICATOR, 1, 2, 1, {"volume",0}, {"short period","long period",0}, {"vosc",0}},
   {"vwma", "Volume Weighted Moving Average", ti_vwma_start, ti_vwma, TI_TYPE_OVERLAY, 2, 1, 1, {"close","volume",0}, {"period",0}, {"vwma",0}},
   {"wad", "Williams Accumulation/Distribution", ti_wad_start, ti_wad, TI_TYPE_INDICATOR, 3, 0, 1, {"high","low","close",0}, {"",0}, {"wad",0}},
   {"wcprice", "Weighted Close Price", ti_wcprice_start, ti_wcprice, TI_TYPE_OVERLAY, 3, 0, 1, {"high","low","close",0}, {"",0}, {"wcprice",0}},
   {"wilders", "Wilders Smoothing", ti_wilders_start, ti_wilders, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"wilders",0}},
   {"willr", "Williams %R", ti_willr_start, ti_willr, TI_TYPE_INDICATOR, 3, 1, 1, {"high","low","close",0}, {"period",0}, {"willr",0}},
   {"wma", "Weighted Moving Average", ti_wma_start, ti_wma, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"wma",0}},
   {"zlema", "Zero-Lag Exponential Moving Average", ti_zlema_start, ti_zlema, TI_TYPE_OVERLAY, 1, 1, 1, {"real",0}, {"period",0}, {"zlema",0}},
*/

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

func acos(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["acos"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_acos_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_acos(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}

func ad(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["ad"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_ad_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_ad(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}

func add(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["add"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_add_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_add(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}

func adosc(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["adosc"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_adosc_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_adosc(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}

func adx(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["adx"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_adx_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_adx(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}

func adxr(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["adxr"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_adxr_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_adxr(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}

func ao(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["ao"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_ao_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_ao(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}

func apo(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["apo"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_apo_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_apo(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}

func aroon(size int, inputs [][]float64, options []float64) (int, [][]float64, error) {

	castSize, castInputs, castOptions := castDoParams(size, inputs, options)

	defer freeC2dDoubleArray(castInputs, len(inputs))
	defer freeCDoubleArray(castOptions)

	var info IndicatorInfo
	var ok bool

	if info, ok = memoizedIndicatorInfo["aroon"]; !ok {
		return 0, [][]float64{}, fmt.Errorf("info hasn't been memoized yet")
	}
	outputs := make([][]float64, info.outputs)

	outputSizeDiff := C.ti_aroon_start(castOptions)

	for i := range outputs {
		outputs[i] = make([]float64, len(inputs[i])-int(outputSizeDiff))
	}

	castOutputs := castToC2dDoubleArray(outputs)

	defer freeC2dDoubleArray(castOutputs, len(outputs))

	doResponse, doError := C.ti_aroon(castSize, castInputs, castOptions, castOutputs)

	extractOutputs(castOutputs, &outputs)

	return int(doResponse), outputs, doError
}
