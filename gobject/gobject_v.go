// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	T "github.com/tHinqa/outside-gtk2/types"

// . "github.com/tHinqa/outside/types"
)

type Value struct {
	Type Type
	//  UNION
	Data [2]uint64 // was union{v_int...}data[2]
	// Int     int
	// Uint    uint
	// Long    T.Glong
	// Ulong   T.Gulong
	// Int64   int64
	// Uint64  uint64
	// Float   float32
	// Double  float64
	// Pointer T.Gpointer
}

var (
	ValueGetType func() Type

	ValueRegisterTransformFunc func(srcType Type, destType Type, transformFunc ValueTransformFunc)
	ValueTypeCompatible        func(srcType Type, destType Type) bool
	ValueTypeTransformable     func(srcType Type, destType Type) bool

	ValueCopy                   func(v *Value, destValue *Value)
	ValueDupBoxed               func(v *Value) T.Gpointer
	ValueDupObject              func(v *Value) T.Gpointer
	ValueDupParam               func(v *Value) *ParamSpec
	ValueDupString              func(v *Value) string
	ValueDupVariant             func(v *Value) *T.GVariant
	ValueFitsPointer            func(v *Value) bool
	ValueGetBoolean             func(v *Value) bool
	ValueGetBoxed               func(v *Value) T.Gpointer
	ValueGetChar                func(v *Value) T.Gchar
	ValueGetDouble              func(v *Value) float64
	ValueGetEnum                func(v *Value) int
	ValueGetFlags               func(v *Value) uint
	ValueGetFloat               func(v *Value) float32
	ValueGetGtype               func(v *Value) Type
	ValueGetInt                 func(v *Value) int
	ValueGetInt64               func(v *Value) int64
	ValueGetLong                func(v *Value) T.Glong
	ValueGetObject              func(v *Value) T.Gpointer
	ValueGetParam               func(v *Value) *ParamSpec
	ValueGetPointer             func(v *Value) T.Gpointer
	ValueGetString              func(v *Value) string
	ValueGetUchar               func(v *Value) T.Guchar
	ValueGetUint                func(v *Value) uint
	ValueGetUint64              func(v *Value) uint64
	ValueGetUlong               func(v *Value) T.Gulong
	ValueGetVariant             func(v *Value) *T.GVariant
	ValueInit                   func(v *Value, gType Type) *Value
	ValuePeekPointer            func(v *Value) T.Gpointer
	ValueReset                  func(v *Value) *Value
	ValueSetBoolean             func(v *Value, vBoolean bool)
	ValueSetBoxed               func(v *Value, vBoxed T.Gconstpointer)
	ValueSetBoxedTakeOwnership  func(v *Value, vBoxed T.Gconstpointer)
	ValueSetChar                func(v *Value, vChar T.Gchar)
	ValueSetDouble              func(v *Value, vDouble float64)
	ValueSetEnum                func(v *Value, vEnum int)
	ValueSetFlags               func(v *Value, vFlags uint)
	ValueSetFloat               func(v *Value, vFloat float32)
	ValueSetGtype               func(v *Value, vGtype Type)
	ValueSetInstance            func(v *Value, instance T.Gpointer)
	ValueSetInt                 func(v *Value, vInt int)
	ValueSetInt64               func(v *Value, vInt64 int64)
	ValueSetLong                func(v *Value, vLong T.Glong)
	ValueSetObject              func(v *Value, vObject T.Gpointer)
	ValueSetObjectTakeOwnership func(v *Value, vObject T.Gpointer)
	ValueSetParam               func(v *Value, param *ParamSpec)
	ValueSetParamTakeOwnership  func(v *Value, param *ParamSpec)
	ValueSetPointer             func(v *Value, vPointer T.Gpointer)
	ValueSetStaticBoxed         func(v *Value, vBoxed T.Gconstpointer)
	ValueSetStaticString        func(v *Value, vString string)
	ValueSetString              func(v *Value, vString string)
	ValueSetStringTakeOwnership func(v *Value, vString string)
	ValueSetUchar               func(v *Value, vUchar T.Guchar)
	ValueSetUint                func(v *Value, vUint uint)
	ValueSetUint64              func(v *Value, vUint64 uint64)
	ValueSetUlong               func(v *Value, vUlong T.Gulong)
	ValueSetVariant             func(v *Value, variant *T.GVariant)
	ValueTakeBoxed              func(v *Value, vBoxed T.Gconstpointer)
	ValueTakeObject             func(v *Value, vObject T.Gpointer)
	ValueTakeParam              func(v *Value, param *ParamSpec)
	ValueTakeString             func(v *Value, vString string)
	ValueTakeVariant            func(v *Value, variant *T.GVariant)
	ValueTransform              func(v *Value, dest *Value) bool //
	ValueUnset                  func(v *Value)
)

