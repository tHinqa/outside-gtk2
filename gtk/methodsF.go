// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type FileChooser struct{}

type FileChooserAction T.Enum

const (
	FILE_CHOOSER_ACTION_OPEN FileChooserAction = iota
	FILE_CHOOSER_ACTION_SAVE
	FILE_CHOOSER_ACTION_SELECT_FOLDER
	FILE_CHOOSER_ACTION_CREATE_FOLDER
)

type FileChooserConfirmation T.Enum

const (
	FILE_CHOOSER_CONFIRMATION_CONFIRM FileChooserConfirmation = iota
	FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME
	FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN
)

type FileChooserError T.Enum

const (
	FILE_CHOOSER_ERROR_NONEXISTENT FileChooserError = iota
	FILE_CHOOSER_ERROR_BAD_FILENAME
	FILE_CHOOSER_ERROR_ALREADY_EXISTS
	FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME
)

var (
	FileChooserGetType    func() T.GType
	FileChooserErrorQuark func() T.GQuark

	FileChooserActionGetType       func() T.GType
	FileChooserConfirmationGetType func() T.GType
	FileChooserDialogGetType       func() T.GType
	FileChooserErrorGetType        func() T.GType
	FileChooserWidgetGetType       func() T.GType

	FileChooserDialogNew            func(title string, parent *Window, action FileChooserAction, firstButtonText string, v ...VArg) *Widget
	FileChooserDialogNewWithBackend func(title string, parent *Window, action FileChooserAction, backend string, firstButtonText string, v ...VArg) *Widget
	FileChooserWidgetNew            func(action FileChooserAction) *Widget
	FileChooserWidgetNewWithBackend func(action FileChooserAction, backend string) *Widget

	fileChooserAddFilter                  func(f *FileChooser, filter *FileFilter)
	fileChooserAddShortcutFolder          func(f *FileChooser, folder string, error **T.GError) T.Gboolean
	fileChooserAddShortcutFolderUri       func(f *FileChooser, uri string, error **T.GError) T.Gboolean
	fileChooserGetAction                  func(f *FileChooser) FileChooserAction
	fileChooserGetCreateFolders           func(f *FileChooser) T.Gboolean
	fileChooserGetCurrentFolder           func(f *FileChooser) string
	fileChooserGetCurrentFolderFile       func(f *FileChooser) *T.GFile
	fileChooserGetCurrentFolderUri        func(f *FileChooser) string
	fileChooserGetDoOverwriteConfirmation func(f *FileChooser) T.Gboolean
	fileChooserGetExtraWidget             func(f *FileChooser) *Widget
	fileChooserGetFile                    func(f *FileChooser) *T.GFile
	fileChooserGetFilename                func(f *FileChooser) string
	fileChooserGetFilenames               func(f *FileChooser) *T.GSList
	fileChooserGetFiles                   func(f *FileChooser) *T.GSList
	fileChooserGetFilter                  func(f *FileChooser) *FileFilter
	fileChooserGetLocalOnly               func(f *FileChooser) T.Gboolean
	fileChooserGetPreviewFile             func(f *FileChooser) *T.GFile
	fileChooserGetPreviewFilename         func(f *FileChooser) string
	fileChooserGetPreviewUri              func(f *FileChooser) string
	fileChooserGetPreviewWidget           func(f *FileChooser) *Widget
	fileChooserGetPreviewWidgetActive     func(f *FileChooser) T.Gboolean
	fileChooserGetSelectMultiple          func(f *FileChooser) T.Gboolean
	fileChooserGetShowHidden              func(f *FileChooser) T.Gboolean
	fileChooserGetUri                     func(f *FileChooser) string
	fileChooserGetUris                    func(f *FileChooser) *T.GSList
	fileChooserGetUsePreviewLabel         func(f *FileChooser) T.Gboolean
	fileChooserListFilters                func(f *FileChooser) *T.GSList
	fileChooserListShortcutFolders        func(f *FileChooser) *T.GSList
	fileChooserListShortcutFolderUris     func(f *FileChooser) *T.GSList
	fileChooserRemoveFilter               func(f *FileChooser, filter *FileFilter)
	fileChooserRemoveShortcutFolder       func(f *FileChooser, folder string, error **T.GError) T.Gboolean
	fileChooserRemoveShortcutFolderUri    func(f *FileChooser, uri string, error **T.GError) T.Gboolean
	fileChooserSelectAll                  func(f *FileChooser)
	fileChooserSelectFile                 func(f *FileChooser, file *T.GFile, error **T.GError) T.Gboolean
	fileChooserSelectFilename             func(f *FileChooser, filename string) T.Gboolean
	fileChooserSelectUri                  func(f *FileChooser, uri string) T.Gboolean
	fileChooserSetAction                  func(f *FileChooser, action FileChooserAction)
	fileChooserSetCreateFolders           func(f *FileChooser, createFolders T.Gboolean)
	fileChooserSetCurrentFolder           func(f *FileChooser, filename string) T.Gboolean
	fileChooserSetCurrentFolderFile       func(f *FileChooser, file *T.GFile, error **T.GError) T.Gboolean
	fileChooserSetCurrentFolderUri        func(f *FileChooser, uri string) T.Gboolean
	fileChooserSetCurrentName             func(f *FileChooser, name string)
	fileChooserSetDoOverwriteConfirmation func(f *FileChooser, doOverwriteConfirmation T.Gboolean)
	fileChooserSetExtraWidget             func(f *FileChooser, extraWidget *Widget)
	fileChooserSetFile                    func(f *FileChooser, file *T.GFile, error **T.GError) T.Gboolean
	fileChooserSetFilename                func(f *FileChooser, filename string) T.Gboolean
	fileChooserSetFilter                  func(f *FileChooser, filter *FileFilter)
	fileChooserSetLocalOnly               func(f *FileChooser, localOnly T.Gboolean)
	fileChooserSetPreviewWidget           func(f *FileChooser, previewWidget *Widget)
	fileChooserSetPreviewWidgetActive     func(f *FileChooser, active T.Gboolean)
	fileChooserSetSelectMultiple          func(f *FileChooser, selectMultiple T.Gboolean)
	fileChooserSetShowHidden              func(f *FileChooser, showHidden T.Gboolean)
	fileChooserSetUri                     func(f *FileChooser, uri string) T.Gboolean
	fileChooserSetUsePreviewLabel         func(f *FileChooser, useLabel T.Gboolean)
	fileChooserUnselectAll                func(f *FileChooser)
	fileChooserUnselectFile               func(f *FileChooser, file *T.GFile)
	fileChooserUnselectFilename           func(f *FileChooser, filename string)
	fileChooserUnselectUri                func(f *FileChooser, uri string)
)

