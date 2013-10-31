// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

var (
	TlsErrorGetType func() O.Type

	TlsAuthenticationModeGetType func() O.Type
)

type TcpConnection struct {
	Parent SocketConnection
	_      *struct{}
}

var (
	TcpConnectionGetType func() O.Type

	TcpConnectionSetGracefulDisconnect func(t *TcpConnection, gracefulDisconnect bool)
	TcpConnectionGetGracefulDisconnect func(t *TcpConnection) bool
)

func (t *TcpConnection) GetGracefulDisconnect() bool {
	return TcpConnectionGetGracefulDisconnect(t)
}
func (t *TcpConnection) SetGracefulDisconnect(gracefulDisconnect bool) {
	TcpConnectionSetGracefulDisconnect(t, gracefulDisconnect)
}

type TcpWrapperConnection struct {
	Parent TcpConnection
	_      *struct{}
}

var (
	TcpWrapperConnectionGetType func() O.Type
	TcpWrapperConnectionNew     func(baseIoStream *IOStream, socket *Socket) *SocketConnection

	TcpWrapperConnectionGetBaseIoStream func(t *TcpWrapperConnection) *IOStream
)

func (t *TcpWrapperConnection) GetBaseIoStream() *IOStream {
	return TcpWrapperConnectionGetBaseIoStream(t)
}

type ThemedIcon struct{}

var (
	ThemedIconGetType                 func() O.Type
	ThemedIconNew                     func(iconname string) *Icon
	ThemedIconNewWithDefaultFallbacks func(iconname string) *Icon
	ThemedIconNewFromNames            func(iconnames **T.Char, len int) *Icon

	ThemedIconAppendName  func(t *ThemedIcon, iconname string)
	ThemedIconGetNames    func(t *ThemedIcon) []string
	ThemedIconPrependName func(t *ThemedIcon, iconname string)
)

func (t *ThemedIcon) AppendName(iconname string)  { ThemedIconAppendName(t, iconname) }
func (t *ThemedIcon) GetNames() []string          { return ThemedIconGetNames(t) }
func (t *ThemedIcon) PrependName(iconname string) { ThemedIconPrependName(t, iconname) }

var (
	ThreadedSocketServiceGetType func() O.Type
	ThreadedSocketServiceNew     func(maxThreads int) *SocketService
)

type TlsBackend struct{}

type TlsCertificate struct {
	Parent O.Object
	_      *struct{}
}

var (
	TlsBackendGetType func() O.Type

	TlsBackendGetDefault func() *TlsBackend

	TlsBackendGetCertificateType      func(t *TlsBackend) O.Type
	TlsBackendGetClientConnectionType func(t *TlsBackend) O.Type
	TlsBackendGetServerConnectionType func(t *TlsBackend) O.Type
	TlsBackendSupportsTls             func(t *TlsBackend) bool
)

func (t *TlsBackend) GetCertificateType() O.Type      { return TlsBackendGetCertificateType(t) }
func (t *TlsBackend) GetClientConnectionType() O.Type { return TlsBackendGetClientConnectionType(t) }
func (t *TlsBackend) GetServerConnectionType() O.Type { return TlsBackendGetServerConnectionType(t) }
func (t *TlsBackend) SupportsTls() bool               { return TlsBackendSupportsTls(t) }

var (
	TlsCertificateGetType func() O.Type

	TlsCertificateNewFromPem      func(data string, length T.Gssize, err **T.GError) *TlsCertificate
	TlsCertificateNewFromFile     func(file string, err **T.GError) *TlsCertificate
	TlsCertificateNewFromFiles    func(certFile string, keyFile string, err **T.GError) *TlsCertificate
	TlsCertificateListNewFromFile func(file string, err **T.GError) *T.GList

	TlsCertificateGetIssuer func(t *TlsCertificate) *TlsCertificate
	TlsCertificateVerify    func(t *TlsCertificate, identity *SocketConnectable, trustedCa *TlsCertificate) TlsCertificateFlags
)

func (t *TlsCertificate) GetIssuer() *TlsCertificate { return TlsCertificateGetIssuer(t) }
func (t *TlsCertificate) Verify(identity *SocketConnectable, trustedCa *TlsCertificate) TlsCertificateFlags {
	return TlsCertificateVerify(t, identity, trustedCa)
}

type TlsConnection struct {
	Parent IOStream
	_      *struct{}
}

