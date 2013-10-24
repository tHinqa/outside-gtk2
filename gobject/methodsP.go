// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type ParamFlags Enum

const (
	PARAM_READABLE ParamFlags = 1 << iota
	PARAM_WRITABLE
	PARAM_CONSTRUCT
	PARAM_CONSTRUCT_ONLY
	PARAM_LAX_VALIDATION
	PARAM_STATIC_NAME
	PARAM_STATIC_NICK
	PARAM_STATIC_BLURB
	PARAM_DEPRECATED ParamFlags = -(1 << 31)
	PARAM_PRIVATE               = PARAM_STATIC_NAME
)

type ParamSpec struct {
	TypeInstance TypeInstance
	Name         *T.Gchar
	Flags        ParamFlags
	Value_type   Type
	Owner_type   Type
	_Nick        *T.Gchar //TODO(t):Use?
	_Blurb       *T.Gchar //TODO(t):Use?
	Qdata        *T.GData
	Ref_count    uint
	Param_id     uint
}

var (
	ParamSpecBoolean    func(name, nick, blurb string, defaultValue T.Gboolean, flags ParamFlags) *ParamSpec
	ParamSpecBoxed      func(name, nick, blurb string, boxedType Type, flags ParamFlags) *ParamSpec
	ParamSpecChar       func(name, nick, blurb string, minimum, maximum, defaultValue int8, flags ParamFlags) *ParamSpec
	ParamSpecDouble     func(name, nick, blurb string, minimum, maximum, defaultValue float64, flags ParamFlags) *ParamSpec
	ParamSpecEnum       func(name, nick, blurb string, enumType Type, defaultValue int, flags ParamFlags) *ParamSpec
	ParamSpecFlags      func(name, nick, blurb string, flagsType Type, defaultValue uint, flags ParamFlags) *ParamSpec
	ParamSpecFloat      func(name, nick, blurb string, minimum, maximum, defaultValue float32, flags ParamFlags) *ParamSpec
	ParamSpecGtype      func(name, nick, blurb string, isAType Type, flags ParamFlags) *ParamSpec
	ParamSpecInt        func(name, nick, blurb string, minimum, maximum, defaultValue int, flags ParamFlags) *ParamSpec
	ParamSpecInt64      func(name, nick, blurb string, minimum, maximum, defaultValue int64, flags ParamFlags) *ParamSpec
	ParamSpecInternal   func(paramType Type, name, nick, blurb string, flags ParamFlags) T.Gpointer
	ParamSpecLong       func(name, nick, blurb string, minimum, maximum, defaultValue T.Glong, flags ParamFlags) *ParamSpec
	ParamSpecObject     func(name, nick, blurb string, objectType Type, flags ParamFlags) *ParamSpec
	ParamSpecOverride   func(name string, overridden *ParamSpec) *ParamSpec
	ParamSpecParam      func(name, nick, blurb string, paramType Type, flags ParamFlags) *ParamSpec
	ParamSpecPointer    func(name, nick, blurb string, flags ParamFlags) *ParamSpec
	ParamSpecString     func(name, nick, blurb, defaultValue string, flags ParamFlags) *ParamSpec
	ParamSpecUchar      func(name, nick, blurb string, minimum, maximum, defaultValue uint8, flags ParamFlags) *ParamSpec
	ParamSpecUint       func(name, nick, blurb string, minimum, maximum, defaultValue uint, flags ParamFlags) *ParamSpec
	ParamSpecUint64     func(name, nick, blurb string, minimum, maximum, defaultValue uint64, flags ParamFlags) *ParamSpec
	ParamSpecUlong      func(name, nick, blurb string, minimum, maximum, defaultValue T.Gulong, flags ParamFlags) *ParamSpec
	ParamSpecUnichar    func(name, nick, blurb string, defaultValue T.Gunichar, flags ParamFlags) *ParamSpec
	ParamSpecValueArray func(name, nick, blurb string, elementSpec *ParamSpec, flags ParamFlags) *ParamSpec
	ParamSpecVariant    func(name, nick, blurb string, t *T.GVariantType, defaultValue *T.GVariant, flags ParamFlags) *ParamSpec

	paramSpecGetBlurb          func(p *ParamSpec) string
	paramSpecGetName           func(p *ParamSpec) string
	paramSpecGetNick           func(p *ParamSpec) string
	paramSpecGetQdata          func(p *ParamSpec, quark T.GQuark) T.Gpointer
	paramSpecGetRedirectTarget func(p *ParamSpec) *ParamSpec
	paramSpecRef               func(p *ParamSpec) *ParamSpec
	paramSpecRefSink           func(p *ParamSpec) *ParamSpec
	paramSpecSetQdata          func(p *ParamSpec, quark T.GQuark, data T.Gpointer)
	paramSpecSetQdataFull      func(p *ParamSpec, quark T.GQuark, data T.Gpointer, destroy T.GDestroyNotify)
	paramSpecSink              func(p *ParamSpec)
	paramSpecStealQdata        func(p *ParamSpec, quark T.GQuark) T.Gpointer
	paramSpecUnref             func(p *ParamSpec)
	paramValueSetDefault       func(p *ParamSpec, value *Value)
	paramValueDefaults         func(p *ParamSpec, value *Value) T.Gboolean
	paramValueValidate         func(p *ParamSpec, value *Value) T.Gboolean
	paramValueConvert          func(p *ParamSpec, srcValue, destValue *Value, strictValidation T.Gboolean) T.Gboolean
	paramValuesCmp             func(p *ParamSpec, value1, value2 *Value) int
)

