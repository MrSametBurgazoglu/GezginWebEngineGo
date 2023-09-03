package structs

import (
	"gezgin_web_engine/utils"
	"strconv"
)

type ColorRGBA struct {
	alpha uint8
	red   uint8
	green uint8
	blue  uint8
}

func HueToRGB(v1 float32, v2 float32, vH float32) float32 {
	if vH < 0 {
		vH += 1
	} else if vH > 1 {
		vH -= 1
	}
	switch {
	case 6*vH < 1:
		return v1 + (v2-v1)*6*vH
	case 2*vH < 1:
		return v2
	case 3*vH < 2:
		return v1 + (v2-v1)*((2.0/3)-vH)*6
	}
	return v1
}

func (receiver *ColorRGBA) GetColorByRGBA() (uint8, uint8, uint8, uint8) {
	if receiver == nil {
		return 255, 0, 0, 0
	}
	return receiver.alpha, receiver.red, receiver.green, receiver.blue
}

func (receiver *ColorRGBA) SetColorByRGBA(red uint8, green uint8, blue uint8, alpha uint8) {
	receiver.alpha = alpha
	receiver.red = red
	receiver.green = green
	receiver.blue = blue
}

func (receiver *ColorRGBA) SetColorByRGB(red uint8, green uint8, blue uint8) {
	receiver.alpha = 255
	receiver.red = red
	receiver.green = green
	receiver.blue = blue
}

func (receiver *ColorRGBA) SetColorByHSL(h float32, s float32, l float32) {
	receiver.alpha = 0
	if s == 0 {
		result := uint8(l * 255)
		receiver.red = result
		receiver.green = result
		receiver.blue = result
	} else {
		hue := h / 360
		var v2 float32
		if l < 0.5 {
			v2 = l * (1 + s)
		} else {
			v2 = l + s - (l * s)
		}
		v1 := 2*l - v2
		receiver.red = uint8(255 * HueToRGB(v1, v2, hue+(1.0/3)))
		receiver.green = uint8(255 * HueToRGB(v1, v2, hue))
		receiver.blue = uint8(255 * HueToRGB(v1, v2, hue-(1.0/3)))
	}
	receiver.alpha = 0
}

func (receiver *ColorRGBA) SetColorByHSLA(h float32, s float32, l float32, alpha uint8) {
	receiver.alpha = alpha
	receiver.SetColorByHSL(h, s, l)
}

func (receiver *ColorRGBA) SetColorByHex(value string) bool {
	if value[0] == '#' {
		values, err := strconv.ParseUint(value[1:], 16, 32)
		//n, err := fmt.Sscanf(value, "#%02x%02x%02x", &r, &g, &b)
		if err != nil {
			return true
		}
		receiver.red = uint8(values >> 16)
		receiver.green = uint8((values >> 8) & 0xFF)
		receiver.blue = uint8(values & 0xFF)
	}
	return false
}

func (receiver *ColorRGBA) SetColor(value string) bool {
	if receiver.SetColorByName(value) {
		return true
	} else if receiver.SetColorByFunction(value) {
		return true
	} else if receiver.SetColorByHex(value) {
		return true
	}
	return false
}

func (receiver *ColorRGBA) SetColorByName(value string) bool {
	if index := utils.IndexFounder(CssColorStrings, value, CssColorStringsCount); index != -1 {
		receiver.alpha = 255
		receiver.red = CssColorRGB[index][0]
		receiver.green = CssColorRGB[index][1]
		receiver.blue = CssColorRGB[index][2]
		return true
	}
	return false
}

func (receiver *ColorRGBA) SetColorByFunction(value string) bool {
	if functionName, functionParameters, ok := utils.ParseFunction(value); ok {
		parameterCount := len(functionParameters)
		if parameterCount == 3 {
			switch functionName {
			case "rgb":
				value1, err := strconv.Atoi(functionParameters[0])
				if err != nil {
					value1 = 0
				}
				value2, err2 := strconv.Atoi(functionParameters[1])
				if err2 != nil {
					value2 = 0
				}
				value3, err3 := strconv.Atoi(functionParameters[2])
				if err3 != nil {
					value3 = 0
				}
				receiver.SetColorByRGB(uint8(value1), uint8(value2), uint8(value3))
				return true
			case "hsl":
				value1, err := strconv.ParseFloat(functionParameters[0], 32)
				if err == nil {
					value1 = 0
				}
				value2, err2 := strconv.ParseFloat(functionParameters[1], 32)
				if err2 == nil {
					value2 = 0
				}
				value3, err3 := strconv.ParseFloat(functionParameters[2], 32)
				if err3 == nil {
					value3 = 0
				}
				receiver.SetColorByHSL(float32(value1), float32(value2), float32(value3))
				return true
			}
		} else if parameterCount == 4 {
			switch functionName {
			case "rgba":
				value1, err := strconv.Atoi(functionParameters[0])
				if err == nil {
					value1 = 0
				}
				value2, err2 := strconv.Atoi(functionParameters[1])
				if err2 == nil {
					value2 = 0
				}
				value3, err3 := strconv.Atoi(functionParameters[2])
				if err3 == nil {
					value3 = 0
				}
				value4, err4 := strconv.Atoi(functionParameters[3])
				if err4 == nil {
					value4 = 0
				}
				receiver.SetColorByRGBA(uint8(value1), uint8(value2), uint8(value3), uint8(value4))
			case "hsla":
				value1, err := strconv.ParseFloat(functionParameters[0], 32)
				if err == nil {
					value1 = 0
				}
				value2, err2 := strconv.ParseFloat(functionParameters[1], 32)
				if err2 == nil {
					value2 = 0
				}
				value3, err3 := strconv.ParseFloat(functionParameters[2], 32)
				if err3 == nil {
					value3 = 0
				}
				value4, err4 := strconv.Atoi(functionParameters[3])
				if err4 == nil {
					value4 = 0
				}
				receiver.SetColorByHSLA(float32(value1), float32(value2), float32(value3), uint8(value4))
			}
		}
	}
	return false
}

func (receiver *ColorRGBA) SyncColor(source *ColorRGBA) {
	receiver.alpha = source.alpha
	receiver.red = source.red
	receiver.green = source.green
	receiver.blue = source.blue
}
