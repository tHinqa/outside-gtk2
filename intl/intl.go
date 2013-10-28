// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package intl provides API definitions for accessing
//intl.dll.
package intl

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	// TODO(t): check documentation for *char usage as msgid, in ngettext etc.
	Gettext func(msgid string) string

	Dgettext func(domainname string, msgid string) string

	Dcgettext func(
		domainname string, msgid string, category int) string

	Ngettext func(
		msgid1 string, msgid2 string, n T.UnsignedLong) string

	Dngettext func(domainname string,
		msgid1 string, msgid2 string, n T.UnsignedLong) string

	Dcngettext func(domainname string, msgid1 string,
		msgid2 string, n T.UnsignedLong, category int) string

	Textdomain func(domainname string) string

	Bindtextdomain func(
		domainname string, dirname string) string

	BindTextdomainCodeset func(
		domainname string, codeset string) string

	Fprintf func(*T.FILE, string, ...VArg) int

	Vfprintf func(*T.FILE, string, T.VaList) int

	Printf func(string, ...VArg) int

	Vprintf func(string, T.VaList) int

	Sprintf func(string, string, ...VArg) int

	Vsprintf func(string, string, T.VaList) int

	Snprintf func(string, T.Size_t, string, ...VArg) int

	Vsnprintf func(string, T.Size_t, string, T.VaList) int

	Setlocale func(int, string) string

	SetRelocationPrefix func(
		origPrefix string, currPrefix string)
)

var dll = "intl.dll"

var apiList = outside.Apis{
	// Undocumented {"_nl_expand_alias",&_nl_expand_alias},
	// Undocumented {"_nl_msg_cat_cntr",&_nl_msg_cat_cntr},
	// Libibtl_ {"bind_textdomain_codeset", &bind_textdomain_codeset},
	// Libibtl_ {"bindtextdomain", &bindtextdomain},
	// Libibtl_ {"dcgettext", &dcgettext},
	// Libibtl_ {"dcngettext", &dcngettext},
	// Libibtl_ {"dgettext", &dgettext},
	// Libibtl_ {"dngettext", &dngettext},
	// Libibtl_ {"gettext", &gettext},
	{"libintl_bind_textdomain_codeset", &BindTextdomainCodeset},
	{"libintl_bindtextdomain", &Bindtextdomain},
	{"libintl_dcgettext", &Dcgettext},
	{"libintl_dcngettext", &Dcngettext},
	{"libintl_dgettext", &Dgettext},
	{"libintl_dngettext", &Dngettext},
	{"libintl_fprintf", &Fprintf},
	{"libintl_gettext", &Gettext},
	{"libintl_ngettext", &Ngettext},
	{"libintl_printf", &Printf},
	{"libintl_set_relocation_prefix", &SetRelocationPrefix},
	{"libintl_setlocale", &Setlocale},
	{"libintl_snprintf", &Snprintf},
	{"libintl_sprintf", &Sprintf},
	{"libintl_textdomain", &Textdomain},
	{"libintl_vfprintf", &Vfprintf},
	{"libintl_vprintf", &Vprintf},
	{"libintl_vsnprintf", &Vsnprintf},
	{"libintl_vsprintf", &Vsprintf},
	// Libibtl_ {"ngettext", &ngettext},
	// Libibtl_ {"textdomain", &textdomain},
}
