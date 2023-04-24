package main

import (
	"fmt"

	"github.com/pandasoli/go-menu"
	"github.com/pandasoli/goterm"
)


func main() {
  termios, _ := goterm.SetRawMode()
  defer goterm.RestoreMode(termios)

  goterm.HideCursor()

  items := []string {
    "item1",
    "2Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
    "item3",
    "item4",
    "5Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
    "item6",
    "item7",
    "item8",
    "9Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  }

  gomenu.OpenList(
    gomenu.Settings {
      Rect: gomenu.Rect {
        W: 40, H: 6,
        X: 2, Y: 2,
      },
      ItemCl: "2",
      HoverItemCl: "1;37",
      Align: "",

      LeftItemPadding: 1,
      RightItemPadding: 1,
      TopItemPadding: 1,
      BottomItemPadding: 1,
    },
    items...
  )

  goterm.ShowCursor()
  fmt.Println("\n")
}
