// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package glib provides API definitions for accessing
//libglib-2.0-0.dll and libgthread-2.0-0.dll.
package glib

import (
	"github.com/tHinqa/outside"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllApis(dllThread, false, apiListThread)
	outside.AddDllData(dll, false, dataList)
}

type (
	Enum int

	//TODO(t):Fix (stat/stat32)
	GStatBuf struct{}

	Error struct {
		Domain  Quark
		Code    int
		Message *T.Gchar
	}
)

const (
	G_ASCII_DTOSTR_BUF_SIZE = 29 + 10
	G_STR_DELIMITERS        = "_-|> <."
)

type (
	StrSlice []string
	Str      string
)

func (StrSlice) Dispose(a **T.Gchar) { Strfreev(a) }

func (Str) Dispose(a T.Gpointer) { Free(a) }

var (
	InternString func(str string) string

	InternStaticString func(str string) string

	ErrorNew func(domain Quark, code int, format string,
		v ...VArg) *Error

	ErrorNewLiteral func(domain Quark,
		code int,
		message string) *Error

	ErrorNewValist func(domain Quark,
		code int,
		format string,
		args VAList) *Error

	ErrorFree func(e *Error)

	ErrorCopy func(e *Error) *Error

	ErrorMatches func(e *Error,
		domain Quark,
		code int) bool

	SetError func(err **Error, domain Quark, code int,
		format string, v ...VArg)

	SetErrorLiteral func(err **Error,
		domain Quark,
		code int,
		message string)

	PropagateError func(dest **Error,
		src *Error)

	ClearError func(err **Error)

	PrefixError func(err **Error, format string, v ...VArg)

	PropagatePrefixedError func(dest **Error, src *Error,
		format string, v ...VArg)

	GetUserName func() string

	GetRealName func() string

	GetHomeDir func() string

	GetTmpDir func() string

	GetHostName func() string

	GetPrgname func() string

	SetPrgname func(prgname string)

	GetApplicationName func() string

	SetApplicationName func(applicationName string)

	ReloadUserSpecialDirsCache func()

	GetUserDataDir func() string

	GetUserConfigDir func() string

	GetUserCacheDir func() string

	GetSystemDataDirs func() []string

	Win32GetSystemDataDirsForModule func(f func()) []string

	ParseDebugString func(
		str string, keys *T.GDebugKey, nkeys uint) uint

	Snprintf func(
		str T.Gchar, n T.Gulong, format string, v ...VArg) int

	Vsnprintf func(str T.Gchar,
		n T.Gulong, format string, args VAList) int

	PathIsAbsolute func(fileName string) bool

	PathSkipRoot func(fileName string) string

	Basename func(fileName string) string

	GetCurrentDir func() string

	PathGetBasename func(fileName string) string

	PathGetDirname func(fileName string) string

	NullifyPointer func(nullifyLocation *T.Gpointer)

	Getenv func(variable string) string

	Setenv func(
		variable, value string, overwrite bool) bool

	Unsetenv func(variable string)

	Listenv func() StrSlice

	GetEnviron func() []string //TODO(t):Documented?

	StaticMutexGetMutexImpl func(mutex **Mutex) *Mutex

	Usleep func(microseconds T.Gulong)

	OnErrorQuery func(prgName string)

	OnErrorStackTrace func(prgName string)

	Base64EncodeStep func(in *T.Guchar, leng T.Gsize,
		breakLines bool, out string, state, save *int) T.Gsize

	Base64EncodeClose func(breakLines bool, out string,
		state, save *int) T.Gsize

	Base64Encode func(data *T.Guchar,
		leng T.Gsize) string

	Base64DecodeStep func(in string,
		leng T.Gsize,
		out *T.Guchar,
		state *int,
		save *uint) T.Gsize

	Base64Decode func(text string, outLen *T.Gsize) *T.Guchar

	Base64DecodeInplace func(
		text string, outLen *T.Gsize) *T.Guchar

	BitLock func(address *int, lockBit int)

	BitTrylock func(address *int, lockBit int) bool

	BitUnlock func(address *int, lockBit int)

	GetSystemConfigDirs func() []string

	GetUserRuntimeDir func() string

	GetLanguageNames func() []string

	GetLocaleVariants func(locale string) []string

	Atexit func(f T.GVoidFunc)

	GlibCheckVersion func(
		requiredMajor, requiredMinor, requiredMicro uint) string

	Free func(mem T.Gpointer)

	Malloc func(nBytes T.Gsize) T.Gpointer

	Malloc0 func(nBytes T.Gsize) T.Gpointer

	Realloc func(mem T.Gpointer,
		nBytes T.Gsize) T.Gpointer

	TryMalloc func(nBytes T.Gsize) T.Gpointer

	TryMalloc0 func(nBytes T.Gsize) T.Gpointer

	TryRealloc func(mem T.Gpointer, nBytes T.Gsize) T.Gpointer

	MallocN func(nBlocks T.Gsize, nBlockBytes T.Gsize) T.Gpointer

	Malloc0N func(nBlocks T.Gsize, nBlockBytes T.Gsize) T.Gpointer

	ReallocN func(mem T.Gpointer,
		nBlocks T.Gsize, nBlockBytes T.Gsize) T.Gpointer

	TryMallocN func(nBlocks T.Gsize, nBlockBytes T.Gsize) T.Gpointer

	TryMalloc0N func(nBlocks T.Gsize, nBlockBytes T.Gsize) T.Gpointer

	TryReallocN func(mem T.Gpointer,
		nBlocks T.Gsize, nBlockBytes T.Gsize) T.Gpointer

	MemIsSystemMalloc func() bool

	MemProfile func()

	BlowChunks func()

	ConvertErrorQuark func() Quark

	Convert func(str string, leng T.Gssize,
		toCodeset, fromCodeset string,
		bytesRead, bytesWritten *T.Gsize,
		e **Error) string

	ConvertWithIconv func(str string, leng T.Gssize,
		converter IConv,
		bytesRead, bytesWritten *T.Gsize, e **Error) string

	ConvertWithFallback func(str string, leng T.Gssize,
		toCodeset, fromCodeset, fallback string,
		bytesRead, bytesWritten *T.Gsize, e **Error) string

	LocaleToUtf8 func(opsysstr string, leng T.Gssize,
		bytesRead, bytesWritten *T.Gsize, e **Error) string

	LocaleFromUtf8 func(utf8str string, leng T.Gssize,
		bytesRead, bytesWritten *T.Gsize, e **Error) string

	FilenameToUtf8 func(opsysstr string, leng T.Gssize,
		bytesRead, bytesWritten *T.Gsize, e **Error) string

	FilenameFromUtf8 func(utf8str string, leng T.Gssize,
		bytesRead, bytesWritten *T.Gsize, e **Error) string

	FilenameFromUri func(uri string,
		hostname **T.Gchar, e **Error) string

	FilenameToUri func(filename, hostname string,
		e **Error) string

	FilenameDisplayName func(filename string) string

	GetFilenameCharsets func(charsets ***T.Gchar) bool

	FilenameDisplayBasename func(filename string) string

	UriListExtractUris func(uriList string) []string

	DatalistInit func(datalist **T.GData)

	DatalistClear func(datalist **T.GData)

	DatalistIdGetData func(datalist **T.GData,
		keyId Quark) T.Gpointer

	DatalistIdSetDataFull func(datalist **T.GData,
		keyId Quark,
		data T.Gpointer,
		destroyFunc O.DestroyNotify)

	DatalistIdRemoveNoNotify func(datalist **T.GData,
		keyId Quark) T.Gpointer

	DatalistForeach func(datalist **T.GData,
		f T.GDataForeachFunc,
		userData T.Gpointer)

	DatalistSetFlags func(datalist **T.GData, flags uint)

	DatalistUnsetFlags func(datalist **T.GData, flags uint)

	DatalistGetFlags func(datalist **T.GData) uint

	DatasetDestroy func(datasetLocation T.Gconstpointer)

	DatasetIdGetData func(datasetLocation T.Gconstpointer,
		keyId Quark) T.Gpointer

	DatasetIdSetDataFull func(datasetLocation T.Gconstpointer,
		keyId Quark, data T.Gpointer,
		destroyFunc O.DestroyNotify)

	DatasetIdRemoveNoNotify func(datasetLocation T.Gconstpointer,
		keyId Quark) T.Gpointer

	DatasetForeach func(datasetLocation T.Gconstpointer,
		f T.GDataForeachFunc, userData T.Gpointer)

	FileErrorQuark func() Quark

	FileErrorFromErrno func(errNo int) T.GFileError

	FileTest func(filename string, test T.GFileTest) bool

	FileGetContents func(filename string, contents **T.Gchar,
		length *T.Gsize, e **Error) bool

	FileSetContents func(filename, contents string,
		length T.Gssize, e **Error) bool

	FileReadLink func(filename string, e **Error) string

	Mkstemp func(tmpl string) int

	MkstempFull func(tmpl string, flags, mode int) int

	FileOpenTmp func(tmpl string, nameUsed **T.Gchar,
		e **Error) int

	FormatSizeForDisplay func(size T.Goffset) string

	BuildPath func(separator, firstElement string,
		v ...VArg) string

	BuildPathv func(separator string, args []string) string

	BuildFilename func(firstElement string, v ...VArg) string

	BuildFilenamev func(args []string) string

	MkdirWithParents func(pathname string, mode int) int

	StrEqual func(v1, v2 T.Gconstpointer) bool

	StrHash func(v T.Gconstpointer) uint

	IntEqual func(v1, v2 T.Gconstpointer) bool

	IntHash func(v T.Gconstpointer) uint

	Int64Equal func(v1, v2 T.Gconstpointer) bool

	Int64Hash func(v T.Gconstpointer) uint

	DoubleEqual func(v1, v2 T.Gconstpointer) bool

	DoubleHash func(v T.Gconstpointer) uint

	DirectHash func(v T.Gconstpointer) uint

	DirectEqual func(v1, v2 T.Gconstpointer) bool

	HostnameIsNonAscii func(hostname string) bool

	HostnameIsAsciiEncoded func(hostname string) bool

	HostnameIsIpAddress func(hostname string) bool

	HostnameToAscii func(hostname string) string

	HostnameToUnicode func(hostname string) string

	Poll func(fds *T.GPollFD, nfds uint, timeout int) int

	IdleSourceNew func() *O.Source

	ChildWatchSourceNew func(pid T.GPid) *O.Source

	GetCurrentTime func(result *TimeVal)

	GetMonotonicTime func() int64

	GetRealTime func() int64

	TimeoutSourceNew func(interval uint) *O.Source

	TimeoutSourceNewSeconds func(interval uint) *O.Source

	TimeoutAddFull func(priority int,
		interval uint,
		function O.SourceFunc,
		data T.Gpointer,
		notify O.DestroyNotify) uint

	TimeoutAdd func(interval uint,
		function O.SourceFunc, data T.Gpointer) uint

	TimeoutAddSecondsFull func(priority int,
		interval uint,
		function O.SourceFunc,
		data T.Gpointer,
		notify O.DestroyNotify) uint

	TimeoutAddSeconds func(interval uint,
		function O.SourceFunc,
		data T.Gpointer) uint

	ChildWatchAddFull func(priority int,
		pid T.GPid,
		function T.GChildWatchFunc,
		data T.Gpointer,
		notify O.DestroyNotify) uint

	ChildWatchAdd func(pid T.GPid,
		function T.GChildWatchFunc,
		data T.Gpointer) uint

	IdleAdd func(function O.SourceFunc,
		data T.Gpointer) uint

	IdleAddFull func(priority int,
		function O.SourceFunc,
		data T.Gpointer,
		notify O.DestroyNotify) uint

	IdleRemoveByData func(data T.Gpointer) bool

	GetCharset func(charset **T.Char) bool

	Utf8GetChar func(p string) Unichar

	Utf8GetCharValidated func(p string,
		maxLen T.Gssize) Unichar

	Utf8OffsetToPointer func(str string,
		offset T.Glong) string

	Utf8PointerToOffset func(str string,
		pos string) T.Glong

	Utf8PrevChar func(p string) string

	Utf8FindNextChar func(p string,
		end string) string

	Utf8FindPrevChar func(str string,
		p string) string

	Utf8Strlen func(p string, max T.Gssize) T.Glong

	Utf8Strncpy func(dest, src string, n T.Gsize) string

	Utf8Strchr func(p string, leng T.Gssize, c Unichar) string

	Utf8Strrchr func(p string, leng T.Gssize, c Unichar) string

	Utf8Strreverse func(str string, leng T.Gssize) string

	Utf8ToUtf16 func(str string, leng T.Glong,
		itemsRead, itemsWritten *T.Glong,
		e **Error) *T.Gunichar2

	Utf8ToUcs4 func(str string, leng T.Glong,
		itemsRead, itemsWritten *T.Glong,
		e **Error) *Unichar

	Utf8ToUcs4Fast func(str string, leng T.Glong,
		itemsWritten *T.Glong) *Unichar

	Utf16ToUcs4 func(str *T.Gunichar2, leng T.Glong,
		itemsRead, itemsWritten *T.Glong,
		e **Error) *Unichar

	Utf16ToUtf8 func(str *T.Gunichar2, leng T.Glong,
		itemsRead, itemsWritten *T.Glong,
		e **Error) string

	Ucs4ToUtf16 func(str *Unichar, leng T.Glong,
		itemsRead, itemsWritten *T.Glong,
		e **Error) *T.Gunichar2

	Ucs4ToUtf8 func(str *Unichar, leng T.Glong,
		itemsRead, itemsWritten *T.Glong,
		e **Error) string

	Utf8Validate func(str string, maxLen T.Gssize,
		end **T.Gchar) bool

	Utf8Strup func(str string, leng T.Gssize) string

	Utf8Strdown func(str string, leng T.Gssize) string

	Utf8Casefold func(str string, leng T.Gssize) string

	Utf8Normalize func(str string,
		leng T.Gssize,
		mode T.GNormalizeMode) string

	Utf8Collate func(str1 string, str2 string) int

	Utf8CollateKey func(str string, leng T.Gssize) string

	Utf8CollateKeyForFilename func(str string,
		leng T.Gssize) string

	LogSetHandler func(logDomain string,
		logLevels T.GLogLevelFlags,
		logFunc T.GLogFunc,
		userData T.Gpointer) uint

	LogRemoveHandler func(logDomain string,
		handlerId uint)

	LogDefaultHandler func(logDomain string,
		logLevel T.GLogLevelFlags,
		message string,
		unusedData T.Gpointer)

	LogSetDefaultHandler func(logFunc T.GLogFunc,
		userData T.Gpointer) T.GLogFunc

	Log func(logDomain string, logLevel T.GLogLevelFlags,
		format string, v ...VArg)

	Logv func(logDomain string,
		logLevel T.GLogLevelFlags,
		format string,
		args VAList)

	LogSetFatalMask func(logDomain string,
		fatalMask T.GLogLevelFlags) T.GLogLevelFlags

	LogSetAlwaysFatal func(fatalMask T.GLogLevelFlags) T.GLogLevelFlags

	ReturnIfFailWarning func(logDomain string,
		prettyFunction string,
		expression string)

	WarnMessage func(domain string,
		file string,
		line int,
		f string,
		warnexpr string)

	AssertWarning func(logDomain string,
		file string,
		line int,
		prettyFunction string,
		expression string)

	Print func(format string, v ...VArg)

	SetPrintHandler func(f T.GPrintFunc) T.GPrintFunc

	Printerr func(format string, v ...VArg)

	SetPrinterrHandler func(f T.GPrintFunc) T.GPrintFunc

	NodeNew func(data T.Gpointer) *T.GNode

	NodeDestroy func(root *T.GNode)

	NodeUnlink func(node *T.GNode)

	NodeCopyDeep func(node *T.GNode,
		copyFunc T.GCopyFunc, data T.Gpointer) *T.GNode

	NodeCopy func(node *T.GNode) *T.GNode

	NodeInsert func(parent *T.GNode,
		position int, node *T.GNode) *T.GNode

	NodeInsertBefore func(parent *T.GNode,
		sibling *T.GNode, node *T.GNode) *T.GNode

	NodeInsertAfter func(parent *T.GNode,
		sibling *T.GNode, node *T.GNode) *T.GNode

	NodePrepend func(parent *T.GNode, node *T.GNode) *T.GNode

	NodeNNodes func(root *T.GNode,
		flags TraverseFlags) uint

	NodeGetRoot func(node *T.GNode) *T.GNode

	NodeIsAncestor func(node *T.GNode,
		descendant *T.GNode) bool

	NodeDepth func(node *T.GNode) uint

	NodeFind func(root *T.GNode,
		order TraverseType,
		flags TraverseFlags,
		data T.Gpointer) *T.GNode

	NodeTraverse func(root *T.GNode,
		order TraverseType,
		flags TraverseFlags,
		maxDepth int,
		f T.GNodeTraverseFunc,
		data T.Gpointer)

	NodeMaxHeight func(root *T.GNode) uint

	NodeChildrenForeach func(node *T.GNode,
		flags TraverseFlags,
		f T.GNodeForeachFunc,
		data T.Gpointer)

	NodeReverseChildren func(node *T.GNode)

	NodeNChildren func(node *T.GNode) uint

	NodeNthChild func(node *T.GNode,
		n uint) *T.GNode

	NodeLastChild func(node *T.GNode) *T.GNode

	NodeFindChild func(node *T.GNode,
		flags TraverseFlags,
		data T.Gpointer) *T.GNode

	NodeChildPosition func(node *T.GNode,
		child *T.GNode) int

	NodeChildIndex func(node *T.GNode,
		data T.Gpointer) int

	NodeFirstSibling func(node *T.GNode) *T.GNode

	NodeLastSibling func(node *T.GNode) *T.GNode

	NodePushAllocator func(dummy T.Gpointer)

	NodePopAllocator func()

	SpacedPrimesClosest func(num uint) uint

	QsortWithData func(pbase T.Gconstpointer,
		totalElems int,
		size T.Gsize,
		compareFunc T.GCompareDataFunc,
		userData T.Gpointer)

	MatchInfoGetRegex func(matchInfo *T.GMatchInfo) *Regex

	MatchInfoGetString func(matchInfo *T.GMatchInfo) string

	MatchInfoFree func(matchInfo *T.GMatchInfo)

	MatchInfoNext func(matchInfo *T.GMatchInfo,
		e **Error) bool

	MatchInfoMatches func(matchInfo *T.GMatchInfo) bool

	MatchInfoGetMatchCount func(matchInfo *T.GMatchInfo) int

	MatchInfoIsPartialMatch func(matchInfo *T.GMatchInfo) bool

	MatchInfoExpandReferences func(matchInfo *T.GMatchInfo,
		stringToExpand string,
		e **Error) string

	MatchInfoFetch func(matchInfo *T.GMatchInfo,
		matchNum int) string

	MatchInfoFetchPos func(matchInfo *T.GMatchInfo,
		matchNum int,
		startPos *int,
		endPos *int) bool

	MatchInfoFetchNamed func(matchInfo *T.GMatchInfo,
		name string) string

	MatchInfoFetchNamedPos func(matchInfo *T.GMatchInfo,
		name string, startPos *int, endPos *int) bool

	MatchInfoFetchAll func(
		matchInfo *T.GMatchInfo) []string
	//
	AsciiTolower func(c T.Gchar) T.Gchar

	AsciiToupper func(c T.Gchar) T.Gchar

	AsciiDigitValue func(c T.Gchar) int

	AsciiXdigitValue func(c T.Gchar) int

	Strdelimit func( // in place
		str, delimiters string, newDelimiter T.Gchar) string

	Strcanon func( // in place
		str, validChars string, substitutor T.Gchar) string

	Strerror func(errnum int) string

	Strsignal func(signum int) string

	Strreverse func(str string) string // in place

	Strlcpy func(
		dest *T.Gchar, src string, destSize T.Gsize) T.Gsize

	Strlcat func(
		dest T.Gchar, src string, destSize T.Gsize) T.Gsize

	StrstrLen func(haystack string,
		haystackLen T.Gssize, needle string) *T.Gchar

	Strrstr func(haystack string, needle string) *T.Gchar

	StrrstrLen func(haystack string,
		haystackLen T.Gssize, needle string) *T.Gchar

	StrHasSuffix func(str string, suffix string) bool

	StrHasPrefix func(str string, prefix string) bool

	Strtod func(nptr string, endptr **T.Gchar) float64

	AsciiStrtod func(nptr string, endptr **T.Gchar) float64

	AsciiStrtoull func(
		nptr string, endptr **T.Gchar, base uint) uint64

	AsciiStrtoll func(
		nptr string, endptr **T.Gchar, base uint) int64

	// suggested bufLen G_ASCII_DTOSTR_BUF_SIZE
	AsciiDtostr func(
		buffer *T.Gchar, bufLen int, d float64) string

	AsciiFormatd func(buffer *T.Gchar,
		bufLen int, format string, d float64) string

	Strchug func(str string) string // in place

	Strchomp func(str string) string // in place

	// Strstrip(s) == Strchomp(Strchug(s))

	AsciiStrcasecmp func(s1 string, s2 string) int

	AsciiStrncasecmp func(s1 string, s2 string, n T.Gsize) int

	AsciiStrdown func(str string, leng T.Gssize) Str

	AsciiStrup func(str string, leng T.Gssize) Str

	//Deprecated 2.2
	// Strcasecmp func(s1 string, s2 string) int
	// Strncasecmp func(s1 string, s2 string, n uint) int
	// Strdown func(str string) Str
	// Strup func(str string) Str

	Strdup func(str string) Str

	StrdupPrintf func(format string, v ...VArg) Str

	StrdupVprintf func(format string, args VAList) Str

	Strndup func(str string, n T.Gsize) Str

	Strnfill func(length T.Gsize, fillChar T.Gchar) Str

	Strconcat func(string1 string, v ...VArg) Str

	Strjoin func(separator string, v ...VArg) Str

	Strcompress func(source string) Str

	Strescape func(source, exceptions string) Str

	Memdup func(mem T.Gconstpointer, byteSize uint) T.Gpointer

	Strsplit func(str, delimiter string, maxTokens int) StrSlice

	StrsplitSet func(
		str, delimiters string, maxTokens int) StrSlice

	Strjoinv func(separator string, strArray []string) Str

	Strfreev func(strArray **T.Gchar) // left as is for Dispose

	Strdupv func(strArray []string) StrSlice

	StrvLength func(strArray **T.Gchar) uint // []string has len()

	Stpcpy func(dest *T.Gchar, src string) *T.Gchar

	StripContext func(msgid, msgval string) string

	Dgettext func(domain, msgid string) string

	Dcgettext func(domain, msgid string, category int) string

	Dngettext func(
		domain, msgid, msgidPlural string, n T.Gulong) string

	Dpgettext func(
		domain, msgctxtid string, msgidoffset T.Gsize) string

	Dpgettext2 func(domain, context, msgid string) string
	//
	Strcmp0 func(str1, str2 string) int

	UriUnescapeString func(escapedString,
		illegalCharacters string) string

	UriUnescapeSegment func(escapedString,
		escapedStringEnd,
		illegalCharacters string) string

	UriParseScheme func(uri string) string

	UriEscapeString func(unescaped,
		reservedCharsAllowed string,
		allowUtf8 bool) string

	Win32Ftruncate func(f int,
		size uint) int

	Win32Getlocale func() string

	Win32ErrorMessage func(err int) string

	Win32GetPackageInstallationDirectory func(
		pkg, dllName string) string

	Win32GetPackageInstallationSubdirectory func(
		pkg, dllName, subdir string) string

	Win32GetPackageInstallationDirectoryOfModule func(
		hmodule T.Gpointer) string

	Win32GetWindowsVersion func() uint

	Win32LocaleFilenameFromUtf8 func(
		utf8filename string) string

	Access func(filename string, mode int) int

	Chmod func(filename string, mode int) int

	Open func(filename string, flags int, mode int) int

	Creat func(filename string, mode int) int

	Rename func(oldfilename, newfilename string) int

	Mkdir func(filename string, mode int) int

	Chdir func(path string) int

	Stat func(filename string, buf *GStatBuf) int

	Lstat func(filename string, buf *GStatBuf) int

	Unlink func(filename string) int

	Remove func(filename string) int

	Rmdir func(filename string) int

	Fopen func(filename, mode string) *T.FILE

	Freopen func(
		filename, mode string, stream *T.FILE) *T.FILE

	Utime func(filename string, utb *T.Utimbuf) int

	FindProgramInPath func(program string) string

	Printf func(format string, v ...VArg) int

	Fprintf func(file *T.FILE, format string, v ...VArg) int

	Sprintf func(str *T.Gchar, format string, v ...VArg) int

	Vprintf func(format string, args VAList) int

	Vfprintf func(
		file *T.FILE, format string, args VAList) int

	Vsprintf func(
		str T.Gchar, format string, args VAList) int

	Vasprintf func(
		str **T.Gchar, format string, args VAList) int

	GetUserSpecialDir func(directory T.GUserDirectory) string

	MarkupCollectAttributes func(elementName string,
		attributeNames []string, attributeValues []string,
		err **Error, firstType T.GMarkupCollectType,
		firstAttr string, v ...VArg) bool

	MarkupErrorQuark func() Quark

	MemSetVtable func(vtable *T.GMemVTable)

	OptionErrorQuark func() Quark

	PrintfStringUpperBound func(
		format string, args VAList) T.Gsize

	RegexErrorQuark func() Quark
)

