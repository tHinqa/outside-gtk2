// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type (
	Variant T.Variant // chicken/egg imports gobject/glib
)

type VariantClass Enum

const (
	VARIANT_CLASS_BOOLEAN     VariantClass = 'b'
	VARIANT_CLASS_BYTE        VariantClass = 'y'
	VARIANT_CLASS_INT16       VariantClass = 'n'
	VARIANT_CLASS_UINT16      VariantClass = 'q'
	VARIANT_CLASS_INT32       VariantClass = 'i'
	VARIANT_CLASS_UINT32      VariantClass = 'u'
	VARIANT_CLASS_INT64       VariantClass = 'x'
	VARIANT_CLASS_UINT64      VariantClass = 't'
	VARIANT_CLASS_HANDLE      VariantClass = 'h'
	VARIANT_CLASS_DOUBLE      VariantClass = 'd'
	VARIANT_CLASS_STRING      VariantClass = 's'
	VARIANT_CLASS_OBJECT_PATH VariantClass = 'o'
	VARIANT_CLASS_SIGNATURE   VariantClass = 'g'
	VARIANT_CLASS_VARIANT     VariantClass = 'v'
	VARIANT_CLASS_MAYBE       VariantClass = 'm'
	VARIANT_CLASS_ARRAY       VariantClass = 'a'
	VARIANT_CLASS_TUPLE       VariantClass = '('
	VARIANT_CLASS_DICT_ENTRY  VariantClass = '{'
)

var (
	VariantNew                func(formatString string, v ...VArg) *Variant
	VariantNewArray           func(childType *VariantType, children **Variant, nChildren T.Gsize) *Variant
	VariantNewBoolean         func(value bool) *Variant
	VariantNewByte            func(value T.Guchar) *Variant
	VariantNewBytestring      func(str string) *Variant
	VariantNewBytestringArray func(strv []string, length T.Gssize) *Variant
	VariantNewDouble          func(value float64) *Variant
	VariantNewFromData        func(typ *VariantType, data T.Gconstpointer, size T.Gsize, trusted bool, notify O.DestroyNotify, userData T.Gpointer) *Variant
	VariantNewHandle          func(value T.GInt32) *Variant
	VariantNewInt16           func(value int16) *Variant
	VariantNewInt32           func(value T.GInt32) *Variant
	VariantNewInt64           func(value int64) *Variant
	VariantNewMaybe           func(childType *VariantType, child *Variant) *Variant
	VariantNewObjectPath      func(objectPath string) *Variant
	VariantNewParsed          func(format string, v ...VArg) *Variant
	VariantNewParsedVa        func(format string, app *VAList) *Variant
	VariantNewSignature       func(signature string) *Variant
	VariantNewString          func(str string) *Variant
	VariantNewStrv            func(strv []string, length T.Gssize) *Variant
	VariantNewTuple           func(children **Variant, nChildren T.Gsize) *Variant
	VariantNewUint16          func(value uint16) *Variant
	VariantNewUint32          func(value T.GUint32) *Variant
	VariantNewUint64          func(value uint64) *Variant
	VariantNewVa              func(formatString string, endptr []string, app *VAList) *Variant

	VariantCompare             func(one T.Gconstpointer, two T.Gconstpointer) int
	VariantEqual               func(one T.Gconstpointer, two T.Gconstpointer) bool
	VariantHash                func(value T.Gconstpointer) uint
	VariantIsObjectPath        func(str string) bool
	VariantIsSignature         func(str string) bool
	VariantParse               func(typ *VariantType, text, limit string, endptr []string, err **Error) *Variant
	VariantParserGetErrorQuark func() Quark

	VariantByteswap           func(v *Variant) *Variant
	VariantClassify           func(v *Variant) VariantClass
	VariantDupBytestring      func(v *Variant, length *T.Gsize) string
	VariantDupBytestringArray func(v *Variant, length *T.Gsize) []string
	VariantDupString          func(v *Variant, length *T.Gsize) string
	VariantDupStrv            func(v *Variant, length *T.Gsize) []string
	VariantGet                func(v *Variant, formatString string, va ...VArg)
	VariantGetBoolean         func(v *Variant) bool
	VariantGetByte            func(v *Variant) T.Guchar
	VariantGetBytestring      func(v *Variant) string
	VariantGetBytestringArray func(v *Variant, length *T.Gsize) []string
	VariantGetChild           func(v *Variant, index T.Gsize, formatString string, va ...VArg)
	VariantGetChildValue      func(v *Variant, index T.Gsize) *Variant
	VariantGetData            func(v *Variant) T.Gconstpointer
	VariantGetDouble          func(v *Variant) float64
	VariantGetFixedArray      func(v *Variant, nElements *T.Gsize, elementSize T.Gsize) T.Gconstpointer
	VariantGetHandle          func(v *Variant) T.GInt32
	VariantGetInt16           func(v *Variant) int16
	VariantGetInt32           func(v *Variant) T.GInt32
	VariantGetInt64           func(v *Variant) int64
	VariantGetMaybe           func(v *Variant) *Variant
	VariantGetNormalForm      func(v *Variant) *Variant
	VariantGetSize            func(v *Variant) T.Gsize
	VariantGetString          func(v *Variant, length *T.Gsize) string
	VariantGetStrv            func(v *Variant, length *T.Gsize) []string
	VariantGetType            func(v *Variant) *VariantType
	VariantGetTypeString      func(v *Variant) string
	VariantGetUint16          func(v *Variant) uint16
	VariantGetUint32          func(v *Variant) T.GUint32
	VariantGetUint64          func(v *Variant) uint64
	VariantGetVa              func(v *Variant, formatString string, endptr []string, app *VAList)
	VariantGetVariant         func(v *Variant) *Variant
	VariantIsContainer        func(v *Variant) bool
	VariantIsFloating         func(v *Variant) bool
	VariantIsNormalForm       func(v *Variant) bool
	VariantIsOfType           func(v *Variant, typ *VariantType) bool
	VariantLookup             func(dictionary *Variant, key string, formatString string, v ...VArg) bool
	VariantLookupValue        func(dictionary *Variant, key string, expectedType *VariantType) *Variant
	VariantNChildren          func(v *Variant) T.Gsize
	VariantNewDictEntry       func(key *Variant, value *Variant) *Variant
	VariantNewVariant         func(v *Variant) *Variant
	VariantPrint              func(v *Variant, typeAnnotate bool) string
	VariantPrintString        func(v *Variant, str *String, typeAnnotate bool) *String
	VariantRef                func(v *Variant) *Variant
	VariantRefSink            func(v *Variant) *Variant
	VariantStore              func(v *Variant, data T.Gpointer)
	VariantUnref              func(v *Variant)
)

