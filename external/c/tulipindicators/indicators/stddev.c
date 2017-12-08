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


int ti_stddev_start(TI_REAL const *options) {
    return (int)options[0]-1;
}


int ti_stddev(int size, TI_REAL const * const *inputs, TI_REAL const *options, TI_REAL * const *outputs) {
    const TI_REAL *input = inputs[0];
    const int period = (int)options[0];
    TI_REAL *output = outputs[0];
    const TI_REAL div = 1.0 / period;

    if (period < 1) return TI_INVALID_OPTION;
    if (size <= ti_stddev_start(options)) return TI_OKAY;

    TI_REAL sum = 0;
    TI_REAL sum2 = 0;

    int i;
    for (i = 0; i < period; ++i) {
        sum += input[i];
        sum2 += input[i] * input[i];
    }

    {
        TI_REAL s2s2 = (sum2 * div - (sum * div) * (sum * div));
        if (s2s2 > 0.0) s2s2 = sqrt(s2s2);
        *output++ = s2s2;
    }

    for (i = period; i < size; ++i) {
        sum += input[i];
        sum2 += input[i] * input[i];

        sum -= input[i-period];
        sum2 -= input[i-period] * input[i-period];

        TI_REAL s2s2 = (sum2 * div - (sum * div) * (sum * div));
        if (s2s2 > 0.0) s2s2 = sqrt(s2s2);
        *output++ = s2s2;
    }

    assert(output - outputs[0] == size - ti_stddev_start(options));
    return TI_OKAY;
}
