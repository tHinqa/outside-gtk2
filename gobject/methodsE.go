// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

var (
	EnumCompleteTypeInfo func(gEnumType Type, info *TypeInfo, constValues *EnumValue)
	EnumRegisterStatic   func(name string, constStaticValues *EnumValue) Type

	EnumGetValue       func(e *EnumClass, value int) *EnumValue
	EnumGetValueByName func(e *EnumClass, name string) *EnumValue
	EnumGetValueByNick func(e *EnumClass, nick string) *EnumValue
)

type EnumClass struct {
	TypeClass TypeClass
	Minimum   int
	Maximum   int
	NValues   uint
	Values    *EnumValue
}

type EnumValue struct {
	Value     int
	ValueName *T.Gchar
	ValueNick *T.Gchar
}
