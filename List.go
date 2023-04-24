package gomenu

import (
	"fmt"
	"strings"

	. "github.com/pandasoli/go-scroll"
	"github.com/pandasoli/goterm"
)


func OpenList(settings Settings, raw_items ...string) (selected int, err error) {
  var scroll int

  // Make items
  var items []Item
  var content_h int

  make_items := func() {
    for _, raw_item := range raw_items {
      // break them into pieces the size of the horizontal space
      lines, h := fitstr(
        raw_item,
        settings.W - (settings.LeftItemPadding + settings.RightItemPadding),
      )

      // Appling paddings
      for i, value := range lines {
        lines[i] = strings.Repeat(" ", settings.LeftItemPadding) + value + strings.Repeat(" ", settings.RightItemPadding)
      }

      blank_line := strings.Repeat(" ", settings.W)

      var top_padding_lines []string
      var bottom_padding_lines []string

      for range make([]int, settings.TopItemPadding) {
        top_padding_lines = append(top_padding_lines, blank_line)
      }

      for range make([]int, settings.BottomItemPadding) {
        bottom_padding_lines = append(bottom_padding_lines, blank_line)
      }

      lines = append(top_padding_lines, lines...)
      lines = append(lines, bottom_padding_lines...)
      h += settings.TopItemPadding + settings.BottomItemPadding

      // Saving
      item := Item {
        Rect {
          W: settings.W,
          H: h,
          X: 0,
          Y: content_h,
        },
        lines,
      }

      content_h += h
      items = append(items, item)
    }
  }

  make_items()
  if content_h > settings.H { // needs scroll
    settings.W--
    content_h = 0
    items = []Item {}

    make_items()
  }

  switch settings.Align {
    case "center":
      for itemi, item := range items {
        for linei, line := range item.Lines {
          items[itemi].Lines[linei] = centerstr(line, item.W)
        }
      }

    case "right":
      for itemi, item := range items {
        for linei, line := range item.Lines {
          items[itemi].Lines[linei] = strings.Repeat(" ", item.W - len(line)) + items[itemi].Lines[linei]
        }
      }

    default:
      for itemi, item := range items {
        for linei, line := range item.Lines {
          items[itemi].Lines[linei] += strings.Repeat(" ", item.W - len(line))
        }
      }
  }

  // Draw stuff
  draw_line := func(line string, y int, selected bool) {
    y -= scroll

    if y >= 0 && y < settings.H { // If inside buffer bounds
      cl := settings.ItemCl

      if selected {
        cl = settings.HoverItemCl
      }

      goterm.GoToXY(settings.X, settings.Y + y)
      fmt.Print("\033[0m\033[" + cl + "m" + line)
    }
  }

  draw := func() {
    ShowYScrollbar(content_h, settings.H, settings.X + settings.W, settings.Y, scroll)

    for itemi, item := range items {
      for linei, line := range item.Lines {
        draw_line(line, item.Y + linei, itemi == selected)
      }
    }
  }

  // Main loop
  for {
    draw()

    key, err := goterm.Getch()
    if err != nil { return -1, err }

    if key == "q" { break }

    switch key {
      case "\033[A" /* Up arrow */:
        if selected > 0 {
          selected--

          if items[selected].Y < scroll {
            scroll = items[selected].Y
          }
        }

      case "\033[B" /* Down arrow */:
        if selected < len(items) - 1 {
          selected++

          item_bottom_y := items[selected].Y + items[selected].H

          if item_bottom_y > scroll + settings.H {
            scroll += item_bottom_y - (scroll + settings.H)
          }
        }
    }
  }

  return selected, err
}
