// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package gtk provides API definitions for accessing
//libgtk-win32-2.0-0.dll.
package gtk

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

var (
	Gtk_accel_group_get_type func() GType

	Gtk_accel_group_new func() *GtkAccelGroup

	Gtk_accel_group_get_is_locked func(
		accel_group *GtkAccelGroup) Gboolean

	Gtk_accel_group_get_modifier_mask func(
		accel_group *GtkAccelGroup) GdkModifierType

	Gtk_accel_group_lock func(accel_group *GtkAccelGroup)

	Gtk_accel_group_unlock func(accel_group *GtkAccelGroup)

	Gtk_accel_group_connect func(
		accel_group *GtkAccelGroup,
		accel_key Guint,
		accel_mods GdkModifierType,
		accel_flags GtkAccelFlags,
		closure *GClosure)

	Gtk_accel_group_connect_by_path func(
		accel_group *GtkAccelGroup,
		accel_path string,
		closure *GClosure)

	Gtk_accel_group_disconnect func(
		accel_group *GtkAccelGroup,
		closure *GClosure) Gboolean

	Gtk_accel_group_disconnect_key func(
		accel_group *GtkAccelGroup,
		accel_key Guint,
		accel_mods GdkModifierType) Gboolean

	Gtk_accel_group_activate func(
		accel_group *GtkAccelGroup,
		accel_quark GQuark,
		acceleratable *GObject,
		accel_key Guint,
		accel_mods GdkModifierType) Gboolean

	Gtk_accel_groups_activate func(
		object *GObject,
		accel_key Guint,
		accel_mods GdkModifierType) Gboolean

	Gtk_accel_groups_from_object func(
		object *GObject) *GSList

	Gtk_accel_group_find func(
		accel_group *GtkAccelGroup,
		find_func GtkAccelGroupFindFunc,
		data Gpointer) *GtkAccelKey

	Gtk_accel_group_from_accel_closure func(
		closure *GClosure) *GtkAccelGroup

	Gtk_accelerator_valid func(
		keyval Guint,
		modifiers GdkModifierType) Gboolean

	Gtk_accelerator_parse func(
		accelerator string,
		accelerator_key *Guint,
		accelerator_mods *GdkModifierType)

	Gtk_accelerator_name func(
		accelerator_key Guint,
		accelerator_mods GdkModifierType) string

	Gtk_accelerator_get_label func(
		accelerator_key Guint,
		accelerator_mods GdkModifierType) string

	Gtk_accelerator_set_default_mod_mask func(
		default_mod_mask GdkModifierType)

	Gtk_accelerator_get_default_mod_mask func() Guint

	Gtk_accel_group_query func(
		accel_group *GtkAccelGroup,
		accel_key Guint,
		accel_mods GdkModifierType,
		n_entries *Guint) *GtkAccelGroupEntry

	Gtk_accel_flags_get_type func() GType

	Gtk_assistant_page_type_get_type func() GType

	Gtk_builder_error_get_type func() GType

	Gtk_calendar_display_options_get_type func() GType

	Gtk_cell_renderer_state_get_type func() GType

	Gtk_cell_renderer_mode_get_type func() GType

	Gtk_cell_renderer_accel_mode_get_type func() GType

	Gtk_debug_flag_get_type func() GType

	Gtk_dialog_flags_get_type func() GType

	Gtk_response_type_get_type func() GType

	Gtk_dest_defaults_get_type func() GType

	Gtk_target_flags_get_type func() GType

	Gtk_entry_icon_position_get_type func() GType

	Gtk_anchor_type_get_type func() GType

	Gtk_arrow_placement_get_type func() GType

	Gtk_arrow_type_get_type func() GType

	Gtk_attach_options_get_type func() GType

	Gtk_button_box_style_get_type func() GType

	Gtk_curve_type_get_type func() GType

	Gtk_delete_type_get_type func() GType

	Gtk_direction_type_get_type func() GType

	Gtk_expander_style_get_type func() GType

	Gtk_icon_size_get_type func() GType

	Gtk_sensitivity_type_get_type func() GType

	Gtk_side_type_get_type func() GType

	Gtk_text_direction_get_type func() GType

	Gtk_justification_get_type func() GType

	Gtk_match_type_get_type func() GType

	Gtk_menu_direction_type_get_type func() GType

	Gtk_message_type_get_type func() GType

	Gtk_metric_type_get_type func() GType

	Gtk_movement_step_get_type func() GType

	Gtk_scroll_step_get_type func() GType

	Gtk_orientation_get_type func() GType

	Gtk_corner_type_get_type func() GType

	Gtk_pack_type_get_type func() GType

	Gtk_path_priority_type_get_type func() GType

	Gtk_path_type_get_type func() GType

	Gtk_policy_type_get_type func() GType

	Gtk_position_type_get_type func() GType

	Gtk_preview_type_get_type func() GType

	Gtk_relief_style_get_type func() GType

	Gtk_resize_mode_get_type func() GType

	Gtk_signal_run_type_get_type func() GType

	Gtk_scroll_type_get_type func() GType

	Gtk_selection_mode_get_type func() GType

	Gtk_shadow_type_get_type func() GType

	Gtk_state_type_get_type func() GType

	Gtk_submenu_direction_get_type func() GType

	Gtk_submenu_placement_get_type func() GType

	Gtk_toolbar_style_get_type func() GType

	Gtk_update_type_get_type func() GType

	Gtk_visibility_get_type func() GType

	Gtk_window_position_get_type func() GType

	Gtk_window_type_get_type func() GType

	Gtk_wrap_mode_get_type func() GType

	Gtk_sort_type_get_type func() GType

	Gtk_im_preedit_style_get_type func() GType

	Gtk_im_status_style_get_type func() GType

	Gtk_pack_direction_get_type func() GType

	Gtk_print_pages_get_type func() GType

	Gtk_page_set_get_type func() GType

	Gtk_number_up_layout_get_type func() GType

	Gtk_page_orientation_get_type func() GType

	Gtk_print_quality_get_type func() GType

	Gtk_print_duplex_get_type func() GType

	Gtk_unit_get_type func() GType

	Gtk_tree_view_grid_lines_get_type func() GType

	Gtk_drag_result_get_type func() GType

	Gtk_file_chooser_action_get_type func() GType

	Gtk_file_chooser_confirmation_get_type func() GType

	Gtk_file_chooser_error_get_type func() GType

	Gtk_file_filter_flags_get_type func() GType

	Gtk_icon_lookup_flags_get_type func() GType

	Gtk_icon_theme_error_get_type func() GType

	Gtk_icon_view_drop_position_get_type func() GType

	Gtk_image_type_get_type func() GType

	Gtk_buttons_type_get_type func() GType

	Gtk_notebook_tab_get_type func() GType

	Gtk_object_flags_get_type func() GType

	Gtk_arg_flags_get_type func() GType

	Gtk_print_status_get_type func() GType

	Gtk_print_operation_result_get_type func() GType

	Gtk_print_operation_action_get_type func() GType

	Gtk_print_error_get_type func() GType

	Gtk_private_flags_get_type func() GType

	Gtk_progress_bar_style_get_type func() GType

	Gtk_progress_bar_orientation_get_type func() GType

	Gtk_rc_flags_get_type func() GType

	Gtk_rc_token_type_get_type func() GType

	Gtk_recent_sort_type_get_type func() GType

	Gtk_recent_chooser_error_get_type func() GType

	Gtk_recent_filter_flags_get_type func() GType

	Gtk_recent_manager_error_get_type func() GType

	Gtk_size_group_mode_get_type func() GType

	Gtk_spin_button_update_policy_get_type func() GType

	Gtk_spin_type_get_type func() GType

	Gtk_text_buffer_target_info_get_type func() GType

	Gtk_text_search_flags_get_type func() GType

	Gtk_text_window_type_get_type func() GType

	Gtk_toolbar_child_type_get_type func() GType

	Gtk_toolbar_space_style_get_type func() GType

	Gtk_tool_palette_drag_targets_get_type func() GType

	Gtk_tree_model_flags_get_type func() GType

	Gtk_tree_view_drop_position_get_type func() GType

	Gtk_tree_view_column_sizing_get_type func() GType

	Gtk_ui_manager_item_type_get_type func() GType

	Gtk_widget_flags_get_type func() GType

	Gtk_widget_help_type_get_type func() GType

	Gtk_tree_view_mode_get_type func() GType

	Gtk_cell_type_get_type func() GType

	Gtk_clist_drag_pos_get_type func() GType

	Gtk_button_action_get_type func() GType

	Gtk_ctree_pos_get_type func() GType

	Gtk_ctree_line_style_get_type func() GType

	Gtk_ctree_expander_style_get_type func() GType

	Gtk_ctree_expansion_type_get_type func() GType

	Gtk_identifier_get_type func() GType

	Gtk_type_init func(
		debug_flags GTypeDebugFlags)

	Gtk_type_unique func(
		parent_type GtkType,
		gtkinfo *GtkTypeInfo) GtkType

	Gtk_type_class func(
		t GtkType) Gpointer

	Gtk_type_new func(
		t GtkType) Gpointer

	Gtk_object_get_type func() GType

	Gtk_object_sink func(
		object *GtkObject)

	Gtk_object_destroy func(
		object *GtkObject)

	Gtk_object_new func(t GType, first_property_name string,
		v ...VArg) *GtkObject

	Gtk_object_ref func(
		object *GtkObject) *GtkObject

	Gtk_object_unref func(
		object *GtkObject)

	Gtk_object_weakref func(
		object *GtkObject,
		notify GDestroyNotify,
		data Gpointer)

	Gtk_object_weakunref func(
		object *GtkObject,
		notify GDestroyNotify,
		data Gpointer)

	Gtk_object_set_data func(
		object *GtkObject,
		key string,
		data Gpointer)

	Gtk_object_set_data_full func(
		object *GtkObject,
		key string,
		data Gpointer,
		destroy GDestroyNotify)

	Gtk_object_remove_data func(
		object *GtkObject,
		key string)

	Gtk_object_get_data func(
		object *GtkObject,
		key string) Gpointer

	Gtk_object_remove_no_notify func(
		object *GtkObject,
		key string)

	Gtk_object_set_user_data func(
		object *GtkObject,
		data Gpointer)

	Gtk_object_get_user_data func(
		object *GtkObject) Gpointer

	Gtk_object_set_data_by_id func(
		object *GtkObject,
		data_id GQuark,
		data Gpointer)

	Gtk_object_set_data_by_id_full func(
		object *GtkObject,
		data_id GQuark,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_object_get_data_by_id func(
		object *GtkObject,
		data_id GQuark) Gpointer

	Gtk_object_remove_data_by_id func(
		object *GtkObject,
		data_id GQuark)

	Gtk_object_remove_no_notify_by_id func(
		object *GtkObject,
		key_id GQuark)

	Gtk_object_get func(object *GtkObject,
		first_property_name string, v ...VArg)

	Gtk_object_set func(object *GtkObject,
		first_property_name string, v ...VArg)

	Gtk_object_add_arg_type func(
		arg_name string,
		arg_type GType,
		arg_flags Guint,
		arg_id Guint)

	Gtk_adjustment_get_type func() GType

	Gtk_adjustment_new func(
		value Gdouble,
		lower Gdouble,
		upper Gdouble,
		step_increment Gdouble,
		page_increment Gdouble,
		page_size Gdouble) *GtkObject

	Gtk_adjustment_changed func(
		adjustment *GtkAdjustment)

	Gtk_adjustment_value_changed func(
		adjustment *GtkAdjustment)

	Gtk_adjustment_clamp_page func(
		adjustment *GtkAdjustment,
		lower Gdouble,
		upper Gdouble)

	Gtk_adjustment_get_value func(
		adjustment *GtkAdjustment) Gdouble

	Gtk_adjustment_set_value func(
		adjustment *GtkAdjustment,
		value Gdouble)

	Gtk_adjustment_get_lower func(
		adjustment *GtkAdjustment) Gdouble

	Gtk_adjustment_set_lower func(
		adjustment *GtkAdjustment,
		lower Gdouble)

	Gtk_adjustment_get_upper func(
		adjustment *GtkAdjustment) Gdouble

	Gtk_adjustment_set_upper func(
		adjustment *GtkAdjustment,
		upper Gdouble)

	Gtk_adjustment_get_step_increment func(
		adjustment *GtkAdjustment) Gdouble

	Gtk_adjustment_set_step_increment func(
		adjustment *GtkAdjustment,
		step_increment Gdouble)

	Gtk_adjustment_get_page_increment func(
		adjustment *GtkAdjustment) Gdouble

	Gtk_adjustment_set_page_increment func(
		adjustment *GtkAdjustment,
		page_increment Gdouble)

	Gtk_adjustment_get_page_size func(
		adjustment *GtkAdjustment) Gdouble

	Gtk_adjustment_set_page_size func(
		adjustment *GtkAdjustment,
		page_size Gdouble)

	Gtk_adjustment_configure func(
		adjustment *GtkAdjustment,
		value Gdouble,
		lower Gdouble,
		upper Gdouble,
		step_increment Gdouble,
		page_increment Gdouble,
		page_size Gdouble)

	Gtk_style_get_type func() GType

	Gtk_style_new func() *GtkStyle

	Gtk_style_copy func(
		style *GtkStyle) *GtkStyle

	Gtk_style_attach func(
		style *GtkStyle,
		window *GdkWindow) *GtkStyle

	Gtk_style_detach func(
		style *GtkStyle)

	Gtk_style_ref func(
		style *GtkStyle) *GtkStyle

	Gtk_style_unref func(
		style *GtkStyle)

	Gtk_style_get_font func(
		style *GtkStyle) *GdkFont

	Gtk_style_set_font func(
		style *GtkStyle,
		font *GdkFont)

	Gtk_style_set_background func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType)

	Gtk_style_apply_default_background func(
		style *GtkStyle,
		window *GdkWindow,
		set_bg Gboolean,
		state_type GtkStateType,
		area *GdkRectangle,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_style_lookup_icon_set func(
		style *GtkStyle,
		stock_id string) *GtkIconSet

	Gtk_style_lookup_color func(
		style *GtkStyle,
		color_name string,
		color *GdkColor) Gboolean

	Gtk_style_render_icon func(
		style *GtkStyle,
		source *GtkIconSource,
		direction GtkTextDirection,
		state GtkStateType,
		size GtkIconSize,
		widget *GtkWidget,
		detail string) *GdkPixbuf

	Gtk_draw_hline func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		x1 Gint,
		x2 Gint,
		y Gint)

	Gtk_draw_vline func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		y1_ Gint,
		y2_ Gint,
		x Gint)

	Gtk_draw_shadow func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_polygon func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		points *GdkPoint,
		npoints Gint,
		fill Gboolean)

	Gtk_draw_arrow func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		arrow_type GtkArrowType,
		fill Gboolean,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_diamond func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_box func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_flat_box func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_check func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_option func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_tab func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_shadow_gap func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		gap_side GtkPositionType,
		gap_x Gint,
		gap_width Gint)

	Gtk_draw_box_gap func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		gap_side GtkPositionType,
		gap_x Gint,
		gap_width Gint)

	Gtk_draw_extension func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		gap_side GtkPositionType)

	Gtk_draw_focus func(
		style *GtkStyle,
		window *GdkWindow,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_draw_slider func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		orientation GtkOrientation)

	Gtk_draw_handle func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		orientation GtkOrientation)

	Gtk_draw_expander func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		x Gint,
		y Gint,
		expander_style GtkExpanderStyle)

	Gtk_draw_layout func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		use_text Gboolean,
		x Gint,
		y Gint,
		layout *PangoLayout)

	Gtk_draw_resize_grip func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		edge GdkWindowEdge,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_hline func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x1 Gint,
		x2 Gint,
		y Gint)

	Gtk_paint_vline func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		y1_ Gint,
		y2_ Gint,
		x Gint)

	Gtk_paint_shadow func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_polygon func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		points *GdkPoint,
		n_points Gint,
		fill Gboolean)

	Gtk_paint_arrow func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		arrow_type GtkArrowType,
		fill Gboolean,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_diamond func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_box func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_flat_box func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_check func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_option func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_tab func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_shadow_gap func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		gap_side GtkPositionType,
		gap_x Gint,
		gap_width Gint)

	Gtk_paint_box_gap func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		gap_side GtkPositionType,
		gap_x Gint,
		gap_width Gint)

	Gtk_paint_extension func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		gap_side GtkPositionType)

	Gtk_paint_focus func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_slider func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		orientation GtkOrientation)

	Gtk_paint_handle func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		shadow_type GtkShadowType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		orientation GtkOrientation)

	Gtk_paint_expander func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		expander_style GtkExpanderStyle)

	Gtk_paint_layout func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		use_text Gboolean,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		layout *PangoLayout)

	Gtk_paint_resize_grip func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		edge GdkWindowEdge,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_paint_spinner func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		step Guint,
		x Gint,
		y Gint,
		width Gint,
		height Gint)

	Gtk_border_get_type func() GType

	Gtk_border_new func() *GtkBorder

	Gtk_border_copy func(
		border_ *GtkBorder) *GtkBorder

	Gtk_border_free func(
		border_ *GtkBorder)

	Gtk_style_get_style_property func(
		style *GtkStyle,
		widget_type GType,
		property_name string,
		value *GValue)

	Gtk_style_get_valist func(
		style *GtkStyle,
		widget_type GType,
		first_property_name string,
		var_args Va_list)

	Gtk_style_get func(style *GtkStyle, widget_type GType,
		first_property_name string, v ...VArg)

	Gtk_draw_string func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		x Gint,
		y Gint,
		string string)

	Gtk_paint_string func(
		style *GtkStyle,
		window *GdkWindow,
		state_type GtkStateType,
		area *GdkRectangle,
		widget *GtkWidget,
		detail string,
		x Gint,
		y Gint,
		string string)

	Gtk_draw_insertion_cursor func(
		widget *GtkWidget,
		drawable *GdkDrawable,
		area *GdkRectangle,
		location *GdkRectangle,
		is_primary Gboolean,
		direction GtkTextDirection,
		draw_arrow Gboolean)

	Gtk_rc_add_default_file func(
		filename string)

	Gtk_rc_set_default_files func(
		filenames **Gchar)

	Gtk_rc_get_default_files func() **Gchar

	Gtk_rc_get_style func(
		widget *GtkWidget) *GtkStyle

	Gtk_rc_get_style_by_paths func(
		settings *GtkSettings,
		widget_path string,
		class_path string,
		t GType) *GtkStyle

	Gtk_rc_reparse_all_for_settings func(
		settings *GtkSettings,
		force_load Gboolean) Gboolean

	Gtk_rc_reset_styles func(
		settings *GtkSettings)

	Gtk_rc_find_pixmap_in_path func(
		settings *GtkSettings,
		scanner *GScanner,
		pixmap_file string) string

	Gtk_rc_parse func(
		filename string)

	Gtk_rc_parse_string func(
		rc_string string)

	Gtk_rc_reparse_all func() Gboolean

	Gtk_rc_add_widget_name_style func(
		rc_style *GtkRcStyle,
		pattern string)

	Gtk_rc_add_widget_class_style func(
		rc_style *GtkRcStyle,
		pattern string)

	Gtk_rc_add_class_style func(
		rc_style *GtkRcStyle,
		pattern string)

	Gtk_rc_style_get_type func() GType

	Gtk_rc_style_new func() *GtkRcStyle

	Gtk_rc_style_copy func(
		orig *GtkRcStyle) *GtkRcStyle

	Gtk_rc_style_ref func(
		rc_style *GtkRcStyle)

	Gtk_rc_style_unref func(
		rc_style *GtkRcStyle)

	Gtk_rc_find_module_in_path func(
		module_file string) string

	Gtk_rc_get_theme_dir func() string

	Gtk_rc_get_module_dir func() string

	Gtk_rc_get_im_module_path func() string

	Gtk_rc_get_im_module_file func() string

	Gtk_rc_scanner_new func() *GScanner

	Gtk_rc_parse_color func(
		scanner *GScanner,
		color *GdkColor) Guint

	Gtk_rc_parse_color_full func(
		scanner *GScanner,
		style *GtkRcStyle,
		color *GdkColor) Guint

	Gtk_rc_parse_state func(
		scanner *GScanner,
		state *GtkStateType) Guint

	Gtk_rc_parse_priority func(
		scanner *GScanner,
		priority *GtkPathPriorityType) Guint

	Gtk_settings_get_type func() GType

	Gtk_settings_get_default func() *GtkSettings

	Gtk_settings_get_for_screen func(
		screen *GdkScreen) *GtkSettings

	Gtk_settings_install_property func(
		pspec *GParamSpec)

	Gtk_settings_install_property_parser func(
		pspec *GParamSpec,
		parser GtkRcPropertyParser)

	Gtk_rc_property_parse_color func(
		pspec *GParamSpec,
		gstring *GString,
		property_value *GValue) Gboolean

	Gtk_rc_property_parse_enum func(
		pspec *GParamSpec,
		gstring *GString,
		property_value *GValue) Gboolean

	Gtk_rc_property_parse_flags func(
		pspec *GParamSpec,
		gstring *GString,
		property_value *GValue) Gboolean

	Gtk_rc_property_parse_requisition func(
		pspec *GParamSpec,
		gstring *GString,
		property_value *GValue) Gboolean

	Gtk_rc_property_parse_border func(
		pspec *GParamSpec,
		gstring *GString,
		property_value *GValue) Gboolean

	Gtk_settings_set_property_value func(
		settings *GtkSettings,
		name string,
		svalue *GtkSettingsValue)

	Gtk_settings_set_string_property func(
		settings *GtkSettings,
		name string,
		v_string string,
		origin string)

	Gtk_settings_set_long_property func(
		settings *GtkSettings,
		name string,
		v_long Glong,
		origin string)

	Gtk_settings_set_double_property func(
		settings *GtkSettings,
		name string,
		v_double Gdouble,
		origin string)

	Gtk_widget_get_type func() GType

	Gtk_widget_new func(t GType, first_property_name string,
		v ...VArg) *GtkWidget

	Gtk_widget_destroy func(
		widget *GtkWidget)

	Gtk_widget_destroyed func(
		widget *GtkWidget,
		widget_pointer **GtkWidget)

	Gtk_widget_ref func(
		widget *GtkWidget) *GtkWidget

	Gtk_widget_unref func(
		widget *GtkWidget)

	Gtk_widget_set func(widget *GtkWidget,
		first_property_name string, v ...VArg)

	Gtk_widget_hide_all func(
		widget *GtkWidget)

	Gtk_widget_unparent func(
		widget *GtkWidget)

	Gtk_widget_show func(
		widget *GtkWidget)

	Gtk_widget_show_now func(
		widget *GtkWidget)

	Gtk_widget_hide func(
		widget *GtkWidget)

	Gtk_widget_show_all func(
		widget *GtkWidget)

	Gtk_widget_set_no_show_all func(
		widget *GtkWidget,
		no_show_all Gboolean)

	Gtk_widget_get_no_show_all func(
		widget *GtkWidget) Gboolean

	Gtk_widget_map func(
		widget *GtkWidget)

	Gtk_widget_unmap func(
		widget *GtkWidget)

	Gtk_widget_realize func(
		widget *GtkWidget)

	Gtk_widget_unrealize func(
		widget *GtkWidget)

	Gtk_widget_queue_draw func(
		widget *GtkWidget)

	Gtk_widget_queue_draw_area func(
		widget *GtkWidget, x, y, width, height Gint)

	Gtk_widget_queue_clear func(
		widget *GtkWidget)

	Gtk_widget_queue_clear_area func(
		widget *GtkWidget, x, y, width, height Gint)

	Gtk_widget_queue_resize func(widget *GtkWidget)

	Gtk_widget_queue_resize_no_redraw func(widget *GtkWidget)

	Gtk_widget_draw func(widget *GtkWidget, area *GdkRectangle)

	Gtk_widget_size_request func(
		widget *GtkWidget,
		requisition *GtkRequisition)

	Gtk_widget_size_allocate func(
		widget *GtkWidget,
		allocation *GtkAllocation)

	Gtk_widget_get_child_requisition func(
		widget *GtkWidget,
		requisition *GtkRequisition)

	Gtk_widget_add_accelerator func(
		widget *GtkWidget,
		accel_signal string,
		accel_group *GtkAccelGroup,
		accel_key Guint,
		accel_mods GdkModifierType,
		accel_flags GtkAccelFlags)

	Gtk_widget_remove_accelerator func(
		widget *GtkWidget,
		accel_group *GtkAccelGroup,
		accel_key Guint,
		accel_mods GdkModifierType) Gboolean

	Gtk_widget_set_accel_path func(
		widget *GtkWidget,
		accel_path string,
		accel_group *GtkAccelGroup)

	Gtk_widget_list_accel_closures func(
		widget *GtkWidget) *GList

	Gtk_widget_can_activate_accel func(
		widget *GtkWidget,
		signal_id Guint) Gboolean

	Gtk_widget_mnemonic_activate func(
		widget *GtkWidget,
		group_cycling Gboolean) Gboolean

	Gtk_widget_event func(
		widget *GtkWidget,
		event *GdkEvent) Gboolean

	Gtk_widget_send_expose func(
		widget *GtkWidget,
		event *GdkEvent) Gint

	Gtk_widget_send_focus_change func(
		widget *GtkWidget,
		event *GdkEvent) Gboolean

	Gtk_widget_activate func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_scroll_adjustments func(
		widget *GtkWidget,
		hadjustment *GtkAdjustment,
		vadjustment *GtkAdjustment) Gboolean

	Gtk_widget_reparent func(
		widget *GtkWidget,
		new_parent *GtkWidget)

	Gtk_widget_intersect func(
		widget *GtkWidget,
		area *GdkRectangle,
		intersection *GdkRectangle) Gboolean

	Gtk_widget_region_intersect func(
		widget *GtkWidget,
		region *GdkRegion) *GdkRegion

	Gtk_widget_freeze_child_notify func(
		widget *GtkWidget)

	Gtk_widget_child_notify func(
		widget *GtkWidget,
		child_property string)

	Gtk_widget_thaw_child_notify func(
		widget *GtkWidget)

	Gtk_widget_set_can_focus func(
		widget *GtkWidget,
		can_focus Gboolean)

	Gtk_widget_get_can_focus func(
		widget *GtkWidget) Gboolean

	Gtk_widget_has_focus func(
		widget *GtkWidget) Gboolean

	Gtk_widget_is_focus func(
		widget *GtkWidget) Gboolean

	Gtk_widget_grab_focus func(
		widget *GtkWidget)

	Gtk_widget_set_can_default func(
		widget *GtkWidget,
		can_default Gboolean)

	Gtk_widget_get_can_default func(
		widget *GtkWidget) Gboolean

	Gtk_widget_has_default func(
		widget *GtkWidget) Gboolean

	Gtk_widget_grab_default func(
		widget *GtkWidget)

	Gtk_widget_set_receives_default func(
		widget *GtkWidget,
		receives_default Gboolean)

	Gtk_widget_get_receives_default func(
		widget *GtkWidget) Gboolean

	Gtk_widget_has_grab func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_name func(
		widget *GtkWidget,
		name string)

	Gtk_widget_get_name func(
		widget *GtkWidget) string

	Gtk_widget_set_state func(
		widget *GtkWidget,
		state GtkStateType)

	Gtk_widget_get_state func(
		widget *GtkWidget) GtkStateType

	Gtk_widget_set_sensitive func(
		widget *GtkWidget,
		sensitive Gboolean)

	Gtk_widget_get_sensitive func(
		widget *GtkWidget) Gboolean

	Gtk_widget_is_sensitive func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_visible func(
		widget *GtkWidget,
		visible Gboolean)

	Gtk_widget_get_visible func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_has_window func(
		widget *GtkWidget,
		has_window Gboolean)

	Gtk_widget_get_has_window func(
		widget *GtkWidget) Gboolean

	Gtk_widget_is_toplevel func(
		widget *GtkWidget) Gboolean

	Gtk_widget_is_drawable func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_realized func(
		widget *GtkWidget,
		realized Gboolean)

	Gtk_widget_get_realized func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_mapped func(
		widget *GtkWidget,
		mapped Gboolean)

	Gtk_widget_get_mapped func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_app_paintable func(
		widget *GtkWidget,
		app_paintable Gboolean)

	Gtk_widget_get_app_paintable func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_double_buffered func(
		widget *GtkWidget,
		double_buffered Gboolean)

	Gtk_widget_get_double_buffered func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_redraw_on_allocate func(
		widget *GtkWidget,
		redraw_on_allocate Gboolean)

	Gtk_widget_set_parent func(
		widget *GtkWidget,
		parent *GtkWidget)

	Gtk_widget_get_parent func(
		widget *GtkWidget) *GtkWidget

	Gtk_widget_set_parent_window func(
		widget *GtkWidget,
		parent_window *GdkWindow)

	Gtk_widget_get_parent_window func(
		widget *GtkWidget) *GdkWindow

	Gtk_widget_set_child_visible func(
		widget *GtkWidget,
		is_visible Gboolean)

	Gtk_widget_get_child_visible func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_window func(
		widget *GtkWidget,
		window *GdkWindow)

	Gtk_widget_get_window func(
		widget *GtkWidget) *GdkWindow

	Gtk_widget_get_allocation func(
		widget *GtkWidget,
		allocation *GtkAllocation)

	Gtk_widget_set_allocation func(
		widget *GtkWidget,
		allocation *GtkAllocation)

	Gtk_widget_get_requisition func(
		widget *GtkWidget,
		requisition *GtkRequisition)

	Gtk_widget_child_focus func(
		widget *GtkWidget,
		direction GtkDirectionType) Gboolean

	Gtk_widget_keynav_failed func(
		widget *GtkWidget,
		direction GtkDirectionType) Gboolean

	Gtk_widget_error_bell func(
		widget *GtkWidget)

	Gtk_widget_set_size_request func(
		widget *GtkWidget,
		width Gint,
		height Gint)

	Gtk_widget_get_size_request func(
		widget *GtkWidget,
		width *Gint,
		height *Gint)

	Gtk_widget_set_uposition func(
		widget *GtkWidget,
		x Gint,
		y Gint)

	Gtk_widget_set_usize func(
		widget *GtkWidget,
		width Gint,
		height Gint)

	Gtk_widget_set_events func(
		widget *GtkWidget,
		events Gint)

	Gtk_widget_add_events func(
		widget *GtkWidget,
		events Gint)

	Gtk_widget_set_extension_events func(
		widget *GtkWidget,
		mode GdkExtensionMode)

	Gtk_widget_get_extension_events func(
		widget *GtkWidget) GdkExtensionMode

	Gtk_widget_get_toplevel func(
		widget *GtkWidget) *GtkWidget

	Gtk_widget_get_ancestor func(
		widget *GtkWidget,
		widget_type GType) *GtkWidget

	Gtk_widget_get_colormap func(
		widget *GtkWidget) *GdkColormap

	Gtk_widget_get_visual func(
		widget *GtkWidget) *GdkVisual

	Gtk_widget_get_screen func(
		widget *GtkWidget) *GdkScreen

	Gtk_widget_has_screen func(
		widget *GtkWidget) Gboolean

	Gtk_widget_get_display func(
		widget *GtkWidget) *GdkDisplay

	Gtk_widget_get_root_window func(
		widget *GtkWidget) *GdkWindow

	Gtk_widget_get_settings func(
		widget *GtkWidget) *GtkSettings

	Gtk_widget_get_clipboard func(
		widget *GtkWidget,
		selection GdkAtom) *GtkClipboard

	Gtk_widget_get_snapshot func(
		widget *GtkWidget,
		clip_rect *GdkRectangle) *GdkPixmap

	Gtk_widget_get_accessible func(
		widget *GtkWidget) *AtkObject

	Gtk_widget_set_colormap func(
		widget *GtkWidget,
		colormap *GdkColormap)

	Gtk_widget_get_events func(
		widget *GtkWidget) Gint

	Gtk_widget_get_pointer func(
		widget *GtkWidget,
		x *Gint,
		y *Gint)

	Gtk_widget_is_ancestor func(
		widget *GtkWidget,
		ancestor *GtkWidget) Gboolean

	Gtk_widget_translate_coordinates func(
		src_widget *GtkWidget,
		dest_widget *GtkWidget,
		src_x Gint,
		src_y Gint,
		dest_x *Gint,
		dest_y *Gint) Gboolean

	Gtk_widget_hide_on_delete func(
		widget *GtkWidget) Gboolean

	Gtk_widget_style_attach func(
		style *GtkWidget)

	Gtk_widget_has_rc_style func(
		widget *GtkWidget) Gboolean

	Gtk_widget_set_style func(
		widget *GtkWidget,
		style *GtkStyle)

	Gtk_widget_ensure_style func(
		widget *GtkWidget)

	Gtk_widget_get_style func(
		widget *GtkWidget) *GtkStyle

	Gtk_widget_modify_style func(
		widget *GtkWidget,
		style *GtkRcStyle)

	Gtk_widget_get_modifier_style func(
		widget *GtkWidget) *GtkRcStyle

	Gtk_widget_modify_fg func(
		widget *GtkWidget,
		state GtkStateType,
		color *GdkColor)

	Gtk_widget_modify_bg func(
		widget *GtkWidget,
		state GtkStateType,
		color *GdkColor)

	Gtk_widget_modify_text func(
		widget *GtkWidget,
		state GtkStateType,
		color *GdkColor)

	Gtk_widget_modify_base func(
		widget *GtkWidget,
		state GtkStateType,
		color *GdkColor)

	Gtk_widget_modify_cursor func(
		widget *GtkWidget,
		primary *GdkColor,
		secondary *GdkColor)

	Gtk_widget_modify_font func(
		widget *GtkWidget,
		font_desc *PangoFontDescription)

	Gtk_widget_create_pango_context func(
		widget *GtkWidget) *PangoContext

	Gtk_widget_get_pango_context func(
		widget *GtkWidget) *PangoContext

	Gtk_widget_create_pango_layout func(
		widget *GtkWidget,
		text string) *PangoLayout

	Gtk_widget_render_icon func(
		widget *GtkWidget,
		stock_id string,
		size GtkIconSize,
		detail string) *GdkPixbuf

	Gtk_widget_set_composite_name func(
		widget *GtkWidget,
		name string)

	Gtk_widget_get_composite_name func(
		widget *GtkWidget) string

	Gtk_widget_reset_rc_styles func(
		widget *GtkWidget)

	Gtk_widget_push_colormap func(
		cmap *GdkColormap)

	Gtk_widget_push_composite_child func()

	Gtk_widget_pop_composite_child func()

	Gtk_widget_pop_colormap func()

	Gtk_widget_class_install_style_property func(
		klass *GtkWidgetClass,
		pspec *GParamSpec)

	Gtk_widget_class_install_style_property_parser func(
		klass *GtkWidgetClass,
		pspec *GParamSpec,
		parser GtkRcPropertyParser)

	Gtk_widget_class_find_style_property func(
		klass *GtkWidgetClass,
		property_name string) *GParamSpec

	Gtk_widget_class_list_style_properties func(
		klass *GtkWidgetClass,
		n_properties *Guint) **GParamSpec

	Gtk_widget_style_get_property func(
		widget *GtkWidget,
		property_name string,
		value *GValue)

	Gtk_widget_style_get_valist func(
		widget *GtkWidget,
		first_property_name string,
		var_args Va_list)

	Gtk_widget_style_get func(widget *GtkWidget,
		first_property_name string, v ...VArg)

	Gtk_widget_set_default_colormap func(
		colormap *GdkColormap)

	Gtk_widget_get_default_style func() *GtkStyle

	Gtk_widget_get_default_colormap func() *GdkColormap

	Gtk_widget_get_default_visual func() *GdkVisual

	Gtk_widget_set_direction func(
		widget *GtkWidget,
		dir GtkTextDirection)

	Gtk_widget_get_direction func(
		widget *GtkWidget) GtkTextDirection

	Gtk_widget_set_default_direction func(
		dir GtkTextDirection)

	Gtk_widget_get_default_direction func() GtkTextDirection

	Gtk_widget_is_composited func(
		widget *GtkWidget) Gboolean

	Gtk_widget_shape_combine_mask func(
		widget *GtkWidget,
		shape_mask *GdkBitmap,
		offset_x Gint,
		offset_y Gint)

	Gtk_widget_input_shape_combine_mask func(
		widget *GtkWidget,
		shape_mask *GdkBitmap,
		offset_x Gint,
		offset_y Gint)

	Gtk_widget_reset_shapes func(
		widget *GtkWidget)

	Gtk_widget_path func(
		widget *GtkWidget,
		path_length *Guint,
		path **Gchar,
		path_reversed **Gchar)

	Gtk_widget_class_path func(
		widget *GtkWidget,
		path_length *Guint,
		path **Gchar,
		path_reversed **Gchar)

	Gtk_widget_list_mnemonic_labels func(
		widget *GtkWidget) *GList

	Gtk_widget_add_mnemonic_label func(
		widget *GtkWidget,
		label *GtkWidget)

	Gtk_widget_remove_mnemonic_label func(
		widget *GtkWidget,
		label *GtkWidget)

	Gtk_widget_set_tooltip_window func(
		widget *GtkWidget,
		custom_window *GtkWindow)

	Gtk_widget_get_tooltip_window func(
		widget *GtkWidget) *GtkWindow

	Gtk_widget_trigger_tooltip_query func(
		widget *GtkWidget)

	Gtk_widget_set_tooltip_text func(
		widget *GtkWidget,
		text string)

	Gtk_widget_get_tooltip_text func(
		widget *GtkWidget) string

	Gtk_widget_set_tooltip_markup func(
		widget *GtkWidget,
		markup string)

	Gtk_widget_get_tooltip_markup func(
		widget *GtkWidget) string

	Gtk_widget_set_has_tooltip func(
		widget *GtkWidget,
		has_tooltip Gboolean)

	Gtk_widget_get_has_tooltip func(
		widget *GtkWidget) Gboolean

	Gtk_requisition_get_type func() GType

	Gtk_requisition_copy func(
		requisition *GtkRequisition) *GtkRequisition

	Gtk_requisition_free func(
		requisition *GtkRequisition)

	Gtk_container_get_type func() GType

	Gtk_container_set_border_width func(
		container *GtkContainer,
		border_width Guint)

	Gtk_container_get_border_width func(
		container *GtkContainer) Guint

	Gtk_container_add func(
		container *GtkContainer,
		widget *GtkWidget)

	Gtk_container_remove func(
		container *GtkContainer,
		widget *GtkWidget)

	Gtk_container_set_resize_mode func(
		container *GtkContainer,
		resize_mode GtkResizeMode)

	Gtk_container_get_resize_mode func(
		container *GtkContainer) GtkResizeMode

	Gtk_container_check_resize func(
		container *GtkContainer)

	Gtk_container_foreach func(
		container *GtkContainer,
		callback GtkCallback,
		callback_data Gpointer)

	Gtk_container_foreach_full func(
		container *GtkContainer,
		callback GtkCallback,
		marshal GtkCallbackMarshal,
		callback_dataGpointer,
		notify GDestroyNotify)

	Gtk_container_get_children func(
		container *GtkContainer) *GList

	Gtk_container_propagate_expose func(
		container *GtkContainer,
		child *GtkWidget,
		event *GdkEventExpose)

	Gtk_container_set_focus_chain func(
		container *GtkContainer,
		focusable_widgets *GList)

	Gtk_container_get_focus_chain func(
		container *GtkContainer,
		focusable_widgets **GList) Gboolean

	Gtk_container_unset_focus_chain func(
		container *GtkContainer)

	Gtk_container_set_reallocate_redraws func(
		container *GtkContainer,
		needs_redraws Gboolean)

	Gtk_container_set_focus_child func(
		container *GtkContainer,
		child *GtkWidget)

	Gtk_container_get_focus_child func(
		container *GtkContainer) *GtkWidget

	Gtk_container_set_focus_vadjustment func(
		container *GtkContainer,
		adjustment *GtkAdjustment)

	Gtk_container_get_focus_vadjustment func(
		container *GtkContainer) *GtkAdjustment

	Gtk_container_set_focus_hadjustment func(
		container *GtkContainer,
		adjustment *GtkAdjustment)

	Gtk_container_get_focus_hadjustment func(
		container *GtkContainer) *GtkAdjustment

	Gtk_container_resize_children func(
		container *GtkContainer)

	Gtk_container_child_type func(
		container *GtkContainer) GType

	Gtk_container_class_install_child_property func(
		cclass *GtkContainerClass,
		property_id Guint,
		pspec *GParamSpec)

	Gtk_container_class_find_child_property func(
		cclass *GObjectClass,
		property_name string) *GParamSpec

	Gtk_container_class_list_child_properties func(
		cclass *GObjectClass,
		n_properties *Guint) **GParamSpec

	Gtk_container_add_with_properties func(
		container *GtkContainer, widget *GtkWidget,
		first_prop_name string, v ...VArg)

	Gtk_container_child_set func(
		container *GtkContainer, child *GtkWidget,
		first_prop_name string, v ...VArg)

	Gtk_container_child_get func(container *GtkContainer,
		child *GtkWidget, first_prop_name string, v ...VArg)

	Gtk_container_child_set_valist func(
		container *GtkContainer,
		child *GtkWidget,
		first_property_name string,
		var_args Va_list)

	Gtk_container_child_get_valist func(
		container *GtkContainer,
		child *GtkWidget,
		first_property_name string,
		var_args Va_list)

	Gtk_container_child_set_property func(
		container *GtkContainer,
		child *GtkWidget,
		property_name string,
		value *GValue)

	Gtk_container_child_get_property func(
		container *GtkContainer,
		child *GtkWidget,
		property_name string,
		value *GValue)

	Gtk_container_forall func(
		container *GtkContainer,
		callback GtkCallback,
		callback_data Gpointer)

	Gtk_bin_get_type func() GType

	Gtk_bin_get_child func(
		bin *GtkBin) *GtkWidget

	Gtk_window_get_type func() GType

	Gtk_window_new func(
		t GtkWindowType) *GtkWidget

	Gtk_window_set_title func(
		window *GtkWindow,
		title string)

	Gtk_window_get_title func(
		window *GtkWindow) string

	Gtk_window_set_wmclass func(
		window *GtkWindow,
		wmclass_name string,
		wmclass_class string)

	Gtk_window_set_role func(
		window *GtkWindow,
		role string)

	Gtk_window_set_startup_id func(
		window *GtkWindow,
		startup_id string)

	Gtk_window_get_role func(
		window *GtkWindow) string

	Gtk_window_add_accel_group func(
		window *GtkWindow,
		accel_group *GtkAccelGroup)

	Gtk_window_remove_accel_group func(
		window *GtkWindow,
		accel_group *GtkAccelGroup)

	Gtk_window_set_position func(
		window *GtkWindow,
		position GtkWindowPosition)

	Gtk_window_activate_focus func(
		window *GtkWindow) Gboolean

	Gtk_window_set_focus func(
		window *GtkWindow,
		focus *GtkWidget)

	Gtk_window_get_focus func(
		window *GtkWindow) *GtkWidget

	Gtk_window_set_default func(
		window *GtkWindow,
		default_widget *GtkWidget)

	Gtk_window_get_default_widget func(
		window *GtkWindow) *GtkWidget

	Gtk_window_activate_default func(
		window *GtkWindow) Gboolean

	Gtk_window_set_transient_for func(
		window *GtkWindow,
		parent *GtkWindow)

	Gtk_window_get_transient_for func(
		window *GtkWindow) *GtkWindow

	Gtk_window_set_opacity func(
		window *GtkWindow,
		opacity Gdouble)

	Gtk_window_get_opacity func(
		window *GtkWindow) Gdouble

	Gtk_window_set_type_hint func(
		window *GtkWindow,
		hint GdkWindowTypeHint)

	Gtk_window_get_type_hint func(
		window *GtkWindow) GdkWindowTypeHint

	Gtk_window_set_skip_taskbar_hint func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_skip_taskbar_hint func(
		window *GtkWindow) Gboolean

	Gtk_window_set_skip_pager_hint func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_skip_pager_hint func(
		window *GtkWindow) Gboolean

	Gtk_window_set_urgency_hint func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_urgency_hint func(
		window *GtkWindow) Gboolean

	Gtk_window_set_accept_focus func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_accept_focus func(
		window *GtkWindow) Gboolean

	Gtk_window_set_focus_on_map func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_focus_on_map func(
		window *GtkWindow) Gboolean

	Gtk_window_set_destroy_with_parent func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_destroy_with_parent func(
		window *GtkWindow) Gboolean

	Gtk_window_set_mnemonics_visible func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_mnemonics_visible func(
		window *GtkWindow) Gboolean

	Gtk_window_set_resizable func(
		window *GtkWindow,
		resizable Gboolean)

	Gtk_window_get_resizable func(
		window *GtkWindow) Gboolean

	Gtk_window_set_gravity func(
		window *GtkWindow,
		gravity GdkGravity)

	Gtk_window_get_gravity func(
		window *GtkWindow) GdkGravity

	Gtk_window_set_geometry_hints func(
		window *GtkWindow,
		geometry_widget *GtkWidget,
		geometry *GdkGeometry,
		geom_mask GdkWindowHints)

	Gtk_window_set_screen func(
		window *GtkWindow,
		screen *GdkScreen)

	Gtk_window_get_screen func(
		window *GtkWindow) *GdkScreen

	Gtk_window_is_active func(
		window *GtkWindow) Gboolean

	Gtk_window_has_toplevel_focus func(
		window *GtkWindow) Gboolean

	Gtk_window_set_has_frame func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_has_frame func(
		window *GtkWindow) Gboolean

	Gtk_window_set_frame_dimensions func(
		window *GtkWindow,
		left Gint,
		top Gint,
		right Gint,
		bottom Gint)

	Gtk_window_get_frame_dimensions func(
		window *GtkWindow,
		left *Gint,
		top *Gint,
		right *Gint,
		bottom *Gint)

	Gtk_window_set_decorated func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_decorated func(
		window *GtkWindow) Gboolean

	Gtk_window_set_deletable func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_get_deletable func(
		window *GtkWindow) Gboolean

	Gtk_window_set_icon_list func(
		window *GtkWindow,
		list *GList)

	Gtk_window_get_icon_list func(
		window *GtkWindow) *GList

	Gtk_window_set_icon func(
		window *GtkWindow,
		icon *GdkPixbuf)

	Gtk_window_set_icon_name func(
		window *GtkWindow,
		name string)

	Gtk_window_set_icon_from_file func(
		window *GtkWindow,
		filename string,
		err **GError) Gboolean

	Gtk_window_get_icon func(
		window *GtkWindow) *GdkPixbuf

	Gtk_window_get_icon_name func(
		window *GtkWindow) string

	Gtk_window_set_default_icon_list func(
		list *GList)

	Gtk_window_get_default_icon_list func() *GList

	Gtk_window_set_default_icon func(
		icon *GdkPixbuf)

	Gtk_window_set_default_icon_name func(
		name string)

	Gtk_window_get_default_icon_name func() string

	Gtk_window_set_default_icon_from_file func(
		filename string,
		err **GError) Gboolean

	Gtk_window_set_auto_startup_notification func(
		setting Gboolean)

	Gtk_window_set_modal func(
		window *GtkWindow,
		modal Gboolean)

	Gtk_window_get_modal func(
		window *GtkWindow) Gboolean

	Gtk_window_list_toplevels func() *GList

	Gtk_window_add_mnemonic func(
		window *GtkWindow,
		keyval Guint,
		target *GtkWidget)

	Gtk_window_remove_mnemonic func(
		window *GtkWindow,
		keyval Guint,
		target *GtkWidget)

	Gtk_window_mnemonic_activate func(
		window *GtkWindow,
		keyval Guint,
		modifier GdkModifierType) Gboolean

	Gtk_window_set_mnemonic_modifier func(
		window *GtkWindow,
		modifier GdkModifierType)

	Gtk_window_get_mnemonic_modifier func(
		window *GtkWindow) GdkModifierType

	Gtk_window_activate_key func(
		window *GtkWindow,
		event *GdkEventKey) Gboolean

	Gtk_window_propagate_key_event func(
		window *GtkWindow,
		event *GdkEventKey) Gboolean

	Gtk_window_present func(
		window *GtkWindow)

	Gtk_window_present_with_time func(
		window *GtkWindow,
		timestamp Guint32)

	Gtk_window_iconify func(
		window *GtkWindow)

	Gtk_window_deiconify func(
		window *GtkWindow)

	Gtk_window_stick func(
		window *GtkWindow)

	Gtk_window_unstick func(
		window *GtkWindow)

	Gtk_window_maximize func(
		window *GtkWindow)

	Gtk_window_unmaximize func(
		window *GtkWindow)

	Gtk_window_fullscreen func(
		window *GtkWindow)

	Gtk_window_unfullscreen func(
		window *GtkWindow)

	Gtk_window_set_keep_above func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_set_keep_below func(
		window *GtkWindow,
		setting Gboolean)

	Gtk_window_begin_resize_drag func(
		window *GtkWindow,
		edge GdkWindowEdge,
		button Gint,
		root_x Gint,
		root_y Gint,
		timestamp Guint32)

	Gtk_window_begin_move_drag func(
		window *GtkWindow,
		button Gint,
		root_x Gint,
		root_y Gint,
		timestamp Guint32)

	Gtk_window_set_policy func(
		window *GtkWindow,
		allow_shrink Gint,
		allow_grow Gint,
		auto_shrink Gint)

	Gtk_window_set_default_size func(
		window *GtkWindow,
		width Gint,
		height Gint)

	Gtk_window_get_default_size func(
		window *GtkWindow,
		width *Gint,
		height *Gint)

	Gtk_window_resize func(
		window *GtkWindow,
		width Gint,
		height Gint)

	Gtk_window_get_size func(
		window *GtkWindow,
		width *Gint,
		height *Gint)

	Gtk_window_move func(
		window *GtkWindow,
		x Gint,
		y Gint)

	Gtk_window_get_position func(
		window *GtkWindow,
		root_x *Gint,
		root_y *Gint)

	Gtk_window_parse_geometry func(
		window *GtkWindow,
		geometry string) Gboolean

	Gtk_window_get_group func(
		window *GtkWindow) *GtkWindowGroup

	Gtk_window_has_group func(
		window *GtkWindow) Gboolean

	Gtk_window_reshow_with_initial_size func(
		window *GtkWindow)

	Gtk_window_get_window_type func(
		window *GtkWindow) GtkWindowType

	Gtk_window_group_get_type func() GType

	Gtk_window_group_new func() *GtkWindowGroup

	Gtk_window_group_add_window func(
		window_group *GtkWindowGroup,
		window *GtkWindow)

	Gtk_window_group_remove_window func(
		window_group *GtkWindowGroup,
		window *GtkWindow)

	Gtk_window_group_list_windows func(
		window_group *GtkWindowGroup) *GList

	Gtk_window_remove_embedded_xid func(
		window *GtkWindow,
		xid GdkNativeWindow)

	Gtk_window_add_embedded_xid func(
		window *GtkWindow,
		xid GdkNativeWindow)

	Gtk_window_group_get_current_grab func(
		window_group *GtkWindowGroup) *GtkWidget

	Gtk_dialog_get_type func() GType

	Gtk_dialog_new func() *GtkWidget

	Gtk_dialog_new_with_buttons func(
		title string, parent *GtkWindow,
		flags GtkDialogFlags, first_button_text string,
		v ...VArg) *GtkWidget

	Gtk_dialog_add_action_widget func(
		dialog *GtkDialog,
		child *GtkWidget,
		response_id Gint)

	Gtk_dialog_add_button func(
		dialog *GtkDialog,
		button_text string,
		response_id Gint) *GtkWidget

	Gtk_dialog_add_buttons func(dialog *GtkDialog,
		first_button_text string, v ...VArg)

	Gtk_dialog_set_response_sensitive func(
		dialog *GtkDialog,
		response_id Gint,
		setting Gboolean)

	Gtk_dialog_set_default_response func(
		dialog *GtkDialog,
		response_id Gint)

	Gtk_dialog_get_widget_for_response func(
		dialog *GtkDialog,
		response_id Gint) *GtkWidget

	Gtk_dialog_get_response_for_widget func(
		dialog *GtkDialog,
		widget *GtkWidget) Gint

	Gtk_dialog_set_has_separator func(
		dialog *GtkDialog,
		setting Gboolean)

	Gtk_dialog_get_has_separator func(
		dialog *GtkDialog) Gboolean

	Gtk_alternative_dialog_button_order func(
		screen *GdkScreen) Gboolean

	Gtk_dialog_set_alternative_button_order func(
		dialog *GtkDialog, first_response_id Gint, v ...VArg)

	Gtk_dialog_set_alternative_button_order_from_array func(
		dialog *GtkDialog,
		n_params Gint,
		new_order *Gint)

	Gtk_dialog_response func(
		dialog *GtkDialog,
		response_id Gint)

	Gtk_dialog_run func(
		dialog *GtkDialog) Gint

	Gtk_dialog_get_action_area func(
		dialog *GtkDialog) *GtkWidget

	Gtk_dialog_get_content_area func(
		dialog *GtkDialog) *GtkWidget

	Gtk_about_dialog_get_type func() GType

	Gtk_about_dialog_new func() *GtkWidget

	Gtk_show_about_dialog func(parent *GtkWindow,
		first_property_name string, v ...VArg)

	Gtk_about_dialog_get_name func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_name func(
		about *GtkAboutDialog,
		name string)

	Gtk_about_dialog_get_program_name func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_program_name func(
		about *GtkAboutDialog,
		name string)

	Gtk_about_dialog_get_version func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_version func(
		about *GtkAboutDialog,
		version string)

	Gtk_about_dialog_get_copyright func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_copyright func(
		about *GtkAboutDialog,
		copyright string)

	Gtk_about_dialog_get_comments func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_comments func(
		about *GtkAboutDialog,
		comments string)

	Gtk_about_dialog_get_license func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_license func(
		about *GtkAboutDialog,
		license string)

	Gtk_about_dialog_get_wrap_license func(
		about *GtkAboutDialog) Gboolean

	Gtk_about_dialog_set_wrap_license func(
		about *GtkAboutDialog,
		wrap_license Gboolean)

	Gtk_about_dialog_get_website func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_website func(
		about *GtkAboutDialog,
		website string)

	Gtk_about_dialog_get_website_label func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_website_label func(
		about *GtkAboutDialog,
		website_label string)

	Gtk_about_dialog_get_authors func(
		about *GtkAboutDialog) **Gchar

	Gtk_about_dialog_set_authors func(
		about *GtkAboutDialog,
		authors **Gchar)

	Gtk_about_dialog_get_documenters func(
		about *GtkAboutDialog) **Gchar

	Gtk_about_dialog_set_documenters func(
		about *GtkAboutDialog,
		documenters **Gchar)

	Gtk_about_dialog_get_artists func(
		about *GtkAboutDialog) **Gchar

	Gtk_about_dialog_set_artists func(
		about *GtkAboutDialog,
		artists **Gchar)

	Gtk_about_dialog_get_translator_credits func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_translator_credits func(
		about *GtkAboutDialog,
		translator_credits string)

	Gtk_about_dialog_get_logo func(
		about *GtkAboutDialog) *GdkPixbuf

	Gtk_about_dialog_set_logo func(
		about *GtkAboutDialog,
		logo *GdkPixbuf)

	Gtk_about_dialog_get_logo_icon_name func(
		about *GtkAboutDialog) string

	Gtk_about_dialog_set_logo_icon_name func(
		about *GtkAboutDialog,
		icon_name string)

	Gtk_about_dialog_set_email_hook func(
		f GtkAboutDialogActivateLinkFunc,
		dataGpointer,
		destroy GDestroyNotify) GtkAboutDialogActivateLinkFunc

	Gtk_about_dialog_set_url_hook func(
		f GtkAboutDialogActivateLinkFunc,
		dataGpointer,
		destroy GDestroyNotify) GtkAboutDialogActivateLinkFunc

	Gtk_misc_get_type func() GType

	Gtk_misc_set_alignment func(
		misc *GtkMisc,
		xalign Gfloat,
		yalign Gfloat)

	Gtk_misc_get_alignment func(
		misc *GtkMisc,
		xalign *Gfloat,
		yalign *Gfloat)

	Gtk_misc_set_padding func(
		misc *GtkMisc,
		xpad Gint,
		ypad Gint)

	Gtk_misc_get_padding func(
		misc *GtkMisc,
		xpad *Gint,
		ypad *Gint)

	Gtk_menu_shell_get_type func() GType

	Gtk_menu_shell_append func(
		menu_shell *GtkMenuShell,
		child *GtkWidget)

	Gtk_menu_shell_prepend func(
		menu_shell *GtkMenuShell,
		child *GtkWidget)

	Gtk_menu_shell_insert func(
		menu_shell *GtkMenuShell,
		child *GtkWidget,
		position Gint)

	Gtk_menu_shell_deactivate func(
		menu_shell *GtkMenuShell)

	Gtk_menu_shell_select_item func(
		menu_shell *GtkMenuShell,
		menu_item *GtkWidget)

	Gtk_menu_shell_deselect func(
		menu_shell *GtkMenuShell)

	Gtk_menu_shell_activate_item func(
		menu_shell *GtkMenuShell,
		menu_item *GtkWidget,
		force_deactivate Gboolean)

	Gtk_menu_shell_select_first func(
		menu_shell *GtkMenuShell,
		search_sensitive Gboolean)

	Gtk_menu_shell_cancel func(
		menu_shell *GtkMenuShell)

	Gtk_menu_shell_get_take_focus func(
		menu_shell *GtkMenuShell) Gboolean

	Gtk_menu_shell_set_take_focus func(
		menu_shell *GtkMenuShell,
		take_focus Gboolean)

	Gtk_menu_get_type func() GType

	Gtk_menu_new func() *GtkWidget

	Gtk_menu_popup func(
		menu *GtkMenu,
		parent_menu_shell *GtkWidget,
		parent_menu_item *GtkWidget,
		f GtkMenuPositionFunc,
		dataGpointer,
		button Guint,
		activate_time Guint32)

	Gtk_menu_reposition func(
		menu *GtkMenu)

	Gtk_menu_popdown func(
		menu *GtkMenu)

	Gtk_menu_get_active func(
		menu *GtkMenu) *GtkWidget

	Gtk_menu_set_active func(
		menu *GtkMenu,
		index_ Guint)

	Gtk_menu_set_accel_group func(
		menu *GtkMenu,
		accel_group *GtkAccelGroup)

	Gtk_menu_get_accel_group func(
		menu *GtkMenu) *GtkAccelGroup

	Gtk_menu_set_accel_path func(
		menu *GtkMenu,
		accel_path string)

	Gtk_menu_get_accel_path func(
		menu *GtkMenu) string

	Gtk_menu_attach_to_widget func(
		menu *GtkMenu,
		attach_widget *GtkWidget,
		detacher GtkMenuDetachFunc)

	Gtk_menu_detach func(
		menu *GtkMenu)

	Gtk_menu_get_attach_widget func(
		menu *GtkMenu) *GtkWidget

	Gtk_menu_set_tearoff_state func(
		menu *GtkMenu,
		torn_off Gboolean)

	Gtk_menu_get_tearoff_state func(
		menu *GtkMenu) Gboolean

	Gtk_menu_set_title func(
		menu *GtkMenu,
		title string)

	Gtk_menu_get_title func(
		menu *GtkMenu) string

	Gtk_menu_reorder_child func(
		menu *GtkMenu,
		child *GtkWidget,
		position Gint)

	Gtk_menu_set_screen func(
		menu *GtkMenu,
		screen *GdkScreen)

	Gtk_menu_attach func(
		menu *GtkMenu,
		child *GtkWidget,
		left_attach Guint,
		right_attach Guint,
		top_attach Guint,
		bottom_attach Guint)

	Gtk_menu_set_monitor func(
		menu *GtkMenu,
		monitor_num Gint)

	Gtk_menu_get_monitor func(
		menu *GtkMenu) Gint

	Gtk_menu_get_for_attach_widget func(
		widget *GtkWidget) *GList

	Gtk_menu_set_reserve_toggle_size func(
		menu *GtkMenu,
		reserve_toggle_size Gboolean)

	Gtk_menu_get_reserve_toggle_size func(
		menu *GtkMenu) Gboolean

	Gtk_label_get_type func() GType

	Gtk_label_new func(
		str string) *GtkWidget

	Gtk_label_new_with_mnemonic func(
		str string) *GtkWidget

	Gtk_label_set_text func(
		label *GtkLabel,
		str string)

	Gtk_label_get_text func(
		label *GtkLabel) string

	Gtk_label_set_attributes func(
		label *GtkLabel,
		attrs *PangoAttrList)

	Gtk_label_get_attributes func(
		label *GtkLabel) *PangoAttrList

	Gtk_label_set_label func(
		label *GtkLabel,
		str string)

	Gtk_label_get_label func(
		label *GtkLabel) string

	Gtk_label_set_markup func(
		label *GtkLabel,
		str string)

	Gtk_label_set_use_markup func(
		label *GtkLabel,
		setting Gboolean)

	Gtk_label_get_use_markup func(
		label *GtkLabel) Gboolean

	Gtk_label_set_use_underline func(
		label *GtkLabel,
		setting Gboolean)

	Gtk_label_get_use_underline func(
		label *GtkLabel) Gboolean

	Gtk_label_set_markup_with_mnemonic func(
		label *GtkLabel,
		str string)

	Gtk_label_get_mnemonic_keyval func(
		label *GtkLabel) Guint

	Gtk_label_set_mnemonic_widget func(
		label *GtkLabel,
		widget *GtkWidget)

	Gtk_label_get_mnemonic_widget func(
		label *GtkLabel) *GtkWidget

	Gtk_label_set_text_with_mnemonic func(
		label *GtkLabel,
		str string)

	Gtk_label_set_justify func(
		label *GtkLabel,
		jtype GtkJustification)

	Gtk_label_get_justify func(
		label *GtkLabel) GtkJustification

	Gtk_label_set_ellipsize func(
		label *GtkLabel,
		mode PangoEllipsizeMode)

	Gtk_label_get_ellipsize func(
		label *GtkLabel) PangoEllipsizeMode

	Gtk_label_set_width_chars func(
		label *GtkLabel,
		n_chars Gint)

	Gtk_label_get_width_chars func(
		label *GtkLabel) Gint

	Gtk_label_set_max_width_chars func(
		label *GtkLabel,
		n_chars Gint)

	Gtk_label_get_max_width_chars func(
		label *GtkLabel) Gint

	Gtk_label_set_pattern func(
		label *GtkLabel,
		pattern string)

	Gtk_label_set_line_wrap func(
		label *GtkLabel,
		wrap Gboolean)

	Gtk_label_get_line_wrap func(
		label *GtkLabel) Gboolean

	Gtk_label_set_line_wrap_mode func(
		label *GtkLabel,
		wrap_mode PangoWrapMode)

	Gtk_label_get_line_wrap_mode func(
		label *GtkLabel) PangoWrapMode

	Gtk_label_set_selectable func(
		label *GtkLabel,
		setting Gboolean)

	Gtk_label_get_selectable func(
		label *GtkLabel) Gboolean

	Gtk_label_set_angle func(
		label *GtkLabel,
		angle Gdouble)

	Gtk_label_get_angle func(
		label *GtkLabel) Gdouble

	Gtk_label_select_region func(
		label *GtkLabel,
		start_offset Gint,
		end_offset Gint)

	Gtk_label_get_selection_bounds func(
		label *GtkLabel,
		start *Gint,
		end *Gint) Gboolean

	Gtk_label_get_layout func(
		label *GtkLabel) *PangoLayout

	Gtk_label_get_layout_offsets func(
		label *GtkLabel,
		x *Gint,
		y *Gint)

	Gtk_label_set_single_line_mode func(
		label *GtkLabel,
		single_line_mode Gboolean)

	Gtk_label_get_single_line_mode func(
		label *GtkLabel) Gboolean

	Gtk_label_get_current_uri func(
		label *GtkLabel) string

	Gtk_label_set_track_visited_links func(
		label *GtkLabel,
		track_links Gboolean)

	Gtk_label_get_track_visited_links func(
		label *GtkLabel) Gboolean

	Gtk_label_get func(
		label *GtkLabel,
		str **Gchar)

	Gtk_label_parse_uline func(
		label *GtkLabel,
		string string) Guint

	Gtk_accel_label_get_type func() GType

	Gtk_accel_label_new func(
		string string) *GtkWidget

	Gtk_accel_label_get_accel_widget func(
		accel_label *GtkAccelLabel) *GtkWidget

	Gtk_accel_label_get_accel_width func(
		accel_label *GtkAccelLabel) Guint

	Gtk_accel_label_set_accel_widget func(
		accel_label *GtkAccelLabel,
		accel_widget *GtkWidget)

	Gtk_accel_label_set_accel_closure func(
		accel_label *GtkAccelLabel,
		accel_closure *GClosure)

	Gtk_accel_label_refetch func(
		accel_label *GtkAccelLabel) Gboolean

	Gtk_accel_map_add_entry func(
		accel_path string,
		accel_key Guint,
		accel_mods GdkModifierType)

	Gtk_accel_map_lookup_entry func(
		accel_path string,
		key *GtkAccelKey) Gboolean

	Gtk_accel_map_change_entry func(
		accel_path string,
		accel_key Guint,
		accel_mods GdkModifierType,
		replace Gboolean) Gboolean

	Gtk_accel_map_load func(
		file_name string)

	Gtk_accel_map_save func(
		file_name string)

	Gtk_accel_map_foreach func(
		dataGpointer,
		foreach_func GtkAccelMapForeach)

	Gtk_accel_map_load_fd func(
		fd Gint)

	Gtk_accel_map_load_scanner func(
		scanner *GScanner)

	Gtk_accel_map_save_fd func(
		fd Gint)

	Gtk_accel_map_lock_path func(
		accel_path string)

	Gtk_accel_map_unlock_path func(
		accel_path string)

	Gtk_accel_map_add_filter func(
		filter_pattern string)

	Gtk_accel_map_foreach_unfiltered func(
		dataGpointer,
		foreach_func GtkAccelMapForeach)

	Gtk_accel_map_get_type func() GType

	Gtk_accel_map_get func() *GtkAccelMap

	Gtk_accessible_get_type func() GType

	Gtk_accessible_set_widget func(
		accessible *GtkAccessible,
		widget *GtkWidget)

	Gtk_accessible_get_widget func(
		accessible *GtkAccessible) *GtkWidget

	Gtk_accessible_connect_widget_destroyed func(
		accessible *GtkAccessible)

	Gtk_action_get_type func() GType

	Gtk_action_new func(
		name string,
		label string,
		tooltip string,
		stock_id string) *GtkAction

	Gtk_action_get_name func(
		action *GtkAction) string

	Gtk_action_is_sensitive func(
		action *GtkAction) Gboolean

	Gtk_action_get_sensitive func(
		action *GtkAction) Gboolean

	Gtk_action_set_sensitive func(
		action *GtkAction,
		sensitive Gboolean)

	Gtk_action_is_visible func(
		action *GtkAction) Gboolean

	Gtk_action_get_visible func(
		action *GtkAction) Gboolean

	Gtk_action_set_visible func(
		action *GtkAction,
		visible Gboolean)

	Gtk_action_activate func(
		action *GtkAction)

	Gtk_action_create_icon func(
		action *GtkAction,
		icon_size GtkIconSize) *GtkWidget

	Gtk_action_create_menu_item func(
		action *GtkAction) *GtkWidget

	Gtk_action_create_tool_item func(
		action *GtkAction) *GtkWidget

	Gtk_action_create_menu func(
		action *GtkAction) *GtkWidget

	Gtk_action_get_proxies func(
		action *GtkAction) *GSList

	Gtk_action_connect_accelerator func(
		action *GtkAction)

	Gtk_action_disconnect_accelerator func(
		action *GtkAction)

	Gtk_action_get_accel_path func(
		action *GtkAction) string

	Gtk_action_get_accel_closure func(
		action *GtkAction) *GClosure

	Gtk_widget_get_action func(
		widget *GtkWidget) *GtkAction

	Gtk_action_connect_proxy func(
		action *GtkAction,
		proxy *GtkWidget)

	Gtk_action_disconnect_proxy func(
		action *GtkAction,
		proxy *GtkWidget)

	Gtk_action_block_activate_from func(
		action *GtkAction,
		proxy *GtkWidget)

	Gtk_action_unblock_activate_from func(
		action *GtkAction,
		proxy *GtkWidget)

	Gtk_action_block_activate func(
		action *GtkAction)

	Gtk_action_unblock_activate func(
		action *GtkAction)

	Gtk_action_set_accel_path func(
		action *GtkAction,
		accel_path string)

	Gtk_action_set_accel_group func(
		action *GtkAction,
		accel_group *GtkAccelGroup)

	Gtk_action_set_label func(
		action *GtkAction,
		label string)

	Gtk_action_get_label func(
		action *GtkAction) string

	Gtk_action_set_short_label func(
		action *GtkAction,
		short_label string)

	Gtk_action_get_short_label func(
		action *GtkAction) string

	Gtk_action_set_tooltip func(
		action *GtkAction,
		tooltip string)

	Gtk_action_get_tooltip func(
		action *GtkAction) string

	Gtk_action_set_stock_id func(
		action *GtkAction,
		stock_id string)

	Gtk_action_get_stock_id func(
		action *GtkAction) string

	Gtk_action_set_gicon func(
		action *GtkAction,
		icon *GIcon)

	Gtk_action_get_gicon func(
		action *GtkAction) *GIcon

	Gtk_action_set_icon_name func(
		action *GtkAction,
		icon_name string)

	Gtk_action_get_icon_name func(
		action *GtkAction) string

	Gtk_action_set_visible_horizontal func(
		action *GtkAction,
		visible_horizontal Gboolean)

	Gtk_action_get_visible_horizontal func(
		action *GtkAction) Gboolean

	Gtk_action_set_visible_vertical func(
		action *GtkAction,
		visible_vertical Gboolean)

	Gtk_action_get_visible_vertical func(
		action *GtkAction) Gboolean

	Gtk_action_set_is_important func(
		action *GtkAction,
		is_important Gboolean)

	Gtk_action_get_is_important func(
		action *GtkAction) Gboolean

	Gtk_action_set_always_show_image func(
		action *GtkAction,
		always_show Gboolean)

	Gtk_action_get_always_show_image func(
		action *GtkAction) Gboolean

	Gtk_action_group_get_type func() GType

	Gtk_action_group_new func(
		name string) *GtkActionGroup

	Gtk_action_group_get_name func(
		action_group *GtkActionGroup) string

	Gtk_action_group_get_sensitive func(
		action_group *GtkActionGroup) Gboolean

	Gtk_action_group_set_sensitive func(
		action_group *GtkActionGroup,
		sensitive Gboolean)

	Gtk_action_group_get_visible func(
		action_group *GtkActionGroup) Gboolean

	Gtk_action_group_set_visible func(
		action_group *GtkActionGroup,
		visible Gboolean)

	Gtk_action_group_get_action func(
		action_group *GtkActionGroup,
		action_name string) *GtkAction

	Gtk_action_group_list_actions func(
		action_group *GtkActionGroup) *GList

	Gtk_action_group_add_action func(
		action_group *GtkActionGroup,
		action *GtkAction)

	Gtk_action_group_add_action_with_accel func(
		action_group *GtkActionGroup,
		action *GtkAction,
		accelerator string)

	Gtk_action_group_remove_action func(
		action_group *GtkActionGroup,
		action *GtkAction)

	Gtk_action_group_add_actions func(
		action_group *GtkActionGroup,
		entries *GtkActionEntry,
		n_entries Guint,
		user_data Gpointer)

	Gtk_action_group_add_toggle_actions func(
		action_group *GtkActionGroup,
		entries *GtkToggleActionEntry,
		n_entries Guint,
		user_data Gpointer)

	Gtk_action_group_add_radio_actions func(
		action_group *GtkActionGroup,
		entries *GtkRadioActionEntry,
		n_entries Guint,
		value Gint,
		on_change GCallback,
		user_data Gpointer)

	Gtk_action_group_add_actions_full func(
		action_group *GtkActionGroup,
		entries *GtkActionEntry,
		n_entries Guint,
		user_dataGpointer,
		destroy GDestroyNotify)

	Gtk_action_group_add_toggle_actions_full func(
		action_group *GtkActionGroup,
		entries *GtkToggleActionEntry,
		n_entries Guint,
		user_dataGpointer,
		destroy GDestroyNotify)

	Gtk_action_group_add_radio_actions_full func(
		action_group *GtkActionGroup,
		entries *GtkRadioActionEntry,
		n_entries Guint,
		value Gint,
		on_change GCallback,
		user_dataGpointer,
		destroy GDestroyNotify)

	Gtk_action_group_set_translate_func func(
		action_group *GtkActionGroup,
		f GtkTranslateFunc,
		dataGpointer,
		notify GDestroyNotify)

	Gtk_action_group_set_translation_domain func(
		action_group *GtkActionGroup,
		domain string)

	Gtk_action_group_translate_string func(
		action_group *GtkActionGroup,
		string string) string

	Gtk_activatable_get_type func() GType

	Gtk_activatable_sync_action_properties func(
		activatable *GtkActivatable,
		action *GtkAction)

	Gtk_activatable_set_related_action func(
		activatable *GtkActivatable,
		action *GtkAction)

	Gtk_activatable_get_related_action func(
		activatable *GtkActivatable) *GtkAction

	Gtk_activatable_set_use_action_appearance func(
		activatable *GtkActivatable,
		use_appearance Gboolean)

	Gtk_activatable_get_use_action_appearance func(
		activatable *GtkActivatable) Gboolean

	Gtk_activatable_do_set_related_action func(
		activatable *GtkActivatable,
		action *GtkAction)

	Gtk_alignment_get_type func() GType

	Gtk_alignment_new func(
		xalign Gfloat,
		yalign Gfloat,
		xscale Gfloat,
		yscale Gfloat) *GtkWidget

	Gtk_alignment_set func(
		alignment *GtkAlignment,
		xalign Gfloat,
		yalign Gfloat,
		xscale Gfloat,
		yscale Gfloat)

	Gtk_alignment_set_padding func(
		alignment *GtkAlignment,
		padding_top Guint,
		padding_bottom Guint,
		padding_left Guint,
		padding_right Guint)

	Gtk_alignment_get_padding func(
		alignment *GtkAlignment,
		padding_top *Guint,
		padding_bottom *Guint,
		padding_left *Guint,
		padding_right *Guint)

	Gtk_arrow_get_type func() GType

	Gtk_arrow_new func(
		arrow_type GtkArrowType,
		shadow_type GtkShadowType) *GtkWidget

	Gtk_arrow_set func(
		arrow *GtkArrow,
		arrow_type GtkArrowType,
		shadow_type GtkShadowType)

	Gtk_frame_get_type func() GType

	Gtk_frame_new func(
		label string) *GtkWidget

	Gtk_frame_set_label func(
		frame *GtkFrame,
		label string)

	Gtk_frame_get_label func(
		frame *GtkFrame) string

	Gtk_frame_set_label_widget func(
		frame *GtkFrame,
		label_widget *GtkWidget)

	Gtk_frame_get_label_widget func(
		frame *GtkFrame) *GtkWidget

	Gtk_frame_set_label_align func(
		frame *GtkFrame,
		xalign Gfloat,
		yalign Gfloat)

	Gtk_frame_get_label_align func(
		frame *GtkFrame,
		xalign *Gfloat,
		yalign *Gfloat)

	Gtk_frame_set_shadow_type func(
		frame *GtkFrame,
		t GtkShadowType)

	Gtk_frame_get_shadow_type func(
		frame *GtkFrame) GtkShadowType

	Gtk_aspect_frame_get_type func() GType

	Gtk_aspect_frame_new func(
		label string,
		xalign Gfloat,
		yalign Gfloat,
		ratio Gfloat,
		obey_child Gboolean) *GtkWidget

	Gtk_aspect_frame_set func(
		aspect_frame *GtkAspectFrame,
		xalign Gfloat,
		yalign Gfloat,
		ratio Gfloat,
		obey_child Gboolean)

	Gtk_assistant_get_type func() GType

	Gtk_assistant_new func() *GtkWidget

	Gtk_assistant_get_current_page func(
		assistant *GtkAssistant) Gint

	Gtk_assistant_set_current_page func(
		assistant *GtkAssistant,
		page_num Gint)

	Gtk_assistant_get_n_pages func(
		assistant *GtkAssistant) Gint

	Gtk_assistant_get_nth_page func(
		assistant *GtkAssistant,
		page_num Gint) *GtkWidget

	Gtk_assistant_prepend_page func(
		assistant *GtkAssistant,
		page *GtkWidget) Gint

	Gtk_assistant_append_page func(
		assistant *GtkAssistant,
		page *GtkWidget) Gint

	Gtk_assistant_insert_page func(
		assistant *GtkAssistant,
		page *GtkWidget,
		position Gint) Gint

	Gtk_assistant_set_forward_page_func func(
		assistant *GtkAssistant,
		page_func GtkAssistantPageFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_assistant_set_page_type func(
		assistant *GtkAssistant,
		page *GtkWidget,
		t GtkAssistantPageType)

	Gtk_assistant_get_page_type func(
		assistant *GtkAssistant,
		page *GtkWidget) GtkAssistantPageType

	Gtk_assistant_set_page_title func(
		assistant *GtkAssistant,
		page *GtkWidget,
		title string)

	Gtk_assistant_get_page_title func(
		assistant *GtkAssistant,
		page *GtkWidget) string

	Gtk_assistant_set_page_header_image func(
		assistant *GtkAssistant,
		page *GtkWidget,
		pixbuf *GdkPixbuf)

	Gtk_assistant_get_page_header_image func(
		assistant *GtkAssistant,
		page *GtkWidget) *GdkPixbuf

	Gtk_assistant_set_page_side_image func(
		assistant *GtkAssistant,
		page *GtkWidget,
		pixbuf *GdkPixbuf)

	Gtk_assistant_get_page_side_image func(
		assistant *GtkAssistant,
		page *GtkWidget) *GdkPixbuf

	Gtk_assistant_set_page_complete func(
		assistant *GtkAssistant,
		page *GtkWidget,
		complete Gboolean)

	Gtk_assistant_get_page_complete func(
		assistant *GtkAssistant,
		page *GtkWidget) Gboolean

	Gtk_assistant_add_action_widget func(
		assistant *GtkAssistant,
		child *GtkWidget)

	Gtk_assistant_remove_action_widget func(
		assistant *GtkAssistant,
		child *GtkWidget)

	Gtk_assistant_update_buttons_state func(
		assistant *GtkAssistant)

	Gtk_assistant_commit func(
		assistant *GtkAssistant)

	Gtk_box_get_type func() GType

	Gtk_box_pack_start func(
		box *GtkBox,
		child *GtkWidget,
		expand Gboolean,
		fill Gboolean,
		padding Guint)

	Gtk_box_pack_end func(
		box *GtkBox,
		child *GtkWidget,
		expand Gboolean,
		fill Gboolean,
		padding Guint)

	Gtk_box_pack_start_defaults func(
		box *GtkBox,
		widget *GtkWidget)

	Gtk_box_pack_end_defaults func(
		box *GtkBox,
		widget *GtkWidget)

	Gtk_box_set_homogeneous func(
		box *GtkBox,
		homogeneous Gboolean)

	Gtk_box_get_homogeneous func(
		box *GtkBox) Gboolean

	Gtk_box_set_spacing func(
		box *GtkBox,
		spacing Gint)

	Gtk_box_get_spacing func(
		box *GtkBox) Gint

	Gtk_box_reorder_child func(
		box *GtkBox,
		child *GtkWidget,
		position Gint)

	Gtk_box_query_child_packing func(
		box *GtkBox,
		child *GtkWidget,
		expand *Gboolean,
		fill *Gboolean,
		padding *Guint,
		pack_type *GtkPackType)

	Gtk_box_set_child_packing func(
		box *GtkBox,
		child *GtkWidget,
		expand Gboolean,
		fill Gboolean,
		padding Guint,
		pack_type GtkPackType)

	Gtk_button_box_get_type func() GType

	Gtk_button_box_get_layout func(
		widget *GtkButtonBox) GtkButtonBoxStyle

	Gtk_button_box_set_layout func(
		widget *GtkButtonBox,
		layout_style GtkButtonBoxStyle)

	Gtk_button_box_get_child_secondary func(
		widget *GtkButtonBox,
		child *GtkWidget) Gboolean

	Gtk_button_box_set_child_secondary func(
		widget *GtkButtonBox,
		child *GtkWidget,
		is_secondary Gboolean)

	Gtk_button_box_set_child_size func(
		widget *GtkButtonBox,
		min_width Gint,
		min_height Gint)

	Gtk_button_box_set_child_ipadding func(
		widget *GtkButtonBox,
		ipad_x Gint,
		ipad_y Gint)

	Gtk_button_box_get_child_size func(
		widget *GtkButtonBox,
		min_width *Gint,
		min_height *Gint)

	Gtk_button_box_get_child_ipadding func(
		widget *GtkButtonBox,
		ipad_x *Gint,
		ipad_y *Gint)

	Gtk_binding_set_new func(
		set_name string) *GtkBindingSet

	Gtk_binding_set_by_class func(
		object_class Gpointer) *GtkBindingSet

	Gtk_binding_set_find func(
		set_name string) *GtkBindingSet

	Gtk_bindings_activate func(
		object *GtkObject,
		keyval Guint,
		modifiers GdkModifierType) Gboolean

	Gtk_bindings_activate_event func(
		object *GtkObject,
		event *GdkEventKey) Gboolean

	Gtk_binding_set_activate func(
		binding_set *GtkBindingSet,
		keyval Guint,
		modifiers GdkModifierType,
		object *GtkObject) Gboolean

	Gtk_binding_entry_clear func(
		binding_set *GtkBindingSet,
		keyval Guint,
		modifiers GdkModifierType)

	Gtk_binding_parse_binding func(
		scanner *GScanner) Guint

	Gtk_binding_entry_skip func(
		binding_set *GtkBindingSet,
		keyval Guint,
		modifiers GdkModifierType)

	Gtk_binding_entry_add_signal func(binding_set *GtkBindingSet,
		keyval Guint, modifiers GdkModifierType,
		signal_name string, n_args Guint, varg ...VArg)

	Gtk_binding_entry_add_signall func(
		binding_set *GtkBindingSet,
		keyval Guint,
		modifiers GdkModifierType,
		signal_name string,
		binding_args *GSList)

	Gtk_binding_entry_remove func(
		binding_set *GtkBindingSet,
		keyval Guint,
		modifiers GdkModifierType)

	Gtk_binding_set_add_path func(
		binding_set *GtkBindingSet,
		path_type GtkPathType,
		path_pattern string,
		priority GtkPathPriorityType)

	Gtk_builder_get_type func() GType

	Gtk_builder_new func() *GtkBuilder

	Gtk_builder_add_from_file func(
		builder *GtkBuilder,
		filename string,
		error **GError) Guint

	Gtk_builder_add_from_string func(
		builder *GtkBuilder,
		buffer string,
		length Gsize,
		error **GError) Guint

	Gtk_builder_add_objects_from_file func(
		builder *GtkBuilder,
		filename string,
		object_ids **Gchar,
		error **GError) Guint

	Gtk_builder_add_objects_from_string func(
		builder *GtkBuilder,
		buffer string,
		length Gsize,
		object_ids **Gchar,
		error **GError) Guint

	Gtk_builder_get_object func(
		builder *GtkBuilder,
		name string) *GObject

	Gtk_builder_get_objects func(
		builder *GtkBuilder) *GSList

	Gtk_builder_connect_signals func(
		builder *GtkBuilder,
		user_data Gpointer)

	Gtk_builder_connect_signals_full func(
		builder *GtkBuilder,
		f GtkBuilderConnectFunc,
		user_data Gpointer)

	Gtk_builder_set_translation_domain func(
		builder *GtkBuilder,
		domain string)

	Gtk_builder_get_translation_domain func(
		builder *GtkBuilder) string

	Gtk_builder_get_type_from_name func(
		builder *GtkBuilder,
		type_name string) GType

	Gtk_builder_value_from_string func(
		builder *GtkBuilder,
		pspec *GParamSpec,
		string string,
		value *GValue,
		error **GError) Gboolean

	Gtk_builder_value_from_string_type func(
		builder *GtkBuilder,
		t GType,
		string string,
		value *GValue,
		error **GError) Gboolean

	Gtk_buildable_get_type func() GType

	Gtk_buildable_set_name func(
		buildable *GtkBuildable,
		name string)

	Gtk_buildable_get_name func(
		buildable *GtkBuildable) string

	Gtk_buildable_add_child func(
		buildable *GtkBuildable,
		builder *GtkBuilder,
		child *GObject,
		typ string)

	Gtk_buildable_set_buildable_property func(
		buildable *GtkBuildable,
		builder *GtkBuilder,
		name string,
		value *GValue)

	Gtk_buildable_construct_child func(
		buildable *GtkBuildable,
		builder *GtkBuilder,
		name string) *GObject

	Gtk_buildable_custom_tag_start func(
		buildable *GtkBuildable,
		builder *GtkBuilder,
		child *GObject,
		tagname string,
		parser *GMarkupParser,
		data *Gpointer) Gboolean

	Gtk_buildable_custom_tag_end func(
		buildable *GtkBuildable,
		builder *GtkBuilder,
		child *GObject,
		tagname string,
		data *Gpointer)

	Gtk_buildable_custom_finished func(
		buildable *GtkBuildable,
		builder *GtkBuilder,
		child *GObject,
		tagname string,
		data Gpointer)

	Gtk_buildable_parser_finished func(
		buildable *GtkBuildable,
		builder *GtkBuilder)

	Gtk_buildable_get_internal_child func(
		buildable *GtkBuildable,
		builder *GtkBuilder,
		childname string) *GObject

	Gtk_image_get_type func() GType

	Gtk_image_new func() *GtkWidget

	Gtk_image_new_from_pixmap func(
		pixmap *GdkPixmap,
		mask *GdkBitmap) *GtkWidget

	Gtk_image_new_from_image func(
		image *GdkImage,
		mask *GdkBitmap) *GtkWidget

	Gtk_image_new_from_file func(
		filename string) *GtkWidget

	Gtk_image_new_from_pixbuf func(
		pixbuf *GdkPixbuf) *GtkWidget

	Gtk_image_new_from_stock func(
		stock_id string,
		size GtkIconSize) *GtkWidget

	Gtk_image_new_from_icon_set func(
		icon_set *GtkIconSet,
		size GtkIconSize) *GtkWidget

	Gtk_image_new_from_animation func(
		animation *GdkPixbufAnimation) *GtkWidget

	Gtk_image_new_from_icon_name func(
		icon_name string,
		size GtkIconSize) *GtkWidget

	Gtk_image_new_from_gicon func(
		icon *GIcon,
		size GtkIconSize) *GtkWidget

	Gtk_image_clear func(
		image *GtkImage)

	Gtk_image_set_from_pixmap func(
		image *GtkImage,
		pixmap *GdkPixmap,
		mask *GdkBitmap)

	Gtk_image_set_from_image func(
		image *GtkImage,
		gdk_image *GdkImage,
		mask *GdkBitmap)

	Gtk_image_set_from_file func(
		image *GtkImage,
		filename string)

	Gtk_image_set_from_pixbuf func(
		image *GtkImage,
		pixbuf *GdkPixbuf)

	Gtk_image_set_from_stock func(
		image *GtkImage,
		stock_id string,
		size GtkIconSize)

	Gtk_image_set_from_icon_set func(
		image *GtkImage,
		icon_set *GtkIconSet,
		size GtkIconSize)

	Gtk_image_set_from_animation func(
		image *GtkImage,
		animation *GdkPixbufAnimation)

	Gtk_image_set_from_icon_name func(
		image *GtkImage,
		icon_name string,
		size GtkIconSize)

	Gtk_image_set_from_gicon func(
		image *GtkImage,
		icon *GIcon,
		size GtkIconSize)

	Gtk_image_set_pixel_size func(
		image *GtkImage,
		pixel_size Gint)

	Gtk_image_get_storage_type func(
		image *GtkImage) GtkImageType

	Gtk_image_get_pixmap func(
		image *GtkImage,
		pixmap **GdkPixmap,
		mask **GdkBitmap)

	Gtk_image_get_image func(
		image *GtkImage,
		gdk_image **GdkImage,
		mask **GdkBitmap)

	Gtk_image_get_pixbuf func(
		image *GtkImage) *GdkPixbuf

	Gtk_image_get_stock func(
		image *GtkImage,
		stock_id **Gchar,
		size *GtkIconSize)

	Gtk_image_get_icon_set func(
		image *GtkImage,
		icon_set **GtkIconSet,
		size *GtkIconSize)

	Gtk_image_get_animation func(
		image *GtkImage) *GdkPixbufAnimation

	Gtk_image_get_icon_name func(
		image *GtkImage,
		icon_name **Gchar,
		size *GtkIconSize)

	Gtk_image_get_gicon func(
		image *GtkImage,
		gicon **GIcon,
		size *GtkIconSize)

	Gtk_image_get_pixel_size func(
		image *GtkImage) Gint

	Gtk_image_set func(
		image *GtkImage,
		val *GdkImage,
		mask *GdkBitmap)

	Gtk_image_get func(
		image *GtkImage,
		val **GdkImage,
		mask **GdkBitmap)

	Gtk_button_get_type func() GType

	Gtk_button_new func() *GtkWidget

	Gtk_button_new_with_label func(
		label string) *GtkWidget

	Gtk_button_new_from_stock func(
		stock_id string) *GtkWidget

	Gtk_button_new_with_mnemonic func(
		label string) *GtkWidget

	Gtk_button_pressed func(
		button *GtkButton)

	Gtk_button_released func(
		button *GtkButton)

	Gtk_button_clicked func(
		button *GtkButton)

	Gtk_button_enter func(
		button *GtkButton)

	Gtk_button_leave func(
		button *GtkButton)

	Gtk_button_set_relief func(
		button *GtkButton,
		newstyle GtkReliefStyle)

	Gtk_button_get_relief func(
		button *GtkButton) GtkReliefStyle

	Gtk_button_set_label func(
		button *GtkButton,
		label string)

	Gtk_button_get_label func(
		button *GtkButton) string

	Gtk_button_set_use_underline func(
		button *GtkButton,
		use_underline Gboolean)

	Gtk_button_get_use_underline func(
		button *GtkButton) Gboolean

	Gtk_button_set_use_stock func(
		button *GtkButton,
		use_stock Gboolean)

	Gtk_button_get_use_stock func(
		button *GtkButton) Gboolean

	Gtk_button_set_focus_on_click func(
		button *GtkButton,
		focus_on_click Gboolean)

	Gtk_button_get_focus_on_click func(
		button *GtkButton) Gboolean

	Gtk_button_set_alignment func(
		button *GtkButton,
		xalign Gfloat,
		yalign Gfloat)

	Gtk_button_get_alignment func(
		button *GtkButton,
		xalign *Gfloat,
		yalign *Gfloat)

	Gtk_button_set_image func(
		button *GtkButton,
		image *GtkWidget)

	Gtk_button_get_image func(
		button *GtkButton) *GtkWidget

	Gtk_button_set_image_position func(
		button *GtkButton,
		position GtkPositionType)

	Gtk_button_get_image_position func(
		button *GtkButton) GtkPositionType

	Gtk_button_get_event_window func(
		button *GtkButton) *GdkWindow

	Gtk_signal_newv func(
		name string,
		signal_flags GtkSignalRunType,
		object_type GType,
		function_offset Guint,
		marshaller GSignalCMarshaller,
		return_val GType,
		n_args Guint,
		args *GType) Guint

	Gtk_signal_new func(name string,
		signal_flags GtkSignalRunType,
		object_type GType, function_offset Guint,
		marshaller GSignalCMarshaller, return_val GType,
		n_args Guint, v ...VArg) Guint

	Gtk_signal_emit_stop_by_name func(
		object *GtkObject, name string)

	Gtk_signal_connect_object_while_alive func(
		object *GtkObject,
		name string,
		f GCallback,
		alive_object *GtkObject)

	Gtk_signal_connect_while_alive func(
		object *GtkObject,
		name string,
		f GCallback,
		func_dataGpointer,
		alive_object *GtkObject)

	Gtk_signal_connect_full func(
		object *GtkObject,
		name string,
		f GCallback,
		unsupported GtkCallbackMarshal,
		dataGpointer,
		destroy_func GDestroyNotify,
		object_signal Gint,
		after Gint) Gulong

	Gtk_signal_emitv func(
		object *GtkObject,
		signal_id Guint,
		args *GtkArg)

	Gtk_signal_emit func(object *GtkObject, signal_id Guint,
		v ...VArg)

	Gtk_signal_emit_by_name func(object *GtkObject,
		name string, v ...VArg)

	Gtk_signal_emitv_by_name func(
		object *GtkObject,
		name string,
		args *GtkArg)

	Gtk_signal_compat_matched func(
		object *GtkObject,
		f GCallback,
		dataGpointer,
		match GSignalMatchType,
		action Guint)

	Gtk_calendar_get_type func() GType

	Gtk_calendar_new func() *GtkWidget

	Gtk_calendar_select_month func(
		calendar *GtkCalendar,
		month Guint,
		year Guint) Gboolean

	Gtk_calendar_select_day func(
		calendar *GtkCalendar,
		day Guint)

	Gtk_calendar_mark_day func(
		calendar *GtkCalendar,
		day Guint) Gboolean

	Gtk_calendar_unmark_day func(
		calendar *GtkCalendar,
		day Guint) Gboolean

	Gtk_calendar_clear_marks func(
		calendar *GtkCalendar)

	Gtk_calendar_set_display_options func(
		calendar *GtkCalendar,
		flags GtkCalendarDisplayOptions)

	Gtk_calendar_get_display_options func(
		calendar *GtkCalendar) GtkCalendarDisplayOptions

	Gtk_calendar_display_options func(
		calendar *GtkCalendar,
		flags GtkCalendarDisplayOptions)

	Gtk_calendar_get_date func(
		calendar *GtkCalendar,
		year *Guint,
		month *Guint,
		day *Guint)

	Gtk_calendar_set_detail_func func(
		calendar *GtkCalendar,
		f GtkCalendarDetailFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_calendar_set_detail_width_chars func(
		calendar *GtkCalendar,
		chars Gint)

	Gtk_calendar_set_detail_height_rows func(
		calendar *GtkCalendar,
		rows Gint)

	Gtk_calendar_get_detail_width_chars func(
		calendar *GtkCalendar) Gint

	Gtk_calendar_get_detail_height_rows func(
		calendar *GtkCalendar) Gint

	Gtk_calendar_freeze func(
		calendar *GtkCalendar)

	Gtk_calendar_thaw func(
		calendar *GtkCalendar)

	Gtk_cell_editable_get_type func() GType

	Gtk_cell_editable_start_editing func(
		cell_editable *GtkCellEditable,
		event *GdkEvent)

	Gtk_cell_editable_editing_done func(
		cell_editable *GtkCellEditable)

	Gtk_cell_editable_remove_widget func(
		cell_editable *GtkCellEditable)

	Gtk_cell_renderer_get_type func() GType

	Gtk_cell_renderer_get_size func(
		cell *GtkCellRenderer,
		widget *GtkWidget,
		cell_area *GdkRectangle,
		x_offset *Gint,
		y_offset *Gint,
		width *Gint,
		height *Gint)

	Gtk_cell_renderer_render func(
		cell *GtkCellRenderer,
		window *GdkWindow,
		widget *GtkWidget,
		background_area *GdkRectangle,
		cell_area *GdkRectangle,
		expose_area *GdkRectangle,
		flags GtkCellRendererState)

	Gtk_cell_renderer_activate func(
		cell *GtkCellRenderer,
		event *GdkEvent,
		widget *GtkWidget,
		path string,
		background_area *GdkRectangle,
		cell_area *GdkRectangle,
		flags GtkCellRendererState) Gboolean

	Gtk_cell_renderer_start_editing func(
		cell *GtkCellRenderer,
		event *GdkEvent,
		widget *GtkWidget,
		path string,
		background_area *GdkRectangle,
		cell_area *GdkRectangle,
		flags GtkCellRendererState) *GtkCellEditable

	Gtk_cell_renderer_set_fixed_size func(
		cell *GtkCellRenderer,
		width Gint,
		height Gint)

	Gtk_cell_renderer_get_fixed_size func(
		cell *GtkCellRenderer,
		width *Gint,
		height *Gint)

	Gtk_cell_renderer_set_alignment func(
		cell *GtkCellRenderer,
		xalign Gfloat,
		yalign Gfloat)

	Gtk_cell_renderer_get_alignment func(
		cell *GtkCellRenderer,
		xalign *Gfloat,
		yalign *Gfloat)

	Gtk_cell_renderer_set_padding func(
		cell *GtkCellRenderer,
		xpad Gint,
		ypad Gint)

	Gtk_cell_renderer_get_padding func(
		cell *GtkCellRenderer,
		xpad *Gint,
		ypad *Gint)

	Gtk_cell_renderer_set_visible func(
		cell *GtkCellRenderer,
		visible Gboolean)

	Gtk_cell_renderer_get_visible func(
		cell *GtkCellRenderer) Gboolean

	Gtk_cell_renderer_set_sensitive func(
		cell *GtkCellRenderer,
		sensitive Gboolean)

	Gtk_cell_renderer_get_sensitive func(
		cell *GtkCellRenderer) Gboolean

	Gtk_cell_renderer_editing_canceled func(
		cell *GtkCellRenderer)

	Gtk_cell_renderer_stop_editing func(
		cell *GtkCellRenderer,
		canceled Gboolean)

	Gtk_tree_path_new func() *GtkTreePath

	Gtk_tree_path_new_from_string func(
		path string) *GtkTreePath

	Gtk_tree_path_new_from_indices func(first_index Gint,
		v ...VArg) *GtkTreePath

	Gtk_tree_path_to_string func(
		path *GtkTreePath) string

	Gtk_tree_path_new_first func() *GtkTreePath

	Gtk_tree_path_append_index func(
		path *GtkTreePath,
		index_ Gint)

	Gtk_tree_path_prepend_index func(
		path *GtkTreePath,
		index_ Gint)

	Gtk_tree_path_get_depth func(
		path *GtkTreePath) Gint

	Gtk_tree_path_get_indices func(
		path *GtkTreePath) *Gint

	Gtk_tree_path_get_indices_with_depth func(
		path *GtkTreePath,
		depth *Gint) *Gint

	Gtk_tree_path_free func(
		path *GtkTreePath)

	Gtk_tree_path_copy func(
		path *GtkTreePath) *GtkTreePath

	Gtk_tree_path_get_type func() GType

	Gtk_tree_path_compare func(
		a *GtkTreePath,
		b *GtkTreePath) Gint

	Gtk_tree_path_next func(
		path *GtkTreePath)

	Gtk_tree_path_prev func(
		path *GtkTreePath) Gboolean

	Gtk_tree_path_up func(
		path *GtkTreePath) Gboolean

	Gtk_tree_path_down func(
		path *GtkTreePath)

	Gtk_tree_path_is_ancestor func(
		path *GtkTreePath,
		descendant *GtkTreePath) Gboolean

	Gtk_tree_path_is_descendant func(
		path *GtkTreePath,
		ancestor *GtkTreePath) Gboolean

	Gtk_tree_row_reference_get_type func() GType

	Gtk_tree_row_reference_new func(
		model *GtkTreeModel,
		path *GtkTreePath) *GtkTreeRowReference

	Gtk_tree_row_reference_new_proxy func(
		proxy *GObject,
		model *GtkTreeModel,
		path *GtkTreePath) *GtkTreeRowReference

	Gtk_tree_row_reference_get_path func(
		reference *GtkTreeRowReference) *GtkTreePath

	Gtk_tree_row_reference_get_model func(
		reference *GtkTreeRowReference) *GtkTreeModel

	Gtk_tree_row_reference_valid func(
		reference *GtkTreeRowReference) Gboolean

	Gtk_tree_row_reference_copy func(
		reference *GtkTreeRowReference) *GtkTreeRowReference

	Gtk_tree_row_reference_free func(
		reference *GtkTreeRowReference)

	Gtk_tree_row_reference_inserted func(
		proxy *GObject,
		path *GtkTreePath)

	Gtk_tree_row_reference_deleted func(
		proxy *GObject,
		path *GtkTreePath)

	Gtk_tree_row_reference_reordered func(
		proxy *GObject,
		path *GtkTreePath,
		iter *GtkTreeIter,
		new_order *Gint)

	Gtk_tree_iter_copy func(
		iter *GtkTreeIter) *GtkTreeIter

	Gtk_tree_iter_free func(
		iter *GtkTreeIter)

	Gtk_tree_iter_get_type func() GType

	Gtk_tree_model_get_type func() GType

	Gtk_tree_model_get_flags func(
		tree_model *GtkTreeModel) GtkTreeModelFlags

	Gtk_tree_model_get_n_columns func(
		tree_model *GtkTreeModel) Gint

	Gtk_tree_model_get_column_type func(
		tree_model *GtkTreeModel,
		index_ Gint) GType

	Gtk_tree_model_get_iter func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter,
		path *GtkTreePath) Gboolean

	Gtk_tree_model_get_iter_from_string func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter,
		path_string string) Gboolean

	Gtk_tree_model_get_string_from_iter func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter) string

	Gtk_tree_model_get_iter_first func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_model_get_path func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter) *GtkTreePath

	Gtk_tree_model_get_value func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter,
		column Gint,
		value *GValue)

	Gtk_tree_model_iter_next func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_model_iter_children func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter,
		parent *GtkTreeIter) Gboolean

	Gtk_tree_model_iter_has_child func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_model_iter_n_children func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter) Gint

	Gtk_tree_model_iter_nth_child func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter,
		parent *GtkTreeIter,
		n Gint) Gboolean

	Gtk_tree_model_iter_parent func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter,
		child *GtkTreeIter) Gboolean

	Gtk_tree_model_ref_node func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter)

	Gtk_tree_model_unref_node func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter)

	Gtk_tree_model_get func(tree_model *GtkTreeModel,
		iter *GtkTreeIter, v ...VArg)

	Gtk_tree_model_get_valist func(
		tree_model *GtkTreeModel,
		iter *GtkTreeIter,
		var_args Va_list)

	Gtk_tree_model_foreach func(
		model *GtkTreeModel,
		f GtkTreeModelForeachFunc,
		user_data Gpointer)

	Gtk_tree_model_row_changed func(
		tree_model *GtkTreeModel,
		path *GtkTreePath,
		iter *GtkTreeIter)

	Gtk_tree_model_row_inserted func(
		tree_model *GtkTreeModel,
		path *GtkTreePath,
		iter *GtkTreeIter)

	Gtk_tree_model_row_has_child_toggled func(
		tree_model *GtkTreeModel,
		path *GtkTreePath,
		iter *GtkTreeIter)

	Gtk_tree_model_row_deleted func(
		tree_model *GtkTreeModel,
		path *GtkTreePath)

	Gtk_tree_model_rows_reordered func(
		tree_model *GtkTreeModel,
		path *GtkTreePath,
		iter *GtkTreeIter,
		new_order *Gint)

	Gtk_tree_sortable_get_type func() GType

	Gtk_tree_sortable_sort_column_changed func(
		sortable *GtkTreeSortable)

	Gtk_tree_sortable_get_sort_column_id func(
		sortable *GtkTreeSortable,
		sort_column_id *Gint,
		order *GtkSortType) Gboolean

	Gtk_tree_sortable_set_sort_column_id func(
		sortable *GtkTreeSortable,
		sort_column_id Gint,
		order GtkSortType)

	Gtk_tree_sortable_set_sort_func func(
		sortable *GtkTreeSortable,
		sort_column_id Gint,
		sort_func GtkTreeIterCompareFunc,
		user_dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_sortable_set_default_sort_func func(
		sortable *GtkTreeSortable,
		sort_func GtkTreeIterCompareFunc,
		user_dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_sortable_has_default_sort_func func(
		sortable *GtkTreeSortable) Gboolean

	Gtk_tree_view_column_get_type func() GType

	Gtk_tree_view_column_new func() *GtkTreeViewColumn

	Gtk_tree_view_column_new_with_attributes func(
		title string, cell *GtkCellRenderer,
		v ...VArg) *GtkTreeViewColumn

	Gtk_tree_view_column_pack_start func(
		tree_column *GtkTreeViewColumn,
		cell *GtkCellRenderer,
		expand Gboolean)

	Gtk_tree_view_column_pack_end func(
		tree_column *GtkTreeViewColumn,
		cell *GtkCellRenderer,
		expand Gboolean)

	Gtk_tree_view_column_clear func(
		tree_column *GtkTreeViewColumn)

	Gtk_tree_view_column_get_cell_renderers func(
		tree_column *GtkTreeViewColumn) *GList

	Gtk_tree_view_column_add_attribute func(
		tree_column *GtkTreeViewColumn,
		cell_renderer *GtkCellRenderer,
		attribute string,
		column Gint)

	Gtk_tree_view_column_set_attributes func(
		tree_column *GtkTreeViewColumn,
		cell_renderer *GtkCellRenderer, v ...VArg)

	Gtk_tree_view_column_set_cell_data_func func(
		tree_column *GtkTreeViewColumn,
		cell_renderer *GtkCellRenderer,
		f GtkTreeCellDataFunc,
		func_dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_view_column_clear_attributes func(
		tree_column *GtkTreeViewColumn,
		cell_renderer *GtkCellRenderer)

	Gtk_tree_view_column_set_spacing func(
		tree_column *GtkTreeViewColumn,
		spacing Gint)

	Gtk_tree_view_column_get_spacing func(
		tree_column *GtkTreeViewColumn) Gint

	Gtk_tree_view_column_set_visible func(
		tree_column *GtkTreeViewColumn,
		visible Gboolean)

	Gtk_tree_view_column_get_visible func(
		tree_column *GtkTreeViewColumn) Gboolean

	Gtk_tree_view_column_set_resizable func(
		tree_column *GtkTreeViewColumn,
		resizable Gboolean)

	Gtk_tree_view_column_get_resizable func(
		tree_column *GtkTreeViewColumn) Gboolean

	Gtk_tree_view_column_set_sizing func(
		tree_column *GtkTreeViewColumn,
		typ GtkTreeViewColumnSizing)

	Gtk_tree_view_column_get_sizing func(
		tree_column *GtkTreeViewColumn) GtkTreeViewColumnSizing

	Gtk_tree_view_column_get_width func(
		tree_column *GtkTreeViewColumn) Gint

	Gtk_tree_view_column_get_fixed_width func(
		tree_column *GtkTreeViewColumn) Gint

	Gtk_tree_view_column_set_fixed_width func(
		tree_column *GtkTreeViewColumn,
		fixed_width Gint)

	Gtk_tree_view_column_set_min_width func(
		tree_column *GtkTreeViewColumn,
		min_width Gint)

	Gtk_tree_view_column_get_min_width func(
		tree_column *GtkTreeViewColumn) Gint

	Gtk_tree_view_column_set_max_width func(
		tree_column *GtkTreeViewColumn,
		max_width Gint)

	Gtk_tree_view_column_get_max_width func(
		tree_column *GtkTreeViewColumn) Gint

	Gtk_tree_view_column_clicked func(
		tree_column *GtkTreeViewColumn)

	Gtk_tree_view_column_set_title func(
		tree_column *GtkTreeViewColumn,
		title string)

	Gtk_tree_view_column_get_title func(
		tree_column *GtkTreeViewColumn) string

	Gtk_tree_view_column_set_expand func(
		tree_column *GtkTreeViewColumn,
		expand Gboolean)

	Gtk_tree_view_column_get_expand func(
		tree_column *GtkTreeViewColumn) Gboolean

	Gtk_tree_view_column_set_clickable func(
		tree_column *GtkTreeViewColumn,
		clickable Gboolean)

	Gtk_tree_view_column_get_clickable func(
		tree_column *GtkTreeViewColumn) Gboolean

	Gtk_tree_view_column_set_widget func(
		tree_column *GtkTreeViewColumn,
		widget *GtkWidget)

	Gtk_tree_view_column_get_widget func(
		tree_column *GtkTreeViewColumn) *GtkWidget

	Gtk_tree_view_column_set_alignment func(
		tree_column *GtkTreeViewColumn,
		xalign Gfloat)

	Gtk_tree_view_column_get_alignment func(
		tree_column *GtkTreeViewColumn) Gfloat

	Gtk_tree_view_column_set_reorderable func(
		tree_column *GtkTreeViewColumn,
		reorderable Gboolean)

	Gtk_tree_view_column_get_reorderable func(
		tree_column *GtkTreeViewColumn) Gboolean

	Gtk_tree_view_column_set_sort_column_id func(
		tree_column *GtkTreeViewColumn,
		sort_column_id Gint)

	Gtk_tree_view_column_get_sort_column_id func(
		tree_column *GtkTreeViewColumn) Gint

	Gtk_tree_view_column_set_sort_indicator func(
		tree_column *GtkTreeViewColumn,
		setting Gboolean)

	Gtk_tree_view_column_get_sort_indicator func(
		tree_column *GtkTreeViewColumn) Gboolean

	Gtk_tree_view_column_set_sort_order func(
		tree_column *GtkTreeViewColumn,
		order GtkSortType)

	Gtk_tree_view_column_get_sort_order func(
		tree_column *GtkTreeViewColumn) GtkSortType

	Gtk_tree_view_column_cell_set_cell_data func(
		tree_column *GtkTreeViewColumn,
		tree_model *GtkTreeModel,
		iter *GtkTreeIter,
		is_expander Gboolean,
		is_expanded Gboolean)

	Gtk_tree_view_column_cell_get_size func(
		tree_column *GtkTreeViewColumn,
		cell_area *GdkRectangle,
		x_offset *Gint,
		y_offset *Gint,
		width *Gint,
		height *Gint)

	Gtk_tree_view_column_cell_is_visible func(
		tree_column *GtkTreeViewColumn) Gboolean

	Gtk_tree_view_column_focus_cell func(
		tree_column *GtkTreeViewColumn,
		cell *GtkCellRenderer)

	Gtk_tree_view_column_cell_get_position func(
		tree_column *GtkTreeViewColumn,
		cell_renderer *GtkCellRenderer,
		start_pos *Gint,
		width *Gint) Gboolean

	Gtk_tree_view_column_queue_resize func(
		tree_column *GtkTreeViewColumn)

	Gtk_tree_view_column_get_tree_view func(
		tree_column *GtkTreeViewColumn) *GtkWidget

	Gtk_cell_layout_get_type func() GType

	Gtk_cell_layout_pack_start func(
		cell_layout *GtkCellLayout,
		cell *GtkCellRenderer,
		expand Gboolean)

	Gtk_cell_layout_pack_end func(
		cell_layout *GtkCellLayout,
		cell *GtkCellRenderer,
		expand Gboolean)

	Gtk_cell_layout_get_cells func(
		cell_layout *GtkCellLayout) *GList

	Gtk_cell_layout_clear func(
		cell_layout *GtkCellLayout)

	Gtk_cell_layout_set_attributes func(cell_layout *GtkCellLayout,
		cell *GtkCellRenderer, v ...VArg)

	Gtk_cell_layout_add_attribute func(
		cell_layout *GtkCellLayout,
		cell *GtkCellRenderer,
		attribute string,
		column Gint)

	Gtk_cell_layout_set_cell_data_func func(
		cell_layout *GtkCellLayout,
		cell *GtkCellRenderer,
		f GtkCellLayoutDataFunc,
		func_dataGpointer,
		destroy GDestroyNotify)

	Gtk_cell_layout_clear_attributes func(
		cell_layout *GtkCellLayout,
		cell *GtkCellRenderer)

	Gtk_cell_layout_reorder func(
		cell_layout *GtkCellLayout,
		cell *GtkCellRenderer,
		position Gint)

	Gtk_cell_renderer_text_get_type func() GType

	Gtk_cell_renderer_text_new func() *GtkCellRenderer

	Gtk_cell_renderer_text_set_fixed_height_from_font func(
		renderer *GtkCellRendererText,
		number_of_rows Gint)

	Gtk_cell_renderer_accel_get_type func() GType

	Gtk_cell_renderer_accel_new func() *GtkCellRenderer

	Gtk_cell_renderer_combo_get_type func() GType

	Gtk_cell_renderer_combo_new func() *GtkCellRenderer

	Gtk_cell_renderer_pixbuf_get_type func() GType

	Gtk_cell_renderer_pixbuf_new func() *GtkCellRenderer

	Gtk_cell_renderer_progress_get_type func() GType

	Gtk_cell_renderer_progress_new func() *GtkCellRenderer

	Gtk_cell_renderer_spin_get_type func() GType

	Gtk_cell_renderer_spin_new func() *GtkCellRenderer

	Gtk_cell_renderer_spinner_get_type func() GType

	Gtk_cell_renderer_spinner_new func() *GtkCellRenderer

	Gtk_cell_renderer_toggle_get_type func() GType

	Gtk_cell_renderer_toggle_new func() *GtkCellRenderer

	Gtk_cell_renderer_toggle_get_radio func(
		toggle *GtkCellRendererToggle) Gboolean

	Gtk_cell_renderer_toggle_set_radio func(
		toggle *GtkCellRendererToggle,
		radio Gboolean)

	Gtk_cell_renderer_toggle_get_active func(
		toggle *GtkCellRendererToggle) Gboolean

	Gtk_cell_renderer_toggle_set_active func(
		toggle *GtkCellRendererToggle,
		setting Gboolean)

	Gtk_cell_renderer_toggle_get_activatable func(
		toggle *GtkCellRendererToggle) Gboolean

	Gtk_cell_renderer_toggle_set_activatable func(
		toggle *GtkCellRendererToggle,
		setting Gboolean)

	Gtk_cell_view_get_type func() GType

	Gtk_cell_view_new func() *GtkWidget

	Gtk_cell_view_new_with_text func(
		text string) *GtkWidget

	Gtk_cell_view_new_with_markup func(
		markup string) *GtkWidget

	Gtk_cell_view_new_with_pixbuf func(
		pixbuf *GdkPixbuf) *GtkWidget

	Gtk_cell_view_set_model func(
		cell_view *GtkCellView,
		model *GtkTreeModel)

	Gtk_cell_view_get_model func(
		cell_view *GtkCellView) *GtkTreeModel

	Gtk_cell_view_set_displayed_row func(
		cell_view *GtkCellView,
		path *GtkTreePath)

	Gtk_cell_view_get_displayed_row func(
		cell_view *GtkCellView) *GtkTreePath

	Gtk_cell_view_get_size_of_row func(
		cell_view *GtkCellView,
		path *GtkTreePath,
		requisition *GtkRequisition) Gboolean

	Gtk_cell_view_set_background_color func(
		cell_view *GtkCellView,
		color *GdkColor)

	Gtk_cell_view_get_cell_renderers func(
		cell_view *GtkCellView) *GList

	Gtk_toggle_button_get_type func() GType

	Gtk_toggle_button_new func() *GtkWidget

	Gtk_toggle_button_new_with_label func(
		label string) *GtkWidget

	Gtk_toggle_button_new_with_mnemonic func(
		label string) *GtkWidget

	Gtk_toggle_button_set_mode func(
		toggle_button *GtkToggleButton,
		draw_indicator Gboolean)

	Gtk_toggle_button_get_mode func(
		toggle_button *GtkToggleButton) Gboolean

	Gtk_toggle_button_set_active func(
		toggle_button *GtkToggleButton,
		is_active Gboolean)

	Gtk_toggle_button_get_active func(
		toggle_button *GtkToggleButton) Gboolean

	Gtk_toggle_button_toggled func(
		toggle_button *GtkToggleButton)

	Gtk_toggle_button_set_inconsistent func(
		toggle_button *GtkToggleButton,
		setting Gboolean)

	Gtk_toggle_button_get_inconsistent func(
		toggle_button *GtkToggleButton) Gboolean

	Gtk_check_button_get_type func() GType

	Gtk_check_button_new func() *GtkWidget

	Gtk_check_button_new_with_label func(
		label string) *GtkWidget

	Gtk_check_button_new_with_mnemonic func(
		label string) *GtkWidget

	Gtk_item_get_type func() GType

	Gtk_item_select func(
		item *GtkItem)

	Gtk_item_deselect func(
		item *GtkItem)

	Gtk_item_toggle func(
		item *GtkItem)

	Gtk_menu_item_get_type func() GType

	Gtk_menu_item_new func() *GtkWidget

	Gtk_menu_item_new_with_label func(
		label string) *GtkWidget

	Gtk_menu_item_new_with_mnemonic func(
		label string) *GtkWidget

	Gtk_menu_item_set_submenu func(
		menu_item *GtkMenuItem,
		submenu *GtkWidget)

	Gtk_menu_item_get_submenu func(
		menu_item *GtkMenuItem) *GtkWidget

	Gtk_menu_item_select func(
		menu_item *GtkMenuItem)

	Gtk_menu_item_deselect func(
		menu_item *GtkMenuItem)

	Gtk_menu_item_activate func(
		menu_item *GtkMenuItem)

	Gtk_menu_item_toggle_size_request func(
		menu_item *GtkMenuItem,
		requisition *Gint)

	Gtk_menu_item_toggle_size_allocate func(
		menu_item *GtkMenuItem,
		allocation Gint)

	Gtk_menu_item_set_right_justified func(
		menu_item *GtkMenuItem,
		right_justified Gboolean)

	Gtk_menu_item_get_right_justified func(
		menu_item *GtkMenuItem) Gboolean

	Gtk_menu_item_set_accel_path func(
		menu_item *GtkMenuItem,
		accel_path string)

	Gtk_menu_item_get_accel_path func(
		menu_item *GtkMenuItem) string

	Gtk_menu_item_set_label func(
		menu_item *GtkMenuItem,
		label string)

	Gtk_menu_item_get_label func(
		menu_item *GtkMenuItem) string

	Gtk_menu_item_set_use_underline func(
		menu_item *GtkMenuItem,
		setting Gboolean)

	Gtk_menu_item_get_use_underline func(
		menu_item *GtkMenuItem) Gboolean

	Gtk_menu_item_remove_submenu func(
		menu_item *GtkMenuItem)

	Gtk_check_menu_item_get_type func() GType

	Gtk_check_menu_item_new func() *GtkWidget

	Gtk_check_menu_item_new_with_label func(
		label string) *GtkWidget

	Gtk_check_menu_item_new_with_mnemonic func(
		label string) *GtkWidget

	Gtk_check_menu_item_set_active func(
		check_menu_item *GtkCheckMenuItem,
		is_active Gboolean)

	Gtk_check_menu_item_get_active func(
		check_menu_item *GtkCheckMenuItem) Gboolean

	Gtk_check_menu_item_toggled func(
		check_menu_item *GtkCheckMenuItem)

	Gtk_check_menu_item_set_inconsistent func(
		check_menu_item *GtkCheckMenuItem,
		setting Gboolean)

	Gtk_check_menu_item_get_inconsistent func(
		check_menu_item *GtkCheckMenuItem) Gboolean

	Gtk_check_menu_item_set_draw_as_radio func(
		check_menu_item *GtkCheckMenuItem,
		draw_as_radio Gboolean)

	Gtk_check_menu_item_get_draw_as_radio func(
		check_menu_item *GtkCheckMenuItem) Gboolean

	Gtk_check_menu_item_set_show_toggle func(
		menu_item *GtkCheckMenuItem,
		always Gboolean)

	Gtk_text_tag_get_type func() GType

	Gtk_text_tag_new func(
		name string) *GtkTextTag

	Gtk_text_tag_get_priority func(
		tag *GtkTextTag) Gint

	Gtk_text_tag_set_priority func(
		tag *GtkTextTag,
		priority Gint)

	Gtk_text_tag_event func(
		tag *GtkTextTag,
		event_object *GObject,
		event *GdkEvent,
		iter *GtkTextIter) Gboolean

	Gtk_text_attributes_new func() *GtkTextAttributes

	Gtk_text_attributes_copy func(
		src *GtkTextAttributes) *GtkTextAttributes

	Gtk_text_attributes_copy_values func(
		src *GtkTextAttributes,
		dest *GtkTextAttributes)

	Gtk_text_attributes_unref func(
		values *GtkTextAttributes)

	Gtk_text_attributes_ref func(
		values *GtkTextAttributes) *GtkTextAttributes

	Gtk_text_attributes_get_type func() GType

	Gtk_text_child_anchor_get_type func() GType

	Gtk_text_child_anchor_new func() *GtkTextChildAnchor

	Gtk_text_child_anchor_get_widgets func(
		anchor *GtkTextChildAnchor) *GList

	Gtk_text_child_anchor_get_deleted func(
		anchor *GtkTextChildAnchor) Gboolean

	Gtk_text_iter_get_buffer func(
		iter *GtkTextIter) *GtkTextBuffer

	Gtk_text_iter_copy func(
		iter *GtkTextIter) *GtkTextIter

	Gtk_text_iter_free func(
		iter *GtkTextIter)

	Gtk_text_iter_get_type func() GType

	Gtk_text_iter_get_offset func(
		iter *GtkTextIter) Gint

	Gtk_text_iter_get_line func(
		iter *GtkTextIter) Gint

	Gtk_text_iter_get_line_offset func(
		iter *GtkTextIter) Gint

	Gtk_text_iter_get_line_index func(
		iter *GtkTextIter) Gint

	Gtk_text_iter_get_visible_line_offset func(
		iter *GtkTextIter) Gint

	Gtk_text_iter_get_visible_line_index func(
		iter *GtkTextIter) Gint

	Gtk_text_iter_get_char func(
		iter *GtkTextIter) Gunichar

	Gtk_text_iter_get_slice func(
		start *GtkTextIter,
		end *GtkTextIter) string

	Gtk_text_iter_get_text func(
		start *GtkTextIter,
		end *GtkTextIter) string

	Gtk_text_iter_get_visible_slice func(
		start *GtkTextIter,
		end *GtkTextIter) string

	Gtk_text_iter_get_visible_text func(
		start *GtkTextIter,
		end *GtkTextIter) string

	Gtk_text_iter_get_pixbuf func(
		iter *GtkTextIter) *GdkPixbuf

	Gtk_text_iter_get_marks func(
		iter *GtkTextIter) *GSList

	Gtk_text_iter_get_child_anchor func(
		iter *GtkTextIter) *GtkTextChildAnchor

	Gtk_text_iter_get_toggled_tags func(
		iter *GtkTextIter,
		toggled_on Gboolean) *GSList

	Gtk_text_iter_begins_tag func(
		iter *GtkTextIter,
		tag *GtkTextTag) Gboolean

	Gtk_text_iter_ends_tag func(
		iter *GtkTextIter,
		tag *GtkTextTag) Gboolean

	Gtk_text_iter_toggles_tag func(
		iter *GtkTextIter,
		tag *GtkTextTag) Gboolean

	Gtk_text_iter_has_tag func(
		iter *GtkTextIter,
		tag *GtkTextTag) Gboolean

	Gtk_text_iter_get_tags func(
		iter *GtkTextIter) *GSList

	Gtk_text_iter_editable func(
		iter *GtkTextIter,
		default_setting Gboolean) Gboolean

	Gtk_text_iter_can_insert func(
		iter *GtkTextIter,
		default_editability Gboolean) Gboolean

	Gtk_text_iter_starts_word func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_ends_word func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_inside_word func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_starts_sentence func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_ends_sentence func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_inside_sentence func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_starts_line func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_ends_line func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_is_cursor_position func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_get_chars_in_line func(
		iter *GtkTextIter) Gint

	Gtk_text_iter_get_bytes_in_line func(
		iter *GtkTextIter) Gint

	Gtk_text_iter_get_attributes func(
		iter *GtkTextIter,
		values *GtkTextAttributes) Gboolean

	Gtk_text_iter_get_language func(
		iter *GtkTextIter) *PangoLanguage

	Gtk_text_iter_is_end func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_is_start func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_char func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_backward_char func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_chars func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_backward_chars func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_forward_line func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_backward_line func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_lines func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_backward_lines func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_forward_word_end func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_backward_word_start func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_word_ends func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_backward_word_starts func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_forward_visible_line func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_backward_visible_line func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_visible_lines func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_backward_visible_lines func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_forward_visible_word_end func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_backward_visible_word_start func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_visible_word_ends func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_backward_visible_word_starts func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_forward_sentence_end func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_backward_sentence_start func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_sentence_ends func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_backward_sentence_starts func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_forward_cursor_position func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_backward_cursor_position func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_cursor_positions func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_backward_cursor_positions func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_forward_visible_cursor_position func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_backward_visible_cursor_position func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_forward_visible_cursor_positions func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_backward_visible_cursor_positions func(
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_iter_set_offset func(
		iter *GtkTextIter,
		char_offset Gint)

	Gtk_text_iter_set_line func(
		iter *GtkTextIter,
		line_number Gint)

	Gtk_text_iter_set_line_offset func(
		iter *GtkTextIter,
		char_on_line Gint)

	Gtk_text_iter_set_line_index func(
		iter *GtkTextIter,
		byte_on_line Gint)

	Gtk_text_iter_forward_to_end func(
		iter *GtkTextIter)

	Gtk_text_iter_forward_to_line_end func(
		iter *GtkTextIter) Gboolean

	Gtk_text_iter_set_visible_line_offset func(
		iter *GtkTextIter,
		char_on_line Gint)

	Gtk_text_iter_set_visible_line_index func(
		iter *GtkTextIter,
		byte_on_line Gint)

	Gtk_text_iter_forward_to_tag_toggle func(
		iter *GtkTextIter,
		tag *GtkTextTag) Gboolean

	Gtk_text_iter_backward_to_tag_toggle func(
		iter *GtkTextIter,
		tag *GtkTextTag) Gboolean

	Gtk_text_iter_forward_find_char func(
		iter *GtkTextIter,
		pred GtkTextCharPredicate,
		user_dataGpointer,
		limit *GtkTextIter) Gboolean

	Gtk_text_iter_backward_find_char func(
		iter *GtkTextIter,
		pred GtkTextCharPredicate,
		user_dataGpointer,
		limit *GtkTextIter) Gboolean

	Gtk_text_iter_forward_search func(
		iter *GtkTextIter,
		str string,
		flags GtkTextSearchFlags,
		match_start *GtkTextIter,
		match_end *GtkTextIter,
		limit *GtkTextIter) Gboolean

	Gtk_text_iter_backward_search func(
		iter *GtkTextIter,
		str string,
		flags GtkTextSearchFlags,
		match_start *GtkTextIter,
		match_end *GtkTextIter,
		limit *GtkTextIter) Gboolean

	Gtk_text_iter_equal func(
		lhs *GtkTextIter,
		rhs *GtkTextIter) Gboolean

	Gtk_text_iter_compare func(
		lhs *GtkTextIter,
		rhs *GtkTextIter) Gint

	Gtk_text_iter_in_range func(
		iter *GtkTextIter,
		start *GtkTextIter,
		end *GtkTextIter) Gboolean

	Gtk_text_iter_order func(
		first *GtkTextIter,
		second *GtkTextIter)

	Gtk_target_list_new func(
		targets *GtkTargetEntry,
		ntargets Guint) *GtkTargetList

	Gtk_target_list_ref func(
		list *GtkTargetList) *GtkTargetList

	Gtk_target_list_unref func(
		list *GtkTargetList)

	Gtk_target_list_add func(
		list *GtkTargetList,
		target GdkAtom,
		flags Guint,
		info Guint)

	Gtk_target_list_add_text_targets func(
		list *GtkTargetList,
		info Guint)

	Gtk_target_list_add_rich_text_targets func(
		list *GtkTargetList,
		info Guint,
		deserializable Gboolean,
		buffer *GtkTextBuffer)

	Gtk_target_list_add_image_targets func(
		list *GtkTargetList,
		info Guint,
		writable Gboolean)

	Gtk_target_list_add_uri_targets func(
		list *GtkTargetList,
		info Guint)

	Gtk_target_list_add_table func(
		list *GtkTargetList,
		targets *GtkTargetEntry,
		ntargets Guint)

	Gtk_target_list_remove func(
		list *GtkTargetList,
		target GdkAtom)

	Gtk_target_list_find func(
		list *GtkTargetList,
		target GdkAtom,
		info *Guint) Gboolean

	Gtk_target_table_new_from_list func(
		list *GtkTargetList,
		n_targets *Gint) *GtkTargetEntry

	Gtk_target_table_free func(
		targets *GtkTargetEntry,
		n_targets Gint)

	Gtk_selection_owner_set func(
		widget *GtkWidget,
		selection GdkAtom,
		time_ Guint32) Gboolean

	Gtk_selection_owner_set_for_display func(
		display *GdkDisplay,
		widget *GtkWidget,
		selection GdkAtom,
		time_ Guint32) Gboolean

	Gtk_selection_add_target func(
		widget *GtkWidget,
		selection GdkAtom,
		target GdkAtom,
		info Guint)

	Gtk_selection_add_targets func(
		widget *GtkWidget,
		selection GdkAtom,
		targets *GtkTargetEntry,
		ntargets Guint)

	Gtk_selection_clear_targets func(
		widget *GtkWidget,
		selection GdkAtom)

	Gtk_selection_convert func(
		widget *GtkWidget,
		selection GdkAtom,
		target GdkAtom,
		time_ Guint32) Gboolean

	Gtk_selection_data_get_selection func(
		selection_data *GtkSelectionData) GdkAtom

	Gtk_selection_data_get_target func(
		selection_data *GtkSelectionData) GdkAtom

	Gtk_selection_data_get_data_type func(
		selection_data *GtkSelectionData) GdkAtom

	Gtk_selection_data_get_format func(
		selection_data *GtkSelectionData) Gint

	Gtk_selection_data_get_data func(
		selection_data *GtkSelectionData) *Guchar

	Gtk_selection_data_get_length func(
		selection_data *GtkSelectionData) Gint

	Gtk_selection_data_get_display func(
		selection_data *GtkSelectionData) *GdkDisplay

	Gtk_selection_data_set func(
		selection_data *GtkSelectionData,
		typ GdkAtom,
		format Gint,
		data *Guchar,
		length Gint)

	Gtk_selection_data_set_text func(
		selection_data *GtkSelectionData,
		str string,
		len Gint) Gboolean

	Gtk_selection_data_get_text func(
		selection_data *GtkSelectionData) *Guchar

	Gtk_selection_data_set_pixbuf func(
		selection_data *GtkSelectionData,
		pixbuf *GdkPixbuf) Gboolean

	Gtk_selection_data_get_pixbuf func(
		selection_data *GtkSelectionData) *GdkPixbuf

	Gtk_selection_data_set_uris func(
		selection_data *GtkSelectionData,
		uris **Gchar) Gboolean

	Gtk_selection_data_get_uris func(
		selection_data *GtkSelectionData) **Gchar

	Gtk_selection_data_get_targets func(
		selection_data *GtkSelectionData,
		targets **GdkAtom,
		n_atoms *Gint) Gboolean

	Gtk_selection_data_targets_include_text func(
		selection_data *GtkSelectionData) Gboolean

	Gtk_selection_data_targets_include_rich_text func(
		selection_data *GtkSelectionData,
		buffer *GtkTextBuffer) Gboolean

	Gtk_selection_data_targets_include_image func(
		selection_data *GtkSelectionData,
		writable Gboolean) Gboolean

	Gtk_selection_data_targets_include_uri func(
		selection_data *GtkSelectionData) Gboolean

	Gtk_targets_include_text func(
		targets *GdkAtom,
		n_targets Gint) Gboolean

	Gtk_targets_include_rich_text func(
		targets *GdkAtom,
		n_targets Gint,
		buffer *GtkTextBuffer) Gboolean

	Gtk_targets_include_image func(
		targets *GdkAtom,
		n_targets Gint,
		writable Gboolean) Gboolean

	Gtk_targets_include_uri func(
		targets *GdkAtom,
		n_targets Gint) Gboolean

	Gtk_selection_remove_all func(
		widget *GtkWidget)

	Gtk_selection_clear func(
		widget *GtkWidget,
		event *GdkEventSelection) Gboolean

	Gtk_selection_data_get_type func() GType

	Gtk_selection_data_copy func(
		data *GtkSelectionData) *GtkSelectionData

	Gtk_selection_data_free func(
		data *GtkSelectionData)

	Gtk_target_list_get_type func() GType

	Gtk_clipboard_get_type func() GType

	Gtk_clipboard_get_for_display func(
		display *GdkDisplay,
		selection GdkAtom) *GtkClipboard

	Gtk_clipboard_get func(
		selection GdkAtom) *GtkClipboard

	Gtk_clipboard_get_display func(
		clipboard *GtkClipboard) *GdkDisplay

	Gtk_clipboard_set_with_data func(
		clipboard *GtkClipboard,
		targets *GtkTargetEntry,
		n_targets Guint,
		get_func GtkClipboardGetFunc,
		clear_func GtkClipboardClearFunc,
		user_data Gpointer) Gboolean

	Gtk_clipboard_set_with_owner func(
		clipboard *GtkClipboard,
		targets *GtkTargetEntry,
		n_targets Guint,
		get_func GtkClipboardGetFunc,
		clear_func GtkClipboardClearFunc,
		owner *GObject) Gboolean

	Gtk_clipboard_get_owner func(
		clipboard *GtkClipboard) *GObject

	Gtk_clipboard_clear func(
		clipboard *GtkClipboard)

	Gtk_clipboard_set_text func(
		clipboard *GtkClipboard,
		text string,
		len Gint)

	Gtk_clipboard_set_image func(
		clipboard *GtkClipboard,
		pixbuf *GdkPixbuf)

	Gtk_clipboard_request_contents func(
		clipboard *GtkClipboard,
		target GdkAtom,
		callback GtkClipboardReceivedFunc,
		user_data Gpointer)

	Gtk_clipboard_request_text func(
		clipboard *GtkClipboard,
		callback GtkClipboardTextReceivedFunc,
		user_data Gpointer)

	Gtk_clipboard_request_rich_text func(
		clipboard *GtkClipboard,
		buffer *GtkTextBuffer,
		callback GtkClipboardRichTextReceivedFunc,
		user_data Gpointer)

	Gtk_clipboard_request_image func(
		clipboard *GtkClipboard,
		callback GtkClipboardImageReceivedFunc,
		user_data Gpointer)

	Gtk_clipboard_request_uris func(
		clipboard *GtkClipboard,
		callback GtkClipboardURIReceivedFunc,
		user_data Gpointer)

	Gtk_clipboard_request_targets func(
		clipboard *GtkClipboard,
		callback GtkClipboardTargetsReceivedFunc,
		user_data Gpointer)

	Gtk_clipboard_wait_for_contents func(
		clipboard *GtkClipboard,
		target GdkAtom) *GtkSelectionData

	Gtk_clipboard_wait_for_text func(
		clipboard *GtkClipboard) string

	Gtk_clipboard_wait_for_rich_text func(
		clipboard *GtkClipboard,
		buffer *GtkTextBuffer,
		format *GdkAtom,
		length *Gsize) *Guint8

	Gtk_clipboard_wait_for_image func(
		clipboard *GtkClipboard) *GdkPixbuf

	Gtk_clipboard_wait_for_uris func(
		clipboard *GtkClipboard) **Gchar

	Gtk_clipboard_wait_for_targets func(
		clipboard *GtkClipboard,
		targets **GdkAtom,
		n_targets *Gint) Gboolean

	Gtk_clipboard_wait_is_text_available func(
		clipboard *GtkClipboard) Gboolean

	Gtk_clipboard_wait_is_rich_text_available func(
		clipboard *GtkClipboard,
		buffer *GtkTextBuffer) Gboolean

	Gtk_clipboard_wait_is_image_available func(
		clipboard *GtkClipboard) Gboolean

	Gtk_clipboard_wait_is_uris_available func(
		clipboard *GtkClipboard) Gboolean

	Gtk_clipboard_wait_is_target_available func(
		clipboard *GtkClipboard,
		target GdkAtom) Gboolean

	Gtk_clipboard_set_can_store func(
		clipboard *GtkClipboard,
		targets *GtkTargetEntry,
		n_targets Gint)

	Gtk_clipboard_store func(
		clipboard *GtkClipboard)

	Gtk_color_button_get_type func() GType

	Gtk_color_button_new func() *GtkWidget

	Gtk_color_button_new_with_color func(
		color *GdkColor) *GtkWidget

	Gtk_color_button_set_color func(
		color_button *GtkColorButton,
		color *GdkColor)

	Gtk_color_button_set_alpha func(
		color_button *GtkColorButton,
		alpha Guint16)

	Gtk_color_button_get_color func(
		color_button *GtkColorButton,
		color *GdkColor)

	Gtk_color_button_get_alpha func(
		color_button *GtkColorButton) Guint16

	Gtk_color_button_set_use_alpha func(
		color_button *GtkColorButton,
		use_alpha Gboolean)

	Gtk_color_button_get_use_alpha func(
		color_button *GtkColorButton) Gboolean

	Gtk_color_button_set_title func(
		color_button *GtkColorButton,
		title string)

	Gtk_color_button_get_title func(
		color_button *GtkColorButton) string

	Gtk_vbox_get_type func() GType

	Gtk_vbox_new func(
		homogeneous Gboolean,
		spacing Gint) *GtkWidget

	Gtk_color_selection_get_type func() GType

	Gtk_color_selection_new func() *GtkWidget

	Gtk_color_selection_get_has_opacity_control func(
		colorsel *GtkColorSelection) Gboolean

	Gtk_color_selection_set_has_opacity_control func(
		colorsel *GtkColorSelection,
		has_opacity Gboolean)

	Gtk_color_selection_get_has_palette func(
		colorsel *GtkColorSelection) Gboolean

	Gtk_color_selection_set_has_palette func(
		colorsel *GtkColorSelection,
		has_palette Gboolean)

	Gtk_color_selection_set_current_color func(
		colorsel *GtkColorSelection,
		color *GdkColor)

	Gtk_color_selection_set_current_alpha func(
		colorsel *GtkColorSelection,
		alpha Guint16)

	Gtk_color_selection_get_current_color func(
		colorsel *GtkColorSelection,
		color *GdkColor)

	Gtk_color_selection_get_current_alpha func(
		colorsel *GtkColorSelection) Guint16

	Gtk_color_selection_set_previous_color func(
		colorsel *GtkColorSelection,
		color *GdkColor)

	Gtk_color_selection_set_previous_alpha func(
		colorsel *GtkColorSelection,
		alpha Guint16)

	Gtk_color_selection_get_previous_color func(
		colorsel *GtkColorSelection,
		color *GdkColor)

	Gtk_color_selection_get_previous_alpha func(
		colorsel *GtkColorSelection) Guint16

	Gtk_color_selection_is_adjusting func(
		colorsel *GtkColorSelection) Gboolean

	Gtk_color_selection_palette_from_string func(
		str string,
		colors **GdkColor,
		n_colors *Gint) Gboolean

	Gtk_color_selection_palette_to_string func(
		colors *GdkColor,
		n_colors Gint) string

	Gtk_color_selection_set_change_palette_hook func(
		f GtkColorSelectionChangePaletteFunc) GtkColorSelectionChangePaletteFunc

	Gtk_color_selection_set_change_palette_with_screen_hook func(
		f GtkColorSelectionChangePaletteWithScreenFunc) GtkColorSelectionChangePaletteWithScreenFunc

	Gtk_color_selection_set_color func(
		colorsel *GtkColorSelection,
		color *Gdouble)

	Gtk_color_selection_get_color func(
		colorsel *GtkColorSelection,
		color *Gdouble)

	Gtk_color_selection_set_update_policy func(
		colorsel *GtkColorSelection,
		policy GtkUpdateType)

	Gtk_color_selection_dialog_get_type func() GType

	Gtk_color_selection_dialog_new func(
		title string) *GtkWidget

	Gtk_color_selection_dialog_get_color_selection func(
		colorsel *GtkColorSelectionDialog) *GtkWidget

	Gtk_drag_get_data func(
		widget *GtkWidget,
		context *GdkDragContext,
		target GdkAtom,
		time_ Guint32)

	Gtk_drag_finish func(
		context *GdkDragContext,
		success Gboolean,
		del Gboolean,
		time_ Guint32)

	Gtk_drag_get_source_widget func(
		context *GdkDragContext) *GtkWidget

	Gtk_drag_highlight func(
		widget *GtkWidget)

	Gtk_drag_unhighlight func(
		widget *GtkWidget)

	Gtk_drag_dest_set func(
		widget *GtkWidget,
		flags GtkDestDefaults,
		targets *GtkTargetEntry,
		n_targets Gint,
		actions GdkDragAction)

	Gtk_drag_dest_set_proxy func(
		widget *GtkWidget,
		proxy_window *GdkWindow,
		protocol GdkDragProtocol,
		use_coordinates Gboolean)

	Gtk_drag_dest_unset func(
		widget *GtkWidget)

	Gtk_drag_dest_find_target func(
		widget *GtkWidget,
		context *GdkDragContext,
		target_list *GtkTargetList) GdkAtom

	Gtk_drag_dest_get_target_list func(
		widget *GtkWidget) *GtkTargetList

	Gtk_drag_dest_set_target_list func(
		widget *GtkWidget,
		target_list *GtkTargetList)

	Gtk_drag_dest_add_text_targets func(
		widget *GtkWidget)

	Gtk_drag_dest_add_image_targets func(
		widget *GtkWidget)

	Gtk_drag_dest_add_uri_targets func(
		widget *GtkWidget)

	Gtk_drag_dest_set_track_motion func(
		widget *GtkWidget,
		track_motion Gboolean)

	Gtk_drag_dest_get_track_motion func(
		widget *GtkWidget) Gboolean

	Gtk_drag_source_set func(
		widget *GtkWidget,
		start_button_mask GdkModifierType,
		targets *GtkTargetEntry,
		n_targets Gint,
		actions GdkDragAction)

	Gtk_drag_source_unset func(
		widget *GtkWidget)

	Gtk_drag_source_get_target_list func(
		widget *GtkWidget) *GtkTargetList

	Gtk_drag_source_set_target_list func(
		widget *GtkWidget,
		target_list *GtkTargetList)

	Gtk_drag_source_add_text_targets func(
		widget *GtkWidget)

	Gtk_drag_source_add_image_targets func(
		widget *GtkWidget)

	Gtk_drag_source_add_uri_targets func(
		widget *GtkWidget)

	Gtk_drag_source_set_icon func(
		widget *GtkWidget,
		colormap *GdkColormap,
		pixmap *GdkPixmap,
		mask *GdkBitmap)

	Gtk_drag_source_set_icon_pixbuf func(
		widget *GtkWidget,
		pixbuf *GdkPixbuf)

	Gtk_drag_source_set_icon_stock func(
		widget *GtkWidget,
		stock_id string)

	Gtk_drag_source_set_icon_name func(
		widget *GtkWidget,
		icon_name string)

	Gtk_drag_begin func(
		widget *GtkWidget,
		targets *GtkTargetList,
		actions GdkDragAction,
		button Gint,
		event *GdkEvent) *GdkDragContext

	Gtk_drag_set_icon_widget func(
		context *GdkDragContext,
		widget *GtkWidget,
		hot_x Gint,
		hot_y Gint)

	Gtk_drag_set_icon_pixmap func(
		context *GdkDragContext,
		colormap *GdkColormap,
		pixmap *GdkPixmap,
		mask *GdkBitmap,
		hot_x Gint,
		hot_y Gint)

	Gtk_drag_set_icon_pixbuf func(
		context *GdkDragContext,
		pixbuf *GdkPixbuf,
		hot_x Gint,
		hot_y Gint)

	Gtk_drag_set_icon_stock func(
		context *GdkDragContext,
		stock_id string,
		hot_x Gint,
		hot_y Gint)

	Gtk_drag_set_icon_name func(
		context *GdkDragContext,
		icon_name string,
		hot_x Gint,
		hot_y Gint)

	Gtk_drag_set_icon_default func(
		context *GdkDragContext)

	Gtk_drag_check_threshold func(
		widget *GtkWidget,
		start_x Gint,
		start_y Gint,
		current_x Gint,
		current_y Gint) Gboolean

	Gtk_drag_set_default_icon func(
		colormap *GdkColormap,
		pixmap *GdkPixmap,
		mask *GdkBitmap,
		hot_x Gint,
		hot_y Gint)

	Gtk_editable_get_type func() GType

	Gtk_editable_select_region func(
		editable *GtkEditable,
		start_pos Gint,
		end_pos Gint)

	Gtk_editable_get_selection_bounds func(
		editable *GtkEditable,
		start_pos *Gint,
		end_pos *Gint) Gboolean

	Gtk_editable_insert_text func(
		editable *GtkEditable,
		new_text string,
		new_text_length Gint,
		position *Gint)

	Gtk_editable_delete_text func(
		editable *GtkEditable,
		start_pos Gint,
		end_pos Gint)

	Gtk_editable_get_chars func(
		editable *GtkEditable,
		start_pos Gint,
		end_pos Gint) string

	Gtk_editable_cut_clipboard func(
		editable *GtkEditable)

	Gtk_editable_copy_clipboard func(
		editable *GtkEditable)

	Gtk_editable_paste_clipboard func(
		editable *GtkEditable)

	Gtk_editable_delete_selection func(
		editable *GtkEditable)

	Gtk_editable_set_position func(
		editable *GtkEditable,
		position Gint)

	Gtk_editable_get_position func(
		editable *GtkEditable) Gint

	Gtk_editable_set_editable func(
		editable *GtkEditable,
		is_editable Gboolean)

	Gtk_editable_get_editable func(
		editable *GtkEditable) Gboolean

	Gtk_im_context_get_type func() GType

	Gtk_im_context_set_client_window func(
		context *GtkIMContext,
		window *GdkWindow)

	Gtk_im_context_get_preedit_string func(
		context *GtkIMContext,
		str **Gchar,
		attrs **PangoAttrList,
		cursor_pos *Gint)

	Gtk_im_context_filter_keypress func(
		context *GtkIMContext,
		event *GdkEventKey) Gboolean

	Gtk_im_context_focus_in func(
		context *GtkIMContext)

	Gtk_im_context_focus_out func(
		context *GtkIMContext)

	Gtk_im_context_reset func(
		context *GtkIMContext)

	Gtk_im_context_set_cursor_location func(
		context *GtkIMContext,
		area *GdkRectangle)

	Gtk_im_context_set_use_preedit func(
		context *GtkIMContext,
		use_preedit Gboolean)

	Gtk_im_context_set_surrounding func(
		context *GtkIMContext,
		text string,
		len Gint,
		cursor_index Gint)

	Gtk_im_context_get_surrounding func(
		context *GtkIMContext,
		text **Gchar,
		cursor_index *Gint) Gboolean

	Gtk_im_context_delete_surrounding func(
		context *GtkIMContext,
		offset Gint,
		n_chars Gint) Gboolean

	Gtk_entry_buffer_get_type func() GType

	Gtk_entry_buffer_new func(
		initial_chars string,
		n_initial_chars Gint) *GtkEntryBuffer

	Gtk_entry_buffer_get_bytes func(
		buffer *GtkEntryBuffer) Gsize

	Gtk_entry_buffer_get_length func(
		buffer *GtkEntryBuffer) Guint

	Gtk_entry_buffer_get_text func(
		buffer *GtkEntryBuffer) string

	Gtk_entry_buffer_set_text func(
		buffer *GtkEntryBuffer,
		chars string,
		n_chars Gint)

	Gtk_entry_buffer_set_max_length func(
		buffer *GtkEntryBuffer,
		max_length Gint)

	Gtk_entry_buffer_get_max_length func(
		buffer *GtkEntryBuffer) Gint

	Gtk_entry_buffer_insert_text func(
		buffer *GtkEntryBuffer,
		position Guint,
		chars string,
		n_chars Gint) Guint

	Gtk_entry_buffer_delete_text func(
		buffer *GtkEntryBuffer,
		position Guint,
		n_chars Gint) Guint

	Gtk_entry_buffer_emit_inserted_text func(
		buffer *GtkEntryBuffer,
		position Guint,
		chars string,
		n_chars Guint)

	Gtk_entry_buffer_emit_deleted_text func(
		buffer *GtkEntryBuffer,
		position Guint,
		n_chars Guint)

	Gtk_list_store_get_type func() GType

	Gtk_list_store_new func(n_columns Gint,
		v ...VArg) *GtkListStore

	Gtk_list_store_newv func(
		n_columns Gint,
		types *GType) *GtkListStore

	Gtk_list_store_set_column_types func(
		list_store *GtkListStore,
		n_columns Gint,
		types *GType)

	Gtk_list_store_set_value func(
		list_store *GtkListStore,
		iter *GtkTreeIter,
		column Gint,
		value *GValue)

	Gtk_list_store_set func(list_store *GtkListStore,
		iter *GtkTreeIter, v ...VArg)

	Gtk_list_store_set_valuesv func(
		list_store *GtkListStore,
		iter *GtkTreeIter,
		columns *Gint,
		values *GValue,
		n_values Gint)

	Gtk_list_store_set_valist func(
		list_store *GtkListStore,
		iter *GtkTreeIter,
		var_args Va_list)

	Gtk_list_store_remove func(
		list_store *GtkListStore,
		iter *GtkTreeIter) Gboolean

	Gtk_list_store_insert func(
		list_store *GtkListStore,
		iter *GtkTreeIter,
		position Gint)

	Gtk_list_store_insert_before func(
		list_store *GtkListStore,
		iter *GtkTreeIter,
		sibling *GtkTreeIter)

	Gtk_list_store_insert_after func(
		list_store *GtkListStore,
		iter *GtkTreeIter,
		sibling *GtkTreeIter)

	Gtk_list_store_insert_with_values func(list_store *GtkListStore,
		iter *GtkTreeIter, position Gint, v ...VArg)

	Gtk_list_store_insert_with_valuesv func(
		list_store *GtkListStore,
		iter *GtkTreeIter,
		position Gint,
		columns *Gint,
		values *GValue,
		n_values Gint)

	Gtk_list_store_prepend func(
		list_store *GtkListStore,
		iter *GtkTreeIter)

	Gtk_list_store_append func(
		list_store *GtkListStore,
		iter *GtkTreeIter)

	Gtk_list_store_clear func(
		list_store *GtkListStore)

	Gtk_list_store_iter_is_valid func(
		list_store *GtkListStore,
		iter *GtkTreeIter) Gboolean

	Gtk_list_store_reorder func(
		store *GtkListStore,
		new_order *Gint)

	Gtk_list_store_swap func(
		store *GtkListStore,
		a *GtkTreeIter,
		b *GtkTreeIter)

	Gtk_list_store_move_after func(
		store *GtkListStore,
		iter *GtkTreeIter,
		position *GtkTreeIter)

	Gtk_list_store_move_before func(
		store *GtkListStore,
		iter *GtkTreeIter,
		position *GtkTreeIter)

	Gtk_tree_model_filter_get_type func() GType

	Gtk_tree_model_filter_new func(
		child_model *GtkTreeModel,
		root *GtkTreePath) *GtkTreeModel

	Gtk_tree_model_filter_set_visible_func func(
		filter *GtkTreeModelFilter,
		f GtkTreeModelFilterVisibleFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_model_filter_set_modify_func func(
		filter *GtkTreeModelFilter,
		n_columns Gint,
		types *GType,
		f GtkTreeModelFilterModifyFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_model_filter_set_visible_column func(
		filter *GtkTreeModelFilter,
		column Gint)

	Gtk_tree_model_filter_get_model func(
		filter *GtkTreeModelFilter) *GtkTreeModel

	Gtk_tree_model_filter_convert_child_iter_to_iter func(
		filter *GtkTreeModelFilter,
		filter_iter *GtkTreeIter,
		child_iter *GtkTreeIter) Gboolean

	Gtk_tree_model_filter_convert_iter_to_child_iter func(
		filter *GtkTreeModelFilter,
		child_iter *GtkTreeIter,
		filter_iter *GtkTreeIter)

	Gtk_tree_model_filter_convert_child_path_to_path func(
		filter *GtkTreeModelFilter,
		child_path *GtkTreePath) *GtkTreePath

	Gtk_tree_model_filter_convert_path_to_child_path func(
		filter *GtkTreeModelFilter,
		filter_path *GtkTreePath) *GtkTreePath

	Gtk_tree_model_filter_refilter func(
		filter *GtkTreeModelFilter)

	Gtk_tree_model_filter_clear_cache func(
		filter *GtkTreeModelFilter)

	Gtk_entry_completion_get_type func() GType

	Gtk_entry_completion_new func() *GtkEntryCompletion

	Gtk_entry_completion_get_entry func(
		completion *GtkEntryCompletion) *GtkWidget

	Gtk_entry_completion_set_model func(
		completion *GtkEntryCompletion,
		model *GtkTreeModel)

	Gtk_entry_completion_get_model func(
		completion *GtkEntryCompletion) *GtkTreeModel

	Gtk_entry_completion_set_match_func func(
		completion *GtkEntryCompletion,
		f GtkEntryCompletionMatchFunc,
		func_dataGpointer,
		func_notify GDestroyNotify)

	Gtk_entry_completion_set_minimum_key_length func(
		completion *GtkEntryCompletion,
		length Gint)

	Gtk_entry_completion_get_minimum_key_length func(
		completion *GtkEntryCompletion) Gint

	Gtk_entry_completion_complete func(
		completion *GtkEntryCompletion)

	Gtk_entry_completion_insert_prefix func(
		completion *GtkEntryCompletion)

	Gtk_entry_completion_insert_action_text func(
		completion *GtkEntryCompletion,
		index_ Gint,
		text string)

	Gtk_entry_completion_insert_action_markup func(
		completion *GtkEntryCompletion,
		index_ Gint,
		markup string)

	Gtk_entry_completion_delete_action func(
		completion *GtkEntryCompletion,
		index_ Gint)

	Gtk_entry_completion_set_inline_completion func(
		completion *GtkEntryCompletion,
		inline_completion Gboolean)

	Gtk_entry_completion_get_inline_completion func(
		completion *GtkEntryCompletion) Gboolean

	Gtk_entry_completion_set_inline_selection func(
		completion *GtkEntryCompletion,
		inline_selection Gboolean)

	Gtk_entry_completion_get_inline_selection func(
		completion *GtkEntryCompletion) Gboolean

	Gtk_entry_completion_set_popup_completion func(
		completion *GtkEntryCompletion,
		popup_completion Gboolean)

	Gtk_entry_completion_get_popup_completion func(
		completion *GtkEntryCompletion) Gboolean

	Gtk_entry_completion_set_popup_set_width func(
		completion *GtkEntryCompletion,
		popup_set_width Gboolean)

	Gtk_entry_completion_get_popup_set_width func(
		completion *GtkEntryCompletion) Gboolean

	Gtk_entry_completion_set_popup_single_match func(
		completion *GtkEntryCompletion,
		popup_single_match Gboolean)

	Gtk_entry_completion_get_popup_single_match func(
		completion *GtkEntryCompletion) Gboolean

	Gtk_entry_completion_get_completion_prefix func(
		completion *GtkEntryCompletion) string

	Gtk_entry_completion_set_text_column func(
		completion *GtkEntryCompletion,
		column Gint)

	Gtk_entry_completion_get_text_column func(
		completion *GtkEntryCompletion) Gint

	Gtk_entry_get_type func() GType

	Gtk_entry_new func() *GtkWidget

	Gtk_entry_new_with_buffer func(
		buffer *GtkEntryBuffer) *GtkWidget

	Gtk_entry_get_buffer func(
		entry *GtkEntry) *GtkEntryBuffer

	Gtk_entry_set_buffer func(
		entry *GtkEntry,
		buffer *GtkEntryBuffer)

	Gtk_entry_get_text_window func(
		entry *GtkEntry) *GdkWindow

	Gtk_entry_set_visibility func(
		entry *GtkEntry,
		visible Gboolean)

	Gtk_entry_get_visibility func(
		entry *GtkEntry) Gboolean

	Gtk_entry_set_invisible_char func(
		entry *GtkEntry,
		ch Gunichar)

	Gtk_entry_get_invisible_char func(
		entry *GtkEntry) Gunichar

	Gtk_entry_unset_invisible_char func(
		entry *GtkEntry)

	Gtk_entry_set_has_frame func(
		entry *GtkEntry,
		setting Gboolean)

	Gtk_entry_get_has_frame func(
		entry *GtkEntry) Gboolean

	Gtk_entry_set_inner_border func(
		entry *GtkEntry,
		border *GtkBorder)

	Gtk_entry_get_inner_border func(
		entry *GtkEntry) *GtkBorder

	Gtk_entry_set_overwrite_mode func(
		entry *GtkEntry,
		overwrite Gboolean)

	Gtk_entry_get_overwrite_mode func(
		entry *GtkEntry) Gboolean

	Gtk_entry_set_max_length func(
		entry *GtkEntry,
		max Gint)

	Gtk_entry_get_max_length func(
		entry *GtkEntry) Gint

	Gtk_entry_get_text_length func(
		entry *GtkEntry) Guint16

	Gtk_entry_set_activates_default func(
		entry *GtkEntry,
		setting Gboolean)

	Gtk_entry_get_activates_default func(
		entry *GtkEntry) Gboolean

	Gtk_entry_set_width_chars func(
		entry *GtkEntry,
		n_chars Gint)

	Gtk_entry_get_width_chars func(
		entry *GtkEntry) Gint

	Gtk_entry_set_text func(
		entry *GtkEntry,
		text string)

	Gtk_entry_get_text func(
		entry *GtkEntry) string

	Gtk_entry_get_layout func(
		entry *GtkEntry) *PangoLayout

	Gtk_entry_get_layout_offsets func(
		entry *GtkEntry,
		x *Gint,
		y *Gint)

	Gtk_entry_set_alignment func(
		entry *GtkEntry,
		xalign Gfloat)

	Gtk_entry_get_alignment func(
		entry *GtkEntry) Gfloat

	Gtk_entry_set_completion func(
		entry *GtkEntry,
		completion *GtkEntryCompletion)

	Gtk_entry_get_completion func(
		entry *GtkEntry) *GtkEntryCompletion

	Gtk_entry_layout_index_to_text_index func(
		entry *GtkEntry,
		layout_index Gint) Gint

	Gtk_entry_text_index_to_layout_index func(
		entry *GtkEntry,
		text_index Gint) Gint

	Gtk_entry_set_cursor_hadjustment func(
		entry *GtkEntry,
		adjustment *GtkAdjustment)

	Gtk_entry_get_cursor_hadjustment func(
		entry *GtkEntry) *GtkAdjustment

	Gtk_entry_set_progress_fraction func(
		entry *GtkEntry,
		fraction Gdouble)

	Gtk_entry_get_progress_fraction func(
		entry *GtkEntry) Gdouble

	Gtk_entry_set_progress_pulse_step func(
		entry *GtkEntry,
		fraction Gdouble)

	Gtk_entry_get_progress_pulse_step func(
		entry *GtkEntry) Gdouble

	Gtk_entry_progress_pulse func(
		entry *GtkEntry)

	Gtk_entry_set_icon_from_pixbuf func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		pixbuf *GdkPixbuf)

	Gtk_entry_set_icon_from_stock func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		stock_id string)

	Gtk_entry_set_icon_from_icon_name func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		icon_name string)

	Gtk_entry_set_icon_from_gicon func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		icon *GIcon)

	Gtk_entry_get_icon_storage_type func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) GtkImageType

	Gtk_entry_get_icon_pixbuf func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) *GdkPixbuf

	Gtk_entry_get_icon_stock func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) string

	Gtk_entry_get_icon_name func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) string

	Gtk_entry_get_icon_gicon func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) *GIcon

	Gtk_entry_set_icon_activatable func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		activatable Gboolean)

	Gtk_entry_get_icon_activatable func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) Gboolean

	Gtk_entry_set_icon_sensitive func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		sensitive Gboolean)

	Gtk_entry_get_icon_sensitive func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) Gboolean

	Gtk_entry_get_icon_at_pos func(
		entry *GtkEntry,
		x Gint,
		y Gint) Gint

	Gtk_entry_set_icon_tooltip_text func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		tooltip string)

	Gtk_entry_get_icon_tooltip_text func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) string

	Gtk_entry_set_icon_tooltip_markup func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		tooltip string)

	Gtk_entry_get_icon_tooltip_markup func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) string

	Gtk_entry_set_icon_drag_source func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition,
		target_list *GtkTargetList,
		actions GdkDragAction)

	Gtk_entry_get_current_icon_drag_source func(
		entry *GtkEntry) Gint

	Gtk_entry_get_icon_window func(
		entry *GtkEntry,
		icon_pos GtkEntryIconPosition) *GdkWindow

	Gtk_entry_im_context_filter_keypress func(
		entry *GtkEntry,
		event *GdkEventKey) Gboolean

	Gtk_entry_reset_im_context func(
		entry *GtkEntry)

	Gtk_entry_new_with_max_length func(
		max Gint) *GtkWidget

	Gtk_entry_append_text func(
		entry *GtkEntry,
		text string)

	Gtk_entry_prepend_text func(
		entry *GtkEntry,
		text string)

	Gtk_entry_set_position func(
		entry *GtkEntry,
		position Gint)

	Gtk_entry_select_region func(
		entry *GtkEntry,
		start Gint,
		end Gint)

	Gtk_entry_set_editable func(
		entry *GtkEntry,
		editable Gboolean)

	Gtk_tree_view_get_type func() GType

	Gtk_tree_view_new func() *GtkWidget

	Gtk_tree_view_new_with_model func(
		model *GtkTreeModel) *GtkWidget

	Gtk_tree_view_get_model func(
		tree_view *GtkTreeView) *GtkTreeModel

	Gtk_tree_view_set_model func(
		tree_view *GtkTreeView,
		model *GtkTreeModel)

	Gtk_tree_view_get_selection func(
		tree_view *GtkTreeView) *GtkTreeSelection

	Gtk_tree_view_get_hadjustment func(
		tree_view *GtkTreeView) *GtkAdjustment

	Gtk_tree_view_set_hadjustment func(
		tree_view *GtkTreeView,
		adjustment *GtkAdjustment)

	Gtk_tree_view_get_vadjustment func(
		tree_view *GtkTreeView) *GtkAdjustment

	Gtk_tree_view_set_vadjustment func(
		tree_view *GtkTreeView,
		adjustment *GtkAdjustment)

	Gtk_tree_view_get_headers_visible func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_set_headers_visible func(
		tree_view *GtkTreeView,
		headers_visible Gboolean)

	Gtk_tree_view_columns_autosize func(
		tree_view *GtkTreeView)

	Gtk_tree_view_get_headers_clickable func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_set_headers_clickable func(
		tree_view *GtkTreeView,
		setting Gboolean)

	Gtk_tree_view_set_rules_hint func(
		tree_view *GtkTreeView,
		setting Gboolean)

	Gtk_tree_view_get_rules_hint func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_append_column func(
		tree_view *GtkTreeView,
		column *GtkTreeViewColumn) Gint

	Gtk_tree_view_remove_column func(
		tree_view *GtkTreeView,
		column *GtkTreeViewColumn) Gint

	Gtk_tree_view_insert_column func(
		tree_view *GtkTreeView,
		column *GtkTreeViewColumn,
		position Gint) Gint

	Gtk_tree_view_insert_column_with_attributes func(
		tree_view *GtkTreeView, position Gint, title string,
		cell *GtkCellRenderer, v ...VArg) Gint

	Gtk_tree_view_insert_column_with_data_func func(
		tree_view *GtkTreeView,
		position Gint,
		title string,
		cell *GtkCellRenderer,
		f GtkTreeCellDataFunc,
		dataGpointer,
		dnotify GDestroyNotify) Gint

	Gtk_tree_view_get_column func(
		tree_view *GtkTreeView,
		n Gint) *GtkTreeViewColumn

	Gtk_tree_view_get_columns func(
		tree_view *GtkTreeView) *GList

	Gtk_tree_view_move_column_after func(
		tree_view *GtkTreeView,
		column *GtkTreeViewColumn,
		base_column *GtkTreeViewColumn)

	Gtk_tree_view_set_expander_column func(
		tree_view *GtkTreeView,
		column *GtkTreeViewColumn)

	Gtk_tree_view_get_expander_column func(
		tree_view *GtkTreeView) *GtkTreeViewColumn

	Gtk_tree_view_set_column_drag_function func(
		tree_view *GtkTreeView,
		f GtkTreeViewColumnDropFunc,
		user_dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_view_scroll_to_point func(
		tree_view *GtkTreeView,
		tree_x Gint,
		tree_y Gint)

	Gtk_tree_view_scroll_to_cell func(
		tree_view *GtkTreeView,
		path *GtkTreePath,
		column *GtkTreeViewColumn,
		use_align Gboolean,
		row_align Gfloat,
		col_align Gfloat)

	Gtk_tree_view_row_activated func(
		tree_view *GtkTreeView,
		path *GtkTreePath,
		column *GtkTreeViewColumn)

	Gtk_tree_view_expand_all func(
		tree_view *GtkTreeView)

	Gtk_tree_view_collapse_all func(
		tree_view *GtkTreeView)

	Gtk_tree_view_expand_to_path func(
		tree_view *GtkTreeView,
		path *GtkTreePath)

	Gtk_tree_view_expand_row func(
		tree_view *GtkTreeView,
		path *GtkTreePath,
		open_all Gboolean) Gboolean

	Gtk_tree_view_collapse_row func(
		tree_view *GtkTreeView,
		path *GtkTreePath) Gboolean

	Gtk_tree_view_map_expanded_rows func(
		tree_view *GtkTreeView,
		f GtkTreeViewMappingFunc,
		data Gpointer)

	Gtk_tree_view_row_expanded func(
		tree_view *GtkTreeView,
		path *GtkTreePath) Gboolean

	Gtk_tree_view_set_reorderable func(
		tree_view *GtkTreeView,
		reorderable Gboolean)

	Gtk_tree_view_get_reorderable func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_set_cursor func(
		tree_view *GtkTreeView,
		path *GtkTreePath,
		focus_column *GtkTreeViewColumn,
		start_editing Gboolean)

	Gtk_tree_view_set_cursor_on_cell func(
		tree_view *GtkTreeView,
		path *GtkTreePath,
		focus_column *GtkTreeViewColumn,
		focus_cell *GtkCellRenderer,
		start_editing Gboolean)

	Gtk_tree_view_get_cursor func(
		tree_view *GtkTreeView,
		path **GtkTreePath,
		focus_column **GtkTreeViewColumn)

	Gtk_tree_view_get_bin_window func(
		tree_view *GtkTreeView) *GdkWindow

	Gtk_tree_view_get_path_at_pos func(
		tree_view *GtkTreeView,
		x Gint,
		y Gint,
		path **GtkTreePath,
		column **GtkTreeViewColumn,
		cell_x *Gint,
		cell_y *Gint) Gboolean

	Gtk_tree_view_get_cell_area func(
		tree_view *GtkTreeView,
		path *GtkTreePath,
		column *GtkTreeViewColumn,
		rect *GdkRectangle)

	Gtk_tree_view_get_background_area func(
		tree_view *GtkTreeView,
		path *GtkTreePath,
		column *GtkTreeViewColumn,
		rect *GdkRectangle)

	Gtk_tree_view_get_visible_rect func(
		tree_view *GtkTreeView,
		visible_rect *GdkRectangle)

	Gtk_tree_view_widget_to_tree_coords func(
		tree_view *GtkTreeView,
		wx Gint,
		wy Gint,
		tx *Gint,
		ty *Gint)

	Gtk_tree_view_tree_to_widget_coords func(
		tree_view *GtkTreeView,
		tx Gint,
		ty Gint,
		wx *Gint,
		wy *Gint)

	Gtk_tree_view_get_visible_range func(
		tree_view *GtkTreeView,
		start_path **GtkTreePath,
		end_path **GtkTreePath) Gboolean

	Gtk_tree_view_enable_model_drag_source func(
		tree_view *GtkTreeView,
		start_button_mask GdkModifierType,
		targets *GtkTargetEntry,
		n_targets Gint,
		actions GdkDragAction)

	Gtk_tree_view_enable_model_drag_dest func(
		tree_view *GtkTreeView,
		targets *GtkTargetEntry,
		n_targets Gint,
		actions GdkDragAction)

	Gtk_tree_view_unset_rows_drag_source func(
		tree_view *GtkTreeView)

	Gtk_tree_view_unset_rows_drag_dest func(
		tree_view *GtkTreeView)

	Gtk_tree_view_set_drag_dest_row func(
		tree_view *GtkTreeView,
		path *GtkTreePath,
		pos GtkTreeViewDropPosition)

	Gtk_tree_view_get_drag_dest_row func(
		tree_view *GtkTreeView,
		path **GtkTreePath,
		pos *GtkTreeViewDropPosition)

	Gtk_tree_view_get_dest_row_at_pos func(
		tree_view *GtkTreeView,
		drag_x Gint,
		drag_y Gint,
		path **GtkTreePath,
		pos *GtkTreeViewDropPosition) Gboolean

	Gtk_tree_view_create_row_drag_icon func(
		tree_view *GtkTreeView,
		path *GtkTreePath) *GdkPixmap

	Gtk_tree_view_set_enable_search func(
		tree_view *GtkTreeView,
		enable_search Gboolean)

	Gtk_tree_view_get_enable_search func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_get_search_column func(
		tree_view *GtkTreeView) Gint

	Gtk_tree_view_set_search_column func(
		tree_view *GtkTreeView,
		column Gint)

	Gtk_tree_view_get_search_equal_func func(
		tree_view *GtkTreeView) GtkTreeViewSearchEqualFunc

	Gtk_tree_view_set_search_equal_func func(
		tree_view *GtkTreeView,
		search_equal_func GtkTreeViewSearchEqualFunc,
		search_user_dataGpointer,
		search_destroy GDestroyNotify)

	Gtk_tree_view_get_search_entry func(
		tree_view *GtkTreeView) *GtkEntry

	Gtk_tree_view_set_search_entry func(
		tree_view *GtkTreeView,
		entry *GtkEntry)

	Gtk_tree_view_get_search_position_func func(
		tree_view *GtkTreeView) GtkTreeViewSearchPositionFunc

	Gtk_tree_view_set_search_position_func func(
		tree_view *GtkTreeView,
		f GtkTreeViewSearchPositionFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_view_convert_widget_to_tree_coords func(
		tree_view *GtkTreeView,
		wx Gint,
		wy Gint,
		tx *Gint,
		ty *Gint)

	Gtk_tree_view_convert_tree_to_widget_coords func(
		tree_view *GtkTreeView,
		tx Gint,
		ty Gint,
		wx *Gint,
		wy *Gint)

	Gtk_tree_view_convert_widget_to_bin_window_coords func(
		tree_view *GtkTreeView,
		wx Gint,
		wy Gint,
		bx *Gint,
		by *Gint)

	Gtk_tree_view_convert_bin_window_to_widget_coords func(
		tree_view *GtkTreeView,
		bx Gint,
		by Gint,
		wx *Gint,
		wy *Gint)

	Gtk_tree_view_convert_tree_to_bin_window_coords func(
		tree_view *GtkTreeView,
		tx Gint,
		ty Gint,
		bx *Gint,
		by *Gint)

	Gtk_tree_view_convert_bin_window_to_tree_coords func(
		tree_view *GtkTreeView,
		bx Gint,
		by Gint,
		tx *Gint,
		ty *Gint)

	Gtk_tree_view_set_destroy_count_func func(
		tree_view *GtkTreeView,
		f GtkTreeDestroyCountFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_view_set_fixed_height_mode func(
		tree_view *GtkTreeView,
		enable Gboolean)

	Gtk_tree_view_get_fixed_height_mode func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_set_hover_selection func(
		tree_view *GtkTreeView,
		hover Gboolean)

	Gtk_tree_view_get_hover_selection func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_set_hover_expand func(
		tree_view *GtkTreeView,
		expand Gboolean)

	Gtk_tree_view_get_hover_expand func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_set_rubber_banding func(
		tree_view *GtkTreeView,
		enable Gboolean)

	Gtk_tree_view_get_rubber_banding func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_is_rubber_banding_active func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_get_row_separator_func func(
		tree_view *GtkTreeView) GtkTreeViewRowSeparatorFunc

	Gtk_tree_view_set_row_separator_func func(
		tree_view *GtkTreeView,
		f GtkTreeViewRowSeparatorFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_view_get_grid_lines func(
		tree_view *GtkTreeView) GtkTreeViewGridLines

	Gtk_tree_view_set_grid_lines func(
		tree_view *GtkTreeView,
		grid_lines GtkTreeViewGridLines)

	Gtk_tree_view_get_enable_tree_lines func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_set_enable_tree_lines func(
		tree_view *GtkTreeView,
		enabled Gboolean)

	Gtk_tree_view_set_show_expanders func(
		tree_view *GtkTreeView,
		enabled Gboolean)

	Gtk_tree_view_get_show_expanders func(
		tree_view *GtkTreeView) Gboolean

	Gtk_tree_view_set_level_indentation func(
		tree_view *GtkTreeView,
		indentation Gint)

	Gtk_tree_view_get_level_indentation func(
		tree_view *GtkTreeView) Gint

	Gtk_tree_view_set_tooltip_row func(
		tree_view *GtkTreeView,
		tooltip *GtkTooltip,
		path *GtkTreePath)

	Gtk_tree_view_set_tooltip_cell func(
		tree_view *GtkTreeView,
		tooltip *GtkTooltip,
		path *GtkTreePath,
		column *GtkTreeViewColumn,
		cell *GtkCellRenderer)

	Gtk_tree_view_get_tooltip_context func(
		tree_view *GtkTreeView,
		x *Gint,
		y *Gint,
		keyboard_tip Gboolean,
		model **GtkTreeModel,
		path **GtkTreePath,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_view_set_tooltip_column func(
		tree_view *GtkTreeView,
		column Gint)

	Gtk_tree_view_get_tooltip_column func(
		tree_view *GtkTreeView) Gint

	Gtk_combo_box_get_type func() GType

	Gtk_combo_box_new func() *GtkWidget

	Gtk_combo_box_new_with_entry func() *GtkWidget

	Gtk_combo_box_new_with_model func(
		model *GtkTreeModel) *GtkWidget

	Gtk_combo_box_new_with_model_and_entry func(
		model *GtkTreeModel) *GtkWidget

	Gtk_combo_box_get_wrap_width func(
		combo_box *GtkComboBox) Gint

	Gtk_combo_box_set_wrap_width func(
		combo_box *GtkComboBox,
		width Gint)

	Gtk_combo_box_get_row_span_column func(
		combo_box *GtkComboBox) Gint

	Gtk_combo_box_set_row_span_column func(
		combo_box *GtkComboBox,
		row_span Gint)

	Gtk_combo_box_get_column_span_column func(
		combo_box *GtkComboBox) Gint

	Gtk_combo_box_set_column_span_column func(
		combo_box *GtkComboBox,
		column_span Gint)

	Gtk_combo_box_get_add_tearoffs func(
		combo_box *GtkComboBox) Gboolean

	Gtk_combo_box_set_add_tearoffs func(
		combo_box *GtkComboBox,
		add_tearoffs Gboolean)

	Gtk_combo_box_get_title func(
		combo_box *GtkComboBox) string

	Gtk_combo_box_set_title func(
		combo_box *GtkComboBox,
		title string)

	Gtk_combo_box_get_focus_on_click func(
		combo *GtkComboBox) Gboolean

	Gtk_combo_box_set_focus_on_click func(
		combo *GtkComboBox,
		focus_on_click Gboolean)

	Gtk_combo_box_get_active func(
		combo_box *GtkComboBox) Gint

	Gtk_combo_box_set_active func(
		combo_box *GtkComboBox,
		index_ Gint)

	Gtk_combo_box_get_active_iter func(
		combo_box *GtkComboBox,
		iter *GtkTreeIter) Gboolean

	Gtk_combo_box_set_active_iter func(
		combo_box *GtkComboBox,
		iter *GtkTreeIter)

	Gtk_combo_box_set_model func(
		combo_box *GtkComboBox,
		model *GtkTreeModel)

	Gtk_combo_box_get_model func(
		combo_box *GtkComboBox) *GtkTreeModel

	Gtk_combo_box_get_row_separator_func func(
		combo_box *GtkComboBox) GtkTreeViewRowSeparatorFunc

	Gtk_combo_box_set_row_separator_func func(
		combo_box *GtkComboBox,
		f GtkTreeViewRowSeparatorFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_combo_box_set_button_sensitivity func(
		combo_box *GtkComboBox,
		sensitivity GtkSensitivityType)

	Gtk_combo_box_get_button_sensitivity func(
		combo_box *GtkComboBox) GtkSensitivityType

	Gtk_combo_box_get_has_entry func(
		combo_box *GtkComboBox) Gboolean

	Gtk_combo_box_set_entry_text_column func(
		combo_box *GtkComboBox,
		text_column Gint)

	Gtk_combo_box_get_entry_text_column func(
		combo_box *GtkComboBox) Gint

	Gtk_combo_box_new_text func() *GtkWidget

	Gtk_combo_box_append_text func(
		combo_box *GtkComboBox,
		text string)

	Gtk_combo_box_insert_text func(
		combo_box *GtkComboBox,
		position Gint,
		text string)

	Gtk_combo_box_prepend_text func(
		combo_box *GtkComboBox,
		text string)

	Gtk_combo_box_remove_text func(
		combo_box *GtkComboBox,
		position Gint)

	Gtk_combo_box_get_active_text func(
		combo_box *GtkComboBox) string

	Gtk_combo_box_popup func(
		combo_box *GtkComboBox)

	Gtk_combo_box_popdown func(
		combo_box *GtkComboBox)

	Gtk_combo_box_get_popup_accessible func(
		combo_box *GtkComboBox) *AtkObject

	Gtk_combo_box_entry_get_type func() GType

	Gtk_combo_box_entry_new func() *GtkWidget

	Gtk_combo_box_entry_new_with_model func(
		model *GtkTreeModel,
		text_column Gint) *GtkWidget

	Gtk_combo_box_entry_set_text_column func(
		entry_box *GtkComboBoxEntry,
		text_column Gint)

	Gtk_combo_box_entry_get_text_column func(
		entry_box *GtkComboBoxEntry) Gint

	Gtk_combo_box_entry_new_text func() *GtkWidget

	Gtk_combo_box_text_get_type func() GType

	Gtk_combo_box_text_new func() *GtkWidget

	Gtk_combo_box_text_new_with_entry func() *GtkWidget

	Gtk_combo_box_text_append_text func(
		combo_box *GtkComboBoxText,
		text string)

	Gtk_combo_box_text_insert_text func(
		combo_box *GtkComboBoxText,
		position Gint,
		text string)

	Gtk_combo_box_text_prepend_text func(
		combo_box *GtkComboBoxText,
		text string)

	Gtk_combo_box_text_remove func(
		combo_box *GtkComboBoxText,
		position Gint)

	Gtk_combo_box_text_get_active_text func(
		combo_box *GtkComboBoxText) string

	Gtk_drawing_area_get_type func() GType

	Gtk_drawing_area_new func() *GtkWidget

	Gtk_drawing_area_size func(
		darea *GtkDrawingArea,
		width Gint,
		height Gint)

	Gtk_event_box_get_type func() GType

	Gtk_event_box_new func() *GtkWidget

	Gtk_event_box_get_visible_window func(
		event_box *GtkEventBox) Gboolean

	Gtk_event_box_set_visible_window func(
		event_box *GtkEventBox,
		visible_window Gboolean)

	Gtk_event_box_get_above_child func(
		event_box *GtkEventBox) Gboolean

	Gtk_event_box_set_above_child func(
		event_box *GtkEventBox,
		above_child Gboolean)

	Gtk_expander_get_type func() GType

	Gtk_expander_new func(
		label string) *GtkWidget

	Gtk_expander_new_with_mnemonic func(
		label string) *GtkWidget

	Gtk_expander_set_expanded func(
		expander *GtkExpander,
		expanded Gboolean)

	Gtk_expander_get_expanded func(
		expander *GtkExpander) Gboolean

	Gtk_expander_set_spacing func(
		expander *GtkExpander,
		spacing Gint)

	Gtk_expander_get_spacing func(
		expander *GtkExpander) Gint

	Gtk_expander_set_label func(
		expander *GtkExpander,
		label string)

	Gtk_expander_get_label func(
		expander *GtkExpander) string

	Gtk_expander_set_use_underline func(
		expander *GtkExpander,
		use_underline Gboolean)

	Gtk_expander_get_use_underline func(
		expander *GtkExpander) Gboolean

	Gtk_expander_set_use_markup func(
		expander *GtkExpander,
		use_markup Gboolean)

	Gtk_expander_get_use_markup func(
		expander *GtkExpander) Gboolean

	Gtk_expander_set_label_widget func(
		expander *GtkExpander,
		label_widget *GtkWidget)

	Gtk_expander_get_label_widget func(
		expander *GtkExpander) *GtkWidget

	Gtk_expander_set_label_fill func(
		expander *GtkExpander,
		label_fill Gboolean)

	Gtk_expander_get_label_fill func(
		expander *GtkExpander) Gboolean

	Gtk_fixed_get_type func() GType

	Gtk_fixed_new func() *GtkWidget

	Gtk_fixed_put func(
		fixed *GtkFixed,
		widget *GtkWidget,
		x Gint,
		y Gint)

	Gtk_fixed_move func(
		fixed *GtkFixed,
		widget *GtkWidget,
		x Gint,
		y Gint)

	Gtk_fixed_set_has_window func(
		fixed *GtkFixed,
		has_window Gboolean)

	Gtk_fixed_get_has_window func(
		fixed *GtkFixed) Gboolean

	Gtk_file_filter_get_type func() GType

	Gtk_file_filter_new func() *GtkFileFilter

	Gtk_file_filter_set_name func(
		filter *GtkFileFilter,
		name string)

	Gtk_file_filter_get_name func(
		filter *GtkFileFilter) string

	Gtk_file_filter_add_mime_type func(
		filter *GtkFileFilter,
		mime_type string)

	Gtk_file_filter_add_pattern func(
		filter *GtkFileFilter,
		pattern string)

	Gtk_file_filter_add_pixbuf_formats func(
		filter *GtkFileFilter)

	Gtk_file_filter_add_custom func(
		filter *GtkFileFilter,
		needed GtkFileFilterFlags,
		f GtkFileFilterFunc,
		dataGpointer,
		notify GDestroyNotify)

	Gtk_file_filter_get_needed func(
		filter *GtkFileFilter) GtkFileFilterFlags

	Gtk_file_filter_filter func(
		filter *GtkFileFilter,
		filter_info *GtkFileFilterInfo) Gboolean

	Gtk_file_chooser_get_type func() GType

	Gtk_file_chooser_error_quark func() GQuark

	Gtk_file_chooser_set_action func(
		chooser *GtkFileChooser,
		action GtkFileChooserAction)

	Gtk_file_chooser_get_action func(
		chooser *GtkFileChooser) GtkFileChooserAction

	Gtk_file_chooser_set_local_only func(
		chooser *GtkFileChooser,
		local_only Gboolean)

	Gtk_file_chooser_get_local_only func(
		chooser *GtkFileChooser) Gboolean

	Gtk_file_chooser_set_select_multiple func(
		chooser *GtkFileChooser,
		select_multiple Gboolean)

	Gtk_file_chooser_get_select_multiple func(
		chooser *GtkFileChooser) Gboolean

	Gtk_file_chooser_set_show_hidden func(
		chooser *GtkFileChooser,
		show_hidden Gboolean)

	Gtk_file_chooser_get_show_hidden func(
		chooser *GtkFileChooser) Gboolean

	Gtk_file_chooser_set_do_overwrite_confirmation func(
		chooser *GtkFileChooser,
		do_overwrite_confirmation Gboolean)

	Gtk_file_chooser_get_do_overwrite_confirmation func(
		chooser *GtkFileChooser) Gboolean

	Gtk_file_chooser_set_create_folders func(
		chooser *GtkFileChooser,
		create_folders Gboolean)

	Gtk_file_chooser_get_create_folders func(
		chooser *GtkFileChooser) Gboolean

	Gtk_file_chooser_set_current_name func(
		chooser *GtkFileChooser,
		name string)

	Gtk_file_chooser_get_filename func(
		chooser *GtkFileChooser) string

	Gtk_file_chooser_set_filename func(
		chooser *GtkFileChooser,
		filename string) Gboolean

	Gtk_file_chooser_select_filename func(
		chooser *GtkFileChooser,
		filename string) Gboolean

	Gtk_file_chooser_unselect_filename func(
		chooser *GtkFileChooser,
		filename string)

	Gtk_file_chooser_select_all func(
		chooser *GtkFileChooser)

	Gtk_file_chooser_unselect_all func(
		chooser *GtkFileChooser)

	Gtk_file_chooser_get_filenames func(
		chooser *GtkFileChooser) *GSList

	Gtk_file_chooser_set_current_folder func(
		chooser *GtkFileChooser,
		filename string) Gboolean

	Gtk_file_chooser_get_current_folder func(
		chooser *GtkFileChooser) string

	Gtk_file_chooser_get_uri func(
		chooser *GtkFileChooser) string

	Gtk_file_chooser_set_uri func(
		chooser *GtkFileChooser,
		uri string) Gboolean

	Gtk_file_chooser_select_uri func(
		chooser *GtkFileChooser,
		uri string) Gboolean

	Gtk_file_chooser_unselect_uri func(
		chooser *GtkFileChooser,
		uri string)

	Gtk_file_chooser_get_uris func(
		chooser *GtkFileChooser) *GSList

	Gtk_file_chooser_set_current_folder_uri func(
		chooser *GtkFileChooser,
		uri string) Gboolean

	Gtk_file_chooser_get_current_folder_uri func(
		chooser *GtkFileChooser) string

	Gtk_file_chooser_get_file func(
		chooser *GtkFileChooser) *GFile

	Gtk_file_chooser_set_file func(
		chooser *GtkFileChooser,
		file *GFile,
		error **GError) Gboolean

	Gtk_file_chooser_select_file func(
		chooser *GtkFileChooser,
		file *GFile,
		error **GError) Gboolean

	Gtk_file_chooser_unselect_file func(
		chooser *GtkFileChooser,
		file *GFile)

	Gtk_file_chooser_get_files func(
		chooser *GtkFileChooser) *GSList

	Gtk_file_chooser_set_current_folder_file func(
		chooser *GtkFileChooser,
		file *GFile,
		error **GError) Gboolean

	Gtk_file_chooser_get_current_folder_file func(
		chooser *GtkFileChooser) *GFile

	Gtk_file_chooser_set_preview_widget func(
		chooser *GtkFileChooser,
		preview_widget *GtkWidget)

	Gtk_file_chooser_get_preview_widget func(
		chooser *GtkFileChooser) *GtkWidget

	Gtk_file_chooser_set_preview_widget_active func(
		chooser *GtkFileChooser,
		active Gboolean)

	Gtk_file_chooser_get_preview_widget_active func(
		chooser *GtkFileChooser) Gboolean

	Gtk_file_chooser_set_use_preview_label func(
		chooser *GtkFileChooser,
		use_label Gboolean)

	Gtk_file_chooser_get_use_preview_label func(
		chooser *GtkFileChooser) Gboolean

	Gtk_file_chooser_get_preview_filename func(
		chooser *GtkFileChooser) string

	Gtk_file_chooser_get_preview_uri func(
		chooser *GtkFileChooser) string

	Gtk_file_chooser_get_preview_file func(
		chooser *GtkFileChooser) *GFile

	Gtk_file_chooser_set_extra_widget func(
		chooser *GtkFileChooser,
		extra_widget *GtkWidget)

	Gtk_file_chooser_get_extra_widget func(
		chooser *GtkFileChooser) *GtkWidget

	Gtk_file_chooser_add_filter func(
		chooser *GtkFileChooser,
		filter *GtkFileFilter)

	Gtk_file_chooser_remove_filter func(
		chooser *GtkFileChooser,
		filter *GtkFileFilter)

	Gtk_file_chooser_list_filters func(
		chooser *GtkFileChooser) *GSList

	Gtk_file_chooser_set_filter func(
		chooser *GtkFileChooser,
		filter *GtkFileFilter)

	Gtk_file_chooser_get_filter func(
		chooser *GtkFileChooser) *GtkFileFilter

	Gtk_file_chooser_add_shortcut_folder func(
		chooser *GtkFileChooser,
		folder string,
		error **GError) Gboolean

	Gtk_file_chooser_remove_shortcut_folder func(
		chooser *GtkFileChooser,
		folder string,
		error **GError) Gboolean

	Gtk_file_chooser_list_shortcut_folders func(
		chooser *GtkFileChooser) *GSList

	Gtk_file_chooser_add_shortcut_folder_uri func(
		chooser *GtkFileChooser,
		uri string,
		error **GError) Gboolean

	Gtk_file_chooser_remove_shortcut_folder_uri func(
		chooser *GtkFileChooser,
		uri string,
		error **GError) Gboolean

	Gtk_file_chooser_list_shortcut_folder_uris func(
		chooser *GtkFileChooser) *GSList

	Gtk_hbox_get_type func() GType

	Gtk_hbox_new func(
		homogeneous Gboolean,
		spacing Gint) *GtkWidget

	Gtk_file_chooser_button_get_type func() GType

	Gtk_file_chooser_button_new func(
		title string,
		action GtkFileChooserAction) *GtkWidget

	Gtk_file_chooser_button_new_with_backend func(
		title string,
		action GtkFileChooserAction,
		backend string) *GtkWidget

	Gtk_file_chooser_button_new_with_dialog func(
		dialog *GtkWidget) *GtkWidget

	Gtk_file_chooser_button_get_title func(
		button *GtkFileChooserButton) string

	Gtk_file_chooser_button_set_title func(
		button *GtkFileChooserButton,
		title string)

	Gtk_file_chooser_button_get_width_chars func(
		button *GtkFileChooserButton) Gint

	Gtk_file_chooser_button_set_width_chars func(
		button *GtkFileChooserButton,
		n_chars Gint)

	Gtk_file_chooser_button_get_focus_on_click func(
		button *GtkFileChooserButton) Gboolean

	Gtk_file_chooser_button_set_focus_on_click func(
		button *GtkFileChooserButton,
		focus_on_click Gboolean)

	Gtk_file_chooser_dialog_get_type func() GType

	Gtk_file_chooser_dialog_new func(title string, parent *GtkWindow,
		action GtkFileChooserAction, first_button_text string,
		v ...VArg) *GtkWidget

	Gtk_file_chooser_dialog_new_with_backend func(title string,
		parent *GtkWindow, action GtkFileChooserAction,
		backend string, first_button_text string,
		v ...VArg) *GtkWidget

	Gtk_file_chooser_widget_get_type func() GType

	Gtk_file_chooser_widget_new func(
		action GtkFileChooserAction) *GtkWidget

	Gtk_file_chooser_widget_new_with_backend func(
		action GtkFileChooserAction,
		backend string) *GtkWidget

	Gtk_font_button_get_type func() GType

	Gtk_font_button_new func() *GtkWidget

	Gtk_font_button_new_with_font func(
		fontname string) *GtkWidget

	Gtk_font_button_get_title func(
		font_button *GtkFontButton) string

	Gtk_font_button_set_title func(
		font_button *GtkFontButton,
		title string)

	Gtk_font_button_get_use_font func(
		font_button *GtkFontButton) Gboolean

	Gtk_font_button_set_use_font func(
		font_button *GtkFontButton,
		use_font Gboolean)

	Gtk_font_button_get_use_size func(
		font_button *GtkFontButton) Gboolean

	Gtk_font_button_set_use_size func(
		font_button *GtkFontButton,
		use_size Gboolean)

	Gtk_font_button_get_font_name func(
		font_button *GtkFontButton) string

	Gtk_font_button_set_font_name func(
		font_button *GtkFontButton,
		fontname string) Gboolean

	Gtk_font_button_get_show_style func(
		font_button *GtkFontButton) Gboolean

	Gtk_font_button_set_show_style func(
		font_button *GtkFontButton,
		show_style Gboolean)

	Gtk_font_button_get_show_size func(
		font_button *GtkFontButton) Gboolean

	Gtk_font_button_set_show_size func(
		font_button *GtkFontButton,
		show_size Gboolean)

	Gtk_font_selection_get_type func() GType

	Gtk_font_selection_new func() *GtkWidget

	Gtk_font_selection_get_family_list func(
		fontsel *GtkFontSelection) *GtkWidget

	Gtk_font_selection_get_face_list func(
		fontsel *GtkFontSelection) *GtkWidget

	Gtk_font_selection_get_size_entry func(
		fontsel *GtkFontSelection) *GtkWidget

	Gtk_font_selection_get_size_list func(
		fontsel *GtkFontSelection) *GtkWidget

	Gtk_font_selection_get_preview_entry func(
		fontsel *GtkFontSelection) *GtkWidget

	Gtk_font_selection_get_family func(
		fontsel *GtkFontSelection) *PangoFontFamily

	Gtk_font_selection_get_face func(
		fontsel *GtkFontSelection) *PangoFontFace

	Gtk_font_selection_get_size func(
		fontsel *GtkFontSelection) Gint

	Gtk_font_selection_get_font_name func(
		fontsel *GtkFontSelection) string

	Gtk_font_selection_get_font func(
		fontsel *GtkFontSelection) *GdkFont

	Gtk_font_selection_set_font_name func(
		fontsel *GtkFontSelection,
		fontname string) Gboolean

	Gtk_font_selection_get_preview_text func(
		fontsel *GtkFontSelection) string

	Gtk_font_selection_set_preview_text func(
		fontsel *GtkFontSelection,
		text string)

	Gtk_font_selection_dialog_get_type func() GType

	Gtk_font_selection_dialog_new func(
		title string) *GtkWidget

	Gtk_font_selection_dialog_get_ok_button func(
		fsd *GtkFontSelectionDialog) *GtkWidget

	Gtk_font_selection_dialog_get_apply_button func(
		fsd *GtkFontSelectionDialog) *GtkWidget

	Gtk_font_selection_dialog_get_cancel_button func(
		fsd *GtkFontSelectionDialog) *GtkWidget

	Gtk_font_selection_dialog_get_font_selection func(
		fsd *GtkFontSelectionDialog) *GtkWidget

	Gtk_font_selection_dialog_get_font_name func(
		fsd *GtkFontSelectionDialog) string

	Gtk_font_selection_dialog_get_font func(
		fsd *GtkFontSelectionDialog) *GdkFont

	Gtk_font_selection_dialog_set_font_name func(
		fsd *GtkFontSelectionDialog,
		fontname string) Gboolean

	Gtk_font_selection_dialog_get_preview_text func(
		fsd *GtkFontSelectionDialog) string

	Gtk_font_selection_dialog_set_preview_text func(
		fsd *GtkFontSelectionDialog,
		text string)

	Gtk_gc_get func(
		depth Gint,
		colormap *GdkColormap,
		values *GdkGCValues,
		values_mask GdkGCValuesMask) *GdkGC

	Gtk_gc_release func(
		gc *GdkGC)

	Gtk_handle_box_get_type func() GType

	Gtk_handle_box_new func() *GtkWidget

	Gtk_handle_box_set_shadow_type func(
		handle_box *GtkHandleBox,
		t GtkShadowType)

	Gtk_handle_box_get_shadow_type func(
		handle_box *GtkHandleBox) GtkShadowType

	Gtk_handle_box_set_handle_position func(
		handle_box *GtkHandleBox,
		position GtkPositionType)

	Gtk_handle_box_get_handle_position func(
		handle_box *GtkHandleBox) GtkPositionType

	Gtk_handle_box_set_snap_edge func(
		handle_box *GtkHandleBox,
		edge GtkPositionType)

	Gtk_handle_box_get_snap_edge func(
		handle_box *GtkHandleBox) GtkPositionType

	Gtk_handle_box_get_child_detached func(
		handle_box *GtkHandleBox) Gboolean

	Gtk_hbutton_box_get_type func() GType

	Gtk_hbutton_box_new func() *GtkWidget

	Gtk_hbutton_box_get_spacing_default func() Gint

	Gtk_hbutton_box_get_layout_default func() GtkButtonBoxStyle

	Gtk_hbutton_box_set_spacing_default func(
		spacing Gint)

	Gtk_hbutton_box_set_layout_default func(
		layout GtkButtonBoxStyle)

	Gtk_paned_get_type func() GType

	Gtk_paned_add1 func(
		paned *GtkPaned,
		child *GtkWidget)

	Gtk_paned_add2 func(
		paned *GtkPaned,
		child *GtkWidget)

	Gtk_paned_pack1 func(
		paned *GtkPaned,
		child *GtkWidget,
		resize Gboolean,
		shrink Gboolean)

	Gtk_paned_pack2 func(
		paned *GtkPaned,
		child *GtkWidget,
		resize Gboolean,
		shrink Gboolean)

	Gtk_paned_get_position func(
		paned *GtkPaned) Gint

	Gtk_paned_set_position func(
		paned *GtkPaned,
		position Gint)

	Gtk_paned_get_child1 func(
		paned *GtkPaned) *GtkWidget

	Gtk_paned_get_child2 func(
		paned *GtkPaned) *GtkWidget

	Gtk_paned_get_handle_window func(
		paned *GtkPaned) *GdkWindow

	Gtk_paned_compute_position func(
		paned *GtkPaned,
		allocation Gint,
		child1_req Gint,
		child2_req Gint)

	Gtk_hpaned_get_type func() GType

	Gtk_hpaned_new func() *GtkWidget

	Gtk_ruler_get_type func() GType

	Gtk_ruler_set_metric func(
		ruler *GtkRuler,
		metric GtkMetricType)

	Gtk_ruler_get_metric func(
		ruler *GtkRuler) GtkMetricType

	Gtk_ruler_set_range func(
		ruler *GtkRuler,
		lower Gdouble,
		upper Gdouble,
		position Gdouble,
		max_size Gdouble)

	Gtk_ruler_get_range func(
		ruler *GtkRuler,
		lower *Gdouble,
		upper *Gdouble,
		position *Gdouble,
		max_size *Gdouble)

	Gtk_ruler_draw_ticks func(
		ruler *GtkRuler)

	Gtk_ruler_draw_pos func(
		ruler *GtkRuler)

	Gtk_hruler_get_type func() GType

	Gtk_hruler_new func() *GtkWidget

	Gtk_range_get_type func() GType

	Gtk_range_set_update_policy func(
		r *GtkRange,
		policy GtkUpdateType)

	Gtk_range_get_update_policy func(
		r *GtkRange) GtkUpdateType

	Gtk_range_set_adjustment func(
		r *GtkRange,
		adjustment *GtkAdjustment)

	Gtk_range_get_adjustment func(
		r *GtkRange) *GtkAdjustment

	Gtk_range_set_inverted func(
		r *GtkRange,
		setting Gboolean)

	Gtk_range_get_inverted func(
		r *GtkRange) Gboolean

	Gtk_range_set_flippable func(
		r *GtkRange,
		flippable Gboolean)

	Gtk_range_get_flippable func(
		r *GtkRange) Gboolean

	Gtk_range_set_slider_size_fixed func(
		r *GtkRange,
		size_fixed Gboolean)

	Gtk_range_get_slider_size_fixed func(
		r *GtkRange) Gboolean

	Gtk_range_set_min_slider_size func(
		r *GtkRange,
		min_size Gboolean)

	Gtk_range_get_min_slider_size func(
		r *GtkRange) Gint

	Gtk_range_get_range_rect func(
		r *GtkRange,
		range_rect *GdkRectangle)

	Gtk_range_get_slider_range func(
		r *GtkRange,
		slider_start *Gint,
		slider_end *Gint)

	Gtk_range_set_lower_stepper_sensitivity func(
		r *GtkRange,
		sensitivity GtkSensitivityType)

	Gtk_range_get_lower_stepper_sensitivity func(
		r *GtkRange) GtkSensitivityType

	Gtk_range_set_upper_stepper_sensitivity func(
		r *GtkRange,
		sensitivity GtkSensitivityType)

	Gtk_range_get_upper_stepper_sensitivity func(
		r *GtkRange) GtkSensitivityType

	Gtk_range_set_increments func(
		r *GtkRange,
		step Gdouble,
		page Gdouble)

	Gtk_range_set_range func(
		r *GtkRange,
		min Gdouble,
		max Gdouble)

	Gtk_range_set_value func(
		r *GtkRange,
		value Gdouble)

	Gtk_range_get_value func(
		r *GtkRange) Gdouble

	Gtk_range_set_show_fill_level func(
		r *GtkRange,
		show_fill_level Gboolean)

	Gtk_range_get_show_fill_level func(
		r *GtkRange) Gboolean

	Gtk_range_set_restrict_to_fill_level func(
		r *GtkRange,
		restrict_to_fill_level Gboolean)

	Gtk_range_get_restrict_to_fill_level func(
		r *GtkRange) Gboolean

	Gtk_range_set_fill_level func(
		r *GtkRange,
		fill_level Gdouble)

	Gtk_range_get_fill_level func(
		r *GtkRange) Gdouble

	Gtk_range_set_round_digits func(
		r *GtkRange,
		round_digits Gint)

	Gtk_range_get_round_digits func(
		r *GtkRange) Gint

	Gtk_scale_get_type func() GType

	Gtk_scale_set_digits func(
		scale *GtkScale,
		digits Gint)

	Gtk_scale_get_digits func(
		scale *GtkScale) Gint

	Gtk_scale_set_draw_value func(
		scale *GtkScale,
		draw_value Gboolean)

	Gtk_scale_get_draw_value func(
		scale *GtkScale) Gboolean

	Gtk_scale_set_value_pos func(
		scale *GtkScale,
		pos GtkPositionType)

	Gtk_scale_get_value_pos func(
		scale *GtkScale) GtkPositionType

	Gtk_scale_get_layout func(
		scale *GtkScale) *PangoLayout

	Gtk_scale_get_layout_offsets func(
		scale *GtkScale,
		x *Gint,
		y *Gint)

	Gtk_scale_add_mark func(
		scale *GtkScale,
		value Gdouble,
		position GtkPositionType,
		markup string)

	Gtk_scale_clear_marks func(
		scale *GtkScale)

	Gtk_hscale_get_type func() GType

	Gtk_hscale_new func(
		adjustment *GtkAdjustment) *GtkWidget

	Gtk_hscale_new_with_range func(
		min Gdouble,
		max Gdouble,
		step Gdouble) *GtkWidget

	Gtk_scrollbar_get_type func() GType

	Gtk_hscrollbar_get_type func() GType

	Gtk_hscrollbar_new func(
		adjustment *GtkAdjustment) *GtkWidget

	Gtk_separator_get_type func() GType

	Gtk_hseparator_get_type func() GType

	Gtk_hseparator_new func() *GtkWidget

	Gtk_hsv_get_type func() GType

	Gtk_hsv_new func() *GtkWidget

	Gtk_hsv_set_color func(
		hsv *GtkHSV,
		h Double,
		s Double,
		v Double)

	Gtk_hsv_get_color func(
		hsv *GtkHSV,
		h *Gdouble,
		s *Gdouble,
		v *Gdouble)

	Gtk_hsv_set_metrics func(
		hsv *GtkHSV,
		size Gint,
		ring_width Gint)

	Gtk_hsv_get_metrics func(
		hsv *GtkHSV,
		size *Gint,
		ring_width *Gint)

	Gtk_hsv_is_adjusting func(
		hsv *GtkHSV) Gboolean

	Gtk_hsv_to_rgb func(
		h Gdouble,
		s Gdouble,
		v Gdouble,
		r *Gdouble,
		g *Gdouble,
		b *Gdouble)

	Gtk_rgb_to_hsv func(
		r Gdouble,
		g Gdouble,
		b Gdouble,
		h *Gdouble,
		s *Gdouble,
		v *Gdouble)

	Gtk_icon_factory_get_type func() GType

	Gtk_icon_factory_new func() *GtkIconFactory

	Gtk_icon_factory_add func(
		factory *GtkIconFactory,
		stock_id string,
		icon_set *GtkIconSet)

	Gtk_icon_factory_lookup func(
		factory *GtkIconFactory,
		stock_id string) *GtkIconSet

	Gtk_icon_factory_add_default func(
		factory *GtkIconFactory)

	Gtk_icon_factory_remove_default func(
		factory *GtkIconFactory)

	Gtk_icon_factory_lookup_default func(
		stock_id string) *GtkIconSet

	Gtk_icon_size_lookup func(
		size GtkIconSize,
		width *Gint,
		height *Gint) Gboolean

	Gtk_icon_size_lookup_for_settings func(
		settings *GtkSettings,
		size GtkIconSize,
		width *Gint,
		height *Gint) Gboolean

	Gtk_icon_size_register func(
		name string,
		width Gint,
		height Gint) GtkIconSize

	Gtk_icon_size_register_alias func(
		alias string,
		target GtkIconSize)

	Gtk_icon_size_from_name func(
		name string) GtkIconSize

	Gtk_icon_size_get_name func(
		size GtkIconSize) string

	Gtk_icon_set_get_type func() GType

	Gtk_icon_set_new func() *GtkIconSet

	Gtk_icon_set_new_from_pixbuf func(
		pixbuf *GdkPixbuf) *GtkIconSet

	Gtk_icon_set_ref func(
		icon_set *GtkIconSet) *GtkIconSet

	Gtk_icon_set_unref func(
		icon_set *GtkIconSet)

	Gtk_icon_set_copy func(
		icon_set *GtkIconSet) *GtkIconSet

	Gtk_icon_set_render_icon func(
		icon_set *GtkIconSet,
		style *GtkStyle,
		direction GtkTextDirection,
		state GtkStateType,
		size GtkIconSize,
		widget *GtkWidget,
		detail string) *GdkPixbuf

	Gtk_icon_set_add_source func(
		icon_set *GtkIconSet,
		source *GtkIconSource)

	Gtk_icon_set_get_sizes func(
		icon_set *GtkIconSet,
		sizes **GtkIconSize,
		n_sizes *Gint)

	Gtk_icon_source_get_type func() GType

	Gtk_icon_source_new func() *GtkIconSource

	Gtk_icon_source_copy func(
		source *GtkIconSource) *GtkIconSource

	Gtk_icon_source_free func(
		source *GtkIconSource)

	Gtk_icon_source_set_filename func(
		source *GtkIconSource,
		filename string)

	Gtk_icon_source_set_icon_name func(
		source *GtkIconSource,
		icon_name string)

	Gtk_icon_source_set_pixbuf func(
		source *GtkIconSource,
		pixbuf *GdkPixbuf)

	Gtk_icon_source_get_filename func(
		source *GtkIconSource) string

	Gtk_icon_source_get_icon_name func(
		source *GtkIconSource) string

	Gtk_icon_source_get_pixbuf func(
		source *GtkIconSource) *GdkPixbuf

	Gtk_icon_source_set_direction_wildcarded func(
		source *GtkIconSource,
		setting Gboolean)

	Gtk_icon_source_set_state_wildcarded func(
		source *GtkIconSource,
		setting Gboolean)

	Gtk_icon_source_set_size_wildcarded func(
		source *GtkIconSource,
		setting Gboolean)

	Gtk_icon_source_get_size_wildcarded func(
		source *GtkIconSource) Gboolean

	Gtk_icon_source_get_state_wildcarded func(
		source *GtkIconSource) Gboolean

	Gtk_icon_source_get_direction_wildcarded func(
		source *GtkIconSource) Gboolean

	Gtk_icon_source_set_direction func(
		source *GtkIconSource,
		direction GtkTextDirection)

	Gtk_icon_source_set_state func(
		source *GtkIconSource,
		state GtkStateType)

	Gtk_icon_source_set_size func(
		source *GtkIconSource,
		size GtkIconSize)

	Gtk_icon_source_get_direction func(
		source *GtkIconSource) GtkTextDirection

	Gtk_icon_source_get_state func(
		source *GtkIconSource) GtkStateType

	Gtk_icon_source_get_size func(
		source *GtkIconSource) GtkIconSize

	Gtk_icon_theme_error_quark func() GQuark

	Gtk_icon_theme_get_type func() GType

	Gtk_icon_theme_new func() *GtkIconTheme

	Gtk_icon_theme_get_default func() *GtkIconTheme

	Gtk_icon_theme_get_for_screen func(
		screen *GdkScreen) *GtkIconTheme

	Gtk_icon_theme_set_screen func(
		icon_theme *GtkIconTheme,
		screen *GdkScreen)

	Gtk_icon_theme_set_search_path func(
		icon_theme *GtkIconTheme,
		path **Gchar,
		n_elements Gint)

	Gtk_icon_theme_get_search_path func(
		icon_theme *GtkIconTheme,
		path ***Gchar,
		n_elements *Gint)

	Gtk_icon_theme_append_search_path func(
		icon_theme *GtkIconTheme,
		path string)

	Gtk_icon_theme_prepend_search_path func(
		icon_theme *GtkIconTheme,
		path string)

	Gtk_icon_theme_set_custom_theme func(
		icon_theme *GtkIconTheme,
		theme_name string)

	Gtk_icon_theme_has_icon func(
		icon_theme *GtkIconTheme,
		icon_name string) Gboolean

	Gtk_icon_theme_get_icon_sizes func(
		icon_theme *GtkIconTheme,
		icon_name string) *Gint

	Gtk_icon_theme_lookup_icon func(
		icon_theme *GtkIconTheme,
		icon_name string,
		size Gint,
		flags GtkIconLookupFlags) *GtkIconInfo

	Gtk_icon_theme_choose_icon func(
		icon_theme *GtkIconTheme,
		icon_names **Gchar,
		size Gint,
		flags GtkIconLookupFlags) *GtkIconInfo

	Gtk_icon_theme_load_icon func(
		icon_theme *GtkIconTheme,
		icon_name string,
		size Gint,
		flags GtkIconLookupFlags,
		error **GError) *GdkPixbuf

	Gtk_icon_theme_lookup_by_gicon func(
		icon_theme *GtkIconTheme,
		icon *GIcon,
		size Gint,
		flags GtkIconLookupFlags) *GtkIconInfo

	Gtk_icon_theme_list_icons func(
		icon_theme *GtkIconTheme,
		context string) *GList

	Gtk_icon_theme_list_contexts func(
		icon_theme *GtkIconTheme) *GList

	Gtk_icon_theme_get_example_icon_name func(
		icon_theme *GtkIconTheme) string

	Gtk_icon_theme_rescan_if_needed func(
		icon_theme *GtkIconTheme) Gboolean

	Gtk_icon_theme_add_builtin_icon func(
		icon_name string,
		size Gint,
		pixbuf *GdkPixbuf)

	Gtk_icon_info_get_type func() GType

	Gtk_icon_info_copy func(
		icon_info *GtkIconInfo) *GtkIconInfo

	Gtk_icon_info_free func(
		icon_info *GtkIconInfo)

	Gtk_icon_info_new_for_pixbuf func(
		icon_theme *GtkIconTheme,
		pixbuf *GdkPixbuf) *GtkIconInfo

	Gtk_icon_info_get_base_size func(
		icon_info *GtkIconInfo) Gint

	Gtk_icon_info_get_filename func(
		icon_info *GtkIconInfo) string

	Gtk_icon_info_get_builtin_pixbuf func(
		icon_info *GtkIconInfo) *GdkPixbuf

	Gtk_icon_info_load_icon func(
		icon_info *GtkIconInfo,
		error **GError) *GdkPixbuf

	Gtk_icon_info_set_raw_coordinates func(
		icon_info *GtkIconInfo,
		raw_coordinates Gboolean)

	Gtk_icon_info_get_embedded_rect func(
		icon_info *GtkIconInfo,
		rectangle *GdkRectangle) Gboolean

	Gtk_icon_info_get_attach_points func(
		icon_info *GtkIconInfo,
		points **GdkPoint,
		n_points *Gint) Gboolean

	Gtk_icon_info_get_display_name func(
		icon_info *GtkIconInfo) string

	Gtk_tooltip_get_type func() GType

	Gtk_tooltip_set_markup func(
		tooltip *GtkTooltip,
		markup string)

	Gtk_tooltip_set_text func(
		tooltip *GtkTooltip,
		text string)

	Gtk_tooltip_set_icon func(
		tooltip *GtkTooltip,
		pixbuf *GdkPixbuf)

	Gtk_tooltip_set_icon_from_stock func(
		tooltip *GtkTooltip,
		stock_id string,
		size GtkIconSize)

	Gtk_tooltip_set_icon_from_icon_name func(
		tooltip *GtkTooltip,
		icon_name string,
		size GtkIconSize)

	Gtk_tooltip_set_icon_from_gicon func(
		tooltip *GtkTooltip,
		gicon *GIcon,
		size GtkIconSize)

	Gtk_tooltip_set_custom func(
		tooltip *GtkTooltip,
		custom_widget *GtkWidget)

	Gtk_tooltip_set_tip_area func(
		tooltip *GtkTooltip,
		rect *GdkRectangle)

	Gtk_tooltip_trigger_tooltip_query func(
		display *GdkDisplay)

	Gtk_icon_view_get_type func() GType

	Gtk_icon_view_new func() *GtkWidget

	Gtk_icon_view_new_with_model func(
		model *GtkTreeModel) *GtkWidget

	Gtk_icon_view_set_model func(
		icon_view *GtkIconView,
		model *GtkTreeModel)

	Gtk_icon_view_get_model func(
		icon_view *GtkIconView) *GtkTreeModel

	Gtk_icon_view_set_text_column func(
		icon_view *GtkIconView,
		column Gint)

	Gtk_icon_view_get_text_column func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_markup_column func(
		icon_view *GtkIconView,
		column Gint)

	Gtk_icon_view_get_markup_column func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_pixbuf_column func(
		icon_view *GtkIconView,
		column Gint)

	Gtk_icon_view_get_pixbuf_column func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_orientation func(
		icon_view *GtkIconView,
		orientation GtkOrientation)

	Gtk_icon_view_get_orientation func(
		icon_view *GtkIconView) GtkOrientation

	Gtk_icon_view_set_item_orientation func(
		icon_view *GtkIconView,
		orientation GtkOrientation)

	Gtk_icon_view_get_item_orientation func(
		icon_view *GtkIconView) GtkOrientation

	Gtk_icon_view_set_columns func(
		icon_view *GtkIconView,
		columns Gint)

	Gtk_icon_view_get_columns func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_item_width func(
		icon_view *GtkIconView,
		item_width Gint)

	Gtk_icon_view_get_item_width func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_spacing func(
		icon_view *GtkIconView,
		spacing Gint)

	Gtk_icon_view_get_spacing func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_row_spacing func(
		icon_view *GtkIconView,
		row_spacing Gint)

	Gtk_icon_view_get_row_spacing func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_column_spacing func(
		icon_view *GtkIconView,
		column_spacing Gint)

	Gtk_icon_view_get_column_spacing func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_margin func(
		icon_view *GtkIconView,
		margin Gint)

	Gtk_icon_view_get_margin func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_set_item_padding func(
		icon_view *GtkIconView,
		item_padding Gint)

	Gtk_icon_view_get_item_padding func(
		icon_view *GtkIconView) Gint

	Gtk_icon_view_get_path_at_pos func(
		icon_view *GtkIconView,
		x Gint,
		y Gint) *GtkTreePath

	Gtk_icon_view_get_item_at_pos func(
		icon_view *GtkIconView,
		x Gint,
		y Gint,
		path **GtkTreePath,
		cell **GtkCellRenderer) Gboolean

	Gtk_icon_view_get_visible_range func(
		icon_view *GtkIconView,
		start_path **GtkTreePath,
		end_path **GtkTreePath) Gboolean

	Gtk_icon_view_selected_foreach func(
		icon_view *GtkIconView,
		f GtkIconViewForeachFunc,
		data Gpointer)

	Gtk_icon_view_set_selection_mode func(
		icon_view *GtkIconView,
		mode GtkSelectionMode)

	Gtk_icon_view_get_selection_mode func(
		icon_view *GtkIconView) GtkSelectionMode

	Gtk_icon_view_select_path func(
		icon_view *GtkIconView,
		path *GtkTreePath)

	Gtk_icon_view_unselect_path func(
		icon_view *GtkIconView,
		path *GtkTreePath)

	Gtk_icon_view_path_is_selected func(
		icon_view *GtkIconView,
		path *GtkTreePath) Gboolean

	Gtk_icon_view_get_item_row func(
		icon_view *GtkIconView,
		path *GtkTreePath) Gint

	Gtk_icon_view_get_item_column func(
		icon_view *GtkIconView,
		path *GtkTreePath) Gint

	Gtk_icon_view_get_selected_items func(
		icon_view *GtkIconView) *GList

	Gtk_icon_view_select_all func(
		icon_view *GtkIconView)

	Gtk_icon_view_unselect_all func(
		icon_view *GtkIconView)

	Gtk_icon_view_item_activated func(
		icon_view *GtkIconView,
		path *GtkTreePath)

	Gtk_icon_view_set_cursor func(
		icon_view *GtkIconView,
		path *GtkTreePath,
		cell *GtkCellRenderer,
		start_editing Gboolean)

	Gtk_icon_view_get_cursor func(
		icon_view *GtkIconView,
		path **GtkTreePath,
		cell **GtkCellRenderer) Gboolean

	Gtk_icon_view_scroll_to_path func(
		icon_view *GtkIconView,
		path *GtkTreePath,
		use_align Gboolean,
		row_align Gfloat,
		col_align Gfloat)

	Gtk_icon_view_enable_model_drag_source func(
		icon_view *GtkIconView,
		start_button_mask GdkModifierType,
		targets *GtkTargetEntry,
		n_targets Gint,
		actions GdkDragAction)

	Gtk_icon_view_enable_model_drag_dest func(
		icon_view *GtkIconView,
		targets *GtkTargetEntry,
		n_targets Gint,
		actions GdkDragAction)

	Gtk_icon_view_unset_model_drag_source func(
		icon_view *GtkIconView)

	Gtk_icon_view_unset_model_drag_dest func(
		icon_view *GtkIconView)

	Gtk_icon_view_set_reorderable func(
		icon_view *GtkIconView,
		reorderable Gboolean)

	Gtk_icon_view_get_reorderable func(
		icon_view *GtkIconView) Gboolean

	Gtk_icon_view_set_drag_dest_item func(
		icon_view *GtkIconView,
		path *GtkTreePath,
		pos GtkIconViewDropPosition)

	Gtk_icon_view_get_drag_dest_item func(
		icon_view *GtkIconView,
		path **GtkTreePath,
		pos *GtkIconViewDropPosition)

	Gtk_icon_view_get_dest_item_at_pos func(
		icon_view *GtkIconView,
		drag_x Gint,
		drag_y Gint,
		path **GtkTreePath,
		pos *GtkIconViewDropPosition) Gboolean

	Gtk_icon_view_create_drag_icon func(
		icon_view *GtkIconView,
		path *GtkTreePath) *GdkPixmap

	Gtk_icon_view_convert_widget_to_bin_window_coords func(
		icon_view *GtkIconView,
		wx Gint,
		wy Gint,
		bx *Gint,
		by *Gint)

	Gtk_icon_view_set_tooltip_item func(
		icon_view *GtkIconView,
		tooltip *GtkTooltip,
		path *GtkTreePath)

	Gtk_icon_view_set_tooltip_cell func(
		icon_view *GtkIconView,
		tooltip *GtkTooltip,
		path *GtkTreePath,
		cell *GtkCellRenderer)

	Gtk_icon_view_get_tooltip_context func(
		icon_view *GtkIconView,
		x *Gint,
		y *Gint,
		keyboard_tip Gboolean,
		model **GtkTreeModel,
		path **GtkTreePath,
		iter *GtkTreeIter) Gboolean

	Gtk_icon_view_set_tooltip_column func(
		icon_view *GtkIconView,
		column Gint)

	Gtk_icon_view_get_tooltip_column func(
		icon_view *GtkIconView) Gint

	Gtk_image_menu_item_get_type func() GType

	Gtk_image_menu_item_new func() *GtkWidget

	Gtk_image_menu_item_new_with_label func(
		label string) *GtkWidget

	Gtk_image_menu_item_new_with_mnemonic func(
		label string) *GtkWidget

	Gtk_image_menu_item_new_from_stock func(
		stock_id string,
		accel_group *GtkAccelGroup) *GtkWidget

	Gtk_image_menu_item_set_always_show_image func(
		image_menu_item *GtkImageMenuItem,
		always_show Gboolean)

	Gtk_image_menu_item_get_always_show_image func(
		image_menu_item *GtkImageMenuItem) Gboolean

	Gtk_image_menu_item_set_image func(
		image_menu_item *GtkImageMenuItem,
		image *GtkWidget)

	Gtk_image_menu_item_get_image func(
		image_menu_item *GtkImageMenuItem) *GtkWidget

	Gtk_image_menu_item_set_use_stock func(
		image_menu_item *GtkImageMenuItem,
		use_stock Gboolean)

	Gtk_image_menu_item_get_use_stock func(
		image_menu_item *GtkImageMenuItem) Gboolean

	Gtk_image_menu_item_set_accel_group func(
		image_menu_item *GtkImageMenuItem,
		accel_group *GtkAccelGroup)

	Gtk_im_context_simple_get_type func() GType

	Gtk_im_context_simple_new func() *GtkIMContext

	Gtk_im_context_simple_add_table func(
		context_simple *GtkIMContextSimple,
		data *Guint16,
		max_seq_len Gint,
		n_seqs Gint)

	Gtk_im_multicontext_get_type func() GType

	Gtk_im_multicontext_new func() *GtkIMContext

	Gtk_im_multicontext_append_menuitems func(
		context *GtkIMMulticontext,
		menushell *GtkMenuShell)

	Gtk_im_multicontext_get_context_id func(
		context *GtkIMMulticontext) string

	Gtk_im_multicontext_set_context_id func(
		context *GtkIMMulticontext,
		context_id string)

	Gtk_info_bar_get_type func() GType

	Gtk_info_bar_new func() *GtkWidget

	Gtk_info_bar_new_with_buttons func(
		first_button_text string, v ...VArg) *GtkWidget

	Gtk_info_bar_get_action_area func(
		info_bar *GtkInfoBar) *GtkWidget

	Gtk_info_bar_get_content_area func(
		info_bar *GtkInfoBar) *GtkWidget

	Gtk_info_bar_add_action_widget func(
		info_bar *GtkInfoBar,
		child *GtkWidget,
		response_id Gint)

	Gtk_info_bar_add_button func(
		info_bar *GtkInfoBar,
		button_text string,
		response_id Gint) *GtkWidget

	Gtk_info_bar_add_buttons func(info_bar *GtkInfoBar,
		first_button_text string, v ...VArg)

	Gtk_info_bar_set_response_sensitive func(
		info_bar *GtkInfoBar,
		response_id Gint,
		setting Gboolean)

	Gtk_info_bar_set_default_response func(
		info_bar *GtkInfoBar,
		response_id Gint)

	Gtk_info_bar_response func(
		info_bar *GtkInfoBar,
		response_id Gint)

	Gtk_info_bar_set_message_type func(
		info_bar *GtkInfoBar,
		message_type GtkMessageType)

	Gtk_info_bar_get_message_type func(
		info_bar *GtkInfoBar) GtkMessageType

	Gtk_invisible_get_type func() GType

	Gtk_invisible_new func() *GtkWidget

	Gtk_invisible_new_for_screen func(
		screen *GdkScreen) *GtkWidget

	Gtk_invisible_set_screen func(
		invisible *GtkInvisible,
		screen *GdkScreen)

	Gtk_invisible_get_screen func(
		invisible *GtkInvisible) *GdkScreen

	Gtk_layout_get_type func() GType

	Gtk_layout_new func(
		hadjustment *GtkAdjustment,
		vadjustment *GtkAdjustment) *GtkWidget

	Gtk_layout_get_bin_window func(
		layout *GtkLayout) *GdkWindow

	Gtk_layout_put func(
		layout *GtkLayout,
		child_widget *GtkWidget,
		x Gint,
		y Gint)

	Gtk_layout_move func(
		layout *GtkLayout,
		child_widget *GtkWidget,
		x Gint,
		y Gint)

	Gtk_layout_set_size func(
		layout *GtkLayout,
		width Guint,
		height Guint)

	Gtk_layout_get_size func(
		layout *GtkLayout,
		width *Guint,
		height *Guint)

	Gtk_layout_get_hadjustment func(
		layout *GtkLayout) *GtkAdjustment

	Gtk_layout_get_vadjustment func(
		layout *GtkLayout) *GtkAdjustment

	Gtk_layout_set_hadjustment func(
		layout *GtkLayout,
		adjustment *GtkAdjustment)

	Gtk_layout_set_vadjustment func(
		layout *GtkLayout,
		adjustment *GtkAdjustment)

	Gtk_layout_freeze func(
		layout *GtkLayout)

	Gtk_layout_thaw func(
		layout *GtkLayout)

	Gtk_link_button_get_type func() GType

	Gtk_link_button_new func(
		uri string) *GtkWidget

	Gtk_link_button_new_with_label func(
		uri string,
		label string) *GtkWidget

	Gtk_link_button_get_uri func(
		link_button *GtkLinkButton) string

	Gtk_link_button_set_uri func(
		link_button *GtkLinkButton,
		uri string)

	Gtk_link_button_set_uri_hook func(
		f GtkLinkButtonUriFunc,
		dataGpointer,
		destroy GDestroyNotify) GtkLinkButtonUriFunc

	Gtk_link_button_get_visited func(
		link_button *GtkLinkButton) Gboolean

	Gtk_link_button_set_visited func(
		link_button *GtkLinkButton,
		visited Gboolean)

	Gtk_check_version func(
		required_major Guint,
		required_minor Guint,
		required_micro Guint) string

	Gtk_parse_args func(
		argc *int,
		argv ***Char) Gboolean

	Gtk_init func(
		argc *int,
		argv ***Char)

	Gtk_init_check func(
		argc *int,
		argv ***Char) Gboolean

	Gtk_init_with_args func(
		argc *int,
		argv ***Char,
		parameter_string string,
		entries *GOptionEntry,
		translation_domain string,
		error **GError) Gboolean

	Gtk_get_option_group func(
		open_default_display Gboolean) *GOptionGroup

	Gtk_init_abi_check func(
		argc *int,
		argv ***Char,
		num_checks int,
		sizeof_GtkWindow Size_t,
		sizeof_GtkBox Size_t)

	Gtk_init_check_abi_check func(
		argc *int,
		argv ***Char,
		num_checks int,
		sizeof_GtkWindow Size_t,
		sizeof_GtkBox Size_t) Gboolean

	Gtk_exit func(
		error_code Gint)

	Gtk_set_locale func() string

	Gtk_disable_setlocale func()

	Gtk_get_default_language func() *PangoLanguage

	Gtk_events_pending func() Gboolean

	Gtk_main_do_event func(
		event *GdkEvent)

	Gtk_main func()

	Gtk_main_level func() Guint

	Gtk_main_quit func()

	Gtk_main_iteration func() Gboolean

	Gtk_main_iteration_do func(
		blocking Gboolean) Gboolean

	Gtk_true func() Gboolean

	Gtk_false func() Gboolean

	Gtk_grab_add func(
		widget *GtkWidget)

	Gtk_grab_get_current func() *GtkWidget

	Gtk_grab_remove func(
		widget *GtkWidget)

	Gtk_init_add func(
		function GtkFunction,
		data Gpointer)

	Gtk_quit_add_destroy func(
		main_level Guint,
		object *GtkObject)

	Gtk_quit_add func(
		main_level Guint,
		function GtkFunction,
		data Gpointer) Guint

	Gtk_quit_add_full func(
		main_level Guint,
		function GtkFunction,
		marshal GtkCallbackMarshal,
		dataGpointer,
		destroy GDestroyNotify) Guint

	Gtk_quit_remove func(
		quit_handler_id Guint)

	Gtk_quit_remove_by_data func(
		data Gpointer)

	Gtk_timeout_add func(
		interval Guint32,
		function GtkFunction,
		data Gpointer) Guint

	Gtk_timeout_add_full func(
		interval Guint32,
		function GtkFunction,
		marshal GtkCallbackMarshal,
		dataGpointer,
		destroy GDestroyNotify) Guint

	Gtk_timeout_remove func(
		timeout_handler_id Guint)

	Gtk_idle_add func(
		function GtkFunction,
		data Gpointer) Guint

	Gtk_idle_add_priority func(
		priority Gint,
		function GtkFunction,
		data Gpointer) Guint

	Gtk_idle_add_full func(
		priority Gint,
		function GtkFunction,
		marshal GtkCallbackMarshal,
		dataGpointer,
		destroy GDestroyNotify) Guint

	Gtk_idle_remove func(
		idle_handler_id Guint)

	Gtk_idle_remove_by_data func(
		data Gpointer)

	Gtk_input_add_full func(
		source Gint,
		condition GdkInputCondition,
		function GdkInputFunction,
		marshal GtkCallbackMarshal,
		dataGpointer,
		destroy GDestroyNotify) Guint

	Gtk_input_remove func(
		input_handler_id Guint)

	Gtk_key_snooper_install func(
		snooper GtkKeySnoopFunc,
		func_data Gpointer) Guint

	Gtk_key_snooper_remove func(
		snooper_handler_id Guint)

	Gtk_get_current_event func() *GdkEvent

	Gtk_get_current_event_time func() Guint32

	Gtk_get_current_event_state func(
		state *GdkModifierType) Gboolean

	Gtk_get_event_widget func(
		event *GdkEvent) *GtkWidget

	Gtk_propagate_event func(
		widget *GtkWidget,
		event *GdkEvent)

	Gtk_menu_bar_get_type func() GType

	Gtk_menu_bar_new func() *GtkWidget

	Gtk_menu_bar_get_pack_direction func(
		menubar *GtkMenuBar) GtkPackDirection

	Gtk_menu_bar_set_pack_direction func(
		menubar *GtkMenuBar,
		pack_dir GtkPackDirection)

	Gtk_menu_bar_get_child_pack_direction func(
		menubar *GtkMenuBar) GtkPackDirection

	Gtk_menu_bar_set_child_pack_direction func(
		menubar *GtkMenuBar,
		child_pack_dir GtkPackDirection)

	Gtk_tooltips_get_type func() GType

	Gtk_tooltips_new func() *GtkTooltips

	Gtk_tooltips_enable func(
		tooltips *GtkTooltips)

	Gtk_tooltips_disable func(
		tooltips *GtkTooltips)

	Gtk_tooltips_set_delay func(
		tooltips *GtkTooltips,
		delay Guint)

	Gtk_tooltips_set_tip func(
		tooltips *GtkTooltips,
		widget *GtkWidget,
		tip_text string,
		tip_private string)

	Gtk_tooltips_data_get func(
		widget *GtkWidget) *GtkTooltipsData

	Gtk_tooltips_force_window func(
		tooltips *GtkTooltips)

	Gtk_tooltips_get_info_from_tip_window func(
		tip_window *GtkWindow,
		tooltips **GtkTooltips,
		current_widget **GtkWidget) Gboolean

	Gtk_size_group_get_type func() GType

	Gtk_size_group_new func(
		mode GtkSizeGroupMode) *GtkSizeGroup

	Gtk_size_group_set_mode func(
		size_group *GtkSizeGroup,
		mode GtkSizeGroupMode)

	Gtk_size_group_get_mode func(
		size_group *GtkSizeGroup) GtkSizeGroupMode

	Gtk_size_group_set_ignore_hidden func(
		size_group *GtkSizeGroup,
		ignore_hidden Gboolean)

	Gtk_size_group_get_ignore_hidden func(
		size_group *GtkSizeGroup) Gboolean

	Gtk_size_group_add_widget func(
		size_group *GtkSizeGroup,
		widget *GtkWidget)

	Gtk_size_group_remove_widget func(
		size_group *GtkSizeGroup,
		widget *GtkWidget)

	Gtk_size_group_get_widgets func(
		size_group *GtkSizeGroup) *GSList

	Gtk_tool_item_get_type func() GType

	Gtk_tool_item_new func() *GtkToolItem

	Gtk_tool_item_set_homogeneous func(
		tool_item *GtkToolItem,
		homogeneous Gboolean)

	Gtk_tool_item_get_homogeneous func(
		tool_item *GtkToolItem) Gboolean

	Gtk_tool_item_set_expand func(
		tool_item *GtkToolItem,
		expand Gboolean)

	Gtk_tool_item_get_expand func(
		tool_item *GtkToolItem) Gboolean

	Gtk_tool_item_set_tooltip func(
		tool_item *GtkToolItem,
		tooltips *GtkTooltips,
		tip_text string,
		tip_private string)

	Gtk_tool_item_set_tooltip_text func(
		tool_item *GtkToolItem,
		text string)

	Gtk_tool_item_set_tooltip_markup func(
		tool_item *GtkToolItem,
		markup string)

	Gtk_tool_item_set_use_drag_window func(
		tool_item *GtkToolItem,
		use_drag_window Gboolean)

	Gtk_tool_item_get_use_drag_window func(
		tool_item *GtkToolItem) Gboolean

	Gtk_tool_item_set_visible_horizontal func(
		tool_item *GtkToolItem,
		visible_horizontal Gboolean)

	Gtk_tool_item_get_visible_horizontal func(
		tool_item *GtkToolItem) Gboolean

	Gtk_tool_item_set_visible_vertical func(
		tool_item *GtkToolItem,
		visible_vertical Gboolean)

	Gtk_tool_item_get_visible_vertical func(
		tool_item *GtkToolItem) Gboolean

	Gtk_tool_item_get_is_important func(
		tool_item *GtkToolItem) Gboolean

	Gtk_tool_item_set_is_important func(
		tool_item *GtkToolItem,
		is_important Gboolean)

	Gtk_tool_item_get_ellipsize_mode func(
		tool_item *GtkToolItem) PangoEllipsizeMode

	Gtk_tool_item_get_icon_size func(
		tool_item *GtkToolItem) GtkIconSize

	Gtk_tool_item_get_orientation func(
		tool_item *GtkToolItem) GtkOrientation

	Gtk_tool_item_get_toolbar_style func(
		tool_item *GtkToolItem) GtkToolbarStyle

	Gtk_tool_item_get_relief_style func(
		tool_item *GtkToolItem) GtkReliefStyle

	Gtk_tool_item_get_text_alignment func(
		tool_item *GtkToolItem) Gfloat

	Gtk_tool_item_get_text_orientation func(
		tool_item *GtkToolItem) GtkOrientation

	Gtk_tool_item_get_text_size_group func(
		tool_item *GtkToolItem) *GtkSizeGroup

	Gtk_tool_item_retrieve_proxy_menu_item func(
		tool_item *GtkToolItem) *GtkWidget

	Gtk_tool_item_get_proxy_menu_item func(
		tool_item *GtkToolItem,
		menu_item_id string) *GtkWidget

	Gtk_tool_item_set_proxy_menu_item func(
		tool_item *GtkToolItem,
		menu_item_id string,
		menu_item *GtkWidget)

	Gtk_tool_item_rebuild_menu func(
		tool_item *GtkToolItem)

	Gtk_tool_item_toolbar_reconfigured func(
		tool_item *GtkToolItem)

	Gtk_tool_button_get_type func() GType

	Gtk_tool_button_new func(
		icon_widget *GtkWidget,
		label string) *GtkToolItem

	Gtk_tool_button_new_from_stock func(
		stock_id string) *GtkToolItem

	Gtk_tool_button_set_label func(
		button *GtkToolButton,
		label string)

	Gtk_tool_button_get_label func(
		button *GtkToolButton) string

	Gtk_tool_button_set_use_underline func(
		button *GtkToolButton,
		use_underline Gboolean)

	Gtk_tool_button_get_use_underline func(
		button *GtkToolButton) Gboolean

	Gtk_tool_button_set_stock_id func(
		button *GtkToolButton,
		stock_id string)

	Gtk_tool_button_get_stock_id func(
		button *GtkToolButton) string

	Gtk_tool_button_set_icon_name func(
		button *GtkToolButton,
		icon_name string)

	Gtk_tool_button_get_icon_name func(
		button *GtkToolButton) string

	Gtk_tool_button_set_icon_widget func(
		button *GtkToolButton,
		icon_widget *GtkWidget)

	Gtk_tool_button_get_icon_widget func(
		button *GtkToolButton) *GtkWidget

	Gtk_tool_button_set_label_widget func(
		button *GtkToolButton,
		label_widget *GtkWidget)

	Gtk_tool_button_get_label_widget func(
		button *GtkToolButton) *GtkWidget

	Gtk_menu_tool_button_get_type func() GType

	Gtk_menu_tool_button_new func(
		icon_widget *GtkWidget,
		label string) *GtkToolItem

	Gtk_menu_tool_button_new_from_stock func(
		stock_id string) *GtkToolItem

	Gtk_menu_tool_button_set_menu func(
		button *GtkMenuToolButton,
		menu *GtkWidget)

	Gtk_menu_tool_button_get_menu func(
		button *GtkMenuToolButton) *GtkWidget

	Gtk_menu_tool_button_set_arrow_tooltip func(
		button *GtkMenuToolButton,
		tooltips *GtkTooltips,
		tip_text string,
		tip_private string)

	Gtk_menu_tool_button_set_arrow_tooltip_text func(
		button *GtkMenuToolButton,
		text string)

	Gtk_menu_tool_button_set_arrow_tooltip_markup func(
		button *GtkMenuToolButton,
		markup string)

	Gtk_message_dialog_get_type func() GType

	Gtk_message_dialog_new func(
		parent *GtkWindow,
		flags GtkDialogFlags,
		t GtkMessageType,
		buttons GtkButtonsType,
		message_format string,
		v ...VArg) *GtkWidget

	Gtk_message_dialog_new_with_markup func(
		parent *GtkWindow,
		flags GtkDialogFlags,
		t GtkMessageType,
		buttons GtkButtonsType,
		message_format string,
		v ...VArg) *GtkWidget

	Gtk_message_dialog_set_image func(
		dialog *GtkMessageDialog,
		image *GtkWidget)

	Gtk_message_dialog_get_image func(
		dialog *GtkMessageDialog) *GtkWidget

	Gtk_message_dialog_set_markup func(
		message_dialog *GtkMessageDialog,
		str string)

	Gtk_message_dialog_format_secondary_text func(
		message_dialog *GtkMessageDialog,
		message_format string,
		v ...VArg)

	Gtk_message_dialog_format_secondary_markup func(
		message_dialog *GtkMessageDialog,
		message_format string,
		v ...VArg)

	Gtk_message_dialog_get_message_area func(
		message_dialog *GtkMessageDialog) *GtkWidget

	Gtk_mount_operation_get_type func() GType

	Gtk_mount_operation_new func(
		parent *GtkWindow) *GMountOperation

	Gtk_mount_operation_is_showing func(
		op *GtkMountOperation) Gboolean

	Gtk_mount_operation_set_parent func(
		op *GtkMountOperation,
		parent *GtkWindow)

	Gtk_mount_operation_get_parent func(
		op *GtkMountOperation) *GtkWindow

	Gtk_mount_operation_set_screen func(
		op *GtkMountOperation,
		screen *GdkScreen)

	Gtk_mount_operation_get_screen func(
		op *GtkMountOperation) *GdkScreen

	Gtk_notebook_get_type func() GType

	Gtk_notebook_new func() *GtkWidget

	Gtk_notebook_append_page func(
		notebook *GtkNotebook,
		child *GtkWidget,
		tab_label *GtkWidget) Gint

	Gtk_notebook_append_page_menu func(
		notebook *GtkNotebook,
		child *GtkWidget,
		tab_label *GtkWidget,
		menu_label *GtkWidget) Gint

	Gtk_notebook_prepend_page func(
		notebook *GtkNotebook,
		child *GtkWidget,
		tab_label *GtkWidget) Gint

	Gtk_notebook_prepend_page_menu func(
		notebook *GtkNotebook,
		child *GtkWidget,
		tab_label *GtkWidget,
		menu_label *GtkWidget) Gint

	Gtk_notebook_insert_page func(
		notebook *GtkNotebook,
		child *GtkWidget,
		tab_label *GtkWidget,
		position Gint) Gint

	Gtk_notebook_insert_page_menu func(
		notebook *GtkNotebook,
		child *GtkWidget,
		tab_label *GtkWidget,
		menu_label *GtkWidget,
		position Gint) Gint

	Gtk_notebook_remove_page func(
		notebook *GtkNotebook,
		page_num Gint)

	Gtk_notebook_set_window_creation_hook func(
		f GtkNotebookWindowCreationFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_notebook_set_group_id func(
		notebook *GtkNotebook,
		group_id Gint)

	Gtk_notebook_get_group_id func(
		notebook *GtkNotebook) Gint

	Gtk_notebook_set_group func(
		notebook *GtkNotebook,
		group Gpointer)

	Gtk_notebook_get_group func(
		notebook *GtkNotebook) Gpointer

	Gtk_notebook_set_group_name func(
		notebook *GtkNotebook,
		group_name string)

	Gtk_notebook_get_group_name func(
		notebook *GtkNotebook) string

	Gtk_notebook_get_current_page func(
		notebook *GtkNotebook) Gint

	Gtk_notebook_get_nth_page func(
		notebook *GtkNotebook,
		page_num Gint) *GtkWidget

	Gtk_notebook_get_n_pages func(
		notebook *GtkNotebook) Gint

	Gtk_notebook_page_num func(
		notebook *GtkNotebook,
		child *GtkWidget) Gint

	Gtk_notebook_set_current_page func(
		notebook *GtkNotebook,
		page_num Gint)

	Gtk_notebook_next_page func(
		notebook *GtkNotebook)

	Gtk_notebook_prev_page func(
		notebook *GtkNotebook)

	Gtk_notebook_set_show_border func(
		notebook *GtkNotebook,
		show_border Gboolean)

	Gtk_notebook_get_show_border func(
		notebook *GtkNotebook) Gboolean

	Gtk_notebook_set_show_tabs func(
		notebook *GtkNotebook,
		show_tabs Gboolean)

	Gtk_notebook_get_show_tabs func(
		notebook *GtkNotebook) Gboolean

	Gtk_notebook_set_tab_pos func(
		notebook *GtkNotebook,
		pos GtkPositionType)

	Gtk_notebook_get_tab_pos func(
		notebook *GtkNotebook) GtkPositionType

	Gtk_notebook_set_homogeneous_tabs func(
		notebook *GtkNotebook,
		homogeneous Gboolean)

	Gtk_notebook_set_tab_border func(
		notebook *GtkNotebook,
		border_width Guint)

	Gtk_notebook_set_tab_hborder func(
		notebook *GtkNotebook,
		tab_hborder Guint)

	Gtk_notebook_set_tab_vborder func(
		notebook *GtkNotebook,
		tab_vborder Guint)

	Gtk_notebook_set_scrollable func(
		notebook *GtkNotebook,
		scrollable Gboolean)

	Gtk_notebook_get_scrollable func(
		notebook *GtkNotebook) Gboolean

	Gtk_notebook_get_tab_hborder func(
		notebook *GtkNotebook) Guint16

	Gtk_notebook_get_tab_vborder func(
		notebook *GtkNotebook) Guint16

	Gtk_notebook_popup_enable func(
		notebook *GtkNotebook)

	Gtk_notebook_popup_disable func(
		notebook *GtkNotebook)

	Gtk_notebook_get_tab_label func(
		notebook *GtkNotebook,
		child *GtkWidget) *GtkWidget

	Gtk_notebook_set_tab_label func(
		notebook *GtkNotebook,
		child *GtkWidget,
		tab_label *GtkWidget)

	Gtk_notebook_set_tab_label_text func(
		notebook *GtkNotebook,
		child *GtkWidget,
		tab_text string)

	Gtk_notebook_get_tab_label_text func(
		notebook *GtkNotebook,
		child *GtkWidget) string

	Gtk_notebook_get_menu_label func(
		notebook *GtkNotebook,
		child *GtkWidget) *GtkWidget

	Gtk_notebook_set_menu_label func(
		notebook *GtkNotebook,
		child *GtkWidget,
		menu_label *GtkWidget)

	Gtk_notebook_set_menu_label_text func(
		notebook *GtkNotebook,
		child *GtkWidget,
		menu_text string)

	Gtk_notebook_get_menu_label_text func(
		notebook *GtkNotebook,
		child *GtkWidget) string

	Gtk_notebook_query_tab_label_packing func(
		notebook *GtkNotebook,
		child *GtkWidget,
		expand *Gboolean,
		fill *Gboolean,
		pack_type *GtkPackType)

	Gtk_notebook_set_tab_label_packing func(
		notebook *GtkNotebook,
		child *GtkWidget,
		expand Gboolean,
		fill Gboolean,
		pack_type GtkPackType)

	Gtk_notebook_reorder_child func(
		notebook *GtkNotebook,
		child *GtkWidget,
		position Gint)

	Gtk_notebook_get_tab_reorderable func(
		notebook *GtkNotebook,
		child *GtkWidget) Gboolean

	Gtk_notebook_set_tab_reorderable func(
		notebook *GtkNotebook,
		child *GtkWidget,
		reorderable Gboolean)

	Gtk_notebook_get_tab_detachable func(
		notebook *GtkNotebook,
		child *GtkWidget) Gboolean

	Gtk_notebook_set_tab_detachable func(
		notebook *GtkNotebook,
		child *GtkWidget,
		detachable Gboolean)

	Gtk_notebook_get_action_widget func(
		notebook *GtkNotebook,
		pack_type GtkPackType) *GtkWidget

	Gtk_notebook_set_action_widget func(
		notebook *GtkNotebook,
		widget *GtkWidget,
		pack_type GtkPackType)

	Gtk_offscreen_window_get_type func() GType

	Gtk_offscreen_window_new func() *GtkWidget

	Gtk_offscreen_window_get_pixmap func(
		offscreen *GtkOffscreenWindow) *GdkPixmap

	Gtk_offscreen_window_get_pixbuf func(
		offscreen *GtkOffscreenWindow) *GdkPixbuf

	Gtk_orientable_get_type func() GType

	Gtk_orientable_set_orientation func(
		orientable *GtkOrientable,
		orientation GtkOrientation)

	Gtk_orientable_get_orientation func(
		orientable *GtkOrientable) GtkOrientation

	Gtk_paper_size_get_type func() GType

	Gtk_paper_size_new func(
		name string) *GtkPaperSize

	Gtk_paper_size_new_from_ppd func(
		ppd_name string,
		ppd_display_name string,
		width Gdouble,
		height Gdouble) *GtkPaperSize

	Gtk_paper_size_new_custom func(
		name string,
		display_name string,
		width Gdouble,
		height Gdouble,
		unit GtkUnit) *GtkPaperSize

	Gtk_paper_size_copy func(
		other *GtkPaperSize) *GtkPaperSize

	Gtk_paper_size_free func(
		size *GtkPaperSize)

	Gtk_paper_size_is_equal func(
		size1 *GtkPaperSize,
		size2 *GtkPaperSize) Gboolean

	Gtk_paper_size_get_paper_sizes func(
		include_custom Gboolean) *GList

	Gtk_paper_size_get_name func(
		size *GtkPaperSize) string

	Gtk_paper_size_get_display_name func(
		size *GtkPaperSize) string

	Gtk_paper_size_get_ppd_name func(
		size *GtkPaperSize) string

	Gtk_paper_size_get_width func(
		size *GtkPaperSize,
		unit GtkUnit) Gdouble

	Gtk_paper_size_get_height func(
		size *GtkPaperSize,
		unit GtkUnit) Gdouble

	Gtk_paper_size_is_custom func(
		size *GtkPaperSize) Gboolean

	Gtk_paper_size_set_size func(
		size *GtkPaperSize,
		width Gdouble,
		height Gdouble,
		unit GtkUnit)

	Gtk_paper_size_get_default_top_margin func(
		size *GtkPaperSize,
		unit GtkUnit) Gdouble

	Gtk_paper_size_get_default_bottom_margin func(
		size *GtkPaperSize,
		unit GtkUnit) Gdouble

	Gtk_paper_size_get_default_left_margin func(
		size *GtkPaperSize,
		unit GtkUnit) Gdouble

	Gtk_paper_size_get_default_right_margin func(
		size *GtkPaperSize,
		unit GtkUnit) Gdouble

	Gtk_paper_size_get_default func() string

	Gtk_paper_size_new_from_key_file func(
		key_file *GKeyFile,
		group_name string,
		error **GError) *GtkPaperSize

	Gtk_paper_size_to_key_file func(
		size *GtkPaperSize,
		key_file *GKeyFile,
		group_name string)

	Gtk_page_setup_get_type func() GType

	Gtk_page_setup_new func() *GtkPageSetup

	Gtk_page_setup_copy func(
		other *GtkPageSetup) *GtkPageSetup

	Gtk_page_setup_get_orientation func(
		setup *GtkPageSetup) GtkPageOrientation

	Gtk_page_setup_set_orientation func(
		setup *GtkPageSetup,
		orientation GtkPageOrientation)

	Gtk_page_setup_get_paper_size func(
		setup *GtkPageSetup) *GtkPaperSize

	Gtk_page_setup_set_paper_size func(
		setup *GtkPageSetup,
		size *GtkPaperSize)

	Gtk_page_setup_get_top_margin func(
		setup *GtkPageSetup,
		unit GtkUnit) Gdouble

	Gtk_page_setup_set_top_margin func(
		setup *GtkPageSetup,
		margin Gdouble,
		unit GtkUnit)

	Gtk_page_setup_get_bottom_margin func(
		setup *GtkPageSetup,
		unit GtkUnit) Gdouble

	Gtk_page_setup_set_bottom_margin func(
		setup *GtkPageSetup,
		margin Gdouble,
		unit GtkUnit)

	Gtk_page_setup_get_left_margin func(
		setup *GtkPageSetup,
		unit GtkUnit) Gdouble

	Gtk_page_setup_set_left_margin func(
		setup *GtkPageSetup,
		margin Gdouble,
		unit GtkUnit)

	Gtk_page_setup_get_right_margin func(
		setup *GtkPageSetup,
		unit GtkUnit) Gdouble

	Gtk_page_setup_set_right_margin func(
		setup *GtkPageSetup,
		margin Gdouble,
		unit GtkUnit)

	Gtk_page_setup_set_paper_size_and_default_margins func(
		setup *GtkPageSetup,
		size *GtkPaperSize)

	Gtk_page_setup_get_paper_width func(
		setup *GtkPageSetup,
		unit GtkUnit) Gdouble

	Gtk_page_setup_get_paper_height func(
		setup *GtkPageSetup,
		unit GtkUnit) Gdouble

	Gtk_page_setup_get_page_width func(
		setup *GtkPageSetup,
		unit GtkUnit) Gdouble

	Gtk_page_setup_get_page_height func(
		setup *GtkPageSetup,
		unit GtkUnit) Gdouble

	Gtk_page_setup_new_from_file func(
		file_name string,
		error **GError) *GtkPageSetup

	Gtk_page_setup_load_file func(
		setup *GtkPageSetup,
		file_name string,
		error **GError) Gboolean

	Gtk_page_setup_to_file func(
		setup *GtkPageSetup,
		file_name string,
		error **GError) Gboolean

	Gtk_page_setup_new_from_key_file func(
		key_file *GKeyFile,
		group_name string,
		error **GError) *GtkPageSetup

	Gtk_page_setup_load_key_file func(
		setup *GtkPageSetup,
		key_file *GKeyFile,
		group_name string,
		error **GError) Gboolean

	Gtk_page_setup_to_key_file func(
		setup *GtkPageSetup,
		key_file *GKeyFile,
		group_name string)

	Gtk_socket_get_type func() GType

	Gtk_socket_new func() *GtkWidget

	Gtk_socket_add_id func(
		socket_ *GtkSocket,
		window_id GdkNativeWindow)

	Gtk_socket_get_id func(
		socket_ *GtkSocket) GdkNativeWindow

	Gtk_socket_get_plug_window func(
		socket_ *GtkSocket) *GdkWindow

	Gtk_socket_steal func(
		socket_ *GtkSocket,
		wid GdkNativeWindow)

	Gtk_plug_get_type func() GType

	Gtk_plug_construct func(
		plug *GtkPlug,
		socket_id GdkNativeWindow)

	Gtk_plug_new func(
		socket_id GdkNativeWindow) *GtkWidget

	Gtk_plug_construct_for_display func(
		plug *GtkPlug,
		display *GdkDisplay,
		socket_id GdkNativeWindow)

	Gtk_plug_new_for_display func(
		display *GdkDisplay,
		socket_id GdkNativeWindow) *GtkWidget

	Gtk_plug_get_id func(
		plug *GtkPlug) GdkNativeWindow

	Gtk_plug_get_embedded func(
		plug *GtkPlug) Gboolean

	Gtk_plug_get_socket_window func(
		plug *GtkPlug) *GdkWindow

	Gtk_print_context_get_type func() GType

	Gtk_print_context_get_cairo_context func(
		context *GtkPrintContext) *Cairo_t

	Gtk_print_context_get_page_setup func(
		context *GtkPrintContext) *GtkPageSetup

	Gtk_print_context_get_width func(
		context *GtkPrintContext) Gdouble

	Gtk_print_context_get_height func(
		context *GtkPrintContext) Gdouble

	Gtk_print_context_get_dpi_x func(
		context *GtkPrintContext) Gdouble

	Gtk_print_context_get_dpi_y func(
		context *GtkPrintContext) Gdouble

	Gtk_print_context_get_hard_margins func(
		context *GtkPrintContext,
		top *Gdouble,
		bottom *Gdouble,
		left *Gdouble,
		right *Gdouble) Gboolean

	Gtk_print_context_get_pango_fontmap func(
		context *GtkPrintContext) *PangoFontMap

	Gtk_print_context_create_pango_context func(
		context *GtkPrintContext) *PangoContext

	Gtk_print_context_create_pango_layout func(
		context *GtkPrintContext) *PangoLayout

	Gtk_print_settings_get_type func() GType

	Gtk_print_settings_new func() *GtkPrintSettings

	Gtk_print_settings_copy func(
		other *GtkPrintSettings) *GtkPrintSettings

	Gtk_print_settings_new_from_file func(
		file_name string,
		error **GError) *GtkPrintSettings

	Gtk_print_settings_load_file func(
		settings *GtkPrintSettings,
		file_name string,
		error **GError) Gboolean

	Gtk_print_settings_to_file func(
		settings *GtkPrintSettings,
		file_name string,
		error **GError) Gboolean

	Gtk_print_settings_new_from_key_file func(
		key_file *GKeyFile,
		group_name string,
		error **GError) *GtkPrintSettings

	Gtk_print_settings_load_key_file func(
		settings *GtkPrintSettings,
		key_file *GKeyFile,
		group_name string,
		error **GError) Gboolean

	Gtk_print_settings_to_key_file func(
		settings *GtkPrintSettings,
		key_file *GKeyFile,
		group_name string)

	Gtk_print_settings_has_key func(
		settings *GtkPrintSettings,
		key string) Gboolean

	Gtk_print_settings_get func(
		settings *GtkPrintSettings,
		key string) string

	Gtk_print_settings_set func(
		settings *GtkPrintSettings,
		key string,
		value string)

	Gtk_print_settings_unset func(
		settings *GtkPrintSettings,
		key string)

	Gtk_print_settings_foreach func(
		settings *GtkPrintSettings,
		f GtkPrintSettingsFunc,
		user_data Gpointer)

	Gtk_print_settings_get_bool func(
		settings *GtkPrintSettings,
		key string) Gboolean

	Gtk_print_settings_set_bool func(
		settings *GtkPrintSettings,
		key string,
		value Gboolean)

	Gtk_print_settings_get_double func(
		settings *GtkPrintSettings,
		key string) Gdouble

	Gtk_print_settings_get_double_with_default func(
		settings *GtkPrintSettings,
		key string,
		def Gdouble) Gdouble

	Gtk_print_settings_set_double func(
		settings *GtkPrintSettings,
		key string,
		value Gdouble)

	Gtk_print_settings_get_length func(
		settings *GtkPrintSettings,
		key string,
		unit GtkUnit) Gdouble

	Gtk_print_settings_set_length func(
		settings *GtkPrintSettings,
		key string,
		value Gdouble,
		unit GtkUnit)

	Gtk_print_settings_get_int func(
		settings *GtkPrintSettings,
		key string) Gint

	Gtk_print_settings_get_int_with_default func(
		settings *GtkPrintSettings,
		key string,
		def Gint) Gint

	Gtk_print_settings_set_int func(
		settings *GtkPrintSettings,
		key string,
		value Gint)

	Gtk_print_settings_get_printer func(
		settings *GtkPrintSettings) string

	Gtk_print_settings_set_printer func(
		settings *GtkPrintSettings,
		printer string)

	Gtk_print_settings_get_orientation func(
		settings *GtkPrintSettings) GtkPageOrientation

	Gtk_print_settings_set_orientation func(
		settings *GtkPrintSettings,
		orientation GtkPageOrientation)

	Gtk_print_settings_get_paper_size func(
		settings *GtkPrintSettings) *GtkPaperSize

	Gtk_print_settings_set_paper_size func(
		settings *GtkPrintSettings,
		paper_size *GtkPaperSize)

	Gtk_print_settings_get_paper_width func(
		settings *GtkPrintSettings,
		unit GtkUnit) Gdouble

	Gtk_print_settings_set_paper_width func(
		settings *GtkPrintSettings,
		width Gdouble,
		unit GtkUnit)

	Gtk_print_settings_get_paper_height func(
		settings *GtkPrintSettings,
		unit GtkUnit) Gdouble

	Gtk_print_settings_set_paper_height func(
		settings *GtkPrintSettings,
		height Gdouble,
		unit GtkUnit)

	Gtk_print_settings_get_use_color func(
		settings *GtkPrintSettings) Gboolean

	Gtk_print_settings_set_use_color func(
		settings *GtkPrintSettings,
		use_color Gboolean)

	Gtk_print_settings_get_collate func(
		settings *GtkPrintSettings) Gboolean

	Gtk_print_settings_set_collate func(
		settings *GtkPrintSettings,
		collate Gboolean)

	Gtk_print_settings_get_reverse func(
		settings *GtkPrintSettings) Gboolean

	Gtk_print_settings_set_reverse func(
		settings *GtkPrintSettings,
		reverse Gboolean)

	Gtk_print_settings_get_duplex func(
		settings *GtkPrintSettings) GtkPrintDuplex

	Gtk_print_settings_set_duplex func(
		settings *GtkPrintSettings,
		duplex GtkPrintDuplex)

	Gtk_print_settings_get_quality func(
		settings *GtkPrintSettings) GtkPrintQuality

	Gtk_print_settings_set_quality func(
		settings *GtkPrintSettings,
		quality GtkPrintQuality)

	Gtk_print_settings_get_n_copies func(
		settings *GtkPrintSettings) Gint

	Gtk_print_settings_set_n_copies func(
		settings *GtkPrintSettings,
		num_copies Gint)

	Gtk_print_settings_get_number_up func(
		settings *GtkPrintSettings) Gint

	Gtk_print_settings_set_number_up func(
		settings *GtkPrintSettings,
		number_up Gint)

	Gtk_print_settings_get_number_up_layout func(
		settings *GtkPrintSettings) GtkNumberUpLayout

	Gtk_print_settings_set_number_up_layout func(
		settings *GtkPrintSettings,
		number_up_layout GtkNumberUpLayout)

	Gtk_print_settings_get_resolution func(
		settings *GtkPrintSettings) Gint

	Gtk_print_settings_set_resolution func(
		settings *GtkPrintSettings,
		resolution Gint)

	Gtk_print_settings_get_resolution_x func(
		settings *GtkPrintSettings) Gint

	Gtk_print_settings_get_resolution_y func(
		settings *GtkPrintSettings) Gint

	Gtk_print_settings_set_resolution_xy func(
		settings *GtkPrintSettings,
		resolution_x Gint,
		resolution_y Gint)

	Gtk_print_settings_get_printer_lpi func(
		settings *GtkPrintSettings) Gdouble

	Gtk_print_settings_set_printer_lpi func(
		settings *GtkPrintSettings,
		lpi Gdouble)

	Gtk_print_settings_get_scale func(
		settings *GtkPrintSettings) Gdouble

	Gtk_print_settings_set_scale func(
		settings *GtkPrintSettings,
		scale Gdouble)

	Gtk_print_settings_get_print_pages func(
		settings *GtkPrintSettings) GtkPrintPages

	Gtk_print_settings_set_print_pages func(
		settings *GtkPrintSettings,
		pages GtkPrintPages)

	Gtk_print_settings_get_page_ranges func(
		settings *GtkPrintSettings,
		num_ranges *Gint) *GtkPageRange

	Gtk_print_settings_set_page_ranges func(
		settings *GtkPrintSettings,
		page_ranges *GtkPageRange,
		num_ranges Gint)

	Gtk_print_settings_get_page_set func(
		settings *GtkPrintSettings) GtkPageSet

	Gtk_print_settings_set_page_set func(
		settings *GtkPrintSettings,
		page_set GtkPageSet)

	Gtk_print_settings_get_default_source func(
		settings *GtkPrintSettings) string

	Gtk_print_settings_set_default_source func(
		settings *GtkPrintSettings,
		default_source string)

	Gtk_print_settings_get_media_type func(
		settings *GtkPrintSettings) string

	Gtk_print_settings_set_media_type func(
		settings *GtkPrintSettings,
		media_type string)

	Gtk_print_settings_get_dither func(
		settings *GtkPrintSettings) string

	Gtk_print_settings_set_dither func(
		settings *GtkPrintSettings,
		dither string)

	Gtk_print_settings_get_finishings func(
		settings *GtkPrintSettings) string

	Gtk_print_settings_set_finishings func(
		settings *GtkPrintSettings,
		finishings string)

	Gtk_print_settings_get_output_bin func(
		settings *GtkPrintSettings) string

	Gtk_print_settings_set_output_bin func(
		settings *GtkPrintSettings,
		output_bin string)

	Gtk_print_operation_preview_get_type func() GType

	Gtk_print_operation_preview_render_page func(
		preview *GtkPrintOperationPreview,
		page_nr Gint)

	Gtk_print_operation_preview_end_preview func(
		preview *GtkPrintOperationPreview)

	Gtk_print_operation_preview_is_selected func(
		preview *GtkPrintOperationPreview,
		page_nr Gint) Gboolean

	Gtk_print_error_quark func() GQuark

	Gtk_print_operation_get_type func() GType

	Gtk_print_operation_new func() *GtkPrintOperation

	Gtk_print_operation_set_default_page_setup func(
		op *GtkPrintOperation,
		default_page_setup *GtkPageSetup)

	Gtk_print_operation_get_default_page_setup func(
		op *GtkPrintOperation) *GtkPageSetup

	Gtk_print_operation_set_print_settings func(
		op *GtkPrintOperation,
		print_settings *GtkPrintSettings)

	Gtk_print_operation_get_print_settings func(
		op *GtkPrintOperation) *GtkPrintSettings

	Gtk_print_operation_set_job_name func(
		op *GtkPrintOperation,
		job_name string)

	Gtk_print_operation_set_n_pages func(
		op *GtkPrintOperation,
		n_pages Gint)

	Gtk_print_operation_set_current_page func(
		op *GtkPrintOperation,
		current_page Gint)

	Gtk_print_operation_set_use_full_page func(
		op *GtkPrintOperation,
		full_page Gboolean)

	Gtk_print_operation_set_unit func(
		op *GtkPrintOperation,
		unit GtkUnit)

	Gtk_print_operation_set_export_filename func(
		op *GtkPrintOperation,
		filename string)

	Gtk_print_operation_set_track_print_status func(
		op *GtkPrintOperation,
		track_status Gboolean)

	Gtk_print_operation_set_show_progress func(
		op *GtkPrintOperation,
		show_progress Gboolean)

	Gtk_print_operation_set_allow_async func(
		op *GtkPrintOperation,
		allow_async Gboolean)

	Gtk_print_operation_set_custom_tab_label func(
		op *GtkPrintOperation,
		label string)

	Gtk_print_operation_run func(
		op *GtkPrintOperation,
		action GtkPrintOperationAction,
		parent *GtkWindow,
		error **GError) GtkPrintOperationResult

	Gtk_print_operation_get_error func(
		op *GtkPrintOperation,
		error **GError)

	Gtk_print_operation_get_status func(
		op *GtkPrintOperation) GtkPrintStatus

	Gtk_print_operation_get_status_string func(
		op *GtkPrintOperation) string

	Gtk_print_operation_is_finished func(
		op *GtkPrintOperation) Gboolean

	Gtk_print_operation_cancel func(
		op *GtkPrintOperation)

	Gtk_print_operation_draw_page_finish func(
		op *GtkPrintOperation)

	Gtk_print_operation_set_defer_drawing func(
		op *GtkPrintOperation)

	Gtk_print_operation_set_support_selection func(
		op *GtkPrintOperation,
		support_selection Gboolean)

	Gtk_print_operation_get_support_selection func(
		op *GtkPrintOperation) Gboolean

	Gtk_print_operation_set_has_selection func(
		op *GtkPrintOperation,
		has_selection Gboolean)

	Gtk_print_operation_get_has_selection func(
		op *GtkPrintOperation) Gboolean

	Gtk_print_operation_set_embed_page_setup func(
		op *GtkPrintOperation,
		embed Gboolean)

	Gtk_print_operation_get_embed_page_setup func(
		op *GtkPrintOperation) Gboolean

	Gtk_print_operation_get_n_pages_to_print func(
		op *GtkPrintOperation) Gint

	Gtk_print_run_page_setup_dialog func(
		parent *GtkWindow,
		page_setup *GtkPageSetup,
		settings *GtkPrintSettings) *GtkPageSetup

	Gtk_print_run_page_setup_dialog_async func(
		parent *GtkWindow,
		page_setup *GtkPageSetup,
		settings *GtkPrintSettings,
		done_cb GtkPageSetupDoneFunc,
		data Gpointer)

	Gtk_progress_get_type func() GType

	Gtk_progress_set_show_text func(
		progress *GtkProgress,
		show_text Gboolean)

	Gtk_progress_set_text_alignment func(
		progress *GtkProgress,
		x_align Gfloat,
		y_align Gfloat)

	Gtk_progress_set_format_string func(
		progress *GtkProgress,
		format string)

	Gtk_progress_set_adjustment func(
		progress *GtkProgress,
		adjustment *GtkAdjustment)

	Gtk_progress_configure func(
		progress *GtkProgress,
		value Gdouble,
		min Gdouble,
		max Gdouble)

	Gtk_progress_set_percentage func(
		progress *GtkProgress,
		percentage Gdouble)

	Gtk_progress_set_value func(
		progress *GtkProgress,
		value Gdouble)

	Gtk_progress_get_value func(
		progress *GtkProgress) Gdouble

	Gtk_progress_set_activity_mode func(
		progress *GtkProgress,
		activity_mode Gboolean)

	Gtk_progress_get_current_text func(
		progress *GtkProgress) string

	Gtk_progress_get_text_from_value func(
		progress *GtkProgress,
		value Gdouble) string

	Gtk_progress_get_current_percentage func(
		progress *GtkProgress) Gdouble

	Gtk_progress_get_percentage_from_value func(
		progress *GtkProgress,
		value Gdouble) Gdouble

	Gtk_progress_bar_get_type func() GType

	Gtk_progress_bar_new func() *GtkWidget

	Gtk_progress_bar_pulse func(
		pbar *GtkProgressBar)

	Gtk_progress_bar_set_text func(
		pbar *GtkProgressBar,
		text string)

	Gtk_progress_bar_set_fraction func(
		pbar *GtkProgressBar,
		fraction Gdouble)

	Gtk_progress_bar_set_pulse_step func(
		pbar *GtkProgressBar,
		fraction Gdouble)

	Gtk_progress_bar_set_orientation func(
		pbar *GtkProgressBar,
		orientation GtkProgressBarOrientation)

	Gtk_progress_bar_get_text func(
		pbar *GtkProgressBar) string

	Gtk_progress_bar_get_fraction func(
		pbar *GtkProgressBar) Gdouble

	Gtk_progress_bar_get_pulse_step func(
		pbar *GtkProgressBar) Gdouble

	Gtk_progress_bar_get_orientation func(
		pbar *GtkProgressBar) GtkProgressBarOrientation

	Gtk_progress_bar_set_ellipsize func(
		pbar *GtkProgressBar,
		mode PangoEllipsizeMode)

	Gtk_progress_bar_get_ellipsize func(
		pbar *GtkProgressBar) PangoEllipsizeMode

	Gtk_progress_bar_new_with_adjustment func(
		adjustment *GtkAdjustment) *GtkWidget

	Gtk_progress_bar_set_bar_style func(
		pbar *GtkProgressBar,
		style GtkProgressBarStyle)

	Gtk_progress_bar_set_discrete_blocks func(
		pbar *GtkProgressBar,
		blocks Guint)

	Gtk_progress_bar_set_activity_step func(
		pbar *GtkProgressBar,
		step Guint)

	Gtk_progress_bar_set_activity_blocks func(
		pbar *GtkProgressBar,
		blocks Guint)

	Gtk_progress_bar_update func(
		pbar *GtkProgressBar,
		percentage Gdouble)

	Gtk_toggle_action_get_type func() GType

	Gtk_toggle_action_new func(
		name string,
		label string,
		tooltip string,
		stock_id string) *GtkToggleAction

	Gtk_toggle_action_toggled func(
		action *GtkToggleAction)

	Gtk_toggle_action_set_active func(
		action *GtkToggleAction,
		is_active Gboolean)

	Gtk_toggle_action_get_active func(
		action *GtkToggleAction) Gboolean

	Gtk_toggle_action_set_draw_as_radio func(
		action *GtkToggleAction,
		draw_as_radio Gboolean)

	Gtk_toggle_action_get_draw_as_radio func(
		action *GtkToggleAction) Gboolean

	Gtk_radio_action_get_type func() GType

	Gtk_radio_action_new func(
		name string,
		label string,
		tooltip string,
		stock_id string,
		value Gint) *GtkRadioAction

	Gtk_radio_action_get_group func(
		action *GtkRadioAction) *GSList

	Gtk_radio_action_set_group func(
		action *GtkRadioAction,
		group *GSList)

	Gtk_radio_action_get_current_value func(
		action *GtkRadioAction) Gint

	Gtk_radio_action_set_current_value func(
		action *GtkRadioAction,
		current_value Gint)

	Gtk_radio_button_get_type func() GType

	Gtk_radio_button_new func(
		group *GSList) *GtkWidget

	Gtk_radio_button_new_from_widget func(
		radio_group_member *GtkRadioButton) *GtkWidget

	Gtk_radio_button_new_with_label func(
		group *GSList,
		label string) *GtkWidget

	Gtk_radio_button_new_with_label_from_widget func(
		radio_group_member *GtkRadioButton,
		label string) *GtkWidget

	Gtk_radio_button_new_with_mnemonic func(
		group *GSList,
		label string) *GtkWidget

	Gtk_radio_button_new_with_mnemonic_from_widget func(
		radio_group_member *GtkRadioButton,
		label string) *GtkWidget

	Gtk_radio_button_get_group func(
		radio_button *GtkRadioButton) *GSList

	Gtk_radio_button_set_group func(
		radio_button *GtkRadioButton,
		group *GSList)

	Gtk_radio_menu_item_get_type func() GType

	Gtk_radio_menu_item_new func(
		group *GSList) *GtkWidget

	Gtk_radio_menu_item_new_with_label func(
		group *GSList,
		label string) *GtkWidget

	Gtk_radio_menu_item_new_with_mnemonic func(
		group *GSList,
		label string) *GtkWidget

	Gtk_radio_menu_item_new_from_widget func(
		group *GtkRadioMenuItem) *GtkWidget

	Gtk_radio_menu_item_new_with_mnemonic_from_widget func(
		group *GtkRadioMenuItem,
		label string) *GtkWidget

	Gtk_radio_menu_item_new_with_label_from_widget func(
		group *GtkRadioMenuItem,
		label string) *GtkWidget

	Gtk_radio_menu_item_get_group func(
		radio_menu_item *GtkRadioMenuItem) *GSList

	Gtk_radio_menu_item_set_group func(
		radio_menu_item *GtkRadioMenuItem,
		group *GSList)

	Gtk_toggle_tool_button_get_type func() GType

	Gtk_toggle_tool_button_new func() *GtkToolItem

	Gtk_toggle_tool_button_new_from_stock func(
		stock_id string) *GtkToolItem

	Gtk_toggle_tool_button_set_active func(
		button *GtkToggleToolButton,
		is_active Gboolean)

	Gtk_toggle_tool_button_get_active func(
		button *GtkToggleToolButton) Gboolean

	Gtk_radio_tool_button_get_type func() GType

	Gtk_radio_tool_button_new func(
		group *GSList) *GtkToolItem

	Gtk_radio_tool_button_new_from_stock func(
		group *GSList,
		stock_id string) *GtkToolItem

	Gtk_radio_tool_button_new_from_widget func(
		group *GtkRadioToolButton) *GtkToolItem

	Gtk_radio_tool_button_new_with_stock_from_widget func(
		group *GtkRadioToolButton,
		stock_id string) *GtkToolItem

	Gtk_radio_tool_button_get_group func(
		button *GtkRadioToolButton) *GSList

	Gtk_radio_tool_button_set_group func(
		button *GtkRadioToolButton,
		group *GSList)

	Gtk_recent_manager_error_quark func() GQuark

	Gtk_recent_manager_get_type func() GType

	Gtk_recent_manager_new func() *GtkRecentManager

	Gtk_recent_manager_get_default func() *GtkRecentManager

	Gtk_recent_manager_get_for_screen func(
		screen *GdkScreen) *GtkRecentManager

	Gtk_recent_manager_set_screen func(
		manager *GtkRecentManager,
		screen *GdkScreen)

	Gtk_recent_manager_add_item func(
		manager *GtkRecentManager,
		uri string) Gboolean

	Gtk_recent_manager_add_full func(
		manager *GtkRecentManager,
		uri string,
		recent_data *GtkRecentData) Gboolean

	Gtk_recent_manager_remove_item func(
		manager *GtkRecentManager,
		uri string,
		error **GError) Gboolean

	Gtk_recent_manager_lookup_item func(
		manager *GtkRecentManager,
		uri string,
		error **GError) *GtkRecentInfo

	Gtk_recent_manager_has_item func(
		manager *GtkRecentManager,
		uri string) Gboolean

	Gtk_recent_manager_move_item func(
		manager *GtkRecentManager,
		uri string,
		new_uri string,
		error **GError) Gboolean

	Gtk_recent_manager_set_limit func(
		manager *GtkRecentManager,
		limit Gint)

	Gtk_recent_manager_get_limit func(
		manager *GtkRecentManager) Gint

	Gtk_recent_manager_get_items func(
		manager *GtkRecentManager) *GList

	Gtk_recent_manager_purge_items func(
		manager *GtkRecentManager,
		error **GError) Gint

	Gtk_recent_info_get_type func() GType

	Gtk_recent_info_ref func(
		info *GtkRecentInfo) *GtkRecentInfo

	Gtk_recent_info_unref func(
		info *GtkRecentInfo)

	Gtk_recent_info_get_uri func(
		info *GtkRecentInfo) string

	Gtk_recent_info_get_display_name func(
		info *GtkRecentInfo) string

	Gtk_recent_info_get_description func(
		info *GtkRecentInfo) string

	Gtk_recent_info_get_mime_type func(
		info *GtkRecentInfo) string

	Gtk_recent_info_get_added func(
		info *GtkRecentInfo) Time_t

	Gtk_recent_info_get_modified func(
		info *GtkRecentInfo) Time_t

	Gtk_recent_info_get_visited func(
		info *GtkRecentInfo) Time_t

	Gtk_recent_info_get_private_hint func(
		info *GtkRecentInfo) Gboolean

	Gtk_recent_info_get_application_info func(
		info *GtkRecentInfo,
		app_name string,
		app_exec **Gchar,
		count *Guint,
		time_ *Time_t) Gboolean

	Gtk_recent_info_get_applications func(
		info *GtkRecentInfo,
		length *Gsize) **Gchar

	Gtk_recent_info_last_application func(
		info *GtkRecentInfo) string

	Gtk_recent_info_has_application func(
		info *GtkRecentInfo,
		app_name string) Gboolean

	Gtk_recent_info_get_groups func(
		info *GtkRecentInfo,
		length *Gsize) **Gchar

	Gtk_recent_info_has_group func(
		info *GtkRecentInfo,
		group_name string) Gboolean

	Gtk_recent_info_get_icon func(
		info *GtkRecentInfo,
		size Gint) *GdkPixbuf

	Gtk_recent_info_get_short_name func(
		info *GtkRecentInfo) string

	Gtk_recent_info_get_uri_display func(
		info *GtkRecentInfo) string

	Gtk_recent_info_get_age func(
		info *GtkRecentInfo) Gint

	Gtk_recent_info_is_local func(
		info *GtkRecentInfo) Gboolean

	Gtk_recent_info_exists func(
		info *GtkRecentInfo) Gboolean

	Gtk_recent_info_match func(
		info_a *GtkRecentInfo,
		info_b *GtkRecentInfo) Gboolean

	Gtk_recent_action_get_type func() GType

	Gtk_recent_action_new func(
		name string,
		label string,
		tooltip string,
		stock_id string) *GtkAction

	Gtk_recent_action_new_for_manager func(
		name string,
		label string,
		tooltip string,
		stock_id string,
		manager *GtkRecentManager) *GtkAction

	Gtk_recent_action_get_show_numbers func(
		action *GtkRecentAction) Gboolean

	Gtk_recent_action_set_show_numbers func(
		action *GtkRecentAction,
		show_numbers Gboolean)

	Gtk_recent_filter_get_type func() GType

	Gtk_recent_filter_new func() *GtkRecentFilter

	Gtk_recent_filter_set_name func(
		filter *GtkRecentFilter,
		name string)

	Gtk_recent_filter_get_name func(
		filter *GtkRecentFilter) string

	Gtk_recent_filter_add_mime_type func(
		filter *GtkRecentFilter,
		mime_type string)

	Gtk_recent_filter_add_pattern func(
		filter *GtkRecentFilter,
		pattern string)

	Gtk_recent_filter_add_pixbuf_formats func(
		filter *GtkRecentFilter)

	Gtk_recent_filter_add_application func(
		filter *GtkRecentFilter,
		application string)

	Gtk_recent_filter_add_group func(
		filter *GtkRecentFilter,
		group string)

	Gtk_recent_filter_add_age func(
		filter *GtkRecentFilter,
		days Gint)

	Gtk_recent_filter_add_custom func(
		filter *GtkRecentFilter,
		needed GtkRecentFilterFlags,
		f GtkRecentFilterFunc,
		dataGpointer,
		data_destroy GDestroyNotify)

	Gtk_recent_filter_get_needed func(
		filter *GtkRecentFilter) GtkRecentFilterFlags

	Gtk_recent_filter_filter func(
		filter *GtkRecentFilter,
		filter_info *GtkRecentFilterInfo) Gboolean

	Gtk_recent_chooser_error_quark func() GQuark

	Gtk_recent_chooser_get_type func() GType

	Gtk_recent_chooser_set_show_private func(
		chooser *GtkRecentChooser,
		show_private Gboolean)

	Gtk_recent_chooser_get_show_private func(
		chooser *GtkRecentChooser) Gboolean

	Gtk_recent_chooser_set_show_not_found func(
		chooser *GtkRecentChooser,
		show_not_found Gboolean)

	Gtk_recent_chooser_get_show_not_found func(
		chooser *GtkRecentChooser) Gboolean

	Gtk_recent_chooser_set_select_multiple func(
		chooser *GtkRecentChooser,
		select_multiple Gboolean)

	Gtk_recent_chooser_get_select_multiple func(
		chooser *GtkRecentChooser) Gboolean

	Gtk_recent_chooser_set_limit func(
		chooser *GtkRecentChooser,
		limit Gint)

	Gtk_recent_chooser_get_limit func(
		chooser *GtkRecentChooser) Gint

	Gtk_recent_chooser_set_local_only func(
		chooser *GtkRecentChooser,
		local_only Gboolean)

	Gtk_recent_chooser_get_local_only func(
		chooser *GtkRecentChooser) Gboolean

	Gtk_recent_chooser_set_show_tips func(
		chooser *GtkRecentChooser,
		show_tips Gboolean)

	Gtk_recent_chooser_get_show_tips func(
		chooser *GtkRecentChooser) Gboolean

	Gtk_recent_chooser_set_show_numbers func(
		chooser *GtkRecentChooser,
		show_numbers Gboolean)

	Gtk_recent_chooser_get_show_numbers func(
		chooser *GtkRecentChooser) Gboolean

	Gtk_recent_chooser_set_show_icons func(
		chooser *GtkRecentChooser,
		show_icons Gboolean)

	Gtk_recent_chooser_get_show_icons func(
		chooser *GtkRecentChooser) Gboolean

	Gtk_recent_chooser_set_sort_type func(
		chooser *GtkRecentChooser,
		sort_type GtkRecentSortType)

	Gtk_recent_chooser_get_sort_type func(
		chooser *GtkRecentChooser) GtkRecentSortType

	Gtk_recent_chooser_set_sort_func func(
		chooser *GtkRecentChooser,
		sort_func GtkRecentSortFunc,
		sort_dataGpointer,
		data_destroy GDestroyNotify)

	Gtk_recent_chooser_set_current_uri func(
		chooser *GtkRecentChooser,
		uri string,
		error **GError) Gboolean

	Gtk_recent_chooser_get_current_uri func(
		chooser *GtkRecentChooser) string

	Gtk_recent_chooser_get_current_item func(
		chooser *GtkRecentChooser) *GtkRecentInfo

	Gtk_recent_chooser_select_uri func(
		chooser *GtkRecentChooser,
		uri string,
		error **GError) Gboolean

	Gtk_recent_chooser_unselect_uri func(
		chooser *GtkRecentChooser,
		uri string)

	Gtk_recent_chooser_select_all func(
		chooser *GtkRecentChooser)

	Gtk_recent_chooser_unselect_all func(
		chooser *GtkRecentChooser)

	Gtk_recent_chooser_get_items func(
		chooser *GtkRecentChooser) *GList

	Gtk_recent_chooser_get_uris func(
		chooser *GtkRecentChooser,
		length *Gsize) **Gchar

	Gtk_recent_chooser_add_filter func(
		chooser *GtkRecentChooser,
		filter *GtkRecentFilter)

	Gtk_recent_chooser_remove_filter func(
		chooser *GtkRecentChooser,
		filter *GtkRecentFilter)

	Gtk_recent_chooser_list_filters func(
		chooser *GtkRecentChooser) *GSList

	Gtk_recent_chooser_set_filter func(
		chooser *GtkRecentChooser,
		filter *GtkRecentFilter)

	Gtk_recent_chooser_get_filter func(
		chooser *GtkRecentChooser) *GtkRecentFilter

	Gtk_recent_chooser_dialog_get_type func() GType

	Gtk_recent_chooser_dialog_new func(
		title string, parent *GtkWindow,
		first_button_text string, v ...VArg) *GtkWidget

	Gtk_recent_chooser_dialog_new_for_manager func(
		title string, parent *GtkWindow,
		manager *GtkRecentManager,
		first_button_text string, v ...VArg) *GtkWidget

	Gtk_recent_chooser_menu_get_type func() GType

	Gtk_recent_chooser_menu_new func() *GtkWidget

	Gtk_recent_chooser_menu_new_for_manager func(
		manager *GtkRecentManager) *GtkWidget

	Gtk_recent_chooser_menu_get_show_numbers func(
		menu *GtkRecentChooserMenu) Gboolean

	Gtk_recent_chooser_menu_set_show_numbers func(
		menu *GtkRecentChooserMenu,
		show_numbers Gboolean)

	Gtk_recent_chooser_widget_get_type func() GType

	Gtk_recent_chooser_widget_new func() *GtkWidget

	Gtk_recent_chooser_widget_new_for_manager func(
		manager *GtkRecentManager) *GtkWidget

	Gtk_scale_button_get_type func() GType

	Gtk_scale_button_new func(
		size GtkIconSize,
		min Gdouble,
		max Gdouble,
		step Gdouble,
		icons **Gchar) *GtkWidget

	Gtk_scale_button_set_icons func(
		button *GtkScaleButton,
		icons **Gchar)

	Gtk_scale_button_get_value func(
		button *GtkScaleButton) Gdouble

	Gtk_scale_button_set_value func(
		button *GtkScaleButton,
		value Gdouble)

	Gtk_scale_button_get_adjustment func(
		button *GtkScaleButton) *GtkAdjustment

	Gtk_scale_button_set_adjustment func(
		button *GtkScaleButton,
		adjustment *GtkAdjustment)

	Gtk_scale_button_get_plus_button func(
		button *GtkScaleButton) *GtkWidget

	Gtk_scale_button_get_minus_button func(
		button *GtkScaleButton) *GtkWidget

	Gtk_scale_button_get_popup func(
		button *GtkScaleButton) *GtkWidget

	Gtk_scale_button_get_orientation func(
		button *GtkScaleButton) GtkOrientation

	Gtk_scale_button_set_orientation func(
		button *GtkScaleButton,
		orientation GtkOrientation)

	Gtk_vscrollbar_get_type func() GType

	Gtk_vscrollbar_new func(
		adjustment *GtkAdjustment) *GtkWidget

	Gtk_viewport_get_type func() GType

	Gtk_viewport_new func(
		hadjustment *GtkAdjustment,
		vadjustment *GtkAdjustment) *GtkWidget

	Gtk_viewport_get_hadjustment func(
		viewport *GtkViewport) *GtkAdjustment

	Gtk_viewport_get_vadjustment func(
		viewport *GtkViewport) *GtkAdjustment

	Gtk_viewport_set_hadjustment func(
		viewport *GtkViewport,
		adjustment *GtkAdjustment)

	Gtk_viewport_set_vadjustment func(
		viewport *GtkViewport,
		adjustment *GtkAdjustment)

	Gtk_viewport_set_shadow_type func(
		viewport *GtkViewport,
		t GtkShadowType)

	Gtk_viewport_get_shadow_type func(
		viewport *GtkViewport) GtkShadowType

	Gtk_viewport_get_bin_window func(
		viewport *GtkViewport) *GdkWindow

	Gtk_viewport_get_view_window func(
		viewport *GtkViewport) *GdkWindow

	Gtk_scrolled_window_get_type func() GType

	Gtk_scrolled_window_new func(
		hadjustment *GtkAdjustment,
		vadjustment *GtkAdjustment) *GtkWidget

	Gtk_scrolled_window_set_hadjustment func(
		scrolled_window *GtkScrolledWindow,
		hadjustment *GtkAdjustment)

	Gtk_scrolled_window_set_vadjustment func(
		scrolled_window *GtkScrolledWindow,
		vadjustment *GtkAdjustment)

	Gtk_scrolled_window_get_hadjustment func(
		scrolled_window *GtkScrolledWindow) *GtkAdjustment

	Gtk_scrolled_window_get_vadjustment func(
		scrolled_window *GtkScrolledWindow) *GtkAdjustment

	Gtk_scrolled_window_get_hscrollbar func(
		scrolled_window *GtkScrolledWindow) *GtkWidget

	Gtk_scrolled_window_get_vscrollbar func(
		scrolled_window *GtkScrolledWindow) *GtkWidget

	Gtk_scrolled_window_set_policy func(
		scrolled_window *GtkScrolledWindow,
		hscrollbar_policy GtkPolicyType,
		vscrollbar_policy GtkPolicyType)

	Gtk_scrolled_window_get_policy func(
		scrolled_window *GtkScrolledWindow,
		hscrollbar_policy *GtkPolicyType,
		vscrollbar_policy *GtkPolicyType)

	Gtk_scrolled_window_set_placement func(
		scrolled_window *GtkScrolledWindow,
		window_placement GtkCornerType)

	Gtk_scrolled_window_unset_placement func(
		scrolled_window *GtkScrolledWindow)

	Gtk_scrolled_window_get_placement func(
		scrolled_window *GtkScrolledWindow) GtkCornerType

	Gtk_scrolled_window_set_shadow_type func(
		scrolled_window *GtkScrolledWindow,
		t GtkShadowType)

	Gtk_scrolled_window_get_shadow_type func(
		scrolled_window *GtkScrolledWindow) GtkShadowType

	Gtk_scrolled_window_add_with_viewport func(
		scrolled_window *GtkScrolledWindow,
		child *GtkWidget)

	Gtk_separator_menu_item_get_type func() GType

	Gtk_separator_menu_item_new func() *GtkWidget

	Gtk_separator_tool_item_get_type func() GType

	Gtk_separator_tool_item_new func() *GtkToolItem

	Gtk_separator_tool_item_get_draw func(
		item *GtkSeparatorToolItem) Gboolean

	Gtk_separator_tool_item_set_draw func(
		item *GtkSeparatorToolItem,
		draw Gboolean)

	Gtk_show_uri func(
		screen *GdkScreen,
		uri string,
		timestamp Guint32,
		error **GError) Gboolean

	Gtk_spin_button_get_type func() GType

	Gtk_spin_button_configure func(
		spin_button *GtkSpinButton,
		adjustment *GtkAdjustment,
		climb_rate Gdouble,
		digits Guint)

	Gtk_spin_button_new func(
		adjustment *GtkAdjustment,
		climb_rate Gdouble,
		digits Guint) *GtkWidget

	Gtk_spin_button_new_with_range func(
		min Gdouble,
		max Gdouble,
		step Gdouble) *GtkWidget

	Gtk_spin_button_set_adjustment func(
		spin_button *GtkSpinButton,
		adjustment *GtkAdjustment)

	Gtk_spin_button_get_adjustment func(
		spin_button *GtkSpinButton) *GtkAdjustment

	Gtk_spin_button_set_digits func(
		spin_button *GtkSpinButton,
		digits Guint)

	Gtk_spin_button_get_digits func(
		spin_button *GtkSpinButton) Guint

	Gtk_spin_button_set_increments func(
		spin_button *GtkSpinButton,
		step Gdouble,
		page Gdouble)

	Gtk_spin_button_get_increments func(
		spin_button *GtkSpinButton,
		step *Gdouble,
		page *Gdouble)

	Gtk_spin_button_set_range func(
		spin_button *GtkSpinButton,
		min Gdouble,
		max Gdouble)

	Gtk_spin_button_get_range func(
		spin_button *GtkSpinButton,
		min *Gdouble,
		max *Gdouble)

	Gtk_spin_button_get_value func(
		spin_button *GtkSpinButton) Gdouble

	Gtk_spin_button_get_value_as_int func(
		spin_button *GtkSpinButton) Gint

	Gtk_spin_button_set_value func(
		spin_button *GtkSpinButton,
		value Gdouble)

	Gtk_spin_button_set_update_policy func(
		spin_button *GtkSpinButton,
		policy GtkSpinButtonUpdatePolicy)

	Gtk_spin_button_get_update_policy func(
		spin_button *GtkSpinButton) GtkSpinButtonUpdatePolicy

	Gtk_spin_button_set_numeric func(
		spin_button *GtkSpinButton,
		numeric Gboolean)

	Gtk_spin_button_get_numeric func(
		spin_button *GtkSpinButton) Gboolean

	Gtk_spin_button_spin func(
		spin_button *GtkSpinButton,
		direction GtkSpinType,
		increment Gdouble)

	Gtk_spin_button_set_wrap func(
		spin_button *GtkSpinButton,
		wrap Gboolean)

	Gtk_spin_button_get_wrap func(
		spin_button *GtkSpinButton) Gboolean

	Gtk_spin_button_set_snap_to_ticks func(
		spin_button *GtkSpinButton,
		snap_to_ticks Gboolean)

	Gtk_spin_button_get_snap_to_ticks func(
		spin_button *GtkSpinButton) Gboolean

	Gtk_spin_button_update func(
		spin_button *GtkSpinButton)

	Gtk_spinner_get_type func() GType

	Gtk_spinner_new func() *GtkWidget

	Gtk_spinner_start func(
		spinner *GtkSpinner)

	Gtk_spinner_stop func(
		spinner *GtkSpinner)

	Gtk_statusbar_get_type func() GType

	Gtk_statusbar_new func() *GtkWidget

	Gtk_statusbar_get_context_id func(
		statusbar *GtkStatusbar,
		context_description string) Guint

	Gtk_statusbar_push func(
		statusbar *GtkStatusbar,
		context_id Guint,
		text string) Guint

	Gtk_statusbar_pop func(
		statusbar *GtkStatusbar,
		context_id Guint)

	Gtk_statusbar_remove func(
		statusbar *GtkStatusbar,
		context_id Guint,
		message_id Guint)

	Gtk_statusbar_remove_all func(
		statusbar *GtkStatusbar,
		context_id Guint)

	Gtk_statusbar_set_has_resize_grip func(
		statusbar *GtkStatusbar,
		setting Gboolean)

	Gtk_statusbar_get_has_resize_grip func(
		statusbar *GtkStatusbar) Gboolean

	Gtk_statusbar_get_message_area func(
		statusbar *GtkStatusbar) *GtkWidget

	Gtk_status_icon_get_type func() GType

	Gtk_status_icon_new func() *GtkStatusIcon

	Gtk_status_icon_new_from_pixbuf func(
		pixbuf *GdkPixbuf) *GtkStatusIcon

	Gtk_status_icon_new_from_file func(
		filename string) *GtkStatusIcon

	Gtk_status_icon_new_from_stock func(
		stock_id string) *GtkStatusIcon

	Gtk_status_icon_new_from_icon_name func(
		icon_name string) *GtkStatusIcon

	Gtk_status_icon_new_from_gicon func(
		icon *GIcon) *GtkStatusIcon

	Gtk_status_icon_set_from_pixbuf func(
		status_icon *GtkStatusIcon,
		pixbuf *GdkPixbuf)

	Gtk_status_icon_set_from_file func(
		status_icon *GtkStatusIcon,
		filename string)

	Gtk_status_icon_set_from_stock func(
		status_icon *GtkStatusIcon,
		stock_id string)

	Gtk_status_icon_set_from_icon_name func(
		status_icon *GtkStatusIcon,
		icon_name string)

	Gtk_status_icon_set_from_gicon func(
		status_icon *GtkStatusIcon,
		icon *GIcon)

	Gtk_status_icon_get_storage_type func(
		status_icon *GtkStatusIcon) GtkImageType

	Gtk_status_icon_get_pixbuf func(
		status_icon *GtkStatusIcon) *GdkPixbuf

	Gtk_status_icon_get_stock func(
		status_icon *GtkStatusIcon) string

	Gtk_status_icon_get_icon_name func(
		status_icon *GtkStatusIcon) string

	Gtk_status_icon_get_gicon func(
		status_icon *GtkStatusIcon) *GIcon

	Gtk_status_icon_get_size func(
		status_icon *GtkStatusIcon) Gint

	Gtk_status_icon_set_screen func(
		status_icon *GtkStatusIcon,
		screen *GdkScreen)

	Gtk_status_icon_get_screen func(
		status_icon *GtkStatusIcon) *GdkScreen

	Gtk_status_icon_set_tooltip func(
		status_icon *GtkStatusIcon,
		tooltip_text string)

	Gtk_status_icon_set_has_tooltip func(
		status_icon *GtkStatusIcon,
		has_tooltip Gboolean)

	Gtk_status_icon_set_tooltip_text func(
		status_icon *GtkStatusIcon,
		text string)

	Gtk_status_icon_set_tooltip_markup func(
		status_icon *GtkStatusIcon,
		markup string)

	Gtk_status_icon_set_title func(
		status_icon *GtkStatusIcon,
		title string)

	Gtk_status_icon_get_title func(
		status_icon *GtkStatusIcon) string

	Gtk_status_icon_set_name func(
		status_icon *GtkStatusIcon,
		name string)

	Gtk_status_icon_set_visible func(
		status_icon *GtkStatusIcon,
		visible Gboolean)

	Gtk_status_icon_get_visible func(
		status_icon *GtkStatusIcon) Gboolean

	Gtk_status_icon_set_blinking func(
		status_icon *GtkStatusIcon,
		blinking Gboolean)

	Gtk_status_icon_get_blinking func(
		status_icon *GtkStatusIcon) Gboolean

	Gtk_status_icon_is_embedded func(
		status_icon *GtkStatusIcon) Gboolean

	Gtk_status_icon_position_menu func(
		menu *GtkMenu,
		x *Gint,
		y *Gint,
		push_in *Gboolean,
		user_data Gpointer)

	Gtk_status_icon_get_geometry func(
		status_icon *GtkStatusIcon,
		screen **GdkScreen,
		area *GdkRectangle,
		orientation *GtkOrientation) Gboolean

	Gtk_status_icon_get_has_tooltip func(
		status_icon *GtkStatusIcon) Gboolean

	Gtk_status_icon_get_tooltip_text func(
		status_icon *GtkStatusIcon) string

	Gtk_status_icon_get_tooltip_markup func(
		status_icon *GtkStatusIcon) string

	Gtk_status_icon_get_x11_window_id func(
		status_icon *GtkStatusIcon) Guint32

	Gtk_stock_add func(
		items *GtkStockItem,
		n_items Guint)

	Gtk_stock_add_static func(
		items *GtkStockItem,
		n_items Guint)

	Gtk_stock_lookup func(
		stock_id string,
		item *GtkStockItem) Gboolean

	Gtk_stock_list_ids func() *GSList

	Gtk_stock_item_copy func(
		item *GtkStockItem) *GtkStockItem

	Gtk_stock_item_free func(
		item *GtkStockItem)

	Gtk_stock_set_translate_func func(
		domain string,
		f GtkTranslateFunc,
		dataGpointer,
		notify GDestroyNotify)

	Gtk_table_get_type func() GType

	Gtk_table_new func(
		rows Guint,
		columns Guint,
		homogeneous Gboolean) *GtkWidget

	Gtk_table_resize func(
		table *GtkTable,
		rows Guint,
		columns Guint)

	Gtk_table_attach func(
		table *GtkTable,
		child *GtkWidget,
		left_attach Guint,
		right_attach Guint,
		top_attach Guint,
		bottom_attach Guint,
		xoptions GtkAttachOptions,
		yoptions GtkAttachOptions,
		xpadding Guint,
		ypadding Guint)

	Gtk_table_attach_defaults func(
		table *GtkTable,
		widget *GtkWidget,
		left_attach Guint,
		right_attach Guint,
		top_attach Guint,
		bottom_attach Guint)

	Gtk_table_set_row_spacing func(
		table *GtkTable,
		row Guint,
		spacing Guint)

	Gtk_table_get_row_spacing func(
		table *GtkTable,
		row Guint) Guint

	Gtk_table_set_col_spacing func(
		table *GtkTable,
		column Guint,
		spacing Guint)

	Gtk_table_get_col_spacing func(
		table *GtkTable,
		column Guint) Guint

	Gtk_table_set_row_spacings func(
		table *GtkTable,
		spacing Guint)

	Gtk_table_get_default_row_spacing func(
		table *GtkTable) Guint

	Gtk_table_set_col_spacings func(
		table *GtkTable,
		spacing Guint)

	Gtk_table_get_default_col_spacing func(
		table *GtkTable) Guint

	Gtk_table_set_homogeneous func(
		table *GtkTable,
		homogeneous Gboolean)

	Gtk_table_get_homogeneous func(
		table *GtkTable) Gboolean

	Gtk_table_get_size func(
		table *GtkTable,
		rows *Guint,
		columns *Guint)

	Gtk_text_tag_table_get_type func() GType

	Gtk_text_tag_table_new func() *GtkTextTagTable

	Gtk_text_tag_table_add func(
		table *GtkTextTagTable,
		tag *GtkTextTag)

	Gtk_text_tag_table_remove func(
		table *GtkTextTagTable,
		tag *GtkTextTag)

	Gtk_text_tag_table_lookup func(
		table *GtkTextTagTable,
		name string) *GtkTextTag

	Gtk_text_tag_table_foreach func(
		table *GtkTextTagTable,
		f GtkTextTagTableForeach,
		data Gpointer)

	Gtk_text_tag_table_get_size func(
		table *GtkTextTagTable) Gint

	Gtk_text_mark_get_type func() GType

	Gtk_text_mark_set_visible func(
		mark *GtkTextMark,
		setting Gboolean)

	Gtk_text_mark_get_visible func(
		mark *GtkTextMark) Gboolean

	Gtk_text_mark_new func(
		name string,
		left_gravity Gboolean) *GtkTextMark

	Gtk_text_mark_get_name func(
		mark *GtkTextMark) string

	Gtk_text_mark_get_deleted func(
		mark *GtkTextMark) Gboolean

	Gtk_text_mark_get_buffer func(
		mark *GtkTextMark) *GtkTextBuffer

	Gtk_text_mark_get_left_gravity func(
		mark *GtkTextMark) Gboolean

	Gtk_text_buffer_get_type func() GType

	Gtk_text_buffer_new func(
		table *GtkTextTagTable) *GtkTextBuffer

	Gtk_text_buffer_get_line_count func(
		buffer *GtkTextBuffer) Gint

	Gtk_text_buffer_get_char_count func(
		buffer *GtkTextBuffer) Gint

	Gtk_text_buffer_get_tag_table func(
		buffer *GtkTextBuffer) *GtkTextTagTable

	Gtk_text_buffer_set_text func(
		buffer *GtkTextBuffer,
		text string,
		len Gint)

	Gtk_text_buffer_insert func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		text string,
		len Gint)

	Gtk_text_buffer_insert_at_cursor func(
		buffer *GtkTextBuffer,
		text string,
		len Gint)

	Gtk_text_buffer_insert_interactive func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		text string,
		len Gint,
		default_editable Gboolean) Gboolean

	Gtk_text_buffer_insert_interactive_at_cursor func(
		buffer *GtkTextBuffer,
		text string,
		len Gint,
		default_editable Gboolean) Gboolean

	Gtk_text_buffer_insert_range func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_text_buffer_insert_range_interactive func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		start *GtkTextIter,
		end *GtkTextIter,
		default_editable Gboolean) Gboolean

	Gtk_text_buffer_insert_with_tags func(buffer *GtkTextBuffer,
		iter *GtkTextIter, text string, len Gint,
		first_tag *GtkTextTag, v ...VArg)

	Gtk_text_buffer_insert_with_tags_by_name func(
		buffer *GtkTextBuffer, iter *GtkTextIter, text string,
		len Gint, first_tag_name string, v ...VArg)

	Gtk_text_buffer_delete func(
		buffer *GtkTextBuffer,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_text_buffer_delete_interactive func(
		buffer *GtkTextBuffer,
		start_iter *GtkTextIter,
		end_iter *GtkTextIter,
		default_editable Gboolean) Gboolean

	Gtk_text_buffer_backspace func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		interactive Gboolean,
		default_editable Gboolean) Gboolean

	Gtk_text_buffer_get_text func(
		buffer *GtkTextBuffer,
		start *GtkTextIter,
		end *GtkTextIter,
		include_hidden_chars Gboolean) string

	Gtk_text_buffer_get_slice func(
		buffer *GtkTextBuffer,
		start *GtkTextIter,
		end *GtkTextIter,
		include_hidden_chars Gboolean) string

	Gtk_text_buffer_insert_pixbuf func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		pixbuf *GdkPixbuf)

	Gtk_text_buffer_insert_child_anchor func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		anchor *GtkTextChildAnchor)

	Gtk_text_buffer_create_child_anchor func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter) *GtkTextChildAnchor

	Gtk_text_buffer_add_mark func(
		buffer *GtkTextBuffer,
		mark *GtkTextMark,
		where *GtkTextIter)

	Gtk_text_buffer_create_mark func(
		buffer *GtkTextBuffer,
		mark_name string,
		where *GtkTextIter,
		left_gravity Gboolean) *GtkTextMark

	Gtk_text_buffer_move_mark func(
		buffer *GtkTextBuffer,
		mark *GtkTextMark,
		where *GtkTextIter)

	Gtk_text_buffer_delete_mark func(
		buffer *GtkTextBuffer,
		mark *GtkTextMark)

	Gtk_text_buffer_get_mark func(
		buffer *GtkTextBuffer,
		name string) *GtkTextMark

	Gtk_text_buffer_move_mark_by_name func(
		buffer *GtkTextBuffer,
		name string,
		where *GtkTextIter)

	Gtk_text_buffer_delete_mark_by_name func(
		buffer *GtkTextBuffer,
		name string)

	Gtk_text_buffer_get_insert func(
		buffer *GtkTextBuffer) *GtkTextMark

	Gtk_text_buffer_get_selection_bound func(
		buffer *GtkTextBuffer) *GtkTextMark

	Gtk_text_buffer_place_cursor func(
		buffer *GtkTextBuffer,
		where *GtkTextIter)

	Gtk_text_buffer_select_range func(
		buffer *GtkTextBuffer,
		ins *GtkTextIter,
		bound *GtkTextIter)

	Gtk_text_buffer_apply_tag func(
		buffer *GtkTextBuffer,
		tag *GtkTextTag,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_text_buffer_remove_tag func(
		buffer *GtkTextBuffer,
		tag *GtkTextTag,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_text_buffer_apply_tag_by_name func(
		buffer *GtkTextBuffer,
		name string,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_text_buffer_remove_tag_by_name func(
		buffer *GtkTextBuffer,
		name string,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_text_buffer_remove_all_tags func(
		buffer *GtkTextBuffer,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_text_buffer_create_tag func(
		buffer *GtkTextBuffer, tag_name string,
		first_property_name string, v ...VArg) *GtkTextTag

	Gtk_text_buffer_get_iter_at_line_offset func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		line_number Gint,
		char_offset Gint)

	Gtk_text_buffer_get_iter_at_line_index func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		line_number Gint,
		byte_index Gint)

	Gtk_text_buffer_get_iter_at_offset func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		char_offset Gint)

	Gtk_text_buffer_get_iter_at_line func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		line_number Gint)

	Gtk_text_buffer_get_start_iter func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter)

	Gtk_text_buffer_get_end_iter func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter)

	Gtk_text_buffer_get_bounds func(
		buffer *GtkTextBuffer,
		start *GtkTextIter,
		end *GtkTextIter)

	Gtk_text_buffer_get_iter_at_mark func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		mark *GtkTextMark)

	Gtk_text_buffer_get_iter_at_child_anchor func(
		buffer *GtkTextBuffer,
		iter *GtkTextIter,
		anchor *GtkTextChildAnchor)

	Gtk_text_buffer_get_modified func(
		buffer *GtkTextBuffer) Gboolean

	Gtk_text_buffer_set_modified func(
		buffer *GtkTextBuffer,
		setting Gboolean)

	Gtk_text_buffer_get_has_selection func(
		buffer *GtkTextBuffer) Gboolean

	Gtk_text_buffer_add_selection_clipboard func(
		buffer *GtkTextBuffer,
		clipboard *GtkClipboard)

	Gtk_text_buffer_remove_selection_clipboard func(
		buffer *GtkTextBuffer,
		clipboard *GtkClipboard)

	Gtk_text_buffer_cut_clipboard func(
		buffer *GtkTextBuffer,
		clipboard *GtkClipboard,
		default_editable Gboolean)

	Gtk_text_buffer_copy_clipboard func(
		buffer *GtkTextBuffer,
		clipboard *GtkClipboard)

	Gtk_text_buffer_paste_clipboard func(
		buffer *GtkTextBuffer,
		clipboard *GtkClipboard,
		override_location *GtkTextIter,
		default_editable Gboolean)

	Gtk_text_buffer_get_selection_bounds func(
		buffer *GtkTextBuffer,
		start *GtkTextIter,
		end *GtkTextIter) Gboolean

	Gtk_text_buffer_delete_selection func(
		buffer *GtkTextBuffer,
		interactive Gboolean,
		default_editable Gboolean) Gboolean

	Gtk_text_buffer_begin_user_action func(
		buffer *GtkTextBuffer)

	Gtk_text_buffer_end_user_action func(
		buffer *GtkTextBuffer)

	Gtk_text_buffer_get_copy_target_list func(
		buffer *GtkTextBuffer) *GtkTargetList

	Gtk_text_buffer_get_paste_target_list func(
		buffer *GtkTextBuffer) *GtkTargetList

	Gtk_text_buffer_register_serialize_format func(
		buffer *GtkTextBuffer,
		mime_type string,
		function GtkTextBufferSerializeFunc,
		user_dataGpointer,
		user_data_destroy GDestroyNotify) GdkAtom

	Gtk_text_buffer_register_serialize_tagset func(
		buffer *GtkTextBuffer,
		tagset_name string) GdkAtom

	Gtk_text_buffer_register_deserialize_format func(
		buffer *GtkTextBuffer,
		mime_type string,
		function GtkTextBufferDeserializeFunc,
		user_dataGpointer,
		user_data_destroy GDestroyNotify) GdkAtom

	Gtk_text_buffer_register_deserialize_tagset func(
		buffer *GtkTextBuffer,
		tagset_name string) GdkAtom

	Gtk_text_buffer_unregister_serialize_format func(
		buffer *GtkTextBuffer,
		format GdkAtom)

	Gtk_text_buffer_unregister_deserialize_format func(
		buffer *GtkTextBuffer,
		format GdkAtom)

	Gtk_text_buffer_deserialize_set_can_create_tags func(
		buffer *GtkTextBuffer,
		format GdkAtom,
		can_create_tags Gboolean)

	Gtk_text_buffer_deserialize_get_can_create_tags func(
		buffer *GtkTextBuffer,
		format GdkAtom) Gboolean

	Gtk_text_buffer_get_serialize_formats func(
		buffer *GtkTextBuffer,
		n_formats *Gint) *GdkAtom

	Gtk_text_buffer_get_deserialize_formats func(
		buffer *GtkTextBuffer,
		n_formats *Gint) *GdkAtom

	Gtk_text_buffer_serialize func(
		register_buffer *GtkTextBuffer,
		content_buffer *GtkTextBuffer,
		format GdkAtom,
		start *GtkTextIter,
		end *GtkTextIter,
		length *Gsize) *Guint8

	Gtk_text_buffer_deserialize func(
		register_buffer *GtkTextBuffer,
		content_buffer *GtkTextBuffer,
		format GdkAtom,
		iter *GtkTextIter,
		data *Guint8,
		length Gsize,
		error **GError) Gboolean

	Gtk_text_view_get_type func() GType

	Gtk_text_view_new func() *GtkWidget

	Gtk_text_view_new_with_buffer func(buffer *GtkTextBuffer) *GtkWidget

	Gtk_text_view_set_buffer func(text_view *GtkTextView,
		buffer *GtkTextBuffer)

	Gtk_text_view_get_buffer func(text_view *GtkTextView) *GtkTextBuffer

	Gtk_text_view_scroll_to_iter func(text_view *GtkTextView,
		iter *GtkTextIter,
		within_margin Gdouble,
		use_align Gboolean,
		xalign Gdouble,
		yalign Gdouble) Gboolean

	Gtk_text_view_scroll_to_mark func(text_view *GtkTextView,
		mark *GtkTextMark,
		within_margin Gdouble,
		use_align Gboolean,
		xalign Gdouble,
		yalign Gdouble)

	Gtk_text_view_scroll_mark_onscreen func(text_view *GtkTextView,
		mark *GtkTextMark)

	Gtk_text_view_move_mark_onscreen func(text_view *GtkTextView,
		mark *GtkTextMark) Gboolean

	Gtk_text_view_place_cursor_onscreen func(text_view *GtkTextView) Gboolean

	Gtk_text_view_get_visible_rect func(text_view *GtkTextView,
		visible_rect *GdkRectangle)

	Gtk_text_view_set_cursor_visible func(text_view *GtkTextView,
		setting Gboolean)

	Gtk_text_view_get_cursor_visible func(text_view *GtkTextView) Gboolean

	Gtk_text_view_get_iter_location func(text_view *GtkTextView,
		iter *GtkTextIter,
		location *GdkRectangle)

	Gtk_text_view_get_iter_at_location func(text_view *GtkTextView,
		iter *GtkTextIter,
		x Gint,
		y Gint)

	Gtk_text_view_get_iter_at_position func(text_view *GtkTextView,
		iter *GtkTextIter,
		trailing *Gint,
		x Gint,
		y Gint)

	Gtk_text_view_get_line_yrange func(text_view *GtkTextView,
		iter *GtkTextIter,
		y *Gint,
		height *Gint)

	Gtk_text_view_get_line_at_y func(text_view *GtkTextView,
		target_iter *GtkTextIter,
		y Gint,
		line_top *Gint)

	Gtk_text_view_buffer_to_window_coords func(text_view *GtkTextView,
		win GtkTextWindowType,
		buffer_x Gint,
		buffer_y Gint,
		window_x *Gint,
		window_y *Gint)

	Gtk_text_view_window_to_buffer_coords func(text_view *GtkTextView,
		win GtkTextWindowType,
		window_x Gint,
		window_y Gint,
		buffer_x *Gint,
		buffer_y *Gint)

	Gtk_text_view_get_hadjustment func(text_view *GtkTextView) *GtkAdjustment

	Gtk_text_view_get_vadjustment func(text_view *GtkTextView) *GtkAdjustment

	Gtk_text_view_get_window func(text_view *GtkTextView,
		win GtkTextWindowType) *GdkWindow

	Gtk_text_view_get_window_type func(text_view *GtkTextView,
		window *GdkWindow) GtkTextWindowType

	Gtk_text_view_set_border_window_size func(text_view *GtkTextView,
		t GtkTextWindowType,
		size Gint)

	Gtk_text_view_get_border_window_size func(text_view *GtkTextView,
		t GtkTextWindowType) Gint

	Gtk_text_view_forward_display_line func(text_view *GtkTextView,
		iter *GtkTextIter) Gboolean

	Gtk_text_view_backward_display_line func(text_view *GtkTextView,
		iter *GtkTextIter) Gboolean

	Gtk_text_view_forward_display_line_end func(text_view *GtkTextView,
		iter *GtkTextIter) Gboolean

	Gtk_text_view_backward_display_line_start func(text_view *GtkTextView,
		iter *GtkTextIter) Gboolean

	Gtk_text_view_starts_display_line func(text_view *GtkTextView,
		iter *GtkTextIter) Gboolean

	Gtk_text_view_move_visually func(text_view *GtkTextView,
		iter *GtkTextIter,
		count Gint) Gboolean

	Gtk_text_view_im_context_filter_keypress func(text_view *GtkTextView,
		event *GdkEventKey) Gboolean

	Gtk_text_view_reset_im_context func(text_view *GtkTextView)

	Gtk_text_view_add_child_at_anchor func(text_view *GtkTextView,
		child *GtkWidget,
		anchor *GtkTextChildAnchor)

	Gtk_text_view_add_child_in_window func(text_view *GtkTextView,
		child *GtkWidget,
		which_window GtkTextWindowType,
		xpos Gint,
		ypos Gint)

	Gtk_text_view_move_child func(text_view *GtkTextView,
		child *GtkWidget,
		xpos Gint,
		ypos Gint)

	Gtk_text_view_set_wrap_mode func(text_view *GtkTextView,
		wrap_mode GtkWrapMode)

	Gtk_text_view_get_wrap_mode func(text_view *GtkTextView) GtkWrapMode

	Gtk_text_view_set_editable func(text_view *GtkTextView,
		setting Gboolean)

	Gtk_text_view_get_editable func(text_view *GtkTextView) Gboolean

	Gtk_text_view_set_overwrite func(text_view *GtkTextView,
		overwrite Gboolean)

	Gtk_text_view_get_overwrite func(text_view *GtkTextView) Gboolean

	Gtk_text_view_set_accepts_tab func(text_view *GtkTextView,
		accepts_tab Gboolean)

	Gtk_text_view_get_accepts_tab func(text_view *GtkTextView) Gboolean

	Gtk_text_view_set_pixels_above_lines func(text_view *GtkTextView,
		pixels_above_lines Gint)

	Gtk_text_view_get_pixels_above_lines func(text_view *GtkTextView) Gint

	Gtk_text_view_set_pixels_below_lines func(text_view *GtkTextView,
		pixels_below_lines Gint)

	Gtk_text_view_get_pixels_below_lines func(text_view *GtkTextView) Gint

	Gtk_text_view_set_pixels_inside_wrap func(text_view *GtkTextView,
		pixels_inside_wrap Gint)

	Gtk_text_view_get_pixels_inside_wrap func(text_view *GtkTextView) Gint

	Gtk_text_view_set_justification func(text_view *GtkTextView,
		justification GtkJustification)

	Gtk_text_view_get_justification func(text_view *GtkTextView) GtkJustification

	Gtk_text_view_set_left_margin func(text_view *GtkTextView,
		left_margin Gint)

	Gtk_text_view_get_left_margin func(text_view *GtkTextView) Gint

	Gtk_text_view_set_right_margin func(text_view *GtkTextView,
		right_margin Gint)

	Gtk_text_view_get_right_margin func(text_view *GtkTextView) Gint

	Gtk_text_view_set_indent func(text_view *GtkTextView,
		indent Gint)

	Gtk_text_view_get_indent func(text_view *GtkTextView) Gint

	Gtk_text_view_set_tabs func(text_view *GtkTextView,
		tabs *PangoTabArray)

	Gtk_text_view_get_tabs func(text_view *GtkTextView) *PangoTabArray

	Gtk_text_view_get_default_attributes func(text_view *GtkTextView) *GtkTextAttributes

	Gtk_pixmap_get_type func() GType

	Gtk_pixmap_new func(pixmap *GdkPixmap,
		mask *GdkBitmap) *GtkWidget

	Gtk_pixmap_set func(pixmap *GtkPixmap,
		val *GdkPixmap,
		mask *GdkBitmap)

	Gtk_pixmap_get func(pixmap *GtkPixmap,
		val **GdkPixmap,
		mask **GdkBitmap)

	Gtk_pixmap_set_build_insensitive func(pixmap *GtkPixmap,
		build Gboolean)

	Gtk_toolbar_get_type func() GType

	Gtk_toolbar_new func() *GtkWidget

	Gtk_toolbar_insert func(toolbar *GtkToolbar,
		item *GtkToolItem,
		pos Gint)

	Gtk_toolbar_get_item_index func(toolbar *GtkToolbar,
		item *GtkToolItem) Gint

	Gtk_toolbar_get_n_items func(toolbar *GtkToolbar) Gint

	Gtk_toolbar_get_nth_item func(toolbar *GtkToolbar,
		n Gint) *GtkToolItem

	Gtk_toolbar_get_show_arrow func(toolbar *GtkToolbar) Gboolean

	Gtk_toolbar_set_show_arrow func(toolbar *GtkToolbar,
		show_arrow Gboolean)

	Gtk_toolbar_get_style func(toolbar *GtkToolbar) GtkToolbarStyle

	Gtk_toolbar_set_style func(toolbar *GtkToolbar,
		style GtkToolbarStyle)

	Gtk_toolbar_unset_style func(toolbar *GtkToolbar)

	Gtk_toolbar_get_icon_size func(toolbar *GtkToolbar) GtkIconSize

	Gtk_toolbar_set_icon_size func(toolbar *GtkToolbar,
		icon_size GtkIconSize)

	Gtk_toolbar_unset_icon_size func(toolbar *GtkToolbar)

	Gtk_toolbar_get_relief_style func(toolbar *GtkToolbar) GtkReliefStyle

	Gtk_toolbar_get_drop_index func(toolbar *GtkToolbar,
		x Gint,
		y Gint) Gint

	Gtk_toolbar_set_drop_highlight_item func(toolbar *GtkToolbar,
		tool_item *GtkToolItem,
		index_ Gint)

	Gtk_toolbar_get_orientation func(toolbar *GtkToolbar) GtkOrientation

	Gtk_toolbar_set_orientation func(toolbar *GtkToolbar,
		orientation GtkOrientation)

	Gtk_toolbar_get_tooltips func(toolbar *GtkToolbar) Gboolean

	Gtk_toolbar_set_tooltips func(toolbar *GtkToolbar,
		enable Gboolean)

	Gtk_toolbar_append_item func(toolbar *GtkToolbar,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *GtkWidget,
		callback GCallback,
		user_data Gpointer) *GtkWidget

	Gtk_toolbar_prepend_item func(toolbar *GtkToolbar,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *GtkWidget,
		callback GCallback,
		user_data Gpointer) *GtkWidget

	Gtk_toolbar_insert_item func(toolbar *GtkToolbar,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *GtkWidget,
		callback GCallback,
		user_dataGpointer,
		position Gint) *GtkWidget

	Gtk_toolbar_insert_stock func(toolbar *GtkToolbar,
		stock_id string,
		tooltip_text string,
		tooltip_private_text string,
		callback GCallback,
		user_dataGpointer,
		position Gint) *GtkWidget

	Gtk_toolbar_append_space func(toolbar *GtkToolbar)

	Gtk_toolbar_prepend_space func(toolbar *GtkToolbar)

	Gtk_toolbar_insert_space func(toolbar *GtkToolbar,
		position Gint)

	Gtk_toolbar_remove_space func(toolbar *GtkToolbar,
		position Gint)

	Gtk_toolbar_append_element func(toolbar *GtkToolbar,
		t GtkToolbarChildType,
		widget *GtkWidget,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *GtkWidget,
		callback GCallback,
		user_data Gpointer) *GtkWidget

	Gtk_toolbar_prepend_element func(toolbar *GtkToolbar,
		t GtkToolbarChildType,
		widget *GtkWidget,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *GtkWidget,
		callback GCallback,
		user_data Gpointer) *GtkWidget

	Gtk_toolbar_insert_element func(toolbar *GtkToolbar,
		t GtkToolbarChildType,
		widget *GtkWidget,
		text string,
		tooltip_text string,
		tooltip_private_text string,
		icon *GtkWidget,
		callback GCallback,
		user_dataGpointer,
		position Gint) *GtkWidget

	Gtk_toolbar_append_widget func(toolbar *GtkToolbar,
		widget *GtkWidget,
		tooltip_text string,
		tooltip_private_text string)

	Gtk_toolbar_prepend_widget func(toolbar *GtkToolbar,
		widget *GtkWidget,
		tooltip_text string,
		tooltip_private_text string)

	Gtk_toolbar_insert_widget func(toolbar *GtkToolbar,
		widget *GtkWidget,
		tooltip_text string,
		tooltip_private_text string,
		position Gint)

	Gtk_tool_item_group_get_type func() GType

	Gtk_tool_item_group_new func(label string) *GtkWidget

	Gtk_tool_item_group_set_label func(group *GtkToolItemGroup,
		label string)

	Gtk_tool_item_group_set_label_widget func(group *GtkToolItemGroup,
		label_widget *GtkWidget)

	Gtk_tool_item_group_set_collapsed func(group *GtkToolItemGroup,
		collapsed Gboolean)

	Gtk_tool_item_group_set_ellipsize func(group *GtkToolItemGroup,
		ellipsize PangoEllipsizeMode)

	Gtk_tool_item_group_set_header_relief func(group *GtkToolItemGroup,
		style GtkReliefStyle)

	Gtk_tool_item_group_get_label func(group *GtkToolItemGroup) string

	Gtk_tool_item_group_get_label_widget func(group *GtkToolItemGroup) *GtkWidget

	Gtk_tool_item_group_get_collapsed func(group *GtkToolItemGroup) Gboolean

	Gtk_tool_item_group_get_ellipsize func(group *GtkToolItemGroup) PangoEllipsizeMode

	Gtk_tool_item_group_get_header_relief func(group *GtkToolItemGroup) GtkReliefStyle

	Gtk_tool_item_group_insert func(group *GtkToolItemGroup,
		item *GtkToolItem,
		position Gint)

	Gtk_tool_item_group_set_item_position func(group *GtkToolItemGroup,
		item *GtkToolItem,
		position Gint)

	Gtk_tool_item_group_get_item_position func(group *GtkToolItemGroup,
		item *GtkToolItem) Gint

	Gtk_tool_item_group_get_n_items func(group *GtkToolItemGroup) Guint

	Gtk_tool_item_group_get_nth_item func(group *GtkToolItemGroup,
		index Guint) *GtkToolItem

	Gtk_tool_item_group_get_drop_item func(group *GtkToolItemGroup,
		x Gint,
		y Gint) *GtkToolItem

	Gtk_tool_palette_get_type func() GType

	Gtk_tool_palette_new func() *GtkWidget

	Gtk_tool_palette_set_group_position func(palette *GtkToolPalette,
		group *GtkToolItemGroup,
		position Gint)

	Gtk_tool_palette_set_exclusive func(palette *GtkToolPalette,
		group *GtkToolItemGroup,
		exclusive Gboolean)

	Gtk_tool_palette_set_expand func(palette *GtkToolPalette,
		group *GtkToolItemGroup,
		expand Gboolean)

	Gtk_tool_palette_get_group_position func(palette *GtkToolPalette,
		group *GtkToolItemGroup) Gint

	Gtk_tool_palette_get_exclusive func(palette *GtkToolPalette,
		group *GtkToolItemGroup) Gboolean

	Gtk_tool_palette_get_expand func(palette *GtkToolPalette,
		group *GtkToolItemGroup) Gboolean

	Gtk_tool_palette_set_icon_size func(palette *GtkToolPalette,
		icon_size GtkIconSize)

	Gtk_tool_palette_unset_icon_size func(palette *GtkToolPalette)

	Gtk_tool_palette_set_style func(palette *GtkToolPalette,
		style GtkToolbarStyle)

	Gtk_tool_palette_unset_style func(palette *GtkToolPalette)

	Gtk_tool_palette_get_icon_size func(palette *GtkToolPalette) GtkIconSize

	Gtk_tool_palette_get_style func(palette *GtkToolPalette) GtkToolbarStyle

	Gtk_tool_palette_get_drop_item func(palette *GtkToolPalette,
		x Gint,
		y Gint) *GtkToolItem

	Gtk_tool_palette_get_drop_group func(palette *GtkToolPalette,
		x Gint,
		y Gint) *GtkToolItemGroup

	Gtk_tool_palette_get_drag_item func(palette *GtkToolPalette,
		selection *GtkSelectionData) *GtkWidget

	Gtk_tool_palette_set_drag_source func(palette *GtkToolPalette,
		targets GtkToolPaletteDragTargets)

	Gtk_tool_palette_add_drag_dest func(palette *GtkToolPalette,
		widget *GtkWidget,
		flags GtkDestDefaults,
		targets GtkToolPaletteDragTargets,
		actions GdkDragAction)

	Gtk_tool_palette_get_hadjustment func(palette *GtkToolPalette) *GtkAdjustment

	Gtk_tool_palette_get_vadjustment func(palette *GtkToolPalette) *GtkAdjustment

	Gtk_tool_palette_get_drag_target_item func() *GtkTargetEntry

	Gtk_tool_palette_get_drag_target_group func() *GtkTargetEntry

	Gtk_tool_shell_get_type func() GType

	Gtk_tool_shell_get_icon_size func(shell *GtkToolShell) GtkIconSize

	Gtk_tool_shell_get_orientation func(shell *GtkToolShell) GtkOrientation

	Gtk_tool_shell_get_style func(shell *GtkToolShell) GtkToolbarStyle

	Gtk_tool_shell_get_relief_style func(shell *GtkToolShell) GtkReliefStyle

	Gtk_tool_shell_rebuild_menu func(shell *GtkToolShell)

	Gtk_tool_shell_get_text_orientation func(shell *GtkToolShell) GtkOrientation

	Gtk_tool_shell_get_text_alignment func(shell *GtkToolShell) Gfloat

	Gtk_tool_shell_get_ellipsize_mode func(shell *GtkToolShell) PangoEllipsizeMode

	Gtk_tool_shell_get_text_size_group func(shell *GtkToolShell) *GtkSizeGroup

	Gtk_test_init func(argcp *int, argvp ***Char, v ...VArg)

	Gtk_test_register_all_types func()

	Gtk_test_list_all_types func(n_types *Guint) *GType

	Gtk_test_find_widget func(widget *GtkWidget,
		label_pattern string,
		widget_type GType) *GtkWidget

	Gtk_test_create_widget func(widget_type GType,
		first_property_name string, v ...VArg) *GtkWidget

	Gtk_test_create_simple_window func(window_title string,
		dialog_text string) *GtkWidget

	Gtk_test_display_button_window func(window_title string,
		dialog_text string, v ...VArg) *GtkWidget

	Gtk_test_slider_set_perc func(widget *GtkWidget,
		percentage Double)

	Gtk_test_slider_get_value func(widget *GtkWidget) Double

	Gtk_test_spin_button_click func(spinner *GtkSpinButton,
		button Guint,
		upwards Gboolean) Gboolean

	Gtk_test_widget_click func(widget *GtkWidget,
		button Guint,
		modifiers GdkModifierType) Gboolean

	Gtk_test_widget_send_key func(widget *GtkWidget,
		keyval Guint,
		modifiers GdkModifierType) Gboolean

	Gtk_test_text_set func(widget *GtkWidget,
		string string)

	Gtk_test_text_get func(widget *GtkWidget) string

	Gtk_test_find_sibling func(base_widget *GtkWidget,
		widget_type GType) *GtkWidget

	Gtk_test_find_label func(widget *GtkWidget,
		label_pattern string) *GtkWidget

	Gtk_tree_drag_source_get_type func() GType

	Gtk_tree_drag_source_row_draggable func(drag_source *GtkTreeDragSource,
		path *GtkTreePath) Gboolean

	Gtk_tree_drag_source_drag_data_delete func(drag_source *GtkTreeDragSource,
		path *GtkTreePath) Gboolean

	Gtk_tree_drag_source_drag_data_get func(drag_source *GtkTreeDragSource,
		path *GtkTreePath,
		selection_data *GtkSelectionData) Gboolean

	Gtk_tree_drag_dest_get_type func() GType

	Gtk_tree_drag_dest_drag_data_received func(drag_dest *GtkTreeDragDest,
		dest *GtkTreePath,
		selection_data *GtkSelectionData) Gboolean

	Gtk_tree_drag_dest_row_drop_possible func(drag_dest *GtkTreeDragDest,
		dest_path *GtkTreePath,
		selection_data *GtkSelectionData) Gboolean

	Gtk_tree_set_row_drag_data func(selection_data *GtkSelectionData,
		tree_model *GtkTreeModel,
		path *GtkTreePath) Gboolean

	Gtk_tree_get_row_drag_data func(selection_data *GtkSelectionData,
		tree_model **GtkTreeModel,
		path **GtkTreePath) Gboolean

	Gtk_tree_model_sort_get_type func() GType

	Gtk_tree_model_sort_new_with_model func(child_model *GtkTreeModel) *GtkTreeModel

	Gtk_tree_model_sort_get_model func(tree_model *GtkTreeModelSort) *GtkTreeModel

	Gtk_tree_model_sort_convert_child_path_to_path func(tree_model_sort *GtkTreeModelSort,
		child_path *GtkTreePath) *GtkTreePath

	Gtk_tree_model_sort_convert_child_iter_to_iter func(tree_model_sort *GtkTreeModelSort,
		sort_iter *GtkTreeIter,
		child_iter *GtkTreeIter) Gboolean

	Gtk_tree_model_sort_convert_path_to_child_path func(tree_model_sort *GtkTreeModelSort,
		sorted_path *GtkTreePath) *GtkTreePath

	Gtk_tree_model_sort_convert_iter_to_child_iter func(tree_model_sort *GtkTreeModelSort,
		child_iter *GtkTreeIter,
		sorted_iter *GtkTreeIter)

	Gtk_tree_model_sort_reset_default_sort_func func(tree_model_sort *GtkTreeModelSort)

	Gtk_tree_model_sort_clear_cache func(tree_model_sort *GtkTreeModelSort)

	Gtk_tree_model_sort_iter_is_valid func(tree_model_sort *GtkTreeModelSort,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_selection_get_type func() GType

	Gtk_tree_selection_set_mode func(selection *GtkTreeSelection,
		typ GtkSelectionMode)

	Gtk_tree_selection_get_mode func(selection *GtkTreeSelection) GtkSelectionMode

	Gtk_tree_selection_set_select_function func(selection *GtkTreeSelection,
		f GtkTreeSelectionFunc,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_tree_selection_get_user_data func(selection *GtkTreeSelection) Gpointer

	Gtk_tree_selection_get_tree_view func(selection *GtkTreeSelection) *GtkTreeView

	Gtk_tree_selection_get_select_function func(selection *GtkTreeSelection) GtkTreeSelectionFunc

	Gtk_tree_selection_get_selected func(selection *GtkTreeSelection,
		model **GtkTreeModel,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_selection_get_selected_rows func(selection *GtkTreeSelection,
		model **GtkTreeModel) *GList

	Gtk_tree_selection_count_selected_rows func(selection *GtkTreeSelection) Gint

	Gtk_tree_selection_selected_foreach func(selection *GtkTreeSelection,
		f GtkTreeSelectionForeachFunc,
		data Gpointer)

	Gtk_tree_selection_select_path func(selection *GtkTreeSelection,
		path *GtkTreePath)

	Gtk_tree_selection_unselect_path func(selection *GtkTreeSelection,
		path *GtkTreePath)

	Gtk_tree_selection_select_iter func(selection *GtkTreeSelection,
		iter *GtkTreeIter)

	Gtk_tree_selection_unselect_iter func(selection *GtkTreeSelection,
		iter *GtkTreeIter)

	Gtk_tree_selection_path_is_selected func(selection *GtkTreeSelection,
		path *GtkTreePath) Gboolean

	Gtk_tree_selection_iter_is_selected func(selection *GtkTreeSelection,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_selection_select_all func(selection *GtkTreeSelection)

	Gtk_tree_selection_unselect_all func(selection *GtkTreeSelection)

	Gtk_tree_selection_select_range func(selection *GtkTreeSelection,
		start_path *GtkTreePath,
		end_path *GtkTreePath)

	Gtk_tree_selection_unselect_range func(selection *GtkTreeSelection,
		start_path *GtkTreePath,
		end_path *GtkTreePath)

	Gtk_tree_store_get_type func() GType

	Gtk_tree_store_new func(
		n_columns Gint, v ...VArg) *GtkTreeStore

	Gtk_tree_store_newv func(n_columns Gint,
		types *GType) *GtkTreeStore

	Gtk_tree_store_set_column_types func(tree_store *GtkTreeStore,
		n_columns Gint,
		types *GType)

	Gtk_tree_store_set_value func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		column Gint,
		value *GValue)

	Gtk_tree_store_set func(tree_store *GtkTreeStore,
		iter *GtkTreeIter, v ...VArg)

	Gtk_tree_store_set_valuesv func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		columns *Gint,
		values *GValue,
		n_values Gint)

	Gtk_tree_store_set_valist func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		var_args Va_list)

	Gtk_tree_store_remove func(tree_store *GtkTreeStore,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_store_insert func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		parent *GtkTreeIter,
		position Gint)

	Gtk_tree_store_insert_before func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		parent *GtkTreeIter,
		sibling *GtkTreeIter)

	Gtk_tree_store_insert_after func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		parent *GtkTreeIter,
		sibling *GtkTreeIter)

	Gtk_tree_store_insert_with_values func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		parent *GtkTreeIter,
		position Gint, v ...VArg)

	Gtk_tree_store_insert_with_valuesv func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		parent *GtkTreeIter,
		position Gint,
		columns *Gint,
		values *GValue,
		n_values Gint)

	Gtk_tree_store_prepend func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		parent *GtkTreeIter)

	Gtk_tree_store_append func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		parent *GtkTreeIter)

	Gtk_tree_store_is_ancestor func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		descendant *GtkTreeIter) Gboolean

	Gtk_tree_store_iter_depth func(tree_store *GtkTreeStore,
		iter *GtkTreeIter) Gint

	Gtk_tree_store_clear func(tree_store *GtkTreeStore)

	Gtk_tree_store_iter_is_valid func(tree_store *GtkTreeStore,
		iter *GtkTreeIter) Gboolean

	Gtk_tree_store_reorder func(tree_store *GtkTreeStore,
		parent *GtkTreeIter,
		new_order *Gint)

	Gtk_tree_store_swap func(tree_store *GtkTreeStore,
		a *GtkTreeIter,
		b *GtkTreeIter)

	Gtk_tree_store_move_before func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		position *GtkTreeIter)

	Gtk_tree_store_move_after func(tree_store *GtkTreeStore,
		iter *GtkTreeIter,
		position *GtkTreeIter)

	Gtk_ui_manager_get_type func() GType

	Gtk_ui_manager_new func() *GtkUIManager

	Gtk_ui_manager_set_add_tearoffs func(self *GtkUIManager,
		add_tearoffs Gboolean)

	Gtk_ui_manager_get_add_tearoffs func(self *GtkUIManager) Gboolean

	Gtk_ui_manager_insert_action_group func(self *GtkUIManager,
		action_group *GtkActionGroup,
		pos Gint)

	Gtk_ui_manager_remove_action_group func(self *GtkUIManager,
		action_group *GtkActionGroup)

	Gtk_ui_manager_get_action_groups func(self *GtkUIManager) *GList

	Gtk_ui_manager_get_accel_group func(self *GtkUIManager) *GtkAccelGroup

	Gtk_ui_manager_get_widget func(self *GtkUIManager,
		path string) *GtkWidget

	Gtk_ui_manager_get_toplevels func(self *GtkUIManager,
		types GtkUIManagerItemType) *GSList

	Gtk_ui_manager_get_action func(self *GtkUIManager,
		path string) *GtkAction

	Gtk_ui_manager_add_ui_from_string func(self *GtkUIManager,
		buffer string,
		length Gssize,
		err **GError) Guint

	Gtk_ui_manager_add_ui_from_file func(self *GtkUIManager,
		filename string,
		err **GError) Guint

	Gtk_ui_manager_add_ui func(self *GtkUIManager,
		merge_id Guint,
		path string,
		name string,
		action string,
		t GtkUIManagerItemType,
		top Gboolean)

	Gtk_ui_manager_remove_ui func(self *GtkUIManager,
		merge_id Guint)

	Gtk_ui_manager_get_ui func(self *GtkUIManager) string

	Gtk_ui_manager_ensure_update func(self *GtkUIManager)

	Gtk_ui_manager_new_merge_id func(self *GtkUIManager) Guint

	Gtk_vbutton_box_get_type func() GType

	Gtk_vbutton_box_new func() *GtkWidget

	Gtk_vbutton_box_get_spacing_default func() Gint

	Gtk_vbutton_box_set_spacing_default func(spacing Gint)

	Gtk_vbutton_box_get_layout_default func() GtkButtonBoxStyle

	Gtk_vbutton_box_set_layout_default func(layout GtkButtonBoxStyle)

	Gtk_volume_button_get_type func() GType

	Gtk_volume_button_new func() *GtkWidget

	Gtk_vpaned_get_type func() GType

	Gtk_vpaned_new func() *GtkWidget

	Gtk_vruler_get_type func() GType

	Gtk_vruler_new func() *GtkWidget

	Gtk_vscale_get_type func() GType

	Gtk_vscale_new func(adjustment *GtkAdjustment) *GtkWidget

	Gtk_vscale_new_with_range func(min Gdouble,
		max Gdouble,
		step Gdouble) *GtkWidget

	Gtk_vseparator_get_type func() GType

	Gtk_vseparator_new func() *GtkWidget

	Gtk_clist_get_type func() GType

	Gtk_clist_new func(columns Gint) *GtkWidget

	Gtk_clist_new_with_titles func(columns Gint,
		titles **Gchar) *GtkWidget

	Gtk_clist_set_hadjustment func(clist *GtkCList,
		adjustment *GtkAdjustment)

	Gtk_clist_set_vadjustment func(clist *GtkCList,
		adjustment *GtkAdjustment)

	Gtk_clist_get_hadjustment func(clist *GtkCList) *GtkAdjustment

	Gtk_clist_get_vadjustment func(clist *GtkCList) *GtkAdjustment

	Gtk_clist_set_shadow_type func(clist *GtkCList,
		t GtkShadowType)

	Gtk_clist_set_selection_mode func(clist *GtkCList,
		mode GtkSelectionMode)

	Gtk_clist_set_reorderable func(clist *GtkCList,
		reorderable Gboolean)

	Gtk_clist_set_use_drag_icons func(clist *GtkCList,
		use_icons Gboolean)

	Gtk_clist_set_button_actions func(clist *GtkCList,
		button Guint,
		button_actions Guint8)

	Gtk_clist_freeze func(clist *GtkCList)

	Gtk_clist_thaw func(clist *GtkCList)

	Gtk_clist_column_titles_show func(clist *GtkCList)

	Gtk_clist_column_titles_hide func(clist *GtkCList)

	Gtk_clist_column_title_active func(clist *GtkCList,
		column Gint)

	Gtk_clist_column_title_passive func(clist *GtkCList,
		column Gint)

	Gtk_clist_column_titles_active func(clist *GtkCList)

	Gtk_clist_column_titles_passive func(clist *GtkCList)

	Gtk_clist_set_column_title func(clist *GtkCList,
		column Gint,
		title string)

	Gtk_clist_get_column_title func(clist *GtkCList,
		column Gint) string

	Gtk_clist_set_column_widget func(clist *GtkCList,
		column Gint,
		widget *GtkWidget)

	Gtk_clist_get_column_widget func(clist *GtkCList,
		column Gint) *GtkWidget

	Gtk_clist_set_column_justification func(clist *GtkCList,
		column Gint,
		justification GtkJustification)

	Gtk_clist_set_column_visibility func(clist *GtkCList,
		column Gint,
		visible Gboolean)

	Gtk_clist_set_column_resizeable func(clist *GtkCList,
		column Gint,
		resizeable Gboolean)

	Gtk_clist_set_column_auto_resize func(clist *GtkCList,
		column Gint,
		auto_resize Gboolean)

	Gtk_clist_columns_autosize func(clist *GtkCList) Gint

	Gtk_clist_optimal_column_width func(clist *GtkCList,
		column Gint) Gint

	Gtk_clist_set_column_width func(clist *GtkCList,
		column Gint,
		width Gint)

	Gtk_clist_set_column_min_width func(clist *GtkCList,
		column Gint,
		min_width Gint)

	Gtk_clist_set_column_max_width func(clist *GtkCList,
		column Gint,
		max_width Gint)

	Gtk_clist_set_row_height func(clist *GtkCList,
		height Guint)

	Gtk_clist_moveto func(clist *GtkCList,
		row Gint,
		column Gint,
		row_align Gfloat,
		col_align Gfloat)

	Gtk_clist_row_is_visible func(clist *GtkCList,
		row Gint) GtkVisibility

	Gtk_clist_get_cell_type func(clist *GtkCList,
		row Gint,
		column Gint) GtkCellType

	Gtk_clist_set_text func(clist *GtkCList,
		row Gint,
		column Gint,
		text string)

	Gtk_clist_get_text func(clist *GtkCList,
		row Gint,
		column Gint,
		text **Gchar) Gint

	Gtk_clist_set_pixmap func(clist *GtkCList,
		row Gint,
		column Gint,
		pixmap *GdkPixmap,
		mask *GdkBitmap)

	Gtk_clist_get_pixmap func(clist *GtkCList,
		row Gint,
		column Gint,
		pixmap **GdkPixmap,
		mask **GdkBitmap) Gint

	Gtk_clist_set_pixtext func(clist *GtkCList,
		row Gint,
		column Gint,
		text string,
		spacing Guint8,
		pixmap *GdkPixmap,
		mask *GdkBitmap)

	Gtk_clist_get_pixtext func(clist *GtkCList,
		row Gint,
		column Gint,
		text **Gchar,
		spacing *Guint8,
		pixmap **GdkPixmap,
		mask **GdkBitmap) Gint

	Gtk_clist_set_foreground func(clist *GtkCList,
		row Gint,
		color *GdkColor)

	Gtk_clist_set_background func(clist *GtkCList,
		row Gint,
		color *GdkColor)

	Gtk_clist_set_cell_style func(clist *GtkCList,
		row Gint,
		column Gint,
		style *GtkStyle)

	Gtk_clist_get_cell_style func(clist *GtkCList,
		row Gint,
		column Gint) *GtkStyle

	Gtk_clist_set_row_style func(clist *GtkCList,
		row Gint,
		style *GtkStyle)

	Gtk_clist_get_row_style func(clist *GtkCList,
		row Gint) *GtkStyle

	Gtk_clist_set_shift func(clist *GtkCList,
		row Gint,
		column Gint,
		vertical Gint,
		horizontal Gint)

	Gtk_clist_set_selectable func(clist *GtkCList,
		row Gint,
		selectable Gboolean)

	Gtk_clist_get_selectable func(clist *GtkCList,
		row Gint) Gboolean

	Gtk_clist_prepend func(clist *GtkCList,
		text **Gchar) Gint

	Gtk_clist_append func(clist *GtkCList,
		text **Gchar) Gint

	Gtk_clist_insert func(clist *GtkCList,
		row Gint,
		text **Gchar) Gint

	Gtk_clist_remove func(clist *GtkCList,
		row Gint)

	Gtk_clist_set_row_data func(clist *GtkCList,
		row Gint,
		data Gpointer)

	Gtk_clist_set_row_data_full func(clist *GtkCList,
		row Gint,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_clist_get_row_data func(clist *GtkCList,
		row Gint) Gpointer

	Gtk_clist_find_row_from_data func(clist *GtkCList,
		data Gpointer) Gint

	Gtk_clist_select_row func(clist *GtkCList,
		row Gint,
		column Gint)

	Gtk_clist_unselect_row func(clist *GtkCList,
		row Gint,
		column Gint)

	Gtk_clist_undo_selection func(clist *GtkCList)

	Gtk_clist_clear func(clist *GtkCList)

	Gtk_clist_get_selection_info func(clist *GtkCList,
		x Gint,
		y Gint,
		row *Gint,
		column *Gint) Gint

	Gtk_clist_select_all func(clist *GtkCList)

	Gtk_clist_unselect_all func(clist *GtkCList)

	Gtk_clist_swap_rows func(clist *GtkCList,
		row1 Gint,
		row2 Gint)

	Gtk_clist_row_move func(clist *GtkCList,
		source_row Gint,
		dest_row Gint)

	Gtk_clist_set_compare_func func(clist *GtkCList,
		cmp_func GtkCListCompareFunc)

	Gtk_clist_set_sort_column func(clist *GtkCList,
		column Gint)

	Gtk_clist_set_sort_type func(clist *GtkCList,
		sort_type GtkSortType)

	Gtk_clist_sort func(clist *GtkCList)

	Gtk_clist_set_auto_sort func(clist *GtkCList,
		auto_sort Gboolean)

	Gtk_combo_get_type func() GType

	Gtk_combo_new func() *GtkWidget

	Gtk_combo_set_value_in_list func(combo *GtkCombo,
		val Gboolean,
		ok_if_empty Gboolean)

	Gtk_combo_set_use_arrows func(combo *GtkCombo,
		val Gboolean)

	Gtk_combo_set_use_arrows_always func(combo *GtkCombo,
		val Gboolean)

	Gtk_combo_set_case_sensitive func(combo *GtkCombo,
		val Gboolean)

	Gtk_combo_set_item_string func(combo *GtkCombo,
		item *GtkItem,
		item_value string)

	Gtk_combo_set_popdown_strings func(combo *GtkCombo,
		strings *GList)

	Gtk_combo_disable_activate func(combo *GtkCombo)

	Gtk_ctree_get_type func() GType

	Gtk_ctree_new_with_titles func(columns Gint,
		tree_column Gint,
		titles **Gchar) *GtkWidget

	Gtk_ctree_new func(columns Gint,
		tree_column Gint) *GtkWidget

	Gtk_ctree_insert_node func(ctree *GtkCTree,
		parent *GtkCTreeNode,
		sibling *GtkCTreeNode,
		text **Gchar,
		spacing Guint8,
		pixmap_closed *GdkPixmap,
		mask_closed *GdkBitmap,
		pixmap_opened *GdkPixmap,
		mask_opened *GdkBitmap,
		is_leaf Gboolean,
		expanded Gboolean) *GtkCTreeNode

	Gtk_ctree_remove_node func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_insert_gnode func(ctree *GtkCTree,
		parent *GtkCTreeNode,
		sibling *GtkCTreeNode,
		gnode *GNode,
		f GtkCTreeGNodeFunc,
		data Gpointer) *GtkCTreeNode

	Gtk_ctree_export_to_gnode func(ctree *GtkCTree,
		parent *GNode,
		sibling *GNode,
		node *GtkCTreeNode,
		f GtkCTreeGNodeFunc,
		data Gpointer) *GNode

	Gtk_ctree_post_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode,
		f GtkCTreeFunc,
		data Gpointer)

	Gtk_ctree_post_recursive_to_depth func(ctree *GtkCTree,
		node *GtkCTreeNode,
		depth Gint,
		f GtkCTreeFunc,
		data Gpointer)

	Gtk_ctree_pre_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode,
		f GtkCTreeFunc,
		data Gpointer)

	Gtk_ctree_pre_recursive_to_depth func(ctree *GtkCTree,
		node *GtkCTreeNode,
		depth Gint,
		f GtkCTreeFunc,
		data Gpointer)

	Gtk_ctree_is_viewable func(ctree *GtkCTree,
		node *GtkCTreeNode) Gboolean

	Gtk_ctree_last func(ctree *GtkCTree,
		node *GtkCTreeNode) *GtkCTreeNode

	Gtk_ctree_find_node_ptr func(ctree *GtkCTree,
		ctree_row *GtkCTreeRow) *GtkCTreeNode

	Gtk_ctree_node_nth func(ctree *GtkCTree,
		row Guint) *GtkCTreeNode

	Gtk_ctree_find func(ctree *GtkCTree,
		node *GtkCTreeNode,
		child *GtkCTreeNode) Gboolean

	Gtk_ctree_is_ancestor func(ctree *GtkCTree,
		node *GtkCTreeNode,
		child *GtkCTreeNode) Gboolean

	Gtk_ctree_find_by_row_data func(ctree *GtkCTree,
		node *GtkCTreeNode,
		data Gpointer) *GtkCTreeNode

	Gtk_ctree_find_all_by_row_data func(ctree *GtkCTree,
		node *GtkCTreeNode,
		data Gpointer) *GList

	Gtk_ctree_find_by_row_data_custom func(ctree *GtkCTree,
		node *GtkCTreeNode,
		dataGpointer,
		f GCompareFunc) *GtkCTreeNode

	Gtk_ctree_find_all_by_row_data_custom func(ctree *GtkCTree,
		node *GtkCTreeNode,
		dataGpointer,
		f GCompareFunc) *GList

	Gtk_ctree_is_hot_spot func(ctree *GtkCTree,
		x Gint,
		y Gint) Gboolean

	Gtk_ctree_move func(ctree *GtkCTree,
		node *GtkCTreeNode,
		new_parent *GtkCTreeNode,
		new_sibling *GtkCTreeNode)

	Gtk_ctree_expand func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_expand_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_expand_to_depth func(ctree *GtkCTree,
		node *GtkCTreeNode,
		depth Gint)

	Gtk_ctree_collapse func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_collapse_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_collapse_to_depth func(ctree *GtkCTree,
		node *GtkCTreeNode,
		depth Gint)

	Gtk_ctree_toggle_expansion func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_toggle_expansion_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_select func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_select_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_unselect func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_unselect_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_real_select_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode,
		state Gint)

	Gtk_ctree_node_set_text func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		text string)

	Gtk_ctree_node_set_pixmap func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		pixmap *GdkPixmap,
		mask *GdkBitmap)

	Gtk_ctree_node_set_pixtext func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		text string,
		spacing Guint8,
		pixmap *GdkPixmap,
		mask *GdkBitmap)

	Gtk_ctree_set_node_info func(ctree *GtkCTree,
		node *GtkCTreeNode,
		text string,
		spacing Guint8,
		pixmap_closed *GdkPixmap,
		mask_closed *GdkBitmap,
		pixmap_opened *GdkPixmap,
		mask_opened *GdkBitmap,
		is_leaf Gboolean,
		expanded Gboolean)

	Gtk_ctree_node_set_shift func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		vertical Gint,
		horizontal Gint)

	Gtk_ctree_node_set_selectable func(ctree *GtkCTree,
		node *GtkCTreeNode,
		selectable Gboolean)

	Gtk_ctree_node_get_selectable func(ctree *GtkCTree,
		node *GtkCTreeNode) Gboolean

	Gtk_ctree_node_get_cell_type func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint) GtkCellType

	Gtk_ctree_node_get_text func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		text **Gchar) Gboolean

	Gtk_ctree_node_get_pixmap func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		pixmap **GdkPixmap,
		mask **GdkBitmap) Gboolean

	Gtk_ctree_node_get_pixtext func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		text **Gchar,
		spacing *Guint8,
		pixmap **GdkPixmap,
		mask **GdkBitmap) Gboolean

	Gtk_ctree_get_node_info func(ctree *GtkCTree,
		node *GtkCTreeNode,
		text **Gchar,
		spacing *Guint8,
		pixmap_closed **GdkPixmap,
		mask_closed **GdkBitmap,
		pixmap_opened **GdkPixmap,
		mask_opened **GdkBitmap,
		is_leaf *Gboolean,
		expanded *Gboolean) Gboolean

	Gtk_ctree_node_set_row_style func(ctree *GtkCTree,
		node *GtkCTreeNode,
		style *GtkStyle)

	Gtk_ctree_node_get_row_style func(ctree *GtkCTree,
		node *GtkCTreeNode) *GtkStyle

	Gtk_ctree_node_set_cell_style func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		style *GtkStyle)

	Gtk_ctree_node_get_cell_style func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint) *GtkStyle

	Gtk_ctree_node_set_foreground func(ctree *GtkCTree,
		node *GtkCTreeNode,
		color *GdkColor)

	Gtk_ctree_node_set_background func(ctree *GtkCTree,
		node *GtkCTreeNode,
		color *GdkColor)

	Gtk_ctree_node_set_row_data func(ctree *GtkCTree,
		node *GtkCTreeNode,
		data Gpointer)

	Gtk_ctree_node_set_row_data_full func(ctree *GtkCTree,
		node *GtkCTreeNode,
		dataGpointer,
		destroy GDestroyNotify)

	Gtk_ctree_node_get_row_data func(ctree *GtkCTree,
		node *GtkCTreeNode) Gpointer

	Gtk_ctree_node_moveto func(ctree *GtkCTree,
		node *GtkCTreeNode,
		column Gint,
		row_align Gfloat,
		col_align Gfloat)

	Gtk_ctree_node_is_visible func(ctree *GtkCTree,
		node *GtkCTreeNode) GtkVisibility

	Gtk_ctree_set_indent func(ctree *GtkCTree,
		indent Gint)

	Gtk_ctree_set_spacing func(ctree *GtkCTree,
		spacing Gint)

	Gtk_ctree_set_show_stub func(ctree *GtkCTree,
		show_stub Gboolean)

	Gtk_ctree_set_line_style func(ctree *GtkCTree,
		line_style GtkCTreeLineStyle)

	Gtk_ctree_set_expander_style func(ctree *GtkCTree,
		expander_style GtkCTreeExpanderStyle)

	Gtk_ctree_set_drag_compare_func func(ctree *GtkCTree,
		cmp_func GtkCTreeCompareDragFunc)

	Gtk_ctree_sort_node func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_sort_recursive func(ctree *GtkCTree,
		node *GtkCTreeNode)

	Gtk_ctree_node_get_type func() GType

	Gtk_curve_get_type func() GType

	Gtk_curve_new func() *GtkWidget

	Gtk_curve_reset func(curve *GtkCurve)

	Gtk_curve_set_gamma func(curve *GtkCurve,
		gamma_ Gfloat)

	Gtk_curve_set_range func(curve *GtkCurve,
		min_x Gfloat,
		max_x Gfloat,
		min_y Gfloat,
		max_y Gfloat)

	Gtk_curve_get_vector func(curve *GtkCurve,
		veclen int,
		vector *Gfloat)

	Gtk_curve_set_vector func(curve *GtkCurve,
		veclen int,
		vector *Gfloat)

	Gtk_curve_set_curve_type func(curve *GtkCurve,
		t GtkCurveType)

	Gtk_file_selection_get_type func() GType

	Gtk_file_selection_new func(title string) *GtkWidget

	Gtk_file_selection_set_filename func(filesel *GtkFileSelection,
		filename string)

	Gtk_file_selection_get_filename func(filesel *GtkFileSelection) string

	Gtk_file_selection_complete func(filesel *GtkFileSelection,
		pattern string)

	Gtk_file_selection_show_fileop_buttons func(filesel *GtkFileSelection)

	Gtk_file_selection_hide_fileop_buttons func(filesel *GtkFileSelection)

	Gtk_file_selection_get_selections func(filesel *GtkFileSelection) **Gchar

	Gtk_file_selection_set_select_multiple func(filesel *GtkFileSelection,
		select_multiple Gboolean)

	Gtk_file_selection_get_select_multiple func(filesel *GtkFileSelection) Gboolean

	Gtk_gamma_curve_get_type func() GType

	Gtk_gamma_curve_new func() *GtkWidget

	Gtk_input_dialog_get_type func() GType

	Gtk_input_dialog_new func() *GtkWidget

	Gtk_item_factory_get_type func() GType

	Gtk_item_factory_new func(container_type GType,
		path string,
		accel_group *GtkAccelGroup) *GtkItemFactory

	Gtk_item_factory_construct func(ifactory *GtkItemFactory,
		container_type GType,
		path string,
		accel_group *GtkAccelGroup)

	Gtk_item_factory_add_foreign func(accel_widget *GtkWidget,
		full_path string,
		accel_group *GtkAccelGroup,
		keyval Guint,
		modifiers GdkModifierType)

	Gtk_item_factory_from_widget func(widget *GtkWidget) *GtkItemFactory

	Gtk_item_factory_path_from_widget func(widget *GtkWidget) string

	Gtk_item_factory_get_item func(ifactory *GtkItemFactory,
		path string) *GtkWidget

	Gtk_item_factory_get_widget func(ifactory *GtkItemFactory,
		path string) *GtkWidget

	Gtk_item_factory_get_widget_by_action func(ifactory *GtkItemFactory,
		action Guint) *GtkWidget

	Gtk_item_factory_get_item_by_action func(ifactory *GtkItemFactory,
		action Guint) *GtkWidget

	Gtk_item_factory_create_item func(ifactory *GtkItemFactory,
		entry *GtkItemFactoryEntry,
		callback_dataGpointer,
		callback_type Guint)

	Gtk_item_factory_create_items func(ifactory *GtkItemFactory,
		n_entries Guint,
		entries *GtkItemFactoryEntry,
		callback_data Gpointer)

	Gtk_item_factory_delete_item func(ifactory *GtkItemFactory,
		path string)

	Gtk_item_factory_delete_entry func(ifactory *GtkItemFactory,
		entry *GtkItemFactoryEntry)

	Gtk_item_factory_delete_entries func(ifactory *GtkItemFactory,
		n_entries Guint,
		entries *GtkItemFactoryEntry)

	Gtk_item_factory_popup func(ifactory *GtkItemFactory,
		x Guint,
		y Guint,
		mouse_button Guint,
		time_ Guint32)

	Gtk_item_factory_popup_with_data func(ifactory *GtkItemFactory,
		popup_dataGpointer,
		destroy GDestroyNotify,
		x Guint,
		y Guint,
		mouse_button Guint,
		time_ Guint32)

	Gtk_item_factory_popup_data func(ifactory *GtkItemFactory) Gpointer

	Gtk_item_factory_popup_data_from_widget func(widget *GtkWidget) Gpointer

	Gtk_item_factory_set_translate_func func(ifactory *GtkItemFactory,
		f GtkTranslateFunc,
		dataGpointer,
		notify GDestroyNotify)

	Gtk_item_factory_create_items_ac func(ifactory *GtkItemFactory,
		n_entries Guint,
		entries *GtkItemFactoryEntry,
		callback_dataGpointer,
		callback_type Guint)

	Gtk_item_factory_from_path func(path string) *GtkItemFactory

	Gtk_item_factory_create_menu_entries func(n_entries Guint,
		entries *GtkMenuEntry)

	Gtk_item_factories_path_delete func(ifactory_path string,
		path string)

	Gtk_list_get_type func() GType

	Gtk_list_new func() *GtkWidget

	Gtk_list_insert_items func(list *GtkList,
		items *GList,
		position Gint)

	Gtk_list_append_items func(list *GtkList,
		items *GList)

	Gtk_list_prepend_items func(list *GtkList,
		items *GList)

	Gtk_list_remove_items func(list *GtkList,
		items *GList)

	Gtk_list_remove_items_no_unref func(list *GtkList,
		items *GList)

	Gtk_list_clear_items func(list *GtkList,
		start Gint,
		end Gint)

	Gtk_list_select_item func(list *GtkList,
		item Gint)

	Gtk_list_unselect_item func(list *GtkList,
		item Gint)

	Gtk_list_select_child func(list *GtkList,
		child *GtkWidget)

	Gtk_list_unselect_child func(list *GtkList,
		child *GtkWidget)

	Gtk_list_child_position func(list *GtkList,
		child *GtkWidget) Gint

	Gtk_list_set_selection_mode func(list *GtkList,
		mode GtkSelectionMode)

	Gtk_list_extend_selection func(list *GtkList,
		scroll_type GtkScrollType,
		position Gfloat,
		auto_start_selection Gboolean)

	Gtk_list_start_selection func(list *GtkList)

	Gtk_list_end_selection func(list *GtkList)

	Gtk_list_select_all func(list *GtkList)

	Gtk_list_unselect_all func(list *GtkList)

	Gtk_list_scroll_horizontal func(list *GtkList,
		scroll_type GtkScrollType,
		position Gfloat)

	Gtk_list_scroll_vertical func(list *GtkList,
		scroll_type GtkScrollType,
		position Gfloat)

	Gtk_list_toggle_add_mode func(list *GtkList)

	Gtk_list_toggle_focus_row func(list *GtkList)

	Gtk_list_toggle_row func(list *GtkList,
		item *GtkWidget)

	Gtk_list_undo_selection func(list *GtkList)

	Gtk_list_end_drag_selection func(list *GtkList)

	Gtk_list_item_get_type func() GType

	Gtk_list_item_new func() *GtkWidget

	Gtk_list_item_new_with_label func(label string) *GtkWidget

	Gtk_list_item_select func(list_item *GtkListItem)

	Gtk_list_item_deselect func(list_item *GtkListItem)

	Gtk_old_editable_get_type func() GType

	Gtk_old_editable_claim_selection func(old_editable *GtkOldEditable,
		claim Gboolean,
		time_ Guint32)

	Gtk_old_editable_changed func(old_editable *GtkOldEditable)

	Gtk_option_menu_get_type func() GType

	Gtk_option_menu_new func() *GtkWidget

	Gtk_option_menu_get_menu func(option_menu *GtkOptionMenu) *GtkWidget

	Gtk_option_menu_set_menu func(option_menu *GtkOptionMenu,
		menu *GtkWidget)

	Gtk_option_menu_remove_menu func(option_menu *GtkOptionMenu)

	Gtk_option_menu_get_history func(option_menu *GtkOptionMenu) Gint

	Gtk_option_menu_set_history func(option_menu *GtkOptionMenu,
		index_ Guint)

	Gtk_preview_get_type func() GType

	Gtk_preview_uninit func()

	Gtk_preview_new func(t GtkPreviewType) *GtkWidget

	Gtk_preview_size func(preview *GtkPreview,
		width Gint,
		height Gint)

	Gtk_preview_put func(preview *GtkPreview,
		window *GdkWindow,
		gc *GdkGC,
		srcx Gint,
		srcy Gint,
		destx Gint,
		desty Gint,
		width Gint,
		height Gint)

	Gtk_preview_draw_row func(preview *GtkPreview,
		data *Guchar,
		x Gint,
		y Gint,
		w Gint)

	Gtk_preview_set_expand func(preview *GtkPreview,
		expand Gboolean)

	Gtk_preview_set_gamma func(gamma_ Double)

	Gtk_preview_set_color_cube func(nred_shades Guint,
		ngreen_shades Guint,
		nblue_shades Guint,
		ngray_shades Guint)

	Gtk_preview_set_install_cmap func(install_cmap Gint)

	Gtk_preview_set_reserved func(nreserved Gint)

	Gtk_preview_set_dither func(preview *GtkPreview,
		dither GdkRgbDither)

	Gtk_preview_get_visual func() *GdkVisual

	Gtk_preview_get_cmap func() *GdkColormap

	Gtk_preview_get_info func() *GtkPreviewInfo

	Gtk_preview_reset func()

	Gtk_tips_query_get_type func() GType

	Gtk_tips_query_new func() *GtkWidget

	Gtk_tips_query_start_query func(tips_query *GtkTipsQuery)

	Gtk_tips_query_stop_query func(tips_query *GtkTipsQuery)

	Gtk_tips_query_set_caller func(tips_query *GtkTipsQuery,
		caller *GtkWidget)

	Gtk_tips_query_set_labels func(tips_query *GtkTipsQuery,
		label_inactive string,
		label_no_tip string)

	Gtk_marshal_BOOLEAN__VOID func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_BOOLEAN__POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_BOOLEAN__POINTER_POINTER_INT_INT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_BOOLEAN__POINTER_INT_INT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_BOOLEAN__POINTER_INT_INT_UINT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_BOOLEAN__POINTER_STRING_STRING_POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_ENUM__ENUM func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_INT__POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_INT__POINTER_CHAR_CHAR func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__ENUM_FLOAT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__ENUM_FLOAT_BOOLEAN func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__INT_INT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__INT_INT_POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_INT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_POINTER_POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_STRING_STRING func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_UINT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_UINT_ENUM func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_POINTER_UINT_UINT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_INT_INT_POINTER_UINT_UINT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__POINTER_UINT_UINT func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__STRING_INT_POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__UINT_POINTER_UINT_ENUM_ENUM_POINTER func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__UINT_POINTER_UINT_UINT_ENUM func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_marshal_VOID__UINT_STRING func(
		closure *GClosure,
		return_value *GValue,
		n_param_values Guint,
		param_values *GValue,
		invocation_hint Gpointer,
		marshal_data Gpointer)

	Gtk_builder_error_quark func() GQuark

	Gtk_print_context_set_cairo_context func(
		context *GtkPrintContext,
		cr *Cairo_t,
		dpi_x, dpi_y Double)

	Gtk_tearoff_menu_item_get_type func() GType

	Gtk_tearoff_menu_item_new func() *GtkWidget

	Gtk_tree_get_type func() GType

	Gtk_tree_new func() *GtkWidget

	Gtk_tree_append func(tree *GtkTree, tree_item *GtkWidget)

	Gtk_tree_prepend func(tree *GtkTree, tree_item *GtkWidget)

	Gtk_tree_insert func(
		tree *GtkTree, tree_item *GtkWidget, position Gint)

	Gtk_tree_remove_items func(tree *GtkTree, items *GList)

	Gtk_tree_clear_items func(
		tree *GtkTree, start Gint, end Gint)

	Gtk_tree_select_item func(tree *GtkTree, item Gint)

	Gtk_tree_unselect_item func(tree *GtkTree, item Gint)

	Gtk_tree_select_child func(
		tree *GtkTree, tree_item *GtkWidget)

	Gtk_tree_unselect_child func(
		tree *GtkTree, tree_item *GtkWidget)

	Gtk_tree_child_position func(
		tree *GtkTree, child *GtkWidget) Gint

	Gtk_tree_set_selection_mode func(
		tree *GtkTree, mode GtkSelectionMode)

	Gtk_tree_set_view_mode func(
		tree *GtkTree, mode GtkTreeViewMode)

	Gtk_tree_set_view_lines func(tree *GtkTree, flag Gboolean)

	Gtk_tree_remove_item func(tree *GtkTree, child *GtkWidget)

	Gtk_tree_item_get_type func() GType

	Gtk_tree_item_new func() *GtkWidget

	Gtk_tree_item_new_with_label func(label string) *GtkWidget

	Gtk_tree_item_set_subtree func(
		tree_item *GtkTreeItem, subtree *GtkWidget)

	Gtk_tree_item_remove_subtree func(tree_item *GtkTreeItem)

	Gtk_tree_item_select func(tree_item *GtkTreeItem)

	Gtk_tree_item_deselect func(tree_item *GtkTreeItem)

	Gtk_tree_item_expand func(tree_item *GtkTreeItem)

	Gtk_tree_item_collapse func(tree_item *GtkTreeItem)

	Gtk_type_enum_get_values func(
		enum_type GtkType) *GtkEnumValue

	Gtk_type_flags_get_values func(
		flags_type GtkType) *GtkFlagValue

	Gtk_type_enum_find_value func(
		enum_type GtkType,
		value_name string) *GtkEnumValue

	Gtk_type_flags_find_value func(
		flags_type GtkType,
		value_name string) *GtkFlagValue
)
