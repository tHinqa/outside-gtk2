// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Seekable struct{}

var (
	SeekableGetType func() O.Type

	seekableCanSeek     func(seekable *Seekable) T.Gboolean
	seekableCanTruncate func(seekable *Seekable) T.Gboolean
	seekableSeek        func(seekable *Seekable, offset T.Goffset, typ SeekType, cancellable *Cancellable, err **T.GError) T.Gboolean
	seekableTell        func(seekable *Seekable) T.Goffset
	seekableTruncate    func(seekable *Seekable, offset T.Goffset, cancellable *Cancellable, err **T.GError) T.Gboolean
)

func (s *Seekable) CanSeek() T.Gboolean     { return seekableCanSeek(s) }
func (s *Seekable) CanTruncate() T.Gboolean { return seekableCanTruncate(s) }
func (s *Seekable) Seek(offset T.Goffset, typ SeekType, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return seekableSeek(s, offset, typ, cancellable, err)
}
func (s *Seekable) Tell() T.Goffset { return seekableTell(s) }
func (s *Seekable) Truncate(offset T.Goffset, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return seekableTruncate(s, offset, cancellable, err)
}

type SeekType Enum

const (
	SEEK_CUR SeekType = iota
	SEEK_SET
	SEEK_END
)

type Settings struct {
	Parent O.Object
	_      *struct{}
}

var (
	SettingsGetType func() O.Type

	SettingsListRelocatableSchemas func() []string
	SettingsListSchemas            func() []string
	SettingsNew                    func(schema string) *Settings
	SettingsNewWithBackend         func(schema string, backend *SettingsBackend) *Settings
	SettingsNewWithBackendAndPath  func(schema string, backend *SettingsBackend, path string) *Settings
	SettingsNewWithPath            func(schema string, path string) *Settings
	SettingsSync                   func()
	SettingsUnbind                 func(object T.Gpointer, property string)

	settingsApply           func(s *Settings)
	settingsBind            func(s *Settings, key string, object T.Gpointer, property string, flags SettingsBindFlags)
	settingsBindWithMapping func(s *Settings, key string, object T.Gpointer, property string, flags SettingsBindFlags, getMapping SettingsBindGetMapping, setMapping SettingsBindSetMapping, userData T.Gpointer, destroy T.GDestroyNotify)
	settingsBindWritable    func(s *Settings, key string, object T.Gpointer, property string, inverted T.Gboolean)
	settingsDelay           func(s *Settings)
	settingsGet             func(s *Settings, key, format string, v ...VArg)
	settingsGetBoolean      func(s *Settings, key string) T.Gboolean
	settingsGetChild        func(s *Settings, name string) *Settings
	settingsGetDouble       func(s *Settings, key string) float64
	settingsGetEnum         func(s *Settings, key string) int
	settingsGetFlags        func(s *Settings, key string) uint
	settingsGetHasUnapplied func(s *Settings) T.Gboolean
	settingsGetInt          func(s *Settings, key string) int
	settingsGetMapped       func(s *Settings, key string, mapping SettingsGetMapping, userData T.Gpointer) T.Gpointer
	settingsGetRange        func(s *Settings, key string) *T.GVariant
	settingsGetString       func(s *Settings, key string) string
	settingsGetStrv         func(s *Settings, key string) []string
	settingsGetValue        func(s *Settings, key string) *T.GVariant
	settingsIsWritable      func(s *Settings, name string) T.Gboolean
	settingsListKeys        func(s *Settings) []string
	settingsListChildren    func(s *Settings) []string
	settingsRangeCheck      func(s *Settings, key string, value *T.GVariant) T.Gboolean
	settingsReset           func(s *Settings, key string)
	settingsRevert          func(s *Settings)
	settingsSet             func(s *Settings, key, format string, v ...VArg) T.Gboolean
	settingsSetBoolean      func(s *Settings, key string, value T.Gboolean) T.Gboolean
	settingsSetDouble       func(s *Settings, key string, value float64) T.Gboolean
	settingsSetEnum         func(s *Settings, key string, value int) T.Gboolean
	settingsSetFlags        func(s *Settings, key string, value uint) T.Gboolean
	settingsSetInt          func(s *Settings, key string, value int) T.Gboolean
	settingsSetString       func(s *Settings, key, value string) T.Gboolean
	settingsSetStrv         func(s *Settings, key string, value **T.Gchar) T.Gboolean
	settingsSetValue        func(s *Settings, key string, value *T.GVariant) T.Gboolean
)

