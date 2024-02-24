package enums

type (
	LandmarkCheck  int
	LandmarkChecks []LandmarkCheck
)

const (
	landMarkCheck_kadPengenalan LandmarkCheck = iota
	landMarkCheck_mykadLogo
	landMarkCheck_flagLogo
	landMarkCheck_mscLogo
	landMarkCheck_ghostFace
	landMarkCheck_hibiscusLogo
	landMarkCheck_coatOfArm
	landMarkCheck_twinTower
)

var landmarkCheckMap = map[LandmarkCheck]string{
	landMarkCheck_kadPengenalan: "kadPengenalan",
	landMarkCheck_mykadLogo:     "mykadLogo",
	landMarkCheck_flagLogo:      "flagLogo",
	landMarkCheck_mscLogo:       "mscLogo",
	landMarkCheck_ghostFace:     "ghostFace",
	landMarkCheck_hibiscusLogo:  "hibiscusLogo",
	landMarkCheck_coatOfArm:     "coatOfArm",
	landMarkCheck_twinTower:     "twinTower",
}

func (c LandmarkCheck) String() string {
	if s, found := landmarkCheckMap[c]; found {
		return s
	}
	return ""
}

func (c LandmarkChecks) ToMap() map[string]any {
	return listStringerToMap(c)
}
