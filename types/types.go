// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package types provides API structs, enums and types for Gtk+2.
package types

type (
	fix uintptr

	Long_double fix
	LongDouble  fix
	Va_list     fix
	VaList      fix
	FILE        fix
	Stat        fix

	CairoBool     int
	Char          int8
	Enum          int
	FT_Bool       uint8
	FT_Byte       uint8
	FT_Char       int8
	FT_Error      int
	FT_Int32      int   // ANOMALLY int != int32
	FT_Long       int32 //TODO(t): signed long SIZE ?
	FT_String     int8
	FT_UInt32     uint   // ANOMALLY unsigned int != uint32
	FT_ULong      uint32 //unsigned long SIZE ?
	Gchar         int8
	Gconstpointer *struct{}
	GInt32        int   // ANOMALLY size?
	Glong         int32 //TODO(t): CHECK
	Gpointer      *struct{}
	Gsize         uint
	Gssize        int
	Guchar        uint8
	GUint32       uint  // ANOMALLY size?
	Long          int32 // TODO(t): Size?
	Size_t        uintptr
	SizeT         uintptr
	Time_t        int
	TimeT         int
	UnsignedChar  uint8
	UnsignedLong  uint32 //TODO(t): check  size

	AtkAttributeSet        GSList
	FT_Angle               FT_Fixed
	FT_Bytes               *FT_Byte
	FT_CharMap             *FT_CharMapRec
	FT_Driver              *FT_DriverRec
	FT_F26Dot6             FT_Long
	FTFace                 *FT_FaceRec
	FT_Face_Internal       *FT_Face_InternalRec
	FT_Fixed               FT_Long
	FT_Glyph               *FT_GlyphRec
	FT_GlyphSlot           *FT_GlyphSlotRec
	FT_Library             *FT_LibraryRec
	FT_List                *FT_ListRec
	FT_ListNode            *FT_ListNodeRec
	FT_Memory              *FT_MemoryRec
	FT_Module              *FT_ModuleRec
	FT_Module_Interface    FT_Pointer
	FT_Pointer             *Void
	FT_Pos                 FT_Long
	FT_Renderer            *FT_RendererRec
	FT_Size                *FT_SizeRec
	FT_Size_Internal       *FT_Size_InternalRec
	FT_Size_Request        *FT_Size_RequestRec
	FT_Slot_Internal       *FT_Slot_InternalRec
	FT_Stream              *FT_StreamRec
	FT_Stroker             *FT_StrokerRec
	FT_SubGlyph            *FT_SubGlyphRec
	FTC_CMapCache          *FTC_CMapCacheRec
	FTC_FaceID             FT_Pointer
	FTC_Font               *FTC_FontRec
	FTC_ImageCache         *FTC_ImageCacheRec
	FTC_ImageType          *FTC_ImageTypeRec
	FTC_Manager            *FTC_ManagerRec
	FTC_Node               *FTC_NodeRec
	FTC_SBit               *FTC_SBitRec
	FTC_SBitCache          *FTC_SBitCacheRec
	FTC_Scaler             *FTC_ScalerRec
	Gboolean               int
	GCacheDestroyFunc      func(value Gpointer)
	GCacheDupFunc          func(value Gpointer) Gpointer
	GCacheNewFunc          func(key Gpointer) Gpointer
	GDateDay               uint8
	GDateTime              GInt32
	GDateYear              uint16
	GdkAtom                *struct{} // REMOVE
	GdkBitmap              GdkDrawable
	GdkNativeWindow        Gpointer
	GdkPixmap              GdkDrawable
	GdkWChar               GUint32
	GdkWindow              GdkDrawable // REMOVE
	GInitiallyUnowned      GObject
	GObjectClass           struct{} //REMOVE
	GInitiallyUnownedClass GObjectClass
	Goffset                int64
	GPid                   *struct{}
	GQuark                 GUint32
	GSequenceIter          GSequenceNode
	GSignalCMarshaller     GClosureMarshal
	GStaticMutex           *GMutex
	GTime                  GInt32
	GTimeSpan              int64
	GtkType                GType // REMOVE
	GType                  Gsize // REMOVE
	Gulong                 UnsignedLong
	Gunichar               GUint32
	Gunichar2              uint16
	GVoidFunc              func()
	PangoGlyph             GUint32
	PangoGlyphUnit         GInt32
	PangoLayoutRun         PangoGlyphItem
	PangoOTTag             GUint32
	PS_FontInfo            *PS_FontInfoRec
	PS_Private             *PS_PrivateRec

	simpleObject struct{ parent GObject }

	AtkHyperlink     simpleObject
	AtkMisc          simpleObject
	AtkObjectFactory simpleObject
	AtkStateSet      simpleObject
	GdkDrawable      simpleObject //REMOVE
	PangoEngine      simpleObject
	PangoFcDecoder   simpleObject

	//TODO(t): Check and Fix
	FT_DriverRec        struct{}
	FT_Face_InternalRec struct{}
	FT_LibraryRec       struct{}
	FT_ModuleRec        struct{}
	FT_RendererRec      struct{}
	FT_Size_InternalRec struct{}
	FT_Slot_InternalRec struct{}
	FT_StreamRec        struct{}
	FT_StrokerRec       struct{}
	FT_SubGlyphRec      struct{}
	FTC_CMapCacheRec    struct{}
	FTC_ImageCacheRec   struct{}
	FTC_ManagerRec      struct{}
	FTC_NodeRec         struct{}
	FTC_SBitCacheRec    struct{}

	AtkAction                      struct{}
	AtkComponent                   struct{}
	AtkDocument                    struct{}
	AtkEditableText                struct{}
	AtkHyperlinkImpl               struct{}
	AtkHypertext                   struct{}
	AtkImage                       struct{}
	AtkImplementor                 struct{}
	AtkSelection                   struct{}
	AtkSocket                      struct{}
	AtkState                       struct{}
	AtkStreamableContent           struct{}
	AtkTable                       struct{}
	AtkText                        struct{}
	AtkValue                       struct{}
	CairoDevice                    struct{}
	CairoFontFace                  struct{}
	CairoFontOptions               struct{}
	CairoPattern                   struct{}
	CairoRegion                    struct{}
	CairoScaledFont                struct{}
	CairoScriptInterpreter         struct{}
	CairoSurface                   struct{}
	Cairo                          struct{}
	FcCharSet                      struct{}
	FcPattern                      struct{}
	GAction                        struct{}
	GActionGroup                   struct{}
	GAllocator                     struct{}
	GAppInfo                       struct{}
	GAppLaunchContextPrivate       struct{}
	GApplicationCommandLinePrivate struct{}
	GApplicationPrivate            struct{}
	GAsyncInitable                 struct{}
	GAsyncQueue                    struct{}
	GAsyncResult                   struct{}
	GBookmarkFile                  struct{}
	GBufferedInputStreamPrivate    struct{}
	GBufferedOutputStreamPrivate   struct{}
	GCache                         struct{}
	GCancellablePrivate            struct{}
	GcharsetConverter              struct{}
	GChecksum                      struct{}
	GCond                          struct{}
	GConverter                     struct{}
	GConverterInputStreamPrivate   struct{}
	GConverterOutputStreamPrivate  struct{}
	GCredentials                   struct{}
	GData                          struct{}
	GDataInputStreamPrivate        struct{}
	GDataOutputStreamPrivate       struct{}
	GDBusAuthObserver              struct{}
	GDBusConnection                struct{}
	GDBusMessage                   struct{}
	GDBusMethodInvocation          struct{}
	GDBusProxyPrivate              struct{}
	GDBusServer                    struct{}
	GDir                           struct{}
	GdkPangoRendererPrivate        struct{}
	GdkXEvent                      struct{}
	GDrive                         struct{}
	GEmblem                        struct{}
	GEmblemedIconPrivate           struct{}
	GFile                          struct{}
	GFileAttributeMatcher          struct{}
	GFileEnumeratorPrivate         struct{}
	GFileIcon                      struct{}
	GFileInfo                      struct{}
	GFileInputStreamPrivate        struct{}
	GFileIOStreamPrivate           struct{}
	GFileMonitorPrivate            struct{}
	GFilenameCompleter             struct{}
	GFileOutputStreamPrivate       struct{}
	GHashTable                     struct{}
	GIcon                          struct{}
	GIConv                         struct{}
	GInetAddressPrivate            struct{}
	GInetSocketAddressPrivate      struct{}
	GInitable                      struct{}
	GInputStreamPrivate            struct{}
	GIOExtension                   struct{}
	GIOExtensionPoint              struct{}
	GIOModule                      struct{}
	GIOSchedulerJob                struct{}
	GIOStreamPrivate               struct{}
	GKeyFile                       struct{}
	GLoadableIcon                  struct{}
	GMainContext                   struct{}
	GMainLoop                      struct{}
	GMappedFile                    struct{}
	GMarkupParseContext            struct{}
	GMatchInfo                     struct{}
	GMemChunk                      struct{}
	GMemoryInputStreamPrivate      struct{}
	GMemoryOutputStreamPrivate     struct{}
	GModule                        struct{}
	GMount                         struct{}
	GMountOperationPrivate         struct{}
	GMutex                         struct{}
	GNetworkAddressPrivate         struct{}
	GNetworkServicePrivate         struct{}
	GOptionContext                 struct{}
	GOptionGroup                   struct{}
	GOutputStreamPrivate           struct{}
	GPatternSpec                   struct{}
	GPermissionPrivate             struct{}
	GPollableInputStream           struct{}
	GPollableOutputStream          struct{}
	GPrivate                       struct{}
	GProxy                         struct{}
	GProxyAddressPrivate           struct{}
	GProxyResolver                 struct{}
	GRand                          struct{}
	GRegex                         struct{}
	GRelation                      struct{}
	GResolverPrivate               struct{}
	GSeekable                      struct{}
	GSequence                      struct{}
	GSequenceNode                  struct{}
	GSettingsBackend               struct{}
	GSettingsPrivate               struct{}
	GSimpleActionGroupPrivate      struct{}
	GSimpleActionPrivate           struct{}
	GSimpleAsyncResult             struct{}
	GSocketClientPrivate           struct{}
	GSocketConnectable             struct{}
	GSocketConnectionPrivate       struct{}
	GSocketControlMessagePrivate   struct{}
	GSocketListenerPrivate         struct{}
	GSocketPrivate                 struct{}
	GSocketServicePrivate          struct{}
	GSrvTarget                     struct{}
	GStringChunk                   struct{}
	GTcpConnectionPrivate          struct{}
	GTcpWrapperConnectionPrivate   struct{}
	GTestCase                      struct{}
	GTestSuite                     struct{}
	GThemedIcon                    struct{}
	GTimer                         struct{}
	GTimeZone                      struct{}
	GtkActionGroupPrivate          struct{}
	GtkActionPrivate               struct{}
	GtkAssistantPrivate            struct{}
	GtkBuilderPrivate              struct{}
	GtkCalendarPrivate             struct{}
	GtkCellViewPrivate             struct{}
	GtkClipboard                   struct{}
	GtkColorButtonPrivate          struct{}
	GtkComboBoxEntryPrivate        struct{}
	GtkComboBoxPrivate             struct{}
	GtkComboBoxTextPrivate         struct{}
	GtkExpanderPrivate             struct{}
	GtkFontButtonPrivate           struct{}
	GtkIconThemePrivate            struct{}
	GtkIconViewPrivate             struct{}
	GtkIMMulticontextPrivate       struct{}
	GtkInfoBarPrivate              struct{}
	GtkLabelSelectionInfo          struct{}
	GtkPageSetup                   struct{} //REMOVE
	GtkRecentActionPrivate         struct{}
	GtkRecentChooserMenuPrivate    struct{}
	GtkRecentManagerPrivate        struct{}
	GtkScaleButtonPrivate          struct{}
	GtkSeparatorToolItemPrivate    struct{}
	GtkSpinnerPrivate              struct{}
	GtkStatusIconPrivate           struct{}
	GtkToggleActionPrivate         struct{}
	GtkToggleToolButtonPrivate     struct{}
	GtkToolButtonPrivate           struct{}
	GtkToolItemGroupPrivate        struct{}
	GtkToolItemPrivate             struct{}
	GtkToolPalettePrivate          struct{}
	GtkUIManagerPrivate            struct{}
	GTlsBackend                    struct{}
	GTlsCertificatePrivate         struct{}
	GTlsClientConnection           struct{}
	GTlsConnectionPrivate          struct{}
	GTree                          struct{}
	GTypeCValue                    struct{}
	GUnixFDList                    struct{}
	GVariant                       struct{}
	GVariantType                   struct{}
	GVolume                        struct{}
	GWin32InputStreamPrivate       struct{}
	GWin32OutputStreamPrivate      struct{}
	GZlibCompressor                struct{}
	GZlibDecompressor              struct{}
	PangoAttrIterator              struct{}
	PangoAttrList                  struct{}
	PangoCairoFont                 struct{}
	PangoCairoFontMap              struct{}
	PangoContext                   struct{}
	PangoCoverage                  struct{}
	PangoEngineLang                struct{}
	PangoEngineShape               struct{}
	PangoFcFontKey                 struct{}
	PangoFcFontMapPrivate          struct{}
	PangoFcFontsetKey              struct{}
	PangoFont                      struct{}
	PangoFontDescription           struct{}
	PangoFontFace                  struct{}
	PangoFontFamily                struct{}
	PangoFontMap                   struct{}
	PangoFontMetrics               struct{}
	PangoFontset                   struct{}
	PangoFontsetSimple             struct{}
	PangoFT2FontMap                struct{}
	PangoLanguage                  struct{}
	PangoLayout                    struct{}
	PangoLayoutIter                struct{}
	PangoMap                       struct{}
	PangoOTBuffer                  struct{}
	PangoOTInfo                    struct{}
	PangoOTRuleset                 struct{}
	PangoRendererPrivate           struct{}
	PangoScriptIter                struct{}
	PangoTabArray                  struct{}
	PangoWin32FontCache            struct{}
	Utimbuf                        struct{}
	Void                           struct{}
)

type GBookmarkFileError Enum

const (
	G_BOOKMARK_FILE_ERROR_INVALID_URI GBookmarkFileError = iota
	G_BOOKMARK_FILE_ERROR_INVALID_VALUE
	G_BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED
	G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND
	G_BOOKMARK_FILE_ERROR_READ
	G_BOOKMARK_FILE_ERROR_UNKNOWN_ENCODING
	G_BOOKMARK_FILE_ERROR_WRITE
	G_BOOKMARK_FILE_ERROR_FILE_NOT_FOUND
)

type GSliceConfig Enum

const (
	G_SLICE_CONFIG_ALWAYS_MALLOC GSliceConfig = iota + 1
	G_SLICE_CONFIG_BYPASS_MAGAZINES
	G_SLICE_CONFIG_WORKING_SET_MSECS
	G_SLICE_CONFIG_COLOR_INCREMENT
	G_SLICE_CONFIG_CHUNK_SIZES
	G_SLICE_CONFIG_CONTENTION_COUNTER
)

type GUserDirectory Enum

const (
	G_USER_DIRECTORY_DESKTOP GUserDirectory = iota
	G_USER_DIRECTORY_DOCUMENTS
	G_USER_DIRECTORY_DOWNLOAD
	G_USER_DIRECTORY_MUSIC
	G_USER_DIRECTORY_PICTURES
	G_USER_DIRECTORY_PUBLIC_SHARE
	G_USER_DIRECTORY_TEMPLATES
	G_USER_DIRECTORY_VIDEOS
	G_USER_N_DIRECTORIES
)

type GThreadError Enum

const (
	G_THREAD_ERROR_AGAIN GThreadError = iota
)

type GThreadPriority Enum

const (
	G_THREAD_PRIORITY_LOW GThreadPriority = iota
	G_THREAD_PRIORITY_NORMAL
	G_THREAD_PRIORITY_HIGH
	G_THREAD_PRIORITY_URGENT
)

type GChecksumType Enum

const (
	G_CHECKSUM_MD5 GChecksumType = iota
	G_CHECKSUM_SHA1
	G_CHECKSUM_SHA256
)

type GOnceStatus Enum

const (
	G_ONCE_STATUS_NOTCALLED GOnceStatus = iota
	G_ONCE_STATUS_PROGRESS
	G_ONCE_STATUS_READY
)

type GConvertError Enum

const (
	G_CONVERT_ERROR_NO_CONVERSION GConvertError = iota
	G_CONVERT_ERROR_ILLEGAL_SEQUENCE
	G_CONVERT_ERROR_FAILED
	G_CONVERT_ERROR_PARTIAL_INPUT
	G_CONVERT_ERROR_BAD_URI
	G_CONVERT_ERROR_NOT_ABSOLUTE_PATH
)

type GDateDMY Enum

const (
	G_DATE_DAY GDateDMY = iota
	G_DATE_MONTH
	G_DATE_YEAR
)

type GDateWeekday Enum

const (
	G_DATE_BAD_WEEKDAY GDateWeekday = iota
	G_DATE_MONDAY
	G_DATE_TUESDAY
	G_DATE_WEDNESDAY
	G_DATE_THURSDAY
	G_DATE_FRIDAY
	G_DATE_SATURDAY
	G_DATE_SUNDAY
)

type GDateMonth Enum

