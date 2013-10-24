// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gdk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Font struct {
	Type    FontType
	Ascent  int
	Descent int
}

var (
	FontGetType func() T.GType

	FontFromDescription           func(fontDesc *T.PangoFontDescription) *Font
	FontFromDescriptionForDisplay func(display *Display, fontDesc *T.PangoFontDescription) *Font
	FontLoad                      func(fontName string) *Font
	FontLoadForDisplay            func(display *Display, fontName string) *Font

	FontsetLoad           func(fontsetName string) *Font
	FontsetLoadForDisplay func(display *Display, fontsetName string) *Font

	fontRef   func(f *Font) *Font
	fontUnref func(f *Font)
	fontId    func(f *Font) int
	fontEqual func(f *Font, fontb *Font) T.Gboolean
)

func (f *Font) Ref() *Font                { return fontRef(f) }
func (f *Font) Unref()                    { fontUnref(f) }
func (f *Font) Id() int                   { return fontId(f) }
func (f *Font) Equal(f2 *Font) T.Gboolean { return fontEqual(f, f2) }

type FontType Enum

const (
	FONT_FONT FontType = iota
	FONT_FONTSET
)
