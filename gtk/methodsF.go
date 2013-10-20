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

	FileChooserDialogNew            func(title string, parent *T.GtkWindow, action FileChooserAction, firstButtonText string, v ...VArg) *T.GtkWidget
	FileChooserDialogNewWithBackend func(title string, parent *T.GtkWindow, action FileChooserAction, backend string, firstButtonText string, v ...VArg) *T.GtkWidget
	FileChooserWidgetNew            func(action FileChooserAction) *T.GtkWidget
	FileChooserWidgetNewWithBackend func(action FileChooserAction, backend string) *T.GtkWidget

	FileChooserAddFilter                  func(f *FileChooser, filter *FileFilter)
	FileChooserAddShortcutFolder          func(f *FileChooser, folder string, error **T.GError) T.Gboolean
	FileChooserAddShortcutFolderUri       func(f *FileChooser, uri string, error **T.GError) T.Gboolean
	FileChooserGetAction                  func(f *FileChooser) FileChooserAction
	FileChooserGetCreateFolders           func(f *FileChooser) T.Gboolean
	FileChooserGetCurrentFolder           func(f *FileChooser) string
	FileChooserGetCurrentFolderFile       func(f *FileChooser) *T.GFile
	FileChooserGetCurrentFolderUri        func(f *FileChooser) string
	FileChooserGetDoOverwriteConfirmation func(f *FileChooser) T.Gboolean
	FileChooserGetExtraWidget             func(f *FileChooser) *T.GtkWidget
	FileChooserGetFile                    func(f *FileChooser) *T.GFile
	FileChooserGetFilename                func(f *FileChooser) string
	FileChooserGetFilenames               func(f *FileChooser) *T.GSList
	FileChooserGetFiles                   func(f *FileChooser) *T.GSList
	FileChooserGetFilter                  func(f *FileChooser) *FileFilter
	FileChooserGetLocalOnly               func(f *FileChooser) T.Gboolean
	FileChooserGetPreviewFile             func(f *FileChooser) *T.GFile
	FileChooserGetPreviewFilename         func(f *FileChooser) string
	FileChooserGetPreviewUri              func(f *FileChooser) string
	FileChooserGetPreviewWidget           func(f *FileChooser) *T.GtkWidget
	FileChooserGetPreviewWidgetActive     func(f *FileChooser) T.Gboolean
	FileChooserGetSelectMultiple          func(f *FileChooser) T.Gboolean
	FileChooserGetShowHidden              func(f *FileChooser) T.Gboolean
	FileChooserGetUri                     func(f *FileChooser) string
	FileChooserGetUris                    func(f *FileChooser) *T.GSList
	FileChooserGetUsePreviewLabel         func(f *FileChooser) T.Gboolean
	FileChooserListFilters                func(f *FileChooser) *T.GSList
	FileChooserListShortcutFolders        func(f *FileChooser) *T.GSList
	FileChooserListShortcutFolderUris     func(f *FileChooser) *T.GSList
	FileChooserRemoveFilter               func(f *FileChooser, filter *FileFilter)
	FileChooserRemoveShortcutFolder       func(f *FileChooser, folder string, error **T.GError) T.Gboolean
	FileChooserRemoveShortcutFolderUri    func(f *FileChooser, uri string, error **T.GError) T.Gboolean
	FileChooserSelectAll                  func(f *FileChooser)
	FileChooserSelectFile                 func(f *FileChooser, file *T.GFile, error **T.GError) T.Gboolean
	FileChooserSelectFilename             func(f *FileChooser, filename string) T.Gboolean
	FileChooserSelectUri                  func(f *FileChooser, uri string) T.Gboolean
	FileChooserSetAction                  func(f *FileChooser, action FileChooserAction)
	FileChooserSetCreateFolders           func(f *FileChooser, createFolders T.Gboolean)
	FileChooserSetCurrentFolder           func(f *FileChooser, filename string) T.Gboolean
	FileChooserSetCurrentFolderFile       func(f *FileChooser, file *T.GFile, error **T.GError) T.Gboolean
	FileChooserSetCurrentFolderUri        func(f *FileChooser, uri string) T.Gboolean
	FileChooserSetCurrentName             func(f *FileChooser, name string)
	FileChooserSetDoOverwriteConfirmation func(f *FileChooser, doOverwriteConfirmation T.Gboolean)
	FileChooserSetExtraWidget             func(f *FileChooser, extraWidget *T.GtkWidget)
	FileChooserSetFile                    func(f *FileChooser, file *T.GFile, error **T.GError) T.Gboolean
	FileChooserSetFilename                func(f *FileChooser, filename string) T.Gboolean
	FileChooserSetFilter                  func(f *FileChooser, filter *FileFilter)
	FileChooserSetLocalOnly               func(f *FileChooser, localOnly T.Gboolean)
	FileChooserSetPreviewWidget           func(f *FileChooser, previewWidget *T.GtkWidget)
	FileChooserSetPreviewWidgetActive     func(f *FileChooser, active T.Gboolean)
	FileChooserSetSelectMultiple          func(f *FileChooser, selectMultiple T.Gboolean)
	FileChooserSetShowHidden              func(f *FileChooser, showHidden T.Gboolean)
	FileChooserSetUri                     func(f *FileChooser, uri string) T.Gboolean
	FileChooserSetUsePreviewLabel         func(f *FileChooser, useLabel T.Gboolean)
	FileChooserUnselectAll                func(f *FileChooser)
	FileChooserUnselectFile               func(f *FileChooser, file *T.GFile)
	FileChooserUnselectFilename           func(f *FileChooser, filename string)
	FileChooserUnselectUri                func(f *FileChooser, uri string)
)

