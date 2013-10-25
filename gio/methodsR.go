// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Resolver struct {
	Parent O.Object
	_      *struct{}
}

var (
	ResolverGetType func() O.Type

	ResolverErrorGetType  func() O.Type
	ResolverErrorQuark    func() T.GQuark
	ResolverFreeAddresses func(addresses *T.GList)
	ResolverFreeTargets   func(targets *T.GList)
	ResolverGetDefault    func() *Resolver

	resolverLookupByAddress       func(r *Resolver, address *InetAddress, cancellable *Cancellable, err **T.GError) string
	resolverLookupByAddressAsync  func(r *Resolver, address *InetAddress, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	resolverLookupByAddressFinish func(r *Resolver, result *AsyncResult, err **T.GError) string
	resolverLookupByName          func(r *Resolver, hostname string, cancellable *Cancellable, err **T.GError) *T.GList
	resolverLookupByNameAsync     func(r *Resolver, hostname string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	resolverLookupByNameFinish    func(r *Resolver, result *AsyncResult, err **T.GError) *T.GList
	resolverLookupService         func(r *Resolver, service, protocol, domain string, cancellable *Cancellable, err **T.GError) *T.GList
	resolverLookupServiceAsync    func(r *Resolver, service, protocol, domain string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	resolverLookupServiceFinish   func(r *Resolver, result *AsyncResult, err **T.GError) *T.GList
	resolverSetDefault            func(r *Resolver)
)

func (r *Resolver) LookupByAddress(address *InetAddress, cancellable *Cancellable, err **T.GError) string {
	return resolverLookupByAddress(r, address, cancellable, err)
}
func (r *Resolver) LookupByAddressAsync(address *InetAddress, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	resolverLookupByAddressAsync(r, address, cancellable, callback, userData)
}
func (r *Resolver) LookupByAddressFinish(result *AsyncResult, err **T.GError) string {
	return resolverLookupByAddressFinish(r, result, err)
}
func (r *Resolver) LookupByName(hostname string, cancellable *Cancellable, err **T.GError) *T.GList {
	return resolverLookupByName(r, hostname, cancellable, err)
}
func (r *Resolver) LookupByNameAsync(hostname string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	resolverLookupByNameAsync(r, hostname, cancellable, callback, userData)
}
func (r *Resolver) LookupByNameFinish(result *AsyncResult, err **T.GError) *T.GList {
	return resolverLookupByNameFinish(r, result, err)
}
func (r *Resolver) LookupService(service, protocol, domain string, cancellable *Cancellable, err **T.GError) *T.GList {
	return resolverLookupService(r, service, protocol, domain, cancellable, err)
}
func (r *Resolver) LookupServiceAsync(service, protocol, domain string, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	resolverLookupServiceAsync(r, service, protocol, domain, cancellable, callback, userData)
}
func (r *Resolver) LookupServiceFinish(result *AsyncResult, err **T.GError) *T.GList {
	return resolverLookupServiceFinish(r, result, err)
}
func (r *Resolver) SetDefault() { resolverSetDefault(r) }