func (v *Variant) Byteswap() *Variant                   { return VariantByteswap(v) }
func (v *Variant) Classify() VariantClass               { return VariantClassify(v) }
func (v *Variant) DupBytestring(length *T.Gsize) string { return VariantDupBytestring(v, length) }
func (v *Variant) DupBytestringArray(length *T.Gsize) []string {
	return VariantDupBytestringArray(v, length)
}
func (v *Variant) DupString(length *T.Gsize) string    { return VariantDupString(v, length) }
func (v *Variant) DupStrv(length *T.Gsize) []string    { return VariantDupStrv(v, length) }
func (v *Variant) Get(formatString string, va ...VArg) { VariantGet(v, formatString, va) }
func (v *Variant) Boolean() bool                       { return VariantGetBoolean(v) }
func (v *Variant) Byte() T.Guchar                      { return VariantGetByte(v) }
func (v *Variant) Bytestring() string                  { return VariantGetBytestring(v) }
func (v *Variant) BytestringArray(length *T.Gsize) []string {
	return VariantGetBytestringArray(v, length)
}
func (v *Variant) Child(index T.Gsize, formatString string, va ...VArg) {
	VariantGetChild(v, index, formatString, va)
}
func (v *Variant) ChildValue(index T.Gsize) *Variant { return VariantGetChildValue(v, index) }
func (v *Variant) Data() T.Gconstpointer             { return VariantGetData(v) }
func (v *Variant) Double() float64                   { return VariantGetDouble(v) }
func (v *Variant) FixedArray(nElements *T.Gsize, elementSize T.Gsize) T.Gconstpointer {
	return VariantGetFixedArray(v, nElements, elementSize)
}
func (v *Variant) Handle() T.GInt32              { return VariantGetHandle(v) }
func (v *Variant) Int16() int16                  { return VariantGetInt16(v) }
func (v *Variant) Int32() T.GInt32               { return VariantGetInt32(v) }
func (v *Variant) Int64() int64                  { return VariantGetInt64(v) }
func (v *Variant) Maybe() *Variant               { return VariantGetMaybe(v) }
func (v *Variant) NormalForm() *Variant          { return VariantGetNormalForm(v) }
func (v *Variant) Size() T.Gsize                 { return VariantGetSize(v) }
func (v *Variant) String(length *T.Gsize) string { return VariantGetString(v, length) }
func (v *Variant) Strv(length *T.Gsize) []string { return VariantGetStrv(v, length) }
func (v *Variant) Type() *VariantType            { return VariantGetType(v) }
func (v *Variant) TypeString() string            { return VariantGetTypeString(v) }
func (v *Variant) Uint16() uint16                { return VariantGetUint16(v) }
func (v *Variant) Uint32() T.GUint32             { return VariantGetUint32(v) }
func (v *Variant) Uint64() uint64                { return VariantGetUint64(v) }
func (v *Variant) Va(formatString string, endptr []string, app *VAList) {
	VariantGetVa(v, formatString, endptr, app)
}
func (v *Variant) Variant() *Variant              { return VariantGetVariant(v) }
func (v *Variant) IsContainer() bool              { return VariantIsContainer(v) }
func (v *Variant) IsFloating() bool               { return VariantIsFloating(v) }
func (v *Variant) IsNormalForm() bool             { return VariantIsNormalForm(v) }
func (v *Variant) IsOfType(typ *VariantType) bool { return VariantIsOfType(v, typ) }
func (d *Variant) Lookup(key string, formatString string, v ...VArg) bool {
	return VariantLookup(d, key, formatString, v)
}
func (d *Variant) LookupValue(key string, expectedType *VariantType) *Variant {
	return VariantLookupValue(d, key, expectedType)
}
func (v *Variant) NChildren() T.Gsize                   { return VariantNChildren(v) }
func (k *Variant) NewDictEntry(value *Variant) *Variant { return VariantNewDictEntry(k, value) }
func (v *Variant) NewVariant() *Variant                 { return VariantNewVariant(v) }
func (v *Variant) Print(typeAnnotate bool) string       { return VariantPrint(v, typeAnnotate) }
func (v *Variant) PrintString(str *String, typeAnnotate bool) *String {
	return VariantPrintString(v, str, typeAnnotate)
}
func (v *Variant) Ref() *Variant         { return VariantRef(v) }
func (v *Variant) RefSink() *Variant     { return VariantRefSink(v) }
func (v *Variant) Store(data T.Gpointer) { VariantStore(v, data) }
func (v *Variant) Unref()                { VariantUnref(v) }

