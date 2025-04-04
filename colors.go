package colors

import (
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"gh.tarampamp.am/colors/internal/isatty"
)

const (
	colorsOff uint32 = iota
	colorsOn
)

var colorsEnabled = initColorsState() //nolint:gochecknoglobals // atomic usage only

// initColorsState returns initialization value for the colors enabled state.
func initColorsState() uint32 {
	if _, exists := os.LookupEnv("FORCE_COLOR"); exists {
		return colorsOn
	} else if _, exists = os.LookupEnv("NO_COLOR"); exists { // docs: <https://no-color.org/>
		return colorsOff
	} else if os.Getenv("TERM") == "dumb" {
		return colorsOff
	} else if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return colorsOff
	}

	return colorsOn
}

// Enabled returns true if colors are enabled. Also, you can set a new state (enable or disable colors).
func Enabled(newState ...bool) bool {
	if len(newState) == 0 {
		return atomic.LoadUint32(&colorsEnabled) == colorsOn
	}

	var set uint32

	if newState[0] {
		set = colorsOn // enable colors
	} else {
		set = colorsOff // disable colors
	}

	atomic.StoreUint32(&colorsEnabled, set)

	return set == colorsOn
}

// TextStyle is a set of text styles.
//
// Developer note:
//
//	uint32 = 0b11111111111111111111111111111111
//	                                  ^^^^^^^^^ - foreground color
//	                                 ^ - bright foreground color bit
//	                        ^^^^^^^^^ - background color
//	                       ^ - bright background color bit
//	               ^^^^^^^^ - text style
//	              ^ - reset style bit
//	           ^^^ - reserved bits
//
// Pretty doc: <https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797>
type TextStyle uint32

const (
	FgBlack   TextStyle = 1 << iota // Black text color
	FgRed                           // Red text color
	FgGreen                         // Green text color
	FgYellow                        // Yellow text color
	FgBlue                          // Blue text color
	FgMagenta                       // Magenta text color
	FgCyan                          // Cyan text color
	FgWhite                         // White text color
	FgDefault                       // Default text color

	FgBright // Bright text color, usage example: (FgRed | FgBright).Wrap("hello world")

	BgBlack   // Black background color
	BgRed     // Red background color
	BgGreen   // Green background color
	BgYellow  // Yellow background color
	BgBlue    // Blue background color
	BgMagenta // Magenta background color
	BgCyan    // Cyan background color
	BgWhite   // White background color
	BgDefault // Default background color

	BgBright // Bright background color, usage example: (BgRed | BgBright).Wrap("hello world")

	Bold      // Bold text
	Faint     // Faint text
	Italic    // Italic text
	Underline // Underline text
	Blinking  // Blinking text
	Reverse   // Reverse text
	Invisible // Invisible text
	Strike    // Strike text

	Reset // Reset text style

	_ // reserved
	_ // reserved
	_ // reserved
)

// Has returns true if provided text style included into this one.
func (ts TextStyle) Has(z TextStyle) bool { return ts&z != 0 }

// Add adds styles to the test style.
func (ts *TextStyle) Add(styles ...TextStyle) {
	for _, style := range styles {
		*ts |= style
	}
}

// Remove removes provided text styles.
func (ts *TextStyle) Remove(styles ...TextStyle) {
	for _, style := range styles {
		*ts &^= style
	}
}

// ternary is a ternary sugar.
func (TextStyle) ternary(cond bool, ok, notOk byte) byte {
	if cond {
		return ok
	}

	return notOk
}

// byteToString converts byte to string.
func (TextStyle) byteToString(t byte) string {
	var (
		s [3]byte
		j = 2
	)

	for i := 0; i < 3; i, j = i+1, j-1 {
		s[j] = '0' + t%10 //nolint:mnd

		if t /= 10; t == 0 {
			break
		}
	}

	return string(s[j:])
}

