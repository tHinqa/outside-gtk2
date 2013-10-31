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

	ActionActivate         func(a *Action, parameter *T.GVariant)
	ActionGetEnabled       func(a *Action) bool
	ActionGetParameterType func(a *Action) *T.GVariantType
	ActionGetState         func(a *Action) *T.GVariant
	ActionGetStateHint     func(a *Action) *T.GVariant
	ActionGetStateType     func(a *Action) *T.GVariantType
	ActionSetState         func(a *Action, value *T.GVariant)
)

func (a *Action) Activate(parameter *T.GVariant)    { ActionActivate(a, parameter) }
func (a *Action) GetEnabled() bool                  { return ActionGetEnabled(a) }
func (a *Action) GetParameterType() *T.GVariantType { return ActionGetParameterType(a) }
func (a *Action) GetState() *T.GVariant             { return ActionGetState(a) }
func (a *Action) GetStateHint() *T.GVariant         { return ActionGetStateHint(a) }
func (a *Action) GetStateType() *T.GVariantType     { return ActionGetStateType(a) }
func (a *Action) SetState(value *T.GVariant)        { ActionSetState(a, value) }

type ActionGroup struct{}

var (
	ActionGroupGetType func() O.Type

	ActionGroupActionAdded            func(a *ActionGroup, actionName string)
	ActionGroupActionEnabledChanged   func(a *ActionGroup, actionName string, enabled bool)
	ActionGroupActionRemoved          func(a *ActionGroup, actionName string)
	ActionGroupActionStateChanged     func(a *ActionGroup, actionName string, state *T.GVariant)
	ActionGroupActivateAction         func(a *ActionGroup, actionName string, parameter *T.GVariant)
	ActionGroupChangeActionState      func(a *ActionGroup, actionName string, value *T.GVariant)
	ActionGroupGetActionEnabled       func(a *ActionGroup, actionName string) bool
	ActionGroupGetActionParameterType func(a *ActionGroup, actionName string) *T.GVariantType
	ActionGroupGetActionState         func(a *ActionGroup, actionName string) *T.GVariant
	ActionGroupGetActionStateHint     func(a *ActionGroup, actionName string) *T.GVariant
	ActionGroupGetActionStateType     func(a *ActionGroup, actionName string) *T.GVariantType
	ActionGroupHasAction              func(a *ActionGroup, actionName string) bool
	ActionGroupListActions            func(a *ActionGroup) []string
)

