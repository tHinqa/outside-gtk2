// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type KeyFile struct{}

var (
	KeyFileNew func() *KeyFile

	KeyFileErrorQuark func() Quark

	KeyFileFree                func(k *KeyFile)
	KeyFileGetBoolean          func(k *KeyFile, groupName, key string, e **Error) bool
	KeyFileGetBooleanList      func(k *KeyFile, groupName, key string, length *T.Gsize, e **Error) *bool
	KeyFileGetComment          func(k *KeyFile, groupName, key string, e **Error) string
	KeyFileGetDouble           func(k *KeyFile, groupName, key string, e **Error) float64
	KeyFileGetDoubleList       func(k *KeyFile, groupName, key string, length *T.Gsize, e **Error) *float64
	KeyFileGetGroups           func(k *KeyFile, length *T.Gsize) []string
	KeyFileGetInt64            func(k *KeyFile, groupName, key string, e **Error) int64
	KeyFileGetInteger          func(k *KeyFile, groupName, key string, e **Error) int
	KeyFileGetIntegerList      func(k *KeyFile, groupName, key string, length *T.Gsize, e **Error) *int
	KeyFileGetKeys             func(k *KeyFile, groupName string, length *T.Gsize, e **Error) []string
	KeyFileGetLocaleString     func(k *KeyFile, groupName, key, locale string, e **Error) string
	KeyFileGetLocaleStringList func(k *KeyFile, groupName, key, locale string, length *T.Gsize, e **Error) []string
	KeyFileGetStartGroup       func(k *KeyFile) string
	KeyFileGetString           func(k *KeyFile, groupName, key string, e **Error) string
	KeyFileGetStringList       func(k *KeyFile, groupName, key string, length *T.Gsize, e **Error) []string
	KeyFileGetUint64           func(k *KeyFile, groupName, key string, e **Error) uint64
	KeyFileGetValue            func(k *KeyFile, groupName, key string, e **Error) string
	KeyFileHasGroup            func(k *KeyFile, groupName string) bool
	KeyFileHasKey              func(k *KeyFile, groupName, key string, e **Error) bool
	KeyFileLoadFromData        func(k *KeyFile, data string, length T.Gsize, flags KeyFileFlags, e **Error) bool
	KeyFileLoadFromDataDirs    func(k *KeyFile, file string, fullPath **T.Gchar, flags KeyFileFlags, e **Error) bool
	KeyFileLoadFromDirs        func(k *KeyFile, file string, searchDirs []string, fullPath **T.Gchar, flags KeyFileFlags, e **Error) bool
	KeyFileLoadFromFile        func(k *KeyFile, file string, flags KeyFileFlags, e **Error) bool
	KeyFileRemoveComment       func(k *KeyFile, groupName, key string, e **Error) bool
	KeyFileRemoveGroup         func(k *KeyFile, groupName string, e **Error) bool
	KeyFileRemoveKey           func(k *KeyFile, groupName, key string, e **Error) bool
	KeyFileSetBoolean          func(k *KeyFile, groupName, key string, value bool)
	KeyFileSetBooleanList      func(k *KeyFile, groupName, key string, list *bool, length T.Gsize)
	KeyFileSetComment          func(k *KeyFile, groupName, key, comment string, e **Error) bool
	KeyFileSetDouble           func(k *KeyFile, groupName, key string, value float64)
	KeyFileSetDoubleList       func(k *KeyFile, groupName, key string, list *float64, length T.Gsize)
	KeyFileSetInt64            func(k *KeyFile, groupName, key string, value int64)
	KeyFileSetInteger          func(k *KeyFile, groupName, key string, value int)
	KeyFileSetIntegerList      func(k *KeyFile, groupName, key string, list *int, length T.Gsize)
	KeyFileSetListSeparator    func(k *KeyFile, separator T.Gchar)
	KeyFileSetLocaleString     func(k *KeyFile, groupName, key string, locale, str string)
	KeyFileSetLocaleStringList func(k *KeyFile, groupName, key, locale string, list []string, length T.Gsize)
	KeyFileSetString           func(k *KeyFile, groupName, key, str string)
	KeyFileSetStringList       func(k *KeyFile, groupName, key string, list []string, length T.Gsize)
	KeyFileSetUint64           func(k *KeyFile, groupName, key string, value uint64)
	KeyFileSetValue            func(k *KeyFile, groupName, key, value string)
	KeyFileToData              func(k *KeyFile, length *T.Gsize, e **Error) string
)