func (v *Value) Copy(destValue *Value)                        { ValueCopy(v, destValue) }
func (v *Value) DupBoxed() T.Gpointer                         { return ValueDupBoxed(v) }
func (v *Value) DupObject() T.Gpointer                        { return ValueDupObject(v) }
func (v *Value) DupParam() *ParamSpec                         { return ValueDupParam(v) }
func (v *Value) DupString() string                            { return ValueDupString(v) }
func (v *Value) DupVariant() *T.GVariant                      { return ValueDupVariant(v) }
func (v *Value) FitsPointer() bool                            { return ValueFitsPointer(v) }
func (v *Value) GetBoolean() bool                             { return ValueGetBoolean(v) }
func (v *Value) GetBoxed() T.Gpointer                         { return ValueGetBoxed(v) }
func (v *Value) GetChar() T.Gchar                             { return ValueGetChar(v) }
func (v *Value) GetDouble() float64                           { return ValueGetDouble(v) }
func (v *Value) GetEnum() int                                 { return ValueGetEnum(v) }
func (v *Value) GetFlags() uint                               { return ValueGetFlags(v) }
func (v *Value) GetFloat() float32                            { return ValueGetFloat(v) }
func (v *Value) GetGtype() Type                               { return ValueGetGtype(v) }
func (v *Value) GetInt() int                                  { return ValueGetInt(v) }
func (v *Value) GetInt64() int64                              { return ValueGetInt64(v) }
func (v *Value) GetLong() T.Glong                             { return ValueGetLong(v) }
func (v *Value) GetObject() T.Gpointer                        { return ValueGetObject(v) }
func (v *Value) GetParam() *ParamSpec                         { return ValueGetParam(v) }
func (v *Value) GetPointer() T.Gpointer                       { return ValueGetPointer(v) }
func (v *Value) GetString() string                            { return ValueGetString(v) }
func (v *Value) GetUchar() T.Guchar                           { return ValueGetUchar(v) }
func (v *Value) GetUint() uint                                { return ValueGetUint(v) }
func (v *Value) GetUint64() uint64                            { return ValueGetUint64(v) }
func (v *Value) GetUlong() T.Gulong                           { return ValueGetUlong(v) }
func (v *Value) GetVariant() *T.GVariant                      { return ValueGetVariant(v) }
func (v *Value) Init(gType Type) *Value                       { return ValueInit(v, gType) }
func (v *Value) PeekPointer() T.Gpointer                      { return ValuePeekPointer(v) }
func (v *Value) Reset() *Value                                { return ValueReset(v) }
func (v *Value) SetBoolean(vBoolean bool)                     { ValueSetBoolean(v, vBoolean) }
func (v *Value) SetBoxed(vBoxed T.Gconstpointer)              { ValueSetBoxed(v, vBoxed) }
func (v *Value) SetBoxedTakeOwnership(vBoxed T.Gconstpointer) { ValueSetBoxedTakeOwnership(v, vBoxed) }
func (v *Value) SetChar(vChar T.Gchar)                        { ValueSetChar(v, vChar) }
func (v *Value) SetDouble(vDouble float64)                    { ValueSetDouble(v, vDouble) }
func (v *Value) SetEnum(vEnum int)                            { ValueSetEnum(v, vEnum) }
func (v *Value) SetFlags(vFlags uint)                         { ValueSetFlags(v, vFlags) }
func (v *Value) SetFloat(vFloat float32)                      { ValueSetFloat(v, vFloat) }
func (v *Value) SetGtype(vGtype Type)                         { ValueSetGtype(v, vGtype) }
func (v *Value) SetInstance(instance T.Gpointer)              { ValueSetInstance(v, instance) }
func (v *Value) SetInt(vInt int)                              { ValueSetInt(v, vInt) }
func (v *Value) SetInt64(vInt64 int64)                        { ValueSetInt64(v, vInt64) }
func (v *Value) SetLong(vLong T.Glong)                        { ValueSetLong(v, vLong) }
func (v *Value) SetObject(vObject T.Gpointer)                 { ValueSetObject(v, vObject) }
func (v *Value) SetObjectTakeOwnership(vObject T.Gpointer)    { ValueSetObjectTakeOwnership(v, vObject) }
func (v *Value) SetParam(param *ParamSpec)                    { ValueSetParam(v, param) }
func (v *Value) SetParamTakeOwnership(param *ParamSpec)       { ValueSetParamTakeOwnership(v, param) }
func (v *Value) SetPointer(vPointer T.Gpointer)               { ValueSetPointer(v, vPointer) }
func (v *Value) SetStaticBoxed(vBoxed T.Gconstpointer)        { ValueSetStaticBoxed(v, vBoxed) }
func (v *Value) SetStaticString(vString string)               { ValueSetStaticString(v, vString) }
func (v *Value) SetString(vString string)                     { ValueSetString(v, vString) }
func (v *Value) SetStringTakeOwnership(vString string)        { ValueSetStringTakeOwnership(v, vString) }
func (v *Value) SetUchar(vUchar T.Guchar)                     { ValueSetUchar(v, vUchar) }
func (v *Value) SetUint(vUint uint)                           { ValueSetUint(v, vUint) }
func (v *Value) SetUint64(vUint64 uint64)                     { ValueSetUint64(v, vUint64) }
func (v *Value) SetUlong(vUlong T.Gulong)                     { ValueSetUlong(v, vUlong) }
func (v *Value) SetVariant(variant *T.GVariant)               { ValueSetVariant(v, variant) }
func (v *Value) TakeBoxed(vBoxed T.Gconstpointer)             { ValueTakeBoxed(v, vBoxed) }
func (v *Value) TakeObject(vObject T.Gpointer)                { ValueTakeObject(v, vObject) }
func (v *Value) TakeParam(param *ParamSpec)                   { ValueTakeParam(v, param) }
func (v *Value) TakeString(vString string)                    { ValueTakeString(v, vString) }
func (v *Value) TakeVariant(variant *T.GVariant)              { ValueTakeVariant(v, variant) }
func (v *Value) Transform(dest *Value) bool                   { return ValueTransform(v, dest) }
func (v *Value) Unset()                                       { ValueUnset(v) }

