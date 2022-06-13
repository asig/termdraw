/*
 * Copyright (c) 2022 Andreas Signer <asigner@gmail.com>
 *
 * This file is part of termdraw.
 *
 * termdraw is free software: you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * termdraw is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with termdraw.  If not, see <http://www.gnu.org/licenses/>.
 */
package termdraw

import (
	"strings"
	"unicode"

	"github.com/asig/termbox-go"
)

type EditField struct {
	x, y, w int
	fg, bg termbox.Attribute
}

func NewEditField(x,y,w int, fg, bg termbox.Attribute) *EditField {
	return &EditField{
		x:  x,
		y:  y,
		w:  w,
		fg: fg,
		bg: bg,
	}
}

func (e *EditField) Run() (string, bool) {
	res := ""
	crsr := 0
	ofs := 0

	Puts(e.x, e.y, strings.Repeat(" ", e.w), e.fg, e.bg)
	termbox.SetCursor(e.x, e.y)

	termbox.Flush()

	var ok bool
	quit := false
	for !quit {
		data := make([]byte, 30)
		termbox.PollRawEvent(data)
		ev := termbox.ParseEvent(data)
		resPos := crsr + ofs
		switch ev.Type {
		case termbox.EventKey:
			switch {
			case ev.Mod == termbox.ModAlt && ev.Key == 0 && ev.Ch == 0:
				// Only ESC pressed, nothing else
				ok = false
				quit = true
			case ev.Key == termbox.KeyArrowRight:
				if resPos < len(res) {
					crsr++
					if crsr >= e.w {
						crsr--
						ofs++
					}
				}
			case ev.Key == termbox.KeyArrowLeft:
				if resPos > 0 {
					crsr--
					if crsr < 0 {
						crsr++
						ofs--
					}
				}
			case unicode.IsPrint(ev.Ch):
				res = res[:resPos] + string(ev.Ch) + res[resPos:]
				crsr++
				if crsr >= e.w {
					crsr--
					ofs++
				}
			case ev.Key == termbox.KeyDelete:
				if resPos < len(res) {
					res = res[:resPos] + res[resPos+1:]
				}
			case ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2:
				if resPos > 0 {
					res = res[:resPos-1] + res[resPos:]
					crsr--
					if crsr < 0 {
						crsr++
						ofs--
					}
				}
			case ev.Key == termbox.KeyEnter:
				ok = true
				quit = true

			}
		}

		// Update res
		padded := res[ofs:]
		if len(padded) <= e.w {
			padded = padded + strings.Repeat(" ", e.w - len(padded))
		} else {
			padded = padded[:e.w]
		}
		Puts(e.x, e.y, padded, e.fg, e.bg)
		termbox.SetCursor(e.x +crsr, e.y)
		termbox.Flush()
	}

	return res, ok
}
