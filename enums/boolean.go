package enums

type (
	YesNo int
)

const (
	YesNo_EMPTY YesNo = iota
	YesNo_N
	YesNo_Y
)

var booleanStrMap = map[YesNo]string{
	YesNo_EMPTY: "",
	YesNo_N:     "N",
	YesNo_Y:     "Y",
}

func (c YesNo) String() string {
	if s, found := booleanStrMap[c]; found {
		return s
	}
	return ""
}