var dll = "libglib-2.0-0.dll"

var apiList = outside.Apis{
	// Undocumented {"_g_debug_flags", &_gDebugFlags},
	// Undocumented {"_g_debug_initialized", &_gDebugInitialized},
	{"g_access", &Access},
	{"g_allocator_free", &AllocatorFree},
	{"g_allocator_new", &AllocatorNew},
	{"g_array_append_vals", &ArrayAppendVals},
	{"g_array_free", &ArrayFree},
	{"g_array_get_element_size", &ArrayGetElementSize},
	{"g_array_insert_vals", &ArrayInsertVals},
	{"g_array_new", &ArrayNew},
	{"g_array_prepend_vals", &ArrayPrependVals},
	{"g_array_ref", &ArrayRef},
	{"g_array_remove_index", &ArrayRemoveIndex},
	{"g_array_remove_index_fast", &ArrayRemoveIndexFast},
	{"g_array_remove_range", &ArrayRemoveRange},
	{"g_array_set_size", &ArraySetSize},
	{"g_array_sized_new", &ArraySizedNew},
	{"g_array_sort", &ArraySort},
	{"g_array_sort_with_data", &ArraySortWithData},
	{"g_array_unref", &ArrayUnref},
	{"g_ascii_digit_value", &AsciiDigitValue},
	{"g_ascii_dtostr", &AsciiDtostr},
	{"g_ascii_formatd", &AsciiFormatd},
	{"g_ascii_strcasecmp", &AsciiStrcasecmp},
	{"g_ascii_strdown", &AsciiStrdown},
	{"g_ascii_strncasecmp", &AsciiStrncasecmp},
	{"g_ascii_strtod", &AsciiStrtod},
	{"g_ascii_strtoll", &AsciiStrtoll},
	{"g_ascii_strtoull", &AsciiStrtoull},
	{"g_ascii_strup", &AsciiStrup},
	{"g_ascii_tolower", &AsciiTolower},
	{"g_ascii_toupper", &AsciiToupper},
	{"g_ascii_xdigit_value", &AsciiXdigitValue},
	{"g_assert_warning", &AssertWarning},
	{"g_assertion_message", &AssertionMessage},
	{"g_assertion_message_cmpnum", &AssertionMessageCmpnum},
	{"g_assertion_message_cmpstr", &AssertionMessageCmpstr},
	{"g_assertion_message_error", &AssertionMessageError},
	{"g_assertion_message_expr", &AssertionMessageExpr},
	{"g_async_queue_length", &AsyncQueueLength},
	{"g_async_queue_length_unlocked", &AsyncQueueLengthUnlocked},
	{"g_async_queue_lock", &AsyncQueueLock},
	{"g_async_queue_new", &AsyncQueueNew},
	{"g_async_queue_new_full", &AsyncQueueNewFull},
	{"g_async_queue_pop", &AsyncQueuePop},
	{"g_async_queue_pop_unlocked", &AsyncQueuePopUnlocked},
	{"g_async_queue_push", &AsyncQueuePush},
	{"g_async_queue_push_sorted", &AsyncQueuePushSorted},
	{"g_async_queue_push_sorted_unlocked", &AsyncQueuePushSortedUnlocked},
	{"g_async_queue_push_unlocked", &AsyncQueuePushUnlocked},
	{"g_async_queue_ref", &AsyncQueueRef},
	{"g_async_queue_ref_unlocked", &AsyncQueueRefUnlocked},
	{"g_async_queue_sort", &AsyncQueueSort},
	{"g_async_queue_sort_unlocked", &AsyncQueueSortUnlocked},
	{"g_async_queue_timed_pop", &AsyncQueueTimedPop},
	{"g_async_queue_timed_pop_unlocked", &AsyncQueueTimedPopUnlocked},
	{"g_async_queue_try_pop", &AsyncQueueTryPop},
	{"g_async_queue_try_pop_unlocked", &AsyncQueueTryPopUnlocked},
	{"g_async_queue_unlock", &AsyncQueueUnlock},
	{"g_async_queue_unref", &AsyncQueueUnref},
	{"g_async_queue_unref_and_unlock", &AsyncQueueUnrefAndUnlock},
	{"g_atexit", &Atexit},
	{"g_atomic_int_add", &AtomicIntAdd},
	{"g_atomic_int_compare_and_exchange", &AtomicIntCompareAndExchange},
	{"g_atomic_int_exchange_and_add", &AtomicIntExchangeAndAdd},
	{"g_atomic_int_get", &AtomicIntGet},
	{"g_atomic_int_set", &AtomicIntSet},
	{"g_atomic_pointer_compare_and_exchange", &AtomicPointerCompareAndExchange},
	{"g_atomic_pointer_get", &AtomicPointerGet},
	{"g_atomic_pointer_set", &AtomicPointerSet},
	{"g_base64_decode", &Base64Decode},
	{"g_base64_decode_inplace", &Base64DecodeInplace},
	{"g_base64_decode_step", &Base64DecodeStep},
	{"g_base64_encode", &Base64Encode},
	{"g_base64_encode_close", &Base64EncodeClose},
	{"g_base64_encode_step", &Base64EncodeStep},
	{"g_basename", &Basename},
	{"g_bit_lock", &BitLock},
	// Inline {"g_bit_nth_lsf", &BitNthLsf},
	// Inline {"g_bit_nth_msf", &BitNthMsf},
	// Inline {"g_bit_storage", &BitStorage},
	{"g_bit_trylock", &BitTrylock},
	{"g_bit_unlock", &BitUnlock},
	{"g_blow_chunks", &BlowChunks},
	{"g_bookmark_file_add_application", &BookmarkFileAddApplication},
	{"g_bookmark_file_add_group", &BookmarkFileAddGroup},
	{"g_bookmark_file_error_quark", &BookmarkFileErrorQuark},
	{"g_bookmark_file_free", &BookmarkFileFree},
	{"g_bookmark_file_get_added", &BookmarkFileGetAdded},
	{"g_bookmark_file_get_app_info", &BookmarkFileGetAppInfo},
	{"g_bookmark_file_get_applications", &BookmarkFileGetApplications},
	{"g_bookmark_file_get_description", &BookmarkFileGetDescription},
	{"g_bookmark_file_get_groups", &BookmarkFileGetGroups},
	{"g_bookmark_file_get_icon", &BookmarkFileGetIcon},
	{"g_bookmark_file_get_is_private", &BookmarkFileGetIsPrivate},
	{"g_bookmark_file_get_mime_type", &BookmarkFileGetMimeType},
	{"g_bookmark_file_get_modified", &BookmarkFileGetModified},
	{"g_bookmark_file_get_size", &BookmarkFileGetSize},
	{"g_bookmark_file_get_title", &BookmarkFileGetTitle},
	{"g_bookmark_file_get_uris", &BookmarkFileGetUris},
	{"g_bookmark_file_get_visited", &BookmarkFileGetVisited},
	{"g_bookmark_file_has_application", &BookmarkFileHasApplication},
	{"g_bookmark_file_has_group", &BookmarkFileHasGroup},
	{"g_bookmark_file_has_item", &BookmarkFileHasItem},
	{"g_bookmark_file_load_from_data", &BookmarkFileLoadFromData},
	{"g_bookmark_file_load_from_data_dirs", &BookmarkFileLoadFromDataDirs},
	{"g_bookmark_file_load_from_file", &BookmarkFileLoadFromFile},
	{"g_bookmark_file_move_item", &BookmarkFileMoveItem},
	{"g_bookmark_file_new", &BookmarkFileNew},
	{"g_bookmark_file_remove_application", &BookmarkFileRemoveApplication},
	{"g_bookmark_file_remove_group", &BookmarkFileRemoveGroup},
	{"g_bookmark_file_remove_item", &BookmarkFileRemoveItem},
	{"g_bookmark_file_set_added", &BookmarkFileSetAdded},
	{"g_bookmark_file_set_app_info", &BookmarkFileSetAppInfo},
	{"g_bookmark_file_set_description", &BookmarkFileSetDescription},
	{"g_bookmark_file_set_groups", &BookmarkFileSetGroups},
	{"g_bookmark_file_set_icon", &BookmarkFileSetIcon},
	{"g_bookmark_file_set_is_private", &BookmarkFileSetIsPrivate},
	{"g_bookmark_file_set_mime_type", &BookmarkFileSetMimeType},
	{"g_bookmark_file_set_modified", &BookmarkFileSetModified},
	{"g_bookmark_file_set_title", &BookmarkFileSetTitle},
	{"g_bookmark_file_set_visited", &BookmarkFileSetVisited},
	{"g_bookmark_file_to_data", &BookmarkFileToData},
	{"g_bookmark_file_to_file", &BookmarkFileToFile},
	{"g_build_filename", &BuildFilename},
	{"g_build_filenamev", &BuildFilenamev},
	{"g_build_path", &BuildPath},
	{"g_build_pathv", &BuildPathv},
	{"g_byte_array_append", &ByteArrayAppend},
	{"g_byte_array_free", &ByteArrayFree},
	{"g_byte_array_new", &ByteArrayNew},
	{"g_byte_array_prepend", &ByteArrayPrepend},
	{"g_byte_array_ref", &ByteArrayRef},
	{"g_byte_array_remove_index", &ByteArrayRemoveIndex},
	{"g_byte_array_remove_index_fast", &ByteArrayRemoveIndexFast},
	{"g_byte_array_remove_range", &ByteArrayRemoveRange},
	{"g_byte_array_set_size", &ByteArraySetSize},
	{"g_byte_array_sized_new", &ByteArraySizedNew},
	{"g_byte_array_sort", &ByteArraySort},
	{"g_byte_array_sort_with_data", &ByteArraySortWithData},
	{"g_byte_array_unref", &ByteArrayUnref},
	{"g_cache_destroy", &CacheDestroy},
	{"g_cache_insert", &CacheInsert},
	{"g_cache_key_foreach", &CacheKeyForeach},
	{"g_cache_new", &CacheNew},
	{"g_cache_remove", &CacheRemove},
	{"g_cache_value_foreach", &CacheValueForeach},
	{"g_chdir", &Chdir},
	{"g_checksum_copy", &ChecksumCopy},
	{"g_checksum_free", &ChecksumFree},
	{"g_checksum_get_digest", &ChecksumGetDigest},
	{"g_checksum_get_string", &ChecksumGetString},
	{"g_checksum_new", &ChecksumNew},
	{"g_checksum_reset", &ChecksumReset},
	{"g_checksum_type_get_length", &ChecksumTypeGetLength},
	{"g_checksum_update", &ChecksumUpdate},
	{"g_child_watch_add", &ChildWatchAdd},
	{"g_child_watch_add_full", &ChildWatchAddFull},
	{"g_child_watch_source_new", &ChildWatchSourceNew},
	{"g_chmod", &Chmod},
	{"g_clear_error", &ClearError},
	{"g_completion_add_items", &CompletionAddItems},
	{"g_completion_clear_items", &CompletionClearItems},
	{"g_completion_complete", &CompletionComplete},
	{"g_completion_complete_utf8", &CompletionCompleteUtf8},
	{"g_completion_free", &CompletionFree},
	{"g_completion_new", &CompletionNew},
	{"g_completion_remove_items", &CompletionRemoveItems},
	{"g_completion_set_compare", &CompletionSetCompare},
	{"g_compute_checksum_for_data", &ComputeChecksumForData},
	{"g_compute_checksum_for_string", &ComputeChecksumForString},
	{"g_convert", &Convert},
	{"g_convert_error_quark", &ConvertErrorQuark},
	{"g_convert_with_fallback", &ConvertWithFallback},
	{"g_convert_with_iconv", &ConvertWithIconv},
	{"g_creat", &Creat},
	{"g_datalist_clear", &DatalistClear},
	{"g_datalist_foreach", &DatalistForeach},
	{"g_datalist_get_flags", &DatalistGetFlags},
	{"g_datalist_id_get_data", &DatalistIdGetData},
	{"g_datalist_id_remove_no_notify", &DatalistIdRemoveNoNotify},
	{"g_datalist_id_set_data_full", &DatalistIdSetDataFull},
	{"g_datalist_init", &DatalistInit},
	{"g_datalist_set_flags", &DatalistSetFlags},
	{"g_datalist_unset_flags", &DatalistUnsetFlags},
	{"g_dataset_destroy", &DatasetDestroy},
	{"g_dataset_foreach", &DatasetForeach},
	{"g_dataset_id_get_data", &DatasetIdGetData},
	{"g_dataset_id_remove_no_notify", &DatasetIdRemoveNoNotify},
	{"g_dataset_id_set_data_full", &DatasetIdSetDataFull},
	{"g_date_add_days", &DateAddDays},
	{"g_date_add_months", &DateAddMonths},
	{"g_date_add_years", &DateAddYears},
	{"g_date_clamp", &DateClamp},
	{"g_date_clear", &DateClear},
	{"g_date_compare", &DateCompare},
	{"g_date_days_between", &DateDaysBetween},
	{"g_date_free", &DateFree},
	{"g_date_get_day", &DateGetDay},
	{"g_date_get_day_of_year", &DateGetDayOfYear},
	{"g_date_get_days_in_month", &DateGetDaysInMonth},
	{"g_date_get_iso8601_week_of_year", &DateGetIso8601WeekOfYear},
	{"g_date_get_julian", &DateGetJulian},
	{"g_date_get_monday_week_of_year", &DateGetMondayWeekOfYear},
	{"g_date_get_monday_weeks_in_year", &DateGetMondayWeeksInYear},
	{"g_date_get_month", &DateGetMonth},
	{"g_date_get_sunday_week_of_year", &DateGetSundayWeekOfYear},
	{"g_date_get_sunday_weeks_in_year", &DateGetSundayWeeksInYear},
	{"g_date_get_weekday", &DateGetWeekday},
	{"g_date_get_year", &DateGetYear},
	{"g_date_is_first_of_month", &DateIsFirstOfMonth},
	{"g_date_is_last_of_month", &DateIsLastOfMonth},
	{"g_date_is_leap_year", &DateIsLeapYear},
	{"g_date_new", &DateNew},
	{"g_date_new_dmy", &DateNewDmy},
	{"g_date_new_julian", &DateNewJulian},
	{"g_date_order", &DateOrder},
	{"g_date_set_day", &DateSetDay},
	{"g_date_set_dmy", &DateSetDmy},
	{"g_date_set_julian", &DateSetJulian},
	{"g_date_set_month", &DateSetMonth},
	{"g_date_set_parse", &DateSetParse},
	{"g_date_set_time", &DateSetTime},
	{"g_date_set_time_t", &DateSetTimeT},
	{"g_date_set_time_val", &DateSetTimeVal},
	{"g_date_set_year", &DateSetYear},
	{"g_date_strftime", &DateStrftime},
	{"g_date_subtract_days", &DateSubtractDays},
	{"g_date_subtract_months", &DateSubtractMonths},
	{"g_date_subtract_years", &DateSubtractYears},
	{"g_date_time_add", &DateTimeAdd},
	{"g_date_time_add_days", &DateTimeAddDays},
	{"g_date_time_add_full", &DateTimeAddFull},
	{"g_date_time_add_hours", &DateTimeAddHours},
	{"g_date_time_add_minutes", &DateTimeAddMinutes},
	{"g_date_time_add_months", &DateTimeAddMonths},
	{"g_date_time_add_seconds", &DateTimeAddSeconds},
	{"g_date_time_add_weeks", &DateTimeAddWeeks},
	{"g_date_time_add_years", &DateTimeAddYears},
	{"g_date_time_compare", &DateTimeCompare},
	{"g_date_time_difference", &DateTimeDifference},
	{"g_date_time_equal", &DateTimeEqual},
	{"g_date_time_format", &DateTimeFormat},
	{"g_date_time_get_day_of_month", &DateTimeGetDayOfMonth},
	{"g_date_time_get_day_of_week", &DateTimeGetDayOfWeek},
	{"g_date_time_get_day_of_year", &DateTimeGetDayOfYear},
	{"g_date_time_get_hour", &DateTimeGetHour},
	{"g_date_time_get_microsecond", &DateTimeGetMicrosecond},
	{"g_date_time_get_minute", &DateTimeGetMinute},
	{"g_date_time_get_month", &DateTimeGetMonth},
	{"g_date_time_get_second", &DateTimeGetSecond},
	{"g_date_time_get_seconds", &DateTimeGetSeconds},
	{"g_date_time_get_timezone_abbreviation", &DateTimeGetTimezoneAbbreviation},
	{"g_date_time_get_utc_offset", &DateTimeGetUtcOffset},
	{"g_date_time_get_week_numbering_year", &DateTimeGetWeekNumberingYear},
	{"g_date_time_get_week_of_year", &DateTimeGetWeekOfYear},
	{"g_date_time_get_year", &DateTimeGetYear},
	{"g_date_time_get_ymd", &DateTimeGetYmd},
	{"g_date_time_hash", &DateTimeHash},
	{"g_date_time_is_daylight_savings", &DateTimeIsDaylightSavings},
	{"g_date_time_new", &DateTimeNew},
	{"g_date_time_new_from_timeval_local", &DateTimeNewFromTimevalLocal},
	{"g_date_time_new_from_timeval_utc", &DateTimeNewFromTimevalUtc},
	{"g_date_time_new_from_unix_local", &DateTimeNewFromUnixLocal},
	{"g_date_time_new_from_unix_utc", &DateTimeNewFromUnixUtc},
	{"g_date_time_new_local", &DateTimeNewLocal},
	{"g_date_time_new_now", &DateTimeNewNow},
	{"g_date_time_new_now_local", &DateTimeNewNowLocal},
	{"g_date_time_new_now_utc", &DateTimeNewNowUtc},
	{"g_date_time_new_utc", &DateTimeNewUtc},
	{"g_date_time_ref", &DateTimeRef},
	{"g_date_time_to_local", &DateTimeToLocal},
	{"g_date_time_to_timeval", &DateTimeToTimeval},
	{"g_date_time_to_timezone", &DateTimeToTimezone},
	{"g_date_time_to_unix", &DateTimeToUnix},
	{"g_date_time_to_utc", &DateTimeToUtc},
	{"g_date_time_unref", &DateTimeUnref},
	{"g_date_to_struct_tm", &DateToStructTm},
	{"g_date_valid", &DateValid},
	{"g_date_valid_day", &DateValidDay},
	{"g_date_valid_dmy", &DateValidDmy},
	{"g_date_valid_julian", &DateValidJulian},
	{"g_date_valid_month", &DateValidMonth},
	{"g_date_valid_weekday", &DateValidWeekday},
	{"g_date_valid_year", &DateValidYear},
	{"g_dcgettext", &Dcgettext},
	{"g_dgettext", &Dgettext},
	{"g_dir_close", &DirClose},
	// Windows: _utf8 {"g_dir_open", &DirOpen},
	{"g_dir_open_utf8", &DirOpen},
	// Windows: _utf8 {"g_dir_read_name", &DirReadName},
	{"g_dir_read_name_utf8", &DirReadName},
	{"g_dir_rewind", &DirRewind},
	{"g_direct_equal", &DirectEqual},
	{"g_direct_hash", &DirectHash},
	{"g_dngettext", &Dngettext},
	{"g_double_equal", &DoubleEqual},
	{"g_double_hash", &DoubleHash},
	{"g_dpgettext", &Dpgettext},
	{"g_dpgettext2", &Dpgettext2},
	{"g_error_copy", &ErrorCopy},
	{"g_error_free", &ErrorFree},
	{"g_error_matches", &ErrorMatches},
	{"g_error_new", &ErrorNew},
	{"g_error_new_literal", &ErrorNewLiteral},
	{"g_error_new_valist", &ErrorNewValist},
	{"g_file_error_from_errno", &FileErrorFromErrno},
	{"g_file_error_quark", &FileErrorQuark},
	// Windows: _utf8 {"g_file_get_contents", &FileGetContents},
	{"g_file_get_contents_utf8", &FileGetContents},
	// Windows: _utf8 {"g_file_open_tmp", &FileOpenTmp},
	{"g_file_open_tmp_utf8", &FileOpenTmp},
	{"g_file_read_link", &FileReadLink},
	{"g_file_set_contents", &FileSetContents},
	// Windows: _utf8 {"g_file_test", &FileTest},
	{"g_file_test_utf8", &FileTest},
	{"g_filename_display_basename", &FilenameDisplayBasename},
	{"g_filename_display_name", &FilenameDisplayName},
	// Windows: _utf8 {"g_filename_from_uri", &FilenameFromUri},
	{"g_filename_from_uri_utf8", &FilenameFromUri},
	// Windows: _utf8 {"g_filename_from_utf8", &FilenameFromUtf8},
	{"g_filename_from_utf8_utf8", &FilenameFromUtf8},
	// Windows: _utf8 {"g_filename_to_uri", &FilenameToUri},
	{"g_filename_to_uri_utf8", &FilenameToUri},
	// Windows: _utf8 {"g_filename_to_utf8", &FilenameToUtf8},
	{"g_filename_to_utf8_utf8", &FilenameToUtf8},
	// Windows: _utf8 {"g_find_program_in_path", &FindProgramInPath},
	{"g_find_program_in_path_utf8", &FindProgramInPath},
	{"g_fopen", &Fopen},
	{"g_format_size_for_display", &FormatSizeForDisplay},
	{"g_fprintf", &Fprintf},
	{"g_free", &Free},
	{"g_freopen", &Freopen},
	{"g_get_application_name", &GetApplicationName},
	{"g_get_charset", &GetCharset},
	// Undocumented {"g_get_codeset", &GetCodeset},
	// Windows: _utf8 {"g_get_current_dir", &GetCurrentDir},
	{"g_get_current_dir_utf8", &GetCurrentDir},
	{"g_get_current_time", &GetCurrentTime},
	{"g_get_environ", &GetEnviron},
	{"g_get_filename_charsets", &GetFilenameCharsets},
	// Windows: _utf8 {"g_get_home_dir", &GetHomeDir},
	{"g_get_home_dir_utf8", &GetHomeDir},
	{"g_get_host_name", &GetHostName},
	{"g_get_language_names", &GetLanguageNames},
	{"g_get_locale_variants", &GetLocaleVariants},
	{"g_get_monotonic_time", &GetMonotonicTime},
	{"g_get_prgname", &GetPrgname},
	// Windows: _utf8 {"g_get_real_name", &GetRealName},
	{"g_get_real_name_utf8", &GetRealName},
	{"g_get_real_time", &GetRealTime},
	{"g_get_system_config_dirs", &GetSystemConfigDirs},
	{"g_get_system_data_dirs", &GetSystemDataDirs},
	// Windows: _utf8 {"g_get_tmp_dir", &GetTmpDir},
	{"g_get_tmp_dir_utf8", &GetTmpDir},
	{"g_get_user_cache_dir", &GetUserCacheDir},
	{"g_get_user_config_dir", &GetUserConfigDir},
	{"g_get_user_data_dir", &GetUserDataDir},
	// Windows: _utf8 {"g_get_user_name", &GetUserName},
	{"g_get_user_name_utf8", &GetUserName},
	{"g_get_user_runtime_dir", &GetUserRuntimeDir},
	{"g_get_user_special_dir", &GetUserSpecialDir},
	// Windows: _utf8 {"g_getenv", &Getenv},
	{"g_getenv_utf8", &Getenv},
	{"g_hash_table_destroy", &HashTableDestroy},
	{"g_hash_table_find", &HashTableFind},
	{"g_hash_table_foreach", &HashTableForeach},
	{"g_hash_table_foreach_remove", &HashTableForeachRemove},
	{"g_hash_table_foreach_steal", &HashTableForeachSteal},
	{"g_hash_table_get_keys", &HashTableGetKeys},
	{"g_hash_table_get_values", &HashTableGetValues},
	{"g_hash_table_insert", &HashTableInsert},
	{"g_hash_table_iter_get_hash_table", &HashTableIterGetHashTable},
	{"g_hash_table_iter_init", &HashTableIterInit},
	{"g_hash_table_iter_next", &HashTableIterNext},
	{"g_hash_table_iter_remove", &HashTableIterRemove},
	{"g_hash_table_iter_steal", &HashTableIterSteal},
	{"g_hash_table_lookup", &HashTableLookup},
	{"g_hash_table_lookup_extended", &HashTableLookupExtended},
	{"g_hash_table_new", &HashTableNew},
	{"g_hash_table_new_full", &HashTableNewFull},
	{"g_hash_table_ref", &HashTableRef},
	{"g_hash_table_remove", &HashTableRemove},
	{"g_hash_table_remove_all", &HashTableRemoveAll},
	{"g_hash_table_replace", &HashTableReplace},
	{"g_hash_table_size", &HashTableSize},
	{"g_hash_table_steal", &HashTableSteal},
	{"g_hash_table_steal_all", &HashTableStealAll},
	{"g_hash_table_unref", &HashTableUnref},
	{"g_hook_alloc", &HookAlloc},
	{"g_hook_compare_ids", &HookCompareIds},
	{"g_hook_destroy", &HookDestroy},
	{"g_hook_destroy_link", &HookDestroyLink},
	{"g_hook_find", &HookFind},
	{"g_hook_find_data", &HookFindData},
	{"g_hook_find_func", &HookFindFunc_},
	{"g_hook_find_func_data", &HookFindFuncData},
	{"g_hook_first_valid", &HookFirstValid},
	{"g_hook_free", &HookFree},
	{"g_hook_get", &HookGet},
	{"g_hook_insert_before", &HookInsertBefore},
	{"g_hook_insert_sorted", &HookInsertSorted},
	{"g_hook_list_clear", &HookListClear},
	{"g_hook_list_init", &HookListInit},
	{"g_hook_list_invoke", &HookListInvoke},
	{"g_hook_list_invoke_check", &HookListInvokeCheck},
	{"g_hook_list_marshal", &HookListMarshal},
	{"g_hook_list_marshal_check", &HookListMarshalCheck},
	{"g_hook_next_valid", &HookNextValid},
	{"g_hook_prepend", &HookPrepend},
	{"g_hook_ref", &HookRef},
	{"g_hook_unref", &HookUnref},
	{"g_hostname_is_ascii_encoded", &HostnameIsAsciiEncoded},
	{"g_hostname_is_ip_address", &HostnameIsIpAddress},
	{"g_hostname_is_non_ascii", &HostnameIsNonAscii},
	{"g_hostname_to_ascii", &HostnameToAscii},
	{"g_hostname_to_unicode", &HostnameToUnicode},
	{"g_iconv", &Iconv},
	{"g_iconv_close", &IconvClose},
	{"g_iconv_open", &IconvOpen},
	{"g_idle_add", &IdleAdd},
	{"g_idle_add_full", &IdleAddFull},
	{"g_idle_remove_by_data", &IdleRemoveByData},
	{"g_idle_source_new", &IdleSourceNew},
	{"g_int64_equal", &Int64Equal},
	{"g_int64_hash", &Int64Hash},
	{"g_int_equal", &IntEqual},
	{"g_int_hash", &IntHash},
	{"g_intern_static_string", &InternStaticString},
	{"g_intern_string", &InternString},
	{"g_io_add_watch", &IoAddWatch},
	{"g_io_add_watch_full", &IoAddWatchFull},
	{"g_io_channel_close", &IoChannelClose},
	{"g_io_channel_error_from_errno", &IoChannelErrorFromErrno},
	{"g_io_channel_error_quark", &IoChannelErrorQuark},
	{"g_io_channel_flush", &IoChannelFlush},
	{"g_io_channel_get_buffer_condition", &IoChannelGetBufferCondition},
	{"g_io_channel_get_buffer_size", &IoChannelGetBufferSize},
	{"g_io_channel_get_buffered", &IoChannelGetBuffered},
	{"g_io_channel_get_close_on_unref", &IoChannelGetCloseOnUnref},
	{"g_io_channel_get_encoding", &IoChannelGetEncoding},
	{"g_io_channel_get_flags", &IoChannelGetFlags},
	{"g_io_channel_get_line_term", &IoChannelGetLineTerm},
	{"g_io_channel_init", &IoChannelInit},
	// Windows: _utf8 {"g_io_channel_new_file", &IoChannelNewFile},
	{"g_io_channel_new_file_utf8", &IoChannelNewFile},
	{"g_io_channel_read", &IoChannelRead},
	{"g_io_channel_read_chars", &IoChannelReadChars},
	{"g_io_channel_read_line", &IoChannelReadLine},
	{"g_io_channel_read_line_string", &IoChannelReadLineString},
	{"g_io_channel_read_to_end", &IoChannelReadToEnd},
	{"g_io_channel_read_unichar", &IoChannelReadUnichar},
	{"g_io_channel_ref", &IoChannelRef},
	{"g_io_channel_seek", &IoChannelSeek},
	{"g_io_channel_seek_position", &IoChannelSeekPosition},
	{"g_io_channel_set_buffer_size", &IoChannelSetBufferSize},
	{"g_io_channel_set_buffered", &IoChannelSetBuffered},
	{"g_io_channel_set_close_on_unref", &IoChannelSetCloseOnUnref},
	{"g_io_channel_set_encoding", &IoChannelSetEncoding},
	{"g_io_channel_set_flags", &IoChannelSetFlags},
	{"g_io_channel_set_line_term", &IoChannelSetLineTerm},
	{"g_io_channel_shutdown", &IoChannelShutdown},
	{"g_io_channel_unix_get_fd", &IoChannelUnixGetFd},
	{"g_io_channel_unix_new", &IoChannelUnixNew},
	{"g_io_channel_unref", &IoChannelUnref},
	{"g_io_channel_win32_get_fd", &IoChannelWin32GetFd},
	{"g_io_channel_win32_make_pollfd", &IoChannelWin32MakePollfd},
	{"g_io_channel_win32_new_fd", &IoChannelWin32NewFd},
	{"g_io_channel_win32_new_messages", &IoChannelWin32NewMessages},
	{"g_io_channel_win32_new_socket", &IoChannelWin32NewSocket},
	// Undocumented {"g_io_channel_win32_new_stream_socket", &IoChannelWin32NewStreamSocket},
	{"g_io_channel_win32_poll", &IoChannelWin32Poll},
	// Undocumented {"g_io_channel_win32_set_debug", &IoChannelWin32SetDebug},
	{"g_io_channel_write", &IoChannelWrite},
	{"g_io_channel_write_chars", &IoChannelWriteChars},
	{"g_io_channel_write_unichar", &IoChannelWriteUnichar},
	{"g_io_create_watch", &IoCreateWatch},
	{"g_key_file_error_quark", &KeyFileErrorQuark},
	{"g_key_file_free", &KeyFileFree},
	{"g_key_file_get_boolean", &KeyFileGetBoolean},
	{"g_key_file_get_boolean_list", &KeyFileGetBooleanList},
	{"g_key_file_get_comment", &KeyFileGetComment},
	{"g_key_file_get_double", &KeyFileGetDouble},
	{"g_key_file_get_double_list", &KeyFileGetDoubleList},
	{"g_key_file_get_groups", &KeyFileGetGroups},
	{"g_key_file_get_int64", &KeyFileGetInt64},
	{"g_key_file_get_integer", &KeyFileGetInteger},
	{"g_key_file_get_integer_list", &KeyFileGetIntegerList},
	{"g_key_file_get_keys", &KeyFileGetKeys},
	{"g_key_file_get_locale_string", &KeyFileGetLocaleString},
	{"g_key_file_get_locale_string_list", &KeyFileGetLocaleStringList},
	{"g_key_file_get_start_group", &KeyFileGetStartGroup},
	{"g_key_file_get_string", &KeyFileGetString},
	{"g_key_file_get_string_list", &KeyFileGetStringList},
	{"g_key_file_get_uint64", &KeyFileGetUint64},
	{"g_key_file_get_value", &KeyFileGetValue},
	{"g_key_file_has_group", &KeyFileHasGroup},
	{"g_key_file_has_key", &KeyFileHasKey},
	{"g_key_file_load_from_data", &KeyFileLoadFromData},
	{"g_key_file_load_from_data_dirs", &KeyFileLoadFromDataDirs},
	{"g_key_file_load_from_dirs", &KeyFileLoadFromDirs},
	{"g_key_file_load_from_file", &KeyFileLoadFromFile},
	{"g_key_file_new", &KeyFileNew},
	{"g_key_file_remove_comment", &KeyFileRemoveComment},
	{"g_key_file_remove_group", &KeyFileRemoveGroup},
	{"g_key_file_remove_key", &KeyFileRemoveKey},
	{"g_key_file_set_boolean", &KeyFileSetBoolean},
	{"g_key_file_set_boolean_list", &KeyFileSetBooleanList},
	{"g_key_file_set_comment", &KeyFileSetComment},
	{"g_key_file_set_double", &KeyFileSetDouble},
	{"g_key_file_set_double_list", &KeyFileSetDoubleList},
	{"g_key_file_set_int64", &KeyFileSetInt64},
	{"g_key_file_set_integer", &KeyFileSetInteger},
	{"g_key_file_set_integer_list", &KeyFileSetIntegerList},
	{"g_key_file_set_list_separator", &KeyFileSetListSeparator},
	{"g_key_file_set_locale_string", &KeyFileSetLocaleString},
	{"g_key_file_set_locale_string_list", &KeyFileSetLocaleStringList},
	{"g_key_file_set_string", &KeyFileSetString},
	{"g_key_file_set_string_list", &KeyFileSetStringList},
	{"g_key_file_set_uint64", &KeyFileSetUint64},
	{"g_key_file_set_value", &KeyFileSetValue},
	{"g_key_file_to_data", &KeyFileToData},
	{"g_list_alloc", &ListAlloc},
	{"g_list_append", &ListAppend},
	{"g_list_concat", &ListConcat},
	{"g_list_copy", &ListCopy},
	{"g_list_delete_link", &ListDeleteLink},
	{"g_list_find", &ListFind},
	{"g_list_find_custom", &ListFindCustom},
	{"g_list_first", &ListFirst},
	{"g_list_foreach", &ListForeach},
	{"g_list_free", &ListFree},
	{"g_list_free_1", &ListFree1},
	{"g_list_free_full", &ListFreeFull},
	{"g_list_index", &ListIndex},
	{"g_list_insert", &ListInsert},
	{"g_list_insert_before", &ListInsertBefore},
	{"g_list_insert_sorted", &ListInsertSorted},
	{"g_list_insert_sorted_with_data", &ListInsertSortedWithData},
	{"g_list_last", &ListLast},
	{"g_list_length", &ListLength},
	{"g_list_nth", &ListNth},
	{"g_list_nth_data", &ListNthData},
	{"g_list_nth_prev", &ListNthPrev},
	{"g_list_pop_allocator", &ListPopAllocator},
	{"g_list_position", &ListPosition},
	{"g_list_prepend", &ListPrepend},
	{"g_list_push_allocator", &ListPushAllocator},
	{"g_list_remove", &ListRemove},
	{"g_list_remove_all", &ListRemoveAll},
	{"g_list_remove_link", &ListRemoveLink},
	{"g_list_reverse", &ListReverse},
	{"g_list_sort", &ListSort},
	{"g_list_sort_with_data", &ListSortWithData},
	{"g_listenv", &Listenv},
	{"g_locale_from_utf8", &LocaleFromUtf8},
	{"g_locale_to_utf8", &LocaleToUtf8},
	{"g_log", &Log},
	{"g_log_default_handler", &LogDefaultHandler},
	{"g_log_remove_handler", &LogRemoveHandler},
	{"g_log_set_always_fatal", &LogSetAlwaysFatal},
	{"g_log_set_default_handler", &LogSetDefaultHandler},
	{"g_log_set_fatal_mask", &LogSetFatalMask},
	{"g_log_set_handler", &LogSetHandler},
	{"g_logv", &Logv},
	{"g_lstat", &Lstat},
	{"g_main_context_acquire", &MainContextAcquire},
	{"g_main_context_add_poll", &MainContextAddPoll},
	{"g_main_context_check", &MainContextCheck},
	{"g_main_context_default", &MainContextDefault},
	{"g_main_context_dispatch", &MainContextDispatch},
	{"g_main_context_find_source_by_funcs_user_data", &MainContextFindSourceByFuncsUserData},
	{"g_main_context_find_source_by_id", &MainContextFindSourceById},
	{"g_main_context_find_source_by_user_data", &MainContextFindSourceByUserData},
	{"g_main_context_get_poll_func", &MainContextGetPollFunc},
	{"g_main_context_get_thread_default", &MainContextGetThreadDefault},
	{"g_main_context_invoke", &MainContextInvoke},
	{"g_main_context_invoke_full", &MainContextInvokeFull},
	{"g_main_context_is_owner", &MainContextIsOwner},
	{"g_main_context_iteration", &MainContextIteration},
	{"g_main_context_new", &MainContextNew},
	{"g_main_context_pending", &MainContextPending},
	{"g_main_context_pop_thread_default", &MainContextPopThreadDefault},
	{"g_main_context_prepare", &MainContextPrepare},
	{"g_main_context_push_thread_default", &MainContextPushThreadDefault},
	{"g_main_context_query", &MainContextQuery},
	{"g_main_context_ref", &MainContextRef},
	{"g_main_context_release", &MainContextRelease},
	{"g_main_context_remove_poll", &MainContextRemovePoll},
	{"g_main_context_set_poll_func", &MainContextSetPollFunc},
	{"g_main_context_unref", &MainContextUnref},
	{"g_main_context_wait", &MainContextWait},
	{"g_main_context_wakeup", &MainContextWakeup},
	{"g_main_current_source", &MainCurrentSource},
	{"g_main_depth", &MainDepth},
	{"g_main_loop_get_context", &MainLoopGetContext},
	{"g_main_loop_is_running", &MainLoopIsRunning},
	{"g_main_loop_new", &MainLoopNew},
	{"g_main_loop_quit", &MainLoopQuit},
	{"g_main_loop_ref", &MainLoopRef},
	{"g_main_loop_run", &MainLoopRun},
	{"g_main_loop_unref", &MainLoopUnref},
	{"g_malloc", &Malloc},
	{"g_malloc0", &Malloc0},
	{"g_malloc0_n", &Malloc0N},
	{"g_malloc_n", &MallocN},
	{"g_mapped_file_free", &MappedFileFree},
	{"g_mapped_file_get_contents", &MappedFileGetContents},
	{"g_mapped_file_get_length", &MappedFileGetLength},
	{"g_mapped_file_new", &MappedFileNew},
	{"g_mapped_file_ref", &MappedFileRef},
	{"g_mapped_file_unref", &MappedFileUnref},
	{"g_markup_collect_attributes", &MarkupCollectAttributes},
	{"g_markup_error_quark", &MarkupErrorQuark},
	{"g_markup_escape_text", &MarkupEscapeText},
	{"g_markup_parse_context_end_parse", &MarkupParseContextEndParse},
	{"g_markup_parse_context_free", &MarkupParseContextFree},
	{"g_markup_parse_context_get_element", &MarkupParseContextGetElement},
	{"g_markup_parse_context_get_element_stack", &MarkupParseContextGetElementStack},
	{"g_markup_parse_context_get_position", &MarkupParseContextGetPosition},
	{"g_markup_parse_context_get_user_data", &MarkupParseContextGetUserData},
	{"g_markup_parse_context_new", &MarkupParseContextNew},
	{"g_markup_parse_context_parse", &MarkupParseContextParse},
	{"g_markup_parse_context_pop", &MarkupParseContextPop},
	{"g_markup_parse_context_push", &MarkupParseContextPush},
	{"g_markup_printf_escaped", &MarkupPrintfEscaped},
	{"g_markup_vprintf_escaped", &MarkupVprintfEscaped},
	{"g_match_info_expand_references", &MatchInfoExpandReferences},
	{"g_match_info_fetch", &MatchInfoFetch},
	{"g_match_info_fetch_all", &MatchInfoFetchAll},
	{"g_match_info_fetch_named", &MatchInfoFetchNamed},
	{"g_match_info_fetch_named_pos", &MatchInfoFetchNamedPos},
	{"g_match_info_fetch_pos", &MatchInfoFetchPos},
	{"g_match_info_free", &MatchInfoFree},
	{"g_match_info_get_match_count", &MatchInfoGetMatchCount},
	{"g_match_info_get_regex", &MatchInfoGetRegex},
	{"g_match_info_get_string", &MatchInfoGetString},
	{"g_match_info_is_partial_match", &MatchInfoIsPartialMatch},
	{"g_match_info_matches", &MatchInfoMatches},
	{"g_match_info_next", &MatchInfoNext},
	{"g_mem_chunk_alloc", &MemChunkAlloc},
	{"g_mem_chunk_alloc0", &MemChunkAlloc0},
	{"g_mem_chunk_clean", &MemChunkClean},
	{"g_mem_chunk_destroy", &MemChunkDestroy},
	{"g_mem_chunk_free", &MemChunkFree},
	{"g_mem_chunk_info", &MemChunkInfo},
	{"g_mem_chunk_new", &MemChunkNew},
	{"g_mem_chunk_print", &MemChunkPrint},
	{"g_mem_chunk_reset", &MemChunkReset},
	{"g_mem_is_system_malloc", &MemIsSystemMalloc},
	{"g_mem_profile", &MemProfile},
	{"g_mem_set_vtable", &MemSetVtable},
	{"g_memdup", &Memdup},
	{"g_mkdir", &Mkdir},
	{"g_mkdir_with_parents", &MkdirWithParents},
	// Windows: _utf8 {"g_mkstemp", &Mkstemp},
	{"g_mkstemp_full", &MkstempFull},
	{"g_mkstemp_utf8", &Mkstemp},
	{"g_node_child_index", &NodeChildIndex},
	{"g_node_child_position", &NodeChildPosition},
	{"g_node_children_foreach", &NodeChildrenForeach},
	{"g_node_copy", &NodeCopy},
	{"g_node_copy_deep", &NodeCopyDeep},
	{"g_node_depth", &NodeDepth},
	{"g_node_destroy", &NodeDestroy},
	{"g_node_find", &NodeFind},
	{"g_node_find_child", &NodeFindChild},
	{"g_node_first_sibling", &NodeFirstSibling},
	{"g_node_get_root", &NodeGetRoot},
	{"g_node_insert", &NodeInsert},
	{"g_node_insert_after", &NodeInsertAfter},
	{"g_node_insert_before", &NodeInsertBefore},
	{"g_node_is_ancestor", &NodeIsAncestor},
	{"g_node_last_child", &NodeLastChild},
	{"g_node_last_sibling", &NodeLastSibling},
	{"g_node_max_height", &NodeMaxHeight},
	{"g_node_n_children", &NodeNChildren},
	{"g_node_n_nodes", &NodeNNodes},
	{"g_node_new", &NodeNew},
	{"g_node_nth_child", &NodeNthChild},
	{"g_node_pop_allocator", &NodePopAllocator},
	{"g_node_prepend", &NodePrepend},
	{"g_node_push_allocator", &NodePushAllocator},
	{"g_node_reverse_children", &NodeReverseChildren},
	{"g_node_traverse", &NodeTraverse},
	{"g_node_unlink", &NodeUnlink},
	{"g_nullify_pointer", &NullifyPointer},
	{"g_on_error_query", &OnErrorQuery},
	{"g_on_error_stack_trace", &OnErrorStackTrace},
	{"g_once_impl", &OnceImpl},
	{"g_once_init_enter", &OnceInitEnter},
	{"g_once_init_enter_impl", &OnceInitEnterImpl},
	{"g_once_init_leave", &OnceInitLeave},
	{"g_open", &Open},
	{"g_option_context_add_group", &OptionContextAddGroup},
	{"g_option_context_add_main_entries", &OptionContextAddMainEntries},
	{"g_option_context_free", &OptionContextFree},
	{"g_option_context_get_description", &OptionContextGetDescription},
	{"g_option_context_get_help", &OptionContextGetHelp},
	{"g_option_context_get_help_enabled", &OptionContextGetHelpEnabled},
	{"g_option_context_get_ignore_unknown_options", &OptionContextGetIgnoreUnknownOptions},
	{"g_option_context_get_main_group", &OptionContextGetMainGroup},
	{"g_option_context_get_summary", &OptionContextGetSummary},
	{"g_option_context_new", &OptionContextNew},
	{"g_option_context_parse", &OptionContextParse},
	{"g_option_context_set_description", &OptionContextSetDescription},
	{"g_option_context_set_help_enabled", &OptionContextSetHelpEnabled},
	{"g_option_context_set_ignore_unknown_options", &OptionContextSetIgnoreUnknownOptions},
	{"g_option_context_set_main_group", &OptionContextSetMainGroup},
	{"g_option_context_set_summary", &OptionContextSetSummary},
	{"g_option_context_set_translate_func", &OptionContextSetTranslateFunc},
	{"g_option_context_set_translation_domain", &OptionContextSetTranslationDomain},
	{"g_option_error_quark", &OptionErrorQuark},
	{"g_option_group_add_entries", &OptionGroupAddEntries},
	{"g_option_group_free", &OptionGroupFree},
	{"g_option_group_new", &OptionGroupNew},
	{"g_option_group_set_error_hook", &OptionGroupSetErrorHook},
	{"g_option_group_set_parse_hooks", &OptionGroupSetParseHooks},
	{"g_option_group_set_translate_func", &OptionGroupSetTranslateFunc},
	{"g_option_group_set_translation_domain", &OptionGroupSetTranslationDomain},
	{"g_parse_debug_string", &ParseDebugString},
	{"g_path_get_basename", &PathGetBasename},
	{"g_path_get_dirname", &PathGetDirname},
	{"g_path_is_absolute", &PathIsAbsolute},
	{"g_path_skip_root", &PathSkipRoot},
	{"g_pattern_match", &PatternMatch},
	{"g_pattern_match_simple", &PatternMatchSimple},
	{"g_pattern_match_string", &PatternMatchString},
	{"g_pattern_spec_equal", &PatternSpecEqual},
	{"g_pattern_spec_free", &PatternSpecFree},
	{"g_pattern_spec_new", &PatternSpecNew},
	{"g_poll", &Poll},
	{"g_prefix_error", &PrefixError},
	{"g_print", &Print},
	{"g_printerr", &Printerr},
	{"g_printf", &Printf},
	{"g_printf_string_upper_bound", &PrintfStringUpperBound},
	{"g_propagate_error", &PropagateError},
	{"g_propagate_prefixed_error", &PropagatePrefixedError},
	{"g_ptr_array_add", &PtrArrayAdd},
	{"g_ptr_array_foreach", &PtrArrayForeach},
	{"g_ptr_array_free", &PtrArrayFree},
	{"g_ptr_array_new", &PtrArrayNew},
	{"g_ptr_array_new_with_free_func", &PtrArrayNewWithFreeFunc},
	{"g_ptr_array_ref", &PtrArrayRef},
	{"g_ptr_array_remove", &PtrArrayRemove},
	{"g_ptr_array_remove_fast", &PtrArrayRemoveFast},
	{"g_ptr_array_remove_index", &PtrArrayRemoveIndex},
	{"g_ptr_array_remove_index_fast", &PtrArrayRemoveIndexFast},
	{"g_ptr_array_remove_range", &PtrArrayRemoveRange},
	{"g_ptr_array_set_free_func", &PtrArraySetFreeFunc},
	{"g_ptr_array_set_size", &PtrArraySetSize},
	{"g_ptr_array_sized_new", &PtrArraySizedNew},
	{"g_ptr_array_sort", &PtrArraySort},
	{"g_ptr_array_sort_with_data", &PtrArraySortWithData},
	{"g_ptr_array_unref", &PtrArrayUnref},
	{"g_qsort_with_data", &QsortWithData},
	{"g_quark_from_static_string", &QuarkFromStaticString},
	{"g_quark_from_string", &QuarkFromString},
	{"g_quark_to_string", &QuarkToString},
	{"g_quark_try_string", &QuarkTryString},
	{"g_queue_clear", &QueueClear},
	{"g_queue_copy", &QueueCopy},
	{"g_queue_delete_link", &QueueDeleteLink},
	{"g_queue_find", &QueueFind},
	{"g_queue_find_custom", &QueueFindCustom},
	{"g_queue_foreach", &QueueForeach},
	{"g_queue_free", &QueueFree},
	{"g_queue_get_length", &QueueGetLength},
	{"g_queue_index", &QueueIndex},
	{"g_queue_init", &QueueInit},
	{"g_queue_insert_after", &QueueInsertAfter},
	{"g_queue_insert_before", &QueueInsertBefore},
	{"g_queue_insert_sorted", &QueueInsertSorted},
	{"g_queue_is_empty", &QueueIsEmpty},
	{"g_queue_link_index", &QueueLinkIndex},
	{"g_queue_new", &QueueNew},
	{"g_queue_peek_head", &QueuePeekHead},
	{"g_queue_peek_head_link", &QueuePeekHeadLink},
	{"g_queue_peek_nth", &QueuePeekNth},
	{"g_queue_peek_nth_link", &QueuePeekNthLink},
	{"g_queue_peek_tail", &QueuePeekTail},
	{"g_queue_peek_tail_link", &QueuePeekTailLink},
	{"g_queue_pop_head", &QueuePopHead},
	{"g_queue_pop_head_link", &QueuePopHeadLink},
	{"g_queue_pop_nth", &QueuePopNth},
	{"g_queue_pop_nth_link", &QueuePopNthLink},
	{"g_queue_pop_tail", &QueuePopTail},
	{"g_queue_pop_tail_link", &QueuePopTailLink},
	{"g_queue_push_head", &QueuePushHead},
	{"g_queue_push_head_link", &QueuePushHeadLink},
	{"g_queue_push_nth", &QueuePushNth},
	{"g_queue_push_nth_link", &QueuePushNthLink},
	{"g_queue_push_tail", &QueuePushTail},
	{"g_queue_push_tail_link", &QueuePushTailLink},
	{"g_queue_remove", &QueueRemove},
	{"g_queue_remove_all", &QueueRemoveAll},
	{"g_queue_reverse", &QueueReverse},
	{"g_queue_sort", &QueueSort},
	{"g_queue_unlink", &QueueUnlink},
	{"g_rand_copy", &RandCopy},
	{"g_rand_double", &RandDouble},
	{"g_rand_double_range", &RandDoubleRange},
	{"g_rand_free", &RandFree},
	{"g_rand_int", &RandInt},
	{"g_rand_int_range", &RandIntRange},
	{"g_rand_new", &RandNew},
	{"g_rand_new_with_seed", &RandNewWithSeed},
	{"g_rand_new_with_seed_array", &RandNewWithSeedArray},
	{"g_rand_set_seed", &RandSetSeed},
	{"g_rand_set_seed_array", &RandSetSeedArray},
	{"g_random_double", &RandomDouble},
	{"g_random_double_range", &RandomDoubleRange},
	{"g_random_int", &RandomInt},
	{"g_random_int_range", &RandomIntRange},
	{"g_random_set_seed", &RandomSetSeed},
	{"g_realloc", &Realloc},
	{"g_realloc_n", &ReallocN},
	{"g_regex_check_replacement", &RegexCheckReplacement},
	{"g_regex_error_quark", &RegexErrorQuark},
	{"g_regex_escape_string", &RegexEscapeString},
	{"g_regex_get_capture_count", &RegexGetCaptureCount},
	{"g_regex_get_compile_flags", &RegexGetCompileFlags},
	{"g_regex_get_match_flags", &RegexGetMatchFlags},
	{"g_regex_get_max_backref", &RegexGetMaxBackref},
	{"g_regex_get_pattern", &RegexGetPattern},
	{"g_regex_get_string_number", &RegexGetStringNumber},
	{"g_regex_match", &RegexMatch},
	{"g_regex_match_all", &RegexMatchAll},
	{"g_regex_match_all_full", &RegexMatchAllFull},
	{"g_regex_match_full", &RegexMatchFull},
	{"g_regex_match_simple", &RegexMatchSimple},
	{"g_regex_new", &RegexNew},
	{"g_regex_ref", &RegexRef},
	{"g_regex_replace", &RegexReplace},
	{"g_regex_replace_eval", &RegexReplaceEval},
	{"g_regex_replace_literal", &RegexReplaceLiteral},
	{"g_regex_split", &RegexSplit},
	{"g_regex_split_full", &RegexSplitFull},
	{"g_regex_split_simple", &RegexSplitSimple},
	{"g_regex_unref", &RegexUnref},
	{"g_relation_count", &RelationCount},
	{"g_relation_delete", &RelationDelete},
	{"g_relation_destroy", &RelationDestroy},
	{"g_relation_exists", &RelationExists},
	{"g_relation_index", &RelationIndex},
	{"g_relation_insert", &RelationInsert},
	{"g_relation_new", &RelationNew},
	{"g_relation_print", &RelationPrint},
	{"g_relation_select", &RelationSelect},
	{"g_reload_user_special_dirs_cache", &ReloadUserSpecialDirsCache},
	{"g_remove", &Remove},
	{"g_rename", &Rename},
	{"g_return_if_fail_warning", &ReturnIfFailWarning},
	{"g_rmdir", &Rmdir},
	{"g_scanner_cur_line", &ScannerCurLine},
	{"g_scanner_cur_position", &ScannerCurPosition},
	{"g_scanner_cur_token", &ScannerCurToken},
	{"g_scanner_cur_value", &ScannerCurValue},
	{"g_scanner_destroy", &ScannerDestroy},
	{"g_scanner_eof", &ScannerEof},
	{"g_scanner_error", &ScannerError},
	{"g_scanner_get_next_token", &ScannerGetNextToken},
	{"g_scanner_input_file", &ScannerInputFile},
	{"g_scanner_input_text", &ScannerInputText},
	{"g_scanner_lookup_symbol", &ScannerLookupSymbol},
	{"g_scanner_new", &ScannerNew},
	{"g_scanner_peek_next_token", &ScannerPeekNextToken},
	{"g_scanner_scope_add_symbol", &ScannerScopeAddSymbol},
	{"g_scanner_scope_foreach_symbol", &ScannerScopeForeachSymbol},
	{"g_scanner_scope_lookup_symbol", &ScannerScopeLookupSymbol},
	{"g_scanner_scope_remove_symbol", &ScannerScopeRemoveSymbol},
	{"g_scanner_set_scope", &ScannerSetScope},
	{"g_scanner_sync_file_offset", &ScannerSyncFileOffset},
	{"g_scanner_unexp_token", &ScannerUnexpToken},
	{"g_scanner_warn", &ScannerWarn},
	{"g_sequence_append", &SequenceAppend},
	{"g_sequence_foreach", &SequenceForeach},
	{"g_sequence_foreach_range", &SequenceForeachRange},
	{"g_sequence_free", &SequenceFree},
	{"g_sequence_get", &SequenceGet},
	{"g_sequence_get_begin_iter", &SequenceGetBeginIter},
	{"g_sequence_get_end_iter", &SequenceGetEndIter},
	{"g_sequence_get_iter_at_pos", &SequenceGetIterAtPos},
	{"g_sequence_get_length", &SequenceGetLength},
	{"g_sequence_insert_before", &SequenceInsertBefore},
	{"g_sequence_insert_sorted", &SequenceInsertSorted},
	{"g_sequence_insert_sorted_iter", &SequenceInsertSortedIter},
	{"g_sequence_iter_compare", &SequenceIterCompare},
	{"g_sequence_iter_get_position", &SequenceIterGetPosition},
	{"g_sequence_iter_get_sequence", &SequenceIterGetSequence},
	{"g_sequence_iter_is_begin", &SequenceIterIsBegin},
	{"g_sequence_iter_is_end", &SequenceIterIsEnd},
	{"g_sequence_iter_move", &SequenceIterMove},
	{"g_sequence_iter_next", &SequenceIterNext},
	{"g_sequence_iter_prev", &SequenceIterPrev},
	{"g_sequence_lookup", &SequenceLookup},
	{"g_sequence_lookup_iter", &SequenceLookupIter},
	{"g_sequence_move", &SequenceMove},
	{"g_sequence_move_range", &SequenceMoveRange},
	{"g_sequence_new", &SequenceNew},
	{"g_sequence_prepend", &SequencePrepend},
	{"g_sequence_range_get_midpoint", &SequenceRangeGetMidpoint},
	{"g_sequence_remove", &SequenceRemove},
	{"g_sequence_remove_range", &SequenceRemoveRange},
	{"g_sequence_search", &SequenceSearch},
	{"g_sequence_search_iter", &SequenceSearchIter},
	{"g_sequence_set", &SequenceSet},
	{"g_sequence_sort", &SequenceSort},
	{"g_sequence_sort_changed", &SequenceSortChanged},
	{"g_sequence_sort_changed_iter", &SequenceSortChangedIter},
	{"g_sequence_sort_iter", &SequenceSortIter},
	{"g_sequence_swap", &SequenceSwap},
	{"g_set_application_name", &SetApplicationName},
	{"g_set_error", &SetError},
	{"g_set_error_literal", &SetErrorLiteral},
	{"g_set_prgname", &SetPrgname},
	{"g_set_print_handler", &SetPrintHandler},
	{"g_set_printerr_handler", &SetPrinterrHandler},
	// Windows: _utf8 {"g_setenv", &Setenv},
	{"g_setenv_utf8", &Setenv},
	{"g_shell_error_quark", &ShellErrorQuark},
	{"g_shell_parse_argv", &ShellParseArgv},
	{"g_shell_quote", &ShellQuote},
	{"g_shell_unquote", &ShellUnquote},
	{"g_slice_alloc", &SliceAlloc},
	{"g_slice_alloc0", &SliceAlloc0},
	{"g_slice_copy", &SliceCopy},
	{"g_slice_free1", &SliceFree1},
	{"g_slice_free_chain_with_offset", &SliceFreeChainWithOffset},
	{"g_slice_get_config", &SliceGetConfig},
	{"g_slice_get_config_state", &SliceGetConfigState},
	{"g_slice_set_config", &SliceSetConfig},
	{"g_slist_alloc", &SlistAlloc},
	{"g_slist_append", &SlistAppend},
	{"g_slist_concat", &SlistConcat},
	{"g_slist_copy", &SlistCopy},
	{"g_slist_delete_link", &SlistDeleteLink},
	{"g_slist_find", &SlistFind},
	{"g_slist_find_custom", &SlistFindCustom},
	{"g_slist_foreach", &SlistForeach},
	{"g_slist_free", &SlistFree},
	{"g_slist_free_1", &SlistFree1},
	{"g_slist_free_full", &SlistFreeFull},
	{"g_slist_index", &SlistIndex},
	{"g_slist_insert", &SlistInsert},
	{"g_slist_insert_before", &SlistInsertBefore},
	{"g_slist_insert_sorted", &SlistInsertSorted},
	{"g_slist_insert_sorted_with_data", &SlistInsertSortedWithData},
	{"g_slist_last", &SlistLast},
	{"g_slist_length", &SlistLength},
	{"g_slist_nth", &SlistNth},
	{"g_slist_nth_data", &SlistNthData},
	{"g_slist_pop_allocator", &SlistPopAllocator},
	{"g_slist_position", &SlistPosition},
	{"g_slist_prepend", &SlistPrepend},
	{"g_slist_push_allocator", &SlistPushAllocator},
	{"g_slist_remove", &SlistRemove},
	{"g_slist_remove_all", &SlistRemoveAll},
	{"g_slist_remove_link", &SlistRemoveLink},
	{"g_slist_reverse", &SlistReverse},
	{"g_slist_sort", &SlistSort},
	{"g_slist_sort_with_data", &SlistSortWithData},
	{"g_snprintf", &Snprintf},
	{"g_source_add_child_source", &SourceAddChildSource},
	{"g_source_add_poll", &SourceAddPoll},
	{"g_source_attach", &SourceAttach},
	{"g_source_destroy", &SourceDestroy},
	{"g_source_get_can_recurse", &SourceGetCanRecurse},
	{"g_source_get_context", &SourceGetContext},
	{"g_source_get_current_time", &SourceGetCurrentTime},
	{"g_source_get_id", &SourceGetId},
	{"g_source_get_name", &SourceGetName},
	{"g_source_get_priority", &SourceGetPriority},
	{"g_source_get_time", &SourceGetTime},
	{"g_source_is_destroyed", &SourceIsDestroyed},
	{"g_source_new", &SourceNew},
	{"g_source_ref", &SourceRef},
	{"g_source_remove", &SourceRemove},
	{"g_source_remove_by_funcs_user_data", &SourceRemoveByFuncsUserData},
	{"g_source_remove_by_user_data", &SourceRemoveByUserData},
	{"g_source_remove_child_source", &SourceRemoveChildSource},
	{"g_source_remove_poll", &SourceRemovePoll},
	{"g_source_set_callback", &SourceSetCallback},
	{"g_source_set_callback_indirect", &SourceSetCallbackIndirect},
	{"g_source_set_can_recurse", &SourceSetCanRecurse},
	{"g_source_set_funcs", &SourceSetFuncs},
	{"g_source_set_name", &SourceSetName},
	{"g_source_set_name_by_id", &SourceSetNameById},
	{"g_source_set_priority", &SourceSetPriority},
	{"g_source_unref", &SourceUnref},
	{"g_spaced_primes_closest", &SpacedPrimesClosest},
	// Windows: _utf8 {"g_spawn_async", &SpawnAsync},
	{"g_spawn_async_utf8", &SpawnAsync},
	// Windows: _utf8 {"g_spawn_async_with_pipes", &SpawnAsyncWithPipes},
	{"g_spawn_async_with_pipes_utf8", &SpawnAsyncWithPipes},
	{"g_spawn_close_pid", &SpawnClosePid},
	// Windows: _utf8 {"g_spawn_command_line_async", &SpawnCommandLineAsync},
	{"g_spawn_command_line_async_utf8", &SpawnCommandLineAsync},
	// Windows: _utf8 {"g_spawn_command_line_sync", &SpawnCommandLineSync},
	{"g_spawn_command_line_sync_utf8", &SpawnCommandLineSync},
	{"g_spawn_error_quark", &SpawnErrorQuark},
	// Windows: _utf8 {"g_spawn_sync", &SpawnSync},
	{"g_spawn_sync_utf8", &SpawnSync},
	{"g_sprintf", &Sprintf},
	{"g_stat", &Stat},
	{"g_static_mutex_free", &StaticMutexFree},
	{"g_static_mutex_get_mutex_impl", &StaticMutexGetMutexImpl},
	{"g_static_mutex_init", &StaticMutexInit},
	{"g_static_private_free", &StaticPrivateFree},
	{"g_static_private_get", &StaticPrivateGet},
	{"g_static_private_init", &StaticPrivateInit},
	{"g_static_private_set", &StaticPrivateSet},
	{"g_static_rec_mutex_free", &StaticRecMutexFree},
	{"g_static_rec_mutex_init", &StaticRecMutexInit},
	{"g_static_rec_mutex_lock", &StaticRecMutexLock},
	{"g_static_rec_mutex_lock_full", &StaticRecMutexLockFull},
	{"g_static_rec_mutex_trylock", &StaticRecMutexTrylock},
	{"g_static_rec_mutex_unlock", &StaticRecMutexUnlock},
	{"g_static_rec_mutex_unlock_full", &StaticRecMutexUnlockFull},
	{"g_static_rw_lock_free", &StaticRwLockFree},
	{"g_static_rw_lock_init", &StaticRwLockInit},
	{"g_static_rw_lock_reader_lock", &StaticRwLockReaderLock},
	{"g_static_rw_lock_reader_trylock", &StaticRwLockReaderTrylock},
	{"g_static_rw_lock_reader_unlock", &StaticRwLockReaderUnlock},
	{"g_static_rw_lock_writer_lock", &StaticRwLockWriterLock},
	{"g_static_rw_lock_writer_trylock", &StaticRwLockWriterTrylock},
	{"g_static_rw_lock_writer_unlock", &StaticRwLockWriterUnlock},
	{"g_stpcpy", &Stpcpy},
	{"g_str_equal", &StrEqual},
	{"g_str_has_prefix", &StrHasPrefix},
	{"g_str_has_suffix", &StrHasSuffix},
	{"g_str_hash", &StrHash},
	{"g_strcanon", &Strcanon},
	// Deprecated 2.2 {"g_strcasecmp", &Strcasecmp},
	{"g_strchomp", &Strchomp},
	{"g_strchug", &Strchug},
	{"g_strcmp0", &Strcmp0},
	{"g_strcompress", &Strcompress},
	{"g_strconcat", &Strconcat},
	{"g_strdelimit", &Strdelimit},
	// Deprecated 2.2 {"g_strdown", &Strdown},
	{"g_strdup", &Strdup},
	{"g_strdup_printf", &StrdupPrintf},
	{"g_strdup_vprintf", &StrdupVprintf},
	{"g_strdupv", &Strdupv},
	{"g_strerror", &Strerror},
	{"g_strescape", &Strescape},
	{"g_strfreev", &Strfreev},
	{"g_string_append", &StringAppend},
	{"g_string_append_c", &StringAppendC},
	{"g_string_append_len", &StringAppendLen},
	{"g_string_append_printf", &StringAppendPrintf},
	{"g_string_append_unichar", &StringAppendUnichar},
	{"g_string_append_uri_escaped", &StringAppendUriEscaped},
	{"g_string_append_vprintf", &StringAppendVprintf},
	{"g_string_ascii_down", &StringAsciiDown},
	{"g_string_ascii_up", &StringAsciiUp},
	{"g_string_assign", &StringAssign},
	{"g_string_chunk_clear", &StringChunkClear},
	{"g_string_chunk_free", &StringChunkFree},
	{"g_string_chunk_insert", &StringChunkInsert},
	{"g_string_chunk_insert_const", &StringChunkInsertConst},
	{"g_string_chunk_insert_len", &StringChunkInsertLen},
	{"g_string_chunk_new", &StringChunkNew},
	{"g_string_down", &StringDown},
	{"g_string_equal", &StringEqual},
	{"g_string_erase", &StringErase},
	{"g_string_free", &StringFree},
	{"g_string_hash", &StringHash},
	{"g_string_insert", &StringInsert},
	{"g_string_insert_c", &StringInsertC},
	{"g_string_insert_len", &StringInsertLen},
	{"g_string_insert_unichar", &StringInsertUnichar},
	{"g_string_new", &StringNew},
	{"g_string_new_len", &StringNewLen},
	{"g_string_overwrite", &StringOverwrite},
	{"g_string_overwrite_len", &StringOverwriteLen},
	{"g_string_prepend", &StringPrepend},
	{"g_string_prepend_c", &StringPrependC},
	{"g_string_prepend_len", &StringPrependLen},
	{"g_string_prepend_unichar", &StringPrependUnichar},
	{"g_string_printf", &StringPrintf},
	{"g_string_set_size", &StringSetSize},
	{"g_string_sized_new", &StringSizedNew},
	{"g_string_truncate", &StringTruncate},
	{"g_string_up", &StringUp},
	{"g_string_vprintf", &StringVprintf},
	{"g_strip_context", &StripContext},
	{"g_strjoin", &Strjoin},
	{"g_strjoinv", &Strjoinv},
	{"g_strlcat", &Strlcat},
	{"g_strlcpy", &Strlcpy},
	// Deprecated 2.2 {"g_strncasecmp", &Strncasecmp},
	{"g_strndup", &Strndup},
	{"g_strnfill", &Strnfill},
	{"g_strreverse", &Strreverse},
	{"g_strrstr", &Strrstr},
	{"g_strrstr_len", &StrrstrLen},
	{"g_strsignal", &Strsignal},
	{"g_strsplit", &Strsplit},
	{"g_strsplit_set", &StrsplitSet},
	{"g_strstr_len", &StrstrLen},
	{"g_strtod", &Strtod},
	// Deprecated 2.2 {"g_strup", &Strup},
	{"g_strv_length", &StrvLength},
	{"g_test_add_data_func", &TestAddDataFunc},
	{"g_test_add_func", &TestAddFunc},
	{"g_test_add_vtable", &TestAddVtable},
	{"g_test_bug", &TestBug},
	{"g_test_bug_base", &TestBugBase},
	{"g_test_create_case", &TestCreateCase},
	{"g_test_create_suite", &TestCreateSuite},
	{"g_test_get_root", &TestGetRoot},
	{"g_test_init", &TestInit},
	{"g_test_log_buffer_free", &TestLogBufferFree},
	{"g_test_log_buffer_new", &TestLogBufferNew},
	{"g_test_log_buffer_pop", &TestLogBufferPop},
	{"g_test_log_buffer_push", &TestLogBufferPush},
	{"g_test_log_msg_free", &TestLogMsgFree},
	{"g_test_log_set_fatal_handler", &TestLogSetFatalHandler},
	{"g_test_log_type_name", &TestLogTypeName},
	{"g_test_maximized_result", &TestMaximizedResult},
	{"g_test_message", &TestMessage},
	{"g_test_minimized_result", &TestMinimizedResult},
	{"g_test_queue_destroy", &TestQueueDestroy},
	{"g_test_queue_free", &TestQueueFree},
	{"g_test_rand_double", &TestRandDouble},
	{"g_test_rand_double_range", &TestRandDoubleRange},
	{"g_test_rand_int", &TestRandInt},
	{"g_test_rand_int_range", &TestRandIntRange},
	{"g_test_run", &TestRun},
	{"g_test_run_suite", &TestRunSuite},
	{"g_test_suite_add", &TestSuiteAdd},
	{"g_test_suite_add_suite", &TestSuiteAddSuite},
	{"g_test_timer_elapsed", &TestTimerElapsed},
	{"g_test_timer_last", &TestTimerLast},
	{"g_test_timer_start", &TestTimerStart},
	{"g_test_trap_assertions", &TestTrapAssertions},
	{"g_test_trap_fork", &TestTrapFork},
	{"g_test_trap_has_passed", &TestTrapHasPassed},
	{"g_test_trap_reached_timeout", &TestTrapReachedTimeout},
	{"g_thread_create_full", &ThreadCreateFull},
	{"g_thread_error_quark", &ThreadErrorQuark},
	{"g_thread_exit", &ThreadExit},
	{"g_thread_foreach", &ThreadForeach},
	{"g_thread_get_initialized", &ThreadGetInitialized},
	// Undocumented {"g_thread_init_glib", &ThreadInitGlib},
	{"g_thread_join", &ThreadJoin},
	{"g_thread_pool_free", &ThreadPoolFree},
	{"g_thread_pool_get_max_idle_time", &ThreadPoolGetMaxIdleTime},
	{"g_thread_pool_get_max_threads", &ThreadPoolGetMaxThreads},
	{"g_thread_pool_get_max_unused_threads", &ThreadPoolGetMaxUnusedThreads},
	{"g_thread_pool_get_num_threads", &ThreadPoolGetNumThreads},
	{"g_thread_pool_get_num_unused_threads", &ThreadPoolGetNumUnusedThreads},
	{"g_thread_pool_new", &ThreadPoolNew},
	{"g_thread_pool_push", &ThreadPoolPush},
	{"g_thread_pool_set_max_idle_time", &ThreadPoolSetMaxIdleTime},
	{"g_thread_pool_set_max_threads", &ThreadPoolSetMaxThreads},
	{"g_thread_pool_set_max_unused_threads", &ThreadPoolSetMaxUnusedThreads},
	{"g_thread_pool_set_sort_function", &ThreadPoolSetSortFunction},
	{"g_thread_pool_stop_unused_threads", &ThreadPoolStopUnusedThreads},
	{"g_thread_pool_unprocessed", &ThreadPoolUnprocessed},
	{"g_thread_self", &ThreadSelf},
	{"g_thread_set_priority", &ThreadSetPriority},
	{"g_time_val_add", &TimeValAdd},
	{"g_time_val_from_iso8601", &TimeValFromIso8601},
	{"g_time_val_to_iso8601", &TimeValToIso8601},
	{"g_time_zone_adjust_time", &TimeZoneAdjustTime},
	{"g_time_zone_find_interval", &TimeZoneFindInterval},
	{"g_time_zone_get_abbreviation", &TimeZoneGetAbbreviation},
	{"g_time_zone_get_offset", &TimeZoneGetOffset},
	{"g_time_zone_is_dst", &TimeZoneIsDst},
	{"g_time_zone_new", &TimeZoneNew},
	{"g_time_zone_new_local", &TimeZoneNewLocal},
	{"g_time_zone_new_utc", &TimeZoneNewUtc},
	{"g_time_zone_ref", &TimeZoneRef},
	{"g_time_zone_unref", &TimeZoneUnref},
	{"g_timeout_add", &TimeoutAdd},
	{"g_timeout_add_full", &TimeoutAddFull},
	{"g_timeout_add_seconds", &TimeoutAddSeconds},
	{"g_timeout_add_seconds_full", &TimeoutAddSecondsFull},
	{"g_timeout_source_new", &TimeoutSourceNew},
	{"g_timeout_source_new_seconds", &TimeoutSourceNewSeconds},
	{"g_timer_continue", &TimerContinue},
	{"g_timer_destroy", &TimerDestroy},
	{"g_timer_elapsed", &TimerElapsed},
	{"g_timer_new", &TimerNew},
	{"g_timer_reset", &TimerReset},
	{"g_timer_start", &TimerStart},
	{"g_timer_stop", &TimerStop},
	{"g_trash_stack_height", &TrashStackHeight},
	{"g_trash_stack_peek", &TrashStackPeek},
	{"g_trash_stack_pop", &TrashStackPop},
	{"g_trash_stack_push", &TrashStackPush},
	{"g_tree_destroy", &TreeDestroy},
	{"g_tree_foreach", &TreeForeach},
	{"g_tree_height", &TreeHeight},
	{"g_tree_insert", &TreeInsert},
	{"g_tree_lookup", &TreeLookup},
	{"g_tree_lookup_extended", &TreeLookupExtended},
	{"g_tree_new", &TreeNew},
	{"g_tree_new_full", &TreeNewFull},
	{"g_tree_new_with_data", &TreeNewWithData},
	{"g_tree_nnodes", &TreeNnodes},
	{"g_tree_ref", &TreeRef},
	{"g_tree_remove", &TreeRemove},
	{"g_tree_replace", &TreeReplace},
	{"g_tree_search", &TreeSearch},
	{"g_tree_steal", &TreeSteal},
	{"g_tree_traverse", &TreeTraverse},
	{"g_tree_unref", &TreeUnref},
	{"g_try_malloc", &TryMalloc},
	{"g_try_malloc0", &TryMalloc0},
	{"g_try_malloc0_n", &TryMalloc0N},
	{"g_try_malloc_n", &TryMallocN},
	{"g_try_realloc", &TryRealloc},
	{"g_try_realloc_n", &TryReallocN},
	{"g_tuples_destroy", &TuplesDestroy},
	{"g_tuples_index", &TuplesIndex},
	{"g_ucs4_to_utf16", &Ucs4ToUtf16},
	{"g_ucs4_to_utf8", &Ucs4ToUtf8},
	{"g_unichar_break_type", &UnicharBreakType},
	{"g_unichar_combining_class", &UnicharCombiningClass},
	{"g_unichar_digit_value", &UnicharDigitValue},
	{"g_unichar_get_mirror_char", &UnicharGetMirrorChar},
	{"g_unichar_get_script", &UnicharGetScript},
	{"g_unichar_isalnum", &UnicharIsalnum},
	{"g_unichar_isalpha", &UnicharIsalpha},
	{"g_unichar_iscntrl", &UnicharIscntrl},
	{"g_unichar_isdefined", &UnicharIsdefined},
	{"g_unichar_isdigit", &UnicharIsdigit},
	{"g_unichar_isgraph", &UnicharIsgraph},
	{"g_unichar_islower", &UnicharIslower},
	{"g_unichar_ismark", &UnicharIsmark},
	{"g_unichar_isprint", &UnicharIsprint},
	{"g_unichar_ispunct", &UnicharIspunct},
	{"g_unichar_isspace", &UnicharIsspace},
	{"g_unichar_istitle", &UnicharIstitle},
	{"g_unichar_isupper", &UnicharIsupper},
	{"g_unichar_iswide", &UnicharIswide},
	{"g_unichar_iswide_cjk", &UnicharIswideCjk},
	{"g_unichar_isxdigit", &UnicharIsxdigit},
	{"g_unichar_iszerowidth", &UnicharIszerowidth},
	{"g_unichar_to_utf8", &UnicharToUtf8},
	{"g_unichar_tolower", &UnicharTolower},
	{"g_unichar_totitle", &UnicharTotitle},
	{"g_unichar_toupper", &UnicharToupper},
	{"g_unichar_type", &UnicharType},
	{"g_unichar_validate", &UnicharValidate},
	{"g_unichar_xdigit_value", &UnicharXdigitValue},
	{"g_unicode_canonical_decomposition", &UnicodeCanonicalDecomposition},
	{"g_unicode_canonical_ordering", &UnicodeCanonicalOrdering},
	{"g_unlink", &Unlink},
	// Windows: _utf8 {"g_unsetenv", &Unsetenv},
	{"g_unsetenv_utf8", &Unsetenv},
	{"g_uri_escape_string", &UriEscapeString},
	{"g_uri_list_extract_uris", &UriListExtractUris},
	{"g_uri_parse_scheme", &UriParseScheme},
	{"g_uri_unescape_segment", &UriUnescapeSegment},
	{"g_uri_unescape_string", &UriUnescapeString},
	{"g_usleep", &Usleep},
	{"g_utf16_to_ucs4", &Utf16ToUcs4},
	{"g_utf16_to_utf8", &Utf16ToUtf8},
	{"g_utf8_casefold", &Utf8Casefold},
	{"g_utf8_collate", &Utf8Collate},
	{"g_utf8_collate_key", &Utf8CollateKey},
	{"g_utf8_collate_key_for_filename", &Utf8CollateKeyForFilename},
	{"g_utf8_find_next_char", &Utf8FindNextChar},
	{"g_utf8_find_prev_char", &Utf8FindPrevChar},
	{"g_utf8_get_char", &Utf8GetChar},
	{"g_utf8_get_char_validated", &Utf8GetCharValidated},
	{"g_utf8_normalize", &Utf8Normalize},
	{"g_utf8_offset_to_pointer", &Utf8OffsetToPointer},
	{"g_utf8_pointer_to_offset", &Utf8PointerToOffset},
	{"g_utf8_prev_char", &Utf8PrevChar},
	{"g_utf8_strchr", &Utf8Strchr},
	{"g_utf8_strdown", &Utf8Strdown},
	{"g_utf8_strlen", &Utf8Strlen},
	{"g_utf8_strncpy", &Utf8Strncpy},
	{"g_utf8_strrchr", &Utf8Strrchr},
	{"g_utf8_strreverse", &Utf8Strreverse},
	{"g_utf8_strup", &Utf8Strup},
	{"g_utf8_to_ucs4", &Utf8ToUcs4},
	{"g_utf8_to_ucs4_fast", &Utf8ToUcs4Fast},
	{"g_utf8_to_utf16", &Utf8ToUtf16},
	{"g_utf8_validate", &Utf8Validate},
	{"g_utime", &Utime},
	{"g_variant_builder_add", &VariantBuilderAdd},
	{"g_variant_builder_add_parsed", &VariantBuilderAddParsed},
	{"g_variant_builder_add_value", &VariantBuilderAddValue},
	{"g_variant_builder_clear", &VariantBuilderClear},
	{"g_variant_builder_close", &VariantBuilderClose},
	{"g_variant_builder_end", &VariantBuilderEnd},
	{"g_variant_builder_init", &VariantBuilderInit},
	{"g_variant_builder_new", &VariantBuilderNew},
	{"g_variant_builder_open", &VariantBuilderOpen},
	{"g_variant_builder_ref", &VariantBuilderRef},
	{"g_variant_builder_unref", &VariantBuilderUnref},
	{"g_variant_byteswap", &VariantByteswap},
	{"g_variant_classify", &VariantClassify},
	{"g_variant_compare", &VariantCompare},
	{"g_variant_dup_bytestring", &VariantDupBytestring},
	{"g_variant_dup_bytestring_array", &VariantDupBytestringArray},
	{"g_variant_dup_string", &VariantDupString},
	{"g_variant_dup_strv", &VariantDupStrv},
	{"g_variant_equal", &VariantEqual},
	// Undocumented {"g_variant_format_string_scan", &VariantFormatStringScan},
	// Undocumented {"g_variant_format_string_scan_type", &VariantFormatStringScanType},
	{"g_variant_get", &VariantGet},
	{"g_variant_get_boolean", &VariantGetBoolean},
	{"g_variant_get_byte", &VariantGetByte},
	{"g_variant_get_bytestring", &VariantGetBytestring},
	{"g_variant_get_bytestring_array", &VariantGetBytestringArray},
	{"g_variant_get_child", &VariantGetChild},
	{"g_variant_get_child_value", &VariantGetChildValue},
	{"g_variant_get_data", &VariantGetData},
	{"g_variant_get_double", &VariantGetDouble},
	{"g_variant_get_fixed_array", &VariantGetFixedArray},
	{"g_variant_get_handle", &VariantGetHandle},
	{"g_variant_get_int16", &VariantGetInt16},
	{"g_variant_get_int32", &VariantGetInt32},
	{"g_variant_get_int64", &VariantGetInt64},
	{"g_variant_get_maybe", &VariantGetMaybe},
	{"g_variant_get_normal_form", &VariantGetNormalForm},
	{"g_variant_get_size", &VariantGetSize},
	{"g_variant_get_string", &VariantGetString},
	{"g_variant_get_strv", &VariantGetStrv},
	{"g_variant_get_type", &VariantGetType},
	{"g_variant_get_type_string", &VariantGetTypeString},
	{"g_variant_get_uint16", &VariantGetUint16},
	{"g_variant_get_uint32", &VariantGetUint32},
	{"g_variant_get_uint64", &VariantGetUint64},
	{"g_variant_get_va", &VariantGetVa},
	{"g_variant_get_variant", &VariantGetVariant},
	{"g_variant_hash", &VariantHash},
	{"g_variant_is_container", &VariantIsContainer},
	{"g_variant_is_floating", &VariantIsFloating},
	{"g_variant_is_normal_form", &VariantIsNormalForm},
	{"g_variant_is_object_path", &VariantIsObjectPath},
	{"g_variant_is_of_type", &VariantIsOfType},
	{"g_variant_is_signature", &VariantIsSignature},
	{"g_variant_iter_copy", &VariantIterCopy},
	{"g_variant_iter_free", &VariantIterFree},
	{"g_variant_iter_init", &VariantIterInit},
	{"g_variant_iter_loop", &VariantIterLoop},
	{"g_variant_iter_n_children", &VariantIterNChildren},
	{"g_variant_iter_new", &VariantIterNew},
	{"g_variant_iter_next", &VariantIterNext},
	{"g_variant_iter_next_value", &VariantIterNextValue},
	{"g_variant_lookup", &VariantLookup},
	{"g_variant_lookup_value", &VariantLookupValue},
	{"g_variant_n_children", &VariantNChildren},
	{"g_variant_new", &VariantNew},
	{"g_variant_new_array", &VariantNewArray},
	{"g_variant_new_boolean", &VariantNewBoolean},
	{"g_variant_new_byte", &VariantNewByte},
	{"g_variant_new_bytestring", &VariantNewBytestring},
	{"g_variant_new_bytestring_array", &VariantNewBytestringArray},
	{"g_variant_new_dict_entry", &VariantNewDictEntry},
	{"g_variant_new_double", &VariantNewDouble},
	{"g_variant_new_from_data", &VariantNewFromData},
	{"g_variant_new_handle", &VariantNewHandle},
	{"g_variant_new_int16", &VariantNewInt16},
	{"g_variant_new_int32", &VariantNewInt32},
	{"g_variant_new_int64", &VariantNewInt64},
	{"g_variant_new_maybe", &VariantNewMaybe},
	{"g_variant_new_object_path", &VariantNewObjectPath},
	{"g_variant_new_parsed", &VariantNewParsed},
	{"g_variant_new_parsed_va", &VariantNewParsedVa},
	{"g_variant_new_signature", &VariantNewSignature},
	{"g_variant_new_string", &VariantNewString},
	{"g_variant_new_strv", &VariantNewStrv},
	{"g_variant_new_tuple", &VariantNewTuple},
	{"g_variant_new_uint16", &VariantNewUint16},
	{"g_variant_new_uint32", &VariantNewUint32},
	{"g_variant_new_uint64", &VariantNewUint64},
	{"g_variant_new_va", &VariantNewVa},
	{"g_variant_new_variant", &VariantNewVariant},
	{"g_variant_parse", &VariantParse},
	{"g_variant_parser_get_error_quark", &VariantParserGetErrorQuark},
	{"g_variant_print", &VariantPrint},
	{"g_variant_print_string", &VariantPrintString},
	{"g_variant_ref", &VariantRef},
	{"g_variant_ref_sink", &VariantRefSink},
	// Undocumented {"g_variant_serialised_byteswap", &VariantSerialisedByteswap},
	// Undocumented {"g_variant_serialised_get_child", &VariantSerialisedGetChild},
	// Undocumented {"g_variant_serialised_is_normal", &VariantSerialisedIsNormal},
	// Undocumented {"g_variant_serialised_n_children", &VariantSerialisedNChildren},
	// Undocumented {"g_variant_serialiser_is_object_path", &VariantSerialiserIsObjectPath},
	// Undocumented {"g_variant_serialiser_is_signature", &VariantSerialiserIsSignature},
	// Undocumented {"g_variant_serialiser_is_string", &VariantSerialiserIsString},
	// Undocumented {"g_variant_serialiser_needed_size", &VariantSerialiserNeededSize},
	// Undocumented {"g_variant_serialiser_serialise", &VariantSerialiserSerialise},
	{"g_variant_store", &VariantStore},
	{"g_variant_type_checked_", &VariantTypeChecked},
	{"g_variant_type_copy", &VariantTypeCopy},
	{"g_variant_type_dup_string", &VariantTypeDupString},
	{"g_variant_type_element", &VariantTypeElement},
	{"g_variant_type_equal", &VariantTypeEqual},
	{"g_variant_type_first", &VariantTypeFirst},
	{"g_variant_type_free", &VariantTypeFree},
	{"g_variant_type_get_string_length", &VariantTypeGetStringLength},
	{"g_variant_type_hash", &VariantTypeHash},
	// Undocumented {"g_variant_type_info_assert_no_infos", &VariantTypeInfoAssertNoInfos},
	// Undocumented {"g_variant_type_info_element", &VariantTypeInfoElement},
	// Undocumented {"g_variant_type_info_get", &VariantTypeInfoGet},
	// Undocumented {"g_variant_type_info_get_type_string", &VariantTypeInfoGetTypeString},
	// Undocumented {"g_variant_type_info_member_info", &VariantTypeInfoMemberInfo},
	// Undocumented {"g_variant_type_info_n_members", &VariantTypeInfoNMembers},
	// Undocumented {"g_variant_type_info_query", &VariantTypeInfoQuery},
	// Undocumented {"g_variant_type_info_query_element", &VariantTypeInfoQueryElement},
	// Undocumented {"g_variant_type_info_ref", &VariantTypeInfoRef},
	// Undocumented {"g_variant_type_info_unref", &VariantTypeInfoUnref},
	{"g_variant_type_is_array", &VariantTypeIsArray},
	{"g_variant_type_is_basic", &VariantTypeIsBasic},
	{"g_variant_type_is_container", &VariantTypeIsContainer},
	{"g_variant_type_is_definite", &VariantTypeIsDefinite},
	{"g_variant_type_is_dict_entry", &VariantTypeIsDictEntry},
	{"g_variant_type_is_maybe", &VariantTypeIsMaybe},
	{"g_variant_type_is_subtype_of", &VariantTypeIsSubtypeOf},
	{"g_variant_type_is_tuple", &VariantTypeIsTuple},
	{"g_variant_type_is_variant", &VariantTypeIsVariant},
	{"g_variant_type_key", &VariantTypeKey},
	{"g_variant_type_n_items", &VariantTypeNItems},
	{"g_variant_type_new", &VariantTypeNew},
	{"g_variant_type_new_array", &VariantTypeNewArray},
	{"g_variant_type_new_dict_entry", &VariantTypeNewDictEntry},
	{"g_variant_type_new_maybe", &VariantTypeNewMaybe},
	{"g_variant_type_new_tuple", &VariantTypeNewTuple},
	{"g_variant_type_next", &VariantTypeNext},
	{"g_variant_type_peek_string", &VariantTypePeekString},
	{"g_variant_type_string_is_valid", &VariantTypeStringIsValid},
	{"g_variant_type_string_scan", &VariantTypeStringScan},
	{"g_variant_type_value", &VariantTypeValue},
	{"g_variant_unref", &VariantUnref},
	{"g_vasprintf", &Vasprintf},
	{"g_vfprintf", &Vfprintf},
	{"g_vprintf", &Vprintf},
	{"g_vsnprintf", &Vsnprintf},
	{"g_vsprintf", &Vsprintf},
	{"g_warn_message", &WarnMessage},
	{"g_win32_error_message", &Win32ErrorMessage},
	{"g_win32_ftruncate", &Win32Ftruncate},
	// Windows: _utf8 {"g_win32_get_package_installation_directory", &Win32GetPackageInstallationDirectory},
	{"g_win32_get_package_installation_directory_of_module", &Win32GetPackageInstallationDirectoryOfModule},
	{"g_win32_get_package_installation_directory_utf8", &Win32GetPackageInstallationDirectory},
	// Windows: _utf8 {"g_win32_get_package_installation_subdirectory", &Win32GetPackageInstallationSubdirectory},
	{"g_win32_get_package_installation_subdirectory_utf8", &Win32GetPackageInstallationSubdirectory},
	{"g_win32_get_system_data_dirs_for_module", &Win32GetSystemDataDirsForModule},
	{"g_win32_get_windows_version", &Win32GetWindowsVersion},
	{"g_win32_getlocale", &Win32Getlocale},
	{"g_win32_locale_filename_from_utf8", &Win32LocaleFilenameFromUtf8},
	{"glib_check_version", &GlibCheckVersion},
	// Undocumented {"glib_gettext", &GlibGettext},
	// Undocumented {"glib_on_error_halt", &GlibOnErrorHalt},
	// Undocumented {"glib_pgettext", &GlibPgettext},
}

