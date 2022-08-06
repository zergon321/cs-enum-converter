package main

func EnumTypeCsToGo(csType CsEnumType) GoEnumType {
	switch csType {
	case CsEnumTypeEmpty, CsEnumTypeInt:
		return GoEnumTypeInt32

	case CsEnumTypeUint:
		return GoEnumTypeUint32

	case CsEnumTypeByte:
		return GoEnumTypeByte

	case CsEnumTypeSbyte:
		return GoEnumTypeInt8

	case CsEnumTypeShort:
		return GoEnumTypeInt16

	case CsEnumTypeUshort:
		return GoEnumTypeUint16

	case CsEnumTypeLong:
		return GoEnumTypeInt64

	case CsEnumTypeUlong:
		return GoEnumTypeUint64
	}

	return ""
}
