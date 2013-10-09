// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package png provides API definitions for accessing
//libpng14-14.dll.
package png

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside-gtk2/zlib"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

//TODO(t):check struct packing/sizes

type (
	Jmp_buf             [16]int
	Png_alloc_size_t    Png_size_t
	Png_byte            Unsigned_char
	Png_bytep           *Png_byte
	Png_bytepp          **Png_byte
	Png_charp           *Char
	Png_charpp          **Char
	Png_color_16p       *Png_color_16
	Png_color_8p        *Png_color_8
	Png_colorp          *Png_color
	Png_const_charp     Png_charp
	Png_doublep         *Double
	Png_FILE_p          *FILE
	Png_fixed_point     Png_int_32
	Png_infop           *Png_info
	Png_infopp          **Png_info
	Png_int_32          int // Anomally: Size?
	Png_row_infop       *Png_row_infop
	Png_size_t          Size_t
	Png_sPLT_entryp     *Png_sPLT_entry
	Png_sPLT_tp         *Png_sPLT_t
	Png_sPLT_tpp        **Png_sPLT_t
	Png_struct          Png_struct_def
	Png_structp         *Png_struct
	Png_structpp        **Png_struct
	Png_textp           *Png_text
	Png_timep           Png_time
	Png_uint_16         Unsigned_short
	Png_uint_16p        *Png_uint_16
	Png_uint_16pp       **Png_uint_16
	Png_uint_32         Unsigned_int // Anomally: Size?
	Png_unknown_chunkp  *Png_unknown_chunk
	Png_unknown_chunkpp **Png_unknown_chunk
	Png_voidp           *Void

	Png_longjmp_ptr func(Jmp_buf, int)

	Png_error_ptr func(Png_structp, Png_const_charp)

	Png_user_transform_ptr func(
		Png_structp, Png_row_infop, Png_bytep)

	Png_flush_ptr func(Png_structp)

	Png_rw_ptr func(Png_structp, Png_bytep, Png_size_t)

	Png_read_status_ptr func(Png_structp, Png_uint_32, int)

	Png_write_status_ptr func(Png_structp, Png_uint_32, int)

	Png_progressive_info_ptr func(Png_structp, Png_infop)

	Png_progressive_end_ptr func(Png_structp, Png_infop)

	Png_progressive_row_ptr func(
		Png_structp, Png_bytep, Png_uint_32, int)

	Png_user_chunk_ptr func(Png_structp, Png_unknown_chunkp) int

	Png_malloc_ptr func(Png_structp, Png_alloc_size_t) Png_voidp

	Png_free_ptr func(Png_structp, Png_voidp)

	//Png_unknown_chunk_ptr func(Png_structp)

	Png_time struct {
		Year   Png_uint_16
		Month  Png_byte
		Day    Png_byte
		Hour   Png_byte
		Minute Png_byte
		Second Png_byte
	}

	Png_color_8 struct {
		Red   Png_byte
		Green Png_byte
		Blue  Png_byte
		Gray  Png_byte
		Alpha Png_byte
	}

	Png_row_info struct {
		Width       Png_uint_32
		Rowbytes    Png_size_t
		Color_type  Png_byte
		Bit_depth   Png_byte
		Channels    Png_byte
		Pixel_depth Png_byte
	}

	Png_color struct {
		Red   Png_byte
		Green Png_byte
		Blue  Png_byte
	}

	Png_color_16 struct {
		Index Png_byte
		Red   Png_uint_16
		Green Png_uint_16
		Blue  Png_uint_16
		Gray  Png_uint_16
	}

	Png_text struct {
		Compression int
		Key         Png_charp
		Text        Png_charp
		Text_length Png_size_t
		Itxt_length Png_size_t
		Lang        Png_charp
		Lang_key    Png_charp
	}

	Png_unknown_chunk struct {
		Name     [5]Png_byte
		Data     *Png_byte
		Size     Png_size_t
		Location Png_byte
	}

	Png_sPLT_t struct {
		Name     Png_charp
		Depth    Png_byte
		Entries  Png_sPLT_entryp
		Nentries Png_int_32
	}

	Png_sPLT_entry struct {
		Red       Png_uint_16
		Green     Png_uint_16
		Blue      Png_uint_16
		Alpha     Png_uint_16
		Frequency Png_uint_16
	}

	Png_struct_def struct { // Deprecated
		Jmpbuf                  Jmp_buf
		Longjmp_fn              Png_longjmp_ptr
		Error_fn                Png_error_ptr
		Warning_fn              Png_error_ptr
		Error_ptr               Png_voidp
		Write_data_fn           Png_rw_ptr
		Read_data_fn            Png_rw_ptr
		Io_ptr                  Png_voidp
		Read_user_transform_fn  Png_user_transform_ptr
		Write_user_transform_fn Png_user_transform_ptr
		User_transform_ptr      Png_voidp
		User_transform_depth    Png_byte
		User_transform_channels Png_byte
		Mode                    Png_uint_32
		Flags                   Png_uint_32
		Transformations         Png_uint_32
		Zstream                 Z_stream
		Zbuf                    Png_bytep
		Zbuf_size               Png_size_t
		Zlib_level              int
		Zlib_method             int
		Zlib_window_bits        int
		Zlib_mem_level          int
		Zlib_strategy           int
		Width                   Png_uint_32
		Height                  Png_uint_32
		Num_rows                Png_uint_32
		Usr_width               Png_uint_32
		Rowbytes                Png_size_t
		User_chunk_malloc_max   Png_alloc_size_t
		Iwidth                  Png_uint_32
		Row_number              Png_uint_32
		Prev_row                Png_bytep
		Row_buf                 Png_bytep
		Sub_row                 Png_bytep
		Up_row                  Png_bytep
		Avg_row                 Png_bytep
		Paeth_row               Png_bytep
		Row_info                Png_row_info
		Idat_size               Png_uint_32
		Crc                     Png_uint_32
		Palette                 Png_colorp
		Num_palette             Png_uint_16
		Num_trans               Png_uint_16
		Chunk_name              [5]Png_byte
		Compression             Png_byte
		Filter                  Png_byte
		Interlaced              Png_byte
		Pass                    Png_byte
		Do_filter               Png_byte
		Color_type              Png_byte
		Bit_depth               Png_byte
		Usr_bit_depth           Png_byte
		Pixel_depth             Png_byte
		Channels                Png_byte
		Usr_channels            Png_byte
		Sig_bytes               Png_byte
		Filler                  Png_uint_16
		Background_gamma_type   Png_byte
		Background_gamma        float32
		Background              Png_color_16
		Background_1            Png_color_16
		Output_flush_fn         Png_flush_ptr
		Flush_dist              Png_uint_32
		Flush_rows              Png_uint_32
		Gamma_shift             int
		Gamma                   float32
		Screen_gamma            float32
		Gamma_table             Png_bytep
		Gamma_from_1            Png_bytep
		Gamma_to_1              Png_bytep
		Gamma_16_table          Png_uint_16pp
		Gamma_16_from_1         Png_uint_16pp
		Gamma_16_to_1           Png_uint_16pp
		Sig_bit                 Png_color_8
		Shift                   Png_color_8
		Trans_alpha             Png_bytep
		Trans_color             Png_color_16
		Read_row_fn             Png_read_status_ptr
		Write_row_fn            Png_write_status_ptr
		Info_fn                 Png_progressive_info_ptr
		Row_fn                  Png_progressive_row_ptr
		End_fn                  Png_progressive_end_ptr
		Save_buffer_ptr         Png_bytep
		Save_buffer             Png_bytep
		Current_buffer_ptr      Png_bytep
		Current_buffer          Png_bytep
		Push_length             Png_uint_32
		Skip_length             Png_uint_32
		Save_buffer_size        Png_size_t
		Save_buffer_max         Png_size_t
		Buffer_size             Png_size_t
		Current_buffer_size     Png_size_t
		Process_mode            int
		Cur_palette             int
		Current_text_size       Png_size_t
		Current_text_left       Png_size_t
		Current_text            Png_charp
		Current_text_ptr        Png_charp
		Palette_lookup          Png_bytep
		Quantize_index          Png_bytep
		Hist                    Png_uint_16p
		Heuristic_method        Png_byte
		Num_prev_filters        Png_byte
		Prev_filters            Png_bytep
		Filter_weights          Png_uint_16p
		Inv_filter_weights      Png_uint_16p
		Filter_costs            Png_uint_16p
		Inv_filter_costs        Png_uint_16p
		Time_buffer             Png_charp
		Free_me                 Png_uint_32
		User_chunk_ptr          Png_voidp
		Read_user_chunk_fn      Png_user_chunk_ptr
		Num_chunk_list          int
		Chunk_list              Png_bytep
		Rgb_to_gray_status      Png_byte
		Rgb_to_gray_red_coeff   Png_uint_16
		Rgb_to_gray_green_coeff Png_uint_16
		Rgb_to_gray_blue_coeff  Png_uint_16
		Mng_features_permitted  Png_uint_32
		Int_gamma               Png_fixed_point
		Filter_type             Png_byte
		Mem_ptr                 Png_voidp
		Malloc_fn               Png_malloc_ptr
		Free_fn                 Png_free_ptr
		Big_row_buf             Png_bytep
		Quantize_sort           Png_bytep
		Index_to_palette        Png_bytep
		Palette_to_index        Png_bytep
		Compression_type        Png_byte
		User_width_max          Png_uint_32
		User_height_max         Png_uint_32
		User_chunk_cache_max    Png_uint_32
		Unknown_chunk           Png_unknown_chunk
		Old_big_row_buf_size    Png_uint_32
		Old_prev_row_size       Png_uint_32
		Chunkdata               Png_charp
		Io_state                Png_uint_32
	}

	Png_info struct { // Deprecated
		Width              Png_uint_32
		Height             Png_uint_32
		Valid              Png_uint_32
		Rowbytes           Png_size_t
		Palette            Png_colorp
		Num_palette        Png_uint_16
		Num_trans          Png_uint_16
		Bit_depth          Png_byte
		Color_type         Png_byte
		Compression_type   Png_byte
		Filter_type        Png_byte
		Interlace_type     Png_byte
		Channels           Png_byte
		Pixel_depth        Png_byte
		Spare_byte         Png_byte
		Signature          [8]Png_byte
		Gamma              float32
		Srgb_intent        Png_byte
		Num_text           int
		Max_text           int
		Text               Png_textp
		Mod_time           Png_time
		Sig_bit            Png_color_8
		Trans_alpha        Png_bytep
		Trans_color        Png_color_16
		Background         Png_color_16
		X_offset           Png_int_32
		Y_offset           Png_int_32
		Offset_unit_type   Png_byte
		X_pixels_per_unit  Png_uint_32
		Y_pixels_per_unit  Png_uint_32
		Phys_unit_type     Png_byte
		Hist               Png_uint_16p
		X_white            float32
		Y_white            float32
		X_red              float32
		Y_red              float32
		X_green            float32
		Y_green            float32
		X_blue             float32
		Y_blue             float32
		Pcal_purpose       Png_charp
		Pcal_X0            Png_int_32
		Pcal_X1            Png_int_32
		Pcal_units         Png_charp
		Pcal_params        Png_charpp
		Pcal_type          Png_byte
		Pcal_nparams       Png_byte
		Free_me            Png_uint_32
		Unknown_chunks     Png_unknown_chunkp
		Unknown_chunks_num Png_size_t
		Iccp_name          Png_charp
		Iccp_profile       Png_charp
		Iccp_proflen       Png_uint_32
		Iccp_compression   Png_byte
		Splt_palettes      Png_sPLT_tp
		Splt_palettes_num  Png_uint_32
		Scal_unit          Png_byte
		Scal_pixel_width   Double
		Scal_pixel_height  Double
		Scal_s_width       Png_charp
		Scal_s_height      Png_charp
		Row_pointers       Png_bytepp
		Int_gamma          Png_fixed_point
		Int_x_white        Png_fixed_point
		Int_y_white        Png_fixed_point
		Int_x_red          Png_fixed_point
		Int_y_red          Png_fixed_point
		Int_x_green        Png_fixed_point
		Int_y_green        Png_fixed_point
		Int_x_blue         Png_fixed_point
		Int_y_blue         Png_fixed_point
	}
)

