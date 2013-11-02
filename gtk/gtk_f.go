// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	D "github.com/tHinqa/outside-gtk2/gdk"
	I "github.com/tHinqa/outside-gtk2/gio"
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	P "github.com/tHinqa/outside-gtk2/pango"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type FileChooser struct{}

type FileChooserAction Enum

const (
	FILE_CHOOSER_ACTION_OPEN FileChooserAction = iota
	FILE_CHOOSER_ACTION_SAVE
	FILE_CHOOSER_ACTION_SELECT_FOLDER
	FILE_CHOOSER_ACTION_CREATE_FOLDER
)

type FileChooserConfirmation Enum

const (
	FILE_CHOOSER_CONFIRMATION_CONFIRM FileChooserConfirmation = iota
	FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME
	FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN
)

type FileChooserError Enum

const (
	FILE_CHOOSER_ERROR_NONEXISTENT FileChooserError = iota
	FILE_CHOOSER_ERROR_BAD_FILENAME
	FILE_CHOOSER_ERROR_ALREADY_EXISTS
	FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME
)

var (
	FileChooserGetType    func() O.Type
	FileChooserErrorQuark func() L.Quark

	FileChooserActionGetType       func() O.Type
	FileChooserConfirmationGetType func() O.Type
	FileChooserDialogGetType       func() O.Type
	FileChooserErrorGetType        func() O.Type
	FileChooserWidgetGetType       func() O.Type

	FileChooserDialogNew            func(title string, parent *Window, action FileChooserAction, firstButtonText string, v ...VArg) *Widget
	FileChooserDialogNewWithBackend func(title string, parent *Window, action FileChooserAction, backend string, firstButtonText string, v ...VArg) *Widget
	FileChooserWidgetNew            func(action FileChooserAction) *Widget
	FileChooserWidgetNewWithBackend func(action FileChooserAction, backend string) *Widget

	FileChooserAddFilter                  func(f *FileChooser, filter *FileFilter)
	FileChooserAddShortcutFolder          func(f *FileChooser, folder string, error **L.Error) bool
	FileChooserAddShortcutFolderUri       func(f *FileChooser, uri string, error **L.Error) bool
	FileChooserGetAction                  func(f *FileChooser) FileChooserAction
	FileChooserGetCreateFolders           func(f *FileChooser) bool
	FileChooserGetCurrentFolder           func(f *FileChooser) string
	FileChooserGetCurrentFolderFile       func(f *FileChooser) *I.File
	FileChooserGetCurrentFolderUri        func(f *FileChooser) string
	FileChooserGetDoOverwriteConfirmation func(f *FileChooser) bool
	FileChooserGetExtraWidget             func(f *FileChooser) *Widget
	FileChooserGetFile                    func(f *FileChooser) *I.File
	FileChooserGetFilename                func(f *FileChooser) string
	FileChooserGetFilenames               func(f *FileChooser) *L.SList
	FileChooserGetFiles                   func(f *FileChooser) *L.SList
	FileChooserGetFilter                  func(f *FileChooser) *FileFilter
	FileChooserGetLocalOnly               func(f *FileChooser) bool
	FileChooserGetPreviewFile             func(f *FileChooser) *I.File
	FileChooserGetPreviewFilename         func(f *FileChooser) string
	FileChooserGetPreviewUri              func(f *FileChooser) string
	FileChooserGetPreviewWidget           func(f *FileChooser) *Widget
	FileChooserGetPreviewWidgetActive     func(f *FileChooser) bool
	FileChooserGetSelectMultiple          func(f *FileChooser) bool
	FileChooserGetShowHidden              func(f *FileChooser) bool
	FileChooserGetUri                     func(f *FileChooser) string
	FileChooserGetUris                    func(f *FileChooser) *L.SList
	FileChooserGetUsePreviewLabel         func(f *FileChooser) bool
	FileChooserListFilters                func(f *FileChooser) *L.SList
	FileChooserListShortcutFolders        func(f *FileChooser) *L.SList
	FileChooserListShortcutFolderUris     func(f *FileChooser) *L.SList
	FileChooserRemoveFilter               func(f *FileChooser, filter *FileFilter)
	FileChooserRemoveShortcutFolder       func(f *FileChooser, folder string, error **L.Error) bool
	FileChooserRemoveShortcutFolderUri    func(f *FileChooser, uri string, error **L.Error) bool
	FileChooserSelectAll                  func(f *FileChooser)
	FileChooserSelectFile                 func(f *FileChooser, file *I.File, error **L.Error) bool
	FileChooserSelectFilename             func(f *FileChooser, filename string) bool
	FileChooserSelectUri                  func(f *FileChooser, uri string) bool
	FileChooserSetAction                  func(f *FileChooser, action FileChooserAction)
	FileChooserSetCreateFolders           func(f *FileChooser, createFolders bool)
	FileChooserSetCurrentFolder           func(f *FileChooser, filename string) bool
	FileChooserSetCurrentFolderFile       func(f *FileChooser, file *I.File, error **L.Error) bool
	FileChooserSetCurrentFolderUri        func(f *FileChooser, uri string) bool
	FileChooserSetCurrentName             func(f *FileChooser, name string)
	FileChooserSetDoOverwriteConfirmation func(f *FileChooser, doOverwriteConfirmation bool)
	FileChooserSetExtraWidget             func(f *FileChooser, extraWidget *Widget)
	FileChooserSetFile                    func(f *FileChooser, file *I.File, error **L.Error) bool
	FileChooserSetFilename                func(f *FileChooser, filename string) bool
	FileChooserSetFilter                  func(f *FileChooser, filter *FileFilter)
	FileChooserSetLocalOnly               func(f *FileChooser, localOnly bool)
	FileChooserSetPreviewWidget           func(f *FileChooser, previewWidget *Widget)
	FileChooserSetPreviewWidgetActive     func(f *FileChooser, active bool)
	FileChooserSetSelectMultiple          func(f *FileChooser, selectMultiple bool)
	FileChooserSetShowHidden              func(f *FileChooser, showHidden bool)
	FileChooserSetUri                     func(f *FileChooser, uri string) bool
	FileChooserSetUsePreviewLabel         func(f *FileChooser, useLabel bool)
	FileChooserUnselectAll                func(f *FileChooser)
	FileChooserUnselectFile               func(f *FileChooser, file *I.File)
	FileChooserUnselectFilename           func(f *FileChooser, filename string)
	FileChooserUnselectUri                func(f *FileChooser, uri string)
)