type VariantBuilder struct {
	X [16]T.Gsize
}

var (
	VariantBuilderNew func(typ *VariantType) *VariantBuilder

	VariantBuilderAdd       func(v *VariantBuilder, formatString string, va ...VArg)
	VariantBuilderAddParsed func(v *VariantBuilder, format string, va ...VArg)
	VariantBuilderAddValue  func(v *VariantBuilder, value *Variant)
	VariantBuilderClear     func(v *VariantBuilder)
	VariantBuilderClose     func(v *VariantBuilder)
	VariantBuilderEnd       func(v *VariantBuilder) *Variant
	VariantBuilderInit      func(v *VariantBuilder, typ *VariantType)
	VariantBuilderOpen      func(v *VariantBuilder, typ *VariantType)
	VariantBuilderRef       func(v *VariantBuilder) *VariantBuilder
	VariantBuilderUnref     func(v *VariantBuilder)
)

func (v *VariantBuilder) Add(formatString string, va ...VArg) { VariantBuilderAdd(v, formatString, va) }
func (v *VariantBuilder) AddParsed(format string, va ...VArg) { VariantBuilderAddParsed(v, format, va) }
func (v *VariantBuilder) AddValue(value *Variant)             { VariantBuilderAddValue(v, value) }
func (v *VariantBuilder) Clear()                              { VariantBuilderClear(v) }
func (v *VariantBuilder) Close()                              { VariantBuilderClose(v) }
func (v *VariantBuilder) End() *Variant                       { return VariantBuilderEnd(v) }
func (v *VariantBuilder) Init(typ *VariantType)               { VariantBuilderInit(v, typ) }
func (v *VariantBuilder) Open(typ *VariantType)               { VariantBuilderOpen(v, typ) }
func (v *VariantBuilder) Ref() *VariantBuilder                { return VariantBuilderRef(v) }
func (v *VariantBuilder) Unref()                              { VariantBuilderUnref(v) }

type VariantIter struct {
	X [16]T.Gsize
}

var (
	VariantIterNew func(value *Variant) *VariantIter

	VariantIterCopy      func(v *VariantIter) *VariantIter
	VariantIterFree      func(v *VariantIter)
	VariantIterInit      func(v *VariantIter, value *Variant) T.Gsize
	VariantIterLoop      func(v *VariantIter, formatString string, va ...VArg) bool
	VariantIterNChildren func(v *VariantIter) T.Gsize
	VariantIterNext      func(v *VariantIter, formatString string, va ...VArg) bool
	VariantIterNextValue func(v *VariantIter) *Variant
)