func (f *FileChooser) SetAction(action FileChooserAction) {
	FileChooserSetAction(f, action)
}

func (f *FileChooser) GetAction() FileChooserAction {
	return FileChooserGetAction(f)
}

func (f *FileChooser) SetLocalOnly(localOnly T.Gboolean) {
	FileChooserSetLocalOnly(f, localOnly)
}

func (f *FileChooser) GetLocalOnly() T.Gboolean {
	return FileChooserGetLocalOnly(f)
}

func (f *FileChooser) SetSelectMultiple(selectMultiple T.Gboolean) {
	FileChooserSetSelectMultiple(f, selectMultiple)
}

func (f *FileChooser) GetSelectMultiple() T.Gboolean {
	return FileChooserGetSelectMultiple(f)
}

func (f *FileChooser) SetShowHidden(showHidden T.Gboolean) {
	FileChooserSetShowHidden(f, showHidden)
}

func (f *FileChooser) GetShowHidden() T.Gboolean {
	return FileChooserGetShowHidden(f)
}

func (f *FileChooser) SetDoOverwriteConfirmation(doOverwriteConfirmation T.Gboolean) {
	FileChooserSetDoOverwriteConfirmation(f, doOverwriteConfirmation)
}

func (f *FileChooser) GetDoOverwriteConfirmation() T.Gboolean {
	return FileChooserGetDoOverwriteConfirmation(f)
}

func (f *FileChooser) SetCreateFolders(createFolders T.Gboolean) {
	FileChooserSetCreateFolders(f, createFolders)
}

func (f *FileChooser) GetCreateFolders() T.Gboolean {
	return FileChooserGetCreateFolders(f)
}

func (f *FileChooser) SetCurrentName(name string) {
	FileChooserSetCurrentName(f, name)
}

func (f *FileChooser) GetFilename() string {
	return FileChooserGetFilename(f)
}

func (f *FileChooser) SetFilename(filename string) T.Gboolean {
	return FileChooserSetFilename(f, filename)
}

func (f *FileChooser) SelectFilename(filename string) T.Gboolean {
	return FileChooserSelectFilename(f, filename)
}

func (f *FileChooser) UnselectFilename(filename string) {
	FileChooserUnselectFilename(f, filename)
}

func (f *FileChooser) SelectAll() {
	FileChooserSelectAll(f)
}

func (f *FileChooser) UnselectAll() {
	FileChooserUnselectAll(f)
}

func (f *FileChooser) GetFilenames() *T.GSList {
	return FileChooserGetFilenames(f)
}

func (f *FileChooser) SetCurrentFolder(filename string) T.Gboolean {
	return FileChooserSetCurrentFolder(f, filename)
}

func (f *FileChooser) GetCurrentFolder() string {
	return FileChooserGetCurrentFolder(f)
}

func (f *FileChooser) GetUri() string {
	return FileChooserGetUri(f)
}

func (f *FileChooser) SetUri(uri string) T.Gboolean {
	return FileChooserSetUri(f, uri)
}

func (f *FileChooser) SelectUri(uri string) T.Gboolean {
	return FileChooserSelectUri(f, uri)
}