func (f *FileChooser) AddFilter(filter *FileFilter) { FileChooserAddFilter(f, filter) }
func (f *FileChooser) AddShortcutFolder(folder string, err **L.Error) bool {
	return FileChooserAddShortcutFolder(f, folder, err)
}
func (f *FileChooser) AddShortcutFolderUri(uri string, err **L.Error) bool {
	return FileChooserAddShortcutFolderUri(f, uri, err)
}
func (f *FileChooser) GetAction() FileChooserAction  { return FileChooserGetAction(f) }
func (f *FileChooser) GetCreateFolders() bool        { return FileChooserGetCreateFolders(f) }
func (f *FileChooser) GetCurrentFolder() string      { return FileChooserGetCurrentFolder(f) }
func (f *FileChooser) GetCurrentFolderFile() *I.File { return FileChooserGetCurrentFolderFile(f) }
func (f *FileChooser) GetCurrentFolderUri() string   { return FileChooserGetCurrentFolderUri(f) }
func (f *FileChooser) GetDoOverwriteConfirmation() bool {
	return FileChooserGetDoOverwriteConfirmation(f)
}
func (f *FileChooser) GetExtraWidget() *Widget          { return FileChooserGetExtraWidget(f) }
func (f *FileChooser) GetFile() *I.File                 { return FileChooserGetFile(f) }
func (f *FileChooser) GetFilename() string              { return FileChooserGetFilename(f) }
func (f *FileChooser) GetFilenames() *L.SList           { return FileChooserGetFilenames(f) }
func (f *FileChooser) GetFiles() *L.SList               { return FileChooserGetFiles(f) }
func (f *FileChooser) GetFilter() *FileFilter           { return FileChooserGetFilter(f) }
func (f *FileChooser) GetLocalOnly() bool               { return FileChooserGetLocalOnly(f) }
func (f *FileChooser) GetPreviewFile() *I.File          { return FileChooserGetPreviewFile(f) }
func (f *FileChooser) GetPreviewFilename() string       { return FileChooserGetPreviewFilename(f) }
func (f *FileChooser) GetPreviewUri() string            { return FileChooserGetPreviewUri(f) }
func (f *FileChooser) GetPreviewWidget() *Widget        { return FileChooserGetPreviewWidget(f) }
func (f *FileChooser) GetPreviewWidgetActive() bool     { return FileChooserGetPreviewWidgetActive(f) }
func (f *FileChooser) GetSelectMultiple() bool          { return FileChooserGetSelectMultiple(f) }
func (f *FileChooser) GetShowHidden() bool              { return FileChooserGetShowHidden(f) }
func (f *FileChooser) GetUri() string                   { return FileChooserGetUri(f) }
func (f *FileChooser) GetUris() *L.SList                { return FileChooserGetUris(f) }
func (f *FileChooser) GetUsePreviewLabel() bool         { return FileChooserGetUsePreviewLabel(f) }
func (f *FileChooser) ListFilters() *L.SList            { return FileChooserListFilters(f) }
func (f *FileChooser) ListShortcutFolders() *L.SList    { return FileChooserListShortcutFolders(f) }
func (f *FileChooser) ListShortcutFolderUris() *L.SList { return FileChooserListShortcutFolderUris(f) }
func (f *FileChooser) RemoveFilter(filter *FileFilter)  { FileChooserRemoveFilter(f, filter) }
func (f *FileChooser) RemoveShortcutFolder(folder string, err **L.Error) bool {
	return FileChooserRemoveShortcutFolder(f, folder, err)
}
func (f *FileChooser) RemoveShortcutFolderUri(uri string, err **L.Error) bool {
	return FileChooserRemoveShortcutFolderUri(f, uri, err)
}
func (f *FileChooser) SelectAll() { FileChooserSelectAll(f) }
func (f *FileChooser) SelectFile(file *I.File, err **L.Error) bool {
	return FileChooserSelectFile(f, file, err)
}
func (f *FileChooser) SelectFilename(filename string) bool {
	return FileChooserSelectFilename(f, filename)
}
func (f *FileChooser) SelectUri(uri string) bool          { return FileChooserSelectUri(f, uri) }
func (f *FileChooser) SetAction(action FileChooserAction) { FileChooserSetAction(f, action) }
func (f *FileChooser) SetCreateFolders(createFolders bool) {
	FileChooserSetCreateFolders(f, createFolders)
}
func (f *FileChooser) SetCurrentFolder(filename string) bool {
	return FileChooserSetCurrentFolder(f, filename)
}
func (f *FileChooser) SetCurrentFolderFile(file *I.File, err **L.Error) bool {
	return FileChooserSetCurrentFolderFile(f, file, err)
}
func (f *FileChooser) SetCurrentFolderUri(uri string) bool {
	return FileChooserSetCurrentFolderUri(f, uri)
}
func (f *FileChooser) SetCurrentName(name string) { FileChooserSetCurrentName(f, name) }
func (f *FileChooser) SetDoOverwriteConfirmation(doOverwriteConfirmation bool) {
	FileChooserSetDoOverwriteConfirmation(f, doOverwriteConfirmation)
}
func (f *FileChooser) SetExtraWidget(extraWidget *Widget) { FileChooserSetExtraWidget(f, extraWidget) }
func (f *FileChooser) SetFile(file *I.File, err **L.Error) bool {
	return FileChooserSetFile(f, file, err)
}
func (f *FileChooser) SetFilename(filename string) bool {
	return FileChooserSetFilename(f, filename)
}
func (f *FileChooser) SetFilter(filter *FileFilter) { FileChooserSetFilter(f, filter) }
func (f *FileChooser) SetLocalOnly(localOnly bool)  { FileChooserSetLocalOnly(f, localOnly) }
func (f *FileChooser) SetPreviewWidget(previewWidget *Widget) {
	FileChooserSetPreviewWidget(f, previewWidget)
}
func (f *FileChooser) SetPreviewWidgetActive(active bool) {
	FileChooserSetPreviewWidgetActive(f, active)
}
func (f *FileChooser) SetSelectMultiple(selectMultiple bool) {
	FileChooserSetSelectMultiple(f, selectMultiple)
}
func (f *FileChooser) SetShowHidden(showHidden bool) { FileChooserSetShowHidden(f, showHidden) }
func (f *FileChooser) SetUri(uri string) bool        { return FileChooserSetUri(f, uri) }
func (f *FileChooser) SetUsePreviewLabel(useLabel bool) {
	FileChooserSetUsePreviewLabel(f, useLabel)
}
func (f *FileChooser) UnselectAll()                     { FileChooserUnselectAll(f) }
func (f *FileChooser) UnselectFile(file *I.File)        { FileChooserUnselectFile(f, file) }
func (f *FileChooser) UnselectFilename(filename string) { FileChooserUnselectFilename(f, filename) }
func (f *FileChooser) UnselectUri(uri string)           { FileChooserUnselectUri(f, uri) }

