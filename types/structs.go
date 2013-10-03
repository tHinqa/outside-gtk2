package types

//TODO(t): add constants for [...]s

type (
	Tm struct {
		Sec, Min, Hour, Mday, Mon, Year, Wday, Yday, Isdst int
	}

	GArray struct {
		data *Gchar
		leng Guint
	}

	GByteArray struct {
		data *Guint8
		leng Guint
	}

	GCompletion struct {
		items        *GList
		fnc          GCompletionFunc
		prefix       *Gchar
		cache        *GList
		strncmp_func GCompletionStrncmpFunc
	}

	GDate struct {
		bits1, bits2 Guint
		// julian_days : 32
		// julian : 1
		// dmy : 1
		// day : 6
		// month : 4
		// year : 16
	}

	GDebugKey struct {
		key   *Gchar
		value Guint
	}

	GError struct {
		domain  GQuark
		code    Gint
		message *Gchar
	}

	GHashTableIter struct {
		dummy1 Gpointer
		dummy2 Gpointer
		dummy3 Gpointer
		dummy4 int
		dummy5 Gboolean
		dummy6 Gpointer
	}

	GHook struct {
		data      Gpointer
		next      *GHook
		prev      *GHook
		ref_count Guint
		hook_id   Gulong
		flags     Guint
		fnc       Gpointer
		destroy   GDestroyNotify
	}

	GHookList struct {
		seq_id Gulong
		bits   Guint
		// hook_size : 16
		// is_setup : 1
		hooks         *GHook
		dummy3        Gpointer
		finalize_hook GHookFinalizeFunc
		dummy         [2]Gpointer
	}

	GIOChannel struct {
		ref_count         Gint
		funcs             *GIOFuncs
		encoding          *Gchar
		read_cd           GIConv
		write_cd          GIConv
		line_term         *Gchar
		line_term_len     Guint
		buf_size          Gsize
		read_buf          *GString
		encoded_read_buf  *GString
		write_buf         *GString
		partial_write_buf [6]Gchar
		bits              Guint
		// use_buffer : 1
		// do_encode : 1
		// close_on_unref : 1
		// is_readable : 1
		// is_writeable : 1
		// is_seekable : 1
		_, _ Gpointer
	}

	GList struct {
		data Gpointer
		next *GList
		prev *GList
	}

	GIOFuncs struct {
		io_read func(
			channel *GIOChannel,
			buf *Gchar,
			count Gsize,
			bytes_read *Gsize,
			err **GError) GIOStatus
		io_write func(
			channel *GIOChannel,
			buf *Gchar,
			count Gsize,
			bytes_written *Gsize,
			err **GError) GIOStatus
		io_seek func(
			channel *GIOChannel,
			offset Gint64,
			type_ GSeekType,
			err **GError) GIOStatus
		io_close func(
			channel *GIOChannel,
			err **GError) GIOStatus
		io_create_watch func(
			channel *GIOChannel,
			condition GIOCondition) *GSource
		io_free func(
			channel *GIOChannel)
		io_set_flags func(
			channel *GIOChannel,
			flags GIOFlags,
			err **GError) GIOStatus
		io_get_flags func(
			channel *GIOChannel) GIOFlags
	}

	GMarkupParser struct {
		start_element func(
			context *GMarkupParseContext,
			element_name *Gchar,
			attribute_names **Gchar,
			attribute_values **Gchar,
			user_data Gpointer,
			error **GError)
		end_element func(
			context *GMarkupParseContext,
			element_name *Gchar,
			user_data Gpointer,
			error **GError)
		text func(
			context *GMarkupParseContext,
			text *Gchar,
			text_len Gsize,
			user_data Gpointer,
			error **GError)
		passthrough func(
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
		data     Gpointer
		next     *GNode
		prev     *GNode
		parent   *GNode
		children *GNode
	}

	GSourceCallbackFuncs struct {
		ref func(
			cb_data Gpointer)
		unref func(
			cb_data Gpointer)
		get func(
			cb_data Gpointer,
			source *GSource,
			fnc *GSourceFunc,
			data *Gpointer)
	}

	GOnce struct {
		status GOnceStatus
		retval Gpointer
	}

	GOptionEntry struct {
		long_name       *Gchar
		short_name      Gchar
		flags           Gint
		arg             GOptionArg
		arg_data        Gpointer
		description     *Gchar
		arg_description *Gchar
	}

	GPollFD struct {
		fd      Gint
		events  Gushort
		revents Gushort
	}

	GPtrArray struct {
		pdata *Gpointer
		len   Guint
	}

	GQueue struct {
		head   *GList
		tail   *GList
		length Guint
	}

	GScanner struct {
		user_data        Gpointer
		max_parse_errors Guint
		parse_errors     Guint
		input_name       *Gchar
		qdata            *GData
		config           *GScannerConfig
		token            GTokenType
		value            GTokenValue
		line             Guint
		position         Guint
		next_token       GTokenType
		next_value       GTokenValue
		next_line        Guint
		next_position    Guint
		symbol_table     *GHashTable
		input_fd         Gint
		text             *Gchar
		text_end         *Gchar
		buffer           *Gchar
		scope_id         Guint
		msg_handler      GScannerMsgFunc
	}

	GScannerConfig struct {
		cset_skip_characters  *Gchar
		cset_identifier_first *Gchar
		cset_identifier_nth   *Gchar
		cpair_comment_single  *Gchar
		bits                  Guint
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
		padding_dummy Guint
	}

	GSList struct {
		data Gpointer
		next *GSList
	}

	GSource struct {
		callback_data  Gpointer
		callback_funcs *GSourceCallbackFuncs
		source_funcs   *GSourceFuncs
		ref_count      Guint
		context        *GMainContext
		priority       Gint
		flags          Guint
		source_id      Guint
		poll_fds       *GSList
		prev           *GSource
		next           *GSource
		name           *Char
		priv           *GSourcePrivate
	}

	GStaticPrivate struct {
		index Guint
	}

	GStaticRecMutex struct {
		mutex GStaticMutex
		depth Guint
		owner GSystemThread
	}

	GSystemThread struct {
		// union
		// data[4] Char
		dummy_double Double
		// dummy_pointer *Void
		// dummy_long Long
	}

	GStaticRWLock struct {
		mutex         GStaticMutex
		read_cond     *GCond
		write_cond    *GCond
		read_counter  Guint
		have_writer   Gboolean
		want_to_read  Guint
		want_to_write Guint
	}

	GString struct {
		str           *Gchar
		len           Gsize
		allocated_len Gsize
	}

	GTestLogBuffer struct {
		data *GString
		msgs *GSList
	}

	GTestLogMsg struct {
		log_type  GTestLogType
		n_strings Guint
		strings   **Gchar
		n_nums    Guint
		nums      *Long_double
	}

	GThread struct {
		Func     GThreadFunc
		data     Gpointer
		joinable Gboolean
		priority GThreadPriority
	}

	GThreadFunctions struct {
		mutex_new       func()
		mutex_lock      func(mutex *GMutex) *GMutex
		mutex_trylock   func(mutex *GMutex) Gboolean
		mutex_unlock    func(mutex *GMutex)
		mutex_free      func(mutex *GMutex)
		cond_new        func() *GCond
		cond_signal     func(cond *GCond)
		cond_broadcast  func(cond *GCond)
		cond_wait       func(cond *GCond, mutex *GMutex)
		cond_timed_wait func(
			cond *GCond,
			mutex *GMutex,
			end_time *GTimeVal) Gboolean
		cond_free     func(cond *GCond)
		private_new   func(destructor GDestroyNotify) *GPrivate
		private_get   func(private_key *GPrivate) Gpointer
		private_set   func(private_key *GPrivate, data Gpointer)
		thread_create func(
			fnc GThreadFunc,
			data Gpointer,
			stack_size Gulong,
			joinable Gboolean,
			bound Gboolean,
			priority GThreadPriority,
			thread Gpointer,
			error **GError)
		thread_yield        func()
		thread_join         func(thread Gpointer)
		thread_exit         func()
		thread_set_priority func(
			thread Gpointer,
			priority GThreadPriority)
		thread_self  func(thread Gpointer)
		thread_equal func(
			thread1 Gpointer,
			thread2 Gpointer) Gboolean
	}

	GThreadPool struct {
		Func      GFunc
		user_data Gpointer
		exclusive Gboolean
	}

	GTimeVal struct {
		tv_sec  Glong
		tv_usec Glong
	}

	GTokenValue struct {
		// union
		// v_symbol  Gpointer;
		// v_identifier  *Gchar;
		// v_binary  Gulong;
		// v_octal  Gulong;
		// v_int  Gulong;
		v_int64 Guint64
		// v_float  Gdouble;
		// v_hex  Gulong;
		// v_string  *Gchar;
		// v_comment  *Gchar;
		// v_char  Guchar;
		// v_error  Guint;
	}

	GTuples struct {
		len Guint
	}

	GVariantBuilder struct {
		x [16]Gsize
	}

	GVariantIter struct {
		x [16]Gsize
	}

	GtkWidget struct {
		object        GtkObject
		private_flags Guint16
		state         Guint8
		saved_state   Guint8
		name          *Gchar
		style         *GtkStyle
		requisition   GtkRequisition
		allocation    GtkAllocation
		window        *GdkWindow
		parent        *GtkWidget
	}

	GdkDrawable struct {
		parent_instance GObject
	}

	GtkRequisition struct {
		width  Gint
		height Gint
	}

	GTypeClass struct {
		g_type GType
	}

	GTypeInstance struct {
		g_class *GTypeClass
	}

	GObject struct {
		g_type_instance GTypeInstance
		ref_count       Guint
		qdata           *GData
	}

	GSourceFuncs struct {
		prepare func(
			source *GSource,
			timeout_ *Gint) Gboolean
		check func(
			source *GSource) Gboolean
		dispatch func(
			source *GSource,
			callback GSourceFunc,
			user_data Gpointer) Gboolean
		finalize func(
			source *GSource)
		closure_callback GSourceFunc
		closure_marshal  GSourceDummyMarshal
	}

	GtkObject struct {
		parent_instance GInitiallyUnowned
		flags           Guint32
	}

	GdkColor struct {
		pixel Guint32
		red   Guint16
		green Guint16
		blue  Guint16
	}

	GtkStyle struct {
		parent_instance   GObject
		fg                [5]GdkColor
		bg                [5]GdkColor
		light             [5]GdkColor
		dark              [5]GdkColor
		mid               [5]GdkColor
		text              [5]GdkColor
		base              [5]GdkColor
		text_aa           [5]GdkColor
		black             GdkColor
		white             GdkColor
		font_desc         *PangoFontDescription
		xthickness        Gint
		ythickness        Gint
		fg_gc             *[5]GdkGC //TODO(t): CHECK
		bg_gc             *[5]GdkGC //TODO(t): CHECK
		light_gc          *[5]GdkGC //TODO(t): CHECK
		dark_gc           *[5]GdkGC //TODO(t): CHECK
		mid_gc            *[5]GdkGC //TODO(t): CHECK
		text_gc           *[5]GdkGC //TODO(t): CHECK
		base_gc           *[5]GdkGC //TODO(t): CHECK
		text_aa_gc        *[5]GdkGC //TODO(t): CHECK
		black_gc          *GdkGC
		white_gc          *GdkGC
		bg_pixmap         *[5]GdkPixmap //TODO(t): CHECK
		attach_count      Gint
		depth             Gint
		colormap          *GdkColormap
		private_font      *GdkFont
		private_font_desc *PangoFontDescription
		rc_style          *GtkRcStyle
		styles            *GSList
		property_cache    *GArray
		icon_factories    *GSList
	}

	GtkRcStyle struct {
		parent_instance  GObject
		name             *Gchar
		bg_pixmap_name   *[5]Gchar //TODO(t): CHECK
		font_desc        *PangoFontDescription
		color_flags      [5]GtkRcFlags
		fg               [5]GdkColor
		bg               [5]GdkColor
		text             [5]GdkColor
		base             [5]GdkColor
		xthickness       Gint
		ythickness       Gint
		rc_properties    *GArray
		rc_style_lists   *GSList
		icon_factories   *GSList
		engine_specified Guint //: 1;
	}

	GdkFont struct {
		Type    GdkFontType
		ascent  Gint
		descent Gint
	}

	GdkGC struct {
		parent_instance GObject
		clip_x_origin   Gint
		clip_y_origin   Gint
		ts_x_origin     Gint
		ts_y_origin     Gint
		colormap        *GdkColormap
	}

	GdkColormap struct {
		parent_instance GObject
		size            Gint
		colors          *GdkColor
		visual          *GdkVisual
		windowing_data  Gpointer
	}

	GdkVisual struct {
		parent_instance GObject
		Type            GdkVisualType
		depth           Gint
		byte_order      GdkByteOrder
		colormap_size   Gint
		bits_per_rgb    Gint
		red_mask        Guint32
		red_shift       Gint
		red_prec        Gint
		green_mask      Guint32
		green_shift     Gint
		green_prec      Gint
		blue_mask       Guint32
		blue_shift      Gint
		blue_prec       Gint
	}

	GdkRectangle struct {
		x      Gint
		y      Gint
		width  Gint
		height Gint
	}

	GtkCurve struct {
		graph         GtkDrawingArea
		cursor_type   Gint
		min_x         Gfloat
		max_x         Gfloat
		min_y         Gfloat
		max_y         Gfloat
		pixmap        *GdkPixmap
		curve_type    GtkCurveType
		height        Gint
		grab_point    Gint
		last          Gint
		num_points    Gint
		point         *GdkPoint
		num_ctlpoints Gint
		ctlpoint      [2]*Gfloat //TODO(t): Gfloat (*ctlpoint)[2]; ???
	}

	GtkDrawingArea struct {
		widget    GtkWidget
		draw_data Gpointer
	}

	GdkPoint struct {
		x Gint
		y Gint
	}

	GtkFileSelection struct {
		parent_instance  GtkDialog
		dir_list         *GtkWidget
		file_list        *GtkWidget
		selection_entry  *GtkWidget
		selection_text   *GtkWidget
		main_vbox        *GtkWidget
		ok_button        *GtkWidget
		cancel_button    *GtkWidget
		help_button      *GtkWidget
		history_pulldown *GtkWidget
		history_menu     *GtkWidget
		history_list     *GList
		fileop_dialog    *GtkWidget
		fileop_entry     *GtkWidget
		fileop_file      *Gchar
		cmpl_state       Gpointer
		fileop_c_dir     *GtkWidget
		fileop_del_file  *GtkWidget
		fileop_ren_file  *GtkWidget
		button_area      *GtkWidget
		action_area      *GtkWidget
		selected_names   *GPtrArray
		last_selected    *Gchar
	}

	GtkDialog struct {
		window      GtkWindow
		vbox        *GtkWidget
		action_area *GtkWidget
		separator   *GtkWidget
	}

	GtkBin struct {
		container GtkContainer
		child     *GtkWidget
	}

	GtkContainer struct {
		widget      GtkWidget
		focus_child *GtkWidget
		bits        Guint
		// border_width : 16
		// need_resize : 1
		// resize_mode : 2
		// reallocate_redraws : 1
		// has_focus_chain : 1
	}

	GtkWindow struct {
		bin                     GtkBin
		title                   *Gchar
		wmclass_name            *Gchar
		wmclass_class           *Gchar
		wm_role                 *Gchar
		focus_widget            *GtkWidget
		default_widget          *GtkWidget
		transient_parent        *GtkWindow
		geometry_info           *GtkWindowGeometryInfo
		frame                   *GdkWindow
		group                   *GtkWindowGroup
		configure_request_count Guint16
		bits                    Guint
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
		frame_left           Guint
		frame_top            Guint
		frame_right          Guint
		frame_bottom         Guint
		keys_changed_handler Guint
		mnemonic_modifier    GdkModifierType
		screen               *GdkScreen
	}

	GtkWindowGroup struct {
		parent_instance GObject
		grabs           *GSList
	}

	GdkScreen struct {
		parent_instance GObject
		closed          Guint      // : 1
		normal_gcs      *[32]GdkGC //TODO(t): CHECK
		exposure_gcs    *[32]GdkGC //TODO(t): CHECK
		subwindow_gcs   *[32]GdkGC //TODO(t): CHECK
		font_options    *Cairo_font_options_t
		resolution      Double
	}

	GtkItemFactory struct {
		object           GtkObject
		path             *Gchar
		accel_group      *GtkAccelGroup
		widget           *GtkWidget
		items            *GSList
		translate_func   GtkTranslateFunc
		translate_data   Gpointer
		translate_notify GDestroyNotify
	}

	GtkAccelGroup struct {
		parent         GObject
		lock_count     Guint
		modifier_mask  GdkModifierType
		acceleratables *GSList
		n_accels       Guint
		priv_accels    *GtkAccelGroupEntry
	}

	GtkAccelGroupEntry struct {
		key              GtkAccelKey
		closure          *GClosure
		accel_path_quark GQuark
	}

	GtkAccelKey struct {
		accel_key   Guint
		accel_mods  GdkModifierType
		accel_flags Guint //: 16
	}

	GClosure struct {
		bits Guint
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
		marshal func(
			closure *GClosure,
			return_value *GValue,
			n_param_values Guint,
			param_values *GValue,
			invocation_hint Gpointer,
			marshal_data Gpointer)
		data      Gpointer
		notifiers *GClosureNotifyData
	}

	GValue struct {
		g_type GType
		//  UNION
		data [2]uint64 // was union{v_int...}data[2]
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
		data   Gpointer
		notify GClosureNotify
	}

	GtkItemFactoryEntry struct {
		path            *Gchar
		accelerator     *Gchar
		callback        GtkItemFactoryCallback
		callback_action Guint
		item_type       *Gchar
		extra_data      Gconstpointer
	}

	GtkMenuEntry struct {
		path          *Gchar
		accelerator   *Gchar
		callback      GtkMenuCallback
		callback_data Gpointer
		widget        *GtkWidget
	}

	GtkList struct {
		container        GtkContainer
		children         *GList
		selection        *GList
		undo_selection   *GList
		undo_unselection *GList
		last_focus_child *GtkWidget
		undo_focus_child *GtkWidget
		htimer           Guint
		vtimer           Guint
		anchor           Gint
		drag_pos         Gint
		anchor_state     GtkStateType
		bits             Guint
		// selection_mode : 2
		// drag_selection : 1
		// add_mode : 1
	}

	GtkListItem struct {
		item GtkItem
	}

	GtkItem struct {
		bin GtkBin
	}

	GtkOldEditable struct {
		widget              GtkWidget
		current_pos         Guint
		selection_start_pos Guint
		selection_end_pos   Guint
		bits                Guint
		// has_selection : 1
		// editable : 1
		// visible : 1
		clipboard_text *Gchar
	}

	GtkOptionMenu struct {
		button    GtkButton
		menu      *GtkWidget
		menu_item *GtkWidget
		width     Guint16
		height    Guint16
	}

	GtkButton struct {
		bin              GtkBin
		event_window     *GdkWindow
		label_text       *Gchar
		activate_timeout Guint
		bits             Guint
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
		widget        GtkWidget
		buffer        *Guchar
		buffer_width  Guint16
		buffer_height Guint16
		bpp           Guint16
		rowstride     Guint16
		dither        GdkRgbDither
		bits          Guint
		// type : 1
		// expand : 1
	}

	GtkPreviewInfo struct {
		lookup *Guchar
		gamma  Gdouble
	}

	GtkTipsQuery struct {
		label GtkLabel
		bits  Guint
		// emit_always : 1
		// in_query : 1
		label_inactive *Gchar
		label_no_tip   *Gchar
		caller         *GtkWidget
		last_crossed   *GtkWidget
		query_cursor   *GdkCursor
	}

	GdkCursor struct {
		Type      GdkCursorType
		ref_count Guint
	}

	GtkLabel struct {
		misc  GtkMisc
		label *Gchar
		bits  Guint
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
		mnemonic_keyval Guint
		text            *Gchar
		attrs           *PangoAttrList
		effective_attrs *PangoAttrList
		layout          *PangoLayout
		mnemonic_widget *GtkWidget
		mnemonic_window *GtkWindow
		select_info     *GtkLabelSelectionInfo
	}

	GtkMisc struct {
		widget GtkWidget
		xalign Gfloat
		yalign Gfloat
		xpad   Guint16
		ypad   Guint16
	}

	GtkCList struct {
		container           GtkContainer
		flags               Guint16
		reserved1           Gpointer
		reserved2           Gpointer
		freeze_count        Guint
		internal_allocation GdkRectangle
		rows                Gint
		row_height          Gint
		row_list            *GList
		row_list_end        *GList
		columns             Gint
		column_title_area   GdkRectangle
		title_window        *GdkWindow
		column              *GtkCListColumn
		clist_window        *GdkWindow
		clist_window_width  Gint
		clist_window_height Gint
		hoffset             Gint
		voffset             Gint
		shadow_type         GtkShadowType
		selection_mode      GtkSelectionMode
		selection           *GList
		selection_end       *GList
		undo_selection      *GList
		undo_unselection    *GList
		undo_anchor         Gint
		button_actions      [5]Guint8
		drag_button         Guint8
		click_cell          GtkCListCellInfo
		hadjustment         *GtkAdjustment
		vadjustment         *GtkAdjustment
		xor_gc              *GdkGC
		fg_gc               *GdkGC
		bg_gc               *GdkGC
		cursor_drag         *GdkCursor
		x_drag              Gint
		focus_row           Gint
		focus_header_column Gint
		anchor              Gint
		anchor_state        GtkStateType
		drag_pos            Gint
		htimer              Gint
		vtimer              Gint
		sort_type           GtkSortType
		compare             GtkCListCompareFunc
		sort_column         Gint
		drag_highlight_row  Gint
		drag_highlight_pos  GtkCListDragPos
	}

	GtkCListColumn struct {
		title         *Gchar
		area          GdkRectangle
		button        *GtkWidget
		window        *GdkWindow
		width         Gint
		min_width     Gint
		max_width     Gint
		justification GtkJustification
		bits          Guint
		// visible : 1
		// width_set : 1
		// resizeable : 1
		// auto_resize : 1
		// button_passive : 1
	}

	GtkCListCellInfo struct {
		row    Gint
		column Gint
	}

	GtkAdjustment struct {
		parent_instance GtkObject
		lower           Gdouble
		upper           Gdouble
		value           Gdouble
		step_increment  Gdouble
		page_increment  Gdouble
		page_size       Gdouble
	}

	GtkCListRow struct {
		cell       *_GtkCell
		state      GtkStateType
		foreground GdkColor
		background GdkColor
		style      *GtkStyle
		data       Gpointer
		destroy    GDestroyNotify
		bits       Guint
		// fg_set : 1
		// bg_set : 1
		// selectable : 1
	}

	_GtkCell struct { // TODO(t):fix
		Type       GtkCellType
		vertical   Gint16
		horizontal Gint16
		style      *GtkStyle
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
		hbox            GtkHBox
		entry           *GtkWidget
		button          *GtkWidget
		popup           *GtkWidget
		popwin          *GtkWidget
		list            *GtkWidget
		entry_change_id Guint
		list_change_id  Guint
		bits            Guint
		// value_in_list:1
		// ok_if_empty:1
		// case_sensitive:1
		// use_arrows:1
		// use_arrows_always:1
		current_button Guint16
		activate_id    Guint
	}

	GtkHBox struct {
		box GtkBox
	}

	GtkBox struct {
		container   GtkContainer
		children    *GList
		spacing     Gint16
		homogeneous Guint // : 1
	}

	GtkCTree struct {
		clist        GtkCList
		lines_gc     *GdkGC
		tree_indent  Gint
		tree_spacing Gint
		tree_column  Gint
		bits         Guint
		// line_style : 2
		// expander_style : 2
		// show_stub : 1
		drag_compare GtkCTreeCompareDragFunc
	}

	GtkCTreeNode struct {
		list GList
	}

	GtkCTreeRow struct {
		row           GtkCListRow
		parent        *GtkCTreeNode
		sibling       *GtkCTreeNode
		children      *GtkCTreeNode
		pixmap_closed *GdkPixmap
		mask_closed   *GdkBitmap
		pixmap_opened *GdkPixmap
		mask_opened   *GdkBitmap
		level         Guint16
		bits          Guint
		// is_leaf : 1
		// expanded : 1
	}

	GtkTextBuffer struct {
		parent_instance            GObject
		tag_table                  *GtkTextTagTable
		btree                      *GtkTextBTree
		clipboard_contents_buffers *GSList
		selection_clipboards       *GSList
		log_attr_cache             *GtkTextLogAttrCache
		user_action_count          Guint
		bits                       Guint
		// modified : 1
		// has_selection : 1
	}

	GtkTextTagTable struct {
		parent_instance GObject
		hash            *GHashTable
		anonymous       *GSList
		anon_count      Gint
		buffers         *GSList
	}

	GtkTextView struct {
		parent_instance        GtkContainer
		layout                 *GtkTextLayout
		buffer                 *GtkTextBuffer
		selection_drag_handler Guint
		scroll_timeout         Guint
		pixels_above_lines     Gint
		pixels_below_lines     Gint
		pixels_inside_wrap     Gint
		wrap_mode              GtkWrapMode
		justify                GtkJustification
		left_margin            Gint
		right_margin           Gint
		indent                 Gint
		tabs                   *PangoTabArray
		bits                   Guint
		// editable : 1
		// overwrite_mode : 1
		// cursor_visible : 1
		// need_im_reset : 1
		// accepts_tab : 1
		// width_changed : 1
		// onscreen_validated : 1
		// mouse_cursor_obscured : 1
		text_window                 *GtkTextWindow
		left_window                 *GtkTextWindow
		right_window                *GtkTextWindow
		top_window                  *GtkTextWindow
		bottom_window               *GtkTextWindow
		hadjustment                 *GtkAdjustment
		vadjustment                 *GtkAdjustment
		xoffset                     Gint
		yoffset                     Gint
		width                       Gint
		height                      Gint
		virtual_cursor_x            Gint
		virtual_cursor_y            Gint
		first_para_mark             *GtkTextMark
		first_para_pixels           Gint
		dnd_mark                    *GtkTextMark
		blink_timeout               Guint
		first_validate_idle         Guint
		incremental_validate_idle   Guint
		im_context                  *GtkIMContext
		popup_menu                  *GtkWidget
		drag_start_x                Gint
		drag_start_y                Gint
		children                    *GSList
		pending_scroll              *GtkTextPendingScroll
		pending_place_cursor_button Gint
	}

	GtkTextMark struct {
		parent_instance GObject
		segment         Gpointer
	}

	GtkIMContext struct {
		parent_instance GObject
	}

	GtkTextIter struct {
		dummy1  Gpointer
		dummy2  Gpointer
		dummy3  Gint
		dummy4  Gint
		dummy5  Gint
		dummy6  Gint
		dummy7  Gint
		dummy8  Gint
		dummy9  Gpointer
		dummy10 Gpointer
		dummy11 Gint
		dummy12 Gint
		dummy13 Gint
		dummy14 Gpointer
	}

	GdkEventKey struct {
		Type             GdkEventType
		window           *GdkWindow
		send_event       Gint8
		time             Guint32
		state            Guint
		keyval           Guint
		length           Gint
		string           *Gchar
		hardware_keycode Guint16
		group            Guint8
		is_modifier      Guint // : 1
	}

	GtkTextChildAnchor struct {
		parent_instance GObject
		segment         Gpointer
	}

	GtkTextAttributes struct {
		refcount           Guint
		appearance         GtkTextAppearance
		justification      GtkJustification
		direction          GtkTextDirection
		font               *PangoFontDescription
		font_scale         Gdouble
		left_margin        Gint
		indent             Gint
		right_margin       Gint
		pixels_above_lines Gint
		pixels_below_lines Gint
		pixels_inside_wrap Gint
		tabs               *PangoTabArray
		wrap_mode          GtkWrapMode
		language           *PangoLanguage
		pg_bg_color        *GdkColor
		bits               Guint
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
		bg_color   GdkColor
		fg_color   GdkColor
		bg_stipple *GdkBitmap
		fg_stipple *GdkBitmap
		rise       Gint
		padding1   Gpointer
		bits       Guint
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
		misc               GtkMisc
		pixmap             *GdkPixmap
		mask               *GdkBitmap
		pixmap_insensitive *GdkPixmap
		build_insensitive  Guint // : 1
	}

	GtkToolbar struct {
		container    GtkContainer
		num_children Gint
		children     *GList
		orientation  GtkOrientation
		style        GtkToolbarStyle
		icon_size    GtkIconSize
		tooltips     *GtkTooltips
		button_maxw  Gint
		button_maxh  Gint
		_, _         Guint
		bits         Guint
		// style_set : 1
		// icon_size_set : 1
	}

	GtkTooltips struct {
		parent_instance  GtkObject
		tip_window       *GtkWidget
		tip_label        *GtkWidget
		active_tips_data *GtkTooltipsData
		tips_data_list   *GList
		bits, bits2      Guint //TODO(t): 33 bits Alignment/size?
		// delay : 30
		// enabled : 1
		// have_grab : 1
		// use_sticky_delay : 1
		timer_tag    Gint
		last_popdown GTimeVal
	}

	GtkTooltipsData struct {
		tooltips    *GtkTooltips
		widget      *GtkWidget
		tip_text    *Gchar
		tip_private *Gchar
	}

	GtkToolItem struct {
		parent GtkBin
		priv   *GtkToolItemPrivate
	}

	GtkToolItemGroup struct {
		parent_instance GtkContainer
		priv            *GtkToolItemGroupPrivate
	}

	GtkToolPalette struct {
		parent_instance GtkContainer
		priv            *GtkToolPalettePrivate
	}

	GtkSelectionData struct {
		selection GdkAtom
		target    GdkAtom
		Type      GdkAtom
		format    Gint
		data      *Guchar
		length    Gint
		display   *GdkDisplay
	}

	GdkDisplay struct {
		parent_instance   GObject
		queued_events     *GList
		queued_tail       *GList
		button_click_time [2]Guint32
		button_window     [2]*GdkWindow
		button_number     [2]Gint
		Double_click_time Guint
		core_pointer      *GdkDevice
		pointer_hooks     *GdkDisplayPointerHooks
		bits              Guint
		// closed : 1
		// ignore_core_events : 1
		Double_click_distance Guint
		button_x              [2]Gint
		button_y              [2]Gint
		pointer_grabs         *GList
		keyboard_grab         GdkKeyboardGrabInfo
		pointer_info          GdkPointerWindowInfo
		last_event_time       Guint32
	}

	GdkKeyboardGrabInfo struct {
		window        *GdkWindow
		native_window *GdkWindow
		serial        Gulong
		owner_events  Gboolean
		time          Guint32
	}

	GdkPointerWindowInfo struct {
		toplevel_under_pointer *GdkWindow
		window_under_pointer   *GdkWindow
		toplevel_x, toplevel_y Gdouble
		state                  Guint32
		button                 Guint32
		motion_hint_serial     Gulong
	}

	GdkDevice struct {
		parent_instance GObject
		name            *Gchar
		source          GdkInputSource
		mode            GdkInputMode
		has_cursor      Gboolean
		num_axes        Gint
		axes            *GdkDeviceAxis
		num_keys        Gint
		keys            *GdkDeviceKey
	}

	GdkDeviceKey struct {
		keyval    Guint
		modifiers GdkModifierType
	}

	GdkDisplayPointerHooks struct {
		get_pointer func(
			display *GdkDisplay,
			screen **GdkScreen,
			x *Gint,
			y *Gint,
			mask *GdkModifierType)
		window_get_pointer func(
			display *GdkDisplay,
			window *GdkWindow,
			x *Gint,
			y *Gint,
			mask *GdkModifierType) *GdkWindow
		window_at_pointer func(
			display *GdkDisplay,
			win_x *Gint,
			win_y *Gint) *GdkWindow
	}

	GdkDeviceAxis struct {
		use GdkAxisUse
		min Gdouble
		max Gdouble
	}

	GtkTargetEntry struct {
		target *Gchar
		flags  Guint
		info   Guint
	}

	GtkSizeGroup struct {
		parent_instance GObject
		widgets         *GSList
		mode            Guint8
		bits            Guint
		// have_width : 1
		// have_height : 1
		// ignore_hidden : 1
		requisition GtkRequisition
	}

	GtkSpinButton struct {
		entry         GtkEntry
		adjustment    *GtkAdjustment
		panel         *GdkWindow
		timer         Guint32
		climb_rate    Gdouble
		timer_step    Gdouble
		update_policy GtkSpinButtonUpdatePolicy
		bits          Guint
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
		widget GtkWidget
		text   *Gchar
		bits   Guint
		// editable : 1
		// visible : 1
		// overwrite_mode : 1
		// in_drag : 1
		text_length     Guint16
		text_max_length Guint16
		text_area       *GdkWindow
		im_context      *GtkIMContext
		popup_menu      *GtkWidget
		current_pos     Gint
		selection_bound Gint
		cached_layout   *PangoLayout
		bits2           Guint
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
		button         Guint
		blink_timeout  Guint
		recompute_idle Guint
		scroll_offset  Gint
		ascent         Gint
		descent        Gint
		x_text_size    Guint16
		x_n_bytes      Guint16
		preedit_length Guint16
		preedit_cursor Guint16
		dnd_position   Gint
		drag_start_x   Gint
		drag_start_y   Gint
		invisible_char Gunichar
		width_chars    Gint
	}

	GtkTreeModelSort struct {
		parent               GObject
		root                 Gpointer
		stamp                Gint
		child_flags          Guint
		child_model          *GtkTreeModel
		zero_ref_count       Gint
		sort_list            *GList
		sort_column_id       Gint
		order                GtkSortType
		default_sort_func    GtkTreeIterCompareFunc
		default_sort_data    Gpointer
		default_sort_destroy GDestroyNotify
		changed_id           Guint
		inserted_id          Guint
		has_child_toggled_id Guint
		deleted_id           Guint
		reordered_id         Guint
	}

	GtkTreeIter struct {
		stamp      Gint
		user_data  Gpointer
		user_data2 Gpointer
		user_data3 Gpointer
	}

	GtkTreeSelection struct {
		parent    GObject
		tree_view *GtkTreeView
		Type      GtkSelectionMode
		user_func GtkTreeSelectionFunc
		user_data Gpointer
		destroy   GDestroyNotify
	}

	GtkTreeView struct {
		parent GtkContainer
		priv   *GtkTreeViewPrivate
	}

	GtkTreeStore struct {
		parent               GObject
		stamp                Gint
		root                 Gpointer
		last                 Gpointer
		n_columns            Gint
		sort_column_id       Gint
		sort_list            *GList
		order                GtkSortType
		column_headers       *GType
		default_sort_func    GtkTreeIterCompareFunc
		default_sort_data    Gpointer
		default_sort_destroy GDestroyNotify
		columns_dirty        Guint // : 1
	}

	GtkUIManager struct {
		parent       GObject
		private_data *GtkUIManagerPrivate
	}

	GtkActionGroup struct {
		parent       GObject
		private_data *GtkActionGroupPrivate
	}

	GtkAction struct {
		object       GObject
		private_data *GtkActionPrivate
	}

	GtkHSV struct {
		parent_instance GtkWidget
		priv            Gpointer
	}

	GtkIconFactory struct {
		parent_instance GObject
		icons           *GHashTable
	}

	GtkSettings struct {
		parent_instance GObject
		queued_settings *GData
		property_values *GtkSettingsPropertyValue
		rc_context      *GtkRcContext
		screen          *GdkScreen
	}

	GtkIconTheme struct {
		parent_instance GObject
		priv            *GtkIconThemePrivate
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
		parent GtkContainer
		priv   *GtkIconViewPrivate
	}

	GtkCellRenderer struct {
		parent GtkObject
		xalign Gfloat
		yalign Gfloat
		width  Gint
		height Gint
		xpad   Guint16
		ypad   Guint16
		bits   Guint
		// mode : 2;
		// visible : 1
		// is_expander : 1
		// is_expanded : 1
		// cell_background_set : 1
		// sensitive : 1
		// editing : 1
	}

	GtkImageMenuItem struct {
		menu_item GtkMenuItem
		image     *GtkWidget
	}

	GtkMenuItem struct {
		item              GtkItem
		submenu           *GtkWidget
		event_window      *GdkWindow
		toggle_size       Guint16
		accelerator_width Guint16
		accel_path        *Gchar
		bits              Guint
		// show_submenu_indicator : 1
		// submenu_placement : 1
		// submenu_direction : 1
		// right_justify: 1
		// timer_from_keypress : 1
		// from_menubar : 1
		timer Guint
	}

	GtkIMContextSimple struct {
		object              GtkIMContext
		tables              *GSList
		compose_buffer      [7 + 1]Guint
		tentative_match     Gunichar
		tentative_match_len Gint
		bits                Guint
		// in_hex_sequence : 1
		// modifiers_dropped : 1
	}

	GtkIMMulticontext struct {
		object     GtkIMContext
		slave      *GtkIMContext
		priv       *GtkIMMulticontextPrivate
		context_id *Gchar
	}

	GtkMenuShell struct {
		container         GtkContainer
		children          *GList
		active_menu_item  *GtkWidget
		parent_menu_shell *GtkWidget
		button            Guint
		activate_time     Guint32
		bits              Guint
		// active : 1
		// have_grab : 1
		// have_xgrab : 1
		// ignore_leave : 1
		// menu_flag : 1
		// ignore_enter : 1
		// keyboard_mode : 1
	}

	GtkInfoBar struct {
		parent GtkHBox
		priv   *GtkInfoBarPrivate
	}

	GtkInvisible struct {
		widget             GtkWidget
		has_user_ref_count Gboolean
		screen             *GdkScreen
	}

	GtkLayout struct {
		container    GtkContainer
		children     *GList
		width        Guint
		height       Guint
		hadjustment  *GtkAdjustment
		vadjustment  *GtkAdjustment
		bin_window   *GdkWindow
		visibility   GdkVisibilityState
		scroll_x     Gint
		scroll_y     Gint
		freeze_count Guint
	}

	GtkLinkButton struct {
		parent_instance GtkButton
		priv            *GtkLinkButtonPrivate
	}

	GtkArg struct {
		Type GType
		name *Gchar
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
		signal_id Guint
		detail    GQuark
		run_type  GSignalFlags
	}

	GtkMenuBar struct {
		menu_shell GtkMenuShell
	}

	GtkToolButton struct {
		parent GtkToolItem
		priv   *GtkToolButtonPrivate
	}

	GtkMenuToolButton struct {
		parent GtkToolButton
		priv   *GtkMenuToolButtonPrivate
	}

	GtkMessageDialog struct {
		parent_instance GtkDialog
		image           *GtkWidget
		label           *GtkWidget
	}

	GMountOperation struct {
		parent_instance GObject
		priv            *GMountOperationPrivate
	}

	GtkMountOperation struct {
		parent_instance GMountOperation
		priv            *GtkMountOperationPrivate
	}

	GtkNotebook struct {
		container    GtkContainer
		cur_page     *GtkNotebookPage
		children     *GList
		first_tab    *GList
		focus_tab    *GList
		menu         *GtkWidget
		event_window *GdkWindow
		timer        Guint32
		tab_hborder  Guint16
		tab_vborder  Guint16
		bits         Guint
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
		parent_object GtkWindow
	}

	GtkSocket struct {
		container      GtkContainer
		request_width  Guint16
		request_height Guint16
		current_width  Guint16
		current_height Guint16
		plug_window    *GdkWindow
		plug_widget    *GtkWidget
		xembed_version Gshort
		bits           Guint
		// same_app : 1
		// focus_in : 1
		// have_size : 1
		// need_map : 1
		// is_mapped : 1
		// active : 1
		accel_group *GtkAccelGroup
		toplevel    *GtkWidget
	}

	GtkPlug struct {
		window          GtkWindow
		socket_window   *GdkWindow
		modality_window *GtkWidget
		modality_group  *GtkWindowGroup
		grabbed_keys    *GHashTable
		same_app        Guint // : 1
	}

	GtkPageRange struct {
		start Gint
		end   Gint
	}

	GtkPrintOperation struct {
		parent_instance GObject
		priv            *GtkPrintOperationPrivate
	}

	GtkProgress struct {
		widget           GtkWidget
		adjustment       *GtkAdjustment
		offscreen_pixmap *GdkPixmap
		format           *Gchar
		x_align          Gfloat
		y_align          Gfloat
		bits             Guint
		// show_text : 1;
		// activity_mode : 1;
		// use_text_format : 1
	}

	GtkProgressBar struct {
		progress        GtkProgress
		bar_style       GtkProgressBarStyle
		orientation     GtkProgressBarOrientation
		blocks          Guint
		in_block        Gint
		activity_pos    Gint
		activity_step   Guint
		activity_blocks Guint
		pulse_fraction  Gdouble
		bits            Guint
		// activity_dir : 1
		// ellipsize : 3
		// dirty : 1
	}

	GtkToggleAction struct {
		parent       GtkAction
		private_data *GtkToggleActionPrivate
	}

	GtkRadioAction struct {
		parent       GtkToggleAction
		private_data *GtkRadioActionPrivate
	}

	GtkRadioButton struct {
		check_button GtkCheckButton
		group        *GSList
	}

	GtkRadioMenuItem struct {
		check_menu_item GtkCheckMenuItem
		group           *GSList
	}

	GtkCheckButton struct {
		toggle_button GtkToggleButton
	}

	GtkCheckMenuItem struct {
		menu_item GtkMenuItem
		bits      Guint
		// active : 1
		// always_show_toggle : 1
		// inconsistent : 1
		// draw_as_radio : 1
	}

	GtkToggleButton struct {
		button GtkButton
		bits   Guint
		// active : 1
		// draw_indicator : 1
		// inconsistent : 1
	}

	GtkToggleToolButton struct {
		parent GtkToolButton
		priv   *GtkToggleToolButtonPrivate
	}

	GtkRadioToolButton struct {
		parent GtkToggleToolButton
	}

	GtkRecentManager struct {
		parent_instance GObject
		priv            *GtkRecentManagerPrivate
	}

	GtkRecentData struct {
		display_name *Gchar
		description  *Gchar
		mime_type    *Gchar
		app_name     *Gchar
		app_exec     *Gchar
		groups       **Gchar
		is_private   Gboolean
	}

	GtkRecentAction struct {
		parent_instance GtkAction
		priv            *GtkRecentActionPrivate
	}

	GtkScaleButton struct {
		parent       GtkButton
		plus_button  *GtkWidget
		minus_button *GtkWidget
		priv         *GtkScaleButtonPrivate
	}

	GtkRecentFilterInfo struct {
		contains     GtkRecentFilterFlags
		uri          *Gchar
		display_name *Gchar
		mime_type    *Gchar
		applications **Gchar
		groups       **Gchar
		age          Gint
	}

	GtkViewport struct {
		bin         GtkBin
		shadow_type GtkShadowType
		view_window *GdkWindow
		bin_window  *GdkWindow
		hadjustment *GtkAdjustment
		vadjustment *GtkAdjustment
	}

	GtkScrolledWindow struct {
		container  GtkBin
		hscrollbar *GtkWidget
		vscrollbar *GtkWidget
		bits       Guint
		// hscrollbar_policy : 2
		// vscrollbar_policy : 2
		// hscrollbar_visible : 1
		// vscrollbar_visible : 1
		// window_placement : 2
		// focus_out : 1
		shadow_type Guint16
	}

	GtkRecentChooserMenu struct {
		parent_instance GtkMenu
		priv            *GtkRecentChooserMenuPrivate
	}

	GtkMenu struct {
		menu_shell           GtkMenuShell
		parent_menu_item     *GtkWidget
		old_active_menu_item *GtkWidget
		accel_group          *GtkAccelGroup
		accel_path           *Gchar
		position_func        GtkMenuPositionFunc
		position_func_data   Gpointer
		toggle_size          Guint
		toplevel             *GtkWidget
		tearoff_window       *GtkWidget
		tearoff_hbox         *GtkWidget
		tearoff_scrollbar    *GtkWidget
		tearoff_adjustment   *GtkAdjustment
		view_window          *GdkWindow
		bin_window           *GdkWindow
		scroll_offset        Gint
		saved_scroll_offset  Gint
		scroll_step          Gint
		timeout_id           Guint
		navigation_region    *GdkRegion
		navigation_timeout   Guint
		bits                 Guint
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
		parent GtkToolItem
		priv   *GtkSeparatorToolItemPrivate
	}

	GtkStatusbar struct {
		parent_widget   GtkHBox
		frame           *GtkWidget
		label           *GtkWidget
		messages        *GSList
		keys            *GSList
		seq_context_id  Guint
		seq_message_id  Guint
		grip_window     *GdkWindow
		has_resize_grip Guint // : 1
	}

	GtkSpinner struct {
		parent GtkDrawingArea
		priv   *GtkSpinnerPrivate
	}

	GtkStatusIcon struct {
		parent_instance GObject
		priv            *GtkStatusIconPrivate
	}

	GtkStockItem struct {
		stock_id           *Gchar
		label              *Gchar
		modifier           GdkModifierType
		keyval             Guint
		translation_domain *Gchar
	}

	GtkTable struct {
		container      GtkContainer
		children       *GList
		rows           *GtkTableRowCol
		cols           *GtkTableRowCol
		nrows          Guint16
		ncols          Guint16
		column_spacing Guint16
		row_spacing    Guint16
		bits           Guint
		// homogeneous : 1
	}

	GtkTableRowCol struct {
		requisition Guint16
		allocation  Guint16
		spacing     Guint16
		bits        Guint
		// need_expand : 1
		// need_shrink : 1
		// expand : 1
		// shrink : 1
		// empty : 1
	}

	GtkTextTag struct {
		parent_instance GObject
		table           *GtkTextTagTable
		name            *Char
		priority        int
		values          *GtkTextAttributes
		bits            Guint
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
		list      *GList
		ref_count Guint
	}

	PangoLogAttr struct {
		bits Guint
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
		parent            GObject
		description       *Gchar
		name              *Gchar
		accessible_parent *AtkObject
		role              AtkRole
		relation_set      *AtkRelationSet
		layer             AtkLayer
	}

	AtkRelationSet struct {
		parent    GObject
		relations *GPtrArray
	}

	GdkDragContext struct {
		parent_instance  GObject
		protocol         GdkDragProtocol
		is_source        Gboolean
		source_window    *GdkWindow
		dest_window      *GdkWindow
		targets          *GList
		actions          GdkDragAction
		suggested_action GdkDragAction
		action           GdkDragAction
		start_time       Guint32
		windowing_data   Gpointer
	}

	GdkEventExpose struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		area       GdkRectangle
		region     *GdkRegion
		count      Gint
	}

	GdkEventOwnerChange struct {
		Type           GdkEventType
		window         *GdkWindow
		send_event     Gint8
		owner          GdkNativeWindow
		reason         GdkOwnerChange
		selection      GdkAtom
		time           Guint32
		selection_time Guint32
	}

	GdkEventProperty struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		atom       GdkAtom
		time       Guint32
		state      Guint
	}

	GdkEventSelection struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		selection  GdkAtom
		target     GdkAtom
		property   GdkAtom
		time       Guint32
		requestor  GdkNativeWindow
	}

	GdkEventSetting struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		action     GdkSettingAction
		name       *Char
	}

	GdkGCValues struct {
		foreground         GdkColor
		background         GdkColor
		font               *GdkFont
		function           GdkFunction
		fill               GdkFill
		tile               *GdkPixmap
		stipple            *GdkPixmap
		clip_mask          *GdkPixmap
		subwindow_mode     GdkSubwindowMode
		ts_x_origin        Gint
		ts_y_origin        Gint
		clip_x_origin      Gint
		clip_y_origin      Gint
		graphics_exposures Gint
		line_width         Gint
		line_style         GdkLineStyle
		cap_style          GdkCapStyle
		join_style         GdkJoinStyle
	}

	GdkGeometry struct {
		min_width   Gint
		min_height  Gint
		max_width   Gint
		max_height  Gint
		base_width  Gint
		base_height Gint
		width_inc   Gint
		height_inc  Gint
		min_aspect  Gdouble
		max_aspect  Gdouble
		win_gravity GdkGravity
	}

	GdkImage struct {
		parent_instance GObject
		Type            GdkImageType
		visual          *GdkVisual
		byte_order      GdkByteOrder
		width           Gint
		height          Gint
		depth           Guint16
		bpp             Guint16
		bpl             Guint16
		bits_per_pixel  Guint16
		mem             Gpointer
		colormap        *GdkColormap
		windowing_data  Gpointer
	}

	GObjectClass struct {
		g_type_class         GTypeClass
		construct_properties *GSList
		constructor          func(
			Type GType,
			n_construct_properties Guint,
			construct_properties *GObjectConstructParam) *GObject
		set_property func(
			object *GObject,
			property_id Guint,
			GValue,
			pspec *GParamSpec)
		get_property func(
			object *GObject,
			property_id Guint,
			value *GValue,
			pspec *GParamSpec)
		dispose func(
			object *GObject)
		finalize func(
			object *GObject)
		dispatch_properties_changed func(
			object *GObject,
			n_pspecs Guint,
			pspecs **GParamSpec)
		notify func(
			object *GObject,
			pspec *GParamSpec)
		constructed func(
			object *GObject)
		flags  Gsize
		pdummy [6]Gpointer
	}

	GParamSpec struct {
		g_type_instance GTypeInstance
		name            *Gchar
		flags           GParamFlags
		value_type      GType
		owner_type      GType
		_nick           *Gchar
		_blurb          *Gchar
		qdata           *GData
		ref_count       Guint
		param_id        Guint
	}

	GObjectConstructParam struct {
		pspec *GParamSpec
		value *GValue
	}

	GtkAboutDialog struct {
		parent_instance GtkDialog
		private_data    Gpointer
	}

	GtkAccelLabel struct {
		label              GtkLabel
		_                  Guint
		accel_padding      Guint
		accel_widget       *GtkWidget
		accel_closure      *GClosure
		accel_group        *GtkAccelGroup
		accel_string       *Gchar
		accel_string_width Guint16
	}

	GtkAccelLabelClass struct {
		parent_class     GtkLabelClass
		signal_quote1    *Gchar
		signal_quote2    *Gchar
		mod_name_shift   *Gchar
		mod_name_control *Gchar
		mod_name_alt     *Gchar
		mod_separator    *Gchar
		accel_seperator  *Gchar
		latin1_to_char   Guint // : 1
		_, _, _, _       func()
	}

	GtkLabelClass struct {
		parent_class GtkMiscClass
		move_cursor  func(
			label *GtkLabel,
			step GtkMovementStep,
			count Gint,
			extend_selection Gboolean)
		copy_clipboard func(
			label *GtkLabel)
		populate_popup func(
			label *GtkLabel,
			menu *GtkMenu)
		activate_link func(
			label *GtkLabel,
			uri *Gchar) Gboolean
		_, _, _ func()
	}

	GtkMiscClass struct {
		parent_class GtkWidgetClass
	}

	GtkWidgetClass struct {
		parent_class                      GtkObjectClass
		activate_signal                   Guint
		set_scroll_adjustments_signal     Guint
		dispatch_child_properties_changed func(
			widget *GtkWidget,
			n_pspecs Guint,
			pspecs **GParamSpec)
		show         func(widget *GtkWidget)
		show_all     func(widget *GtkWidget)
		hide         func(widget *GtkWidget)
		hide_all     func(widget *GtkWidget)
		Map          func(widget *GtkWidget)
		unmap        func(widget *GtkWidget)
		realize      func(widget *GtkWidget)
		unrealize    func(widget *GtkWidget)
		size_request func(widget *GtkWidget,
			requisition *GtkRequisition)
		size_allocate func(widget *GtkWidget,
			allocation *GtkAllocation)
		state_changed func(widget *GtkWidget,
			previous_state GtkStateType)
		parent_set func(widget *GtkWidget,
			previous_parent *GtkWidget)
		hierarchy_changed func(widget *GtkWidget,
			previous_toplevel *GtkWidget)
		style_set func(widget *GtkWidget,
			previous_style *GtkStyle)
		direction_changed func(widget *GtkWidget,
			previous_direction GtkTextDirection)
		grab_notify func(widget *GtkWidget,
			was_grabbed Gboolean)
		child_notify func(widget *GtkWidget,
			pspec *GParamSpec)
		mnemonic_activate func(widget *GtkWidget,
			group_cycling Gboolean) Gboolean
		grab_focus func(widget *GtkWidget)
		focus      func(widget *GtkWidget,
			direction GtkDirectionType) Gboolean
		event func(widget *GtkWidget,
			event *GdkEvent) Gboolean
		button_press_event func(widget *GtkWidget,
			event *GdkEventButton) Gboolean
		button_release_event func(widget *GtkWidget,
			event *GdkEventButton) Gboolean
		scroll_event func(widget *GtkWidget,
			event *GdkEventScroll) Gboolean
		motion_notify_event func(widget *GtkWidget,
			event *GdkEventMotion) Gboolean
		delete_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		destroy_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		expose_event func(widget *GtkWidget,
			event *GdkEventExpose) Gboolean
		key_press_event func(widget *GtkWidget,
			event *GdkEventKey) Gboolean
		key_release_event func(widget *GtkWidget,
			event *GdkEventKey) Gboolean
		enter_notify_event func(widget *GtkWidget,
			event *GdkEventCrossing) Gboolean
		leave_notify_event func(widget *GtkWidget,
			event *GdkEventCrossing) Gboolean
		configure_event func(widget *GtkWidget,
			event *GdkEventConfigure) Gboolean
		focus_in_event func(widget *GtkWidget,
			event *GdkEventFocus) Gboolean
		focus_out_event func(widget *GtkWidget,
			event *GdkEventFocus) Gboolean
		map_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		unmap_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		property_notify_event func(widget *GtkWidget,
			event *GdkEventProperty) Gboolean
		selection_clear_event func(widget *GtkWidget,
			event *GdkEventSelection) Gboolean
		selection_request_event func(widget *GtkWidget,
			event *GdkEventSelection) Gboolean
		selection_notify_event func(widget *GtkWidget,
			event *GdkEventSelection) Gboolean
		proximity_in_event func(widget *GtkWidget,
			event *GdkEventProximity) Gboolean
		proximity_out_event func(widget *GtkWidget,
			event *GdkEventProximity) Gboolean
		visibility_notify_event func(widget *GtkWidget,
			event *GdkEventVisibility) Gboolean
		client_event func(widget *GtkWidget,
			event *GdkEventClient) Gboolean
		no_expose_event func(widget *GtkWidget,
			event *GdkEventAny) Gboolean
		window_state_event func(widget *GtkWidget,
			event *GdkEventWindowState) Gboolean
		selection_get func(widget *GtkWidget,
			selection_data *GtkSelectionData,
			info, time_ Guint)
		selection_received func(widget *GtkWidget,
			selection_data *GtkSelectionData,
			time_ Guint)
		drag_begin func(widget *GtkWidget,
			context *GdkDragContext)
		drag_end func(widget *GtkWidget,
			context *GdkDragContext)
		drag_data_get func(widget *GtkWidget,
			context *GdkDragContext,
			selection_data *GtkSelectionData,
			info, time_ Guint)
		drag_data_delete func(widget *GtkWidget,
			context *GdkDragContext)
		drag_leave func(widget *GtkWidget,
			context *GdkDragContext,
			time_ Guint)
		drag_motion func(widget *GtkWidget,
			context *GdkDragContext,
			x, y Gint, time_ Guint) Gboolean
		drag_drop func(widget *GtkWidget,
			context *GdkDragContext,
			x, y Gint, time_ Guint) Gboolean
		drag_data_received func(widget *GtkWidget,
			context *GdkDragContext,
			x, y Gint,
			selection_data *GtkSelectionData,
			info, time_ Guint)
		popup_menu func(widget *GtkWidget) Gboolean
		show_help  func(widget *GtkWidget,
			help_type GtkWidgetHelpType) Gboolean
		get_accessible func(widget *GtkWidget) *AtkObject
		screen_changed func(widget *GtkWidget,
			previous_screen *GdkScreen) Gboolean
		can_activate_accel func(widget *GtkWidget,
			signal_id Guint) Gboolean
		grab_broken_event func(widget *GtkWidget,
			event *GdkEventGrabBroken) Gboolean
		composited_changed func(widget *GtkWidget)
		query_tooltip      func(widget *GtkWidget,
			x, y Gint, keyboard_tooltip Gboolean,
			tooltip *GtkTooltip) Gboolean
		_, _, _ func()
	}

	GtkObjectClass struct {
		parent_class GInitiallyUnownedClass
		set_arg      func(object *GtkObject,
			arg *GtkArg,
			arg_id Guint)
		get_arg func(object *GtkObject,
			arg *GtkArg,
			arg_id Guint)
		destroy func(object *GtkObject)
	}

	GdkEventButton struct {
		Type           GdkEventType
		window         *GdkWindow
		send_event     Gint8
		time           Guint32
		x              Gdouble
		y              Gdouble
		axes           *Gdouble
		state          Guint
		button         Guint
		device         *GdkDevice
		x_root, y_root Gdouble
	}

	GdkEventScroll struct {
		Type           GdkEventType
		window         *GdkWindow
		send_event     Gint8
		time           Guint32
		x              Gdouble
		y              Gdouble
		state          Guint
		direction      GdkScrollDirection
		device         *GdkDevice
		x_root, y_root Gdouble
	}

	GdkEventMotion struct {
		Type           GdkEventType
		window         *GdkWindow
		send_event     Gint8
		time           Guint32
		x              Gdouble
		y              Gdouble
		axes           *Gdouble
		state          Guint
		is_hint        Gint16
		device         *GdkDevice
		x_root, y_root Gdouble
	}

	GdkEventAny struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
	}

	GdkEventCrossing struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		subwindow  *GdkWindow
		time       Guint32
		x          Gdouble
		y          Gdouble
		x_root     Gdouble
		y_root     Gdouble
		mode       GdkCrossingMode
		detail     GdkNotifyType
		focus      Gboolean
		state      Guint
	}

	GdkEventConfigure struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		x, y       Gint
		width      Gint
		height     Gint
	}

	GdkEventFocus struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		in         Gint16
	}

	GdkEventProximity struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		time       Guint32
		device     *GdkDevice
	}

	GdkEventVisibility struct {
		Type       GdkEventType
		window     *GdkWindow
		send_event Gint8
		state      GdkVisibilityState
	}

	GdkEventClient struct {
		Type         GdkEventType
		window       *GdkWindow
		send_event   Gint8
		message_type GdkAtom
		data_format  Gushort
		// Union
		b [20]Char
		// s [10]Short
		// l [5]Long
	}

	GdkEventWindowState struct {
		Type             GdkEventType
		window           *GdkWindow
		send_event       Gint8
		changed_mask     GdkWindowState
		new_window_state GdkWindowState
	}

	GdkEventGrabBroken struct {
		Yype        GdkEventType
		window      *GdkWindow
		send_event  Gint8
		keyboard    Gboolean
		implicit    Gboolean
		grab_window *GdkWindow
	}

	GtkAccessible struct {
		parent AtkObject
		widget *GtkWidget
	}

	GtkActionEntry struct {
		name        *Gchar
		stock_id    *Gchar
		label       *Gchar
		accelerator *Gchar
		tooltip     *Gchar
		callback    GCallback
	}

	GtkAlignment struct {
		bin    GtkBin
		xalign Gfloat
		yalign Gfloat
		xscale Gfloat
		yscale Gfloat
	}

	GtkArrow struct {
		misc        GtkMisc
		arrow_type  Gint16
		shadow_type Gint16
	}

	GtkAspectFrame struct {
		frame             GtkFrame
		xalign            Gfloat
		yalign            Gfloat
		ratio             Gfloat
		obey_child        Gboolean
		center_allocation GtkAllocation
	}

	GtkAssistant struct {
		parent  GtkWindow
		cancel  *GtkWidget
		forward *GtkWidget
		back    *GtkWidget
		apply   *GtkWidget
		close   *GtkWidget
		last    *GtkWidget
		priv    *GtkAssistantPrivate
	}

	GtkBindingSet struct {
		set_name            *Gchar
		priority            Gint
		widget_path_pspecs  *GSList
		widget_class_pspecs *GSList
		class_branch_pspecs *GSList
		entries             *GtkBindingEntry
		current             *GtkBindingEntry
		parsed              Guint // : 1
	}

	GtkBorder struct {
		left   Gint
		right  Gint
		top    Gint
		bottom Gint
	}

	GtkBuilder struct {
		parent_instance GObject
		priv            *GtkBuilderPrivate
	}

	GtkButtonBox struct {
		box              GtkBox
		child_min_width  Gint
		child_min_height Gint
		child_ipad_x     Gint
		child_ipad_y     Gint
		layout_style     GtkButtonBoxStyle
	}

	GtkCalendar struct {
		widget            GtkWidget
		header_style      *GtkStyle
		label_style       *GtkStyle
		month             Gint
		year              Gint
		selected_day      Gint
		day_month         [6][7]Gint
		day               [6][7]Gint
		num_marked_dates  Gint
		marked_date       [31]Gint
		display_flags     GtkCalendarDisplayOptions
		marked_date_color [31]GdkColor
		gc                *GdkGC
		xor_gc            *GdkGC
		focus_row         Gint
		focus_col         Gint
		highlight_row     Gint
		highlight_col     Gint
		priv              *GtkCalendarPrivate
		grow_space        [32]Gchar
		_, _, _, _        func()
	}

	GtkCellRendererText struct {
		parent            GtkCellRenderer
		text              *Gchar
		font              *PangoFontDescription
		font_scale        Gdouble
		foreground        PangoColor
		background        PangoColor
		extra_attrs       *PangoAttrList
		underline_style   PangoUnderline
		rise              Gint
		fixed_height_rows Gint
		bits              Guint
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
		parent GtkCellRenderer
		bits   Guint
		// active : 1;
		// activatable : 1;
		// radio : 1;
	}

	GtkCellView struct {
		parent_instance GtkWidget
		priv            *GtkCellViewPrivate
	}

	GtkColorButton struct {
		button GtkButton
		priv   *GtkColorButtonPrivate
	}

	GtkColorSelection struct {
		parent_instance GtkVBox
		private_data    Gpointer
	}

	GtkColorSelectionDialog struct {
		parent_instance GtkDialog
		colorsel        *GtkWidget
		ok_button       *GtkWidget
		cancel_button   *GtkWidget
		help_button     *GtkWidget
	}

	GtkComboBox struct {
		parent_instance GtkBin
		priv            *GtkComboBoxPrivate
	}

	GtkComboBoxEntry struct {
		parent_instance GtkComboBox
		priv            *GtkComboBoxEntryPrivate
	}

	GtkComboBoxText struct {
		parent_instance GtkComboBox
		priv            *GtkComboBoxTextPrivate
	}

	GtkContainerClass struct {
		parent_class GtkWidgetClass
		add          func(container *GtkContainer,
			widget *GtkWidget)
		remove func(container *GtkContainer,
			widget *GtkWidget)
		check_resize func(container *GtkContainer)
		forall       func(container *GtkContainer,
			include_internals Gboolean,
			callback GtkCallback,
			callback_data Gpointer)
		set_focus_child func(container *GtkContainer,
			widget *GtkWidget)
		child_type     func(container *GtkContainer) GType
		composite_name func(container *GtkContainer,
			child *GtkWidget) *Gchar
		set_child_property func(container *GtkContainer,
			child *GtkWidget,
			property_id Guint,
			value *GValue,
			pspec *GParamSpec)
		get_child_property func(container *GtkContainer,
			child *GtkWidget,
			property_id Guint,
			value *GValue,
			pspec *GParamSpec)
		_, _, _, _ func()
	}

	GtkEntryBuffer struct {
		parent_instance GObject
		priv            *GtkEntryBufferPrivate
	}

	GtkEntryCompletion struct {
		parent_instance GObject
		priv            *GtkEntryCompletionPrivate
	}

	GtkEventBox struct {
		bin GtkBin
	}

	GtkExpander struct {
		bin  GtkBin
		priv *GtkExpanderPrivate
	}

	GtkFileChooserButton struct {
		parent GtkHBox
		priv   *GtkFileChooserButtonPrivate
	}

	GtkFileFilterInfo struct {
		contains     GtkFileFilterFlags
		filename     *Gchar
		uri          *Gchar
		display_name *Gchar
		mime_type    *Gchar
	}

	GtkFixed struct {
		container GtkContainer
		children  *GList
	}

	GtkFontButton struct {
		button GtkButton
		priv   *GtkFontButtonPrivate
	}

	GtkFontSelection struct {
		parent_instance  GtkVBox
		font_entry       *GtkWidget
		family_list      *GtkWidget
		font_style_entry *GtkWidget
		face_list        *GtkWidget
		size_entry       *GtkWidget
		size_list        *GtkWidget
		pixels_button    *GtkWidget
		points_button    *GtkWidget
		filter_button    *GtkWidget
		preview_entry    *GtkWidget
		family           *PangoFontFamily
		face             *PangoFontFace
		size             Gint
		font             *GdkFont
	}

	GtkFontSelectionDialog struct {
		parent_instance GtkDialog
		fontsel         *GtkWidget
		main_vbox       *GtkWidget
		action_area     *GtkWidget
		ok_button       *GtkWidget
		apply_button    *GtkWidget
		cancel_button   *GtkWidget
		dialog_width    Gint
		auto_resize     Gboolean
	}

	GtkFrame struct {
		bin              GtkBin
		label_widget     *GtkWidget
		shadow_type      Gint16
		label_xalign     Gfloat
		label_yalign     Gfloat
		child_allocation GtkAllocation
	}

	GtkHandleBox struct {
		bin          GtkBin
		bin_window   *GdkWindow
		float_window *GdkWindow
		shadow_type  GtkShadowType
		bits         Guint
		// handle_position : 2;
		// float_window_mapped : 1;
		// child_detached : 1;
		// in_drag : 1;
		// shrink_on_detach : 1;
		// snap_edge : 3 // signed
		deskoff_x         Gint
		deskoff_y         Gint
		attach_allocation GtkAllocation
		float_allocation  GtkAllocation
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
		parent               GObject
		stamp                Gint
		seq                  Gpointer
		_                    Gpointer
		sort_list            *GList
		n_columns            Gint
		sort_column_id       Gint
		order                GtkSortType
		column_headers       *GType
		length               Gint
		default_sort_func    GtkTreeIterCompareFunc
		default_sort_data    Gpointer
		default_sort_destroy GDestroyNotify
		columns_dirty        Guint // : 1
	}

	GtkPaned struct {
		container       GtkContainer
		child1          *GtkWidget
		child2          *GtkWidget
		handle          *GdkWindow
		xor_gc          *GdkGC
		cursor_type     GdkCursorType
		handle_pos      GdkRectangle
		child1_size     Gint
		last_allocation Gint
		min_position    Gint
		max_position    Gint
		bits            Guint
		// position_set : 1
		// in_drag : 1
		// child1_shrink : 1
		// child1_resize : 1
		// child2_shrink : 1
		// child2_resize : 1
		// orientation : 1
		// in_recursion : 1
		// handle_prelit : 1
		last_child1_focus *GtkWidget
		last_child2_focus *GtkWidget
		priv              *GtkPanedPrivate
		drag_pos          Gint
		original_position Gint
	}

	GtkToggleActionEntry struct {
		name        *Gchar
		stock_id    *Gchar
		label       *Gchar
		accelerator *Gchar
		tooltip     *Gchar
		callback    GCallback
		is_active   Gboolean
	}

	GtkRange struct {
		widget        GtkWidget
		adjustment    *GtkAdjustment
		update_policy GtkUpdateType
		bits          Guint
		// inverted : 1
		// flippable : 1
		// has_stepper_a : 1
		// has_stepper_b : 1
		// has_stepper_c : 1
		// has_stepper_d : 1
		// need_recalc : 1
		// slider_size_fixed : 1
		min_slider_size Gint
		orientation     GtkOrientation
		range_rect      GdkRectangle
		slider_start    Gint
		slider_end      Gint
		round_digits    Gint
		bits2           Guint
		// trough_click_forward : 1
		// update_pending : 1
		layout                        *GtkRangeLayout
		timer                         *GtkRangeStepTimer
		slide_initial_slider_position Gint
		slide_initial_coordinate      Gint
		update_timeout_id             Guint
		event_window                  *GdkWindow
	}

	GtkRcProperty struct {
		type_name     GQuark
		property_name GQuark
		origin        *Gchar
		value         GValue
	}

	GtkRuler struct {
		widget        GtkWidget
		backing_store *GdkPixmap
		non_gr_exp_gc *GdkGC
		metric        *GtkRulerMetric
		xsrc          Gint
		ysrc          Gint
		slider_size   Gint
		lower         Gdouble
		upper         Gdouble
		position      Gdouble
		max_size      Gdouble
	}

	GtkScale struct {
		Range  GtkRange
		digits Gint
		bits   Guint
		// draw_value : 1
		// value_pos : 2
	}

	GtkSettingsValue struct {
		origin *Gchar
		value  GValue
	}

	GtkTreeModelFilter struct {
		parent GObject
		priv   *GtkTreeModelFilterPrivate
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
		type_name            *Gchar
		object_size          Guint
		class_size           Guint
		class_init_func      GtkClassInitFunc
		object_init_func     GtkObjectInitFunc
		reserved_1           Gpointer
		reserved_2           Gpointer
		base_class_init_func GtkClassInitFunc
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
		keyval      Guint
		modifiers   GdkModifierType
		binding_set *GtkBindingSet
		bits        Guint
		// destroyed : 1
		// in_emission : 1
		// marks_unbound : 1
		set_next  *GtkBindingEntry
		hash_next *GtkBindingEntry
		signals   *GtkBindingSignal
	}

	PangoColor struct {
		red   Guint16
		green Guint16
		blue  Guint16
	}

	GtkVBox struct {
		box GtkBox
	}

	GtkRulerMetric struct {
		metric_name     *Gchar
		abbrev          *Gchar
		pixels_per_unit Gdouble
		ruler_scale     [10]Gdouble
		subdivide       [5]Gint
	}

	GtkBindingSignal struct {
		next        *GtkBindingSignal
		signal_name *Gchar
		n_args      Guint
		args        *GtkBindingArg
	}

	GtkBindingArg struct {
		arg_type GType
		//Union {
		// long_data  Glong;
		Double_data Gdouble
		// string_data  *Gchar;
		//}
	}

	GtkRadioActionEntry struct {
		name        *Gchar
		stock_id    *Gchar
		label       *Gchar
		accelerator *Gchar
		tooltip     *Gchar
		value       Gint
	}

	GCancellable struct {
		parent_instance GObject
		priv            *GCancellablePrivate
	}

	GdkAppLaunchContext struct {
		parent_instance GAppLaunchContext
		priv            *GdkAppLaunchContextPrivate
	}

	GdkKeymap struct {
		parent_instance GObject
		display         *GdkDisplay
	}

	GdkKeymapKey struct {
		keycode Guint
		group   Gint
		level   Gint
	}

	GdkPangoRenderer struct {
		parent_instance PangoRenderer
		priv            *GdkPangoRendererPrivate
	}

	GdkPixbufLoader struct {
		parent_instance GObject
		priv            Gpointer
	}

	GAppLaunchContext struct {
		parent_instance GObject
		priv            *GAppLaunchContextPrivate
	}

	PangoRenderer struct {
		parent_instance GObject
		underline       PangoUnderline
		strikethrough   Gboolean
		active_count    int
		matrix          *PangoMatrix
		priv            *PangoRendererPrivate
	}

	PangoMatrix struct {
		xx Double
		xy Double
		yx Double
		yy Double
		x0 Double
		y0 Double
	}

	PangoAttrShape struct {
		attr         PangoAttribute
		ink_rect     PangoRectangle
		logical_rect PangoRectangle
		data         Gpointer
		copy_func    PangoAttrDataCopyFunc
		destroy_func GDestroyNotify
	}

	PangoRectangle struct {
		x      int
		y      int
		width  int
		height int
	}

	GdkPointerHooks struct {
		get_pointer func(
			window *GdkWindow,
			x *Gint,
			y *Gint,
			mask *GdkModifierType) *GdkWindow
		window_at_pointer func(screen *GdkScreen,
			win_x *Gint,
			win_y *Gint) *GdkWindow
	}

	GdkRgbCmap struct {
		colors    [256]Guint32
		n_colors  Gint
		info_list *GSList
	}

	GdkSegment struct {
		x1 Gint
		y1 Gint
		x2 Gint
		y2 Gint
	}

	GdkSpan struct {
		x     Gint
		y     Gint
		width Gint
	}

	GdkTimeCoord struct {
		time Guint32
		axes [128]Gdouble
	}

	GdkTrapezoid struct {
		y1, x11, x21, y2, x12, x22 Double
	}

	GdkWindowAttr struct {
		title             *Gchar
		event_mask        Gint
		x, y              Gint
		width             Gint
		height            Gint
		wclass            GdkWindowClass
		visual            *GdkVisual
		colormap          *GdkColormap
		window_type       GdkWindowType
		cursor            *GdkCursor
		wmclass_name      *Gchar
		wmclass_class     *Gchar
		override_redirect Gboolean
		type_hint         GdkWindowTypeHint
	}

	GInputStream struct {
		parent_instance GObject
		priv            *GInputStreamPrivate
	}

	PangoAttribute struct {
		klass       *PangoAttrClass
		start_index Guint
		end_index   Guint
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
		item   *PangoItem
		glyphs *PangoGlyphString
	}

	PangoItem struct {
		offset    Gint
		length    Gint
		num_chars Gint
		analysis  PangoAnalysis
	}

	PangoAnalysis struct {
		shape_engine *PangoEngineShape
		lang_engine  *PangoEngineLang
		font         *PangoFont
		level        Guint8
		gravity      Guint8
		flags        Guint8
		script       Guint8
		language     *PangoLanguage
		extra_attrs  *GSList
	}

	PangoGlyphString struct {
		num_glyphs   Gint
		glyphs       *PangoGlyphInfo
		log_clusters *Gint
		space        Gint
	}

	PangoLayoutLine struct {
		layout      *PangoLayout
		start_index Gint
		length      Gint
		runs        *GSList
		bits        Guint
		// is_paragraph_start : 1
		// resolved_dir : 3
	}

	PangoGlyphInfo struct {
		glyph    PangoGlyph
		geometry PangoGlyphGeometry
		attr     PangoGlyphVisAttr
	}

	PangoGlyphGeometry struct {
		width    PangoGlyphUnit
		x_offset PangoGlyphUnit
		y_offset PangoGlyphUnit
	}

	PangoGlyphVisAttr struct {
		is_cluster_start Guint // : 1
	}

	GdkPixdata struct {
		magic        Guint32
		length       Gint32
		pixdata_type GdkPixdataType
		rowstride    Guint32
		width        Guint32
		height       Guint32
		pixel_data   *Guint8
	}

	PangoGlyphItemIter struct {
		glyph_item  *PangoGlyphItem
		text        *Gchar
		start_glyph int
		start_index int
		start_char  int
		end_glyph   int
		end_index   int
		end_char    int
	}

	Cairo_user_data_key_t struct {
		unused int
	}

	Cairo_matrix_t struct {
		xx Double
		yx Double
		xy Double
		yy Double
		x0 Double
		y0 Double
	}

	Cairo_rectangle_list_t struct {
		status         Cairo_status_t
		rectangles     *Cairo_rectangle_t
		num_rectangles int
	}

	Cairo_rectangle_t struct {
		x, y, width, height Double
	}

	Cairo_glyph_t struct {
		index Unsigned_long
		x     Double
		y     Double
	}

	Cairo_text_cluster_t struct {
		num_bytes  int
		num_glyphs int
	}

	Cairo_text_extents_t struct {
		x_bearing Double
		y_bearing Double
		width     Double
		height    Double
		x_advance Double
		y_advance Double
	}

	Cairo_font_extents_t struct {
		ascent        Double
		descent       Double
		height        Double
		max_x_advance Double
		max_y_advance Double
	}

	Cairo_path_t struct {
		status   Cairo_status_t
		data     *Cairo_path_data_t
		num_data int
	}

	Cairo_path_data_t struct {
		// todo(t): Union
		// header struct {
		// 	Type   Cairo_path_data_type_t
		// 	length int
		// }
		point struct {
			x, y Double
		}
	}

	Cairo_rectangle_int_t struct {
		x, y          int
		width, height int
	}

	Cairo_script_interpreter_hooks_t struct {
		closure         *Void
		surface_create  Csi_surface_create_func_t
		surface_destroy Csi_destroy_func_t
		context_create  Csi_context_create_func_t
		context_destroy Csi_destroy_func_t
		show_page       Csi_show_page_func_t
		copy_page       Csi_copy_page_func_t
	}

	GMemVTable struct {
		malloc  func(n_bytes Gsize) Gpointer
		realloc func(mem Gpointer,
			n_bytes Gsize) Gpointer
		free   func(mem Gpointer)
		calloc func(n_blocks Gsize,
			n_block_bytes Gsize) Gpointer
		try_malloc  func(n_bytes Gsize) Gpointer
		try_realloc func(mem Gpointer,
			n_bytes Gsize) Gpointer
	}

	GTrashStack struct {
		next *GTrashStack
	}

	GailTextUtil struct {
		parent GObject
		buffer *GtkTextBuffer
	}
)
