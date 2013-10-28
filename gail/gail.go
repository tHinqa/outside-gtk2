// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gail provides API definitions for accessing
//libgailutil-18.dll.
package gail

import (
	"github.com/tHinqa/outside"
	A "github.com/tHinqa/outside-gtk2/atk"
	O "github.com/tHinqa/outside-gtk2/gobject"
	G "github.com/tHinqa/outside-gtk2/gtk"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type Enum int

var (
	MiscAddAttribute func(
		attribSet *A.AttributeSet,
		attr A.TextAttribute,
		value string) *A.AttributeSet

	MiscLayoutGetRunAttributes func(
		attribSet *A.AttributeSet,
		layout *P.Layout,
		text string,
		offset int,
		startOffset, endOffset *int) *A.AttributeSet

	MiscGetDefaultAttributes func(
		attribSet *A.AttributeSet,
		layout *P.Layout,
		widget *G.Widget) *A.AttributeSet

	MiscGetExtentsFromPangoRectangle func(
		widget *G.Widget,
		charRect *P.Rectangle,
		xLayout, yLayout int,
		x, y, width, height *int,
		coords A.CoordType)

	MiscGetIndexAtPointInLayout func(
		widget *G.Widget,
		layout *P.Layout,
		xLayout, yLayout, x, y int,
		coords A.CoordType) int

	MiscGetOrigins func(
		widget *G.Widget,
		xWindow, yWindow, xToplevel, yToplevel *int)

	MiscAddToAttrSet func(
		attribSet *A.AttributeSet,
		attrs *G.TextAttributes,
		attr A.TextAttribute) *A.AttributeSet

	MiscBufferGetRunAttributes func(
		buffer *G.TextBuffer,
		offset int,
		startOffset, endOffset *int) *A.AttributeSet
)

type TextUtil struct {
	Parent O.Object
	Buffer *G.TextBuffer
}

type OffsetType Enum

const (
	BEFORE_OFFSET OffsetType = iota
	AT_OFFSET
	AFTER_OFFSET
)

var (
	TextUtilGetType func() O.Type
	TextUtilNew     func() *TextUtil

	textUtilTextSetup    func(t *TextUtil, text string)
	textUtilBufferSetup  func(t *TextUtil, buffer *G.TextBuffer)
	textUtilGetText      func(t *TextUtil, layout T.Gpointer, function OffsetType, boundaryType A.TextBoundary, offset int, startOffset, endOffset *int) string
	textUtilGetSubstring func(t *TextUtil, startPos, endPos int) string
)

func (t *TextUtil) TextSetup(text string)            { textUtilTextSetup(t, text) }
func (t *TextUtil) BufferSetup(buffer *G.TextBuffer) { textUtilBufferSetup(t, buffer) }
func (t *TextUtil) GetText(layout T.Gpointer, function OffsetType, boundaryType A.TextBoundary, offset int, startOffset, endOffset *int) string {
	return textUtilGetText(t, layout, function, boundaryType, offset, startOffset, endOffset)
}
func (t *TextUtil) GetSubstring(startPos, endPos int) string {
	return textUtilGetSubstring(t, startPos, endPos)
}

var dll = "libgailutil-18.dll"

var apiList = outside.Apis{
	{"gail_misc_add_attribute", &MiscAddAttribute},
	{"gail_misc_add_to_attr_set", &MiscAddToAttrSet},
	{"gail_misc_buffer_get_run_attributes", &MiscBufferGetRunAttributes},
	{"gail_misc_get_default_attributes", &MiscGetDefaultAttributes},
	{"gail_misc_get_extents_from_pango_rectangle", &MiscGetExtentsFromPangoRectangle},
	{"gail_misc_get_index_at_point_in_layout", &MiscGetIndexAtPointInLayout},
	{"gail_misc_get_origins", &MiscGetOrigins},
	{"gail_misc_layout_get_run_attributes", &MiscLayoutGetRunAttributes},
	{"gail_text_util_buffer_setup", &textUtilBufferSetup},
	{"gail_text_util_get_substring", &textUtilGetSubstring},
	{"gail_text_util_get_text", &textUtilGetText},
	{"gail_text_util_get_type", &TextUtilGetType},
	{"gail_text_util_new", &TextUtilNew},
	{"gail_text_util_text_setup", &textUtilTextSetup},
}
