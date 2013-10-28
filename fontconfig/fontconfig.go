// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package fontconfig provides API definitions for accessing
//libfontconfig-1.dll.
package fontconfig

import (
	"github.com/tHinqa/outside"
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

	FreeTypeCharIndex func(face T.FTFace, ucs4 Char32) uint

	FreeTypeCharSetAndSpacing func(
		face T.FTFace, blanks *Blanks, spacing *int) *CharSet

	FreeTypeCharSet func(
		face T.FTFace, blanks *Blanks) *CharSet

	FreeTypeQueryFace func(face T.FTFace,
		file string, id int, blanks *Blanks) *Pattern
)

type Atomic struct{}

var (
	AtomicCreate func(file string) *Atomic

	atomicDeleteNew   func(a *Atomic)
	atomicDestroy     func(a *Atomic)
	atomicLock        func(a *Atomic) bool
	atomicNewFile     func(a *Atomic) string
	atomicOrigFile    func(a *Atomic) string
	atomicReplaceOrig func(a *Atomic) bool
	atomicUnlock      func(a *Atomic)
)

func (a *Atomic) DeleteNew()        { atomicDeleteNew(a) }
func (a *Atomic) Destroy()          { atomicDestroy(a) }
func (a *Atomic) Lock() bool        { return atomicLock(a) }
func (a *Atomic) NewFile() string   { return atomicNewFile(a) }
func (a *Atomic) OrigFile() string  { return atomicOrigFile(a) }
func (a *Atomic) ReplaceOrig() bool { return atomicReplaceOrig(a) }
func (a *Atomic) Unlock()           { atomicUnlock(a) }

type Blanks struct{}

var (
	BlanksCreate func() *Blanks

	blanksAdd      func(b *Blanks, ucs4 Char32) bool
	blanksDestroy  func(b *Blanks)
	blanksIsMember func(b *Blanks, ucs4 Char32) bool
)

func (b *Blanks) Add(ucs4 Char32) bool      { return blanksAdd(b, ucs4) }
func (b *Blanks) Destroy()                  { blanksDestroy(b) }
func (b *Blanks) IsMember(ucs4 Char32) bool { return blanksIsMember(b, ucs4) }

type Cache struct{}

var (
	DirCacheUnload func(c *Cache)

	cacheCopySet   func(c *Cache) *FontSet
	cacheDir       func(c *Cache) string
	cacheNumFont   func(c *Cache) int
	cacheNumSubdir func(c *Cache) int
	cacheSubdir    func(c *Cache, i int) string
)

func (c *Cache) CopySet() *FontSet   { return cacheCopySet(c) }
func (c *Cache) Dir() string         { return cacheDir(c) }
func (c *Cache) NumFont() int        { return cacheNumFont(c) }
func (c *Cache) NumSubdir() int      { return cacheNumSubdir(c) }
func (c *Cache) Subdir(i int) string { return cacheSubdir(c, i) }

type CharSet struct{}

var (
	CharSetCreate func() *CharSet
	CharSetNew    func() *CharSet

	charSetAddChar        func(c *CharSet, ucs4 Char32) bool
	charSetCopy           func(c *CharSet) *CharSet
	charSetCount          func(c *CharSet) Char32
	charSetCoverage       func(c *CharSet, page Char32, result *Char32) Char32
	charSetDestroy        func(c *CharSet)
	charSetEqual          func(c *CharSet, c2 *CharSet) bool
	charSetFirstPage      func(c *CharSet, Map [FC_CHARSET_MAP_SIZE]Char32, next *Char32) Char32
	charSetHasChar        func(c *CharSet, ucs4 Char32) bool
	charSetIntersect      func(c, c2 *CharSet) *CharSet
	charSetIntersectCount func(c, c2 *CharSet) Char32
	charSetIsSubset       func(c, c2 *CharSet) bool
	charSetMerge          func(c, c2 *CharSet, changed *bool) bool
	charSetNextPage       func(c *CharSet, Map [FC_CHARSET_MAP_SIZE]Char32, next *Char32) Char32
	charSetSubtract       func(c, c2 *CharSet) *CharSet
	charSetSubtractCount  func(c, c2 *CharSet) Char32
	charSetUnion          func(c, c2 *CharSet) *CharSet
)