func (f *FileChooser) AddFilter(filter *FileFilter) { fileChooserAddFilter(f, filter) }
func (f *FileChooser) AddShortcutFolder(folder string, err **T.GError) T.Gboolean {
	return fileChooserAddShortcutFolder(f, folder, err)
}
func (f *FileChooser) AddShortcutFolderUri(uri string, err **T.GError) T.Gboolean {
	return fileChooserAddShortcutFolderUri(f, uri, err)
}
func (f *FileChooser) GetAction() FileChooserAction   { return fileChooserGetAction(f) }
func (f *FileChooser) GetCreateFolders() T.Gboolean   { return fileChooserGetCreateFolders(f) }
func (f *FileChooser) GetCurrentFolder() string       { return fileChooserGetCurrentFolder(f) }
func (f *FileChooser) GetCurrentFolderFile() *T.GFile { return fileChooserGetCurrentFolderFile(f) }
func (f *FileChooser) GetCurrentFolderUri() string    { return fileChooserGetCurrentFolderUri(f) }
func (f *FileChooser) GetDoOverwriteConfirmation() T.Gboolean {
	return fileChooserGetDoOverwriteConfirmation(f)
}
func (f *FileChooser) GetExtraWidget() *Widget            { return fileChooserGetExtraWidget(f) }
func (f *FileChooser) GetFile() *T.GFile                  { return fileChooserGetFile(f) }
func (f *FileChooser) GetFilename() string                { return fileChooserGetFilename(f) }
func (f *FileChooser) GetFilenames() *T.GSList            { return fileChooserGetFilenames(f) }
func (f *FileChooser) GetFiles() *T.GSList                { return fileChooserGetFiles(f) }
func (f *FileChooser) GetFilter() *FileFilter             { return fileChooserGetFilter(f) }
func (f *FileChooser) GetLocalOnly() T.Gboolean           { return fileChooserGetLocalOnly(f) }
func (f *FileChooser) GetPreviewFile() *T.GFile           { return fileChooserGetPreviewFile(f) }
func (f *FileChooser) GetPreviewFilename() string         { return fileChooserGetPreviewFilename(f) }
func (f *FileChooser) GetPreviewUri() string              { return fileChooserGetPreviewUri(f) }
func (f *FileChooser) GetPreviewWidget() *Widget          { return fileChooserGetPreviewWidget(f) }
func (f *FileChooser) GetPreviewWidgetActive() T.Gboolean { return fileChooserGetPreviewWidgetActive(f) }
func (f *FileChooser) GetSelectMultiple() T.Gboolean      { return fileChooserGetSelectMultiple(f) }
func (f *FileChooser) GetShowHidden() T.Gboolean          { return fileChooserGetShowHidden(f) }
func (f *FileChooser) GetUri() string                     { return fileChooserGetUri(f) }
func (f *FileChooser) GetUris() *T.GSList                 { return fileChooserGetUris(f) }
func (f *FileChooser) GetUsePreviewLabel() T.Gboolean     { return fileChooserGetUsePreviewLabel(f) }
func (f *FileChooser) ListFilters() *T.GSList             { return fileChooserListFilters(f) }
func (f *FileChooser) ListShortcutFolders() *T.GSList     { return fileChooserListShortcutFolders(f) }
func (f *FileChooser) ListShortcutFolderUris() *T.GSList  { return fileChooserListShortcutFolderUris(f) }
func (f *FileChooser) RemoveFilter(filter *FileFilter)    { fileChooserRemoveFilter(f, filter) }
func (f *FileChooser) RemoveShortcutFolder(folder string, err **T.GError) T.Gboolean {
	return fileChooserRemoveShortcutFolder(f, folder, err)
}
func (f *FileChooser) RemoveShortcutFolderUri(uri string, err **T.GError) T.Gboolean {
	return fileChooserRemoveShortcutFolderUri(f, uri, err)
}
func (f *FileChooser) SelectAll() { fileChooserSelectAll(f) }
func (f *FileChooser) SelectFile(file *T.GFile, err **T.GError) T.Gboolean {
	return fileChooserSelectFile(f, file, err)
}
func (f *FileChooser) SelectFilename(filename string) T.Gboolean {
	return fileChooserSelectFilename(f, filename)
}
func (f *FileChooser) SelectUri(uri string) T.Gboolean    { return fileChooserSelectUri(f, uri) }
func (f *FileChooser) SetAction(action FileChooserAction) { fileChooserSetAction(f, action) }
func (f *FileChooser) SetCreateFolders(createFolders T.Gboolean) {
	fileChooserSetCreateFolders(f, createFolders)
}
func (f *FileChooser) SetCurrentFolder(filename string) T.Gboolean {
	return fileChooserSetCurrentFolder(f, filename)
}
func (f *FileChooser) SetCurrentFolderFile(file *T.GFile, err **T.GError) T.Gboolean {
	return fileChooserSetCurrentFolderFile(f, file, err)
}
func (f *FileChooser) SetCurrentFolderUri(uri string) T.Gboolean {
	return fileChooserSetCurrentFolderUri(f, uri)
}
func (f *FileChooser) SetCurrentName(name string) { fileChooserSetCurrentName(f, name) }
func (f *FileChooser) SetDoOverwriteConfirmation(doOverwriteConfirmation T.Gboolean) {
	fileChooserSetDoOverwriteConfirmation(f, doOverwriteConfirmation)
}
func (f *FileChooser) SetExtraWidget(extraWidget *Widget) { fileChooserSetExtraWidget(f, extraWidget) }
func (f *FileChooser) SetFile(file *T.GFile, err **T.GError) T.Gboolean {
	return fileChooserSetFile(f, file, err)
}
func (f *FileChooser) SetFilename(filename string) T.Gboolean {
	return fileChooserSetFilename(f, filename)
}
func (f *FileChooser) SetFilter(filter *FileFilter)      { fileChooserSetFilter(f, filter) }
func (f *FileChooser) SetLocalOnly(localOnly T.Gboolean) { fileChooserSetLocalOnly(f, localOnly) }
func (f *FileChooser) SetPreviewWidget(previewWidget *Widget) {
	fileChooserSetPreviewWidget(f, previewWidget)
}
func (f *FileChooser) SetPreviewWidgetActive(active T.Gboolean) {
	fileChooserSetPreviewWidgetActive(f, active)
}
func (f *FileChooser) SetSelectMultiple(selectMultiple T.Gboolean) {
	fileChooserSetSelectMultiple(f, selectMultiple)
}
func (f *FileChooser) SetShowHidden(showHidden T.Gboolean) { fileChooserSetShowHidden(f, showHidden) }
func (f *FileChooser) SetUri(uri string) T.Gboolean        { return fileChooserSetUri(f, uri) }
func (f *FileChooser) SetUsePreviewLabel(useLabel T.Gboolean) {
	fileChooserSetUsePreviewLabel(f, useLabel)
}
func (f *FileChooser) UnselectAll()                     { fileChooserUnselectAll(f) }
func (f *FileChooser) UnselectFile(file *T.GFile)       { fileChooserUnselectFile(f, file) }
func (f *FileChooser) UnselectFilename(filename string) { fileChooserUnselectFilename(f, filename) }
func (f *FileChooser) UnselectUri(uri string)           { fileChooserUnselectUri(f, uri) }