const (
	G_DATE_BAD_MONTH GDateMonth = iota
	G_DATE_JANUARY
	G_DATE_FEBRUARY
	G_DATE_MARCH
	G_DATE_APRIL
	G_DATE_MAY
	G_DATE_JUNE
	G_DATE_JULY
	G_DATE_AUGUST
	G_DATE_SEPTEMBER
	G_DATE_OCTOBER
	G_DATE_NOVEMBER
	G_DATE_DECEMBER
)

type GTimeType Enum

const (
	G_TIME_TYPE_STANDARD GTimeType = iota
	G_TIME_TYPE_DAYLIGHT
	G_TIME_TYPE_UNIVERSAL
)

type GFileError Enum

const (
	G_FILE_ERROR_EXIST GFileError = iota
	G_FILE_ERROR_ISDIR
	G_FILE_ERROR_ACCES
	G_FILE_ERROR_NAMETOOLONG
	G_FILE_ERROR_NOENT
	G_FILE_ERROR_NOTDIR
	G_FILE_ERROR_NXIO
	G_FILE_ERROR_NODEV
	G_FILE_ERROR_ROFS
	G_FILE_ERROR_TXTBSY
	G_FILE_ERROR_FAULT
	G_FILE_ERROR_LOOP
	G_FILE_ERROR_NOSPC
	G_FILE_ERROR_NOMEM
	G_FILE_ERROR_MFILE
	G_FILE_ERROR_NFILE
	G_FILE_ERROR_BADF
	G_FILE_ERROR_INVAL
	G_FILE_ERROR_PIPE
	G_FILE_ERROR_AGAIN
	G_FILE_ERROR_INTR
	G_FILE_ERROR_IO
	G_FILE_ERROR_PERM
	G_FILE_ERROR_NOSYS
	G_FILE_ERROR_FAILED
)

type GFileTest Enum

const (
	G_FILE_TEST_IS_REGULAR GFileTest = 1 << iota
	G_FILE_TEST_IS_SYMLINK
	G_FILE_TEST_IS_DIR
	G_FILE_TEST_IS_EXECUTABLE
	G_FILE_TEST_EXISTS
)

type GHookFlagMask Enum

const (
	G_HOOK_FLAG_ACTIVE GHookFlagMask = 1 << iota
	G_HOOK_FLAG_IN_CALL
	G_HOOK_FLAG_MASK GHookFlagMask = 0x0F
)

type GUnicodeType Enum

const (
	G_UNICODE_CONTROL GUnicodeType = iota
	G_UNICODE_FORMAT
	G_UNICODE_UNASSIGNED
	G_UNICODE_PRIVATE_USE
	G_UNICODE_SURROGATE
	G_UNICODE_LOWERCASE_LETTER
	G_UNICODE_MODIFIER_LETTER
	G_UNICODE_OTHER_LETTER
	G_UNICODE_TITLECASE_LETTER
	G_UNICODE_UPPERCASE_LETTER
	G_UNICODE_COMBINING_MARK
	G_UNICODE_ENCLOSING_MARK
	G_UNICODE_NON_SPACING_MARK
	G_UNICODE_DECIMAL_NUMBER
	G_UNICODE_LETTER_NUMBER
	G_UNICODE_OTHER_NUMBER
	G_UNICODE_CONNECT_PUNCTUATION
	G_UNICODE_DASH_PUNCTUATION
	G_UNICODE_CLOSE_PUNCTUATION
	G_UNICODE_FINAL_PUNCTUATION
	G_UNICODE_INITIAL_PUNCTUATION
	G_UNICODE_OTHER_PUNCTUATION
	G_UNICODE_OPEN_PUNCTUATION
	G_UNICODE_CURRENCY_SYMBOL
	G_UNICODE_MODIFIER_SYMBOL
	G_UNICODE_MATH_SYMBOL
	G_UNICODE_OTHER_SYMBOL
	G_UNICODE_LINE_SEPARATOR
	G_UNICODE_PARAGRAPH_SEPARATOR
	G_UNICODE_SPACE_SEPARATOR
)

type GUnicodeBreakType Enum

const (
	G_UNICODE_BREAK_MANDATORY GUnicodeBreakType = iota
	G_UNICODE_BREAK_CARRIAGE_RETURN
	G_UNICODE_BREAK_LINE_FEED
	G_UNICODE_BREAK_COMBINING_MARK
	G_UNICODE_BREAK_SURROGATE
	G_UNICODE_BREAK_ZERO_WIDTH_SPACE
	G_UNICODE_BREAK_INSEPARABLE
	G_UNICODE_BREAK_NON_BREAKING_GLUE
	G_UNICODE_BREAK_CONTINGENT
	G_UNICODE_BREAK_SPACE
	G_UNICODE_BREAK_AFTER
	G_UNICODE_BREAK_BEFORE
	G_UNICODE_BREAK_BEFORE_AND_AFTER
	G_UNICODE_BREAK_HYPHEN
	G_UNICODE_BREAK_NON_STARTER
	G_UNICODE_BREAK_OPEN_PUNCTUATION
	G_UNICODE_BREAK_CLOSE_PUNCTUATION
	G_UNICODE_BREAK_QUOTATION
	G_UNICODE_BREAK_EXCLAMATION
	G_UNICODE_BREAK_IDEOGRAPHIC
	G_UNICODE_BREAK_NUMERIC
	G_UNICODE_BREAK_INFIX_SEPARATOR
	G_UNICODE_BREAK_SYMBOL
	G_UNICODE_BREAK_ALPHABETIC
	G_UNICODE_BREAK_PREFIX
	G_UNICODE_BREAK_POSTFIX
	G_UNICODE_BREAK_COMPLEX_CONTEXT
	G_UNICODE_BREAK_AMBIGUOUS
	G_UNICODE_BREAK_UNKNOWN
	G_UNICODE_BREAK_NEXT_LINE
	G_UNICODE_BREAK_WORD_JOINER
	G_UNICODE_BREAK_HANGUL_L_JAMO
	G_UNICODE_BREAK_HANGUL_V_JAMO
	G_UNICODE_BREAK_HANGUL_T_JAMO
	G_UNICODE_BREAK_HANGUL_LV_SYLLABLE
	G_UNICODE_BREAK_HANGUL_LVT_SYLLABLE
	G_UNICODE_BREAK_CLOSE_PARANTHESIS
)

type GUnicodeScript Enum

const (
	G_UNICODE_SCRIPT_INVALID_CODE GUnicodeScript = iota - 1
	G_UNICODE_SCRIPT_COMMON
	G_UNICODE_SCRIPT_INHERITED
	G_UNICODE_SCRIPT_ARABIC
	G_UNICODE_SCRIPT_ARMENIAN
	G_UNICODE_SCRIPT_BENGALI
	G_UNICODE_SCRIPT_BOPOMOFO
	G_UNICODE_SCRIPT_CHEROKEE
	G_UNICODE_SCRIPT_COPTIC
	G_UNICODE_SCRIPT_CYRILLIC
	G_UNICODE_SCRIPT_DESERET
	G_UNICODE_SCRIPT_DEVANAGARI
	G_UNICODE_SCRIPT_ETHIOPIC
	G_UNICODE_SCRIPT_GEORGIAN
	G_UNICODE_SCRIPT_GOTHIC
	G_UNICODE_SCRIPT_GREEK
	G_UNICODE_SCRIPT_GUJARATI
	G_UNICODE_SCRIPT_GURMUKHI
	G_UNICODE_SCRIPT_HAN
	G_UNICODE_SCRIPT_HANGUL
	G_UNICODE_SCRIPT_HEBREW
	G_UNICODE_SCRIPT_HIRAGANA
	G_UNICODE_SCRIPT_KANNADA
	G_UNICODE_SCRIPT_KATAKANA
	G_UNICODE_SCRIPT_KHMER
	G_UNICODE_SCRIPT_LAO
	G_UNICODE_SCRIPT_LATIN
	G_UNICODE_SCRIPT_MALAYALAM
	G_UNICODE_SCRIPT_MONGOLIAN
	G_UNICODE_SCRIPT_MYANMAR
	G_UNICODE_SCRIPT_OGHAM
	G_UNICODE_SCRIPT_OLD_ITALIC
	G_UNICODE_SCRIPT_ORIYA
	G_UNICODE_SCRIPT_RUNIC
	G_UNICODE_SCRIPT_SINHALA
	G_UNICODE_SCRIPT_SYRIAC
	G_UNICODE_SCRIPT_TAMIL
	G_UNICODE_SCRIPT_TELUGU
	G_UNICODE_SCRIPT_THAANA
	G_UNICODE_SCRIPT_THAI
	G_UNICODE_SCRIPT_TIBETAN
	G_UNICODE_SCRIPT_CANADIAN_ABORIGINAL
	G_UNICODE_SCRIPT_YI
	G_UNICODE_SCRIPT_TAGALOG
	G_UNICODE_SCRIPT_HANUNOO
	G_UNICODE_SCRIPT_BUHID
	G_UNICODE_SCRIPT_TAGBANWA
	G_UNICODE_SCRIPT_BRAILLE
	G_UNICODE_SCRIPT_CYPRIOT
	G_UNICODE_SCRIPT_LIMBU
	G_UNICODE_SCRIPT_OSMANYA
	G_UNICODE_SCRIPT_SHAVIAN
	G_UNICODE_SCRIPT_LINEAR_B
	G_UNICODE_SCRIPT_TAI_LE
	G_UNICODE_SCRIPT_UGARITIC
	G_UNICODE_SCRIPT_NEW_TAI_LUE
	G_UNICODE_SCRIPT_BUGINESE
	G_UNICODE_SCRIPT_GLAGOLITIC
	G_UNICODE_SCRIPT_TIFINAGH
	G_UNICODE_SCRIPT_SYLOTI_NAGRI
	G_UNICODE_SCRIPT_OLD_PERSIAN
	G_UNICODE_SCRIPT_KHAROSHTHI
	G_UNICODE_SCRIPT_UNKNOWN
	G_UNICODE_SCRIPT_BALINESE
	G_UNICODE_SCRIPT_CUNEIFORM
	G_UNICODE_SCRIPT_PHOENICIAN
	G_UNICODE_SCRIPT_PHAGS_PA
	G_UNICODE_SCRIPT_NKO
	G_UNICODE_SCRIPT_KAYAH_LI
	G_UNICODE_SCRIPT_LEPCHA
	G_UNICODE_SCRIPT_REJANG
	G_UNICODE_SCRIPT_SUNDANESE
	G_UNICODE_SCRIPT_SAURASHTRA
	G_UNICODE_SCRIPT_CHAM
	G_UNICODE_SCRIPT_OL_CHIKI
	G_UNICODE_SCRIPT_VAI
	G_UNICODE_SCRIPT_CARIAN
	G_UNICODE_SCRIPT_LYCIAN
	G_UNICODE_SCRIPT_LYDIAN
	G_UNICODE_SCRIPT_AVESTAN
	G_UNICODE_SCRIPT_BAMUM
	G_UNICODE_SCRIPT_EGYPTIAN_HIEROGLYPHS
	G_UNICODE_SCRIPT_IMPERIAL_ARAMAIC
	G_UNICODE_SCRIPT_INSCRIPTIONAL_PAHLAVI
	G_UNICODE_SCRIPT_INSCRIPTIONAL_PARTHIAN
	G_UNICODE_SCRIPT_JAVANESE
	G_UNICODE_SCRIPT_KAITHI
	G_UNICODE_SCRIPT_LISU
	G_UNICODE_SCRIPT_MEETEI_MAYEK
	G_UNICODE_SCRIPT_OLD_SOUTH_ARABIAN
	G_UNICODE_SCRIPT_OLD_TURKIC
	G_UNICODE_SCRIPT_SAMARITAN
	G_UNICODE_SCRIPT_TAI_THAM
	G_UNICODE_SCRIPT_TAI_VIET
	G_UNICODE_SCRIPT_BATAK
	G_UNICODE_SCRIPT_BRAHMI
	G_UNICODE_SCRIPT_MANDAIC
)

type GNormalizeMode Enum

const (
	G_NORMALIZE_DEFAULT GNormalizeMode = iota
	G_NORMALIZE_DEFAULT_COMPOSE
	G_NORMALIZE_ALL
	G_NORMALIZE_ALL_COMPOSE
	G_NORMALIZE_NFD  = G_NORMALIZE_DEFAULT
	G_NORMALIZE_NFC  = G_NORMALIZE_DEFAULT_COMPOSE
	G_NORMALIZE_NFKD = G_NORMALIZE_ALL
	G_NORMALIZE_NFKC = G_NORMALIZE_ALL_COMPOSE
)

type GIOError Enum

const (
	G_IO_ERROR_NONE GIOError = iota
	G_IO_ERROR_AGAIN
	G_IO_ERROR_INVAL
	G_IO_ERROR_UNKNOWN
)

type GIOChannelError Enum

const (
	G_IO_CHANNEL_ERROR_FBIG GIOChannelError = iota
	G_IO_CHANNEL_ERROR_INVAL
	G_IO_CHANNEL_ERROR_IO
	G_IO_CHANNEL_ERROR_ISDIR
	G_IO_CHANNEL_ERROR_NOSPC
	G_IO_CHANNEL_ERROR_NXIO
	G_IO_CHANNEL_ERROR_OVERFLOW
	G_IO_CHANNEL_ERROR_PIPE
	G_IO_CHANNEL_ERROR_FAILED
)

type GIOStatus Enum

const (
	G_IO_STATUS_ERROR GIOStatus = iota
	G_IO_STATUS_NORMAL
	G_IO_STATUS_EOF
	G_IO_STATUS_AGAIN
)

type GSeekType Enum

const (
	G_SEEK_CUR GSeekType = iota
	G_SEEK_SET
	G_SEEK_END
)

type GIOCondition Enum

const (
	G_IO_IN GIOCondition = 1 << iota
	G_IO_PRI
	G_IO_OUT
	G_IO_ERR
	G_IO_HUP
	G_IO_NVAL
)

type GIOFlags Enum

const (
	G_IO_FLAG_APPEND GIOFlags = 1 << iota
	G_IO_FLAG_NONBLOCK
	G_IO_FLAG_IS_READABLE
	G_IO_FLAG_IS_WRITEABLE
	G_IO_FLAG_IS_SEEKABLE
	G_IO_FLAG_MASK     GIOFlags = (1 << iota) - 1
	G_IO_FLAG_GET_MASK          = G_IO_FLAG_MASK
	G_IO_FLAG_SET_MASK          = G_IO_FLAG_APPEND | G_IO_FLAG_NONBLOCK
)

type GKeyFileError Enum

const (
	G_KEY_FILE_ERROR_UNKNOWN_ENCODING GKeyFileError = iota
	G_KEY_FILE_ERROR_PARSE
	G_KEY_FILE_ERROR_NOT_FOUND
	G_KEY_FILE_ERROR_KEY_NOT_FOUND
	G_KEY_FILE_ERROR_GROUP_NOT_FOUND
	G_KEY_FILE_ERROR_INVALID_VALUE
)

type GKeyFileFlags Enum

const (
	G_KEY_FILE_KEEP_COMMENTS GKeyFileFlags = 1 << iota
	G_KEY_FILE_KEEP_TRANSLATIONS
	G_KEY_FILE_NONE GKeyFileFlags = 0
)

type GMarkupError Enum

const (
	G_MARKUP_ERROR_BAD_UTF8 GMarkupError = iota
	G_MARKUP_ERROR_EMPTY
	G_MARKUP_ERROR_PARSE
	G_MARKUP_ERROR_UNKNOWN_ELEMENT
	G_MARKUP_ERROR_UNKNOWN_ATTRIBUTE
	G_MARKUP_ERROR_INVALID_CONTENT
	G_MARKUP_ERROR_MISSING_ATTRIBUTE
)

type GMarkupParseFlags Enum

const (
	G_MARKUP_DO_NOT_USE_THIS_UNSUPPORTED_FLAG GMarkupParseFlags = 1 << iota
	G_MARKUP_TREAT_CDATA_AS_TEXT
	G_MARKUP_PREFIX_ERROR_POSITION
)

type GMarkupCollectType Enum

const (
	G_MARKUP_COLLECT_INVALID GMarkupCollectType = iota
	G_MARKUP_COLLECT_STRING
	G_MARKUP_COLLECT_STRDUP
	G_MARKUP_COLLECT_BOOLEAN
	G_MARKUP_COLLECT_TRISTATE
	G_MARKUP_COLLECT_OPTIONAL GMarkupCollectType = 1 << 16
)

type GLogLevelFlags Enum

const (
	G_LOG_FLAG_RECURSION GLogLevelFlags = 1 << iota
	G_LOG_FLAG_FATAL
	G_LOG_LEVEL_ERROR
	G_LOG_LEVEL_CRITICAL
	G_LOG_LEVEL_WARNING
	G_LOG_LEVEL_MESSAGE
	G_LOG_LEVEL_INFO
	G_LOG_LEVEL_DEBUG
	G_LOG_LEVEL_MASK = ^(G_LOG_FLAG_RECURSION | G_LOG_FLAG_FATAL)
)

type GTraverseFlags Enum

const (
	G_TRAVERSE_LEAVES GTraverseFlags = 1 << iota
	G_TRAVERSE_NON_LEAVES
	G_TRAVERSE_ALL       = G_TRAVERSE_LEAVES | G_TRAVERSE_NON_LEAVES
	G_TRAVERSE_MASK      = G_TRAVERSE_ALL // 0x03
	G_TRAVERSE_LEAFS     = G_TRAVERSE_LEAVES
	G_TRAVERSE_NON_LEAFS = G_TRAVERSE_NON_LEAVES
)

