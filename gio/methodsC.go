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

	cancellableCancel              func(c *Cancellable)
	cancellableConnect             func(c *Cancellable, callback T.GCallback, data T.Gpointer, dataDestroyFunc T.GDestroyNotify) T.Gulong
	cancellableDisconnect          func(c *Cancellable, handlerId T.Gulong)
	cancellableGetFd               func(c *Cancellable) int
	cancellableIsCancelled         func(c *Cancellable) T.Gboolean
	cancellableMakePollfd          func(c *Cancellable, pollfd *T.GPollFD) T.Gboolean
	cancellablePopCurrent          func(c *Cancellable)
	cancellablePushCurrent         func(c *Cancellable)
	cancellableReleaseFd           func(c *Cancellable)
	cancellableReset               func(c *Cancellable)
	cancellableSetErrorIfCancelled func(c *Cancellable, err **T.GError) T.Gboolean
	cancellableSourceNew           func(c *Cancellable) *T.GSource
)

func (c *Cancellable) Cancel() { cancellableCancel(c) }
func (c *Cancellable) Connect(callback T.GCallback, data T.Gpointer, dataDestroyFunc T.GDestroyNotify) T.Gulong {
	return cancellableConnect(c, callback, data, dataDestroyFunc)
}
func (c *Cancellable) Disconnect(handlerId T.Gulong) { cancellableDisconnect(c, handlerId) }
func (c *Cancellable) GetFd() int                    { return cancellableGetFd(c) }
func (c *Cancellable) IsCancelled() T.Gboolean       { return cancellableIsCancelled(c) }
func (c *Cancellable) MakePollfd(pollfd *T.GPollFD) T.Gboolean {
	return cancellableMakePollfd(c, pollfd)
}
func (c *Cancellable) PopCurrent()  { cancellablePopCurrent(c) }
func (c *Cancellable) PushCurrent() { cancellablePushCurrent(c) }
func (c *Cancellable) ReleaseFd()   { cancellableReleaseFd(c) }
func (c *Cancellable) Reset()       { cancellableReset(c) }
func (c *Cancellable) SetErrorIfCancelled(err **T.GError) T.Gboolean {
	return cancellableSetErrorIfCancelled(c, err)
}
func (c *Cancellable) SourceNew() *T.GSource { return cancellableSourceNew(c) }

type CharsetConverter struct{}

var (
	charsetConverterGetType func() O.Type
	charsetConverterNew     func(toCharset string, fromCharset string, err **T.GError) *CharsetConverter

	charsetConverterGetNumFallbacks func(c *CharsetConverter) uint
	charsetConverterGetUseFallback  func(c *CharsetConverter) T.Gboolean
	charsetConverterSetUseFallback  func(c *CharsetConverter, useFallback T.Gboolean)
)

func (c *CharsetConverter) GetNumFallbacks() uint      { return charsetConverterGetNumFallbacks(c) }
func (c *CharsetConverter) GetUseFallback() T.Gboolean { return charsetConverterGetUseFallback(c) }
func (c *CharsetConverter) SetUseFallback(useFallback T.Gboolean) {
	charsetConverterSetUseFallback(c, useFallback)
}

var (
	ContentTypeCanBeExecutable func(typ string) T.Gboolean
	ContentTypeEquals          func(type1 string, type2 string) T.Gboolean
	ContentTypeFromMimeType    func(mimeType string) string
	ContentTypeGetDescription  func(typ string) string
	ContentTypeGetIcon         func(typ string) *Icon
	ContentTypeGetMimeType     func(typ string) string
	ContentTypeGuess           func(filename string, data *T.Guchar, dataSize T.Gsize, resultUncertain *T.Gboolean) string
	ContentTypeGuessForTree    func(root *File) **T.Gchar
	ContentTypeIsA             func(typ string, supertype string) T.Gboolean
	ContentTypeIsUnknown       func(typ string) T.Gboolean
	ContentTypesGetRegistered  func() *T.GList
)

type Converter struct{}

var (
	ConverterGetType func() O.Type

	converterConvert func(c *Converter, inbuf *T.Void, inbufSize T.Gsize, outbuf *T.Void, outbufSize T.Gsize, flags ConverterFlags, bytesRead, bytesWritten *T.Gsize, err **T.GError) ConverterResult
	converterReset   func(c *Converter)
)

func (c *Converter) Convert(inbuf *T.Void, inbufSize T.Gsize, outbuf *T.Void, outbufSize T.Gsize, flags ConverterFlags, bytesRead, bytesWritten *T.Gsize, err **T.GError) ConverterResult {
	return converterConvert(c, inbuf, inbufSize, outbuf, outbufSize, flags, bytesRead, bytesWritten, err)
}
func (c *Converter) Reset() { converterReset(c) }

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

	converterInputStreamGetConverter func(c *ConverterInputStream) *Converter
)

func (c *ConverterInputStream) GetConverter() *Converter { return converterInputStreamGetConverter(c) }

type ConverterOutputStream struct {
	Parent FilterOutputStream
	_      *struct{}
}

var (
	ConverterOutputStreamGetType func() O.Type
	ConverterOutputStreamNew     func(baseStream *OutputStream, converter *Converter) *OutputStream

	converterOutputStreamGetConverter func(c *ConverterOutputStream) *Converter
)

func (c *ConverterOutputStream) GetConverter() *Converter { return converterOutputStreamGetConverter(c) }

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

	credentialsGetNative  func(c *Credentials, nativeType CredentialsType) T.Gpointer
	credentialsIsSameUser func(c *Credentials, otherCredentials *Credentials, err **T.GError) T.Gboolean
	credentialsSetNative  func(c *Credentials, nativeType CredentialsType, native T.Gpointer)
	credentialsToString   func(c *Credentials) string
)

func (c *Credentials) GetNative(nativeType CredentialsType) T.Gpointer {
	return credentialsGetNative(c, nativeType)
}
func (c *Credentials) IsSameUser(otherCredentials *Credentials, err **T.GError) T.Gboolean {
	return credentialsIsSameUser(c, otherCredentials, err)
}
func (c *Credentials) SetNative(nativeType CredentialsType, native T.Gpointer) {
	credentialsSetNative(c, nativeType, native)
}
func (c *Credentials) ToString() string { return credentialsToString(c) }

type CredentialsType Enum

const (
	CREDENTIALS_TYPE_INVALID CredentialsType = iota
	CREDENTIALS_TYPE_LINUX_UCRED
	CREDENTIALS_TYPE_FREEBSD_CMSGCRED
)

var CredentialsTypeGetType func() O.Type
