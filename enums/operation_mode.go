package enums

type (
	OperationMode int
)

const (
	OperationMode_EMPTY OperationMode = iota
	OperationMode_STANDARD
	OperationMode_CLOSED
	OperationMode_LOOSE
	OperationMode_STRICT
	OperationMode_STANDARD_VC_CLOSED
	OperationMode_STANDARD_IDN_CLOSED
	OperationMode_STANDARD_VC_IDN_CLOSED
)

var operationModeMap = map[OperationMode]string{
	OperationMode_EMPTY:                  "",
	OperationMode_STANDARD:               "STANDARD",
	OperationMode_CLOSED:                 "CLOSED",
	OperationMode_LOOSE:                  "LOOSE",
	OperationMode_STRICT:                 "STRICT",
	OperationMode_STANDARD_VC_CLOSED:     "STANDARD_VC_CLOSED",
	OperationMode_STANDARD_IDN_CLOSED:    "STANDARD_IDN_CLOSED",
	OperationMode_STANDARD_VC_IDN_CLOSED: "STANDARD_VC_IDN_CLOSED",
}

func (c OperationMode) String() string {
	if s, found := operationModeMap[c]; found {
		return s
	}
	return ""
}
