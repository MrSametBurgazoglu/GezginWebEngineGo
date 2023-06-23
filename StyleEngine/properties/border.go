package properties

import (
	"gezgin_web_engine/StyleEngine"
	"gezgin_web_engine/StyleEngine/enums"
	structs2 "gezgin_web_engine/StyleEngine/structs"
	"gezgin_web_engine/utils"
	"log"
	"strconv"
	"strings"
)

const BorderStyleStringCount = 10
const BorderWidthStringCount = 3
const BorderImageRepeatStringCount = 4

var borderStyleStrings = []string{
	"dashed",
	"dotted",
	"double",
	"groove",
	"hidden",
	"inset",
	"none",
	"outset",
	"ridge",
	"solid",
}

var borderWidthStrings = []string{
	"medium",
	"thick",
	"thin",
}

var border_image_repeat_strings = []string{
	"repeat",
	"round",
	"space",
	"stretch",
}

func setBorderTopStyle(border *structs2.Border, value string) {
	index := utils.IndexFounder(borderStyleStrings, value, BorderStyleStringCount)
	if index != -1 {
		border.BorderStyle.BorderStyleTypeTop = enums.CssBorderStyleType(index)
	} else {
		border.BorderStyle.BorderStyleTypeTop = enums.CSS_BORDER_STYLE_TYPE_NONE
	}
}

func setBorderBottomStyle(border *structs2.Border, value string) {
	index := utils.IndexFounder(borderStyleStrings, value, BorderStyleStringCount)
	if index != -1 {
		border.BorderStyle.BorderStyleTypeBottom = enums.CssBorderStyleType(index)
	} else {
		border.BorderStyle.BorderStyleTypeBottom = enums.CSS_BORDER_STYLE_TYPE_NONE
	}
}

func setBorderLeftStyle(border *structs2.Border, value string) {
	index := utils.IndexFounder(borderStyleStrings, value, BorderStyleStringCount)
	if index != -1 {
		border.BorderStyle.BorderStyleTypeLeft = enums.CssBorderStyleType(index)
	} else {
		border.BorderStyle.BorderStyleTypeLeft = enums.CSS_BORDER_STYLE_TYPE_NONE
	}
}

func setBorderRightStyle(border *structs2.Border, value string) {
	index := utils.IndexFounder(borderStyleStrings, value, BorderStyleStringCount)
	if index != -1 {
		border.BorderStyle.BorderStyleTypeRight = enums.CssBorderStyleType(index)
	} else {
		border.BorderStyle.BorderStyleTypeRight = enums.CSS_BORDER_STYLE_TYPE_NONE
	}
}

func setBorderStyle(border *structs2.Border, value string) {
	values := strings.Split(value, " ")
	index := len(values)
	switch index {
	case 1:
		setBorderBottomStyle(border, values[0])
		setBorderTopStyle(border, values[0])
		setBorderLeftStyle(border, values[0])
		setBorderRightStyle(border, values[0])
	case 2:
		setBorderBottomStyle(border, values[0])
		setBorderTopStyle(border, values[0])
		setBorderLeftStyle(border, values[1])
		setBorderRightStyle(border, values[1])
	case 3:
		setBorderBottomStyle(border, values[2])
		setBorderTopStyle(border, values[0])
		setBorderLeftStyle(border, values[1])
		setBorderRightStyle(border, values[1])
	case 4:
		setBorderBottomStyle(border, values[2])
		setBorderTopStyle(border, values[0])
		setBorderLeftStyle(border, values[3])
		setBorderRightStyle(border, values[1])
	default:
		break
	}
}

func setBorderTopWidth(border *structs2.Border, value string) {
	index := utils.IndexFounder(borderWidthStrings, value, BorderWidthStringCount)
	if index != -1 {
		border.BorderWidth.TopBorder.BorderLineType = enums.CssBorderlineType(index)
	} else {
		firstValue, err := strconv.Atoi(value)
		if err == nil {
			if strings.HasSuffix(value, "px") {
				border.BorderWidth.TopBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_LENGTH
				border.BorderWidth.TopBorder.Value = firstValue
			} else {
				border.BorderWidth.TopBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
			}
		} else {
			border.BorderWidth.TopBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
		}
	}
}

func setBorderBottomWidth(border *structs2.Border, value string) {
	index := utils.IndexFounder(borderWidthStrings, value, BorderWidthStringCount)
	if index != -1 {
		border.BorderWidth.BottomBorder.BorderLineType = enums.CssBorderlineType(index)
	} else {
		firstValue, err := strconv.Atoi(value)
		if err == nil {
			if strings.HasSuffix(value, "px") {
				border.BorderWidth.BottomBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_LENGTH
				border.BorderWidth.BottomBorder.Value = firstValue
			} else {
				border.BorderWidth.BottomBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
			}
		} else {
			border.BorderWidth.BottomBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
		}
	}
}

