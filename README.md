# 🦜 colors

![go_version][badge_go_version]
[![tests][badge_tests]][actions]
[![coverage][badge_coverage]][coverage]
[![docs][badge_docs]][docs]

<div align="center">

![screenshot](https://user-images.githubusercontent.com/7326800/197354081-4a083eb2-5f94-4475-9ac0-e23a6de9a378.png)

</div>

One more Go library for using colors in the terminal console. The most important features are:

- ANSI colors support (using Escape Sequences)
- Multi-thread safe
- Support `FORCE_COLOR`, `NO_COLOR` and `TERM` variables out of the box
- Super-lightweight and extremely fast
- Color codes are not pre-allocated, but cached (in memory) and re-used upon further usage
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

[badge_tests]:https://img.shields.io/github/actions/workflow/status/tarampampam/colors/tests.yml?branch=master
[badge_coverage]:https://img.shields.io/codecov/c/github/tarampampam/colors/master.svg?maxAge=30
[badge_docs]:https://pkg.go.dev/badge/mod/github.com/tarampampam/colors
[badge_go_version]:https://img.shields.io/badge/go%20version-%3E=1.16-61CFDD.svg
[actions]:https://github.com/tarampampam/colors/actions
[coverage]:https://codecov.io/gh/tarampampam/colors
[docs]:https://pkg.go.dev/github.com/tarampampam/colors