type FileChooserButton struct {
	Parent HBox
	_      *struct{}
}

var (
	FileChooserButtonGetType        func() O.Type
	FileChooserButtonNew            func(title string, action FileChooserAction) *Widget
	FileChooserButtonNewWithBackend func(title string, action FileChooserAction, backend string) *Widget
	FileChooserButtonNewWithDialog  func(dialog *Widget) *Widget

	FileChooserButtonGetTitle        func(f *FileChooserButton) string
	FileChooserButtonSetTitle        func(f *FileChooserButton, title string)
	FileChooserButtonGetWidthChars   func(f *FileChooserButton) int
	FileChooserButtonSetWidthChars   func(f *FileChooserButton, nChars int)
	FileChooserButtonGetFocusOnClick func(f *FileChooserButton) bool
	FileChooserButtonSetFocusOnClick func(f *FileChooserButton, focusOnClick bool)
)

func (f *FileChooserButton) ButtonGetTitle() string { return FileChooserButtonGetTitle(f) }
func (f *FileChooserButton) GetFocusOnClick() bool  { return FileChooserButtonGetFocusOnClick(f) }
func (f *FileChooserButton) GetWidthChars() int     { return FileChooserButtonGetWidthChars(f) }
func (f *FileChooserButton) SetFocusOnClick(focusOnClick bool) {
	FileChooserButtonSetFocusOnClick(f, focusOnClick)
}
func (f *FileChooserButton) SetTitle(title string)    { FileChooserButtonSetTitle(f, title) }
func (f *FileChooserButton) SetWidthChars(nChars int) { FileChooserButtonSetWidthChars(f, nChars) }

