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

	tcpConnectionSetGracefulDisconnect func(t *TcpConnection, gracefulDisconnect T.Gboolean)
	tcpConnectionGetGracefulDisconnect func(t *TcpConnection) T.Gboolean
)

func (t *TcpConnection) GetGracefulDisconnect() T.Gboolean {
	return tcpConnectionGetGracefulDisconnect(t)
}
func (t *TcpConnection) SetGracefulDisconnect(gracefulDisconnect T.Gboolean) {
	tcpConnectionSetGracefulDisconnect(t, gracefulDisconnect)
}

type TcpWrapperConnection struct {
	Parent TcpConnection
	_      *struct{}
}

var (
	TcpWrapperConnectionGetType func() O.Type
	TcpWrapperConnectionNew     func(baseIoStream *IOStream, socket *Socket) *SocketConnection

	tcpWrapperConnectionGetBaseIoStream func(t *TcpWrapperConnection) *IOStream
)

func (t *TcpWrapperConnection) GetBaseIoStream() *IOStream {
	return tcpWrapperConnectionGetBaseIoStream(t)
}

type ThemedIcon struct{}

var (
	ThemedIconGetType                 func() O.Type
	ThemedIconNew                     func(iconname string) *Icon
	ThemedIconNewWithDefaultFallbacks func(iconname string) *Icon
	ThemedIconNewFromNames            func(iconnames **T.Char, len int) *Icon

	themedIconAppendName  func(t *ThemedIcon, iconname string)
	themedIconGetNames    func(t *ThemedIcon) []string
	themedIconPrependName func(t *ThemedIcon, iconname string)
)

func (t *ThemedIcon) AppendName(iconname string)  { themedIconAppendName(t, iconname) }
func (t *ThemedIcon) GetNames() []string          { return themedIconGetNames(t) }
func (t *ThemedIcon) PrependName(iconname string) { themedIconPrependName(t, iconname) }

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

	tlsBackendGetCertificateType      func(t *TlsBackend) O.Type
	tlsBackendGetClientConnectionType func(t *TlsBackend) O.Type
	tlsBackendGetServerConnectionType func(t *TlsBackend) O.Type
	tlsBackendSupportsTls             func(t *TlsBackend) T.Gboolean
)

func (t *TlsBackend) GetCertificateType() O.Type      { return tlsBackendGetCertificateType(t) }
func (t *TlsBackend) GetClientConnectionType() O.Type { return tlsBackendGetClientConnectionType(t) }
func (t *TlsBackend) GetServerConnectionType() O.Type { return tlsBackendGetServerConnectionType(t) }
func (t *TlsBackend) SupportsTls() T.Gboolean         { return tlsBackendSupportsTls(t) }

var (
	TlsCertificateGetType func() O.Type

	TlsCertificateNewFromPem      func(data string, length T.Gssize, err **T.GError) *TlsCertificate
	TlsCertificateNewFromFile     func(file string, err **T.GError) *TlsCertificate
	TlsCertificateNewFromFiles    func(certFile string, keyFile string, err **T.GError) *TlsCertificate
	TlsCertificateListNewFromFile func(file string, err **T.GError) *T.GList

	tlsCertificateGetIssuer func(t *TlsCertificate) *TlsCertificate
	tlsCertificateVerify    func(t *TlsCertificate, identity *SocketConnectable, trustedCa *TlsCertificate) TlsCertificateFlags
)

func (t *TlsCertificate) GetIssuer() *TlsCertificate { return tlsCertificateGetIssuer(t) }
func (t *TlsCertificate) Verify(identity *SocketConnectable, trustedCa *TlsCertificate) TlsCertificateFlags {
	return tlsCertificateVerify(t, identity, trustedCa)
}

type TlsConnection struct {
	Parent IOStream
	_      *struct{}
}

