// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Action struct{}

var (
	ActionGetType func() O.Type

	ActionGetName func(a *Action) string

	actionActivate         func(a *Action, parameter *T.GVariant)
	actionGetEnabled       func(a *Action) bool
	actionGetParameterType func(a *Action) *T.GVariantType
	actionGetState         func(a *Action) *T.GVariant
	actionGetStateHint     func(a *Action) *T.GVariant
	actionGetStateType     func(a *Action) *T.GVariantType
	actionSetState         func(a *Action, value *T.GVariant)
)

func (a *Action) Activate(parameter *T.GVariant)    { actionActivate(a, parameter) }
func (a *Action) GetEnabled() bool                  { return actionGetEnabled(a) }
func (a *Action) GetParameterType() *T.GVariantType { return actionGetParameterType(a) }
func (a *Action) GetState() *T.GVariant             { return actionGetState(a) }
func (a *Action) GetStateHint() *T.GVariant         { return actionGetStateHint(a) }
func (a *Action) GetStateType() *T.GVariantType     { return actionGetStateType(a) }
func (a *Action) SetState(value *T.GVariant)        { actionSetState(a, value) }

type ActionGroup struct{}

var (
	ActionGroupGetType func() O.Type

	actionGroupActionAdded            func(a *ActionGroup, actionName string)
	actionGroupActionEnabledChanged   func(a *ActionGroup, actionName string, enabled bool)
	actionGroupActionRemoved          func(a *ActionGroup, actionName string)
	actionGroupActionStateChanged     func(a *ActionGroup, actionName string, state *T.GVariant)
	actionGroupActivateAction         func(a *ActionGroup, actionName string, parameter *T.GVariant)
	actionGroupChangeActionState      func(a *ActionGroup, actionName string, value *T.GVariant)
	actionGroupGetActionEnabled       func(a *ActionGroup, actionName string) bool
	actionGroupGetActionParameterType func(a *ActionGroup, actionName string) *T.GVariantType
	actionGroupGetActionState         func(a *ActionGroup, actionName string) *T.GVariant
	actionGroupGetActionStateHint     func(a *ActionGroup, actionName string) *T.GVariant
	actionGroupGetActionStateType     func(a *ActionGroup, actionName string) *T.GVariantType
	actionGroupHasAction              func(a *ActionGroup, actionName string) bool
	actionGroupListActions            func(a *ActionGroup) []string
)

func (a *ActionGroup) ActionAdded(actionName string) { actionGroupActionAdded(a, actionName) }
func (a *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	actionGroupActionEnabledChanged(a, actionName, enabled)
}
func (a *ActionGroup) ActionRemoved(actionName string) { actionGroupActionRemoved(a, actionName) }
func (a *ActionGroup) ActionStateChanged(actionName string, state *T.GVariant) {
	actionGroupActionStateChanged(a, actionName, state)
}
func (a *ActionGroup) ActivateAction(actionName string, parameter *T.GVariant) {
	actionGroupActivateAction(a, actionName, parameter)
}
func (a *ActionGroup) ChangeActionState(actionName string, value *T.GVariant) {
	actionGroupChangeActionState(a, actionName, value)
}
func (a *ActionGroup) GetActionEnabled(actionName string) bool {
	return actionGroupGetActionEnabled(a, actionName)
}
func (a *ActionGroup) GetActionParameterType(actionName string) *T.GVariantType {
	return actionGroupGetActionParameterType(a, actionName)
}
func (a *ActionGroup) GetActionState(actionName string) *T.GVariant {
	return actionGroupGetActionState(a, actionName)
}
func (a *ActionGroup) GetActionStateHint(actionName string) *T.GVariant {
	return actionGroupGetActionStateHint(a, actionName)
}
func (a *ActionGroup) GetActionStateType(actionName string) *T.GVariantType {
	return actionGroupGetActionStateType(a, actionName)
}
func (a *ActionGroup) HasAction(actionName string) bool {
	return actionGroupHasAction(a, actionName)
}
func (a *ActionGroup) ListActions() []string { return actionGroupListActions(a) }

type AppInfo struct{}