func (c *CharSet) AddChar(ucs4 Char32) bool { return charSetAddChar(c, ucs4) }
func (c *CharSet) Copy() *CharSet           { return charSetCopy(c) }
func (c *CharSet) Count() Char32            { return charSetCount(c) }
func (c *CharSet) Coverage(page Char32, result *Char32) Char32 {
	return charSetCoverage(c, page, result)
}
func (c *CharSet) Destroy()               { charSetDestroy(c) }
func (c *CharSet) Equal(c2 *CharSet) bool { return charSetEqual(c, c2) }
func (c *CharSet) FirstPage(Map [FC_CHARSET_MAP_SIZE]Char32, next *Char32) Char32 {
	return charSetFirstPage(c, Map, next)
}
func (c *CharSet) HasChar(ucs4 Char32) bool              { return charSetHasChar(c, ucs4) }
func (c *CharSet) Intersect(c2 *CharSet) *CharSet        { return charSetIntersect(c, c2) }
func (c *CharSet) IntersectCount(c2 *CharSet) Char32     { return charSetIntersectCount(c, c2) }
func (c *CharSet) IsSubset(c2 *CharSet) bool             { return charSetIsSubset(c, c2) }
func (c *CharSet) Merge(c2 *CharSet, changed *bool) bool { return charSetMerge(c, c2, changed) }
func (c *CharSet) NextPage(Map [FC_CHARSET_MAP_SIZE]Char32, next *Char32) Char32 {
	return charSetNextPage(c, Map, next)
}
func (c *CharSet) Subtract(c2 *CharSet) *CharSet    { return charSetSubtract(c, c2) }
func (c *CharSet) SubtractCount(c2 *CharSet) Char32 { return charSetSubtractCount(c, c2) }
func (c *CharSet) Union(c2 *CharSet) *CharSet       { return charSetUnion(c, c2) }

type Config struct{}

var (
	ConfigCreate     func() *Config
	ConfigEnableHome func(enable bool) bool
	ConfigFilename   func(url string) string
	ConfigGetCurrent func() *Config
	ConfigHome       func() string

	configAppFontAddDir     func(c *Config, dir string) bool
	configAppFontAddFile    func(c *Config, file string) bool
	configAppFontClear      func(c *Config)
	configBuildFonts        func(c *Config) bool
	configDestroy           func(c *Config)
	configGetBlanks         func(c *Config) *Blanks
	configGetCache          func(c *Config) string
	configGetCacheDirs      func(c *Config) *StrList
	configGetConfigDirs     func(c *Config) *StrList
	configGetConfigFiles    func(c *Config) *StrList
	configGetFontDirs       func(c *Config) *StrList
	configGetFonts          func(c *Config, set SetName) *FontSet
	configGetRescanInterval func(c *Config) int
	configParseAndLoad      func(c *Config, file string, complain bool) bool
	configReference         func(c *Config) *Config
	configSetCurrent        func(c *Config) bool
	configSetRescanInterval func(c *Config, rescanInterval int) bool
	configSubstitute        func(c *Config, p *Pattern, kind MatchKind) bool
	configSubstituteWithPat func(c *Config, p, pPat *Pattern, kind MatchKind) bool
	configUptoDate          func(c *Config) bool
)

func (c *Config) AppFontAddDir(dir string) bool   { return configAppFontAddDir(c, dir) }
func (c *Config) AppFontAddFile(file string) bool { return configAppFontAddFile(c, file) }
func (c *Config) AppFontClear()                   { configAppFontClear(c) }
func (c *Config) Blanks() *Blanks                 { return configGetBlanks(c) }
func (c *Config) BuildFonts() bool                { return configBuildFonts(c) }
func (c *Config) Destroy()                        { configDestroy(c) }
func (c *Config) Cache() string                   { return configGetCache(c) }
func (c *Config) CacheDirs() *StrList             { return configGetCacheDirs(c) }
func (c *Config) ConfigDirs() *StrList            { return configGetConfigDirs(c) }
func (c *Config) ConfigFiles() *StrList           { return configGetConfigFiles(c) }
func (c *Config) Current() bool                   { return configSetCurrent(c) }
func (c *Config) FontDirs() *StrList              { return configGetFontDirs(c) }
func (c *Config) Fonts(set SetName) *FontSet      { return configGetFonts(c, set) }
func (c *Config) ParseAndLoad(file string, complain bool) bool {
	return configParseAndLoad(c, file, complain)
}
func (c *Config) Reference() *Config  { return configReference(c) }
func (c *Config) RescanInterval() int { return configGetRescanInterval(c) }
func (c *Config) SetRescanInterval(rescanInterval int) bool {
	return configSetRescanInterval(c, rescanInterval)
}
func (c *Config) Substitute(p *Pattern, kind MatchKind) bool { return configSubstitute(c, p, kind) }
func (c *Config) SubstituteWithPat(p, pPat *Pattern, kind MatchKind) bool {
	return configSubstituteWithPat(c, p, pPat, kind)
}
func (c *Config) UptoDate() bool { return configUptoDate(c) }

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

	fontSetAdd         func(f *FontSet, font *Pattern) bool
	fontSetDestroy     func(f *FontSet)
	fontSetPrint       func(f *FontSet)
	fontSetSortDestroy func(f *FontSet)
)

