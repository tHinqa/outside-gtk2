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

	fileAppendTo                            func(f *File, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream
	fileAppendToAsync                       func(f *File, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileAppendToFinish                      func(f *File, res *AsyncResult, err **T.GError) *FileOutputStream
	fileCopy                                func(f, destination *File, flags FileCopyFlags, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, err **T.GError) T.Gboolean
	fileCopyAsync                           func(f, destination *File, flags FileCopyFlags, ioPriority int, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, callback AsyncReadyCallback, userData T.Gpointer)
	fileCopyAttributes                      func(f, destination *File, flags FileCopyFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileCopyFinish                          func(f *File, res *AsyncResult, err **T.GError) T.Gboolean
	fileCreate                              func(f *File, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream
	fileCreateAsync                         func(f *File, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileCreateFinish                        func(f *File, res *AsyncResult, err **T.GError) *FileOutputStream
	fileCreateReadwrite                     func(f *File, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileIOStream
	fileCreateReadwriteAsync                func(f *File, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileCreateReadwriteFinish               func(f *File, res *AsyncResult, err **T.GError) *FileIOStream
	fileDelete                              func(f *File, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileDup                                 func(f *File) *File
	fileEjectMountable                      func(f *File, flags T.GMountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileEjectMountableFinish                func(f *File, result *AsyncResult, err **T.GError) T.Gboolean
	fileEjectMountableWithOperation         func(f *File, flags T.GMountUnmountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileEjectMountableWithOperationFinish   func(f *File, result *AsyncResult, err **T.GError) T.Gboolean
	fileEnumerateChildren                   func(f *File, attributes string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) *FileEnumerator
	fileEnumerateChildrenAsync              func(f *File, attributes string, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileEnumerateChildrenFinish             func(f *File, res *AsyncResult, err **T.GError) *FileEnumerator
	fileEqual                               func(f, file2 *File) T.Gboolean
	fileFindEnclosingMount                  func(f *File, cancellable *Cancellable, err **T.GError) *T.GMount
	fileFindEnclosingMountAsync             func(f *File, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileFindEnclosingMountFinish            func(f *File, res *AsyncResult, err **T.GError) *T.GMount
	fileGetBasename                         func(f *File) string
	fileGetChild                            func(f *File, name string) *File
	fileGetChildForDisplayName              func(f *File, displayName string, err **T.GError) *File
	fileGetParent                           func(f *File) *File
	fileGetParseName                        func(f *File) string
	fileGetPath                             func(f *File) string
	fileGetRelativePath                     func(f *File, descendant *File) string
	fileGetUri                              func(f *File) string
	fileGetUriScheme                        func(f *File) string
	fileHasParent                           func(f, parent *File) T.Gboolean
	fileHasPrefix                           func(f, prefix *File) T.Gboolean
	fileHasUriScheme                        func(f *File, uriScheme string) T.Gboolean
	fileIsNative                            func(f *File) T.Gboolean
	fileLoadContents                        func(f *File, cancellable *Cancellable, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) T.Gboolean
	fileLoadContentsAsync                   func(f *File, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileLoadContentsFinish                  func(f *File, res *AsyncResult, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) T.Gboolean
	fileLoadPartialContentsAsync            func(f *File, cancellable *Cancellable, readMoreCallback FileReadMoreCallback, callback AsyncReadyCallback, userData T.Gpointer)
	fileLoadPartialContentsFinish           func(f *File, res *AsyncResult, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) T.Gboolean
	fileMakeDirectory                       func(f *File, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileMakeDirectoryWithParents            func(f *File, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileMakeSymbolicLink                    func(f *File, symlinkValue string, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileMonitor                             func(f *File, flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor // AMBIGUITY FILE/TYPE
	fileMonitorDirectory                    func(f *File, flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor
	fileMonitorFile                         func(f *File, flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor
	fileMountEnclosingVolume                func(f *File, flags T.GMountMountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileMountEnclosingVolumeFinish          func(f *File, result *AsyncResult, err **T.GError) T.Gboolean
	fileMountMountable                      func(f *File, flags T.GMountMountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileMountMountableFinish                func(f *File, result *AsyncResult, err **T.GError) *File
	fileMove                                func(f *File, destination *File, flags FileCopyFlags, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, err **T.GError) T.Gboolean
	fileOpenReadwrite                       func(f *File, cancellable *Cancellable, err **T.GError) *FileIOStream
	fileOpenReadwriteAsync                  func(f *File, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileOpenReadwriteFinish                 func(f *File, res *AsyncResult, err **T.GError) *FileIOStream
	filePollMountable                       func(f *File, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	filePollMountableFinish                 func(f *File, result *AsyncResult, err **T.GError) T.Gboolean
	fileQueryDefaultHandler                 func(f *File, cancellable *Cancellable, err **T.GError) *AppInfo
	fileQueryExists                         func(f *File, cancellable *Cancellable) T.Gboolean
	fileQueryFilesystemInfo                 func(f *File, attributes string, cancellable *Cancellable, err **T.GError) *FileInfo
	fileQueryFilesystemInfoAsync            func(f *File, attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileQueryFilesystemInfoFinish           func(f *File, res *AsyncResult, err **T.GError) *FileInfo
	fileQueryFileType                       func(f *File, flags FileQueryInfoFlags, cancellable *Cancellable) FileType
	fileQueryInfo                           func(f *File, attributes string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) *FileInfo
	fileQueryInfoAsync                      func(f *File, attributes string, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileQueryInfoFinish                     func(f *File, res *AsyncResult, err **T.GError) *FileInfo
	fileQuerySettableAttributes             func(f *File, cancellable *Cancellable, err **T.GError) *FileAttributeInfoList
	fileQueryWritableNamespaces             func(f *File, cancellable *Cancellable, err **T.GError) *FileAttributeInfoList
	fileRead                                func(f *File, cancellable *Cancellable, err **T.GError) *FileInputStream
	fileReadAsync                           func(f *File, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileReadFinish                          func(f *File, res *AsyncResult, err **T.GError) *FileInputStream
	fileReplace                             func(f *File, etag string, makeBackup T.Gboolean, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream
	fileReplaceAsync                        func(f *File, etag string, makeBackup T.Gboolean, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileReplaceContents                     func(f *File, contents string, length T.Gsize, etag string, makeBackup T.Gboolean, flags FileCreateFlags, newEtag **T.Char, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileReplaceContentsAsync                func(f *File, contents string, length T.Gsize, etag string, makeBackup T.Gboolean, flags FileCreateFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileReplaceContentsFinish               func(f *File, res *AsyncResult, newEtag **T.Char, err **T.GError) T.Gboolean
	fileReplaceFinish                       func(f *File, res *AsyncResult, err **T.GError) *FileOutputStream
	fileReplaceReadwrite                    func(f *File, etag string, makeBackup T.Gboolean, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileIOStream
	fileReplaceReadwriteAsync               func(f *File, etag string, makeBackup T.Gboolean, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileReplaceReadwriteFinish              func(f *File, res *AsyncResult, err **T.GError) *FileIOStream
	fileResolveRelativePath                 func(f *File, relativePath string) *File
	fileSetAttribute                        func(f *File, attribute string, typ FileAttributeType, valueP T.Gpointer, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileSetAttributeByteString              func(f *File, attribute, value string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileSetAttributeInt32                   func(f *File, attribute string, value T.GInt32, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileSetAttributeInt64                   func(f *File, attribute string, value int64, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileSetAttributesAsync                  func(f *File, info *FileInfo, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileSetAttributesFinish                 func(f *File, result *AsyncResult, info **FileInfo, err **T.GError) T.Gboolean
	fileSetAttributesFromInfo               func(f *File, info *FileInfo, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileSetAttributeString                  func(f *File, attribute, value string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileSetAttributeUint32                  func(f *File, attribute string, value T.GUint32, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileSetAttributeUint64                  func(f *File, attribute string, value uint64, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileSetDisplayName                      func(f *File, displayName string, cancellable *Cancellable, err **T.GError) *File
	fileSetDisplayNameAsync                 func(f *File, displayName string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileSetDisplayNameFinish                func(f *File, res *AsyncResult, err **T.GError) *File
	fileStartMountable                      func(f *File, flags DriveStartFlags, startOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileStartMountableFinish                func(f *File, result *AsyncResult, err **T.GError) T.Gboolean
	fileStopMountable                       func(f *File, flags T.GMountUnmountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileStopMountableFinish                 func(f *File, result *AsyncResult, err **T.GError) T.Gboolean
	fileSupportsThreadContexts              func(f *File) T.Gboolean
	fileTrash                               func(f *File, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileUnmountMountable                    func(f *File, flags T.GMountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileUnmountMountableFinish              func(f *File, result *AsyncResult, err **T.GError) T.Gboolean
	fileUnmountMountableWithOperation       func(f *File, flags T.GMountUnmountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileUnmountMountableWithOperationFinish func(f *File, result *AsyncResult, err **T.GError) T.Gboolean
)

func (f *File) AppendTo(flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream {
	return fileAppendTo(f, flags, cancellable, err)
}
func (f *File) AppendToAsync(flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileAppendToAsync(f, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) AppendToFinish(res *AsyncResult, err **T.GError) *FileOutputStream {
	return fileAppendToFinish(f, res, err)
}
func (f *File) Copy(destination *File, flags FileCopyFlags, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, err **T.GError) T.Gboolean {
	return fileCopy(f, destination, flags, cancellable, progressCallback, progressCallbackData, err)
}
func (f *File) CopyAsync(destination *File, flags FileCopyFlags, ioPriority int, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, callback AsyncReadyCallback, userData T.Gpointer) {
	fileCopyAsync(f, destination, flags, ioPriority, cancellable, progressCallback, progressCallbackData, callback, userData)
}
func (f *File) CopyAttributes(destination *File, flags FileCopyFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileCopyAttributes(f, destination, flags, cancellable, err)
}
func (f *File) CopyFinish(res *AsyncResult, err **T.GError) T.Gboolean {
	return fileCopyFinish(f, res, err)
}
func (f *File) Create(flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream {
	return fileCreate(f, flags, cancellable, err)
}
func (f *File) CreateAsync(flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileCreateAsync(f, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) CreateFinish(res *AsyncResult, err **T.GError) *FileOutputStream {
	return fileCreateFinish(f, res, err)
}
func (f *File) CreateReadwrite(flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileIOStream {
	return fileCreateReadwrite(f, flags, cancellable, err)
}
func (f *File) CreateReadwriteAsync(flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileCreateReadwriteAsync(f, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) CreateReadwriteFinish(res *AsyncResult, err **T.GError) *FileIOStream {
	return fileCreateReadwriteFinish(f, res, err)
}
func (f *File) Delete(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileDelete(f, cancellable, err)
}
func (f *File) Dup() *File { return fileDup(f) }
func (f *File) EjectMountable(flags T.GMountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileEjectMountable(f, flags, cancellable, callback, userData)
}
func (f *File) EjectMountableFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return fileEjectMountableFinish(f, result, err)
}
func (f *File) EjectMountableWithOperation(flags T.GMountUnmountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileEjectMountableWithOperation(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) EjectMountableWithOperationFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return fileEjectMountableWithOperationFinish(f, result, err)
}
func (f *File) EnumerateChildren(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) *FileEnumerator {
	return fileEnumerateChildren(f, attributes, flags, cancellable, err)
}
func (f *File) EnumerateChildrenAsync(attributes string, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileEnumerateChildrenAsync(f, attributes, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) EnumerateChildrenFinish(res *AsyncResult, err **T.GError) *FileEnumerator {
	return fileEnumerateChildrenFinish(f, res, err)
}
func (f *File) Equal(file2 *File) T.Gboolean { return fileEqual(f, file2) }
func (f *File) FindEnclosingMount(cancellable *Cancellable, err **T.GError) *T.GMount {
	return fileFindEnclosingMount(f, cancellable, err)
}
func (f *File) FindEnclosingMountAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileFindEnclosingMountAsync(f, ioPriority, cancellable, callback, userData)
}
func (f *File) FindEnclosingMountFinish(res *AsyncResult, err **T.GError) *T.GMount {
	return fileFindEnclosingMountFinish(f, res, err)
}
func (f *File) GetBasename() string        { return fileGetBasename(f) }
func (f *File) GetChild(name string) *File { return fileGetChild(f, name) }
func (f *File) GetChildForDisplayName(displayName string, err **T.GError) *File {
	return fileGetChildForDisplayName(f, displayName, err)
}
func (f *File) GetParent() *File                         { return fileGetParent(f) }
func (f *File) GetParseName() string                     { return fileGetParseName(f) }
func (f *File) GetPath() string                          { return fileGetPath(f) }
func (f *File) GetRelativePath(descendant *File) string  { return fileGetRelativePath(f, descendant) }
func (f *File) GetUri() string                           { return fileGetUri(f) }
func (f *File) GetUriScheme() string                     { return fileGetUriScheme(f) }
func (f *File) HasParent(parent *File) T.Gboolean        { return fileHasParent(f, parent) }
func (f *File) HasPrefix(prefix *File) T.Gboolean        { return fileHasPrefix(f, prefix) }
func (f *File) HasUriScheme(uriScheme string) T.Gboolean { return fileHasUriScheme(f, uriScheme) }
func (f *File) IsNative() T.Gboolean                     { return fileIsNative(f) }
func (f *File) LoadContents(cancellable *Cancellable, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) T.Gboolean {
	return fileLoadContents(f, cancellable, contents, length, etagOut, err)
}
func (f *File) LoadContentsAsync(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileLoadContentsAsync(f, cancellable, callback, userData)
}
func (f *File) LoadContentsFinish(res *AsyncResult, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) T.Gboolean {
	return fileLoadContentsFinish(f, res, contents, length, etagOut, err)
}
func (f *File) LoadPartialContentsAsync(cancellable *Cancellable, readMoreCallback FileReadMoreCallback, callback AsyncReadyCallback, userData T.Gpointer) {
	fileLoadPartialContentsAsync(f, cancellable, readMoreCallback, callback, userData)
}
func (f *File) LoadPartialContentsFinish(res *AsyncResult, contents **T.Char, length *T.Gsize, etagOut **T.Char, err **T.GError) T.Gboolean {
	return fileLoadPartialContentsFinish(f, res, contents, length, etagOut, err)
}
func (f *File) MakeDirectory(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileMakeDirectory(f, cancellable, err)
}
func (f *File) MakeDirectoryWithParents(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileMakeDirectoryWithParents(f, cancellable, err)
}
func (f *File) MakeSymbolicLink(symlinkValue string, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileMakeSymbolicLink(f, symlinkValue, cancellable, err)
}
func (f *File) Monitor(flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor {
	return fileMonitor(f, flags, cancellable, err)
}
func (f *File) MonitorDirectory(flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor {
	return fileMonitorDirectory(f, flags, cancellable, err)
}
func (f *File) MonitorFile(flags FileMonitorFlags, cancellable *Cancellable, err **T.GError) *FileMonitor {
	return fileMonitorFile(f, flags, cancellable, err)
}
func (f *File) MountEnclosingVolume(flags T.GMountMountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileMountEnclosingVolume(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) MountEnclosingVolumeFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return fileMountEnclosingVolumeFinish(f, result, err)
}
func (f *File) MountMountable(flags T.GMountMountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileMountMountable(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) MountMountableFinish(result *AsyncResult, err **T.GError) *File {
	return fileMountMountableFinish(f, result, err)
}
func (f *File) Move(destination *File, flags FileCopyFlags, cancellable *Cancellable, progressCallback FileProgressCallback, progressCallbackData T.Gpointer, err **T.GError) T.Gboolean {
	return fileMove(f, destination, flags, cancellable, progressCallback, progressCallbackData, err)
}
func (f *File) OpenReadwrite(cancellable *Cancellable, err **T.GError) *FileIOStream {
	return fileOpenReadwrite(f, cancellable, err)
}
func (f *File) OpenReadwriteAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileOpenReadwriteAsync(f, ioPriority, cancellable, callback, userData)
}
func (f *File) OpenReadwriteFinish(res *AsyncResult, err **T.GError) *FileIOStream {
	return fileOpenReadwriteFinish(f, res, err)
}
func (f *File) PollMountable(cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	filePollMountable(f, cancellable, callback, userData)
}
func (f *File) PollMountableFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return filePollMountableFinish(f, result, err)
}
func (f *File) QueryDefaultHandler(cancellable *Cancellable, err **T.GError) *AppInfo {
	return fileQueryDefaultHandler(f, cancellable, err)
}
func (f *File) QueryExists(cancellable *Cancellable) T.Gboolean {
	return fileQueryExists(f, cancellable)
}
func (f *File) QueryFilesystemInfo(attributes string, cancellable *Cancellable, err **T.GError) *FileInfo {
	return fileQueryFilesystemInfo(f, attributes, cancellable, err)
}
func (f *File) QueryFilesystemInfoAsync(attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileQueryFilesystemInfoAsync(f, attributes, ioPriority, cancellable, callback, userData)
}
func (f *File) QueryFilesystemInfoFinish(res *AsyncResult, err **T.GError) *FileInfo {
	return fileQueryFilesystemInfoFinish(f, res, err)
}
func (f *File) QueryFileType(flags FileQueryInfoFlags, cancellable *Cancellable) FileType {
	return fileQueryFileType(f, flags, cancellable)
}
func (f *File) QueryInfo(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) *FileInfo {
	return fileQueryInfo(f, attributes, flags, cancellable, err)
}
func (f *File) QueryInfoAsync(attributes string, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileQueryInfoAsync(f, attributes, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) QueryInfoFinish(res *AsyncResult, err **T.GError) *FileInfo {
	return fileQueryInfoFinish(f, res, err)
}
func (f *File) QuerySettableAttributes(cancellable *Cancellable, err **T.GError) *FileAttributeInfoList {
	return fileQuerySettableAttributes(f, cancellable, err)
}
func (f *File) QueryWritableNamespaces(cancellable *Cancellable, err **T.GError) *FileAttributeInfoList {
	return fileQueryWritableNamespaces(f, cancellable, err)
}
func (f *File) Read(cancellable *Cancellable, err **T.GError) *FileInputStream {
	return fileRead(f, cancellable, err)
}
func (f *File) ReadAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileReadAsync(f, ioPriority, cancellable, callback, userData)
}
func (f *File) ReadFinish(res *AsyncResult, err **T.GError) *FileInputStream {
	return fileReadFinish(f, res, err)
}
func (f *File) Replace(etag string, makeBackup T.Gboolean, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileOutputStream {
	return fileReplace(f, etag, makeBackup, flags, cancellable, err)
}
func (f *File) ReplaceAsync(etag string, makeBackup T.Gboolean, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileReplaceAsync(f, etag, makeBackup, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) ReplaceContents(contents string, length T.Gsize, etag string, makeBackup T.Gboolean, flags FileCreateFlags, newEtag **T.Char, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileReplaceContents(f, contents, length, etag, makeBackup, flags, newEtag, cancellable, err)
}
func (f *File) ReplaceContentsAsync(contents string, length T.Gsize, etag string, makeBackup T.Gboolean, flags FileCreateFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileReplaceContentsAsync(f, contents, length, etag, makeBackup, flags, cancellable, callback, userData)
}
func (f *File) ReplaceContentsFinish(res *AsyncResult, newEtag **T.Char, err **T.GError) T.Gboolean {
	return fileReplaceContentsFinish(f, res, newEtag, err)
}
func (f *File) ReplaceFinish(res *AsyncResult, err **T.GError) *FileOutputStream {
	return fileReplaceFinish(f, res, err)
}
func (f *File) ReplaceReadwrite(etag string, makeBackup T.Gboolean, flags FileCreateFlags, cancellable *Cancellable, err **T.GError) *FileIOStream {
	return fileReplaceReadwrite(f, etag, makeBackup, flags, cancellable, err)
}
func (f *File) ReplaceReadwriteAsync(etag string, makeBackup T.Gboolean, flags FileCreateFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileReplaceReadwriteAsync(f, etag, makeBackup, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) ReplaceReadwriteFinish(res *AsyncResult, err **T.GError) *FileIOStream {
	return fileReplaceReadwriteFinish(f, res, err)
}
func (f *File) ResolveRelativePath(relativePath string) *File {
	return fileResolveRelativePath(f, relativePath)
}
func (f *File) SetAttribute(attribute string, typ FileAttributeType, valueP T.Gpointer, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileSetAttribute(f, attribute, typ, valueP, flags, cancellable, err)
}
func (f *File) SetAttributeByteString(attribute, value string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileSetAttributeByteString(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributeInt32(attribute string, value T.GInt32, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileSetAttributeInt32(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributeInt64(attribute string, value int64, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileSetAttributeInt64(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributesAsync(info *FileInfo, flags FileQueryInfoFlags, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileSetAttributesAsync(f, info, flags, ioPriority, cancellable, callback, userData)
}
func (f *File) SetAttributesFinish(result *AsyncResult, info **FileInfo, err **T.GError) T.Gboolean {
	return fileSetAttributesFinish(f, result, info, err)
}
func (f *File) SetAttributesFromInfo(info *FileInfo, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileSetAttributesFromInfo(f, info, flags, cancellable, err)
}
func (f *File) SetAttributeString(attribute, value string, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileSetAttributeString(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributeUint32(attribute string, value T.GUint32, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileSetAttributeUint32(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetAttributeUint64(attribute string, value uint64, flags FileQueryInfoFlags, cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileSetAttributeUint64(f, attribute, value, flags, cancellable, err)
}
func (f *File) SetDisplayName(displayName string, cancellable *Cancellable, err **T.GError) *File {
	return fileSetDisplayName(f, displayName, cancellable, err)
}
func (f *File) SetDisplayNameAsync(displayName string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileSetDisplayNameAsync(f, displayName, ioPriority, cancellable, callback, userData)
}
func (f *File) SetDisplayNameFinish(res *AsyncResult, err **T.GError) *File {
	return fileSetDisplayNameFinish(f, res, err)
}
func (f *File) StartMountable(flags DriveStartFlags, startOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileStartMountable(f, flags, startOperation, cancellable, callback, userData)
}
func (f *File) StartMountableFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return fileStartMountableFinish(f, result, err)
}
func (f *File) StopMountable(flags T.GMountUnmountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileStopMountable(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) StopMountableFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return fileStopMountableFinish(f, result, err)
}
func (f *File) SupportsThreadContexts() T.Gboolean { return fileSupportsThreadContexts(f) }
func (f *File) Trash(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileTrash(f, cancellable, err)
}
func (f *File) UnmountMountable(flags T.GMountUnmountFlags, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileUnmountMountable(f, flags, cancellable, callback, userData)
}
func (f *File) UnmountMountableFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return fileUnmountMountableFinish(f, result, err)
}
func (f *File) UnmountMountableWithOperation(flags T.GMountUnmountFlags, mountOperation *T.GMountOperation, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileUnmountMountableWithOperation(f, flags, mountOperation, cancellable, callback, userData)
}
func (f *File) UnmountMountableWithOperationFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return fileUnmountMountableWithOperationFinish(f, result, err)
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

	fileAttributeInfoListAdd    func(f *FileAttributeInfoList, name string, typ FileAttributeType, flags FileAttributeInfoFlags)
	fileAttributeInfoListDup    func(f *FileAttributeInfoList) *FileAttributeInfoList
	fileAttributeInfoListLookup func(f *FileAttributeInfoList, name string) *FileAttributeInfo
	fileAttributeInfoListRef    func(f *FileAttributeInfoList) *FileAttributeInfoList
	fileAttributeInfoListUnref  func(f *FileAttributeInfoList)
)

func (f *FileAttributeInfoList) Add(name string, typ FileAttributeType, flags FileAttributeInfoFlags) {
	fileAttributeInfoListAdd(f, name, typ, flags)
}
func (f *FileAttributeInfoList) Dup() *FileAttributeInfoList { return fileAttributeInfoListDup(f) }
func (f *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	return fileAttributeInfoListLookup(f, name)
}
func (f *FileAttributeInfoList) Ref() *FileAttributeInfoList { return fileAttributeInfoListRef(f) }
func (f *FileAttributeInfoList) Unref()                      { fileAttributeInfoListUnref(f) }

type FileAttributeMatcher struct{}

var (
	FileAttributeMatcherGetType func() O.Type
	FileAttributeMatcherNew     func(attributes string) *FileAttributeMatcher

	fileAttributeMatcherEnumerateNamespace func(f *FileAttributeMatcher, ns string) T.Gboolean
	fileAttributeMatcherEnumerateNext      func(f *FileAttributeMatcher) string
	fileAttributeMatcherMatches            func(f *FileAttributeMatcher, attribute string) T.Gboolean
	fileAttributeMatcherMatchesOnly        func(f *FileAttributeMatcher, attribute string) T.Gboolean
	fileAttributeMatcherRef                func(f *FileAttributeMatcher) *FileAttributeMatcher
	fileAttributeMatcherUnref              func(f *FileAttributeMatcher)
)

func (f *FileAttributeMatcher) EnumerateNamespace(ns string) T.Gboolean {
	return fileAttributeMatcherEnumerateNamespace(f, ns)
}
func (f *FileAttributeMatcher) EnumerateNext() string { return fileAttributeMatcherEnumerateNext(f) }
func (f *FileAttributeMatcher) Matches(attribute string) T.Gboolean {
	return fileAttributeMatcherMatches(f, attribute)
}
func (f *FileAttributeMatcher) MatchesOnly(attribute string) T.Gboolean {
	return fileAttributeMatcherMatchesOnly(f, attribute)
}
func (f *FileAttributeMatcher) Ref() *FileAttributeMatcher { return fileAttributeMatcherRef(f) }
func (f *FileAttributeMatcher) Unref()                     { fileAttributeMatcherUnref(f) }

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

	fileEnumeratorNextFile        func(f *FileEnumerator, cancellable *Cancellable, err **T.GError) *FileInfo
	fileEnumeratorClose           func(f *FileEnumerator, cancellable *Cancellable, err **T.GError) T.Gboolean
	fileEnumeratorNextFilesAsync  func(f *FileEnumerator, numFiles, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileEnumeratorNextFilesFinish func(f *FileEnumerator, result *AsyncResult, err **T.GError) *T.GList
	fileEnumeratorCloseAsync      func(f *FileEnumerator, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileEnumeratorCloseFinish     func(f *FileEnumerator, result *AsyncResult, err **T.GError) T.Gboolean
	fileEnumeratorIsClosed        func(f *FileEnumerator) T.Gboolean
	fileEnumeratorHasPending      func(f *FileEnumerator) T.Gboolean
	fileEnumeratorSetPending      func(f *FileEnumerator, pending T.Gboolean)
	fileEnumeratorGetContainer    func(f *FileEnumerator) *File
)

func (f *FileEnumerator) NextFile(cancellable *Cancellable, err **T.GError) *FileInfo {
	return fileEnumeratorNextFile(f, cancellable, err)
}
func (f *FileEnumerator) Close(cancellable *Cancellable, err **T.GError) T.Gboolean {
	return fileEnumeratorClose(f, cancellable, err)
}
func (f *FileEnumerator) NextFilesAsync(numFiles, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileEnumeratorNextFilesAsync(f, numFiles, ioPriority, cancellable, callback, userData)
}
func (f *FileEnumerator) NextFilesFinish(result *AsyncResult, err **T.GError) *T.GList {
	return fileEnumeratorNextFilesFinish(f, result, err)
}
func (f *FileEnumerator) CloseAsync(ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileEnumeratorCloseAsync(f, ioPriority, cancellable, callback, userData)
}
func (f *FileEnumerator) CloseFinish(result *AsyncResult, err **T.GError) T.Gboolean {
	return fileEnumeratorCloseFinish(f, result, err)
}
func (f *FileEnumerator) IsClosed() T.Gboolean          { return fileEnumeratorIsClosed(f) }
func (f *FileEnumerator) HasPending() T.Gboolean        { return fileEnumeratorHasPending(f) }
func (f *FileEnumerator) SetPending(pending T.Gboolean) { fileEnumeratorSetPending(f, pending) }
func (f *FileEnumerator) GetContainer() *File           { return fileEnumeratorGetContainer(f) }

type FileIcon struct{}

var (
	FileIconGetType func() O.Type
	FileIconNew     func(file *File) *Icon

	fileIconGetFile func(f *FileIcon) *File
)

func (f *FileIcon) GetFile() *File { return fileIconGetFile(f) }

type FileInfo struct{}

var (
	FileInfoGetType func() O.Type
	FileInfoNew     func() *FileInfo

	FileInfoDup      func(other *FileInfo) *FileInfo
	FileInfoCopyInto func(srcInfo, destInfo *FileInfo)

	fileInfoClearStatus            func(f *FileInfo)
	fileInfoGetAttributeAsString   func(f *FileInfo, attribute string) string
	fileInfoGetAttributeBoolean    func(f *FileInfo, attribute string) T.Gboolean
	fileInfoGetAttributeByteString func(f *FileInfo, attribute string) string
	fileInfoGetAttributeData       func(f *FileInfo, attribute string, typ *FileAttributeType, valuePp *T.Gpointer, status *FileAttributeStatus) T.Gboolean
	fileInfoGetAttributeInt32      func(f *FileInfo, attribute string) T.GInt32
	fileInfoGetAttributeInt64      func(f *FileInfo, attribute string) int64
	fileInfoGetAttributeObject     func(f *FileInfo, attribute string) *O.Object
	fileInfoGetAttributeStatus     func(f *FileInfo, attribute string) FileAttributeStatus
	fileInfoGetAttributeString     func(f *FileInfo, attribute string) string
	fileInfoGetAttributeStringv    func(f *FileInfo, attribute string) **T.Char
	fileInfoGetAttributeType       func(f *FileInfo, attribute string) FileAttributeType
	fileInfoGetAttributeUint32     func(f *FileInfo, attribute string) T.GUint32
	fileInfoGetAttributeUint64     func(f *FileInfo, attribute string) uint64
	fileInfoGetContentType         func(f *FileInfo) string
	fileInfoGetDisplayName         func(f *FileInfo) string
	fileInfoGetEditName            func(f *FileInfo) string
	fileInfoGetEtag                func(f *FileInfo) string
	fileInfoGetFileType            func(f *FileInfo) FileType
	fileInfoGetIcon                func(f *FileInfo) *Icon
	fileInfoGetIsBackup            func(f *FileInfo) T.Gboolean
	fileInfoGetIsHidden            func(f *FileInfo) T.Gboolean
	fileInfoGetIsSymlink           func(f *FileInfo) T.Gboolean
	fileInfoGetModificationTime    func(f *FileInfo, result *T.GTimeVal)
	fileInfoGetName                func(f *FileInfo) string
	fileInfoGetSize                func(f *FileInfo) T.Goffset
	fileInfoGetSortOrder           func(f *FileInfo) T.GInt32
	fileInfoGetSymlinkTarget       func(f *FileInfo) string
	fileInfoHasAttribute           func(f *FileInfo, attribute string) T.Gboolean
	fileInfoHasNamespace           func(f *FileInfo, nameSpace string) T.Gboolean
	fileInfoListAttributes         func(f *FileInfo, nameSpace string) **T.Char
	fileInfoRemoveAttribute        func(f *FileInfo, attribute string)
	fileInfoSetAttribute           func(f *FileInfo, attribute string, typ FileAttributeType, valueP T.Gpointer)
	fileInfoSetAttributeBoolean    func(f *FileInfo, attribute string, attrValue T.Gboolean)
	fileInfoSetAttributeByteString func(f *FileInfo, attribute, attrValue string)
	fileInfoSetAttributeInt32      func(f *FileInfo, attribute string, attrValue T.GInt32)
	fileInfoSetAttributeInt64      func(f *FileInfo, attribute string, attrValue int64)
	fileInfoSetAttributeMask       func(f *FileInfo, mask *FileAttributeMatcher)
	fileInfoSetAttributeObject     func(f *FileInfo, attribute string, attrValue *O.Object)
	fileInfoSetAttributeStatus     func(f *FileInfo, attribute string, status FileAttributeStatus) T.Gboolean
	fileInfoSetAttributeString     func(f *FileInfo, attribute, attrValue string)
	fileInfoSetAttributeStringv    func(f *FileInfo, attribute string, attrValue **T.Char)
	fileInfoSetAttributeUint32     func(f *FileInfo, attribute string, attrValue T.GUint32)
	fileInfoSetAttributeUint64     func(f *FileInfo, attribute string, attrValue uint64)
	fileInfoSetContentType         func(f *FileInfo, contentType string)
	fileInfoSetDisplayName         func(f *FileInfo, displayName string)
	fileInfoSetEditName            func(f *FileInfo, editName string)
	fileInfoSetFileType            func(f *FileInfo, typ FileType)
	fileInfoSetIcon                func(f *FileInfo, icon *Icon)
	fileInfoSetIsHidden            func(f *FileInfo, isHidden T.Gboolean)
	fileInfoSetIsSymlink           func(f *FileInfo, isSymlink T.Gboolean)
	fileInfoSetModificationTime    func(f *FileInfo, mtime *T.GTimeVal)
	fileInfoSetName                func(f *FileInfo, name string)
	fileInfoSetSize                func(f *FileInfo, size T.Goffset)
	fileInfoSetSortOrder           func(f *FileInfo, sortOrder T.GInt32)
	fileInfoSetSymlinkTarget       func(f *FileInfo, symlinkTarget string)
	fileInfoUnsetAttributeMask     func(f *FileInfo)
)

func (f *FileInfo) ClearStatus() { fileInfoClearStatus(f) }
func (f *FileInfo) GetAttributeAsString(attribute string) string {
	return fileInfoGetAttributeAsString(f, attribute)
}
func (f *FileInfo) GetAttributeBoolean(attribute string) T.Gboolean {
	return fileInfoGetAttributeBoolean(f, attribute)
}
func (f *FileInfo) GetAttributeByteString(attribute string) string {
	return fileInfoGetAttributeByteString(f, attribute)
}
func (f *FileInfo) GetAttributeData(attribute string, typ *FileAttributeType, valuePp *T.Gpointer, status *FileAttributeStatus) T.Gboolean {
	return fileInfoGetAttributeData(f, attribute, typ, valuePp, status)
}
func (f *FileInfo) GetAttributeInt32(attribute string) T.GInt32 {
	return fileInfoGetAttributeInt32(f, attribute)
}
func (f *FileInfo) GetAttributeInt64(attribute string) int64 {
	return fileInfoGetAttributeInt64(f, attribute)
}
func (f *FileInfo) GetAttributeObject(attribute string) *O.Object {
	return fileInfoGetAttributeObject(f, attribute)
}
func (f *FileInfo) GetAttributeStatus(attribute string) FileAttributeStatus {
	return fileInfoGetAttributeStatus(f, attribute)
}
func (f *FileInfo) GetAttributeString(attribute string) string {
	return fileInfoGetAttributeString(f, attribute)
}
func (f *FileInfo) GetAttributeStringv(attribute string) **T.Char {
	return fileInfoGetAttributeStringv(f, attribute)
}
func (f *FileInfo) GetAttributeType(attribute string) FileAttributeType {
	return fileInfoGetAttributeType(f, attribute)
}
func (f *FileInfo) GetAttributeUint32(attribute string) T.GUint32 {
	return fileInfoGetAttributeUint32(f, attribute)
}
func (f *FileInfo) GetAttributeUint64(attribute string) uint64 {
	return fileInfoGetAttributeUint64(f, attribute)
}
func (f *FileInfo) GetContentType() string                 { return fileInfoGetContentType(f) }
func (f *FileInfo) GetDisplayName() string                 { return fileInfoGetDisplayName(f) }
func (f *FileInfo) GetEditName() string                    { return fileInfoGetEditName(f) }
func (f *FileInfo) GetEtag() string                        { return fileInfoGetEtag(f) }
func (f *FileInfo) GetFileType() FileType                  { return fileInfoGetFileType(f) }
func (f *FileInfo) GetIcon() *Icon                         { return fileInfoGetIcon(f) }
func (f *FileInfo) GetIsBackup() T.Gboolean                { return fileInfoGetIsBackup(f) }
func (f *FileInfo) GetIsHidden() T.Gboolean                { return fileInfoGetIsHidden(f) }
func (f *FileInfo) GetIsSymlink() T.Gboolean               { return fileInfoGetIsSymlink(f) }
func (f *FileInfo) GetModificationTime(result *T.GTimeVal) { fileInfoGetModificationTime(f, result) }
func (f *FileInfo) GetName() string                        { return fileInfoGetName(f) }
func (f *FileInfo) GetSize() T.Goffset                     { return fileInfoGetSize(f) }
func (f *FileInfo) GetSortOrder() T.GInt32                 { return fileInfoGetSortOrder(f) }
func (f *FileInfo) GetSymlinkTarget() string               { return fileInfoGetSymlinkTarget(f) }
func (f *FileInfo) HasAttribute(attribute string) T.Gboolean {
	return fileInfoHasAttribute(f, attribute)
}
func (f *FileInfo) HasNamespace(nameSpace string) T.Gboolean {
	return fileInfoHasNamespace(f, nameSpace)
}
func (f *FileInfo) ListAttributes(nameSpace string) **T.Char {
	return fileInfoListAttributes(f, nameSpace)
}
func (f *FileInfo) RemoveAttribute(attribute string) { fileInfoRemoveAttribute(f, attribute) }
func (f *FileInfo) SetAttribute(attribute string, typ FileAttributeType, valueP T.Gpointer) {
	fileInfoSetAttribute(f, attribute, typ, valueP)
}
func (f *FileInfo) SetAttributeBoolean(attribute string, attrValue T.Gboolean) {
	fileInfoSetAttributeBoolean(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeByteString(attribute, attrValue string) {
	fileInfoSetAttributeByteString(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeInt32(attribute string, attrValue T.GInt32) {
	fileInfoSetAttributeInt32(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeInt64(attribute string, attrValue int64) {
	fileInfoSetAttributeInt64(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeMask(mask *FileAttributeMatcher) { fileInfoSetAttributeMask(f, mask) }
func (f *FileInfo) SetAttributeObject(attribute string, attrValue *O.Object) {
	fileInfoSetAttributeObject(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeStatus(attribute string, status FileAttributeStatus) T.Gboolean {
	return fileInfoSetAttributeStatus(f, attribute, status)
}
func (f *FileInfo) SetAttributeString(attribute, attrValue string) {
	fileInfoSetAttributeString(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeStringv(attribute string, attrValue **T.Char) {
	fileInfoSetAttributeStringv(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeUint32(attribute string, attrValue T.GUint32) {
	fileInfoSetAttributeUint32(f, attribute, attrValue)
}
func (f *FileInfo) SetAttributeUint64(attribute string, attrValue uint64) {
	fileInfoSetAttributeUint64(f, attribute, attrValue)
}
func (f *FileInfo) SetContentType(contentType string)     { fileInfoSetContentType(f, contentType) }
func (f *FileInfo) SetDisplayName(displayName string)     { fileInfoSetDisplayName(f, displayName) }
func (f *FileInfo) SetEditName(editName string)           { fileInfoSetEditName(f, editName) }
func (f *FileInfo) SetFileType(typ FileType)              { fileInfoSetFileType(f, typ) }
func (f *FileInfo) SetIcon(icon *Icon)                    { fileInfoSetIcon(f, icon) }
func (f *FileInfo) SetIsHidden(isHidden T.Gboolean)       { fileInfoSetIsHidden(f, isHidden) }
func (f *FileInfo) SetIsSymlink(isSymlink T.Gboolean)     { fileInfoSetIsSymlink(f, isSymlink) }
func (f *FileInfo) SetModificationTime(mtime *T.GTimeVal) { fileInfoSetModificationTime(f, mtime) }
func (f *FileInfo) SetName(name string)                   { fileInfoSetName(f, name) }
func (f *FileInfo) SetSize(size T.Goffset)                { fileInfoSetSize(f, size) }
func (f *FileInfo) SetSortOrder(sortOrder T.GInt32)       { fileInfoSetSortOrder(f, sortOrder) }
func (f *FileInfo) SetSymlinkTarget(symlinkTarget string) { fileInfoSetSymlinkTarget(f, symlinkTarget) }
func (f *FileInfo) UnsetAttributeMask()                   { fileInfoUnsetAttributeMask(f) }

type FileInputStream struct {
	Parent InputStream
	_      *struct{}
}

var (
	FileInputStreamGetType func() O.Type

	fileInputStreamQueryInfo       func(f *FileInputStream, attributes string, cancellable *Cancellable, err **T.GError) *FileInfo
	fileInputStreamQueryInfoAsync  func(f *FileInputStream, attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileInputStreamQueryInfoFinish func(f *FileInputStream, result *AsyncResult, err **T.GError) *FileInfo
)

func (f *FileInputStream) QueryInfo(attributes string, cancellable *Cancellable, err **T.GError) *FileInfo {
	return fileInputStreamQueryInfo(f, attributes, cancellable, err)
}
func (f *FileInputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileInputStreamQueryInfoAsync(f, attributes, ioPriority, cancellable, callback, userData)
}
func (f *FileInputStream) QueryInfoFinish(result *AsyncResult, err **T.GError) *FileInfo {
	return fileInputStreamQueryInfoFinish(f, result, err)
}

type FileIOStream struct {
	Parent IOStream
	_      *struct{}
}

var (
	FileIoStreamGetType func() O.Type

	fileIoStreamGetEtag         func(f *FileIOStream) string
	fileIoStreamQueryInfo       func(f *FileIOStream, attributes string, cancellable *Cancellable, err **T.GError) *FileInfo
	fileIoStreamQueryInfoAsync  func(f *FileIOStream, attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileIoStreamQueryInfoFinish func(f *FileIOStream, result *AsyncResult, err **T.GError) *FileInfo
)

func (f *FileIOStream) GetEtag() string { return fileIoStreamGetEtag(f) }
func (f *FileIOStream) QueryInfo(attributes string, cancellable *Cancellable, err **T.GError) *FileInfo {
	return fileIoStreamQueryInfo(f, attributes, cancellable, err)
}
func (f *FileIOStream) QueryInfoAsync(attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileIoStreamQueryInfoAsync(f, attributes, ioPriority, cancellable, callback, userData)
}
func (f *FileIOStream) QueryInfoFinish(result *AsyncResult, err **T.GError) *FileInfo {
	return fileIoStreamQueryInfoFinish(f, result, err)
}

type FileMonitor struct {
	Parent O.Object
	_      *struct{}
}

var (
	FileMonitorGetType func() O.Type

	fileMonitorCancel       func(f *FileMonitor) T.Gboolean
	fileMonitorEmitEvent    func(f *FileMonitor, child, otherFile *File, eventType FileMonitorEvent)
	fileMonitorIsCancelled  func(f *FileMonitor) T.Gboolean
	fileMonitorSetRateLimit func(f *FileMonitor, limitMsecs int)
)

func (f *FileMonitor) Cancel() T.Gboolean { return fileMonitorCancel(f) }
func (f *FileMonitor) EmitEvent(child, otherFile *File, eventType FileMonitorEvent) {
	fileMonitorEmitEvent(f, child, otherFile, eventType)
}
func (f *FileMonitor) IsCancelled() T.Gboolean     { return fileMonitorIsCancelled(f) }
func (f *FileMonitor) SetRateLimit(limitMsecs int) { fileMonitorSetRateLimit(f, limitMsecs) }

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

	filenameCompleterGetCompletions      func(f *FilenameCompleter, initialText string) **T.Char
	filenameCompleterGetCompletionSuffix func(f *FilenameCompleter, initialText string) string
	filenameCompleterSetDirsOnly         func(f *FilenameCompleter, dirsOnly T.Gboolean)
)

func (f *FilenameCompleter) GetCompletions(initialText string) **T.Char {
	return filenameCompleterGetCompletions(f, initialText)
}
func (f *FilenameCompleter) GetCompletionSuffix(initialText string) string {
	return filenameCompleterGetCompletionSuffix(f, initialText)
}
func (f *FilenameCompleter) SetDirsOnly(dirsOnly T.Gboolean) {
	filenameCompleterSetDirsOnly(f, dirsOnly)
}

type FileOutputStream struct {
	Parent OutputStream
	_      *struct{}
}

var (
	FileOutputStreamGetType func() O.Type

	fileOutputStreamGetEtag         func(f *FileOutputStream) string
	fileOutputStreamQueryInfo       func(f *FileOutputStream, attributes string, cancellable *Cancellable, err **T.GError) *FileInfo
	fileOutputStreamQueryInfoAsync  func(f *FileOutputStream, attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer)
	fileOutputStreamQueryInfoFinish func(f *FileOutputStream, result *AsyncResult, err **T.GError) *FileInfo
)

func (f *FileOutputStream) GetEtag() string { return fileOutputStreamGetEtag(f) }
func (f *FileOutputStream) QueryInfo(attributes string, cancellable *Cancellable, err **T.GError) *FileInfo {
	return fileOutputStreamQueryInfo(f, attributes, cancellable, err)
}
func (f *FileOutputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable *Cancellable, callback AsyncReadyCallback, userData T.Gpointer) {
	fileOutputStreamQueryInfoAsync(f, attributes, ioPriority, cancellable, callback, userData)
}
func (f *FileOutputStream) QueryInfoFinish(result *AsyncResult, err **T.GError) *FileInfo {
	return fileOutputStreamQueryInfoFinish(f, result, err)
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
	fileSize T.Goffset, callbackData T.Gpointer) T.Gboolean

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

	filterInputStreamGetBaseStream      func(f *FilterInputStream) *InputStream
	filterInputStreamGetCloseBaseStream func(f *FilterInputStream) T.Gboolean
	filterInputStreamSetCloseBaseStream func(f *FilterInputStream, closeBase T.Gboolean)
)

func (f *FilterInputStream) GetBaseStream() *InputStream { return filterInputStreamGetBaseStream(f) }
func (f *FilterInputStream) GetCloseBaseStream() T.Gboolean {
	return filterInputStreamGetCloseBaseStream(f)
}
func (f *FilterInputStream) SetCloseBaseStream(closeBase T.Gboolean) {
	filterInputStreamSetCloseBaseStream(f, closeBase)
}

type FilterOutputStream struct {
	Parent     OutputStream
	BaseStream *OutputStream
}

var (
	FilterOutputStreamGetType func() O.Type

	filterOutputStreamGetBaseStream      func(f *FilterOutputStream) *OutputStream
	filterOutputStreamGetCloseBaseStream func(f *FilterOutputStream) T.Gboolean
	filterOutputStreamSetCloseBaseStream func(f *FilterOutputStream, closeBase T.Gboolean)
)

func (f *FilterOutputStream) GetBaseStream() *OutputStream {
	return filterOutputStreamGetBaseStream(f)
}
func (f *FilterOutputStream) GetCloseBaseStream() T.Gboolean {
	return filterOutputStreamGetCloseBaseStream(f)
}
func (f *FilterOutputStream) SetCloseBaseStream(closeBase T.Gboolean) {
	filterOutputStreamSetCloseBaseStream(f, closeBase)
}