func setBorderLeftWidth(border *structs2.Border, value string) {
	index := utils.IndexFounder(borderWidthStrings, value, BorderWidthStringCount)
	if index != -1 {
		border.BorderWidth.LeftBorder.BorderLineType = enums.CssBorderlineType(index)
	} else {
		firstValue, err := strconv.Atoi(value)
		if err == nil {
			if strings.HasSuffix(value, "px") {
				border.BorderWidth.LeftBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_LENGTH
				border.BorderWidth.LeftBorder.Value = firstValue
			} else {
				border.BorderWidth.LeftBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
			}
		} else {
			border.BorderWidth.LeftBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
		}
	}
}

func setBorderRightWidth(border *structs2.Border, value string) {
	index := utils.IndexFounder(borderWidthStrings, value, BorderWidthStringCount)
	if index != -1 {
		border.BorderWidth.RightBorder.BorderLineType = enums.CssBorderlineType(index)
	} else {
		firstValue, err := strconv.Atoi(value)
		if err == nil {
			if strings.HasSuffix(value, "px") {
				border.BorderWidth.RightBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_LENGTH
				border.BorderWidth.RightBorder.Value = firstValue
			} else {
				border.BorderWidth.RightBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
			}
		} else {
			border.BorderWidth.RightBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
		}
	}
}

func setBorderWidth(border *structs2.Border, value string) {
	values := strings.Split(value, " ")
	switch len(values) {
	case 1:
		setBorderBottomWidth(border, values[0])
		setBorderTopWidth(border, values[0])
		setBorderLeftWidth(border, values[0])
		setBorderRightWidth(border, values[0])
	case 2:
		setBorderBottomWidth(border, values[0])
		setBorderTopWidth(border, values[0])
		setBorderLeftWidth(border, values[1])
		setBorderRightWidth(border, values[1])
	case 3:
		setBorderBottomWidth(border, values[2])
		setBorderTopWidth(border, values[0])
		setBorderLeftWidth(border, values[1])
		setBorderRightWidth(border, values[1])
	case 4:
		setBorderBottomWidth(border, values[2])
		setBorderTopWidth(border, values[0])
		setBorderLeftWidth(border, values[3])
		setBorderRightWidth(border, values[1])
	default:
		break
	}
}

func setBorderTopColor(border *structs2.Border, value string) {
	border.BorderColor.TopBorderColor.SetColor(value)
}

func setBorderLeftColor(border *structs2.Border, value string) {
	border.BorderColor.LeftBorderColor.SetColor(value)
}

func setBorderRightColor(border *structs2.Border, value string) {
	border.BorderColor.RightBorderColor.SetColor(value)
}

func setBorderBottomColor(border *structs2.Border, value string) {
	border.BorderColor.BottomBorderColor.SetColor(value)
}

func setBorderColor(border *structs2.Border, value string) {
	values := strings.Split(value, " ")
	switch len(values) {
	case 1:
		setBorderBottomColor(border, values[0])
		setBorderTopColor(border, values[0])
		setBorderLeftColor(border, values[0])
		setBorderRightColor(border, values[0])
	case 2:
		setBorderBottomColor(border, values[0])
		setBorderTopColor(border, values[0])
		setBorderLeftColor(border, values[1])
		setBorderRightColor(border, values[1])
	case 3:
		setBorderBottomColor(border, values[2])
		setBorderTopColor(border, values[0])
		setBorderLeftColor(border, values[1])
		setBorderRightColor(border, values[1])
	case 4:
		setBorderBottomColor(border, values[2])
		setBorderTopColor(border, values[0])
		setBorderLeftColor(border, values[3])
		setBorderRightColor(border, values[1])
	default:
		break
	}
}

func setBorder(border *structs2.Border, value string) {
	values := strings.Split(value, " ")
	switch len(values) {
	case 1:
		setBorderStyle(border, values[0])
	case 2:
		setBorderWidth(border, values[0])
		setBorderStyle(border, values[1])
	case 3:
		setBorderWidth(border, values[0])
		setBorderStyle(border, values[1])
		setBorderColor(border, values[2])
	default:
		break
	}
}

func setBorderBottomLeftRadius(borderRadius *structs2.BorderRadius, value string) {
	number, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	borderRadius.BorderBottomLeftRadiusValue = number
	if strings.HasSuffix(value, "%") {
		borderRadius.BorderBottomLeftRadiusType = enums.CSS_BORDER_RADIUS_TYPE_PERCENTAGE
	} else {
		borderRadius.BorderBottomLeftRadiusType = enums.CSS_BORDER_RADIUS_TYPE_LENGTH
	}
}

