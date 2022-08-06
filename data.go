package main

type CsEnumType string

const (
	CsEnumTypeEmpty  CsEnumType = ""
	CsEnumTypeUint   CsEnumType = "uint"
	CsEnumTypeInt    CsEnumType = "int"
	CsEnumTypeByte   CsEnumType = "byte"
	CsEnumTypeSbyte  CsEnumType = "sbyte"
	CsEnumTypeUshort CsEnumType = "ushort"
	CsEnumTypeShort  CsEnumType = "short"
	CsEnumTypeUlong  CsEnumType = "ulong"
	CsEnumTypeLong   CsEnumType = "long"
)

type GoEnumType string

const (
	GoEnumTypeUint32 GoEnumType = "uint32"
	GoEnumTypeInt32  GoEnumType = "int32"
	GoEnumTypeByte   GoEnumType = "byte"
	GoEnumTypeInt8   GoEnumType = "int8"
	GoEnumTypeUint16 GoEnumType = "uint16"
	GoEnumTypeInt16  GoEnumType = "int16"
	GoEnumTypeUint64 GoEnumType = "uint64"
	GoEnumTypeInt64  GoEnumType = "int64"
)

type Enum struct {
	Name          string
	CsType        CsEnumType
	GoType        GoEnumType
	KeyValuePairs []KeyValuePair
}

type KeyValuePair struct {
	Name  string
	Value string
}