func (v *VariantIter) Copy() *VariantIter          { return VariantIterCopy(v) }
func (v *VariantIter) Free()                       { VariantIterFree(v) }
func (v *VariantIter) Init(value *Variant) T.Gsize { return VariantIterInit(v, value) }
func (v *VariantIter) Loop(formatString string, va ...VArg) bool {
	return VariantIterLoop(v, formatString, va)
}
func (v *VariantIter) NChildren() T.Gsize { return VariantIterNChildren(v) }
func (v *VariantIter) Next(formatString string, va ...VArg) bool {
	return VariantIterNext(v, formatString, va)
}
func (v *VariantIter) NextValue() *Variant { return VariantIterNextValue(v) }

type VariantType T.VariantType // chicken/egg imports gobject/glib

var (
	VariantTypeNew      func(typeString string) *VariantType
	VariantTypeNewTuple func(items **VariantType, length int) *VariantType

	VariantTypeChecked       func(string) *VariantType
	VariantTypeEqual         func(type1 T.Gconstpointer, type2 T.Gconstpointer) bool
	VariantTypeHash          func(typ T.Gconstpointer) uint
	VariantTypeStringIsValid func(typeString string) bool
	VariantTypeStringScan    func(str, limit string, endptr **T.Gchar) bool

	VariantTypeCopy            func(v *VariantType) *VariantType
	VariantTypeDupString       func(v *VariantType) string
	VariantTypeElement         func(v *VariantType) *VariantType
	VariantTypeFirst           func(v *VariantType) *VariantType
	VariantTypeFree            func(v *VariantType)
	VariantTypeGetStringLength func(v *VariantType) T.Gsize
	VariantTypeIsArray         func(v *VariantType) bool
	VariantTypeIsBasic         func(v *VariantType) bool
	VariantTypeIsContainer     func(v *VariantType) bool
	VariantTypeIsDefinite      func(v *VariantType) bool
	VariantTypeIsDictEntry     func(v *VariantType) bool
	VariantTypeIsMaybe         func(v *VariantType) bool
	VariantTypeIsSubtypeOf     func(v *VariantType, superv *VariantType) bool
	VariantTypeIsTuple         func(v *VariantType) bool
	VariantTypeIsVariant       func(v *VariantType) bool
	VariantTypeKey             func(v *VariantType) *VariantType
	VariantTypeNewArray        func(element *VariantType) *VariantType
	VariantTypeNewDictEntry    func(key *VariantType, value *VariantType) *VariantType
	VariantTypeNewMaybe        func(element *VariantType) *VariantType
	VariantTypeNext            func(v *VariantType) *VariantType
	VariantTypeNItems          func(v *VariantType) T.Gsize
	VariantTypePeekString      func(v *VariantType) string
	VariantTypeValue           func(v *VariantType) *VariantType
)

func (v *VariantType) Copy() *VariantType    { return VariantTypeCopy(v) }
func (v *VariantType) DupString() string     { return VariantTypeDupString(v) }
func (v *VariantType) Element() *VariantType { return VariantTypeElement(v) }
func (v *VariantType) First() *VariantType   { return VariantTypeFirst(v) }
func (v *VariantType) Free()                 { VariantTypeFree(v) }
func (v *VariantType) StringLength() T.Gsize { return VariantTypeGetStringLength(v) }
func (v *VariantType) IsArray() bool         { return VariantTypeIsArray(v) }
func (v *VariantType) IsBasic() bool         { return VariantTypeIsBasic(v) }
func (v *VariantType) IsContainer() bool     { return VariantTypeIsContainer(v) }
func (v *VariantType) IsDefinite() bool      { return VariantTypeIsDefinite(v) }
func (v *VariantType) IsDictEntry() bool     { return VariantTypeIsDictEntry(v) }
func (v *VariantType) IsMaybe() bool         { return VariantTypeIsMaybe(v) }
func (v *VariantType) IsSubtypeOf(superv *VariantType) bool {
	return VariantTypeIsSubtypeOf(v, superv)
}
func (v *VariantType) IsTuple() bool          { return VariantTypeIsTuple(v) }
func (v *VariantType) IsVariant() bool        { return VariantTypeIsVariant(v) }
func (v *VariantType) Key() *VariantType      { return VariantTypeKey(v) }
func (e *VariantType) NewArray() *VariantType { return VariantTypeNewArray(e) }
func (k *VariantType) NewDictEntry(value *VariantType) *VariantType {
	return VariantTypeNewDictEntry(k, value)
}
func (e *VariantType) NewMaybe() *VariantType { return VariantTypeNewMaybe(e) }
func (v *VariantType) Next() *VariantType     { return VariantTypeNext(v) }
func (v *VariantType) NItems() T.Gsize        { return VariantTypeNItems(v) }
func (v *VariantType) PeekString() string     { return VariantTypePeekString(v) }
func (v *VariantType) Value() *VariantType    { return VariantTypeValue(v) }