var (
	Png_access_version_number func() Png_uint_32

	Png_set_sig_bytes func(
		png_ptr Png_structp,
		num_bytes int)

	Png_sig_cmp func(
		sig Png_bytep,
		start Png_size_t,
		num_to_check Png_size_t) int

	Png_create_read_struct func(
		user_Png_ver Png_const_charp,
		error_ptr Png_voidp,
		error_fn Png_error_ptr,
		warn_fn Png_error_ptr) Png_structp

	Png_create_write_struct func(
		user_Png_ver Png_const_charp,
		error_ptr Png_voidp,
		error_fn Png_error_ptr,
		warn_fn Png_error_ptr) Png_structp

	Png_get_compression_buffer_size func(
		png_ptr Png_structp) Png_size_t

	Png_set_compression_buffer_size func(
		png_ptr Png_structp,
		size Png_size_t)

	Png_set_longjmp_fn func(
		png_ptr Png_structp,
		longjmp_fn Png_longjmp_ptr,
		Jmp_buf_size Size_t) *Jmp_buf

	Png_reset_zstream func(
		png_ptr Png_structp) int

	Png_create_read_struct_2 func(
		user_Png_ver Png_const_charp,
		error_ptr Png_voidp,
		error_fn Png_error_ptr,
		warn_fn Png_error_ptr,
		mem_ptr Png_voidp,
		malloc_fn Png_malloc_ptr,
		free_fn Png_free_ptr) Png_structp

	Png_create_write_struct_2 func(
		user_Png_ver Png_const_charp,
		error_ptr Png_voidp,
		error_fn Png_error_ptr,
		warn_fn Png_error_ptr,
		mem_ptr Png_voidp,
		malloc_fn Png_malloc_ptr,
		free_fn Png_free_ptr) Png_structp

	Png_write_sig func(
		png_ptr Png_structp)

	Png_write_chunk func(
		png_ptr Png_structp,
		chunk_name Png_bytep,
		data Png_bytep,
		length Png_size_t)

	Png_write_chunk_start func(
		png_ptr Png_structp,
		chunk_name Png_bytep,
		length Png_uint_32)

	Png_write_chunk_data func(
		png_ptr Png_structp,
		data Png_bytep,
		length Png_size_t)

	Png_write_chunk_end func(
		png_ptr Png_structp)

	Png_create_info_struct func(
		png_ptr Png_structp) Png_infop

	Png_info_init_3 func(
		info_ptr Png_infopp,
		png_info_struct_size Png_size_t)

	Png_write_info_before_PLTE func(
		png_ptr Png_structp,
		info_ptr Png_infop)

	Png_write_info func(
		png_ptr Png_structp,
		info_ptr Png_infop)

	Png_read_info func(
		png_ptr Png_structp,
		info_ptr Png_infop)

	Png_convert_to_rfc1123 func(
		png_ptr Png_structp,
		ptime Png_timep) Png_charp

	Png_convert_from_struct_tm func(
		ptime Png_timep,
		ttime *Tm)

	Png_convert_from_time_t func(
		ptime Png_timep,
		ttime Time_t)

	Png_set_expand func(
		png_ptr Png_structp)

	Png_set_expand_gray_1_2_4_to_8 func(
		png_ptr Png_structp)

	Png_set_palette_to_rgb func(
		png_ptr Png_structp)

	Png_set_tRNS_to_alpha func(
		png_ptr Png_structp)

	Png_set_bgr func(
		png_ptr Png_structp)

	Png_set_gray_to_rgb func(
		png_ptr Png_structp)

	Png_set_rgb_to_gray func(
		png_ptr Png_structp,
		error_action int,
		red,
		green Double)

	Png_set_rgb_to_gray_fixed func(
		png_ptr Png_structp,
		error_action int,
		red, green Png_fixed_point)

	Png_get_rgb_to_gray_status func(
		png_ptr Png_structp) Png_byte

	Png_build_grayscale_palette func(
		bit_depth int,
		palette Png_colorp)

	Png_set_strip_alpha func(
		png_ptr Png_structp)

	Png_set_swap_alpha func(
		png_ptr Png_structp)

	Png_set_invert_alpha func(
		png_ptr Png_structp)

	Png_set_filler func(
		png_ptr Png_structp,
		filler Png_uint_32,
		flags int)

	Png_set_add_alpha func(
		png_ptr Png_structp,
		filler Png_uint_32,
		flags int)

	Png_set_swap func(
		png_ptr Png_structp)

	Png_set_packing func(
		png_ptr Png_structp)

	Png_set_packswap func(
		png_ptr Png_structp)

	Png_set_shift func(
		png_ptr Png_structp,
		true_bits Png_color_8p)

	Png_set_interlace_handling func(
		png_ptr Png_structp) int

	Png_set_invert_mono func(
		png_ptr Png_structp)

	Png_set_background func(
		png_ptr Png_structp,
		background_color Png_color_16p,
		background_gamma_code int,
		need_expand int,
		background_gamma Double)

	Png_set_strip_16 func(
		png_ptr Png_structp)

	Png_set_quantize func(
		png_ptr Png_structp,
		palette Png_colorp,
		num_palette int,
		maximum_colors int,
		histogram Png_uint_16p,
		full_quantize int)

	Png_set_gamma func(
		png_ptr Png_structp,
		screen_gamma Double,
		default_file_gamma Double)

	Png_set_flush func(
		png_ptr Png_structp,
		nrows int)

	Png_write_flush func(
		png_ptr Png_structp)

	Png_start_read_image func(
		png_ptr Png_structp)

	Png_read_update_info func(
		png_ptr Png_structp,
		info_ptr Png_infop)

	Png_read_rows func(
		png_ptr Png_structp,
		row Png_bytepp,
		display_row Png_bytepp,
		num_rows Png_uint_32)

	Png_read_row func(
		png_ptr Png_structp,
		row Png_bytep,
		display_row Png_bytep)

	Png_read_image func(
		png_ptr Png_structp,
		image Png_bytepp)

	Png_write_row func(
		png_ptr Png_structp,
		row Png_bytep)

	Png_write_rows func(
		png_ptr Png_structp,
		row Png_bytepp,
		num_rows Png_uint_32)

	Png_write_image func(
		png_ptr Png_structp,
		image Png_bytepp)

	Png_write_end func(
		png_ptr Png_structp,
		info_ptr Png_infop)

	Png_read_end func(
		png_ptr Png_structp,
		info_ptr Png_infop)

	Png_destroy_info_struct func(
		png_ptr Png_structp,
		info_ptr_ptr Png_infopp)

	Png_destroy_read_struct func(
		png_ptr_ptr Png_structpp,
		info_ptr_ptr Png_infopp,
		end_info_ptr_ptr Png_infopp)

	Png_destroy_write_struct func(
		png_ptr_ptr Png_structpp,
		info_ptr_ptr Png_infopp)

	Png_set_crc_action func(
		png_ptr Png_structp,
		crit_action int,
		ancil_action int)

	Png_set_filter func(
		png_ptr Png_structp,
		method int,
		filters int)

	Png_set_filter_heuristics func(
		png_ptr Png_structp,
		heuristic_method int,
		num_weights int,
		filter_weights Png_doublep,
		filter_costs Png_doublep)

	Png_set_compression_level func(
		png_ptr Png_structp,
		level int)

	Png_set_compression_mem_level func(
		png_ptr Png_structp,
		mem_level int)

	Png_set_compression_strategy func(
		png_ptr Png_structp,
		strategy int)

	Png_set_compression_window_bits func(
		png_ptr Png_structp,
		window_bits int)

	Png_set_compression_method func(
		png_ptr Png_structp,
		method int)

	Png_init_io func(
		png_ptr Png_structp,
		fp Png_FILE_p)

	Png_set_error_fn func(
		png_ptr Png_structp,
		error_ptr Png_voidp,
		error_fn Png_error_ptr,
		warning_fn Png_error_ptr)

	Png_get_error_ptr func(
		png_ptr Png_structp) Png_voidp

	Png_set_write_fn func(
		png_ptr Png_structp,
		io_ptr Png_voidp,
		write_data_fn Png_rw_ptr,
		output_flush_fn Png_flush_ptr)

	Png_set_read_fn func(
		png_ptr Png_structp,
		io_ptr Png_voidp,
		read_data_fn Png_rw_ptr)

	Png_get_io_ptr func(
		png_ptr Png_structp) Png_voidp

	Png_set_read_status_fn func(
		png_ptr Png_structp,
		read_row_fn Png_read_status_ptr)

	Png_set_write_status_fn func(
		png_ptr Png_structp,
		write_row_fn Png_write_status_ptr)

	Png_set_mem_fn func(
		png_ptr Png_structp,
		mem_ptr Png_voidp,
		malloc_fn Png_malloc_ptr,
		free_fn Png_free_ptr)

	Png_get_mem_ptr func(
		png_ptr Png_structp) Png_voidp

	Png_set_read_user_transform_fn func(
		png_ptr Png_structp,
		read_user_transform_fn Png_user_transform_ptr)

	Png_set_write_user_transform_fn func(
		png_ptr Png_structp,
		write_user_transform_fn Png_user_transform_ptr)

	Png_set_user_transform_info func(
		png_ptr Png_structp,
		user_transform_ptr Png_voidp,
		user_transform_depth int,
		user_transform_channels int)

	Png_get_user_transform_ptr func(
		png_ptr Png_structp) Png_voidp

	Png_set_read_user_chunk_fn func(
		png_ptr Png_structp,
		user_chunk_ptr Png_voidp,
		read_user_chunk_fn Png_user_chunk_ptr)
	Png_get_user_chunk_ptr func(
		png_ptr Png_structp) Png_voidp

	Png_set_progressive_read_fn func(
		png_ptr Png_structp,
		progressive_ptr Png_voidp,
		info_fn Png_progressive_info_ptr,
		row_fn Png_progressive_row_ptr,
		end_fn Png_progressive_end_ptr)

	Png_get_progressive_ptr func(
		png_ptr Png_structp) Png_voidp

	Png_process_data func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		buffer Png_bytep,
		buffer_size Png_size_t)

	Png_progressive_combine_row func(
		png_ptr Png_structp,
		old_row Png_bytep,
		new_row Png_bytep)

	Png_malloc func(
		png_ptr Png_structp,
		size Png_alloc_size_t) Png_voidp

	Png_calloc func(
		png_ptr Png_structp,
		size Png_alloc_size_t) Png_voidp

	Png_malloc_warn func(
		png_ptr Png_structp,
		size Png_alloc_size_t) Png_voidp

	Png_free func(
		png_ptr Png_structp,
		ptr Png_voidp)

	Png_free_data func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		free_me Png_uint_32,
		num int)

	Png_data_freer func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		freer int,
		mask Png_uint_32)

	Png_malloc_default func(
		png_ptr Png_structp,
		size Png_alloc_size_t) Png_voidp

	Png_free_default func(
		png_ptr Png_structp,
		ptr Png_voidp)

	Png_error func(
		png_ptr Png_structp,
		error_message Png_const_charp)

	Png_chunk_error func(
		png_ptr Png_structp,
		error_message Png_const_charp)

	Png_warning func(
		png_ptr Png_structp,
		warning_message Png_const_charp)

	Png_chunk_warning func(
		png_ptr Png_structp,
		warning_message Png_const_charp)

	Png_get_valid func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		flag Png_uint_32) Png_uint_32

	Png_get_rowbytes func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_size_t

	Png_get_rows func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_bytepp

	Png_set_rows func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		row_pointers Png_bytepp)

	Png_get_channels func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_byte

	Png_get_image_width func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_uint_32

	Png_get_image_height func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_uint_32

	Png_get_bit_depth func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_byte

	Png_get_color_type func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_byte

	Png_get_filter_type func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_byte

	Png_get_interlace_type func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_byte

	Png_get_compression_type func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_byte

	Png_get_pixels_per_meter func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_uint_32

	Png_get_x_pixels_per_meter func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_uint_32

	Png_get_y_pixels_per_meter func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_uint_32

	Png_get_pixel_aspect_ratio func(
		png_ptr Png_structp,
		info_ptr Png_infop) float32

	Png_get_x_offset_pixels func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_int_32

	Png_get_y_offset_pixels func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_int_32

	Png_get_x_offset_microns func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_int_32

	Png_get_y_offset_microns func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_int_32

	Png_get_signature func(
		png_ptr Png_structp,
		info_ptr Png_infop) Png_bytep

	Png_get_bKGD func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		background *Png_color_16p) Png_uint_32

	Png_set_bKGD func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		background Png_color_16p)

	Png_get_cHRM func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		white_x, white_y,
		red_x, red_y,
		green_x, green_y,
		blue_x, blue_y *Double) Png_uint_32

	Png_get_cHRM_fixed func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		int_white_x *Png_fixed_point,
		int_white_y *Png_fixed_point,
		int_red_x *Png_fixed_point,
		int_red_y *Png_fixed_point,
		int_green_x *Png_fixed_point,
		int_green_y *Png_fixed_point,
		int_blue_x *Png_fixed_point,
		int_blue_y *Png_fixed_point) Png_uint_32

	Png_set_cHRM func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		white_x, white_y,
		red_x, red_y,
		green_x, green_y,
		blue_x, blue_y Double)

	Png_set_cHRM_fixed func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		int_white_x Png_fixed_point,
		int_white_y Png_fixed_point,
		int_red_x Png_fixed_point,
		int_red_y Png_fixed_point,
		int_green_x Png_fixed_point,
		int_green_y Png_fixed_point,
		int_blue_x Png_fixed_point,
		int_blue_y Png_fixed_point)

	Png_get_gAMA func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		file_gamma *Double) Png_uint_32

	Png_get_gAMA_fixed func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		int_file_gamma *Png_fixed_point) Png_uint_32

	Png_set_gAMA func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		file_gamma Double)

	Png_set_gAMA_fixed func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		int_file_gamma Png_fixed_point)

	Png_get_hIST func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		hist *Png_uint_16p) Png_uint_32

	Png_set_hIST func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		hist Png_uint_16p)

	Png_get_IHDR func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		width *Png_uint_32,
		height *Png_uint_32,
		bit_depth *int,
		color_type *int,
		interlace_method *int,
		compression_method *int,
		filter_method *int) Png_uint_32

	Png_set_IHDR func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		width Png_uint_32,
		height Png_uint_32,
		bit_depth int,
		color_type int,
		interlace_method int,
		compression_method int,
		filter_method int)

	Png_get_oFFs func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		offset_x *Png_int_32,
		offset_y *Png_int_32,
		unit_type *int) Png_uint_32

	Png_set_oFFs func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		offset_x Png_int_32,
		offset_y Png_int_32,
		unit_type int)

	Png_get_pCAL func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		purpose *Png_charp,
		X0 *Png_int_32,
		X1 *Png_int_32,
		typ *int,
		nparams *int,
		units *Png_charp,
		params *Png_charpp) Png_uint_32

	Png_set_pCAL func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		purpose Png_charp,
		X0 Png_int_32,
		X1 Png_int_32,
		typ int,
		nparams int,
		units Png_charp,
		params Png_charpp)

	Png_get_pHYs func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		res_x *Png_uint_32,
		res_y *Png_uint_32,
		unit_type *int) Png_uint_32

	Png_set_pHYs func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		res_x Png_uint_32,
		res_y Png_uint_32,
		unit_type int)

	Png_get_PLTE func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		palette *Png_colorp,
		num_palette *int) Png_uint_32

	Png_set_PLTE func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		palette Png_colorp,
		num_palette int)

	Png_get_sBIT func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		sig_bit *Png_color_8p) Png_uint_32

	Png_set_sBIT func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		sig_bit Png_color_8p)

	Png_get_sRGB func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		intent *int) Png_uint_32

	Png_set_sRGB func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		intent int)

	Png_set_sRGB_gAMA_and_cHRM func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		intent int)

	Png_get_iCCP func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		name Png_charpp,
		compression_type *int,
		profile Png_charpp,
		proflen *Png_uint_32) Png_uint_32

	Png_set_iCCP func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		name Png_charp,
		compression_type int,
		profile Png_charp,
		proflen Png_uint_32)

	Png_get_sPLT func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		entries Png_sPLT_tpp) Png_uint_32

	Png_set_sPLT func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		entries Png_sPLT_tp,
		nentries int)

	Png_get_text func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		text_ptr *Png_textp,
		num_text *int) Png_uint_32

	Png_set_text func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		text_ptr Png_textp,
		num_text int)

	Png_get_tIME func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		mod_time *Png_timep) Png_uint_32

	Png_set_tIME func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		mod_time Png_timep)

	Png_get_tRNS func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		trans_alpha *Png_bytep,
		num_trans *int,
		trans_color *Png_color_16p) Png_uint_32

	Png_set_tRNS func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		trans_alpha Png_bytep,
		num_trans int,
		trans_color Png_color_16p)

	Png_get_sCAL func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		unit *int,
		width, height *Double) Png_uint_32

	Png_set_sCAL func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		unit int,
		width, height Double)

	Png_set_keep_unknown_chunks func(
		png_ptr Png_structp,
		Keep int,
		chunk_list Png_bytep,
		num_chunks int)

	Png_handle_as_unknown func(
		png_ptr Png_structp,
		chunk_name Png_bytep) int

	Png_set_unknown_chunks func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		unknowns Png_unknown_chunkp,
		num_unknowns int)

	Png_set_unknown_chunk_location func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		chunk int,
		location int)

	Png_get_unknown_chunks func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		entries Png_unknown_chunkpp) Png_uint_32

	Png_set_invalid func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		mask int)

	Png_read_png func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		transforms int,
		params Png_voidp)

	Png_write_png func(
		png_ptr Png_structp,
		info_ptr Png_infop,
		transforms int,
		params Png_voidp)

	Png_get_copyright func(
		png_ptr Png_structp) Png_charp

	Png_get_header_ver func(
		png_ptr Png_structp) Png_charp

	Png_get_header_version func(
		png_ptr Png_structp) Png_charp

	Png_get_libpng_ver func(
		png_ptr Png_structp) Png_charp

	Png_permit_mng_features func(
		png_ptr Png_structp,
		mng_features_permitted Png_uint_32) Png_uint_32

	Png_set_user_limits func(
		png_ptr Png_structp,
		user_width_max Png_uint_32,
		user_height_max Png_uint_32)

	Png_get_user_width_max func(
		png_ptr Png_structp) Png_uint_32

	Png_get_user_height_max func(
		png_ptr Png_structp) Png_uint_32

	Png_set_chunk_cache_max func(
		png_ptr Png_structp,
		user_chunk_cache_max Png_uint_32)

	Png_get_chunk_cache_max func(
		png_ptr Png_structp) Png_uint_32

	Png_set_chunk_malloc_max func(
		png_ptr Png_structp,
		user_chunk_cache_max Png_alloc_size_t)

	Png_get_chunk_malloc_max func(
		png_ptr Png_structp) Png_alloc_size_t

	Png_get_io_state func(
		png_ptr Png_structp) Png_uint_32

	Png_get_io_chunk_name func(
		png_ptr Png_structp) Png_bytep

	Png_get_uint_31 func(
		png_ptr Png_structp,
		buf Png_bytep) Png_uint_32

	Png_save_uint_32 func(
		buf Png_bytep,
		i Png_uint_32)

	Png_save_int_32 func(
		buf Png_bytep,
		i Png_int_32)

	Png_save_uint_16 func(
		buf Png_bytep,
		i Unsigned_int)
)

