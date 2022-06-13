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
	"unicode/utf8"

	"github.com/asig/termbox-go"
)

type block struct {
	x, y, w, h int
	data       []termbox.Cell
}

func saveBlock(x, y, w, h int) block {
	b := block{
		x:    x,
		y:    y,
		w:    w,
		h:    h,
		data: make([]termbox.Cell, w*h),
	}

	termW, _ := termbox.Size()
	cells := termbox.CellBuffer()
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			b.data[i*w+j] = cells[(y+i)*termW+x+j]
		}
	}
	return b
}

func restoreBlock(b block) {
	termW, _ := termbox.Size()
	cells := termbox.CellBuffer()
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			cells[(b.y+i)*termW+b.x+j] = b.data[i*b.w+j]
		}
	}
}

func DrawBox(x, y, w, h int, fg termbox.Attribute, bg termbox.Attribute, bs BorderStyle) {
	border := borderMap[bs]
	termbox.SetCell(x, y, border[0][0], fg, bg)
	termbox.SetCell(x+w-1, y, border[0][4], fg, bg)
	termbox.SetCell(x, y+h-1, border[4][0], fg, bg)
	termbox.SetCell(x+w-1, y+h-1, border[4][4], fg, bg)

	for i := x + 1; i < x+w-1; i++ {
		termbox.SetCell(i, y, border[0][1], fg, bg)
		termbox.SetCell(i, y+h-1, border[4][1], fg, bg)
	}

	for i := y + 1; i < y+h-1; i++ {
		termbox.SetCell(x, i, border[1][0], fg, bg)
		termbox.SetCell(x+w-1, i, border[1][4], fg, bg)
	}
}

func FillBox(x, y, w, h int, fg termbox.Attribute, bg termbox.Attribute, bs BorderStyle) {
	DrawBox(x, y, w, h, fg, bg, bs)
	for j := y + 1; j < y+h-1; j++ {
		for i := x + 1; i < x+w-1; i++ {
			termbox.SetCell(i, j, ' ', fg, bg)
		}
	}
}

func Puts(x, y int, s string, fg, bg termbox.Attribute) {
	p := 0
	for p < len(s) {
		r, w := utf8.DecodeRune([]byte(s[p:]))
		termbox.SetCell(x, y, r, fg, bg)
		x++
		p += w
	}
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func FileDialog(title string) (string, bool) {
	label := "Filename: "
	editW := 50
	termW, termH := termbox.Size()
	h := 4
	w := min(termW, 4+len(label)+editW)

	px := (termW - w) / 2
	py := (termH - h) / 2

	// Save background
	buf := saveBlock(px, py, w, h)
	savedCrsrX, savedCrsrY := termbox.GetCursor()

	FillBox(px, py, w, h, ColLightCyan, ColBlue, BorderStyle_Double)

	Puts(px+5, py+2, "<Esc> to cancel", ColLightBlue, ColBlue)
	s := "<Enter> to confirm"
	Puts(px+w-5-len(s), py+2, s, ColLightBlue, ColBlue)

	Puts(px+int((w-len(title)+2)/2), py, " "+title+" ", ColWhite, ColBlue)
	Puts(px+2, py+1, label, ColWhite, ColBlue)
	termbox.Flush()

	editField := NewEditField(px+2+len(label), py+1, editW, ColWhite, ColBlack)
	filename, ok := editField.Run()

	restoreBlock(buf)
	termbox.SetCursor(savedCrsrX, savedCrsrY)
	termbox.Flush()

	return filename, ok
}

func YesNoCancelDialog(title, message string) (res bool, valid bool) {
	buttons := "<Y> or <N>. <Esc> to cancel"
	textW := max(max(len(buttons), len(title)), len(message))
	termW, termH := termbox.Size()
	h := 5
	w := min(termW, 4+textW)

	px := (termW - w) / 2
	py := (termH - h) / 2

	// Save background
	buf := saveBlock(px, py, w, h)
	savedCrsrX, savedCrsrY := termbox.GetCursor()

	FillBox(px, py, w, h, ColLightCyan, ColBlue, BorderStyle_Double)

	Puts(px+int((w-len(title))/2), py, " "+title+" ", ColWhite, ColBlue)
	Puts(px+int((w-len(message))/2), py+1, message, ColWhite, ColBlue)
	Puts(px+int((w-len(buttons))/2), py+3, buttons, ColLightBlue, ColBlue)
	termbox.Flush()

	quit := false
	for !quit {
		data := make([]byte, 30)
		termbox.PollRawEvent(data)
		ev := termbox.ParseEvent(data)
		switch ev.Type {
		case termbox.EventKey:
			switch {
			case ev.Mod == termbox.ModAlt && ev.Key == 0 && ev.Ch == 0:
				// Only ESC pressed, nothing else
				res = false
				valid = false
				quit = true
			case ev.Ch == 'y' || ev.Ch == 'Y':
				res = true
				valid = true
				quit = true
			case ev.Ch == 'n' || ev.Ch == 'N':
				res = false
				valid = true
				quit = true
			}
		}
	}

	restoreBlock(buf)
	termbox.SetCursor(savedCrsrX, savedCrsrY)
	termbox.Flush()

	return
}

func ErrorDialog(message string) {
	title := " Somethings's gone wrong... "
	buttons := "<Return>"
	termW, termH := termbox.Size()
	h := 6
	w := min(termW, 4+max(len(title), len(message)))

	px := (termW - w) / 2
	py := (termH - h) / 2

	// Save background
	buf := saveBlock(px, py, w, h)
	savedCrsrX, savedCrsrY := termbox.GetCursor()

	FillBox(px, py, w, h, ColLightRed, ColRed, BorderStyle_Double)
	Puts(px+int((w-len(title))/2), py, " "+title+" ", ColWhite, ColRed)
	Puts(px+int((w-len(message))/2), py+2, message, ColYellow, ColRed)
	Puts(px+int((w-len(buttons))/2), py+4, buttons, ColWhite, ColRed)
	termbox.Flush()

	quit := false
	for !quit {
		data := make([]byte, 30)
		termbox.PollRawEvent(data)
		ev := termbox.ParseEvent(data)
		switch ev.Type {
		case termbox.EventKey:
			switch {
			case ev.Mod == termbox.ModAlt && ev.Key == 0 && ev.Ch == 0:
				// Only ESC pressed, nothing else
				quit = true
			case ev.Key == termbox.KeyEnter:
				quit = true
			}
		}
	}

	restoreBlock(buf)
	termbox.SetCursor(savedCrsrX, savedCrsrY)
	termbox.Flush()
}
