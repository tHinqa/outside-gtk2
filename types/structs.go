package types

//TODO(t): add constants for [...]s

type Tm struct {
	Sec, Min, Hour, Mday, Mon, Year, Wday, Yday, Isdst int
}

type GArray struct {
	data *Gchar
	leng Guint
}

type GByteArray struct {
	data *Guint8
	leng Guint
}

type GCompletion struct {
	items        *GList
	fnc          GCompletionFunc
	prefix       *Gchar
	cache        *GList
	strncmp_func GCompletionStrncmpFunc
}

type GDate struct {
	bits1, bits2 Guint
	/*
	   julian_days : 32
	   julian : 1
	   dmy : 1
	   day : 6
	   month : 4
	   year : 16
	*/
}

type GDebugKey struct {
	key   *Gchar
	value Guint
}

type GError struct {
	domain  GQuark
	code    Gint
	message *Gchar
}

type GHashTableIter struct {
	dummy1 Gpointer
	dummy2 Gpointer
	dummy3 Gpointer
	dummy4 int
	dummy5 Gboolean
	dummy6 Gpointer
}

type GHook struct {
	data      Gpointer
	next      *GHook
	prev      *GHook
	ref_count Guint
	hook_id   Gulong
	flags     Guint
	fnc       Gpointer
	destroy   GDestroyNotify
}

type GHookList struct {
	seq_id Gulong
	bits   Guint
	/*
	   hook_size : 16
	   is_setup : 1
	*/
	hooks         *GHook
	dummy3        Gpointer
	finalize_hook GHookFinalizeFunc
	dummy         [2]Gpointer
}

type GIOChannel struct {
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
	/*
	   use_buffer : 1
	   do_encode : 1
	   close_on_unref : 1
	   is_readable : 1
	   is_writeable : 1
	   is_seekable : 1
	*/
	_, _ Gpointer
}

type GList struct {
	data Gpointer
	next *GList
	prev *GList
}

