package structs

import (
	"gezgin_web_engine/StyleProperty/enums"
)

type Font struct {
	FontSizeInherit   bool
	FontSizeValueType enums.CssFontSizeType
	FontSizeValue     int
}
