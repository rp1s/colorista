package colorista

import (
	"fmt"
	"strings"
)

type GradientPos struct {
	Pos   float64 // 0.0 - 1.0
	Color RGB
}

func (self *Colorista) FullGradient(
	text string,
	fg []GradientPos,
	bg []GradientPos,
	styles ...Style,
) string {
	if !self.enabled {
		return text
	}

	runes := []rune(text)

	var out strings.Builder

	for i, ch := range runes {
		pos := 0.0

		if len(runes) > 1 {
			pos = float64(i) / float64(len(runes)-1)
		}

		codes := make([]Style, 0, len(styles)+2)

		codes = append(codes, styles...)

		if len(fg) > 0 {
			fgColor := colorAt(fg, pos)
			codes = append(codes, self.rgbStyle(fgColor))
		}

		if len(bg) > 0 {
			bgColor := colorAt(bg, pos)
			codes = append(codes, self.bgRGBStyle(bgColor))
		}

		out.WriteString(self.Apply(string(ch), codes...))
	}

	return out.String()
}

func (self *Colorista) Gradient(text string, stops []GradientPos, styles ...Style) string {
	if !self.enabled {
		return text
	}

	r := []rune(text)

	var out strings.Builder

	for i, ch := range r {

		p := 0.0
		if len(r) > 1 {
			p = float64(i) / float64(len(r)-1)
		}

		c := colorAt(stops, p)

		s := append([]Style{}, styles...)
		s = append(s, self.rgbStyle(c))

		out.WriteString(self.Apply(string(ch), s...))
	}

	return out.String()
}

func (self *Colorista) BgGradient(text string, stops []GradientPos, styles ...Style) string {
	if !self.enabled {
		return text
	}

	r := []rune(text)

	var out strings.Builder

	for i, ch := range r {

		p := 0.0
		if len(r) > 1 {
			p = float64(i) / float64(len(r)-1)
		}

		c := colorAt(stops, p)

		s := append([]Style{}, styles...)
		s = append(s, self.bgRGBStyle(c))

		out.WriteString(self.Apply(string(ch), s...))
	}

	return out.String()
}

func (self *Colorista) rgbStyle(c RGB) Style {

	c = self.CorrectContrast(c)

	return Style(fmt.Sprintf(
		"38;2;%d;%d;%d",
		c.R,
		c.G,
		c.B,
	))
}

func (self *Colorista) bgRGBStyle(c RGB) Style {

	c = self.CorrectContrast(c)

	return Style(fmt.Sprintf(
		"48;2;%d;%d;%d",
		c.R,
		c.G,
		c.B,
	))
}

func colorAt(gp []GradientPos, pos float64) RGB {
	if len(gp) == 0 {
		return RGB{}
	}

	if pos <= gp[0].Pos {
		return gp[0].Color
	}

	for i := 0; i < len(gp)-1; i++ {
		a := gp[i]
		b := gp[i+1]

		if pos >= a.Pos && pos <= b.Pos {
			t := (pos - a.Pos) / (b.Pos - a.Pos)
			return interpolate(a.Color, b.Color, t)
		}
	}

	return gp[len(gp)-1].Color
}

func interpolate(a, b RGB, t float64) RGB {
	return RGB{
		R: lerp(a.R, b.R, t),
		G: lerp(a.G, b.G, t),
		B: lerp(a.B, b.B, t),
	}
}

func lerp(a, b uint8, t float64) uint8 {
	return uint8(float64(a) + (float64(b)-float64(a))*t)
}

func luminance(c RGB) float64 {
	return 0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B)

}

func adjustBrightness(c RGB, factor float64) RGB {
	return RGB{
		R: uint8(clamp(float64(c.R) * factor)),
		G: uint8(clamp(float64(c.G) * factor)),
		B: uint8(clamp(float64(c.B) * factor)),
	}
}

func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}

	if v > 255 {
		return 255
	}

	return v
}

func (self *Colorista) CorrectContrast(c RGB) RGB {

	switch self.theme {

	case ThemeLight:
		if luminance(c) > 120 {
			return adjustBrightness(c, 0.55)
		}

		return c

	case ThemeDark:
		if luminance(c) < 80 {
			return adjustBrightness(c, 1.7)
		}

		return c
	}

	return c
}
