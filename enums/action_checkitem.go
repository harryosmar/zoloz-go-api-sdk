package enums

type (
	ActionCheckItem  int
	ActionsCheckItem []ActionCheckItem
)

const (
	ActionCheckItem_FACEBLINK ActionCheckItem = iota
	ActionCheckItem_MOUTHOPEN
	ActionCheckItem_HEADSHAKE
	ActionCheckItem_HEADLOWER
	ActionCheckItem_HEADRAISE
)

var actionCheckItemMap = map[ActionCheckItem]string{
	ActionCheckItem_FACEBLINK: "FACEBLINK",
	ActionCheckItem_MOUTHOPEN: "MOUTHOPEN",
	ActionCheckItem_HEADSHAKE: "HEADSHAKE",
	ActionCheckItem_HEADLOWER: "HEADLOWER",
	ActionCheckItem_HEADRAISE: "HEADRAISE",
}

func (c ActionCheckItem) String() string {
	if s, found := actionCheckItemMap[c]; found {
		return s
	}
	return ""
}

func (c ActionsCheckItem) ToStrings() []string {
	return listStringerToListString(c)
}
