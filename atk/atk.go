// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package atk provides API definitions for accessing
//libatk-1.0-0.dll.
package atk

import (
	"github.com/tHinqa/outside"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type (
	Enum int
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllData(dll, false, dataList)
}

type (
	AttributeSet T.GSList

	EventListenerInit func()

	FocusHandler func(*Object, bool)

	PropertyValues struct {
		PropertyName *T.Gchar
		OldValue     O.Value
		NewValue     O.Value
	}

	PropertyChangeHandler func(*Object, *PropertyValues)
)

type Layer Enum

const (
	LAYER_INVALID Layer = iota
	LAYER_BACKGROUND
	LAYER_CANVAS
	LAYER_WIDGET
	LAYER_MDI
	LAYER_POPUP
	LAYER_OVERLAY
	LAYER_WINDOW
)

var (
	UtilGetType func() O.Type

	FocusTrackerInit func(init EventListenerInit)

	FocusTrackerNotify func(object *Object)

	//atk_get_version func() int //NOTE: Not exposed

	NoOpObjectNew func(obj *O.Object) *Object

	NoOpObjectFactoryGetType func() O.Type

	NoOpObjectFactoryNew func() *ObjectFactory

	LayerGetType func() O.Type

	NoOpObjectGetType func() O.Type
)

//NOTE(t): Unreferenced
type KeyEventType Enum

const (
	KEY_EVENT_PRESS KeyEventType = iota
	KEY_EVENT_RELEASE
	KEY_EVENT_LAST_DEFINED
)

var KeyEventTypeGetType func() O.Type

type Misc struct{ parent O.Object }

var (
	MiscGetType func() O.Type

	MiscGetInstance func() *Misc

	miscThreadsEnter func(m *Misc)
	miscThreadsLeave func(m *Misc)
)

func (m *Misc) ThreadsEnter() { miscThreadsEnter(m) }
func (m *Misc) ThreadsLeave() { miscThreadsLeave(m) }

/*NOTE(t): atl_plug_... Not exposed
type Plug struct {
	Parent Object
}

var (
	PlugGetType func() O.Type
	PlugNew     func() *Object

	plugGetId func(p *Plug) string
)

func (p *Plug) GetId() string {	return plugGetId(p)}
*/

var dll = "libatk-1.0-0.dll"

var apiList = outside.Apis{
	{"atk_action_do_action", &actionDoAction},
	{"atk_action_get_description", &actionGetDescription},
	{"atk_action_get_keybinding", &actionGetKeybinding},
	{"atk_action_get_localized_name", &actionGetLocalizedName},
	{"atk_action_get_n_actions", &actionGetNActions},
	{"atk_action_get_name", &actionGetName},
	{"atk_action_get_type", &ActionGetType},
	{"atk_action_set_description", &actionSetDescription},
	{"atk_add_focus_tracker", &AddFocusTracker},
	{"atk_add_global_event_listener", &AddGlobalEventListener},
	{"atk_add_key_event_listener", &AddKeyEventListener},
	{"atk_attribute_set_free", &AttributeSetFree},
	{"atk_component_add_focus_handler", &componentAddFocusHandler},
	{"atk_component_contains", &componentContains},
	{"atk_component_get_alpha", &componentGetAlpha},
	{"atk_component_get_extents", &componentGetExtents},
	{"atk_component_get_layer", &componentGetLayer},
	{"atk_component_get_mdi_zorder", &componentGetMdiZorder},
	{"atk_component_get_position", &componentGetPosition},
	{"atk_component_get_size", &componentGetSize},
	{"atk_component_get_type", &ComponentGetType},
	{"atk_component_grab_focus", &componentGrabFocus},
	{"atk_component_ref_accessible_at_point", &componentRefAccessibleAtPoint},
	{"atk_component_remove_focus_handler", &componentRemoveFocusHandler},
	{"atk_component_set_extents", &componentSetExtents},
	{"atk_component_set_position", &componentSetPosition},
	{"atk_component_set_size", &componentSetSize},
	{"atk_coord_type_get_type", &CoordTypeGetType},
	{"atk_document_get_attribute_value", &documentGetAttributeValue},
	{"atk_document_get_attributes", &documentGetAttributes},
	{"atk_document_get_document", &documentGetDocument},
	{"atk_document_get_document_type", &documentGetDocumentType},
	{"atk_document_get_locale", &documentGetLocale},
	{"atk_document_get_type", &DocumentGetType},
	{"atk_document_set_attribute_value", &documentSetAttributeValue},
	{"atk_editable_text_copy_text", &editableTextCopyText},
	{"atk_editable_text_cut_text", &editableTextCutText},
	{"atk_editable_text_delete_text", &editableTextDeleteText},
	{"atk_editable_text_get_type", &EditableTextGetType},
	{"atk_editable_text_insert_text", &editableTextInsertText},
	{"atk_editable_text_paste_text", &editableTextPasteText},
	{"atk_editable_text_set_run_attributes", &editableTextSetRunAttributes},
	{"atk_editable_text_set_text_contents", &editableTextSetTextContents},
	{"atk_focus_tracker_init", &FocusTrackerInit},
	{"atk_focus_tracker_notify", &FocusTrackerNotify},
	{"atk_get_default_registry", &GetDefaultRegistry},
	{"atk_get_focus_object", &GetFocusObject},
	{"atk_get_root", &GetRoot},
	{"atk_get_toolkit_name", &GetToolkitName},
	{"atk_get_toolkit_version", &GetToolkitVersion},
	{"atk_gobject_accessible_for_object", &GobjectAccessibleForObject},
	{"atk_gobject_accessible_get_object", &gobjectAccessibleGetObject},
	{"atk_gobject_accessible_get_type", &GobjectAccessibleGetType},
	{"atk_hyperlink_get_end_index", &hyperlinkGetEndIndex},
	{"atk_hyperlink_get_n_anchors", &hyperlinkGetNAnchors},
	{"atk_hyperlink_get_object", &hyperlinkGetObject},
	{"atk_hyperlink_get_start_index", &hyperlinkGetStartIndex},
	{"atk_hyperlink_get_type", &HyperlinkGetType},
	{"atk_hyperlink_get_uri", &hyperlinkGetUri},
	{"atk_hyperlink_impl_get_hyperlink", &hyperlinkImplGetHyperlink},
	{"atk_hyperlink_impl_get_type", &HyperlinkImplGetType},
	{"atk_hyperlink_is_inline", &hyperlinkIsInline},
	{"atk_hyperlink_is_selected_link", &hyperlinkIsSelectedLink},
	{"atk_hyperlink_is_valid", &hyperlinkIsValid},
	{"atk_hyperlink_state_flags_get_type", &HyperlinkStateFlagsGetType},
	{"atk_hypertext_get_link", &hypertextGetLink},
	{"atk_hypertext_get_link_index", &hypertextGetLinkIndex},
	{"atk_hypertext_get_n_links", &hypertextGetNLinks},
	{"atk_hypertext_get_type", &HypertextGetType},
	{"atk_image_get_image_description", &imageGetImageDescription},
	{"atk_image_get_image_locale", &imageGetImageLocale},
	{"atk_image_get_image_position", &imageGetImagePosition},
	{"atk_image_get_image_size", &imageGetImageSize},
	{"atk_image_get_type", &ImageGetType},
	{"atk_image_set_image_description", &imageSetImageDescription},
	{"atk_implementor_get_type", &ImplementorGetType},
	{"atk_implementor_ref_accessible", &implementorRefAccessible},
	{"atk_key_event_type_get_type", &KeyEventTypeGetType},
	{"atk_layer_get_type", &LayerGetType},
	{"atk_misc_get_instance", &MiscGetInstance},
	{"atk_misc_get_type", &MiscGetType},
	{"atk_misc_threads_enter", &miscThreadsEnter},
	{"atk_misc_threads_leave", &miscThreadsLeave},
	{"atk_no_op_object_factory_get_type", &NoOpObjectFactoryGetType},
	{"atk_no_op_object_factory_new", &NoOpObjectFactoryNew},
	{"atk_no_op_object_get_type", &NoOpObjectGetType},
	{"atk_no_op_object_new", &NoOpObjectNew},
	{"atk_object_add_relationship", &objectAddRelationship},
	{"atk_object_connect_property_change_handler", &objectConnectPropertyChangeHandler},
	{"atk_object_factory_create_accessible", &objectFactoryCreateAccessible},
	{"atk_object_factory_get_accessible_type", &objectFactoryGetAccessibleType},
	{"atk_object_factory_get_type", &ObjectFactoryGetType},
	{"atk_object_factory_invalidate", &objectFactoryInvalidate},
	{"atk_object_get_attributes", &objectGetAttributes},
	{"atk_object_get_description", &objectGetDescription},
	{"atk_object_get_index_in_parent", &objectGetIndexInParent},
	{"atk_object_get_layer", &objectGetLayer},
	{"atk_object_get_mdi_zorder", &objectGetMdiZorder},
	{"atk_object_get_n_accessible_children", &objectGetNAccessibleChildren},
	{"atk_object_get_name", &objectGetName},
	{"atk_object_get_parent", &objectGetParent},
	{"atk_object_get_role", &objectGetRole},
	{"atk_object_get_type", &ObjectGetType},
	{"atk_object_initialize", &objectInitialize},
	{"atk_object_notify_state_change", &objectNotifyStateChange},
	{"atk_object_ref_accessible_child", &objectRefAccessibleChild},
	{"atk_object_ref_relation_set", &objectRefRelationSet},
	{"atk_object_ref_state_set", &objectRefStateSet},
	{"atk_object_remove_property_change_handler", &objectRemovePropertyChangeHandler},
	{"atk_object_remove_relationship", &objectRemoveRelationship},
	{"atk_object_set_description", &objectSetDescription},
	{"atk_object_set_name", &objectSetName},
	{"atk_object_set_parent", &objectSetParent},
	{"atk_object_set_role", &objectSetRole},
	{"atk_rectangle_get_type", &RectangleGetType},
	{"atk_registry_get_factory", &registryGetFactory},
	{"atk_registry_get_factory_type", &registryGetFactoryType},
	{"atk_registry_get_type", &RegistryGetType},
	{"atk_registry_set_factory_type", &registrySetFactoryType},
	{"atk_relation_add_target", &relationAddTarget},
	{"atk_relation_get_relation_type", &relationGetRelationType},
	{"atk_relation_get_target", &relationGetTarget},
	{"atk_relation_get_type", &RelationGetType},
	{"atk_relation_new", &RelationNew},
	{"atk_relation_set_add", &relationSetAdd},
	{"atk_relation_set_add_relation_by_type", &relationSetAddRelationByType},
	{"atk_relation_set_contains", &relationSetContains},
	{"atk_relation_set_get_n_relations", &relationSetGetNRelations},
	{"atk_relation_set_get_relation", &relationSetGetRelation},
	{"atk_relation_set_get_relation_by_type", &relationSetGetRelationByType},
	{"atk_relation_set_get_type", &RelationSetGetType},
	{"atk_relation_set_new", &RelationSetNew},
	{"atk_relation_set_remove", &relationSetRemove},
	{"atk_relation_type_for_name", &RelationTypeForName},
	{"atk_relation_type_get_name", &relationTypeGetName},
	{"atk_relation_type_get_type", &RelationTypeGetType},
	{"atk_relation_type_register", &RelationTypeRegister},
	{"atk_remove_focus_tracker", &RemoveFocusTracker},
	{"atk_remove_global_event_listener", &RemoveGlobalEventListener},
	{"atk_remove_key_event_listener", &RemoveKeyEventListener},
	{"atk_role_for_name", &RoleForName},
	{"atk_role_get_localized_name", &roleGetLocalizedName},
	{"atk_role_get_name", &roleGetName},
	{"atk_role_get_type", &RoleGetType},
	{"atk_role_register", &RoleRegister},
	{"atk_selection_add_selection", &selectionAddSelection},
	{"atk_selection_clear_selection", &selectionClearSelection},
	{"atk_selection_get_selection_count", &selectionGetSelectionCount},
	{"atk_selection_get_type", &SelectionGetType},
	{"atk_selection_is_child_selected", &selectionIsChildSelected},
	{"atk_selection_ref_selection", &selectionRefSelection},
	{"atk_selection_remove_selection", &selectionRemoveSelection},
	{"atk_selection_select_all_selection", &selectionSelectAllSelection},
	{"atk_state_set_add_state", &stateSetAddState},
	{"atk_state_set_add_states", &stateSetAddStates},
	{"atk_state_set_and_sets", &stateSetAndSets},
	{"atk_state_set_clear_states", &stateSetClearStates},
	{"atk_state_set_contains_state", &stateSetContainsState},
	{"atk_state_set_contains_states", &stateSetContainsStates},
	{"atk_state_set_get_type", &StateSetGetType},
	{"atk_state_set_is_empty", &stateSetIsEmpty},
	{"atk_state_set_new", &StateSetNew},
	{"atk_state_set_or_sets", &stateSetOrSets},
	{"atk_state_set_remove_state", &stateSetRemoveState},
	{"atk_state_set_xor_sets", &stateSetXorSets},
	{"atk_state_type_for_name", &StateTypeForName},
	{"atk_state_type_get_name", &stateTypeGetName},
	{"atk_state_type_get_type", &StateTypeGetType},
	{"atk_state_type_register", &StateTypeRegister},
	{"atk_streamable_content_get_mime_type", &streamableContentGetMimeType},
	{"atk_streamable_content_get_n_mime_types", &streamableContentGetNMimeTypes},
	{"atk_streamable_content_get_stream", &streamableContentGetStream},
	{"atk_streamable_content_get_type", &StreamableContentGetType},
	{"atk_streamable_content_get_uri", &streamableContentGetUri},
	{"atk_table_add_column_selection", &tableAddColumnSelection},
	{"atk_table_add_row_selection", &tableAddRowSelection},
	{"atk_table_get_caption", &tableGetCaption},
	{"atk_table_get_column_at_index", &tableGetColumnAtIndex},
	{"atk_table_get_column_description", &tableGetColumnDescription},
	{"atk_table_get_column_extent_at", &tableGetColumnExtentAt},
	{"atk_table_get_column_header", &tableGetColumnHeader},
	{"atk_table_get_index_at", &tableGetIndexAt},
	{"atk_table_get_n_columns", &tableGetNColumns},
	{"atk_table_get_n_rows", &tableGetNRows},
	{"atk_table_get_row_at_index", &tableGetRowAtIndex},
	{"atk_table_get_row_description", &tableGetRowDescription},
	{"atk_table_get_row_extent_at", &tableGetRowExtentAt},
	{"atk_table_get_row_header", &tableGetRowHeader},
	{"atk_table_get_selected_columns", &tableGetSelectedColumns},
	{"atk_table_get_selected_rows", &tableGetSelectedRows},
	{"atk_table_get_summary", &tableGetSummary},
	{"atk_table_get_type", &TableGetType},
	{"atk_table_is_column_selected", &tableIsColumnSelected},
	{"atk_table_is_row_selected", &tableIsRowSelected},
	{"atk_table_is_selected", &tableIsSelected},
	{"atk_table_ref_at", &tableRefAt},
	{"atk_table_remove_column_selection", &tableRemoveColumnSelection},
	{"atk_table_remove_row_selection", &tableRemoveRowSelection},
	{"atk_table_set_caption", &tableSetCaption},
	{"atk_table_set_column_description", &tableSetColumnDescription},
	{"atk_table_set_column_header", &tableSetColumnHeader},
	{"atk_table_set_row_description", &tableSetRowDescription},
	{"atk_table_set_row_header", &tableSetRowHeader},
	{"atk_table_set_summary", &tableSetSummary},
	{"atk_text_add_selection", &textAddSelection},
	{"atk_text_attribute_for_name", &TextAttributeForName},
	{"atk_text_attribute_get_name", &textAttributeGetName},
	{"atk_text_attribute_get_type", &TextAttributeGetType},
	{"atk_text_attribute_get_value", &textAttributeGetValue},
	{"atk_text_attribute_register", &TextAttributeRegister},
	{"atk_text_boundary_get_type", &TextBoundaryGetType},
	{"atk_text_clip_type_get_type", &TextClipTypeGetType},
	{"atk_text_free_ranges", &TextFreeRanges},
	{"atk_text_get_bounded_ranges", &textGetBoundedRanges},
	{"atk_text_get_caret_offset", &textGetCaretOffset},
	{"atk_text_get_character_at_offset", &textGetCharacterAtOffset},
	{"atk_text_get_character_count", &textGetCharacterCount},
	{"atk_text_get_character_extents", &textGetCharacterExtents},
	{"atk_text_get_default_attributes", &textGetDefaultAttributes},
	{"atk_text_get_n_selections", &textGetNSelections},
	{"atk_text_get_offset_at_point", &textGetOffsetAtPoint},
	{"atk_text_get_range_extents", &textGetRangeExtents},
	{"atk_text_get_run_attributes", &textGetRunAttributes},
	{"atk_text_get_selection", &textGetSelection},
	{"atk_text_get_text", &textGetText},
	{"atk_text_get_text_after_offset", &textGetTextAfterOffset},
	{"atk_text_get_text_at_offset", &textGetTextAtOffset},
	{"atk_text_get_text_before_offset", &textGetTextBeforeOffset},
	{"atk_text_get_type", &TextGetType},
	{"atk_text_remove_selection", &textRemoveSelection},
	{"atk_text_set_caret_offset", &textSetCaretOffset},
	{"atk_text_set_selection", &textSetSelection},
	{"atk_util_get_type", &UtilGetType},
	{"atk_value_get_current_value", &valueGetCurrentValue},
	{"atk_value_get_maximum_value", &valueGetMaximumValue},
	{"atk_value_get_minimum_increment", &valueGetMinimumIncrement},
	{"atk_value_get_minimum_value", &valueGetMinimumValue},
	{"atk_value_get_type", &ValueGetType},
	{"atk_value_set_current_value", &valueSetCurrentValue},
}

var dataList = outside.Data{
	{"atk_misc_instance", (*Misc)(nil)},
}