var (
	TlsConnectionGetType func() O.Type

	TlsConnectionGetCertificate           func(t *TlsConnection) *TlsCertificate
	TlsConnectionGetPeerCertificate       func(t *TlsConnection) *TlsCertificate
	TlsConnectionGetPeerCertificateErrors func(t *TlsConnection) TlsCertificateFlags
	TlsConnectionGetRehandshakeMode       func(t *TlsConnection) TlsRehandshakeMode
	TlsConnectionGetRequireCloseNotify    func(t *TlsConnection) bool
	TlsConnectionGetUseSystemCertdb       func(t *TlsConnection) bool
	TlsConnectionHandshake                func(t *TlsConnection, cancellable *Cancellable, err **T.GError) bool
	TlsConnectionHandshakeAsync           func(t *TlsConnection, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	TlsConnectionHandshakeFinish          func(t *TlsConnection, result *AsyncResult, err **T.GError) bool
	TlsConnectionSetCertificate           func(t *TlsConnection, certificate *TlsCertificate)
	TlsConnectionSetRehandshakeMode       func(t *TlsConnection, mode TlsRehandshakeMode)
	TlsConnectionSetRequireCloseNotify    func(t *TlsConnection, requireCloseNotify bool)
	TlsConnectionSetUseSystemCertdb       func(t *TlsConnection, useSystemCertdb bool)
)

func (t *TlsConnection) GetCertificate() *TlsCertificate { return TlsConnectionGetCertificate(t) }
func (t *TlsConnection) GetPeerCertificate() *TlsCertificate {
	return TlsConnectionGetPeerCertificate(t)
}
func (t *TlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	return TlsConnectionGetPeerCertificateErrors(t)
}
func (t *TlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	return TlsConnectionGetRehandshakeMode(t)
}
func (t *TlsConnection) GetRequireCloseNotify() bool {
	return TlsConnectionGetRequireCloseNotify(t)
}
func (t *TlsConnection) GetUseSystemCertdb() bool { return TlsConnectionGetUseSystemCertdb(t) }
func (t *TlsConnection) Handshake(cancellable *Cancellable, err **T.GError) bool {
	return TlsConnectionHandshake(t, cancellable, err)
}
func (t *TlsConnection) HandshakeAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	TlsConnectionHandshakeAsync(t, ioPriority, cancellable, callback, userData)
}
func (t *TlsConnection) HandshakeFinish(result *AsyncResult, err **T.GError) bool {
	return TlsConnectionHandshakeFinish(t, result, err)
}
func (t *TlsConnection) SetCertificate(certificate *TlsCertificate) {
	TlsConnectionSetCertificate(t, certificate)
}
func (t *TlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	TlsConnectionSetRehandshakeMode(t, mode)
}
func (t *TlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	TlsConnectionSetRequireCloseNotify(t, requireCloseNotify)
}
func (t *TlsConnection) SetUseSystemCertdb(useSystemCertdb bool) {
	TlsConnectionSetUseSystemCertdb(t, useSystemCertdb)
}

var (
	TlsErrorQuark func() T.GQuark

	TlsConnectionEmitAcceptCertificate func(t *TlsConnection, peerCert *TlsCertificate, errors TlsCertificateFlags) bool
)

func (t *TlsConnection) EmitAcceptCertificate(peerCert *TlsCertificate, errors TlsCertificateFlags) bool {
	return TlsConnectionEmitAcceptCertificate(t, peerCert, errors)
}

type TlsClientConnection struct{}

var (
	TlsClientConnectionGetType func() O.Type
	TlsClientConnectionNew     func(baseIoStream *IOStream, serverIdentity *SocketConnectable, err **T.GError) *IOStream

	TlsClientConnectionGetAcceptedCas     func(t *TlsClientConnection) *T.GList
	TlsClientConnectionGetServerIdentity  func(t *TlsClientConnection) *SocketConnectable
	TlsClientConnectionGetUseSsl3         func(t *TlsClientConnection) bool
	TlsClientConnectionGetValidationFlags func(t *TlsClientConnection) TlsCertificateFlags
	TlsClientConnectionSetServerIdentity  func(t *TlsClientConnection, identity *SocketConnectable)
	TlsClientConnectionSetUseSsl3         func(t *TlsClientConnection, useSsl3 bool)
	TlsClientConnectionSetValidationFlags func(t *TlsClientConnection, flags TlsCertificateFlags)
)

func (t *TlsClientConnection) GetAcceptedCas() *T.GList { return TlsClientConnectionGetAcceptedCas(t) }
func (t *TlsClientConnection) GetServerIdentity() *SocketConnectable {
	return TlsClientConnectionGetServerIdentity(t)
}
func (t *TlsClientConnection) GetUseSsl3() bool { return TlsClientConnectionGetUseSsl3(t) }
func (t *TlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	return TlsClientConnectionGetValidationFlags(t)
}
func (t *TlsClientConnection) SetServerIdentity(identity *SocketConnectable) {
	TlsClientConnectionSetServerIdentity(t, identity)
}
func (t *TlsClientConnection) SetUseSsl3(useSsl3 bool) {
	TlsClientConnectionSetUseSsl3(t, useSsl3)
}
func (t *TlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	TlsClientConnectionSetValidationFlags(t, flags)
}

var (
	TlsServerConnectionGetType func() O.Type
	TlsServerConnectionNew     func(t *IOStream, certificate *TlsCertificate, err **T.GError) *IOStream
)

type TlsCertificateFlags Enum

const (
	TLS_CERTIFICATE_UNKNOWN_CA TlsCertificateFlags = 1 << iota
	TLS_CERTIFICATE_BAD_IDENTITY
	TLS_CERTIFICATE_NOT_ACTIVATED
	TLS_CERTIFICATE_EXPIRED
	TLS_CERTIFICATE_REVOKED
	TLS_CERTIFICATE_INSECURE
	TLS_CERTIFICATE_GENERIC_ERROR
	TLS_CERTIFICATE_VALIDATE_ALL TlsCertificateFlags = 0x007f
)

var TlsCertificateFlagsGetType func() O.Type

type TlsRehandshakeMode Enum

const (
	TLS_REHANDSHAKE_NEVER TlsRehandshakeMode = iota
	TLS_REHANDSHAKE_SAFELY
	TLS_REHANDSHAKE_UNSAFELY
)

var TlsRehandshakeModeGetType func() O.Type