// rawColorCodes returns raw color codes (as a slice of bytes).
func (ts TextStyle) rawColorCodes() (start, reset []byte) { //nolint:funlen,gocyclo
	const resetByte byte = 0

	if ts.Has(Reset) {
		start = append(start, resetByte)

		return //nolint:nakedret
	}

	const (
		boldByte, boldResetByte           byte = 1, 22
		faintByte, faintResetByte         byte = 2, 22
		italicByte, italicResetByte       byte = 3, 23
		underlineByte, underlineResetByte byte = 4, 24
		blinkingByte, blinkingResetByte   byte = 5, 25
		reverseByte, reverseResetByte     byte = 7, 27
		invisibleByte, invisibleResetByte byte = 8, 28
		strikeByte, strikeResetByte       byte = 9, 29
	)

	if ts.Has(Bold) {
		start, reset = append(start, boldByte), append([]byte{boldResetByte}, reset...)
	}

	if ts.Has(Faint) {
		start, reset = append(start, faintByte), append([]byte{faintResetByte}, reset...)
	}

	if ts.Has(Italic) {
		start, reset = append(start, italicByte), append([]byte{italicResetByte}, reset...)
	}

	if ts.Has(Underline) {
		start, reset = append(start, underlineByte), append([]byte{underlineResetByte}, reset...)
	}

	if ts.Has(Blinking) {
		start, reset = append(start, blinkingByte), append([]byte{blinkingResetByte}, reset...)
	}

	if ts.Has(Reverse) {
		start, reset = append(start, reverseByte), append([]byte{reverseResetByte}, reset...)
	}

	if ts.Has(Invisible) {
		start, reset = append(start, invisibleByte), append([]byte{invisibleResetByte}, reset...)
	}

	if ts.Has(Strike) {
		start, reset = append(start, strikeByte), append([]byte{strikeResetByte}, reset...)
	}

	var fgCode, fgBright = byte(0), ts.Has(FgBright)

	const (
		fgDefaultByte                      byte = 39
		fgBlackByte, fgBlackBrightByte     byte = 30, 90
		fgRedByte, fgRedBrightByte         byte = 31, 91
		fgGreenByte, fgGreenBrightByte     byte = 32, 92
		fgYellowByte, fgYellowBrightByte   byte = 33, 93
		fgBlueByte, fgBlueBrightByte       byte = 34, 94
		fgMagentaByte, fgMagentaBrightByte byte = 35, 95
		fgCyanByte, fgCyanBrightByte       byte = 36, 96
		fgWhiteByte, fgWhiteBrightByte     byte = 37, 97
	)

	switch { //nolint:dupl
	case ts.Has(FgDefault):
		fgCode = fgDefaultByte
	case ts.Has(FgBlack):
		fgCode = ts.ternary(fgBright, fgBlackBrightByte, fgBlackByte)
	case ts.Has(FgRed):
		fgCode = ts.ternary(fgBright, fgRedBrightByte, fgRedByte)
	case ts.Has(FgGreen):
		fgCode = ts.ternary(fgBright, fgGreenBrightByte, fgGreenByte)
	case ts.Has(FgYellow):
		fgCode = ts.ternary(fgBright, fgYellowBrightByte, fgYellowByte)
	case ts.Has(FgBlue):
		fgCode = ts.ternary(fgBright, fgBlueBrightByte, fgBlueByte)
	case ts.Has(FgMagenta):
		fgCode = ts.ternary(fgBright, fgMagentaBrightByte, fgMagentaByte)
	case ts.Has(FgCyan):
		fgCode = ts.ternary(fgBright, fgCyanBrightByte, fgCyanByte)
	case ts.Has(FgWhite):
		fgCode = ts.ternary(fgBright, fgWhiteBrightByte, fgWhiteByte)
	}

	if fgCode != 0 {
		start = append(start, fgCode)

		if fgCode != fgDefaultByte {
			reset = append([]byte{fgDefaultByte}, reset...) // prepend fg reset color code
		}
	}

	var bgCode, bgBright = byte(0), ts.Has(BgBright)

	const (
		bgDefaultByte                      byte = 49
		bgBlackByte, bgBlackBrightByte     byte = 40, 100
		bgRedByte, bgRedBrightByte         byte = 41, 101
		bgGreenByte, bgGreenBrightByte     byte = 42, 102
		bgYellowByte, bgYellowBrightByte   byte = 43, 103
		bgBlueByte, bgBlueBrightByte       byte = 44, 104
		bgMagentaByte, bgMagentaBrightByte byte = 45, 105
		bgCyanByte, bgCyanBrightByte       byte = 46, 106
		bgWhiteByte, bgWhiteBrightByte     byte = 47, 107
	)

	switch { //nolint:dupl
	case ts.Has(BgDefault):
		bgCode = bgDefaultByte
	case ts.Has(BgBlack):
		bgCode = ts.ternary(bgBright, bgBlackBrightByte, bgBlackByte)
	case ts.Has(BgRed):
		bgCode = ts.ternary(bgBright, bgRedBrightByte, bgRedByte)
	case ts.Has(BgGreen):
		bgCode = ts.ternary(bgBright, bgGreenBrightByte, bgGreenByte)
	case ts.Has(BgYellow):
		bgCode = ts.ternary(bgBright, bgYellowBrightByte, bgYellowByte)
	case ts.Has(BgBlue):
		bgCode = ts.ternary(bgBright, bgBlueBrightByte, bgBlueByte)
	case ts.Has(BgMagenta):
		bgCode = ts.ternary(bgBright, bgMagentaBrightByte, bgMagentaByte)
	case ts.Has(BgCyan):
		bgCode = ts.ternary(bgBright, bgCyanBrightByte, bgCyanByte)
	case ts.Has(BgWhite):
		bgCode = ts.ternary(bgBright, bgWhiteBrightByte, bgWhiteByte)
	}

	if bgCode != 0 {
		start = append(start, bgCode)

		if bgCode != bgDefaultByte {
			reset = append([]byte{bgDefaultByte}, reset...) // prepend bg reset color code
		}
	}

	return start, reset
}

