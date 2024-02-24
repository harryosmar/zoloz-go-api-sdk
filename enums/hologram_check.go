package enums

type (
	HologramCheck  int
	HologramChecks []HologramCheck
)

const (
	hologramCheck_hologram HologramCheck = iota
)

var hologramCheckMap = map[HologramCheck]string{
	hologramCheck_hologram: "hologram",
}

func (c HologramCheck) String() string {
	if s, found := hologramCheckMap[c]; found {
		return s
	}
	return ""
}

func (c HologramChecks) ToMap() map[string]any {
	return listStringerToMap(c)
}
