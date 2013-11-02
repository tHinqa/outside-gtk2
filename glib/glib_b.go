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

	BookmarkFileErrorQuark func() Quark

	BookmarkFileAddApplication    func(bookmark *BookmarkFile, uri, name, exec string)
	BookmarkFileAddGroup          func(bookmark *BookmarkFile, uri, group string)
	BookmarkFileFree              func(bookmark *BookmarkFile)
	BookmarkFileGetAdded          func(bookmark *BookmarkFile, uri string, e **Error) T.TimeT
	BookmarkFileGetAppInfo        func(bookmark *BookmarkFile, uri, name string, exec **T.Gchar, count *uint, stamp *T.TimeT, e **Error) bool
	BookmarkFileGetApplications   func(bookmark *BookmarkFile, uri string, length *T.Gsize, e **Error) []string
	BookmarkFileGetDescription    func(bookmark *BookmarkFile, uri string, e **Error) string
	BookmarkFileGetGroups         func(bookmark *BookmarkFile, uri string, length *T.Gsize, e **Error) []string
	BookmarkFileGetIcon           func(bookmark *BookmarkFile, uri string, href, mimeType **T.Gchar, e **Error) bool
	BookmarkFileGetIsPrivate      func(bookmark *BookmarkFile, uri string, e **Error) bool
	BookmarkFileGetMimeType       func(bookmark *BookmarkFile, uri string, e **Error) string
	BookmarkFileGetModified       func(bookmark *BookmarkFile, uri string, e **Error) T.TimeT
	BookmarkFileGetSize           func(bookmark *BookmarkFile) int
	BookmarkFileGetTitle          func(bookmark *BookmarkFile, uri, e **Error) string
	BookmarkFileGetUris           func(bookmark *BookmarkFile, length *T.Gsize) []string
	BookmarkFileGetVisited        func(bookmark *BookmarkFile, uri string, e **Error) T.TimeT
	BookmarkFileHasApplication    func(bookmark *BookmarkFile, uri, name string, e **Error) bool
	BookmarkFileHasGroup          func(bookmark *BookmarkFile, uri, group string, e **Error) bool
	BookmarkFileHasItem           func(bookmark *BookmarkFile, uri string) bool
	BookmarkFileLoadFromData      func(bookmark *BookmarkFile, data string, length T.Gsize, e **Error) bool
	BookmarkFileLoadFromDataDirs  func(bookmark *BookmarkFile, file string, fullPath **T.Gchar, e **Error) bool
	BookmarkFileLoadFromFile      func(bookmark *BookmarkFile, filename string, e **Error) bool
	BookmarkFileMoveItem          func(bookmark *BookmarkFile, oldUri, newUri string, e **Error) bool
	BookmarkFileRemoveApplication func(bookmark *BookmarkFile, uri, name string, e **Error) bool
	BookmarkFileRemoveGroup       func(bookmark *BookmarkFile, uri, group string, e **Error) bool
	BookmarkFileRemoveItem        func(bookmark *BookmarkFile, uri string, e **Error) bool
	BookmarkFileSetAdded          func(bookmark *BookmarkFile, uri string, added T.TimeT)
	BookmarkFileSetAppInfo        func(bookmark *BookmarkFile, uri, name, exec string, count int, stamp T.TimeT, e **Error) bool
	BookmarkFileSetDescription    func(bookmark *BookmarkFile, uri string, description string)
	BookmarkFileSetGroups         func(bookmark *BookmarkFile, uri string, groups **T.Gchar, length T.Gsize)
	BookmarkFileSetIcon           func(bookmark *BookmarkFile, uri, href, mimeType string)
	BookmarkFileSetIsPrivate      func(bookmark *BookmarkFile, uri string, isPrivate bool)
	BookmarkFileSetMimeType       func(bookmark *BookmarkFile, uri, mimeType string)
	BookmarkFileSetModified       func(bookmark *BookmarkFile, uri string, modified T.TimeT)
	BookmarkFileSetTitle          func(bookmark *BookmarkFile, uri, title string)
	BookmarkFileSetVisited        func(bookmark *BookmarkFile, uri string, visited T.TimeT)
	BookmarkFileToData            func(bookmark *BookmarkFile, length *T.Gsize, e **Error) string
	BookmarkFileToFile            func(bookmark *BookmarkFile, filename string, e **Error) bool
)

