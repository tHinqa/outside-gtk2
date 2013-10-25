// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

var PasswordSaveGetType func() O.Type

type Permission struct {
	Parent O.Object
	_      *struct{}
}

var (
	PermissionGetType func() O.Type

	permissionAcquire       func(p *Permission, cancellable *Cancellable, err **T.GError) T.Gboolean
	permissionAcquireAsync  func(p *Permission, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	permissionAcquireFinish func(p *Permission, result *AsyncResult, err **T.GError) T.Gboolean
	permissionGetAllowed    func(p *Permission) T.Gboolean
	permissionGetCanAcquire func(p *Permission) T.Gboolean
	permissionGetCanRelease func(p *Permission) T.Gboolean
	permissionImplUpdate    func(p *Permission, allowed T.Gboolean, canAcquire T.Gboolean, canRelease T.Gboolean)
	permissionRelease       func(p *Permission, cancellable *Cancellable, err **T.GError) T.Gboolean
	permissionReleaseAsync  func(p *Permission, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	permissionReleaseFinish func(p *Permission, result *AsyncResult, err **T.GError) T.Gboolean
)

func (p *Permission) Acquire(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return permissionAcquire(p, cancellable, err)
}
func (p *Permission) AcquireAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	permissionAcquireAsync(p, cancellable, callback, userData)
}
func (p *Permission) AcquireFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return permissionAcquireFinish(p, result, err)
}
func (p *Permission) GetAllowed() T.Gboolean    { return permissionGetAllowed(p) }
func (p *Permission) GetCanAcquire() T.Gboolean { return permissionGetCanAcquire(p) }
func (p *Permission) GetCanRelease() T.Gboolean { return permissionGetCanRelease(p) }
func (p *Permission) ImplUpdate(allowed T.Gboolean, canAcquire T.Gboolean, canRelease T.Gboolean) {
	permissionImplUpdate(p, allowed, canAcquire, canRelease)
}
func (p *Permission) Release(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return permissionRelease(p, cancellable, err)
}
func (p *Permission) ReleaseAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	permissionReleaseAsync(p, cancellable, callback, userData)
}
func (p *Permission) ReleaseFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return permissionReleaseFinish(p, result, err)
}

type PollableInputStream struct{}

