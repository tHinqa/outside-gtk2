// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type BookmarkFile struct{}

var (
	BookmarkFileNew func() *BookmarkFile

	BookmarkFileErrorQuark func() T.GQuark

	bookmarkFileAddApplication    func(bookmark *BookmarkFile, uri, name, exec string)
	bookmarkFileAddGroup          func(bookmark *BookmarkFile, uri, group string)
	bookmarkFileFree              func(bookmark *BookmarkFile)
	bookmarkFileGetAdded          func(bookmark *BookmarkFile, uri string, e **T.GError) T.TimeT
	bookmarkFileGetAppInfo        func(bookmark *BookmarkFile, uri, name string, exec **T.Gchar, count *uint, stamp *T.TimeT, e **T.GError) bool
	bookmarkFileGetApplications   func(bookmark *BookmarkFile, uri string, length *T.Gsize, e **T.GError) []string
	bookmarkFileGetDescription    func(bookmark *BookmarkFile, uri string, e **T.GError) string
	bookmarkFileGetGroups         func(bookmark *BookmarkFile, uri string, length *T.Gsize, e **T.GError) []string
	bookmarkFileGetIcon           func(bookmark *BookmarkFile, uri string, href, mimeType **T.Gchar, e **T.GError) bool
	bookmarkFileGetIsPrivate      func(bookmark *BookmarkFile, uri string, e **T.GError) bool
	bookmarkFileGetMimeType       func(bookmark *BookmarkFile, uri string, e **T.GError) string
	bookmarkFileGetModified       func(bookmark *BookmarkFile, uri string, e **T.GError) T.TimeT
	bookmarkFileGetSize           func(bookmark *BookmarkFile) int
	bookmarkFileGetTitle          func(bookmark *BookmarkFile, uri, e **T.GError) string
	bookmarkFileGetUris           func(bookmark *BookmarkFile, length *T.Gsize) []string
	bookmarkFileGetVisited        func(bookmark *BookmarkFile, uri string, e **T.GError) T.TimeT
	bookmarkFileHasApplication    func(bookmark *BookmarkFile, uri, name string, e **T.GError) bool
	bookmarkFileHasGroup          func(bookmark *BookmarkFile, uri, group string, e **T.GError) bool
	bookmarkFileHasItem           func(bookmark *BookmarkFile, uri string) bool
	bookmarkFileLoadFromData      func(bookmark *BookmarkFile, data string, length T.Gsize, e **T.GError) bool
	bookmarkFileLoadFromDataDirs  func(bookmark *BookmarkFile, file string, fullPath **T.Gchar, e **T.GError) bool
	bookmarkFileLoadFromFile      func(bookmark *BookmarkFile, filename string, e **T.GError) bool
	bookmarkFileMoveItem          func(bookmark *BookmarkFile, oldUri, newUri string, e **T.GError) bool
	bookmarkFileRemoveApplication func(bookmark *BookmarkFile, uri, name string, e **T.GError) bool
	bookmarkFileRemoveGroup       func(bookmark *BookmarkFile, uri, group string, e **T.GError) bool
	bookmarkFileRemoveItem        func(bookmark *BookmarkFile, uri string, e **T.GError) bool
	bookmarkFileSetAdded          func(bookmark *BookmarkFile, uri string, added T.TimeT)
	bookmarkFileSetAppInfo        func(bookmark *BookmarkFile, uri, name, exec string, count int, stamp T.TimeT, e **T.GError) bool
	bookmarkFileSetDescription    func(bookmark *BookmarkFile, uri string, description string)
	bookmarkFileSetGroups         func(bookmark *BookmarkFile, uri string, groups **T.Gchar, length T.Gsize)
	bookmarkFileSetIcon           func(bookmark *BookmarkFile, uri, href, mimeType string)
	bookmarkFileSetIsPrivate      func(bookmark *BookmarkFile, uri string, isPrivate bool)
	bookmarkFileSetMimeType       func(bookmark *BookmarkFile, uri, mimeType string)
	bookmarkFileSetModified       func(bookmark *BookmarkFile, uri string, modified T.TimeT)
	bookmarkFileSetTitle          func(bookmark *BookmarkFile, uri, title string)
	bookmarkFileSetVisited        func(bookmark *BookmarkFile, uri string, visited T.TimeT)
	bookmarkFileToData            func(bookmark *BookmarkFile, length *T.Gsize, e **T.GError) string
	bookmarkFileToFile            func(bookmark *BookmarkFile, filename string, e **T.GError) bool
)

