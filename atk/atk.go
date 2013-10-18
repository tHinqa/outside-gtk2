// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package atk provides API definitions for accessing
//libatk-1.0-0.dll.
package atk

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
	outside.AddDllData(dll, false, dataList)
}

var (
	StateTypeRegister func(
		name string) T.AtkStateType

	StateTypeGetName func(
		t T.AtkStateType) string

	StateTypeForName func(
		name string) T.AtkStateType

	RoleRegister func(
		name string) T.AtkRole

	ObjectGetType func() T.GType

	ImplementorGetType func() T.GType

	ImplementorRefAccessible func(
		implementor *T.AtkImplementor) *T.AtkObject

	ObjectGetName func(
		accessible *T.AtkObject) string

	ObjectGetDescription func(
		accessible *T.AtkObject) string

	ObjectGetParent func(
		accessible *T.AtkObject) *T.AtkObject

	ObjectGetNAccessibleChildren func(
		accessible *T.AtkObject) int

	ObjectRefAccessibleChild func(
		accessible *T.AtkObject,
		i int) *T.AtkObject

	ObjectRefRelationSet func(
		accessible *T.AtkObject) *T.AtkRelationSet

	ObjectGetRole func(
		accessible *T.AtkObject) T.AtkRole

	ObjectGetLayer func(
		accessible *T.AtkObject) T.AtkLayer

	ObjectGetMdiZorder func(
		accessible *T.AtkObject) int

	ObjectGetAttributes func(
		accessible *T.AtkObject) *T.AtkAttributeSet

	ObjectRefStateSet func(
		accessible *T.AtkObject) *T.AtkStateSet

	ObjectGetIndexInParent func(
		accessible *T.AtkObject) int

	ObjectSetName func(
		accessible *T.AtkObject,
		name string)

	ObjectSetDescription func(
		accessible *T.AtkObject,
		description string)

	ObjectSetParent func(
		accessible *T.AtkObject,
		parent *T.AtkObject)

	ObjectSetRole func(
		accessible *T.AtkObject,
		role T.AtkRole)

	ObjectConnectPropertyChangeHandler func(
		accessible *T.AtkObject,
		handler *T.AtkPropertyChangeHandler) uint

	ObjectRemovePropertyChangeHandler func(
		accessible *T.AtkObject,
		handlerId uint)

	ObjectNotifyStateChange func(
		accessible *T.AtkObject,
		state T.AtkState,
		value T.Gboolean)

	ObjectInitialize func(
		accessible *T.AtkObject,
		data T.Gpointer)

	RoleGetName func(
		role T.AtkRole) string

	RoleForName func(
		name string) T.AtkRole

	ObjectAddRelationship func(
		object *T.AtkObject,
		relationship T.AtkRelationType,
		target *T.AtkObject) T.Gboolean

	ObjectRemoveRelationship func(
		object *T.AtkObject,
		relationship T.AtkRelationType,
		target *T.AtkObject) T.Gboolean

	RoleGetLocalizedName func(
		role T.AtkRole) string

	ActionGetType func() T.GType

	ActionDoAction func(
		action *T.AtkAction,
		i int) T.Gboolean

	ActionGetNActions func(
		action *T.AtkAction) int

	ActionGetDescription func(
		action *T.AtkAction,
		i int) string

	ActionGetName func(
		action *T.AtkAction,
		i int) string

	ActionGetKeybinding func(
		action *T.AtkAction,
		i int) string

	ActionSetDescription func(
		action *T.AtkAction,
		i int,
		desc string) T.Gboolean

	ActionGetLocalizedName func(
		action *T.AtkAction,
		i int) string

	UtilGetType func() T.GType

	AddFocusTracker func(
		focusTracker T.AtkEventListener) uint

	RemoveFocusTracker func(
		trackerId uint)

	FocusTrackerInit func(
		init T.AtkEventListenerInit)

	FocusTrackerNotify func(
		object *T.AtkObject)

	AddGlobalEventListener func(
		listener T.GSignalEmissionHook,
		eventType string) uint

	RemoveGlobalEventListener func(
		listenerId uint)

	AddKeyEventListener func(
		listener T.AtkKeySnoopFunc,
		data T.Gpointer) uint

	RemoveKeyEventListener func(
		listenerId uint)

	GetRoot func() *T.AtkObject

	GetFocusObject func() *T.AtkObject

	GetToolkitName func() string

	GetToolkitVersion func() string

	//Atk_get_version func() int //NOTE: Not exposed

	RectangleGetType func() T.GType

	ComponentGetType func() T.GType

	ComponentAddFocusHandler func(
		component *T.AtkComponent,
		handler T.AtkFocusHandler) uint

	ComponentContains func(
		component *T.AtkComponent,
		x int,
		y int,
		coordType T.AtkCoordType) T.Gboolean

	ComponentRefAccessibleAtPoint func(
		component *T.AtkComponent,
		x int,
		y int,
		coordType T.AtkCoordType) *T.AtkObject

	ComponentGetExtents func(
		component *T.AtkComponent,
		x *int,
		y *int,
		width *int,
		height *int,
		coordType T.AtkCoordType)

	ComponentGetPosition func(
		component *T.AtkComponent,
		x *int,
		y *int,
		coordType T.AtkCoordType)

	ComponentGetSize func(
		component *T.AtkComponent,
		width *int,
		height *int)

	ComponentGetLayer func(
		component *T.AtkComponent) T.AtkLayer

	ComponentGetMdiZorder func(
		component *T.AtkComponent) int

	ComponentGrabFocus func(
		component *T.AtkComponent) T.Gboolean

	ComponentRemoveFocusHandler func(
		component *T.AtkComponent,
		handlerId uint)

	ComponentSetExtents func(
		component *T.AtkComponent,
		x int,
		y int,
		width int,
		height int,
		coordType T.AtkCoordType) T.Gboolean

	ComponentSetPosition func(
		component *T.AtkComponent,
		x int,
		y int,
		coordType T.AtkCoordType) T.Gboolean

	ComponentSetSize func(
		component *T.AtkComponent,
		width int,
		height int) T.Gboolean

	ComponentGetAlpha func(
		component *T.AtkComponent) float64

	DocumentGetType func() T.GType

	DocumentGetDocumentType func(
		document *T.AtkDocument) string

	DocumentGetDocument func(
		document *T.AtkDocument) T.Gpointer

	DocumentGetLocale func(
		document *T.AtkDocument) string

	DocumentGetAttributes func(
		document *T.AtkDocument) *T.AtkAttributeSet

	DocumentGetAttributeValue func(
		document *T.AtkDocument,
		attributeName string) string

	DocumentSetAttributeValue func(
		document *T.AtkDocument,
		attributeName string,
		attributeValue string) T.Gboolean

	TextAttributeRegister func(
		name string) T.AtkTextAttribute

	TextGetType func() T.GType

	TextGetText func(
		text *T.AtkText,
		startOffset int,
		endOffset int) string

	TextGetCharacterAtOffset func(
		text *T.AtkText,
		offset int) T.Gunichar

	TextGetTextAfterOffset func(
		text *T.AtkText,
		offset int,
		boundaryType T.AtkTextBoundary,
		startOffset *int,
		endOffset *int) string

	TextGetTextAtOffset func(
		text *T.AtkText,
		offset int,
		boundaryType T.AtkTextBoundary,
		startOffset *int,
		endOffset *int) string

	TextGetTextBeforeOffset func(
		text *T.AtkText,
		offset int,
		boundaryType T.AtkTextBoundary,
		startOffset *int,
		endOffset *int) string

	TextGetCaretOffset func(
		text *T.AtkText) int

	TextGetCharacterExtents func(
		text *T.AtkText,
		offset int,
		x *int,
		y *int,
		width *int,
		height *int,
		coords T.AtkCoordType)

	TextGetRunAttributes func(
		text *T.AtkText,
		offset int,
		startOffset *int,
		endOffset *int) *T.AtkAttributeSet

	TextGetDefaultAttributes func(
		text *T.AtkText) *T.AtkAttributeSet

	TextGetCharacterCount func(
		text *T.AtkText) int

	TextGetOffsetAtPoint func(
		text *T.AtkText,
		x int,
		y int,
		coords T.AtkCoordType) int

	TextGetNSelections func(
		text *T.AtkText) int

	TextGetSelection func(
		text *T.AtkText,
		selectionNum int,
		startOffset *int,
		endOffset *int) string

	TextAddSelection func(
		text *T.AtkText,
		startOffset int,
		endOffset int) T.Gboolean

	TextRemoveSelection func(
		text *T.AtkText,
		selectionNum int) T.Gboolean

	TextSetSelection func(
		text *T.AtkText,
		selectionNum int,
		startOffset int,
		endOffset int) T.Gboolean

	TextSetCaretOffset func(
		text *T.AtkText, offset int) T.Gboolean

	TextGetRangeExtents func(
		text *T.AtkText,
		startOffset, endOffset int,
		coordType T.AtkCoordType,
		rect *T.AtkTextRectangle)

	TextGetBoundedRanges func(
		text *T.AtkText,
		rect *T.AtkTextRectangle,
		coordType T.AtkCoordType,
		xClipType,
		yClipType T.AtkTextClipType) **T.AtkTextRange

	TextFreeRanges func(ranges **T.AtkTextRange)

	AttributeSetFree func(attribSet *T.AtkAttributeSet)

	TextAttributeGetName func(
		attr T.AtkTextAttribute) string

	TextAttributeForName func(
		name string) T.AtkTextAttribute

	TextAttributeGetValue func(
		attr T.AtkTextAttribute, index int) string

	EditableTextGetType func() T.GType

	EditableTextSetRunAttributes func(
		text *T.AtkEditableText,
		attribSet *T.AtkAttributeSet,
		startOffset int,
		endOffset int) T.Gboolean

	EditableTextSetTextContents func(
		text *T.AtkEditableText,
		string string)

	EditableTextInsertText func(
		text *T.AtkEditableText,
		string string,
		length int,
		position *int)

	EditableTextCopyText func(
		text *T.AtkEditableText,
		startPos int,
		endPos int)

	EditableTextCutText func(
		text *T.AtkEditableText,
		startPos int,
		endPos int)

	EditableTextDeleteText func(
		text *T.AtkEditableText,
		startPos int,
		endPos int)

	EditableTextPasteText func(
		text *T.AtkEditableText,
		position int)

	GobjectAccessibleForObject func(
		obj *T.GObject) *T.AtkObject

	GobjectAccessibleGetObject func(
		obj *T.AtkGObjectAccessible) *T.GObject

	HyperlinkGetType func() T.GType

	HyperlinkGetUri func(
		link *T.AtkHyperlink,
		i int) string

	HyperlinkGetObject func(
		link *T.AtkHyperlink,
		i int) *T.AtkObject

	HyperlinkGetEndIndex func(
		link *T.AtkHyperlink) int

	HyperlinkGetStartIndex func(
		link *T.AtkHyperlink) int

	HyperlinkIsValid func(
		link *T.AtkHyperlink) T.Gboolean

	HyperlinkIsInline func(
		link *T.AtkHyperlink) T.Gboolean

	HyperlinkGetNAnchors func(
		link *T.AtkHyperlink) int

	HyperlinkIsSelectedLink func(
		link *T.AtkHyperlink) T.Gboolean

	HyperlinkImplGetType func() T.GType

	HyperlinkImplGetHyperlink func(
		obj *T.AtkHyperlinkImpl) *T.AtkHyperlink

	HypertextGetType func() T.GType

	HypertextGetLink func(
		hypertext *T.AtkHypertext,
		linkIndex int) *T.AtkHyperlink

	HypertextGetNLinks func(
		hypertext *T.AtkHypertext) int

	HypertextGetLinkIndex func(
		hypertext *T.AtkHypertext,
		charIndex int) int

	ImageGetType func() T.GType

	ImageGetImageDescription func(
		image *T.AtkImage) string

	ImageGetImageSize func(
		image *T.AtkImage,
		width *int,
		height *int)

	ImageSetImageDescription func(
		image *T.AtkImage,
		description string) T.Gboolean

	ImageGetImagePosition func(
		image *T.AtkImage,
		x *int,
		y *int,
		coordType T.AtkCoordType)

	ImageGetImageLocale func(
		image *T.AtkImage) string

	NoOpObjectNew func(
		obj *T.GObject) *T.AtkObject

	ObjectFactoryGetType func() T.GType

	ObjectFactoryCreateAccessible func(
		factory *T.AtkObjectFactory,
		obj *T.GObject) *T.AtkObject

	ObjectFactoryInvalidate func(
		factory *T.AtkObjectFactory)

	ObjectFactoryGetAccessibleType func(
		factory *T.AtkObjectFactory) T.GType

	NoOpObjectFactoryGetType func() T.GType

	NoOpObjectFactoryNew func() *T.AtkObjectFactory

	PlugGetType func() T.GType

	PlugNew func() *T.AtkObject

	PlugGetId func(plug *T.AtkPlug) string

	RegistryGetType func() T.GType

	RegistrySetFactoryType func(
		registry *T.AtkRegistry,
		t T.GType,
		factoryType T.GType)

	RegistryGetFactoryType func(
		registry *T.AtkRegistry,
		t T.GType) T.GType

	RegistryGetFactory func(
		registry *T.AtkRegistry,
		t T.GType) *T.AtkObjectFactory

	GetDefaultRegistry func() *T.AtkRegistry

	RelationGetType func() T.GType

	RelationTypeRegister func(
		name string) T.AtkRelationType

	RelationTypeGetName func(
		t T.AtkRelationType) string

	RelationTypeForName func(
		name string) T.AtkRelationType

	RelationNew func(
		targets **T.AtkObject,
		nTargets int,
		relationship T.AtkRelationType) *T.AtkRelation

	RelationGetRelationType func(
		relation *T.AtkRelation) T.AtkRelationType

	RelationGetTarget func(
		relation *T.AtkRelation) *T.GPtrArray

	RelationAddTarget func(
		relation *T.AtkRelation,
		target *T.AtkObject)

	RelationRemoveTarget func(
		relation *T.AtkRelation,
		target *T.AtkObject) T.Gboolean

	RelationSetGetType func() T.GType

	RelationSetNew func() *T.AtkRelationSet

	RelationSetContains func(
		set *T.AtkRelationSet,
		relationship T.AtkRelationType) T.Gboolean

	RelationSetRemove func(
		set *T.AtkRelationSet,
		relation *T.AtkRelation)

	RelationSetAdd func(
		set *T.AtkRelationSet,
		relation *T.AtkRelation)

	RelationSetGetNRelations func(
		set *T.AtkRelationSet) int

	RelationSetGetRelation func(
		set *T.AtkRelationSet,
		i int) *T.AtkRelation

	RelationSetGetRelationByType func(
		set *T.AtkRelationSet,
		relationship T.AtkRelationType) *T.AtkRelation

	RelationSetAddRelationByType func(
		set *T.AtkRelationSet,
		relationship T.AtkRelationType,
		target *T.AtkObject)

	SelectionGetType func() T.GType

	SelectionAddSelection func(
		selection *T.AtkSelection,
		i int) T.Gboolean

	SelectionClearSelection func(
		selection *T.AtkSelection) T.Gboolean

	SelectionRefSelection func(
		selection *T.AtkSelection,
		i int) *T.AtkObject

	SelectionGetSelectionCount func(
		selection *T.AtkSelection) int

	SelectionIsChildSelected func(
		selection *T.AtkSelection,
		i int) T.Gboolean

	SelectionRemoveSelection func(
		selection *T.AtkSelection,
		i int) T.Gboolean

	SelectionSelectAllSelection func(
		selection *T.AtkSelection) T.Gboolean

	SocketNew func() *T.AtkObject

	SocketEmbed func(
		obj *T.AtkSocket,
		plugId string)

	SocketIsOccupied func(
		obj *T.AtkSocket) T.Gboolean

	StateSetGetType func() T.GType

	StateSetNew func() *T.AtkStateSet

	StateSetIsEmpty func(
		set *T.AtkStateSet) T.Gboolean

	StateSetAddState func(
		set *T.AtkStateSet,
		t T.AtkStateType) T.Gboolean

	StateSetAddStates func(
		set *T.AtkStateSet,
		types *T.AtkStateType,
		nTypes int)

	StateSetClearStates func(
		set *T.AtkStateSet)

	StateSetContainsState func(
		set *T.AtkStateSet,
		t T.AtkStateType) T.Gboolean

	StateSetContainsStates func(
		set *T.AtkStateSet,
		types *T.AtkStateType,
		nTypes int) T.Gboolean

	StateSetRemoveState func(
		set *T.AtkStateSet,
		t T.AtkStateType) T.Gboolean

	StateSetAndSets func(
		set *T.AtkStateSet,
		compareSet *T.AtkStateSet) *T.AtkStateSet

	StateSetOrSets func(
		set *T.AtkStateSet,
		compareSet *T.AtkStateSet) *T.AtkStateSet

	StateSetXorSets func(
		set *T.AtkStateSet,
		compareSet *T.AtkStateSet) *T.AtkStateSet

	StreamableContentGetType func() T.GType

	StreamableContentGetNMimeTypes func(
		streamable *T.AtkStreamableContent) int

	StreamableContentGetMimeType func(
		streamable *T.AtkStreamableContent,
		i int) string

	StreamableContentGetStream func(
		streamable *T.AtkStreamableContent,
		mimeType string) *T.GIOChannel

	StreamableContentGetUri func(
		streamable *T.AtkStreamableContent,
		mimeType string) string

	TableGetType func() T.GType

	TableRefAt func(
		table *T.AtkTable,
		row int,
		column int) *T.AtkObject

	TableGetIndexAt func(
		table *T.AtkTable,
		row int,
		column int) int

	TableGetColumnAtIndex func(
		table *T.AtkTable,
		index int) int

	TableGetRowAtIndex func(
		table *T.AtkTable,
		index int) int

	TableGetNColumns func(
		table *T.AtkTable) int

	TableGetNRows func(
		table *T.AtkTable) int

	TableGetColumnExtentAt func(
		table *T.AtkTable,
		row int,
		column int) int

	TableGetRowExtentAt func(
		table *T.AtkTable,
		row int,
		column int) int

	TableGetCaption func(
		table *T.AtkTable) *T.AtkObject

	TableGetColumnDescription func(
		table *T.AtkTable,
		column int) string

	TableGetColumnHeader func(
		table *T.AtkTable,
		column int) *T.AtkObject

	TableGetRowDescription func(
		table *T.AtkTable,
		row int) string

	TableGetRowHeader func(
		table *T.AtkTable,
		row int) *T.AtkObject

	TableGetSummary func(
		table *T.AtkTable) *T.AtkObject

	TableSetCaption func(
		table *T.AtkTable,
		caption *T.AtkObject)

	TableSetColumnDescription func(
		table *T.AtkTable,
		column int,
		description string)

	TableSetColumnHeader func(
		table *T.AtkTable,
		column int,
		header *T.AtkObject)

	TableSetRowDescription func(
		table *T.AtkTable,
		row int,
		description string)

	TableSetRowHeader func(
		table *T.AtkTable,
		row int,
		header *T.AtkObject)

	TableSetSummary func(
		table *T.AtkTable,
		accessible *T.AtkObject)

	TableGetSelectedColumns func(
		table *T.AtkTable,
		selected **int) int

	TableGetSelectedRows func(
		table *T.AtkTable,
		selected **int) int

	TableIsColumnSelected func(
		table *T.AtkTable,
		column int) T.Gboolean

	TableIsRowSelected func(
		table *T.AtkTable,
		row int) T.Gboolean

	TableIsSelected func(
		table *T.AtkTable,
		row int,
		column int) T.Gboolean

	TableAddRowSelection func(
		table *T.AtkTable,
		row int) T.Gboolean

	TableRemoveRowSelection func(
		table *T.AtkTable,
		row int) T.Gboolean

	TableAddColumnSelection func(
		table *T.AtkTable,
		column int) T.Gboolean

	TableRemoveColumnSelection func(
		table *T.AtkTable,
		column int) T.Gboolean

	MiscGetType func() T.GType

	MiscThreadsEnter func(
		misc *T.AtkMisc)

	MiscThreadsLeave func(
		misc *T.AtkMisc)

	MiscGetInstance func() *T.AtkMisc

	ValueGetType func() T.GType

	ValueGetCurrentValue func(
		obj *T.AtkValue, value *T.GValue)

	ValueGetMaximumValue func(
		obj *T.AtkValue, value *T.GValue)

	ValueGetMinimumValue func(
		obj *T.AtkValue, value *T.GValue)

	ValueSetCurrentValue func(
		obj *T.AtkValue, value *T.GValue) T.Gboolean

	ValueGetMinimumIncrement func(
		obj *T.AtkValue, value *T.GValue)

	CoordTypeGetType func() T.GType

	GobjectAccessibleGetType func() T.GType

	HyperlinkStateFlagsGetType func() T.GType

	KeyEventTypeGetType func() T.GType

	LayerGetType func() T.GType

	NoOpObjectGetType func() T.GType

	RelationTypeGetType func() T.GType

	RoleGetType func() T.GType

	StateTypeGetType func() T.GType

	TextAttributeGetType func() T.GType

	TextBoundaryGetType func() T.GType

	TextClipTypeGetType func() T.GType
)

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
	{"atk_misc_instance", (*T.AtkMisc)(nil)},
}
