// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package gail provides API definitions for accessing
//libgailutil-18.dll.
package gail

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

var (
	Gail_misc_add_attribute func(
		attrib_set *AtkAttributeSet,
		attr AtkTextAttribute,
		value string) *AtkAttributeSet

	Gail_misc_layout_get_run_attributes func(
		attrib_set *AtkAttributeSet,
		layout *PangoLayout,
		text string,
		offset Gint,
		start_offset *Gint,
		end_offset *Gint) *AtkAttributeSet

	Gail_misc_get_default_attributes func(
		attrib_set *AtkAttributeSet,
		layout *PangoLayout,
		widget *GtkWidget) *AtkAttributeSet

	Gail_misc_get_extents_from_pango_rectangle func(
		widget *GtkWidget,
		char_rect *PangoRectangle,
		x_layout Gint,
		y_layout Gint,
		x *Gint,
		y *Gint,
		width *Gint,
		height *Gint,
		coords AtkCoordType)

	Gail_misc_get_index_at_point_in_layout func(
		widget *GtkWidget,
		layout *PangoLayout,
		x_layout Gint,
		y_layout Gint,
		x Gint,
		y Gint,
		coords AtkCoordType) Gint

	Gail_misc_get_origins func(
		widget *GtkWidget,
		x_window *Gint,
		y_window *Gint,
		x_toplevel *Gint,
		y_toplevel *Gint)

	Gail_misc_add_to_attr_set func(
		attrib_set *AtkAttributeSet,
		attrs *GtkTextAttributes,
		attr AtkTextAttribute) *AtkAttributeSet

	Gail_misc_buffer_get_run_attributes func(
		buffer *GtkTextBuffer,
		offset Gint,
		start_offset *Gint,
		end_offset *Gint) *AtkAttributeSet

	Gail_text_util_get_type func() GType

	Gail_text_util_new func() *GailTextUtil

	Gail_text_util_text_setup func(
		textutil *GailTextUtil,
		text string)

	Gail_text_util_buffer_setup func(
		textutil *GailTextUtil,
		buffer *GtkTextBuffer)

	Gail_text_util_get_text func(
		textutil *GailTextUtil,
		layout Gpointer,
		function GailOffsetType,
		boundary_type AtkTextBoundary,
		offset Gint,
		start_offset *Gint,
		end_offset *Gint) string

	Gail_text_util_get_substring func(
		textutil *GailTextUtil,
		start_pos Gint,
		end_pos Gint) string
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