type (
	FileFilter struct{}

	FileFilterFlags Enum

	FileFilterInfo struct {
		Contains    FileFilterFlags
		Filename    *T.Gchar
		Uri         *T.Gchar
		DisplayName *T.Gchar
		MimeType    *T.Gchar
	}

	FileFilterFunc func(
		filterInfo *FileFilterInfo,
		data T.Gpointer) bool
)

const (
	FILE_FILTER_FILENAME FileFilterFlags = 1 << iota
	FILE_FILTER_URI
	FILE_FILTER_DISPLAY_NAME
	FILE_FILTER_MIME_TYPE
)

var (
	FileFilterGetType      func() O.Type
	FileFilterFlagsGetType func() O.Type
	FileFilterNew          func() *FileFilter

	FileFilterAddCustom        func(filter *FileFilter, needed FileFilterFlags, f FileFilterFunc, data T.Gpointer, notify O.DestroyNotify)
	FileFilterAddMimeType      func(filter *FileFilter, mimeType string)
	FileFilterAddPattern       func(filter *FileFilter, pattern string)
	FileFilterAddPixbufFormats func(filter *FileFilter)
	FileFilterFilter           func(filter *FileFilter, filterInfo *FileFilterInfo) bool
	FileFilterGetName          func(filter *FileFilter) string
	FileFilterGetNeeded        func(filter *FileFilter) FileFilterFlags
	FileFilterSetName          func(filter *FileFilter, name string)
)

func (f *FileFilter) AddCustom(needed FileFilterFlags, fnc FileFilterFunc, data T.Gpointer, notify O.DestroyNotify) {
	FileFilterAddCustom(f, needed, fnc, data, notify)
}
func (f *FileFilter) AddMimeType(mimeType string) { FileFilterAddMimeType(f, mimeType) }
func (f *FileFilter) AddPattern(pattern string)   { FileFilterAddPattern(f, pattern) }
func (f *FileFilter) AddPixbufFormats()           { FileFilterAddPixbufFormats(f) }
func (f *FileFilter) Filter(filterInfo *FileFilterInfo) bool {
	return FileFilterFilter(f, filterInfo)
}
func (f *FileFilter) GetName() string            { return FileFilterGetName(f) }
func (f *FileFilter) GetNeeded() FileFilterFlags { return FileFilterGetNeeded(f) }
func (f *FileFilter) SetName(name string)        { FileFilterSetName(f, name) }

