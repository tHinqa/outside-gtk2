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

	GtkWidget struct {
		Object        GtkObject
		Private_flags uint16
		State         uint8
		Saved_state   uint8
		Name          *Gchar
		Style         *GtkStyle
		Requisition   GtkRequisition
		Allocation    GtkAllocation
		Window        *GdkWindow
		Parent        *GtkWidget
	}

	GtkRequisition struct {
		Width  int
		Height int
	}

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

	GtkObject struct {
		Parent_instance GInitiallyUnowned
		Flags           GUint32
	}

	GdkColor struct {
		Pixel GUint32
		Red   uint16
		Green uint16
		Blue  uint16
	}

	GtkStyle struct {
		Parent_instance   GObject
		Fg                [5]GdkColor
		Bg                [5]GdkColor
		Light             [5]GdkColor
		Dark              [5]GdkColor
		Mid               [5]GdkColor
		Text              [5]GdkColor
		Base              [5]GdkColor
		Text_aa           [5]GdkColor
		Black             GdkColor
		White             GdkColor
		Font_desc         *PangoFontDescription
		Xthickness        int
		Ythickness        int
		Fg_gc             *[5]GdkGC //TODO(t): CHECK
		Bg_gc             *[5]GdkGC //TODO(t): CHECK
		Light_gc          *[5]GdkGC //TODO(t): CHECK
		Dark_gc           *[5]GdkGC //TODO(t): CHECK
		Mid_gc            *[5]GdkGC //TODO(t): CHECK
		Text_gc           *[5]GdkGC //TODO(t): CHECK
		Base_gc           *[5]GdkGC //TODO(t): CHECK
		Text_aa_gc        *[5]GdkGC //TODO(t): CHECK
		Black_gc          *GdkGC
		White_gc          *GdkGC
		Bg_pixmap         *[5]GdkPixmap //TODO(t): CHECK
		Attach_count      int
		Depth             int
		Colormap          *GdkColormap
		Private_font      *GdkFont
		Private_font_desc *PangoFontDescription
		Rc_style          *GtkRcStyle
		Styles            *GSList
		Property_cache    *GArray
		Icon_factories    *GSList
	}

	GtkRcStyle struct {
		Parent_instance  GObject
		Name             *Gchar
		Bg_pixmap_name   *[5]Gchar //TODO(t): CHECK
		Font_desc        *PangoFontDescription
		Color_flags      [5]GtkRcFlags
		Fg               [5]GdkColor
		Bg               [5]GdkColor
		Text             [5]GdkColor
		Base             [5]GdkColor
		Xthickness       int
		Ythickness       int
		Rc_properties    *GArray
		Rc_style_lists   *GSList
		Icon_factories   *GSList
		Engine_specified uint //: 1;
	}

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

	GdkColormap struct {
		Parent_instance GObject
		Size            int
		Colors          *GdkColor
		Visual          *GdkVisual
		Windowing_data  Gpointer
	}

	GdkVisual struct {
		Parent_instance GObject
		Type            GdkVisualType
		Depth           int
		Byte_order      GdkByteOrder
		Colormap_size   int
		Bits_per_rgb    int
		Red_mask        GUint32
		Red_shift       int
		Red_prec        int
		Green_mask      GUint32
		Green_shift     int
		Green_prec      int
		Blue_mask       GUint32
		Blue_shift      int
		Blue_prec       int
	}

	GdkRectangle struct {
		X      int
		Y      int
		Width  int
		Height int
	}

	GtkCurve struct {
		Graph         GtkDrawingArea
		Cursor_type   int
		Min_x         float32
		Max_x         float32
		Min_y         float32
		Max_y         float32
		Pixmap        *GdkPixmap
		Curve_type    GtkCurveType
		Height        int
		Grab_point    int
		Last          int
		Num_points    int
		Point         *GdkPoint
		Num_ctlpoints int
		Ctlpoint      [2]*float32 //TODO(t): float32 (*ctlpoint)[2]; ???
	}

	GtkDrawingArea struct {
		Widget    GtkWidget
		Draw_data Gpointer
	}

	GdkPoint struct {
		X int
		Y int
	}

	GtkFileSelection struct {
		Parent_instance  GtkDialog
		Dir_list         *GtkWidget
		File_list        *GtkWidget
		Selection_entry  *GtkWidget
		Selection_text   *GtkWidget
		Main_vbox        *GtkWidget
		Ok_button        *GtkWidget
		Cancel_button    *GtkWidget
		Help_button      *GtkWidget
		History_pulldown *GtkWidget
		History_menu     *GtkWidget
		History_list     *GList
		Fileop_dialog    *GtkWidget
		Fileop_entry     *GtkWidget
		Fileop_file      *Gchar
		Cmpl_state       Gpointer
		Fileop_c_dir     *GtkWidget
		Fileop_del_file  *GtkWidget
		Fileop_ren_file  *GtkWidget
		Button_area      *GtkWidget
		Action_area      *GtkWidget
		Selected_names   *GPtrArray
		Last_selected    *Gchar
	}

	GtkDialog struct {
		Window      GtkWindow
		Vbox        *GtkWidget
		Action_area *GtkWidget
		Separator   *GtkWidget
	}

	GtkBin struct { // REMOVE
		Container GtkContainer
		Child     *GtkWidget
	}

	GtkContainer struct {
		Widget      GtkWidget
		Focus_child *GtkWidget
		Bits        uint
		// border_width : 16
		// need_resize : 1
		// resize_mode : 2
		// reallocate_redraws : 1
		// has_focus_chain : 1
	}

	GtkWindow struct {
		Bin                     GtkBin
		Title                   *Gchar
		Wmclass_name            *Gchar
		Wmclass_class           *Gchar
		Wm_role                 *Gchar
		Focus_widget            *GtkWidget
		Default_widget          *GtkWidget
		Transient_parent        *GtkWindow
		Geometry_info           *GtkWindowGeometryInfo
		Frame                   *GdkWindow
		Group                   *GtkWindowGroup
		Configure_request_count uint16
		Bits                    uint
		// allow_shrink : 1
		// allow_grow : 1
		// configure_notify_received : 1
		// need_default_position : 1
		// need_default_size : 1
		// position : 3
		// type : 4
		// has_user_ref_count : 1
		// has_focus : 1
		// modal : 1
		// destroy_with_parent : 1
		// has_frame : 1
		// iconify_initially : 1
		// stick_initially : 1
		// maximize_initially : 1
		// decorated : 1
		// type_hint : 3
		// gravity : 5
		// is_active : 1
		// has_toplevel_focus : 1
		Frame_left           uint
		Frame_top            uint
		Frame_right          uint
		Frame_bottom         uint
		Keys_changed_handler uint
		Mnemonic_modifier    GdkModifierType
		Screen               *GdkScreen
	}

	GtkWindowGroup struct {
		Parent_instance GObject
		Grabs           *GSList
	}

	GdkScreen struct {
		Parent_instance GObject
		Closed          uint       // : 1
		Normal_gcs      *[32]GdkGC //TODO(t): CHECK
		Exposure_gcs    *[32]GdkGC //TODO(t): CHECK
		Subwindow_gcs   *[32]GdkGC //TODO(t): CHECK
		Font_options    *CairoFontOptions
		Resolution      float64
	}

	GtkItemFactory struct {
		Object           GtkObject
		Path             *Gchar
		Accel_group      *GtkAccelGroup
		Widget           *GtkWidget
		Items            *GSList
		Translate_func   GtkTranslateFunc
		Translate_data   Gpointer
		Translate_notify GDestroyNotify
	}

	GtkAccelGroup struct { // REMOVE
		Parent         GObject
		Lock_count     uint
		Modifier_mask  GdkModifierType
		Acceleratables *GSList
		N_accels       uint
		Priv_accels    *GtkAccelGroupEntry
	}

	GtkAccelGroupEntry struct { // REMOVE
		Key              GtkAccelKey
		Closure          *GClosure
		Accel_path_quark GQuark
	}

	GtkAccelKey struct {
		Accel_key   uint
		Accel_mods  GdkModifierType
		Accel_flags uint //: 16
	}

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

	GtkItemFactoryEntry struct {
		Path            *Gchar
		Accelerator     *Gchar
		Callback        GtkItemFactoryCallback
		Callback_action uint
		Item_type       *Gchar
		Extra_data      Gconstpointer
	}

	GtkMenuEntry struct {
		Path          *Gchar
		Accelerator   *Gchar
		Callback      GtkMenuCallback
		Callback_data Gpointer
		Widget        *GtkWidget
	}

	GtkItem struct {
		Bin GtkBin
	}

	GtkOldEditable struct {
		Widget              GtkWidget
		Current_pos         uint
		Selection_start_pos uint
		Selection_end_pos   uint
		Bits                uint
		// has_selection : 1
		// editable : 1
		// visible : 1
		Clipboard_text *Gchar
	}

	GtkOptionMenu struct {
		Button    GtkButton
		Menu      *GtkWidget
		Menu_item *GtkWidget
		Width     uint16
		Height    uint16
	}

	GtkButton struct { // REMOVE
		Bin              GtkBin
		Event_window     *GdkWindow
		Label_text       *Gchar
		Activate_timeout uint
		Bits             uint
		// constructed : 1
		// in_button : 1
		// button_down : 1
		// relief : 2
		// use_underline : 1
		// use_stock : 1
		// depressed : 1
		// depress_on_activate : 1
		// focus_on_click : 1
	}

	GtkPreview struct {
		Widget        GtkWidget
		Buffer        *Guchar
		Buffer_width  uint16
		Buffer_height uint16
		Bpp           uint16
		Rowstride     uint16
		Dither        GdkRgbDither
		Bits          uint
		// type : 1
		// expand : 1
	}

	GtkPreviewInfo struct {
		Lookup *Guchar
		Gamma  float64
	}

	GtkTipsQuery struct {
		Label GtkLabel
		Bits  uint
		// emit_always : 1
		// in_query : 1
		Label_inactive *Gchar
		Label_no_tip   *Gchar
		Caller         *GtkWidget
		Last_crossed   *GtkWidget
		Query_cursor   *GdkCursor
	}

	GdkCursor struct {
		Type      GdkCursorType
		Ref_count uint
	}

	GtkLabel struct {
		Misc  GtkMisc
		Label *Gchar
		Bits  uint
		// jtype : 2
		// wrap : 1
		// use_underline : 1
		// use_markup : 1
		// ellipsize : 3
		// single_line_mode : 1
		// have_transform : 1
		// in_click : 1
		// wrap_mode : 3
		// pattern_set : 1
		// track_links : 1
		Mnemonic_keyval uint
		Text            *Gchar
		Attrs           *PangoAttrList
		Effective_attrs *PangoAttrList
		Layout          *PangoLayout
		Mnemonic_widget *GtkWidget
		Mnemonic_window *GtkWindow
		Select_info     *GtkLabelSelectionInfo
	}

	GtkMisc struct {
		Widget GtkWidget
		Xalign float32
		Yalign float32
		Xpad   uint16
		Ypad   uint16
	}

	GtkCList struct { // REMOVE
		Container           GtkContainer
		Flags               uint16
		Reserved1           Gpointer
		Reserved2           Gpointer
		Freeze_count        uint
		Internal_allocation GdkRectangle
		Rows                int
		Row_height          int
		Row_list            *GList
		Row_list_end        *GList
		Columns             int
		Column_title_area   GdkRectangle
		Title_window        *GdkWindow
		Column              *GtkCListColumn
		Clist_window        *GdkWindow
		Clist_window_width  int
		Clist_window_height int
		Hoffset             int
		Voffset             int
		Shadow_type         GtkShadowType
		Selection_mode      GtkSelectionMode
		Selection           *GList
		Selection_end       *GList
		Undo_selection      *GList
		Undo_unselection    *GList
		Undo_anchor         int
		Button_actions      [5]uint8
		Drag_button         uint8
		Click_cell          GtkCListCellInfo
		Hadjustment         *GtkAdjustment
		Vadjustment         *GtkAdjustment
		Xor_gc              *GdkGC
		Fg_gc               *GdkGC
		Bg_gc               *GdkGC
		Cursor_drag         *GdkCursor
		X_drag              int
		Focus_row           int
		Focus_header_column int
		Anchor              int
		Anchor_state        GtkStateType
		Drag_pos            int
		Htimer              int
		Vtimer              int
		Sort_type           GtkSortType
		Compare             GtkCListCompareFunc
		Sort_column         int
		Drag_highlight_row  int
		Drag_highlight_pos  GtkCListDragPos
	}

	GtkCListColumn struct {
		Title         *Gchar
		Area          GdkRectangle
		Button        *GtkWidget
		Window        *GdkWindow
		Width         int
		Min_width     int
		Max_width     int
		Justification GtkJustification
		Bits          uint
		// visible : 1
		// width_set : 1
		// resizeable : 1
		// auto_resize : 1
		// button_passive : 1
	}

	GtkCListCellInfo struct {
		Row    int
		Column int
	}

	GtkAdjustment struct { // REMOVE
		Parent_instance GtkObject
		Lower           float64
		Upper           float64
		Value           float64
		Step_increment  float64
		Page_increment  float64
		Page_size       float64
	}

	GtkCListRow struct {
		Cell       *_GtkCell
		State      GtkStateType
		Foreground GdkColor
		Background GdkColor
		Style      *GtkStyle
		Data       Gpointer
		Destroy    GDestroyNotify
		Bits       uint
		// fg_set : 1
		// bg_set : 1
		// selectable : 1
	}

	_GtkCell struct { // TODO(t):fix
		Type       GtkCellType
		Vertical   int16
		Horizontal int16
		Style      *GtkStyle
		/* union
		   text  *Gchar;
		   struct {
				pixmap  *GdkPixmap;
				mask  *GdkBitmap;
		   } pm;
		   struct {
				text  *Gchar;
				spacing  uint8;
				pixmap  *GdkPixmap;
				mask  *GdkBitmap;
		   } pt;
		   widget  *GtkWidget;
		*/
	}

	GtkCombo struct {
		Hbox            GtkHBox
		Entry           *GtkWidget
		Button          *GtkWidget
		Popup           *GtkWidget
		Popwin          *GtkWidget
		List            *GtkWidget
		Entry_change_id uint
		List_change_id  uint
		Bits            uint
		// value_in_list:1
		// ok_if_empty:1
		// case_sensitive:1
		// use_arrows:1
		// use_arrows_always:1
		Current_button uint16
		Activate_id    uint
	}

	GtkHBox struct {
		Box GtkBox
	}

	GtkBox struct { // REMOVE
		Container   GtkContainer
		Children    *GList
		Spacing     int16
		Homogeneous uint // : 1
	}

	GtkCTree struct {
		Clist        GtkCList
		Lines_gc     *GdkGC
		Tree_indent  int
		Tree_spacing int
		Tree_column  int
		Bits         uint
		// line_style : 2
		// expander_style : 2
		// show_stub : 1
		Drag_compare GtkCTreeCompareDragFunc
	}

	GtkCTreeNode struct {
		List GList
	}

	GtkCTreeRow struct {
		Row           GtkCListRow
		Parent        *GtkCTreeNode
		Sibling       *GtkCTreeNode
		Children      *GtkCTreeNode
		Pixmap_closed *GdkPixmap
		Mask_closed   *GdkBitmap
		Pixmap_opened *GdkPixmap
		Mask_opened   *GdkBitmap
		Level         uint16
		Bits          uint
		// is_leaf : 1
		// expanded : 1
	}

	GtkTextBuffer struct {
		Parent_instance            GObject
		Tag_table                  *GtkTextTagTable
		Btree                      *GtkTextBTree
		Clipboard_contents_buffers *GSList
		Selection_clipboards       *GSList
		Log_attr_cache             *GtkTextLogAttrCache
		User_action_count          uint
		Bits                       uint
		// modified : 1
		// has_selection : 1
	}

	GtkTextTagTable struct {
		Parent_instance GObject
		Hash            *GHashTable
		Anonymous       *GSList
		Anon_count      int
		Buffers         *GSList
	}

	GtkTextView struct {
		Parent_instance        GtkContainer
		Layout                 *GtkTextLayout
		Buffer                 *GtkTextBuffer
		Selection_drag_handler uint
		Scroll_timeout         uint
		Pixels_above_lines     int
		Pixels_below_lines     int
		Pixels_inside_wrap     int
		Wrap_mode              GtkWrapMode
		Justify                GtkJustification
		Left_margin            int
		Right_margin           int
		Indent                 int
		Tabs                   *PangoTabArray
		Bits                   uint
		// editable : 1
		// overwrite_mode : 1
		// cursor_visible : 1
		// need_im_reset : 1
		// accepts_tab : 1
		// width_changed : 1
		// onscreen_validated : 1
		// mouse_cursor_obscured : 1
		Text_window                 *GtkTextWindow
		Left_window                 *GtkTextWindow
		Right_window                *GtkTextWindow
		Top_window                  *GtkTextWindow
		Bottom_window               *GtkTextWindow
		Hadjustment                 *GtkAdjustment
		Vadjustment                 *GtkAdjustment
		Xoffset                     int
		Yoffset                     int
		Width                       int
		Height                      int
		Virtual_cursor_x            int
		Virtual_cursor_y            int
		First_para_mark             *GtkTextMark
		First_para_pixels           int
		Dnd_mark                    *GtkTextMark
		Blink_timeout               uint
		First_validate_idle         uint
		Incremental_validate_idle   uint
		Im_context                  *GtkIMContext
		Popup_menu                  *GtkWidget
		Drag_start_x                int
		Drag_start_y                int
		Children                    *GSList
		Pending_scroll              *GtkTextPendingScroll
		Pending_place_cursor_button int
	}

	GtkTextMark struct {
		Parent_instance GObject
		Segment         Gpointer
	}

	GtkTextIter struct {
		Dummy1  Gpointer
		Dummy2  Gpointer
		Dummy3  int
		Dummy4  int
		Dummy5  int
		Dummy6  int
		Dummy7  int
		Dummy8  int
		Dummy9  Gpointer
		Dummy10 Gpointer
		Dummy11 int
		Dummy12 int
		Dummy13 int
		Dummy14 Gpointer
	}

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

	GtkTextChildAnchor struct {
		Parent_instance GObject
		Segment         Gpointer
	}

	GtkTextAttributes struct {
		Refcount           uint
		Appearance         GtkTextAppearance
		Justification      GtkJustification
		Direction          GtkTextDirection
		Font               *PangoFontDescription
		Font_scale         float64
		Left_margin        int
		Indent             int
		Right_margin       int
		Pixels_above_lines int
		Pixels_below_lines int
		Pixels_inside_wrap int
		Tabs               *PangoTabArray
		Wrap_mode          GtkWrapMode
		Language           *PangoLanguage
		Pg_bg_color        *GdkColor
		Bits               uint
		// invisible : 1
		// bg_full_height : 1
		// editable : 1
		// realized : 1
		// pad1 : 1
		// pad2 : 1
		// pad3 : 1
		// pad4 : 1
	}

	GtkTextAppearance struct {
		Bg_color   GdkColor
		Fg_color   GdkColor
		Bg_stipple *GdkBitmap
		Fg_stipple *GdkBitmap
		Rise       int
		Padding1   Gpointer
		Bits       uint
		// underline : 4
		// strikethrough : 1
		// draw_bg : 1
		// inside_selection : 1
		// is_text : 1
		// pad1 : 1
		// pad2 : 1
		// pad3 : 1
		// pad4 : 1
	}

	GtkPixmap struct {
		Misc               GtkMisc
		Pixmap             *GdkPixmap
		Mask               *GdkBitmap
		Pixmap_insensitive *GdkPixmap
		Build_insensitive  uint // : 1
	}

	GtkToolbar struct {
		Container    GtkContainer
		Num_children int
		Children     *GList
		Orientation  GtkOrientation
		Style        GtkToolbarStyle
		Icon_size    GtkIconSize
		Tooltips     *GtkTooltips
		Button_maxw  int
		Button_maxh  int
		_, _         uint
		Bits         uint
		// style_set : 1
		// icon_size_set : 1
	}

	GtkTooltips struct {
		Parent_instance  GtkObject
		Tip_window       *GtkWidget
		Tip_label        *GtkWidget
		Active_tips_data *GtkTooltipsData
		Tips_data_list   *GList
		Bits, Bits2      uint //TODO(t): 33 bits Alignment/size?
		// delay : 30
		// enabled : 1
		// have_grab : 1
		// use_sticky_delay : 1
		Timer_tag    int
		Last_popdown GTimeVal
	}

	GtkTooltipsData struct {
		Tooltips    *GtkTooltips
		Widget      *GtkWidget
		Tip_text    *Gchar
		Tip_private *Gchar
	}

	GtkToolItem struct {
		Parent GtkBin
		Priv   *GtkToolItemPrivate
	}

	GtkToolItemGroup struct {
		Parent_instance GtkContainer
		Priv            *GtkToolItemGroupPrivate
	}

	GtkToolPalette struct {
		Parent_instance GtkContainer
		Priv            *GtkToolPalettePrivate
	}

	GtkSelectionData struct {
		Selection GdkAtom
		Target    GdkAtom
		Type      GdkAtom
		Format    int
		Data      *Guchar
		Length    int
		Display   *GdkDisplay
	}

	GdkDisplay struct {
		Parent_instance   GObject
		Queued_events     *GList
		Queued_tail       *GList
		Button_click_time [2]GUint32
		Button_window     [2]*GdkWindow
		Button_number     [2]int
		Double_click_time uint
		Core_pointer      *GdkDevice
		Pointer_hooks     *GdkDisplayPointerHooks
		Bits              uint
		// closed : 1
		// ignore_core_events : 1
		Double_click_distance uint
		Button_x              [2]int
		Button_y              [2]int
		Pointer_grabs         *GList
		Keyboard_grab         GdkKeyboardGrabInfo
		Pointer_info          GdkPointerWindowInfo
		Last_event_time       GUint32
	}

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

	GdkDisplayPointerHooks struct {
		Get_pointer func(
			display *GdkDisplay,
			screen **GdkScreen,
			x *int,
			y *int,
			mask *GdkModifierType)
		Window_get_pointer func(
			display *GdkDisplay,
			window *GdkWindow,
			x *int,
			y *int,
			mask *GdkModifierType) *GdkWindow
		Window_at_pointer func(
			display *GdkDisplay,
			win_x *int,
			win_y *int) *GdkWindow
	}

	GdkDeviceAxis struct {
		Use GdkAxisUse
		Min float64
		Max float64
	}

	GtkTargetEntry struct {
		Target *Gchar
		Flags  uint
		Info   uint
	}

	GtkSizeGroup struct {
		Parent_instance GObject
		Widgets         *GSList
		Mode            uint8
		Bits            uint
		// have_width : 1
		// have_height : 1
		// ignore_hidden : 1
		Requisition GtkRequisition
	}

	GtkSpinButton struct {
		Entry         GtkEntry
		Adjustment    *GtkAdjustment
		Panel         *GdkWindow
		Timer         GUint32
		Climb_rate    float64
		Timer_step    float64
		Update_policy GtkSpinButtonUpdatePolicy
		Bits          uint
		// in_child : 2
		// click_child : 2
		// button : 2
		// need_timer : 1
		// timer_calls : 3
		// digits : 10
		// numeric : 1
		// wrap : 1
		// snap_to_ticks : 1
	}

	GtkEntry struct {
		Widget GtkWidget
		Text   *Gchar
		Bits   uint
		// editable : 1
		// visible : 1
		// overwrite_mode : 1
		// in_drag : 1
		Text_length     uint16
		Text_max_length uint16
		Text_area       *GdkWindow
		Im_context      *GtkIMContext
		Popup_menu      *GtkWidget
		Current_pos     int
		Selection_bound int
		Cached_layout   *PangoLayout
		Bits2           uint
		// cache_includes_preedit : 1
		// need_im_reset : 1
		// has_frame : 1
		// activates_default : 1
		// cursor_visible : 1
		// in_click : 1
		// is_cell_renderer : 1
		// editing_canceled : 1
		// mouse_cursor_obscured : 1
		// select_words : 1
		// select_lines : 1
		// resolved_dir : 4
		// truncate_multiline : 1
		Button         uint
		Blink_timeout  uint
		Recompute_idle uint
		Scroll_offset  int
		Ascent         int
		Descent        int
		X_text_size    uint16
		X_n_bytes      uint16
		Preedit_length uint16
		Preedit_cursor uint16
		Dnd_position   int
		Drag_start_x   int
		Drag_start_y   int
		Invisible_char Gunichar
		Width_chars    int
	}

	GtkTreeModelSort struct {
		Parent               GObject
		Root                 Gpointer
		Stamp                int
		Child_flags          uint
		Child_model          *GtkTreeModel
		Zero_ref_count       int
		Sort_list            *GList
		Sort_column_id       int
		Order                GtkSortType
		Default_sort_func    GtkTreeIterCompareFunc
		Default_sort_data    Gpointer
		Default_sort_destroy GDestroyNotify
		Changed_id           uint
		Inserted_id          uint
		Has_child_toggled_id uint
		Deleted_id           uint
		Reordered_id         uint
	}

	GtkTreeIter struct {
		Stamp      int
		User_data  Gpointer
		User_data2 Gpointer
		User_data3 Gpointer
	}

	GtkTreeSelection struct {
		Parent    GObject
		Tree_view *GtkTreeView
		Type      GtkSelectionMode
		User_func GtkTreeSelectionFunc
		User_data Gpointer
		Destroy   GDestroyNotify
	}

	GtkTreeView struct {
		Parent GtkContainer
		Priv   *GtkTreeViewPrivate
	}

	GtkTreeStore struct {
		Parent               GObject
		Stamp                int
		Root                 Gpointer
		Last                 Gpointer
		N_columns            int
		Sort_column_id       int
		Sort_list            *GList
		Order                GtkSortType
		Column_headers       *GType
		Default_sort_func    GtkTreeIterCompareFunc
		Default_sort_data    Gpointer
		Default_sort_destroy GDestroyNotify
		Columns_dirty        uint // : 1
	}

	GtkUIManager struct {
		Parent       GObject
		Private_data *GtkUIManagerPrivate
	}

	GtkActionGroup struct { // REMOVE
		Parent       GObject
		Private_data *GtkActionGroupPrivate
	}

	GtkAction struct { // REMOVE
		Object       GObject
		Private_data *GtkActionPrivate
	}

	GtkHSV struct {
		Parent_instance GtkWidget
		Priv            Gpointer
	}

	GtkIconFactory struct {
		Parent_instance GObject
		Icons           *GHashTable
	}

	GtkSettings struct {
		Parent_instance GObject
		Queued_settings *GData
		Property_values *GtkSettingsPropertyValue
		Rc_context      *GtkRcContext
		Screen          *GdkScreen
	}

	GtkIconTheme struct {
		Parent_instance GObject
		Priv            *GtkIconThemePrivate
	}

	GdkEvent struct {
		Type GdkEventType
		// Union
		// any  GdkEventAny;
		// expose  GdkEventExpose;
		// no_expose  GdkEventNoExpose;
		// visibility  GdkEventVisibility;
		// motion  GdkEventMotion;
		// button  GdkEventButton;
		// scroll  GdkEventScroll;
		// key  GdkEventKey;
		// crossing  GdkEventCrossing;
		// focus_change  GdkEventFocus;
		// configure  GdkEventConfigure;
		// property  GdkEventProperty;
		// selection  GdkEventSelection;
		// owner_change  GdkEventOwnerChange;
		// proximity  GdkEventProximity;
		// client  GdkEventClient;
		// dnd  GdkEventDND;
		// window_state  GdkEventWindowState;
		// setting  GdkEventSetting;
		// grab_broken  GdkEventGrabBroken;
	}

	GtkIconView struct {
		Parent GtkContainer
		Priv   *GtkIconViewPrivate
	}

	GtkCellRenderer struct { // REMOVE
		Parent GtkObject
		Xalign float32
		Yalign float32
		Width  int
		Height int
		Xpad   uint16
		Ypad   uint16
		Bits   uint
		// mode : 2;
		// visible : 1
		// is_expander : 1
		// is_expanded : 1
		// cell_background_set : 1
		// sensitive : 1
		// editing : 1
	}

	GtkImageMenuItem struct {
		Menu_item GtkMenuItem
		Image     *GtkWidget
	}

	GtkMenuItem struct {
		Item              GtkItem
		Submenu           *GtkWidget
		Event_window      *GdkWindow
		Toggle_size       uint16
		Accelerator_width uint16
		Accel_path        *Gchar
		Bits              uint
		// show_submenu_indicator : 1
		// submenu_placement : 1
		// submenu_direction : 1
		// right_justify: 1
		// timer_from_keypress : 1
		// from_menubar : 1
		Timer uint
	}

	GtkIMContextSimple struct {
		Object              GtkIMContext
		Tables              *GSList
		Compose_buffer      [7 + 1]uint
		Tentative_match     Gunichar
		Tentative_match_len int
		Bits                uint
		// in_hex_sequence : 1
		// modifiers_dropped : 1
	}

	GtkIMMulticontext struct {
		Object     GtkIMContext
		Slave      *GtkIMContext
		Priv       *GtkIMMulticontextPrivate
		Context_id *Gchar
	}

	GtkMenuShell struct {
		Container         GtkContainer
		Children          *GList
		Active_menu_item  *GtkWidget
		Parent_menu_shell *GtkWidget
		Button            uint
		Activate_time     GUint32
		Bits              uint
		// active : 1
		// have_grab : 1
		// have_xgrab : 1
		// ignore_leave : 1
		// menu_flag : 1
		// ignore_enter : 1
		// keyboard_mode : 1
	}

	GtkInfoBar struct {
		Parent GtkHBox
		Priv   *GtkInfoBarPrivate
	}

	GtkInvisible struct {
		Widget             GtkWidget
		Has_user_ref_count Gboolean
		Screen             *GdkScreen
	}

	GtkLayout struct {
		Container    GtkContainer
		Children     *GList
		Width        uint
		Height       uint
		Hadjustment  *GtkAdjustment
		Vadjustment  *GtkAdjustment
		Bin_window   *GdkWindow
		Visibility   GdkVisibilityState
		Scroll_x     int
		Scroll_y     int
		Freeze_count uint
	}

	GtkLinkButton struct {
		Parent_instance GtkButton
		Priv            *GtkLinkButtonPrivate
	}

	GtkArg struct {
		Type GType
		Name *Gchar
		// Union
		// Gchar char_data;
		// uchar_data  Guchar;
		// bool_data  Gboolean;
		// int_data  int;
		// uint_data  uint;
		// long_data  glong;
		// ulong_data  Gulong;
		// float_data  float32;
		// Double_data  float64;
		// string_data  *Gchar;
		// object_data  *GtkObject;
		// pointer_data  Gpointer;
		// signal_data struct{f GCallback; d Gpointer}
	}

	GSignalInvocationHint struct {
		Signal_id uint
		Detail    GQuark
		Run_type  GSignalFlags
	}

	GtkMenuBar struct {
		Menu_shell GtkMenuShell
	}

	GtkToolButton struct {
		Parent GtkToolItem
		Priv   *GtkToolButtonPrivate
	}

	GtkMenuToolButton struct {
		Parent GtkToolButton
		Priv   *GtkMenuToolButtonPrivate
	}

	GtkMessageDialog struct {
		Parent_instance GtkDialog
		Image           *GtkWidget
		Label           *GtkWidget
	}

	GMountOperation struct {
		Parent_instance GObject
		Priv            *GMountOperationPrivate
	}

	GtkMountOperation struct {
		Parent_instance GMountOperation
		Priv            *GtkMountOperationPrivate
	}

	GtkNotebook struct {
		Container    GtkContainer
		Cur_page     *GtkNotebookPage
		Children     *GList
		First_tab    *GList
		Focus_tab    *GList
		Menu         *GtkWidget
		Event_window *GdkWindow
		Timer        GUint32
		Tab_hborder  uint16
		Tab_vborder  uint16
		Bits         uint
		// show_tabs : 1
		// homogeneous : 1
		// show_border : 1
		// tab_pos : 2
		// scrollable : 1
		// in_child : 3
		// click_child : 3
		// button : 2
		// need_timer : 1
		// child_has_focus : 1
		// have_visible_child : 1
		// focus_out : 1
		// has_before_previous : 1
		// has_before_next : 1
		// has_after_previous : 1
		// has_after_next : 1
	}

	GtkOffscreenWindow struct {
		Parent_object GtkWindow
	}

	GtkSocket struct {
		Container      GtkContainer
		Request_width  uint16
		Request_height uint16
		Current_width  uint16
		Current_height uint16
		Plug_window    *GdkWindow
		Plug_widget    *GtkWidget
		Xembed_version int16
		Bits           uint
		// same_app : 1
		// focus_in : 1
		// have_size : 1
		// need_map : 1
		// is_mapped : 1
		// active : 1
		Accel_group *GtkAccelGroup
		Toplevel    *GtkWidget
	}

	GtkPlug struct {
		Window          GtkWindow
		Socket_window   *GdkWindow
		Modality_window *GtkWidget
		Modality_group  *GtkWindowGroup
		Grabbed_keys    *GHashTable
		Same_app        uint // : 1
	}

	GtkPageRange struct {
		Start int
		End   int
	}

	GtkPrintOperation struct {
		Parent_instance GObject
		Priv            *GtkPrintOperationPrivate
	}

	GtkProgress struct {
		Widget           GtkWidget
		Adjustment       *GtkAdjustment
		Offscreen_pixmap *GdkPixmap
		Format           *Gchar
		X_align          float32
		Y_align          float32
		Bits             uint
		// show_text : 1;
		// activity_mode : 1;
		// use_text_format : 1
	}

	GtkProgressBar struct {
		Progress        GtkProgress
		Bar_style       GtkProgressBarStyle
		Orientation     GtkProgressBarOrientation
		Blocks          uint
		In_block        int
		Activity_pos    int
		Activity_step   uint
		Activity_blocks uint
		Pulse_fraction  float64
		Bits            uint
		// activity_dir : 1
		// ellipsize : 3
		// dirty : 1
	}

	GtkToggleAction struct {
		Parent       GtkAction
		Private_data *GtkToggleActionPrivate
	}

	GtkRadioAction struct {
		Parent       GtkToggleAction
		Private_data *GtkRadioActionPrivate
	}

	GtkRadioButton struct {
		Check_button GtkCheckButton
		Group        *GSList
	}

	GtkRadioMenuItem struct {
		Check_menu_item GtkCheckMenuItem
		Group           *GSList
	}

	GtkCheckButton struct {
		Toggle_button GtkToggleButton
	}

	GtkCheckMenuItem struct { //REMOVE
		Menu_item GtkMenuItem
		Bits      uint
		// active : 1
		// always_show_toggle : 1
		// inconsistent : 1
		// draw_as_radio : 1
	}

	GtkToggleButton struct {
		Button GtkButton
		Bits   uint
		// active : 1
		// draw_indicator : 1
		// inconsistent : 1
	}

	GtkToggleToolButton struct {
		Parent GtkToolButton
		Priv   *GtkToggleToolButtonPrivate
	}

	GtkRadioToolButton struct {
		Parent GtkToggleToolButton
	}

	GtkRecentManager struct {
		Parent_instance GObject
		Priv            *GtkRecentManagerPrivate
	}

	GtkRecentData struct {
		Display_name *Gchar
		Description  *Gchar
		Mime_type    *Gchar
		App_name     *Gchar
		App_exec     *Gchar
		Groups       **Gchar
		Is_private   Gboolean
	}

	GtkRecentAction struct {
		Parent_instance GtkAction
		Priv            *GtkRecentActionPrivate
	}

	GtkScaleButton struct {
		Parent       GtkButton
		Plus_button  *GtkWidget
		Minus_button *GtkWidget
		Priv         *GtkScaleButtonPrivate
	}

	GtkRecentFilterInfo struct {
		Contains     GtkRecentFilterFlags
		Uri          *Gchar
		Display_name *Gchar
		Mime_type    *Gchar
		Applications **Gchar
		Groups       **Gchar
		Age          int
	}

	GtkViewport struct {
		Bin         GtkBin
		Shadow_type GtkShadowType
		View_window *GdkWindow
		Bin_window  *GdkWindow
		Hadjustment *GtkAdjustment
		Vadjustment *GtkAdjustment
	}

	GtkScrolledWindow struct {
		Container  GtkBin
		Hscrollbar *GtkWidget
		Vscrollbar *GtkWidget
		Bits       uint
		// hscrollbar_policy : 2
		// vscrollbar_policy : 2
		// hscrollbar_visible : 1
		// vscrollbar_visible : 1
		// window_placement : 2
		// focus_out : 1
		Shadow_type uint16
	}

	GtkRecentChooserMenu struct {
		Parent_instance GtkMenu
		Priv            *GtkRecentChooserMenuPrivate
	}

	GtkMenu struct {
		Menu_shell           GtkMenuShell
		Parent_menu_item     *GtkWidget
		Old_active_menu_item *GtkWidget
		Accel_group          *GtkAccelGroup
		Accel_path           *Gchar
		Position_func        GtkMenuPositionFunc
		Position_func_data   Gpointer
		Toggle_size          uint
		Toplevel             *GtkWidget
		Tearoff_window       *GtkWidget
		Tearoff_hbox         *GtkWidget
		Tearoff_scrollbar    *GtkWidget
		Tearoff_adjustment   *GtkAdjustment
		View_window          *GdkWindow
		Bin_window           *GdkWindow
		Scroll_offset        int
		Saved_scroll_offset  int
		Scroll_step          int
		Timeout_id           uint
		Navigation_region    *GdkRegion
		Navigation_timeout   uint
		Bits                 uint
		// needs_destruction_ref_count : 1
		// torn_off : 1
		// tearoff_active : 1
		// scroll_fast : 1
		// upper_arrow_visible : 1
		// lower_arrow_visible : 1
		// upper_arrow_prelight : 1
		// lower_arrow_prelight : 1
	}

	GtkSeparatorToolItem struct {
		Parent GtkToolItem
		Priv   *GtkSeparatorToolItemPrivate
	}

	GtkStatusbar struct {
		Parent_widget   GtkHBox
		Frame           *GtkWidget
		Label           *GtkWidget
		Messages        *GSList
		Keys            *GSList
		Seq_context_id  uint
		Seq_message_id  uint
		Grip_window     *GdkWindow
		Has_resize_grip uint // : 1
	}

	GtkSpinner struct {
		Parent GtkDrawingArea
		Priv   *GtkSpinnerPrivate
	}

	GtkStatusIcon struct {
		Parent_instance GObject
		Priv            *GtkStatusIconPrivate
	}

	GtkStockItem struct {
		Stock_id           *Gchar
		Label              *Gchar
		Modifier           GdkModifierType
		Keyval             uint
		Translation_domain *Gchar
	}

	GtkTable struct {
		Container      GtkContainer
		Children       *GList
		Rows           *GtkTableRowCol
		Cols           *GtkTableRowCol
		Nrows          uint16
		Ncols          uint16
		Column_spacing uint16
		Row_spacing    uint16
		Bits           uint
		// homogeneous : 1
	}

	GtkTableRowCol struct {
		Requisition uint16
		Allocation  uint16
		Spacing     uint16
		Bits        uint
		// need_expand : 1
		// need_shrink : 1
		// expand : 1
		// shrink : 1
		// empty : 1
	}

	GtkTextTag struct {
		Parent_instance GObject
		Table           *GtkTextTagTable
		Name            *Char
		Priority        int
		Values          *GtkTextAttributes
		Bits            uint
		// bg_color_set : 1
		// bg_stipple_set : 1
		// fg_color_set : 1
		// scale_set : 1
		// fg_stipple_set : 1
		// justification_set : 1
		// left_margin_set : 1
		// indent_set : 1
		// rise_set : 1
		// strikethrough_set : 1
		// right_margin_set : 1
		// pixels_above_lines_set : 1
		// pixels_below_lines_set : 1
		// pixels_inside_wrap_set : 1
		// tabs_set : 1
		// underline_set : 1
		// wrap_mode_set : 1
		// bg_full_height_set : 1
		// invisible_set : 1
		// editable_set : 1
		// language_set : 1
		// pg_bg_color_set : 1
		// accumulative_margin : 1
		// pad1 : 1
	}

	GtkTargetList struct {
		List      *GList
		Ref_count uint
	}

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

	GtkAboutDialog struct { // REMOVE
		Parent_instance GtkDialog
		Private_data    Gpointer
	}

	GtkAccelLabel struct { //REMOVE
		Label              GtkLabel
		_                  uint
		Accel_padding      uint
		Accel_widget       *GtkWidget
		Accel_closure      *GClosure
		Accel_group        *GtkAccelGroup
		Accel_string       *Gchar
		Accel_string_width uint16
	}

	GtkAccelLabelClass struct {
		Parent_class     GtkLabelClass
		Signal_quote1    *Gchar
		Signal_quote2    *Gchar
		Mod_name_shift   *Gchar
		Mod_name_control *Gchar
		Mod_name_alt     *Gchar
		Mod_separator    *Gchar
		Accel_seperator  *Gchar
		Latin1_to_char   uint // : 1
		_, _, _, _       func()
	}

	GtkLabelClass struct {
		Parent_class GtkMiscClass
		Move_cursor  func(
			label *GtkLabel,
			step GtkMovementStep,
			count int,
			extend_selection Gboolean)
		Copy_clipboard func(
			label *GtkLabel)
		Populate_popup func(
			label *GtkLabel,
			menu *GtkMenu)
		Activate_link func(
			label *GtkLabel,
			uri *Gchar) Gboolean
		_, _, _ func()
	}

	GtkMiscClass struct {
		Parent_class GtkWidgetClass
	}

	GtkWidgetClass struct {
		Parent_class                      GtkObjectClass
		Activate_signal                   uint
		Set_scroll_adjustments_signal     uint
		Dispatch_child_properties_changed func(
			widget *GtkWidget,
			n_pspecs uint,
			pspecs **GParamSpec)
		Show         func(widget *GtkWidget)
		Show_all     func(widget *GtkWidget)
		Hide         func(widget *GtkWidget)
		Hide_all     func(widget *GtkWidget)
		Map          func(widget *GtkWidget)
		Unmap        func(widget *GtkWidget)
		Realize      func(widget *GtkWidget)
		Unrealize    func(widget *GtkWidget)
		Size_request func(widget *GtkWidget,
			requisition *GtkRequisition)
		Size_allocate func(widget *GtkWidget,
			allocation *GtkAllocation)
		State_changed func(widget *GtkWidget,
			previous_state GtkStateType)
		Parent_set func(widget *GtkWidget,
			previous_parent *GtkWidget)
		Hierarchy_changed func(widget *GtkWidget,
			previous_toplevel *GtkWidget)
		Style_set func(widget *GtkWidget,
			previous_style *GtkStyle)
		Direction_changed func(widget *GtkWidget,
			previous_direction GtkTextDirection)
		Grab_notify func(widget *GtkWidget,
			was_grabbed Gboolean)
		Child_notify func(widget *GtkWidget,
			pspec *GParamSpec)
		Mnemonic_activate func(widget *GtkWidget,
			group_cycling Gboolean) Gboolean
		Grab_focus func(widget *GtkWidget)
		Focus      func(widget *GtkWidget,
			direction GtkDirectionType) Gboolean
		Event func(widget *GtkWidget,
			event *GdkEvent) Gboolean
		Button_press_event func(widget *GtkWidget,
			event *GdkEventButton) Gboolean
		Button_release_event func(widget *GtkWidget,
			event *GdkEventButton) Gboolean
		Scroll_event func(widget *GtkWidget,
			event *GdkEventScroll) Gboolean
		Motion_notify_event func(widget *GtkWidget,
			event *GdkEventMotion) Gboolean
		Delete_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		Destroy_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		Expose_event func(widget *GtkWidget,
			event *GdkEventExpose) Gboolean
		Key_press_event func(widget *GtkWidget,
			event *GdkEventKey) Gboolean
		Key_release_event func(widget *GtkWidget,
			event *GdkEventKey) Gboolean
		Enter_notify_event func(widget *GtkWidget,
			event *GdkEventCrossing) Gboolean
		Leave_notify_event func(widget *GtkWidget,
			event *GdkEventCrossing) Gboolean
		Configure_event func(widget *GtkWidget,
			event *GdkEventConfigure) Gboolean
		Focus_in_event func(widget *GtkWidget,
			event *GdkEventFocus) Gboolean
		Focus_out_event func(widget *GtkWidget,
			event *GdkEventFocus) Gboolean
		Map_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		Unmap_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		Property_notify_event func(widget *GtkWidget,
			event *GdkEventProperty) Gboolean
		Selection_clear_event func(widget *GtkWidget,
			event *GdkEventSelection) Gboolean
		Selection_request_event func(widget *GtkWidget,
			event *GdkEventSelection) Gboolean
		Selection_notify_event func(widget *GtkWidget,
			event *GdkEventSelection) Gboolean
		Proximity_in_event func(widget *GtkWidget,
			event *GdkEventProximity) Gboolean
		Proximity_out_event func(widget *GtkWidget,
			event *GdkEventProximity) Gboolean
		Visibility_notify_event func(widget *GtkWidget,
			event *GdkEventVisibility) Gboolean
		Client_event func(widget *GtkWidget,
			event *GdkEventClient) Gboolean
		No_expose_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		Window_state_event func(widget *GtkWidget,
			event *GdkEventWindowState) Gboolean
		Selection_get func(widget *GtkWidget,
			selection_data *GtkSelectionData,
			info, time_ uint)
		Selection_received func(widget *GtkWidget,
			selection_data *GtkSelectionData,
			time_ uint)
		Drag_begin func(widget *GtkWidget,
			context *GdkDragContext)
		Drag_end func(widget *GtkWidget,
			context *GdkDragContext)
		Drag_data_get func(widget *GtkWidget,
			context *GdkDragContext,
			selection_data *GtkSelectionData,
			info, time_ uint)
		Drag_data_delete func(widget *GtkWidget,
			context *GdkDragContext)
		Drag_leave func(widget *GtkWidget,
			context *GdkDragContext,
			time_ uint)
		Drag_motion func(widget *GtkWidget,
			context *GdkDragContext,
			x, y int, time_ uint) Gboolean
		Drag_drop func(widget *GtkWidget,
			context *GdkDragContext,
			x, y int, time_ uint) Gboolean
		Drag_data_received func(widget *GtkWidget,
			context *GdkDragContext,
			x, y int,
			selection_data *GtkSelectionData,
			info, time_ uint)
		Popup_menu func(widget *GtkWidget) Gboolean
		Show_help  func(widget *GtkWidget,
			help_type GtkWidgetHelpType) Gboolean
		Get_accessible func(widget *GtkWidget) *AtkObject
		Screen_changed func(widget *GtkWidget,
			previous_screen *GdkScreen) Gboolean
		Can_activate_accel func(widget *GtkWidget,
			signal_id uint) Gboolean
		Grab_broken_event func(widget *GtkWidget,
			event *GdkEventGrabBroken) Gboolean
		Composited_changed func(widget *GtkWidget)
		Query_tooltip      func(widget *GtkWidget,
			x, y int, keyboard_tooltip Gboolean,
			tooltip *GtkTooltip) Gboolean
		_, _, _ func()
	}

	GtkObjectClass struct {
		Parent_class GInitiallyUnownedClass
		Set_arg      func(object *GtkObject,
			arg *GtkArg,
			arg_id uint)
		Get_arg func(object *GtkObject,
			arg *GtkArg,
			arg_id uint)
		Destroy func(object *GtkObject)
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

	GdkEventGrabBroken struct {
		Yype        GdkEventType
		Window      *GdkWindow
		Send_event  int8
		Keyboard    Gboolean
		Implicit    Gboolean
		Grab_window *GdkWindow
	}

	GtkAlignment struct {
		Bin    GtkBin
		Xalign float32
		Yalign float32
		Xscale float32
		Yscale float32
	}

	GtkAssistant struct { //REMOVE
		Parent  GtkWindow
		Cancel  *GtkWidget
		Forward *GtkWidget
		Back    *GtkWidget
		Apply   *GtkWidget
		Close   *GtkWidget
		Last    *GtkWidget
		Priv    *GtkAssistantPrivate
	}

	GtkBindingSet struct {
		Set_name            *Gchar
		Priority            int
		Widget_path_pspecs  *GSList
		Widget_class_pspecs *GSList
		Class_branch_pspecs *GSList
		Entries             *GtkBindingEntry
		Current             *GtkBindingEntry
		Parsed              uint // : 1
	}

	GtkBuilder struct { //REMOVE
		Parent_instance GObject
		Priv            *GtkBuilderPrivate
	}

	GtkButtonBox struct { // REMOVE
		Box              GtkBox
		Child_min_width  int
		Child_min_height int
		Child_ipad_x     int
		Child_ipad_y     int
		Layout_style     GtkButtonBoxStyle
	}

	GtkCalendar struct { // REMOVE
		Widget            GtkWidget
		Header_style      *GtkStyle
		Label_style       *GtkStyle
		Month             int
		Year              int
		Selected_day      int
		Day_month         [6][7]int
		Day               [6][7]int
		Num_marked_dates  int
		Marked_date       [31]int
		Display_flags     GtkCalendarDisplayOptions
		Marked_date_color [31]GdkColor
		Gc                *GdkGC
		Xor_gc            *GdkGC
		Focus_row         int
		Focus_col         int
		Highlight_row     int
		Highlight_col     int
		Priv              *GtkCalendarPrivate
		Grow_space        [32]Gchar
		_, _, _, _        func()
	}

	GtkColorButton struct {
		Button GtkButton
		Priv   *GtkColorButtonPrivate
	}

	GtkColorSelection struct {
		Parent_instance GtkVBox
		Private_data    Gpointer
	}

	GtkColorSelectionDialog struct {
		Parent_instance GtkDialog
		Colorsel        *GtkWidget
		Ok_button       *GtkWidget
		Cancel_button   *GtkWidget
		Help_button     *GtkWidget
	}

	GtkComboBox struct { // REMOVE
		Parent_instance GtkBin
		Priv            *GtkComboBoxPrivate
	}

	GtkComboBoxEntry struct { // REMOVE
		Parent_instance GtkComboBox
		Priv            *GtkComboBoxEntryPrivate
	}

	GtkComboBoxText struct { //REMOVE
		Parent_instance GtkComboBox
		Priv            *GtkComboBoxTextPrivate
	}

	GtkContainerClass struct {
		Parent_class GtkWidgetClass
		Add          func(container *GtkContainer,
			widget *GtkWidget)
		Remove func(container *GtkContainer,
			widget *GtkWidget)
		Check_resize func(container *GtkContainer)
		Forall       func(container *GtkContainer,
			include_internals Gboolean,
			callback GtkCallback,
			callback_data Gpointer)
		Set_focus_child func(container *GtkContainer,
			widget *GtkWidget)
		Child_type     func(container *GtkContainer) GType
		Composite_name func(container *GtkContainer,
			child *GtkWidget) *Gchar
		Set_child_property func(container *GtkContainer,
			child *GtkWidget,
			property_id uint,
			value *GValue,
			pspec *GParamSpec)
		Get_child_property func(container *GtkContainer,
			child *GtkWidget,
			property_id uint,
			value *GValue,
			pspec *GParamSpec)
		_, _, _, _ func()
	}

	GtkEntryBuffer struct {
		Parent_instance GObject
		Priv            *GtkEntryBufferPrivate
	}

	GtkEntryCompletion struct {
		Parent_instance GObject
		Priv            *GtkEntryCompletionPrivate
	}

	GtkEventBox struct {
		Bin GtkBin
	}

	GtkExpander struct {
		Bin  GtkBin
		Priv *GtkExpanderPrivate
	}

	GtkFileChooserButton struct {
		Parent GtkHBox
		Priv   *GtkFileChooserButtonPrivate
	}

	GtkFileFilterInfo struct {
		Contains     GtkFileFilterFlags
		Filename     *Gchar
		Uri          *Gchar
		Display_name *Gchar
		Mime_type    *Gchar
	}

	GtkFixed struct {
		Container GtkContainer
		Children  *GList
	}

	GtkFontButton struct {
		Button GtkButton
		Priv   *GtkFontButtonPrivate
	}

	GtkFontSelection struct {
		Parent_instance  GtkVBox
		Font_entry       *GtkWidget
		Family_list      *GtkWidget
		Font_style_entry *GtkWidget
		Face_list        *GtkWidget
		Size_entry       *GtkWidget
		Size_list        *GtkWidget
		Pixels_button    *GtkWidget
		Points_button    *GtkWidget
		Filter_button    *GtkWidget
		Preview_entry    *GtkWidget
		Family           *PangoFontFamily
		Face             *PangoFontFace
		Size             int
		Font             *GdkFont
	}

	GtkFontSelectionDialog struct {
		Parent_instance GtkDialog
		Fontsel         *GtkWidget
		Main_vbox       *GtkWidget
		Action_area     *GtkWidget
		Ok_button       *GtkWidget
		Apply_button    *GtkWidget
		Cancel_button   *GtkWidget
		Dialog_width    int
		Auto_resize     Gboolean
	}

	GtkFrame struct {
		Bin              GtkBin
		Label_widget     *GtkWidget
		Shadow_type      int16
		Label_xalign     float32
		Label_yalign     float32
		Child_allocation GtkAllocation
	}

	GtkHandleBox struct {
		Bin          GtkBin
		Bin_window   *GdkWindow
		Float_window *GdkWindow
		Shadow_type  GtkShadowType
		Bits         uint
		// handle_position : 2;
		// float_window_mapped : 1;
		// child_detached : 1;
		// in_drag : 1;
		// shrink_on_detach : 1;
		// snap_edge : 3 // signed
		Deskoff_x         int
		Deskoff_y         int
		Attach_allocation GtkAllocation
		Float_allocation  GtkAllocation
	}
	/*
	   GtkImage  struct
	   {
	   misc  GtkMisc;
	   storage_type  GtkImageType;
	    union
	    {
	   pixmap  GtkImagePixmapData;
	   image  GtkImageImageData;
	   pixbuf  GtkImagePixbufData;
	   stock  GtkImageStockData;
	   icon_set  GtkImageIconSetData;
	   anim  GtkImageAnimationData;
	   name  GtkImageIconNameData;
	   gicon  GtkImageGIconData;
	    } data;
	   mask  *GdkBitmap;
	   icon_size  GtkIconSize;
	   };
	*/
	GtkListStore struct {
		Parent               GObject
		Stamp                int
		Seq                  Gpointer
		_                    Gpointer
		Sort_list            *GList
		N_columns            int
		Sort_column_id       int
		Order                GtkSortType
		Column_headers       *GType
		Length               int
		Default_sort_func    GtkTreeIterCompareFunc
		Default_sort_data    Gpointer
		Default_sort_destroy GDestroyNotify
		Columns_dirty        uint // : 1
	}

	GtkPaned struct {
		Container       GtkContainer
		Child1          *GtkWidget
		Child2          *GtkWidget
		Handle          *GdkWindow
		Xor_gc          *GdkGC
		Cursor_type     GdkCursorType
		Handle_pos      GdkRectangle
		Child1_size     int
		Last_allocation int
		Min_position    int
		Max_position    int
		Bits            uint
		// position_set : 1
		// in_drag : 1
		// child1_shrink : 1
		// child1_resize : 1
		// child2_shrink : 1
		// child2_resize : 1
		// orientation : 1
		// in_recursion : 1
		// handle_prelit : 1
		Last_child1_focus *GtkWidget
		Last_child2_focus *GtkWidget
		Priv              *GtkPanedPrivate
		Drag_pos          int
		Original_position int
	}

	GtkToggleActionEntry struct {
		Name        *Gchar
		Stock_id    *Gchar
		Label       *Gchar
		Accelerator *Gchar
		Tooltip     *Gchar
		Callback    GCallback
		Is_active   Gboolean
	}

	GtkRange struct {
		Widget        GtkWidget
		Adjustment    *GtkAdjustment
		Update_policy GtkUpdateType
		Bits          uint
		// inverted : 1
		// flippable : 1
		// has_stepper_a : 1
		// has_stepper_b : 1
		// has_stepper_c : 1
		// has_stepper_d : 1
		// need_recalc : 1
		// slider_size_fixed : 1
		Min_slider_size int
		Orientation     GtkOrientation
		Range_rect      GdkRectangle
		Slider_start    int
		Slider_end      int
		Round_digits    int
		Bits2           uint
		// trough_click_forward : 1
		// update_pending : 1
		Layout                        *GtkRangeLayout
		Timer                         *GtkRangeStepTimer
		Slide_initial_slider_position int
		Slide_initial_coordinate      int
		Update_timeout_id             uint
		Event_window                  *GdkWindow
	}

	GtkRcProperty struct {
		Type_name     GQuark
		Property_name GQuark
		Origin        *Gchar
		Value         GValue
	}

	GtkRuler struct {
		Widget        GtkWidget
		Backing_store *GdkPixmap
		Non_gr_exp_gc *GdkGC
		Metric        *GtkRulerMetric
		Xsrc          int
		Ysrc          int
		Slider_size   int
		Lower         float64
		Upper         float64
		Position      float64
		Max_size      float64
	}

	GtkScale struct {
		Range  GtkRange
		Digits int
		Bits   uint
		// draw_value : 1
		// value_pos : 2
	}

	GtkSettingsValue struct {
		Origin *Gchar
		Value  GValue
	}

	GtkTreeModelFilter struct {
		Parent GObject
		Priv   *GtkTreeModelFilterPrivate
	}

	/*
	   GtkTreeViewColumn  struct
	   {
	   parent  GtkObject;
	   tree_view  *GtkWidget;
	   button  *GtkWidget;
	   child  *GtkWidget;
	   arrow  *GtkWidget;
	   alignment  *GtkWidget;
	   window  *GdkWindow;
	   editable_widget  *GtkCellEditable;
	   xalign  float32;
	   property_changed_signal  uint;
	   spacing  int;
	   column_type  GtkTreeViewColumnSizing;
	   requested_width  int;
	   button_request  int;
	   resized_width  int;
	   width  int;
	   fixed_width  int;
	   min_width  int;
	   max_width  int;
	   drag_x  int;
	   drag_y  int;
	   title  *Gchar;
	   cell_list  *GList;
	   sort_clicked_signal  uint;
	   sort_column_changed_signal  uint;
	   sort_column_id  int;
	   sort_order  GtkSortType;
	    uint visible : 1;
	    uint resizable : 1;
	    uint clickable : 1;
	    uint dirty : 1;
	    uint show_sort_indicator : 1;
	    uint maybe_reordered : 1;
	    uint reorderable : 1;
	    uint use_resized_width : 1;
	    uint expand : 1;
	   };
	*/

	GtkTypeInfo struct {
		Type_name            *Gchar
		Object_size          uint
		Class_size           uint
		Class_init_func      GtkClassInitFunc
		Object_init_func     GtkObjectInitFunc
		Reserved_1           Gpointer
		Reserved_2           Gpointer
		Base_class_init_func GtkClassInitFunc
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

	GtkBindingEntry struct {
		Keyval      uint
		Modifiers   GdkModifierType
		Binding_set *GtkBindingSet
		Bits        uint
		// destroyed : 1
		// in_emission : 1
		// marks_unbound : 1
		Set_next  *GtkBindingEntry
		Hash_next *GtkBindingEntry
		Signals   *GtkBindingSignal
	}

	PangoColor struct {
		Red   uint16
		Green uint16
		Blue  uint16
	}

	GtkVBox struct {
		Box GtkBox
	}

	GtkRulerMetric struct {
		Metric_name     *Gchar
		Abbrev          *Gchar
		Pixels_per_unit float64
		Ruler_scale     [10]float64
		Subdivide       [5]int
	}

	GtkBindingSignal struct {
		Next        *GtkBindingSignal
		Signal_name *Gchar
		N_args      uint
		Args        *GtkBindingArg
	}

	GtkBindingArg struct {
		Arg_type GType
		//Union {
		// long_data  Glong;
		Double_data float64
		// string_data  *Gchar;
		//}
	}

	GtkRadioActionEntry struct {
		Name        *Gchar
		Stock_id    *Gchar
		Label       *Gchar
		Accelerator *Gchar
		Tooltip     *Gchar
		Value       int
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

	GdkPixbufLoader struct {
		Parent_instance GObject
		Priv            Gpointer
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

	GdkWindowAttr struct {
		Title             *Gchar
		Event_mask        int
		X, Y              int
		Width             int
		Height            int
		Wclass            GdkWindowClass
		Visual            *GdkVisual
		Colormap          *GdkColormap
		Window_type       GdkWindowType
		Cursor            *GdkCursor
		Wmclass_name      *Gchar
		Wmclass_class     *Gchar
		Override_redirect Gboolean
		Type_hint         GdkWindowTypeHint
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

	GtkTree struct {
		Container      GtkContainer
		Children       *GList
		Root_tree      *GtkTree
		Tree_owner     *GtkWidget
		Selection      *GList
		Level          uint
		Indent_value   uint
		Current_indent uint
		Bits           uint
		// selection_mode : 2
		// view_mode : 1
		// view_line : 1
	}

	GtkTreeItem struct {
		Item             GtkItem
		Subtree          *GtkWidget
		Pixmaps_box      *GtkWidget
		Plus_pix_widget  *GtkWidget
		Minus_pix_widget *GtkWidget
		Pixmaps          *GList
		Expanded         uint // : 1
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
