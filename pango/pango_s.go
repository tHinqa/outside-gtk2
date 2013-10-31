// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

var (
	ScanInt    func(pos **T.Char, out *int) bool
	ScanString func(pos **T.Char, out *L.String) bool
	ScanWord   func(pos **T.Char, out *L.String) bool
)

type Script Enum

const (
	SCRIPT_INVALID_CODE Script = iota - 1
	SCRIPT_COMMON
	SCRIPT_INHERITED
	SCRIPT_ARABIC
	SCRIPT_ARMENIAN
	SCRIPT_BENGALI
	SCRIPT_BOPOMOFO
	SCRIPT_CHEROKEE
	SCRIPT_COPTIC
	SCRIPT_CYRILLIC
	SCRIPT_DESERET
	SCRIPT_DEVANAGARI
	SCRIPT_ETHIOPIC
	SCRIPT_GEORGIAN
	SCRIPT_GOTHIC
	SCRIPT_GREEK
	SCRIPT_GUJARATI
	SCRIPT_GURMUKHI
	SCRIPT_HAN
	SCRIPT_HANGUL
	SCRIPT_HEBREW
	SCRIPT_HIRAGANA
	SCRIPT_KANNADA
	SCRIPT_KATAKANA
	SCRIPT_KHMER
	SCRIPT_LAO
	SCRIPT_LATIN
	SCRIPT_MALAYALAM
	SCRIPT_MONGOLIAN
	SCRIPT_MYANMAR
	SCRIPT_OGHAM
	SCRIPT_OLD_ITALIC
	SCRIPT_ORIYA
	SCRIPT_RUNIC
	SCRIPT_SINHALA
	SCRIPT_SYRIAC
	SCRIPT_TAMIL
	SCRIPT_TELUGU
	SCRIPT_THAANA
	SCRIPT_THAI
	SCRIPT_TIBETAN
	SCRIPT_CANADIAN_ABORIGINAL
	SCRIPT_YI
	SCRIPT_TAGALOG
	SCRIPT_HANUNOO
	SCRIPT_BUHID
	SCRIPT_TAGBANWA
	SCRIPT_BRAILLE
	SCRIPT_CYPRIOT
	SCRIPT_LIMBU
	SCRIPT_OSMANYA
	SCRIPT_SHAVIAN
	SCRIPT_LINEAR_B
	SCRIPT_TAI_LE
	SCRIPT_UGARITIC
	SCRIPT_NEW_TAI_LUE
	SCRIPT_BUGINESE
	SCRIPT_GLAGOLITIC
	SCRIPT_TIFINAGH
	SCRIPT_SYLOTI_NAGRI
	SCRIPT_OLD_PERSIAN
	SCRIPT_KHAROSHTHI
	SCRIPT_UNKNOWN
	SCRIPT_BALINESE
	SCRIPT_CUNEIFORM
	SCRIPT_PHOENICIAN
	SCRIPT_PHAGS_PA
	SCRIPT_NKO
	SCRIPT_KAYAH_LI
	SCRIPT_LEPCHA
	SCRIPT_REJANG
	SCRIPT_SUNDANESE
	SCRIPT_SAURASHTRA
	SCRIPT_CHAM
	SCRIPT_OL_CHIKI
	SCRIPT_VAI
	SCRIPT_CARIAN
	SCRIPT_LYCIAN
	SCRIPT_LYDIAN
)

var (
	ScriptGetType func() O.Type

	ScriptForUnichar func(ch T.Gunichar) Script

	ScriptGetSampleLanguage func(s Script) *Language
)

func (s Script) GetSampleLanguage() *Language { return ScriptGetSampleLanguage(s) }

type ScriptIter struct{}

var (
	ScriptIterNew func(text string, length int) *ScriptIter

	ScriptIterFree     func(s *ScriptIter)
	ScriptIterGetRange func(s *ScriptIter, start, end **T.Char, script *Script)
	ScriptIterNext     func(s *ScriptIter) bool
)

func (s *ScriptIter) Free() { ScriptIterFree(s) }
func (s *ScriptIter) GetRange(start, end **T.Char, script *Script) {
	ScriptIterGetRange(s, start, end, script)
}
func (s *ScriptIter) Next() bool { return ScriptIterNext(s) }

var (
	Shape         func(text string, length int, Analysis *Analysis, glyphs *GlyphString)
	SkipSpace     func(pos **T.Char) bool
	SplitFileList func(str string) []string
)

type Stretch Enum

const (
	STRETCH_ULTRA_CONDENSED Stretch = iota
	STRETCH_EXTRA_CONDENSED
	STRETCH_CONDENSED
	STRETCH_SEMI_CONDENSED
	STRETCH_NORMAL
	STRETCH_SEMI_EXPANDED
	STRETCH_EXPANDED
	STRETCH_EXTRA_EXPANDED
	STRETCH_ULTRA_EXPANDED
)

var StretchGetType func() O.Type

type Style Enum

const (
	STYLE_NORMAL Style = iota
	STYLE_OBLIQUE
	STYLE_ITALIC
)

var StyleGetType func() O.Type
