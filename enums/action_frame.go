package enums

type (
	ActionFrame  int
	ActionsFrame []ActionFrame
)

const (
	ActionFrame_EYECLOSE ActionFrame = iota
)

var actionFrameMap = map[ActionFrame]string{
	ActionFrame_EYECLOSE: "EYECLOSE",
}

func (c ActionFrame) String() string {
	if s, found := actionFrameMap[c]; found {
		return s
	}
	return ""
}

func (c ActionsFrame) ToStrings() []string {
	return listStringerToListString(c)
}
