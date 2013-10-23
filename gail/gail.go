// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package gail provides API definitions for accessing
//libgailutil-18.dll.
package gail

import (
	"github.com/tHinqa/outside"
	G "github.com/tHinqa/outside-gtk2/gtk"
	T "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	Gail_misc_add_attribute func(
		attrib_set *T.AtkAttributeSet,
		attr T.AtkTextAttribute,
		value string) *T.AtkAttributeSet

	Gail_misc_layout_get_run_attributes func(
		attrib_set *T.AtkAttributeSet,
		layout *T.PangoLayout,
		text string,
		offset int,
		start_offset *int,
		end_offset *int) *T.AtkAttributeSet

	Gail_misc_get_default_attributes func(
		attrib_set *T.AtkAttributeSet,
		layout *T.PangoLayout,
		widget *G.Widget) *T.AtkAttributeSet

	Gail_misc_get_extents_from_pango_rectangle func(
		widget *G.Widget,
		char_rect *T.PangoRectangle,
		x_layout int,
		y_layout int,
		x *int,
		y *int,
		width *int,
		height *int,
		coords T.AtkCoordType)

	Gail_misc_get_index_at_point_in_layout func(
		widget *G.Widget,
		layout *T.PangoLayout,
		x_layout int,
		y_layout int,
		x int,
		y int,
		coords T.AtkCoordType) int

	Gail_misc_get_origins func(
		widget *G.Widget,
		x_window *int,
		y_window *int,
		x_toplevel *int,
		y_toplevel *int)

	Gail_misc_add_to_attr_set func(
		attrib_set *T.AtkAttributeSet,
		attrs *G.TextAttributes,
		attr T.AtkTextAttribute) *T.AtkAttributeSet

	Gail_misc_buffer_get_run_attributes func(
		buffer *G.TextBuffer,
		offset int,
		start_offset *int,
		end_offset *int) *T.AtkAttributeSet

	Gail_text_util_get_type func() T.GType

	Gail_text_util_new func() *T.GailTextUtil

	Gail_text_util_text_setup func(
		textutil *T.GailTextUtil,
		text string)

	Gail_text_util_buffer_setup func(
		textutil *T.GailTextUtil,
		buffer *G.TextBuffer)

	Gail_text_util_get_text func(
		textutil *T.GailTextUtil,
		layout T.Gpointer,
		function T.GailOffsetType,
		boundary_type T.AtkTextBoundary,
		offset int,
		start_offset *int,
		end_offset *int) string

	Gail_text_util_get_substring func(
		textutil *T.GailTextUtil,
		start_pos int,
		end_pos int) string
)

var dll = "libgailutil-18.dll"

var apiList = outside.Apis{
	{"gail_misc_add_attribute", &Gail_misc_add_attribute},
	{"gail_misc_add_to_attr_set", &Gail_misc_add_to_attr_set},
	{"gail_misc_buffer_get_run_attributes", &Gail_misc_buffer_get_run_attributes},
	{"gail_misc_get_default_attributes", &Gail_misc_get_default_attributes},
	{"gail_misc_get_extents_from_pango_rectangle", &Gail_misc_get_extents_from_pango_rectangle},
	{"gail_misc_get_index_at_point_in_layout", &Gail_misc_get_index_at_point_in_layout},
	{"gail_misc_get_origins", &Gail_misc_get_origins},
	{"gail_misc_layout_get_run_attributes", &Gail_misc_layout_get_run_attributes},
	{"gail_text_util_buffer_setup", &Gail_text_util_buffer_setup},
	{"gail_text_util_get_substring", &Gail_text_util_get_substring},
	{"gail_text_util_get_text", &Gail_text_util_get_text},
	{"gail_text_util_get_type", &Gail_text_util_get_type},
	{"gail_text_util_new", &Gail_text_util_new},
	{"gail_text_util_text_setup", &Gail_text_util_text_setup},
}
