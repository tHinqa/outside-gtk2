// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
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

	dataInputStreamGetByteOrder    func(d *DataInputStream) DataStreamByteOrder
	dataInputStreamGetNewlineType  func(d *DataInputStream) DataStreamNewlineType
	dataInputStreamReadByte        func(d *DataInputStream, cancellable *Cancellable, err **T.GError) T.Guchar
	dataInputStreamReadInt16       func(d *DataInputStream, cancellable *Cancellable, err **T.GError) int16
	dataInputStreamReadInt32       func(d *DataInputStream, cancellable *Cancellable, err **T.GError) T.GInt32
	dataInputStreamReadInt64       func(d *DataInputStream, cancellable *Cancellable, err **T.GError) int64
	dataInputStreamReadLine        func(d *DataInputStream, length *T.Gsize, cancellable *Cancellable, err **T.GError) string
	dataInputStreamReadLineAsync   func(d *DataInputStream, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	dataInputStreamReadLineFinish  func(d *DataInputStream, result *AsyncResult, length *T.Gsize, err **T.GError) string
	dataInputStreamReadUint16      func(d *DataInputStream, cancellable *Cancellable, err **T.GError) uint16
	dataInputStreamReadUint32      func(d *DataInputStream, cancellable *Cancellable, err **T.GError) T.GUint32
	dataInputStreamReadUint64      func(d *DataInputStream, cancellable *Cancellable, err **T.GError) uint64
	dataInputStreamReadUntil       func(d *DataInputStream, stopChars string, length *T.Gsize, cancellable *Cancellable, err **T.GError) string
	dataInputStreamReadUntilAsync  func(d *DataInputStream, stopChars string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	dataInputStreamReadUntilFinish func(d *DataInputStream, result *AsyncResult, length *T.Gsize, err **T.GError) string
	dataInputStreamReadUpto        func(d *DataInputStream, stopChars string, stopCharsLen T.Gssize, length *T.Gsize, cancellable *Cancellable, err **T.GError) string
	dataInputStreamReadUptoAsync   func(d *DataInputStream, stopChars string, stopCharsLen T.Gssize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	dataInputStreamReadUptoFinish  func(d *DataInputStream, result *AsyncResult, length *T.Gsize, err **T.GError) string
	dataInputStreamSetByteOrder    func(d *DataInputStream, order DataStreamByteOrder)
	dataInputStreamSetNewlineType  func(d *DataInputStream, typ DataStreamNewlineType)
)

func (d *DataInputStream) GetByteOrder() DataStreamByteOrder { return dataInputStreamGetByteOrder(d) }
func (d *DataInputStream) GetNewlineType() DataStreamNewlineType {
	return dataInputStreamGetNewlineType(d)
}
func (d *DataInputStream) ReadByte(cancellable *Cancellable, err **T.GError) T.Guchar {
	return dataInputStreamReadByte(d, cancellable, err)
}
func (d *DataInputStream) ReadInt16(cancellable *Cancellable, err **T.GError) int16 {
	return dataInputStreamReadInt16(d, cancellable, err)
}
func (d *DataInputStream) ReadInt32(cancellable *Cancellable, err **T.GError) T.GInt32 {
	return dataInputStreamReadInt32(d, cancellable, err)
}
func (d *DataInputStream) ReadInt64(cancellable *Cancellable, err **T.GError) int64 {
	return dataInputStreamReadInt64(d, cancellable, err)
}
func (d *DataInputStream) ReadLine(length *T.Gsize, cancellable *Cancellable, err **T.GError) string {
	return dataInputStreamReadLine(d, length, cancellable, err)
}
func (d *DataInputStream) ReadLineAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	dataInputStreamReadLineAsync(d, ioPriority, cancellable, callback, userData)
}
func (d *DataInputStream) ReadLineFinish(result *AsyncResult, length *T.Gsize, err **T.GError) string {
	return dataInputStreamReadLineFinish(d, result, length, err)
}
func (d *DataInputStream) ReadUint16(cancellable *Cancellable, err **T.GError) uint16 {
	return dataInputStreamReadUint16(d, cancellable, err)
}
func (d *DataInputStream) ReadUint32(cancellable *Cancellable, err **T.GError) T.GUint32 {
	return dataInputStreamReadUint32(d, cancellable, err)
}
func (d *DataInputStream) ReadUint64(cancellable *Cancellable, err **T.GError) uint64 {
	return dataInputStreamReadUint64(d, cancellable, err)
}
func (d *DataInputStream) ReadUntil(stopChars string, length *T.Gsize, cancellable *Cancellable, err **T.GError) string {
	return dataInputStreamReadUntil(d, stopChars, length, cancellable, err)
}
func (d *DataInputStream) ReadUntilAsync(stopChars string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	dataInputStreamReadUntilAsync(d, stopChars, ioPriority, cancellable, callback, userData)
}
func (d *DataInputStream) ReadUntilFinish(result *AsyncResult, length *T.Gsize, err **T.GError) string {
	return dataInputStreamReadUntilFinish(d, result, length, err)
}
func (d *DataInputStream) ReadUpto(stopChars string, stopCharsLen T.Gssize, length *T.Gsize, cancellable *Cancellable, err **T.GError) string {
	return dataInputStreamReadUpto(d, stopChars, stopCharsLen, length, cancellable, err)
}
func (d *DataInputStream) ReadUptoAsync(stopChars string, stopCharsLen T.Gssize, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	dataInputStreamReadUptoAsync(d, stopChars, stopCharsLen, ioPriority, cancellable, callback, userData)
}
func (d *DataInputStream) ReadUptoFinish(result *AsyncResult, length *T.Gsize, err **T.GError) string {
	return dataInputStreamReadUptoFinish(d, result, length, err)
}
func (d *DataInputStream) SetByteOrder(order DataStreamByteOrder) {
	dataInputStreamSetByteOrder(d, order)
}
func (d *DataInputStream) SetNewlineType(typ DataStreamNewlineType) {
	dataInputStreamSetNewlineType(d, typ)
}

type DataOutputStream struct {
	Parent FilterOutputStream
	_      *struct{}
}

var (
	DataOutputStreamGetType func() O.Type
	DataOutputStreamNew     func(baseStream *OutputStream) *DataOutputStream

	dataOutputStreamGetByteOrder func(d *DataOutputStream) DataStreamByteOrder
	dataOutputStreamPutByte      func(d *DataOutputStream, data T.Guchar, cancellable *Cancellable, err **T.GError) T.Gboolean
	dataOutputStreamPutInt16     func(d *DataOutputStream, data int16, cancellable *Cancellable, err **T.GError) T.Gboolean
	dataOutputStreamPutInt32     func(d *DataOutputStream, data T.GInt32, cancellable *Cancellable, err **T.GError) T.Gboolean
	dataOutputStreamPutInt64     func(d *DataOutputStream, data int64, cancellable *Cancellable, err **T.GError) T.Gboolean
	dataOutputStreamPutString    func(d *DataOutputStream, str string, cancellable *Cancellable, err **T.GError) T.Gboolean
	dataOutputStreamPutUint16    func(d *DataOutputStream, data uint16, cancellable *Cancellable, err **T.GError) T.Gboolean
	dataOutputStreamPutUint32    func(d *DataOutputStream, data T.GUint32, cancellable *Cancellable, err **T.GError) T.Gboolean
	dataOutputStreamPutUint64    func(d *DataOutputStream, data uint64, cancellable *Cancellable, err **T.GError) T.Gboolean
	dataOutputStreamSetByteOrder func(d *DataOutputStream, order DataStreamByteOrder)
)

func (d *DataOutputStream) GetByteOrder() DataStreamByteOrder { return dataOutputStreamGetByteOrder(d) }
func (d *DataOutputStream) PutByte(data T.Guchar, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dataOutputStreamPutByte(d, data, cancellable, err)
}
func (d *DataOutputStream) PutInt16(data int16, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dataOutputStreamPutInt16(d, data, cancellable, err)
}
func (d *DataOutputStream) PutInt32(data T.GInt32, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dataOutputStreamPutInt32(d, data, cancellable, err)
}
func (d *DataOutputStream) PutInt64(data int64, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dataOutputStreamPutInt64(d, data, cancellable, err)
}
func (d *DataOutputStream) PutString(str string, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dataOutputStreamPutString(d, str, cancellable, err)
}
func (d *DataOutputStream) PutUint16(data uint16, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dataOutputStreamPutUint16(d, data, cancellable, err)
}
func (d *DataOutputStream) PutUint32(data T.GUint32, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dataOutputStreamPutUint32(d, data, cancellable, err)
}
func (d *DataOutputStream) PutUint64(data uint64, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dataOutputStreamPutUint64(d, data, cancellable, err)
}
func (d *DataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	dataOutputStreamSetByteOrder(d, order)
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

	dBusAnnotationInfoRef   func(d *DBusAnnotationInfo) *DBusAnnotationInfo
	dBusAnnotationInfoUnref func(d *DBusAnnotationInfo)
)

func (d *DBusAnnotationInfo) Ref() *DBusAnnotationInfo { return dBusAnnotationInfoRef(d) }
func (d *DBusAnnotationInfo) Unref()                   { dBusAnnotationInfoUnref(d) }

type DBusArgInfo struct {
	RefCount    int
	Name        *T.Gchar
	Signature   *T.Gchar
	Annotations **DBusAnnotationInfo
}

var (
	DBusArgInfoGetType func() O.Type

	dBusArgInfoRef   func(d *DBusArgInfo) *DBusArgInfo
	dBusArgInfoUnref func(d *DBusArgInfo)
)

func (d *DBusArgInfo) Ref() *DBusArgInfo { return dBusArgInfoRef(d) }
func (d *DBusArgInfo) Unref()            { dBusArgInfoUnref(d) }

type DBusAuthObserver struct{}

var (
	DBusAuthObserverGetType func() O.Type
	DBusAuthObserverNew     func() *DBusAuthObserver

	dBusAuthObserverAuthorizeAuthenticatedPeer func(d *DBusAuthObserver, stream *IOStream, credentials *Credentials) T.Gboolean
)

func (d *DBusAuthObserver) AuthorizeAuthenticatedPeer(stream *IOStream, credentials *Credentials) T.Gboolean {
	return dBusAuthObserverAuthorizeAuthenticatedPeer(d, stream, credentials)
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

	dBusConnectionAddFilter                  func(d *DBusConnection, filterFunction DBusMessageFilterFunction, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint
	dBusConnectionCall                       func(d *DBusConnection, busName, objectPath, interfaceName, methodName string, parameters *T.GVariant, replyType *T.GVariantType, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	dBusConnectionCallFinish                 func(d *DBusConnection, res *AsyncResult, err **T.GError) *T.GVariant
	dBusConnectionCallSync                   func(d *DBusConnection, busName, objectPath, interfaceName, methodName string, parameters *T.GVariant, replyType *T.GVariantType, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, err **T.GError) *T.GVariant
	dBusConnectionClose                      func(d *DBusConnection, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	dBusConnectionCloseFinish                func(d *DBusConnection, res *AsyncResult, err **T.GError) T.Gboolean
	dBusConnectionCloseSync                  func(d *DBusConnection, cancellable *Cancellable, err **T.GError) T.Gboolean
	dBusConnectionEmitSignal                 func(d *DBusConnection, destinationBusName, objectPath, interfaceName string, signalName string, parameters *T.GVariant, err **T.GError) T.Gboolean
	dBusConnectionFlush                      func(d *DBusConnection, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	dBusConnectionFlushFinish                func(d *DBusConnection, res *AsyncResult, err **T.GError) T.Gboolean
	dBusConnectionFlushSync                  func(d *DBusConnection, cancellable *Cancellable, err **T.GError) T.Gboolean
	dBusConnectionGetCapabilities            func(d *DBusConnection) DBusCapabilityFlags
	dBusConnectionGetExitOnClose             func(d *DBusConnection) T.Gboolean
	dBusConnectionGetGuid                    func(d *DBusConnection) string
	dBusConnectionGetPeerCredentials         func(d *DBusConnection) *Credentials
	dBusConnectionGetStream                  func(d *DBusConnection) *IOStream
	dBusConnectionGetUniqueName              func(d *DBusConnection) string
	dBusConnectionIsClosed                   func(d *DBusConnection) T.Gboolean
	dBusConnectionRegisterObject             func(d *DBusConnection, objectPath string, interfaceInfo *DBusInterfaceInfo, vtable *DBusInterfaceVTable, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify, err **T.GError) uint
	dBusConnectionRegisterSubtree            func(d *DBusConnection, objectPath string, vtable *DBusSubtreeVTable, flags DBusSubtreeFlags, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify, err **T.GError) uint
	dBusConnectionRemoveFilter               func(d *DBusConnection, filterId uint)
	dBusConnectionSendMessage                func(d *DBusConnection, message *DBusMessage, flags DBusSendMessageFlags, outSerial *T.GUint32, err **T.GError) T.Gboolean
	dBusConnectionSendMessageWithReply       func(d *DBusConnection, message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int, outSerial *T.GUint32, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	dBusConnectionSendMessageWithReplyFinish func(d *DBusConnection, res *AsyncResult, err **T.GError) *DBusMessage
	dBusConnectionSendMessageWithReplySync   func(d *DBusConnection, message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int, outSerial *T.GUint32, cancellable *Cancellable, err **T.GError) *DBusMessage
	dBusConnectionSetExitOnClose             func(d *DBusConnection, exitOnClose T.Gboolean)
	dBusConnectionSignalSubscribe            func(d *DBusConnection, sender, interfaceName, member, objectPath, arg0 string, flags DBusSignalFlags, callback DBusSignalCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint
	dBusConnectionSignalUnsubscribe          func(d *DBusConnection, subscriptionId uint)
	dBusConnectionStartMessageProcessing     func(d *DBusConnection)
	dBusConnectionUnregisterObject           func(d *DBusConnection, registrationId uint) T.Gboolean
	dBusConnectionUnregisterSubtree          func(d *DBusConnection, registrationId uint) T.Gboolean
)

func (d *DBusConnection) AddFilter(filterFunction DBusMessageFilterFunction, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint {
	return dBusConnectionAddFilter(d, filterFunction, userData, userDataFreeFunc)
}
func (d *DBusConnection) Call(busName, objectPath, interfaceName, methodName string, parameters *T.GVariant, replyType *T.GVariantType, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	dBusConnectionCall(d, busName, objectPath, interfaceName, methodName, parameters, replyType, flags, timeoutMsec, cancellable, callback, userData)
}
func (d *DBusConnection) CallFinish(res *AsyncResult, err **T.GError) *T.GVariant {
	return dBusConnectionCallFinish(d, res, err)
}
func (d *DBusConnection) CallSync(busName, objectPath, interfaceName, methodName string, parameters *T.GVariant, replyType *T.GVariantType, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, err **T.GError) *T.GVariant {
	return dBusConnectionCallSync(d, busName, objectPath, interfaceName, methodName, parameters, replyType, flags, timeoutMsec, cancellable, err)
}
func (d *DBusConnection) Close(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	dBusConnectionClose(d, cancellable, callback, userData)
}
func (d *DBusConnection) CloseFinish(res *AsyncResult, err **T.GError) T.Gboolean {
	return dBusConnectionCloseFinish(d, res, err)
}
func (d *DBusConnection) CloseSync(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dBusConnectionCloseSync(d, cancellable, err)
}
func (d *DBusConnection) EmitSignal(destinationBusName, objectPath, interfaceName string, signalName string, parameters *T.GVariant, err **T.GError) T.Gboolean {
	return dBusConnectionEmitSignal(d, destinationBusName, objectPath, interfaceName, signalName, parameters, err)
}
func (d *DBusConnection) Flush(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	dBusConnectionFlush(d, cancellable, callback, userData)
}
func (d *DBusConnection) FlushFinish(res *AsyncResult, err **T.GError) T.Gboolean {
	return dBusConnectionFlushFinish(d, res, err)
}
func (d *DBusConnection) FlushSync(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return dBusConnectionFlushSync(d, cancellable, err)
}
func (d *DBusConnection) GetCapabilities() DBusCapabilityFlags {
	return dBusConnectionGetCapabilities(d)
}
func (d *DBusConnection) GetExitOnClose() T.Gboolean       { return dBusConnectionGetExitOnClose(d) }
func (d *DBusConnection) GetGuid() string                  { return dBusConnectionGetGuid(d) }
func (d *DBusConnection) GetPeerCredentials() *Credentials { return dBusConnectionGetPeerCredentials(d) }
func (d *DBusConnection) GetStream() *IOStream             { return dBusConnectionGetStream(d) }
func (d *DBusConnection) GetUniqueName() string            { return dBusConnectionGetUniqueName(d) }
func (d *DBusConnection) IsClosed() T.Gboolean             { return dBusConnectionIsClosed(d) }
func (d *DBusConnection) RegisterObject(objectPath string, interfaceInfo *DBusInterfaceInfo, vtable *DBusInterfaceVTable, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify, err **T.GError) uint {
	return dBusConnectionRegisterObject(d, objectPath, interfaceInfo, vtable, userData, userDataFreeFunc, err)
}
func (d *DBusConnection) RegisterSubtree(objectPath string, vtable *DBusSubtreeVTable, flags DBusSubtreeFlags, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify, err **T.GError) uint {
	return dBusConnectionRegisterSubtree(d, objectPath, vtable, flags, userData, userDataFreeFunc, err)
}
func (d *DBusConnection) RemoveFilter(filterId uint) { dBusConnectionRemoveFilter(d, filterId) }
func (d *DBusConnection) SendMessage(message *DBusMessage, flags DBusSendMessageFlags, outSerial *T.GUint32, err **T.GError) T.Gboolean {
	return dBusConnectionSendMessage(d, message, flags, outSerial, err)
}
func (d *DBusConnection) SendMessageWithReply(message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int, outSerial *T.GUint32, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	dBusConnectionSendMessageWithReply(d, message, flags, timeoutMsec, outSerial, cancellable, callback, userData)
}
func (d *DBusConnection) SendMessageWithReplyFinish(res *AsyncResult, err **T.GError) *DBusMessage {
	return dBusConnectionSendMessageWithReplyFinish(d, res, err)
}
func (d *DBusConnection) SendMessageWithReplySync(message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int, outSerial *T.GUint32, cancellable *Cancellable, err **T.GError) *DBusMessage {
	return dBusConnectionSendMessageWithReplySync(d, message, flags, timeoutMsec, outSerial, cancellable, err)
}
func (d *DBusConnection) SetExitOnClose(exitOnClose T.Gboolean) {
	dBusConnectionSetExitOnClose(d, exitOnClose)
}
func (d *DBusConnection) SignalSubscribe(sender, interfaceName, member, objectPath, arg0 string, flags DBusSignalFlags, callback DBusSignalCallback, userData T.Gpointer, userDataFreeFunc T.GDestroyNotify) uint {
	return dBusConnectionSignalSubscribe(d, sender, interfaceName, member, objectPath, arg0, flags, callback, userData, userDataFreeFunc)
}
func (d *DBusConnection) SignalUnsubscribe(subscriptionId uint) {
	dBusConnectionSignalUnsubscribe(d, subscriptionId)
}
func (d *DBusConnection) StartMessageProcessing() { dBusConnectionStartMessageProcessing(d) }
func (d *DBusConnection) UnregisterObject(registrationId uint) T.Gboolean {
	return dBusConnectionUnregisterObject(d, registrationId)
}
func (d *DBusConnection) UnregisterSubtree(registrationId uint) T.Gboolean {
	return dBusConnectionUnregisterSubtree(d, registrationId)
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
	parameters *T.GVariant, invocation *DBusMethodInvocation,
	userData T.Gpointer)

type DBusInterfaceSetPropertyFunc func(connection *DBusConnection,
	sender, objectPath, interfaceName, propertyName string,
	value *T.GVariant, err **T.GError, userData T.Gpointer)

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
	DBusErrorIsRemoteError       func(err *T.GError) T.Gboolean
	DBusErrorNewForDBusError     func(dbusErrorName string, dbusErrorMessage string) *T.GError
	DBusErrorQuark               func() T.GQuark
	DBusErrorRegisterError       func(errorDomain T.GQuark, errorCode int, dbusErrorName string) T.Gboolean
	DBusErrorRegisterErrorDomain func(errorDomainQuarkName string, quarkVolatile *T.Gsize, entries *DBusErrorEntry, numEntries uint)
	DBusErrorSetDBusError        func(e **T.GError, dbusErrorName, dbusErrorMessage, format string, v ...VArg)
	DBusErrorSetDBusErrorValist  func(err **T.GError, dbusErrorName string, dbusErrorMessage string, format string, varArgs T.VaList)
	DBusErrorStripRemoteError    func(err *T.GError) T.Gboolean
	DBusErrorUnregisterError     func(errorDomain T.GQuark, errorCode int, dbusErrorName string) T.Gboolean
)

var DBusGenerateGuid func() string

var (
	DBusInterfaceInfoGetType func() O.Type

	dBusInterfaceInfoGenerateXml    func(d *DBusInterfaceInfo, indent uint, stringBuilder *T.GString)
	dBusInterfaceInfoLookupMethod   func(d *DBusInterfaceInfo, name string) *DBusMethodInfo
	dBusInterfaceInfoLookupProperty func(d *DBusInterfaceInfo, name string) *DBusPropertyInfo
	dBusInterfaceInfoLookupSignal   func(d *DBusInterfaceInfo, name string) *DBusSignalInfo
	dBusInterfaceInfoRef            func(d *DBusInterfaceInfo) *DBusInterfaceInfo
	dBusInterfaceInfoUnref          func(d *DBusInterfaceInfo)
)

func (d *DBusInterfaceInfo) GenerateXml(indent uint, stringBuilder *T.GString) {
	dBusInterfaceInfoGenerateXml(d, indent, stringBuilder)
}
func (d *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	return dBusInterfaceInfoLookupMethod(d, name)
}
func (d *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	return dBusInterfaceInfoLookupProperty(d, name)
}
func (d *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	return dBusInterfaceInfoLookupSignal(d, name)
}
func (d *DBusInterfaceInfo) Ref() *DBusInterfaceInfo { return dBusInterfaceInfoRef(d) }
func (d *DBusInterfaceInfo) Unref()                  { dBusInterfaceInfoUnref(d) }

var (
	DBusIsAddress          func(string string) T.Gboolean
	DBusIsGuid             func(string string) T.Gboolean
	DBusIsInterfaceName    func(string string) T.Gboolean
	DBusIsMemberName       func(string string) T.Gboolean
	DBusIsName             func(string string) T.Gboolean
	DBusIsSupportedAddress func(string string, err **T.GError) T.Gboolean
	DBusIsUniqueName       func(string string) T.Gboolean
)

type DBusMessage struct{}

var (
	DBusMessageGetType       func() O.Type
	DBusMessageNew           func() *DBusMessage
	DBusMessageNewFromBlob   func(blob *T.Guchar, blobLen T.Gsize, capabilities DBusCapabilityFlags, err **T.GError) *DBusMessage
	DBusMessageNewMethodCall func(name string, path string, interface_ string, method string) *DBusMessage
	DBusMessageNewSignal     func(path string, interface_ string, signal string) *DBusMessage

	DBusMessageBytesNeeded func(blob *T.Guchar, blobLen T.Gsize, err **T.GError) T.Gssize

	dBusMessageCopy                  func(d *DBusMessage, err **T.GError) *DBusMessage
	dBusMessageGetArg0               func(d *DBusMessage) string
	dBusMessageGetBody               func(d *DBusMessage) *T.GVariant
	dBusMessageGetByteOrder          func(d *DBusMessage) DBusMessageByteOrder
	dBusMessageGetDestination        func(d *DBusMessage) string
	dBusMessageGetErrorName          func(d *DBusMessage) string
	dBusMessageGetFlags              func(d *DBusMessage) DBusMessageFlags
	dBusMessageGetHeader             func(d *DBusMessage, headerField DBusMessageHeaderField) *T.GVariant
	dBusMessageGetHeaderFields       func(d *DBusMessage) *T.Guchar
	dBusMessageGetInterface          func(d *DBusMessage) string
	dBusMessageGetLocked             func(d *DBusMessage) T.Gboolean
	dBusMessageGetMember             func(d *DBusMessage) string
	dBusMessageGetMessageType        func(d *DBusMessage) DBusMessageType
	dBusMessageGetNumUnixFds         func(d *DBusMessage) T.GUint32
	dBusMessageGetPath               func(d *DBusMessage) string
	dBusMessageGetReplySerial        func(d *DBusMessage) T.GUint32
	dBusMessageGetSender             func(d *DBusMessage) string
	dBusMessageGetSerial             func(d *DBusMessage) T.GUint32
	dBusMessageGetSignature          func(d *DBusMessage) string
	dBusMessageGetUnixFdList         func(d *DBusMessage) *T.GUnixFDList
	dBusMessageLock                  func(d *DBusMessage)
	dBusMessageNewMethodError        func(d *DBusMessage, errorName, errorMessageFormat string, v ...VArg) *DBusMessage
	dBusMessageNewMethodErrorLiteral func(d *DBusMessage, errorName, errorMessage string) *DBusMessage
	dBusMessageNewMethodErrorValist  func(d *DBusMessage, errorName, errorMessageFormat string, varArgs T.VaList) *DBusMessage
	dBusMessageNewMethodReply        func(d *DBusMessage) *DBusMessage
	dBusMessagePrint                 func(d *DBusMessage, indent uint) string
	dBusMessageSetBody               func(d *DBusMessage, body *T.GVariant)
	dBusMessageSetByteOrder          func(d *DBusMessage, byteOrder DBusMessageByteOrder)
	dBusMessageSetDestination        func(d *DBusMessage, value string)
	dBusMessageSetErrorName          func(d *DBusMessage, value string)
	dBusMessageSetFlags              func(d *DBusMessage, flags DBusMessageFlags)
	dBusMessageSetHeader             func(d *DBusMessage, headerField DBusMessageHeaderField, value *T.GVariant)
	dBusMessageSetInterface          func(d *DBusMessage, value string)
	dBusMessageSetMember             func(d *DBusMessage, value string)
	dBusMessageSetMessageType        func(d *DBusMessage, typ DBusMessageType)
	dBusMessageSetNumUnixFds         func(d *DBusMessage, value T.GUint32)
	dBusMessageSetPath               func(d *DBusMessage, value string)
	dBusMessageSetReplySerial        func(d *DBusMessage, value T.GUint32)
	dBusMessageSetSender             func(d *DBusMessage, value string)
	dBusMessageSetSerial             func(d *DBusMessage, serial T.GUint32)
	dBusMessageSetSignature          func(d *DBusMessage, value string)
	dBusMessageSetUnixFdList         func(d *DBusMessage, fdList *T.GUnixFDList)
	dBusMessageToBlob                func(d *DBusMessage, outSize *T.Gsize, capabilities DBusCapabilityFlags, err **T.GError) *T.Guchar
	dBusMessageToGerror              func(d *DBusMessage, err **T.GError) T.Gboolean
)

func (d *DBusMessage) Copy(err **T.GError) *DBusMessage   { return dBusMessageCopy(d, err) }
func (d *DBusMessage) GetArg0() string                    { return dBusMessageGetArg0(d) }
func (d *DBusMessage) GetBody() *T.GVariant               { return dBusMessageGetBody(d) }
func (d *DBusMessage) GetByteOrder() DBusMessageByteOrder { return dBusMessageGetByteOrder(d) }
func (d *DBusMessage) GetDestination() string             { return dBusMessageGetDestination(d) }
func (d *DBusMessage) GetErrorName() string               { return dBusMessageGetErrorName(d) }
func (d *DBusMessage) GetFlags() DBusMessageFlags         { return dBusMessageGetFlags(d) }
func (d *DBusMessage) GetHeader(headerField DBusMessageHeaderField) *T.GVariant {
	return dBusMessageGetHeader(d, headerField)
}
func (d *DBusMessage) GetHeaderFields() *T.Guchar      { return dBusMessageGetHeaderFields(d) }
func (d *DBusMessage) GetInterface() string            { return dBusMessageGetInterface(d) }
func (d *DBusMessage) GetLocked() T.Gboolean           { return dBusMessageGetLocked(d) }
func (d *DBusMessage) GetMember() string               { return dBusMessageGetMember(d) }
func (d *DBusMessage) GetMessageType() DBusMessageType { return dBusMessageGetMessageType(d) }
func (d *DBusMessage) GetNumUnixFds() T.GUint32        { return dBusMessageGetNumUnixFds(d) }
func (d *DBusMessage) GetPath() string                 { return dBusMessageGetPath(d) }
func (d *DBusMessage) GetReplySerial() T.GUint32       { return dBusMessageGetReplySerial(d) }
func (d *DBusMessage) GetSender() string               { return dBusMessageGetSender(d) }
func (d *DBusMessage) GetSerial() T.GUint32            { return dBusMessageGetSerial(d) }
func (d *DBusMessage) GetSignature() string            { return dBusMessageGetSignature(d) }
func (d *DBusMessage) GetUnixFdList() *T.GUnixFDList   { return dBusMessageGetUnixFdList(d) }
func (d *DBusMessage) Lock()                           { dBusMessageLock(d) }
func (d *DBusMessage) NewMethodError(errorName, errorMessageFormat string, v ...VArg) *DBusMessage {
	return dBusMessageNewMethodError(d, errorName, errorMessageFormat, v)
}
func (d *DBusMessage) NewMethodErrorLiteral(errorName, errorMessage string) *DBusMessage {
	return dBusMessageNewMethodErrorLiteral(d, errorName, errorMessage)
}
func (d *DBusMessage) NewMethodErrorValist(errorName, errorMessageFormat string, varArgs T.VaList) *DBusMessage {
	return dBusMessageNewMethodErrorValist(d, errorName, errorMessageFormat, varArgs)
}
func (d *DBusMessage) NewMethodReply() *DBusMessage { return dBusMessageNewMethodReply(d) }
func (d *DBusMessage) Print(indent uint) string     { return dBusMessagePrint(d, indent) }
func (d *DBusMessage) SetBody(body *T.GVariant)     { dBusMessageSetBody(d, body) }
func (d *DBusMessage) SetByteOrder(byteOrder DBusMessageByteOrder) {
	dBusMessageSetByteOrder(d, byteOrder)
}
func (d *DBusMessage) SetDestination(value string)     { dBusMessageSetDestination(d, value) }
func (d *DBusMessage) SetErrorName(value string)       { dBusMessageSetErrorName(d, value) }
func (d *DBusMessage) SetFlags(flags DBusMessageFlags) { dBusMessageSetFlags(d, flags) }
func (d *DBusMessage) SetHeader(headerField DBusMessageHeaderField, value *T.GVariant) {
	dBusMessageSetHeader(d, headerField, value)
}
func (d *DBusMessage) SetInterface(value string)           { dBusMessageSetInterface(d, value) }
func (d *DBusMessage) SetMember(value string)              { dBusMessageSetMember(d, value) }
func (d *DBusMessage) SetMessageType(typ DBusMessageType)  { dBusMessageSetMessageType(d, typ) }
func (d *DBusMessage) SetNumUnixFds(value T.GUint32)       { dBusMessageSetNumUnixFds(d, value) }
func (d *DBusMessage) SetPath(value string)                { dBusMessageSetPath(d, value) }
func (d *DBusMessage) SetReplySerial(value T.GUint32)      { dBusMessageSetReplySerial(d, value) }
func (d *DBusMessage) SetSender(value string)              { dBusMessageSetSender(d, value) }
func (d *DBusMessage) SetSerial(serial T.GUint32)          { dBusMessageSetSerial(d, serial) }
func (d *DBusMessage) SetSignature(value string)           { dBusMessageSetSignature(d, value) }
func (d *DBusMessage) SetUnixFdList(fdList *T.GUnixFDList) { dBusMessageSetUnixFdList(d, fdList) }
func (d *DBusMessage) ToBlob(outSize *T.Gsize, capabilities DBusCapabilityFlags, err **T.GError) *T.Guchar {
	return dBusMessageToBlob(d, outSize, capabilities, err)
}
func (d *DBusMessage) ToGerror(err **T.GError) T.Gboolean { return dBusMessageToGerror(d, err) }

type DBusMessageByteOrder Enum

const (
	DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN    DBusMessageByteOrder = 'B'
	DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN DBusMessageByteOrder = 'l'
)

var DBusMessageByteOrderGetType func() O.Type

type DBusMessageFilterFunction func(
	connection *DBusConnection,
	message *DBusMessage,
	incoming T.Gboolean,
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

	dBusMethodInfoRef   func(d *DBusMethodInfo) *DBusMethodInfo
	dBusMethodInfoUnref func(d *DBusMethodInfo)
)

func (d *DBusMethodInfo) Ref() *DBusMethodInfo { return dBusMethodInfoRef(d) }
func (d *DBusMethodInfo) Unref()               { dBusMethodInfoUnref(d) }

type DBusMethodInvocation struct{}

var (
	DBusMethodInvocationGetType func() O.Type

	dBusMethodInvocationGetConnection      func(d *DBusMethodInvocation) *DBusConnection
	dBusMethodInvocationGetInterfaceName   func(d *DBusMethodInvocation) string
	dBusMethodInvocationGetMessage         func(d *DBusMethodInvocation) *DBusMessage
	dBusMethodInvocationGetMethodInfo      func(d *DBusMethodInvocation) *DBusMethodInfo
	dBusMethodInvocationGetMethodName      func(d *DBusMethodInvocation) string
	dBusMethodInvocationGetObjectPath      func(d *DBusMethodInvocation) string
	dBusMethodInvocationGetParameters      func(d *DBusMethodInvocation) *T.GVariant
	dBusMethodInvocationGetSender          func(d *DBusMethodInvocation) string
	dBusMethodInvocationGetUserData        func(d *DBusMethodInvocation) T.Gpointer
	dBusMethodInvocationReturnDBusError    func(d *DBusMethodInvocation, errorName, errorMessage string)
	dBusMethodInvocationReturnError        func(d *DBusMethodInvocation, domain T.GQuark, code int, format string, v ...VArg)
	dBusMethodInvocationReturnErrorLiteral func(d *DBusMethodInvocation, domain T.GQuark, code int, message string)
	dBusMethodInvocationReturnErrorValist  func(d *DBusMethodInvocation, domain T.GQuark, code int, format string, varArgs T.VaList)
	dBusMethodInvocationReturnGerror       func(d *DBusMethodInvocation, err *T.GError)
	dBusMethodInvocationReturnValue        func(d *DBusMethodInvocation, parameters *T.GVariant)
)

func (d *DBusMethodInvocation) GetConnection() *DBusConnection {
	return dBusMethodInvocationGetConnection(d)
}
func (d *DBusMethodInvocation) GetInterfaceName() string {
	return dBusMethodInvocationGetInterfaceName(d)
}
func (d *DBusMethodInvocation) GetMessage() *DBusMessage { return dBusMethodInvocationGetMessage(d) }
func (d *DBusMethodInvocation) GetMethodInfo() *DBusMethodInfo {
	return dBusMethodInvocationGetMethodInfo(d)
}
func (d *DBusMethodInvocation) GetMethodName() string { return dBusMethodInvocationGetMethodName(d) }
func (d *DBusMethodInvocation) GetObjectPath() string { return dBusMethodInvocationGetObjectPath(d) }
func (d *DBusMethodInvocation) GetParameters() *T.GVariant {
	return dBusMethodInvocationGetParameters(d)
}
func (d *DBusMethodInvocation) GetSender() string       { return dBusMethodInvocationGetSender(d) }
func (d *DBusMethodInvocation) GetUserData() T.Gpointer { return dBusMethodInvocationGetUserData(d) }
func (d *DBusMethodInvocation) ReturnDBusError(errorName, errorMessage string) {
	dBusMethodInvocationReturnDBusError(d, errorName, errorMessage)
}
func (d *DBusMethodInvocation) ReturnError(domain T.GQuark, code int, format string, v ...VArg) {
	dBusMethodInvocationReturnError(d, domain, code, format, v)
}
func (d *DBusMethodInvocation) ReturnErrorLiteral(domain T.GQuark, code int, message string) {
	dBusMethodInvocationReturnErrorLiteral(d, domain, code, message)
}
func (d *DBusMethodInvocation) ReturnErrorValist(domain T.GQuark, code int, format string, varArgs T.VaList) {
	dBusMethodInvocationReturnErrorValist(d, domain, code, format, varArgs)
}
func (d *DBusMethodInvocation) ReturnGerror(err *T.GError) { dBusMethodInvocationReturnGerror(d, err) }
func (d *DBusMethodInvocation) ReturnValue(parameters *T.GVariant) {
	dBusMethodInvocationReturnValue(d, parameters)
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

	dBusNodeInfoGenerateXml     func(d *DBusNodeInfo, indent uint, stringBuilder *T.GString)
	dBusNodeInfoLookupInterface func(d *DBusNodeInfo, name string) *DBusInterfaceInfo
	dBusNodeInfoRef             func(d *DBusNodeInfo) *DBusNodeInfo
	dBusNodeInfoUnref           func(d *DBusNodeInfo)
)

func (d *DBusNodeInfo) GenerateXml(indent uint, stringBuilder *T.GString) {
	dBusNodeInfoGenerateXml(d, indent, stringBuilder)
}
func (d *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	return dBusNodeInfoLookupInterface(d, name)
}
func (d *DBusNodeInfo) Ref() *DBusNodeInfo { return dBusNodeInfoRef(d) }
func (d *DBusNodeInfo) Unref()             { dBusNodeInfoUnref(d) }

type DBusPropertyInfo struct {
	RefCount    int
	Name        *T.Gchar
	Signature   *T.Gchar
	Flags       DBusPropertyInfoFlags
	Annotations **DBusAnnotationInfo
}

var (
	DBusPropertyInfoGetType func() O.Type

	dBusPropertyInfoRef   func(d *DBusPropertyInfo) *DBusPropertyInfo
	dBusPropertyInfoUnref func(d *DBusPropertyInfo)
)

func (d *DBusPropertyInfo) Ref() *DBusPropertyInfo { return dBusPropertyInfoRef(d) }
func (d *DBusPropertyInfo) Unref()                 { dBusPropertyInfoUnref(d) }

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

	dBusProxyCall                   func(d *DBusProxy, methodName string, parameters *T.GVariant, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	dBusProxyCallFinish             func(d *DBusProxy, res *AsyncResult, err **T.GError) *T.GVariant
	dBusProxyCallSync               func(d *DBusProxy, methodName string, parameters *T.GVariant, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, err **T.GError) *T.GVariant
	dBusProxyGetCachedProperty      func(d *DBusProxy, propertyName string) *T.GVariant
	dBusProxyGetCachedPropertyNames func(d *DBusProxy) []string
	dBusProxyGetConnection          func(d *DBusProxy) *DBusConnection
	dBusProxyGetDefaultTimeout      func(d *DBusProxy) int
	dBusProxyGetFlags               func(d *DBusProxy) DBusProxyFlags
	dBusProxyGetInterfaceInfo       func(d *DBusProxy) *DBusInterfaceInfo
	dBusProxyGetInterfaceName       func(d *DBusProxy) string
	dBusProxyGetName                func(d *DBusProxy) string
	dBusProxyGetNameOwner           func(d *DBusProxy) string
	dBusProxyGetObjectPath          func(d *DBusProxy) string
	dBusProxySetCachedProperty      func(d *DBusProxy, propertyName string, value *T.GVariant)
	dBusProxySetDefaultTimeout      func(d *DBusProxy, timeoutMsec int)
	dBusProxySetInterfaceInfo       func(d *DBusProxy, info *DBusInterfaceInfo)
)

func (d *DBusProxy) Call(methodName string, parameters *T.GVariant, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	dBusProxyCall(d, methodName, parameters, flags, timeoutMsec, cancellable, callback, userData)
}
func (d *DBusProxy) CallFinish(res *AsyncResult, err **T.GError) *T.GVariant {
	return dBusProxyCallFinish(d, res, err)
}
func (d *DBusProxy) CallSync(methodName string, parameters *T.GVariant, flags DBusCallFlags, timeoutMsec int, cancellable *Cancellable, err **T.GError) *T.GVariant {
	return dBusProxyCallSync(d, methodName, parameters, flags, timeoutMsec, cancellable, err)
}
func (d *DBusProxy) GetCachedProperty(propertyName string) *T.GVariant {
	return dBusProxyGetCachedProperty(d, propertyName)
}
func (d *DBusProxy) GetCachedPropertyNames() []string     { return dBusProxyGetCachedPropertyNames(d) }
func (d *DBusProxy) GetConnection() *DBusConnection       { return dBusProxyGetConnection(d) }
func (d *DBusProxy) GetDefaultTimeout() int               { return dBusProxyGetDefaultTimeout(d) }
func (d *DBusProxy) GetFlags() DBusProxyFlags             { return dBusProxyGetFlags(d) }
func (d *DBusProxy) GetInterfaceInfo() *DBusInterfaceInfo { return dBusProxyGetInterfaceInfo(d) }
func (d *DBusProxy) GetInterfaceName() string             { return dBusProxyGetInterfaceName(d) }
func (d *DBusProxy) GetName() string                      { return dBusProxyGetName(d) }
func (d *DBusProxy) GetNameOwner() string                 { return dBusProxyGetNameOwner(d) }
func (d *DBusProxy) GetObjectPath() string                { return dBusProxyGetObjectPath(d) }
func (d *DBusProxy) SetCachedProperty(propertyName string, value *T.GVariant) {
	dBusProxySetCachedProperty(d, propertyName, value)
}
func (d *DBusProxy) SetDefaultTimeout(timeoutMsec int)        { dBusProxySetDefaultTimeout(d, timeoutMsec) }
func (d *DBusProxy) SetInterfaceInfo(info *DBusInterfaceInfo) { dBusProxySetInterfaceInfo(d, info) }

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

	dBusServerGetClientAddress func(d *DBusServer) string
	dBusServerGetFlags         func(d *DBusServer) DBusServerFlags
	dBusServerGetGuid          func(d *DBusServer) string
	dBusServerIsActive         func(d *DBusServer) T.Gboolean
	dBusServerStart            func(d *DBusServer)
	dBusServerStop             func(d *DBusServer)
)

func (d *DBusServer) GetClientAddress() string  { return dBusServerGetClientAddress(d) }
func (d *DBusServer) GetFlags() DBusServerFlags { return dBusServerGetFlags(d) }
func (d *DBusServer) GetGuid() string           { return dBusServerGetGuid(d) }
func (d *DBusServer) IsActive() T.Gboolean      { return dBusServerIsActive(d) }
func (d *DBusServer) Start()                    { dBusServerStart(d) }
func (d *DBusServer) Stop()                     { dBusServerStop(d) }

type DBusServerFlags Enum

const (
	DBUS_SERVER_FLAGS_RUN_IN_THREAD DBusServerFlags = 1 << iota
	DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS
	DBUS_SERVER_FLAGS_NONE DBusServerFlags = 0
)

var DBusServerFlagsGetType func() O.Type

type DBusSignalCallback func(connection *DBusConnection,
	senderName, objectPath, interfaceName, signalName string,
	parameters *T.GVariant, userData T.Gpointer)

type DBusSignalInfo struct {
	RefCount    int
	Name        *T.Gchar
	Args        **DBusArgInfo
	Annotations **DBusAnnotationInfo
}

var (
	DBusSignalInfoGetType func() O.Type

	dBusSignalInfoRef   func(d *DBusSignalInfo) *DBusSignalInfo
	dBusSignalInfoUnref func(d *DBusSignalInfo)
)

func (d *DBusSignalInfo) Ref() *DBusSignalInfo { return dBusSignalInfoRef(d) }
func (d *DBusSignalInfo) Unref()               { dBusSignalInfoUnref(d) }

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

	driveCanEject                 func(d *Drive) T.Gboolean
	driveCanPollForMedia          func(d *Drive) T.Gboolean
	driveCanStart                 func(d *Drive) T.Gboolean
	driveCanStartDegraded         func(d *Drive) T.Gboolean
	driveCanStop                  func(d *Drive) T.Gboolean
	driveEject                    func(d *Drive, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	driveEjectFinish              func(d *Drive, result *AsyncResult, err **T.GError) T.Gboolean
	driveEjectWithOperation       func(d *Drive, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	driveEjectWithOperationFinish func(d *Drive, result *AsyncResult, err **T.GError) T.Gboolean
	driveEnumerateIdentifiers     func(d *Drive) []string
	driveGetIcon                  func(d *Drive) *Icon
	driveGetIdentifier            func(d *Drive, kind string) string
	driveGetName                  func(d *Drive) string
	driveGetStartStopType         func(d *Drive) DriveStartStopType
	driveGetVolumes               func(d *Drive) *T.GList
	driveHasMedia                 func(d *Drive) T.Gboolean
	driveHasVolumes               func(d *Drive) T.Gboolean
	driveIsMediaCheckAutomatic    func(d *Drive) T.Gboolean
	driveIsMediaRemovable         func(d *Drive) T.Gboolean
	drivePollForMedia             func(d *Drive, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	drivePollForMediaFinish       func(d *Drive, result *AsyncResult, err **T.GError) T.Gboolean
	driveStart                    func(d *Drive, flags DriveStartFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	driveStartFinish              func(d *Drive, result *AsyncResult, err **T.GError) T.Gboolean
	driveStop                     func(d *Drive, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	driveStopFinish               func(d *Drive, result *AsyncResult, err **T.GError) T.Gboolean
)

func (d *Drive) CanEject() T.Gboolean         { return driveCanEject(d) }
func (d *Drive) CanPollForMedia() T.Gboolean  { return driveCanPollForMedia(d) }
func (d *Drive) CanStart() T.Gboolean         { return driveCanStart(d) }
func (d *Drive) CanStartDegraded() T.Gboolean { return driveCanStartDegraded(d) }
func (d *Drive) CanStop() T.Gboolean          { return driveCanStop(d) }
func (d *Drive) Eject(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	driveEject(d, flags, cancellable, callback, userData)
}
func (d *Drive) EjectFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return driveEjectFinish(d, result, err)
}
func (d *Drive) EjectWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	driveEjectWithOperation(d, flags, mountOperation, cancellable, callback, userData)
}
func (d *Drive) EjectWithOperationFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return driveEjectWithOperationFinish(d, result, err)
}
func (d *Drive) EnumerateIdentifiers() []string       { return driveEnumerateIdentifiers(d) }
func (d *Drive) GetIcon() *Icon                       { return driveGetIcon(d) }
func (d *Drive) GetIdentifier(kind string) string     { return driveGetIdentifier(d, kind) }
func (d *Drive) GetName() string                      { return driveGetName(d) }
func (d *Drive) GetStartStopType() DriveStartStopType { return driveGetStartStopType(d) }
func (d *Drive) GetVolumes() *T.GList                 { return driveGetVolumes(d) }
func (d *Drive) HasMedia() T.Gboolean                 { return driveHasMedia(d) }
func (d *Drive) HasVolumes() T.Gboolean               { return driveHasVolumes(d) }
func (d *Drive) IsMediaCheckAutomatic() T.Gboolean    { return driveIsMediaCheckAutomatic(d) }
func (d *Drive) IsMediaRemovable() T.Gboolean         { return driveIsMediaRemovable(d) }
func (d *Drive) PollForMedia(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	drivePollForMedia(d, cancellable, callback, userData)
}
func (d *Drive) PollForMediaFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return drivePollForMediaFinish(d, result, err)
}
func (d *Drive) Start(flags DriveStartFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	driveStart(d, flags, mountOperation, cancellable, callback, userData)
}
func (d *Drive) StartFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return driveStartFinish(d, result, err)
}
func (d *Drive) Stop(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	driveStop(d, flags, mountOperation, cancellable, callback, userData)
}
func (d *Drive) StopFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return driveStopFinish(d, result, err)
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
