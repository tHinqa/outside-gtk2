// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gio

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type File struct{}

var (
	FileGetType              func() O.Type
	FileNewForPath           func(path string) *File
	FileNewForUri            func(uri string) *File
	FileNewForCommandlineArg func(arg string) *File

	FileParseName func(parseName string) *File
	FileHash      func(file T.Gconstpointer) uint

	FileAppendTo                            func(f *File, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream
	FileAppendToAsync                       func(f *File, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileAppendToFinish                      func(f *File, res *AsyncResult, err **T.GError) *FileOutputStream
	FileCopy                                func(f, destination *File, flags FileCopyFlags, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, err **T.GError) bool
	FileCopyAsync                           func(f, destination *File, flags FileCopyFlags, ioPriority int, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, callback AsyncReadyCallback, userData T.Gpointer)
	FileCopyAttributes                      func(f, destination *File, flags FileCopyFlags, cancellable *Cancellable, err **T.GError) bool
	FileCopyFinish                          func(f *File, res *AsyncResult, err **T.GError) bool
	FileCreate                              func(f *File, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream
	FileCreateAsync                         func(f *File, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileCreateFinish                        func(f *File, res *AsyncResult, err **T.GError) *FileOutputStream
	FileCreateReadwrite                     func(f *File, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileIOStream
	FileCreateReadwriteAsync                func(f *File, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileCreateReadwriteFinish               func(f *File, res *AsyncResult, err **T.GError) *FileIOStream
	FileDelete                              func(f *File, cancellable *Cancellable, err **T.GError) bool
	FileDup                                 func(f *File) *File
	FileEjectMountable                      func(f *File, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileEjectMountableFinish                func(f *File, result *AsyncResult, err **T.GError) bool
	FileEjectMountableWithOperation         func(f *File, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileEjectMountableWithOperationFinish   func(f *File, result *AsyncResult, err **T.GError) bool
	FileEnumerateChildren                   func(f *File, attributes string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) *FileEnumerator
	FileEnumerateChildrenAsync              func(f *File, attributes string, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileEnumerateChildrenFinish             func(f *File, res *AsyncResult, err **T.GError) *FileEnumerator
	FileEqual                               func(f, file2 *File) bool
	FileFindEnclosingMount                  func(f *File, cancellable *Cancellable, err **T.GError) *Mount
	FileFindEnclosingMountAsync             func(f *File, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileFindEnclosingMountFinish            func(f *File, res *AsyncResult, err **T.GError) *Mount
	FileGetBasename                         func(f *File) string
	FileGetChild                            func(f *File, name string) *File
	FileGetChildForDisplayName              func(f *File, displayName string, err **T.GError) *File
	FileGetParent                           func(f *File) *File
	FileGetParseName                        func(f *File) string
	FileGetPath                             func(f *File) string
	FileGetRelativePath                     func(f *File, descendant *File) string
	FileGetUri                              func(f *File) string
	FileGetUriScheme                        func(f *File) string
	FileHasParent                           func(f, parent *File) bool
	FileHasPrefix                           func(f, prefix *File) bool
	FileHasUriScheme                        func(f *File, uriScheme string) bool
	FileIsNative                            func(f *File) bool
	FileLoadContents                        func(f *File, cancellable *Cancellable, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) bool
	FileLoadContentsAsync                   func(f *File, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileLoadContentsFinish                  func(f *File, res *AsyncResult, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) bool
	FileLoadPartialContentsAsync            func(f *File, cancellable *Cancellable, readMoreCallback FileReadMoreCallback, callback AsyncReadyCallback, userData T.Gpointer)
	FileLoadPartialContentsFinish           func(f *File, res *AsyncResult, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) bool
	FileMakeDirectory                       func(f *File, cancellable *Cancellable, err **T.GError) bool
	FileMakeDirectoryWithParents            func(f *File, cancellable *Cancellable, err **T.GError) bool
	FileMakeSymbolicLink                    func(f *File, symlinkValue string, cancellable *Cancellable, err **T.GError) bool
	FileMonitor_                            func(f *File, flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor
	FileMonitorDirectory                    func(f *File, flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor
	FileMonitorFile                         func(f *File, flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor
	FileMountEnclosingVolume                func(f *File, flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileMountEnclosingVolumeFinish          func(f *File, result *AsyncResult, err **T.GError) bool
	FileMountMountable                      func(f *File, flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileMountMountableFinish                func(f *File, result *AsyncResult, err **T.GError) *File
	FileMove                                func(f *File, destination *File, flags FileCopyFlags, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, err **T.GError) bool
	FileOpenReadwrite                       func(f *File, cancellable *Cancellable, err **T.GError) *FileIOStream
	FileOpenReadwriteAsync                  func(f *File, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileOpenReadwriteFinish                 func(f *File, res *AsyncResult, err **T.GError) *FileIOStream
	FilePollMountable                       func(f *File, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FilePollMountableFinish                 func(f *File, result *AsyncResult, err **T.GError) bool
	FileQueryDefaultHandler                 func(f *File, cancellable *Cancellable, err **T.GError) *AppInfo
	FileQueryExists                         func(f *File, cancellable *Cancellable) bool
	FileQueryFilesystemInfo                 func(f *File, attributes string, cancellable *Cancellable, err **T.GError) *FileInfo
	FileQueryFilesystemInfoAsync            func(f *File, attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileQueryFilesystemInfoFinish           func(f *File, res *AsyncResult, err **T.GError) *FileInfo
	FileQueryFileType                       func(f *File, flags FileQueryInfoFlags, cancellable *Cancellable) FileType
	FileQueryInfo                           func(f *File, attributes string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) *FileInfo
	FileQueryInfoAsync                      func(f *File, attributes string, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileQueryInfoFinish                     func(f *File, res *AsyncResult, err **T.GError) *FileInfo
	FileQuerySettableAttributes             func(f *File, cancellable *Cancellable, err **T.GError) *FileAttributeInfoList
	FileQueryWritableNamespaces             func(f *File, cancellable *Cancellable, err **T.GError) *FileAttributeInfoList
	FileRead                                func(f *File, cancellable *Cancellable, err **T.GError) *FileInputStream
	FileReadAsync                           func(f *File, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileReadFinish                          func(f *File, res *AsyncResult, err **T.GError) *FileInputStream
	FileReplace                             func(f *File, etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream
	FileReplaceAsync                        func(f *File, etag string, makeBackup bool, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileReplaceContents                     func(f *File, contents string, length T.Gsize, etag string, makeBackup bool, flags FileCreateFlags, newEtag **T.Char, cancellable *Cancellable, err **T.GError) bool
	FileReplaceContentsAsync                func(f *File, contents string, length T.Gsize, etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileReplaceContentsFinish               func(f *File, res *AsyncResult, newEtag **T.Char, err **T.GError) bool
	FileReplaceFinish                       func(f *File, res *AsyncResult, err **T.GError) *FileOutputStream
	FileReplaceReadwrite                    func(f *File, etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileIOStream
	FileReplaceReadwriteAsync               func(f *File, etag string, makeBackup bool, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileReplaceReadwriteFinish              func(f *File, res *AsyncResult, err **T.GError) *FileIOStream
	FileResolveRelativePath                 func(f *File, relativePath string) *File
	FileSetAttribute                        func(f *File, attribute string, typ FileAttributeType, valueP T.Gpointer, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool
	FileSetAttributeByteString              func(f *File, attribute, value string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool
	FileSetAttributeInt32                   func(f *File, attribute string, value T.GInt32, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool
	FileSetAttributeInt64                   func(f *File, attribute string, value int64, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool
	FileSetAttributesAsync                  func(f *File, info *FileInfo, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileSetAttributesFinish                 func(f *File, result *AsyncResult, info **FileInfo, err **T.GError) bool
	FileSetAttributesFromInfo               func(f *File, info *FileInfo, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool
	FileSetAttributeString                  func(f *File, attribute, value string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool
	FileSetAttributeUint32                  func(f *File, attribute string, value T.GUint32, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool
	FileSetAttributeUint64                  func(f *File, attribute string, value uint64, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool
	FileSetDisplayName                      func(f *File, displayName string, cancellable *Cancellable, err **T.GError) *File
	FileSetDisplayNameAsync                 func(f *File, displayName string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileSetDisplayNameFinish                func(f *File, res *AsyncResult, err **T.GError) *File
	FileStartMountable                      func(f *File, flags DriveStartFlags, startOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileStartMountableFinish                func(f *File, result *AsyncResult, err **T.GError) bool
	FileStopMountable                       func(f *File, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileStopMountableFinish                 func(f *File, result *AsyncResult, err **T.GError) bool
	FileSupportsThreadContexts              func(f *File) bool
	FileTrash                               func(f *File, cancellable *Cancellable, err **T.GError) bool
	FileUnmountMountable                    func(f *File, flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileUnmountMountableFinish              func(f *File, result *AsyncResult, err **T.GError) bool
	FileUnmountMountableWithOperation       func(f *File, flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileUnmountMountableWithOperationFinish func(f *File, result *AsyncResult, err **T.GError) bool
)

func (f *File) AppendTo(flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream {
	return FileAppendTo(f, flags, cancellable, err)
}
func (f *File) AppendToAsync(flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileAppendToAsync(f, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) AppendToFinish(res *AsyncResult, err **T.GError) *FileOutputStream {
	return FileAppendToFinish(f, res, err)
}
func (f *File) Copy(destination *File, flags FileCopyFlags, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, err **T.GError) bool {
	return FileCopy(f, destination, flags, cancellable, progressCallback, progressCallbackData, err)
}
func (f *File) CopyAsync(destination *File, flags FileCopyFlags, ioPriority int, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, callback AsyncReadyCallback, userData T.Gpointer) {
	FileCopyAsync(f, destination, flags, ioPriority, cancellable, progressCallback, progressCallbackData, callback, userData)
}
func (f *File) CopyAttributes(destination *File, flags FileCopyFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileCopyAttributes(f, destination, flags, cancellable, err)
}
func (f *File) CopyFinish(res *AsyncResult, err **T.GError) bool {
	return FileCopyFinish(f, res, err)
}
func (f *File) Create(flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream {
	return FileCreate(f, flags, cancellable, err)
}
func (f *File) CreateAsync(flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileCreateAsync(f, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) CreateFinish(res *AsyncResult, err **T.GError) *FileOutputStream {
	return FileCreateFinish(f, res, err)
}
func (f *File) CreateReadwrite(flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileIOStream {
	return FileCreateReadwrite(f, flags, cancellable, err)
}
func (f *File) CreateReadwriteAsync(flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileCreateReadwriteAsync(f, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) CreateReadwriteFinish(res *AsyncResult, err **T.GError) *FileIOStream {
	return FileCreateReadwriteFinish(f, res, err)
}
func (f *File) Delete(cancellable *Cancellable, err **T.GError) bool {
	return FileDelete(f, cancellable, err)
}
func (f *File) Dup() *File { return FileDup(f) }
func (f *File) EjectMountable(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileEjectMountable(f, flags, cancellable, callback, userData)
}
func (f *File) EjectMountableFinish(result *AsyncResult, err **T.GError) bool {
	return FileEjectMountableFinish(f, result, err)
}
func (f *File) EjectMountableWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileEjectMountableWithOperation(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) EjectMountableWithOperationFinish(result *AsyncResult, err **T.GError) bool {
	return FileEjectMountableWithOperationFinish(f, result, err)
}
func (f *File) EnumerateChildren(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) *FileEnumerator {
	return FileEnumerateChildren(f, attributes, flags, cancellable, err)
}
func (f *File) EnumerateChildrenAsync(attributes string, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileEnumerateChildrenAsync(f, attributes, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) EnumerateChildrenFinish(res *AsyncResult, err **T.GError) *FileEnumerator {
	return FileEnumerateChildrenFinish(f, res, err)
}
func (f *File) Equal(file2 *File) bool { return FileEqual(f, file2) }
func (f *File) FindEnclosingMount(cancellable *Cancellable, err **T.GError) *Mount {
	return FileFindEnclosingMount(f, cancellable, err)
}
func (f *File) FindEnclosingMountAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileFindEnclosingMountAsync(f, ioPriority, cancellable, callback, userData)
}
func (f *File) FindEnclosingMountFinish(res *AsyncResult, err **T.GError) *Mount {
	return FileFindEnclosingMountFinish(f, res, err)
}
func (f *File) GetBasename() string        { return FileGetBasename(f) }
func (f *File) GetChild(name string) *File { return FileGetChild(f, name) }
func (f *File) GetChildForDisplayName(displayName string, err **T.GError) *File {
	return FileGetChildForDisplayName(f, displayName, err)
}
func (f *File) GetParent() *File                        { return FileGetParent(f) }
func (f *File) GetParseName() string                    { return FileGetParseName(f) }
func (f *File) GetPath() string                         { return FileGetPath(f) }
func (f *File) GetRelativePath(descendant *File) string { return FileGetRelativePath(f, descendant) }
func (f *File) GetUri() string                          { return FileGetUri(f) }
func (f *File) GetUriScheme() string                    { return FileGetUriScheme(f) }
func (f *File) HasParent(parent *File) bool             { return FileHasParent(f, parent) }
func (f *File) HasPrefix(prefix *File) bool             { return FileHasPrefix(f, prefix) }
func (f *File) HasUriScheme(uriScheme string) bool      { return FileHasUriScheme(f, uriScheme) }
func (f *File) IsNative() bool                          { return FileIsNative(f) }
func (f *File) LoadContents(cancellable *Cancellable, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) bool {
	return FileLoadContents(f, cancellable, contents, length, etagOut, err)
}
func (f *File) LoadContentsAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileLoadContentsAsync(f, cancellable, callback, userData)
}
func (f *File) LoadContentsFinish(res *AsyncResult, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) bool {
	return FileLoadContentsFinish(f, res, contents, length, etagOut, err)
}
func (f *File) LoadPartialContentsAsync(cancellable *Cancellable, readMoreCallback FileReadMoreCallback, callback AsyncReadyCallback, userData T.Gpointer) {
	FileLoadPartialContentsAsync(f, cancellable, readMoreCallback, callback, userData)
}
func (f *File) LoadPartialContentsFinish(res *AsyncResult, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) bool {
	return FileLoadPartialContentsFinish(f, res, contents, length, etagOut, err)
}
func (f *File) MakeDirectory(cancellable *Cancellable, err **T.GError) bool {
	return FileMakeDirectory(f, cancellable, err)
}
func (f *File) MakeDirectoryWithParents(cancellable *Cancellable, err **T.GError) bool {
	return FileMakeDirectoryWithParents(f, cancellable, err)
}
func (f *File) MakeSymbolicLink(symlinkValue string, cancellable *Cancellable, err **T.GError) bool {
	return FileMakeSymbolicLink(f, symlinkValue, cancellable, err)
}
func (f *File) Monitor(flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor {
	return FileMonitor_(f, flags, cancellable, err)
}
func (f *File) MonitorDirectory(flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor {
	return FileMonitorDirectory(f, flags, cancellable, err)
}
func (f *File) MonitorFile(flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor {
	return FileMonitorFile(f, flags, cancellable, err)
}
func (f *File) MountEnclosingVolume(flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileMountEnclosingVolume(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) MountEnclosingVolumeFinish(result *AsyncResult, err **T.GError) bool {
	return FileMountEnclosingVolumeFinish(f, result, err)
}
func (f *File) MountMountable(flags MountMountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileMountMountable(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) MountMountableFinish(result *AsyncResult, err **T.GError) *File {
	return FileMountMountableFinish(f, result, err)
}
func (f *File) Move(destination *File, flags FileCopyFlags, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, err **T.GError) bool {
	return FileMove(f, destination, flags, cancellable, progressCallback, progressCallbackData, err)
}
func (f *File) OpenReadwrite(cancellable *Cancellable, err **T.GError) *FileIOStream {
	return FileOpenReadwrite(f, cancellable, err)
}
func (f *File) OpenReadwriteAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileOpenReadwriteAsync(f, ioPriority, cancellable, callback, userData)
}
func (f *File) OpenReadwriteFinish(res *AsyncResult, err **T.GError) *FileIOStream {
	return FileOpenReadwriteFinish(f, res, err)
}
func (f *File) PollMountable(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FilePollMountable(f, cancellable, callback, userData)
}
func (f *File) PollMountableFinish(result *AsyncResult, err **T.GError) bool {
	return FilePollMountableFinish(f, result, err)
}
func (f *File) QueryDefaultHandler(cancellable *Cancellable, err **T.GError) *AppInfo {
	return FileQueryDefaultHandler(f, cancellable, err)
}
func (f *File) QueryExists(cancellable *Cancellable) bool {
	return FileQueryExists(f, cancellable)
}
func (f *File) QueryFilesystemInfo(attributes string, cancellable *Cancellable, err **T.GError) *FileInfo {
	return FileQueryFilesystemInfo(f, attributes, cancellable, err)
}
func (f *File) QueryFilesystemInfoAsync(attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileQueryFilesystemInfoAsync(f, attributes, ioPriority, cancellable, callback, userData)
}
func (f *File) QueryFilesystemInfoFinish(res *AsyncResult, err **T.GError) *FileInfo {
	return FileQueryFilesystemInfoFinish(f, res, err)
}
func (f *File) QueryFileType(flags FileQueryInfoFlags, cancellable *Cancellable) FileType {
	return FileQueryFileType(f, flags, cancellable)
}
func (f *File) QueryInfo(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) *FileInfo {
	return FileQueryInfo(f, attributes, flags, cancellable, err)
}
func (f *File) QueryInfoAsync(attributes string, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileQueryInfoAsync(f, attributes, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) QueryInfoFinish(res *AsyncResult, err **T.GError) *FileInfo {
	return FileQueryInfoFinish(f, res, err)
}
func (f *File) QuerySettableAttributes(cancellable *Cancellable, err **T.GError) *FileAttributeInfoList {
	return FileQuerySettableAttributes(f, cancellable, err)
}
func (f *File) QueryWritableNamespaces(cancellable *Cancellable, err **T.GError) *FileAttributeInfoList {
	return FileQueryWritableNamespaces(f, cancellable, err)
}
func (f *File) Read(cancellable *Cancellable, err **T.GError) *FileInputStream {
	return FileRead(f, cancellable, err)
}
func (f *File) ReadAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileReadAsync(f, ioPriority, cancellable, callback, userData)
}
func (f *File) ReadFinish(res *AsyncResult, err **T.GError) *FileInputStream {
	return FileReadFinish(f, res, err)
}
func (f *File) Replace(etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream {
	return FileReplace(f, etag, makeBackup, flags, cancellable, err)
}
func (f *File) ReplaceAsync(etag string, makeBackup bool, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileReplaceAsync(f, etag, makeBackup, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) ReplaceContents(contents string, length T.Gsize, etag string, makeBackup bool, flags FileCreateFlags, newEtag **T.Char, cancellable *Cancellable, err **T.GError) bool {
	return FileReplaceContents(f, contents, length, etag, makeBackup, flags, newEtag, cancellable, err)
}
func (f *File) ReplaceContentsAsync(contents string, length T.Gsize, etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileReplaceContentsAsync(f, contents, length, etag, makeBackup, flags, cancellable, callback, userData)
}
func (f *File) ReplaceContentsFinish(res *AsyncResult, newEtag **T.Char, err **T.GError) bool {
	return FileReplaceContentsFinish(f, res, newEtag, err)
}
func (f *File) ReplaceFinish(res *AsyncResult, err **T.GError) *FileOutputStream {
	return FileReplaceFinish(f, res, err)
}
func (f *File) ReplaceReadwrite(etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileIOStream {
	return FileReplaceReadwrite(f, etag, makeBackup, flags, cancellable, err)
}
func (f *File) ReplaceReadwriteAsync(etag string, makeBackup bool, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileReplaceReadwriteAsync(f, etag, makeBackup, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) ReplaceReadwriteFinish(res *AsyncResult, err **T.GError) *FileIOStream {
	return FileReplaceReadwriteFinish(f, res, err)
}
func (f *File) ResolveRelativePath(relativePath string) *File {
	return FileResolveRelativePath(f, relativePath)
}
func (f *File) SetAttribute(attribute string, typ FileAttributeType, valueP T.Gpointer, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileSetAttribute(f, attribute, typ, valueP, flags, cancellable, err)
}
func (f *File) SetAttributeByteString(attribute, value string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileSetAttributeByteString(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributeInt32(attribute string, value T.GInt32, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileSetAttributeInt32(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributeInt64(attribute string, value int64, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileSetAttributeInt64(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributesAsync(info *FileInfo, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileSetAttributesAsync(f, info, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) SetAttributesFinish(result *AsyncResult, info **FileInfo, err **T.GError) bool {
	return FileSetAttributesFinish(f, result, info, err)
}
func (f *File) SetAttributesFromInfo(info *FileInfo, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileSetAttributesFromInfo(f, info, flags, cancellable, err)
}
func (f *File) SetAttributeString(attribute, value string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileSetAttributeString(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributeUint32(attribute string, value T.GUint32, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileSetAttributeUint32(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributeUint64(attribute string, value uint64, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) bool {
	return FileSetAttributeUint64(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetDisplayName(displayName string, cancellable *Cancellable, err **T.GError) *File {
	return FileSetDisplayName(f, displayName, cancellable, err)
}
func (f *File) SetDisplayNameAsync(displayName string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileSetDisplayNameAsync(f, displayName, ioPriority, cancellable, callback, userData)
}
func (f *File) SetDisplayNameFinish(res *AsyncResult, err **T.GError) *File {
	return FileSetDisplayNameFinish(f, res, err)
}
func (f *File) StartMountable(flags DriveStartFlags, startOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileStartMountable(f, flags, startOperation, cancellable, callback, userData)
}
func (f *File) StartMountableFinish(result *AsyncResult, err **T.GError) bool {
	return FileStartMountableFinish(f, result, err)
}
func (f *File) StopMountable(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileStopMountable(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) StopMountableFinish(result *AsyncResult, err **T.GError) bool {
	return FileStopMountableFinish(f, result, err)
}
func (f *File) SupportsThreadContexts() bool { return FileSupportsThreadContexts(f) }
func (f *File) Trash(cancellable *Cancellable, err **T.GError) bool {
	return FileTrash(f, cancellable, err)
}
func (f *File) UnmountMountable(flags MountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileUnmountMountable(f, flags, cancellable, callback, userData)
}
func (f *File) UnmountMountableFinish(result *AsyncResult, err **T.GError) bool {
	return FileUnmountMountableFinish(f, result, err)
}
func (f *File) UnmountMountableWithOperation(flags MountUnmountFlags, mountOperation *MountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileUnmountMountableWithOperation(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) UnmountMountableWithOperationFinish(result *AsyncResult, err **T.GError) bool {
	return FileUnmountMountableWithOperationFinish(f, result, err)
}

type FileAttributeInfo struct {
	Name  *T.Char
	Type  FileAttributeType
	Flags FileAttributeInfoFlags
}

type FileAttributeInfoFlags Enum

const (
	FILE_ATTRIBUTE_INFO_COPY_WITH_FILE FileAttributeInfoFlags = 1 << iota
	FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED
	FILE_ATTRIBUTE_INFO_NONE FileAttributeInfoFlags = 0
)

var FileAttributeInfoFlagsGetType func() O.Type

type FileAttributeInfoList struct {
	Infos  *FileAttributeInfo
	NInfos int
}

var (
	FileAttributeInfoListGetType func() O.Type
	FileAttributeInfoListNew     func() *FileAttributeInfoList

	FileAttributeInfoListAdd    func(f *FileAttributeInfoList, name string, typ FileAttributeType, flags FileAttributeInfoFlags)
	FileAttributeInfoListDup    func(f *FileAttributeInfoList) *FileAttributeInfoList
	FileAttributeInfoListLookup func(f *FileAttributeInfoList, name string) *FileAttributeInfo
	FileAttributeInfoListRef    func(f *FileAttributeInfoList) *FileAttributeInfoList
	FileAttributeInfoListUnref  func(f *FileAttributeInfoList)
)

func (f *FileAttributeInfoList) Add(name string, typ FileAttributeType, flags FileAttributeInfoFlags) {
	FileAttributeInfoListAdd(f, name, typ, flags)
}
func (f *FileAttributeInfoList) Dup() *FileAttributeInfoList { return FileAttributeInfoListDup(f) }
func (f *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	return FileAttributeInfoListLookup(f, name)
}
func (f *FileAttributeInfoList) Ref() *FileAttributeInfoList { return FileAttributeInfoListRef(f) }
func (f *FileAttributeInfoList) Unref()                      { FileAttributeInfoListUnref(f) }

type FileAttributeMatcher struct{}

var (
	FileAttributeMatcherGetType func() O.Type
	FileAttributeMatcherNew     func(attributes string) *FileAttributeMatcher

	FileAttributeMatcherEnumerateNamespace func(f *FileAttributeMatcher, ns string) bool
	FileAttributeMatcherEnumerateNext      func(f *FileAttributeMatcher) string
	FileAttributeMatcherMatches            func(f *FileAttributeMatcher, attribute string) bool
	FileAttributeMatcherMatchesOnly        func(f *FileAttributeMatcher, attribute string) bool
	FileAttributeMatcherRef                func(f *FileAttributeMatcher) *FileAttributeMatcher
	FileAttributeMatcherUnref              func(f *FileAttributeMatcher)
)

func (f *FileAttributeMatcher) EnumerateNamespace(ns string) bool {
	return FileAttributeMatcherEnumerateNamespace(f, ns)
}
func (f *FileAttributeMatcher) EnumerateNext() string { return FileAttributeMatcherEnumerateNext(f) }
func (f *FileAttributeMatcher) Matches(attribute string) bool {
	return FileAttributeMatcherMatches(f, attribute)
}
func (f *FileAttributeMatcher) MatchesOnly(attribute string) bool {
	return FileAttributeMatcherMatchesOnly(f, attribute)
}
func (f *FileAttributeMatcher) Ref() *FileAttributeMatcher { return FileAttributeMatcherRef(f) }
func (f *FileAttributeMatcher) Unref()                     { FileAttributeMatcherUnref(f) }

type FileAttributeType Enum

const (
	FILE_ATTRIBUTE_TYPE_INVALID FileAttributeType = iota
	FILE_ATTRIBUTE_TYPE_STRING
	FILE_ATTRIBUTE_TYPE_BYTE_STRING
	FILE_ATTRIBUTE_TYPE_BOOLEAN
	FILE_ATTRIBUTE_TYPE_UINT32
	FILE_ATTRIBUTE_TYPE_INT32
	FILE_ATTRIBUTE_TYPE_UINT64
	FILE_ATTRIBUTE_TYPE_INT64
	FILE_ATTRIBUTE_TYPE_OBJECT
	FILE_ATTRIBUTE_TYPE_STRINGV
)

var FileAttributeTypeGetType func() O.Type

type FileAttributeStatus Enum

const (
	FILE_ATTRIBUTE_STATUS_UNSET FileAttributeStatus = iota
	FILE_ATTRIBUTE_STATUS_SET
	FILE_ATTRIBUTE_STATUS_ERROR_SETTING
)

var FileAttributeStatusGetType func() O.Type

type FileCreateFlags Enum

const (
	FILE_CREATE_PRIVATE FileCreateFlags = 1 << iota
	FILE_CREATE_REPLACE_DESTINATION
	FILE_CREATE_NONE FileCreateFlags = 0
)

var FileCreateFlagsGetType func() O.Type

type FileCopyFlags Enum

const (
	FILE_COPY_OVERWRITE FileCopyFlags = 1 << iota
	FILE_COPY_BACKUP
	FILE_COPY_NOFOLLOW_SYMLINKS
	FILE_COPY_ALL_METADATA
	FILE_COPY_NO_FALLBACK_FOR_MOVE
	FILE_COPY_TARGET_DEFAULT_PERMS
	FILE_COPY_NONE FileCopyFlags = 0
)

var FileCopyFlagsGetType func() O.Type

type FileEnumerator struct {
	Parent O.Object
	_      *struct{}
}

var (
	FileEnumeratorGetType func() O.Type

	FileEnumeratorNextFile        func(f *FileEnumerator, cancellable *Cancellable, err **T.GError) *FileInfo
	FileEnumeratorClose           func(f *FileEnumerator, cancellable *Cancellable, err **T.GError) bool
	FileEnumeratorNextFilesAsync  func(f *FileEnumerator, numFiles, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileEnumeratorNextFilesFinish func(f *FileEnumerator, result *AsyncResult, err **T.GError) *T.GList
	FileEnumeratorCloseAsync      func(f *FileEnumerator, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileEnumeratorCloseFinish     func(f *FileEnumerator, result *AsyncResult, err **T.GError) bool
	FileEnumeratorIsClosed        func(f *FileEnumerator) bool
	FileEnumeratorHasPending      func(f *FileEnumerator) bool
	FileEnumeratorSetPending      func(f *FileEnumerator, pending bool)
	FileEnumeratorGetContainer    func(f *FileEnumerator) *File
)

func (f *FileEnumerator) NextFile(cancellable *Cancellable, err **T.GError) *FileInfo {
	return FileEnumeratorNextFile(f, cancellable, err)
}
func (f *FileEnumerator) Close(cancellable *Cancellable, err **T.GError) bool {
	return FileEnumeratorClose(f, cancellable, err)
}
func (f *FileEnumerator) NextFilesAsync(numFiles, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileEnumeratorNextFilesAsync(f, numFiles, ioPriority, cancellable, callback, userData)
}
func (f *FileEnumerator) NextFilesFinish(result *AsyncResult, err **T.GError) *T.GList {
	return FileEnumeratorNextFilesFinish(f, result, err)
}
func (f *FileEnumerator) CloseAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileEnumeratorCloseAsync(f, ioPriority, cancellable, callback, userData)
}
func (f *FileEnumerator) CloseFinish(result *AsyncResult, err **T.GError) bool {
	return FileEnumeratorCloseFinish(f, result, err)
}
func (f *FileEnumerator) IsClosed() bool          { return FileEnumeratorIsClosed(f) }
func (f *FileEnumerator) HasPending() bool        { return FileEnumeratorHasPending(f) }
func (f *FileEnumerator) SetPending(pending bool) { FileEnumeratorSetPending(f, pending) }
func (f *FileEnumerator) GetContainer() *File     { return FileEnumeratorGetContainer(f) }

type FileIcon struct{}

var (
	FileIconGetType func() O.Type
	FileIconNew     func(file *File) *Icon

	FileIconGetFile func(f *FileIcon) *File
)

func (f *FileIcon) GetFile() *File { return FileIconGetFile(f) }

type FileInfo struct{}

var (
	FileInfoGetType func() O.Type
	FileInfoNew     func() *FileInfo

	FileInfoDup      func(other *FileInfo) *FileInfo
	FileInfoCopyInto func(srcInfo, destInfo *FileInfo)

	FileInfoClearStatus            func(f *FileInfo)
	FileInfoGetAttributeAsString   func(f *FileInfo, attribute string) string
	FileInfoGetAttributeBoolean    func(f *FileInfo, attribute string) bool
	FileInfoGetAttributeByteString func(f *FileInfo, attribute string) string
	FileInfoGetAttributeData       func(f *FileInfo, attribute string, typ *FileAttributeType, valuePp *T.Gpointer, status *FileAttributeStatus) bool
	FileInfoGetAttributeInt32      func(f *FileInfo, attribute string) T.GInt32
	FileInfoGetAttributeInt64      func(f *FileInfo, attribute string) int64
	FileInfoGetAttributeObject     func(f *FileInfo, attribute string) *O.Object
	FileInfoGetAttributeStatus     func(f *FileInfo, attribute string) FileAttributeStatus
	FileInfoGetAttributeString     func(f *FileInfo, attribute string) string
	FileInfoGetAttributeStringv    func(f *FileInfo, attribute string) []string
	FileInfoGetAttributeType       func(f *FileInfo, attribute string) FileAttributeType
	FileInfoGetAttributeUint32     func(f *FileInfo, attribute string) T.GUint32
	FileInfoGetAttributeUint64     func(f *FileInfo, attribute string) uint64
	FileInfoGetContentType         func(f *FileInfo) string
	FileInfoGetDisplayName         func(f *FileInfo) string
	FileInfoGetEditName            func(f *FileInfo) string
	FileInfoGetEtag                func(f *FileInfo) string
	FileInfoGetFileType            func(f *FileInfo) FileType
	FileInfoGetIcon                func(f *FileInfo) *Icon
	FileInfoGetIsBackup            func(f *FileInfo) bool
	FileInfoGetIsHidden            func(f *FileInfo) bool
	FileInfoGetIsSymlink           func(f *FileInfo) bool
	FileInfoGetModificationTime    func(f *FileInfo, result *T.GTimeVal)
	FileInfoGetName                func(f *FileInfo) string
	FileInfoGetSize                func(f *FileInfo) T.Goffset
	FileInfoGetSortOrder           func(f *FileInfo) T.GInt32
	FileInfoGetSymlinkTarget       func(f *FileInfo) string
	FileInfoHasAttribute           func(f *FileInfo, attribute string) bool
	FileInfoHasNamespace           func(f *FileInfo, nameSpace string) bool
	FileInfoListAttributes         func(f *FileInfo, nameSpace string) []string
	FileInfoRemoveAttribute        func(f *FileInfo, attribute string)
	FileInfoSetAttribute           func(f *FileInfo, attribute string, typ FileAttributeType, valueP T.Gpointer)
	FileInfoSetAttributeBoolean    func(f *FileInfo, attribute string, attrValue bool)
	FileInfoSetAttributeByteString func(f *FileInfo, attribute, attrValue string)
	FileInfoSetAttributeInt32      func(f *FileInfo, attribute string, attrValue T.GInt32)
	FileInfoSetAttributeInt64      func(f *FileInfo, attribute string, attrValue int64)
	FileInfoSetAttributeMask       func(f *FileInfo, mask *FileAttributeMatcher)
	FileInfoSetAttributeObject     func(f *FileInfo, attribute string, attrValue *O.Object)
	FileInfoSetAttributeStatus     func(f *FileInfo, attribute string, status FileAttributeStatus) bool
	FileInfoSetAttributeString     func(f *FileInfo, attribute, attrValue string)
	FileInfoSetAttributeStringv    func(f *FileInfo, attribute string, attrValue **T.Char)
	FileInfoSetAttributeUint32     func(f *FileInfo, attribute string, attrValue T.GUint32)
	FileInfoSetAttributeUint64     func(f *FileInfo, attribute string, attrValue uint64)
	FileInfoSetContentType         func(f *FileInfo, contentType string)
	FileInfoSetDisplayName         func(f *FileInfo, displayName string)
	FileInfoSetEditName            func(f *FileInfo, editName string)
	FileInfoSetFileType            func(f *FileInfo, typ FileType)
	FileInfoSetIcon                func(f *FileInfo, icon *Icon)
	FileInfoSetIsHidden            func(f *FileInfo, isHidden bool)
	FileInfoSetIsSymlink           func(f *FileInfo, isSymlink bool)
	FileInfoSetModificationTime    func(f *FileInfo, mtime *T.GTimeVal)
	FileInfoSetName                func(f *FileInfo, name string)
	FileInfoSetSize                func(f *FileInfo, size T.Goffset)
	FileInfoSetSortOrder           func(f *FileInfo, sortOrder T.GInt32)
	FileInfoSetSymlinkTarget       func(f *FileInfo, symlinkTarget string)
	FileInfoUnsetAttributeMask     func(f *FileInfo)
)

func (f *FileInfo) ClearStatus() { FileInfoClearStatus(f) }
func (f *FileInfo) GetAttributeAsString(attribute string) string {
	return FileInfoGetAttributeAsString(f, attribute)
}
func (f *FileInfo) GetAttributeBoolean(attribute string) bool {
	return FileInfoGetAttributeBoolean(f, attribute)
}
func (f *FileInfo) GetAttributeByteString(attribute string) string {
	return FileInfoGetAttributeByteString(f, attribute)
}
func (f *FileInfo) GetAttributeData(attribute string, typ *FileAttributeType, valuePp *T.Gpointer, status *FileAttributeStatus) bool {
	return FileInfoGetAttributeData(f, attribute, typ, valuePp, status)
}
func (f *FileInfo) GetAttributeInt32(attribute string) T.GInt32 {
	return FileInfoGetAttributeInt32(f, attribute)
}
func (f *FileInfo) GetAttributeInt64(attribute string) int64 {
	return FileInfoGetAttributeInt64(f, attribute)
}
func (f *FileInfo) GetAttributeObject(attribute string) *O.Object {
	return FileInfoGetAttributeObject(f, attribute)
}
func (f *FileInfo) GetAttributeStatus(attribute string) FileAttributeStatus {
	return FileInfoGetAttributeStatus(f, attribute)
}
func (f *FileInfo) GetAttributeString(attribute string) string {
	return FileInfoGetAttributeString(f, attribute)
}
func (f *FileInfo) GetAttributeStringv(attribute string) []string {
	return FileInfoGetAttributeStringv(f, attribute)
}
func (f *FileInfo) GetAttributeType(attribute string) FileAttributeType {
	return FileInfoGetAttributeType(f, attribute)
}
func (f *FileInfo) GetAttributeUint32(attribute string) T.GUint32 {
	return FileInfoGetAttributeUint32(f, attribute)
}
func (f *FileInfo) GetAttributeUint64(attribute string) uint64 {
	return FileInfoGetAttributeUint64(f, attribute)
}
func (f *FileInfo) GetContentType() string                 { return FileInfoGetContentType(f) }
func (f *FileInfo) GetDisplayName() string                 { return FileInfoGetDisplayName(f) }
func (f *FileInfo) GetEditName() string                    { return FileInfoGetEditName(f) }
func (f *FileInfo) GetEtag() string                        { return FileInfoGetEtag(f) }
func (f *FileInfo) GetFileType() FileType                  { return FileInfoGetFileType(f) }
func (f *FileInfo) GetIcon() *Icon                         { return FileInfoGetIcon(f) }
func (f *FileInfo) GetIsBackup() bool                      { return FileInfoGetIsBackup(f) }
func (f *FileInfo) GetIsHidden() bool                      { return FileInfoGetIsHidden(f) }
func (f *FileInfo) GetIsSymlink() bool                     { return FileInfoGetIsSymlink(f) }
func (f *FileInfo) GetModificationTime(result *T.GTimeVal) { FileInfoGetModificationTime(f, result) }
func (f *FileInfo) GetName() string                        { return FileInfoGetName(f) }
func (f *FileInfo) GetSize() T.Goffset                     { return FileInfoGetSize(f) }
func (f *FileInfo) GetSortOrder() T.GInt32                 { return FileInfoGetSortOrder(f) }
func (f *FileInfo) GetSymlinkTarget() string               { return FileInfoGetSymlinkTarget(f) }
func (f *FileInfo) HasAttribute(attribute string) bool {
	return FileInfoHasAttribute(f, attribute)
}
func (f *FileInfo) HasNamespace(nameSpace string) bool {
	return FileInfoHasNamespace(f, nameSpace)
}
func (f *FileInfo) ListAttributes(nameSpace string) []string {
	return FileInfoListAttributes(f, nameSpace)
}
func (f *FileInfo) RemoveAttribute(attribute string) { FileInfoRemoveAttribute(f, attribute) }
func (f *FileInfo) SetAttribute(attribute string, typ FileAttributeType, valueP T.Gpointer) {
	FileInfoSetAttribute(f, attribute, typ, valueP)
}
func (f *FileInfo) SetAttributeBoolean(attribute string, attrValue bool) {
	FileInfoSetAttributeBoolean(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeByteString(attribute, attrValue string) {
	FileInfoSetAttributeByteString(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeInt32(attribute string, attrValue T.GInt32) {
	FileInfoSetAttributeInt32(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeInt64(attribute string, attrValue int64) {
	FileInfoSetAttributeInt64(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeMask(mask *FileAttributeMatcher) { FileInfoSetAttributeMask(f, mask) }
func (f *FileInfo) SetAttributeObject(attribute string, attrValue *O.Object) {
	FileInfoSetAttributeObject(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeStatus(attribute string, status FileAttributeStatus) bool {
	return FileInfoSetAttributeStatus(f, attribute, status)
}
func (f *FileInfo) SetAttributeString(attribute, attrValue string) {
	FileInfoSetAttributeString(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeStringv(attribute string, attrValue **T.Char) {
	FileInfoSetAttributeStringv(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeUint32(attribute string, attrValue T.GUint32) {
	FileInfoSetAttributeUint32(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeUint64(attribute string, attrValue uint64) {
	FileInfoSetAttributeUint64(f, attribute, attrValue)
}
func (f *FileInfo) SetContentType(contentType string)     { FileInfoSetContentType(f, contentType) }
func (f *FileInfo) SetDisplayName(displayName string)     { FileInfoSetDisplayName(f, displayName) }
func (f *FileInfo) SetEditName(editName string)           { FileInfoSetEditName(f, editName) }
func (f *FileInfo) SetFileType(typ FileType)              { FileInfoSetFileType(f, typ) }
func (f *FileInfo) SetIcon(icon *Icon)                    { FileInfoSetIcon(f, icon) }
func (f *FileInfo) SetIsHidden(isHidden bool)             { FileInfoSetIsHidden(f, isHidden) }
func (f *FileInfo) SetIsSymlink(isSymlink bool)           { FileInfoSetIsSymlink(f, isSymlink) }
func (f *FileInfo) SetModificationTime(mtime *T.GTimeVal) { FileInfoSetModificationTime(f, mtime) }
func (f *FileInfo) SetName(name string)                   { FileInfoSetName(f, name) }
func (f *FileInfo) SetSize(size T.Goffset)                { FileInfoSetSize(f, size) }
func (f *FileInfo) SetSortOrder(sortOrder T.GInt32)       { FileInfoSetSortOrder(f, sortOrder) }
func (f *FileInfo) SetSymlinkTarget(symlinkTarget string) { FileInfoSetSymlinkTarget(f, symlinkTarget) }
func (f *FileInfo) UnsetAttributeMask()                   { FileInfoUnsetAttributeMask(f) }

type FileInputStream struct {
	Parent InputStream
	_      *struct{}
}

var (
	FileInputStreamGetType func() O.Type

	FileInputStreamQueryInfo       func(f *FileInputStream, attributes string, cancellable *Cancellable, err **T.GError) *FileInfo
	FileInputStreamQueryInfoAsync  func(f *FileInputStream, attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileInputStreamQueryInfoFinish func(f *FileInputStream, result *AsyncResult, err **T.GError) *FileInfo
)

func (f *FileInputStream) QueryInfo(attributes string, cancellable *Cancellable, err **T.GError) *FileInfo {
	return FileInputStreamQueryInfo(f, attributes, cancellable, err)
}
func (f *FileInputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileInputStreamQueryInfoAsync(f, attributes, ioPriority, cancellable, callback, userData)
}
func (f *FileInputStream) QueryInfoFinish(result *AsyncResult, err **T.GError) *FileInfo {
	return FileInputStreamQueryInfoFinish(f, result, err)
}

type FileIOStream struct {
	Parent IOStream
	_      *struct{}
}

var (
	FileIoStreamGetType func() O.Type

	FileIoStreamGetEtag         func(f *FileIOStream) string
	FileIoStreamQueryInfo       func(f *FileIOStream, attributes string, cancellable *Cancellable, err **T.GError) *FileInfo
	FileIoStreamQueryInfoAsync  func(f *FileIOStream, attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileIoStreamQueryInfoFinish func(f *FileIOStream, result *AsyncResult, err **T.GError) *FileInfo
)

func (f *FileIOStream) GetEtag() string { return FileIoStreamGetEtag(f) }
func (f *FileIOStream) QueryInfo(attributes string, cancellable *Cancellable, err **T.GError) *FileInfo {
	return FileIoStreamQueryInfo(f, attributes, cancellable, err)
}
func (f *FileIOStream) QueryInfoAsync(attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileIoStreamQueryInfoAsync(f, attributes, ioPriority, cancellable, callback, userData)
}
func (f *FileIOStream) QueryInfoFinish(result *AsyncResult, err **T.GError) *FileInfo {
	return FileIoStreamQueryInfoFinish(f, result, err)
}

type FileMonitor struct {
	Parent O.Object
	_      *struct{}
}

var (
	FileMonitorGetType func() O.Type

	FileMonitorCancel       func(f *FileMonitor) bool
	FileMonitorEmitEvent    func(f *FileMonitor, child, otherFile *File, eventType FileMonitorEvent)
	FileMonitorIsCancelled  func(f *FileMonitor) bool
	FileMonitorSetRateLimit func(f *FileMonitor, limitMsecs int)
)

func (f *FileMonitor) Cancel() bool { return FileMonitorCancel(f) }
func (f *FileMonitor) EmitEvent(child, otherFile *File, eventType FileMonitorEvent) {
	FileMonitorEmitEvent(f, child, otherFile, eventType)
}
func (f *FileMonitor) IsCancelled() bool           { return FileMonitorIsCancelled(f) }
func (f *FileMonitor) SetRateLimit(limitMsecs int) { FileMonitorSetRateLimit(f, limitMsecs) }

type FileMonitorEvent Enum

const (
	FILE_MONITOR_EVENT_CHANGED FileMonitorEvent = iota
	FILE_MONITOR_EVENT_CHANGES_DONE_HINT
	FILE_MONITOR_EVENT_DELETED
	FILE_MONITOR_EVENT_CREATED
	FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED
	FILE_MONITOR_EVENT_PRE_UNMOUNT
	FILE_MONITOR_EVENT_UNMOUNTED
	FILE_MONITOR_EVENT_MOVED
)

var FileMonitorEventGetType func() O.Type

type FileMonitorFlags Enum

const (
	FILE_MONITOR_WATCH_MOUNTS FileMonitorFlags = 1 << iota
	FILE_MONITOR_SEND_MOVED
	FILE_MONITOR_NONE FileMonitorFlags = 0
)

var FileMonitorFlagsGetType func() O.Type

type FilenameCompleter struct{}

var (
	FilenameCompleterGetType func() O.Type
	FilenameCompleterNew     func() *FilenameCompleter

	FilenameCompleterGetCompletions      func(f *FilenameCompleter, initialText string) []string
	FilenameCompleterGetCompletionSuffix func(f *FilenameCompleter, initialText string) string
	FilenameCompleterSetDirsOnly         func(f *FilenameCompleter, dirsOnly bool)
)

func (f *FilenameCompleter) GetCompletions(initialText string) []string {
	return FilenameCompleterGetCompletions(f, initialText)
}
func (f *FilenameCompleter) GetCompletionSuffix(initialText string) string {
	return FilenameCompleterGetCompletionSuffix(f, initialText)
}
func (f *FilenameCompleter) SetDirsOnly(dirsOnly bool) {
	FilenameCompleterSetDirsOnly(f, dirsOnly)
}

type FileOutputStream struct {
	Parent OutputStream
	_      *struct{}
}

var (
	FileOutputStreamGetType func() O.Type

	FileOutputStreamGetEtag         func(f *FileOutputStream) string
	FileOutputStreamQueryInfo       func(f *FileOutputStream, attributes string, cancellable *Cancellable, err **T.GError) *FileInfo
	FileOutputStreamQueryInfoAsync  func(f *FileOutputStream, attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	FileOutputStreamQueryInfoFinish func(f *FileOutputStream, result *AsyncResult, err **T.GError) *FileInfo
)

func (f *FileOutputStream) GetEtag() string { return FileOutputStreamGetEtag(f) }
func (f *FileOutputStream) QueryInfo(attributes string, cancellable *Cancellable, err **T.GError) *FileInfo {
	return FileOutputStreamQueryInfo(f, attributes, cancellable, err)
}
func (f *FileOutputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	FileOutputStreamQueryInfoAsync(f, attributes, ioPriority, cancellable, callback, userData)
}
func (f *FileOutputStream) QueryInfoFinish(result *AsyncResult, err **T.GError) *FileInfo {
	return FileOutputStreamQueryInfoFinish(f, result, err)
}

type FileProgressCallback func(
	currentNumBytes, totalNumBytes T.Goffset, userData T.Gpointer)

type FileQueryInfoFlags Enum

const (
	FILE_QUERY_INFO_NOFOLLOW_SYMLINKS FileQueryInfoFlags = 1 << iota
	FILE_QUERY_INFO_NONE              FileQueryInfoFlags = 0
)

var FileQueryInfoFlagsGetType func() O.Type

type FileReadMoreCallback func(fileContents string,
	FileSize T.Goffset, callbackData T.Gpointer) bool

var FilesystemPreviewTypeGetType func() O.Type

type FileType Enum

const (
	FILE_TYPE_UNKNOWN FileType = iota
	FILE_TYPE_REGULAR
	FILE_TYPE_DIRECTORY
	FILE_TYPE_SYMBOLIC_LINK
	FILE_TYPE_SPECIAL
	FILE_TYPE_SHORTCUT
	FILE_TYPE_MOUNTABLE
)

var FileTypeGetType func() O.Type

type FilterInputStream struct {
	Parent     InputStream
	BaseStream *InputStream
}

var (
	FilterInputStreamGetType func() O.Type

	FilterInputStreamGetBaseStream      func(f *FilterInputStream) *InputStream
	FilterInputStreamGetCloseBaseStream func(f *FilterInputStream) bool
	FilterInputStreamSetCloseBaseStream func(f *FilterInputStream, closeBase bool)
)

func (f *FilterInputStream) GetBaseStream() *InputStream { return FilterInputStreamGetBaseStream(f) }
func (f *FilterInputStream) GetCloseBaseStream() bool {
	return FilterInputStreamGetCloseBaseStream(f)
}
func (f *FilterInputStream) SetCloseBaseStream(closeBase bool) {
	FilterInputStreamSetCloseBaseStream(f, closeBase)
}

type FilterOutputStream struct {
	Parent     OutputStream
	BaseStream *OutputStream
}

var (
	FilterOutputStreamGetType func() O.Type

	FilterOutputStreamGetBaseStream      func(f *FilterOutputStream) *OutputStream
	FilterOutputStreamGetCloseBaseStream func(f *FilterOutputStream) bool
	FilterOutputStreamSetCloseBaseStream func(f *FilterOutputStream, closeBase bool)
)

func (f *FilterOutputStream) GetBaseStream() *OutputStream {
	return FilterOutputStreamGetBaseStream(f)
}
func (f *FilterOutputStream) GetCloseBaseStream() bool {
	return FilterOutputStreamGetCloseBaseStream(f)
}
func (f *FilterOutputStream) SetCloseBaseStream(closeBase bool) {
	FilterOutputStreamSetCloseBaseStream(f, closeBase)
}
