package glib

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
	AddDllApis(dllThread, false, apiListThread)
	AddDllData(dll, false, dataList)
}

type (
	//TODO(t):Fix (stat/stat32)
	GStatBuf         struct{}
	FILE             struct{}
	Utimbuf          struct{}
	G_thread_gettime func() Guint64
)

var (
	G_array_new func(zero_terminated Gboolean,
		clear_ Gboolean,
		element_size Guint) *GArray

	G_array_sized_new func(zero_terminated Gboolean,
		clear_ Gboolean,
		element_size Guint,
		reserved_size Guint) *GArray

	G_array_free func(array *GArray,
		free_segment Gboolean) string

	G_array_ref func(array *GArray) *GArray

	G_array_unref func(array *GArray)

	G_array_get_element_size func(array *GArray) Guint

	G_array_append_vals func(array *GArray,
		data Gconstpointer,
		leng Guint) *GArray

	G_array_prepend_vals func(array *GArray,
		data Gconstpointer,
		leng Guint) *GArray

	G_array_insert_vals func(array *GArray,
		index_ Guint,
		data Gconstpointer,
		leng Guint) *GArray

	G_array_set_size func(array *GArray,
		length Guint) *GArray

	G_array_remove_index func(array *GArray,
		index_ Guint) *GArray

	G_array_remove_index_fast func(array *GArray,
		index_ Guint) *GArray

	G_array_remove_range func(array *GArray,
		index_ Guint,
		length Guint) *GArray

	G_array_sort func(array *GArray,
		compare_func GCompareFunc)

	G_array_sort_with_data func(array *GArray,
		compare_func GCompareDataFunc,
		user_data Gpointer)

	G_ptr_array_new func() *GPtrArray

	G_ptr_array_new_with_free_func func(element_free_func GDestroyNotify) *GPtrArray

	G_ptr_array_sized_new func(reserved_size Guint) *GPtrArray

	G_ptr_array_free func(array *GPtrArray,
		free_seg Gboolean) *Gpointer

	G_ptr_array_ref func(array *GPtrArray) *GPtrArray

	G_ptr_array_unref func(array *GPtrArray)

	G_ptr_array_set_free_func func(array *GPtrArray,
		element_free_func GDestroyNotify)

	G_ptr_array_set_size func(array *GPtrArray,
		length Gint)

	G_ptr_array_remove_index func(array *GPtrArray,
		index_ Guint) Gpointer

	G_ptr_array_remove_index_fast func(array *GPtrArray,
		index_ Guint) Gpointer

	G_ptr_array_remove func(array *GPtrArray,
		data Gpointer) Gboolean

	G_ptr_array_remove_fast func(array *GPtrArray,
		data Gpointer) Gboolean

	G_ptr_array_remove_range func(array *GPtrArray,
		index_ Guint,
		length Guint)

	G_ptr_array_add func(array *GPtrArray,
		data Gpointer)

	G_ptr_array_sort func(array *GPtrArray,
		compare_func GCompareFunc)

	G_ptr_array_sort_with_data func(array *GPtrArray,
		compare_func GCompareDataFunc,
		user_data Gpointer)

	G_ptr_array_foreach func(array *GPtrArray,
		f GFunc,
		user_data Gpointer)

	G_byte_array_new func() *GByteArray

	G_byte_array_sized_new func(reserved_size Guint) *GByteArray

	G_byte_array_free func(array *GByteArray,
		free_segment Gboolean) *Guint8

	G_byte_array_ref func(array *GByteArray) *GByteArray

	G_byte_array_unref func(array *GByteArray)

	G_byte_array_append func(array *GByteArray,
		data *Guint8,
		leng Guint) *GByteArray

	G_byte_array_prepend func(array *GByteArray,
		data *Guint8,
		leng Guint) *GByteArray

	G_byte_array_set_size func(array *GByteArray,
		length Guint) *GByteArray

	G_byte_array_remove_index func(array *GByteArray,
		index_ Guint) *GByteArray

	G_byte_array_remove_index_fast func(array *GByteArray,
		index_ Guint) *GByteArray

	G_byte_array_remove_range func(array *GByteArray,
		index_ Guint,
		length Guint) *GByteArray

	G_byte_array_sort func(array *GByteArray,
		compare_func GCompareFunc)

	G_byte_array_sort_with_data func(array *GByteArray,
		compare_func GCompareDataFunc,
		user_data Gpointer)

	G_quark_try_string func(str string) GQuark

	G_quark_from_static_string func(str string) GQuark

	G_quark_from_string func(str string) GQuark

	G_quark_to_string func(quark GQuark) string

	G_intern_string func(str string) string

	G_intern_static_string func(str string) string

	G_error_new func(domain GQuark, code Gint, format string,
		v ...VArg) *GError

	G_error_new_literal func(domain GQuark,
		code Gint,
		message string) *GError

	G_error_new_valist func(domain GQuark,
		code Gint,
		format string,
		args Va_list) *GError

	G_error_free func(e *GError)

	G_error_copy func(e *GError) *GError

	G_error_matches func(e *GError,
		domain GQuark,
		code Gint) Gboolean

	G_set_error func(err **GError, domain GQuark, code Gint,
		format string, v ...VArg)

	G_set_error_literal func(err **GError,
		domain GQuark,
		code Gint,
		message string)

	G_propagate_error func(dest **GError,
		src *GError)

	G_clear_error func(err **GError)

	G_prefix_error func(err **GError, format string, v ...VArg)

	G_propagate_prefixed_error func(dest **GError, src *GError,
		format string, v ...VArg)

	G_get_user_name func() string

	G_get_real_name func() string

	G_get_home_dir func() string

	G_get_tmp_dir func() string

	G_get_host_name func() string

	G_get_prgname func() string

	G_set_prgname func(prgname string)

	G_get_application_name func() string

	G_set_application_name func(application_name string)

	G_reload_user_special_dirs_cache func()

	G_get_user_data_dir func() string

	G_get_user_config_dir func() string

	G_get_user_cache_dir func() string

	G_get_system_data_dirs func() **Gchar

	G_win32_get_system_data_dirs_for_module func(
		f func()) **Gchar

	G_parse_debug_string func(
		str string, keys *GDebugKey, nkeys Guint) Guint

	G_snprintf func(str string, n Gulong, format string,
		v ...VArg) Gint

	G_vsnprintf func(str string,
		n Gulong,
		format string,
		args Va_list) Gint

	G_path_is_absolute func(file_name string) Gboolean

	G_path_skip_root func(file_name string) string

	G_basename func(file_name string) string

	G_get_current_dir func() string

	G_path_get_basename func(file_name string) string

	G_path_get_dirname func(file_name string) string

	G_nullify_pointer func(nullify_location *Gpointer)

	G_getenv func(variable string) string

	G_setenv func(variable string,
		value string,
		overwrite Gboolean) Gboolean

	G_unsetenv func(variable string)

	G_listenv func() **Gchar

	G_get_environ func() **Gchar

	G_thread_init func(vtable *GThreadFunctions)

	G_thread_init_with_errorcheck_mutexes func(vtable *GThreadFunctions)

	G_thread_get_initialized func() Gboolean

	G_static_mutex_get_mutex_impl func(mutex **GMutex) *GMutex

	G_thread_create_full func(f GThreadFunc,
		data Gpointer,
		stack_size Gulong,
		joinable Gboolean,
		bound Gboolean,
		priority GThreadPriority,
		e **GError) *GThread

	G_thread_self func() *GThread

	G_thread_exit func(retval Gpointer)

	G_thread_join func(thread *GThread) Gpointer

	G_thread_set_priority func(thread *GThread,
		priority GThreadPriority)

	G_static_mutex_init func(mutex *GStaticMutex)

	G_static_mutex_free func(mutex *GStaticMutex)

	G_static_rec_mutex_init func(mutex *GStaticRecMutex)

	G_static_rec_mutex_lock func(mutex *GStaticRecMutex)

	G_static_rec_mutex_trylock func(mutex *GStaticRecMutex) Gboolean

	G_static_rec_mutex_unlock func(mutex *GStaticRecMutex)

	G_static_rec_mutex_lock_full func(mutex *GStaticRecMutex,
		depth Guint)

	G_static_rec_mutex_unlock_full func(mutex *GStaticRecMutex) Guint

	G_static_rec_mutex_free func(mutex *GStaticRecMutex)

	G_static_rw_lock_init func(
		lock *GStaticRWLock)

	G_static_rw_lock_reader_lock func(
		lock *GStaticRWLock)

	G_static_rw_lock_reader_trylock func(
		lock *GStaticRWLock) Gboolean

	G_static_rw_lock_reader_unlock func(
		lock *GStaticRWLock)

	G_static_rw_lock_writer_lock func(
		lock *GStaticRWLock)

	G_static_rw_lock_writer_trylock func(
		lock *GStaticRWLock) Gboolean

	G_static_rw_lock_writer_unlock func(
		lock *GStaticRWLock)

	G_static_rw_lock_free func(
		lock *GStaticRWLock)

	G_thread_foreach func(thread_func GFunc,
		user_data Gpointer)

	G_once_impl func(
		once *GOnce,
		f GThreadFunc,
		arg Gpointer) Gpointer

	G_once_init_enter func(
		value_location *Gsize) Gboolean

	G_once_init_enter_impl func(
		value_location *Gsize) Gboolean

	G_once_init_leave func(
		value_location *Gsize,
		initialization_value Gsize)

	G_async_queue_new func() *GAsyncQueue

	G_async_queue_new_full func(item_free_func GDestroyNotify) *GAsyncQueue

	G_async_queue_lock func(queue *GAsyncQueue)

	G_async_queue_unlock func(queue *GAsyncQueue)

	G_async_queue_ref func(queue *GAsyncQueue) *GAsyncQueue

	G_async_queue_unref func(queue *GAsyncQueue)

	G_async_queue_ref_unlocked func(queue *GAsyncQueue)

	G_async_queue_unref_and_unlock func(queue *GAsyncQueue)

	G_async_queue_push func(queue *GAsyncQueue,
		data Gpointer)

	G_async_queue_push_unlocked func(queue *GAsyncQueue,
		data Gpointer)

	G_async_queue_push_sorted func(queue *GAsyncQueue,
		data Gpointer,
		f GCompareDataFunc,
		user_data Gpointer)

	G_async_queue_push_sorted_unlocked func(queue *GAsyncQueue,
		data Gpointer,
		f GCompareDataFunc,
		user_data Gpointer)

	G_async_queue_pop func(queue *GAsyncQueue) Gpointer

	G_async_queue_pop_unlocked func(queue *GAsyncQueue) Gpointer

	G_async_queue_try_pop func(queue *GAsyncQueue) Gpointer

	G_async_queue_try_pop_unlocked func(queue *GAsyncQueue) Gpointer

	G_async_queue_timed_pop func(queue *GAsyncQueue,
		end_time *GTimeVal) Gpointer

	G_async_queue_timed_pop_unlocked func(queue *GAsyncQueue,
		end_time *GTimeVal) Gpointer

	G_async_queue_length func(queue *GAsyncQueue) Gint

	G_async_queue_length_unlocked func(queue *GAsyncQueue) Gint

	G_async_queue_sort func(queue *GAsyncQueue,
		f GCompareDataFunc,
		user_data Gpointer)

	G_async_queue_sort_unlocked func(queue *GAsyncQueue,
		f GCompareDataFunc,
		user_data Gpointer)

	G_on_error_query func(prg_name string)

	G_on_error_stack_trace func(prg_name string)

	G_base64_encode_step func(in *Guchar,
		leng Gsize,
		break_lines Gboolean,
		out string,
		state *Gint,
		save *Gint) Gsize

	G_base64_encode_close func(break_lines Gboolean,
		out string,
		state *Gint,
		save *Gint) Gsize

	G_base64_encode func(data *Guchar,
		leng Gsize) string

	G_base64_decode_step func(in string,
		leng Gsize,
		out *Guchar,
		state *Gint,
		save *Guint) Gsize

	G_base64_decode func(text string,
		out_len *Gsize) *Guchar

	G_base64_decode_inplace func(text string,
		out_len *Gsize) *Guchar

	G_bit_lock func(
		address *Gint,
		lock_bit Gint)

	G_bit_trylock func(
		address *Gint,
		lock_bit Gint) Gboolean

	G_bit_unlock func(
		address *Gint,
		lock_bit Gint)

	G_get_system_config_dirs func() **Gchar

	G_get_user_runtime_dir func() string

	G_get_language_names func() **Gchar

	G_get_locale_variants func(locale string) **Gchar

	G_atexit func(f GVoidFunc)

	Glib_check_version func(required_major Guint,
		required_minor Guint,
		required_micro Guint) string

	G_atomic_int_exchange_and_add func(atomic *Gint,
		val Gint) Gint

	G_atomic_int_add func(atomic *Gint,
		val Gint)

	G_atomic_int_compare_and_exchange func(atomic *Gint,
		oldval Gint,
		newval Gint) Gboolean

	G_atomic_pointer_compare_and_exchange func(atomic *Gpointer,
		oldval Gpointer,
		newval Gpointer) Gboolean

	G_atomic_int_get func(atomic *Gint) Gint

	G_atomic_int_set func(atomic *Gint,
		newval Gint)

	G_atomic_pointer_get func(atomic *Gpointer) Gpointer

	G_atomic_pointer_set func(atomic *Gpointer,
		newval Gpointer)

	G_thread_error_quark func() GQuark

	G_bookmark_file_error_quark func() GQuark

	G_bookmark_file_new func() *GBookmarkFile

	G_bookmark_file_free func(bookmark *GBookmarkFile)

	G_bookmark_file_load_from_file func(bookmark *GBookmarkFile,
		filename string,
		e **GError) Gboolean

	G_bookmark_file_load_from_data func(bookmark *GBookmarkFile,
		data string,
		length Gsize,
		e **GError) Gboolean

	G_bookmark_file_load_from_data_dirs func(bookmark *GBookmarkFile,
		file string,
		full_path **Gchar,
		e **GError) Gboolean

	G_bookmark_file_to_data func(bookmark *GBookmarkFile,
		length *Gsize,
		e **GError) string

	G_bookmark_file_to_file func(bookmark *GBookmarkFile,
		filename string,
		e **GError) Gboolean

	G_bookmark_file_set_title func(bookmark *GBookmarkFile,
		uri string,
		title string)

	G_bookmark_file_get_title func(bookmark *GBookmarkFile,
		uri string,
		e **GError) string

	G_bookmark_file_set_description func(bookmark *GBookmarkFile,
		uri string,
		description string)

	G_bookmark_file_get_description func(bookmark *GBookmarkFile,
		uri string,
		e **GError) string

	G_bookmark_file_set_mime_type func(bookmark *GBookmarkFile,
		uri string,
		mime_type string)

	G_bookmark_file_get_mime_type func(bookmark *GBookmarkFile,
		uri string,
		e **GError) string

	G_bookmark_file_set_groups func(bookmark *GBookmarkFile,
		uri string,
		groups **Gchar,
		length Gsize)

	G_bookmark_file_add_group func(bookmark *GBookmarkFile,
		uri string,
		group string)

	G_bookmark_file_has_group func(bookmark *GBookmarkFile,
		uri string,
		group string,
		e **GError) Gboolean

	G_bookmark_file_get_groups func(bookmark *GBookmarkFile,
		uri string,
		length *Gsize,
		e **GError) **Gchar

	G_bookmark_file_add_application func(bookmark *GBookmarkFile,
		uri string,
		name string,
		exec string)

	G_bookmark_file_has_application func(bookmark *GBookmarkFile,
		uri string,
		name string,
		e **GError) Gboolean

	G_bookmark_file_get_applications func(bookmark *GBookmarkFile,
		uri string,
		length *Gsize,
		e **GError) **Gchar

	G_bookmark_file_set_app_info func(bookmark *GBookmarkFile,
		uri string,
		name string,
		exec string,
		count Gint,
		stamp Time_t,
		e **GError) Gboolean

	G_bookmark_file_get_app_info func(bookmark *GBookmarkFile,
		uri string,
		name string,
		exec **Gchar,
		count *Guint,
		stamp *Time_t,
		e **GError) Gboolean

	G_bookmark_file_set_is_private func(bookmark *GBookmarkFile,
		uri string,
		is_private Gboolean)

	G_bookmark_file_get_is_private func(bookmark *GBookmarkFile,
		uri string,
		e **GError) Gboolean

	G_bookmark_file_set_icon func(bookmark *GBookmarkFile,
		uri string,
		href string,
		mime_type string)

	G_bookmark_file_get_icon func(bookmark *GBookmarkFile,
		uri string,
		href **Gchar,
		mime_type **Gchar,
		e **GError) Gboolean

	G_bookmark_file_set_added func(bookmark *GBookmarkFile,
		uri string,
		added Time_t)

	G_bookmark_file_get_added func(bookmark *GBookmarkFile,
		uri string,
		e **GError) Time_t

	G_bookmark_file_set_modified func(bookmark *GBookmarkFile,
		uri string,
		modified Time_t)

	G_bookmark_file_get_modified func(bookmark *GBookmarkFile,
		uri string,
		e **GError) Time_t

	G_bookmark_file_set_visited func(bookmark *GBookmarkFile,
		uri string,
		visited Time_t)

	G_bookmark_file_get_visited func(bookmark *GBookmarkFile,
		uri string,
		e **GError) Time_t

	G_bookmark_file_has_item func(bookmark *GBookmarkFile,
		uri string) Gboolean

	G_bookmark_file_get_size func(bookmark *GBookmarkFile) Gint

	G_bookmark_file_get_uris func(bookmark *GBookmarkFile,
		length *Gsize) **Gchar

	G_bookmark_file_remove_group func(bookmark *GBookmarkFile,
		uri string,
		group string,
		e **GError) Gboolean

	G_bookmark_file_remove_application func(bookmark *GBookmarkFile,
		uri string,
		name string,
		e **GError) Gboolean

	G_bookmark_file_remove_item func(bookmark *GBookmarkFile,
		uri string,
		e **GError) Gboolean

	G_bookmark_file_move_item func(bookmark *GBookmarkFile,
		old_uri string,
		new_uri string,
		e **GError) Gboolean

	G_slice_alloc func(block_size Gsize) Gpointer

	G_slice_alloc0 func(block_size Gsize) Gpointer

	G_slice_copy func(block_size Gsize,
		mem_block Gconstpointer) Gpointer

	G_slice_free1 func(block_size Gsize,
		mem_block Gpointer)

	G_slice_free_chain_with_offset func(block_size Gsize,
		mem_chain Gpointer,
		next_offset Gsize)

	G_slice_set_config func(ckey GSliceConfig, value Gint64)

	G_slice_get_config func(ckey GSliceConfig) Gint64

	G_slice_get_config_state func(ckey GSliceConfig, address Gint64, n_values *Guint) *Gint64

	G_free func(mem Gpointer)

	G_malloc func(n_bytes Gsize) Gpointer

	G_malloc0 func(n_bytes Gsize) Gpointer

	G_realloc func(mem Gpointer,
		n_bytes Gsize) Gpointer

	G_try_malloc func(n_bytes Gsize) Gpointer

	G_try_malloc0 func(n_bytes Gsize) Gpointer

	G_try_realloc func(mem Gpointer,
		n_bytes Gsize) Gpointer

	G_malloc_n func(n_blocks Gsize,
		n_block_bytes Gsize) Gpointer

	G_malloc0_n func(n_blocks Gsize,
		n_block_bytes Gsize) Gpointer

	G_realloc_n func(mem Gpointer,
		n_blocks Gsize,
		n_block_bytes Gsize) Gpointer

	G_try_malloc_n func(n_blocks Gsize,
		n_block_bytes Gsize) Gpointer

	G_try_malloc0_n func(n_blocks Gsize,
		n_block_bytes Gsize) Gpointer

	G_try_realloc_n func(mem Gpointer,
		n_blocks Gsize,
		n_block_bytes Gsize) Gpointer

	G_mem_is_system_malloc func() Gboolean

	G_mem_profile func()

	G_mem_chunk_new func(name string,
		atom_size Gint,
		area_size Gsize,
		typ Gint) *GMemChunk

	G_mem_chunk_destroy func(mem_chunk *GMemChunk)

	G_mem_chunk_alloc func(mem_chunk *GMemChunk) Gpointer

	G_mem_chunk_alloc0 func(mem_chunk *GMemChunk) Gpointer

	G_mem_chunk_free func(mem_chunk *GMemChunk,
		mem Gpointer)

	G_mem_chunk_clean func(mem_chunk *GMemChunk)

	G_mem_chunk_reset func(mem_chunk *GMemChunk)

	G_mem_chunk_print func(mem_chunk *GMemChunk)

	G_mem_chunk_info func()

	G_blow_chunks func()

	G_allocator_new func(name string,
		n_preallocs Guint) *GAllocator

	G_allocator_free func(allocator *GAllocator)

	G_list_alloc func() *GList

	G_list_free func(list *GList)

	G_list_free_1 func(list *GList)

	G_list_free_full func(list *GList,
		free_func GDestroyNotify)

	G_list_append func(list *GList,
		data Gpointer) *GList

	G_list_prepend func(list *GList,
		data Gpointer) *GList

	G_list_insert func(list *GList,
		data Gpointer,
		position Gint) *GList

	G_list_insert_sorted func(list *GList,
		data Gpointer,
		f GCompareFunc) *GList

	G_list_insert_sorted_with_data func(list *GList,
		data Gpointer,
		fGCompareDataFunc,
		user_data Gpointer) *GList

	G_list_insert_before func(list *GList,
		sibling *GList,
		data Gpointer) *GList

	G_list_concat func(list1 *GList,
		list2 *GList) *GList

	G_list_remove func(list *GList,
		data Gconstpointer) *GList

	G_list_remove_all func(list *GList,
		data Gconstpointer) *GList

	G_list_remove_link func(list *GList,
		llink *GList) *GList

	G_list_delete_link func(list *GList,
		link_ *GList) *GList

	G_list_reverse func(list *GList) *GList

	G_list_copy func(list *GList) *GList

	G_list_nth func(list *GList,
		n Guint) *GList

	G_list_nth_prev func(list *GList,
		n Guint) *GList

	G_list_find func(list *GList,
		data Gconstpointer) *GList

	G_list_find_custom func(list *GList,
		data Gconstpointer,
		f GCompareFunc) *GList

	G_list_position func(list *GList,
		llink *GList) Gint

	G_list_index func(list *GList,
		data Gconstpointer) Gint

	G_list_last func(list *GList) *GList

	G_list_first func(list *GList) *GList

	G_list_length func(list *GList) Guint

	G_list_foreach func(list *GList,
		f GFunc,
		user_data Gpointer)

	G_list_sort func(list *GList,
		compare_func GCompareFunc) *GList

	G_list_sort_with_data func(list *GList,
		compare_func GCompareDataFunc,
		user_data Gpointer) *GList

	G_list_nth_data func(list *GList,
		n Guint) Gpointer

	G_list_push_allocator func(allocator Gpointer)

	G_list_pop_allocator func()

	G_cache_new func(value_new_func GCacheNewFunc,
		value_destroy_func GCacheDestroyFunc,
		key_dup_func GCacheDupFunc,
		key_destroy_func GCacheDestroyFunc,
		hash_key_func GHashFunc,
		hash_value_func GHashFunc,
		key_equal_func GEqualFunc) *GCache

	G_cache_destroy func(cache *GCache)

	G_cache_insert func(cache *GCache,
		key Gpointer) Gpointer

	G_cache_remove func(cache *GCache,
		value Gconstpointer)

	G_cache_key_foreach func(cache *GCache,
		f GHFunc,
		user_data Gpointer)

	G_cache_value_foreach func(cache *GCache,
		f GHFunc,
		user_data Gpointer)

	G_checksum_type_get_length func(checksum_type GChecksumType) Gssize

	G_checksum_new func(checksum_type GChecksumType) *GChecksum

	G_checksum_reset func(checksum *GChecksum)

	G_checksum_copy func(checksum *GChecksum) *GChecksum

	G_checksum_free func(checksum *GChecksum)

	G_checksum_update func(checksum *GChecksum,
		data *Guchar,
		length Gssize)

	G_checksum_get_string func(checksum *GChecksum) string

	G_checksum_get_digest func(checksum *GChecksum,
		buffer *Guint8,
		digest_len *Gsize)

	G_compute_checksum_for_data func(checksum_type GChecksumType,
		data *Guchar,
		length Gsize) string

	G_compute_checksum_for_string func(checksum_type GChecksumType,
		str string,
		length Gssize) string

	G_completion_new func(f GCompletionFunc) *GCompletion

	G_completion_add_items func(cmp *GCompletion,
		items *GList)

	G_completion_remove_items func(cmp *GCompletion,
		items *GList)

	G_completion_clear_items func(cmp *GCompletion)

	G_completion_complete func(cmp *GCompletion,
		prefix string,
		new_prefix **Gchar) *GList

	G_completion_complete_utf8 func(cmp *GCompletion,
		prefix string,
		new_prefix **Gchar) *GList

	G_completion_set_compare func(cmp *GCompletion,
		strncmp_func GCompletionStrncmpFunc)

	G_completion_free func(cmp *GCompletion)

	G_convert_error_quark func() GQuark

	G_iconv_open func(to_codeset string,
		from_codeset string) GIConv

	G_iconv func(converter GIConv,
		inbuf **Gchar,
		inbytes_left *Gsize,
		outbuf **Gchar,
		outbytes_left *Gsize) Gsize

	G_iconv_close func(converter GIConv) Gint

	G_convert func(str string,
		leng Gssize,
		to_codeset string,
		from_codeset string,
		bytes_read *Gsize,
		bytes_written *Gsize,
		e **GError) string

	G_convert_with_iconv func(str string,
		leng Gssize,
		converter GIConv,
		bytes_read *Gsize,
		bytes_written *Gsize,
		e **GError) string

	G_convert_with_fallback func(str string,
		leng Gssize,
		to_codeset string,
		from_codeset string,
		fallback string,
		bytes_read *Gsize,
		bytes_written *Gsize,
		e **GError) string

	G_locale_to_utf8 func(opsysstr string,
		leng Gssize,
		bytes_read *Gsize,
		bytes_written *Gsize,
		e **GError) string

	G_locale_from_utf8 func(utf8str string,
		leng Gssize,
		bytes_read *Gsize,
		bytes_written *Gsize,
		e **GError) string

	G_filename_to_utf8 func(opsysstr string,
		leng Gssize,
		bytes_read *Gsize,
		bytes_written *Gsize,
		e **GError) string

	G_filename_from_utf8 func(utf8str string,
		leng Gssize,
		bytes_read *Gsize,
		bytes_written *Gsize,
		e **GError) string

	G_filename_from_uri func(uri string,
		hostname **Gchar,
		e **GError) string

	G_filename_to_uri func(filename string,
		hostname string,
		e **GError) string

	G_filename_display_name func(filename string) string

	G_get_filename_charsets func(charsets ***Gchar) Gboolean

	G_filename_display_basename func(filename string) string

	G_uri_list_extract_uris func(uri_list string) **Gchar

	G_datalist_init func(datalist **GData)

	G_datalist_clear func(datalist **GData)

	G_datalist_id_get_data func(datalist **GData,
		key_id GQuark) Gpointer

	G_datalist_id_set_data_full func(datalist **GData,
		key_id GQuark,
		data Gpointer,
		destroy_func GDestroyNotify)

	G_datalist_id_remove_no_notify func(datalist **GData,
		key_id GQuark) Gpointer

	G_datalist_foreach func(datalist **GData,
		f GDataForeachFunc,
		user_data Gpointer)

	G_datalist_set_flags func(datalist **GData,
		flags Guint)

	G_datalist_unset_flags func(datalist **GData,
		flags Guint)

	G_datalist_get_flags func(datalist **GData) Guint

	G_dataset_destroy func(dataset_location Gconstpointer)

	G_dataset_id_get_data func(dataset_location Gconstpointer,
		key_id GQuark) Gpointer

	G_dataset_id_set_data_full func(dataset_location Gconstpointer,
		key_id GQuark,
		data Gpointer,
		destroy_func GDestroyNotify)

	G_dataset_id_remove_no_notify func(dataset_location Gconstpointer,
		key_id GQuark) Gpointer

	G_dataset_foreach func(dataset_location Gconstpointer,
		f GDataForeachFunc,
		user_data Gpointer)

	G_date_new func() *GDate

	G_date_new_dmy func(day GDateDay,
		month GDateMonth,
		year GDateYear) *GDate

	G_date_new_julian func(julian_day Guint32) *GDate

	G_date_free func(date *GDate)

	G_date_valid func(date *GDate) Gboolean

	G_date_valid_day func(day GDateDay) Gboolean

	G_date_valid_month func(month GDateMonth) Gboolean

	G_date_valid_year func(year GDateYear) Gboolean

	G_date_valid_weekday func(weekday GDateWeekday) Gboolean

	G_date_valid_julian func(julian_date Guint32) Gboolean

	G_date_valid_dmy func(day GDateDay,
		month GDateMonth,
		year GDateYear) Gboolean

	G_date_get_weekday func(date *GDate) GDateWeekday

	G_date_get_month func(date *GDate) GDateMonth

	G_date_get_year func(date *GDate) GDateYear

	G_date_get_day func(date *GDate) GDateDay

	G_date_get_julian func(date *GDate) Guint32

	G_date_get_day_of_year func(date *GDate) Guint

	G_date_get_monday_week_of_year func(date *GDate) Guint

	G_date_get_sunday_week_of_year func(date *GDate) Guint

	G_date_get_iso8601_week_of_year func(date *GDate) Guint

	G_date_clear func(date *GDate,
		n_dates Guint)

	G_date_set_parse func(date *GDate,
		str string)

	G_date_set_time_t func(date *GDate,
		timet Time_t)

	G_date_set_time_val func(date *GDate,
		timeval *GTimeVal)

	G_date_set_time func(date *GDate,
		time_ GTime)

	G_date_set_month func(date *GDate,
		month GDateMonth)

	G_date_set_day func(date *GDate,
		day GDateDay)

	G_date_set_year func(date *GDate,
		year GDateYear)

	G_date_set_dmy func(date *GDate,
		day GDateDay,
		month GDateMonth,
		y GDateYear)

	G_date_set_julian func(date *GDate,
		julian_date Guint32)

	G_date_is_first_of_month func(date *GDate) Gboolean

	G_date_is_last_of_month func(date *GDate) Gboolean

	G_date_add_days func(date *GDate,
		n_days Guint)

	G_date_subtract_days func(date *GDate,
		n_days Guint)

	G_date_add_months func(date *GDate,
		n_months Guint)

	G_date_subtract_months func(date *GDate,
		n_months Guint)

	G_date_add_years func(date *GDate,
		n_years Guint)

	G_date_subtract_years func(date *GDate,
		n_years Guint)

	G_date_is_leap_year func(year GDateYear) Gboolean

	G_date_get_days_in_month func(month GDateMonth,
		year GDateYear) Guint8

	G_date_get_monday_weeks_in_year func(year GDateYear) Guint8

	G_date_get_sunday_weeks_in_year func(year GDateYear) Guint8

	G_date_days_between func(date1 *GDate,
		date2 *GDate) Gint

	G_date_compare func(lhs *GDate,
		rhs *GDate) Gint

	G_date_to_struct_tm func(date *GDate,
		tm *Tm)

	G_date_clamp func(date *GDate,
		min_date *GDate,
		max_date *GDate)

	G_date_order func(date1 *GDate, date2 *GDate)

	G_date_strftime func(s string,
		slen Gsize,
		format string,
		date *GDate) Gsize

	G_time_zone_new func(identifier string) *GTimeZone

	G_time_zone_new_utc func() *GTimeZone

	G_time_zone_new_local func() *GTimeZone

	G_time_zone_ref func(tz *GTimeZone) *GTimeZone

	G_time_zone_unref func(tz *GTimeZone)

	G_time_zone_find_interval func(tz *GTimeZone,
		typ GTimeType,
		time Gint64) Gint

	G_time_zone_adjust_time func(tz *GTimeZone,
		typ GTimeType,
		time *Gint64) Gint

	G_time_zone_get_abbreviation func(tz *GTimeZone,
		interval Gint) string

	G_time_zone_get_offset func(tz *GTimeZone,
		interval Gint) Gint32

	G_time_zone_is_dst func(tz *GTimeZone,
		interval Gint) Gboolean

	G_date_time_unref func(datetime *GDateTime)

	G_date_time_ref func(datetime *GDateTime) *GDateTime

	G_date_time_new_now func(tz *GTimeZone) *GDateTime

	G_date_time_new_now_local func() *GDateTime

	G_date_time_new_now_utc func() *GDateTime

	G_date_time_new_from_unix_local func(t Gint64) *GDateTime

	G_date_time_new_from_unix_utc func(t Gint64) *GDateTime

	G_date_time_new_from_timeval_local func(tv *GTimeVal) *GDateTime

	G_date_time_new_from_timeval_utc func(tv *GTimeVal) *GDateTime

	G_date_time_new func(tz *GTimeZone,
		year Gint,
		month Gint,
		day Gint,
		hour Gint,
		minute Gint,
		seconds Gdouble) *GDateTime

	G_date_time_new_local func(year Gint,
		month Gint,
		day Gint,
		hour Gint,
		minute Gint,
		seconds Gdouble) *GDateTime

	G_date_time_new_utc func(year Gint,
		month Gint,
		day Gint,
		hour Gint,
		minute Gint,
		seconds Gdouble) *GDateTime

	G_date_time_add func(datetime *GDateTime,
		timespan GTimeSpan) *GDateTime

	G_date_time_add_years func(datetime *GDateTime,
		years Gint) *GDateTime

	G_date_time_add_months func(datetime *GDateTime,
		months Gint) *GDateTime

	G_date_time_add_weeks func(datetime *GDateTime,
		weeks Gint) *GDateTime

	G_date_time_add_days func(datetime *GDateTime,
		days Gint) *GDateTime

	G_date_time_add_hours func(datetime *GDateTime,
		hours Gint) *GDateTime

	G_date_time_add_minutes func(datetime *GDateTime,
		minutes Gint) *GDateTime

	G_date_time_add_seconds func(datetime *GDateTime,
		seconds Gdouble) *GDateTime

	G_date_time_add_full func(datetime *GDateTime,
		years Gint,
		months Gint,
		days Gint,
		hours Gint,
		minutes Gint,
		seconds Gdouble) *GDateTime

	G_date_time_compare func(dt1 Gconstpointer,
		dt2 Gconstpointer) Gint

	G_date_time_difference func(end *GDateTime,
		begin *GDateTime) GTimeSpan

	G_date_time_hash func(datetime Gconstpointer) Guint

	G_date_time_equal func(dt1 Gconstpointer,
		dt2 Gconstpointer) Gboolean

	G_date_time_get_ymd func(datetime *GDateTime,
		year *Gint,
		month *Gint,
		day *Gint)

	G_date_time_get_year func(datetime *GDateTime) Gint

	G_date_time_get_month func(datetime *GDateTime) Gint

	G_date_time_get_day_of_month func(datetime *GDateTime) Gint

	G_date_time_get_week_numbering_year func(datetime *GDateTime) Gint

	G_date_time_get_week_of_year func(datetime *GDateTime) Gint

	G_date_time_get_day_of_week func(datetime *GDateTime) Gint

	G_date_time_get_day_of_year func(datetime *GDateTime) Gint

	G_date_time_get_hour func(datetime *GDateTime) Gint

	G_date_time_get_minute func(datetime *GDateTime) Gint

	G_date_time_get_second func(datetime *GDateTime) Gint

	G_date_time_get_microsecond func(datetime *GDateTime) Gint

	G_date_time_get_seconds func(datetime *GDateTime) Gdouble

	G_date_time_to_unix func(datetime *GDateTime) Gint64

	G_date_time_to_timeval func(datetime *GDateTime,
		tv *GTimeVal) Gboolean

	G_date_time_get_utc_offset func(datetime *GDateTime) GTimeSpan

	G_date_time_get_timezone_abbreviation func(datetime *GDateTime) string

	G_date_time_is_daylight_savings func(datetime *GDateTime) Gboolean

	G_date_time_to_timezone func(datetime *GDateTime,
		tz *GTimeZone) *GDateTime

	G_date_time_to_local func(datetime *GDateTime) *GDateTime

	G_date_time_to_utc func(datetime *GDateTime) *GDateTime

	G_date_time_format func(datetime *GDateTime,
		format string) string

	G_dir_open_utf8 func(path string,
		flags Guint,
		e **GError) *GDir

	G_dir_read_name_utf8 func(dir *GDir) string

	G_dir_rewind func(dir *GDir)

	G_dir_close func(dir *GDir)

	G_file_error_quark func() GQuark

	G_file_error_from_errno func(err_no Gint) GFileError

	G_file_test func(filename string,
		test GFileTest) Gboolean

	G_file_get_contents func(filename string,
		contents **Gchar,
		length *Gsize,
		e **GError) Gboolean

	G_file_set_contents func(filename string,
		contents string,
		length Gssize,
		e **GError) Gboolean

	G_file_read_link func(filename string,
		e **GError) string

	G_mkstemp func(tmpl string) Gint

	G_mkstemp_full func(tmpl string,
		flags int,
		mode int) Gint

	G_file_open_tmp func(tmpl string,
		name_used **Gchar,
		e **GError) Gint

	G_format_size_for_display func(size Goffset) *Char

	G_build_path func(separator string, first_element string,
		v ...VArg) string

	G_build_pathv func(separator string,
		args **Gchar) string

	G_build_filename func(first_element string, v ...VArg) string

	G_build_filenamev func(args **Gchar) string

	G_mkdir_with_parents func(pathname string,
		mode int) int

	G_hash_table_new func(hash_func GHashFunc,
		key_equal_func GEqualFunc) *GHashTable

	G_hash_table_new_full func(hash_func GHashFunc,
		key_equal_func GEqualFunc,
		key_destroy_func GDestroyNotify,
		value_destroy_func GDestroyNotify) *GHashTable

	G_hash_table_destroy func(hash_table *GHashTable)

	G_hash_table_insert func(hash_table *GHashTable,
		key Gpointer,
		value Gpointer)

	G_hash_table_replace func(hash_table *GHashTable,
		key Gpointer,
		value Gpointer)

	G_hash_table_remove func(hash_table *GHashTable,
		key Gconstpointer) Gboolean

	G_hash_table_remove_all func(hash_table *GHashTable)

	G_hash_table_steal func(hash_table *GHashTable,
		key Gconstpointer) Gboolean

	G_hash_table_steal_all func(hash_table *GHashTable)

	G_hash_table_lookup func(hash_table *GHashTable,
		key Gconstpointer) Gpointer

	G_hash_table_lookup_extended func(hash_table *GHashTable,
		lookup_key Gconstpointer,
		orig_key *Gpointer,
		value *Gpointer) Gboolean

	G_hash_table_foreach func(hash_table *GHashTable,
		f GHFunc,
		user_data Gpointer)

	G_hash_table_find func(hash_table *GHashTable,
		predicate GHRFunc,
		user_data Gpointer) Gpointer

	G_hash_table_foreach_remove func(hash_table *GHashTable,
		f GHRFunc,
		user_data Gpointer) Guint

	G_hash_table_foreach_steal func(hash_table *GHashTable,
		f GHRFunc,
		user_data Gpointer) Guint

	G_hash_table_size func(hash_table *GHashTable) Guint

	G_hash_table_get_keys func(hash_table *GHashTable) *GList

	G_hash_table_get_values func(hash_table *GHashTable) *GList

	G_hash_table_iter_init func(iter *GHashTableIter,
		hash_table *GHashTable)

	G_hash_table_iter_next func(iter *GHashTableIter,
		key *Gpointer,
		value *Gpointer) Gboolean

	G_hash_table_iter_get_hash_table func(iter *GHashTableIter) *GHashTable

	G_hash_table_iter_remove func(iter *GHashTableIter)

	G_hash_table_iter_steal func(iter *GHashTableIter)

	G_hash_table_ref func(hash_table *GHashTable) *GHashTable

	G_hash_table_unref func(hash_table *GHashTable)

	G_str_equal func(v1 Gconstpointer,
		v2 Gconstpointer) Gboolean

	G_str_hash func(v Gconstpointer) Guint

	G_int_equal func(v1 Gconstpointer,
		v2 Gconstpointer) Gboolean

	G_int_hash func(v Gconstpointer) Guint

	G_int64_equal func(v1 Gconstpointer,
		v2 Gconstpointer) Gboolean

	G_int64_hash func(v Gconstpointer) Guint

	G_double_equal func(v1 Gconstpointer,
		v2 Gconstpointer) Gboolean

	G_double_hash func(v Gconstpointer) Guint

	G_direct_hash func(v Gconstpointer) Guint

	G_direct_equal func(v1 Gconstpointer,
		v2 Gconstpointer) Gboolean

	G_hook_list_init func(hook_list *GHookList,
		hook_size Guint)

	G_hook_list_clear func(hook_list *GHookList)

	G_hook_alloc func(hook_list *GHookList) *GHook

	G_hook_free func(hook_list *GHookList,
		hook *GHook)

	G_hook_ref func(hook_list *GHookList,
		hook *GHook) *GHook

	G_hook_unref func(hook_list *GHookList,
		hook *GHook)

	G_hook_destroy func(hook_list *GHookList,
		hook_id Gulong) Gboolean

	G_hook_destroy_link func(hook_list *GHookList,
		hook *GHook)

	G_hook_prepend func(hook_list *GHookList,
		hook *GHook)

	G_hook_insert_before func(hook_list *GHookList,
		sibling *GHook,
		hook *GHook)

	G_hook_insert_sorted func(hook_list *GHookList,
		hook *GHook,
		f GHookCompareFunc)

	G_hook_get func(hook_list *GHookList,
		hook_id Gulong) *GHook

	G_hook_find func(hook_list *GHookList,
		need_valids Gboolean,
		f GHookFindFunc,
		data Gpointer) *GHook

	G_hook_find_data func(hook_list *GHookList,
		need_valids Gboolean,
		data Gpointer) *GHook

	G_hook_find_func func(hook_list *GHookList,
		need_valids Gboolean,
		f Gpointer) *GHook

	G_hook_find_func_data func(hook_list *GHookList,
		need_valids Gboolean,
		f Gpointer,
		data Gpointer) *GHook

	G_hook_first_valid func(hook_list *GHookList,
		may_be_in_call Gboolean) *GHook

	G_hook_next_valid func(hook_list *GHookList,
		hook *GHook,
		may_be_in_call Gboolean) *GHook

	G_hook_compare_ids func(new_hook *GHook,
		sibling *GHook) Gint

	G_hook_list_invoke func(hook_list *GHookList,
		may_recurse Gboolean)

	G_hook_list_invoke_check func(hook_list *GHookList,
		may_recurse Gboolean)

	G_hook_list_marshal func(hook_list *GHookList,
		may_recurse Gboolean,
		marshaller GHookMarshaller,
		marshal_data Gpointer)

	G_hook_list_marshal_check func(hook_list *GHookList,
		may_recurse Gboolean,
		marshaller GHookCheckMarshaller,
		marshal_data Gpointer)

	G_hostname_is_non_ascii func(hostname string) Gboolean

	G_hostname_is_ascii_encoded func(hostname string) Gboolean

	G_hostname_is_ip_address func(hostname string) Gboolean

	G_hostname_to_ascii func(hostname string) string

	G_hostname_to_unicode func(hostname string) string

	G_poll func(fds *GPollFD,
		nfds Guint,
		timeout Gint) Gint

	G_slist_alloc func() *GSList

	G_slist_free func(list *GSList)

	G_slist_free_1 func(list *GSList)

	G_slist_free_full func(list *GSList,
		free_func GDestroyNotify)

	G_slist_append func(list *GSList,
		data Gpointer) *GSList

	G_slist_prepend func(list *GSList,
		data Gpointer) *GSList

	G_slist_insert func(list *GSList,
		data Gpointer,
		position Gint) *GSList

	G_slist_insert_sorted func(list *GSList,
		data Gpointer,
		f GCompareFunc) *GSList

	G_slist_insert_sorted_with_data func(list *GSList,
		data Gpointer,
		f GCompareDataFunc,
		user_data Gpointer) *GSList

	G_slist_insert_before func(slist *GSList,
		sibling *GSList,
		data Gpointer) *GSList

	G_slist_concat func(list1 *GSList,
		list2 *GSList) *GSList

	G_slist_remove func(list *GSList,
		data Gconstpointer) *GSList

	G_slist_remove_all func(list *GSList,
		data Gconstpointer) *GSList

	G_slist_remove_link func(list *GSList,
		link_ *GSList) *GSList

	G_slist_delete_link func(list *GSList,
		link_ *GSList) *GSList

	G_slist_reverse func(list *GSList) *GSList

	G_slist_copy func(list *GSList) *GSList

	G_slist_nth func(list *GSList,
		n Guint) *GSList

	G_slist_find func(list *GSList,
		data Gconstpointer) *GSList

	G_slist_find_custom func(list *GSList,
		data Gconstpointer,
		f GCompareFunc) *GSList

	G_slist_position func(list *GSList,
		llink *GSList) Gint

	G_slist_index func(list *GSList,
		data Gconstpointer) Gint

	G_slist_last func(list *GSList) *GSList

	G_slist_length func(list *GSList) Guint

	G_slist_foreach func(list *GSList,
		f GFunc,
		user_data Gpointer)

	G_slist_sort func(list *GSList,
		compare_func GCompareFunc) *GSList

	G_slist_sort_with_data func(list *GSList,
		compare_func GCompareDataFunc,
		user_data Gpointer) *GSList

	G_slist_nth_data func(list *GSList,
		n Guint) Gpointer

	G_slist_push_allocator func(dummy Gpointer)

	G_slist_pop_allocator func()

	G_main_context_new func() *GMainContext

	G_main_context_ref func(context *GMainContext) *GMainContext

	G_main_context_unref func(context *GMainContext)

	G_main_context_default func() *GMainContext

	G_main_context_iteration func(context *GMainContext,
		may_block Gboolean) Gboolean

	G_main_context_pending func(context *GMainContext) Gboolean

	G_main_context_find_source_by_id func(context *GMainContext,
		source_id Guint) *GSource

	G_main_context_find_source_by_user_data func(context *GMainContext,
		user_data Gpointer) *GSource

	G_main_context_find_source_by_funcs_user_data func(context *GMainContext,
		funcs *GSourceFuncs,
		user_data Gpointer) *GSource

	G_main_context_wakeup func(context *GMainContext)

	G_main_context_acquire func(context *GMainContext) Gboolean

	G_main_context_release func(context *GMainContext)

	G_main_context_is_owner func(context *GMainContext) Gboolean

	G_main_context_wait func(context *GMainContext,
		cond *GCond,
		mutex *GMutex) Gboolean

	G_main_context_prepare func(context *GMainContext,
		priority *Gint) Gboolean

	G_main_context_query func(context *GMainContext,
		max_priority Gint,
		timeout_ *Gint,
		fds *GPollFD,
		n_fds Gint) Gint

	G_main_context_check func(context *GMainContext,
		max_priority Gint,
		fds *GPollFD,
		n_fds Gint) Gint

	G_main_context_dispatch func(context *GMainContext)

	G_main_context_set_poll_func func(context *GMainContext,
		f GPollFunc)

	G_main_context_get_poll_func func(context *GMainContext) GPollFunc

	G_main_context_add_poll func(context *GMainContext,
		fd *GPollFD,
		priority Gint)

	G_main_context_remove_poll func(context *GMainContext,
		fd *GPollFD)

	G_main_depth func() Gint

	G_main_current_source func() *GSource

	G_main_context_push_thread_default func(context *GMainContext)

	G_main_context_pop_thread_default func(context *GMainContext)

	G_main_context_get_thread_default func() *GMainContext

	G_main_loop_new func(context *GMainContext,
		is_running Gboolean) *GMainLoop

	G_main_loop_run func(loop *GMainLoop)

	G_main_loop_quit func(loop *GMainLoop)

	G_main_loop_ref func(loop *GMainLoop) *GMainLoop

	G_main_loop_unref func(loop *GMainLoop)

	G_main_loop_is_running func(loop *GMainLoop) Gboolean

	G_main_loop_get_context func(loop *GMainLoop) *GMainContext

	G_source_new func(source_funcs *GSourceFuncs,

		struct_size Guint) *GSource

	G_source_ref func(source *GSource) *GSource

	G_source_unref func(source *GSource)

	G_source_attach func(source *GSource,
		context *GMainContext) Guint

	G_source_destroy func(source *GSource)

	G_source_set_priority func(source *GSource,
		priority Gint)

	G_source_get_priority func(source *GSource) Gint

	G_source_set_can_recurse func(source *GSource,
		can_recurse Gboolean)

	G_source_get_can_recurse func(source *GSource) Gboolean

	G_source_get_id func(source *GSource) Guint

	G_source_get_context func(source *GSource) *GMainContext

	G_source_set_callback func(source *GSource,
		f GSourceFunc,
		data Gpointer,
		notify GDestroyNotify)

	G_source_set_funcs func(source *GSource,
		funcs *GSourceFuncs)

	G_source_is_destroyed func(source *GSource) Gboolean

	G_source_set_name func(source *GSource,
		name *Char)

	G_source_get_name func(source *GSource) *Char

	G_source_set_name_by_id func(tag Guint,
		name *Char)

	G_source_set_callback_indirect func(source *GSource,
		callback_data Gpointer,
		callback_funcs *GSourceCallbackFuncs)

	G_source_add_poll func(source *GSource,
		fd *GPollFD)

	G_source_remove_poll func(source *GSource,
		fd *GPollFD)

	G_source_add_child_source func(source *GSource,
		child_source *GSource)

	G_source_remove_child_source func(source *GSource,
		child_source *GSource)

	G_source_get_current_time func(source *GSource,
		timeval *GTimeVal)

	G_source_get_time func(source *GSource) Gint64

	G_idle_source_new func() *GSource

	G_child_watch_source_new func(pid GPid) *GSource

	G_timeout_source_new func(interval Guint) *GSource

	G_timeout_source_new_seconds func(interval Guint) *GSource

	G_get_current_time func(result *GTimeVal)

	G_get_monotonic_time func() Gint64

	G_get_real_time func() Gint64

	G_source_remove func(tag Guint) Gboolean

	G_source_remove_by_user_data func(user_data Gpointer) Gboolean

	G_source_remove_by_funcs_user_data func(funcs *GSourceFuncs,
		user_data Gpointer) Gboolean

	G_timeout_add_full func(priority Gint,
		interval Guint,
		function GSourceFunc,
		data Gpointer,
		notify GDestroyNotify) Guint

	G_timeout_add func(interval Guint,
		function GSourceFunc,
		data Gpointer) Guint

	G_timeout_add_seconds_full func(priority Gint,
		interval Guint,
		function GSourceFunc,
		data Gpointer,
		notify GDestroyNotify) Guint

	G_timeout_add_seconds func(interval Guint,
		function GSourceFunc,
		data Gpointer) Guint

	G_child_watch_add_full func(priority Gint,
		pid GPid,
		function GChildWatchFunc,
		data Gpointer,
		notify GDestroyNotify) Guint

	G_child_watch_add func(pid GPid,
		function GChildWatchFunc,
		data Gpointer) Guint

	G_idle_add func(function GSourceFunc,
		data Gpointer) Guint

	G_idle_add_full func(priority Gint,
		function GSourceFunc,
		data Gpointer,
		notify GDestroyNotify) Guint

	G_idle_remove_by_data func(data Gpointer) Gboolean

	G_main_context_invoke_full func(context *GMainContext,
		priority Gint,
		function GSourceFunc,
		data Gpointer,
		notify GDestroyNotify)

	G_main_context_invoke func(context *GMainContext,
		function GSourceFunc,
		data Gpointer)

	G_get_charset func(charset **Char) Gboolean

	G_unichar_isalnum func(c Gunichar) Gboolean

	G_unichar_isalpha func(c Gunichar) Gboolean

	G_unichar_iscntrl func(c Gunichar) Gboolean

	G_unichar_isdigit func(c Gunichar) Gboolean

	G_unichar_isgraph func(c Gunichar) Gboolean

	G_unichar_islower func(c Gunichar) Gboolean

	G_unichar_isprint func(c Gunichar) Gboolean

	G_unichar_ispunct func(c Gunichar) Gboolean

	G_unichar_isspace func(c Gunichar) Gboolean

	G_unichar_isupper func(c Gunichar) Gboolean

	G_unichar_isxdigit func(c Gunichar) Gboolean

	G_unichar_istitle func(c Gunichar) Gboolean

	G_unichar_isdefined func(c Gunichar) Gboolean

	G_unichar_iswide func(c Gunichar) Gboolean

	G_unichar_iswide_cjk func(c Gunichar) Gboolean

	G_unichar_iszerowidth func(c Gunichar) Gboolean

	G_unichar_ismark func(c Gunichar) Gboolean

	G_unichar_toupper func(c Gunichar) Gunichar

	G_unichar_tolower func(c Gunichar) Gunichar

	G_unichar_totitle func(c Gunichar) Gunichar

	G_unichar_digit_value func(c Gunichar) Gint

	G_unichar_xdigit_value func(c Gunichar) Gint

	G_unichar_type func(c Gunichar) GUnicodeType

	G_unichar_break_type func(c Gunichar) GUnicodeBreakType

	G_unichar_combining_class func(uc Gunichar) Gint

	G_unicode_canonical_ordering func(str *Gunichar,
		leng Gsize)

	G_unicode_canonical_decomposition func(ch Gunichar,
		result_len *Gsize) *Gunichar

	G_utf8_get_char func(p string) Gunichar

	G_utf8_get_char_validated func(p string,
		max_len Gssize) Gunichar

	G_utf8_offset_to_pointer func(str string,
		offset Glong) string

	G_utf8_pointer_to_offset func(str string,
		pos string) Glong

	G_utf8_prev_char func(p string) string

	G_utf8_find_next_char func(p string,
		end string) string

	G_utf8_find_prev_char func(str string,
		p string) string

	G_utf8_strlen func(p string,
		max Gssize) Glong

	G_utf8_strncpy func(dest string,
		src string,
		n Gsize) string

	G_utf8_strchr func(p string,
		leng Gssize,
		c Gunichar) string

	G_utf8_strrchr func(p string,
		leng Gssize,
		c Gunichar) string

	G_utf8_strreverse func(str string,
		leng Gssize) string

	G_utf8_to_utf16 func(str string,
		leng Glong,
		items_read *Glong,
		items_written *Glong,
		e **GError) *Gunichar2

	G_utf8_to_ucs4 func(str string,
		leng Glong,
		items_read *Glong,
		items_written *Glong,
		e **GError) *Gunichar

	G_utf8_to_ucs4_fast func(str string,
		leng Glong,
		items_written *Glong) *Gunichar

	G_utf16_to_ucs4 func(str *Gunichar2,
		leng Glong,
		items_read *Glong,
		items_written *Glong,
		e **GError) *Gunichar

	G_utf16_to_utf8 func(str *Gunichar2,
		leng Glong,
		items_read *Glong,
		items_written *Glong,
		e **GError) string

	G_ucs4_to_utf16 func(str *Gunichar,
		leng Glong,
		items_read *Glong,
		items_written *Glong,
		e **GError) *Gunichar2

	G_ucs4_to_utf8 func(str *Gunichar,
		leng Glong,
		items_read *Glong,
		items_written *Glong,
		e **GError) string

	G_unichar_to_utf8 func(c Gunichar,
		outbuf string) Gint

	G_utf8_validate func(str string,
		max_len Gssize,
		end **Gchar) Gboolean

	G_unichar_validate func(ch Gunichar) Gboolean
	G_utf8_strup       func(str string,
		leng Gssize) string

	G_utf8_strdown func(str string,
		leng Gssize) string

	G_utf8_casefold func(str string,
		leng Gssize) string

	G_utf8_normalize func(str string,
		leng Gssize,
		mode GNormalizeMode) string

	G_utf8_collate func(str1 string,
		str2 string) Gint

	G_utf8_collate_key func(str string,
		leng Gssize) string

	G_utf8_collate_key_for_filename func(str string,
		leng Gssize) string

	G_unichar_get_mirror_char func(ch Gunichar,
		mirrored_ch *Gunichar) Gboolean

	G_unichar_get_script func(ch Gunichar) GUnicodeScript

	G_string_chunk_new func(size Gsize) *GStringChunk

	G_string_chunk_free func(chunk *GStringChunk)

	G_string_chunk_clear func(chunk *GStringChunk)

	G_string_chunk_insert func(chunk *GStringChunk,
		str string) string

	G_string_chunk_insert_len func(chunk *GStringChunk,
		str string,
		leng Gssize) string

	G_string_chunk_insert_ func(chunk *GStringChunk,
		str string) string

	G_string_new func(init string) *GString

	G_string_new_len func(init string,
		leng Gssize) *GString

	G_string_sized_new func(dfl_size Gsize) *GString

	G_string_free func(str *GString,
		free_segment Gboolean) string

	G_string_equal func(v *GString,
		v2 *GString) Gboolean

	G_string_hash func(str *GString) Guint

	G_string_assign func(str *GString,
		rval string) *GString

	G_string_truncate func(str *GString,
		leng Gsize) *GString

	G_string_set_size func(str *GString,
		leng Gsize) *GString

	G_string_insert_len func(str *GString,
		pos Gssize,
		val string,
		leng Gssize) *GString

	G_string_append func(str *GString,
		val string) *GString

	G_string_append_len func(str *GString,
		val string,
		leng Gssize) *GString

	G_string_append_c func(str *GString,
		c Gchar) *GString

	G_string_append_unichar func(str *GString,
		wc Gunichar) *GString

	G_string_prepend func(str *GString,
		val string) *GString

	G_string_prepend_c func(str *GString,
		c Gchar) *GString

	G_string_prepend_unichar func(str *GString,
		wc Gunichar) *GString

	G_string_prepend_len func(str *GString,
		val string,
		leng Gssize) *GString

	G_string_insert func(str *GString,
		pos Gssize,
		val string) *GString

	G_string_insert_c func(str *GString,
		pos Gssize,
		c Gchar) *GString

	G_string_insert_unichar func(str *GString,
		pos Gssize,
		wc Gunichar) *GString

	G_string_overwrite func(str *GString,
		pos Gsize,
		val string) *GString

	G_string_overwrite_len func(str *GString,
		pos Gsize,
		val string,
		leng Gssize) *GString

	G_string_erase func(str *GString,
		pos Gssize,
		leng Gssize) *GString

	G_string_ascii_down func(str *GString) *GString

	G_string_ascii_up func(str *GString) *GString

	G_string_vprintf func(str *GString,
		format string,
		args Va_list)

	G_string_printf func(str *GString, format string, v ...VArg)

	G_string_append_vprintf func(str *GString,
		format string,
		args Va_list)

	G_string_append_printf func(str *GString,
		format string, v ...VArg)

	G_string_append_uri_escaped func(str *GString,
		unescaped *Char,
		reserved_chars_allowed *Char,
		allow_utf8 Gboolean) *GString

	G_io_channel_init func(channel *GIOChannel)

	G_io_channel_ref func(channel *GIOChannel) *GIOChannel

	G_io_channel_unref func(channel *GIOChannel)

	G_io_channel_read func(channel *GIOChannel,
		buf string,
		count Gsize,
		bytes_read *Gsize) GIOError

	G_io_channel_write func(channel *GIOChannel,
		buf string,
		count Gsize,
		bytes_written *Gsize) GIOError

	G_io_channel_seek func(channel *GIOChannel,
		offset Gint64,
		typ GSeekType) GIOError

	G_io_channel_close func(channel *GIOChannel)

	G_io_channel_shutdown func(channel *GIOChannel,
		flush Gboolean,
		err **GError) GIOStatus

	G_io_add_watch_full func(channel *GIOChannel,
		priority Gint,
		condition GIOCondition,
		f GIOFunc,
		user_data Gpointer,
		notify GDestroyNotify) Guint

	G_io_create_watch func(channel *GIOChannel,
		condition GIOCondition) *GSource

	G_io_add_watch func(channel *GIOChannel,
		condition GIOCondition,
		f GIOFunc,
		user_data Gpointer) Guint

	G_io_channel_set_buffer_size func(channel *GIOChannel,
		size Gsize)

	G_io_channel_get_buffer_size func(channel *GIOChannel) Gsize

	G_io_channel_get_buffer_condition func(channel *GIOChannel) GIOCondition

	G_io_channel_set_flags func(channel *GIOChannel,
		flags GIOFlags,
		e **GError) GIOStatus

	G_io_channel_get_flags func(channel *GIOChannel) GIOFlags

	G_io_channel_set_line_term func(channel *GIOChannel,
		line_term string,
		length Gint)

	G_io_channel_get_line_term func(channel *GIOChannel,
		length *Gint) string

	G_io_channel_set_buffered func(channel *GIOChannel,
		buffered Gboolean)

	G_io_channel_get_buffered func(channel *GIOChannel) Gboolean

	G_io_channel_set_encoding func(channel *GIOChannel,
		encoding string,
		e **GError) GIOStatus

	G_io_channel_get_encoding func(channel *GIOChannel) string

	G_io_channel_set_close_on_unref func(channel *GIOChannel,
		do_close Gboolean)

	G_io_channel_get_close_on_unref func(channel *GIOChannel) Gboolean

	G_io_channel_flush func(channel *GIOChannel,
		e **GError) GIOStatus

	G_io_channel_read_line func(channel *GIOChannel,
		str_return **Gchar,
		length *Gsize,
		terminator_pos *Gsize,
		e **GError) GIOStatus

	G_io_channel_read_line_string func(channel *GIOChannel,
		buffer *GString,
		terminator_pos *Gsize,
		e **GError) GIOStatus

	G_io_channel_read_to_end func(channel *GIOChannel,
		str_return **Gchar,
		length *Gsize,
		e **GError) GIOStatus

	G_io_channel_read_chars func(channel *GIOChannel,
		buf string,
		count Gsize,
		bytes_read *Gsize,
		e **GError) GIOStatus

	G_io_channel_read_unichar func(channel *GIOChannel,
		thechar *Gunichar,
		e **GError) GIOStatus

	G_io_channel_write_chars func(channel *GIOChannel,
		buf string,
		count Gssize,
		bytes_written *Gsize,
		e **GError) GIOStatus

	G_io_channel_write_unichar func(channel *GIOChannel,
		thechar Gunichar,
		e **GError) GIOStatus

	G_io_channel_seek_position func(channel *GIOChannel,
		offset Gint64,
		typ GSeekType,
		e **GError) GIOStatus

	G_io_channel_new_file func(filename string,
		mode string,
		e **GError) *GIOChannel

	G_io_channel_error_quark func() GQuark

	G_io_channel_error_from_errno func(en Gint) GIOChannelError

	G_io_channel_unix_new func(fd int) *GIOChannel

	G_io_channel_unix_get_fd func(channel *GIOChannel) Gint

	G_io_channel_win32_make_pollfd func(channel *GIOChannel,
		condition GIOCondition,
		fd *GPollFD)

	G_io_channel_win32_poll func(fds *GPollFD,
		n_fds Gint,
		timeout_ Gint) Gint

	G_io_channel_win32_new_messages func(hwnd Guint) *GIOChannel

	G_io_channel_win32_new_fd func(fd Gint) *GIOChannel

	G_io_channel_win32_get_fd func(channel *GIOChannel) Gint

	G_io_channel_win32_new_socket func(socket Gint) *GIOChannel

	G_key_file_new func() *GKeyFile

	G_key_file_free func(key_file *GKeyFile)

	G_key_file_set_list_separator func(key_file *GKeyFile,
		separator Gchar)

	G_key_file_load_from_file func(key_file *GKeyFile,
		file string,
		flags GKeyFileFlags,
		e **GError) Gboolean

	G_key_file_load_from_data func(key_file *GKeyFile,
		data string,
		length Gsize,
		flags GKeyFileFlags,
		e **GError) Gboolean

	G_key_file_load_from_dirs func(key_file *GKeyFile,
		file string,
		search_dirs **Gchar,
		full_path **Gchar,
		flags GKeyFileFlags,
		e **GError) Gboolean

	G_key_file_load_from_data_dirs func(key_file *GKeyFile,
		file string,
		full_path **Gchar,
		flags GKeyFileFlags,
		e **GError) Gboolean

	G_key_file_to_data func(key_file *GKeyFile,
		length *Gsize,
		e **GError) string

	G_key_file_get_start_group func(key_file *GKeyFile) string

	G_key_file_get_groups func(key_file *GKeyFile,
		length *Gsize) **Gchar

	G_key_file_get_keys func(key_file *GKeyFile,
		group_name string,
		length *Gsize,
		e **GError) **Gchar

	G_key_file_has_group func(key_file *GKeyFile,
		group_name string) Gboolean

	G_key_file_has_key func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) Gboolean

	G_key_file_get_value func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) string

	G_key_file_set_value func(key_file *GKeyFile,
		group_name string,
		key string,
		value string)

	G_key_file_get_string func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) string

	G_key_file_set_string func(key_file *GKeyFile,
		group_name string,
		key string,
		str string)

	G_key_file_get_locale_string func(key_file *GKeyFile,
		group_name string,
		key string,
		locale string,
		e **GError) string

	G_key_file_set_locale_string func(key_file *GKeyFile,
		group_name string,
		key string,
		locale string,
		str string)

	G_key_file_get_boolean func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) Gboolean

	G_key_file_set_boolean func(key_file *GKeyFile,
		group_name string,
		key string,
		value Gboolean)

	G_key_file_get_integer func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) Gint

	G_key_file_set_integer func(key_file *GKeyFile,
		group_name string,
		key string,
		value Gint)

	G_key_file_get_int64 func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) Gint64

	G_key_file_set_int64 func(key_file *GKeyFile,
		group_name string,
		key string,
		value Gint64)

	G_key_file_get_uint64 func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) Guint64

	G_key_file_set_uint64 func(key_file *GKeyFile,
		group_name string,
		key string,
		value Guint64)

	G_key_file_get_double func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) Gdouble

	G_key_file_set_double func(key_file *GKeyFile,
		group_name string,
		key string,
		value Gdouble)

	G_key_file_get_string_list func(key_file *GKeyFile,
		group_name string,
		key string,
		length *Gsize,
		e **GError) **Gchar

	G_key_file_set_string_list func(key_file *GKeyFile,
		group_name string,
		key string,
		list **Gchar,
		length Gsize)

	G_key_file_get_locale_string_list func(key_file *GKeyFile,
		group_name string,
		key string,
		locale string,
		length *Gsize,
		e **GError) **Gchar

	G_key_file_set_locale_string_list func(key_file *GKeyFile,
		group_name string,
		key string,
		locale string,
		list **Gchar,
		length Gsize)

	G_key_file_get_boolean_list func(key_file *GKeyFile,
		group_name string,
		key string,
		length *Gsize,
		e **GError) *Gboolean

	G_key_file_set_boolean_list func(key_file *GKeyFile,
		group_name string,
		key string,
		list *Gboolean,
		length Gsize)

	G_key_file_get_integer_list func(key_file *GKeyFile,
		group_name string,
		key string,
		length *Gsize,
		e **GError) *Gint

	G_key_file_set_double_list func(key_file *GKeyFile,
		group_name string,
		key string,
		list *Gdouble,
		length Gsize)

	G_key_file_get_double_list func(key_file *GKeyFile,
		group_name string,
		key string,
		length *Gsize,
		e **GError) *Gdouble

	G_key_file_set_integer_list func(key_file *GKeyFile,
		group_name string,
		key string,
		list *Gint,
		length Gsize)

	G_key_file_set_comment func(key_file *GKeyFile,
		group_name string,
		key string,
		comment string,
		e **GError) Gboolean

	G_key_file_get_comment func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) string

	G_key_file_remove_comment func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) Gboolean

	G_key_file_remove_key func(key_file *GKeyFile,
		group_name string,
		key string,
		e **GError) Gboolean

	G_key_file_remove_group func(key_file *GKeyFile,
		group_name string,
		e **GError) Gboolean

	G_mapped_file_new func(filename string,
		writable Gboolean,
		e **GError) *GMappedFile

	G_mapped_file_get_length func(file *GMappedFile) Gsize

	G_mapped_file_get_contents func(file *GMappedFile) string

	G_mapped_file_ref func(file *GMappedFile) *GMappedFile

	G_mapped_file_unref func(file *GMappedFile)

	G_mapped_file_free func(file *GMappedFile)

	G_markup_parse_context_new func(parser *GMarkupParser,
		flags GMarkupParseFlags,
		user_data Gpointer,
		user_data_dnotify GDestroyNotify) *GMarkupParseContext

	G_markup_parse_context_free func(context *GMarkupParseContext)

	G_markup_parse_context_parse func(context *GMarkupParseContext,
		text string,
		text_len Gssize,
		e **GError) Gboolean

	G_markup_parse_context_push func(context *GMarkupParseContext,
		parser *GMarkupParser,
		user_data Gpointer)

	G_markup_parse_context_pop func(context *GMarkupParseContext) Gpointer

	G_markup_parse_context_end_parse func(context *GMarkupParseContext,
		e **GError) Gboolean

	G_markup_parse_context_get_element func(context *GMarkupParseContext) string

	G_markup_parse_context_get_element_stack func(context *GMarkupParseContext) *GSList

	G_markup_parse_context_get_position func(context *GMarkupParseContext,
		line_number *Gint,
		char_number *Gint)

	G_markup_parse_context_get_user_data func(context *GMarkupParseContext) Gpointer

	G_markup_escape_text func(text string,
		length Gssize) string

	G_markup_printf_escaped func(format *Char, v ...VArg) string

	G_markup_vprintf_escaped func(format *Char,
		args Va_list) string

	G_log_set_handler func(log_domain string,
		log_levels GLogLevelFlags,
		log_func GLogFunc,
		user_data Gpointer) Guint

	G_log_remove_handler func(log_domain string,
		handler_id Guint)

	G_log_default_handler func(log_domain string,
		log_level GLogLevelFlags,
		message string,
		unused_data Gpointer)

	G_log_set_default_handler func(log_func GLogFunc,
		user_data Gpointer) GLogFunc

	G_log func(log_domain string, log_level GLogLevelFlags,
		format string, v ...VArg)

	G_logv func(log_domain string,
		log_level GLogLevelFlags,
		format string,
		args Va_list)

	G_log_set_fatal_mask func(log_domain string,
		fatal_mask GLogLevelFlags) GLogLevelFlags

	G_log_set_always_fatal func(fatal_mask GLogLevelFlags) GLogLevelFlags

	G_return_if_fail_warning func(log_domain *Char,
		pretty_function *Char,
		expression *Char)

	G_warn_message func(domain *Char,
		file *Char,
		line int,
		f *Char,
		warnexpr *Char)

	G_assert_warning func(log_domain *Char,
		file *Char,
		line int,
		pretty_function *Char,
		expression *Char)

	G_print func(format string, v ...VArg)

	G_set_print_handler func(f GPrintFunc) GPrintFunc

	G_printerr func(format string, v ...VArg)

	G_set_printerr_handler func(f GPrintFunc) GPrintFunc

	G_node_new func(data Gpointer) *GNode

	G_node_destroy func(root *GNode)

	G_node_unlink func(node *GNode)

	G_node_copy_deep func(node *GNode,
		copy_func GCopyFunc, data Gpointer) *GNode

	G_node_copy func(node *GNode) *GNode

	G_node_insert func(parent *GNode,
		position Gint, node *GNode) *GNode

	G_node_insert_before func(parent *GNode,
		sibling *GNode, node *GNode) *GNode

	G_node_insert_after func(parent *GNode,
		sibling *GNode, node *GNode) *GNode

	G_node_prepend func(parent *GNode, node *GNode) *GNode

	G_node_n_nodes func(root *GNode,
		flags GTraverseFlags) Guint

	G_node_get_root func(node *GNode) *GNode

	G_node_is_ancestor func(node *GNode,
		descendant *GNode) Gboolean

	G_node_depth func(node *GNode) Guint

	G_node_find func(root *GNode,
		order GTraverseType,
		flags GTraverseFlags,
		data Gpointer) *GNode

	G_node_traverse func(root *GNode,
		order GTraverseType,
		flags GTraverseFlags,
		max_depth Gint,
		f GNodeTraverseFunc,
		data Gpointer)

	G_node_max_height func(root *GNode) Guint

	G_node_children_foreach func(node *GNode,
		flags GTraverseFlags,
		f GNodeForeachFunc,
		data Gpointer)

	G_node_reverse_children func(node *GNode)

	G_node_n_children func(node *GNode) Guint

	G_node_nth_child func(node *GNode,
		n Guint) *GNode

	G_node_last_child func(node *GNode) *GNode

	G_node_find_child func(node *GNode,
		flags GTraverseFlags,
		data Gpointer) *GNode

	G_node_child_position func(node *GNode,
		child *GNode) Gint

	G_node_child_index func(node *GNode,
		data Gpointer) Gint

	G_node_first_sibling func(node *GNode) *GNode

	G_node_last_sibling func(node *GNode) *GNode

	G_node_push_allocator func(dummy Gpointer)

	G_node_pop_allocator func()

	G_option_context_new func(parameter_string string) *GOptionContext

	G_option_context_set_summary func(context *GOptionContext,
		summary string)

	G_option_context_get_summary func(context *GOptionContext) string

	G_option_context_set_description func(context *GOptionContext,
		description string)

	G_option_context_get_description func(context *GOptionContext) string

	G_option_context_free func(context *GOptionContext)

	G_option_context_set_help_enabled func(context *GOptionContext,
		help_enabled Gboolean)

	G_option_context_get_help_enabled func(context *GOptionContext) Gboolean

	G_option_context_set_ignore_unknown_options func(context *GOptionContext,
		ignore_unknown Gboolean)

	G_option_context_get_ignore_unknown_options func(context *GOptionContext) Gboolean

	G_option_context_add_main_entries func(context *GOptionContext,
		entries *GOptionEntry,
		translation_domain string)

	G_option_context_parse func(context *GOptionContext,
		argc *Gint,
		argv ***Gchar,
		e **GError) Gboolean

	G_option_context_set_translate_func func(context *GOptionContext,
		f GTranslateFunc,
		data Gpointer,
		destroy_notify GDestroyNotify)

	G_option_context_set_translation_domain func(context *GOptionContext,
		domain string)

	G_option_context_add_group func(context *GOptionContext,
		group *GOptionGroup)

	G_option_context_set_main_group func(context *GOptionContext,
		group *GOptionGroup)
	G_option_context_get_main_group func(context *GOptionContext) *GOptionGroup

	G_option_context_get_help func(context *GOptionContext,
		main_help Gboolean,
		group *GOptionGroup) string

	G_option_group_new func(name string,
		description string,
		help_description string,
		user_data Gpointer,
		destroy GDestroyNotify) *GOptionGroup

	G_option_group_set_parse_hooks func(group *GOptionGroup,
		pre_parse_func GOptionParseFunc,
		post_parse_func GOptionParseFunc)

	G_option_group_set_error_hook func(group *GOptionGroup,
		error_func GOptionErrorFunc)

	G_option_group_free func(group *GOptionGroup)

	G_option_group_add_entries func(group *GOptionGroup,
		entries *GOptionEntry)

	G_option_group_set_translate_func func(group *GOptionGroup,
		f GTranslateFunc,
		data Gpointer,
		destroy_notify GDestroyNotify)

	G_option_group_set_translation_domain func(group *GOptionGroup,
		domain string)

	G_pattern_spec_new func(pattern string) *GPatternSpec

	G_pattern_spec_free func(pspec *GPatternSpec)

	G_pattern_spec_equal func(pspec1 *GPatternSpec,
		pspec2 *GPatternSpec) Gboolean

	G_pattern_match func(pspec *GPatternSpec,
		string_length Guint,
		str string,
		string_reversed string) Gboolean

	G_pattern_match_string func(pspec *GPatternSpec,
		str string) Gboolean

	G_pattern_match_simple func(pattern string,
		str string) Gboolean

	G_spaced_primes_closest func(num Guint) Guint

	G_qsort_with_data func(pbase Gconstpointer,
		total_elems Gint,
		size Gsize,
		compare_func GCompareDataFunc,
		user_data Gpointer)

	G_queue_new func() *GQueue

	G_queue_free func(queue *GQueue)

	G_queue_init func(queue *GQueue)

	G_queue_clear func(queue *GQueue)

	G_queue_is_empty func(queue *GQueue) Gboolean

	G_queue_get_length func(queue *GQueue) Guint

	G_queue_reverse func(queue *GQueue)

	G_queue_copy func(queue *GQueue) *GQueue

	G_queue_foreach func(queue *GQueue,
		f GFunc,
		user_data Gpointer)

	G_queue_find func(queue *GQueue,
		data Gconstpointer) *GList

	G_queue_find_custom func(queue *GQueue,
		data Gconstpointer,
		f GCompareFunc) *GList

	G_queue_sort func(queue *GQueue,
		compare_func GCompareDataFunc,
		user_data Gpointer)

	G_queue_push_head func(queue *GQueue,
		data Gpointer)

	G_queue_push_tail func(queue *GQueue,
		data Gpointer)

	G_queue_push_nth func(queue *GQueue,
		data Gpointer,
		n Gint)

	G_queue_pop_head func(queue *GQueue) Gpointer

	G_queue_pop_tail func(queue *GQueue) Gpointer

	G_queue_pop_nth func(queue *GQueue,
		n Guint) Gpointer

	G_queue_peek_head func(queue *GQueue) Gpointer

	G_queue_peek_tail func(queue *GQueue) Gpointer

	G_queue_peek_nth func(queue *GQueue,
		n Guint) Gpointer

	G_queue_index func(queue *GQueue,
		data Gconstpointer) Gint

	G_queue_remove func(queue *GQueue,
		data Gconstpointer)

	G_queue_remove_all func(queue *GQueue,
		data Gconstpointer)

	G_queue_insert_before func(queue *GQueue,
		sibling *GList,
		data Gpointer)

	G_queue_insert_after func(queue *GQueue,
		sibling *GList,
		data Gpointer)

	G_queue_insert_sorted func(queue *GQueue,
		data Gpointer,
		f GCompareDataFunc,
		user_data Gpointer)

	G_queue_push_head_link func(queue *GQueue,
		link_ *GList)

	G_queue_push_tail_link func(queue *GQueue,
		link_ *GList)

	G_queue_push_nth_link func(queue *GQueue,
		n Gint,
		link_ *GList)

	G_queue_pop_head_link func(queue *GQueue) *GList

	G_queue_pop_tail_link func(queue *GQueue) *GList

	G_queue_pop_nth_link func(queue *GQueue,
		n Guint) *GList

	G_queue_peek_head_link func(queue *GQueue) *GList

	G_queue_peek_tail_link func(queue *GQueue) *GList

	G_queue_peek_nth_link func(queue *GQueue,
		n Guint) *GList

	G_queue_link_index func(queue *GQueue,
		link_ *GList) Gint

	G_queue_unlink func(queue *GQueue,
		link_ *GList)

	G_queue_delete_link func(queue *GQueue,
		link_ *GList)

	G_rand_new_with_seed func(seed Guint32) *GRand

	G_rand_new_with_seed_array func(seed *Guint32,
		seed_length Guint) *GRand

	G_rand_new func() *GRand

	G_rand_free func(rand_ *GRand)

	G_rand_copy func(rand_ *GRand) *GRand

	G_rand_set_seed func(rand_ *GRand,
		seed Guint32)

	G_rand_set_seed_array func(rand_ *GRand,
		seed *Guint32,
		seed_length Guint)

	G_rand_int func(rand_ *GRand) Guint32

	G_rand_int_range func(rand_ *GRand,
		begin Gint32,
		end Gint32) Gint32

	G_rand_double func(rand_ *GRand) Gdouble

	G_rand_double_range func(rand_ *GRand,
		begin Gdouble,
		end Gdouble) Gdouble

	G_random_set_seed func(seed Guint32)

	G_random_int func() Guint32

	G_random_int_range func(begin Gint32,
		end Gint32) Gint32

	G_random_double func() Gdouble

	G_random_double_range func(begin Gdouble,
		end Gdouble) Gdouble

	G_relation_new func(fields Gint) *GRelation

	G_relation_destroy func(relation *GRelation)

	G_relation_index func(relation *GRelation,
		field Gint,
		hash_func GHashFunc,
		key_equal_func GEqualFunc)

	G_relation_insert func(relation *GRelation, v ...VArg)

	G_relation_delete func(relation *GRelation,
		key Gconstpointer,
		field Gint) Gint

	G_relation_select func(relation *GRelation,
		key Gconstpointer,
		field Gint) *GTuples

	G_relation_count func(relation *GRelation,
		key Gconstpointer,
		field Gint) Gint

	G_relation_exists func(
		relation *GRelation, v ...VArg) Gboolean

	G_relation_print func(relation *GRelation)

	G_tuples_destroy func(tuples *GTuples)

	G_tuples_index func(tuples *GTuples,
		index_ Gint,
		field Gint) Gpointer

	G_regex_new func(pattern string,
		compile_options GRegexCompileFlags,
		match_options GRegexMatchFlags,
		e **GError) *GRegex

	G_regex_ref func(regex *GRegex) *GRegex

	G_regex_unref func(regex *GRegex)

	G_regex_get_pattern func(regex *GRegex) string

	G_regex_get_max_backref func(regex *GRegex) Gint

	G_regex_get_capture_count func(regex *GRegex) Gint

	G_regex_get_string_number func(regex *GRegex,
		name string) Gint

	G_regex_escape_string func(str string,
		length Gint) string

	G_regex_get_compile_flags func(regex *GRegex) GRegexCompileFlags

	G_regex_get_match_flags func(regex *GRegex) GRegexMatchFlags

	G_regex_match_simple func(pattern string,
		str string,
		compile_options GRegexCompileFlags,
		match_options GRegexMatchFlags) Gboolean

	G_regex_match func(regex *GRegex,
		str string,
		match_options GRegexMatchFlags,
		match_info **GMatchInfo) Gboolean

	G_regex_match_full func(regex *GRegex,
		str string,
		string_len Gssize,
		start_position Gint,
		match_options GRegexMatchFlags,
		match_info **GMatchInfo,
		e **GError) Gboolean

	G_regex_match_all func(regex *GRegex,
		str string,
		match_options GRegexMatchFlags,
		match_info **GMatchInfo) Gboolean

	G_regex_match_all_full func(regex *GRegex,
		str string,
		string_len Gssize,
		start_position Gint,
		match_options GRegexMatchFlags,
		match_info **GMatchInfo,
		e **GError) Gboolean

	G_regex_split_simple func(pattern string,
		str string,
		compile_options GRegexCompileFlags,
		match_options GRegexMatchFlags) **Gchar

	G_regex_split func(regex *GRegex,
		str string,
		match_options GRegexMatchFlags) **Gchar

	G_regex_split_full func(regex *GRegex,
		str string,
		string_len Gssize,
		start_position Gint,
		match_options GRegexMatchFlags,
		max_tokens Gint,
		e **GError) **Gchar

	G_regex_replace func(regex *GRegex,
		str string,
		string_len Gssize,
		start_position Gint,
		replacement string,
		match_options GRegexMatchFlags,
		e **GError) string

	G_regex_replace_literal func(regex *GRegex,
		str string,
		string_len Gssize,
		start_position Gint,
		replacement string,
		match_options GRegexMatchFlags,
		e **GError) string

	G_regex_replace_eval func(regex *GRegex,
		str string,
		string_len Gssize,
		start_position Gint,
		match_options GRegexMatchFlags,
		eval GRegexEvalCallback,
		user_data Gpointer,
		e **GError) string

	G_regex_check_replacement func(replacement string,
		has_references *Gboolean,
		e **GError) Gboolean

	G_match_info_get_regex func(match_info *GMatchInfo) *GRegex

	G_match_info_get_string func(match_info *GMatchInfo) string

	G_match_info_free func(match_info *GMatchInfo)

	G_match_info_next func(match_info *GMatchInfo,
		e **GError) Gboolean

	G_match_info_matches func(match_info *GMatchInfo) Gboolean

	G_match_info_get_match_count func(match_info *GMatchInfo) Gint

	G_match_info_is_partial_match func(match_info *GMatchInfo) Gboolean

	G_match_info_expand_references func(match_info *GMatchInfo,
		string_to_expand string,
		e **GError) string

	G_match_info_fetch func(match_info *GMatchInfo,
		match_num Gint) string

	G_match_info_fetch_pos func(match_info *GMatchInfo,
		match_num Gint,
		start_pos *Gint,
		end_pos *Gint) Gboolean

	G_match_info_fetch_named func(match_info *GMatchInfo,
		name string) string

	G_match_info_fetch_named_pos func(match_info *GMatchInfo,
		name string,
		start_pos *Gint,
		end_pos *Gint) Gboolean

	G_match_info_fetch_all func(match_info *GMatchInfo) **Gchar

	G_scanner_new func(config_templ *GScannerConfig) *GScanner

	G_scanner_destroy func(scanner *GScanner)

	G_scanner_input_file func(scanner *GScanner,
		input_fd Gint)

	G_scanner_sync_file_offset func(scanner *GScanner)

	G_scanner_input_text func(scanner *GScanner,
		text string,
		text_len Guint)

	G_scanner_get_next_token func(scanner *GScanner) GTokenType

	G_scanner_peek_next_token func(scanner *GScanner) GTokenType

	G_scanner_cur_token func(scanner *GScanner) GTokenType

	G_scanner_cur_value func(scanner *GScanner) GTokenValue

	G_scanner_cur_line func(scanner *GScanner) Guint

	G_scanner_cur_position func(scanner *GScanner) Guint

	G_scanner_eof func(scanner *GScanner) Gboolean

	G_scanner_set_scope func(scanner *GScanner,
		scope_id Guint) Guint

	G_scanner_scope_add_symbol func(scanner *GScanner,
		scope_id Guint,
		symbol string,
		value Gpointer)

	G_scanner_scope_remove_symbol func(scanner *GScanner,
		scope_id Guint,
		symbol string)

	G_scanner_scope_lookup_symbol func(scanner *GScanner,
		scope_id Guint,
		symbol string) Gpointer

	G_scanner_scope_foreach_symbol func(scanner *GScanner,
		scope_id Guint,
		f GHFunc,
		user_data Gpointer)

	G_scanner_lookup_symbol func(scanner *GScanner,
		symbol string) Gpointer

	G_scanner_unexp_token func(scanner *GScanner,
		expected_token GTokenType,
		identifier_spec string,
		symbol_spec string,
		symbol_name string,
		message string,
		is_error Gint)

	G_scanner_error func(scanner *GScanner,
		format string, v ...VArg)

	G_scanner_warn func(scanner *GScanner,
		format string, v ...VArg)

	G_sequence_new func(data_destroy GDestroyNotify) *GSequence

	G_sequence_free func(seq *GSequence)

	G_sequence_get_length func(seq *GSequence) Gint

	G_sequence_foreach func(seq *GSequence,
		f GFunc,
		user_data Gpointer)

	G_sequence_foreach_range func(begin *GSequenceIter,
		end *GSequenceIter,
		f GFunc,
		user_data Gpointer)

	G_sequence_sort func(seq *GSequence,
		cmp_func GCompareDataFunc,
		cmp_data Gpointer)

	G_sequence_sort_iter func(seq *GSequence,
		cmp_func GSequenceIterCompareFunc,
		cmp_data Gpointer)

	G_sequence_get_begin_iter func(seq *GSequence) *GSequenceIter

	G_sequence_get_end_iter func(seq *GSequence) *GSequenceIter

	G_sequence_get_iter_at_pos func(seq *GSequence,
		pos Gint) *GSequenceIter

	G_sequence_append func(seq *GSequence,
		data Gpointer) *GSequenceIter

	G_sequence_prepend func(seq *GSequence,
		data Gpointer) *GSequenceIter

	G_sequence_insert_before func(iter *GSequenceIter,
		data Gpointer) *GSequenceIter

	G_sequence_move func(src *GSequenceIter,
		dest *GSequenceIter)

	G_sequence_swap func(a *GSequenceIter,
		b *GSequenceIter)

	G_sequence_insert_sorted func(seq *GSequence,
		data Gpointer,
		cmp_func GCompareDataFunc,
		cmp_data Gpointer) *GSequenceIter

	G_sequence_insert_sorted_iter func(seq *GSequence,
		data Gpointer,
		iter_cmp GSequenceIterCompareFunc,
		cmp_data Gpointer) *GSequenceIter

	G_sequence_sort_changed func(iter *GSequenceIter,
		cmp_func GCompareDataFunc,
		cmp_data Gpointer)

	G_sequence_sort_changed_iter func(iter *GSequenceIter,
		iter_cmp GSequenceIterCompareFunc,
		cmp_data Gpointer)

	G_sequence_remove func(iter *GSequenceIter)

	G_sequence_remove_range func(begin *GSequenceIter,
		end *GSequenceIter)

	G_sequence_move_range func(dest *GSequenceIter,
		begin *GSequenceIter,
		end *GSequenceIter)

	G_sequence_search func(seq *GSequence,
		data Gpointer,
		cmp_func GCompareDataFunc,
		cmp_data Gpointer) *GSequenceIter

	G_sequence_search_iter func(seq *GSequence,
		data Gpointer,
		iter_cmp GSequenceIterCompareFunc,
		cmp_data Gpointer) *GSequenceIter

	G_sequence_lookup func(seq *GSequence,
		data Gpointer,
		cmp_func GCompareDataFunc,
		cmp_data Gpointer) *GSequenceIter

	G_sequence_lookup_iter func(seq *GSequence,
		data Gpointer,
		iter_cmp GSequenceIterCompareFunc,
		cmp_data Gpointer) *GSequenceIter

	G_sequence_get func(iter *GSequenceIter) Gpointer

	G_sequence_set func(iter *GSequenceIter,
		data Gpointer)

	G_sequence_iter_is_begin func(iter *GSequenceIter) Gboolean

	G_sequence_iter_is_end func(iter *GSequenceIter) Gboolean

	G_sequence_iter_next func(iter *GSequenceIter) *GSequenceIter

	G_sequence_iter_prev func(iter *GSequenceIter) *GSequenceIter

	G_sequence_iter_get_position func(iter *GSequenceIter) Gint

	G_sequence_iter_move func(iter *GSequenceIter,
		delta Gint) *GSequenceIter

	G_sequence_iter_get_sequence func(iter *GSequenceIter) *GSequence

	G_sequence_iter_compare func(a *GSequenceIter,
		b *GSequenceIter) Gint

	G_sequence_range_get_midpoint func(begin *GSequenceIter,
		end *GSequenceIter) *GSequenceIter

	G_shell_error_quark func() GQuark

	G_shell_quote func(unquoted_string string) string

	G_shell_unquote func(quoted_string string,
		e **GError) string

	G_shell_parse_argv func(command_line string,
		argcp *Gint,
		argvp ***Gchar,
		e **GError) Gboolean

	G_spawn_error_quark func() GQuark

	G_spawn_async func(working_directory string,
		argv **Gchar,
		envp **Gchar,
		flags GSpawnFlags,
		child_setup GSpawnChildSetupFunc,
		user_data Gpointer,
		child_pid *GPid,
		e **GError) Gboolean

	G_spawn_async_with_pipes func(working_directory string,
		argv **Gchar,
		envp **Gchar,
		flags GSpawnFlags,
		child_setup GSpawnChildSetupFunc,
		user_data Gpointer,
		child_pid *GPid,
		standard_input *Gint,
		standard_output *Gint,
		standard_error *Gint,
		e **GError) Gboolean

	G_spawn_sync func(working_directory string,
		argv **Gchar,
		envp **Gchar,
		flags GSpawnFlags,
		child_setup GSpawnChildSetupFunc,
		user_data Gpointer,
		standard_output **Gchar,
		standard_error **Gchar,
		exit_status *Gint,
		e **GError) Gboolean

	G_spawn_command_line_sync func(command_line string,
		standard_output **Gchar,
		standard_error **Gchar,
		exit_status *Gint,
		e **GError) Gboolean

	G_spawn_command_line_async func(command_line string,
		e **GError) Gboolean

	G_spawn_close_pid func(pid GPid)

	G_ascii_tolower func(c Gchar) Gchar

	G_ascii_toupper func(c Gchar) Gchar

	G_ascii_digit_value func(c Gchar) Gint

	G_ascii_xdigit_value func(c Gchar) Gint

	G_strdelimit func(str string,
		delimiters string,
		new_delimiter Gchar) string

	G_strcanon func(str string,
		valid_chars string,
		substitutor Gchar) string

	G_strerror func(errnum Gint) string

	G_strsignal func(signum Gint) string

	G_strreverse func(str string) string

	G_strlcpy func(dest string,
		src string,
		dest_size Gsize) Gsize

	G_strlcat func(dest string,
		src string,
		dest_size Gsize) Gsize

	G_strstr_len func(haystack string,
		haystack_len Gssize,
		needle string) string

	G_strrstr func(haystack string,
		needle string) string

	G_strrstr_len func(haystack string,
		haystack_len Gssize,
		needle string) string

	G_str_has_suffix func(str string,
		suffix string) Gboolean

	G_str_has_prefix func(str string,
		prefix string) Gboolean

	G_strtod func(nptr string,
		endptr **Gchar) Gdouble

	G_ascii_strtod func(nptr string,
		endptr **Gchar) Gdouble

	G_ascii_strtoull func(nptr string,
		endptr **Gchar,
		base Guint) Guint64

	G_ascii_strtoll func(nptr string,
		endptr **Gchar,
		base Guint) Gint64

	G_ascii_dtostr func(buffer string,
		buf_len Gint,
		d Gdouble) string

	G_ascii_formatd func(buffer string,
		buf_len Gint,
		format string,
		d Gdouble) string

	G_strchug func(str string) string

	G_strchomp func(str string) string

	G_ascii_strcasecmp func(s1 string,
		s2 string) Gint

	G_ascii_strncasecmp func(s1 string,
		s2 string,
		n Gsize) Gint

	G_ascii_strdown func(str string,
		leng Gssize) string

	G_ascii_strup func(str string,
		leng Gssize) string

	G_strcasecmp func(s1 string,
		s2 string) Gint

	G_strncasecmp func(s1 string,
		s2 string,
		n Guint) Gint

	G_strdown func(str string) string

	G_strup func(str string) string

	G_strdup func(str string) string

	G_strdup_printf func(format string, v ...VArg) string

	G_strdup_vprintf func(format string,
		args Va_list) string

	G_strndup func(str string,
		n Gsize) string

	G_strnfill func(length Gsize,
		fill_char Gchar) string

	G_strconcat func(string1 string, v ...VArg) string

	G_strjoin func(separator string, v ...VArg) string

	G_strcompress func(source string) string

	G_strescape func(source string,
		exceptions string) string

	G_memdup func(mem Gconstpointer,
		byte_size Guint) Gpointer

	G_strsplit func(str string,
		delimiter string,
		max_tokens Gint) **Gchar

	G_strsplit_set func(str string,
		delimiters string,
		max_tokens Gint) **Gchar

	G_strjoinv func(separator string,
		str_array **Gchar) string

	G_strfreev func(str_array **Gchar)

	G_strdupv func(str_array **Gchar) **Gchar

	G_strv_length func(str_array **Gchar) Guint

	G_stpcpy func(dest string,
		src *Char) string

	G_strip_context func(msgid string,
		msgval string) string

	G_dgettext func(domain string,
		msgid string) string

	G_dcgettext func(domain string,
		msgid string,
		category int) string

	G_dngettext func(domain string,
		msgid string,
		msgid_plural string,
		n Gulong) string

	G_dpgettext func(domain string,
		msgctxtid string,
		msgidoffset Gsize) string

	G_dpgettext2 func(domain string,
		context string,
		msgid string) string

	G_strcmp0 func(str1 *Char,
		str2 *Char) int

	G_test_minimized_result func(minimized_quantity Double,
		format *Char, v ...VArg)

	G_test_maximized_result func(maximized_quantity Double,
		format *Char, v ...VArg)

	G_test_init func(argc *int, argv ***Char, v ...VArg)

	G_test_run func() int

	G_test_add_func func(testpath *Char,
		test_func GTestFunc)

	G_test_add_data_func func(testpath *Char,
		test_data Gconstpointer,
		test_func GTestDataFunc)

	G_test_message func(format *Char, v ...VArg)

	G_test_bug_base func(uri_pattern *Char)

	G_test_bug func(bug_uri_snippet *Char)

	G_test_timer_start func()

	G_test_timer_elapsed func() Double

	G_test_timer_last func() Double

	G_test_queue_free func(gfree_pointer Gpointer)

	G_test_queue_destroy func(destroy_func GDestroyNotify,
		destroy_data Gpointer)

	G_test_trap_fork func(usec_timeout Guint64,
		test_trap_flags GTestTrapFlags) Gboolean

	G_test_trap_has_passed func() Gboolean

	G_test_trap_reached_timeout func() Gboolean

	G_test_rand_int func() Gint32

	G_test_rand_int_range func(begin Gint32,
		end Gint32) Gint32

	G_test_rand_double func() Double

	G_test_rand_double_range func(range_start Double,
		range_end Double) Double

	G_test_create_case func(test_name *Char,
		data_size Gsize,
		test_data Gconstpointer,
		data_setup GTestFixtureFunc,
		data_test GTestFixtureFunc,
		data_teardown GTestFixtureFunc) *GTestCase

	G_test_create_suite func(suite_name *Char) *GTestSuite

	G_test_get_root func() *GTestSuite

	G_test_suite_add func(suite *GTestSuite,
		test_case *GTestCase)

	G_test_suite_add_suite func(suite *GTestSuite,
		nestedsuite *GTestSuite)

	G_test_run_suite func(suite *GTestSuite) int

	G_test_trap_assertions func(domain *Char,
		file *Char,
		line int,
		f *Char,
		assertion_flags Guint64,
		pattern *Char)

	G_assertion_message func(domain *Char,
		file *Char,
		line int,
		f *Char,
		message *Char)

	G_assertion_message_expr func(domain *Char,
		file *Char,
		line int,
		f *Char,
		expr *Char)

	G_assertion_message_cmpstr func(domain *Char,
		file *Char,
		line int,
		f *Char,
		expr *Char,
		arg1 *Char,
		cmp *Char,
		arg2 *Char)

	G_assertion_message_cmpnum func(domain *Char,
		file *Char,
		line int,
		f *Char,
		expr *Char,
		arg1 Long_double,
		cmp *Char,
		arg2 Long_double,
		numtype Char)

	G_assertion_message_error func(domain *Char,
		file *Char,
		line int,
		f *Char,
		expr *Char,
		e *GError,
		error_domain GQuark,
		error_code int)

	G_test_add_vtable func(testpath *Char,
		data_size Gsize,
		test_data Gconstpointer,
		data_setup GTestFixtureFunc,
		data_test GTestFixtureFunc,
		data_teardown GTestFixtureFunc)

	G_test_log_type_name func(log_type GTestLogType) *Char

	G_test_log_buffer_new func() *GTestLogBuffer

	G_test_log_buffer_free func(tbuffer *GTestLogBuffer)

	G_test_log_buffer_push func(tbuffer *GTestLogBuffer,
		n_bytes Guint,
		bytes *Guint8)

	G_test_log_buffer_pop func(tbuffer *GTestLogBuffer) *GTestLogMsg

	G_test_log_msg_free func(tmsg *GTestLogMsg)

	G_test_log_set_fatal_handler func(log_func GTestLogFatalFunc,
		user_data Gpointer)

	G_thread_pool_new func(f GFunc,
		user_data Gpointer,
		max_threads Gint,
		exclusive Gboolean,
		e **GError) *GThreadPool

	G_thread_pool_push func(pool *GThreadPool,
		data Gpointer,
		e **GError)

	G_thread_pool_set_max_threads func(pool *GThreadPool,
		max_threads Gint,
		e **GError)

	G_thread_pool_get_max_threads func(pool *GThreadPool) Gint

	G_thread_pool_get_num_threads func(pool *GThreadPool) Guint

	G_thread_pool_unprocessed func(pool *GThreadPool) Guint

	G_thread_pool_free func(pool *GThreadPool,
		immediate Gboolean,
		wait_ Gboolean)

	G_thread_pool_set_max_unused_threads func(max_threads Gint)

	G_thread_pool_get_max_unused_threads func() Gint

	G_thread_pool_get_num_unused_threads func() Guint

	G_thread_pool_stop_unused_threads func()

	G_thread_pool_set_sort_function func(pool *GThreadPool,
		f GCompareDataFunc,
		user_data Gpointer)

	G_thread_pool_set_max_idle_time func(interval Guint)

	G_thread_pool_get_max_idle_time func() Guint

	G_timer_new func() *GTimer

	G_timer_destroy func(timer *GTimer)

	G_timer_start func(timer *GTimer)

	G_timer_stop func(timer *GTimer)

	G_timer_reset func(timer *GTimer)

	G_timer_continue func(timer *GTimer)

	G_timer_elapsed func(timer *GTimer,
		microseconds *Gulong) Gdouble

	G_usleep func(microseconds Gulong)

	G_time_val_add func(time_ *GTimeVal,
		microseconds Glong)

	G_time_val_from_iso8601 func(iso_date string,
		time_ *GTimeVal) Gboolean

	G_time_val_to_iso8601 func(time_ *GTimeVal) string

	G_tree_new func(key_compare_func GCompareFunc) *GTree

	G_tree_new_with_data func(key_compare_func GCompareDataFunc,
		key_compare_data Gpointer) *GTree

	G_tree_new_full func(key_compare_func GCompareDataFunc,
		key_compare_data Gpointer,
		key_destroy_func GDestroyNotify,
		value_destroy_func GDestroyNotify) *GTree

	G_tree_ref func(tree *GTree) *GTree

	G_tree_unref func(tree *GTree)

	G_tree_destroy func(tree *GTree)

	G_tree_insert func(tree *GTree,
		key Gpointer,
		value Gpointer)

	G_tree_replace func(tree *GTree,
		key Gpointer,
		value Gpointer)

	G_tree_remove func(tree *GTree,
		key Gconstpointer) Gboolean

	G_tree_steal func(tree *GTree,
		key Gconstpointer) Gboolean

	G_tree_lookup func(tree *GTree,
		key Gconstpointer) Gpointer

	G_tree_lookup_extended func(tree *GTree,
		lookup_key Gconstpointer,
		orig_key *Gpointer,
		value *Gpointer) Gboolean

	G_tree_foreach func(tree *GTree,
		f GTraverseFunc,
		user_data Gpointer)

	G_tree_traverse func(tree *GTree,
		traverse_func GTraverseFunc,
		traverse_type GTraverseType,
		user_data Gpointer)

	G_tree_search func(tree *GTree,
		search_func GCompareFunc,
		user_data Gconstpointer) Gpointer

	G_tree_height func(tree *GTree) Gint

	G_tree_nnodes func(tree *GTree) Gint

	G_uri_unescape_string func(escaped_string *Char,
		illegal_characters *Char) *Char

	G_uri_unescape_segment func(escaped_string *Char,
		escaped_string_end *Char,
		illegal_characters *Char) *Char

	G_uri_parse_scheme func(uri *Char) *Char

	G_uri_escape_string func(unescaped *Char,
		reserved_chars_allowed *Char,
		allow_utf8 Gboolean) *Char

	G_variant_type_string_is_valid func(
		type_string string) Gboolean

	G_variant_type_string_scan func(str string,
		limit string,
		endptr **Gchar) Gboolean

	G_variant_type_free func(typ *GVariantType)

	G_variant_type_copy func(typ *GVariantType) *GVariantType

	G_variant_type_new func(type_string string) *GVariantType

	G_variant_type_get_string_length func(typ *GVariantType) Gsize

	G_variant_type_peek_string func(typ *GVariantType) string

	G_variant_type_dup_string func(typ *GVariantType) string

	G_variant_type_is_definite func(typ *GVariantType) Gboolean

	G_variant_type_is_container func(typ *GVariantType) Gboolean

	G_variant_type_is_basic func(typ *GVariantType) Gboolean

	G_variant_type_is_maybe func(typ *GVariantType) Gboolean

	G_variant_type_is_array func(typ *GVariantType) Gboolean

	G_variant_type_is_tuple func(typ *GVariantType) Gboolean

	G_variant_type_is_dict_entry func(typ *GVariantType) Gboolean

	G_variant_type_is_variant func(typ *GVariantType) Gboolean

	G_variant_type_hash func(typ Gconstpointer) Guint

	G_variant_type_equal func(type1 Gconstpointer,
		type2 Gconstpointer) Gboolean

	G_variant_type_is_subtype_of func(typ *GVariantType,
		supertyp *GVariantType) Gboolean

	G_variant_type_element func(typ *GVariantType) *GVariantType

	G_variant_type_first func(typ *GVariantType) *GVariantType

	G_variant_type_next func(typ *GVariantType) *GVariantType

	G_variant_type_n_items func(typ *GVariantType) Gsize

	G_variant_type_key func(typ *GVariantType) *GVariantType

	G_variant_type_value func(typ *GVariantType) *GVariantType

	G_variant_type_new_array func(element *GVariantType) *GVariantType

	G_variant_type_new_maybe func(element *GVariantType) *GVariantType

	G_variant_type_new_tuple func(items **GVariantType,
		length Gint) *GVariantType

	G_variant_type_new_dict_entry func(key *GVariantType,
		value *GVariantType) *GVariantType

	G_variant_type_checked_ func(string) *GVariantType

	G_variant_unref func(value *GVariant)

	G_variant_ref func(value *GVariant) *GVariant

	G_variant_ref_sink func(value *GVariant) *GVariant

	G_variant_is_floating func(value *GVariant) Gboolean

	G_variant_get_type func(value *GVariant) *GVariantType

	G_variant_get_type_string func(value *GVariant) string

	G_variant_is_of_type func(value *GVariant,
		typ *GVariantType) Gboolean

	G_variant_is_container func(value *GVariant) Gboolean

	G_variant_classify func(value *GVariant) GVariantClass

	G_variant_new_boolean func(value Gboolean) *GVariant

	G_variant_new_byte func(value Guchar) *GVariant

	G_variant_new_int16 func(value Gint16) *GVariant

	G_variant_new_uint16 func(value Guint16) *GVariant

	G_variant_new_int32 func(value Gint32) *GVariant

	G_variant_new_uint32 func(value Guint32) *GVariant

	G_variant_new_int64 func(value Gint64) *GVariant

	G_variant_new_uint64 func(value Guint64) *GVariant

	G_variant_new_handle func(value Gint32) *GVariant

	G_variant_new_double func(value Gdouble) *GVariant

	G_variant_new_string func(str string) *GVariant

	G_variant_new_object_path func(object_path string) *GVariant

	G_variant_is_object_path func(str string) Gboolean

	G_variant_new_signature func(signature string) *GVariant

	G_variant_is_signature func(str string) Gboolean

	G_variant_new_variant func(value *GVariant) *GVariant

	G_variant_new_strv func(strv **Gchar,
		length Gssize) *GVariant

	G_variant_new_bytestring func(str string) *GVariant

	G_variant_new_bytestring_array func(strv **Gchar,
		length Gssize) *GVariant

	G_variant_get_boolean func(value *GVariant) Gboolean

	G_variant_get_byte func(value *GVariant) Guchar

	G_variant_get_int16 func(value *GVariant) Gint16

	G_variant_get_uint16 func(value *GVariant) Guint16

	G_variant_get_int32 func(value *GVariant) Gint32

	G_variant_get_uint32 func(value *GVariant) Guint32

	G_variant_get_int64 func(value *GVariant) Gint64

	G_variant_get_uint64 func(value *GVariant) Guint64

	G_variant_get_handle func(value *GVariant) Gint32

	G_variant_get_double func(value *GVariant) Gdouble

	G_variant_get_variant func(value *GVariant) *GVariant

	G_variant_get_string func(value *GVariant,
		length *Gsize) string

	G_variant_dup_string func(value *GVariant,
		length *Gsize) string

	G_variant_get_strv func(value *GVariant,
		length *Gsize) **Gchar

	G_variant_dup_strv func(value *GVariant,
		length *Gsize) **Gchar

	G_variant_get_bytestring func(value *GVariant) string

	G_variant_dup_bytestring func(value *GVariant,
		length *Gsize) string

	G_variant_get_bytestring_array func(value *GVariant,
		length *Gsize) **Gchar

	G_variant_dup_bytestring_array func(value *GVariant,
		length *Gsize) **Gchar

	G_variant_new_maybe func(child_type *GVariantType,
		child *GVariant) *GVariant

	G_variant_new_array func(child_type *GVariantType,
		children **GVariant,
		n_children Gsize) *GVariant

	G_variant_new_tuple func(children **GVariant,
		n_children Gsize) *GVariant

	G_variant_new_dict_entry func(key *GVariant,
		value *GVariant) *GVariant

	G_variant_get_maybe func(value *GVariant) *GVariant

	G_variant_n_children func(value *GVariant) Gsize

	G_variant_get_child func(value *GVariant, index_ Gsize,
		format_string string, v ...VArg)

	G_variant_get_child_value func(value *GVariant,
		index_ Gsize) *GVariant

	G_variant_lookup func(dictionary *GVariant, key string,
		format_string string, v ...VArg) Gboolean

	G_variant_lookup_value func(dictionary *GVariant,
		key string,
		expected_type *GVariantType) *GVariant

	G_variant_get_fixed_array func(value *GVariant,
		n_elements *Gsize,
		element_size Gsize) Gconstpointer

	G_variant_get_size func(value *GVariant) Gsize

	G_variant_get_data func(value *GVariant) Gconstpointer

	G_variant_store func(value *GVariant,
		data Gpointer)

	G_variant_print func(value *GVariant,
		type_annotate Gboolean) string

	G_variant_print_string func(value *GVariant,
		str *GString,
		type_annotate Gboolean) *GString

	G_variant_hash func(value Gconstpointer) Guint

	G_variant_equal func(one Gconstpointer,
		two Gconstpointer) Gboolean

	G_variant_get_normal_form func(value *GVariant) *GVariant

	G_variant_is_normal_form func(value *GVariant) Gboolean

	G_variant_byteswap func(value *GVariant) *GVariant

	G_variant_new_from_data func(typ *GVariantType,
		data Gconstpointer,
		size Gsize,
		trusted Gboolean,
		notify GDestroyNotify,
		user_data Gpointer) *GVariant

	G_variant_iter_new func(value *GVariant) *GVariantIter

	G_variant_iter_init func(iter *GVariantIter,
		value *GVariant) Gsize

	G_variant_iter_copy func(iter *GVariantIter) *GVariantIter

	G_variant_iter_n_children func(iter *GVariantIter) Gsize

	G_variant_iter_free func(iter *GVariantIter)

	G_variant_iter_next_value func(iter *GVariantIter) *GVariant

	G_variant_iter_next func(iter *GVariantIter,
		format_string string, v ...VArg) Gboolean

	G_variant_iter_loop func(iter *GVariantIter,
		format_string string, v ...VArg) Gboolean

	G_variant_parser_get_error_quark func() GQuark

	G_variant_builder_new func(typ *GVariantType) *GVariantBuilder

	G_variant_builder_unref func(builder *GVariantBuilder)

	G_variant_builder_ref func(builder *GVariantBuilder) *GVariantBuilder

	G_variant_builder_init func(builder *GVariantBuilder,
		typ *GVariantType)

	G_variant_builder_end func(builder *GVariantBuilder) *GVariant

	G_variant_builder_clear func(builder *GVariantBuilder)

	G_variant_builder_open func(builder *GVariantBuilder,
		typ *GVariantType)

	G_variant_builder_close func(builder *GVariantBuilder)

	G_variant_builder_add_value func(builder *GVariantBuilder,
		value *GVariant)

	G_variant_builder_add func(builder *GVariantBuilder,
		format_string string, v ...VArg)

	G_variant_builder_add_parsed func(builder *GVariantBuilder,
		format string, v ...VArg)

	G_variant_new func(format_string string, v ...VArg) *GVariant

	G_variant_get func(value *GVariant, format_string string,
		v ...VArg)

	G_variant_new_va func(format_string string,
		endptr **Gchar,
		app *Va_list) *GVariant

	G_variant_get_va func(value *GVariant,
		format_string string,
		endptr **Gchar,
		app *Va_list)

	G_variant_parse func(typ *GVariantType,
		text string,
		limit string,
		endptr **Gchar,
		err **GError) *GVariant

	G_variant_new_parsed func(format string, v ...VArg) *GVariant

	G_variant_new_parsed_va func(format string,
		app *Va_list) *GVariant

	G_variant_compare func(one Gconstpointer,
		two Gconstpointer) Gint

	G_win32_ftruncate func(f Gint,
		size Guint) Gint

	G_win32_getlocale func() string

	G_win32_error_message func(err Gint) string

	G_win32_get_package_installation_directory func(pkg string,
		dll_name string) string

	G_win32_get_package_installation_subdirectory func(pkg string,
		dll_name string,
		subdir string) string

	G_win32_get_package_installation_directory_of_module func(hmodule Gpointer) string

	G_win32_get_windows_version func() Guint

	G_win32_locale_filename_from_utf8 func(utf8filename string) string

	G_access func(filename string, mode int) int

	G_chmod func(filename string, mode int) int

	G_open func(filename string, flags int, mode int) int

	G_creat func(filename string, mode int) int

	G_rename func(oldfilename string, newfilename string) int

	G_mkdir func(filename string, mode int) int

	G_chdir func(path string) int

	G_stat func(filename string, buf *GStatBuf) int

	G_lstat func(filename string, buf *GStatBuf) int

	G_unlink func(filename string) int

	G_remove func(filename string) int

	G_rmdir func(filename string) int

	G_fopen func(filename string, mode string) *FILE

	G_freopen func(
		filename string, mode string, stream *FILE) *FILE

	G_utime func(filename string, utb *Utimbuf) int

	G_find_program_in_path func(program string) string

	G_printf func(format string, v ...VArg) Gint

	G_fprintf func(file *FILE, format string, v ...VArg) Gint

	G_sprintf func(str string, format string, v ...VArg) Gint

	G_vprintf func(format string, args Va_list) Gint

	G_vfprintf func(
		file *FILE, format string, args Va_list) Gint

	G_vsprintf func(
		str string, format string, args Va_list) Gint

	G_vasprintf func(
		string **Gchar, format string, args Va_list) Gint

	G_get_user_special_dir func(directory GUserDirectory) string

	G_key_file_error_quark func() GQuark

	G_markup_collect_attributes func(element_name string,
		attribute_names **Gchar, attribute_values **Gchar,
		error **GError, first_type GMarkupCollectType,
		first_attr string, v ...VArg) Gboolean

	G_markup_error_quark func() GQuark

	G_mem_set_vtable func(vtable *GMemVTable)

	G_option_error_quark func() GQuark

	G_printf_string_upper_bound func(
		format string,
		args Va_list) Gsize

	G_regex_error_quark func() GQuark

	G_static_private_free func(
		private_key *GStaticPrivate)

	G_static_private_get func(
		private_key *GStaticPrivate) Gpointer

	G_static_private_init func(
		private_key *GStaticPrivate)

	G_static_private_set func(
		private_key *GStaticPrivate,
		data Gpointer,
		notify GDestroyNotify)

	G_string_chunk_insert_const func(
		chunk *GStringChunk,
		str string) string

	G_string_down func(str *GString) *GString

	G_string_up func(str *GString) *GString

	G_trash_stack_push func(
		stack_p **GTrashStack,
		data_p Gpointer)

	G_trash_stack_pop func(
		stack_p **GTrashStack) Gpointer

	G_trash_stack_peek func(
		stack_p **GTrashStack) Gpointer

	G_trash_stack_height func(
		stack_p **GTrashStack) Guint
)

