// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Resolver struct {
	Parent O.Object
	_      *struct{}
}

var (
	ResolverGetType func() O.Type

	ResolverErrorGetType  func() O.Type
	ResolverErrorQuark    func() L.Quark
	ResolverFreeAddresses func(addresses *L.List)
	ResolverFreeTargets   func(targets *L.List)
	ResolverGetDefault    func() *Resolver

	ResolverLookupByAddress       func(r *Resolver, address *InetAddress, cancellable *Cancellable, err **L.Error) string
	ResolverLookupByAddressAsync  func(r *Resolver, address *InetAddress, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	ResolverLookupByAddressFinish func(r *Resolver, result *AsyncResult, err **L.Error) string
	ResolverLookupByName          func(r *Resolver, hostname string, cancellable *Cancellable, err **L.Error) *L.List
	ResolverLookupByNameAsync     func(r *Resolver, hostname string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	ResolverLookupByNameFinish    func(r *Resolver, result *AsyncResult, err **L.Error) *L.List
	ResolverLookupService         func(r *Resolver, service, protocol, domain string, cancellable *Cancellable, err **L.Error) *L.List
	ResolverLookupServiceAsync    func(r *Resolver, service, protocol, domain string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	ResolverLookupServiceFinish   func(r *Resolver, result *AsyncResult, err **L.Error) *L.List
	ResolverSetDefault            func(r *Resolver)
)

func (r *Resolver) LookupByAddress(address *InetAddress, cancellable *Cancellable, err **L.Error) string {
	return ResolverLookupByAddress(r, address, cancellable, err)
}
func (r *Resolver) LookupByAddressAsync(address *InetAddress, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	ResolverLookupByAddressAsync(r, address, cancellable, callback, userData)
}
func (r *Resolver) LookupByAddressFinish(result *AsyncResult, err **L.Error) string {
	return ResolverLookupByAddressFinish(r, result, err)
}
func (r *Resolver) LookupByName(hostname string, cancellable *Cancellable, err **L.Error) *L.List {
	return ResolverLookupByName(r, hostname, cancellable, err)
}
func (r *Resolver) LookupByNameAsync(hostname string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	ResolverLookupByNameAsync(r, hostname, cancellable, callback, userData)
}
func (r *Resolver) LookupByNameFinish(result *AsyncResult, err **L.Error) *L.List {
	return ResolverLookupByNameFinish(r, result, err)
}
func (r *Resolver) LookupService(service, protocol, domain string, cancellable *Cancellable, err **L.Error) *L.List {
	return ResolverLookupService(r, service, protocol, domain, cancellable, err)
}
func (r *Resolver) LookupServiceAsync(service, protocol, domain string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	ResolverLookupServiceAsync(r, service, protocol, domain, cancellable, callback, userData)
}
func (r *Resolver) LookupServiceFinish(result *AsyncResult, err **L.Error) *L.List {
	return ResolverLookupServiceFinish(r, result, err)
}
func (r *Resolver) SetDefault() { ResolverSetDefault(r) }
