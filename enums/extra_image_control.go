package enums

type (
	ExtraImageControl     int
	ExtraImageControlList []ExtraImageControl
)

const (
	ExtraImageControl_FACE_EYE_CLOSE ExtraImageControl = iota
	ExtraImageControl_DOC_FRONT_ANGLE
	ExtraImageControl_DOC_FRONT_FLASH
	ExtraImageControl_DOC_BACK_ANGLE
	ExtraImageControl_DOC_BACK_FLASH
	ExtraImageControl_CROPPED_FACE
)

var extraImageControlMap = map[ExtraImageControl]string{
	ExtraImageControl_FACE_EYE_CLOSE:  "FACE_EYE_CLOSE",
	ExtraImageControl_DOC_FRONT_ANGLE: "DOC_FRONT_ANGLE",
	ExtraImageControl_DOC_FRONT_FLASH: "DOC_FRONT_FLASH",
	ExtraImageControl_DOC_BACK_ANGLE:  "DOC_BACK_ANGLE",
	ExtraImageControl_DOC_BACK_FLASH:  "DOC_BACK_FLASH",
	ExtraImageControl_CROPPED_FACE:    "CROPPED_FACE",
}

func (c ExtraImageControl) String() string {
	if s, found := extraImageControlMap[c]; found {
		return s
	}
	return ""
}

func (c ExtraImageControlList) ToStringList() []string {
	return listStringerToListString(c)
}