type FileSelection struct {
	Parent_instance  Dialog
	Dir_list         *Widget
	File_list        *Widget
	Selection_entry  *Widget
	Selection_text   *Widget
	Main_vbox        *Widget
	Ok_button        *Widget
	Cancel_button    *Widget
	Help_button      *Widget
	History_pulldown *Widget
	History_menu     *Widget
	History_list     *L.List
	Fileop_dialog    *Widget
	Fileop_entry     *Widget
	Fileop_file      *T.Gchar
	Cmpl_state       T.Gpointer
	Fileop_c_dir     *Widget
	Fileop_del_file  *Widget
	Fileop_ren_file  *Widget
	Button_area      *Widget
	Action_area      *Widget
	Selected_names   *L.PtrArray
	Last_selected    *T.Gchar
}

var (
	FileSelectionGetType func() O.Type
	FileSelectionNew     func(title string) *Widget

	FileSelectionComplete          func(f *FileSelection, pattern string)
	FileSelectionGetFilename       func(f *FileSelection) string
	FileSelectionGetSelections     func(f *FileSelection) []string
	FileSelectionGetSelectMultiple func(f *FileSelection) bool
	FileSelectionHideFileopButtons func(f *FileSelection)
	FileSelectionSetFilename       func(f *FileSelection, filename string)
	FileSelectionSetSelectMultiple func(f *FileSelection, selectMultiple bool)
	FileSelectionShowFileopButtons func(f *FileSelection)
)

func (f *FileSelection) Complete(pattern string)     { FileSelectionComplete(f, pattern) }
func (f *FileSelection) GetFilename() string         { return FileSelectionGetFilename(f) }
func (f *FileSelection) GetSelections() []string     { return FileSelectionGetSelections(f) }
func (f *FileSelection) GetSelectMultiple() bool     { return FileSelectionGetSelectMultiple(f) }
func (f *FileSelection) HideFileopButtons()          { FileSelectionHideFileopButtons(f) }
func (f *FileSelection) SetFilename(filename string) { FileSelectionSetFilename(f, filename) }
func (f *FileSelection) SetSelectMultiple(selectMultiple bool) {
	FileSelectionSetSelectMultiple(f, selectMultiple)
}
func (f *FileSelection) ShowFileopButtons() { FileSelectionShowFileopButtons(f) }

type Fixed struct {
	Container Container
	Children  *L.List
}

var (
	FixedGetType func() O.Type
	FixedNew     func() *Widget

	FixedGetHasWindow func(f *Fixed) bool
	FixedMove         func(f *Fixed, widget *Widget, x, y int)
	FixedPut          func(f *Fixed, widget *Widget, x, y int)
	FixedSetHasWindow func(f *Fixed, hasWindow bool)
)

func (f *Fixed) GetHasWindow() bool            { return FixedGetHasWindow(f) }
func (f *Fixed) Move(widget *Widget, x, y int) { FixedMove(f, widget, x, y) }
func (f *Fixed) Put(widget *Widget, x, y int)  { FixedPut(f, widget, x, y) }
func (f *Fixed) SetHasWindow(hasWindow bool)   { FixedSetHasWindow(f, hasWindow) }

type FlagValue O.FlagsValue

type FontButton struct {
	Button Button
	_      *struct{}
}

var (
	FontButtonGetType     func() O.Type
	FontButtonNew         func() *Widget
	FontButtonNewWithFont func(fontname string) *Widget

	FontButtonGetFontName  func(f *FontButton) string
	FontButtonGetShowSize  func(f *FontButton) bool
	FontButtonGetShowStyle func(f *FontButton) bool
	FontButtonGetTitle     func(f *FontButton) string
	FontButtonGetUseFont   func(f *FontButton) bool
	FontButtonGetUseSize   func(f *FontButton) bool
	FontButtonSetFontName  func(f *FontButton, fontname string) bool
	FontButtonSetShowSize  func(f *FontButton, showSize bool)
	FontButtonSetShowStyle func(f *FontButton, showStyle bool)
	FontButtonSetTitle     func(f *FontButton, title string)
	FontButtonSetUseFont   func(f *FontButton, useFont bool)
	FontButtonSetUseSize   func(f *FontButton, useSize bool)
)

