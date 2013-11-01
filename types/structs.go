package types

//TODO(t): add constants for [...]s

type (
	Tm struct {
		Sec, Min, Hour, Mday, Mon, Year, Wday, Yday, Isdst int
	}

	GDebugKey struct {
		Key   *Gchar
		Value uint
	}

	GError struct {
		Domain  Quark
		Code    int
		Message *Gchar
	}

	GList struct{} //REMOVE

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

	GPollFD struct {
		Fd      int
		Events  uint16
		Revents uint16
	}

	SList struct { // For functions in gobject and glib (import cycle
		Data Gpointer
		Next *SList
	}

	GSource struct{} //REMOVE

	GString struct{} // REMOVE

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
		Detail   Quark
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

	GtkRcProperty struct {
		TypeName     Quark
		PropertyName Quark
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
		InfoList *SList
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
