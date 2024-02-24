package enums

type (
	Level int
)

const (
	Level_EMPTY    Level = iota
	Level_STANDARD Level = iota
	Level_CLOSED
	Level_LOOSE
	Level_STRICT
)

var levelMap = map[Level]string{
	Level_EMPTY:    "",
	Level_STANDARD: "STANDARD",
	Level_CLOSED:   "CLOSED",
	Level_LOOSE:    "LOOSE",
	Level_STRICT:   "STRICT",
}

func (c Level) String() string {
	if s, found := levelMap[c]; found {
		return s
	}
	return ""
}
