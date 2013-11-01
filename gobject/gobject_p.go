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
	ParamSpecBoolean    func(name, nick, blurb string, defaultValue bool, flags ParamFlags) *ParamSpec
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
	ParamSpecVariant    func(name, nick, blurb string, t *T.VariantType, defaultValue *T.Variant, flags ParamFlags) *ParamSpec

	ParamSpecGetBlurb          func(p *ParamSpec) string
	ParamSpecGetName           func(p *ParamSpec) string
	ParamSpecGetNick           func(p *ParamSpec) string
	ParamSpecGetQdata          func(p *ParamSpec, quark T.Quark) T.Gpointer
	ParamSpecGetRedirectTarget func(p *ParamSpec) *ParamSpec
	ParamSpecRef               func(p *ParamSpec) *ParamSpec
	ParamSpecRefSink           func(p *ParamSpec) *ParamSpec
	ParamSpecSetQdata          func(p *ParamSpec, quark T.Quark, data T.Gpointer)
	ParamSpecSetQdataFull      func(p *ParamSpec, quark T.Quark, data T.Gpointer, destroy T.GDestroyNotify)
	ParamSpecSink              func(p *ParamSpec)
	ParamSpecStealQdata        func(p *ParamSpec, quark T.Quark) T.Gpointer
	ParamSpecUnref             func(p *ParamSpec)
	ParamValueSetDefault       func(p *ParamSpec, value *Value)
	ParamValueDefaults         func(p *ParamSpec, value *Value) bool
	ParamValueValidate         func(p *ParamSpec, value *Value) bool
	ParamValueConvert          func(p *ParamSpec, srcValue, destValue *Value, strictValidation bool) bool
	ParamValuesCmp             func(p *ParamSpec, value1, value2 *Value) int
)

func (p *ParamSpec) Convert(srcValue, destValue *Value, strictValidation bool) bool {
	return ParamValueConvert(p, srcValue, destValue, strictValidation)
}
func (p *ParamSpec) Defaults(value *Value) bool              { return ParamValueDefaults(p, value) }
func (p *ParamSpec) GetBlurb() string                        { return ParamSpecGetBlurb(p) }
func (p *ParamSpec) GetName() string                         { return ParamSpecGetName(p) }
func (p *ParamSpec) GetNick() string                         { return ParamSpecGetNick(p) }
func (p *ParamSpec) GetQdata(quark T.Quark) T.Gpointer       { return ParamSpecGetQdata(p, quark) }
func (p *ParamSpec) GetRedirectTarget() *ParamSpec           { return ParamSpecGetRedirectTarget(p) }
func (p *ParamSpec) Ref() *ParamSpec                         { return ParamSpecRef(p) }
func (p *ParamSpec) RefSink() *ParamSpec                     { return ParamSpecRefSink(p) }
func (p *ParamSpec) SetDefault(value *Value)                 { ParamValueSetDefault(p, value) }
func (p *ParamSpec) SetQdata(quark T.Quark, data T.Gpointer) { ParamSpecSetQdata(p, quark, data) }
func (p *ParamSpec) SetQdataFull(quark T.Quark, data T.Gpointer, destroy T.GDestroyNotify) {
	ParamSpecSetQdataFull(p, quark, data, destroy)
}
func (p *ParamSpec) Sink()                               { ParamSpecSink(p) }
func (p *ParamSpec) StealQdata(quark T.Quark) T.Gpointer { return ParamSpecStealQdata(p, quark) }
func (p *ParamSpec) Unref()                              { ParamSpecUnref(p) }
func (p *ParamSpec) Validate(value *Value) bool          { return ParamValueValidate(p, value) }
func (p *ParamSpec) ValuesCmp(value1, value2 *Value) int { return ParamValuesCmp(p, value1, value2) }

type ParamSpecPool struct{}

var (
	ParamSpecPoolNew func(typePrefixing bool) *ParamSpecPool

	ParamSpecPoolInsert    func(p *ParamSpecPool, pspec *ParamSpec, ownerType Type)
	ParamSpecPoolList      func(p *ParamSpecPool, ownerType Type, nPspecsP *uint) **ParamSpec
	ParamSpecPoolListOwned func(p *ParamSpecPool, ownerType Type) *T.GList
	ParamSpecPoolLookup    func(p *ParamSpecPool, paramName string, ownerType Type, walkAncestors bool) *ParamSpec
	ParamSpecPoolRemove    func(p *ParamSpecPool, pspec *ParamSpec)
)

func (p *ParamSpecPool) Insert(pspec *ParamSpec, ownerType Type) {
	ParamSpecPoolInsert(p, pspec, ownerType)
}
func (p *ParamSpecPool) List(ownerType Type, nPspecsP *uint) **ParamSpec {
	return ParamSpecPoolList(p, ownerType, nPspecsP)
}
func (p *ParamSpecPool) ListOwned(ownerType Type) *T.GList {
	return ParamSpecPoolListOwned(p, ownerType)
}
func (p *ParamSpecPool) Lookup(paramName string, ownerType Type, walkAncestors bool) *ParamSpec {
	return ParamSpecPoolLookup(p, paramName, ownerType, walkAncestors)
}
func (p *ParamSpecPool) Remove(pspec *ParamSpec) { ParamSpecPoolRemove(p, pspec) }

type ParamSpecTypeInfo struct {
	InstanceSize      uint16
	NPreallocs        uint16
	InstanceInit      func(pspec *ParamSpec)
	ValueType         Type
	Finalize          func(pspec *ParamSpec)
	Value_set_default func(
		pspec *ParamSpec, value *T.GValue)
	Value_validate func(
		pspec *ParamSpec, value *T.GValue) bool
	Values_cmp func(pspec *ParamSpec,
		value1 *T.GValue, value2 *T.GValue) int
}

var (
	ParamTypeRegisterStatic func(
		name string, pspecInfo *ParamSpecTypeInfo) Type
)
