// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"

// . "github.com/tHinqa/outside/types"
)

type Binding struct{}

var (
	BindingGetType func() Type

	BindingGetFlags          func(b *Binding) BindingFlags
	BindingGetSource         func(b *Binding) *Object
	BindingGetTarget         func(b *Binding) *Object
	BindingGetSourceProperty func(b *Binding) string
	BindingGetTargetProperty func(b *Binding) string
)

type BindingFlags Enum

const (
	BINDING_BIDIRECTIONAL BindingFlags = 1 << iota
	BINDING_SYNC_CREATE
	BINDING_INVERT_BOOLEAN
	BINDING_DEFAULT BindingFlags = 0
)

var BindingFlagsGetType func() Type

type BindingTransformFunc func(
	binding *Binding,
	sourceValue, targetValue *T.GValue,
	userData T.Gpointer) T.Gboolean
