// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Icon struct{}

var (
	IconGetType      func() O.Type
	IconNewForString func(str string, err **T.GError) *Icon

	IconHash func(icon T.Gconstpointer) uint

	IconEqual    func(i *Icon, icon2 *Icon) bool
	IconToString func(i *Icon) string
)

type InetAddress struct {
	Parent O.Object
	_      *struct{}
}

var (
	InetAddressGetType       func() O.Type
	InetAddressNewAny        func(family SocketFamily) *InetAddress
	InetAddressNewFromBytes  func(bytes *uint8, family SocketFamily) *InetAddress
	InetAddressNewFromString func(string string) *InetAddress
	InetAddressNewLoopback   func(family SocketFamily) *InetAddress

	InetAddressGetFamily        func(i *InetAddress) SocketFamily
	InetAddressGetIsAny         func(i *InetAddress) bool
	InetAddressGetIsLinkLocal   func(i *InetAddress) bool
	InetAddressGetIsLoopback    func(i *InetAddress) bool
	InetAddressGetIsMcGlobal    func(i *InetAddress) bool
	InetAddressGetIsMcLinkLocal func(i *InetAddress) bool
	InetAddressGetIsMcNodeLocal func(i *InetAddress) bool
	InetAddressGetIsMcOrgLocal  func(i *InetAddress) bool
	InetAddressGetIsMcSiteLocal func(i *InetAddress) bool
	InetAddressGetIsMulticast   func(i *InetAddress) bool
	InetAddressGetIsSiteLocal   func(i *InetAddress) bool
	InetAddressGetNativeSize    func(i *InetAddress) T.Gsize
	InetAddressToBytes          func(i *InetAddress) *uint8
	InetAddressToString         func(i *InetAddress) string
)

func (i *InetAddress) GetFamily() SocketFamily { return InetAddressGetFamily(i) }
func (i *InetAddress) GetIsAny() bool          { return InetAddressGetIsAny(i) }
func (i *InetAddress) GetIsLinkLocal() bool    { return InetAddressGetIsLinkLocal(i) }
func (i *InetAddress) GetIsLoopback() bool     { return InetAddressGetIsLoopback(i) }
func (i *InetAddress) GetIsMcGlobal() bool     { return InetAddressGetIsMcGlobal(i) }
func (i *InetAddress) GetIsMcLinkLocal() bool  { return InetAddressGetIsMcLinkLocal(i) }
func (i *InetAddress) GetIsMcNodeLocal() bool  { return InetAddressGetIsMcNodeLocal(i) }
func (i *InetAddress) GetIsMcOrgLocal() bool   { return InetAddressGetIsMcOrgLocal(i) }
func (i *InetAddress) GetIsMcSiteLocal() bool  { return InetAddressGetIsMcSiteLocal(i) }
func (i *InetAddress) GetIsMulticast() bool    { return InetAddressGetIsMulticast(i) }
func (i *InetAddress) GetIsSiteLocal() bool    { return InetAddressGetIsSiteLocal(i) }
func (i *InetAddress) GetNativeSize() T.Gsize  { return InetAddressGetNativeSize(i) }
func (i *InetAddress) ToBytes() *uint8         { return InetAddressToBytes(i) }
func (i *InetAddress) ToString() string        { return InetAddressToString(i) }

type InetSocketAddress struct {
	Parent SocketAddress
	_      *struct{}
}

var (
	InetSocketAddressGetType func() O.Type
	InetSocketAddressNew     func(address *InetAddress, port uint16) *SocketAddress

	InetSocketAddressGetAddress func(i *InetSocketAddress) *InetAddress
	InetSocketAddressGetPort    func(i *InetSocketAddress) uint16
)

func (i *InetSocketAddress) GetAddress() *InetAddress { return InetSocketAddressGetAddress(i) }
func (i *InetSocketAddress) GetPort() uint16          { return InetSocketAddressGetPort(i) }

type Initable struct{}