type GTraverseType Enum

const (
	G_IN_ORDER GTraverseType = iota
	G_PRE_ORDER
	G_POST_ORDER
	G_LEVEL_ORDER
)

type GOptionFlags Enum

const (
	G_OPTION_FLAG_HIDDEN GOptionFlags = 1 << iota
	G_OPTION_FLAG_IN_MAIN
	G_OPTION_FLAG_REVERSE
	G_OPTION_FLAG_NO_ARG
	G_OPTION_FLAG_FILENAME
	G_OPTION_FLAG_OPTIONAL_ARG
	G_OPTION_FLAG_NOALIAS
)

type GOptionArg Enum

const (
	G_OPTION_ARG_NONE GOptionArg = iota
	G_OPTION_ARG_STRING
	G_OPTION_ARG_INT
	G_OPTION_ARG_CALLBACK
	G_OPTION_ARG_FILENAME
	G_OPTION_ARG_STRING_ARRAY
	G_OPTION_ARG_FILENAME_ARRAY
	G_OPTION_ARG_DOUBLE
	G_OPTION_ARG_INT64
)

type GOptionError Enum

const (
	G_OPTION_ERROR_UNKNOWN_OPTION GOptionError = iota
	G_OPTION_ERROR_BAD_VALUE
	G_OPTION_ERROR_FAILED
)

type GRegexError Enum

const (
	G_REGEX_ERROR_COMPILE GRegexError = iota
	G_REGEX_ERROR_OPTIMIZE
	G_REGEX_ERROR_REPLACE
	G_REGEX_ERROR_MATCH
	G_REGEX_ERROR_INTERNAL
	//TODO(t): fix
)
const (
	G_REGEX_ERROR_STRAY_BACKSLASH GRegexError = 101 + iota
	G_REGEX_ERROR_MISSING_CONTROL_CHAR
	G_REGEX_ERROR_UNRECOGNIZED_ESCAPE
	G_REGEX_ERROR_QUANTIFIERS_OUT_OF_ORDER
	G_REGEX_ERROR_QUANTIFIER_TOO_BIG
	G_REGEX_ERROR_UNTERMINATED_CHARACTER_CLASS
	G_REGEX_ERROR_INVALID_ESCAPE_IN_CHARACTER_CLASS
	G_REGEX_ERROR_RANGE_OUT_OF_ORDER
	G_REGEX_ERROR_NOTHING_TO_REPEAT
	_ // 110
	_
	G_REGEX_ERROR_UNRECOGNIZED_CHARACTER
	G_REGEX_ERROR_POSIX_NAMED_CLASS_OUTSIDE_CLASS
	G_REGEX_ERROR_UNMATCHED_PARENTHESIS
	G_REGEX_ERROR_INEXISTENT_SUBPATTERN_REFERENCE
	_
	_
	G_REGEX_ERROR_UNTERMINATED_COMMENT
	_
	G_REGEX_ERROR_EXPRESSION_TOO_LARGE // 120
	G_REGEX_ERROR_MEMORY_ERROR
	_
	_
	_
	G_REGEX_ERROR_VARIABLE_LENGTH_LOOKBEHIND
	G_REGEX_ERROR_MALFORMED_CONDITION
	G_REGEX_ERROR_TOO_MANY_CONDITIONAL_BRANCHES
	G_REGEX_ERROR_ASSERTION_EXPECTED
	_
	G_REGEX_ERROR_UNKNOWN_POSIX_CLASS_NAME // 130
	G_REGEX_ERROR_POSIX_COLLATING_ELEMENTS_NOT_SUPPORTED
	_
	_
	G_REGEX_ERROR_HEX_CODE_TOO_LARGE
	G_REGEX_ERROR_INVALID_CONDITION
	G_REGEX_ERROR_SINGLE_BYTE_MATCH_IN_LOOKBEHIND
	_
	_
	_
	G_REGEX_ERROR_INFINITE_LOOP // 140
	_
	G_REGEX_ERROR_MISSING_SUBPATTERN_NAME_TERMINATOR
	G_REGEX_ERROR_DUPLICATE_SUBPATTERN_NAME
	_
	_
	G_REGEX_ERROR_MALFORMED_PROPERTY
	G_REGEX_ERROR_UNKNOWN_PROPERTY
	G_REGEX_ERROR_SUBPATTERN_NAME_TOO_LONG
	G_REGEX_ERROR_TOO_MANY_SUBPATTERNS
	_ // 150
	G_REGEX_ERROR_INVALID_OCTAL_VALUE
	_
	_
	G_REGEX_ERROR_TOO_MANY_BRANCHES_IN_DEFINE
	G_REGEX_ERROR_DEFINE_REPETION
	G_REGEX_ERROR_INCONSISTENT_NEWLINE_OPTIONS
	G_REGEX_ERROR_MISSING_BACK_REFERENCE //157
)

type GRegexCompileFlags Enum

const (
	G_REGEX_CASELESS GRegexCompileFlags = 1 << iota
	G_REGEX_MULTILINE
	G_REGEX_DOTALL
	G_REGEX_EXTENDED
	G_REGEX_ANCHORED
	G_REGEX_DOLLAR_ENDONLY
	_
	_
	_
	G_REGEX_UNGREEDY
	_
	G_REGEX_RAW
	G_REGEX_NO_AUTO_CAPTURE
	G_REGEX_OPTIMIZE
	_
	_
	_
	_
	_

	G_REGEX_DUPNAMES
	G_REGEX_NEWLINE_CR
	G_REGEX_NEWLINE_LF
	G_REGEX_NEWLINE_CRLF = G_REGEX_NEWLINE_CR | G_REGEX_NEWLINE_LF
)

type GRegexMatchFlags Enum

const (
	G_REGEX_MATCH_ANCHORED GRegexMatchFlags = 1 << (iota + 4)
	_
	_
	G_REGEX_MATCH_NOTBOL
	G_REGEX_MATCH_NOTEOL
	_
	G_REGEX_MATCH_NOTEMPTY
	_
	_
	_
	_
	G_REGEX_MATCH_PARTIAL
	_
	_
	_
	_
	G_REGEX_MATCH_NEWLINE_CR
	G_REGEX_MATCH_NEWLINE_LF
	G_REGEX_MATCH_NEWLINE_ANY
	G_REGEX_MATCH_NEWLINE_CRLF = G_REGEX_MATCH_NEWLINE_CR | G_REGEX_MATCH_NEWLINE_LF
)

type GErrorType Enum

const (
	G_ERR_UNKNOWN GErrorType = iota
	G_ERR_UNEXP_EOF
	G_ERR_UNEXP_EOF_IN_STRING
	G_ERR_UNEXP_EOF_IN_COMMENT
	G_ERR_NON_DIGIT_IN_CONST
	G_ERR_DIGIT_RADIX
	G_ERR_FLOAT_RADIX
	G_ERR_FLOAT_MALFORMED
)

type GTokenType Enum

const (
	G_TOKEN_EOF         GTokenType = 0
	G_TOKEN_LEFT_PAREN  GTokenType = '('
	G_TOKEN_RIGHT_PAREN GTokenType = ')'
	G_TOKEN_LEFT_CURLY  GTokenType = '{'
	G_TOKEN_RIGHT_CURLY GTokenType = '}'
	G_TOKEN_LEFT_BRACE  GTokenType = '['
	G_TOKEN_RIGHT_BRACE GTokenType = ']'
	G_TOKEN_EQUAL_SIGN  GTokenType = '='
	G_TOKEN_COMMA       GTokenType = ','
)
const (
	G_TOKEN_NONE GTokenType = 0x100 + iota
	G_TOKEN_ERROR
	G_TOKEN_CHAR
	G_TOKEN_BINARY
	G_TOKEN_OCTAL
	G_TOKEN_INT
	G_TOKEN_HEX
	G_TOKEN_FLOAT
	G_TOKEN_STRING
	G_TOKEN_SYMBOL
	G_TOKEN_IDENTIFIER
	G_TOKEN_IDENTIFIER_NULL
	G_TOKEN_COMMENT_SINGLE
	G_TOKEN_COMMENT_MULTI
	G_TOKEN_LAST
)

type GShellError Enum

const (
	G_SHELL_ERROR_BAD_QUOTING GShellError = iota
	G_SHELL_ERROR_EMPTY_STRING
	G_SHELL_ERROR_FAILED
)

type GSpawnError Enum

const (
	G_SPAWN_ERROR_FORK GSpawnError = iota
	G_SPAWN_ERROR_READ
	G_SPAWN_ERROR_CHDIR
	G_SPAWN_ERROR_ACCES
	G_SPAWN_ERROR_PERM
	G_SPAWN_ERROR_2BIG
	G_SPAWN_ERROR_NOEXEC
	G_SPAWN_ERROR_NAMETOOLONG
	G_SPAWN_ERROR_NOENT
	G_SPAWN_ERROR_NOMEM
	G_SPAWN_ERROR_NOTDIR
	G_SPAWN_ERROR_LOOP
	G_SPAWN_ERROR_TXTBUSY
	G_SPAWN_ERROR_IO
	G_SPAWN_ERROR_NFILE
	G_SPAWN_ERROR_MFILE
	G_SPAWN_ERROR_INVAL
	G_SPAWN_ERROR_ISDIR
	G_SPAWN_ERROR_LIBBAD
	G_SPAWN_ERROR_FAILED
)

type GSpawnFlags Enum

const (
	G_SPAWN_LEAVE_DESCRIPTORS_OPEN GSpawnFlags = 1 << iota
	G_SPAWN_DO_NOT_REAP_CHILD
	G_SPAWN_SEARCH_PATH
	G_SPAWN_STDOUT_TO_DEV_NULL
	G_SPAWN_STDERR_TO_DEV_NULL
	G_SPAWN_CHILD_INHERITS_STDIN
	G_SPAWN_FILE_AND_ARGV_ZERO
)

type GAsciiType Enum

const (
	G_ASCII_ALNUM GAsciiType = 1 << iota
	G_ASCII_ALPHA
	G_ASCII_CNTRL
	G_ASCII_DIGIT
	G_ASCII_GRAPH
	G_ASCII_LOWER
	G_ASCII_PRINT
	G_ASCII_PUNCT
	G_ASCII_SPACE
	G_ASCII_UPPER
	G_ASCII_XDIGIT
)

type GTestTrapFlags Enum

const (
	G_TEST_TRAP_SILENCE_STDOUT GTestTrapFlags = 1 << (7 + iota)
	G_TEST_TRAP_SILENCE_STDERR
	G_TEST_TRAP_INHERIT_STDIN
)

type GTestLogType Enum

const (
	G_TEST_LOG_NONE GTestLogType = iota
	G_TEST_LOG_ERROR
	G_TEST_LOG_START_BINARY
	G_TEST_LOG_LIST_CASE
	G_TEST_LOG_SKIP_CASE
	G_TEST_LOG_START_CASE
	G_TEST_LOG_STOP_CASE
	G_TEST_LOG_MIN_RESULT
	G_TEST_LOG_MAX_RESULT
	G_TEST_LOG_MESSAGE
)

type GVariantClass Enum

const (
	G_VARIANT_CLASS_BOOLEAN     GVariantClass = 'b'
	G_VARIANT_CLASS_BYTE        GVariantClass = 'y'
	G_VARIANT_CLASS_INT16       GVariantClass = 'n'
	G_VARIANT_CLASS_UINT16      GVariantClass = 'q'
	G_VARIANT_CLASS_INT32       GVariantClass = 'i'
	G_VARIANT_CLASS_UINT32      GVariantClass = 'u'
	G_VARIANT_CLASS_INT64       GVariantClass = 'x'
	G_VARIANT_CLASS_UINT64      GVariantClass = 't'
	G_VARIANT_CLASS_HANDLE      GVariantClass = 'h'
	G_VARIANT_CLASS_DOUBLE      GVariantClass = 'd'
	G_VARIANT_CLASS_STRING      GVariantClass = 's'
	G_VARIANT_CLASS_OBJECT_PATH GVariantClass = 'o'
	G_VARIANT_CLASS_SIGNATURE   GVariantClass = 'g'
	G_VARIANT_CLASS_VARIANT     GVariantClass = 'v'
	G_VARIANT_CLASS_MAYBE       GVariantClass = 'm'
	G_VARIANT_CLASS_ARRAY       GVariantClass = 'a'
	G_VARIANT_CLASS_TUPLE       GVariantClass = '('
	G_VARIANT_CLASS_DICT_ENTRY  GVariantClass = '{'
)

type GVariantParseError Enum

const (
	G_VARIANT_PARSE_ERROR_FAILED GVariantParseError = iota
	G_VARIANT_PARSE_ERROR_BASIC_TYPE_EXPECTED
	G_VARIANT_PARSE_ERROR_CANNOT_INFER_TYPE
	G_VARIANT_PARSE_ERROR_DEFINITE_TYPE_EXPECTED
	G_VARIANT_PARSE_ERROR_INPUT_NOT_AT_END
	G_VARIANT_PARSE_ERROR_INVALID_CHARACTER
	G_VARIANT_PARSE_ERROR_INVALID_FORMAT_STRING
	G_VARIANT_PARSE_ERROR_INVALID_OBJECT_PATH
	G_VARIANT_PARSE_ERROR_INVALID_SIGNATURE
	G_VARIANT_PARSE_ERROR_INVALID_TYPE_STRING
	G_VARIANT_PARSE_ERROR_NO_COMMON_TYPE
	G_VARIANT_PARSE_ERROR_NUMBER_OUT_OF_RANGE
	G_VARIANT_PARSE_ERROR_NUMBER_TOO_BIG
	G_VARIANT_PARSE_ERROR_TYPE_ERROR
	G_VARIANT_PARSE_ERROR_UNEXPECTED_TOKEN
	G_VARIANT_PARSE_ERROR_UNKNOWN_KEYWORD
	G_VARIANT_PARSE_ERROR_UNTERMINATED_STRING_CONSTANT
	G_VARIANT_PARSE_ERROR_VALUE_EXPECTED
)

type GTypeFundamentalFlags Enum

const (
	G_TYPE_FLAG_CLASSED GTypeFundamentalFlags = 1 << iota
	G_TYPE_FLAG_INSTANTIATABLE
	G_TYPE_FLAG_DERIVABLE
	G_TYPE_FLAG_DEEP_DERIVABLE
)

type GSignalFlags Enum

const (
	G_SIGNAL_RUN_FIRST GSignalFlags = 1 << iota
	G_SIGNAL_RUN_LAST
	G_SIGNAL_RUN_CLEANUP
	G_SIGNAL_NO_RECURSE
	G_SIGNAL_DETAILED
	G_SIGNAL_ACTION
	G_SIGNAL_NO_HOOKS
)

type GConnectFlags Enum

const (
	G_CONNECT_AFTER GConnectFlags = 1 << iota
	G_CONNECT_SWAPPED
)

type GSignalMatchType Enum

const (
	G_SIGNAL_MATCH_ID GSignalMatchType = 1 << iota
	G_SIGNAL_MATCH_DETAIL
	G_SIGNAL_MATCH_CLOSURE
	G_SIGNAL_MATCH_FUNC
	G_SIGNAL_MATCH_DATA
	G_SIGNAL_MATCH_UNBLOCKED
)

type GAppInfoCreateFlags Enum

const (
	G_APP_INFO_CREATE_NEEDS_TERMINAL GAppInfoCreateFlags = 1 << iota
	G_APP_INFO_CREATE_SUPPORTS_URIS
	G_APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION
	G_APP_INFO_CREATE_NONE GAppInfoCreateFlags = 0
)

type GConverterFlags Enum

const (
	G_CONVERTER_INPUT_AT_END GConverterFlags = 1 << iota
	G_CONVERTER_FLUSH
	G_CONVERTER_NO_FLAGS GConverterFlags = 0
)

type GConverterResult Enum

const (
	G_CONVERTER_ERROR GConverterResult = iota
	G_CONVERTER_CONVERTED
	G_CONVERTER_FINISHED
	G_CONVERTER_FLUSHED
)

type GDataStreamByteOrder Enum

const (
	G_DATA_STREAM_BYTE_ORDER_BIG_ENDIAN GDataStreamByteOrder = iota
	G_DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN
	G_DATA_STREAM_BYTE_ORDER_HOST_ENDIAN
)

type GDataStreamNewlineType Enum

const (
	G_DATA_STREAM_NEWLINE_TYPE_LF GDataStreamNewlineType = iota
	G_DATA_STREAM_NEWLINE_TYPE_CR
	G_DATA_STREAM_NEWLINE_TYPE_CR_LF
	G_DATA_STREAM_NEWLINE_TYPE_ANY
)

type GFileAttributeType Enum

const (
	G_FILE_ATTRIBUTE_TYPE_INVALID GFileAttributeType = iota
	G_FILE_ATTRIBUTE_TYPE_STRING
	G_FILE_ATTRIBUTE_TYPE_BYTE_STRING
	G_FILE_ATTRIBUTE_TYPE_BOOLEAN
	G_FILE_ATTRIBUTE_TYPE_UINT32
	G_FILE_ATTRIBUTE_TYPE_INT32
	G_FILE_ATTRIBUTE_TYPE_UINT64
	G_FILE_ATTRIBUTE_TYPE_INT64
	G_FILE_ATTRIBUTE_TYPE_OBJECT
	G_FILE_ATTRIBUTE_TYPE_STRINGV
)

