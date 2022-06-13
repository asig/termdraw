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
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/asig/termbox-go"

	"github.com/asig/termdraw/pkg/termdraw"
)

var (
	termW  int
	termH  int
	canvas *termdraw.Canvas

	curBorderStyle termdraw.BorderStyle
	curFilename    string
	insert         bool
	dirty          bool
)

func showWelcome() {
	y := termdraw.ColYellow
	yb := y | termbox.AttrBold
	lr := termdraw.ColLightRed
	bg := termdraw.ColRed
	t := &termdraw.TextCard{
		Fg: lr,
		Bg: bg,
		Bs: termdraw.BorderStyle_Rounded,
		Content: []termdraw.Line{
			{{"                               ", y, bg}},
			{{"    ***********************    ", yb, bg}},
			{{"    *** T E R M D R A W ***    ", yb, bg}},
			{{"    ***********************    ", yb, bg}},
			{{"                               ", y, bg}},
			{{"  Draw like it's the nineties  ", lr, bg}},
			{{"                               ", lr, bg}},
			{{"   Press <Ctrl>+<H> for help   ", lr, bg}},
			{{"   Press <Ctrl>+<X> to quit    ", lr, bg}},
			{{"                               ", lr, bg}},
		},
	}
	t.Show()
}

func showHelp() {
	y := termdraw.ColYellow
	w := termdraw.ColWhite | termbox.AttrBold
	yb := y | termbox.AttrBold
	bg := termdraw.ColBrown
	t := &termdraw.TextCard{
		Fg: y,
		Bg: bg,
		Bs: termdraw.BorderStyle_Rounded,
		Content: []termdraw.Line{
			{{"                      * * * H E L P * * *                      ", yb, bg}},
			{{"", y, bg}},
			{{"Termdraw is focussed on drawing Unicode-based borders in text", y, bg}},
			{{"files. To do so, press ", y, bg}, {"Ctrl-B", w, bg}, {" to pick the border style, and then", y, bg}},
			{{"use the ", y, bg}, {"Cursor keys", w, bg}, {" to draw the border.", y, bg}},
			{{"", y, bg}},
			{{"Besides that, it pretty  much works like a regular text editor.", y, bg}},
			{{"", y, bg}},
			{{"Commands", y, bg}},
			{{"────────", y, bg}},
			{{"Ctrl-B ", w, bg}, {"Select border style", y, bg}},
			{{"Ctrl-Q ", w, bg}, {"Quit", y, bg}},
			{{"", y, bg}},
			{{"Ctrl-H ", w, bg}, {"Show this dialog", y, bg}},
			{{"", y, bg}},
			{{"Ctrl-O ", w, bg}, {"Load a text file", y, bg}},
			{{"Ctrl-S ", w, bg}, {"Save as a text file", y, bg}},
			{{"Ctrl-I ", w, bg}, {"Insert a line", y, bg}},
			{{"Ctrl-D ", w, bg}, {"Delete current line", y, bg}},
		},
	}
	t.Show()
}

func drawStatusbar() {
	p := canvas.Pos()
	ins := "OVW"
	if insert {
		ins = "INS"
	}
	status := fmt.Sprintf(" Pos: %d/%d | %s | Border style: %s ", p.X, p.Y, ins, curBorderStyle)
	if curFilename != "" || dirty {
		var filepart string
		if dirty {
			filepart += "*"
		}
		if curFilename != "" {
			filepart += curFilename
		}
		status = " " + filepart + " |" + status
	}
	if len(status) < termW {
		status = status + strings.Repeat(" ", termW-len(status))
	}

	termdraw.Puts(0, termH-1, status, termdraw.ColCyan, termdraw.ColBlue)
}

func handleMove(dir termdraw.Direction) {
	oldPos, newPos := canvas.Move(dir)
	if curBorderStyle == termdraw.BorderStyle_None {
		return
	}
	if oldPos != newPos {
		// Leaving the current tile in the direction we're moving
		canvas.SetTile(oldPos, canvas.Tile(oldPos).WithDir(dir, curBorderStyle))
		// Entering the new tile from the inverse diurection
		canvas.SetTile(newPos, canvas.Tile(newPos).WithDir(dir.Inverse(), curBorderStyle))
		dirty = true
	}
}

func handleChar(ch rune) {
	p := canvas.Pos()
	if insert {
		canvas.Insert(p)
	}
	canvas.SetRune(p, ch)
	canvas.Move(termdraw.DirRight)
	dirty = true
}

func handleBackspace() {
	p := canvas.Pos()
	if p.X == 0 {
		return
	}
	canvas.Delete(termdraw.Pos{p.X - 1, p.Y})
	canvas.Move(termdraw.DirLeft)
	dirty = true
}

func handleEnter() {
	p := canvas.Pos()
	canvas.SetPos(termdraw.Pos{
		X: 0,
		Y: p.Y + 1,
	})
}

func handleDelete() {
	p := canvas.Pos()
	canvas.Delete(p)
	dirty = true
}

