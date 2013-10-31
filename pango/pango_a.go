// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Alignment Enum

const (
	ALIGN_LEFT Alignment = iota
	ALIGN_CENTER
	ALIGN_RIGHT
)

var AlignmentGetType func() O.Type

type Analysis struct {
	ShapeEngine *EngineShape
	LangEngine  *EngineLang
	Font        *Font
	Level       uint8
	Gravity     uint8
	Flags       uint8
	Script      uint8
	Language    *Language
	ExtraAttrs  *L.SList
}

type Attribute struct {
	Class      *AttrClass
	StartIndex uint
	EndIndex   uint
}

var (
	AttributeInit    func(Attr *Attribute, class *AttrClass)
	AttributeCopy    func(Attr *Attribute) *Attribute
	AttributeDestroy func(Attr *Attribute)
	AttributeEqual   func(Attr1 *Attribute, Attr2 *Attribute) bool
)

func (a *Attribute) Init(class *AttrClass)       { AttributeInit(a, class) }
func (a *Attribute) Copy() *Attribute            { return AttributeCopy(a) }
func (a *Attribute) Destroy()                    { AttributeDestroy(a) }
func (a *Attribute) Equal(Attr2 *Attribute) bool { return AttributeEqual(a, Attr2) }

type AttrClass struct {
	Type    AttrType
	Copy    func(attr *Attribute) *Attribute
	Destroy func(attr *Attribute)
	Equal   func(attr1, attr2 *Attribute) bool
}

var (
	AttrBackgroundNew         func(red, green, blue uint16) *Attribute
	AttrFallbackNew           func(enableFallback bool) *Attribute
	AttrFamilyNew             func(family string) *Attribute
	AttrFontDescNew           func(desc *FontDescription) *Attribute
	AttrForegroundNew         func(red, green, blue uint16) *Attribute
	AttrGravityHintNew        func(hint GravityHint) *Attribute
	AttrGravityNew            func(gravity Gravity) *Attribute
	AttrLanguageNew           func(language *Language) *Attribute
	AttrLetterSpacingNew      func(letterSpacing int) *Attribute
	AttrRiseNew               func(rise int) *Attribute
	AttrScaleNew              func(scaleFactor float64) *Attribute
	AttrShapeNew              func(inkRect *Rectangle, logicalRect *Rectangle) *Attribute
	AttrShapeNewWithData      func(inkRect *Rectangle, logicalRect *Rectangle, data T.Gpointer, copyFunc AttrDataCopyFunc, destroyFunc T.GDestroyNotify) *Attribute
	AttrSizeNew               func(size int) *Attribute
	AttrSizeNewAbsolute       func(size int) *Attribute
	AttrStretchNew            func(stretch Stretch) *Attribute
	AttrStrikethroughColorNew func(red, green, blue uint16) *Attribute
	AttrStrikethroughNew      func(strikethrough bool) *Attribute
	AttrStyleNew              func(style Style) *Attribute
	AttrTypeGetName           func(t AttrType) string
	AttrTypeRegister          func(name string) AttrType
	AttrUnderlineColorNew     func(red, green, blue uint16) *Attribute
	AttrUnderlineNew          func(underline Underline) *Attribute
	AttrVariantNew            func(variant Variant) *Attribute
	AttrWeightNew             func(weight Weight) *Attribute
)

type AttrDataCopyFunc func(data T.Gconstpointer) T.Gpointer

type AttrFilterFunc func(
	Attribute *Attribute, data T.Gpointer) bool

type AttrList struct{}