func (s *Settings) Apply() { settingsApply(s) }
func (s *Settings) Bind(key string, object T.Gpointer, property string, flags SettingsBindFlags) {
	settingsBind(s, key, object, property, flags)
}
func (s *Settings) BindWithMapping(key string, object T.Gpointer, property string, flags SettingsBindFlags, getMapping SettingsBindGetMapping, setMapping SettingsBindSetMapping, userData T.Gpointer, destroy T.GDestroyNotify) {
	settingsBindWithMapping(s, key, object, property, flags, getMapping, setMapping, userData, destroy)
}
func (s *Settings) BindWritable(key string, object T.Gpointer, property string, inverted T.Gboolean) {
	settingsBindWritable(s, key, object, property, inverted)
}
func (s *Settings) Delay()                            { settingsDelay(s) }
func (s *Settings) Get(key, format string, v ...VArg) { settingsGet(s, key, format, v) }
func (s *Settings) GetBoolean(key string) T.Gboolean  { return settingsGetBoolean(s, key) }
func (s *Settings) GetChild(name string) *Settings    { return settingsGetChild(s, name) }
func (s *Settings) GetDouble(key string) float64      { return settingsGetDouble(s, key) }
func (s *Settings) GetEnum(key string) int            { return settingsGetEnum(s, key) }
func (s *Settings) GetFlags(key string) uint          { return settingsGetFlags(s, key) }
func (s *Settings) GetHasUnapplied() T.Gboolean       { return settingsGetHasUnapplied(s) }
func (s *Settings) GetInt(key string) int             { return settingsGetInt(s, key) }
func (s *Settings) GetMapped(key string, mapping SettingsGetMapping, userData T.Gpointer) T.Gpointer {
	return settingsGetMapped(s, key, mapping, userData)
}
func (s *Settings) GetRange(key string) *T.GVariant   { return settingsGetRange(s, key) }
func (s *Settings) GetString(key string) string       { return settingsGetString(s, key) }
func (s *Settings) GetStrv(key string) []string       { return settingsGetStrv(s, key) }
func (s *Settings) GetValue(key string) *T.GVariant   { return settingsGetValue(s, key) }
func (s *Settings) IsWritable(name string) T.Gboolean { return settingsIsWritable(s, name) }
func (s *Settings) ListKeys() []string                { return settingsListKeys(s) }
func (s *Settings) ListChildren() []string            { return settingsListChildren(s) }
func (s *Settings) RangeCheck(key string, value *T.GVariant) T.Gboolean {
	return settingsRangeCheck(s, key, value)
}
func (s *Settings) Reset(key string) { settingsReset(s, key) }
func (s *Settings) Revert()          { settingsRevert(s) }
func (s *Settings) Set(key, format string, v ...VArg) T.Gboolean {
	return settingsSet(s, key, format, v)
}
func (s *Settings) SetBoolean(key string, value T.Gboolean) T.Gboolean {
	return settingsSetBoolean(s, key, value)
}
func (s *Settings) SetDouble(key string, value float64) T.Gboolean {
	return settingsSetDouble(s, key, value)
}
func (s *Settings) SetEnum(key string, value int) T.Gboolean   { return settingsSetEnum(s, key, value) }
func (s *Settings) SetFlags(key string, value uint) T.Gboolean { return settingsSetFlags(s, key, value) }
func (s *Settings) SetInt(key string, value int) T.Gboolean    { return settingsSetInt(s, key, value) }
func (s *Settings) SetString(key, value string) T.Gboolean     { return settingsSetString(s, key, value) }
func (s *Settings) SetStrv(key string, value **T.Gchar) T.Gboolean {
	return settingsSetStrv(s, key, value)
}
func (s *Settings) SetValue(key string, value *T.GVariant) T.Gboolean {
	return settingsSetValue(s, key, value)
}

type SettingsBackend struct{}

var (
	SettingsBackendGetType func() O.Type

	SettingsBackendFlattenTree func(tree *T.GTree, path **T.Gchar, keys ***T.Gchar, values ***T.GVariant)
	SettingsBackendGetDefault  func() *SettingsBackend

	settingsBackendKeysChanged         func(s *SettingsBackend, path string, items **T.Gchar, originTag T.Gpointer)
	settingsBackendChanged             func(s *SettingsBackend, key string, originTag T.Gpointer)
	settingsBackendChangedTree         func(s *SettingsBackend, tree *T.GTree, originTag T.Gpointer)
	settingsBackendPathChanged         func(s *SettingsBackend, path string, originTag T.Gpointer)
	settingsBackendPathWritableChanged func(s *SettingsBackend, path string)
	settingsBackendWritableChanged     func(s *SettingsBackend, key string)
)

func (s *SettingsBackend) KeysChanged(path string, items **T.Gchar, originTag T.Gpointer) {
	settingsBackendKeysChanged(s, path, items, originTag)
}
func (s *SettingsBackend) Changed(key string, originTag T.Gpointer) {
	settingsBackendChanged(s, key, originTag)
}
func (s *SettingsBackend) ChangedTree(tree *T.GTree, originTag T.Gpointer) {
	settingsBackendChangedTree(s, tree, originTag)
}
func (s *SettingsBackend) PathChanged(path string, originTag T.Gpointer) {
	settingsBackendPathChanged(s, path, originTag)
}
func (s *SettingsBackend) PathWritableChanged(path string) {
	settingsBackendPathWritableChanged(s, path)
}
func (s *SettingsBackend) WritableChanged(key string) { settingsBackendWritableChanged(s, key) }

type SettingsBindFlags Enum

const (
	SETTINGS_BIND_GET SettingsBindFlags = 1 << iota
	SETTINGS_BIND_SET
	SETTINGS_BIND_NO_SENSITIVITY
	SETTINGS_BIND_GET_NO_CHANGES
	SETTINGS_BIND_INVERT_BOOLEAN
	SETTINGS_BIND_DEFAULT SettingsBindFlags = 0
)

var SettingsBindFlagsGetType func() O.Type

type SettingsBindGetMapping func(
	value *O.Value,
	variant *T.GVariant,
	userData T.Gpointer) T.Gboolean

type SettingsBindSetMapping func(
	value *O.Value,
	expectedType *T.GVariantType,
	userData T.Gpointer) *T.GVariant

type SettingsGetMapping func(
	value *T.GVariant,
	result *T.Gpointer,
	userData T.Gpointer) T.Gboolean

type SimpleAction struct {
	Parent O.Object
	_      *struct{}
}

var (
	SimpleActionGetType     func() O.Type
	SimpleActionNew         func(name string, parameterType *T.GVariantType) *SimpleAction
	SimpleActionNewStateful func(name string, parameterType *T.GVariantType, state *T.GVariant) *SimpleAction

	simpleActionSetEnabled func(s *SimpleAction, enabled T.Gboolean)
)

