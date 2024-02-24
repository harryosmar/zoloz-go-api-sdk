package enums

type (
	ServiceLevel int
)

const (
	ServiceLevel_EMPTY ServiceLevel = iota
	ServiceLevel_REALID0001
	ServiceLevel_REALID0002
	ServiceLevel_REALID0003
	ServiceLevel_REALID0004
	ServiceLevel_REALID0009
	ServiceLevel_REALID00011
	ServiceLevel_REALID00012
)

var serviceLevelMap = map[ServiceLevel]string{
	ServiceLevel_EMPTY:       "",
	ServiceLevel_REALID0001:  "REALID0001",
	ServiceLevel_REALID0002:  "REALID0002",
	ServiceLevel_REALID0003:  "REALID0003",
	ServiceLevel_REALID0004:  "REALID0004",
	ServiceLevel_REALID0009:  "REALID0009",
	ServiceLevel_REALID00011: "REALID00011",
	ServiceLevel_REALID00012: "REALID00012",
}

func (c ServiceLevel) String() string {
	if s, found := serviceLevelMap[c]; found {
		return s
	}
	return ""
}