var (
	AppInfoGetType func() O.Type

	AppInfoGetAll                 func() *T.GList
	AppInfoGetAllForType          func(contentType string) *T.GList
	AppInfoGetRecommendedForType  func(contentType string) *T.GList
	AppInfoGetFallbackForType     func(contentType string) *T.GList
	AppInfoResetTypeAssociations  func(contentType string)
	AppInfoGetDefaultForType      func(contentType string, mustSupportUris bool) *AppInfo
	AppInfoGetDefaultForUriScheme func(uriScheme string) *AppInfo
	AppInfoLaunchDefaultForUri    func(uri string, launchContext *AppLaunchContext, err **T.GError) bool
	AppInfoCreateFromCommandline  func(commandline string, applicationName string, flags AppInfoCreateFlags, err **T.GError) *AppInfo

	appInfoAddSupportsType          func(a *AppInfo, contentType string, err **T.GError) bool
	appInfoCanDelete                func(a *AppInfo) bool
	appInfoCanRemoveSupportsType    func(a *AppInfo) bool
	appInfoDelete                   func(a *AppInfo) bool
	appInfoDup                      func(a *AppInfo) *AppInfo
	appInfoEqual                    func(a *AppInfo, appinfo2 *AppInfo) bool
	appInfoGetCommandline           func(a *AppInfo) string
	appInfoGetDescription           func(a *AppInfo) string
	appInfoGetDisplayName           func(a *AppInfo) string
	appInfoGetExecutable            func(a *AppInfo) string
	appInfoGetIcon                  func(a *AppInfo) *Icon
	appInfoGetId                    func(a *AppInfo) string
	appInfoGetName                  func(a *AppInfo) string
	appInfoLaunch                   func(a *AppInfo, files *T.GList, launchContext *AppLaunchContext, err **T.GError) bool
	appInfoLaunchUris               func(a *AppInfo, uris *T.GList, launchContext *AppLaunchContext, err **T.GError) bool
	appInfoRemoveSupportsType       func(a *AppInfo, contentType string, err **T.GError) bool
	appInfoSetAsDefaultForExtension func(a *AppInfo, extension string, err **T.GError) bool
	appInfoSetAsDefaultForType      func(a *AppInfo, contentType string, err **T.GError) bool
	appInfoSetAsLastUsedForType     func(a *AppInfo, contentType string, err **T.GError) bool
	appInfoShouldShow               func(a *AppInfo) bool
	appInfoSupportsFiles            func(a *AppInfo) bool
	appInfoSupportsUris             func(a *AppInfo) bool
)

func (a *AppInfo) AddSupportsType(contentType string, err **T.GError) bool {
	return appInfoAddSupportsType(a, contentType, err)
}
func (a *AppInfo) CanDelete() bool              { return appInfoCanDelete(a) }
func (a *AppInfo) CanRemoveSupportsType() bool  { return appInfoCanRemoveSupportsType(a) }
func (a *AppInfo) Delete() bool                 { return appInfoDelete(a) }
func (a *AppInfo) Dup() *AppInfo                { return appInfoDup(a) }
func (a *AppInfo) Equal(appinfo2 *AppInfo) bool { return appInfoEqual(a, appinfo2) }
func (a *AppInfo) GetCommandline() string       { return appInfoGetCommandline(a) }
func (a *AppInfo) GetDescription() string       { return appInfoGetDescription(a) }
func (a *AppInfo) GetDisplayName() string       { return appInfoGetDisplayName(a) }
func (a *AppInfo) GetExecutable() string        { return appInfoGetExecutable(a) }
func (a *AppInfo) GetIcon() *Icon               { return appInfoGetIcon(a) }
func (a *AppInfo) GetId() string                { return appInfoGetId(a) }
func (a *AppInfo) GetName() string              { return appInfoGetName(a) }
func (a *AppInfo) Launch(files *T.GList, launchContext *AppLaunchContext, err **T.GError) bool {
	return appInfoLaunch(a, files, launchContext, err)
}
func (a *AppInfo) LaunchUris(uris *T.GList, launchContext *AppLaunchContext, err **T.GError) bool {
	return appInfoLaunchUris(a, uris, launchContext, err)
}
func (a *AppInfo) RemoveSupportsType(contentType string, err **T.GError) bool {
	return appInfoRemoveSupportsType(a, contentType, err)
}
func (a *AppInfo) SetAsDefaultForExtension(extension string, err **T.GError) bool {
	return appInfoSetAsDefaultForExtension(a, extension, err)
}
func (a *AppInfo) SetAsDefaultForType(contentType string, err **T.GError) bool {
	return appInfoSetAsDefaultForType(a, contentType, err)
}
func (a *AppInfo) SetAsLastUsedForType(contentType string, err **T.GError) bool {
	return appInfoSetAsLastUsedForType(a, contentType, err)
}
func (a *AppInfo) ShouldShow() bool    { return appInfoShouldShow(a) }
func (a *AppInfo) SupportsFiles() bool { return appInfoSupportsFiles(a) }
func (a *AppInfo) SupportsUris() bool  { return appInfoSupportsUris(a) }

