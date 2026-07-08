package colorista

import (
	"fmt"
	"strings"
)

type Colorista struct {
	theme Theme

	enabled bool
}

func NewColorista(t Theme) *Colorista {
	self := &Colorista{
		enabled: true,
	}
	self.SetTheme(t)
	return self
}

func (self *Colorista) SetTheme(theme Theme) {
	switch theme {
	case ThemeAuto:
		self.theme = detectTheme()
	default:
		self.theme = theme
	}
}

type Style string

const reset = "\033[0m"

type RGB struct {
	R, G, B uint8
}

// Text colors
const (
	Black   Style = "30"
	Red     Style = "31"
	Green   Style = "32"
	Yellow  Style = "33"
	Blue    Style = "34"
	Magenta Style = "35"
	Cyan    Style = "36"
	White   Style = "37"

	BrightBlack   Style = "90"
	BrightRed     Style = "91"
	BrightGreen   Style = "92"
	BrightYellow  Style = "93"
	BrightBlue    Style = "94"
	BrightMagenta Style = "95"
	BrightCyan    Style = "96"
	BrightWhite   Style = "97"
)

// Background colors
const (
	BgBlack   Style = "40"
	BgRed     Style = "41"
	BgGreen   Style = "42"
	BgYellow  Style = "43"
	BgBlue    Style = "44"
	BgMagenta Style = "45"
	BgCyan    Style = "46"
	BgWhite   Style = "47"
)

// Effects
const (
	ResetStyle Style = "0"

	Bold          Style = "1"
	Faint         Style = "2"
	Italic        Style = "3"
	Underline     Style = "4"
	SlowBlink     Style = "5"
	RapidBlink    Style = "6"
	Reverse       Style = "7"
	Hidden        Style = "8"
	StrikeThrough Style = "9"

	DoubleUnderline Style = "21"

	NormalIntensity Style = "22"
	NoItalic        Style = "23"
	NoUnderline     Style = "24"
	NoBlink         Style = "25"
	NoReverse       Style = "27"
	Reveal          Style = "28"
	NoStrike        Style = "29"

	Frame    Style = "51"
	Encircle Style = "52"
	Overline Style = "53"

	NoFrame    Style = "54"
	NoOverline Style = "55"
)

func (self *Colorista) Apply(text any, styles ...Style) string {
	s := fmt.Sprint(text)

	if !self.enabled || len(styles) == 0 {
		return s
	}

	codes := make([]string, len(styles))
	for i, st := range styles {
		codes[i] = string(st)
	}

	return "\033[" + strings.Join(codes, ";") + "m" + s + reset
}

func Rgb(color RGB) Style {
	return Style(fmt.Sprintf("38;2;%d;%d;%d", color.R, color.G, color.B))
}

func BgRgb(color RGB) Style {
	return Style(fmt.Sprintf("48;2;%d;%d;%d", color.R, color.G, color.B))
}

func ANSI256(i uint8) Style {
	return Style(fmt.Sprintf("38;5;%d", i))
}

func BgANSI256(i uint8) Style {
	return Style(fmt.Sprintf("48;5;%d", i))
}