var (
	AttrListGetType func() O.Type
	AttrListNew     func() *AttrList

	AttrListChange       func(list *AttrList, Attr *Attribute)
	AttrListCopy         func(list *AttrList) *AttrList
	AttrListFilter       func(list *AttrList, f AttrFilterFunc, data T.Gpointer) *AttrList
	AttrListGetIterator  func(list *AttrList) *AttrIterator
	AttrListInsert       func(list *AttrList, Attr *Attribute)
	AttrListInsertBefore func(list *AttrList, Attr *Attribute)
	AttrListRef          func(list *AttrList) *AttrList
	AttrListSplice       func(list *AttrList, other *AttrList, pos int, len int)
	AttrListUnref        func(list *AttrList)
)

func (a *AttrList) Change(Attr *Attribute) { AttrListChange(a, Attr) }
func (a *AttrList) Copy() *AttrList        { return AttrListCopy(a) }
func (a *AttrList) Filter(f AttrFilterFunc, data T.Gpointer) *AttrList {
	return AttrListFilter(a, f, data)
}
func (a *AttrList) GetIterator() *AttrIterator                { return AttrListGetIterator(a) }
func (a *AttrList) Insert(Attr *Attribute)                    { AttrListInsert(a, Attr) }
func (a *AttrList) InsertBefore(Attr *Attribute)              { AttrListInsertBefore(a, Attr) }
func (a *AttrList) Ref() *AttrList                            { return AttrListRef(a) }
func (a *AttrList) Splice(other *AttrList, pos int, leng int) { AttrListSplice(a, other, pos, leng) }
func (a *AttrList) Unref()                                    { AttrListUnref(a) }

type AttrIterator struct{}

var (
	AttrIteratorCopy     func(iterator *AttrIterator) *AttrIterator
	AttrIteratorDestroy  func(iterator *AttrIterator)
	AttrIteratorGet      func(iterator *AttrIterator, t AttrType) *Attribute
	AttrIteratorGetAttrs func(iterator *AttrIterator) *L.SList
	AttrIteratorGetFont  func(iterator *AttrIterator, desc *FontDescription, language **Language, extraAttrs **L.SList)
	AttrIteratorNext     func(iterator *AttrIterator) bool
	AttrIteratorRange    func(iterator *AttrIterator, start, end *int)
)

func (a *AttrIterator) Copy() *AttrIterator       { return AttrIteratorCopy(a) }
func (a *AttrIterator) Destroy()                  { AttrIteratorDestroy(a) }
func (a *AttrIterator) Get(t AttrType) *Attribute { return AttrIteratorGet(a, t) }
func (a *AttrIterator) GetAttrs() *L.SList        { return AttrIteratorGetAttrs(a) }
func (a *AttrIterator) GetFont(desc *FontDescription, language **Language, extraAttrs **L.SList) {
	AttrIteratorGetFont(a, desc, language, extraAttrs)
}
func (a *AttrIterator) Next() bool            { return AttrIteratorNext(a) }
func (a *AttrIterator) Range(start, end *int) { AttrIteratorRange(a, start, end) }

type AttrShape struct {
	Attr        Attribute
	InkRect     Rectangle
	LogicalRect Rectangle
	Data        T.Gpointer
	CopyFunc    AttrDataCopyFunc
	DestroyFunc T.GDestroyNotify
}

type AttrType Enum

const (
	ATTR_INVALID AttrType = iota
	ATTR_LANGUAGE
	ATTR_FAMILY
	ATTR_STYLE
	ATTR_WEIGHT
	ATTR_VARIANT
	ATTR_STRETCH
	ATTR_SIZE
	ATTR_FONT_DESC
	ATTR_FOREGROUND
	ATTR_BACKGROUND
	ATTR_UNDERLINE
	ATTR_STRIKETHROUGH
	ATTR_RISE
	ATTR_SHAPE
	ATTR_SCALE
	ATTR_FALLBACK
	ATTR_LETTER_SPACING
	ATTR_UNDERLINE_COLOR
	ATTR_STRIKETHROUGH_COLOR
	ATTR_ABSOLUTE_SIZE
	ATTR_GRAVITY
	ATTR_GRAVITY_HINT
)

var AttrTypeGetType func() O.Type
