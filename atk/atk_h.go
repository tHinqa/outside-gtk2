// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
)

type Hyperlink struct{ parent O.Object }

var (
	HyperlinkGetType func() O.Type

	HyperlinkGetEndIndex    func(h *Hyperlink) int
	HyperlinkGetNAnchors    func(h *Hyperlink) int
	HyperlinkGetObject      func(h *Hyperlink, i int) *Object
	HyperlinkGetStartIndex  func(h *Hyperlink) int
	HyperlinkGetUri         func(h *Hyperlink, i int) string
	HyperlinkIsInline       func(h *Hyperlink) bool
	HyperlinkIsSelectedLink func(h *Hyperlink) bool
	HyperlinkIsValid        func(h *Hyperlink) bool
)

func (h *Hyperlink) GetEndIndex() int        { return HyperlinkGetEndIndex(h) }
func (h *Hyperlink) GetNAnchors() int        { return HyperlinkGetNAnchors(h) }
func (h *Hyperlink) GetObject(i int) *Object { return HyperlinkGetObject(h, i) }
func (h *Hyperlink) GetStartIndex() int      { return HyperlinkGetStartIndex(h) }
func (h *Hyperlink) GetUri(i int) string     { return HyperlinkGetUri(h, i) }
func (h *Hyperlink) IsInline() bool          { return HyperlinkIsInline(h) }
func (h *Hyperlink) IsSelectedLink() bool    { return HyperlinkIsSelectedLink(h) }
func (h *Hyperlink) IsValid() bool           { return HyperlinkIsValid(h) }

type HyperlinkImpl struct{}

var (
	HyperlinkImplGetType func() O.Type

	HyperlinkImplGetHyperlink func(h *HyperlinkImpl) *Hyperlink
)

func (h *HyperlinkImpl) GetHyperlink() *Hyperlink { return HyperlinkImplGetHyperlink(h) }

type Hypertext struct{}

var (
	HypertextGetType func() O.Type

	HypertextGetLink      func(h *Hypertext, linkIndex int) *Hyperlink
	HypertextGetLinkIndex func(h *Hypertext, charIndex int) int
	HypertextGetNLinks    func(h *Hypertext) int
)

func (h *Hypertext) GetLink(linkIndex int) *Hyperlink { return HypertextGetLink(h, linkIndex) }
func (h *Hypertext) GetLinkIndex(charIndex int) int   { return HypertextGetLinkIndex(h, charIndex) }
func (h *Hypertext) GetNLinks() int                   { return HypertextGetNLinks(h) }

//NOTE(t): Unreferenced
type HyperlinkStateFlags Enum

const HYPERLINK_IS_INLINE HyperlinkStateFlags = 1 << 0

var HyperlinkStateFlagsGetType func() O.Type