func (a *ActionGroup) ActionAdded(actionName string) { ActionGroupActionAdded(a, actionName) }
func (a *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	ActionGroupActionEnabledChanged(a, actionName, enabled)
}
func (a *ActionGroup) ActionRemoved(actionName string) { ActionGroupActionRemoved(a, actionName) }
func (a *ActionGroup) ActionStateChanged(actionName string, state *T.GVariant) {
	ActionGroupActionStateChanged(a, actionName, state)
}
func (a *ActionGroup) ActivateAction(actionName string, parameter *T.GVariant) {
	ActionGroupActivateAction(a, actionName, parameter)
}
func (a *ActionGroup) ChangeActionState(actionName string, value *T.GVariant) {
	ActionGroupChangeActionState(a, actionName, value)
}
func (a *ActionGroup) GetActionEnabled(actionName string) bool {
	return ActionGroupGetActionEnabled(a, actionName)
}
func (a *ActionGroup) GetActionParameterType(actionName string) *T.GVariantType {
	return ActionGroupGetActionParameterType(a, actionName)
}
func (a *ActionGroup) GetActionState(actionName string) *T.GVariant {
	return ActionGroupGetActionState(a, actionName)
}
func (a *ActionGroup) GetActionStateHint(actionName string) *T.GVariant {
	return ActionGroupGetActionStateHint(a, actionName)
}
func (a *ActionGroup) GetActionStateType(actionName string) *T.GVariantType {
	return ActionGroupGetActionStateType(a, actionName)
}
func (a *ActionGroup) HasAction(actionName string) bool {
	return ActionGroupHasAction(a, actionName)
}
func (a *ActionGroup) ListActions() []string { return ActionGroupListActions(a) }

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

	AppInfoAddSupportsType          func(a *AppInfo, contentType string, err **T.GError) bool
	AppInfoCanDelete                func(a *AppInfo) bool
	AppInfoCanRemoveSupportsType    func(a *AppInfo) bool
	AppInfoDelete                   func(a *AppInfo) bool
	AppInfoDup                      func(a *AppInfo) *AppInfo
	AppInfoEqual                    func(a *AppInfo, appinfo2 *AppInfo) bool
	AppInfoGetCommandline           func(a *AppInfo) string
	AppInfoGetDescription           func(a *AppInfo) string
	AppInfoGetDisplayName           func(a *AppInfo) string
	AppInfoGetExecutable            func(a *AppInfo) string
	AppInfoGetIcon                  func(a *AppInfo) *Icon
	AppInfoGetId                    func(a *AppInfo) string
	AppInfoGetName                  func(a *AppInfo) string
	AppInfoLaunch                   func(a *AppInfo, files *T.GList, launchContext *AppLaunchContext, err **T.GError) bool
	AppInfoLaunchUris               func(a *AppInfo, uris *T.GList, launchContext *AppLaunchContext, err **T.GError) bool
	AppInfoRemoveSupportsType       func(a *AppInfo, contentType string, err **T.GError) bool
	AppInfoSetAsDefaultForExtension func(a *AppInfo, extension string, err **T.GError) bool
	AppInfoSetAsDefaultForType      func(a *AppInfo, contentType string, err **T.GError) bool
	AppInfoSetAsLastUsedForType     func(a *AppInfo, contentType string, err **T.GError) bool
	AppInfoShouldShow               func(a *AppInfo) bool
	AppInfoSupportsFiles            func(a *AppInfo) bool
	AppInfoSupportsUris             func(a *AppInfo) bool
)

func (a *AppInfo) AddSupportsType(contentType string, err **T.GError) bool {
	return AppInfoAddSupportsType(a, contentType, err)
}
func (a *AppInfo) CanDelete() bool              { return AppInfoCanDelete(a) }
func (a *AppInfo) CanRemoveSupportsType() bool  { return AppInfoCanRemoveSupportsType(a) }
func (a *AppInfo) Delete() bool                 { return AppInfoDelete(a) }
func (a *AppInfo) Dup() *AppInfo                { return AppInfoDup(a) }
func (a *AppInfo) Equal(appinfo2 *AppInfo) bool { return AppInfoEqual(a, appinfo2) }
func (a *AppInfo) GetCommandline() string       { return AppInfoGetCommandline(a) }
func (a *AppInfo) GetDescription() string       { return AppInfoGetDescription(a) }
func (a *AppInfo) GetDisplayName() string       { return AppInfoGetDisplayName(a) }
func (a *AppInfo) GetExecutable() string        { return AppInfoGetExecutable(a) }
func (a *AppInfo) GetIcon() *Icon               { return AppInfoGetIcon(a) }
func (a *AppInfo) GetId() string                { return AppInfoGetId(a) }
func (a *AppInfo) GetName() string              { return AppInfoGetName(a) }
func (a *AppInfo) Launch(files *T.GList, launchContext *AppLaunchContext, err **T.GError) bool {
	return AppInfoLaunch(a, files, launchContext, err)
}
func (a *AppInfo) LaunchUris(uris *T.GList, launchContext *AppLaunchContext, err **T.GError) bool {
	return AppInfoLaunchUris(a, uris, launchContext, err)
}
func (a *AppInfo) RemoveSupportsType(contentType string, err **T.GError) bool {
	return AppInfoRemoveSupportsType(a, contentType, err)
}
func (a *AppInfo) SetAsDefaultForExtension(extension string, err **T.GError) bool {
	return AppInfoSetAsDefaultForExtension(a, extension, err)
}
func (a *AppInfo) SetAsDefaultForType(contentType string, err **T.GError) bool {
	return AppInfoSetAsDefaultForType(a, contentType, err)
}
func (a *AppInfo) SetAsLastUsedForType(contentType string, err **T.GError) bool {
	return AppInfoSetAsLastUsedForType(a, contentType, err)
}
func (a *AppInfo) ShouldShow() bool    { return AppInfoShouldShow(a) }
func (a *AppInfo) SupportsFiles() bool { return AppInfoSupportsFiles(a) }
func (a *AppInfo) SupportsUris() bool  { return AppInfoSupportsUris(a) }

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

	AppLaunchContextGetDisplay         func(a *AppLaunchContext, info *AppInfo, files *T.GList) string
	AppLaunchContextGetStartupNotifyId func(a *AppLaunchContext, info *AppInfo, files *T.GList) string
	AppLaunchContextLaunchFailed       func(a *AppLaunchContext, startupNotifyId string)
)