func (b *BookmarkFile) AddApplication(uri, name, exec string) {
	BookmarkFileAddApplication(b, uri, name, exec)
}
func (b *BookmarkFile) AddGroup(uri, group string) { BookmarkFileAddGroup(b, uri, group) }
func (b *BookmarkFile) Free()                      { BookmarkFileFree(b) }
func (b *BookmarkFile) GetAdded(uri string, e **Error) T.TimeT {
	return BookmarkFileGetAdded(b, uri, e)
}
func (b *BookmarkFile) GetAppInfo(uri, name string, exec **T.Gchar, count *uint, stamp *T.TimeT, e **Error) bool {
	return BookmarkFileGetAppInfo(b, uri, name, exec, count, stamp, e)
}
func (b *BookmarkFile) GetApplications(uri string, length *T.Gsize, e **Error) []string {
	return BookmarkFileGetApplications(b, uri, length, e)
}
func (b *BookmarkFile) GetDescription(uri string, e **Error) string {
	return BookmarkFileGetDescription(b, uri, e)
}
func (b *BookmarkFile) GetGroups(uri string, length *T.Gsize, e **Error) []string {
	return BookmarkFileGetGroups(b, uri, length, e)
}
func (b *BookmarkFile) GetIcon(uri string, href, mimeType **T.Gchar, e **Error) bool {
	return BookmarkFileGetIcon(b, uri, href, mimeType, e)
}
func (b *BookmarkFile) GetIsPrivate(uri string, e **Error) bool {
	return BookmarkFileGetIsPrivate(b, uri, e)
}
func (b *BookmarkFile) GetMimeType(uri string, e **Error) string {
	return BookmarkFileGetMimeType(b, uri, e)
}
func (b *BookmarkFile) GetModified(uri string, e **Error) T.TimeT {
	return BookmarkFileGetModified(b, uri, e)
}
func (b *BookmarkFile) GetSize() int                     { return BookmarkFileGetSize(b) }
func (b *BookmarkFile) GetTitle(uri, e **Error) string   { return BookmarkFileGetTitle(b, uri, e) }
func (b *BookmarkFile) GetUris(length *T.Gsize) []string { return BookmarkFileGetUris(b, length) }
func (b *BookmarkFile) GetVisited(uri string, e **Error) T.TimeT {
	return BookmarkFileGetVisited(b, uri, e)
}
func (b *BookmarkFile) HasApplication(uri, name string, e **Error) bool {
	return BookmarkFileHasApplication(b, uri, name, e)
}
func (b *BookmarkFile) HasGroup(uri, group string, e **Error) bool {
	return BookmarkFileHasGroup(b, uri, group, e)
}
func (b *BookmarkFile) HasItem(uri string) bool { return BookmarkFileHasItem(b, uri) }
func (b *BookmarkFile) LoadFromData(data string, length T.Gsize, e **Error) bool {
	return BookmarkFileLoadFromData(b, data, length, e)
}
func (b *BookmarkFile) LoadFromDataDirs(file string, fullPath **T.Gchar, e **Error) bool {
	return BookmarkFileLoadFromDataDirs(b, file, fullPath, e)
}
func (b *BookmarkFile) LoadFromFile(filename string, e **Error) bool {
	return BookmarkFileLoadFromFile(b, filename, e)
}
func (b *BookmarkFile) MoveItem(oldUri, newUri string, e **Error) bool {
	return BookmarkFileMoveItem(b, oldUri, newUri, e)
}
func (b *BookmarkFile) RemoveApplication(uri, name string, e **Error) bool {
	return BookmarkFileRemoveApplication(b, uri, name, e)
}
func (b *BookmarkFile) RemoveGroup(uri, group string, e **Error) bool {
	return BookmarkFileRemoveGroup(b, uri, group, e)
}
func (b *BookmarkFile) RemoveItem(uri string, e **Error) bool {
	return BookmarkFileRemoveItem(b, uri, e)
}
func (b *BookmarkFile) SetAdded(uri string, added T.TimeT) { BookmarkFileSetAdded(b, uri, added) }
func (b *BookmarkFile) SetAppInfo(uri, name, exec string, count int, stamp T.TimeT, e **Error) bool {
	return BookmarkFileSetAppInfo(b, uri, name, exec, count, stamp, e)
}
func (b *BookmarkFile) SetDescription(uri string, description string) {
	BookmarkFileSetDescription(b, uri, description)
}
func (b *BookmarkFile) SetGroups(uri string, groups **T.Gchar, length T.Gsize) {
	BookmarkFileSetGroups(b, uri, groups, length)
}
func (b *BookmarkFile) SetIcon(uri, href, mimeType string) {
	BookmarkFileSetIcon(b, uri, href, mimeType)
}
func (b *BookmarkFile) SetIsPrivate(uri string, isPrivate bool) {
	BookmarkFileSetIsPrivate(b, uri, isPrivate)
}
func (b *BookmarkFile) SetMimeType(uri, mimeType string) { BookmarkFileSetMimeType(b, uri, mimeType) }
func (b *BookmarkFile) SetModified(uri string, modified T.TimeT) {
	BookmarkFileSetModified(b, uri, modified)
}
func (b *BookmarkFile) SetTitle(uri, title string) { BookmarkFileSetTitle(b, uri, title) }
func (b *BookmarkFile) SetVisited(uri string, visited T.TimeT) {
	BookmarkFileSetVisited(b, uri, visited)
}
func (b *BookmarkFile) ToData(length *T.Gsize, e **Error) string {
	return BookmarkFileToData(b, length, e)
}
func (b *BookmarkFile) ToFile(filename string, e **Error) bool {
	return BookmarkFileToFile(b, filename, e)
}