type AppInfoCreateFlags Enum

const (
	APP_INFO_CREATE_NEEDS_TERMINAL AppInfoCreateFlags = 1 << iota
	APP_INFO_CREATE_SUPPORTS_URIS
	APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION
	APP_INFO_CREATE_NONE AppInfoCreateFlags = 0
)

var AppInfoCreateFlagsGetType func() O.Type

type AppLaunchContext struct {
	Parent O.Object
	_      *struct{}
}

var (
	AppLaunchContextGetType func() O.Type

	AppLaunchContextNew func() *AppLaunchContext

	appLaunchContextGetDisplay         func(a *AppLaunchContext, info *AppInfo, files *T.GList) string
	appLaunchContextGetStartupNotifyId func(a *AppLaunchContext, info *AppInfo, files *T.GList) string
	appLaunchContextLaunchFailed       func(a *AppLaunchContext, startupNotifyId string)
)

func (a *AppLaunchContext) GetDisplay(info *AppInfo, files *T.GList) string {
	return appLaunchContextGetDisplay(a, info, files)
}
func (a *AppLaunchContext) GetStartupNotifyId(info *AppInfo, files *T.GList) string {
	return appLaunchContextGetStartupNotifyId(a, info, files)
}
func (a *AppLaunchContext) LaunchFailed(startupNotifyId string) {
	appLaunchContextLaunchFailed(a, startupNotifyId)
}

type Application struct {
	Parent O.Object
	_      *struct{}
}

var (
	ApplicationGetType func() O.Type

	ApplicationIdIsValid func(applicationId string) bool
	ApplicationNew       func(applicationId string, flags ApplicationFlags) *Application

	applicationActivate             func(a *Application)
	applicationGetApplicationId     func(a *Application) string
	applicationGetFlags             func(a *Application) ApplicationFlags
	applicationGetInactivityTimeout func(a *Application) uint
	applicationGetIsRegistered      func(a *Application) bool
	applicationGetIsRemote          func(a *Application) bool
	applicationHold                 func(a *Application)
	applicationOpen                 func(a *Application, files **File, nFiles int, hint string)
	applicationRegister             func(a *Application, cancellable *Cancellable, err **T.GError) bool
	applicationRelease              func(a *Application)
	applicationRun                  func(a *Application, argc int, argv **T.Char) int
	applicationSetActionGroup       func(a *Application, ag *ActionGroup)
	applicationSetApplicationId     func(a *Application, applicationId string)
	applicationSetFlags             func(a *Application, flags ApplicationFlags)
	applicationSetInactivityTimeout func(a *Application, inactivityTimeout uint)
)

func (a *Application) Activate()                  { applicationActivate(a) }
func (a *Application) GetApplicationId() string   { return applicationGetApplicationId(a) }
func (a *Application) GetFlags() ApplicationFlags { return applicationGetFlags(a) }
func (a *Application) GetInactivityTimeout() uint { return applicationGetInactivityTimeout(a) }
func (a *Application) GetIsRegistered() bool      { return applicationGetIsRegistered(a) }
func (a *Application) GetIsRemote() bool          { return applicationGetIsRemote(a) }
func (a *Application) Hold()                      { applicationHold(a) }
func (a *Application) Open(files **File, nFiles int, hint string) {
	applicationOpen(a, files, nFiles, hint)
}
func (a *Application) Register(cancellable *Cancellable, err **T.GError) bool {
	return applicationRegister(a, cancellable, err)
}
func (a *Application) Release()                        { applicationRelease(a) }
func (a *Application) Run(argc int, argv **T.Char) int { return applicationRun(a, argc, argv) }
func (a *Application) SetActionGroup(ag *ActionGroup)  { applicationSetActionGroup(a, ag) }
func (a *Application) SetApplicationId(applicationId string) {
	applicationSetApplicationId(a, applicationId)
}
func (a *Application) SetFlags(flags ApplicationFlags) { applicationSetFlags(a, flags) }
func (a *Application) SetInactivityTimeout(inactivityTimeout uint) {
	applicationSetInactivityTimeout(a, inactivityTimeout)
}

type ApplicationCommandLine struct {
	Parent O.Object
	_      *struct{}
}

