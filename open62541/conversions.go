// SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
//
// SPDX-License-Identifier: MIT

package open62541

import "fmt"
import "C"

func boolToCUshort(b bool) C.ushort {
	if b {
		return C.ushort(1)
	}
	return C.ushort(0)
}

func float64ToBool(b float64) (bool, error) {
	if b == 0 {
		return false, nil
	} else if b == 1 {
		return true, nil
	}

	return false, fmt.Errorf("Failed to convert float64 to bool, value is not [0, 1]: %v", b)
}
