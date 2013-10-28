// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gmodule provides API definitions for accessing
//libgmodule-2.0-0.dll.
package gmodule

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type Enum int

type Module struct{}

var (
	ModuleError     func() string
	ModuleSupported func() T.Gboolean
	ModuleOpen      func(
		fileName string, flags ModuleFlags) *Module
	ModuleBuildPath func(directory, moduleName string) string

	moduleClose        func(m *Module) T.Gboolean
	moduleMakeResident func(m *Module)
	moduleName         func(m *Module) string
	moduleSymbol       func(m *Module, symbolName string, symbol *T.Gpointer) T.Gboolean
)

func (m *Module) Close() T.Gboolean { return moduleClose(m) }
func (m *Module) MakeResident()     { moduleMakeResident(m) }
func (m *Module) Name() string      { return moduleName(m) }
func (m *Module) Symbol(symbolName string, symbol *T.Gpointer) T.Gboolean {
	return moduleSymbol(m, symbolName, symbol)
}

type ModuleFlags Enum

const (
	BIND_LAZY ModuleFlags = 1 << iota
	BIND_LOCAL
	BIND_MASK ModuleFlags = 0x03
)

var dll = "libgmodule-2.0-0.dll"

var apiList = outside.Apis{
	{"g_module_build_path", &ModuleBuildPath},
	{"g_module_close", &moduleClose},
	{"g_module_error", &ModuleError},
	{"g_module_make_resident", &moduleMakeResident},
	// Windows _utf8 {"g_module_name", &ModuleName},
	{"g_module_name_utf8", &moduleName},
	// Windows _utf8 {"g_module_open", &ModuleOpen},
	{"g_module_open_utf8", &ModuleOpen},
	{"g_module_supported", &ModuleSupported},
	{"g_module_symbol", &moduleSymbol},
}