func (f *FileChooser) UnselectUri(uri string) {
	FileChooserUnselectUri(f, uri)
}

func (f *FileChooser) GetUris() *T.GSList {
	return FileChooserGetUris(f)
}

func (f *FileChooser) SetCurrentFolderUri(uri string) T.Gboolean {
	return FileChooserSetCurrentFolderUri(f, uri)
}

func (f *FileChooser) GetCurrentFolderUri() string {
	return FileChooserGetCurrentFolderUri(f)
}

func (f *FileChooser) GetFile() *T.GFile {
	return FileChooserGetFile(f)
}

func (f *FileChooser) SetFile(file *T.GFile, err **T.GError) T.Gboolean {
	return FileChooserSetFile(f, file, err)
}

func (f *FileChooser) SelectFile(file *T.GFile, err **T.GError) T.Gboolean {
	return FileChooserSelectFile(f, file, err)
}

func (f *FileChooser) UnselectFile(file *T.GFile) {
	FileChooserUnselectFile(f, file)
}

func (f *FileChooser) GetFiles() *T.GSList {
	return FileChooserGetFiles(f)
}

func (f *FileChooser) SetCurrentFolderFile(file *T.GFile, err **T.GError) T.Gboolean {
	return FileChooserSetCurrentFolderFile(f, file, err)
}

func (f *FileChooser) GetCurrentFolderFile() *T.GFile {
	return FileChooserGetCurrentFolderFile(f)
}

func (f *FileChooser) SetPreviewWidget(previewWidget *T.GtkWidget) {
	FileChooserSetPreviewWidget(f, previewWidget)
}

func (f *FileChooser) GetPreviewWidget() *T.GtkWidget {
	return FileChooserGetPreviewWidget(f)
}

func (f *FileChooser) SetPreviewWidgetActive(active T.Gboolean) {
	FileChooserSetPreviewWidgetActive(f, active)
}

func (f *FileChooser) GetPreviewWidgetActive() T.Gboolean {
	return FileChooserGetPreviewWidgetActive(f)
}

func (f *FileChooser) SetUsePreviewLabel(useLabel T.Gboolean) {
	FileChooserSetUsePreviewLabel(f, useLabel)
}

func (f *FileChooser) GetUsePreviewLabel() T.Gboolean {
	return FileChooserGetUsePreviewLabel(f)
}

func (f *FileChooser) GetPreviewFilename() string {
	return FileChooserGetPreviewFilename(f)
}

func (f *FileChooser) GetPreviewUri() string {
	return FileChooserGetPreviewUri(f)
}

func (f *FileChooser) GetPreviewFile() *T.GFile {
	return FileChooserGetPreviewFile(f)
}

func (f *FileChooser) SetExtraWidget(extraWidget *T.GtkWidget) {
	FileChooserSetExtraWidget(f, extraWidget)
}

func (f *FileChooser) GetExtraWidget() *T.GtkWidget {
	return FileChooserGetExtraWidget(f)
}

func (f *FileChooser) AddFilter(filter *FileFilter) {
	FileChooserAddFilter(f, filter)
}

func (f *FileChooser) RemoveFilter(filter *FileFilter) {
	FileChooserRemoveFilter(f, filter)
}

func (f *FileChooser) ListFilters() *T.GSList {
	return FileChooserListFilters(f)
}

func (f *FileChooser) SetFilter(filter *FileFilter) {
	FileChooserSetFilter(f, filter)
}

func (f *FileChooser) GetFilter() *FileFilter {
	return FileChooserGetFilter(f)
}

func (f *FileChooser) AddShortcutFolder(folder string, err **T.GError) T.Gboolean {
	return FileChooserAddShortcutFolder(f, folder, err)
}

func (f *FileChooser) RemoveShortcutFolder(folder string, err **T.GError) T.Gboolean {
	return FileChooserRemoveShortcutFolder(f, folder, err)
}

func (f *FileChooser) ListShortcutFolders() *T.GSList {
	return FileChooserListShortcutFolders(f)
}

func (f *FileChooser) AddShortcutFolderUri(uri string, err **T.GError) T.Gboolean {
	return FileChooserAddShortcutFolderUri(f, uri, err)
}

func (f *FileChooser) RemoveShortcutFolderUri(uri string, err **T.GError) T.Gboolean {
	return FileChooserRemoveShortcutFolderUri(f, uri, err)
}