type FileChooserButton struct {
	Parent HBox
	_      *struct{}
}

var (
	FileChooserButtonGetType        func() T.GType
	FileChooserButtonNew            func(title string, action FileChooserAction) *Widget
	FileChooserButtonNewWithBackend func(title string, action FileChooserAction, backend string) *Widget
	FileChooserButtonNewWithDialog  func(dialog *Widget) *Widget

	fileChooserButtonGetTitle        func(f *FileChooserButton) string
	fileChooserButtonSetTitle        func(f *FileChooserButton, title string)
	fileChooserButtonGetWidthChars   func(f *FileChooserButton) int
	fileChooserButtonSetWidthChars   func(f *FileChooserButton, nChars int)
	fileChooserButtonGetFocusOnClick func(f *FileChooserButton) T.Gboolean
	fileChooserButtonSetFocusOnClick func(f *FileChooserButton, focusOnClick T.Gboolean)
)

func (f *FileChooserButton) ButtonGetTitle() string      { return fileChooserButtonGetTitle(f) }
func (f *FileChooserButton) GetFocusOnClick() T.Gboolean { return fileChooserButtonGetFocusOnClick(f) }
func (f *FileChooserButton) GetWidthChars() int          { return fileChooserButtonGetWidthChars(f) }
func (f *FileChooserButton) SetFocusOnClick(focusOnClick T.Gboolean) {
	fileChooserButtonSetFocusOnClick(f, focusOnClick)
}
func (f *FileChooserButton) SetTitle(title string)    { fileChooserButtonSetTitle(f, title) }
func (f *FileChooserButton) SetWidthChars(nChars int) { fileChooserButtonSetWidthChars(f, nChars) }