var ccCache = struct { //nolint:gochecknoglobals // color codes in-memory cache
	sync.Mutex
	m map[TextStyle][2]string
}{
	m: make(map[TextStyle][2]string),
}

// ColorCodes returns color codes for the text style. Important note: the result of this function working does not
// depend on the colors enabling state.
func (ts TextStyle) ColorCodes() (start, reset string) {
	if ts == 0 {
		return
	}

	ccCache.Lock()
	cached, ok := ccCache.m[ts] // read from cache
	ccCache.Unlock()

	if ok {
		return cached[0], cached[1]
	}

	var (
		buf                strings.Builder
		rawStart, rawReset = ts.rawColorCodes()
	)

	const esc = "\x1b["

	if len(rawStart) != 0 {
		buf.WriteString(esc)

		for i := 0; i < len(rawStart); i++ {
			buf.WriteString(ts.byteToString(rawStart[i]))

			if i < len(rawStart)-1 {
				buf.WriteRune(';')
			}
		}

		buf.WriteRune('m')

		start = buf.String()
	}

	if len(rawReset) != 0 {
		buf.Reset()

		buf.WriteString(esc)

		for i := 0; i < len(rawReset); i++ {
			buf.WriteString(ts.byteToString(rawReset[i]))

			if i < len(rawReset)-1 {
				buf.WriteRune(';')
			}
		}

		buf.WriteRune('m')

		reset = buf.String()
	}

	ccCache.Lock()
	ccCache.m[ts] = [2]string{start, reset} // put into cache
	ccCache.Unlock()

	return start, reset
}

// String returns a string starting text styling (useful for usage with fmt.Sprintf).
// Note: Don't forget to use Reset() to reset the styling (resting is NOT needed for FgDefault, BgDefault and Reset).
func (ts TextStyle) String() string { return ts.Start() }

// Start returns current text style starting code. An empty string will return when colors are disabled.
func (ts TextStyle) Start() (start string) {
	if ts == 0 {
		return
	}

	if !Enabled() {
		return
	}

	start, _ = ts.ColorCodes()

	return
}

// Reset returns current text style resetting code. An empty string will return when colors are disabled or when called
// on FgDefault, BgDefault, or Reset.
func (ts TextStyle) Reset() (reset string) {
	if ts == 0 {
		return
	}

	if !Enabled() {
		return
	}

	_, reset = ts.ColorCodes()

	return
}

// Wrap wraps provided string with staring and reset color codes. The provided string will return without any
// modifications when colors are disabled.
func (ts TextStyle) Wrap(s string) string {
	if ts == 0 {
		return s
	}

	if !Enabled() {
		return s
	}

	var (
		start, reset = ts.ColorCodes()
		buf          strings.Builder
	)

	buf.Grow(len(start) + len(s) + len(reset))

	buf.WriteString(start)
	buf.WriteString(s)
	buf.WriteString(reset)

	return buf.String()
}