type GFileAttributeInfoFlags Enum

const (
	G_FILE_ATTRIBUTE_INFO_COPY_WITH_FILE GFileAttributeInfoFlags = 1 << iota
	G_FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED
	G_FILE_ATTRIBUTE_INFO_NONE GFileAttributeInfoFlags = 0
)

type GFileAttributeStatus Enum

const (
	G_FILE_ATTRIBUTE_STATUS_UNSET GFileAttributeStatus = iota
	G_FILE_ATTRIBUTE_STATUS_SET
	G_FILE_ATTRIBUTE_STATUS_ERROR_SETTING
)

type GFileQueryInfoFlags Enum

const (
	G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS GFileQueryInfoFlags = 1 << iota
	G_FILE_QUERY_INFO_NONE              GFileQueryInfoFlags = 0
)

type GFileCreateFlags Enum

const (
	G_FILE_CREATE_PRIVATE GFileCreateFlags = 1 << iota
	G_FILE_CREATE_REPLACE_DESTINATION
	G_FILE_CREATE_NONE GFileCreateFlags = 0
)

type GMountMountFlags Enum

const G_MOUNT_MOUNT_NONE GMountMountFlags = 0

type GMountUnmountFlags Enum

const (
	G_MOUNT_UNMOUNT_FORCE GMountUnmountFlags = 1 << iota
	G_MOUNT_UNMOUNT_NONE  GMountUnmountFlags = 0
)

type GDriveStartFlags Enum

const G_DRIVE_START_NONE GDriveStartFlags = 0

type GDriveStartStopType Enum

const (
	G_DRIVE_START_STOP_TYPE_UNKNOWN GDriveStartStopType = iota
	G_DRIVE_START_STOP_TYPE_SHUTDOWN
	G_DRIVE_START_STOP_TYPE_NETWORK
	G_DRIVE_START_STOP_TYPE_MULTIDISK
	G_DRIVE_START_STOP_TYPE_PASSWORD
)

type GFileCopyFlags Enum

const (
	G_FILE_COPY_OVERWRITE GFileCopyFlags = 1 << iota
	G_FILE_COPY_BACKUP
	G_FILE_COPY_NOFOLLOW_SYMLINKS
	G_FILE_COPY_ALL_METADATA
	G_FILE_COPY_NO_FALLBACK_FOR_MOVE
	G_FILE_COPY_TARGET_DEFAULT_PERMS
	G_FILE_COPY_NONE GFileCopyFlags = 0
)

type GFileMonitorFlags Enum

const (
	G_FILE_MONITOR_WATCH_MOUNTS GFileMonitorFlags = 1 << iota
	G_FILE_MONITOR_SEND_MOVED
	G_FILE_MONITOR_NONE GFileMonitorFlags = 0
)

type GFileType Enum

const (
	G_FILE_TYPE_UNKNOWN GFileType = iota
	G_FILE_TYPE_REGULAR
	G_FILE_TYPE_DIRECTORY
	G_FILE_TYPE_SYMBOLIC_LINK
	G_FILE_TYPE_SPECIAL
	G_FILE_TYPE_SHORTCUT
	G_FILE_TYPE_MOUNTABLE
)

type GFilesystemPreviewType Enum

const (
	G_FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS GFilesystemPreviewType = iota
	G_FILESYSTEM_PREVIEW_TYPE_IF_LOCAL
	G_FILESYSTEM_PREVIEW_TYPE_NEVER
)

type GFileMonitorEvent Enum

const (
	G_FILE_MONITOR_EVENT_CHANGED GFileMonitorEvent = iota
	G_FILE_MONITOR_EVENT_CHANGES_DONE_HINT
	G_FILE_MONITOR_EVENT_DELETED
	G_FILE_MONITOR_EVENT_CREATED
	G_FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED
	G_FILE_MONITOR_EVENT_PRE_UNMOUNT
	G_FILE_MONITOR_EVENT_UNMOUNTED
	G_FILE_MONITOR_EVENT_MOVED
)

type GIOErrorEnum Enum

const (
	G_IO_ERROR_FAILED GIOErrorEnum = iota
	G_IO_ERROR_NOT_FOUND
	G_IO_ERROR_EXISTS
	G_IO_ERROR_IS_DIRECTORY
	G_IO_ERROR_NOT_DIRECTORY
	G_IO_ERROR_NOT_EMPTY
	G_IO_ERROR_NOT_REGULAR_FILE
	G_IO_ERROR_NOT_SYMBOLIC_LINK
	G_IO_ERROR_NOT_MOUNTABLE_FILE
	G_IO_ERROR_FILENAME_TOO_LONG
	G_IO_ERROR_INVALID_FILENAME
	G_IO_ERROR_TOO_MANY_LINKS
	G_IO_ERROR_NO_SPACE
	G_IO_ERROR_INVALID_ARGUMENT
	G_IO_ERROR_PERMISSION_DENIED
	G_IO_ERROR_NOT_SUPPORTED
	G_IO_ERROR_NOT_MOUNTED
	G_IO_ERROR_ALREADY_MOUNTED
	G_IO_ERROR_CLOSED
	G_IO_ERROR_CANCELLED
	G_IO_ERROR_PENDING
	G_IO_ERROR_READ_ONLY
	G_IO_ERROR_CANT_CREATE_BACKUP
	G_IO_ERROR_WRONG_ETAG
	G_IO_ERROR_TIMED_OUT
	G_IO_ERROR_WOULD_RECURSE
	G_IO_ERROR_BUSY
	G_IO_ERROR_WOULD_BLOCK
	G_IO_ERROR_HOST_NOT_FOUND
	G_IO_ERROR_WOULD_MERGE
	G_IO_ERROR_FAILED_HANDLED
	G_IO_ERROR_TOO_MANY_OPEN_FILES
	G_IO_ERROR_NOT_INITIALIZED
	G_IO_ERROR_ADDRESS_IN_USE
	G_IO_ERROR_PARTIAL_INPUT
	G_IO_ERROR_INVALID_DATA
	G_IO_ERROR_DBUS_ERROR
	G_IO_ERROR_HOST_UNREACHABLE
	G_IO_ERROR_NETWORK_UNREACHABLE
	G_IO_ERROR_CONNECTION_REFUSED
	G_IO_ERROR_PROXY_FAILED
	G_IO_ERROR_PROXY_AUTH_FAILED
	G_IO_ERROR_PROXY_NEED_AUTH
	G_IO_ERROR_PROXY_NOT_ALLOWED
)

type GAskPasswordFlags Enum

const (
	G_ASK_PASSWORD_NEED_PASSWORD GAskPasswordFlags = 1 << iota
	G_ASK_PASSWORD_NEED_USERNAME
	G_ASK_PASSWORD_NEED_DOMAIN
	G_ASK_PASSWORD_SAVING_SUPPORTED
	G_ASK_PASSWORD_ANONYMOUS_SUPPORTED
)

type GPasswordSave Enum

const (
	G_PASSWORD_SAVE_NEVER GPasswordSave = iota
	G_PASSWORD_SAVE_FOR_SESSION
	G_PASSWORD_SAVE_PERMANENTLY
)

type GMountOperationResult Enum

const (
	G_MOUNT_OPERATION_HANDLED GMountOperationResult = iota
	G_MOUNT_OPERATION_ABORTED
	G_MOUNT_OPERATION_UNHANDLED
)

type GOutputStreamSpliceFlags Enum

const (
	G_OUTPUT_STREAM_SPLICE_CLOSE_SOURCE GOutputStreamSpliceFlags = 1 << iota
	G_OUTPUT_STREAM_SPLICE_CLOSE_TARGET
	G_OUTPUT_STREAM_SPLICE_NONE GOutputStreamSpliceFlags = 0
)

type GIOStreamSpliceFlags Enum

const (
	G_IO_STREAM_SPLICE_CLOSE_STREAM1 GIOStreamSpliceFlags = 1 << iota
	G_IO_STREAM_SPLICE_CLOSE_STREAM2
	G_IO_STREAM_SPLICE_WAIT_FOR_BOTH
	G_IO_STREAM_SPLICE_NONE GIOStreamSpliceFlags = 0
)

type GEmblemOrigin Enum

const (
	G_EMBLEM_ORIGIN_UNKNOWN GEmblemOrigin = iota
	G_EMBLEM_ORIGIN_DEVICE
	G_EMBLEM_ORIGIN_LIVEMETADATA
	G_EMBLEM_ORIGIN_TAG
)

type GResolverError Enum

const (
	G_RESOLVER_ERROR_NOT_FOUND GResolverError = iota
	G_RESOLVER_ERROR_TEMPORARY_FAILURE
	G_RESOLVER_ERROR_INTERNAL
)

type GSocketFamily Enum

const (
	G_SOCKET_FAMILY_INVALID GSocketFamily = 0
	G_SOCKET_FAMILY_UNIX    GSocketFamily = 1
	G_SOCKET_FAMILY_IPV4    GSocketFamily = 2
	G_SOCKET_FAMILY_IPV6    GSocketFamily = 23
)

type GSocketType Enum

const (
	G_SOCKET_TYPE_INVALID GSocketType = iota
	G_SOCKET_TYPE_STREAM
	G_SOCKET_TYPE_DATAGRAM
	G_SOCKET_TYPE_SEQPACKET
)

type GSocketMsgFlags Enum

const (
	G_SOCKET_MSG_OOB GSocketMsgFlags = 1 << iota
	G_SOCKET_MSG_PEEK
	G_SOCKET_MSG_DONTROUTE
	G_SOCKET_MSG_NONE GSocketMsgFlags = 0
)

type GSocketProtocol Enum

const (
	G_SOCKET_PROTOCOL_UNKNOWN GSocketProtocol = -1
	G_SOCKET_PROTOCOL_DEFAULT GSocketProtocol = 0
	G_SOCKET_PROTOCOL_TCP     GSocketProtocol = 6
	G_SOCKET_PROTOCOL_UDP     GSocketProtocol = 17
	G_SOCKET_PROTOCOL_SCTP    GSocketProtocol = 132
)

type GZlibCompressorFormat Enum

const (
	G_ZLIB_COMPRESSOR_FORMAT_ZLIB GZlibCompressorFormat = iota
	G_ZLIB_COMPRESSOR_FORMAT_GZIP
	G_ZLIB_COMPRESSOR_FORMAT_RAW
)

type GUnixSocketAddressType Enum

const (
	G_UNIX_SOCKET_ADDRESS_INVALID GUnixSocketAddressType = iota
	G_UNIX_SOCKET_ADDRESS_ANONYMOUS
	G_UNIX_SOCKET_ADDRESS_PATH
	G_UNIX_SOCKET_ADDRESS_ABSTRACT
	G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED
)

type GBusType Enum

const (
	G_BUS_TYPE_STARTER GBusType = iota - 1
	G_BUS_TYPE_NONE
	G_BUS_TYPE_SYSTEM
	G_BUS_TYPE_SESSION
)

type GBusNameOwnerFlags Enum

const (
	G_BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT GBusNameOwnerFlags = 1 << iota
	G_BUS_NAME_OWNER_FLAGS_REPLACE
	G_BUS_NAME_OWNER_FLAGS_NONE GBusNameOwnerFlags = 0
)

type GBusNameWatcherFlags Enum

const (
	G_BUS_NAME_WATCHER_FLAGS_AUTO_START GBusNameWatcherFlags = 1 << iota
	G_BUS_NAME_WATCHER_FLAGS_NONE       GBusNameWatcherFlags = 0
)

type GDBusProxyFlags Enum

const (
	G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES GDBusProxyFlags = 1 << iota
	G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS
	G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START
	G_DBUS_PROXY_FLAGS_NONE GDBusProxyFlags = 0
)

type GDBusError Enum

const (
	G_DBUS_ERROR_FAILED GDBusError = iota
	G_DBUS_ERROR_NO_MEMORY
	G_DBUS_ERROR_SERVICE_UNKNOWN
	G_DBUS_ERROR_NAME_HAS_NO_OWNER
	G_DBUS_ERROR_NO_REPLY
	G_DBUS_ERROR_IO_ERROR
	G_DBUS_ERROR_BAD_ADDRESS
	G_DBUS_ERROR_NOT_SUPPORTED
	G_DBUS_ERROR_LIMITS_EXCEEDED
	G_DBUS_ERROR_ACCESS_DENIED
	G_DBUS_ERROR_AUTH_FAILED
	G_DBUS_ERROR_NO_SERVER
	G_DBUS_ERROR_TIMEOUT
	G_DBUS_ERROR_NO_NETWORK
	G_DBUS_ERROR_ADDRESS_IN_USE
	G_DBUS_ERROR_DISCONNECTED
	G_DBUS_ERROR_INVALID_ARGS
	G_DBUS_ERROR_FILE_NOT_FOUND
	G_DBUS_ERROR_FILE_EXISTS
	G_DBUS_ERROR_UNKNOWN_METHOD
	G_DBUS_ERROR_TIMED_OUT
	G_DBUS_ERROR_MATCH_RULE_NOT_FOUND
	G_DBUS_ERROR_MATCH_RULE_INVALID
	G_DBUS_ERROR_SPAWN_EXEC_FAILED
	G_DBUS_ERROR_SPAWN_FORK_FAILED
	G_DBUS_ERROR_SPAWN_CHILD_EXITED
	G_DBUS_ERROR_SPAWN_CHILD_SIGNALED
	G_DBUS_ERROR_SPAWN_FAILED
	G_DBUS_ERROR_SPAWN_SETUP_FAILED
	G_DBUS_ERROR_SPAWN_CONFIG_INVALID
	G_DBUS_ERROR_SPAWN_SERVICE_INVALID
	G_DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND
	G_DBUS_ERROR_SPAWN_PERMISSIONS_INVALID
	G_DBUS_ERROR_SPAWN_FILE_INVALID
	G_DBUS_ERROR_SPAWN_NO_MEMORY
	G_DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN
	G_DBUS_ERROR_INVALID_SIGNATURE
	G_DBUS_ERROR_INVALID_FILE_CONTENT
	G_DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN
	G_DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN
	G_DBUS_ERROR_OBJECT_PATH_IN_USE
)

type GDBusConnectionFlags Enum

const (
	G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT GDBusConnectionFlags = 1 << iota
	G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER
	G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS
	G_DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION
	G_DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING
	G_DBUS_CONNECTION_FLAGS_NONE GDBusConnectionFlags = 0
)

type GDBusCapabilityFlags Enum

const (
	G_DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING GDBusCapabilityFlags = 1 << iota
	G_DBUS_CAPABILITY_FLAGS_NONE                                 = 0
)

type GDBusCallFlags Enum

const (
	G_DBUS_CALL_FLAGS_NO_AUTO_START GDBusCallFlags = 1 << iota
	G_DBUS_CALL_FLAGS_NONE          GDBusCallFlags = 0
)

type GDBusMessageType Enum

const (
	G_DBUS_MESSAGE_TYPE_INVALID GDBusMessageType = iota
	G_DBUS_MESSAGE_TYPE_METHOD_CALL
	G_DBUS_MESSAGE_TYPE_METHOD_RETURN
	G_DBUS_MESSAGE_TYPE_ERROR
	G_DBUS_MESSAGE_TYPE_SIGNAL
)

type GDBusMessageFlags Enum

const (
	G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED GDBusMessageFlags = 1 << iota
	G_DBUS_MESSAGE_FLAGS_NO_AUTO_START
	G_DBUS_MESSAGE_FLAGS_NONE GDBusMessageFlags = 0
)

type GDBusMessageHeaderField Enum

const (
	G_DBUS_MESSAGE_HEADER_FIELD_INVALID GDBusMessageHeaderField = iota
	G_DBUS_MESSAGE_HEADER_FIELD_PATH
	G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE
	G_DBUS_MESSAGE_HEADER_FIELD_MEMBER
	G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME
	G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL
	G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION
	G_DBUS_MESSAGE_HEADER_FIELD_SENDER
	G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE
	G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS
)

type GDBusPropertyInfoFlags Enum

const (
	G_DBUS_PROPERTY_INFO_FLAGS_READABLE GDBusPropertyInfoFlags = 1 << iota
	G_DBUS_PROPERTY_INFO_FLAGS_WRITABLE
	G_DBUS_PROPERTY_INFO_FLAGS_NONE GDBusPropertyInfoFlags = 0
)

type GDBusSubtreeFlags Enum

const (
	G_DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES GDBusSubtreeFlags = 1 << iota
	G_DBUS_SUBTREE_FLAGS_NONE                           GDBusSubtreeFlags = 0
)

type GDBusServerFlags Enum

const (
	G_DBUS_SERVER_FLAGS_RUN_IN_THREAD GDBusServerFlags = 1 << iota
	G_DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS
	G_DBUS_SERVER_FLAGS_NONE GDBusServerFlags = 0
)

type GDBusSignalFlags Enum

const (
	G_DBUS_SIGNAL_FLAGS_NO_MATCH_RULE GDBusSignalFlags = 1 << iota
	G_DBUS_SIGNAL_FLAGS_NONE          GDBusSignalFlags = 0
)

