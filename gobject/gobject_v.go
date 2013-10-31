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

	valueCopy                   func(v *Value, destValue *Value)
	valueDupBoxed               func(v *Value) T.Gpointer
	valueDupObject              func(v *Value) T.Gpointer
	valueDupParam               func(v *Value) *ParamSpec
	valueDupString              func(v *Value) string
	valueDupVariant             func(v *Value) *T.GVariant
	valueFitsPointer            func(v *Value) bool
	valueGetBoolean             func(v *Value) bool
	valueGetBoxed               func(v *Value) T.Gpointer
	valueGetChar                func(v *Value) T.Gchar
	valueGetDouble              func(v *Value) float64
	valueGetEnum                func(v *Value) int
	valueGetFlags               func(v *Value) uint
	valueGetFloat               func(v *Value) float32
	valueGetGtype               func(v *Value) Type
	valueGetInt                 func(v *Value) int
	valueGetInt64               func(v *Value) int64
	valueGetLong                func(v *Value) T.Glong
	valueGetObject              func(v *Value) T.Gpointer
	valueGetParam               func(v *Value) *ParamSpec
	valueGetPointer             func(v *Value) T.Gpointer
	valueGetString              func(v *Value) string
	valueGetUchar               func(v *Value) T.Guchar
	valueGetUint                func(v *Value) uint
	valueGetUint64              func(v *Value) uint64
	valueGetUlong               func(v *Value) T.Gulong
	valueGetVariant             func(v *Value) *T.GVariant
	valueInit                   func(v *Value, gType Type) *Value
	valuePeekPointer            func(v *Value) T.Gpointer
	valueReset                  func(v *Value) *Value
	valueSetBoolean             func(v *Value, vBoolean bool)
	valueSetBoxed               func(v *Value, vBoxed T.Gconstpointer)
	valueSetBoxedTakeOwnership  func(v *Value, vBoxed T.Gconstpointer)
	valueSetChar                func(v *Value, vChar T.Gchar)
	valueSetDouble              func(v *Value, vDouble float64)
	valueSetEnum                func(v *Value, vEnum int)
	valueSetFlags               func(v *Value, vFlags uint)
	valueSetFloat               func(v *Value, vFloat float32)
	valueSetGtype               func(v *Value, vGtype Type)
	valueSetInstance            func(v *Value, instance T.Gpointer)
	valueSetInt                 func(v *Value, vInt int)
	valueSetInt64               func(v *Value, vInt64 int64)
	valueSetLong                func(v *Value, vLong T.Glong)
	valueSetObject              func(v *Value, vObject T.Gpointer)
	valueSetObjectTakeOwnership func(v *Value, vObject T.Gpointer)
	valueSetParam               func(v *Value, param *ParamSpec)
	valueSetParamTakeOwnership  func(v *Value, param *ParamSpec)
	valueSetPointer             func(v *Value, vPointer T.Gpointer)
	valueSetStaticBoxed         func(v *Value, vBoxed T.Gconstpointer)
	valueSetStaticString        func(v *Value, vString string)
	valueSetString              func(v *Value, vString string)
	valueSetStringTakeOwnership func(v *Value, vString string)
	valueSetUchar               func(v *Value, vUchar T.Guchar)
	valueSetUint                func(v *Value, vUint uint)
	valueSetUint64              func(v *Value, vUint64 uint64)
	valueSetUlong               func(v *Value, vUlong T.Gulong)
	valueSetVariant             func(v *Value, variant *T.GVariant)
	valueTakeBoxed              func(v *Value, vBoxed T.Gconstpointer)
	valueTakeObject             func(v *Value, vObject T.Gpointer)
	valueTakeParam              func(v *Value, param *ParamSpec)
	valueTakeString             func(v *Value, vString string)
	valueTakeVariant            func(v *Value, variant *T.GVariant)
	valueTransform              func(v *Value, dest *Value) bool //
	valueUnset                  func(v *Value)
)

