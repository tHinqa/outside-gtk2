// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Unichar T.Gunichar

var (
	UnicodeCanonicalOrdering func(str *Unichar, leng T.Gsize)

	UnicodeCanonicalDecomposition func(u Unichar, resultLen *T.Gsize) *Unichar

	UnicharBreakType      func(u Unichar) UnicodeBreakType
	UnicharCombiningClass func(u Unichar) int
	UnicharDigitValue     func(u Unichar) int
	UnicharGetMirrorChar  func(u Unichar, mirroredCh *Unichar) bool
	UnicharGetScript      func(u Unichar) UnicodeScript
	UnicharIsalnum        func(u Unichar) bool
	UnicharIsalpha        func(u Unichar) bool
	UnicharIscntrl        func(u Unichar) bool
	UnicharIsdefined      func(u Unichar) bool
	UnicharIsdigit        func(u Unichar) bool
	UnicharIsgraph        func(u Unichar) bool
	UnicharIslower        func(u Unichar) bool
	UnicharIsmark         func(u Unichar) bool
	UnicharIsprint        func(u Unichar) bool
	UnicharIspunct        func(u Unichar) bool
	UnicharIsspace        func(u Unichar) bool
	UnicharIstitle        func(u Unichar) bool
	UnicharIsupper        func(u Unichar) bool
	UnicharIswide         func(u Unichar) bool
	UnicharIswideCjk      func(u Unichar) bool
	UnicharIsxdigit       func(u Unichar) bool
	UnicharIszerowidth    func(u Unichar) bool
	UnicharTolower        func(u Unichar) Unichar
	UnicharTotitle        func(u Unichar) Unichar
	UnicharToupper        func(u Unichar) Unichar
	UnicharToUtf8         func(u Unichar, outbuf string) int
	UnicharType           func(u Unichar) UnicodeType
	UnicharValidate       func(u Unichar) bool
	UnicharXdigitValue    func(u Unichar) int
)

func (u Unichar) BreakType() UnicodeBreakType { return UnicharBreakType(u) }
func (u Unichar) CanonicalDecomposition(resultLen *T.Gsize) *Unichar {
	return UnicodeCanonicalDecomposition(u, resultLen)
}
func (u Unichar) CombiningClass() int                    { return UnicharCombiningClass(u) }
func (u Unichar) DigitValue() int                        { return UnicharDigitValue(u) }
func (u Unichar) GetMirrorChar(mirroredCh *Unichar) bool { return UnicharGetMirrorChar(u, mirroredCh) }
func (u Unichar) GetScript() UnicodeScript               { return UnicharGetScript(u) }
func (u Unichar) Isalnum() bool                          { return UnicharIsalnum(u) }
func (u Unichar) Isalpha() bool                          { return UnicharIsalpha(u) }
func (u Unichar) Iscntrl() bool                          { return UnicharIscntrl(u) }
func (u Unichar) Isdefined() bool                        { return UnicharIsdefined(u) }
func (u Unichar) Isdigit() bool                          { return UnicharIsdigit(u) }
func (u Unichar) Isgraph() bool                          { return UnicharIsgraph(u) }
func (u Unichar) Islower() bool                          { return UnicharIslower(u) }
func (u Unichar) Ismark() bool                           { return UnicharIsmark(u) }
func (u Unichar) Isprint() bool                          { return UnicharIsprint(u) }
func (u Unichar) Ispunct() bool                          { return UnicharIspunct(u) }
func (u Unichar) Isspace() bool                          { return UnicharIsspace(u) }
func (u Unichar) Istitle() bool                          { return UnicharIstitle(u) }
func (u Unichar) Isupper() bool                          { return UnicharIsupper(u) }
func (u Unichar) Iswide() bool                           { return UnicharIswide(u) }
func (u Unichar) IswideCjk() bool                        { return UnicharIswideCjk(u) }
func (u Unichar) Isxdigit() bool                         { return UnicharIsxdigit(u) }
func (u Unichar) Iszerowidth() bool                      { return UnicharIszerowidth(u) }
func (u Unichar) Tolower() Unichar                       { return UnicharTolower(u) }
func (u Unichar) Totitle() Unichar                       { return UnicharTotitle(u) }
func (u Unichar) Toupper() Unichar                       { return UnicharToupper(u) }
func (u Unichar) ToUtf8(outbuf string) int               { return UnicharToUtf8(u, outbuf) }
func (u Unichar) Type() UnicodeType                      { return UnicharType(u) }
func (u Unichar) Validate() bool                         { return UnicharValidate(u) }
func (u Unichar) XdigitValue() int                       { return UnicharXdigitValue(u) }

