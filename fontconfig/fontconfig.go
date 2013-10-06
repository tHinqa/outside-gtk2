package fontconfig

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
}

const (
	FC_CHARSET_MAP_SIZE = 256 / 32
	FC_UTF8_MAX_LEN     = 6
)

type (
	FcChar8  Unsigned_char
	FcChar16 Unsigned_short
	FcChar32 Unsigned_int
	FcBool   int

	FcAtomic    struct{}
	FcBlanks    struct{}
	FcCache     struct{}
	FcConfig    struct{}
	FcStrList   struct{}
	FcStrSet    struct{}
	FcFileCache struct{}
	FcLangSet   struct{}
	_           struct{}
	_           struct{}

	FcFontSet struct {
		nfont int
		sfont int
		fonts **FcPattern
	}

	FcObjectSet struct {
		nobject int
		sobject int
		objects **Char
	}

	FcMatrix struct {
		xx, xy, yx, yy Double
	}

	FcObjectType struct {
		object *Char
		Type   FcType
	}

	FcConstant struct {
		name   *FcChar8
		object *Char
		value  int
	}

	FcValue struct {
		Type FcType
		// Union
		s *FcChar8
		// i int
		// b FcBool
		// d double
		// m *FcMatrix
		// c *FcCharSet
		// f *void
		// l *FcLangSet
	}
)

type FcType Enum

const (
	FcTypeVoid FcType = iota
	FcTypeInteger
	FcTypeDouble
	FcTypeString
	FcTypeBool
	FcTypeMatrix
	FcTypeCharSet
	FcTypeFTFace
	FcTypeLangSet
)

type FcSetName Enum

const (
	FcSetSystem FcSetName = iota
	FcSetApplication
)

type FcMatchKind Enum

const (
	FcMatchPattern FcMatchKind = iota
	FcMatchFont
	FcMatchScan
)

type FcResult Enum

const (
	FcResultMatch FcResult = iota
	FcResultNoMatch
	FcResultTypeMismatch
	FcResultNoId
	FcResultOutOfMemory
)

type FcLangResult Enum

const (
	FcLangEqual FcLangResult = 0
	FcLangDifferentCountry
	FcLangDifferentLang
	FcLangDifferentTerritory = FcLangDifferentCountry
)

type FcEndian Enum

const (
	FcEndianBig FcEndian = iota
	FcEndianLittle
)

