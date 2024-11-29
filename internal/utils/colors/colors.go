package colors


var (
	ColorBlack       = "30"
	ColorRed         = "31"
	ColorGreen       = "32"
	ColorYellow      = "33"
	ColorBlue        = "34"
	ColorMagenta     = "35"
	ColorCyan        = "36"
	ColorWhite       = "37"

	ColorBlackBold   = "1;" + ColorBlack
	ColorRedBold     = "1;" + ColorRed
	ColorGreenBold   = "1;" + ColorGreen
	ColorYellowBold  = "1;" + ColorYellow
	ColorBlueBold    = "1;" + ColorBlue
	ColorMagentaBold = "1;" + ColorMagenta
	ColorCyanBold    = "1;" + ColorCyan
	ColorWhiteBold   = "1;" + ColorWhite

	BgBlack          = "40"
	BgRed            = "41"
	BgGreen          = "42"
	BgYellow         = "43"
	BgBlue           = "44"
	BgMagenta        = "45"
	BgCyan           = "46"
	BgWhite          = "47"
)

func ColorString(str string, color string) string {
	return "\x1b[" + color + "m" + str + "\x1b[0m"
}