var dll = "libglib-2.0-0.dll"

var apiList = Apis{
	// Undocumented {"_g_debug_flags", &_g_debug_flags},
	// Undocumented {"_g_debug_initialized", &_g_debug_initialized},
	{"g_access", &G_access},
	{"g_allocator_free", &G_allocator_free},
	{"g_allocator_new", &G_allocator_new},
	{"g_array_append_vals", &G_array_append_vals},
	{"g_array_free", &G_array_free},
	{"g_array_get_element_size", &G_array_get_element_size},
	{"g_array_insert_vals", &G_array_insert_vals},
	{"g_array_new", &G_array_new},
	{"g_array_prepend_vals", &G_array_prepend_vals},
	{"g_array_ref", &G_array_ref},
	{"g_array_remove_index", &G_array_remove_index},
	{"g_array_remove_index_fast", &G_array_remove_index_fast},
	{"g_array_remove_range", &G_array_remove_range},
	{"g_array_set_size", &G_array_set_size},
	{"g_array_sized_new", &G_array_sized_new},
	{"g_array_sort", &G_array_sort},
	{"g_array_sort_with_data", &G_array_sort_with_data},
	{"g_array_unref", &G_array_unref},
	{"g_ascii_digit_value", &G_ascii_digit_value},
	{"g_ascii_dtostr", &G_ascii_dtostr},
	{"g_ascii_formatd", &G_ascii_formatd},
	{"g_ascii_strcasecmp", &G_ascii_strcasecmp},
	{"g_ascii_strdown", &G_ascii_strdown},
	{"g_ascii_strncasecmp", &G_ascii_strncasecmp},
	{"g_ascii_strtod", &G_ascii_strtod},
	{"g_ascii_strtoll", &G_ascii_strtoll},
	{"g_ascii_strtoull", &G_ascii_strtoull},
	{"g_ascii_strup", &G_ascii_strup},
	{"g_ascii_tolower", &G_ascii_tolower},
	{"g_ascii_toupper", &G_ascii_toupper},
	{"g_ascii_xdigit_value", &G_ascii_xdigit_value},
	{"g_assert_warning", &G_assert_warning},
	{"g_assertion_message", &G_assertion_message},
	{"g_assertion_message_cmpnum", &G_assertion_message_cmpnum},
	{"g_assertion_message_cmpstr", &G_assertion_message_cmpstr},
	{"g_assertion_message_error", &G_assertion_message_error},
	{"g_assertion_message_expr", &G_assertion_message_expr},
	{"g_async_queue_length", &G_async_queue_length},
	{"g_async_queue_length_unlocked", &G_async_queue_length_unlocked},
	{"g_async_queue_lock", &G_async_queue_lock},
	{"g_async_queue_new", &G_async_queue_new},
	{"g_async_queue_new_full", &G_async_queue_new_full},
	{"g_async_queue_pop", &G_async_queue_pop},
	{"g_async_queue_pop_unlocked", &G_async_queue_pop_unlocked},
	{"g_async_queue_push", &G_async_queue_push},
	{"g_async_queue_push_sorted", &G_async_queue_push_sorted},
	{"g_async_queue_push_sorted_unlocked", &G_async_queue_push_sorted_unlocked},
	{"g_async_queue_push_unlocked", &G_async_queue_push_unlocked},
	{"g_async_queue_ref", &G_async_queue_ref},
	{"g_async_queue_ref_unlocked", &G_async_queue_ref_unlocked},
	{"g_async_queue_sort", &G_async_queue_sort},
	{"g_async_queue_sort_unlocked", &G_async_queue_sort_unlocked},
	{"g_async_queue_timed_pop", &G_async_queue_timed_pop},
	{"g_async_queue_timed_pop_unlocked", &G_async_queue_timed_pop_unlocked},
	{"g_async_queue_try_pop", &G_async_queue_try_pop},
	{"g_async_queue_try_pop_unlocked", &G_async_queue_try_pop_unlocked},
	{"g_async_queue_unlock", &G_async_queue_unlock},
	{"g_async_queue_unref", &G_async_queue_unref},
	{"g_async_queue_unref_and_unlock", &G_async_queue_unref_and_unlock},
	{"g_atexit", &G_atexit},
	{"g_atomic_int_add", &G_atomic_int_add},
	{"g_atomic_int_compare_and_exchange", &G_atomic_int_compare_and_exchange},
	{"g_atomic_int_exchange_and_add", &G_atomic_int_exchange_and_add},
	{"g_atomic_int_get", &G_atomic_int_get},
	{"g_atomic_int_set", &G_atomic_int_set},
	{"g_atomic_pointer_compare_and_exchange", &G_atomic_pointer_compare_and_exchange},
	{"g_atomic_pointer_get", &G_atomic_pointer_get},
	{"g_atomic_pointer_set", &G_atomic_pointer_set},
	{"g_base64_decode", &G_base64_decode},
	{"g_base64_decode_inplace", &G_base64_decode_inplace},
	{"g_base64_decode_step", &G_base64_decode_step},
	{"g_base64_encode", &G_base64_encode},
	{"g_base64_encode_close", &G_base64_encode_close},
	{"g_base64_encode_step", &G_base64_encode_step},
	{"g_basename", &G_basename},
	{"g_bit_lock", &G_bit_lock},
	// Inline {"g_bit_nth_lsf", &G_bit_nth_lsf},
	// Inline {"g_bit_nth_msf", &G_bit_nth_msf},
	// Inline {"g_bit_storage", &G_bit_storage},
	{"g_bit_trylock", &G_bit_trylock},
	{"g_bit_unlock", &G_bit_unlock},
	{"g_blow_chunks", &G_blow_chunks},
	{"g_bookmark_file_add_application", &G_bookmark_file_add_application},
	{"g_bookmark_file_add_group", &G_bookmark_file_add_group},
	{"g_bookmark_file_error_quark", &G_bookmark_file_error_quark},
	{"g_bookmark_file_free", &G_bookmark_file_free},
	{"g_bookmark_file_get_added", &G_bookmark_file_get_added},
	{"g_bookmark_file_get_app_info", &G_bookmark_file_get_app_info},
	{"g_bookmark_file_get_applications", &G_bookmark_file_get_applications},
	{"g_bookmark_file_get_description", &G_bookmark_file_get_description},
	{"g_bookmark_file_get_groups", &G_bookmark_file_get_groups},
	{"g_bookmark_file_get_icon", &G_bookmark_file_get_icon},
	{"g_bookmark_file_get_is_private", &G_bookmark_file_get_is_private},
	{"g_bookmark_file_get_mime_type", &G_bookmark_file_get_mime_type},
	{"g_bookmark_file_get_modified", &G_bookmark_file_get_modified},
	{"g_bookmark_file_get_size", &G_bookmark_file_get_size},
	{"g_bookmark_file_get_title", &G_bookmark_file_get_title},
	{"g_bookmark_file_get_uris", &G_bookmark_file_get_uris},
	{"g_bookmark_file_get_visited", &G_bookmark_file_get_visited},
	{"g_bookmark_file_has_application", &G_bookmark_file_has_application},
	{"g_bookmark_file_has_group", &G_bookmark_file_has_group},
	{"g_bookmark_file_has_item", &G_bookmark_file_has_item},
	{"g_bookmark_file_load_from_data", &G_bookmark_file_load_from_data},
	{"g_bookmark_file_load_from_data_dirs", &G_bookmark_file_load_from_data_dirs},
	{"g_bookmark_file_load_from_file", &G_bookmark_file_load_from_file},
	{"g_bookmark_file_move_item", &G_bookmark_file_move_item},
	{"g_bookmark_file_new", &G_bookmark_file_new},
	{"g_bookmark_file_remove_application", &G_bookmark_file_remove_application},
	{"g_bookmark_file_remove_group", &G_bookmark_file_remove_group},
	{"g_bookmark_file_remove_item", &G_bookmark_file_remove_item},
	{"g_bookmark_file_set_added", &G_bookmark_file_set_added},
	{"g_bookmark_file_set_app_info", &G_bookmark_file_set_app_info},
	{"g_bookmark_file_set_description", &G_bookmark_file_set_description},
	{"g_bookmark_file_set_groups", &G_bookmark_file_set_groups},
	{"g_bookmark_file_set_icon", &G_bookmark_file_set_icon},
	{"g_bookmark_file_set_is_private", &G_bookmark_file_set_is_private},
	{"g_bookmark_file_set_mime_type", &G_bookmark_file_set_mime_type},
	{"g_bookmark_file_set_modified", &G_bookmark_file_set_modified},
	{"g_bookmark_file_set_title", &G_bookmark_file_set_title},
	{"g_bookmark_file_set_visited", &G_bookmark_file_set_visited},
	{"g_bookmark_file_to_data", &G_bookmark_file_to_data},
	{"g_bookmark_file_to_file", &G_bookmark_file_to_file},
	{"g_build_filename", &G_build_filename},
	{"g_build_filenamev", &G_build_filenamev},
	{"g_build_path", &G_build_path},
	{"g_build_pathv", &G_build_pathv},
	{"g_byte_array_append", &G_byte_array_append},
	{"g_byte_array_free", &G_byte_array_free},
	{"g_byte_array_new", &G_byte_array_new},
	{"g_byte_array_prepend", &G_byte_array_prepend},
	{"g_byte_array_ref", &G_byte_array_ref},
	{"g_byte_array_remove_index", &G_byte_array_remove_index},
	{"g_byte_array_remove_index_fast", &G_byte_array_remove_index_fast},
	{"g_byte_array_remove_range", &G_byte_array_remove_range},
	{"g_byte_array_set_size", &G_byte_array_set_size},
	{"g_byte_array_sized_new", &G_byte_array_sized_new},
	{"g_byte_array_sort", &G_byte_array_sort},
	{"g_byte_array_sort_with_data", &G_byte_array_sort_with_data},
	{"g_byte_array_unref", &G_byte_array_unref},
	{"g_cache_destroy", &G_cache_destroy},
	{"g_cache_insert", &G_cache_insert},
	{"g_cache_key_foreach", &G_cache_key_foreach},
	{"g_cache_new", &G_cache_new},
	{"g_cache_remove", &G_cache_remove},
	{"g_cache_value_foreach", &G_cache_value_foreach},
	{"g_chdir", &G_chdir},
	{"g_checksum_copy", &G_checksum_copy},
	{"g_checksum_free", &G_checksum_free},
	{"g_checksum_get_digest", &G_checksum_get_digest},
	{"g_checksum_get_string", &G_checksum_get_string},
	{"g_checksum_new", &G_checksum_new},
	{"g_checksum_reset", &G_checksum_reset},
	{"g_checksum_type_get_length", &G_checksum_type_get_length},
	{"g_checksum_update", &G_checksum_update},
	{"g_child_watch_add", &G_child_watch_add},
	{"g_child_watch_add_full", &G_child_watch_add_full},
	{"g_child_watch_source_new", &G_child_watch_source_new},
	{"g_chmod", &G_chmod},
	{"g_clear_error", &G_clear_error},
	{"g_completion_add_items", &G_completion_add_items},
	{"g_completion_clear_items", &G_completion_clear_items},
	{"g_completion_complete", &G_completion_complete},
	{"g_completion_complete_utf8", &G_completion_complete_utf8},
	{"g_completion_free", &G_completion_free},
	{"g_completion_new", &G_completion_new},
	{"g_completion_remove_items", &G_completion_remove_items},
	{"g_completion_set_compare", &G_completion_set_compare},
	{"g_compute_checksum_for_data", &G_compute_checksum_for_data},
	{"g_compute_checksum_for_string", &G_compute_checksum_for_string},
	{"g_convert", &G_convert},
	{"g_convert_error_quark", &G_convert_error_quark},
	{"g_convert_with_fallback", &G_convert_with_fallback},
	{"g_convert_with_iconv", &G_convert_with_iconv},
	{"g_creat", &G_creat},
	{"g_datalist_clear", &G_datalist_clear},
	{"g_datalist_foreach", &G_datalist_foreach},
	{"g_datalist_get_flags", &G_datalist_get_flags},
	{"g_datalist_id_get_data", &G_datalist_id_get_data},
	{"g_datalist_id_remove_no_notify", &G_datalist_id_remove_no_notify},
	{"g_datalist_id_set_data_full", &G_datalist_id_set_data_full},
	{"g_datalist_init", &G_datalist_init},
	{"g_datalist_set_flags", &G_datalist_set_flags},
	{"g_datalist_unset_flags", &G_datalist_unset_flags},
	{"g_dataset_destroy", &G_dataset_destroy},
	{"g_dataset_foreach", &G_dataset_foreach},
	{"g_dataset_id_get_data", &G_dataset_id_get_data},
	{"g_dataset_id_remove_no_notify", &G_dataset_id_remove_no_notify},
	{"g_dataset_id_set_data_full", &G_dataset_id_set_data_full},
	{"g_date_add_days", &G_date_add_days},
	{"g_date_add_months", &G_date_add_months},
	{"g_date_add_years", &G_date_add_years},
	{"g_date_clamp", &G_date_clamp},
	{"g_date_clear", &G_date_clear},
	{"g_date_compare", &G_date_compare},
	{"g_date_days_between", &G_date_days_between},
	{"g_date_free", &G_date_free},
	{"g_date_get_day", &G_date_get_day},
	{"g_date_get_day_of_year", &G_date_get_day_of_year},
	{"g_date_get_days_in_month", &G_date_get_days_in_month},
	{"g_date_get_iso8601_week_of_year", &G_date_get_iso8601_week_of_year},
	{"g_date_get_julian", &G_date_get_julian},
	{"g_date_get_monday_week_of_year", &G_date_get_monday_week_of_year},
	{"g_date_get_monday_weeks_in_year", &G_date_get_monday_weeks_in_year},
	{"g_date_get_month", &G_date_get_month},
	{"g_date_get_sunday_week_of_year", &G_date_get_sunday_week_of_year},
	{"g_date_get_sunday_weeks_in_year", &G_date_get_sunday_weeks_in_year},
	{"g_date_get_weekday", &G_date_get_weekday},
	{"g_date_get_year", &G_date_get_year},
	{"g_date_is_first_of_month", &G_date_is_first_of_month},
	{"g_date_is_last_of_month", &G_date_is_last_of_month},
	{"g_date_is_leap_year", &G_date_is_leap_year},
	{"g_date_new", &G_date_new},
	{"g_date_new_dmy", &G_date_new_dmy},
	{"g_date_new_julian", &G_date_new_julian},
	{"g_date_order", &G_date_order},
	{"g_date_set_day", &G_date_set_day},
	{"g_date_set_dmy", &G_date_set_dmy},
	{"g_date_set_julian", &G_date_set_julian},
	{"g_date_set_month", &G_date_set_month},
	{"g_date_set_parse", &G_date_set_parse},
	{"g_date_set_time", &G_date_set_time},
	{"g_date_set_time_t", &G_date_set_time_t},
	{"g_date_set_time_val", &G_date_set_time_val},
	{"g_date_set_year", &G_date_set_year},
	{"g_date_strftime", &G_date_strftime},
	{"g_date_subtract_days", &G_date_subtract_days},
	{"g_date_subtract_months", &G_date_subtract_months},
	{"g_date_subtract_years", &G_date_subtract_years},
	{"g_date_time_add", &G_date_time_add},
	{"g_date_time_add_days", &G_date_time_add_days},
	{"g_date_time_add_full", &G_date_time_add_full},
	{"g_date_time_add_hours", &G_date_time_add_hours},
	{"g_date_time_add_minutes", &G_date_time_add_minutes},
	{"g_date_time_add_months", &G_date_time_add_months},
	{"g_date_time_add_seconds", &G_date_time_add_seconds},
	{"g_date_time_add_weeks", &G_date_time_add_weeks},
	{"g_date_time_add_years", &G_date_time_add_years},
	{"g_date_time_compare", &G_date_time_compare},
	{"g_date_time_difference", &G_date_time_difference},
	{"g_date_time_equal", &G_date_time_equal},
	{"g_date_time_format", &G_date_time_format},
	{"g_date_time_get_day_of_month", &G_date_time_get_day_of_month},
	{"g_date_time_get_day_of_week", &G_date_time_get_day_of_week},
	{"g_date_time_get_day_of_year", &G_date_time_get_day_of_year},
	{"g_date_time_get_hour", &G_date_time_get_hour},
	{"g_date_time_get_microsecond", &G_date_time_get_microsecond},
	{"g_date_time_get_minute", &G_date_time_get_minute},
	{"g_date_time_get_month", &G_date_time_get_month},
	{"g_date_time_get_second", &G_date_time_get_second},
	{"g_date_time_get_seconds", &G_date_time_get_seconds},
	{"g_date_time_get_timezone_abbreviation", &G_date_time_get_timezone_abbreviation},
	{"g_date_time_get_utc_offset", &G_date_time_get_utc_offset},
	{"g_date_time_get_week_numbering_year", &G_date_time_get_week_numbering_year},
	{"g_date_time_get_week_of_year", &G_date_time_get_week_of_year},
	{"g_date_time_get_year", &G_date_time_get_year},
	{"g_date_time_get_ymd", &G_date_time_get_ymd},
	{"g_date_time_hash", &G_date_time_hash},
	{"g_date_time_is_daylight_savings", &G_date_time_is_daylight_savings},
	{"g_date_time_new", &G_date_time_new},
	{"g_date_time_new_from_timeval_local", &G_date_time_new_from_timeval_local},
	{"g_date_time_new_from_timeval_utc", &G_date_time_new_from_timeval_utc},
	{"g_date_time_new_from_unix_local", &G_date_time_new_from_unix_local},
	{"g_date_time_new_from_unix_utc", &G_date_time_new_from_unix_utc},
	{"g_date_time_new_local", &G_date_time_new_local},
	{"g_date_time_new_now", &G_date_time_new_now},
	{"g_date_time_new_now_local", &G_date_time_new_now_local},
	{"g_date_time_new_now_utc", &G_date_time_new_now_utc},
	{"g_date_time_new_utc", &G_date_time_new_utc},
	{"g_date_time_ref", &G_date_time_ref},
	{"g_date_time_to_local", &G_date_time_to_local},
	{"g_date_time_to_timeval", &G_date_time_to_timeval},
	{"g_date_time_to_timezone", &G_date_time_to_timezone},
	{"g_date_time_to_unix", &G_date_time_to_unix},
	{"g_date_time_to_utc", &G_date_time_to_utc},
	{"g_date_time_unref", &G_date_time_unref},
	{"g_date_to_struct_tm", &G_date_to_struct_tm},
	{"g_date_valid", &G_date_valid},
	{"g_date_valid_day", &G_date_valid_day},
	{"g_date_valid_dmy", &G_date_valid_dmy},
	{"g_date_valid_julian", &G_date_valid_julian},
	{"g_date_valid_month", &G_date_valid_month},
	{"g_date_valid_weekday", &G_date_valid_weekday},
	{"g_date_valid_year", &G_date_valid_year},
	{"g_dcgettext", &G_dcgettext},
	{"g_dgettext", &G_dgettext},
	{"g_dir_close", &G_dir_close},
	// Windows: _utf8 {"g_dir_open", &G_dir_open},
	{"g_dir_open_utf8", &G_dir_open_utf8},
	// Windows: _utf8 {"g_dir_read_name", &G_dir_read_name},
	{"g_dir_read_name_utf8", &G_dir_read_name_utf8},
	{"g_dir_rewind", &G_dir_rewind},
	{"g_direct_equal", &G_direct_equal},
	{"g_direct_hash", &G_direct_hash},
	{"g_dngettext", &G_dngettext},
	{"g_double_equal", &G_double_equal},
	{"g_double_hash", &G_double_hash},
	{"g_dpgettext", &G_dpgettext},
	{"g_dpgettext2", &G_dpgettext2},
	{"g_error_copy", &G_error_copy},
	{"g_error_free", &G_error_free},
	{"g_error_matches", &G_error_matches},
	{"g_error_new", &G_error_new},
	{"g_error_new_literal", &G_error_new_literal},
	{"g_error_new_valist", &G_error_new_valist},
	{"g_file_error_from_errno", &G_file_error_from_errno},
	{"g_file_error_quark", &G_file_error_quark},
	// Windows: _utf8 {"g_file_get_contents", &G_file_get_contents},
	{"g_file_get_contents_utf8", &G_file_get_contents},
	// Windows: _utf8 {"g_file_open_tmp", &G_file_open_tmp},
	{"g_file_open_tmp_utf8", &G_file_open_tmp},
	{"g_file_read_link", &G_file_read_link},
	{"g_file_set_contents", &G_file_set_contents},
	// Windows: _utf8 {"g_file_test", &G_file_test},
	{"g_file_test_utf8", &G_file_test},
	{"g_filename_display_basename", &G_filename_display_basename},
	{"g_filename_display_name", &G_filename_display_name},
	// Windows: _utf8 {"g_filename_from_uri", &G_filename_from_uri},
	{"g_filename_from_uri_utf8", &G_filename_from_uri},
	// Windows: _utf8 {"g_filename_from_utf8", &G_filename_from_utf8},
	{"g_filename_from_utf8_utf8", &G_filename_from_utf8},
	// Windows: _utf8 {"g_filename_to_uri", &G_filename_to_uri},
	{"g_filename_to_uri_utf8", &G_filename_to_uri},
	// Windows: _utf8 {"g_filename_to_utf8", &G_filename_to_utf8},
	{"g_filename_to_utf8_utf8", &G_filename_to_utf8},
	// Windows: _utf8 {"g_find_program_in_path", &G_find_program_in_path},
	{"g_find_program_in_path_utf8", &G_find_program_in_path},
	{"g_fopen", &G_fopen},
	{"g_format_size_for_display", &G_format_size_for_display},
	{"g_fprintf", &G_fprintf},
	{"g_free", &G_free},
	{"g_freopen", &G_freopen},
	{"g_get_application_name", &G_get_application_name},
	{"g_get_charset", &G_get_charset},
	// Undocumented {"g_get_codeset", &G_get_codeset},
	// Windows: _utf8 {"g_get_current_dir", &G_get_current_dir},
	{"g_get_current_dir_utf8", &G_get_current_dir},
	{"g_get_current_time", &G_get_current_time},
	{"g_get_environ", &G_get_environ},
	{"g_get_filename_charsets", &G_get_filename_charsets},
	// Windows: _utf8 {"g_get_home_dir", &G_get_home_dir},
	{"g_get_home_dir_utf8", &G_get_home_dir},
	{"g_get_host_name", &G_get_host_name},
	{"g_get_language_names", &G_get_language_names},
	{"g_get_locale_variants", &G_get_locale_variants},
	{"g_get_monotonic_time", &G_get_monotonic_time},
	{"g_get_prgname", &G_get_prgname},
	// Windows: _utf8 {"g_get_real_name", &G_get_real_name},
	{"g_get_real_name_utf8", &G_get_real_name},
	{"g_get_real_time", &G_get_real_time},
	{"g_get_system_config_dirs", &G_get_system_config_dirs},
	{"g_get_system_data_dirs", &G_get_system_data_dirs},
	// Windows: _utf8 {"g_get_tmp_dir", &G_get_tmp_dir},
	{"g_get_tmp_dir_utf8", &G_get_tmp_dir},
	{"g_get_user_cache_dir", &G_get_user_cache_dir},
	{"g_get_user_config_dir", &G_get_user_config_dir},
	{"g_get_user_data_dir", &G_get_user_data_dir},
	// Windows: _utf8 {"g_get_user_name", &G_get_user_name},
	{"g_get_user_name_utf8", &G_get_user_name},
	{"g_get_user_runtime_dir", &G_get_user_runtime_dir},
	{"g_get_user_special_dir", &G_get_user_special_dir},
	// Windows: _utf8 {"g_getenv", &G_getenv},
	{"g_getenv_utf8", &G_getenv},
	{"g_hash_table_destroy", &G_hash_table_destroy},
	{"g_hash_table_find", &G_hash_table_find},
	{"g_hash_table_foreach", &G_hash_table_foreach},
	{"g_hash_table_foreach_remove", &G_hash_table_foreach_remove},
	{"g_hash_table_foreach_steal", &G_hash_table_foreach_steal},
	{"g_hash_table_get_keys", &G_hash_table_get_keys},
	{"g_hash_table_get_values", &G_hash_table_get_values},
	{"g_hash_table_insert", &G_hash_table_insert},
	{"g_hash_table_iter_get_hash_table", &G_hash_table_iter_get_hash_table},
	{"g_hash_table_iter_init", &G_hash_table_iter_init},
	{"g_hash_table_iter_next", &G_hash_table_iter_next},
	{"g_hash_table_iter_remove", &G_hash_table_iter_remove},
	{"g_hash_table_iter_steal", &G_hash_table_iter_steal},
	{"g_hash_table_lookup", &G_hash_table_lookup},
	{"g_hash_table_lookup_extended", &G_hash_table_lookup_extended},
	{"g_hash_table_new", &G_hash_table_new},
	{"g_hash_table_new_full", &G_hash_table_new_full},
	{"g_hash_table_ref", &G_hash_table_ref},
	{"g_hash_table_remove", &G_hash_table_remove},
	{"g_hash_table_remove_all", &G_hash_table_remove_all},
	{"g_hash_table_replace", &G_hash_table_replace},
	{"g_hash_table_size", &G_hash_table_size},
	{"g_hash_table_steal", &G_hash_table_steal},
	{"g_hash_table_steal_all", &G_hash_table_steal_all},
	{"g_hash_table_unref", &G_hash_table_unref},
	{"g_hook_alloc", &G_hook_alloc},
	{"g_hook_compare_ids", &G_hook_compare_ids},
	{"g_hook_destroy", &G_hook_destroy},
	{"g_hook_destroy_link", &G_hook_destroy_link},
	{"g_hook_find", &G_hook_find},
	{"g_hook_find_data", &G_hook_find_data},
	{"g_hook_find_func", &G_hook_find_func},
	{"g_hook_find_func_data", &G_hook_find_func_data},
	{"g_hook_first_valid", &G_hook_first_valid},
	{"g_hook_free", &G_hook_free},
	{"g_hook_get", &G_hook_get},
	{"g_hook_insert_before", &G_hook_insert_before},
	{"g_hook_insert_sorted", &G_hook_insert_sorted},
	{"g_hook_list_clear", &G_hook_list_clear},
	{"g_hook_list_init", &G_hook_list_init},
	{"g_hook_list_invoke", &G_hook_list_invoke},
	{"g_hook_list_invoke_check", &G_hook_list_invoke_check},
	{"g_hook_list_marshal", &G_hook_list_marshal},
	{"g_hook_list_marshal_check", &G_hook_list_marshal_check},
	{"g_hook_next_valid", &G_hook_next_valid},
	{"g_hook_prepend", &G_hook_prepend},
	{"g_hook_ref", &G_hook_ref},
	{"g_hook_unref", &G_hook_unref},
	{"g_hostname_is_ascii_encoded", &G_hostname_is_ascii_encoded},
	{"g_hostname_is_ip_address", &G_hostname_is_ip_address},
	{"g_hostname_is_non_ascii", &G_hostname_is_non_ascii},
	{"g_hostname_to_ascii", &G_hostname_to_ascii},
	{"g_hostname_to_unicode", &G_hostname_to_unicode},
	{"g_iconv", &G_iconv},
	{"g_iconv_close", &G_iconv_close},
	{"g_iconv_open", &G_iconv_open},
	{"g_idle_add", &G_idle_add},
	{"g_idle_add_full", &G_idle_add_full},
	{"g_idle_remove_by_data", &G_idle_remove_by_data},
	{"g_idle_source_new", &G_idle_source_new},
	{"g_int64_equal", &G_int64_equal},
	{"g_int64_hash", &G_int64_hash},
	{"g_int_equal", &G_int_equal},
	{"g_int_hash", &G_int_hash},
	{"g_intern_static_string", &G_intern_static_string},
	{"g_intern_string", &G_intern_string},
	{"g_io_add_watch", &G_io_add_watch},
	{"g_io_add_watch_full", &G_io_add_watch_full},
	{"g_io_channel_close", &G_io_channel_close},
	{"g_io_channel_error_from_errno", &G_io_channel_error_from_errno},
	{"g_io_channel_error_quark", &G_io_channel_error_quark},
	{"g_io_channel_flush", &G_io_channel_flush},
	{"g_io_channel_get_buffer_condition", &G_io_channel_get_buffer_condition},
	{"g_io_channel_get_buffer_size", &G_io_channel_get_buffer_size},
	{"g_io_channel_get_buffered", &G_io_channel_get_buffered},
	{"g_io_channel_get_close_on_unref", &G_io_channel_get_close_on_unref},
	{"g_io_channel_get_encoding", &G_io_channel_get_encoding},
	{"g_io_channel_get_flags", &G_io_channel_get_flags},
	{"g_io_channel_get_line_term", &G_io_channel_get_line_term},
	{"g_io_channel_init", &G_io_channel_init},
	// Windows: _utf8 {"g_io_channel_new_file", &G_io_channel_new_file},
	{"g_io_channel_new_file_utf8", &G_io_channel_new_file},
	{"g_io_channel_read", &G_io_channel_read},
	{"g_io_channel_read_chars", &G_io_channel_read_chars},
	{"g_io_channel_read_line", &G_io_channel_read_line},
	{"g_io_channel_read_line_string", &G_io_channel_read_line_string},
	{"g_io_channel_read_to_end", &G_io_channel_read_to_end},
	{"g_io_channel_read_unichar", &G_io_channel_read_unichar},
	{"g_io_channel_ref", &G_io_channel_ref},
	{"g_io_channel_seek", &G_io_channel_seek},
	{"g_io_channel_seek_position", &G_io_channel_seek_position},
	{"g_io_channel_set_buffer_size", &G_io_channel_set_buffer_size},
	{"g_io_channel_set_buffered", &G_io_channel_set_buffered},
	{"g_io_channel_set_close_on_unref", &G_io_channel_set_close_on_unref},
	{"g_io_channel_set_encoding", &G_io_channel_set_encoding},
	{"g_io_channel_set_flags", &G_io_channel_set_flags},
	{"g_io_channel_set_line_term", &G_io_channel_set_line_term},
	{"g_io_channel_shutdown", &G_io_channel_shutdown},
	{"g_io_channel_unix_get_fd", &G_io_channel_unix_get_fd},
	{"g_io_channel_unix_new", &G_io_channel_unix_new},
	{"g_io_channel_unref", &G_io_channel_unref},
	{"g_io_channel_win32_get_fd", &G_io_channel_win32_get_fd},
	{"g_io_channel_win32_make_pollfd", &G_io_channel_win32_make_pollfd},
	{"g_io_channel_win32_new_fd", &G_io_channel_win32_new_fd},
	{"g_io_channel_win32_new_messages", &G_io_channel_win32_new_messages},
	{"g_io_channel_win32_new_socket", &G_io_channel_win32_new_socket},
	// Undocumented {"g_io_channel_win32_new_stream_socket", &G_io_channel_win32_new_stream_socket},
	{"g_io_channel_win32_poll", &G_io_channel_win32_poll},
	// Undocumented {"g_io_channel_win32_set_debug", &G_io_channel_win32_set_debug},
	{"g_io_channel_write", &G_io_channel_write},
	{"g_io_channel_write_chars", &G_io_channel_write_chars},
	{"g_io_channel_write_unichar", &G_io_channel_write_unichar},
	{"g_io_create_watch", &G_io_create_watch},
	{"g_key_file_error_quark", &G_key_file_error_quark},
	{"g_key_file_free", &G_key_file_free},
	{"g_key_file_get_boolean", &G_key_file_get_boolean},
	{"g_key_file_get_boolean_list", &G_key_file_get_boolean_list},
	{"g_key_file_get_comment", &G_key_file_get_comment},
	{"g_key_file_get_double", &G_key_file_get_double},
	{"g_key_file_get_double_list", &G_key_file_get_double_list},
	{"g_key_file_get_groups", &G_key_file_get_groups},
	{"g_key_file_get_int64", &G_key_file_get_int64},
	{"g_key_file_get_integer", &G_key_file_get_integer},
	{"g_key_file_get_integer_list", &G_key_file_get_integer_list},
	{"g_key_file_get_keys", &G_key_file_get_keys},
	{"g_key_file_get_locale_string", &G_key_file_get_locale_string},
	{"g_key_file_get_locale_string_list", &G_key_file_get_locale_string_list},
	{"g_key_file_get_start_group", &G_key_file_get_start_group},
	{"g_key_file_get_string", &G_key_file_get_string},
	{"g_key_file_get_string_list", &G_key_file_get_string_list},
	{"g_key_file_get_uint64", &G_key_file_get_uint64},
	{"g_key_file_get_value", &G_key_file_get_value},
	{"g_key_file_has_group", &G_key_file_has_group},
	{"g_key_file_has_key", &G_key_file_has_key},
	{"g_key_file_load_from_data", &G_key_file_load_from_data},
	{"g_key_file_load_from_data_dirs", &G_key_file_load_from_data_dirs},
	{"g_key_file_load_from_dirs", &G_key_file_load_from_dirs},
	{"g_key_file_load_from_file", &G_key_file_load_from_file},
	{"g_key_file_new", &G_key_file_new},
	{"g_key_file_remove_comment", &G_key_file_remove_comment},
	{"g_key_file_remove_group", &G_key_file_remove_group},
	{"g_key_file_remove_key", &G_key_file_remove_key},
	{"g_key_file_set_boolean", &G_key_file_set_boolean},
	{"g_key_file_set_boolean_list", &G_key_file_set_boolean_list},
	{"g_key_file_set_comment", &G_key_file_set_comment},
	{"g_key_file_set_double", &G_key_file_set_double},
	{"g_key_file_set_double_list", &G_key_file_set_double_list},
	{"g_key_file_set_int64", &G_key_file_set_int64},
	{"g_key_file_set_integer", &G_key_file_set_integer},
	{"g_key_file_set_integer_list", &G_key_file_set_integer_list},
	{"g_key_file_set_list_separator", &G_key_file_set_list_separator},
	{"g_key_file_set_locale_string", &G_key_file_set_locale_string},
	{"g_key_file_set_locale_string_list", &G_key_file_set_locale_string_list},
	{"g_key_file_set_string", &G_key_file_set_string},
	{"g_key_file_set_string_list", &G_key_file_set_string_list},
	{"g_key_file_set_uint64", &G_key_file_set_uint64},
	{"g_key_file_set_value", &G_key_file_set_value},
	{"g_key_file_to_data", &G_key_file_to_data},
	{"g_list_alloc", &G_list_alloc},
	{"g_list_append", &G_list_append},
	{"g_list_concat", &G_list_concat},
	{"g_list_copy", &G_list_copy},
	{"g_list_delete_link", &G_list_delete_link},
	{"g_list_find", &G_list_find},
	{"g_list_find_custom", &G_list_find_custom},
	{"g_list_first", &G_list_first},
	{"g_list_foreach", &G_list_foreach},
	{"g_list_free", &G_list_free},
	{"g_list_free_1", &G_list_free_1},
	{"g_list_free_full", &G_list_free_full},
	{"g_list_index", &G_list_index},
	{"g_list_insert", &G_list_insert},
	{"g_list_insert_before", &G_list_insert_before},
	{"g_list_insert_sorted", &G_list_insert_sorted},
	{"g_list_insert_sorted_with_data", &G_list_insert_sorted_with_data},
	{"g_list_last", &G_list_last},
	{"g_list_length", &G_list_length},
	{"g_list_nth", &G_list_nth},
	{"g_list_nth_data", &G_list_nth_data},
	{"g_list_nth_prev", &G_list_nth_prev},
	{"g_list_pop_allocator", &G_list_pop_allocator},
	{"g_list_position", &G_list_position},
	{"g_list_prepend", &G_list_prepend},
	{"g_list_push_allocator", &G_list_push_allocator},
	{"g_list_remove", &G_list_remove},
	{"g_list_remove_all", &G_list_remove_all},
	{"g_list_remove_link", &G_list_remove_link},
	{"g_list_reverse", &G_list_reverse},
	{"g_list_sort", &G_list_sort},
	{"g_list_sort_with_data", &G_list_sort_with_data},
	{"g_listenv", &G_listenv},
	{"g_locale_from_utf8", &G_locale_from_utf8},
	{"g_locale_to_utf8", &G_locale_to_utf8},
	{"g_log", &G_log},
	{"g_log_default_handler", &G_log_default_handler},
	{"g_log_remove_handler", &G_log_remove_handler},
	{"g_log_set_always_fatal", &G_log_set_always_fatal},
	{"g_log_set_default_handler", &G_log_set_default_handler},
	{"g_log_set_fatal_mask", &G_log_set_fatal_mask},
	{"g_log_set_handler", &G_log_set_handler},
	{"g_logv", &G_logv},
	{"g_lstat", &G_lstat},
	{"g_main_context_acquire", &G_main_context_acquire},
	{"g_main_context_add_poll", &G_main_context_add_poll},
	{"g_main_context_check", &G_main_context_check},
	{"g_main_context_default", &G_main_context_default},
	{"g_main_context_dispatch", &G_main_context_dispatch},
	{"g_main_context_find_source_by_funcs_user_data", &G_main_context_find_source_by_funcs_user_data},
	{"g_main_context_find_source_by_id", &G_main_context_find_source_by_id},
	{"g_main_context_find_source_by_user_data", &G_main_context_find_source_by_user_data},
	{"g_main_context_get_poll_func", &G_main_context_get_poll_func},
	{"g_main_context_get_thread_default", &G_main_context_get_thread_default},
	{"g_main_context_invoke", &G_main_context_invoke},
	{"g_main_context_invoke_full", &G_main_context_invoke_full},
	{"g_main_context_is_owner", &G_main_context_is_owner},
	{"g_main_context_iteration", &G_main_context_iteration},
	{"g_main_context_new", &G_main_context_new},
	{"g_main_context_pending", &G_main_context_pending},
	{"g_main_context_pop_thread_default", &G_main_context_pop_thread_default},
	{"g_main_context_prepare", &G_main_context_prepare},
	{"g_main_context_push_thread_default", &G_main_context_push_thread_default},
	{"g_main_context_query", &G_main_context_query},
	{"g_main_context_ref", &G_main_context_ref},
	{"g_main_context_release", &G_main_context_release},
	{"g_main_context_remove_poll", &G_main_context_remove_poll},
	{"g_main_context_set_poll_func", &G_main_context_set_poll_func},
	{"g_main_context_unref", &G_main_context_unref},
	{"g_main_context_wait", &G_main_context_wait},
	{"g_main_context_wakeup", &G_main_context_wakeup},
	{"g_main_current_source", &G_main_current_source},
	{"g_main_depth", &G_main_depth},
	{"g_main_loop_get_context", &G_main_loop_get_context},
	{"g_main_loop_is_running", &G_main_loop_is_running},
	{"g_main_loop_new", &G_main_loop_new},
	{"g_main_loop_quit", &G_main_loop_quit},
	{"g_main_loop_ref", &G_main_loop_ref},
	{"g_main_loop_run", &G_main_loop_run},
	{"g_main_loop_unref", &G_main_loop_unref},
	{"g_malloc", &G_malloc},
	{"g_malloc0", &G_malloc0},
	{"g_malloc0_n", &G_malloc0_n},
	{"g_malloc_n", &G_malloc_n},
	{"g_mapped_file_free", &G_mapped_file_free},
	{"g_mapped_file_get_contents", &G_mapped_file_get_contents},
	{"g_mapped_file_get_length", &G_mapped_file_get_length},
	{"g_mapped_file_new", &G_mapped_file_new},
	{"g_mapped_file_ref", &G_mapped_file_ref},
	{"g_mapped_file_unref", &G_mapped_file_unref},
	{"g_markup_collect_attributes", &G_markup_collect_attributes},
	{"g_markup_error_quark", &G_markup_error_quark},
	{"g_markup_escape_text", &G_markup_escape_text},
	{"g_markup_parse_context_end_parse", &G_markup_parse_context_end_parse},
	{"g_markup_parse_context_free", &G_markup_parse_context_free},
	{"g_markup_parse_context_get_element", &G_markup_parse_context_get_element},
	{"g_markup_parse_context_get_element_stack", &G_markup_parse_context_get_element_stack},
	{"g_markup_parse_context_get_position", &G_markup_parse_context_get_position},
	{"g_markup_parse_context_get_user_data", &G_markup_parse_context_get_user_data},
	{"g_markup_parse_context_new", &G_markup_parse_context_new},
	{"g_markup_parse_context_parse", &G_markup_parse_context_parse},
	{"g_markup_parse_context_pop", &G_markup_parse_context_pop},
	{"g_markup_parse_context_push", &G_markup_parse_context_push},
	{"g_markup_printf_escaped", &G_markup_printf_escaped},
	{"g_markup_vprintf_escaped", &G_markup_vprintf_escaped},
	{"g_match_info_expand_references", &G_match_info_expand_references},
	{"g_match_info_fetch", &G_match_info_fetch},
	{"g_match_info_fetch_all", &G_match_info_fetch_all},
	{"g_match_info_fetch_named", &G_match_info_fetch_named},
	{"g_match_info_fetch_named_pos", &G_match_info_fetch_named_pos},
	{"g_match_info_fetch_pos", &G_match_info_fetch_pos},
	{"g_match_info_free", &G_match_info_free},
	{"g_match_info_get_match_count", &G_match_info_get_match_count},
	{"g_match_info_get_regex", &G_match_info_get_regex},
	{"g_match_info_get_string", &G_match_info_get_string},
	{"g_match_info_is_partial_match", &G_match_info_is_partial_match},
	{"g_match_info_matches", &G_match_info_matches},
	{"g_match_info_next", &G_match_info_next},
	{"g_mem_chunk_alloc", &G_mem_chunk_alloc},
	{"g_mem_chunk_alloc0", &G_mem_chunk_alloc0},
	{"g_mem_chunk_clean", &G_mem_chunk_clean},
	{"g_mem_chunk_destroy", &G_mem_chunk_destroy},
	{"g_mem_chunk_free", &G_mem_chunk_free},
	{"g_mem_chunk_info", &G_mem_chunk_info},
	{"g_mem_chunk_new", &G_mem_chunk_new},
	{"g_mem_chunk_print", &G_mem_chunk_print},
	{"g_mem_chunk_reset", &G_mem_chunk_reset},
	{"g_mem_is_system_malloc", &G_mem_is_system_malloc},
	{"g_mem_profile", &G_mem_profile},
	{"g_mem_set_vtable", &G_mem_set_vtable},
	{"g_memdup", &G_memdup},
	{"g_mkdir", &G_mkdir},
	{"g_mkdir_with_parents", &G_mkdir_with_parents},
	// Windows: _utf8 {"g_mkstemp", &G_mkstemp},
	{"g_mkstemp_full", &G_mkstemp_full},
	{"g_mkstemp_utf8", &G_mkstemp},
	{"g_node_child_index", &G_node_child_index},
	{"g_node_child_position", &G_node_child_position},
	{"g_node_children_foreach", &G_node_children_foreach},
	{"g_node_copy", &G_node_copy},
	{"g_node_copy_deep", &G_node_copy_deep},
	{"g_node_depth", &G_node_depth},
	{"g_node_destroy", &G_node_destroy},
	{"g_node_find", &G_node_find},
	{"g_node_find_child", &G_node_find_child},
	{"g_node_first_sibling", &G_node_first_sibling},
	{"g_node_get_root", &G_node_get_root},
	{"g_node_insert", &G_node_insert},
	{"g_node_insert_after", &G_node_insert_after},
	{"g_node_insert_before", &G_node_insert_before},
	{"g_node_is_ancestor", &G_node_is_ancestor},
	{"g_node_last_child", &G_node_last_child},
	{"g_node_last_sibling", &G_node_last_sibling},
	{"g_node_max_height", &G_node_max_height},
	{"g_node_n_children", &G_node_n_children},
	{"g_node_n_nodes", &G_node_n_nodes},
	{"g_node_new", &G_node_new},
	{"g_node_nth_child", &G_node_nth_child},
	{"g_node_pop_allocator", &G_node_pop_allocator},
	{"g_node_prepend", &G_node_prepend},
	{"g_node_push_allocator", &G_node_push_allocator},
	{"g_node_reverse_children", &G_node_reverse_children},
	{"g_node_traverse", &G_node_traverse},
	{"g_node_unlink", &G_node_unlink},
	{"g_nullify_pointer", &G_nullify_pointer},
	{"g_on_error_query", &G_on_error_query},
	{"g_on_error_stack_trace", &G_on_error_stack_trace},
	{"g_once_impl", &G_once_impl},
	{"g_once_init_enter", &G_once_init_enter},
	{"g_once_init_enter_impl", &G_once_init_enter_impl},
	{"g_once_init_leave", &G_once_init_leave},
	{"g_open", &G_open},
	{"g_option_context_add_group", &G_option_context_add_group},
	{"g_option_context_add_main_entries", &G_option_context_add_main_entries},
	{"g_option_context_free", &G_option_context_free},
	{"g_option_context_get_description", &G_option_context_get_description},
	{"g_option_context_get_help", &G_option_context_get_help},
	{"g_option_context_get_help_enabled", &G_option_context_get_help_enabled},
	{"g_option_context_get_ignore_unknown_options", &G_option_context_get_ignore_unknown_options},
	{"g_option_context_get_main_group", &G_option_context_get_main_group},
	{"g_option_context_get_summary", &G_option_context_get_summary},
	{"g_option_context_new", &G_option_context_new},
	{"g_option_context_parse", &G_option_context_parse},
	{"g_option_context_set_description", &G_option_context_set_description},
	{"g_option_context_set_help_enabled", &G_option_context_set_help_enabled},
	{"g_option_context_set_ignore_unknown_options", &G_option_context_set_ignore_unknown_options},
	{"g_option_context_set_main_group", &G_option_context_set_main_group},
	{"g_option_context_set_summary", &G_option_context_set_summary},
	{"g_option_context_set_translate_func", &G_option_context_set_translate_func},
	{"g_option_context_set_translation_domain", &G_option_context_set_translation_domain},
	{"g_option_error_quark", &G_option_error_quark},
	{"g_option_group_add_entries", &G_option_group_add_entries},
	{"g_option_group_free", &G_option_group_free},
	{"g_option_group_new", &G_option_group_new},
	{"g_option_group_set_error_hook", &G_option_group_set_error_hook},
	{"g_option_group_set_parse_hooks", &G_option_group_set_parse_hooks},
	{"g_option_group_set_translate_func", &G_option_group_set_translate_func},
	{"g_option_group_set_translation_domain", &G_option_group_set_translation_domain},
	{"g_parse_debug_string", &G_parse_debug_string},
	{"g_path_get_basename", &G_path_get_basename},
	{"g_path_get_dirname", &G_path_get_dirname},
	{"g_path_is_absolute", &G_path_is_absolute},
	{"g_path_skip_root", &G_path_skip_root},
	{"g_pattern_match", &G_pattern_match},
	{"g_pattern_match_simple", &G_pattern_match_simple},
	{"g_pattern_match_string", &G_pattern_match_string},
	{"g_pattern_spec_equal", &G_pattern_spec_equal},
	{"g_pattern_spec_free", &G_pattern_spec_free},
	{"g_pattern_spec_new", &G_pattern_spec_new},
	{"g_poll", &G_poll},
	{"g_prefix_error", &G_prefix_error},
	{"g_print", &G_print},
	{"g_printerr", &G_printerr},
	{"g_printf", &G_printf},
	{"g_printf_string_upper_bound", &G_printf_string_upper_bound},
	{"g_propagate_error", &G_propagate_error},
	{"g_propagate_prefixed_error", &G_propagate_prefixed_error},
	{"g_ptr_array_add", &G_ptr_array_add},
	{"g_ptr_array_foreach", &G_ptr_array_foreach},
	{"g_ptr_array_free", &G_ptr_array_free},
	{"g_ptr_array_new", &G_ptr_array_new},
	{"g_ptr_array_new_with_free_func", &G_ptr_array_new_with_free_func},
	{"g_ptr_array_ref", &G_ptr_array_ref},
	{"g_ptr_array_remove", &G_ptr_array_remove},
	{"g_ptr_array_remove_fast", &G_ptr_array_remove_fast},
	{"g_ptr_array_remove_index", &G_ptr_array_remove_index},
	{"g_ptr_array_remove_index_fast", &G_ptr_array_remove_index_fast},
	{"g_ptr_array_remove_range", &G_ptr_array_remove_range},
	{"g_ptr_array_set_free_func", &G_ptr_array_set_free_func},
	{"g_ptr_array_set_size", &G_ptr_array_set_size},
	{"g_ptr_array_sized_new", &G_ptr_array_sized_new},
	{"g_ptr_array_sort", &G_ptr_array_sort},
	{"g_ptr_array_sort_with_data", &G_ptr_array_sort_with_data},
	{"g_ptr_array_unref", &G_ptr_array_unref},
	{"g_qsort_with_data", &G_qsort_with_data},
	{"g_quark_from_static_string", &G_quark_from_static_string},
	{"g_quark_from_string", &G_quark_from_string},
	{"g_quark_to_string", &G_quark_to_string},
	{"g_quark_try_string", &G_quark_try_string},
	{"g_queue_clear", &G_queue_clear},
	{"g_queue_copy", &G_queue_copy},
	{"g_queue_delete_link", &G_queue_delete_link},
	{"g_queue_find", &G_queue_find},
	{"g_queue_find_custom", &G_queue_find_custom},
	{"g_queue_foreach", &G_queue_foreach},
	{"g_queue_free", &G_queue_free},
	{"g_queue_get_length", &G_queue_get_length},
	{"g_queue_index", &G_queue_index},
	{"g_queue_init", &G_queue_init},
	{"g_queue_insert_after", &G_queue_insert_after},
	{"g_queue_insert_before", &G_queue_insert_before},
	{"g_queue_insert_sorted", &G_queue_insert_sorted},
	{"g_queue_is_empty", &G_queue_is_empty},
	{"g_queue_link_index", &G_queue_link_index},
	{"g_queue_new", &G_queue_new},
	{"g_queue_peek_head", &G_queue_peek_head},
	{"g_queue_peek_head_link", &G_queue_peek_head_link},
	{"g_queue_peek_nth", &G_queue_peek_nth},
	{"g_queue_peek_nth_link", &G_queue_peek_nth_link},
	{"g_queue_peek_tail", &G_queue_peek_tail},
	{"g_queue_peek_tail_link", &G_queue_peek_tail_link},
	{"g_queue_pop_head", &G_queue_pop_head},
	{"g_queue_pop_head_link", &G_queue_pop_head_link},
	{"g_queue_pop_nth", &G_queue_pop_nth},
	{"g_queue_pop_nth_link", &G_queue_pop_nth_link},
	{"g_queue_pop_tail", &G_queue_pop_tail},
	{"g_queue_pop_tail_link", &G_queue_pop_tail_link},
	{"g_queue_push_head", &G_queue_push_head},
	{"g_queue_push_head_link", &G_queue_push_head_link},
	{"g_queue_push_nth", &G_queue_push_nth},
	{"g_queue_push_nth_link", &G_queue_push_nth_link},
	{"g_queue_push_tail", &G_queue_push_tail},
	{"g_queue_push_tail_link", &G_queue_push_tail_link},
	{"g_queue_remove", &G_queue_remove},
	{"g_queue_remove_all", &G_queue_remove_all},
	{"g_queue_reverse", &G_queue_reverse},
	{"g_queue_sort", &G_queue_sort},
	{"g_queue_unlink", &G_queue_unlink},
	{"g_rand_copy", &G_rand_copy},
	{"g_rand_double", &G_rand_double},
	{"g_rand_double_range", &G_rand_double_range},
	{"g_rand_free", &G_rand_free},
	{"g_rand_int", &G_rand_int},
	{"g_rand_int_range", &G_rand_int_range},
	{"g_rand_new", &G_rand_new},
	{"g_rand_new_with_seed", &G_rand_new_with_seed},
	{"g_rand_new_with_seed_array", &G_rand_new_with_seed_array},
	{"g_rand_set_seed", &G_rand_set_seed},
	{"g_rand_set_seed_array", &G_rand_set_seed_array},
	{"g_random_double", &G_random_double},
	{"g_random_double_range", &G_random_double_range},
	{"g_random_int", &G_random_int},
	{"g_random_int_range", &G_random_int_range},
	{"g_random_set_seed", &G_random_set_seed},
	{"g_realloc", &G_realloc},
	{"g_realloc_n", &G_realloc_n},
	{"g_regex_check_replacement", &G_regex_check_replacement},
	{"g_regex_error_quark", &G_regex_error_quark},
	{"g_regex_escape_string", &G_regex_escape_string},
	{"g_regex_get_capture_count", &G_regex_get_capture_count},
	{"g_regex_get_compile_flags", &G_regex_get_compile_flags},
	{"g_regex_get_match_flags", &G_regex_get_match_flags},
	{"g_regex_get_max_backref", &G_regex_get_max_backref},
	{"g_regex_get_pattern", &G_regex_get_pattern},
	{"g_regex_get_string_number", &G_regex_get_string_number},
	{"g_regex_match", &G_regex_match},
	{"g_regex_match_all", &G_regex_match_all},
	{"g_regex_match_all_full", &G_regex_match_all_full},
	{"g_regex_match_full", &G_regex_match_full},
	{"g_regex_match_simple", &G_regex_match_simple},
	{"g_regex_new", &G_regex_new},
	{"g_regex_ref", &G_regex_ref},
	{"g_regex_replace", &G_regex_replace},
	{"g_regex_replace_eval", &G_regex_replace_eval},
	{"g_regex_replace_literal", &G_regex_replace_literal},
	{"g_regex_split", &G_regex_split},
	{"g_regex_split_full", &G_regex_split_full},
	{"g_regex_split_simple", &G_regex_split_simple},
	{"g_regex_unref", &G_regex_unref},
	{"g_relation_count", &G_relation_count},
	{"g_relation_delete", &G_relation_delete},
	{"g_relation_destroy", &G_relation_destroy},
	{"g_relation_exists", &G_relation_exists},
	{"g_relation_index", &G_relation_index},
	{"g_relation_insert", &G_relation_insert},
	{"g_relation_new", &G_relation_new},
	{"g_relation_print", &G_relation_print},
	{"g_relation_select", &G_relation_select},
	{"g_reload_user_special_dirs_cache", &G_reload_user_special_dirs_cache},
	{"g_remove", &G_remove},
	{"g_rename", &G_rename},
	{"g_return_if_fail_warning", &G_return_if_fail_warning},
	{"g_rmdir", &G_rmdir},
	{"g_scanner_cur_line", &G_scanner_cur_line},
	{"g_scanner_cur_position", &G_scanner_cur_position},
	{"g_scanner_cur_token", &G_scanner_cur_token},
	{"g_scanner_cur_value", &G_scanner_cur_value},
	{"g_scanner_destroy", &G_scanner_destroy},
	{"g_scanner_eof", &G_scanner_eof},
	{"g_scanner_error", &G_scanner_error},
	{"g_scanner_get_next_token", &G_scanner_get_next_token},
	{"g_scanner_input_file", &G_scanner_input_file},
	{"g_scanner_input_text", &G_scanner_input_text},
	{"g_scanner_lookup_symbol", &G_scanner_lookup_symbol},
	{"g_scanner_new", &G_scanner_new},
	{"g_scanner_peek_next_token", &G_scanner_peek_next_token},
	{"g_scanner_scope_add_symbol", &G_scanner_scope_add_symbol},
	{"g_scanner_scope_foreach_symbol", &G_scanner_scope_foreach_symbol},
	{"g_scanner_scope_lookup_symbol", &G_scanner_scope_lookup_symbol},
	{"g_scanner_scope_remove_symbol", &G_scanner_scope_remove_symbol},
	{"g_scanner_set_scope", &G_scanner_set_scope},
	{"g_scanner_sync_file_offset", &G_scanner_sync_file_offset},
	{"g_scanner_unexp_token", &G_scanner_unexp_token},
	{"g_scanner_warn", &G_scanner_warn},
	{"g_sequence_append", &G_sequence_append},
	{"g_sequence_foreach", &G_sequence_foreach},
	{"g_sequence_foreach_range", &G_sequence_foreach_range},
	{"g_sequence_free", &G_sequence_free},
	{"g_sequence_get", &G_sequence_get},
	{"g_sequence_get_begin_iter", &G_sequence_get_begin_iter},
	{"g_sequence_get_end_iter", &G_sequence_get_end_iter},
	{"g_sequence_get_iter_at_pos", &G_sequence_get_iter_at_pos},
	{"g_sequence_get_length", &G_sequence_get_length},
	{"g_sequence_insert_before", &G_sequence_insert_before},
	{"g_sequence_insert_sorted", &G_sequence_insert_sorted},
	{"g_sequence_insert_sorted_iter", &G_sequence_insert_sorted_iter},
	{"g_sequence_iter_compare", &G_sequence_iter_compare},
	{"g_sequence_iter_get_position", &G_sequence_iter_get_position},
	{"g_sequence_iter_get_sequence", &G_sequence_iter_get_sequence},
	{"g_sequence_iter_is_begin", &G_sequence_iter_is_begin},
	{"g_sequence_iter_is_end", &G_sequence_iter_is_end},
	{"g_sequence_iter_move", &G_sequence_iter_move},
	{"g_sequence_iter_next", &G_sequence_iter_next},
	{"g_sequence_iter_prev", &G_sequence_iter_prev},
	{"g_sequence_lookup", &G_sequence_lookup},
	{"g_sequence_lookup_iter", &G_sequence_lookup_iter},
	{"g_sequence_move", &G_sequence_move},
	{"g_sequence_move_range", &G_sequence_move_range},
	{"g_sequence_new", &G_sequence_new},
	{"g_sequence_prepend", &G_sequence_prepend},
	{"g_sequence_range_get_midpoint", &G_sequence_range_get_midpoint},
	{"g_sequence_remove", &G_sequence_remove},
	{"g_sequence_remove_range", &G_sequence_remove_range},
	{"g_sequence_search", &G_sequence_search},
	{"g_sequence_search_iter", &G_sequence_search_iter},
	{"g_sequence_set", &G_sequence_set},
	{"g_sequence_sort", &G_sequence_sort},
	{"g_sequence_sort_changed", &G_sequence_sort_changed},
	{"g_sequence_sort_changed_iter", &G_sequence_sort_changed_iter},
	{"g_sequence_sort_iter", &G_sequence_sort_iter},
	{"g_sequence_swap", &G_sequence_swap},
	{"g_set_application_name", &G_set_application_name},
	{"g_set_error", &G_set_error},
	{"g_set_error_literal", &G_set_error_literal},
	{"g_set_prgname", &G_set_prgname},
	{"g_set_print_handler", &G_set_print_handler},
	{"g_set_printerr_handler", &G_set_printerr_handler},
	// Windows: _utf8 {"g_setenv", &G_setenv},
	{"g_setenv_utf8", &G_setenv},
	{"g_shell_error_quark", &G_shell_error_quark},
	{"g_shell_parse_argv", &G_shell_parse_argv},
	{"g_shell_quote", &G_shell_quote},
	{"g_shell_unquote", &G_shell_unquote},
	{"g_slice_alloc", &G_slice_alloc},
	{"g_slice_alloc0", &G_slice_alloc0},
	{"g_slice_copy", &G_slice_copy},
	{"g_slice_free1", &G_slice_free1},
	{"g_slice_free_chain_with_offset", &G_slice_free_chain_with_offset},
	{"g_slice_get_config", &G_slice_get_config},
	{"g_slice_get_config_state", &G_slice_get_config_state},
	{"g_slice_set_config", &G_slice_set_config},
	{"g_slist_alloc", &G_slist_alloc},
	{"g_slist_append", &G_slist_append},
	{"g_slist_concat", &G_slist_concat},
	{"g_slist_copy", &G_slist_copy},
	{"g_slist_delete_link", &G_slist_delete_link},
	{"g_slist_find", &G_slist_find},
	{"g_slist_find_custom", &G_slist_find_custom},
	{"g_slist_foreach", &G_slist_foreach},
	{"g_slist_free", &G_slist_free},
	{"g_slist_free_1", &G_slist_free_1},
	{"g_slist_free_full", &G_slist_free_full},
	{"g_slist_index", &G_slist_index},
	{"g_slist_insert", &G_slist_insert},
	{"g_slist_insert_before", &G_slist_insert_before},
	{"g_slist_insert_sorted", &G_slist_insert_sorted},
	{"g_slist_insert_sorted_with_data", &G_slist_insert_sorted_with_data},
	{"g_slist_last", &G_slist_last},
	{"g_slist_length", &G_slist_length},
	{"g_slist_nth", &G_slist_nth},
	{"g_slist_nth_data", &G_slist_nth_data},
	{"g_slist_pop_allocator", &G_slist_pop_allocator},
	{"g_slist_position", &G_slist_position},
	{"g_slist_prepend", &G_slist_prepend},
	{"g_slist_push_allocator", &G_slist_push_allocator},
	{"g_slist_remove", &G_slist_remove},
	{"g_slist_remove_all", &G_slist_remove_all},
	{"g_slist_remove_link", &G_slist_remove_link},
	{"g_slist_reverse", &G_slist_reverse},
	{"g_slist_sort", &G_slist_sort},
	{"g_slist_sort_with_data", &G_slist_sort_with_data},
	{"g_snprintf", &G_snprintf},
	{"g_source_add_child_source", &G_source_add_child_source},
	{"g_source_add_poll", &G_source_add_poll},
	{"g_source_attach", &G_source_attach},
	{"g_source_destroy", &G_source_destroy},
	{"g_source_get_can_recurse", &G_source_get_can_recurse},
	{"g_source_get_context", &G_source_get_context},
	{"g_source_get_current_time", &G_source_get_current_time},
	{"g_source_get_id", &G_source_get_id},
	{"g_source_get_name", &G_source_get_name},
	{"g_source_get_priority", &G_source_get_priority},
	{"g_source_get_time", &G_source_get_time},
	{"g_source_is_destroyed", &G_source_is_destroyed},
	{"g_source_new", &G_source_new},
	{"g_source_ref", &G_source_ref},
	{"g_source_remove", &G_source_remove},
	{"g_source_remove_by_funcs_user_data", &G_source_remove_by_funcs_user_data},
	{"g_source_remove_by_user_data", &G_source_remove_by_user_data},
	{"g_source_remove_child_source", &G_source_remove_child_source},
	{"g_source_remove_poll", &G_source_remove_poll},
	{"g_source_set_callback", &G_source_set_callback},
	{"g_source_set_callback_indirect", &G_source_set_callback_indirect},
	{"g_source_set_can_recurse", &G_source_set_can_recurse},
	{"g_source_set_funcs", &G_source_set_funcs},
	{"g_source_set_name", &G_source_set_name},
	{"g_source_set_name_by_id", &G_source_set_name_by_id},
	{"g_source_set_priority", &G_source_set_priority},
	{"g_source_unref", &G_source_unref},
	{"g_spaced_primes_closest", &G_spaced_primes_closest},
	// Windows: _utf8 {"g_spawn_async", &G_spawn_async},
	{"g_spawn_async_utf8", &G_spawn_async},
	// Windows: _utf8 {"g_spawn_async_with_pipes", &G_spawn_async_with_pipes},
	{"g_spawn_async_with_pipes_utf8", &G_spawn_async_with_pipes},
	{"g_spawn_close_pid", &G_spawn_close_pid},
	// Windows: _utf8 {"g_spawn_command_line_async", &G_spawn_command_line_async},
	{"g_spawn_command_line_async_utf8", &G_spawn_command_line_async},
	// Windows: _utf8 {"g_spawn_command_line_sync", &G_spawn_command_line_sync},
	{"g_spawn_command_line_sync_utf8", &G_spawn_command_line_sync},
	{"g_spawn_error_quark", &G_spawn_error_quark},
	// Windows: _utf8 {"g_spawn_sync", &G_spawn_sync},
	{"g_spawn_sync_utf8", &G_spawn_sync},
	{"g_sprintf", &G_sprintf},
	{"g_stat", &G_stat},
	{"g_static_mutex_free", &G_static_mutex_free},
	{"g_static_mutex_get_mutex_impl", &G_static_mutex_get_mutex_impl},
	{"g_static_mutex_init", &G_static_mutex_init},
	{"g_static_private_free", &G_static_private_free},
	{"g_static_private_get", &G_static_private_get},
	{"g_static_private_init", &G_static_private_init},
	{"g_static_private_set", &G_static_private_set},
	{"g_static_rec_mutex_free", &G_static_rec_mutex_free},
	{"g_static_rec_mutex_init", &G_static_rec_mutex_init},
	{"g_static_rec_mutex_lock", &G_static_rec_mutex_lock},
	{"g_static_rec_mutex_lock_full", &G_static_rec_mutex_lock_full},
	{"g_static_rec_mutex_trylock", &G_static_rec_mutex_trylock},
	{"g_static_rec_mutex_unlock", &G_static_rec_mutex_unlock},
	{"g_static_rec_mutex_unlock_full", &G_static_rec_mutex_unlock_full},
	{"g_static_rw_lock_free", &G_static_rw_lock_free},
	{"g_static_rw_lock_init", &G_static_rw_lock_init},
	{"g_static_rw_lock_reader_lock", &G_static_rw_lock_reader_lock},
	{"g_static_rw_lock_reader_trylock", &G_static_rw_lock_reader_trylock},
	{"g_static_rw_lock_reader_unlock", &G_static_rw_lock_reader_unlock},
	{"g_static_rw_lock_writer_lock", &G_static_rw_lock_writer_lock},
	{"g_static_rw_lock_writer_trylock", &G_static_rw_lock_writer_trylock},
	{"g_static_rw_lock_writer_unlock", &G_static_rw_lock_writer_unlock},
	{"g_stpcpy", &G_stpcpy},
	{"g_str_equal", &G_str_equal},
	{"g_str_has_prefix", &G_str_has_prefix},
	{"g_str_has_suffix", &G_str_has_suffix},
	{"g_str_hash", &G_str_hash},
	{"g_strcanon", &G_strcanon},
	{"g_strcasecmp", &G_strcasecmp},
	{"g_strchomp", &G_strchomp},
	{"g_strchug", &G_strchug},
	{"g_strcmp0", &G_strcmp0},
	{"g_strcompress", &G_strcompress},
	{"g_strconcat", &G_strconcat},
	{"g_strdelimit", &G_strdelimit},
	{"g_strdown", &G_strdown},
	{"g_strdup", &G_strdup},
	{"g_strdup_printf", &G_strdup_printf},
	{"g_strdup_vprintf", &G_strdup_vprintf},
	{"g_strdupv", &G_strdupv},
	{"g_strerror", &G_strerror},
	{"g_strescape", &G_strescape},
	{"g_strfreev", &G_strfreev},
	{"g_string_append", &G_string_append},
	{"g_string_append_c", &G_string_append_c},
	{"g_string_append_len", &G_string_append_len},
	{"g_string_append_printf", &G_string_append_printf},
	{"g_string_append_unichar", &G_string_append_unichar},
	{"g_string_append_uri_escaped", &G_string_append_uri_escaped},
	{"g_string_append_vprintf", &G_string_append_vprintf},
	{"g_string_ascii_down", &G_string_ascii_down},
	{"g_string_ascii_up", &G_string_ascii_up},
	{"g_string_assign", &G_string_assign},
	{"g_string_chunk_clear", &G_string_chunk_clear},
	{"g_string_chunk_free", &G_string_chunk_free},
	{"g_string_chunk_insert", &G_string_chunk_insert},
	{"g_string_chunk_insert_const", &G_string_chunk_insert_const},
	{"g_string_chunk_insert_len", &G_string_chunk_insert_len},
	{"g_string_chunk_new", &G_string_chunk_new},
	{"g_string_down", &G_string_down},
	{"g_string_equal", &G_string_equal},
	{"g_string_erase", &G_string_erase},
	{"g_string_free", &G_string_free},
	{"g_string_hash", &G_string_hash},
	{"g_string_insert", &G_string_insert},
	{"g_string_insert_c", &G_string_insert_c},
	{"g_string_insert_len", &G_string_insert_len},
	{"g_string_insert_unichar", &G_string_insert_unichar},
	{"g_string_new", &G_string_new},
	{"g_string_new_len", &G_string_new_len},
	{"g_string_overwrite", &G_string_overwrite},
	{"g_string_overwrite_len", &G_string_overwrite_len},
	{"g_string_prepend", &G_string_prepend},
	{"g_string_prepend_c", &G_string_prepend_c},
	{"g_string_prepend_len", &G_string_prepend_len},
	{"g_string_prepend_unichar", &G_string_prepend_unichar},
	{"g_string_printf", &G_string_printf},
	{"g_string_set_size", &G_string_set_size},
	{"g_string_sized_new", &G_string_sized_new},
	{"g_string_truncate", &G_string_truncate},
	{"g_string_up", &G_string_up},
	{"g_string_vprintf", &G_string_vprintf},
	{"g_strip_context", &G_strip_context},
	{"g_strjoin", &G_strjoin},
	{"g_strjoinv", &G_strjoinv},
	{"g_strlcat", &G_strlcat},
	{"g_strlcpy", &G_strlcpy},
	{"g_strncasecmp", &G_strncasecmp},
	{"g_strndup", &G_strndup},
	{"g_strnfill", &G_strnfill},
	{"g_strreverse", &G_strreverse},
	{"g_strrstr", &G_strrstr},
	{"g_strrstr_len", &G_strrstr_len},
	{"g_strsignal", &G_strsignal},
	{"g_strsplit", &G_strsplit},
	{"g_strsplit_set", &G_strsplit_set},
	{"g_strstr_len", &G_strstr_len},
	{"g_strtod", &G_strtod},
	{"g_strup", &G_strup},
	{"g_strv_length", &G_strv_length},
	{"g_test_add_data_func", &G_test_add_data_func},
	{"g_test_add_func", &G_test_add_func},
	{"g_test_add_vtable", &G_test_add_vtable},
	{"g_test_bug", &G_test_bug},
	{"g_test_bug_base", &G_test_bug_base},
	{"g_test_create_case", &G_test_create_case},
	{"g_test_create_suite", &G_test_create_suite},
	{"g_test_get_root", &G_test_get_root},
	{"g_test_init", &G_test_init},
	{"g_test_log_buffer_free", &G_test_log_buffer_free},
	{"g_test_log_buffer_new", &G_test_log_buffer_new},
	{"g_test_log_buffer_pop", &G_test_log_buffer_pop},
	{"g_test_log_buffer_push", &G_test_log_buffer_push},
	{"g_test_log_msg_free", &G_test_log_msg_free},
	{"g_test_log_set_fatal_handler", &G_test_log_set_fatal_handler},
	{"g_test_log_type_name", &G_test_log_type_name},
	{"g_test_maximized_result", &G_test_maximized_result},
	{"g_test_message", &G_test_message},
	{"g_test_minimized_result", &G_test_minimized_result},
	{"g_test_queue_destroy", &G_test_queue_destroy},
	{"g_test_queue_free", &G_test_queue_free},
	{"g_test_rand_double", &G_test_rand_double},
	{"g_test_rand_double_range", &G_test_rand_double_range},
	{"g_test_rand_int", &G_test_rand_int},
	{"g_test_rand_int_range", &G_test_rand_int_range},
	{"g_test_run", &G_test_run},
	{"g_test_run_suite", &G_test_run_suite},
	{"g_test_suite_add", &G_test_suite_add},
	{"g_test_suite_add_suite", &G_test_suite_add_suite},
	{"g_test_timer_elapsed", &G_test_timer_elapsed},
	{"g_test_timer_last", &G_test_timer_last},
	{"g_test_timer_start", &G_test_timer_start},
	{"g_test_trap_assertions", &G_test_trap_assertions},
	{"g_test_trap_fork", &G_test_trap_fork},
	{"g_test_trap_has_passed", &G_test_trap_has_passed},
	{"g_test_trap_reached_timeout", &G_test_trap_reached_timeout},
	{"g_thread_create_full", &G_thread_create_full},
	{"g_thread_error_quark", &G_thread_error_quark},
	{"g_thread_exit", &G_thread_exit},
	{"g_thread_foreach", &G_thread_foreach},
	{"g_thread_get_initialized", &G_thread_get_initialized},
	// Undocumented {"g_thread_init_glib", &G_thread_init_glib},
	{"g_thread_join", &G_thread_join},
	{"g_thread_pool_free", &G_thread_pool_free},
	{"g_thread_pool_get_max_idle_time", &G_thread_pool_get_max_idle_time},
	{"g_thread_pool_get_max_threads", &G_thread_pool_get_max_threads},
	{"g_thread_pool_get_max_unused_threads", &G_thread_pool_get_max_unused_threads},
	{"g_thread_pool_get_num_threads", &G_thread_pool_get_num_threads},
	{"g_thread_pool_get_num_unused_threads", &G_thread_pool_get_num_unused_threads},
	{"g_thread_pool_new", &G_thread_pool_new},
	{"g_thread_pool_push", &G_thread_pool_push},
	{"g_thread_pool_set_max_idle_time", &G_thread_pool_set_max_idle_time},
	{"g_thread_pool_set_max_threads", &G_thread_pool_set_max_threads},
	{"g_thread_pool_set_max_unused_threads", &G_thread_pool_set_max_unused_threads},
	{"g_thread_pool_set_sort_function", &G_thread_pool_set_sort_function},
	{"g_thread_pool_stop_unused_threads", &G_thread_pool_stop_unused_threads},
	{"g_thread_pool_unprocessed", &G_thread_pool_unprocessed},
	{"g_thread_self", &G_thread_self},
	{"g_thread_set_priority", &G_thread_set_priority},
	{"g_time_val_add", &G_time_val_add},
	{"g_time_val_from_iso8601", &G_time_val_from_iso8601},
	{"g_time_val_to_iso8601", &G_time_val_to_iso8601},
	{"g_time_zone_adjust_time", &G_time_zone_adjust_time},
	{"g_time_zone_find_interval", &G_time_zone_find_interval},
	{"g_time_zone_get_abbreviation", &G_time_zone_get_abbreviation},
	{"g_time_zone_get_offset", &G_time_zone_get_offset},
	{"g_time_zone_is_dst", &G_time_zone_is_dst},
	{"g_time_zone_new", &G_time_zone_new},
	{"g_time_zone_new_local", &G_time_zone_new_local},
	{"g_time_zone_new_utc", &G_time_zone_new_utc},
	{"g_time_zone_ref", &G_time_zone_ref},
	{"g_time_zone_unref", &G_time_zone_unref},
	{"g_timeout_add", &G_timeout_add},
	{"g_timeout_add_full", &G_timeout_add_full},
	{"g_timeout_add_seconds", &G_timeout_add_seconds},
	{"g_timeout_add_seconds_full", &G_timeout_add_seconds_full},
	{"g_timeout_source_new", &G_timeout_source_new},
	{"g_timeout_source_new_seconds", &G_timeout_source_new_seconds},
	{"g_timer_continue", &G_timer_continue},
	{"g_timer_destroy", &G_timer_destroy},
	{"g_timer_elapsed", &G_timer_elapsed},
	{"g_timer_new", &G_timer_new},
	{"g_timer_reset", &G_timer_reset},
	{"g_timer_start", &G_timer_start},
	{"g_timer_stop", &G_timer_stop},
	{"g_trash_stack_height", &G_trash_stack_height},
	{"g_trash_stack_peek", &G_trash_stack_peek},
	{"g_trash_stack_pop", &G_trash_stack_pop},
	{"g_trash_stack_push", &G_trash_stack_push},
	{"g_tree_destroy", &G_tree_destroy},
	{"g_tree_foreach", &G_tree_foreach},
	{"g_tree_height", &G_tree_height},
	{"g_tree_insert", &G_tree_insert},
	{"g_tree_lookup", &G_tree_lookup},
	{"g_tree_lookup_extended", &G_tree_lookup_extended},
	{"g_tree_new", &G_tree_new},
	{"g_tree_new_full", &G_tree_new_full},
	{"g_tree_new_with_data", &G_tree_new_with_data},
	{"g_tree_nnodes", &G_tree_nnodes},
	{"g_tree_ref", &G_tree_ref},
	{"g_tree_remove", &G_tree_remove},
	{"g_tree_replace", &G_tree_replace},
	{"g_tree_search", &G_tree_search},
	{"g_tree_steal", &G_tree_steal},
	{"g_tree_traverse", &G_tree_traverse},
	{"g_tree_unref", &G_tree_unref},
	{"g_try_malloc", &G_try_malloc},
	{"g_try_malloc0", &G_try_malloc0},
	{"g_try_malloc0_n", &G_try_malloc0_n},
	{"g_try_malloc_n", &G_try_malloc_n},
	{"g_try_realloc", &G_try_realloc},
	{"g_try_realloc_n", &G_try_realloc_n},
	{"g_tuples_destroy", &G_tuples_destroy},
	{"g_tuples_index", &G_tuples_index},
	{"g_ucs4_to_utf16", &G_ucs4_to_utf16},
	{"g_ucs4_to_utf8", &G_ucs4_to_utf8},
	{"g_unichar_break_type", &G_unichar_break_type},
	{"g_unichar_combining_class", &G_unichar_combining_class},
	{"g_unichar_digit_value", &G_unichar_digit_value},
	{"g_unichar_get_mirror_char", &G_unichar_get_mirror_char},
	{"g_unichar_get_script", &G_unichar_get_script},
	{"g_unichar_isalnum", &G_unichar_isalnum},
	{"g_unichar_isalpha", &G_unichar_isalpha},
	{"g_unichar_iscntrl", &G_unichar_iscntrl},
	{"g_unichar_isdefined", &G_unichar_isdefined},
	{"g_unichar_isdigit", &G_unichar_isdigit},
	{"g_unichar_isgraph", &G_unichar_isgraph},
	{"g_unichar_islower", &G_unichar_islower},
	{"g_unichar_ismark", &G_unichar_ismark},
	{"g_unichar_isprint", &G_unichar_isprint},
	{"g_unichar_ispunct", &G_unichar_ispunct},
	{"g_unichar_isspace", &G_unichar_isspace},
	{"g_unichar_istitle", &G_unichar_istitle},
	{"g_unichar_isupper", &G_unichar_isupper},
	{"g_unichar_iswide", &G_unichar_iswide},
	{"g_unichar_iswide_cjk", &G_unichar_iswide_cjk},
	{"g_unichar_isxdigit", &G_unichar_isxdigit},
	{"g_unichar_iszerowidth", &G_unichar_iszerowidth},
	{"g_unichar_to_utf8", &G_unichar_to_utf8},
	{"g_unichar_tolower", &G_unichar_tolower},
	{"g_unichar_totitle", &G_unichar_totitle},
	{"g_unichar_toupper", &G_unichar_toupper},
	{"g_unichar_type", &G_unichar_type},
	{"g_unichar_validate", &G_unichar_validate},
	{"g_unichar_xdigit_value", &G_unichar_xdigit_value},
	{"g_unicode_canonical_decomposition", &G_unicode_canonical_decomposition},
	{"g_unicode_canonical_ordering", &G_unicode_canonical_ordering},
	{"g_unlink", &G_unlink},
	// Windows: _utf8 {"g_unsetenv", &G_unsetenv},
	{"g_unsetenv_utf8", &G_unsetenv},
	{"g_uri_escape_string", &G_uri_escape_string},
	{"g_uri_list_extract_uris", &G_uri_list_extract_uris},
	{"g_uri_parse_scheme", &G_uri_parse_scheme},
	{"g_uri_unescape_segment", &G_uri_unescape_segment},
	{"g_uri_unescape_string", &G_uri_unescape_string},
	{"g_usleep", &G_usleep},
	{"g_utf16_to_ucs4", &G_utf16_to_ucs4},
	{"g_utf16_to_utf8", &G_utf16_to_utf8},
	{"g_utf8_casefold", &G_utf8_casefold},
	{"g_utf8_collate", &G_utf8_collate},
	{"g_utf8_collate_key", &G_utf8_collate_key},
	{"g_utf8_collate_key_for_filename", &G_utf8_collate_key_for_filename},
	{"g_utf8_find_next_char", &G_utf8_find_next_char},
	{"g_utf8_find_prev_char", &G_utf8_find_prev_char},
	{"g_utf8_get_char", &G_utf8_get_char},
	{"g_utf8_get_char_validated", &G_utf8_get_char_validated},
	{"g_utf8_normalize", &G_utf8_normalize},
	{"g_utf8_offset_to_pointer", &G_utf8_offset_to_pointer},
	{"g_utf8_pointer_to_offset", &G_utf8_pointer_to_offset},
	{"g_utf8_prev_char", &G_utf8_prev_char},
	{"g_utf8_strchr", &G_utf8_strchr},
	{"g_utf8_strdown", &G_utf8_strdown},
	{"g_utf8_strlen", &G_utf8_strlen},
	{"g_utf8_strncpy", &G_utf8_strncpy},
	{"g_utf8_strrchr", &G_utf8_strrchr},
	{"g_utf8_strreverse", &G_utf8_strreverse},
	{"g_utf8_strup", &G_utf8_strup},
	{"g_utf8_to_ucs4", &G_utf8_to_ucs4},
	{"g_utf8_to_ucs4_fast", &G_utf8_to_ucs4_fast},
	{"g_utf8_to_utf16", &G_utf8_to_utf16},
	{"g_utf8_validate", &G_utf8_validate},
	{"g_utime", &G_utime},
	{"g_variant_builder_add", &G_variant_builder_add},
	{"g_variant_builder_add_parsed", &G_variant_builder_add_parsed},
	{"g_variant_builder_add_value", &G_variant_builder_add_value},
	{"g_variant_builder_clear", &G_variant_builder_clear},
	{"g_variant_builder_close", &G_variant_builder_close},
	{"g_variant_builder_end", &G_variant_builder_end},
	{"g_variant_builder_init", &G_variant_builder_init},
	{"g_variant_builder_new", &G_variant_builder_new},
	{"g_variant_builder_open", &G_variant_builder_open},
	{"g_variant_builder_ref", &G_variant_builder_ref},
	{"g_variant_builder_unref", &G_variant_builder_unref},
	{"g_variant_byteswap", &G_variant_byteswap},
	{"g_variant_classify", &G_variant_classify},
	{"g_variant_compare", &G_variant_compare},
	{"g_variant_dup_bytestring", &G_variant_dup_bytestring},
	{"g_variant_dup_bytestring_array", &G_variant_dup_bytestring_array},
	{"g_variant_dup_string", &G_variant_dup_string},
	{"g_variant_dup_strv", &G_variant_dup_strv},
	{"g_variant_equal", &G_variant_equal},
	// Undocumented {"g_variant_format_string_scan", &G_variant_format_string_scan},
	// Undocumented {"g_variant_format_string_scan_type", &G_variant_format_string_scan_type},
	{"g_variant_get", &G_variant_get},
	{"g_variant_get_boolean", &G_variant_get_boolean},
	{"g_variant_get_byte", &G_variant_get_byte},
	{"g_variant_get_bytestring", &G_variant_get_bytestring},
	{"g_variant_get_bytestring_array", &G_variant_get_bytestring_array},
	{"g_variant_get_child", &G_variant_get_child},
	{"g_variant_get_child_value", &G_variant_get_child_value},
	{"g_variant_get_data", &G_variant_get_data},
	{"g_variant_get_double", &G_variant_get_double},
	{"g_variant_get_fixed_array", &G_variant_get_fixed_array},
	{"g_variant_get_handle", &G_variant_get_handle},
	{"g_variant_get_int16", &G_variant_get_int16},
	{"g_variant_get_int32", &G_variant_get_int32},
	{"g_variant_get_int64", &G_variant_get_int64},
	{"g_variant_get_maybe", &G_variant_get_maybe},
	{"g_variant_get_normal_form", &G_variant_get_normal_form},
	{"g_variant_get_size", &G_variant_get_size},
	{"g_variant_get_string", &G_variant_get_string},
	{"g_variant_get_strv", &G_variant_get_strv},
	{"g_variant_get_type", &G_variant_get_type},
	{"g_variant_get_type_string", &G_variant_get_type_string},
	{"g_variant_get_uint16", &G_variant_get_uint16},
	{"g_variant_get_uint32", &G_variant_get_uint32},
	{"g_variant_get_uint64", &G_variant_get_uint64},
	{"g_variant_get_va", &G_variant_get_va},
	{"g_variant_get_variant", &G_variant_get_variant},
	{"g_variant_hash", &G_variant_hash},
	{"g_variant_is_container", &G_variant_is_container},
	{"g_variant_is_floating", &G_variant_is_floating},
	{"g_variant_is_normal_form", &G_variant_is_normal_form},
	{"g_variant_is_object_path", &G_variant_is_object_path},
	{"g_variant_is_of_type", &G_variant_is_of_type},
	{"g_variant_is_signature", &G_variant_is_signature},
	{"g_variant_iter_copy", &G_variant_iter_copy},
	{"g_variant_iter_free", &G_variant_iter_free},
	{"g_variant_iter_init", &G_variant_iter_init},
	{"g_variant_iter_loop", &G_variant_iter_loop},
	{"g_variant_iter_n_children", &G_variant_iter_n_children},
	{"g_variant_iter_new", &G_variant_iter_new},
	{"g_variant_iter_next", &G_variant_iter_next},
	{"g_variant_iter_next_value", &G_variant_iter_next_value},
	{"g_variant_lookup", &G_variant_lookup},
	{"g_variant_lookup_value", &G_variant_lookup_value},
	{"g_variant_n_children", &G_variant_n_children},
	{"g_variant_new", &G_variant_new},
	{"g_variant_new_array", &G_variant_new_array},
	{"g_variant_new_boolean", &G_variant_new_boolean},
	{"g_variant_new_byte", &G_variant_new_byte},
	{"g_variant_new_bytestring", &G_variant_new_bytestring},
	{"g_variant_new_bytestring_array", &G_variant_new_bytestring_array},
	{"g_variant_new_dict_entry", &G_variant_new_dict_entry},
	{"g_variant_new_double", &G_variant_new_double},
	{"g_variant_new_from_data", &G_variant_new_from_data},
	{"g_variant_new_handle", &G_variant_new_handle},
	{"g_variant_new_int16", &G_variant_new_int16},
	{"g_variant_new_int32", &G_variant_new_int32},
	{"g_variant_new_int64", &G_variant_new_int64},
	{"g_variant_new_maybe", &G_variant_new_maybe},
	{"g_variant_new_object_path", &G_variant_new_object_path},
	{"g_variant_new_parsed", &G_variant_new_parsed},
	{"g_variant_new_parsed_va", &G_variant_new_parsed_va},
	{"g_variant_new_signature", &G_variant_new_signature},
	{"g_variant_new_string", &G_variant_new_string},
	{"g_variant_new_strv", &G_variant_new_strv},
	{"g_variant_new_tuple", &G_variant_new_tuple},
	{"g_variant_new_uint16", &G_variant_new_uint16},
	{"g_variant_new_uint32", &G_variant_new_uint32},
	{"g_variant_new_uint64", &G_variant_new_uint64},
	{"g_variant_new_va", &G_variant_new_va},
	{"g_variant_new_variant", &G_variant_new_variant},
	{"g_variant_parse", &G_variant_parse},
	{"g_variant_parser_get_error_quark", &G_variant_parser_get_error_quark},
	{"g_variant_print", &G_variant_print},
	{"g_variant_print_string", &G_variant_print_string},
	{"g_variant_ref", &G_variant_ref},
	{"g_variant_ref_sink", &G_variant_ref_sink},
	// Undocumented {"g_variant_serialised_byteswap", &G_variant_serialised_byteswap},
	// Undocumented {"g_variant_serialised_get_child", &G_variant_serialised_get_child},
	// Undocumented {"g_variant_serialised_is_normal", &G_variant_serialised_is_normal},
	// Undocumented {"g_variant_serialised_n_children", &G_variant_serialised_n_children},
	// Undocumented {"g_variant_serialiser_is_object_path", &G_variant_serialiser_is_object_path},
	// Undocumented {"g_variant_serialiser_is_signature", &G_variant_serialiser_is_signature},
	// Undocumented {"g_variant_serialiser_is_string", &G_variant_serialiser_is_string},
	// Undocumented {"g_variant_serialiser_needed_size", &G_variant_serialiser_needed_size},
	// Undocumented {"g_variant_serialiser_serialise", &G_variant_serialiser_serialise},
	{"g_variant_store", &G_variant_store},
	{"g_variant_type_checked_", &G_variant_type_checked_},
	{"g_variant_type_copy", &G_variant_type_copy},
	{"g_variant_type_dup_string", &G_variant_type_dup_string},
	{"g_variant_type_element", &G_variant_type_element},
	{"g_variant_type_equal", &G_variant_type_equal},
	{"g_variant_type_first", &G_variant_type_first},
	{"g_variant_type_free", &G_variant_type_free},
	{"g_variant_type_get_string_length", &G_variant_type_get_string_length},
	{"g_variant_type_hash", &G_variant_type_hash},
	// Undocumented {"g_variant_type_info_assert_no_infos", &G_variant_type_info_assert_no_infos},
	// Undocumented {"g_variant_type_info_element", &G_variant_type_info_element},
	// Undocumented {"g_variant_type_info_get", &G_variant_type_info_get},
	// Undocumented {"g_variant_type_info_get_type_string", &G_variant_type_info_get_type_string},
	// Undocumented {"g_variant_type_info_member_info", &G_variant_type_info_member_info},
	// Undocumented {"g_variant_type_info_n_members", &G_variant_type_info_n_members},
	// Undocumented {"g_variant_type_info_query", &G_variant_type_info_query},
	// Undocumented {"g_variant_type_info_query_element", &G_variant_type_info_query_element},
	// Undocumented {"g_variant_type_info_ref", &G_variant_type_info_ref},
	// Undocumented {"g_variant_type_info_unref", &G_variant_type_info_unref},
	{"g_variant_type_is_array", &G_variant_type_is_array},
	{"g_variant_type_is_basic", &G_variant_type_is_basic},
	{"g_variant_type_is_container", &G_variant_type_is_container},
	{"g_variant_type_is_definite", &G_variant_type_is_definite},
	{"g_variant_type_is_dict_entry", &G_variant_type_is_dict_entry},
	{"g_variant_type_is_maybe", &G_variant_type_is_maybe},
	{"g_variant_type_is_subtype_of", &G_variant_type_is_subtype_of},
	{"g_variant_type_is_tuple", &G_variant_type_is_tuple},
	{"g_variant_type_is_variant", &G_variant_type_is_variant},
	{"g_variant_type_key", &G_variant_type_key},
	{"g_variant_type_n_items", &G_variant_type_n_items},
	{"g_variant_type_new", &G_variant_type_new},
	{"g_variant_type_new_array", &G_variant_type_new_array},
	{"g_variant_type_new_dict_entry", &G_variant_type_new_dict_entry},
	{"g_variant_type_new_maybe", &G_variant_type_new_maybe},
	{"g_variant_type_new_tuple", &G_variant_type_new_tuple},
	{"g_variant_type_next", &G_variant_type_next},
	{"g_variant_type_peek_string", &G_variant_type_peek_string},
	{"g_variant_type_string_is_valid", &G_variant_type_string_is_valid},
	{"g_variant_type_string_scan", &G_variant_type_string_scan},
	{"g_variant_type_value", &G_variant_type_value},
	{"g_variant_unref", &G_variant_unref},
	{"g_vasprintf", &G_vasprintf},
	{"g_vfprintf", &G_vfprintf},
	{"g_vprintf", &G_vprintf},
	{"g_vsnprintf", &G_vsnprintf},
	{"g_vsprintf", &G_vsprintf},
	{"g_warn_message", &G_warn_message},
	{"g_win32_error_message", &G_win32_error_message},
	{"g_win32_ftruncate", &G_win32_ftruncate},
	// Windows: _utf8 {"g_win32_get_package_installation_directory", &G_win32_get_package_installation_directory},
	{"g_win32_get_package_installation_directory_of_module", &G_win32_get_package_installation_directory_of_module},
	{"g_win32_get_package_installation_directory_utf8", &G_win32_get_package_installation_directory},
	// Windows: _utf8 {"g_win32_get_package_installation_subdirectory", &G_win32_get_package_installation_subdirectory},
	{"g_win32_get_package_installation_subdirectory_utf8", &G_win32_get_package_installation_subdirectory},
	{"g_win32_get_system_data_dirs_for_module", &G_win32_get_system_data_dirs_for_module},
	{"g_win32_get_windows_version", &G_win32_get_windows_version},
	{"g_win32_getlocale", &G_win32_getlocale},
	{"g_win32_locale_filename_from_utf8", &G_win32_locale_filename_from_utf8},
	{"glib_check_version", &Glib_check_version},
	// Undocumented {"glib_gettext", &glib_gettext},
	// Undocumented {"glib_on_error_halt", &glib_on_error_halt},
	// Undocumented {"glib_pgettext", &glib_pgettext},
}

