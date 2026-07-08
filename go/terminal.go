package colorista

import (
	"os"
	"strconv"
	"strings"
)

type Theme uint8

const (
	ThemeAuto Theme = iota
	ThemeDark
	ThemeLight
)

func detectTheme() Theme {
	if fg := os.Getenv("COLORFGBG"); fg != "" {
		parts := strings.Split(fg, ";")

		if len(parts) > 0 {

			bg, err := strconv.Atoi(parts[len(parts)-1])

			if err == nil {
				if bg <= 7 {
					return ThemeDark
				}

				return ThemeLight
			}
		}
	}

	if os.Getenv("WT_SESSION") != "" {
		return ThemeDark
	}

	if os.Getenv("TERM_PROGRAM") == "iTerm.app" {
		return ThemeDark
	}

	return ThemeDark
}