type GIOFuncs struct {
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

type GMarkupParser struct {
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

type GNode struct {
	data     Gpointer
	next     *GNode
	prev     *GNode
	parent   *GNode
	children *GNode
}

type GSourceCallbackFuncs struct {
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

type GOnce struct {
	status GOnceStatus
	retval Gpointer
}

type GOptionEntry struct {
	long_name       *Gchar
	short_name      Gchar
	flags           Gint
	arg             GOptionArg
	arg_data        Gpointer
	description     *Gchar
	arg_description *Gchar
}

type GPollFD struct {
	fd      Gint
	events  Gushort
	revents Gushort
}

type GPtrArray struct {
	pdata *Gpointer
	len   Guint
}

type GQueue struct {
	head   *GList
	tail   *GList
	length Guint
}

type GScanner struct {
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

type GScannerConfig struct {
	cset_skip_characters  *Gchar
	cset_identifier_first *Gchar
	cset_identifier_nth   *Gchar
	cpair_comment_single  *Gchar
	bits                  Guint
	/*
	   case_sensitive : 1
	   skip_comment_multi : 1
	   skip_comment_single : 1
	   scan_comment_multi : 1
	   scan_identifier : 1
	   scan_identifier_1char : 1
	   scan_identifier_NULL : 1
	   scan_symbols : 1
	   scan_binary : 1
	   scan_octal : 1
	   scan_float : 1
	   scan_hex : 1
	   scan_hex_dollar : 1
	   scan_string_sq : 1
	   scan_string_dq : 1
	   numbers_2_int : 1
	   int_2_float : 1
	   identifier_2_string : 1
	   char_2_token : 1
	   symbol_2_token : 1
	   scope_0_fallback : 1
	   store_int64 : 1
	*/
	padding_dummy Guint
}

type GSList struct {
	data Gpointer
	next *GSList
}

type GSource struct {
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

type GStaticPrivate struct {
	index Guint
}

type GStaticRecMutex struct {
	mutex GStaticMutex
	depth Guint
	owner GSystemThread
}

type GSystemThread struct {
	// union
	// data[4] Char
	dummy_double Double
	// dummy_pointer *Void
	// dummy_long Long
}

type GStaticRWLock struct {
	mutex         GStaticMutex
	read_cond     *GCond
	write_cond    *GCond
	read_counter  Guint
	have_writer   Gboolean
	want_to_read  Guint
	want_to_write Guint
}

type GString struct {
	str           *Gchar
	len           Gsize
	allocated_len Gsize
}

type GTestLogBuffer struct {
	data *GString
	msgs *GSList
}

type GTestLogMsg struct {
	log_type  GTestLogType
	n_strings Guint
	strings   **Gchar
	n_nums    Guint
	nums      *Long_double
}

type GThread struct {
	Func     GThreadFunc
	data     Gpointer
	joinable Gboolean
	priority GThreadPriority
}

type GThreadFunctions struct {
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

type GThreadPool struct {
	Func      GFunc
	user_data Gpointer
	exclusive Gboolean
}

type GTimeVal struct {
	tv_sec  Glong
	tv_usec Glong
}

type GTokenValue struct {
	/*union
	  v_symbol  Gpointer;
	  v_identifier  *Gchar;
	  v_binary  Gulong;
	  v_octal  Gulong;
	  v_int  Gulong;
	*/
	v_int64 Guint64
	/*
	   v_float  Gdouble;
	   v_hex  Gulong;
	   v_string  *Gchar;
	   v_comment  *Gchar;
	   v_char  Guchar;
	   v_error  Guint;
	*/
}

type GTuples struct {
	len Guint
}

type GVariantBuilder struct {
	x [16]Gsize
}

type GVariantIter struct {
	x [16]Gsize
}

type GtkWidget struct {
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

type GdkDrawable struct {
	parent_instance GObject
}

type GtkRequisition struct {
	width  Gint
	height Gint
}

type GTypeClass struct {
	g_type GType
}

type GTypeInstance struct {
	g_class *GTypeClass
}

type GObject struct {
	g_type_instance GTypeInstance
	ref_count       Guint
	qdata           *GData
}

type GtkObject struct {
	parent_instance GInitiallyUnowned
	flags           Guint32
}

type GdkColor struct {
	pixel Guint32
	red   Guint16
	green Guint16
	blue  Guint16
}

type GtkStyle struct {
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
	fg_gc             *[5]GdkGC // CHECK
	bg_gc             *[5]GdkGC // CHECK
	light_gc          *[5]GdkGC // CHECK
	dark_gc           *[5]GdkGC // CHECK
	mid_gc            *[5]GdkGC // CHECK
	text_gc           *[5]GdkGC // CHECK
	base_gc           *[5]GdkGC // CHECK
	text_aa_gc        *[5]GdkGC // CHECK
	black_gc          *GdkGC
	white_gc          *GdkGC
	bg_pixmap         *[5]GdkPixmap // CHECK
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

type GtkRcStyle struct {
	parent_instance  GObject
	name             *Gchar
	bg_pixmap_name   *[5]Gchar // CHECK
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

type GdkFont struct {
	Type    GdkFontType
	ascent  Gint
	descent Gint
}

type GdkGC struct {
	parent_instance GObject
	clip_x_origin   Gint
	clip_y_origin   Gint
	ts_x_origin     Gint
	ts_y_origin     Gint
	colormap        *GdkColormap
}

type GdkColormap struct {
	parent_instance GObject
	size            Gint
	colors          *GdkColor
	visual          *GdkVisual
	windowing_data  Gpointer
}

type GdkVisual struct {
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

type GdkRectangle struct {
	x      Gint
	y      Gint
	width  Gint
	height Gint
}

type GtkCurve struct {
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
	ctlpoint      [2]*Gfloat // Gfloat (*ctlpoint)[2]; ???
}

type GtkDrawingArea struct {
	widget    GtkWidget
	draw_data Gpointer
}

type GdkPoint struct {
	x Gint
	y Gint
}

// typedef struct _GtkFileSelectionClass GtkFileSelectionClass;
type GtkFileSelection struct {
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

type GtkDialog struct {
	window      GtkWindow
	vbox        *GtkWidget
	action_area *GtkWidget
	separator   *GtkWidget
}

type GtkBin struct {
	container GtkContainer
	child     *GtkWidget
}

type GtkContainer struct {
	widget      GtkWidget
	focus_child *GtkWidget
	bits        Guint
	/*
	  border_width : 16
	  need_resize : 1
	  resize_mode : 2
	  reallocate_redraws : 1
	  has_focus_chain : 1
	*/
}

type GtkWindow struct {
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
	/*
	   allow_shrink : 1
	   allow_grow : 1
	   configure_notify_received : 1
	   need_default_position : 1
	   need_default_size : 1
	   position : 3
	   type : 4
	   has_user_ref_count : 1
	   has_focus : 1
	   modal : 1
	   destroy_with_parent : 1
	   has_frame : 1
	   iconify_initially : 1
	   stick_initially : 1
	   maximize_initially : 1
	   decorated : 1
	   type_hint : 3
	   gravity : 5
	   is_active : 1
	   has_toplevel_focus : 1
	*/
	frame_left           Guint
	frame_top            Guint
	frame_right          Guint
	frame_bottom         Guint
	keys_changed_handler Guint
	mnemonic_modifier    GdkModifierType
	screen               *GdkScreen
}

type GtkWindowGroup struct {
	parent_instance GObject
	grabs           *GSList
}

type GdkScreen struct {
	parent_instance GObject
	closed          Guint      // : 1
	normal_gcs      *[32]GdkGC // CHECK
	exposure_gcs    *[32]GdkGC // CHECK
	subwindow_gcs   *[32]GdkGC // CHECK
	font_options    *Cairo_font_options_t
	resolution      Double
}

type GtkItemFactory struct {
	object           GtkObject
	path             *Gchar
	accel_group      *GtkAccelGroup
	widget           *GtkWidget
	items            *GSList
	translate_func   GtkTranslateFunc
	translate_data   Gpointer
	translate_notify GDestroyNotify
}

type GtkAccelGroup struct {
	parent         GObject
	lock_count     Guint
	modifier_mask  GdkModifierType
	acceleratables *GSList
	n_accels       Guint
	priv_accels    *GtkAccelGroupEntry
}

type GtkAccelGroupEntry struct {
	key              GtkAccelKey
	closure          *GClosure
	accel_path_quark GQuark
}

type GtkAccelKey struct {
	accel_key   Guint
	accel_mods  GdkModifierType
	accel_flags Guint //: 16
}

type GClosure struct {
	bits Guint
	/*
	  ref_count : 15
	  meta_marshal : 1
	  n_guards : 1
	  n_fnotifiers : 2
	  n_inotifiers : 8
	  in_inotify : 1
	  floating : 1
	  derivative_flag : 1
	  in_marshal : 1
	  is_invalid : 1
	*/
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

type GValue struct {
	g_type GType
	data   [2]uint64
	/* UNION
	v_int     Gint
	v_uint    Guint
	v_long    glong
	v_ulong   Gulong
	v_int64   Gint64
	v_uint64  Guint64
	v_float   Gfloat
	v_double  Gdouble
	v_pointer Gpointer
	*/
}
type GClosureNotifyData struct {
	data   Gpointer
	notify GClosureNotify
}

type GtkItemFactoryEntry struct {
	path            *Gchar
	accelerator     *Gchar
	callback        GtkItemFactoryCallback
	callback_action Guint
	item_type       *Gchar
	extra_data      Gconstpointer
}

type GtkMenuEntry struct {
	path          *Gchar
	accelerator   *Gchar
	callback      GtkMenuCallback
	callback_data Gpointer
	widget        *GtkWidget
}

type GtkList struct {
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
	/*
		selection_mode : 2
		drag_selection : 1
		add_mode : 1
	*/
}

type GtkListItem struct {
	item GtkItem
}

type GtkItem struct {
	bin GtkBin
}

type GtkOldEditable struct {
	widget              GtkWidget
	current_pos         Guint
	selection_start_pos Guint
	selection_end_pos   Guint
	bits                Guint
	/*
		has_selection : 1
		editable : 1
		visible : 1
	*/
	clipboard_text *Gchar
}

type GtkOptionMenu struct {
	button    GtkButton
	menu      *GtkWidget
	menu_item *GtkWidget
	width     Guint16
	height    Guint16
}

type GtkButton struct {
	bin              GtkBin
	event_window     *GdkWindow
	label_text       *Gchar
	activate_timeout Guint
	bits             Guint
	/*
		constructed : 1
		in_button : 1
		button_down : 1
		relief : 2
		use_underline : 1
		use_stock : 1
		depressed : 1
		depress_on_activate : 1
		focus_on_click : 1
	*/
}

type GtkPreview struct {
	widget        GtkWidget
	buffer        *Guchar
	buffer_width  Guint16
	buffer_height Guint16
	bpp           Guint16
	rowstride     Guint16
	dither        GdkRgbDither
	bits          Guint
	/*
	   type : 1
	   expand : 1
	*/
}

type GtkPreviewInfo struct {
	lookup *Guchar
	gamma  Gdouble
}

type GtkTipsQuery struct {
	label GtkLabel
	bits  Guint
	/*
	   emit_always : 1
	   in_query : 1
	*/
	label_inactive *Gchar
	label_no_tip   *Gchar
	caller         *GtkWidget
	last_crossed   *GtkWidget
	query_cursor   *GdkCursor
}

type GdkCursor struct {
	Type      GdkCursorType
	ref_count Guint
}

type GtkLabel struct {
	misc  GtkMisc
	label *Gchar
	bits  Guint
	/*
		jtype : 2
		wrap : 1
		use_underline : 1
		use_markup : 1
		ellipsize : 3
		single_line_mode : 1
		have_transform : 1
		in_click : 1
		wrap_mode : 3
		pattern_set : 1
		track_links : 1
	*/
	mnemonic_keyval Guint
	text            *Gchar
	attrs           *PangoAttrList
	effective_attrs *PangoAttrList
	layout          *PangoLayout
	mnemonic_widget *GtkWidget
	mnemonic_window *GtkWindow
	select_info     *GtkLabelSelectionInfo
}

type GtkMisc struct {
	widget GtkWidget
	xalign Gfloat
	yalign Gfloat
	xpad   Guint16
	ypad   Guint16
}

type GtkCList struct {
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

type GtkCListColumn struct {
	title         *Gchar
	area          GdkRectangle
	button        *GtkWidget
	window        *GdkWindow
	width         Gint
	min_width     Gint
	max_width     Gint
	justification GtkJustification
	bits          Guint
	/*
	   visible : 1
	   width_set : 1
	   resizeable : 1
	   auto_resize : 1
	   button_passive : 1
	*/
}

type GtkCListCellInfo struct {
	row    Gint
	column Gint
}

type GtkAdjustment struct {
	parent_instance GtkObject
	lower           Gdouble
	upper           Gdouble
	value           Gdouble
	step_increment  Gdouble
	page_increment  Gdouble
	page_size       Gdouble
}

type GtkCListRow struct {
	cell       *GtkCell
	state      GtkStateType
	foreground GdkColor
	background GdkColor
	style      *GtkStyle
	data       Gpointer
	destroy    GDestroyNotify
	bits       Guint
	/*
	   fg_set : 1
	   bg_set : 1
	   selectable : 1
	*/
}

type GtkCell struct {
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

type GtkCombo struct {
	hbox            GtkHBox
	entry           *GtkWidget
	button          *GtkWidget
	popup           *GtkWidget
	popwin          *GtkWidget
	list            *GtkWidget
	entry_change_id Guint
	list_change_id  Guint
	bits            Guint
	/*
	   value_in_list:1
	   ok_if_empty:1
	   case_sensitive:1
	   use_arrows:1
	   use_arrows_always:1
	*/
	current_button Guint16
	activate_id    Guint
}

type GtkHBox struct {
	box GtkBox
}

type GtkBox struct {
	container   GtkContainer
	children    *GList
	spacing     Gint16
	homogeneous Guint // : 1
}

type GtkCTree struct {
	clist        GtkCList
	lines_gc     *GdkGC
	tree_indent  Gint
	tree_spacing Gint
	tree_column  Gint
	bits         Guint
	/*
	   line_style : 2
	   expander_style : 2
	   show_stub : 1
	*/
	drag_compare GtkCTreeCompareDragFunc
}

type GtkCTreeNode struct {
	list GList
}

type GtkCTreeRow struct {
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
	/*
	   is_leaf : 1
	   expanded : 1
	*/
}

type GtkTextBuffer struct {
	parent_instance            GObject
	tag_table                  *GtkTextTagTable
	btree                      *GtkTextBTree
	clipboard_contents_buffers *GSList
	selection_clipboards       *GSList
	log_attr_cache             *GtkTextLogAttrCache
	user_action_count          Guint
	bits                       Guint
	/*
	   modified : 1
	   has_selection : 1
	*/
}

type GtkTextTagTable struct {
	parent_instance GObject
	hash            *GHashTable
	anonymous       *GSList
	anon_count      Gint
	buffers         *GSList
}

type GtkTextView struct {
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
	/*
	   editable : 1
	   overwrite_mode : 1
	   cursor_visible : 1
	   need_im_reset : 1
	   accepts_tab : 1
	   width_changed : 1
	   onscreen_validated : 1
	   mouse_cursor_obscured : 1
	*/
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

type GtkTextMark struct {
	parent_instance GObject
	segment         Gpointer
}

type GtkIMContext struct {
	parent_instance GObject
}

type GtkTextIter struct {
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

type GdkEventKey struct {
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

type GtkTextChildAnchor struct {
	parent_instance GObject
	segment         Gpointer
}

type GtkTextAttributes struct {
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
	/*
	   invisible : 1
	   bg_full_height : 1
	   editable : 1
	   realized : 1
	   pad1 : 1
	   pad2 : 1
	   pad3 : 1
	   pad4 : 1
	*/
}

type GtkTextAppearance struct {
	bg_color   GdkColor
	fg_color   GdkColor
	bg_stipple *GdkBitmap
	fg_stipple *GdkBitmap
	rise       Gint
	padding1   Gpointer
	bits       Guint
	/*
	   underline : 4
	   strikethrough : 1
	   draw_bg : 1
	   inside_selection : 1
	   is_text : 1
	   pad1 : 1
	   pad2 : 1
	   pad3 : 1
	   pad4 : 1
	*/
}

type GtkPixmap struct {
	misc               GtkMisc
	pixmap             *GdkPixmap
	mask               *GdkBitmap
	pixmap_insensitive *GdkPixmap
	build_insensitive  Guint // : 1
}

type GtkToolbar struct {
	container    GtkContainer
	num_children Gint
	children     *GList
	orientation  GtkOrientation
	style        GtkToolbarStyle
	icon_size    GtkIconSize
	tooltips     *GtkTooltips
	button_maxw  Gint
	button_maxh  Gint
	_            Guint
	_            Guint
	bits         Guint
	/*
	   style_set : 1
	   icon_size_set : 1
	*/
}

type GtkTooltips struct {
	parent_instance  GtkObject
	tip_window       *GtkWidget
	tip_label        *GtkWidget
	active_tips_data *GtkTooltipsData
	tips_data_list   *GList
	bits             Guint // UTTER MESS 32/64
	/*
	   Guint delay : 30
	   Guint enabled : 1
	   Guint have_grab : 1
	*/
	bits2 Guint
	/*
	   Guint use_sticky_delay : 1
	*/
	timer_tag    Gint
	last_popdown GTimeVal
}

type GtkTooltipsData struct {
	tooltips    *GtkTooltips
	widget      *GtkWidget
	tip_text    *Gchar
	tip_private *Gchar
}

type GtkToolItem struct {
	parent GtkBin
	priv   *GtkToolItemPrivate
}

type GtkToolItemGroup struct {
	parent_instance GtkContainer
	priv            *GtkToolItemGroupPrivate
}

type GtkToolPalette struct {
	parent_instance GtkContainer
	priv            *GtkToolPalettePrivate
}

type GtkSelectionData struct {
	selection GdkAtom
	target    GdkAtom
	Type      GdkAtom
	format    Gint
	data      *Guchar
	length    Gint
	display   *GdkDisplay
}

type GdkDisplay struct {
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
	/*
	   closed : 1
	   ignore_core_events : 1
	*/
	Double_click_distance Guint
	button_x              [2]Gint
	button_y              [2]Gint
	pointer_grabs         *GList
	keyboard_grab         GdkKeyboardGrabInfo
	pointer_info          GdkPointerWindowInfo
	last_event_time       Guint32
}

type GdkKeyboardGrabInfo struct {
	window        *GdkWindow
	native_window *GdkWindow
	serial        Gulong
	owner_events  Gboolean
	time          Guint32
}

type GdkPointerWindowInfo struct {
	toplevel_under_pointer *GdkWindow
	window_under_pointer   *GdkWindow
	toplevel_x, toplevel_y Gdouble
	state                  Guint32
	button                 Guint32
	motion_hint_serial     Gulong
}

type GdkDevice struct {
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

type GdkDeviceKey struct {
	keyval    Guint
	modifiers GdkModifierType
}

type GdkDisplayPointerHooks struct {
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

type GdkDeviceAxis struct {
	use GdkAxisUse
	min Gdouble
	max Gdouble
}

type GtkTargetEntry struct {
	target *Gchar
	flags  Guint
	info   Guint
}

type GtkSizeGroup struct {
	parent_instance GObject
	widgets         *GSList
	mode            Guint8
	bits            Guint
	/*
	   have_width : 1
	   have_height : 1
	   ignore_hidden : 1
	*/
	requisition GtkRequisition
}

type GtkSpinButton struct {
	entry         GtkEntry
	adjustment    *GtkAdjustment
	panel         *GdkWindow
	timer         Guint32
	climb_rate    Gdouble
	timer_step    Gdouble
	update_policy GtkSpinButtonUpdatePolicy
	bits          Guint
	/*
	   in_child : 2
	   click_child : 2
	   button : 2
	   need_timer : 1
	   timer_calls : 3
	   digits : 10
	   numeric : 1
	   wrap : 1
	   snap_to_ticks : 1
	*/
}

type GtkEntry struct {
	widget GtkWidget
	text   *Gchar
	bits   Guint
	/*
	   editable : 1
	   visible : 1
	   overwrite_mode : 1
	   in_drag : 1
	*/
	text_length     Guint16
	text_max_length Guint16
	text_area       *GdkWindow
	im_context      *GtkIMContext
	popup_menu      *GtkWidget
	current_pos     Gint
	selection_bound Gint
	cached_layout   *PangoLayout
	bits2           Guint
	/*
	   cache_includes_preedit : 1
	   need_im_reset : 1
	   has_frame : 1
	   activates_default : 1
	   cursor_visible : 1
	   in_click : 1
	   is_cell_renderer : 1
	   editing_canceled : 1
	   mouse_cursor_obscured : 1
	   select_words : 1
	   select_lines : 1
	   resolved_dir : 4
	   truncate_multiline : 1
	*/
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

type GtkTreeModelSort struct {
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

type GtkTreeIter struct {
	stamp      Gint
	user_data  Gpointer
	user_data2 Gpointer
	user_data3 Gpointer
}

type GtkTreeSelection struct {
	parent    GObject
	tree_view *GtkTreeView
	Type      GtkSelectionMode
	user_func GtkTreeSelectionFunc
	user_data Gpointer
	destroy   GDestroyNotify
}

type GtkTreeView struct {
	parent GtkContainer
	priv   *GtkTreeViewPrivate
}

type GtkTreeStore struct {
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

type GtkUIManager struct {
	parent       GObject
	private_data *GtkUIManagerPrivate
}

type GtkActionGroup struct {
	parent       GObject
	private_data *GtkActionGroupPrivate
}

type GtkAction struct {
	object       GObject
	private_data *GtkActionPrivate
}

type GtkHSV struct {
	parent_instance GtkWidget
	priv            Gpointer
}

type GtkIconFactory struct {
	parent_instance GObject
	icons           *GHashTable
}

type GtkSettings struct {
	parent_instance GObject
	queued_settings *GData
	property_values *GtkSettingsPropertyValue
	rc_context      *GtkRcContext
	screen          *GdkScreen
}

type GtkIconTheme struct {
	parent_instance GObject
	priv            *GtkIconThemePrivate
}

type GdkEvent struct {
	Type GdkEventType
	/* Union
	   any  GdkEventAny;
	   expose  GdkEventExpose;
	   no_expose  GdkEventNoExpose;
	   visibility  GdkEventVisibility;
	   motion  GdkEventMotion;
	   button  GdkEventButton;
	   scroll  GdkEventScroll;
	   key  GdkEventKey;
	   crossing  GdkEventCrossing;
	   focus_change  GdkEventFocus;
	   configure  GdkEventConfigure;
	   property  GdkEventProperty;
	   selection  GdkEventSelection;
	   owner_change  GdkEventOwnerChange;
	   proximity  GdkEventProximity;
	   client  GdkEventClient;
	   dnd  GdkEventDND;
	   window_state  GdkEventWindowState;
	   setting  GdkEventSetting;
	   grab_broken  GdkEventGrabBroken;
	*/
}

type GtkIconView struct {
	parent GtkContainer
	priv   *GtkIconViewPrivate
}

type GtkCellRenderer struct {
	parent GtkObject
	xalign Gfloat
	yalign Gfloat
	width  Gint
	height Gint
	xpad   Guint16
	ypad   Guint16
	bits   Guint
	/*
	   mode : 2;
	   visible : 1
	   is_expander : 1
	   is_expanded : 1
	   cell_background_set : 1
	   sensitive : 1
	   editing : 1
	*/
}

type GtkImageMenuItem struct {
	menu_item GtkMenuItem
	image     *GtkWidget
}

type GtkMenuItem struct {
	item              GtkItem
	submenu           *GtkWidget
	event_window      *GdkWindow
	toggle_size       Guint16
	accelerator_width Guint16
	accel_path        *Gchar
	bits              Guint
	/*
	   show_submenu_indicator : 1
	   submenu_placement : 1
	   submenu_direction : 1
	   right_justify: 1
	   timer_from_keypress : 1
	   from_menubar : 1
	*/
	timer Guint
}

type GtkIMContextSimple struct {
	object              GtkIMContext
	tables              *GSList
	compose_buffer      [7 + 1]Guint
	tentative_match     Gunichar
	tentative_match_len Gint
	bits                Guint
	/*
	   in_hex_sequence : 1
	   modifiers_dropped : 1
	*/
}

type GtkIMMulticontext struct {
	object     GtkIMContext
	slave      *GtkIMContext
	priv       *GtkIMMulticontextPrivate
	context_id *Gchar
}

type GtkMenuShell struct {
	container         GtkContainer
	children          *GList
	active_menu_item  *GtkWidget
	parent_menu_shell *GtkWidget
	button            Guint
	activate_time     Guint32
	bits              Guint
	/*
	   active : 1
	   have_grab : 1
	   have_xgrab : 1
	   ignore_leave : 1
	   menu_flag : 1
	   ignore_enter : 1
	   keyboard_mode : 1
	*/
}

type GtkInfoBar struct {
	parent GtkHBox
	priv   *GtkInfoBarPrivate
}

type GtkInvisible struct {
	widget             GtkWidget
	has_user_ref_count Gboolean
	screen             *GdkScreen
}

type GtkLayout struct {
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

type GtkLinkButton struct {
	parent_instance GtkButton
	priv            *GtkLinkButtonPrivate
}

type GtkArg struct {
	Type GType
	name *Gchar

	/* union
	   Gchar char_data;
	   uchar_data  Guchar;
	   bool_data  Gboolean;
	   int_data  Gint;
	   uint_data  Guint;
	   long_data  glong;
	   ulong_data  Gulong;
	   float_data  Gfloat;
	   Double_data  Gdouble;
	   string_data  *Gchar;
	   object_data  *GtkObject;
	   pointer_data  Gpointer;
	   // signal_data struct
	   		f  GCallback;
	   		d  Gpointer;
	   //
	*/
}

type GSignalInvocationHint struct {
	signal_id Guint
	detail    GQuark
	run_type  GSignalFlags
}

type GtkMenuBar struct {
	menu_shell GtkMenuShell
}

type GtkToolButton struct {
	parent GtkToolItem
	priv   *GtkToolButtonPrivate
}

type GtkMenuToolButton struct {
	parent GtkToolButton
	priv   *GtkMenuToolButtonPrivate
}

type GtkMessageDialog struct {
	parent_instance GtkDialog
	image           *GtkWidget
	label           *GtkWidget
}

type GMountOperation struct {
	parent_instance GObject
	priv            *GMountOperationPrivate
}

type GtkMountOperation struct {
	parent_instance GMountOperation
	priv            *GtkMountOperationPrivate
}

type GtkNotebook struct {
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
	/*
	   show_tabs : 1
	   homogeneous : 1
	   show_border : 1
	   tab_pos : 2
	   scrollable : 1
	   in_child : 3
	   click_child : 3
	   button : 2
	   need_timer : 1
	   child_has_focus : 1
	   have_visible_child : 1
	   focus_out : 1
	   has_before_previous : 1
	   has_before_next : 1
	   has_after_previous : 1
	   has_after_next : 1
	*/
}

type GtkOffscreenWindow struct {
	parent_object GtkWindow
}

type GtkSocket struct {
	container      GtkContainer
	request_width  Guint16
	request_height Guint16
	current_width  Guint16
	current_height Guint16
	plug_window    *GdkWindow
	plug_widget    *GtkWidget
	xembed_version Gshort
	bits           Guint
	/*
	   same_app : 1
	   focus_in : 1
	   have_size : 1
	   need_map : 1
	   is_mapped : 1
	   active : 1
	*/
	accel_group *GtkAccelGroup
	toplevel    *GtkWidget
}

type GtkPlug struct {
	window          GtkWindow
	socket_window   *GdkWindow
	modality_window *GtkWidget
	modality_group  *GtkWindowGroup
	grabbed_keys    *GHashTable
	same_app        Guint // : 1
}

type GtkPageRange struct {
	start Gint
	end   Gint
}

type GtkPrintOperation struct {
	parent_instance GObject
	priv            *GtkPrintOperationPrivate
}

type GtkProgress struct {
	widget           GtkWidget
	adjustment       *GtkAdjustment
	offscreen_pixmap *GdkPixmap
	format           *Gchar
	x_align          Gfloat
	y_align          Gfloat
	bits             Guint
	/*
	   show_text : 1;
	   activity_mode : 1;
	   use_text_format : 1
	*/
}

type GtkProgressBar struct {
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
	/*
	   activity_dir : 1
	   ellipsize : 3
	   dirty : 1
	*/
}

type GtkToggleAction struct {
	parent       GtkAction
	private_data *GtkToggleActionPrivate
}

type GtkRadioAction struct {
	parent       GtkToggleAction
	private_data *GtkRadioActionPrivate
}

type GtkRadioButton struct {
	check_button GtkCheckButton
	group        *GSList
}

type GtkRadioMenuItem struct {
	check_menu_item GtkCheckMenuItem
	group           *GSList
}

type GtkCheckButton struct {
	toggle_button GtkToggleButton
}

type GtkCheckMenuItem struct {
	menu_item GtkMenuItem
	bits      Guint
	/*
	   active : 1
	   always_show_toggle : 1
	   inconsistent : 1
	   draw_as_radio : 1
	*/
}

type GtkToggleButton struct {
	button GtkButton
	bits   Guint
	/*
	   active : 1
	   draw_indicator : 1
	   inconsistent : 1
	*/
}

type GtkToggleToolButton struct {
	parent GtkToolButton
	priv   *GtkToggleToolButtonPrivate
}

type GtkRadioToolButton struct {
	parent GtkToggleToolButton
}

type GtkRecentManager struct {
	parent_instance GObject
	priv            *GtkRecentManagerPrivate
}

type GtkRecentData struct {
	display_name *Gchar
	description  *Gchar
	mime_type    *Gchar
	app_name     *Gchar
	app_exec     *Gchar
	groups       **Gchar
	is_private   Gboolean
}

type GtkRecentAction struct {
	parent_instance GtkAction
	priv            *GtkRecentActionPrivate
}

type GtkScaleButton struct {
	parent       GtkButton
	plus_button  *GtkWidget
	minus_button *GtkWidget
	priv         *GtkScaleButtonPrivate
}

type GtkRecentFilterInfo struct {
	contains     GtkRecentFilterFlags
	uri          *Gchar
	display_name *Gchar
	mime_type    *Gchar
	applications **Gchar
	groups       **Gchar
	age          Gint
}

type GtkViewport struct {
	bin         GtkBin
	shadow_type GtkShadowType
	view_window *GdkWindow
	bin_window  *GdkWindow
	hadjustment *GtkAdjustment
	vadjustment *GtkAdjustment
}

type GtkScrolledWindow struct {
	container  GtkBin
	hscrollbar *GtkWidget
	vscrollbar *GtkWidget
	bits       Guint
	/*
	   hscrollbar_policy : 2
	   vscrollbar_policy : 2
	   hscrollbar_visible : 1
	   vscrollbar_visible : 1
	   window_placement : 2
	   focus_out : 1
	*/
	shadow_type Guint16
}

type GtkRecentChooserMenu struct {
	parent_instance GtkMenu
	priv            *GtkRecentChooserMenuPrivate
}

type GtkMenu struct {
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
	/*
	   needs_destruction_ref_count : 1
	   torn_off : 1
	   tearoff_active : 1
	   scroll_fast : 1
	   upper_arrow_visible : 1
	   lower_arrow_visible : 1
	   upper_arrow_prelight : 1
	   lower_arrow_prelight : 1
	*/
}

type GtkSeparatorToolItem struct {
	parent GtkToolItem
	priv   *GtkSeparatorToolItemPrivate
}

type GtkStatusbar struct {
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

type GtkSpinner struct {
	parent GtkDrawingArea
	priv   *GtkSpinnerPrivate
}

type GtkStatusIcon struct {
	parent_instance GObject
	priv            *GtkStatusIconPrivate
}

type GtkStockItem struct {
	stock_id           *Gchar
	label              *Gchar
	modifier           GdkModifierType
	keyval             Guint
	translation_domain *Gchar
}

type GtkTable struct {
	container      GtkContainer
	children       *GList
	rows           *GtkTableRowCol
	cols           *GtkTableRowCol
	nrows          Guint16
	ncols          Guint16
	column_spacing Guint16
	row_spacing    Guint16
	bits           Guint
	/*
	   homogeneous : 1
	*/
}

type GtkTableRowCol struct {
	requisition Guint16
	allocation  Guint16
	spacing     Guint16
	bits        Guint
	/*
	   need_expand : 1
	   need_shrink : 1
	   expand : 1
	   shrink : 1
	   empty : 1
	*/
}

type GtkTextTag struct {
	parent_instance GObject
	table           *GtkTextTagTable
	name            *Char
	priority        int
	values          *GtkTextAttributes
	bits            Guint
	/*
	   bg_color_set : 1
	   bg_stipple_set : 1
	   fg_color_set : 1
	   scale_set : 1
	   fg_stipple_set : 1
	   justification_set : 1
	   left_margin_set : 1
	   indent_set : 1
	   rise_set : 1
	   strikethrough_set : 1
	   right_margin_set : 1
	   pixels_above_lines_set : 1
	   pixels_below_lines_set : 1
	   pixels_inside_wrap_set : 1
	   tabs_set : 1
	   underline_set : 1
	   wrap_mode_set : 1
	   bg_full_height_set : 1
	   invisible_set : 1
	   editable_set : 1
	   language_set : 1
	   pg_bg_color_set : 1
	   accumulative_margin : 1
	   pad1 : 1
	*/
}

type GtkTargetList struct {
	list      *GList
	ref_count Guint
}

type PangoLogAttr struct {
	bits Guint
	/*
	   is_line_break : 1
	   is_mandatory_break : 1
	   is_char_break : 1
	   is_white : 1
	   is_cursor_position : 1
	   is_word_start : 1
	   is_word_end : 1
	   is_sentence_boundary : 1
	   is_sentence_start : 1
	   is_sentence_end : 1
	   backspace_deletes_character : 1
	   is_expandable_space : 1
	   is_word_boundary : 1
	*/
}