var dllThread = "libgthread-2.0-0.dll"

var apiListThread = Apis{
	{"g_thread_init", &G_thread_init},
	{"g_thread_init_with_errorcheck_mutexes", &G_thread_init_with_errorcheck_mutexes},
}

var dataList = Data{
// {"g_ascii_table", new(G_ascii_table)},
// {"g_child_watch_funcs", new(G_child_watch_funcs)},
// {"g_idle_funcs", new(G_idle_funcs)},
// {"g_io_watch_funcs", new(G_io_watch_funcs)},
// {"g_mem_gc_friendly", new(G_mem_gc_friendly)},
// {"g_test_config_vars", new(G_test_config_vars)},
// {"g_thread_functions_for_glib_use", new(G_thread_functions_for_glib_use)},
// {"g_thread_gettime", new(G_thread_gettime)},
// {"g_thread_use_default_impl", new(G_thread_use_default_impl)},
// {"g_threads_got_initialized", new(G_threads_got_initialized)},
// {"g_timeout_funcs", new(G_timeout_funcs)},
// {"g_utf8_skip", new(G_utf8_skip)},
// {"glib_binary_age", new(glib_binary_age)},
// {"glib_interface_age", new(glib_interface_age)},
// {"glib_major_version", new(glib_major_version)},
// {"glib_mem_profiler_table", new(glib_mem_profiler_table)},
// {"glib_micro_version", new(glib_micro_version)},
// {"glib_minor_version", new(glib_minor_version)},
)}
