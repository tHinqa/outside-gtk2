// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gtk provides API definitions for accessing
//libgtk-win32-2.0-0.dll.
package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

var (
	Gtk_accel_group_get_type func() T.GType

	Gtk_accel_group_new func() *T.GtkAccelGroup

	Gtk_accel_group_get_is_locked func(
		accel_group *T.GtkAccelGroup) T.Gboolean

	Gtk_accel_group_get_modifier_mask func(
		accel_group *T.GtkAccelGroup) T.GdkModifierType

	Gtk_accel_group_lock func(accel_group *T.GtkAccelGroup)

	Gtk_accel_group_unlock func(accel_group *T.GtkAccelGroup)

	Gtk_accel_group_connect func(
		accel_group *T.GtkAccelGroup,
		accel_key T.Guint,
		accel_mods T.GdkModifierType,
		accel_flags T.GtkAccelFlags,
		closure *T.GClosure)

	Gtk_accel_group_connect_by_path func(
		accel_group *T.GtkAccelGroup,
		accel_path string,
		closure *T.GClosure)

	Gtk_accel_group_disconnect func(
		accel_group *T.GtkAccelGroup,
		closure *T.GClosure) T.Gboolean

	Gtk_accel_group_disconnect_key func(
		accel_group *T.GtkAccelGroup,
		accel_key T.Guint,
		accel_mods T.GdkModifierType) T.Gboolean

	Gtk_accel_group_activate func(
		accel_group *T.GtkAccelGroup,
		accel_quark T.GQuark,
		acceleratable *T.GObject,
		accel_key T.Guint,
		accel_mods T.GdkModifierType) T.Gboolean

	Gtk_accel_groups_activate func(
		object *T.GObject,
		accel_key T.Guint,
		accel_mods T.GdkModifierType) T.Gboolean

	Gtk_accel_groups_from_object func(
		object *T.GObject) *T.GSList

	Gtk_accel_group_find func(
		accel_group *T.GtkAccelGroup,
		find_func T.GtkAccelGroupFindFunc,
		data T.Gpointer) *T.GtkAccelKey

	Gtk_accel_group_from_accel_closure func(
		closure *T.GClosure) *T.GtkAccelGroup

	Gtk_accelerator_valid func(
		keyval T.Guint,
		modifiers T.GdkModifierType) T.Gboolean

	Gtk_accelerator_parse func(
		accelerator string,
		accelerator_key *T.Guint,
		accelerator_mods *T.GdkModifierType)

	Gtk_accelerator_name func(
		accelerator_key T.Guint,
		accelerator_mods T.GdkModifierType) string

	Gtk_accelerator_get_label func(
		accelerator_key T.Guint,
		accelerator_mods T.GdkModifierType) string

	Gtk_accelerator_set_default_mod_mask func(
		default_mod_mask T.GdkModifierType)

	Gtk_accelerator_get_default_mod_mask func() T.Guint

	Gtk_accel_group_query func(
		accel_group *T.GtkAccelGroup,
		accel_key T.Guint,
		accel_mods T.GdkModifierType,
		n_entries *T.Guint) *T.GtkAccelGroupEntry

	Gtk_accel_flags_get_type func() T.GType

	Gtk_assistant_page_type_get_type func() T.GType

	Gtk_builder_error_get_type func() T.GType

	Gtk_calendar_display_options_get_type func() T.GType

	Gtk_cell_renderer_state_get_type func() T.GType

	Gtk_cell_renderer_mode_get_type func() T.GType

	Gtk_cell_renderer_accel_mode_get_type func() T.GType

	Gtk_debug_flag_get_type func() T.GType

	Gtk_dialog_flags_get_type func() T.GType

	Gtk_response_type_get_type func() T.GType

	Gtk_dest_defaults_get_type func() T.GType

	Gtk_target_flags_get_type func() T.GType

	Gtk_entry_icon_position_get_type func() T.GType

	Gtk_anchor_type_get_type func() T.GType

	Gtk_arrow_placement_get_type func() T.GType

	Gtk_arrow_type_get_type func() T.GType

	Gtk_attach_options_get_type func() T.GType

	Gtk_button_box_style_get_type func() T.GType

	Gtk_curve_type_get_type func() T.GType

	Gtk_delete_type_get_type func() T.GType

	Gtk_direction_type_get_type func() T.GType

	Gtk_expander_style_get_type func() T.GType

	Gtk_icon_size_get_type func() T.GType

	Gtk_sensitivity_type_get_type func() T.GType

	Gtk_side_type_get_type func() T.GType

	Gtk_text_direction_get_type func() T.GType

	Gtk_justification_get_type func() T.GType

	Gtk_match_type_get_type func() T.GType

	Gtk_menu_direction_type_get_type func() T.GType

	Gtk_message_type_get_type func() T.GType

	Gtk_metric_type_get_type func() T.GType

	Gtk_movement_step_get_type func() T.GType

	Gtk_scroll_step_get_type func() T.GType

	Gtk_orientation_get_type func() T.GType

	Gtk_corner_type_get_type func() T.GType

	Gtk_pack_type_get_type func() T.GType

	Gtk_path_priority_type_get_type func() T.GType

	Gtk_path_type_get_type func() T.GType

	Gtk_policy_type_get_type func() T.GType

	Gtk_position_type_get_type func() T.GType

	Gtk_preview_type_get_type func() T.GType

	Gtk_relief_style_get_type func() T.GType

	Gtk_resize_mode_get_type func() T.GType

	Gtk_signal_run_type_get_type func() T.GType

	Gtk_scroll_type_get_type func() T.GType

	Gtk_selection_mode_get_type func() T.GType

	Gtk_shadow_type_get_type func() T.GType

	Gtk_state_type_get_type func() T.GType

	Gtk_submenu_direction_get_type func() T.GType

	Gtk_submenu_placement_get_type func() T.GType

	Gtk_toolbar_style_get_type func() T.GType

	Gtk_update_type_get_type func() T.GType

	Gtk_visibility_get_type func() T.GType

	Gtk_window_position_get_type func() T.GType

	Gtk_window_type_get_type func() T.GType

	Gtk_wrap_mode_get_type func() T.GType

	Gtk_sort_type_get_type func() T.GType

	Gtk_im_preedit_style_get_type func() T.GType

	Gtk_im_status_style_get_type func() T.GType

	Gtk_pack_direction_get_type func() T.GType

	Gtk_print_pages_get_type func() T.GType

	Gtk_page_set_get_type func() T.GType

	Gtk_number_up_layout_get_type func() T.GType

	Gtk_page_orientation_get_type func() T.GType

	Gtk_print_quality_get_type func() T.GType

	Gtk_print_duplex_get_type func() T.GType

	Gtk_unit_get_type func() T.GType

	Gtk_tree_view_grid_lines_get_type func() T.GType

	Gtk_drag_result_get_type func() T.GType

	Gtk_file_chooser_action_get_type func() T.GType

	Gtk_file_chooser_confirmation_get_type func() T.GType

	Gtk_file_chooser_error_get_type func() T.GType

	Gtk_file_filter_flags_get_type func() T.GType

	Gtk_icon_lookup_flags_get_type func() T.GType

	Gtk_icon_theme_error_get_type func() T.GType

	Gtk_icon_view_drop_position_get_type func() T.GType

	Gtk_image_type_get_type func() T.GType

	Gtk_buttons_type_get_type func() T.GType

	Gtk_notebook_tab_get_type func() T.GType

	Gtk_object_flags_get_type func() T.GType

	Gtk_arg_flags_get_type func() T.GType

	Gtk_print_status_get_type func() T.GType

	Gtk_print_operation_result_get_type func() T.GType

	Gtk_print_operation_action_get_type func() T.GType

	Gtk_print_error_get_type func() T.GType

	Gtk_private_flags_get_type func() T.GType

	Gtk_progress_bar_style_get_type func() T.GType

	Gtk_progress_bar_orientation_get_type func() T.GType

	Gtk_rc_flags_get_type func() T.GType

	Gtk_rc_token_type_get_type func() T.GType

	Gtk_recent_sort_type_get_type func() T.GType

	Gtk_recent_chooser_error_get_type func() T.GType

	Gtk_recent_filter_flags_get_type func() T.GType

	Gtk_recent_manager_error_get_type func() T.GType

	Gtk_size_group_mode_get_type func() T.GType

	Gtk_spin_button_update_policy_get_type func() T.GType

	Gtk_spin_type_get_type func() T.GType

	Gtk_text_buffer_target_info_get_type func() T.GType

	Gtk_text_search_flags_get_type func() T.GType

	Gtk_text_window_type_get_type func() T.GType

	Gtk_toolbar_child_type_get_type func() T.GType

	Gtk_toolbar_space_style_get_type func() T.GType

	Gtk_tool_palette_drag_targets_get_type func() T.GType

	Gtk_tree_model_flags_get_type func() T.GType

	Gtk_tree_view_drop_position_get_type func() T.GType

	Gtk_tree_view_column_sizing_get_type func() T.GType

	Gtk_ui_manager_item_type_get_type func() T.GType

	Gtk_widget_flags_get_type func() T.GType

	Gtk_widget_help_type_get_type func() T.GType

	Gtk_tree_view_mode_get_type func() T.GType

	Gtk_cell_type_get_type func() T.GType

	Gtk_clist_drag_pos_get_type func() T.GType

	Gtk_button_action_get_type func() T.GType

	Gtk_ctree_pos_get_type func() T.GType

	Gtk_ctree_line_style_get_type func() T.GType

	Gtk_ctree_expander_style_get_type func() T.GType

	Gtk_ctree_expansion_type_get_type func() T.GType

	Gtk_identifier_get_type func() T.GType

	Gtk_type_init func(
		debug_flags T.GTypeDebugFlags)

	Gtk_type_unique func(
		parent_type T.GtkType,
		gtkinfo *T.GtkTypeInfo) T.GtkType

	Gtk_type_class func(
		t T.GtkType) T.Gpointer

	Gtk_type_new func(
		t T.GtkType) T.Gpointer

	Gtk_object_get_type func() T.GType

	Gtk_object_sink func(
		object *T.GtkObject)

	Gtk_object_destroy func(
		object *T.GtkObject)

	Gtk_object_new func(t T.GType, first_property_name string,
		v ...VArg) *T.GtkObject

	Gtk_object_ref func(
		object *T.GtkObject) *T.GtkObject

	Gtk_object_unref func(
		object *T.GtkObject)

	Gtk_object_weakref func(
		object *T.GtkObject,
		notify T.GDestroyNotify,
		data T.Gpointer)

	Gtk_object_weakunref func(
		object *T.GtkObject,
		notify T.GDestroyNotify,
		data T.Gpointer)

	Gtk_object_set_data func(
		object *T.GtkObject,
		key string,
		data T.Gpointer)

	Gtk_object_set_data_full func(
		object *T.GtkObject,
		key string,
		data T.Gpointer,
		destroy T.GDestroyNotify)

	Gtk_object_remove_data func(
		object *T.GtkObject,
		key string)

	Gtk_object_get_data func(
		object *T.GtkObject,
		key string) T.Gpointer

	Gtk_object_remove_no_notify func(
		object *T.GtkObject,
		key string)

	Gtk_object_set_user_data func(
		object *T.GtkObject,
		data T.Gpointer)

	Gtk_object_get_user_data func(
		object *T.GtkObject) T.Gpointer

	Gtk_object_set_data_by_id func(
		object *T.GtkObject,
		data_id T.GQuark,
		data T.Gpointer)

	Gtk_object_set_data_by_id_full func(
		object *T.GtkObject,
		data_id T.GQuark,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_object_get_data_by_id func(
		object *T.GtkObject,
		data_id T.GQuark) T.Gpointer

	Gtk_object_remove_data_by_id func(
		object *T.GtkObject,
		data_id T.GQuark)

	Gtk_object_remove_no_notify_by_id func(
		object *T.GtkObject,
		key_id T.GQuark)

	Gtk_object_get func(object *T.GtkObject,
		first_property_name string, v ...VArg)

	Gtk_object_set func(object *T.GtkObject,
		first_property_name string, v ...VArg)

	Gtk_object_add_arg_type func(
		arg_name string,
		arg_type T.GType,
		arg_flags T.Guint,
		arg_id T.Guint)

	Gtk_adjustment_get_type func() T.GType

	Gtk_adjustment_new func(
		value T.Gdouble,
		lower T.Gdouble,
		upper T.Gdouble,
		step_increment T.Gdouble,
		page_increment T.Gdouble,
		page_size T.Gdouble) *T.GtkObject

	Gtk_adjustment_changed func(
		adjustment *T.GtkAdjustment)

	Gtk_adjustment_value_changed func(
		adjustment *T.GtkAdjustment)

	Gtk_adjustment_clamp_page func(
		adjustment *T.GtkAdjustment, lower, upper T.Gdouble)

	Gtk_adjustment_get_value func(
		adjustment *T.GtkAdjustment) T.Gdouble

	Gtk_adjustment_set_value func(
		adjustment *T.GtkAdjustment, value T.Gdouble)

	Gtk_adjustment_get_lower func(
		adjustment *T.GtkAdjustment) T.Gdouble

	Gtk_adjustment_set_lower func(
		adjustment *T.GtkAdjustment, lower T.Gdouble)

	Gtk_adjustment_get_upper func(
		adjustment *T.GtkAdjustment) T.Gdouble

	Gtk_adjustment_set_upper func(
		adjustment *T.GtkAdjustment, upper T.Gdouble)

	Gtk_adjustment_get_step_increment func(
		adjustment *T.GtkAdjustment) T.Gdouble

	Gtk_adjustment_set_step_increment func(
		adjustment *T.GtkAdjustment,
		step_increment T.Gdouble)

	Gtk_adjustment_get_page_increment func(
		adjustment *T.GtkAdjustment) T.Gdouble

	Gtk_adjustment_set_page_increment func(
		adjustment *T.GtkAdjustment,
		page_increment T.Gdouble)

	Gtk_adjustment_get_page_size func(
		adjustment *T.GtkAdjustment) T.Gdouble

	Gtk_adjustment_set_page_size func(
		adjustment *T.GtkAdjustment,
		page_size T.Gdouble)

	Gtk_adjustment_configure func(
		adjustment *T.GtkAdjustment,
		value T.Gdouble,
		lower T.Gdouble,
		upper T.Gdouble,
		step_increment T.Gdouble,
		page_increment T.Gdouble,
		page_size T.Gdouble)

	Gtk_style_get_type func() T.GType

	Gtk_style_new func() *T.GtkStyle

	Gtk_style_copy func(
		style *T.GtkStyle) *T.GtkStyle

	Gtk_style_attach func(
		style *T.GtkStyle,
		window *T.GdkWindow) *T.GtkStyle

	Gtk_style_detach func(
		style *T.GtkStyle)

	Gtk_style_ref func(
		style *T.GtkStyle) *T.GtkStyle

	Gtk_style_unref func(
		style *T.GtkStyle)

	Gtk_style_get_font func(
		style *T.GtkStyle) *T.GdkFont

	Gtk_style_set_font func(
		style *T.GtkStyle,
		font *T.GdkFont)

	Gtk_style_set_background func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType)

	Gtk_style_apply_default_background func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		set_bg T.Gboolean,
		state_type T.GtkStateType,
		area *T.GdkRectangle,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_style_lookup_icon_set func(
		style *T.GtkStyle,
		stock_id string) *T.GtkIconSet

	Gtk_style_lookup_color func(
		style *T.GtkStyle,
		color_name string,
		color *T.GdkColor) T.Gboolean

	Gtk_style_render_icon func(
		style *T.GtkStyle,
		source *T.GtkIconSource,
		direction T.GtkTextDirection,
		state T.GtkStateType,
		size T.GtkIconSize,
		widget *T.GtkWidget,
		detail string) *T.GdkPixbuf

	Gtk_draw_hline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		x1 T.Gint,
		x2 T.Gint,
		y T.Gint)

	Gtk_draw_vline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		y1_ T.Gint,
		y2_ T.Gint,
		x T.Gint)

	Gtk_draw_shadow func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_polygon func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		points *T.GdkPoint,
		npoints T.Gint,
		fill T.Gboolean)

	Gtk_draw_arrow func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		arrow_type T.GtkArrowType,
		fill T.Gboolean,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_diamond func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_box func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_flat_box func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_check func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_option func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_tab func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_shadow_gap func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		gap_side T.GtkPositionType,
		gap_x T.Gint,
		gap_width T.Gint)

	Gtk_draw_box_gap func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		gap_side T.GtkPositionType,
		gap_x T.Gint,
		gap_width T.Gint)

	Gtk_draw_extension func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		gap_side T.GtkPositionType)

	Gtk_draw_focus func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_draw_slider func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		orientation T.GtkOrientation)

	Gtk_draw_handle func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		orientation T.GtkOrientation)

	Gtk_draw_expander func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		x T.Gint,
		y T.Gint,
		expander_style T.GtkExpanderStyle)

	Gtk_draw_layout func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		use_text T.Gboolean,
		x T.Gint,
		y T.Gint,
		layout *T.PangoLayout)

	Gtk_draw_resize_grip func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		edge T.GdkWindowEdge,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_hline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x1 T.Gint,
		x2 T.Gint,
		y T.Gint)

	Gtk_paint_vline func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		y1_ T.Gint,
		y2_ T.Gint,
		x T.Gint)

	Gtk_paint_shadow func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_polygon func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		points *T.GdkPoint,
		n_points T.Gint,
		fill T.Gboolean)

	Gtk_paint_arrow func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		arrow_type T.GtkArrowType,
		fill T.Gboolean,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_diamond func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_box func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_flat_box func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_check func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_option func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_tab func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_shadow_gap func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		gap_side T.GtkPositionType,
		gap_x T.Gint,
		gap_width T.Gint)

	Gtk_paint_box_gap func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		gap_side T.GtkPositionType,
		gap_x T.Gint,
		gap_width T.Gint)

	Gtk_paint_extension func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		gap_side T.GtkPositionType)

	Gtk_paint_focus func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_slider func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		orientation T.GtkOrientation)

	Gtk_paint_handle func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		shadow_type T.GtkShadowType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint,
		orientation T.GtkOrientation)

	Gtk_paint_expander func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		expander_style T.GtkExpanderStyle)

	Gtk_paint_layout func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		use_text T.Gboolean,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		layout *T.PangoLayout)

	Gtk_paint_resize_grip func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		edge T.GdkWindowEdge,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_paint_spinner func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		step T.Guint,
		x T.Gint,
		y T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_border_get_type func() T.GType

	Gtk_border_new func() *T.GtkBorder

	Gtk_border_copy func(
		border_ *T.GtkBorder) *T.GtkBorder

	Gtk_border_free func(
		border_ *T.GtkBorder)

	Gtk_style_get_style_property func(
		style *T.GtkStyle,
		widget_type T.GType,
		property_name string,
		value *T.GValue)

	Gtk_style_get_valist func(
		style *T.GtkStyle,
		widget_type T.GType,
		first_property_name string,
		var_args T.Va_list)

	Gtk_style_get func(style *T.GtkStyle, widget_type T.GType,
		first_property_name string, v ...VArg)

	Gtk_draw_string func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		x T.Gint,
		y T.Gint,
		string string)

	Gtk_paint_string func(
		style *T.GtkStyle,
		window *T.GdkWindow,
		state_type T.GtkStateType,
		area *T.GdkRectangle,
		widget *T.GtkWidget,
		detail string,
		x T.Gint,
		y T.Gint,
		string string)

	Gtk_draw_insertion_cursor func(
		widget *T.GtkWidget,
		drawable *T.GdkDrawable,
		area *T.GdkRectangle,
		location *T.GdkRectangle,
		is_primary T.Gboolean,
		direction T.GtkTextDirection,
		draw_arrow T.Gboolean)

	Gtk_rc_add_default_file func(
		filename string)

	Gtk_rc_set_default_files func(
		filenames **T.Gchar)

	Gtk_rc_get_default_files func() **T.Gchar

	Gtk_rc_get_style func(
		widget *T.GtkWidget) *T.GtkStyle

	Gtk_rc_get_style_by_paths func(
		settings *T.GtkSettings,
		widget_path string,
		class_path string,
		t T.GType) *T.GtkStyle

	Gtk_rc_reparse_all_for_settings func(
		settings *T.GtkSettings,
		force_load T.Gboolean) T.Gboolean

	Gtk_rc_reset_styles func(
		settings *T.GtkSettings)

	Gtk_rc_find_pixmap_in_path func(
		settings *T.GtkSettings,
		scanner *T.GScanner,
		pixmap_file string) string

	Gtk_rc_parse func(
		filename string)

	Gtk_rc_parse_string func(
		rc_string string)

	Gtk_rc_reparse_all func() T.Gboolean

	Gtk_rc_add_widget_name_style func(
		rc_style *T.GtkRcStyle,
		pattern string)

	Gtk_rc_add_widget_class_style func(
		rc_style *T.GtkRcStyle,
		pattern string)

	Gtk_rc_add_class_style func(
		rc_style *T.GtkRcStyle,
		pattern string)

	Gtk_rc_style_get_type func() T.GType

	Gtk_rc_style_new func() *T.GtkRcStyle

	Gtk_rc_style_copy func(
		orig *T.GtkRcStyle) *T.GtkRcStyle

	Gtk_rc_style_ref func(
		rc_style *T.GtkRcStyle)

	Gtk_rc_style_unref func(
		rc_style *T.GtkRcStyle)

	Gtk_rc_find_module_in_path func(
		module_file string) string

	Gtk_rc_get_theme_dir func() string

	Gtk_rc_get_module_dir func() string

	Gtk_rc_get_im_module_path func() string

	Gtk_rc_get_im_module_file func() string

	Gtk_rc_scanner_new func() *T.GScanner

	Gtk_rc_parse_color func(
		scanner *T.GScanner,
		color *T.GdkColor) T.Guint

	Gtk_rc_parse_color_full func(
		scanner *T.GScanner,
		style *T.GtkRcStyle,
		color *T.GdkColor) T.Guint

	Gtk_rc_parse_state func(
		scanner *T.GScanner,
		state *T.GtkStateType) T.Guint

	Gtk_rc_parse_priority func(
		scanner *T.GScanner,
		priority *T.GtkPathPriorityType) T.Guint

	Gtk_settings_get_type func() T.GType

	Gtk_settings_get_default func() *T.GtkSettings

	Gtk_settings_get_for_screen func(
		screen *T.GdkScreen) *T.GtkSettings

	Gtk_settings_install_property func(
		pspec *T.GParamSpec)

	Gtk_settings_install_property_parser func(
		pspec *T.GParamSpec,
		parser T.GtkRcPropertyParser)

	Gtk_rc_property_parse_color func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		property_value *T.GValue) T.Gboolean

	Gtk_rc_property_parse_enum func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		property_value *T.GValue) T.Gboolean

	Gtk_rc_property_parse_flags func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		property_value *T.GValue) T.Gboolean

	Gtk_rc_property_parse_requisition func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		property_value *T.GValue) T.Gboolean

	Gtk_rc_property_parse_border func(
		pspec *T.GParamSpec,
		gstring *T.GString,
		property_value *T.GValue) T.Gboolean

	Gtk_settings_set_property_value func(
		settings *T.GtkSettings,
		name string,
		svalue *T.GtkSettingsValue)

	Gtk_settings_set_string_property func(
		settings *T.GtkSettings,
		name string,
		v_string string,
		origin string)

	Gtk_settings_set_long_property func(
		settings *T.GtkSettings,
		name string,
		v_long T.Glong,
		origin string)

	Gtk_settings_set_double_property func(
		settings *T.GtkSettings,
		name string,
		v_double T.Gdouble,
		origin string)

	Gtk_widget_get_type func() T.GType

	Gtk_widget_new func(t T.GType, first_property_name string,
		v ...VArg) *T.GtkWidget

	Gtk_widget_destroy func(
		widget *T.GtkWidget)

	Gtk_widget_destroyed func(
		widget *T.GtkWidget,
		widget_pointer **T.GtkWidget)

	Gtk_widget_ref func(
		widget *T.GtkWidget) *T.GtkWidget

	Gtk_widget_unref func(
		widget *T.GtkWidget)

	Gtk_widget_set func(widget *T.GtkWidget,
		first_property_name string, v ...VArg)

	Gtk_widget_hide_all func(
		widget *T.GtkWidget)

	Gtk_widget_unparent func(
		widget *T.GtkWidget)

	Gtk_widget_show func(
		widget *T.GtkWidget)

	Gtk_widget_show_now func(
		widget *T.GtkWidget)

	Gtk_widget_hide func(
		widget *T.GtkWidget)

	Gtk_widget_show_all func(
		widget *T.GtkWidget)

	Gtk_widget_set_no_show_all func(
		widget *T.GtkWidget,
		no_show_all T.Gboolean)

	Gtk_widget_get_no_show_all func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_map func(
		widget *T.GtkWidget)

	Gtk_widget_unmap func(
		widget *T.GtkWidget)

	Gtk_widget_realize func(
		widget *T.GtkWidget)

	Gtk_widget_unrealize func(
		widget *T.GtkWidget)

	Gtk_widget_queue_draw func(
		widget *T.GtkWidget)

	Gtk_widget_queue_draw_area func(
		widget *T.GtkWidget, x, y, width, height T.Gint)

	Gtk_widget_queue_clear func(
		widget *T.GtkWidget)

	Gtk_widget_queue_clear_area func(
		widget *T.GtkWidget, x, y, width, height T.Gint)

	Gtk_widget_queue_resize func(widget *T.GtkWidget)

	Gtk_widget_queue_resize_no_redraw func(widget *T.GtkWidget)

	Gtk_widget_draw func(widget *T.GtkWidget, area *T.GdkRectangle)

	Gtk_widget_size_request func(
		widget *T.GtkWidget,
		requisition *T.GtkRequisition)

	Gtk_widget_size_allocate func(
		widget *T.GtkWidget,
		allocation *T.GtkAllocation)

	Gtk_widget_get_child_requisition func(
		widget *T.GtkWidget,
		requisition *T.GtkRequisition)

	Gtk_widget_add_accelerator func(
		widget *T.GtkWidget,
		accel_signal string,
		accel_group *T.GtkAccelGroup,
		accel_key T.Guint,
		accel_mods T.GdkModifierType,
		accel_flags T.GtkAccelFlags)

	Gtk_widget_remove_accelerator func(
		widget *T.GtkWidget,
		accel_group *T.GtkAccelGroup,
		accel_key T.Guint,
		accel_mods T.GdkModifierType) T.Gboolean

	Gtk_widget_set_accel_path func(
		widget *T.GtkWidget,
		accel_path string,
		accel_group *T.GtkAccelGroup)

	Gtk_widget_list_accel_closures func(
		widget *T.GtkWidget) *T.GList

	Gtk_widget_can_activate_accel func(
		widget *T.GtkWidget,
		signal_id T.Guint) T.Gboolean

	Gtk_widget_mnemonic_activate func(
		widget *T.GtkWidget,
		group_cycling T.Gboolean) T.Gboolean

	Gtk_widget_event func(
		widget *T.GtkWidget,
		event *T.GdkEvent) T.Gboolean

	Gtk_widget_send_expose func(
		widget *T.GtkWidget,
		event *T.GdkEvent) T.Gint

	Gtk_widget_send_focus_change func(
		widget *T.GtkWidget,
		event *T.GdkEvent) T.Gboolean

	Gtk_widget_activate func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_scroll_adjustments func(
		widget *T.GtkWidget,
		hadjustment *T.GtkAdjustment,
		vadjustment *T.GtkAdjustment) T.Gboolean

	Gtk_widget_reparent func(
		widget *T.GtkWidget,
		new_parent *T.GtkWidget)

	Gtk_widget_intersect func(
		widget *T.GtkWidget,
		area *T.GdkRectangle,
		intersection *T.GdkRectangle) T.Gboolean

	Gtk_widget_region_intersect func(
		widget *T.GtkWidget,
		region *T.GdkRegion) *T.GdkRegion

	Gtk_widget_freeze_child_notify func(
		widget *T.GtkWidget)

	Gtk_widget_child_notify func(
		widget *T.GtkWidget,
		child_property string)

	Gtk_widget_thaw_child_notify func(
		widget *T.GtkWidget)

	Gtk_widget_set_can_focus func(
		widget *T.GtkWidget,
		can_focus T.Gboolean)

	Gtk_widget_get_can_focus func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_has_focus func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_is_focus func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_grab_focus func(
		widget *T.GtkWidget)

	Gtk_widget_set_can_default func(
		widget *T.GtkWidget,
		can_default T.Gboolean)

	Gtk_widget_get_can_default func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_has_default func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_grab_default func(
		widget *T.GtkWidget)

	Gtk_widget_set_receives_default func(
		widget *T.GtkWidget,
		receives_default T.Gboolean)

	Gtk_widget_get_receives_default func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_has_grab func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_name func(
		widget *T.GtkWidget,
		name string)

	Gtk_widget_get_name func(
		widget *T.GtkWidget) string

	Gtk_widget_set_state func(
		widget *T.GtkWidget,
		state T.GtkStateType)

	Gtk_widget_get_state func(
		widget *T.GtkWidget) T.GtkStateType

	Gtk_widget_set_sensitive func(
		widget *T.GtkWidget,
		sensitive T.Gboolean)

	Gtk_widget_get_sensitive func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_is_sensitive func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_visible func(
		widget *T.GtkWidget,
		visible T.Gboolean)

	Gtk_widget_get_visible func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_has_window func(
		widget *T.GtkWidget,
		has_window T.Gboolean)

	Gtk_widget_get_has_window func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_is_toplevel func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_is_drawable func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_realized func(
		widget *T.GtkWidget,
		realized T.Gboolean)

	Gtk_widget_get_realized func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_mapped func(
		widget *T.GtkWidget,
		mapped T.Gboolean)

	Gtk_widget_get_mapped func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_app_paintable func(
		widget *T.GtkWidget,
		app_paintable T.Gboolean)

	Gtk_widget_get_app_paintable func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_double_buffered func(
		widget *T.GtkWidget,
		double_buffered T.Gboolean)

	Gtk_widget_get_double_buffered func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_redraw_on_allocate func(
		widget *T.GtkWidget,
		redraw_on_allocate T.Gboolean)

	Gtk_widget_set_parent func(
		widget *T.GtkWidget,
		parent *T.GtkWidget)

	Gtk_widget_get_parent func(
		widget *T.GtkWidget) *T.GtkWidget

	Gtk_widget_set_parent_window func(
		widget *T.GtkWidget,
		parent_window *T.GdkWindow)

	Gtk_widget_get_parent_window func(
		widget *T.GtkWidget) *T.GdkWindow

	Gtk_widget_set_child_visible func(
		widget *T.GtkWidget,
		is_visible T.Gboolean)

	Gtk_widget_get_child_visible func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_window func(
		widget *T.GtkWidget,
		window *T.GdkWindow)

	Gtk_widget_get_window func(
		widget *T.GtkWidget) *T.GdkWindow

	Gtk_widget_get_allocation func(
		widget *T.GtkWidget,
		allocation *T.GtkAllocation)

	Gtk_widget_set_allocation func(
		widget *T.GtkWidget,
		allocation *T.GtkAllocation)

	Gtk_widget_get_requisition func(
		widget *T.GtkWidget,
		requisition *T.GtkRequisition)

	Gtk_widget_child_focus func(
		widget *T.GtkWidget,
		direction T.GtkDirectionType) T.Gboolean

	Gtk_widget_keynav_failed func(
		widget *T.GtkWidget,
		direction T.GtkDirectionType) T.Gboolean

	Gtk_widget_error_bell func(
		widget *T.GtkWidget)

	Gtk_widget_set_size_request func(
		widget *T.GtkWidget,
		width T.Gint,
		height T.Gint)

	Gtk_widget_get_size_request func(
		widget *T.GtkWidget,
		width *T.Gint,
		height *T.Gint)

	Gtk_widget_set_uposition func(
		widget *T.GtkWidget,
		x T.Gint,
		y T.Gint)

	Gtk_widget_set_usize func(
		widget *T.GtkWidget,
		width T.Gint,
		height T.Gint)

	Gtk_widget_set_events func(
		widget *T.GtkWidget,
		events T.Gint)

	Gtk_widget_add_events func(
		widget *T.GtkWidget,
		events T.Gint)

	Gtk_widget_set_extension_events func(
		widget *T.GtkWidget,
		mode T.GdkExtensionMode)

	Gtk_widget_get_extension_events func(
		widget *T.GtkWidget) T.GdkExtensionMode

	Gtk_widget_get_toplevel func(
		widget *T.GtkWidget) *T.GtkWidget

	Gtk_widget_get_ancestor func(
		widget *T.GtkWidget,
		widget_type T.GType) *T.GtkWidget

	Gtk_widget_get_colormap func(
		widget *T.GtkWidget) *T.GdkColormap

	Gtk_widget_get_visual func(
		widget *T.GtkWidget) *T.GdkVisual

	Gtk_widget_get_screen func(
		widget *T.GtkWidget) *T.GdkScreen

	Gtk_widget_has_screen func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_get_display func(
		widget *T.GtkWidget) *T.GdkDisplay

	Gtk_widget_get_root_window func(
		widget *T.GtkWidget) *T.GdkWindow

	Gtk_widget_get_settings func(
		widget *T.GtkWidget) *T.GtkSettings

	Gtk_widget_get_clipboard func(
		widget *T.GtkWidget,
		selection T.GdkAtom) *T.GtkClipboard

	Gtk_widget_get_snapshot func(
		widget *T.GtkWidget,
		clip_rect *T.GdkRectangle) *T.GdkPixmap

	Gtk_widget_get_accessible func(
		widget *T.GtkWidget) *T.AtkObject

	Gtk_widget_set_colormap func(
		widget *T.GtkWidget,
		colormap *T.GdkColormap)

	Gtk_widget_get_events func(
		widget *T.GtkWidget) T.Gint

	Gtk_widget_get_pointer func(
		widget *T.GtkWidget,
		x *T.Gint,
		y *T.Gint)

	Gtk_widget_is_ancestor func(
		widget *T.GtkWidget,
		ancestor *T.GtkWidget) T.Gboolean

	Gtk_widget_translate_coordinates func(
		src_widget *T.GtkWidget,
		dest_widget *T.GtkWidget,
		src_x T.Gint,
		src_y T.Gint,
		dest_x *T.Gint,
		dest_y *T.Gint) T.Gboolean

	Gtk_widget_hide_on_delete func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_style_attach func(
		style *T.GtkWidget)

	Gtk_widget_has_rc_style func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_set_style func(
		widget *T.GtkWidget,
		style *T.GtkStyle)

	Gtk_widget_ensure_style func(
		widget *T.GtkWidget)

	Gtk_widget_get_style func(
		widget *T.GtkWidget) *T.GtkStyle

	Gtk_widget_modify_style func(
		widget *T.GtkWidget,
		style *T.GtkRcStyle)

	Gtk_widget_get_modifier_style func(
		widget *T.GtkWidget) *T.GtkRcStyle

	Gtk_widget_modify_fg func(
		widget *T.GtkWidget,
		state T.GtkStateType,
		color *T.GdkColor)

	Gtk_widget_modify_bg func(
		widget *T.GtkWidget,
		state T.GtkStateType,
		color *T.GdkColor)

	Gtk_widget_modify_text func(
		widget *T.GtkWidget,
		state T.GtkStateType,
		color *T.GdkColor)

	Gtk_widget_modify_base func(
		widget *T.GtkWidget,
		state T.GtkStateType,
		color *T.GdkColor)

	Gtk_widget_modify_cursor func(
		widget *T.GtkWidget,
		primary *T.GdkColor,
		secondary *T.GdkColor)

	Gtk_widget_modify_font func(
		widget *T.GtkWidget,
		font_desc *T.PangoFontDescription)

	Gtk_widget_create_pango_context func(
		widget *T.GtkWidget) *T.PangoContext

	Gtk_widget_get_pango_context func(
		widget *T.GtkWidget) *T.PangoContext

	Gtk_widget_create_pango_layout func(
		widget *T.GtkWidget,
		text string) *T.PangoLayout

	Gtk_widget_render_icon func(
		widget *T.GtkWidget,
		stock_id string,
		size T.GtkIconSize,
		detail string) *T.GdkPixbuf

	Gtk_widget_set_composite_name func(
		widget *T.GtkWidget,
		name string)

	Gtk_widget_get_composite_name func(
		widget *T.GtkWidget) string

	Gtk_widget_reset_rc_styles func(
		widget *T.GtkWidget)

	Gtk_widget_push_colormap func(
		cmap *T.GdkColormap)

	Gtk_widget_push_composite_child func()

	Gtk_widget_pop_composite_child func()

	Gtk_widget_pop_colormap func()

	Gtk_widget_class_install_style_property func(
		klass *T.GtkWidgetClass,
		pspec *T.GParamSpec)

	Gtk_widget_class_install_style_property_parser func(
		klass *T.GtkWidgetClass,
		pspec *T.GParamSpec,
		parser T.GtkRcPropertyParser)

	Gtk_widget_class_find_style_property func(
		klass *T.GtkWidgetClass,
		property_name string) *T.GParamSpec

	Gtk_widget_class_list_style_properties func(
		klass *T.GtkWidgetClass,
		n_properties *T.Guint) **T.GParamSpec

	Gtk_widget_style_get_property func(
		widget *T.GtkWidget,
		property_name string,
		value *T.GValue)

	Gtk_widget_style_get_valist func(
		widget *T.GtkWidget,
		first_property_name string,
		var_args T.Va_list)

	Gtk_widget_style_get func(widget *T.GtkWidget,
		first_property_name string, v ...VArg)

	Gtk_widget_set_default_colormap func(
		colormap *T.GdkColormap)

	Gtk_widget_get_default_style func() *T.GtkStyle

	Gtk_widget_get_default_colormap func() *T.GdkColormap

	Gtk_widget_get_default_visual func() *T.GdkVisual

	Gtk_widget_set_direction func(
		widget *T.GtkWidget,
		dir T.GtkTextDirection)

	Gtk_widget_get_direction func(
		widget *T.GtkWidget) T.GtkTextDirection

	Gtk_widget_set_default_direction func(
		dir T.GtkTextDirection)

	Gtk_widget_get_default_direction func() T.GtkTextDirection

	Gtk_widget_is_composited func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_widget_shape_combine_mask func(
		widget *T.GtkWidget,
		shape_mask *T.GdkBitmap,
		offset_x T.Gint,
		offset_y T.Gint)

	Gtk_widget_input_shape_combine_mask func(
		widget *T.GtkWidget,
		shape_mask *T.GdkBitmap,
		offset_x T.Gint,
		offset_y T.Gint)

	Gtk_widget_reset_shapes func(
		widget *T.GtkWidget)

	Gtk_widget_path func(
		widget *T.GtkWidget,
		path_length *T.Guint,
		path **T.Gchar,
		path_reversed **T.Gchar)

	Gtk_widget_class_path func(
		widget *T.GtkWidget,
		path_length *T.Guint,
		path **T.Gchar,
		path_reversed **T.Gchar)

	Gtk_widget_list_mnemonic_labels func(
		widget *T.GtkWidget) *T.GList

	Gtk_widget_add_mnemonic_label func(
		widget *T.GtkWidget,
		label *T.GtkWidget)

	Gtk_widget_remove_mnemonic_label func(
		widget *T.GtkWidget,
		label *T.GtkWidget)

	Gtk_widget_set_tooltip_window func(
		widget *T.GtkWidget,
		custom_window *T.GtkWindow)

	Gtk_widget_get_tooltip_window func(
		widget *T.GtkWidget) *T.GtkWindow

	Gtk_widget_trigger_tooltip_query func(
		widget *T.GtkWidget)

	Gtk_widget_set_tooltip_text func(
		widget *T.GtkWidget,
		text string)

	Gtk_widget_get_tooltip_text func(
		widget *T.GtkWidget) string

	Gtk_widget_set_tooltip_markup func(
		widget *T.GtkWidget,
		markup string)

	Gtk_widget_get_tooltip_markup func(
		widget *T.GtkWidget) string

	Gtk_widget_set_has_tooltip func(
		widget *T.GtkWidget,
		has_tooltip T.Gboolean)

	Gtk_widget_get_has_tooltip func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_requisition_get_type func() T.GType

	Gtk_requisition_copy func(
		requisition *T.GtkRequisition) *T.GtkRequisition

	Gtk_requisition_free func(
		requisition *T.GtkRequisition)

	Gtk_container_get_type func() T.GType

	Gtk_container_set_border_width func(
		container *T.GtkContainer,
		border_width T.Guint)

	Gtk_container_get_border_width func(
		container *T.GtkContainer) T.Guint

	Gtk_container_add func(
		container *T.GtkContainer,
		widget *T.GtkWidget)

	Gtk_container_remove func(
		container *T.GtkContainer,
		widget *T.GtkWidget)

	Gtk_container_set_resize_mode func(
		container *T.GtkContainer,
		resize_mode T.GtkResizeMode)

	Gtk_container_get_resize_mode func(
		container *T.GtkContainer) T.GtkResizeMode

	Gtk_container_check_resize func(
		container *T.GtkContainer)

	Gtk_container_foreach func(
		container *T.GtkContainer,
		callback T.GtkCallback,
		callback_data T.Gpointer)

	Gtk_container_foreach_full func(
		container *T.GtkContainer,
		callback T.GtkCallback,
		marshal T.GtkCallbackMarshal,
		callback_dataGpointer,
		notify T.GDestroyNotify)

	Gtk_container_get_children func(
		container *T.GtkContainer) *T.GList

	Gtk_container_propagate_expose func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		event *T.GdkEventExpose)

	Gtk_container_set_focus_chain func(
		container *T.GtkContainer,
		focusable_widgets *T.GList)

	Gtk_container_get_focus_chain func(
		container *T.GtkContainer,
		focusable_widgets **T.GList) T.Gboolean

	Gtk_container_unset_focus_chain func(
		container *T.GtkContainer)

	Gtk_container_set_reallocate_redraws func(
		container *T.GtkContainer,
		needs_redraws T.Gboolean)

	Gtk_container_set_focus_child func(
		container *T.GtkContainer,
		child *T.GtkWidget)

	Gtk_container_get_focus_child func(
		container *T.GtkContainer) *T.GtkWidget

	Gtk_container_set_focus_vadjustment func(
		container *T.GtkContainer,
		adjustment *T.GtkAdjustment)

	Gtk_container_get_focus_vadjustment func(
		container *T.GtkContainer) *T.GtkAdjustment

	Gtk_container_set_focus_hadjustment func(
		container *T.GtkContainer,
		adjustment *T.GtkAdjustment)

	Gtk_container_get_focus_hadjustment func(
		container *T.GtkContainer) *T.GtkAdjustment

	Gtk_container_resize_children func(
		container *T.GtkContainer)

	Gtk_container_child_type func(
		container *T.GtkContainer) T.GType

	Gtk_container_class_install_child_property func(
		cclass *T.GtkContainerClass,
		property_id T.Guint,
		pspec *T.GParamSpec)

	Gtk_container_class_find_child_property func(
		cclass *T.GObjectClass,
		property_name string) *T.GParamSpec

	Gtk_container_class_list_child_properties func(
		cclass *T.GObjectClass,
		n_properties *T.Guint) **T.GParamSpec

	Gtk_container_add_with_properties func(
		container *T.GtkContainer, widget *T.GtkWidget,
		first_prop_name string, v ...VArg)

	Gtk_container_child_set func(
		container *T.GtkContainer, child *T.GtkWidget,
		first_prop_name string, v ...VArg)

	Gtk_container_child_get func(container *T.GtkContainer,
		child *T.GtkWidget, first_prop_name string, v ...VArg)

	Gtk_container_child_set_valist func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		first_property_name string,
		var_args T.Va_list)

	Gtk_container_child_get_valist func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		first_property_name string,
		var_args T.Va_list)

	Gtk_container_child_set_property func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		property_name string,
		value *T.GValue)

	Gtk_container_child_get_property func(
		container *T.GtkContainer,
		child *T.GtkWidget,
		property_name string,
		value *T.GValue)

	Gtk_container_forall func(
		container *T.GtkContainer,
		callback T.GtkCallback,
		callback_data T.Gpointer)

	Gtk_bin_get_type func() T.GType

	Gtk_bin_get_child func(
		bin *T.GtkBin) *T.GtkWidget

	Gtk_window_get_type func() T.GType

	Gtk_window_new func(
		t T.GtkWindowType) *T.GtkWidget

	Gtk_window_set_title func(
		window *T.GtkWindow,
		title string)

	Gtk_window_get_title func(
		window *T.GtkWindow) string

	Gtk_window_set_wmclass func(
		window *T.GtkWindow,
		wmclass_name string,
		wmclass_class string)

	Gtk_window_set_role func(
		window *T.GtkWindow,
		role string)

	Gtk_window_set_startup_id func(
		window *T.GtkWindow,
		startup_id string)

	Gtk_window_get_role func(
		window *T.GtkWindow) string

	Gtk_window_add_accel_group func(
		window *T.GtkWindow,
		accel_group *T.GtkAccelGroup)

	Gtk_window_remove_accel_group func(
		window *T.GtkWindow,
		accel_group *T.GtkAccelGroup)

	Gtk_window_set_position func(
		window *T.GtkWindow,
		position T.GtkWindowPosition)

	Gtk_window_activate_focus func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_focus func(
		window *T.GtkWindow,
		focus *T.GtkWidget)

	Gtk_window_get_focus func(
		window *T.GtkWindow) *T.GtkWidget

	Gtk_window_set_default func(
		window *T.GtkWindow,
		default_widget *T.GtkWidget)

	Gtk_window_get_default_widget func(
		window *T.GtkWindow) *T.GtkWidget

	Gtk_window_activate_default func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_transient_for func(
		window *T.GtkWindow,
		parent *T.GtkWindow)

	Gtk_window_get_transient_for func(
		window *T.GtkWindow) *T.GtkWindow

	Gtk_window_set_opacity func(
		window *T.GtkWindow,
		opacity T.Gdouble)

	Gtk_window_get_opacity func(
		window *T.GtkWindow) T.Gdouble

	Gtk_window_set_type_hint func(
		window *T.GtkWindow,
		hint T.GdkWindowTypeHint)

	Gtk_window_get_type_hint func(
		window *T.GtkWindow) T.GdkWindowTypeHint

	Gtk_window_set_skip_taskbar_hint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_skip_taskbar_hint func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_skip_pager_hint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_skip_pager_hint func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_urgency_hint func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_urgency_hint func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_accept_focus func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_accept_focus func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_focus_on_map func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_focus_on_map func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_destroy_with_parent func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_destroy_with_parent func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_mnemonics_visible func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_mnemonics_visible func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_resizable func(
		window *T.GtkWindow,
		resizable T.Gboolean)

	Gtk_window_get_resizable func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_gravity func(
		window *T.GtkWindow,
		gravity T.GdkGravity)

	Gtk_window_get_gravity func(
		window *T.GtkWindow) T.GdkGravity

	Gtk_window_set_geometry_hints func(
		window *T.GtkWindow,
		geometry_widget *T.GtkWidget,
		geometry *T.GdkGeometry,
		geom_mask T.GdkWindowHints)

	Gtk_window_set_screen func(
		window *T.GtkWindow,
		screen *T.GdkScreen)

	Gtk_window_get_screen func(
		window *T.GtkWindow) *T.GdkScreen

	Gtk_window_is_active func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_has_toplevel_focus func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_has_frame func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_has_frame func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_frame_dimensions func(
		window *T.GtkWindow,
		left T.Gint,
		top T.Gint,
		right T.Gint,
		bottom T.Gint)

	Gtk_window_get_frame_dimensions func(
		window *T.GtkWindow,
		left *T.Gint,
		top *T.Gint,
		right *T.Gint,
		bottom *T.Gint)

	Gtk_window_set_decorated func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_decorated func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_deletable func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_get_deletable func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_set_icon_list func(
		window *T.GtkWindow,
		list *T.GList)

	Gtk_window_get_icon_list func(
		window *T.GtkWindow) *T.GList

	Gtk_window_set_icon func(
		window *T.GtkWindow,
		icon *T.GdkPixbuf)

	Gtk_window_set_icon_name func(
		window *T.GtkWindow,
		name string)

	Gtk_window_set_icon_from_file func(
		window *T.GtkWindow,
		filename string,
		err **T.GError) T.Gboolean

	Gtk_window_get_icon func(
		window *T.GtkWindow) *T.GdkPixbuf

	Gtk_window_get_icon_name func(
		window *T.GtkWindow) string

	Gtk_window_set_default_icon_list func(
		list *T.GList)

	Gtk_window_get_default_icon_list func() *T.GList

	Gtk_window_set_default_icon func(
		icon *T.GdkPixbuf)

	Gtk_window_set_default_icon_name func(
		name string)

	Gtk_window_get_default_icon_name func() string

	Gtk_window_set_default_icon_from_file func(
		filename string,
		err **T.GError) T.Gboolean

	Gtk_window_set_auto_startup_notification func(
		setting T.Gboolean)

	Gtk_window_set_modal func(
		window *T.GtkWindow,
		modal T.Gboolean)

	Gtk_window_get_modal func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_list_toplevels func() *T.GList

	Gtk_window_add_mnemonic func(
		window *T.GtkWindow,
		keyval T.Guint,
		target *T.GtkWidget)

	Gtk_window_remove_mnemonic func(
		window *T.GtkWindow,
		keyval T.Guint,
		target *T.GtkWidget)

	Gtk_window_mnemonic_activate func(
		window *T.GtkWindow,
		keyval T.Guint,
		modifier T.GdkModifierType) T.Gboolean

	Gtk_window_set_mnemonic_modifier func(
		window *T.GtkWindow,
		modifier T.GdkModifierType)

	Gtk_window_get_mnemonic_modifier func(
		window *T.GtkWindow) T.GdkModifierType

	Gtk_window_activate_key func(
		window *T.GtkWindow,
		event *T.GdkEventKey) T.Gboolean

	Gtk_window_propagate_key_event func(
		window *T.GtkWindow,
		event *T.GdkEventKey) T.Gboolean

	Gtk_window_present func(
		window *T.GtkWindow)

	Gtk_window_present_with_time func(
		window *T.GtkWindow,
		timestamp T.Guint32)

	Gtk_window_iconify func(
		window *T.GtkWindow)

	Gtk_window_deiconify func(
		window *T.GtkWindow)

	Gtk_window_stick func(
		window *T.GtkWindow)

	Gtk_window_unstick func(
		window *T.GtkWindow)

	Gtk_window_maximize func(
		window *T.GtkWindow)

	Gtk_window_unmaximize func(
		window *T.GtkWindow)

	Gtk_window_fullscreen func(
		window *T.GtkWindow)

	Gtk_window_unfullscreen func(
		window *T.GtkWindow)

	Gtk_window_set_keep_above func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_set_keep_below func(
		window *T.GtkWindow,
		setting T.Gboolean)

	Gtk_window_begin_resize_drag func(
		window *T.GtkWindow,
		edge T.GdkWindowEdge,
		button T.Gint,
		root_x T.Gint,
		root_y T.Gint,
		timestamp T.Guint32)

	Gtk_window_begin_move_drag func(
		window *T.GtkWindow,
		button T.Gint,
		root_x T.Gint,
		root_y T.Gint,
		timestamp T.Guint32)

	Gtk_window_set_policy func(
		window *T.GtkWindow,
		allow_shrink T.Gint,
		allow_grow T.Gint,
		auto_shrink T.Gint)

	Gtk_window_set_default_size func(
		window *T.GtkWindow,
		width T.Gint,
		height T.Gint)

	Gtk_window_get_default_size func(
		window *T.GtkWindow,
		width *T.Gint,
		height *T.Gint)

	Gtk_window_resize func(
		window *T.GtkWindow,
		width T.Gint,
		height T.Gint)

	Gtk_window_get_size func(
		window *T.GtkWindow,
		width *T.Gint,
		height *T.Gint)

	Gtk_window_move func(
		window *T.GtkWindow,
		x T.Gint,
		y T.Gint)

	Gtk_window_get_position func(
		window *T.GtkWindow,
		root_x *T.Gint,
		root_y *T.Gint)

	Gtk_window_parse_geometry func(
		window *T.GtkWindow,
		geometry string) T.Gboolean

	Gtk_window_get_group func(
		window *T.GtkWindow) *T.GtkWindowGroup

	Gtk_window_has_group func(
		window *T.GtkWindow) T.Gboolean

	Gtk_window_reshow_with_initial_size func(
		window *T.GtkWindow)

	Gtk_window_get_window_type func(
		window *T.GtkWindow) T.GtkWindowType

	Gtk_window_group_get_type func() T.GType

	Gtk_window_group_new func() *T.GtkWindowGroup

	Gtk_window_group_add_window func(
		window_group *T.GtkWindowGroup,
		window *T.GtkWindow)

	Gtk_window_group_remove_window func(
		window_group *T.GtkWindowGroup,
		window *T.GtkWindow)

	Gtk_window_group_list_windows func(
		window_group *T.GtkWindowGroup) *T.GList

	Gtk_window_remove_embedded_xid func(
		window *T.GtkWindow,
		xid T.GdkNativeWindow)

	Gtk_window_add_embedded_xid func(
		window *T.GtkWindow,
		xid T.GdkNativeWindow)

	Gtk_window_group_get_current_grab func(
		window_group *T.GtkWindowGroup) *T.GtkWidget

	Gtk_dialog_get_type func() T.GType

	Gtk_dialog_new func() *T.GtkWidget

	Gtk_dialog_new_with_buttons func(
		title string, parent *T.GtkWindow,
		flags T.GtkDialogFlags, first_button_text string,
		v ...VArg) *T.GtkWidget

	Gtk_dialog_add_action_widget func(
		dialog *T.GtkDialog,
		child *T.GtkWidget,
		response_id T.Gint)

	Gtk_dialog_add_button func(
		dialog *T.GtkDialog,
		button_text string,
		response_id T.Gint) *T.GtkWidget

	Gtk_dialog_add_buttons func(dialog *T.GtkDialog,
		first_button_text string, v ...VArg)

	Gtk_dialog_set_response_sensitive func(
		dialog *T.GtkDialog,
		response_id T.Gint,
		setting T.Gboolean)

	Gtk_dialog_set_default_response func(
		dialog *T.GtkDialog,
		response_id T.Gint)

	Gtk_dialog_get_widget_for_response func(
		dialog *T.GtkDialog,
		response_id T.Gint) *T.GtkWidget

	Gtk_dialog_get_response_for_widget func(
		dialog *T.GtkDialog,
		widget *T.GtkWidget) T.Gint

	Gtk_dialog_set_has_separator func(
		dialog *T.GtkDialog,
		setting T.Gboolean)

	Gtk_dialog_get_has_separator func(
		dialog *T.GtkDialog) T.Gboolean

	Gtk_alternative_dialog_button_order func(
		screen *T.GdkScreen) T.Gboolean

	Gtk_dialog_set_alternative_button_order func(
		dialog *T.GtkDialog, first_response_id T.Gint, v ...VArg)

	Gtk_dialog_set_alternative_button_order_from_array func(
		dialog *T.GtkDialog,
		n_params T.Gint,
		new_order *T.Gint)

	Gtk_dialog_response func(
		dialog *T.GtkDialog,
		response_id T.Gint)

	Gtk_dialog_run func(
		dialog *T.GtkDialog) T.Gint

	Gtk_dialog_get_action_area func(
		dialog *T.GtkDialog) *T.GtkWidget

	Gtk_dialog_get_content_area func(
		dialog *T.GtkDialog) *T.GtkWidget

	Gtk_about_dialog_get_type func() T.GType

	Gtk_about_dialog_new func() *T.GtkWidget

	Gtk_show_about_dialog func(parent *T.GtkWindow,
		first_property_name string, v ...VArg)

	Gtk_about_dialog_get_name func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_name func(
		about *T.GtkAboutDialog,
		name string)

	Gtk_about_dialog_get_program_name func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_program_name func(
		about *T.GtkAboutDialog,
		name string)

	Gtk_about_dialog_get_version func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_version func(
		about *T.GtkAboutDialog,
		version string)

	Gtk_about_dialog_get_copyright func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_copyright func(
		about *T.GtkAboutDialog,
		copyright string)

	Gtk_about_dialog_get_comments func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_comments func(
		about *T.GtkAboutDialog,
		comments string)

	Gtk_about_dialog_get_license func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_license func(
		about *T.GtkAboutDialog,
		license string)

	Gtk_about_dialog_get_wrap_license func(
		about *T.GtkAboutDialog) T.Gboolean

	Gtk_about_dialog_set_wrap_license func(
		about *T.GtkAboutDialog,
		wrap_license T.Gboolean)

	Gtk_about_dialog_get_website func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_website func(
		about *T.GtkAboutDialog,
		website string)

	Gtk_about_dialog_get_website_label func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_website_label func(
		about *T.GtkAboutDialog,
		website_label string)

	Gtk_about_dialog_get_authors func(
		about *T.GtkAboutDialog) **T.Gchar

	Gtk_about_dialog_set_authors func(
		about *T.GtkAboutDialog,
		authors **T.Gchar)

	Gtk_about_dialog_get_documenters func(
		about *T.GtkAboutDialog) **T.Gchar

	Gtk_about_dialog_set_documenters func(
		about *T.GtkAboutDialog,
		documenters **T.Gchar)

	Gtk_about_dialog_get_artists func(
		about *T.GtkAboutDialog) **T.Gchar

	Gtk_about_dialog_set_artists func(
		about *T.GtkAboutDialog,
		artists **T.Gchar)

	Gtk_about_dialog_get_translator_credits func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_translator_credits func(
		about *T.GtkAboutDialog,
		translator_credits string)

	Gtk_about_dialog_get_logo func(
		about *T.GtkAboutDialog) *T.GdkPixbuf

	Gtk_about_dialog_set_logo func(
		about *T.GtkAboutDialog,
		logo *T.GdkPixbuf)

	Gtk_about_dialog_get_logo_icon_name func(
		about *T.GtkAboutDialog) string

	Gtk_about_dialog_set_logo_icon_name func(
		about *T.GtkAboutDialog,
		icon_name string)

	Gtk_about_dialog_set_email_hook func(
		f T.GtkAboutDialogActivateLinkFunc,
		dataGpointer,
		destroy T.GDestroyNotify) T.GtkAboutDialogActivateLinkFunc

	Gtk_about_dialog_set_url_hook func(
		f T.GtkAboutDialogActivateLinkFunc,
		dataGpointer,
		destroy T.GDestroyNotify) T.GtkAboutDialogActivateLinkFunc

	Gtk_misc_get_type func() T.GType

	Gtk_misc_set_alignment func(
		misc *T.GtkMisc,
		xalign T.Gfloat,
		yalign T.Gfloat)

	Gtk_misc_get_alignment func(
		misc *T.GtkMisc,
		xalign *T.Gfloat,
		yalign *T.Gfloat)

	Gtk_misc_set_padding func(
		misc *T.GtkMisc,
		xpad T.Gint,
		ypad T.Gint)

	Gtk_misc_get_padding func(
		misc *T.GtkMisc,
		xpad *T.Gint,
		ypad *T.Gint)

	Gtk_menu_shell_get_type func() T.GType

	Gtk_menu_shell_append func(
		menu_shell *T.GtkMenuShell,
		child *T.GtkWidget)

	Gtk_menu_shell_prepend func(
		menu_shell *T.GtkMenuShell,
		child *T.GtkWidget)

	Gtk_menu_shell_insert func(
		menu_shell *T.GtkMenuShell,
		child *T.GtkWidget,
		position T.Gint)

	Gtk_menu_shell_deactivate func(
		menu_shell *T.GtkMenuShell)

	Gtk_menu_shell_select_item func(
		menu_shell *T.GtkMenuShell,
		menu_item *T.GtkWidget)

	Gtk_menu_shell_deselect func(
		menu_shell *T.GtkMenuShell)

	Gtk_menu_shell_activate_item func(
		menu_shell *T.GtkMenuShell,
		menu_item *T.GtkWidget,
		force_deactivate T.Gboolean)

	Gtk_menu_shell_select_first func(
		menu_shell *T.GtkMenuShell,
		search_sensitive T.Gboolean)

	Gtk_menu_shell_cancel func(
		menu_shell *T.GtkMenuShell)

	Gtk_menu_shell_get_take_focus func(
		menu_shell *T.GtkMenuShell) T.Gboolean

	Gtk_menu_shell_set_take_focus func(
		menu_shell *T.GtkMenuShell,
		take_focus T.Gboolean)

	Gtk_menu_get_type func() T.GType

	Gtk_menu_new func() *T.GtkWidget

	Gtk_menu_popup func(
		menu *T.GtkMenu,
		parent_menu_shell *T.GtkWidget,
		parent_menu_item *T.GtkWidget,
		f T.GtkMenuPositionFunc,
		dataGpointer,
		button T.Guint,
		activate_time T.Guint32)

	Gtk_menu_reposition func(
		menu *T.GtkMenu)

	Gtk_menu_popdown func(
		menu *T.GtkMenu)

	Gtk_menu_get_active func(
		menu *T.GtkMenu) *T.GtkWidget

	Gtk_menu_set_active func(
		menu *T.GtkMenu,
		index_ T.Guint)

	Gtk_menu_set_accel_group func(
		menu *T.GtkMenu,
		accel_group *T.GtkAccelGroup)

	Gtk_menu_get_accel_group func(
		menu *T.GtkMenu) *T.GtkAccelGroup

	Gtk_menu_set_accel_path func(
		menu *T.GtkMenu,
		accel_path string)

	Gtk_menu_get_accel_path func(
		menu *T.GtkMenu) string

	Gtk_menu_attach_to_widget func(
		menu *T.GtkMenu,
		attach_widget *T.GtkWidget,
		detacher T.GtkMenuDetachFunc)

	Gtk_menu_detach func(
		menu *T.GtkMenu)

	Gtk_menu_get_attach_widget func(
		menu *T.GtkMenu) *T.GtkWidget

	Gtk_menu_set_tearoff_state func(
		menu *T.GtkMenu,
		torn_off T.Gboolean)

	Gtk_menu_get_tearoff_state func(
		menu *T.GtkMenu) T.Gboolean

	Gtk_menu_set_title func(
		menu *T.GtkMenu,
		title string)

	Gtk_menu_get_title func(
		menu *T.GtkMenu) string

	Gtk_menu_reorder_child func(
		menu *T.GtkMenu,
		child *T.GtkWidget,
		position T.Gint)

	Gtk_menu_set_screen func(
		menu *T.GtkMenu,
		screen *T.GdkScreen)

	Gtk_menu_attach func(
		menu *T.GtkMenu,
		child *T.GtkWidget,
		left_attach T.Guint,
		right_attach T.Guint,
		top_attach T.Guint,
		bottom_attach T.Guint)

	Gtk_menu_set_monitor func(
		menu *T.GtkMenu,
		monitor_num T.Gint)

	Gtk_menu_get_monitor func(
		menu *T.GtkMenu) T.Gint

	Gtk_menu_get_for_attach_widget func(
		widget *T.GtkWidget) *T.GList

	Gtk_menu_set_reserve_toggle_size func(
		menu *T.GtkMenu,
		reserve_toggle_size T.Gboolean)

	Gtk_menu_get_reserve_toggle_size func(
		menu *T.GtkMenu) T.Gboolean

	Gtk_label_get_type func() T.GType

	Gtk_label_new func(
		str string) *T.GtkWidget

	Gtk_label_new_with_mnemonic func(
		str string) *T.GtkWidget

	Gtk_label_set_text func(
		label *T.GtkLabel,
		str string)

	Gtk_label_get_text func(
		label *T.GtkLabel) string

	Gtk_label_set_attributes func(
		label *T.GtkLabel,
		attrs *T.PangoAttrList)

	Gtk_label_get_attributes func(
		label *T.GtkLabel) *T.PangoAttrList

	Gtk_label_set_label func(
		label *T.GtkLabel,
		str string)

	Gtk_label_get_label func(
		label *T.GtkLabel) string

	Gtk_label_set_markup func(
		label *T.GtkLabel,
		str string)

	Gtk_label_set_use_markup func(
		label *T.GtkLabel,
		setting T.Gboolean)

	Gtk_label_get_use_markup func(
		label *T.GtkLabel) T.Gboolean

	Gtk_label_set_use_underline func(
		label *T.GtkLabel,
		setting T.Gboolean)

	Gtk_label_get_use_underline func(
		label *T.GtkLabel) T.Gboolean

	Gtk_label_set_markup_with_mnemonic func(
		label *T.GtkLabel,
		str string)

	Gtk_label_get_mnemonic_keyval func(
		label *T.GtkLabel) T.Guint

	Gtk_label_set_mnemonic_widget func(
		label *T.GtkLabel,
		widget *T.GtkWidget)

	Gtk_label_get_mnemonic_widget func(
		label *T.GtkLabel) *T.GtkWidget

	Gtk_label_set_text_with_mnemonic func(
		label *T.GtkLabel,
		str string)

	Gtk_label_set_justify func(
		label *T.GtkLabel,
		jtype T.GtkJustification)

	Gtk_label_get_justify func(
		label *T.GtkLabel) T.GtkJustification

	Gtk_label_set_ellipsize func(
		label *T.GtkLabel,
		mode T.PangoEllipsizeMode)

	Gtk_label_get_ellipsize func(
		label *T.GtkLabel) T.PangoEllipsizeMode

	Gtk_label_set_width_chars func(
		label *T.GtkLabel,
		n_chars T.Gint)

	Gtk_label_get_width_chars func(
		label *T.GtkLabel) T.Gint

	Gtk_label_set_max_width_chars func(
		label *T.GtkLabel,
		n_chars T.Gint)

	Gtk_label_get_max_width_chars func(
		label *T.GtkLabel) T.Gint

	Gtk_label_set_pattern func(
		label *T.GtkLabel,
		pattern string)

	Gtk_label_set_line_wrap func(
		label *T.GtkLabel,
		wrap T.Gboolean)

	Gtk_label_get_line_wrap func(
		label *T.GtkLabel) T.Gboolean

	Gtk_label_set_line_wrap_mode func(
		label *T.GtkLabel,
		wrap_mode T.PangoWrapMode)

	Gtk_label_get_line_wrap_mode func(
		label *T.GtkLabel) T.PangoWrapMode

	Gtk_label_set_selectable func(
		label *T.GtkLabel,
		setting T.Gboolean)

	Gtk_label_get_selectable func(
		label *T.GtkLabel) T.Gboolean

	Gtk_label_set_angle func(
		label *T.GtkLabel,
		angle T.Gdouble)

	Gtk_label_get_angle func(
		label *T.GtkLabel) T.Gdouble

	Gtk_label_select_region func(
		label *T.GtkLabel,
		start_offset T.Gint,
		end_offset T.Gint)

	Gtk_label_get_selection_bounds func(
		label *T.GtkLabel,
		start *T.Gint,
		end *T.Gint) T.Gboolean

	Gtk_label_get_layout func(
		label *T.GtkLabel) *T.PangoLayout

	Gtk_label_get_layout_offsets func(
		label *T.GtkLabel,
		x *T.Gint,
		y *T.Gint)

	Gtk_label_set_single_line_mode func(
		label *T.GtkLabel,
		single_line_mode T.Gboolean)

	Gtk_label_get_single_line_mode func(
		label *T.GtkLabel) T.Gboolean

	Gtk_label_get_current_uri func(
		label *T.GtkLabel) string

	Gtk_label_set_track_visited_links func(
		label *T.GtkLabel,
		track_links T.Gboolean)

	Gtk_label_get_track_visited_links func(
		label *T.GtkLabel) T.Gboolean

	Gtk_label_get func(
		label *T.GtkLabel,
		str **T.Gchar)

	Gtk_label_parse_uline func(
		label *T.GtkLabel,
		string string) T.Guint

	Gtk_accel_label_get_type func() T.GType

	Gtk_accel_label_new func(
		string string) *T.GtkWidget

	Gtk_accel_label_get_accel_widget func(
		accel_label *T.GtkAccelLabel) *T.GtkWidget

	Gtk_accel_label_get_accel_width func(
		accel_label *T.GtkAccelLabel) T.Guint

	Gtk_accel_label_set_accel_widget func(
		accel_label *T.GtkAccelLabel,
		accel_widget *T.GtkWidget)

	Gtk_accel_label_set_accel_closure func(
		accel_label *T.GtkAccelLabel,
		accel_closure *T.GClosure)

	Gtk_accel_label_refetch func(
		accel_label *T.GtkAccelLabel) T.Gboolean

	Gtk_accel_map_add_entry func(
		accel_path string,
		accel_key T.Guint,
		accel_mods T.GdkModifierType)

	Gtk_accel_map_lookup_entry func(
		accel_path string,
		key *T.GtkAccelKey) T.Gboolean

	Gtk_accel_map_change_entry func(
		accel_path string,
		accel_key T.Guint,
		accel_mods T.GdkModifierType,
		replace T.Gboolean) T.Gboolean

	Gtk_accel_map_load func(
		file_name string)

	Gtk_accel_map_save func(
		file_name string)

	Gtk_accel_map_foreach func(
		dataGpointer,
		foreach_func T.GtkAccelMapForeach)

	Gtk_accel_map_load_fd func(
		fd T.Gint)

	Gtk_accel_map_load_scanner func(
		scanner *T.GScanner)

	Gtk_accel_map_save_fd func(
		fd T.Gint)

	Gtk_accel_map_lock_path func(
		accel_path string)

	Gtk_accel_map_unlock_path func(
		accel_path string)

	Gtk_accel_map_add_filter func(
		filter_pattern string)

	Gtk_accel_map_foreach_unfiltered func(
		dataGpointer,
		foreach_func T.GtkAccelMapForeach)

	Gtk_accel_map_get_type func() T.GType

	Gtk_accel_map_get func() *T.GtkAccelMap

	Gtk_accessible_get_type func() T.GType

	Gtk_accessible_set_widget func(
		accessible *T.GtkAccessible,
		widget *T.GtkWidget)

	Gtk_accessible_get_widget func(
		accessible *T.GtkAccessible) *T.GtkWidget

	Gtk_accessible_connect_widget_destroyed func(
		accessible *T.GtkAccessible)

	Gtk_action_get_type func() T.GType

	Gtk_action_new func(
		name string,
		label string,
		tooltip string,
		stock_id string) *T.GtkAction

	Gtk_action_get_name func(
		action *T.GtkAction) string

	Gtk_action_is_sensitive func(
		action *T.GtkAction) T.Gboolean

	Gtk_action_get_sensitive func(
		action *T.GtkAction) T.Gboolean

	Gtk_action_set_sensitive func(
		action *T.GtkAction,
		sensitive T.Gboolean)

	Gtk_action_is_visible func(
		action *T.GtkAction) T.Gboolean

	Gtk_action_get_visible func(
		action *T.GtkAction) T.Gboolean

	Gtk_action_set_visible func(
		action *T.GtkAction,
		visible T.Gboolean)

	Gtk_action_activate func(
		action *T.GtkAction)

	Gtk_action_create_icon func(
		action *T.GtkAction,
		icon_size T.GtkIconSize) *T.GtkWidget

	Gtk_action_create_menu_item func(
		action *T.GtkAction) *T.GtkWidget

	Gtk_action_create_tool_item func(
		action *T.GtkAction) *T.GtkWidget

	Gtk_action_create_menu func(
		action *T.GtkAction) *T.GtkWidget

	Gtk_action_get_proxies func(
		action *T.GtkAction) *T.GSList

	Gtk_action_connect_accelerator func(
		action *T.GtkAction)

	Gtk_action_disconnect_accelerator func(
		action *T.GtkAction)

	Gtk_action_get_accel_path func(
		action *T.GtkAction) string

	Gtk_action_get_accel_closure func(
		action *T.GtkAction) *T.GClosure

	Gtk_widget_get_action func(
		widget *T.GtkWidget) *T.GtkAction

	Gtk_action_connect_proxy func(
		action *T.GtkAction,
		proxy *T.GtkWidget)

	Gtk_action_disconnect_proxy func(
		action *T.GtkAction,
		proxy *T.GtkWidget)

	Gtk_action_block_activate_from func(
		action *T.GtkAction,
		proxy *T.GtkWidget)

	Gtk_action_unblock_activate_from func(
		action *T.GtkAction,
		proxy *T.GtkWidget)

	Gtk_action_block_activate func(
		action *T.GtkAction)

	Gtk_action_unblock_activate func(
		action *T.GtkAction)

	Gtk_action_set_accel_path func(
		action *T.GtkAction,
		accel_path string)

	Gtk_action_set_accel_group func(
		action *T.GtkAction,
		accel_group *T.GtkAccelGroup)

	Gtk_action_set_label func(
		action *T.GtkAction,
		label string)

	Gtk_action_get_label func(
		action *T.GtkAction) string

	Gtk_action_set_short_label func(
		action *T.GtkAction,
		short_label string)

	Gtk_action_get_short_label func(
		action *T.GtkAction) string

	Gtk_action_set_tooltip func(
		action *T.GtkAction,
		tooltip string)

	Gtk_action_get_tooltip func(
		action *T.GtkAction) string

	Gtk_action_set_stock_id func(
		action *T.GtkAction,
		stock_id string)

	Gtk_action_get_stock_id func(
		action *T.GtkAction) string

	Gtk_action_set_gicon func(
		action *T.GtkAction,
		icon *T.GIcon)

	Gtk_action_get_gicon func(
		action *T.GtkAction) *T.GIcon

	Gtk_action_set_icon_name func(
		action *T.GtkAction,
		icon_name string)

	Gtk_action_get_icon_name func(
		action *T.GtkAction) string

	Gtk_action_set_visible_horizontal func(
		action *T.GtkAction,
		visible_horizontal T.Gboolean)

	Gtk_action_get_visible_horizontal func(
		action *T.GtkAction) T.Gboolean

	Gtk_action_set_visible_vertical func(
		action *T.GtkAction,
		visible_vertical T.Gboolean)

	Gtk_action_get_visible_vertical func(
		action *T.GtkAction) T.Gboolean

	Gtk_action_set_is_important func(
		action *T.GtkAction,
		is_important T.Gboolean)

	Gtk_action_get_is_important func(
		action *T.GtkAction) T.Gboolean

	Gtk_action_set_always_show_image func(
		action *T.GtkAction,
		always_show T.Gboolean)

	Gtk_action_get_always_show_image func(
		action *T.GtkAction) T.Gboolean

	Gtk_action_group_get_type func() T.GType

	Gtk_action_group_new func(
		name string) *T.GtkActionGroup

	Gtk_action_group_get_name func(
		action_group *T.GtkActionGroup) string

	Gtk_action_group_get_sensitive func(
		action_group *T.GtkActionGroup) T.Gboolean

	Gtk_action_group_set_sensitive func(
		action_group *T.GtkActionGroup,
		sensitive T.Gboolean)

	Gtk_action_group_get_visible func(
		action_group *T.GtkActionGroup) T.Gboolean

	Gtk_action_group_set_visible func(
		action_group *T.GtkActionGroup,
		visible T.Gboolean)

	Gtk_action_group_get_action func(
		action_group *T.GtkActionGroup,
		action_name string) *T.GtkAction

	Gtk_action_group_list_actions func(
		action_group *T.GtkActionGroup) *T.GList

	Gtk_action_group_add_action func(
		action_group *T.GtkActionGroup,
		action *T.GtkAction)

	Gtk_action_group_add_action_with_accel func(
		action_group *T.GtkActionGroup,
		action *T.GtkAction,
		accelerator string)

	Gtk_action_group_remove_action func(
		action_group *T.GtkActionGroup,
		action *T.GtkAction)

	Gtk_action_group_add_actions func(
		action_group *T.GtkActionGroup,
		entries *T.GtkActionEntry,
		n_entries T.Guint,
		user_data T.Gpointer)

	Gtk_action_group_add_toggle_actions func(
		action_group *T.GtkActionGroup,
		entries *T.GtkToggleActionEntry,
		n_entries T.Guint,
		user_data T.Gpointer)

	Gtk_action_group_add_radio_actions func(
		action_group *T.GtkActionGroup,
		entries *T.GtkRadioActionEntry,
		n_entries T.Guint,
		value T.Gint,
		on_change T.GCallback,
		user_data T.Gpointer)

	Gtk_action_group_add_actions_full func(
		action_group *T.GtkActionGroup,
		entries *T.GtkActionEntry,
		n_entries T.Guint,
		user_dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_action_group_add_toggle_actions_full func(
		action_group *T.GtkActionGroup,
		entries *T.GtkToggleActionEntry,
		n_entries T.Guint,
		user_dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_action_group_add_radio_actions_full func(
		action_group *T.GtkActionGroup,
		entries *T.GtkRadioActionEntry,
		n_entries T.Guint,
		value T.Gint,
		on_change T.GCallback,
		user_dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_action_group_set_translate_func func(
		action_group *T.GtkActionGroup,
		f T.GtkTranslateFunc,
		dataGpointer,
		notify T.GDestroyNotify)

	Gtk_action_group_set_translation_domain func(
		action_group *T.GtkActionGroup,
		domain string)

	Gtk_action_group_translate_string func(
		action_group *T.GtkActionGroup,
		string string) string

	Gtk_activatable_get_type func() T.GType

	Gtk_activatable_sync_action_properties func(
		activatable *T.GtkActivatable,
		action *T.GtkAction)

	Gtk_activatable_set_related_action func(
		activatable *T.GtkActivatable,
		action *T.GtkAction)

	Gtk_activatable_get_related_action func(
		activatable *T.GtkActivatable) *T.GtkAction

	Gtk_activatable_set_use_action_appearance func(
		activatable *T.GtkActivatable,
		use_appearance T.Gboolean)

	Gtk_activatable_get_use_action_appearance func(
		activatable *T.GtkActivatable) T.Gboolean

	Gtk_activatable_do_set_related_action func(
		activatable *T.GtkActivatable,
		action *T.GtkAction)

	Gtk_alignment_get_type func() T.GType

	Gtk_alignment_new func(
		xalign T.Gfloat,
		yalign T.Gfloat,
		xscale T.Gfloat,
		yscale T.Gfloat) *T.GtkWidget

	Gtk_alignment_set func(
		alignment *T.GtkAlignment,
		xalign T.Gfloat,
		yalign T.Gfloat,
		xscale T.Gfloat,
		yscale T.Gfloat)

	Gtk_alignment_set_padding func(
		alignment *T.GtkAlignment,
		padding_top T.Guint,
		padding_bottom T.Guint,
		padding_left T.Guint,
		padding_right T.Guint)

	Gtk_alignment_get_padding func(
		alignment *T.GtkAlignment,
		padding_top *T.Guint,
		padding_bottom *T.Guint,
		padding_left *T.Guint,
		padding_right *T.Guint)

	Gtk_arrow_get_type func() T.GType

	Gtk_arrow_new func(
		arrow_type T.GtkArrowType,
		shadow_type T.GtkShadowType) *T.GtkWidget

	Gtk_arrow_set func(
		arrow *T.GtkArrow,
		arrow_type T.GtkArrowType,
		shadow_type T.GtkShadowType)

	Gtk_frame_get_type func() T.GType

	Gtk_frame_new func(
		label string) *T.GtkWidget

	Gtk_frame_set_label func(
		frame *T.GtkFrame,
		label string)

	Gtk_frame_get_label func(
		frame *T.GtkFrame) string

	Gtk_frame_set_label_widget func(
		frame *T.GtkFrame,
		label_widget *T.GtkWidget)

	Gtk_frame_get_label_widget func(
		frame *T.GtkFrame) *T.GtkWidget

	Gtk_frame_set_label_align func(
		frame *T.GtkFrame,
		xalign T.Gfloat,
		yalign T.Gfloat)

	Gtk_frame_get_label_align func(
		frame *T.GtkFrame,
		xalign *T.Gfloat,
		yalign *T.Gfloat)

	Gtk_frame_set_shadow_type func(
		frame *T.GtkFrame,
		t T.GtkShadowType)

	Gtk_frame_get_shadow_type func(
		frame *T.GtkFrame) T.GtkShadowType

	Gtk_aspect_frame_get_type func() T.GType

	Gtk_aspect_frame_new func(
		label string,
		xalign T.Gfloat,
		yalign T.Gfloat,
		ratio T.Gfloat,
		obey_child T.Gboolean) *T.GtkWidget

	Gtk_aspect_frame_set func(
		aspect_frame *T.GtkAspectFrame,
		xalign T.Gfloat,
		yalign T.Gfloat,
		ratio T.Gfloat,
		obey_child T.Gboolean)

	Gtk_assistant_get_type func() T.GType

	Gtk_assistant_new func() *T.GtkWidget

	Gtk_assistant_get_current_page func(
		assistant *T.GtkAssistant) T.Gint

	Gtk_assistant_set_current_page func(
		assistant *T.GtkAssistant,
		page_num T.Gint)

	Gtk_assistant_get_n_pages func(
		assistant *T.GtkAssistant) T.Gint

	Gtk_assistant_get_nth_page func(
		assistant *T.GtkAssistant,
		page_num T.Gint) *T.GtkWidget

	Gtk_assistant_prepend_page func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget) T.Gint

	Gtk_assistant_append_page func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget) T.Gint

	Gtk_assistant_insert_page func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget,
		position T.Gint) T.Gint

	Gtk_assistant_set_forward_page_func func(
		assistant *T.GtkAssistant,
		page_func T.GtkAssistantPageFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_assistant_set_page_type func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget,
		t T.GtkAssistantPageType)

	Gtk_assistant_get_page_type func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget) T.GtkAssistantPageType

	Gtk_assistant_set_page_title func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget,
		title string)

	Gtk_assistant_get_page_title func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget) string

	Gtk_assistant_set_page_header_image func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget,
		pixbuf *T.GdkPixbuf)

	Gtk_assistant_get_page_header_image func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget) *T.GdkPixbuf

	Gtk_assistant_set_page_side_image func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget,
		pixbuf *T.GdkPixbuf)

	Gtk_assistant_get_page_side_image func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget) *T.GdkPixbuf

	Gtk_assistant_set_page_complete func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget,
		complete T.Gboolean)

	Gtk_assistant_get_page_complete func(
		assistant *T.GtkAssistant,
		page *T.GtkWidget) T.Gboolean

	Gtk_assistant_add_action_widget func(
		assistant *T.GtkAssistant,
		child *T.GtkWidget)

	Gtk_assistant_remove_action_widget func(
		assistant *T.GtkAssistant,
		child *T.GtkWidget)

	Gtk_assistant_update_buttons_state func(
		assistant *T.GtkAssistant)

	Gtk_assistant_commit func(
		assistant *T.GtkAssistant)

	Gtk_box_get_type func() T.GType

	Gtk_box_pack_start func(
		box *T.GtkBox,
		child *T.GtkWidget,
		expand T.Gboolean,
		fill T.Gboolean,
		padding T.Guint)

	Gtk_box_pack_end func(
		box *T.GtkBox,
		child *T.GtkWidget,
		expand T.Gboolean,
		fill T.Gboolean,
		padding T.Guint)

	Gtk_box_pack_start_defaults func(
		box *T.GtkBox,
		widget *T.GtkWidget)

	Gtk_box_pack_end_defaults func(
		box *T.GtkBox,
		widget *T.GtkWidget)

	Gtk_box_set_homogeneous func(
		box *T.GtkBox,
		homogeneous T.Gboolean)

	Gtk_box_get_homogeneous func(
		box *T.GtkBox) T.Gboolean

	Gtk_box_set_spacing func(
		box *T.GtkBox,
		spacing T.Gint)

	Gtk_box_get_spacing func(
		box *T.GtkBox) T.Gint

	Gtk_box_reorder_child func(
		box *T.GtkBox,
		child *T.GtkWidget,
		position T.Gint)

	Gtk_box_query_child_packing func(
		box *T.GtkBox,
		child *T.GtkWidget,
		expand *T.Gboolean,
		fill *T.Gboolean,
		padding *T.Guint,
		pack_type *T.GtkPackType)

	Gtk_box_set_child_packing func(
		box *T.GtkBox,
		child *T.GtkWidget,
		expand T.Gboolean,
		fill T.Gboolean,
		padding T.Guint,
		pack_type T.GtkPackType)

	Gtk_button_box_get_type func() T.GType

	Gtk_button_box_get_layout func(
		widget *T.GtkButtonBox) T.GtkButtonBoxStyle

	Gtk_button_box_set_layout func(
		widget *T.GtkButtonBox,
		layout_style T.GtkButtonBoxStyle)

	Gtk_button_box_get_child_secondary func(
		widget *T.GtkButtonBox,
		child *T.GtkWidget) T.Gboolean

	Gtk_button_box_set_child_secondary func(
		widget *T.GtkButtonBox,
		child *T.GtkWidget,
		is_secondary T.Gboolean)

	Gtk_button_box_set_child_size func(
		widget *T.GtkButtonBox,
		min_width T.Gint,
		min_height T.Gint)

	Gtk_button_box_set_child_ipadding func(
		widget *T.GtkButtonBox,
		ipad_x T.Gint,
		ipad_y T.Gint)

	Gtk_button_box_get_child_size func(
		widget *T.GtkButtonBox,
		min_width *T.Gint,
		min_height *T.Gint)

	Gtk_button_box_get_child_ipadding func(
		widget *T.GtkButtonBox,
		ipad_x *T.Gint,
		ipad_y *T.Gint)

	Gtk_binding_set_new func(
		set_name string) *T.GtkBindingSet

	Gtk_binding_set_by_class func(
		object_class T.Gpointer) *T.GtkBindingSet

	Gtk_binding_set_find func(
		set_name string) *T.GtkBindingSet

	Gtk_bindings_activate func(
		object *T.GtkObject,
		keyval T.Guint,
		modifiers T.GdkModifierType) T.Gboolean

	Gtk_bindings_activate_event func(
		object *T.GtkObject,
		event *T.GdkEventKey) T.Gboolean

	Gtk_binding_set_activate func(
		binding_set *T.GtkBindingSet,
		keyval T.Guint,
		modifiers T.GdkModifierType,
		object *T.GtkObject) T.Gboolean

	Gtk_binding_entry_clear func(
		binding_set *T.GtkBindingSet,
		keyval T.Guint,
		modifiers T.GdkModifierType)

	Gtk_binding_parse_binding func(
		scanner *T.GScanner) T.Guint

	Gtk_binding_entry_skip func(
		binding_set *T.GtkBindingSet,
		keyval T.Guint,
		modifiers T.GdkModifierType)

	Gtk_binding_entry_add_signal func(binding_set *T.GtkBindingSet,
		keyval T.Guint, modifiers T.GdkModifierType,
		signal_name string, n_args T.Guint, varg ...VArg)

	Gtk_binding_entry_add_signall func(
		binding_set *T.GtkBindingSet,
		keyval T.Guint,
		modifiers T.GdkModifierType,
		signal_name string,
		binding_args *T.GSList)

	Gtk_binding_entry_remove func(
		binding_set *T.GtkBindingSet,
		keyval T.Guint,
		modifiers T.GdkModifierType)

	Gtk_binding_set_add_path func(
		binding_set *T.GtkBindingSet,
		path_type T.GtkPathType,
		path_pattern string,
		priority T.GtkPathPriorityType)

	Gtk_builder_get_type func() T.GType

	Gtk_builder_new func() *T.GtkBuilder

	Gtk_builder_add_from_file func(
		builder *T.GtkBuilder,
		filename string,
		error **T.GError) T.Guint

	Gtk_builder_add_from_string func(
		builder *T.GtkBuilder,
		buffer string,
		length T.Gsize,
		error **T.GError) T.Guint

	Gtk_builder_add_objects_from_file func(
		builder *T.GtkBuilder,
		filename string,
		object_ids **T.Gchar,
		error **T.GError) T.Guint

	Gtk_builder_add_objects_from_string func(
		builder *T.GtkBuilder,
		buffer string,
		length T.Gsize,
		object_ids **T.Gchar,
		error **T.GError) T.Guint

	Gtk_builder_get_object func(
		builder *T.GtkBuilder,
		name string) *T.GObject

	Gtk_builder_get_objects func(
		builder *T.GtkBuilder) *T.GSList

	Gtk_builder_connect_signals func(
		builder *T.GtkBuilder,
		user_data T.Gpointer)

	Gtk_builder_connect_signals_full func(
		builder *T.GtkBuilder,
		f T.GtkBuilderConnectFunc,
		user_data T.Gpointer)

	Gtk_builder_set_translation_domain func(
		builder *T.GtkBuilder,
		domain string)

	Gtk_builder_get_translation_domain func(
		builder *T.GtkBuilder) string

	Gtk_builder_get_type_from_name func(
		builder *T.GtkBuilder,
		type_name string) T.GType

	Gtk_builder_value_from_string func(
		builder *T.GtkBuilder,
		pspec *T.GParamSpec,
		string string,
		value *T.GValue,
		error **T.GError) T.Gboolean

	Gtk_builder_value_from_string_type func(
		builder *T.GtkBuilder,
		t T.GType,
		string string,
		value *T.GValue,
		error **T.GError) T.Gboolean

	Gtk_buildable_get_type func() T.GType

	Gtk_buildable_set_name func(
		buildable *T.GtkBuildable,
		name string)

	Gtk_buildable_get_name func(
		buildable *T.GtkBuildable) string

	Gtk_buildable_add_child func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		child *T.GObject,
		typ string)

	Gtk_buildable_set_buildable_property func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		name string,
		value *T.GValue)

	Gtk_buildable_construct_child func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		name string) *T.GObject

	Gtk_buildable_custom_tag_start func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		child *T.GObject,
		tagname string,
		parser *T.GMarkupParser,
		data *T.Gpointer) T.Gboolean

	Gtk_buildable_custom_tag_end func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		child *T.GObject,
		tagname string,
		data *T.Gpointer)

	Gtk_buildable_custom_finished func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		child *T.GObject,
		tagname string,
		data T.Gpointer)

	Gtk_buildable_parser_finished func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder)

	Gtk_buildable_get_internal_child func(
		buildable *T.GtkBuildable,
		builder *T.GtkBuilder,
		childname string) *T.GObject

	Gtk_image_get_type func() T.GType

	Gtk_image_new func() *T.GtkWidget

	Gtk_image_new_from_pixmap func(
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap) *T.GtkWidget

	Gtk_image_new_from_image func(
		image *T.GdkImage,
		mask *T.GdkBitmap) *T.GtkWidget

	Gtk_image_new_from_file func(
		filename string) *T.GtkWidget

	Gtk_image_new_from_pixbuf func(
		pixbuf *T.GdkPixbuf) *T.GtkWidget

	Gtk_image_new_from_stock func(
		stock_id string,
		size T.GtkIconSize) *T.GtkWidget

	Gtk_image_new_from_icon_set func(
		icon_set *T.GtkIconSet,
		size T.GtkIconSize) *T.GtkWidget

	Gtk_image_new_from_animation func(
		animation *T.GdkPixbufAnimation) *T.GtkWidget

	Gtk_image_new_from_icon_name func(
		icon_name string,
		size T.GtkIconSize) *T.GtkWidget

	Gtk_image_new_from_gicon func(
		icon *T.GIcon,
		size T.GtkIconSize) *T.GtkWidget

	Gtk_image_clear func(
		image *T.GtkImage)

	Gtk_image_set_from_pixmap func(
		image *T.GtkImage,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	Gtk_image_set_from_image func(
		image *T.GtkImage,
		gdk_image *T.GdkImage,
		mask *T.GdkBitmap)

	Gtk_image_set_from_file func(
		image *T.GtkImage,
		filename string)

	Gtk_image_set_from_pixbuf func(
		image *T.GtkImage,
		pixbuf *T.GdkPixbuf)

	Gtk_image_set_from_stock func(
		image *T.GtkImage,
		stock_id string,
		size T.GtkIconSize)

	Gtk_image_set_from_icon_set func(
		image *T.GtkImage,
		icon_set *T.GtkIconSet,
		size T.GtkIconSize)

	Gtk_image_set_from_animation func(
		image *T.GtkImage,
		animation *T.GdkPixbufAnimation)

	Gtk_image_set_from_icon_name func(
		image *T.GtkImage,
		icon_name string,
		size T.GtkIconSize)

	Gtk_image_set_from_gicon func(
		image *T.GtkImage,
		icon *T.GIcon,
		size T.GtkIconSize)

	Gtk_image_set_pixel_size func(
		image *T.GtkImage,
		pixel_size T.Gint)

	Gtk_image_get_storage_type func(
		image *T.GtkImage) T.GtkImageType

	Gtk_image_get_pixmap func(
		image *T.GtkImage,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap)

	Gtk_image_get_image func(
		image *T.GtkImage,
		gdk_image **T.GdkImage,
		mask **T.GdkBitmap)

	Gtk_image_get_pixbuf func(
		image *T.GtkImage) *T.GdkPixbuf

	Gtk_image_get_stock func(
		image *T.GtkImage,
		stock_id **T.Gchar,
		size *T.GtkIconSize)

	Gtk_image_get_icon_set func(
		image *T.GtkImage,
		icon_set **T.GtkIconSet,
		size *T.GtkIconSize)

	Gtk_image_get_animation func(
		image *T.GtkImage) *T.GdkPixbufAnimation

	Gtk_image_get_icon_name func(
		image *T.GtkImage,
		icon_name **T.Gchar,
		size *T.GtkIconSize)

	Gtk_image_get_gicon func(
		image *T.GtkImage,
		gicon **T.GIcon,
		size *T.GtkIconSize)

	Gtk_image_get_pixel_size func(
		image *T.GtkImage) T.Gint

	Gtk_image_set func(
		image *T.GtkImage,
		val *T.GdkImage,
		mask *T.GdkBitmap)

	Gtk_image_get func(
		image *T.GtkImage,
		val **T.GdkImage,
		mask **T.GdkBitmap)

	Gtk_button_get_type func() T.GType

	Gtk_button_new func() *T.GtkWidget

	Gtk_button_new_with_label func(
		label string) *T.GtkWidget

	Gtk_button_new_from_stock func(
		stock_id string) *T.GtkWidget

	Gtk_button_new_with_mnemonic func(
		label string) *T.GtkWidget

	Gtk_button_pressed func(
		button *T.GtkButton)

	Gtk_button_released func(
		button *T.GtkButton)

	Gtk_button_clicked func(
		button *T.GtkButton)

	Gtk_button_enter func(
		button *T.GtkButton)

	Gtk_button_leave func(
		button *T.GtkButton)

	Gtk_button_set_relief func(
		button *T.GtkButton,
		newstyle T.GtkReliefStyle)

	Gtk_button_get_relief func(
		button *T.GtkButton) T.GtkReliefStyle

	Gtk_button_set_label func(
		button *T.GtkButton,
		label string)

	Gtk_button_get_label func(
		button *T.GtkButton) string

	Gtk_button_set_use_underline func(
		button *T.GtkButton,
		use_underline T.Gboolean)

	Gtk_button_get_use_underline func(
		button *T.GtkButton) T.Gboolean

	Gtk_button_set_use_stock func(
		button *T.GtkButton,
		use_stock T.Gboolean)

	Gtk_button_get_use_stock func(
		button *T.GtkButton) T.Gboolean

	Gtk_button_set_focus_on_click func(
		button *T.GtkButton,
		focus_on_click T.Gboolean)

	Gtk_button_get_focus_on_click func(
		button *T.GtkButton) T.Gboolean

	Gtk_button_set_alignment func(
		button *T.GtkButton,
		xalign T.Gfloat,
		yalign T.Gfloat)

	Gtk_button_get_alignment func(
		button *T.GtkButton,
		xalign *T.Gfloat,
		yalign *T.Gfloat)

	Gtk_button_set_image func(
		button *T.GtkButton,
		image *T.GtkWidget)

	Gtk_button_get_image func(
		button *T.GtkButton) *T.GtkWidget

	Gtk_button_set_image_position func(
		button *T.GtkButton,
		position T.GtkPositionType)

	Gtk_button_get_image_position func(
		button *T.GtkButton) T.GtkPositionType

	Gtk_button_get_event_window func(
		button *T.GtkButton) *T.GdkWindow

	Gtk_signal_newv func(
		name string,
		signal_flags T.GtkSignalRunType,
		object_type T.GType,
		function_offset T.Guint,
		marshaller T.GSignalCMarshaller,
		return_val T.GType,
		n_args T.Guint,
		args *T.GType) T.Guint

	Gtk_signal_new func(name string,
		signal_flags T.GtkSignalRunType,
		object_type T.GType, function_offset T.Guint,
		marshaller T.GSignalCMarshaller, return_val T.GType,
		n_args T.Guint, v ...VArg) T.Guint

	Gtk_signal_emit_stop_by_name func(
		object *T.GtkObject, name string)

	Gtk_signal_connect_object_while_alive func(
		object *T.GtkObject,
		name string,
		f T.GCallback,
		alive_object *T.GtkObject)

	Gtk_signal_connect_while_alive func(
		object *T.GtkObject,
		name string,
		f T.GCallback,
		func_dataGpointer,
		alive_object *T.GtkObject)

	Gtk_signal_connect_full func(
		object *T.GtkObject,
		name string,
		f T.GCallback,
		unsupported T.GtkCallbackMarshal,
		dataGpointer,
		destroy_func T.GDestroyNotify,
		object_signal T.Gint,
		after T.Gint) T.Gulong

	Gtk_signal_emitv func(
		object *T.GtkObject,
		signal_id T.Guint,
		args *T.GtkArg)

	Gtk_signal_emit func(object *T.GtkObject, signal_id T.Guint,
		v ...VArg)

	Gtk_signal_emit_by_name func(object *T.GtkObject,
		name string, v ...VArg)

	Gtk_signal_emitv_by_name func(
		object *T.GtkObject,
		name string,
		args *T.GtkArg)

	Gtk_signal_compat_matched func(
		object *T.GtkObject,
		f T.GCallback,
		dataGpointer,
		match T.GSignalMatchType,
		action T.Guint)

	Gtk_calendar_get_type func() T.GType

	Gtk_calendar_new func() *T.GtkWidget

	Gtk_calendar_select_month func(
		calendar *T.GtkCalendar,
		month T.Guint,
		year T.Guint) T.Gboolean

	Gtk_calendar_select_day func(
		calendar *T.GtkCalendar,
		day T.Guint)

	Gtk_calendar_mark_day func(
		calendar *T.GtkCalendar,
		day T.Guint) T.Gboolean

	Gtk_calendar_unmark_day func(
		calendar *T.GtkCalendar,
		day T.Guint) T.Gboolean

	Gtk_calendar_clear_marks func(
		calendar *T.GtkCalendar)

	Gtk_calendar_set_display_options func(
		calendar *T.GtkCalendar,
		flags T.GtkCalendarDisplayOptions)

	Gtk_calendar_get_display_options func(
		calendar *T.GtkCalendar) T.GtkCalendarDisplayOptions

	Gtk_calendar_display_options func(
		calendar *T.GtkCalendar,
		flags T.GtkCalendarDisplayOptions)

	Gtk_calendar_get_date func(
		calendar *T.GtkCalendar,
		year *T.Guint,
		month *T.Guint,
		day *T.Guint)

	Gtk_calendar_set_detail_func func(
		calendar *T.GtkCalendar,
		f T.GtkCalendarDetailFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_calendar_set_detail_width_chars func(
		calendar *T.GtkCalendar,
		chars T.Gint)

	Gtk_calendar_set_detail_height_rows func(
		calendar *T.GtkCalendar,
		rows T.Gint)

	Gtk_calendar_get_detail_width_chars func(
		calendar *T.GtkCalendar) T.Gint

	Gtk_calendar_get_detail_height_rows func(
		calendar *T.GtkCalendar) T.Gint

	Gtk_calendar_freeze func(
		calendar *T.GtkCalendar)

	Gtk_calendar_thaw func(
		calendar *T.GtkCalendar)

	Gtk_cell_editable_get_type func() T.GType

	Gtk_cell_editable_start_editing func(
		cell_editable *T.GtkCellEditable,
		event *T.GdkEvent)

	Gtk_cell_editable_editing_done func(
		cell_editable *T.GtkCellEditable)

	Gtk_cell_editable_remove_widget func(
		cell_editable *T.GtkCellEditable)

	Gtk_cell_renderer_get_type func() T.GType

	Gtk_cell_renderer_get_size func(
		cell *T.GtkCellRenderer,
		widget *T.GtkWidget,
		cell_area *T.GdkRectangle,
		x_offset *T.Gint,
		y_offset *T.Gint,
		width *T.Gint,
		height *T.Gint)

	Gtk_cell_renderer_render func(
		cell *T.GtkCellRenderer,
		window *T.GdkWindow,
		widget *T.GtkWidget,
		background_area *T.GdkRectangle,
		cell_area *T.GdkRectangle,
		expose_area *T.GdkRectangle,
		flags T.GtkCellRendererState)

	Gtk_cell_renderer_activate func(
		cell *T.GtkCellRenderer,
		event *T.GdkEvent,
		widget *T.GtkWidget,
		path string,
		background_area *T.GdkRectangle,
		cell_area *T.GdkRectangle,
		flags T.GtkCellRendererState) T.Gboolean

	Gtk_cell_renderer_start_editing func(
		cell *T.GtkCellRenderer,
		event *T.GdkEvent,
		widget *T.GtkWidget,
		path string,
		background_area *T.GdkRectangle,
		cell_area *T.GdkRectangle,
		flags T.GtkCellRendererState) *T.GtkCellEditable

	Gtk_cell_renderer_set_fixed_size func(
		cell *T.GtkCellRenderer,
		width T.Gint,
		height T.Gint)

	Gtk_cell_renderer_get_fixed_size func(
		cell *T.GtkCellRenderer,
		width *T.Gint,
		height *T.Gint)

	Gtk_cell_renderer_set_alignment func(
		cell *T.GtkCellRenderer,
		xalign T.Gfloat,
		yalign T.Gfloat)

	Gtk_cell_renderer_get_alignment func(
		cell *T.GtkCellRenderer,
		xalign *T.Gfloat,
		yalign *T.Gfloat)

	Gtk_cell_renderer_set_padding func(
		cell *T.GtkCellRenderer,
		xpad T.Gint,
		ypad T.Gint)

	Gtk_cell_renderer_get_padding func(
		cell *T.GtkCellRenderer,
		xpad *T.Gint,
		ypad *T.Gint)

	Gtk_cell_renderer_set_visible func(
		cell *T.GtkCellRenderer,
		visible T.Gboolean)

	Gtk_cell_renderer_get_visible func(
		cell *T.GtkCellRenderer) T.Gboolean

	Gtk_cell_renderer_set_sensitive func(
		cell *T.GtkCellRenderer,
		sensitive T.Gboolean)

	Gtk_cell_renderer_get_sensitive func(
		cell *T.GtkCellRenderer) T.Gboolean

	Gtk_cell_renderer_editing_canceled func(
		cell *T.GtkCellRenderer)

	Gtk_cell_renderer_stop_editing func(
		cell *T.GtkCellRenderer,
		canceled T.Gboolean)

	Gtk_tree_path_new func() *T.GtkTreePath

	Gtk_tree_path_new_from_string func(
		path string) *T.GtkTreePath

	Gtk_tree_path_new_from_indices func(first_index T.Gint,
		v ...VArg) *T.GtkTreePath

	Gtk_tree_path_to_string func(
		path *T.GtkTreePath) string

	Gtk_tree_path_new_first func() *T.GtkTreePath

	Gtk_tree_path_append_index func(
		path *T.GtkTreePath,
		index_ T.Gint)

	Gtk_tree_path_prepend_index func(
		path *T.GtkTreePath,
		index_ T.Gint)

	Gtk_tree_path_get_depth func(
		path *T.GtkTreePath) T.Gint

	Gtk_tree_path_get_indices func(
		path *T.GtkTreePath) *T.Gint

	Gtk_tree_path_get_indices_with_depth func(
		path *T.GtkTreePath,
		depth *T.Gint) *T.Gint

	Gtk_tree_path_free func(
		path *T.GtkTreePath)

	Gtk_tree_path_copy func(
		path *T.GtkTreePath) *T.GtkTreePath

	Gtk_tree_path_get_type func() T.GType

	Gtk_tree_path_compare func(
		a *T.GtkTreePath,
		b *T.GtkTreePath) T.Gint

	Gtk_tree_path_next func(
		path *T.GtkTreePath)

	Gtk_tree_path_prev func(
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_path_up func(
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_path_down func(
		path *T.GtkTreePath)

	Gtk_tree_path_is_ancestor func(
		path *T.GtkTreePath,
		descendant *T.GtkTreePath) T.Gboolean

	Gtk_tree_path_is_descendant func(
		path *T.GtkTreePath,
		ancestor *T.GtkTreePath) T.Gboolean

	Gtk_tree_row_reference_get_type func() T.GType

	Gtk_tree_row_reference_new func(
		model *T.GtkTreeModel,
		path *T.GtkTreePath) *T.GtkTreeRowReference

	Gtk_tree_row_reference_new_proxy func(
		proxy *T.GObject,
		model *T.GtkTreeModel,
		path *T.GtkTreePath) *T.GtkTreeRowReference

	Gtk_tree_row_reference_get_path func(
		reference *T.GtkTreeRowReference) *T.GtkTreePath

	Gtk_tree_row_reference_get_model func(
		reference *T.GtkTreeRowReference) *T.GtkTreeModel

	Gtk_tree_row_reference_valid func(
		reference *T.GtkTreeRowReference) T.Gboolean

	Gtk_tree_row_reference_copy func(
		reference *T.GtkTreeRowReference) *T.GtkTreeRowReference

	Gtk_tree_row_reference_free func(
		reference *T.GtkTreeRowReference)

	Gtk_tree_row_reference_inserted func(
		proxy *T.GObject,
		path *T.GtkTreePath)

	Gtk_tree_row_reference_deleted func(
		proxy *T.GObject,
		path *T.GtkTreePath)

	Gtk_tree_row_reference_reordered func(
		proxy *T.GObject,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter,
		new_order *T.Gint)

	Gtk_tree_iter_copy func(
		iter *T.GtkTreeIter) *T.GtkTreeIter

	Gtk_tree_iter_free func(
		iter *T.GtkTreeIter)

	Gtk_tree_iter_get_type func() T.GType

	Gtk_tree_model_get_type func() T.GType

	Gtk_tree_model_get_flags func(
		tree_model *T.GtkTreeModel) T.GtkTreeModelFlags

	Gtk_tree_model_get_n_columns func(
		tree_model *T.GtkTreeModel) T.Gint

	Gtk_tree_model_get_column_type func(
		tree_model *T.GtkTreeModel,
		index_ T.Gint) T.GType

	Gtk_tree_model_get_iter func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_model_get_iter_from_string func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		path_string string) T.Gboolean

	Gtk_tree_model_get_string_from_iter func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter) string

	Gtk_tree_model_get_iter_first func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_model_get_path func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter) *T.GtkTreePath

	Gtk_tree_model_get_value func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		column T.Gint,
		value *T.GValue)

	Gtk_tree_model_iter_next func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_model_iter_children func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter) T.Gboolean

	Gtk_tree_model_iter_has_child func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_model_iter_n_children func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gint

	Gtk_tree_model_iter_nth_child func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		n T.Gint) T.Gboolean

	Gtk_tree_model_iter_parent func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		child *T.GtkTreeIter) T.Gboolean

	Gtk_tree_model_ref_node func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter)

	Gtk_tree_model_unref_node func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter)

	Gtk_tree_model_get func(tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter, v ...VArg)

	Gtk_tree_model_get_valist func(
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		var_args T.Va_list)

	Gtk_tree_model_foreach func(
		model *T.GtkTreeModel,
		f T.GtkTreeModelForeachFunc,
		user_data T.Gpointer)

	Gtk_tree_model_row_changed func(
		tree_model *T.GtkTreeModel,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter)

	Gtk_tree_model_row_inserted func(
		tree_model *T.GtkTreeModel,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter)

	Gtk_tree_model_row_has_child_toggled func(
		tree_model *T.GtkTreeModel,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter)

	Gtk_tree_model_row_deleted func(
		tree_model *T.GtkTreeModel,
		path *T.GtkTreePath)

	Gtk_tree_model_rows_reordered func(
		tree_model *T.GtkTreeModel,
		path *T.GtkTreePath,
		iter *T.GtkTreeIter,
		new_order *T.Gint)

	Gtk_tree_sortable_get_type func() T.GType

	Gtk_tree_sortable_sort_column_changed func(
		sortable *T.GtkTreeSortable)

	Gtk_tree_sortable_get_sort_column_id func(
		sortable *T.GtkTreeSortable,
		sort_column_id *T.Gint,
		order *T.GtkSortType) T.Gboolean

	Gtk_tree_sortable_set_sort_column_id func(
		sortable *T.GtkTreeSortable,
		sort_column_id T.Gint,
		order T.GtkSortType)

	Gtk_tree_sortable_set_sort_func func(
		sortable *T.GtkTreeSortable,
		sort_column_id T.Gint,
		sort_func T.GtkTreeIterCompareFunc,
		user_dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_sortable_set_default_sort_func func(
		sortable *T.GtkTreeSortable,
		sort_func T.GtkTreeIterCompareFunc,
		user_dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_sortable_has_default_sort_func func(
		sortable *T.GtkTreeSortable) T.Gboolean

	Gtk_tree_view_column_get_type func() T.GType

	Gtk_tree_view_column_new func() *T.GtkTreeViewColumn

	Gtk_tree_view_column_new_with_attributes func(
		title string, cell *T.GtkCellRenderer,
		v ...VArg) *T.GtkTreeViewColumn

	Gtk_tree_view_column_pack_start func(
		tree_column *T.GtkTreeViewColumn,
		cell *T.GtkCellRenderer,
		expand T.Gboolean)

	Gtk_tree_view_column_pack_end func(
		tree_column *T.GtkTreeViewColumn,
		cell *T.GtkCellRenderer,
		expand T.Gboolean)

	Gtk_tree_view_column_clear func(
		tree_column *T.GtkTreeViewColumn)

	Gtk_tree_view_column_get_cell_renderers func(
		tree_column *T.GtkTreeViewColumn) *T.GList

	Gtk_tree_view_column_add_attribute func(
		tree_column *T.GtkTreeViewColumn,
		cell_renderer *T.GtkCellRenderer,
		attribute string,
		column T.Gint)

	Gtk_tree_view_column_set_attributes func(
		tree_column *T.GtkTreeViewColumn,
		cell_renderer *T.GtkCellRenderer, v ...VArg)

	Gtk_tree_view_column_set_cell_data_func func(
		tree_column *T.GtkTreeViewColumn,
		cell_renderer *T.GtkCellRenderer,
		f T.GtkTreeCellDataFunc,
		func_dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_view_column_clear_attributes func(
		tree_column *T.GtkTreeViewColumn,
		cell_renderer *T.GtkCellRenderer)

	Gtk_tree_view_column_set_spacing func(
		tree_column *T.GtkTreeViewColumn,
		spacing T.Gint)

	Gtk_tree_view_column_get_spacing func(
		tree_column *T.GtkTreeViewColumn) T.Gint

	Gtk_tree_view_column_set_visible func(
		tree_column *T.GtkTreeViewColumn,
		visible T.Gboolean)

	Gtk_tree_view_column_get_visible func(
		tree_column *T.GtkTreeViewColumn) T.Gboolean

	Gtk_tree_view_column_set_resizable func(
		tree_column *T.GtkTreeViewColumn,
		resizable T.Gboolean)

	Gtk_tree_view_column_get_resizable func(
		tree_column *T.GtkTreeViewColumn) T.Gboolean

	Gtk_tree_view_column_set_sizing func(
		tree_column *T.GtkTreeViewColumn,
		typ T.GtkTreeViewColumnSizing)

	Gtk_tree_view_column_get_sizing func(
		tree_column *T.GtkTreeViewColumn) T.GtkTreeViewColumnSizing

	Gtk_tree_view_column_get_width func(
		tree_column *T.GtkTreeViewColumn) T.Gint

	Gtk_tree_view_column_get_fixed_width func(
		tree_column *T.GtkTreeViewColumn) T.Gint

	Gtk_tree_view_column_set_fixed_width func(
		tree_column *T.GtkTreeViewColumn,
		fixed_width T.Gint)

	Gtk_tree_view_column_set_min_width func(
		tree_column *T.GtkTreeViewColumn,
		min_width T.Gint)

	Gtk_tree_view_column_get_min_width func(
		tree_column *T.GtkTreeViewColumn) T.Gint

	Gtk_tree_view_column_set_max_width func(
		tree_column *T.GtkTreeViewColumn,
		max_width T.Gint)

	Gtk_tree_view_column_get_max_width func(
		tree_column *T.GtkTreeViewColumn) T.Gint

	Gtk_tree_view_column_clicked func(
		tree_column *T.GtkTreeViewColumn)

	Gtk_tree_view_column_set_title func(
		tree_column *T.GtkTreeViewColumn,
		title string)

	Gtk_tree_view_column_get_title func(
		tree_column *T.GtkTreeViewColumn) string

	Gtk_tree_view_column_set_expand func(
		tree_column *T.GtkTreeViewColumn,
		expand T.Gboolean)

	Gtk_tree_view_column_get_expand func(
		tree_column *T.GtkTreeViewColumn) T.Gboolean

	Gtk_tree_view_column_set_clickable func(
		tree_column *T.GtkTreeViewColumn,
		clickable T.Gboolean)

	Gtk_tree_view_column_get_clickable func(
		tree_column *T.GtkTreeViewColumn) T.Gboolean

	Gtk_tree_view_column_set_widget func(
		tree_column *T.GtkTreeViewColumn,
		widget *T.GtkWidget)

	Gtk_tree_view_column_get_widget func(
		tree_column *T.GtkTreeViewColumn) *T.GtkWidget

	Gtk_tree_view_column_set_alignment func(
		tree_column *T.GtkTreeViewColumn,
		xalign T.Gfloat)

	Gtk_tree_view_column_get_alignment func(
		tree_column *T.GtkTreeViewColumn) T.Gfloat

	Gtk_tree_view_column_set_reorderable func(
		tree_column *T.GtkTreeViewColumn,
		reorderable T.Gboolean)

	Gtk_tree_view_column_get_reorderable func(
		tree_column *T.GtkTreeViewColumn) T.Gboolean

	Gtk_tree_view_column_set_sort_column_id func(
		tree_column *T.GtkTreeViewColumn,
		sort_column_id T.Gint)

	Gtk_tree_view_column_get_sort_column_id func(
		tree_column *T.GtkTreeViewColumn) T.Gint

	Gtk_tree_view_column_set_sort_indicator func(
		tree_column *T.GtkTreeViewColumn,
		setting T.Gboolean)

	Gtk_tree_view_column_get_sort_indicator func(
		tree_column *T.GtkTreeViewColumn) T.Gboolean

	Gtk_tree_view_column_set_sort_order func(
		tree_column *T.GtkTreeViewColumn,
		order T.GtkSortType)

	Gtk_tree_view_column_get_sort_order func(
		tree_column *T.GtkTreeViewColumn) T.GtkSortType

	Gtk_tree_view_column_cell_set_cell_data func(
		tree_column *T.GtkTreeViewColumn,
		tree_model *T.GtkTreeModel,
		iter *T.GtkTreeIter,
		is_expander T.Gboolean,
		is_expanded T.Gboolean)

	Gtk_tree_view_column_cell_get_size func(
		tree_column *T.GtkTreeViewColumn,
		cell_area *T.GdkRectangle,
		x_offset *T.Gint,
		y_offset *T.Gint,
		width *T.Gint,
		height *T.Gint)

	Gtk_tree_view_column_cell_is_visible func(
		tree_column *T.GtkTreeViewColumn) T.Gboolean

	Gtk_tree_view_column_focus_cell func(
		tree_column *T.GtkTreeViewColumn,
		cell *T.GtkCellRenderer)

	Gtk_tree_view_column_cell_get_position func(
		tree_column *T.GtkTreeViewColumn,
		cell_renderer *T.GtkCellRenderer,
		start_pos *T.Gint,
		width *T.Gint) T.Gboolean

	Gtk_tree_view_column_queue_resize func(
		tree_column *T.GtkTreeViewColumn)

	Gtk_tree_view_column_get_tree_view func(
		tree_column *T.GtkTreeViewColumn) *T.GtkWidget

	Gtk_cell_layout_get_type func() T.GType

	Gtk_cell_layout_pack_start func(
		cell_layout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		expand T.Gboolean)

	Gtk_cell_layout_pack_end func(
		cell_layout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		expand T.Gboolean)

	Gtk_cell_layout_get_cells func(
		cell_layout *T.GtkCellLayout) *T.GList

	Gtk_cell_layout_clear func(
		cell_layout *T.GtkCellLayout)

	Gtk_cell_layout_set_attributes func(cell_layout *T.GtkCellLayout,
		cell *T.GtkCellRenderer, v ...VArg)

	Gtk_cell_layout_add_attribute func(
		cell_layout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		attribute string,
		column T.Gint)

	Gtk_cell_layout_set_cell_data_func func(
		cell_layout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		f T.GtkCellLayoutDataFunc,
		func_dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_cell_layout_clear_attributes func(
		cell_layout *T.GtkCellLayout,
		cell *T.GtkCellRenderer)

	Gtk_cell_layout_reorder func(
		cell_layout *T.GtkCellLayout,
		cell *T.GtkCellRenderer,
		position T.Gint)

	Gtk_cell_renderer_text_get_type func() T.GType

	Gtk_cell_renderer_text_new func() *T.GtkCellRenderer

	Gtk_cell_renderer_text_set_fixed_height_from_font func(
		renderer *T.GtkCellRendererText,
		number_of_rows T.Gint)

	Gtk_cell_renderer_accel_get_type func() T.GType

	Gtk_cell_renderer_accel_new func() *T.GtkCellRenderer

	Gtk_cell_renderer_combo_get_type func() T.GType

	Gtk_cell_renderer_combo_new func() *T.GtkCellRenderer

	Gtk_cell_renderer_pixbuf_get_type func() T.GType

	Gtk_cell_renderer_pixbuf_new func() *T.GtkCellRenderer

	Gtk_cell_renderer_progress_get_type func() T.GType

	Gtk_cell_renderer_progress_new func() *T.GtkCellRenderer

	Gtk_cell_renderer_spin_get_type func() T.GType

	Gtk_cell_renderer_spin_new func() *T.GtkCellRenderer

	Gtk_cell_renderer_spinner_get_type func() T.GType

	Gtk_cell_renderer_spinner_new func() *T.GtkCellRenderer

	Gtk_cell_renderer_toggle_get_type func() T.GType

	Gtk_cell_renderer_toggle_new func() *T.GtkCellRenderer

	Gtk_cell_renderer_toggle_get_radio func(
		toggle *T.GtkCellRendererToggle) T.Gboolean

	Gtk_cell_renderer_toggle_set_radio func(
		toggle *T.GtkCellRendererToggle,
		radio T.Gboolean)

	Gtk_cell_renderer_toggle_get_active func(
		toggle *T.GtkCellRendererToggle) T.Gboolean

	Gtk_cell_renderer_toggle_set_active func(
		toggle *T.GtkCellRendererToggle,
		setting T.Gboolean)

	Gtk_cell_renderer_toggle_get_activatable func(
		toggle *T.GtkCellRendererToggle) T.Gboolean

	Gtk_cell_renderer_toggle_set_activatable func(
		toggle *T.GtkCellRendererToggle,
		setting T.Gboolean)

	Gtk_cell_view_get_type func() T.GType

	Gtk_cell_view_new func() *T.GtkWidget

	Gtk_cell_view_new_with_text func(
		text string) *T.GtkWidget

	Gtk_cell_view_new_with_markup func(
		markup string) *T.GtkWidget

	Gtk_cell_view_new_with_pixbuf func(
		pixbuf *T.GdkPixbuf) *T.GtkWidget

	Gtk_cell_view_set_model func(
		cell_view *T.GtkCellView,
		model *T.GtkTreeModel)

	Gtk_cell_view_get_model func(
		cell_view *T.GtkCellView) *T.GtkTreeModel

	Gtk_cell_view_set_displayed_row func(
		cell_view *T.GtkCellView,
		path *T.GtkTreePath)

	Gtk_cell_view_get_displayed_row func(
		cell_view *T.GtkCellView) *T.GtkTreePath

	Gtk_cell_view_get_size_of_row func(
		cell_view *T.GtkCellView,
		path *T.GtkTreePath,
		requisition *T.GtkRequisition) T.Gboolean

	Gtk_cell_view_set_background_color func(
		cell_view *T.GtkCellView,
		color *T.GdkColor)

	Gtk_cell_view_get_cell_renderers func(
		cell_view *T.GtkCellView) *T.GList

	Gtk_toggle_button_get_type func() T.GType

	Gtk_toggle_button_new func() *T.GtkWidget

	Gtk_toggle_button_new_with_label func(
		label string) *T.GtkWidget

	Gtk_toggle_button_new_with_mnemonic func(
		label string) *T.GtkWidget

	Gtk_toggle_button_set_mode func(
		toggle_button *T.GtkToggleButton,
		draw_indicator T.Gboolean)

	Gtk_toggle_button_get_mode func(
		toggle_button *T.GtkToggleButton) T.Gboolean

	Gtk_toggle_button_set_active func(
		toggle_button *T.GtkToggleButton,
		is_active T.Gboolean)

	Gtk_toggle_button_get_active func(
		toggle_button *T.GtkToggleButton) T.Gboolean

	Gtk_toggle_button_toggled func(
		toggle_button *T.GtkToggleButton)

	Gtk_toggle_button_set_inconsistent func(
		toggle_button *T.GtkToggleButton,
		setting T.Gboolean)

	Gtk_toggle_button_get_inconsistent func(
		toggle_button *T.GtkToggleButton) T.Gboolean

	Gtk_check_button_get_type func() T.GType

	Gtk_check_button_new func() *T.GtkWidget

	Gtk_check_button_new_with_label func(
		label string) *T.GtkWidget

	Gtk_check_button_new_with_mnemonic func(
		label string) *T.GtkWidget

	Gtk_item_get_type func() T.GType

	Gtk_item_select func(
		item *T.GtkItem)

	Gtk_item_deselect func(
		item *T.GtkItem)

	Gtk_item_toggle func(
		item *T.GtkItem)

	Gtk_menu_item_get_type func() T.GType

	Gtk_menu_item_new func() *T.GtkWidget

	Gtk_menu_item_new_with_label func(
		label string) *T.GtkWidget

	Gtk_menu_item_new_with_mnemonic func(
		label string) *T.GtkWidget

	Gtk_menu_item_set_submenu func(
		menu_item *T.GtkMenuItem,
		submenu *T.GtkWidget)

	Gtk_menu_item_get_submenu func(
		menu_item *T.GtkMenuItem) *T.GtkWidget

	Gtk_menu_item_select func(
		menu_item *T.GtkMenuItem)

	Gtk_menu_item_deselect func(
		menu_item *T.GtkMenuItem)

	Gtk_menu_item_activate func(
		menu_item *T.GtkMenuItem)

	Gtk_menu_item_toggle_size_request func(
		menu_item *T.GtkMenuItem,
		requisition *T.Gint)

	Gtk_menu_item_toggle_size_allocate func(
		menu_item *T.GtkMenuItem,
		allocation T.Gint)

	Gtk_menu_item_set_right_justified func(
		menu_item *T.GtkMenuItem,
		right_justified T.Gboolean)

	Gtk_menu_item_get_right_justified func(
		menu_item *T.GtkMenuItem) T.Gboolean

	Gtk_menu_item_set_accel_path func(
		menu_item *T.GtkMenuItem,
		accel_path string)

	Gtk_menu_item_get_accel_path func(
		menu_item *T.GtkMenuItem) string

	Gtk_menu_item_set_label func(
		menu_item *T.GtkMenuItem,
		label string)

	Gtk_menu_item_get_label func(
		menu_item *T.GtkMenuItem) string

	Gtk_menu_item_set_use_underline func(
		menu_item *T.GtkMenuItem,
		setting T.Gboolean)

	Gtk_menu_item_get_use_underline func(
		menu_item *T.GtkMenuItem) T.Gboolean

	Gtk_menu_item_remove_submenu func(
		menu_item *T.GtkMenuItem)

	Gtk_check_menu_item_get_type func() T.GType

	Gtk_check_menu_item_new func() *T.GtkWidget

	Gtk_check_menu_item_new_with_label func(
		label string) *T.GtkWidget

	Gtk_check_menu_item_new_with_mnemonic func(
		label string) *T.GtkWidget

	Gtk_check_menu_item_set_active func(
		check_menu_item *T.GtkCheckMenuItem,
		is_active T.Gboolean)

	Gtk_check_menu_item_get_active func(
		check_menu_item *T.GtkCheckMenuItem) T.Gboolean

	Gtk_check_menu_item_toggled func(
		check_menu_item *T.GtkCheckMenuItem)

	Gtk_check_menu_item_set_inconsistent func(
		check_menu_item *T.GtkCheckMenuItem,
		setting T.Gboolean)

	Gtk_check_menu_item_get_inconsistent func(
		check_menu_item *T.GtkCheckMenuItem) T.Gboolean

	Gtk_check_menu_item_set_draw_as_radio func(
		check_menu_item *T.GtkCheckMenuItem,
		draw_as_radio T.Gboolean)

	Gtk_check_menu_item_get_draw_as_radio func(
		check_menu_item *T.GtkCheckMenuItem) T.Gboolean

	Gtk_check_menu_item_set_show_toggle func(
		menu_item *T.GtkCheckMenuItem,
		always T.Gboolean)

	Gtk_text_tag_get_type func() T.GType

	Gtk_text_tag_new func(
		name string) *T.GtkTextTag

	Gtk_text_tag_get_priority func(
		tag *T.GtkTextTag) T.Gint

	Gtk_text_tag_set_priority func(
		tag *T.GtkTextTag,
		priority T.Gint)

	Gtk_text_tag_event func(
		tag *T.GtkTextTag,
		event_object *T.GObject,
		event *T.GdkEvent,
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_attributes_new func() *T.GtkTextAttributes

	Gtk_text_attributes_copy func(
		src *T.GtkTextAttributes) *T.GtkTextAttributes

	Gtk_text_attributes_copy_values func(
		src *T.GtkTextAttributes,
		dest *T.GtkTextAttributes)

	Gtk_text_attributes_unref func(
		values *T.GtkTextAttributes)

	Gtk_text_attributes_ref func(
		values *T.GtkTextAttributes) *T.GtkTextAttributes

	Gtk_text_attributes_get_type func() T.GType

	Gtk_text_child_anchor_get_type func() T.GType

	Gtk_text_child_anchor_new func() *T.GtkTextChildAnchor

	Gtk_text_child_anchor_get_widgets func(
		anchor *T.GtkTextChildAnchor) *T.GList

	Gtk_text_child_anchor_get_deleted func(
		anchor *T.GtkTextChildAnchor) T.Gboolean

	Gtk_text_iter_get_buffer func(
		iter *T.GtkTextIter) *T.GtkTextBuffer

	Gtk_text_iter_copy func(
		iter *T.GtkTextIter) *T.GtkTextIter

	Gtk_text_iter_free func(
		iter *T.GtkTextIter)

	Gtk_text_iter_get_type func() T.GType

	Gtk_text_iter_get_offset func(
		iter *T.GtkTextIter) T.Gint

	Gtk_text_iter_get_line func(
		iter *T.GtkTextIter) T.Gint

	Gtk_text_iter_get_line_offset func(
		iter *T.GtkTextIter) T.Gint

	Gtk_text_iter_get_line_index func(
		iter *T.GtkTextIter) T.Gint

	Gtk_text_iter_get_visible_line_offset func(
		iter *T.GtkTextIter) T.Gint

	Gtk_text_iter_get_visible_line_index func(
		iter *T.GtkTextIter) T.Gint

	Gtk_text_iter_get_char func(
		iter *T.GtkTextIter) T.Gunichar

	Gtk_text_iter_get_slice func(
		start *T.GtkTextIter,
		end *T.GtkTextIter) string

	Gtk_text_iter_get_text func(
		start *T.GtkTextIter,
		end *T.GtkTextIter) string

	Gtk_text_iter_get_visible_slice func(
		start *T.GtkTextIter,
		end *T.GtkTextIter) string

	Gtk_text_iter_get_visible_text func(
		start *T.GtkTextIter,
		end *T.GtkTextIter) string

	Gtk_text_iter_get_pixbuf func(
		iter *T.GtkTextIter) *T.GdkPixbuf

	Gtk_text_iter_get_marks func(
		iter *T.GtkTextIter) *T.GSList

	Gtk_text_iter_get_child_anchor func(
		iter *T.GtkTextIter) *T.GtkTextChildAnchor

	Gtk_text_iter_get_toggled_tags func(
		iter *T.GtkTextIter,
		toggled_on T.Gboolean) *T.GSList

	Gtk_text_iter_begins_tag func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	Gtk_text_iter_ends_tag func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	Gtk_text_iter_toggles_tag func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	Gtk_text_iter_has_tag func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	Gtk_text_iter_get_tags func(
		iter *T.GtkTextIter) *T.GSList

	Gtk_text_iter_editable func(
		iter *T.GtkTextIter,
		default_setting T.Gboolean) T.Gboolean

	Gtk_text_iter_can_insert func(
		iter *T.GtkTextIter,
		default_editability T.Gboolean) T.Gboolean

	Gtk_text_iter_starts_word func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_ends_word func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_inside_word func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_starts_sentence func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_ends_sentence func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_inside_sentence func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_starts_line func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_ends_line func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_is_cursor_position func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_get_chars_in_line func(
		iter *T.GtkTextIter) T.Gint

	Gtk_text_iter_get_bytes_in_line func(
		iter *T.GtkTextIter) T.Gint

	Gtk_text_iter_get_attributes func(
		iter *T.GtkTextIter,
		values *T.GtkTextAttributes) T.Gboolean

	Gtk_text_iter_get_language func(
		iter *T.GtkTextIter) *T.PangoLanguage

	Gtk_text_iter_is_end func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_is_start func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_char func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_char func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_chars func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_backward_chars func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_forward_line func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_line func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_lines func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_backward_lines func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_forward_word_end func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_word_start func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_word_ends func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_backward_word_starts func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_forward_visible_line func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_visible_line func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_visible_lines func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_backward_visible_lines func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_forward_visible_word_end func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_visible_word_start func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_visible_word_ends func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_backward_visible_word_starts func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_forward_sentence_end func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_sentence_start func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_sentence_ends func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_backward_sentence_starts func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_forward_cursor_position func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_cursor_position func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_cursor_positions func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_backward_cursor_positions func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_forward_visible_cursor_position func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_visible_cursor_position func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_visible_cursor_positions func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_backward_visible_cursor_positions func(
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_iter_set_offset func(
		iter *T.GtkTextIter,
		char_offset T.Gint)

	Gtk_text_iter_set_line func(
		iter *T.GtkTextIter,
		line_number T.Gint)

	Gtk_text_iter_set_line_offset func(
		iter *T.GtkTextIter,
		char_on_line T.Gint)

	Gtk_text_iter_set_line_index func(
		iter *T.GtkTextIter,
		byte_on_line T.Gint)

	Gtk_text_iter_forward_to_end func(
		iter *T.GtkTextIter)

	Gtk_text_iter_forward_to_line_end func(
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_set_visible_line_offset func(
		iter *T.GtkTextIter,
		char_on_line T.Gint)

	Gtk_text_iter_set_visible_line_index func(
		iter *T.GtkTextIter,
		byte_on_line T.Gint)

	Gtk_text_iter_forward_to_tag_toggle func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	Gtk_text_iter_backward_to_tag_toggle func(
		iter *T.GtkTextIter,
		tag *T.GtkTextTag) T.Gboolean

	Gtk_text_iter_forward_find_char func(
		iter *T.GtkTextIter,
		pred T.GtkTextCharPredicate,
		user_dataGpointer,
		limit *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_find_char func(
		iter *T.GtkTextIter,
		pred T.GtkTextCharPredicate,
		user_dataGpointer,
		limit *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_forward_search func(
		iter *T.GtkTextIter,
		str string,
		flags T.GtkTextSearchFlags,
		match_start *T.GtkTextIter,
		match_end *T.GtkTextIter,
		limit *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_backward_search func(
		iter *T.GtkTextIter,
		str string,
		flags T.GtkTextSearchFlags,
		match_start *T.GtkTextIter,
		match_end *T.GtkTextIter,
		limit *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_equal func(
		lhs *T.GtkTextIter,
		rhs *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_compare func(
		lhs *T.GtkTextIter,
		rhs *T.GtkTextIter) T.Gint

	Gtk_text_iter_in_range func(
		iter *T.GtkTextIter,
		start *T.GtkTextIter,
		end *T.GtkTextIter) T.Gboolean

	Gtk_text_iter_order func(
		first *T.GtkTextIter,
		second *T.GtkTextIter)

	Gtk_target_list_new func(
		targets *T.GtkTargetEntry,
		ntargets T.Guint) *T.GtkTargetList

	Gtk_target_list_ref func(
		list *T.GtkTargetList) *T.GtkTargetList

	Gtk_target_list_unref func(
		list *T.GtkTargetList)

	Gtk_target_list_add func(
		list *T.GtkTargetList,
		target T.GdkAtom,
		flags T.Guint,
		info T.Guint)

	Gtk_target_list_add_text_targets func(
		list *T.GtkTargetList,
		info T.Guint)

	Gtk_target_list_add_rich_text_targets func(
		list *T.GtkTargetList,
		info T.Guint,
		deserializable T.Gboolean,
		buffer *T.GtkTextBuffer)

	Gtk_target_list_add_image_targets func(
		list *T.GtkTargetList,
		info T.Guint,
		writable T.Gboolean)

	Gtk_target_list_add_uri_targets func(
		list *T.GtkTargetList,
		info T.Guint)

	Gtk_target_list_add_table func(
		list *T.GtkTargetList,
		targets *T.GtkTargetEntry,
		ntargets T.Guint)

	Gtk_target_list_remove func(
		list *T.GtkTargetList,
		target T.GdkAtom)

	Gtk_target_list_find func(
		list *T.GtkTargetList,
		target T.GdkAtom,
		info *T.Guint) T.Gboolean

	Gtk_target_table_new_from_list func(
		list *T.GtkTargetList,
		n_targets *T.Gint) *T.GtkTargetEntry

	Gtk_target_table_free func(
		targets *T.GtkTargetEntry,
		n_targets T.Gint)

	Gtk_selection_owner_set func(
		widget *T.GtkWidget,
		selection T.GdkAtom,
		time_ T.Guint32) T.Gboolean

	Gtk_selection_owner_set_for_display func(
		display *T.GdkDisplay,
		widget *T.GtkWidget,
		selection T.GdkAtom,
		time_ T.Guint32) T.Gboolean

	Gtk_selection_add_target func(
		widget *T.GtkWidget,
		selection T.GdkAtom,
		target T.GdkAtom,
		info T.Guint)

	Gtk_selection_add_targets func(
		widget *T.GtkWidget,
		selection T.GdkAtom,
		targets *T.GtkTargetEntry,
		ntargets T.Guint)

	Gtk_selection_clear_targets func(
		widget *T.GtkWidget,
		selection T.GdkAtom)

	Gtk_selection_convert func(
		widget *T.GtkWidget,
		selection T.GdkAtom,
		target T.GdkAtom,
		time_ T.Guint32) T.Gboolean

	Gtk_selection_data_get_selection func(
		selection_data *T.GtkSelectionData) T.GdkAtom

	Gtk_selection_data_get_target func(
		selection_data *T.GtkSelectionData) T.GdkAtom

	Gtk_selection_data_get_data_type func(
		selection_data *T.GtkSelectionData) T.GdkAtom

	Gtk_selection_data_get_format func(
		selection_data *T.GtkSelectionData) T.Gint

	Gtk_selection_data_get_data func(
		selection_data *T.GtkSelectionData) *T.Guchar

	Gtk_selection_data_get_length func(
		selection_data *T.GtkSelectionData) T.Gint

	Gtk_selection_data_get_display func(
		selection_data *T.GtkSelectionData) *T.GdkDisplay

	Gtk_selection_data_set func(
		selection_data *T.GtkSelectionData,
		typ T.GdkAtom,
		format T.Gint,
		data *T.Guchar,
		length T.Gint)

	Gtk_selection_data_set_text func(
		selection_data *T.GtkSelectionData,
		str string,
		len T.Gint) T.Gboolean

	Gtk_selection_data_get_text func(
		selection_data *T.GtkSelectionData) *T.Guchar

	Gtk_selection_data_set_pixbuf func(
		selection_data *T.GtkSelectionData,
		pixbuf *T.GdkPixbuf) T.Gboolean

	Gtk_selection_data_get_pixbuf func(
		selection_data *T.GtkSelectionData) *T.GdkPixbuf

	Gtk_selection_data_set_uris func(
		selection_data *T.GtkSelectionData,
		uris **T.Gchar) T.Gboolean

	Gtk_selection_data_get_uris func(
		selection_data *T.GtkSelectionData) **T.Gchar

	Gtk_selection_data_get_targets func(
		selection_data *T.GtkSelectionData,
		targets **T.GdkAtom,
		n_atoms *T.Gint) T.Gboolean

	Gtk_selection_data_targets_include_text func(
		selection_data *T.GtkSelectionData) T.Gboolean

	Gtk_selection_data_targets_include_rich_text func(
		selection_data *T.GtkSelectionData,
		buffer *T.GtkTextBuffer) T.Gboolean

	Gtk_selection_data_targets_include_image func(
		selection_data *T.GtkSelectionData,
		writable T.Gboolean) T.Gboolean

	Gtk_selection_data_targets_include_uri func(
		selection_data *T.GtkSelectionData) T.Gboolean

	Gtk_targets_include_text func(
		targets *T.GdkAtom,
		n_targets T.Gint) T.Gboolean

	Gtk_targets_include_rich_text func(
		targets *T.GdkAtom,
		n_targets T.Gint,
		buffer *T.GtkTextBuffer) T.Gboolean

	Gtk_targets_include_image func(
		targets *T.GdkAtom,
		n_targets T.Gint,
		writable T.Gboolean) T.Gboolean

	Gtk_targets_include_uri func(
		targets *T.GdkAtom,
		n_targets T.Gint) T.Gboolean

	Gtk_selection_remove_all func(
		widget *T.GtkWidget)

	Gtk_selection_clear func(
		widget *T.GtkWidget,
		event *T.GdkEventSelection) T.Gboolean

	Gtk_selection_data_get_type func() T.GType

	Gtk_selection_data_copy func(
		data *T.GtkSelectionData) *T.GtkSelectionData

	Gtk_selection_data_free func(
		data *T.GtkSelectionData)

	Gtk_target_list_get_type func() T.GType

	Gtk_clipboard_get_type func() T.GType

	Gtk_clipboard_get_for_display func(
		display *T.GdkDisplay,
		selection T.GdkAtom) *T.GtkClipboard

	Gtk_clipboard_get func(
		selection T.GdkAtom) *T.GtkClipboard

	Gtk_clipboard_get_display func(
		clipboard *T.GtkClipboard) *T.GdkDisplay

	Gtk_clipboard_set_with_data func(
		clipboard *T.GtkClipboard,
		targets *T.GtkTargetEntry,
		n_targets T.Guint,
		get_func T.GtkClipboardGetFunc,
		clear_func T.GtkClipboardClearFunc,
		user_data T.Gpointer) T.Gboolean

	Gtk_clipboard_set_with_owner func(
		clipboard *T.GtkClipboard,
		targets *T.GtkTargetEntry,
		n_targets T.Guint,
		get_func T.GtkClipboardGetFunc,
		clear_func T.GtkClipboardClearFunc,
		owner *T.GObject) T.Gboolean

	Gtk_clipboard_get_owner func(
		clipboard *T.GtkClipboard) *T.GObject

	Gtk_clipboard_clear func(
		clipboard *T.GtkClipboard)

	Gtk_clipboard_set_text func(
		clipboard *T.GtkClipboard,
		text string,
		len T.Gint)

	Gtk_clipboard_set_image func(
		clipboard *T.GtkClipboard,
		pixbuf *T.GdkPixbuf)

	Gtk_clipboard_request_contents func(
		clipboard *T.GtkClipboard,
		target T.GdkAtom,
		callback T.GtkClipboardReceivedFunc,
		user_data T.Gpointer)

	Gtk_clipboard_request_text func(
		clipboard *T.GtkClipboard,
		callback T.GtkClipboardTextReceivedFunc,
		user_data T.Gpointer)

	Gtk_clipboard_request_rich_text func(
		clipboard *T.GtkClipboard,
		buffer *T.GtkTextBuffer,
		callback T.GtkClipboardRichTextReceivedFunc,
		user_data T.Gpointer)

	Gtk_clipboard_request_image func(
		clipboard *T.GtkClipboard,
		callback T.GtkClipboardImageReceivedFunc,
		user_data T.Gpointer)

	Gtk_clipboard_request_uris func(
		clipboard *T.GtkClipboard,
		callback T.GtkClipboardURIReceivedFunc,
		user_data T.Gpointer)

	Gtk_clipboard_request_targets func(
		clipboard *T.GtkClipboard,
		callback T.GtkClipboardTargetsReceivedFunc,
		user_data T.Gpointer)

	Gtk_clipboard_wait_for_contents func(
		clipboard *T.GtkClipboard,
		target T.GdkAtom) *T.GtkSelectionData

	Gtk_clipboard_wait_for_text func(
		clipboard *T.GtkClipboard) string

	Gtk_clipboard_wait_for_rich_text func(
		clipboard *T.GtkClipboard,
		buffer *T.GtkTextBuffer,
		format *T.GdkAtom,
		length *T.Gsize) *T.Guint8

	Gtk_clipboard_wait_for_image func(
		clipboard *T.GtkClipboard) *T.GdkPixbuf

	Gtk_clipboard_wait_for_uris func(
		clipboard *T.GtkClipboard) **T.Gchar

	Gtk_clipboard_wait_for_targets func(
		clipboard *T.GtkClipboard,
		targets **T.GdkAtom,
		n_targets *T.Gint) T.Gboolean

	Gtk_clipboard_wait_is_text_available func(
		clipboard *T.GtkClipboard) T.Gboolean

	Gtk_clipboard_wait_is_rich_text_available func(
		clipboard *T.GtkClipboard,
		buffer *T.GtkTextBuffer) T.Gboolean

	Gtk_clipboard_wait_is_image_available func(
		clipboard *T.GtkClipboard) T.Gboolean

	Gtk_clipboard_wait_is_uris_available func(
		clipboard *T.GtkClipboard) T.Gboolean

	Gtk_clipboard_wait_is_target_available func(
		clipboard *T.GtkClipboard,
		target T.GdkAtom) T.Gboolean

	Gtk_clipboard_set_can_store func(
		clipboard *T.GtkClipboard,
		targets *T.GtkTargetEntry,
		n_targets T.Gint)

	Gtk_clipboard_store func(
		clipboard *T.GtkClipboard)

	Gtk_color_button_get_type func() T.GType

	Gtk_color_button_new func() *T.GtkWidget

	Gtk_color_button_new_with_color func(
		color *T.GdkColor) *T.GtkWidget

	Gtk_color_button_set_color func(
		color_button *T.GtkColorButton,
		color *T.GdkColor)

	Gtk_color_button_set_alpha func(
		color_button *T.GtkColorButton,
		alpha T.Guint16)

	Gtk_color_button_get_color func(
		color_button *T.GtkColorButton,
		color *T.GdkColor)

	Gtk_color_button_get_alpha func(
		color_button *T.GtkColorButton) T.Guint16

	Gtk_color_button_set_use_alpha func(
		color_button *T.GtkColorButton,
		use_alpha T.Gboolean)

	Gtk_color_button_get_use_alpha func(
		color_button *T.GtkColorButton) T.Gboolean

	Gtk_color_button_set_title func(
		color_button *T.GtkColorButton,
		title string)

	Gtk_color_button_get_title func(
		color_button *T.GtkColorButton) string

	Gtk_vbox_get_type func() T.GType

	Gtk_vbox_new func(
		homogeneous T.Gboolean,
		spacing T.Gint) *T.GtkWidget

	Gtk_color_selection_get_type func() T.GType

	Gtk_color_selection_new func() *T.GtkWidget

	Gtk_color_selection_get_has_opacity_control func(
		colorsel *T.GtkColorSelection) T.Gboolean

	Gtk_color_selection_set_has_opacity_control func(
		colorsel *T.GtkColorSelection,
		has_opacity T.Gboolean)

	Gtk_color_selection_get_has_palette func(
		colorsel *T.GtkColorSelection) T.Gboolean

	Gtk_color_selection_set_has_palette func(
		colorsel *T.GtkColorSelection,
		has_palette T.Gboolean)

	Gtk_color_selection_set_current_color func(
		colorsel *T.GtkColorSelection,
		color *T.GdkColor)

	Gtk_color_selection_set_current_alpha func(
		colorsel *T.GtkColorSelection,
		alpha T.Guint16)

	Gtk_color_selection_get_current_color func(
		colorsel *T.GtkColorSelection,
		color *T.GdkColor)

	Gtk_color_selection_get_current_alpha func(
		colorsel *T.GtkColorSelection) T.Guint16

	Gtk_color_selection_set_previous_color func(
		colorsel *T.GtkColorSelection,
		color *T.GdkColor)

	Gtk_color_selection_set_previous_alpha func(
		colorsel *T.GtkColorSelection,
		alpha T.Guint16)

	Gtk_color_selection_get_previous_color func(
		colorsel *T.GtkColorSelection,
		color *T.GdkColor)

	Gtk_color_selection_get_previous_alpha func(
		colorsel *T.GtkColorSelection) T.Guint16

	Gtk_color_selection_is_adjusting func(
		colorsel *T.GtkColorSelection) T.Gboolean

	Gtk_color_selection_palette_from_string func(
		str string,
		colors **T.GdkColor,
		n_colors *T.Gint) T.Gboolean

	Gtk_color_selection_palette_to_string func(
		colors *T.GdkColor,
		n_colors T.Gint) string

	Gtk_color_selection_set_change_palette_hook func(
		f T.GtkColorSelectionChangePaletteFunc) T.GtkColorSelectionChangePaletteFunc

	Gtk_color_selection_set_change_palette_with_screen_hook func(
		f T.GtkColorSelectionChangePaletteWithScreenFunc) T.GtkColorSelectionChangePaletteWithScreenFunc

	Gtk_color_selection_set_color func(
		colorsel *T.GtkColorSelection,
		color *T.Gdouble)

	Gtk_color_selection_get_color func(
		colorsel *T.GtkColorSelection,
		color *T.Gdouble)

	Gtk_color_selection_set_update_policy func(
		colorsel *T.GtkColorSelection,
		policy T.GtkUpdateType)

	Gtk_color_selection_dialog_get_type func() T.GType

	Gtk_color_selection_dialog_new func(
		title string) *T.GtkWidget

	Gtk_color_selection_dialog_get_color_selection func(
		colorsel *T.GtkColorSelectionDialog) *T.GtkWidget

	Gtk_drag_get_data func(
		widget *T.GtkWidget,
		context *T.GdkDragContext,
		target T.GdkAtom,
		time_ T.Guint32)

	Gtk_drag_finish func(
		context *T.GdkDragContext,
		success T.Gboolean,
		del T.Gboolean,
		time_ T.Guint32)

	Gtk_drag_get_source_widget func(
		context *T.GdkDragContext) *T.GtkWidget

	Gtk_drag_highlight func(
		widget *T.GtkWidget)

	Gtk_drag_unhighlight func(
		widget *T.GtkWidget)

	Gtk_drag_dest_set func(
		widget *T.GtkWidget,
		flags T.GtkDestDefaults,
		targets *T.GtkTargetEntry,
		n_targets T.Gint,
		actions T.GdkDragAction)

	Gtk_drag_dest_set_proxy func(
		widget *T.GtkWidget,
		proxy_window *T.GdkWindow,
		protocol T.GdkDragProtocol,
		use_coordinates T.Gboolean)

	Gtk_drag_dest_unset func(
		widget *T.GtkWidget)

	Gtk_drag_dest_find_target func(
		widget *T.GtkWidget,
		context *T.GdkDragContext,
		target_list *T.GtkTargetList) T.GdkAtom

	Gtk_drag_dest_get_target_list func(
		widget *T.GtkWidget) *T.GtkTargetList

	Gtk_drag_dest_set_target_list func(
		widget *T.GtkWidget,
		target_list *T.GtkTargetList)

	Gtk_drag_dest_add_text_targets func(
		widget *T.GtkWidget)

	Gtk_drag_dest_add_image_targets func(
		widget *T.GtkWidget)

	Gtk_drag_dest_add_uri_targets func(
		widget *T.GtkWidget)

	Gtk_drag_dest_set_track_motion func(
		widget *T.GtkWidget,
		track_motion T.Gboolean)

	Gtk_drag_dest_get_track_motion func(
		widget *T.GtkWidget) T.Gboolean

	Gtk_drag_source_set func(
		widget *T.GtkWidget,
		start_button_mask T.GdkModifierType,
		targets *T.GtkTargetEntry,
		n_targets T.Gint,
		actions T.GdkDragAction)

	Gtk_drag_source_unset func(
		widget *T.GtkWidget)

	Gtk_drag_source_get_target_list func(
		widget *T.GtkWidget) *T.GtkTargetList

	Gtk_drag_source_set_target_list func(
		widget *T.GtkWidget,
		target_list *T.GtkTargetList)

	Gtk_drag_source_add_text_targets func(
		widget *T.GtkWidget)

	Gtk_drag_source_add_image_targets func(
		widget *T.GtkWidget)

	Gtk_drag_source_add_uri_targets func(
		widget *T.GtkWidget)

	Gtk_drag_source_set_icon func(
		widget *T.GtkWidget,
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	Gtk_drag_source_set_icon_pixbuf func(
		widget *T.GtkWidget,
		pixbuf *T.GdkPixbuf)

	Gtk_drag_source_set_icon_stock func(
		widget *T.GtkWidget,
		stock_id string)

	Gtk_drag_source_set_icon_name func(
		widget *T.GtkWidget,
		icon_name string)

	Gtk_drag_begin func(
		widget *T.GtkWidget,
		targets *T.GtkTargetList,
		actions T.GdkDragAction,
		button T.Gint,
		event *T.GdkEvent) *T.GdkDragContext

	Gtk_drag_set_icon_widget func(
		context *T.GdkDragContext,
		widget *T.GtkWidget,
		hot_x T.Gint,
		hot_y T.Gint)

	Gtk_drag_set_icon_pixmap func(
		context *T.GdkDragContext,
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap,
		hot_x T.Gint,
		hot_y T.Gint)

	Gtk_drag_set_icon_pixbuf func(
		context *T.GdkDragContext,
		pixbuf *T.GdkPixbuf,
		hot_x T.Gint,
		hot_y T.Gint)

	Gtk_drag_set_icon_stock func(
		context *T.GdkDragContext,
		stock_id string,
		hot_x T.Gint,
		hot_y T.Gint)

	Gtk_drag_set_icon_name func(
		context *T.GdkDragContext,
		icon_name string,
		hot_x T.Gint,
		hot_y T.Gint)

	Gtk_drag_set_icon_default func(
		context *T.GdkDragContext)

	Gtk_drag_check_threshold func(
		widget *T.GtkWidget,
		start_x T.Gint,
		start_y T.Gint,
		current_x T.Gint,
		current_y T.Gint) T.Gboolean

	Gtk_drag_set_default_icon func(
		colormap *T.GdkColormap,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap,
		hot_x T.Gint,
		hot_y T.Gint)

	Gtk_editable_get_type func() T.GType

	Gtk_editable_select_region func(
		editable *T.GtkEditable,
		start_pos T.Gint,
		end_pos T.Gint)

	Gtk_editable_get_selection_bounds func(
		editable *T.GtkEditable,
		start_pos *T.Gint,
		end_pos *T.Gint) T.Gboolean

	Gtk_editable_insert_text func(
		editable *T.GtkEditable,
		new_text string,
		new_text_length T.Gint,
		position *T.Gint)

	Gtk_editable_delete_text func(
		editable *T.GtkEditable,
		start_pos T.Gint,
		end_pos T.Gint)

	Gtk_editable_get_chars func(
		editable *T.GtkEditable,
		start_pos T.Gint,
		end_pos T.Gint) string

	Gtk_editable_cut_clipboard func(
		editable *T.GtkEditable)

	Gtk_editable_copy_clipboard func(
		editable *T.GtkEditable)

	Gtk_editable_paste_clipboard func(
		editable *T.GtkEditable)

	Gtk_editable_delete_selection func(
		editable *T.GtkEditable)

	Gtk_editable_set_position func(
		editable *T.GtkEditable,
		position T.Gint)

	Gtk_editable_get_position func(
		editable *T.GtkEditable) T.Gint

	Gtk_editable_set_editable func(
		editable *T.GtkEditable,
		is_editable T.Gboolean)

	Gtk_editable_get_editable func(
		editable *T.GtkEditable) T.Gboolean

	Gtk_im_context_get_type func() T.GType

	Gtk_im_context_set_client_window func(
		context *T.GtkIMContext,
		window *T.GdkWindow)

	Gtk_im_context_get_preedit_string func(
		context *T.GtkIMContext,
		str **T.Gchar,
		attrs **T.PangoAttrList,
		cursor_pos *T.Gint)

	Gtk_im_context_filter_keypress func(
		context *T.GtkIMContext,
		event *T.GdkEventKey) T.Gboolean

	Gtk_im_context_focus_in func(
		context *T.GtkIMContext)

	Gtk_im_context_focus_out func(
		context *T.GtkIMContext)

	Gtk_im_context_reset func(
		context *T.GtkIMContext)

	Gtk_im_context_set_cursor_location func(
		context *T.GtkIMContext,
		area *T.GdkRectangle)

	Gtk_im_context_set_use_preedit func(
		context *T.GtkIMContext,
		use_preedit T.Gboolean)

	Gtk_im_context_set_surrounding func(
		context *T.GtkIMContext,
		text string,
		len T.Gint,
		cursor_index T.Gint)

	Gtk_im_context_get_surrounding func(
		context *T.GtkIMContext,
		text **T.Gchar,
		cursor_index *T.Gint) T.Gboolean

	Gtk_im_context_delete_surrounding func(
		context *T.GtkIMContext,
		offset T.Gint,
		n_chars T.Gint) T.Gboolean

	Gtk_entry_buffer_get_type func() T.GType

	Gtk_entry_buffer_new func(
		initial_chars string,
		n_initial_chars T.Gint) *T.GtkEntryBuffer

	Gtk_entry_buffer_get_bytes func(
		buffer *T.GtkEntryBuffer) T.Gsize

	Gtk_entry_buffer_get_length func(
		buffer *T.GtkEntryBuffer) T.Guint

	Gtk_entry_buffer_get_text func(
		buffer *T.GtkEntryBuffer) string

	Gtk_entry_buffer_set_text func(
		buffer *T.GtkEntryBuffer,
		chars string,
		n_chars T.Gint)

	Gtk_entry_buffer_set_max_length func(
		buffer *T.GtkEntryBuffer,
		max_length T.Gint)

	Gtk_entry_buffer_get_max_length func(
		buffer *T.GtkEntryBuffer) T.Gint

	Gtk_entry_buffer_insert_text func(
		buffer *T.GtkEntryBuffer,
		position T.Guint,
		chars string,
		n_chars T.Gint) T.Guint

	Gtk_entry_buffer_delete_text func(
		buffer *T.GtkEntryBuffer,
		position T.Guint,
		n_chars T.Gint) T.Guint

	Gtk_entry_buffer_emit_inserted_text func(
		buffer *T.GtkEntryBuffer,
		position T.Guint,
		chars string,
		n_chars T.Guint)

	Gtk_entry_buffer_emit_deleted_text func(
		buffer *T.GtkEntryBuffer,
		position T.Guint,
		n_chars T.Guint)

	Gtk_list_store_get_type func() T.GType

	Gtk_list_store_new func(n_columns T.Gint,
		v ...VArg) *T.GtkListStore

	Gtk_list_store_newv func(
		n_columns T.Gint,
		types *T.GType) *T.GtkListStore

	Gtk_list_store_set_column_types func(
		list_store *T.GtkListStore,
		n_columns T.Gint,
		types *T.GType)

	Gtk_list_store_set_value func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter,
		column T.Gint,
		value *T.GValue)

	Gtk_list_store_set func(list_store *T.GtkListStore,
		iter *T.GtkTreeIter, v ...VArg)

	Gtk_list_store_set_valuesv func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter,
		columns *T.Gint,
		values *T.GValue,
		n_values T.Gint)

	Gtk_list_store_set_valist func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter,
		var_args T.Va_list)

	Gtk_list_store_remove func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_list_store_insert func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter,
		position T.Gint)

	Gtk_list_store_insert_before func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter,
		sibling *T.GtkTreeIter)

	Gtk_list_store_insert_after func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter,
		sibling *T.GtkTreeIter)

	Gtk_list_store_insert_with_values func(list_store *T.GtkListStore,
		iter *T.GtkTreeIter, position T.Gint, v ...VArg)

	Gtk_list_store_insert_with_valuesv func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter,
		position T.Gint,
		columns *T.Gint,
		values *T.GValue,
		n_values T.Gint)

	Gtk_list_store_prepend func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter)

	Gtk_list_store_append func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter)

	Gtk_list_store_clear func(
		list_store *T.GtkListStore)

	Gtk_list_store_iter_is_valid func(
		list_store *T.GtkListStore,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_list_store_reorder func(
		store *T.GtkListStore,
		new_order *T.Gint)

	Gtk_list_store_swap func(
		store *T.GtkListStore,
		a *T.GtkTreeIter,
		b *T.GtkTreeIter)

	Gtk_list_store_move_after func(
		store *T.GtkListStore,
		iter *T.GtkTreeIter,
		position *T.GtkTreeIter)

	Gtk_list_store_move_before func(
		store *T.GtkListStore,
		iter *T.GtkTreeIter,
		position *T.GtkTreeIter)

	Gtk_tree_model_filter_get_type func() T.GType

	Gtk_tree_model_filter_new func(
		child_model *T.GtkTreeModel,
		root *T.GtkTreePath) *T.GtkTreeModel

	Gtk_tree_model_filter_set_visible_func func(
		filter *T.GtkTreeModelFilter,
		f T.GtkTreeModelFilterVisibleFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_model_filter_set_modify_func func(
		filter *T.GtkTreeModelFilter,
		n_columns T.Gint,
		types *T.GType,
		f T.GtkTreeModelFilterModifyFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_model_filter_set_visible_column func(
		filter *T.GtkTreeModelFilter,
		column T.Gint)

	Gtk_tree_model_filter_get_model func(
		filter *T.GtkTreeModelFilter) *T.GtkTreeModel

	Gtk_tree_model_filter_convert_child_iter_to_iter func(
		filter *T.GtkTreeModelFilter,
		filter_iter *T.GtkTreeIter,
		child_iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_model_filter_convert_iter_to_child_iter func(
		filter *T.GtkTreeModelFilter,
		child_iter *T.GtkTreeIter,
		filter_iter *T.GtkTreeIter)

	Gtk_tree_model_filter_convert_child_path_to_path func(
		filter *T.GtkTreeModelFilter,
		child_path *T.GtkTreePath) *T.GtkTreePath

	Gtk_tree_model_filter_convert_path_to_child_path func(
		filter *T.GtkTreeModelFilter,
		filter_path *T.GtkTreePath) *T.GtkTreePath

	Gtk_tree_model_filter_refilter func(
		filter *T.GtkTreeModelFilter)

	Gtk_tree_model_filter_clear_cache func(
		filter *T.GtkTreeModelFilter)

	Gtk_entry_completion_get_type func() T.GType

	Gtk_entry_completion_new func() *T.GtkEntryCompletion

	Gtk_entry_completion_get_entry func(
		completion *T.GtkEntryCompletion) *T.GtkWidget

	Gtk_entry_completion_set_model func(
		completion *T.GtkEntryCompletion,
		model *T.GtkTreeModel)

	Gtk_entry_completion_get_model func(
		completion *T.GtkEntryCompletion) *T.GtkTreeModel

	Gtk_entry_completion_set_match_func func(
		completion *T.GtkEntryCompletion,
		f T.GtkEntryCompletionMatchFunc,
		func_dataGpointer,
		func_notify T.GDestroyNotify)

	Gtk_entry_completion_set_minimum_key_length func(
		completion *T.GtkEntryCompletion,
		length T.Gint)

	Gtk_entry_completion_get_minimum_key_length func(
		completion *T.GtkEntryCompletion) T.Gint

	Gtk_entry_completion_complete func(
		completion *T.GtkEntryCompletion)

	Gtk_entry_completion_insert_prefix func(
		completion *T.GtkEntryCompletion)

	Gtk_entry_completion_insert_action_text func(
		completion *T.GtkEntryCompletion,
		index_ T.Gint,
		text string)

	Gtk_entry_completion_insert_action_markup func(
		completion *T.GtkEntryCompletion,
		index_ T.Gint,
		markup string)

	Gtk_entry_completion_delete_action func(
		completion *T.GtkEntryCompletion,
		index_ T.Gint)

	Gtk_entry_completion_set_inline_completion func(
		completion *T.GtkEntryCompletion,
		inline_completion T.Gboolean)

	Gtk_entry_completion_get_inline_completion func(
		completion *T.GtkEntryCompletion) T.Gboolean

	Gtk_entry_completion_set_inline_selection func(
		completion *T.GtkEntryCompletion,
		inline_selection T.Gboolean)

	Gtk_entry_completion_get_inline_selection func(
		completion *T.GtkEntryCompletion) T.Gboolean

	Gtk_entry_completion_set_popup_completion func(
		completion *T.GtkEntryCompletion,
		popup_completion T.Gboolean)

	Gtk_entry_completion_get_popup_completion func(
		completion *T.GtkEntryCompletion) T.Gboolean

	Gtk_entry_completion_set_popup_set_width func(
		completion *T.GtkEntryCompletion,
		popup_set_width T.Gboolean)

	Gtk_entry_completion_get_popup_set_width func(
		completion *T.GtkEntryCompletion) T.Gboolean

	Gtk_entry_completion_set_popup_single_match func(
		completion *T.GtkEntryCompletion,
		popup_single_match T.Gboolean)

	Gtk_entry_completion_get_popup_single_match func(
		completion *T.GtkEntryCompletion) T.Gboolean

	Gtk_entry_completion_get_completion_prefix func(
		completion *T.GtkEntryCompletion) string

	Gtk_entry_completion_set_text_column func(
		completion *T.GtkEntryCompletion,
		column T.Gint)

	Gtk_entry_completion_get_text_column func(
		completion *T.GtkEntryCompletion) T.Gint

	Gtk_entry_get_type func() T.GType

	Gtk_entry_new func() *T.GtkWidget

	Gtk_entry_new_with_buffer func(
		buffer *T.GtkEntryBuffer) *T.GtkWidget

	Gtk_entry_get_buffer func(
		entry *T.GtkEntry) *T.GtkEntryBuffer

	Gtk_entry_set_buffer func(
		entry *T.GtkEntry,
		buffer *T.GtkEntryBuffer)

	Gtk_entry_get_text_window func(
		entry *T.GtkEntry) *T.GdkWindow

	Gtk_entry_set_visibility func(
		entry *T.GtkEntry,
		visible T.Gboolean)

	Gtk_entry_get_visibility func(
		entry *T.GtkEntry) T.Gboolean

	Gtk_entry_set_invisible_char func(
		entry *T.GtkEntry,
		ch T.Gunichar)

	Gtk_entry_get_invisible_char func(
		entry *T.GtkEntry) T.Gunichar

	Gtk_entry_unset_invisible_char func(
		entry *T.GtkEntry)

	Gtk_entry_set_has_frame func(
		entry *T.GtkEntry,
		setting T.Gboolean)

	Gtk_entry_get_has_frame func(
		entry *T.GtkEntry) T.Gboolean

	Gtk_entry_set_inner_border func(
		entry *T.GtkEntry,
		border *T.GtkBorder)

	Gtk_entry_get_inner_border func(
		entry *T.GtkEntry) *T.GtkBorder

	Gtk_entry_set_overwrite_mode func(
		entry *T.GtkEntry,
		overwrite T.Gboolean)

	Gtk_entry_get_overwrite_mode func(
		entry *T.GtkEntry) T.Gboolean

	Gtk_entry_set_max_length func(
		entry *T.GtkEntry,
		max T.Gint)

	Gtk_entry_get_max_length func(
		entry *T.GtkEntry) T.Gint

	Gtk_entry_get_text_length func(
		entry *T.GtkEntry) T.Guint16

	Gtk_entry_set_activates_default func(
		entry *T.GtkEntry,
		setting T.Gboolean)

	Gtk_entry_get_activates_default func(
		entry *T.GtkEntry) T.Gboolean

	Gtk_entry_set_width_chars func(
		entry *T.GtkEntry,
		n_chars T.Gint)

	Gtk_entry_get_width_chars func(
		entry *T.GtkEntry) T.Gint

	Gtk_entry_set_text func(
		entry *T.GtkEntry,
		text string)

	Gtk_entry_get_text func(
		entry *T.GtkEntry) string

	Gtk_entry_get_layout func(
		entry *T.GtkEntry) *T.PangoLayout

	Gtk_entry_get_layout_offsets func(
		entry *T.GtkEntry,
		x *T.Gint,
		y *T.Gint)

	Gtk_entry_set_alignment func(
		entry *T.GtkEntry,
		xalign T.Gfloat)

	Gtk_entry_get_alignment func(
		entry *T.GtkEntry) T.Gfloat

	Gtk_entry_set_completion func(
		entry *T.GtkEntry,
		completion *T.GtkEntryCompletion)

	Gtk_entry_get_completion func(
		entry *T.GtkEntry) *T.GtkEntryCompletion

	Gtk_entry_layout_index_to_text_index func(
		entry *T.GtkEntry,
		layout_index T.Gint) T.Gint

	Gtk_entry_text_index_to_layout_index func(
		entry *T.GtkEntry,
		text_index T.Gint) T.Gint

	Gtk_entry_set_cursor_hadjustment func(
		entry *T.GtkEntry,
		adjustment *T.GtkAdjustment)

	Gtk_entry_get_cursor_hadjustment func(
		entry *T.GtkEntry) *T.GtkAdjustment

	Gtk_entry_set_progress_fraction func(
		entry *T.GtkEntry,
		fraction T.Gdouble)

	Gtk_entry_get_progress_fraction func(
		entry *T.GtkEntry) T.Gdouble

	Gtk_entry_set_progress_pulse_step func(
		entry *T.GtkEntry,
		fraction T.Gdouble)

	Gtk_entry_get_progress_pulse_step func(
		entry *T.GtkEntry) T.Gdouble

	Gtk_entry_progress_pulse func(
		entry *T.GtkEntry)

	Gtk_entry_set_icon_from_pixbuf func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		pixbuf *T.GdkPixbuf)

	Gtk_entry_set_icon_from_stock func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		stock_id string)

	Gtk_entry_set_icon_from_icon_name func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		icon_name string)

	Gtk_entry_set_icon_from_gicon func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		icon *T.GIcon)

	Gtk_entry_get_icon_storage_type func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) T.GtkImageType

	Gtk_entry_get_icon_pixbuf func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) *T.GdkPixbuf

	Gtk_entry_get_icon_stock func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) string

	Gtk_entry_get_icon_name func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) string

	Gtk_entry_get_icon_gicon func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) *T.GIcon

	Gtk_entry_set_icon_activatable func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		activatable T.Gboolean)

	Gtk_entry_get_icon_activatable func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) T.Gboolean

	Gtk_entry_set_icon_sensitive func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		sensitive T.Gboolean)

	Gtk_entry_get_icon_sensitive func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) T.Gboolean

	Gtk_entry_get_icon_at_pos func(
		entry *T.GtkEntry,
		x T.Gint,
		y T.Gint) T.Gint

	Gtk_entry_set_icon_tooltip_text func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		tooltip string)

	Gtk_entry_get_icon_tooltip_text func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) string

	Gtk_entry_set_icon_tooltip_markup func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		tooltip string)

	Gtk_entry_get_icon_tooltip_markup func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) string

	Gtk_entry_set_icon_drag_source func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition,
		target_list *T.GtkTargetList,
		actions T.GdkDragAction)

	Gtk_entry_get_current_icon_drag_source func(
		entry *T.GtkEntry) T.Gint

	Gtk_entry_get_icon_window func(
		entry *T.GtkEntry,
		icon_pos T.GtkEntryIconPosition) *T.GdkWindow

	Gtk_entry_im_context_filter_keypress func(
		entry *T.GtkEntry,
		event *T.GdkEventKey) T.Gboolean

	Gtk_entry_reset_im_context func(
		entry *T.GtkEntry)

	Gtk_entry_new_with_max_length func(
		max T.Gint) *T.GtkWidget

	Gtk_entry_append_text func(
		entry *T.GtkEntry,
		text string)

	Gtk_entry_prepend_text func(
		entry *T.GtkEntry,
		text string)

	Gtk_entry_set_position func(
		entry *T.GtkEntry,
		position T.Gint)

	Gtk_entry_select_region func(
		entry *T.GtkEntry,
		start T.Gint,
		end T.Gint)

	Gtk_entry_set_editable func(
		entry *T.GtkEntry,
		editable T.Gboolean)

	Gtk_tree_view_get_type func() T.GType

	Gtk_tree_view_new func() *T.GtkWidget

	Gtk_tree_view_new_with_model func(
		model *T.GtkTreeModel) *T.GtkWidget

	Gtk_tree_view_get_model func(
		tree_view *T.GtkTreeView) *T.GtkTreeModel

	Gtk_tree_view_set_model func(
		tree_view *T.GtkTreeView,
		model *T.GtkTreeModel)

	Gtk_tree_view_get_selection func(
		tree_view *T.GtkTreeView) *T.GtkTreeSelection

	Gtk_tree_view_get_hadjustment func(
		tree_view *T.GtkTreeView) *T.GtkAdjustment

	Gtk_tree_view_set_hadjustment func(
		tree_view *T.GtkTreeView,
		adjustment *T.GtkAdjustment)

	Gtk_tree_view_get_vadjustment func(
		tree_view *T.GtkTreeView) *T.GtkAdjustment

	Gtk_tree_view_set_vadjustment func(
		tree_view *T.GtkTreeView,
		adjustment *T.GtkAdjustment)

	Gtk_tree_view_get_headers_visible func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_set_headers_visible func(
		tree_view *T.GtkTreeView,
		headers_visible T.Gboolean)

	Gtk_tree_view_columns_autosize func(
		tree_view *T.GtkTreeView)

	Gtk_tree_view_get_headers_clickable func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_set_headers_clickable func(
		tree_view *T.GtkTreeView,
		setting T.Gboolean)

	Gtk_tree_view_set_rules_hint func(
		tree_view *T.GtkTreeView,
		setting T.Gboolean)

	Gtk_tree_view_get_rules_hint func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_append_column func(
		tree_view *T.GtkTreeView,
		column *T.GtkTreeViewColumn) T.Gint

	Gtk_tree_view_remove_column func(
		tree_view *T.GtkTreeView,
		column *T.GtkTreeViewColumn) T.Gint

	Gtk_tree_view_insert_column func(
		tree_view *T.GtkTreeView,
		column *T.GtkTreeViewColumn,
		position T.Gint) T.Gint

	Gtk_tree_view_insert_column_with_attributes func(
		tree_view *T.GtkTreeView, position T.Gint, title string,
		cell *T.GtkCellRenderer, v ...VArg) T.Gint

	Gtk_tree_view_insert_column_with_data_func func(
		tree_view *T.GtkTreeView,
		position T.Gint,
		title string,
		cell *T.GtkCellRenderer,
		f T.GtkTreeCellDataFunc,
		dataGpointer,
		dnotify T.GDestroyNotify) T.Gint

	Gtk_tree_view_get_column func(
		tree_view *T.GtkTreeView,
		n T.Gint) *T.GtkTreeViewColumn

	Gtk_tree_view_get_columns func(
		tree_view *T.GtkTreeView) *T.GList

	Gtk_tree_view_move_column_after func(
		tree_view *T.GtkTreeView,
		column *T.GtkTreeViewColumn,
		base_column *T.GtkTreeViewColumn)

	Gtk_tree_view_set_expander_column func(
		tree_view *T.GtkTreeView,
		column *T.GtkTreeViewColumn)

	Gtk_tree_view_get_expander_column func(
		tree_view *T.GtkTreeView) *T.GtkTreeViewColumn

	Gtk_tree_view_set_column_drag_function func(
		tree_view *T.GtkTreeView,
		f T.GtkTreeViewColumnDropFunc,
		user_dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_view_scroll_to_point func(
		tree_view *T.GtkTreeView,
		tree_x T.Gint,
		tree_y T.Gint)

	Gtk_tree_view_scroll_to_cell func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath,
		column *T.GtkTreeViewColumn,
		use_align T.Gboolean,
		row_align T.Gfloat,
		col_align T.Gfloat)

	Gtk_tree_view_row_activated func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath,
		column *T.GtkTreeViewColumn)

	Gtk_tree_view_expand_all func(
		tree_view *T.GtkTreeView)

	Gtk_tree_view_collapse_all func(
		tree_view *T.GtkTreeView)

	Gtk_tree_view_expand_to_path func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath)

	Gtk_tree_view_expand_row func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath,
		open_all T.Gboolean) T.Gboolean

	Gtk_tree_view_collapse_row func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_view_map_expanded_rows func(
		tree_view *T.GtkTreeView,
		f T.GtkTreeViewMappingFunc,
		data T.Gpointer)

	Gtk_tree_view_row_expanded func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_view_set_reorderable func(
		tree_view *T.GtkTreeView,
		reorderable T.Gboolean)

	Gtk_tree_view_get_reorderable func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_set_cursor func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath,
		focus_column *T.GtkTreeViewColumn,
		start_editing T.Gboolean)

	Gtk_tree_view_set_cursor_on_cell func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath,
		focus_column *T.GtkTreeViewColumn,
		focus_cell *T.GtkCellRenderer,
		start_editing T.Gboolean)

	Gtk_tree_view_get_cursor func(
		tree_view *T.GtkTreeView,
		path **T.GtkTreePath,
		focus_column **T.GtkTreeViewColumn)

	Gtk_tree_view_get_bin_window func(
		tree_view *T.GtkTreeView) *T.GdkWindow

	Gtk_tree_view_get_path_at_pos func(
		tree_view *T.GtkTreeView,
		x T.Gint,
		y T.Gint,
		path **T.GtkTreePath,
		column **T.GtkTreeViewColumn,
		cell_x *T.Gint,
		cell_y *T.Gint) T.Gboolean

	Gtk_tree_view_get_cell_area func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath,
		column *T.GtkTreeViewColumn,
		rect *T.GdkRectangle)

	Gtk_tree_view_get_background_area func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath,
		column *T.GtkTreeViewColumn,
		rect *T.GdkRectangle)

	Gtk_tree_view_get_visible_rect func(
		tree_view *T.GtkTreeView,
		visible_rect *T.GdkRectangle)

	Gtk_tree_view_widget_to_tree_coords func(
		tree_view *T.GtkTreeView,
		wx T.Gint,
		wy T.Gint,
		tx *T.Gint,
		ty *T.Gint)

	Gtk_tree_view_tree_to_widget_coords func(
		tree_view *T.GtkTreeView,
		tx T.Gint,
		ty T.Gint,
		wx *T.Gint,
		wy *T.Gint)

	Gtk_tree_view_get_visible_range func(
		tree_view *T.GtkTreeView,
		start_path **T.GtkTreePath,
		end_path **T.GtkTreePath) T.Gboolean

	Gtk_tree_view_enable_model_drag_source func(
		tree_view *T.GtkTreeView,
		start_button_mask T.GdkModifierType,
		targets *T.GtkTargetEntry,
		n_targets T.Gint,
		actions T.GdkDragAction)

	Gtk_tree_view_enable_model_drag_dest func(
		tree_view *T.GtkTreeView,
		targets *T.GtkTargetEntry,
		n_targets T.Gint,
		actions T.GdkDragAction)

	Gtk_tree_view_unset_rows_drag_source func(
		tree_view *T.GtkTreeView)

	Gtk_tree_view_unset_rows_drag_dest func(
		tree_view *T.GtkTreeView)

	Gtk_tree_view_set_drag_dest_row func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath,
		pos T.GtkTreeViewDropPosition)

	Gtk_tree_view_get_drag_dest_row func(
		tree_view *T.GtkTreeView,
		path **T.GtkTreePath,
		pos *T.GtkTreeViewDropPosition)

	Gtk_tree_view_get_dest_row_at_pos func(
		tree_view *T.GtkTreeView,
		drag_x T.Gint,
		drag_y T.Gint,
		path **T.GtkTreePath,
		pos *T.GtkTreeViewDropPosition) T.Gboolean

	Gtk_tree_view_create_row_drag_icon func(
		tree_view *T.GtkTreeView,
		path *T.GtkTreePath) *T.GdkPixmap

	Gtk_tree_view_set_enable_search func(
		tree_view *T.GtkTreeView,
		enable_search T.Gboolean)

	Gtk_tree_view_get_enable_search func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_get_search_column func(
		tree_view *T.GtkTreeView) T.Gint

	Gtk_tree_view_set_search_column func(
		tree_view *T.GtkTreeView,
		column T.Gint)

	Gtk_tree_view_get_search_equal_func func(
		tree_view *T.GtkTreeView) T.GtkTreeViewSearchEqualFunc

	Gtk_tree_view_set_search_equal_func func(
		tree_view *T.GtkTreeView,
		search_equal_func T.GtkTreeViewSearchEqualFunc,
		search_user_dataGpointer,
		search_destroy T.GDestroyNotify)

	Gtk_tree_view_get_search_entry func(
		tree_view *T.GtkTreeView) *T.GtkEntry

	Gtk_tree_view_set_search_entry func(
		tree_view *T.GtkTreeView,
		entry *T.GtkEntry)

	Gtk_tree_view_get_search_position_func func(
		tree_view *T.GtkTreeView) T.GtkTreeViewSearchPositionFunc

	Gtk_tree_view_set_search_position_func func(
		tree_view *T.GtkTreeView,
		f T.GtkTreeViewSearchPositionFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_view_convert_widget_to_tree_coords func(
		tree_view *T.GtkTreeView,
		wx T.Gint,
		wy T.Gint,
		tx *T.Gint,
		ty *T.Gint)

	Gtk_tree_view_convert_tree_to_widget_coords func(
		tree_view *T.GtkTreeView,
		tx T.Gint,
		ty T.Gint,
		wx *T.Gint,
		wy *T.Gint)

	Gtk_tree_view_convert_widget_to_bin_window_coords func(
		tree_view *T.GtkTreeView,
		wx T.Gint,
		wy T.Gint,
		bx *T.Gint,
		by *T.Gint)

	Gtk_tree_view_convert_bin_window_to_widget_coords func(
		tree_view *T.GtkTreeView,
		bx T.Gint,
		by T.Gint,
		wx *T.Gint,
		wy *T.Gint)

	Gtk_tree_view_convert_tree_to_bin_window_coords func(
		tree_view *T.GtkTreeView,
		tx T.Gint,
		ty T.Gint,
		bx *T.Gint,
		by *T.Gint)

	Gtk_tree_view_convert_bin_window_to_tree_coords func(
		tree_view *T.GtkTreeView,
		bx T.Gint,
		by T.Gint,
		tx *T.Gint,
		ty *T.Gint)

	Gtk_tree_view_set_destroy_count_func func(
		tree_view *T.GtkTreeView,
		f T.GtkTreeDestroyCountFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_view_set_fixed_height_mode func(
		tree_view *T.GtkTreeView,
		enable T.Gboolean)

	Gtk_tree_view_get_fixed_height_mode func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_set_hover_selection func(
		tree_view *T.GtkTreeView,
		hover T.Gboolean)

	Gtk_tree_view_get_hover_selection func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_set_hover_expand func(
		tree_view *T.GtkTreeView,
		expand T.Gboolean)

	Gtk_tree_view_get_hover_expand func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_set_rubber_banding func(
		tree_view *T.GtkTreeView,
		enable T.Gboolean)

	Gtk_tree_view_get_rubber_banding func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_is_rubber_banding_active func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_get_row_separator_func func(
		tree_view *T.GtkTreeView) T.GtkTreeViewRowSeparatorFunc

	Gtk_tree_view_set_row_separator_func func(
		tree_view *T.GtkTreeView,
		f T.GtkTreeViewRowSeparatorFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_view_get_grid_lines func(
		tree_view *T.GtkTreeView) T.GtkTreeViewGridLines

	Gtk_tree_view_set_grid_lines func(
		tree_view *T.GtkTreeView,
		grid_lines T.GtkTreeViewGridLines)

	Gtk_tree_view_get_enable_tree_lines func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_set_enable_tree_lines func(
		tree_view *T.GtkTreeView,
		enabled T.Gboolean)

	Gtk_tree_view_set_show_expanders func(
		tree_view *T.GtkTreeView,
		enabled T.Gboolean)

	Gtk_tree_view_get_show_expanders func(
		tree_view *T.GtkTreeView) T.Gboolean

	Gtk_tree_view_set_level_indentation func(
		tree_view *T.GtkTreeView,
		indentation T.Gint)

	Gtk_tree_view_get_level_indentation func(
		tree_view *T.GtkTreeView) T.Gint

	Gtk_tree_view_set_tooltip_row func(
		tree_view *T.GtkTreeView,
		tooltip *T.GtkTooltip,
		path *T.GtkTreePath)

	Gtk_tree_view_set_tooltip_cell func(
		tree_view *T.GtkTreeView,
		tooltip *T.GtkTooltip,
		path *T.GtkTreePath,
		column *T.GtkTreeViewColumn,
		cell *T.GtkCellRenderer)

	Gtk_tree_view_get_tooltip_context func(
		tree_view *T.GtkTreeView,
		x *T.Gint,
		y *T.Gint,
		keyboard_tip T.Gboolean,
		model **T.GtkTreeModel,
		path **T.GtkTreePath,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_view_set_tooltip_column func(
		tree_view *T.GtkTreeView,
		column T.Gint)

	Gtk_tree_view_get_tooltip_column func(
		tree_view *T.GtkTreeView) T.Gint

	Gtk_combo_box_get_type func() T.GType

	Gtk_combo_box_new func() *T.GtkWidget

	Gtk_combo_box_new_with_entry func() *T.GtkWidget

	Gtk_combo_box_new_with_model func(
		model *T.GtkTreeModel) *T.GtkWidget

	Gtk_combo_box_new_with_model_and_entry func(
		model *T.GtkTreeModel) *T.GtkWidget

	Gtk_combo_box_get_wrap_width func(
		combo_box *T.GtkComboBox) T.Gint

	Gtk_combo_box_set_wrap_width func(
		combo_box *T.GtkComboBox,
		width T.Gint)

	Gtk_combo_box_get_row_span_column func(
		combo_box *T.GtkComboBox) T.Gint

	Gtk_combo_box_set_row_span_column func(
		combo_box *T.GtkComboBox,
		row_span T.Gint)

	Gtk_combo_box_get_column_span_column func(
		combo_box *T.GtkComboBox) T.Gint

	Gtk_combo_box_set_column_span_column func(
		combo_box *T.GtkComboBox,
		column_span T.Gint)

	Gtk_combo_box_get_add_tearoffs func(
		combo_box *T.GtkComboBox) T.Gboolean

	Gtk_combo_box_set_add_tearoffs func(
		combo_box *T.GtkComboBox,
		add_tearoffs T.Gboolean)

	Gtk_combo_box_get_title func(
		combo_box *T.GtkComboBox) string

	Gtk_combo_box_set_title func(
		combo_box *T.GtkComboBox,
		title string)

	Gtk_combo_box_get_focus_on_click func(
		combo *T.GtkComboBox) T.Gboolean

	Gtk_combo_box_set_focus_on_click func(
		combo *T.GtkComboBox,
		focus_on_click T.Gboolean)

	Gtk_combo_box_get_active func(
		combo_box *T.GtkComboBox) T.Gint

	Gtk_combo_box_set_active func(
		combo_box *T.GtkComboBox,
		index_ T.Gint)

	Gtk_combo_box_get_active_iter func(
		combo_box *T.GtkComboBox,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_combo_box_set_active_iter func(
		combo_box *T.GtkComboBox,
		iter *T.GtkTreeIter)

	Gtk_combo_box_set_model func(
		combo_box *T.GtkComboBox,
		model *T.GtkTreeModel)

	Gtk_combo_box_get_model func(
		combo_box *T.GtkComboBox) *T.GtkTreeModel

	Gtk_combo_box_get_row_separator_func func(
		combo_box *T.GtkComboBox) T.GtkTreeViewRowSeparatorFunc

	Gtk_combo_box_set_row_separator_func func(
		combo_box *T.GtkComboBox,
		f T.GtkTreeViewRowSeparatorFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_combo_box_set_button_sensitivity func(
		combo_box *T.GtkComboBox,
		sensitivity T.GtkSensitivityType)

	Gtk_combo_box_get_button_sensitivity func(
		combo_box *T.GtkComboBox) T.GtkSensitivityType

	Gtk_combo_box_get_has_entry func(
		combo_box *T.GtkComboBox) T.Gboolean

	Gtk_combo_box_set_entry_text_column func(
		combo_box *T.GtkComboBox,
		text_column T.Gint)

	Gtk_combo_box_get_entry_text_column func(
		combo_box *T.GtkComboBox) T.Gint

	Gtk_combo_box_new_text func() *T.GtkWidget

	Gtk_combo_box_append_text func(
		combo_box *T.GtkComboBox,
		text string)

	Gtk_combo_box_insert_text func(
		combo_box *T.GtkComboBox,
		position T.Gint,
		text string)

	Gtk_combo_box_prepend_text func(
		combo_box *T.GtkComboBox,
		text string)

	Gtk_combo_box_remove_text func(
		combo_box *T.GtkComboBox,
		position T.Gint)

	Gtk_combo_box_get_active_text func(
		combo_box *T.GtkComboBox) string

	Gtk_combo_box_popup func(
		combo_box *T.GtkComboBox)

	Gtk_combo_box_popdown func(
		combo_box *T.GtkComboBox)

	Gtk_combo_box_get_popup_accessible func(
		combo_box *T.GtkComboBox) *T.AtkObject

	Gtk_combo_box_entry_get_type func() T.GType

	Gtk_combo_box_entry_new func() *T.GtkWidget

	Gtk_combo_box_entry_new_with_model func(
		model *T.GtkTreeModel,
		text_column T.Gint) *T.GtkWidget

	Gtk_combo_box_entry_set_text_column func(
		entry_box *T.GtkComboBoxEntry,
		text_column T.Gint)

	Gtk_combo_box_entry_get_text_column func(
		entry_box *T.GtkComboBoxEntry) T.Gint

	Gtk_combo_box_entry_new_text func() *T.GtkWidget

	Gtk_combo_box_text_get_type func() T.GType

	Gtk_combo_box_text_new func() *T.GtkWidget

	Gtk_combo_box_text_new_with_entry func() *T.GtkWidget

	Gtk_combo_box_text_append_text func(
		combo_box *T.GtkComboBoxText,
		text string)

	Gtk_combo_box_text_insert_text func(
		combo_box *T.GtkComboBoxText,
		position T.Gint,
		text string)

	Gtk_combo_box_text_prepend_text func(
		combo_box *T.GtkComboBoxText,
		text string)

	Gtk_combo_box_text_remove func(
		combo_box *T.GtkComboBoxText,
		position T.Gint)

	Gtk_combo_box_text_get_active_text func(
		combo_box *T.GtkComboBoxText) string

	Gtk_drawing_area_get_type func() T.GType

	Gtk_drawing_area_new func() *T.GtkWidget

	Gtk_drawing_area_size func(
		darea *T.GtkDrawingArea,
		width T.Gint,
		height T.Gint)

	Gtk_event_box_get_type func() T.GType

	Gtk_event_box_new func() *T.GtkWidget

	Gtk_event_box_get_visible_window func(
		event_box *T.GtkEventBox) T.Gboolean

	Gtk_event_box_set_visible_window func(
		event_box *T.GtkEventBox,
		visible_window T.Gboolean)

	Gtk_event_box_get_above_child func(
		event_box *T.GtkEventBox) T.Gboolean

	Gtk_event_box_set_above_child func(
		event_box *T.GtkEventBox,
		above_child T.Gboolean)

	Gtk_expander_get_type func() T.GType

	Gtk_expander_new func(
		label string) *T.GtkWidget

	Gtk_expander_new_with_mnemonic func(
		label string) *T.GtkWidget

	Gtk_expander_set_expanded func(
		expander *T.GtkExpander,
		expanded T.Gboolean)

	Gtk_expander_get_expanded func(
		expander *T.GtkExpander) T.Gboolean

	Gtk_expander_set_spacing func(
		expander *T.GtkExpander,
		spacing T.Gint)

	Gtk_expander_get_spacing func(
		expander *T.GtkExpander) T.Gint

	Gtk_expander_set_label func(
		expander *T.GtkExpander,
		label string)

	Gtk_expander_get_label func(
		expander *T.GtkExpander) string

	Gtk_expander_set_use_underline func(
		expander *T.GtkExpander,
		use_underline T.Gboolean)

	Gtk_expander_get_use_underline func(
		expander *T.GtkExpander) T.Gboolean

	Gtk_expander_set_use_markup func(
		expander *T.GtkExpander,
		use_markup T.Gboolean)

	Gtk_expander_get_use_markup func(
		expander *T.GtkExpander) T.Gboolean

	Gtk_expander_set_label_widget func(
		expander *T.GtkExpander,
		label_widget *T.GtkWidget)

	Gtk_expander_get_label_widget func(
		expander *T.GtkExpander) *T.GtkWidget

	Gtk_expander_set_label_fill func(
		expander *T.GtkExpander,
		label_fill T.Gboolean)

	Gtk_expander_get_label_fill func(
		expander *T.GtkExpander) T.Gboolean

	Gtk_fixed_get_type func() T.GType

	Gtk_fixed_new func() *T.GtkWidget

	Gtk_fixed_put func(
		fixed *T.GtkFixed,
		widget *T.GtkWidget,
		x T.Gint,
		y T.Gint)

	Gtk_fixed_move func(
		fixed *T.GtkFixed,
		widget *T.GtkWidget,
		x T.Gint,
		y T.Gint)

	Gtk_fixed_set_has_window func(
		fixed *T.GtkFixed,
		has_window T.Gboolean)

	Gtk_fixed_get_has_window func(
		fixed *T.GtkFixed) T.Gboolean

	Gtk_file_filter_get_type func() T.GType

	Gtk_file_filter_new func() *T.GtkFileFilter

	Gtk_file_filter_set_name func(
		filter *T.GtkFileFilter,
		name string)

	Gtk_file_filter_get_name func(
		filter *T.GtkFileFilter) string

	Gtk_file_filter_add_mime_type func(
		filter *T.GtkFileFilter,
		mime_type string)

	Gtk_file_filter_add_pattern func(
		filter *T.GtkFileFilter,
		pattern string)

	Gtk_file_filter_add_pixbuf_formats func(
		filter *T.GtkFileFilter)

	Gtk_file_filter_add_custom func(
		filter *T.GtkFileFilter,
		needed T.GtkFileFilterFlags,
		f T.GtkFileFilterFunc,
		dataGpointer,
		notify T.GDestroyNotify)

	Gtk_file_filter_get_needed func(
		filter *T.GtkFileFilter) T.GtkFileFilterFlags

	Gtk_file_filter_filter func(
		filter *T.GtkFileFilter,
		filter_info *T.GtkFileFilterInfo) T.Gboolean

	Gtk_file_chooser_get_type func() T.GType

	Gtk_file_chooser_error_quark func() T.GQuark

	Gtk_file_chooser_set_action func(
		chooser *T.GtkFileChooser,
		action T.GtkFileChooserAction)

	Gtk_file_chooser_get_action func(
		chooser *T.GtkFileChooser) T.GtkFileChooserAction

	Gtk_file_chooser_set_local_only func(
		chooser *T.GtkFileChooser,
		local_only T.Gboolean)

	Gtk_file_chooser_get_local_only func(
		chooser *T.GtkFileChooser) T.Gboolean

	Gtk_file_chooser_set_select_multiple func(
		chooser *T.GtkFileChooser,
		select_multiple T.Gboolean)

	Gtk_file_chooser_get_select_multiple func(
		chooser *T.GtkFileChooser) T.Gboolean

	Gtk_file_chooser_set_show_hidden func(
		chooser *T.GtkFileChooser,
		show_hidden T.Gboolean)

	Gtk_file_chooser_get_show_hidden func(
		chooser *T.GtkFileChooser) T.Gboolean

	Gtk_file_chooser_set_do_overwrite_confirmation func(
		chooser *T.GtkFileChooser,
		do_overwrite_confirmation T.Gboolean)

	Gtk_file_chooser_get_do_overwrite_confirmation func(
		chooser *T.GtkFileChooser) T.Gboolean

	Gtk_file_chooser_set_create_folders func(
		chooser *T.GtkFileChooser,
		create_folders T.Gboolean)

	Gtk_file_chooser_get_create_folders func(
		chooser *T.GtkFileChooser) T.Gboolean

	Gtk_file_chooser_set_current_name func(
		chooser *T.GtkFileChooser,
		name string)

	Gtk_file_chooser_get_filename func(
		chooser *T.GtkFileChooser) string

	Gtk_file_chooser_set_filename func(
		chooser *T.GtkFileChooser,
		filename string) T.Gboolean

	Gtk_file_chooser_select_filename func(
		chooser *T.GtkFileChooser,
		filename string) T.Gboolean

	Gtk_file_chooser_unselect_filename func(
		chooser *T.GtkFileChooser,
		filename string)

	Gtk_file_chooser_select_all func(
		chooser *T.GtkFileChooser)

	Gtk_file_chooser_unselect_all func(
		chooser *T.GtkFileChooser)

	Gtk_file_chooser_get_filenames func(
		chooser *T.GtkFileChooser) *T.GSList

	Gtk_file_chooser_set_current_folder func(
		chooser *T.GtkFileChooser,
		filename string) T.Gboolean

	Gtk_file_chooser_get_current_folder func(
		chooser *T.GtkFileChooser) string

	Gtk_file_chooser_get_uri func(
		chooser *T.GtkFileChooser) string

	Gtk_file_chooser_set_uri func(
		chooser *T.GtkFileChooser,
		uri string) T.Gboolean

	Gtk_file_chooser_select_uri func(
		chooser *T.GtkFileChooser,
		uri string) T.Gboolean

	Gtk_file_chooser_unselect_uri func(
		chooser *T.GtkFileChooser,
		uri string)

	Gtk_file_chooser_get_uris func(
		chooser *T.GtkFileChooser) *T.GSList

	Gtk_file_chooser_set_current_folder_uri func(
		chooser *T.GtkFileChooser,
		uri string) T.Gboolean

	Gtk_file_chooser_get_current_folder_uri func(
		chooser *T.GtkFileChooser) string

	Gtk_file_chooser_get_file func(
		chooser *T.GtkFileChooser) *T.GFile

	Gtk_file_chooser_set_file func(
		chooser *T.GtkFileChooser,
		file *T.GFile,
		error **T.GError) T.Gboolean

	Gtk_file_chooser_select_file func(
		chooser *T.GtkFileChooser,
		file *T.GFile,
		error **T.GError) T.Gboolean

	Gtk_file_chooser_unselect_file func(
		chooser *T.GtkFileChooser,
		file *T.GFile)

	Gtk_file_chooser_get_files func(
		chooser *T.GtkFileChooser) *T.GSList

	Gtk_file_chooser_set_current_folder_file func(
		chooser *T.GtkFileChooser,
		file *T.GFile,
		error **T.GError) T.Gboolean

	Gtk_file_chooser_get_current_folder_file func(
		chooser *T.GtkFileChooser) *T.GFile

	Gtk_file_chooser_set_preview_widget func(
		chooser *T.GtkFileChooser,
		preview_widget *T.GtkWidget)

	Gtk_file_chooser_get_preview_widget func(
		chooser *T.GtkFileChooser) *T.GtkWidget

	Gtk_file_chooser_set_preview_widget_active func(
		chooser *T.GtkFileChooser,
		active T.Gboolean)

	Gtk_file_chooser_get_preview_widget_active func(
		chooser *T.GtkFileChooser) T.Gboolean

	Gtk_file_chooser_set_use_preview_label func(
		chooser *T.GtkFileChooser,
		use_label T.Gboolean)

	Gtk_file_chooser_get_use_preview_label func(
		chooser *T.GtkFileChooser) T.Gboolean

	Gtk_file_chooser_get_preview_filename func(
		chooser *T.GtkFileChooser) string

	Gtk_file_chooser_get_preview_uri func(
		chooser *T.GtkFileChooser) string

	Gtk_file_chooser_get_preview_file func(
		chooser *T.GtkFileChooser) *T.GFile

	Gtk_file_chooser_set_extra_widget func(
		chooser *T.GtkFileChooser,
		extra_widget *T.GtkWidget)

	Gtk_file_chooser_get_extra_widget func(
		chooser *T.GtkFileChooser) *T.GtkWidget

	Gtk_file_chooser_add_filter func(
		chooser *T.GtkFileChooser,
		filter *T.GtkFileFilter)

	Gtk_file_chooser_remove_filter func(
		chooser *T.GtkFileChooser,
		filter *T.GtkFileFilter)

	Gtk_file_chooser_list_filters func(
		chooser *T.GtkFileChooser) *T.GSList

	Gtk_file_chooser_set_filter func(
		chooser *T.GtkFileChooser,
		filter *T.GtkFileFilter)

	Gtk_file_chooser_get_filter func(
		chooser *T.GtkFileChooser) *T.GtkFileFilter

	Gtk_file_chooser_add_shortcut_folder func(
		chooser *T.GtkFileChooser,
		folder string,
		error **T.GError) T.Gboolean

	Gtk_file_chooser_remove_shortcut_folder func(
		chooser *T.GtkFileChooser,
		folder string,
		error **T.GError) T.Gboolean

	Gtk_file_chooser_list_shortcut_folders func(
		chooser *T.GtkFileChooser) *T.GSList

	Gtk_file_chooser_add_shortcut_folder_uri func(
		chooser *T.GtkFileChooser,
		uri string,
		error **T.GError) T.Gboolean

	Gtk_file_chooser_remove_shortcut_folder_uri func(
		chooser *T.GtkFileChooser,
		uri string,
		error **T.GError) T.Gboolean

	Gtk_file_chooser_list_shortcut_folder_uris func(
		chooser *T.GtkFileChooser) *T.GSList

	Gtk_hbox_get_type func() T.GType

	Gtk_hbox_new func(
		homogeneous T.Gboolean,
		spacing T.Gint) *T.GtkWidget

	Gtk_file_chooser_button_get_type func() T.GType

	Gtk_file_chooser_button_new func(
		title string,
		action T.GtkFileChooserAction) *T.GtkWidget

	Gtk_file_chooser_button_new_with_backend func(
		title string,
		action T.GtkFileChooserAction,
		backend string) *T.GtkWidget

	Gtk_file_chooser_button_new_with_dialog func(
		dialog *T.GtkWidget) *T.GtkWidget

	Gtk_file_chooser_button_get_title func(
		button *T.GtkFileChooserButton) string

	Gtk_file_chooser_button_set_title func(
		button *T.GtkFileChooserButton,
		title string)

	Gtk_file_chooser_button_get_width_chars func(
		button *T.GtkFileChooserButton) T.Gint

	Gtk_file_chooser_button_set_width_chars func(
		button *T.GtkFileChooserButton,
		n_chars T.Gint)

	Gtk_file_chooser_button_get_focus_on_click func(
		button *T.GtkFileChooserButton) T.Gboolean

	Gtk_file_chooser_button_set_focus_on_click func(
		button *T.GtkFileChooserButton,
		focus_on_click T.Gboolean)

	Gtk_file_chooser_dialog_get_type func() T.GType

	Gtk_file_chooser_dialog_new func(title string, parent *T.GtkWindow,
		action T.GtkFileChooserAction, first_button_text string,
		v ...VArg) *T.GtkWidget

	Gtk_file_chooser_dialog_new_with_backend func(title string,
		parent *T.GtkWindow, action T.GtkFileChooserAction,
		backend string, first_button_text string,
		v ...VArg) *T.GtkWidget

	Gtk_file_chooser_widget_get_type func() T.GType

	Gtk_file_chooser_widget_new func(
		action T.GtkFileChooserAction) *T.GtkWidget

	Gtk_file_chooser_widget_new_with_backend func(
		action T.GtkFileChooserAction,
		backend string) *T.GtkWidget

	Gtk_font_button_get_type func() T.GType

	Gtk_font_button_new func() *T.GtkWidget

	Gtk_font_button_new_with_font func(
		fontname string) *T.GtkWidget

	Gtk_font_button_get_title func(
		font_button *T.GtkFontButton) string

	Gtk_font_button_set_title func(
		font_button *T.GtkFontButton,
		title string)

	Gtk_font_button_get_use_font func(
		font_button *T.GtkFontButton) T.Gboolean

	Gtk_font_button_set_use_font func(
		font_button *T.GtkFontButton,
		use_font T.Gboolean)

	Gtk_font_button_get_use_size func(
		font_button *T.GtkFontButton) T.Gboolean

	Gtk_font_button_set_use_size func(
		font_button *T.GtkFontButton,
		use_size T.Gboolean)

	Gtk_font_button_get_font_name func(
		font_button *T.GtkFontButton) string

	Gtk_font_button_set_font_name func(
		font_button *T.GtkFontButton,
		fontname string) T.Gboolean

	Gtk_font_button_get_show_style func(
		font_button *T.GtkFontButton) T.Gboolean

	Gtk_font_button_set_show_style func(
		font_button *T.GtkFontButton,
		show_style T.Gboolean)

	Gtk_font_button_get_show_size func(
		font_button *T.GtkFontButton) T.Gboolean

	Gtk_font_button_set_show_size func(
		font_button *T.GtkFontButton,
		show_size T.Gboolean)

	Gtk_font_selection_get_type func() T.GType

	Gtk_font_selection_new func() *T.GtkWidget

	Gtk_font_selection_get_family_list func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	Gtk_font_selection_get_face_list func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	Gtk_font_selection_get_size_entry func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	Gtk_font_selection_get_size_list func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	Gtk_font_selection_get_preview_entry func(
		fontsel *T.GtkFontSelection) *T.GtkWidget

	Gtk_font_selection_get_family func(
		fontsel *T.GtkFontSelection) *T.PangoFontFamily

	Gtk_font_selection_get_face func(
		fontsel *T.GtkFontSelection) *T.PangoFontFace

	Gtk_font_selection_get_size func(
		fontsel *T.GtkFontSelection) T.Gint

	Gtk_font_selection_get_font_name func(
		fontsel *T.GtkFontSelection) string

	Gtk_font_selection_get_font func(
		fontsel *T.GtkFontSelection) *T.GdkFont

	Gtk_font_selection_set_font_name func(
		fontsel *T.GtkFontSelection,
		fontname string) T.Gboolean

	Gtk_font_selection_get_preview_text func(
		fontsel *T.GtkFontSelection) string

	Gtk_font_selection_set_preview_text func(
		fontsel *T.GtkFontSelection,
		text string)

	Gtk_font_selection_dialog_get_type func() T.GType

	Gtk_font_selection_dialog_new func(
		title string) *T.GtkWidget

	Gtk_font_selection_dialog_get_ok_button func(
		fsd *T.GtkFontSelectionDialog) *T.GtkWidget

	Gtk_font_selection_dialog_get_apply_button func(
		fsd *T.GtkFontSelectionDialog) *T.GtkWidget

	Gtk_font_selection_dialog_get_cancel_button func(
		fsd *T.GtkFontSelectionDialog) *T.GtkWidget

	Gtk_font_selection_dialog_get_font_selection func(
		fsd *T.GtkFontSelectionDialog) *T.GtkWidget

	Gtk_font_selection_dialog_get_font_name func(
		fsd *T.GtkFontSelectionDialog) string

	Gtk_font_selection_dialog_get_font func(
		fsd *T.GtkFontSelectionDialog) *T.GdkFont

	Gtk_font_selection_dialog_set_font_name func(
		fsd *T.GtkFontSelectionDialog,
		fontname string) T.Gboolean

	Gtk_font_selection_dialog_get_preview_text func(
		fsd *T.GtkFontSelectionDialog) string

	Gtk_font_selection_dialog_set_preview_text func(
		fsd *T.GtkFontSelectionDialog,
		text string)

	Gtk_gc_get func(
		depth T.Gint,
		colormap *T.GdkColormap,
		values *T.GdkGCValues,
		values_mask T.GdkGCValuesMask) *T.GdkGC

	Gtk_gc_release func(
		gc *T.GdkGC)

	Gtk_handle_box_get_type func() T.GType

	Gtk_handle_box_new func() *T.GtkWidget

	Gtk_handle_box_set_shadow_type func(
		handle_box *T.GtkHandleBox,
		t T.GtkShadowType)

	Gtk_handle_box_get_shadow_type func(
		handle_box *T.GtkHandleBox) T.GtkShadowType

	Gtk_handle_box_set_handle_position func(
		handle_box *T.GtkHandleBox,
		position T.GtkPositionType)

	Gtk_handle_box_get_handle_position func(
		handle_box *T.GtkHandleBox) T.GtkPositionType

	Gtk_handle_box_set_snap_edge func(
		handle_box *T.GtkHandleBox,
		edge T.GtkPositionType)

	Gtk_handle_box_get_snap_edge func(
		handle_box *T.GtkHandleBox) T.GtkPositionType

	Gtk_handle_box_get_child_detached func(
		handle_box *T.GtkHandleBox) T.Gboolean

	Gtk_hbutton_box_get_type func() T.GType

	Gtk_hbutton_box_new func() *T.GtkWidget

	Gtk_hbutton_box_get_spacing_default func() T.Gint

	Gtk_hbutton_box_get_layout_default func() T.GtkButtonBoxStyle

	Gtk_hbutton_box_set_spacing_default func(
		spacing T.Gint)

	Gtk_hbutton_box_set_layout_default func(
		layout T.GtkButtonBoxStyle)

	Gtk_paned_get_type func() T.GType

	Gtk_paned_add1 func(
		paned *T.GtkPaned,
		child *T.GtkWidget)

	Gtk_paned_add2 func(
		paned *T.GtkPaned,
		child *T.GtkWidget)

	Gtk_paned_pack1 func(
		paned *T.GtkPaned,
		child *T.GtkWidget,
		resize T.Gboolean,
		shrink T.Gboolean)

	Gtk_paned_pack2 func(
		paned *T.GtkPaned,
		child *T.GtkWidget,
		resize T.Gboolean,
		shrink T.Gboolean)

	Gtk_paned_get_position func(
		paned *T.GtkPaned) T.Gint

	Gtk_paned_set_position func(
		paned *T.GtkPaned,
		position T.Gint)

	Gtk_paned_get_child1 func(
		paned *T.GtkPaned) *T.GtkWidget

	Gtk_paned_get_child2 func(
		paned *T.GtkPaned) *T.GtkWidget

	Gtk_paned_get_handle_window func(
		paned *T.GtkPaned) *T.GdkWindow

	Gtk_paned_compute_position func(
		paned *T.GtkPaned,
		allocation T.Gint,
		child1_req T.Gint,
		child2_req T.Gint)

	Gtk_hpaned_get_type func() T.GType

	Gtk_hpaned_new func() *T.GtkWidget

	Gtk_ruler_get_type func() T.GType

	Gtk_ruler_set_metric func(
		ruler *T.GtkRuler,
		metric T.GtkMetricType)

	Gtk_ruler_get_metric func(
		ruler *T.GtkRuler) T.GtkMetricType

	Gtk_ruler_set_range func(
		ruler *T.GtkRuler,
		lower T.Gdouble,
		upper T.Gdouble,
		position T.Gdouble,
		max_size T.Gdouble)

	Gtk_ruler_get_range func(
		ruler *T.GtkRuler,
		lower *T.Gdouble,
		upper *T.Gdouble,
		position *T.Gdouble,
		max_size *T.Gdouble)

	Gtk_ruler_draw_ticks func(
		ruler *T.GtkRuler)

	Gtk_ruler_draw_pos func(
		ruler *T.GtkRuler)

	Gtk_hruler_get_type func() T.GType

	Gtk_hruler_new func() *T.GtkWidget

	Gtk_range_get_type func() T.GType

	Gtk_range_set_update_policy func(
		r *T.GtkRange,
		policy T.GtkUpdateType)

	Gtk_range_get_update_policy func(
		r *T.GtkRange) T.GtkUpdateType

	Gtk_range_set_adjustment func(
		r *T.GtkRange,
		adjustment *T.GtkAdjustment)

	Gtk_range_get_adjustment func(
		r *T.GtkRange) *T.GtkAdjustment

	Gtk_range_set_inverted func(
		r *T.GtkRange,
		setting T.Gboolean)

	Gtk_range_get_inverted func(
		r *T.GtkRange) T.Gboolean

	Gtk_range_set_flippable func(
		r *T.GtkRange,
		flippable T.Gboolean)

	Gtk_range_get_flippable func(
		r *T.GtkRange) T.Gboolean

	Gtk_range_set_slider_size_fixed func(
		r *T.GtkRange,
		size_fixed T.Gboolean)

	Gtk_range_get_slider_size_fixed func(
		r *T.GtkRange) T.Gboolean

	Gtk_range_set_min_slider_size func(
		r *T.GtkRange,
		min_size T.Gboolean)

	Gtk_range_get_min_slider_size func(
		r *T.GtkRange) T.Gint

	Gtk_range_get_range_rect func(
		r *T.GtkRange,
		range_rect *T.GdkRectangle)

	Gtk_range_get_slider_range func(
		r *T.GtkRange,
		slider_start *T.Gint,
		slider_end *T.Gint)

	Gtk_range_set_lower_stepper_sensitivity func(
		r *T.GtkRange,
		sensitivity T.GtkSensitivityType)

	Gtk_range_get_lower_stepper_sensitivity func(
		r *T.GtkRange) T.GtkSensitivityType

	Gtk_range_set_upper_stepper_sensitivity func(
		r *T.GtkRange,
		sensitivity T.GtkSensitivityType)

	Gtk_range_get_upper_stepper_sensitivity func(
		r *T.GtkRange) T.GtkSensitivityType

	Gtk_range_set_increments func(
		r *T.GtkRange,
		step T.Gdouble,
		page T.Gdouble)

	Gtk_range_set_range func(
		r *T.GtkRange,
		min T.Gdouble,
		max T.Gdouble)

	Gtk_range_set_value func(
		r *T.GtkRange,
		value T.Gdouble)

	Gtk_range_get_value func(
		r *T.GtkRange) T.Gdouble

	Gtk_range_set_show_fill_level func(
		r *T.GtkRange,
		show_fill_level T.Gboolean)

	Gtk_range_get_show_fill_level func(
		r *T.GtkRange) T.Gboolean

	Gtk_range_set_restrict_to_fill_level func(
		r *T.GtkRange,
		restrict_to_fill_level T.Gboolean)

	Gtk_range_get_restrict_to_fill_level func(
		r *T.GtkRange) T.Gboolean

	Gtk_range_set_fill_level func(
		r *T.GtkRange,
		fill_level T.Gdouble)

	Gtk_range_get_fill_level func(
		r *T.GtkRange) T.Gdouble

	Gtk_range_set_round_digits func(
		r *T.GtkRange,
		round_digits T.Gint)

	Gtk_range_get_round_digits func(
		r *T.GtkRange) T.Gint

	Gtk_scale_get_type func() T.GType

	Gtk_scale_set_digits func(
		scale *T.GtkScale,
		digits T.Gint)

	Gtk_scale_get_digits func(
		scale *T.GtkScale) T.Gint

	Gtk_scale_set_draw_value func(
		scale *T.GtkScale,
		draw_value T.Gboolean)

	Gtk_scale_get_draw_value func(
		scale *T.GtkScale) T.Gboolean

	Gtk_scale_set_value_pos func(
		scale *T.GtkScale,
		pos T.GtkPositionType)

	Gtk_scale_get_value_pos func(
		scale *T.GtkScale) T.GtkPositionType

	Gtk_scale_get_layout func(
		scale *T.GtkScale) *T.PangoLayout

	Gtk_scale_get_layout_offsets func(
		scale *T.GtkScale,
		x *T.Gint,
		y *T.Gint)

	Gtk_scale_add_mark func(
		scale *T.GtkScale,
		value T.Gdouble,
		position T.GtkPositionType,
		markup string)

	Gtk_scale_clear_marks func(
		scale *T.GtkScale)

	Gtk_hscale_get_type func() T.GType

	Gtk_hscale_new func(
		adjustment *T.GtkAdjustment) *T.GtkWidget

	Gtk_hscale_new_with_range func(
		min T.Gdouble,
		max T.Gdouble,
		step T.Gdouble) *T.GtkWidget

	Gtk_scrollbar_get_type func() T.GType

	Gtk_hscrollbar_get_type func() T.GType

	Gtk_hscrollbar_new func(
		adjustment *T.GtkAdjustment) *T.GtkWidget

	Gtk_separator_get_type func() T.GType

	Gtk_hseparator_get_type func() T.GType

	Gtk_hseparator_new func() *T.GtkWidget

	Gtk_hsv_get_type func() T.GType

	Gtk_hsv_new func() *T.GtkWidget

	Gtk_hsv_set_color func(
		hsv *T.GtkHSV, h, s, v T.Double)

	Gtk_hsv_get_color func(
		hsv *T.GtkHSV, h, s, v *T.Gdouble)

	Gtk_hsv_set_metrics func(
		hsv *T.GtkHSV,
		size T.Gint,
		ring_width T.Gint)

	Gtk_hsv_get_metrics func(
		hsv *T.GtkHSV,
		size *T.Gint,
		ring_width *T.Gint)

	Gtk_hsv_is_adjusting func(
		hsv *T.GtkHSV) T.Gboolean

	Gtk_hsv_to_rgb func(h, s, v T.Gdouble, r, g, b *T.Gdouble)

	Gtk_rgb_to_hsv func(r, g, b T.Gdouble, h, s, v *T.Gdouble)

	Gtk_icon_factory_get_type func() T.GType

	Gtk_icon_factory_new func() *T.GtkIconFactory

	Gtk_icon_factory_add func(
		factory *T.GtkIconFactory,
		stock_id string,
		icon_set *T.GtkIconSet)

	Gtk_icon_factory_lookup func(
		factory *T.GtkIconFactory,
		stock_id string) *T.GtkIconSet

	Gtk_icon_factory_add_default func(
		factory *T.GtkIconFactory)

	Gtk_icon_factory_remove_default func(
		factory *T.GtkIconFactory)

	Gtk_icon_factory_lookup_default func(
		stock_id string) *T.GtkIconSet

	Gtk_icon_size_lookup func(
		size T.GtkIconSize,
		width *T.Gint,
		height *T.Gint) T.Gboolean

	Gtk_icon_size_lookup_for_settings func(
		settings *T.GtkSettings,
		size T.GtkIconSize,
		width *T.Gint,
		height *T.Gint) T.Gboolean

	Gtk_icon_size_register func(
		name string,
		width T.Gint,
		height T.Gint) T.GtkIconSize

	Gtk_icon_size_register_alias func(
		alias string,
		target T.GtkIconSize)

	Gtk_icon_size_from_name func(
		name string) T.GtkIconSize

	Gtk_icon_size_get_name func(
		size T.GtkIconSize) string

	Gtk_icon_set_get_type func() T.GType

	Gtk_icon_set_new func() *T.GtkIconSet

	Gtk_icon_set_new_from_pixbuf func(
		pixbuf *T.GdkPixbuf) *T.GtkIconSet

	Gtk_icon_set_ref func(
		icon_set *T.GtkIconSet) *T.GtkIconSet

	Gtk_icon_set_unref func(
		icon_set *T.GtkIconSet)

	Gtk_icon_set_copy func(
		icon_set *T.GtkIconSet) *T.GtkIconSet

	Gtk_icon_set_render_icon func(
		icon_set *T.GtkIconSet,
		style *T.GtkStyle,
		direction T.GtkTextDirection,
		state T.GtkStateType,
		size T.GtkIconSize,
		widget *T.GtkWidget,
		detail string) *T.GdkPixbuf

	Gtk_icon_set_add_source func(
		icon_set *T.GtkIconSet,
		source *T.GtkIconSource)

	Gtk_icon_set_get_sizes func(
		icon_set *T.GtkIconSet,
		sizes **T.GtkIconSize,
		n_sizes *T.Gint)

	Gtk_icon_source_get_type func() T.GType

	Gtk_icon_source_new func() *T.GtkIconSource

	Gtk_icon_source_copy func(
		source *T.GtkIconSource) *T.GtkIconSource

	Gtk_icon_source_free func(
		source *T.GtkIconSource)

	Gtk_icon_source_set_filename func(
		source *T.GtkIconSource,
		filename string)

	Gtk_icon_source_set_icon_name func(
		source *T.GtkIconSource,
		icon_name string)

	Gtk_icon_source_set_pixbuf func(
		source *T.GtkIconSource,
		pixbuf *T.GdkPixbuf)

	Gtk_icon_source_get_filename func(
		source *T.GtkIconSource) string

	Gtk_icon_source_get_icon_name func(
		source *T.GtkIconSource) string

	Gtk_icon_source_get_pixbuf func(
		source *T.GtkIconSource) *T.GdkPixbuf

	Gtk_icon_source_set_direction_wildcarded func(
		source *T.GtkIconSource,
		setting T.Gboolean)

	Gtk_icon_source_set_state_wildcarded func(
		source *T.GtkIconSource,
		setting T.Gboolean)

	Gtk_icon_source_set_size_wildcarded func(
		source *T.GtkIconSource,
		setting T.Gboolean)

	Gtk_icon_source_get_size_wildcarded func(
		source *T.GtkIconSource) T.Gboolean

	Gtk_icon_source_get_state_wildcarded func(
		source *T.GtkIconSource) T.Gboolean

	Gtk_icon_source_get_direction_wildcarded func(
		source *T.GtkIconSource) T.Gboolean

	Gtk_icon_source_set_direction func(
		source *T.GtkIconSource,
		direction T.GtkTextDirection)

	Gtk_icon_source_set_state func(
		source *T.GtkIconSource,
		state T.GtkStateType)

	Gtk_icon_source_set_size func(
		source *T.GtkIconSource,
		size T.GtkIconSize)

	Gtk_icon_source_get_direction func(
		source *T.GtkIconSource) T.GtkTextDirection

	Gtk_icon_source_get_state func(
		source *T.GtkIconSource) T.GtkStateType

	Gtk_icon_source_get_size func(
		source *T.GtkIconSource) T.GtkIconSize

	Gtk_icon_theme_error_quark func() T.GQuark

	Gtk_icon_theme_get_type func() T.GType

	Gtk_icon_theme_new func() *T.GtkIconTheme

	Gtk_icon_theme_get_default func() *T.GtkIconTheme

	Gtk_icon_theme_get_for_screen func(
		screen *T.GdkScreen) *T.GtkIconTheme

	Gtk_icon_theme_set_screen func(
		icon_theme *T.GtkIconTheme,
		screen *T.GdkScreen)

	Gtk_icon_theme_set_search_path func(
		icon_theme *T.GtkIconTheme,
		path **T.Gchar,
		n_elements T.Gint)

	Gtk_icon_theme_get_search_path func(
		icon_theme *T.GtkIconTheme,
		path ***T.Gchar,
		n_elements *T.Gint)

	Gtk_icon_theme_append_search_path func(
		icon_theme *T.GtkIconTheme,
		path string)

	Gtk_icon_theme_prepend_search_path func(
		icon_theme *T.GtkIconTheme,
		path string)

	Gtk_icon_theme_set_custom_theme func(
		icon_theme *T.GtkIconTheme,
		theme_name string)

	Gtk_icon_theme_has_icon func(
		icon_theme *T.GtkIconTheme,
		icon_name string) T.Gboolean

	Gtk_icon_theme_get_icon_sizes func(
		icon_theme *T.GtkIconTheme,
		icon_name string) *T.Gint

	Gtk_icon_theme_lookup_icon func(
		icon_theme *T.GtkIconTheme,
		icon_name string,
		size T.Gint,
		flags T.GtkIconLookupFlags) *T.GtkIconInfo

	Gtk_icon_theme_choose_icon func(
		icon_theme *T.GtkIconTheme,
		icon_names **T.Gchar,
		size T.Gint,
		flags T.GtkIconLookupFlags) *T.GtkIconInfo

	Gtk_icon_theme_load_icon func(
		icon_theme *T.GtkIconTheme,
		icon_name string,
		size T.Gint,
		flags T.GtkIconLookupFlags,
		error **T.GError) *T.GdkPixbuf

	Gtk_icon_theme_lookup_by_gicon func(
		icon_theme *T.GtkIconTheme,
		icon *T.GIcon,
		size T.Gint,
		flags T.GtkIconLookupFlags) *T.GtkIconInfo

	Gtk_icon_theme_list_icons func(
		icon_theme *T.GtkIconTheme,
		context string) *T.GList

	Gtk_icon_theme_list_contexts func(
		icon_theme *T.GtkIconTheme) *T.GList

	Gtk_icon_theme_get_example_icon_name func(
		icon_theme *T.GtkIconTheme) string

	Gtk_icon_theme_rescan_if_needed func(
		icon_theme *T.GtkIconTheme) T.Gboolean

	Gtk_icon_theme_add_builtin_icon func(
		icon_name string,
		size T.Gint,
		pixbuf *T.GdkPixbuf)

	Gtk_icon_info_get_type func() T.GType

	Gtk_icon_info_copy func(
		icon_info *T.GtkIconInfo) *T.GtkIconInfo

	Gtk_icon_info_free func(
		icon_info *T.GtkIconInfo)

	Gtk_icon_info_new_for_pixbuf func(
		icon_theme *T.GtkIconTheme,
		pixbuf *T.GdkPixbuf) *T.GtkIconInfo

	Gtk_icon_info_get_base_size func(
		icon_info *T.GtkIconInfo) T.Gint

	Gtk_icon_info_get_filename func(
		icon_info *T.GtkIconInfo) string

	Gtk_icon_info_get_builtin_pixbuf func(
		icon_info *T.GtkIconInfo) *T.GdkPixbuf

	Gtk_icon_info_load_icon func(
		icon_info *T.GtkIconInfo,
		error **T.GError) *T.GdkPixbuf

	Gtk_icon_info_set_raw_coordinates func(
		icon_info *T.GtkIconInfo,
		raw_coordinates T.Gboolean)

	Gtk_icon_info_get_embedded_rect func(
		icon_info *T.GtkIconInfo,
		rectangle *T.GdkRectangle) T.Gboolean

	Gtk_icon_info_get_attach_points func(
		icon_info *T.GtkIconInfo,
		points **T.GdkPoint,
		n_points *T.Gint) T.Gboolean

	Gtk_icon_info_get_display_name func(
		icon_info *T.GtkIconInfo) string

	Gtk_tooltip_get_type func() T.GType

	Gtk_tooltip_set_markup func(
		tooltip *T.GtkTooltip,
		markup string)

	Gtk_tooltip_set_text func(
		tooltip *T.GtkTooltip,
		text string)

	Gtk_tooltip_set_icon func(
		tooltip *T.GtkTooltip,
		pixbuf *T.GdkPixbuf)

	Gtk_tooltip_set_icon_from_stock func(
		tooltip *T.GtkTooltip,
		stock_id string,
		size T.GtkIconSize)

	Gtk_tooltip_set_icon_from_icon_name func(
		tooltip *T.GtkTooltip,
		icon_name string,
		size T.GtkIconSize)

	Gtk_tooltip_set_icon_from_gicon func(
		tooltip *T.GtkTooltip,
		gicon *T.GIcon,
		size T.GtkIconSize)

	Gtk_tooltip_set_custom func(
		tooltip *T.GtkTooltip,
		custom_widget *T.GtkWidget)

	Gtk_tooltip_set_tip_area func(
		tooltip *T.GtkTooltip,
		rect *T.GdkRectangle)

	Gtk_tooltip_trigger_tooltip_query func(
		display *T.GdkDisplay)

	Gtk_icon_view_get_type func() T.GType

	Gtk_icon_view_new func() *T.GtkWidget

	Gtk_icon_view_new_with_model func(
		model *T.GtkTreeModel) *T.GtkWidget

	Gtk_icon_view_set_model func(
		icon_view *T.GtkIconView,
		model *T.GtkTreeModel)

	Gtk_icon_view_get_model func(
		icon_view *T.GtkIconView) *T.GtkTreeModel

	Gtk_icon_view_set_text_column func(
		icon_view *T.GtkIconView,
		column T.Gint)

	Gtk_icon_view_get_text_column func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_markup_column func(
		icon_view *T.GtkIconView,
		column T.Gint)

	Gtk_icon_view_get_markup_column func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_pixbuf_column func(
		icon_view *T.GtkIconView,
		column T.Gint)

	Gtk_icon_view_get_pixbuf_column func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_orientation func(
		icon_view *T.GtkIconView,
		orientation T.GtkOrientation)

	Gtk_icon_view_get_orientation func(
		icon_view *T.GtkIconView) T.GtkOrientation

	Gtk_icon_view_set_item_orientation func(
		icon_view *T.GtkIconView,
		orientation T.GtkOrientation)

	Gtk_icon_view_get_item_orientation func(
		icon_view *T.GtkIconView) T.GtkOrientation

	Gtk_icon_view_set_columns func(
		icon_view *T.GtkIconView,
		columns T.Gint)

	Gtk_icon_view_get_columns func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_item_width func(
		icon_view *T.GtkIconView,
		item_width T.Gint)

	Gtk_icon_view_get_item_width func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_spacing func(
		icon_view *T.GtkIconView,
		spacing T.Gint)

	Gtk_icon_view_get_spacing func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_row_spacing func(
		icon_view *T.GtkIconView,
		row_spacing T.Gint)

	Gtk_icon_view_get_row_spacing func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_column_spacing func(
		icon_view *T.GtkIconView,
		column_spacing T.Gint)

	Gtk_icon_view_get_column_spacing func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_margin func(
		icon_view *T.GtkIconView,
		margin T.Gint)

	Gtk_icon_view_get_margin func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_set_item_padding func(
		icon_view *T.GtkIconView,
		item_padding T.Gint)

	Gtk_icon_view_get_item_padding func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_icon_view_get_path_at_pos func(
		icon_view *T.GtkIconView,
		x T.Gint,
		y T.Gint) *T.GtkTreePath

	Gtk_icon_view_get_item_at_pos func(
		icon_view *T.GtkIconView,
		x T.Gint,
		y T.Gint,
		path **T.GtkTreePath,
		cell **T.GtkCellRenderer) T.Gboolean

	Gtk_icon_view_get_visible_range func(
		icon_view *T.GtkIconView,
		start_path **T.GtkTreePath,
		end_path **T.GtkTreePath) T.Gboolean

	Gtk_icon_view_selected_foreach func(
		icon_view *T.GtkIconView,
		f T.GtkIconViewForeachFunc,
		data T.Gpointer)

	Gtk_icon_view_set_selection_mode func(
		icon_view *T.GtkIconView,
		mode T.GtkSelectionMode)

	Gtk_icon_view_get_selection_mode func(
		icon_view *T.GtkIconView) T.GtkSelectionMode

	Gtk_icon_view_select_path func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath)

	Gtk_icon_view_unselect_path func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath)

	Gtk_icon_view_path_is_selected func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath) T.Gboolean

	Gtk_icon_view_get_item_row func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath) T.Gint

	Gtk_icon_view_get_item_column func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath) T.Gint

	Gtk_icon_view_get_selected_items func(
		icon_view *T.GtkIconView) *T.GList

	Gtk_icon_view_select_all func(
		icon_view *T.GtkIconView)

	Gtk_icon_view_unselect_all func(
		icon_view *T.GtkIconView)

	Gtk_icon_view_item_activated func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath)

	Gtk_icon_view_set_cursor func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath,
		cell *T.GtkCellRenderer,
		start_editing T.Gboolean)

	Gtk_icon_view_get_cursor func(
		icon_view *T.GtkIconView,
		path **T.GtkTreePath,
		cell **T.GtkCellRenderer) T.Gboolean

	Gtk_icon_view_scroll_to_path func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath,
		use_align T.Gboolean,
		row_align T.Gfloat,
		col_align T.Gfloat)

	Gtk_icon_view_enable_model_drag_source func(
		icon_view *T.GtkIconView,
		start_button_mask T.GdkModifierType,
		targets *T.GtkTargetEntry,
		n_targets T.Gint,
		actions T.GdkDragAction)

	Gtk_icon_view_enable_model_drag_dest func(
		icon_view *T.GtkIconView,
		targets *T.GtkTargetEntry,
		n_targets T.Gint,
		actions T.GdkDragAction)

	Gtk_icon_view_unset_model_drag_source func(
		icon_view *T.GtkIconView)

	Gtk_icon_view_unset_model_drag_dest func(
		icon_view *T.GtkIconView)

	Gtk_icon_view_set_reorderable func(
		icon_view *T.GtkIconView,
		reorderable T.Gboolean)

	Gtk_icon_view_get_reorderable func(
		icon_view *T.GtkIconView) T.Gboolean

	Gtk_icon_view_set_drag_dest_item func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath,
		pos T.GtkIconViewDropPosition)

	Gtk_icon_view_get_drag_dest_item func(
		icon_view *T.GtkIconView,
		path **T.GtkTreePath,
		pos *T.GtkIconViewDropPosition)

	Gtk_icon_view_get_dest_item_at_pos func(
		icon_view *T.GtkIconView,
		drag_x T.Gint,
		drag_y T.Gint,
		path **T.GtkTreePath,
		pos *T.GtkIconViewDropPosition) T.Gboolean

	Gtk_icon_view_create_drag_icon func(
		icon_view *T.GtkIconView,
		path *T.GtkTreePath) *T.GdkPixmap

	Gtk_icon_view_convert_widget_to_bin_window_coords func(
		icon_view *T.GtkIconView,
		wx T.Gint,
		wy T.Gint,
		bx *T.Gint,
		by *T.Gint)

	Gtk_icon_view_set_tooltip_item func(
		icon_view *T.GtkIconView,
		tooltip *T.GtkTooltip,
		path *T.GtkTreePath)

	Gtk_icon_view_set_tooltip_cell func(
		icon_view *T.GtkIconView,
		tooltip *T.GtkTooltip,
		path *T.GtkTreePath,
		cell *T.GtkCellRenderer)

	Gtk_icon_view_get_tooltip_context func(
		icon_view *T.GtkIconView,
		x *T.Gint,
		y *T.Gint,
		keyboard_tip T.Gboolean,
		model **T.GtkTreeModel,
		path **T.GtkTreePath,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_icon_view_set_tooltip_column func(
		icon_view *T.GtkIconView,
		column T.Gint)

	Gtk_icon_view_get_tooltip_column func(
		icon_view *T.GtkIconView) T.Gint

	Gtk_image_menu_item_get_type func() T.GType

	Gtk_image_menu_item_new func() *T.GtkWidget

	Gtk_image_menu_item_new_with_label func(
		label string) *T.GtkWidget

	Gtk_image_menu_item_new_with_mnemonic func(
		label string) *T.GtkWidget

	Gtk_image_menu_item_new_from_stock func(
		stock_id string,
		accel_group *T.GtkAccelGroup) *T.GtkWidget

	Gtk_image_menu_item_set_always_show_image func(
		image_menu_item *T.GtkImageMenuItem,
		always_show T.Gboolean)

	Gtk_image_menu_item_get_always_show_image func(
		image_menu_item *T.GtkImageMenuItem) T.Gboolean

	Gtk_image_menu_item_set_image func(
		image_menu_item *T.GtkImageMenuItem,
		image *T.GtkWidget)

	Gtk_image_menu_item_get_image func(
		image_menu_item *T.GtkImageMenuItem) *T.GtkWidget

	Gtk_image_menu_item_set_use_stock func(
		image_menu_item *T.GtkImageMenuItem,
		use_stock T.Gboolean)

	Gtk_image_menu_item_get_use_stock func(
		image_menu_item *T.GtkImageMenuItem) T.Gboolean

	Gtk_image_menu_item_set_accel_group func(
		image_menu_item *T.GtkImageMenuItem,
		accel_group *T.GtkAccelGroup)

	Gtk_im_context_simple_get_type func() T.GType

	Gtk_im_context_simple_new func() *T.GtkIMContext

	Gtk_im_context_simple_add_table func(
		context_simple *T.GtkIMContextSimple,
		data *T.Guint16,
		max_seq_len T.Gint,
		n_seqs T.Gint)

	Gtk_im_multicontext_get_type func() T.GType

	Gtk_im_multicontext_new func() *T.GtkIMContext

	Gtk_im_multicontext_append_menuitems func(
		context *T.GtkIMMulticontext,
		menushell *T.GtkMenuShell)

	Gtk_im_multicontext_get_context_id func(
		context *T.GtkIMMulticontext) string

	Gtk_im_multicontext_set_context_id func(
		context *T.GtkIMMulticontext,
		context_id string)

	Gtk_info_bar_get_type func() T.GType

	Gtk_info_bar_new func() *T.GtkWidget

	Gtk_info_bar_new_with_buttons func(
		first_button_text string, v ...VArg) *T.GtkWidget

	Gtk_info_bar_get_action_area func(
		info_bar *T.GtkInfoBar) *T.GtkWidget

	Gtk_info_bar_get_content_area func(
		info_bar *T.GtkInfoBar) *T.GtkWidget

	Gtk_info_bar_add_action_widget func(
		info_bar *T.GtkInfoBar,
		child *T.GtkWidget,
		response_id T.Gint)

	Gtk_info_bar_add_button func(
		info_bar *T.GtkInfoBar,
		button_text string,
		response_id T.Gint) *T.GtkWidget

	Gtk_info_bar_add_buttons func(info_bar *T.GtkInfoBar,
		first_button_text string, v ...VArg)

	Gtk_info_bar_set_response_sensitive func(
		info_bar *T.GtkInfoBar,
		response_id T.Gint,
		setting T.Gboolean)

	Gtk_info_bar_set_default_response func(
		info_bar *T.GtkInfoBar,
		response_id T.Gint)

	Gtk_info_bar_response func(
		info_bar *T.GtkInfoBar,
		response_id T.Gint)

	Gtk_info_bar_set_message_type func(
		info_bar *T.GtkInfoBar,
		message_type T.GtkMessageType)

	Gtk_info_bar_get_message_type func(
		info_bar *T.GtkInfoBar) T.GtkMessageType

	Gtk_invisible_get_type func() T.GType

	Gtk_invisible_new func() *T.GtkWidget

	Gtk_invisible_new_for_screen func(
		screen *T.GdkScreen) *T.GtkWidget

	Gtk_invisible_set_screen func(
		invisible *T.GtkInvisible,
		screen *T.GdkScreen)

	Gtk_invisible_get_screen func(
		invisible *T.GtkInvisible) *T.GdkScreen

	Gtk_layout_get_type func() T.GType

	Gtk_layout_new func(
		hadjustment *T.GtkAdjustment,
		vadjustment *T.GtkAdjustment) *T.GtkWidget

	Gtk_layout_get_bin_window func(
		layout *T.GtkLayout) *T.GdkWindow

	Gtk_layout_put func(
		layout *T.GtkLayout,
		child_widget *T.GtkWidget,
		x T.Gint,
		y T.Gint)

	Gtk_layout_move func(
		layout *T.GtkLayout,
		child_widget *T.GtkWidget,
		x T.Gint,
		y T.Gint)

	Gtk_layout_set_size func(
		layout *T.GtkLayout,
		width T.Guint,
		height T.Guint)

	Gtk_layout_get_size func(
		layout *T.GtkLayout,
		width *T.Guint,
		height *T.Guint)

	Gtk_layout_get_hadjustment func(
		layout *T.GtkLayout) *T.GtkAdjustment

	Gtk_layout_get_vadjustment func(
		layout *T.GtkLayout) *T.GtkAdjustment

	Gtk_layout_set_hadjustment func(
		layout *T.GtkLayout,
		adjustment *T.GtkAdjustment)

	Gtk_layout_set_vadjustment func(
		layout *T.GtkLayout,
		adjustment *T.GtkAdjustment)

	Gtk_layout_freeze func(
		layout *T.GtkLayout)

	Gtk_layout_thaw func(
		layout *T.GtkLayout)

	Gtk_link_button_get_type func() T.GType

	Gtk_link_button_new func(
		uri string) *T.GtkWidget

	Gtk_link_button_new_with_label func(
		uri string,
		label string) *T.GtkWidget

	Gtk_link_button_get_uri func(
		link_button *T.GtkLinkButton) string

	Gtk_link_button_set_uri func(
		link_button *T.GtkLinkButton,
		uri string)

	Gtk_link_button_set_uri_hook func(
		f T.GtkLinkButtonUriFunc,
		dataGpointer,
		destroy T.GDestroyNotify) T.GtkLinkButtonUriFunc

	Gtk_link_button_get_visited func(
		link_button *T.GtkLinkButton) T.Gboolean

	Gtk_link_button_set_visited func(
		link_button *T.GtkLinkButton,
		visited T.Gboolean)

	Gtk_check_version func(
		required_major T.Guint,
		required_minor T.Guint,
		required_micro T.Guint) string

	Gtk_parse_args func(
		argc *int,
		argv ***T.Char) T.Gboolean

	Gtk_init func(
		argc *int,
		argv ***T.Char)

	Gtk_init_check func(
		argc *int,
		argv ***T.Char) T.Gboolean

	Gtk_init_with_args func(
		argc *int,
		argv ***T.Char,
		parameter_string string,
		entries *T.GOptionEntry,
		translation_domain string,
		error **T.GError) T.Gboolean

	Gtk_get_option_group func(
		open_default_display T.Gboolean) *T.GOptionGroup

	Gtk_init_abi_check func(
		argc *int,
		argv ***T.Char,
		num_checks int,
		sizeof_GtkWindow T.Size_t,
		sizeof_GtkBox T.Size_t)

	Gtk_init_check_abi_check func(
		argc *int,
		argv ***T.Char,
		num_checks int,
		sizeof_GtkWindow T.Size_t,
		sizeof_GtkBox T.Size_t) T.Gboolean

	Gtk_exit func(
		error_code T.Gint)

	Gtk_set_locale func() string

	Gtk_disable_setlocale func()

	Gtk_get_default_language func() *T.PangoLanguage

	Gtk_events_pending func() T.Gboolean

	Gtk_main_do_event func(
		event *T.GdkEvent)

	Gtk_main func()

	Gtk_main_level func() T.Guint

	Gtk_main_quit func()

	Gtk_main_iteration func() T.Gboolean

	Gtk_main_iteration_do func(
		blocking T.Gboolean) T.Gboolean

	Gtk_true func() T.Gboolean

	Gtk_false func() T.Gboolean

	Gtk_grab_add func(
		widget *T.GtkWidget)

	Gtk_grab_get_current func() *T.GtkWidget

	Gtk_grab_remove func(
		widget *T.GtkWidget)

	Gtk_init_add func(
		function T.GtkFunction,
		data T.Gpointer)

	Gtk_quit_add_destroy func(
		main_level T.Guint,
		object *T.GtkObject)

	Gtk_quit_add func(
		main_level T.Guint,
		function T.GtkFunction,
		data T.Gpointer) T.Guint

	Gtk_quit_add_full func(
		main_level T.Guint,
		function T.GtkFunction,
		marshal T.GtkCallbackMarshal,
		dataGpointer,
		destroy T.GDestroyNotify) T.Guint

	Gtk_quit_remove func(
		quit_handler_id T.Guint)

	Gtk_quit_remove_by_data func(
		data T.Gpointer)

	Gtk_timeout_add func(
		interval T.Guint32,
		function T.GtkFunction,
		data T.Gpointer) T.Guint

	Gtk_timeout_add_full func(
		interval T.Guint32,
		function T.GtkFunction,
		marshal T.GtkCallbackMarshal,
		dataGpointer,
		destroy T.GDestroyNotify) T.Guint

	Gtk_timeout_remove func(
		timeout_handler_id T.Guint)

	Gtk_idle_add func(
		function T.GtkFunction,
		data T.Gpointer) T.Guint

	Gtk_idle_add_priority func(
		priority T.Gint,
		function T.GtkFunction,
		data T.Gpointer) T.Guint

	Gtk_idle_add_full func(
		priority T.Gint,
		function T.GtkFunction,
		marshal T.GtkCallbackMarshal,
		dataGpointer,
		destroy T.GDestroyNotify) T.Guint

	Gtk_idle_remove func(
		idle_handler_id T.Guint)

	Gtk_idle_remove_by_data func(
		data T.Gpointer)

	Gtk_input_add_full func(
		source T.Gint,
		condition T.GdkInputCondition,
		function T.GdkInputFunction,
		marshal T.GtkCallbackMarshal,
		dataGpointer,
		destroy T.GDestroyNotify) T.Guint

	Gtk_input_remove func(
		input_handler_id T.Guint)

	Gtk_key_snooper_install func(
		snooper T.GtkKeySnoopFunc,
		func_data T.Gpointer) T.Guint

	Gtk_key_snooper_remove func(
		snooper_handler_id T.Guint)

	Gtk_get_current_event func() *T.GdkEvent

	Gtk_get_current_event_time func() T.Guint32

	Gtk_get_current_event_state func(
		state *T.GdkModifierType) T.Gboolean

	Gtk_get_event_widget func(
		event *T.GdkEvent) *T.GtkWidget

	Gtk_propagate_event func(
		widget *T.GtkWidget,
		event *T.GdkEvent)

	Gtk_menu_bar_get_type func() T.GType

	Gtk_menu_bar_new func() *T.GtkWidget

	Gtk_menu_bar_get_pack_direction func(
		menubar *T.GtkMenuBar) T.GtkPackDirection

	Gtk_menu_bar_set_pack_direction func(
		menubar *T.GtkMenuBar,
		pack_dir T.GtkPackDirection)

	Gtk_menu_bar_get_child_pack_direction func(
		menubar *T.GtkMenuBar) T.GtkPackDirection

	Gtk_menu_bar_set_child_pack_direction func(
		menubar *T.GtkMenuBar,
		child_pack_dir T.GtkPackDirection)

	Gtk_tooltips_get_type func() T.GType

	Gtk_tooltips_new func() *T.GtkTooltips

	Gtk_tooltips_enable func(
		tooltips *T.GtkTooltips)

	Gtk_tooltips_disable func(
		tooltips *T.GtkTooltips)

	Gtk_tooltips_set_delay func(
		tooltips *T.GtkTooltips,
		delay T.Guint)

	Gtk_tooltips_set_tip func(
		tooltips *T.GtkTooltips,
		widget *T.GtkWidget,
		tip_text string,
		tip_private string)

	Gtk_tooltips_data_get func(
		widget *T.GtkWidget) *T.GtkTooltipsData

	Gtk_tooltips_force_window func(
		tooltips *T.GtkTooltips)

	Gtk_tooltips_get_info_from_tip_window func(
		tip_window *T.GtkWindow,
		tooltips **T.GtkTooltips,
		current_widget **T.GtkWidget) T.Gboolean

	Gtk_size_group_get_type func() T.GType

	Gtk_size_group_new func(
		mode T.GtkSizeGroupMode) *T.GtkSizeGroup

	Gtk_size_group_set_mode func(
		size_group *T.GtkSizeGroup,
		mode T.GtkSizeGroupMode)

	Gtk_size_group_get_mode func(
		size_group *T.GtkSizeGroup) T.GtkSizeGroupMode

	Gtk_size_group_set_ignore_hidden func(
		size_group *T.GtkSizeGroup,
		ignore_hidden T.Gboolean)

	Gtk_size_group_get_ignore_hidden func(
		size_group *T.GtkSizeGroup) T.Gboolean

	Gtk_size_group_add_widget func(
		size_group *T.GtkSizeGroup,
		widget *T.GtkWidget)

	Gtk_size_group_remove_widget func(
		size_group *T.GtkSizeGroup,
		widget *T.GtkWidget)

	Gtk_size_group_get_widgets func(
		size_group *T.GtkSizeGroup) *T.GSList

	Gtk_tool_item_get_type func() T.GType

	Gtk_tool_item_new func() *T.GtkToolItem

	Gtk_tool_item_set_homogeneous func(
		tool_item *T.GtkToolItem,
		homogeneous T.Gboolean)

	Gtk_tool_item_get_homogeneous func(
		tool_item *T.GtkToolItem) T.Gboolean

	Gtk_tool_item_set_expand func(
		tool_item *T.GtkToolItem,
		expand T.Gboolean)

	Gtk_tool_item_get_expand func(
		tool_item *T.GtkToolItem) T.Gboolean

	Gtk_tool_item_set_tooltip func(
		tool_item *T.GtkToolItem,
		tooltips *T.GtkTooltips,
		tip_text string,
		tip_private string)

	Gtk_tool_item_set_tooltip_text func(
		tool_item *T.GtkToolItem,
		text string)

	Gtk_tool_item_set_tooltip_markup func(
		tool_item *T.GtkToolItem,
		markup string)

	Gtk_tool_item_set_use_drag_window func(
		tool_item *T.GtkToolItem,
		use_drag_window T.Gboolean)

	Gtk_tool_item_get_use_drag_window func(
		tool_item *T.GtkToolItem) T.Gboolean

	Gtk_tool_item_set_visible_horizontal func(
		tool_item *T.GtkToolItem,
		visible_horizontal T.Gboolean)

	Gtk_tool_item_get_visible_horizontal func(
		tool_item *T.GtkToolItem) T.Gboolean

	Gtk_tool_item_set_visible_vertical func(
		tool_item *T.GtkToolItem,
		visible_vertical T.Gboolean)

	Gtk_tool_item_get_visible_vertical func(
		tool_item *T.GtkToolItem) T.Gboolean

	Gtk_tool_item_get_is_important func(
		tool_item *T.GtkToolItem) T.Gboolean

	Gtk_tool_item_set_is_important func(
		tool_item *T.GtkToolItem,
		is_important T.Gboolean)

	Gtk_tool_item_get_ellipsize_mode func(
		tool_item *T.GtkToolItem) T.PangoEllipsizeMode

	Gtk_tool_item_get_icon_size func(
		tool_item *T.GtkToolItem) T.GtkIconSize

	Gtk_tool_item_get_orientation func(
		tool_item *T.GtkToolItem) T.GtkOrientation

	Gtk_tool_item_get_toolbar_style func(
		tool_item *T.GtkToolItem) T.GtkToolbarStyle

	Gtk_tool_item_get_relief_style func(
		tool_item *T.GtkToolItem) T.GtkReliefStyle

	Gtk_tool_item_get_text_alignment func(
		tool_item *T.GtkToolItem) T.Gfloat

	Gtk_tool_item_get_text_orientation func(
		tool_item *T.GtkToolItem) T.GtkOrientation

	Gtk_tool_item_get_text_size_group func(
		tool_item *T.GtkToolItem) *T.GtkSizeGroup

	Gtk_tool_item_retrieve_proxy_menu_item func(
		tool_item *T.GtkToolItem) *T.GtkWidget

	Gtk_tool_item_get_proxy_menu_item func(
		tool_item *T.GtkToolItem,
		menu_item_id string) *T.GtkWidget

	Gtk_tool_item_set_proxy_menu_item func(
		tool_item *T.GtkToolItem,
		menu_item_id string,
		menu_item *T.GtkWidget)

	Gtk_tool_item_rebuild_menu func(
		tool_item *T.GtkToolItem)

	Gtk_tool_item_toolbar_reconfigured func(
		tool_item *T.GtkToolItem)

	Gtk_tool_button_get_type func() T.GType

	Gtk_tool_button_new func(
		icon_widget *T.GtkWidget,
		label string) *T.GtkToolItem

	Gtk_tool_button_new_from_stock func(
		stock_id string) *T.GtkToolItem

	Gtk_tool_button_set_label func(
		button *T.GtkToolButton,
		label string)

	Gtk_tool_button_get_label func(
		button *T.GtkToolButton) string

	Gtk_tool_button_set_use_underline func(
		button *T.GtkToolButton,
		use_underline T.Gboolean)

	Gtk_tool_button_get_use_underline func(
		button *T.GtkToolButton) T.Gboolean

	Gtk_tool_button_set_stock_id func(
		button *T.GtkToolButton,
		stock_id string)

	Gtk_tool_button_get_stock_id func(
		button *T.GtkToolButton) string

	Gtk_tool_button_set_icon_name func(
		button *T.GtkToolButton,
		icon_name string)

	Gtk_tool_button_get_icon_name func(
		button *T.GtkToolButton) string

	Gtk_tool_button_set_icon_widget func(
		button *T.GtkToolButton,
		icon_widget *T.GtkWidget)

	Gtk_tool_button_get_icon_widget func(
		button *T.GtkToolButton) *T.GtkWidget

	Gtk_tool_button_set_label_widget func(
		button *T.GtkToolButton,
		label_widget *T.GtkWidget)

	Gtk_tool_button_get_label_widget func(
		button *T.GtkToolButton) *T.GtkWidget

	Gtk_menu_tool_button_get_type func() T.GType

	Gtk_menu_tool_button_new func(
		icon_widget *T.GtkWidget,
		label string) *T.GtkToolItem

	Gtk_menu_tool_button_new_from_stock func(
		stock_id string) *T.GtkToolItem

	Gtk_menu_tool_button_set_menu func(
		button *T.GtkMenuToolButton,
		menu *T.GtkWidget)

	Gtk_menu_tool_button_get_menu func(
		button *T.GtkMenuToolButton) *T.GtkWidget

	Gtk_menu_tool_button_set_arrow_tooltip func(
		button *T.GtkMenuToolButton,
		tooltips *T.GtkTooltips,
		tip_text string,
		tip_private string)

	Gtk_menu_tool_button_set_arrow_tooltip_text func(
		button *T.GtkMenuToolButton,
		text string)

	Gtk_menu_tool_button_set_arrow_tooltip_markup func(
		button *T.GtkMenuToolButton,
		markup string)

	Gtk_message_dialog_get_type func() T.GType

	Gtk_message_dialog_new func(
		parent *T.GtkWindow,
		flags T.GtkDialogFlags,
		t T.GtkMessageType,
		buttons T.GtkButtonsType,
		message_format string,
		v ...VArg) *T.GtkWidget

	Gtk_message_dialog_new_with_markup func(
		parent *T.GtkWindow,
		flags T.GtkDialogFlags,
		t T.GtkMessageType,
		buttons T.GtkButtonsType,
		message_format string,
		v ...VArg) *T.GtkWidget

	Gtk_message_dialog_set_image func(
		dialog *T.GtkMessageDialog,
		image *T.GtkWidget)

	Gtk_message_dialog_get_image func(
		dialog *T.GtkMessageDialog) *T.GtkWidget

	Gtk_message_dialog_set_markup func(
		message_dialog *T.GtkMessageDialog,
		str string)

	Gtk_message_dialog_format_secondary_text func(
		message_dialog *T.GtkMessageDialog,
		message_format string,
		v ...VArg)

	Gtk_message_dialog_format_secondary_markup func(
		message_dialog *T.GtkMessageDialog,
		message_format string,
		v ...VArg)

	Gtk_message_dialog_get_message_area func(
		message_dialog *T.GtkMessageDialog) *T.GtkWidget

	Gtk_mount_operation_get_type func() T.GType

	Gtk_mount_operation_new func(
		parent *T.GtkWindow) *T.GMountOperation

	Gtk_mount_operation_is_showing func(
		op *T.GtkMountOperation) T.Gboolean

	Gtk_mount_operation_set_parent func(
		op *T.GtkMountOperation,
		parent *T.GtkWindow)

	Gtk_mount_operation_get_parent func(
		op *T.GtkMountOperation) *T.GtkWindow

	Gtk_mount_operation_set_screen func(
		op *T.GtkMountOperation,
		screen *T.GdkScreen)

	Gtk_mount_operation_get_screen func(
		op *T.GtkMountOperation) *T.GdkScreen

	Gtk_notebook_get_type func() T.GType

	Gtk_notebook_new func() *T.GtkWidget

	Gtk_notebook_append_page func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tab_label *T.GtkWidget) T.Gint

	Gtk_notebook_append_page_menu func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tab_label *T.GtkWidget,
		menu_label *T.GtkWidget) T.Gint

	Gtk_notebook_prepend_page func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tab_label *T.GtkWidget) T.Gint

	Gtk_notebook_prepend_page_menu func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tab_label *T.GtkWidget,
		menu_label *T.GtkWidget) T.Gint

	Gtk_notebook_insert_page func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tab_label *T.GtkWidget,
		position T.Gint) T.Gint

	Gtk_notebook_insert_page_menu func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tab_label *T.GtkWidget,
		menu_label *T.GtkWidget,
		position T.Gint) T.Gint

	Gtk_notebook_remove_page func(
		notebook *T.GtkNotebook,
		page_num T.Gint)

	Gtk_notebook_set_window_creation_hook func(
		f T.GtkNotebookWindowCreationFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_notebook_set_group_id func(
		notebook *T.GtkNotebook,
		group_id T.Gint)

	Gtk_notebook_get_group_id func(
		notebook *T.GtkNotebook) T.Gint

	Gtk_notebook_set_group func(
		notebook *T.GtkNotebook,
		group T.Gpointer)

	Gtk_notebook_get_group func(
		notebook *T.GtkNotebook) T.Gpointer

	Gtk_notebook_set_group_name func(
		notebook *T.GtkNotebook,
		group_name string)

	Gtk_notebook_get_group_name func(
		notebook *T.GtkNotebook) string

	Gtk_notebook_get_current_page func(
		notebook *T.GtkNotebook) T.Gint

	Gtk_notebook_get_nth_page func(
		notebook *T.GtkNotebook,
		page_num T.Gint) *T.GtkWidget

	Gtk_notebook_get_n_pages func(
		notebook *T.GtkNotebook) T.Gint

	Gtk_notebook_page_num func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) T.Gint

	Gtk_notebook_set_current_page func(
		notebook *T.GtkNotebook,
		page_num T.Gint)

	Gtk_notebook_next_page func(
		notebook *T.GtkNotebook)

	Gtk_notebook_prev_page func(
		notebook *T.GtkNotebook)

	Gtk_notebook_set_show_border func(
		notebook *T.GtkNotebook,
		show_border T.Gboolean)

	Gtk_notebook_get_show_border func(
		notebook *T.GtkNotebook) T.Gboolean

	Gtk_notebook_set_show_tabs func(
		notebook *T.GtkNotebook,
		show_tabs T.Gboolean)

	Gtk_notebook_get_show_tabs func(
		notebook *T.GtkNotebook) T.Gboolean

	Gtk_notebook_set_tab_pos func(
		notebook *T.GtkNotebook,
		pos T.GtkPositionType)

	Gtk_notebook_get_tab_pos func(
		notebook *T.GtkNotebook) T.GtkPositionType

	Gtk_notebook_set_homogeneous_tabs func(
		notebook *T.GtkNotebook,
		homogeneous T.Gboolean)

	Gtk_notebook_set_tab_border func(
		notebook *T.GtkNotebook,
		border_width T.Guint)

	Gtk_notebook_set_tab_hborder func(
		notebook *T.GtkNotebook,
		tab_hborder T.Guint)

	Gtk_notebook_set_tab_vborder func(
		notebook *T.GtkNotebook,
		tab_vborder T.Guint)

	Gtk_notebook_set_scrollable func(
		notebook *T.GtkNotebook,
		scrollable T.Gboolean)

	Gtk_notebook_get_scrollable func(
		notebook *T.GtkNotebook) T.Gboolean

	Gtk_notebook_get_tab_hborder func(
		notebook *T.GtkNotebook) T.Guint16

	Gtk_notebook_get_tab_vborder func(
		notebook *T.GtkNotebook) T.Guint16

	Gtk_notebook_popup_enable func(
		notebook *T.GtkNotebook)

	Gtk_notebook_popup_disable func(
		notebook *T.GtkNotebook)

	Gtk_notebook_get_tab_label func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) *T.GtkWidget

	Gtk_notebook_set_tab_label func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tab_label *T.GtkWidget)

	Gtk_notebook_set_tab_label_text func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		tab_text string)

	Gtk_notebook_get_tab_label_text func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) string

	Gtk_notebook_get_menu_label func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) *T.GtkWidget

	Gtk_notebook_set_menu_label func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		menu_label *T.GtkWidget)

	Gtk_notebook_set_menu_label_text func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		menu_text string)

	Gtk_notebook_get_menu_label_text func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) string

	Gtk_notebook_query_tab_label_packing func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		expand *T.Gboolean,
		fill *T.Gboolean,
		pack_type *T.GtkPackType)

	Gtk_notebook_set_tab_label_packing func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		expand T.Gboolean,
		fill T.Gboolean,
		pack_type T.GtkPackType)

	Gtk_notebook_reorder_child func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		position T.Gint)

	Gtk_notebook_get_tab_reorderable func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) T.Gboolean

	Gtk_notebook_set_tab_reorderable func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		reorderable T.Gboolean)

	Gtk_notebook_get_tab_detachable func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget) T.Gboolean

	Gtk_notebook_set_tab_detachable func(
		notebook *T.GtkNotebook,
		child *T.GtkWidget,
		detachable T.Gboolean)

	Gtk_notebook_get_action_widget func(
		notebook *T.GtkNotebook,
		pack_type T.GtkPackType) *T.GtkWidget

	Gtk_notebook_set_action_widget func(
		notebook *T.GtkNotebook,
		widget *T.GtkWidget,
		pack_type T.GtkPackType)

	Gtk_offscreen_window_get_type func() T.GType

	Gtk_offscreen_window_new func() *T.GtkWidget

	Gtk_offscreen_window_get_pixmap func(
		offscreen *T.GtkOffscreenWindow) *T.GdkPixmap

	Gtk_offscreen_window_get_pixbuf func(
		offscreen *T.GtkOffscreenWindow) *T.GdkPixbuf

	Gtk_orientable_get_type func() T.GType

	Gtk_orientable_set_orientation func(
		orientable *T.GtkOrientable,
		orientation T.GtkOrientation)

	Gtk_orientable_get_orientation func(
		orientable *T.GtkOrientable) T.GtkOrientation

	Gtk_paper_size_get_type func() T.GType

	Gtk_paper_size_new func(
		name string) *T.GtkPaperSize

	Gtk_paper_size_new_from_ppd func(
		ppd_name string,
		ppd_display_name string,
		width T.Gdouble,
		height T.Gdouble) *T.GtkPaperSize

	Gtk_paper_size_new_custom func(
		name string,
		display_name string,
		width T.Gdouble,
		height T.Gdouble,
		unit T.GtkUnit) *T.GtkPaperSize

	Gtk_paper_size_copy func(
		other *T.GtkPaperSize) *T.GtkPaperSize

	Gtk_paper_size_free func(
		size *T.GtkPaperSize)

	Gtk_paper_size_is_equal func(
		size1 *T.GtkPaperSize,
		size2 *T.GtkPaperSize) T.Gboolean

	Gtk_paper_size_get_paper_sizes func(
		include_custom T.Gboolean) *T.GList

	Gtk_paper_size_get_name func(
		size *T.GtkPaperSize) string

	Gtk_paper_size_get_display_name func(
		size *T.GtkPaperSize) string

	Gtk_paper_size_get_ppd_name func(
		size *T.GtkPaperSize) string

	Gtk_paper_size_get_width func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) T.Gdouble

	Gtk_paper_size_get_height func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) T.Gdouble

	Gtk_paper_size_is_custom func(
		size *T.GtkPaperSize) T.Gboolean

	Gtk_paper_size_set_size func(
		size *T.GtkPaperSize,
		width T.Gdouble,
		height T.Gdouble,
		unit T.GtkUnit)

	Gtk_paper_size_get_default_top_margin func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) T.Gdouble

	Gtk_paper_size_get_default_bottom_margin func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) T.Gdouble

	Gtk_paper_size_get_default_left_margin func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) T.Gdouble

	Gtk_paper_size_get_default_right_margin func(
		size *T.GtkPaperSize,
		unit T.GtkUnit) T.Gdouble

	Gtk_paper_size_get_default func() string

	Gtk_paper_size_new_from_key_file func(
		key_file *T.GKeyFile,
		group_name string,
		error **T.GError) *T.GtkPaperSize

	Gtk_paper_size_to_key_file func(
		size *T.GtkPaperSize,
		key_file *T.GKeyFile,
		group_name string)

	Gtk_page_setup_get_type func() T.GType

	Gtk_page_setup_new func() *T.GtkPageSetup

	Gtk_page_setup_copy func(
		other *T.GtkPageSetup) *T.GtkPageSetup

	Gtk_page_setup_get_orientation func(
		setup *T.GtkPageSetup) T.GtkPageOrientation

	Gtk_page_setup_set_orientation func(
		setup *T.GtkPageSetup,
		orientation T.GtkPageOrientation)

	Gtk_page_setup_get_paper_size func(
		setup *T.GtkPageSetup) *T.GtkPaperSize

	Gtk_page_setup_set_paper_size func(
		setup *T.GtkPageSetup,
		size *T.GtkPaperSize)

	Gtk_page_setup_get_top_margin func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) T.Gdouble

	Gtk_page_setup_set_top_margin func(
		setup *T.GtkPageSetup,
		margin T.Gdouble,
		unit T.GtkUnit)

	Gtk_page_setup_get_bottom_margin func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) T.Gdouble

	Gtk_page_setup_set_bottom_margin func(
		setup *T.GtkPageSetup,
		margin T.Gdouble,
		unit T.GtkUnit)

	Gtk_page_setup_get_left_margin func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) T.Gdouble

	Gtk_page_setup_set_left_margin func(
		setup *T.GtkPageSetup,
		margin T.Gdouble,
		unit T.GtkUnit)

	Gtk_page_setup_get_right_margin func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) T.Gdouble

	Gtk_page_setup_set_right_margin func(
		setup *T.GtkPageSetup,
		margin T.Gdouble,
		unit T.GtkUnit)

	Gtk_page_setup_set_paper_size_and_default_margins func(
		setup *T.GtkPageSetup,
		size *T.GtkPaperSize)

	Gtk_page_setup_get_paper_width func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) T.Gdouble

	Gtk_page_setup_get_paper_height func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) T.Gdouble

	Gtk_page_setup_get_page_width func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) T.Gdouble

	Gtk_page_setup_get_page_height func(
		setup *T.GtkPageSetup,
		unit T.GtkUnit) T.Gdouble

	Gtk_page_setup_new_from_file func(
		file_name string,
		error **T.GError) *T.GtkPageSetup

	Gtk_page_setup_load_file func(
		setup *T.GtkPageSetup,
		file_name string,
		error **T.GError) T.Gboolean

	Gtk_page_setup_to_file func(
		setup *T.GtkPageSetup,
		file_name string,
		error **T.GError) T.Gboolean

	Gtk_page_setup_new_from_key_file func(
		key_file *T.GKeyFile,
		group_name string,
		error **T.GError) *T.GtkPageSetup

	Gtk_page_setup_load_key_file func(
		setup *T.GtkPageSetup,
		key_file *T.GKeyFile,
		group_name string,
		error **T.GError) T.Gboolean

	Gtk_page_setup_to_key_file func(
		setup *T.GtkPageSetup,
		key_file *T.GKeyFile,
		group_name string)

	Gtk_socket_get_type func() T.GType

	Gtk_socket_new func() *T.GtkWidget

	Gtk_socket_add_id func(
		socket_ *T.GtkSocket,
		window_id T.GdkNativeWindow)

	Gtk_socket_get_id func(
		socket_ *T.GtkSocket) T.GdkNativeWindow

	Gtk_socket_get_plug_window func(
		socket_ *T.GtkSocket) *T.GdkWindow

	Gtk_socket_steal func(
		socket_ *T.GtkSocket,
		wid T.GdkNativeWindow)

	Gtk_plug_get_type func() T.GType

	Gtk_plug_construct func(
		plug *T.GtkPlug,
		socket_id T.GdkNativeWindow)

	Gtk_plug_new func(
		socket_id T.GdkNativeWindow) *T.GtkWidget

	Gtk_plug_construct_for_display func(
		plug *T.GtkPlug,
		display *T.GdkDisplay,
		socket_id T.GdkNativeWindow)

	Gtk_plug_new_for_display func(
		display *T.GdkDisplay,
		socket_id T.GdkNativeWindow) *T.GtkWidget

	Gtk_plug_get_id func(
		plug *T.GtkPlug) T.GdkNativeWindow

	Gtk_plug_get_embedded func(
		plug *T.GtkPlug) T.Gboolean

	Gtk_plug_get_socket_window func(
		plug *T.GtkPlug) *T.GdkWindow

	Gtk_print_context_get_type func() T.GType

	Gtk_print_context_get_cairo_context func(
		context *T.GtkPrintContext) *T.Cairo_t

	Gtk_print_context_get_page_setup func(
		context *T.GtkPrintContext) *T.GtkPageSetup

	Gtk_print_context_get_width func(
		context *T.GtkPrintContext) T.Gdouble

	Gtk_print_context_get_height func(
		context *T.GtkPrintContext) T.Gdouble

	Gtk_print_context_get_dpi_x func(
		context *T.GtkPrintContext) T.Gdouble

	Gtk_print_context_get_dpi_y func(
		context *T.GtkPrintContext) T.Gdouble

	Gtk_print_context_get_hard_margins func(
		context *T.GtkPrintContext,
		top *T.Gdouble,
		bottom *T.Gdouble,
		left *T.Gdouble,
		right *T.Gdouble) T.Gboolean

	Gtk_print_context_get_pango_fontmap func(
		context *T.GtkPrintContext) *T.PangoFontMap

	Gtk_print_context_create_pango_context func(
		context *T.GtkPrintContext) *T.PangoContext

	Gtk_print_context_create_pango_layout func(
		context *T.GtkPrintContext) *T.PangoLayout

	Gtk_print_settings_get_type func() T.GType

	Gtk_print_settings_new func() *T.GtkPrintSettings

	Gtk_print_settings_copy func(
		other *T.GtkPrintSettings) *T.GtkPrintSettings

	Gtk_print_settings_new_from_file func(
		file_name string,
		error **T.GError) *T.GtkPrintSettings

	Gtk_print_settings_load_file func(
		settings *T.GtkPrintSettings,
		file_name string,
		error **T.GError) T.Gboolean

	Gtk_print_settings_to_file func(
		settings *T.GtkPrintSettings,
		file_name string,
		error **T.GError) T.Gboolean

	Gtk_print_settings_new_from_key_file func(
		key_file *T.GKeyFile,
		group_name string,
		error **T.GError) *T.GtkPrintSettings

	Gtk_print_settings_load_key_file func(
		settings *T.GtkPrintSettings,
		key_file *T.GKeyFile,
		group_name string,
		error **T.GError) T.Gboolean

	Gtk_print_settings_to_key_file func(
		settings *T.GtkPrintSettings,
		key_file *T.GKeyFile,
		group_name string)

	Gtk_print_settings_has_key func(
		settings *T.GtkPrintSettings,
		key string) T.Gboolean

	Gtk_print_settings_get func(
		settings *T.GtkPrintSettings,
		key string) string

	Gtk_print_settings_set func(
		settings *T.GtkPrintSettings,
		key string,
		value string)

	Gtk_print_settings_unset func(
		settings *T.GtkPrintSettings,
		key string)

	Gtk_print_settings_foreach func(
		settings *T.GtkPrintSettings,
		f T.GtkPrintSettingsFunc,
		user_data T.Gpointer)

	Gtk_print_settings_get_bool func(
		settings *T.GtkPrintSettings,
		key string) T.Gboolean

	Gtk_print_settings_set_bool func(
		settings *T.GtkPrintSettings,
		key string,
		value T.Gboolean)

	Gtk_print_settings_get_double func(
		settings *T.GtkPrintSettings,
		key string) T.Gdouble

	Gtk_print_settings_get_double_with_default func(
		settings *T.GtkPrintSettings,
		key string,
		def T.Gdouble) T.Gdouble

	Gtk_print_settings_set_double func(
		settings *T.GtkPrintSettings,
		key string,
		value T.Gdouble)

	Gtk_print_settings_get_length func(
		settings *T.GtkPrintSettings,
		key string,
		unit T.GtkUnit) T.Gdouble

	Gtk_print_settings_set_length func(
		settings *T.GtkPrintSettings,
		key string,
		value T.Gdouble,
		unit T.GtkUnit)

	Gtk_print_settings_get_int func(
		settings *T.GtkPrintSettings,
		key string) T.Gint

	Gtk_print_settings_get_int_with_default func(
		settings *T.GtkPrintSettings,
		key string,
		def T.Gint) T.Gint

	Gtk_print_settings_set_int func(
		settings *T.GtkPrintSettings,
		key string,
		value T.Gint)

	Gtk_print_settings_get_printer func(
		settings *T.GtkPrintSettings) string

	Gtk_print_settings_set_printer func(
		settings *T.GtkPrintSettings,
		printer string)

	Gtk_print_settings_get_orientation func(
		settings *T.GtkPrintSettings) T.GtkPageOrientation

	Gtk_print_settings_set_orientation func(
		settings *T.GtkPrintSettings,
		orientation T.GtkPageOrientation)

	Gtk_print_settings_get_paper_size func(
		settings *T.GtkPrintSettings) *T.GtkPaperSize

	Gtk_print_settings_set_paper_size func(
		settings *T.GtkPrintSettings,
		paper_size *T.GtkPaperSize)

	Gtk_print_settings_get_paper_width func(
		settings *T.GtkPrintSettings,
		unit T.GtkUnit) T.Gdouble

	Gtk_print_settings_set_paper_width func(
		settings *T.GtkPrintSettings,
		width T.Gdouble,
		unit T.GtkUnit)

	Gtk_print_settings_get_paper_height func(
		settings *T.GtkPrintSettings,
		unit T.GtkUnit) T.Gdouble

	Gtk_print_settings_set_paper_height func(
		settings *T.GtkPrintSettings,
		height T.Gdouble,
		unit T.GtkUnit)

	Gtk_print_settings_get_use_color func(
		settings *T.GtkPrintSettings) T.Gboolean

	Gtk_print_settings_set_use_color func(
		settings *T.GtkPrintSettings,
		use_color T.Gboolean)

	Gtk_print_settings_get_collate func(
		settings *T.GtkPrintSettings) T.Gboolean

	Gtk_print_settings_set_collate func(
		settings *T.GtkPrintSettings,
		collate T.Gboolean)

	Gtk_print_settings_get_reverse func(
		settings *T.GtkPrintSettings) T.Gboolean

	Gtk_print_settings_set_reverse func(
		settings *T.GtkPrintSettings,
		reverse T.Gboolean)

	Gtk_print_settings_get_duplex func(
		settings *T.GtkPrintSettings) T.GtkPrintDuplex

	Gtk_print_settings_set_duplex func(
		settings *T.GtkPrintSettings,
		duplex T.GtkPrintDuplex)

	Gtk_print_settings_get_quality func(
		settings *T.GtkPrintSettings) T.GtkPrintQuality

	Gtk_print_settings_set_quality func(
		settings *T.GtkPrintSettings,
		quality T.GtkPrintQuality)

	Gtk_print_settings_get_n_copies func(
		settings *T.GtkPrintSettings) T.Gint

	Gtk_print_settings_set_n_copies func(
		settings *T.GtkPrintSettings,
		num_copies T.Gint)

	Gtk_print_settings_get_number_up func(
		settings *T.GtkPrintSettings) T.Gint

	Gtk_print_settings_set_number_up func(
		settings *T.GtkPrintSettings,
		number_up T.Gint)

	Gtk_print_settings_get_number_up_layout func(
		settings *T.GtkPrintSettings) T.GtkNumberUpLayout

	Gtk_print_settings_set_number_up_layout func(
		settings *T.GtkPrintSettings,
		number_up_layout T.GtkNumberUpLayout)

	Gtk_print_settings_get_resolution func(
		settings *T.GtkPrintSettings) T.Gint

	Gtk_print_settings_set_resolution func(
		settings *T.GtkPrintSettings,
		resolution T.Gint)

	Gtk_print_settings_get_resolution_x func(
		settings *T.GtkPrintSettings) T.Gint

	Gtk_print_settings_get_resolution_y func(
		settings *T.GtkPrintSettings) T.Gint

	Gtk_print_settings_set_resolution_xy func(
		settings *T.GtkPrintSettings,
		resolution_x T.Gint,
		resolution_y T.Gint)

	Gtk_print_settings_get_printer_lpi func(
		settings *T.GtkPrintSettings) T.Gdouble

	Gtk_print_settings_set_printer_lpi func(
		settings *T.GtkPrintSettings,
		lpi T.Gdouble)

	Gtk_print_settings_get_scale func(
		settings *T.GtkPrintSettings) T.Gdouble

	Gtk_print_settings_set_scale func(
		settings *T.GtkPrintSettings,
		scale T.Gdouble)

	Gtk_print_settings_get_print_pages func(
		settings *T.GtkPrintSettings) T.GtkPrintPages

	Gtk_print_settings_set_print_pages func(
		settings *T.GtkPrintSettings,
		pages T.GtkPrintPages)

	Gtk_print_settings_get_page_ranges func(
		settings *T.GtkPrintSettings,
		num_ranges *T.Gint) *T.GtkPageRange

	Gtk_print_settings_set_page_ranges func(
		settings *T.GtkPrintSettings,
		page_ranges *T.GtkPageRange,
		num_ranges T.Gint)

	Gtk_print_settings_get_page_set func(
		settings *T.GtkPrintSettings) T.GtkPageSet

	Gtk_print_settings_set_page_set func(
		settings *T.GtkPrintSettings,
		page_set T.GtkPageSet)

	Gtk_print_settings_get_default_source func(
		settings *T.GtkPrintSettings) string

	Gtk_print_settings_set_default_source func(
		settings *T.GtkPrintSettings,
		default_source string)

	Gtk_print_settings_get_media_type func(
		settings *T.GtkPrintSettings) string

	Gtk_print_settings_set_media_type func(
		settings *T.GtkPrintSettings,
		media_type string)

	Gtk_print_settings_get_dither func(
		settings *T.GtkPrintSettings) string

	Gtk_print_settings_set_dither func(
		settings *T.GtkPrintSettings,
		dither string)

	Gtk_print_settings_get_finishings func(
		settings *T.GtkPrintSettings) string

	Gtk_print_settings_set_finishings func(
		settings *T.GtkPrintSettings,
		finishings string)

	Gtk_print_settings_get_output_bin func(
		settings *T.GtkPrintSettings) string

	Gtk_print_settings_set_output_bin func(
		settings *T.GtkPrintSettings,
		output_bin string)

	Gtk_print_operation_preview_get_type func() T.GType

	Gtk_print_operation_preview_render_page func(
		preview *T.GtkPrintOperationPreview,
		page_nr T.Gint)

	Gtk_print_operation_preview_end_preview func(
		preview *T.GtkPrintOperationPreview)

	Gtk_print_operation_preview_is_selected func(
		preview *T.GtkPrintOperationPreview,
		page_nr T.Gint) T.Gboolean

	Gtk_print_error_quark func() T.GQuark

	Gtk_print_operation_get_type func() T.GType

	Gtk_print_operation_new func() *T.GtkPrintOperation

	Gtk_print_operation_set_default_page_setup func(
		op *T.GtkPrintOperation,
		default_page_setup *T.GtkPageSetup)

	Gtk_print_operation_get_default_page_setup func(
		op *T.GtkPrintOperation) *T.GtkPageSetup

	Gtk_print_operation_set_print_settings func(
		op *T.GtkPrintOperation,
		print_settings *T.GtkPrintSettings)

	Gtk_print_operation_get_print_settings func(
		op *T.GtkPrintOperation) *T.GtkPrintSettings

	Gtk_print_operation_set_job_name func(
		op *T.GtkPrintOperation,
		job_name string)

	Gtk_print_operation_set_n_pages func(
		op *T.GtkPrintOperation,
		n_pages T.Gint)

	Gtk_print_operation_set_current_page func(
		op *T.GtkPrintOperation,
		current_page T.Gint)

	Gtk_print_operation_set_use_full_page func(
		op *T.GtkPrintOperation,
		full_page T.Gboolean)

	Gtk_print_operation_set_unit func(
		op *T.GtkPrintOperation,
		unit T.GtkUnit)

	Gtk_print_operation_set_export_filename func(
		op *T.GtkPrintOperation,
		filename string)

	Gtk_print_operation_set_track_print_status func(
		op *T.GtkPrintOperation,
		track_status T.Gboolean)

	Gtk_print_operation_set_show_progress func(
		op *T.GtkPrintOperation,
		show_progress T.Gboolean)

	Gtk_print_operation_set_allow_async func(
		op *T.GtkPrintOperation,
		allow_async T.Gboolean)

	Gtk_print_operation_set_custom_tab_label func(
		op *T.GtkPrintOperation,
		label string)

	Gtk_print_operation_run func(
		op *T.GtkPrintOperation,
		action T.GtkPrintOperationAction,
		parent *T.GtkWindow,
		error **T.GError) T.GtkPrintOperationResult

	Gtk_print_operation_get_error func(
		op *T.GtkPrintOperation,
		error **T.GError)

	Gtk_print_operation_get_status func(
		op *T.GtkPrintOperation) T.GtkPrintStatus

	Gtk_print_operation_get_status_string func(
		op *T.GtkPrintOperation) string

	Gtk_print_operation_is_finished func(
		op *T.GtkPrintOperation) T.Gboolean

	Gtk_print_operation_cancel func(
		op *T.GtkPrintOperation)

	Gtk_print_operation_draw_page_finish func(
		op *T.GtkPrintOperation)

	Gtk_print_operation_set_defer_drawing func(
		op *T.GtkPrintOperation)

	Gtk_print_operation_set_support_selection func(
		op *T.GtkPrintOperation,
		support_selection T.Gboolean)

	Gtk_print_operation_get_support_selection func(
		op *T.GtkPrintOperation) T.Gboolean

	Gtk_print_operation_set_has_selection func(
		op *T.GtkPrintOperation,
		has_selection T.Gboolean)

	Gtk_print_operation_get_has_selection func(
		op *T.GtkPrintOperation) T.Gboolean

	Gtk_print_operation_set_embed_page_setup func(
		op *T.GtkPrintOperation,
		embed T.Gboolean)

	Gtk_print_operation_get_embed_page_setup func(
		op *T.GtkPrintOperation) T.Gboolean

	Gtk_print_operation_get_n_pages_to_print func(
		op *T.GtkPrintOperation) T.Gint

	Gtk_print_run_page_setup_dialog func(
		parent *T.GtkWindow,
		page_setup *T.GtkPageSetup,
		settings *T.GtkPrintSettings) *T.GtkPageSetup

	Gtk_print_run_page_setup_dialog_async func(
		parent *T.GtkWindow,
		page_setup *T.GtkPageSetup,
		settings *T.GtkPrintSettings,
		done_cb T.GtkPageSetupDoneFunc,
		data T.Gpointer)

	Gtk_progress_get_type func() T.GType

	Gtk_progress_set_show_text func(
		progress *T.GtkProgress,
		show_text T.Gboolean)

	Gtk_progress_set_text_alignment func(
		progress *T.GtkProgress,
		x_align T.Gfloat,
		y_align T.Gfloat)

	Gtk_progress_set_format_string func(
		progress *T.GtkProgress,
		format string)

	Gtk_progress_set_adjustment func(
		progress *T.GtkProgress,
		adjustment *T.GtkAdjustment)

	Gtk_progress_configure func(
		progress *T.GtkProgress,
		value T.Gdouble,
		min T.Gdouble,
		max T.Gdouble)

	Gtk_progress_set_percentage func(
		progress *T.GtkProgress,
		percentage T.Gdouble)

	Gtk_progress_set_value func(
		progress *T.GtkProgress,
		value T.Gdouble)

	Gtk_progress_get_value func(
		progress *T.GtkProgress) T.Gdouble

	Gtk_progress_set_activity_mode func(
		progress *T.GtkProgress,
		activity_mode T.Gboolean)

	Gtk_progress_get_current_text func(
		progress *T.GtkProgress) string

	Gtk_progress_get_text_from_value func(
		progress *T.GtkProgress,
		value T.Gdouble) string

	Gtk_progress_get_current_percentage func(
		progress *T.GtkProgress) T.Gdouble

	Gtk_progress_get_percentage_from_value func(
		progress *T.GtkProgress,
		value T.Gdouble) T.Gdouble

	Gtk_progress_bar_get_type func() T.GType

	Gtk_progress_bar_new func() *T.GtkWidget

	Gtk_progress_bar_pulse func(
		pbar *T.GtkProgressBar)

	Gtk_progress_bar_set_text func(
		pbar *T.GtkProgressBar,
		text string)

	Gtk_progress_bar_set_fraction func(
		pbar *T.GtkProgressBar,
		fraction T.Gdouble)

	Gtk_progress_bar_set_pulse_step func(
		pbar *T.GtkProgressBar,
		fraction T.Gdouble)

	Gtk_progress_bar_set_orientation func(
		pbar *T.GtkProgressBar,
		orientation T.GtkProgressBarOrientation)

	Gtk_progress_bar_get_text func(
		pbar *T.GtkProgressBar) string

	Gtk_progress_bar_get_fraction func(
		pbar *T.GtkProgressBar) T.Gdouble

	Gtk_progress_bar_get_pulse_step func(
		pbar *T.GtkProgressBar) T.Gdouble

	Gtk_progress_bar_get_orientation func(
		pbar *T.GtkProgressBar) T.GtkProgressBarOrientation

	Gtk_progress_bar_set_ellipsize func(
		pbar *T.GtkProgressBar,
		mode T.PangoEllipsizeMode)

	Gtk_progress_bar_get_ellipsize func(
		pbar *T.GtkProgressBar) T.PangoEllipsizeMode

	Gtk_progress_bar_new_with_adjustment func(
		adjustment *T.GtkAdjustment) *T.GtkWidget

	Gtk_progress_bar_set_bar_style func(
		pbar *T.GtkProgressBar,
		style T.GtkProgressBarStyle)

	Gtk_progress_bar_set_discrete_blocks func(
		pbar *T.GtkProgressBar,
		blocks T.Guint)

	Gtk_progress_bar_set_activity_step func(
		pbar *T.GtkProgressBar,
		step T.Guint)

	Gtk_progress_bar_set_activity_blocks func(
		pbar *T.GtkProgressBar,
		blocks T.Guint)

	Gtk_progress_bar_update func(
		pbar *T.GtkProgressBar,
		percentage T.Gdouble)

	Gtk_toggle_action_get_type func() T.GType

	Gtk_toggle_action_new func(
		name string,
		label string,
		tooltip string,
		stock_id string) *T.GtkToggleAction

	Gtk_toggle_action_toggled func(
		action *T.GtkToggleAction)

	Gtk_toggle_action_set_active func(
		action *T.GtkToggleAction,
		is_active T.Gboolean)

	Gtk_toggle_action_get_active func(
		action *T.GtkToggleAction) T.Gboolean

	Gtk_toggle_action_set_draw_as_radio func(
		action *T.GtkToggleAction,
		draw_as_radio T.Gboolean)

	Gtk_toggle_action_get_draw_as_radio func(
		action *T.GtkToggleAction) T.Gboolean

	Gtk_radio_action_get_type func() T.GType

	Gtk_radio_action_new func(
		name string,
		label string,
		tooltip string,
		stock_id string,
		value T.Gint) *T.GtkRadioAction

	Gtk_radio_action_get_group func(
		action *T.GtkRadioAction) *T.GSList

	Gtk_radio_action_set_group func(
		action *T.GtkRadioAction,
		group *T.GSList)

	Gtk_radio_action_get_current_value func(
		action *T.GtkRadioAction) T.Gint

	Gtk_radio_action_set_current_value func(
		action *T.GtkRadioAction,
		current_value T.Gint)

	Gtk_radio_button_get_type func() T.GType

	Gtk_radio_button_new func(
		group *T.GSList) *T.GtkWidget

	Gtk_radio_button_new_from_widget func(
		radio_group_member *T.GtkRadioButton) *T.GtkWidget

	Gtk_radio_button_new_with_label func(
		group *T.GSList,
		label string) *T.GtkWidget

	Gtk_radio_button_new_with_label_from_widget func(
		radio_group_member *T.GtkRadioButton,
		label string) *T.GtkWidget

	Gtk_radio_button_new_with_mnemonic func(
		group *T.GSList,
		label string) *T.GtkWidget

	Gtk_radio_button_new_with_mnemonic_from_widget func(
		radio_group_member *T.GtkRadioButton,
		label string) *T.GtkWidget

	Gtk_radio_button_get_group func(
		radio_button *T.GtkRadioButton) *T.GSList

	Gtk_radio_button_set_group func(
		radio_button *T.GtkRadioButton,
		group *T.GSList)

	Gtk_radio_menu_item_get_type func() T.GType

	Gtk_radio_menu_item_new func(
		group *T.GSList) *T.GtkWidget

	Gtk_radio_menu_item_new_with_label func(
		group *T.GSList,
		label string) *T.GtkWidget

	Gtk_radio_menu_item_new_with_mnemonic func(
		group *T.GSList,
		label string) *T.GtkWidget

	Gtk_radio_menu_item_new_from_widget func(
		group *T.GtkRadioMenuItem) *T.GtkWidget

	Gtk_radio_menu_item_new_with_mnemonic_from_widget func(
		group *T.GtkRadioMenuItem,
		label string) *T.GtkWidget

	Gtk_radio_menu_item_new_with_label_from_widget func(
		group *T.GtkRadioMenuItem,
		label string) *T.GtkWidget

	Gtk_radio_menu_item_get_group func(
		radio_menu_item *T.GtkRadioMenuItem) *T.GSList

	Gtk_radio_menu_item_set_group func(
		radio_menu_item *T.GtkRadioMenuItem,
		group *T.GSList)

	Gtk_toggle_tool_button_get_type func() T.GType

	Gtk_toggle_tool_button_new func() *T.GtkToolItem

	Gtk_toggle_tool_button_new_from_stock func(
		stock_id string) *T.GtkToolItem

	Gtk_toggle_tool_button_set_active func(
		button *T.GtkToggleToolButton,
		is_active T.Gboolean)

	Gtk_toggle_tool_button_get_active func(
		button *T.GtkToggleToolButton) T.Gboolean

	Gtk_radio_tool_button_get_type func() T.GType

	Gtk_radio_tool_button_new func(
		group *T.GSList) *T.GtkToolItem

	Gtk_radio_tool_button_new_from_stock func(
		group *T.GSList,
		stock_id string) *T.GtkToolItem

	Gtk_radio_tool_button_new_from_widget func(
		group *T.GtkRadioToolButton) *T.GtkToolItem

	Gtk_radio_tool_button_new_with_stock_from_widget func(
		group *T.GtkRadioToolButton,
		stock_id string) *T.GtkToolItem

	Gtk_radio_tool_button_get_group func(
		button *T.GtkRadioToolButton) *T.GSList

	Gtk_radio_tool_button_set_group func(
		button *T.GtkRadioToolButton,
		group *T.GSList)

	Gtk_recent_manager_error_quark func() T.GQuark

	Gtk_recent_manager_get_type func() T.GType

	Gtk_recent_manager_new func() *T.GtkRecentManager

	Gtk_recent_manager_get_default func() *T.GtkRecentManager

	Gtk_recent_manager_get_for_screen func(
		screen *T.GdkScreen) *T.GtkRecentManager

	Gtk_recent_manager_set_screen func(
		manager *T.GtkRecentManager,
		screen *T.GdkScreen)

	Gtk_recent_manager_add_item func(
		manager *T.GtkRecentManager,
		uri string) T.Gboolean

	Gtk_recent_manager_add_full func(
		manager *T.GtkRecentManager,
		uri string,
		recent_data *T.GtkRecentData) T.Gboolean

	Gtk_recent_manager_remove_item func(
		manager *T.GtkRecentManager,
		uri string,
		error **T.GError) T.Gboolean

	Gtk_recent_manager_lookup_item func(
		manager *T.GtkRecentManager,
		uri string,
		error **T.GError) *T.GtkRecentInfo

	Gtk_recent_manager_has_item func(
		manager *T.GtkRecentManager,
		uri string) T.Gboolean

	Gtk_recent_manager_move_item func(
		manager *T.GtkRecentManager,
		uri string,
		new_uri string,
		error **T.GError) T.Gboolean

	Gtk_recent_manager_set_limit func(
		manager *T.GtkRecentManager,
		limit T.Gint)

	Gtk_recent_manager_get_limit func(
		manager *T.GtkRecentManager) T.Gint

	Gtk_recent_manager_get_items func(
		manager *T.GtkRecentManager) *T.GList

	Gtk_recent_manager_purge_items func(
		manager *T.GtkRecentManager,
		error **T.GError) T.Gint

	Gtk_recent_info_get_type func() T.GType

	Gtk_recent_info_ref func(
		info *T.GtkRecentInfo) *T.GtkRecentInfo

	Gtk_recent_info_unref func(
		info *T.GtkRecentInfo)

	Gtk_recent_info_get_uri func(
		info *T.GtkRecentInfo) string

	Gtk_recent_info_get_display_name func(
		info *T.GtkRecentInfo) string

	Gtk_recent_info_get_description func(
		info *T.GtkRecentInfo) string

	Gtk_recent_info_get_mime_type func(
		info *T.GtkRecentInfo) string

	Gtk_recent_info_get_added func(
		info *T.GtkRecentInfo) T.Time_t

	Gtk_recent_info_get_modified func(
		info *T.GtkRecentInfo) T.Time_t

	Gtk_recent_info_get_visited func(
		info *T.GtkRecentInfo) T.Time_t

	Gtk_recent_info_get_private_hint func(
		info *T.GtkRecentInfo) T.Gboolean

	Gtk_recent_info_get_application_info func(
		info *T.GtkRecentInfo,
		app_name string,
		app_exec **T.Gchar,
		count *T.Guint,
		time_ *T.Time_t) T.Gboolean

	Gtk_recent_info_get_applications func(
		info *T.GtkRecentInfo,
		length *T.Gsize) **T.Gchar

	Gtk_recent_info_last_application func(
		info *T.GtkRecentInfo) string

	Gtk_recent_info_has_application func(
		info *T.GtkRecentInfo,
		app_name string) T.Gboolean

	Gtk_recent_info_get_groups func(
		info *T.GtkRecentInfo,
		length *T.Gsize) **T.Gchar

	Gtk_recent_info_has_group func(
		info *T.GtkRecentInfo,
		group_name string) T.Gboolean

	Gtk_recent_info_get_icon func(
		info *T.GtkRecentInfo,
		size T.Gint) *T.GdkPixbuf

	Gtk_recent_info_get_short_name func(
		info *T.GtkRecentInfo) string

	Gtk_recent_info_get_uri_display func(
		info *T.GtkRecentInfo) string

	Gtk_recent_info_get_age func(
		info *T.GtkRecentInfo) T.Gint

	Gtk_recent_info_is_local func(
		info *T.GtkRecentInfo) T.Gboolean

	Gtk_recent_info_exists func(
		info *T.GtkRecentInfo) T.Gboolean

	Gtk_recent_info_match func(
		info_a *T.GtkRecentInfo,
		info_b *T.GtkRecentInfo) T.Gboolean

	Gtk_recent_action_get_type func() T.GType

	Gtk_recent_action_new func(
		name string,
		label string,
		tooltip string,
		stock_id string) *T.GtkAction

	Gtk_recent_action_new_for_manager func(
		name string,
		label string,
		tooltip string,
		stock_id string,
		manager *T.GtkRecentManager) *T.GtkAction

	Gtk_recent_action_get_show_numbers func(
		action *T.GtkRecentAction) T.Gboolean

	Gtk_recent_action_set_show_numbers func(
		action *T.GtkRecentAction,
		show_numbers T.Gboolean)

	Gtk_recent_filter_get_type func() T.GType

	Gtk_recent_filter_new func() *T.GtkRecentFilter

	Gtk_recent_filter_set_name func(
		filter *T.GtkRecentFilter,
		name string)

	Gtk_recent_filter_get_name func(
		filter *T.GtkRecentFilter) string

	Gtk_recent_filter_add_mime_type func(
		filter *T.GtkRecentFilter,
		mime_type string)

	Gtk_recent_filter_add_pattern func(
		filter *T.GtkRecentFilter,
		pattern string)

	Gtk_recent_filter_add_pixbuf_formats func(
		filter *T.GtkRecentFilter)

	Gtk_recent_filter_add_application func(
		filter *T.GtkRecentFilter,
		application string)

	Gtk_recent_filter_add_group func(
		filter *T.GtkRecentFilter,
		group string)

	Gtk_recent_filter_add_age func(
		filter *T.GtkRecentFilter,
		days T.Gint)

	Gtk_recent_filter_add_custom func(
		filter *T.GtkRecentFilter,
		needed T.GtkRecentFilterFlags,
		f T.GtkRecentFilterFunc,
		dataGpointer,
		data_destroy T.GDestroyNotify)

	Gtk_recent_filter_get_needed func(
		filter *T.GtkRecentFilter) T.GtkRecentFilterFlags

	Gtk_recent_filter_filter func(
		filter *T.GtkRecentFilter,
		filter_info *T.GtkRecentFilterInfo) T.Gboolean

	Gtk_recent_chooser_error_quark func() T.GQuark

	Gtk_recent_chooser_get_type func() T.GType

	Gtk_recent_chooser_set_show_private func(
		chooser *T.GtkRecentChooser,
		show_private T.Gboolean)

	Gtk_recent_chooser_get_show_private func(
		chooser *T.GtkRecentChooser) T.Gboolean

	Gtk_recent_chooser_set_show_not_found func(
		chooser *T.GtkRecentChooser,
		show_not_found T.Gboolean)

	Gtk_recent_chooser_get_show_not_found func(
		chooser *T.GtkRecentChooser) T.Gboolean

	Gtk_recent_chooser_set_select_multiple func(
		chooser *T.GtkRecentChooser,
		select_multiple T.Gboolean)

	Gtk_recent_chooser_get_select_multiple func(
		chooser *T.GtkRecentChooser) T.Gboolean

	Gtk_recent_chooser_set_limit func(
		chooser *T.GtkRecentChooser,
		limit T.Gint)

	Gtk_recent_chooser_get_limit func(
		chooser *T.GtkRecentChooser) T.Gint

	Gtk_recent_chooser_set_local_only func(
		chooser *T.GtkRecentChooser,
		local_only T.Gboolean)

	Gtk_recent_chooser_get_local_only func(
		chooser *T.GtkRecentChooser) T.Gboolean

	Gtk_recent_chooser_set_show_tips func(
		chooser *T.GtkRecentChooser,
		show_tips T.Gboolean)

	Gtk_recent_chooser_get_show_tips func(
		chooser *T.GtkRecentChooser) T.Gboolean

	Gtk_recent_chooser_set_show_numbers func(
		chooser *T.GtkRecentChooser,
		show_numbers T.Gboolean)

	Gtk_recent_chooser_get_show_numbers func(
		chooser *T.GtkRecentChooser) T.Gboolean

	Gtk_recent_chooser_set_show_icons func(
		chooser *T.GtkRecentChooser,
		show_icons T.Gboolean)

	Gtk_recent_chooser_get_show_icons func(
		chooser *T.GtkRecentChooser) T.Gboolean

	Gtk_recent_chooser_set_sort_type func(
		chooser *T.GtkRecentChooser,
		sort_type T.GtkRecentSortType)

	Gtk_recent_chooser_get_sort_type func(
		chooser *T.GtkRecentChooser) T.GtkRecentSortType

	Gtk_recent_chooser_set_sort_func func(
		chooser *T.GtkRecentChooser,
		sort_func T.GtkRecentSortFunc,
		sort_dataGpointer,
		data_destroy T.GDestroyNotify)

	Gtk_recent_chooser_set_current_uri func(
		chooser *T.GtkRecentChooser,
		uri string,
		error **T.GError) T.Gboolean

	Gtk_recent_chooser_get_current_uri func(
		chooser *T.GtkRecentChooser) string

	Gtk_recent_chooser_get_current_item func(
		chooser *T.GtkRecentChooser) *T.GtkRecentInfo

	Gtk_recent_chooser_select_uri func(
		chooser *T.GtkRecentChooser,
		uri string,
		error **T.GError) T.Gboolean

	Gtk_recent_chooser_unselect_uri func(
		chooser *T.GtkRecentChooser,
		uri string)

	Gtk_recent_chooser_select_all func(
		chooser *T.GtkRecentChooser)

	Gtk_recent_chooser_unselect_all func(
		chooser *T.GtkRecentChooser)

	Gtk_recent_chooser_get_items func(
		chooser *T.GtkRecentChooser) *T.GList

	Gtk_recent_chooser_get_uris func(
		chooser *T.GtkRecentChooser,
		length *T.Gsize) **T.Gchar

	Gtk_recent_chooser_add_filter func(
		chooser *T.GtkRecentChooser,
		filter *T.GtkRecentFilter)

	Gtk_recent_chooser_remove_filter func(
		chooser *T.GtkRecentChooser,
		filter *T.GtkRecentFilter)

	Gtk_recent_chooser_list_filters func(
		chooser *T.GtkRecentChooser) *T.GSList

	Gtk_recent_chooser_set_filter func(
		chooser *T.GtkRecentChooser,
		filter *T.GtkRecentFilter)

	Gtk_recent_chooser_get_filter func(
		chooser *T.GtkRecentChooser) *T.GtkRecentFilter

	Gtk_recent_chooser_dialog_get_type func() T.GType

	Gtk_recent_chooser_dialog_new func(
		title string, parent *T.GtkWindow,
		first_button_text string, v ...VArg) *T.GtkWidget

	Gtk_recent_chooser_dialog_new_for_manager func(
		title string, parent *T.GtkWindow,
		manager *T.GtkRecentManager,
		first_button_text string, v ...VArg) *T.GtkWidget

	Gtk_recent_chooser_menu_get_type func() T.GType

	Gtk_recent_chooser_menu_new func() *T.GtkWidget

	Gtk_recent_chooser_menu_new_for_manager func(
		manager *T.GtkRecentManager) *T.GtkWidget

	Gtk_recent_chooser_menu_get_show_numbers func(
		menu *T.GtkRecentChooserMenu) T.Gboolean

	Gtk_recent_chooser_menu_set_show_numbers func(
		menu *T.GtkRecentChooserMenu,
		show_numbers T.Gboolean)

	Gtk_recent_chooser_widget_get_type func() T.GType

	Gtk_recent_chooser_widget_new func() *T.GtkWidget

	Gtk_recent_chooser_widget_new_for_manager func(
		manager *T.GtkRecentManager) *T.GtkWidget

	Gtk_scale_button_get_type func() T.GType

	Gtk_scale_button_new func(
		size T.GtkIconSize,
		min T.Gdouble,
		max T.Gdouble,
		step T.Gdouble,
		icons **T.Gchar) *T.GtkWidget

	Gtk_scale_button_set_icons func(
		button *T.GtkScaleButton,
		icons **T.Gchar)

	Gtk_scale_button_get_value func(
		button *T.GtkScaleButton) T.Gdouble

	Gtk_scale_button_set_value func(
		button *T.GtkScaleButton,
		value T.Gdouble)

	Gtk_scale_button_get_adjustment func(
		button *T.GtkScaleButton) *T.GtkAdjustment

	Gtk_scale_button_set_adjustment func(
		button *T.GtkScaleButton,
		adjustment *T.GtkAdjustment)

	Gtk_scale_button_get_plus_button func(
		button *T.GtkScaleButton) *T.GtkWidget

	Gtk_scale_button_get_minus_button func(
		button *T.GtkScaleButton) *T.GtkWidget

	Gtk_scale_button_get_popup func(
		button *T.GtkScaleButton) *T.GtkWidget

	Gtk_scale_button_get_orientation func(
		button *T.GtkScaleButton) T.GtkOrientation

	Gtk_scale_button_set_orientation func(
		button *T.GtkScaleButton,
		orientation T.GtkOrientation)

	Gtk_vscrollbar_get_type func() T.GType

	Gtk_vscrollbar_new func(
		adjustment *T.GtkAdjustment) *T.GtkWidget

	Gtk_viewport_get_type func() T.GType

	Gtk_viewport_new func(
		hadjustment *T.GtkAdjustment,
		vadjustment *T.GtkAdjustment) *T.GtkWidget

	Gtk_viewport_get_hadjustment func(
		viewport *T.GtkViewport) *T.GtkAdjustment

	Gtk_viewport_get_vadjustment func(
		viewport *T.GtkViewport) *T.GtkAdjustment

	Gtk_viewport_set_hadjustment func(
		viewport *T.GtkViewport,
		adjustment *T.GtkAdjustment)

	Gtk_viewport_set_vadjustment func(
		viewport *T.GtkViewport,
		adjustment *T.GtkAdjustment)

	Gtk_viewport_set_shadow_type func(
		viewport *T.GtkViewport,
		t T.GtkShadowType)

	Gtk_viewport_get_shadow_type func(
		viewport *T.GtkViewport) T.GtkShadowType

	Gtk_viewport_get_bin_window func(
		viewport *T.GtkViewport) *T.GdkWindow

	Gtk_viewport_get_view_window func(
		viewport *T.GtkViewport) *T.GdkWindow

	Gtk_scrolled_window_get_type func() T.GType

	Gtk_scrolled_window_new func(
		hadjustment *T.GtkAdjustment,
		vadjustment *T.GtkAdjustment) *T.GtkWidget

	Gtk_scrolled_window_set_hadjustment func(
		scrolled_window *T.GtkScrolledWindow,
		hadjustment *T.GtkAdjustment)

	Gtk_scrolled_window_set_vadjustment func(
		scrolled_window *T.GtkScrolledWindow,
		vadjustment *T.GtkAdjustment)

	Gtk_scrolled_window_get_hadjustment func(
		scrolled_window *T.GtkScrolledWindow) *T.GtkAdjustment

	Gtk_scrolled_window_get_vadjustment func(
		scrolled_window *T.GtkScrolledWindow) *T.GtkAdjustment

	Gtk_scrolled_window_get_hscrollbar func(
		scrolled_window *T.GtkScrolledWindow) *T.GtkWidget

	Gtk_scrolled_window_get_vscrollbar func(
		scrolled_window *T.GtkScrolledWindow) *T.GtkWidget

	Gtk_scrolled_window_set_policy func(
		scrolled_window *T.GtkScrolledWindow,
		hscrollbar_policy T.GtkPolicyType,
		vscrollbar_policy T.GtkPolicyType)

	Gtk_scrolled_window_get_policy func(
		scrolled_window *T.GtkScrolledWindow,
		hscrollbar_policy *T.GtkPolicyType,
		vscrollbar_policy *T.GtkPolicyType)

	Gtk_scrolled_window_set_placement func(
		scrolled_window *T.GtkScrolledWindow,
		window_placement T.GtkCornerType)

	Gtk_scrolled_window_unset_placement func(
		scrolled_window *T.GtkScrolledWindow)

	Gtk_scrolled_window_get_placement func(
		scrolled_window *T.GtkScrolledWindow) T.GtkCornerType

	Gtk_scrolled_window_set_shadow_type func(
		scrolled_window *T.GtkScrolledWindow,
		t T.GtkShadowType)

	Gtk_scrolled_window_get_shadow_type func(
		scrolled_window *T.GtkScrolledWindow) T.GtkShadowType

	Gtk_scrolled_window_add_with_viewport func(
		scrolled_window *T.GtkScrolledWindow,
		child *T.GtkWidget)

	Gtk_separator_menu_item_get_type func() T.GType

	Gtk_separator_menu_item_new func() *T.GtkWidget

	Gtk_separator_tool_item_get_type func() T.GType

	Gtk_separator_tool_item_new func() *T.GtkToolItem

	Gtk_separator_tool_item_get_draw func(
		item *T.GtkSeparatorToolItem) T.Gboolean

	Gtk_separator_tool_item_set_draw func(
		item *T.GtkSeparatorToolItem,
		draw T.Gboolean)

	Gtk_show_uri func(
		screen *T.GdkScreen,
		uri string,
		timestamp T.Guint32,
		error **T.GError) T.Gboolean

	Gtk_spin_button_get_type func() T.GType

	Gtk_spin_button_configure func(
		spin_button *T.GtkSpinButton,
		adjustment *T.GtkAdjustment,
		climb_rate T.Gdouble,
		digits T.Guint)

	Gtk_spin_button_new func(
		adjustment *T.GtkAdjustment,
		climb_rate T.Gdouble,
		digits T.Guint) *T.GtkWidget

	Gtk_spin_button_new_with_range func(
		min T.Gdouble,
		max T.Gdouble,
		step T.Gdouble) *T.GtkWidget

	Gtk_spin_button_set_adjustment func(
		spin_button *T.GtkSpinButton,
		adjustment *T.GtkAdjustment)

	Gtk_spin_button_get_adjustment func(
		spin_button *T.GtkSpinButton) *T.GtkAdjustment

	Gtk_spin_button_set_digits func(
		spin_button *T.GtkSpinButton,
		digits T.Guint)

	Gtk_spin_button_get_digits func(
		spin_button *T.GtkSpinButton) T.Guint

	Gtk_spin_button_set_increments func(
		spin_button *T.GtkSpinButton,
		step T.Gdouble,
		page T.Gdouble)

	Gtk_spin_button_get_increments func(
		spin_button *T.GtkSpinButton,
		step *T.Gdouble,
		page *T.Gdouble)

	Gtk_spin_button_set_range func(
		spin_button *T.GtkSpinButton,
		min T.Gdouble,
		max T.Gdouble)

	Gtk_spin_button_get_range func(
		spin_button *T.GtkSpinButton,
		min *T.Gdouble,
		max *T.Gdouble)

	Gtk_spin_button_get_value func(
		spin_button *T.GtkSpinButton) T.Gdouble

	Gtk_spin_button_get_value_as_int func(
		spin_button *T.GtkSpinButton) T.Gint

	Gtk_spin_button_set_value func(
		spin_button *T.GtkSpinButton,
		value T.Gdouble)

	Gtk_spin_button_set_update_policy func(
		spin_button *T.GtkSpinButton,
		policy T.GtkSpinButtonUpdatePolicy)

	Gtk_spin_button_get_update_policy func(
		spin_button *T.GtkSpinButton) T.GtkSpinButtonUpdatePolicy

	Gtk_spin_button_set_numeric func(
		spin_button *T.GtkSpinButton,
		numeric T.Gboolean)

	Gtk_spin_button_get_numeric func(
		spin_button *T.GtkSpinButton) T.Gboolean

	Gtk_spin_button_spin func(
		spin_button *T.GtkSpinButton,
		direction T.GtkSpinType,
		increment T.Gdouble)

	Gtk_spin_button_set_wrap func(
		spin_button *T.GtkSpinButton,
		wrap T.Gboolean)

	Gtk_spin_button_get_wrap func(
		spin_button *T.GtkSpinButton) T.Gboolean

	Gtk_spin_button_set_snap_to_ticks func(
		spin_button *T.GtkSpinButton,
		snap_to_ticks T.Gboolean)

	Gtk_spin_button_get_snap_to_ticks func(
		spin_button *T.GtkSpinButton) T.Gboolean

	Gtk_spin_button_update func(
		spin_button *T.GtkSpinButton)

	Gtk_spinner_get_type func() T.GType

	Gtk_spinner_new func() *T.GtkWidget

	Gtk_spinner_start func(
		spinner *T.GtkSpinner)

	Gtk_spinner_stop func(
		spinner *T.GtkSpinner)

	Gtk_statusbar_get_type func() T.GType

	Gtk_statusbar_new func() *T.GtkWidget

	Gtk_statusbar_get_context_id func(
		statusbar *T.GtkStatusbar,
		context_description string) T.Guint

	Gtk_statusbar_push func(
		statusbar *T.GtkStatusbar,
		context_id T.Guint,
		text string) T.Guint

	Gtk_statusbar_pop func(
		statusbar *T.GtkStatusbar,
		context_id T.Guint)

	Gtk_statusbar_remove func(
		statusbar *T.GtkStatusbar,
		context_id T.Guint,
		message_id T.Guint)

	Gtk_statusbar_remove_all func(
		statusbar *T.GtkStatusbar,
		context_id T.Guint)

	Gtk_statusbar_set_has_resize_grip func(
		statusbar *T.GtkStatusbar,
		setting T.Gboolean)

	Gtk_statusbar_get_has_resize_grip func(
		statusbar *T.GtkStatusbar) T.Gboolean

	Gtk_statusbar_get_message_area func(
		statusbar *T.GtkStatusbar) *T.GtkWidget

	Gtk_status_icon_get_type func() T.GType

	Gtk_status_icon_new func() *T.GtkStatusIcon

	Gtk_status_icon_new_from_pixbuf func(
		pixbuf *T.GdkPixbuf) *T.GtkStatusIcon

	Gtk_status_icon_new_from_file func(
		filename string) *T.GtkStatusIcon

	Gtk_status_icon_new_from_stock func(
		stock_id string) *T.GtkStatusIcon

	Gtk_status_icon_new_from_icon_name func(
		icon_name string) *T.GtkStatusIcon

	Gtk_status_icon_new_from_gicon func(
		icon *T.GIcon) *T.GtkStatusIcon

	Gtk_status_icon_set_from_pixbuf func(
		status_icon *T.GtkStatusIcon,
		pixbuf *T.GdkPixbuf)

	Gtk_status_icon_set_from_file func(
		status_icon *T.GtkStatusIcon,
		filename string)

	Gtk_status_icon_set_from_stock func(
		status_icon *T.GtkStatusIcon,
		stock_id string)

	Gtk_status_icon_set_from_icon_name func(
		status_icon *T.GtkStatusIcon,
		icon_name string)

	Gtk_status_icon_set_from_gicon func(
		status_icon *T.GtkStatusIcon,
		icon *T.GIcon)

	Gtk_status_icon_get_storage_type func(
		status_icon *T.GtkStatusIcon) T.GtkImageType

	Gtk_status_icon_get_pixbuf func(
		status_icon *T.GtkStatusIcon) *T.GdkPixbuf

	Gtk_status_icon_get_stock func(
		status_icon *T.GtkStatusIcon) string

	Gtk_status_icon_get_icon_name func(
		status_icon *T.GtkStatusIcon) string

	Gtk_status_icon_get_gicon func(
		status_icon *T.GtkStatusIcon) *T.GIcon

	Gtk_status_icon_get_size func(
		status_icon *T.GtkStatusIcon) T.Gint

	Gtk_status_icon_set_screen func(
		status_icon *T.GtkStatusIcon,
		screen *T.GdkScreen)

	Gtk_status_icon_get_screen func(
		status_icon *T.GtkStatusIcon) *T.GdkScreen

	Gtk_status_icon_set_tooltip func(
		status_icon *T.GtkStatusIcon,
		tooltip_text string)

	Gtk_status_icon_set_has_tooltip func(
		status_icon *T.GtkStatusIcon,
		has_tooltip T.Gboolean)

	Gtk_status_icon_set_tooltip_text func(
		status_icon *T.GtkStatusIcon,
		text string)

	Gtk_status_icon_set_tooltip_markup func(
		status_icon *T.GtkStatusIcon,
		markup string)

	Gtk_status_icon_set_title func(
		status_icon *T.GtkStatusIcon,
		title string)

	Gtk_status_icon_get_title func(
		status_icon *T.GtkStatusIcon) string

	Gtk_status_icon_set_name func(
		status_icon *T.GtkStatusIcon,
		name string)

	Gtk_status_icon_set_visible func(
		status_icon *T.GtkStatusIcon,
		visible T.Gboolean)

	Gtk_status_icon_get_visible func(
		status_icon *T.GtkStatusIcon) T.Gboolean

	Gtk_status_icon_set_blinking func(
		status_icon *T.GtkStatusIcon,
		blinking T.Gboolean)

	Gtk_status_icon_get_blinking func(
		status_icon *T.GtkStatusIcon) T.Gboolean

	Gtk_status_icon_is_embedded func(
		status_icon *T.GtkStatusIcon) T.Gboolean

	Gtk_status_icon_position_menu func(
		menu *T.GtkMenu,
		x *T.Gint,
		y *T.Gint,
		push_in *T.Gboolean,
		user_data T.Gpointer)

	Gtk_status_icon_get_geometry func(
		status_icon *T.GtkStatusIcon,
		screen **T.GdkScreen,
		area *T.GdkRectangle,
		orientation *T.GtkOrientation) T.Gboolean

	Gtk_status_icon_get_has_tooltip func(
		status_icon *T.GtkStatusIcon) T.Gboolean

	Gtk_status_icon_get_tooltip_text func(
		status_icon *T.GtkStatusIcon) string

	Gtk_status_icon_get_tooltip_markup func(
		status_icon *T.GtkStatusIcon) string

	Gtk_status_icon_get_x11_window_id func(
		status_icon *T.GtkStatusIcon) T.Guint32

	Gtk_stock_add func(
		items *T.GtkStockItem,
		n_items T.Guint)

	Gtk_stock_add_static func(
		items *T.GtkStockItem,
		n_items T.Guint)

	Gtk_stock_lookup func(
		stock_id string,
		item *T.GtkStockItem) T.Gboolean

	Gtk_stock_list_ids func() *T.GSList

	Gtk_stock_item_copy func(
		item *T.GtkStockItem) *T.GtkStockItem

	Gtk_stock_item_free func(
		item *T.GtkStockItem)

	Gtk_stock_set_translate_func func(
		domain string,
		f T.GtkTranslateFunc,
		dataGpointer,
		notify T.GDestroyNotify)

	Gtk_table_get_type func() T.GType

	Gtk_table_new func(
		rows T.Guint,
		columns T.Guint,
		homogeneous T.Gboolean) *T.GtkWidget

	Gtk_table_resize func(
		table *T.GtkTable,
		rows T.Guint,
		columns T.Guint)

	Gtk_table_attach func(
		table *T.GtkTable,
		child *T.GtkWidget,
		left_attach T.Guint,
		right_attach T.Guint,
		top_attach T.Guint,
		bottom_attach T.Guint,
		xoptions T.GtkAttachOptions,
		yoptions T.GtkAttachOptions,
		xpadding T.Guint,
		ypadding T.Guint)

	Gtk_table_attach_defaults func(
		table *T.GtkTable,
		widget *T.GtkWidget,
		left_attach T.Guint,
		right_attach T.Guint,
		top_attach T.Guint,
		bottom_attach T.Guint)

	Gtk_table_set_row_spacing func(
		table *T.GtkTable,
		row T.Guint,
		spacing T.Guint)

	Gtk_table_get_row_spacing func(
		table *T.GtkTable,
		row T.Guint) T.Guint

	Gtk_table_set_col_spacing func(
		table *T.GtkTable,
		column T.Guint,
		spacing T.Guint)

	Gtk_table_get_col_spacing func(
		table *T.GtkTable,
		column T.Guint) T.Guint

	Gtk_table_set_row_spacings func(
		table *T.GtkTable,
		spacing T.Guint)

	Gtk_table_get_default_row_spacing func(
		table *T.GtkTable) T.Guint

	Gtk_table_set_col_spacings func(
		table *T.GtkTable,
		spacing T.Guint)

	Gtk_table_get_default_col_spacing func(
		table *T.GtkTable) T.Guint

	Gtk_table_set_homogeneous func(
		table *T.GtkTable,
		homogeneous T.Gboolean)

	Gtk_table_get_homogeneous func(
		table *T.GtkTable) T.Gboolean

	Gtk_table_get_size func(
		table *T.GtkTable,
		rows *T.Guint,
		columns *T.Guint)

	Gtk_text_tag_table_get_type func() T.GType

	Gtk_text_tag_table_new func() *T.GtkTextTagTable

	Gtk_text_tag_table_add func(
		table *T.GtkTextTagTable,
		tag *T.GtkTextTag)

	Gtk_text_tag_table_remove func(
		table *T.GtkTextTagTable,
		tag *T.GtkTextTag)

	Gtk_text_tag_table_lookup func(
		table *T.GtkTextTagTable,
		name string) *T.GtkTextTag

	Gtk_text_tag_table_foreach func(
		table *T.GtkTextTagTable,
		f T.GtkTextTagTableForeach,
		data T.Gpointer)

	Gtk_text_tag_table_get_size func(
		table *T.GtkTextTagTable) T.Gint

	Gtk_text_mark_get_type func() T.GType

	Gtk_text_mark_set_visible func(
		mark *T.GtkTextMark,
		setting T.Gboolean)

	Gtk_text_mark_get_visible func(
		mark *T.GtkTextMark) T.Gboolean

	Gtk_text_mark_new func(
		name string,
		left_gravity T.Gboolean) *T.GtkTextMark

	Gtk_text_mark_get_name func(
		mark *T.GtkTextMark) string

	Gtk_text_mark_get_deleted func(
		mark *T.GtkTextMark) T.Gboolean

	Gtk_text_mark_get_buffer func(
		mark *T.GtkTextMark) *T.GtkTextBuffer

	Gtk_text_mark_get_left_gravity func(
		mark *T.GtkTextMark) T.Gboolean

	Gtk_text_buffer_get_type func() T.GType

	Gtk_text_buffer_new func(
		table *T.GtkTextTagTable) *T.GtkTextBuffer

	Gtk_text_buffer_get_line_count func(
		buffer *T.GtkTextBuffer) T.Gint

	Gtk_text_buffer_get_char_count func(
		buffer *T.GtkTextBuffer) T.Gint

	Gtk_text_buffer_get_tag_table func(
		buffer *T.GtkTextBuffer) *T.GtkTextTagTable

	Gtk_text_buffer_set_text func(
		buffer *T.GtkTextBuffer,
		text string,
		len T.Gint)

	Gtk_text_buffer_insert func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		text string,
		len T.Gint)

	Gtk_text_buffer_insert_at_cursor func(
		buffer *T.GtkTextBuffer,
		text string,
		len T.Gint)

	Gtk_text_buffer_insert_interactive func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		text string,
		len T.Gint,
		default_editable T.Gboolean) T.Gboolean

	Gtk_text_buffer_insert_interactive_at_cursor func(
		buffer *T.GtkTextBuffer,
		text string,
		len T.Gint,
		default_editable T.Gboolean) T.Gboolean

	Gtk_text_buffer_insert_range func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	Gtk_text_buffer_insert_range_interactive func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		start *T.GtkTextIter,
		end *T.GtkTextIter,
		default_editable T.Gboolean) T.Gboolean

	Gtk_text_buffer_insert_with_tags func(buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter, text string, len T.Gint,
		first_tag *T.GtkTextTag, v ...VArg)

	Gtk_text_buffer_insert_with_tags_by_name func(
		buffer *T.GtkTextBuffer, iter *T.GtkTextIter, text string,
		len T.Gint, first_tag_name string, v ...VArg)

	Gtk_text_buffer_delete func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	Gtk_text_buffer_delete_interactive func(
		buffer *T.GtkTextBuffer,
		start_iter *T.GtkTextIter,
		end_iter *T.GtkTextIter,
		default_editable T.Gboolean) T.Gboolean

	Gtk_text_buffer_backspace func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		interactive T.Gboolean,
		default_editable T.Gboolean) T.Gboolean

	Gtk_text_buffer_get_text func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter,
		include_hidden_chars T.Gboolean) string

	Gtk_text_buffer_get_slice func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter,
		include_hidden_chars T.Gboolean) string

	Gtk_text_buffer_insert_pixbuf func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		pixbuf *T.GdkPixbuf)

	Gtk_text_buffer_insert_child_anchor func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		anchor *T.GtkTextChildAnchor)

	Gtk_text_buffer_create_child_anchor func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter) *T.GtkTextChildAnchor

	Gtk_text_buffer_add_mark func(
		buffer *T.GtkTextBuffer,
		mark *T.GtkTextMark,
		where *T.GtkTextIter)

	Gtk_text_buffer_create_mark func(
		buffer *T.GtkTextBuffer,
		mark_name string,
		where *T.GtkTextIter,
		left_gravity T.Gboolean) *T.GtkTextMark

	Gtk_text_buffer_move_mark func(
		buffer *T.GtkTextBuffer,
		mark *T.GtkTextMark,
		where *T.GtkTextIter)

	Gtk_text_buffer_delete_mark func(
		buffer *T.GtkTextBuffer,
		mark *T.GtkTextMark)

	Gtk_text_buffer_get_mark func(
		buffer *T.GtkTextBuffer,
		name string) *T.GtkTextMark

	Gtk_text_buffer_move_mark_by_name func(
		buffer *T.GtkTextBuffer,
		name string,
		where *T.GtkTextIter)

	Gtk_text_buffer_delete_mark_by_name func(
		buffer *T.GtkTextBuffer,
		name string)

	Gtk_text_buffer_get_insert func(
		buffer *T.GtkTextBuffer) *T.GtkTextMark

	Gtk_text_buffer_get_selection_bound func(
		buffer *T.GtkTextBuffer) *T.GtkTextMark

	Gtk_text_buffer_place_cursor func(
		buffer *T.GtkTextBuffer,
		where *T.GtkTextIter)

	Gtk_text_buffer_select_range func(
		buffer *T.GtkTextBuffer,
		ins *T.GtkTextIter,
		bound *T.GtkTextIter)

	Gtk_text_buffer_apply_tag func(
		buffer *T.GtkTextBuffer,
		tag *T.GtkTextTag,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	Gtk_text_buffer_remove_tag func(
		buffer *T.GtkTextBuffer,
		tag *T.GtkTextTag,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	Gtk_text_buffer_apply_tag_by_name func(
		buffer *T.GtkTextBuffer,
		name string,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	Gtk_text_buffer_remove_tag_by_name func(
		buffer *T.GtkTextBuffer,
		name string,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	Gtk_text_buffer_remove_all_tags func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	Gtk_text_buffer_create_tag func(
		buffer *T.GtkTextBuffer, tag_name string,
		first_property_name string, v ...VArg) *T.GtkTextTag

	Gtk_text_buffer_get_iter_at_line_offset func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		line_number T.Gint,
		char_offset T.Gint)

	Gtk_text_buffer_get_iter_at_line_index func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		line_number T.Gint,
		byte_index T.Gint)

	Gtk_text_buffer_get_iter_at_offset func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		char_offset T.Gint)

	Gtk_text_buffer_get_iter_at_line func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		line_number T.Gint)

	Gtk_text_buffer_get_start_iter func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter)

	Gtk_text_buffer_get_end_iter func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter)

	Gtk_text_buffer_get_bounds func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter)

	Gtk_text_buffer_get_iter_at_mark func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		mark *T.GtkTextMark)

	Gtk_text_buffer_get_iter_at_child_anchor func(
		buffer *T.GtkTextBuffer,
		iter *T.GtkTextIter,
		anchor *T.GtkTextChildAnchor)

	Gtk_text_buffer_get_modified func(
		buffer *T.GtkTextBuffer) T.Gboolean

	Gtk_text_buffer_set_modified func(
		buffer *T.GtkTextBuffer,
		setting T.Gboolean)

	Gtk_text_buffer_get_has_selection func(
		buffer *T.GtkTextBuffer) T.Gboolean

	Gtk_text_buffer_add_selection_clipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard)

	Gtk_text_buffer_remove_selection_clipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard)

	Gtk_text_buffer_cut_clipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard,
		default_editable T.Gboolean)

	Gtk_text_buffer_copy_clipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard)

	Gtk_text_buffer_paste_clipboard func(
		buffer *T.GtkTextBuffer,
		clipboard *T.GtkClipboard,
		override_location *T.GtkTextIter,
		default_editable T.Gboolean)

	Gtk_text_buffer_get_selection_bounds func(
		buffer *T.GtkTextBuffer,
		start *T.GtkTextIter,
		end *T.GtkTextIter) T.Gboolean

	Gtk_text_buffer_delete_selection func(
		buffer *T.GtkTextBuffer,
		interactive T.Gboolean,
		default_editable T.Gboolean) T.Gboolean

	Gtk_text_buffer_begin_user_action func(
		buffer *T.GtkTextBuffer)

	Gtk_text_buffer_end_user_action func(
		buffer *T.GtkTextBuffer)

	Gtk_text_buffer_get_copy_target_list func(
		buffer *T.GtkTextBuffer) *T.GtkTargetList

	Gtk_text_buffer_get_paste_target_list func(
		buffer *T.GtkTextBuffer) *T.GtkTargetList

	Gtk_text_buffer_register_serialize_format func(
		buffer *T.GtkTextBuffer,
		mime_type string,
		function T.GtkTextBufferSerializeFunc,
		user_dataGpointer,
		user_data_destroy T.GDestroyNotify) T.GdkAtom

	Gtk_text_buffer_register_serialize_tagset func(
		buffer *T.GtkTextBuffer,
		tagset_name string) T.GdkAtom

	Gtk_text_buffer_register_deserialize_format func(
		buffer *T.GtkTextBuffer,
		mime_type string,
		function T.GtkTextBufferDeserializeFunc,
		user_dataGpointer,
		user_data_destroy T.GDestroyNotify) T.GdkAtom

	Gtk_text_buffer_register_deserialize_tagset func(
		buffer *T.GtkTextBuffer,
		tagset_name string) T.GdkAtom

	Gtk_text_buffer_unregister_serialize_format func(
		buffer *T.GtkTextBuffer,
		format T.GdkAtom)

	Gtk_text_buffer_unregister_deserialize_format func(
		buffer *T.GtkTextBuffer,
		format T.GdkAtom)

	Gtk_text_buffer_deserialize_set_can_create_tags func(
		buffer *T.GtkTextBuffer,
		format T.GdkAtom,
		can_create_tags T.Gboolean)

	Gtk_text_buffer_deserialize_get_can_create_tags func(
		buffer *T.GtkTextBuffer,
		format T.GdkAtom) T.Gboolean

	Gtk_text_buffer_get_serialize_formats func(
		buffer *T.GtkTextBuffer,
		n_formats *T.Gint) *T.GdkAtom

	Gtk_text_buffer_get_deserialize_formats func(
		buffer *T.GtkTextBuffer,
		n_formats *T.Gint) *T.GdkAtom

	Gtk_text_buffer_serialize func(
		register_buffer *T.GtkTextBuffer,
		content_buffer *T.GtkTextBuffer,
		format T.GdkAtom,
		start *T.GtkTextIter,
		end *T.GtkTextIter,
		length *T.Gsize) *T.Guint8

	Gtk_text_buffer_deserialize func(
		register_buffer *T.GtkTextBuffer,
		content_buffer *T.GtkTextBuffer,
		format T.GdkAtom,
		iter *T.GtkTextIter,
		data *T.Guint8,
		length T.Gsize,
		error **T.GError) T.Gboolean

	Gtk_text_view_get_type func() T.GType

	Gtk_text_view_new func() *T.GtkWidget

	Gtk_text_view_new_with_buffer func(buffer *T.GtkTextBuffer) *T.GtkWidget

	Gtk_text_view_set_buffer func(text_view *T.GtkTextView,
		buffer *T.GtkTextBuffer)

	Gtk_text_view_get_buffer func(text_view *T.GtkTextView) *T.GtkTextBuffer

	Gtk_text_view_scroll_to_iter func(text_view *T.GtkTextView,
		iter *T.GtkTextIter,
		within_margin T.Gdouble,
		use_align T.Gboolean,
		xalign T.Gdouble,
		yalign T.Gdouble) T.Gboolean

	Gtk_text_view_scroll_to_mark func(text_view *T.GtkTextView,
		mark *T.GtkTextMark,
		within_margin T.Gdouble,
		use_align T.Gboolean,
		xalign T.Gdouble,
		yalign T.Gdouble)

	Gtk_text_view_scroll_mark_onscreen func(text_view *T.GtkTextView,
		mark *T.GtkTextMark)

	Gtk_text_view_move_mark_onscreen func(text_view *T.GtkTextView,
		mark *T.GtkTextMark) T.Gboolean

	Gtk_text_view_place_cursor_onscreen func(text_view *T.GtkTextView) T.Gboolean

	Gtk_text_view_get_visible_rect func(text_view *T.GtkTextView,
		visible_rect *T.GdkRectangle)

	Gtk_text_view_set_cursor_visible func(text_view *T.GtkTextView,
		setting T.Gboolean)

	Gtk_text_view_get_cursor_visible func(text_view *T.GtkTextView) T.Gboolean

	Gtk_text_view_get_iter_location func(text_view *T.GtkTextView,
		iter *T.GtkTextIter,
		location *T.GdkRectangle)

	Gtk_text_view_get_iter_at_location func(text_view *T.GtkTextView,
		iter *T.GtkTextIter,
		x T.Gint,
		y T.Gint)

	Gtk_text_view_get_iter_at_position func(text_view *T.GtkTextView,
		iter *T.GtkTextIter,
		trailing *T.Gint,
		x T.Gint,
		y T.Gint)

	Gtk_text_view_get_line_yrange func(text_view *T.GtkTextView,
		iter *T.GtkTextIter,
		y *T.Gint,
		height *T.Gint)

	Gtk_text_view_get_line_at_y func(text_view *T.GtkTextView,
		target_iter *T.GtkTextIter,
		y T.Gint,
		line_top *T.Gint)

	Gtk_text_view_buffer_to_window_coords func(text_view *T.GtkTextView,
		win T.GtkTextWindowType,
		buffer_x T.Gint,
		buffer_y T.Gint,
		window_x *T.Gint,
		window_y *T.Gint)

	Gtk_text_view_window_to_buffer_coords func(text_view *T.GtkTextView,
		win T.GtkTextWindowType,
		window_x T.Gint,
		window_y T.Gint,
		buffer_x *T.Gint,
		buffer_y *T.Gint)

	Gtk_text_view_get_hadjustment func(text_view *T.GtkTextView) *T.GtkAdjustment

	Gtk_text_view_get_vadjustment func(text_view *T.GtkTextView) *T.GtkAdjustment

	Gtk_text_view_get_window func(text_view *T.GtkTextView,
		win T.GtkTextWindowType) *T.GdkWindow

	Gtk_text_view_get_window_type func(text_view *T.GtkTextView,
		window *T.GdkWindow) T.GtkTextWindowType

	Gtk_text_view_set_border_window_size func(text_view *T.GtkTextView,
		t T.GtkTextWindowType,
		size T.Gint)

	Gtk_text_view_get_border_window_size func(text_view *T.GtkTextView,
		t T.GtkTextWindowType) T.Gint

	Gtk_text_view_forward_display_line func(text_view *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_view_backward_display_line func(text_view *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_view_forward_display_line_end func(text_view *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_view_backward_display_line_start func(text_view *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_view_starts_display_line func(text_view *T.GtkTextView,
		iter *T.GtkTextIter) T.Gboolean

	Gtk_text_view_move_visually func(text_view *T.GtkTextView,
		iter *T.GtkTextIter,
		count T.Gint) T.Gboolean

	Gtk_text_view_im_context_filter_keypress func(text_view *T.GtkTextView,
		event *T.GdkEventKey) T.Gboolean

	Gtk_text_view_reset_im_context func(text_view *T.GtkTextView)

	Gtk_text_view_add_child_at_anchor func(text_view *T.GtkTextView,
		child *T.GtkWidget,
		anchor *T.GtkTextChildAnchor)

	Gtk_text_view_add_child_in_window func(text_view *T.GtkTextView,
		child *T.GtkWidget,
		which_window T.GtkTextWindowType,
		xpos T.Gint,
		ypos T.Gint)

	Gtk_text_view_move_child func(text_view *T.GtkTextView,
		child *T.GtkWidget,
		xpos T.Gint,
		ypos T.Gint)

	Gtk_text_view_set_wrap_mode func(text_view *T.GtkTextView,
		wrap_mode T.GtkWrapMode)

	Gtk_text_view_get_wrap_mode func(text_view *T.GtkTextView) T.GtkWrapMode

	Gtk_text_view_set_editable func(text_view *T.GtkTextView,
		setting T.Gboolean)

	Gtk_text_view_get_editable func(text_view *T.GtkTextView) T.Gboolean

	Gtk_text_view_set_overwrite func(text_view *T.GtkTextView,
		overwrite T.Gboolean)

	Gtk_text_view_get_overwrite func(text_view *T.GtkTextView) T.Gboolean

	Gtk_text_view_set_accepts_tab func(text_view *T.GtkTextView,
		accepts_tab T.Gboolean)

	Gtk_text_view_get_accepts_tab func(text_view *T.GtkTextView) T.Gboolean

	Gtk_text_view_set_pixels_above_lines func(text_view *T.GtkTextView,
		pixels_above_lines T.Gint)

	Gtk_text_view_get_pixels_above_lines func(text_view *T.GtkTextView) T.Gint

	Gtk_text_view_set_pixels_below_lines func(text_view *T.GtkTextView,
		pixels_below_lines T.Gint)

	Gtk_text_view_get_pixels_below_lines func(text_view *T.GtkTextView) T.Gint

	Gtk_text_view_set_pixels_inside_wrap func(text_view *T.GtkTextView,
		pixels_inside_wrap T.Gint)

	Gtk_text_view_get_pixels_inside_wrap func(text_view *T.GtkTextView) T.Gint

	Gtk_text_view_set_justification func(text_view *T.GtkTextView,
		justification T.GtkJustification)

	Gtk_text_view_get_justification func(text_view *T.GtkTextView) T.GtkJustification

	Gtk_text_view_set_left_margin func(text_view *T.GtkTextView,
		left_margin T.Gint)

	Gtk_text_view_get_left_margin func(text_view *T.GtkTextView) T.Gint

	Gtk_text_view_set_right_margin func(text_view *T.GtkTextView,
		right_margin T.Gint)

	Gtk_text_view_get_right_margin func(text_view *T.GtkTextView) T.Gint

	Gtk_text_view_set_indent func(text_view *T.GtkTextView,
		indent T.Gint)

	Gtk_text_view_get_indent func(text_view *T.GtkTextView) T.Gint

	Gtk_text_view_set_tabs func(text_view *T.GtkTextView,
		tabs *T.PangoTabArray)

	Gtk_text_view_get_tabs func(text_view *T.GtkTextView) *T.PangoTabArray

	Gtk_text_view_get_default_attributes func(text_view *T.GtkTextView) *T.GtkTextAttributes

	Gtk_pixmap_get_type func() T.GType

	Gtk_pixmap_new func(pixmap *T.GdkPixmap,
		mask *T.GdkBitmap) *T.GtkWidget

	Gtk_pixmap_set func(pixmap *T.GtkPixmap,
		val *T.GdkPixmap,
		mask *T.GdkBitmap)

	Gtk_pixmap_get func(pixmap *T.GtkPixmap,
		val **T.GdkPixmap,
		mask **T.GdkBitmap)

	Gtk_pixmap_set_build_insensitive func(pixmap *T.GtkPixmap,
		build T.Gboolean)

	Gtk_toolbar_get_type func() T.GType

	Gtk_toolbar_new func() *T.GtkWidget

	Gtk_toolbar_insert func(toolbar *T.GtkToolbar,
		item *T.GtkToolItem,
		pos T.Gint)

	Gtk_toolbar_get_item_index func(toolbar *T.GtkToolbar,
		item *T.GtkToolItem) T.Gint

	Gtk_toolbar_get_n_items func(toolbar *T.GtkToolbar) T.Gint

	Gtk_toolbar_get_nth_item func(toolbar *T.GtkToolbar,
		n T.Gint) *T.GtkToolItem

	Gtk_toolbar_get_show_arrow func(toolbar *T.GtkToolbar) T.Gboolean

	Gtk_toolbar_set_show_arrow func(toolbar *T.GtkToolbar,
		show_arrow T.Gboolean)

	Gtk_toolbar_get_style func(toolbar *T.GtkToolbar) T.GtkToolbarStyle

	Gtk_toolbar_set_style func(toolbar *T.GtkToolbar,
		style T.GtkToolbarStyle)

	Gtk_toolbar_unset_style func(toolbar *T.GtkToolbar)

	Gtk_toolbar_get_icon_size func(toolbar *T.GtkToolbar) T.GtkIconSize

	Gtk_toolbar_set_icon_size func(toolbar *T.GtkToolbar,
		icon_size T.GtkIconSize)

	Gtk_toolbar_unset_icon_size func(toolbar *T.GtkToolbar)

	Gtk_toolbar_get_relief_style func(toolbar *T.GtkToolbar) T.GtkReliefStyle

	Gtk_toolbar_get_drop_index func(toolbar *T.GtkToolbar,
		x T.Gint,
		y T.Gint) T.Gint

	Gtk_toolbar_set_drop_highlight_item func(toolbar *T.GtkToolbar,
		tool_item *T.GtkToolItem,
		index_ T.Gint)

	Gtk_toolbar_get_orientation func(toolbar *T.GtkToolbar) T.GtkOrientation

	Gtk_toolbar_set_orientation func(toolbar *T.GtkToolbar,
		orientation T.GtkOrientation)

	Gtk_toolbar_get_tooltips func(toolbar *T.GtkToolbar) T.Gboolean

	Gtk_toolbar_set_tooltips func(toolbar *T.GtkToolbar,
		enable T.Gboolean)

	Gtk_toolbar_append_item func(toolbar *T.GtkToolbar,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *T.GtkWidget,
		callback T.GCallback,
		user_data T.Gpointer) *T.GtkWidget

	Gtk_toolbar_prepend_item func(toolbar *T.GtkToolbar,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *T.GtkWidget,
		callback T.GCallback,
		user_data T.Gpointer) *T.GtkWidget

	Gtk_toolbar_insert_item func(toolbar *T.GtkToolbar,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *T.GtkWidget,
		callback T.GCallback,
		user_dataGpointer,
		position T.Gint) *T.GtkWidget

	Gtk_toolbar_insert_stock func(toolbar *T.GtkToolbar,
		stock_id string,
		tooltip_text string,
		tooltip_private_text string,
		callback T.GCallback,
		user_dataGpointer,
		position T.Gint) *T.GtkWidget

	Gtk_toolbar_append_space func(toolbar *T.GtkToolbar)

	Gtk_toolbar_prepend_space func(toolbar *T.GtkToolbar)

	Gtk_toolbar_insert_space func(toolbar *T.GtkToolbar,
		position T.Gint)

	Gtk_toolbar_remove_space func(toolbar *T.GtkToolbar,
		position T.Gint)

	Gtk_toolbar_append_element func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *T.GtkWidget,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *T.GtkWidget,
		callback T.GCallback,
		user_data T.Gpointer) *T.GtkWidget

	Gtk_toolbar_prepend_element func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *T.GtkWidget,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *T.GtkWidget,
		callback T.GCallback,
		user_data T.Gpointer) *T.GtkWidget

	Gtk_toolbar_insert_element func(toolbar *T.GtkToolbar,
		t T.GtkToolbarChildType,
		widget *T.GtkWidget,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *T.GtkWidget,
		callback T.GCallback,
		user_dataGpointer,
		position T.Gint) *T.GtkWidget

	Gtk_toolbar_append_widget func(toolbar *T.GtkToolbar,
		widget *T.GtkWidget,
		tooltip_text string,
		tooltip_private_text string)

	Gtk_toolbar_prepend_widget func(toolbar *T.GtkToolbar,
		widget *T.GtkWidget,
		tooltip_text string,
		tooltip_private_text string)

	Gtk_toolbar_insert_widget func(toolbar *T.GtkToolbar,
		widget *T.GtkWidget,
		tooltip_text string,
		tooltip_private_text string,
		position T.Gint)

	Gtk_tool_item_group_get_type func() T.GType

	Gtk_tool_item_group_new func(label string) *T.GtkWidget

	Gtk_tool_item_group_set_label func(group *T.GtkToolItemGroup,
		label string)

	Gtk_tool_item_group_set_label_widget func(group *T.GtkToolItemGroup,
		label_widget *T.GtkWidget)

	Gtk_tool_item_group_set_collapsed func(group *T.GtkToolItemGroup,
		collapsed T.Gboolean)

	Gtk_tool_item_group_set_ellipsize func(group *T.GtkToolItemGroup,
		ellipsize T.PangoEllipsizeMode)

	Gtk_tool_item_group_set_header_relief func(group *T.GtkToolItemGroup,
		style T.GtkReliefStyle)

	Gtk_tool_item_group_get_label func(group *T.GtkToolItemGroup) string

	Gtk_tool_item_group_get_label_widget func(group *T.GtkToolItemGroup) *T.GtkWidget

	Gtk_tool_item_group_get_collapsed func(group *T.GtkToolItemGroup) T.Gboolean

	Gtk_tool_item_group_get_ellipsize func(group *T.GtkToolItemGroup) T.PangoEllipsizeMode

	Gtk_tool_item_group_get_header_relief func(group *T.GtkToolItemGroup) T.GtkReliefStyle

	Gtk_tool_item_group_insert func(group *T.GtkToolItemGroup,
		item *T.GtkToolItem,
		position T.Gint)

	Gtk_tool_item_group_set_item_position func(group *T.GtkToolItemGroup,
		item *T.GtkToolItem,
		position T.Gint)

	Gtk_tool_item_group_get_item_position func(group *T.GtkToolItemGroup,
		item *T.GtkToolItem) T.Gint

	Gtk_tool_item_group_get_n_items func(group *T.GtkToolItemGroup) T.Guint

	Gtk_tool_item_group_get_nth_item func(group *T.GtkToolItemGroup,
		index T.Guint) *T.GtkToolItem

	Gtk_tool_item_group_get_drop_item func(group *T.GtkToolItemGroup,
		x T.Gint,
		y T.Gint) *T.GtkToolItem

	Gtk_tool_palette_get_type func() T.GType

	Gtk_tool_palette_new func() *T.GtkWidget

	Gtk_tool_palette_set_group_position func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup,
		position T.Gint)

	Gtk_tool_palette_set_exclusive func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup,
		exclusive T.Gboolean)

	Gtk_tool_palette_set_expand func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup,
		expand T.Gboolean)

	Gtk_tool_palette_get_group_position func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup) T.Gint

	Gtk_tool_palette_get_exclusive func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup) T.Gboolean

	Gtk_tool_palette_get_expand func(palette *T.GtkToolPalette,
		group *T.GtkToolItemGroup) T.Gboolean

	Gtk_tool_palette_set_icon_size func(palette *T.GtkToolPalette,
		icon_size T.GtkIconSize)

	Gtk_tool_palette_unset_icon_size func(palette *T.GtkToolPalette)

	Gtk_tool_palette_set_style func(palette *T.GtkToolPalette,
		style T.GtkToolbarStyle)

	Gtk_tool_palette_unset_style func(palette *T.GtkToolPalette)

	Gtk_tool_palette_get_icon_size func(palette *T.GtkToolPalette) T.GtkIconSize

	Gtk_tool_palette_get_style func(palette *T.GtkToolPalette) T.GtkToolbarStyle

	Gtk_tool_palette_get_drop_item func(palette *T.GtkToolPalette,
		x T.Gint,
		y T.Gint) *T.GtkToolItem

	Gtk_tool_palette_get_drop_group func(palette *T.GtkToolPalette,
		x T.Gint,
		y T.Gint) *T.GtkToolItemGroup

	Gtk_tool_palette_get_drag_item func(palette *T.GtkToolPalette,
		selection *T.GtkSelectionData) *T.GtkWidget

	Gtk_tool_palette_set_drag_source func(palette *T.GtkToolPalette,
		targets T.GtkToolPaletteDragTargets)

	Gtk_tool_palette_add_drag_dest func(palette *T.GtkToolPalette,
		widget *T.GtkWidget,
		flags T.GtkDestDefaults,
		targets T.GtkToolPaletteDragTargets,
		actions T.GdkDragAction)

	Gtk_tool_palette_get_hadjustment func(palette *T.GtkToolPalette) *T.GtkAdjustment

	Gtk_tool_palette_get_vadjustment func(palette *T.GtkToolPalette) *T.GtkAdjustment

	Gtk_tool_palette_get_drag_target_item func() *T.GtkTargetEntry

	Gtk_tool_palette_get_drag_target_group func() *T.GtkTargetEntry

	Gtk_tool_shell_get_type func() T.GType

	Gtk_tool_shell_get_icon_size func(shell *T.GtkToolShell) T.GtkIconSize

	Gtk_tool_shell_get_orientation func(shell *T.GtkToolShell) T.GtkOrientation

	Gtk_tool_shell_get_style func(shell *T.GtkToolShell) T.GtkToolbarStyle

	Gtk_tool_shell_get_relief_style func(shell *T.GtkToolShell) T.GtkReliefStyle

	Gtk_tool_shell_rebuild_menu func(shell *T.GtkToolShell)

	Gtk_tool_shell_get_text_orientation func(shell *T.GtkToolShell) T.GtkOrientation

	Gtk_tool_shell_get_text_alignment func(shell *T.GtkToolShell) T.Gfloat

	Gtk_tool_shell_get_ellipsize_mode func(shell *T.GtkToolShell) T.PangoEllipsizeMode

	Gtk_tool_shell_get_text_size_group func(shell *T.GtkToolShell) *T.GtkSizeGroup

	Gtk_test_init func(argcp *int, argvp ***T.Char, v ...VArg)

	Gtk_test_register_all_types func()

	Gtk_test_list_all_types func(n_types *T.Guint) *T.GType

	Gtk_test_find_widget func(widget *T.GtkWidget,
		label_pattern string,
		widget_type T.GType) *T.GtkWidget

	Gtk_test_create_widget func(widget_type T.GType,
		first_property_name string, v ...VArg) *T.GtkWidget

	Gtk_test_create_simple_window func(window_title string,
		dialog_text string) *T.GtkWidget

	Gtk_test_display_button_window func(window_title string,
		dialog_text string, v ...VArg) *T.GtkWidget

	Gtk_test_slider_set_perc func(widget *T.GtkWidget,
		percentage T.Double)

	Gtk_test_slider_get_value func(widget *T.GtkWidget) T.Double

	Gtk_test_spin_button_click func(spinner *T.GtkSpinButton,
		button T.Guint,
		upwards T.Gboolean) T.Gboolean

	Gtk_test_widget_click func(widget *T.GtkWidget,
		button T.Guint,
		modifiers T.GdkModifierType) T.Gboolean

	Gtk_test_widget_send_key func(widget *T.GtkWidget,
		keyval T.Guint,
		modifiers T.GdkModifierType) T.Gboolean

	Gtk_test_text_set func(widget *T.GtkWidget,
		string string)

	Gtk_test_text_get func(widget *T.GtkWidget) string

	Gtk_test_find_sibling func(base_widget *T.GtkWidget,
		widget_type T.GType) *T.GtkWidget

	Gtk_test_find_label func(widget *T.GtkWidget,
		label_pattern string) *T.GtkWidget

	Gtk_tree_drag_source_get_type func() T.GType

	Gtk_tree_drag_source_row_draggable func(drag_source *T.GtkTreeDragSource,
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_drag_source_drag_data_delete func(drag_source *T.GtkTreeDragSource,
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_drag_source_drag_data_get func(drag_source *T.GtkTreeDragSource,
		path *T.GtkTreePath,
		selection_data *T.GtkSelectionData) T.Gboolean

	Gtk_tree_drag_dest_get_type func() T.GType

	Gtk_tree_drag_dest_drag_data_received func(drag_dest *T.GtkTreeDragDest,
		dest *T.GtkTreePath,
		selection_data *T.GtkSelectionData) T.Gboolean

	Gtk_tree_drag_dest_row_drop_possible func(drag_dest *T.GtkTreeDragDest,
		dest_path *T.GtkTreePath,
		selection_data *T.GtkSelectionData) T.Gboolean

	Gtk_tree_set_row_drag_data func(selection_data *T.GtkSelectionData,
		tree_model *T.GtkTreeModel,
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_get_row_drag_data func(selection_data *T.GtkSelectionData,
		tree_model **T.GtkTreeModel,
		path **T.GtkTreePath) T.Gboolean

	Gtk_tree_model_sort_get_type func() T.GType

	Gtk_tree_model_sort_new_with_model func(child_model *T.GtkTreeModel) *T.GtkTreeModel

	Gtk_tree_model_sort_get_model func(tree_model *T.GtkTreeModelSort) *T.GtkTreeModel

	Gtk_tree_model_sort_convert_child_path_to_path func(tree_model_sort *T.GtkTreeModelSort,
		child_path *T.GtkTreePath) *T.GtkTreePath

	Gtk_tree_model_sort_convert_child_iter_to_iter func(tree_model_sort *T.GtkTreeModelSort,
		sort_iter *T.GtkTreeIter,
		child_iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_model_sort_convert_path_to_child_path func(tree_model_sort *T.GtkTreeModelSort,
		sorted_path *T.GtkTreePath) *T.GtkTreePath

	Gtk_tree_model_sort_convert_iter_to_child_iter func(tree_model_sort *T.GtkTreeModelSort,
		child_iter *T.GtkTreeIter,
		sorted_iter *T.GtkTreeIter)

	Gtk_tree_model_sort_reset_default_sort_func func(tree_model_sort *T.GtkTreeModelSort)

	Gtk_tree_model_sort_clear_cache func(tree_model_sort *T.GtkTreeModelSort)

	Gtk_tree_model_sort_iter_is_valid func(tree_model_sort *T.GtkTreeModelSort,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_selection_get_type func() T.GType

	Gtk_tree_selection_set_mode func(selection *T.GtkTreeSelection,
		typ T.GtkSelectionMode)

	Gtk_tree_selection_get_mode func(selection *T.GtkTreeSelection) T.GtkSelectionMode

	Gtk_tree_selection_set_select_function func(selection *T.GtkTreeSelection,
		f T.GtkTreeSelectionFunc,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_tree_selection_get_user_data func(selection *T.GtkTreeSelection) T.Gpointer

	Gtk_tree_selection_get_tree_view func(selection *T.GtkTreeSelection) *T.GtkTreeView

	Gtk_tree_selection_get_select_function func(selection *T.GtkTreeSelection) T.GtkTreeSelectionFunc

	Gtk_tree_selection_get_selected func(selection *T.GtkTreeSelection,
		model **T.GtkTreeModel,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_selection_get_selected_rows func(selection *T.GtkTreeSelection,
		model **T.GtkTreeModel) *T.GList

	Gtk_tree_selection_count_selected_rows func(selection *T.GtkTreeSelection) T.Gint

	Gtk_tree_selection_selected_foreach func(selection *T.GtkTreeSelection,
		f T.GtkTreeSelectionForeachFunc,
		data T.Gpointer)

	Gtk_tree_selection_select_path func(selection *T.GtkTreeSelection,
		path *T.GtkTreePath)

	Gtk_tree_selection_unselect_path func(selection *T.GtkTreeSelection,
		path *T.GtkTreePath)

	Gtk_tree_selection_select_iter func(selection *T.GtkTreeSelection,
		iter *T.GtkTreeIter)

	Gtk_tree_selection_unselect_iter func(selection *T.GtkTreeSelection,
		iter *T.GtkTreeIter)

	Gtk_tree_selection_path_is_selected func(selection *T.GtkTreeSelection,
		path *T.GtkTreePath) T.Gboolean

	Gtk_tree_selection_iter_is_selected func(selection *T.GtkTreeSelection,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_selection_select_all func(selection *T.GtkTreeSelection)

	Gtk_tree_selection_unselect_all func(selection *T.GtkTreeSelection)

	Gtk_tree_selection_select_range func(selection *T.GtkTreeSelection,
		start_path *T.GtkTreePath,
		end_path *T.GtkTreePath)

	Gtk_tree_selection_unselect_range func(selection *T.GtkTreeSelection,
		start_path *T.GtkTreePath,
		end_path *T.GtkTreePath)

	Gtk_tree_store_get_type func() T.GType

	Gtk_tree_store_new func(
		n_columns T.Gint, v ...VArg) *T.GtkTreeStore

	Gtk_tree_store_newv func(n_columns T.Gint,
		types *T.GType) *T.GtkTreeStore

	Gtk_tree_store_set_column_types func(tree_store *T.GtkTreeStore,
		n_columns T.Gint,
		types *T.GType)

	Gtk_tree_store_set_value func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		column T.Gint,
		value *T.GValue)

	Gtk_tree_store_set func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter, v ...VArg)

	Gtk_tree_store_set_valuesv func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		columns *T.Gint,
		values *T.GValue,
		n_values T.Gint)

	Gtk_tree_store_set_valist func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		var_args T.Va_list)

	Gtk_tree_store_remove func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_store_insert func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		position T.Gint)

	Gtk_tree_store_insert_before func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		sibling *T.GtkTreeIter)

	Gtk_tree_store_insert_after func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		sibling *T.GtkTreeIter)

	Gtk_tree_store_insert_with_values func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		position T.Gint, v ...VArg)

	Gtk_tree_store_insert_with_valuesv func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter,
		position T.Gint,
		columns *T.Gint,
		values *T.GValue,
		n_values T.Gint)

	Gtk_tree_store_prepend func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter)

	Gtk_tree_store_append func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		parent *T.GtkTreeIter)

	Gtk_tree_store_is_ancestor func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		descendant *T.GtkTreeIter) T.Gboolean

	Gtk_tree_store_iter_depth func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter) T.Gint

	Gtk_tree_store_clear func(tree_store *T.GtkTreeStore)

	Gtk_tree_store_iter_is_valid func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter) T.Gboolean

	Gtk_tree_store_reorder func(tree_store *T.GtkTreeStore,
		parent *T.GtkTreeIter,
		new_order *T.Gint)

	Gtk_tree_store_swap func(tree_store *T.GtkTreeStore,
		a *T.GtkTreeIter,
		b *T.GtkTreeIter)

	Gtk_tree_store_move_before func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		position *T.GtkTreeIter)

	Gtk_tree_store_move_after func(tree_store *T.GtkTreeStore,
		iter *T.GtkTreeIter,
		position *T.GtkTreeIter)

	Gtk_ui_manager_get_type func() T.GType

	Gtk_ui_manager_new func() *T.GtkUIManager

	Gtk_ui_manager_set_add_tearoffs func(self *T.GtkUIManager,
		add_tearoffs T.Gboolean)

	Gtk_ui_manager_get_add_tearoffs func(self *T.GtkUIManager) T.Gboolean

	Gtk_ui_manager_insert_action_group func(self *T.GtkUIManager,
		action_group *T.GtkActionGroup,
		pos T.Gint)

	Gtk_ui_manager_remove_action_group func(self *T.GtkUIManager,
		action_group *T.GtkActionGroup)

	Gtk_ui_manager_get_action_groups func(self *T.GtkUIManager) *T.GList

	Gtk_ui_manager_get_accel_group func(self *T.GtkUIManager) *T.GtkAccelGroup

	Gtk_ui_manager_get_widget func(self *T.GtkUIManager,
		path string) *T.GtkWidget

	Gtk_ui_manager_get_toplevels func(self *T.GtkUIManager,
		types T.GtkUIManagerItemType) *T.GSList

	Gtk_ui_manager_get_action func(self *T.GtkUIManager,
		path string) *T.GtkAction

	Gtk_ui_manager_add_ui_from_string func(self *T.GtkUIManager,
		buffer string,
		length T.Gssize,
		err **T.GError) T.Guint

	Gtk_ui_manager_add_ui_from_file func(self *T.GtkUIManager,
		filename string,
		err **T.GError) T.Guint

	Gtk_ui_manager_add_ui func(self *T.GtkUIManager,
		merge_id T.Guint,
		path string,
		name string,
		action string,
		t T.GtkUIManagerItemType,
		top T.Gboolean)

	Gtk_ui_manager_remove_ui func(self *T.GtkUIManager,
		merge_id T.Guint)

	Gtk_ui_manager_get_ui func(self *T.GtkUIManager) string

	Gtk_ui_manager_ensure_update func(self *T.GtkUIManager)

	Gtk_ui_manager_new_merge_id func(self *T.GtkUIManager) T.Guint

	Gtk_vbutton_box_get_type func() T.GType

	Gtk_vbutton_box_new func() *T.GtkWidget

	Gtk_vbutton_box_get_spacing_default func() T.Gint

	Gtk_vbutton_box_set_spacing_default func(spacing T.Gint)

	Gtk_vbutton_box_get_layout_default func() T.GtkButtonBoxStyle

	Gtk_vbutton_box_set_layout_default func(layout T.GtkButtonBoxStyle)

	Gtk_volume_button_get_type func() T.GType

	Gtk_volume_button_new func() *T.GtkWidget

	Gtk_vpaned_get_type func() T.GType

	Gtk_vpaned_new func() *T.GtkWidget

	Gtk_vruler_get_type func() T.GType

	Gtk_vruler_new func() *T.GtkWidget

	Gtk_vscale_get_type func() T.GType

	Gtk_vscale_new func(adjustment *T.GtkAdjustment) *T.GtkWidget

	Gtk_vscale_new_with_range func(min T.Gdouble,
		max T.Gdouble,
		step T.Gdouble) *T.GtkWidget

	Gtk_vseparator_get_type func() T.GType

	Gtk_vseparator_new func() *T.GtkWidget

	Gtk_clist_get_type func() T.GType

	Gtk_clist_new func(columns T.Gint) *T.GtkWidget

	Gtk_clist_new_with_titles func(columns T.Gint,
		titles **T.Gchar) *T.GtkWidget

	Gtk_clist_set_hadjustment func(clist *T.GtkCList,
		adjustment *T.GtkAdjustment)

	Gtk_clist_set_vadjustment func(clist *T.GtkCList,
		adjustment *T.GtkAdjustment)

	Gtk_clist_get_hadjustment func(clist *T.GtkCList) *T.GtkAdjustment

	Gtk_clist_get_vadjustment func(clist *T.GtkCList) *T.GtkAdjustment

	Gtk_clist_set_shadow_type func(clist *T.GtkCList,
		t T.GtkShadowType)

	Gtk_clist_set_selection_mode func(clist *T.GtkCList,
		mode T.GtkSelectionMode)

	Gtk_clist_set_reorderable func(clist *T.GtkCList,
		reorderable T.Gboolean)

	Gtk_clist_set_use_drag_icons func(clist *T.GtkCList,
		use_icons T.Gboolean)

	Gtk_clist_set_button_actions func(clist *T.GtkCList,
		button T.Guint,
		button_actions T.Guint8)

	Gtk_clist_freeze func(clist *T.GtkCList)

	Gtk_clist_thaw func(clist *T.GtkCList)

	Gtk_clist_column_titles_show func(clist *T.GtkCList)

	Gtk_clist_column_titles_hide func(clist *T.GtkCList)

	Gtk_clist_column_title_active func(clist *T.GtkCList,
		column T.Gint)

	Gtk_clist_column_title_passive func(clist *T.GtkCList,
		column T.Gint)

	Gtk_clist_column_titles_active func(clist *T.GtkCList)

	Gtk_clist_column_titles_passive func(clist *T.GtkCList)

	Gtk_clist_set_column_title func(clist *T.GtkCList,
		column T.Gint,
		title string)

	Gtk_clist_get_column_title func(clist *T.GtkCList,
		column T.Gint) string

	Gtk_clist_set_column_widget func(clist *T.GtkCList,
		column T.Gint,
		widget *T.GtkWidget)

	Gtk_clist_get_column_widget func(clist *T.GtkCList,
		column T.Gint) *T.GtkWidget

	Gtk_clist_set_column_justification func(clist *T.GtkCList,
		column T.Gint,
		justification T.GtkJustification)

	Gtk_clist_set_column_visibility func(clist *T.GtkCList,
		column T.Gint,
		visible T.Gboolean)

	Gtk_clist_set_column_resizeable func(clist *T.GtkCList,
		column T.Gint,
		resizeable T.Gboolean)

	Gtk_clist_set_column_auto_resize func(clist *T.GtkCList,
		column T.Gint,
		auto_resize T.Gboolean)

	Gtk_clist_columns_autosize func(clist *T.GtkCList) T.Gint

	Gtk_clist_optimal_column_width func(clist *T.GtkCList,
		column T.Gint) T.Gint

	Gtk_clist_set_column_width func(clist *T.GtkCList,
		column T.Gint,
		width T.Gint)

	Gtk_clist_set_column_min_width func(clist *T.GtkCList,
		column T.Gint,
		min_width T.Gint)

	Gtk_clist_set_column_max_width func(clist *T.GtkCList,
		column T.Gint,
		max_width T.Gint)

	Gtk_clist_set_row_height func(clist *T.GtkCList,
		height T.Guint)

	Gtk_clist_moveto func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		row_align T.Gfloat,
		col_align T.Gfloat)

	Gtk_clist_row_is_visible func(clist *T.GtkCList,
		row T.Gint) T.GtkVisibility

	Gtk_clist_get_cell_type func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint) T.GtkCellType

	Gtk_clist_set_text func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		text string)

	Gtk_clist_get_text func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		text **T.Gchar) T.Gint

	Gtk_clist_set_pixmap func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	Gtk_clist_get_pixmap func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap) T.Gint

	Gtk_clist_set_pixtext func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		text string,
		spacing T.Guint8,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	Gtk_clist_get_pixtext func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		text **T.Gchar,
		spacing *T.Guint8,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap) T.Gint

	Gtk_clist_set_foreground func(clist *T.GtkCList,
		row T.Gint,
		color *T.GdkColor)

	Gtk_clist_set_background func(clist *T.GtkCList,
		row T.Gint,
		color *T.GdkColor)

	Gtk_clist_set_cell_style func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		style *T.GtkStyle)

	Gtk_clist_get_cell_style func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint) *T.GtkStyle

	Gtk_clist_set_row_style func(clist *T.GtkCList,
		row T.Gint,
		style *T.GtkStyle)

	Gtk_clist_get_row_style func(clist *T.GtkCList,
		row T.Gint) *T.GtkStyle

	Gtk_clist_set_shift func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint,
		vertical T.Gint,
		horizontal T.Gint)

	Gtk_clist_set_selectable func(clist *T.GtkCList,
		row T.Gint,
		selectable T.Gboolean)

	Gtk_clist_get_selectable func(clist *T.GtkCList,
		row T.Gint) T.Gboolean

	Gtk_clist_prepend func(clist *T.GtkCList,
		text **T.Gchar) T.Gint

	Gtk_clist_append func(clist *T.GtkCList,
		text **T.Gchar) T.Gint

	Gtk_clist_insert func(clist *T.GtkCList,
		row T.Gint,
		text **T.Gchar) T.Gint

	Gtk_clist_remove func(clist *T.GtkCList,
		row T.Gint)

	Gtk_clist_set_row_data func(clist *T.GtkCList,
		row T.Gint,
		data T.Gpointer)

	Gtk_clist_set_row_data_full func(clist *T.GtkCList,
		row T.Gint,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_clist_get_row_data func(clist *T.GtkCList,
		row T.Gint) T.Gpointer

	Gtk_clist_find_row_from_data func(clist *T.GtkCList,
		data T.Gpointer) T.Gint

	Gtk_clist_select_row func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint)

	Gtk_clist_unselect_row func(clist *T.GtkCList,
		row T.Gint,
		column T.Gint)

	Gtk_clist_undo_selection func(clist *T.GtkCList)

	Gtk_clist_clear func(clist *T.GtkCList)

	Gtk_clist_get_selection_info func(clist *T.GtkCList,
		x T.Gint,
		y T.Gint,
		row *T.Gint,
		column *T.Gint) T.Gint

	Gtk_clist_select_all func(clist *T.GtkCList)

	Gtk_clist_unselect_all func(clist *T.GtkCList)

	Gtk_clist_swap_rows func(clist *T.GtkCList,
		row1 T.Gint,
		row2 T.Gint)

	Gtk_clist_row_move func(clist *T.GtkCList,
		source_row T.Gint,
		dest_row T.Gint)

	Gtk_clist_set_compare_func func(clist *T.GtkCList,
		cmp_func T.GtkCListCompareFunc)

	Gtk_clist_set_sort_column func(clist *T.GtkCList,
		column T.Gint)

	Gtk_clist_set_sort_type func(clist *T.GtkCList,
		sort_type T.GtkSortType)

	Gtk_clist_sort func(clist *T.GtkCList)

	Gtk_clist_set_auto_sort func(clist *T.GtkCList,
		auto_sort T.Gboolean)

	Gtk_combo_get_type func() T.GType

	Gtk_combo_new func() *T.GtkWidget

	Gtk_combo_set_value_in_list func(combo *T.GtkCombo,
		val T.Gboolean,
		ok_if_empty T.Gboolean)

	Gtk_combo_set_use_arrows func(combo *T.GtkCombo,
		val T.Gboolean)

	Gtk_combo_set_use_arrows_always func(combo *T.GtkCombo,
		val T.Gboolean)

	Gtk_combo_set_case_sensitive func(combo *T.GtkCombo,
		val T.Gboolean)

	Gtk_combo_set_item_string func(combo *T.GtkCombo,
		item *T.GtkItem,
		item_value string)

	Gtk_combo_set_popdown_strings func(combo *T.GtkCombo,
		strings *T.GList)

	Gtk_combo_disable_activate func(combo *T.GtkCombo)

	Gtk_ctree_get_type func() T.GType

	Gtk_ctree_new_with_titles func(columns T.Gint,
		tree_column T.Gint,
		titles **T.Gchar) *T.GtkWidget

	Gtk_ctree_new func(columns T.Gint,
		tree_column T.Gint) *T.GtkWidget

	Gtk_ctree_insert_node func(ctree *T.GtkCTree,
		parent *T.GtkCTreeNode,
		sibling *T.GtkCTreeNode,
		text **T.Gchar,
		spacing T.Guint8,
		pixmap_closed *T.GdkPixmap,
		mask_closed *T.GdkBitmap,
		pixmap_opened *T.GdkPixmap,
		mask_opened *T.GdkBitmap,
		is_leaf T.Gboolean,
		expanded T.Gboolean) *T.GtkCTreeNode

	Gtk_ctree_remove_node func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_insert_gnode func(ctree *T.GtkCTree,
		parent *T.GtkCTreeNode,
		sibling *T.GtkCTreeNode,
		gnode *T.GNode,
		f T.GtkCTreeGNodeFunc,
		data T.Gpointer) *T.GtkCTreeNode

	Gtk_ctree_export_to_gnode func(ctree *T.GtkCTree,
		parent *T.GNode,
		sibling *T.GNode,
		node *T.GtkCTreeNode,
		f T.GtkCTreeGNodeFunc,
		data T.Gpointer) *T.GNode

	Gtk_ctree_post_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		f T.GtkCTreeFunc,
		data T.Gpointer)

	Gtk_ctree_post_recursive_to_depth func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		depth T.Gint,
		f T.GtkCTreeFunc,
		data T.Gpointer)

	Gtk_ctree_pre_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		f T.GtkCTreeFunc,
		data T.Gpointer)

	Gtk_ctree_pre_recursive_to_depth func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		depth T.Gint,
		f T.GtkCTreeFunc,
		data T.Gpointer)

	Gtk_ctree_is_viewable func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) T.Gboolean

	Gtk_ctree_last func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) *T.GtkCTreeNode

	Gtk_ctree_find_node_ptr func(ctree *T.GtkCTree,
		ctree_row *T.GtkCTreeRow) *T.GtkCTreeNode

	Gtk_ctree_node_nth func(ctree *T.GtkCTree,
		row T.Guint) *T.GtkCTreeNode

	Gtk_ctree_find func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		child *T.GtkCTreeNode) T.Gboolean

	Gtk_ctree_is_ancestor func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		child *T.GtkCTreeNode) T.Gboolean

	Gtk_ctree_find_by_row_data func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		data T.Gpointer) *T.GtkCTreeNode

	Gtk_ctree_find_all_by_row_data func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		data T.Gpointer) *T.GList

	Gtk_ctree_find_by_row_data_custom func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		dataGpointer,
		f T.GCompareFunc) *T.GtkCTreeNode

	Gtk_ctree_find_all_by_row_data_custom func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		dataGpointer,
		f T.GCompareFunc) *T.GList

	Gtk_ctree_is_hot_spot func(ctree *T.GtkCTree,
		x T.Gint,
		y T.Gint) T.Gboolean

	Gtk_ctree_move func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		new_parent *T.GtkCTreeNode,
		new_sibling *T.GtkCTreeNode)

	Gtk_ctree_expand func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_expand_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_expand_to_depth func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		depth T.Gint)

	Gtk_ctree_collapse func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_collapse_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_collapse_to_depth func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		depth T.Gint)

	Gtk_ctree_toggle_expansion func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_toggle_expansion_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_select func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_select_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_unselect func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_unselect_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_real_select_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		state T.Gint)

	Gtk_ctree_node_set_text func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		text string)

	Gtk_ctree_node_set_pixmap func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	Gtk_ctree_node_set_pixtext func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		text string,
		spacing T.Guint8,
		pixmap *T.GdkPixmap,
		mask *T.GdkBitmap)

	Gtk_ctree_set_node_info func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		text string,
		spacing T.Guint8,
		pixmap_closed *T.GdkPixmap,
		mask_closed *T.GdkBitmap,
		pixmap_opened *T.GdkPixmap,
		mask_opened *T.GdkBitmap,
		is_leaf T.Gboolean,
		expanded T.Gboolean)

	Gtk_ctree_node_set_shift func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		vertical T.Gint,
		horizontal T.Gint)

	Gtk_ctree_node_set_selectable func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		selectable T.Gboolean)

	Gtk_ctree_node_get_selectable func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) T.Gboolean

	Gtk_ctree_node_get_cell_type func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint) T.GtkCellType

	Gtk_ctree_node_get_text func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		text **T.Gchar) T.Gboolean

	Gtk_ctree_node_get_pixmap func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap) T.Gboolean

	Gtk_ctree_node_get_pixtext func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		text **T.Gchar,
		spacing *T.Guint8,
		pixmap **T.GdkPixmap,
		mask **T.GdkBitmap) T.Gboolean

	Gtk_ctree_get_node_info func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		text **T.Gchar,
		spacing *T.Guint8,
		pixmap_closed **T.GdkPixmap,
		mask_closed **T.GdkBitmap,
		pixmap_opened **T.GdkPixmap,
		mask_opened **T.GdkBitmap,
		is_leaf *T.Gboolean,
		expanded *T.Gboolean) T.Gboolean

	Gtk_ctree_node_set_row_style func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		style *T.GtkStyle)

	Gtk_ctree_node_get_row_style func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) *T.GtkStyle

	Gtk_ctree_node_set_cell_style func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		style *T.GtkStyle)

	Gtk_ctree_node_get_cell_style func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint) *T.GtkStyle

	Gtk_ctree_node_set_foreground func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		color *T.GdkColor)

	Gtk_ctree_node_set_background func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		color *T.GdkColor)

	Gtk_ctree_node_set_row_data func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		data T.Gpointer)

	Gtk_ctree_node_set_row_data_full func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		dataGpointer,
		destroy T.GDestroyNotify)

	Gtk_ctree_node_get_row_data func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) T.Gpointer

	Gtk_ctree_node_moveto func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode,
		column T.Gint,
		row_align T.Gfloat,
		col_align T.Gfloat)

	Gtk_ctree_node_is_visible func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode) T.GtkVisibility

	Gtk_ctree_set_indent func(ctree *T.GtkCTree,
		indent T.Gint)

	Gtk_ctree_set_spacing func(ctree *T.GtkCTree,
		spacing T.Gint)

	Gtk_ctree_set_show_stub func(ctree *T.GtkCTree,
		show_stub T.Gboolean)

	Gtk_ctree_set_line_style func(ctree *T.GtkCTree,
		line_style T.GtkCTreeLineStyle)

	Gtk_ctree_set_expander_style func(ctree *T.GtkCTree,
		expander_style T.GtkCTreeExpanderStyle)

	Gtk_ctree_set_drag_compare_func func(ctree *T.GtkCTree,
		cmp_func T.GtkCTreeCompareDragFunc)

	Gtk_ctree_sort_node func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_sort_recursive func(ctree *T.GtkCTree,
		node *T.GtkCTreeNode)

	Gtk_ctree_node_get_type func() T.GType

	Gtk_curve_get_type func() T.GType

	Gtk_curve_new func() *T.GtkWidget

	Gtk_curve_reset func(curve *T.GtkCurve)

	Gtk_curve_set_gamma func(curve *T.GtkCurve,
		gamma_ T.Gfloat)

	Gtk_curve_set_range func(curve *T.GtkCurve,
		min_x T.Gfloat,
		max_x T.Gfloat,
		min_y T.Gfloat,
		max_y T.Gfloat)

	Gtk_curve_get_vector func(curve *T.GtkCurve,
		veclen int,
		vector *T.Gfloat)

	Gtk_curve_set_vector func(curve *T.GtkCurve,
		veclen int,
		vector *T.Gfloat)

	Gtk_curve_set_curve_type func(curve *T.GtkCurve,
		t T.GtkCurveType)

	Gtk_file_selection_get_type func() T.GType

	Gtk_file_selection_new func(title string) *T.GtkWidget

	Gtk_file_selection_set_filename func(filesel *T.GtkFileSelection,
		filename string)

	Gtk_file_selection_get_filename func(filesel *T.GtkFileSelection) string

	Gtk_file_selection_complete func(filesel *T.GtkFileSelection,
		pattern string)

	Gtk_file_selection_show_fileop_buttons func(filesel *T.GtkFileSelection)

	Gtk_file_selection_hide_fileop_buttons func(filesel *T.GtkFileSelection)

	Gtk_file_selection_get_selections func(filesel *T.GtkFileSelection) **T.Gchar

	Gtk_file_selection_set_select_multiple func(filesel *T.GtkFileSelection,
		select_multiple T.Gboolean)

	Gtk_file_selection_get_select_multiple func(filesel *T.GtkFileSelection) T.Gboolean

	Gtk_gamma_curve_get_type func() T.GType

	Gtk_gamma_curve_new func() *T.GtkWidget

	Gtk_input_dialog_get_type func() T.GType

	Gtk_input_dialog_new func() *T.GtkWidget

	Gtk_item_factory_get_type func() T.GType

	Gtk_item_factory_new func(container_type T.GType,
		path string,
		accel_group *T.GtkAccelGroup) *T.GtkItemFactory

	Gtk_item_factory_construct func(ifactory *T.GtkItemFactory,
		container_type T.GType,
		path string,
		accel_group *T.GtkAccelGroup)

	Gtk_item_factory_add_foreign func(accel_widget *T.GtkWidget,
		full_path string,
		accel_group *T.GtkAccelGroup,
		keyval T.Guint,
		modifiers T.GdkModifierType)

	Gtk_item_factory_from_widget func(widget *T.GtkWidget) *T.GtkItemFactory

	Gtk_item_factory_path_from_widget func(widget *T.GtkWidget) string

	Gtk_item_factory_get_item func(ifactory *T.GtkItemFactory,
		path string) *T.GtkWidget

	Gtk_item_factory_get_widget func(ifactory *T.GtkItemFactory,
		path string) *T.GtkWidget

	Gtk_item_factory_get_widget_by_action func(ifactory *T.GtkItemFactory,
		action T.Guint) *T.GtkWidget

	Gtk_item_factory_get_item_by_action func(ifactory *T.GtkItemFactory,
		action T.Guint) *T.GtkWidget

	Gtk_item_factory_create_item func(ifactory *T.GtkItemFactory,
		entry *T.GtkItemFactoryEntry,
		callback_dataGpointer,
		callback_type T.Guint)

	Gtk_item_factory_create_items func(ifactory *T.GtkItemFactory,
		n_entries T.Guint,
		entries *T.GtkItemFactoryEntry,
		callback_data T.Gpointer)

	Gtk_item_factory_delete_item func(ifactory *T.GtkItemFactory,
		path string)

	Gtk_item_factory_delete_entry func(ifactory *T.GtkItemFactory,
		entry *T.GtkItemFactoryEntry)

	Gtk_item_factory_delete_entries func(ifactory *T.GtkItemFactory,
		n_entries T.Guint,
		entries *T.GtkItemFactoryEntry)

	Gtk_item_factory_popup func(ifactory *T.GtkItemFactory,
		x T.Guint,
		y T.Guint,
		mouse_button T.Guint,
		time_ T.Guint32)

	Gtk_item_factory_popup_with_data func(ifactory *T.GtkItemFactory,
		popup_dataGpointer,
		destroy T.GDestroyNotify,
		x T.Guint,
		y T.Guint,
		mouse_button T.Guint,
		time_ T.Guint32)

	Gtk_item_factory_popup_data func(ifactory *T.GtkItemFactory) T.Gpointer

	Gtk_item_factory_popup_data_from_widget func(widget *T.GtkWidget) T.Gpointer

	Gtk_item_factory_set_translate_func func(ifactory *T.GtkItemFactory,
		f T.GtkTranslateFunc,
		dataGpointer,
		notify T.GDestroyNotify)

	Gtk_item_factory_create_items_ac func(ifactory *T.GtkItemFactory,
		n_entries T.Guint,
		entries *T.GtkItemFactoryEntry,
		callback_dataGpointer,
		callback_type T.Guint)

	Gtk_item_factory_from_path func(path string) *T.GtkItemFactory

	Gtk_item_factory_create_menu_entries func(n_entries T.Guint,
		entries *T.GtkMenuEntry)

	Gtk_item_factories_path_delete func(ifactory_path string,
		path string)

	Gtk_list_get_type func() T.GType

	Gtk_list_new func() *T.GtkWidget

	Gtk_list_insert_items func(list *T.GtkList,
		items *T.GList,
		position T.Gint)

	Gtk_list_append_items func(list *T.GtkList,
		items *T.GList)

	Gtk_list_prepend_items func(list *T.GtkList,
		items *T.GList)

	Gtk_list_remove_items func(list *T.GtkList,
		items *T.GList)

	Gtk_list_remove_items_no_unref func(list *T.GtkList,
		items *T.GList)

	Gtk_list_clear_items func(list *T.GtkList,
		start T.Gint,
		end T.Gint)

	Gtk_list_select_item func(list *T.GtkList,
		item T.Gint)

	Gtk_list_unselect_item func(list *T.GtkList,
		item T.Gint)

	Gtk_list_select_child func(list *T.GtkList,
		child *T.GtkWidget)

	Gtk_list_unselect_child func(list *T.GtkList,
		child *T.GtkWidget)

	Gtk_list_child_position func(list *T.GtkList,
		child *T.GtkWidget) T.Gint

	Gtk_list_set_selection_mode func(list *T.GtkList,
		mode T.GtkSelectionMode)

	Gtk_list_extend_selection func(list *T.GtkList,
		scroll_type T.GtkScrollType,
		position T.Gfloat,
		auto_start_selection T.Gboolean)

	Gtk_list_start_selection func(list *T.GtkList)

	Gtk_list_end_selection func(list *T.GtkList)

	Gtk_list_select_all func(list *T.GtkList)

	Gtk_list_unselect_all func(list *T.GtkList)

	Gtk_list_scroll_horizontal func(list *T.GtkList,
		scroll_type T.GtkScrollType,
		position T.Gfloat)

	Gtk_list_scroll_vertical func(list *T.GtkList,
		scroll_type T.GtkScrollType,
		position T.Gfloat)

	Gtk_list_toggle_add_mode func(list *T.GtkList)

	Gtk_list_toggle_focus_row func(list *T.GtkList)

	Gtk_list_toggle_row func(list *T.GtkList,
		item *T.GtkWidget)

	Gtk_list_undo_selection func(list *T.GtkList)

	Gtk_list_end_drag_selection func(list *T.GtkList)

	Gtk_list_item_get_type func() T.GType

	Gtk_list_item_new func() *T.GtkWidget

	Gtk_list_item_new_with_label func(label string) *T.GtkWidget

	Gtk_list_item_select func(list_item *T.GtkListItem)

	Gtk_list_item_deselect func(list_item *T.GtkListItem)

	Gtk_old_editable_get_type func() T.GType

	Gtk_old_editable_claim_selection func(old_editable *T.GtkOldEditable,
		claim T.Gboolean,
		time_ T.Guint32)

	Gtk_old_editable_changed func(old_editable *T.GtkOldEditable)

	Gtk_option_menu_get_type func() T.GType

	Gtk_option_menu_new func() *T.GtkWidget

	Gtk_option_menu_get_menu func(option_menu *T.GtkOptionMenu) *T.GtkWidget

	Gtk_option_menu_set_menu func(option_menu *T.GtkOptionMenu,
		menu *T.GtkWidget)

	Gtk_option_menu_remove_menu func(option_menu *T.GtkOptionMenu)

	Gtk_option_menu_get_history func(option_menu *T.GtkOptionMenu) T.Gint

	Gtk_option_menu_set_history func(option_menu *T.GtkOptionMenu,
		index_ T.Guint)

	Gtk_preview_get_type func() T.GType

	Gtk_preview_uninit func()

	Gtk_preview_new func(t T.GtkPreviewType) *T.GtkWidget

	Gtk_preview_size func(preview *T.GtkPreview,
		width T.Gint,
		height T.Gint)

	Gtk_preview_put func(preview *T.GtkPreview,
		window *T.GdkWindow,
		gc *T.GdkGC,
		srcx T.Gint,
		srcy T.Gint,
		destx T.Gint,
		desty T.Gint,
		width T.Gint,
		height T.Gint)

	Gtk_preview_draw_row func(preview *T.GtkPreview,
		data *T.Guchar,
		x T.Gint,
		y T.Gint,
		w T.Gint)

	Gtk_preview_set_expand func(preview *T.GtkPreview,
		expand T.Gboolean)

	Gtk_preview_set_gamma func(gamma_ T.Double)

	Gtk_preview_set_color_cube func(nred_shades T.Guint,
		ngreen_shades T.Guint,
		nblue_shades T.Guint,
		ngray_shades T.Guint)

	Gtk_preview_set_install_cmap func(install_cmap T.Gint)

	Gtk_preview_set_reserved func(nreserved T.Gint)

	Gtk_preview_set_dither func(preview *T.GtkPreview,
		dither T.GdkRgbDither)

	Gtk_preview_get_visual func() *T.GdkVisual

	Gtk_preview_get_cmap func() *T.GdkColormap

	Gtk_preview_get_info func() *T.GtkPreviewInfo

	Gtk_preview_reset func()

	Gtk_tips_query_get_type func() T.GType

	Gtk_tips_query_new func() *T.GtkWidget

	Gtk_tips_query_start_query func(tips_query *T.GtkTipsQuery)

	Gtk_tips_query_stop_query func(tips_query *T.GtkTipsQuery)

	Gtk_tips_query_set_caller func(tips_query *T.GtkTipsQuery,
		caller *T.GtkWidget)

	Gtk_tips_query_set_labels func(tips_query *T.GtkTipsQuery,
		label_inactive string,
		label_no_tip string)

	Gtk_marshal_BOOLEAN__VOID func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_BOOLEAN__POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_BOOLEAN__POINTER_POINTER_INT_INT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_BOOLEAN__POINTER_INT_INT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_BOOLEAN__POINTER_INT_INT_UINT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_BOOLEAN__POINTER_STRING_STRING_POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_ENUM__ENUM func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_INT__POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_INT__POINTER_CHAR_CHAR func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__ENUM_FLOAT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__ENUM_FLOAT_BOOLEAN func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__INT_INT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__INT_INT_POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_INT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_POINTER_POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_STRING_STRING func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_UINT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_UINT_ENUM func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_POINTER_UINT_UINT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_INT_INT_POINTER_UINT_UINT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__POINTER_UINT_UINT func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__STRING_INT_POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__UINT_POINTER_UINT_ENUM_ENUM_POINTER func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__UINT_POINTER_UINT_UINT_ENUM func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_marshal_VOID__UINT_STRING func(
		closure *T.GClosure,
		return_value *T.GValue,
		n_param_values T.Guint,
		param_values *T.GValue,
		invocation_hint T.Gpointer,
		marshal_data T.Gpointer)

	Gtk_builder_error_quark func() T.GQuark

	Gtk_print_context_set_cairo_context func(
		context *T.GtkPrintContext,
		cr *T.Cairo_t,
		dpi_x, dpi_y T.Double)

	Gtk_tearoff_menu_item_get_type func() T.GType

	Gtk_tearoff_menu_item_new func() *T.GtkWidget

	Gtk_tree_get_type func() T.GType

	Gtk_tree_new func() *T.GtkWidget

	Gtk_tree_append func(tree *T.GtkTree, tree_item *T.GtkWidget)

	Gtk_tree_prepend func(tree *T.GtkTree, tree_item *T.GtkWidget)

	Gtk_tree_insert func(
		tree *T.GtkTree, tree_item *T.GtkWidget, position T.Gint)

	Gtk_tree_remove_items func(tree *T.GtkTree, items *T.GList)

	Gtk_tree_clear_items func(
		tree *T.GtkTree, start T.Gint, end T.Gint)

	Gtk_tree_select_item func(tree *T.GtkTree, item T.Gint)

	Gtk_tree_unselect_item func(tree *T.GtkTree, item T.Gint)

	Gtk_tree_select_child func(
		tree *T.GtkTree, tree_item *T.GtkWidget)

	Gtk_tree_unselect_child func(
		tree *T.GtkTree, tree_item *T.GtkWidget)

	Gtk_tree_child_position func(
		tree *T.GtkTree, child *T.GtkWidget) T.Gint

	Gtk_tree_set_selection_mode func(
		tree *T.GtkTree, mode T.GtkSelectionMode)

	Gtk_tree_set_view_mode func(
		tree *T.GtkTree, mode T.GtkTreeViewMode)

	Gtk_tree_set_view_lines func(tree *T.GtkTree, flag T.Gboolean)

	Gtk_tree_remove_item func(tree *T.GtkTree, child *T.GtkWidget)

	Gtk_tree_item_get_type func() T.GType

	Gtk_tree_item_new func() *T.GtkWidget

	Gtk_tree_item_new_with_label func(label string) *T.GtkWidget

	Gtk_tree_item_set_subtree func(
		tree_item *T.GtkTreeItem, subtree *T.GtkWidget)

	Gtk_tree_item_remove_subtree func(tree_item *T.GtkTreeItem)

	Gtk_tree_item_select func(tree_item *T.GtkTreeItem)

	Gtk_tree_item_deselect func(tree_item *T.GtkTreeItem)

	Gtk_tree_item_expand func(tree_item *T.GtkTreeItem)

	Gtk_tree_item_collapse func(tree_item *T.GtkTreeItem)

	Gtk_type_enum_get_values func(
		enum_type T.GtkType) *T.GtkEnumValue

	Gtk_type_flags_get_values func(
		flags_type T.GtkType) *T.GtkFlagValue

	Gtk_type_enum_find_value func(
		enum_type T.GtkType,
		value_name string) *T.GtkEnumValue

	Gtk_type_flags_find_value func(
		flags_type T.GtkType,
		value_name string) *T.GtkFlagValue
)
