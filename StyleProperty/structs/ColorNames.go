package structs

const CssColorStringsCount = 9

var CssColorStrings = []string{
	"black",
	"blue",
	"cyan",
	"darkblue",
	"gray",
	"green",
	"purple",
	"red",
	"white",
}

var CssColorRGB = [][3]uint8{
	{0, 0, 0},
	{0, 0, 255},
	{0, 255, 255},
	{0, 0, 139},
	{128, 128, 128},
	{0, 255, 0},
	{128, 0, 128},
	{255, 0, 0},
	{255, 255, 255},
}