var (
	TlsConnectionGetType func() O.Type

	tlsConnectionGetCertificate           func(t *TlsConnection) *TlsCertificate
	tlsConnectionGetPeerCertificate       func(t *TlsConnection) *TlsCertificate
	tlsConnectionGetPeerCertificateErrors func(t *TlsConnection) TlsCertificateFlags
	tlsConnectionGetRehandshakeMode       func(t *TlsConnection) TlsRehandshakeMode
	tlsConnectionGetRequireCloseNotify    func(t *TlsConnection) T.Gboolean
	tlsConnectionGetUseSystemCertdb       func(t *TlsConnection) T.Gboolean
	tlsConnectionHandshake                func(t *TlsConnection, cancellable *Cancellable, err **T.GError) T.Gboolean
	tlsConnectionHandshakeAsync           func(t *TlsConnection, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	tlsConnectionHandshakeFinish          func(t *TlsConnection, result *AsyncResult, err **T.GError) T.Gboolean
	tlsConnectionSetCertificate           func(t *TlsConnection, certificate *TlsCertificate)
	tlsConnectionSetRehandshakeMode       func(t *TlsConnection, mode TlsRehandshakeMode)
	tlsConnectionSetRequireCloseNotify    func(t *TlsConnection, requireCloseNotify T.Gboolean)
	tlsConnectionSetUseSystemCertdb       func(t *TlsConnection, useSystemCertdb T.Gboolean)
)

func (t *TlsConnection) GetCertificate() *TlsCertificate { return tlsConnectionGetCertificate(t) }
func (t *TlsConnection) GetPeerCertificate() *TlsCertificate {
	return tlsConnectionGetPeerCertificate(t)
}
func (t *TlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	return tlsConnectionGetPeerCertificateErrors(t)
}
func (t *TlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	return tlsConnectionGetRehandshakeMode(t)
}
func (t *TlsConnection) GetRequireCloseNotify() T.Gboolean {
	return tlsConnectionGetRequireCloseNotify(t)
}
func (t *TlsConnection) GetUseSystemCertdb() T.Gboolean { return tlsConnectionGetUseSystemCertdb(t) }
func (t *TlsConnection) Handshake(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return tlsConnectionHandshake(t, cancellable, err)
}
func (t *TlsConnection) HandshakeAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	tlsConnectionHandshakeAsync(t, ioPriority, cancellable, callback, userData)
}
func (t *TlsConnection) HandshakeFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return tlsConnectionHandshakeFinish(t, result, err)
}
func (t *TlsConnection) SetCertificate(certificate *TlsCertificate) {
	tlsConnectionSetCertificate(t, certificate)
}
func (t *TlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	tlsConnectionSetRehandshakeMode(t, mode)
}
func (t *TlsConnection) SetRequireCloseNotify(requireCloseNotify T.Gboolean) {
	tlsConnectionSetRequireCloseNotify(t, requireCloseNotify)
}
func (t *TlsConnection) SetUseSystemCertdb(useSystemCertdb T.Gboolean) {
	tlsConnectionSetUseSystemCertdb(t, useSystemCertdb)
}

var (
	TlsErrorQuark func() T.GQuark

	tlsConnectionEmitAcceptCertificate func(t *TlsConnection, peerCert *TlsCertificate, errors TlsCertificateFlags) T.Gboolean
)

func (t *TlsConnection) EmitAcceptCertificate(peerCert *TlsCertificate, errors TlsCertificateFlags) T.Gboolean {
	return tlsConnectionEmitAcceptCertificate(t, peerCert, errors)
}

type TlsClientConnection struct{}

var (
	TlsClientConnectionGetType func() O.Type
	TlsClientConnectionNew     func(baseIoStream *IOStream, serverIdentity *SocketConnectable, err **T.GError) *IOStream

	tlsClientConnectionGetAcceptedCas     func(t *TlsClientConnection) *T.GList
	tlsClientConnectionGetServerIdentity  func(t *TlsClientConnection) *SocketConnectable
	tlsClientConnectionGetUseSsl3         func(t *TlsClientConnection) T.Gboolean
	tlsClientConnectionGetValidationFlags func(t *TlsClientConnection) TlsCertificateFlags
	tlsClientConnectionSetServerIdentity  func(t *TlsClientConnection, identity *SocketConnectable)
	tlsClientConnectionSetUseSsl3         func(t *TlsClientConnection, useSsl3 T.Gboolean)
	tlsClientConnectionSetValidationFlags func(t *TlsClientConnection, flags TlsCertificateFlags)
)

func (t *TlsClientConnection) GetAcceptedCas() *T.GList { return tlsClientConnectionGetAcceptedCas(t) }
func (t *TlsClientConnection) GetServerIdentity() *SocketConnectable {
	return tlsClientConnectionGetServerIdentity(t)
}
func (t *TlsClientConnection) GetUseSsl3() T.Gboolean { return tlsClientConnectionGetUseSsl3(t) }
func (t *TlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	return tlsClientConnectionGetValidationFlags(t)
}
func (t *TlsClientConnection) SetServerIdentity(identity *SocketConnectable) {
	tlsClientConnectionSetServerIdentity(t, identity)
}
func (t *TlsClientConnection) SetUseSsl3(useSsl3 T.Gboolean) {
	tlsClientConnectionSetUseSsl3(t, useSsl3)
}
func (t *TlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	tlsClientConnectionSetValidationFlags(t, flags)
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