type GDBusSendMessageFlags Enum

const (
	G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL GDBusSendMessageFlags = 1 << iota
	G_DBUS_SEND_MESSAGE_FLAGS_NONE            GDBusSendMessageFlags = 0
)

type GCredentialsType Enum

const (
	G_CREDENTIALS_TYPE_INVALID GCredentialsType = iota
	G_CREDENTIALS_TYPE_LINUX_UCRED
	G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED
)

type GDBusMessageByteOrder Enum

const (
	G_DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN    GDBusMessageByteOrder = 'B'
	G_DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN GDBusMessageByteOrder = 'l'
)

type GApplicationFlags Enum

const (
	G_APPLICATION_IS_SERVICE GApplicationFlags = 1 << iota
	G_APPLICATION_IS_LAUNCHER
	G_APPLICATION_HANDLES_OPEN
	G_APPLICATION_HANDLES_COMMAND_LINE
	G_APPLICATION_SEND_ENVIRONMENT
	G_APPLICATION_FLAGS_NONE GApplicationFlags = 0
)

type GTlsError Enum

const (
	G_TLS_ERROR_UNAVAILABLE GTlsError = iota
	G_TLS_ERROR_MISC
	G_TLS_ERROR_BAD_CERTIFICATE
	G_TLS_ERROR_NOT_TLS
	G_TLS_ERROR_HANDSHAKE
	G_TLS_ERROR_CERTIFICATE_REQUIRED
	G_TLS_ERROR_EOF
)

type GTlsCertificateFlags Enum

const (
	G_TLS_CERTIFICATE_UNKNOWN_CA GTlsCertificateFlags = 1 << iota
	G_TLS_CERTIFICATE_BAD_IDENTITY
	G_TLS_CERTIFICATE_NOT_ACTIVATED
	G_TLS_CERTIFICATE_EXPIRED
	G_TLS_CERTIFICATE_REVOKED
	G_TLS_CERTIFICATE_INSECURE
	G_TLS_CERTIFICATE_GENERIC_ERROR
	G_TLS_CERTIFICATE_VALIDATE_ALL GTlsCertificateFlags = 0x007f
)

type GTlsAuthenticationMode Enum

const (
	G_TLS_AUTHENTICATION_NONE GTlsAuthenticationMode = iota
	G_TLS_AUTHENTICATION_REQUESTED
	G_TLS_AUTHENTICATION_REQUIRED
)

type GTlsRehandshakeMode Enum

const (
	G_TLS_REHANDSHAKE_NEVER GTlsRehandshakeMode = iota
	G_TLS_REHANDSHAKE_SAFELY
	G_TLS_REHANDSHAKE_UNSAFELY
)

type GModuleFlags Enum

const (
	G_MODULE_BIND_LAZY GModuleFlags = 1 << iota
	G_MODULE_BIND_LOCAL
	G_MODULE_BIND_MASK GModuleFlags = 0x03
)

type GSettingsBindFlags Enum

const (
	G_SETTINGS_BIND_GET GSettingsBindFlags = 1 << iota
	G_SETTINGS_BIND_SET
	G_SETTINGS_BIND_NO_SENSITIVITY
	G_SETTINGS_BIND_GET_NO_CHANGES
	G_SETTINGS_BIND_INVERT_BOOLEAN
	G_SETTINGS_BIND_DEFAULT GSettingsBindFlags = 0
)

type CairoStatus Enum

const (
	CAIRO_STATUS_SUCCESS CairoStatus = iota
	CAIRO_STATUS_NO_MEMORY
	CAIRO_STATUS_INVALID_RESTORE
	CAIRO_STATUS_INVALID_POP_GROUP
	CAIRO_STATUS_NO_CURRENT_POINT
	CAIRO_STATUS_INVALID_MATRIX
	CAIRO_STATUS_INVALID_STATUS
	CAIRO_STATUS_NULL_POINTER
	CAIRO_STATUS_INVALID_STRING
	CAIRO_STATUS_INVALID_PATH_DATA
	CAIRO_STATUS_READ_ERROR
	CAIRO_STATUS_WRITE_ERROR
	CAIRO_STATUS_SURFACE_FINISHED
	CAIRO_STATUS_SURFACE_TYPE_MISMATCH
	CAIRO_STATUS_PATTERN_TYPE_MISMATCH
	CAIRO_STATUS_INVALID_CONTENT
	CAIRO_STATUS_INVALID_FORMAT
	CAIRO_STATUS_INVALID_VISUAL
	CAIRO_STATUS_FILE_NOT_FOUND
	CAIRO_STATUS_INVALID_DASH
	CAIRO_STATUS_INVALID_DSC_COMMENT
	CAIRO_STATUS_INVALID_INDEX
	CAIRO_STATUS_CLIP_NOT_REPRESENTABLE
	CAIRO_STATUS_TEMP_FILE_ERROR
	CAIRO_STATUS_INVALID_STRIDE
	CAIRO_STATUS_FONT_TYPE_MISMATCH
	CAIRO_STATUS_USER_FONT_IMMUTABLE
	CAIRO_STATUS_USER_FONT_ERROR
	CAIRO_STATUS_NEGATIVE_COUNT
	CAIRO_STATUS_INVALID_CLUSTERS
	CAIRO_STATUS_INVALID_SLANT
	CAIRO_STATUS_INVALID_WEIGHT
	CAIRO_STATUS_INVALID_SIZE
	CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED
	CAIRO_STATUS_DEVICE_TYPE_MISMATCH
	CAIRO_STATUS_DEVICE_ERROR
	CAIRO_STATUS_LAST_STATUS
)

type CairoContent Enum

const (
	CAIRO_CONTENT_COLOR       CairoContent = 0x1000
	CAIRO_CONTENT_ALPHA       CairoContent = 0x2000
	CAIRO_CONTENT_COLOR_ALPHA CairoContent = 0x3000
)

type CairoOperator Enum

const (
	CAIRO_OPERATOR_CLEAR CairoOperator = iota
	CAIRO_OPERATOR_SOURCE
	CAIRO_OPERATOR_OVER
	CAIRO_OPERATOR_IN
	CAIRO_OPERATOR_OUT
	CAIRO_OPERATOR_ATOP
	CAIRO_OPERATOR_DEST
	CAIRO_OPERATOR_DEST_OVER
	CAIRO_OPERATOR_DEST_IN
	CAIRO_OPERATOR_DEST_OUT
	CAIRO_OPERATOR_DEST_ATOP
	CAIRO_OPERATOR_XOR
	CAIRO_OPERATOR_ADD
	CAIRO_OPERATOR_SATURATE
	CAIRO_OPERATOR_MULTIPLY
	CAIRO_OPERATOR_SCREEN
	CAIRO_OPERATOR_OVERLAY
	CAIRO_OPERATOR_DARKEN
	CAIRO_OPERATOR_LIGHTEN
	CAIRO_OPERATOR_COLOR_DODGE
	CAIRO_OPERATOR_COLOR_BURN
	CAIRO_OPERATOR_HARD_LIGHT
	CAIRO_OPERATOR_SOFT_LIGHT
	CAIRO_OPERATOR_DIFFERENCE
	CAIRO_OPERATOR_EXCLUSION
	CAIRO_OPERATOR_HSL_HUE
	CAIRO_OPERATOR_HSL_SATURATION
	CAIRO_OPERATOR_HSL_COLOR
	CAIRO_OPERATOR_HSL_LUMINOSITY
)

type CairoAntialias Enum

const (
	CAIRO_ANTIALIAS_DEFAULT CairoAntialias = iota
	CAIRO_ANTIALIAS_NONE
	CAIRO_ANTIALIAS_GRAY
	CAIRO_ANTIALIAS_SUBPIXEL
)

type CairoFillRule Enum

const (
	CAIRO_FILL_RULE_WINDING CairoFillRule = iota
	CAIRO_FILL_RULE_EVEN_ODD
)

type CairoLineCap Enum

const (
	CAIRO_LINE_CAP_BUTT CairoLineCap = iota
	CAIRO_LINE_CAP_ROUND
	CAIRO_LINE_CAP_SQUARE
)

type CairoLineJoin Enum

const (
	CAIRO_LINE_JOIN_MITER CairoLineJoin = iota
	CAIRO_LINE_JOIN_ROUND
	CAIRO_LINE_JOIN_BEVEL
)

type CairoTextClusterFlags Enum

const (
	CAIRO_TEXT_CLUSTER_FLAG_BACKWARD CairoTextClusterFlags = 1 << iota
)

type CairoFontSlant Enum

const (
	CAIRO_FONT_SLANT_NORMAL CairoFontSlant = iota
	CAIRO_FONT_SLANT_ITALIC
	CAIRO_FONT_SLANT_OBLIQUE
)

type CairoFontWeight Enum

const (
	CAIRO_FONT_WEIGHT_NORMAL CairoFontWeight = iota
	CAIRO_FONT_WEIGHT_BOLD
)

type CairoSubpixelOrder Enum

const (
	CAIRO_SUBPIXEL_ORDER_DEFAULT CairoSubpixelOrder = iota
	CAIRO_SUBPIXEL_ORDER_RGB
	CAIRO_SUBPIXEL_ORDER_BGR
	CAIRO_SUBPIXEL_ORDER_VRGB
	CAIRO_SUBPIXEL_ORDER_VBGR
)

type CairoHintStyle Enum

const (
	CAIRO_HINT_STYLE_DEFAULT CairoHintStyle = iota
	CAIRO_HINT_STYLE_NONE
	CAIRO_HINT_STYLE_SLIGHT
	CAIRO_HINT_STYLE_MEDIUM
	CAIRO_HINT_STYLE_FULL
)

type CairoHintMetrics Enum

const (
	CAIRO_HINT_METRICS_DEFAULT CairoHintMetrics = iota
	CAIRO_HINT_METRICS_OFF
	CAIRO_HINT_METRICS_ON
)

type CairoFontType Enum

const (
	CAIRO_FONT_TYPE_TOY CairoFontType = iota
	CAIRO_FONT_TYPE_FT
	CAIRO_FONT_TYPE_WIN32
	CAIRO_FONT_TYPE_QUARTZ
	CAIRO_FONT_TYPE_USER
)

type Cairo_path_data_type_t Enum

const (
	CAIRO_PATH_MOVE_TO Cairo_path_data_type_t = iota
	CAIRO_PATH_LINE_TO
	CAIRO_PATH_CURVE_TO
	CAIRO_PATH_CLOSE_PATH
)

type CairoDeviceType Enum

const (
	CAIRO_DEVICE_TYPE_DRM CairoDeviceType = iota
	CAIRO_DEVICE_TYPE_GL
	CAIRO_DEVICE_TYPE_SCRIPT
	CAIRO_DEVICE_TYPE_XCB
	CAIRO_DEVICE_TYPE_XLIB
	CAIRO_DEVICE_TYPE_XML
)

type CairoSurfaceType Enum

const (
	CAIRO_SURFACE_TYPE_IMAGE CairoSurfaceType = iota
	CAIRO_SURFACE_TYPE_PDF
	CAIRO_SURFACE_TYPE_PS
	CAIRO_SURFACE_TYPE_XLIB
	CAIRO_SURFACE_TYPE_XCB
	CAIRO_SURFACE_TYPE_GLITZ
	CAIRO_SURFACE_TYPE_QUARTZ
	CAIRO_SURFACE_TYPE_WIN32
	CAIRO_SURFACE_TYPE_BEOS
	CAIRO_SURFACE_TYPE_DIRECTFB
	CAIRO_SURFACE_TYPE_SVG
	CAIRO_SURFACE_TYPE_OS2
	CAIRO_SURFACE_TYPE_WIN32_PRINTING
	CAIRO_SURFACE_TYPE_QUARTZ_IMAGE
	CAIRO_SURFACE_TYPE_SCRIPT
	CAIRO_SURFACE_TYPE_QT
	CAIRO_SURFACE_TYPE_RECORDING
	CAIRO_SURFACE_TYPE_VG
	CAIRO_SURFACE_TYPE_GL
	CAIRO_SURFACE_TYPE_DRM
	CAIRO_SURFACE_TYPE_TEE
	CAIRO_SURFACE_TYPE_XML
	CAIRO_SURFACE_TYPE_SKIA
	CAIRO_SURFACE_TYPE_SUBSURFACE
)

type CairoFormat Enum

const (
	CAIRO_FORMAT_INVALID CairoFormat = iota - 1
	CAIRO_FORMAT_ARGB32
	CAIRO_FORMAT_RGB24
	CAIRO_FORMAT_A8
	CAIRO_FORMAT_A1
	CAIRO_FORMAT_RGB16_565
)

type CairoPatternType Enum

const (
	CAIRO_PATTERN_TYPE_SOLID CairoPatternType = iota
	CAIRO_PATTERN_TYPE_SURFACE
	CAIRO_PATTERN_TYPE_LINEAR
	CAIRO_PATTERN_TYPE_RADIAL
)

type CairoExtend Enum

const (
	CAIRO_EXTEND_NONE CairoExtend = iota
	CAIRO_EXTEND_REPEAT
	CAIRO_EXTEND_REFLECT
	CAIRO_EXTEND_PAD
)

type CairoFilter Enum

const (
	CAIRO_FILTER_FAST CairoFilter = iota
	CAIRO_FILTER_GOOD
	CAIRO_FILTER_BEST
	CAIRO_FILTER_NEAREST
	CAIRO_FILTER_BILINEAR
	CAIRO_FILTER_GAUSSIAN
)

type CairoRegionOverlap Enum

const (
	CAIRO_REGION_OVERLAP_IN CairoRegionOverlap = iota
	CAIRO_REGION_OVERLAP_OUT
	CAIRO_REGION_OVERLAP_PART
)

type PangoCoverageLevel Enum

const (
	PANGO_COVERAGE_NONE PangoCoverageLevel = iota
	PANGO_COVERAGE_FALLBACK
	PANGO_COVERAGE_APPROXIMATE
	PANGO_COVERAGE_EXACT
)

type PangoGravity Enum

const (
	PANGO_GRAVITY_SOUTH PangoGravity = iota
	PANGO_GRAVITY_EAST
	PANGO_GRAVITY_NORTH
	PANGO_GRAVITY_WEST
	PANGO_GRAVITY_AUTO
)

type PangoGravityHint Enum

const (
	PANGO_GRAVITY_HINT_NATURAL PangoGravityHint = iota
	PANGO_GRAVITY_HINT_STRONG
	PANGO_GRAVITY_HINT_LINE
)

type PangoScript Enum

const (
	PANGO_SCRIPT_INVALID_CODE PangoScript = iota - 1
	PANGO_SCRIPT_COMMON
	PANGO_SCRIPT_INHERITED
	PANGO_SCRIPT_ARABIC
	PANGO_SCRIPT_ARMENIAN
	PANGO_SCRIPT_BENGALI
	PANGO_SCRIPT_BOPOMOFO
	PANGO_SCRIPT_CHEROKEE
	PANGO_SCRIPT_COPTIC
	PANGO_SCRIPT_CYRILLIC
	PANGO_SCRIPT_DESERET
	PANGO_SCRIPT_DEVANAGARI
	PANGO_SCRIPT_ETHIOPIC
	PANGO_SCRIPT_GEORGIAN
	PANGO_SCRIPT_GOTHIC
	PANGO_SCRIPT_GREEK
	PANGO_SCRIPT_GUJARATI
	PANGO_SCRIPT_GURMUKHI
	PANGO_SCRIPT_HAN
	PANGO_SCRIPT_HANGUL
	PANGO_SCRIPT_HEBREW
	PANGO_SCRIPT_HIRAGANA
	PANGO_SCRIPT_KANNADA
	PANGO_SCRIPT_KATAKANA
	PANGO_SCRIPT_KHMER
	PANGO_SCRIPT_LAO
	PANGO_SCRIPT_LATIN
	PANGO_SCRIPT_MALAYALAM
	PANGO_SCRIPT_MONGOLIAN
	PANGO_SCRIPT_MYANMAR
	PANGO_SCRIPT_OGHAM
	PANGO_SCRIPT_OLD_ITALIC
	PANGO_SCRIPT_ORIYA
	PANGO_SCRIPT_RUNIC
	PANGO_SCRIPT_SINHALA
	PANGO_SCRIPT_SYRIAC
	PANGO_SCRIPT_TAMIL
	PANGO_SCRIPT_TELUGU
	PANGO_SCRIPT_THAANA
	PANGO_SCRIPT_THAI
	PANGO_SCRIPT_TIBETAN
	PANGO_SCRIPT_CANADIAN_ABORIGINAL
	PANGO_SCRIPT_YI
	PANGO_SCRIPT_TAGALOG
	PANGO_SCRIPT_HANUNOO
	PANGO_SCRIPT_BUHID
	PANGO_SCRIPT_TAGBANWA
	PANGO_SCRIPT_BRAILLE
	PANGO_SCRIPT_CYPRIOT
	PANGO_SCRIPT_LIMBU
	PANGO_SCRIPT_OSMANYA
	PANGO_SCRIPT_SHAVIAN
	PANGO_SCRIPT_LINEAR_B
	PANGO_SCRIPT_TAI_LE
	PANGO_SCRIPT_UGARITIC
	PANGO_SCRIPT_NEW_TAI_LUE
	PANGO_SCRIPT_BUGINESE
	PANGO_SCRIPT_GLAGOLITIC
	PANGO_SCRIPT_TIFINAGH
	PANGO_SCRIPT_SYLOTI_NAGRI
	PANGO_SCRIPT_OLD_PERSIAN
	PANGO_SCRIPT_KHAROSHTHI
	PANGO_SCRIPT_UNKNOWN
	PANGO_SCRIPT_BALINESE
	PANGO_SCRIPT_CUNEIFORM
	PANGO_SCRIPT_PHOENICIAN
	PANGO_SCRIPT_PHAGS_PA
	PANGO_SCRIPT_NKO
	PANGO_SCRIPT_KAYAH_LI
	PANGO_SCRIPT_LEPCHA
	PANGO_SCRIPT_REJANG
	PANGO_SCRIPT_SUNDANESE
	PANGO_SCRIPT_SAURASHTRA
	PANGO_SCRIPT_CHAM
	PANGO_SCRIPT_OL_CHIKI
	PANGO_SCRIPT_VAI
	PANGO_SCRIPT_CARIAN
	PANGO_SCRIPT_LYCIAN
	PANGO_SCRIPT_LYDIAN
)