var (
	FcBlanksCreate func() *FcBlanks

	FcBlanksDestroy func(b *FcBlanks)

	FcBlanksAdd func(b *FcBlanks, ucs4 FcChar32) FcBool

	FcBlanksIsMember func(b *FcBlanks, ucs4 FcChar32) FcBool

	FcCacheDir func(c *FcCache) string

	FcCacheCopySet func(c *FcCache) *FcFontSet

	FcCacheSubdir func(c *FcCache, i int) string

	FcCacheNumSubdir func(c *FcCache) int

	FcCacheNumFont func(c *FcCache) int

	FcDirCacheUnlink func(dir string, config *FcConfig) FcBool

	FcDirCacheValid func(cache_file string) FcBool

	FcConfigHome func() string

	FcConfigEnableHome func(enable FcBool) FcBool

	FcConfigFilename func(url string) string

	FcConfigCreate func() *FcConfig

	FcConfigReference func(config *FcConfig) *FcConfig

	FcConfigDestroy func(config *FcConfig)

	FcConfigSetCurrent func(config *FcConfig) FcBool

	FcConfigGetCurrent func() *FcConfig

	FcConfigUptoDate func(config *FcConfig) FcBool

	FcConfigBuildFonts func(config *FcConfig) FcBool

	FcConfigGetFontDirs func(config *FcConfig) *FcStrList

	FcConfigGetConfigDirs func(config *FcConfig) *FcStrList

	FcConfigGetConfigFiles func(config *FcConfig) *FcStrList

	FcConfigGetCache func(config *FcConfig) string

	FcConfigGetBlanks func(config *FcConfig) *FcBlanks

	FcConfigGetCacheDirs func(config *FcConfig) *FcStrList

	FcConfigGetRescanInterval func(config *FcConfig) int

	FcConfigSetRescanInterval func(
		config *FcConfig, rescanInterval int) FcBool

	FcConfigGetFonts func(
		config *FcConfig, set FcSetName) *FcFontSet

	FcConfigAppFontAddFile func(
		config *FcConfig, file string) FcBool

	FcConfigAppFontAddDir func(
		config *FcConfig, dir string) FcBool

	FcConfigAppFontClear func(config *FcConfig)

	FcConfigSubstituteWithPat func(config *FcConfig,
		p *FcPattern, p_pat *FcPattern, kind FcMatchKind) FcBool

	FcConfigSubstitute func(
		config *FcConfig, p *FcPattern, kind FcMatchKind) FcBool

	FcCharSetCreate func() *FcCharSet

	FcCharSetNew func() *FcCharSet

	FcCharSetDestroy func(fcs *FcCharSet)

	FcCharSetAddChar func(fcs *FcCharSet, ucs4 FcChar32) FcBool

	FcCharSetCopy func(src *FcCharSet) *FcCharSet

	FcCharSetEqual func(a, b *FcCharSet) FcBool

	FcCharSetIntersect func(a, b *FcCharSet) *FcCharSet

	FcCharSetUnion func(a, b *FcCharSet) *FcCharSet

	FcCharSetSubtract func(a, b *FcCharSet) *FcCharSet

	FcCharSetMerge func(a, b *FcCharSet, changed *FcBool) FcBool

	FcCharSetHasChar func(fcs *FcCharSet, ucs4 FcChar32) FcBool

	FcCharSetCount func(a *FcCharSet) FcChar32

	FcCharSetIntersectCount func(a, b *FcCharSet) FcChar32

	FcCharSetSubtractCount func(a, b *FcCharSet) FcChar32

	FcCharSetIsSubset func(a, b *FcCharSet) FcBool

	FcCharSetFirstPage func(a *FcCharSet,
		Map [FC_CHARSET_MAP_SIZE]FcChar32,
		next *FcChar32) FcChar32

	FcCharSetNextPage func(a *FcCharSet,
		Map [FC_CHARSET_MAP_SIZE]FcChar32,
		next *FcChar32) FcChar32

	FcCharSetCoverage func(
		a *FcCharSet, page FcChar32, result *FcChar32) FcChar32

	FcValuePrint func(v FcValue)

	FcPatternPrint func(p *FcPattern)

	FcFontSetPrint func(s *FcFontSet)

	FcDefaultSubstitute func(pattern *FcPattern)

	FcFileIsDir func(file string) FcBool

	FcFileScan func(set *FcFontSet, dirs *FcStrSet,
		cache *FcFileCache, blanks *FcBlanks,
		file string, force FcBool) FcBool

	FcDirScan func(set *FcFontSet, dirs *FcStrSet,
		cache *FcFileCache, blanks *FcBlanks, dir string,
		force FcBool) FcBool

	FcDirSave func(
		set *FcFontSet, dirs *FcStrSet, dir string) FcBool

	FcDirCacheLoad func(dir string,
		config *FcConfig, cache_file **FcChar8) *FcCache

	FcDirCacheRead func(
		dir string, force FcBool, config *FcConfig) *FcCache

	FcDirCacheLoadFile func(
		cache_file string, file_stat *Stat) *FcCache

	FcDirCacheUnload func(cache *FcCache)

	FcFreeTypeQuery func(file string,
		id int, blanks *FcBlanks, count *int) *FcPattern

	FcFontSetCreate func() *FcFontSet

	FcFontSetDestroy func(s *FcFontSet)

	FcFontSetAdd func(s *FcFontSet, font *FcPattern) FcBool

	FcInitLoadConfig func() *FcConfig

	FcInitLoadConfigAndFonts func() *FcConfig

	FcInit func() FcBool

	FcFini func()

	FcGetVersion func() int

	FcInitReinitialize func() FcBool

	FcInitBringUptoDate func() FcBool

	FcGetLangs func() *FcStrSet

	FcLangGetCharSet func(lang string) *FcCharSet

	FcLangSetCreate func() *FcLangSet

	FcLangSetDestroy func(ls *FcLangSet)

	FcLangSetCopy func(ls *FcLangSet) *FcLangSet

	FcLangSetAdd func(ls *FcLangSet, lang string) FcBool

	FcLangSetHasLang func(ls *FcLangSet, lang string) FcLangResult

	FcLangSetCompare func(lsa, lsb *FcLangSet) FcLangResult

	FcLangSetContains func(lsa, lsb *FcLangSet) FcBool

	FcLangSetEqual func(lsa, lsb *FcLangSet) FcBool

	FcLangSetHash func(ls *FcLangSet) FcChar32

	FcLangSetGetLangs func(ls *FcLangSet) *FcStrSet

	FcObjectSetCreate func() *FcObjectSet

	FcObjectSetAdd func(os *FcObjectSet, object string) FcBool

	FcObjectSetDestroy func(os *FcObjectSet)

	FcObjectSetVaBuild func(first string, va Va_list) *FcObjectSet

	FcObjectSetBuild func(first string, v ...VArg) *FcObjectSet

	FcFontSetList func(config *FcConfig, sets **FcFontSet,
		nsets int, p *FcPattern, os *FcObjectSet) *FcFontSet

	FcFontList func(config *FcConfig,
		p *FcPattern, os *FcObjectSet) *FcFontSet

	FcAtomicCreate func(file string) *FcAtomic

	FcAtomicLock func(atomic *FcAtomic) FcBool

	FcAtomicNewFile func(atomic *FcAtomic) string

	FcAtomicOrigFile func(atomic *FcAtomic) string

	FcAtomicReplaceOrig func(atomic *FcAtomic) FcBool

	FcAtomicDeleteNew func(atomic *FcAtomic)

	FcAtomicUnlock func(atomic *FcAtomic)

	FcAtomicDestroy func(atomic *FcAtomic)

	FcFontSetMatch func(config *FcConfig, sets **FcFontSet,
		nsets int, p *FcPattern, result *FcResult) *FcPattern

	FcFontMatch func(config *FcConfig,
		p *FcPattern, result *FcResult) *FcPattern

	FcFontRenderPrepare func(config *FcConfig,
		pat *FcPattern, font *FcPattern) *FcPattern

	FcFontSetSort func(config *FcConfig, sets **FcFontSet,
		nsets int, p *FcPattern, trim FcBool,
		csp **FcCharSet, result *FcResult) *FcFontSet

	FcFontSort func(config *FcConfig, p *FcPattern,
		trim FcBool, csp **FcCharSet, result *FcResult) *FcFontSet

	FcFontSetSortDestroy func(fs *FcFontSet)

	FcMatrixCopy func(mat *FcMatrix) *FcMatrix

	FcMatrixEqual func(mat1, mat2 *FcMatrix) FcBool

	FcMatrixMultiply func(result, a, b *FcMatrix)

	FcMatrixRotate func(m *FcMatrix, c, s Double)

	FcMatrixScale func(m *FcMatrix, sx, sy Double)

	FcMatrixShear func(m *FcMatrix, sh, sv Double)

	FcNameRegisterObjectTypes func(
		types *FcObjectType, ntype int) FcBool

	FcNameUnregisterObjectTypes func(
		types *FcObjectType, ntype int) FcBool

	FcNameGetObjectType func(object string) *FcObjectType

	FcNameRegisterConstants func(
		consts *FcConstant, nconsts int) FcBool

	FcNameUnregisterConstants func(
		consts *FcConstant, nconsts int) FcBool

	FcNameGetConstant func(str string) *FcConstant

	FcNameConstant func(st string, result *int) FcBool

	FcNameParse func(name string) *FcPattern

	FcNameUnparse func(pat *FcPattern) string

	FcPatternCreate func() *FcPattern

	FcPatternDuplicate func(p *FcPattern) *FcPattern

	FcPatternReference func(p *FcPattern)

	FcPatternFilter func(p *FcPattern, os *FcObjectSet) *FcPattern

	FcValueDestroy func(v FcValue)

	FcValueEqual func(va, vb FcValue) FcBool

	FcValueSave func(v FcValue) FcValue

	FcPatternDestroy func(p *FcPattern)

	FcPatternEqual func(pa, pb *FcPattern) FcBool

	FcPatternEqualSubset func(
		pa, pb *FcPattern, os *FcObjectSet) FcBool

	FcPatternHash func(p *FcPattern) FcChar32

	FcPatternAdd func(p *FcPattern,
		object string, value FcValue, append FcBool) FcBool

	FcPatternAddWeak func(p *FcPattern,
		object string, value FcValue, append FcBool) FcBool

	FcPatternGet func(
		p *FcPattern, object string, id int, v *FcValue) FcResult

	FcPatternDel func(p *FcPattern, object string) FcBool

	FcPatternRemove func(
		p *FcPattern, object string, id int) FcBool

	FcPatternAddInteger func(
		p *FcPattern, object string, i int) FcBool

	FcPatternAddDouble func(
		p *FcPattern, object string, d Double) FcBool

	FcPatternAddString func(
		p *FcPattern, object string, s string) FcBool

	FcPatternAddMatrix func(
		p *FcPattern, object string, s *FcMatrix) FcBool

	FcPatternAddCharSet func(
		p *FcPattern, object string, c *FcCharSet) FcBool

	FcPatternAddBool func(
		p *FcPattern, object string, b FcBool) FcBool

	FcPatternAddLangSet func(
		p *FcPattern, object string, ls *FcLangSet) FcBool

	FcPatternGetInteger func(
		p *FcPattern, object string, n int, i *int) FcResult

	FcPatternGetDouble func(
		p *FcPattern, object string, n int, d *Double) FcResult

	FcPatternGetString func(
		p *FcPattern, object string, n int, s *string) FcResult

	FcPatternGetMatrix func(
		p *FcPattern, object string, n int, s **FcMatrix) FcResult

	FcPatternGetCharSet func(p *FcPattern,
		object string, n int, c **FcCharSet) FcResult

	FcPatternGetBool func(
		p *FcPattern, object string, n int, b *FcBool) FcResult

	FcPatternGetLangSet func(p *FcPattern,
		object string, n int, ls **FcLangSet) FcResult

	FcPatternVaBuild func(p *FcPattern, va Va_list) *FcPattern

	FcPatternBuild func(p *FcPattern, v ...VArg) *FcPattern

	FcPatternFormat func(pat *FcPattern, format string) string

	FcStrCopy func(s string) string

	FcStrCopyFilename func(s string) string

	FcStrPlus func(s1, s2 string) string

	FcStrFree func(s string)

	FcStrDowncase func(s string) string

	FcStrCmpIgnoreCase func(s1, s2 string) int

	FcStrCmp func(s1, s2 string) int

	FcStrStrIgnoreCase func(s1, s2 string) string

	FcStrStr func(s1, s2 string) string

	FcUtf8ToUcs4 func(
		src_orig string, dst *FcChar32, leng int) int

	FcUtf8Len func(
		FcChar8 *string, leng int, nchar, wchar *int) FcBool

	FcUcs4ToUtf8 func(
		ucs4 FcChar32, dest [FC_UTF8_MAX_LEN]FcChar8) int

	FcUtf16ToUcs4 func(src_orig string, endian FcEndian,
		dst *FcChar32, leng int) int

	FcUtf16Len func(FcChar8 *string, endian FcEndian,
		leng int, nchar, wchar *int) FcBool

	FcStrDirname func(file string) string

	FcStrBasename func(file string) string

	FcStrSetCreate func() *FcStrSet

	FcStrSetMember func(set *FcStrSet, s string) FcBool

	FcStrSetEqual func(sa, sb *FcStrSet) FcBool

	FcStrSetAdd func(set *FcStrSet, s string) FcBool

	FcStrSetAddFilename func(set *FcStrSet, s string) FcBool

	FcStrSetDel func(set *FcStrSet, s string) FcBool

	FcStrSetDestroy func(set *FcStrSet)

	FcStrListCreate func(set *FcStrSet) *FcStrList

	FcStrListNext func(list *FcStrList) string

	FcStrListDone func(list *FcStrList)

	FcConfigParseAndLoad func(
		config *FcConfig, file string, complain FcBool) FcBool

	FcFreeTypeCharIndex func(face FT_Face, ucs4 FcChar32) FT_UInt

	FcFreeTypeCharSetAndSpacing func(
		face FT_Face, blanks *FcBlanks, spacing *int) *FcCharSet

	FcFreeTypeCharSet func(
		face FT_Face, blanks *FcBlanks) *FcCharSet

	FcPatternGetFTFace func(
		p *FcPattern, object string, n int, f *FT_Face) FcResult

	FcPatternAddFTFace func(
		p *FcPattern, object string, f FT_Face) FcBool

	FcFreeTypeQueryFace func(face FT_Face,
		file string, id int, blanks *FcBlanks) *FcPattern
)

