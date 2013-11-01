// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Seekable struct{}

var (
	SeekableGetType func() O.Type

	SeekableCanSeek     func(seekable *Seekable) bool
	SeekableCanTruncate func(seekable *Seekable) bool
	SeekableSeek        func(seekable *Seekable, offset T.Goffset, typ L.SeekType, cancellable *Cancellable, err **T.GError) bool
	SeekableTell        func(seekable *Seekable) T.Goffset
	SeekableTruncate    func(seekable *Seekable, offset T.Goffset, cancellable *Cancellable, err **T.GError) bool
)

func (s *Seekable) CanSeek() bool     { return SeekableCanSeek(s) }
func (s *Seekable) CanTruncate() bool { return SeekableCanTruncate(s) }
func (s *Seekable) Seek(offset T.Goffset, typ L.SeekType, cancellable *Cancellable, err **T.GError) bool {
	return SeekableSeek(s, offset, typ, cancellable, err)
}
func (s *Seekable) Tell() T.Goffset { return SeekableTell(s) }
func (s *Seekable) Truncate(offset T.Goffset, cancellable *Cancellable, err **T.GError) bool {
	return SeekableTruncate(s, offset, cancellable, err)
}

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

	SettingsApply           func(s *Settings)
	SettingsBind            func(s *Settings, key string, object T.Gpointer, property string, flags SettingsBindFlags)
	SettingsBindWithMapping func(s *Settings, key string, object T.Gpointer, property string, flags SettingsBindFlags, getMapping SettingsBindGetMapping, setMapping SettingsBindSetMapping, userData T.Gpointer, destroy T.GDestroyNotify)
	SettingsBindWritable    func(s *Settings, key string, object T.Gpointer, property string, inverted bool)
	SettingsDelay           func(s *Settings)
	SettingsGet             func(s *Settings, key, format string, v ...VArg)
	SettingsGetBoolean      func(s *Settings, key string) bool
	SettingsGetChild        func(s *Settings, name string) *Settings
	SettingsGetDouble       func(s *Settings, key string) float64
	SettingsGetEnum         func(s *Settings, key string) int
	SettingsGetFlags        func(s *Settings, key string) uint
	SettingsGetHasUnapplied func(s *Settings) bool
	SettingsGetInt          func(s *Settings, key string) int
	SettingsGetMapped       func(s *Settings, key string, mapping SettingsGetMapping, userData T.Gpointer) T.Gpointer
	SettingsGetRange        func(s *Settings, key string) *L.Variant
	SettingsGetString       func(s *Settings, key string) string
	SettingsGetStrv         func(s *Settings, key string) []string
	SettingsGetValue        func(s *Settings, key string) *L.Variant
	SettingsIsWritable      func(s *Settings, name string) bool
	SettingsListKeys        func(s *Settings) []string
	SettingsListChildren    func(s *Settings) []string
	SettingsRangeCheck      func(s *Settings, key string, value *L.Variant) bool
	SettingsReset           func(s *Settings, key string)
	SettingsRevert          func(s *Settings)
	SettingsSet             func(s *Settings, key, format string, v ...VArg) bool
	SettingsSetBoolean      func(s *Settings, key string, value bool) bool
	SettingsSetDouble       func(s *Settings, key string, value float64) bool
	SettingsSetEnum         func(s *Settings, key string, value int) bool
	SettingsSetFlags        func(s *Settings, key string, value uint) bool
	SettingsSetInt          func(s *Settings, key string, value int) bool
	SettingsSetString       func(s *Settings, key, value string) bool
	SettingsSetStrv         func(s *Settings, key string, value []string) bool
	SettingsSetValue        func(s *Settings, key string, value *L.Variant) bool
)

