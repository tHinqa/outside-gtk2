// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Icon struct{}

var (
	IconGetType      func() O.Type
	IconNewForString func(str string, err **T.GError) *Icon

	IconHash func(icon T.Gconstpointer) uint

	IconEqual    func(i *Icon, icon2 *Icon) T.Gboolean
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

	inetAddressGetFamily        func(i *InetAddress) SocketFamily
	inetAddressGetIsAny         func(i *InetAddress) T.Gboolean
	inetAddressGetIsLinkLocal   func(i *InetAddress) T.Gboolean
	inetAddressGetIsLoopback    func(i *InetAddress) T.Gboolean
	inetAddressGetIsMcGlobal    func(i *InetAddress) T.Gboolean
	inetAddressGetIsMcLinkLocal func(i *InetAddress) T.Gboolean
	inetAddressGetIsMcNodeLocal func(i *InetAddress) T.Gboolean
	inetAddressGetIsMcOrgLocal  func(i *InetAddress) T.Gboolean
	inetAddressGetIsMcSiteLocal func(i *InetAddress) T.Gboolean
	inetAddressGetIsMulticast   func(i *InetAddress) T.Gboolean
	inetAddressGetIsSiteLocal   func(i *InetAddress) T.Gboolean
	inetAddressGetNativeSize    func(i *InetAddress) T.Gsize
	inetAddressToBytes          func(i *InetAddress) *uint8
	inetAddressToString         func(i *InetAddress) string
)

func (i *InetAddress) GetFamily() SocketFamily      { return inetAddressGetFamily(i) }
func (i *InetAddress) GetIsAny() T.Gboolean         { return inetAddressGetIsAny(i) }
func (i *InetAddress) GetIsLinkLocal() T.Gboolean   { return inetAddressGetIsLinkLocal(i) }
func (i *InetAddress) GetIsLoopback() T.Gboolean    { return inetAddressGetIsLoopback(i) }
func (i *InetAddress) GetIsMcGlobal() T.Gboolean    { return inetAddressGetIsMcGlobal(i) }
func (i *InetAddress) GetIsMcLinkLocal() T.Gboolean { return inetAddressGetIsMcLinkLocal(i) }
func (i *InetAddress) GetIsMcNodeLocal() T.Gboolean { return inetAddressGetIsMcNodeLocal(i) }
func (i *InetAddress) GetIsMcOrgLocal() T.Gboolean  { return inetAddressGetIsMcOrgLocal(i) }
func (i *InetAddress) GetIsMcSiteLocal() T.Gboolean { return inetAddressGetIsMcSiteLocal(i) }
func (i *InetAddress) GetIsMulticast() T.Gboolean   { return inetAddressGetIsMulticast(i) }
func (i *InetAddress) GetIsSiteLocal() T.Gboolean   { return inetAddressGetIsSiteLocal(i) }
func (i *InetAddress) GetNativeSize() T.Gsize       { return inetAddressGetNativeSize(i) }
func (i *InetAddress) ToBytes() *uint8              { return inetAddressToBytes(i) }
func (i *InetAddress) ToString() string             { return inetAddressToString(i) }

type InetSocketAddress struct {
	Parent SocketAddress
	_      *struct{}
}

var (
	InetSocketAddressGetType func() O.Type
	InetSocketAddressNew     func(address *InetAddress, port uint16) *SocketAddress

	inetSocketAddressGetAddress func(i *InetSocketAddress) *InetAddress
	inetSocketAddressGetPort    func(i *InetSocketAddress) uint16
)

func (i *InetSocketAddress) GetAddress() *InetAddress { return inetSocketAddressGetAddress(i) }
func (i *InetSocketAddress) GetPort() uint16          { return inetSocketAddressGetPort(i) }

type Initable struct{}

var (
	InitableGetType   func() O.Type
	InitableNew       func(objectType O.Type, cancellable *Cancellable, e **T.GError, firstPropertyName string, v ...VArg) T.Gpointer
	InitableNewv      func(objectType O.Type, nParameters uint, parameters *T.GParameter, cancellable *Cancellable, err **T.GError) T.Gpointer
	InitableNewValist func(objectType O.Type, firstPropertyName string, varArgs T.VaList, cancellable *Cancellable, err **T.GError) *O.Object

	initableInit func(i *Initable, cancellable *Cancellable, err **T.GError) T.Gboolean
)

func (i *Initable) Init(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return initableInit(i, cancellable, err)
}

type InputStream struct {
	Parent O.Object
	_      *struct{}
}

