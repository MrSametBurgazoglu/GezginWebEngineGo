package structs

import (
	"gezgin_web_engine/StyleEngine/enums"
)

type Font struct {
	FontSizeInherit   bool
	FontSizeValueType enums.CssFontSizeType
	FontSizeValue     int
}
