// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type TabAlign Enum

const TAB_LEFT TabAlign = 0

var TabAlignGetType func() O.Type

type TabArray struct{}

var (
	TabArrayGetType          func() O.Type
	TabArrayNew              func(initialSize int, positionsInPixels bool) *TabArray
	TabArrayNewWithPositions func(size int, positionsInPixels bool, firstAlignment TabAlign, firstPosition int, v ...VArg) *TabArray

	TabArrayCopy                 func(t *TabArray) *TabArray
	TabArrayFree                 func(t *TabArray)
	TabArrayGetPositionsInPixels func(t *TabArray) bool
	TabArrayGetSize              func(t *TabArray) int
	TabArrayGetTab               func(t *TabArray, tabIndex int, Alignment *TabAlign, location *int)
	TabArrayGetTabs              func(t *TabArray, Alignments **TabAlign, locations **int)
	TabArrayResize               func(t *TabArray, newSize int)
	TabArraySetTab               func(t *TabArray, tabIndex int, Alignment TabAlign, location int)
)

func (t *TabArray) Copy() *TabArray            { return TabArrayCopy(t) }
func (t *TabArray) Free()                      { TabArrayFree(t) }
func (t *TabArray) GetPositionsInPixels() bool { return TabArrayGetPositionsInPixels(t) }
func (t *TabArray) GetSize() int               { return TabArrayGetSize(t) }
func (t *TabArray) GetTab(tabIndex int, Alignment *TabAlign, location *int) {
	TabArrayGetTab(t, tabIndex, Alignment, location)
}
func (t *TabArray) GetTabs(Alignments **TabAlign, locations **int) {
	TabArrayGetTabs(t, Alignments, locations)
}
func (t *TabArray) Resize(newSize int) { TabArrayResize(t, newSize) }
func (t *TabArray) SetTab(tabIndex int, Alignment TabAlign, location int) {
	TabArraySetTab(t, tabIndex, Alignment, location)
}

var TrimString func(str string) string
