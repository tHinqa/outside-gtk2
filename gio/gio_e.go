// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Emblem struct{}

var (
	EmblemGetType       func() O.Type
	EmblemNew           func(icon *Icon) *Emblem
	EmblemNewWithOrigin func(icon *Icon, origin EmblemOrigin) *Emblem

	emblemGetIcon   func(e *Emblem) *Icon
	emblemGetOrigin func(e *Emblem) EmblemOrigin
)

func (e *Emblem) GetIcon() *Icon          { return emblemGetIcon(e) }
func (e *Emblem) GetOrigin() EmblemOrigin { return emblemGetOrigin(e) }

type EmblemedIcon struct {
	Parent O.Object
	_      *struct{}
}

var (
	EmblemedIconGetType func() O.Type
	EmblemedIconNew     func(icon *Icon, emblem *Emblem) *Icon

	emblemedIconGetIcon      func(e *EmblemedIcon) *Icon
	emblemedIconGetEmblems   func(e *EmblemedIcon) *T.GList
	emblemedIconAddEmblem    func(e *EmblemedIcon, emblem *Emblem)
	emblemedIconClearEmblems func(e *EmblemedIcon)
)

func (e *EmblemedIcon) GetIcon() *Icon           { return emblemedIconGetIcon(e) }
func (e *EmblemedIcon) GetEmblems() *T.GList     { return emblemedIconGetEmblems(e) }
func (e *EmblemedIcon) AddEmblem(emblem *Emblem) { emblemedIconAddEmblem(e, emblem) }
func (e *EmblemedIcon) ClearEmblems()            { emblemedIconClearEmblems(e) }

type EmblemOrigin Enum

const (
	EMBLEM_ORIGIN_UNKNOWN EmblemOrigin = iota
	EMBLEM_ORIGIN_DEVICE
	EMBLEM_ORIGIN_LIVEMETADATA
	EMBLEM_ORIGIN_TAG
)

var EmblemOriginGetType func() O.Type
