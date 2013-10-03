package atk

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	Atk_state_type_register func(
		name *Gchar) AtkStateType

	Atk_state_type_get_name func(
		t AtkStateType) *Gchar

	Atk_state_type_for_name func(
		name *Gchar) AtkStateType

	Atk_role_register func(
		name *Gchar) AtkRole

	Atk_object_get_type func() GType

	Atk_implementor_get_type func() GType

	Atk_implementor_ref_accessible func(
		implementor *AtkImplementor) *AtkObject

	Atk_object_get_name func(
		accessible *AtkObject) *Gchar

	Atk_object_get_description func(
		accessible *AtkObject) *Gchar

	Atk_object_get_parent func(
		accessible *AtkObject) *AtkObject

	Atk_object_get_n_accessible_children func(
		accessible *AtkObject) Gint

	Atk_object_ref_accessible_child func(
		accessible *AtkObject,
		i Gint) *AtkObject

	Atk_object_ref_relation_set func(
		accessible *AtkObject) *AtkRelationSet

	Atk_object_get_role func(
		accessible *AtkObject) AtkRole

	Atk_object_get_layer func(
		accessible *AtkObject) AtkLayer

	Atk_object_get_mdi_zorder func(
		accessible *AtkObject) Gint

	Atk_object_get_attributes func(
		accessible *AtkObject) *AtkAttributeSet

	Atk_object_ref_state_set func(
		accessible *AtkObject) *AtkStateSet

	Atk_object_get_index_in_parent func(
		accessible *AtkObject) Gint

	Atk_object_set_name func(
		accessible *AtkObject,
		name *Gchar)

	Atk_object_set_description func(
		accessible *AtkObject,
		description *Gchar)

	Atk_object_set_parent func(
		accessible *AtkObject,
		parent *AtkObject)

	Atk_object_set_role func(
		accessible *AtkObject,
		role AtkRole)

	Atk_object_connect_property_change_handler func(
		accessible *AtkObject,
		handler *AtkPropertyChangeHandler) Guint

	Atk_object_remove_property_change_handler func(
		accessible *AtkObject,
		handler_id Guint)

	Atk_object_notify_state_change func(
		accessible *AtkObject,
		state AtkState,
		value Gboolean)

	Atk_object_initialize func(
		accessible *AtkObject,
		data Gpointer)

	Atk_role_get_name func(
		role AtkRole) *Gchar

	Atk_role_for_name func(
		name *Gchar) AtkRole

	Atk_object_add_relationship func(
		object *AtkObject,
		relationship AtkRelationType,
		target *AtkObject) Gboolean

	Atk_object_remove_relationship func(
		object *AtkObject,
		relationship AtkRelationType,
		target *AtkObject) Gboolean

	Atk_role_get_localized_name func(
		role AtkRole) *Gchar

	Atk_action_get_type func() GType

	Atk_action_do_action func(
		action *AtkAction,
		i Gint) Gboolean

	Atk_action_get_n_actions func(
		action *AtkAction) Gint

	Atk_action_get_description func(
		action *AtkAction,
		i Gint) *Gchar

	Atk_action_get_name func(
		action *AtkAction,
		i Gint) *Gchar

	Atk_action_get_keybinding func(
		action *AtkAction,
		i Gint) *Gchar

	Atk_action_set_description func(
		action *AtkAction,
		i Gint,
		desc *Gchar) Gboolean

	Atk_action_get_localized_name func(
		action *AtkAction,
		i Gint) *Gchar

	Atk_util_get_type func() GType

	Atk_add_focus_tracker func(
		focus_tracker AtkEventListener) Guint

	Atk_remove_focus_tracker func(
		tracker_id Guint)

	Atk_focus_tracker_init func(
		init AtkEventListenerInit)

	Atk_focus_tracker_notify func(
		object *AtkObject)

	Atk_add_global_event_listener func(
		listener GSignalEmissionHook,
		event_type *Gchar) Guint

	Atk_remove_global_event_listener func(
		listener_id Guint)

	Atk_add_key_event_listener func(
		listener AtkKeySnoopFunc,
		data Gpointer) Guint

	Atk_remove_key_event_listener func(
		listener_id Guint)

	Atk_get_root func() *AtkObject

	Atk_get_focus_object func() *AtkObject

	Atk_get_toolkit_name func() *Gchar

	Atk_get_toolkit_version func() *Gchar

	Atk_get_version func() *Gchar

	Atk_rectangle_get_type func() GType

	Atk_component_get_type func() GType

	Atk_component_add_focus_handler func(
		component *AtkComponent,
		handler AtkFocusHandler) Guint

	Atk_component_contains func(
		component *AtkComponent,
		x Gint,
		y Gint,
		coord_type AtkCoordType) Gboolean

	Atk_component_ref_accessible_at_point func(
		component *AtkComponent,
		x Gint,
		y Gint,
		coord_type AtkCoordType) *AtkObject

	Atk_component_get_extents func(
		component *AtkComponent,
		x *Gint,
		y *Gint,
		width *Gint,
		height *Gint,
		coord_type AtkCoordType)

	Atk_component_get_position func(
		component *AtkComponent,
		x *Gint,
		y *Gint,
		coord_type AtkCoordType)

	Atk_component_get_size func(
		component *AtkComponent,
		width *Gint,
		height *Gint)

	Atk_component_get_layer func(
		component *AtkComponent) AtkLayer

	Atk_component_get_mdi_zorder func(
		component *AtkComponent) Gint

	Atk_component_grab_focus func(
		component *AtkComponent) Gboolean

	Atk_component_remove_focus_handler func(
		component *AtkComponent,
		handler_id Guint)

	Atk_component_set_extents func(
		component *AtkComponent,
		x Gint,
		y Gint,
		width Gint,
		height Gint,
		coord_type AtkCoordType) Gboolean

	Atk_component_set_position func(
		component *AtkComponent,
		x Gint,
		y Gint,
		coord_type AtkCoordType) Gboolean

	Atk_component_set_size func(
		component *AtkComponent,
		width Gint,
		height Gint) Gboolean

	Atk_component_get_alpha func(
		component *AtkComponent) Gdouble

	Atk_document_get_type func() GType

	Atk_document_get_document_type func(
		document *AtkDocument) *Gchar

	Atk_document_get_document func(
		document *AtkDocument) Gpointer

	Atk_document_get_locale func(
		document *AtkDocument) *Gchar

	Atk_document_get_attributes func(
		document *AtkDocument) *AtkAttributeSet

	Atk_document_get_attribute_value func(
		document *AtkDocument,
		attribute_name *Gchar) *Gchar

	Atk_document_set_attribute_value func(
		document *AtkDocument,
		attribute_name *Gchar,
		attribute_value *Gchar) Gboolean

	Atk_text_attribute_register func(
		name *Gchar) AtkTextAttribute

	Atk_text_get_type func() GType

	Atk_text_get_text func(
		text *AtkText,
		start_offset Gint,
		end_offset Gint) *Gchar

	Atk_text_get_character_at_offset func(
		text *AtkText,
		offset Gint) Gunichar

	Atk_text_get_text_after_offset func(
		text *AtkText,
		offset Gint,
		boundary_type AtkTextBoundary,
		start_offset *Gint,
		end_offset *Gint) *Gchar

	Atk_text_get_text_at_offset func(
		text *AtkText,
		offset Gint,
		boundary_type AtkTextBoundary,
		start_offset *Gint,
		end_offset *Gint) *Gchar

	Atk_text_get_text_before_offset func(
		text *AtkText,
		offset Gint,
		boundary_type AtkTextBoundary,
		start_offset *Gint,
		end_offset *Gint) *Gchar

	Atk_text_get_caret_offset func(
		text *AtkText) Gint

	Atk_text_get_character_extents func(
		text *AtkText,
		offset Gint,
		x *Gint,
		y *Gint,
		width *Gint,
		height *Gint,
		coords AtkCoordType)

	Atk_text_get_run_attributes func(
		text *AtkText,
		offset Gint,
		start_offset *Gint,
		end_offset *Gint) *AtkAttributeSet

	Atk_text_get_default_attributes func(
		text *AtkText) *AtkAttributeSet

	Atk_text_get_character_count func(
		text *AtkText) Gint

	Atk_text_get_offset_at_point func(
		text *AtkText,
		x Gint,
		y Gint,
		coords AtkCoordType) Gint

	Atk_text_get_n_selections func(
		text *AtkText) Gint

	Atk_text_get_selection func(
		text *AtkText,
		selection_num Gint,
		start_offset *Gint,
		end_offset *Gint) *Gchar

	Atk_text_add_selection func(
		text *AtkText,
		start_offset Gint,
		end_offset Gint) Gboolean

	Atk_text_remove_selection func(
		text *AtkText,
		selection_num Gint) Gboolean

	Atk_text_set_selection func(
		text *AtkText,
		selection_num Gint,
		start_offset Gint,
		end_offset Gint) Gboolean

	Atk_text_set_caret_offset func(
		text *AtkText,
		offset Gint) Gboolean

	Atk_text_get_range_extents func(
		text *AtkText,
		start_offset Gint,
		end_offset Gint,
		coord_type AtkCoordType,
		rect *AtkTextRectangle)

	Atk_text_get_bounded_ranges func(
		text *AtkText,
		rect *AtkTextRectangle,
		coord_type AtkCoordType,
		x_clip_type AtkTextClipType,
		y_clip_type AtkTextClipType) **AtkTextRange

	Atk_text_free_ranges func(
		ranges **AtkTextRange)

	Atk_attribute_set_free func(
		attrib_set *AtkAttributeSet)

	Atk_text_attribute_get_name func(
		attr AtkTextAttribute) *Gchar

	Atk_text_attribute_for_name func(
		name *Gchar) AtkTextAttribute

	Atk_text_attribute_get_value func(
		attr AtkTextAttribute,
		index_ Gint) *Gchar

	Atk_editable_text_get_type func() GType

	Atk_editable_text_set_run_attributes func(
		text *AtkEditableText,
		attrib_set *AtkAttributeSet,
		start_offset Gint,
		end_offset Gint) Gboolean

	Atk_editable_text_set_text_contents func(
		text *AtkEditableText,
		string *Gchar)

	Atk_editable_text_insert_text func(
		text *AtkEditableText,
		string *Gchar,
		length Gint,
		position *Gint)

	Atk_editable_text_copy_text func(
		text *AtkEditableText,
		start_pos Gint,
		end_pos Gint)

	Atk_editable_text_cut_text func(
		text *AtkEditableText,
		start_pos Gint,
		end_pos Gint)

	Atk_editable_text_delete_text func(
		text *AtkEditableText,
		start_pos Gint,
		end_pos Gint)

	Atk_editable_text_paste_text func(
		text *AtkEditableText,
		position Gint)

	Atk_gobject_accessible_for_object func(
		obj *GObject) *AtkObject

	Atk_gobject_accessible_get_object func(
		obj *AtkGObjectAccessible) *GObject

	Atk_hyperlink_get_type func() GType

	Atk_hyperlink_get_uri func(
		link_ *AtkHyperlink,
		i Gint) *Gchar

	Atk_hyperlink_get_object func(
		link_ *AtkHyperlink,
		i Gint) *AtkObject

	Atk_hyperlink_get_end_index func(
		link_ *AtkHyperlink) Gint

	Atk_hyperlink_get_start_index func(
		link_ *AtkHyperlink) Gint

	Atk_hyperlink_is_valid func(
		link_ *AtkHyperlink) Gboolean

	Atk_hyperlink_is_inline func(
		link_ *AtkHyperlink) Gboolean

	Atk_hyperlink_get_n_anchors func(
		link_ *AtkHyperlink) Gint

	Atk_hyperlink_is_selected_link func(
		link_ *AtkHyperlink) Gboolean

	Atk_hyperlink_impl_get_type func() GType

	Atk_hyperlink_impl_get_hyperlink func(
		obj *AtkHyperlinkImpl) *AtkHyperlink

	Atk_hypertext_get_type func() GType

	Atk_hypertext_get_link func(
		hypertext *AtkHypertext,
		link_index Gint) *AtkHyperlink

	Atk_hypertext_get_n_links func(
		hypertext *AtkHypertext) Gint

	Atk_hypertext_get_link_index func(
		hypertext *AtkHypertext,
		char_index Gint) Gint

	Atk_image_get_type func() GType

	Atk_image_get_image_description func(
		image *AtkImage) *Gchar

	Atk_image_get_image_size func(
		image *AtkImage,
		width *Gint,
		height *Gint)

	Atk_image_set_image_description func(
		image *AtkImage,
		description *Gchar) Gboolean

	Atk_image_get_image_position func(
		image *AtkImage,
		x *Gint,
		y *Gint,
		coord_type AtkCoordType)

	Atk_image_get_image_locale func(
		image *AtkImage) *Gchar

	Atk_no_op_object_new func(
		obj *GObject) *AtkObject

	Atk_object_factory_get_type func() GType

	Atk_object_factory_create_accessible func(
		factory *AtkObjectFactory,
		obj *GObject) *AtkObject

	Atk_object_factory_invalidate func(
		factory *AtkObjectFactory)

	Atk_object_factory_get_accessible_type func(
		factory *AtkObjectFactory) GType

	Atk_no_op_object_factory_get_type func() GType

	Atk_no_op_object_factory_new func() *AtkObjectFactory

	Atk_plug_get_type func() GType

	Atk_plug_new func() *AtkObject

	Atk_plug_get_id func(plug *AtkPlug) *Gchar

	Atk_registry_get_type func() GType

	Atk_registry_set_factory_type func(
		registry *AtkRegistry,
		t GType,
		factory_type GType)

	Atk_registry_get_factory_type func(
		registry *AtkRegistry,
		t GType) GType

	Atk_registry_get_factory func(
		registry *AtkRegistry,
		t GType) *AtkObjectFactory

	Atk_get_default_registry func() *AtkRegistry

	Atk_relation_get_type func() GType

	Atk_relation_type_register func(
		name *Gchar) AtkRelationType

	Atk_relation_type_get_name func(
		t AtkRelationType) *Gchar

	Atk_relation_type_for_name func(
		name *Gchar) AtkRelationType

	Atk_relation_new func(
		targets **AtkObject,
		n_targets Gint,
		relationship AtkRelationType) *AtkRelation

	Atk_relation_get_relation_type func(
		relation *AtkRelation) AtkRelationType

	Atk_relation_get_target func(
		relation *AtkRelation) *GPtrArray

	Atk_relation_add_target func(
		relation *AtkRelation,
		target *AtkObject)

	Atk_relation_remove_target func(
		relation *AtkRelation,
		target *AtkObject) Gboolean

	Atk_relation_set_get_type func() GType

	Atk_relation_set_new func() *AtkRelationSet

	Atk_relation_set_contains func(
		set *AtkRelationSet,
		relationship AtkRelationType) Gboolean

	Atk_relation_set_remove func(
		set *AtkRelationSet,
		relation *AtkRelation)

	Atk_relation_set_add func(
		set *AtkRelationSet,
		relation *AtkRelation)

	Atk_relation_set_get_n_relations func(
		set *AtkRelationSet) Gint

	Atk_relation_set_get_relation func(
		set *AtkRelationSet,
		i Gint) *AtkRelation

	Atk_relation_set_get_relation_by_type func(
		set *AtkRelationSet,
		relationship AtkRelationType) *AtkRelation

	Atk_relation_set_add_relation_by_type func(
		set *AtkRelationSet,
		relationship AtkRelationType,
		target *AtkObject)

	Atk_selection_get_type func() GType

	Atk_selection_add_selection func(
		selection *AtkSelection,
		i Gint) Gboolean

	Atk_selection_clear_selection func(
		selection *AtkSelection) Gboolean

	Atk_selection_ref_selection func(
		selection *AtkSelection,
		i Gint) *AtkObject

	Atk_selection_get_selection_count func(
		selection *AtkSelection) Gint

	Atk_selection_is_child_selected func(
		selection *AtkSelection,
		i Gint) Gboolean

	Atk_selection_remove_selection func(
		selection *AtkSelection,
		i Gint) Gboolean

	Atk_selection_select_all_selection func(
		selection *AtkSelection) Gboolean

	Atk_socket_new func() *AtkObject

	Atk_socket_embed func(
		obj *AtkSocket,
		plug_id *Gchar)

	Atk_socket_is_occupied func(
		obj *AtkSocket) Gboolean

	Atk_state_set_get_type func() GType

	Atk_state_set_new func() *AtkStateSet

	Atk_state_set_is_empty func(
		set *AtkStateSet) Gboolean

	Atk_state_set_add_state func(
		set *AtkStateSet,
		t AtkStateType) Gboolean

	Atk_state_set_add_states func(
		set *AtkStateSet,
		types *AtkStateType,
		n_types Gint)

	Atk_state_set_clear_states func(
		set *AtkStateSet)

	Atk_state_set_contains_state func(
		set *AtkStateSet,
		t AtkStateType) Gboolean

	Atk_state_set_contains_states func(
		set *AtkStateSet,
		types *AtkStateType,
		n_types Gint) Gboolean

	Atk_state_set_remove_state func(
		set *AtkStateSet,
		t AtkStateType) Gboolean

	Atk_state_set_and_sets func(
		set *AtkStateSet,
		compare_set *AtkStateSet) *AtkStateSet

	Atk_state_set_or_sets func(
		set *AtkStateSet,
		compare_set *AtkStateSet) *AtkStateSet

	Atk_state_set_xor_sets func(
		set *AtkStateSet,
		compare_set *AtkStateSet) *AtkStateSet

	Atk_streamable_content_get_type func() GType

	Atk_streamable_content_get_n_mime_types func(
		streamable *AtkStreamableContent) Gint

	Atk_streamable_content_get_mime_type func(
		streamable *AtkStreamableContent,
		i Gint) *Gchar

	Atk_streamable_content_get_stream func(
		streamable *AtkStreamableContent,
		mime_type *Gchar) *GIOChannel

	Atk_streamable_content_get_uri func(
		streamable *AtkStreamableContent,
		mime_type *Gchar) *Gchar

	Atk_table_get_type func() GType

	Atk_table_ref_at func(
		table *AtkTable,
		row Gint,
		column Gint) *AtkObject

	Atk_table_get_index_at func(
		table *AtkTable,
		row Gint,
		column Gint) Gint

	Atk_table_get_column_at_index func(
		table *AtkTable,
		index_ Gint) Gint

	Atk_table_get_row_at_index func(
		table *AtkTable,
		index_ Gint) Gint

	Atk_table_get_n_columns func(
		table *AtkTable) Gint

	Atk_table_get_n_rows func(
		table *AtkTable) Gint

	Atk_table_get_column_extent_at func(
		table *AtkTable,
		row Gint,
		column Gint) Gint

	Atk_table_get_row_extent_at func(
		table *AtkTable,
		row Gint,
		column Gint) Gint

	Atk_table_get_caption func(
		table *AtkTable) *AtkObject

	Atk_table_get_column_description func(
		table *AtkTable,
		column Gint) *Gchar

	Atk_table_get_column_header func(
		table *AtkTable,
		column Gint) *AtkObject

	Atk_table_get_row_description func(
		table *AtkTable,
		row Gint) *Gchar

	Atk_table_get_row_header func(
		table *AtkTable,
		row Gint) *AtkObject

	Atk_table_get_summary func(
		table *AtkTable) *AtkObject

	Atk_table_set_caption func(
		table *AtkTable,
		caption *AtkObject)

	Atk_table_set_column_description func(
		table *AtkTable,
		column Gint,
		description *Gchar)

	Atk_table_set_column_header func(
		table *AtkTable,
		column Gint,
		header *AtkObject)

	Atk_table_set_row_description func(
		table *AtkTable,
		row Gint,
		description *Gchar)

	Atk_table_set_row_header func(
		table *AtkTable,
		row Gint,
		header *AtkObject)

	Atk_table_set_summary func(
		table *AtkTable,
		accessible *AtkObject)

	Atk_table_get_selected_columns func(
		table *AtkTable,
		selected **Gint) Gint

	Atk_table_get_selected_rows func(
		table *AtkTable,
		selected **Gint) Gint

	Atk_table_is_column_selected func(
		table *AtkTable,
		column Gint) Gboolean

	Atk_table_is_row_selected func(
		table *AtkTable,
		row Gint) Gboolean

	Atk_table_is_selected func(
		table *AtkTable,
		row Gint,
		column Gint) Gboolean

	Atk_table_add_row_selection func(
		table *AtkTable,
		row Gint) Gboolean

	Atk_table_remove_row_selection func(
		table *AtkTable,
		row Gint) Gboolean

	Atk_table_add_column_selection func(
		table *AtkTable,
		column Gint) Gboolean

	Atk_table_remove_column_selection func(
		table *AtkTable,
		column Gint) Gboolean

	Atk_misc_get_type func() GType

	Atk_misc_threads_enter func(
		misc *AtkMisc)

	Atk_misc_threads_leave func(
		misc *AtkMisc)

	Atk_misc_get_instance func() *AtkMisc

	Atk_value_get_type func() GType

	Atk_value_get_current_value func(
		obj *AtkValue, value *GValue)

	Atk_value_get_maximum_value func(
		obj *AtkValue, value *GValue)

	Atk_value_get_minimum_value func(
		obj *AtkValue, value *GValue)

	Atk_value_set_current_value func(
		obj *AtkValue, value *GValue) Gboolean

	Atk_value_get_minimum_increment func(
		obj *AtkValue, value *GValue)

	Atk_coord_type_get_type func() GType

	Atk_gobject_accessible_get_type func() GType

	Atk_hyperlink_state_flags_get_type func() GType

	Atk_key_event_type_get_type func() GType

	Atk_layer_get_type func() GType

	Atk_no_op_object_get_type func() GType

	Atk_relation_type_get_type func() GType

	Atk_role_get_type func() GType

	Atk_state_type_get_type func() GType

	Atk_text_attribute_get_type func() GType

	Atk_text_boundary_get_type func() GType

	Atk_text_clip_type_get_type func() GType
)