func (p *ParamSpec) Convert(srcValue, destValue *Value, strictValidation T.Gboolean) T.Gboolean {
	return paramValueConvert(p, srcValue, destValue, strictValidation)
}
func (p *ParamSpec) Defaults(value *Value) T.Gboolean         { return paramValueDefaults(p, value) }
func (p *ParamSpec) GetBlurb() string                         { return paramSpecGetBlurb(p) }
func (p *ParamSpec) GetName() string                          { return paramSpecGetName(p) }
func (p *ParamSpec) GetNick() string                          { return paramSpecGetNick(p) }
func (p *ParamSpec) GetQdata(quark T.GQuark) T.Gpointer       { return paramSpecGetQdata(p, quark) }
func (p *ParamSpec) GetRedirectTarget() *ParamSpec            { return paramSpecGetRedirectTarget(p) }
func (p *ParamSpec) Ref() *ParamSpec                          { return paramSpecRef(p) }
func (p *ParamSpec) RefSink() *ParamSpec                      { return paramSpecRefSink(p) }
func (p *ParamSpec) SetDefault(value *Value)                  { paramValueSetDefault(p, value) }
func (p *ParamSpec) SetQdata(quark T.GQuark, data T.Gpointer) { paramSpecSetQdata(p, quark, data) }
func (p *ParamSpec) SetQdataFull(quark T.GQuark, data T.Gpointer, destroy T.GDestroyNotify) {
	paramSpecSetQdataFull(p, quark, data, destroy)
}
func (p *ParamSpec) Sink()                                { paramSpecSink(p) }
func (p *ParamSpec) StealQdata(quark T.GQuark) T.Gpointer { return paramSpecStealQdata(p, quark) }
func (p *ParamSpec) Unref()                               { paramSpecUnref(p) }
func (p *ParamSpec) Validate(value *Value) T.Gboolean     { return paramValueValidate(p, value) }
func (p *ParamSpec) ValuesCmp(value1, value2 *Value) int  { return paramValuesCmp(p, value1, value2) }

type ParamSpecPool struct{}

var (
	ParamSpecPoolNew func(typePrefixing T.Gboolean) *ParamSpecPool

	paramSpecPoolInsert    func(p *ParamSpecPool, pspec *ParamSpec, ownerType Type)
	paramSpecPoolList      func(p *ParamSpecPool, ownerType Type, nPspecsP *uint) **ParamSpec
	paramSpecPoolListOwned func(p *ParamSpecPool, ownerType Type) *T.GList
	paramSpecPoolLookup    func(p *ParamSpecPool, paramName string, ownerType Type, walkAncestors T.Gboolean) *ParamSpec
	paramSpecPoolRemove    func(p *ParamSpecPool, pspec *ParamSpec)
)

func (p *ParamSpecPool) Insert(pspec *ParamSpec, ownerType Type) {
	paramSpecPoolInsert(p, pspec, ownerType)
}
func (p *ParamSpecPool) List(ownerType Type, nPspecsP *uint) **ParamSpec {
	return paramSpecPoolList(p, ownerType, nPspecsP)
}
func (p *ParamSpecPool) ListOwned(ownerType Type) *T.GList {
	return paramSpecPoolListOwned(p, ownerType)
}
func (p *ParamSpecPool) Lookup(paramName string, ownerType Type, walkAncestors T.Gboolean) *ParamSpec {
	return paramSpecPoolLookup(p, paramName, ownerType, walkAncestors)
}
func (p *ParamSpecPool) Remove(pspec *ParamSpec) { paramSpecPoolRemove(p, pspec) }

type ParamSpecTypeInfo struct {
	InstanceSize      uint16
	NPreallocs        uint16
	InstanceInit      func(pspec *ParamSpec)
	ValueType         Type
	Finalize          func(pspec *ParamSpec)
	Value_set_default func(
		pspec *ParamSpec, value *T.GValue)
	Value_validate func(
		pspec *ParamSpec, value *T.GValue) T.Gboolean
	Values_cmp func(pspec *ParamSpec,
		value1 *T.GValue, value2 *T.GValue) int
}

var (
	ParamTypeRegisterStatic func(
		name string, pspecInfo *ParamSpecTypeInfo) Type
)
