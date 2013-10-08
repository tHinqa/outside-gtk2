package types

//TODO(t): add constants for [...]s

type (
	Tm struct {
		Sec, Min, Hour, Mday, Mon, Year, Wday, Yday, Isdst int
	}

	GArray struct {
		Data *Gchar
		Leng Guint
	}

	GByteArray struct {
		Data *Guint8
		Leng Guint
	}

	GCompletion struct {
		Items        *GList
		Fnc          GCompletionFunc
		Prefix       *Gchar
		Cache        *GList
		Strncmp_func GCompletionStrncmpFunc
	}

	GDate struct {
		Bits1, Bits2 Guint
		// julian_days : 32
		// julian : 1
		// dmy : 1
		// day : 6
		// month : 4
		// year : 16
	}

	GDebugKey struct {
		Key   *Gchar
		Value Guint
	}

	GError struct {
		Domain  GQuark
		Code    Gint
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
		Ref_count Guint
		Hook_id   Gulong
		Flags     Guint
		Fnc       Gpointer
		Destroy   GDestroyNotify
	}

	GHookList struct {
		Seq_id Gulong
		Bits   Guint
		// hook_size : 16
		// is_setup : 1
		Hooks         *GHook
		Dummy3        Gpointer
		Finalize_hook GHookFinalizeFunc
		Dummy         [2]Gpointer
	}

	GIOChannel struct {
		Ref_count         Gint
		Funcs             *GIOFuncs
		Encoding          *Gchar
		Read_cd           GIConv
		Write_cd          GIConv
		Line_term         *Gchar
		Line_term_len     Guint
		Buf_size          Gsize
		Read_buf          *GString
		Encoded_read_buf  *GString
		Write_buf         *GString
		Partial_write_buf [6]Gchar
		Bits              Guint
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
			offset Gint64,
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
		Flags           Gint
		Arg             GOptionArg
		Arg_data        Gpointer
		Description     *Gchar
		Arg_description *Gchar
	}

	GPollFD struct {
		Fd      Gint
		Events  Gushort
		Revents Gushort
	}

	GPtrArray struct {
		Pdata *Gpointer
		Len   Guint
	}

	GQueue struct {
		Head   *GList
		Tail   *GList
		Length Guint
	}

	GScanner struct {
		User_data        Gpointer
		Max_parse_errors Guint
		Parse_errors     Guint
		Input_name       *Gchar
		Qdata            *GData
		Config           *GScannerConfig
		Token            GTokenType
		Value            GTokenValue
		Line             Guint
		Position         Guint
		Next_token       GTokenType
		Next_value       GTokenValue
		Next_line        Guint
		Next_position    Guint
		Symbol_table     *GHashTable
		Input_fd         Gint
		Text             *Gchar
		Text_end         *Gchar
		Buffer           *Gchar
		Scope_id         Guint
		Msg_handler      GScannerMsgFunc
	}

	GScannerConfig struct {
		Cset_skip_characters  *Gchar
		Cset_identifier_first *Gchar
		Cset_identifier_nth   *Gchar
		Cpair_comment_single  *Gchar
		Bits                  Guint
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
		Padding_dummy Guint
	}

	GSList struct {
		Data Gpointer
		Next *GSList
	}

	GSource struct {
		Callback_data  Gpointer
		Callback_funcs *GSourceCallbackFuncs
		Source_funcs   *GSourceFuncs
		Ref_count      Guint
		Context        *GMainContext
		Priority       Gint
		Flags          Guint
		Source_id      Guint
		Poll_fds       *GSList
		Prev           *GSource
		Next           *GSource
		Name           *Char
		Priv           *GSourcePrivate
	}

	GStaticPrivate struct {
		Index Guint
	}

	GStaticRecMutex struct {
		Mutex GStaticMutex
		Depth Guint
		Owner GSystemThread
	}

	GSystemThread struct {
		// union
		// data[4] Char
		Dummy_double Double
		// dummy_pointer *Void
		// dummy_long Long
	}

	GStaticRWLock struct {
		Mutex         GStaticMutex
		Read_cond     *GCond
		Write_cond    *GCond
		Read_counter  Guint
		Have_writer   Gboolean
		Want_to_read  Guint
		Want_to_write Guint
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
		N_strings Guint
		Strings   **Gchar
		N_nums    Guint
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
		V_int64 Guint64
		// v_float  Gdouble;
		// v_hex  Gulong;
		// v_string  *Gchar;
		// v_comment  *Gchar;
		// v_char  Guchar;
		// v_error  Guint;
	}

	GTuples struct {
		Len Guint
	}

	GVariantBuilder struct {
		X [16]Gsize
	}

	GVariantIter struct {
		X [16]Gsize
	}

	GtkWidget struct {
		Object        GtkObject
		Private_flags Guint16
		State         Guint8
		Saved_state   Guint8
		Name          *Gchar
		Style         *GtkStyle
		Requisition   GtkRequisition
		Allocation    GtkAllocation
		Window        *GdkWindow
		Parent        *GtkWidget
	}

	GtkRequisition struct {
		Width  Gint
		Height Gint
	}

	GTypeClass struct {
		G_type GType
	}

	GTypeInstance struct {
		G_class *GTypeClass
	}

	GObject struct {
		G_type_instance GTypeInstance
		Ref_count       Guint
		Qdata           *GData
	}

	GSourceFuncs struct {
		Prepare func(
			source *GSource,
			timeout_ *Gint) Gboolean
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
		Flags           Guint32
	}

	GdkColor struct {
		Pixel Guint32
		Red   Guint16
		Green Guint16
		Blue  Guint16
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
		Xthickness        Gint
		Ythickness        Gint
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
		Attach_count      Gint
		Depth             Gint
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
		Xthickness       Gint
		Ythickness       Gint
		Rc_properties    *GArray
		Rc_style_lists   *GSList
		Icon_factories   *GSList
		Engine_specified Guint //: 1;
	}

	GdkFont struct {
		Type    GdkFontType
		Ascent  Gint
		Descent Gint
	}

	GdkGC struct {
		Parent_instance GObject
		Clip_x_origin   Gint
		Clip_y_origin   Gint
		Ts_x_origin     Gint
		Ts_y_origin     Gint
		Colormap        *GdkColormap
	}

	GdkColormap struct {
		Parent_instance GObject
		Size            Gint
		Colors          *GdkColor
		Visual          *GdkVisual
		Windowing_data  Gpointer
	}

	GdkVisual struct {
		Parent_instance GObject
		Type            GdkVisualType
		Depth           Gint
		Byte_order      GdkByteOrder
		Colormap_size   Gint
		Bits_per_rgb    Gint
		Red_mask        Guint32
		Red_shift       Gint
		Red_prec        Gint
		Green_mask      Guint32
		Green_shift     Gint
		Green_prec      Gint
		Blue_mask       Guint32
		Blue_shift      Gint
		Blue_prec       Gint
	}

	GdkRectangle struct {
		X      Gint
		Y      Gint
		Width  Gint
		Height Gint
	}

	GtkCurve struct {
		Graph         GtkDrawingArea
		Cursor_type   Gint
		Min_x         Gfloat
		Max_x         Gfloat
		Min_y         Gfloat
		Max_y         Gfloat
		Pixmap        *GdkPixmap
		Curve_type    GtkCurveType
		Height        Gint
		Grab_point    Gint
		Last          Gint
		Num_points    Gint
		Point         *GdkPoint
		Num_ctlpoints Gint
		Ctlpoint      [2]*Gfloat //TODO(t): Gfloat (*ctlpoint)[2]; ???
	}

	GtkDrawingArea struct {
		Widget    GtkWidget
		Draw_data Gpointer
	}

	GdkPoint struct {
		X Gint
		Y Gint
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

	GtkBin struct {
		Container GtkContainer
		Child     *GtkWidget
	}

	GtkContainer struct {
		Widget      GtkWidget
		Focus_child *GtkWidget
		Bits        Guint
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
		Configure_request_count Guint16
		Bits                    Guint
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
		Frame_left           Guint
		Frame_top            Guint
		Frame_right          Guint
		Frame_bottom         Guint
		Keys_changed_handler Guint
		Mnemonic_modifier    GdkModifierType
		Screen               *GdkScreen
	}

	GtkWindowGroup struct {
		Parent_instance GObject
		Grabs           *GSList
	}

	GdkScreen struct {
		Parent_instance GObject
		Closed          Guint      // : 1
		Normal_gcs      *[32]GdkGC //TODO(t): CHECK
		Exposure_gcs    *[32]GdkGC //TODO(t): CHECK
		Subwindow_gcs   *[32]GdkGC //TODO(t): CHECK
		Font_options    *Cairo_font_options_t
		Resolution      Double
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

	GtkAccelGroup struct {
		Parent         GObject
		Lock_count     Guint
		Modifier_mask  GdkModifierType
		Acceleratables *GSList
		N_accels       Guint
		Priv_accels    *GtkAccelGroupEntry
	}

	GtkAccelGroupEntry struct {
		Key              GtkAccelKey
		Closure          *GClosure
		Accel_path_quark GQuark
	}

	GtkAccelKey struct {
		Accel_key   Guint
		Accel_mods  GdkModifierType
		Accel_flags Guint //: 16
	}

	GClosure struct {
		Bits Guint
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
			n_param_values Guint,
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
		// v_int     Gint
		// v_uint    Guint
		// v_long    glong
		// v_ulong   Gulong
		// v_int64   Gint64
		// v_uint64  Guint64
		// v_float   Gfloat
		// v_double  Gdouble
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
		Callback_action Guint
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

	GtkList struct {
		Container        GtkContainer
		Children         *GList
		Selection        *GList
		Undo_selection   *GList
		Undo_unselection *GList
		Last_focus_child *GtkWidget
		Undo_focus_child *GtkWidget
		Htimer           Guint
		Vtimer           Guint
		Anchor           Gint
		Drag_pos         Gint
		Anchor_state     GtkStateType
		Bits             Guint
		// selection_mode : 2
		// drag_selection : 1
		// add_mode : 1
	}

	GtkListItem struct {
		Item GtkItem
	}

	GtkItem struct {
		Bin GtkBin
	}

	GtkOldEditable struct {
		Widget              GtkWidget
		Current_pos         Guint
		Selection_start_pos Guint
		Selection_end_pos   Guint
		Bits                Guint
		// has_selection : 1
		// editable : 1
		// visible : 1
		Clipboard_text *Gchar
	}

	GtkOptionMenu struct {
		Button    GtkButton
		Menu      *GtkWidget
		Menu_item *GtkWidget
		Width     Guint16
		Height    Guint16
	}

	GtkButton struct {
		Bin              GtkBin
		Event_window     *GdkWindow
		Label_text       *Gchar
		Activate_timeout Guint
		Bits             Guint
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
		Buffer_width  Guint16
		Buffer_height Guint16
		Bpp           Guint16
		Rowstride     Guint16
		Dither        GdkRgbDither
		Bits          Guint
		// type : 1
		// expand : 1
	}

	GtkPreviewInfo struct {
		Lookup *Guchar
		Gamma  Gdouble
	}

	GtkTipsQuery struct {
		Label GtkLabel
		Bits  Guint
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
		Ref_count Guint
	}

	GtkLabel struct {
		Misc  GtkMisc
		Label *Gchar
		Bits  Guint
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
		Mnemonic_keyval Guint
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
		Xalign Gfloat
		Yalign Gfloat
		Xpad   Guint16
		Ypad   Guint16
	}

	GtkCList struct {
		Container           GtkContainer
		Flags               Guint16
		Reserved1           Gpointer
		Reserved2           Gpointer
		Freeze_count        Guint
		Internal_allocation GdkRectangle
		Rows                Gint
		Row_height          Gint
		Row_list            *GList
		Row_list_end        *GList
		Columns             Gint
		Column_title_area   GdkRectangle
		Title_window        *GdkWindow
		Column              *GtkCListColumn
		Clist_window        *GdkWindow
		Clist_window_width  Gint
		Clist_window_height Gint
		Hoffset             Gint
		Voffset             Gint
		Shadow_type         GtkShadowType
		Selection_mode      GtkSelectionMode
		Selection           *GList
		Selection_end       *GList
		Undo_selection      *GList
		Undo_unselection    *GList
		Undo_anchor         Gint
		Button_actions      [5]Guint8
		Drag_button         Guint8
		Click_cell          GtkCListCellInfo
		Hadjustment         *GtkAdjustment
		Vadjustment         *GtkAdjustment
		Xor_gc              *GdkGC
		Fg_gc               *GdkGC
		Bg_gc               *GdkGC
		Cursor_drag         *GdkCursor
		X_drag              Gint
		Focus_row           Gint
		Focus_header_column Gint
		Anchor              Gint
		Anchor_state        GtkStateType
		Drag_pos            Gint
		Htimer              Gint
		Vtimer              Gint
		Sort_type           GtkSortType
		Compare             GtkCListCompareFunc
		Sort_column         Gint
		Drag_highlight_row  Gint
		Drag_highlight_pos  GtkCListDragPos
	}

	GtkCListColumn struct {
		Title         *Gchar
		Area          GdkRectangle
		Button        *GtkWidget
		Window        *GdkWindow
		Width         Gint
		Min_width     Gint
		Max_width     Gint
		Justification GtkJustification
		Bits          Guint
		// visible : 1
		// width_set : 1
		// resizeable : 1
		// auto_resize : 1
		// button_passive : 1
	}

	GtkCListCellInfo struct {
		Row    Gint
		Column Gint
	}

	GtkAdjustment struct {
		Parent_instance GtkObject
		Lower           Gdouble
		Upper           Gdouble
		Value           Gdouble
		Step_increment  Gdouble
		Page_increment  Gdouble
		Page_size       Gdouble
	}

	GtkCListRow struct {
		Cell       *_GtkCell
		State      GtkStateType
		Foreground GdkColor
		Background GdkColor
		Style      *GtkStyle
		Data       Gpointer
		Destroy    GDestroyNotify
		Bits       Guint
		// fg_set : 1
		// bg_set : 1
		// selectable : 1
	}

	_GtkCell struct { // TODO(t):fix
		Type       GtkCellType
		Vertical   Gint16
		Horizontal Gint16
		Style      *GtkStyle
		/* union
		   text  *Gchar;
		   struct {
				pixmap  *GdkPixmap;
				mask  *GdkBitmap;
		   } pm;
		   struct {
				text  *Gchar;
				spacing  Guint8;
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
		Entry_change_id Guint
		List_change_id  Guint
		Bits            Guint
		// value_in_list:1
		// ok_if_empty:1
		// case_sensitive:1
		// use_arrows:1
		// use_arrows_always:1
		Current_button Guint16
		Activate_id    Guint
	}

	GtkHBox struct {
		Box GtkBox
	}

	GtkBox struct {
		Container   GtkContainer
		Children    *GList
		Spacing     Gint16
		Homogeneous Guint // : 1
	}

	GtkCTree struct {
		Clist        GtkCList
		Lines_gc     *GdkGC
		Tree_indent  Gint
		Tree_spacing Gint
		Tree_column  Gint
		Bits         Guint
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
		Level         Guint16
		Bits          Guint
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
		User_action_count          Guint
		Bits                       Guint
		// modified : 1
		// has_selection : 1
	}

	GtkTextTagTable struct {
		Parent_instance GObject
		Hash            *GHashTable
		Anonymous       *GSList
		Anon_count      Gint
		Buffers         *GSList
	}

	GtkTextView struct {
		Parent_instance        GtkContainer
		Layout                 *GtkTextLayout
		Buffer                 *GtkTextBuffer
		Selection_drag_handler Guint
		Scroll_timeout         Guint
		Pixels_above_lines     Gint
		Pixels_below_lines     Gint
		Pixels_inside_wrap     Gint
		Wrap_mode              GtkWrapMode
		Justify                GtkJustification
		Left_margin            Gint
		Right_margin           Gint
		Indent                 Gint
		Tabs                   *PangoTabArray
		Bits                   Guint
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
		Xoffset                     Gint
		Yoffset                     Gint
		Width                       Gint
		Height                      Gint
		Virtual_cursor_x            Gint
		Virtual_cursor_y            Gint
		First_para_mark             *GtkTextMark
		First_para_pixels           Gint
		Dnd_mark                    *GtkTextMark
		Blink_timeout               Guint
		First_validate_idle         Guint
		Incremental_validate_idle   Guint
		Im_context                  *GtkIMContext
		Popup_menu                  *GtkWidget
		Drag_start_x                Gint
		Drag_start_y                Gint
		Children                    *GSList
		Pending_scroll              *GtkTextPendingScroll
		Pending_place_cursor_button Gint
	}

	GtkTextMark struct {
		Parent_instance GObject
		Segment         Gpointer
	}

	GtkTextIter struct {
		Dummy1  Gpointer
		Dummy2  Gpointer
		Dummy3  Gint
		Dummy4  Gint
		Dummy5  Gint
		Dummy6  Gint
		Dummy7  Gint
		Dummy8  Gint
		Dummy9  Gpointer
		Dummy10 Gpointer
		Dummy11 Gint
		Dummy12 Gint
		Dummy13 Gint
		Dummy14 Gpointer
	}

	GdkEventKey struct {
		Type             GdkEventType
		Window           *GdkWindow
		Send_event       Gint8
		Time             Guint32
		State            Guint
		Keyval           Guint
		Length           Gint
		String           *Gchar
		Hardware_keycode Guint16
		Group            Guint8
		Is_modifier      Guint // : 1
	}

	GtkTextChildAnchor struct {
		Parent_instance GObject
		Segment         Gpointer
	}

	GtkTextAttributes struct {
		Refcount           Guint
		Appearance         GtkTextAppearance
		Justification      GtkJustification
		Direction          GtkTextDirection
		Font               *PangoFontDescription
		Font_scale         Gdouble
		Left_margin        Gint
		Indent             Gint
		Right_margin       Gint
		Pixels_above_lines Gint
		Pixels_below_lines Gint
		Pixels_inside_wrap Gint
		Tabs               *PangoTabArray
		Wrap_mode          GtkWrapMode
		Language           *PangoLanguage
		Pg_bg_color        *GdkColor
		Bits               Guint
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
		Rise       Gint
		Padding1   Gpointer
		Bits       Guint
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
		Build_insensitive  Guint // : 1
	}

	GtkToolbar struct {
		Container    GtkContainer
		Num_children Gint
		Children     *GList
		Orientation  GtkOrientation
		Style        GtkToolbarStyle
		Icon_size    GtkIconSize
		Tooltips     *GtkTooltips
		Button_maxw  Gint
		Button_maxh  Gint
		_, _         Guint
		Bits         Guint
		// style_set : 1
		// icon_size_set : 1
	}

	GtkTooltips struct {
		Parent_instance  GtkObject
		Tip_window       *GtkWidget
		Tip_label        *GtkWidget
		Active_tips_data *GtkTooltipsData
		Tips_data_list   *GList
		Bits, Bits2      Guint //TODO(t): 33 bits Alignment/size?
		// delay : 30
		// enabled : 1
		// have_grab : 1
		// use_sticky_delay : 1
		Timer_tag    Gint
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
		Format    Gint
		Data      *Guchar
		Length    Gint
		Display   *GdkDisplay
	}

	GdkDisplay struct {
		Parent_instance   GObject
		Queued_events     *GList
		Queued_tail       *GList
		Button_click_time [2]Guint32
		Button_window     [2]*GdkWindow
		Button_number     [2]Gint
		Double_click_time Guint
		Core_pointer      *GdkDevice
		Pointer_hooks     *GdkDisplayPointerHooks
		Bits              Guint
		// closed : 1
		// ignore_core_events : 1
		Double_click_distance Guint
		Button_x              [2]Gint
		Button_y              [2]Gint
		Pointer_grabs         *GList
		Keyboard_grab         GdkKeyboardGrabInfo
		Pointer_info          GdkPointerWindowInfo
		Last_event_time       Guint32
	}

	GdkKeyboardGrabInfo struct {
		Window        *GdkWindow
		Native_window *GdkWindow
		Serial        Gulong
		Owner_events  Gboolean
		Time          Guint32
	}

	GdkPointerWindowInfo struct {
		Toplevel_under_pointer *GdkWindow
		Window_under_pointer   *GdkWindow
		Toplevel_x, Toplevel_y Gdouble
		State                  Guint32
		Button                 Guint32
		Motion_hint_serial     Gulong
	}

	GdkDevice struct {
		Parent_instance GObject
		Name            *Gchar
		Source          GdkInputSource
		Mode            GdkInputMode
		Has_cursor      Gboolean
		Num_axes        Gint
		Axes            *GdkDeviceAxis
		Num_keys        Gint
		Keys            *GdkDeviceKey
	}

	GdkDeviceKey struct {
		Keyval    Guint
		Modifiers GdkModifierType
	}

	GdkDisplayPointerHooks struct {
		Get_pointer func(
			display *GdkDisplay,
			screen **GdkScreen,
			x *Gint,
			y *Gint,
			mask *GdkModifierType)
		Window_get_pointer func(
			display *GdkDisplay,
			window *GdkWindow,
			x *Gint,
			y *Gint,
			mask *GdkModifierType) *GdkWindow
		Window_at_pointer func(
			display *GdkDisplay,
			win_x *Gint,
			win_y *Gint) *GdkWindow
	}

	GdkDeviceAxis struct {
		Use GdkAxisUse
		Min Gdouble
		Max Gdouble
	}

	GtkTargetEntry struct {
		Target *Gchar
		Flags  Guint
		Info   Guint
	}

	GtkSizeGroup struct {
		Parent_instance GObject
		Widgets         *GSList
		Mode            Guint8
		Bits            Guint
		// have_width : 1
		// have_height : 1
		// ignore_hidden : 1
		Requisition GtkRequisition
	}

	GtkSpinButton struct {
		Entry         GtkEntry
		Adjustment    *GtkAdjustment
		Panel         *GdkWindow
		Timer         Guint32
		Climb_rate    Gdouble
		Timer_step    Gdouble
		Update_policy GtkSpinButtonUpdatePolicy
		Bits          Guint
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
		Bits   Guint
		// editable : 1
		// visible : 1
		// overwrite_mode : 1
		// in_drag : 1
		Text_length     Guint16
		Text_max_length Guint16
		Text_area       *GdkWindow
		Im_context      *GtkIMContext
		Popup_menu      *GtkWidget
		Current_pos     Gint
		Selection_bound Gint
		Cached_layout   *PangoLayout
		Bits2           Guint
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
		Button         Guint
		Blink_timeout  Guint
		Recompute_idle Guint
		Scroll_offset  Gint
		Ascent         Gint
		Descent        Gint
		X_text_size    Guint16
		X_n_bytes      Guint16
		Preedit_length Guint16
		Preedit_cursor Guint16
		Dnd_position   Gint
		Drag_start_x   Gint
		Drag_start_y   Gint
		Invisible_char Gunichar
		Width_chars    Gint
	}

	GtkTreeModelSort struct {
		Parent               GObject
		Root                 Gpointer
		Stamp                Gint
		Child_flags          Guint
		Child_model          *GtkTreeModel
		Zero_ref_count       Gint
		Sort_list            *GList
		Sort_column_id       Gint
		Order                GtkSortType
		Default_sort_func    GtkTreeIterCompareFunc
		Default_sort_data    Gpointer
		Default_sort_destroy GDestroyNotify
		Changed_id           Guint
		Inserted_id          Guint
		Has_child_toggled_id Guint
		Deleted_id           Guint
		Reordered_id         Guint
	}

	GtkTreeIter struct {
		Stamp      Gint
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
		Stamp                Gint
		Root                 Gpointer
		Last                 Gpointer
		N_columns            Gint
		Sort_column_id       Gint
		Sort_list            *GList
		Order                GtkSortType
		Column_headers       *GType
		Default_sort_func    GtkTreeIterCompareFunc
		Default_sort_data    Gpointer
		Default_sort_destroy GDestroyNotify
		Columns_dirty        Guint // : 1
	}

	GtkUIManager struct {
		Parent       GObject
		Private_data *GtkUIManagerPrivate
	}

	GtkActionGroup struct {
		Parent       GObject
		Private_data *GtkActionGroupPrivate
	}

	GtkAction struct {
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

	GtkCellRenderer struct {
		Parent GtkObject
		Xalign Gfloat
		Yalign Gfloat
		Width  Gint
		Height Gint
		Xpad   Guint16
		Ypad   Guint16
		Bits   Guint
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
		Toggle_size       Guint16
		Accelerator_width Guint16
		Accel_path        *Gchar
		Bits              Guint
		// show_submenu_indicator : 1
		// submenu_placement : 1
		// submenu_direction : 1
		// right_justify: 1
		// timer_from_keypress : 1
		// from_menubar : 1
		Timer Guint
	}

	GtkIMContextSimple struct {
		Object              GtkIMContext
		Tables              *GSList
		Compose_buffer      [7 + 1]Guint
		Tentative_match     Gunichar
		Tentative_match_len Gint
		Bits                Guint
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
		Button            Guint
		Activate_time     Guint32
		Bits              Guint
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
		Width        Guint
		Height       Guint
		Hadjustment  *GtkAdjustment
		Vadjustment  *GtkAdjustment
		Bin_window   *GdkWindow
		Visibility   GdkVisibilityState
		Scroll_x     Gint
		Scroll_y     Gint
		Freeze_count Guint
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
		// int_data  Gint;
		// uint_data  Guint;
		// long_data  glong;
		// ulong_data  Gulong;
		// float_data  Gfloat;
		// Double_data  Gdouble;
		// string_data  *Gchar;
		// object_data  *GtkObject;
		// pointer_data  Gpointer;
		// signal_data struct{f GCallback; d Gpointer}
	}

	GSignalInvocationHint struct {
		Signal_id Guint
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
		Timer        Guint32
		Tab_hborder  Guint16
		Tab_vborder  Guint16
		Bits         Guint
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
		Request_width  Guint16
		Request_height Guint16
		Current_width  Guint16
		Current_height Guint16
		Plug_window    *GdkWindow
		Plug_widget    *GtkWidget
		Xembed_version Gshort
		Bits           Guint
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
		Same_app        Guint // : 1
	}

	GtkPageRange struct {
		Start Gint
		End   Gint
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
		X_align          Gfloat
		Y_align          Gfloat
		Bits             Guint
		// show_text : 1;
		// activity_mode : 1;
		// use_text_format : 1
	}

	GtkProgressBar struct {
		Progress        GtkProgress
		Bar_style       GtkProgressBarStyle
		Orientation     GtkProgressBarOrientation
		Blocks          Guint
		In_block        Gint
		Activity_pos    Gint
		Activity_step   Guint
		Activity_blocks Guint
		Pulse_fraction  Gdouble
		Bits            Guint
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

	GtkCheckMenuItem struct {
		Menu_item GtkMenuItem
		Bits      Guint
		// active : 1
		// always_show_toggle : 1
		// inconsistent : 1
		// draw_as_radio : 1
	}

	GtkToggleButton struct {
		Button GtkButton
		Bits   Guint
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
		Age          Gint
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
		Bits       Guint
		// hscrollbar_policy : 2
		// vscrollbar_policy : 2
		// hscrollbar_visible : 1
		// vscrollbar_visible : 1
		// window_placement : 2
		// focus_out : 1
		Shadow_type Guint16
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
		Toggle_size          Guint
		Toplevel             *GtkWidget
		Tearoff_window       *GtkWidget
		Tearoff_hbox         *GtkWidget
		Tearoff_scrollbar    *GtkWidget
		Tearoff_adjustment   *GtkAdjustment
		View_window          *GdkWindow
		Bin_window           *GdkWindow
		Scroll_offset        Gint
		Saved_scroll_offset  Gint
		Scroll_step          Gint
		Timeout_id           Guint
		Navigation_region    *GdkRegion
		Navigation_timeout   Guint
		Bits                 Guint
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
		Seq_context_id  Guint
		Seq_message_id  Guint
		Grip_window     *GdkWindow
		Has_resize_grip Guint // : 1
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
		Keyval             Guint
		Translation_domain *Gchar
	}

	GtkTable struct {
		Container      GtkContainer
		Children       *GList
		Rows           *GtkTableRowCol
		Cols           *GtkTableRowCol
		Nrows          Guint16
		Ncols          Guint16
		Column_spacing Guint16
		Row_spacing    Guint16
		Bits           Guint
		// homogeneous : 1
	}

	GtkTableRowCol struct {
		Requisition Guint16
		Allocation  Guint16
		Spacing     Guint16
		Bits        Guint
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
		Bits            Guint
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
		Ref_count Guint
	}

	PangoLogAttr struct {
		Bits Guint
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
		Start_time       Guint32
		Windowing_data   Gpointer
	}

	GdkEventExpose struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
		Area       GdkRectangle
		Region     *GdkRegion
		Count      Gint
	}

	GdkEventOwnerChange struct {
		Type           GdkEventType
		Window         *GdkWindow
		Send_event     Gint8
		Owner          GdkNativeWindow
		Reason         GdkOwnerChange
		Selection      GdkAtom
		Time           Guint32
		Selection_time Guint32
	}

	GdkEventProperty struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
		Atom       GdkAtom
		Time       Guint32
		State      Guint
	}

	GdkEventSelection struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
		Selection  GdkAtom
		Target     GdkAtom
		Property   GdkAtom
		Time       Guint32
		Requestor  GdkNativeWindow
	}

	GdkEventSetting struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
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
		Ts_x_origin        Gint
		Ts_y_origin        Gint
		Clip_x_origin      Gint
		Clip_y_origin      Gint
		Graphics_exposures Gint
		Line_width         Gint
		Line_style         GdkLineStyle
		Cap_style          GdkCapStyle
		Join_style         GdkJoinStyle
	}

	GdkGeometry struct {
		Min_width   Gint
		Min_height  Gint
		Max_width   Gint
		Max_height  Gint
		Base_width  Gint
		Base_height Gint
		Width_inc   Gint
		Height_inc  Gint
		Min_aspect  Gdouble
		Max_aspect  Gdouble
		Win_gravity GdkGravity
	}

	GdkImage struct {
		Parent_instance GObject
		Type            GdkImageType
		Visual          *GdkVisual
		Byte_order      GdkByteOrder
		Width           Gint
		Height          Gint
		Depth           Guint16
		Bpp             Guint16
		Bpl             Guint16
		Bits_per_pixel  Guint16
		Mem             Gpointer
		Colormap        *GdkColormap
		Windowing_data  Gpointer
	}

	GObjectClass struct {
		G_type_class         GTypeClass
		Construct_properties *GSList
		Constructor          func(
			Type GType,
			n_construct_properties Guint,
			construct_properties *GObjectConstructParam) *GObject
		Set_property func(
			object *GObject,
			property_id Guint,
			GValue,
			pspec *GParamSpec)
		Get_property func(
			object *GObject,
			property_id Guint,
			value *GValue,
			pspec *GParamSpec)
		Dispose func(
			object *GObject)
		Finalize func(
			object *GObject)
		Dispatch_properties_changed func(
			object *GObject,
			n_pspecs Guint,
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
		Ref_count       Guint
		Param_id        Guint
	}

	GObjectConstructParam struct {
		Pspec *GParamSpec
		Value *GValue
	}

	GtkAboutDialog struct {
		Parent_instance GtkDialog
		Private_data    Gpointer
	}

	GtkAccelLabel struct {
		Label              GtkLabel
		_                  Guint
		Accel_padding      Guint
		Accel_widget       *GtkWidget
		Accel_closure      *GClosure
		Accel_group        *GtkAccelGroup
		Accel_string       *Gchar
		Accel_string_width Guint16
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
		Latin1_to_char   Guint // : 1
		_, _, _, _       func()
	}

	GtkLabelClass struct {
		Parent_class GtkMiscClass
		Move_cursor  func(
			label *GtkLabel,
			step GtkMovementStep,
			count Gint,
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
		Activate_signal                   Guint
		Set_scroll_adjustments_signal     Guint
		Dispatch_child_properties_changed func(
			widget *GtkWidget,
			n_pspecs Guint,
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
			info, time_ Guint)
		Selection_received func(widget *GtkWidget,
			selection_data *GtkSelectionData,
			time_ Guint)
		Drag_begin func(widget *GtkWidget,
			context *GdkDragContext)
		Drag_end func(widget *GtkWidget,
			context *GdkDragContext)
		Drag_data_get func(widget *GtkWidget,
			context *GdkDragContext,
			selection_data *GtkSelectionData,
			info, time_ Guint)
		Drag_data_delete func(widget *GtkWidget,
			context *GdkDragContext)
		Drag_leave func(widget *GtkWidget,
			context *GdkDragContext,
			time_ Guint)
		Drag_motion func(widget *GtkWidget,
			context *GdkDragContext,
			x, y Gint, time_ Guint) Gboolean
		Drag_drop func(widget *GtkWidget,
			context *GdkDragContext,
			x, y Gint, time_ Guint) Gboolean
		Drag_data_received func(widget *GtkWidget,
			context *GdkDragContext,
			x, y Gint,
			selection_data *GtkSelectionData,
			info, time_ Guint)
		Popup_menu func(widget *GtkWidget) Gboolean
		Show_help  func(widget *GtkWidget,
			help_type GtkWidgetHelpType) Gboolean
		Get_accessible func(widget *GtkWidget) *AtkObject
		Screen_changed func(widget *GtkWidget,
			previous_screen *GdkScreen) Gboolean
		Can_activate_accel func(widget *GtkWidget,
			signal_id Guint) Gboolean
		Grab_broken_event func(widget *GtkWidget,
			event *GdkEventGrabBroken) Gboolean
		Composited_changed func(widget *GtkWidget)
		Query_tooltip      func(widget *GtkWidget,
			x, y Gint, keyboard_tooltip Gboolean,
			tooltip *GtkTooltip) Gboolean
		_, _, _ func()
	}

	GtkObjectClass struct {
		Parent_class GInitiallyUnownedClass
		Set_arg      func(object *GtkObject,
			arg *GtkArg,
			arg_id Guint)
		Get_arg func(object *GtkObject,
			arg *GtkArg,
			arg_id Guint)
		Destroy func(object *GtkObject)
	}

	GdkEventButton struct {
		Type           GdkEventType
		Window         *GdkWindow
		Send_event     Gint8
		Time           Guint32
		X              Gdouble
		Y              Gdouble
		Axes           *Gdouble
		State          Guint
		Button         Guint
		Device         *GdkDevice
		X_root, Y_root Gdouble
	}

	GdkEventScroll struct {
		Type           GdkEventType
		Window         *GdkWindow
		Send_event     Gint8
		Time           Guint32
		X              Gdouble
		Y              Gdouble
		State          Guint
		Direction      GdkScrollDirection
		Device         *GdkDevice
		X_root, Y_root Gdouble
	}

	GdkEventMotion struct {
		Type           GdkEventType
		Window         *GdkWindow
		Send_event     Gint8
		Time           Guint32
		X              Gdouble
		Y              Gdouble
		Axes           *Gdouble
		State          Guint
		Is_hint        Gint16
		Device         *GdkDevice
		X_root, Y_root Gdouble
	}

	GdkEventAny struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
	}

	GdkEventCrossing struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
		Subwindow  *GdkWindow
		Time       Guint32
		X          Gdouble
		Y          Gdouble
		X_root     Gdouble
		Y_root     Gdouble
		Mode       GdkCrossingMode
		Detail     GdkNotifyType
		Focus      Gboolean
		State      Guint
	}

	GdkEventConfigure struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
		X, Y       Gint
		Width      Gint
		Height     Gint
	}

	GdkEventFocus struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
		In         Gint16
	}

	GdkEventProximity struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
		Time       Guint32
		Device     *GdkDevice
	}

	GdkEventVisibility struct {
		Type       GdkEventType
		Window     *GdkWindow
		Send_event Gint8
		State      GdkVisibilityState
	}

	GdkEventClient struct {
		Type         GdkEventType
		Window       *GdkWindow
		Send_event   Gint8
		Message_type GdkAtom
		Data_format  Gushort
		// Union
		B [20]Char
		// s [10]Short
		// l [5]Long
	}

	GdkEventWindowState struct {
		Type             GdkEventType
		Window           *GdkWindow
		Send_event       Gint8
		Changed_mask     GdkWindowState
		New_window_state GdkWindowState
	}

	GdkEventGrabBroken struct {
		Yype        GdkEventType
		Window      *GdkWindow
		Send_event  Gint8
		Keyboard    Gboolean
		Implicit    Gboolean
		Grab_window *GdkWindow
	}

	GtkAccessible struct {
		Parent AtkObject
		Widget *GtkWidget
	}

	GtkActionEntry struct {
		Name        *Gchar
		Stock_id    *Gchar
		Label       *Gchar
		Accelerator *Gchar
		Tooltip     *Gchar
		Callback    GCallback
	}

	GtkAlignment struct {
		Bin    GtkBin
		Xalign Gfloat
		Yalign Gfloat
		Xscale Gfloat
		Yscale Gfloat
	}

	GtkArrow struct {
		Misc        GtkMisc
		Arrow_type  Gint16
		Shadow_type Gint16
	}

	GtkAspectFrame struct {
		Frame             GtkFrame
		Xalign            Gfloat
		Yalign            Gfloat
		Ratio             Gfloat
		Obey_child        Gboolean
		Center_allocation GtkAllocation
	}

	GtkAssistant struct {
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
		Priority            Gint
		Widget_path_pspecs  *GSList
		Widget_class_pspecs *GSList
		Class_branch_pspecs *GSList
		Entries             *GtkBindingEntry
		Current             *GtkBindingEntry
		Parsed              Guint // : 1
	}

	GtkBorder struct {
		Left   Gint
		Right  Gint
		Top    Gint
		Bottom Gint
	}

	GtkBuilder struct {
		Parent_instance GObject
		Priv            *GtkBuilderPrivate
	}

	GtkButtonBox struct {
		Box              GtkBox
		Child_min_width  Gint
		Child_min_height Gint
		Child_ipad_x     Gint
		Child_ipad_y     Gint
		Layout_style     GtkButtonBoxStyle
	}

	GtkCalendar struct {
		Widget            GtkWidget
		Header_style      *GtkStyle
		Label_style       *GtkStyle
		Month             Gint
		Year              Gint
		Selected_day      Gint
		Day_month         [6][7]Gint
		Day               [6][7]Gint
		Num_marked_dates  Gint
		Marked_date       [31]Gint
		Display_flags     GtkCalendarDisplayOptions
		Marked_date_color [31]GdkColor
		Gc                *GdkGC
		Xor_gc            *GdkGC
		Focus_row         Gint
		Focus_col         Gint
		Highlight_row     Gint
		Highlight_col     Gint
		Priv              *GtkCalendarPrivate
		Grow_space        [32]Gchar
		_, _, _, _        func()
	}

	GtkCellRendererText struct {
		Parent            GtkCellRenderer
		Text              *Gchar
		Font              *PangoFontDescription
		Font_scale        Gdouble
		Foreground        PangoColor
		Background        PangoColor
		Extra_attrs       *PangoAttrList
		Underline_style   PangoUnderline
		Rise              Gint
		Fixed_height_rows Gint
		Bits              Guint
		// strikethrough : 1
		// editable : 1
		// scale_set : 1
		// foreground_set : 1
		// background_set : 1
		// underline_set : 1
		// rise_set : 1
		// strikethrough_set : 1
		// editable_set : 1
		// calc_fixed_height : 1
	}

	GtkCellRendererToggle struct {
		Parent GtkCellRenderer
		Bits   Guint
		// active : 1;
		// activatable : 1;
		// radio : 1;
	}

	GtkCellView struct {
		Parent_instance GtkWidget
		Priv            *GtkCellViewPrivate
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

	GtkComboBox struct {
		Parent_instance GtkBin
		Priv            *GtkComboBoxPrivate
	}

	GtkComboBoxEntry struct {
		Parent_instance GtkComboBox
		Priv            *GtkComboBoxEntryPrivate
	}

	GtkComboBoxText struct {
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
			property_id Guint,
			value *GValue,
			pspec *GParamSpec)
		Get_child_property func(container *GtkContainer,
			child *GtkWidget,
			property_id Guint,
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
		Size             Gint
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
		Dialog_width    Gint
		Auto_resize     Gboolean
	}

	GtkFrame struct {
		Bin              GtkBin
		Label_widget     *GtkWidget
		Shadow_type      Gint16
		Label_xalign     Gfloat
		Label_yalign     Gfloat
		Child_allocation GtkAllocation
	}

	GtkHandleBox struct {
		Bin          GtkBin
		Bin_window   *GdkWindow
		Float_window *GdkWindow
		Shadow_type  GtkShadowType
		Bits         Guint
		// handle_position : 2;
		// float_window_mapped : 1;
		// child_detached : 1;
		// in_drag : 1;
		// shrink_on_detach : 1;
		// snap_edge : 3 // signed
		Deskoff_x         Gint
		Deskoff_y         Gint
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
		Stamp                Gint
		Seq                  Gpointer
		_                    Gpointer
		Sort_list            *GList
		N_columns            Gint
		Sort_column_id       Gint
		Order                GtkSortType
		Column_headers       *GType
		Length               Gint
		Default_sort_func    GtkTreeIterCompareFunc
		Default_sort_data    Gpointer
		Default_sort_destroy GDestroyNotify
		Columns_dirty        Guint // : 1
	}

	GtkPaned struct {
		Container       GtkContainer
		Child1          *GtkWidget
		Child2          *GtkWidget
		Handle          *GdkWindow
		Xor_gc          *GdkGC
		Cursor_type     GdkCursorType
		Handle_pos      GdkRectangle
		Child1_size     Gint
		Last_allocation Gint
		Min_position    Gint
		Max_position    Gint
		Bits            Guint
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
		Drag_pos          Gint
		Original_position Gint
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
		Bits          Guint
		// inverted : 1
		// flippable : 1
		// has_stepper_a : 1
		// has_stepper_b : 1
		// has_stepper_c : 1
		// has_stepper_d : 1
		// need_recalc : 1
		// slider_size_fixed : 1
		Min_slider_size Gint
		Orientation     GtkOrientation
		Range_rect      GdkRectangle
		Slider_start    Gint
		Slider_end      Gint
		Round_digits    Gint
		Bits2           Guint
		// trough_click_forward : 1
		// update_pending : 1
		Layout                        *GtkRangeLayout
		Timer                         *GtkRangeStepTimer
		Slide_initial_slider_position Gint
		Slide_initial_coordinate      Gint
		Update_timeout_id             Guint
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
		Xsrc          Gint
		Ysrc          Gint
		Slider_size   Gint
		Lower         Gdouble
		Upper         Gdouble
		Position      Gdouble
		Max_size      Gdouble
	}

	GtkScale struct {
		Range  GtkRange
		Digits Gint
		Bits   Guint
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
	   xalign  Gfloat;
	   property_changed_signal  Guint;
	   spacing  Gint;
	   column_type  GtkTreeViewColumnSizing;
	   requested_width  Gint;
	   button_request  Gint;
	   resized_width  Gint;
	   width  Gint;
	   fixed_width  Gint;
	   min_width  Gint;
	   max_width  Gint;
	   drag_x  Gint;
	   drag_y  Gint;
	   title  *Gchar;
	   cell_list  *GList;
	   sort_clicked_signal  Guint;
	   sort_column_changed_signal  Guint;
	   sort_column_id  Gint;
	   sort_order  GtkSortType;
	    Guint visible : 1;
	    Guint resizable : 1;
	    Guint clickable : 1;
	    Guint dirty : 1;
	    Guint show_sort_indicator : 1;
	    Guint maybe_reordered : 1;
	    Guint reorderable : 1;
	    Guint use_resized_width : 1;
	    Guint expand : 1;
	   };
	*/

	GtkTypeInfo struct {
		Type_name            *Gchar
		Object_size          Guint
		Class_size           Guint
		Class_init_func      GtkClassInitFunc
		Object_init_func     GtkObjectInitFunc
		Reserved_1           Gpointer
		Reserved_2           Gpointer
		Base_class_init_func GtkClassInitFunc
	}

	/*
	   GtkWidgetAuxInfo  struct
	   {
	   x  Gint;
	   y  Gint;
	   width  Gint;
	   height  Gint;
	    Guint x_set : 1;
	    Guint y_set : 1;
	   };
	*/

	GtkBindingEntry struct {
		Keyval      Guint
		Modifiers   GdkModifierType
		Binding_set *GtkBindingSet
		Bits        Guint
		// destroyed : 1
		// in_emission : 1
		// marks_unbound : 1
		Set_next  *GtkBindingEntry
		Hash_next *GtkBindingEntry
		Signals   *GtkBindingSignal
	}

	PangoColor struct {
		Red   Guint16
		Green Guint16
		Blue  Guint16
	}

	GtkVBox struct {
		Box GtkBox
	}

	GtkRulerMetric struct {
		Metric_name     *Gchar
		Abbrev          *Gchar
		Pixels_per_unit Gdouble
		Ruler_scale     [10]Gdouble
		Subdivide       [5]Gint
	}

	GtkBindingSignal struct {
		Next        *GtkBindingSignal
		Signal_name *Gchar
		N_args      Guint
		Args        *GtkBindingArg
	}

	GtkBindingArg struct {
		Arg_type GType
		//Union {
		// long_data  Glong;
		Double_data Gdouble
		// string_data  *Gchar;
		//}
	}

	GtkRadioActionEntry struct {
		Name        *Gchar
		Stock_id    *Gchar
		Label       *Gchar
		Accelerator *Gchar
		Tooltip     *Gchar
		Value       Gint
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
		Keycode Guint
		Group   Gint
		Level   Gint
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
		Xx Double
		Xy Double
		Yx Double
		Yy Double
		X0 Double
		Y0 Double
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
			x *Gint,
			y *Gint,
			mask *GdkModifierType) *GdkWindow
		Window_at_pointer func(screen *GdkScreen,
			win_x *Gint,
			win_y *Gint) *GdkWindow
	}

	GdkRgbCmap struct {
		Colors    [256]Guint32
		N_colors  Gint
		Info_list *GSList
	}

	GdkSegment struct {
		X1 Gint
		Y1 Gint
		X2 Gint
		Y2 Gint
	}

	GdkSpan struct {
		X     Gint
		Y     Gint
		Width Gint
	}

	GdkTimeCoord struct {
		Time Guint32
		Axes [128]Gdouble
	}

	GdkTrapezoid struct {
		Y1, X11, X21, Y2, X12, X22 Double
	}

	GdkWindowAttr struct {
		Title             *Gchar
		Event_mask        Gint
		X, Y              Gint
		Width             Gint
		Height            Gint
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
		Start_index Guint
		End_index   Guint
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
		Offset    Gint
		Length    Gint
		Num_chars Gint
		Analysis  PangoAnalysis
	}

	PangoAnalysis struct {
		Shape_engine *PangoEngineShape
		Lang_engine  *PangoEngineLang
		Font         *PangoFont
		Level        Guint8
		Gravity      Guint8
		Flags        Guint8
		Script       Guint8
		Language     *PangoLanguage
		Extra_attrs  *GSList
	}

	PangoGlyphString struct {
		Num_glyphs   Gint
		Glyphs       *PangoGlyphInfo
		Log_clusters *Gint
		Space        Gint
	}

	PangoLayoutLine struct {
		Layout      *PangoLayout
		Start_index Gint
		Length      Gint
		Runs        *GSList
		Bits        Guint
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
		Is_cluster_start Guint // : 1
	}

	GdkPixdata struct {
		Magic        Guint32
		Length       Gint32
		Pixdata_type GdkPixdataType
		Rowstride    Guint32
		Width        Guint32
		Height       Guint32
		Pixel_data   *Guint8
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

	Cairo_user_data_key_t struct {
		Unused int
	}

	Cairo_matrix_t struct {
		Xx Double
		Yx Double
		Xy Double
		Yy Double
		X0 Double
		Y0 Double
	}

	Cairo_rectangle_list_t struct {
		Status         Cairo_status_t
		Rectangles     *Cairo_rectangle_t
		Num_rectangles int
	}

	Cairo_rectangle_t struct {
		X, Y, Width, Height Double
	}

	Cairo_glyph_t struct {
		Index Unsigned_long
		X     Double
		Y     Double
	}

	Cairo_text_cluster_t struct {
		Num_bytes  int
		Num_glyphs int
	}

	Cairo_text_extents_t struct {
		X_bearing Double
		Y_bearing Double
		Width     Double
		Height    Double
		X_advance Double
		Y_advance Double
	}

	Cairo_font_extents_t struct {
		Ascent        Double
		Descent       Double
		Height        Double
		Max_x_advance Double
		Max_y_advance Double
	}

	Cairo_path_t struct {
		Status   Cairo_status_t
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
			X, Y Double
		}
	}

	Cairo_rectangle_int_t struct {
		X, Y          int
		Width, Height int
	}

	Cairo_script_interpreter_hooks_t struct {
		Closure         *Void
		Surface_create  Csi_surface_create_func_t
		Surface_destroy Csi_destroy_func_t
		Context_create  Csi_context_create_func_t
		Context_destroy Csi_destroy_func_t
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
		Type      Gint
		State     Guint
		Keyval    Guint
		Length    Gint
		String    *Gchar
		Keycode   Guint16
		Timestamp Guint32
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
		X      Gint
		Y      Gint
		Width  Gint
		Height Gint
	}

	AtkTextRange struct {
		Bounds       AtkTextRectangle
		Start_offset Gint
		End_offset   Gint
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
		Bits            Guint
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
		N_static_gsub_features Guint
		Static_gpos_features   *PangoOTFeatureMap
		N_static_gpos_features Guint
		Other_features         *PangoOTFeatureMap
		N_other_features       Guint
	}

	PangoOTFeatureMap struct {
		Feature_name [5]Char
		Property_bit Gulong
	}

	PangoOTGlyph struct {
		Glyph      Guint32
		Properties Guint
		Cluster    Guint
		Component  Gushort
		LigID      Gushort
		Internal   Guint
	}

	FT_FaceRec struct {
		Num_faces           FT_Long
		Face_index          FT_Long
		Face_flags          FT_Long
		Style_flags         FT_Long
		Num_glyphs          FT_Long
		Family_name         *FT_String
		Style_name          *FT_String
		Num_fixed_sizes     FT_Int
		Available_sizes     *FT_Bitmap_Size
		Num_charmaps        FT_Int
		Charmaps            *FT_CharMap
		Generic             FT_Generic
		Bbox                FT_BBox
		Units_per_EM        FT_UShort
		Ascender            FT_Short
		Descender           FT_Short
		Height              FT_Short
		Max_advance_width   FT_Short
		Max_advance_height  FT_Short
		Underline_position  FT_Short
		Underline_thickness FT_Short
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
		Height FT_Short
		Width  FT_Short
		Size   FT_Pos
		X_ppem FT_Pos
		Y_ppem FT_Pos
	}

	FT_Generic_Finalizer func(
		Object *Void)

	FT_Open_Args struct {
		Flags       FT_UInt
		Memory_base *FT_Byte
		Memory_size FT_Long
		Pathname    *FT_String
		Stream      FT_Stream
		Driver      FT_Module
		Num_params  FT_Int
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
		HoriResolution FT_UInt
		VertResolution FT_UInt
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
		Face        FT_Face
		Encoding    FT_Encoding
		Platform_id FT_UShort
		Encoding_id FT_UShort
	}

	FT_GlyphSlotRec struct {
		Library           FT_Library
		Face              FT_Face
		Next              FT_GlyphSlot
		Reserved          FT_UInt
		Generic           FT_Generic
		Metrics           FT_Glyph_Metrics
		LinearHoriAdvance FT_Fixed
		LinearVertAdvance FT_Fixed
		Advance           FT_Vector
		Format            FT_Glyph_Format
		Bitmap            FT_Bitmap
		Bitmap_left       FT_Int
		Bitmap_top        FT_Int
		Outline           FT_Outline
		Num_subglyphs     FT_UInt
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
		N_contours Short
		N_points   Short
		Points     *FT_Vector
		Tags       *Char
		Contours   *Short
		Flags      int
	}

	FT_BBox struct {
		XMin, YMin FT_Pos
		XMax, YMax FT_Pos
	}

	FT_SizeRec struct {
		Face     FT_Face
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
		X_ppem      FT_UShort
		Y_ppem      FT_UShort
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
		Underline_position  FT_Short
		Underline_thickness FT_UShort
	}

	PS_PrivateRec struct {
		Unique_id              FT_Int
		LenIV                  FT_Int
		Num_blue_values        FT_Byte
		Num_other_blues        FT_Byte
		Num_family_blues       FT_Byte
		Num_family_other_blues FT_Byte
		Blue_values            [14]FT_Short
		Other_blues            [10]FT_Short
		Family_blues           [14]FT_Short
		Family_other_blues     [10]FT_Short
		Blue_scale             FT_Fixed
		Blue_shift             FT_Int
		Blue_fuzz              FT_Int
		Standard_width         [1]FT_UShort
		Standard_height        [1]FT_UShort
		Num_snap_widths        FT_Byte
		Num_snap_heights       FT_Byte
		Force_bold             FT_Bool
		Round_stem_up          FT_Bool
		Snap_widths            [13]FT_Short
		Snap_heights           [13]FT_Short
		Expansion_factor       FT_Fixed
		Language_group         FT_Long
		Password               FT_Long
		Min_feature            [2]FT_Short
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
		Num_axis    FT_UInt
		Num_designs FT_UInt
		Axis        [4]FT_MM_Axis
	}

	FT_MM_Axis struct {
		Name    *FT_String
		Minimum FT_Long
		Maximum FT_Long
	}

	FT_MM_Var struct {
		Num_axis        FT_UInt
		Num_designs     FT_UInt
		Num_namedstyles FT_UInt
		Axis            *FT_Var_Axis
		Namedstyle      *FT_Var_Named_Style
	}

	FT_Var_Axis struct {
		Name    *FT_String
		Minimum FT_Fixed
		Def     FT_Fixed
		Maximum FT_Fixed
		Tag     FT_ULong
		Strid   FT_UInt
	}

	FT_Var_Named_Style struct {
		Coords *FT_Fixed
		Strid  FT_UInt
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
		X        Short
		Len      Unsigned_short
		Coverage Unsigned_char
	}

	BDF_PropertyRec struct {
		Type BDF_PropertyType
		// Union
		Atom *Char
		// integer  FT_Int32
		// cardinal FT_UInt32
	}

	FT_WinFNT_HeaderRec struct {
		Version               FT_UShort
		File_size             FT_ULong
		Copyright             [60]FT_Byte
		File_type             FT_UShort
		Nominal_point_size    FT_UShort
		Vertical_resolution   FT_UShort
		Horizontal_resolution FT_UShort
		Ascent                FT_UShort
		Internal_leading      FT_UShort
		External_leading      FT_UShort
		Italic                FT_Byte
		Underline             FT_Byte
		Strike_out            FT_Byte
		Weight                FT_UShort
		Charset               FT_Byte
		Pixel_width           FT_UShort
		Pixel_height          FT_UShort
		Pitch_and_family      FT_Byte
		Avg_width             FT_UShort
		Max_width             FT_UShort
		First_char            FT_Byte
		Last_char             FT_Byte
		Default_char          FT_Byte
		Break_char            FT_Byte
		Bytes_per_row         FT_UShort
		Device_offset         FT_ULong
		Face_name_offset      FT_ULong
		Bits_pointer          FT_ULong
		Bits_offset           FT_ULong
		Reserved              FT_Byte
		Flags                 FT_ULong
		A_space               FT_UShort
		B_space               FT_UShort
		C_space               FT_UShort
		Color_table_offset    FT_UShort
		Reserved1             [4]FT_ULong
	}

	FT_SfntName struct {
		Platform_id FT_UShort
		Encoding_id FT_UShort
		Language_id FT_UShort
		Name_id     FT_UShort
		String      *FT_Byte
		String_len  FT_UInt
	}

	FTC_ScalerRec struct {
		Face_id FTC_FaceID
		Width   FT_UInt
		Height  FT_UInt
		Pixel   FT_Int
		X_res   FT_UInt
		Y_res   FT_UInt
	}

	FTC_ImageTypeRec struct {
		Face_id FTC_FaceID
		Width   FT_Int
		Height  FT_Int
		Flags   FT_Int32
	}

	FTC_SBitRec struct {
		Width     FT_Byte
		Height    FT_Byte
		Left      FT_Char
		Top       FT_Char
		Format    FT_Byte
		Max_grays FT_Byte
		Pitch     FT_Short
		Xadvance  FT_Char
		Yadvance  FT_Char
		Buffer    *FT_Byte
	}

	FTC_FontRec struct {
		Face_id    FTC_FaceID
		Pix_width  FT_UShort
		Pix_height FT_UShort
	}

	FT_Bitmap struct {
		Rows         int
		Width        int
		Pitch        int
		Buffer       *Unsigned_char
		Num_grays    Short
		Pixel_mode   Char
		Palette_mode Char
		Palette      *Void
	}

	GParameter struct {
		Name  string
		Value GValue
	}

	GSignalQuery struct {
		Signal_id    Guint
		Signal_name  string
		Itype        GType
		Signal_flags GSignalFlags
		Return_type  GType
		N_params     Guint
		Param_types  *GType
	}

	GTypeQuery struct {
		Type          GType
		Type_name     string
		Class_size    Guint
		Instance_size Guint
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
		Ref_count   Gint
		Name        *Gchar
		Methods     **GDBusMethodInfo
		Signals     **GDBusSignalInfo
		Properties  **GDBusPropertyInfo
		Annotations **GDBusAnnotationInfo
	}

	GDBusMethodInfo struct {
		Ref_count   Gint
		Name        *Gchar
		In_args     **GDBusArgInfo
		Out_args    **GDBusArgInfo
		Annotations **GDBusAnnotationInfo
	}

	GDBusSignalInfo struct {
		Ref_count   Gint
		Name        *Gchar
		Args        **GDBusArgInfo
		Annotations **GDBusAnnotationInfo
	}

	GDBusPropertyInfo struct {
		Ref_count   Gint
		Name        *Gchar
		Signature   *Gchar
		Flags       GDBusPropertyInfoFlags
		Annotations **GDBusAnnotationInfo
	}

	GDBusArgInfo struct {
		Ref_count   Gint
		Name        *Gchar
		Signature   *Gchar
		Annotations **GDBusAnnotationInfo
	}

	GDBusAnnotationInfo struct {
		Ref_count   Gint
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
		Ref_count   Gint
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
		Error_code      Gint
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
		Minimum      Gint
		Maximum      Gint
		N_values     Guint
		Values       *GEnumValue
	}

	GEnumValue struct {
		Value      Gint
		Value_name *Gchar
		Value_nick *Gchar
	}

	GFlagsClass struct {
		G_type_class GTypeClass
		Mask         Guint
		N_values     Guint
		Values       *GFlagsValue
	}

	GFlagsValue struct {
		Value      Guint
		Value_name *Gchar
		Value_nick *Gchar
	}

	GTypeInfo struct {
		Class_size     Guint16
		Base_init      GBaseInitFunc
		Base_finalize  GBaseFinalizeFunc
		Class_init     GClassInitFunc
		Class_finalize GClassFinalizeFunc
		Class_data     Gconstpointer
		Instance_size  Guint16
		N_preallocs    Guint16
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
			n_collect_values Guint,
			collect_values *GTypeCValue,
			collect_flags Guint) *Gchar
		Lcopy_format *Gchar
		Lcopy_value  func(
			value *GValue,
			n_collect_values Guint,
			collect_values *GTypeCValue,
			collect_flags Guint) *Gchar
	}

	GValueArray struct {
		N_values     Guint
		Values       *GValue
		N_prealloced Guint
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
		Use_count       Guint
		Type_infos      *GSList
		Interface_infos *GSList
		Name            *Gchar
	}

	GParamSpecTypeInfo struct {
		Instance_size     Guint16
		N_preallocs       Guint16
		Instance_init     func(pspec *GParamSpec)
		Value_type        GType
		Finalize          func(pspec *GParamSpec)
		Value_set_default func(
			pspec *GParamSpec, value *GValue)
		Value_validate func(
			pspec *GParamSpec, value *GValue) Gboolean
		Values_cmp func(pspec *GParamSpec,
			value1 *GValue, value2 *GValue) Gint
	}

	GtkTree struct {
		Container      GtkContainer
		Children       *GList
		Root_tree      *GtkTree
		Tree_owner     *GtkWidget
		Selection      *GList
		Level          Guint
		Indent_value   Guint
		Current_indent Guint
		Bits           Guint
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
		Expanded         Guint // : 1
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
		N_scripts   Gint
	}
)