type PangoBidiType Enum

const (
	PANGO_BIDI_TYPE_L PangoBidiType = iota
	PANGO_BIDI_TYPE_LRE
	PANGO_BIDI_TYPE_LRO
	PANGO_BIDI_TYPE_R
	PANGO_BIDI_TYPE_AL
	PANGO_BIDI_TYPE_RLE
	PANGO_BIDI_TYPE_RLO
	PANGO_BIDI_TYPE_PDF
	PANGO_BIDI_TYPE_EN
	PANGO_BIDI_TYPE_ES
	PANGO_BIDI_TYPE_ET
	PANGO_BIDI_TYPE_AN
	PANGO_BIDI_TYPE_CS
	PANGO_BIDI_TYPE_NSM
	PANGO_BIDI_TYPE_BN
	PANGO_BIDI_TYPE_B
	PANGO_BIDI_TYPE_S
	PANGO_BIDI_TYPE_WS
	PANGO_BIDI_TYPE_ON
)

type PangoDirection Enum

const (
	PANGO_DIRECTION_LTR PangoDirection = iota
	PANGO_DIRECTION_RTL
	PANGO_DIRECTION_TTB_LTR
	PANGO_DIRECTION_TTB_RTL
	PANGO_DIRECTION_WEAK_LTR
	PANGO_DIRECTION_WEAK_RTL
	PANGO_DIRECTION_NEUTRAL
)

type PangoStyle Enum

const (
	PANGO_STYLE_NORMAL PangoStyle = iota
	PANGO_STYLE_OBLIQUE
	PANGO_STYLE_ITALIC
)

type PangoVariant Enum

const (
	PANGO_VARIANT_NORMAL PangoVariant = iota
	PANGO_VARIANT_SMALL_CAPS
)

type PangoWeight Enum

const (
	PANGO_WEIGHT_THIN       PangoWeight = 100
	PANGO_WEIGHT_ULTRALIGHT PangoWeight = 200
	PANGO_WEIGHT_LIGHT      PangoWeight = 300
	PANGO_WEIGHT_BOOK       PangoWeight = 380
	PANGO_WEIGHT_NORMAL     PangoWeight = 400
	PANGO_WEIGHT_MEDIUM     PangoWeight = 500
	PANGO_WEIGHT_SEMIBOLD   PangoWeight = 600
	PANGO_WEIGHT_BOLD       PangoWeight = 700
	PANGO_WEIGHT_ULTRABOLD  PangoWeight = 800
	PANGO_WEIGHT_HEAVY      PangoWeight = 900
	PANGO_WEIGHT_ULTRAHEAVY PangoWeight = 1000
)

type PangoStretch Enum

const (
	PANGO_STRETCH_ULTRA_CONDENSED PangoStretch = iota
	PANGO_STRETCH_EXTRA_CONDENSED
	PANGO_STRETCH_CONDENSED
	PANGO_STRETCH_SEMI_CONDENSED
	PANGO_STRETCH_NORMAL
	PANGO_STRETCH_SEMI_EXPANDED
	PANGO_STRETCH_EXPANDED
	PANGO_STRETCH_EXTRA_EXPANDED
	PANGO_STRETCH_ULTRA_EXPANDED
)

type PangoFontMask Enum

const (
	PANGO_FONT_MASK_FAMILY PangoFontMask = 1 << iota
	PANGO_FONT_MASK_STYLE
	PANGO_FONT_MASK_VARIANT
	PANGO_FONT_MASK_WEIGHT
	PANGO_FONT_MASK_STRETCH
	PANGO_FONT_MASK_SIZE
	PANGO_FONT_MASK_GRAVITY
)

type PangoAttrType Enum

const (
	PANGO_ATTR_INVALID PangoAttrType = iota
	PANGO_ATTR_LANGUAGE
	PANGO_ATTR_FAMILY
	PANGO_ATTR_STYLE
	PANGO_ATTR_WEIGHT
	PANGO_ATTR_VARIANT
	PANGO_ATTR_STRETCH
	PANGO_ATTR_SIZE
	PANGO_ATTR_FONT_DESC
	PANGO_ATTR_FOREGROUND
	PANGO_ATTR_BACKGROUND
	PANGO_ATTR_UNDERLINE
	PANGO_ATTR_STRIKETHROUGH
	PANGO_ATTR_RISE
	PANGO_ATTR_SHAPE
	PANGO_ATTR_SCALE
	PANGO_ATTR_FALLBACK
	PANGO_ATTR_LETTER_SPACING
	PANGO_ATTR_UNDERLINE_COLOR
	PANGO_ATTR_STRIKETHROUGH_COLOR
	PANGO_ATTR_ABSOLUTE_SIZE
	PANGO_ATTR_GRAVITY
	PANGO_ATTR_GRAVITY_HINT
)

type PangoUnderline Enum

const (
	PANGO_UNDERLINE_NONE PangoUnderline = iota
	PANGO_UNDERLINE_SINGLE
	PANGO_UNDERLINE_DOUBLE
	PANGO_UNDERLINE_LOW
	PANGO_UNDERLINE_ERROR
)

type PangoTabAlign Enum

const PANGO_TAB_LEFT PangoTabAlign = 0

type PangoAlignment Enum

const (
	PANGO_ALIGN_LEFT PangoAlignment = iota
	PANGO_ALIGN_CENTER
	PANGO_ALIGN_RIGHT
)

type PangoWrapMode Enum

const (
	PANGO_WRAP_WORD PangoWrapMode = iota
	PANGO_WRAP_CHAR
	PANGO_WRAP_WORD_CHAR
)

type PangoEllipsizeMode Enum

const (
	PANGO_ELLIPSIZE_NONE PangoEllipsizeMode = iota
	PANGO_ELLIPSIZE_START
	PANGO_ELLIPSIZE_MIDDLE
	PANGO_ELLIPSIZE_END
)

type PangoRenderPart Enum

const (
	PANGO_RENDER_PART_FOREGROUND PangoRenderPart = iota
	PANGO_RENDER_PART_BACKGROUND
	PANGO_RENDER_PART_UNDERLINE
	PANGO_RENDER_PART_STRIKETHROUGH
)

type GdkByteOrder Enum

const (
	GDK_LSB_FIRST GdkByteOrder = iota
	GDK_MSB_FIRST
)

type GdkModifierType Enum

const (
	GDK_SHIFT_MASK GdkModifierType = 1 << 0
	GDK_LOCK_MASK
	GDK_CONTROL_MASK
	GDK_MOD1_MASK
	GDK_MOD2_MASK
	GDK_MOD3_MASK
	GDK_MOD4_MASK
	GDK_MOD5_MASK
	GDK_BUTTON1_MASK
	GDK_BUTTON2_MASK
	GDK_BUTTON3_MASK
	GDK_BUTTON4_MASK
	GDK_BUTTON5_MASK
)
const (
	GDK_SUPER_MASK GdkModifierType = 1 << (26 + iota)
	GDK_HYPER_MASK
	GDK_META_MASK
	GDK_RELEASE_MASK
	GDK_MODIFIER_MASK GdkModifierType = 0x5c001fff
)

type GdkInputCondition Enum

const (
	GDK_INPUT_READ GdkInputCondition = 1 << iota
	GDK_INPUT_WRITE
	GDK_INPUT_EXCEPTION
)

type GdkStatus Enum

const (
	GDK_OK    GdkStatus = 0
	GDK_ERROR GdkStatus = -iota
	GDK_ERROR_PARAM
	GDK_ERROR_FILE
	GDK_ERROR_MEM
)

type GdkInputSource Enum

const (
	GDK_SOURCE_MOUSE GdkInputSource = iota
	GDK_SOURCE_PEN
	GDK_SOURCE_ERASER
	GDK_SOURCE_CURSOR
)

type GdkInputMode Enum

const (
	GDK_MODE_DISABLED GdkInputMode = iota
	GDK_MODE_SCREEN
	GDK_MODE_WINDOW
)

type GdkAxisUse Enum

const (
	GDK_AXIS_IGNORE GdkAxisUse = iota
	GDK_AXIS_X
	GDK_AXIS_Y
	GDK_AXIS_PRESSURE
	GDK_AXIS_XTILT
	GDK_AXIS_YTILT
	GDK_AXIS_WHEEL
	GDK_AXIS_LAST
)

type GdkFilterReturn Enum

const (
	GDK_FILTER_CONTINUE GdkFilterReturn = iota
	GDK_FILTER_TRANSLATE
	GDK_FILTER_REMOVE
)

type GdkEventType Enum // REMOVE

type GdkScrollDirection Enum

const (
	GDK_SCROLL_UP GdkScrollDirection = iota
	GDK_SCROLL_DOWN
	GDK_SCROLL_LEFT
	GDK_SCROLL_RIGHT
)

type GdkNotifyType Enum

const (
	GDK_NOTIFY_ANCESTOR GdkNotifyType = iota
	GDK_NOTIFY_VIRTUAL
	GDK_NOTIFY_INFERIOR
	GDK_NOTIFY_NONLINEAR
	GDK_NOTIFY_NONLINEAR_VIRTUAL
	GDK_NOTIFY_UNKNOWN
)

type GdkSettingAction Enum

const (
	GDK_SETTING_ACTION_NEW GdkSettingAction = iota
	GDK_SETTING_ACTION_CHANGED
	GDK_SETTING_ACTION_DELETED
)

type GdkOwnerChange Enum

const (
	GDK_OWNER_CHANGE_NEW_OWNER GdkOwnerChange = iota
	GDK_OWNER_CHANGE_DESTROY
	GDK_OWNER_CHANGE_CLOSE
)

type GdkRgbDither Enum

const (
	GDK_RGB_DITHER_NONE GdkRgbDither = iota
	GDK_RGB_DITHER_NORMAL
	GDK_RGB_DITHER_MAX
)

type GdkPixbufError Enum

const (
	GDK_PIXBUF_ERROR_CORRUPT_IMAGE GdkPixbufError = iota
	GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY
	GDK_PIXBUF_ERROR_BAD_OPTION
	GDK_PIXBUF_ERROR_UNKNOWN_TYPE
	GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION
	GDK_PIXBUF_ERROR_FAILED
)

type GdkInterpType Enum

const (
	GDK_INTERP_NEAREST GdkInterpType = iota
	GDK_INTERP_TILES
	GDK_INTERP_BILINEAR
	GDK_INTERP_HYPER
)

type GdkFill Enum

const (
	GDK_SOLID GdkFill = iota
	GDK_TILED
	GDK_STIPPLED
	GDK_OPAQUE_STIPPLED
)

type GdkFunction Enum

const (
	GDK_COPY GdkFunction = iota
	GDK_INVERT
	GDK_XOR
	GDK_CLEAR
	GDK_AND
	GDK_AND_REVERSE
	GDK_AND_INVERT
	GDK_NOOP
	GDK_OR
	GDK_EQUIV
	GDK_OR_REVERSE
	GDK_COPY_INVERT
	GDK_OR_INVERT
	GDK_NAND
	GDK_NOR
	GDK_SET
)

type GdkJoinStyle Enum

const (
	GDK_JOIN_MITER GdkJoinStyle = iota
	GDK_JOIN_ROUND
	GDK_JOIN_BEVEL
)

type GdkLineStyle Enum

const (
	GDK_LINE_SOLID GdkLineStyle = iota
	GDK_LINE_ON_OFF_DASH
	GDK_LINE_DOUBLE_DASH
)

type GdkSubwindowMode Enum

const (
	GDK_CLIP_BY_CHILDREN GdkSubwindowMode = iota
	GDK_INCLUDE_INFERIORS
)

type GdkFillRule Enum

const (
	GDK_EVEN_ODD_RULE GdkFillRule = iota
	GDK_WINDING_RULE
)

type GdkOverlapType Enum

const (
	GDK_OVERLAP_RECTANGLE_IN GdkOverlapType = iota
	GDK_OVERLAP_RECTANGLE_OUT
	GDK_OVERLAP_RECTANGLE_PART
)

type GdkWindowAttributesType Enum

const (
	GDK_WA_TITLE GdkWindowAttributesType = 1 << (iota + 1)
	GDK_WA_X
	GDK_WA_Y
	GDK_WA_CURSOR
	GDK_WA_COLORMAP
	GDK_WA_VISUAL
	GDK_WA_WMCLASS
	GDK_WA_NOREDIR
	GDK_WA_TYPE_HINT
)

type GdkWMDecoration Enum

const (
	GDK_DECOR_ALL GdkWMDecoration = 1 << iota
	GDK_DECOR_BORDER
	GDK_DECOR_RESIZEH
	GDK_DECOR_TITLE
	GDK_DECOR_MENU
	GDK_DECOR_MINIMIZE
	GDK_DECOR_MAXIMIZE
)

type GdkWMFunction Enum

const (
	GDK_FUNC_ALL GdkWMFunction = 1 << iota
	GDK_FUNC_RESIZE
	GDK_FUNC_MOVE
	GDK_FUNC_MINIMIZE
	GDK_FUNC_MAXIMIZE
	GDK_FUNC_CLOSE
)

type GtkScrollStep Enum

const (
	GTK_SCROLL_STEPS GtkScrollStep = iota
	GTK_SCROLL_PAGES
	GTK_SCROLL_ENDS
	GTK_SCROLL_HORIZONTAL_STEPS
	GTK_SCROLL_HORIZONTAL_PAGES
	GTK_SCROLL_HORIZONTAL_ENDS
)

type GtkIMPreeditStyle Enum

const (
	GTK_IM_PREEDIT_NOTHING GtkIMPreeditStyle = iota
	GTK_IM_PREEDIT_CALLBACK
	GTK_IM_PREEDIT_NONE
)

type GtkIMStatusStyle Enum

const (
	GTK_IM_STATUS_NOTHING GtkIMStatusStyle = iota
	GTK_IM_STATUS_CALLBACK
	GTK_IM_STATUS_NONE
)

type GtkRcTokenType Enum

const (
	GTK_RC_TOKEN_INVALID GtkRcTokenType = GtkRcTokenType(G_TOKEN_LAST) + iota
	GTK_RC_TOKEN_INCLUDE
	GTK_RC_TOKEN_NORMAL
	GTK_RC_TOKEN_ACTIVE
	GTK_RC_TOKEN_PRELIGHT
	GTK_RC_TOKEN_SELECTED
	GTK_RC_TOKEN_INSENSITIVE
	GTK_RC_TOKEN_FG
	GTK_RC_TOKEN_BG
	GTK_RC_TOKEN_TEXT
	GTK_RC_TOKEN_BASE
	GTK_RC_TOKEN_XTHICKNESS
	GTK_RC_TOKEN_YTHICKNESS
	GTK_RC_TOKEN_FONT
	GTK_RC_TOKEN_FONTSET
	GTK_RC_TOKEN_FONT_NAME
	GTK_RC_TOKEN_BG_PIXMAP
	GTK_RC_TOKEN_PIXMAP_PATH
	GTK_RC_TOKEN_STYLE
	GTK_RC_TOKEN_BINDING
	GTK_RC_TOKEN_BIND
	GTK_RC_TOKEN_WIDGET
	GTK_RC_TOKEN_WIDGET_CLASS
	GTK_RC_TOKEN_CLASS
	GTK_RC_TOKEN_LOWEST
	GTK_RC_TOKEN_GTK
	GTK_RC_TOKEN_APPLICATION
	GTK_RC_TOKEN_THEME
	GTK_RC_TOKEN_RC
	GTK_RC_TOKEN_HIGHEST
	GTK_RC_TOKEN_ENGINE
	GTK_RC_TOKEN_MODULE_PATH
	GTK_RC_TOKEN_IM_MODULE_PATH
	GTK_RC_TOKEN_IM_MODULE_FILE
	GTK_RC_TOKEN_STOCK
	GTK_RC_TOKEN_LTR
	GTK_RC_TOKEN_RTL
	GTK_RC_TOKEN_COLOR
	GTK_RC_TOKEN_UNBIND
	GTK_RC_TOKEN_LAST
)

type AtkStateType Enum