func (a *AppLaunchContext) GetDisplay(info *AppInfo, files *T.GList) string {
	return AppLaunchContextGetDisplay(a, info, files)
}
func (a *AppLaunchContext) GetStartupNotifyId(info *AppInfo, files *T.GList) string {
	return AppLaunchContextGetStartupNotifyId(a, info, files)
}
func (a *AppLaunchContext) LaunchFailed(startupNotifyId string) {
	AppLaunchContextLaunchFailed(a, startupNotifyId)
}

type Application struct {
	Parent O.Object
	_      *struct{}
}

var (
	ApplicationGetType func() O.Type

	ApplicationIdIsValid func(applicationId string) bool
	ApplicationNew       func(applicationId string, flags ApplicationFlags) *Application

	ApplicationActivate             func(a *Application)
	ApplicationGetApplicationId     func(a *Application) string
	ApplicationGetFlags             func(a *Application) ApplicationFlags
	ApplicationGetInactivityTimeout func(a *Application) uint
	ApplicationGetIsRegistered      func(a *Application) bool
	ApplicationGetIsRemote          func(a *Application) bool
	ApplicationHold                 func(a *Application)
	ApplicationOpen                 func(a *Application, files **File, nFiles int, hint string)
	ApplicationRegister             func(a *Application, cancellable *Cancellable, err **T.GError) bool
	ApplicationRelease              func(a *Application)
	ApplicationRun                  func(a *Application, argc int, argv **T.Char) int
	ApplicationSetActionGroup       func(a *Application, ag *ActionGroup)
	ApplicationSetApplicationId     func(a *Application, applicationId string)
	ApplicationSetFlags             func(a *Application, flags ApplicationFlags)
	ApplicationSetInactivityTimeout func(a *Application, inactivityTimeout uint)
)

func (a *Application) Activate()                  { ApplicationActivate(a) }
func (a *Application) GetApplicationId() string   { return ApplicationGetApplicationId(a) }
func (a *Application) GetFlags() ApplicationFlags { return ApplicationGetFlags(a) }
func (a *Application) GetInactivityTimeout() uint { return ApplicationGetInactivityTimeout(a) }
func (a *Application) GetIsRegistered() bool      { return ApplicationGetIsRegistered(a) }
func (a *Application) GetIsRemote() bool          { return ApplicationGetIsRemote(a) }
func (a *Application) Hold()                      { ApplicationHold(a) }
func (a *Application) Open(files **File, nFiles int, hint string) {
	ApplicationOpen(a, files, nFiles, hint)
}
func (a *Application) Register(cancellable *Cancellable, err **T.GError) bool {
	return ApplicationRegister(a, cancellable, err)
}
func (a *Application) Release()                        { ApplicationRelease(a) }
func (a *Application) Run(argc int, argv **T.Char) int { return ApplicationRun(a, argc, argv) }
func (a *Application) SetActionGroup(ag *ActionGroup)  { ApplicationSetActionGroup(a, ag) }
func (a *Application) SetApplicationId(applicationId string) {
	ApplicationSetApplicationId(a, applicationId)
}
func (a *Application) SetFlags(flags ApplicationFlags) { ApplicationSetFlags(a, flags) }
func (a *Application) SetInactivityTimeout(inactivityTimeout uint) {
	ApplicationSetInactivityTimeout(a, inactivityTimeout)
}

