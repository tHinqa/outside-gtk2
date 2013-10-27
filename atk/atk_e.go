// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type EditableText struct{}

var (
	EditableTextGetType func() O.Type

	editableTextCopyText         func(e *EditableText, startPos, endPos int)
	editableTextCutText          func(e *EditableText, startPos, endPos int)
	editableTextDeleteText       func(e *EditableText, startPos, endPos int)
	editableTextInsertText       func(e *EditableText, str string, length int, position *int)
	editableTextPasteText        func(e *EditableText, position int)
	editableTextSetRunAttributes func(e *EditableText, attribSet *AttributeSet, startOffset, endOffset int) T.Gboolean
	editableTextSetTextContents  func(e *EditableText, str string)
)

func (e *EditableText) CopyText(startPos, endPos int)   { editableTextCopyText(e, startPos, endPos) }
func (e *EditableText) CutText(startPos, endPos int)    { editableTextCutText(e, startPos, endPos) }
func (e *EditableText) DeleteText(startPos, endPos int) { editableTextDeleteText(e, startPos, endPos) }
func (e *EditableText) InsertText(str string, length int, position *int) {
	editableTextInsertText(e, str, length, position)
}
func (e *EditableText) PasteText(position int) { editableTextPasteText(e, position) }
func (e *EditableText) SetRunAttributes(attribSet *AttributeSet, startOffset, endOffset int) T.Gboolean {
	return editableTextSetRunAttributes(e, attribSet, startOffset, endOffset)
}
func (e *EditableText) SetTextContents(str string) { editableTextSetTextContents(e, str) }