var dll = "libfontconfig-1.dll"

var apiList = Apis{
	{"FcAtomicCreate", &FcAtomicCreate},
	{"FcAtomicDeleteNew", &FcAtomicDeleteNew},
	{"FcAtomicDestroy", &FcAtomicDestroy},
	{"FcAtomicLock", &FcAtomicLock},
	{"FcAtomicNewFile", &FcAtomicNewFile},
	{"FcAtomicOrigFile", &FcAtomicOrigFile},
	{"FcAtomicReplaceOrig", &FcAtomicReplaceOrig},
	{"FcAtomicUnlock", &FcAtomicUnlock},
	{"FcBlanksAdd", &FcBlanksAdd},
	{"FcBlanksCreate", &FcBlanksCreate},
	{"FcBlanksDestroy", &FcBlanksDestroy},
	{"FcBlanksIsMember", &FcBlanksIsMember},
	{"FcCacheCopySet", &FcCacheCopySet},
	{"FcCacheDir", &FcCacheDir},
	{"FcCacheNumFont", &FcCacheNumFont},
	{"FcCacheNumSubdir", &FcCacheNumSubdir},
	{"FcCacheSubdir", &FcCacheSubdir},
	{"FcCharSetAddChar", &FcCharSetAddChar},
	{"FcCharSetCopy", &FcCharSetCopy},
	{"FcCharSetCount", &FcCharSetCount},
	{"FcCharSetCoverage", &FcCharSetCoverage},
	{"FcCharSetCreate", &FcCharSetCreate},
	{"FcCharSetDestroy", &FcCharSetDestroy},
	{"FcCharSetEqual", &FcCharSetEqual},
	{"FcCharSetFirstPage", &FcCharSetFirstPage},
	{"FcCharSetHasChar", &FcCharSetHasChar},
	{"FcCharSetIntersect", &FcCharSetIntersect},
	{"FcCharSetIntersectCount", &FcCharSetIntersectCount},
	{"FcCharSetIsSubset", &FcCharSetIsSubset},
	{"FcCharSetMerge", &FcCharSetMerge},
	{"FcCharSetNew", &FcCharSetNew},
	{"FcCharSetNextPage", &FcCharSetNextPage},
	{"FcCharSetSubtract", &FcCharSetSubtract},
	{"FcCharSetSubtractCount", &FcCharSetSubtractCount},
	{"FcCharSetUnion", &FcCharSetUnion},
	{"FcConfigAppFontAddDir", &FcConfigAppFontAddDir},
	{"FcConfigAppFontAddFile", &FcConfigAppFontAddFile},
	{"FcConfigAppFontClear", &FcConfigAppFontClear},
	{"FcConfigBuildFonts", &FcConfigBuildFonts},
	{"FcConfigCreate", &FcConfigCreate},
	{"FcConfigDestroy", &FcConfigDestroy},
	{"FcConfigEnableHome", &FcConfigEnableHome},
	{"FcConfigFilename", &FcConfigFilename},
	{"FcConfigGetBlanks", &FcConfigGetBlanks},
	{"FcConfigGetCache", &FcConfigGetCache},
	{"FcConfigGetCacheDirs", &FcConfigGetCacheDirs},
	{"FcConfigGetConfigDirs", &FcConfigGetConfigDirs},
	{"FcConfigGetConfigFiles", &FcConfigGetConfigFiles},
	{"FcConfigGetCurrent", &FcConfigGetCurrent},
	{"FcConfigGetFontDirs", &FcConfigGetFontDirs},
	{"FcConfigGetFonts", &FcConfigGetFonts},
	// Deprecated/Undocumented {"FcConfigGetRescanInterval", &FcConfigGetRescanInterval},
	{"FcConfigHome", &FcConfigHome},
	{"FcConfigParseAndLoad", &FcConfigParseAndLoad},
	{"FcConfigReference", &FcConfigReference},
	{"FcConfigSetCurrent", &FcConfigSetCurrent},
	// Deprecated/Undocumented {"FcConfigSetRescanInterval", &FcConfigSetRescanInterval},
	{"FcConfigSubstitute", &FcConfigSubstitute},
	{"FcConfigSubstituteWithPat", &FcConfigSubstituteWithPat},
	{"FcConfigUptoDate", &FcConfigUptoDate},
	{"FcDefaultSubstitute", &FcDefaultSubstitute},
	{"FcDirCacheLoad", &FcDirCacheLoad},
	{"FcDirCacheLoadFile", &FcDirCacheLoadFile},
	{"FcDirCacheRead", &FcDirCacheRead},
	{"FcDirCacheUnlink", &FcDirCacheUnlink},
	{"FcDirCacheUnload", &FcDirCacheUnload},
	{"FcDirCacheValid", &FcDirCacheValid},
	{"FcDirSave", &FcDirSave},
	{"FcDirScan", &FcDirScan},
	{"FcFileIsDir", &FcFileIsDir},
	{"FcFileScan", &FcFileScan},
	{"FcFini", &FcFini},
	{"FcFontList", &FcFontList},
	{"FcFontMatch", &FcFontMatch},
	{"FcFontRenderPrepare", &FcFontRenderPrepare},
	{"FcFontSetAdd", &FcFontSetAdd},
	{"FcFontSetCreate", &FcFontSetCreate},
	{"FcFontSetDestroy", &FcFontSetDestroy},
	{"FcFontSetList", &FcFontSetList},
	{"FcFontSetMatch", &FcFontSetMatch},
	{"FcFontSetPrint", &FcFontSetPrint},
	{"FcFontSetSort", &FcFontSetSort},
	{"FcFontSetSortDestroy", &FcFontSetSortDestroy},
	{"FcFontSort", &FcFontSort},
	{"FcFreeTypeCharIndex", &FcFreeTypeCharIndex},
	{"FcFreeTypeCharSet", &FcFreeTypeCharSet},
	{"FcFreeTypeCharSetAndSpacing", &FcFreeTypeCharSetAndSpacing},
	{"FcFreeTypeQuery", &FcFreeTypeQuery},
	{"FcFreeTypeQueryFace", &FcFreeTypeQueryFace},
	{"FcGetLangs", &FcGetLangs},
	{"FcGetVersion", &FcGetVersion},
	{"FcInit", &FcInit},
	{"FcInitBringUptoDate", &FcInitBringUptoDate},
	{"FcInitLoadConfig", &FcInitLoadConfig},
	{"FcInitLoadConfigAndFonts", &FcInitLoadConfigAndFonts},
	{"FcInitReinitialize", &FcInitReinitialize},
	{"FcLangGetCharSet", &FcLangGetCharSet},
	{"FcLangSetAdd", &FcLangSetAdd},
	{"FcLangSetCompare", &FcLangSetCompare},
	{"FcLangSetContains", &FcLangSetContains},
	{"FcLangSetCopy", &FcLangSetCopy},
	{"FcLangSetCreate", &FcLangSetCreate},
	{"FcLangSetDestroy", &FcLangSetDestroy},
	{"FcLangSetEqual", &FcLangSetEqual},
	{"FcLangSetGetLangs", &FcLangSetGetLangs},
	{"FcLangSetHasLang", &FcLangSetHasLang},
	{"FcLangSetHash", &FcLangSetHash},
	{"FcMatrixCopy", &FcMatrixCopy},
	{"FcMatrixEqual", &FcMatrixEqual},
	{"FcMatrixMultiply", &FcMatrixMultiply},
	{"FcMatrixRotate", &FcMatrixRotate},
	{"FcMatrixScale", &FcMatrixScale},
	{"FcMatrixShear", &FcMatrixShear},
	{"FcNameConstant", &FcNameConstant},
	{"FcNameGetConstant", &FcNameGetConstant},
	{"FcNameGetObjectType", &FcNameGetObjectType},
	{"FcNameParse", &FcNameParse},
	{"FcNameRegisterConstants", &FcNameRegisterConstants},
	{"FcNameRegisterObjectTypes", &FcNameRegisterObjectTypes},
	{"FcNameUnparse", &FcNameUnparse},
	{"FcNameUnregisterConstants", &FcNameUnregisterConstants},
	{"FcNameUnregisterObjectTypes", &FcNameUnregisterObjectTypes},
	{"FcObjectSetAdd", &FcObjectSetAdd},
	{"FcObjectSetBuild", &FcObjectSetBuild},
	{"FcObjectSetCreate", &FcObjectSetCreate},
	{"FcObjectSetDestroy", &FcObjectSetDestroy},
	{"FcObjectSetVaBuild", &FcObjectSetVaBuild},
	{"FcPatternAdd", &FcPatternAdd},
	{"FcPatternAddBool", &FcPatternAddBool},
	{"FcPatternAddCharSet", &FcPatternAddCharSet},
	{"FcPatternAddDouble", &FcPatternAddDouble},
	{"FcPatternAddFTFace", &FcPatternAddFTFace},
	{"FcPatternAddInteger", &FcPatternAddInteger},
	{"FcPatternAddLangSet", &FcPatternAddLangSet},
	{"FcPatternAddMatrix", &FcPatternAddMatrix},
	{"FcPatternAddString", &FcPatternAddString},
	{"FcPatternAddWeak", &FcPatternAddWeak},
	{"FcPatternBuild", &FcPatternBuild},
	{"FcPatternCreate", &FcPatternCreate},
	{"FcPatternDel", &FcPatternDel},
	{"FcPatternDestroy", &FcPatternDestroy},
	{"FcPatternDuplicate", &FcPatternDuplicate},
	{"FcPatternEqual", &FcPatternEqual},
	{"FcPatternEqualSubset", &FcPatternEqualSubset},
	{"FcPatternFilter", &FcPatternFilter},
	{"FcPatternFormat", &FcPatternFormat},
	{"FcPatternGet", &FcPatternGet},
	{"FcPatternGetBool", &FcPatternGetBool},
	{"FcPatternGetCharSet", &FcPatternGetCharSet},
	{"FcPatternGetDouble", &FcPatternGetDouble},
	{"FcPatternGetFTFace", &FcPatternGetFTFace},
	{"FcPatternGetInteger", &FcPatternGetInteger},
	{"FcPatternGetLangSet", &FcPatternGetLangSet},
	{"FcPatternGetMatrix", &FcPatternGetMatrix},
	{"FcPatternGetString", &FcPatternGetString},
	{"FcPatternHash", &FcPatternHash},
	{"FcPatternPrint", &FcPatternPrint},
	{"FcPatternReference", &FcPatternReference},
	{"FcPatternRemove", &FcPatternRemove},
	{"FcPatternVaBuild", &FcPatternVaBuild},
	{"FcStrBasename", &FcStrBasename},
	{"FcStrCmp", &FcStrCmp},
	{"FcStrCmpIgnoreCase", &FcStrCmpIgnoreCase},
	{"FcStrCopy", &FcStrCopy},
	{"FcStrCopyFilename", &FcStrCopyFilename},
	{"FcStrDirname", &FcStrDirname},
	{"FcStrDowncase", &FcStrDowncase},
	{"FcStrFree", &FcStrFree},
	{"FcStrListCreate", &FcStrListCreate},
	{"FcStrListDone", &FcStrListDone},
	{"FcStrListNext", &FcStrListNext},
	{"FcStrPlus", &FcStrPlus},
	{"FcStrSetAdd", &FcStrSetAdd},
	{"FcStrSetAddFilename", &FcStrSetAddFilename},
	{"FcStrSetCreate", &FcStrSetCreate},
	{"FcStrSetDel", &FcStrSetDel},
	{"FcStrSetDestroy", &FcStrSetDestroy},
	{"FcStrSetEqual", &FcStrSetEqual},
	{"FcStrSetMember", &FcStrSetMember},
	{"FcStrStr", &FcStrStr},
	{"FcStrStrIgnoreCase", &FcStrStrIgnoreCase},
	{"FcUcs4ToUtf8", &FcUcs4ToUtf8},
	{"FcUtf16Len", &FcUtf16Len},
	{"FcUtf16ToUcs4", &FcUtf16ToUcs4},
	{"FcUtf8Len", &FcUtf8Len},
	{"FcUtf8ToUcs4", &FcUtf8ToUcs4},
	{"FcValueDestroy", &FcValueDestroy},
	{"FcValueEqual", &FcValueEqual},
	{"FcValuePrint", &FcValuePrint},
	{"FcValueSave", &FcValueSave},
}