var (
	InputStreamGetType func() O.Type

	inputStreamClearPending func(i *InputStream)
	inputStreamClose        func(i *InputStream, cancellable *Cancellable, err **T.GError) T.Gboolean
	inputStreamCloseAsync   func(i *InputStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	inputStreamCloseFinish  func(i *InputStream, result *AsyncResult, err **T.GError) T.Gboolean
	inputStreamHasPending   func(i *InputStream) T.Gboolean
	inputStreamIsClosed     func(i *InputStream) T.Gboolean
	inputStreamRead         func(i *InputStream, buffer *T.Void, count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	inputStreamReadAll      func(i *InputStream, buffer *T.Void, count T.Gsize, bytesRead *T.Gsize, cancellable *Cancellable, err **T.GError) T.Gboolean
	inputStreamReadAsync    func(i *InputStream, buffer *T.Void, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	inputStreamReadFinish   func(i *InputStream, result *AsyncResult, err **T.GError) T.Gssize
	inputStreamSetPending   func(i *InputStream, err **T.GError) T.Gboolean
	inputStreamSkip         func(i *InputStream, count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
	inputStreamSkipAsync    func(i *InputStream, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	inputStreamSkipFinish   func(i *InputStream, result *AsyncResult, err **T.GError) T.Gssize
)

func (i *InputStream) ClearPending() { inputStreamClearPending(i) }
func (i *InputStream) Close(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return inputStreamClose(i, cancellable, err)
}
func (i *InputStream) CloseAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	inputStreamCloseAsync(i, ioPriority, cancellable, callback, userData)
}
func (i *InputStream) CloseFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return inputStreamCloseFinish(i, result, err)
}
func (i *InputStream) HasPending() T.Gboolean { return inputStreamHasPending(i) }
func (i *InputStream) IsClosed() T.Gboolean   { return inputStreamIsClosed(i) }
func (i *InputStream) Read(buffer *T.Void, count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return inputStreamRead(i, buffer, count, cancellable, err)
}
func (i *InputStream) ReadAll(buffer *T.Void, count T.Gsize, bytesRead *T.Gsize, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return inputStreamReadAll(i, buffer, count, bytesRead, cancellable, err)
}
func (i *InputStream) ReadAsync(buffer *T.Void, count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	inputStreamReadAsync(i, buffer, count, ioPriority, cancellable, callback, userData)
}
func (i *InputStream) ReadFinish(result *AsyncResult, err **T.GError) T.Gssize {
	return inputStreamReadFinish(i, result, err)
}
func (i *InputStream) SetPending(err **T.GError) T.Gboolean { return inputStreamSetPending(i, err) }
func (i *InputStream) Skip(count T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return inputStreamSkip(i, count, cancellable, err)
}
func (i *InputStream) SkipAsync(count T.Gsize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	inputStreamSkipAsync(i, count, ioPriority, cancellable, callback, userData)
}
func (i *InputStream) SkipFinish(result *AsyncResult, err **T.GError) T.Gssize {
	return inputStreamSkipFinish(i, result, err)
}

type InputVector struct {
	Buffer T.Gpointer
	Size   T.Gsize
}

var (
	IoErrorQuark          func() T.GQuark
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
	ioExtensionGetName     func(i *IOExtension) string
	ioExtensionGetPriority func(i *IOExtension) int
	ioExtensionGetType     func(i *IOExtension) O.Type
	ioExtensionRefClass    func(i *IOExtension) *O.TypeClass
)

func (i *IOExtension) GetName() string        { return ioExtensionGetName(i) }
func (i *IOExtension) GetPriority() int       { return ioExtensionGetPriority(i) }
func (i *IOExtension) GetType() O.Type        { return ioExtensionGetType(i) }
func (i *IOExtension) RefClass() *O.TypeClass { return ioExtensionRefClass(i) }

type IOExtensionPoint struct{}

var (
	IoExtensionPointLookup   func(name string) *IOExtensionPoint
	IoExtensionPointRegister func(name string) *IOExtensionPoint

	IoExtensionPointImplement func(extensionPointName string, typ O.Type, extensionName string, priority int) *IOExtension

	ioExtensionPointGetExtensionByName func(i *IOExtensionPoint, name string) *IOExtension
	ioExtensionPointGetExtensions      func(i *IOExtensionPoint) *T.GList
	ioExtensionPointGetRequiredType    func(i *IOExtensionPoint) O.Type
	ioExtensionPointSetRequiredType    func(i *IOExtensionPoint, typ O.Type)
)

func (i *IOExtensionPoint) GetExtensionByName(name string) *IOExtension {
	return ioExtensionPointGetExtensionByName(i, name)
}
func (i *IOExtensionPoint) GetExtensions() *T.GList    { return ioExtensionPointGetExtensions(i) }
func (i *IOExtensionPoint) GetRequiredType() O.Type    { return ioExtensionPointGetRequiredType(i) }
func (i *IOExtensionPoint) SetRequiredType(typ O.Type) { ioExtensionPointSetRequiredType(i, typ) }

type IOModule struct{}

var (
	IoModuleGetType func() O.Type
	IoModuleNew     func(filename string) *IOModule

	IoModuleQuery               func() []string
	IoModulesLoadAllInDirectory func(dirname string) *T.GList
	IoModulesScanAllInDirectory func(dirname string)

	ioModuleLoad   func(i *IOModule)
	ioModuleUnload func(i *IOModule)
)

func (i *IOModule) Load()   { ioModuleLoad(i) }
func (i *IOModule) Unload() { ioModuleUnload(i) }

var (
	IoSchedulerCancelAllJobs func()

	IoSchedulerPushJob func(i IOSchedulerJobFunc, userData T.Gpointer, notify T.GDestroyNotify, ioPriority int, cancellable *Cancellable)
)

type IOSchedulerJobFunc func(
	job *IOSchedulerJob,
	cancellable *Cancellable,
	userData T.Gpointer) T.Gboolean

type IOSchedulerJob struct{}

var (
	ioSchedulerJobSendToMainloop      func(i *IOSchedulerJob, f O.SourceFunc, userData T.Gpointer, notify T.GDestroyNotify) T.Gboolean
	ioSchedulerJobSendToMainloopAsync func(i *IOSchedulerJob, f O.SourceFunc, userData T.Gpointer, notify T.GDestroyNotify)
)

func (i *IOSchedulerJob) SendToMainloop(f O.SourceFunc, userData T.Gpointer, notify T.GDestroyNotify) T.Gboolean {
	return ioSchedulerJobSendToMainloop(i, f, userData, notify)
}
func (i *IOSchedulerJob) SendToMainloopAsync(f O.SourceFunc, userData T.Gpointer, notify T.GDestroyNotify) {
	ioSchedulerJobSendToMainloopAsync(i, f, userData, notify)
}

type IOStream struct {
	Parent O.Object
	_      *struct{}
}

var (
	IoStreamGetType func() O.Type

	IoStreamSpliceFinish func(result *AsyncResult, err **T.GError) T.Gboolean

	ioStreamClearPending    func(i *IOStream)
	ioStreamClose           func(i *IOStream, cancellable *Cancellable, err **T.GError) T.Gboolean
	ioStreamCloseAsync      func(i *IOStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	ioStreamCloseFinish     func(i *IOStream, result *AsyncResult, err **T.GError) T.Gboolean
	ioStreamGetInputStream  func(i *IOStream) *InputStream
	ioStreamGetOutputStream func(i *IOStream) *OutputStream
	ioStreamHasPending      func(i *IOStream) T.Gboolean
	ioStreamIsClosed        func(i *IOStream) T.Gboolean
	ioStreamSetPending      func(i *IOStream, err **T.GError) T.Gboolean
	ioStreamSpliceAsync     func(i *IOStream, stream2 *IOStream, flags IOStreamSpliceFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
)

func (i *IOStream) ClearPending() { ioStreamClearPending(i) }
func (i *IOStream) Close(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return ioStreamClose(i, cancellable, err)
}
func (i *IOStream) CloseAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	ioStreamCloseAsync(i, ioPriority, cancellable, callback, userData)
}
func (i *IOStream) CloseFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return ioStreamCloseFinish(i, result, err)
}
func (i *IOStream) GetInputStream() *InputStream         { return ioStreamGetInputStream(i) }
func (i *IOStream) GetOutputStream() *OutputStream       { return ioStreamGetOutputStream(i) }
func (i *IOStream) HasPending() T.Gboolean               { return ioStreamHasPending(i) }
func (i *IOStream) IsClosed() T.Gboolean                 { return ioStreamIsClosed(i) }
func (i *IOStream) SetPending(err **T.GError) T.Gboolean { return ioStreamSetPending(i, err) }
func (i *IOStream) SpliceAsync(stream2 *IOStream, flags IOStreamSpliceFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	ioStreamSpliceAsync(i, stream2, flags, ioPriority, cancellable, callback, userData)
}

type IOStreamSpliceFlags Enum

const (
	IO_STREAM_SPLICE_CLOSE_STREAM1 IOStreamSpliceFlags = 1 << iota
	IO_STREAM_SPLICE_CLOSE_STREAM2
	IO_STREAM_SPLICE_WAIT_FOR_BOTH
	IO_STREAM_SPLICE_NONE IOStreamSpliceFlags = 0
)

var IoStreamSpliceFlagsGetType func() O.Type
