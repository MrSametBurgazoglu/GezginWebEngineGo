package structs

type ColorRGBA struct {
	alpha uint8
	red   uint8
	green uint8
	blue  uint8
}

func GetColorByName(color *ColorRGBA, name string) bool {
	/*
		index := IndexFounder(css_color_strings, name, CSS_COLOR_NAME_STRING_COUNT);
		    if (index != -1){
		        int* values = css_color_name_rgb[index]
		        color_struct.alpha = 0
		        color_struct.red = values[0]
		        color_struct.green = values[1]
		        color_struct.blue = values[2]
		        return true
		    }
		    return false
	*/
	return true
}

func (receiver *ColorRGBA) SetColor(value string) bool {
	println(value)
	if GetColorByName(receiver, value) {
		return true
	}
	return false
}