type ValueArray struct {
	NValues     uint
	Values      *Value
	NPrealloced uint
}

var (
	ValueArrayGetType func() Type
	ValueArrayNew     func(nPrealloced uint) *ValueArray

	ValueArrayGetNth       func(v *ValueArray, index uint) *Value
	ValueArrayFree         func(v *ValueArray)
	ValueArrayCopy         func(v *ValueArray) *ValueArray
	ValueArrayPrepend      func(v *ValueArray, value *Value) *ValueArray
	ValueArrayAppend       func(v *ValueArray, value *Value) *ValueArray
	ValueArrayInsert       func(v *ValueArray, index uint, value *Value) *ValueArray
	ValueArrayRemove       func(v *ValueArray, index uint) *ValueArray
	ValueArraySort         func(v *ValueArray, compareFunc T.GCompareFunc) *ValueArray
	ValueArraySortWithData func(v *ValueArray, compareFunc T.GCompareDataFunc, userData T.Gpointer) *ValueArray
)

func (v *ValueArray) Append(value *Value) *ValueArray { return ValueArrayAppend(v, value) }
func (v *ValueArray) Copy() *ValueArray               { return ValueArrayCopy(v) }
func (v *ValueArray) Free()                           { ValueArrayFree(v) }
func (v *ValueArray) GetNth(index uint) *Value        { return ValueArrayGetNth(v, index) }
func (v *ValueArray) Insert(index uint, value *Value) *ValueArray {
	return ValueArrayInsert(v, index, value)
}
func (v *ValueArray) Prepend(value *Value) *ValueArray { return ValueArrayPrepend(v, value) }
func (v *ValueArray) Remove(index uint) *ValueArray    { return ValueArrayRemove(v, index) }
func (v *ValueArray) Sort(compareFunc T.GCompareFunc) *ValueArray {
	return ValueArraySort(v, compareFunc)
}
func (v *ValueArray) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) *ValueArray {
	return ValueArraySortWithData(v, compareFunc, userData)
}

type ValueTransformFunc func(srcValue, destValue *Value) //NOTE(t): was ambiguity
