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
	"strings"
)

var (
	// Unicode info taken from https://www.fileformat.info/info/unicode/block/box_drawing/list.htm
	names = map[string]rune{
		"LIGHT HORIZONTAL":                      '─',
		"HEAVY HORIZONTAL":                      '━',
		"LIGHT VERTICAL":                        '│',
		"HEAVY VERTICAL":                        '┃',
		"LIGHT DOWN AND RIGHT":                  '┌',
		"DOWN LIGHT AND RIGHT HEAVY":            '┍',
		"DOWN HEAVY AND RIGHT LIGHT":            '┎',
		"HEAVY DOWN AND RIGHT":                  '┏',
		"LIGHT DOWN AND LEFT":                   '┐',
		"DOWN LIGHT AND LEFT HEAVY":             '┑',
		"DOWN HEAVY AND LEFT LIGHT":             '┒',
		"HEAVY DOWN AND LEFT":                   '┓',
		"LIGHT UP AND RIGHT":                    '└',
		"UP LIGHT AND RIGHT HEAVY":              '┕',
		"UP HEAVY AND RIGHT LIGHT":              '┖',
		"HEAVY UP AND RIGHT":                    '┗',
		"LIGHT UP AND LEFT":                     '┘',
		"UP LIGHT AND LEFT HEAVY":               '┙',
		"UP HEAVY AND LEFT LIGHT":               '┚',
		"HEAVY UP AND LEFT":                     '┛',
		"LIGHT VERTICAL AND RIGHT":              '├',
		"VERTICAL LIGHT AND RIGHT HEAVY":        '┝',
		"UP HEAVY AND RIGHT DOWN LIGHT":         '┞',
		"DOWN HEAVY AND RIGHT UP LIGHT":         '┟',
		"VERTICAL HEAVY AND RIGHT LIGHT":        '┠',
		"DOWN LIGHT AND RIGHT UP HEAVY":         '┡',
		"UP LIGHT AND RIGHT DOWN HEAVY":         '┢',
		"HEAVY VERTICAL AND RIGHT":              '┣',
		"LIGHT VERTICAL AND LEFT":               '┤',
		"VERTICAL LIGHT AND LEFT HEAVY":         '┥',
		"UP HEAVY AND LEFT DOWN LIGHT":          '┦',
		"DOWN HEAVY AND LEFT UP LIGHT":          '┧',
		"VERTICAL HEAVY AND LEFT LIGHT":         '┨',
		"DOWN LIGHT AND LEFT UP HEAVY":          '┩',
		"UP LIGHT AND LEFT DOWN HEAVY":          '┪',
		"HEAVY VERTICAL AND LEFT":               '┫',
		"LIGHT DOWN AND HORIZONTAL":             '┬',
		"LEFT HEAVY AND RIGHT DOWN LIGHT":       '┭',
		"RIGHT HEAVY AND LEFT DOWN LIGHT":       '┮',
		"DOWN LIGHT AND HORIZONTAL HEAVY":       '┯',
		"DOWN HEAVY AND HORIZONTAL LIGHT":       '┰',
		"RIGHT LIGHT AND LEFT DOWN HEAVY":       '┱',
		"LEFT LIGHT AND RIGHT DOWN HEAVY":       '┲',
		"HEAVY DOWN AND HORIZONTAL":             '┳',
		"LIGHT UP AND HORIZONTAL":               '┴',
		"LEFT HEAVY AND RIGHT UP LIGHT":         '┵',
		"RIGHT HEAVY AND LEFT UP LIGHT":         '┶',
		"UP LIGHT AND HORIZONTAL HEAVY":         '┷',
		"UP HEAVY AND HORIZONTAL LIGHT":         '┸',
		"RIGHT LIGHT AND LEFT UP HEAVY":         '┹',
		"LEFT LIGHT AND RIGHT UP HEAVY":         '┺',
		"HEAVY UP AND HORIZONTAL":               '┻',
		"LIGHT VERTICAL AND HORIZONTAL":         '┼',
		"LEFT HEAVY AND RIGHT VERTICAL LIGHT":   '┽',
		"RIGHT HEAVY AND LEFT VERTICAL LIGHT":   '┾',
		"VERTICAL LIGHT AND HORIZONTAL HEAVY":   '┿',
		"UP HEAVY AND DOWN HORIZONTAL LIGHT":    '╀',
		"DOWN HEAVY AND UP HORIZONTAL LIGHT":    '╁',
		"VERTICAL HEAVY AND HORIZONTAL LIGHT":   '╂',
		"LEFT UP HEAVY AND RIGHT DOWN LIGHT":    '╃',
		"RIGHT UP HEAVY AND LEFT DOWN LIGHT":    '╄',
		"LEFT DOWN HEAVY AND RIGHT UP LIGHT":    '╅',
		"RIGHT DOWN HEAVY AND LEFT UP LIGHT":    '╆',
		"DOWN LIGHT AND UP HORIZONTAL HEAVY":    '╇',
		"UP LIGHT AND DOWN HORIZONTAL HEAVY":    '╈',
		"RIGHT LIGHT AND LEFT VERTICAL HEAVY":   '╉',
		"LEFT LIGHT AND RIGHT VERTICAL HEAVY":   '╊',
		"HEAVY VERTICAL AND HORIZONTAL":         '╋',
		"DOUBLE HORIZONTAL":                     '═',
		"DOUBLE VERTICAL":                       '║',
		"DOWN SINGLE AND RIGHT DOUBLE":          '╒',
		"DOWN DOUBLE AND RIGHT SINGLE":          '╓',
		"DOUBLE DOWN AND RIGHT":                 '╔',
		"DOWN SINGLE AND LEFT DOUBLE":           '╕',
		"DOWN DOUBLE AND LEFT SINGLE":           '╖',
		"DOUBLE DOWN AND LEFT":                  '╗',
		"UP SINGLE AND RIGHT DOUBLE":            '╘',
		"UP DOUBLE AND RIGHT SINGLE":            '╙',
		"DOUBLE UP AND RIGHT":                   '╚',
		"UP SINGLE AND LEFT DOUBLE":             '╛',
		"UP DOUBLE AND LEFT SINGLE":             '╜',
		"DOUBLE UP AND LEFT":                    '╝',
		"VERTICAL SINGLE AND RIGHT DOUBLE":      '╞',
		"VERTICAL DOUBLE AND RIGHT SINGLE":      '╟',
		"DOUBLE VERTICAL AND RIGHT":             '╠',
		"VERTICAL SINGLE AND LEFT DOUBLE":       '╡',
		"VERTICAL DOUBLE AND LEFT SINGLE":       '╢',
		"DOUBLE VERTICAL AND LEFT":              '╣',
		"DOWN SINGLE AND HORIZONTAL DOUBLE":     '╤',
		"DOWN DOUBLE AND HORIZONTAL SINGLE":     '╥',
		"DOUBLE DOWN AND HORIZONTAL":            '╦',
		"UP SINGLE AND HORIZONTAL DOUBLE":       '╧',
		"UP DOUBLE AND HORIZONTAL SINGLE":       '╨',
		"DOUBLE UP AND HORIZONTAL":              '╩',
		"VERTICAL SINGLE AND HORIZONTAL DOUBLE": '╪',
		"VERTICAL DOUBLE AND HORIZONTAL SINGLE": '╫',
		"DOUBLE VERTICAL AND HORIZONTAL":        '╬',
		"LIGHT ARC DOWN AND RIGHT":              '╭',
		"LIGHT ARC DOWN AND LEFT":               '╮',
		"LIGHT ARC UP AND LEFT":                 '╯',
		"LIGHT ARC UP AND RIGHT":                '╰',
		"LIGHT LEFT":                            '╴',
		"LIGHT UP":                              '╵',
		"LIGHT RIGHT":                           '╶',
		"LIGHT DOWN":                            '╷',
		"HEAVY LEFT":                            '╸',
		"HEAVY UP":                              '╹',
		"HEAVY RIGHT":                           '╺',
		"HEAVY DOWN":                            '╻',
		"LIGHT LEFT AND HEAVY RIGHT":            '╼',
		"LIGHT UP AND HEAVY DOWN":               '╽',
		"HEAVY LEFT AND LIGHT RIGHT":            '╾',
		"HEAVY UP AND LIGHT DOWN":               '╿',
		/*
			"LIGHT TRIPLE DASH HORIZONTAL":             '┄',
			"HEAVY TRIPLE DASH HORIZONTAL":             '┅',
			"LIGHT TRIPLE DASH VERTICAL":               '┆',
			"HEAVY TRIPLE DASH VERTICAL":               '┇',
			"LIGHT QUADRUPLE DASH HORIZONTAL":          '┈',
			"HEAVY QUADRUPLE DASH HORIZONTAL":          '┉',
			"LIGHT QUADRUPLE DASH VERTICAL":            '┊',
			"HEAVY QUADRUPLE DASH VERTICAL":            '┋',
			"LIGHT DOUBLE DASH HORIZONTAL":             '╌',
			"HEAVY DOUBLE DASH HORIZONTAL":             '╍',
			"LIGHT DOUBLE DASH VERTICAL":               '╎',
			"HEAVY DOUBLE DASH VERTICAL":               '╏',
			"LIGHT DIAGONAL UPPER RIGHT TO LOWER LEFT": '╱',
			"LIGHT DIAGONAL UPPER LEFT TO LOWER RIGHT": '╲',
			"LIGHT DIAGONAL CROSS":                     '╳',

		*/
	}
)