type (
	FileFilter struct{}

	FileFilterFlags T.Enum

	FileFilterInfo struct {
		Contains    FileFilterFlags
		Filename    *T.Gchar
		Uri         *T.Gchar
		DisplayName *T.Gchar
		MimeType    *T.Gchar
	}

	FileFilterFunc func(
		filterInfo *FileFilterInfo,
		data T.Gpointer) T.Gboolean
)

const (
	FILE_FILTER_FILENAME FileFilterFlags = 1 << iota
	FILE_FILTER_URI
	FILE_FILTER_DISPLAY_NAME
	FILE_FILTER_MIME_TYPE
)

var (
	FileFilterGetType      func() T.GType
	FileFilterFlagsGetType func() T.GType
	FileFilterNew          func() *FileFilter

	fileFilterAddCustom        func(filter *FileFilter, needed FileFilterFlags, f FileFilterFunc, data T.Gpointer, notify T.GDestroyNotify)
	fileFilterAddMimeType      func(filter *FileFilter, mimeType string)
	fileFilterAddPattern       func(filter *FileFilter, pattern string)
	fileFilterAddPixbufFormats func(filter *FileFilter)
	fileFilterFilter           func(filter *FileFilter, filterInfo *FileFilterInfo) T.Gboolean
	fileFilterGetName          func(filter *FileFilter) string
	fileFilterGetNeeded        func(filter *FileFilter) FileFilterFlags
	fileFilterSetName          func(filter *FileFilter, name string)
)