var dll = "libpng14-14.dll"

var apiList = outside.Apis{
	{"png_access_version_number", &Png_access_version_number},
	{"png_build_grayscale_palette", &Png_build_grayscale_palette},
	{"png_calloc", &Png_calloc},
	{"png_chunk_error", &Png_chunk_error},
	{"png_chunk_warning", &Png_chunk_warning},
	{"png_convert_from_struct_tm", &Png_convert_from_struct_tm},
	{"png_convert_from_time_t", &Png_convert_from_time_t},
	{"png_convert_to_rfc1123", &Png_convert_to_rfc1123},
	{"png_create_info_struct", &Png_create_info_struct},
	{"png_create_read_struct", &Png_create_read_struct},
	{"png_create_read_struct_2", &Png_create_read_struct_2},
	{"png_create_write_struct", &Png_create_write_struct},
	{"png_create_write_struct_2", &Png_create_write_struct_2},
	{"png_data_freer", &Png_data_freer},
	{"png_destroy_info_struct", &Png_destroy_info_struct},
	{"png_destroy_read_struct", &Png_destroy_read_struct},
	{"png_destroy_write_struct", &Png_destroy_write_struct},
	{"png_error", &Png_error},
	{"png_free", &Png_free},
	{"png_free_data", &Png_free_data},
	{"png_free_default", &Png_free_default},
	{"png_get_IHDR", &Png_get_IHDR},
	{"png_get_PLTE", &Png_get_PLTE},
	{"png_get_bKGD", &Png_get_bKGD},
	{"png_get_bit_depth", &Png_get_bit_depth},
	{"png_get_cHRM", &Png_get_cHRM},
	{"png_get_cHRM_fixed", &Png_get_cHRM_fixed},
	{"png_get_channels", &Png_get_channels},
	{"png_get_chunk_cache_max", &Png_get_chunk_cache_max},
	{"png_get_chunk_malloc_max", &Png_get_chunk_malloc_max},
	{"png_get_color_type", &Png_get_color_type},
	{"png_get_compression_buffer_size", &Png_get_compression_buffer_size},
	{"png_get_compression_type", &Png_get_compression_type},
	{"png_get_copyright", &Png_get_copyright},
	{"png_get_error_ptr", &Png_get_error_ptr},
	{"png_get_filter_type", &Png_get_filter_type},
	{"png_get_gAMA", &Png_get_gAMA},
	{"png_get_gAMA_fixed", &Png_get_gAMA_fixed},
	{"png_get_hIST", &Png_get_hIST},
	{"png_get_header_ver", &Png_get_header_ver},
	{"png_get_header_version", &Png_get_header_version},
	{"png_get_iCCP", &Png_get_iCCP},
	{"png_get_image_height", &Png_get_image_height},
	{"png_get_image_width", &Png_get_image_width},
	{"png_get_interlace_type", &Png_get_interlace_type},
	{"png_get_io_chunk_name", &Png_get_io_chunk_name},
	{"png_get_io_ptr", &Png_get_io_ptr},
	{"png_get_io_state", &Png_get_io_state},
	{"png_get_libpng_ver", &Png_get_libpng_ver},
	{"png_get_mem_ptr", &Png_get_mem_ptr},
	{"png_get_oFFs", &Png_get_oFFs},
	{"png_get_pCAL", &Png_get_pCAL},
	{"png_get_pHYs", &Png_get_pHYs},
	{"png_get_pixel_aspect_ratio", &Png_get_pixel_aspect_ratio},
	{"png_get_pixels_per_meter", &Png_get_pixels_per_meter},
	{"png_get_progressive_ptr", &Png_get_progressive_ptr},
	{"png_get_rgb_to_gray_status", &Png_get_rgb_to_gray_status},
	{"png_get_rowbytes", &Png_get_rowbytes},
	{"png_get_rows", &Png_get_rows},
	{"png_get_sBIT", &Png_get_sBIT},
	{"png_get_sCAL", &Png_get_sCAL},
	{"png_get_sPLT", &Png_get_sPLT},
	{"png_get_sRGB", &Png_get_sRGB},
	{"png_get_signature", &Png_get_signature},
	{"png_get_tIME", &Png_get_tIME},
	{"png_get_tRNS", &Png_get_tRNS},
	{"png_get_text", &Png_get_text},
	{"png_get_uint_31", &Png_get_uint_31},
	{"png_get_unknown_chunks", &Png_get_unknown_chunks},
	{"png_get_user_chunk_ptr", &Png_get_user_chunk_ptr},
	{"png_get_user_height_max", &Png_get_user_height_max},
	{"png_get_user_transform_ptr", &Png_get_user_transform_ptr},
	{"png_get_user_width_max", &Png_get_user_width_max},
	{"png_get_valid", &Png_get_valid},
	{"png_get_x_offset_microns", &Png_get_x_offset_microns},
	{"png_get_x_offset_pixels", &Png_get_x_offset_pixels},
	{"png_get_x_pixels_per_meter", &Png_get_x_pixels_per_meter},
	{"png_get_y_offset_microns", &Png_get_y_offset_microns},
	{"png_get_y_offset_pixels", &Png_get_y_offset_pixels},
	{"png_get_y_pixels_per_meter", &Png_get_y_pixels_per_meter},
	{"png_handle_as_unknown", &Png_handle_as_unknown},
	{"png_info_init_3", &Png_info_init_3},
	{"png_init_io", &Png_init_io},
	{"png_malloc", &Png_malloc},
	{"png_malloc_default", &Png_malloc_default},
	{"png_malloc_warn", &Png_malloc_warn},
	{"png_permit_mng_features", &Png_permit_mng_features},
	{"png_process_data", &Png_process_data},
	{"png_progressive_combine_row", &Png_progressive_combine_row},
	{"png_read_end", &Png_read_end},
	{"png_read_image", &Png_read_image},
	{"png_read_info", &Png_read_info},
	{"png_read_png", &Png_read_png},
	{"png_read_row", &Png_read_row},
	{"png_read_rows", &Png_read_rows},
	{"png_read_update_info", &Png_read_update_info},
	{"png_reset_zstream", &Png_reset_zstream},
	{"png_save_int_32", &Png_save_int_32},
	{"png_save_uint_16", &Png_save_uint_16},
	{"png_save_uint_32", &Png_save_uint_32},
	{"png_set_IHDR", &Png_set_IHDR},
	{"png_set_PLTE", &Png_set_PLTE},
	{"png_set_add_alpha", &Png_set_add_alpha},
	{"png_set_bKGD", &Png_set_bKGD},
	{"png_set_background", &Png_set_background},
	{"png_set_bgr", &Png_set_bgr},
	{"png_set_cHRM", &Png_set_cHRM},
	{"png_set_cHRM_fixed", &Png_set_cHRM_fixed},
	{"png_set_chunk_cache_max", &Png_set_chunk_cache_max},
	{"png_set_chunk_malloc_max", &Png_set_chunk_malloc_max},
	{"png_set_compression_buffer_size", &Png_set_compression_buffer_size},
	{"png_set_compression_level", &Png_set_compression_level},
	{"png_set_compression_mem_level", &Png_set_compression_mem_level},
	{"png_set_compression_method", &Png_set_compression_method},
	{"png_set_compression_strategy", &Png_set_compression_strategy},
	{"png_set_compression_window_bits", &Png_set_compression_window_bits},
	{"png_set_crc_action", &Png_set_crc_action},
	{"png_set_error_fn", &Png_set_error_fn},
	{"png_set_expand", &Png_set_expand},
	{"png_set_expand_gray_1_2_4_to_8", &Png_set_expand_gray_1_2_4_to_8},
	{"png_set_filler", &Png_set_filler},
	{"png_set_filter", &Png_set_filter},
	{"png_set_filter_heuristics", &Png_set_filter_heuristics},
	{"png_set_flush", &Png_set_flush},
	{"png_set_gAMA", &Png_set_gAMA},
	{"png_set_gAMA_fixed", &Png_set_gAMA_fixed},
	{"png_set_gamma", &Png_set_gamma},
	{"png_set_gray_to_rgb", &Png_set_gray_to_rgb},
	{"png_set_hIST", &Png_set_hIST},
	{"png_set_iCCP", &Png_set_iCCP},
	{"png_set_interlace_handling", &Png_set_interlace_handling},
	{"png_set_invalid", &Png_set_invalid},
	{"png_set_invert_alpha", &Png_set_invert_alpha},
	{"png_set_invert_mono", &Png_set_invert_mono},
	{"png_set_keep_unknown_chunks", &Png_set_keep_unknown_chunks},
	{"png_set_longjmp_fn", &Png_set_longjmp_fn},
	{"png_set_mem_fn", &Png_set_mem_fn},
	{"png_set_oFFs", &Png_set_oFFs},
	{"png_set_pCAL", &Png_set_pCAL},
	{"png_set_pHYs", &Png_set_pHYs},
	{"png_set_packing", &Png_set_packing},
	{"png_set_packswap", &Png_set_packswap},
	{"png_set_palette_to_rgb", &Png_set_palette_to_rgb},
	{"png_set_progressive_read_fn", &Png_set_progressive_read_fn},
	{"png_set_quantize", &Png_set_quantize},
	{"png_set_read_fn", &Png_set_read_fn},
	{"png_set_read_status_fn", &Png_set_read_status_fn},
	{"png_set_read_user_chunk_fn", &Png_set_read_user_chunk_fn},
	{"png_set_read_user_transform_fn", &Png_set_read_user_transform_fn},
	{"png_set_rgb_to_gray", &Png_set_rgb_to_gray},
	{"png_set_rgb_to_gray_fixed", &Png_set_rgb_to_gray_fixed},
	{"png_set_rows", &Png_set_rows},
	{"png_set_sBIT", &Png_set_sBIT},
	{"png_set_sCAL", &Png_set_sCAL},
	{"png_set_sPLT", &Png_set_sPLT},
	{"png_set_sRGB", &Png_set_sRGB},
	{"png_set_sRGB_gAMA_and_cHRM", &Png_set_sRGB_gAMA_and_cHRM},
	{"png_set_shift", &Png_set_shift},
	{"png_set_sig_bytes", &Png_set_sig_bytes},
	{"png_set_strip_16", &Png_set_strip_16},
	{"png_set_strip_alpha", &Png_set_strip_alpha},
	{"png_set_swap", &Png_set_swap},
	{"png_set_swap_alpha", &Png_set_swap_alpha},
	{"png_set_tIME", &Png_set_tIME},
	{"png_set_tRNS", &Png_set_tRNS},
	{"png_set_tRNS_to_alpha", &Png_set_tRNS_to_alpha},
	{"png_set_text", &Png_set_text},
	{"png_set_unknown_chunk_location", &Png_set_unknown_chunk_location},
	{"png_set_unknown_chunks", &Png_set_unknown_chunks},
	{"png_set_user_limits", &Png_set_user_limits},
	{"png_set_user_transform_info", &Png_set_user_transform_info},
	{"png_set_write_fn", &Png_set_write_fn},
	{"png_set_write_status_fn", &Png_set_write_status_fn},
	{"png_set_write_user_transform_fn", &Png_set_write_user_transform_fn},
	{"png_sig_cmp", &Png_sig_cmp},
	{"png_start_read_image", &Png_start_read_image},
	{"png_warning", &Png_warning},
	{"png_write_chunk", &Png_write_chunk},
	{"png_write_chunk_data", &Png_write_chunk_data},
	{"png_write_chunk_end", &Png_write_chunk_end},
	{"png_write_chunk_start", &Png_write_chunk_start},
	{"png_write_end", &Png_write_end},
	{"png_write_flush", &Png_write_flush},
	{"png_write_image", &Png_write_image},
	{"png_write_info", &Png_write_info},
	{"png_write_info_before_PLTE", &Png_write_info_before_PLTE},
	{"png_write_png", &Png_write_png},
	{"png_write_row", &Png_write_row},
	{"png_write_rows", &Png_write_rows},
	{"png_write_sig", &Png_write_sig},
}
