package enums

type (
	FlowType int
)

const (
	FlowType_Empty FlowType = iota
	FlowType_REALIDLITE_KYC
	FlowType_H5_REALIDLITE_KYC
)

var flowTypeMap = map[FlowType]string{
	FlowType_Empty:             "",
	FlowType_REALIDLITE_KYC:    "REALIDLITE_KYC",
	FlowType_H5_REALIDLITE_KYC: "H5_REALIDLITE_KYC",
}

func (c FlowType) String() string {
	if s, found := flowTypeMap[c]; found {
		return s
	}
	return ""
}
