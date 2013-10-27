// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

var NativeVolumeMonitorGetType func() O.Type

type NetworkAddress struct {
	Parent O.Object
	_b     *struct{}
}

var (
	NetworkAddressGetType func() O.Type
	NetworkAddressNew     func(hostname string, port uint16) *SocketConnectable

	NetworkAddressParse    func(hostAndPort string, defaultPort uint16, err **T.GError) *SocketConnectable
	NetworkAddressParseUri func(uri string, defaultPort uint16, err **T.GError) *SocketConnectable

	networkAddressGetHostname func(n *NetworkAddress) string
	networkAddressGetPort     func(n *NetworkAddress) uint16
	networkAddressGetScheme   func(n *NetworkAddress) string
)

func (n *NetworkAddress) GetHostname() string { return networkAddressGetHostname(n) }
func (n *NetworkAddress) GetPort() uint16     { return networkAddressGetPort(n) }
func (n *NetworkAddress) GetScheme() string   { return networkAddressGetScheme(n) }

type NetworkService struct {
	Parent O.Object
	_      *struct{}
}

var (
	NetworkServiceGetType func() O.Type
	NetworkServiceNew     func(service string, protocol string, domain string) *SocketConnectable

	networkServiceGetService  func(n *NetworkService) string
	networkServiceGetProtocol func(n *NetworkService) string
	networkServiceGetDomain   func(n *NetworkService) string
	networkServiceGetScheme   func(n *NetworkService) string
	networkServiceSetScheme   func(n *NetworkService, scheme string)
)

func (n *NetworkService) GetService() string      { return networkServiceGetService(n) }
func (n *NetworkService) GetProtocol() string     { return networkServiceGetProtocol(n) }
func (n *NetworkService) GetDomain() string       { return networkServiceGetDomain(n) }
func (n *NetworkService) GetScheme() string       { return networkServiceGetScheme(n) }
func (n *NetworkService) SetScheme(scheme string) { networkServiceSetScheme(n, scheme) }

var NullSettingsBackendNew func() *SettingsBackend