const (
	ATK_STATE_INVALID AtkStateType = iota
	ATK_STATE_ACTIVE
	ATK_STATE_ARMED
	ATK_STATE_BUSY
	ATK_STATE_CHECKED
	ATK_STATE_DEFUNCT
	ATK_STATE_EDITABLE
	ATK_STATE_ENABLED
	ATK_STATE_EXPANDABLE
	ATK_STATE_EXPANDED
	ATK_STATE_FOCUSABLE
	ATK_STATE_FOCUSED
	ATK_STATE_HORIZONTAL
	ATK_STATE_ICONIFIED
	ATK_STATE_MODAL
	ATK_STATE_MULTI_LINE
	ATK_STATE_MULTISELECTABLE
	ATK_STATE_OPAQUE
	ATK_STATE_PRESSED
	ATK_STATE_RESIZABLE
	ATK_STATE_SELECTABLE
	ATK_STATE_SELECTED
	ATK_STATE_SENSITIVE
	ATK_STATE_SHOWING
	ATK_STATE_SINGLE_LINE
	ATK_STATE_STALE
	ATK_STATE_TRANSIENT
	ATK_STATE_VERTICAL
	ATK_STATE_VISIBLE
	ATK_STATE_MANAGES_DESCENDANTS
	ATK_STATE_INDETERMINATE
	ATK_STATE_TRUNCATED
	ATK_STATE_REQUIRED
	ATK_STATE_INVALID_ENTRY
	ATK_STATE_SUPPORTS_AUTOCOMPLETION
	ATK_STATE_SELECTABLE_TEXT
	ATK_STATE_DEFAULT
	ATK_STATE_ANIMATED
	ATK_STATE_VISITED
	ATK_STATE_LAST_DEFINED
)

type AtkRelationType Enum

const (
	ATK_RELATION_NULL AtkRelationType = iota
	ATK_RELATION_CONTROLLED_BY
	ATK_RELATION_CONTROLLER_FOR
	ATK_RELATION_LABEL_FOR
	ATK_RELATION_LABELLED_BY
	ATK_RELATION_MEMBER_OF
	ATK_RELATION_NODE_CHILD_OF
	ATK_RELATION_FLOWS_TO
	ATK_RELATION_FLOWS_FROM
	ATK_RELATION_SUBWINDOW_OF
	ATK_RELATION_EMBEDS
	ATK_RELATION_EMBEDDED_BY
	ATK_RELATION_POPUP_FOR
	ATK_RELATION_PARENT_WINDOW_OF
	ATK_RELATION_DESCRIBED_BY
	ATK_RELATION_DESCRIPTION_FOR
	ATK_RELATION_NODE_PARENT_OF
	ATK_RELATION_LAST_DEFINED
)

type AtkRole Enum

const (
	ATK_ROLE_INVALID AtkRole = iota
	ATK_ROLE_ACCEL_LABEL
	ATK_ROLE_ALERT
	ATK_ROLE_ANIMATION
	ATK_ROLE_ARROW
	ATK_ROLE_CALENDAR
	ATK_ROLE_CANVAS
	ATK_ROLE_CHECK_BOX
	ATK_ROLE_CHECK_MENU_ITEM
	ATK_ROLE_COLOR_CHOOSER
	ATK_ROLE_COLUMN_HEADER
	ATK_ROLE_COMBO_BOX
	ATK_ROLE_DATE_EDITOR
	ATK_ROLE_DESKTOP_ICON
	ATK_ROLE_DESKTOP_FRAME
	ATK_ROLE_DIAL
	ATK_ROLE_DIALOG
	ATK_ROLE_DIRECTORY_PANE
	ATK_ROLE_DRAWING_AREA
	ATK_ROLE_FILE_CHOOSER
	ATK_ROLE_FILLER
	ATK_ROLE_FONT_CHOOSER
	ATK_ROLE_FRAME
	ATK_ROLE_GLASS_PANE
	ATK_ROLE_HTML_CONTAINER
	ATK_ROLE_ICON
	ATK_ROLE_IMAGE
	ATK_ROLE_INTERNAL_FRAME
	ATK_ROLE_LABEL
	ATK_ROLE_LAYERED_PANE
	ATK_ROLE_LIST
	ATK_ROLE_LIST_ITEM
	ATK_ROLE_MENU
	ATK_ROLE_MENU_BAR
	ATK_ROLE_MENU_ITEM
	ATK_ROLE_OPTION_PANE
	ATK_ROLE_PAGE_TAB
	ATK_ROLE_PAGE_TAB_LIST
	ATK_ROLE_PANEL
	ATK_ROLE_PASSWORD_TEXT
	ATK_ROLE_POPUP_MENU
	ATK_ROLE_PROGRESS_BAR
	ATK_ROLE_PUSH_BUTTON
	ATK_ROLE_RADIO_BUTTON
	ATK_ROLE_RADIO_MENU_ITEM
	ATK_ROLE_ROOT_PANE
	ATK_ROLE_ROW_HEADER
	ATK_ROLE_SCROLL_BAR
	ATK_ROLE_SCROLL_PANE
	ATK_ROLE_SEPARATOR
	ATK_ROLE_SLIDER
	ATK_ROLE_SPLIT_PANE
	ATK_ROLE_SPIN_BUTTON
	ATK_ROLE_STATUSBAR
	ATK_ROLE_TABLE
	ATK_ROLE_TABLE_CELL
	ATK_ROLE_TABLE_COLUMN_HEADER
	ATK_ROLE_TABLE_ROW_HEADER
	ATK_ROLE_TEAR_OFF_MENU_ITEM
	ATK_ROLE_TERMINAL
	ATK_ROLE_TEXT
	ATK_ROLE_TOGGLE_BUTTON
	ATK_ROLE_TOOL_BAR
	ATK_ROLE_TOOL_TIP
	ATK_ROLE_TREE
	ATK_ROLE_TREE_TABLE
	ATK_ROLE_UNKNOWN
	ATK_ROLE_VIEWPORT
	ATK_ROLE_WINDOW
	ATK_ROLE_HEADER
	ATK_ROLE_FOOTER
	ATK_ROLE_PARAGRAPH
	ATK_ROLE_RULER
	ATK_ROLE_APPLICATION
	ATK_ROLE_AUTOCOMPLETE
	ATK_ROLE_EDITBAR
	ATK_ROLE_EMBEDDED
	ATK_ROLE_ENTRY
	ATK_ROLE_CHART
	ATK_ROLE_CAPTION
	ATK_ROLE_DOCUMENT_FRAME
	ATK_ROLE_HEADING
	ATK_ROLE_PAGE
	ATK_ROLE_SECTION
	ATK_ROLE_REDUNDANT_OBJECT
	ATK_ROLE_FORM
	ATK_ROLE_LINK
	ATK_ROLE_INPUT_METHOD_WINDOW
	ATK_ROLE_LAST_DEFINED
)

type AtkLayer Enum

const (
	ATK_LAYER_INVALID AtkLayer = iota
	ATK_LAYER_BACKGROUND
	ATK_LAYER_CANVAS
	ATK_LAYER_WIDGET
	ATK_LAYER_MDI
	ATK_LAYER_POPUP
	ATK_LAYER_OVERLAY
	ATK_LAYER_WINDOW
)

type AtkKeyEventType Enum

const (
	ATK_KEY_EVENT_PRESS AtkKeyEventType = iota
	ATK_KEY_EVENT_RELEASE
	ATK_KEY_EVENT_LAST_DEFINED
)

type AtkCoordType Enum

const (
	ATK_XY_SCREEN AtkCoordType = iota
	ATK_XY_WINDOW
)

type AtkTextAttribute Enum

const (
	ATK_TEXT_ATTR_INVALID AtkTextAttribute = iota
	ATK_TEXT_ATTR_LEFT_MARGIN
	ATK_TEXT_ATTR_RIGHT_MARGIN
	ATK_TEXT_ATTR_INDENT
	ATK_TEXT_ATTR_INVISIBLE
	ATK_TEXT_ATTR_EDITABLE
	ATK_TEXT_ATTR_PIXELS_ABOVE_LINES
	ATK_TEXT_ATTR_PIXELS_BELOW_LINES
	ATK_TEXT_ATTR_PIXELS_INSIDE_WRAP
	ATK_TEXT_ATTR_BG_FULL_HEIGHT
	ATK_TEXT_ATTR_RISE
	ATK_TEXT_ATTR_UNDERLINE
	ATK_TEXT_ATTR_STRIKETHROUGH
	ATK_TEXT_ATTR_SIZE
	ATK_TEXT_ATTR_SCALE
	ATK_TEXT_ATTR_WEIGHT
	ATK_TEXT_ATTR_LANGUAGE
	ATK_TEXT_ATTR_FAMILY_NAME
	ATK_TEXT_ATTR_BG_COLOR
	ATK_TEXT_ATTR_FG_COLOR
	ATK_TEXT_ATTR_BG_STIPPLE
	ATK_TEXT_ATTR_FG_STIPPLE
	ATK_TEXT_ATTR_WRAP_MODE
	ATK_TEXT_ATTR_DIRECTION
	ATK_TEXT_ATTR_JUSTIFICATION
	ATK_TEXT_ATTR_STRETCH
	ATK_TEXT_ATTR_VARIANT
	ATK_TEXT_ATTR_STYLE
	ATK_TEXT_ATTR_LAST_DEFINED
)

type AtkTextBoundary Enum

const (
	ATK_TEXT_BOUNDARY_CHAR AtkTextBoundary = iota
	ATK_TEXT_BOUNDARY_WORD_START
	ATK_TEXT_BOUNDARY_WORD_END
	ATK_TEXT_BOUNDARY_SENTENCE_START
	ATK_TEXT_BOUNDARY_SENTENCE_END
	ATK_TEXT_BOUNDARY_LINE_START
	ATK_TEXT_BOUNDARY_LINE_END
)

type AtkTextClipType Enum

const (
	ATK_TEXT_CLIP_NONE AtkTextClipType = iota
	ATK_TEXT_CLIP_MIN
	ATK_TEXT_CLIP_MAX
	ATK_TEXT_CLIP_BOTH
)

type AtkHyperlinkStateFlags Enum

const ATK_HYPERLINK_IS_INLINE AtkHyperlinkStateFlags = 1 << 0

type GtkWidgetFlags Enum

const (
	GTK_TOPLEVEL GtkWidgetFlags = 1 << (iota + 4)
	GTK_NO_WINDOW
	GTK_REALIZED
	GTK_MAPPED
	GTK_VISIBLE
	GTK_SENSITIVE
	GTK_PARENT_SENSITIVE
	GTK_CAN_FOCUS
	GTK_HAS_FOCUS
	GTK_CAN_DEFAULT
	GTK_HAS_DEFAULT
	GTK_HAS_GRAB
	GTK_RC_STYLE
	GTK_COMPOSITE_CHILD
	GTK_NO_REPARENT
	GTK_APP_PAINTABLE
	GTK_RECEIVES_DEFAULT
	GTK_DOUBLE_BUFFERED
	GTK_NO_SHOW_ALL
)

type GtkBuilderError Enum

const (
	GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION GtkBuilderError = iota
	GTK_BUILDER_ERROR_UNHANDLED_TAG
	GTK_BUILDER_ERROR_MISSING_ATTRIBUTE
	GTK_BUILDER_ERROR_INVALID_ATTRIBUTE
	GTK_BUILDER_ERROR_INVALID_TAG
	GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE
	GTK_BUILDER_ERROR_INVALID_VALUE
	GTK_BUILDER_ERROR_VERSION_MISMATCH
	GTK_BUILDER_ERROR_DUPLICATE_ID
)

type GtkCellRendererMode Enum

const (
	GTK_CELL_RENDERER_MODE_INERT GtkCellRendererMode = iota
	GTK_CELL_RENDERER_MODE_ACTIVATABLE
	GTK_CELL_RENDERER_MODE_EDITABLE
)

type GtkCellRendererAccelMode Enum

const (
	GTK_CELL_RENDERER_ACCEL_MODE_GTK GtkCellRendererAccelMode = iota
	GTK_CELL_RENDERER_ACCEL_MODE_OTHER
)

type GtkIconThemeError Enum

const (
	GTK_ICON_THEME_NOT_FOUND GtkIconThemeError = iota
	GTK_ICON_THEME_FAILED
)

type GtkNotebookTab Enum

const (
	GTK_NOTEBOOK_TAB_FIRST GtkNotebookTab = iota
	GTK_NOTEBOOK_TAB_LAST
)

type GtkRecentManagerError Enum

const (
	GTK_RECENT_MANAGER_ERROR_NOT_FOUND GtkRecentManagerError = iota
	GTK_RECENT_MANAGER_ERROR_INVALID_URI
	GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING
	GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED
	GTK_RECENT_MANAGER_ERROR_READ
	GTK_RECENT_MANAGER_ERROR_WRITE
	GTK_RECENT_MANAGER_ERROR_UNKNOWN
)

type GtkRecentChooserError Enum

const (
	GTK_RECENT_CHOOSER_ERROR_NOT_FOUND GtkRecentChooserError = iota
	GTK_RECENT_CHOOSER_ERROR_INVALID_URI
)

type GtkTextBufferTargetInfo Enum

const (
	GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS GtkTextBufferTargetInfo = -(iota + 1)
	GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT
	GTK_TEXT_BUFFER_TARGET_INFO_TEXT
)

type GtkToolbarSpaceStyle Enum

const (
	GTK_TOOLBAR_SPACE_EMPTY GtkToolbarSpaceStyle = iota
	GTK_TOOLBAR_SPACE_LINE
)

type GtkButtonAction Enum

const (
	GTK_BUTTON_SELECTS GtkButtonAction = 1 << iota
	GTK_BUTTON_DRAGS
	GTK_BUTTON_EXPANDS
	GTK_BUTTON_IGNORED GtkButtonAction = 0
)

type GtkCTreePos Enum

const (
	GTK_CTREE_POS_BEFORE GtkCTreePos = iota
	GTK_CTREE_POS_AS_CHILD
	GTK_CTREE_POS_AFTER
)

type GtkCTreeExpansionType Enum

const (
	GTK_CTREE_EXPANSION_EXPAND GtkCTreeExpansionType = iota
	GTK_CTREE_EXPANSION_EXPAND_RECURSIVE
	GTK_CTREE_EXPANSION_COLLAPSE
	GTK_CTREE_EXPANSION_COLLAPSE_RECURSIVE
	GTK_CTREE_EXPANSION_TOGGLE
	GTK_CTREE_EXPANSION_TOGGLE_RECURSIVE
)

type GtkCListFlags Enum

const (
	GTK_CLIST_IN_DRAG GtkCListFlags = 1 << iota
	GTK_CLIST_ROW_HEIGHT_SET
	GTK_CLIST_SHOW_TITLES
	_
	GTK_CLIST_ADD_MODE
	GTK_CLIST_AUTO_SORT
	GTK_CLIST_AUTO_RESIZE_BLOCKED
	GTK_CLIST_REORDERABLE
	GTK_CLIST_USE_DRAG_ICONS
	GTK_CLIST_DRAW_DRAG_LINE
	GTK_CLIST_DRAW_DRAG_RECT
)

const (
	GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID  = -1
	GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID = -2
)

type GailOffsetType Enum

const (
	GAIL_BEFORE_OFFSET GailOffsetType = iota
	GAIL_AT_OFFSET
	GAIL_AFTER_OFFSET
)

type PangoOTTableType Enum

const (
	PANGO_OT_TABLE_GSUB PangoOTTableType = iota
	PANGO_OT_TABLE_GPOS
)

type FT_Size_Request_Type Enum

const (
	FT_SIZE_REQUEST_TYPE_NOMINAL FT_Size_Request_Type = iota
	FT_SIZE_REQUEST_TYPE_REAL_DIM
	FT_SIZE_REQUEST_TYPE_BBOX
	FT_SIZE_REQUEST_TYPE_CELL
	FT_SIZE_REQUEST_TYPE_SCALES
	FT_SIZE_REQUEST_TYPE_MAX
)

type FT_Encoding uint32 // TODO(t): enum/Size?

const (
	FT_ENCODING_NONE           FT_Encoding = 0
	FT_ENCODING_MS_SYMBOL      FT_Encoding = (('s' << 24) | ('y' << 16) | ('m' << 8) | 'b')
	FT_ENCODING_UNICODE        FT_Encoding = (('u' << 24) | ('n' << 16) | ('i' << 8) | 'c')
	FT_ENCODING_SJIS           FT_Encoding = (('s' << 24) | ('j' << 16) | ('i' << 8) | 's')
	FT_ENCODING_GB2312         FT_Encoding = (('g' << 24) | ('b' << 16) | (' ' << 8) | ' ')
	FT_ENCODING_BIG5           FT_Encoding = (('b' << 24) | ('i' << 16) | ('g' << 8) | '5')
	FT_ENCODING_WANSUNG        FT_Encoding = (('w' << 24) | ('a' << 16) | ('n' << 8) | 's')
	FT_ENCODING_JOHAB          FT_Encoding = (('j' << 24) | ('o' << 16) | ('h' << 8) | 'a')
	FT_ENCODING_MS_SJIS                    = FT_ENCODING_SJIS
	FT_ENCODING_MS_GB2312                  = FT_ENCODING_GB2312
	FT_ENCODING_MS_BIG5                    = FT_ENCODING_BIG5
	FT_ENCODING_MS_WANSUNG                 = FT_ENCODING_WANSUNG
	FT_ENCODING_MS_JOHAB                   = FT_ENCODING_JOHAB
	FT_ENCODING_ADOBE_STANDARD FT_Encoding = (('A' << 24) | ('D' << 16) | ('O' << 8) | 'B')
	FT_ENCODING_ADOBE_EXPERT   FT_Encoding = (('A' << 24) | ('D' << 16) | ('B' << 8) | 'E')
	FT_ENCODING_ADOBE_CUSTOM   FT_Encoding = (('A' << 24) | ('D' << 16) | ('B' << 8) | 'C')
	FT_ENCODING_ADOBE_LATIN_1  FT_Encoding = (('l' << 24) | ('a' << 16) | ('t' << 8) | '1')
	FT_ENCODING_OLD_LATIN_2    FT_Encoding = (('l' << 24) | ('a' << 16) | ('t' << 8) | '2')
	FT_ENCODING_APPLE_ROMAN    FT_Encoding = (('a' << 24) | ('r' << 16) | ('m' << 8) | 'n')
)