func isWeight(s string) bool {
	return s == "LIGHT" || s == "HEAVY" || s == "DOUBLE"
}

func splitFirst(s string) (first, remainder string) {
	parts := strings.SplitN(s, " ", 2)
	return parts[0], parts[1]
}

func parsePart(part string, curWeight byte, pattern []byte) byte {
	elems := strings.Split(part, " ")

	// Find weight
	for _, e := range elems {
		switch e {
		case "SINGLE", "LIGHT":
			curWeight = 'l'
		case "ARC":
			curWeight = 'r'
		case "HEAVY":
			curWeight = 'h'
		case "DOUBLE":
			curWeight = 'd'
		}
	}

	// handle directions
	for _, e := range elems {
		switch e {
		case "LEFT":
			pattern[0] = curWeight
		case "UP":
			pattern[1] = curWeight
		case "RIGHT":
			pattern[2] = curWeight
		case "DOWN":
			pattern[3] = curWeight
		case "HORIZONTAL":
			pattern[0] = curWeight
			pattern[2] = curWeight
		case "VERTICAL":
			pattern[1] = curWeight
			pattern[3] = curWeight
		}
	}

	return curWeight
}

func substituteRound(pattern []byte) []string {
	var seen = make(map[string]bool)

	s := string(pattern)
	if s == "  rr" || s == " rr " || s == "rr  " || s == "r  r" {
		// no need to substitute
		return []string { s }
	}

	// all 4 bit combinations, for every bit set, substitute 'l' with 'r'
	for i := 0; i < 32; i++ {
		p := make([]byte, 4)
		copy(p, pattern)

		for bit := 0; bit < 4; bit++ {
			if (i & (1<<bit)) == (1<<bit) && p[bit] == 'l' {
				p[bit] = 'r'
			}
		}
		s := string(p)
		if s == "  rr" || s == " rr " || s == "rr  " || s == "r  r" {
			// Not a valid substitution
			continue
		}
		seen[s] = true
	}


	var res []string
	for key, _ := range seen {
		res = append(res, key)
	}
	return res
}



func parse(name string, ch rune) {
	pattern := []byte{' ', ' ', ' ', ' '}

	var curWeight byte = 'l'
	parts := strings.Split(name, " AND ")
	for _, p := range parts {
		curWeight = parsePart(p, curWeight, pattern)
	}

	patterns := substituteRound(pattern);
	for _, p := range patterns {
		fmt.Printf("newTile(\"%s\"): '%c',\n", p, ch)
	}
}

func main() {
	for key, val := range names {
		parse(key, val)
	}
}
