// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
)

var (
	GetRoot func() *Object

	GetFocusObject func() *Object

	GetToolkitName func() string

	GetToolkitVersion func() string

	GobjectAccessibleForObject func(obj *O.Object) *Object
)

type GObjectAccessible struct {
	Parent Object
}

var GobjectAccessibleGetObject func(obj *GObjectAccessible) *O.Object

func (o *GObjectAccessible) GetObject() *O.Object { return GobjectAccessibleGetObject(o) }

var (
	GetDefaultRegistry func() *Registry

	GobjectAccessibleGetType func() O.Type
)