func (s *SimpleAction) SetEnabled(enabled T.Gboolean) { simpleActionSetEnabled(s, enabled) }

type SimpleActionGroup struct {
	Parent O.Object
	_      *struct{}
}

var (
	SimpleActionGroupGetType func() O.Type
	SimpleActionGroupNew     func() *SimpleActionGroup

	simpleActionGroupInsert func(s *SimpleActionGroup, action *Action)
	simpleActionGroupLookup func(s *SimpleActionGroup, actionName string) *Action
	simpleActionGroupRemove func(s *SimpleActionGroup, actionName string)
)

func (s *SimpleActionGroup) Insert(action *Action) { simpleActionGroupInsert(s, action) }
func (s *SimpleActionGroup) Lookup(actionName string) *Action {
	return simpleActionGroupLookup(s, actionName)
}
func (s *SimpleActionGroup) Remove(actionName string) { simpleActionGroupRemove(s, actionName) }

var (
	SimpleAsyncReportErrorInIdle      func(object *O.Object, callback AsyncReadyCallback, userData T.Gpointer, domain T.GQuark, code int, format string, v ...VArg)
	SimpleAsyncReportGerrorInIdle     func(object *O.Object, callback AsyncReadyCallback, userData T.Gpointer, err *T.GError)
	SimpleAsyncReportTakeGerrorInIdle func(object *O.Object, callback AsyncReadyCallback, userData T.Gpointer, err *T.GError)
)

type SimpleAsyncResult struct{}

var (
	SimpleAsyncResultGetType      func() O.Type
	SimpleAsyncResultNew          func(sourceObject *O.Object, callback AsyncReadyCallback, userData T.Gpointer, sourceTag T.Gpointer) *SimpleAsyncResult
	SimpleAsyncResultNewError     func(sourceObject *O.Object, callback AsyncReadyCallback, userData T.Gpointer, domain T.GQuark, code int, format string, v ...VArg) *SimpleAsyncResult
	SimpleAsyncResultNewFromError func(sourceObject *O.Object, callback AsyncReadyCallback, userData T.Gpointer, err *T.GError) *SimpleAsyncResult
	SimpleAsyncResultNewTakeError func(sourceObject *O.Object, callback AsyncReadyCallback, userData T.Gpointer, err *T.GError) *SimpleAsyncResult

	SimpleAsyncResultIsValid func(result *AsyncResult, source *O.Object, sourceTag T.Gpointer) T.Gboolean

	simpleAsyncResultComplete              func(s *SimpleAsyncResult)
	simpleAsyncResultCompleteInIdle        func(s *SimpleAsyncResult)
	simpleAsyncResultGetOpResGboolean      func(s *SimpleAsyncResult) T.Gboolean
	simpleAsyncResultGetOpResGpointer      func(s *SimpleAsyncResult) T.Gpointer
	simpleAsyncResultGetOpResGssize        func(s *SimpleAsyncResult) T.Gssize
	simpleAsyncResultGetSourceTag          func(s *SimpleAsyncResult) T.Gpointer
	simpleAsyncResultPropagateError        func(s *SimpleAsyncResult, dest **T.GError) T.Gboolean
	simpleAsyncResultRunInThread           func(s *SimpleAsyncResult, f SimpleAsyncThreadFunc, ioPriority int, cancellable *Cancellable)
	simpleAsyncResultSetError              func(s *SimpleAsyncResult, domain T.GQuark, code int, format string, v ...VArg)
	simpleAsyncResultSetErrorVa            func(s *SimpleAsyncResult, domain T.GQuark, code int, format string, args T.VaList)
	simpleAsyncResultSetFromError          func(s *SimpleAsyncResult, err *T.GError)
	simpleAsyncResultSetHandleCancellation func(s *SimpleAsyncResult, handleCancellation T.Gboolean)
	simpleAsyncResultSetOpResGboolean      func(s *SimpleAsyncResult, opRes T.Gboolean)
	simpleAsyncResultSetOpResGpointer      func(s *SimpleAsyncResult, opRes T.Gpointer, destroyOpRes T.GDestroyNotify)
	simpleAsyncResultSetOpResGssize        func(s *SimpleAsyncResult, opRes T.Gssize)
	simpleAsyncResultTakeError             func(s *SimpleAsyncResult, err *T.GError)
)

func (s *SimpleAsyncResult) Complete()                    { simpleAsyncResultComplete(s) }
func (s *SimpleAsyncResult) CompleteInIdle()              { simpleAsyncResultCompleteInIdle(s) }
func (s *SimpleAsyncResult) GetOpResGboolean() T.Gboolean { return simpleAsyncResultGetOpResGboolean(s) }
func (s *SimpleAsyncResult) GetOpResGpointer() T.Gpointer { return simpleAsyncResultGetOpResGpointer(s) }
func (s *SimpleAsyncResult) GetOpResGssize() T.Gssize     { return simpleAsyncResultGetOpResGssize(s) }
func (s *SimpleAsyncResult) GetSourceTag() T.Gpointer     { return simpleAsyncResultGetSourceTag(s) }
func (s *SimpleAsyncResult) PropagateError(dest **T.GError) T.Gboolean {
	return simpleAsyncResultPropagateError(s, dest)
}
func (s *SimpleAsyncResult) RunInThread(f SimpleAsyncThreadFunc, ioPriority int, cancellable *Cancellable) {
	simpleAsyncResultRunInThread(s, f, ioPriority, cancellable)
}
func (s *SimpleAsyncResult) SetError(domain T.GQuark, code int, format string, v ...VArg) {
	simpleAsyncResultSetError(s, domain, code, format, v)
}
func (s *SimpleAsyncResult) SetErrorVa(domain T.GQuark, code int, format string, args T.VaList) {
	simpleAsyncResultSetErrorVa(s, domain, code, format, args)
}
func (s *SimpleAsyncResult) SetFromError(err *T.GError) { simpleAsyncResultSetFromError(s, err) }
func (s *SimpleAsyncResult) SetHandleCancellation(handleCancellation T.Gboolean) {
	simpleAsyncResultSetHandleCancellation(s, handleCancellation)
}
func (s *SimpleAsyncResult) SetOpResGboolean(opRes T.Gboolean) {
	simpleAsyncResultSetOpResGboolean(s, opRes)
}
func (s *SimpleAsyncResult) SetOpResGpointer(opRes T.Gpointer, destroyOpRes T.GDestroyNotify) {
	simpleAsyncResultSetOpResGpointer(s, opRes, destroyOpRes)
}
func (s *SimpleAsyncResult) SetOpResGssize(opRes T.Gssize) { simpleAsyncResultSetOpResGssize(s, opRes) }
func (s *SimpleAsyncResult) TakeError(err *T.GError)       { simpleAsyncResultTakeError(s, err) }

