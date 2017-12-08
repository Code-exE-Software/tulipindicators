/*
 * Tulip Indicators
 * https://tulipindicators.org/
 * Copyright (c) 2010-2016 Tulip Charts LLC
 * Lewis Van Winkle (LV@tulipcharts.org)
 *
 * This file is part of Tulip Indicators.
 *
 * Tulip Indicators is free software: you can redistribute it and/or modify it
 * under the terms of the GNU Lesser General Public License as published by the
 * Free Software Foundation, either version 3 of the License, or (at your
 * option) any later version.
 *
 * Tulip Indicators is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Lesser General Public License
 * for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Tulip Indicators.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

#include "../indicators.h"
#include "truerange.h"



int ti_atr_start(TI_REAL const *options) {
    return (int)options[0]-1;
}


int ti_atr(int size, TI_REAL const * const *inputs, TI_REAL const *options, TI_REAL * const *outputs) {
    const TI_REAL *high = inputs[0];
    const TI_REAL *low = inputs[1];
    const TI_REAL *close = inputs[2];

    const int period = (int)options[0];

    TI_REAL *output = outputs[0];

    if (period < 1) return TI_INVALID_OPTION;
    if (size <= ti_atr_start(options)) return TI_OKAY;

    const TI_REAL per = 1.0 / ((TI_REAL)period);

    TI_REAL sum = 0;
    TI_REAL truerange;

    sum += high[0] - low[0];
    int i;
    for (i = 1; i < period; ++i) {
        CALC_TRUERANGE();
        sum += truerange;
    }


    TI_REAL val = sum / period;
    *output++ = val;

    for (i = period; i < size; ++i) {
        CALC_TRUERANGE();
        val = (truerange-val) * per + val;
        *output++ = val;
    }


    assert(output - outputs[0] == size - ti_atr_start(options));
    return TI_OKAY;
}