func (b *BookmarkFile) AddApplication(uri, name, exec string) {
	bookmarkFileAddApplication(b, uri, name, exec)
}
func (b *BookmarkFile) AddGroup(uri, group string) { bookmarkFileAddGroup(b, uri, group) }
func (b *BookmarkFile) Free()                      { bookmarkFileFree(b) }
func (b *BookmarkFile) GetAdded(uri string, e **T.GError) T.TimeT {
	return bookmarkFileGetAdded(b, uri, e)
}
func (b *BookmarkFile) GetAppInfo(uri, name string, exec **T.Gchar, count *uint, stamp *T.TimeT, e **T.GError) bool {
	return bookmarkFileGetAppInfo(b, uri, name, exec, count, stamp, e)
}
func (b *BookmarkFile) GetApplications(uri string, length *T.Gsize, e **T.GError) []string {
	return bookmarkFileGetApplications(b, uri, length, e)
}
func (b *BookmarkFile) GetDescription(uri string, e **T.GError) string {
	return bookmarkFileGetDescription(b, uri, e)
}
func (b *BookmarkFile) GetGroups(uri string, length *T.Gsize, e **T.GError) []string {
	return bookmarkFileGetGroups(b, uri, length, e)
}
func (b *BookmarkFile) GetIcon(uri string, href, mimeType **T.Gchar, e **T.GError) bool {
	return bookmarkFileGetIcon(b, uri, href, mimeType, e)
}
func (b *BookmarkFile) GetIsPrivate(uri string, e **T.GError) bool {
	return bookmarkFileGetIsPrivate(b, uri, e)
}
func (b *BookmarkFile) GetMimeType(uri string, e **T.GError) string {
	return bookmarkFileGetMimeType(b, uri, e)
}
func (b *BookmarkFile) GetModified(uri string, e **T.GError) T.TimeT {
	return bookmarkFileGetModified(b, uri, e)
}
func (b *BookmarkFile) GetSize() int                      { return bookmarkFileGetSize(b) }
func (b *BookmarkFile) GetTitle(uri, e **T.GError) string { return bookmarkFileGetTitle(b, uri, e) }
func (b *BookmarkFile) GetUris(length *T.Gsize) []string  { return bookmarkFileGetUris(b, length) }
func (b *BookmarkFile) GetVisited(uri string, e **T.GError) T.TimeT {
	return bookmarkFileGetVisited(b, uri, e)
}
func (b *BookmarkFile) HasApplication(uri, name string, e **T.GError) bool {
	return bookmarkFileHasApplication(b, uri, name, e)
}
func (b *BookmarkFile) HasGroup(uri, group string, e **T.GError) bool {
	return bookmarkFileHasGroup(b, uri, group, e)
}
func (b *BookmarkFile) HasItem(uri string) bool { return bookmarkFileHasItem(b, uri) }
func (b *BookmarkFile) LoadFromData(data string, length T.Gsize, e **T.GError) bool {
	return bookmarkFileLoadFromData(b, data, length, e)
}
func (b *BookmarkFile) LoadFromDataDirs(file string, fullPath **T.Gchar, e **T.GError) bool {
	return bookmarkFileLoadFromDataDirs(b, file, fullPath, e)
}
func (b *BookmarkFile) LoadFromFile(filename string, e **T.GError) bool {
	return bookmarkFileLoadFromFile(b, filename, e)
}
func (b *BookmarkFile) MoveItem(oldUri, newUri string, e **T.GError) bool {
	return bookmarkFileMoveItem(b, oldUri, newUri, e)
}
func (b *BookmarkFile) RemoveApplication(uri, name string, e **T.GError) bool {
	return bookmarkFileRemoveApplication(b, uri, name, e)
}
func (b *BookmarkFile) RemoveGroup(uri, group string, e **T.GError) bool {
	return bookmarkFileRemoveGroup(b, uri, group, e)
}
func (b *BookmarkFile) RemoveItem(uri string, e **T.GError) bool {
	return bookmarkFileRemoveItem(b, uri, e)
}
func (b *BookmarkFile) SetAdded(uri string, added T.TimeT) { bookmarkFileSetAdded(b, uri, added) }
func (b *BookmarkFile) SetAppInfo(uri, name, exec string, count int, stamp T.TimeT, e **T.GError) bool {
	return bookmarkFileSetAppInfo(b, uri, name, exec, count, stamp, e)
}
func (b *BookmarkFile) SetDescription(uri string, description string) {
	bookmarkFileSetDescription(b, uri, description)
}
func (b *BookmarkFile) SetGroups(uri string, groups **T.Gchar, length T.Gsize) {
	bookmarkFileSetGroups(b, uri, groups, length)
}
func (b *BookmarkFile) SetIcon(uri, href, mimeType string) {
	bookmarkFileSetIcon(b, uri, href, mimeType)
}
func (b *BookmarkFile) SetIsPrivate(uri string, isPrivate bool) {
	bookmarkFileSetIsPrivate(b, uri, isPrivate)
}
func (b *BookmarkFile) SetMimeType(uri, mimeType string) { bookmarkFileSetMimeType(b, uri, mimeType) }
func (b *BookmarkFile) SetModified(uri string, modified T.TimeT) {
	bookmarkFileSetModified(b, uri, modified)
}
func (b *BookmarkFile) SetTitle(uri, title string) { bookmarkFileSetTitle(b, uri, title) }
func (b *BookmarkFile) SetVisited(uri string, visited T.TimeT) {
	bookmarkFileSetVisited(b, uri, visited)
}
func (b *BookmarkFile) ToData(length *T.Gsize, e **T.GError) string {
	return bookmarkFileToData(b, length, e)
}
func (b *BookmarkFile) ToFile(filename string, e **T.GError) bool {
	return bookmarkFileToFile(b, filename, e)
}

