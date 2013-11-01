// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type DataInputStream struct {
	Parent BufferedInputStream
	_      *struct{}
}

var (
	DataInputStreamGetType func() O.Type
	DataInputStreamNew     func(baseStream *InputStream) *DataInputStream

	DataInputStreamGetByteOrder    func(d *DataInputStream) DataStreamByteOrder
	DataInputStreamGetNewlineType  func(d *DataInputStream) DataStreamNewlineType
	DataInputStreamReadByte        func(d *DataInputStream, cancellable *Cancellable, err **T.GError) T.Guchar
	DataInputStreamReadInt16       func(d *DataInputStream, cancellable *Cancellable, err **T.GError) int16
	DataInputStreamReadInt32       func(d *DataInputStream, cancellable *Cancellable, err **T.GError) T.GInt32
	DataInputStreamReadInt64       func(d *DataInputStream, cancellable *Cancellable, err **T.GError) int64
	DataInputStreamReadLine        func(d *DataInputStream, length *T.Gsize, cancellable *Cancellable, err **T.GError) string
	DataInputStreamReadLineAsync   func(d *DataInputStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DataInputStreamReadLineFinish  func(d *DataInputStream, result *AsyncResult, length *T.Gsize, err **T.GError) string
	DataInputStreamReadUint16      func(d *DataInputStream, cancellable *Cancellable, err **T.GError) uint16
	DataInputStreamReadUint32      func(d *DataInputStream, cancellable *Cancellable, err **T.GError) T.GUint32
	DataInputStreamReadUint64      func(d *DataInputStream, cancellable *Cancellable, err **T.GError) uint64
	DataInputStreamReadUntil       func(d *DataInputStream, stopChars string, length *T.Gsize, cancellable *Cancellable, err **T.GError) string
	DataInputStreamReadUntilAsync  func(d *DataInputStream, stopChars string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DataInputStreamReadUntilFinish func(d *DataInputStream, result *AsyncResult, length *T.Gsize, err **T.GError) string
	DataInputStreamReadUpto        func(d *DataInputStream, stopChars string, stopCharsLen T.Gssize, length *T.Gsize, cancellable *Cancellable, err **T.GError) string
	DataInputStreamReadUptoAsync   func(d *DataInputStream, stopChars string, stopCharsLen T.Gssize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DataInputStreamReadUptoFinish  func(d *DataInputStream, result *AsyncResult, length *T.Gsize, err **T.GError) string
	DataInputStreamSetByteOrder    func(d *DataInputStream, order DataStreamByteOrder)
	DataInputStreamSetNewlineType  func(d *DataInputStream, typ DataStreamNewlineType)
)

func (d *DataInputStream) GetByteOrder() DataStreamByteOrder { return DataInputStreamGetByteOrder(d) }
func (d *DataInputStream) GetNewlineType() DataStreamNewlineType {
	return DataInputStreamGetNewlineType(d)
}
func (d *DataInputStream) ReadByte(cancellable *Cancellable, err **T.GError) T.Guchar {
	return DataInputStreamReadByte(d, cancellable, err)
}
func (d *DataInputStream) ReadInt16(cancellable *Cancellable, err **T.GError) int16 {
	return DataInputStreamReadInt16(d, cancellable, err)
}
func (d *DataInputStream) ReadInt32(cancellable *Cancellable, err **T.GError) T.GInt32 {
	return DataInputStreamReadInt32(d, cancellable, err)
}
func (d *DataInputStream) ReadInt64(cancellable *Cancellable, err **T.GError) int64 {
	return DataInputStreamReadInt64(d, cancellable, err)
}
func (d *DataInputStream) ReadLine(length *T.Gsize, cancellable *Cancellable, err **T.GError) string {
	return DataInputStreamReadLine(d, length, cancellable, err)
}
func (d *DataInputStream) ReadLineAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DataInputStreamReadLineAsync(d, ioPriority, cancellable, callback, userData)
}
func (d *DataInputStream) ReadLineFinish(result *AsyncResult, length *T.Gsize, err **T.GError) string {
	return DataInputStreamReadLineFinish(d, result, length, err)
}
func (d *DataInputStream) ReadUint16(cancellable *Cancellable, err **T.GError) uint16 {
	return DataInputStreamReadUint16(d, cancellable, err)
}
func (d *DataInputStream) ReadUint32(cancellable *Cancellable, err **T.GError) T.GUint32 {
	return DataInputStreamReadUint32(d, cancellable, err)
}
func (d *DataInputStream) ReadUint64(cancellable *Cancellable, err **T.GError) uint64 {
	return DataInputStreamReadUint64(d, cancellable, err)
}
func (d *DataInputStream) ReadUntil(stopChars string, length *T.Gsize, cancellable *Cancellable, err **T.GError) string {
	return DataInputStreamReadUntil(d, stopChars, length, cancellable, err)
}
func (d *DataInputStream) ReadUntilAsync(stopChars string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DataInputStreamReadUntilAsync(d, stopChars, ioPriority, cancellable, callback, userData)
}
func (d *DataInputStream) ReadUntilFinish(result *AsyncResult, length *T.Gsize, err **T.GError) string {
	return DataInputStreamReadUntilFinish(d, result, length, err)
}
func (d *DataInputStream) ReadUpto(stopChars string, stopCharsLen T.Gssize, length *T.Gsize, cancellable *Cancellable, err **T.GError) string {
	return DataInputStreamReadUpto(d, stopChars, stopCharsLen, length, cancellable, err)
}
func (d *DataInputStream) ReadUptoAsync(stopChars string, stopCharsLen T.Gssize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DataInputStreamReadUptoAsync(d, stopChars, stopCharsLen, ioPriority, cancellable, callback, userData)
}
func (d *DataInputStream) ReadUptoFinish(result *AsyncResult, length *T.Gsize, err **T.GError) string {
	return DataInputStreamReadUptoFinish(d, result, length, err)
}
func (d *DataInputStream) SetByteOrder(order DataStreamByteOrder) {
	DataInputStreamSetByteOrder(d, order)
}
func (d *DataInputStream) SetNewlineType(typ DataStreamNewlineType) {
	DataInputStreamSetNewlineType(d, typ)
}

type DataOutputStream struct {
	Parent FilterOutputStream
	_      *struct{}
}

var (
	DataOutputStreamGetType func() O.Type
	DataOutputStreamNew     func(baseStream *OutputStream) *DataOutputStream

	DataOutputStreamGetByteOrder func(d *DataOutputStream) DataStreamByteOrder
	DataOutputStreamPutByte      func(d *DataOutputStream, data T.Guchar, cancellable *Cancellable, err **T.GError) bool
	DataOutputStreamPutInt16     func(d *DataOutputStream, data int16, cancellable *Cancellable, err **T.GError) bool
	DataOutputStreamPutInt32     func(d *DataOutputStream, data T.GInt32, cancellable *Cancellable, err **T.GError) bool
	DataOutputStreamPutInt64     func(d *DataOutputStream, data int64, cancellable *Cancellable, err **T.GError) bool
	DataOutputStreamPutString    func(d *DataOutputStream, str string, cancellable *Cancellable, err **T.GError) bool
	DataOutputStreamPutUint16    func(d *DataOutputStream, data uint16, cancellable *Cancellable, err **T.GError) bool
	DataOutputStreamPutUint32    func(d *DataOutputStream, data T.GUint32, cancellable *Cancellable, err **T.GError) bool
	DataOutputStreamPutUint64    func(d *DataOutputStream, data uint64, cancellable *Cancellable, err **T.GError) bool
	DataOutputStreamSetByteOrder func(d *DataOutputStream, order DataStreamByteOrder)
)

func (d *DataOutputStream) GetByteOrder() DataStreamByteOrder { return DataOutputStreamGetByteOrder(d) }
func (d *DataOutputStream) PutByte(data T.Guchar, cancellable *Cancellable, err **T.GError) bool {
	return DataOutputStreamPutByte(d, data, cancellable, err)
}
func (d *DataOutputStream) PutInt16(data int16, cancellable *Cancellable, err **T.GError) bool {
	return DataOutputStreamPutInt16(d, data, cancellable, err)
}
func (d *DataOutputStream) PutInt32(data T.GInt32, cancellable *Cancellable, err **T.GError) bool {
	return DataOutputStreamPutInt32(d, data, cancellable, err)
}
func (d *DataOutputStream) PutInt64(data int64, cancellable *Cancellable, err **T.GError) bool {
	return DataOutputStreamPutInt64(d, data, cancellable, err)
}
func (d *DataOutputStream) PutString(str string, cancellable *Cancellable, err **T.GError) bool {
	return DataOutputStreamPutString(d, str, cancellable, err)
}
func (d *DataOutputStream) PutUint16(data uint16, cancellable *Cancellable, err **T.GError) bool {
	return DataOutputStreamPutUint16(d, data, cancellable, err)
}
func (d *DataOutputStream) PutUint32(data T.GUint32, cancellable *Cancellable, err **T.GError) bool {
	return DataOutputStreamPutUint32(d, data, cancellable, err)
}
func (d *DataOutputStream) PutUint64(data uint64, cancellable *Cancellable, err **T.GError) bool {
	return DataOutputStreamPutUint64(d, data, cancellable, err)
}
func (d *DataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	DataOutputStreamSetByteOrder(d, order)
}

type DataStreamByteOrder Enum

const (
	DATA_STREAM_BYTE_ORDER_BIG_ENDIAN DataStreamByteOrder = iota
	DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN
	DATA_STREAM_BYTE_ORDER_HOST_ENDIAN
)

var DataStreamByteOrderGetType func() O.Type

type DataStreamNewlineType Enum

const (
	DATA_STREAM_NEWLINE_TYPE_LF DataStreamNewlineType = iota
	DATA_STREAM_NEWLINE_TYPE_CR
	DATA_STREAM_NEWLINE_TYPE_CR_LF
	DATA_STREAM_NEWLINE_TYPE_ANY
)

var DataStreamNewlineTypeGetType func() O.Type

var (
	DBusAddressGetForBusSync   func(busType BusType, cancellable *Cancellable, err **T.GError) string
	DBusAddressGetStream       func(address string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusAddressGetStreamFinish func(res *AsyncResult, outGuid **T.Gchar, err **T.GError) *IOStream
	DBusAddressGetStreamSync   func(address string, outGuid **T.Gchar, cancellable *Cancellable, err **T.GError) *IOStream
)

type DBusAnnotationInfo struct {
	RefCount    int
	Key         *T.Gchar
	Value       *T.Gchar
	Annotations **DBusAnnotationInfo
}

var (
	DBusAnnotationInfoGetType func() O.Type

	DBusAnnotationInfoLookup func(annotations **DBusAnnotationInfo, name string) string

	DBusAnnotationInfoRef   func(d *DBusAnnotationInfo) *DBusAnnotationInfo
	DBusAnnotationInfoUnref func(d *DBusAnnotationInfo)
)

func (d *DBusAnnotationInfo) Ref() *DBusAnnotationInfo { return DBusAnnotationInfoRef(d) }
func (d *DBusAnnotationInfo) Unref()                   { DBusAnnotationInfoUnref(d) }

type DBusArgInfo struct {
	RefCount    int
	Name        *T.Gchar
	Signature   *T.Gchar
	Annotations **DBusAnnotationInfo
}

var (
	DBusArgInfoGetType func() O.Type

	DBusArgInfoRef   func(d *DBusArgInfo) *DBusArgInfo
	DBusArgInfoUnref func(d *DBusArgInfo)
)

func (d *DBusArgInfo) Ref() *DBusArgInfo { return DBusArgInfoRef(d) }
func (d *DBusArgInfo) Unref()            { DBusArgInfoUnref(d) }

type DBusAuthObserver struct{}

var (
	DBusAuthObserverGetType func() O.Type
	DBusAuthObserverNew     func() *DBusAuthObserver

	DBusAuthObserverAuthorizeAuthenticatedPeer func(d *DBusAuthObserver, stream *IOStream, credentials *Credentials) bool
)

func (d *DBusAuthObserver) AuthorizeAuthenticatedPeer(stream *IOStream, credentials *Credentials) bool {
	return DBusAuthObserverAuthorizeAuthenticatedPeer(d, stream, credentials)
}

type DBusCallFlags Enum

const (
	DBUS_CALL_FLAGS_NO_AUTO_START DBusCallFlags = 1 << iota
	DBUS_CALL_FLAGS_NONE          DBusCallFlags = 0
)

var DBusCallFlagsGetType func() O.Type

type DBusCapabilityFlags Enum

const (
	DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING DBusCapabilityFlags = 1 << iota
	DBUS_CAPABILITY_FLAGS_NONE            DBusCapabilityFlags = 0
)

var DBusCapabilityFlagsGetType func() O.Type

type DBusConnection struct{}

var (
	DBusConnectionGetType             func() O.Type
	DBusConnectionNew                 func(d *IOStream, guid string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusConnectionNewFinish           func(res *AsyncResult, err **T.GError) *DBusConnection
	DBusConnectionNewForAddress       func(address string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusConnectionNewForAddressFinish func(res *AsyncResult, err **T.GError) *DBusConnection
	DBusConnectionNewForAddressSync   func(address string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable, err **T.GError) *DBusConnection
	DBusConnectionNewSync             func(d *IOStream, guid string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable, err **T.GError) *DBusConnection

	DBusConnectionAddFilter                  func(d *DBusConnection, filterFunction DBusMessageFilterFunction, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint
	DBusConnectionCall                       func(d *DBusConnection, busName, objectPath, interfaceName, methodName string, parameters *L.Variant, replyType *L.VariantType, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusConnectionCallFinish                 func(d *DBusConnection, res *AsyncResult, err **T.GError) *L.Variant
	DBusConnectionCallSync                   func(d *DBusConnection, busName, objectPath, interfaceName, methodName string, parameters *L.Variant, replyType *L.VariantType, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, err **T.GError) *L.Variant
	DBusConnectionClose                      func(d *DBusConnection, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusConnectionCloseFinish                func(d *DBusConnection, res *AsyncResult, err **T.GError) bool
	DBusConnectionCloseSync                  func(d *DBusConnection, cancellable *Cancellable, err **T.GError) bool
	DBusConnectionEmitSignal                 func(d *DBusConnection, destinationBusName, objectPath, interfaceName string, signalName string, parameters *L.Variant, err **T.GError) bool
	DBusConnectionFlush                      func(d *DBusConnection, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusConnectionFlushFinish                func(d *DBusConnection, res *AsyncResult, err **T.GError) bool
	DBusConnectionFlushSync                  func(d *DBusConnection, cancellable *Cancellable, err **T.GError) bool
	DBusConnectionGetCapabilities            func(d *DBusConnection) DBusCapabilityFlags
	DBusConnectionGetExitOnClose             func(d *DBusConnection) bool
	DBusConnectionGetGuid                    func(d *DBusConnection) string
	DBusConnectionGetPeerCredentials         func(d *DBusConnection) *Credentials
	DBusConnectionGetStream                  func(d *DBusConnection) *IOStream
	DBusConnectionGetUniqueName              func(d *DBusConnection) string
	DBusConnectionIsClosed                   func(d *DBusConnection) bool
	DBusConnectionRegisterObject             func(d *DBusConnection, objectPath string, interfaceInfo *DBusInterfaceInfo, vtable *DBusInterfaceVTable, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify, err **T.GError) uint
	DBusConnectionRegisterSubtree            func(d *DBusConnection, objectPath string, vtable *DBusSubtreeVTable, flags DBusSubtreeFlags, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify, err **T.GError) uint
	DBusConnectionRemoveFilter               func(d *DBusConnection, filterId uint)
	DBusConnectionSendMessage                func(d *DBusConnection, message *DBusMessage, flags DBusSendMessageFlags, outSerial *T.GUint32, err **T.GError) bool
	DBusConnectionSendMessageWithReply       func(d *DBusConnection, message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int, outSerial *T.GUint32, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusConnectionSendMessageWithReplyFinish func(d *DBusConnection, res *AsyncResult, err **T.GError) *DBusMessage
	DBusConnectionSendMessageWithReplySync   func(d *DBusConnection, message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int, outSerial *T.GUint32, cancellable *Cancellable, err **T.GError) *DBusMessage
	DBusConnectionSetExitOnClose             func(d *DBusConnection, exitOnClose bool)
	DBusConnectionSignalSubscribe            func(d *DBusConnection, sender, interfaceName, member, objectPath, arg0 string, flags DBusSignalFlags, callback DBusSignalCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint
	DBusConnectionSignalUnsubscribe          func(d *DBusConnection, subscriptionId uint)
	DBusConnectionStartMessageProcessing     func(d *DBusConnection)
	DBusConnectionUnregisterObject           func(d *DBusConnection, registrationId uint) bool
	DBusConnectionUnregisterSubtree          func(d *DBusConnection, registrationId uint) bool
)

func (d *DBusConnection) AddFilter(filterFunction DBusMessageFilterFunction, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint {
	return DBusConnectionAddFilter(d, filterFunction, userData, userDataFreeFunc)
}
func (d *DBusConnection) Call(busName, objectPath, interfaceName, methodName string, parameters *L.Variant, replyType *L.VariantType, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DBusConnectionCall(d, busName, objectPath, interfaceName, methodName, parameters, replyType, flags, timeoutMsec, cancellable, callback, userData)
}
func (d *DBusConnection) CallFinish(res *AsyncResult, err **T.GError) *L.Variant {
	return DBusConnectionCallFinish(d, res, err)
}
func (d *DBusConnection) CallSync(busName, objectPath, interfaceName, methodName string, parameters *L.Variant, replyType *L.VariantType, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, err **T.GError) *L.Variant {
	return DBusConnectionCallSync(d, busName, objectPath, interfaceName, methodName, parameters, replyType, flags, timeoutMsec, cancellable, err)
}
func (d *DBusConnection) Close(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DBusConnectionClose(d, cancellable, callback, userData)
}
func (d *DBusConnection) CloseFinish(res *AsyncResult, err **T.GError) bool {
	return DBusConnectionCloseFinish(d, res, err)
}
func (d *DBusConnection) CloseSync(cancellable *Cancellable, err **T.GError) bool {
	return DBusConnectionCloseSync(d, cancellable, err)
}
func (d *DBusConnection) EmitSignal(destinationBusName, objectPath, interfaceName string, signalName string, parameters *L.Variant, err **T.GError) bool {
	return DBusConnectionEmitSignal(d, destinationBusName, objectPath, interfaceName, signalName, parameters, err)
}
func (d *DBusConnection) Flush(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DBusConnectionFlush(d, cancellable, callback, userData)
}
func (d *DBusConnection) FlushFinish(res *AsyncResult, err **T.GError) bool {
	return DBusConnectionFlushFinish(d, res, err)
}
func (d *DBusConnection) FlushSync(cancellable *Cancellable, err **T.GError) bool {
	return DBusConnectionFlushSync(d, cancellable, err)
}
func (d *DBusConnection) GetCapabilities() DBusCapabilityFlags {
	return DBusConnectionGetCapabilities(d)
}
func (d *DBusConnection) GetExitOnClose() bool             { return DBusConnectionGetExitOnClose(d) }
func (d *DBusConnection) GetGuid() string                  { return DBusConnectionGetGuid(d) }
func (d *DBusConnection) GetPeerCredentials() *Credentials { return DBusConnectionGetPeerCredentials(d) }
func (d *DBusConnection) GetStream() *IOStream             { return DBusConnectionGetStream(d) }
func (d *DBusConnection) GetUniqueName() string            { return DBusConnectionGetUniqueName(d) }
func (d *DBusConnection) IsClosed() bool                   { return DBusConnectionIsClosed(d) }
func (d *DBusConnection) RegisterObject(objectPath string, interfaceInfo *DBusInterfaceInfo, vtable *DBusInterfaceVTable, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify, err **T.GError) uint {
	return DBusConnectionRegisterObject(d, objectPath, interfaceInfo, vtable, userData, userDataFreeFunc, err)
}
func (d *DBusConnection) RegisterSubtree(objectPath string, vtable *DBusSubtreeVTable, flags DBusSubtreeFlags, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify, err **T.GError) uint {
	return DBusConnectionRegisterSubtree(d, objectPath, vtable, flags, userData, userDataFreeFunc, err)
}
func (d *DBusConnection) RemoveFilter(filterId uint) { DBusConnectionRemoveFilter(d, filterId) }
func (d *DBusConnection) SendMessage(message *DBusMessage, flags DBusSendMessageFlags, outSerial *T.GUint32, err **T.GError) bool {
	return DBusConnectionSendMessage(d, message, flags, outSerial, err)
}
func (d *DBusConnection) SendMessageWithReply(message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int, outSerial *T.GUint32, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DBusConnectionSendMessageWithReply(d, message, flags, timeoutMsec, outSerial, cancellable, callback, userData)
}
func (d *DBusConnection) SendMessageWithReplyFinish(res *AsyncResult, err **T.GError) *DBusMessage {
	return DBusConnectionSendMessageWithReplyFinish(d, res, err)
}
func (d *DBusConnection) SendMessageWithReplySync(message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int, outSerial *T.GUint32, cancellable *Cancellable, err **T.GError) *DBusMessage {
	return DBusConnectionSendMessageWithReplySync(d, message, flags, timeoutMsec, outSerial, cancellable, err)
}
func (d *DBusConnection) SetExitOnClose(exitOnClose bool) {
	DBusConnectionSetExitOnClose(d, exitOnClose)
}
func (d *DBusConnection) SignalSubscribe(sender, interfaceName, member, objectPath, arg0 string, flags DBusSignalFlags, callback DBusSignalCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint {
	return DBusConnectionSignalSubscribe(d, sender, interfaceName, member, objectPath, arg0, flags, callback, userData, userDataFreeFunc)
}
func (d *DBusConnection) SignalUnsubscribe(subscriptionId uint) {
	DBusConnectionSignalUnsubscribe(d, subscriptionId)
}
func (d *DBusConnection) StartMessageProcessing() { DBusConnectionStartMessageProcessing(d) }
func (d *DBusConnection) UnregisterObject(registrationId uint) bool {
	return DBusConnectionUnregisterObject(d, registrationId)
}
func (d *DBusConnection) UnregisterSubtree(registrationId uint) bool {
	return DBusConnectionUnregisterSubtree(d, registrationId)
}

type DBusConnectionFlags Enum

const (
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT DBusConnectionFlags = 1 << iota
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS
	DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION
	DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING
	DBUS_CONNECTION_FLAGS_NONE DBusConnectionFlags = 0
)

var DBusConnectionFlagsGetType func() O.Type

var DBusErrorGetType func() O.Type

type DBusInterfaceInfo struct {
	RefCount    int
	Name        *T.Gchar
	Methods     **DBusMethodInfo
	Signals     **DBusSignalInfo
	Properties  **DBusPropertyInfo
	Annotations **DBusAnnotationInfo
}

type DBusInterfaceGetPropertyFunc func(connection *DBusConnection,
	sender, objectPath, interfaceName, propertyName string,
	err **T.GError, userData T.Gpointer)

type DBusInterfaceMethodCallFunc func(connection *DBusConnection,
	sender, objectPath, interfaceName, methodName string,
	parameters *L.Variant, invocation *DBusMethodInvocation,
	userData T.Gpointer)

type DBusInterfaceSetPropertyFunc func(connection *DBusConnection,
	sender, objectPath, interfaceName, propertyName string,
	value *L.Variant, err **T.GError, userData T.Gpointer)

type DBusInterfaceVTable struct {
	MethodCall  DBusInterfaceMethodCallFunc
	GetProperty DBusInterfaceGetPropertyFunc
	SetProperty DBusInterfaceSetPropertyFunc
	Padding     [8]T.Gpointer
}

type DBusErrorEntry struct {
	ErrorCode     int
	DbusErrorName *T.Gchar
}

var (
	DBusErrorEncodeGerror        func(err *T.GError) string
	DBusErrorGetRemoteError      func(err *T.GError) string
	DBusErrorIsRemoteError       func(err *T.GError) bool
	DBusErrorNewForDBusError     func(dbusErrorName string, dbusErrorMessage string) *T.GError
	DBusErrorQuark               func() L.Quark
	DBusErrorRegisterError       func(errorDomain L.Quark, errorCode int, dbusErrorName string) bool
	DBusErrorRegisterErrorDomain func(errorDomainQuarkName string, quarkVolatile *T.Gsize, entries *DBusErrorEntry, numEntries uint)
	DBusErrorSetDBusError        func(e **T.GError, dbusErrorName, dbusErrorMessage, format string, v ...VArg)
	DBusErrorSetDBusErrorValist  func(err **T.GError, dbusErrorName string, dbusErrorMessage string, format string, varArgs T.VaList)
	DBusErrorStripRemoteError    func(err *T.GError) bool
	DBusErrorUnregisterError     func(errorDomain L.Quark, errorCode int, dbusErrorName string) bool
)

var DBusGenerateGuid func() string

var (
	DBusInterfaceInfoGetType func() O.Type

	DBusInterfaceInfoGenerateXml    func(d *DBusInterfaceInfo, indent uint, stringBuilder *L.String)
	DBusInterfaceInfoLookupMethod   func(d *DBusInterfaceInfo, name string) *DBusMethodInfo
	DBusInterfaceInfoLookupProperty func(d *DBusInterfaceInfo, name string) *DBusPropertyInfo
	DBusInterfaceInfoLookupSignal   func(d *DBusInterfaceInfo, name string) *DBusSignalInfo
	DBusInterfaceInfoRef            func(d *DBusInterfaceInfo) *DBusInterfaceInfo
	DBusInterfaceInfoUnref          func(d *DBusInterfaceInfo)
)

func (d *DBusInterfaceInfo) GenerateXml(indent uint, stringBuilder *L.String) {
	DBusInterfaceInfoGenerateXml(d, indent, stringBuilder)
}
func (d *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	return DBusInterfaceInfoLookupMethod(d, name)
}
func (d *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	return DBusInterfaceInfoLookupProperty(d, name)
}
func (d *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	return DBusInterfaceInfoLookupSignal(d, name)
}
func (d *DBusInterfaceInfo) Ref() *DBusInterfaceInfo { return DBusInterfaceInfoRef(d) }
func (d *DBusInterfaceInfo) Unref()                  { DBusInterfaceInfoUnref(d) }

var (
	DBusIsAddress          func(string string) bool
	DBusIsGuid             func(string string) bool
	DBusIsInterfaceName    func(string string) bool
	DBusIsMemberName       func(string string) bool
	DBusIsName             func(string string) bool
	DBusIsSupportedAddress func(string string, err **T.GError) bool
	DBusIsUniqueName       func(string string) bool
)

type DBusMessage struct{}

var (
	DBusMessageGetType       func() O.Type
	DBusMessageNew           func() *DBusMessage
	DBusMessageNewFromBlob   func(blob *T.Guchar, blobLen T.Gsize, capabilities DBusCapabilityFlags, err **T.GError) *DBusMessage
	DBusMessageNewMethodCall func(name string, path string, interface_ string, method string) *DBusMessage
	DBusMessageNewSignal     func(path string, interface_ string, signal string) *DBusMessage

	DBusMessageBytesNeeded func(blob *T.Guchar, blobLen T.Gsize, err **T.GError) T.Gssize

	DBusMessageCopy                  func(d *DBusMessage, err **T.GError) *DBusMessage
	DBusMessageGetArg0               func(d *DBusMessage) string
	DBusMessageGetBody               func(d *DBusMessage) *L.Variant
	DBusMessageGetByteOrder          func(d *DBusMessage) DBusMessageByteOrder
	DBusMessageGetDestination        func(d *DBusMessage) string
	DBusMessageGetErrorName          func(d *DBusMessage) string
	DBusMessageGetFlags              func(d *DBusMessage) DBusMessageFlags
	DBusMessageGetHeader             func(d *DBusMessage, headerField DBusMessageHeaderField) *L.Variant
	DBusMessageGetHeaderFields       func(d *DBusMessage) *T.Guchar
	DBusMessageGetInterface          func(d *DBusMessage) string
	DBusMessageGetLocked             func(d *DBusMessage) bool
	DBusMessageGetMember             func(d *DBusMessage) string
	DBusMessageGetMessageType        func(d *DBusMessage) DBusMessageType
	DBusMessageGetNumUnixFds         func(d *DBusMessage) T.GUint32
	DBusMessageGetPath               func(d *DBusMessage) string
	DBusMessageGetReplySerial        func(d *DBusMessage) T.GUint32
	DBusMessageGetSender             func(d *DBusMessage) string
	DBusMessageGetSerial             func(d *DBusMessage) T.GUint32
	DBusMessageGetSignature          func(d *DBusMessage) string
	DBusMessageGetUnixFdList         func(d *DBusMessage) *T.GUnixFDList
	DBusMessageLock                  func(d *DBusMessage)
	DBusMessageNewMethodError        func(d *DBusMessage, errorName, errorMessageFormat string, v ...VArg) *DBusMessage
	DBusMessageNewMethodErrorLiteral func(d *DBusMessage, errorName, errorMessage string) *DBusMessage
	DBusMessageNewMethodErrorValist  func(d *DBusMessage, errorName, errorMessageFormat string, varArgs T.VaList) *DBusMessage
	DBusMessageNewMethodReply        func(d *DBusMessage) *DBusMessage
	DBusMessagePrint                 func(d *DBusMessage, indent uint) string
	DBusMessageSetBody               func(d *DBusMessage, body *L.Variant)
	DBusMessageSetByteOrder          func(d *DBusMessage, byteOrder DBusMessageByteOrder)
	DBusMessageSetDestination        func(d *DBusMessage, value string)
	DBusMessageSetErrorName          func(d *DBusMessage, value string)
	DBusMessageSetFlags              func(d *DBusMessage, flags DBusMessageFlags)
	DBusMessageSetHeader             func(d *DBusMessage, headerField DBusMessageHeaderField, value *L.Variant)
	DBusMessageSetInterface          func(d *DBusMessage, value string)
	DBusMessageSetMember             func(d *DBusMessage, value string)
	DBusMessageSetMessageType        func(d *DBusMessage, typ DBusMessageType)
	DBusMessageSetNumUnixFds         func(d *DBusMessage, value T.GUint32)
	DBusMessageSetPath               func(d *DBusMessage, value string)
	DBusMessageSetReplySerial        func(d *DBusMessage, value T.GUint32)
	DBusMessageSetSender             func(d *DBusMessage, value string)
	DBusMessageSetSerial             func(d *DBusMessage, serial T.GUint32)
	DBusMessageSetSignature          func(d *DBusMessage, value string)
	DBusMessageSetUnixFdList         func(d *DBusMessage, fdList *T.GUnixFDList)
	DBusMessageToBlob                func(d *DBusMessage, outSize *T.Gsize, capabilities DBusCapabilityFlags, err **T.GError) *T.Guchar
	DBusMessageToGerror              func(d *DBusMessage, err **T.GError) bool
)

func (d *DBusMessage) Copy(err **T.GError) *DBusMessage   { return DBusMessageCopy(d, err) }
func (d *DBusMessage) GetArg0() string                    { return DBusMessageGetArg0(d) }
func (d *DBusMessage) GetBody() *L.Variant                { return DBusMessageGetBody(d) }
func (d *DBusMessage) GetByteOrder() DBusMessageByteOrder { return DBusMessageGetByteOrder(d) }
func (d *DBusMessage) GetDestination() string             { return DBusMessageGetDestination(d) }
func (d *DBusMessage) GetErrorName() string               { return DBusMessageGetErrorName(d) }
func (d *DBusMessage) GetFlags() DBusMessageFlags         { return DBusMessageGetFlags(d) }
func (d *DBusMessage) GetHeader(headerField DBusMessageHeaderField) *L.Variant {
	return DBusMessageGetHeader(d, headerField)
}
func (d *DBusMessage) GetHeaderFields() *T.Guchar      { return DBusMessageGetHeaderFields(d) }
func (d *DBusMessage) GetInterface() string            { return DBusMessageGetInterface(d) }
func (d *DBusMessage) GetLocked() bool                 { return DBusMessageGetLocked(d) }
func (d *DBusMessage) GetMember() string               { return DBusMessageGetMember(d) }
func (d *DBusMessage) GetMessageType() DBusMessageType { return DBusMessageGetMessageType(d) }
func (d *DBusMessage) GetNumUnixFds() T.GUint32        { return DBusMessageGetNumUnixFds(d) }
func (d *DBusMessage) GetPath() string                 { return DBusMessageGetPath(d) }
func (d *DBusMessage) GetReplySerial() T.GUint32       { return DBusMessageGetReplySerial(d) }
func (d *DBusMessage) GetSender() string               { return DBusMessageGetSender(d) }
func (d *DBusMessage) GetSerial() T.GUint32            { return DBusMessageGetSerial(d) }
func (d *DBusMessage) GetSignature() string            { return DBusMessageGetSignature(d) }
func (d *DBusMessage) GetUnixFdList() *T.GUnixFDList   { return DBusMessageGetUnixFdList(d) }
func (d *DBusMessage) Lock()                           { DBusMessageLock(d) }
func (d *DBusMessage) NewMethodError(errorName, errorMessageFormat string, v ...VArg) *DBusMessage {
	return DBusMessageNewMethodError(d, errorName, errorMessageFormat, v)
}
func (d *DBusMessage) NewMethodErrorLiteral(errorName, errorMessage string) *DBusMessage {
	return DBusMessageNewMethodErrorLiteral(d, errorName, errorMessage)
}
func (d *DBusMessage) NewMethodErrorValist(errorName, errorMessageFormat string, varArgs T.VaList) *DBusMessage {
	return DBusMessageNewMethodErrorValist(d, errorName, errorMessageFormat, varArgs)
}
func (d *DBusMessage) NewMethodReply() *DBusMessage { return DBusMessageNewMethodReply(d) }
func (d *DBusMessage) Print(indent uint) string     { return DBusMessagePrint(d, indent) }
func (d *DBusMessage) SetBody(body *L.Variant)      { DBusMessageSetBody(d, body) }
func (d *DBusMessage) SetByteOrder(byteOrder DBusMessageByteOrder) {
	DBusMessageSetByteOrder(d, byteOrder)
}
func (d *DBusMessage) SetDestination(value string)     { DBusMessageSetDestination(d, value) }
func (d *DBusMessage) SetErrorName(value string)       { DBusMessageSetErrorName(d, value) }
func (d *DBusMessage) SetFlags(flags DBusMessageFlags) { DBusMessageSetFlags(d, flags) }
func (d *DBusMessage) SetHeader(headerField DBusMessageHeaderField, value *L.Variant) {
	DBusMessageSetHeader(d, headerField, value)
}
func (d *DBusMessage) SetInterface(value string)           { DBusMessageSetInterface(d, value) }
func (d *DBusMessage) SetMember(value string)              { DBusMessageSetMember(d, value) }
func (d *DBusMessage) SetMessageType(typ DBusMessageType)  { DBusMessageSetMessageType(d, typ) }
func (d *DBusMessage) SetNumUnixFds(value T.GUint32)       { DBusMessageSetNumUnixFds(d, value) }
func (d *DBusMessage) SetPath(value string)                { DBusMessageSetPath(d, value) }
func (d *DBusMessage) SetReplySerial(value T.GUint32)      { DBusMessageSetReplySerial(d, value) }
func (d *DBusMessage) SetSender(value string)              { DBusMessageSetSender(d, value) }
func (d *DBusMessage) SetSerial(serial T.GUint32)          { DBusMessageSetSerial(d, serial) }
func (d *DBusMessage) SetSignature(value string)           { DBusMessageSetSignature(d, value) }
func (d *DBusMessage) SetUnixFdList(fdList *T.GUnixFDList) { DBusMessageSetUnixFdList(d, fdList) }
func (d *DBusMessage) ToBlob(outSize *T.Gsize, capabilities DBusCapabilityFlags, err **T.GError) *T.Guchar {
	return DBusMessageToBlob(d, outSize, capabilities, err)
}
func (d *DBusMessage) ToGerror(err **T.GError) bool { return DBusMessageToGerror(d, err) }

type DBusMessageByteOrder Enum

const (
	DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN    DBusMessageByteOrder = 'B'
	DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN DBusMessageByteOrder = 'l'
)

var DBusMessageByteOrderGetType func() O.Type

type DBusMessageFilterFunction func(
	connection *DBusConnection,
	message *DBusMessage,
	incoming bool,
	userData T.Gpointer) *DBusMessage

type DBusMessageFlags Enum

const (
	DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED DBusMessageFlags = 1 << iota
	DBUS_MESSAGE_FLAGS_NO_AUTO_START
	DBUS_MESSAGE_FLAGS_NONE DBusMessageFlags = 0
)

var DBusMessageFlagsGetType func() O.Type

type DBusMessageHeaderField Enum

const (
	DBUS_MESSAGE_HEADER_FIELD_INVALID DBusMessageHeaderField = iota
	DBUS_MESSAGE_HEADER_FIELD_PATH
	DBUS_MESSAGE_HEADER_FIELD_INTERFACE
	DBUS_MESSAGE_HEADER_FIELD_MEMBER
	DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME
	DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL
	DBUS_MESSAGE_HEADER_FIELD_DESTINATION
	DBUS_MESSAGE_HEADER_FIELD_SENDER
	DBUS_MESSAGE_HEADER_FIELD_SIGNATURE
	DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS
)

var DBusMessageHeaderFieldGetType func() O.Type

type DBusMessageType Enum

const (
	DBUS_MESSAGE_TYPE_INVALID DBusMessageType = iota
	DBUS_MESSAGE_TYPE_METHOD_CALL
	DBUS_MESSAGE_TYPE_METHOD_RETURN
	DBUS_MESSAGE_TYPE_ERROR
	DBUS_MESSAGE_TYPE_SIGNAL
)

var DBusMessageTypeGetType func() O.Type

type DBusMethodInfo struct {
	RefCount    int
	Name        *T.Gchar
	InArgs      **DBusArgInfo
	OutArgs     **DBusArgInfo
	Annotations **DBusAnnotationInfo
}

var (
	DBusMethodInfoGetType func() O.Type

	DBusMethodInfoRef   func(d *DBusMethodInfo) *DBusMethodInfo
	DBusMethodInfoUnref func(d *DBusMethodInfo)
)

func (d *DBusMethodInfo) Ref() *DBusMethodInfo { return DBusMethodInfoRef(d) }
func (d *DBusMethodInfo) Unref()               { DBusMethodInfoUnref(d) }

type DBusMethodInvocation struct{}

var (
	DBusMethodInvocationGetType func() O.Type

	DBusMethodInvocationGetConnection      func(d *DBusMethodInvocation) *DBusConnection
	DBusMethodInvocationGetInterfaceName   func(d *DBusMethodInvocation) string
	DBusMethodInvocationGetMessage         func(d *DBusMethodInvocation) *DBusMessage
	DBusMethodInvocationGetMethodInfo      func(d *DBusMethodInvocation) *DBusMethodInfo
	DBusMethodInvocationGetMethodName      func(d *DBusMethodInvocation) string
	DBusMethodInvocationGetObjectPath      func(d *DBusMethodInvocation) string
	DBusMethodInvocationGetParameters      func(d *DBusMethodInvocation) *L.Variant
	DBusMethodInvocationGetSender          func(d *DBusMethodInvocation) string
	DBusMethodInvocationGetUserData        func(d *DBusMethodInvocation) T.Gpointer
	DBusMethodInvocationReturnDBusError    func(d *DBusMethodInvocation, errorName, errorMessage string)
	DBusMethodInvocationReturnError        func(d *DBusMethodInvocation, domain L.Quark, code int, format string, v ...VArg)
	DBusMethodInvocationReturnErrorLiteral func(d *DBusMethodInvocation, domain L.Quark, code int, message string)
	DBusMethodInvocationReturnErrorValist  func(d *DBusMethodInvocation, domain L.Quark, code int, format string, varArgs T.VaList)
	DBusMethodInvocationReturnGerror       func(d *DBusMethodInvocation, err *T.GError)
	DBusMethodInvocationReturnValue        func(d *DBusMethodInvocation, parameters *L.Variant)
)

func (d *DBusMethodInvocation) GetConnection() *DBusConnection {
	return DBusMethodInvocationGetConnection(d)
}
func (d *DBusMethodInvocation) GetInterfaceName() string {
	return DBusMethodInvocationGetInterfaceName(d)
}
func (d *DBusMethodInvocation) GetMessage() *DBusMessage { return DBusMethodInvocationGetMessage(d) }
func (d *DBusMethodInvocation) GetMethodInfo() *DBusMethodInfo {
	return DBusMethodInvocationGetMethodInfo(d)
}
func (d *DBusMethodInvocation) GetMethodName() string { return DBusMethodInvocationGetMethodName(d) }
func (d *DBusMethodInvocation) GetObjectPath() string { return DBusMethodInvocationGetObjectPath(d) }
func (d *DBusMethodInvocation) GetParameters() *L.Variant {
	return DBusMethodInvocationGetParameters(d)
}
func (d *DBusMethodInvocation) GetSender() string       { return DBusMethodInvocationGetSender(d) }
func (d *DBusMethodInvocation) GetUserData() T.Gpointer { return DBusMethodInvocationGetUserData(d) }
func (d *DBusMethodInvocation) ReturnDBusError(errorName, errorMessage string) {
	DBusMethodInvocationReturnDBusError(d, errorName, errorMessage)
}
func (d *DBusMethodInvocation) ReturnError(domain L.Quark, code int, format string, v ...VArg) {
	DBusMethodInvocationReturnError(d, domain, code, format, v)
}
func (d *DBusMethodInvocation) ReturnErrorLiteral(domain L.Quark, code int, message string) {
	DBusMethodInvocationReturnErrorLiteral(d, domain, code, message)
}
func (d *DBusMethodInvocation) ReturnErrorValist(domain L.Quark, code int, format string, varArgs T.VaList) {
	DBusMethodInvocationReturnErrorValist(d, domain, code, format, varArgs)
}
func (d *DBusMethodInvocation) ReturnGerror(err *T.GError) { DBusMethodInvocationReturnGerror(d, err) }
func (d *DBusMethodInvocation) ReturnValue(parameters *L.Variant) {
	DBusMethodInvocationReturnValue(d, parameters)
}

type DBusNodeInfo struct {
	Ref_count   int
	Path        *T.Gchar
	Interfaces  **DBusInterfaceInfo
	Nodes       **DBusNodeInfo
	Annotations **DBusAnnotationInfo
}

var (
	DBusNodeInfoGetType   func() O.Type
	DBusNodeInfoNewForXml func(xmlData string, err **T.GError) *DBusNodeInfo

	DBusNodeInfoGenerateXml     func(d *DBusNodeInfo, indent uint, stringBuilder *L.String)
	DBusNodeInfoLookupInterface func(d *DBusNodeInfo, name string) *DBusInterfaceInfo
	DBusNodeInfoRef             func(d *DBusNodeInfo) *DBusNodeInfo
	DBusNodeInfoUnref           func(d *DBusNodeInfo)
)

func (d *DBusNodeInfo) GenerateXml(indent uint, stringBuilder *L.String) {
	DBusNodeInfoGenerateXml(d, indent, stringBuilder)
}
func (d *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	return DBusNodeInfoLookupInterface(d, name)
}
func (d *DBusNodeInfo) Ref() *DBusNodeInfo { return DBusNodeInfoRef(d) }
func (d *DBusNodeInfo) Unref()             { DBusNodeInfoUnref(d) }

type DBusPropertyInfo struct {
	RefCount    int
	Name        *T.Gchar
	Signature   *T.Gchar
	Flags       DBusPropertyInfoFlags
	Annotations **DBusAnnotationInfo
}

var (
	DBusPropertyInfoGetType func() O.Type

	DBusPropertyInfoRef   func(d *DBusPropertyInfo) *DBusPropertyInfo
	DBusPropertyInfoUnref func(d *DBusPropertyInfo)
)

func (d *DBusPropertyInfo) Ref() *DBusPropertyInfo { return DBusPropertyInfoRef(d) }
func (d *DBusPropertyInfo) Unref()                 { DBusPropertyInfoUnref(d) }

type DBusPropertyInfoFlags Enum

const (
	DBUS_PROPERTY_INFO_FLAGS_READABLE DBusPropertyInfoFlags = 1 << iota
	DBUS_PROPERTY_INFO_FLAGS_WRITABLE
	DBUS_PROPERTY_INFO_FLAGS_NONE DBusPropertyInfoFlags = 0
)

var DBusPropertyInfoFlagsGetType func() O.Type

type DBusProxy struct {
	Parent O.Object
	Priv   *struct{}
}

var (
	DBusProxyGetType         func() O.Type
	DBusProxyNew             func(d *DBusConnection, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusProxyNewFinish       func(res *AsyncResult, err **T.GError) *DBusProxy
	DBusProxyNewForBus       func(busType BusType, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusProxyNewForBusFinish func(res *AsyncResult, err **T.GError) *DBusProxy
	DBusProxyNewForBusSync   func(busType BusType, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable, err **T.GError) *DBusProxy
	DBusProxyNewSync         func(d *DBusConnection, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable, err **T.GError) *DBusProxy

	DBusProxyCall                   func(d *DBusProxy, methodName string, parameters *L.Variant, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DBusProxyCallFinish             func(d *DBusProxy, res *AsyncResult, err **T.GError) *L.Variant
	DBusProxyCallSync               func(d *DBusProxy, methodName string, parameters *L.Variant, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, err **T.GError) *L.Variant
	DBusProxyGetCachedProperty      func(d *DBusProxy, propertyName string) *L.Variant
	DBusProxyGetCachedPropertyNames func(d *DBusProxy) []string
	DBusProxyGetConnection          func(d *DBusProxy) *DBusConnection
	DBusProxyGetDefaultTimeout      func(d *DBusProxy) int
	DBusProxyGetFlags               func(d *DBusProxy) DBusProxyFlags
	DBusProxyGetInterfaceInfo       func(d *DBusProxy) *DBusInterfaceInfo
	DBusProxyGetInterfaceName       func(d *DBusProxy) string
	DBusProxyGetName                func(d *DBusProxy) string
	DBusProxyGetNameOwner           func(d *DBusProxy) string
	DBusProxyGetObjectPath          func(d *DBusProxy) string
	DBusProxySetCachedProperty      func(d *DBusProxy, propertyName string, value *L.Variant)
	DBusProxySetDefaultTimeout      func(d *DBusProxy, timeoutMsec int)
	DBusProxySetInterfaceInfo       func(d *DBusProxy, info *DBusInterfaceInfo)
)

func (d *DBusProxy) Call(methodName string, parameters *L.Variant, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DBusProxyCall(d, methodName, parameters, flags, timeoutMsec, cancellable, callback, userData)
}
func (d *DBusProxy) CallFinish(res *AsyncResult, err **T.GError) *L.Variant {
	return DBusProxyCallFinish(d, res, err)
}
func (d *DBusProxy) CallSync(methodName string, parameters *L.Variant, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, err **T.GError) *L.Variant {
	return DBusProxyCallSync(d, methodName, parameters, flags, timeoutMsec, cancellable, err)
}
func (d *DBusProxy) GetCachedProperty(propertyName string) *L.Variant {
	return DBusProxyGetCachedProperty(d, propertyName)
}
func (d *DBusProxy) GetCachedPropertyNames() []string     { return DBusProxyGetCachedPropertyNames(d) }
func (d *DBusProxy) GetConnection() *DBusConnection       { return DBusProxyGetConnection(d) }
func (d *DBusProxy) GetDefaultTimeout() int               { return DBusProxyGetDefaultTimeout(d) }
func (d *DBusProxy) GetFlags() DBusProxyFlags             { return DBusProxyGetFlags(d) }
func (d *DBusProxy) GetInterfaceInfo() *DBusInterfaceInfo { return DBusProxyGetInterfaceInfo(d) }
func (d *DBusProxy) GetInterfaceName() string             { return DBusProxyGetInterfaceName(d) }
func (d *DBusProxy) GetName() string                      { return DBusProxyGetName(d) }
func (d *DBusProxy) GetNameOwner() string                 { return DBusProxyGetNameOwner(d) }
func (d *DBusProxy) GetObjectPath() string                { return DBusProxyGetObjectPath(d) }
func (d *DBusProxy) SetCachedProperty(propertyName string, value *L.Variant) {
	DBusProxySetCachedProperty(d, propertyName, value)
}
func (d *DBusProxy) SetDefaultTimeout(timeoutMsec int)        { DBusProxySetDefaultTimeout(d, timeoutMsec) }
func (d *DBusProxy) SetInterfaceInfo(info *DBusInterfaceInfo) { DBusProxySetInterfaceInfo(d, info) }

type DBusProxyFlags Enum

const (
	DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES DBusProxyFlags = 1 << iota
	DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS
	DBUS_PROXY_FLAGS_DO_NOT_AUTO_START
	DBUS_PROXY_FLAGS_NONE DBusProxyFlags = 0
)

var DBusProxyFlagsGetType func() O.Type

type DBusSendMessageFlags Enum

const (
	DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL DBusSendMessageFlags = 1 << iota
	DBUS_SEND_MESSAGE_FLAGS_NONE            DBusSendMessageFlags = 0
)

var DBusSendMessageFlagsGetType func() O.Type

type DBusServer struct{}

var (
	DBusServerGetType func() O.Type
	DBusServerNewSync func(address string, flags DBusServerFlags, guid string, observer *DBusAuthObserver, cancellable *Cancellable, err **T.GError) *DBusServer

	DBusServerGetClientAddress func(d *DBusServer) string
	DBusServerGetFlags         func(d *DBusServer) DBusServerFlags
	DBusServerGetGuid          func(d *DBusServer) string
	DBusServerIsActive         func(d *DBusServer) bool
	DBusServerStart            func(d *DBusServer)
	DBusServerStop             func(d *DBusServer)
)

func (d *DBusServer) GetClientAddress() string  { return DBusServerGetClientAddress(d) }
func (d *DBusServer) GetFlags() DBusServerFlags { return DBusServerGetFlags(d) }
func (d *DBusServer) GetGuid() string           { return DBusServerGetGuid(d) }
func (d *DBusServer) IsActive() bool            { return DBusServerIsActive(d) }
func (d *DBusServer) Start()                    { DBusServerStart(d) }
func (d *DBusServer) Stop()                     { DBusServerStop(d) }

type DBusServerFlags Enum

const (
	DBUS_SERVER_FLAGS_RUN_IN_THREAD DBusServerFlags = 1 << iota
	DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS
	DBUS_SERVER_FLAGS_NONE DBusServerFlags = 0
)

var DBusServerFlagsGetType func() O.Type

type DBusSignalCallback func(connection *DBusConnection,
	senderName, objectPath, interfaceName, signalName string,
	parameters *L.Variant, userData T.Gpointer)

type DBusSignalInfo struct {
	RefCount    int
	Name        *T.Gchar
	Args        **DBusArgInfo
	Annotations **DBusAnnotationInfo
}

var (
	DBusSignalInfoGetType func() O.Type

	DBusSignalInfoRef   func(d *DBusSignalInfo) *DBusSignalInfo
	DBusSignalInfoUnref func(d *DBusSignalInfo)
)

func (d *DBusSignalInfo) Ref() *DBusSignalInfo { return DBusSignalInfoRef(d) }
func (d *DBusSignalInfo) Unref()               { DBusSignalInfoUnref(d) }

type DBusSignalFlags Enum

const (
	DBUS_SIGNAL_FLAGS_NO_MATCH_RULE DBusSignalFlags = 1 << iota
	DBUS_SIGNAL_FLAGS_NONE          DBusSignalFlags = 0
)

var DBusSignalFlagsGetType func() O.Type

type DBusSubtreeDispatchFunc func(connection *DBusConnection,
	sender, objectPath, interfaceName, node string,
	outUserData *T.Gpointer,
	userData T.Gpointer) *DBusInterfaceVTable

type DBusSubtreeEnumerateFunc func(connection *DBusConnection,
	sender, objectPath string, userData T.Gpointer) []string

type DBusSubtreeFlags Enum

const (
	DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES DBusSubtreeFlags = 1 << iota
	DBUS_SUBTREE_FLAGS_NONE                           DBusSubtreeFlags = 0
)

var DBusSubtreeFlagsGetType func() O.Type

type DBusSubtreeIntrospectFunc func(connection *DBusConnection,
	sender, objectPath, node string,
	userData T.Gpointer) **DBusInterfaceInfo

type DBusSubtreeVTable struct {
	Enumerate  DBusSubtreeEnumerateFunc
	Introspect DBusSubtreeIntrospectFunc
	Dispatch   DBusSubtreeDispatchFunc
	Padding    [8]T.Gpointer
}

type Drive struct{}

var (
	DriveGetType func() O.Type

	DriveCanEject                 func(d *Drive) bool
	DriveCanPollForMedia          func(d *Drive) bool
	DriveCanStart                 func(d *Drive) bool
	DriveCanStartDegraded         func(d *Drive) bool
	DriveCanStop                  func(d *Drive) bool
	DriveEject                    func(d *Drive, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DriveEjectFinish              func(d *Drive, result *AsyncResult, err **T.GError) bool
	DriveEjectWithOperation       func(d *Drive, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DriveEjectWithOperationFinish func(d *Drive, result *AsyncResult, err **T.GError) bool
	DriveEnumerateIdentifiers     func(d *Drive) []string
	DriveGetIcon                  func(d *Drive) *Icon
	DriveGetIdentifier            func(d *Drive, kind string) string
	DriveGetName                  func(d *Drive) string
	DriveGetStartStopType         func(d *Drive) DriveStartStopType
	DriveGetVolumes               func(d *Drive) *T.GList
	DriveHasMedia                 func(d *Drive) bool
	DriveHasVolumes               func(d *Drive) bool
	DriveIsMediaCheckAutomatic    func(d *Drive) bool
	DriveIsMediaRemovable         func(d *Drive) bool
	DrivePollForMedia             func(d *Drive, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DrivePollForMediaFinish       func(d *Drive, result *AsyncResult, err **T.GError) bool
	DriveStart                    func(d *Drive, flags DriveStartFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DriveStartFinish              func(d *Drive, result *AsyncResult, err **T.GError) bool
	DriveStop                     func(d *Drive, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	DriveStopFinish               func(d *Drive, result *AsyncResult, err **T.GError) bool
)

func (d *Drive) CanEject() bool         { return DriveCanEject(d) }
func (d *Drive) CanPollForMedia() bool  { return DriveCanPollForMedia(d) }
func (d *Drive) CanStart() bool         { return DriveCanStart(d) }
func (d *Drive) CanStartDegraded() bool { return DriveCanStartDegraded(d) }
func (d *Drive) CanStop() bool          { return DriveCanStop(d) }
func (d *Drive) Eject(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DriveEject(d, flags, cancellable, callback, userData)
}
func (d *Drive) EjectFinish(result *AsyncResult, err **T.GError) bool {
	return DriveEjectFinish(d, result, err)
}
func (d *Drive) EjectWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DriveEjectWithOperation(d, flags, mountOperation, cancellable, callback, userData)
}
func (d *Drive) EjectWithOperationFinish(result *AsyncResult, err **T.GError) bool {
	return DriveEjectWithOperationFinish(d, result, err)
}
func (d *Drive) EnumerateIdentifiers() []string       { return DriveEnumerateIdentifiers(d) }
func (d *Drive) GetIcon() *Icon                       { return DriveGetIcon(d) }
func (d *Drive) GetIdentifier(kind string) string     { return DriveGetIdentifier(d, kind) }
func (d *Drive) GetName() string                      { return DriveGetName(d) }
func (d *Drive) GetStartStopType() DriveStartStopType { return DriveGetStartStopType(d) }
func (d *Drive) GetVolumes() *T.GList                 { return DriveGetVolumes(d) }
func (d *Drive) HasMedia() bool                       { return DriveHasMedia(d) }
func (d *Drive) HasVolumes() bool                     { return DriveHasVolumes(d) }
func (d *Drive) IsMediaCheckAutomatic() bool          { return DriveIsMediaCheckAutomatic(d) }
func (d *Drive) IsMediaRemovable() bool               { return DriveIsMediaRemovable(d) }
func (d *Drive) PollForMedia(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DrivePollForMedia(d, cancellable, callback, userData)
}
func (d *Drive) PollForMediaFinish(result *AsyncResult, err **T.GError) bool {
	return DrivePollForMediaFinish(d, result, err)
}
func (d *Drive) Start(flags DriveStartFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DriveStart(d, flags, mountOperation, cancellable, callback, userData)
}
func (d *Drive) StartFinish(result *AsyncResult, err **T.GError) bool {
	return DriveStartFinish(d, result, err)
}
func (d *Drive) Stop(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	DriveStop(d, flags, mountOperation, cancellable, callback, userData)
}
func (d *Drive) StopFinish(result *AsyncResult, err **T.GError) bool {
	return DriveStopFinish(d, result, err)
}

type DriveStartFlags Enum

const DRIVE_START_NONE DriveStartFlags = 0

var DriveStartFlagsGetType func() O.Type

type DriveStartStopType Enum

const (
	DRIVE_START_STOP_TYPE_UNKNOWN DriveStartStopType = iota
	DRIVE_START_STOP_TYPE_SHUTDOWN
	DRIVE_START_STOP_TYPE_NETWORK
	DRIVE_START_STOP_TYPE_MULTIDISK
	DRIVE_START_STOP_TYPE_PASSWORD
)

var DriveStartStopTypeGetType func() O.Type