func (f *FontButton) GetFontName() string { return FontButtonGetFontName(f) }
func (f *FontButton) GetShowSize() bool   { return FontButtonGetShowSize(f) }
func (f *FontButton) GetShowStyle() bool  { return FontButtonGetShowStyle(f) }
func (f *FontButton) GetTitle() string    { return FontButtonGetTitle(f) }
func (f *FontButton) GetUseFont() bool    { return FontButtonGetUseFont(f) }
func (f *FontButton) GetUseSize() bool    { return FontButtonGetUseSize(f) }
func (f *FontButton) SetFontName(fontname string) bool {
	return FontButtonSetFontName(f, fontname)
}
func (f *FontButton) SetShowSize(showSize bool)   { FontButtonSetShowSize(f, showSize) }
func (f *FontButton) SetShowStyle(showStyle bool) { FontButtonSetShowStyle(f, showStyle) }
func (f *FontButton) SetTitle(title string)       { FontButtonSetTitle(f, title) }
func (f *FontButton) SetUseFont(useFont bool)     { FontButtonSetUseFont(f, useFont) }
func (f *FontButton) SetUseSize(useSize bool)     { FontButtonSetUseSize(f, useSize) }

type FontSelection struct {
	Parent_instance  VBox
	Font_entry       *Widget
	Family_list      *Widget
	Font_style_entry *Widget
	Face_list        *Widget
	Size_entry       *Widget
	Size_list        *Widget
	Pixels_button    *Widget
	Points_button    *Widget
	Filter_button    *Widget
	Preview_entry    *Widget
	Family           *P.FontFamily
	Face             *P.FontFace
	Size             int
	Font             *D.Font
}

var (
	FontSelectionGetType func() O.Type
	FontSelectionNew     func() *Widget

	FontSelectionGetFace         func(f *FontSelection) *P.FontFace
	FontSelectionGetFaceList     func(f *FontSelection) *Widget
	FontSelectionGetFamily       func(f *FontSelection) *P.FontFamily
	FontSelectionGetFamilyList   func(f *FontSelection) *Widget
	FontSelectionGetFont         func(f *FontSelection) *D.Font
	FontSelectionGetFontName     func(f *FontSelection) string
	FontSelectionGetPreviewEntry func(f *FontSelection) *Widget
	FontSelectionGetPreviewText  func(f *FontSelection) string
	FontSelectionGetSize         func(f *FontSelection) int
	FontSelectionGetSizeEntry    func(f *FontSelection) *Widget
	FontSelectionGetSizeList     func(f *FontSelection) *Widget
	FontSelectionSetFontName     func(f *FontSelection, fontname string) bool
	FontSelectionSetPreviewText  func(f *FontSelection, text string)
)

func (f *FontSelection) GetFace() *P.FontFace     { return FontSelectionGetFace(f) }
func (f *FontSelection) GetFaceList() *Widget     { return FontSelectionGetFaceList(f) }
func (f *FontSelection) GetFamily() *P.FontFamily { return FontSelectionGetFamily(f) }
func (f *FontSelection) GetFamilyList() *Widget   { return FontSelectionGetFamilyList(f) }
func (f *FontSelection) GetFont() *D.Font         { return FontSelectionGetFont(f) }
func (f *FontSelection) GetFontName() string      { return FontSelectionGetFontName(f) }
func (f *FontSelection) GetPreviewEntry() *Widget { return FontSelectionGetPreviewEntry(f) }
func (f *FontSelection) GetPreviewText() string   { return FontSelectionGetPreviewText(f) }
func (f *FontSelection) GetSize() int             { return FontSelectionGetSize(f) }
func (f *FontSelection) GetSizeEntry() *Widget    { return FontSelectionGetSizeEntry(f) }
func (f *FontSelection) GetSizeList() *Widget     { return FontSelectionGetSizeList(f) }
func (f *FontSelection) SetFontName(fontname string) bool {
	return FontSelectionSetFontName(f, fontname)
}
func (f *FontSelection) SetPreviewText(text string) { FontSelectionSetPreviewText(f, text) }

