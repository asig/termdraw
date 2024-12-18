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
	"log"
	"unicode/utf8"

	"github.com/asig/termbox-go"
)

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

func (d Direction) Inverse() Direction {
	switch d {
	case DirUp:
		return DirDown
	case DirDown:
		return DirUp
	case DirLeft:
		return DirRight
	case DirRight:
		return DirLeft
	}
	panic("Bad Direction!")
}

const (
	canvasWidth  = 2048
	canvasHeight = 2048
)

type Pos struct {
	X, Y int
}

type cell struct {
	ch     rune
	tile   Tile
	fg, bg termbox.Attribute
}

type Canvas struct {
	// Width and Height of the gadget
	w, h int

	// Position of the gadget on the screen
	pX, pY int

	// Cursor coords relative to canvas
	cX, cY int

	// Top-Left coords of the visible canvas
	ofsX, ofsY int

	cells [][]cell
	//tiles [][]Tile
	//cells [][]termbox.Cell

	fg, bg termbox.Attribute
}

func NewCanvas(x, y int, w, h int) *Canvas {
	c := &Canvas{
		pX:    x,
		pY:    y,
		w:     w,
		h:     h,
		ofsX:  0,
		ofsY:  0,
		fg:    ColLightGrey,
		bg:    ColBlack,
		cells: make([][]cell, canvasHeight),
	}
	c.Clear()
	return c
}

func (c *Canvas) Clear() {
	for i := 0; i < canvasHeight; i++ {
		c.cells[i] = make([]cell, canvasWidth)
		for j := 0; j < canvasWidth; j++ {
			c.cells[i][j] = cell{
				ch:   ' ',
				tile: 0,
				fg:   c.fg,
				bg:   c.bg,
			}
		}
	}
}

func (c *Canvas) AsText() []string {
	var text []string

	for y := 0; y < canvasHeight; y++ {
		line := ""
		for x := 0; x < canvasWidth; x++ {
			line = line + string(c.cells[y][x].ch)
		}
		text = append(text, line)
	}
	return text
}

func (c *Canvas) SetText(text []string) {
	c.Clear()
	if len(text) > canvasHeight {
		text = text[:canvasHeight]
	}

	for y, l := range text {
		for x, i, w := 0, 0, 0; i < len(l); i += w {
			ch, width := utf8.DecodeRuneInString(l[i:])
			c.cells[y][x].ch = ch
			c.cells[y][x].tile = TileFromRune(ch)
			c.cells[y][x].fg = c.fg
			c.cells[y][x].bg = c.bg
			x++
			w = width
		}
	}
}

func (c *Canvas) IncSize(dw, dh int) {
	c.w += dw
	c.h += dh
}

func (c *Canvas) Draw() {
	for y := 0; y < c.h; y++ {
		for x := 0; x < c.w; x++ {
			cell := c.cells[c.ofsY+y][c.ofsX+x]
			termbox.SetCell(c.pX+x, c.pY+y, cell.ch, cell.fg, cell.bg)
		}
	}
	termbox.SetCursor(c.pX+(c.cX-c.ofsX), c.pY+(c.cY-c.ofsY))
}

func (c *Canvas) SetTile(p Pos, t Tile) {
	c.cells[p.Y][p.X].tile = t
	ch := t.Rune()
	if ch != ' ' {
		c.cells[p.Y][p.X].ch = ch
	}
}

func (c *Canvas) Tile(p Pos) Tile {
	return c.cells[p.Y][p.X].tile
}

func (c *Canvas) Pos() Pos {
	return Pos{c.cX, c.cY}
}

//
// Commands
//

func (c *Canvas) Move(d Direction) (oldPos, newPos Pos) {
	oldPos = Pos{X: c.cX, Y: c.cY}
	switch d {
	case DirUp:
		if c.cY > 0 {
			c.cY--
		}
	case DirDown:
		if c.cY < canvasHeight-1 {
			c.cY++
		}
	case DirLeft:
		if c.cX > 0 {
			c.cX--
		}
	case DirRight:
		if c.cX < canvasWidth-1 {
			c.cX++
		}
	}
	newPos = Pos{X: c.cX, Y: c.cY}

	// Adjust the "camera"
	if c.cX < c.ofsX {
		c.ofsX = c.cX
	}
	if c.cX >= c.ofsX+c.w {
		c.ofsX = c.cX - c.w + 1
	}
	if c.cY < c.ofsY {
		c.ofsY = c.cY
	}
	if c.cY >= c.ofsY+c.h {
		c.ofsY = c.cY - c.h + 1
	}

	return oldPos, newPos
}

func (c *Canvas) SetPos(p Pos) {
	c.cX = p.X
	c.cY = p.Y
	if c.cX < 0 {
		c.cX = 0
	}
	if c.cX >= canvasWidth {
		c.cX = canvasWidth - 1
	}
	if c.cY < 0 {
		c.cY = 0
	}
	if c.cY >= canvasHeight {
		c.cY = canvasHeight - 1
	}
}

func (c *Canvas) Insert(p Pos) {
	l := len(c.cells[p.Y])
	for i := l - 2; i >= p.X; i-- {
		c.cells[p.Y][i+1] = c.cells[p.Y][i]
	}
}

func (c *Canvas) Delete(p Pos) {
	l := len(c.cells[p.Y])
	for i := p.X + 1; i < l; i++ {
		c.cells[p.Y][i-1] = c.cells[p.Y][i]
	}
	c.cells[p.Y][l-1] = cell{ch: ' ', fg: c.fg, bg: c.bg, tile: 0}

}

func (c *Canvas) InsertLine(p Pos) {
	l := len(c.cells)
	for i := l - 2; i >= p.Y; i-- {
		c.cells[i+1] = c.cells[i]
	}
	c.cells[p.Y] = make([]cell, canvasWidth)
	for x := 0; x < len(c.cells[p.Y]); x++ {
		c.cells[p.Y][x] = cell{ch: ' ', fg: c.fg, bg: c.bg, tile: 0}
	}
}

func (c *Canvas) DeleteLine(p Pos) {
	l := len(c.cells)
	for i := p.Y + 1; i < l; i++ {
		c.cells[i-1] = c.cells[i]
	}
	y := canvasHeight - 1
	c.cells[y] = make([]cell, canvasWidth)
	for x := 0; x < len(c.cells[y]); x++ {
		c.cells[y][x] = cell{ch: ' ', fg: c.fg, bg: c.bg, tile: 0}
	}

}

func (c *Canvas) SetRune(p Pos, ch rune) {
	if ch == 0 {
		log.Fatalf("NULL!!!")
	}
	c.cells[p.Y][p.X].ch = ch
	c.cells[p.Y][p.X].tile = 0
}