func (f *FileChooser) ListShortcutFolderUris() *T.GSList {
	return FileChooserListShortcutFolderUris(f)
}

type FileChooserButton struct {
	Parent T.GtkHBox
	_      *struct{}
}

var (
	FileChooserButtonGetType        func() T.GType
	FileChooserButtonNew            func(title string, action FileChooserAction) *T.GtkWidget
	FileChooserButtonNewWithBackend func(title string, action FileChooserAction, backend string) *T.GtkWidget
	FileChooserButtonNewWithDialog  func(dialog *T.GtkWidget) *T.GtkWidget

	FileChooserButtonGetTitle        func(f *FileChooserButton) string
	FileChooserButtonSetTitle        func(f *FileChooserButton, title string)
	FileChooserButtonGetWidthChars   func(f *FileChooserButton) int
	FileChooserButtonSetWidthChars   func(f *FileChooserButton, nChars int)
	FileChooserButtonGetFocusOnClick func(f *FileChooserButton) T.Gboolean
	FileChooserButtonSetFocusOnClick func(f *FileChooserButton, focusOnClick T.Gboolean)
)

func (f *FileChooserButton) ButtonGetTitle() string {
	return FileChooserButtonGetTitle(f)
}

func (f *FileChooserButton) SetTitle(title string) {
	FileChooserButtonSetTitle(f, title)
}

func (f *FileChooserButton) GetWidthChars() int {
	return FileChooserButtonGetWidthChars(f)
}

func (f *FileChooserButton) SetWidthChars(nChars int) {
	FileChooserButtonSetWidthChars(f, nChars)
}

func (f *FileChooserButton) GetFocusOnClick() T.Gboolean {
	return FileChooserButtonGetFocusOnClick(f)
}

func (f *FileChooserButton) SetFocusOnClick(focusOnClick T.Gboolean) {
	FileChooserButtonSetFocusOnClick(f, focusOnClick)
}

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

	FileFilterAddCustom        func(filter *FileFilter, needed FileFilterFlags, f FileFilterFunc, dataGpointer, notify T.GDestroyNotify)
	FileFilterAddMimeType      func(filter *FileFilter, mimeType string)
	FileFilterAddPattern       func(filter *FileFilter, pattern string)
	FileFilterAddPixbufFormats func(filter *FileFilter)
	FileFilterFilter           func(filter *FileFilter, filterInfo *FileFilterInfo) T.Gboolean
	FileFilterGetName          func(filter *FileFilter) string
	FileFilterGetNeeded        func(filter *FileFilter) FileFilterFlags
	FileFilterSetName          func(filter *FileFilter, name string)
)

func (f *FileFilter) SetName(name string) {
	FileFilterSetName(f, name)
}

func (f *FileFilter) GetName() string {
	return FileFilterGetName(f)
}

func (f *FileFilter) AddMimeType(mimeType string) {
	FileFilterAddMimeType(f, mimeType)
}

func (f *FileFilter) AddPattern(pattern string) {
	FileFilterAddPattern(f, pattern)
}

func (f *FileFilter) AddPixbufFormats() {
	FileFilterAddPixbufFormats(f)
}

func (f *FileFilter) AddCustom(needed FileFilterFlags, fnc FileFilterFunc, dataGpointer, notify T.GDestroyNotify) {
	FileFilterAddCustom(f, needed, fnc, dataGpointer, notify)
}

func (f *FileFilter) GetNeeded() FileFilterFlags {
	return FileFilterGetNeeded(f)
}

func (f *FileFilter) Filter(filterInfo *FileFilterInfo) T.Gboolean {
	return FileFilterFilter(f, filterInfo)
}

type FileSelection struct {
	Parent_instance  T.GtkDialog
	Dir_list         *T.GtkWidget
	File_list        *T.GtkWidget
	Selection_entry  *T.GtkWidget
	Selection_text   *T.GtkWidget
	Main_vbox        *T.GtkWidget
	Ok_button        *T.GtkWidget
	Cancel_button    *T.GtkWidget
	Help_button      *T.GtkWidget
	History_pulldown *T.GtkWidget
	History_menu     *T.GtkWidget
	History_list     *T.GList
	Fileop_dialog    *T.GtkWidget
	Fileop_entry     *T.GtkWidget
	Fileop_file      *T.Gchar
	Cmpl_state       T.Gpointer
	Fileop_c_dir     *T.GtkWidget
	Fileop_del_file  *T.GtkWidget
	Fileop_ren_file  *T.GtkWidget
	Button_area      *T.GtkWidget
	Action_area      *T.GtkWidget
	Selected_names   *T.GPtrArray
	Last_selected    *T.Gchar
}