func (f *FontSet) Add(font *Pattern) bool { return fontSetAdd(f, font) }
func (f *FontSet) Destroy()               { fontSetDestroy(f) }
func (f *FontSet) Print()                 { fontSetPrint(f) }
func (f *FontSet) SortDestroy()           { fontSetSortDestroy(f) }

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

	langSetAdd      func(l *LangSet, lang string) bool
	langSetCompare  func(l, l2 *LangSet) LangResult
	langSetContains func(l, l2 *LangSet) bool
	langSetCopy     func(l *LangSet) *LangSet
	langSetDestroy  func(l *LangSet)
	langSetEqual    func(l, l2 *LangSet) bool
	langSetGetLangs func(l *LangSet) *StrSet
	langSetHash     func(l *LangSet) Char32
	langSetHasLang  func(l *LangSet, lang string) LangResult
)

func (l *LangSet) Add(lang string) bool           { return langSetAdd(l, lang) }
func (l *LangSet) Compare(l2 *LangSet) LangResult { return langSetCompare(l, l2) }
func (l *LangSet) Contains(l2 *LangSet) bool      { return langSetContains(l, l2) }
func (l *LangSet) Copy() *LangSet                 { return langSetCopy(l) }
func (l *LangSet) Destroy()                       { langSetDestroy(l) }
func (l *LangSet) Equal(l2 *LangSet) bool         { return langSetEqual(l, l2) }
func (l *LangSet) GetLangs() *StrSet              { return langSetGetLangs(l) }
func (l *LangSet) Hash() Char32                   { return langSetHash(l) }
func (l *LangSet) HasLang(lang string) LangResult { return langSetHasLang(l, lang) }

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

	matrixCopy   func(m *Matrix) *Matrix
	matrixEqual  func(m1, m2 *Matrix) bool
	matrixRotate func(m *Matrix, c, s float64)
	matrixScale  func(m *Matrix, sx, sy float64)
	matrixShear  func(m *Matrix, sh, sv float64)
)

func (m *Matrix) Copy() *Matrix         { return matrixCopy(m) }
func (m *Matrix) Equal(m2 *Matrix) bool { return matrixEqual(m, m2) }
func (m *Matrix) Rotate(c, s float64)   { matrixRotate(m, c, s) }
func (m *Matrix) Scale(sx, sy float64)  { matrixScale(m, sx, sy) }
func (m *Matrix) Shear(sh, sv float64)  { matrixShear(m, sh, sv) }

type ObjectSet struct {
	Nobject int
	Sobject int
	Objects **T.Char
}

var (
	ObjectSetBuild   func(first string, v ...VArg) *ObjectSet
	ObjectSetCreate  func() *ObjectSet
	ObjectSetVaBuild func(first string, va T.VaList) *ObjectSet

	objectSetAdd     func(os *ObjectSet, object string) bool
	objectSetDestroy func(os *ObjectSet)
)

func (o *ObjectSet) Add(object string) bool { return objectSetAdd(o, object) }
func (o *ObjectSet) Destroy()               { objectSetDestroy(o) }

type ObjectType struct {
	Object *T.Char
	Type   Type
}

type Pattern struct{}

