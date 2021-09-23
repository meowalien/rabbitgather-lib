package text

import (
	"fmt"
	"regexp"
)
const AllASCII = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLM_NOPQRSTUVWXYZ[\\]^`abcdefghijklmnopqrstuvwxyz{|}~"
const AllASCIICount = 95

const (
	Reset ColorCode = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack ColorCode = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

type ColorCode int

// Foreground Hi-Intensity text colors
const (
	FgHiBlack ColorCode = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack ColorCode = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack ColorCode = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// ColorSting Color the given string to the given color
func ColorSting(s string, color ColorCode) string {
	return fmt.Sprintf("\033[%dm%s\033[00m", color, s)
}

//var englishLetter = []rune(AllASCII)

var PlainEnglishOnlyRegexp = regexp.MustCompile(fmt.Sprintf(`^[%s%s]+$`, AllASCII, "_"))
