package gomenu

import (
	"github.com/pandasoli/colorstring/v2"
)


type Color colorstring.Color

type Rect struct {
  W, H,
  X, Y int
}

type Settings struct {
  Rect

  ItemCl,
  HoverItemCl,

  Align string
}

type Item struct {
  Rect

  Lines []string
}