func (f *FileFilter) AddCustom(needed FileFilterFlags, fnc FileFilterFunc, data T.Gpointer, notify T.GDestroyNotify) {
	fileFilterAddCustom(f, needed, fnc, data, notify)
}
func (f *FileFilter) AddMimeType(mimeType string) { fileFilterAddMimeType(f, mimeType) }
func (f *FileFilter) AddPattern(pattern string)   { fileFilterAddPattern(f, pattern) }
func (f *FileFilter) AddPixbufFormats()           { fileFilterAddPixbufFormats(f) }
func (f *FileFilter) Filter(filterInfo *FileFilterInfo) T.Gboolean {
	return fileFilterFilter(f, filterInfo)
}
func (f *FileFilter) GetName() string            { return fileFilterGetName(f) }
func (f *FileFilter) GetNeeded() FileFilterFlags { return fileFilterGetNeeded(f) }
func (f *FileFilter) SetName(name string)        { fileFilterSetName(f, name) }

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
	History_list     *T.GList
	Fileop_dialog    *Widget
	Fileop_entry     *Widget
	Fileop_file      *T.Gchar
	Cmpl_state       T.Gpointer
	Fileop_c_dir     *Widget
	Fileop_del_file  *Widget
	Fileop_ren_file  *Widget
	Button_area      *Widget
	Action_area      *Widget
	Selected_names   *T.GPtrArray
	Last_selected    *T.Gchar
}

var (
	FileSelectionGetType func() T.GType
	FileSelectionNew     func(title string) *Widget

	fileSelectionComplete          func(f *FileSelection, pattern string)
	fileSelectionGetFilename       func(f *FileSelection) string
	fileSelectionGetSelections     func(f *FileSelection) **T.Gchar
	fileSelectionGetSelectMultiple func(f *FileSelection) T.Gboolean
	fileSelectionHideFileopButtons func(f *FileSelection)
	fileSelectionSetFilename       func(f *FileSelection, filename string)
	fileSelectionSetSelectMultiple func(f *FileSelection, selectMultiple T.Gboolean)
	fileSelectionShowFileopButtons func(f *FileSelection)
)

func (f *FileSelection) Complete(pattern string)       { fileSelectionComplete(f, pattern) }
func (f *FileSelection) GetFilename() string           { return fileSelectionGetFilename(f) }
func (f *FileSelection) GetSelections() **T.Gchar      { return fileSelectionGetSelections(f) }
func (f *FileSelection) GetSelectMultiple() T.Gboolean { return fileSelectionGetSelectMultiple(f) }
func (f *FileSelection) HideFileopButtons()            { fileSelectionHideFileopButtons(f) }
func (f *FileSelection) SetFilename(filename string)   { fileSelectionSetFilename(f, filename) }
func (f *FileSelection) SetSelectMultiple(selectMultiple T.Gboolean) {
	fileSelectionSetSelectMultiple(f, selectMultiple)
}
func (f *FileSelection) ShowFileopButtons() { fileSelectionShowFileopButtons(f) }

type Fixed struct {
	Container Container
	Children  *T.GList
}

var (
	FixedGetType func() T.GType
	FixedNew     func() *Widget

	fixedGetHasWindow func(f *Fixed) T.Gboolean
	fixedMove         func(f *Fixed, widget *Widget, x, y int)
	fixedPut          func(f *Fixed, widget *Widget, x, y int)
	fixedSetHasWindow func(f *Fixed, hasWindow T.Gboolean)
)

func (f *Fixed) GetHasWindow() T.Gboolean          { return fixedGetHasWindow(f) }
func (f *Fixed) Move(widget *Widget, x, y int)     { fixedMove(f, widget, x, y) }
func (f *Fixed) Put(widget *Widget, x, y int)      { fixedPut(f, widget, x, y) }
func (f *Fixed) SetHasWindow(hasWindow T.Gboolean) { fixedSetHasWindow(f, hasWindow) }

type FlagValue T.GFlagsValue

type FontButton struct {
	Button Button
	_      *struct{}
}