func (s *Settings) Apply() { SettingsApply(s) }
func (s *Settings) Bind(key string, object T.Gpointer, property string, flags SettingsBindFlags) {
	SettingsBind(s, key, object, property, flags)
}
func (s *Settings) BindWithMapping(key string, object T.Gpointer, property string, flags SettingsBindFlags, getMapping SettingsBindGetMapping, setMapping SettingsBindSetMapping, userData T.Gpointer, destroy T.GDestroyNotify) {
	SettingsBindWithMapping(s, key, object, property, flags, getMapping, setMapping, userData, destroy)
}
func (s *Settings) BindWritable(key string, object T.Gpointer, property string, inverted bool) {
	SettingsBindWritable(s, key, object, property, inverted)
}
func (s *Settings) Delay()                            { SettingsDelay(s) }
func (s *Settings) Get(key, format string, v ...VArg) { SettingsGet(s, key, format, v) }
func (s *Settings) GetBoolean(key string) bool        { return SettingsGetBoolean(s, key) }
func (s *Settings) GetChild(name string) *Settings    { return SettingsGetChild(s, name) }
func (s *Settings) GetDouble(key string) float64      { return SettingsGetDouble(s, key) }
func (s *Settings) GetEnum(key string) int            { return SettingsGetEnum(s, key) }
func (s *Settings) GetFlags(key string) uint          { return SettingsGetFlags(s, key) }
func (s *Settings) GetHasUnapplied() bool             { return SettingsGetHasUnapplied(s) }
func (s *Settings) GetInt(key string) int             { return SettingsGetInt(s, key) }
func (s *Settings) GetMapped(key string, mapping SettingsGetMapping, userData T.Gpointer) T.Gpointer {
	return SettingsGetMapped(s, key, mapping, userData)
}
func (s *Settings) GetRange(key string) *L.Variant { return SettingsGetRange(s, key) }
func (s *Settings) GetString(key string) string    { return SettingsGetString(s, key) }
func (s *Settings) GetStrv(key string) []string    { return SettingsGetStrv(s, key) }
func (s *Settings) GetValue(key string) *L.Variant { return SettingsGetValue(s, key) }
func (s *Settings) IsWritable(name string) bool    { return SettingsIsWritable(s, name) }
func (s *Settings) ListKeys() []string             { return SettingsListKeys(s) }
func (s *Settings) ListChildren() []string         { return SettingsListChildren(s) }
func (s *Settings) RangeCheck(key string, value *L.Variant) bool {
	return SettingsRangeCheck(s, key, value)
}
func (s *Settings) Reset(key string) { SettingsReset(s, key) }
func (s *Settings) Revert()          { SettingsRevert(s) }
func (s *Settings) Set(key, format string, v ...VArg) bool {
	return SettingsSet(s, key, format, v)
}
func (s *Settings) SetBoolean(key string, value bool) bool {
	return SettingsSetBoolean(s, key, value)
}
func (s *Settings) SetDouble(key string, value float64) bool {
	return SettingsSetDouble(s, key, value)
}
func (s *Settings) SetEnum(key string, value int) bool   { return SettingsSetEnum(s, key, value) }
func (s *Settings) SetFlags(key string, value uint) bool { return SettingsSetFlags(s, key, value) }
func (s *Settings) SetInt(key string, value int) bool    { return SettingsSetInt(s, key, value) }
func (s *Settings) SetString(key, value string) bool     { return SettingsSetString(s, key, value) }
func (s *Settings) SetStrv(key string, value []string) bool {
	return SettingsSetStrv(s, key, value)
}
func (s *Settings) SetValue(key string, value *L.Variant) bool {
	return SettingsSetValue(s, key, value)
}

type SettingsBackend struct{}

var (
	SettingsBackendGetType func() O.Type

	SettingsBackendFlattenTree func(tree *L.Tree, path []string, keys ***T.Gchar, values ***L.Variant)
	SettingsBackendGetDefault  func() *SettingsBackend

	SettingsBackendKeysChanged         func(s *SettingsBackend, path string, items []string, originTag T.Gpointer)
	SettingsBackendChanged             func(s *SettingsBackend, key string, originTag T.Gpointer)
	SettingsBackendChangedTree         func(s *SettingsBackend, tree *L.Tree, originTag T.Gpointer)
	SettingsBackendPathChanged         func(s *SettingsBackend, path string, originTag T.Gpointer)
	SettingsBackendPathWritableChanged func(s *SettingsBackend, path string)
	SettingsBackendWritableChanged     func(s *SettingsBackend, key string)
)

func (s *SettingsBackend) KeysChanged(path string, items []string, originTag T.Gpointer) {
	SettingsBackendKeysChanged(s, path, items, originTag)
}
func (s *SettingsBackend) Changed(key string, originTag T.Gpointer) {
	SettingsBackendChanged(s, key, originTag)
}
func (s *SettingsBackend) ChangedTree(tree *L.Tree, originTag T.Gpointer) {
	SettingsBackendChangedTree(s, tree, originTag)
}
func (s *SettingsBackend) PathChanged(path string, originTag T.Gpointer) {
	SettingsBackendPathChanged(s, path, originTag)
}
func (s *SettingsBackend) PathWritableChanged(path string) {
	SettingsBackendPathWritableChanged(s, path)
}
func (s *SettingsBackend) WritableChanged(key string) { SettingsBackendWritableChanged(s, key) }

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
	variant *L.Variant,
	userData T.Gpointer) bool

type SettingsBindSetMapping func(
	value *O.Value,
	expectedType *L.VariantType,
	userData T.Gpointer) *L.Variant

type SettingsGetMapping func(
	value *L.Variant,
	result *T.Gpointer,
	userData T.Gpointer) bool

type SimpleAction struct {
	Parent O.Object
	_      *struct{}
}

var (
	SimpleActionGetType     func() O.Type
	SimpleActionNew         func(name string, parameterType *L.VariantType) *SimpleAction
	SimpleActionNewStateful func(name string, parameterType *L.VariantType, state *L.Variant) *SimpleAction

	SimpleActionSetEnabled func(s *SimpleAction, enabled bool)
)

func (s *SimpleAction) SetEnabled(enabled bool) { SimpleActionSetEnabled(s, enabled) }

type SimpleActionGroup struct {
	Parent O.Object
	_      *struct{}
}

var (
	SimpleActionGroupGetType func() O.Type
	SimpleActionGroupNew     func() *SimpleActionGroup

	SimpleActionGroupInsert func(s *SimpleActionGroup, action *Action)
	SimpleActionGroupLookup func(s *SimpleActionGroup, actionName string) *Action
	SimpleActionGroupRemove func(s *SimpleActionGroup, actionName string)
)

