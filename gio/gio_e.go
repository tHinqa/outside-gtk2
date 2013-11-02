// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
)

type Emblem struct{}

var (
	EmblemGetType       func() O.Type
	EmblemNew           func(icon *Icon) *Emblem
	EmblemNewWithOrigin func(icon *Icon, origin EmblemOrigin) *Emblem

	EmblemGetIcon   func(e *Emblem) *Icon
	EmblemGetOrigin func(e *Emblem) EmblemOrigin
)

func (e *Emblem) GetIcon() *Icon          { return EmblemGetIcon(e) }
func (e *Emblem) GetOrigin() EmblemOrigin { return EmblemGetOrigin(e) }

type EmblemedIcon struct {
	Parent O.Object
	_      *struct{}
}

var (
	EmblemedIconGetType func() O.Type
	EmblemedIconNew     func(icon *Icon, emblem *Emblem) *Icon

	EmblemedIconGetIcon      func(e *EmblemedIcon) *Icon
	EmblemedIconGetEmblems   func(e *EmblemedIcon) *L.List
	EmblemedIconAddEmblem    func(e *EmblemedIcon, emblem *Emblem)
	EmblemedIconClearEmblems func(e *EmblemedIcon)
)

func (e *EmblemedIcon) GetIcon() *Icon           { return EmblemedIconGetIcon(e) }
func (e *EmblemedIcon) GetEmblems() *L.List      { return EmblemedIconGetEmblems(e) }
func (e *EmblemedIcon) AddEmblem(emblem *Emblem) { EmblemedIconAddEmblem(e, emblem) }
func (e *EmblemedIcon) ClearEmblems()            { EmblemedIconClearEmblems(e) }

type EmblemOrigin Enum

const (
	EMBLEM_ORIGIN_UNKNOWN EmblemOrigin = iota
	EMBLEM_ORIGIN_DEVICE
	EMBLEM_ORIGIN_LIVEMETADATA
	EMBLEM_ORIGIN_TAG
)

var EmblemOriginGetType func() O.Type