var (
	PatternCreate func() *Pattern

	DefaultSubstitute func(p *Pattern)
	NameUnparse       func(p *Pattern) string

	patternAdd         func(p *Pattern, object string, value Value, append bool) bool
	patternAddBool     func(p *Pattern, object string, b bool) bool
	patternAddCharSet  func(p *Pattern, object string, c *CharSet) bool
	patternAddDouble   func(p *Pattern, object string, d float64) bool
	patternAddFTFace   func(p *Pattern, object string, f T.FTFace) bool
	patternAddInteger  func(p *Pattern, object string, i int) bool
	patternAddLangSet  func(p *Pattern, object string, ls *LangSet) bool
	patternAddMatrix   func(p *Pattern, object string, s *Matrix) bool
	patternAddString   func(p *Pattern, object string, s string) bool
	patternAddWeak     func(p *Pattern, object string, value Value, append bool) bool
	patternBuild       func(p *Pattern, v ...VArg) *Pattern
	patternDel         func(p *Pattern, object string) bool
	patternDestroy     func(p *Pattern)
	patternDuplicate   func(p *Pattern) *Pattern
	patternEqual       func(p *Pattern, pb *Pattern) bool
	patternEqualSubset func(p *Pattern, pb *Pattern, os *ObjectSet) bool
	patternFilter      func(p *Pattern, os *ObjectSet) *Pattern
	patternFormat      func(p *Pattern, format string) string
	patternGet         func(p *Pattern, object string, id int, v *Value) Result
	patternGetBool     func(p *Pattern, object string, n int, b *bool) Result
	patternGetCharSet  func(p *Pattern, object string, n int, c **CharSet) Result
	patternGetDouble   func(p *Pattern, object string, n int, d *float64) Result
	patternGetFTFace   func(p *Pattern, object string, n int, f *T.FTFace) Result
	patternGetInteger  func(p *Pattern, object string, n int, i *int) Result
	patternGetLangSet  func(p *Pattern, object string, n int, ls **LangSet) Result
	patternGetMatrix   func(p *Pattern, object string, n int, s **Matrix) Result
	patternGetString   func(p *Pattern, object string, n int, s *string) Result
	patternHash        func(p *Pattern) Char32
	patternPrint       func(p *Pattern)
	patternReference   func(p *Pattern)
	patternRemove      func(p *Pattern, object string, id int) bool
	patternVaBuild     func(p *Pattern, va T.VaList) *Pattern
)

func (p *Pattern) Add(object string, value Value, append bool) bool {
	return patternAdd(p, object, value, append)
}
func (p *Pattern) Addbool(object string, b bool) bool         { return patternAddBool(p, object, b) }
func (p *Pattern) AddCharSet(object string, c *CharSet) bool  { return patternAddCharSet(p, object, c) }
func (p *Pattern) AddDouble(object string, d float64) bool    { return patternAddDouble(p, object, d) }
func (p *Pattern) AddFTFace(object string, f T.FTFace) bool   { return patternAddFTFace(p, object, f) }
func (p *Pattern) AddInteger(object string, i int) bool       { return patternAddInteger(p, object, i) }
func (p *Pattern) AddLangSet(object string, ls *LangSet) bool { return patternAddLangSet(p, object, ls) }
func (p *Pattern) AddMatrix(object string, s *Matrix) bool    { return patternAddMatrix(p, object, s) }
func (p *Pattern) AddString(object string, s string) bool     { return patternAddString(p, object, s) }
func (p *Pattern) AddWeak(object string, value Value, append bool) bool {
	return patternAddWeak(p, object, value, append)
}
func (p *Pattern) Build(v ...VArg) *Pattern                    { return patternBuild(p, v) }
func (p *Pattern) Del(object string) bool                      { return patternDel(p, object) }
func (p *Pattern) Destroy()                                    { patternDestroy(p) }
func (p *Pattern) Duplicate() *Pattern                         { return patternDuplicate(p) }
func (p *Pattern) Equal(pb *Pattern) bool                      { return patternEqual(p, pb) }
func (p *Pattern) EqualSubset(pb *Pattern, os *ObjectSet) bool { return patternEqualSubset(p, pb, os) }
func (p *Pattern) Filter(os *ObjectSet) *Pattern               { return patternFilter(p, os) }
func (p *Pattern) Format(format string) string                 { return patternFormat(p, format) }
func (p *Pattern) Get(object string, id int, v *Value) Result  { return patternGet(p, object, id, v) }
func (p *Pattern) GetBool(object string, n int, b *bool) Result {
	return patternGetBool(p, object, n, b)
}
func (p *Pattern) GetCharSet(object string, n int, c **CharSet) Result {
	return patternGetCharSet(p, object, n, c)
}
func (p *Pattern) GetDouble(object string, n int, d *float64) Result {
	return patternGetDouble(p, object, n, d)
}
func (p *Pattern) GetFTFace(object string, n int, f *T.FTFace) Result {
	return patternGetFTFace(p, object, n, f)
}
func (p *Pattern) GetInteger(object string, n int, i *int) Result {
	return patternGetInteger(p, object, n, i)
}
func (p *Pattern) GetLangSet(object string, n int, ls **LangSet) Result {
	return patternGetLangSet(p, object, n, ls)
}
func (p *Pattern) GetMatrix(object string, n int, s **Matrix) Result {
	return patternGetMatrix(p, object, n, s)
}
func (p *Pattern) GetString(object string, n int, s *string) Result {
	return patternGetString(p, object, n, s)
}
func (p *Pattern) Hash() Char32                      { return patternHash(p) }
func (p *Pattern) Print()                            { patternPrint(p) }
func (p *Pattern) Reference()                        { patternReference(p) }
func (p *Pattern) Remove(object string, id int) bool { return patternRemove(p, object, id) }
func (p *Pattern) VaBuild(va T.VaList) *Pattern      { return patternVaBuild(p, va) }

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

	strListDone func(s *StrList)
	strListNext func(s *StrList) string
)

