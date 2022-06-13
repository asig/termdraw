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

var (
	// Color values generated with https://github.com/canidlogic/vgapal
	ColBlack        = termbox.RGBToAttribute(0, 0, 0)
	ColBlue         = termbox.RGBToAttribute(0, 0, 170)
	ColGreen        = termbox.RGBToAttribute(0, 170, 0)
	ColCyan         = termbox.RGBToAttribute(0, 170, 170)
	ColRed          = termbox.RGBToAttribute(170, 0, 0)
	ColMagenta      = termbox.RGBToAttribute(170, 0, 170)
	ColBrown        = termbox.RGBToAttribute(170, 85, 0)
	ColLightGrey    = termbox.RGBToAttribute(170, 170, 170)
	ColGrey         = termbox.RGBToAttribute(85, 85, 85)
	ColLightBlue    = termbox.RGBToAttribute(85, 85, 255)
	ColLightGreen   = termbox.RGBToAttribute(85, 255, 85)
	ColLightCyan    = termbox.RGBToAttribute(85, 255, 255)
	ColLightRed     = termbox.RGBToAttribute(255, 85, 85)
	ColLightMagenta = termbox.RGBToAttribute(255, 85, 255)
	ColYellow       = termbox.RGBToAttribute(255, 255, 85)
	ColWhite        = termbox.RGBToAttribute(255, 255, 255)
)

