// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
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

	PermissionAcquire       func(p *Permission, cancellable *Cancellable, err **L.Error) bool
	PermissionAcquireAsync  func(p *Permission, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	PermissionAcquireFinish func(p *Permission, result *AsyncResult, err **L.Error) bool
	PermissionGetAllowed    func(p *Permission) bool
	PermissionGetCanAcquire func(p *Permission) bool
	PermissionGetCanRelease func(p *Permission) bool
	PermissionImplUpdate    func(p *Permission, allowed bool, canAcquire bool, canRelease bool)
	PermissionRelease       func(p *Permission, cancellable *Cancellable, err **L.Error) bool
	PermissionReleaseAsync  func(p *Permission, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	PermissionReleaseFinish func(p *Permission, result *AsyncResult, err **L.Error) bool
)

func (p *Permission) Acquire(cancellable *Cancellable, err **L.Error) bool {
	return PermissionAcquire(p, cancellable, err)
}
func (p *Permission) AcquireAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	PermissionAcquireAsync(p, cancellable, callback, userData)
}
func (p *Permission) AcquireFinish(result *AsyncResult, err **L.Error) bool {
	return PermissionAcquireFinish(p, result, err)
}
func (p *Permission) GetAllowed() bool    { return PermissionGetAllowed(p) }
func (p *Permission) GetCanAcquire() bool { return PermissionGetCanAcquire(p) }
func (p *Permission) GetCanRelease() bool { return PermissionGetCanRelease(p) }
func (p *Permission) ImplUpdate(allowed bool, canAcquire bool, canRelease bool) {
	PermissionImplUpdate(p, allowed, canAcquire, canRelease)
}
func (p *Permission) Release(cancellable *Cancellable, err **L.Error) bool {
	return PermissionRelease(p, cancellable, err)
}
func (p *Permission) ReleaseAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	PermissionReleaseAsync(p, cancellable, callback, userData)
}
func (p *Permission) ReleaseFinish(result *AsyncResult, err **L.Error) bool {
	return PermissionReleaseFinish(p, result, err)
}

type PollableInputStream struct{}

var (
	PollableInputStreamGetType func() O.Type

	PollableInputStreamCanPoll         func(p *PollableInputStream) bool
	PollableInputStreamIsReadable      func(p *PollableInputStream) bool
	PollableInputStreamCreateSource    func(p *PollableInputStream, cancellable *Cancellable) *T.GSource
	PollableInputStreamReadNonblocking func(p *PollableInputStream, buffer *T.Void, size T.Gsize, cancellable *Cancellable, err **L.Error) T.Gssize
)

func (p *PollableInputStream) CanPoll() bool    { return PollableInputStreamCanPoll(p) }
func (p *PollableInputStream) IsReadable() bool { return PollableInputStreamIsReadable(p) }
func (p *PollableInputStream) CreateSource(cancellable *Cancellable) *T.GSource {
	return PollableInputStreamCreateSource(p, cancellable)
}
func (p *PollableInputStream) ReadNonblocking(buffer *T.Void, size T.Gsize, cancellable *Cancellable, err **L.Error) T.Gssize {
	return PollableInputStreamReadNonblocking(p, buffer, size, cancellable, err)
}

var PollableSourceNew func(pollableStream *O.Object) *T.GSource

type PollableOutputStream struct{}

var (
	PollableOutputStreamGetType func() O.Type

	PollableOutputStreamCanPoll          func(p *PollableOutputStream) bool
	PollableOutputStreamIsWritable       func(p *PollableOutputStream) bool
	PollableOutputStreamCreateSource     func(p *PollableOutputStream, cancellable *Cancellable) *T.GSource
	PollableOutputStreamWriteNonblocking func(p *PollableOutputStream, buffer *T.Void, size T.Gsize, cancellable *Cancellable, err **L.Error) T.Gssize
)

