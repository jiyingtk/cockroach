// Code generated by "stringer -type=Type"; DO NOT EDIT.

package encoding

import "fmt"

const _Type_name = "UnknownNullNotNullIntFloatDecimalBytesBytesDescTimeDurationTrueFalseUUIDArrayIPAddrSentinelType"

var _Type_index = [...]uint8{0, 7, 11, 18, 21, 26, 33, 38, 47, 51, 59, 63, 68, 72, 77, 83, 95}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
