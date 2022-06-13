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

type (
	Border      [][]rune
	BorderStyle uint8
)

const (
	BorderStyle_None BorderStyle = iota
	BorderStyle_Light
	BorderStyle_Rounded
	BorderStyle_Heavy
	BorderStyle_Double

	BorderStyle_Max = BorderStyle_Double
)

var (
	borderMap = map[BorderStyle]Border{
		BorderStyle_Light: {
			{'┌', '─', '┬', '─', '┐'},
			{'│', ' ', '│', ' ', '│'},
			{'├', '─', '┼', '─', '┤'},
			{'│', ' ', '│', ' ', '│'},
			{'└', '─', '┴', '─', '┘'},
		},
		BorderStyle_Rounded: {
			{'╭', '─', '┬', '─', '╮'},
			{'│', ' ', '│', ' ', '│'},
			{'├', '─', '┼', '─', '┤'},
			{'│', ' ', '│', ' ', '│'},
			{'╰', '─', '┴', '─', '╯'},
		},
		BorderStyle_Heavy: {
			{'┏', '━', '┳', '━', '┓'},
			{'┃', ' ', '┃', ' ', '┃'},
			{'┣', '━', '╋', '━', '┫'},
			{'┃', ' ', '┃', ' ', '┃'},
			{'┗', '━', '┻', '━', '┛'},
		},
		BorderStyle_Double: {
			{'╔', '═', '╦', '═', '╗'},
			{'║', ' ', '║', ' ', '║'},
			{'╠', '═', '╬', '═', '╣'},
			{'║', ' ', '║', ' ', '║'},
			{'╚', '═', '╩', '═', '╝'},
		},
		BorderStyle_None:
		{
			{' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' '},
		},
	}

	borderNames = map[BorderStyle]string{
		BorderStyle_None:    "None",
		BorderStyle_Light:   "Light",
		BorderStyle_Rounded: "Rounded",
		BorderStyle_Heavy:   "Heavy",
		BorderStyle_Double:  "Double",
	}
)

func (b BorderStyle) Prev() BorderStyle {
	prev := b - 1
	if prev < 0 {
		prev = BorderStyle_Max
	}
	return prev
}

func (b BorderStyle) Next() BorderStyle {
	next := b + 1
	if next > BorderStyle_Max {
		next = 0
	}
	return next
}

func (b BorderStyle) Runes() Border {
	return borderMap[b]
}

func (b BorderStyle) String() string {
	return borderNames[b]
}
