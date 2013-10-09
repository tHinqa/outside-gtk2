// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package intl provides API definitions for accessing
//intl.dll.
package intl

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
}

var (
	// TODO(t): check documentation for *char usage as msgid, in ngettext etc.
	Libintl_gettext func(msgid string) string

	Libintl_dgettext func(domainname string, msgid string) string

	Libintl_dcgettext func(
		domainname string, msgid string, category int) string

	Libintl_ngettext func(
		msgid1 string, msgid2 string, n Unsigned_long) string

	Libintl_dngettext func(domainname string,
		msgid1 string, msgid2 string, n Unsigned_long) string

	Libintl_dcngettext func(domainname string, msgid1 string,
		msgid2 string, n Unsigned_long, category int) string

	Libintl_textdomain func(domainname string) string

	Libintl_bindtextdomain func(
		domainname string, dirname string) string

	Libintl_bind_textdomain_codeset func(
		domainname string, codeset string) string

	Libintl_fprintf func(*FILE, string, ...VArg) int

	Libintl_vfprintf func(*FILE, string, Va_list) int

	Libintl_printf func(string, ...VArg) int

	Libintl_vprintf func(string, Va_list) int

	Libintl_sprintf func(string, string, ...VArg) int

	Libintl_vsprintf func(string, string, Va_list) int

	Libintl_snprintf func(string, Size_t, string, ...VArg) int

	Libintl_vsnprintf func(string, Size_t, string, Va_list) int

	Libintl_setlocale func(int, string) string

	Libintl_set_relocation_prefix func(
		orig_prefix string, curr_prefix string)
)

var dll = "intl.dll"

var apiList = Apis{
	// Undocumented {"_nl_expand_alias",&_nl_expand_alias},
	// Undocumented {"_nl_msg_cat_cntr",&_nl_msg_cat_cntr},
	// Libibtl_ {"bind_textdomain_codeset", &bind_textdomain_codeset},
	// Libibtl_ {"bindtextdomain", &bindtextdomain},
	// Libibtl_ {"dcgettext", &dcgettext},
	// Libibtl_ {"dcngettext", &dcngettext},
	// Libibtl_ {"dgettext", &dgettext},
	// Libibtl_ {"dngettext", &dngettext},
	// Libibtl_ {"gettext", &gettext},
	{"libintl_bind_textdomain_codeset", &Libintl_bind_textdomain_codeset},
	{"libintl_bindtextdomain", &Libintl_bindtextdomain},
	{"libintl_dcgettext", &Libintl_dcgettext},
	{"libintl_dcngettext", &Libintl_dcngettext},
	{"libintl_dgettext", &Libintl_dgettext},
	{"libintl_dngettext", &Libintl_dngettext},
	{"libintl_fprintf", &Libintl_fprintf},
	{"libintl_gettext", &Libintl_gettext},
	{"libintl_ngettext", &Libintl_ngettext},
	{"libintl_printf", &Libintl_printf},
	{"libintl_set_relocation_prefix", &Libintl_set_relocation_prefix},
	{"libintl_setlocale", &Libintl_setlocale},
	{"libintl_snprintf", &Libintl_snprintf},
	{"libintl_sprintf", &Libintl_sprintf},
	{"libintl_textdomain", &Libintl_textdomain},
	{"libintl_vfprintf", &Libintl_vfprintf},
	{"libintl_vprintf", &Libintl_vprintf},
	{"libintl_vsnprintf", &Libintl_vsnprintf},
	{"libintl_vsprintf", &Libintl_vsprintf},
	// Libibtl_ {"ngettext", &ngettext},
	// Libibtl_ {"textdomain", &textdomain},
}
