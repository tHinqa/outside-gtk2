// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
)

var NativeVolumeMonitorGetType func() O.Type

type NetworkAddress struct {
	Parent O.Object
	_      *struct{}
}

var (
	NetworkAddressGetType func() O.Type
	NetworkAddressNew     func(hostname string, port uint16) *SocketConnectable

	NetworkAddressParse    func(hostAndPort string, defaultPort uint16, err **L.Error) *SocketConnectable
	NetworkAddressParseUri func(uri string, defaultPort uint16, err **L.Error) *SocketConnectable

	NetworkAddressGetHostname func(n *NetworkAddress) string
	NetworkAddressGetPort     func(n *NetworkAddress) uint16
	NetworkAddressGetScheme   func(n *NetworkAddress) string
)

func (n *NetworkAddress) GetHostname() string { return NetworkAddressGetHostname(n) }
func (n *NetworkAddress) GetPort() uint16     { return NetworkAddressGetPort(n) }
func (n *NetworkAddress) GetScheme() string   { return NetworkAddressGetScheme(n) }

type NetworkService struct {
	Parent O.Object
	_      *struct{}
}

var (
	NetworkServiceGetType func() O.Type
	NetworkServiceNew     func(service string, protocol string, domain string) *SocketConnectable

	NetworkServiceGetService  func(n *NetworkService) string
	NetworkServiceGetProtocol func(n *NetworkService) string
	NetworkServiceGetDomain   func(n *NetworkService) string
	NetworkServiceGetScheme   func(n *NetworkService) string
	NetworkServiceSetScheme   func(n *NetworkService, scheme string)
)

func (n *NetworkService) GetService() string      { return NetworkServiceGetService(n) }
func (n *NetworkService) GetProtocol() string     { return NetworkServiceGetProtocol(n) }
func (n *NetworkService) GetDomain() string       { return NetworkServiceGetDomain(n) }
func (n *NetworkService) GetScheme() string       { return NetworkServiceGetScheme(n) }
func (n *NetworkService) SetScheme(scheme string) { NetworkServiceSetScheme(n, scheme) }

var NullSettingsBackendNew func() *SettingsBackend
