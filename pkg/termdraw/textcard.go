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
	"github.com/asig/termbox-go"
)

type Segment struct {
	S      string
	Fg, Bg termbox.Attribute
}

type Line []Segment

func (l Line) Width() int {
	w := 0
	for _, s := range l {
		w += len(s.S)
	}
	return w
}

type TextCard struct {
	Fg, Bg  termbox.Attribute
	Bs      BorderStyle
	Content []Line
}

func (t *TextCard) Show() {
	contentWidth := 0
	for _, s := range t.Content {
		w := s.Width()
		if w > contentWidth {
			contentWidth = w
		}
	}

	termW, termH := termbox.Size()
	w := contentWidth + 4
	h := len(t.Content) + 2
	x := (termW - w) / 2
	y := (termH - h) / 2

	FillBox(x, y, w, h, t.Fg, t.Bg, t.Bs)
	for i, l := range t.Content {
		px := x + 2
		for _, s := range l {
			Puts(px, y+1+i, s.S, s.Fg, s.Bg)
			px += len(s.S)
		}
	}
	termbox.Flush()
}

