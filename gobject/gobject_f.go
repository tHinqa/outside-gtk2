// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

var (
	FlagsCompleteTypeInfo func(gFlagsType Type, info *TypeInfo, constValues *FlagsValue)
	FlagsRegisterStatic   func(name string, constStaticValues *FlagsValue) Type
)

type FlagsClass struct {
	TypeClass TypeClass
	Mask      uint
	NValues   uint
	Values    *FlagsValue
}

var (
	flagsGetFirstValue  func(f *FlagsClass, value uint) *FlagsValue
	flagsGetValueByName func(f *FlagsClass, name string) *FlagsValue
	flagsGetValueByNick func(f *FlagsClass, nick string) *FlagsValue
)

func (f *FlagsClass) GetFirstValue(value uint) *FlagsValue { return flagsGetFirstValue(f, value) }
func (f *FlagsClass) GetValueByName(name string) *FlagsValue {
	return flagsGetValueByName(f, name)
}
func (f *FlagsClass) GetValueByNick(nick string) *FlagsValue {
	return flagsGetValueByNick(f, nick)
}

type FlagsValue struct {
	Value     uint
	ValueName *T.Gchar
	ValueNick *T.Gchar
}