var dllThread = "libgthread-2.0-0.dll"

var apiListThread = outside.Apis{
	{"g_thread_init", &ThreadInit},
	{"g_thread_init_with_errorcheck_mutexes", &ThreadInitWithErrorcheckMutexes},
}

var dataList = outside.Data{
	{"g_ascii_table", (*uint16)(nil)},
	{"g_child_watch_funcs", (*O.SourceFuncs)(nil)},
	{"g_idle_funcs", (*O.SourceFuncs)(nil)},
	{"g_io_watch_funcs", (*O.SourceFuncs)(nil)},
	{"g_mem_gc_friendly", (*T.Gboolean)(nil)},
	{"g_test_config_vars", (*T.GTestConfig)(nil)},
	{"g_thread_functions_for_glib_use", (*ThreadFunctions)(nil)},
	{"g_thread_gettime", (*T.G_thread_gettime)(nil)},
	{"g_thread_use_default_impl", (*T.Gboolean)(nil)},
	{"g_threads_got_initialized", (*T.Gboolean)(nil)},
	{"g_timeout_funcs", (*O.SourceFuncs)(nil)},
	{"g_utf8_skip", (*T.Gchar)(nil)},
	{"glib_binary_age", (*uint)(nil)},
	{"glib_interface_age", (*uint)(nil)},
	{"glib_major_version", (*uint)(nil)},
	{"glib_mem_profiler_table", (*T.GMemVTable)(nil)},
	{"glib_micro_version", (*uint)(nil)},
	{"glib_minor_version", (*uint)(nil)},
}