type SimpleAsyncThreadFunc func(
	res *SimpleAsyncResult,
	object *O.Object,
	cancellable *Cancellable)

var (
	SimplePermissionGetType func() O.Type
	SimplePermissionNew     func(allowed T.Gboolean) *Permission
)

type Socket struct {
	Parent O.Object
	_      *struct{}
}

var (
	SocketGetType   func() O.Type
	SocketNew       func(family SocketFamily, t SocketType, protocol SocketProtocol, err **T.GError) *Socket
	SocketNewFromFd func(fd int, err **T.GError) *Socket

	socketGetFd               func(s *Socket) int
	socketGetFamily           func(s *Socket) SocketFamily
	socketGetSocketType       func(s *Socket) SocketType
	socketGetProtocol         func(s *Socket) SocketProtocol
	socketGetLocalAddress     func(s *Socket, err **T.GError) *SocketAddress
	socketGetRemoteAddress    func(s *Socket, err **T.GError) *SocketAddress
	socketSetBlocking         func(s *Socket, blocking T.Gboolean)
	socketGetBlocking         func(s *Socket) T.Gboolean
	socketSetKeepalive        func(s *Socket, keepalive T.Gboolean)
	socketGetKeepalive        func(s *Socket) T.Gboolean
	socketGetListenBacklog    func(s *Socket) int
	socketSetListenBacklog    func(s *Socket, backlog int)
	socketGetTimeout          func(s *Socket) uint
	socketSetTimeout          func(s *Socket, timeout uint)
	socketIsConnected         func(s *Socket) T.Gboolean
	socketBind                func(s *Socket, address *SocketAddress, allowReuse T.Gboolean, err **T.GError) T.Gboolean
	socketConnect             func(s *Socket, address *SocketAddress, cancellable *Cancellable, err **T.GError) T.Gboolean
	socketCheckConnectResult  func(s *Socket, err **T.GError) T.Gboolean
	socketConditionCheck      func(s *Socket, condition T.GIOCondition) T.GIOCondition
	socketConditionWait       func(s *Socket, condition T.GIOCondition, cancellable *Cancellable, err **T.GError) T.Gboolean
	socketAccept              func(s *Socket, cancellable *Cancellable, err **T.GError) *Socket
	socketListen              func(s *Socket, err **T.GError) T.Gboolean
	socketReceive             func(s *Socket, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	socketReceiveFrom         func(s *Socket, address **SocketAddress, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	socketSend                func(s *Socket, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	socketSendTo              func(s *Socket, address *SocketAddress, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	socketReceiveMessage      func(s *Socket, address **SocketAddress, vectors *InputVector, numVectors int, messages ***SocketControlMessage, numMessages *int, flags *int, cancellable *Cancellable, err **T.GError) T.Gssize
	socketSendMessage         func(s *Socket, address *SocketAddress, vectors *T.GOutputVector, numVectors int, messages **SocketControlMessage, numMessages int, flags int, cancellable *Cancellable, err **T.GError) T.Gssize
	socketClose               func(s *Socket, err **T.GError) T.Gboolean
	socketShutdown            func(s *Socket, shutdownRead T.Gboolean, shutdownWrite T.Gboolean, err **T.GError) T.Gboolean
	socketIsClosed            func(s *Socket) T.Gboolean
	socketCreateSource        func(s *Socket, condition T.GIOCondition, cancellable *Cancellable) *O.Source
	socketSpeaksIpv4          func(s *Socket) T.Gboolean
	socketGetCredentials      func(s *Socket, err **T.GError) *Credentials
	socketReceiveWithBlocking func(s *Socket, buffer string, size T.Gsize, blocking T.Gboolean, cancellable *Cancellable, err **T.GError) T.Gssize
	socketSendWithBlocking    func(s *Socket, buffer string, size T.Gsize, blocking T.Gboolean, cancellable *Cancellable, err **T.GError) T.Gssize
)

func (s *Socket) GetFd() int                                    { return socketGetFd(s) }
func (s *Socket) GetFamily() SocketFamily                       { return socketGetFamily(s) }
func (s *Socket) GetSocketType() SocketType                     { return socketGetSocketType(s) }
func (s *Socket) GetProtocol() SocketProtocol                   { return socketGetProtocol(s) }
func (s *Socket) GetLocalAddress(err **T.GError) *SocketAddress { return socketGetLocalAddress(s, err) }
func (s *Socket) GetRemoteAddress(err **T.GError) *SocketAddress {
	return socketGetRemoteAddress(s, err)
}
func (s *Socket) SetBlocking(blocking T.Gboolean)   { socketSetBlocking(s, blocking) }
func (s *Socket) GetBlocking() T.Gboolean           { return socketGetBlocking(s) }
func (s *Socket) SetKeepalive(keepalive T.Gboolean) { socketSetKeepalive(s, keepalive) }
func (s *Socket) GetKeepalive() T.Gboolean          { return socketGetKeepalive(s) }
func (s *Socket) GetListenBacklog() int             { return socketGetListenBacklog(s) }
func (s *Socket) SetListenBacklog(backlog int)      { socketSetListenBacklog(s, backlog) }
func (s *Socket) GetTimeout() uint                  { return socketGetTimeout(s) }
func (s *Socket) SetTimeout(timeout uint)           { socketSetTimeout(s, timeout) }
func (s *Socket) IsConnected() T.Gboolean           { return socketIsConnected(s) }
func (s *Socket) Bind(address *SocketAddress, allowReuse T.Gboolean, err **T.GError) T.Gboolean {
	return socketBind(s, address, allowReuse, err)
}
func (s *Socket) Connect(address *SocketAddress, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return socketConnect(s, address, cancellable, err)
}
func (s *Socket) CheckConnectResult(err **T.GError) T.Gboolean {
	return socketCheckConnectResult(s, err)
}
func (s *Socket) ConditionCheck(condition T.GIOCondition) T.GIOCondition {
	return socketConditionCheck(s, condition)
}
func (s *Socket) ConditionWait(condition T.GIOCondition, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return socketConditionWait(s, condition, cancellable, err)
}
func (s *Socket) Accept(cancellable *Cancellable, err **T.GError) *Socket {
	return socketAccept(s, cancellable, err)
}
func (s *Socket) Listen(err **T.GError) T.Gboolean { return socketListen(s, err) }
func (s *Socket) Receive(buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return socketReceive(s, buffer, size, cancellable, err)
}
func (s *Socket) ReceiveFrom(address **SocketAddress, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return socketReceiveFrom(s, address, buffer, size, cancellable, err)
}
func (s *Socket) Send(buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return socketSend(s, buffer, size, cancellable, err)
}
func (s *Socket) SendTo(address *SocketAddress, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return socketSendTo(s, address, buffer, size, cancellable, err)
}
func (s *Socket) ReceiveMessage(address **SocketAddress, vectors *InputVector, numVectors int, messages ***SocketControlMessage, numMessages, flags *int, cancellable *Cancellable, err **T.GError) T.Gssize {
	return socketReceiveMessage(s, address, vectors, numVectors, messages, numMessages, flags, cancellable, err)
}
func (s *Socket) SendMessage(address *SocketAddress, vectors *T.GOutputVector, numVectors int, messages **SocketControlMessage, numMessages, flags int, cancellable *Cancellable, err **T.GError) T.Gssize {
	return socketSendMessage(s, address, vectors, numVectors, messages, numMessages, flags, cancellable, err)
}
func (s *Socket) Close(err **T.GError) T.Gboolean { return socketClose(s, err) }
func (s *Socket) Shutdown(shutdownRead, shutdownWrite T.Gboolean, err **T.GError) T.Gboolean {
	return socketShutdown(s, shutdownRead, shutdownWrite, err)
}
func (s *Socket) IsClosed() T.Gboolean { return socketIsClosed(s) }
func (s *Socket) CreateSource(condition T.GIOCondition, cancellable *Cancellable) *O.Source {
	return socketCreateSource(s, condition, cancellable)
}
func (s *Socket) SpeaksIpv4() T.Gboolean                     { return socketSpeaksIpv4(s) }
func (s *Socket) GetCredentials(err **T.GError) *Credentials { return socketGetCredentials(s, err) }
func (s *Socket) ReceiveWithBlocking(buffer string, size T.Gsize, blocking T.Gboolean, cancellable *Cancellable, err **T.GError) T.Gssize {
	return socketReceiveWithBlocking(s, buffer, size, blocking, cancellable, err)
}
func (s *Socket) SendWithBlocking(buffer string, size T.Gsize, blocking T.Gboolean, cancellable *Cancellable, err **T.GError) T.Gssize {
	return socketSendWithBlocking(s, buffer, size, blocking, cancellable, err)
}

type SocketAddress struct {
	Parent O.Object
}

var (
	SocketAddressGetType       func() O.Type
	SocketAddressNewFromNative func(native T.Gpointer, len T.Gsize) *SocketAddress

	socketAddressGetFamily     func(s *SocketAddress) SocketFamily
	socketAddressGetNativeSize func(s *SocketAddress) T.Gssize
	socketAddressToNative      func(s *SocketAddress, dest T.Gpointer, destlen T.Gsize, err **T.GError) T.Gboolean
)

func (s *SocketAddress) GetFamily() SocketFamily { return socketAddressGetFamily(s) }
func (s *SocketAddress) GetNativeSize() T.Gssize { return socketAddressGetNativeSize(s) }
func (s *SocketAddress) ToNative(dest T.Gpointer, destlen T.Gsize, err **T.GError) T.Gboolean {
	return socketAddressToNative(s, dest, destlen, err)
}

type SocketAddressEnumerator struct {
	Parent O.Object
}

var (
	SocketAddressEnumeratorGetType func() O.Type

	socketAddressEnumeratorNext       func(s *SocketAddressEnumerator, cancellable *Cancellable, err **T.GError) *SocketAddress
	socketAddressEnumeratorNextAsync  func(s *SocketAddressEnumerator, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	socketAddressEnumeratorNextFinish func(s *SocketAddressEnumerator, result *AsyncResult, err **T.GError) *SocketAddress
)

func (s *SocketAddressEnumerator) Next(cancellable *Cancellable, err **T.GError) *SocketAddress {
	return socketAddressEnumeratorNext(s, cancellable, err)
}
func (s *SocketAddressEnumerator) NextAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	socketAddressEnumeratorNextAsync(s, cancellable, callback, userData)
}
func (s *SocketAddressEnumerator) NextFinish(result *AsyncResult, err **T.GError) *SocketAddress {
	return socketAddressEnumeratorNextFinish(s, result, err)
}

type SocketClient struct {
	Parent O.Object
	_      *struct{}
}

var (
	SocketClientGetType func() O.Type
	SocketClientNew     func() *SocketClient

	socketClientAddApplicationProxy    func(s *SocketClient, protocol string)
	socketClientConnect                func(s *SocketClient, connectable *SocketConnectable, cancellable *Cancellable, err **T.GError) *SocketConnection
	socketClientConnectAsync           func(s *SocketClient, connectable *SocketConnectable, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	socketClientConnectFinish          func(s *SocketClient, result *AsyncResult, err **T.GError) *SocketConnection
	socketClientConnectToHost          func(s *SocketClient, hostAndPort string, defaultPort uint16, cancellable *Cancellable, err **T.GError) *SocketConnection
	socketClientConnectToHostAsync     func(s *SocketClient, hostAndPort string, defaultPort uint16, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	socketClientConnectToHostFinish    func(s *SocketClient, result *AsyncResult, err **T.GError) *SocketConnection
	socketClientConnectToService       func(s *SocketClient, domain string, service string, cancellable *Cancellable, err **T.GError) *SocketConnection
	socketClientConnectToServiceAsync  func(s *SocketClient, domain string, service string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	socketClientConnectToServiceFinish func(s *SocketClient, result *AsyncResult, err **T.GError) *SocketConnection
	socketClientConnectToUri           func(s *SocketClient, uri string, defaultPort uint16, cancellable *Cancellable, err **T.GError) *SocketConnection
	socketClientConnectToUriAsync      func(s *SocketClient, uri string, defaultPort uint16, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	socketClientConnectToUriFinish     func(s *SocketClient, result *AsyncResult, err **T.GError) *SocketConnection
	socketClientGetEnableProxy         func(s *SocketClient) T.Gboolean
	socketClientGetFamily              func(s *SocketClient) SocketFamily
	socketClientGetLocalAddress        func(s *SocketClient) *SocketAddress
	socketClientGetProtocol            func(s *SocketClient) SocketProtocol
	socketClientGetSocketType          func(s *SocketClient) SocketType
	socketClientGetTimeout             func(s *SocketClient) uint
	socketClientGetTls                 func(s *SocketClient) T.Gboolean
	socketClientGetTlsValidationFlags  func(s *SocketClient) TlsCertificateFlags
	socketClientSetEnableProxy         func(s *SocketClient, enable T.Gboolean)
	socketClientSetFamily              func(s *SocketClient, family SocketFamily)
	socketClientSetLocalAddress        func(s *SocketClient, address *SocketAddress)
	socketClientSetProtocol            func(s *SocketClient, protocol SocketProtocol)
	socketClientSetSocketType          func(s *SocketClient, t SocketType)
	socketClientSetTimeout             func(s *SocketClient, timeout uint)
	socketClientSetTls                 func(s *SocketClient, tls T.Gboolean)
	socketClientSetTlsValidationFlags  func(s *SocketClient, flags TlsCertificateFlags)
)

func (s *SocketClient) AddApplicationProxy(protocol string) {
	socketClientAddApplicationProxy(s, protocol)
}
func (s *SocketClient) Connect(connectable *SocketConnectable, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return socketClientConnect(s, connectable, cancellable, err)
}
func (s *SocketClient) ConnectAsync(connectable *SocketConnectable, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	socketClientConnectAsync(s, connectable, cancellable, callback, userData)
}
func (s *SocketClient) ConnectFinish(result *AsyncResult, err **T.GError) *SocketConnection {
	return socketClientConnectFinish(s, result, err)
}
func (s *SocketClient) ConnectToHost(hostAndPort string, defaultPort uint16, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return socketClientConnectToHost(s, hostAndPort, defaultPort, cancellable, err)
}
func (s *SocketClient) ConnectToHostAsync(hostAndPort string, defaultPort uint16, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	socketClientConnectToHostAsync(s, hostAndPort, defaultPort, cancellable, callback, userData)
}
func (s *SocketClient) ConnectToHostFinish(result *AsyncResult, err **T.GError) *SocketConnection {
	return socketClientConnectToHostFinish(s, result, err)
}
func (s *SocketClient) ConnectToService(domain, service string, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return socketClientConnectToService(s, domain, service, cancellable, err)
}
func (s *SocketClient) ConnectToServiceAsync(domain, service string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	socketClientConnectToServiceAsync(s, domain, service, cancellable, callback, userData)
}
func (s *SocketClient) ConnectToServiceFinish(result *AsyncResult, err **T.GError) *SocketConnection {
	return socketClientConnectToServiceFinish(s, result, err)
}
func (s *SocketClient) ConnectToUri(uri string, defaultPort uint16, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return socketClientConnectToUri(s, uri, defaultPort, cancellable, err)
}
func (s *SocketClient) ConnectToUriAsync(uri string, defaultPort uint16, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	socketClientConnectToUriAsync(s, uri, defaultPort, cancellable, callback, userData)
}
func (s *SocketClient) ConnectToUriFinish(result *AsyncResult, err **T.GError) *SocketConnection {
	return socketClientConnectToUriFinish(s, result, err)
}
func (s *SocketClient) GetEnableProxy() T.Gboolean      { return socketClientGetEnableProxy(s) }
func (s *SocketClient) GetFamily() SocketFamily         { return socketClientGetFamily(s) }
func (s *SocketClient) GetLocalAddress() *SocketAddress { return socketClientGetLocalAddress(s) }
func (s *SocketClient) GetProtocol() SocketProtocol     { return socketClientGetProtocol(s) }
func (s *SocketClient) GetSocketType() SocketType       { return socketClientGetSocketType(s) }
func (s *SocketClient) GetTimeout() uint                { return socketClientGetTimeout(s) }
func (s *SocketClient) GetTls() T.Gboolean              { return socketClientGetTls(s) }
func (s *SocketClient) GetTlsValidationFlags() TlsCertificateFlags {
	return socketClientGetTlsValidationFlags(s)
}
func (s *SocketClient) SetEnableProxy(enable T.Gboolean) { socketClientSetEnableProxy(s, enable) }
func (s *SocketClient) SetFamily(family SocketFamily)    { socketClientSetFamily(s, family) }
func (s *SocketClient) SetLocalAddress(address *SocketAddress) {
	socketClientSetLocalAddress(s, address)
}
func (s *SocketClient) SetProtocol(protocol SocketProtocol) { socketClientSetProtocol(s, protocol) }
func (s *SocketClient) SetSocketType(t SocketType)          { socketClientSetSocketType(s, t) }
func (s *SocketClient) SetTimeout(timeout uint)             { socketClientSetTimeout(s, timeout) }
func (s *SocketClient) SetTls(tls T.Gboolean)               { socketClientSetTls(s, tls) }
func (s *SocketClient) SetTlsValidationFlags(flags TlsCertificateFlags) {
	socketClientSetTlsValidationFlags(s, flags)
}

type SocketConnectable struct{}

var (
	SocketConnectableGetType func() O.Type

	socketConnectableEnumerate      func(s *SocketConnectable) *SocketAddressEnumerator
	socketConnectableProxyEnumerate func(s *SocketConnectable) *SocketAddressEnumerator
)

func (s *SocketConnectable) Enumerate() *SocketAddressEnumerator { return socketConnectableEnumerate(s) }
func (s *SocketConnectable) ProxyEnumerate() *SocketAddressEnumerator {
	return socketConnectableProxyEnumerate(s)
}

type SocketConnection struct {
	Parent IOStream
	_      *struct{}
}

var (
	SocketConnectionGetType func() O.Type

	SocketConnectionFactoryCreateConnection func(socket *Socket) *SocketConnection
	SocketConnectionFactoryLookupType       func(family SocketFamily, t SocketType, protocolId int) O.Type
	SocketConnectionFactoryRegisterType     func(GType O.Type, family SocketFamily, t SocketType, protocol int)

	socketConnectionGetLocalAddress  func(s *SocketConnection, err **T.GError) *SocketAddress
	socketConnectionGetRemoteAddress func(s *SocketConnection, err **T.GError) *SocketAddress
	socketConnectionGetSocket        func(s *SocketConnection) *Socket
)

func (s *SocketConnection) GetLocalAddress(err **T.GError) *SocketAddress {
	return socketConnectionGetLocalAddress(s, err)
}
func (s *SocketConnection) GetRemoteAddress(err **T.GError) *SocketAddress {
	return socketConnectionGetRemoteAddress(s, err)
}
func (s *SocketConnection) GetSocket() *Socket { return socketConnectionGetSocket(s) }

type SocketControlMessage struct {
	Parent O.Object
	_      *struct{}
}

var (
	SocketControlMessageGetType func() O.Type

	SocketControlMessageDeserialize func(level int, typ int, size T.Gsize, data T.Gpointer) *SocketControlMessage

	socketControlMessageGetLevel   func(s *SocketControlMessage) int
	socketControlMessageGetMsgType func(s *SocketControlMessage) int
	socketControlMessageGetSize    func(s *SocketControlMessage) T.Gsize
	socketControlMessageSerialize  func(s *SocketControlMessage, data T.Gpointer)
)

func (s *SocketControlMessage) GetLevel() int             { return socketControlMessageGetLevel(s) }
func (s *SocketControlMessage) GetMsgType() int           { return socketControlMessageGetMsgType(s) }
func (s *SocketControlMessage) GetSize() T.Gsize          { return socketControlMessageGetSize(s) }
func (s *SocketControlMessage) Serialize(data T.Gpointer) { socketControlMessageSerialize(s, data) }

type SocketFamily Enum

const (
	SOCKET_FAMILY_INVALID SocketFamily = iota
	SOCKET_FAMILY_UNIX
	SOCKET_FAMILY_IPV4
	SOCKET_FAMILY_IPV6 SocketFamily = 23
)

var SocketFamilyGetType func() O.Type

type SocketListener struct {
	Parent O.Object
	_      *struct{}
}

var (
	SocketListenerGetType func() O.Type
	SocketListenerNew     func() *SocketListener

	socketListenerAccept             func(s *SocketListener, sourceObject **O.Object, cancellable *Cancellable, err **T.GError) *SocketConnection
	socketListenerAcceptAsync        func(s *SocketListener, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	socketListenerAcceptFinish       func(s *SocketListener, result *AsyncResult, sourceObject **O.Object, err **T.GError) *SocketConnection
	socketListenerAcceptSocket       func(s *SocketListener, sourceObject **O.Object, cancellable *Cancellable, err **T.GError) *Socket
	socketListenerAcceptSocketAsync  func(s *SocketListener, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	socketListenerAcceptSocketFinish func(s *SocketListener, result *AsyncResult, sourceObject **O.Object, err **T.GError) *Socket
	socketListenerAddAddress         func(s *SocketListener, address *SocketAddress, t SocketType, protocol SocketProtocol, sourceObject *O.Object, effectiveAddress **SocketAddress, err **T.GError) T.Gboolean
	socketListenerAddAnyInetPort     func(s *SocketListener, sourceObject *O.Object, err **T.GError) uint16
	socketListenerAddInetPort        func(s *SocketListener, port uint16, sourceObject *O.Object, err **T.GError) T.Gboolean
	socketListenerAddSocket          func(s *SocketListener, socket *Socket, sourceObject *O.Object, err **T.GError) T.Gboolean
	socketListenerClose              func(s *SocketListener)
	socketListenerSetBacklog         func(s *SocketListener, listenBacklog int)
)

func (s *SocketListener) Accept(sourceObject **O.Object, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return socketListenerAccept(s, sourceObject, cancellable, err)
}
func (s *SocketListener) AcceptAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	socketListenerAcceptAsync(s, cancellable, callback, userData)
}
func (s *SocketListener) AcceptFinish(result *AsyncResult, sourceObject **O.Object, err **T.GError) *SocketConnection {
	return socketListenerAcceptFinish(s, result, sourceObject, err)
}
func (s *SocketListener) AcceptSocket(sourceObject **O.Object, cancellable *Cancellable, err **T.GError) *Socket {
	return socketListenerAcceptSocket(s, sourceObject, cancellable, err)
}
func (s *SocketListener) AcceptSocketAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	socketListenerAcceptSocketAsync(s, cancellable, callback, userData)
}
func (s *SocketListener) AcceptSocketFinish(result *AsyncResult, sourceObject **O.Object, err **T.GError) *Socket {
	return socketListenerAcceptSocketFinish(s, result, sourceObject, err)
}
func (s *SocketListener) AddAddress(address *SocketAddress, t SocketType, protocol SocketProtocol, sourceObject *O.Object, effectiveAddress **SocketAddress, err **T.GError) T.Gboolean {
	return socketListenerAddAddress(s, address, t, protocol, sourceObject, effectiveAddress, err)
}
func (s *SocketListener) AddAnyInetPort(sourceObject *O.Object, err **T.GError) uint16 {
	return socketListenerAddAnyInetPort(s, sourceObject, err)
}
func (s *SocketListener) AddInetPort(port uint16, sourceObject *O.Object, err **T.GError) T.Gboolean {
	return socketListenerAddInetPort(s, port, sourceObject, err)
}
func (s *SocketListener) AddSocket(socket *Socket, sourceObject *O.Object, err **T.GError) T.Gboolean {
	return socketListenerAddSocket(s, socket, sourceObject, err)
}
func (s *SocketListener) Close()                       { socketListenerClose(s) }
func (s *SocketListener) SetBacklog(listenBacklog int) { socketListenerSetBacklog(s, listenBacklog) }

var SocketMsgFlagsGetType func() O.Type

type SocketProtocol Enum

const (
	SOCKET_PROTOCOL_UNKNOWN SocketProtocol = -1
	SOCKET_PROTOCOL_DEFAULT SocketProtocol = 0
	SOCKET_PROTOCOL_TCP     SocketProtocol = 6
	SOCKET_PROTOCOL_UDP     SocketProtocol = 17
	SOCKET_PROTOCOL_SCTP    SocketProtocol = 132
)

var SocketProtocolGetType func() O.Type

type SocketService struct {
	Parent SocketListener
	_      *struct{}
}

var (
	SocketServiceGetType func() O.Type
	SocketServiceNew     func() *SocketService

	socketServiceIsActive func(s *SocketService) T.Gboolean
	socketServiceStart    func(s *SocketService)
	socketServiceStop     func(s *SocketService)
)

func (s *SocketService) IsActive() T.Gboolean { return socketServiceIsActive(s) }
func (s *SocketService) Start()               { socketServiceStart(s) }
func (s *SocketService) Stop()                { socketServiceStop(s) }

type SocketType Enum

const (
	SOCKET_TYPE_INVALID SocketType = iota
	SOCKET_TYPE_STREAM
	SOCKET_TYPE_DATAGRAM
	SOCKET_TYPE_SEQPACKET
)

var SocketTypeGetType func() O.Type

type SrvTarget struct{}

var (
	SrvTargetGetType func() O.Type
	SrvTargetNew     func(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget

	SrvTargetListSort func(targets *T.GList) *T.GList

	srvTargetCopy        func(s *SrvTarget) *SrvTarget
	srvTargetFree        func(s *SrvTarget)
	srvTargetGetHostname func(s *SrvTarget) string
	srvTargetGetPort     func(s *SrvTarget) uint16
	srvTargetGetPriority func(s *SrvTarget) uint16
	srvTargetGetWeight   func(s *SrvTarget) uint16
)

func (s *SrvTarget) Copy() *SrvTarget    { return srvTargetCopy(s) }
func (s *SrvTarget) Free()               { srvTargetFree(s) }
func (s *SrvTarget) GetHostname() string { return srvTargetGetHostname(s) }
func (s *SrvTarget) GetPort() uint16     { return srvTargetGetPort(s) }
func (s *SrvTarget) GetPriority() uint16 { return srvTargetGetPriority(s) }
func (s *SrvTarget) GetWeight() uint16   { return srvTargetGetWeight(s) }
