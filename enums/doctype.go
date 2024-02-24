package enums

type (
	DocType int
)

const (
	DocType_IndonesiaEKtp DocType = iota
	DocType_IndonesiaTaxID
)

var docTypeMap = map[DocType]string{
	DocType_IndonesiaEKtp:  "00620000001",
	DocType_IndonesiaTaxID: "00620000004",
}

func (c DocType) String() string {
	if s, found := docTypeMap[c]; found {
		return s
	}
	return ""
}
