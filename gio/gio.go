// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gio provides API definitions for accessing
//libgio-2.0-0.dll.
package gio

import (
	"github.com/tHinqa/outside"
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	AppInfoGetType func() O.Type

	AppInfoCreateFromCommandline func(
		commandline string,
		applicationName string,
		flags T.GAppInfoCreateFlags,
		err **T.GError) *T.GAppInfo

	AppInfoDup func(appinfo *T.GAppInfo) *T.GAppInfo

	AppInfoEqual func(
		appinfo1 *T.GAppInfo, appinfo2 *T.GAppInfo) T.Gboolean

	AppInfoGetId func(appinfo *T.GAppInfo) string

	AppInfoGetName func(appinfo *T.GAppInfo) string

	AppInfoGetDisplayName func(appinfo *T.GAppInfo) string

	AppInfoGetDescription func(appinfo *T.GAppInfo) string

	AppInfoGetExecutable func(appinfo *T.GAppInfo) string

	AppInfoGetCommandline func(appinfo *T.GAppInfo) string

	AppInfoGetIcon func(appinfo *T.GAppInfo) *T.GIcon

	AppInfoLaunch func(
		appinfo *T.GAppInfo,
		files *T.GList,
		launchContext *T.GAppLaunchContext,
		err **T.GError) T.Gboolean

	AppInfoSupportsUris func(appinfo *T.GAppInfo) T.Gboolean

	AppInfoSupportsFiles func(appinfo *T.GAppInfo) T.Gboolean

	AppInfoLaunchUris func(
		appinfo *T.GAppInfo,
		uris *T.GList,
		launchContext *T.GAppLaunchContext,
		err **T.GError) T.Gboolean

	AppInfoShouldShow func(appinfo *T.GAppInfo) T.Gboolean

	AppInfoSetAsDefaultForType func(
		appinfo *T.GAppInfo,
		contentType string,
		err **T.GError) T.Gboolean

	AppInfoSetAsDefaultForExtension func(
		appinfo *T.GAppInfo,
		extension string,
		err **T.GError) T.Gboolean

	AppInfoAddSupportsType func(
		appinfo *T.GAppInfo,
		contentType string,
		err **T.GError) T.Gboolean

	AppInfoCanRemoveSupportsType func(
		appinfo *T.GAppInfo) T.Gboolean

	AppInfoRemoveSupportsType func(
		appinfo *T.GAppInfo,
		contentType string,
		err **T.GError) T.Gboolean

	AppInfoCanDelete func(appinfo *T.GAppInfo) T.Gboolean

	AppInfoDelete func(appinfo *T.GAppInfo) T.Gboolean

	AppInfoSetAsLastUsedForType func(
		appinfo *T.GAppInfo,
		contentType string,
		err **T.GError) T.Gboolean

	AppInfoGetAll func() *T.GList

	AppInfoGetAllForType func(contentType string) *T.GList

	AppInfoGetRecommendedForType func(
		contentType string) *T.GList

	AppInfoGetFallbackForType func(
		contentType string) *T.GList

	AppInfoResetTypeAssociations func(contentType string)

	AppInfoGetDefaultForType func(
		contentType string,
		mustSupportUris T.Gboolean) *T.GAppInfo

	AppInfoGetDefaultForUriScheme func(
		uriScheme string) *T.GAppInfo

	AppInfoLaunchDefaultForUri func(
		uri string,
		launchContext *T.GAppLaunchContext,
		err **T.GError) T.Gboolean

	AppLaunchContextGetType func() O.Type

	AppLaunchContextNew func() *T.GAppLaunchContext

	AppLaunchContextGetDisplay func(
		context *T.GAppLaunchContext,
		info *T.GAppInfo,
		files *T.GList) string

	AppLaunchContextGetStartupNotifyId func(
		context *T.GAppLaunchContext,
		info *T.GAppInfo,
		files *T.GList) string

	AppLaunchContextLaunchFailed func(
		context *T.GAppLaunchContext,
		startupNotifyId string)

	ActionGetType func() O.Type

	ActionGetName func(action *T.GAction) string

	ActionGetParameterType func(action *T.GAction) *T.GVariantType

	ActionGetStateType func(action *T.GAction) *T.GVariantType

	ActionGetStateHint func(action *T.GAction) *T.GVariant

	ActionGetEnabled func(action *T.GAction) T.Gboolean

	ActionGetState func(action *T.GAction) *T.GVariant

	ActionSetState func(action *T.GAction, value *T.GVariant)

	ActionActivate func(action *T.GAction, parameter *T.GVariant)

	SimpleActionGetType func() O.Type

	SimpleActionNew func(
		name string, parameterType *T.GVariantType) *T.GSimpleAction

	SimpleActionNewStateful func(
		name string,
		parameterType *T.GVariantType,
		state *T.GVariant) *T.GSimpleAction

	SimpleActionSetEnabled func(
		simple *T.GSimpleAction,
		enabled T.Gboolean)

	ActionGroupGetType func() O.Type

	ActionGroupHasAction func(
		actionGroup *T.GActionGroup,
		actionName string) T.Gboolean

	ActionGroupListActions func(
		actionGroup *T.GActionGroup) **T.Gchar

	ActionGroupGetActionParameterType func(
		actionGroup *T.GActionGroup,
		actionName string) *T.GVariantType

	ActionGroupGetActionStateType func(
		actionGroup *T.GActionGroup,
		actionName string) *T.GVariantType

	ActionGroupGetActionStateHint func(
		actionGroup *T.GActionGroup,
		actionName string) *T.GVariant

	ActionGroupGetActionEnabled func(
		actionGroup *T.GActionGroup,
		actionName string) T.Gboolean

	ActionGroupGetActionState func(
		actionGroup *T.GActionGroup,
		actionName string) *T.GVariant

	ActionGroupChangeActionState func(
		actionGroup *T.GActionGroup,
		actionName string,
		value *T.GVariant)

	ActionGroupActivateAction func(
		actionGroup *T.GActionGroup,
		actionName string,
		parameter *T.GVariant)

	ActionGroupActionAdded func(
		actionGroup *T.GActionGroup,
		actionName string)

	ActionGroupActionRemoved func(
		actionGroup *T.GActionGroup,
		actionName string)

	ActionGroupActionEnabledChanged func(
		actionGroup *T.GActionGroup,
		actionName string,
		enabled T.Gboolean)

	ActionGroupActionStateChanged func(
		actionGroup *T.GActionGroup,
		actionName string,
		state *T.GVariant)

	SimpleActionGroupGetType func() O.Type

	SimpleActionGroupNew func() *T.GSimpleActionGroup

	SimpleActionGroupLookup func(
		simple *T.GSimpleActionGroup,
		actionName string) *T.GAction

	SimpleActionGroupInsert func(
		simple *T.GSimpleActionGroup,
		action *T.GAction)

	SimpleActionGroupRemove func(
		simple *T.GSimpleActionGroup,
		actionName string)

	ApplicationGetType func() O.Type

	ApplicationIdIsValid func(
		applicationId string) T.Gboolean

	ApplicationNew func(
		applicationId string,
		flags T.GApplicationFlags) *T.GApplication

	ApplicationGetApplicationId func(
		application *T.GApplication) string

	ApplicationSetApplicationId func(
		application *T.GApplication,
		applicationId string)

	ApplicationGetInactivityTimeout func(
		application *T.GApplication) uint

	ApplicationSetInactivityTimeout func(
		application *T.GApplication,
		inactivityTimeout uint)

	ApplicationGetFlags func(
		application *T.GApplication) T.GApplicationFlags

	ApplicationSetFlags func(
		application *T.GApplication,
		flags T.GApplicationFlags)

	ApplicationSetActionGroup func(
		application *T.GApplication,
		actionGroup *T.GActionGroup)

	ApplicationGetIsRegistered func(
		application *T.GApplication) T.Gboolean

	ApplicationGetIsRemote func(
		application *T.GApplication) T.Gboolean

	ApplicationRegister func(
		application *T.GApplication,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	ApplicationHold func(
		application *T.GApplication)

	ApplicationRelease func(
		application *T.GApplication)

	ApplicationActivate func(
		application *T.GApplication)

	ApplicationOpen func(
		application *T.GApplication,
		files **T.GFile,
		nFiles int,
		hint string)

	ApplicationRun func(
		application *T.GApplication,
		argc int,
		argv **T.Char) int

	ApplicationCommandLineGetType func() O.Type

	ApplicationCommandLineGetArguments func(
		cmdline *T.GApplicationCommandLine,
		argc *int) **T.Gchar

	ApplicationCommandLineGetEnviron func(
		cmdline *T.GApplicationCommandLine) **T.Gchar

	ApplicationCommandLineGetenv func(
		cmdline *T.GApplicationCommandLine,
		name string) string

	ApplicationCommandLineGetCwd func(
		cmdline *T.GApplicationCommandLine) string

	ApplicationCommandLineGetIsRemote func(
		cmdline *T.GApplicationCommandLine) T.Gboolean

	ApplicationCommandLinePrint func(
		cmdline *T.GApplicationCommandLine, format string, v ...VArg)
	ApplicationCommandLinePrinterr func(
		cmdline *T.GApplicationCommandLine, format string, v ...VArg)

	ApplicationCommandLineGetExitStatus func(
		cmdline *T.GApplicationCommandLine) int

	ApplicationCommandLineSetExitStatus func(
		cmdline *T.GApplicationCommandLine,
		exitStatus int)

	ApplicationCommandLineGetPlatformData func(
		cmdline *T.GApplicationCommandLine) *T.GVariant

	InitableGetType func() O.Type

	InitableInit func(
		initable *T.GInitable,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	InitableNew func(objectType O.Type,
		cancellable *T.GCancellable, e **T.GError,
		firstPropertyName string, v ...VArg) T.Gpointer

	InitableNewv func(
		objectType O.Type,
		nParameters uint,
		parameters *T.GParameter,
		cancellable *T.GCancellable,
		err **T.GError) T.Gpointer

	InitableNewValist func(
		objectType O.Type,
		firstPropertyName string,
		varArgs T.VaList,
		cancellable *T.GCancellable,
		err **T.GError) *O.Object

	AsyncInitableGetType func() O.Type

	AsyncInitableInitAsync func(
		initable *T.GAsyncInitable,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	AsyncInitableInitFinish func(
		initable *T.GAsyncInitable,
		res *T.GAsyncResult,
		err **T.GError) T.Gboolean

	AsyncInitableNewAsync func(objectType O.Type,
		ioPriority int, cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback, userData T.Gpointer,
		firstPropertyName string, v ...VArg)

	AsyncInitableNewvAsync func(
		objectType O.Type,
		nParameters uint,
		parameters *T.GParameter,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	AsyncInitableNewValistAsync func(
		objectType O.Type,
		firstPropertyName string,
		varArgs T.VaList,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	AsyncInitableNewFinish func(
		initable *T.GAsyncInitable,
		res *T.GAsyncResult,
		err **T.GError) *O.Object

	AsyncResultGetType func() O.Type

	AsyncResultGetUserData func(
		res *T.GAsyncResult) T.Gpointer

	AsyncResultGetSourceObject func(
		res *T.GAsyncResult) *O.Object

	InputStreamGetType func() O.Type

	InputStreamRead func(
		stream *T.GInputStream,
		buffer *T.Void,
		count T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	InputStreamReadAll func(
		stream *T.GInputStream,
		buffer *T.Void,
		count T.Gsize,
		bytesRead *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	InputStreamSkip func(
		stream *T.GInputStream,
		count T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	InputStreamClose func(
		stream *T.GInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	InputStreamReadAsync func(
		stream *T.GInputStream,
		buffer *T.Void,
		count T.Gsize,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	InputStreamReadFinish func(
		stream *T.GInputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	InputStreamSkipAsync func(
		stream *T.GInputStream,
		count T.Gsize,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	InputStreamSkipFinish func(
		stream *T.GInputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	InputStreamCloseAsync func(
		stream *T.GInputStream,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	InputStreamCloseFinish func(
		stream *T.GInputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	InputStreamIsClosed func(
		stream *T.GInputStream) T.Gboolean

	InputStreamHasPending func(
		stream *T.GInputStream) T.Gboolean

	InputStreamSetPending func(
		stream *T.GInputStream,
		err **T.GError) T.Gboolean

	InputStreamClearPending func(
		stream *T.GInputStream)

	FilterInputStreamGetType func() O.Type

	FilterInputStreamGetBaseStream func(
		stream *T.GFilterInputStream) *T.GInputStream

	FilterInputStreamGetCloseBaseStream func(
		stream *T.GFilterInputStream) T.Gboolean

	FilterInputStreamSetCloseBaseStream func(
		stream *T.GFilterInputStream,
		closeBase T.Gboolean)

	BufferedInputStreamGetType func() O.Type

	BufferedInputStreamNew func(
		baseStream *T.GInputStream) *T.GInputStream

	BufferedInputStreamNewSized func(
		baseStream *T.GInputStream,
		size T.Gsize) *T.GInputStream

	BufferedInputStreamGetBufferSize func(
		stream *T.GBufferedInputStream) T.Gsize

	BufferedInputStreamSetBufferSize func(
		stream *T.GBufferedInputStream,
		size T.Gsize)

	BufferedInputStreamGetAvailable func(
		stream *T.GBufferedInputStream) T.Gsize

	BufferedInputStreamPeek func(
		stream *T.GBufferedInputStream,
		buffer *T.Void,
		offset T.Gsize,
		count T.Gsize) T.Gsize

	BufferedInputStreamPeekBuffer func(
		stream *T.GBufferedInputStream,
		count *T.Gsize) *T.Void

	BufferedInputStreamFill func(
		stream *T.GBufferedInputStream,
		count T.Gssize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	BufferedInputStreamFillAsync func(
		stream *T.GBufferedInputStream,
		count T.Gssize,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	BufferedInputStreamFillFinish func(
		stream *T.GBufferedInputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	BufferedInputStreamReadByte func(
		stream *T.GBufferedInputStream,
		cancellable *T.GCancellable,
		err **T.GError) int

	OutputStreamGetType func() O.Type

	OutputStreamWrite func(
		stream *T.GOutputStream,
		buffer *T.Void,
		count T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	OutputStreamWriteAll func(
		stream *T.GOutputStream,
		buffer *T.Void,
		count T.Gsize,
		bytesWritten *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	OutputStreamSplice func(
		stream *T.GOutputStream,
		source *T.GInputStream,
		flags T.GOutputStreamSpliceFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	OutputStreamFlush func(
		stream *T.GOutputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	OutputStreamClose func(
		stream *T.GOutputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	OutputStreamWriteAsync func(
		stream *T.GOutputStream,
		buffer *T.Void,
		count T.Gsize,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	OutputStreamWriteFinish func(
		stream *T.GOutputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	OutputStreamSpliceAsync func(
		stream *T.GOutputStream,
		source *T.GInputStream,
		flags T.GOutputStreamSpliceFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	OutputStreamSpliceFinish func(
		stream *T.GOutputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gssize

	OutputStreamFlushAsync func(
		stream *T.GOutputStream,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	OutputStreamFlushFinish func(
		stream *T.GOutputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	OutputStreamCloseAsync func(
		stream *T.GOutputStream,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	OutputStreamCloseFinish func(
		stream *T.GOutputStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	OutputStreamIsClosed func(
		stream *T.GOutputStream) T.Gboolean

	OutputStreamIsClosing func(
		stream *T.GOutputStream) T.Gboolean

	OutputStreamHasPending func(
		stream *T.GOutputStream) T.Gboolean

	OutputStreamSetPending func(
		stream *T.GOutputStream,
		err **T.GError) T.Gboolean

	OutputStreamClearPending func(
		stream *T.GOutputStream)

	FilterOutputStreamGetType func() O.Type

	FilterOutputStreamGetBaseStream func(
		stream *T.GFilterOutputStream) *T.GOutputStream

	FilterOutputStreamGetCloseBaseStream func(
		stream *T.GFilterOutputStream) T.Gboolean

	FilterOutputStreamSetCloseBaseStream func(
		stream *T.GFilterOutputStream,
		closeBase T.Gboolean)

	BufferedOutputStreamGetType func() O.Type

	BufferedOutputStreamNew func(
		baseStream *T.GOutputStream) *T.GOutputStream

	BufferedOutputStreamNewSized func(
		baseStream *T.GOutputStream,
		size T.Gsize) *T.GOutputStream

	BufferedOutputStreamGetBufferSize func(
		stream *T.GBufferedOutputStream) T.Gsize

	BufferedOutputStreamSetBufferSize func(
		stream *T.GBufferedOutputStream,
		size T.Gsize)

	BufferedOutputStreamGetAutoGrow func(
		stream *T.GBufferedOutputStream) T.Gboolean

	BufferedOutputStreamSetAutoGrow func(
		stream *T.GBufferedOutputStream,
		autoGrow T.Gboolean)

	CancellableGetType func() O.Type

	CancellableNew func() *T.GCancellable

	CancellableIsCancelled func(
		cancellable *T.GCancellable) T.Gboolean

	CancellableSetErrorIfCancelled func(
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	CancellableGetFd func(
		cancellable *T.GCancellable) int

	CancellableMakePollfd func(
		cancellable *T.GCancellable,
		pollfd *T.GPollFD) T.Gboolean

	CancellableReleaseFd func(
		cancellable *T.GCancellable)

	CancellableSourceNew func(
		cancellable *T.GCancellable) *T.GSource

	CancellableGetCurrent func() *T.GCancellable

	CancellablePushCurrent func(
		cancellable *T.GCancellable)

	CancellablePopCurrent func(
		cancellable *T.GCancellable)

	CancellableReset func(
		cancellable *T.GCancellable)

	CancellableConnect func(
		cancellable *T.GCancellable,
		callback T.GCallback,
		data T.Gpointer,
		dataDestroyFunc T.GDestroyNotify) T.Gulong

	CancellableDisconnect func(
		cancellable *T.GCancellable,
		handlerId T.Gulong)

	CancellableCancel func(
		cancellable *T.GCancellable)

	ConverterGetType func() O.Type

	ConverterConvert func(
		converter *T.GConverter,
		inbuf *T.Void,
		inbufSize T.Gsize,
		outbuf *T.Void,
		outbufSize T.Gsize,
		flags T.GConverterFlags,
		bytesRead *T.Gsize,
		bytesWritten *T.Gsize,
		err **T.GError) T.GConverterResult

	ConverterReset func(
		converter *T.GConverter)

	CharsetConverterGetType func() O.Type

	CharsetConverterNew func(
		toCharset string,
		fromCharset string,
		err **T.GError) *T.GcharsetConverter

	CharsetConverterSetUseFallback func(
		converter *T.GcharsetConverter,
		useFallback T.Gboolean)

	CharsetConverterGetUseFallback func(
		converter *T.GcharsetConverter) T.Gboolean

	CharsetConverterGetNumFallbacks func(
		converter *T.GcharsetConverter) uint

	ContentTypeEquals func(
		type1 string,
		type2 string) T.Gboolean

	ContentTypeIsA func(
		typ string,
		supertype string) T.Gboolean

	ContentTypeIsUnknown func(
		typ string) T.Gboolean

	ContentTypeGetDescription func(
		typ string) string

	ContentTypeGetMimeType func(
		typ string) string

	ContentTypeGetIcon func(
		typ string) *T.GIcon

	ContentTypeCanBeExecutable func(
		typ string) T.Gboolean

	ContentTypeFromMimeType func(
		mimeType string) string

	ContentTypeGuess func(
		filename string,
		data *T.Guchar,
		dataSize T.Gsize,
		resultUncertain *T.Gboolean) string

	ContentTypeGuessForTree func(
		root *T.GFile) **T.Gchar

	ContentTypesGetRegistered func() *T.GList

	ConverterInputStreamGetType func() O.Type

	ConverterInputStreamNew func(
		baseStream *T.GInputStream,
		converter *T.GConverter) *T.GInputStream

	ConverterInputStreamGetConverter func(
		converterStream *T.GConverterInputStream) *T.GConverter

	ConverterOutputStreamGetType func() O.Type

	ConverterOutputStreamNew func(
		baseStream *T.GOutputStream,
		converter *T.GConverter) *T.GOutputStream

	ConverterOutputStreamGetConverter func(
		converterStream *T.GConverterOutputStream) *T.GConverter

	CredentialsGetType func() O.Type

	CredentialsNew func() *T.GCredentials

	CredentialsToString func(
		credentials *T.GCredentials) string

	CredentialsGetNative func(
		credentials *T.GCredentials,
		nativeType T.GCredentialsType) T.Gpointer

	CredentialsSetNative func(
		credentials *T.GCredentials,
		nativeType T.GCredentialsType,
		native T.Gpointer)

	CredentialsIsSameUser func(
		credentials *T.GCredentials,
		otherCredentials *T.GCredentials,
		err **T.GError) T.Gboolean

	DataInputStreamGetType func() O.Type

	DataInputStreamNew func(
		baseStream *T.GInputStream) *T.GDataInputStream

	DataInputStreamSetByteOrder func(
		stream *T.GDataInputStream,
		order T.GDataStreamByteOrder)

	DataInputStreamGetByteOrder func(
		stream *T.GDataInputStream) T.GDataStreamByteOrder

	DataInputStreamSetNewlineType func(
		stream *T.GDataInputStream,
		typ T.GDataStreamNewlineType)

	DataInputStreamGetNewlineType func(
		stream *T.GDataInputStream) T.GDataStreamNewlineType

	DataInputStreamReadByte func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Guchar

	DataInputStreamReadInt16 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) int16

	DataInputStreamReadUint16 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) uint16

	DataInputStreamReadInt32 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.GInt32

	DataInputStreamReadUint32 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) T.GUint32

	DataInputStreamReadInt64 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) int64

	DataInputStreamReadUint64 func(
		stream *T.GDataInputStream,
		cancellable *T.GCancellable,
		err **T.GError) uint64

	DataInputStreamReadLine func(
		stream *T.GDataInputStream,
		length *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) string

	DataInputStreamReadLineAsync func(
		stream *T.GDataInputStream,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DataInputStreamReadLineFinish func(
		stream *T.GDataInputStream,
		result *T.GAsyncResult,
		length *T.Gsize,
		err **T.GError) string

	DataInputStreamReadUntil func(
		stream *T.GDataInputStream,
		stopChars string,
		length *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) string

	DataInputStreamReadUntilAsync func(
		stream *T.GDataInputStream,
		stopChars string,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DataInputStreamReadUntilFinish func(
		stream *T.GDataInputStream,
		result *T.GAsyncResult,
		length *T.Gsize,
		err **T.GError) string

	DataInputStreamReadUpto func(
		stream *T.GDataInputStream,
		stopChars string,
		stopCharsLen T.Gssize,
		length *T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) string

	DataInputStreamReadUptoAsync func(
		stream *T.GDataInputStream,
		stopChars string,
		stopCharsLen T.Gssize,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DataInputStreamReadUptoFinish func(
		stream *T.GDataInputStream,
		result *T.GAsyncResult,
		length *T.Gsize,
		err **T.GError) string

	DataOutputStreamGetType func() O.Type

	DataOutputStreamNew func(
		baseStream *T.GOutputStream) *T.GDataOutputStream

	DataOutputStreamSetByteOrder func(
		stream *T.GDataOutputStream,
		order T.GDataStreamByteOrder)

	DataOutputStreamGetByteOrder func(
		stream *T.GDataOutputStream) T.GDataStreamByteOrder

	DataOutputStreamPutByte func(
		stream *T.GDataOutputStream,
		data T.Guchar,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DataOutputStreamPutInt16 func(
		stream *T.GDataOutputStream,
		data int16,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DataOutputStreamPutUint16 func(
		stream *T.GDataOutputStream,
		data uint16,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DataOutputStreamPutInt32 func(
		stream *T.GDataOutputStream,
		data T.GInt32,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DataOutputStreamPutUint32 func(
		stream *T.GDataOutputStream,
		data T.GUint32,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DataOutputStreamPutInt64 func(
		stream *T.GDataOutputStream,
		data int64,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DataOutputStreamPutUint64 func(
		stream *T.GDataOutputStream,
		data uint64,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DataOutputStreamPutString func(
		stream *T.GDataOutputStream,
		str string,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DbusIsAddress func(
		string string) T.Gboolean

	DbusIsSupportedAddress func(
		string string,
		err **T.GError) T.Gboolean

	DbusAddressGetStream func(
		address string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusAddressGetStreamFinish func(
		res *T.GAsyncResult,
		outGuid **T.Gchar,
		err **T.GError) *T.GIOStream

	DbusAddressGetStreamSync func(
		address string,
		outGuid **T.Gchar,
		cancellable *T.GCancellable,
		err **T.GError) *T.GIOStream

	DbusAddressGetForBusSync func(
		busType T.GBusType,
		cancellable *T.GCancellable,
		err **T.GError) string

	DbusAuthObserverGetType func() O.Type

	DbusAuthObserverNew func() *T.GDBusAuthObserver

	DbusAuthObserverAuthorizeAuthenticatedPeer func(
		observer *T.GDBusAuthObserver,
		stream *T.GIOStream,
		credentials *T.GCredentials) T.Gboolean

	DbusConnectionGetType func() O.Type

	BusGet func(
		busType T.GBusType,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	BusGetFinish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusConnection

	BusGetSync func(
		busType T.GBusType,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusConnection

	DbusConnectionNew func(
		stream *T.GIOStream,
		guid string,
		flags T.GDBusConnectionFlags,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusConnectionNewFinish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusConnection

	DbusConnectionNewSync func(
		stream *T.GIOStream,
		guid string,
		flags T.GDBusConnectionFlags,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusConnection

	DbusConnectionNewForAddress func(
		address string,
		flags T.GDBusConnectionFlags,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusConnectionNewForAddressFinish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusConnection

	DbusConnectionNewForAddressSync func(
		address string,
		flags T.GDBusConnectionFlags,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusConnection

	DbusConnectionStartMessageProcessing func(
		connection *T.GDBusConnection)

	DbusConnectionIsClosed func(
		connection *T.GDBusConnection) T.Gboolean

	DbusConnectionGetStream func(
		connection *T.GDBusConnection) *T.GIOStream

	DbusConnectionGetGuid func(
		connection *T.GDBusConnection) string

	DbusConnectionGetUniqueName func(
		connection *T.GDBusConnection) string

	DbusConnectionGetPeerCredentials func(
		connection *T.GDBusConnection) *T.GCredentials

	DbusConnectionGetExitOnClose func(
		connection *T.GDBusConnection) T.Gboolean

	DbusConnectionSetExitOnClose func(
		connection *T.GDBusConnection,
		exitOnClose T.Gboolean)

	DbusConnectionGetCapabilities func(
		connection *T.GDBusConnection) T.GDBusCapabilityFlags

	DbusConnectionClose func(
		connection *T.GDBusConnection,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusConnectionCloseFinish func(
		connection *T.GDBusConnection,
		res *T.GAsyncResult,
		err **T.GError) T.Gboolean

	DbusConnectionCloseSync func(
		connection *T.GDBusConnection,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DbusConnectionFlush func(
		connection *T.GDBusConnection,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusConnectionFlushFinish func(
		connection *T.GDBusConnection,
		res *T.GAsyncResult,
		err **T.GError) T.Gboolean

	DbusConnectionFlushSync func(
		connection *T.GDBusConnection,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	DbusConnectionSendMessage func(
		connection *T.GDBusConnection,
		message *T.GDBusMessage,
		flags T.GDBusSendMessageFlags,
		outSerial *T.GUint32,
		err **T.GError) T.Gboolean

	DbusConnectionSendMessageWithReply func(
		connection *T.GDBusConnection,
		message *T.GDBusMessage,
		flags T.GDBusSendMessageFlags,
		timeoutMsec int,
		outSerial *T.GUint32,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusConnectionSendMessageWithReplyFinish func(
		connection *T.GDBusConnection,
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusMessage

	DbusConnectionSendMessageWithReplySync func(
		connection *T.GDBusConnection,
		message *T.GDBusMessage,
		flags T.GDBusSendMessageFlags,
		timeoutMsec int,
		outSerial *T.GUint32,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusMessage

	DbusConnectionEmitSignal func(
		connection *T.GDBusConnection,
		destinationBusName string,
		objectPath string,
		interfaceName string,
		signalName string,
		parameters *T.GVariant,
		err **T.GError) T.Gboolean

	DbusConnectionCall func(
		connection *T.GDBusConnection,
		busName string,
		objectPath string,
		interfaceName string,
		methodName string,
		parameters *T.GVariant,
		replyType *T.GVariantType,
		flags T.GDBusCallFlags,
		timeoutMsec int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusConnectionCallFinish func(
		connection *T.GDBusConnection,
		res *T.GAsyncResult,
		err **T.GError) *T.GVariant

	DbusConnectionCallSync func(
		connection *T.GDBusConnection,
		busName string,
		objectPath string,
		interfaceName string,
		methodName string,
		parameters *T.GVariant,
		replyType *T.GVariantType,
		flags T.GDBusCallFlags,
		timeoutMsec int,
		cancellable *T.GCancellable,
		err **T.GError) *T.GVariant

	DbusConnectionRegisterObject func(
		connection *T.GDBusConnection,
		objectPath string,
		interfaceInfo *T.GDBusInterfaceInfo,
		vtable *T.GDBusInterfaceVTable,
		userData T.Gpointer,
		userDataFreeFunc T.GDestroyNotify,
		err **T.GError) uint

	DbusConnectionUnregisterObject func(
		connection *T.GDBusConnection,
		registrationId uint) T.Gboolean

	DbusConnectionRegisterSubtree func(
		connection *T.GDBusConnection,
		objectPath string,
		vtable *T.GDBusSubtreeVTable,
		flags T.GDBusSubtreeFlags,
		userData T.Gpointer,
		userDataFreeFunc T.GDestroyNotify,
		err **T.GError) uint

	DbusConnectionUnregisterSubtree func(
		connection *T.GDBusConnection,
		registrationId uint) T.Gboolean

	DbusConnectionSignalSubscribe func(
		connection *T.GDBusConnection,
		sender string,
		interfaceName string,
		member string,
		objectPath string,
		arg0 string,
		flags T.GDBusSignalFlags,
		callback T.GDBusSignalCallback,
		userData T.Gpointer,
		userDataFreeFunc T.GDestroyNotify) uint

	DbusConnectionSignalUnsubscribe func(
		connection *T.GDBusConnection,
		subscriptionId uint)

	DbusConnectionAddFilter func(
		connection *T.GDBusConnection,
		filterFunction T.GDBusMessageFilterFunction,
		userData T.Gpointer,
		userDataFreeFunc T.GDestroyNotify) uint

	DbusConnectionRemoveFilter func(
		connection *T.GDBusConnection,
		filterId uint)

	DbusErrorQuark func() T.GQuark

	DbusErrorIsRemoteError func(
		err *T.GError) T.Gboolean

	DbusErrorGetRemoteError func(
		err *T.GError) string

	DbusErrorStripRemoteError func(
		err *T.GError) T.Gboolean

	DbusErrorRegisterError func(
		errorDomain T.GQuark,
		errorCode int,
		dbusErrorName string) T.Gboolean

	DbusErrorUnregisterError func(
		errorDomain T.GQuark,
		errorCode int,
		dbusErrorName string) T.Gboolean

	DbusErrorRegisterErrorDomain func(
		errorDomainQuarkName string,
		quarkVolatile *T.Gsize,
		entries *T.GDBusErrorEntry,
		numEntries uint)

	DbusErrorNewForDbusError func(
		dbusErrorName string,
		dbusErrorMessage string) *T.GError

	DbusErrorSetDbusError func(e **T.GError,
		dbusErrorName, dbusErrorMessage,
		format string, v ...VArg)

	DbusErrorSetDbusErrorValist func(
		err **T.GError,
		dbusErrorName string,
		dbusErrorMessage string,
		format string,
		varArgs T.VaList)

	DbusErrorEncodeGerror func(
		err *T.GError) string

	DbusAnnotationInfoLookup func(
		annotations **T.GDBusAnnotationInfo,
		name string) string

	DbusInterfaceInfoLookupMethod func(
		info *T.GDBusInterfaceInfo,
		name string) *T.GDBusMethodInfo

	DbusInterfaceInfoLookupSignal func(
		info *T.GDBusInterfaceInfo,
		name string) *T.GDBusSignalInfo

	DbusInterfaceInfoLookupProperty func(
		info *T.GDBusInterfaceInfo,
		name string) *T.GDBusPropertyInfo

	DbusInterfaceInfoGenerateXml func(
		info *T.GDBusInterfaceInfo,
		indent uint,
		stringBuilder *T.GString)

	DbusNodeInfoNewForXml func(
		xmlData string,
		err **T.GError) *T.GDBusNodeInfo

	DbusNodeInfoLookupInterface func(
		info *T.GDBusNodeInfo,
		name string) *T.GDBusInterfaceInfo

	DbusNodeInfoGenerateXml func(
		info *T.GDBusNodeInfo,
		indent uint,
		stringBuilder *T.GString)

	DbusNodeInfoRef func(
		info *T.GDBusNodeInfo) *T.GDBusNodeInfo

	DbusInterfaceInfoRef func(
		info *T.GDBusInterfaceInfo) *T.GDBusInterfaceInfo

	DbusMethodInfoRef func(
		info *T.GDBusMethodInfo) *T.GDBusMethodInfo

	DbusSignalInfoRef func(
		info *T.GDBusSignalInfo) *T.GDBusSignalInfo

	DbusPropertyInfoRef func(
		info *T.GDBusPropertyInfo) *T.GDBusPropertyInfo

	DbusArgInfoRef func(
		info *T.GDBusArgInfo) *T.GDBusArgInfo

	DbusAnnotationInfoRef func(
		info *T.GDBusAnnotationInfo) *T.GDBusAnnotationInfo

	DbusNodeInfoUnref func(
		info *T.GDBusNodeInfo)

	DbusInterfaceInfoUnref func(
		info *T.GDBusInterfaceInfo)

	DbusMethodInfoUnref func(
		info *T.GDBusMethodInfo)

	DbusSignalInfoUnref func(
		info *T.GDBusSignalInfo)

	DbusPropertyInfoUnref func(
		info *T.GDBusPropertyInfo)

	DbusArgInfoUnref func(
		info *T.GDBusArgInfo)

	DbusAnnotationInfoUnref func(
		info *T.GDBusAnnotationInfo)

	DbusNodeInfoGetType func() O.Type

	DbusInterfaceInfoGetType func() O.Type

	DbusMethodInfoGetType func() O.Type

	DbusSignalInfoGetType func() O.Type

	DbusPropertyInfoGetType func() O.Type

	DbusArgInfoGetType func() O.Type

	DbusAnnotationInfoGetType func() O.Type

	DbusMessageGetType func() O.Type

	DbusMessageNew func() *T.GDBusMessage

	DbusMessageNewSignal func(
		path string,
		interface_ string,
		signal string) *T.GDBusMessage

	DbusMessageNewMethodCall func(
		name string,
		path string,
		interface_ string,
		method string) *T.GDBusMessage

	DbusMessageNewMethodReply func(
		methodCallMessage *T.GDBusMessage) *T.GDBusMessage

	DbusMessageNewMethodError func(
		methodCallMessage *T.GDBusMessage,
		errorName, errorMessageFormat string,
		v ...VArg) *T.GDBusMessage

	DbusMessageNewMethodErrorValist func(
		methodCallMessage *T.GDBusMessage,
		errorName string,
		errorMessageFormat string,
		varArgs T.VaList) *T.GDBusMessage

	DbusMessageNewMethodErrorLiteral func(
		methodCallMessage *T.GDBusMessage,
		errorName string,
		errorMessage string) *T.GDBusMessage

	DbusMessagePrint func(
		message *T.GDBusMessage,
		indent uint) string

	DbusMessageGetLocked func(
		message *T.GDBusMessage) T.Gboolean

	DbusMessageLock func(
		message *T.GDBusMessage)

	DbusMessageCopy func(
		message *T.GDBusMessage,
		err **T.GError) *T.GDBusMessage

	DbusMessageGetByteOrder func(
		message *T.GDBusMessage) T.GDBusMessageByteOrder

	DbusMessageSetByteOrder func(
		message *T.GDBusMessage,
		byteOrder T.GDBusMessageByteOrder)

	DbusMessageGetMessageType func(
		message *T.GDBusMessage) T.GDBusMessageType

	DbusMessageSetMessageType func(
		message *T.GDBusMessage,
		typ T.GDBusMessageType)

	DbusMessageGetFlags func(
		message *T.GDBusMessage) T.GDBusMessageFlags

	DbusMessageSetFlags func(
		message *T.GDBusMessage,
		flags T.GDBusMessageFlags)

	DbusMessageGetSerial func(
		message *T.GDBusMessage) T.GUint32

	DbusMessageSetSerial func(
		message *T.GDBusMessage,
		serial T.GUint32)

	DbusMessageGetHeader func(
		message *T.GDBusMessage,
		headerField T.GDBusMessageHeaderField) *T.GVariant

	DbusMessageSetHeader func(
		message *T.GDBusMessage,
		headerField T.GDBusMessageHeaderField,
		value *T.GVariant)

	DbusMessageGetHeaderFields func(
		message *T.GDBusMessage) *T.Guchar

	DbusMessageGetBody func(
		message *T.GDBusMessage) *T.GVariant

	DbusMessageSetBody func(
		message *T.GDBusMessage,
		body *T.GVariant)

	DbusMessageGetUnixFdList func(
		message *T.GDBusMessage) *T.GUnixFDList

	DbusMessageSetUnixFdList func(
		message *T.GDBusMessage,
		fdList *T.GUnixFDList)

	DbusMessageGetReplySerial func(
		message *T.GDBusMessage) T.GUint32

	DbusMessageSetReplySerial func(
		message *T.GDBusMessage,
		value T.GUint32)

	DbusMessageGetInterface func(
		message *T.GDBusMessage) string

	DbusMessageSetInterface func(
		message *T.GDBusMessage,
		value string)

	DbusMessageGetMember func(
		message *T.GDBusMessage) string

	DbusMessageSetMember func(
		message *T.GDBusMessage,
		value string)

	DbusMessageGetPath func(
		message *T.GDBusMessage) string

	DbusMessageSetPath func(
		message *T.GDBusMessage,
		value string)

	DbusMessageGetSender func(
		message *T.GDBusMessage) string

	DbusMessageSetSender func(
		message *T.GDBusMessage,
		value string)

	DbusMessageGetDestination func(
		message *T.GDBusMessage) string

	DbusMessageSetDestination func(
		message *T.GDBusMessage,
		value string)

	DbusMessageGetErrorName func(
		message *T.GDBusMessage) string

	DbusMessageSetErrorName func(
		message *T.GDBusMessage,
		value string)

	DbusMessageGetSignature func(
		message *T.GDBusMessage) string

	DbusMessageSetSignature func(
		message *T.GDBusMessage,
		value string)

	DbusMessageGetNumUnixFds func(
		message *T.GDBusMessage) T.GUint32

	DbusMessageSetNumUnixFds func(
		message *T.GDBusMessage,
		value T.GUint32)

	DbusMessageGetArg0 func(
		message *T.GDBusMessage) string

	DbusMessageNewFromBlob func(
		blob *T.Guchar,
		blobLen T.Gsize,
		capabilities T.GDBusCapabilityFlags,
		err **T.GError) *T.GDBusMessage

	DbusMessageBytesNeeded func(
		blob *T.Guchar,
		blobLen T.Gsize,
		err **T.GError) T.Gssize

	DbusMessageToBlob func(
		message *T.GDBusMessage,
		outSize *T.Gsize,
		capabilities T.GDBusCapabilityFlags,
		err **T.GError) *T.Guchar

	DbusMessageToGerror func(
		message *T.GDBusMessage,
		err **T.GError) T.Gboolean

	DbusMethodInvocationGetType func() O.Type

	DbusMethodInvocationGetSender func(
		invocation *T.GDBusMethodInvocation) string

	DbusMethodInvocationGetObjectPath func(
		invocation *T.GDBusMethodInvocation) string

	DbusMethodInvocationGetInterfaceName func(
		invocation *T.GDBusMethodInvocation) string

	DbusMethodInvocationGetMethodName func(
		invocation *T.GDBusMethodInvocation) string

	DbusMethodInvocationGetMethodInfo func(
		invocation *T.GDBusMethodInvocation) *T.GDBusMethodInfo

	DbusMethodInvocationGetConnection func(
		invocation *T.GDBusMethodInvocation) *T.GDBusConnection

	DbusMethodInvocationGetMessage func(
		invocation *T.GDBusMethodInvocation) *T.GDBusMessage

	DbusMethodInvocationGetParameters func(
		invocation *T.GDBusMethodInvocation) *T.GVariant

	DbusMethodInvocationGetUserData func(
		invocation *T.GDBusMethodInvocation) T.Gpointer

	DbusMethodInvocationReturnValue func(
		invocation *T.GDBusMethodInvocation,
		parameters *T.GVariant)

	DbusMethodInvocationReturnError func(
		invocation *T.GDBusMethodInvocation,
		domain T.GQuark, code int, format string, v ...VArg)

	DbusMethodInvocationReturnErrorValist func(
		invocation *T.GDBusMethodInvocation,
		domain T.GQuark,
		code int,
		format string,
		varArgs T.VaList)

	DbusMethodInvocationReturnErrorLiteral func(
		invocation *T.GDBusMethodInvocation,
		domain T.GQuark,
		code int,
		message string)

	DbusMethodInvocationReturnGerror func(
		invocation *T.GDBusMethodInvocation,
		err *T.GError)

	DbusMethodInvocationReturnDbusError func(
		invocation *T.GDBusMethodInvocation,
		errorName string,
		errorMessage string)

	BusOwnName func(
		busType T.GBusType,
		name string,
		flags T.GBusNameOwnerFlags,
		busAcquiredHandler T.GBusAcquiredCallback,
		nameAcquiredHandler T.GBusNameAcquiredCallback,
		nameLostHandler T.GBusNameLostCallback,
		userData T.Gpointer,
		userDataFreeFunc T.GDestroyNotify) uint

	BusOwnNameOnConnection func(
		connection *T.GDBusConnection,
		name string,
		flags T.GBusNameOwnerFlags,
		nameAcquiredHandler T.GBusNameAcquiredCallback,
		nameLostHandler T.GBusNameLostCallback,
		userData T.Gpointer,
		userDataFreeFunc T.GDestroyNotify) uint

	BusOwnNameWithClosures func(
		busType T.GBusType,
		name string,
		flags T.GBusNameOwnerFlags,
		busAcquiredClosure *T.GClosure,
		nameAcquiredClosure *T.GClosure,
		nameLostClosure *T.GClosure) uint

	BusOwnNameOnConnectionWithClosures func(

		connection *T.GDBusConnection,
		name string,
		flags T.GBusNameOwnerFlags,
		nameAcquiredClosure *T.GClosure,
		nameLostClosure *T.GClosure) uint

	BusUnownName func(
		ownerId uint)

	BusWatchName func(
		busType T.GBusType,
		name string,
		flags T.GBusNameWatcherFlags,
		nameAppearedHandler T.GBusNameAppearedCallback,
		nameVanishedHandler T.GBusNameVanishedCallback,
		userData T.Gpointer,
		userDataFreeFunc T.GDestroyNotify) uint

	BusWatchNameOnConnection func(
		connection *T.GDBusConnection,
		name string,
		flags T.GBusNameWatcherFlags,
		nameAppearedHandler T.GBusNameAppearedCallback,
		nameVanishedHandler T.GBusNameVanishedCallback,
		userData T.Gpointer,
		userDataFreeFunc T.GDestroyNotify) uint

	BusWatchNameWithClosures func(
		busType T.GBusType,
		name string,
		flags T.GBusNameWatcherFlags,
		nameAppearedClosure *T.GClosure,
		nameVanishedClosure *T.GClosure) uint

	BusWatchNameOnConnectionWithClosures func(

		connection *T.GDBusConnection,
		name string,
		flags T.GBusNameWatcherFlags,
		nameAppearedClosure *T.GClosure,
		nameVanishedClosure *T.GClosure) uint

	BusUnwatchName func(
		watcherId uint)

	DbusProxyGetType func() O.Type

	DbusProxyNew func(
		connection *T.GDBusConnection,
		flags T.GDBusProxyFlags,
		info *T.GDBusInterfaceInfo,
		name string,
		objectPath string,
		interfaceName string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusProxyNewFinish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusProxy

	DbusProxyNewSync func(
		connection *T.GDBusConnection,
		flags T.GDBusProxyFlags,
		info *T.GDBusInterfaceInfo,
		name string,
		objectPath string,
		interfaceName string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusProxy

	DbusProxyNewForBus func(
		busType T.GBusType,
		flags T.GDBusProxyFlags,
		info *T.GDBusInterfaceInfo,
		name string,
		objectPath string,
		interfaceName string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusProxyNewForBusFinish func(
		res *T.GAsyncResult,
		err **T.GError) *T.GDBusProxy

	DbusProxyNewForBusSync func(
		busType T.GBusType,
		flags T.GDBusProxyFlags,
		info *T.GDBusInterfaceInfo,
		name string,
		objectPath string,
		interfaceName string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusProxy

	DbusProxyGetConnection func(
		proxy *T.GDBusProxy) *T.GDBusConnection

	DbusProxyGetFlags func(
		proxy *T.GDBusProxy) T.GDBusProxyFlags

	DbusProxyGetName func(
		proxy *T.GDBusProxy) string

	DbusProxyGetNameOwner func(
		proxy *T.GDBusProxy) string

	DbusProxyGetObjectPath func(
		proxy *T.GDBusProxy) string

	DbusProxyGetInterfaceName func(
		proxy *T.GDBusProxy) string

	DbusProxyGetDefaultTimeout func(
		proxy *T.GDBusProxy) int

	DbusProxySetDefaultTimeout func(
		proxy *T.GDBusProxy,
		timeoutMsec int)

	DbusProxyGetInterfaceInfo func(
		proxy *T.GDBusProxy) *T.GDBusInterfaceInfo

	DbusProxySetInterfaceInfo func(
		proxy *T.GDBusProxy,
		info *T.GDBusInterfaceInfo)

	DbusProxyGetCachedProperty func(
		proxy *T.GDBusProxy,
		propertyName string) *T.GVariant

	DbusProxySetCachedProperty func(
		proxy *T.GDBusProxy,
		propertyName string,
		value *T.GVariant)

	DbusProxyGetCachedPropertyNames func(
		proxy *T.GDBusProxy) **T.Gchar

	DbusProxyCall func(
		proxy *T.GDBusProxy,
		methodName string,
		parameters *T.GVariant,
		flags T.GDBusCallFlags,
		timeoutMsec int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DbusProxyCallFinish func(
		proxy *T.GDBusProxy,
		res *T.GAsyncResult,
		err **T.GError) *T.GVariant

	DbusProxyCallSync func(
		proxy *T.GDBusProxy,
		methodName string,
		parameters *T.GVariant,
		flags T.GDBusCallFlags,
		timeoutMsec int,
		cancellable *T.GCancellable,
		err **T.GError) *T.GVariant

	DbusServerGetType func() O.Type

	DbusServerNewSync func(
		address string,
		flags T.GDBusServerFlags,
		guid string,
		observer *T.GDBusAuthObserver,
		cancellable *T.GCancellable,
		err **T.GError) *T.GDBusServer

	DbusServerGetClientAddress func(
		server *T.GDBusServer) string

	DbusServerGetGuid func(
		server *T.GDBusServer) string

	DbusServerGetFlags func(
		server *T.GDBusServer) T.GDBusServerFlags

	DbusServerStart func(
		server *T.GDBusServer)

	DbusServerStop func(
		server *T.GDBusServer)

	DbusServerIsActive func(
		server *T.GDBusServer) T.Gboolean

	DbusIsGuid func(
		string string) T.Gboolean

	DbusGenerateGuid func() string

	DbusIsName func(
		string string) T.Gboolean

	DbusIsUniqueName func(
		string string) T.Gboolean

	DbusIsMemberName func(
		string string) T.Gboolean

	DbusIsInterfaceName func(
		string string) T.Gboolean

	DriveGetType func() O.Type

	DriveGetName func(drive *T.GDrive) string

	DriveGetIcon func(drive *T.GDrive) *T.GIcon

	DriveHasVolumes func(drive *T.GDrive) T.Gboolean

	DriveGetVolumes func(drive *T.GDrive) *T.GList

	DriveIsMediaRemovable func(drive *T.GDrive) T.Gboolean

	DriveHasMedia func(drive *T.GDrive) T.Gboolean

	DriveIsMediaCheckAutomatic func(drive *T.GDrive) T.Gboolean

	DriveCanPollForMedia func(drive *T.GDrive) T.Gboolean

	DriveCanEject func(drive *T.GDrive) T.Gboolean

	DriveEject func(
		drive *T.GDrive,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DriveEjectFinish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	DrivePollForMedia func(
		drive *T.GDrive,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DrivePollForMediaFinish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	DriveGetIdentifier func(drive *T.GDrive, kind string) string

	DriveEnumerateIdentifiers func(
		drive *T.GDrive) **T.Char

	DriveGetStartStopType func(
		drive *T.GDrive) T.GDriveStartStopType

	DriveCanStart func(drive *T.GDrive) T.Gboolean

	DriveCanStartDegraded func(drive *T.GDrive) T.Gboolean

	DriveStart func(
		drive *T.GDrive,
		flags T.GDriveStartFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DriveStartFinish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	DriveCanStop func(drive *T.GDrive) T.Gboolean

	DriveStop func(
		drive *T.GDrive,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DriveStopFinish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	DriveEjectWithOperation func(
		drive *T.GDrive,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	DriveEjectWithOperationFinish func(
		drive *T.GDrive,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	IconGetType func() O.Type

	IconHash func(icon T.Gconstpointer) uint

	IconEqual func(icon1 *T.GIcon, icon2 *T.GIcon) T.Gboolean

	IconToString func(icon *T.GIcon) string

	IconNewForString func(str string, err **T.GError) *T.GIcon

	EmblemGetType func() O.Type

	EmblemNew func(icon *T.GIcon) *T.GEmblem

	EmblemNewWithOrigin func(
		icon *T.GIcon, origin T.GEmblemOrigin) *T.GEmblem

	EmblemGetIcon func(emblem *T.GEmblem) *T.GIcon

	EmblemGetOrigin func(emblem *T.GEmblem) T.GEmblemOrigin

	EmblemedIconGetType func() O.Type

	EmblemedIconNew func(icon *T.GIcon, emblem *T.GEmblem) *T.GIcon

	EmblemedIconGetIcon func(emblemed *T.GEmblemedIcon) *T.GIcon

	EmblemedIconGetEmblems func(
		emblemed *T.GEmblemedIcon) *T.GList

	EmblemedIconAddEmblem func(
		emblemed *T.GEmblemedIcon, emblem *T.GEmblem)

	EmblemedIconClearEmblems func(emblemed *T.GEmblemedIcon)

	FileAttributeInfoListGetType func() O.Type

	FileAttributeInfoListNew func() *T.GFileAttributeInfoList

	FileAttributeInfoListRef func(
		list *T.GFileAttributeInfoList) *T.GFileAttributeInfoList

	FileAttributeInfoListUnref func(
		list *T.GFileAttributeInfoList)

	FileAttributeInfoListDup func(
		list *T.GFileAttributeInfoList) *T.GFileAttributeInfoList

	FileAttributeInfoListLookup func(
		list *T.GFileAttributeInfoList,
		name string) *T.GFileAttributeInfo

	FileAttributeInfoListAdd func(
		list *T.GFileAttributeInfoList,
		name string,
		typ T.GFileAttributeType,
		flags T.GFileAttributeInfoFlags)

	FileEnumeratorGetType func() O.Type

	FileEnumeratorNextFile func(
		enumerator *T.GFileEnumerator,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	FileEnumeratorClose func(
		enumerator *T.GFileEnumerator,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileEnumeratorNextFilesAsync func(
		enumerator *T.GFileEnumerator,
		numFiles int,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileEnumeratorNextFilesFinish func(
		enumerator *T.GFileEnumerator,
		result *T.GAsyncResult,
		err **T.GError) *T.GList

	FileEnumeratorCloseAsync func(
		enumerator *T.GFileEnumerator,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileEnumeratorCloseFinish func(
		enumerator *T.GFileEnumerator,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileEnumeratorIsClosed func(
		enumerator *T.GFileEnumerator) T.Gboolean

	FileEnumeratorHasPending func(
		enumerator *T.GFileEnumerator) T.Gboolean

	FileEnumeratorSetPending func(
		enumerator *T.GFileEnumerator,
		pending T.Gboolean)

	FileEnumeratorGetContainer func(
		enumerator *T.GFileEnumerator) *T.GFile

	FileGetType func() O.Type

	FileNewForPath func(path string) *T.GFile

	FileNewForUri func(uri string) *T.GFile

	FileNewForCommandlineArg func(arg string) *T.GFile

	FileParseName func(parseName string) *T.GFile

	FileDup func(file *T.GFile) *T.GFile

	FileHash func(file T.Gconstpointer) uint

	FileEqual func(file1 *T.GFile, file2 *T.GFile) T.Gboolean

	FileGetBasename func(file *T.GFile) string

	FileGetPath func(file *T.GFile) string

	FileGetUri func(file *T.GFile) string

	FileGetParseName func(file *T.GFile) string

	FileGetParent func(file *T.GFile) *T.GFile

	FileHasParent func(file *T.GFile, parent *T.GFile) T.Gboolean

	FileGetChild func(file *T.GFile, name string) *T.GFile

	FileGetChildForDisplayName func(
		file *T.GFile, displayName string, err **T.GError) *T.GFile

	FileHasPrefix func(file *T.GFile, prefix *T.GFile) T.Gboolean

	FileGetRelativePath func(
		parent *T.GFile, descendant *T.GFile) string

	FileResolveRelativePath func(
		file *T.GFile, relativePath string) *T.GFile

	FileIsNative func(file *T.GFile) T.Gboolean

	FileHasUriScheme func(
		file *T.GFile, uriScheme string) T.Gboolean

	FileGetUriScheme func(file *T.GFile) string

	FileRead func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInputStream

	FileReadAsync func(
		file *T.GFile,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileReadFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileInputStream

	FileAppendTo func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileOutputStream

	FileCreate func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileOutputStream

	FileReplace func(
		file *T.GFile,
		etag string,
		makeBackup T.Gboolean,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileOutputStream

	FileAppendToAsync func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileAppendToFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileOutputStream

	FileCreateAsync func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileCreateFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileOutputStream

	FileReplaceAsync func(
		file *T.GFile,
		etag string,
		makeBackup T.Gboolean,
		flags T.GFileCreateFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileReplaceFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileOutputStream

	FileOpenReadwrite func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileIOStream

	FileOpenReadwriteAsync func(
		file *T.GFile,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileOpenReadwriteFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileIOStream

	FileCreateReadwrite func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileIOStream

	FileCreateReadwriteAsync func(
		file *T.GFile,
		flags T.GFileCreateFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileCreateReadwriteFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileIOStream

	FileReplaceReadwrite func(
		file *T.GFile,
		etag string,
		makeBackup T.Gboolean,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileIOStream

	FileReplaceReadwriteAsync func(
		file *T.GFile,
		etag string,
		makeBackup T.Gboolean,
		flags T.GFileCreateFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileReplaceReadwriteFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileIOStream

	FileQueryExists func(
		file *T.GFile,
		cancellable *T.GCancellable) T.Gboolean

	FileQueryFileType func(
		file *T.GFile,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable) T.GFileType

	FileQueryInfo func(
		file *T.GFile,
		attributes string,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	FileQueryInfoAsync func(
		file *T.GFile,
		attributes string,
		flags T.GFileQueryInfoFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileQueryInfoFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	FileQueryFilesystemInfo func(
		file *T.GFile,
		attributes string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	FileQueryFilesystemInfoAsync func(
		file *T.GFile,
		attributes string,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileQueryFilesystemInfoFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	FileFindEnclosingMount func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GMount

	FileFindEnclosingMountAsync func(
		file *T.GFile,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileFindEnclosingMountFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GMount

	FileEnumerateChildren func(
		file *T.GFile,
		attributes string,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileEnumerator

	FileEnumerateChildrenAsync func(
		file *T.GFile,
		attributes string,
		flags T.GFileQueryInfoFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileEnumerateChildrenFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFileEnumerator

	FileSetDisplayName func(
		file *T.GFile,
		displayName string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFile

	FileSetDisplayNameAsync func(
		file *T.GFile,
		displayName string,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileSetDisplayNameFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) *T.GFile

	FileDelete func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileTrash func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileCopy func(
		source *T.GFile,
		destination *T.GFile,
		flags T.GFileCopyFlags,
		cancellable *T.GCancellable,
		progressCallback T.GFileProgressCallback,
		progressCallbackData T.Gpointer,
		err **T.GError) T.Gboolean

	FileCopyAsync func(
		source *T.GFile,
		destination *T.GFile,
		flags T.GFileCopyFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		progressCallback T.GFileProgressCallback,
		progressCallbackData T.Gpointer,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileCopyFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileMove func(
		source *T.GFile,
		destination *T.GFile,
		flags T.GFileCopyFlags,
		cancellable *T.GCancellable,
		progressCallback T.GFileProgressCallback,
		progressCallbackData T.Gpointer,
		err **T.GError) T.Gboolean

	FileMakeDirectory func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileMakeDirectoryWithParents func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileMakeSymbolicLink func(
		file *T.GFile,
		symlinkValue string,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileQuerySettableAttributes func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileAttributeInfoList

	FileQueryWritableNamespaces func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileAttributeInfoList

	FileSetAttribute func(
		file *T.GFile,
		attribute string,
		typ T.GFileAttributeType,
		valueP T.Gpointer,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileSetAttributesFromInfo func(
		file *T.GFile,
		info *T.GFileInfo,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileSetAttributesAsync func(
		file *T.GFile,
		info *T.GFileInfo,
		flags T.GFileQueryInfoFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileSetAttributesFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		info **T.GFileInfo,
		err **T.GError) T.Gboolean

	FileSetAttributeString func(
		file *T.GFile,
		attribute string,
		value string,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileSetAttributeByteString func(
		file *T.GFile,
		attribute string,
		value string,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileSetAttributeUint32 func(
		file *T.GFile,
		attribute string,
		value T.GUint32,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileSetAttributeInt32 func(
		file *T.GFile,
		attribute string,
		value T.GInt32,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileSetAttributeUint64 func(
		file *T.GFile,
		attribute string,
		value uint64,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileSetAttributeInt64 func(
		file *T.GFile,
		attribute string,
		value int64,
		flags T.GFileQueryInfoFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileMountEnclosingVolume func(
		location *T.GFile,
		flags T.GMountMountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileMountEnclosingVolumeFinish func(
		location *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileMountMountable func(
		file *T.GFile,
		flags T.GMountMountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileMountMountableFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) *T.GFile

	FileUnmountMountable func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileUnmountMountableFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileUnmountMountableWithOperation func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileUnmountMountableWithOperationFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileEjectMountable func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileEjectMountableFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileEjectMountableWithOperation func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileEjectMountableWithOperationFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileCopyAttributes func(
		source *T.GFile,
		destination *T.GFile,
		flags T.GFileCopyFlags,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileMonitorDirectory func(
		file *T.GFile,
		flags T.GFileMonitorFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileMonitor

	FileMonitorFile func(
		file *T.GFile,
		flags T.GFileMonitorFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileMonitor

	FileMonitor func(
		file *T.GFile,
		flags T.GFileMonitorFlags,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileMonitor

	FileStartMountable func(
		file *T.GFile,
		flags T.GDriveStartFlags,
		startOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileStartMountableFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileStopMountable func(
		file *T.GFile,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileStopMountableFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FilePollMountable func(
		file *T.GFile,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FilePollMountableFinish func(
		file *T.GFile,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	FileQueryDefaultHandler func(
		file *T.GFile,
		cancellable *T.GCancellable,
		err **T.GError) *T.GAppInfo

	FileLoadContents func(
		file *T.GFile,
		cancellable *T.GCancellable,
		contents **T.Char,
		length *T.Gsize,
		etagOut **T.Char,
		err **T.GError) T.Gboolean

	FileLoadContentsAsync func(
		file *T.GFile,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileLoadContentsFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		contents **T.Char,
		length *T.Gsize,
		etagOut **T.Char,
		err **T.GError) T.Gboolean

	FileLoadPartialContentsAsync func(
		file *T.GFile,
		cancellable *T.GCancellable,
		readMoreCallback T.GFileReadMoreCallback,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileLoadPartialContentsFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		contents **T.Char,
		length *T.Gsize,
		etagOut **T.Char,
		err **T.GError) T.Gboolean

	FileReplaceContents func(
		file *T.GFile,
		contents string,
		length T.Gsize,
		etag string,
		makeBackup T.Gboolean,
		flags T.GFileCreateFlags,
		newEtag **T.Char,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	FileReplaceContentsAsync func(
		file *T.GFile,
		contents string,
		length T.Gsize,
		etag string,
		makeBackup T.Gboolean,
		flags T.GFileCreateFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileReplaceContentsFinish func(
		file *T.GFile,
		res *T.GAsyncResult,
		newEtag **T.Char,
		err **T.GError) T.Gboolean

	FileSupportsThreadContexts func(file *T.GFile) T.Gboolean

	FileIconGetType func() O.Type

	FileIconNew func(file *T.GFile) *T.GIcon

	FileIconGetFile func(icon *T.GFileIcon) *T.GFile

	FileInfoGetType func() O.Type

	FileInfoNew func() *T.GFileInfo

	FileInfoDup func(other *T.GFileInfo) *T.GFileInfo

	FileInfoCopyInto func(srcInfo, destInfo *T.GFileInfo)

	FileInfoHasAttribute func(
		info *T.GFileInfo, attribute string) T.Gboolean

	FileInfoHasNamespace func(
		info *T.GFileInfo, nameSpace string) T.Gboolean

	FileInfoListAttributes func(
		info *T.GFileInfo, nameSpace string) **T.Char

	FileInfoGetAttributeData func(
		info *T.GFileInfo,
		attribute string,
		typ *T.GFileAttributeType,
		valuePp *T.Gpointer,
		status *T.GFileAttributeStatus) T.Gboolean

	FileInfoGetAttributeType func(
		info *T.GFileInfo,
		attribute string) T.GFileAttributeType

	FileInfoRemoveAttribute func(
		info *T.GFileInfo,
		attribute string)

	FileInfoGetAttributeStatus func(
		info *T.GFileInfo,
		attribute string) T.GFileAttributeStatus

	FileInfoSetAttributeStatus func(
		info *T.GFileInfo,
		attribute string,
		status T.GFileAttributeStatus) T.Gboolean

	FileInfoGetAttributeAsString func(
		info *T.GFileInfo,
		attribute string) string

	FileInfoGetAttributeString func(
		info *T.GFileInfo,
		attribute string) string

	FileInfoGetAttributeByteString func(
		info *T.GFileInfo,
		attribute string) string

	FileInfoGetAttributeBoolean func(
		info *T.GFileInfo,
		attribute string) T.Gboolean

	FileInfoGetAttributeUint32 func(
		info *T.GFileInfo,
		attribute string) T.GUint32

	FileInfoGetAttributeInt32 func(
		info *T.GFileInfo,
		attribute string) T.GInt32

	FileInfoGetAttributeUint64 func(
		info *T.GFileInfo,
		attribute string) uint64

	FileInfoGetAttributeInt64 func(
		info *T.GFileInfo,
		attribute string) int64

	FileInfoGetAttributeObject func(
		info *T.GFileInfo,
		attribute string) *O.Object

	FileInfoGetAttributeStringv func(
		info *T.GFileInfo,
		attribute string) **T.Char

	FileInfoSetAttribute func(
		info *T.GFileInfo,
		attribute string,
		typ T.GFileAttributeType,
		valueP T.Gpointer)

	FileInfoSetAttributeString func(
		info *T.GFileInfo,
		attribute string,
		attrValue string)

	FileInfoSetAttributeByteString func(
		info *T.GFileInfo,
		attribute string,
		attrValue string)

	FileInfoSetAttributeBoolean func(
		info *T.GFileInfo,
		attribute string,
		attrValue T.Gboolean)

	FileInfoSetAttributeUint32 func(
		info *T.GFileInfo,
		attribute string,
		attrValue T.GUint32)

	FileInfoSetAttributeInt32 func(
		info *T.GFileInfo,
		attribute string,
		attrValue T.GInt32)

	FileInfoSetAttributeUint64 func(
		info *T.GFileInfo,
		attribute string,
		attrValue uint64)

	FileInfoSetAttributeInt64 func(
		info *T.GFileInfo,
		attribute string,
		attrValue int64)

	FileInfoSetAttributeObject func(
		info *T.GFileInfo,
		attribute string,
		attrValue *O.Object)

	FileInfoSetAttributeStringv func(
		info *T.GFileInfo,
		attribute string,
		attrValue **T.Char)

	FileInfoClearStatus func(
		info *T.GFileInfo)

	FileInfoGetFileType func(
		info *T.GFileInfo) T.GFileType

	FileInfoGetIsHidden func(
		info *T.GFileInfo) T.Gboolean

	FileInfoGetIsBackup func(
		info *T.GFileInfo) T.Gboolean

	FileInfoGetIsSymlink func(
		info *T.GFileInfo) T.Gboolean

	FileInfoGetName func(
		info *T.GFileInfo) string

	FileInfoGetDisplayName func(
		info *T.GFileInfo) string

	FileInfoGetEditName func(
		info *T.GFileInfo) string

	FileInfoGetIcon func(
		info *T.GFileInfo) *T.GIcon

	FileInfoGetContentType func(
		info *T.GFileInfo) string

	FileInfoGetSize func(
		info *T.GFileInfo) T.Goffset

	FileInfoGetModificationTime func(
		info *T.GFileInfo,
		result *T.GTimeVal)

	FileInfoGetSymlinkTarget func(
		info *T.GFileInfo) string

	FileInfoGetEtag func(
		info *T.GFileInfo) string

	FileInfoGetSortOrder func(
		info *T.GFileInfo) T.GInt32

	FileInfoSetAttributeMask func(
		info *T.GFileInfo,
		mask *T.GFileAttributeMatcher)

	FileInfoUnsetAttributeMask func(
		info *T.GFileInfo)

	FileInfoSetFileType func(
		info *T.GFileInfo,
		typ T.GFileType)

	FileInfoSetIsHidden func(
		info *T.GFileInfo,
		isHidden T.Gboolean)

	FileInfoSetIsSymlink func(
		info *T.GFileInfo,
		isSymlink T.Gboolean)

	FileInfoSetName func(
		info *T.GFileInfo,
		name string)

	FileInfoSetDisplayName func(
		info *T.GFileInfo,
		displayName string)

	FileInfoSetEditName func(
		info *T.GFileInfo,
		editName string)

	FileInfoSetIcon func(
		info *T.GFileInfo,
		icon *T.GIcon)

	FileInfoSetContentType func(
		info *T.GFileInfo,
		contentType string)

	FileInfoSetSize func(
		info *T.GFileInfo,
		size T.Goffset)

	FileInfoSetModificationTime func(
		info *T.GFileInfo,
		mtime *T.GTimeVal)

	FileInfoSetSymlinkTarget func(
		info *T.GFileInfo,
		symlinkTarget string)

	FileInfoSetSortOrder func(
		info *T.GFileInfo,
		sortOrder T.GInt32)

	FileAttributeMatcherGetType func() O.Type

	FileAttributeMatcherNew func(
		attributes string) *T.GFileAttributeMatcher

	FileAttributeMatcherRef func(
		matcher *T.GFileAttributeMatcher) *T.GFileAttributeMatcher

	FileAttributeMatcherUnref func(
		matcher *T.GFileAttributeMatcher)

	FileAttributeMatcherMatches func(
		matcher *T.GFileAttributeMatcher,
		attribute string) T.Gboolean

	FileAttributeMatcherMatchesOnly func(
		matcher *T.GFileAttributeMatcher,
		attribute string) T.Gboolean

	FileAttributeMatcherEnumerateNamespace func(
		matcher *T.GFileAttributeMatcher,
		ns string) T.Gboolean

	FileAttributeMatcherEnumerateNext func(
		matcher *T.GFileAttributeMatcher) string

	FileInputStreamGetType func() O.Type

	FileInputStreamQueryInfo func(
		stream *T.GFileInputStream,
		attributes string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	FileInputStreamQueryInfoAsync func(
		stream *T.GFileInputStream,
		attributes string,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileInputStreamQueryInfoFinish func(
		stream *T.GFileInputStream,
		result *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	IoErrorQuark func() T.GQuark

	IoErrorFromErrno func(
		errNo int) T.GIOErrorEnum

	IoErrorFromWin32Error func(
		errorCode int) T.GIOErrorEnum

	IoStreamGetType func() O.Type

	IoStreamGetInputStream func(
		stream *T.GIOStream) *T.GInputStream

	IoStreamGetOutputStream func(
		stream *T.GIOStream) *T.GOutputStream

	IoStreamSpliceAsync func(
		stream1 *T.GIOStream,
		stream2 *T.GIOStream,
		flags T.GIOStreamSpliceFlags,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	IoStreamSpliceFinish func(
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	IoStreamClose func(
		stream *T.GIOStream,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	IoStreamCloseAsync func(
		stream *T.GIOStream,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	IoStreamCloseFinish func(
		stream *T.GIOStream,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	IoStreamIsClosed func(
		stream *T.GIOStream) T.Gboolean

	IoStreamHasPending func(
		stream *T.GIOStream) T.Gboolean

	IoStreamSetPending func(
		stream *T.GIOStream,
		err **T.GError) T.Gboolean

	IoStreamClearPending func(
		stream *T.GIOStream)

	FileIoStreamGetType func() O.Type

	FileIoStreamQueryInfo func(
		stream *T.GFileIOStream,
		attributes string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	FileIoStreamQueryInfoAsync func(
		stream *T.GFileIOStream,
		attributes string,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileIoStreamQueryInfoFinish func(
		stream *T.GFileIOStream,
		result *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	FileIoStreamGetEtag func(
		stream *T.GFileIOStream) string

	FileMonitorGetType func() O.Type

	FileMonitorCancel func(
		monitor *T.GFileMonitor) T.Gboolean

	FileMonitorIsCancelled func(
		monitor *T.GFileMonitor) T.Gboolean

	FileMonitorSetRateLimit func(
		monitor *T.GFileMonitor,
		limitMsecs int)

	FileMonitorEmitEvent func(
		monitor *T.GFileMonitor,
		child *T.GFile,
		otherFile *T.GFile,
		eventType T.GFileMonitorEvent)

	FilenameCompleterGetType func() O.Type

	FilenameCompleterNew func() *T.GFilenameCompleter

	FilenameCompleterGetCompletionSuffix func(
		completer *T.GFilenameCompleter,
		initialText string) string

	FilenameCompleterGetCompletions func(
		completer *T.GFilenameCompleter,
		initialText string) **T.Char

	FilenameCompleterSetDirsOnly func(
		completer *T.GFilenameCompleter,
		dirsOnly T.Gboolean)

	FileOutputStreamGetType func() O.Type

	FileOutputStreamQueryInfo func(
		stream *T.GFileOutputStream,
		attributes string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GFileInfo

	FileOutputStreamQueryInfoAsync func(
		stream *T.GFileOutputStream,
		attributes string,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	FileOutputStreamQueryInfoFinish func(
		stream *T.GFileOutputStream,
		result *T.GAsyncResult,
		err **T.GError) *T.GFileInfo

	FileOutputStreamGetEtag func(
		stream *T.GFileOutputStream) string

	InetAddressGetType func() O.Type

	InetAddressNewFromString func(
		string string) *T.GInetAddress

	InetAddressNewFromBytes func(
		bytes *uint8,
		family T.GSocketFamily) *T.GInetAddress

	InetAddressNewLoopback func(
		family T.GSocketFamily) *T.GInetAddress

	InetAddressNewAny func(
		family T.GSocketFamily) *T.GInetAddress

	InetAddressToString func(
		address *T.GInetAddress) string

	InetAddressToBytes func(
		address *T.GInetAddress) *uint8

	InetAddressGetNativeSize func(
		address *T.GInetAddress) T.Gsize

	InetAddressGetFamily func(
		address *T.GInetAddress) T.GSocketFamily

	InetAddressGetIsAny func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsLoopback func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsLinkLocal func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsSiteLocal func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsMulticast func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsMcGlobal func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsMcLinkLocal func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsMcNodeLocal func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsMcOrgLocal func(
		address *T.GInetAddress) T.Gboolean

	InetAddressGetIsMcSiteLocal func(
		address *T.GInetAddress) T.Gboolean

	SocketAddressGetType func() O.Type

	SocketAddressGetFamily func(
		address *T.GSocketAddress) T.GSocketFamily

	SocketAddressNewFromNative func(
		native T.Gpointer,
		len T.Gsize) *T.GSocketAddress

	SocketAddressToNative func(
		address *T.GSocketAddress,
		dest T.Gpointer,
		destlen T.Gsize,
		err **T.GError) T.Gboolean

	SocketAddressGetNativeSize func(
		address *T.GSocketAddress) T.Gssize

	InetSocketAddressGetType func() O.Type

	InetSocketAddressNew func(
		address *T.GInetAddress,
		port uint16) *T.GSocketAddress

	InetSocketAddressGetAddress func(
		address *T.GInetSocketAddress) *T.GInetAddress

	InetSocketAddressGetPort func(
		address *T.GInetSocketAddress) uint16

	AppInfoCreateFlagsGetType func() O.Type

	ConverterFlagsGetType func() O.Type

	ConverterResultGetType func() O.Type

	DataStreamByteOrderGetType func() O.Type

	DataStreamNewlineTypeGetType func() O.Type

	FileAttributeTypeGetType func() O.Type

	FileAttributeInfoFlagsGetType func() O.Type

	FileAttributeStatusGetType func() O.Type

	FileQueryInfoFlagsGetType func() O.Type

	FileCreateFlagsGetType func() O.Type

	MountMountFlagsGetType func() O.Type

	MountUnmountFlagsGetType func() O.Type

	DriveStartFlagsGetType func() O.Type

	DriveStartStopTypeGetType func() O.Type

	FileCopyFlagsGetType func() O.Type

	FileMonitorFlagsGetType func() O.Type

	FileTypeGetType func() O.Type

	FilesystemPreviewTypeGetType func() O.Type

	FileMonitorEventGetType func() O.Type

	IoErrorEnumGetType func() O.Type

	AskPasswordFlagsGetType func() O.Type

	PasswordSaveGetType func() O.Type

	MountOperationResultGetType func() O.Type

	OutputStreamSpliceFlagsGetType func() O.Type

	IoStreamSpliceFlagsGetType func() O.Type

	EmblemOriginGetType func() O.Type

	ResolverErrorGetType func() O.Type

	SocketFamilyGetType func() O.Type

	SocketTypeGetType func() O.Type

	SocketMsgFlagsGetType func() O.Type

	SocketProtocolGetType func() O.Type

	ZlibCompressorFormatGetType func() O.Type

	UnixSocketAddressTypeGetType func() O.Type

	BusTypeGetType func() O.Type

	BusNameOwnerFlagsGetType func() O.Type

	BusNameWatcherFlagsGetType func() O.Type

	DbusProxyFlagsGetType func() O.Type

	DbusErrorGetType func() O.Type

	DbusConnectionFlagsGetType func() O.Type

	DbusCapabilityFlagsGetType func() O.Type

	DbusCallFlagsGetType func() O.Type

	DbusMessageTypeGetType func() O.Type

	DbusMessageFlagsGetType func() O.Type

	DbusMessageHeaderFieldGetType func() O.Type

	DbusPropertyInfoFlagsGetType func() O.Type

	DbusSubtreeFlagsGetType func() O.Type

	DbusServerFlagsGetType func() O.Type

	DbusSignalFlagsGetType func() O.Type

	DbusSendMessageFlagsGetType func() O.Type

	CredentialsTypeGetType func() O.Type

	DbusMessageByteOrderGetType func() O.Type

	ApplicationFlagsGetType func() O.Type

	TlsErrorGetType func() O.Type

	TlsCertificateFlagsGetType func() O.Type

	TlsAuthenticationModeGetType func() O.Type

	TlsRehandshakeModeGetType func() O.Type

	SettingsBindFlagsGetType func() O.Type

	IoModuleGetType func() O.Type

	IoModuleNew func(
		filename string) *T.GIOModule

	IoModulesScanAllInDirectory func(
		dirname string)

	IoModulesLoadAllInDirectory func(
		dirname string) *T.GList

	IoExtensionPointRegister func(
		name string) *T.GIOExtensionPoint

	IoExtensionPointLookup func(
		name string) *T.GIOExtensionPoint

	IoExtensionPointSetRequiredType func(
		extensionPoint *T.GIOExtensionPoint,
		typ O.Type)

	IoExtensionPointGetRequiredType func(
		extensionPoint *T.GIOExtensionPoint) O.Type

	IoExtensionPointGetExtensions func(
		extensionPoint *T.GIOExtensionPoint) *T.GList

	IoExtensionPointGetExtensionByName func(
		extensionPoint *T.GIOExtensionPoint,
		name string) *T.GIOExtension

	IoExtensionPointImplement func(
		extensionPointName string,
		typ O.Type,
		extensionName string,
		priority int) *T.GIOExtension

	IoExtensionGetType func(
		extension *T.GIOExtension) O.Type

	IoExtensionGetName func(
		extension *T.GIOExtension) string

	IoExtensionGetPriority func(
		extension *T.GIOExtension) int

	IoExtensionRefClass func(
		extension *T.GIOExtension) *O.TypeClass

	IoModuleLoad func(
		module *T.GIOModule)

	IoModuleUnload func(
		module *T.GIOModule)

	IoModuleQuery func() **T.Char

	IoSchedulerPushJob func(
		jobFunc T.GIOSchedulerJobFunc,
		userData T.Gpointer,
		notify T.GDestroyNotify,
		ioPriority int,
		cancellable *T.GCancellable)

	IoSchedulerCancelAllJobs func()

	IoSchedulerJobSendToMainloop func(
		job *T.GIOSchedulerJob,
		f O.SourceFunc,
		userData T.Gpointer,
		notify T.GDestroyNotify) T.Gboolean

	IoSchedulerJobSendToMainloopAsync func(
		job *T.GIOSchedulerJob,
		f O.SourceFunc,
		userData T.Gpointer,
		notify T.GDestroyNotify)

	LoadableIconGetType func() O.Type

	LoadableIconLoad func(
		icon *T.GLoadableIcon,
		size int,
		t **T.Char,
		cancellable *T.GCancellable,
		err **T.GError) *T.GInputStream

	LoadableIconLoadAsync func(
		icon *T.GLoadableIcon,
		size int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	LoadableIconLoadFinish func(
		icon *T.GLoadableIcon,
		res *T.GAsyncResult,
		typ **T.Char,
		err **T.GError) *T.GInputStream

	MemoryInputStreamGetType func() O.Type

	MemoryInputStreamNew func() *T.GInputStream

	MemoryInputStreamNewFromData func(
		data *T.Void,
		len T.Gssize,
		destroy T.GDestroyNotify) *T.GInputStream

	MemoryInputStreamAddData func(
		stream *T.GMemoryInputStream,
		data *T.Void,
		len T.Gssize,
		destroy T.GDestroyNotify)

	MemoryOutputStreamGetType func() O.Type

	MemoryOutputStreamNew func(
		data T.Gpointer,
		size T.Gsize,
		reallocFunction T.GReallocFunc,
		destroyFunction T.GDestroyNotify) *T.GOutputStream

	MemoryOutputStreamGetData func(
		ostream *T.GMemoryOutputStream) T.Gpointer

	MemoryOutputStreamGetSize func(
		ostream *T.GMemoryOutputStream) T.Gsize

	MemoryOutputStreamGetDataSize func(
		ostream *T.GMemoryOutputStream) T.Gsize

	MemoryOutputStreamStealData func(
		ostream *T.GMemoryOutputStream) T.Gpointer

	MountGetType func() O.Type

	MountGetRoot func(
		mount *T.GMount) *T.GFile

	MountGetDefaultLocation func(
		mount *T.GMount) *T.GFile

	MountGetName func(
		mount *T.GMount) string

	MountGetIcon func(
		mount *T.GMount) *T.GIcon

	MountGetUuid func(
		mount *T.GMount) string

	MountGetVolume func(
		mount *T.GMount) *T.GVolume

	MountGetDrive func(
		mount *T.GMount) *T.GDrive

	MountCanUnmount func(
		mount *T.GMount) T.Gboolean

	MountCanEject func(
		mount *T.GMount) T.Gboolean

	MountUnmount func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	MountUnmountFinish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	MountEject func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	MountEjectFinish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	MountRemount func(
		mount *T.GMount,
		flags T.GMountMountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	MountRemountFinish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	MountGuessContentType func(
		mount *T.GMount,
		forceRescan T.Gboolean,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	MountGuessContentTypeFinish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) **T.Gchar

	MountGuessContentTypeSync func(
		mount *T.GMount,
		forceRescan T.Gboolean,
		cancellable *T.GCancellable,
		err **T.GError) **T.Gchar

	MountIsShadowed func(
		mount *T.GMount) T.Gboolean

	MountShadow func(
		mount *T.GMount)

	MountUnshadow func(
		mount *T.GMount)

	MountUnmountWithOperation func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	MountUnmountWithOperationFinish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	MountEjectWithOperation func(
		mount *T.GMount,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	MountEjectWithOperationFinish func(
		mount *T.GMount,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	MountOperationGetType func() O.Type

	MountOperationNew func() *T.GMountOperation

	MountOperationGetUsername func(
		op *T.GMountOperation) string

	MountOperationSetUsername func(
		op *T.GMountOperation,
		username string)

	MountOperationGetPassword func(
		op *T.GMountOperation) string

	MountOperationSetPassword func(
		op *T.GMountOperation,
		password string)

	MountOperationGetAnonymous func(
		op *T.GMountOperation) T.Gboolean

	MountOperationSetAnonymous func(
		op *T.GMountOperation,
		anonymous T.Gboolean)

	MountOperationGetDomain func(
		op *T.GMountOperation) string

	MountOperationSetDomain func(
		op *T.GMountOperation,
		domain string)

	MountOperationGetPasswordSave func(
		op *T.GMountOperation) T.GPasswordSave

	MountOperationSetPasswordSave func(
		op *T.GMountOperation,
		save T.GPasswordSave)

	MountOperationGetChoice func(
		op *T.GMountOperation) int

	MountOperationSetChoice func(
		op *T.GMountOperation,
		choice int)

	MountOperationReply func(
		op *T.GMountOperation,
		result T.GMountOperationResult)

	VolumeMonitorGetType func() O.Type

	VolumeMonitorGet func() *T.GVolumeMonitor

	VolumeMonitorGetConnectedDrives func(
		volumeMonitor *T.GVolumeMonitor) *T.GList

	VolumeMonitorGetVolumes func(
		volumeMonitor *T.GVolumeMonitor) *T.GList

	VolumeMonitorGetMounts func(
		volumeMonitor *T.GVolumeMonitor) *T.GList

	VolumeMonitorGetVolumeForUuid func(
		volumeMonitor *T.GVolumeMonitor,
		uuid string) *T.GVolume

	VolumeMonitorGetMountForUuid func(
		volumeMonitor *T.GVolumeMonitor,
		uuid string) *T.GMount

	VolumeMonitorAdoptOrphanMount func(
		mount *T.GMount) *T.GVolume

	NativeVolumeMonitorGetType func() O.Type

	NetworkAddressGetType func() O.Type

	NetworkAddressNew func(
		hostname string,
		port uint16) *T.GSocketConnectable

	NetworkAddressParse func(
		hostAndPort string,
		defaultPort uint16,
		err **T.GError) *T.GSocketConnectable

	NetworkAddressParseUri func(
		uri string,
		defaultPort uint16,
		err **T.GError) *T.GSocketConnectable

	NetworkAddressGetHostname func(
		addr *T.GNetworkAddress) string

	NetworkAddressGetPort func(
		addr *T.GNetworkAddress) uint16

	NetworkAddressGetScheme func(
		addr *T.GNetworkAddress) string

	NetworkServiceGetType func() O.Type

	NetworkServiceNew func(
		service string,
		protocol string,
		domain string) *T.GSocketConnectable

	NetworkServiceGetService func(
		srv *T.GNetworkService) string

	NetworkServiceGetProtocol func(
		srv *T.GNetworkService) string

	NetworkServiceGetDomain func(
		srv *T.GNetworkService) string

	NetworkServiceGetScheme func(
		srv *T.GNetworkService) string

	NetworkServiceSetScheme func(
		srv *T.GNetworkService,
		scheme string)

	PermissionGetType func() O.Type

	PermissionAcquire func(
		permission *T.GPermission,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	PermissionAcquireAsync func(
		permission *T.GPermission,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	PermissionAcquireFinish func(
		permission *T.GPermission,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	PermissionRelease func(
		permission *T.GPermission,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	PermissionReleaseAsync func(
		permission *T.GPermission,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	PermissionReleaseFinish func(
		permission *T.GPermission,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	PermissionGetAllowed func(
		permission *T.GPermission) T.Gboolean

	PermissionGetCanAcquire func(
		permission *T.GPermission) T.Gboolean

	PermissionGetCanRelease func(
		permission *T.GPermission) T.Gboolean

	PermissionImplUpdate func(
		permission *T.GPermission,
		allowed T.Gboolean,
		canAcquire T.Gboolean,
		canRelease T.Gboolean)

	PollableInputStreamGetType func() O.Type

	PollableInputStreamCanPoll func(
		stream *T.GPollableInputStream) T.Gboolean

	PollableInputStreamIsReadable func(
		stream *T.GPollableInputStream) T.Gboolean

	PollableInputStreamCreateSource func(
		stream *T.GPollableInputStream,
		cancellable *T.GCancellable) *T.GSource

	PollableInputStreamReadNonblocking func(
		stream *T.GPollableInputStream,
		buffer *T.Void,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	PollableSourceNew func(
		pollableStream *O.Object) *T.GSource

	PollableOutputStreamGetType func() O.Type

	PollableOutputStreamCanPoll func(
		stream *T.GPollableOutputStream) T.Gboolean

	PollableOutputStreamIsWritable func(
		stream *T.GPollableOutputStream) T.Gboolean

	PollableOutputStreamCreateSource func(
		stream *T.GPollableOutputStream,
		cancellable *T.GCancellable) *T.GSource

	PollableOutputStreamWriteNonblocking func(
		stream *T.GPollableOutputStream,
		buffer *T.Void,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	ProxyGetType func() O.Type

	ProxyGetDefaultForProtocol func(
		protocol string) *T.GProxy

	ProxyConnect func(
		proxy *T.GProxy,
		connection *T.GIOStream,
		proxyAddress *T.GProxyAddress,
		cancellable *T.GCancellable,
		err **T.GError) *T.GIOStream

	ProxyConnectAsync func(
		proxy *T.GProxy,
		connection *T.GIOStream,
		proxyAddress *T.GProxyAddress,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	ProxyConnectFinish func(
		proxy *T.GProxy,
		result *T.GAsyncResult,
		err **T.GError) *T.GIOStream

	ProxySupportsHostname func(
		proxy *T.GProxy) T.Gboolean

	ProxyAddressGetType func() O.Type

	ProxyAddressNew func(
		inetaddr *T.GInetAddress,
		port uint16,
		protocol string,
		destHostname string,
		destPort uint16,
		username string,
		password string) *T.GSocketAddress

	ProxyAddressGetProtocol func(
		proxy *T.GProxyAddress) string

	ProxyAddressGetDestinationHostname func(
		proxy *T.GProxyAddress) string

	ProxyAddressGetDestinationPort func(
		proxy *T.GProxyAddress) uint16

	ProxyAddressGetUsername func(
		proxy *T.GProxyAddress) string

	ProxyAddressGetPassword func(
		proxy *T.GProxyAddress) string

	SocketAddressEnumeratorGetType func() O.Type

	SocketAddressEnumeratorNext func(
		enumerator *T.GSocketAddressEnumerator,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketAddress

	SocketAddressEnumeratorNextAsync func(
		enumerator *T.GSocketAddressEnumerator,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	SocketAddressEnumeratorNextFinish func(
		enumerator *T.GSocketAddressEnumerator,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketAddress

	ProxyAddressEnumeratorGetType func() O.Type

	ProxyResolverGetType func() O.Type

	ProxyResolverGetDefault func() *T.GProxyResolver

	ProxyResolverIsSupported func(
		resolver *T.GProxyResolver) T.Gboolean

	ProxyResolverLookup func(
		resolver *T.GProxyResolver,
		uri string,
		cancellable *T.GCancellable,
		err **T.GError) **T.Gchar

	ProxyResolverLookupAsync func(
		resolver *T.GProxyResolver,
		uri string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	ProxyResolverLookupFinish func(
		resolver *T.GProxyResolver,
		result *T.GAsyncResult,
		err **T.GError) **T.Gchar

	ResolverGetType func() O.Type

	ResolverGetDefault func() *T.GResolver

	ResolverSetDefault func(
		resolver *T.GResolver)

	ResolverLookupByName func(
		resolver *T.GResolver,
		hostname string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GList

	ResolverLookupByNameAsync func(
		resolver *T.GResolver,
		hostname string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	ResolverLookupByNameFinish func(
		resolver *T.GResolver,
		result *T.GAsyncResult,
		err **T.GError) *T.GList

	ResolverFreeAddresses func(
		addresses *T.GList)

	ResolverLookupByAddress func(
		resolver *T.GResolver,
		address *T.GInetAddress,
		cancellable *T.GCancellable,
		err **T.GError) string

	ResolverLookupByAddressAsync func(
		resolver *T.GResolver,
		address *T.GInetAddress,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	ResolverLookupByAddressFinish func(
		resolver *T.GResolver,
		result *T.GAsyncResult,
		err **T.GError) string

	ResolverLookupService func(
		resolver *T.GResolver,
		service string,
		protocol string,
		domain string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GList

	ResolverLookupServiceAsync func(
		resolver *T.GResolver,
		service string,
		protocol string,
		domain string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	ResolverLookupServiceFinish func(
		resolver *T.GResolver,
		result *T.GAsyncResult,
		err **T.GError) *T.GList

	ResolverFreeTargets func(
		targets *T.GList)

	ResolverErrorQuark func() T.GQuark

	SeekableGetType func() O.Type

	SeekableTell func(
		seekable *T.GSeekable) T.Goffset

	SeekableCanSeek func(
		seekable *T.GSeekable) T.Gboolean

	SeekableSeek func(
		seekable *T.GSeekable,
		offset T.Goffset,
		typ T.GSeekType,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	SeekableCanTruncate func(
		seekable *T.GSeekable) T.Gboolean

	SeekableTruncate func(
		seekable *T.GSeekable,
		offset T.Goffset,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	SettingsGetType func() O.Type

	SettingsListSchemas func() **T.Gchar

	SettingsListRelocatableSchemas func() **T.Gchar

	SettingsNew func(
		schema string) *T.GSettings

	SettingsNewWithPath func(
		schema string,
		path string) *T.GSettings

	SettingsNewWithBackend func(
		schema string,
		backend *T.GSettingsBackend) *T.GSettings

	SettingsNewWithBackendAndPath func(
		schema string,
		backend *T.GSettingsBackend,
		path string) *T.GSettings

	SettingsListChildren func(
		settings *T.GSettings) **T.Gchar

	SettingsList_keys func(
		settings *T.GSettings) **T.Gchar

	SettingsGetRange func(
		settings *T.GSettings,
		key string) *T.GVariant

	SettingsRangeCheck func(
		settings *T.GSettings,
		key string,
		value *T.GVariant) T.Gboolean

	SettingsSetValue func(
		settings *T.GSettings,
		key string,
		value *T.GVariant) T.Gboolean

	SettingsGetValue func(
		settings *T.GSettings,
		key string) *T.GVariant

	SettingsSet func(settings *T.GSettings, key, format string,
		v ...VArg) T.Gboolean

	SettingsGet func(settings *T.GSettings, key, format string,
		v ...VArg)

	SettingsReset func(settings *T.GSettings, key string)

	SettingsGetInt func(settings *T.GSettings, key string) int

	SettingsSetInt func(
		settings *T.GSettings, key string, value int) T.Gboolean

	SettingsGetString func(
		settings *T.GSettings, key string) string

	SettingsSetString func(
		settings *T.GSettings, key, value string) T.Gboolean

	SettingsGetBoolean func(
		settings *T.GSettings, key string) T.Gboolean

	SettingsSetBoolean func(
		settings *T.GSettings, key string, value T.Gboolean) T.Gboolean

	SettingsGetDouble func(
		settings *T.GSettings, key string) float64

	SettingsSetDouble func(settings *T.GSettings,
		key string, value float64) T.Gboolean

	SettingsGetStrv func(
		settings *T.GSettings, key string) **T.Gchar

	SettingsSetStrv func(settings *T.GSettings,
		key string, value **T.Gchar) T.Gboolean

	SettingsGetEnum func(
		settings *T.GSettings, key string) int

	SettingsSetEnum func(settings *T.GSettings,
		key string, value int) T.Gboolean

	SettingsGetFlags func(
		settings *T.GSettings, key string) uint

	SettingsSetFlags func(settings *T.GSettings,
		key string, value uint) T.Gboolean

	SettingsGetChild func(
		settings *T.GSettings,
		name string) *T.GSettings

	SettingsIsWritable func(
		settings *T.GSettings, name string) T.Gboolean

	SettingsDelay func(settings *T.GSettings)

	SettingsApply func(settings *T.GSettings)

	SettingsRevert func(settings *T.GSettings)

	SettingsGetHasUnapplied func(
		settings *T.GSettings) T.Gboolean

	SettingsSync func()

	SettingsBind func(
		settings *T.GSettings,
		key string,
		object T.Gpointer,
		property string,
		flags T.GSettingsBindFlags)

	SettingsBindWithMapping func(
		settings *T.GSettings,
		key string,
		object T.Gpointer,
		property string,
		flags T.GSettingsBindFlags,
		getMapping T.GSettingsBindGetMapping,
		setMapping T.GSettingsBindSetMapping,
		userData T.Gpointer,
		destroy T.GDestroyNotify)

	SettingsBindWritable func(
		settings *T.GSettings,
		key string,
		object T.Gpointer,
		property string,
		inverted T.Gboolean)

	SettingsUnbind func(object T.Gpointer, property string)

	SettingsGetMapped func(
		settings *T.GSettings,
		key string,
		mapping T.GSettingsGetMapping,
		userData T.Gpointer) T.Gpointer

	SimpleAsyncResultGetType func() O.Type

	SimpleAsyncResultNew func(
		sourceObject *O.Object,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer,
		sourceTag T.Gpointer) *T.GSimpleAsyncResult

	SimpleAsyncResultNewError func(
		sourceObject *O.Object, callback T.GAsyncReadyCallback,
		userData T.Gpointer, domain T.GQuark, code int,
		format string, v ...VArg) *T.GSimpleAsyncResult

	SimpleAsyncResultNewFromError func(
		sourceObject *O.Object,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer,
		err *T.GError) *T.GSimpleAsyncResult

	SimpleAsyncResultNewTakeError func(
		sourceObject *O.Object,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer,
		err *T.GError) *T.GSimpleAsyncResult

	SimpleAsyncResultSetOpResGpointer func(
		simple *T.GSimpleAsyncResult,
		opRes T.Gpointer,
		destroyOpRes T.GDestroyNotify)

	SimpleAsyncResultGetOpResGpointer func(
		simple *T.GSimpleAsyncResult) T.Gpointer

	SimpleAsyncResultSetOpResGssize func(
		simple *T.GSimpleAsyncResult,
		opRes T.Gssize)

	SimpleAsyncResultGetOpResGssize func(
		simple *T.GSimpleAsyncResult) T.Gssize

	SimpleAsyncResultSetOpResGboolean func(
		simple *T.GSimpleAsyncResult,
		opRes T.Gboolean)

	SimpleAsyncResultGetOpResGboolean func(
		simple *T.GSimpleAsyncResult) T.Gboolean

	SimpleAsyncResultGetSourceTag func(
		simple *T.GSimpleAsyncResult) T.Gpointer

	SimpleAsyncResultSetHandleCancellation func(
		simple *T.GSimpleAsyncResult,
		handleCancellation T.Gboolean)

	SimpleAsyncResultComplete func(
		simple *T.GSimpleAsyncResult)

	SimpleAsyncResultCompleteInIdle func(
		simple *T.GSimpleAsyncResult)

	SimpleAsyncResultRunInThread func(
		simple *T.GSimpleAsyncResult,
		f T.GSimpleAsyncThreadFunc,
		ioPriority int,
		cancellable *T.GCancellable)

	SimpleAsyncResultSetFromError func(
		simple *T.GSimpleAsyncResult,
		err *T.GError)

	SimpleAsyncResultTakeError func(
		simple *T.GSimpleAsyncResult,
		err *T.GError)

	SimpleAsyncResultPropagateError func(
		simple *T.GSimpleAsyncResult,
		dest **T.GError) T.Gboolean

	SimpleAsyncResultSetError func(
		simple *T.GSimpleAsyncResult, domain T.GQuark,
		code int, format string, v ...VArg)

	SimpleAsyncResultSetErrorVa func(
		simple *T.GSimpleAsyncResult,
		domain T.GQuark,
		code int,
		format string,
		args T.VaList)

	SimpleAsyncResultIsValid func(
		result *T.GAsyncResult,
		source *O.Object,
		sourceTag T.Gpointer) T.Gboolean

	SimpleAsyncReportErrorInIdle func(object *O.Object,
		callback T.GAsyncReadyCallback, userData T.Gpointer,
		domain T.GQuark, code int, format string, v ...VArg)

	SimpleAsyncReportGerrorInIdle func(
		object *O.Object,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer,
		err *T.GError)

	SimpleAsyncReportTakeGerrorInIdle func(
		object *O.Object,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer,
		err *T.GError)

	SimplePermissionGetType func() O.Type

	SimplePermissionNew func(
		allowed T.Gboolean) *T.GPermission

	SocketClientGetType func() O.Type

	SocketClientNew func() *T.GSocketClient

	SocketClientGetFamily func(
		client *T.GSocketClient) T.GSocketFamily

	SocketClientSetFamily func(
		client *T.GSocketClient,
		family T.GSocketFamily)

	SocketClientGetSocketType func(
		client *T.GSocketClient) T.GSocketType

	SocketClientSetSocketType func(
		client *T.GSocketClient,
		t T.GSocketType)

	SocketClientGetProtocol func(
		client *T.GSocketClient) T.GSocketProtocol

	SocketClientSetProtocol func(
		client *T.GSocketClient,
		protocol T.GSocketProtocol)

	SocketClientGetLocalAddress func(
		client *T.GSocketClient) *T.GSocketAddress

	SocketClientSetLocalAddress func(
		client *T.GSocketClient,
		address *T.GSocketAddress)

	SocketClientGetTimeout func(
		client *T.GSocketClient) uint

	SocketClientSetTimeout func(
		client *T.GSocketClient,
		timeout uint)

	SocketClientGetEnableProxy func(
		client *T.GSocketClient) T.Gboolean

	SocketClientSetEnableProxy func(
		client *T.GSocketClient,
		enable T.Gboolean)

	SocketClientGetTls func(
		client *T.GSocketClient) T.Gboolean

	SocketClientSetTls func(
		client *T.GSocketClient,
		tls T.Gboolean)

	SocketClientGetTlsValidationFlags func(
		client *T.GSocketClient) T.GTlsCertificateFlags

	SocketClientSetTlsValidationFlags func(
		client *T.GSocketClient,
		flags T.GTlsCertificateFlags)

	SocketClientConnect func(
		client *T.GSocketClient,
		connectable *T.GSocketConnectable,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	SocketClientConnectToHost func(
		client *T.GSocketClient,
		hostAndPort string,
		defaultPort uint16,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	SocketClientConnectToService func(
		client *T.GSocketClient,
		domain string,
		service string,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	SocketClientConnectToUri func(
		client *T.GSocketClient,
		uri string,
		defaultPort uint16,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	SocketClientConnectAsync func(
		client *T.GSocketClient,
		connectable *T.GSocketConnectable,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	SocketClientConnectFinish func(
		client *T.GSocketClient,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketConnection

	SocketClientConnectToHostAsync func(
		client *T.GSocketClient,
		hostAndPort string,
		defaultPort uint16,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	SocketClientConnectToHostFinish func(
		client *T.GSocketClient,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketConnection

	SocketClientConnectToServiceAsync func(
		client *T.GSocketClient,
		domain string,
		service string,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	SocketClientConnectToServiceFinish func(
		client *T.GSocketClient,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketConnection

	SocketClientConnectToUriAsync func(
		client *T.GSocketClient,
		uri string,
		defaultPort uint16,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	SocketClientConnectToUriFinish func(
		client *T.GSocketClient,
		result *T.GAsyncResult,
		err **T.GError) *T.GSocketConnection

	SocketClientAddApplicationProxy func(
		client *T.GSocketClient,
		protocol string)

	SocketConnectableGetType func() O.Type

	SocketConnectableEnumerate func(
		connectable *T.GSocketConnectable) *T.GSocketAddressEnumerator

	SocketConnectableProxyEnumerate func(
		connectable *T.GSocketConnectable) *T.GSocketAddressEnumerator

	SocketGetType func() O.Type

	SocketNew func(
		family T.GSocketFamily,
		t T.GSocketType,
		protocol T.GSocketProtocol,
		err **T.GError) *T.GSocket

	SocketNewFromFd func(
		fd int,
		err **T.GError) *T.GSocket

	SocketGetFd func(
		socket *T.GSocket) int

	SocketGetFamily func(
		socket *T.GSocket) T.GSocketFamily

	SocketGetSocketType func(
		socket *T.GSocket) T.GSocketType

	SocketGetProtocol func(
		socket *T.GSocket) T.GSocketProtocol

	SocketGetLocalAddress func(
		socket *T.GSocket,
		err **T.GError) *T.GSocketAddress

	SocketGetRemoteAddress func(
		socket *T.GSocket,
		err **T.GError) *T.GSocketAddress

	SocketSetBlocking func(
		socket *T.GSocket,
		blocking T.Gboolean)

	SocketGetBlocking func(
		socket *T.GSocket) T.Gboolean

	SocketSet_keepalive func(
		socket *T.GSocket,
		keepalive T.Gboolean)

	SocketGet_keepalive func(
		socket *T.GSocket) T.Gboolean

	SocketGetListenBacklog func(
		socket *T.GSocket) int

	SocketSetListenBacklog func(
		socket *T.GSocket,
		backlog int)

	SocketGetTimeout func(
		socket *T.GSocket) uint

	SocketSetTimeout func(
		socket *T.GSocket,
		timeout uint)

	SocketIsConnected func(
		socket *T.GSocket) T.Gboolean

	SocketBind func(
		socket *T.GSocket,
		address *T.GSocketAddress,
		allowReuse T.Gboolean,
		err **T.GError) T.Gboolean

	SocketConnect func(
		socket *T.GSocket,
		address *T.GSocketAddress,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	SocketCheckConnectResult func(
		socket *T.GSocket,
		err **T.GError) T.Gboolean

	SocketConditionCheck func(
		socket *T.GSocket,
		condition T.GIOCondition) T.GIOCondition

	SocketConditionWait func(
		socket *T.GSocket,
		condition T.GIOCondition,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	SocketAccept func(
		socket *T.GSocket,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocket

	SocketListen func(
		socket *T.GSocket,
		err **T.GError) T.Gboolean

	SocketReceive func(
		socket *T.GSocket,
		buffer string,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	SocketReceiveFrom func(
		socket *T.GSocket,
		address **T.GSocketAddress,
		buffer string,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	SocketSend func(
		socket *T.GSocket,
		buffer string,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	SocketSendTo func(
		socket *T.GSocket,
		address *T.GSocketAddress,
		buffer string,
		size T.Gsize,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	SocketReceiveMessage func(
		socket *T.GSocket,
		address **T.GSocketAddress,
		vectors *T.GInputVector,
		numVectors int,
		messages ***T.GSocketControlMessage,
		numMessages *int,
		flags *int,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	SocketSendMessage func(
		socket *T.GSocket,
		address *T.GSocketAddress,
		vectors *T.GOutputVector,
		numVectors int,
		messages **T.GSocketControlMessage,
		numMessages int,
		flags int,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	SocketClose func(
		socket *T.GSocket,
		err **T.GError) T.Gboolean

	SocketShutdown func(
		socket *T.GSocket,
		shutdownRead T.Gboolean,
		shutdownWrite T.Gboolean,
		err **T.GError) T.Gboolean

	SocketIsClosed func(
		socket *T.GSocket) T.Gboolean

	SocketCreateSource func(
		socket *T.GSocket,
		condition T.GIOCondition,
		cancellable *T.GCancellable) *T.GSource

	SocketSpeaksIpv4 func(
		socket *T.GSocket) T.Gboolean

	SocketGetCredentials func(
		socket *T.GSocket,
		err **T.GError) *T.GCredentials

	SocketReceiveWithBlocking func(
		socket *T.GSocket,
		buffer string,
		size T.Gsize,
		blocking T.Gboolean,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	SocketSendWithBlocking func(
		socket *T.GSocket,
		buffer string,
		size T.Gsize,
		blocking T.Gboolean,
		cancellable *T.GCancellable,
		err **T.GError) T.Gssize

	SocketConnectionGetType func() O.Type

	SocketConnectionGetSocket func(
		connection *T.GSocketConnection) *T.GSocket

	SocketConnectionGetLocalAddress func(
		connection *T.GSocketConnection,
		err **T.GError) *T.GSocketAddress

	SocketConnectionGetRemoteAddress func(
		connection *T.GSocketConnection,
		err **T.GError) *T.GSocketAddress

	SocketConnectionFactoryRegisterType func(

		GType O.Type,
		family T.GSocketFamily,
		t T.GSocketType,
		protocol int)

	SocketConnectionFactoryLookupType func(
		family T.GSocketFamily,
		t T.GSocketType,
		protocolId int) O.Type

	SocketConnectionFactoryCreateConnection func(
		socket *T.GSocket) *T.GSocketConnection

	SocketControlMessageGetType func() O.Type

	SocketControlMessageGetSize func(
		message *T.GSocketControlMessage) T.Gsize

	SocketControlMessageGetLevel func(
		message *T.GSocketControlMessage) int

	SocketControlMessageGetMsgType func(
		message *T.GSocketControlMessage) int

	SocketControlMessageSerialize func(
		message *T.GSocketControlMessage,
		data T.Gpointer)

	SocketControlMessageDeserialize func(
		level int,
		typ int,
		size T.Gsize,
		data T.Gpointer) *T.GSocketControlMessage

	SocketListenerGetType func() O.Type

	SocketListenerNew func() *T.GSocketListener

	SocketListenerSetBacklog func(
		listener *T.GSocketListener,
		listenBacklog int)

	SocketListenerAddSocket func(
		listener *T.GSocketListener,
		socket *T.GSocket,
		sourceObject *O.Object,
		err **T.GError) T.Gboolean

	SocketListenerAddAddress func(
		listener *T.GSocketListener,
		address *T.GSocketAddress,
		t T.GSocketType,
		protocol T.GSocketProtocol,
		sourceObject *O.Object,
		effectiveAddress **T.GSocketAddress,
		err **T.GError) T.Gboolean

	SocketListenerAddInetPort func(
		listener *T.GSocketListener,
		port uint16,
		sourceObject *O.Object,
		err **T.GError) T.Gboolean

	SocketListenerAddAnyInetPort func(
		listener *T.GSocketListener,
		sourceObject *O.Object,
		err **T.GError) uint16

	SocketListenerAcceptSocket func(
		listener *T.GSocketListener,
		sourceObject **O.Object,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocket

	SocketListenerAcceptSocketAsync func(
		listener *T.GSocketListener,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	SocketListenerAcceptSocketFinish func(
		listener *T.GSocketListener,
		result *T.GAsyncResult,
		sourceObject **O.Object,
		err **T.GError) *T.GSocket

	SocketListenerAccept func(
		listener *T.GSocketListener,
		sourceObject **O.Object,
		cancellable *T.GCancellable,
		err **T.GError) *T.GSocketConnection

	SocketListenerAcceptAsync func(
		listener *T.GSocketListener,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	SocketListenerAcceptFinish func(
		listener *T.GSocketListener,
		result *T.GAsyncResult,
		sourceObject **O.Object,
		err **T.GError) *T.GSocketConnection

	SocketListenerClose func(
		listener *T.GSocketListener)

	SocketServiceGetType func() O.Type

	SocketServiceNew func() *T.GSocketService

	SocketServiceStart func(
		service *T.GSocketService)

	SocketServiceStop func(
		service *T.GSocketService)

	SocketServiceIsActive func(
		service *T.GSocketService) T.Gboolean

	SrvTargetGetType func() O.Type

	SrvTargetNew func(
		hostname string,
		port uint16,
		priority uint16,
		weight uint16) *T.GSrvTarget

	SrvTargetCopy func(
		target *T.GSrvTarget) *T.GSrvTarget

	SrvTargetFree func(
		target *T.GSrvTarget)

	SrvTargetGetHostname func(
		target *T.GSrvTarget) string

	SrvTargetGetPort func(
		target *T.GSrvTarget) uint16

	SrvTargetGetPriority func(
		target *T.GSrvTarget) uint16

	SrvTargetGetWeight func(
		target *T.GSrvTarget) uint16

	SrvTargetListSort func(
		targets *T.GList) *T.GList

	TcpConnectionGetType func() O.Type

	TcpConnectionSetGracefulDisconnect func(
		connection *T.GTcpConnection,
		gracefulDisconnect T.Gboolean)

	TcpConnectionGetGracefulDisconnect func(
		connection *T.GTcpConnection) T.Gboolean

	TcpWrapperConnectionGetType func() O.Type

	TcpWrapperConnectionNew func(
		baseIoStream *T.GIOStream,
		socket *T.GSocket) *T.GSocketConnection

	TcpWrapperConnectionGetBaseIoStream func(
		conn *T.GTcpWrapperConnection) *T.GIOStream

	ThemedIconGetType func() O.Type

	ThemedIconNew func(
		iconname string) *T.GIcon

	ThemedIconNewWithDefaultFallbacks func(
		iconname string) *T.GIcon

	ThemedIconNewFromNames func(
		iconnames **T.Char,
		len int) *T.GIcon

	ThemedIconPrependName func(
		icon *T.GThemedIcon,
		iconname string)

	ThemedIconAppendName func(
		icon *T.GThemedIcon,
		iconname string)

	ThemedIconGetNames func(
		icon *T.GThemedIcon) **T.Gchar

	ThreadedSocketServiceGetType func() O.Type

	ThreadedSocketServiceNew func(
		maxThreads int) *T.GSocketService

	TlsBackendGetType func() O.Type

	TlsBackendGetDefault func() *T.GTlsBackend

	TlsBackendSupportsTls func(
		backend *T.GTlsBackend) T.Gboolean

	TlsBackendGetCertificateType func(
		backend *T.GTlsBackend) O.Type

	TlsBackendGetClientConnectionType func(
		backend *T.GTlsBackend) O.Type

	TlsBackendGetServerConnectionType func(
		backend *T.GTlsBackend) O.Type

	TlsCertificateGetType func() O.Type

	TlsCertificateNewFromPem func(
		data string,
		length T.Gssize,
		err **T.GError) *T.GTlsCertificate

	TlsCertificateNewFromFile func(
		file string,
		err **T.GError) *T.GTlsCertificate

	TlsCertificateNewFromFiles func(
		certFile string,
		keyFile string,
		err **T.GError) *T.GTlsCertificate

	TlsCertificateListNewFromFile func(
		file string,
		err **T.GError) *T.GList

	TlsCertificateGetIssuer func(
		cert *T.GTlsCertificate) *T.GTlsCertificate

	TlsCertificateVerify func(
		cert *T.GTlsCertificate,
		identity *T.GSocketConnectable,
		trustedCa *T.GTlsCertificate) T.GTlsCertificateFlags

	TlsConnectionGetType func() O.Type

	TlsConnectionSetUseSystemCertdb func(
		conn *T.GTlsConnection,
		useSystemCertdb T.Gboolean)

	TlsConnectionGetUseSystemCertdb func(
		conn *T.GTlsConnection) T.Gboolean

	TlsConnectionSetCertificate func(
		conn *T.GTlsConnection,
		certificate *T.GTlsCertificate)

	TlsConnectionGetCertificate func(
		conn *T.GTlsConnection) *T.GTlsCertificate

	TlsConnectionGetPeerCertificate func(
		conn *T.GTlsConnection) *T.GTlsCertificate

	TlsConnectionGetPeerCertificateErrors func(
		conn *T.GTlsConnection) T.GTlsCertificateFlags

	TlsConnectionSetRequireCloseNotify func(
		conn *T.GTlsConnection,
		requireCloseNotify T.Gboolean)

	TlsConnectionGetRequireCloseNotify func(
		conn *T.GTlsConnection) T.Gboolean

	TlsConnectionSetRehandshakeMode func(
		conn *T.GTlsConnection,
		mode T.GTlsRehandshakeMode)

	TlsConnectionGetRehandshakeMode func(
		conn *T.GTlsConnection) T.GTlsRehandshakeMode

	TlsConnectionHandshake func(
		conn *T.GTlsConnection,
		cancellable *T.GCancellable,
		err **T.GError) T.Gboolean

	TlsConnectionHandshakeAsync func(
		conn *T.GTlsConnection,
		ioPriority int,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	TlsConnectionHandshakeFinish func(
		conn *T.GTlsConnection,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	TlsErrorQuark func() T.GQuark

	TlsConnectionEmitAcceptCertificate func(
		conn *T.GTlsConnection,
		peerCert *T.GTlsCertificate,
		errors T.GTlsCertificateFlags) T.Gboolean

	TlsClientConnectionGetType func() O.Type

	TlsClientConnectionNew func(
		baseIoStream *T.GIOStream,
		serverIdentity *T.GSocketConnectable,
		err **T.GError) *T.GIOStream

	TlsClientConnectionGetValidationFlags func(
		conn *T.GTlsClientConnection) T.GTlsCertificateFlags

	TlsClientConnectionSetValidationFlags func(
		conn *T.GTlsClientConnection,
		flags T.GTlsCertificateFlags)

	TlsClientConnectionGetServerIdentity func(
		conn *T.GTlsClientConnection) *T.GSocketConnectable

	TlsClientConnectionSetServerIdentity func(
		conn *T.GTlsClientConnection,
		identity *T.GSocketConnectable)

	TlsClientConnectionGetUseSsl3 func(
		conn *T.GTlsClientConnection) T.Gboolean

	TlsClientConnectionSetUseSsl3 func(
		conn *T.GTlsClientConnection,
		useSsl3 T.Gboolean)

	TlsClientConnectionGetAcceptedCas func(
		conn *T.GTlsClientConnection) *T.GList

	TlsServerConnectionGetType func() O.Type

	TlsServerConnectionNew func(
		baseIoStream *T.GIOStream,
		certificate *T.GTlsCertificate,
		err **T.GError) *T.GIOStream

	VfsGetType func() O.Type

	VfsIsActive func(
		vfs *T.GVfs) T.Gboolean

	VfsGetFileForPath func(
		vfs *T.GVfs,
		path string) *T.GFile

	VfsGetFileForUri func(
		vfs *T.GVfs,
		uri string) *T.GFile

	VfsGetSupportedUriSchemes func(
		vfs *T.GVfs) **T.Gchar

	VfsParseName func(
		vfs *T.GVfs,
		parseName string) *T.GFile

	VfsGetDefault func() *T.GVfs

	VfsGetLocal func() *T.GVfs

	VolumeGetType func() O.Type

	VolumeGetName func(
		volume *T.GVolume) string

	VolumeGetIcon func(
		volume *T.GVolume) *T.GIcon

	VolumeGetUuid func(
		volume *T.GVolume) string

	VolumeGetDrive func(
		volume *T.GVolume) *T.GDrive

	VolumeGetMount func(
		volume *T.GVolume) *T.GMount

	VolumeCanMount func(
		volume *T.GVolume) T.Gboolean

	VolumeCanEject func(
		volume *T.GVolume) T.Gboolean

	VolumeShouldAutomount func(
		volume *T.GVolume) T.Gboolean

	VolumeMount func(
		volume *T.GVolume,
		flags T.GMountMountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	VolumeMountFinish func(
		volume *T.GVolume,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	VolumeEject func(
		volume *T.GVolume,
		flags T.GMountUnmountFlags,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	VolumeEjectFinish func(
		volume *T.GVolume,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	VolumeGetIdentifier func(
		volume *T.GVolume,
		kind string) string

	VolumeEnumerateIdentifiers func(
		volume *T.GVolume) **T.Char

	VolumeGetActivationRoot func(
		volume *T.GVolume) *T.GFile

	VolumeEjectWithOperation func(
		volume *T.GVolume,
		flags T.GMountUnmountFlags,
		mountOperation *T.GMountOperation,
		cancellable *T.GCancellable,
		callback T.GAsyncReadyCallback,
		userData T.Gpointer)

	VolumeEjectWithOperationFinish func(
		volume *T.GVolume,
		result *T.GAsyncResult,
		err **T.GError) T.Gboolean

	ZlibCompressorGetType func() O.Type

	ZlibCompressorNew func(
		format T.GZlibCompressorFormat,
		level int) *T.GZlibCompressor

	ZlibCompressorGetFileInfo func(
		compressor *T.GZlibCompressor) *T.GFileInfo

	ZlibCompressorSetFileInfo func(
		compressor *T.GZlibCompressor,
		fileInfo *T.GFileInfo)

	ZlibDecompressorGetType func() O.Type

	ZlibDecompressorNew func(
		format T.GZlibCompressorFormat) *T.GZlibDecompressor

	ZlibDecompressorGetFileInfo func(
		decompressor *T.GZlibDecompressor) *T.GFileInfo

	Win32InputStreamGetType func() O.Type

	Win32InputStreamNew func(
		handle *T.Void,
		closeHandle T.Gboolean) *T.GInputStream

	Win32InputStreamSetCloseHandle func(
		stream *T.GWin32InputStream,
		closeHandle T.Gboolean)

	Win32InputStreamGetCloseHandle func(
		stream *T.GWin32InputStream) T.Gboolean

	Win32InputStreamGetHandle func(
		stream *T.GWin32InputStream) *T.Void

	Win32OutputStreamGetType func() O.Type

	Win32OutputStreamNew func(
		handle *T.Void,
		closeHandle T.Gboolean) *T.GOutputStream

	Win32OutputStreamSetCloseHandle func(
		stream *T.GWin32OutputStream,
		closeHandle T.Gboolean)

	Win32OutputStreamGetCloseHandle func(
		stream *T.GWin32OutputStream) T.Gboolean

	Win32OutputStreamGetHandle func(
		stream *T.GWin32OutputStream) *T.Void

	SettingsBackendGetType func() O.Type

	SettingsBackendChanged func(
		backend *T.GSettingsBackend,
		key string,
		originTag T.Gpointer)

	SettingsBackendPathChanged func(
		backend *T.GSettingsBackend,
		path string,
		originTag T.Gpointer)

	SettingsBackendFlattenTree func(
		tree *T.GTree,
		path **T.Gchar,
		keys ***T.Gchar,
		values ***T.GVariant)

	SettingsBackend_keysChanged func(
		backend *T.GSettingsBackend,
		path string,
		items **T.Gchar,
		originTag T.Gpointer)

	SettingsBackendPathWritableChanged func(
		backend *T.GSettingsBackend,
		path string)

	SettingsBackendWritableChanged func(
		backend *T.GSettingsBackend,
		key string)

	SettingsBackendChangedTree func(
		backend *T.GSettingsBackend,
		tree *T.GTree,
		originTag T.Gpointer)

	SettingsBackendGetDefault func() *T.GSettingsBackend

	KeyfileSettingsBackendNew func(
		filename string,
		rootPath string,
		rootGroup string) *T.GSettingsBackend

	NullSettingsBackendNew func() *T.GSettingsBackend

	MemorySettingsBackendNew func() *T.GSettingsBackend
)
var dll = "libgio-2.0-0.dll"

var apiList = outside.Apis{
	{"g_action_activate", &ActionActivate},
	{"g_action_get_enabled", &ActionGetEnabled},
	{"g_action_get_name", &ActionGetName},
	{"g_action_get_parameter_type", &ActionGetParameterType},
	{"g_action_get_state", &ActionGetState},
	{"g_action_get_state_hint", &ActionGetStateHint},
	{"g_action_get_state_type", &ActionGetStateType},
	{"g_action_get_type", &ActionGetType},
	{"g_action_group_action_added", &ActionGroupActionAdded},
	{"g_action_group_action_enabled_changed", &ActionGroupActionEnabledChanged},
	{"g_action_group_action_removed", &ActionGroupActionRemoved},
	{"g_action_group_action_state_changed", &ActionGroupActionStateChanged},
	{"g_action_group_activate_action", &ActionGroupActivateAction},
	{"g_action_group_change_action_state", &ActionGroupChangeActionState},
	{"g_action_group_get_action_enabled", &ActionGroupGetActionEnabled},
	{"g_action_group_get_action_parameter_type", &ActionGroupGetActionParameterType},
	{"g_action_group_get_action_state", &ActionGroupGetActionState},
	{"g_action_group_get_action_state_hint", &ActionGroupGetActionStateHint},
	{"g_action_group_get_action_state_type", &ActionGroupGetActionStateType},
	{"g_action_group_get_type", &ActionGroupGetType},
	{"g_action_group_has_action", &ActionGroupHasAction},
	{"g_action_group_list_actions", &ActionGroupListActions},
	{"g_action_set_state", &ActionSetState},
	{"g_app_info_add_supports_type", &AppInfoAddSupportsType},
	{"g_app_info_can_delete", &AppInfoCanDelete},
	{"g_app_info_can_remove_supports_type", &AppInfoCanRemoveSupportsType},
	{"g_app_info_create_flags_get_type", &AppInfoCreateFlagsGetType},
	{"g_app_info_create_from_commandline", &AppInfoCreateFromCommandline},
	{"g_app_info_delete", &AppInfoDelete},
	{"g_app_info_dup", &AppInfoDup},
	{"g_app_info_equal", &AppInfoEqual},
	{"g_app_info_get_all", &AppInfoGetAll},
	{"g_app_info_get_all_for_type", &AppInfoGetAllForType},
	{"g_app_info_get_commandline", &AppInfoGetCommandline},
	{"g_app_info_get_default_for_type", &AppInfoGetDefaultForType},
	{"g_app_info_get_default_for_uri_scheme", &AppInfoGetDefaultForUriScheme},
	{"g_app_info_get_description", &AppInfoGetDescription},
	{"g_app_info_get_display_name", &AppInfoGetDisplayName},
	{"g_app_info_get_executable", &AppInfoGetExecutable},
	{"g_app_info_get_fallback_for_type", &AppInfoGetFallbackForType},
	{"g_app_info_get_icon", &AppInfoGetIcon},
	{"g_app_info_get_id", &AppInfoGetId},
	{"g_app_info_get_name", &AppInfoGetName},
	{"g_app_info_get_recommended_for_type", &AppInfoGetRecommendedForType},
	{"g_app_info_get_type", &AppInfoGetType},
	{"g_app_info_launch", &AppInfoLaunch},
	{"g_app_info_launch_default_for_uri", &AppInfoLaunchDefaultForUri},
	{"g_app_info_launch_uris", &AppInfoLaunchUris},
	{"g_app_info_remove_supports_type", &AppInfoRemoveSupportsType},
	{"g_app_info_reset_type_associations", &AppInfoResetTypeAssociations},
	{"g_app_info_set_as_default_for_extension", &AppInfoSetAsDefaultForExtension},
	{"g_app_info_set_as_default_for_type", &AppInfoSetAsDefaultForType},
	{"g_app_info_set_as_last_used_for_type", &AppInfoSetAsLastUsedForType},
	{"g_app_info_should_show", &AppInfoShouldShow},
	{"g_app_info_supports_files", &AppInfoSupportsFiles},
	{"g_app_info_supports_uris", &AppInfoSupportsUris},
	{"g_app_launch_context_get_display", &AppLaunchContextGetDisplay},
	{"g_app_launch_context_get_startup_notify_id", &AppLaunchContextGetStartupNotifyId},
	{"g_app_launch_context_get_type", &AppLaunchContextGetType},
	{"g_app_launch_context_launch_failed", &AppLaunchContextLaunchFailed},
	{"g_app_launch_context_new", &AppLaunchContextNew},
	{"g_application_activate", &ApplicationActivate},
	{"g_application_command_line_get_arguments", &ApplicationCommandLineGetArguments},
	{"g_application_command_line_get_cwd", &ApplicationCommandLineGetCwd},
	{"g_application_command_line_get_environ", &ApplicationCommandLineGetEnviron},
	{"g_application_command_line_get_exit_status", &ApplicationCommandLineGetExitStatus},
	{"g_application_command_line_get_is_remote", &ApplicationCommandLineGetIsRemote},
	{"g_application_command_line_get_platform_data", &ApplicationCommandLineGetPlatformData},
	{"g_application_command_line_get_type", &ApplicationCommandLineGetType},
	{"g_application_command_line_getenv", &ApplicationCommandLineGetenv},
	{"g_application_command_line_print", &ApplicationCommandLinePrint},
	{"g_application_command_line_printerr", &ApplicationCommandLinePrinterr},
	{"g_application_command_line_set_exit_status", &ApplicationCommandLineSetExitStatus},
	{"g_application_flags_get_type", &ApplicationFlagsGetType},
	{"g_application_get_application_id", &ApplicationGetApplicationId},
	{"g_application_get_flags", &ApplicationGetFlags},
	{"g_application_get_inactivity_timeout", &ApplicationGetInactivityTimeout},
	{"g_application_get_is_registered", &ApplicationGetIsRegistered},
	{"g_application_get_is_remote", &ApplicationGetIsRemote},
	{"g_application_get_type", &ApplicationGetType},
	{"g_application_hold", &ApplicationHold},
	{"g_application_id_is_valid", &ApplicationIdIsValid},
	{"g_application_new", &ApplicationNew},
	{"g_application_open", &ApplicationOpen},
	{"g_application_register", &ApplicationRegister},
	{"g_application_release", &ApplicationRelease},
	{"g_application_run", &ApplicationRun},
	{"g_application_set_action_group", &ApplicationSetActionGroup},
	{"g_application_set_application_id", &ApplicationSetApplicationId},
	{"g_application_set_flags", &ApplicationSetFlags},
	{"g_application_set_inactivity_timeout", &ApplicationSetInactivityTimeout},
	{"g_ask_password_flags_get_type", &AskPasswordFlagsGetType},
	{"g_async_initable_get_type", &AsyncInitableGetType},
	{"g_async_initable_init_async", &AsyncInitableInitAsync},
	{"g_async_initable_init_finish", &AsyncInitableInitFinish},
	{"g_async_initable_new_async", &AsyncInitableNewAsync},
	{"g_async_initable_new_finish", &AsyncInitableNewFinish},
	{"g_async_initable_new_valist_async", &AsyncInitableNewValistAsync},
	{"g_async_initable_newv_async", &AsyncInitableNewvAsync},
	{"g_async_result_get_source_object", &AsyncResultGetSourceObject},
	{"g_async_result_get_type", &AsyncResultGetType},
	{"g_async_result_get_user_data", &AsyncResultGetUserData},
	{"g_buffered_input_stream_fill", &BufferedInputStreamFill},
	{"g_buffered_input_stream_fill_async", &BufferedInputStreamFillAsync},
	{"g_buffered_input_stream_fill_finish", &BufferedInputStreamFillFinish},
	{"g_buffered_input_stream_get_available", &BufferedInputStreamGetAvailable},
	{"g_buffered_input_stream_get_buffer_size", &BufferedInputStreamGetBufferSize},
	{"g_buffered_input_stream_get_type", &BufferedInputStreamGetType},
	{"g_buffered_input_stream_new", &BufferedInputStreamNew},
	{"g_buffered_input_stream_new_sized", &BufferedInputStreamNewSized},
	{"g_buffered_input_stream_peek", &BufferedInputStreamPeek},
	{"g_buffered_input_stream_peek_buffer", &BufferedInputStreamPeekBuffer},
	{"g_buffered_input_stream_read_byte", &BufferedInputStreamReadByte},
	{"g_buffered_input_stream_set_buffer_size", &BufferedInputStreamSetBufferSize},
	{"g_buffered_output_stream_get_auto_grow", &BufferedOutputStreamGetAutoGrow},
	{"g_buffered_output_stream_get_buffer_size", &BufferedOutputStreamGetBufferSize},
	{"g_buffered_output_stream_get_type", &BufferedOutputStreamGetType},
	{"g_buffered_output_stream_new", &BufferedOutputStreamNew},
	{"g_buffered_output_stream_new_sized", &BufferedOutputStreamNewSized},
	{"g_buffered_output_stream_set_auto_grow", &BufferedOutputStreamSetAutoGrow},
	{"g_buffered_output_stream_set_buffer_size", &BufferedOutputStreamSetBufferSize},
	{"g_bus_get", &BusGet},
	{"g_bus_get_finish", &BusGetFinish},
	{"g_bus_get_sync", &BusGetSync},
	{"g_bus_name_owner_flags_get_type", &BusNameOwnerFlagsGetType},
	{"g_bus_name_watcher_flags_get_type", &BusNameWatcherFlagsGetType},
	{"g_bus_own_name", &BusOwnName},
	{"g_bus_own_name_on_connection", &BusOwnNameOnConnection},
	{"g_bus_own_name_on_connection_with_closures", &BusOwnNameOnConnectionWithClosures},
	{"g_bus_own_name_with_closures", &BusOwnNameWithClosures},
	{"g_bus_type_get_type", &BusTypeGetType},
	{"g_bus_unown_name", &BusUnownName},
	{"g_bus_unwatch_name", &BusUnwatchName},
	{"g_bus_watch_name", &BusWatchName},
	{"g_bus_watch_name_on_connection", &BusWatchNameOnConnection},
	{"g_bus_watch_name_on_connection_with_closures", &BusWatchNameOnConnectionWithClosures},
	{"g_bus_watch_name_with_closures", &BusWatchNameWithClosures},
	{"g_cancellable_cancel", &CancellableCancel},
	{"g_cancellable_connect", &CancellableConnect},
	{"g_cancellable_disconnect", &CancellableDisconnect},
	{"g_cancellable_get_current", &CancellableGetCurrent},
	{"g_cancellable_get_fd", &CancellableGetFd},
	{"g_cancellable_get_type", &CancellableGetType},
	{"g_cancellable_is_cancelled", &CancellableIsCancelled},
	{"g_cancellable_make_pollfd", &CancellableMakePollfd},
	{"g_cancellable_new", &CancellableNew},
	{"g_cancellable_pop_current", &CancellablePopCurrent},
	{"g_cancellable_push_current", &CancellablePushCurrent},
	{"g_cancellable_release_fd", &CancellableReleaseFd},
	{"g_cancellable_reset", &CancellableReset},
	{"g_cancellable_set_error_if_cancelled", &CancellableSetErrorIfCancelled},
	{"g_cancellable_source_new", &CancellableSourceNew},
	{"g_charset_converter_get_num_fallbacks", &CharsetConverterGetNumFallbacks},
	{"g_charset_converter_get_type", &CharsetConverterGetType},
	{"g_charset_converter_get_use_fallback", &CharsetConverterGetUseFallback},
	{"g_charset_converter_new", &CharsetConverterNew},
	{"g_charset_converter_set_use_fallback", &CharsetConverterSetUseFallback},
	{"g_content_type_can_be_executable", &ContentTypeCanBeExecutable},
	{"g_content_type_equals", &ContentTypeEquals},
	{"g_content_type_from_mime_type", &ContentTypeFromMimeType},
	{"g_content_type_get_description", &ContentTypeGetDescription},
	{"g_content_type_get_icon", &ContentTypeGetIcon},
	{"g_content_type_get_mime_type", &ContentTypeGetMimeType},
	{"g_content_type_guess", &ContentTypeGuess},
	{"g_content_type_guess_for_tree", &ContentTypeGuessForTree},
	{"g_content_type_is_a", &ContentTypeIsA},
	{"g_content_type_is_unknown", &ContentTypeIsUnknown},
	{"g_content_types_get_registered", &ContentTypesGetRegistered},
	{"g_converter_convert", &ConverterConvert},
	{"g_converter_flags_get_type", &ConverterFlagsGetType},
	{"g_converter_get_type", &ConverterGetType},
	{"g_converter_input_stream_get_converter", &ConverterInputStreamGetConverter},
	{"g_converter_input_stream_get_type", &ConverterInputStreamGetType},
	{"g_converter_input_stream_new", &ConverterInputStreamNew},
	{"g_converter_output_stream_get_converter", &ConverterOutputStreamGetConverter},
	{"g_converter_output_stream_get_type", &ConverterOutputStreamGetType},
	{"g_converter_output_stream_new", &ConverterOutputStreamNew},
	{"g_converter_reset", &ConverterReset},
	{"g_converter_result_get_type", &ConverterResultGetType},
	{"g_credentials_get_native", &CredentialsGetNative},
	{"g_credentials_get_type", &CredentialsGetType},
	{"g_credentials_is_same_user", &CredentialsIsSameUser},
	{"g_credentials_new", &CredentialsNew},
	{"g_credentials_set_native", &CredentialsSetNative},
	{"g_credentials_to_string", &CredentialsToString},
	{"g_credentials_type_get_type", &CredentialsTypeGetType},
	{"g_data_input_stream_get_byte_order", &DataInputStreamGetByteOrder},
	{"g_data_input_stream_get_newline_type", &DataInputStreamGetNewlineType},
	{"g_data_input_stream_get_type", &DataInputStreamGetType},
	{"g_data_input_stream_new", &DataInputStreamNew},
	{"g_data_input_stream_read_byte", &DataInputStreamReadByte},
	{"g_data_input_stream_read_int16", &DataInputStreamReadInt16},
	{"g_data_input_stream_read_int32", &DataInputStreamReadInt32},
	{"g_data_input_stream_read_int64", &DataInputStreamReadInt64},
	{"g_data_input_stream_read_line", &DataInputStreamReadLine},
	{"g_data_input_stream_read_line_async", &DataInputStreamReadLineAsync},
	{"g_data_input_stream_read_line_finish", &DataInputStreamReadLineFinish},
	{"g_data_input_stream_read_uint16", &DataInputStreamReadUint16},
	{"g_data_input_stream_read_uint32", &DataInputStreamReadUint32},
	{"g_data_input_stream_read_uint64", &DataInputStreamReadUint64},
	{"g_data_input_stream_read_until", &DataInputStreamReadUntil},
	{"g_data_input_stream_read_until_async", &DataInputStreamReadUntilAsync},
	{"g_data_input_stream_read_until_finish", &DataInputStreamReadUntilFinish},
	{"g_data_input_stream_read_upto", &DataInputStreamReadUpto},
	{"g_data_input_stream_read_upto_async", &DataInputStreamReadUptoAsync},
	{"g_data_input_stream_read_upto_finish", &DataInputStreamReadUptoFinish},
	{"g_data_input_stream_set_byte_order", &DataInputStreamSetByteOrder},
	{"g_data_input_stream_set_newline_type", &DataInputStreamSetNewlineType},
	{"g_data_output_stream_get_byte_order", &DataOutputStreamGetByteOrder},
	{"g_data_output_stream_get_type", &DataOutputStreamGetType},
	{"g_data_output_stream_new", &DataOutputStreamNew},
	{"g_data_output_stream_put_byte", &DataOutputStreamPutByte},
	{"g_data_output_stream_put_int16", &DataOutputStreamPutInt16},
	{"g_data_output_stream_put_int32", &DataOutputStreamPutInt32},
	{"g_data_output_stream_put_int64", &DataOutputStreamPutInt64},
	{"g_data_output_stream_put_string", &DataOutputStreamPutString},
	{"g_data_output_stream_put_uint16", &DataOutputStreamPutUint16},
	{"g_data_output_stream_put_uint32", &DataOutputStreamPutUint32},
	{"g_data_output_stream_put_uint64", &DataOutputStreamPutUint64},
	{"g_data_output_stream_set_byte_order", &DataOutputStreamSetByteOrder},
	{"g_data_stream_byte_order_get_type", &DataStreamByteOrderGetType},
	{"g_data_stream_newline_type_get_type", &DataStreamNewlineTypeGetType},
	{"g_dbus_address_get_for_bus_sync", &DbusAddressGetForBusSync},
	{"g_dbus_address_get_stream", &DbusAddressGetStream},
	{"g_dbus_address_get_stream_finish", &DbusAddressGetStreamFinish},
	{"g_dbus_address_get_stream_sync", &DbusAddressGetStreamSync},
	{"g_dbus_annotation_info_get_type", &DbusAnnotationInfoGetType},
	{"g_dbus_annotation_info_lookup", &DbusAnnotationInfoLookup},
	{"g_dbus_annotation_info_ref", &DbusAnnotationInfoRef},
	{"g_dbus_annotation_info_unref", &DbusAnnotationInfoUnref},
	{"g_dbus_arg_info_get_type", &DbusArgInfoGetType},
	{"g_dbus_arg_info_ref", &DbusArgInfoRef},
	{"g_dbus_arg_info_unref", &DbusArgInfoUnref},
	{"g_dbus_auth_observer_authorize_authenticated_peer", &DbusAuthObserverAuthorizeAuthenticatedPeer},
	{"g_dbus_auth_observer_get_type", &DbusAuthObserverGetType},
	{"g_dbus_auth_observer_new", &DbusAuthObserverNew},
	{"g_dbus_call_flags_get_type", &DbusCallFlagsGetType},
	{"g_dbus_capability_flags_get_type", &DbusCapabilityFlagsGetType},
	{"g_dbus_connection_add_filter", &DbusConnectionAddFilter},
	{"g_dbus_connection_call", &DbusConnectionCall},
	{"g_dbus_connection_call_finish", &DbusConnectionCallFinish},
	{"g_dbus_connection_call_sync", &DbusConnectionCallSync},
	{"g_dbus_connection_close", &DbusConnectionClose},
	{"g_dbus_connection_close_finish", &DbusConnectionCloseFinish},
	{"g_dbus_connection_close_sync", &DbusConnectionCloseSync},
	{"g_dbus_connection_emit_signal", &DbusConnectionEmitSignal},
	{"g_dbus_connection_flags_get_type", &DbusConnectionFlagsGetType},
	{"g_dbus_connection_flush", &DbusConnectionFlush},
	{"g_dbus_connection_flush_finish", &DbusConnectionFlushFinish},
	{"g_dbus_connection_flush_sync", &DbusConnectionFlushSync},
	{"g_dbus_connection_get_capabilities", &DbusConnectionGetCapabilities},
	{"g_dbus_connection_get_exit_on_close", &DbusConnectionGetExitOnClose},
	{"g_dbus_connection_get_guid", &DbusConnectionGetGuid},
	{"g_dbus_connection_get_peer_credentials", &DbusConnectionGetPeerCredentials},
	{"g_dbus_connection_get_stream", &DbusConnectionGetStream},
	{"g_dbus_connection_get_type", &DbusConnectionGetType},
	{"g_dbus_connection_get_unique_name", &DbusConnectionGetUniqueName},
	{"g_dbus_connection_is_closed", &DbusConnectionIsClosed},
	{"g_dbus_connection_new", &DbusConnectionNew},
	{"g_dbus_connection_new_finish", &DbusConnectionNewFinish},
	{"g_dbus_connection_new_for_address", &DbusConnectionNewForAddress},
	{"g_dbus_connection_new_for_address_finish", &DbusConnectionNewForAddressFinish},
	{"g_dbus_connection_new_for_address_sync", &DbusConnectionNewForAddressSync},
	{"g_dbus_connection_new_sync", &DbusConnectionNewSync},
	{"g_dbus_connection_register_object", &DbusConnectionRegisterObject},
	{"g_dbus_connection_register_subtree", &DbusConnectionRegisterSubtree},
	{"g_dbus_connection_remove_filter", &DbusConnectionRemoveFilter},
	{"g_dbus_connection_send_message", &DbusConnectionSendMessage},
	{"g_dbus_connection_send_message_with_reply", &DbusConnectionSendMessageWithReply},
	{"g_dbus_connection_send_message_with_reply_finish", &DbusConnectionSendMessageWithReplyFinish},
	{"g_dbus_connection_send_message_with_reply_sync", &DbusConnectionSendMessageWithReplySync},
	{"g_dbus_connection_set_exit_on_close", &DbusConnectionSetExitOnClose},
	{"g_dbus_connection_signal_subscribe", &DbusConnectionSignalSubscribe},
	{"g_dbus_connection_signal_unsubscribe", &DbusConnectionSignalUnsubscribe},
	{"g_dbus_connection_start_message_processing", &DbusConnectionStartMessageProcessing},
	{"g_dbus_connection_unregister_object", &DbusConnectionUnregisterObject},
	{"g_dbus_connection_unregister_subtree", &DbusConnectionUnregisterSubtree},
	{"g_dbus_error_encode_gerror", &DbusErrorEncodeGerror},
	{"g_dbus_error_get_remote_error", &DbusErrorGetRemoteError},
	{"g_dbus_error_get_type", &DbusErrorGetType},
	{"g_dbus_error_is_remote_error", &DbusErrorIsRemoteError},
	{"g_dbus_error_new_for_dbus_error", &DbusErrorNewForDbusError},
	{"g_dbus_error_quark", &DbusErrorQuark},
	{"g_dbus_error_register_error", &DbusErrorRegisterError},
	{"g_dbus_error_register_error_domain", &DbusErrorRegisterErrorDomain},
	{"g_dbus_error_set_dbus_error", &DbusErrorSetDbusError},
	{"g_dbus_error_set_dbus_error_valist", &DbusErrorSetDbusErrorValist},
	{"g_dbus_error_strip_remote_error", &DbusErrorStripRemoteError},
	{"g_dbus_error_unregister_error", &DbusErrorUnregisterError},
	{"g_dbus_generate_guid", &DbusGenerateGuid},
	{"g_dbus_interface_info_generate_xml", &DbusInterfaceInfoGenerateXml},
	{"g_dbus_interface_info_get_type", &DbusInterfaceInfoGetType},
	{"g_dbus_interface_info_lookup_method", &DbusInterfaceInfoLookupMethod},
	{"g_dbus_interface_info_lookup_property", &DbusInterfaceInfoLookupProperty},
	{"g_dbus_interface_info_lookup_signal", &DbusInterfaceInfoLookupSignal},
	{"g_dbus_interface_info_ref", &DbusInterfaceInfoRef},
	{"g_dbus_interface_info_unref", &DbusInterfaceInfoUnref},
	{"g_dbus_is_address", &DbusIsAddress},
	{"g_dbus_is_guid", &DbusIsGuid},
	{"g_dbus_is_interface_name", &DbusIsInterfaceName},
	{"g_dbus_is_member_name", &DbusIsMemberName},
	{"g_dbus_is_name", &DbusIsName},
	{"g_dbus_is_supported_address", &DbusIsSupportedAddress},
	{"g_dbus_is_unique_name", &DbusIsUniqueName},
	{"g_dbus_message_byte_order_get_type", &DbusMessageByteOrderGetType},
	{"g_dbus_message_bytes_needed", &DbusMessageBytesNeeded},
	{"g_dbus_message_copy", &DbusMessageCopy},
	{"g_dbus_message_flags_get_type", &DbusMessageFlagsGetType},
	{"g_dbus_message_get_arg0", &DbusMessageGetArg0},
	{"g_dbus_message_get_body", &DbusMessageGetBody},
	{"g_dbus_message_get_byte_order", &DbusMessageGetByteOrder},
	{"g_dbus_message_get_destination", &DbusMessageGetDestination},
	{"g_dbus_message_get_error_name", &DbusMessageGetErrorName},
	{"g_dbus_message_get_flags", &DbusMessageGetFlags},
	{"g_dbus_message_get_header", &DbusMessageGetHeader},
	{"g_dbus_message_get_header_fields", &DbusMessageGetHeaderFields},
	{"g_dbus_message_get_interface", &DbusMessageGetInterface},
	{"g_dbus_message_get_locked", &DbusMessageGetLocked},
	{"g_dbus_message_get_member", &DbusMessageGetMember},
	{"g_dbus_message_get_message_type", &DbusMessageGetMessageType},
	{"g_dbus_message_get_num_unix_fds", &DbusMessageGetNumUnixFds},
	{"g_dbus_message_get_path", &DbusMessageGetPath},
	{"g_dbus_message_get_reply_serial", &DbusMessageGetReplySerial},
	{"g_dbus_message_get_sender", &DbusMessageGetSender},
	{"g_dbus_message_get_serial", &DbusMessageGetSerial},
	{"g_dbus_message_get_signature", &DbusMessageGetSignature},
	{"g_dbus_message_get_type", &DbusMessageGetType},
	{"g_dbus_message_header_field_get_type", &DbusMessageHeaderFieldGetType},
	{"g_dbus_message_lock", &DbusMessageLock},
	{"g_dbus_message_new", &DbusMessageNew},
	{"g_dbus_message_new_from_blob", &DbusMessageNewFromBlob},
	{"g_dbus_message_new_method_call", &DbusMessageNewMethodCall},
	{"g_dbus_message_new_method_error", &DbusMessageNewMethodError},
	{"g_dbus_message_new_method_error_literal", &DbusMessageNewMethodErrorLiteral},
	{"g_dbus_message_new_method_error_valist", &DbusMessageNewMethodErrorValist},
	{"g_dbus_message_new_method_reply", &DbusMessageNewMethodReply},
	{"g_dbus_message_new_signal", &DbusMessageNewSignal},
	{"g_dbus_message_print", &DbusMessagePrint},
	{"g_dbus_message_set_body", &DbusMessageSetBody},
	{"g_dbus_message_set_byte_order", &DbusMessageSetByteOrder},
	{"g_dbus_message_set_destination", &DbusMessageSetDestination},
	{"g_dbus_message_set_error_name", &DbusMessageSetErrorName},
	{"g_dbus_message_set_flags", &DbusMessageSetFlags},
	{"g_dbus_message_set_header", &DbusMessageSetHeader},
	{"g_dbus_message_set_interface", &DbusMessageSetInterface},
	{"g_dbus_message_set_member", &DbusMessageSetMember},
	{"g_dbus_message_set_message_type", &DbusMessageSetMessageType},
	{"g_dbus_message_set_num_unix_fds", &DbusMessageSetNumUnixFds},
	{"g_dbus_message_set_path", &DbusMessageSetPath},
	{"g_dbus_message_set_reply_serial", &DbusMessageSetReplySerial},
	{"g_dbus_message_set_sender", &DbusMessageSetSender},
	{"g_dbus_message_set_serial", &DbusMessageSetSerial},
	{"g_dbus_message_set_signature", &DbusMessageSetSignature},
	{"g_dbus_message_to_blob", &DbusMessageToBlob},
	{"g_dbus_message_to_gerror", &DbusMessageToGerror},
	{"g_dbus_message_type_get_type", &DbusMessageTypeGetType},
	{"g_dbus_method_info_get_type", &DbusMethodInfoGetType},
	{"g_dbus_method_info_ref", &DbusMethodInfoRef},
	{"g_dbus_method_info_unref", &DbusMethodInfoUnref},
	{"g_dbus_method_invocation_get_connection", &DbusMethodInvocationGetConnection},
	{"g_dbus_method_invocation_get_interface_name", &DbusMethodInvocationGetInterfaceName},
	{"g_dbus_method_invocation_get_message", &DbusMethodInvocationGetMessage},
	{"g_dbus_method_invocation_get_method_info", &DbusMethodInvocationGetMethodInfo},
	{"g_dbus_method_invocation_get_method_name", &DbusMethodInvocationGetMethodName},
	{"g_dbus_method_invocation_get_object_path", &DbusMethodInvocationGetObjectPath},
	{"g_dbus_method_invocation_get_parameters", &DbusMethodInvocationGetParameters},
	{"g_dbus_method_invocation_get_sender", &DbusMethodInvocationGetSender},
	{"g_dbus_method_invocation_get_type", &DbusMethodInvocationGetType},
	{"g_dbus_method_invocation_get_user_data", &DbusMethodInvocationGetUserData},
	{"g_dbus_method_invocation_return_dbus_error", &DbusMethodInvocationReturnDbusError},
	{"g_dbus_method_invocation_return_error", &DbusMethodInvocationReturnError},
	{"g_dbus_method_invocation_return_error_literal", &DbusMethodInvocationReturnErrorLiteral},
	{"g_dbus_method_invocation_return_error_valist", &DbusMethodInvocationReturnErrorValist},
	{"g_dbus_method_invocation_return_gerror", &DbusMethodInvocationReturnGerror},
	{"g_dbus_method_invocation_return_value", &DbusMethodInvocationReturnValue},
	{"g_dbus_node_info_generate_xml", &DbusNodeInfoGenerateXml},
	{"g_dbus_node_info_get_type", &DbusNodeInfoGetType},
	{"g_dbus_node_info_lookup_interface", &DbusNodeInfoLookupInterface},
	{"g_dbus_node_info_new_for_xml", &DbusNodeInfoNewForXml},
	{"g_dbus_node_info_ref", &DbusNodeInfoRef},
	{"g_dbus_node_info_unref", &DbusNodeInfoUnref},
	{"g_dbus_property_info_flags_get_type", &DbusPropertyInfoFlagsGetType},
	{"g_dbus_property_info_get_type", &DbusPropertyInfoGetType},
	{"g_dbus_property_info_ref", &DbusPropertyInfoRef},
	{"g_dbus_property_info_unref", &DbusPropertyInfoUnref},
	{"g_dbus_proxy_call", &DbusProxyCall},
	{"g_dbus_proxy_call_finish", &DbusProxyCallFinish},
	{"g_dbus_proxy_call_sync", &DbusProxyCallSync},
	{"g_dbus_proxy_flags_get_type", &DbusProxyFlagsGetType},
	{"g_dbus_proxy_get_cached_property", &DbusProxyGetCachedProperty},
	{"g_dbus_proxy_get_cached_property_names", &DbusProxyGetCachedPropertyNames},
	{"g_dbus_proxy_get_connection", &DbusProxyGetConnection},
	{"g_dbus_proxy_get_default_timeout", &DbusProxyGetDefaultTimeout},
	{"g_dbus_proxy_get_flags", &DbusProxyGetFlags},
	{"g_dbus_proxy_get_interface_info", &DbusProxyGetInterfaceInfo},
	{"g_dbus_proxy_get_interface_name", &DbusProxyGetInterfaceName},
	{"g_dbus_proxy_get_name", &DbusProxyGetName},
	{"g_dbus_proxy_get_name_owner", &DbusProxyGetNameOwner},
	{"g_dbus_proxy_get_object_path", &DbusProxyGetObjectPath},
	{"g_dbus_proxy_get_type", &DbusProxyGetType},
	{"g_dbus_proxy_new", &DbusProxyNew},
	{"g_dbus_proxy_new_finish", &DbusProxyNewFinish},
	{"g_dbus_proxy_new_for_bus", &DbusProxyNewForBus},
	{"g_dbus_proxy_new_for_bus_finish", &DbusProxyNewForBusFinish},
	{"g_dbus_proxy_new_for_bus_sync", &DbusProxyNewForBusSync},
	{"g_dbus_proxy_new_sync", &DbusProxyNewSync},
	{"g_dbus_proxy_set_cached_property", &DbusProxySetCachedProperty},
	{"g_dbus_proxy_set_default_timeout", &DbusProxySetDefaultTimeout},
	{"g_dbus_proxy_set_interface_info", &DbusProxySetInterfaceInfo},
	{"g_dbus_send_message_flags_get_type", &DbusSendMessageFlagsGetType},
	{"g_dbus_server_flags_get_type", &DbusServerFlagsGetType},
	{"g_dbus_server_get_client_address", &DbusServerGetClientAddress},
	{"g_dbus_server_get_flags", &DbusServerGetFlags},
	{"g_dbus_server_get_guid", &DbusServerGetGuid},
	{"g_dbus_server_get_type", &DbusServerGetType},
	{"g_dbus_server_is_active", &DbusServerIsActive},
	{"g_dbus_server_new_sync", &DbusServerNewSync},
	{"g_dbus_server_start", &DbusServerStart},
	{"g_dbus_server_stop", &DbusServerStop},
	{"g_dbus_signal_flags_get_type", &DbusSignalFlagsGetType},
	{"g_dbus_signal_info_get_type", &DbusSignalInfoGetType},
	{"g_dbus_signal_info_ref", &DbusSignalInfoRef},
	{"g_dbus_signal_info_unref", &DbusSignalInfoUnref},
	{"g_dbus_subtree_flags_get_type", &DbusSubtreeFlagsGetType},
	{"g_drive_can_eject", &DriveCanEject},
	{"g_drive_can_poll_for_media", &DriveCanPollForMedia},
	{"g_drive_can_start", &DriveCanStart},
	{"g_drive_can_start_degraded", &DriveCanStartDegraded},
	{"g_drive_can_stop", &DriveCanStop},
	{"g_drive_eject", &DriveEject},
	{"g_drive_eject_finish", &DriveEjectFinish},
	{"g_drive_eject_with_operation", &DriveEjectWithOperation},
	{"g_drive_eject_with_operation_finish", &DriveEjectWithOperationFinish},
	{"g_drive_enumerate_identifiers", &DriveEnumerateIdentifiers},
	{"g_drive_get_icon", &DriveGetIcon},
	{"g_drive_get_identifier", &DriveGetIdentifier},
	{"g_drive_get_name", &DriveGetName},
	{"g_drive_get_start_stop_type", &DriveGetStartStopType},
	{"g_drive_get_type", &DriveGetType},
	{"g_drive_get_volumes", &DriveGetVolumes},
	{"g_drive_has_media", &DriveHasMedia},
	{"g_drive_has_volumes", &DriveHasVolumes},
	{"g_drive_is_media_check_automatic", &DriveIsMediaCheckAutomatic},
	{"g_drive_is_media_removable", &DriveIsMediaRemovable},
	{"g_drive_poll_for_media", &DrivePollForMedia},
	{"g_drive_poll_for_media_finish", &DrivePollForMediaFinish},
	{"g_drive_start", &DriveStart},
	{"g_drive_start_finish", &DriveStartFinish},
	{"g_drive_start_flags_get_type", &DriveStartFlagsGetType},
	{"g_drive_start_stop_type_get_type", &DriveStartStopTypeGetType},
	{"g_drive_stop", &DriveStop},
	{"g_drive_stop_finish", &DriveStopFinish},
	{"g_emblem_get_icon", &EmblemGetIcon},
	{"g_emblem_get_origin", &EmblemGetOrigin},
	{"g_emblem_get_type", &EmblemGetType},
	{"g_emblem_new", &EmblemNew},
	{"g_emblem_new_with_origin", &EmblemNewWithOrigin},
	{"g_emblem_origin_get_type", &EmblemOriginGetType},
	{"g_emblemed_icon_add_emblem", &EmblemedIconAddEmblem},
	{"g_emblemed_icon_clear_emblems", &EmblemedIconClearEmblems},
	{"g_emblemed_icon_get_emblems", &EmblemedIconGetEmblems},
	{"g_emblemed_icon_get_icon", &EmblemedIconGetIcon},
	{"g_emblemed_icon_get_type", &EmblemedIconGetType},
	{"g_emblemed_icon_new", &EmblemedIconNew},
	{"g_file_append_to", &FileAppendTo},
	{"g_file_append_to_async", &FileAppendToAsync},
	{"g_file_append_to_finish", &FileAppendToFinish},
	{"g_file_attribute_info_flags_get_type", &FileAttributeInfoFlagsGetType},
	{"g_file_attribute_info_list_add", &FileAttributeInfoListAdd},
	{"g_file_attribute_info_list_dup", &FileAttributeInfoListDup},
	{"g_file_attribute_info_list_get_type", &FileAttributeInfoListGetType},
	{"g_file_attribute_info_list_lookup", &FileAttributeInfoListLookup},
	{"g_file_attribute_info_list_new", &FileAttributeInfoListNew},
	{"g_file_attribute_info_list_ref", &FileAttributeInfoListRef},
	{"g_file_attribute_info_list_unref", &FileAttributeInfoListUnref},
	{"g_file_attribute_matcher_enumerate_namespace", &FileAttributeMatcherEnumerateNamespace},
	{"g_file_attribute_matcher_enumerate_next", &FileAttributeMatcherEnumerateNext},
	{"g_file_attribute_matcher_get_type", &FileAttributeMatcherGetType},
	{"g_file_attribute_matcher_matches", &FileAttributeMatcherMatches},
	{"g_file_attribute_matcher_matches_only", &FileAttributeMatcherMatchesOnly},
	{"g_file_attribute_matcher_new", &FileAttributeMatcherNew},
	{"g_file_attribute_matcher_ref", &FileAttributeMatcherRef},
	{"g_file_attribute_matcher_unref", &FileAttributeMatcherUnref},
	{"g_file_attribute_status_get_type", &FileAttributeStatusGetType},
	{"g_file_attribute_type_get_type", &FileAttributeTypeGetType},
	{"g_file_copy", &FileCopy},
	{"g_file_copy_async", &FileCopyAsync},
	{"g_file_copy_attributes", &FileCopyAttributes},
	{"g_file_copy_finish", &FileCopyFinish},
	{"g_file_copy_flags_get_type", &FileCopyFlagsGetType},
	{"g_file_create", &FileCreate},
	{"g_file_create_async", &FileCreateAsync},
	{"g_file_create_finish", &FileCreateFinish},
	{"g_file_create_flags_get_type", &FileCreateFlagsGetType},
	{"g_file_create_readwrite", &FileCreateReadwrite},
	{"g_file_create_readwrite_async", &FileCreateReadwriteAsync},
	{"g_file_create_readwrite_finish", &FileCreateReadwriteFinish},
	{"g_file_delete", &FileDelete},
	{"g_file_dup", &FileDup},
	{"g_file_eject_mountable", &FileEjectMountable},
	{"g_file_eject_mountable_finish", &FileEjectMountableFinish},
	{"g_file_eject_mountable_with_operation", &FileEjectMountableWithOperation},
	{"g_file_eject_mountable_with_operation_finish", &FileEjectMountableWithOperationFinish},
	{"g_file_enumerate_children", &FileEnumerateChildren},
	{"g_file_enumerate_children_async", &FileEnumerateChildrenAsync},
	{"g_file_enumerate_children_finish", &FileEnumerateChildrenFinish},
	{"g_file_enumerator_close", &FileEnumeratorClose},
	{"g_file_enumerator_close_async", &FileEnumeratorCloseAsync},
	{"g_file_enumerator_close_finish", &FileEnumeratorCloseFinish},
	{"g_file_enumerator_get_container", &FileEnumeratorGetContainer},
	{"g_file_enumerator_get_type", &FileEnumeratorGetType},
	{"g_file_enumerator_has_pending", &FileEnumeratorHasPending},
	{"g_file_enumerator_is_closed", &FileEnumeratorIsClosed},
	{"g_file_enumerator_next_file", &FileEnumeratorNextFile},
	{"g_file_enumerator_next_files_async", &FileEnumeratorNextFilesAsync},
	{"g_file_enumerator_next_files_finish", &FileEnumeratorNextFilesFinish},
	{"g_file_enumerator_set_pending", &FileEnumeratorSetPending},
	{"g_file_equal", &FileEqual},
	{"g_file_find_enclosing_mount", &FileFindEnclosingMount},
	{"g_file_find_enclosing_mount_async", &FileFindEnclosingMountAsync},
	{"g_file_find_enclosing_mount_finish", &FileFindEnclosingMountFinish},
	{"g_file_get_basename", &FileGetBasename},
	{"g_file_get_child", &FileGetChild},
	{"g_file_get_child_for_display_name", &FileGetChildForDisplayName},
	{"g_file_get_parent", &FileGetParent},
	{"g_file_get_parse_name", &FileGetParseName},
	{"g_file_get_path", &FileGetPath},
	{"g_file_get_relative_path", &FileGetRelativePath},
	{"g_file_get_type", &FileGetType},
	{"g_file_get_uri", &FileGetUri},
	{"g_file_get_uri_scheme", &FileGetUriScheme},
	{"g_file_has_parent", &FileHasParent},
	{"g_file_has_prefix", &FileHasPrefix},
	{"g_file_has_uri_scheme", &FileHasUriScheme},
	{"g_file_hash", &FileHash},
	{"g_file_icon_get_file", &FileIconGetFile},
	{"g_file_icon_get_type", &FileIconGetType},
	{"g_file_icon_new", &FileIconNew},
	{"g_file_info_clear_status", &FileInfoClearStatus},
	{"g_file_info_copy_into", &FileInfoCopyInto},
	{"g_file_info_dup", &FileInfoDup},
	{"g_file_info_get_attribute_as_string", &FileInfoGetAttributeAsString},
	{"g_file_info_get_attribute_boolean", &FileInfoGetAttributeBoolean},
	{"g_file_info_get_attribute_byte_string", &FileInfoGetAttributeByteString},
	{"g_file_info_get_attribute_data", &FileInfoGetAttributeData},
	{"g_file_info_get_attribute_int32", &FileInfoGetAttributeInt32},
	{"g_file_info_get_attribute_int64", &FileInfoGetAttributeInt64},
	{"g_file_info_get_attribute_object", &FileInfoGetAttributeObject},
	{"g_file_info_get_attribute_status", &FileInfoGetAttributeStatus},
	{"g_file_info_get_attribute_string", &FileInfoGetAttributeString},
	{"g_file_info_get_attribute_stringv", &FileInfoGetAttributeStringv},
	{"g_file_info_get_attribute_type", &FileInfoGetAttributeType},
	{"g_file_info_get_attribute_uint32", &FileInfoGetAttributeUint32},
	{"g_file_info_get_attribute_uint64", &FileInfoGetAttributeUint64},
	{"g_file_info_get_content_type", &FileInfoGetContentType},
	{"g_file_info_get_display_name", &FileInfoGetDisplayName},
	{"g_file_info_get_edit_name", &FileInfoGetEditName},
	{"g_file_info_get_etag", &FileInfoGetEtag},
	{"g_file_info_get_file_type", &FileInfoGetFileType},
	{"g_file_info_get_icon", &FileInfoGetIcon},
	{"g_file_info_get_is_backup", &FileInfoGetIsBackup},
	{"g_file_info_get_is_hidden", &FileInfoGetIsHidden},
	{"g_file_info_get_is_symlink", &FileInfoGetIsSymlink},
	{"g_file_info_get_modification_time", &FileInfoGetModificationTime},
	{"g_file_info_get_name", &FileInfoGetName},
	{"g_file_info_get_size", &FileInfoGetSize},
	{"g_file_info_get_sort_order", &FileInfoGetSortOrder},
	{"g_file_info_get_symlink_target", &FileInfoGetSymlinkTarget},
	{"g_file_info_get_type", &FileInfoGetType},
	{"g_file_info_has_attribute", &FileInfoHasAttribute},
	{"g_file_info_has_namespace", &FileInfoHasNamespace},
	{"g_file_info_list_attributes", &FileInfoListAttributes},
	{"g_file_info_new", &FileInfoNew},
	{"g_file_info_remove_attribute", &FileInfoRemoveAttribute},
	{"g_file_info_set_attribute", &FileInfoSetAttribute},
	{"g_file_info_set_attribute_boolean", &FileInfoSetAttributeBoolean},
	{"g_file_info_set_attribute_byte_string", &FileInfoSetAttributeByteString},
	{"g_file_info_set_attribute_int32", &FileInfoSetAttributeInt32},
	{"g_file_info_set_attribute_int64", &FileInfoSetAttributeInt64},
	{"g_file_info_set_attribute_mask", &FileInfoSetAttributeMask},
	{"g_file_info_set_attribute_object", &FileInfoSetAttributeObject},
	{"g_file_info_set_attribute_status", &FileInfoSetAttributeStatus},
	{"g_file_info_set_attribute_string", &FileInfoSetAttributeString},
	{"g_file_info_set_attribute_stringv", &FileInfoSetAttributeStringv},
	{"g_file_info_set_attribute_uint32", &FileInfoSetAttributeUint32},
	{"g_file_info_set_attribute_uint64", &FileInfoSetAttributeUint64},
	{"g_file_info_set_content_type", &FileInfoSetContentType},
	{"g_file_info_set_display_name", &FileInfoSetDisplayName},
	{"g_file_info_set_edit_name", &FileInfoSetEditName},
	{"g_file_info_set_file_type", &FileInfoSetFileType},
	{"g_file_info_set_icon", &FileInfoSetIcon},
	{"g_file_info_set_is_hidden", &FileInfoSetIsHidden},
	{"g_file_info_set_is_symlink", &FileInfoSetIsSymlink},
	{"g_file_info_set_modification_time", &FileInfoSetModificationTime},
	{"g_file_info_set_name", &FileInfoSetName},
	{"g_file_info_set_size", &FileInfoSetSize},
	{"g_file_info_set_sort_order", &FileInfoSetSortOrder},
	{"g_file_info_set_symlink_target", &FileInfoSetSymlinkTarget},
	{"g_file_info_unset_attribute_mask", &FileInfoUnsetAttributeMask},
	{"g_file_input_stream_get_type", &FileInputStreamGetType},
	{"g_file_input_stream_query_info", &FileInputStreamQueryInfo},
	{"g_file_input_stream_query_info_async", &FileInputStreamQueryInfoAsync},
	{"g_file_input_stream_query_info_finish", &FileInputStreamQueryInfoFinish},
	{"g_file_io_stream_get_etag", &FileIoStreamGetEtag},
	{"g_file_io_stream_get_type", &FileIoStreamGetType},
	{"g_file_io_stream_query_info", &FileIoStreamQueryInfo},
	{"g_file_io_stream_query_info_async", &FileIoStreamQueryInfoAsync},
	{"g_file_io_stream_query_info_finish", &FileIoStreamQueryInfoFinish},
	{"g_file_is_native", &FileIsNative},
	{"g_file_load_contents", &FileLoadContents},
	{"g_file_load_contents_async", &FileLoadContentsAsync},
	{"g_file_load_contents_finish", &FileLoadContentsFinish},
	{"g_file_load_partial_contents_async", &FileLoadPartialContentsAsync},
	{"g_file_load_partial_contents_finish", &FileLoadPartialContentsFinish},
	{"g_file_make_directory", &FileMakeDirectory},
	{"g_file_make_directory_with_parents", &FileMakeDirectoryWithParents},
	{"g_file_make_symbolic_link", &FileMakeSymbolicLink},
	{"g_file_monitor", &FileMonitor},
	{"g_file_monitor_cancel", &FileMonitorCancel},
	{"g_file_monitor_directory", &FileMonitorDirectory},
	{"g_file_monitor_emit_event", &FileMonitorEmitEvent},
	{"g_file_monitor_event_get_type", &FileMonitorEventGetType},
	{"g_file_monitor_file", &FileMonitorFile},
	{"g_file_monitor_flags_get_type", &FileMonitorFlagsGetType},
	{"g_file_monitor_get_type", &FileMonitorGetType},
	{"g_file_monitor_is_cancelled", &FileMonitorIsCancelled},
	{"g_file_monitor_set_rate_limit", &FileMonitorSetRateLimit},
	{"g_file_mount_enclosing_volume", &FileMountEnclosingVolume},
	{"g_file_mount_enclosing_volume_finish", &FileMountEnclosingVolumeFinish},
	{"g_file_mount_mountable", &FileMountMountable},
	{"g_file_mount_mountable_finish", &FileMountMountableFinish},
	{"g_file_move", &FileMove},
	{"g_file_new_for_commandline_arg", &FileNewForCommandlineArg},
	{"g_file_new_for_path", &FileNewForPath},
	{"g_file_new_for_uri", &FileNewForUri},
	{"g_file_open_readwrite", &FileOpenReadwrite},
	{"g_file_open_readwrite_async", &FileOpenReadwriteAsync},
	{"g_file_open_readwrite_finish", &FileOpenReadwriteFinish},
	{"g_file_output_stream_get_etag", &FileOutputStreamGetEtag},
	{"g_file_output_stream_get_type", &FileOutputStreamGetType},
	{"g_file_output_stream_query_info", &FileOutputStreamQueryInfo},
	{"g_file_output_stream_query_info_async", &FileOutputStreamQueryInfoAsync},
	{"g_file_output_stream_query_info_finish", &FileOutputStreamQueryInfoFinish},
	{"g_file_parse_name", &FileParseName},
	{"g_file_poll_mountable", &FilePollMountable},
	{"g_file_poll_mountable_finish", &FilePollMountableFinish},
	{"g_file_query_default_handler", &FileQueryDefaultHandler},
	{"g_file_query_exists", &FileQueryExists},
	{"g_file_query_file_type", &FileQueryFileType},
	{"g_file_query_filesystem_info", &FileQueryFilesystemInfo},
	{"g_file_query_filesystem_info_async", &FileQueryFilesystemInfoAsync},
	{"g_file_query_filesystem_info_finish", &FileQueryFilesystemInfoFinish},
	{"g_file_query_info", &FileQueryInfo},
	{"g_file_query_info_async", &FileQueryInfoAsync},
	{"g_file_query_info_finish", &FileQueryInfoFinish},
	{"g_file_query_info_flags_get_type", &FileQueryInfoFlagsGetType},
	{"g_file_query_settable_attributes", &FileQuerySettableAttributes},
	{"g_file_query_writable_namespaces", &FileQueryWritableNamespaces},
	{"g_file_read", &FileRead},
	{"g_file_read_async", &FileReadAsync},
	{"g_file_read_finish", &FileReadFinish},
	{"g_file_replace", &FileReplace},
	{"g_file_replace_async", &FileReplaceAsync},
	{"g_file_replace_contents", &FileReplaceContents},
	{"g_file_replace_contents_async", &FileReplaceContentsAsync},
	{"g_file_replace_contents_finish", &FileReplaceContentsFinish},
	{"g_file_replace_finish", &FileReplaceFinish},
	{"g_file_replace_readwrite", &FileReplaceReadwrite},
	{"g_file_replace_readwrite_async", &FileReplaceReadwriteAsync},
	{"g_file_replace_readwrite_finish", &FileReplaceReadwriteFinish},
	{"g_file_resolve_relative_path", &FileResolveRelativePath},
	{"g_file_set_attribute", &FileSetAttribute},
	{"g_file_set_attribute_byte_string", &FileSetAttributeByteString},
	{"g_file_set_attribute_int32", &FileSetAttributeInt32},
	{"g_file_set_attribute_int64", &FileSetAttributeInt64},
	{"g_file_set_attribute_string", &FileSetAttributeString},
	{"g_file_set_attribute_uint32", &FileSetAttributeUint32},
	{"g_file_set_attribute_uint64", &FileSetAttributeUint64},
	{"g_file_set_attributes_async", &FileSetAttributesAsync},
	{"g_file_set_attributes_finish", &FileSetAttributesFinish},
	{"g_file_set_attributes_from_info", &FileSetAttributesFromInfo},
	{"g_file_set_display_name", &FileSetDisplayName},
	{"g_file_set_display_name_async", &FileSetDisplayNameAsync},
	{"g_file_set_display_name_finish", &FileSetDisplayNameFinish},
	{"g_file_start_mountable", &FileStartMountable},
	{"g_file_start_mountable_finish", &FileStartMountableFinish},
	{"g_file_stop_mountable", &FileStopMountable},
	{"g_file_stop_mountable_finish", &FileStopMountableFinish},
	{"g_file_supports_thread_contexts", &FileSupportsThreadContexts},
	{"g_file_trash", &FileTrash},
	{"g_file_type_get_type", &FileTypeGetType},
	{"g_file_unmount_mountable", &FileUnmountMountable},
	{"g_file_unmount_mountable_finish", &FileUnmountMountableFinish},
	{"g_file_unmount_mountable_with_operation", &FileUnmountMountableWithOperation},
	{"g_file_unmount_mountable_with_operation_finish", &FileUnmountMountableWithOperationFinish},
	{"g_filename_completer_get_completion_suffix", &FilenameCompleterGetCompletionSuffix},
	{"g_filename_completer_get_completions", &FilenameCompleterGetCompletions},
	{"g_filename_completer_get_type", &FilenameCompleterGetType},
	{"g_filename_completer_new", &FilenameCompleterNew},
	{"g_filename_completer_set_dirs_only", &FilenameCompleterSetDirsOnly},
	{"g_filesystem_preview_type_get_type", &FilesystemPreviewTypeGetType},
	{"g_filter_input_stream_get_base_stream", &FilterInputStreamGetBaseStream},
	{"g_filter_input_stream_get_close_base_stream", &FilterInputStreamGetCloseBaseStream},
	{"g_filter_input_stream_get_type", &FilterInputStreamGetType},
	{"g_filter_input_stream_set_close_base_stream", &FilterInputStreamSetCloseBaseStream},
	{"g_filter_output_stream_get_base_stream", &FilterOutputStreamGetBaseStream},
	{"g_filter_output_stream_get_close_base_stream", &FilterOutputStreamGetCloseBaseStream},
	{"g_filter_output_stream_get_type", &FilterOutputStreamGetType},
	{"g_filter_output_stream_set_close_base_stream", &FilterOutputStreamSetCloseBaseStream},
	{"g_icon_equal", &IconEqual},
	{"g_icon_get_type", &IconGetType},
	{"g_icon_hash", &IconHash},
	{"g_icon_new_for_string", &IconNewForString},
	{"g_icon_to_string", &IconToString},
	{"g_inet_address_get_family", &InetAddressGetFamily},
	{"g_inet_address_get_is_any", &InetAddressGetIsAny},
	{"g_inet_address_get_is_link_local", &InetAddressGetIsLinkLocal},
	{"g_inet_address_get_is_loopback", &InetAddressGetIsLoopback},
	{"g_inet_address_get_is_mc_global", &InetAddressGetIsMcGlobal},
	{"g_inet_address_get_is_mc_link_local", &InetAddressGetIsMcLinkLocal},
	{"g_inet_address_get_is_mc_node_local", &InetAddressGetIsMcNodeLocal},
	{"g_inet_address_get_is_mc_org_local", &InetAddressGetIsMcOrgLocal},
	{"g_inet_address_get_is_mc_site_local", &InetAddressGetIsMcSiteLocal},
	{"g_inet_address_get_is_multicast", &InetAddressGetIsMulticast},
	{"g_inet_address_get_is_site_local", &InetAddressGetIsSiteLocal},
	{"g_inet_address_get_native_size", &InetAddressGetNativeSize},
	{"g_inet_address_get_type", &InetAddressGetType},
	{"g_inet_address_new_any", &InetAddressNewAny},
	{"g_inet_address_new_from_bytes", &InetAddressNewFromBytes},
	{"g_inet_address_new_from_string", &InetAddressNewFromString},
	{"g_inet_address_new_loopback", &InetAddressNewLoopback},
	{"g_inet_address_to_bytes", &InetAddressToBytes},
	{"g_inet_address_to_string", &InetAddressToString},
	{"g_inet_socket_address_get_address", &InetSocketAddressGetAddress},
	{"g_inet_socket_address_get_port", &InetSocketAddressGetPort},
	{"g_inet_socket_address_get_type", &InetSocketAddressGetType},
	{"g_inet_socket_address_new", &InetSocketAddressNew},
	{"g_initable_get_type", &InitableGetType},
	{"g_initable_init", &InitableInit},
	{"g_initable_new", &InitableNew},
	{"g_initable_new_valist", &InitableNewValist},
	{"g_initable_newv", &InitableNewv},
	{"g_input_stream_clear_pending", &InputStreamClearPending},
	{"g_input_stream_close", &InputStreamClose},
	{"g_input_stream_close_async", &InputStreamCloseAsync},
	{"g_input_stream_close_finish", &InputStreamCloseFinish},
	{"g_input_stream_get_type", &InputStreamGetType},
	{"g_input_stream_has_pending", &InputStreamHasPending},
	{"g_input_stream_is_closed", &InputStreamIsClosed},
	{"g_input_stream_read", &InputStreamRead},
	{"g_input_stream_read_all", &InputStreamReadAll},
	{"g_input_stream_read_async", &InputStreamReadAsync},
	{"g_input_stream_read_finish", &InputStreamReadFinish},
	{"g_input_stream_set_pending", &InputStreamSetPending},
	{"g_input_stream_skip", &InputStreamSkip},
	{"g_input_stream_skip_async", &InputStreamSkipAsync},
	{"g_input_stream_skip_finish", &InputStreamSkipFinish},
	{"g_io_error_enum_get_type", &IoErrorEnumGetType},
	{"g_io_error_from_errno", &IoErrorFromErrno},
	{"g_io_error_from_win32_error", &IoErrorFromWin32Error},
	{"g_io_error_quark", &IoErrorQuark},
	{"g_io_extension_get_name", &IoExtensionGetName},
	{"g_io_extension_get_priority", &IoExtensionGetPriority},
	{"g_io_extension_get_type", &IoExtensionGetType},
	{"g_io_extension_point_get_extension_by_name", &IoExtensionPointGetExtensionByName},
	{"g_io_extension_point_get_extensions", &IoExtensionPointGetExtensions},
	{"g_io_extension_point_get_required_type", &IoExtensionPointGetRequiredType},
	{"g_io_extension_point_implement", &IoExtensionPointImplement},
	{"g_io_extension_point_lookup", &IoExtensionPointLookup},
	{"g_io_extension_point_register", &IoExtensionPointRegister},
	{"g_io_extension_point_set_required_type", &IoExtensionPointSetRequiredType},
	{"g_io_extension_ref_class", &IoExtensionRefClass},
	{"g_io_module_get_type", &IoModuleGetType},
	{"g_io_module_new", &IoModuleNew},
	{"g_io_modules_load_all_in_directory", &IoModulesLoadAllInDirectory},
	{"g_io_modules_scan_all_in_directory", &IoModulesScanAllInDirectory},
	{"g_io_scheduler_cancel_all_jobs", &IoSchedulerCancelAllJobs},
	{"g_io_scheduler_job_send_to_mainloop", &IoSchedulerJobSendToMainloop},
	{"g_io_scheduler_job_send_to_mainloop_async", &IoSchedulerJobSendToMainloopAsync},
	{"g_io_scheduler_push_job", &IoSchedulerPushJob},
	{"g_io_stream_clear_pending", &IoStreamClearPending},
	{"g_io_stream_close", &IoStreamClose},
	{"g_io_stream_close_async", &IoStreamCloseAsync},
	{"g_io_stream_close_finish", &IoStreamCloseFinish},
	{"g_io_stream_get_input_stream", &IoStreamGetInputStream},
	{"g_io_stream_get_output_stream", &IoStreamGetOutputStream},
	{"g_io_stream_get_type", &IoStreamGetType},
	{"g_io_stream_has_pending", &IoStreamHasPending},
	{"g_io_stream_is_closed", &IoStreamIsClosed},
	{"g_io_stream_set_pending", &IoStreamSetPending},
	{"g_io_stream_splice_async", &IoStreamSpliceAsync},
	{"g_io_stream_splice_finish", &IoStreamSpliceFinish},
	{"g_io_stream_splice_flags_get_type", &IoStreamSpliceFlagsGetType},
	{"g_keyfile_settings_backend_new", &KeyfileSettingsBackendNew},
	{"g_loadable_icon_get_type", &LoadableIconGetType},
	{"g_loadable_icon_load", &LoadableIconLoad},
	{"g_loadable_icon_load_async", &LoadableIconLoadAsync},
	{"g_loadable_icon_load_finish", &LoadableIconLoadFinish},
	{"g_memory_input_stream_add_data", &MemoryInputStreamAddData},
	{"g_memory_input_stream_get_type", &MemoryInputStreamGetType},
	{"g_memory_input_stream_new", &MemoryInputStreamNew},
	{"g_memory_input_stream_new_from_data", &MemoryInputStreamNewFromData},
	{"g_memory_output_stream_get_data", &MemoryOutputStreamGetData},
	{"g_memory_output_stream_get_data_size", &MemoryOutputStreamGetDataSize},
	{"g_memory_output_stream_get_size", &MemoryOutputStreamGetSize},
	{"g_memory_output_stream_get_type", &MemoryOutputStreamGetType},
	{"g_memory_output_stream_new", &MemoryOutputStreamNew},
	{"g_memory_output_stream_steal_data", &MemoryOutputStreamStealData},
	{"g_memory_settings_backend_new", &MemorySettingsBackendNew},
	{"g_mount_can_eject", &MountCanEject},
	{"g_mount_can_unmount", &MountCanUnmount},
	{"g_mount_eject", &MountEject},
	{"g_mount_eject_finish", &MountEjectFinish},
	{"g_mount_eject_with_operation", &MountEjectWithOperation},
	{"g_mount_eject_with_operation_finish", &MountEjectWithOperationFinish},
	{"g_mount_get_default_location", &MountGetDefaultLocation},
	{"g_mount_get_drive", &MountGetDrive},
	{"g_mount_get_icon", &MountGetIcon},
	{"g_mount_get_name", &MountGetName},
	{"g_mount_get_root", &MountGetRoot},
	{"g_mount_get_type", &MountGetType},
	{"g_mount_get_uuid", &MountGetUuid},
	{"g_mount_get_volume", &MountGetVolume},
	{"g_mount_guess_content_type", &MountGuessContentType},
	{"g_mount_guess_content_type_finish", &MountGuessContentTypeFinish},
	{"g_mount_guess_content_type_sync", &MountGuessContentTypeSync},
	{"g_mount_is_shadowed", &MountIsShadowed},
	{"g_mount_mount_flags_get_type", &MountMountFlagsGetType},
	{"g_mount_operation_get_anonymous", &MountOperationGetAnonymous},
	{"g_mount_operation_get_choice", &MountOperationGetChoice},
	{"g_mount_operation_get_domain", &MountOperationGetDomain},
	{"g_mount_operation_get_password", &MountOperationGetPassword},
	{"g_mount_operation_get_password_save", &MountOperationGetPasswordSave},
	{"g_mount_operation_get_type", &MountOperationGetType},
	{"g_mount_operation_get_username", &MountOperationGetUsername},
	{"g_mount_operation_new", &MountOperationNew},
	{"g_mount_operation_reply", &MountOperationReply},
	{"g_mount_operation_result_get_type", &MountOperationResultGetType},
	{"g_mount_operation_set_anonymous", &MountOperationSetAnonymous},
	{"g_mount_operation_set_choice", &MountOperationSetChoice},
	{"g_mount_operation_set_domain", &MountOperationSetDomain},
	{"g_mount_operation_set_password", &MountOperationSetPassword},
	{"g_mount_operation_set_password_save", &MountOperationSetPasswordSave},
	{"g_mount_operation_set_username", &MountOperationSetUsername},
	{"g_mount_remount", &MountRemount},
	{"g_mount_remount_finish", &MountRemountFinish},
	{"g_mount_shadow", &MountShadow},
	{"g_mount_unmount", &MountUnmount},
	{"g_mount_unmount_finish", &MountUnmountFinish},
	{"g_mount_unmount_flags_get_type", &MountUnmountFlagsGetType},
	{"g_mount_unmount_with_operation", &MountUnmountWithOperation},
	{"g_mount_unmount_with_operation_finish", &MountUnmountWithOperationFinish},
	{"g_mount_unshadow", &MountUnshadow},
	{"g_native_volume_monitor_get_type", &NativeVolumeMonitorGetType},
	{"g_network_address_get_hostname", &NetworkAddressGetHostname},
	{"g_network_address_get_port", &NetworkAddressGetPort},
	{"g_network_address_get_scheme", &NetworkAddressGetScheme},
	{"g_network_address_get_type", &NetworkAddressGetType},
	{"g_network_address_new", &NetworkAddressNew},
	{"g_network_address_parse", &NetworkAddressParse},
	{"g_network_address_parse_uri", &NetworkAddressParseUri},
	{"g_network_service_get_domain", &NetworkServiceGetDomain},
	{"g_network_service_get_protocol", &NetworkServiceGetProtocol},
	{"g_network_service_get_scheme", &NetworkServiceGetScheme},
	{"g_network_service_get_service", &NetworkServiceGetService},
	{"g_network_service_get_type", &NetworkServiceGetType},
	{"g_network_service_new", &NetworkServiceNew},
	{"g_network_service_set_scheme", &NetworkServiceSetScheme},
	{"g_null_settings_backend_new", &NullSettingsBackendNew},
	{"g_output_stream_clear_pending", &OutputStreamClearPending},
	{"g_output_stream_close", &OutputStreamClose},
	{"g_output_stream_close_async", &OutputStreamCloseAsync},
	{"g_output_stream_close_finish", &OutputStreamCloseFinish},
	{"g_output_stream_flush", &OutputStreamFlush},
	{"g_output_stream_flush_async", &OutputStreamFlushAsync},
	{"g_output_stream_flush_finish", &OutputStreamFlushFinish},
	{"g_output_stream_get_type", &OutputStreamGetType},
	{"g_output_stream_has_pending", &OutputStreamHasPending},
	{"g_output_stream_is_closed", &OutputStreamIsClosed},
	{"g_output_stream_is_closing", &OutputStreamIsClosing},
	{"g_output_stream_set_pending", &OutputStreamSetPending},
	{"g_output_stream_splice", &OutputStreamSplice},
	{"g_output_stream_splice_async", &OutputStreamSpliceAsync},
	{"g_output_stream_splice_finish", &OutputStreamSpliceFinish},
	{"g_output_stream_splice_flags_get_type", &OutputStreamSpliceFlagsGetType},
	{"g_output_stream_write", &OutputStreamWrite},
	{"g_output_stream_write_all", &OutputStreamWriteAll},
	{"g_output_stream_write_async", &OutputStreamWriteAsync},
	{"g_output_stream_write_finish", &OutputStreamWriteFinish},
	{"g_password_save_get_type", &PasswordSaveGetType},
	{"g_permission_acquire", &PermissionAcquire},
	{"g_permission_acquire_async", &PermissionAcquireAsync},
	{"g_permission_acquire_finish", &PermissionAcquireFinish},
	{"g_permission_get_allowed", &PermissionGetAllowed},
	{"g_permission_get_can_acquire", &PermissionGetCanAcquire},
	{"g_permission_get_can_release", &PermissionGetCanRelease},
	{"g_permission_get_type", &PermissionGetType},
	{"g_permission_impl_update", &PermissionImplUpdate},
	{"g_permission_release", &PermissionRelease},
	{"g_permission_release_async", &PermissionReleaseAsync},
	{"g_permission_release_finish", &PermissionReleaseFinish},
	{"g_pollable_input_stream_can_poll", &PollableInputStreamCanPoll},
	{"g_pollable_input_stream_create_source", &PollableInputStreamCreateSource},
	{"g_pollable_input_stream_get_type", &PollableInputStreamGetType},
	{"g_pollable_input_stream_is_readable", &PollableInputStreamIsReadable},
	{"g_pollable_input_stream_read_nonblocking", &PollableInputStreamReadNonblocking},
	{"g_pollable_output_stream_can_poll", &PollableOutputStreamCanPoll},
	{"g_pollable_output_stream_create_source", &PollableOutputStreamCreateSource},
	{"g_pollable_output_stream_get_type", &PollableOutputStreamGetType},
	{"g_pollable_output_stream_is_writable", &PollableOutputStreamIsWritable},
	{"g_pollable_output_stream_write_nonblocking", &PollableOutputStreamWriteNonblocking},
	{"g_pollable_source_new", &PollableSourceNew},
	{"g_proxy_address_enumerator_get_type", &ProxyAddressEnumeratorGetType},
	{"g_proxy_address_get_destination_hostname", &ProxyAddressGetDestinationHostname},
	{"g_proxy_address_get_destination_port", &ProxyAddressGetDestinationPort},
	{"g_proxy_address_get_password", &ProxyAddressGetPassword},
	{"g_proxy_address_get_protocol", &ProxyAddressGetProtocol},
	{"g_proxy_address_get_type", &ProxyAddressGetType},
	{"g_proxy_address_get_username", &ProxyAddressGetUsername},
	{"g_proxy_address_new", &ProxyAddressNew},
	{"g_proxy_connect", &ProxyConnect},
	{"g_proxy_connect_async", &ProxyConnectAsync},
	{"g_proxy_connect_finish", &ProxyConnectFinish},
	{"g_proxy_get_default_for_protocol", &ProxyGetDefaultForProtocol},
	{"g_proxy_get_type", &ProxyGetType},
	{"g_proxy_resolver_get_default", &ProxyResolverGetDefault},
	{"g_proxy_resolver_get_type", &ProxyResolverGetType},
	{"g_proxy_resolver_is_supported", &ProxyResolverIsSupported},
	{"g_proxy_resolver_lookup", &ProxyResolverLookup},
	{"g_proxy_resolver_lookup_async", &ProxyResolverLookupAsync},
	{"g_proxy_resolver_lookup_finish", &ProxyResolverLookupFinish},
	{"g_proxy_supports_hostname", &ProxySupportsHostname},
	{"g_resolver_error_get_type", &ResolverErrorGetType},
	{"g_resolver_error_quark", &ResolverErrorQuark},
	{"g_resolver_free_addresses", &ResolverFreeAddresses},
	{"g_resolver_free_targets", &ResolverFreeTargets},
	{"g_resolver_get_default", &ResolverGetDefault},
	{"g_resolver_get_type", &ResolverGetType},
	{"g_resolver_lookup_by_address", &ResolverLookupByAddress},
	{"g_resolver_lookup_by_address_async", &ResolverLookupByAddressAsync},
	{"g_resolver_lookup_by_address_finish", &ResolverLookupByAddressFinish},
	{"g_resolver_lookup_by_name", &ResolverLookupByName},
	{"g_resolver_lookup_by_name_async", &ResolverLookupByNameAsync},
	{"g_resolver_lookup_by_name_finish", &ResolverLookupByNameFinish},
	{"g_resolver_lookup_service", &ResolverLookupService},
	{"g_resolver_lookup_service_async", &ResolverLookupServiceAsync},
	{"g_resolver_lookup_service_finish", &ResolverLookupServiceFinish},
	{"g_resolver_set_default", &ResolverSetDefault},
	{"g_seekable_can_seek", &SeekableCanSeek},
	{"g_seekable_can_truncate", &SeekableCanTruncate},
	{"g_seekable_get_type", &SeekableGetType},
	{"g_seekable_seek", &SeekableSeek},
	{"g_seekable_tell", &SeekableTell},
	{"g_seekable_truncate", &SeekableTruncate},
	{"g_settings_apply", &SettingsApply},
	{"g_settings_backend_changed", &SettingsBackendChanged},
	{"g_settings_backend_changed_tree", &SettingsBackendChangedTree},
	{"g_settings_backend_flatten_tree", &SettingsBackendFlattenTree},
	{"g_settings_backend_get_default", &SettingsBackendGetDefault},
	{"g_settings_backend_get_type", &SettingsBackendGetType},
	{"g_settings_backend_keys_changed", &SettingsBackend_keysChanged},
	{"g_settings_backend_path_changed", &SettingsBackendPathChanged},
	{"g_settings_backend_path_writable_changed", &SettingsBackendPathWritableChanged},
	{"g_settings_backend_writable_changed", &SettingsBackendWritableChanged},
	{"g_settings_bind", &SettingsBind},
	{"g_settings_bind_flags_get_type", &SettingsBindFlagsGetType},
	{"g_settings_bind_with_mapping", &SettingsBindWithMapping},
	{"g_settings_bind_writable", &SettingsBindWritable},
	{"g_settings_delay", &SettingsDelay},
	{"g_settings_get", &SettingsGet},
	{"g_settings_get_boolean", &SettingsGetBoolean},
	{"g_settings_get_child", &SettingsGetChild},
	{"g_settings_get_double", &SettingsGetDouble},
	{"g_settings_get_enum", &SettingsGetEnum},
	{"g_settings_get_flags", &SettingsGetFlags},
	{"g_settings_get_has_unapplied", &SettingsGetHasUnapplied},
	{"g_settings_get_int", &SettingsGetInt},
	{"g_settings_get_mapped", &SettingsGetMapped},
	{"g_settings_get_range", &SettingsGetRange},
	{"g_settings_get_string", &SettingsGetString},
	{"g_settings_get_strv", &SettingsGetStrv},
	{"g_settings_get_type", &SettingsGetType},
	{"g_settings_get_value", &SettingsGetValue},
	{"g_settings_is_writable", &SettingsIsWritable},
	{"g_settings_list_children", &SettingsListChildren},
	{"g_settings_list_keys", &SettingsList_keys},
	{"g_settings_list_relocatable_schemas", &SettingsListRelocatableSchemas},
	{"g_settings_list_schemas", &SettingsListSchemas},
	{"g_settings_new", &SettingsNew},
	{"g_settings_new_with_backend", &SettingsNewWithBackend},
	{"g_settings_new_with_backend_and_path", &SettingsNewWithBackendAndPath},
	{"g_settings_new_with_path", &SettingsNewWithPath},
	{"g_settings_range_check", &SettingsRangeCheck},
	{"g_settings_reset", &SettingsReset},
	{"g_settings_revert", &SettingsRevert},
	{"g_settings_set", &SettingsSet},
	{"g_settings_set_boolean", &SettingsSetBoolean},
	{"g_settings_set_double", &SettingsSetDouble},
	{"g_settings_set_enum", &SettingsSetEnum},
	{"g_settings_set_flags", &SettingsSetFlags},
	{"g_settings_set_int", &SettingsSetInt},
	{"g_settings_set_string", &SettingsSetString},
	{"g_settings_set_strv", &SettingsSetStrv},
	{"g_settings_set_value", &SettingsSetValue},
	{"g_settings_sync", &SettingsSync},
	{"g_settings_unbind", &SettingsUnbind},
	// Undocumented {"g_simple_action_get_parameter_type", &SimpleActionGetParameterType},
	{"g_simple_action_get_type", &SimpleActionGetType},
	{"g_simple_action_group_get_type", &SimpleActionGroupGetType},
	{"g_simple_action_group_insert", &SimpleActionGroupInsert},
	{"g_simple_action_group_lookup", &SimpleActionGroupLookup},
	{"g_simple_action_group_new", &SimpleActionGroupNew},
	{"g_simple_action_group_remove", &SimpleActionGroupRemove},
	{"g_simple_action_new", &SimpleActionNew},
	{"g_simple_action_new_stateful", &SimpleActionNewStateful},
	{"g_simple_action_set_enabled", &SimpleActionSetEnabled},
	{"g_simple_async_report_error_in_idle", &SimpleAsyncReportErrorInIdle},
	{"g_simple_async_report_gerror_in_idle", &SimpleAsyncReportGerrorInIdle},
	{"g_simple_async_report_take_gerror_in_idle", &SimpleAsyncReportTakeGerrorInIdle},
	{"g_simple_async_result_complete", &SimpleAsyncResultComplete},
	{"g_simple_async_result_complete_in_idle", &SimpleAsyncResultCompleteInIdle},
	{"g_simple_async_result_get_op_res_gboolean", &SimpleAsyncResultGetOpResGboolean},
	{"g_simple_async_result_get_op_res_gpointer", &SimpleAsyncResultGetOpResGpointer},
	{"g_simple_async_result_get_op_res_gssize", &SimpleAsyncResultGetOpResGssize},
	{"g_simple_async_result_get_source_tag", &SimpleAsyncResultGetSourceTag},
	{"g_simple_async_result_get_type", &SimpleAsyncResultGetType},
	{"g_simple_async_result_is_valid", &SimpleAsyncResultIsValid},
	{"g_simple_async_result_new", &SimpleAsyncResultNew},
	{"g_simple_async_result_new_error", &SimpleAsyncResultNewError},
	{"g_simple_async_result_new_from_error", &SimpleAsyncResultNewFromError},
	{"g_simple_async_result_new_take_error", &SimpleAsyncResultNewTakeError},
	{"g_simple_async_result_propagate_error", &SimpleAsyncResultPropagateError},
	{"g_simple_async_result_run_in_thread", &SimpleAsyncResultRunInThread},
	{"g_simple_async_result_set_error", &SimpleAsyncResultSetError},
	{"g_simple_async_result_set_error_va", &SimpleAsyncResultSetErrorVa},
	{"g_simple_async_result_set_from_error", &SimpleAsyncResultSetFromError},
	{"g_simple_async_result_set_handle_cancellation", &SimpleAsyncResultSetHandleCancellation},
	{"g_simple_async_result_set_op_res_gboolean", &SimpleAsyncResultSetOpResGboolean},
	{"g_simple_async_result_set_op_res_gpointer", &SimpleAsyncResultSetOpResGpointer},
	{"g_simple_async_result_set_op_res_gssize", &SimpleAsyncResultSetOpResGssize},
	{"g_simple_async_result_take_error", &SimpleAsyncResultTakeError},
	{"g_simple_permission_get_type", &SimplePermissionGetType},
	{"g_simple_permission_new", &SimplePermissionNew},
	{"g_socket_accept", &SocketAccept},
	{"g_socket_address_enumerator_get_type", &SocketAddressEnumeratorGetType},
	{"g_socket_address_enumerator_next", &SocketAddressEnumeratorNext},
	{"g_socket_address_enumerator_next_async", &SocketAddressEnumeratorNextAsync},
	{"g_socket_address_enumerator_next_finish", &SocketAddressEnumeratorNextFinish},
	{"g_socket_address_get_family", &SocketAddressGetFamily},
	{"g_socket_address_get_native_size", &SocketAddressGetNativeSize},
	{"g_socket_address_get_type", &SocketAddressGetType},
	{"g_socket_address_new_from_native", &SocketAddressNewFromNative},
	{"g_socket_address_to_native", &SocketAddressToNative},
	{"g_socket_bind", &SocketBind},
	{"g_socket_check_connect_result", &SocketCheckConnectResult},
	{"g_socket_client_add_application_proxy", &SocketClientAddApplicationProxy},
	{"g_socket_client_connect", &SocketClientConnect},
	{"g_socket_client_connect_async", &SocketClientConnectAsync},
	{"g_socket_client_connect_finish", &SocketClientConnectFinish},
	{"g_socket_client_connect_to_host", &SocketClientConnectToHost},
	{"g_socket_client_connect_to_host_async", &SocketClientConnectToHostAsync},
	{"g_socket_client_connect_to_host_finish", &SocketClientConnectToHostFinish},
	{"g_socket_client_connect_to_service", &SocketClientConnectToService},
	{"g_socket_client_connect_to_service_async", &SocketClientConnectToServiceAsync},
	{"g_socket_client_connect_to_service_finish", &SocketClientConnectToServiceFinish},
	{"g_socket_client_connect_to_uri", &SocketClientConnectToUri},
	{"g_socket_client_connect_to_uri_async", &SocketClientConnectToUriAsync},
	{"g_socket_client_connect_to_uri_finish", &SocketClientConnectToUriFinish},
	{"g_socket_client_get_enable_proxy", &SocketClientGetEnableProxy},
	{"g_socket_client_get_family", &SocketClientGetFamily},
	{"g_socket_client_get_local_address", &SocketClientGetLocalAddress},
	{"g_socket_client_get_protocol", &SocketClientGetProtocol},
	{"g_socket_client_get_socket_type", &SocketClientGetSocketType},
	{"g_socket_client_get_timeout", &SocketClientGetTimeout},
	{"g_socket_client_get_tls", &SocketClientGetTls},
	{"g_socket_client_get_tls_validation_flags", &SocketClientGetTlsValidationFlags},
	{"g_socket_client_get_type", &SocketClientGetType},
	{"g_socket_client_new", &SocketClientNew},
	{"g_socket_client_set_enable_proxy", &SocketClientSetEnableProxy},
	{"g_socket_client_set_family", &SocketClientSetFamily},
	{"g_socket_client_set_local_address", &SocketClientSetLocalAddress},
	{"g_socket_client_set_protocol", &SocketClientSetProtocol},
	{"g_socket_client_set_socket_type", &SocketClientSetSocketType},
	{"g_socket_client_set_timeout", &SocketClientSetTimeout},
	{"g_socket_client_set_tls", &SocketClientSetTls},
	{"g_socket_client_set_tls_validation_flags", &SocketClientSetTlsValidationFlags},
	{"g_socket_close", &SocketClose},
	{"g_socket_condition_check", &SocketConditionCheck},
	{"g_socket_condition_wait", &SocketConditionWait},
	{"g_socket_connect", &SocketConnect},
	{"g_socket_connectable_enumerate", &SocketConnectableEnumerate},
	{"g_socket_connectable_get_type", &SocketConnectableGetType},
	{"g_socket_connectable_proxy_enumerate", &SocketConnectableProxyEnumerate},
	{"g_socket_connection_factory_create_connection", &SocketConnectionFactoryCreateConnection},
	{"g_socket_connection_factory_lookup_type", &SocketConnectionFactoryLookupType},
	{"g_socket_connection_factory_register_type", &SocketConnectionFactoryRegisterType},
	{"g_socket_connection_get_local_address", &SocketConnectionGetLocalAddress},
	{"g_socket_connection_get_remote_address", &SocketConnectionGetRemoteAddress},
	{"g_socket_connection_get_socket", &SocketConnectionGetSocket},
	{"g_socket_connection_get_type", &SocketConnectionGetType},
	{"g_socket_control_message_deserialize", &SocketControlMessageDeserialize},
	{"g_socket_control_message_get_level", &SocketControlMessageGetLevel},
	{"g_socket_control_message_get_msg_type", &SocketControlMessageGetMsgType},
	{"g_socket_control_message_get_size", &SocketControlMessageGetSize},
	{"g_socket_control_message_get_type", &SocketControlMessageGetType},
	{"g_socket_control_message_serialize", &SocketControlMessageSerialize},
	{"g_socket_create_source", &SocketCreateSource},
	{"g_socket_family_get_type", &SocketFamilyGetType},
	{"g_socket_get_blocking", &SocketGetBlocking},
	{"g_socket_get_credentials", &SocketGetCredentials},
	{"g_socket_get_family", &SocketGetFamily},
	{"g_socket_get_fd", &SocketGetFd},
	{"g_socket_get_keepalive", &SocketGet_keepalive},
	{"g_socket_get_listen_backlog", &SocketGetListenBacklog},
	{"g_socket_get_local_address", &SocketGetLocalAddress},
	{"g_socket_get_protocol", &SocketGetProtocol},
	{"g_socket_get_remote_address", &SocketGetRemoteAddress},
	{"g_socket_get_socket_type", &SocketGetSocketType},
	{"g_socket_get_timeout", &SocketGetTimeout},
	{"g_socket_get_type", &SocketGetType},
	{"g_socket_is_closed", &SocketIsClosed},
	{"g_socket_is_connected", &SocketIsConnected},
	{"g_socket_listen", &SocketListen},
	{"g_socket_listener_accept", &SocketListenerAccept},
	{"g_socket_listener_accept_async", &SocketListenerAcceptAsync},
	{"g_socket_listener_accept_finish", &SocketListenerAcceptFinish},
	{"g_socket_listener_accept_socket", &SocketListenerAcceptSocket},
	{"g_socket_listener_accept_socket_async", &SocketListenerAcceptSocketAsync},
	{"g_socket_listener_accept_socket_finish", &SocketListenerAcceptSocketFinish},
	{"g_socket_listener_add_address", &SocketListenerAddAddress},
	{"g_socket_listener_add_any_inet_port", &SocketListenerAddAnyInetPort},
	{"g_socket_listener_add_inet_port", &SocketListenerAddInetPort},
	{"g_socket_listener_add_socket", &SocketListenerAddSocket},
	{"g_socket_listener_close", &SocketListenerClose},
	{"g_socket_listener_get_type", &SocketListenerGetType},
	{"g_socket_listener_new", &SocketListenerNew},
	{"g_socket_listener_set_backlog", &SocketListenerSetBacklog},
	{"g_socket_msg_flags_get_type", &SocketMsgFlagsGetType},
	{"g_socket_new", &SocketNew},
	{"g_socket_new_from_fd", &SocketNewFromFd},
	{"g_socket_protocol_get_type", &SocketProtocolGetType},
	{"g_socket_receive", &SocketReceive},
	{"g_socket_receive_from", &SocketReceiveFrom},
	{"g_socket_receive_message", &SocketReceiveMessage},
	{"g_socket_receive_with_blocking", &SocketReceiveWithBlocking},
	{"g_socket_send", &SocketSend},
	{"g_socket_send_message", &SocketSendMessage},
	{"g_socket_send_to", &SocketSendTo},
	{"g_socket_send_with_blocking", &SocketSendWithBlocking},
	{"g_socket_service_get_type", &SocketServiceGetType},
	{"g_socket_service_is_active", &SocketServiceIsActive},
	{"g_socket_service_new", &SocketServiceNew},
	{"g_socket_service_start", &SocketServiceStart},
	{"g_socket_service_stop", &SocketServiceStop},
	{"g_socket_set_blocking", &SocketSetBlocking},
	{"g_socket_set_keepalive", &SocketSet_keepalive},
	{"g_socket_set_listen_backlog", &SocketSetListenBacklog},
	{"g_socket_set_timeout", &SocketSetTimeout},
	{"g_socket_shutdown", &SocketShutdown},
	{"g_socket_speaks_ipv4", &SocketSpeaksIpv4},
	{"g_socket_type_get_type", &SocketTypeGetType},
	{"g_srv_target_copy", &SrvTargetCopy},
	{"g_srv_target_free", &SrvTargetFree},
	{"g_srv_target_get_hostname", &SrvTargetGetHostname},
	{"g_srv_target_get_port", &SrvTargetGetPort},
	{"g_srv_target_get_priority", &SrvTargetGetPriority},
	{"g_srv_target_get_type", &SrvTargetGetType},
	{"g_srv_target_get_weight", &SrvTargetGetWeight},
	{"g_srv_target_list_sort", &SrvTargetListSort},
	{"g_srv_target_new", &SrvTargetNew},
	{"g_tcp_connection_get_graceful_disconnect", &TcpConnectionGetGracefulDisconnect},
	{"g_tcp_connection_get_type", &TcpConnectionGetType},
	{"g_tcp_connection_set_graceful_disconnect", &TcpConnectionSetGracefulDisconnect},
	{"g_tcp_wrapper_connection_get_base_io_stream", &TcpWrapperConnectionGetBaseIoStream},
	{"g_tcp_wrapper_connection_get_type", &TcpWrapperConnectionGetType},
	{"g_tcp_wrapper_connection_new", &TcpWrapperConnectionNew},
	{"g_themed_icon_append_name", &ThemedIconAppendName},
	{"g_themed_icon_get_names", &ThemedIconGetNames},
	{"g_themed_icon_get_type", &ThemedIconGetType},
	{"g_themed_icon_new", &ThemedIconNew},
	{"g_themed_icon_new_from_names", &ThemedIconNewFromNames},
	{"g_themed_icon_new_with_default_fallbacks", &ThemedIconNewWithDefaultFallbacks},
	{"g_themed_icon_prepend_name", &ThemedIconPrependName},
	// Undocumented {"g_threaded_resolver_get_type", &ThreadedResolverGetType},
	{"g_threaded_socket_service_get_type", &ThreadedSocketServiceGetType},
	{"g_threaded_socket_service_new", &ThreadedSocketServiceNew},
	{"g_tls_authentication_mode_get_type", &TlsAuthenticationModeGetType},
	{"g_tls_backend_get_certificate_type", &TlsBackendGetCertificateType},
	{"g_tls_backend_get_client_connection_type", &TlsBackendGetClientConnectionType},
	{"g_tls_backend_get_default", &TlsBackendGetDefault},
	{"g_tls_backend_get_server_connection_type", &TlsBackendGetServerConnectionType},
	{"g_tls_backend_get_type", &TlsBackendGetType},
	{"g_tls_backend_supports_tls", &TlsBackendSupportsTls},
	{"g_tls_certificate_flags_get_type", &TlsCertificateFlagsGetType},
	{"g_tls_certificate_get_issuer", &TlsCertificateGetIssuer},
	{"g_tls_certificate_get_type", &TlsCertificateGetType},
	{"g_tls_certificate_list_new_from_file", &TlsCertificateListNewFromFile},
	{"g_tls_certificate_new_from_file", &TlsCertificateNewFromFile},
	{"g_tls_certificate_new_from_files", &TlsCertificateNewFromFiles},
	{"g_tls_certificate_new_from_pem", &TlsCertificateNewFromPem},
	{"g_tls_certificate_verify", &TlsCertificateVerify},
	{"g_tls_client_connection_get_accepted_cas", &TlsClientConnectionGetAcceptedCas},
	{"g_tls_client_connection_get_server_identity", &TlsClientConnectionGetServerIdentity},
	{"g_tls_client_connection_get_type", &TlsClientConnectionGetType},
	{"g_tls_client_connection_get_use_ssl3", &TlsClientConnectionGetUseSsl3},
	{"g_tls_client_connection_get_validation_flags", &TlsClientConnectionGetValidationFlags},
	{"g_tls_client_connection_new", &TlsClientConnectionNew},
	{"g_tls_client_connection_set_server_identity", &TlsClientConnectionSetServerIdentity},
	{"g_tls_client_connection_set_use_ssl3", &TlsClientConnectionSetUseSsl3},
	{"g_tls_client_connection_set_validation_flags", &TlsClientConnectionSetValidationFlags},
	{"g_tls_connection_emit_accept_certificate", &TlsConnectionEmitAcceptCertificate},
	{"g_tls_connection_get_certificate", &TlsConnectionGetCertificate},
	{"g_tls_connection_get_peer_certificate", &TlsConnectionGetPeerCertificate},
	{"g_tls_connection_get_peer_certificate_errors", &TlsConnectionGetPeerCertificateErrors},
	{"g_tls_connection_get_rehandshake_mode", &TlsConnectionGetRehandshakeMode},
	{"g_tls_connection_get_require_close_notify", &TlsConnectionGetRequireCloseNotify},
	{"g_tls_connection_get_type", &TlsConnectionGetType},
	{"g_tls_connection_get_use_system_certdb", &TlsConnectionGetUseSystemCertdb},
	{"g_tls_connection_handshake", &TlsConnectionHandshake},
	{"g_tls_connection_handshake_async", &TlsConnectionHandshakeAsync},
	{"g_tls_connection_handshake_finish", &TlsConnectionHandshakeFinish},
	{"g_tls_connection_set_certificate", &TlsConnectionSetCertificate},
	{"g_tls_connection_set_rehandshake_mode", &TlsConnectionSetRehandshakeMode},
	{"g_tls_connection_set_require_close_notify", &TlsConnectionSetRequireCloseNotify},
	{"g_tls_connection_set_use_system_certdb", &TlsConnectionSetUseSystemCertdb},
	{"g_tls_error_get_type", &TlsErrorGetType},
	{"g_tls_error_quark", &TlsErrorQuark},
	{"g_tls_rehandshake_mode_get_type", &TlsRehandshakeModeGetType},
	{"g_tls_server_connection_get_type", &TlsServerConnectionGetType},
	{"g_tls_server_connection_new", &TlsServerConnectionNew},
	{"g_unix_socket_address_type_get_type", &UnixSocketAddressTypeGetType},
	{"g_vfs_get_default", &VfsGetDefault},
	{"g_vfs_get_file_for_path", &VfsGetFileForPath},
	{"g_vfs_get_file_for_uri", &VfsGetFileForUri},
	{"g_vfs_get_local", &VfsGetLocal},
	{"g_vfs_get_supported_uri_schemes", &VfsGetSupportedUriSchemes},
	{"g_vfs_get_type", &VfsGetType},
	{"g_vfs_is_active", &VfsIsActive},
	{"g_vfs_parse_name", &VfsParseName},
	{"g_volume_can_eject", &VolumeCanEject},
	{"g_volume_can_mount", &VolumeCanMount},
	{"g_volume_eject", &VolumeEject},
	{"g_volume_eject_finish", &VolumeEjectFinish},
	{"g_volume_eject_with_operation", &VolumeEjectWithOperation},
	{"g_volume_eject_with_operation_finish", &VolumeEjectWithOperationFinish},
	{"g_volume_enumerate_identifiers", &VolumeEnumerateIdentifiers},
	{"g_volume_get_activation_root", &VolumeGetActivationRoot},
	{"g_volume_get_drive", &VolumeGetDrive},
	{"g_volume_get_icon", &VolumeGetIcon},
	{"g_volume_get_identifier", &VolumeGetIdentifier},
	{"g_volume_get_mount", &VolumeGetMount},
	{"g_volume_get_name", &VolumeGetName},
	{"g_volume_get_type", &VolumeGetType},
	{"g_volume_get_uuid", &VolumeGetUuid},
	{"g_volume_monitor_adopt_orphan_mount", &VolumeMonitorAdoptOrphanMount},
	{"g_volume_monitor_get", &VolumeMonitorGet},
	{"g_volume_monitor_get_connected_drives", &VolumeMonitorGetConnectedDrives},
	{"g_volume_monitor_get_mount_for_uuid", &VolumeMonitorGetMountForUuid},
	{"g_volume_monitor_get_mounts", &VolumeMonitorGetMounts},
	{"g_volume_monitor_get_type", &VolumeMonitorGetType},
	{"g_volume_monitor_get_volume_for_uuid", &VolumeMonitorGetVolumeForUuid},
	{"g_volume_monitor_get_volumes", &VolumeMonitorGetVolumes},
	{"g_volume_mount", &VolumeMount},
	{"g_volume_mount_finish", &VolumeMountFinish},
	{"g_volume_should_automount", &VolumeShouldAutomount},
	{"g_win32_input_stream_get_close_handle", &Win32InputStreamGetCloseHandle},
	{"g_win32_input_stream_get_handle", &Win32InputStreamGetHandle},
	{"g_win32_input_stream_get_type", &Win32InputStreamGetType},
	{"g_win32_input_stream_new", &Win32InputStreamNew},
	{"g_win32_input_stream_set_close_handle", &Win32InputStreamSetCloseHandle},
	{"g_win32_output_stream_get_close_handle", &Win32OutputStreamGetCloseHandle},
	{"g_win32_output_stream_get_handle", &Win32OutputStreamGetHandle},
	{"g_win32_output_stream_get_type", &Win32OutputStreamGetType},
	{"g_win32_output_stream_new", &Win32OutputStreamNew},
	{"g_win32_output_stream_set_close_handle", &Win32OutputStreamSetCloseHandle},
	// Undocumented {"g_win32_resolver_get_type", &Win32ResolverGetType},
	{"g_zlib_compressor_format_get_type", &ZlibCompressorFormatGetType},
	{"g_zlib_compressor_get_file_info", &ZlibCompressorGetFileInfo},
	{"g_zlib_compressor_get_type", &ZlibCompressorGetType},
	{"g_zlib_compressor_new", &ZlibCompressorNew},
	{"g_zlib_compressor_set_file_info", &ZlibCompressorSetFileInfo},
	{"g_zlib_decompressor_get_file_info", &ZlibDecompressorGetFileInfo},
	{"g_zlib_decompressor_get_type", &ZlibDecompressorGetType},
	{"g_zlib_decompressor_new", &ZlibDecompressorNew},
}