type UnicodeBreakType Enum

const (
	UNICODE_BREAK_MANDATORY UnicodeBreakType = iota
	UNICODE_BREAK_CARRIAGE_RETURN
	UNICODE_BREAK_LINE_FEED
	UNICODE_BREAK_COMBINING_MARK
	UNICODE_BREAK_SURROGATE
	UNICODE_BREAK_ZERO_WIDTH_SPACE
	UNICODE_BREAK_INSEPARABLE
	UNICODE_BREAK_NON_BREAKING_GLUE
	UNICODE_BREAK_CONTINGENT
	UNICODE_BREAK_SPACE
	UNICODE_BREAK_AFTER
	UNICODE_BREAK_BEFORE
	UNICODE_BREAK_BEFORE_AND_AFTER
	UNICODE_BREAK_HYPHEN
	UNICODE_BREAK_NON_STARTER
	UNICODE_BREAK_OPEN_PUNCTUATION
	UNICODE_BREAK_CLOSE_PUNCTUATION
	UNICODE_BREAK_QUOTATION
	UNICODE_BREAK_EXCLAMATION
	UNICODE_BREAK_IDEOGRAPHIC
	UNICODE_BREAK_NUMERIC
	UNICODE_BREAK_INFIX_SEPARATOR
	UNICODE_BREAK_SYMBOL
	UNICODE_BREAK_ALPHABETIC
	UNICODE_BREAK_PREFIX
	UNICODE_BREAK_POSTFIX
	UNICODE_BREAK_COMPLEX_CONTEXT
	UNICODE_BREAK_AMBIGUOUS
	UNICODE_BREAK_UNKNOWN
	UNICODE_BREAK_NEXT_LINE
	UNICODE_BREAK_WORD_JOINER
	UNICODE_BREAK_HANGUL_L_JAMO
	UNICODE_BREAK_HANGUL_V_JAMO
	UNICODE_BREAK_HANGUL_T_JAMO
	UNICODE_BREAK_HANGUL_LV_SYLLABLE
	UNICODE_BREAK_HANGUL_LVT_SYLLABLE
	UNICODE_BREAK_CLOSE_PARANTHESIS
)

type UnicodeScript Enum

