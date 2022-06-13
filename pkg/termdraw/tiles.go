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
	"fmt"
)

type (
	Tile uint32
)

var (
	// generated with gentiles.go
	tileToRune = map[Tile]rune{
		newTile(" rh "): '┕',
		newTile(" lh "): '┕',
		newTile(" rdl"): '╞',
		newTile(" ldr"): '╞',
		newTile(" rdr"): '╞',
		newTile(" ldl"): '╞',
		newTile("h ll"): '┭',
		newTile("h rl"): '┭',
		newTile("h lr"): '┭',
		newTile("h rr"): '┭',
		newTile("r hl"): '┮',
		newTile("l hr"): '┮',
		newTile("r hr"): '┮',
		newTile("l hl"): '┮',
		newTile("d dd"): '╦',
		newTile("h   "): '╸',
		newTile("l  h"): '┒',
		newTile("r  h"): '┒',
		newTile("d dl"): '╤',
		newTile("d dr"): '╤',
		newTile("rhrl"): '╀',
		newTile("lhlr"): '╀',
		newTile("rhlr"): '╀',
		newTile("lhrr"): '╀',
		newTile("rhrr"): '╀',
		newTile("lhll"): '╀',
		newTile("rhll"): '╀',
		newTile("lhrl"): '╀',
		newTile("hhrr"): '╃',
		newTile("hhll"): '╃',
		newTile("hhrl"): '╃',
		newTile("hhlr"): '╃',
		newTile("lhhl"): '╄',
		newTile("rhhl"): '╄',
		newTile("lhhr"): '╄',
		newTile("rhhr"): '╄',
		newTile(" ld "): '╘',
		newTile(" rd "): '╘',
		newTile("  lh"): '┎',
		newTile("  rh"): '┎',
		newTile(" hh "): '┗',
		newTile("ll  "): '┘',
		newTile("rl  "): '┘',
		newTile("lr  "): '┘',
		newTile("hhhh"): '╋',
		newTile("l h "): '╼',
		newTile("r h "): '╼',
		newTile(" l l"): '│',
		newTile(" r l"): '│',
		newTile(" l r"): '│',
		newTile(" r r"): '│',
		newTile("r rr"): '┬',
		newTile("l ll"): '┬',
		newTile("r ll"): '┬',
		newTile("l rl"): '┬',
		newTile("r rl"): '┬',
		newTile("l lr"): '┬',
		newTile("r lr"): '┬',
		newTile("l rr"): '┬',
		newTile(" rlr"): '├',
		newTile(" lrr"): '├',
		newTile(" rrr"): '├',
		newTile(" lll"): '├',
		newTile(" rll"): '├',
		newTile(" lrl"): '├',
		newTile(" rrl"): '├',
		newTile(" llr"): '├',
		newTile("hhlh"): '╉',
		newTile("hhrh"): '╉',
		newTile("dddd"): '╬',
		newTile("hl l"): '┥',
		newTile("hr l"): '┥',
		newTile("hl r"): '┥',
		newTile("hr r"): '┥',
		newTile("lhr "): '┸',
		newTile("rhr "): '┸',
		newTile("lhl "): '┸',
		newTile("rhl "): '┸',
		newTile("hllh"): '╅',
		newTile("hrlh"): '╅',
		newTile("hlrh"): '╅',
		newTile("hrrh"): '╅',
		newTile("ld  "): '╜',
		newTile("rd  "): '╜',
		newTile("dldl"): '╪',
		newTile("drdl"): '╪',
		newTile("dldr"): '╪',
		newTile("drdr"): '╪',
		newTile(" rhh"): '┢',
		newTile(" lhh"): '┢',
		newTile("d  l"): '╕',
		newTile("d  r"): '╕',
		newTile("   l"): '╷',
		newTile("   r"): '╷',
		newTile("hh h"): '┫',
		newTile("h hh"): '┳',
		newTile("ldl "): '╨',
		newTile("rdl "): '╨',
		newTile("ldr "): '╨',
		newTile("rdr "): '╨',
		newTile("l   "): '╴',
		newTile("r   "): '╴',
		newTile("lrll"): '┼',
		newTile("rrlr"): '┼',
		newTile("llrl"): '┼',
		newTile("rlrl"): '┼',
		newTile("lrlr"): '┼',
		newTile("llrr"): '┼',
		newTile("rlrr"): '┼',
		newTile("llll"): '┼',
		newTile("rlll"): '┼',
		newTile("rrll"): '┼',
		newTile("lrrl"): '┼',
		newTile("lllr"): '┼',
		newTile("rrrr"): '┼',
		newTile("rrrl"): '┼',
		newTile("rllr"): '┼',
		newTile("lrrr"): '┼',
		newTile("  dl"): '╒',
		newTile("  dr"): '╒',
		newTile("   h"): '╻',
		newTile(" lhl"): '┝',
		newTile(" rhl"): '┝',
		newTile(" lhr"): '┝',
		newTile(" rhr"): '┝',
		newTile("ll r"): '┤',
		newTile("rl r"): '┤',
		newTile("lr r"): '┤',
		newTile("rr r"): '┤',
		newTile("ll l"): '┤',
		newTile("rl l"): '┤',
		newTile("lr l"): '┤',
		newTile("rr l"): '┤',
		newTile("hlhl"): '┿',
		newTile("hrhl"): '┿',
		newTile("hlhr"): '┿',
		newTile("hrhr"): '┿',
		newTile("rhrh"): '╂',
		newTile("lhlh"): '╂',
		newTile("rhlh"): '╂',
		newTile("lhrh"): '╂',
		newTile("h  l"): '┑',
		newTile("h  r"): '┑',
		newTile("rlr "): '┴',
		newTile("lrr "): '┴',
		newTile("rrr "): '┴',
		newTile("lll "): '┴',
		newTile("rll "): '┴',
		newTile("lrl "): '┴',
		newTile("rrl "): '┴',
		newTile("llr "): '┴',
		newTile("hll "): '┵',
		newTile("hrl "): '┵',
		newTile("hlr "): '┵',
		newTile("hrr "): '┵',
		newTile("d  d"): '╗',
		newTile(" rr "): '╰',
		newTile("hh  "): '┛',
		newTile(" hll"): '┞',
		newTile(" hrl"): '┞',
		newTile(" hlr"): '┞',
		newTile(" hrr"): '┞',
		newTile("hh l"): '┩',
		newTile("hh r"): '┩',
		newTile("  l "): '╶',
		newTile("  r "): '╶',
		newTile("  hl"): '┍',
		newTile("  hr"): '┍',
		newTile("hrlr"): '┽',
		newTile("hlrr"): '┽',
		newTile("hrrr"): '┽',
		newTile("hlll"): '┽',
		newTile("hrll"): '┽',
		newTile("hlrl"): '┽',
		newTile("hrrl"): '┽',
		newTile("hllr"): '┽',
		newTile("ldld"): '╫',
		newTile("rdld"): '╫',
		newTile("ldrd"): '╫',
		newTile("rdrd"): '╫',
		newTile("  hh"): '┏',
		newTile("lh h"): '┨',
		newTile("rh h"): '┨',
		newTile("r  r"): '╮',
		newTile(" l  "): '╵',
		newTile(" r  "): '╵',
		newTile(" hhl"): '┡',
		newTile(" hhr"): '┡',
		newTile("hhl "): '┹',
		newTile("hhr "): '┹',
		newTile("lrhl"): '┾',
		newTile("rrhl"): '┾',
		newTile("llhr"): '┾',
		newTile("rlhr"): '┾',
		newTile("lrhr"): '┾',
		newTile("rrhr"): '┾',
		newTile("llhl"): '┾',
		newTile("rlhl"): '┾',
		newTile("l  d"): '╖',
		newTile("r  d"): '╖',
		newTile("rr  "): '╯',
		newTile(" ll "): '└',
		newTile(" rl "): '└',
		newTile(" lr "): '└',
		newTile(" dd "): '╚',
		newTile("lhh "): '┺',
		newTile("rhh "): '┺',
		newTile(" dld"): '╟',
		newTile(" drd"): '╟',
		newTile("dl l"): '╡',
		newTile("dr l"): '╡',
		newTile("dl r"): '╡',
		newTile("dr r"): '╡',
		newTile("drd "): '╧',
		newTile("dld "): '╧',
		newTile("hl  "): '┙',
		newTile("hr  "): '┙',
		newTile("hlh "): '┷',
		newTile("hrh "): '┷',
		newTile("lrlh"): '╁',
		newTile("rrlh"): '╁',
		newTile("llrh"): '╁',
		newTile("rlrh"): '╁',
		newTile("lrrh"): '╁',
		newTile("rrrh"): '╁',
		newTile("lllh"): '╁',
		newTile("rllh"): '╁',
		newTile("r ld"): '╥',
		newTile("l rd"): '╥',
		newTile("r rd"): '╥',
		newTile("l ld"): '╥',
		newTile("rl h"): '┧',
		newTile("lr h"): '┧',
		newTile("rr h"): '┧',
		newTile("ll h"): '┧',
		newTile("hl h"): '┪',
		newTile("hr h"): '┪',
		newTile("h lh"): '┱',
		newTile("h rh"): '┱',
		newTile("d d "): '═',
		newTile("dd  "): '╝',
		newTile("hhh "): '┻',
		newTile("rrhh"): '╆',
		newTile("llhh"): '╆',
		newTile("rlhh"): '╆',
		newTile("lrhh"): '╆',
		newTile("  dd"): '╔',
		newTile(" h l"): '╿',
		newTile(" h r"): '╿',
		newTile("l  l"): '┐',
		newTile("r  l"): '┐',
		newTile("l  r"): '┐',
		newTile(" llh"): '┟',
		newTile(" rlh"): '┟',
		newTile(" lrh"): '┟',
		newTile(" rrh"): '┟',
		newTile(" hhh"): '┣',
		newTile("ld d"): '╢',
		newTile("rd d"): '╢',
		newTile("  h "): '╺',
		newTile(" h  "): '╹',
		newTile("ddd "): '╩',
		newTile("  ll"): '┌',
		newTile("  rl"): '┌',
		newTile("  lr"): '┌',
		newTile(" hrh"): '┠',
		newTile(" hlh"): '┠',
		newTile(" d d"): '║',
		newTile("  ld"): '╓',
		newTile("  rd"): '╓',
		newTile(" ddd"): '╠',
		newTile("h  h"): '┓',
		newTile("lh  "): '┚',
		newTile("rh  "): '┚',
		newTile("l rh"): '┰',
		newTile("r rh"): '┰',
		newTile("l lh"): '┰',
		newTile("r lh"): '┰',
		newTile("dd d"): '╣',
		newTile("  rr"): '╭',
		newTile("l l "): '─',
		newTile("r l "): '─',
		newTile("l r "): '─',
		newTile("r r "): '─',
		newTile(" h h"): '┃',
		newTile("hlhh"): '╈',
		newTile("hrhh"): '╈',
		newTile("lhhh"): '╊',
		newTile("rhhh"): '╊',
		newTile(" dl "): '╙',
		newTile(" dr "): '╙',
		newTile("rh r"): '┦',
		newTile("lh l"): '┦',
		newTile("rh l"): '┦',
		newTile("lh r"): '┦',
		newTile("h r "): '╾',
		newTile("h l "): '╾',
		newTile("hhhl"): '╇',
		newTile("hhhr"): '╇',
		newTile("dl  "): '╛',
		newTile("dr  "): '╛',
		newTile(" l h"): '╽',
		newTile(" r h"): '╽',
		newTile("h h "): '━',
		newTile(" hr "): '┖',
		newTile(" hl "): '┖',
		newTile("h hr"): '┯',
		newTile("h hl"): '┯',
		newTile("l hh"): '┲',
		newTile("r hh"): '┲',
		newTile("llh "): '┶',
		newTile("rlh "): '┶',
		newTile("lrh "): '┶',
		newTile("rrh "): '┶',
	}
)