var (
	FileSelectionGetType func() T.GType
	FileSelectionNew     func(title string) *T.GtkWidget

	FileSelectionComplete          func(f *FileSelection, pattern string)
	FileSelectionGetFilename       func(f *FileSelection) string
	FileSelectionGetSelections     func(f *FileSelection) **T.Gchar
	FileSelectionGetSelectMultiple func(f *FileSelection) T.Gboolean
	FileSelectionHideFileopButtons func(f *FileSelection)
	FileSelectionSetFilename       func(f *FileSelection, filename string)
	FileSelectionSetSelectMultiple func(f *FileSelection, selectMultiple T.Gboolean)
	FileSelectionShowFileopButtons func(f *FileSelection)
)

func (f *FileSelection) SetFilename(filename string) {
	FileSelectionSetFilename(f, filename)
}

func (f *FileSelection) GetFilename() string {
	return FileSelectionGetFilename(f)
}

func (f *FileSelection) Complete(pattern string) {
	FileSelectionComplete(f, pattern)
}

func (f *FileSelection) ShowFileopButtons() {
	FileSelectionShowFileopButtons(f)
}

func (f *FileSelection) HideFileopButtons() {
	FileSelectionHideFileopButtons(f)
}

func (f *FileSelection) GetSelections() **T.Gchar {
	return FileSelectionGetSelections(f)
}

func (f *FileSelection) SetSelectMultiple(selectMultiple T.Gboolean) {
	FileSelectionSetSelectMultiple(f, selectMultiple)
}

func (f *FileSelection) GetSelectMultiple() T.Gboolean {
	return FileSelectionGetSelectMultiple(f)
}

type Fixed struct {
	Container T.GtkContainer
	Children  *T.GList
}

var (
	FixedGetType func() T.GType
	FixedNew     func() *T.GtkWidget

	FixedGetHasWindow func(f *Fixed) T.Gboolean
	FixedMove         func(f *Fixed, widget *T.GtkWidget, x, y int)
	FixedPut          func(f *Fixed, widget *T.GtkWidget, x, y int)
	FixedSetHasWindow func(f *Fixed, hasWindow T.Gboolean)
)

func (f *Fixed) Put(widget *T.GtkWidget, x, y int) {
	FixedPut(f, widget, x, y)
}

func (f *Fixed) Move(widget *T.GtkWidget, x, y int) {
	FixedMove(f, widget, x, y)
}

func (f *Fixed) SetHasWindow(hasWindow T.Gboolean) {
	FixedSetHasWindow(f, hasWindow)
}

func (f *Fixed) GetHasWindow() T.Gboolean {
	return FixedGetHasWindow(f)
}

type FontButton struct {
	Button Button
	_      *struct{}
}

var (
	FontButtonGetType     func() T.GType
	FontButtonNew         func() *T.GtkWidget
	FontButtonNewWithFont func(fontname string) *T.GtkWidget

	FontButtonGetFontName  func(f *FontButton) string
	FontButtonGetShowSize  func(f *FontButton) T.Gboolean
	FontButtonGetShowStyle func(f *FontButton) T.Gboolean
	FontButtonGetTitle     func(f *FontButton) string
	FontButtonGetUseFont   func(f *FontButton) T.Gboolean
	FontButtonGetUseSize   func(f *FontButton) T.Gboolean
	FontButtonSetFontName  func(f *FontButton, fontname string) T.Gboolean
	FontButtonSetShowSize  func(f *FontButton, showSize T.Gboolean)
	FontButtonSetShowStyle func(f *FontButton, showStyle T.Gboolean)
	FontButtonSetTitle     func(f *FontButton, title string)
	FontButtonSetUseFont   func(f *FontButton, useFont T.Gboolean)
	FontButtonSetUseSize   func(f *FontButton, useSize T.Gboolean)
)

func (f *FontButton) GetTitle() string {
	return FontButtonGetTitle(f)
}

func (f *FontButton) SetTitle(title string) {
	FontButtonSetTitle(f, title)
}

func (f *FontButton) GetUseFont() T.Gboolean {
	return FontButtonGetUseFont(f)
}