func (s *SimpleActionGroup) Insert(action *Action) { SimpleActionGroupInsert(s, action) }
func (s *SimpleActionGroup) Lookup(actionName string) *Action {
	return SimpleActionGroupLookup(s, actionName)
}
func (s *SimpleActionGroup) Remove(actionName string) { SimpleActionGroupRemove(s, actionName) }

var (
	SimpleAsyncReportErrorInIdle      func(object *O.Object, callback AsyncReadyCallback, userData T.Gpointer, domain L.Quark, code int, format string, v ...VArg)
	SimpleAsyncReportGerrorInIdle     func(object *O.Object, callback AsyncReadyCallback, userData T.Gpointer, err *T.GError)
	SimpleAsyncReportTakeGerrorInIdle func(object *O.Object, callback AsyncReadyCallback, userData T.Gpointer, err *T.GError)
)

type SimpleAsyncResult struct{}

var (
	SimpleAsyncResultGetType      func() O.Type
	SimpleAsyncResultNew          func(sourceObject *O.Object, callback AsyncReadyCallback, userData T.Gpointer, sourceTag T.Gpointer) *SimpleAsyncResult
	SimpleAsyncResultNewError     func(sourceObject *O.Object, callback AsyncReadyCallback, userData T.Gpointer, domain L.Quark, code int, format string, v ...VArg) *SimpleAsyncResult
	SimpleAsyncResultNewFromError func(sourceObject *O.Object, callback AsyncReadyCallback, userData T.Gpointer, err *T.GError) *SimpleAsyncResult
	SimpleAsyncResultNewTakeError func(sourceObject *O.Object, callback AsyncReadyCallback, userData T.Gpointer, err *T.GError) *SimpleAsyncResult

	SimpleAsyncResultIsValid func(result *AsyncResult, source *O.Object, sourceTag T.Gpointer) bool

	SimpleAsyncResultComplete              func(s *SimpleAsyncResult)
	SimpleAsyncResultCompleteInIdle        func(s *SimpleAsyncResult)
	SimpleAsyncResultGetOpResGboolean      func(s *SimpleAsyncResult) bool
	SimpleAsyncResultGetOpResGpointer      func(s *SimpleAsyncResult) T.Gpointer
	SimpleAsyncResultGetOpResGssize        func(s *SimpleAsyncResult) T.Gssize
	SimpleAsyncResultGetSourceTag          func(s *SimpleAsyncResult) T.Gpointer
	SimpleAsyncResultPropagateError        func(s *SimpleAsyncResult, dest **T.GError) bool
	SimpleAsyncResultRunInThread           func(s *SimpleAsyncResult, f SimpleAsyncThreadFunc, ioPriority int, cancellable *Cancellable)
	SimpleAsyncResultSetError              func(s *SimpleAsyncResult, domain L.Quark, code int, format string, v ...VArg)
	SimpleAsyncResultSetErrorVa            func(s *SimpleAsyncResult, domain L.Quark, code int, format string, args T.VaList)
	SimpleAsyncResultSetFromError          func(s *SimpleAsyncResult, err *T.GError)
	SimpleAsyncResultSetHandleCancellation func(s *SimpleAsyncResult, handleCancellation bool)
	SimpleAsyncResultSetOpResGboolean      func(s *SimpleAsyncResult, opRes bool)
	SimpleAsyncResultSetOpResGpointer      func(s *SimpleAsyncResult, opRes T.Gpointer, destroyOpRes T.GDestroyNotify)
	SimpleAsyncResultSetOpResGssize        func(s *SimpleAsyncResult, opRes T.Gssize)
	SimpleAsyncResultTakeError             func(s *SimpleAsyncResult, err *T.GError)
)

func (s *SimpleAsyncResult) Complete()                    { SimpleAsyncResultComplete(s) }
func (s *SimpleAsyncResult) CompleteInIdle()              { SimpleAsyncResultCompleteInIdle(s) }
func (s *SimpleAsyncResult) GetOpResGboolean() bool       { return SimpleAsyncResultGetOpResGboolean(s) }
func (s *SimpleAsyncResult) GetOpResGpointer() T.Gpointer { return SimpleAsyncResultGetOpResGpointer(s) }
func (s *SimpleAsyncResult) GetOpResGssize() T.Gssize     { return SimpleAsyncResultGetOpResGssize(s) }
func (s *SimpleAsyncResult) GetSourceTag() T.Gpointer     { return SimpleAsyncResultGetSourceTag(s) }
func (s *SimpleAsyncResult) PropagateError(dest **T.GError) bool {
	return SimpleAsyncResultPropagateError(s, dest)
}
func (s *SimpleAsyncResult) RunInThread(f SimpleAsyncThreadFunc, ioPriority int, cancellable *Cancellable) {
	SimpleAsyncResultRunInThread(s, f, ioPriority, cancellable)
}
func (s *SimpleAsyncResult) SetError(domain L.Quark, code int, format string, v ...VArg) {
	SimpleAsyncResultSetError(s, domain, code, format, v)
}
func (s *SimpleAsyncResult) SetErrorVa(domain L.Quark, code int, format string, args T.VaList) {
	SimpleAsyncResultSetErrorVa(s, domain, code, format, args)
}
func (s *SimpleAsyncResult) SetFromError(err *T.GError) { SimpleAsyncResultSetFromError(s, err) }
func (s *SimpleAsyncResult) SetHandleCancellation(handleCancellation bool) {
	SimpleAsyncResultSetHandleCancellation(s, handleCancellation)
}
func (s *SimpleAsyncResult) SetOpResGboolean(opRes bool) {
	SimpleAsyncResultSetOpResGboolean(s, opRes)
}
func (s *SimpleAsyncResult) SetOpResGpointer(opRes T.Gpointer, destroyOpRes T.GDestroyNotify) {
	SimpleAsyncResultSetOpResGpointer(s, opRes, destroyOpRes)
}
func (s *SimpleAsyncResult) SetOpResGssize(opRes T.Gssize) { SimpleAsyncResultSetOpResGssize(s, opRes) }
func (s *SimpleAsyncResult) TakeError(err *T.GError)       { SimpleAsyncResultTakeError(s, err) }

