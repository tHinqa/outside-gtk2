package gtksourceview

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type GtkSourceCompletionActivation Enum

const (
	GTK_SOURCE_COMPLETION_ACTIVATION_INTERACTIVE GtkSourceCompletionActivation = 1 << iota
	GTK_SOURCE_COMPLETION_ACTIVATION_USER_REQUESTED
	GTK_SOURCE_COMPLETION_ACTIVATION_NONE GtkSourceCompletionActivation = 0
)

type GtkSourceSmartHomeEndType Enum

const (
	GTK_SOURCE_SMART_HOME_END_DISABLED GtkSourceSmartHomeEndType = iota
	GTK_SOURCE_SMART_HOME_END_BEFORE
	GTK_SOURCE_SMART_HOME_END_AFTER
	GTK_SOURCE_SMART_HOME_END_ALWAYS
)

type GtkSourceDrawSpacesFlags Enum

const (
	GTK_SOURCE_DRAW_SPACES_SPACE GtkSourceDrawSpacesFlags = 1 << iota
	GTK_SOURCE_DRAW_SPACES_TAB
	GTK_SOURCE_DRAW_SPACES_NEWLINE
	GTK_SOURCE_DRAW_SPACES_NBSP
	GTK_SOURCE_DRAW_SPACES_LEADING
	GTK_SOURCE_DRAW_SPACES_TEXT
	GTK_SOURCE_DRAW_SPACES_TRAILING
	GTK_SOURCE_DRAW_SPACES_ALL = GTK_SOURCE_DRAW_SPACES_SPACE |
		GTK_SOURCE_DRAW_SPACES_TAB |
		GTK_SOURCE_DRAW_SPACES_NEWLINE |
		GTK_SOURCE_DRAW_SPACES_NBSP |
		GTK_SOURCE_DRAW_SPACES_LEADING |
		GTK_SOURCE_DRAW_SPACES_TEXT |
		GTK_SOURCE_DRAW_SPACES_TRAILING
)

type GtkSourceSearchFlags Enum

const (
	GTK_SOURCE_SEARCH_VISIBLE_ONLY GtkSourceSearchFlags = 1 << iota
	GTK_SOURCE_SEARCH_TEXT_ONLY
	GTK_SOURCE_SEARCH_CASE_INSENSITIVE
)

type (
	GtkSourceBufferPrivate             struct{}
	GtkSourceCompletionContextPrivate  struct{}
	GtkSourceCompletionInfoPrivate     struct{}
	GtkSourceCompletionItemPrivate     struct{}
	GtkSourceCompletionPrivate         struct{}
	GtkSourceCompletionProposal        struct{}
	GtkSourceCompletionProvider        struct{}
	GtkSourceCompletionWordsPrivate    struct{}
	GtkSourceGutterPrivate             struct{}
	GtkSourceLanguageManagerPrivate    struct{}
	GtkSourceLanguagePrivate           struct{}
	GtkSourceMarkPrivate               struct{}
	GtkSourcePrintCompositorPrivate    struct{}
	GtkSourceStyle                     struct{}
	GtkSourceStyleSchemePrivate        struct{}
	GtkSourceUndoManager               struct{}
	GtkSourceViewPrivate               struct{}
	GtkSourceStyleSchemeManagerPrivate struct{}

	GtkSourceGutterDataFunc func(
		gutter *GtkSourceGutter,
		cell *GtkCellRenderer,
		line_number Gint,
		current_line Gboolean,
		data Gpointer)

	GtkSourceGutterSizeFunc func(
		gutter *GtkSourceGutter,
		cell *GtkCellRenderer,
		data Gpointer)

	GtkSourceViewMarkTooltipFunc func(
		mark *GtkSourceMark,
		user_data Gpointer) *Gchar

	GtkSourceLanguage struct {
		parent_instance GObject
		priv            *GtkSourceLanguagePrivate
	}

	GtkSourceMark struct {
		parent_instance GtkTextMark
		priv            *GtkSourceMarkPrivate
	}

	GtkSourceStyleScheme struct {
		base GObject
		priv *GtkSourceStyleSchemePrivate
	}

	GtkSourceBuffer struct {
		parent_instance GtkTextBuffer
		priv            *GtkSourceBufferPrivate
	}

	GtkSourceCompletionInfo struct {
		parent GtkWindow
		priv   *GtkSourceCompletionInfoPrivate
	}

	GtkSourceCompletionContext struct {
		parent GInitiallyUnowned
		priv   *GtkSourceCompletionContextPrivate
	}

	GtkSourceCompletion struct {
		parent GtkObject
		priv   *GtkSourceCompletionPrivate
	}

	GtkSourceGutter struct {
		parent GObject
		priv   *GtkSourceGutterPrivate
	}

	GtkSourceView struct {
		parent GtkTextView
		priv   *GtkSourceViewPrivate
	}

	GtkSourceCompletionItem struct {
		parent GObject
		priv   *GtkSourceCompletionItemPrivate
	}

	GtkSourceCompletionWords struct {
		parent GObject
		priv   *GtkSourceCompletionWordsPrivate
	}

	GtkSourceLanguageManager struct {
		parent_instance GObject
		priv            *GtkSourceLanguageManagerPrivate
	}

	GtkSourcePrintCompositor struct {
		parent_instance GObject
		priv            *GtkSourcePrintCompositorPrivate
	}

	GtkSourceStyleSchemeManager struct {
		parent GObject
		priv   *GtkSourceStyleSchemeManagerPrivate
	}
)

