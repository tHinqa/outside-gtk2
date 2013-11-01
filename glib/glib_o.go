// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type OptionArg Enum

const (
	OPTION_ARG_NONE OptionArg = iota
	OPTION_ARG_STRING
	OPTION_ARG_INT
	OPTION_ARG_CALLBACK
	OPTION_ARG_FILENAME
	OPTION_ARG_STRING_ARRAY
	OPTION_ARG_FILENAME_ARRAY
	OPTION_ARG_DOUBLE
	OPTION_ARG_INT64
)

type OptionContext struct{}

var (
	OptionContextNew func(parameterString string) *OptionContext

	OptionContextAddGroup                func(o *OptionContext, group *OptionGroup)
	OptionContextAddMainEntries          func(o *OptionContext, entries *OptionEntry, translationDomain string)
	OptionContextFree                    func(o *OptionContext)
	OptionContextGetDescription          func(o *OptionContext) string
	OptionContextGetHelp                 func(o *OptionContext, mainHelp bool, group *OptionGroup) string
	OptionContextGetHelpEnabled          func(o *OptionContext) bool
	OptionContextGetIgnoreUnknownOptions func(o *OptionContext) bool
	OptionContextGetMainGroup            func(o *OptionContext) *OptionGroup
	OptionContextGetSummary              func(o *OptionContext) string
	OptionContextParse                   func(o *OptionContext, argc *int, argv ***T.Gchar, e **T.GError) bool
	OptionContextSetDescription          func(o *OptionContext, description string)
	OptionContextSetHelpEnabled          func(o *OptionContext, helpEnabled bool)
	OptionContextSetIgnoreUnknownOptions func(o *OptionContext, ignoreUnknown bool)
	OptionContextSetMainGroup            func(o *OptionContext, group *OptionGroup)
	OptionContextSetSummary              func(o *OptionContext, summary string)
	OptionContextSetTranslateFunc        func(o *OptionContext, f T.GTranslateFunc, data T.Gpointer, destroyNotify T.GDestroyNotify)
	OptionContextSetTranslationDomain    func(o *OptionContext, domain string)
)

func (o *OptionContext) AddGroup(group *OptionGroup) { OptionContextAddGroup(o, group) }
func (o *OptionContext) AddMainEntries(entries *OptionEntry, translationDomain string) {
	OptionContextAddMainEntries(o, entries, translationDomain)
}
func (o *OptionContext) Free()                  { OptionContextFree(o) }
func (o *OptionContext) GetDescription() string { return OptionContextGetDescription(o) }
func (o *OptionContext) GetHelp(mainHelp bool, group *OptionGroup) string {
	return OptionContextGetHelp(o, mainHelp, group)
}
func (o *OptionContext) GetHelpEnabled() bool          { return OptionContextGetHelpEnabled(o) }
func (o *OptionContext) GetIgnoreUnknownOptions() bool { return OptionContextGetIgnoreUnknownOptions(o) }
func (o *OptionContext) GetMainGroup() *OptionGroup    { return OptionContextGetMainGroup(o) }
func (o *OptionContext) GetSummary() string            { return OptionContextGetSummary(o) }
func (o *OptionContext) Parse(argc *int, argv ***T.Gchar, e **T.GError) bool {
	return OptionContextParse(o, argc, argv, e)
}
func (o *OptionContext) SetDescription(description string) {
	OptionContextSetDescription(o, description)
}
func (o *OptionContext) SetHelpEnabled(helpEnabled bool) { OptionContextSetHelpEnabled(o, helpEnabled) }
func (o *OptionContext) SetIgnoreUnknownOptions(ignoreUnknown bool) {
	OptionContextSetIgnoreUnknownOptions(o, ignoreUnknown)
}
func (o *OptionContext) SetMainGroup(group *OptionGroup) { OptionContextSetMainGroup(o, group) }
func (o *OptionContext) SetSummary(summary string)       { OptionContextSetSummary(o, summary) }
func (o *OptionContext) SetTranslateFunc(f T.GTranslateFunc, data T.Gpointer, destroyNotify T.GDestroyNotify) {
	OptionContextSetTranslateFunc(o, f, data, destroyNotify)
}
func (o *OptionContext) SetTranslationDomain(domain string) {
	OptionContextSetTranslationDomain(o, domain)
}

type OptionEntry struct {
	LongName       *T.Gchar
	ShortName      T.Gchar
	Flags          int
	Arg            OptionArg
	ArgData        T.Gpointer
	Description    *T.Gchar
	ArgDescription *T.Gchar
}

type OptionErrorFunc func(
	context *OptionContext,
	group *OptionGroup,
	data T.Gpointer,
	err **T.GError)

type OptionGroup struct{}

var (
	OptionGroupNew func(name string, description string, helpDescription string, userData T.Gpointer, destroy T.GDestroyNotify) *OptionGroup

	OptionGroupAddEntries           func(o *OptionGroup, entries *OptionEntry)
	OptionGroupFree                 func(o *OptionGroup)
	OptionGroupSetErrorHook         func(o *OptionGroup, errorFunc OptionErrorFunc)
	OptionGroupSetParseHooks        func(o *OptionGroup, preParseFunc OptionParseFunc, postParseFunc OptionParseFunc)
	OptionGroupSetTranslateFunc     func(o *OptionGroup, f T.GTranslateFunc, data T.Gpointer, destroyNotify T.GDestroyNotify)
	OptionGroupSetTranslationDomain func(o *OptionGroup, domain string)
)

func (o *OptionGroup) AddEntries(entries *OptionEntry)        { OptionGroupAddEntries(o, entries) }
func (o *OptionGroup) Free()                                  { OptionGroupFree(o) }
func (o *OptionGroup) SetErrorHook(errorFunc OptionErrorFunc) { OptionGroupSetErrorHook(o, errorFunc) }
func (o *OptionGroup) SetParseHooks(preParseFunc OptionParseFunc, postParseFunc OptionParseFunc) {
	OptionGroupSetParseHooks(o, preParseFunc, postParseFunc)
}
func (o *OptionGroup) SetTranslateFunc(f T.GTranslateFunc, data T.Gpointer, destroyNotify T.GDestroyNotify) {
	OptionGroupSetTranslateFunc(o, f, data, destroyNotify)
}
func (o *OptionGroup) SetTranslationDomain(domain string) { OptionGroupSetTranslationDomain(o, domain) }

type OptionParseFunc func(
	context *OptionContext,
	group *OptionGroup,
	data T.Gpointer,
	err **T.GError) T.Gboolean