var (
	ApplicationCommandLineGetType func() O.Type

	applicationCommandLineGetArguments    func(a *ApplicationCommandLine, argc *int) []string
	applicationCommandLineGetCwd          func(a *ApplicationCommandLine) string
	applicationCommandLineGetenv          func(a *ApplicationCommandLine, name string) string
	applicationCommandLineGetEnviron      func(a *ApplicationCommandLine) []string
	applicationCommandLineGetExitStatus   func(a *ApplicationCommandLine) int
	applicationCommandLineGetIsRemote     func(a *ApplicationCommandLine) bool
	applicationCommandLineGetPlatformData func(a *ApplicationCommandLine) *T.GVariant
	applicationCommandLinePrint           func(a *ApplicationCommandLine, format string, v ...VArg)
	applicationCommandLinePrinterr        func(a *ApplicationCommandLine, format string, v ...VArg)
	applicationCommandLineSetExitStatus   func(a *ApplicationCommandLine, exitStatus int)
)

func (a *ApplicationCommandLine) GetArguments(argc *int) []string {
	return applicationCommandLineGetArguments(a, argc)
}
func (a *ApplicationCommandLine) GetCwd() string { return applicationCommandLineGetCwd(a) }
func (a *ApplicationCommandLine) Getenv(name string) string {
	return applicationCommandLineGetenv(a, name)
}
func (a *ApplicationCommandLine) GetEnviron() []string { return applicationCommandLineGetEnviron(a) }
func (a *ApplicationCommandLine) GetExitStatus() int   { return applicationCommandLineGetExitStatus(a) }
func (a *ApplicationCommandLine) GetIsRemote() bool    { return applicationCommandLineGetIsRemote(a) }
func (a *ApplicationCommandLine) GetPlatformData() *T.GVariant {
	return applicationCommandLineGetPlatformData(a)
}
func (a *ApplicationCommandLine) Print(format string, v ...VArg) {
	applicationCommandLinePrint(a, format, v)
}
func (a *ApplicationCommandLine) Printerr(format string, v ...VArg) {
	applicationCommandLinePrinterr(a, format, v)
}
func (a *ApplicationCommandLine) SetExitStatus(exitStatus int) {
	applicationCommandLineSetExitStatus(a, exitStatus)
}

type ApplicationFlags Enum

const (
	APPLICATION_IS_SERVICE ApplicationFlags = 1 << iota
	APPLICATION_IS_LAUNCHER
	APPLICATION_HANDLES_OPEN
	APPLICATION_HANDLES_COMMAND_LINE
	APPLICATION_SEND_ENVIRONMENT
	APPLICATION_FLAGS_NONE ApplicationFlags = 0
)

var ApplicationFlagsGetType func() O.Type

var AskPasswordFlagsGetType func() O.Type

type AsyncInitable struct{}

var (
	AsyncInitableGetType func() O.Type

	AsyncInitableNewAsync       func(objectType O.Type, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer, firstPropertyName string, v ...VArg)
	AsyncInitableNewvAsync      func(objectType O.Type, nParameters uint, parameters *T.GParameter, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	AsyncInitableNewValistAsync func(objectType O.Type, firstPropertyName string, varArgs T.VaList, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)

	asyncInitableInitAsync  func(a *AsyncInitable, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	asyncInitableInitFinish func(a *AsyncInitable, res *AsyncResult, err **T.GError) bool
	asyncInitableNewFinish  func(a *AsyncInitable, res *AsyncResult, err **T.GError) *O.Object
)

func (a *AsyncInitable) InitAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	asyncInitableInitAsync(a, ioPriority, cancellable, callback, userData)
}
func (a *AsyncInitable) InitFinish(res *AsyncResult, err **T.GError) bool {
	return asyncInitableInitFinish(a, res, err)
}
func (a *AsyncInitable) NewFinish(res *AsyncResult, err **T.GError) *O.Object {
	return asyncInitableNewFinish(a, res, err)
}

type AsyncReadyCallback func(
	sourceObject *O.Object,
	res *AsyncResult,
	userData T.Gpointer)

type AsyncResult struct{}

var (
	AsyncResultGetType func() O.Type

	asyncResultGetSourceObject func(a *AsyncResult) *O.Object
	asyncResultGetUserData     func(a *AsyncResult) T.Gpointer
)

func (a *AsyncResult) GetSourceObject() *O.Object { return asyncResultGetSourceObject(a) }
func (a *AsyncResult) GetUserData() T.Gpointer    { return asyncResultGetUserData(a) }
