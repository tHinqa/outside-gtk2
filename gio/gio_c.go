// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Cancellable struct {
	Parent O.Object
	_      *struct{}
}

var (
	CancellableGetType func() O.Type
	CancellableNew     func() *Cancellable

	CancellableGetCurrent func() *Cancellable

	CancellableCancel              func(c *Cancellable)
	CancellableConnect             func(c *Cancellable, callback T.GCallback, data T.Gpointer, dataDestroyFunc T.GDestroyNotify) T.Gulong
	CancellableDisconnect          func(c *Cancellable, handlerId T.Gulong)
	CancellableGetFd               func(c *Cancellable) int
	CancellableIsCancelled         func(c *Cancellable) bool
	CancellableMakePollfd          func(c *Cancellable, pollfd *T.GPollFD) bool
	CancellablePopCurrent          func(c *Cancellable)
	CancellablePushCurrent         func(c *Cancellable)
	CancellableReleaseFd           func(c *Cancellable)
	CancellableReset               func(c *Cancellable)
	CancellableSetErrorIfCancelled func(c *Cancellable, err **T.GError) bool
	CancellableSourceNew           func(c *Cancellable) *T.GSource
)

func (c *Cancellable) Cancel() { CancellableCancel(c) }
func (c *Cancellable) Connect(callback T.GCallback, data T.Gpointer, dataDestroyFunc T.GDestroyNotify) T.Gulong {
	return CancellableConnect(c, callback, data, dataDestroyFunc)
}
func (c *Cancellable) Disconnect(handlerId T.Gulong) { CancellableDisconnect(c, handlerId) }
func (c *Cancellable) GetFd() int                    { return CancellableGetFd(c) }
func (c *Cancellable) IsCancelled() bool             { return CancellableIsCancelled(c) }
func (c *Cancellable) MakePollfd(pollfd *T.GPollFD) bool {
	return CancellableMakePollfd(c, pollfd)
}
func (c *Cancellable) PopCurrent()  { CancellablePopCurrent(c) }
func (c *Cancellable) PushCurrent() { CancellablePushCurrent(c) }
func (c *Cancellable) ReleaseFd()   { CancellableReleaseFd(c) }
func (c *Cancellable) Reset()       { CancellableReset(c) }
func (c *Cancellable) SetErrorIfCancelled(err **T.GError) bool {
	return CancellableSetErrorIfCancelled(c, err)
}
func (c *Cancellable) SourceNew() *T.GSource { return CancellableSourceNew(c) }

type CharsetConverter struct{}

var (
	CharsetConverterGetType func() O.Type
	CharsetConverterNew     func(toCharset string, fromCharset string, err **T.GError) *CharsetConverter

	CharsetConverterGetNumFallbacks func(c *CharsetConverter) uint
	CharsetConverterGetUseFallback  func(c *CharsetConverter) bool
	CharsetConverterSetUseFallback  func(c *CharsetConverter, useFallback bool)
)

func (c *CharsetConverter) GetNumFallbacks() uint { return CharsetConverterGetNumFallbacks(c) }
func (c *CharsetConverter) GetUseFallback() bool  { return CharsetConverterGetUseFallback(c) }
func (c *CharsetConverter) SetUseFallback(useFallback bool) {
	CharsetConverterSetUseFallback(c, useFallback)
}

var (
	ContentTypeCanBeExecutable func(typ string) bool
	ContentTypeEquals          func(type1 string, type2 string) bool
	ContentTypeFromMimeType    func(mimeType string) string
	ContentTypeGetDescription  func(typ string) string
	ContentTypeGetIcon         func(typ string) *Icon
	ContentTypeGetMimeType     func(typ string) string
	ContentTypeGuess           func(filename string, data *T.Guchar, dataSize T.Gsize, resultUncertain *bool) string
	ContentTypeGuessForTree    func(root *File) []string
	ContentTypeIsA             func(typ string, supertype string) bool
	ContentTypeIsUnknown       func(typ string) bool
	ContentTypesGetRegistered  func() *T.GList
)

type Converter struct{}

var (
	ConverterGetType func() O.Type

	ConverterConvert func(c *Converter, inbuf *T.Void, inbufSize T.Gsize, outbuf *T.Void, outbufSize T.Gsize, flags ConverterFlags, bytesRead, bytesWritten *T.Gsize, err **T.GError) ConverterResult
	ConverterReset   func(c *Converter)
)

func (c *Converter) Convert(inbuf *T.Void, inbufSize T.Gsize, outbuf *T.Void, outbufSize T.Gsize, flags ConverterFlags, bytesRead, bytesWritten *T.Gsize, err **T.GError) ConverterResult {
	return ConverterConvert(c, inbuf, inbufSize, outbuf, outbufSize, flags, bytesRead, bytesWritten, err)
}
func (c *Converter) Reset() { ConverterReset(c) }

type ConverterFlags Enum

const (
	CONVERTER_INPUT_AT_END ConverterFlags = 1 << iota
	CONVERTER_FLUSH
	CONVERTER_NO_FLAGS ConverterFlags = 0
)

var ConverterFlagsGetType func() O.Type

type ConverterInputStream struct {
	Parent FilterInputStream
	_      *struct{}
}

var (
	ConverterInputStreamGetType func() O.Type
	ConverterInputStreamNew     func(baseStream *InputStream, converter *Converter) *InputStream

	ConverterInputStreamGetConverter func(c *ConverterInputStream) *Converter
)

func (c *ConverterInputStream) GetConverter() *Converter { return ConverterInputStreamGetConverter(c) }

type ConverterOutputStream struct {
	Parent FilterOutputStream
	_      *struct{}
}

var (
	ConverterOutputStreamGetType func() O.Type
	ConverterOutputStreamNew     func(baseStream *OutputStream, converter *Converter) *OutputStream

	ConverterOutputStreamGetConverter func(c *ConverterOutputStream) *Converter
)

func (c *ConverterOutputStream) GetConverter() *Converter { return ConverterOutputStreamGetConverter(c) }

type ConverterResult Enum

const (
	CONVERTER_ERROR ConverterResult = iota
	CONVERTER_CONVERTED
	CONVERTER_FINISHED
	CONVERTER_FLUSHED
)

var ConverterResultGetType func() O.Type

type Credentials struct{}

var (
	CredentialsGetType func() O.Type
	CredentialsNew     func() *Credentials

	CredentialsGetNative  func(c *Credentials, nativeType CredentialsType) T.Gpointer
	CredentialsIsSameUser func(c *Credentials, otherCredentials *Credentials, err **T.GError) bool
	CredentialsSetNative  func(c *Credentials, nativeType CredentialsType, native T.Gpointer)
	CredentialsToString   func(c *Credentials) string
)

func (c *Credentials) GetNative(nativeType CredentialsType) T.Gpointer {
	return CredentialsGetNative(c, nativeType)
}
func (c *Credentials) IsSameUser(otherCredentials *Credentials, err **T.GError) bool {
	return CredentialsIsSameUser(c, otherCredentials, err)
}
func (c *Credentials) SetNative(nativeType CredentialsType, native T.Gpointer) {
	CredentialsSetNative(c, nativeType, native)
}
func (c *Credentials) ToString() string { return CredentialsToString(c) }

type CredentialsType Enum

const (
	CREDENTIALS_TYPE_INVALID CredentialsType = iota
	CREDENTIALS_TYPE_LINUX_UCRED
	CREDENTIALS_TYPE_FREEBSD_CMSGCRED
)

var CredentialsTypeGetType func() O.Type
