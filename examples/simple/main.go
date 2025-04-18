package main

import (
	"fmt"

	"gh.tarampamp.am/colors"
)

func main() { //nolint:funlen
	fmt.Println( //nolint:forbidigo
		(colors.FgGreen | colors.Bold).Wrap("tarampampam/colors:"),
		(colors.FgBlue | colors.FgBright).Wrap("an"),
		(colors.FgWhite | colors.FgBright).Wrap("ANSI"),
		(colors.FgYellow | colors.FgBright).Wrap("colors"),
		(colors.FgRed | colors.FgBright).Wrap("are"),
		(colors.FgMagenta | colors.FgBright).Wrap("included"),
	)

	println() //nolint:forbidigo // empty line

	for _, set := range []struct {
		name                          string
		fg, bg, bgBright, customStyle colors.TextStyle
	}{
		{
			"Red",
			colors.FgRed,
			colors.BgRed | colors.FgBlack,
			colors.BgRed | colors.BgBright | colors.FgBlack,
			colors.FgRed | colors.FgBright | colors.Bold | colors.Italic | colors.Underline,
		},
		{
			"Green",
			colors.FgGreen,
			colors.BgGreen | colors.FgBlack,
			colors.BgGreen | colors.BgBright | colors.FgBlack,
			colors.FgGreen | colors.FgBright | colors.Bold | colors.Italic | colors.Underline,
		},
		{
			"Yellow",
			colors.FgYellow,
			colors.BgYellow | colors.FgBlack,
			colors.BgYellow | colors.BgBright | colors.FgBlack,
			colors.FgYellow | colors.FgBright | colors.Bold | colors.Italic | colors.Underline,
		},
		{
			"Blue",
			colors.FgBlue,
			colors.BgBlue | colors.FgBlack,
			colors.BgBlue | colors.BgBright | colors.FgBlack,
			colors.FgBlue | colors.FgBright | colors.Bold | colors.Italic | colors.Underline,
		},
		{
			"Magenta",
			colors.FgMagenta,
			colors.BgMagenta | colors.FgBlack,
			colors.BgMagenta | colors.BgBright | colors.FgBlack,
			colors.FgMagenta | colors.FgBright | colors.Bold | colors.Italic | colors.Underline,
		},
		{
			"Cyan",
			colors.FgCyan,
			colors.BgCyan | colors.FgBlack,
			colors.BgCyan | colors.BgBright | colors.FgBlack,
			colors.FgCyan | colors.FgBright | colors.Bold | colors.Italic | colors.Underline,
		},
		{
			"White",
			colors.FgWhite,
			colors.BgWhite | colors.FgBlack,
			colors.BgWhite | colors.BgBright | colors.FgBlack,
			colors.FgWhite | colors.FgBright | colors.Bold | colors.Italic | colors.Underline,
		},
	} {
		fmt.Printf("%s\t", set.fg.Wrap(set.name))                                                 //nolint:forbidigo
		fmt.Printf("%s normal %s\t", set.bg.Start(), set.bg.Reset())                              //nolint:forbidigo
		fmt.Printf("%s bright %s\t", set.bgBright.Start(), set.bgBright.Reset())                  //nolint:forbidigo
		fmt.Printf("%sStyled %s%s\n", set.customStyle.Start(), set.name, set.customStyle.Reset()) //nolint:forbidigo
	}
}