type SimpleAsyncThreadFunc func(
	res *SimpleAsyncResult,
	object *O.Object,
	cancellable *Cancellable)

var (
	SimplePermissionGetType func() O.Type
	SimplePermissionNew     func(allowed bool) *Permission
)

type Socket struct {
	Parent O.Object
	_      *struct{}
}

var (
	SocketGetType   func() O.Type
	SocketNew       func(family SocketFamily, t SocketType, protocol SocketProtocol, err **T.GError) *Socket
	SocketNewFromFd func(fd int, err **T.GError) *Socket

	SocketGetFd               func(s *Socket) int
	SocketGetFamily           func(s *Socket) SocketFamily
	SocketGetSocketType       func(s *Socket) SocketType
	SocketGetProtocol         func(s *Socket) SocketProtocol
	SocketGetLocalAddress     func(s *Socket, err **T.GError) *SocketAddress
	SocketGetRemoteAddress    func(s *Socket, err **T.GError) *SocketAddress
	SocketSetBlocking         func(s *Socket, blocking bool)
	SocketGetBlocking         func(s *Socket) bool
	SocketSetKeepalive        func(s *Socket, keepalive bool)
	SocketGetKeepalive        func(s *Socket) bool
	SocketGetListenBacklog    func(s *Socket) int
	SocketSetListenBacklog    func(s *Socket, backlog int)
	SocketGetTimeout          func(s *Socket) uint
	SocketSetTimeout          func(s *Socket, timeout uint)
	SocketIsConnected         func(s *Socket) bool
	SocketBind                func(s *Socket, address *SocketAddress, allowReuse bool, err **T.GError) bool
	SocketConnect             func(s *Socket, address *SocketAddress, cancellable *Cancellable, err **T.GError) bool
	SocketCheckConnectResult  func(s *Socket, err **T.GError) bool
	SocketConditionCheck      func(s *Socket, condition L.IOCondition) L.IOCondition
	SocketConditionWait       func(s *Socket, condition L.IOCondition, cancellable *Cancellable, err **T.GError) bool
	SocketAccept              func(s *Socket, cancellable *Cancellable, err **T.GError) *Socket
	SocketListen              func(s *Socket, err **T.GError) bool
	SocketReceive             func(s *Socket, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	SocketReceiveFrom         func(s *Socket, address **SocketAddress, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	SocketSend                func(s *Socket, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	SocketSendTo              func(s *Socket, address *SocketAddress, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	SocketReceiveMessage      func(s *Socket, address **SocketAddress, vectors *InputVector, numVectors int, messages ***SocketControlMessage, numMessages *int, flags *int, cancellable *Cancellable, err **T.GError) T.Gssize
	SocketSendMessage         func(s *Socket, address *SocketAddress, vectors *T.GOutputVector, numVectors int, messages **SocketControlMessage, numMessages int, flags int, cancellable *Cancellable, err **T.GError) T.Gssize
	SocketClose               func(s *Socket, err **T.GError) bool
	SocketShutdown            func(s *Socket, shutdownRead bool, shutdownWrite bool, err **T.GError) bool
	SocketIsClosed            func(s *Socket) bool
	SocketCreateSource        func(s *Socket, condition L.IOCondition, cancellable *Cancellable) *O.Source
	SocketSpeaksIpv4          func(s *Socket) bool
	SocketGetCredentials      func(s *Socket, err **T.GError) *Credentials
	SocketReceiveWithBlocking func(s *Socket, buffer string, size T.Gsize, blocking bool, cancellable *Cancellable, err **T.GError) T.Gssize
	SocketSendWithBlocking    func(s *Socket, buffer string, size T.Gsize, blocking bool, cancellable *Cancellable, err **T.GError) T.Gssize
)

func (s *Socket) GetFd() int                                    { return SocketGetFd(s) }
func (s *Socket) GetFamily() SocketFamily                       { return SocketGetFamily(s) }
func (s *Socket) GetSocketType() SocketType                     { return SocketGetSocketType(s) }
func (s *Socket) GetProtocol() SocketProtocol                   { return SocketGetProtocol(s) }
func (s *Socket) GetLocalAddress(err **T.GError) *SocketAddress { return SocketGetLocalAddress(s, err) }
func (s *Socket) GetRemoteAddress(err **T.GError) *SocketAddress {
	return SocketGetRemoteAddress(s, err)
}
func (s *Socket) SetBlocking(blocking bool)    { SocketSetBlocking(s, blocking) }
func (s *Socket) GetBlocking() bool            { return SocketGetBlocking(s) }
func (s *Socket) SetKeepalive(keepalive bool)  { SocketSetKeepalive(s, keepalive) }
func (s *Socket) GetKeepalive() bool           { return SocketGetKeepalive(s) }
func (s *Socket) GetListenBacklog() int        { return SocketGetListenBacklog(s) }
func (s *Socket) SetListenBacklog(backlog int) { SocketSetListenBacklog(s, backlog) }
func (s *Socket) GetTimeout() uint             { return SocketGetTimeout(s) }
func (s *Socket) SetTimeout(timeout uint)      { SocketSetTimeout(s, timeout) }
func (s *Socket) IsConnected() bool            { return SocketIsConnected(s) }
func (s *Socket) Bind(address *SocketAddress, allowReuse bool, err **T.GError) bool {
	return SocketBind(s, address, allowReuse, err)
}
func (s *Socket) Connect(address *SocketAddress, cancellable *Cancellable, err **T.GError) bool {
	return SocketConnect(s, address, cancellable, err)
}
func (s *Socket) CheckConnectResult(err **T.GError) bool {
	return SocketCheckConnectResult(s, err)
}
func (s *Socket) ConditionCheck(condition L.IOCondition) L.IOCondition {
	return SocketConditionCheck(s, condition)
}
func (s *Socket) ConditionWait(condition L.IOCondition, cancellable *Cancellable, err **T.GError) bool {
	return SocketConditionWait(s, condition, cancellable, err)
}
func (s *Socket) Accept(cancellable *Cancellable, err **T.GError) *Socket {
	return SocketAccept(s, cancellable, err)
}
func (s *Socket) Listen(err **T.GError) bool { return SocketListen(s, err) }
func (s *Socket) Receive(buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return SocketReceive(s, buffer, size, cancellable, err)
}
func (s *Socket) ReceiveFrom(address **SocketAddress, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return SocketReceiveFrom(s, address, buffer, size, cancellable, err)
}
func (s *Socket) Send(buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return SocketSend(s, buffer, size, cancellable, err)
}
func (s *Socket) SendTo(address *SocketAddress, buffer string, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return SocketSendTo(s, address, buffer, size, cancellable, err)
}
func (s *Socket) ReceiveMessage(address **SocketAddress, vectors *InputVector, numVectors int, messages ***SocketControlMessage, numMessages, flags *int, cancellable *Cancellable, err **T.GError) T.Gssize {
	return SocketReceiveMessage(s, address, vectors, numVectors, messages, numMessages, flags, cancellable, err)
}
func (s *Socket) SendMessage(address *SocketAddress, vectors *T.GOutputVector, numVectors int, messages **SocketControlMessage, numMessages, flags int, cancellable *Cancellable, err **T.GError) T.Gssize {
	return SocketSendMessage(s, address, vectors, numVectors, messages, numMessages, flags, cancellable, err)
}
func (s *Socket) Close(err **T.GError) bool { return SocketClose(s, err) }
func (s *Socket) Shutdown(shutdownRead, shutdownWrite bool, err **T.GError) bool {
	return SocketShutdown(s, shutdownRead, shutdownWrite, err)
}
func (s *Socket) IsClosed() bool { return SocketIsClosed(s) }
func (s *Socket) CreateSource(condition L.IOCondition, cancellable *Cancellable) *O.Source {
	return SocketCreateSource(s, condition, cancellable)
}
func (s *Socket) SpeaksIpv4() bool                           { return SocketSpeaksIpv4(s) }
func (s *Socket) GetCredentials(err **T.GError) *Credentials { return SocketGetCredentials(s, err) }
func (s *Socket) ReceiveWithBlocking(buffer string, size T.Gsize, blocking bool, cancellable *Cancellable, err **T.GError) T.Gssize {
	return SocketReceiveWithBlocking(s, buffer, size, blocking, cancellable, err)
}
func (s *Socket) SendWithBlocking(buffer string, size T.Gsize, blocking bool, cancellable *Cancellable, err **T.GError) T.Gssize {
	return SocketSendWithBlocking(s, buffer, size, blocking, cancellable, err)
}

type SocketAddress struct {
	Parent O.Object
}

var (
	SocketAddressGetType       func() O.Type
	SocketAddressNewFromNative func(native T.Gpointer, len T.Gsize) *SocketAddress

	SocketAddressGetFamily     func(s *SocketAddress) SocketFamily
	SocketAddressGetNativeSize func(s *SocketAddress) T.Gssize
	SocketAddressToNative      func(s *SocketAddress, dest T.Gpointer, destlen T.Gsize, err **T.GError) bool
)

func (s *SocketAddress) GetFamily() SocketFamily { return SocketAddressGetFamily(s) }
func (s *SocketAddress) GetNativeSize() T.Gssize { return SocketAddressGetNativeSize(s) }
func (s *SocketAddress) ToNative(dest T.Gpointer, destlen T.Gsize, err **T.GError) bool {
	return SocketAddressToNative(s, dest, destlen, err)
}

type SocketAddressEnumerator struct {
	Parent O.Object
}

var (
	SocketAddressEnumeratorGetType func() O.Type

	SocketAddressEnumeratorNext       func(s *SocketAddressEnumerator, cancellable *Cancellable, err **T.GError) *SocketAddress
	SocketAddressEnumeratorNextAsync  func(s *SocketAddressEnumerator, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	SocketAddressEnumeratorNextFinish func(s *SocketAddressEnumerator, result *AsyncResult, err **T.GError) *SocketAddress
)

func (s *SocketAddressEnumerator) Next(cancellable *Cancellable, err **T.GError) *SocketAddress {
	return SocketAddressEnumeratorNext(s, cancellable, err)
}
func (s *SocketAddressEnumerator) NextAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	SocketAddressEnumeratorNextAsync(s, cancellable, callback, userData)
}
func (s *SocketAddressEnumerator) NextFinish(result *AsyncResult, err **T.GError) *SocketAddress {
	return SocketAddressEnumeratorNextFinish(s, result, err)
}

type SocketClient struct {
	Parent O.Object
	_      *struct{}
}

var (
	SocketClientGetType func() O.Type
	SocketClientNew     func() *SocketClient

	SocketClientAddApplicationProxy    func(s *SocketClient, protocol string)
	SocketClientConnect                func(s *SocketClient, connectable *SocketConnectable, cancellable *Cancellable, err **T.GError) *SocketConnection
	SocketClientConnectAsync           func(s *SocketClient, connectable *SocketConnectable, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	SocketClientConnectFinish          func(s *SocketClient, result *AsyncResult, err **T.GError) *SocketConnection
	SocketClientConnectToHost          func(s *SocketClient, hostAndPort string, defaultPort uint16, cancellable *Cancellable, err **T.GError) *SocketConnection
	SocketClientConnectToHostAsync     func(s *SocketClient, hostAndPort string, defaultPort uint16, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	SocketClientConnectToHostFinish    func(s *SocketClient, result *AsyncResult, err **T.GError) *SocketConnection
	SocketClientConnectToService       func(s *SocketClient, domain string, service string, cancellable *Cancellable, err **T.GError) *SocketConnection
	SocketClientConnectToServiceAsync  func(s *SocketClient, domain string, service string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	SocketClientConnectToServiceFinish func(s *SocketClient, result *AsyncResult, err **T.GError) *SocketConnection
	SocketClientConnectToUri           func(s *SocketClient, uri string, defaultPort uint16, cancellable *Cancellable, err **T.GError) *SocketConnection
	SocketClientConnectToUriAsync      func(s *SocketClient, uri string, defaultPort uint16, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	SocketClientConnectToUriFinish     func(s *SocketClient, result *AsyncResult, err **T.GError) *SocketConnection
	SocketClientGetEnableProxy         func(s *SocketClient) bool
	SocketClientGetFamily              func(s *SocketClient) SocketFamily
	SocketClientGetLocalAddress        func(s *SocketClient) *SocketAddress
	SocketClientGetProtocol            func(s *SocketClient) SocketProtocol
	SocketClientGetSocketType          func(s *SocketClient) SocketType
	SocketClientGetTimeout             func(s *SocketClient) uint
	SocketClientGetTls                 func(s *SocketClient) bool
	SocketClientGetTlsValidationFlags  func(s *SocketClient) TlsCertificateFlags
	SocketClientSetEnableProxy         func(s *SocketClient, enable bool)
	SocketClientSetFamily              func(s *SocketClient, family SocketFamily)
	SocketClientSetLocalAddress        func(s *SocketClient, address *SocketAddress)
	SocketClientSetProtocol            func(s *SocketClient, protocol SocketProtocol)
	SocketClientSetSocketType          func(s *SocketClient, t SocketType)
	SocketClientSetTimeout             func(s *SocketClient, timeout uint)
	SocketClientSetTls                 func(s *SocketClient, tls bool)
	SocketClientSetTlsValidationFlags  func(s *SocketClient, flags TlsCertificateFlags)
)

func (s *SocketClient) AddApplicationProxy(protocol string) {
	SocketClientAddApplicationProxy(s, protocol)
}
func (s *SocketClient) Connect(connectable *SocketConnectable, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return SocketClientConnect(s, connectable, cancellable, err)
}
func (s *SocketClient) ConnectAsync(connectable *SocketConnectable, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	SocketClientConnectAsync(s, connectable, cancellable, callback, userData)
}
func (s *SocketClient) ConnectFinish(result *AsyncResult, err **T.GError) *SocketConnection {
	return SocketClientConnectFinish(s, result, err)
}
func (s *SocketClient) ConnectToHost(hostAndPort string, defaultPort uint16, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return SocketClientConnectToHost(s, hostAndPort, defaultPort, cancellable, err)
}
func (s *SocketClient) ConnectToHostAsync(hostAndPort string, defaultPort uint16, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	SocketClientConnectToHostAsync(s, hostAndPort, defaultPort, cancellable, callback, userData)
}
func (s *SocketClient) ConnectToHostFinish(result *AsyncResult, err **T.GError) *SocketConnection {
	return SocketClientConnectToHostFinish(s, result, err)
}
func (s *SocketClient) ConnectToService(domain, service string, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return SocketClientConnectToService(s, domain, service, cancellable, err)
}
func (s *SocketClient) ConnectToServiceAsync(domain, service string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	SocketClientConnectToServiceAsync(s, domain, service, cancellable, callback, userData)
}
func (s *SocketClient) ConnectToServiceFinish(result *AsyncResult, err **T.GError) *SocketConnection {
	return SocketClientConnectToServiceFinish(s, result, err)
}
func (s *SocketClient) ConnectToUri(uri string, defaultPort uint16, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return SocketClientConnectToUri(s, uri, defaultPort, cancellable, err)
}
func (s *SocketClient) ConnectToUriAsync(uri string, defaultPort uint16, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	SocketClientConnectToUriAsync(s, uri, defaultPort, cancellable, callback, userData)
}
func (s *SocketClient) ConnectToUriFinish(result *AsyncResult, err **T.GError) *SocketConnection {
	return SocketClientConnectToUriFinish(s, result, err)
}
func (s *SocketClient) GetEnableProxy() bool            { return SocketClientGetEnableProxy(s) }
func (s *SocketClient) GetFamily() SocketFamily         { return SocketClientGetFamily(s) }
func (s *SocketClient) GetLocalAddress() *SocketAddress { return SocketClientGetLocalAddress(s) }
func (s *SocketClient) GetProtocol() SocketProtocol     { return SocketClientGetProtocol(s) }
func (s *SocketClient) GetSocketType() SocketType       { return SocketClientGetSocketType(s) }
func (s *SocketClient) GetTimeout() uint                { return SocketClientGetTimeout(s) }
func (s *SocketClient) GetTls() bool                    { return SocketClientGetTls(s) }
func (s *SocketClient) GetTlsValidationFlags() TlsCertificateFlags {
	return SocketClientGetTlsValidationFlags(s)
}
func (s *SocketClient) SetEnableProxy(enable bool)    { SocketClientSetEnableProxy(s, enable) }
func (s *SocketClient) SetFamily(family SocketFamily) { SocketClientSetFamily(s, family) }
func (s *SocketClient) SetLocalAddress(address *SocketAddress) {
	SocketClientSetLocalAddress(s, address)
}
func (s *SocketClient) SetProtocol(protocol SocketProtocol) { SocketClientSetProtocol(s, protocol) }
func (s *SocketClient) SetSocketType(t SocketType)          { SocketClientSetSocketType(s, t) }
func (s *SocketClient) SetTimeout(timeout uint)             { SocketClientSetTimeout(s, timeout) }
func (s *SocketClient) SetTls(tls bool)                     { SocketClientSetTls(s, tls) }
func (s *SocketClient) SetTlsValidationFlags(flags TlsCertificateFlags) {
	SocketClientSetTlsValidationFlags(s, flags)
}

type SocketConnectable struct{}

var (
	SocketConnectableGetType func() O.Type

	SocketConnectableEnumerate      func(s *SocketConnectable) *SocketAddressEnumerator
	SocketConnectableProxyEnumerate func(s *SocketConnectable) *SocketAddressEnumerator
)

func (s *SocketConnectable) Enumerate() *SocketAddressEnumerator { return SocketConnectableEnumerate(s) }
func (s *SocketConnectable) ProxyEnumerate() *SocketAddressEnumerator {
	return SocketConnectableProxyEnumerate(s)
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

	SocketConnectionGetLocalAddress  func(s *SocketConnection, err **T.GError) *SocketAddress
	SocketConnectionGetRemoteAddress func(s *SocketConnection, err **T.GError) *SocketAddress
	SocketConnectionGetSocket        func(s *SocketConnection) *Socket
)

func (s *SocketConnection) GetLocalAddress(err **T.GError) *SocketAddress {
	return SocketConnectionGetLocalAddress(s, err)
}
func (s *SocketConnection) GetRemoteAddress(err **T.GError) *SocketAddress {
	return SocketConnectionGetRemoteAddress(s, err)
}
func (s *SocketConnection) GetSocket() *Socket { return SocketConnectionGetSocket(s) }

type SocketControlMessage struct {
	Parent O.Object
	_      *struct{}
}

var (
	SocketControlMessageGetType func() O.Type

	SocketControlMessageDeserialize func(level int, typ int, size T.Gsize, data T.Gpointer) *SocketControlMessage

	SocketControlMessageGetLevel   func(s *SocketControlMessage) int
	SocketControlMessageGetMsgType func(s *SocketControlMessage) int
	SocketControlMessageGetSize    func(s *SocketControlMessage) T.Gsize
	SocketControlMessageSerialize  func(s *SocketControlMessage, data T.Gpointer)
)

func (s *SocketControlMessage) GetLevel() int             { return SocketControlMessageGetLevel(s) }
func (s *SocketControlMessage) GetMsgType() int           { return SocketControlMessageGetMsgType(s) }
func (s *SocketControlMessage) GetSize() T.Gsize          { return SocketControlMessageGetSize(s) }
func (s *SocketControlMessage) Serialize(data T.Gpointer) { SocketControlMessageSerialize(s, data) }

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

	SocketListenerAccept             func(s *SocketListener, sourceObject **O.Object, cancellable *Cancellable, err **T.GError) *SocketConnection
	SocketListenerAcceptAsync        func(s *SocketListener, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	SocketListenerAcceptFinish       func(s *SocketListener, result *AsyncResult, sourceObject **O.Object, err **T.GError) *SocketConnection
	SocketListenerAcceptSocket       func(s *SocketListener, sourceObject **O.Object, cancellable *Cancellable, err **T.GError) *Socket
	SocketListenerAcceptSocketAsync  func(s *SocketListener, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	SocketListenerAcceptSocketFinish func(s *SocketListener, result *AsyncResult, sourceObject **O.Object, err **T.GError) *Socket
	SocketListenerAddAddress         func(s *SocketListener, address *SocketAddress, t SocketType, protocol SocketProtocol, sourceObject *O.Object, effectiveAddress **SocketAddress, err **T.GError) bool
	SocketListenerAddAnyInetPort     func(s *SocketListener, sourceObject *O.Object, err **T.GError) uint16
	SocketListenerAddInetPort        func(s *SocketListener, port uint16, sourceObject *O.Object, err **T.GError) bool
	SocketListenerAddSocket          func(s *SocketListener, socket *Socket, sourceObject *O.Object, err **T.GError) bool
	SocketListenerClose              func(s *SocketListener)
	SocketListenerSetBacklog         func(s *SocketListener, listenBacklog int)
)

func (s *SocketListener) Accept(sourceObject **O.Object, cancellable *Cancellable, err **T.GError) *SocketConnection {
	return SocketListenerAccept(s, sourceObject, cancellable, err)
}
func (s *SocketListener) AcceptAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	SocketListenerAcceptAsync(s, cancellable, callback, userData)
}
func (s *SocketListener) AcceptFinish(result *AsyncResult, sourceObject **O.Object, err **T.GError) *SocketConnection {
	return SocketListenerAcceptFinish(s, result, sourceObject, err)
}
func (s *SocketListener) AcceptSocket(sourceObject **O.Object, cancellable *Cancellable, err **T.GError) *Socket {
	return SocketListenerAcceptSocket(s, sourceObject, cancellable, err)
}
func (s *SocketListener) AcceptSocketAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	SocketListenerAcceptSocketAsync(s, cancellable, callback, userData)
}
func (s *SocketListener) AcceptSocketFinish(result *AsyncResult, sourceObject **O.Object, err **T.GError) *Socket {
	return SocketListenerAcceptSocketFinish(s, result, sourceObject, err)
}
func (s *SocketListener) AddAddress(address *SocketAddress, t SocketType, protocol SocketProtocol, sourceObject *O.Object, effectiveAddress **SocketAddress, err **T.GError) bool {
	return SocketListenerAddAddress(s, address, t, protocol, sourceObject, effectiveAddress, err)
}
func (s *SocketListener) AddAnyInetPort(sourceObject *O.Object, err **T.GError) uint16 {
	return SocketListenerAddAnyInetPort(s, sourceObject, err)
}
func (s *SocketListener) AddInetPort(port uint16, sourceObject *O.Object, err **T.GError) bool {
	return SocketListenerAddInetPort(s, port, sourceObject, err)
}
func (s *SocketListener) AddSocket(socket *Socket, sourceObject *O.Object, err **T.GError) bool {
	return SocketListenerAddSocket(s, socket, sourceObject, err)
}
func (s *SocketListener) Close()                       { SocketListenerClose(s) }
func (s *SocketListener) SetBacklog(listenBacklog int) { SocketListenerSetBacklog(s, listenBacklog) }

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

	SocketServiceIsActive func(s *SocketService) bool
	SocketServiceStart    func(s *SocketService)
	SocketServiceStop     func(s *SocketService)
)

func (s *SocketService) IsActive() bool { return SocketServiceIsActive(s) }
func (s *SocketService) Start()         { SocketServiceStart(s) }
func (s *SocketService) Stop()          { SocketServiceStop(s) }

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

	SrvTargetCopy        func(s *SrvTarget) *SrvTarget
	SrvTargetFree        func(s *SrvTarget)
	SrvTargetGetHostname func(s *SrvTarget) string
	SrvTargetGetPort     func(s *SrvTarget) uint16
	SrvTargetGetPriority func(s *SrvTarget) uint16
	SrvTargetGetWeight   func(s *SrvTarget) uint16
)

func (s *SrvTarget) Copy() *SrvTarget    { return SrvTargetCopy(s) }
func (s *SrvTarget) Free()               { SrvTargetFree(s) }
func (s *SrvTarget) GetHostname() string { return SrvTargetGetHostname(s) }
func (s *SrvTarget) GetPort() uint16     { return SrvTargetGetPort(s) }
func (s *SrvTarget) GetPriority() uint16 { return SrvTargetGetPriority(s) }
func (s *SrvTarget) GetWeight() uint16   { return SrvTargetGetWeight(s) }
