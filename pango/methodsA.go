// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
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
	ExtraAttrs  *T.GSList
}

type Attribute struct {
	Class      *AttrClass
	StartIndex uint
	EndIndex   uint
}

var (
	attributeInit    func(Attr *Attribute, class *AttrClass)
	attributeCopy    func(Attr *Attribute) *Attribute
	attributeDestroy func(Attr *Attribute)
	attributeEqual   func(Attr1 *Attribute, Attr2 *Attribute) T.Gboolean
)

func (a *Attribute) Init(class *AttrClass)             { attributeInit(a, class) }
func (a *Attribute) Copy() *Attribute                  { return attributeCopy(a) }
func (a *Attribute) Destroy()                          { attributeDestroy(a) }
func (a *Attribute) Equal(Attr2 *Attribute) T.Gboolean { return attributeEqual(a, Attr2) }

type AttrClass struct {
	Type    AttrType
	Copy    func(attr *Attribute) *Attribute
	Destroy func(attr *Attribute)
	Equal   func(attr1, attr2 *Attribute) T.Gboolean
}

var (
	AttrBackgroundNew         func(red, green, blue uint16) *Attribute
	AttrFallbackNew           func(enableFallback T.Gboolean) *Attribute
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
	AttrStrikethroughNew      func(strikethrough T.Gboolean) *Attribute
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
	attribute *Attribute, data T.Gpointer) T.Gboolean

type AttrList struct{}

var (
	AttrListGetType func() O.Type
	AttrListNew     func() *AttrList

	attrListChange       func(list *AttrList, Attr *Attribute)
	attrListCopy         func(list *AttrList) *AttrList
	attrListFilter       func(list *AttrList, f AttrFilterFunc, data T.Gpointer) *AttrList
	attrListGetIterator  func(list *AttrList) *AttrIterator
	attrListInsert       func(list *AttrList, Attr *Attribute)
	attrListInsertBefore func(list *AttrList, Attr *Attribute)
	attrListRef          func(list *AttrList) *AttrList
	attrListSplice       func(list *AttrList, other *AttrList, pos int, len int)
	attrListUnref        func(list *AttrList)
)

func (a *AttrList) Change(Attr *Attribute) { attrListChange(a, Attr) }
func (a *AttrList) Copy() *AttrList        { return attrListCopy(a) }
func (a *AttrList) Filter(f AttrFilterFunc, data T.Gpointer) *AttrList {
	return attrListFilter(a, f, data)
}
func (a *AttrList) GetIterator() *AttrIterator                { return attrListGetIterator(a) }
func (a *AttrList) Insert(Attr *Attribute)                    { attrListInsert(a, Attr) }
func (a *AttrList) InsertBefore(Attr *Attribute)              { attrListInsertBefore(a, Attr) }
func (a *AttrList) Ref() *AttrList                            { return attrListRef(a) }
func (a *AttrList) Splice(other *AttrList, pos int, leng int) { attrListSplice(a, other, pos, leng) }
func (a *AttrList) Unref()                                    { attrListUnref(a) }

type AttrIterator struct{}

var (
	attrIteratorCopy     func(iterator *AttrIterator) *AttrIterator
	attrIteratorDestroy  func(iterator *AttrIterator)
	attrIteratorGet      func(iterator *AttrIterator, t AttrType) *Attribute
	attrIteratorGetAttrs func(iterator *AttrIterator) *T.GSList
	attrIteratorGetFont  func(iterator *AttrIterator, desc *FontDescription, language **Language, extraAttrs **T.GSList)
	attrIteratorNext     func(iterator *AttrIterator) T.Gboolean
	attrIteratorRange    func(iterator *AttrIterator, start, end *int)
)

func (a *AttrIterator) Copy() *AttrIterator       { return attrIteratorCopy(a) }
func (a *AttrIterator) Destroy()                  { attrIteratorDestroy(a) }
func (a *AttrIterator) Get(t AttrType) *Attribute { return attrIteratorGet(a, t) }
func (a *AttrIterator) GetAttrs() *T.GSList       { return attrIteratorGetAttrs(a) }
func (a *AttrIterator) GetFont(desc *FontDescription, language **Language, extraAttrs **T.GSList) {
	attrIteratorGetFont(a, desc, language, extraAttrs)
}
func (a *AttrIterator) Next() T.Gboolean      { return attrIteratorNext(a) }
func (a *AttrIterator) Range(start, end *int) { attrIteratorRange(a, start, end) }

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