func setBorderBottomRightRadius(borderRadius *structs2.BorderRadius, value string) {
	number, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	borderRadius.BorderBottomRightRadiusValue = number
	if strings.HasSuffix(value, "%") {
		borderRadius.BorderBottomRightRadiusType = enums.CSS_BORDER_RADIUS_TYPE_PERCENTAGE
	} else {
		borderRadius.BorderBottomRightRadiusType = enums.CSS_BORDER_RADIUS_TYPE_LENGTH
	}
}

func setBorderTopLeftRadius(borderRadius *structs2.BorderRadius, value string) {
	number, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	borderRadius.BorderTopLeftRadiusValue = number
	if strings.HasSuffix(value, "%") {
		borderRadius.BorderTopLeftRadiusType = enums.CSS_BORDER_RADIUS_TYPE_PERCENTAGE
	} else {
		borderRadius.BorderTopLeftRadiusType = enums.CSS_BORDER_RADIUS_TYPE_LENGTH
	}
}

func setBorderTopRightRadius(borderRadius *structs2.BorderRadius, value string) {
	number, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	borderRadius.BorderTopRightRadiusValue = number
	if strings.HasSuffix(value, "%") {
		borderRadius.BorderTopRightRadiusType = enums.CSS_BORDER_RADIUS_TYPE_PERCENTAGE
	} else {
		borderRadius.BorderTopRightRadiusType = enums.CSS_BORDER_RADIUS_TYPE_LENGTH
	}
}

func setBorderRadius(borderRadius *structs2.BorderRadius, value string) {
	values := strings.Split(value, " ")
	switch len(values) {
	case 1:
		setBorderBottomLeftRadius(borderRadius, values[0])
		setBorderBottomRightRadius(borderRadius, values[0])
		setBorderTopLeftRadius(borderRadius, values[0])
		setBorderTopRightRadius(borderRadius, values[0])
	case 2:
		setBorderBottomLeftRadius(borderRadius, values[1])
		setBorderBottomRightRadius(borderRadius, values[0])
		setBorderTopLeftRadius(borderRadius, values[0])
		setBorderTopRightRadius(borderRadius, values[1])
	case 3:
		setBorderBottomLeftRadius(borderRadius, values[1])
		setBorderBottomRightRadius(borderRadius, values[2])
		setBorderTopLeftRadius(borderRadius, values[0])
		setBorderTopRightRadius(borderRadius, values[1])
	case 4:
		setBorderBottomLeftRadius(borderRadius, values[3])
		setBorderBottomRightRadius(borderRadius, values[2])
		setBorderTopLeftRadius(borderRadius, values[0])
		setBorderTopRightRadius(borderRadius, values[1])
	default:
		break
	}
}

func setBorderSpacing(borderSpacing *structs2.BorderSpacing, value string) {
	values := strings.Split(value, " ")
	switch len(values) {
	case 1:
		hSpacing, err := strconv.Atoi(values[0])
		if err == nil {
			borderSpacing.HSpacing = hSpacing
			borderSpacing.VSpacing = hSpacing
		}
	case 2:
		hSpacing, err := strconv.Atoi(values[0])
		vSpacing, err2 := strconv.Atoi(values[1])
		if err == nil && err2 == nil {
			borderSpacing.HSpacing = hSpacing
			borderSpacing.VSpacing = vSpacing
		}
	default:
		break
	}
}

func setBorderBottom(border *structs2.Border, value string) {
	values := strings.Split(value, " ")
	index := len(values)

	switch index {
	case 1:
		setBorderBottomStyle(border, values[0])
	case 2:
		setBorderBottomStyle(border, values[1])
		setBorderBottomWidth(border, values[0])
	case 3:
		setBorderBottomStyle(border, values[1])
		setBorderBottomWidth(border, values[0])
		setBorderBottomColor(border, values[2])
	default:
		break
	}
}

func setBorderTop(border *structs2.Border, value string) {
	values := strings.Split(value, " ")
	index := len(values)

	switch index {
	case 1:
		setBorderTopStyle(border, values[0])
	case 2:
		setBorderTopStyle(border, values[1])
		setBorderTopWidth(border, values[0])
	case 3:
		setBorderTopStyle(border, values[1])
		setBorderTopWidth(border, values[0])
		setBorderTopColor(border, values[2])
	default:
		break
	}
}

func setBorderLeft(border *structs2.Border, value string) {
	values := strings.Split(value, " ")
	index := len(values)

	switch index {
	case 1:
		setBorderLeftStyle(border, values[0])
	case 2:
		setBorderLeftStyle(border, values[1])
		setBorderLeftWidth(border, values[0])
	case 3:
		setBorderLeftStyle(border, values[1])
		setBorderLeftWidth(border, values[0])
		setBorderLeftColor(border, values[2])
	default:
		break
	}
}