type FT_Glyph_Format uint32 // TODO(t): enum/Size?

const (
	FT_GLYPH_FORMAT_NONE      FT_Glyph_Format = ((0 << 24) | (0 << 16) | (0 << 8) | 0)
	FT_GLYPH_FORMAT_COMPOSITE FT_Glyph_Format = (('c' << 24) | ('o' << 16) | ('m' << 8) | 'p')
	FT_GLYPH_FORMAT_BITMAP    FT_Glyph_Format = (('b' << 24) | ('i' << 16) | ('t' << 8) | 's')
	FT_GLYPH_FORMAT_OUTLINE   FT_Glyph_Format = (('o' << 24) | ('u' << 16) | ('t' << 8) | 'l')
	FT_GLYPH_FORMAT_PLOTTER   FT_Glyph_Format = (('p' << 24) | ('l' << 16) | ('o' << 8) | 't')
)

type FT_Render_Mode Enum

const (
	FT_RENDER_MODE_NORMAL FT_Render_Mode = iota
	FT_RENDER_MODE_LIGHT
	FT_RENDER_MODE_MONO
	FT_RENDER_MODE_LCD
	FT_RENDER_MODE_LCD_V
	FT_RENDER_MODE_MAX
)

type FT_LcdFilter Enum

const (
	FT_LCD_FILTER_DEFAULT FT_LcdFilter = 1 << iota
	FT_LCD_FILTER_LIGHT
	_
	_
	FT_LCD_FILTER_LEGACY
	FT_LCD_FILTER_MAX               = FT_LCD_FILTER_LEGACY + 1
	FT_LCD_FILTER_NONE FT_LcdFilter = 0
)

type FT_TrueTypeEngineType Enum

const (
	FT_TRUETYPE_ENGINE_TYPE_NONE FT_TrueTypeEngineType = iota
	FT_TRUETYPE_ENGINE_TYPE_UNPATENTED
	FT_TRUETYPE_ENGINE_TYPE_PATENTED
)

type FT_Orientation Enum

const (
	FT_ORIENTATION_TRUETYPE FT_Orientation = iota
	FT_ORIENTATION_POSTSCRIPT
	FT_ORIENTATION_NONE
	FT_ORIENTATION_FILL_RIGHT = FT_ORIENTATION_TRUETYPE
	FT_ORIENTATION_FILL_LEFT  = FT_ORIENTATION_POSTSCRIPT
)

type FT_Sfnt_Tag Enum

const (
	Ft_sfnt_head FT_Sfnt_Tag = iota
	Ft_sfnt_maxp
	Ft_sfnt_os2
	Ft_sfnt_hhea
	Ft_sfnt_vhea
	Ft_sfnt_post
	Ft_sfnt_pclt
	Sfnt_max
)

type BDF_PropertyType Enum

const (
	BDF_PROPERTY_TYPE_NONE BDF_PropertyType = iota
	BDF_PROPERTY_TYPE_ATOM
	BDF_PROPERTY_TYPE_INTEGER
	BDF_PROPERTY_TYPE_CARDINAL
)

type FT_StrokerBorder Enum

const (
	FT_STROKER_BORDER_LEFT FT_StrokerBorder = iota
	FT_STROKER_BORDER_RIGHT
)

type FT_Stroker_LineCap Enum

const (
	FT_STROKER_LINECAP_BUTT FT_Stroker_LineCap = iota
	FT_STROKER_LINECAP_ROUND
	FT_STROKER_LINECAP_SQUARE
)

type FT_Stroker_LineJoin Enum

const (
	FT_STROKER_LINEJOIN_ROUND FT_Stroker_LineJoin = iota
	FT_STROKER_LINEJOIN_BEVEL
	FT_STROKER_LINEJOIN_MITER
)

type (
	GChildWatchFunc func(pid GPid, status int, data Gpointer)

	GCallback func()

	GdkInputFunction func(data Gpointer,
		source int, condition GdkInputCondition)

	GCompareDataFunc func(
		a, b Gconstpointer, user_data Gpointer) int

	GCompareFunc func(a, b Gconstpointer) int

	GCompletionFunc func(Gpointer) string

	GCompletionStrncmpFunc func(s1, s2 string, n Gsize) int

	GCopyFunc func(src Gconstpointer, data Gpointer) Gpointer

	GDataForeachFunc func(
		key_id GQuark, data Gpointer, user_data Gpointer)

	GDestroyNotify func(data Gpointer)

	GEqualFunc func(a, b Gconstpointer) Gboolean

	GFunc func(data Gpointer, user_data Gpointer)

	GHashFunc func(key Gconstpointer) uint

	GHFunc func(key, value, user_data Gpointer)

	GHookCheckMarshaller func(
		hook *GHook, marshal_data Gpointer) Gboolean

	GHookCompareFunc func(new_hook, sibling *GHook) int

	GHookFinalizeFunc func(hook_list *GHookList, hook *GHook)

	GHookFindFunc func(hook *GHook, data Gpointer) Gboolean

	GHookMarshaller func(hook *GHook, marshal_data Gpointer)

	GHRFunc func(key, value, user_data Gpointer) Gboolean

	GIOFunc func(source *GIOChannel,
		condition GIOCondition, data Gpointer) Gboolean

	GLogFunc func(
		log_domain string,
		log_level GLogLevelFlags,
		message string,
		user_data Gpointer)

	GNodeForeachFunc func(node *GNode, data Gpointer)

	GNodeTraverseFunc func(node *GNode, data Gpointer) Gboolean

	GOptionErrorFunc func(
		context *GOptionContext,
		group *GOptionGroup,
		data Gpointer,
		err **GError)

	GOptionParseFunc func(
		context *GOptionContext,
		group *GOptionGroup,
		data Gpointer,
		err **GError) Gboolean

	GPollFunc func(
		ufds *GPollFD, nfsd uint, timeout int) int

	GPrintFunc func(str string)

	GRegexEvalCallback func(match_info *GMatchInfo,
		result *GString, user_data Gpointer) Gboolean

	GScannerMsgFunc func(
		scanner *GScanner, message string, err Gboolean)

	GSequenceIterCompareFunc func(
		a, b *GSequenceIter, data Gpointer) int

	GSpawnChildSetupFunc func(user_data Gpointer)

	GTestDataFunc func(user_data Gconstpointer)

	GTestFixtureFunc func(
		fixture Gpointer, user_data Gconstpointer)

	GTestFunc func()

	GTestLogFatalFunc func(
		log_domain string, log_level GLogLevelFlags,
		message string, user_data Gpointer) Gboolean

	GThreadFunc func(data Gpointer) Gpointer

	GTranslateFunc func(str string, data Gpointer) string

	GTraverseFunc func(key, value, data Gpointer) Gboolean

	GClosureMarshal func() //REMOVE

	GBaseInitFunc func(
		g_class Gpointer)

	GInstanceInitFunc func(
		instance *GTypeInstance,
		g_class Gpointer)

	GtkWindowKeysForeachFunc func(
		window *GtkWindow,
		keyval uint,
		modifiers GdkModifierType,
		is_mnemonic Gboolean,
		data Gpointer)

	GAsyncReadyCallback func(
		source_object *GObject,
		res *GAsyncResult,
		user_data Gpointer)

	GdkFilterFunc func(
		xevent *GdkXEvent,
		event *GdkEvent,
		data Gpointer) GdkFilterReturn

	GdkSpanFunc func(
		span *GdkSpan,
		data Gpointer)

	PangoCairoShapeRendererFunc func(
		cr *Cairo,
		attr *PangoAttrShape,
		do_path Gboolean,
		data Gpointer)

	PangoAttrDataCopyFunc func(data Gconstpointer) Gpointer

	PangoAttrFilterFunc func(
		attribute *PangoAttribute,
		data Gpointer) Gboolean

	PangoFontsetForeachFunc func(
		fontset *PangoFontset,
		font *PangoFont,
		data Gpointer) Gboolean

	CairoDestroyFunc func(data *Void)

	CairoUserScaledFontInitFunc func(
		scaled_font *CairoScaledFont,
		cr *Cairo,
		extents *CairoFontExtents) CairoStatus

	CairoUserScaledFontRenderGlyphFunc func(
		scaled_font *CairoScaledFont,
		glyph UnsignedLong,
		cr *Cairo,
		extents *CairoTextExtents) CairoStatus

	CairoUserScaledFontTextToGlyphsFunc func(
		scaled_font *CairoScaledFont,
		utf8 string,
		utf8_len int,
		glyphs **CairoGlyph,
		num_glyphs *int,
		clusters **CairoTextCluster,
		num_clusters *int,
		cluster_flags *CairoTextClusterFlags) CairoStatus

	CairoUserScaledFontUnicodeToGlyphFunc func(
		scaled_font *CairoScaledFont,
		unicode UnsignedLong,
		glyph_index *UnsignedLong) CairoStatus

	CairoWriteFunc func(
		closure *Void,
		data *UnsignedChar,
		length uint) CairoStatus

	CairoReadFunc func(
		closure *Void,
		data *UnsignedChar,
		length uint) CairoStatus

	CsiSurfaceCreateFunc func(
		closure *Void,
		content CairoContent,
		width, height float64,
		uid Long) *CairoSurface

	CsiDestroyFunc func(closure, ptr *Void)

	Csi_context_create_func_t func(
		closure *Void,
		surface *CairoSurface) *Cairo

	Csi_show_page_func_t func(closure *Void, cr *Cairo)

	Csi_copy_page_func_t func(closure *Void, cr *Cairo)

	AtkEventListener func(obj *AtkObject)

	GSignalEmissionHook func(
		ihint *GSignalInvocationHint,
		n_param_values uint,
		param_values *GValue,
		data Gpointer) Gboolean

	AtkKeySnoopFunc func(
		event *AtkKeyEventStruct, func_data Gpointer) int

	AtkPropertyChangeHandler func(*AtkObject, *AtkPropertyValues)

	AtkEventListenerInit func()

	AtkFocusHandler func(*AtkObject, Gboolean)

	PangoFcDecoderFindFunc func(
		pattern *FcPattern,
		user_data Gpointer) *PangoFcDecoder

	PangoFT2SubstituteFunc func(
		pattern *FcPattern,
		data Gpointer)

	FT_Alloc_Func func(
		memory FT_Memory,
		size Long) *Void

	FT_Free_Func func(
		memory FT_Memory,
		block *Void)

	FT_Realloc_Func func(
		memory FT_Memory,
		cur_size, new_size Long,
		block *Void) *Void

	FT_Glyph_InitFunc func(
		glyph FT_Glyph,
		slot FT_GlyphSlot) FT_Error

	FT_Glyph_DoneFunc func(
		glyph FT_Glyph)

	FT_Glyph_TransformFunc func(
		glyph FT_Glyph,
		matrix *FT_Matrix,
		delta *FT_Vector)

	FT_Glyph_GetBBoxFunc func(glyph FT_Glyph, abbox *FT_BBox)

	FT_Glyph_CopyFunc func(source, target FT_Glyph) FT_Error

	FT_Glyph_PrepareFunc func(
		glyph FT_Glyph, slot FT_GlyphSlot) FT_Error

	FT_Module_Constructor func(module FT_Module) FT_Error

	FT_Module_Destructor func(module FT_Module)

	FT_Module_Requester func(
		module FT_Module, name string) FT_Module_Interface

	FT_List_Iterator func(node FT_ListNode, user *Void) FT_Error

	FT_List_Destructor func(memory FT_Memory, data, user *Void)

	FT_Outline_MoveToFunc func(to *FT_Vector, user *Void) int

	FT_Outline_LineToFunc func(to *FT_Vector, user *Void) int

	FT_Outline_ConicToFunc func(
		control, to *FT_Vector, user *Void) int

	FT_Outline_CubicToFunc func(
		control1, control2 *FT_Vector,
		to *FT_Vector, user *Void) int

	FT_SpanFunc func(
		y, count int, spans *FT_Span, user *Void)

	FT_Raster_BitTest_Func func(y, x int, user *Void) int

	FT_Raster_BitSet_Func func(y, x int, user *Void)

	FT_DebugHook_Func func(arg *Void)

	FTC_Face_Requester func(
		face_id FTC_FaceID,
		library FT_Library,
		request_data FT_Pointer,
		aface *FTFace) FT_Error

	GBoxedCopyFunc func(boxed Gpointer) Gpointer

	GBoxedFreeFunc func(boxed Gpointer)

	GWeakNotify func(
		data Gpointer, where_the_object_was *GObject)

	GToggleNotify func(
		data Gpointer, object *GObject, is_last_ref Gboolean)

	GSignalAccumulator func(
		ihint *GSignalInvocationHint,
		return_accu *GValue,
		handler_return *GValue,
		data Gpointer) Gboolean

	GBusAcquiredCallback func(
		connection *GDBusConnection,
		name string,
		user_data Gpointer)

	GBusNameAcquiredCallback func(
		connection *GDBusConnection,
		name string,
		user_data Gpointer)

	GBusNameLostCallback func(
		connection *GDBusConnection,
		name string,
		user_data Gpointer)

	GBusNameAppearedCallback func(
		connection *GDBusConnection,
		name string,
		name_owner string,
		user_data Gpointer)

	GBusNameVanishedCallback func(
		connection *GDBusConnection,
		name string,
		user_data Gpointer)

	GDBusInterfaceMethodCallFunc func(
		connection *GDBusConnection,
		sender string,
		object_path string,
		interface_name string,
		method_name string,
		parameters *GVariant,
		invocation *GDBusMethodInvocation,
		user_data Gpointer)

	GDBusInterfaceGetPropertyFunc func(
		connection *GDBusConnection,
		sender string,
		object_path string,
		interface_name string,
		property_name string,
		err **GError,
		user_data Gpointer)

	GDBusInterfaceSetPropertyFunc func(
		connection *GDBusConnection,
		sender string,
		object_path string,
		interface_name string,
		property_name string,
		value *GVariant,
		err **GError,
		user_data Gpointer)

	GDBusSubtreeEnumerateFunc func(
		connection *GDBusConnection,
		sender string,
		object_path string,
		user_data Gpointer) **Gchar

	GDBusSubtreeIntrospectFunc func(
		connection *GDBusConnection,
		sender string,
		object_path string,
		node string,
		user_data Gpointer) **GDBusInterfaceInfo

	GDBusSubtreeDispatchFunc func(
		connection *GDBusConnection,
		sender string,
		object_path string,
		interface_name string,
		node string,
		out_user_data *Gpointer,
		user_data Gpointer) *GDBusInterfaceVTable

	GDBusSignalCallback func(
		connection *GDBusConnection,
		sender_name string,
		object_path string,
		interface_name string,
		signal_name string,
		parameters *GVariant,
		user_data Gpointer)

	GDBusMessageFilterFunction func(
		connection *GDBusConnection,
		message *GDBusMessage,
		incoming Gboolean,
		user_data Gpointer) *GDBusMessage

	GFileProgressCallback func(
		current_num_bytes Goffset,
		total_num_bytes Goffset,
		user_data Gpointer)

	GIOSchedulerJobFunc func(
		job *GIOSchedulerJob,
		cancellable *GCancellable,
		user_data Gpointer) Gboolean

	GFileReadMoreCallback func(
		file_contents string,
		file_size Goffset,
		callback_data Gpointer) Gboolean

	GReallocFunc func(
		data Gpointer,
		size Gsize) Gpointer

	GSettingsBindSetMapping func(
		value *GValue,
		expected_type *GVariantType,
		user_data Gpointer) *GVariant

	GSettingsBindGetMapping func(
		value *GValue,
		variant *GVariant,
		user_data Gpointer) Gboolean

	GSettingsGetMapping func(
		value *GVariant,
		result *Gpointer,
		user_data Gpointer) Gboolean

	GSimpleAsyncThreadFunc func(
		res *GSimpleAsyncResult,
		object *GObject,
		cancellable *GCancellable)

	GBaseFinalizeFunc func(g_class Gpointer)

	GClassInitFunc func(g_class, class_data Gpointer)

	GClassFinalizeFunc func(g_class, class_data Gpointer)

	GInterfaceFinalizeFunc func(g_iface, iface_data Gpointer)

	GInterfaceInitFunc func(g_iface, iface_data Gpointer)

	GTypeClassCacheFunc func(
		cache_data Gpointer, g_class *GTypeClass) Gboolean

	GTypeInterfaceCheckFunc func(check_data, g_iface Gpointer)

	G_thread_gettime func() uint64
)
