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

	hyperlinkGetEndIndex    func(h *Hyperlink) int
	hyperlinkGetNAnchors    func(h *Hyperlink) int
	hyperlinkGetObject      func(h *Hyperlink, i int) *Object
	hyperlinkGetStartIndex  func(h *Hyperlink) int
	hyperlinkGetUri         func(h *Hyperlink, i int) string
	hyperlinkIsInline       func(h *Hyperlink) bool
	hyperlinkIsSelectedLink func(h *Hyperlink) bool
	hyperlinkIsValid        func(h *Hyperlink) bool
)

func (h *Hyperlink) GetEndIndex() int        { return hyperlinkGetEndIndex(h) }
func (h *Hyperlink) GetNAnchors() int        { return hyperlinkGetNAnchors(h) }
func (h *Hyperlink) GetObject(i int) *Object { return hyperlinkGetObject(h, i) }
func (h *Hyperlink) GetStartIndex() int      { return hyperlinkGetStartIndex(h) }
func (h *Hyperlink) GetUri(i int) string     { return hyperlinkGetUri(h, i) }
func (h *Hyperlink) IsInline() bool          { return hyperlinkIsInline(h) }
func (h *Hyperlink) IsSelectedLink() bool    { return hyperlinkIsSelectedLink(h) }
func (h *Hyperlink) IsValid() bool           { return hyperlinkIsValid(h) }

type HyperlinkImpl struct{}

var (
	HyperlinkImplGetType func() O.Type

	hyperlinkImplGetHyperlink func(h *HyperlinkImpl) *Hyperlink
)

func (h *HyperlinkImpl) GetHyperlink() *Hyperlink { return hyperlinkImplGetHyperlink(h) }

type Hypertext struct{}

var (
	HypertextGetType func() O.Type

	hypertextGetLink      func(h *Hypertext, linkIndex int) *Hyperlink
	hypertextGetLinkIndex func(h *Hypertext, charIndex int) int
	hypertextGetNLinks    func(h *Hypertext) int
)

func (h *Hypertext) GetLink(linkIndex int) *Hyperlink { return hypertextGetLink(h, linkIndex) }
func (h *Hypertext) GetLinkIndex(charIndex int) int   { return hypertextGetLinkIndex(h, charIndex) }
func (h *Hypertext) GetNLinks() int                   { return hypertextGetNLinks(h) }

//NOTE(t): Unreferenced
type HyperlinkStateFlags Enum

const HYPERLINK_IS_INLINE HyperlinkStateFlags = 1 << 0

var HyperlinkStateFlagsGetType func() O.Type
