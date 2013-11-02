// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package types provides API structs, enums and types for Gtk+2.
package types

type (
	fix uintptr

	Long_double fix
	LongDouble  fix
	VaList      fix
	FILE        fix
	Stat        fix

	Char          int8
	Enum          int
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

	Gboolean               int
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
	GTime                  GInt32
	GTimeSpan              int64
	GType                  Gsize // REMOVE
	Gulong                 UnsignedLong
	Gunichar2              uint16
	GVoidFunc              func()

	GdkDrawable struct{ parent GObject } //REMOVE

	GCancellablePrivate      struct{}
	GData                    struct{}
	GdkXEvent                struct{}
	GFileInputStreamPrivate  struct{}
	GInputStreamPrivate      struct{}
	GMatchInfo               struct{}
	GResolverPrivate         struct{}
	GSocketConnectionPrivate struct{}
	GSocketServicePrivate    struct{}
	GtkClipboard             struct{}
	GtkLabelSelectionInfo    struct{}
	GUnixFDList              struct{}
	Utimbuf                  struct{}
	Void                     struct{}
	Dummy                    *struct{}
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

type GConvertError Enum

const (
	G_CONVERT_ERROR_NO_CONVERSION GConvertError = iota
	G_CONVERT_ERROR_ILLEGAL_SEQUENCE
	G_CONVERT_ERROR_FAILED
	G_CONVERT_ERROR_PARTIAL_INPUT
	G_CONVERT_ERROR_BAD_URI
	G_CONVERT_ERROR_NOT_ABSOLUTE_PATH
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

// type GSeekType Enum // REMOVE

type GKeyFileError Enum

const (
	G_KEY_FILE_ERROR_UNKNOWN_ENCODING GKeyFileError = iota
	G_KEY_FILE_ERROR_PARSE
	G_KEY_FILE_ERROR_NOT_FOUND
	G_KEY_FILE_ERROR_KEY_NOT_FOUND
	G_KEY_FILE_ERROR_GROUP_NOT_FOUND
	G_KEY_FILE_ERROR_INVALID_VALUE
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

type GFilesystemPreviewType Enum

const (
	G_FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS GFilesystemPreviewType = iota
	G_FILESYSTEM_PREVIEW_TYPE_IF_LOCAL
	G_FILESYSTEM_PREVIEW_TYPE_NEVER
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

type GResolverError Enum

const (
	G_RESOLVER_ERROR_NOT_FOUND GResolverError = iota
	G_RESOLVER_ERROR_TEMPORARY_FAILURE
	G_RESOLVER_ERROR_INTERNAL
)

type GSocketMsgFlags Enum

const (
	G_SOCKET_MSG_OOB GSocketMsgFlags = 1 << iota
	G_SOCKET_MSG_PEEK
	G_SOCKET_MSG_DONTROUTE
	G_SOCKET_MSG_NONE GSocketMsgFlags = 0
)

type GUnixSocketAddressType Enum

const (
	G_UNIX_SOCKET_ADDRESS_INVALID GUnixSocketAddressType = iota
	G_UNIX_SOCKET_ADDRESS_ANONYMOUS
	G_UNIX_SOCKET_ADDRESS_PATH
	G_UNIX_SOCKET_ADDRESS_ABSTRACT
	G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED
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

type GTlsAuthenticationMode Enum

const (
	G_TLS_AUTHENTICATION_NONE GTlsAuthenticationMode = iota
	G_TLS_AUTHENTICATION_REQUESTED
	G_TLS_AUTHENTICATION_REQUIRED
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

type (
	GChildWatchFunc func(pid GPid, status int, data Gpointer)

	GdkInputFunction func(data Gpointer,
		source int, condition GdkInputCondition)

	GCompareDataFunc func(
		a, b Gconstpointer, user_data Gpointer) int

	GCompareFunc func(a, b Gconstpointer) int

	GCopyFunc func(src Gconstpointer, data Gpointer) Gpointer

	GDataForeachFunc func(
		keyId Quark, data Gpointer, userData Gpointer)

	GEqualFunc func(a, b Gconstpointer) Gboolean

	GFunc func(data Gpointer, user_data Gpointer)

	GLogFunc func(
		log_domain string,
		log_level GLogLevelFlags,
		message string,
		user_data Gpointer)

	GNodeForeachFunc func(node *GNode, data Gpointer)

	GNodeTraverseFunc func(node *GNode, data Gpointer) Gboolean

	GPollFunc func(
		ufds *GPollFD, nfsd uint, timeout int) int

	GPrintFunc func(str string)

	GTranslateFunc func(str string, data Gpointer) string

	GClosureMarshal func() //REMOVE

	GBaseInitFunc func(
		gClass Gpointer)

	GInstanceInitFunc func(
		instance *GTypeInstance,
		gClass Gpointer)

	GtkWindowKeysForeachFunc func(
		window *GtkWindow,
		keyval uint,
		modifiers GdkModifierType,
		is_mnemonic Gboolean,
		data Gpointer)

	GdkFilterFunc func(
		xevent *GdkXEvent,
		event *GdkEvent,
		data Gpointer) GdkFilterReturn

	GdkSpanFunc func(
		span *GdkSpan,
		data Gpointer)

	GBoxedCopyFunc func(boxed Gpointer) Gpointer

	GBoxedFreeFunc func(boxed Gpointer)

	GWeakNotify func(
		data Gpointer, where_the_object_was *GObject)

	GToggleNotify func(
		data Gpointer, object *GObject, is_last_ref Gboolean)

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
