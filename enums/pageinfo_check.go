package enums

type (
	PageInfoCheck  int
	PageInfoChecks []PageInfoCheck
)

const (
	pageInfoCheck_id PageInfoCheck = iota
	pageInfoCheck_symbol
	pageInfoCheck_name
)

var pageInfoCheckMap = map[PageInfoCheck]string{
	pageInfoCheck_id:     "id",
	pageInfoCheck_symbol: "symbol",
	pageInfoCheck_name:   "name",
}

func (c PageInfoCheck) String() string {
	if s, found := pageInfoCheckMap[c]; found {
		return s
	}
	return ""
}

func (c PageInfoChecks) ToMap() map[string]any {
	return listStringerToMap(c)
}
