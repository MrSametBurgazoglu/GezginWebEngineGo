package structs

const CssColorStringsCount = 5

var CssColorStrings = []string{
	"black",
	"blue",
	"green",
	"red",
	"white",
}

var CssColorRGB = [][3]uint8{
	{0, 0, 0},
	{0, 0, 255},
	{0, 255, 0},
	{255, 0, 0},
	{255, 255, 255},
}