var (
	FontButtonGetType     func() T.GType
	FontButtonNew         func() *Widget
	FontButtonNewWithFont func(fontname string) *Widget

	fontButtonGetFontName  func(f *FontButton) string
	fontButtonGetShowSize  func(f *FontButton) T.Gboolean
	fontButtonGetShowStyle func(f *FontButton) T.Gboolean
	fontButtonGetTitle     func(f *FontButton) string
	fontButtonGetUseFont   func(f *FontButton) T.Gboolean
	fontButtonGetUseSize   func(f *FontButton) T.Gboolean
	fontButtonSetFontName  func(f *FontButton, fontname string) T.Gboolean
	fontButtonSetShowSize  func(f *FontButton, showSize T.Gboolean)
	fontButtonSetShowStyle func(f *FontButton, showStyle T.Gboolean)
	fontButtonSetTitle     func(f *FontButton, title string)
	fontButtonSetUseFont   func(f *FontButton, useFont T.Gboolean)
	fontButtonSetUseSize   func(f *FontButton, useSize T.Gboolean)
)

func (f *FontButton) GetFontName() string      { return fontButtonGetFontName(f) }
func (f *FontButton) GetShowSize() T.Gboolean  { return fontButtonGetShowSize(f) }
func (f *FontButton) GetShowStyle() T.Gboolean { return fontButtonGetShowStyle(f) }
func (f *FontButton) GetTitle() string         { return fontButtonGetTitle(f) }
func (f *FontButton) GetUseFont() T.Gboolean   { return fontButtonGetUseFont(f) }
func (f *FontButton) GetUseSize() T.Gboolean   { return fontButtonGetUseSize(f) }
func (f *FontButton) SetFontName(fontname string) T.Gboolean {
	return fontButtonSetFontName(f, fontname)
}
func (f *FontButton) SetShowSize(showSize T.Gboolean)   { fontButtonSetShowSize(f, showSize) }
func (f *FontButton) SetShowStyle(showStyle T.Gboolean) { fontButtonSetShowStyle(f, showStyle) }
func (f *FontButton) SetTitle(title string)             { fontButtonSetTitle(f, title) }
func (f *FontButton) SetUseFont(useFont T.Gboolean)     { fontButtonSetUseFont(f, useFont) }
func (f *FontButton) SetUseSize(useSize T.Gboolean)     { fontButtonSetUseSize(f, useSize) }

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
	Family           *T.PangoFontFamily
	Face             *T.PangoFontFace
	Size             int
	Font             *T.GdkFont
}

var (
	FontSelectionGetType func() T.GType
	FontSelectionNew     func() *Widget

	fontSelectionGetFace         func(f *FontSelection) *T.PangoFontFace
	fontSelectionGetFaceList     func(f *FontSelection) *Widget
	fontSelectionGetFamily       func(f *FontSelection) *T.PangoFontFamily
	fontSelectionGetFamilyList   func(f *FontSelection) *Widget
	fontSelectionGetFont         func(f *FontSelection) *T.GdkFont
	fontSelectionGetFontName     func(f *FontSelection) string
	fontSelectionGetPreviewEntry func(f *FontSelection) *Widget
	fontSelectionGetPreviewText  func(f *FontSelection) string
	fontSelectionGetSize         func(f *FontSelection) int
	fontSelectionGetSizeEntry    func(f *FontSelection) *Widget
	fontSelectionGetSizeList     func(f *FontSelection) *Widget
	fontSelectionSetFontName     func(f *FontSelection, fontname string) T.Gboolean
	fontSelectionSetPreviewText  func(f *FontSelection, text string)
)

func (f *FontSelection) GetFace() *T.PangoFontFace     { return fontSelectionGetFace(f) }
func (f *FontSelection) GetFaceList() *Widget          { return fontSelectionGetFaceList(f) }
func (f *FontSelection) GetFamily() *T.PangoFontFamily { return fontSelectionGetFamily(f) }
func (f *FontSelection) GetFamilyList() *Widget        { return fontSelectionGetFamilyList(f) }
func (f *FontSelection) GetFont() *T.GdkFont           { return fontSelectionGetFont(f) }
func (f *FontSelection) GetFontName() string           { return fontSelectionGetFontName(f) }
func (f *FontSelection) GetPreviewEntry() *Widget      { return fontSelectionGetPreviewEntry(f) }
func (f *FontSelection) GetPreviewText() string        { return fontSelectionGetPreviewText(f) }
func (f *FontSelection) GetSize() int                  { return fontSelectionGetSize(f) }
func (f *FontSelection) GetSizeEntry() *Widget         { return fontSelectionGetSizeEntry(f) }
func (f *FontSelection) GetSizeList() *Widget          { return fontSelectionGetSizeList(f) }
func (f *FontSelection) SetFontName(fontname string) T.Gboolean {
	return fontSelectionSetFontName(f, fontname)
}
func (f *FontSelection) SetPreviewText(text string) { fontSelectionSetPreviewText(f, text) }

