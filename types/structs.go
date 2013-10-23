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
		Items        *GList
		Fnc          GCompletionFunc
		Prefix       *Gchar
		Cache        *GList
		Strncmp_func GCompletionStrncmpFunc
	}

	GDate struct {
		Bits1, Bits2 uint
		// julian_days : 32
		// julian : 1
		// dmy : 1
		// day : 6
		// month : 4
		// year : 16
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
		Data      Gpointer
		Next      *GHook
		Prev      *GHook
		Ref_count uint
		Hook_id   Gulong
		Flags     uint
		Fnc       Gpointer
		Destroy   GDestroyNotify
	}

	GHookList struct {
		Seq_id Gulong
		Bits   uint
		// hook_size : 16
		// is_setup : 1
		Hooks         *GHook
		Dummy3        Gpointer
		Finalize_hook GHookFinalizeFunc
		Dummy         [2]Gpointer
	}

	GIOChannel struct {
		Ref_count         int
		Funcs             *GIOFuncs
		Encoding          *Gchar
		Read_cd           GIConv
		Write_cd          GIConv
		Line_term         *Gchar
		Line_term_len     uint
		Buf_size          Gsize
		Read_buf          *GString
		Encoded_read_buf  *GString
		Write_buf         *GString
		Partial_write_buf [6]Gchar
		Bits              uint
		// use_buffer : 1
		// do_encode : 1
		// close_on_unref : 1
		// is_readable : 1
		// is_writeable : 1
		// is_seekable : 1
		_, _ Gpointer
	}

	GList struct {
		Data Gpointer
		Next *GList
		Prev *GList
	}

	GIOFuncs struct {
		Io_read func(
			channel *GIOChannel,
			buf *Gchar,
			count Gsize,
			bytes_read *Gsize,
			err **GError) GIOStatus
		Io_write func(
			channel *GIOChannel,
			buf *Gchar,
			count Gsize,
			bytes_written *Gsize,
			err **GError) GIOStatus
		Io_seek func(
			channel *GIOChannel,
			offset int64,
			type_ GSeekType,
			err **GError) GIOStatus
		Io_close func(
			channel *GIOChannel,
			err **GError) GIOStatus
		Io_create_watch func(
			channel *GIOChannel,
			condition GIOCondition) *GSource
		Io_free func(
			channel *GIOChannel)
		Io_set_flags func(
			channel *GIOChannel,
			flags GIOFlags,
			err **GError) GIOStatus
		Io_get_flags func(
			channel *GIOChannel) GIOFlags
	}

	GMarkupParser struct {
		Start_element func(
			context *GMarkupParseContext,
			element_name *Gchar,
			attribute_names **Gchar,
			attribute_values **Gchar,
			user_data Gpointer,
			error **GError)
		End_element func(
			context *GMarkupParseContext,
			element_name *Gchar,
			user_data Gpointer,
			error **GError)
		Text func(
			context *GMarkupParseContext,
			text *Gchar,
			text_len Gsize,
			user_data Gpointer,
			error **GError)
		Passthrough func(
			context *GMarkupParseContext,
			passthrough_text *Gchar,
			text_len Gsize,
			user_data Gpointer,
			error **GError)
		Error func(
			context *GMarkupParseContext,
			error *GError,
			user_data Gpointer)
	}

	GNode struct {
		Data     Gpointer
		Next     *GNode
		Prev     *GNode
		Parent   *GNode
		Children *GNode
	}

	GSourceCallbackFuncs struct {
		Ref func(
			cb_data Gpointer)
		Unref func(
			cb_data Gpointer)
		Get func(
			cb_data Gpointer,
			source *GSource,
			fnc *GSourceFunc,
			data *Gpointer)
	}

	GOnce struct {
		Status GOnceStatus
		Retval Gpointer
	}

	GOptionEntry struct {
		Long_name       *Gchar
		Short_name      Gchar
		Flags           int
		Arg             GOptionArg
		Arg_data        Gpointer
		Description     *Gchar
		Arg_description *Gchar
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
		User_data        Gpointer
		Max_parse_errors uint
		Parse_errors     uint
		Input_name       *Gchar
		Qdata            *GData
		Config           *GScannerConfig
		Token            GTokenType
		Value            GTokenValue
		Line             uint
		Position         uint
		Next_token       GTokenType
		Next_value       GTokenValue
		Next_line        uint
		Next_position    uint
		Symbol_table     *GHashTable
		Input_fd         int
		Text             *Gchar
		Text_end         *Gchar
		Buffer           *Gchar
		Scope_id         uint
		Msg_handler      GScannerMsgFunc
	}

	GScannerConfig struct {
		Cset_skip_characters  *Gchar
		Cset_identifier_first *Gchar
		Cset_identifier_nth   *Gchar
		Cpair_comment_single  *Gchar
		Bits                  uint
		// case_sensitive : 1
		// skip_comment_multi : 1
		// skip_comment_single : 1
		// scan_comment_multi : 1
		// scan_identifier : 1
		// scan_identifier_1char : 1
		// scan_identifier_NULL : 1
		// scan_symbols : 1
		// scan_binary : 1
		// scan_octal : 1
		// scan_float : 1
		// scan_hex : 1
		// scan_hex_dollar : 1
		// scan_string_sq : 1
		// scan_string_dq : 1
		// numbers_2_int : 1
		// int_2_float : 1
		// identifier_2_string : 1
		// char_2_token : 1
		// symbol_2_token : 1
		// scope_0_fallback : 1
		// store_int64 : 1
		Padding_dummy uint
	}

	GSList struct {
		Data Gpointer
		Next *GSList
	}

	GSource struct {
		Callback_data  Gpointer
		Callback_funcs *GSourceCallbackFuncs
		Source_funcs   *GSourceFuncs
		Ref_count      uint
		Context        *GMainContext
		Priority       int
		Flags          uint
		Source_id      uint
		Poll_fds       *GSList
		Prev           *GSource
		Next           *GSource
		Name           *Char
		Priv           *GSourcePrivate
	}

	GStaticPrivate struct {
		Index uint
	}

	GStaticRecMutex struct {
		Mutex GStaticMutex
		Depth uint
		Owner GSystemThread
	}

	GSystemThread struct {
		// union
		// data[4] Char
		Dummy_double float64
		// dummy_pointer *Void
		// dummy_long Long
	}

	GStaticRWLock struct {
		Mutex         GStaticMutex
		Read_cond     *GCond
		Write_cond    *GCond
		Read_counter  uint
		Have_writer   Gboolean
		Want_to_read  uint
		Want_to_write uint
	}

	GString struct {
		Str           *Gchar
		Len           Gsize
		Allocated_len Gsize
	}

	GTestLogBuffer struct {
		Data *GString
		Msgs *GSList
	}

	GTestLogMsg struct {
		Log_type  GTestLogType
		N_strings uint
		Strings   **Gchar
		N_nums    uint
		Nums      *Long_double
	}

	GThread struct {
		Func     GThreadFunc
		Data     Gpointer
		Joinable Gboolean
		Priority GThreadPriority
	}

	GThreadFunctions struct {
		Mutex_new       func()
		Mutex_lock      func(mutex *GMutex) *GMutex
		Mutex_trylock   func(mutex *GMutex) Gboolean
		Mutex_unlock    func(mutex *GMutex)
		Mutex_free      func(mutex *GMutex)
		Cond_new        func() *GCond
		Cond_signal     func(cond *GCond)
		Cond_broadcast  func(cond *GCond)
		Cond_wait       func(cond *GCond, mutex *GMutex)
		Cond_timed_wait func(
			cond *GCond,
			mutex *GMutex,
			end_time *GTimeVal) Gboolean
		Cond_free     func(cond *GCond)
		Private_new   func(destructor GDestroyNotify) *GPrivate
		Private_get   func(private_key *GPrivate) Gpointer
		Private_set   func(private_key *GPrivate, data Gpointer)
		Thread_create func(
			fnc GThreadFunc,
			data Gpointer,
			stack_size Gulong,
			joinable Gboolean,
			bound Gboolean,
			priority GThreadPriority,
			thread Gpointer,
			error **GError)
		Thread_yield        func()
		Thread_join         func(thread Gpointer)
		Thread_exit         func()
		Thread_set_priority func(
			thread Gpointer,
			priority GThreadPriority)
		Thread_self  func(thread Gpointer)
		Thread_equal func(
			thread1 Gpointer,
			thread2 Gpointer) Gboolean
	}

	GThreadPool struct {
		Func      GFunc
		User_data Gpointer
		Exclusive Gboolean
	}

	GTimeVal struct {
		Tv_sec  Glong
		Tv_usec Glong
	}

	GTokenValue struct {
		// union
		// v_symbol  Gpointer;
		// v_identifier  *Gchar;
		// v_binary  Gulong;
		// v_octal  Gulong;
		// v_int  Gulong;
		V_int64 uint64
		// v_float  float64;
		// v_hex  Gulong;
		// v_string  *Gchar;
		// v_comment  *Gchar;
		// v_char  Guchar;
		// v_error  uint;
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

	GtkWidget struct{} //REMOVE

	GTypeClass struct {
		G_type GType
	}

	GTypeInstance struct {
		G_class *GTypeClass
	}

	GObject struct {
		G_type_instance GTypeInstance
		Ref_count       uint
		Qdata           *GData
	}

	GSourceFuncs struct {
		Prepare func(
			source *GSource,
			timeout_ *int) Gboolean
		Check func(
			source *GSource) Gboolean
		Dispatch func(
			source *GSource,
			callback GSourceFunc,
			user_data Gpointer) Gboolean
		Finalize func(
			source *GSource)
		Closure_callback GSourceFunc
		Closure_marshal  GSourceDummyMarshal
	}

	GdkColor struct{} //REMOVE

	GdkFont struct {
		Type    GdkFontType
		Ascent  int
		Descent int
	}

	GdkGC struct {
		Parent_instance GObject
		Clip_x_origin   int
		Clip_y_origin   int
		Ts_x_origin     int
		Ts_y_origin     int
		Colormap        *GdkColormap
	}

	GdkColormap struct{} //REMOVE

	GdkVisual struct{} //REMOVE

	GdkRectangle struct {
		X      int
		Y      int
		Width  int
		Height int
	}

	GdkPoint struct {
		X int
		Y int
	}

	GtkWindow struct{} //REMOVE

	GdkScreen struct{} //REMOVE

	GClosure struct {
		Bits uint
		// ref_count : 15
		// meta_marshal : 1
		// n_guards : 1
		// n_fnotifiers : 2
		// n_inotifiers : 8
		// in_inotify : 1
		// floating : 1
		// derivative_flag : 1
		// in_marshal : 1
		// is_invalid : 1
		Marshal func(
			closure *GClosure,
			return_value *GValue,
			n_param_values uint,
			param_values *GValue,
			invocation_hint Gpointer,
			marshal_data Gpointer)
		Data      Gpointer
		Notifiers *GClosureNotifyData
	}

	GValue struct {
		G_type GType
		//  UNION
		Data [2]uint64 // was union{v_int...}data[2]
		// v_int     int
		// v_uint    uint
		// v_long    glong
		// v_ulong   Gulong
		// v_int64   int64
		// v_uint64  uint64
		// v_float   float32
		// v_double  float64
		// v_pointer Gpointer
	}

	GClosureNotifyData struct {
		Data   Gpointer
		Notify GClosureNotify
	}

	GdkCursor struct{} //REMOVE

	GtkLabel struct{} // REMOVE

	GtkBox struct{} // REMOVE

	GtkTextBuffer struct{} //REMOVE

	GdkEventKey struct {
		Type             GdkEventType
		Window           *GdkWindow
		Send_event       int8
		Time             GUint32
		State            uint
		Keyval           uint
		Length           int
		String           *Gchar
		Hardware_keycode uint16
		Group            uint8
		Is_modifier      uint // : 1
	}

	GdkDisplay struct{} //REMOVE

	GdkKeyboardGrabInfo struct {
		Window        *GdkWindow
		Native_window *GdkWindow
		Serial        Gulong
		Owner_events  Gboolean
		Time          GUint32
	}

	GdkPointerWindowInfo struct {
		Toplevel_under_pointer *GdkWindow
		Window_under_pointer   *GdkWindow
		Toplevel_x, Toplevel_y float64
		State                  GUint32
		Button                 GUint32
		Motion_hint_serial     Gulong
	}

	GdkDevice struct {
		Parent_instance GObject
		Name            *Gchar
		Source          GdkInputSource
		Mode            GdkInputMode
		Has_cursor      Gboolean
		Num_axes        int
		Axes            *GdkDeviceAxis
		Num_keys        int
		Keys            *GdkDeviceKey
	}

	GdkDeviceKey struct {
		Keyval    uint
		Modifiers GdkModifierType
	}

	GdkDeviceAxis struct {
		Use GdkAxisUse
		Min float64
		Max float64
	}

	GdkEvent struct{} //REMOTE

	GSignalInvocationHint struct {
		Signal_id uint
		Detail    GQuark
		Run_type  GSignalFlags
	}

	GMountOperation struct {
		Parent_instance GObject
		Priv            *GMountOperationPrivate
	}

	GtkMenu struct{} //REMOVE

	PangoLogAttr struct {
		Bits uint
		// is_line_break : 1
		// is_mandatory_break : 1
		// is_char_break : 1
		// is_white : 1
		// is_cursor_position : 1
		// is_word_start : 1
		// is_word_end : 1
		// is_sentence_boundary : 1
		// is_sentence_start : 1
		// is_sentence_end : 1
		// backspace_deletes_character : 1
		// is_expandable_space : 1
		// is_word_boundary : 1
	}

	AtkObject struct {
		Parent            GObject
		Description       *Gchar
		Name              *Gchar
		Accessible_parent *AtkObject
		Role              AtkRole
		Relation_set      *AtkRelationSet
		Layer             AtkLayer
	}

	AtkRelationSet struct {
		Parent    GObject
		Relations *GPtrArray
	}

	GdkDragContext struct {
		Parent_instance  GObject
		Protocol         GdkDragProtocol
		Is_source        Gboolean
		Source_window    *GdkWindow
		Dest_window      *GdkWindow
		Targets          *GList
		Actions          GdkDragAction
		Suggested_action GdkDragAction
		Action           GdkDragAction
		Start_time       GUint32
		Windowing_data   Gpointer
	}

	GdkEventExpose struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		Area       GdkRectangle
		Region     *GdkRegion
		Count      int
	}

	GdkEventOwnerChange struct {
		Type           GdkEventType
		Window         *GdkWindow
		Send_event     int8
		Owner          GdkNativeWindow
		Reason         GdkOwnerChange
		Selection      GdkAtom
		Time           GUint32
		Selection_time GUint32
	}

	GdkEventProperty struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		Atom       GdkAtom
		Time       GUint32
		State      uint
	}

	GdkEventSelection struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		Selection  GdkAtom
		Target     GdkAtom
		Property   GdkAtom
		Time       GUint32
		Requestor  GdkNativeWindow
	}

	GdkEventSetting struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		Action     GdkSettingAction
		Name       *Char
	}

	GdkGCValues struct {
		Foreground         GdkColor
		Background         GdkColor
		Font               *GdkFont
		Function           GdkFunction
		Fill               GdkFill
		Tile               *GdkPixmap
		Stipple            *GdkPixmap
		Clip_mask          *GdkPixmap
		Subwindow_mode     GdkSubwindowMode
		Ts_x_origin        int
		Ts_y_origin        int
		Clip_x_origin      int
		Clip_y_origin      int
		Graphics_exposures int
		Line_width         int
		Line_style         GdkLineStyle
		Cap_style          GdkCapStyle
		Join_style         GdkJoinStyle
	}
	GdkCapStyle Enum //REMOVE

	GdkGeometry struct {
		Min_width   int
		Min_height  int
		Max_width   int
		Max_height  int
		Base_width  int
		Base_height int
		Width_inc   int
		Height_inc  int
		Min_aspect  float64
		Max_aspect  float64
		Win_gravity GdkGravity
	}

	GdkImage struct {
		Parent_instance GObject
		Type            GdkImageType
		Visual          *GdkVisual
		Byte_order      GdkByteOrder
		Width           int
		Height          int
		Depth           uint16
		Bpp             uint16
		Bpl             uint16
		Bits_per_pixel  uint16
		Mem             Gpointer
		Colormap        *GdkColormap
		Windowing_data  Gpointer
	}

	GObjectClass struct {
		G_type_class         GTypeClass
		Construct_properties *GSList
		Constructor          func(
			Type GType,
			n_construct_properties uint,
			construct_properties *GObjectConstructParam) *GObject
		Set_property func(
			object *GObject,
			property_id uint,
			GValue,
			pspec *GParamSpec)
		Get_property func(
			object *GObject,
			property_id uint,
			value *GValue,
			pspec *GParamSpec)
		Dispose func(
			object *GObject)
		Finalize func(
			object *GObject)
		Dispatch_properties_changed func(
			object *GObject,
			n_pspecs uint,
			pspecs **GParamSpec)
		Notify func(
			object *GObject,
			pspec *GParamSpec)
		Constructed func(
			object *GObject)
		Flags  Gsize
		Pdummy [6]Gpointer
	}

	GParamSpec struct {
		G_type_instance GTypeInstance
		Name            *Gchar
		Flags           GParamFlags
		Value_type      GType
		Owner_type      GType
		_nick           *Gchar
		_blurb          *Gchar
		Qdata           *GData
		Ref_count       uint
		Param_id        uint
	}

	GObjectConstructParam struct {
		Pspec *GParamSpec
		Value *GValue
	}

	GdkEventButton struct {
		Type           GdkEventType
		Window         *GdkWindow
		Send_event     int8
		Time           GUint32
		X              float64
		Y              float64
		Axes           *float64
		State          uint
		Button         uint
		Device         *GdkDevice
		X_root, Y_root float64
	}

	GdkEventScroll struct {
		Type           GdkEventType
		Window         *GdkWindow
		Send_event     int8
		Time           GUint32
		X              float64
		Y              float64
		State          uint
		Direction      GdkScrollDirection
		Device         *GdkDevice
		X_root, Y_root float64
	}

	GdkEventMotion struct {
		Type           GdkEventType
		Window         *GdkWindow
		Send_event     int8
		Time           GUint32
		X              float64
		Y              float64
		Axes           *float64
		State          uint
		Is_hint        int16
		Device         *GdkDevice
		X_root, Y_root float64
	}

	GdkEventAny struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
	}

	GdkEventCrossing struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		Subwindow  *GdkWindow
		Time       GUint32
		X          float64
		Y          float64
		X_root     float64
		Y_root     float64
		Mode       GdkCrossingMode
		Detail     GdkNotifyType
		Focus      Gboolean
		State      uint
	}
	GdkCrossingMode Enum //REMOVE

	GdkEventConfigure struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		X, Y       int
		Width      int
		Height     int
	}

	GdkEventFocus struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		In         int16
	}

	GdkEventProximity struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		Time       GUint32
		Device     *GdkDevice
	}

	GdkEventVisibility struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event int8
		State      GdkVisibilityState
	}
	GdkVisibilityState Enum //REMOVE

	GdkEventClient struct {
		Type         GdkEventType
		Window       *GdkWindow
		Send_event   int8
		Message_type GdkAtom
		Data_format  uint16
		// Union
		B [20]Char
		// s [10]Short
		// l [5]Long
	}

	GdkEventWindowState struct {
		Type             GdkEventType
		Window           *GdkWindow
		Send_event       int8
		Changed_mask     GdkWindowState
		New_window_state GdkWindowState
	}
	GdkWindowState Enum //REMOVE

	GdkEventGrabBroken struct {
		Yype        GdkEventType
		Window      *GdkWindow
		Send_event  int8
		Keyboard    Gboolean
		Implicit    Gboolean
		Grab_window *GdkWindow
	}

	GtkRcProperty struct {
		Type_name     GQuark
		Property_name GQuark
		Origin        *Gchar
		Value         GValue
	}

	/*
	   GtkWidgetAuxInfo  struct
	   {
	   x  int;
	   y  int;
	   width  int;
	   height  int;
	    uint x_set : 1;
	    uint y_set : 1;
	   };
	*/

	PangoColor struct {
		Red   uint16
		Green uint16
		Blue  uint16
	}

	GCancellable struct {
		Parent_instance GObject
		Priv            *GCancellablePrivate
	}

	GdkAppLaunchContext struct {
		Parent_instance GAppLaunchContext
		Priv            *GdkAppLaunchContextPrivate
	}

	GdkKeymap struct {
		Parent_instance GObject
		Display         *GdkDisplay
	}

	GdkKeymapKey struct {
		Keycode uint
		Group   int
		Level   int
	}

	GdkPangoRenderer struct {
		Parent_instance PangoRenderer
		Priv            *GdkPangoRendererPrivate
	}

	GAppLaunchContext struct {
		Parent_instance GObject
		Priv            *GAppLaunchContextPrivate
	}

	PangoRenderer struct {
		Parent_instance GObject
		Underline       PangoUnderline
		Strikethrough   Gboolean
		Active_count    int
		Matrix          *PangoMatrix
		Priv            *PangoRendererPrivate
	}

	PangoMatrix struct {
		Xx float64
		Xy float64
		Yx float64
		Yy float64
		X0 float64
		Y0 float64
	}

	PangoAttrShape struct {
		Attr         PangoAttribute
		Ink_rect     PangoRectangle
		Logical_rect PangoRectangle
		Data         Gpointer
		Copy_func    PangoAttrDataCopyFunc
		Destroy_func GDestroyNotify
	}

	PangoRectangle struct {
		X      int
		Y      int
		Width  int
		Height int
	}

	GdkPointerHooks struct {
		Get_pointer func(
			window *GdkWindow,
			x *int,
			y *int,
			mask *GdkModifierType) *GdkWindow
		Window_at_pointer func(screen *GdkScreen,
			win_x *int,
			win_y *int) *GdkWindow
	}

	GdkRgbCmap struct {
		Colors    [256]GUint32
		N_colors  int
		Info_list *GSList
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

	GInputStream struct {
		Parent_instance GObject
		Priv            *GInputStreamPrivate
	}

	PangoAttribute struct {
		Klass       *PangoAttrClass
		Start_index uint
		End_index   uint
	}

	PangoAttrClass struct {
		Type PangoAttrType
		Copy func(
			attr *PangoAttribute) *PangoAttribute
		Destroy func(
			attr *PangoAttribute)
		Equal func(
			attr1, attr2 *PangoAttribute) Gboolean
	}

	PangoGlyphItem struct {
		Item   *PangoItem
		Glyphs *PangoGlyphString
	}

	PangoItem struct {
		Offset    int
		Length    int
		Num_chars int
		Analysis  PangoAnalysis
	}

	PangoAnalysis struct {
		Shape_engine *PangoEngineShape
		Lang_engine  *PangoEngineLang
		Font         *PangoFont
		Level        uint8
		Gravity      uint8
		Flags        uint8
		Script       uint8
		Language     *PangoLanguage
		Extra_attrs  *GSList
	}

	PangoGlyphString struct {
		Num_glyphs   int
		Glyphs       *PangoGlyphInfo
		Log_clusters *int
		Space        int
	}

	PangoLayoutLine struct {
		Layout      *PangoLayout
		Start_index int
		Length      int
		Runs        *GSList
		Bits        uint
		// is_paragraph_start : 1
		// resolved_dir : 3
	}

	PangoGlyphInfo struct {
		Glyph    PangoGlyph
		Geometry PangoGlyphGeometry
		Attr     PangoGlyphVisAttr
	}

	PangoGlyphGeometry struct {
		Width    PangoGlyphUnit
		X_offset PangoGlyphUnit
		Y_offset PangoGlyphUnit
	}

	PangoGlyphVisAttr struct {
		Is_cluster_start uint // : 1
	}

	GdkPixdata struct {
		Magic        GUint32
		Length       GInt32
		Pixdata_type GdkPixdataType
		Rowstride    GUint32
		Width        GUint32
		Height       GUint32
		Pixel_data   *uint8
	}

	PangoGlyphItemIter struct {
		Glyph_item  *PangoGlyphItem
		Text        *Gchar
		Start_glyph int
		Start_index int
		Start_char  int
		End_glyph   int
		End_index   int
		End_char    int
	}

	CairoUserDataKey struct {
		Unused int
	}

	CairoMatrix struct {
		Xx, Yx, Xy, Yy, X0, Y0 float64
	}

	CairoRectangleList struct {
		Status         CairoStatus
		Rectangles     *CairoRectangle
		Num_rectangles int
	}

	CairoRectangle struct {
		X, Y, Width, Height float64
	}

	CairoGlyph struct {
		Index UnsignedLong
		X     float64
		Y     float64
	}

	CairoTextCluster struct {
		Num_bytes  int
		Num_glyphs int
	}

	CairoTextExtents struct {
		X_bearing float64
		Y_bearing float64
		Width     float64
		Height    float64
		X_advance float64
		Y_advance float64
	}

	CairoFontExtents struct {
		Ascent        float64
		Descent       float64
		Height        float64
		Max_x_advance float64
		Max_y_advance float64
	}

	CairoPath struct {
		Status   CairoStatus
		Data     *Cairo_path_data_t
		Num_data int
	}

	Cairo_path_data_t struct {
		// todo(t): Union
		// header struct {
		// 	Type   Cairo_path_data_type_t
		// 	length int
		// }
		Point struct {
			X, Y float64
		}
	}

	CairoRectangleInt struct {
		X, Y          int
		Width, Height int
	}

	CairoScriptInterpreterHooks struct {
		Closure         *Void
		Surface_create  CsiSurfaceCreateFunc
		Surface_destroy CsiDestroyFunc
		Context_create  Csi_context_create_func_t
		Context_destroy CsiDestroyFunc
		Show_page       Csi_show_page_func_t
		Copy_page       Csi_copy_page_func_t
	}

	GMemVTable struct {
		Malloc  func(n_bytes Gsize) Gpointer
		Realloc func(mem Gpointer,
			n_bytes Gsize) Gpointer
		Free   func(mem Gpointer)
		Calloc func(n_blocks Gsize,
			n_block_bytes Gsize) Gpointer
		Try_malloc  func(n_bytes Gsize) Gpointer
		Try_realloc func(mem Gpointer,
			n_bytes Gsize) Gpointer
	}

	GTrashStack struct {
		Next *GTrashStack
	}

	GailTextUtil struct {
		Parent GObject
		Buffer *GtkTextBuffer
	}

	AtkKeyEventStruct struct {
		Type      int
		State     uint
		Keyval    uint
		Length    int
		String    *Gchar
		Keycode   uint16
		Timestamp GUint32
	}

	AtkRegistry struct {
		Parent                  GObject
		Factory_type_registry   *GHashTable
		Factory_singleton_cache *GHashTable
	}

	AtkRelation struct {
		Parent       GObject
		Target       *GPtrArray
		Relationship AtkRelationType
	}

	AtkPropertyValues struct {
		Property_name *Gchar
		Old_value     GValue
		New_value     GValue
	}

	AtkGObjectAccessible struct {
		Parent AtkObject
	}

	AtkTextRectangle struct {
		X      int
		Y      int
		Width  int
		Height int
	}

	AtkTextRange struct {
		Bounds       AtkTextRectangle
		Start_offset int
		End_offset   int
		Content      *Gchar
	}

	AtkPlug struct {
		Parent AtkObject
	}

	PangoFcFont struct {
		Parent_instance PangoFont
		Font_pattern    *FcPattern
		Fontmap         *PangoFontMap
		Priv            Gpointer
		Matrix          PangoMatrix
		Description     *PangoFontDescription
		Metrics_by_lang *GSList
		Bits            uint
		// is_hinted : 1
		// is_transformed : 1
	}

	PangoFcFontMap struct {
		Parent_instance PangoFontMap
		Priv            *PangoFcFontMapPrivate
	}

	PangoOTRulesetDescription struct {
		Script                 PangoScript
		Language               *PangoLanguage
		Static_gsub_features   *PangoOTFeatureMap
		N_static_gsub_features uint
		Static_gpos_features   *PangoOTFeatureMap
		N_static_gpos_features uint
		Other_features         *PangoOTFeatureMap
		N_other_features       uint
	}

	PangoOTFeatureMap struct {
		Feature_name [5]Char
		Property_bit Gulong
	}

	PangoOTGlyph struct {
		Glyph      GUint32
		Properties uint
		Cluster    uint
		Component  uint16
		LigID      uint16
		Internal   uint
	}

	FT_FaceRec struct {
		Num_faces           FT_Long
		Face_index          FT_Long
		Face_flags          FT_Long
		Style_flags         FT_Long
		Num_glyphs          FT_Long
		Family_name         *FT_String
		Style_name          *FT_String
		Num_fixed_sizes     int
		Available_sizes     *FT_Bitmap_Size
		Num_charmaps        int
		Charmaps            *FT_CharMap
		Generic             FT_Generic
		Bbox                FT_BBox
		Units_per_EM        uint16
		Ascender            int16
		Descender           int16
		Height              int16
		Max_advance_width   int16
		Max_advance_height  int16
		Underline_position  int16
		Underline_thickness int16
		Glyph               FT_GlyphSlot
		Size                FT_Size
		Charmap             FT_CharMap
		Driver              FT_Driver
		Memory              FT_Memory
		Stream              FT_Stream
		Sizes_list          FT_ListRec
		Autohint            FT_Generic
		Extensions          *Void
		Internal            FT_Face_Internal
	}

	FT_Bitmap_Size struct {
		Height int16
		Width  int16
		Size   FT_Pos
		X_ppem FT_Pos
		Y_ppem FT_Pos
	}

	FT_Generic_Finalizer func(
		Object *Void)

	FT_Open_Args struct {
		Flags       uint
		Memory_base *FT_Byte
		Memory_size FT_Long
		Pathname    *FT_String
		Stream      FT_Stream
		Driver      FT_Module
		Num_params  int
		Params      *FT_Parameter
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
		Face        FTFace
		Encoding    FT_Encoding
		Platform_id uint16
		Encoding_id uint16
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
		Bitmap_left       int
		Bitmap_top        int
		Outline           FT_Outline
		Num_subglyphs     uint
		Subglyphs         FT_SubGlyph
		Control_data      *Void
		Control_len       Long
		Lsb_delta         FT_Pos
		Rsb_delta         FT_Pos
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
		N_contours int16
		N_points   int16
		Points     *FT_Vector
		Tags       *Char
		Contours   *int16
		Flags      int
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
		X_ppem      uint16
		Y_ppem      uint16
		X_scale     FT_Fixed
		Y_scale     FT_Fixed
		Ascender    FT_Pos
		Descender   FT_Pos
		Height      FT_Pos
		Max_advance FT_Pos
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
		Glyph_size      FT_Long
		Glyph_format    FT_Glyph_Format
		Glyph_init      FT_Glyph_InitFunc
		Glyph_done      FT_Glyph_DoneFunc
		Glyph_copy      FT_Glyph_CopyFunc
		Glyph_transform FT_Glyph_TransformFunc
		Glyph_bbox      FT_Glyph_GetBBoxFunc
		Glyph_prepare   FT_Glyph_PrepareFunc
	}

	PS_FontInfoRec struct {
		Version             *FT_String
		Notice              *FT_String
		Full_name           *FT_String
		Family_name         *FT_String
		Weight              *FT_String
		Italic_angle        FT_Long
		Is_fixed_pitch      FT_Bool
		Underline_position  int16
		Underline_thickness uint16
	}

	PS_PrivateRec struct {
		Unique_id              int
		LenIV                  int
		Num_blue_values        FT_Byte
		Num_other_blues        FT_Byte
		Num_family_blues       FT_Byte
		Num_family_other_blues FT_Byte
		Blue_values            [14]int16
		Other_blues            [10]int16
		Family_blues           [14]int16
		Family_other_blues     [10]int16
		Blue_scale             FT_Fixed
		Blue_shift             int
		Blue_fuzz              int
		Standard_width         [1]uint16
		Standard_height        [1]uint16
		Num_snap_widths        FT_Byte
		Num_snap_heights       FT_Byte
		Force_bold             FT_Bool
		Round_stem_up          FT_Bool
		Snap_widths            [13]int16
		Snap_heights           [13]int16
		Expansion_factor       FT_Fixed
		Language_group         FT_Long
		Password               FT_Long
		Min_feature            [2]int16
	}

	FT_Module_Class struct {
		Module_flags     FT_ULong
		Module_size      FT_Long
		Module_name      *FT_String
		Module_version   FT_Fixed
		Module_requires  FT_Fixed
		Module_interface *Void
		Module_init      FT_Module_Constructor
		Module_done      FT_Module_Destructor
		Get_interface    FT_Module_Requester
	}

	FT_Multi_Master struct {
		Num_axis    uint
		Num_designs uint
		Axis        [4]FT_MM_Axis
	}

	FT_MM_Axis struct {
		Name    *FT_String
		Minimum FT_Long
		Maximum FT_Long
	}

	FT_MM_Var struct {
		Num_axis        uint
		Num_designs     uint
		Num_namedstyles uint
		Axis            *FT_Var_Axis
		Namedstyle      *FT_Var_Named_Style
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
		Move_to  FT_Outline_MoveToFunc
		Line_to  FT_Outline_LineToFunc
		Conic_to FT_Outline_ConicToFunc
		Cubic_to FT_Outline_CubicToFunc
		Shift    int
		Delta    FT_Pos
	}

	FT_Raster_Params struct {
		Target      *FT_Bitmap
		Source      *Void
		Flags       int
		Gray_spans  FT_SpanFunc
		Black_spans FT_SpanFunc
		Bit_test    FT_Raster_BitTest_Func
		Bit_set     FT_Raster_BitSet_Func
		User        *Void
		Clip_box    FT_BBox
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
		// integer  FT_Int32
		// cardinal FT_UInt32
	}

	FT_WinFNT_HeaderRec struct {
		Version               uint16
		File_size             FT_ULong
		Copyright             [60]FT_Byte
		File_type             uint16
		Nominal_point_size    uint16
		Vertical_resolution   uint16
		Horizontal_resolution uint16
		Ascent                uint16
		Internal_leading      uint16
		External_leading      uint16
		Italic                FT_Byte
		Underline             FT_Byte
		Strike_out            FT_Byte
		Weight                uint16
		Charset               FT_Byte
		Pixel_width           uint16
		Pixel_height          uint16
		Pitch_and_family      FT_Byte
		Avg_width             uint16
		Max_width             uint16
		First_char            FT_Byte
		Last_char             FT_Byte
		Default_char          FT_Byte
		Break_char            FT_Byte
		Bytes_per_row         uint16
		Device_offset         FT_ULong
		Face_name_offset      FT_ULong
		Bits_pointer          FT_ULong
		Bits_offset           FT_ULong
		Reserved              FT_Byte
		Flags                 FT_ULong
		A_space               uint16
		B_space               uint16
		C_space               uint16
		Color_table_offset    uint16
		Reserved1             [4]FT_ULong
	}

	FT_SfntName struct {
		Platform_id uint16
		Encoding_id uint16
		Language_id uint16
		Name_id     uint16
		String      *FT_Byte
		String_len  uint
	}

	FTC_ScalerRec struct {
		Face_id FTC_FaceID
		Width   uint
		Height  uint
		Pixel   int
		X_res   uint
		Y_res   uint
	}

	FTC_ImageTypeRec struct {
		Face_id FTC_FaceID
		Width   int
		Height  int
		Flags   FT_Int32
	}

	FTC_SBitRec struct {
		Width     FT_Byte
		Height    FT_Byte
		Left      FT_Char
		Top       FT_Char
		Format    FT_Byte
		Max_grays FT_Byte
		Pitch     int16
		Xadvance  FT_Char
		Yadvance  FT_Char
		Buffer    *FT_Byte
	}

	FTC_FontRec struct {
		Face_id    FTC_FaceID
		Pix_width  uint16
		Pix_height uint16
	}

	FT_Bitmap struct {
		Rows         int
		Width        int
		Pitch        int
		Buffer       *UnsignedChar
		Num_grays    int16
		Pixel_mode   Char
		Palette_mode Char
		Palette      *Void
	}

	GParameter struct {
		Name  string
		Value GValue
	}

	GSignalQuery struct {
		Signal_id    uint
		Signal_name  string
		Itype        GType
		Signal_flags GSignalFlags
		Return_type  GType
		N_params     uint
		Param_types  *GType
	}

	GTypeQuery struct {
		Type          GType
		Type_name     string
		Class_size    uint
		Instance_size uint
	}

	GApplication struct {
		Parent_instance GObject
		Priv            *GApplicationPrivate
	}

	GApplicationCommandLine struct {
		Parent_instance GObject
		Priv            *GApplicationCommandLinePrivate
	}

	GBufferedInputStream struct {
		Parent_instance GFilterInputStream
		Priv            *GBufferedInputStreamPrivate
	}

	GFilterInputStream struct {
		Parent_instance GInputStream
		Base_stream     *GInputStream
	}

	GOutputStream struct {
		Parent_instance GObject
		Priv            *GOutputStreamPrivate
	}

	GBufferedOutputStream struct {
		Parent_instance GFilterOutputStream
		Priv            *GBufferedOutputStreamPrivate
	}

	GFilterOutputStream struct {
		Parent_instance GOutputStream
		Base_stream     *GOutputStream
	}

	GConverterInputStream struct {
		Parent_instance GFilterInputStream
		Priv            *GConverterInputStreamPrivate
	}

	GConverterOutputStream struct {
		Parent_instance GFilterOutputStream
		Priv            *GConverterOutputStreamPrivate
	}

	GDataInputStream struct {
		Parent_instance GBufferedInputStream
		Priv            *GDataInputStreamPrivate
	}

	GDataOutputStream struct {
		Parent_instance GFilterOutputStream
		Priv            *GDataOutputStreamPrivate
	}

	GIOStream struct {
		Parent_instance GObject
		Priv            *GIOStreamPrivate
	}

	GDBusInterfaceInfo struct {
		Ref_count   int
		Name        *Gchar
		Methods     **GDBusMethodInfo
		Signals     **GDBusSignalInfo
		Properties  **GDBusPropertyInfo
		Annotations **GDBusAnnotationInfo
	}

	GDBusMethodInfo struct {
		Ref_count   int
		Name        *Gchar
		In_args     **GDBusArgInfo
		Out_args    **GDBusArgInfo
		Annotations **GDBusAnnotationInfo
	}

	GDBusSignalInfo struct {
		Ref_count   int
		Name        *Gchar
		Args        **GDBusArgInfo
		Annotations **GDBusAnnotationInfo
	}

	GDBusPropertyInfo struct {
		Ref_count   int
		Name        *Gchar
		Signature   *Gchar
		Flags       GDBusPropertyInfoFlags
		Annotations **GDBusAnnotationInfo
	}

	GDBusArgInfo struct {
		Ref_count   int
		Name        *Gchar
		Signature   *Gchar
		Annotations **GDBusAnnotationInfo
	}

	GDBusAnnotationInfo struct {
		Ref_count   int
		Key         *Gchar
		Value       *Gchar
		Annotations **GDBusAnnotationInfo
	}

	GDBusInterfaceVTable struct {
		Method_call  GDBusInterfaceMethodCallFunc
		Get_property GDBusInterfaceGetPropertyFunc
		Set_property GDBusInterfaceSetPropertyFunc
		Padding      [8]Gpointer
	}

	GDBusNodeInfo struct {
		Ref_count   int
		Path        *Gchar
		Interfaces  **GDBusInterfaceInfo
		Nodes       **GDBusNodeInfo
		Annotations **GDBusAnnotationInfo
	}

	GDBusProxy struct {
		Parent_instance GObject
		Priv            *GDBusProxyPrivate
	}

	GDBusSubtreeVTable struct {
		Enumerate  GDBusSubtreeEnumerateFunc
		Introspect GDBusSubtreeIntrospectFunc
		Dispatch   GDBusSubtreeDispatchFunc
		Padding    [8]Gpointer
	}

	GDBusErrorEntry struct {
		Error_code      int
		Dbus_error_name *Gchar
	}

	GEmblemedIcon struct {
		Parent_instance GObject
		Priv            *GEmblemedIconPrivate
	}

	GFileAttributeInfoList struct {
		Infos   *GFileAttributeInfo
		N_infos int
	}

	GFileAttributeInfo struct {
		Name  *Char
		Type  GFileAttributeType
		Flags GFileAttributeInfoFlags
	}

	GFileOutputStream struct {
		Parent_instance GOutputStream
		Priv            *GFileOutputStreamPrivate
	}

	GFileIOStream struct {
		Parent_instance GIOStream
		Priv            *GFileIOStreamPrivate
	}

	GFileEnumerator struct {
		Parent_instance GObject
		Priv            *GFileEnumeratorPrivate
	}

	GFileInputStream struct {
		Parent_instance GInputStream
		Priv            *GFileInputStreamPrivate
	}

	GFileMonitor struct {
		Parent_instance GObject
		Priv            *GFileMonitorPrivate
	}

	GInetAddress struct {
		Parent_instance GObject
		Priv            *GInetAddressPrivate
	}

	GInetSocketAddress struct {
		Parent_instance GSocketAddress
		Priv            *GInetSocketAddressPrivate
	}

	GSocketAddress struct {
		Parent_instance GObject
	}

	GMemoryOutputStream struct {
		Parent_instance GOutputStream
		Priv            *GMemoryOutputStreamPrivate
	}

	GNetworkAddress struct {
		Parent_instance GObject
		Priv            *GNetworkAddressPrivate
	}

	GNetworkService struct {
		Parent_instance GObject
		Priv            *GNetworkServicePrivate
	}

	GMemoryInputStream struct {
		Parent_instance GInputStream
		Priv            *GMemoryInputStreamPrivate
	}

	GPermission struct {
		Parent_instance GObject
		Priv            *GPermissionPrivate
	}

	GProxyAddress struct {
		Parent_instance GInetSocketAddress
		Priv            *GProxyAddressPrivate
	}

	GResolver struct {
		Parent_instance GObject
		Priv            *GResolverPrivate
	}

	GSettings struct {
		Parent_instance GObject
		Priv            *GSettingsPrivate
	}

	GSimpleActionGroup struct {
		Parent_instance GObject
		Priv            *GSimpleActionGroupPrivate
	}

	GSimpleAction struct {
		Parent_instance GObject
		Priv            *GSimpleActionPrivate
	}

	GSocketAddressEnumerator struct {
		Parent_instance GObject
	}

	GSocket struct {
		Parent_instance GObject
		Priv            *GSocketPrivate
	}

	GSocketClient struct {
		Parent_instance GObject
		Priv            *GSocketClientPrivate
	}

	GSocketConnection struct {
		Parent_instance GIOStream
		Priv            *GSocketConnectionPrivate
	}

	GSocketControlMessage struct {
		Parent_instance GObject
		Priv            *GSocketControlMessagePrivate
	}

	GSocketListener struct {
		Parent_instance GObject
		Priv            *GSocketListenerPrivate
	}

	GSocketService struct {
		Parent_instance GSocketListener
		Priv            *GSocketServicePrivate
	}

	GTlsCertificate struct {
		Parent_instance GObject
		Priv            *GTlsCertificatePrivate
	}

	GTcpConnection struct {
		Parent_instance GSocketConnection
		Priv            *GTcpConnectionPrivate
	}

	GTlsConnection struct {
		Parent_instance GIOStream
		Priv            *GTlsConnectionPrivate
	}

	GVfs struct {
		Parent_instance GObject
	}

	GVolumeMonitor struct {
		Parent_instance GObject
		Priv            Gpointer
	}

	GInputVector struct {
		Buffer Gpointer
		Size   Gsize
	}

	GOutputVector struct {
		Buffer Gconstpointer
		Size   Gsize
	}

	GTcpWrapperConnection struct {
		Parent_instance GTcpConnection
		Priv            *GTcpWrapperConnectionPrivate
	}

	GWin32InputStream struct {
		Parent_instance GInputStream
		Priv            *GWin32InputStreamPrivate
	}

	GWin32OutputStream struct {
		Parent_instance GOutputStream
		Priv            *GWin32OutputStreamPrivate
	}

	GEnumClass struct {
		G_type_class GTypeClass
		Minimum      int
		Maximum      int
		N_values     uint
		Values       *GEnumValue
	}

	GEnumValue struct {
		Value      int
		Value_name *Gchar
		Value_nick *Gchar
	}

	GFlagsClass struct {
		G_type_class GTypeClass
		Mask         uint
		N_values     uint
		Values       *GFlagsValue
	}

	GFlagsValue struct {
		Value      uint
		Value_name *Gchar
		Value_nick *Gchar
	}

	GTypeInfo struct {
		Class_size     uint16
		Base_init      GBaseInitFunc
		Base_finalize  GBaseFinalizeFunc
		Class_init     GClassInitFunc
		Class_finalize GClassFinalizeFunc
		Class_data     Gconstpointer
		Instance_size  uint16
		N_preallocs    uint16
		Instance_init  GInstanceInitFunc
		Value_table    *GTypeValueTable
	}

	GTypeValueTable struct {
		Value_init func(value *GValue)
		Value_free func(value *GValue)
		Value_copy func(
			src_value *GValue, dest_value *GValue)
		Value_peek_pointer func(value *GValue) Gpointer
		Collect_format     *Gchar
		Collect_value      func(
			value *GValue,
			n_collect_values uint,
			collect_values *GTypeCValue,
			collect_flags uint) *Gchar
		Lcopy_format *Gchar
		Lcopy_value  func(
			value *GValue,
			n_collect_values uint,
			collect_values *GTypeCValue,
			collect_flags uint) *Gchar
	}

	GValueArray struct {
		N_values     uint
		Values       *GValue
		N_prealloced uint
	}

	GTypeFundamentalInfo struct {
		Type_flags GTypeFundamentalFlags
	}

	GInterfaceInfo struct {
		Interface_init     GInterfaceInitFunc
		Interface_finalize GInterfaceFinalizeFunc
		Interface_data     Gpointer
	}

	GTypeModule struct {
		Parent_instance GObject
		Use_count       uint
		Type_infos      *GSList
		Interface_infos *GSList
		Name            *Gchar
	}

	GParamSpecTypeInfo struct {
		Instance_size     uint16
		N_preallocs       uint16
		Instance_init     func(pspec *GParamSpec)
		Value_type        GType
		Finalize          func(pspec *GParamSpec)
		Value_set_default func(
			pspec *GParamSpec, value *GValue)
		Value_validate func(
			pspec *GParamSpec, value *GValue) Gboolean
		Values_cmp func(pspec *GParamSpec,
			value1 *GValue, value2 *GValue) int
	}

	PangoIncludedModule struct {
		List   func(engines **PangoEngineInfo, n_engines *int)
		Init   func(module *GTypeModule)
		Exit   func()
		Create func(id *Char) *PangoEngine
	}

	PangoEngineScriptInfo struct {
		Script PangoScript
		Langs  *Gchar
	}

	PangoEngineInfo struct {
		Id          *Gchar
		Engine_type *Gchar
		Render_type *Gchar
		Scripts     *PangoEngineScriptInfo
		N_scripts   int
	}

	GTestConfig struct {
		Test_initialized Gboolean
		Test_quick       Gboolean
		Test_perf        Gboolean
		Test_verbose     Gboolean
		Test_quiet       Gboolean
	}
)
