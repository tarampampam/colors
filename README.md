# colors

[![tests][badge_tests]][actions]
[![release][badge_release]][actions]
[![Coverage][badge_coverage]][coverage]

One more Go library for using colors in the terminal console. The most important features are:

- ANSI colors support (using Escape Sequences)
- Multi-thread safe
- Support `FORCE_COLOR`, `NO_COLOR` and `TERM` variables out of the box
- Super-lightweight
- Color codes are not pre-allocated, but cached (in memory) and re-used where possible
- Easy to integrate with the existing code-base

## Usage examples

```go
package main

import (
  "fmt"

  "github.com/tarampampam/colors"
)

func main() {
  fmt.Println((colors.FgGreen | colors.Bold).Wrap("green color + bold text"))

  var bg = colors.BgRed

  fmt.Printf("%s red background %s\n", bg.Start(), bg.Reset())

  colors.Enabled(false) // disable colors
  colors.Enabled(true)  // enable colors
}
```

For more examples see [examples](./examples) directory.

[badge_tests]:https://github.com/tarampampam/colors/actions/workflows/tests.yml/badge.svg
[badge_release]:https://github.com/tarampampam/colors/actions/workflows/release.yml/badge.svg
[badge_coverage]:https://img.shields.io/codecov/c/github/tarampampam/colors/master.svg?maxAge=30
[actions]:https://github.com/tarampampam/colors/actions
[coverage]:https://codecov.io/gh/tarampampam/colors