var (
	Gtk_source_language_get_type func() GType

	Gtk_source_language_get_id func(
		language *GtkSourceLanguage) *Gchar

	Gtk_source_language_get_name func(
		language *GtkSourceLanguage) *Gchar

	Gtk_source_language_get_section func(
		language *GtkSourceLanguage) *Gchar

	Gtk_source_language_get_hidden func(
		language *GtkSourceLanguage) Gboolean

	Gtk_source_language_get_metadata func(
		language *GtkSourceLanguage,
		name *Gchar) *Gchar

	Gtk_source_language_get_mime_types func(
		language *GtkSourceLanguage) **Gchar

	Gtk_source_language_get_globs func(
		language *GtkSourceLanguage) **Gchar

	Gtk_source_language_get_style_ids func(
		language *GtkSourceLanguage) **Gchar

	Gtk_source_language_get_style_name func(
		language *GtkSourceLanguage,
		style_id *Gchar) *Gchar

	Gtk_source_mark_get_type func() GType

	Gtk_source_mark_new func(
		name *Gchar,
		category *Gchar) *GtkSourceMark

	Gtk_source_mark_get_category func(
		mark *GtkSourceMark) *Gchar

	Gtk_source_mark_next func(
		mark *GtkSourceMark,
		category *Gchar) *GtkSourceMark

	Gtk_source_mark_prev func(
		mark *GtkSourceMark,
		category *Gchar) *GtkSourceMark

	Gtk_source_style_get_type func() GType

	Gtk_source_style_copy func(
		style *GtkSourceStyle) *GtkSourceStyle

	Gtk_source_style_scheme_get_type func() GType

	_gtk_source_style_scheme_new func(
		id *Gchar,
		name *Gchar) *GtkSourceStyleScheme

	Gtk_source_style_scheme_get_id func(
		scheme *GtkSourceStyleScheme) *Gchar

	Gtk_source_style_scheme_get_name func(
		scheme *GtkSourceStyleScheme) *Gchar

	Gtk_source_style_scheme_get_description func(
		scheme *GtkSourceStyleScheme) *Gchar

	Gtk_source_style_scheme_get_authors func(
		scheme *GtkSourceStyleScheme) **Gchar

	Gtk_source_style_scheme_get_filename func(
		scheme *GtkSourceStyleScheme) *Gchar

	Gtk_source_style_scheme_get_style func(
		scheme *GtkSourceStyleScheme,
		style_id *Gchar) *GtkSourceStyle

	_gtk_source_style_scheme_new_from_file func(
		filename *Gchar) *GtkSourceStyleScheme

	_gtk_source_style_scheme_get_default func() *GtkSourceStyleScheme

	_gtk_source_style_scheme_get_parent_id func(
		scheme *GtkSourceStyleScheme) *Gchar

	_gtk_source_style_scheme_set_parent func(
		scheme *GtkSourceStyleScheme,
		parent_scheme *GtkSourceStyleScheme)

	_gtk_source_style_scheme_apply func(
		scheme *GtkSourceStyleScheme,
		widget *GtkWidget)

	_gtk_source_style_scheme_get_matching_brackets_style func(
		scheme *GtkSourceStyleScheme) *GtkSourceStyle

	_gtk_source_style_scheme_get_right_margin_style func(
		scheme *GtkSourceStyleScheme) *GtkSourceStyle

	_gtk_source_style_scheme_get_draw_spaces_style func(
		scheme *GtkSourceStyleScheme) *GtkSourceStyle

	_gtk_source_style_scheme_get_current_line_color func(
		scheme *GtkSourceStyleScheme,
		color *GdkColor) Gboolean

	Gtk_source_undo_manager_get_type func() GType

	Gtk_source_undo_manager_can_undo func(
		manager *GtkSourceUndoManager) Gboolean

	Gtk_source_undo_manager_can_redo func(
		manager *GtkSourceUndoManager) Gboolean

	Gtk_source_undo_manager_undo func(
		manager *GtkSourceUndoManager)

	Gtk_source_undo_manager_redo func(
		manager *GtkSourceUndoManager)

	Gtk_source_undo_manager_begin_not_undoable_action func(
		manager *GtkSourceUndoManager)

	Gtk_source_undo_manager_end_not_undoable_action func(
		manager *GtkSourceUndoManager)

	Gtk_source_undo_manager_can_undo_changed func(
		manager *GtkSourceUndoManager)

	Gtk_source_undo_manager_can_redo_changed func(
		manager *GtkSourceUndoManager)

	Gtk_source_buffer_get_type func() GType

	Gtk_source_buffer_new func(
		table *GtkTextTagTable) *GtkSourceBuffer

	Gtk_source_buffer_new_with_language func(
		language *GtkSourceLanguage) *GtkSourceBuffer

	Gtk_source_buffer_get_highlight_syntax func(
		buffer *GtkSourceBuffer) Gboolean

	Gtk_source_buffer_set_highlight_syntax func(
		buffer *GtkSourceBuffer,
		highlight Gboolean)

	Gtk_source_buffer_get_highlight_matching_brackets func(
		buffer *GtkSourceBuffer) Gboolean

	Gtk_source_buffer_set_highlight_matching_brackets func(
		buffer *GtkSourceBuffer,
		highlight Gboolean)

	Gtk_source_buffer_get_max_undo_levels func(
		buffer *GtkSourceBuffer) Gint

	Gtk_source_buffer_set_max_undo_levels func(
		buffer *GtkSourceBuffer,
		max_undo_levels Gint)

	Gtk_source_buffer_get_language func(
		buffer *GtkSourceBuffer) *GtkSourceLanguage

	Gtk_source_buffer_set_language func(
		buffer *GtkSourceBuffer,
		language *GtkSourceLanguage)

	Gtk_source_buffer_can_undo func(
		buffer *GtkSourceBuffer) Gboolean

	Gtk_source_buffer_can_redo func(
		buffer *GtkSourceBuffer) Gboolean

	Gtk_source_buffer_get_style_scheme func(
		buffer *GtkSourceBuffer) *GtkSourceStyleScheme

	Gtk_source_buffer_set_style_scheme func(
		buffer *GtkSourceBuffer,
		scheme *GtkSourceStyleScheme)

	Gtk_source_buffer_ensure_highlight func(
		buffer *GtkSourceBuffer,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_source_buffer_undo func(
		buffer *GtkSourceBuffer)

	Gtk_source_buffer_redo func(
		buffer *GtkSourceBuffer)

	Gtk_source_buffer_begin_not_undoable_action func(
		buffer *GtkSourceBuffer)

	Gtk_source_buffer_end_not_undoable_action func(
		buffer *GtkSourceBuffer)

	Gtk_source_buffer_create_source_mark func(
		buffer *GtkSourceBuffer,
		name *Gchar,
		category *Gchar,
		where *GtkTextIter) *GtkSourceMark

	Gtk_source_buffer_forward_iter_to_source_mark func(
		buffer *GtkSourceBuffer,
		iter *GtkTextIter,
		category *Gchar) Gboolean

	Gtk_source_buffer_backward_iter_to_source_mark func(
		buffer *GtkSourceBuffer,
		iter *GtkTextIter,
		category *Gchar) Gboolean

	Gtk_source_buffer_get_source_marks_at_iter func(
		buffer *GtkSourceBuffer,
		iter *GtkTextIter,
		category *Gchar) *GSList

	Gtk_source_buffer_get_source_marks_at_line func(
		buffer *GtkSourceBuffer,
		line Gint,
		category *Gchar) *GSList

	Gtk_source_buffer_remove_source_marks func(
		buffer *GtkSourceBuffer,
		start *GtkTextIter,
		end *GtkTextIter,
		category *Gchar)

	Gtk_source_buffer_iter_has_context_class func(
		buffer *GtkSourceBuffer,
		iter *GtkTextIter,
		context_class *Gchar) Gboolean

	Gtk_source_buffer_get_context_classes_at_iter func(
		buffer *GtkSourceBuffer,
		iter *GtkTextIter) **Gchar

	Gtk_source_buffer_iter_forward_to_context_class_toggle func(
		buffer *GtkSourceBuffer,
		iter *GtkTextIter,
		context_class *Gchar) Gboolean

	Gtk_source_buffer_iter_backward_to_context_class_toggle func(
		buffer *GtkSourceBuffer,
		iter *GtkTextIter,
		context_class *Gchar) Gboolean

	Gtk_source_buffer_get_undo_manager func(
		buffer *GtkSourceBuffer) *GtkSourceUndoManager

	Gtk_source_buffer_set_undo_manager func(
		buffer *GtkSourceBuffer,
		manager *GtkSourceUndoManager)

	_gtk_source_buffer_update_highlight func(
		buffer *GtkSourceBuffer,
		start *GtkTextIter,
		end *GtkTextIter,
		synchronous Gboolean)

	_gtk_source_buffer_source_mark_next func(
		buffer *GtkSourceBuffer,
		mark *GtkSourceMark,
		category *Gchar) *GtkSourceMark

	_gtk_source_buffer_source_mark_prev func(
		buffer *GtkSourceBuffer,
		mark *GtkSourceMark,
		category *Gchar) *GtkSourceMark

	_gtk_source_buffer_get_bracket_match_tag func(
		buffer *GtkSourceBuffer) *GtkTextTag

	Gtk_source_completion_info_get_type func() GType

	Gtk_source_completion_info_new func() *GtkSourceCompletionInfo

	Gtk_source_completion_info_move_to_iter func(
		info *GtkSourceCompletionInfo,
		view *GtkTextView,
		iter *GtkTextIter)

	Gtk_source_completion_info_set_sizing func(
		info *GtkSourceCompletionInfo,
		width Gint,
		height Gint,
		shrink_width Gboolean,
		shrink_height Gboolean)

	Gtk_source_completion_info_set_widget func(
		info *GtkSourceCompletionInfo,
		widget *GtkWidget)

	Gtk_source_completion_info_get_widget func(
		info *GtkSourceCompletionInfo) *GtkWidget

	Gtk_source_completion_info_process_resize func(
		info *GtkSourceCompletionInfo)

	Gtk_source_completion_proposal_get_type func() GType

	Gtk_source_completion_proposal_get_label func(
		proposal *GtkSourceCompletionProposal) *Gchar

	Gtk_source_completion_proposal_get_markup func(
		proposal *GtkSourceCompletionProposal) *Gchar

	Gtk_source_completion_proposal_get_text func(
		proposal *GtkSourceCompletionProposal) *Gchar

	Gtk_source_completion_proposal_get_icon func(
		proposal *GtkSourceCompletionProposal) *GdkPixbuf

	Gtk_source_completion_proposal_get_info func(
		proposal *GtkSourceCompletionProposal) *Gchar

	Gtk_source_completion_proposal_changed func(
		proposal *GtkSourceCompletionProposal)

	Gtk_source_completion_proposal_hash func(
		proposal *GtkSourceCompletionProposal) Guint

	Gtk_source_completion_proposal_equal func(
		proposal *GtkSourceCompletionProposal,
		other *GtkSourceCompletionProposal) Gboolean

	Gtk_source_completion_context_get_type func() GType

	Gtk_source_completion_context_add_proposals func(
		context *GtkSourceCompletionContext,
		provider *GtkSourceCompletionProvider,
		proposals *GList,
		finished Gboolean)

	Gtk_source_completion_context_get_iter func(
		context *GtkSourceCompletionContext,
		iter *GtkTextIter)

	Gtk_source_completion_context_get_activation func(
		context *GtkSourceCompletionContext) GtkSourceCompletionActivation

	_gtk_source_completion_context_new func(
		completion *GtkSourceCompletion,
		position *GtkTextIter) *GtkSourceCompletionContext

	_gtk_source_completion_context_cancel func(
		context *GtkSourceCompletionContext)

	Gtk_source_completion_provider_get_type func() GType

	Gtk_source_completion_provider_get_name func(
		provider *GtkSourceCompletionProvider) *Gchar

	Gtk_source_completion_provider_get_icon func(
		provider *GtkSourceCompletionProvider) *GdkPixbuf

	Gtk_source_completion_provider_populate func(
		provider *GtkSourceCompletionProvider,
		context *GtkSourceCompletionContext)

	Gtk_source_completion_provider_get_activation func(
		provider *GtkSourceCompletionProvider) GtkSourceCompletionActivation

	Gtk_source_completion_provider_match func(
		provider *GtkSourceCompletionProvider,
		context *GtkSourceCompletionContext) Gboolean

	Gtk_source_completion_provider_get_info_widget func(
		provider *GtkSourceCompletionProvider,
		proposal *GtkSourceCompletionProposal) *GtkWidget

	Gtk_source_completion_provider_update_info func(
		provider *GtkSourceCompletionProvider,
		proposal *GtkSourceCompletionProposal,
		info *GtkSourceCompletionInfo)

	Gtk_source_completion_provider_get_start_iter func(
		provider *GtkSourceCompletionProvider,
		context *GtkSourceCompletionContext,
		proposal *GtkSourceCompletionProposal,
		iter *GtkTextIter) Gboolean

	Gtk_source_completion_provider_activate_proposal func(
		provider *GtkSourceCompletionProvider,
		proposal *GtkSourceCompletionProposal,
		iter *GtkTextIter) Gboolean

	Gtk_source_completion_provider_get_interactive_delay func(
		provider *GtkSourceCompletionProvider) Gint

	Gtk_source_completion_provider_get_priority func(
		provider *GtkSourceCompletionProvider) Gint

	Gtk_source_completion_get_type func() GType

	Gtk_source_completion_error_quark func() GQuark

	Gtk_source_completion_add_provider func(
		completion *GtkSourceCompletion,
		provider *GtkSourceCompletionProvider,
		error **GError) Gboolean

	Gtk_source_completion_remove_provider func(
		completion *GtkSourceCompletion,
		provider *GtkSourceCompletionProvider,
		error **GError) Gboolean

	Gtk_source_completion_get_providers func(
		completion *GtkSourceCompletion) *GList

	Gtk_source_completion_show func(
		completion *GtkSourceCompletion,
		providers *GList,
		context *GtkSourceCompletionContext) Gboolean

	Gtk_source_completion_hide func(
		completion *GtkSourceCompletion)

	Gtk_source_completion_get_info_window func(
		completion *GtkSourceCompletion) *GtkSourceCompletionInfo

	Gtk_source_completion_get_view func(
		completion *GtkSourceCompletion) *GtkSourceView

	Gtk_source_completion_create_context func(
		completion *GtkSourceCompletion,
		position *GtkTextIter) *GtkSourceCompletionContext

	Gtk_source_completion_move_window func(
		completion *GtkSourceCompletion,
		iter *GtkTextIter)

	_gtk_source_completion_add_proposals func(
		completion *GtkSourceCompletion,
		context *GtkSourceCompletionContext,
		provider *GtkSourceCompletionProvider,
		proposals *GList,
		finished Gboolean)

	Gtk_source_completion_block_interactive func(
		completion *GtkSourceCompletion)

	Gtk_source_completion_unblock_interactive func(
		completion *GtkSourceCompletion)

	Gtk_source_gutter_get_type func() GType

	Gtk_source_gutter_get_window func(
		gutter *GtkSourceGutter) *GdkWindow

	Gtk_source_gutter_insert func(
		gutter *GtkSourceGutter,
		renderer *GtkCellRenderer,
		position Gint)

	Gtk_source_gutter_reorder func(
		gutter *GtkSourceGutter,
		renderer *GtkCellRenderer,
		position Gint)

	Gtk_source_gutter_remove func(
		gutter *GtkSourceGutter,
		renderer *GtkCellRenderer)

	Gtk_source_gutter_set_cell_data_func func(
		gutter *GtkSourceGutter,
		renderer *GtkCellRenderer,
		f GtkSourceGutterDataFunc,
		func_data Gpointer,
		destroy GDestroyNotify)

	Gtk_source_gutter_set_cell_size_func func(
		gutter *GtkSourceGutter,
		renderer *GtkCellRenderer,
		f GtkSourceGutterSizeFunc,
		func_data Gpointer,
		destroy GDestroyNotify)

	Gtk_source_gutter_queue_draw func(
		gutter *GtkSourceGutter)

	Gtk_source_view_get_type func() GType

	Gtk_source_view_new func() *GtkWidget

	Gtk_source_view_new_with_buffer func(
		buffer *GtkSourceBuffer) *GtkWidget

	Gtk_source_view_set_show_line_numbers func(
		view *GtkSourceView,
		show Gboolean)

	Gtk_source_view_get_show_line_numbers func(
		view *GtkSourceView) Gboolean

	Gtk_source_view_set_tab_width func(
		view *GtkSourceView,
		width Guint)

	Gtk_source_view_get_tab_width func(
		view *GtkSourceView) Guint

	Gtk_source_view_set_indent_width func(
		view *GtkSourceView,
		width Gint)

	Gtk_source_view_get_indent_width func(
		view *GtkSourceView) Gint

	Gtk_source_view_set_auto_indent func(
		view *GtkSourceView,
		enable Gboolean)

	Gtk_source_view_get_auto_indent func(
		view *GtkSourceView) Gboolean

	Gtk_source_view_set_insert_spaces_instead_of_tabs func(
		view *GtkSourceView,
		enable Gboolean)

	Gtk_source_view_get_insert_spaces_instead_of_tabs func(
		view *GtkSourceView) Gboolean

	Gtk_source_view_set_indent_on_tab func(
		view *GtkSourceView,
		enable Gboolean)

	Gtk_source_view_get_indent_on_tab func(
		view *GtkSourceView) Gboolean

	Gtk_source_view_set_highlight_current_line func(
		view *GtkSourceView,
		show Gboolean)

	Gtk_source_view_get_highlight_current_line func(
		view *GtkSourceView) Gboolean

	Gtk_source_view_set_show_right_margin func(
		view *GtkSourceView,
		show Gboolean)

	Gtk_source_view_get_show_right_margin func(
		view *GtkSourceView) Gboolean

	Gtk_source_view_set_right_margin_position func(
		view *GtkSourceView,
		pos Guint)

	Gtk_source_view_get_right_margin_position func(
		view *GtkSourceView) Guint

	Gtk_source_view_set_show_line_marks func(
		view *GtkSourceView,
		show Gboolean)

	Gtk_source_view_get_show_line_marks func(
		view *GtkSourceView) Gboolean

	Gtk_source_view_set_mark_category_pixbuf func(
		view *GtkSourceView,
		category *Gchar,
		pixbuf *GdkPixbuf)

	Gtk_source_view_set_mark_category_icon_from_pixbuf func(
		view *GtkSourceView,
		category *Gchar,
		pixbuf *GdkPixbuf)

	Gtk_source_view_set_mark_category_icon_from_stock func(
		view *GtkSourceView,
		category *Gchar,
		stock_id *Gchar)

	Gtk_source_view_set_mark_category_icon_from_icon_name func(
		view *GtkSourceView,
		category *Gchar,
		name *Gchar)

	Gtk_source_view_get_mark_category_pixbuf func(
		view *GtkSourceView,
		category *Gchar) *GdkPixbuf

	Gtk_source_view_set_mark_category_background func(
		view *GtkSourceView,
		category *Gchar,
		color *GdkColor)

	Gtk_source_view_get_mark_category_background func(
		view *GtkSourceView,
		category *Gchar,
		dest *GdkColor) Gboolean

	Gtk_source_view_set_mark_category_tooltip_func func(
		view *GtkSourceView,
		category *Gchar,
		f GtkSourceViewMarkTooltipFunc,
		user_data Gpointer,
		user_data_notify GDestroyNotify)

	Gtk_source_view_set_mark_category_tooltip_markup_func func(
		view *GtkSourceView,
		category *Gchar,
		markup_func GtkSourceViewMarkTooltipFunc,
		user_data Gpointer,
		user_data_notify GDestroyNotify)

	Gtk_source_view_set_mark_category_priority func(
		view *GtkSourceView,
		category *Gchar,
		priority Gint)

	Gtk_source_view_get_mark_category_priority func(
		view *GtkSourceView,
		category *Gchar) Gint

	Gtk_source_view_set_smart_home_end func(
		view *GtkSourceView,
		smart_he GtkSourceSmartHomeEndType)

	Gtk_source_view_get_smart_home_end func(
		view *GtkSourceView) GtkSourceSmartHomeEndType

	Gtk_source_view_set_draw_spaces func(
		view *GtkSourceView,
		flags GtkSourceDrawSpacesFlags)

	Gtk_source_view_get_draw_spaces func(
		view *GtkSourceView) GtkSourceDrawSpacesFlags

	Gtk_source_view_get_completion func(
		view *GtkSourceView) *GtkSourceCompletion

	Gtk_source_view_get_gutter func(
		view *GtkSourceView,
		window_type GtkTextWindowType) *GtkSourceGutter

	Gtk_source_completion_item_get_type func() GType

	Gtk_source_completion_item_new func(
		label *Gchar,
		text *Gchar,
		icon *GdkPixbuf,
		info *Gchar) *GtkSourceCompletionItem

	Gtk_source_completion_item_new_with_markup func(
		markup *Gchar,
		text *Gchar,
		icon *GdkPixbuf,
		info *Gchar) *GtkSourceCompletionItem

	Gtk_source_completion_item_new_from_stock func(
		label *Gchar,
		text *Gchar,
		stock *Gchar,
		info *Gchar) *GtkSourceCompletionItem

	Gtk_source_search_flags_get_type func() GType

	Gtk_source_view_gutter_position_get_type func() GType

	Gtk_source_smart_home_end_type_get_type func() GType

	Gtk_source_draw_spaces_flags_get_type func() GType

	Gtk_source_completion_error_get_type func() GType

	Gtk_source_completion_activation_get_type func() GType

	Gtk_source_completion_words_get_type func() GType

	Gtk_source_completion_words_new func(
		name *Gchar,
		icon *GdkPixbuf) *GtkSourceCompletionWords

	Gtk_source_completion_words_register func(
		words *GtkSourceCompletionWords,
		buffer *GtkTextBuffer)

	Gtk_source_completion_words_unregister func(
		words *GtkSourceCompletionWords,
		buffer *GtkTextBuffer)

	Gtk_source_iter_forward_search func(
		iter *GtkTextIter,
		str *Gchar,
		flags GtkSourceSearchFlags,
		match_start *GtkTextIter,
		match_end *GtkTextIter,
		limit *GtkTextIter) Gboolean

	Gtk_source_iter_backward_search func(
		iter *GtkTextIter,
		str *Gchar,
		flags GtkSourceSearchFlags,
		match_start *GtkTextIter,
		match_end *GtkTextIter,
		limit *GtkTextIter) Gboolean

	Gtk_source_language_manager_get_type func() GType

	Gtk_source_language_manager_new func() *GtkSourceLanguageManager

	Gtk_source_language_manager_get_default func() *GtkSourceLanguageManager

	Gtk_source_language_manager_get_search_path func(
		lm *GtkSourceLanguageManager) **Gchar

	Gtk_source_language_manager_set_search_path func(
		lm *GtkSourceLanguageManager,
		dirs **Gchar)

	Gtk_source_language_manager_get_language_ids func(
		lm *GtkSourceLanguageManager) **Gchar

	Gtk_source_language_manager_get_language func(
		lm *GtkSourceLanguageManager,
		id *Gchar) *GtkSourceLanguage

	Gtk_source_language_manager_guess_language func(
		lm *GtkSourceLanguageManager,
		filename *Gchar,
		content_type *Gchar) *GtkSourceLanguage

	Gtk_source_print_compositor_get_type func() GType

	Gtk_source_print_compositor_new func(
		buffer *GtkSourceBuffer) *GtkSourcePrintCompositor

	Gtk_source_print_compositor_new_from_view func(
		view *GtkSourceView) *GtkSourcePrintCompositor

	Gtk_source_print_compositor_get_buffer func(
		compositor *GtkSourcePrintCompositor) *GtkSourceBuffer

	Gtk_source_print_compositor_set_tab_width func(
		compositor *GtkSourcePrintCompositor,
		width Guint)

	Gtk_source_print_compositor_get_tab_width func(
		compositor *GtkSourcePrintCompositor) Guint

	Gtk_source_print_compositor_set_wrap_mode func(
		compositor *GtkSourcePrintCompositor,
		wrap_mode GtkWrapMode)

	Gtk_source_print_compositor_get_wrap_mode func(
		compositor *GtkSourcePrintCompositor) GtkWrapMode

	Gtk_source_print_compositor_set_highlight_syntax func(
		compositor *GtkSourcePrintCompositor,
		highlight Gboolean)

	Gtk_source_print_compositor_get_highlight_syntax func(
		compositor *GtkSourcePrintCompositor) Gboolean

	Gtk_source_print_compositor_set_print_line_numbers func(
		compositor *GtkSourcePrintCompositor,
		interval Guint)

	Gtk_source_print_compositor_get_print_line_numbers func(
		compositor *GtkSourcePrintCompositor) Guint

	Gtk_source_print_compositor_set_body_font_name func(
		compositor *GtkSourcePrintCompositor,
		font_name *Gchar)

	Gtk_source_print_compositor_get_body_font_name func(
		compositor *GtkSourcePrintCompositor) *Gchar

	Gtk_source_print_compositor_set_line_numbers_font_name func(
		compositor *GtkSourcePrintCompositor,
		font_name *Gchar)

	Gtk_source_print_compositor_get_line_numbers_font_name func(
		compositor *GtkSourcePrintCompositor) *Gchar

	Gtk_source_print_compositor_set_header_font_name func(
		compositor *GtkSourcePrintCompositor,
		font_name *Gchar)

	Gtk_source_print_compositor_get_header_font_name func(
		compositor *GtkSourcePrintCompositor) *Gchar

	Gtk_source_print_compositor_set_footer_font_name func(
		compositor *GtkSourcePrintCompositor,
		font_name *Gchar)

	Gtk_source_print_compositor_get_footer_font_name func(
		compositor *GtkSourcePrintCompositor) *Gchar

	Gtk_source_print_compositor_get_top_margin func(
		compositor *GtkSourcePrintCompositor,
		unit GtkUnit) Gdouble

	Gtk_source_print_compositor_set_top_margin func(
		compositor *GtkSourcePrintCompositor,
		margin Gdouble,
		unit GtkUnit)

	Gtk_source_print_compositor_get_bottom_margin func(
		compositor *GtkSourcePrintCompositor,
		unit GtkUnit) Gdouble

	Gtk_source_print_compositor_set_bottom_margin func(
		compositor *GtkSourcePrintCompositor,
		margin Gdouble,
		unit GtkUnit)

	Gtk_source_print_compositor_get_left_margin func(
		compositor *GtkSourcePrintCompositor,
		unit GtkUnit) Gdouble

	Gtk_source_print_compositor_set_left_margin func(
		compositor *GtkSourcePrintCompositor,
		margin Gdouble,
		unit GtkUnit)

	Gtk_source_print_compositor_get_right_margin func(
		compositor *GtkSourcePrintCompositor,
		unit GtkUnit) Gdouble

	Gtk_source_print_compositor_set_right_margin func(
		compositor *GtkSourcePrintCompositor,
		margin Gdouble,
		unit GtkUnit)

	Gtk_source_print_compositor_set_print_header func(
		compositor *GtkSourcePrintCompositor,
		print Gboolean)

	Gtk_source_print_compositor_get_print_header func(
		compositor *GtkSourcePrintCompositor) Gboolean

	Gtk_source_print_compositor_set_print_footer func(
		compositor *GtkSourcePrintCompositor,
		print Gboolean)

	Gtk_source_print_compositor_get_print_footer func(
		compositor *GtkSourcePrintCompositor) Gboolean

	Gtk_source_print_compositor_set_header_format func(
		compositor *GtkSourcePrintCompositor,
		separator Gboolean,
		left *Gchar,
		center *Gchar,
		right *Gchar)

	Gtk_source_print_compositor_set_footer_format func(
		compositor *GtkSourcePrintCompositor,
		separator Gboolean,
		left *Gchar,
		center *Gchar,
		right *Gchar)

	Gtk_source_print_compositor_get_n_pages func(
		compositor *GtkSourcePrintCompositor) Gint

	Gtk_source_print_compositor_paginate func(
		compositor *GtkSourcePrintCompositor,
		context *GtkPrintContext) Gboolean

	Gtk_source_print_compositor_get_pagination_progress func(
		compositor *GtkSourcePrintCompositor) Gdouble

	Gtk_source_print_compositor_draw_page func(
		compositor *GtkSourcePrintCompositor,
		context *GtkPrintContext,
		page_nr Gint)

	Gtk_source_style_scheme_manager_get_type func() GType

	Gtk_source_style_scheme_manager_new func() *GtkSourceStyleSchemeManager

	Gtk_source_style_scheme_manager_get_default func() *GtkSourceStyleSchemeManager

	Gtk_source_style_scheme_manager_set_search_path func(
		manager *GtkSourceStyleSchemeManager,
		path **Gchar)

	Gtk_source_style_scheme_manager_append_search_path func(
		manager *GtkSourceStyleSchemeManager,
		path *Gchar)

	Gtk_source_style_scheme_manager_prepend_search_path func(
		manager *GtkSourceStyleSchemeManager,
		path *Gchar)

	Gtk_source_style_scheme_manager_get_search_path func(
		manager *GtkSourceStyleSchemeManager) **Gchar

	Gtk_source_style_scheme_manager_force_rescan func(
		manager *GtkSourceStyleSchemeManager)

	Gtk_source_style_scheme_manager_get_scheme_ids func(
		manager *GtkSourceStyleSchemeManager) **Gchar

	Gtk_source_style_scheme_manager_get_scheme func(
		manager *GtkSourceStyleSchemeManager,
		scheme_id *Gchar) *GtkSourceStyleScheme
)

var dll = "libgtksourceview-2.0-0.dll"

var apiList = outside.Apis{
	{"gtk_source_buffer_backward_iter_to_source_mark", &Gtk_source_buffer_backward_iter_to_source_mark},
	{"gtk_source_buffer_begin_not_undoable_action", &Gtk_source_buffer_begin_not_undoable_action},
	{"gtk_source_buffer_can_redo", &Gtk_source_buffer_can_redo},
	{"gtk_source_buffer_can_undo", &Gtk_source_buffer_can_undo},
	{"gtk_source_buffer_create_source_mark", &Gtk_source_buffer_create_source_mark},
	{"gtk_source_buffer_end_not_undoable_action", &Gtk_source_buffer_end_not_undoable_action},
	{"gtk_source_buffer_ensure_highlight", &Gtk_source_buffer_ensure_highlight},
	{"gtk_source_buffer_forward_iter_to_source_mark", &Gtk_source_buffer_forward_iter_to_source_mark},
	{"gtk_source_buffer_get_context_classes_at_iter", &Gtk_source_buffer_get_context_classes_at_iter},
	{"gtk_source_buffer_get_highlight_matching_brackets", &Gtk_source_buffer_get_highlight_matching_brackets},
	{"gtk_source_buffer_get_highlight_syntax", &Gtk_source_buffer_get_highlight_syntax},
	{"gtk_source_buffer_get_language", &Gtk_source_buffer_get_language},
	{"gtk_source_buffer_get_max_undo_levels", &Gtk_source_buffer_get_max_undo_levels},
	{"gtk_source_buffer_get_source_marks_at_iter", &Gtk_source_buffer_get_source_marks_at_iter},
	{"gtk_source_buffer_get_source_marks_at_line", &Gtk_source_buffer_get_source_marks_at_line},
	{"gtk_source_buffer_get_style_scheme", &Gtk_source_buffer_get_style_scheme},
	{"gtk_source_buffer_get_type", &Gtk_source_buffer_get_type},
	{"gtk_source_buffer_get_undo_manager", &Gtk_source_buffer_get_undo_manager},
	{"gtk_source_buffer_iter_backward_to_context_class_toggle", &Gtk_source_buffer_iter_backward_to_context_class_toggle},
	{"gtk_source_buffer_iter_forward_to_context_class_toggle", &Gtk_source_buffer_iter_forward_to_context_class_toggle},
	{"gtk_source_buffer_iter_has_context_class", &Gtk_source_buffer_iter_has_context_class},
	{"gtk_source_buffer_new", &Gtk_source_buffer_new},
	{"gtk_source_buffer_new_with_language", &Gtk_source_buffer_new_with_language},
	{"gtk_source_buffer_redo", &Gtk_source_buffer_redo},
	{"gtk_source_buffer_remove_source_marks", &Gtk_source_buffer_remove_source_marks},
	{"gtk_source_buffer_set_highlight_matching_brackets", &Gtk_source_buffer_set_highlight_matching_brackets},
	{"gtk_source_buffer_set_highlight_syntax", &Gtk_source_buffer_set_highlight_syntax},
	{"gtk_source_buffer_set_language", &Gtk_source_buffer_set_language},
	{"gtk_source_buffer_set_max_undo_levels", &Gtk_source_buffer_set_max_undo_levels},
	{"gtk_source_buffer_set_style_scheme", &Gtk_source_buffer_set_style_scheme},
	{"gtk_source_buffer_set_undo_manager", &Gtk_source_buffer_set_undo_manager},
	{"gtk_source_buffer_undo", &Gtk_source_buffer_undo},
	{"gtk_source_completion_activation_get_type", &Gtk_source_completion_activation_get_type},
	{"gtk_source_completion_add_provider", &Gtk_source_completion_add_provider},
	{"gtk_source_completion_block_interactive", &Gtk_source_completion_block_interactive},
	{"gtk_source_completion_context_add_proposals", &Gtk_source_completion_context_add_proposals},
	{"gtk_source_completion_context_get_activation", &Gtk_source_completion_context_get_activation},
	{"gtk_source_completion_context_get_iter", &Gtk_source_completion_context_get_iter},
	{"gtk_source_completion_context_get_type", &Gtk_source_completion_context_get_type},
	{"gtk_source_completion_create_context", &Gtk_source_completion_create_context},
	{"gtk_source_completion_error_get_type", &Gtk_source_completion_error_get_type},
	{"gtk_source_completion_error_quark", &Gtk_source_completion_error_quark},
	{"gtk_source_completion_get_info_window", &Gtk_source_completion_get_info_window},
	{"gtk_source_completion_get_providers", &Gtk_source_completion_get_providers},
	{"gtk_source_completion_get_type", &Gtk_source_completion_get_type},
	{"gtk_source_completion_get_view", &Gtk_source_completion_get_view},
	{"gtk_source_completion_hide", &Gtk_source_completion_hide},
	{"gtk_source_completion_info_get_type", &Gtk_source_completion_info_get_type},
	{"gtk_source_completion_info_get_widget", &Gtk_source_completion_info_get_widget},
	{"gtk_source_completion_info_move_to_iter", &Gtk_source_completion_info_move_to_iter},
	{"gtk_source_completion_info_new", &Gtk_source_completion_info_new},
	{"gtk_source_completion_info_process_resize", &Gtk_source_completion_info_process_resize},
	{"gtk_source_completion_info_set_sizing", &Gtk_source_completion_info_set_sizing},
	{"gtk_source_completion_info_set_widget", &Gtk_source_completion_info_set_widget},
	{"gtk_source_completion_item_get_type", &Gtk_source_completion_item_get_type},
	{"gtk_source_completion_item_new", &Gtk_source_completion_item_new},
	{"gtk_source_completion_item_new_from_stock", &Gtk_source_completion_item_new_from_stock},
	{"gtk_source_completion_item_new_with_markup", &Gtk_source_completion_item_new_with_markup},
	// {"gtk_source_completion_model_append", &Gtk_source_completion_model_append},
	// {"gtk_source_completion_model_begin", &Gtk_source_completion_model_begin},
	// {"gtk_source_completion_model_cancel", &Gtk_source_completion_model_cancel},
	// {"gtk_source_completion_model_clear", &Gtk_source_completion_model_clear},
	// {"gtk_source_completion_model_end", &Gtk_source_completion_model_end},
	// {"gtk_source_completion_model_get_providers", &Gtk_source_completion_model_get_providers},
	// {"gtk_source_completion_model_get_type", &Gtk_source_completion_model_get_type},
	// {"gtk_source_completion_model_get_visible_providers", &Gtk_source_completion_model_get_visible_providers},
	// {"gtk_source_completion_model_is_empty", &Gtk_source_completion_model_is_empty},
	// {"gtk_source_completion_model_iter_equal", &Gtk_source_completion_model_iter_equal},
	// {"gtk_source_completion_model_iter_is_header", &Gtk_source_completion_model_iter_is_header},
	// {"gtk_source_completion_model_iter_last", &Gtk_source_completion_model_iter_last},
	// {"gtk_source_completion_model_iter_previous", &Gtk_source_completion_model_iter_previous},
	// {"gtk_source_completion_model_n_proposals", &Gtk_source_completion_model_n_proposals},
	// {"gtk_source_completion_model_new", &Gtk_source_completion_model_new},
	// {"gtk_source_completion_model_set_show_headers", &Gtk_source_completion_model_set_show_headers},
	// {"gtk_source_completion_model_set_visible_providers", &Gtk_source_completion_model_set_visible_providers},
	{"gtk_source_completion_move_window", &Gtk_source_completion_move_window},
	// {"gtk_source_completion_new", &Gtk_source_completion_new},
	{"gtk_source_completion_proposal_changed", &Gtk_source_completion_proposal_changed},
	{"gtk_source_completion_proposal_equal", &Gtk_source_completion_proposal_equal},
	{"gtk_source_completion_proposal_get_icon", &Gtk_source_completion_proposal_get_icon},
	{"gtk_source_completion_proposal_get_info", &Gtk_source_completion_proposal_get_info},
	{"gtk_source_completion_proposal_get_label", &Gtk_source_completion_proposal_get_label},
	{"gtk_source_completion_proposal_get_markup", &Gtk_source_completion_proposal_get_markup},
	{"gtk_source_completion_proposal_get_text", &Gtk_source_completion_proposal_get_text},
	{"gtk_source_completion_proposal_get_type", &Gtk_source_completion_proposal_get_type},
	{"gtk_source_completion_proposal_hash", &Gtk_source_completion_proposal_hash},
	{"gtk_source_completion_provider_activate_proposal", &Gtk_source_completion_provider_activate_proposal},
	{"gtk_source_completion_provider_get_activation", &Gtk_source_completion_provider_get_activation},
	{"gtk_source_completion_provider_get_icon", &Gtk_source_completion_provider_get_icon},
	{"gtk_source_completion_provider_get_info_widget", &Gtk_source_completion_provider_get_info_widget},
	{"gtk_source_completion_provider_get_interactive_delay", &Gtk_source_completion_provider_get_interactive_delay},
	{"gtk_source_completion_provider_get_name", &Gtk_source_completion_provider_get_name},
	{"gtk_source_completion_provider_get_priority", &Gtk_source_completion_provider_get_priority},
	{"gtk_source_completion_provider_get_start_iter", &Gtk_source_completion_provider_get_start_iter},
	{"gtk_source_completion_provider_get_type", &Gtk_source_completion_provider_get_type},
	{"gtk_source_completion_provider_match", &Gtk_source_completion_provider_match},
	{"gtk_source_completion_provider_populate", &Gtk_source_completion_provider_populate},
	{"gtk_source_completion_provider_update_info", &Gtk_source_completion_provider_update_info},
	{"gtk_source_completion_remove_provider", &Gtk_source_completion_remove_provider},
	{"gtk_source_completion_show", &Gtk_source_completion_show},
	{"gtk_source_completion_unblock_interactive", &Gtk_source_completion_unblock_interactive},
	// {"gtk_source_completion_utils_get_word", &Gtk_source_completion_utils_get_word},
	// {"gtk_source_completion_utils_get_word_iter", &Gtk_source_completion_utils_get_word_iter},
	// {"gtk_source_completion_utils_is_separator", &Gtk_source_completion_utils_is_separator},
	// {"gtk_source_completion_utils_move_to_cursor", &Gtk_source_completion_utils_move_to_cursor},
	// {"gtk_source_completion_utils_move_to_iter", &Gtk_source_completion_utils_move_to_iter},
	// {"gtk_source_completion_utils_replace_current_word", &Gtk_source_completion_utils_replace_current_word},
	// {"gtk_source_completion_utils_replace_word", &Gtk_source_completion_utils_replace_word},
	// {"gtk_source_completion_words_buffer_get_buffer", &Gtk_source_completion_words_buffer_get_buffer},
	// {"gtk_source_completion_words_buffer_get_mark", &Gtk_source_completion_words_buffer_get_mark},
	// {"gtk_source_completion_words_buffer_get_type", &Gtk_source_completion_words_buffer_get_type},
	// {"gtk_source_completion_words_buffer_new", &Gtk_source_completion_words_buffer_new},
	// {"gtk_source_completion_words_buffer_set_minimum_word_size", &Gtk_source_completion_words_buffer_set_minimum_word_size},
	// {"gtk_source_completion_words_buffer_set_scan_batch_size", &Gtk_source_completion_words_buffer_set_scan_batch_size},
	{"gtk_source_completion_words_get_type", &Gtk_source_completion_words_get_type},
	// {"gtk_source_completion_words_library_add_word", &Gtk_source_completion_words_library_add_word},
	// {"gtk_source_completion_words_library_find", &Gtk_source_completion_words_library_find},
	// {"gtk_source_completion_words_library_find_first", &Gtk_source_completion_words_library_find_first},
	// {"gtk_source_completion_words_library_find_next", &Gtk_source_completion_words_library_find_next},
	// {"gtk_source_completion_words_library_get_proposal", &Gtk_source_completion_words_library_get_proposal},
	// {"gtk_source_completion_words_library_get_type", &Gtk_source_completion_words_library_get_type},
	// {"gtk_source_completion_words_library_is_locked", &Gtk_source_completion_words_library_is_locked},
	// {"gtk_source_completion_words_library_lock", &Gtk_source_completion_words_library_lock},
	// {"gtk_source_completion_words_library_new", &Gtk_source_completion_words_library_new},
	// {"gtk_source_completion_words_library_remove_word", &Gtk_source_completion_words_library_remove_word},
	// {"gtk_source_completion_words_library_unlock", &Gtk_source_completion_words_library_unlock},
	{"gtk_source_completion_words_new", &Gtk_source_completion_words_new},
	// {"gtk_source_completion_words_proposal_get_type", &Gtk_source_completion_words_proposal_get_type},
	// {"gtk_source_completion_words_proposal_get_word", &Gtk_source_completion_words_proposal_get_word},
	// {"gtk_source_completion_words_proposal_new", &Gtk_source_completion_words_proposal_new},
	// {"gtk_source_completion_words_proposal_unuse", &Gtk_source_completion_words_proposal_unuse},
	// {"gtk_source_completion_words_proposal_use", &Gtk_source_completion_words_proposal_use},
	{"gtk_source_completion_words_register", &Gtk_source_completion_words_register},
	{"gtk_source_completion_words_unregister", &Gtk_source_completion_words_unregister},
	// {"gtk_source_completion_words_utils_backward_word_start", &Gtk_source_completion_words_utils_backward_word_start},
	// {"gtk_source_completion_words_utils_forward_word_end", &Gtk_source_completion_words_utils_forward_word_end},
	// {"gtk_source_context_class_free", &Gtk_source_context_class_free},
	// {"gtk_source_context_class_new", &Gtk_source_context_class_new},
	{"gtk_source_draw_spaces_flags_get_type", &Gtk_source_draw_spaces_flags_get_type},
	{"gtk_source_gutter_get_type", &Gtk_source_gutter_get_type},
	{"gtk_source_gutter_get_window", &Gtk_source_gutter_get_window},
	{"gtk_source_gutter_insert", &Gtk_source_gutter_insert},
	// {"gtk_source_gutter_new", &Gtk_source_gutter_new},
	{"gtk_source_gutter_queue_draw", &Gtk_source_gutter_queue_draw},
	{"gtk_source_gutter_remove", &Gtk_source_gutter_remove},
	{"gtk_source_gutter_reorder", &Gtk_source_gutter_reorder},
	{"gtk_source_gutter_set_cell_data_func", &Gtk_source_gutter_set_cell_data_func},
	{"gtk_source_gutter_set_cell_size_func", &Gtk_source_gutter_set_cell_size_func},
	{"gtk_source_iter_backward_search", &Gtk_source_iter_backward_search},
	{"gtk_source_iter_forward_search", &Gtk_source_iter_forward_search},
	{"gtk_source_language_get_globs", &Gtk_source_language_get_globs},
	{"gtk_source_language_get_hidden", &Gtk_source_language_get_hidden},
	{"gtk_source_language_get_id", &Gtk_source_language_get_id},
	{"gtk_source_language_get_metadata", &Gtk_source_language_get_metadata},
	{"gtk_source_language_get_mime_types", &Gtk_source_language_get_mime_types},
	{"gtk_source_language_get_name", &Gtk_source_language_get_name},
	{"gtk_source_language_get_section", &Gtk_source_language_get_section},
	{"gtk_source_language_get_style_ids", &Gtk_source_language_get_style_ids},
	{"gtk_source_language_get_style_name", &Gtk_source_language_get_style_name},
	{"gtk_source_language_get_type", &Gtk_source_language_get_type},
	{"gtk_source_language_manager_get_default", &Gtk_source_language_manager_get_default},
	{"gtk_source_language_manager_get_language", &Gtk_source_language_manager_get_language},
	{"gtk_source_language_manager_get_language_ids", &Gtk_source_language_manager_get_language_ids},
	{"gtk_source_language_manager_get_search_path", &Gtk_source_language_manager_get_search_path},
	{"gtk_source_language_manager_get_type", &Gtk_source_language_manager_get_type},
	{"gtk_source_language_manager_guess_language", &Gtk_source_language_manager_guess_language},
	{"gtk_source_language_manager_new", &Gtk_source_language_manager_new},
	{"gtk_source_language_manager_set_search_path", &Gtk_source_language_manager_set_search_path},
	{"gtk_source_mark_get_category", &Gtk_source_mark_get_category},
	{"gtk_source_mark_get_type", &Gtk_source_mark_get_type},
	{"gtk_source_mark_new", &Gtk_source_mark_new},
	{"gtk_source_mark_next", &Gtk_source_mark_next},
	{"gtk_source_mark_prev", &Gtk_source_mark_prev},
	{"gtk_source_print_compositor_draw_page", &Gtk_source_print_compositor_draw_page},
	{"gtk_source_print_compositor_get_body_font_name", &Gtk_source_print_compositor_get_body_font_name},
	{"gtk_source_print_compositor_get_bottom_margin", &Gtk_source_print_compositor_get_bottom_margin},
	{"gtk_source_print_compositor_get_buffer", &Gtk_source_print_compositor_get_buffer},
	{"gtk_source_print_compositor_get_footer_font_name", &Gtk_source_print_compositor_get_footer_font_name},
	{"gtk_source_print_compositor_get_header_font_name", &Gtk_source_print_compositor_get_header_font_name},
	{"gtk_source_print_compositor_get_highlight_syntax", &Gtk_source_print_compositor_get_highlight_syntax},
	{"gtk_source_print_compositor_get_left_margin", &Gtk_source_print_compositor_get_left_margin},
	{"gtk_source_print_compositor_get_line_numbers_font_name", &Gtk_source_print_compositor_get_line_numbers_font_name},
	{"gtk_source_print_compositor_get_n_pages", &Gtk_source_print_compositor_get_n_pages},
	{"gtk_source_print_compositor_get_pagination_progress", &Gtk_source_print_compositor_get_pagination_progress},
	{"gtk_source_print_compositor_get_print_footer", &Gtk_source_print_compositor_get_print_footer},
	{"gtk_source_print_compositor_get_print_header", &Gtk_source_print_compositor_get_print_header},
	{"gtk_source_print_compositor_get_print_line_numbers", &Gtk_source_print_compositor_get_print_line_numbers},
	{"gtk_source_print_compositor_get_right_margin", &Gtk_source_print_compositor_get_right_margin},
	{"gtk_source_print_compositor_get_tab_width", &Gtk_source_print_compositor_get_tab_width},
	{"gtk_source_print_compositor_get_top_margin", &Gtk_source_print_compositor_get_top_margin},
	{"gtk_source_print_compositor_get_type", &Gtk_source_print_compositor_get_type},
	{"gtk_source_print_compositor_get_wrap_mode", &Gtk_source_print_compositor_get_wrap_mode},
	{"gtk_source_print_compositor_new", &Gtk_source_print_compositor_new},
	{"gtk_source_print_compositor_new_from_view", &Gtk_source_print_compositor_new_from_view},
	{"gtk_source_print_compositor_paginate", &Gtk_source_print_compositor_paginate},
	{"gtk_source_print_compositor_set_body_font_name", &Gtk_source_print_compositor_set_body_font_name},
	{"gtk_source_print_compositor_set_bottom_margin", &Gtk_source_print_compositor_set_bottom_margin},
	{"gtk_source_print_compositor_set_footer_font_name", &Gtk_source_print_compositor_set_footer_font_name},
	{"gtk_source_print_compositor_set_footer_format", &Gtk_source_print_compositor_set_footer_format},
	{"gtk_source_print_compositor_set_header_font_name", &Gtk_source_print_compositor_set_header_font_name},
	{"gtk_source_print_compositor_set_header_format", &Gtk_source_print_compositor_set_header_format},
	{"gtk_source_print_compositor_set_highlight_syntax", &Gtk_source_print_compositor_set_highlight_syntax},
	{"gtk_source_print_compositor_set_left_margin", &Gtk_source_print_compositor_set_left_margin},
	{"gtk_source_print_compositor_set_line_numbers_font_name", &Gtk_source_print_compositor_set_line_numbers_font_name},
	{"gtk_source_print_compositor_set_print_footer", &Gtk_source_print_compositor_set_print_footer},
	{"gtk_source_print_compositor_set_print_header", &Gtk_source_print_compositor_set_print_header},
	{"gtk_source_print_compositor_set_print_line_numbers", &Gtk_source_print_compositor_set_print_line_numbers},
	{"gtk_source_print_compositor_set_right_margin", &Gtk_source_print_compositor_set_right_margin},
	{"gtk_source_print_compositor_set_tab_width", &Gtk_source_print_compositor_set_tab_width},
	{"gtk_source_print_compositor_set_top_margin", &Gtk_source_print_compositor_set_top_margin},
	{"gtk_source_print_compositor_set_wrap_mode", &Gtk_source_print_compositor_set_wrap_mode},
	{"gtk_source_search_flags_get_type", &Gtk_source_search_flags_get_type},
	{"gtk_source_smart_home_end_type_get_type", &Gtk_source_smart_home_end_type_get_type},
	{"gtk_source_style_copy", &Gtk_source_style_copy},
	{"gtk_source_style_get_type", &Gtk_source_style_get_type},
	{"gtk_source_style_scheme_get_authors", &Gtk_source_style_scheme_get_authors},
	{"gtk_source_style_scheme_get_description", &Gtk_source_style_scheme_get_description},
	{"gtk_source_style_scheme_get_filename", &Gtk_source_style_scheme_get_filename},
	{"gtk_source_style_scheme_get_id", &Gtk_source_style_scheme_get_id},
	{"gtk_source_style_scheme_get_name", &Gtk_source_style_scheme_get_name},
	{"gtk_source_style_scheme_get_style", &Gtk_source_style_scheme_get_style},
	{"gtk_source_style_scheme_get_type", &Gtk_source_style_scheme_get_type},
	{"gtk_source_style_scheme_manager_append_search_path", &Gtk_source_style_scheme_manager_append_search_path},
	{"gtk_source_style_scheme_manager_force_rescan", &Gtk_source_style_scheme_manager_force_rescan},
	{"gtk_source_style_scheme_manager_get_default", &Gtk_source_style_scheme_manager_get_default},
	{"gtk_source_style_scheme_manager_get_scheme", &Gtk_source_style_scheme_manager_get_scheme},
	{"gtk_source_style_scheme_manager_get_scheme_ids", &Gtk_source_style_scheme_manager_get_scheme_ids},
	{"gtk_source_style_scheme_manager_get_search_path", &Gtk_source_style_scheme_manager_get_search_path},
	{"gtk_source_style_scheme_manager_get_type", &Gtk_source_style_scheme_manager_get_type},
	{"gtk_source_style_scheme_manager_new", &Gtk_source_style_scheme_manager_new},
	{"gtk_source_style_scheme_manager_prepend_search_path", &Gtk_source_style_scheme_manager_prepend_search_path},
	{"gtk_source_style_scheme_manager_set_search_path", &Gtk_source_style_scheme_manager_set_search_path},
	{"gtk_source_undo_manager_begin_not_undoable_action", &Gtk_source_undo_manager_begin_not_undoable_action},
	{"gtk_source_undo_manager_can_redo", &Gtk_source_undo_manager_can_redo},
	{"gtk_source_undo_manager_can_redo_changed", &Gtk_source_undo_manager_can_redo_changed},
	{"gtk_source_undo_manager_can_undo", &Gtk_source_undo_manager_can_undo},
	{"gtk_source_undo_manager_can_undo_changed", &Gtk_source_undo_manager_can_undo_changed},
	// {"gtk_source_undo_manager_default_get_type", &Gtk_source_undo_manager_default_get_type},
	// {"gtk_source_undo_manager_default_set_max_undo_levels", &Gtk_source_undo_manager_default_set_max_undo_levels},
	{"gtk_source_undo_manager_end_not_undoable_action", &Gtk_source_undo_manager_end_not_undoable_action},
	{"gtk_source_undo_manager_get_type", &Gtk_source_undo_manager_get_type},
	{"gtk_source_undo_manager_redo", &Gtk_source_undo_manager_redo},
	{"gtk_source_undo_manager_undo", &Gtk_source_undo_manager_undo},
	{"gtk_source_view_get_auto_indent", &Gtk_source_view_get_auto_indent},
	{"gtk_source_view_get_completion", &Gtk_source_view_get_completion},
	{"gtk_source_view_get_draw_spaces", &Gtk_source_view_get_draw_spaces},
	{"gtk_source_view_get_gutter", &Gtk_source_view_get_gutter},
	{"gtk_source_view_get_highlight_current_line", &Gtk_source_view_get_highlight_current_line},
	{"gtk_source_view_get_indent_on_tab", &Gtk_source_view_get_indent_on_tab},
	{"gtk_source_view_get_indent_width", &Gtk_source_view_get_indent_width},
	{"gtk_source_view_get_insert_spaces_instead_of_tabs", &Gtk_source_view_get_insert_spaces_instead_of_tabs},
	{"gtk_source_view_get_mark_category_background", &Gtk_source_view_get_mark_category_background},
	{"gtk_source_view_get_mark_category_pixbuf", &Gtk_source_view_get_mark_category_pixbuf},
	{"gtk_source_view_get_mark_category_priority", &Gtk_source_view_get_mark_category_priority},
	{"gtk_source_view_get_right_margin_position", &Gtk_source_view_get_right_margin_position},
	{"gtk_source_view_get_show_line_marks", &Gtk_source_view_get_show_line_marks},
	{"gtk_source_view_get_show_line_numbers", &Gtk_source_view_get_show_line_numbers},
	{"gtk_source_view_get_show_right_margin", &Gtk_source_view_get_show_right_margin},
	{"gtk_source_view_get_smart_home_end", &Gtk_source_view_get_smart_home_end},
	{"gtk_source_view_get_tab_width", &Gtk_source_view_get_tab_width},
	{"gtk_source_view_get_type", &Gtk_source_view_get_type},
	{"gtk_source_view_gutter_position_get_type", &Gtk_source_view_gutter_position_get_type},
	{"gtk_source_view_new", &Gtk_source_view_new},
	{"gtk_source_view_new_with_buffer", &Gtk_source_view_new_with_buffer},
	{"gtk_source_view_set_auto_indent", &Gtk_source_view_set_auto_indent},
	{"gtk_source_view_set_draw_spaces", &Gtk_source_view_set_draw_spaces},
	{"gtk_source_view_set_highlight_current_line", &Gtk_source_view_set_highlight_current_line},
	{"gtk_source_view_set_indent_on_tab", &Gtk_source_view_set_indent_on_tab},
	{"gtk_source_view_set_indent_width", &Gtk_source_view_set_indent_width},
	{"gtk_source_view_set_insert_spaces_instead_of_tabs", &Gtk_source_view_set_insert_spaces_instead_of_tabs},
	{"gtk_source_view_set_mark_category_background", &Gtk_source_view_set_mark_category_background},
	{"gtk_source_view_set_mark_category_icon_from_icon_name", &Gtk_source_view_set_mark_category_icon_from_icon_name},
	{"gtk_source_view_set_mark_category_icon_from_pixbuf", &Gtk_source_view_set_mark_category_icon_from_pixbuf},
	{"gtk_source_view_set_mark_category_icon_from_stock", &Gtk_source_view_set_mark_category_icon_from_stock},
	{"gtk_source_view_set_mark_category_pixbuf", &Gtk_source_view_set_mark_category_pixbuf},
	{"gtk_source_view_set_mark_category_priority", &Gtk_source_view_set_mark_category_priority},
	{"gtk_source_view_set_mark_category_tooltip_func", &Gtk_source_view_set_mark_category_tooltip_func},
	{"gtk_source_view_set_mark_category_tooltip_markup_func", &Gtk_source_view_set_mark_category_tooltip_markup_func},
	{"gtk_source_view_set_right_margin_position", &Gtk_source_view_set_right_margin_position},
	{"gtk_source_view_set_show_line_marks", &Gtk_source_view_set_show_line_marks},
	{"gtk_source_view_set_show_line_numbers", &Gtk_source_view_set_show_line_numbers},
	{"gtk_source_view_set_show_right_margin", &Gtk_source_view_set_show_right_margin},
	{"gtk_source_view_set_smart_home_end", &Gtk_source_view_set_smart_home_end},
	{"gtk_source_view_set_tab_width", &Gtk_source_view_set_tab_width},
}
