package types

//TODO(t): add constants for [...]s

type (
	Tm struct {
		Sec, Min, Hour, Mday, Mon, Year, Wday, Yday, Isdst int
	}

	GArray struct {
		Data *Gchar
		Leng uint
	}

	GByteArray struct {
		Data *uint8
		Leng uint
	}

	GCompletion struct {
		Items       *GList
		Fnc         GCompletionFunc
		Prefix      *Gchar
		Cache       *GList
		StrncmpFunc GCompletionStrncmpFunc
	}

	GDate struct {
		Bits1, Bits2 uint
		// JulianDays : 32
		// Julian : 1
		// Dmy : 1
		// Day : 6
		// Month : 4
		// Year : 16
	}

	GDebugKey struct {
		Key   *Gchar
		Value uint
	}

	GError struct {
		Domain  GQuark
		Code    int
		Message *Gchar
	}

	GHashTableIter struct {
		Dummy1 Gpointer
		Dummy2 Gpointer
		Dummy3 Gpointer
		Dummy4 int
		Dummy5 Gboolean
		Dummy6 Gpointer
	}

	GHook struct {
		Data     Gpointer
		Next     *GHook
		Prev     *GHook
		RefCount uint
		HookId   Gulong
		Flags    uint
		Fnc      Gpointer
		Destroy  GDestroyNotify
	}

	GHookList struct {
		SeqId Gulong
		Bits  uint
		// HookSize : 16
		// IsSetup : 1
		Hooks        *GHook
		Dummy3       Gpointer
		FinalizeHook GHookFinalizeFunc
		Dummy        [2]Gpointer
	}

	GIOChannel struct {
		RefCount        int
		Funcs           *GIOFuncs
		Encoding        *Gchar
		ReadCd          GIConv
		WriteCd         GIConv
		LineTerm        *Gchar
		LineTermLen     uint
		BufSize         Gsize
		ReadBuf         *GString
		EncodedReadBuf  *GString
		WriteBuf        *GString
		PartialWriteBuf [6]Gchar
		Bits            uint
		// UseBuffer : 1
		// DoEncode : 1
		// CloseOnUnref : 1
		// IsReadable : 1
		// IsWriteable : 1
		// IsSeekable : 1
		_, _ Gpointer
	}

	GList struct {
		Data Gpointer
		Next *GList
		Prev *GList
	}

	GIOFuncs struct {
		IoRead func(
			channel *GIOChannel,
			buf *Gchar,
			count Gsize,
			bytesRead *Gsize,
			err **GError) GIOStatus
		IoWrite func(
			channel *GIOChannel,
			buf *Gchar,
			count Gsize,
			bytesWritten *Gsize,
			err **GError) GIOStatus
		IoSeek func(
			channel *GIOChannel,
			offset int64,
			type_ GSeekType,
			err **GError) GIOStatus
		IoClose func(
			channel *GIOChannel,
			err **GError) GIOStatus
		IoCreateWatch func(
			channel *GIOChannel,
			condition GIOCondition) *GSource
		IoFree func(
			channel *GIOChannel)
		IoSetFlags func(
			channel *GIOChannel,
			flags GIOFlags,
			err **GError) GIOStatus
		IoGetFlags func(
			channel *GIOChannel) GIOFlags
	}

	GMarkupParser struct {
		StartElement func(
			context *GMarkupParseContext,
			elementName *Gchar,
			attributeNames **Gchar,
			attributeValues **Gchar,
			userData Gpointer,
			error **GError)
		EndElement func(
			context *GMarkupParseContext,
			elementName *Gchar,
			userData Gpointer,
			error **GError)
		Text func(
			context *GMarkupParseContext,
			text *Gchar,
			textLen Gsize,
			userData Gpointer,
			error **GError)
		Passthrough func(
			context *GMarkupParseContext,
			passthroughText *Gchar,
			textLen Gsize,
			userData Gpointer,
			error **GError)
		Error func(
			context *GMarkupParseContext,
			error *GError,
			userData Gpointer)
	}

	GNode struct {
		Data     Gpointer
		Next     *GNode
		Prev     *GNode
		Parent   *GNode
		Children *GNode
	}

	GOnce struct {
		Status GOnceStatus
		Retval Gpointer
	}

	GOptionEntry struct {
		LongName       *Gchar
		ShortName      Gchar
		Flags          int
		Arg            GOptionArg
		ArgData        Gpointer
		Description    *Gchar
		ArgDescription *Gchar
	}

	GPollFD struct {
		Fd      int
		Events  uint16
		Revents uint16
	}

	GPtrArray struct {
		Pdata *Gpointer
		Len   uint
	}

	GQueue struct {
		Head   *GList
		Tail   *GList
		Length uint
	}

	GScanner struct {
		UserData       Gpointer
		MaxParseErrors uint
		ParseErrors    uint
		InputName      *Gchar
		Qdata          *GData
		Config         *GScannerConfig
		Token          GTokenType
		Value          GTokenValue
		Line           uint
		Position       uint
		NextToken      GTokenType
		NextValue      GTokenValue
		NextLine       uint
		NextPosition   uint
		SymbolTable    *GHashTable
		InputFd        int
		Text           *Gchar
		TextEnd        *Gchar
		Buffer         *Gchar
		ScopeId        uint
		MsgHandler     GScannerMsgFunc
	}

	GScannerConfig struct {
		CsetSkipCharacters  *Gchar
		CsetIdentifierFirst *Gchar
		CsetIdentifierNth   *Gchar
		CpairCommentSingle  *Gchar
		Bits                uint
		// CaseSensitive : 1
		// SkipCommentMulti : 1
		// SkipCommentSingle : 1
		// ScanCommentMulti : 1
		// ScanIdentifier : 1
		// ScanIdentifier1char : 1
		// ScanIdentifier_NULL : 1
		// ScanSymbols : 1
		// ScanBinary : 1
		// ScanOctal : 1
		// ScanFloat : 1
		// ScanHex : 1
		// ScanHexDollar : 1
		// ScanStringSq : 1
		// ScanStringDq : 1
		// Numbers2Int : 1
		// Int2Float : 1
		// Identifier2String : 1
		// Char2Token : 1
		// Symbol2Token : 1
		// Scope0Fallback : 1
		// StoreInt64 : 1
		PaddingDummy uint
	}

	GSList struct {
		Data Gpointer
		Next *GSList
	}

	GSource struct{} //REMOVE

	GStaticPrivate struct {
		Index uint
	}

	GStaticRecMutex struct {
		Mutex GStaticMutex
		Depth uint
		Owner GSystemThread
	}

	GSystemThread struct {
		// Union
		// Data[4] Char
		DummyDouble float64
		// DummyPointer *Void
		// DummyLong Long
	}

	GStaticRWLock struct {
		Mutex       GStaticMutex
		ReadCond    *GCond
		WriteCond   *GCond
		ReadCounter uint
		HaveWriter  Gboolean
		WantToRead  uint
		WantToWrite uint
	}

	GString struct {
		Str          *Gchar
		Len          Gsize
		AllocatedLen Gsize
	}

	GTestLogBuffer struct {
		Data *GString
		Msgs *GSList
	}

	GTestLogMsg struct {
		LogType  GTestLogType
		NStrings uint
		Strings  **Gchar
		NNums    uint
		Nums     *LongDouble
	}

	GThread struct {
		Func     GThreadFunc
		Data     Gpointer
		Joinable Gboolean
		Priority GThreadPriority
	}

	GThreadFunctions struct {
		MutexNew      func()
		MutexLock     func(mutex *GMutex) *GMutex
		MutexTrylock  func(mutex *GMutex) Gboolean
		MutexUnlock   func(mutex *GMutex)
		MutexFree     func(mutex *GMutex)
		CondNew       func() *GCond
		CondSignal    func(cond *GCond)
		CondBroadcast func(cond *GCond)
		CondWait      func(cond *GCond, mutex *GMutex)
		CondTimedWait func(
			cond *GCond,
			mutex *GMutex,
			endTime *GTimeVal) Gboolean
		CondFree     func(cond *GCond)
		PrivateNew   func(destructor GDestroyNotify) *GPrivate
		PrivateGet   func(privateKey *GPrivate) Gpointer
		PrivateSet   func(privateKey *GPrivate, data Gpointer)
		ThreadCreate func(
			fnc GThreadFunc,
			data Gpointer,
			stackSize Gulong,
			joinable Gboolean,
			bound Gboolean,
			priority GThreadPriority,
			thread Gpointer,
			error **GError)
		Threadield        func()
		ThreadJoin        func(thread Gpointer)
		ThreadExit        func()
		ThreadSetPriority func(
			thread Gpointer,
			priority GThreadPriority)
		ThreadSelf  func(thread Gpointer)
		ThreadEqual func(
			thread1 Gpointer,
			thread2 Gpointer) Gboolean
	}

	GThreadPool struct {
		Func      GFunc
		UserData  Gpointer
		Exclusive Gboolean
	}

	GTimeVal struct {
		TvSec  Glong
		TvUsec Glong
	}

	GTokenValue struct {
		// Union
		// VSymbol  Gpointer;
		// VIdentifier  *Gchar;
		// VBinary  Gulong;
		// VOctal  Gulong;
		// VInt  Gulong;
		VInt64 uint64
		// VFloat  float64;
		// VHex  Gulong;
		// VString  *Gchar;
		// VComment  *Gchar;
		// VChar  Guchar;
		// VError  uint;
	}

	GTuples struct {
		Len uint
	}

	GVariantBuilder struct {
		X [16]Gsize
	}

	GVariantIter struct {
		X [16]Gsize
	}

	GTypeClass struct{} //REMOVE

	GTypeInstance struct{} //REMOVE

	GObject struct{} //REMOVE

	GtkWindow struct{} //REMOVE

	GdkScreen struct{} //REMOVE

	GValue struct{} //REMOVE

	GdkKeyboardGrabInfo struct {
		Window       *GdkWindow
		NativeWindow *GdkWindow
		Serial       Gulong
		OwnerEvents  Gboolean
		Time         GUint32
	}

	GdkPointerWindowInfo struct {
		ToplevelUnderPointer *GdkWindow
		WindowUnderPointer   *GdkWindow
		ToplevelX, Toplevel  float64
		State                GUint32
		Button               GUint32
		MotionHintSerial     Gulong
	}

	GdkEvent struct{} //REMOTE

	GSignalInvocationHint struct {
		SignalId uint
		Detail   GQuark
		RunType  GSignalFlags
	}

	GdkEventOwnerChange struct {
		Type          GdkEventType
		Window        *GdkWindow
		SendEvent     int8
		Owner         GdkNativeWindow
		Reason        GdkOwnerChange
		Selection     GdkAtom
		Time          GUint32
		SelectionTime GUint32
	}

	GdkEventSetting struct {
		Type      GdkEventType
		Window    *GdkWindow
		SendEvent int8
		Action    GdkSettingAction
		Name      *Char
	}

	GParamSpec struct{} //REMOVE

	GObjectConstructParam struct {
		Pspec *GParamSpec
		Value *GValue
	}

	GtkRcProperty struct {
		TypeName     GQuark
		PropertyName GQuark
		Origin       *Gchar
		Value        GValue
	}

	GtkWidgetAuxInfo struct {
		X      int
		Y      int
		Width  int
		Height int
		Bits   uint
		// XSet : 1
		// YSet : 1
	}

	GdkPointerHooks struct {
		GetPointer func(
			window *GdkWindow,
			x *int,
			y *int,
			mask *GdkModifierType) *GdkWindow
		WindowAtPointer func(screen *GdkScreen,
			winX *int,
			win *int) *GdkWindow
	}

	GdkRgbCmap struct {
		Colors   [256]GUint32
		NColors  int
		InfoList *GSList
	}

	GdkSegment struct {
		X1 int
		Y1 int
		X2 int
		Y2 int
	}

	GdkSpan struct {
		X     int
		Y     int
		Width int
	}

	GdkTimeCoord struct {
		Time GUint32
		Axes [128]float64
	}

	GdkTrapezoid struct {
		Y1, X11, X21, Y2, X12, X22 float64
	}

	GMemVTable struct {
		Malloc     func(nBytes Gsize) Gpointer
		Realloc    func(mem Gpointer, nBytes Gsize) Gpointer
		Free       func(mem Gpointer)
		Calloc     func(nBlocks Gsize, nBlockBytes Gsize) Gpointer
		TryMalloc  func(nBytes Gsize) Gpointer
		TryRealloc func(mem Gpointer, nBytes Gsize) Gpointer
	}

	GTrashStack struct {
		Next *GTrashStack
	}

	GParameter struct {
		Name  string
		Value GValue
	}

	GSignalQuery struct {
		SignalId    uint
		SignalName  string
		Itype       GType
		SignalFlags GSignalFlags
		ReturnType  GType
		NParams     uint
		ParamTypes  *GType
	}

	GTypeQuery struct {
		Type         GType
		TypeName     string
		ClassSize    uint
		InstanceSize uint
	}

	GOutputVector struct {
		Buffer Gconstpointer
		Size   Gsize
	}

	GTypeValueTable struct {
		ValueInit func(value *GValue)
		ValueFree func(value *GValue)
		ValueCopy func(
			srcValue *GValue, destValue *GValue)
		ValuePeekPointer func(value *GValue) Gpointer
		CollectFormat    *Gchar
		CollectValue     func(
			value *GValue,
			nCollectValues uint,
			collectValues *GTypeCValue,
			collectFlags uint) *Gchar
		LcopyFormat *Gchar
		LcopyValue  func(
			value *GValue,
			nCollectValues uint,
			collectValues *GTypeCValue,
			collectFlags uint) *Gchar
	}

	GTypeFundamentalInfo struct {
		TypeFlags GTypeFundamentalFlags
	}

	GInterfaceInfo struct {
		InterfaceInit     GInterfaceInitFunc
		InterfaceFinalize GInterfaceFinalizeFunc
		InterfaceData     Gpointer
	}

	GTestConfig struct {
		TestInitialized Gboolean
		TestQuick       Gboolean
		TestPerf        Gboolean
		TestVerbose     Gboolean
		TestQuiet       Gboolean
	}
)