type FontSelectionDialog struct {
	Parent       Dialog
	Fontsel      *Widget
	MainVbox     *Widget
	ActionArea   *Widget
	OkButton     *Widget
	ApplyButton  *Widget
	CancelButton *Widget
	DialogWidth  int
	AutoResize   T.Gboolean
}

var (
	FontSelectionDialogGetType func() T.GType
	FontSelectionDialogNew     func(title string) *Widget

	fontSelectionDialogGetApplyButton   func(f *FontSelectionDialog) *Widget
	fontSelectionDialogGetCancelButton  func(f *FontSelectionDialog) *Widget
	fontSelectionDialogGetFont          func(f *FontSelectionDialog) *T.GdkFont
	fontSelectionDialogGetFontName      func(f *FontSelectionDialog) string
	fontSelectionDialogGetFontSelection func(f *FontSelectionDialog) *Widget
	fontSelectionDialogGetOkButton      func(f *FontSelectionDialog) *Widget
	fontSelectionDialogGetPreviewText   func(f *FontSelectionDialog) string
	fontSelectionDialogSetFontName      func(f *FontSelectionDialog, fontname string) T.Gboolean
	fontSelectionDialogSetPreviewText   func(f *FontSelectionDialog, text string)
)

func (f *FontSelectionDialog) GetApplyButton() *Widget  { return fontSelectionDialogGetApplyButton(f) }
func (f *FontSelectionDialog) GetCancelButton() *Widget { return fontSelectionDialogGetCancelButton(f) }
func (f *FontSelectionDialog) GetFont() *T.GdkFont      { return fontSelectionDialogGetFont(f) }
func (f *FontSelectionDialog) GetFontName() string      { return fontSelectionDialogGetFontName(f) }
func (f *FontSelectionDialog) GetFontSelection() *Widget {
	return fontSelectionDialogGetFontSelection(f)
}
func (f *FontSelectionDialog) GetOkButton() *Widget   { return fontSelectionDialogGetOkButton(f) }
func (f *FontSelectionDialog) GetPreviewText() string { return fontSelectionDialogGetPreviewText(f) }
func (f *FontSelectionDialog) SetFontName(fontname string) T.Gboolean {
	return fontSelectionDialogSetFontName(f, fontname)
}
func (f *FontSelectionDialog) SetPreviewText(text string) { fontSelectionDialogSetPreviewText(f, text) }

type Frame struct {
	Bin             Bin
	LabelWidget     *Widget
	ShadowType      int16
	LabelXalign     float32
	LabelYalign     float32
	ChildAllocation Allocation
}

var (
	FrameGetType func() T.GType
	FrameNew     func(label string) *Widget

	frameGetLabel       func(f *Frame) string
	frameGetLabelAlign  func(f *Frame, xalign, yalign *float32)
	frameGetLabelWidget func(f *Frame) *Widget
	frameGetShadowType  func(f *Frame) ShadowType
	frameSetLabel       func(f *Frame, label string)
	frameSetLabelAlign  func(f *Frame, xalign, yalign float32)
	frameSetLabelWidget func(f *Frame, labelWidget *Widget)
	frameSetShadowType  func(f *Frame, t ShadowType)
)

func (f *Frame) GetLabel() string                      { return frameGetLabel(f) }
func (f *Frame) GetLabelAlign(xalign, yalign *float32) { frameGetLabelAlign(f, xalign, yalign) }
func (f *Frame) GetLabelWidget() *Widget               { return frameGetLabelWidget(f) }
func (f *Frame) GetShadowType() ShadowType             { return frameGetShadowType(f) }
func (f *Frame) SetLabel(label string)                 { frameSetLabel(f, label) }
func (f *Frame) SetLabelAlign(xalign, yalign float32)  { frameSetLabelAlign(f, xalign, yalign) }
func (f *Frame) SetLabelWidget(labelWidget *Widget)    { frameSetLabelWidget(f, labelWidget) }
func (f *Frame) SetShadowType(t ShadowType)            { frameSetShadowType(f, t) }

type Function func(data T.Gpointer) T.Gboolean
