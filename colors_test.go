package colors_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tarampampam/colors"
)

func ExampleTextStyle_Wrap() {
	colors.Enabled(false) // change to true to see colors

	fmt.Println((colors.FgRed | colors.Bold).Wrap("Foo Bar"))

	// output:
	// Foo Bar
}

func ExampleTextStyle_Start() {
	colors.Enabled(false) // change to true to see colors

	var style = colors.FgRed | colors.Bold

	fmt.Println(style.Start(), "Foo Bar", style.Reset())

	// output:
	// Foo Bar
}

func TestColorsEnabled(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			colors.Enabled(false)
			colors.Enabled(true)
			colors.Enabled()
		}()
	}

	wg.Wait()

	colors.Enabled(false)

	assert.False(t, colors.Enabled())
	colors.Enabled(false)
	assert.False(t, colors.Enabled())

	colors.Enabled(true)
	assert.True(t, colors.Enabled())
}

func TestTextStyle_Has(t *testing.T) {
	var s = colors.BgBlack | colors.FgWhite | colors.Bold

	assert.False(t, s.Has(colors.BgWhite))
	assert.False(t, s.Has(colors.FgBlack))
	assert.False(t, s.Has(colors.Italic))
	assert.False(t, s.Has(colors.Reset))
	assert.False(t, s.Has(colors.BgDefault))

	assert.True(t, s.Has(colors.BgBlack))
	assert.True(t, s.Has(colors.FgWhite))
	assert.True(t, s.Has(colors.Bold))
}

func TestTextStyle_Add(t *testing.T) {
	var s = colors.BgBlack

	assert.False(t, s.Has(colors.BgWhite))
	assert.False(t, s.Has(colors.Bold))
	assert.True(t, s.Has(colors.BgBlack))
	assert.False(t, s.Has(colors.Underline))

	s.Add(colors.BgWhite, colors.Bold)

	assert.True(t, s.Has(colors.BgWhite))
	assert.True(t, s.Has(colors.Bold))
	assert.False(t, s.Has(colors.Underline))
}

func TestTextStyle_Remove(t *testing.T) {
	var s = colors.BgBlack | colors.BgWhite | colors.Bold

	assert.True(t, s.Has(colors.BgWhite))
	assert.True(t, s.Has(colors.Bold))
	assert.True(t, s.Has(colors.BgBlack))
	assert.False(t, s.Has(colors.Underline))

	s.Remove(colors.BgWhite, colors.Bold, colors.Underline)

	assert.False(t, s.Has(colors.BgWhite))
	assert.False(t, s.Has(colors.Bold))
	assert.False(t, s.Has(colors.Underline))
}