type ByteArray struct {
	Data *uint8
	Leng uint
}

var (
	ByteArrayNew      func() *ByteArray
	ByteArraySizedNew func(reservedSize uint) *ByteArray

	ByteArrayAppend          func(array *ByteArray, data *uint8, leng uint) *ByteArray
	ByteArrayFree            func(array *ByteArray, freeSegment bool) *uint8
	ByteArrayPrepend         func(array *ByteArray, data *uint8, leng uint) *ByteArray
	ByteArrayRef             func(array *ByteArray) *ByteArray
	ByteArrayRemoveIndex     func(array *ByteArray, index uint) *ByteArray
	ByteArrayRemoveIndexFast func(array *ByteArray, index uint) *ByteArray
	ByteArrayRemoveRange     func(array *ByteArray, index, length uint) *ByteArray
	ByteArraySetSize         func(array *ByteArray, length uint) *ByteArray
	ByteArraySort            func(array *ByteArray, compareFunc T.GCompareFunc)
	ByteArraySortWithData    func(array *ByteArray, compareFunc T.GCompareDataFunc, userData T.Gpointer)
	ByteArrayUnref           func(array *ByteArray)
)

func (b *ByteArray) Append(data *uint8, leng uint) *ByteArray  { return ByteArrayAppend(b, data, leng) }
func (b *ByteArray) Free(freeSegment bool) *uint8              { return ByteArrayFree(b, freeSegment) }
func (b *ByteArray) Prepend(data *uint8, leng uint) *ByteArray { return ByteArrayPrepend(b, data, leng) }
func (b *ByteArray) Ref() *ByteArray                           { return ByteArrayRef(b) }
func (b *ByteArray) RemoveIndex(index uint) *ByteArray         { return ByteArrayRemoveIndex(b, index) }
func (b *ByteArray) RemoveIndexFast(index uint) *ByteArray     { return ByteArrayRemoveIndexFast(b, index) }
func (b *ByteArray) RemoveRange(index, length uint) *ByteArray {
	return ByteArrayRemoveRange(b, index, length)
}
func (b *ByteArray) SetSize(length uint) *ByteArray  { return ByteArraySetSize(b, length) }
func (b *ByteArray) Sort(compareFunc T.GCompareFunc) { ByteArraySort(b, compareFunc) }
func (b *ByteArray) SortWithData(compareFunc T.GCompareDataFunc, userData T.Gpointer) {
	ByteArraySortWithData(b, compareFunc, userData)
}
func (b *ByteArray) Unref() { ByteArrayUnref(b) }