type ByteArray struct {
	Data *uint8
	Leng uint
}

var (
	ByteArrayNew      func() *ByteArray
	ByteArraySizedNew func(reservedSize uint) *ByteArray

	byteArrayAppend          func(array *ByteArray, data *uint8, leng uint) *ByteArray
	byteArrayFree            func(array *ByteArray, freeSegment bool) *uint8
	byteArrayPrepend         func(array *ByteArray, data *uint8, leng uint) *ByteArray
	byteArrayRef             func(array *ByteArray) *ByteArray
	byteArrayRemoveIndex     func(array *ByteArray, index uint) *ByteArray
	byteArrayRemoveIndexFast func(array *ByteArray, index uint) *ByteArray
	byteArrayRemoveRange     func(array *ByteArray, index, length uint) *ByteArray
	byteArraySetSize         func(array *ByteArray, length uint) *ByteArray
	byteArraySort            func(array *ByteArray, compareFunc T.GCompareFunc)
	byteArraySortWithData    func(array *ByteArray, compareFunc T.GCompareDataFunc, userData T.Gpointer)
	byteArrayUnref           func(array *ByteArray)
)

func (b *ByteArray) Append(data *uint8, leng uint) *ByteArray  { return byteArrayAppend(b, data, leng) }
func (b *ByteArray) Free(freeSegment bool) *uint8              { return byteArrayFree(b, freeSegment) }
func (b *ByteArray) Prepend(data *uint8, leng uint) *ByteArray { return byteArrayPrepend(b, data, leng) }
func (b *ByteArray) Ref() *ByteArray                           { return byteArrayRef(b) }
func (b *ByteArray) RemoveIndex(index uint) *ByteArray         { return byteArrayRemoveIndex(b, index) }
func (b *ByteArray) RemoveIndexFast(index uint) *ByteArray     { return byteArrayRemoveIndexFast(b, index) }
func (b *ByteArray) RemoveRange(index, length uint) *ByteArray {
	return byteArrayRemoveRange(b, index, length)
}
func (b *ByteArray) SetSize(length uint) *ByteArray  { return byteArraySetSize(b, length) }
func (b *ByteArray) Sort(compareFunc T.GCompareFunc) { byteArraySort(b, compareFunc) }
func (b *ByteArray) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) {
	byteArraySortWithData(b, compareFunc, userData)
}
func (b *ByteArray) Unref() { byteArrayUnref(b) }