func TestTextStyle_ColorCodes(t *testing.T) {
	var colorsState = colors.Enabled()

	defer colors.Enabled(colorsState)

	for name, tt := range map[string]struct {
		giveTextStyle        colors.TextStyle
		wantStart, wantReset string
	}{
		"Reset":                      {colors.Reset, "\x1b[0m", ""},
		"Reset | FgBlack | BgYellow": {colors.Reset | colors.FgBlack | colors.BgYellow, "\x1b[0m", ""},

		"FgBlack":   {colors.FgBlack, "\x1b[30m", "\x1b[39m"},
		"FgRed":     {colors.FgRed, "\x1b[31m", "\x1b[39m"},
		"FgGreen":   {colors.FgGreen, "\x1b[32m", "\x1b[39m"},
		"FgYellow":  {colors.FgYellow, "\x1b[33m", "\x1b[39m"},
		"FgBlue":    {colors.FgBlue, "\x1b[34m", "\x1b[39m"},
		"FgMagenta": {colors.FgMagenta, "\x1b[35m", "\x1b[39m"},
		"FgCyan":    {colors.FgCyan, "\x1b[36m", "\x1b[39m"},
		"FgWhite":   {colors.FgWhite, "\x1b[37m", "\x1b[39m"},
		"FgDefault": {colors.FgDefault, "\x1b[39m", ""},

		"FgBlack | FgBright":   {colors.FgBlack | colors.FgBright, "\x1b[90m", "\x1b[39m"},
		"FgRed | FgBright":     {colors.FgRed | colors.FgBright, "\x1b[91m", "\x1b[39m"},
		"FgGreen | FgBright":   {colors.FgGreen | colors.FgBright, "\x1b[92m", "\x1b[39m"},
		"FgYellow | FgBright":  {colors.FgYellow | colors.FgBright, "\x1b[93m", "\x1b[39m"},
		"FgBlue | FgBright":    {colors.FgBlue | colors.FgBright, "\x1b[94m", "\x1b[39m"},
		"FgMagenta | FgBright": {colors.FgMagenta | colors.FgBright, "\x1b[95m", "\x1b[39m"},
		"FgCyan | FgBright":    {colors.FgCyan | colors.FgBright, "\x1b[96m", "\x1b[39m"},
		"FgWhite | FgBright":   {colors.FgWhite | colors.FgBright, "\x1b[97m", "\x1b[39m"},
		"FgDefault | FgBright": {colors.FgDefault | colors.FgBright, "\x1b[39m", ""},

		"BgBlack":   {colors.BgBlack, "\x1b[40m", "\x1b[49m"},
		"BgRed":     {colors.BgRed, "\x1b[41m", "\x1b[49m"},
		"BgGreen":   {colors.BgGreen, "\x1b[42m", "\x1b[49m"},
		"BgYellow":  {colors.BgYellow, "\x1b[43m", "\x1b[49m"},
		"BgBlue":    {colors.BgBlue, "\x1b[44m", "\x1b[49m"},
		"BgMagenta": {colors.BgMagenta, "\x1b[45m", "\x1b[49m"},
		"BgCyan":    {colors.BgCyan, "\x1b[46m", "\x1b[49m"},
		"BgWhite":   {colors.BgWhite, "\x1b[47m", "\x1b[49m"},
		"BgDefault": {colors.BgDefault, "\x1b[49m", ""},

		"BgBlack | BgBright":   {colors.BgBlack | colors.BgBright, "\x1b[100m", "\x1b[49m"},
		"BgRed | BgBright":     {colors.BgRed | colors.BgBright, "\x1b[101m", "\x1b[49m"},
		"BgGreen | BgBright":   {colors.BgGreen | colors.BgBright, "\x1b[102m", "\x1b[49m"},
		"BgYellow | BgBright":  {colors.BgYellow | colors.BgBright, "\x1b[103m", "\x1b[49m"},
		"BgBlue | BgBright":    {colors.BgBlue | colors.BgBright, "\x1b[104m", "\x1b[49m"},
		"BgMagenta | BgBright": {colors.BgMagenta | colors.BgBright, "\x1b[105m", "\x1b[49m"},
		"BgCyan | BgBright":    {colors.BgCyan | colors.BgBright, "\x1b[106m", "\x1b[49m"},
		"BgWhite | BgBright":   {colors.BgWhite | colors.BgBright, "\x1b[107m", "\x1b[49m"},
		"BgDefault | BgBright": {colors.BgDefault | colors.BgBright, "\x1b[49m", ""},

		"Bold":      {colors.Bold, "\x1b[1m", "\x1b[22m"},
		"Faint":     {colors.Faint, "\x1b[2m", "\x1b[22m"},
		"Italic":    {colors.Italic, "\x1b[3m", "\x1b[23m"},
		"Underline": {colors.Underline, "\x1b[4m", "\x1b[24m"},
		"Blinking":  {colors.Blinking, "\x1b[5m", "\x1b[25m"},
		"Reverse":   {colors.Reverse, "\x1b[7m", "\x1b[27m"},
		"Invisible": {colors.Invisible, "\x1b[8m", "\x1b[28m"},
		"Strike":    {colors.Strike, "\x1b[9m", "\x1b[29m"},

		"FgBlack(2) | FgBright | Bold | Underline": {
			colors.FgBlack | colors.FgBlack | colors.FgBright | colors.Bold | colors.Underline, //nolint:gocritic
			"\x1b[1;4;90m",
			"\x1b[39;24;22m",
		},

		"<zero>": {0, "", ""},
	} {
		t.Run(name, func(t *testing.T) {
			colors.Enabled(true) // enable colors

			var start, reset = tt.giveTextStyle.ColorCodes()

			assert.EqualValues(t, tt.wantStart, start)
			assert.EqualValues(t, tt.wantReset, reset)

			assert.EqualValues(t, tt.wantStart, tt.giveTextStyle.Start())
			assert.EqualValues(t, tt.wantStart, tt.giveTextStyle.String())
			assert.EqualValues(t, tt.wantReset, tt.giveTextStyle.Reset())

			colors.Enabled(false) // disable colors

			start, reset = tt.giveTextStyle.ColorCodes()

			assert.EqualValues(t, tt.wantStart, start) // not changed
			assert.EqualValues(t, tt.wantReset, reset) // not changed

			assert.EqualValues(t, "", tt.giveTextStyle.Start())  // empty
			assert.EqualValues(t, "", tt.giveTextStyle.String()) // empty
			assert.EqualValues(t, "", tt.giveTextStyle.Reset())  // empty
		})
	}
}

func TestTextStyle_Wrap(t *testing.T) {
	var (
		colorsState = colors.Enabled()
		testStyle   = colors.FgBlack | colors.FgBright | colors.Bold | colors.Underline
	)

	defer colors.Enabled(colorsState)

	colors.Enabled(true) // enable colors

	assert.EqualValues(t, "\x1b[1;4;90mFOOBAR\x1b[39;24;22m", testStyle.Wrap("FOOBAR"))

	colors.Enabled(false) // disable colors

	assert.EqualValues(t, "FOOBAR", testStyle.Wrap("FOOBAR"))
}

var bmWrapRes string

func BenchmarkColorCodes(b *testing.B) {
	var colorsState = colors.Enabled()

	defer colors.Enabled(colorsState)

	colors.Enabled(true)
	b.ReportAllocs()
	_ = bmWrapRes //nolint:wsl

	for i := 0; i < b.N; i++ {
		bmWrapRes = (colors.FgGreen | colors.BgRed | colors.Bold).Wrap("FOOBAR")
	}
}