func (v *Value) Copy(destValue *Value)                        { valueCopy(v, destValue) }
func (v *Value) DupBoxed() T.Gpointer                         { return valueDupBoxed(v) }
func (v *Value) DupObject() T.Gpointer                        { return valueDupObject(v) }
func (v *Value) DupParam() *ParamSpec                         { return valueDupParam(v) }
func (v *Value) DupString() string                            { return valueDupString(v) }
func (v *Value) DupVariant() *T.GVariant                      { return valueDupVariant(v) }
func (v *Value) FitsPointer() bool                            { return valueFitsPointer(v) }
func (v *Value) GetBoolean() bool                             { return valueGetBoolean(v) }
func (v *Value) GetBoxed() T.Gpointer                         { return valueGetBoxed(v) }
func (v *Value) GetChar() T.Gchar                             { return valueGetChar(v) }
func (v *Value) GetDouble() float64                           { return valueGetDouble(v) }
func (v *Value) GetEnum() int                                 { return valueGetEnum(v) }
func (v *Value) GetFlags() uint                               { return valueGetFlags(v) }
func (v *Value) GetFloat() float32                            { return valueGetFloat(v) }
func (v *Value) GetGtype() Type                               { return valueGetGtype(v) }
func (v *Value) GetInt() int                                  { return valueGetInt(v) }
func (v *Value) GetInt64() int64                              { return valueGetInt64(v) }
func (v *Value) GetLong() T.Glong                             { return valueGetLong(v) }
func (v *Value) GetObject() T.Gpointer                        { return valueGetObject(v) }
func (v *Value) GetParam() *ParamSpec                         { return valueGetParam(v) }
func (v *Value) GetPointer() T.Gpointer                       { return valueGetPointer(v) }
func (v *Value) GetString() string                            { return valueGetString(v) }
func (v *Value) GetUchar() T.Guchar                           { return valueGetUchar(v) }
func (v *Value) GetUint() uint                                { return valueGetUint(v) }
func (v *Value) GetUint64() uint64                            { return valueGetUint64(v) }
func (v *Value) GetUlong() T.Gulong                           { return valueGetUlong(v) }
func (v *Value) GetVariant() *T.GVariant                      { return valueGetVariant(v) }
func (v *Value) Init(gType Type) *Value                       { return valueInit(v, gType) }
func (v *Value) PeekPointer() T.Gpointer                      { return valuePeekPointer(v) }
func (v *Value) Reset() *Value                                { return valueReset(v) }
func (v *Value) SetBoolean(vBoolean bool)                     { valueSetBoolean(v, vBoolean) }
func (v *Value) SetBoxed(vBoxed T.Gconstpointer)              { valueSetBoxed(v, vBoxed) }
func (v *Value) SetBoxedTakeOwnership(vBoxed T.Gconstpointer) { valueSetBoxedTakeOwnership(v, vBoxed) }
func (v *Value) SetChar(vChar T.Gchar)                        { valueSetChar(v, vChar) }
func (v *Value) SetDouble(vDouble float64)                    { valueSetDouble(v, vDouble) }
func (v *Value) SetEnum(vEnum int)                            { valueSetEnum(v, vEnum) }
func (v *Value) SetFlags(vFlags uint)                         { valueSetFlags(v, vFlags) }
func (v *Value) SetFloat(vFloat float32)                      { valueSetFloat(v, vFloat) }
func (v *Value) SetGtype(vGtype Type)                         { valueSetGtype(v, vGtype) }
func (v *Value) SetInstance(instance T.Gpointer)              { valueSetInstance(v, instance) }
func (v *Value) SetInt(vInt int)                              { valueSetInt(v, vInt) }
func (v *Value) SetInt64(vInt64 int64)                        { valueSetInt64(v, vInt64) }
func (v *Value) SetLong(vLong T.Glong)                        { valueSetLong(v, vLong) }
func (v *Value) SetObject(vObject T.Gpointer)                 { valueSetObject(v, vObject) }
func (v *Value) SetObjectTakeOwnership(vObject T.Gpointer)    { valueSetObjectTakeOwnership(v, vObject) }
func (v *Value) SetParam(param *ParamSpec)                    { valueSetParam(v, param) }
func (v *Value) SetParamTakeOwnership(param *ParamSpec)       { valueSetParamTakeOwnership(v, param) }
func (v *Value) SetPointer(vPointer T.Gpointer)               { valueSetPointer(v, vPointer) }
func (v *Value) SetStaticBoxed(vBoxed T.Gconstpointer)        { valueSetStaticBoxed(v, vBoxed) }
func (v *Value) SetStaticString(vString string)               { valueSetStaticString(v, vString) }
func (v *Value) SetString(vString string)                     { valueSetString(v, vString) }
func (v *Value) SetStringTakeOwnership(vString string)        { valueSetStringTakeOwnership(v, vString) }
func (v *Value) SetUchar(vUchar T.Guchar)                     { valueSetUchar(v, vUchar) }
func (v *Value) SetUint(vUint uint)                           { valueSetUint(v, vUint) }
func (v *Value) SetUint64(vUint64 uint64)                     { valueSetUint64(v, vUint64) }
func (v *Value) SetUlong(vUlong T.Gulong)                     { valueSetUlong(v, vUlong) }
func (v *Value) SetVariant(variant *T.GVariant)               { valueSetVariant(v, variant) }
func (v *Value) TakeBoxed(vBoxed T.Gconstpointer)             { valueTakeBoxed(v, vBoxed) }
func (v *Value) TakeObject(vObject T.Gpointer)                { valueTakeObject(v, vObject) }
func (v *Value) TakeParam(param *ParamSpec)                   { valueTakeParam(v, param) }
func (v *Value) TakeString(vString string)                    { valueTakeString(v, vString) }
func (v *Value) TakeVariant(variant *T.GVariant)              { valueTakeVariant(v, variant) }
func (v *Value) Transform(dest *Value) bool                   { return valueTransform(v, dest) }
func (v *Value) Unset()                                       { valueUnset(v) }

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