func borderStyleFromChar(c byte) BorderStyle {
	switch c {
	case ' ':
		return BorderStyle_None
	case 'r':
		return BorderStyle_Rounded
	case 'l':
		return BorderStyle_Light
	case 'h':
		return BorderStyle_Heavy
	case 'd':
		return BorderStyle_Double
	default:
		panic(fmt.Sprintf("Invalid border style character %c", c))
	}
}

func newTile(s string) Tile {
	return Tile(uint32(borderStyleFromChar(s[0]))<<24 | uint32(borderStyleFromChar(s[1]))<<16 | uint32(borderStyleFromChar(s[2]))<<8 | uint32(borderStyleFromChar(s[3])))
}

func tileExchange(t Tile, shift int, border BorderStyle) Tile {
	mask := ^(uint32(0xff) << shift)
	return Tile((uint32(t) & mask) | uint32(border) << shift)
}

func TileFromRune(r rune) Tile {
	for key, val := range tileToRune {
		if val == r {
			return key
		}
	}
	return 0
}

func (t Tile) WithDir(dir Direction, border BorderStyle) Tile {
	shift := 0
	switch dir {
	case DirDown: shift = 0
	case DirRight: shift = 8
	case DirUp: shift = 16
	case DirLeft: shift = 24
	}
	return tileExchange(t, shift, border)
}

func (t Tile) Rune() rune {
	if r, ok := tileToRune[t]; ok {
		return r
	}
	return ' '
}

