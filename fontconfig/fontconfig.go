// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package fontconfig provides API definitions for accessing
//libfontconfig-1.dll.
package fontconfig

import (
	"github.com/tHinqa/outside"
	FT "github.com/tHinqa/outside-gtk2/freetype2"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

const (
	FC_CHARSET_MAP_SIZE = 256 / 32
	FC_UTF8_MAX_LEN     = 6
)

type (
	Enum int

	Char32 uint // ANOMALLY

	FileCache struct{}
)

var (
	DirCacheUnlink func(dir string, config *Config) bool

	DirCacheValid func(cache_file string) bool

	FileIsDir func(file string) bool

	FileScan func(set *FontSet, dirs *StrSet,
		cache *FileCache, blanks *Blanks,
		file string, force bool) bool

	DirScan func(set *FontSet, dirs *StrSet,
		cache *FileCache, blanks *Blanks, dir string,
		force bool) bool

	DirSave func(
		set *FontSet, dirs *StrSet, dir string) bool

	DirCacheLoad func(dir string,
		config *Config, cache_file **uint8) *Cache

	DirCacheRead func(
		dir string, force bool, config *Config) *Cache

	DirCacheLoadFile func(
		cache_file string, file_stat *T.Stat) *Cache

	FreeTypeQuery func(file string,
		id int, blanks *Blanks, count *int) *Pattern

	InitLoadConfig func() *Config

	InitLoadConfigAndFonts func() *Config

	Init func() bool

	Fini func()

	GetVersion func() int

	InitReinitialize func() bool

	InitBringUptoDate func() bool

	GetLangs func() *StrSet

	FontList func(config *Config,
		p *Pattern, os *ObjectSet) *FontSet

	FontMatch func(config *Config,
		p *Pattern, result *Result) *Pattern

	FontRenderPrepare func(config *Config,
		pat *Pattern, font *Pattern) *Pattern

	FontSort func(config *Config, p *Pattern,
		trim bool, csp **CharSet, result *Result) *FontSet

	NameRegisterObjectTypes func(
		types *ObjectType, ntype int) bool

	NameUnregisterObjectTypes func(
		types *ObjectType, ntype int) bool

	NameGetObjectType func(object string) *ObjectType

	NameRegisterConstants func(
		consts *Constant, nconsts int) bool

	NameUnregisterConstants func(
		consts *Constant, nconsts int) bool

	NameGetConstant func(str string) *Constant

	NameConstant func(st string, result *int) bool

	NameParse func(name string) *Pattern

	StrCopy func(s string) string

	StrCopyFilename func(s string) string

	StrPlus func(s1, s2 string) string

	StrFree func(s string)

	StrDowncase func(s string) string

	StrCmpIgnoreCase func(s1, s2 string) int

	StrCmp func(s1, s2 string) int

	StrStrIgnoreCase func(s1, s2 string) string

	StrStr func(s1, s2 string) string

	Utf8ToUcs4 func(src_orig string, dst *Char32, leng int) int

	Utf8Len func(str *uint8, leng int, nchar, wchar *int) bool

	Ucs4ToUtf8 func(ucs4 Char32, dest [FC_UTF8_MAX_LEN]uint8) int

	Utf16ToUcs4 func(srcOrig string, endian Endian,
		dst *Char32, leng int) int

	Utf16Len func(str *uint8, endian Endian,
		leng int, nchar, wchar *int) bool

	StrDirname func(file string) string

	StrBasename func(file string) string

	FreeTypeCharIndex func(face FT.Face, ucs4 Char32) uint

	FreeTypeCharSetAndSpacing func(
		face FT.Face, blanks *Blanks, spacing *int) *CharSet

	FreeTypeCharSet func(
		face FT.Face, blanks *Blanks) *CharSet

	FreeTypeQueryFace func(face FT.Face,
		file string, id int, blanks *Blanks) *Pattern
)

type Atomic struct{}

var (
	AtomicCreate func(file string) *Atomic

	AtomicDeleteNew   func(a *Atomic)
	AtomicDestroy     func(a *Atomic)
	AtomicLock        func(a *Atomic) bool
	AtomicNewFile     func(a *Atomic) string
	AtomicOrigFile    func(a *Atomic) string
	AtomicReplaceOrig func(a *Atomic) bool
	AtomicUnlock      func(a *Atomic)
)

func (a *Atomic) DeleteNew()        { AtomicDeleteNew(a) }
func (a *Atomic) Destroy()          { AtomicDestroy(a) }
func (a *Atomic) Lock() bool        { return AtomicLock(a) }
func (a *Atomic) NewFile() string   { return AtomicNewFile(a) }
func (a *Atomic) OrigFile() string  { return AtomicOrigFile(a) }
func (a *Atomic) ReplaceOrig() bool { return AtomicReplaceOrig(a) }
func (a *Atomic) Unlock()           { AtomicUnlock(a) }

type Blanks struct{}

var (
	BlanksCreate func() *Blanks

	BlanksAdd      func(b *Blanks, ucs4 Char32) bool
	BlanksDestroy  func(b *Blanks)
	BlanksIsMember func(b *Blanks, ucs4 Char32) bool
)

func (b *Blanks) Add(ucs4 Char32) bool      { return BlanksAdd(b, ucs4) }
func (b *Blanks) Destroy()                  { BlanksDestroy(b) }
func (b *Blanks) IsMember(ucs4 Char32) bool { return BlanksIsMember(b, ucs4) }

type Cache struct{}

var (
	DirCacheUnload func(c *Cache)

	CacheCopySet   func(c *Cache) *FontSet
	CacheDir       func(c *Cache) string
	CacheNumFont   func(c *Cache) int
	CacheNumSubdir func(c *Cache) int
	CacheSubdir    func(c *Cache, i int) string
)

func (c *Cache) CopySet() *FontSet   { return CacheCopySet(c) }
func (c *Cache) Dir() string         { return CacheDir(c) }
func (c *Cache) NumFont() int        { return CacheNumFont(c) }
func (c *Cache) NumSubdir() int      { return CacheNumSubdir(c) }
func (c *Cache) Subdir(i int) string { return CacheSubdir(c, i) }

type CharSet struct{}

var (
	CharSetCreate func() *CharSet
	CharSetNew    func() *CharSet

	CharSetAddChar        func(c *CharSet, ucs4 Char32) bool
	CharSetCopy           func(c *CharSet) *CharSet
	CharSetCount          func(c *CharSet) Char32
	CharSetCoverage       func(c *CharSet, page Char32, result *Char32) Char32
	CharSetDestroy        func(c *CharSet)
	CharSetEqual          func(c *CharSet, c2 *CharSet) bool
	CharSetFirstPage      func(c *CharSet, Map [FC_CHARSET_MAP_SIZE]Char32, next *Char32) Char32
	CharSetHasChar        func(c *CharSet, ucs4 Char32) bool
	CharSetIntersect      func(c, c2 *CharSet) *CharSet
	CharSetIntersectCount func(c, c2 *CharSet) Char32
	CharSetIsSubset       func(c, c2 *CharSet) bool
	CharSetMerge          func(c, c2 *CharSet, changed *bool) bool
	CharSetNextPage       func(c *CharSet, Map [FC_CHARSET_MAP_SIZE]Char32, next *Char32) Char32
	CharSetSubtract       func(c, c2 *CharSet) *CharSet
	CharSetSubtractCount  func(c, c2 *CharSet) Char32
	CharSetUnion          func(c, c2 *CharSet) *CharSet
)

func (c *CharSet) AddChar(ucs4 Char32) bool { return CharSetAddChar(c, ucs4) }
func (c *CharSet) Copy() *CharSet           { return CharSetCopy(c) }
func (c *CharSet) Count() Char32            { return CharSetCount(c) }
func (c *CharSet) Coverage(page Char32, result *Char32) Char32 {
	return CharSetCoverage(c, page, result)
}
func (c *CharSet) Destroy()               { CharSetDestroy(c) }
func (c *CharSet) Equal(c2 *CharSet) bool { return CharSetEqual(c, c2) }
func (c *CharSet) FirstPage(Map [FC_CHARSET_MAP_SIZE]Char32, next *Char32) Char32 {
	return CharSetFirstPage(c, Map, next)
}
func (c *CharSet) HasChar(ucs4 Char32) bool              { return CharSetHasChar(c, ucs4) }
func (c *CharSet) Intersect(c2 *CharSet) *CharSet        { return CharSetIntersect(c, c2) }
func (c *CharSet) IntersectCount(c2 *CharSet) Char32     { return CharSetIntersectCount(c, c2) }
func (c *CharSet) IsSubset(c2 *CharSet) bool             { return CharSetIsSubset(c, c2) }
func (c *CharSet) Merge(c2 *CharSet, changed *bool) bool { return CharSetMerge(c, c2, changed) }
func (c *CharSet) NextPage(Map [FC_CHARSET_MAP_SIZE]Char32, next *Char32) Char32 {
	return CharSetNextPage(c, Map, next)
}
func (c *CharSet) Subtract(c2 *CharSet) *CharSet    { return CharSetSubtract(c, c2) }
func (c *CharSet) SubtractCount(c2 *CharSet) Char32 { return CharSetSubtractCount(c, c2) }
func (c *CharSet) Union(c2 *CharSet) *CharSet       { return CharSetUnion(c, c2) }

type Config struct{}

var (
	ConfigCreate     func() *Config
	ConfigEnableHome func(enable bool) bool
	ConfigFilename   func(url string) string
	ConfigGetCurrent func() *Config
	ConfigHome       func() string

	ConfigAppFontAddDir     func(c *Config, dir string) bool
	ConfigAppFontAddFile    func(c *Config, file string) bool
	ConfigAppFontClear      func(c *Config)
	ConfigBuildFonts        func(c *Config) bool
	ConfigDestroy           func(c *Config)
	ConfigGetBlanks         func(c *Config) *Blanks
	ConfigGetCache          func(c *Config) string
	ConfigGetCacheDirs      func(c *Config) *StrList
	ConfigGetConfigDirs     func(c *Config) *StrList
	ConfigGetConfigFiles    func(c *Config) *StrList
	ConfigGetFontDirs       func(c *Config) *StrList
	ConfigGetFonts          func(c *Config, set SetName) *FontSet
	ConfigGetRescanInterval func(c *Config) int
	ConfigParseAndLoad      func(c *Config, file string, complain bool) bool
	ConfigReference         func(c *Config) *Config
	ConfigSetCurrent        func(c *Config) bool
	ConfigSetRescanInterval func(c *Config, rescanInterval int) bool
	ConfigSubstitute        func(c *Config, p *Pattern, kind MatchKind) bool
	ConfigSubstituteWithPat func(c *Config, p, pPat *Pattern, kind MatchKind) bool
	ConfigUptoDate          func(c *Config) bool
)

func (c *Config) AppFontAddDir(dir string) bool   { return ConfigAppFontAddDir(c, dir) }
func (c *Config) AppFontAddFile(file string) bool { return ConfigAppFontAddFile(c, file) }
func (c *Config) AppFontClear()                   { ConfigAppFontClear(c) }
func (c *Config) Blanks() *Blanks                 { return ConfigGetBlanks(c) }
func (c *Config) BuildFonts() bool                { return ConfigBuildFonts(c) }
func (c *Config) Destroy()                        { ConfigDestroy(c) }
func (c *Config) Cache() string                   { return ConfigGetCache(c) }
func (c *Config) CacheDirs() *StrList             { return ConfigGetCacheDirs(c) }
func (c *Config) ConfigDirs() *StrList            { return ConfigGetConfigDirs(c) }
func (c *Config) ConfigFiles() *StrList           { return ConfigGetConfigFiles(c) }
func (c *Config) Current() bool                   { return ConfigSetCurrent(c) }
func (c *Config) FontDirs() *StrList              { return ConfigGetFontDirs(c) }
func (c *Config) Fonts(set SetName) *FontSet      { return ConfigGetFonts(c, set) }
func (c *Config) ParseAndLoad(file string, complain bool) bool {
	return ConfigParseAndLoad(c, file, complain)
}
func (c *Config) Reference() *Config  { return ConfigReference(c) }
func (c *Config) RescanInterval() int { return ConfigGetRescanInterval(c) }
func (c *Config) SetRescanInterval(rescanInterval int) bool {
	return ConfigSetRescanInterval(c, rescanInterval)
}
func (c *Config) Substitute(p *Pattern, kind MatchKind) bool { return ConfigSubstitute(c, p, kind) }
func (c *Config) SubstituteWithPat(p, pPat *Pattern, kind MatchKind) bool {
	return ConfigSubstituteWithPat(c, p, pPat, kind)
}
func (c *Config) UptoDate() bool { return ConfigUptoDate(c) }

type Constant struct {
	Name   *uint8
	Object *T.Char
	Value  int
}

type Endian Enum

const (
	EndianBig Endian = iota
	EndianLittle
)

type FontSet struct {
	Nfont int
	Sfont int
	Fonts **Pattern
}

var (
	FontSetCreate func() *FontSet

	FontSetList func(config *Config, sets **FontSet,
		nsets int, p *Pattern, os *ObjectSet) *FontSet

	FontSetMatch func(config *Config, sets **FontSet,
		nsets int, p *Pattern, result *Result) *Pattern

	FontSetSort func(config *Config, sets **FontSet,
		nsets int, p *Pattern, trim bool,
		csp **CharSet, result *Result) *FontSet

	FontSetAdd         func(f *FontSet, font *Pattern) bool
	FontSetDestroy     func(f *FontSet)
	FontSetPrint       func(f *FontSet)
	FontSetSortDestroy func(f *FontSet)
)

func (f *FontSet) Add(font *Pattern) bool { return FontSetAdd(f, font) }
func (f *FontSet) Destroy()               { FontSetDestroy(f) }
func (f *FontSet) Print()                 { FontSetPrint(f) }
func (f *FontSet) SortDestroy()           { FontSetSortDestroy(f) }

type LangResult Enum

const (
	LangEqual LangResult = 0
	LangDifferentCountry
	LangDifferentLang
	LangDifferentTerritory = LangDifferentCountry
)

type LangSet struct{}

var (
	LangSetCreate  func() *LangSet
	LangGetCharSet func(lang string) *CharSet

	LangSetAdd      func(l *LangSet, lang string) bool
	LangSetCompare  func(l, l2 *LangSet) LangResult
	LangSetContains func(l, l2 *LangSet) bool
	LangSetCopy     func(l *LangSet) *LangSet
	LangSetDestroy  func(l *LangSet)
	LangSetEqual    func(l, l2 *LangSet) bool
	LangSetGetLangs func(l *LangSet) *StrSet
	LangSetHash     func(l *LangSet) Char32
	LangSetHasLang  func(l *LangSet, lang string) LangResult
)

func (l *LangSet) Add(lang string) bool           { return LangSetAdd(l, lang) }
func (l *LangSet) Compare(l2 *LangSet) LangResult { return LangSetCompare(l, l2) }
func (l *LangSet) Contains(l2 *LangSet) bool      { return LangSetContains(l, l2) }
func (l *LangSet) Copy() *LangSet                 { return LangSetCopy(l) }
func (l *LangSet) Destroy()                       { LangSetDestroy(l) }
func (l *LangSet) Equal(l2 *LangSet) bool         { return LangSetEqual(l, l2) }
func (l *LangSet) GetLangs() *StrSet              { return LangSetGetLangs(l) }
func (l *LangSet) Hash() Char32                   { return LangSetHash(l) }
func (l *LangSet) HasLang(lang string) LangResult { return LangSetHasLang(l, lang) }

type MatchKind Enum

const (
	MatchPattern MatchKind = iota
	MatchFont
	MatchScan
)

type Matrix struct {
	XX, XY, YX, YY float64
}

var (
	MatrixMultiply func(result, a, b *Matrix)

	MatrixCopy   func(m *Matrix) *Matrix
	MatrixEqual  func(m1, m2 *Matrix) bool
	MatrixRotate func(m *Matrix, c, s float64)
	MatrixScale  func(m *Matrix, sx, sy float64)
	MatrixShear  func(m *Matrix, sh, sv float64)
)

func (m *Matrix) Copy() *Matrix         { return MatrixCopy(m) }
func (m *Matrix) Equal(m2 *Matrix) bool { return MatrixEqual(m, m2) }
func (m *Matrix) Rotate(c, s float64)   { MatrixRotate(m, c, s) }
func (m *Matrix) Scale(sx, sy float64)  { MatrixScale(m, sx, sy) }
func (m *Matrix) Shear(sh, sv float64)  { MatrixShear(m, sh, sv) }

type ObjectSet struct {
	Nobject int
	Sobject int
	Objects **T.Char
}

var (
	ObjectSetBuild   func(first string, v ...VArg) *ObjectSet
	ObjectSetCreate  func() *ObjectSet
	ObjectSetVaBuild func(first string, va T.VaList) *ObjectSet

	ObjectSetAdd     func(os *ObjectSet, object string) bool
	ObjectSetDestroy func(os *ObjectSet)
)

func (o *ObjectSet) Add(object string) bool { return ObjectSetAdd(o, object) }
func (o *ObjectSet) Destroy()               { ObjectSetDestroy(o) }

type ObjectType struct {
	Object *T.Char
	Type   Type
}

type Pattern struct{}

var (
	PatternCreate func() *Pattern

	DefaultSubstitute func(p *Pattern)
	NameUnparse       func(p *Pattern) string

	PatternAdd         func(p *Pattern, object string, value Value, append bool) bool
	PatternAddBool     func(p *Pattern, object string, b bool) bool
	PatternAddCharSet  func(p *Pattern, object string, c *CharSet) bool
	PatternAddDouble   func(p *Pattern, object string, d float64) bool
	PatternAddFTFace   func(p *Pattern, object string, f FT.Face) bool
	PatternAddInteger  func(p *Pattern, object string, i int) bool
	PatternAddLangSet  func(p *Pattern, object string, ls *LangSet) bool
	PatternAddMatrix   func(p *Pattern, object string, s *Matrix) bool
	PatternAddString   func(p *Pattern, object string, s string) bool
	PatternAddWeak     func(p *Pattern, object string, value Value, append bool) bool
	PatternBuild       func(p *Pattern, v ...VArg) *Pattern
	PatternDel         func(p *Pattern, object string) bool
	PatternDestroy     func(p *Pattern)
	PatternDuplicate   func(p *Pattern) *Pattern
	PatternEqual       func(p *Pattern, pb *Pattern) bool
	PatternEqualSubset func(p *Pattern, pb *Pattern, os *ObjectSet) bool
	PatternFilter      func(p *Pattern, os *ObjectSet) *Pattern
	PatternFormat      func(p *Pattern, format string) string
	PatternGet         func(p *Pattern, object string, id int, v *Value) Result
	PatternGetBool     func(p *Pattern, object string, n int, b *bool) Result
	PatternGetCharSet  func(p *Pattern, object string, n int, c **CharSet) Result
	PatternGetDouble   func(p *Pattern, object string, n int, d *float64) Result
	PatternGetFTFace   func(p *Pattern, object string, n int, f *FT.Face) Result
	PatternGetInteger  func(p *Pattern, object string, n int, i *int) Result
	PatternGetLangSet  func(p *Pattern, object string, n int, ls **LangSet) Result
	PatternGetMatrix   func(p *Pattern, object string, n int, s **Matrix) Result
	PatternGetString   func(p *Pattern, object string, n int, s *string) Result
	PatternHash        func(p *Pattern) Char32
	PatternPrint       func(p *Pattern)
	PatternReference   func(p *Pattern)
	PatternRemove      func(p *Pattern, object string, id int) bool
	PatternVaBuild     func(p *Pattern, va T.VaList) *Pattern
)

func (p *Pattern) Add(object string, value Value, append bool) bool {
	return PatternAdd(p, object, value, append)
}
func (p *Pattern) Addbool(object string, b bool) bool         { return PatternAddBool(p, object, b) }
func (p *Pattern) AddCharSet(object string, c *CharSet) bool  { return PatternAddCharSet(p, object, c) }
func (p *Pattern) AddDouble(object string, d float64) bool    { return PatternAddDouble(p, object, d) }
func (p *Pattern) AddFTFace(object string, f FT.Face) bool    { return PatternAddFTFace(p, object, f) }
func (p *Pattern) AddInteger(object string, i int) bool       { return PatternAddInteger(p, object, i) }
func (p *Pattern) AddLangSet(object string, ls *LangSet) bool { return PatternAddLangSet(p, object, ls) }
func (p *Pattern) AddMatrix(object string, s *Matrix) bool    { return PatternAddMatrix(p, object, s) }
func (p *Pattern) AddString(object string, s string) bool     { return PatternAddString(p, object, s) }
func (p *Pattern) AddWeak(object string, value Value, append bool) bool {
	return PatternAddWeak(p, object, value, append)
}
func (p *Pattern) Build(v ...VArg) *Pattern                    { return PatternBuild(p, v) }
func (p *Pattern) Del(object string) bool                      { return PatternDel(p, object) }
func (p *Pattern) Destroy()                                    { PatternDestroy(p) }
func (p *Pattern) Duplicate() *Pattern                         { return PatternDuplicate(p) }
func (p *Pattern) Equal(pb *Pattern) bool                      { return PatternEqual(p, pb) }
func (p *Pattern) EqualSubset(pb *Pattern, os *ObjectSet) bool { return PatternEqualSubset(p, pb, os) }
func (p *Pattern) Filter(os *ObjectSet) *Pattern               { return PatternFilter(p, os) }
func (p *Pattern) Format(format string) string                 { return PatternFormat(p, format) }
func (p *Pattern) Get(object string, id int, v *Value) Result  { return PatternGet(p, object, id, v) }
func (p *Pattern) GetBool(object string, n int, b *bool) Result {
	return PatternGetBool(p, object, n, b)
}
func (p *Pattern) GetCharSet(object string, n int, c **CharSet) Result {
	return PatternGetCharSet(p, object, n, c)
}
func (p *Pattern) GetDouble(object string, n int, d *float64) Result {
	return PatternGetDouble(p, object, n, d)
}
func (p *Pattern) GetFTFace(object string, n int, f *FT.Face) Result {
	return PatternGetFTFace(p, object, n, f)
}
func (p *Pattern) GetInteger(object string, n int, i *int) Result {
	return PatternGetInteger(p, object, n, i)
}
func (p *Pattern) GetLangSet(object string, n int, ls **LangSet) Result {
	return PatternGetLangSet(p, object, n, ls)
}
func (p *Pattern) GetMatrix(object string, n int, s **Matrix) Result {
	return PatternGetMatrix(p, object, n, s)
}
func (p *Pattern) GetString(object string, n int, s *string) Result {
	return PatternGetString(p, object, n, s)
}
func (p *Pattern) Hash() Char32                      { return PatternHash(p) }
func (p *Pattern) Print()                            { PatternPrint(p) }
func (p *Pattern) Reference()                        { PatternReference(p) }
func (p *Pattern) Remove(object string, id int) bool { return PatternRemove(p, object, id) }
func (p *Pattern) VaBuild(va T.VaList) *Pattern      { return PatternVaBuild(p, va) }

type Result Enum

const (
	ResultMatch Result = iota
	ResultNoMatch
	ResultTypeMismatch
	ResultNoId
	ResultOutOfMemory
)

type SetName Enum

const (
	SetSystem SetName = iota
	SetApplication
)

type StrList struct{}

var (
	StrListCreate func(set *StrSet) *StrList

	StrListDone func(s *StrList)
	StrListNext func(s *StrList) string
)

func (s *StrList) Done()        { StrListDone(s) }
func (s *StrList) Next() string { return StrListNext(s) }

type StrSet struct{}

var (
	StrSetCreate func() *StrSet

	StrSetAdd         func(s *StrSet, str string) bool
	StrSetAddFilename func(s *StrSet, str string) bool
	StrSetDel         func(s *StrSet, str string) bool
	StrSetDestroy     func(s *StrSet)
	StrSetEqual       func(s, s2 *StrSet) bool
	StrSetMember      func(s *StrSet, str string) bool
)

func (s *StrSet) Add(str string) bool         { return StrSetAdd(s, str) }
func (s *StrSet) AddFilename(str string) bool { return StrSetAddFilename(s, str) }
func (s *StrSet) Del(str string) bool         { return StrSetDel(s, str) }
func (s *StrSet) Destroy()                    { StrSetDestroy(s) }
func (s *StrSet) Equal(s2 *StrSet) bool       { return StrSetEqual(s, s2) }
func (s *StrSet) Member(str string) bool      { return StrSetMember(s, str) }

type Type Enum

const (
	TypeVoid Type = iota
	TypeInteger
	TypeDouble
	TypeString
	Typebool
	TypeMatrix
	TypeCharSet
	TypeFTFace
	TypeLangSet
)

type Value struct {
	Type Type
	// Union
	S *uint8
	// i int
	// b bool
	// d double
	// m *Matrix
	// c *CharSet
	// f *void
	// l *LangSet
}

var (
	ValueDestroy func(v Value)
	ValueEqual   func(va, vb Value) bool
	ValuePrint   func(v Value)
	ValueSave    func(v Value) Value
)

func (v Value) Destroy()            { ValueDestroy(v) }
func (v Value) Equal(v2 Value) bool { return ValueEqual(v, v2) }
func (v Value) Print()              { ValuePrint(v) }
func (v Value) Save() Value         { return ValueSave(v) }

var dll = "libfontconfig-1.dll"

var apiList = outside.Apis{
	{"FcAtomicCreate", &AtomicCreate},
	{"FcAtomicDeleteNew", &AtomicDeleteNew},
	{"FcAtomicDestroy", &AtomicDestroy},
	{"FcAtomicLock", &AtomicLock},
	{"FcAtomicNewFile", &AtomicNewFile},
	{"FcAtomicOrigFile", &AtomicOrigFile},
	{"FcAtomicReplaceOrig", &AtomicReplaceOrig},
	{"FcAtomicUnlock", &AtomicUnlock},
	{"FcBlanksAdd", &BlanksAdd},
	{"FcBlanksCreate", &BlanksCreate},
	{"FcBlanksDestroy", &BlanksDestroy},
	{"FcBlanksIsMember", &BlanksIsMember},
	{"FcCacheCopySet", &CacheCopySet},
	{"FcCacheDir", &CacheDir},
	{"FcCacheNumFont", &CacheNumFont},
	{"FcCacheNumSubdir", &CacheNumSubdir},
	{"FcCacheSubdir", &CacheSubdir},
	{"FcCharSetAddChar", &CharSetAddChar},
	{"FcCharSetCopy", &CharSetCopy},
	{"FcCharSetCount", &CharSetCount},
	{"FcCharSetCoverage", &CharSetCoverage},
	{"FcCharSetCreate", &CharSetCreate},
	{"FcCharSetDestroy", &CharSetDestroy},
	{"FcCharSetEqual", &CharSetEqual},
	{"FcCharSetFirstPage", &CharSetFirstPage},
	{"FcCharSetHasChar", &CharSetHasChar},
	{"FcCharSetIntersect", &CharSetIntersect},
	{"FcCharSetIntersectCount", &CharSetIntersectCount},
	{"FcCharSetIsSubset", &CharSetIsSubset},
	{"FcCharSetMerge", &CharSetMerge},
	{"FcCharSetNew", &CharSetNew},
	{"FcCharSetNextPage", &CharSetNextPage},
	{"FcCharSetSubtract", &CharSetSubtract},
	{"FcCharSetSubtractCount", &CharSetSubtractCount},
	{"FcCharSetUnion", &CharSetUnion},
	{"FcConfigAppFontAddDir", &ConfigAppFontAddDir},
	{"FcConfigAppFontAddFile", &ConfigAppFontAddFile},
	{"FcConfigAppFontClear", &ConfigAppFontClear},
	{"FcConfigBuildFonts", &ConfigBuildFonts},
	{"FcConfigCreate", &ConfigCreate},
	{"FcConfigDestroy", &ConfigDestroy},
	{"FcConfigEnableHome", &ConfigEnableHome},
	{"FcConfigFilename", &ConfigFilename},
	{"FcConfigGetBlanks", &ConfigGetBlanks},
	{"FcConfigGetCache", &ConfigGetCache},
	{"FcConfigGetCacheDirs", &ConfigGetCacheDirs},
	{"FcConfigGetConfigDirs", &ConfigGetConfigDirs},
	{"FcConfigGetConfigFiles", &ConfigGetConfigFiles},
	{"FcConfigGetCurrent", &ConfigGetCurrent},
	{"FcConfigGetFontDirs", &ConfigGetFontDirs},
	{"FcConfigGetFonts", &ConfigGetFonts},
	// Deprecated/Undocumented {"FcConfigGetRescanInterval", &ConfigGetRescanInterval},
	{"FcConfigHome", &ConfigHome},
	{"FcConfigParseAndLoad", &ConfigParseAndLoad},
	{"FcConfigReference", &ConfigReference},
	{"FcConfigSetCurrent", &ConfigSetCurrent},
	// Deprecated/Undocumented {"FcConfigSetRescanInterval", &ConfigSetRescanInterval},
	{"FcConfigSubstitute", &ConfigSubstitute},
	{"FcConfigSubstituteWithPat", &ConfigSubstituteWithPat},
	{"FcConfigUptoDate", &ConfigUptoDate},
	{"FcDefaultSubstitute", &DefaultSubstitute},
	{"FcDirCacheLoad", &DirCacheLoad},
	{"FcDirCacheLoadFile", &DirCacheLoadFile},
	{"FcDirCacheRead", &DirCacheRead},
	{"FcDirCacheUnlink", &DirCacheUnlink},
	{"FcDirCacheUnload", &DirCacheUnload},
	{"FcDirCacheValid", &DirCacheValid},
	{"FcDirSave", &DirSave},
	{"FcDirScan", &DirScan},
	{"FcFileIsDir", &FileIsDir},
	{"FcFileScan", &FileScan},
	{"FcFini", &Fini},
	{"FcFontList", &FontList},
	{"FcFontMatch", &FontMatch},
	{"FcFontRenderPrepare", &FontRenderPrepare},
	{"FcFontSetAdd", &FontSetAdd},
	{"FcFontSetCreate", &FontSetCreate},
	{"FcFontSetDestroy", &FontSetDestroy},
	{"FcFontSetList", &FontSetList},
	{"FcFontSetMatch", &FontSetMatch},
	{"FcFontSetPrint", &FontSetPrint},
	{"FcFontSetSort", &FontSetSort},
	{"FcFontSetSortDestroy", &FontSetSortDestroy},
	{"FcFontSort", &FontSort},
	{"FcFreeTypeCharIndex", &FreeTypeCharIndex},
	{"FcFreeTypeCharSet", &FreeTypeCharSet},
	{"FcFreeTypeCharSetAndSpacing", &FreeTypeCharSetAndSpacing},
	{"FcFreeTypeQuery", &FreeTypeQuery},
	{"FcFreeTypeQueryFace", &FreeTypeQueryFace},
	{"FcGetLangs", &GetLangs},
	{"FcGetVersion", &GetVersion},
	{"FcInit", &Init},
	{"FcInitBringUptoDate", &InitBringUptoDate},
	{"FcInitLoadConfig", &InitLoadConfig},
	{"FcInitLoadConfigAndFonts", &InitLoadConfigAndFonts},
	{"FcInitReinitialize", &InitReinitialize},
	{"FcLangGetCharSet", &LangGetCharSet},
	{"FcLangSetAdd", &LangSetAdd},
	{"FcLangSetCompare", &LangSetCompare},
	{"FcLangSetContains", &LangSetContains},
	{"FcLangSetCopy", &LangSetCopy},
	{"FcLangSetCreate", &LangSetCreate},
	{"FcLangSetDestroy", &LangSetDestroy},
	{"FcLangSetEqual", &LangSetEqual},
	{"FcLangSetGetLangs", &LangSetGetLangs},
	{"FcLangSetHasLang", &LangSetHasLang},
	{"FcLangSetHash", &LangSetHash},
	{"FcMatrixCopy", &MatrixCopy},
	{"FcMatrixEqual", &MatrixEqual},
	{"FcMatrixMultiply", &MatrixMultiply},
	{"FcMatrixRotate", &MatrixRotate},
	{"FcMatrixScale", &MatrixScale},
	{"FcMatrixShear", &MatrixShear},
	{"FcNameConstant", &NameConstant},
	{"FcNameGetConstant", &NameGetConstant},
	{"FcNameGetObjectType", &NameGetObjectType},
	{"FcNameParse", &NameParse},
	{"FcNameRegisterConstants", &NameRegisterConstants},
	{"FcNameRegisterObjectTypes", &NameRegisterObjectTypes},
	{"FcNameUnparse", &NameUnparse},
	{"FcNameUnregisterConstants", &NameUnregisterConstants},
	{"FcNameUnregisterObjectTypes", &NameUnregisterObjectTypes},
	{"FcObjectSetAdd", &ObjectSetAdd},
	{"FcObjectSetBuild", &ObjectSetBuild},
	{"FcObjectSetCreate", &ObjectSetCreate},
	{"FcObjectSetDestroy", &ObjectSetDestroy},
	{"FcObjectSetVaBuild", &ObjectSetVaBuild},
	{"FcPatternAdd", &PatternAdd},
	{"FcPatternAddBool", &PatternAddBool},
	{"FcPatternAddCharSet", &PatternAddCharSet},
	{"FcPatternAddDouble", &PatternAddDouble},
	{"FcPatternAddFTFace", &PatternAddFTFace},
	{"FcPatternAddInteger", &PatternAddInteger},
	{"FcPatternAddLangSet", &PatternAddLangSet},
	{"FcPatternAddMatrix", &PatternAddMatrix},
	{"FcPatternAddString", &PatternAddString},
	{"FcPatternAddWeak", &PatternAddWeak},
	{"FcPatternBuild", &PatternBuild},
	{"FcPatternCreate", &PatternCreate},
	{"FcPatternDel", &PatternDel},
	{"FcPatternDestroy", &PatternDestroy},
	{"FcPatternDuplicate", &PatternDuplicate},
	{"FcPatternEqual", &PatternEqual},
	{"FcPatternEqualSubset", &PatternEqualSubset},
	{"FcPatternFilter", &PatternFilter},
	{"FcPatternFormat", &PatternFormat},
	{"FcPatternGet", &PatternGet},
	{"FcPatternGetBool", &PatternGetBool},
	{"FcPatternGetCharSet", &PatternGetCharSet},
	{"FcPatternGetDouble", &PatternGetDouble},
	{"FcPatternGetFTFace", &PatternGetFTFace},
	{"FcPatternGetInteger", &PatternGetInteger},
	{"FcPatternGetLangSet", &PatternGetLangSet},
	{"FcPatternGetMatrix", &PatternGetMatrix},
	{"FcPatternGetString", &PatternGetString},
	{"FcPatternHash", &PatternHash},
	{"FcPatternPrint", &PatternPrint},
	{"FcPatternReference", &PatternReference},
	{"FcPatternRemove", &PatternRemove},
	{"FcPatternVaBuild", &PatternVaBuild},
	{"FcStrBasename", &StrBasename},
	{"FcStrCmp", &StrCmp},
	{"FcStrCmpIgnoreCase", &StrCmpIgnoreCase},
	{"FcStrCopy", &StrCopy},
	{"FcStrCopyFilename", &StrCopyFilename},
	{"FcStrDirname", &StrDirname},
	{"FcStrDowncase", &StrDowncase},
	{"FcStrFree", &StrFree},
	{"FcStrListCreate", &StrListCreate},
	{"FcStrListDone", &StrListDone},
	{"FcStrListNext", &StrListNext},
	{"FcStrPlus", &StrPlus},
	{"FcStrSetAdd", &StrSetAdd},
	{"FcStrSetAddFilename", &StrSetAddFilename},
	{"FcStrSetCreate", &StrSetCreate},
	{"FcStrSetDel", &StrSetDel},
	{"FcStrSetDestroy", &StrSetDestroy},
	{"FcStrSetEqual", &StrSetEqual},
	{"FcStrSetMember", &StrSetMember},
	{"FcStrStr", &StrStr},
	{"FcStrStrIgnoreCase", &StrStrIgnoreCase},
	{"FcUcs4ToUtf8", &Ucs4ToUtf8},
	{"FcUtf16Len", &Utf16Len},
	{"FcUtf16ToUcs4", &Utf16ToUcs4},
	{"FcUtf8Len", &Utf8Len},
	{"FcUtf8ToUcs4", &Utf8ToUcs4},
	{"FcValueDestroy", &ValueDestroy},
	{"FcValueEqual", &ValueEqual},
	{"FcValuePrint", &ValuePrint},
	{"FcValueSave", &ValueSave},
}