func (s *StrList) Done()        { strListDone(s) }
func (s *StrList) Next() string { return strListNext(s) }

type StrSet struct{}

var (
	StrSetCreate func() *StrSet

	strSetAdd         func(s *StrSet, str string) bool
	strSetAddFilename func(s *StrSet, str string) bool
	strSetDel         func(s *StrSet, str string) bool
	strSetDestroy     func(s *StrSet)
	strSetEqual       func(s, s2 *StrSet) bool
	strSetMember      func(s *StrSet, str string) bool
)

func (s *StrSet) Add(str string) bool         { return strSetAdd(s, str) }
func (s *StrSet) AddFilename(str string) bool { return strSetAddFilename(s, str) }
func (s *StrSet) Del(str string) bool         { return strSetDel(s, str) }
func (s *StrSet) Destroy()                    { strSetDestroy(s) }
func (s *StrSet) Equal(s2 *StrSet) bool       { return strSetEqual(s, s2) }
func (s *StrSet) Member(str string) bool      { return strSetMember(s, str) }

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
	valueDestroy func(v Value)
	valueEqual   func(va, vb Value) bool
	valuePrint   func(v Value)
	valueSave    func(v Value) Value
)

func (v Value) Destroy()            { valueDestroy(v) }
func (v Value) Equal(v2 Value) bool { return valueEqual(v, v2) }
func (v Value) Print()              { valuePrint(v) }
func (v Value) Save() Value         { return valueSave(v) }

var dll = "libfontconfig-1.dll"