func setBorderRight(border *structs2.Border, value string) {
	values := strings.Split(value, " ")
	index := len(values)

	switch index {
	case 1:
		setBorderRightStyle(border, values[0])
	case 2:
		setBorderRightStyle(border, values[1])
		setBorderRightWidth(border, values[0])
	case 3:
		setBorderRightStyle(border, values[1])
		setBorderRightWidth(border, values[0])
		setBorderRightColor(border, values[2])
	default:
		break
	}
}

func setBorderCollapse(cssProperties *StyleEngine.StyleProperty, value string) {
	if value == "collapse" {
		cssProperties.BorderCollapseType = enums.CSS_BORDER_COLLAPSE_TYPE_COLLAPSE
	} else {
		cssProperties.BorderCollapseType = enums.CSS_BORDER_COLLAPSE_TYPE_SEPARATE
	}
}

func BorderPropertySetValue(cssProperties *StyleEngine.StyleProperty, value string) {
	if value == "inherit" {
		cssProperties.BorderInherit = true
	} else {
		cssProperties.BorderInherit = false
		if cssProperties.Border == nil || cssProperties.BorderInherit {
			cssProperties.Border = new(structs2.Border)
			cssProperties.Border.BorderStyle = new(structs2.BorderStyle)
			cssProperties.Border.BorderWidth = new(structs2.BorderWidth)
			cssProperties.Border.BorderWidth.TopBorder = new(structs2.BorderLineWidth)
			cssProperties.Border.BorderWidth.BottomBorder = new(structs2.BorderLineWidth)
			cssProperties.Border.BorderWidth.LeftBorder = new(structs2.BorderLineWidth)
			cssProperties.Border.BorderWidth.RightBorder = new(structs2.BorderLineWidth)
			cssProperties.Border.BorderColor = new(structs2.BorderColor)
			cssProperties.Border.BorderColor.TopBorderColor = new(structs2.ColorRGBA)
			cssProperties.Border.BorderColor.BottomBorderColor = new(structs2.ColorRGBA)
			cssProperties.Border.BorderColor.LeftBorderColor = new(structs2.ColorRGBA)
			cssProperties.Border.BorderColor.RightBorderColor = new(structs2.ColorRGBA)
		} else {
			if cssProperties.Border.BorderStyle == nil || cssProperties.Border.BorderStyleInherit {
				cssProperties.Border.BorderStyle = new(structs2.BorderStyle)
			}
			if cssProperties.Border.BorderWidth == nil || cssProperties.Border.BorderWidthInherit {
				cssProperties.Border.BorderWidth = new(structs2.BorderWidth)
				cssProperties.Border.BorderWidth.TopBorder = new(structs2.BorderLineWidth)
				cssProperties.Border.BorderWidth.BottomBorder = new(structs2.BorderLineWidth)
				cssProperties.Border.BorderWidth.RightBorder = new(structs2.BorderLineWidth)
				cssProperties.Border.BorderWidth.LeftBorder = new(structs2.BorderLineWidth)
			}
			if cssProperties.Border.BorderColor == nil || cssProperties.Border.BorderColorInherit {
				cssProperties.Border.BorderColor = new(structs2.BorderColor)
				cssProperties.Border.BorderColor.TopBorderColor = new(structs2.ColorRGBA)
				cssProperties.Border.BorderColor.BottomBorderColor = new(structs2.ColorRGBA)
				cssProperties.Border.BorderColor.LeftBorderColor = new(structs2.ColorRGBA)
				cssProperties.Border.BorderColor.RightBorderColor = new(structs2.ColorRGBA)
			}
		}
		if value == "initial" {
			cssProperties.Border.BorderStyle.BorderStyleTypeTop = enums.CSS_BORDER_STYLE_TYPE_NONE
			cssProperties.Border.BorderStyle.BorderStyleTypeBottom = enums.CSS_BORDER_STYLE_TYPE_NONE
			cssProperties.Border.BorderStyle.BorderStyleTypeLeft = enums.CSS_BORDER_STYLE_TYPE_NONE
			cssProperties.Border.BorderStyle.BorderStyleTypeRight = enums.CSS_BORDER_STYLE_TYPE_NONE
			cssProperties.Border.BorderWidth.TopBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
			cssProperties.Border.BorderWidth.BottomBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
			cssProperties.Border.BorderWidth.LeftBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
			cssProperties.Border.BorderWidth.RightBorder.BorderLineType = enums.CSS_BORDER_LINE_TYPE_MEDIUM
			cssProperties.Border.BorderColor.TopBorderColor.SyncColor(cssProperties.Color)
			cssProperties.Border.BorderColor.BottomBorderColor.SyncColor(cssProperties.Color)
			cssProperties.Border.BorderColor.LeftBorderColor.SyncColor(cssProperties.Color)
			cssProperties.Border.BorderColor.RightBorderColor.SyncColor(cssProperties.Color)
		} else {
			setBorder(cssProperties.Border, value)
		}
	}
}