func (f *FontButton) SetUseFont(useFont T.Gboolean) {
	FontButtonSetUseFont(f, useFont)
}

func (f *FontButton) GetUseSize() T.Gboolean {
	return FontButtonGetUseSize(f)
}

func (f *FontButton) SetUseSize(useSize T.Gboolean) {
	FontButtonSetUseSize(f, useSize)
}

func (f *FontButton) GetFontName() string {
	return FontButtonGetFontName(f)
}

func (f *FontButton) SetFontName(fontname string) T.Gboolean {
	return FontButtonSetFontName(f, fontname)
}

func (f *FontButton) GetShowStyle() T.Gboolean {
	return FontButtonGetShowStyle(f)
}

func (f *FontButton) SetShowStyle(showStyle T.Gboolean) {
	FontButtonSetShowStyle(f, showStyle)
}

func (f *FontButton) GetShowSize() T.Gboolean {
	return FontButtonGetShowSize(f)
}

func (f *FontButton) SetShowSize(showSize T.Gboolean) {
	FontButtonSetShowSize(f, showSize)
}

type FontSelection struct {
	Parent_instance  T.GtkVBox
	Font_entry       *T.GtkWidget
	Family_list      *T.GtkWidget
	Font_style_entry *T.GtkWidget
	Face_list        *T.GtkWidget
	Size_entry       *T.GtkWidget
	Size_list        *T.GtkWidget
	Pixels_button    *T.GtkWidget
	Points_button    *T.GtkWidget
	Filter_button    *T.GtkWidget
	Preview_entry    *T.GtkWidget
	Family           *T.PangoFontFamily
	Face             *T.PangoFontFace
	Size             int
	Font             *T.GdkFont
}

var (
	FontSelectionGetType func() T.GType
	FontSelectionNew     func() *T.GtkWidget

	FontSelectionGetFace         func(f *FontSelection) *T.PangoFontFace
	FontSelectionGetFaceList     func(f *FontSelection) *T.GtkWidget
	FontSelectionGetFamily       func(f *FontSelection) *T.PangoFontFamily
	FontSelectionGetFamilyList   func(f *FontSelection) *T.GtkWidget
	FontSelectionGetFont         func(f *FontSelection) *T.GdkFont
	FontSelectionGetFontName     func(f *FontSelection) string
	FontSelectionGetPreviewEntry func(f *FontSelection) *T.GtkWidget
	FontSelectionGetPreviewText  func(f *FontSelection) string
	FontSelectionGetSize         func(f *FontSelection) int
	FontSelectionGetSizeEntry    func(f *FontSelection) *T.GtkWidget
	FontSelectionGetSizeList     func(f *FontSelection) *T.GtkWidget
	FontSelectionSetFontName     func(f *FontSelection, fontname string) T.Gboolean
	FontSelectionSetPreviewText  func(f *FontSelection, text string)
)

func (f *FontSelection) GetFamilyList() *T.GtkWidget {
	return FontSelectionGetFamilyList(f)
}

func (f *FontSelection) GetFaceList() *T.GtkWidget {
	return FontSelectionGetFaceList(f)
}

func (f *FontSelection) GetSizeEntry() *T.GtkWidget {
	return FontSelectionGetSizeEntry(f)
}

func (f *FontSelection) GetSizeList() *T.GtkWidget {
	return FontSelectionGetSizeList(f)
}

func (f *FontSelection) GetPreviewEntry() *T.GtkWidget {
	return FontSelectionGetPreviewEntry(f)
}

func (f *FontSelection) GetFamily() *T.PangoFontFamily {
	return FontSelectionGetFamily(f)
}

func (f *FontSelection) GetFace() *T.PangoFontFace {
	return FontSelectionGetFace(f)
}

func (f *FontSelection) GetSize() int {
	return FontSelectionGetSize(f)
}

func (f *FontSelection) GetFontName() string {
	return FontSelectionGetFontName(f)
}

func (f *FontSelection) GetFont() *T.GdkFont {
	return FontSelectionGetFont(f)
}

func (f *FontSelection) SetFontName(fontname string) T.Gboolean {
	return FontSelectionSetFontName(f, fontname)
}

func (f *FontSelection) GetPreviewText() string {
	return FontSelectionGetPreviewText(f)
}

func (f *FontSelection) SetPreviewText(text string) {
	FontSelectionSetPreviewText(f, text)
}

