// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gio provides API definitions for accessing
//libgio-2.0-0.dll.
package gio

import (
	"github.com/tHinqa/outside"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type Enum int

var (
	EmblemGetType func() O.Type

	EmblemNew func(icon *Icon) *T.GEmblem

	EmblemNewWithOrigin func(
		icon *Icon, origin T.GEmblemOrigin) *T.GEmblem

	EmblemGetIcon func(emblem *T.GEmblem) *Icon

	EmblemGetOrigin func(emblem *T.GEmblem) T.GEmblemOrigin

	EmblemedIconGetType func() O.Type

	EmblemedIconNew func(icon *Icon, emblem *T.GEmblem) *Icon

	EmblemedIconGetIcon func(emblemed *T.GEmblemedIcon) *Icon

	EmblemedIconGetEmblems func(
		emblemed *T.GEmblemedIcon) *T.GList

	EmblemedIconAddEmblem func(
		emblemed *T.GEmblemedIcon, emblem *T.GEmblem)

	EmblemedIconClearEmblems func(emblemed *T.GEmblemedIcon)

	MountMountFlagsGetType func() O.Type

	MountUnmountFlagsGetType func() O.Type

	AskPasswordFlagsGetType func() O.Type

	PasswordSaveGetType func() O.Type

	MountOperationResultGetType func() O.Type

	EmblemOriginGetType func() O.Type

	ResolverErrorGetType func() O.Type

	ZlibCompressorFormatGetType func() O.Type

	UnixSocketAddressTypeGetType func() O.Type

	TlsErrorGetType func() O.Type

	TlsCertificateFlagsGetType func() O.Type

	TlsAuthenticationModeGetType func() O.Type

	TlsRehandshakeModeGetType func() O.Type

	LoadableIconGetType func() O.Type

	LoadableIconLoad func(
		icon *T.GLoadableIcon,
		size int,
		t **T.Char,
		cancellable *Cancellable,
		err **T.GError) *InputStream

	LoadableIconLoadAsync func(
		icon *T.GLoadableIcon,
		size int,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	LoadableIconLoadFinish func(
		icon *T.GLoadableIcon,
		res *AsyncResult,
		typ **T.Char,
		err **T.GError) *InputStream

	MemoryInputStreamGetType func() O.Type

	MemoryInputStreamNew func() *InputStream

	MemoryInputStreamNewFromData func(
		data *T.Void,
		len T.Gssize,
		destroy T.GDestroyNotify) *InputStream

	MemoryInputStreamAddData func(
		stream *T.GMemoryInputStream,
		data *T.Void,
		len T.Gssize,
		destroy T.GDestroyNotify)

	MemoryOutputStreamGetType func() O.Type

	MemoryOutputStreamNew func(
		data T.Gpointer,
		size T.Gsize,
		reallocFunction T.GReallocFunc,
		destroyFunction T.GDestroyNotify) *OutputStream

	MemoryOutputStreamGetData func(
		ostream *T.GMemoryOutputStream) T.Gpointer

	MemoryOutputStreamGetSize func(
		ostream *T.GMemoryOutputStream) T.Gsize

	MemoryOutputStreamGetDataSize func(
		ostream *T.GMemoryOutputStream) T.Gsize

	MemoryOutputStreamStealData func(
		ostream *T.GMemoryOutputStream) T.Gpointer

	MountGetType func() O.Type

	MountGetRoot func(
		mount *T.GMount) *File

	MountGetDefaultLocation func(
		mount *T.GMount) *File

	MountGetName func(
		mount *T.GMount) string

	MountGetIcon func(
		mount *T.GMount) *Icon

	MountGetUuid func(
		mount *T.GMount) string

	MountGetVolume func(
		mount *T.GMount) *T.GVolume

	MountGetDrive func(
		mount *T.GMount) *Drive

	MountCanUnmount func(
		mount *T.GMount) T.Gboolean

	MountCanEject func(
		mount *T.GMount) T.Gboolean

	MountUnmount func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	MountUnmountFinish func(
		mount *T.GMount,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	MountEject func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	MountEjectFinish func(
		mount *T.GMount,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	MountRemount func(
		mount *T.GMount,
		flags T.GMountMountFlags,
		mountOperation *T.GMountOperation,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	MountRemountFinish func(
		mount *T.GMount,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	MountGuessContentType func(
		mount *T.GMount,
		forceRescan T.Gboolean,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	MountGuessContentTypeFinish func(
		mount *T.GMount,
		result *AsyncResult,
		err **T.GError) **T.Gchar

	MountGuessContentTypeSync func(
		mount *T.GMount,
		forceRescan T.Gboolean,
		cancellable *Cancellable,
		err **T.GError) **T.Gchar

	MountIsShadowed func(
		mount *T.GMount) T.Gboolean

	MountShadow func(
		mount *T.GMount)

	MountUnshadow func(
		mount *T.GMount)

	MountUnmountWithOperation func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	MountUnmountWithOperationFinish func(
		mount *T.GMount,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	MountEjectWithOperation func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	MountEjectWithOperationFinish func(
		mount *T.GMount,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	MountOperationGetType func() O.Type

	MountOperationNew func() *T.GMountOperation

	MountOperationGetUsername func(
		op *T.GMountOperation) string

	MountOperationSetUsername func(
		op *T.GMountOperation,
		username string)

	MountOperationGetPassword func(
		op *T.GMountOperation) string

	MountOperationSetPassword func(
		op *T.GMountOperation,
		password string)

	MountOperationGetAnonymous func(
		op *T.GMountOperation) T.Gboolean

	MountOperationSetAnonymous func(
		op *T.GMountOperation,
		anonymous T.Gboolean)

	MountOperationGetDomain func(
		op *T.GMountOperation) string

	MountOperationSetDomain func(
		op *T.GMountOperation,
		domain string)

	MountOperationGetPasswordSave func(
		op *T.GMountOperation) T.GPasswordSave

	MountOperationSetPasswordSave func(
		op *T.GMountOperation,
		save T.GPasswordSave)

	MountOperationGetChoice func(
		op *T.GMountOperation) int

	MountOperationSetChoice func(
		op *T.GMountOperation,
		choice int)

	MountOperationReply func(
		op *T.GMountOperation,
		result T.GMountOperationResult)

	VolumeMonitorGetType func() O.Type

	VolumeMonitorGet func() *T.GVolumeMonitor

	VolumeMonitorGetConnectedDrives func(
		volumeMonitor *T.GVolumeMonitor) *T.GList

	VolumeMonitorGetVolumes func(
		volumeMonitor *T.GVolumeMonitor) *T.GList

	VolumeMonitorGetMounts func(
		volumeMonitor *T.GVolumeMonitor) *T.GList

	VolumeMonitorGetVolumeForUuid func(
		volumeMonitor *T.GVolumeMonitor,
		uuid string) *T.GVolume

	VolumeMonitorGetMountForUuid func(
		volumeMonitor *T.GVolumeMonitor,
		uuid string) *T.GMount

	VolumeMonitorAdoptOrphanMount func(
		mount *T.GMount) *T.GVolume

	NativeVolumeMonitorGetType func() O.Type

	NetworkAddressGetType func() O.Type

	NetworkAddressNew func(
		hostname string,
		port uint16) *SocketConnectable

	NetworkAddressParse func(
		hostAndPort string,
		defaultPort uint16,
		err **T.GError) *SocketConnectable

	NetworkAddressParseUri func(
		uri string,
		defaultPort uint16,
		err **T.GError) *SocketConnectable

	NetworkAddressGetHostname func(
		addr *T.GNetworkAddress) string

	NetworkAddressGetPort func(
		addr *T.GNetworkAddress) uint16

	NetworkAddressGetScheme func(
		addr *T.GNetworkAddress) string

	NetworkServiceGetType func() O.Type

	NetworkServiceNew func(
		service string,
		protocol string,
		domain string) *SocketConnectable

	NetworkServiceGetService func(
		srv *T.GNetworkService) string

	NetworkServiceGetProtocol func(
		srv *T.GNetworkService) string

	NetworkServiceGetDomain func(
		srv *T.GNetworkService) string

	NetworkServiceGetScheme func(
		srv *T.GNetworkService) string

	NetworkServiceSetScheme func(
		srv *T.GNetworkService,
		scheme string)

	PermissionGetType func() O.Type

	PermissionAcquire func(
		permission *T.GPermission,
		cancellable *Cancellable,
		err **T.GError) T.Gboolean

	PermissionAcquireAsync func(
		permission *T.GPermission,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	PermissionAcquireFinish func(
		permission *T.GPermission,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	PermissionRelease func(
		permission *T.GPermission,
		cancellable *Cancellable,
		err **T.GError) T.Gboolean

	PermissionReleaseAsync func(
		permission *T.GPermission,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	PermissionReleaseFinish func(
		permission *T.GPermission,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	PermissionGetAllowed func(
		permission *T.GPermission) T.Gboolean

	PermissionGetCanAcquire func(
		permission *T.GPermission) T.Gboolean

	PermissionGetCanRelease func(
		permission *T.GPermission) T.Gboolean

	PermissionImplUpdate func(
		permission *T.GPermission,
		allowed T.Gboolean,
		canAcquire T.Gboolean,
		canRelease T.Gboolean)

	PollableInputStreamGetType func() O.Type

	PollableInputStreamCanPoll func(
		stream *T.GPollableInputStream) T.Gboolean

	PollableInputStreamIsReadable func(
		stream *T.GPollableInputStream) T.Gboolean

	PollableInputStreamCreateSource func(
		stream *T.GPollableInputStream,
		cancellable *Cancellable) *T.GSource

	PollableInputStreamReadNonblocking func(
		stream *T.GPollableInputStream,
		buffer *T.Void,
		size T.Gsize,
		cancellable *Cancellable,
		err **T.GError) T.Gssize

	PollableSourceNew func(
		pollableStream *O.Object) *T.GSource

	PollableOutputStreamGetType func() O.Type

	PollableOutputStreamCanPoll func(
		stream *T.GPollableOutputStream) T.Gboolean

	PollableOutputStreamIsWritable func(
		stream *T.GPollableOutputStream) T.Gboolean

	PollableOutputStreamCreateSource func(
		stream *T.GPollableOutputStream,
		cancellable *Cancellable) *T.GSource

	PollableOutputStreamWriteNonblocking func(
		stream *T.GPollableOutputStream,
		buffer *T.Void,
		size T.Gsize,
		cancellable *Cancellable,
		err **T.GError) T.Gssize

	ProxyGetType func() O.Type

	ProxyGetDefaultForProtocol func(
		protocol string) *T.GProxy

	ProxyConnect func(
		proxy *T.GProxy,
		connection *T.GIOStream,
		proxyAddress *T.GProxyAddress,
		cancellable *Cancellable,
		err **T.GError) *T.GIOStream

	ProxyConnectAsync func(
		proxy *T.GProxy,
		connection *T.GIOStream,
		proxyAddress *T.GProxyAddress,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	ProxyConnectFinish func(
		proxy *T.GProxy,
		result *AsyncResult,
		err **T.GError) *T.GIOStream

	ProxySupportsHostname func(
		proxy *T.GProxy) T.Gboolean

	ProxyAddressGetType func() O.Type

	ProxyAddressNew func(
		inetaddr *InetAddress,
		port uint16,
		protocol,
		destHostname string,
		destPort uint16,
		username,
		password string) *T.GSocketAddress

	ProxyAddressGetProtocol func(
		proxy *T.GProxyAddress) string

	ProxyAddressGetDestinationHostname func(
		proxy *T.GProxyAddress) string

	ProxyAddressGetDestinationPort func(
		proxy *T.GProxyAddress) uint16

	ProxyAddressGetUsername func(
		proxy *T.GProxyAddress) string

	ProxyAddressGetPassword func(
		proxy *T.GProxyAddress) string

	ProxyAddressEnumeratorGetType func() O.Type

	ProxyResolverGetType func() O.Type

	ProxyResolverGetDefault func() *T.GProxyResolver

	ProxyResolverIsSupported func(
		resolver *T.GProxyResolver) T.Gboolean

	ProxyResolverLookup func(
		resolver *T.GProxyResolver,
		uri string,
		cancellable *Cancellable,
		err **T.GError) **T.Gchar

	ProxyResolverLookupAsync func(
		resolver *T.GProxyResolver,
		uri string,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	ProxyResolverLookupFinish func(
		resolver *T.GProxyResolver,
		result *AsyncResult,
		err **T.GError) **T.Gchar

	ResolverGetType func() O.Type

	ResolverGetDefault func() *T.GResolver

	ResolverSetDefault func(
		resolver *T.GResolver)

	ResolverLookupByName func(
		resolver *T.GResolver,
		hostname string,
		cancellable *Cancellable,
		err **T.GError) *T.GList

	ResolverLookupByNameAsync func(
		resolver *T.GResolver,
		hostname string,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	ResolverLookupByNameFinish func(
		resolver *T.GResolver,
		result *AsyncResult,
		err **T.GError) *T.GList

	ResolverFreeAddresses func(
		addresses *T.GList)

	ResolverLookupByAddress func(
		resolver *T.GResolver,
		address *InetAddress,
		cancellable *Cancellable,
		err **T.GError) string

	ResolverLookupByAddressAsync func(
		resolver *T.GResolver,
		address *InetAddress,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	ResolverLookupByAddressFinish func(
		resolver *T.GResolver,
		result *AsyncResult,
		err **T.GError) string

	ResolverLookupService func(
		resolver *T.GResolver,
		service string,
		protocol string,
		domain string,
		cancellable *Cancellable,
		err **T.GError) *T.GList

	ResolverLookupServiceAsync func(
		resolver *T.GResolver,
		service string,
		protocol string,
		domain string,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	ResolverLookupServiceFinish func(
		resolver *T.GResolver,
		result *AsyncResult,
		err **T.GError) *T.GList

	ResolverFreeTargets func(
		targets *T.GList)

	ResolverErrorQuark func() T.GQuark

	TcpConnectionGetType func() O.Type

	TcpConnectionSetGracefulDisconnect func(
		connection *T.GTcpConnection,
		gracefulDisconnect T.Gboolean)

	TcpConnectionGetGracefulDisconnect func(
		connection *T.GTcpConnection) T.Gboolean

	TcpWrapperConnectionGetType func() O.Type

	TcpWrapperConnectionNew func(
		baseIoStream *T.GIOStream,
		socket *Socket) *SocketConnection

	TcpWrapperConnectionGetBaseIoStream func(
		conn *T.GTcpWrapperConnection) *T.GIOStream

	ThemedIconGetType func() O.Type

	ThemedIconNew func(
		iconname string) *Icon

	ThemedIconNewWithDefaultFallbacks func(
		iconname string) *Icon

	ThemedIconNewFromNames func(
		iconnames **T.Char,
		len int) *Icon

	ThemedIconPrependName func(
		icon *T.GThemedIcon,
		iconname string)

	ThemedIconAppendName func(
		icon *T.GThemedIcon,
		iconname string)

	ThemedIconGetNames func(
		icon *T.GThemedIcon) **T.Gchar

	ThreadedSocketServiceGetType func() O.Type

	ThreadedSocketServiceNew func(
		maxThreads int) *SocketService

	TlsBackendGetType func() O.Type

	TlsBackendGetDefault func() *T.GTlsBackend

	TlsBackendSupportsTls func(
		backend *T.GTlsBackend) T.Gboolean

	TlsBackendGetCertificateType func(
		backend *T.GTlsBackend) O.Type

	TlsBackendGetClientConnectionType func(
		backend *T.GTlsBackend) O.Type

	TlsBackendGetServerConnectionType func(
		backend *T.GTlsBackend) O.Type

	TlsCertificateGetType func() O.Type

	TlsCertificateNewFromPem func(
		data string,
		length T.Gssize,
		err **T.GError) *T.GTlsCertificate

	TlsCertificateNewFromFile func(
		file string,
		err **T.GError) *T.GTlsCertificate

	TlsCertificateNewFromFiles func(
		certFile string,
		keyFile string,
		err **T.GError) *T.GTlsCertificate

	TlsCertificateListNewFromFile func(
		file string,
		err **T.GError) *T.GList

	TlsCertificateGetIssuer func(
		cert *T.GTlsCertificate) *T.GTlsCertificate

	TlsCertificateVerify func(
		cert *T.GTlsCertificate,
		identity *SocketConnectable,
		trustedCa *T.GTlsCertificate) T.GTlsCertificateFlags

	TlsConnectionGetType func() O.Type

	TlsConnectionSetUseSystemCertdb func(
		conn *T.GTlsConnection,
		useSystemCertdb T.Gboolean)

	TlsConnectionGetUseSystemCertdb func(
		conn *T.GTlsConnection) T.Gboolean

	TlsConnectionSetCertificate func(
		conn *T.GTlsConnection,
		certificate *T.GTlsCertificate)

	TlsConnectionGetCertificate func(
		conn *T.GTlsConnection) *T.GTlsCertificate

	TlsConnectionGetPeerCertificate func(
		conn *T.GTlsConnection) *T.GTlsCertificate

	TlsConnectionGetPeerCertificateErrors func(
		conn *T.GTlsConnection) T.GTlsCertificateFlags

	TlsConnectionSetRequireCloseNotify func(
		conn *T.GTlsConnection,
		requireCloseNotify T.Gboolean)

	TlsConnectionGetRequireCloseNotify func(
		conn *T.GTlsConnection) T.Gboolean

	TlsConnectionSetRehandshakeMode func(
		conn *T.GTlsConnection,
		mode T.GTlsRehandshakeMode)

	TlsConnectionGetRehandshakeMode func(
		conn *T.GTlsConnection) T.GTlsRehandshakeMode

	TlsConnectionHandshake func(
		conn *T.GTlsConnection,
		cancellable *Cancellable,
		err **T.GError) T.Gboolean

	TlsConnectionHandshakeAsync func(
		conn *T.GTlsConnection,
		ioPriority int,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	TlsConnectionHandshakeFinish func(
		conn *T.GTlsConnection,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	TlsErrorQuark func() T.GQuark

	TlsConnectionEmitAcceptCertificate func(
		conn *T.GTlsConnection,
		peerCert *T.GTlsCertificate,
		errors T.GTlsCertificateFlags) T.Gboolean

	TlsClientConnectionGetType func() O.Type

	TlsClientConnectionNew func(
		baseIoStream *T.GIOStream,
		serverIdentity *SocketConnectable,
		err **T.GError) *T.GIOStream

	TlsClientConnectionGetValidationFlags func(
		conn *T.GTlsClientConnection) T.GTlsCertificateFlags

	TlsClientConnectionSetValidationFlags func(
		conn *T.GTlsClientConnection,
		flags T.GTlsCertificateFlags)

	TlsClientConnectionGetServerIdentity func(
		conn *T.GTlsClientConnection) *SocketConnectable

	TlsClientConnectionSetServerIdentity func(
		conn *T.GTlsClientConnection,
		identity *SocketConnectable)

	TlsClientConnectionGetUseSsl3 func(
		conn *T.GTlsClientConnection) T.Gboolean

	TlsClientConnectionSetUseSsl3 func(
		conn *T.GTlsClientConnection,
		useSsl3 T.Gboolean)

	TlsClientConnectionGetAcceptedCas func(
		conn *T.GTlsClientConnection) *T.GList

	TlsServerConnectionGetType func() O.Type

	TlsServerConnectionNew func(
		baseIoStream *T.GIOStream,
		certificate *T.GTlsCertificate,
		err **T.GError) *T.GIOStream

	VfsGetType func() O.Type

	VfsIsActive func(
		vfs *T.GVfs) T.Gboolean

	VfsGetFileForPath func(
		vfs *T.GVfs,
		path string) *File

	VfsGetFileForUri func(
		vfs *T.GVfs,
		uri string) *File

	VfsGetSupportedUriSchemes func(
		vfs *T.GVfs) **T.Gchar

	VfsParseName func(
		vfs *T.GVfs,
		parseName string) *File

	VfsGetDefault func() *T.GVfs

	VfsGetLocal func() *T.GVfs

	VolumeGetType func() O.Type

	VolumeGetName func(
		volume *T.GVolume) string

	VolumeGetIcon func(
		volume *T.GVolume) *Icon

	VolumeGetUuid func(
		volume *T.GVolume) string

	VolumeGetDrive func(
		volume *T.GVolume) *Drive

	VolumeGetMount func(
		volume *T.GVolume) *T.GMount

	VolumeCanMount func(
		volume *T.GVolume) T.Gboolean

	VolumeCanEject func(
		volume *T.GVolume) T.Gboolean

	VolumeShouldAutomount func(
		volume *T.GVolume) T.Gboolean

	VolumeMount func(
		volume *T.GVolume,
		flags T.GMountMountFlags,
		mountOperation *T.GMountOperation,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	VolumeMountFinish func(
		volume *T.GVolume,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	VolumeEject func(
		volume *T.GVolume,
		flags T.GMountUnmountFlags,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	VolumeEjectFinish func(
		volume *T.GVolume,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	VolumeGetIdentifier func(
		volume *T.GVolume,
		kind string) string

	VolumeEnumerateIdentifiers func(
		volume *T.GVolume) **T.Char

	VolumeGetActivationRoot func(
		volume *T.GVolume) *File

	VolumeEjectWithOperation func(
		volume *T.GVolume,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *Cancellable,
		callback AsyncReadyCallback,
		userData T.Gpointer)

	VolumeEjectWithOperationFinish func(
		volume *T.GVolume,
		result *AsyncResult,
		err **T.GError) T.Gboolean

	ZlibCompressorGetType func() O.Type

	ZlibCompressorNew func(
		format T.GZlibCompressorFormat,
		level int) *T.GZlibCompressor

	ZlibCompressorGetFileInfo func(
		compressor *T.GZlibCompressor) *FileInfo

	ZlibCompressorSetFileInfo func(
		compressor *T.GZlibCompressor,
		fileInfo *FileInfo)

	ZlibDecompressorGetType func() O.Type

	ZlibDecompressorNew func(
		format T.GZlibCompressorFormat) *T.GZlibDecompressor

	ZlibDecompressorGetFileInfo func(
		decompressor *T.GZlibDecompressor) *FileInfo

	Win32InputStreamGetType func() O.Type

	Win32InputStreamNew func(
		handle *T.Void,
		closeHandle T.Gboolean) *InputStream

	Win32InputStreamSetCloseHandle func(
		stream *T.GWin32InputStream,
		closeHandle T.Gboolean)

	Win32InputStreamGetCloseHandle func(
		stream *T.GWin32InputStream) T.Gboolean

	Win32InputStreamGetHandle func(
		stream *T.GWin32InputStream) *T.Void

	Win32OutputStreamGetType func() O.Type

	Win32OutputStreamNew func(
		handle *T.Void,
		closeHandle T.Gboolean) *OutputStream

	Win32OutputStreamSetCloseHandle func(
		stream *T.GWin32OutputStream,
		closeHandle T.Gboolean)

	Win32OutputStreamGetCloseHandle func(
		stream *T.GWin32OutputStream) T.Gboolean

	Win32OutputStreamGetHandle func(
		stream *T.GWin32OutputStream) *T.Void

	KeyfileSettingsBackendNew func(
		filename string,
		rootPath string,
		rootGroup string) *SettingsBackend

	NullSettingsBackendNew func() *SettingsBackend

	MemorySettingsBackendNew func() *SettingsBackend
)
var dll = "libgio-2.0-0.dll"

var apiList = outside.Apis{
	{"g_action_activate", &actionActivate},
	{"g_action_get_enabled", &actionGetEnabled},
	{"g_action_get_name", &ActionGetName},
	{"g_action_get_parameter_type", &actionGetParameterType},
	{"g_action_get_state", &actionGetState},
	{"g_action_get_state_hint", &actionGetStateHint},
	{"g_action_get_state_type", &actionGetStateType},
	{"g_action_get_type", &ActionGetType},
	{"g_action_group_action_added", &actionGroupActionAdded},
	{"g_action_group_action_enabled_changed", &actionGroupActionEnabledChanged},
	{"g_action_group_action_removed", &actionGroupActionRemoved},
	{"g_action_group_action_state_changed", &actionGroupActionStateChanged},
	{"g_action_group_activate_action", &actionGroupActivateAction},
	{"g_action_group_change_action_state", &actionGroupChangeActionState},
	{"g_action_group_get_action_enabled", &actionGroupGetActionEnabled},
	{"g_action_group_get_action_parameter_type", &actionGroupGetActionParameterType},
	{"g_action_group_get_action_state", &actionGroupGetActionState},
	{"g_action_group_get_action_state_hint", &actionGroupGetActionStateHint},
	{"g_action_group_get_action_state_type", &actionGroupGetActionStateType},
	{"g_action_group_get_type", &ActionGroupGetType},
	{"g_action_group_has_action", &actionGroupHasAction},
	{"g_action_group_list_actions", &actionGroupListActions},
	{"g_action_set_state", &actionSetState},
	{"g_app_info_add_supports_type", &appInfoAddSupportsType},
	{"g_app_info_can_delete", &appInfoCanDelete},
	{"g_app_info_can_remove_supports_type", &appInfoCanRemoveSupportsType},
	{"g_app_info_create_flags_get_type", &AppInfoCreateFlagsGetType},
	{"g_app_info_create_from_commandline", &AppInfoCreateFromCommandline},
	{"g_app_info_delete", &appInfoDelete},
	{"g_app_info_dup", &appInfoDup},
	{"g_app_info_equal", &appInfoEqual},
	{"g_app_info_get_all", &AppInfoGetAll},
	{"g_app_info_get_all_for_type", &AppInfoGetAllForType},
	{"g_app_info_get_commandline", &appInfoGetCommandline},
	{"g_app_info_get_default_for_type", &AppInfoGetDefaultForType},
	{"g_app_info_get_default_for_uri_scheme", &AppInfoGetDefaultForUriScheme},
	{"g_app_info_get_description", &appInfoGetDescription},
	{"g_app_info_get_display_name", &appInfoGetDisplayName},
	{"g_app_info_get_executable", &appInfoGetExecutable},
	{"g_app_info_get_fallback_for_type", &AppInfoGetFallbackForType},
	{"g_app_info_get_icon", &appInfoGetIcon},
	{"g_app_info_get_id", &appInfoGetId},
	{"g_app_info_get_name", &appInfoGetName},
	{"g_app_info_get_recommended_for_type", &AppInfoGetRecommendedForType},
	{"g_app_info_get_type", &AppInfoGetType},
	{"g_app_info_launch", &appInfoLaunch},
	{"g_app_info_launch_default_for_uri", &AppInfoLaunchDefaultForUri},
	{"g_app_info_launch_uris", &appInfoLaunchUris},
	{"g_app_info_remove_supports_type", &appInfoRemoveSupportsType},
	{"g_app_info_reset_type_associations", &AppInfoResetTypeAssociations},
	{"g_app_info_set_as_default_for_extension", &appInfoSetAsDefaultForExtension},
	{"g_app_info_set_as_default_for_type", &appInfoSetAsDefaultForType},
	{"g_app_info_set_as_last_used_for_type", &appInfoSetAsLastUsedForType},
	{"g_app_info_should_show", &appInfoShouldShow},
	{"g_app_info_supports_files", &appInfoSupportsFiles},
	{"g_app_info_supports_uris", &appInfoSupportsUris},
	{"g_app_launch_context_get_display", &appLaunchContextGetDisplay},
	{"g_app_launch_context_get_startup_notify_id", &appLaunchContextGetStartupNotifyId},
	{"g_app_launch_context_get_type", &AppLaunchContextGetType},
	{"g_app_launch_context_launch_failed", &appLaunchContextLaunchFailed},
	{"g_app_launch_context_new", &AppLaunchContextNew},
	{"g_application_activate", &applicationActivate},
	{"g_application_command_line_get_arguments", &applicationCommandLineGetArguments},
	{"g_application_command_line_get_cwd", &applicationCommandLineGetCwd},
	{"g_application_command_line_get_environ", &applicationCommandLineGetEnviron},
	{"g_application_command_line_get_exit_status", &applicationCommandLineGetExitStatus},
	{"g_application_command_line_get_is_remote", &applicationCommandLineGetIsRemote},
	{"g_application_command_line_get_platform_data", &applicationCommandLineGetPlatformData},
	{"g_application_command_line_get_type", &ApplicationCommandLineGetType},
	{"g_application_command_line_getenv", &applicationCommandLineGetenv},
	{"g_application_command_line_print", &applicationCommandLinePrint},
	{"g_application_command_line_printerr", &applicationCommandLinePrinterr},
	{"g_application_command_line_set_exit_status", &applicationCommandLineSetExitStatus},
	{"g_application_flags_get_type", &ApplicationFlagsGetType},
	{"g_application_get_application_id", &applicationGetApplicationId},
	{"g_application_get_flags", &applicationGetFlags},
	{"g_application_get_inactivity_timeout", &applicationGetInactivityTimeout},
	{"g_application_get_is_registered", &applicationGetIsRegistered},
	{"g_application_get_is_remote", &applicationGetIsRemote},
	{"g_application_get_type", &ApplicationGetType},
	{"g_application_hold", &applicationHold},
	{"g_application_id_is_valid", &ApplicationIdIsValid},
	{"g_application_new", &ApplicationNew},
	{"g_application_open", &applicationOpen},
	{"g_application_register", &applicationRegister},
	{"g_application_release", &applicationRelease},
	{"g_application_run", &applicationRun},
	{"g_application_set_action_group", &applicationSetActionGroup},
	{"g_application_set_application_id", &applicationSetApplicationId},
	{"g_application_set_flags", &applicationSetFlags},
	{"g_application_set_inactivity_timeout", &applicationSetInactivityTimeout},
	{"g_ask_password_flags_get_type", &AskPasswordFlagsGetType},
	{"g_async_initable_get_type", &AsyncInitableGetType},
	{"g_async_initable_init_async", &asyncInitableInitAsync},
	{"g_async_initable_init_finish", &asyncInitableInitFinish},
	{"g_async_initable_new_async", &AsyncInitableNewAsync},
	{"g_async_initable_new_finish", &asyncInitableNewFinish},
	{"g_async_initable_new_valist_async", &AsyncInitableNewValistAsync},
	{"g_async_initable_newv_async", &AsyncInitableNewvAsync},
	{"g_async_result_get_source_object", &asyncResultGetSourceObject},
	{"g_async_result_get_type", &AsyncResultGetType},
	{"g_async_result_get_user_data", &asyncResultGetUserData},
	{"g_buffered_input_stream_fill", &bufferedInputStreamFill},
	{"g_buffered_input_stream_fill_async", &bufferedInputStreamFillAsync},
	{"g_buffered_input_stream_fill_finish", &bufferedInputStreamFillFinish},
	{"g_buffered_input_stream_get_available", &bufferedInputStreamGetAvailable},
	{"g_buffered_input_stream_get_buffer_size", &bufferedInputStreamGetBufferSize},
	{"g_buffered_input_stream_get_type", &BufferedInputStreamGetType},
	{"g_buffered_input_stream_new", &BufferedInputStreamNew},
	{"g_buffered_input_stream_new_sized", &BufferedInputStreamNewSized},
	{"g_buffered_input_stream_peek", &bufferedInputStreamPeek},
	{"g_buffered_input_stream_peek_buffer", &bufferedInputStreamPeekBuffer},
	{"g_buffered_input_stream_read_byte", &bufferedInputStreamReadByte},
	{"g_buffered_input_stream_set_buffer_size", &bufferedInputStreamSetBufferSize},
	{"g_buffered_output_stream_get_auto_grow", &bufferedOutputStreamGetAutoGrow},
	{"g_buffered_output_stream_get_buffer_size", &bufferedOutputStreamGetBufferSize},
	{"g_buffered_output_stream_get_type", &BufferedOutputStreamGetType},
	{"g_buffered_output_stream_new", &BufferedOutputStreamNew},
	{"g_buffered_output_stream_new_sized", &BufferedOutputStreamNewSized},
	{"g_buffered_output_stream_set_auto_grow", &bufferedOutputStreamSetAutoGrow},
	{"g_buffered_output_stream_set_buffer_size", &bufferedOutputStreamSetBufferSize},
	{"g_bus_get", &busGet},
	{"g_bus_get_finish", &BusGetFinish},
	{"g_bus_get_sync", &busGetSync},
	{"g_bus_name_owner_flags_get_type", &BusNameOwnerFlagsGetType},
	{"g_bus_name_watcher_flags_get_type", &BusNameWatcherFlagsGetType},
	{"g_bus_own_name", &busOwnName},
	{"g_bus_own_name_on_connection", &BusOwnNameOnConnection},
	{"g_bus_own_name_on_connection_with_closures", &BusOwnNameOnConnectionWithClosures},
	{"g_bus_own_name_with_closures", &busOwnNameWithClosures},
	{"g_bus_type_get_type", &BusTypeGetType},
	{"g_bus_unown_name", &BusUnownName},
	{"g_bus_unwatch_name", &BusUnwatchName},
	{"g_bus_watch_name", &busWatchName},
	{"g_bus_watch_name_on_connection", &BusWatchNameOnConnection},
	{"g_bus_watch_name_on_connection_with_closures", &BusWatchNameOnConnectionWithClosures},
	{"g_bus_watch_name_with_closures", &busWatchNameWithClosures},
	{"g_cancellable_cancel", &cancellableCancel},
	{"g_cancellable_connect", &cancellableConnect},
	{"g_cancellable_disconnect", &cancellableDisconnect},
	{"g_cancellable_get_current", &CancellableGetCurrent},
	{"g_cancellable_get_fd", &cancellableGetFd},
	{"g_cancellable_get_type", &CancellableGetType},
	{"g_cancellable_is_cancelled", &cancellableIsCancelled},
	{"g_cancellable_make_pollfd", &cancellableMakePollfd},
	{"g_cancellable_new", &CancellableNew},
	{"g_cancellable_pop_current", &cancellablePopCurrent},
	{"g_cancellable_push_current", &cancellablePushCurrent},
	{"g_cancellable_release_fd", &cancellableReleaseFd},
	{"g_cancellable_reset", &cancellableReset},
	{"g_cancellable_set_error_if_cancelled", &cancellableSetErrorIfCancelled},
	{"g_cancellable_source_new", &cancellableSourceNew},
	{"g_charset_converter_get_num_fallbacks", &charsetConverterGetNumFallbacks},
	{"g_charset_converter_get_type", &charsetConverterGetType},
	{"g_charset_converter_get_use_fallback", &charsetConverterGetUseFallback},
	{"g_charset_converter_new", &charsetConverterNew},
	{"g_charset_converter_set_use_fallback", &charsetConverterSetUseFallback},
	{"g_content_type_can_be_executable", &ContentTypeCanBeExecutable},
	{"g_content_type_equals", &ContentTypeEquals},
	{"g_content_type_from_mime_type", &ContentTypeFromMimeType},
	{"g_content_type_get_description", &ContentTypeGetDescription},
	{"g_content_type_get_icon", &ContentTypeGetIcon},
	{"g_content_type_get_mime_type", &ContentTypeGetMimeType},
	{"g_content_type_guess", &ContentTypeGuess},
	{"g_content_type_guess_for_tree", &ContentTypeGuessForTree},
	{"g_content_type_is_a", &ContentTypeIsA},
	{"g_content_type_is_unknown", &ContentTypeIsUnknown},
	{"g_content_types_get_registered", &ContentTypesGetRegistered},
	{"g_converter_convert", &converterConvert},
	{"g_converter_flags_get_type", &ConverterFlagsGetType},
	{"g_converter_get_type", &ConverterGetType},
	{"g_converter_input_stream_get_converter", &converterInputStreamGetConverter},
	{"g_converter_input_stream_get_type", &ConverterInputStreamGetType},
	{"g_converter_input_stream_new", &ConverterInputStreamNew},
	{"g_converter_output_stream_get_converter", &converterOutputStreamGetConverter},
	{"g_converter_output_stream_get_type", &ConverterOutputStreamGetType},
	{"g_converter_output_stream_new", &ConverterOutputStreamNew},
	{"g_converter_reset", &converterReset},
	{"g_converter_result_get_type", &ConverterResultGetType},
	{"g_credentials_get_native", &credentialsGetNative},
	{"g_credentials_get_type", &CredentialsGetType},
	{"g_credentials_is_same_user", &credentialsIsSameUser},
	{"g_credentials_new", &CredentialsNew},
	{"g_credentials_set_native", &credentialsSetNative},
	{"g_credentials_to_string", &credentialsToString},
	{"g_credentials_type_get_type", &CredentialsTypeGetType},
	{"g_data_input_stream_get_byte_order", &dataInputStreamGetByteOrder},
	{"g_data_input_stream_get_newline_type", &dataInputStreamGetNewlineType},
	{"g_data_input_stream_get_type", &DataInputStreamGetType},
	{"g_data_input_stream_new", &DataInputStreamNew},
	{"g_data_input_stream_read_byte", &dataInputStreamReadByte},
	{"g_data_input_stream_read_int16", &dataInputStreamReadInt16},
	{"g_data_input_stream_read_int32", &dataInputStreamReadInt32},
	{"g_data_input_stream_read_int64", &dataInputStreamReadInt64},
	{"g_data_input_stream_read_line", &dataInputStreamReadLine},
	{"g_data_input_stream_read_line_async", &dataInputStreamReadLineAsync},
	{"g_data_input_stream_read_line_finish", &dataInputStreamReadLineFinish},
	{"g_data_input_stream_read_uint16", &dataInputStreamReadUint16},
	{"g_data_input_stream_read_uint32", &dataInputStreamReadUint32},
	{"g_data_input_stream_read_uint64", &dataInputStreamReadUint64},
	{"g_data_input_stream_read_until", &dataInputStreamReadUntil},
	{"g_data_input_stream_read_until_async", &dataInputStreamReadUntilAsync},
	{"g_data_input_stream_read_until_finish", &dataInputStreamReadUntilFinish},
	{"g_data_input_stream_read_upto", &dataInputStreamReadUpto},
	{"g_data_input_stream_read_upto_async", &dataInputStreamReadUptoAsync},
	{"g_data_input_stream_read_upto_finish", &dataInputStreamReadUptoFinish},
	{"g_data_input_stream_set_byte_order", &dataInputStreamSetByteOrder},
	{"g_data_input_stream_set_newline_type", &dataInputStreamSetNewlineType},
	{"g_data_output_stream_get_byte_order", &dataOutputStreamGetByteOrder},
	{"g_data_output_stream_get_type", &DataOutputStreamGetType},
	{"g_data_output_stream_new", &DataOutputStreamNew},
	{"g_data_output_stream_put_byte", &dataOutputStreamPutByte},
	{"g_data_output_stream_put_int16", &dataOutputStreamPutInt16},
	{"g_data_output_stream_put_int32", &dataOutputStreamPutInt32},
	{"g_data_output_stream_put_int64", &dataOutputStreamPutInt64},
	{"g_data_output_stream_put_string", &dataOutputStreamPutString},
	{"g_data_output_stream_put_uint16", &dataOutputStreamPutUint16},
	{"g_data_output_stream_put_uint32", &dataOutputStreamPutUint32},
	{"g_data_output_stream_put_uint64", &dataOutputStreamPutUint64},
	{"g_data_output_stream_set_byte_order", &dataOutputStreamSetByteOrder},
	{"g_data_stream_byte_order_get_type", &DataStreamByteOrderGetType},
	{"g_data_stream_newline_type_get_type", &DataStreamNewlineTypeGetType},
	{"g_dbus_address_get_for_bus_sync", &DBusAddressGetForBusSync},
	{"g_dbus_address_get_stream", &DBusAddressGetStream},
	{"g_dbus_address_get_stream_finish", &DBusAddressGetStreamFinish},
	{"g_dbus_address_get_stream_sync", &DBusAddressGetStreamSync},
	{"g_dbus_annotation_info_get_type", &DBusAnnotationInfoGetType},
	{"g_dbus_annotation_info_lookup", &DBusAnnotationInfoLookup},
	{"g_dbus_annotation_info_ref", &dBusAnnotationInfoRef},
	{"g_dbus_annotation_info_unref", &dBusAnnotationInfoUnref},
	{"g_dbus_arg_info_get_type", &DBusArgInfoGetType},
	{"g_dbus_arg_info_ref", &dBusArgInfoRef},
	{"g_dbus_arg_info_unref", &dBusArgInfoUnref},
	{"g_dbus_auth_observer_authorize_authenticated_peer", &dBusAuthObserverAuthorizeAuthenticatedPeer},
	{"g_dbus_auth_observer_get_type", &DBusAuthObserverGetType},
	{"g_dbus_auth_observer_new", &DBusAuthObserverNew},
	{"g_dbus_call_flags_get_type", &DBusCallFlagsGetType},
	{"g_dbus_capability_flags_get_type", &DBusCapabilityFlagsGetType},
	{"g_dbus_connection_add_filter", &dBusConnectionAddFilter},
	{"g_dbus_connection_call", &dBusConnectionCall},
	{"g_dbus_connection_call_finish", &dBusConnectionCallFinish},
	{"g_dbus_connection_call_sync", &dBusConnectionCallSync},
	{"g_dbus_connection_close", &dBusConnectionClose},
	{"g_dbus_connection_close_finish", &dBusConnectionCloseFinish},
	{"g_dbus_connection_close_sync", &dBusConnectionCloseSync},
	{"g_dbus_connection_emit_signal", &dBusConnectionEmitSignal},
	{"g_dbus_connection_flags_get_type", &DBusConnectionFlagsGetType},
	{"g_dbus_connection_flush", &dBusConnectionFlush},
	{"g_dbus_connection_flush_finish", &dBusConnectionFlushFinish},
	{"g_dbus_connection_flush_sync", &dBusConnectionFlushSync},
	{"g_dbus_connection_get_capabilities", &dBusConnectionGetCapabilities},
	{"g_dbus_connection_get_exit_on_close", &dBusConnectionGetExitOnClose},
	{"g_dbus_connection_get_guid", &dBusConnectionGetGuid},
	{"g_dbus_connection_get_peer_credentials", &dBusConnectionGetPeerCredentials},
	{"g_dbus_connection_get_stream", &dBusConnectionGetStream},
	{"g_dbus_connection_get_type", &DBusConnectionGetType},
	{"g_dbus_connection_get_unique_name", &dBusConnectionGetUniqueName},
	{"g_dbus_connection_is_closed", &dBusConnectionIsClosed},
	{"g_dbus_connection_new", &DBusConnectionNew},
	{"g_dbus_connection_new_finish", &DBusConnectionNewFinish},
	{"g_dbus_connection_new_for_address", &DBusConnectionNewForAddress},
	{"g_dbus_connection_new_for_address_finish", &DBusConnectionNewForAddressFinish},
	{"g_dbus_connection_new_for_address_sync", &DBusConnectionNewForAddressSync},
	{"g_dbus_connection_new_sync", &DBusConnectionNewSync},
	{"g_dbus_connection_register_object", &dBusConnectionRegisterObject},
	{"g_dbus_connection_register_subtree", &dBusConnectionRegisterSubtree},
	{"g_dbus_connection_remove_filter", &dBusConnectionRemoveFilter},
	{"g_dbus_connection_send_message", &dBusConnectionSendMessage},
	{"g_dbus_connection_send_message_with_reply", &dBusConnectionSendMessageWithReply},
	{"g_dbus_connection_send_message_with_reply_finish", &dBusConnectionSendMessageWithReplyFinish},
	{"g_dbus_connection_send_message_with_reply_sync", &dBusConnectionSendMessageWithReplySync},
	{"g_dbus_connection_set_exit_on_close", &dBusConnectionSetExitOnClose},
	{"g_dbus_connection_signal_subscribe", &dBusConnectionSignalSubscribe},
	{"g_dbus_connection_signal_unsubscribe", &dBusConnectionSignalUnsubscribe},
	{"g_dbus_connection_start_message_processing", &dBusConnectionStartMessageProcessing},
	{"g_dbus_connection_unregister_object", &dBusConnectionUnregisterObject},
	{"g_dbus_connection_unregister_subtree", &dBusConnectionUnregisterSubtree},
	{"g_dbus_error_encode_gerror", &DBusErrorEncodeGerror},
	{"g_dbus_error_get_remote_error", &DBusErrorGetRemoteError},
	{"g_dbus_error_get_type", &DBusErrorGetType},
	{"g_dbus_error_is_remote_error", &DBusErrorIsRemoteError},
	{"g_dbus_error_new_for_dbus_error", &DBusErrorNewForDBusError},
	{"g_dbus_error_quark", &DBusErrorQuark},
	{"g_dbus_error_register_error", &DBusErrorRegisterError},
	{"g_dbus_error_register_error_domain", &DBusErrorRegisterErrorDomain},
	{"g_dbus_error_set_dbus_error", &DBusErrorSetDBusError},
	{"g_dbus_error_set_dbus_error_valist", &DBusErrorSetDBusErrorValist},
	{"g_dbus_error_strip_remote_error", &DBusErrorStripRemoteError},
	{"g_dbus_error_unregister_error", &DBusErrorUnregisterError},
	{"g_dbus_generate_guid", &DBusGenerateGuid},
	{"g_dbus_interface_info_generate_xml", &dBusInterfaceInfoGenerateXml},
	{"g_dbus_interface_info_get_type", &DBusInterfaceInfoGetType},
	{"g_dbus_interface_info_lookup_method", &dBusInterfaceInfoLookupMethod},
	{"g_dbus_interface_info_lookup_property", &dBusInterfaceInfoLookupProperty},
	{"g_dbus_interface_info_lookup_signal", &dBusInterfaceInfoLookupSignal},
	{"g_dbus_interface_info_ref", &dBusInterfaceInfoRef},
	{"g_dbus_interface_info_unref", &dBusInterfaceInfoUnref},
	{"g_dbus_is_address", &DBusIsAddress},
	{"g_dbus_is_guid", &DBusIsGuid},
	{"g_dbus_is_interface_name", &DBusIsInterfaceName},
	{"g_dbus_is_member_name", &DBusIsMemberName},
	{"g_dbus_is_name", &DBusIsName},
	{"g_dbus_is_supported_address", &DBusIsSupportedAddress},
	{"g_dbus_is_unique_name", &DBusIsUniqueName},
	{"g_dbus_message_byte_order_get_type", &DBusMessageByteOrderGetType},
	{"g_dbus_message_bytes_needed", &DBusMessageBytesNeeded},
	{"g_dbus_message_copy", &dBusMessageCopy},
	{"g_dbus_message_flags_get_type", &DBusMessageFlagsGetType},
	{"g_dbus_message_get_arg0", &dBusMessageGetArg0},
	{"g_dbus_message_get_body", &dBusMessageGetBody},
	{"g_dbus_message_get_byte_order", &dBusMessageGetByteOrder},
	{"g_dbus_message_get_destination", &dBusMessageGetDestination},
	{"g_dbus_message_get_error_name", &dBusMessageGetErrorName},
	{"g_dbus_message_get_flags", &dBusMessageGetFlags},
	{"g_dbus_message_get_header", &dBusMessageGetHeader},
	{"g_dbus_message_get_header_fields", &dBusMessageGetHeaderFields},
	{"g_dbus_message_get_interface", &dBusMessageGetInterface},
	{"g_dbus_message_get_locked", &dBusMessageGetLocked},
	{"g_dbus_message_get_member", &dBusMessageGetMember},
	{"g_dbus_message_get_message_type", &dBusMessageGetMessageType},
	{"g_dbus_message_get_num_unix_fds", &dBusMessageGetNumUnixFds},
	{"g_dbus_message_get_path", &dBusMessageGetPath},
	{"g_dbus_message_get_reply_serial", &dBusMessageGetReplySerial},
	{"g_dbus_message_get_sender", &dBusMessageGetSender},
	{"g_dbus_message_get_serial", &dBusMessageGetSerial},
	{"g_dbus_message_get_signature", &dBusMessageGetSignature},
	{"g_dbus_message_get_type", &DBusMessageGetType},
	{"g_dbus_message_header_field_get_type", &DBusMessageHeaderFieldGetType},
	{"g_dbus_message_lock", &dBusMessageLock},
	{"g_dbus_message_new", &DBusMessageNew},
	{"g_dbus_message_new_from_blob", &DBusMessageNewFromBlob},
	{"g_dbus_message_new_method_call", &DBusMessageNewMethodCall},
	{"g_dbus_message_new_method_error", &dBusMessageNewMethodError},
	{"g_dbus_message_new_method_error_literal", &dBusMessageNewMethodErrorLiteral},
	{"g_dbus_message_new_method_error_valist", &dBusMessageNewMethodErrorValist},
	{"g_dbus_message_new_method_reply", &dBusMessageNewMethodReply},
	{"g_dbus_message_new_signal", &DBusMessageNewSignal},
	{"g_dbus_message_print", &dBusMessagePrint},
	{"g_dbus_message_set_body", &dBusMessageSetBody},
	{"g_dbus_message_set_byte_order", &dBusMessageSetByteOrder},
	{"g_dbus_message_set_destination", &dBusMessageSetDestination},
	{"g_dbus_message_set_error_name", &dBusMessageSetErrorName},
	{"g_dbus_message_set_flags", &dBusMessageSetFlags},
	{"g_dbus_message_set_header", &dBusMessageSetHeader},
	{"g_dbus_message_set_interface", &dBusMessageSetInterface},
	{"g_dbus_message_set_member", &dBusMessageSetMember},
	{"g_dbus_message_set_message_type", &dBusMessageSetMessageType},
	{"g_dbus_message_set_num_unix_fds", &dBusMessageSetNumUnixFds},
	{"g_dbus_message_set_path", &dBusMessageSetPath},
	{"g_dbus_message_set_reply_serial", &dBusMessageSetReplySerial},
	{"g_dbus_message_set_sender", &dBusMessageSetSender},
	{"g_dbus_message_set_serial", &dBusMessageSetSerial},
	{"g_dbus_message_set_signature", &dBusMessageSetSignature},
	{"g_dbus_message_to_blob", &dBusMessageToBlob},
	{"g_dbus_message_to_gerror", &dBusMessageToGerror},
	{"g_dbus_message_type_get_type", &DBusMessageTypeGetType},
	{"g_dbus_method_info_get_type", &DBusMethodInfoGetType},
	{"g_dbus_method_info_ref", &dBusMethodInfoRef},
	{"g_dbus_method_info_unref", &dBusMethodInfoUnref},
	{"g_dbus_method_invocation_get_connection", &dBusMethodInvocationGetConnection},
	{"g_dbus_method_invocation_get_interface_name", &dBusMethodInvocationGetInterfaceName},
	{"g_dbus_method_invocation_get_message", &dBusMethodInvocationGetMessage},
	{"g_dbus_method_invocation_get_method_info", &dBusMethodInvocationGetMethodInfo},
	{"g_dbus_method_invocation_get_method_name", &dBusMethodInvocationGetMethodName},
	{"g_dbus_method_invocation_get_object_path", &dBusMethodInvocationGetObjectPath},
	{"g_dbus_method_invocation_get_parameters", &dBusMethodInvocationGetParameters},
	{"g_dbus_method_invocation_get_sender", &dBusMethodInvocationGetSender},
	{"g_dbus_method_invocation_get_type", &DBusMethodInvocationGetType},
	{"g_dbus_method_invocation_get_user_data", &dBusMethodInvocationGetUserData},
	{"g_dbus_method_invocation_return_dbus_error", &dBusMethodInvocationReturnDBusError},
	{"g_dbus_method_invocation_return_error", &dBusMethodInvocationReturnError},
	{"g_dbus_method_invocation_return_error_literal", &dBusMethodInvocationReturnErrorLiteral},
	{"g_dbus_method_invocation_return_error_valist", &dBusMethodInvocationReturnErrorValist},
	{"g_dbus_method_invocation_return_gerror", &dBusMethodInvocationReturnGerror},
	{"g_dbus_method_invocation_return_value", &dBusMethodInvocationReturnValue},
	{"g_dbus_node_info_generate_xml", &dBusNodeInfoGenerateXml},
	{"g_dbus_node_info_get_type", &DBusNodeInfoGetType},
	{"g_dbus_node_info_lookup_interface", &dBusNodeInfoLookupInterface},
	{"g_dbus_node_info_new_for_xml", &DBusNodeInfoNewForXml},
	{"g_dbus_node_info_ref", &dBusNodeInfoRef},
	{"g_dbus_node_info_unref", &dBusNodeInfoUnref},
	{"g_dbus_property_info_flags_get_type", &DBusPropertyInfoFlagsGetType},
	{"g_dbus_property_info_get_type", &DBusPropertyInfoGetType},
	{"g_dbus_property_info_ref", &dBusPropertyInfoRef},
	{"g_dbus_property_info_unref", &dBusPropertyInfoUnref},
	{"g_dbus_proxy_call", &dBusProxyCall},
	{"g_dbus_proxy_call_finish", &dBusProxyCallFinish},
	{"g_dbus_proxy_call_sync", &dBusProxyCallSync},
	{"g_dbus_proxy_flags_get_type", &DBusProxyFlagsGetType},
	{"g_dbus_proxy_get_cached_property", &dBusProxyGetCachedProperty},
	{"g_dbus_proxy_get_cached_property_names", &dBusProxyGetCachedPropertyNames},
	{"g_dbus_proxy_get_connection", &dBusProxyGetConnection},
	{"g_dbus_proxy_get_default_timeout", &dBusProxyGetDefaultTimeout},
	{"g_dbus_proxy_get_flags", &dBusProxyGetFlags},
	{"g_dbus_proxy_get_interface_info", &dBusProxyGetInterfaceInfo},
	{"g_dbus_proxy_get_interface_name", &dBusProxyGetInterfaceName},
	{"g_dbus_proxy_get_name", &dBusProxyGetName},
	{"g_dbus_proxy_get_name_owner", &dBusProxyGetNameOwner},
	{"g_dbus_proxy_get_object_path", &dBusProxyGetObjectPath},
	{"g_dbus_proxy_get_type", &DBusProxyGetType},
	{"g_dbus_proxy_new", &DBusProxyNew},
	{"g_dbus_proxy_new_finish", &DBusProxyNewFinish},
	{"g_dbus_proxy_new_for_bus", &DBusProxyNewForBus},
	{"g_dbus_proxy_new_for_bus_finish", &DBusProxyNewForBusFinish},
	{"g_dbus_proxy_new_for_bus_sync", &DBusProxyNewForBusSync},
	{"g_dbus_proxy_new_sync", &DBusProxyNewSync},
	{"g_dbus_proxy_set_cached_property", &dBusProxySetCachedProperty},
	{"g_dbus_proxy_set_default_timeout", &dBusProxySetDefaultTimeout},
	{"g_dbus_proxy_set_interface_info", &dBusProxySetInterfaceInfo},
	{"g_dbus_send_message_flags_get_type", &DBusSendMessageFlagsGetType},
	{"g_dbus_server_flags_get_type", &DBusServerFlagsGetType},
	{"g_dbus_server_get_client_address", &dBusServerGetClientAddress},
	{"g_dbus_server_get_flags", &dBusServerGetFlags},
	{"g_dbus_server_get_guid", &dBusServerGetGuid},
	{"g_dbus_server_get_type", &DBusServerGetType},
	{"g_dbus_server_is_active", &dBusServerIsActive},
	{"g_dbus_server_new_sync", &DBusServerNewSync},
	{"g_dbus_server_start", &dBusServerStart},
	{"g_dbus_server_stop", &dBusServerStop},
	{"g_dbus_signal_flags_get_type", &DBusSignalFlagsGetType},
	{"g_dbus_signal_info_get_type", &DBusSignalInfoGetType},
	{"g_dbus_signal_info_ref", &dBusSignalInfoRef},
	{"g_dbus_signal_info_unref", &dBusSignalInfoUnref},
	{"g_dbus_subtree_flags_get_type", &DBusSubtreeFlagsGetType},
	{"g_drive_can_eject", &driveCanEject},
	{"g_drive_can_poll_for_media", &driveCanPollForMedia},
	{"g_drive_can_start", &driveCanStart},
	{"g_drive_can_start_degraded", &driveCanStartDegraded},
	{"g_drive_can_stop", &driveCanStop},
	{"g_drive_eject", &driveEject},
	{"g_drive_eject_finish", &driveEjectFinish},
	{"g_drive_eject_with_operation", &driveEjectWithOperation},
	{"g_drive_eject_with_operation_finish", &driveEjectWithOperationFinish},
	{"g_drive_enumerate_identifiers", &driveEnumerateIdentifiers},
	{"g_drive_get_icon", &driveGetIcon},
	{"g_drive_get_identifier", &driveGetIdentifier},
	{"g_drive_get_name", &driveGetName},
	{"g_drive_get_start_stop_type", &driveGetStartStopType},
	{"g_drive_get_type", &DriveGetType},
	{"g_drive_get_volumes", &driveGetVolumes},
	{"g_drive_has_media", &driveHasMedia},
	{"g_drive_has_volumes", &driveHasVolumes},
	{"g_drive_is_media_check_automatic", &driveIsMediaCheckAutomatic},
	{"g_drive_is_media_removable", &driveIsMediaRemovable},
	{"g_drive_poll_for_media", &drivePollForMedia},
	{"g_drive_poll_for_media_finish", &drivePollForMediaFinish},
	{"g_drive_start", &driveStart},
	{"g_drive_start_finish", &driveStartFinish},
	{"g_drive_start_flags_get_type", &DriveStartFlagsGetType},
	{"g_drive_start_stop_type_get_type", &DriveStartStopTypeGetType},
	{"g_drive_stop", &driveStop},
	{"g_drive_stop_finish", &driveStopFinish},
	{"g_emblem_get_icon", &EmblemGetIcon},
	{"g_emblem_get_origin", &EmblemGetOrigin},
	{"g_emblem_get_type", &EmblemGetType},
	{"g_emblem_new", &EmblemNew},
	{"g_emblem_new_with_origin", &EmblemNewWithOrigin},
	{"g_emblem_origin_get_type", &EmblemOriginGetType},
	{"g_emblemed_icon_add_emblem", &EmblemedIconAddEmblem},
	{"g_emblemed_icon_clear_emblems", &EmblemedIconClearEmblems},
	{"g_emblemed_icon_get_emblems", &EmblemedIconGetEmblems},
	{"g_emblemed_icon_get_icon", &EmblemedIconGetIcon},
	{"g_emblemed_icon_get_type", &EmblemedIconGetType},
	{"g_emblemed_icon_new", &EmblemedIconNew},
	{"g_file_append_to", &fileAppendTo},
	{"g_file_append_to_async", &fileAppendToAsync},
	{"g_file_append_to_finish", &fileAppendToFinish},
	{"g_file_attribute_info_flags_get_type", &FileAttributeInfoFlagsGetType},
	{"g_file_attribute_info_list_add", &fileAttributeInfoListAdd},
	{"g_file_attribute_info_list_dup", &fileAttributeInfoListDup},
	{"g_file_attribute_info_list_get_type", &FileAttributeInfoListGetType},
	{"g_file_attribute_info_list_lookup", &fileAttributeInfoListLookup},
	{"g_file_attribute_info_list_new", &FileAttributeInfoListNew},
	{"g_file_attribute_info_list_ref", &fileAttributeInfoListRef},
	{"g_file_attribute_info_list_unref", &fileAttributeInfoListUnref},
	{"g_file_attribute_matcher_enumerate_namespace", &fileAttributeMatcherEnumerateNamespace},
	{"g_file_attribute_matcher_enumerate_next", &fileAttributeMatcherEnumerateNext},
	{"g_file_attribute_matcher_get_type", &FileAttributeMatcherGetType},
	{"g_file_attribute_matcher_matches", &fileAttributeMatcherMatches},
	{"g_file_attribute_matcher_matches_only", &fileAttributeMatcherMatchesOnly},
	{"g_file_attribute_matcher_new", &FileAttributeMatcherNew},
	{"g_file_attribute_matcher_ref", &fileAttributeMatcherRef},
	{"g_file_attribute_matcher_unref", &fileAttributeMatcherUnref},
	{"g_file_attribute_status_get_type", &FileAttributeStatusGetType},
	{"g_file_attribute_type_get_type", &FileAttributeTypeGetType},
	{"g_file_copy", &fileCopy},
	{"g_file_copy_async", &fileCopyAsync},
	{"g_file_copy_attributes", &fileCopyAttributes},
	{"g_file_copy_finish", &fileCopyFinish},
	{"g_file_copy_flags_get_type", &FileCopyFlagsGetType},
	{"g_file_create", &fileCreate},
	{"g_file_create_async", &fileCreateAsync},
	{"g_file_create_finish", &fileCreateFinish},
	{"g_file_create_flags_get_type", &FileCreateFlagsGetType},
	{"g_file_create_readwrite", &fileCreateReadwrite},
	{"g_file_create_readwrite_async", &fileCreateReadwriteAsync},
	{"g_file_create_readwrite_finish", &fileCreateReadwriteFinish},
	{"g_file_delete", &fileDelete},
	{"g_file_dup", &fileDup},
	{"g_file_eject_mountable", &fileEjectMountable},
	{"g_file_eject_mountable_finish", &fileEjectMountableFinish},
	{"g_file_eject_mountable_with_operation", &fileEjectMountableWithOperation},
	{"g_file_eject_mountable_with_operation_finish", &fileEjectMountableWithOperationFinish},
	{"g_file_enumerate_children", &fileEnumerateChildren},
	{"g_file_enumerate_children_async", &fileEnumerateChildrenAsync},
	{"g_file_enumerate_children_finish", &fileEnumerateChildrenFinish},
	{"g_file_enumerator_close", &fileEnumeratorClose},
	{"g_file_enumerator_close_async", &fileEnumeratorCloseAsync},
	{"g_file_enumerator_close_finish", &fileEnumeratorCloseFinish},
	{"g_file_enumerator_get_container", &fileEnumeratorGetContainer},
	{"g_file_enumerator_get_type", &FileEnumeratorGetType},
	{"g_file_enumerator_has_pending", &fileEnumeratorHasPending},
	{"g_file_enumerator_is_closed", &fileEnumeratorIsClosed},
	{"g_file_enumerator_next_file", &fileEnumeratorNextFile},
	{"g_file_enumerator_next_files_async", &fileEnumeratorNextFilesAsync},
	{"g_file_enumerator_next_files_finish", &fileEnumeratorNextFilesFinish},
	{"g_file_enumerator_set_pending", &fileEnumeratorSetPending},
	{"g_file_equal", &fileEqual},
	{"g_file_find_enclosing_mount", &fileFindEnclosingMount},
	{"g_file_find_enclosing_mount_async", &fileFindEnclosingMountAsync},
	{"g_file_find_enclosing_mount_finish", &fileFindEnclosingMountFinish},
	{"g_file_get_basename", &fileGetBasename},
	{"g_file_get_child", &fileGetChild},
	{"g_file_get_child_for_display_name", &fileGetChildForDisplayName},
	{"g_file_get_parent", &fileGetParent},
	{"g_file_get_parse_name", &fileGetParseName},
	{"g_file_get_path", &fileGetPath},
	{"g_file_get_relative_path", &fileGetRelativePath},
	{"g_file_get_type", &FileGetType},
	{"g_file_get_uri", &fileGetUri},
	{"g_file_get_uri_scheme", &fileGetUriScheme},
	{"g_file_has_parent", &fileHasParent},
	{"g_file_has_prefix", &fileHasPrefix},
	{"g_file_has_uri_scheme", &fileHasUriScheme},
	{"g_file_hash", &FileHash},
	{"g_file_icon_get_file", &fileIconGetFile},
	{"g_file_icon_get_type", &FileIconGetType},
	{"g_file_icon_new", &FileIconNew},
	{"g_file_info_clear_status", &fileInfoClearStatus},
	{"g_file_info_copy_into", &FileInfoCopyInto},
	{"g_file_info_dup", &FileInfoDup},
	{"g_file_info_get_attribute_as_string", &fileInfoGetAttributeAsString},
	{"g_file_info_get_attribute_boolean", &fileInfoGetAttributeBoolean},
	{"g_file_info_get_attribute_byte_string", &fileInfoGetAttributeByteString},
	{"g_file_info_get_attribute_data", &fileInfoGetAttributeData},
	{"g_file_info_get_attribute_int32", &fileInfoGetAttributeInt32},
	{"g_file_info_get_attribute_int64", &fileInfoGetAttributeInt64},
	{"g_file_info_get_attribute_object", &fileInfoGetAttributeObject},
	{"g_file_info_get_attribute_status", &fileInfoGetAttributeStatus},
	{"g_file_info_get_attribute_string", &fileInfoGetAttributeString},
	{"g_file_info_get_attribute_stringv", &fileInfoGetAttributeStringv},
	{"g_file_info_get_attribute_type", &fileInfoGetAttributeType},
	{"g_file_info_get_attribute_uint32", &fileInfoGetAttributeUint32},
	{"g_file_info_get_attribute_uint64", &fileInfoGetAttributeUint64},
	{"g_file_info_get_content_type", &fileInfoGetContentType},
	{"g_file_info_get_display_name", &fileInfoGetDisplayName},
	{"g_file_info_get_edit_name", &fileInfoGetEditName},
	{"g_file_info_get_etag", &fileInfoGetEtag},
	{"g_file_info_get_file_type", &fileInfoGetFileType},
	{"g_file_info_get_icon", &fileInfoGetIcon},
	{"g_file_info_get_is_backup", &fileInfoGetIsBackup},
	{"g_file_info_get_is_hidden", &fileInfoGetIsHidden},
	{"g_file_info_get_is_symlink", &fileInfoGetIsSymlink},
	{"g_file_info_get_modification_time", &fileInfoGetModificationTime},
	{"g_file_info_get_name", &fileInfoGetName},
	{"g_file_info_get_size", &fileInfoGetSize},
	{"g_file_info_get_sort_order", &fileInfoGetSortOrder},
	{"g_file_info_get_symlink_target", &fileInfoGetSymlinkTarget},
	{"g_file_info_get_type", &FileInfoGetType},
	{"g_file_info_has_attribute", &fileInfoHasAttribute},
	{"g_file_info_has_namespace", &fileInfoHasNamespace},
	{"g_file_info_list_attributes", &fileInfoListAttributes},
	{"g_file_info_new", &FileInfoNew},
	{"g_file_info_remove_attribute", &fileInfoRemoveAttribute},
	{"g_file_info_set_attribute", &fileInfoSetAttribute},
	{"g_file_info_set_attribute_boolean", &fileInfoSetAttributeBoolean},
	{"g_file_info_set_attribute_byte_string", &fileInfoSetAttributeByteString},
	{"g_file_info_set_attribute_int32", &fileInfoSetAttributeInt32},
	{"g_file_info_set_attribute_int64", &fileInfoSetAttributeInt64},
	{"g_file_info_set_attribute_mask", &fileInfoSetAttributeMask},
	{"g_file_info_set_attribute_object", &fileInfoSetAttributeObject},
	{"g_file_info_set_attribute_status", &fileInfoSetAttributeStatus},
	{"g_file_info_set_attribute_string", &fileInfoSetAttributeString},
	{"g_file_info_set_attribute_stringv", &fileInfoSetAttributeStringv},
	{"g_file_info_set_attribute_uint32", &fileInfoSetAttributeUint32},
	{"g_file_info_set_attribute_uint64", &fileInfoSetAttributeUint64},
	{"g_file_info_set_content_type", &fileInfoSetContentType},
	{"g_file_info_set_display_name", &fileInfoSetDisplayName},
	{"g_file_info_set_edit_name", &fileInfoSetEditName},
	{"g_file_info_set_file_type", &fileInfoSetFileType},
	{"g_file_info_set_icon", &fileInfoSetIcon},
	{"g_file_info_set_is_hidden", &fileInfoSetIsHidden},
	{"g_file_info_set_is_symlink", &fileInfoSetIsSymlink},
	{"g_file_info_set_modification_time", &fileInfoSetModificationTime},
	{"g_file_info_set_name", &fileInfoSetName},
	{"g_file_info_set_size", &fileInfoSetSize},
	{"g_file_info_set_sort_order", &fileInfoSetSortOrder},
	{"g_file_info_set_symlink_target", &fileInfoSetSymlinkTarget},
	{"g_file_info_unset_attribute_mask", &fileInfoUnsetAttributeMask},
	{"g_file_input_stream_get_type", &FileInputStreamGetType},
	{"g_file_input_stream_query_info", &fileInputStreamQueryInfo},
	{"g_file_input_stream_query_info_async", &fileInputStreamQueryInfoAsync},
	{"g_file_input_stream_query_info_finish", &fileInputStreamQueryInfoFinish},
	{"g_file_io_stream_get_etag", &fileIoStreamGetEtag},
	{"g_file_io_stream_get_type", &FileIoStreamGetType},
	{"g_file_io_stream_query_info", &fileIoStreamQueryInfo},
	{"g_file_io_stream_query_info_async", &fileIoStreamQueryInfoAsync},
	{"g_file_io_stream_query_info_finish", &fileIoStreamQueryInfoFinish},
	{"g_file_is_native", &fileIsNative},
	{"g_file_load_contents", &fileLoadContents},
	{"g_file_load_contents_async", &fileLoadContentsAsync},
	{"g_file_load_contents_finish", &fileLoadContentsFinish},
	{"g_file_load_partial_contents_async", &fileLoadPartialContentsAsync},
	{"g_file_load_partial_contents_finish", &fileLoadPartialContentsFinish},
	{"g_file_make_directory", &fileMakeDirectory},
	{"g_file_make_directory_with_parents", &fileMakeDirectoryWithParents},
	{"g_file_make_symbolic_link", &fileMakeSymbolicLink},
	{"g_file_monitor", &fileMonitor},
	{"g_file_monitor_cancel", &fileMonitorCancel},
	{"g_file_monitor_directory", &fileMonitorDirectory},
	{"g_file_monitor_emit_event", &fileMonitorEmitEvent},
	{"g_file_monitor_event_get_type", &FileMonitorEventGetType},
	{"g_file_monitor_file", &fileMonitorFile},
	{"g_file_monitor_flags_get_type", &FileMonitorFlagsGetType},
	{"g_file_monitor_get_type", &FileMonitorGetType},
	{"g_file_monitor_is_cancelled", &fileMonitorIsCancelled},
	{"g_file_monitor_set_rate_limit", &fileMonitorSetRateLimit},
	{"g_file_mount_enclosing_volume", &fileMountEnclosingVolume},
	{"g_file_mount_enclosing_volume_finish", &fileMountEnclosingVolumeFinish},
	{"g_file_mount_mountable", &fileMountMountable},
	{"g_file_mount_mountable_finish", &fileMountMountableFinish},
	{"g_file_move", &fileMove},
	{"g_file_new_for_commandline_arg", &FileNewForCommandlineArg},
	{"g_file_new_for_path", &FileNewForPath},
	{"g_file_new_for_uri", &FileNewForUri},
	{"g_file_open_readwrite", &fileOpenReadwrite},
	{"g_file_open_readwrite_async", &fileOpenReadwriteAsync},
	{"g_file_open_readwrite_finish", &fileOpenReadwriteFinish},
	{"g_file_output_stream_get_etag", &fileOutputStreamGetEtag},
	{"g_file_output_stream_get_type", &FileOutputStreamGetType},
	{"g_file_output_stream_query_info", &fileOutputStreamQueryInfo},
	{"g_file_output_stream_query_info_async", &fileOutputStreamQueryInfoAsync},
	{"g_file_output_stream_query_info_finish", &fileOutputStreamQueryInfoFinish},
	{"g_file_parse_name", &FileParseName},
	{"g_file_poll_mountable", &filePollMountable},
	{"g_file_poll_mountable_finish", &filePollMountableFinish},
	{"g_file_query_default_handler", &fileQueryDefaultHandler},
	{"g_file_query_exists", &fileQueryExists},
	{"g_file_query_file_type", &fileQueryFileType},
	{"g_file_query_filesystem_info", &fileQueryFilesystemInfo},
	{"g_file_query_filesystem_info_async", &fileQueryFilesystemInfoAsync},
	{"g_file_query_filesystem_info_finish", &fileQueryFilesystemInfoFinish},
	{"g_file_query_info", &fileQueryInfo},
	{"g_file_query_info_async", &fileQueryInfoAsync},
	{"g_file_query_info_finish", &fileQueryInfoFinish},
	{"g_file_query_info_flags_get_type", &FileQueryInfoFlagsGetType},
	{"g_file_query_settable_attributes", &fileQuerySettableAttributes},
	{"g_file_query_writable_namespaces", &fileQueryWritableNamespaces},
	{"g_file_read", &fileRead},
	{"g_file_read_async", &fileReadAsync},
	{"g_file_read_finish", &fileReadFinish},
	{"g_file_replace", &fileReplace},
	{"g_file_replace_async", &fileReplaceAsync},
	{"g_file_replace_contents", &fileReplaceContents},
	{"g_file_replace_contents_async", &fileReplaceContentsAsync},
	{"g_file_replace_contents_finish", &fileReplaceContentsFinish},
	{"g_file_replace_finish", &fileReplaceFinish},
	{"g_file_replace_readwrite", &fileReplaceReadwrite},
	{"g_file_replace_readwrite_async", &fileReplaceReadwriteAsync},
	{"g_file_replace_readwrite_finish", &fileReplaceReadwriteFinish},
	{"g_file_resolve_relative_path", &fileResolveRelativePath},
	{"g_file_set_attribute", &fileSetAttribute},
	{"g_file_set_attribute_byte_string", &fileSetAttributeByteString},
	{"g_file_set_attribute_int32", &fileSetAttributeInt32},
	{"g_file_set_attribute_int64", &fileSetAttributeInt64},
	{"g_file_set_attribute_string", &fileSetAttributeString},
	{"g_file_set_attribute_uint32", &fileSetAttributeUint32},
	{"g_file_set_attribute_uint64", &fileSetAttributeUint64},
	{"g_file_set_attributes_async", &fileSetAttributesAsync},
	{"g_file_set_attributes_finish", &fileSetAttributesFinish},
	{"g_file_set_attributes_from_info", &fileSetAttributesFromInfo},
	{"g_file_set_display_name", &fileSetDisplayName},
	{"g_file_set_display_name_async", &fileSetDisplayNameAsync},
	{"g_file_set_display_name_finish", &fileSetDisplayNameFinish},
	{"g_file_start_mountable", &fileStartMountable},
	{"g_file_start_mountable_finish", &fileStartMountableFinish},
	{"g_file_stop_mountable", &fileStopMountable},
	{"g_file_stop_mountable_finish", &fileStopMountableFinish},
	{"g_file_supports_thread_contexts", &fileSupportsThreadContexts},
	{"g_file_trash", &fileTrash},
	{"g_file_type_get_type", &FileTypeGetType},
	{"g_file_unmount_mountable", &fileUnmountMountable},
	{"g_file_unmount_mountable_finish", &fileUnmountMountableFinish},
	{"g_file_unmount_mountable_with_operation", &fileUnmountMountableWithOperation},
	{"g_file_unmount_mountable_with_operation_finish", &fileUnmountMountableWithOperationFinish},
	{"g_filename_completer_get_completion_suffix", &filenameCompleterGetCompletionSuffix},
	{"g_filename_completer_get_completions", &filenameCompleterGetCompletions},
	{"g_filename_completer_get_type", &FilenameCompleterGetType},
	{"g_filename_completer_new", &FilenameCompleterNew},
	{"g_filename_completer_set_dirs_only", &filenameCompleterSetDirsOnly},
	{"g_filesystem_preview_type_get_type", &FilesystemPreviewTypeGetType},
	{"g_filter_input_stream_get_base_stream", &filterInputStreamGetBaseStream},
	{"g_filter_input_stream_get_close_base_stream", &filterInputStreamGetCloseBaseStream},
	{"g_filter_input_stream_get_type", &FilterInputStreamGetType},
	{"g_filter_input_stream_set_close_base_stream", &filterInputStreamSetCloseBaseStream},
	{"g_filter_output_stream_get_base_stream", &filterOutputStreamGetBaseStream},
	{"g_filter_output_stream_get_close_base_stream", &filterOutputStreamGetCloseBaseStream},
	{"g_filter_output_stream_get_type", &FilterOutputStreamGetType},
	{"g_filter_output_stream_set_close_base_stream", &filterOutputStreamSetCloseBaseStream},
	{"g_icon_equal", &IconEqual},
	{"g_icon_get_type", &IconGetType},
	{"g_icon_hash", &IconHash},
	{"g_icon_new_for_string", &IconNewForString},
	{"g_icon_to_string", &IconToString},
	{"g_inet_address_get_family", &inetAddressGetFamily},
	{"g_inet_address_get_is_any", &inetAddressGetIsAny},
	{"g_inet_address_get_is_link_local", &inetAddressGetIsLinkLocal},
	{"g_inet_address_get_is_loopback", &inetAddressGetIsLoopback},
	{"g_inet_address_get_is_mc_global", &inetAddressGetIsMcGlobal},
	{"g_inet_address_get_is_mc_link_local", &inetAddressGetIsMcLinkLocal},
	{"g_inet_address_get_is_mc_node_local", &inetAddressGetIsMcNodeLocal},
	{"g_inet_address_get_is_mc_org_local", &inetAddressGetIsMcOrgLocal},
	{"g_inet_address_get_is_mc_site_local", &inetAddressGetIsMcSiteLocal},
	{"g_inet_address_get_is_multicast", &inetAddressGetIsMulticast},
	{"g_inet_address_get_is_site_local", &inetAddressGetIsSiteLocal},
	{"g_inet_address_get_native_size", &inetAddressGetNativeSize},
	{"g_inet_address_get_type", &InetAddressGetType},
	{"g_inet_address_new_any", &InetAddressNewAny},
	{"g_inet_address_new_from_bytes", &InetAddressNewFromBytes},
	{"g_inet_address_new_from_string", &InetAddressNewFromString},
	{"g_inet_address_new_loopback", &InetAddressNewLoopback},
	{"g_inet_address_to_bytes", &inetAddressToBytes},
	{"g_inet_address_to_string", &inetAddressToString},
	{"g_inet_socket_address_get_address", &inetSocketAddressGetAddress},
	{"g_inet_socket_address_get_port", &inetSocketAddressGetPort},
	{"g_inet_socket_address_get_type", &InetSocketAddressGetType},
	{"g_inet_socket_address_new", &InetSocketAddressNew},
	{"g_initable_get_type", &InitableGetType},
	{"g_initable_init", &initableInit},
	{"g_initable_new", &InitableNew},
	{"g_initable_new_valist", &InitableNewValist},
	{"g_initable_newv", &InitableNewv},
	{"g_input_stream_clear_pending", &inputStreamClearPending},
	{"g_input_stream_close", &inputStreamClose},
	{"g_input_stream_close_async", &inputStreamCloseAsync},
	{"g_input_stream_close_finish", &inputStreamCloseFinish},
	{"g_input_stream_get_type", &InputStreamGetType},
	{"g_input_stream_has_pending", &inputStreamHasPending},
	{"g_input_stream_is_closed", &inputStreamIsClosed},
	{"g_input_stream_read", &inputStreamRead},
	{"g_input_stream_read_all", &inputStreamReadAll},
	{"g_input_stream_read_async", &inputStreamReadAsync},
	{"g_input_stream_read_finish", &inputStreamReadFinish},
	{"g_input_stream_set_pending", &inputStreamSetPending},
	{"g_input_stream_skip", &inputStreamSkip},
	{"g_input_stream_skip_async", &inputStreamSkipAsync},
	{"g_input_stream_skip_finish", &inputStreamSkipFinish},
	{"g_io_error_enum_get_type", &IoErrorEnumGetType},
	{"g_io_error_from_errno", &IoErrorFromErrno},
	{"g_io_error_from_win32_error", &IoErrorFromWin32Error},
	{"g_io_error_quark", &IoErrorQuark},
	{"g_io_extension_get_name", &ioExtensionGetName},
	{"g_io_extension_get_priority", &ioExtensionGetPriority},
	{"g_io_extension_get_type", &ioExtensionGetType},
	{"g_io_extension_point_get_extension_by_name", &ioExtensionPointGetExtensionByName},
	{"g_io_extension_point_get_extensions", &ioExtensionPointGetExtensions},
	{"g_io_extension_point_get_required_type", &ioExtensionPointGetRequiredType},
	{"g_io_extension_point_implement", &IoExtensionPointImplement},
	{"g_io_extension_point_lookup", &IoExtensionPointLookup},
	{"g_io_extension_point_register", &IoExtensionPointRegister},
	{"g_io_extension_point_set_required_type", &ioExtensionPointSetRequiredType},
	{"g_io_extension_ref_class", &ioExtensionRefClass},
	{"g_io_module_get_type", &IoModuleGetType},
	{"g_io_module_new", &IoModuleNew},
	{"g_io_modules_load_all_in_directory", &IoModulesLoadAllInDirectory},
	{"g_io_modules_scan_all_in_directory", &IoModulesScanAllInDirectory},
	{"g_io_scheduler_cancel_all_jobs", &IoSchedulerCancelAllJobs},
	{"g_io_scheduler_job_send_to_mainloop", &ioSchedulerJobSendToMainloop},
	{"g_io_scheduler_job_send_to_mainloop_async", &ioSchedulerJobSendToMainloopAsync},
	{"g_io_scheduler_push_job", &IoSchedulerPushJob},
	{"g_io_stream_clear_pending", &ioStreamClearPending},
	{"g_io_stream_close", &ioStreamClose},
	{"g_io_stream_close_async", &ioStreamCloseAsync},
	{"g_io_stream_close_finish", &ioStreamCloseFinish},
	{"g_io_stream_get_input_stream", &ioStreamGetInputStream},
	{"g_io_stream_get_output_stream", &ioStreamGetOutputStream},
	{"g_io_stream_get_type", &IoStreamGetType},
	{"g_io_stream_has_pending", &ioStreamHasPending},
	{"g_io_stream_is_closed", &ioStreamIsClosed},
	{"g_io_stream_set_pending", &ioStreamSetPending},
	{"g_io_stream_splice_async", &ioStreamSpliceAsync},
	{"g_io_stream_splice_finish", &IoStreamSpliceFinish},
	{"g_io_stream_splice_flags_get_type", &IoStreamSpliceFlagsGetType},
	{"g_keyfile_settings_backend_new", &KeyfileSettingsBackendNew},
	{"g_loadable_icon_get_type", &LoadableIconGetType},
	{"g_loadable_icon_load", &LoadableIconLoad},
	{"g_loadable_icon_load_async", &LoadableIconLoadAsync},
	{"g_loadable_icon_load_finish", &LoadableIconLoadFinish},
	{"g_memory_input_stream_add_data", &MemoryInputStreamAddData},
	{"g_memory_input_stream_get_type", &MemoryInputStreamGetType},
	{"g_memory_input_stream_new", &MemoryInputStreamNew},
	{"g_memory_input_stream_new_from_data", &MemoryInputStreamNewFromData},
	{"g_memory_output_stream_get_data", &MemoryOutputStreamGetData},
	{"g_memory_output_stream_get_data_size", &MemoryOutputStreamGetDataSize},
	{"g_memory_output_stream_get_size", &MemoryOutputStreamGetSize},
	{"g_memory_output_stream_get_type", &MemoryOutputStreamGetType},
	{"g_memory_output_stream_new", &MemoryOutputStreamNew},
	{"g_memory_output_stream_steal_data", &MemoryOutputStreamStealData},
	{"g_memory_settings_backend_new", &MemorySettingsBackendNew},
	{"g_mount_can_eject", &MountCanEject},
	{"g_mount_can_unmount", &MountCanUnmount},
	{"g_mount_eject", &MountEject},
	{"g_mount_eject_finish", &MountEjectFinish},
	{"g_mount_eject_with_operation", &MountEjectWithOperation},
	{"g_mount_eject_with_operation_finish", &MountEjectWithOperationFinish},
	{"g_mount_get_default_location", &MountGetDefaultLocation},
	{"g_mount_get_drive", &MountGetDrive},
	{"g_mount_get_icon", &MountGetIcon},
	{"g_mount_get_name", &MountGetName},
	{"g_mount_get_root", &MountGetRoot},
	{"g_mount_get_type", &MountGetType},
	{"g_mount_get_uuid", &MountGetUuid},
	{"g_mount_get_volume", &MountGetVolume},
	{"g_mount_guess_content_type", &MountGuessContentType},
	{"g_mount_guess_content_type_finish", &MountGuessContentTypeFinish},
	{"g_mount_guess_content_type_sync", &MountGuessContentTypeSync},
	{"g_mount_is_shadowed", &MountIsShadowed},
	{"g_mount_mount_flags_get_type", &MountMountFlagsGetType},
	{"g_mount_operation_get_anonymous", &MountOperationGetAnonymous},
	{"g_mount_operation_get_choice", &MountOperationGetChoice},
	{"g_mount_operation_get_domain", &MountOperationGetDomain},
	{"g_mount_operation_get_password", &MountOperationGetPassword},
	{"g_mount_operation_get_password_save", &MountOperationGetPasswordSave},
	{"g_mount_operation_get_type", &MountOperationGetType},
	{"g_mount_operation_get_username", &MountOperationGetUsername},
	{"g_mount_operation_new", &MountOperationNew},
	{"g_mount_operation_reply", &MountOperationReply},
	{"g_mount_operation_result_get_type", &MountOperationResultGetType},
	{"g_mount_operation_set_anonymous", &MountOperationSetAnonymous},
	{"g_mount_operation_set_choice", &MountOperationSetChoice},
	{"g_mount_operation_set_domain", &MountOperationSetDomain},
	{"g_mount_operation_set_password", &MountOperationSetPassword},
	{"g_mount_operation_set_password_save", &MountOperationSetPasswordSave},
	{"g_mount_operation_set_username", &MountOperationSetUsername},
	{"g_mount_remount", &MountRemount},
	{"g_mount_remount_finish", &MountRemountFinish},
	{"g_mount_shadow", &MountShadow},
	{"g_mount_unmount", &MountUnmount},
	{"g_mount_unmount_finish", &MountUnmountFinish},
	{"g_mount_unmount_flags_get_type", &MountUnmountFlagsGetType},
	{"g_mount_unmount_with_operation", &MountUnmountWithOperation},
	{"g_mount_unmount_with_operation_finish", &MountUnmountWithOperationFinish},
	{"g_mount_unshadow", &MountUnshadow},
	{"g_native_volume_monitor_get_type", &NativeVolumeMonitorGetType},
	{"g_network_address_get_hostname", &NetworkAddressGetHostname},
	{"g_network_address_get_port", &NetworkAddressGetPort},
	{"g_network_address_get_scheme", &NetworkAddressGetScheme},
	{"g_network_address_get_type", &NetworkAddressGetType},
	{"g_network_address_new", &NetworkAddressNew},
	{"g_network_address_parse", &NetworkAddressParse},
	{"g_network_address_parse_uri", &NetworkAddressParseUri},
	{"g_network_service_get_domain", &NetworkServiceGetDomain},
	{"g_network_service_get_protocol", &NetworkServiceGetProtocol},
	{"g_network_service_get_scheme", &NetworkServiceGetScheme},
	{"g_network_service_get_service", &NetworkServiceGetService},
	{"g_network_service_get_type", &NetworkServiceGetType},
	{"g_network_service_new", &NetworkServiceNew},
	{"g_network_service_set_scheme", &NetworkServiceSetScheme},
	{"g_null_settings_backend_new", &NullSettingsBackendNew},
	{"g_output_stream_clear_pending", &outputStreamClearPending},
	{"g_output_stream_close", &outputStreamClose},
	{"g_output_stream_close_async", &outputStreamCloseAsync},
	{"g_output_stream_close_finish", &outputStreamCloseFinish},
	{"g_output_stream_flush", &outputStreamFlush},
	{"g_output_stream_flush_async", &outputStreamFlushAsync},
	{"g_output_stream_flush_finish", &outputStreamFlushFinish},
	{"g_output_stream_get_type", &OutputStreamGetType},
	{"g_output_stream_has_pending", &outputStreamHasPending},
	{"g_output_stream_is_closed", &outputStreamIsClosed},
	{"g_output_stream_is_closing", &outputStreamIsClosing},
	{"g_output_stream_set_pending", &outputStreamSetPending},
	{"g_output_stream_splice", &outputStreamSplice},
	{"g_output_stream_splice_async", &outputStreamSpliceAsync},
	{"g_output_stream_splice_finish", &outputStreamSpliceFinish},
	{"g_output_stream_splice_flags_get_type", &OutputStreamSpliceFlagsGetType},
	{"g_output_stream_write", &outputStreamWrite},
	{"g_output_stream_write_all", &outputStreamWriteAll},
	{"g_output_stream_write_async", &outputStreamWriteAsync},
	{"g_output_stream_write_finish", &outputStreamWriteFinish},
	{"g_password_save_get_type", &PasswordSaveGetType},
	{"g_permission_acquire", &PermissionAcquire},
	{"g_permission_acquire_async", &PermissionAcquireAsync},
	{"g_permission_acquire_finish", &PermissionAcquireFinish},
	{"g_permission_get_allowed", &PermissionGetAllowed},
	{"g_permission_get_can_acquire", &PermissionGetCanAcquire},
	{"g_permission_get_can_release", &PermissionGetCanRelease},
	{"g_permission_get_type", &PermissionGetType},
	{"g_permission_impl_update", &PermissionImplUpdate},
	{"g_permission_release", &PermissionRelease},
	{"g_permission_release_async", &PermissionReleaseAsync},
	{"g_permission_release_finish", &PermissionReleaseFinish},
	{"g_pollable_input_stream_can_poll", &PollableInputStreamCanPoll},
	{"g_pollable_input_stream_create_source", &PollableInputStreamCreateSource},
	{"g_pollable_input_stream_get_type", &PollableInputStreamGetType},
	{"g_pollable_input_stream_is_readable", &PollableInputStreamIsReadable},
	{"g_pollable_input_stream_read_nonblocking", &PollableInputStreamReadNonblocking},
	{"g_pollable_output_stream_can_poll", &PollableOutputStreamCanPoll},
	{"g_pollable_output_stream_create_source", &PollableOutputStreamCreateSource},
	{"g_pollable_output_stream_get_type", &PollableOutputStreamGetType},
	{"g_pollable_output_stream_is_writable", &PollableOutputStreamIsWritable},
	{"g_pollable_output_stream_write_nonblocking", &PollableOutputStreamWriteNonblocking},
	{"g_pollable_source_new", &PollableSourceNew},
	{"g_proxy_address_enumerator_get_type", &ProxyAddressEnumeratorGetType},
	{"g_proxy_address_get_destination_hostname", &ProxyAddressGetDestinationHostname},
	{"g_proxy_address_get_destination_port", &ProxyAddressGetDestinationPort},
	{"g_proxy_address_get_password", &ProxyAddressGetPassword},
	{"g_proxy_address_get_protocol", &ProxyAddressGetProtocol},
	{"g_proxy_address_get_type", &ProxyAddressGetType},
	{"g_proxy_address_get_username", &ProxyAddressGetUsername},
	{"g_proxy_address_new", &ProxyAddressNew},
	{"g_proxy_connect", &ProxyConnect},
	{"g_proxy_connect_async", &ProxyConnectAsync},
	{"g_proxy_connect_finish", &ProxyConnectFinish},
	{"g_proxy_get_default_for_protocol", &ProxyGetDefaultForProtocol},
	{"g_proxy_get_type", &ProxyGetType},
	{"g_proxy_resolver_get_default", &ProxyResolverGetDefault},
	{"g_proxy_resolver_get_type", &ProxyResolverGetType},
	{"g_proxy_resolver_is_supported", &ProxyResolverIsSupported},
	{"g_proxy_resolver_lookup", &ProxyResolverLookup},
	{"g_proxy_resolver_lookup_async", &ProxyResolverLookupAsync},
	{"g_proxy_resolver_lookup_finish", &ProxyResolverLookupFinish},
	{"g_proxy_supports_hostname", &ProxySupportsHostname},
	{"g_resolver_error_get_type", &ResolverErrorGetType},
	{"g_resolver_error_quark", &ResolverErrorQuark},
	{"g_resolver_free_addresses", &ResolverFreeAddresses},
	{"g_resolver_free_targets", &ResolverFreeTargets},
	{"g_resolver_get_default", &ResolverGetDefault},
	{"g_resolver_get_type", &ResolverGetType},
	{"g_resolver_lookup_by_address", &ResolverLookupByAddress},
	{"g_resolver_lookup_by_address_async", &ResolverLookupByAddressAsync},
	{"g_resolver_lookup_by_address_finish", &ResolverLookupByAddressFinish},
	{"g_resolver_lookup_by_name", &ResolverLookupByName},
	{"g_resolver_lookup_by_name_async", &ResolverLookupByNameAsync},
	{"g_resolver_lookup_by_name_finish", &ResolverLookupByNameFinish},
	{"g_resolver_lookup_service", &ResolverLookupService},
	{"g_resolver_lookup_service_async", &ResolverLookupServiceAsync},
	{"g_resolver_lookup_service_finish", &ResolverLookupServiceFinish},
	{"g_resolver_set_default", &ResolverSetDefault},
	{"g_seekable_can_seek", &seekableCanSeek},
	{"g_seekable_can_truncate", &seekableCanTruncate},
	{"g_seekable_get_type", &SeekableGetType},
	{"g_seekable_seek", &seekableSeek},
	{"g_seekable_tell", &seekableTell},
	{"g_seekable_truncate", &seekableTruncate},
	{"g_settings_apply", &settingsApply},
	{"g_settings_backend_changed", &settingsBackendChanged},
	{"g_settings_backend_changed_tree", &settingsBackendChangedTree},
	{"g_settings_backend_flatten_tree", &SettingsBackendFlattenTree},
	{"g_settings_backend_get_default", &SettingsBackendGetDefault},
	{"g_settings_backend_get_type", &SettingsBackendGetType},
	{"g_settings_backend_keys_changed", &settingsBackendKeysChanged},
	{"g_settings_backend_path_changed", &settingsBackendPathChanged},
	{"g_settings_backend_path_writable_changed", &settingsBackendPathWritableChanged},
	{"g_settings_backend_writable_changed", &settingsBackendWritableChanged},
	{"g_settings_bind", &settingsBind},
	{"g_settings_bind_flags_get_type", &SettingsBindFlagsGetType},
	{"g_settings_bind_with_mapping", &settingsBindWithMapping},
	{"g_settings_bind_writable", &settingsBindWritable},
	{"g_settings_delay", &settingsDelay},
	{"g_settings_get", &settingsGet},
	{"g_settings_get_boolean", &settingsGetBoolean},
	{"g_settings_get_child", &settingsGetChild},
	{"g_settings_get_double", &settingsGetDouble},
	{"g_settings_get_enum", &settingsGetEnum},
	{"g_settings_get_flags", &settingsGetFlags},
	{"g_settings_get_has_unapplied", &settingsGetHasUnapplied},
	{"g_settings_get_int", &settingsGetInt},
	{"g_settings_get_mapped", &settingsGetMapped},
	{"g_settings_get_range", &settingsGetRange},
	{"g_settings_get_string", &settingsGetString},
	{"g_settings_get_strv", &settingsGetStrv},
	{"g_settings_get_type", &SettingsGetType},
	{"g_settings_get_value", &settingsGetValue},
	{"g_settings_is_writable", &settingsIsWritable},
	{"g_settings_list_children", &settingsListChildren},
	{"g_settings_list_keys", &settingsListKeys},
	{"g_settings_list_relocatable_schemas", &SettingsListRelocatableSchemas},
	{"g_settings_list_schemas", &SettingsListSchemas},
	{"g_settings_new", &SettingsNew},
	{"g_settings_new_with_backend", &SettingsNewWithBackend},
	{"g_settings_new_with_backend_and_path", &SettingsNewWithBackendAndPath},
	{"g_settings_new_with_path", &SettingsNewWithPath},
	{"g_settings_range_check", &settingsRangeCheck},
	{"g_settings_reset", &settingsReset},
	{"g_settings_revert", &settingsRevert},
	{"g_settings_set", &settingsSet},
	{"g_settings_set_boolean", &settingsSetBoolean},
	{"g_settings_set_double", &settingsSetDouble},
	{"g_settings_set_enum", &settingsSetEnum},
	{"g_settings_set_flags", &settingsSetFlags},
	{"g_settings_set_int", &settingsSetInt},
	{"g_settings_set_string", &settingsSetString},
	{"g_settings_set_strv", &settingsSetStrv},
	{"g_settings_set_value", &settingsSetValue},
	{"g_settings_sync", &SettingsSync},
	{"g_settings_unbind", &SettingsUnbind},
	// Undocumented {"g_simple_action_get_parameter_type", &SimpleActionGetParameterType},
	{"g_simple_action_get_type", &SimpleActionGetType},
	{"g_simple_action_group_get_type", &SimpleActionGroupGetType},
	{"g_simple_action_group_insert", &simpleActionGroupInsert},
	{"g_simple_action_group_lookup", &simpleActionGroupLookup},
	{"g_simple_action_group_new", &SimpleActionGroupNew},
	{"g_simple_action_group_remove", &simpleActionGroupRemove},
	{"g_simple_action_new", &SimpleActionNew},
	{"g_simple_action_new_stateful", &SimpleActionNewStateful},
	{"g_simple_action_set_enabled", &simpleActionSetEnabled},
	{"g_simple_async_report_error_in_idle", &SimpleAsyncReportErrorInIdle},
	{"g_simple_async_report_gerror_in_idle", &SimpleAsyncReportGerrorInIdle},
	{"g_simple_async_report_take_gerror_in_idle", &SimpleAsyncReportTakeGerrorInIdle},
	{"g_simple_async_result_complete", &simpleAsyncResultComplete},
	{"g_simple_async_result_complete_in_idle", &simpleAsyncResultCompleteInIdle},
	{"g_simple_async_result_get_op_res_gboolean", &simpleAsyncResultGetOpResGboolean},
	{"g_simple_async_result_get_op_res_gpointer", &simpleAsyncResultGetOpResGpointer},
	{"g_simple_async_result_get_op_res_gssize", &simpleAsyncResultGetOpResGssize},
	{"g_simple_async_result_get_source_tag", &simpleAsyncResultGetSourceTag},
	{"g_simple_async_result_get_type", &SimpleAsyncResultGetType},
	{"g_simple_async_result_is_valid", &SimpleAsyncResultIsValid},
	{"g_simple_async_result_new", &SimpleAsyncResultNew},
	{"g_simple_async_result_new_error", &SimpleAsyncResultNewError},
	{"g_simple_async_result_new_from_error", &SimpleAsyncResultNewFromError},
	{"g_simple_async_result_new_take_error", &SimpleAsyncResultNewTakeError},
	{"g_simple_async_result_propagate_error", &simpleAsyncResultPropagateError},
	{"g_simple_async_result_run_in_thread", &simpleAsyncResultRunInThread},
	{"g_simple_async_result_set_error", &simpleAsyncResultSetError},
	{"g_simple_async_result_set_error_va", &simpleAsyncResultSetErrorVa},
	{"g_simple_async_result_set_from_error", &simpleAsyncResultSetFromError},
	{"g_simple_async_result_set_handle_cancellation", &simpleAsyncResultSetHandleCancellation},
	{"g_simple_async_result_set_op_res_gboolean", &simpleAsyncResultSetOpResGboolean},
	{"g_simple_async_result_set_op_res_gpointer", &simpleAsyncResultSetOpResGpointer},
	{"g_simple_async_result_set_op_res_gssize", &simpleAsyncResultSetOpResGssize},
	{"g_simple_async_result_take_error", &simpleAsyncResultTakeError},
	{"g_simple_permission_get_type", &SimplePermissionGetType},
	{"g_simple_permission_new", &SimplePermissionNew},
	{"g_socket_accept", &socketAccept},
	{"g_socket_address_enumerator_get_type", &SocketAddressEnumeratorGetType},
	{"g_socket_address_enumerator_next", &socketAddressEnumeratorNext},
	{"g_socket_address_enumerator_next_async", &socketAddressEnumeratorNextAsync},
	{"g_socket_address_enumerator_next_finish", &socketAddressEnumeratorNextFinish},
	{"g_socket_address_get_family", &socketAddressGetFamily},
	{"g_socket_address_get_native_size", &socketAddressGetNativeSize},
	{"g_socket_address_get_type", &SocketAddressGetType},
	{"g_socket_address_new_from_native", &SocketAddressNewFromNative},
	{"g_socket_address_to_native", &socketAddressToNative},
	{"g_socket_bind", &socketBind},
	{"g_socket_check_connect_result", &socketCheckConnectResult},
	{"g_socket_client_add_application_proxy", &socketClientAddApplicationProxy},
	{"g_socket_client_connect", &socketClientConnect},
	{"g_socket_client_connect_async", &socketClientConnectAsync},
	{"g_socket_client_connect_finish", &socketClientConnectFinish},
	{"g_socket_client_connect_to_host", &socketClientConnectToHost},
	{"g_socket_client_connect_to_host_async", &socketClientConnectToHostAsync},
	{"g_socket_client_connect_to_host_finish", &socketClientConnectToHostFinish},
	{"g_socket_client_connect_to_service", &socketClientConnectToService},
	{"g_socket_client_connect_to_service_async", &socketClientConnectToServiceAsync},
	{"g_socket_client_connect_to_service_finish", &socketClientConnectToServiceFinish},
	{"g_socket_client_connect_to_uri", &socketClientConnectToUri},
	{"g_socket_client_connect_to_uri_async", &socketClientConnectToUriAsync},
	{"g_socket_client_connect_to_uri_finish", &socketClientConnectToUriFinish},
	{"g_socket_client_get_enable_proxy", &socketClientGetEnableProxy},
	{"g_socket_client_get_family", &socketClientGetFamily},
	{"g_socket_client_get_local_address", &socketClientGetLocalAddress},
	{"g_socket_client_get_protocol", &socketClientGetProtocol},
	{"g_socket_client_get_socket_type", &socketClientGetSocketType},
	{"g_socket_client_get_timeout", &socketClientGetTimeout},
	{"g_socket_client_get_tls", &socketClientGetTls},
	{"g_socket_client_get_tls_validation_flags", &socketClientGetTlsValidationFlags},
	{"g_socket_client_get_type", &SocketClientGetType},
	{"g_socket_client_new", &SocketClientNew},
	{"g_socket_client_set_enable_proxy", &socketClientSetEnableProxy},
	{"g_socket_client_set_family", &socketClientSetFamily},
	{"g_socket_client_set_local_address", &socketClientSetLocalAddress},
	{"g_socket_client_set_protocol", &socketClientSetProtocol},
	{"g_socket_client_set_socket_type", &socketClientSetSocketType},
	{"g_socket_client_set_timeout", &socketClientSetTimeout},
	{"g_socket_client_set_tls", &socketClientSetTls},
	{"g_socket_client_set_tls_validation_flags", &socketClientSetTlsValidationFlags},
	{"g_socket_close", &socketClose},
	{"g_socket_condition_check", &socketConditionCheck},
	{"g_socket_condition_wait", &socketConditionWait},
	{"g_socket_connect", &socketConnect},
	{"g_socket_connectable_enumerate", &socketConnectableEnumerate},
	{"g_socket_connectable_get_type", &SocketConnectableGetType},
	{"g_socket_connectable_proxy_enumerate", &socketConnectableProxyEnumerate},
	{"g_socket_connection_factory_create_connection", &SocketConnectionFactoryCreateConnection},
	{"g_socket_connection_factory_lookup_type", &SocketConnectionFactoryLookupType},
	{"g_socket_connection_factory_register_type", &SocketConnectionFactoryRegisterType},
	{"g_socket_connection_get_local_address", &socketConnectionGetLocalAddress},
	{"g_socket_connection_get_remote_address", &socketConnectionGetRemoteAddress},
	{"g_socket_connection_get_socket", &socketConnectionGetSocket},
	{"g_socket_connection_get_type", &SocketConnectionGetType},
	{"g_socket_control_message_deserialize", &SocketControlMessageDeserialize},
	{"g_socket_control_message_get_level", &socketControlMessageGetLevel},
	{"g_socket_control_message_get_msg_type", &socketControlMessageGetMsgType},
	{"g_socket_control_message_get_size", &socketControlMessageGetSize},
	{"g_socket_control_message_get_type", &SocketControlMessageGetType},
	{"g_socket_control_message_serialize", &socketControlMessageSerialize},
	{"g_socket_create_source", &socketCreateSource},
	{"g_socket_family_get_type", &SocketFamilyGetType},
	{"g_socket_get_blocking", &socketGetBlocking},
	{"g_socket_get_credentials", &socketGetCredentials},
	{"g_socket_get_family", &socketGetFamily},
	{"g_socket_get_fd", &socketGetFd},
	{"g_socket_get_keepalive", &socketGetKeepalive},
	{"g_socket_get_listen_backlog", &socketGetListenBacklog},
	{"g_socket_get_local_address", &socketGetLocalAddress},
	{"g_socket_get_protocol", &socketGetProtocol},
	{"g_socket_get_remote_address", &socketGetRemoteAddress},
	{"g_socket_get_socket_type", &socketGetSocketType},
	{"g_socket_get_timeout", &socketGetTimeout},
	{"g_socket_get_type", &SocketGetType},
	{"g_socket_is_closed", &socketIsClosed},
	{"g_socket_is_connected", &socketIsConnected},
	{"g_socket_listen", &socketListen},
	{"g_socket_listener_accept", &socketListenerAccept},
	{"g_socket_listener_accept_async", &socketListenerAcceptAsync},
	{"g_socket_listener_accept_finish", &socketListenerAcceptFinish},
	{"g_socket_listener_accept_socket", &socketListenerAcceptSocket},
	{"g_socket_listener_accept_socket_async", &socketListenerAcceptSocketAsync},
	{"g_socket_listener_accept_socket_finish", &socketListenerAcceptSocketFinish},
	{"g_socket_listener_add_address", &socketListenerAddAddress},
	{"g_socket_listener_add_any_inet_port", &socketListenerAddAnyInetPort},
	{"g_socket_listener_add_inet_port", &socketListenerAddInetPort},
	{"g_socket_listener_add_socket", &socketListenerAddSocket},
	{"g_socket_listener_close", &socketListenerClose},
	{"g_socket_listener_get_type", &SocketListenerGetType},
	{"g_socket_listener_new", &SocketListenerNew},
	{"g_socket_listener_set_backlog", &socketListenerSetBacklog},
	{"g_socket_msg_flags_get_type", &SocketMsgFlagsGetType},
	{"g_socket_new", &SocketNew},
	{"g_socket_new_from_fd", &SocketNewFromFd},
	{"g_socket_protocol_get_type", &SocketProtocolGetType},
	{"g_socket_receive", &socketReceive},
	{"g_socket_receive_from", &socketReceiveFrom},
	{"g_socket_receive_message", &socketReceiveMessage},
	{"g_socket_receive_with_blocking", &socketReceiveWithBlocking},
	{"g_socket_send", &socketSend},
	{"g_socket_send_message", &socketSendMessage},
	{"g_socket_send_to", &socketSendTo},
	{"g_socket_send_with_blocking", &socketSendWithBlocking},
	{"g_socket_service_get_type", &SocketServiceGetType},
	{"g_socket_service_is_active", &socketServiceIsActive},
	{"g_socket_service_new", &SocketServiceNew},
	{"g_socket_service_start", &socketServiceStart},
	{"g_socket_service_stop", &socketServiceStop},
	{"g_socket_set_blocking", &socketSetBlocking},
	{"g_socket_set_keepalive", &socketSetKeepalive},
	{"g_socket_set_listen_backlog", &socketSetListenBacklog},
	{"g_socket_set_timeout", &socketSetTimeout},
	{"g_socket_shutdown", &socketShutdown},
	{"g_socket_speaks_ipv4", &socketSpeaksIpv4},
	{"g_socket_type_get_type", &SocketTypeGetType},
	{"g_srv_target_copy", &srvTargetCopy},
	{"g_srv_target_free", &srvTargetFree},
	{"g_srv_target_get_hostname", &srvTargetGetHostname},
	{"g_srv_target_get_port", &srvTargetGetPort},
	{"g_srv_target_get_priority", &srvTargetGetPriority},
	{"g_srv_target_get_type", &SrvTargetGetType},
	{"g_srv_target_get_weight", &srvTargetGetWeight},
	{"g_srv_target_list_sort", &SrvTargetListSort},
	{"g_srv_target_new", &SrvTargetNew},
	{"g_tcp_connection_get_graceful_disconnect", &TcpConnectionGetGracefulDisconnect},
	{"g_tcp_connection_get_type", &TcpConnectionGetType},
	{"g_tcp_connection_set_graceful_disconnect", &TcpConnectionSetGracefulDisconnect},
	{"g_tcp_wrapper_connection_get_base_io_stream", &TcpWrapperConnectionGetBaseIoStream},
	{"g_tcp_wrapper_connection_get_type", &TcpWrapperConnectionGetType},
	{"g_tcp_wrapper_connection_new", &TcpWrapperConnectionNew},
	{"g_themed_icon_append_name", &ThemedIconAppendName},
	{"g_themed_icon_get_names", &ThemedIconGetNames},
	{"g_themed_icon_get_type", &ThemedIconGetType},
	{"g_themed_icon_new", &ThemedIconNew},
	{"g_themed_icon_new_from_names", &ThemedIconNewFromNames},
	{"g_themed_icon_new_with_default_fallbacks", &ThemedIconNewWithDefaultFallbacks},
	{"g_themed_icon_prepend_name", &ThemedIconPrependName},
	// Undocumented {"g_threaded_resolver_get_type", &ThreadedResolverGetType},
	{"g_threaded_socket_service_get_type", &ThreadedSocketServiceGetType},
	{"g_threaded_socket_service_new", &ThreadedSocketServiceNew},
	{"g_tls_authentication_mode_get_type", &TlsAuthenticationModeGetType},
	{"g_tls_backend_get_certificate_type", &TlsBackendGetCertificateType},
	{"g_tls_backend_get_client_connection_type", &TlsBackendGetClientConnectionType},
	{"g_tls_backend_get_default", &TlsBackendGetDefault},
	{"g_tls_backend_get_server_connection_type", &TlsBackendGetServerConnectionType},
	{"g_tls_backend_get_type", &TlsBackendGetType},
	{"g_tls_backend_supports_tls", &TlsBackendSupportsTls},
	{"g_tls_certificate_flags_get_type", &TlsCertificateFlagsGetType},
	{"g_tls_certificate_get_issuer", &TlsCertificateGetIssuer},
	{"g_tls_certificate_get_type", &TlsCertificateGetType},
	{"g_tls_certificate_list_new_from_file", &TlsCertificateListNewFromFile},
	{"g_tls_certificate_new_from_file", &TlsCertificateNewFromFile},
	{"g_tls_certificate_new_from_files", &TlsCertificateNewFromFiles},
	{"g_tls_certificate_new_from_pem", &TlsCertificateNewFromPem},
	{"g_tls_certificate_verify", &TlsCertificateVerify},
	{"g_tls_client_connection_get_accepted_cas", &TlsClientConnectionGetAcceptedCas},
	{"g_tls_client_connection_get_server_identity", &TlsClientConnectionGetServerIdentity},
	{"g_tls_client_connection_get_type", &TlsClientConnectionGetType},
	{"g_tls_client_connection_get_use_ssl3", &TlsClientConnectionGetUseSsl3},
	{"g_tls_client_connection_get_validation_flags", &TlsClientConnectionGetValidationFlags},
	{"g_tls_client_connection_new", &TlsClientConnectionNew},
	{"g_tls_client_connection_set_server_identity", &TlsClientConnectionSetServerIdentity},
	{"g_tls_client_connection_set_use_ssl3", &TlsClientConnectionSetUseSsl3},
	{"g_tls_client_connection_set_validation_flags", &TlsClientConnectionSetValidationFlags},
	{"g_tls_connection_emit_accept_certificate", &TlsConnectionEmitAcceptCertificate},
	{"g_tls_connection_get_certificate", &TlsConnectionGetCertificate},
	{"g_tls_connection_get_peer_certificate", &TlsConnectionGetPeerCertificate},
	{"g_tls_connection_get_peer_certificate_errors", &TlsConnectionGetPeerCertificateErrors},
	{"g_tls_connection_get_rehandshake_mode", &TlsConnectionGetRehandshakeMode},
	{"g_tls_connection_get_require_close_notify", &TlsConnectionGetRequireCloseNotify},
	{"g_tls_connection_get_type", &TlsConnectionGetType},
	{"g_tls_connection_get_use_system_certdb", &TlsConnectionGetUseSystemCertdb},
	{"g_tls_connection_handshake", &TlsConnectionHandshake},
	{"g_tls_connection_handshake_async", &TlsConnectionHandshakeAsync},
	{"g_tls_connection_handshake_finish", &TlsConnectionHandshakeFinish},
	{"g_tls_connection_set_certificate", &TlsConnectionSetCertificate},
	{"g_tls_connection_set_rehandshake_mode", &TlsConnectionSetRehandshakeMode},
	{"g_tls_connection_set_require_close_notify", &TlsConnectionSetRequireCloseNotify},
	{"g_tls_connection_set_use_system_certdb", &TlsConnectionSetUseSystemCertdb},
	{"g_tls_error_get_type", &TlsErrorGetType},
	{"g_tls_error_quark", &TlsErrorQuark},
	{"g_tls_rehandshake_mode_get_type", &TlsRehandshakeModeGetType},
	{"g_tls_server_connection_get_type", &TlsServerConnectionGetType},
	{"g_tls_server_connection_new", &TlsServerConnectionNew},
	{"g_unix_socket_address_type_get_type", &UnixSocketAddressTypeGetType},
	{"g_vfs_get_default", &VfsGetDefault},
	{"g_vfs_get_file_for_path", &VfsGetFileForPath},
	{"g_vfs_get_file_for_uri", &VfsGetFileForUri},
	{"g_vfs_get_local", &VfsGetLocal},
	{"g_vfs_get_supported_uri_schemes", &VfsGetSupportedUriSchemes},
	{"g_vfs_get_type", &VfsGetType},
	{"g_vfs_is_active", &VfsIsActive},
	{"g_vfs_parse_name", &VfsParseName},
	{"g_volume_can_eject", &VolumeCanEject},
	{"g_volume_can_mount", &VolumeCanMount},
	{"g_volume_eject", &VolumeEject},
	{"g_volume_eject_finish", &VolumeEjectFinish},
	{"g_volume_eject_with_operation", &VolumeEjectWithOperation},
	{"g_volume_eject_with_operation_finish", &VolumeEjectWithOperationFinish},
	{"g_volume_enumerate_identifiers", &VolumeEnumerateIdentifiers},
	{"g_volume_get_activation_root", &VolumeGetActivationRoot},
	{"g_volume_get_drive", &VolumeGetDrive},
	{"g_volume_get_icon", &VolumeGetIcon},
	{"g_volume_get_identifier", &VolumeGetIdentifier},
	{"g_volume_get_mount", &VolumeGetMount},
	{"g_volume_get_name", &VolumeGetName},
	{"g_volume_get_type", &VolumeGetType},
	{"g_volume_get_uuid", &VolumeGetUuid},
	{"g_volume_monitor_adopt_orphan_mount", &VolumeMonitorAdoptOrphanMount},
	{"g_volume_monitor_get", &VolumeMonitorGet},
	{"g_volume_monitor_get_connected_drives", &VolumeMonitorGetConnectedDrives},
	{"g_volume_monitor_get_mount_for_uuid", &VolumeMonitorGetMountForUuid},
	{"g_volume_monitor_get_mounts", &VolumeMonitorGetMounts},
	{"g_volume_monitor_get_type", &VolumeMonitorGetType},
	{"g_volume_monitor_get_volume_for_uuid", &VolumeMonitorGetVolumeForUuid},
	{"g_volume_monitor_get_volumes", &VolumeMonitorGetVolumes},
	{"g_volume_mount", &VolumeMount},
	{"g_volume_mount_finish", &VolumeMountFinish},
	{"g_volume_should_automount", &VolumeShouldAutomount},
	{"g_win32_input_stream_get_close_handle", &Win32InputStreamGetCloseHandle},
	{"g_win32_input_stream_get_handle", &Win32InputStreamGetHandle},
	{"g_win32_input_stream_get_type", &Win32InputStreamGetType},
	{"g_win32_input_stream_new", &Win32InputStreamNew},
	{"g_win32_input_stream_set_close_handle", &Win32InputStreamSetCloseHandle},
	{"g_win32_output_stream_get_close_handle", &Win32OutputStreamGetCloseHandle},
	{"g_win32_output_stream_get_handle", &Win32OutputStreamGetHandle},
	{"g_win32_output_stream_get_type", &Win32OutputStreamGetType},
	{"g_win32_output_stream_new", &Win32OutputStreamNew},
	{"g_win32_output_stream_set_close_handle", &Win32OutputStreamSetCloseHandle},
	// Undocumented {"g_win32_resolver_get_type", &Win32ResolverGetType},
	{"g_zlib_compressor_format_get_type", &ZlibCompressorFormatGetType},
	{"g_zlib_compressor_get_file_info", &ZlibCompressorGetFileInfo},
	{"g_zlib_compressor_get_type", &ZlibCompressorGetType},
	{"g_zlib_compressor_new", &ZlibCompressorNew},
	{"g_zlib_compressor_set_file_info", &ZlibCompressorSetFileInfo},
	{"g_zlib_decompressor_get_file_info", &ZlibDecompressorGetFileInfo},
	{"g_zlib_decompressor_get_type", &ZlibDecompressorGetType},
	{"g_zlib_decompressor_new", &ZlibDecompressorNew},
}
