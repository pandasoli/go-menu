package gomenu

import (
	"fmt"
	"strings"

	"github.com/pandasoli/goterm"
)


func fitstr(str string, w int) (res []string, size int) {
  for i := 0; i < len(str); i += w {
    part := str[i:]

    if len(part) > w {
      part = part[:w]
    }

    size++
    res = append(res, part)
  }

  return res, size
}

func centerstr(str string, w int) string {
  usable_space := w - len([]rune(str))

  padding_left := strings.Repeat(" ", usable_space / 2)
  padding_right := strings.Repeat(" ", (usable_space + 1) / 2)

  return padding_left + str + padding_right
}

func makeBoundsBorder(bounds Rect) {
  goterm.GoToXY(bounds.X - 1, bounds.Y - 1)
  fmt.Print("\033[44m" + strings.Repeat(" ", bounds.W + 2))

  goterm.GoToXY(bounds.X - 1, bounds.Y + bounds.H)
  fmt.Print("\033[44m" + strings.Repeat(" ", bounds.W + 2))

  for y := range make([]int, bounds.H) {
    y += bounds.Y

    goterm.GoToXY(bounds.X - 1, y)
    fmt.Print("\033[44m ")

    goterm.GoToXY(bounds.X + bounds.W, y)
    fmt.Print("\033[44m ")
  }
}

var debug_line int
func debug(a ...any) {
  goterm.GoToXY(70, 1 + debug_line)
  debug_line++

  for i, item := range a {
    if str, ok := item.(string); ok {
      a[i] = "\"" + strings.ReplaceAll(str, "\033", "\\033") + "\""
    }
  }
  
  fmt.Print(a...)
}
