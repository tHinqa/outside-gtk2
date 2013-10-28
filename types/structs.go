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

	FT_FaceRec struct {
		NumFaces           FT_Long
		FaceIndex          FT_Long
		FaceFlags          FT_Long
		StyleFlags         FT_Long
		NumGlyphs          FT_Long
		FamilyName         *FT_String
		StyleName          *FT_String
		NumFixedSizes      int
		AvailableSizes     *FT_Bitmap_Size
		NumCharmaps        int
		Charmaps           *FT_CharMap
		Generic            FT_Generic
		Bbox               FT_BBox
		UnitsPer_EM        uint16
		Ascender           int16
		Descender          int16
		Height             int16
		MaxAdvanceWidth    int16
		MaxAdvanceHeight   int16
		UnderlinePosition  int16
		UnderlineThickness int16
		Glyph              FT_GlyphSlot
		Size               FT_Size
		Charmap            FT_CharMap
		Driver             FT_Driver
		Memory             FT_Memory
		Stream             FT_Stream
		SizesList          FT_ListRec
		Autohint           FT_Generic
		Extensions         *Void
		Internal           FT_Face_Internal
	}

	FT_Bitmap_Size struct {
		Height int16
		Width  int16
		Size   FT_Pos
		XPpem  FT_Pos
		YPpem  FT_Pos
	}

	FT_Generic_Finalizer func(
		Object *Void)

	FT_Open_Args struct {
		Flags      uint
		MemoryBase *FT_Byte
		MemorySize FT_Long
		Pathname   *FT_String
		Stream     FT_Stream
		Driver     FT_Module
		NumParams  int
		Params     *FT_Parameter
	}

	FT_StreamDesc struct {
		//Union
		Value   Long
		Pointer *Void
	}

	FT_Parameter struct {
		Tag  FT_ULong
		Data FT_Pointer
	}

	FT_Size_RequestRec struct {
		Type           FT_Size_Request_Type
		Width          FT_Long
		Height         FT_Long
		HoriResolution uint
		VertResolution uint
	}

	FT_Matrix struct {
		XX, XY FT_Fixed
		YX, YY FT_Fixed
	}

	FT_Vector struct {
		X FT_Pos
		Y FT_Pos
	}

	FT_CharMapRec struct {
		Face       FTFace
		Encoding   FT_Encoding
		PlatformId uint16
		EncodingId uint16
	}

	FT_GlyphSlotRec struct {
		Library           FT_Library
		Face              FTFace
		Next              FT_GlyphSlot
		Reserved          uint
		Generic           FT_Generic
		Metrics           FT_Glyph_Metrics
		LinearHoriAdvance FT_Fixed
		LinearVertAdvance FT_Fixed
		Advance           FT_Vector
		Format            FT_Glyph_Format
		Bitmap            FT_Bitmap
		BitmapLeft        int
		BitmapTop         int
		Outline           FT_Outline
		NumSubglyphs      uint
		Subglyphs         FT_SubGlyph
		ControlData       *Void
		ControlLen        Long
		LsbDelta          FT_Pos
		RsbDelta          FT_Pos
		Other             *Void
		Internal          FT_Slot_Internal
	}

	FT_Generic struct {
		Data      *Void
		Finalizer FT_Generic_Finalizer
	}

	FT_Glyph_Metrics struct {
		Width        FT_Pos
		Height       FT_Pos
		HoriBearingX FT_Pos
		HoriBearingY FT_Pos
		HoriAdvance  FT_Pos
		VertBearingX FT_Pos
		VertBearingY FT_Pos
		VertAdvance  FT_Pos
	}

	FT_Outline struct {
		NContours int16
		NPoints   int16
		Points    *FT_Vector
		Tags      *Char
		Contours  *int16
		Flags     int
	}

	FT_BBox struct {
		XMin, YMin FT_Pos
		XMax, YMax FT_Pos
	}

	FT_SizeRec struct {
		Face     FTFace
		Generic  FT_Generic
		Metrics  FT_Size_Metrics
		Internal FT_Size_Internal
	}

	FT_MemoryRec struct {
		User    *Void
		Alloc   FT_Alloc_Func
		Free    FT_Free_Func
		Realloc FT_Realloc_Func
	}

	FT_Size_Metrics struct {
		XPpem      uint16
		YPpem      uint16
		XScale     FT_Fixed
		YScale     FT_Fixed
		Ascender   FT_Pos
		Descender  FT_Pos
		Height     FT_Pos
		MaxAdvance FT_Pos
	}

	FT_ListRec struct {
		Head FT_ListNode
		Tail FT_ListNode
	}

	FT_ListNodeRec struct {
		Prev FT_ListNode
		Next FT_ListNode
		Data *Void
	}

	FT_GlyphRec struct {
		Library FT_Library
		Clazz   *FT_Glyph_Class
		Format  FT_Glyph_Format
		Advance FT_Vector
	}

	FT_Glyph_Class struct {
		GlyphSize      FT_Long
		GlyphFormat    FT_Glyph_Format
		GlyphInit      FT_Glyph_InitFunc
		GlyphDone      FT_Glyph_DoneFunc
		GlyphCopy      FT_Glyph_CopyFunc
		GlyphTransform FT_Glyph_TransformFunc
		GlyphBbox      FT_Glyph_GetBBoxFunc
		GlyphPrepare   FT_Glyph_PrepareFunc
	}

	PS_FontInfoRec struct {
		Version            *FT_String
		Notice             *FT_String
		FullName           *FT_String
		FamilyName         *FT_String
		Weight             *FT_String
		ItalicAngle        FT_Long
		IsFixedPitch       FT_Bool
		UnderlinePosition  int16
		UnderlineThickness uint16
	}

	PS_PrivateRec struct {
		UniqueId            int
		LenIV               int
		NumBlueValues       FT_Byte
		NumOtherBlues       FT_Byte
		NumFamilyBlues      FT_Byte
		NumFamilyOtherBlues FT_Byte
		BlueValues          [14]int16
		OtherBlues          [10]int16
		FamilyBlues         [14]int16
		FamilyOtherBlues    [10]int16
		BlueScale           FT_Fixed
		BlueShift           int
		BlueFuzz            int
		StandardWidth       [1]uint16
		StandardHeight      [1]uint16
		NumSnapWidths       FT_Byte
		NumSnapHeights      FT_Byte
		ForceBold           FT_Bool
		RoundStemUp         FT_Bool
		SnapWidths          [13]int16
		SnapHeights         [13]int16
		ExpansionFactor     FT_Fixed
		LanguageGroup       FT_Long
		Password            FT_Long
		MinFeature          [2]int16
	}

	FT_Module_Class struct {
		ModuleFlags     FT_ULong
		ModuleSize      FT_Long
		ModuleName      *FT_String
		ModuleVersion   FT_Fixed
		ModuleRequires  FT_Fixed
		ModuleInterface *Void
		ModuleInit      FT_Module_Constructor
		ModuleDone      FT_Module_Destructor
		GetInterface    FT_Module_Requester
	}

	FT_Multi_Master struct {
		NumAxis    uint
		NumDesigns uint
		Axis       [4]FT_MM_Axis
	}

	FT_MM_Axis struct {
		Name    *FT_String
		Minimum FT_Long
		Maximum FT_Long
	}

	FT_MM_Var struct {
		NumAxis        uint
		NumDesigns     uint
		NumNamedstyles uint
		Axis           *FT_Var_Axis
		Namedstyle     *FT_Var_Named_Style
	}

	FT_Var_Axis struct {
		Name    *FT_String
		Minimum FT_Fixed
		Def     FT_Fixed
		Maximum FT_Fixed
		Tag     FT_ULong
		Strid   uint
	}

	FT_Var_Named_Style struct {
		Coords *FT_Fixed
		Strid  uint
	}

	FT_Outline_Funcs struct {
		MoveTo  FT_Outline_MoveToFunc
		LineTo  FT_Outline_LineToFunc
		ConicTo FT_Outline_ConicToFunc
		CubicTo FT_Outline_CubicToFunc
		Shift   int
		Delta   FT_Pos
	}

	FT_Raster_Params struct {
		Target     *FT_Bitmap
		Source     *Void
		Flags      int
		GraySpans  FT_SpanFunc
		BlackSpans FT_SpanFunc
		BitTest    FT_Raster_BitTest_Func
		BitSet     FT_Raster_BitSet_Func
		User       *Void
		ClipBox    FT_BBox
	}

	FT_Span struct {
		X        int16
		Len      uint16
		Coverage uint8
	}

	BDF_PropertyRec struct {
		Type BDF_PropertyType
		// Union
		Atom *Char
		// Integer  FT_Int32
		// Cardinal FT_UInt32
	}

	FT_WinFNT_HeaderRec struct {
		Version              uint16
		FileSize             FT_ULong
		Copyright            [60]FT_Byte
		FileType             uint16
		NominalPointSize     uint16
		VerticalResolution   uint16
		HorizontalResolution uint16
		Ascent               uint16
		InternalLeading      uint16
		ExternalLeading      uint16
		Italic               FT_Byte
		Underline            FT_Byte
		StrikeOut            FT_Byte
		Weight               uint16
		Charset              FT_Byte
		PixelWidth           uint16
		PixelHeight          uint16
		PitchAndFamily       FT_Byte
		AvgWidth             uint16
		MaxWidth             uint16
		FirstChar            FT_Byte
		LastChar             FT_Byte
		DefaultChar          FT_Byte
		BreakChar            FT_Byte
		BytesPerRow          uint16
		DeviceOffset         FT_ULong
		FaceNameOffset       FT_ULong
		BitsPointer          FT_ULong
		BitsOffset           FT_ULong
		Reserved             FT_Byte
		Flags                FT_ULong
		ASpace               uint16
		BSpace               uint16
		CSpace               uint16
		ColorTableOffset     uint16
		Reserved1            [4]FT_ULong
	}

	FT_SfntName struct {
		PlatformId uint16
		EncodingId uint16
		LanguageId uint16
		NameId     uint16
		String     *FT_Byte
		StringLen  uint
	}

	FTC_ScalerRec struct {
		FaceId FTC_FaceID
		Width  uint
		Height uint
		Pixel  int
		XRes   uint
		YRes   uint
	}

	FTC_ImageTypeRec struct {
		FaceId FTC_FaceID
		Width  int
		Height int
		Flags  FT_Int32
	}

	FTC_SBitRec struct {
		Width    FT_Byte
		Height   FT_Byte
		Left     FT_Char
		Top      FT_Char
		Format   FT_Byte
		MaxGrays FT_Byte
		Pitch    int16
		Xadvance FT_Char
		Yadvance FT_Char
		Buffer   *FT_Byte
	}

	FTC_FontRec struct {
		FaceId    FTC_FaceID
		PixWidth  uint16
		PixHeight uint16
	}

	FT_Bitmap struct {
		Rows        int
		Width       int
		Pitch       int
		Buffer      *UnsignedChar
		NumGrays    int16
		PixelMode   Char
		PaletteMode Char
		Palette     *Void
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
