// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	// T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Font struct {
	Type    FontType
	Ascent  int
	Descent int
}

var (
	FontGetType func() O.Type

	FontFromDescription           func(fontDesc *P.FontDescription) *Font
	FontFromDescriptionForDisplay func(display *Display, fontDesc *P.FontDescription) *Font
	FontLoad                      func(fontName string) *Font
	FontLoadForDisplay            func(display *Display, fontName string) *Font

	FontsetLoad           func(fontsetName string) *Font
	FontsetLoadForDisplay func(display *Display, fontsetName string) *Font

	FontRef   func(f *Font) *Font
	FontUnref func(f *Font)
	FontId    func(f *Font) int
	FontEqual func(f *Font, fontb *Font) bool
)

func (f *Font) Ref() *Font          { return FontRef(f) }
func (f *Font) Unref()              { FontUnref(f) }
func (f *Font) Id() int             { return FontId(f) }
func (f *Font) Equal(f2 *Font) bool { return FontEqual(f, f2) }

type FontType Enum

const (
	FONT_FONT FontType = iota
	FONT_FONTSET
)