const (
	UNICODE_SCRIPT_INVALID_CODE UnicodeScript = iota - 1
	UNICODE_SCRIPT_COMMON
	UNICODE_SCRIPT_INHERITED
	UNICODE_SCRIPT_ARABIC
	UNICODE_SCRIPT_ARMENIAN
	UNICODE_SCRIPT_BENGALI
	UNICODE_SCRIPT_BOPOMOFO
	UNICODE_SCRIPT_CHEROKEE
	UNICODE_SCRIPT_COPTIC
	UNICODE_SCRIPT_CYRILLIC
	UNICODE_SCRIPT_DESERET
	UNICODE_SCRIPT_DEVANAGARI
	UNICODE_SCRIPT_ETHIOPIC
	UNICODE_SCRIPT_GEORGIAN
	UNICODE_SCRIPT_GOTHIC
	UNICODE_SCRIPT_GREEK
	UNICODE_SCRIPT_GUJARATI
	UNICODE_SCRIPT_GURMUKHI
	UNICODE_SCRIPT_HAN
	UNICODE_SCRIPT_HANGUL
	UNICODE_SCRIPT_HEBREW
	UNICODE_SCRIPT_HIRAGANA
	UNICODE_SCRIPT_KANNADA
	UNICODE_SCRIPT_KATAKANA
	UNICODE_SCRIPT_KHMER
	UNICODE_SCRIPT_LAO
	UNICODE_SCRIPT_LATIN
	UNICODE_SCRIPT_MALAYALAM
	UNICODE_SCRIPT_MONGOLIAN
	UNICODE_SCRIPT_MYANMAR
	UNICODE_SCRIPT_OGHAM
	UNICODE_SCRIPT_OLD_ITALIC
	UNICODE_SCRIPT_ORIYA
	UNICODE_SCRIPT_RUNIC
	UNICODE_SCRIPT_SINHALA
	UNICODE_SCRIPT_SYRIAC
	UNICODE_SCRIPT_TAMIL
	UNICODE_SCRIPT_TELUGU
	UNICODE_SCRIPT_THAANA
	UNICODE_SCRIPT_THAI
	UNICODE_SCRIPT_TIBETAN
	UNICODE_SCRIPT_CANADIAN_ABORIGINAL
	UNICODE_SCRIPT_YI
	UNICODE_SCRIPT_TAGALOG
	UNICODE_SCRIPT_HANUNOO
	UNICODE_SCRIPT_BUHID
	UNICODE_SCRIPT_TAGBANWA
	UNICODE_SCRIPT_BRAILLE
	UNICODE_SCRIPT_CYPRIOT
	UNICODE_SCRIPT_LIMBU
	UNICODE_SCRIPT_OSMANYA
	UNICODE_SCRIPT_SHAVIAN
	UNICODE_SCRIPT_LINEAR_B
	UNICODE_SCRIPT_TAI_LE
	UNICODE_SCRIPT_UGARITIC
	UNICODE_SCRIPT_NEW_TAI_LUE
	UNICODE_SCRIPT_BUGINESE
	UNICODE_SCRIPT_GLAGOLITIC
	UNICODE_SCRIPT_TIFINAGH
	UNICODE_SCRIPT_SYLOTI_NAGRI
	UNICODE_SCRIPT_OLD_PERSIAN
	UNICODE_SCRIPT_KHAROSHTHI
	UNICODE_SCRIPT_UNKNOWN
	UNICODE_SCRIPT_BALINESE
	UNICODE_SCRIPT_CUNEIFORM
	UNICODE_SCRIPT_PHOENICIAN
	UNICODE_SCRIPT_PHAGS_PA
	UNICODE_SCRIPT_NKO
	UNICODE_SCRIPT_KAYAH_LI
	UNICODE_SCRIPT_LEPCHA
	UNICODE_SCRIPT_REJANG
	UNICODE_SCRIPT_SUNDANESE
	UNICODE_SCRIPT_SAURASHTRA
	UNICODE_SCRIPT_CHAM
	UNICODE_SCRIPT_OL_CHIKI
	UNICODE_SCRIPT_VAI
	UNICODE_SCRIPT_CARIAN
	UNICODE_SCRIPT_LYCIAN
	UNICODE_SCRIPT_LYDIAN
	UNICODE_SCRIPT_AVESTAN
	UNICODE_SCRIPT_BAMUM
	UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS
	UNICODE_SCRIPT_IMPERIAL_ARAMAIC
	UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI
	UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN
	UNICODE_SCRIPT_JAVANESE
	UNICODE_SCRIPT_KAITHI
	UNICODE_SCRIPT_LISU
	UNICODE_SCRIPT_MEETEI_MAYEK
	UNICODE_SCRIPT_OLD_SOUTH_ARABIAN
	UNICODE_SCRIPT_OLD_TURKIC
	UNICODE_SCRIPT_SAMARITAN
	UNICODE_SCRIPT_TAI_THAM
	UNICODE_SCRIPT_TAI_VIET
	UNICODE_SCRIPT_BATAK
	UNICODE_SCRIPT_BRAHMI
	UNICODE_SCRIPT_MANDAIC
)

type UnicodeType Enum

const (
	UNICODE_CONTROL UnicodeType = iota
	UNICODE_FORMAT
	UNICODE_UNASSIGNED
	UNICODE_PRIVATE_USE
	UNICODE_SURROGATE
	UNICODE_LOWERCASE_LETTER
	UNICODE_MODIFIER_LETTER
	UNICODE_OTHER_LETTER
	UNICODE_TITLECASE_LETTER
	UNICODE_UPPERCASE_LETTER
	UNICODE_COMBINING_MARK
	UNICODE_ENCLOSING_MARK
	UNICODE_NON_SPACING_MARK
	UNICODE_DECIMAL_NUMBER
	UNICODE_LETTER_NUMBER
	UNICODE_OTHER_NUMBER
	UNICODE_CONNECT_PUNCTUATION
	UNICODE_DASH_PUNCTUATION
	UNICODE_CLOSE_PUNCTUATION
	UNICODE_FINAL_PUNCTUATION
	UNICODE_INITIAL_PUNCTUATION
	UNICODE_OTHER_PUNCTUATION
	UNICODE_OPEN_PUNCTUATION
	UNICODE_CURRENCY_SYMBOL
	UNICODE_MODIFIER_SYMBOL
	UNICODE_MATH_SYMBOL
	UNICODE_OTHER_SYMBOL
	UNICODE_LINE_SEPARATOR
	UNICODE_PARAGRAPH_SEPARATOR
	UNICODE_SPACE_SEPARATOR
)