type ApplicationCommandLine struct {
	Parent O.Object
	_      *struct{}
}

var (
	ApplicationCommandLineGetType func() O.Type

	ApplicationCommandLineGetArguments    func(a *ApplicationCommandLine, argc *int) []string
	ApplicationCommandLineGetCwd          func(a *ApplicationCommandLine) string
	ApplicationCommandLineGetenv          func(a *ApplicationCommandLine, name string) string
	ApplicationCommandLineGetEnviron      func(a *ApplicationCommandLine) []string
	ApplicationCommandLineGetExitStatus   func(a *ApplicationCommandLine) int
	ApplicationCommandLineGetIsRemote     func(a *ApplicationCommandLine) bool
	ApplicationCommandLineGetPlatformData func(a *ApplicationCommandLine) *T.GVariant
	ApplicationCommandLinePrint           func(a *ApplicationCommandLine, format string, v ...VArg)
	ApplicationCommandLinePrinterr        func(a *ApplicationCommandLine, format string, v ...VArg)
	ApplicationCommandLineSetExitStatus   func(a *ApplicationCommandLine, exitStatus int)
)

func (a *ApplicationCommandLine) GetArguments(argc *int) []string {
	return ApplicationCommandLineGetArguments(a, argc)
}
func (a *ApplicationCommandLine) GetCwd() string { return ApplicationCommandLineGetCwd(a) }
func (a *ApplicationCommandLine) Getenv(name string) string {
	return ApplicationCommandLineGetenv(a, name)
}
func (a *ApplicationCommandLine) GetEnviron() []string { return ApplicationCommandLineGetEnviron(a) }
func (a *ApplicationCommandLine) GetExitStatus() int   { return ApplicationCommandLineGetExitStatus(a) }
func (a *ApplicationCommandLine) GetIsRemote() bool    { return ApplicationCommandLineGetIsRemote(a) }
func (a *ApplicationCommandLine) GetPlatformData() *T.GVariant {
	return ApplicationCommandLineGetPlatformData(a)
}
func (a *ApplicationCommandLine) Print(format string, v ...VArg) {
	ApplicationCommandLinePrint(a, format, v)
}
func (a *ApplicationCommandLine) Printerr(format string, v ...VArg) {
	ApplicationCommandLinePrinterr(a, format, v)
}
func (a *ApplicationCommandLine) SetExitStatus(exitStatus int) {
	ApplicationCommandLineSetExitStatus(a, exitStatus)
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

	AsyncInitableInitAsync  func(a *AsyncInitable, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	AsyncInitableInitFinish func(a *AsyncInitable, res *AsyncResult, err **T.GError) bool
	AsyncInitableNewFinish  func(a *AsyncInitable, res *AsyncResult, err **T.GError) *O.Object
)

func (a *AsyncInitable) InitAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	AsyncInitableInitAsync(a, ioPriority, cancellable, callback, userData)
}
func (a *AsyncInitable) InitFinish(res *AsyncResult, err **T.GError) bool {
	return AsyncInitableInitFinish(a, res, err)
}
func (a *AsyncInitable) NewFinish(res *AsyncResult, err **T.GError) *O.Object {
	return AsyncInitableNewFinish(a, res, err)
}

type AsyncReadyCallback func(
	sourceObject *O.Object,
	res *AsyncResult,
	userData T.Gpointer)

type AsyncResult struct{}

var (
	AsyncResultGetType func() O.Type

	AsyncResultGetSourceObject func(a *AsyncResult) *O.Object
	AsyncResultGetUserData     func(a *AsyncResult) T.Gpointer
)

func (a *AsyncResult) GetSourceObject() *O.Object { return AsyncResultGetSourceObject(a) }
func (a *AsyncResult) GetUserData() T.Gpointer    { return AsyncResultGetUserData(a) }
