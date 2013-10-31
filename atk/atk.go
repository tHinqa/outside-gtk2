// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package atk provides API definitions for accessing
//libatk-1.0-0.dll.
package atk

import (
	"github.com/tHinqa/outside"
	L "github.com/tHinqa/outside-gtk2/glib"
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
	AttributeSet L.SList

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

	MiscThreadsEnter func(m *Misc)
	MiscThreadsLeave func(m *Misc)
)

func (m *Misc) ThreadsEnter() { MiscThreadsEnter(m) }
func (m *Misc) ThreadsLeave() { MiscThreadsLeave(m) }

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
	{"atk_action_do_action", &ActionDoAction},
	{"atk_action_get_description", &ActionGetDescription},
	{"atk_action_get_keybinding", &ActionGetKeybinding},
	{"atk_action_get_localized_name", &ActionGetLocalizedName},
	{"atk_action_get_n_actions", &ActionGetNActions},
	{"atk_action_get_name", &ActionGetName},
	{"atk_action_get_type", &ActionGetType},
	{"atk_action_set_description", &ActionSetDescription},
	{"atk_add_focus_tracker", &AddFocusTracker},
	{"atk_add_global_event_listener", &AddGlobalEventListener},
	{"atk_add_key_event_listener", &AddKeyEventListener},
	{"atk_attribute_set_free", &AttributeSetFree},
	{"atk_component_add_focus_handler", &ComponentAddFocusHandler},
	{"atk_component_contains", &ComponentContains},
	{"atk_component_get_alpha", &ComponentGetAlpha},
	{"atk_component_get_extents", &ComponentGetExtents},
	{"atk_component_get_layer", &ComponentGetLayer},
	{"atk_component_get_mdi_zorder", &ComponentGetMdiZorder},
	{"atk_component_get_position", &ComponentGetPosition},
	{"atk_component_get_size", &ComponentGetSize},
	{"atk_component_get_type", &ComponentGetType},
	{"atk_component_grab_focus", &ComponentGrabFocus},
	{"atk_component_ref_accessible_at_point", &ComponentRefAccessibleAtPoint},
	{"atk_component_remove_focus_handler", &ComponentRemoveFocusHandler},
	{"atk_component_set_extents", &ComponentSetExtents},
	{"atk_component_set_position", &ComponentSetPosition},
	{"atk_component_set_size", &ComponentSetSize},
	{"atk_coord_type_get_type", &CoordTypeGetType},
	{"atk_document_get_attribute_value", &DocumentGetAttributeValue},
	{"atk_document_get_attributes", &DocumentGetAttributes},
	{"atk_document_get_document", &DocumentGetDocument},
	{"atk_document_get_document_type", &DocumentGetDocumentType},
	{"atk_document_get_locale", &DocumentGetLocale},
	{"atk_document_get_type", &DocumentGetType},
	{"atk_document_set_attribute_value", &DocumentSetAttributeValue},
	{"atk_editable_text_copy_text", &EditableTextCopyText},
	{"atk_editable_text_cut_text", &EditableTextCutText},
	{"atk_editable_text_delete_text", &EditableTextDeleteText},
	{"atk_editable_text_get_type", &EditableTextGetType},
	{"atk_editable_text_insert_text", &EditableTextInsertText},
	{"atk_editable_text_paste_text", &EditableTextPasteText},
	{"atk_editable_text_set_run_attributes", &EditableTextSetRunAttributes},
	{"atk_editable_text_set_text_contents", &EditableTextSetTextContents},
	{"atk_focus_tracker_init", &FocusTrackerInit},
	{"atk_focus_tracker_notify", &FocusTrackerNotify},
	{"atk_get_default_registry", &GetDefaultRegistry},
	{"atk_get_focus_object", &GetFocusObject},
	{"atk_get_root", &GetRoot},
	{"atk_get_toolkit_name", &GetToolkitName},
	{"atk_get_toolkit_version", &GetToolkitVersion},
	{"atk_gobject_accessible_for_object", &GobjectAccessibleForObject},
	{"atk_gobject_accessible_get_object", &GobjectAccessibleGetObject},
	{"atk_gobject_accessible_get_type", &GobjectAccessibleGetType},
	{"atk_hyperlink_get_end_index", &HyperlinkGetEndIndex},
	{"atk_hyperlink_get_n_anchors", &HyperlinkGetNAnchors},
	{"atk_hyperlink_get_object", &HyperlinkGetObject},
	{"atk_hyperlink_get_start_index", &HyperlinkGetStartIndex},
	{"atk_hyperlink_get_type", &HyperlinkGetType},
	{"atk_hyperlink_get_uri", &HyperlinkGetUri},
	{"atk_hyperlink_impl_get_hyperlink", &HyperlinkImplGetHyperlink},
	{"atk_hyperlink_impl_get_type", &HyperlinkImplGetType},
	{"atk_hyperlink_is_inline", &HyperlinkIsInline},
	{"atk_hyperlink_is_selected_link", &HyperlinkIsSelectedLink},
	{"atk_hyperlink_is_valid", &HyperlinkIsValid},
	{"atk_hyperlink_state_flags_get_type", &HyperlinkStateFlagsGetType},
	{"atk_hypertext_get_link", &HypertextGetLink},
	{"atk_hypertext_get_link_index", &HypertextGetLinkIndex},
	{"atk_hypertext_get_n_links", &HypertextGetNLinks},
	{"atk_hypertext_get_type", &HypertextGetType},
	{"atk_image_get_image_description", &ImageGetImageDescription},
	{"atk_image_get_image_locale", &ImageGetImageLocale},
	{"atk_image_get_image_position", &ImageGetImagePosition},
	{"atk_image_get_image_size", &ImageGetImageSize},
	{"atk_image_get_type", &ImageGetType},
	{"atk_image_set_image_description", &ImageSetImageDescription},
	{"atk_implementor_get_type", &ImplementorGetType},
	{"atk_implementor_ref_accessible", &ImplementorRefAccessible},
	{"atk_key_event_type_get_type", &KeyEventTypeGetType},
	{"atk_layer_get_type", &LayerGetType},
	{"atk_misc_get_instance", &MiscGetInstance},
	{"atk_misc_get_type", &MiscGetType},
	{"atk_misc_threads_enter", &MiscThreadsEnter},
	{"atk_misc_threads_leave", &MiscThreadsLeave},
	{"atk_no_op_object_factory_get_type", &NoOpObjectFactoryGetType},
	{"atk_no_op_object_factory_new", &NoOpObjectFactoryNew},
	{"atk_no_op_object_get_type", &NoOpObjectGetType},
	{"atk_no_op_object_new", &NoOpObjectNew},
	{"atk_object_add_relationship", &ObjectAddRelationship},
	{"atk_object_connect_property_change_handler", &ObjectConnectPropertyChangeHandler},
	{"atk_object_factory_create_accessible", &ObjectFactoryCreateAccessible},
	{"atk_object_factory_get_accessible_type", &ObjectFactoryGetAccessibleType},
	{"atk_object_factory_get_type", &ObjectFactoryGetType},
	{"atk_object_factory_invalidate", &ObjectFactoryInvalidate},
	{"atk_object_get_attributes", &ObjectGetAttributes},
	{"atk_object_get_description", &ObjectGetDescription},
	{"atk_object_get_index_in_parent", &ObjectGetIndexInParent},
	{"atk_object_get_layer", &ObjectGetLayer},
	{"atk_object_get_mdi_zorder", &ObjectGetMdiZorder},
	{"atk_object_get_n_accessible_children", &ObjectGetNAccessibleChildren},
	{"atk_object_get_name", &ObjectGetName},
	{"atk_object_get_parent", &ObjectGetParent},
	{"atk_object_get_role", &ObjectGetRole},
	{"atk_object_get_type", &ObjectGetType},
	{"atk_object_initialize", &ObjectInitialize},
	{"atk_object_notify_state_change", &ObjectNotifyStateChange},
	{"atk_object_ref_accessible_child", &ObjectRefAccessibleChild},
	{"atk_object_ref_relation_set", &ObjectRefRelationSet},
	{"atk_object_ref_state_set", &ObjectRefStateSet},
	{"atk_object_remove_property_change_handler", &ObjectRemovePropertyChangeHandler},
	{"atk_object_remove_relationship", &ObjectRemoveRelationship},
	{"atk_object_set_description", &ObjectSetDescription},
	{"atk_object_set_name", &ObjectSetName},
	{"atk_object_set_parent", &ObjectSetParent},
	{"atk_object_set_role", &ObjectSetRole},
	{"atk_rectangle_get_type", &RectangleGetType},
	{"atk_registry_get_factory", &RegistryGetFactory},
	{"atk_registry_get_factory_type", &RegistryGetFactoryType},
	{"atk_registry_get_type", &RegistryGetType},
	{"atk_registry_set_factory_type", &RegistrySetFactoryType},
	{"atk_relation_add_target", &RelationAddTarget},
	{"atk_relation_get_relation_type", &RelationGetRelationType},
	{"atk_relation_get_target", &RelationGetTarget},
	{"atk_relation_get_type", &RelationGetType},
	{"atk_relation_new", &RelationNew},
	{"atk_relation_set_add", &RelationSetAdd},
	{"atk_relation_set_add_relation_by_type", &RelationSetAddRelationByType},
	{"atk_relation_set_contains", &RelationSetContains},
	{"atk_relation_set_get_n_relations", &RelationSetGetNRelations},
	{"atk_relation_set_get_relation", &RelationSetGetRelation},
	{"atk_relation_set_get_relation_by_type", &RelationSetGetRelationByType},
	{"atk_relation_set_get_type", &RelationSetGetType},
	{"atk_relation_set_new", &RelationSetNew},
	{"atk_relation_set_remove", &RelationSetRemove},
	{"atk_relation_type_for_name", &RelationTypeForName},
	{"atk_relation_type_get_name", &RelationTypeGetName},
	{"atk_relation_type_get_type", &RelationTypeGetType},
	{"atk_relation_type_register", &RelationTypeRegister},
	{"atk_remove_focus_tracker", &RemoveFocusTracker},
	{"atk_remove_global_event_listener", &RemoveGlobalEventListener},
	{"atk_remove_key_event_listener", &RemoveKeyEventListener},
	{"atk_role_for_name", &RoleForName},
	{"atk_role_get_localized_name", &RoleGetLocalizedName},
	{"atk_role_get_name", &RoleGetName},
	{"atk_role_get_type", &RoleGetType},
	{"atk_role_register", &RoleRegister},
	{"atk_selection_add_selection", &SelectionAddSelection},
	{"atk_selection_clear_selection", &SelectionClearSelection},
	{"atk_selection_get_selection_count", &SelectionGetSelectionCount},
	{"atk_selection_get_type", &SelectionGetType},
	{"atk_selection_is_child_selected", &SelectionIsChildSelected},
	{"atk_selection_ref_selection", &SelectionRefSelection},
	{"atk_selection_remove_selection", &SelectionRemoveSelection},
	{"atk_selection_select_all_selection", &SelectionSelectAllSelection},
	{"atk_state_set_add_state", &StateSetAddState},
	{"atk_state_set_add_states", &StateSetAddStates},
	{"atk_state_set_and_sets", &StateSetAndSets},
	{"atk_state_set_clear_states", &StateSetClearStates},
	{"atk_state_set_contains_state", &StateSetContainsState},
	{"atk_state_set_contains_states", &StateSetContainsStates},
	{"atk_state_set_get_type", &StateSetGetType},
	{"atk_state_set_is_empty", &StateSetIsEmpty},
	{"atk_state_set_new", &StateSetNew},
	{"atk_state_set_or_sets", &StateSetOrSets},
	{"atk_state_set_remove_state", &StateSetRemoveState},
	{"atk_state_set_xor_sets", &StateSetXorSets},
	{"atk_state_type_for_name", &StateTypeForName},
	{"atk_state_type_get_name", &StateTypeGetName},
	{"atk_state_type_get_type", &StateTypeGetType},
	{"atk_state_type_register", &StateTypeRegister},
	{"atk_streamable_content_get_mime_type", &StreamableContentGetMimeType},
	{"atk_streamable_content_get_n_mime_types", &StreamableContentGetNMimeTypes},
	{"atk_streamable_content_get_stream", &StreamableContentGetStream},
	{"atk_streamable_content_get_type", &StreamableContentGetType},
	{"atk_streamable_content_get_uri", &StreamableContentGetUri},
	{"atk_table_add_column_selection", &TableAddColumnSelection},
	{"atk_table_add_row_selection", &TableAddRowSelection},
	{"atk_table_get_caption", &TableGetCaption},
	{"atk_table_get_column_at_index", &TableGetColumnAtIndex},
	{"atk_table_get_column_description", &TableGetColumnDescription},
	{"atk_table_get_column_extent_at", &TableGetColumnExtentAt},
	{"atk_table_get_column_header", &TableGetColumnHeader},
	{"atk_table_get_index_at", &TableGetIndexAt},
	{"atk_table_get_n_columns", &TableGetNColumns},
	{"atk_table_get_n_rows", &TableGetNRows},
	{"atk_table_get_row_at_index", &TableGetRowAtIndex},
	{"atk_table_get_row_description", &TableGetRowDescription},
	{"atk_table_get_row_extent_at", &TableGetRowExtentAt},
	{"atk_table_get_row_header", &TableGetRowHeader},
	{"atk_table_get_selected_columns", &TableGetSelectedColumns},
	{"atk_table_get_selected_rows", &TableGetSelectedRows},
	{"atk_table_get_summary", &TableGetSummary},
	{"atk_table_get_type", &TableGetType},
	{"atk_table_is_column_selected", &TableIsColumnSelected},
	{"atk_table_is_row_selected", &TableIsRowSelected},
	{"atk_table_is_selected", &TableIsSelected},
	{"atk_table_ref_at", &TableRefAt},
	{"atk_table_remove_column_selection", &TableRemoveColumnSelection},
	{"atk_table_remove_row_selection", &TableRemoveRowSelection},
	{"atk_table_set_caption", &TableSetCaption},
	{"atk_table_set_column_description", &TableSetColumnDescription},
	{"atk_table_set_column_header", &TableSetColumnHeader},
	{"atk_table_set_row_description", &TableSetRowDescription},
	{"atk_table_set_row_header", &TableSetRowHeader},
	{"atk_table_set_summary", &TableSetSummary},
	{"atk_text_add_selection", &TextAddSelection},
	{"atk_text_attribute_for_name", &TextAttributeForName},
	{"atk_text_attribute_get_name", &TextAttributeGetName},
	{"atk_text_attribute_get_type", &TextAttributeGetType},
	{"atk_text_attribute_get_value", &TextAttributeGetValue},
	{"atk_text_attribute_register", &TextAttributeRegister},
	{"atk_text_boundary_get_type", &TextBoundaryGetType},
	{"atk_text_clip_type_get_type", &TextClipTypeGetType},
	{"atk_text_free_ranges", &TextFreeRanges},
	{"atk_text_get_bounded_ranges", &TextGetBoundedRanges},
	{"atk_text_get_caret_offset", &TextGetCaretOffset},
	{"atk_text_get_character_at_offset", &TextGetCharacterAtOffset},
	{"atk_text_get_character_count", &TextGetCharacterCount},
	{"atk_text_get_character_extents", &TextGetCharacterExtents},
	{"atk_text_get_default_attributes", &TextGetDefaultAttributes},
	{"atk_text_get_n_selections", &TextGetNSelections},
	{"atk_text_get_offset_at_point", &TextGetOffsetAtPoint},
	{"atk_text_get_range_extents", &TextGetRangeExtents},
	{"atk_text_get_run_attributes", &TextGetRunAttributes},
	{"atk_text_get_selection", &TextGetSelection},
	{"atk_text_get_text", &TextGetText},
	{"atk_text_get_text_after_offset", &TextGetTextAfterOffset},
	{"atk_text_get_text_at_offset", &TextGetTextAtOffset},
	{"atk_text_get_text_before_offset", &TextGetTextBeforeOffset},
	{"atk_text_get_type", &TextGetType},
	{"atk_text_remove_selection", &TextRemoveSelection},
	{"atk_text_set_caret_offset", &TextSetCaretOffset},
	{"atk_text_set_selection", &TextSetSelection},
	{"atk_util_get_type", &UtilGetType},
	{"atk_value_get_current_value", &ValueGetCurrentValue},
	{"atk_value_get_maximum_value", &ValueGetMaximumValue},
	{"atk_value_get_minimum_increment", &ValueGetMinimumIncrement},
	{"atk_value_get_minimum_value", &ValueGetMinimumValue},
	{"atk_value_get_type", &ValueGetType},
	{"atk_value_set_current_value", &ValueSetCurrentValue},
}

var dataList = outside.Data{
	{"atk_misc_instance", (*Misc)(nil)},
}