var (
	PollableInputStreamGetType func() O.Type

	pollableInputStreamCanPoll         func(p *PollableInputStream) T.Gboolean
	pollableInputStreamIsReadable      func(p *PollableInputStream) T.Gboolean
	pollableInputStreamCreateSource    func(p *PollableInputStream, cancellable *Cancellable) *T.GSource
	pollableInputStreamReadNonblocking func(p *PollableInputStream, buffer *T.Void, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
)

func (p *PollableInputStream) CanPoll() T.Gboolean    { return pollableInputStreamCanPoll(p) }
func (p *PollableInputStream) IsReadable() T.Gboolean { return pollableInputStreamIsReadable(p) }
func (p *PollableInputStream) CreateSource(cancellable *Cancellable) *T.GSource {
	return pollableInputStreamCreateSource(p, cancellable)
}
func (p *PollableInputStream) ReadNonblocking(buffer *T.Void, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return pollableInputStreamReadNonblocking(p, buffer, size, cancellable, err)
}

var PollableSourceNew func(pollableStream *O.Object) *T.GSource

type PollableOutputStream struct{}

var (
	PollableOutputStreamGetType func() O.Type

	pollableOutputStreamCanPoll          func(p *PollableOutputStream) T.Gboolean
	pollableOutputStreamIsWritable       func(p *PollableOutputStream) T.Gboolean
	pollableOutputStreamCreateSource     func(p *PollableOutputStream, cancellable *Cancellable) *T.GSource
	pollableOutputStreamWriteNonblocking func(p *PollableOutputStream, buffer *T.Void, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize
)

func (p *PollableOutputStream) CanPoll() T.Gboolean    { return pollableOutputStreamCanPoll(p) }
func (p *PollableOutputStream) IsWritable() T.Gboolean { return pollableOutputStreamIsWritable(p) }
func (p *PollableOutputStream) CreateSource(cancellable *Cancellable) *T.GSource {
	return pollableOutputStreamCreateSource(p, cancellable)
}
func (p *PollableOutputStream) WriteNonblocking(buffer *T.Void, size T.Gsize, cancellable *Cancellable, err **T.GError) T.Gssize {
	return pollableOutputStreamWriteNonblocking(p, buffer, size, cancellable, err)
}

type Proxy struct{}

var (
	ProxyGetType func() O.Type

	ProxyGetDefaultForProtocol func(protocol string) *Proxy

	proxyConnect          func(p *Proxy, connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable, err **T.GError) *IOStream
	proxyConnectAsync     func(p *Proxy, connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	proxyConnectFinish    func(p *Proxy, result *AsyncResult, err **T.GError) *IOStream
	proxySupportsHostname func(p *Proxy) T.Gboolean
)

func (p *Proxy) Connect(connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable, err **T.GError) *IOStream {
	return proxyConnect(p, connection, proxyAddress, cancellable, err)
}
func (p *Proxy) ConnectAsync(connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	proxyConnectAsync(p, connection, proxyAddress, cancellable, callback, userData)
}
func (p *Proxy) ConnectFinish(result *AsyncResult, err **T.GError) *IOStream {
	return proxyConnectFinish(p, result, err)
}
func (p *Proxy) SupportsHostname() T.Gboolean { return proxySupportsHostname(p) }

type ProxyAddress struct {
	Parent InetSocketAddress
	_      *struct{}
}

var (
	ProxyAddressGetType func() O.Type
	ProxyAddressNew     func(inetaddr *InetAddress, port uint16, protocol, destHostname string, destPort uint16, username, password string) *SocketAddress

	proxyAddressGetProtocol            func(p *ProxyAddress) string
	proxyAddressGetDestinationHostname func(p *ProxyAddress) string
	proxyAddressGetDestinationPort     func(p *ProxyAddress) uint16
	proxyAddressGetUsername            func(p *ProxyAddress) string
	proxyAddressGetPassword            func(p *ProxyAddress) string
)

func (p *ProxyAddress) GetProtocol() string            { return proxyAddressGetProtocol(p) }
func (p *ProxyAddress) GetDestinationHostname() string { return proxyAddressGetDestinationHostname(p) }
func (p *ProxyAddress) GetDestinationPort() uint16     { return proxyAddressGetDestinationPort(p) }
func (p *ProxyAddress) GetUsername() string            { return proxyAddressGetUsername(p) }
func (p *ProxyAddress) GetPassword() string            { return proxyAddressGetPassword(p) }

var ProxyAddressEnumeratorGetType func() O.Type

type ProxyResolver struct{}

var (
	ProxyResolverGetType func() O.Type

	ProxyResolverGetDefault func() *ProxyResolver

	proxyResolverIsSupported  func(p *ProxyResolver) T.Gboolean
	proxyResolverLookup       func(p *ProxyResolver, uri string, cancellable *Cancellable, err **T.GError) **T.Gchar
	proxyResolverLookupAsync  func(p *ProxyResolver, uri string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	proxyResolverLookupFinish func(p *ProxyResolver, result *AsyncResult, err **T.GError) **T.Gchar
)

func (p *ProxyResolver) IsSupported() T.Gboolean { return proxyResolverIsSupported(p) }
func (p *ProxyResolver) Lookup(uri string, cancellable *Cancellable, err **T.GError) **T.Gchar {
	return proxyResolverLookup(p, uri, cancellable, err)
}
func (p *ProxyResolver) LookupAsync(uri string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	proxyResolverLookupAsync(p, uri, cancellable, callback, userData)
}
func (p *ProxyResolver) LookupFinish(result *AsyncResult, err **T.GError) **T.Gchar {
	return proxyResolverLookupFinish(p, result, err)
}