type FontSelectionDialog struct {
	Parent       Dialog
	Fontsel      *Widget
	MainVbox     *Widget
	ActionArea   *Widget
	OkButton     *Widget
	ApplyButton  *Widget
	CancelButton *Widget
	DialogWidth  int
	AutoResize   bool
}

var (
	FontSelectionDialogGetType func() O.Type
	FontSelectionDialogNew     func(title string) *Widget

	FontSelectionDialogGetApplyButton   func(f *FontSelectionDialog) *Widget
	FontSelectionDialogGetCancelButton  func(f *FontSelectionDialog) *Widget
	FontSelectionDialogGetFont          func(f *FontSelectionDialog) *D.Font
	FontSelectionDialogGetFontName      func(f *FontSelectionDialog) string
	FontSelectionDialogGetFontSelection func(f *FontSelectionDialog) *Widget
	FontSelectionDialogGetOkButton      func(f *FontSelectionDialog) *Widget
	FontSelectionDialogGetPreviewText   func(f *FontSelectionDialog) string
	FontSelectionDialogSetFontName      func(f *FontSelectionDialog, fontname string) bool
	FontSelectionDialogSetPreviewText   func(f *FontSelectionDialog, text string)
)

func (f *FontSelectionDialog) GetApplyButton() *Widget  { return FontSelectionDialogGetApplyButton(f) }
func (f *FontSelectionDialog) GetCancelButton() *Widget { return FontSelectionDialogGetCancelButton(f) }
func (f *FontSelectionDialog) GetFont() *D.Font         { return FontSelectionDialogGetFont(f) }
func (f *FontSelectionDialog) GetFontName() string      { return FontSelectionDialogGetFontName(f) }
func (f *FontSelectionDialog) GetFontSelection() *Widget {
	return FontSelectionDialogGetFontSelection(f)
}
func (f *FontSelectionDialog) GetOkButton() *Widget   { return FontSelectionDialogGetOkButton(f) }
func (f *FontSelectionDialog) GetPreviewText() string { return FontSelectionDialogGetPreviewText(f) }
func (f *FontSelectionDialog) SetFontName(fontname string) bool {
	return FontSelectionDialogSetFontName(f, fontname)
}
func (f *FontSelectionDialog) SetPreviewText(text string) { FontSelectionDialogSetPreviewText(f, text) }

type Frame struct {
	Bin             Bin
	LabelWidget     *Widget
	ShadowType      int16
	LabelXalign     float32
	LabelYalign     float32
	ChildAllocation Allocation
}

var (
	FrameGetType func() O.Type
	FrameNew     func(label string) *Widget

	FrameGetLabel       func(f *Frame) string
	FrameGetLabelAlign  func(f *Frame, xalign, yalign *float32)
	FrameGetLabelWidget func(f *Frame) *Widget
	FrameGetShadowType  func(f *Frame) ShadowType
	FrameSetLabel       func(f *Frame, label string)
	FrameSetLabelAlign  func(f *Frame, xalign, yalign float32)
	FrameSetLabelWidget func(f *Frame, labelWidget *Widget)
	FrameSetShadowType  func(f *Frame, t ShadowType)
)

func (f *Frame) GetLabel() string                      { return FrameGetLabel(f) }
func (f *Frame) GetLabelAlign(xalign, yalign *float32) { FrameGetLabelAlign(f, xalign, yalign) }
func (f *Frame) GetLabelWidget() *Widget               { return FrameGetLabelWidget(f) }
func (f *Frame) GetShadowType() ShadowType             { return FrameGetShadowType(f) }
func (f *Frame) SetLabel(label string)                 { FrameSetLabel(f, label) }
func (f *Frame) SetLabelAlign(xalign, yalign float32)  { FrameSetLabelAlign(f, xalign, yalign) }
func (f *Frame) SetLabelWidget(labelWidget *Widget)    { FrameSetLabelWidget(f, labelWidget) }
func (f *Frame) SetShadowType(t ShadowType)            { FrameSetShadowType(f, t) }

type Function func(data T.Gpointer) bool