type FontSelectionDialog struct {
	Parent       T.GtkDialog
	Fontsel      *T.GtkWidget
	MainVbox     *T.GtkWidget
	ActionArea   *T.GtkWidget
	OkButton     *T.GtkWidget
	ApplyButton  *T.GtkWidget
	CancelButton *T.GtkWidget
	DialogWidth  int
	AutoResize   T.Gboolean
}

var (
	FontSelectionDialogGetType func() T.GType
	FontSelectionDialogNew     func(title string) *T.GtkWidget

	FontSelectionDialogGetApplyButton   func(f *FontSelectionDialog) *T.GtkWidget
	FontSelectionDialogGetCancelButton  func(f *FontSelectionDialog) *T.GtkWidget
	FontSelectionDialogGetFont          func(f *FontSelectionDialog) *T.GdkFont
	FontSelectionDialogGetFontName      func(f *FontSelectionDialog) string
	FontSelectionDialogGetFontSelection func(f *FontSelectionDialog) *T.GtkWidget
	FontSelectionDialogGetOkButton      func(f *FontSelectionDialog) *T.GtkWidget
	FontSelectionDialogGetPreviewText   func(f *FontSelectionDialog) string
	FontSelectionDialogSetFontName      func(f *FontSelectionDialog, fontname string) T.Gboolean
	FontSelectionDialogSetPreviewText   func(f *FontSelectionDialog, text string)
)

func (f *FontSelectionDialog) GetOkButton() *T.GtkWidget {
	return FontSelectionDialogGetOkButton(f)
}

func (f *FontSelectionDialog) GetApplyButton() *T.GtkWidget {
	return FontSelectionDialogGetApplyButton(f)
}

func (f *FontSelectionDialog) GetCancelButton() *T.GtkWidget {
	return FontSelectionDialogGetCancelButton(f)
}

func (f *FontSelectionDialog) GetFontSelection() *T.GtkWidget {
	return FontSelectionDialogGetFontSelection(f)
}

func (f *FontSelectionDialog) GetFontName() string {
	return FontSelectionDialogGetFontName(f)
}

func (f *FontSelectionDialog) GetFont() *T.GdkFont {
	return FontSelectionDialogGetFont(f)
}

func (f *FontSelectionDialog) SetFontName(fontname string) T.Gboolean {
	return FontSelectionDialogSetFontName(f, fontname)
}

func (f *FontSelectionDialog) GetPreviewText() string {
	return FontSelectionDialogGetPreviewText(f)
}

func (f *FontSelectionDialog) SetPreviewText(text string) {
	FontSelectionDialogSetPreviewText(f, text)
}

type Frame struct {
	Bin             T.GtkBin
	LabelWidget     *T.GtkWidget
	ShadowType      int16
	LabelXalign     float32
	LabelYalign     float32
	ChildAllocation T.GtkAllocation
}

var (
	FrameGetType func() T.GType
	FrameNew     func(label string) *T.GtkWidget

	FrameGetLabel       func(f *Frame) string
	FrameGetLabelAlign  func(f *Frame, xalign, yalign *float32)
	FrameGetLabelWidget func(f *Frame) *T.GtkWidget
	FrameGetShadowType  func(f *Frame) T.GtkShadowType
	FrameSetLabel       func(f *Frame, label string)
	FrameSetLabelAlign  func(f *Frame, xalign, yalign float32)
	FrameSetLabelWidget func(f *Frame, labelWidget *T.GtkWidget)
	FrameSetShadowType  func(f *Frame, t T.GtkShadowType)
)

func (f *Frame) SetLabel(label string) {
	FrameSetLabel(f, label)
}

func (f *Frame) GetLabel() string {
	return FrameGetLabel(f)
}

func (f *Frame) SetLabelWidget(labelWidget *T.GtkWidget) {
	FrameSetLabelWidget(f, labelWidget)
}

func (f *Frame) GetLabelWidget() *T.GtkWidget {
	return FrameGetLabelWidget(f)
}

func (f *Frame) SetLabelAlign(xalign, yalign float32) {
	FrameSetLabelAlign(f, xalign, yalign)
}

func (f *Frame) GetLabelAlign(xalign, yalign *float32) {
	FrameGetLabelAlign(f, xalign, yalign)
}

func (f *Frame) SetShadowType(t T.GtkShadowType) {
	FrameSetShadowType(f, t)
}

func (f *Frame) GetShadowType() T.GtkShadowType {
	return FrameGetShadowType(f)
}
