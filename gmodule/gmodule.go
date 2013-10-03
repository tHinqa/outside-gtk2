package gmodule

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	G_module_supported func() Gboolean

	G_module_open func(
		file_name *Gchar,
		flags GModuleFlags) *GModule

	G_module_close func(
		module *GModule) Gboolean

	G_module_make_resident func(
		module *GModule)

	G_module_error func() *Gchar

	G_module_symbol func(
		module *GModule,
		symbol_name *Gchar,
		symbol *Gpointer) Gboolean

	G_module_name func(
		module *GModule) *Gchar

	G_module_build_path func(
		directory *Gchar,
		module_name *Gchar) *Gchar
)

var dll = "libgmodule-2.0-0.dll"

var apiList = outside.Apis{
	{"g_module_build_path", &G_module_build_path},
	{"g_module_close", &G_module_close},
	{"g_module_error", &G_module_error},
	{"g_module_make_resident", &G_module_make_resident},
	// Windows _utf8 {"g_module_name", &G_module_name},
	{"g_module_name_utf8", &G_module_name},
	// Windows _utf8 {"g_module_open", &G_module_open},
	{"g_module_open_utf8", &G_module_open},
	{"g_module_supported", &G_module_supported},
	{"g_module_symbol", &G_module_symbol},
}