var (
	InitableGetType   func() O.Type
	InitableNew       func(objectType O.Type, cancellable *Cancellable, e **T.GError, firstPropertyName string, v ...VArg) T.Gpointer
	InitableNewv      func(objectType O.Type, nParameters uint, parameters *T.GParameter, cancellable *Cancellable, err **T.GError) T.Gpointer
	InitableNewValist func(objectType O.Type, firstPropertyName string, varArgs T.VaList, cancellable *Cancellable, err **T.GError) *O.Object

	InitableInit func(i *Initable, cancellable *Cancellable, err **T.GError) bool
)

func (i *Initable) Init(cancellable *Cancellable, err **T.GError) bool {
	return InitableInit(i, cancellable, err)
}

type InputStream struct {
	Parent O.Object
	_      *struct{}
}

var (
	InputStreamGetType func() O.Type

	InputStreamClearPending func(i *InputStream)
	InputStreamClose        func(i *InputStream, cancellable *Cancellable, err **T.GError) bool
	InputStreamCloseAsync   func(i *InputStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	InputStreamCloseFinish  func(i *InputStream, result *AsyncResult, err **T.GError) bool
	InputStreamHasPending   func(i *InputStream) bool
	InputStreamIsClosed     func(i *InputStream) bool
	InputStreamRead         func(i *InputStream, buffer *T.Void, count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	InputStreamReadAll      func(i *InputStream, buffer *T.Void, count T.Gsize, bytesRead *T.Gsize, cancellable *Cancellable, err **T.GError) bool
	InputStreamReadAsync    func(i *InputStream, buffer *T.Void, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	InputStreamReadFinish   func(i *InputStream, result *AsyncResult, err **T.GError) T.Gssize
	InputStreamSetPending   func(i *InputStream, err **T.GError) bool
	InputStreamSkip         func(i *InputStream, count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	InputStreamSkipAsync    func(i *InputStream, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	InputStreamSkipFinish   func(i *InputStream, result *AsyncResult, err **T.GError) T.Gssize
)

func (i *InputStream) ClearPending() { InputStreamClearPending(i) }
func (i *InputStream) Close(cancellable *Cancellable, err **T.GError) bool {
	return InputStreamClose(i, cancellable, err)
}
func (i *InputStream) CloseAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	InputStreamCloseAsync(i, ioPriority, cancellable, callback, userData)
}
func (i *InputStream) CloseFinish(result *AsyncResult, err **T.GError) bool {
	return InputStreamCloseFinish(i, result, err)
}
func (i *InputStream) HasPending() bool { return InputStreamHasPending(i) }
func (i *InputStream) IsClosed() bool   { return InputStreamIsClosed(i) }
func (i *InputStream) Read(buffer *T.Void, count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return InputStreamRead(i, buffer, count, cancellable, err)
}
func (i *InputStream) ReadAll(buffer *T.Void, count T.Gsize, bytesRead *T.Gsize, cancellable *Cancellable, err **T.GError) bool {
	return InputStreamReadAll(i, buffer, count, bytesRead, cancellable, err)
}
func (i *InputStream) ReadAsync(buffer *T.Void, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	InputStreamReadAsync(i, buffer, count, ioPriority, cancellable, callback, userData)
}
func (i *InputStream) ReadFinish(result *AsyncResult, err **T.GError) T.Gssize {
	return InputStreamReadFinish(i, result, err)
}
func (i *InputStream) SetPending(err **T.GError) bool { return InputStreamSetPending(i, err) }
func (i *InputStream) Skip(count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return InputStreamSkip(i, count, cancellable, err)
}
func (i *InputStream) SkipAsync(count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	InputStreamSkipAsync(i, count, ioPriority, cancellable, callback, userData)
}
func (i *InputStream) SkipFinish(result *AsyncResult, err **T.GError) T.Gssize {
	return InputStreamSkipFinish(i, result, err)
}

type InputVector struct {
	Buffer T.Gpointer
	Size   T.Gsize
}

var (
	IoErrorQuark          func() L.Quark
	IoErrorFromErrno      func(errNo int) IOErrorEnum
	IoErrorFromWin32Error func(errorCode int) IOErrorEnum
)

type IOErrorEnum Enum

const (
	IO_ERROR_FAILED IOErrorEnum = iota
	IO_ERROR_NOT_FOUND
	IO_ERROR_EXISTS
	IO_ERROR_IS_DIRECTORY
	IO_ERROR_NOT_DIRECTORY
	IO_ERROR_NOT_EMPTY
	IO_ERROR_NOT_REGULAR_FILE
	IO_ERROR_NOT_SYMBOLIC_LINK
	IO_ERROR_NOT_MOUNTABLE_FILE
	IO_ERROR_FILENAME_TOO_LONG
	IO_ERROR_INVALID_FILENAME
	IO_ERROR_TOO_MANY_LINKS
	IO_ERROR_NO_SPACE
	IO_ERROR_INVALID_ARGUMENT
	IO_ERROR_PERMISSION_DENIED
	IO_ERROR_NOT_SUPPORTED
	IO_ERROR_NOT_MOUNTED
	IO_ERROR_ALREADY_MOUNTED
	IO_ERROR_CLOSED
	IO_ERROR_CANCELLED
	IO_ERROR_PENDING
	IO_ERROR_READ_ONLY
	IO_ERROR_CANT_CREATE_BACKUP
	IO_ERROR_WRONG_ETAG
	IO_ERROR_TIMED_OUT
	IO_ERROR_WOULD_RECURSE
	IO_ERROR_BUSY
	IO_ERROR_WOULD_BLOCK
	IO_ERROR_HOST_NOT_FOUND
	IO_ERROR_WOULD_MERGE
	IO_ERROR_FAILED_HANDLED
	IO_ERROR_TOO_MANY_OPEN_FILES
	IO_ERROR_NOT_INITIALIZED
	IO_ERROR_ADDRESS_IN_USE
	IO_ERROR_PARTIAL_INPUT
	IO_ERROR_INVALID_DATA
	IO_ERROR_DBUS_ERROR
	IO_ERROR_HOST_UNREACHABLE
	IO_ERROR_NETWORK_UNREACHABLE
	IO_ERROR_CONNECTION_REFUSED
	IO_ERROR_PROXY_FAILED
	IO_ERROR_PROXY_AUTH_FAILED
	IO_ERROR_PROXY_NEED_AUTH
	IO_ERROR_PROXY_NOT_ALLOWED
)

var IoErrorEnumGetType func() O.Type

type IOExtension struct{}

var (
	IoExtensionGetName     func(i *IOExtension) string
	IoExtensionGetPriority func(i *IOExtension) int
	IoExtensionGetType     func(i *IOExtension) O.Type
	IoExtensionRefClass    func(i *IOExtension) *O.TypeClass
)

func (i *IOExtension) GetName() string        { return IoExtensionGetName(i) }
func (i *IOExtension) GetPriority() int       { return IoExtensionGetPriority(i) }
func (i *IOExtension) GetType() O.Type        { return IoExtensionGetType(i) }
func (i *IOExtension) RefClass() *O.TypeClass { return IoExtensionRefClass(i) }

type IOExtensionPoint struct{}

var (
	IoExtensionPointLookup   func(name string) *IOExtensionPoint
	IoExtensionPointRegister func(name string) *IOExtensionPoint

	IoExtensionPointImplement func(extensionPointName string, typ O.Type, extensionName string, priority int) *IOExtension

	IoExtensionPointGetExtensionByName func(i *IOExtensionPoint, name string) *IOExtension
	IoExtensionPointGetExtensions      func(i *IOExtensionPoint) *T.GList
	IoExtensionPointGetRequiredType    func(i *IOExtensionPoint) O.Type
	IoExtensionPointSetRequiredType    func(i *IOExtensionPoint, typ O.Type)
)

func (i *IOExtensionPoint) GetExtensionByName(name string) *IOExtension {
	return IoExtensionPointGetExtensionByName(i, name)
}
func (i *IOExtensionPoint) GetExtensions() *T.GList    { return IoExtensionPointGetExtensions(i) }
func (i *IOExtensionPoint) GetRequiredType() O.Type    { return IoExtensionPointGetRequiredType(i) }
func (i *IOExtensionPoint) SetRequiredType(typ O.Type) { IoExtensionPointSetRequiredType(i, typ) }

type IOModule struct{}

var (
	IoModuleGetType func() O.Type
	IoModuleNew     func(filename string) *IOModule

	IoModuleQuery               func() []string
	IoModulesLoadAllInDirectory func(dirname string) *T.GList
	IoModulesScanAllInDirectory func(dirname string)

	IoModuleLoad   func(i *IOModule)
	IoModuleUnload func(i *IOModule)
)

func (i *IOModule) Load()   { IoModuleLoad(i) }
func (i *IOModule) Unload() { IoModuleUnload(i) }

var (
	IoSchedulerCancelAllJobs func()

	IoSchedulerPushJob func(i IOSchedulerJobFunc, userData T.Gpointer, notify T.GDestroyNotify, ioPriority int, cancellable *Cancellable)
)

type IOSchedulerJobFunc func(
	job *IOSchedulerJob,
	cancellable *Cancellable,
	userData T.Gpointer) bool

type IOSchedulerJob struct{}

var (
	IoSchedulerJobSendToMainloop      func(i *IOSchedulerJob, f O.SourceFunc, userData T.Gpointer, notify T.GDestroyNotify) bool
	IoSchedulerJobSendToMainloopAsync func(i *IOSchedulerJob, f O.SourceFunc, userData T.Gpointer, notify T.GDestroyNotify)
)

func (i *IOSchedulerJob) SendToMainloop(f O.SourceFunc, userData T.Gpointer, notify T.GDestroyNotify) bool {
	return IoSchedulerJobSendToMainloop(i, f, userData, notify)
}
func (i *IOSchedulerJob) SendToMainloopAsync(f O.SourceFunc, userData T.Gpointer, notify T.GDestroyNotify) {
	IoSchedulerJobSendToMainloopAsync(i, f, userData, notify)
}

type IOStream struct {
	Parent O.Object
	_      *struct{}
}

var (
	IoStreamGetType func() O.Type

	IoStreamSpliceFinish func(result *AsyncResult, err **T.GError) bool

	IoStreamClearPending    func(i *IOStream)
	IoStreamClose           func(i *IOStream, cancellable *Cancellable, err **T.GError) bool
	IoStreamCloseAsync      func(i *IOStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	IoStreamCloseFinish     func(i *IOStream, result *AsyncResult, err **T.GError) bool
	IoStreamGetInputStream  func(i *IOStream) *InputStream
	IoStreamGetOutputStream func(i *IOStream) *OutputStream
	IoStreamHasPending      func(i *IOStream) bool
	IoStreamIsClosed        func(i *IOStream) bool
	IoStreamSetPending      func(i *IOStream, err **T.GError) bool
	IoStreamSpliceAsync     func(i *IOStream, stream2 *IOStream, flags IOStreamSpliceFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
)

func (i *IOStream) ClearPending() { IoStreamClearPending(i) }
func (i *IOStream) Close(cancellable *Cancellable, err **T.GError) bool {
	return IoStreamClose(i, cancellable, err)
}
func (i *IOStream) CloseAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	IoStreamCloseAsync(i, ioPriority, cancellable, callback, userData)
}
func (i *IOStream) CloseFinish(result *AsyncResult, err **T.GError) bool {
	return IoStreamCloseFinish(i, result, err)
}
func (i *IOStream) GetInputStream() *InputStream   { return IoStreamGetInputStream(i) }
func (i *IOStream) GetOutputStream() *OutputStream { return IoStreamGetOutputStream(i) }
func (i *IOStream) HasPending() bool               { return IoStreamHasPending(i) }
func (i *IOStream) IsClosed() bool                 { return IoStreamIsClosed(i) }
func (i *IOStream) SetPending(err **T.GError) bool { return IoStreamSetPending(i, err) }
func (i *IOStream) SpliceAsync(stream2 *IOStream, flags IOStreamSpliceFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	IoStreamSpliceAsync(i, stream2, flags, ioPriority, cancellable, callback, userData)
}

type IOStreamSpliceFlags Enum

const (
	IO_STREAM_SPLICE_CLOSE_STREAM1 IOStreamSpliceFlags = 1 << iota
	IO_STREAM_SPLICE_CLOSE_STREAM2
	IO_STREAM_SPLICE_WAIT_FOR_BOTH
	IO_STREAM_SPLICE_NONE IOStreamSpliceFlags = 0
)

var IoStreamSpliceFlagsGetType func() O.Type