var apiList = outside.Apis{
	{"FcAtomicCreate", &AtomicCreate},
	{"FcAtomicDeleteNew", &atomicDeleteNew},
	{"FcAtomicDestroy", &atomicDestroy},
	{"FcAtomicLock", &atomicLock},
	{"FcAtomicNewFile", &atomicNewFile},
	{"FcAtomicOrigFile", &atomicOrigFile},
	{"FcAtomicReplaceOrig", &atomicReplaceOrig},
	{"FcAtomicUnlock", &atomicUnlock},
	{"FcBlanksAdd", &blanksAdd},
	{"FcBlanksCreate", &BlanksCreate},
	{"FcBlanksDestroy", &blanksDestroy},
	{"FcBlanksIsMember", &blanksIsMember},
	{"FcCacheCopySet", &cacheCopySet},
	{"FcCacheDir", &cacheDir},
	{"FcCacheNumFont", &cacheNumFont},
	{"FcCacheNumSubdir", &cacheNumSubdir},
	{"FcCacheSubdir", &cacheSubdir},
	{"FcCharSetAddChar", &charSetAddChar},
	{"FcCharSetCopy", &charSetCopy},
	{"FcCharSetCount", &charSetCount},
	{"FcCharSetCoverage", &charSetCoverage},
	{"FcCharSetCreate", &CharSetCreate},
	{"FcCharSetDestroy", &charSetDestroy},
	{"FcCharSetEqual", &charSetEqual},
	{"FcCharSetFirstPage", &charSetFirstPage},
	{"FcCharSetHasChar", &charSetHasChar},
	{"FcCharSetIntersect", &charSetIntersect},
	{"FcCharSetIntersectCount", &charSetIntersectCount},
	{"FcCharSetIsSubset", &charSetIsSubset},
	{"FcCharSetMerge", &charSetMerge},
	{"FcCharSetNew", &CharSetNew},
	{"FcCharSetNextPage", &charSetNextPage},
	{"FcCharSetSubtract", &charSetSubtract},
	{"FcCharSetSubtractCount", &charSetSubtractCount},
	{"FcCharSetUnion", &charSetUnion},
	{"FcConfigAppFontAddDir", &configAppFontAddDir},
	{"FcConfigAppFontAddFile", &configAppFontAddFile},
	{"FcConfigAppFontClear", &configAppFontClear},
	{"FcConfigBuildFonts", &configBuildFonts},
	{"FcConfigCreate", &ConfigCreate},
	{"FcConfigDestroy", &configDestroy},
	{"FcConfigEnableHome", &ConfigEnableHome},
	{"FcConfigFilename", &ConfigFilename},
	{"FcConfigGetBlanks", &configGetBlanks},
	{"FcConfigGetCache", &configGetCache},
	{"FcConfigGetCacheDirs", &configGetCacheDirs},
	{"FcConfigGetConfigDirs", &configGetConfigDirs},
	{"FcConfigGetConfigFiles", &configGetConfigFiles},
	{"FcConfigGetCurrent", &ConfigGetCurrent},
	{"FcConfigGetFontDirs", &configGetFontDirs},
	{"FcConfigGetFonts", &configGetFonts},
	// Deprecated/Undocumented {"FcConfigGetRescanInterval", &ConfigGetRescanInterval},
	{"FcConfigHome", &ConfigHome},
	{"FcConfigParseAndLoad", &configParseAndLoad},
	{"FcConfigReference", &configReference},
	{"FcConfigSetCurrent", &configSetCurrent},
	// Deprecated/Undocumented {"FcConfigSetRescanInterval", &ConfigSetRescanInterval},
	{"FcConfigSubstitute", &configSubstitute},
	{"FcConfigSubstituteWithPat", &configSubstituteWithPat},
	{"FcConfigUptoDate", &configUptoDate},
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
	{"FcFontSetAdd", &fontSetAdd},
	{"FcFontSetCreate", &FontSetCreate},
	{"FcFontSetDestroy", &fontSetDestroy},
	{"FcFontSetList", &FontSetList},
	{"FcFontSetMatch", &FontSetMatch},
	{"FcFontSetPrint", &fontSetPrint},
	{"FcFontSetSort", &FontSetSort},
	{"FcFontSetSortDestroy", &fontSetSortDestroy},
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
	{"FcLangSetAdd", &langSetAdd},
	{"FcLangSetCompare", &langSetCompare},
	{"FcLangSetContains", &langSetContains},
	{"FcLangSetCopy", &langSetCopy},
	{"FcLangSetCreate", &LangSetCreate},
	{"FcLangSetDestroy", &langSetDestroy},
	{"FcLangSetEqual", &langSetEqual},
	{"FcLangSetGetLangs", &langSetGetLangs},
	{"FcLangSetHasLang", &langSetHasLang},
	{"FcLangSetHash", &langSetHash},
	{"FcMatrixCopy", &matrixCopy},
	{"FcMatrixEqual", &matrixEqual},
	{"FcMatrixMultiply", &MatrixMultiply},
	{"FcMatrixRotate", &matrixRotate},
	{"FcMatrixScale", &matrixScale},
	{"FcMatrixShear", &matrixShear},
	{"FcNameConstant", &NameConstant},
	{"FcNameGetConstant", &NameGetConstant},
	{"FcNameGetObjectType", &NameGetObjectType},
	{"FcNameParse", &NameParse},
	{"FcNameRegisterConstants", &NameRegisterConstants},
	{"FcNameRegisterObjectTypes", &NameRegisterObjectTypes},
	{"FcNameUnparse", &NameUnparse},
	{"FcNameUnregisterConstants", &NameUnregisterConstants},
	{"FcNameUnregisterObjectTypes", &NameUnregisterObjectTypes},
	{"FcObjectSetAdd", &objectSetAdd},
	{"FcObjectSetBuild", &ObjectSetBuild},
	{"FcObjectSetCreate", &ObjectSetCreate},
	{"FcObjectSetDestroy", &objectSetDestroy},
	{"FcObjectSetVaBuild", &ObjectSetVaBuild},
	{"FcPatternAdd", &patternAdd},
	{"FcPatternAddBool", &patternAddBool},
	{"FcPatternAddCharSet", &patternAddCharSet},
	{"FcPatternAddDouble", &patternAddDouble},
	{"FcPatternAddFTFace", &patternAddFTFace},
	{"FcPatternAddInteger", &patternAddInteger},
	{"FcPatternAddLangSet", &patternAddLangSet},
	{"FcPatternAddMatrix", &patternAddMatrix},
	{"FcPatternAddString", &patternAddString},
	{"FcPatternAddWeak", &patternAddWeak},
	{"FcPatternBuild", &patternBuild},
	{"FcPatternCreate", &PatternCreate},
	{"FcPatternDel", &patternDel},
	{"FcPatternDestroy", &patternDestroy},
	{"FcPatternDuplicate", &patternDuplicate},
	{"FcPatternEqual", &patternEqual},
	{"FcPatternEqualSubset", &patternEqualSubset},
	{"FcPatternFilter", &patternFilter},
	{"FcPatternFormat", &patternFormat},
	{"FcPatternGet", &patternGet},
	{"FcPatternGetBool", &patternGetBool},
	{"FcPatternGetCharSet", &patternGetCharSet},
	{"FcPatternGetDouble", &patternGetDouble},
	{"FcPatternGetFTFace", &patternGetFTFace},
	{"FcPatternGetInteger", &patternGetInteger},
	{"FcPatternGetLangSet", &patternGetLangSet},
	{"FcPatternGetMatrix", &patternGetMatrix},
	{"FcPatternGetString", &patternGetString},
	{"FcPatternHash", &patternHash},
	{"FcPatternPrint", &patternPrint},
	{"FcPatternReference", &patternReference},
	{"FcPatternRemove", &patternRemove},
	{"FcPatternVaBuild", &patternVaBuild},
	{"FcStrBasename", &StrBasename},
	{"FcStrCmp", &StrCmp},
	{"FcStrCmpIgnoreCase", &StrCmpIgnoreCase},
	{"FcStrCopy", &StrCopy},
	{"FcStrCopyFilename", &StrCopyFilename},
	{"FcStrDirname", &StrDirname},
	{"FcStrDowncase", &StrDowncase},
	{"FcStrFree", &StrFree},
	{"FcStrListCreate", &StrListCreate},
	{"FcStrListDone", &strListDone},
	{"FcStrListNext", &strListNext},
	{"FcStrPlus", &StrPlus},
	{"FcStrSetAdd", &strSetAdd},
	{"FcStrSetAddFilename", &strSetAddFilename},
	{"FcStrSetCreate", &StrSetCreate},
	{"FcStrSetDel", &strSetDel},
	{"FcStrSetDestroy", &strSetDestroy},
	{"FcStrSetEqual", &strSetEqual},
	{"FcStrSetMember", &strSetMember},
	{"FcStrStr", &StrStr},
	{"FcStrStrIgnoreCase", &StrStrIgnoreCase},
	{"FcUcs4ToUtf8", &Ucs4ToUtf8},
	{"FcUtf16Len", &Utf16Len},
	{"FcUtf16ToUcs4", &Utf16ToUcs4},
	{"FcUtf8Len", &Utf8Len},
	{"FcUtf8ToUcs4", &Utf8ToUcs4},
	{"FcValueDestroy", &valueDestroy},
	{"FcValueEqual", &valueEqual},
	{"FcValuePrint", &valuePrint},
	{"FcValueSave", &valueSave},
}