func (k *KeyFile) Free() { KeyFileFree(k) }
func (k *KeyFile) GetBoolean(groupName, key string, e **Error) bool {
	return KeyFileGetBoolean(k, groupName, key, e)
}
func (k *KeyFile) GetBooleanList(groupName, key string, length *T.Gsize, e **Error) *bool {
	return KeyFileGetBooleanList(k, groupName, key, length, e)
}
func (k *KeyFile) GetComment(groupName, key string, e **Error) string {
	return KeyFileGetComment(k, groupName, key, e)
}
func (k *KeyFile) GetDouble(groupName, key string, e **Error) float64 {
	return KeyFileGetDouble(k, groupName, key, e)
}
func (k *KeyFile) GetDoubleList(groupName, key string, length *T.Gsize, e **Error) *float64 {
	return KeyFileGetDoubleList(k, groupName, key, length, e)
}
func (k *KeyFile) GetGroups(length *T.Gsize) []string { return KeyFileGetGroups(k, length) }
func (k *KeyFile) GetInt64(groupName, key string, e **Error) int64 {
	return KeyFileGetInt64(k, groupName, key, e)
}
func (k *KeyFile) GetInteger(groupName, key string, e **Error) int {
	return KeyFileGetInteger(k, groupName, key, e)
}
func (k *KeyFile) GetIntegerList(groupName, key string, length *T.Gsize, e **Error) *int {
	return KeyFileGetIntegerList(k, groupName, key, length, e)
}
func (k *KeyFile) GetKeys(groupName string, length *T.Gsize, e **Error) []string {
	return KeyFileGetKeys(k, groupName, length, e)
}
func (k *KeyFile) GetLocaleString(groupName, key, locale string, e **Error) string {
	return KeyFileGetLocaleString(k, groupName, key, locale, e)
}
func (k *KeyFile) GetLocaleStringList(groupName, key, locale string, length *T.Gsize, e **Error) []string {
	return KeyFileGetLocaleStringList(k, groupName, key, locale, length, e)
}
func (k *KeyFile) GetStartGroup() string { return KeyFileGetStartGroup(k) }
func (k *KeyFile) GetString(groupName, key string, e **Error) string {
	return KeyFileGetString(k, groupName, key, e)
}
func (k *KeyFile) GetStringList(groupName, key string, length *T.Gsize, e **Error) []string {
	return KeyFileGetStringList(k, groupName, key, length, e)
}
func (k *KeyFile) GetUint64(groupName, key string, e **Error) uint64 {
	return KeyFileGetUint64(k, groupName, key, e)
}
func (k *KeyFile) GetValue(groupName, key string, e **Error) string {
	return KeyFileGetValue(k, groupName, key, e)
}
func (k *KeyFile) HasGroup(groupName string) bool { return KeyFileHasGroup(k, groupName) }
func (k *KeyFile) HasKey(groupName, key string, e **Error) bool {
	return KeyFileHasKey(k, groupName, key, e)
}
func (k *KeyFile) LoadFromData(data string, length T.Gsize, flags KeyFileFlags, e **Error) bool {
	return KeyFileLoadFromData(k, data, length, flags, e)
}
func (k *KeyFile) LoadFromDataDirs(file string, fullPath **T.Gchar, flags KeyFileFlags, e **Error) bool {
	return KeyFileLoadFromDataDirs(k, file, fullPath, flags, e)
}
func (k *KeyFile) LoadFromDirs(file string, searchDirs []string, fullPath **T.Gchar, flags KeyFileFlags, e **Error) bool {
	return KeyFileLoadFromDirs(k, file, searchDirs, fullPath, flags, e)
}
func (k *KeyFile) LoadFromFile(file string, flags KeyFileFlags, e **Error) bool {
	return KeyFileLoadFromFile(k, file, flags, e)
}
func (k *KeyFile) RemoveComment(groupName, key string, e **Error) bool {
	return KeyFileRemoveComment(k, groupName, key, e)
}
func (k *KeyFile) RemoveGroup(groupName string, e **Error) bool {
	return KeyFileRemoveGroup(k, groupName, e)
}
func (k *KeyFile) RemoveKey(groupName, key string, e **Error) bool {
	return KeyFileRemoveKey(k, groupName, key, e)
}
func (k *KeyFile) SetBoolean(groupName, key string, value bool) {
	KeyFileSetBoolean(k, groupName, key, value)
}
func (k *KeyFile) SetBooleanList(groupName, key string, list *bool, length T.Gsize) {
	KeyFileSetBooleanList(k, groupName, key, list, length)
}
func (k *KeyFile) SetComment(groupName, key, comment string, e **Error) bool {
	return KeyFileSetComment(k, groupName, key, comment, e)
}
func (k *KeyFile) SetDouble(groupName, key string, value float64) {
	KeyFileSetDouble(k, groupName, key, value)
}
func (k *KeyFile) SetDoubleList(groupName, key string, list *float64, length T.Gsize) {
	KeyFileSetDoubleList(k, groupName, key, list, length)
}
func (k *KeyFile) SetInt64(groupName, key string, value int64) {
	KeyFileSetInt64(k, groupName, key, value)
}
func (k *KeyFile) SetInteger(groupName, key string, value int) {
	KeyFileSetInteger(k, groupName, key, value)
}
func (k *KeyFile) SetIntegerList(groupName, key string, list *int, length T.Gsize) {
	KeyFileSetIntegerList(k, groupName, key, list, length)
}
func (k *KeyFile) SetListSeparator(separator T.Gchar) { KeyFileSetListSeparator(k, separator) }
func (k *KeyFile) SetLocaleString(groupName, key, locale, str string) {
	KeyFileSetLocaleString(k, groupName, key, locale, str)
}
func (k *KeyFile) SetLocaleStringList(groupName, key, locale string, list []string, length T.Gsize) {
	KeyFileSetLocaleStringList(k, groupName, key, locale, list, length)
}
func (k *KeyFile) SetString(groupName, key, str string) { KeyFileSetString(k, groupName, key, str) }
func (k *KeyFile) SetStringList(groupName, key string, list []string, length T.Gsize) {
	KeyFileSetStringList(k, groupName, key, list, length)
}
func (k *KeyFile) SetUint64(groupName, key string, value uint64) {
	KeyFileSetUint64(k, groupName, key, value)
}
func (k *KeyFile) SetValue(groupName, key, value string)    { KeyFileSetValue(k, groupName, key, value) }
func (k *KeyFile) ToData(length *T.Gsize, e **Error) string { return KeyFileToData(k, length, e) }

type KeyFileFlags Enum

const (
	KEY_FILE_KEEP_COMMENTS KeyFileFlags = 1 << iota
	KEY_FILE_KEEP_TRANSLATIONS
	KEY_FILE_NONE KeyFileFlags = 0
)