var dll = "libatk-1.0-0.dll"

var apiList = outside.Apis{
	{"atk_action_do_action", &Atk_action_do_action},
	{"atk_action_get_description", &Atk_action_get_description},
	{"atk_action_get_keybinding", &Atk_action_get_keybinding},
	{"atk_action_get_localized_name", &Atk_action_get_localized_name},
	{"atk_action_get_n_actions", &Atk_action_get_n_actions},
	{"atk_action_get_name", &Atk_action_get_name},
	{"atk_action_get_type", &Atk_action_get_type},
	{"atk_action_set_description", &Atk_action_set_description},
	{"atk_add_focus_tracker", &Atk_add_focus_tracker},
	{"atk_add_global_event_listener", &Atk_add_global_event_listener},
	{"atk_add_key_event_listener", &Atk_add_key_event_listener},
	{"atk_attribute_set_free", &Atk_attribute_set_free},
	{"atk_component_add_focus_handler", &Atk_component_add_focus_handler},
	{"atk_component_contains", &Atk_component_contains},
	{"atk_component_get_alpha", &Atk_component_get_alpha},
	{"atk_component_get_extents", &Atk_component_get_extents},
	{"atk_component_get_layer", &Atk_component_get_layer},
	{"atk_component_get_mdi_zorder", &Atk_component_get_mdi_zorder},
	{"atk_component_get_position", &Atk_component_get_position},
	{"atk_component_get_size", &Atk_component_get_size},
	{"atk_component_get_type", &Atk_component_get_type},
	{"atk_component_grab_focus", &Atk_component_grab_focus},
	{"atk_component_ref_accessible_at_point", &Atk_component_ref_accessible_at_point},
	{"atk_component_remove_focus_handler", &Atk_component_remove_focus_handler},
	{"atk_component_set_extents", &Atk_component_set_extents},
	{"atk_component_set_position", &Atk_component_set_position},
	{"atk_component_set_size", &Atk_component_set_size},
	{"atk_coord_type_get_type", &Atk_coord_type_get_type},
	{"atk_document_get_attribute_value", &Atk_document_get_attribute_value},
	{"atk_document_get_attributes", &Atk_document_get_attributes},
	{"atk_document_get_document", &Atk_document_get_document},
	{"atk_document_get_document_type", &Atk_document_get_document_type},
	{"atk_document_get_locale", &Atk_document_get_locale},
	{"atk_document_get_type", &Atk_document_get_type},
	{"atk_document_set_attribute_value", &Atk_document_set_attribute_value},
	{"atk_editable_text_copy_text", &Atk_editable_text_copy_text},
	{"atk_editable_text_cut_text", &Atk_editable_text_cut_text},
	{"atk_editable_text_delete_text", &Atk_editable_text_delete_text},
	{"atk_editable_text_get_type", &Atk_editable_text_get_type},
	{"atk_editable_text_insert_text", &Atk_editable_text_insert_text},
	{"atk_editable_text_paste_text", &Atk_editable_text_paste_text},
	{"atk_editable_text_set_run_attributes", &Atk_editable_text_set_run_attributes},
	{"atk_editable_text_set_text_contents", &Atk_editable_text_set_text_contents},
	{"atk_focus_tracker_init", &Atk_focus_tracker_init},
	{"atk_focus_tracker_notify", &Atk_focus_tracker_notify},
	{"atk_get_default_registry", &Atk_get_default_registry},
	{"atk_get_focus_object", &Atk_get_focus_object},
	{"atk_get_root", &Atk_get_root},
	{"atk_get_toolkit_name", &Atk_get_toolkit_name},
	{"atk_get_toolkit_version", &Atk_get_toolkit_version},
	{"atk_gobject_accessible_for_object", &Atk_gobject_accessible_for_object},
	{"atk_gobject_accessible_get_object", &Atk_gobject_accessible_get_object},
	{"atk_gobject_accessible_get_type", &Atk_gobject_accessible_get_type},
	{"atk_hyperlink_get_end_index", &Atk_hyperlink_get_end_index},
	{"atk_hyperlink_get_n_anchors", &Atk_hyperlink_get_n_anchors},
	{"atk_hyperlink_get_object", &Atk_hyperlink_get_object},
	{"atk_hyperlink_get_start_index", &Atk_hyperlink_get_start_index},
	{"atk_hyperlink_get_type", &Atk_hyperlink_get_type},
	{"atk_hyperlink_get_uri", &Atk_hyperlink_get_uri},
	{"atk_hyperlink_impl_get_hyperlink", &Atk_hyperlink_impl_get_hyperlink},
	{"atk_hyperlink_impl_get_type", &Atk_hyperlink_impl_get_type},
	{"atk_hyperlink_is_inline", &Atk_hyperlink_is_inline},
	{"atk_hyperlink_is_selected_link", &Atk_hyperlink_is_selected_link},
	{"atk_hyperlink_is_valid", &Atk_hyperlink_is_valid},
	{"atk_hyperlink_state_flags_get_type", &Atk_hyperlink_state_flags_get_type},
	{"atk_hypertext_get_link", &Atk_hypertext_get_link},
	{"atk_hypertext_get_link_index", &Atk_hypertext_get_link_index},
	{"atk_hypertext_get_n_links", &Atk_hypertext_get_n_links},
	{"atk_hypertext_get_type", &Atk_hypertext_get_type},
	{"atk_image_get_image_description", &Atk_image_get_image_description},
	{"atk_image_get_image_locale", &Atk_image_get_image_locale},
	{"atk_image_get_image_position", &Atk_image_get_image_position},
	{"atk_image_get_image_size", &Atk_image_get_image_size},
	{"atk_image_get_type", &Atk_image_get_type},
	{"atk_image_set_image_description", &Atk_image_set_image_description},
	{"atk_implementor_get_type", &Atk_implementor_get_type},
	{"atk_implementor_ref_accessible", &Atk_implementor_ref_accessible},
	{"atk_key_event_type_get_type", &Atk_key_event_type_get_type},
	{"atk_layer_get_type", &Atk_layer_get_type},
	{"atk_misc_get_instance", &Atk_misc_get_instance},
	{"atk_misc_get_type", &Atk_misc_get_type},
	// Data {"atk_misc_instance", &Atk_misc_instance},
	{"atk_misc_threads_enter", &Atk_misc_threads_enter},
	{"atk_misc_threads_leave", &Atk_misc_threads_leave},
	{"atk_no_op_object_factory_get_type", &Atk_no_op_object_factory_get_type},
	{"atk_no_op_object_factory_new", &Atk_no_op_object_factory_new},
	{"atk_no_op_object_get_type", &Atk_no_op_object_get_type},
	{"atk_no_op_object_new", &Atk_no_op_object_new},
	{"atk_object_add_relationship", &Atk_object_add_relationship},
	{"atk_object_connect_property_change_handler", &Atk_object_connect_property_change_handler},
	{"atk_object_factory_create_accessible", &Atk_object_factory_create_accessible},
	{"atk_object_factory_get_accessible_type", &Atk_object_factory_get_accessible_type},
	{"atk_object_factory_get_type", &Atk_object_factory_get_type},
	{"atk_object_factory_invalidate", &Atk_object_factory_invalidate},
	{"atk_object_get_attributes", &Atk_object_get_attributes},
	{"atk_object_get_description", &Atk_object_get_description},
	{"atk_object_get_index_in_parent", &Atk_object_get_index_in_parent},
	{"atk_object_get_layer", &Atk_object_get_layer},
	{"atk_object_get_mdi_zorder", &Atk_object_get_mdi_zorder},
	{"atk_object_get_n_accessible_children", &Atk_object_get_n_accessible_children},
	{"atk_object_get_name", &Atk_object_get_name},
	{"atk_object_get_parent", &Atk_object_get_parent},
	{"atk_object_get_role", &Atk_object_get_role},
	{"atk_object_get_type", &Atk_object_get_type},
	{"atk_object_initialize", &Atk_object_initialize},
	{"atk_object_notify_state_change", &Atk_object_notify_state_change},
	{"atk_object_ref_accessible_child", &Atk_object_ref_accessible_child},
	{"atk_object_ref_relation_set", &Atk_object_ref_relation_set},
	{"atk_object_ref_state_set", &Atk_object_ref_state_set},
	{"atk_object_remove_property_change_handler", &Atk_object_remove_property_change_handler},
	{"atk_object_remove_relationship", &Atk_object_remove_relationship},
	{"atk_object_set_description", &Atk_object_set_description},
	{"atk_object_set_name", &Atk_object_set_name},
	{"atk_object_set_parent", &Atk_object_set_parent},
	{"atk_object_set_role", &Atk_object_set_role},
	{"atk_rectangle_get_type", &Atk_rectangle_get_type},
	{"atk_registry_get_factory", &Atk_registry_get_factory},
	{"atk_registry_get_factory_type", &Atk_registry_get_factory_type},
	{"atk_registry_get_type", &Atk_registry_get_type},
	{"atk_registry_set_factory_type", &Atk_registry_set_factory_type},
	{"atk_relation_add_target", &Atk_relation_add_target},
	{"atk_relation_get_relation_type", &Atk_relation_get_relation_type},
	{"atk_relation_get_target", &Atk_relation_get_target},
	{"atk_relation_get_type", &Atk_relation_get_type},
	{"atk_relation_new", &Atk_relation_new},
	{"atk_relation_set_add", &Atk_relation_set_add},
	{"atk_relation_set_add_relation_by_type", &Atk_relation_set_add_relation_by_type},
	{"atk_relation_set_contains", &Atk_relation_set_contains},
	{"atk_relation_set_get_n_relations", &Atk_relation_set_get_n_relations},
	{"atk_relation_set_get_relation", &Atk_relation_set_get_relation},
	{"atk_relation_set_get_relation_by_type", &Atk_relation_set_get_relation_by_type},
	{"atk_relation_set_get_type", &Atk_relation_set_get_type},
	{"atk_relation_set_new", &Atk_relation_set_new},
	{"atk_relation_set_remove", &Atk_relation_set_remove},
	{"atk_relation_type_for_name", &Atk_relation_type_for_name},
	{"atk_relation_type_get_name", &Atk_relation_type_get_name},
	{"atk_relation_type_get_type", &Atk_relation_type_get_type},
	{"atk_relation_type_register", &Atk_relation_type_register},
	{"atk_remove_focus_tracker", &Atk_remove_focus_tracker},
	{"atk_remove_global_event_listener", &Atk_remove_global_event_listener},
	{"atk_remove_key_event_listener", &Atk_remove_key_event_listener},
	{"atk_role_for_name", &Atk_role_for_name},
	{"atk_role_get_localized_name", &Atk_role_get_localized_name},
	{"atk_role_get_name", &Atk_role_get_name},
	{"atk_role_get_type", &Atk_role_get_type},
	{"atk_role_register", &Atk_role_register},
	{"atk_selection_add_selection", &Atk_selection_add_selection},
	{"atk_selection_clear_selection", &Atk_selection_clear_selection},
	{"atk_selection_get_selection_count", &Atk_selection_get_selection_count},
	{"atk_selection_get_type", &Atk_selection_get_type},
	{"atk_selection_is_child_selected", &Atk_selection_is_child_selected},
	{"atk_selection_ref_selection", &Atk_selection_ref_selection},
	{"atk_selection_remove_selection", &Atk_selection_remove_selection},
	{"atk_selection_select_all_selection", &Atk_selection_select_all_selection},
	{"atk_state_set_add_state", &Atk_state_set_add_state},
	{"atk_state_set_add_states", &Atk_state_set_add_states},
	{"atk_state_set_and_sets", &Atk_state_set_and_sets},
	{"atk_state_set_clear_states", &Atk_state_set_clear_states},
	{"atk_state_set_contains_state", &Atk_state_set_contains_state},
	{"atk_state_set_contains_states", &Atk_state_set_contains_states},
	{"atk_state_set_get_type", &Atk_state_set_get_type},
	{"atk_state_set_is_empty", &Atk_state_set_is_empty},
	{"atk_state_set_new", &Atk_state_set_new},
	{"atk_state_set_or_sets", &Atk_state_set_or_sets},
	{"atk_state_set_remove_state", &Atk_state_set_remove_state},
	{"atk_state_set_xor_sets", &Atk_state_set_xor_sets},
	{"atk_state_type_for_name", &Atk_state_type_for_name},
	{"atk_state_type_get_name", &Atk_state_type_get_name},
	{"atk_state_type_get_type", &Atk_state_type_get_type},
	{"atk_state_type_register", &Atk_state_type_register},
	{"atk_streamable_content_get_mime_type", &Atk_streamable_content_get_mime_type},
	{"atk_streamable_content_get_n_mime_types", &Atk_streamable_content_get_n_mime_types},
	{"atk_streamable_content_get_stream", &Atk_streamable_content_get_stream},
	{"atk_streamable_content_get_type", &Atk_streamable_content_get_type},
	{"atk_streamable_content_get_uri", &Atk_streamable_content_get_uri},
	{"atk_table_add_column_selection", &Atk_table_add_column_selection},
	{"atk_table_add_row_selection", &Atk_table_add_row_selection},
	{"atk_table_get_caption", &Atk_table_get_caption},
	{"atk_table_get_column_at_index", &Atk_table_get_column_at_index},
	{"atk_table_get_column_description", &Atk_table_get_column_description},
	{"atk_table_get_column_extent_at", &Atk_table_get_column_extent_at},
	{"atk_table_get_column_header", &Atk_table_get_column_header},
	{"atk_table_get_index_at", &Atk_table_get_index_at},
	{"atk_table_get_n_columns", &Atk_table_get_n_columns},
	{"atk_table_get_n_rows", &Atk_table_get_n_rows},
	{"atk_table_get_row_at_index", &Atk_table_get_row_at_index},
	{"atk_table_get_row_description", &Atk_table_get_row_description},
	{"atk_table_get_row_extent_at", &Atk_table_get_row_extent_at},
	{"atk_table_get_row_header", &Atk_table_get_row_header},
	{"atk_table_get_selected_columns", &Atk_table_get_selected_columns},
	{"atk_table_get_selected_rows", &Atk_table_get_selected_rows},
	{"atk_table_get_summary", &Atk_table_get_summary},
	{"atk_table_get_type", &Atk_table_get_type},
	{"atk_table_is_column_selected", &Atk_table_is_column_selected},
	{"atk_table_is_row_selected", &Atk_table_is_row_selected},
	{"atk_table_is_selected", &Atk_table_is_selected},
	{"atk_table_ref_at", &Atk_table_ref_at},
	{"atk_table_remove_column_selection", &Atk_table_remove_column_selection},
	{"atk_table_remove_row_selection", &Atk_table_remove_row_selection},
	{"atk_table_set_caption", &Atk_table_set_caption},
	{"atk_table_set_column_description", &Atk_table_set_column_description},
	{"atk_table_set_column_header", &Atk_table_set_column_header},
	{"atk_table_set_row_description", &Atk_table_set_row_description},
	{"atk_table_set_row_header", &Atk_table_set_row_header},
	{"atk_table_set_summary", &Atk_table_set_summary},
	{"atk_text_add_selection", &Atk_text_add_selection},
	{"atk_text_attribute_for_name", &Atk_text_attribute_for_name},
	{"atk_text_attribute_get_name", &Atk_text_attribute_get_name},
	{"atk_text_attribute_get_type", &Atk_text_attribute_get_type},
	{"atk_text_attribute_get_value", &Atk_text_attribute_get_value},
	{"atk_text_attribute_register", &Atk_text_attribute_register},
	{"atk_text_boundary_get_type", &Atk_text_boundary_get_type},
	{"atk_text_clip_type_get_type", &Atk_text_clip_type_get_type},
	{"atk_text_free_ranges", &Atk_text_free_ranges},
	{"atk_text_get_bounded_ranges", &Atk_text_get_bounded_ranges},
	{"atk_text_get_caret_offset", &Atk_text_get_caret_offset},
	{"atk_text_get_character_at_offset", &Atk_text_get_character_at_offset},
	{"atk_text_get_character_count", &Atk_text_get_character_count},
	{"atk_text_get_character_extents", &Atk_text_get_character_extents},
	{"atk_text_get_default_attributes", &Atk_text_get_default_attributes},
	{"atk_text_get_n_selections", &Atk_text_get_n_selections},
	{"atk_text_get_offset_at_point", &Atk_text_get_offset_at_point},
	{"atk_text_get_range_extents", &Atk_text_get_range_extents},
	{"atk_text_get_run_attributes", &Atk_text_get_run_attributes},
	{"atk_text_get_selection", &Atk_text_get_selection},
	{"atk_text_get_text", &Atk_text_get_text},
	{"atk_text_get_text_after_offset", &Atk_text_get_text_after_offset},
	{"atk_text_get_text_at_offset", &Atk_text_get_text_at_offset},
	{"atk_text_get_text_before_offset", &Atk_text_get_text_before_offset},
	{"atk_text_get_type", &Atk_text_get_type},
	{"atk_text_remove_selection", &Atk_text_remove_selection},
	{"atk_text_set_caret_offset", &Atk_text_set_caret_offset},
	{"atk_text_set_selection", &Atk_text_set_selection},
	{"atk_util_get_type", &Atk_util_get_type},
	{"atk_value_get_current_value", &Atk_value_get_current_value},
	{"atk_value_get_maximum_value", &Atk_value_get_maximum_value},
	{"atk_value_get_minimum_increment", &Atk_value_get_minimum_increment},
	{"atk_value_get_minimum_value", &Atk_value_get_minimum_value},
	{"atk_value_get_type", &Atk_value_get_type},
	{"atk_value_set_current_value", &Atk_value_set_current_value},
}
