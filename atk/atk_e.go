// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
)

type EditableText struct{}

var (
	EditableTextGetType func() O.Type

	EditableTextCopyText         func(e *EditableText, startPos, endPos int)
	EditableTextCutText          func(e *EditableText, startPos, endPos int)
	EditableTextDeleteText       func(e *EditableText, startPos, endPos int)
	EditableTextInsertText       func(e *EditableText, str string, length int, position *int)
	EditableTextPasteText        func(e *EditableText, position int)
	EditableTextSetRunAttributes func(e *EditableText, attribSet *AttributeSet, startOffset, endOffset int) bool
	EditableTextSetTextContents  func(e *EditableText, str string)
)

func (e *EditableText) CopyText(startPos, endPos int)   { EditableTextCopyText(e, startPos, endPos) }
func (e *EditableText) CutText(startPos, endPos int)    { EditableTextCutText(e, startPos, endPos) }
func (e *EditableText) DeleteText(startPos, endPos int) { EditableTextDeleteText(e, startPos, endPos) }
func (e *EditableText) InsertText(str string, length int, position *int) {
	EditableTextInsertText(e, str, length, position)
}
func (e *EditableText) PasteText(position int) { EditableTextPasteText(e, position) }
func (e *EditableText) SetRunAttributes(attribSet *AttributeSet, startOffset, endOffset int) bool {
	return EditableTextSetRunAttributes(e, attribSet, startOffset, endOffset)
}
func (e *EditableText) SetTextContents(str string) { EditableTextSetTextContents(e, str) }