func handleInsertLine() {
	p := canvas.Pos()
	canvas.InsertLine(p)
	dirty = true
}

func handleDeleteLine() {
	p := canvas.Pos()
	canvas.DeleteLine(p)
	dirty = true
}

func saveCanvas() {
	text := canvas.AsText()
	for i, _ := range text {
		text[i] = strings.TrimRightFunc(text[i], unicode.IsSpace)
	}
	end := len(text)
	for end > 0 && len(text[end-1]) == 0 {
		end--
	}
	text = text[:end]
	ioutil.WriteFile(curFilename, []byte(strings.Join(text, "\n")), 0644)
}

func loadCanvas(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	canvas.SetText(lines)
	return nil
}

func handleSave() {
	if curFilename == "" {
		f, ok := termdraw.FileDialog("Save File")
		if !ok {
			return
		}
		curFilename = f
	}

	saveCanvas()

	dirty = false
}

func handleLoad() {
	if dirty {
		res, valid := termdraw.YesNoCancelDialog("Save?", "Text is modified. Save it first?")
		if !valid {
			return
		}
		if res {
			handleSave()
		}
	}
	f, ok := termdraw.FileDialog("Open File")
	if !ok {
		return
	}
	err := loadCanvas(f)
	if err != nil {
		termdraw.ErrorDialog(err.Error())
		return
	}

	dirty = false
}

func maybeSave() bool {
	res, valid := termdraw.YesNoCancelDialog("Save?", "Text is modified. Save it?")
	if !valid {
		return false
	}
	if res {
		handleSave()
	}
	return true
}

func handleEvent(ev termbox.Event) (quit bool, helpShown bool) {
	switch ev.Type {
	case termbox.EventResize:
		newW, newH := termbox.Size()
		deltaW := newW - termW
		deltaH := newH - termH
		s := fmt.Sprintf("old size: %d x %d; new size %d x %d,  deltaW = %d, deltaH = %d", termW, termH, newW, newH, deltaW, deltaH)
		termdraw.Puts(10, 0, s, termbox.AttrBold|termbox.ColorWhite, termbox.ColorBlue)
		termW = newW
		termH = newH
		canvas.IncSize(deltaW, deltaH)

	case termbox.EventKey:
		switch {
		case ev.Key == termbox.KeyCtrlX:
			quit = true
			if dirty {
				quit = maybeSave()
			}
		case ev.Key == termbox.KeyCtrlB:
			curBorderStyle = curBorderStyle.Next()
		case ev.Key == termbox.KeyCtrlH:
			helpShown = true
			showHelp()
		case ev.Key == termbox.KeyArrowDown:
			handleMove(termdraw.DirDown)
		case ev.Key == termbox.KeyArrowUp:
			handleMove(termdraw.DirUp)
		case ev.Key == termbox.KeyArrowLeft:
			handleMove(termdraw.DirLeft)
		case ev.Key == termbox.KeyArrowRight:
			handleMove(termdraw.DirRight)
		case ev.Key == termbox.KeyInsert:
			insert = !insert
		case ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2:
			handleBackspace()
		case ev.Key == termbox.KeyDelete:
			handleDelete()
		case ev.Key == termbox.KeyEnter:
			handleEnter()
		case ev.Key == termbox.KeyCtrlS:
			handleSave()
		case ev.Key == termbox.KeyCtrlO:
			handleLoad()
		case ev.Key == termbox.KeyCtrlI:
			handleInsertLine()
		case ev.Key == termbox.KeyCtrlD:
			handleDeleteLine()
		case unicode.IsPrint(ev.Ch) || ev.Key == ' ':
			handleChar(ev.Ch)
		}
		//case termbox.EventMouse:
		//	if ev.Key == termbox.MouseLeft {
		//		//mx, _ = ev.MouseX, ev.MouseY
		//		//count = mx
		//	}
		// case termbox.EventResize:
		// 	reallocBackBuffer(ev.Width, ev.Height)
	}
	return quit, helpShown
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	termbox.SetOutputMode(termbox.OutputRGB)

	termW, termH = termbox.Size()
	canvas = termdraw.NewCanvas(0, 0, termW, termH-1)

	curFilename = ""

	drawStatusbar()
	if len(os.Args) > 1 {
		err = loadCanvas(os.Args[1])
		if err == nil {
			curFilename = os.Args[1]
		}
	}
	canvas.Draw()
	if err != nil {
		termdraw.ErrorDialog(err.Error())
	} else {
		showWelcome()
		termbox.Flush()
	}

	curBorderStyle = termdraw.BorderStyle_None
	dirty = false

	eventBuf := make([]byte, 20)
	quit := false
	for !quit {
		helpShown := false

		ev := termbox.PollRawEvent(eventBuf)
		if ev.Type == termbox.EventRaw {
			ev = termbox.ParseEvent(eventBuf)
		}
		quit, helpShown = handleEvent(ev)
		canvas.Draw()
		drawStatusbar()
		if helpShown {
			showHelp()
		}
		termbox.Flush()
	}
}
