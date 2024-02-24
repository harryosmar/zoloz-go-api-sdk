package enums

type (
	DocUiType        int
)

const (
	DocUiType_1 DocUiType = iota
	DocUiType_2
	DocUiType_3
	DocUiType_4
	DocUiType_5
)

var docUiTypeMap = map[DocUiType]string{
	DocUiType_1: "1",
	DocUiType_2: "2",
	DocUiType_3: "3",
	DocUiType_4: "4",
	DocUiType_5: "5",
}

func (c DocUiType) String() string {
	if s, found := docUiTypeMap[c]; found {
		return s
	}
	return ""
}