func (p *PollableOutputStream) CanPoll() bool    { return PollableOutputStreamCanPoll(p) }
func (p *PollableOutputStream) IsWritable() bool { return PollableOutputStreamIsWritable(p) }
func (p *PollableOutputStream) CreateSource(cancellable *Cancellable) *T.GSource {
	return PollableOutputStreamCreateSource(p, cancellable)
}
func (p *PollableOutputStream) WriteNonblocking(buffer *T.Void, size T.Gsize, cancellable *Cancellable, err **L.Error) T.Gssize {
	return PollableOutputStreamWriteNonblocking(p, buffer, size, cancellable, err)
}

type Proxy struct{}

var (
	ProxyGetType func() O.Type

	ProxyGetDefaultForProtocol func(protocol string) *Proxy

	ProxyConnect          func(p *Proxy, connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable, err **L.Error) *IOStream
	ProxyConnectAsync     func(p *Proxy, connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	ProxyConnectFinish    func(p *Proxy, result *AsyncResult, err **L.Error) *IOStream
	ProxySupportsHostname func(p *Proxy) bool
)

func (p *Proxy) Connect(connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable, err **L.Error) *IOStream {
	return ProxyConnect(p, connection, proxyAddress, cancellable, err)
}
func (p *Proxy) ConnectAsync(connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	ProxyConnectAsync(p, connection, proxyAddress, cancellable, callback, userData)
}
func (p *Proxy) ConnectFinish(result *AsyncResult, err **L.Error) *IOStream {
	return ProxyConnectFinish(p, result, err)
}
func (p *Proxy) SupportsHostname() bool { return ProxySupportsHostname(p) }

type ProxyAddress struct {
	Parent InetSocketAddress
	_      *struct{}
}

var (
	ProxyAddressGetType func() O.Type
	ProxyAddressNew     func(inetaddr *InetAddress, port uint16, protocol, destHostname string, destPort uint16, username, password string) *SocketAddress

	ProxyAddressGetProtocol            func(p *ProxyAddress) string
	ProxyAddressGetDestinationHostname func(p *ProxyAddress) string
	ProxyAddressGetDestinationPort     func(p *ProxyAddress) uint16
	ProxyAddressGetUsername            func(p *ProxyAddress) string
	ProxyAddressGetPassword            func(p *ProxyAddress) string
)

func (p *ProxyAddress) GetProtocol() string            { return ProxyAddressGetProtocol(p) }
func (p *ProxyAddress) GetDestinationHostname() string { return ProxyAddressGetDestinationHostname(p) }
func (p *ProxyAddress) GetDestinationPort() uint16     { return ProxyAddressGetDestinationPort(p) }
func (p *ProxyAddress) GetUsername() string            { return ProxyAddressGetUsername(p) }
func (p *ProxyAddress) GetPassword() string            { return ProxyAddressGetPassword(p) }

var ProxyAddressEnumeratorGetType func() O.Type

type ProxyResolver struct{}

var (
	ProxyResolverGetType func() O.Type

	ProxyResolverGetDefault func() *ProxyResolver

	ProxyResolverIsSupported  func(p *ProxyResolver) bool
	ProxyResolverLookup       func(p *ProxyResolver, uri string, cancellable *Cancellable, err **L.Error) []string
	ProxyResolverLookupAsync  func(p *ProxyResolver, uri string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	ProxyResolverLookupFinish func(p *ProxyResolver, result *AsyncResult, err **L.Error) []string
)

func (p *ProxyResolver) IsSupported() bool { return ProxyResolverIsSupported(p) }
func (p *ProxyResolver) Lookup(uri string, cancellable *Cancellable, err **L.Error) []string {
	return ProxyResolverLookup(p, uri, cancellable, err)
}
func (p *ProxyResolver) LookupAsync(uri string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	ProxyResolverLookupAsync(p, uri, cancellable, callback, userData)
}
func (p *ProxyResolver) LookupFinish(result *AsyncResult, err **L.Error) []string {
	return ProxyResolverLookupFinish(p, result, err)
}
