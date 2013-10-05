package expat

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
}

type (
	//XML_Parser *struct{}
	XML_Parser struct{}
	XML_Char   Char
	XML_LChar  Char
	XML_Bool   Unsigned_char
	XML_Size   Unsigned_long
	XML_Index  Long

	XML_ElementDeclHandler func(
		userData *Void, name string, model *XML_Content)

	XML_AttlistDeclHandler func(userData *Void,
		elname, attname, att_type, dflt string, isrequired int)

	XML_XmlDeclHandler func(userData *Void,
		version, encoding string, standalone int)

	XML_StartCdataSectionHandler func(userData *Void)

	XML_EndCdataSectionHandler func(userData *Void)

	XML_DefaultHandler func(
		userData *Void, s *XML_Char, leng int)

	XML_CharacterDataHandler func(
		userData *Void, s *XML_Char, leng int)

	XML_ProcessingInstructionHandler func(
		userData *Void, target, data string)

	XML_CommentHandler func(
		userData *Void, data string)

	XML_StartDoctypeDeclHandler func(
		userData *Void,
		doctypeName, sysid, pubid string,
		has_internal_subset int)

	XML_StartElementHandler func(
		userData *Void, name string, atts **XML_Char)

	XML_EndElementHandler func(
		userData *Void, name string)

	XML_EndDoctypeDeclHandler func(userData *Void)

	XML_EntityDeclHandler func(
		userData *Void,
		entityName string,
		is_parameter_entity int,
		value string,
		value_length int,
		base, systemId, publicId, notationName string)

	XML_EndNamespaceDeclHandler func(
		userData *Void, prefix string)

	XML_StartNamespaceDeclHandler func(
		userData *Void, prefix, uri string)

	XML_NotStandaloneHandler func(userData *Void) int

	XML_NotationDeclHandler func(
		userData *Void,
		notationName, base, systemId, publicId string)

	XML_ExternalEntityRefHandler func(parser XML_Parser,
		context, base, systemId, publicId string) int

	XML_SkippedEntityHandler func(
		userData *Void, entityName string, is_parameter_entity int)

	XML_UnparsedEntityDeclHandler func(userData *Void,
		entityName, base, systemId, publicId, notationName string)

	XML_UnknownEncodingHandler func(encodingHandlerData *Void,
		name string, info *XML_Encoding) int

	XML_Memory_Handling_Suite struct {
		malloc_fcn  func(size Size_t) *Void
		realloc_fcn func(ptr *Void, size Size_t) *Void
		free_fcn    func(ptr *Void)
	}

	XML_Content struct {
		Type        XML_Content_Type
		quant       XML_Content_Quant
		name        *XML_Char
		numchildren Unsigned_int
		children    *XML_Content
	}

	XML_ParsingStatus struct {
		parsing     XML_Parsing
		finalBuffer XML_Bool
	}

	XML_Encoding struct {
		Map     [256]int
		data    *Void
		convert func(data *Void, s *Char) int
		release func(data *Void)
	}

	XML_Expat_Version struct {
		major, minor, micro int
	}

	XML_Feature struct {
		feature XML_FeatureEnum
		name    POVString
		value   Long
	}
)

type XML_Status Enum

const (
	XML_STATUS_ERROR XML_Status = iota
	XML_STATUS_OK
	XML_STATUS_SUSPENDED
)

type XML_Content_Type Enum

const (
	XML_CTYPE_EMPTY XML_Content_Type = iota + 1
	XML_CTYPE_ANY
	XML_CTYPE_MIXED
	XML_CTYPE_NAME
	XML_CTYPE_CHOICE
	XML_CTYPE_SEQ
)

type XML_Content_Quant Enum

const (
	XML_CQUANT_NONE XML_Content_Quant = iota
	XML_CQUANT_OPT
	XML_CQUANT_REP
	XML_CQUANT_PLUS
)

type XML_Parsing Enum

const (
	XML_INITIALIZED XML_Parsing = iota
	XML_PARSING
	XML_FINISHED
	XML_SUSPENDED
)

type XML_Error Enum

const (
	XML_ERROR_NONE XML_Error = iota
	XML_ERROR_NO_MEMORY
	XML_ERROR_SYNTAX
	XML_ERROR_NO_ELEMENTS
	XML_ERROR_INVALID_TOKEN
	XML_ERROR_UNCLOSED_TOKEN
	XML_ERROR_PARTIAL_CHAR
	XML_ERROR_TAG_MISMATCH
	XML_ERROR_DUPLICATE_ATTRIBUTE
	XML_ERROR_JUNK_AFTER_DOC_ELEMENT
	XML_ERROR_PARAM_ENTITY_REF
	XML_ERROR_UNDEFINED_ENTITY
	XML_ERROR_RECURSIVE_ENTITY_REF
	XML_ERROR_ASYNC_ENTITY
	XML_ERROR_BAD_CHAR_REF
	XML_ERROR_BINARY_ENTITY_REF
	XML_ERROR_ATTRIBUTE_EXTERNAL_ENTITY_REF
	XML_ERROR_MISPLACED_XML_PI
	XML_ERROR_UNKNOWN_ENCODING
	XML_ERROR_INCORRECT_ENCODING
	XML_ERROR_UNCLOSED_CDATA_SECTION
	XML_ERROR_EXTERNAL_ENTITY_HANDLING
	XML_ERROR_NOT_STANDALONE
	XML_ERROR_UNEXPECTED_STATE
	XML_ERROR_ENTITY_DECLARED_IN_PE
	XML_ERROR_FEATURE_REQUIRES_XML_DTD
	XML_ERROR_CANT_CHANGE_FEATURE_ONCE_PARSING
	XML_ERROR_UNBOUND_PREFIX
	XML_ERROR_UNDECLARING_PREFIX
	XML_ERROR_INCOMPLETE_PE
	XML_ERROR_XML_DECL
	XML_ERROR_TEXT_DECL
	XML_ERROR_PUBLICID
	XML_ERROR_SUSPENDED
	XML_ERROR_NOT_SUSPENDED
	XML_ERROR_ABORTED
	XML_ERROR_FINISHED
	XML_ERROR_SUSPEND_PE
	XML_ERROR_RESERVED_PREFIX_XML
	XML_ERROR_RESERVED_PREFIX_XMLNS
	XML_ERROR_RESERVED_NAMESPACE_URI
)

type XML_ParamEntityParsing Enum

const (
	XML_PARAM_ENTITY_PARSING_NEVER XML_ParamEntityParsing = iota
	XML_PARAM_ENTITY_PARSING_UNLESS_STANDALONE
	XML_PARAM_ENTITY_PARSING_ALWAYS
)

type XML_FeatureEnum Enum

const (
	XML_FEATURE_END XML_FeatureEnum = iota
	XML_FEATURE_UNICODE
	XML_FEATURE_UNICODE_WCHAR_T
	XML_FEATURE_DTD
	XML_FEATURE_CONTEXT_BYTES
	XML_FEATURE_MIN_SIZE
	XML_FEATURE_SIZEOF_XML_CHAR
	XML_FEATURE_SIZEOF_XML_LCHAR
	XML_FEATURE_NS
	XML_FEATURE_LARGE_SIZE
)

var (
	XML_SetElementDeclHandler func(
		parser XML_Parser, eldecl XML_ElementDeclHandler)

	XML_SetAttlistDeclHandler func(
		parser XML_Parser, attdecl XML_AttlistDeclHandler)

	XML_SetXmlDeclHandler func(
		parser XML_Parser, xmldecl XML_XmlDeclHandler)

	XML_ParserCreate func(encoding string) *XML_Parser

	XML_ParserCreateNS func(
		encoding *XML_Char,
		namespaceSeparator XML_Char) XML_Parser

	XML_ParserCreate_MM func(
		encoding *XML_Char,
		memsuite *XML_Memory_Handling_Suite,
		namespaceSeparator *XML_Char) XML_Parser

	XML_ParserReset func(
		parser XML_Parser, encoding *XML_Char) XML_Bool

	XML_SetEntityDeclHandler func(
		parser XML_Parser,
		handler XML_EntityDeclHandler)

	XML_SetElementHandler func(
		parser XML_Parser,
		start XML_StartElementHandler,
		end XML_EndElementHandler)

	XML_SetStartElementHandler func(
		parser XML_Parser,
		handler XML_StartElementHandler)

	XML_SetEndElementHandler func(
		parser XML_Parser,
		handler XML_EndElementHandler)

	XML_SetCharacterDataHandler func(
		parser XML_Parser,
		handler XML_CharacterDataHandler)

	XML_SetProcessingInstructionHandler func(
		parser XML_Parser,
		handler XML_ProcessingInstructionHandler)

	XML_SetCommentHandler func(
		parser XML_Parser,
		handler XML_CommentHandler)

	XML_SetCdataSectionHandler func(
		parser XML_Parser,
		start XML_StartCdataSectionHandler,
		end XML_EndCdataSectionHandler)

	XML_SetStartCdataSectionHandler func(
		parser XML_Parser,
		start XML_StartCdataSectionHandler)

	XML_SetEndCdataSectionHandler func(
		parser XML_Parser,
		end XML_EndCdataSectionHandler)

	XML_SetDefaultHandler func(
		parser XML_Parser,
		handler XML_DefaultHandler)

	XML_SetDefaultHandlerExpand func(
		parser XML_Parser,
		handler XML_DefaultHandler)

	XML_SetDoctypeDeclHandler func(
		parser XML_Parser,
		start XML_StartDoctypeDeclHandler,
		end XML_EndDoctypeDeclHandler)

	XML_SetStartDoctypeDeclHandler func(
		parser XML_Parser,
		start XML_StartDoctypeDeclHandler)

	XML_SetEndDoctypeDeclHandler func(
		parser XML_Parser,
		end XML_EndDoctypeDeclHandler)

	XML_SetUnparsedEntityDeclHandler func(
		parser XML_Parser,
		handler XML_UnparsedEntityDeclHandler)

	XML_SetNotationDeclHandler func(
		parser XML_Parser,
		handler XML_NotationDeclHandler)

	XML_SetNamespaceDeclHandler func(
		parser XML_Parser,
		start XML_StartNamespaceDeclHandler,
		end XML_EndNamespaceDeclHandler)

	XML_SetStartNamespaceDeclHandler func(
		parser XML_Parser,
		start XML_StartNamespaceDeclHandler)

	XML_SetEndNamespaceDeclHandler func(
		parser XML_Parser,
		end XML_EndNamespaceDeclHandler)

	XML_SetNotStandaloneHandler func(
		parser XML_Parser,
		handler XML_NotStandaloneHandler)

	XML_SetExternalEntityRefHandler func(
		parser XML_Parser,
		handler XML_ExternalEntityRefHandler)

	XML_SetExternalEntityRefHandlerArg func(
		parser XML_Parser, arg *Void)

	XML_SetSkippedEntityHandler func(
		parser XML_Parser,
		handler XML_SkippedEntityHandler)

	XML_SetUnknownEncodingHandler func(
		parser XML_Parser,
		handler XML_UnknownEncodingHandler,
		encodingHandlerData *Void)

	XML_DefaultCurrent func(
		parser XML_Parser)

	XML_SetReturnNSTriplet func(parser XML_Parser, do_nst int)

	XML_SetUserData func(parser XML_Parser, userData *Void)

	XML_SetEncoding func(
		parser XML_Parser, encoding string) XML_Status

	XML_UseParserAsHandlerArg func(
		parser XML_Parser)

	XML_UseForeignDTD func(
		parser XML_Parser, useDTD XML_Bool) XML_Error

	XML_SetBase func(
		parser XML_Parser, base string) XML_Status

	XML_GetBase func(parser XML_Parser) string

	XML_GetSpecifiedAttributeCount func(
		parser XML_Parser) int

	XML_GetIdAttributeIndex func(
		parser XML_Parser) int

	XML_Parse func(parser XML_Parser,
		s *Char, leng, isFinal int) XML_Status

	XML_GetBuffer func(parser XML_Parser, leng int) *Void

	XML_ParseBuffer func(
		parser XML_Parser, leng, isFinal int) XML_Status

	XML_StopParser func(parser XML_Parser,
		resumable XML_Bool) XML_Status

	XML_ResumeParser func(parser XML_Parser) XML_Status

	XML_GetParsingStatus func(
		parser XML_Parser, status *XML_ParsingStatus)

	XML_ExternalEntityParserCreate func(
		parser XML_Parser,
		context, encoding *XML_Char) XML_Parser

	XML_SetParamEntityParsing func(
		parser XML_Parser, parsing XML_ParamEntityParsing) int

	XML_GetErrorCode func(parser XML_Parser) XML_Error

	XML_GetCurrentLineNumber func(parser XML_Parser) XML_Size

	XML_GetCurrentColumnNumber func(parser XML_Parser) XML_Size

	XML_GetCurrentByteIndex func(parser XML_Parser) XML_Index

	XML_GetCurrentByteCount func(parser XML_Parser) int

	XML_GetInputContext func(
		parser XML_Parser, offset, size *int) string

	XML_FreeContentModel func(
		parser XML_Parser, model *XML_Content)

	XML_MemMalloc func(parser XML_Parser, size Size_t) *Void

	XML_MemRealloc func(
		parser XML_Parser, ptr *Void, size Size_t) *Void

	XML_MemFree func(parser XML_Parser, ptr *Void)

	XML_ParserFree func(parser XML_Parser)

	XML_ErrorString func(code XML_Error) string

	XML_ExpatVersion func() string

	XML_ExpatVersionInfo func() XML_Expat_Version // NOTE: CRASHES

	XML_GetFeatureList func() *XML_Feature
)

var dll = "libexpat-1.dll"

var apiList = Apis{
	{"XML_DefaultCurrent", &XML_DefaultCurrent},
	{"XML_ErrorString", &XML_ErrorString},
	{"XML_ExpatVersion", &XML_ExpatVersion},
	{"XML_ExpatVersionInfo", &XML_ExpatVersionInfo},
	{"XML_ExternalEntityParserCreate", &XML_ExternalEntityParserCreate},
	{"XML_FreeContentModel", &XML_FreeContentModel},
	{"XML_GetBase", &XML_GetBase},
	{"XML_GetBuffer", &XML_GetBuffer},
	{"XML_GetCurrentByteCount", &XML_GetCurrentByteCount},
	{"XML_GetCurrentByteIndex", &XML_GetCurrentByteIndex},
	{"XML_GetCurrentColumnNumber", &XML_GetCurrentColumnNumber},
	{"XML_GetCurrentLineNumber", &XML_GetCurrentLineNumber},
	{"XML_GetErrorCode", &XML_GetErrorCode},
	{"XML_GetFeatureList", &XML_GetFeatureList},
	{"XML_GetIdAttributeIndex", &XML_GetIdAttributeIndex},
	{"XML_GetInputContext", &XML_GetInputContext},
	{"XML_GetParsingStatus", &XML_GetParsingStatus},
	{"XML_GetSpecifiedAttributeCount", &XML_GetSpecifiedAttributeCount},
	{"XML_MemFree", &XML_MemFree},
	{"XML_MemMalloc", &XML_MemMalloc},
	{"XML_MemRealloc", &XML_MemRealloc},
	{"XML_Parse", &XML_Parse},
	{"XML_ParseBuffer", &XML_ParseBuffer},
	{"XML_ParserCreate", &XML_ParserCreate},
	{"XML_ParserCreateNS", &XML_ParserCreateNS},
	{"XML_ParserCreate_MM", &XML_ParserCreate_MM},
	{"XML_ParserFree", &XML_ParserFree},
	{"XML_ParserReset", &XML_ParserReset},
	{"XML_ResumeParser", &XML_ResumeParser},
	{"XML_SetAttlistDeclHandler", &XML_SetAttlistDeclHandler},
	{"XML_SetBase", &XML_SetBase},
	{"XML_SetCdataSectionHandler", &XML_SetCdataSectionHandler},
	{"XML_SetCharacterDataHandler", &XML_SetCharacterDataHandler},
	{"XML_SetCommentHandler", &XML_SetCommentHandler},
	{"XML_SetDefaultHandler", &XML_SetDefaultHandler},
	{"XML_SetDefaultHandlerExpand", &XML_SetDefaultHandlerExpand},
	{"XML_SetDoctypeDeclHandler", &XML_SetDoctypeDeclHandler},
	{"XML_SetElementDeclHandler", &XML_SetElementDeclHandler},
	{"XML_SetElementHandler", &XML_SetElementHandler},
	{"XML_SetEncoding", &XML_SetEncoding},
	{"XML_SetEndCdataSectionHandler", &XML_SetEndCdataSectionHandler},
	{"XML_SetEndDoctypeDeclHandler", &XML_SetEndDoctypeDeclHandler},
	{"XML_SetEndElementHandler", &XML_SetEndElementHandler},
	{"XML_SetEndNamespaceDeclHandler", &XML_SetEndNamespaceDeclHandler},
	{"XML_SetEntityDeclHandler", &XML_SetEntityDeclHandler},
	{"XML_SetExternalEntityRefHandler", &XML_SetExternalEntityRefHandler},
	{"XML_SetExternalEntityRefHandlerArg", &XML_SetExternalEntityRefHandlerArg},
	{"XML_SetNamespaceDeclHandler", &XML_SetNamespaceDeclHandler},
	{"XML_SetNotStandaloneHandler", &XML_SetNotStandaloneHandler},
	{"XML_SetNotationDeclHandler", &XML_SetNotationDeclHandler},
	{"XML_SetParamEntityParsing", &XML_SetParamEntityParsing},
	{"XML_SetProcessingInstructionHandler", &XML_SetProcessingInstructionHandler},
	{"XML_SetReturnNSTriplet", &XML_SetReturnNSTriplet},
	{"XML_SetSkippedEntityHandler", &XML_SetSkippedEntityHandler},
	{"XML_SetStartCdataSectionHandler", &XML_SetStartCdataSectionHandler},
	{"XML_SetStartDoctypeDeclHandler", &XML_SetStartDoctypeDeclHandler},
	{"XML_SetStartElementHandler", &XML_SetStartElementHandler},
	{"XML_SetStartNamespaceDeclHandler", &XML_SetStartNamespaceDeclHandler},
	{"XML_SetUnknownEncodingHandler", &XML_SetUnknownEncodingHandler},
	{"XML_SetUnparsedEntityDeclHandler", &XML_SetUnparsedEntityDeclHandler},
	{"XML_SetUserData", &XML_SetUserData},
	{"XML_SetXmlDeclHandler", &XML_SetXmlDeclHandler},
	{"XML_StopParser", &XML_StopParser},
	{"XML_UseForeignDTD", &XML_UseForeignDTD},
	{"XML_UseParserAsHandlerArg", &XML_UseParserAsHandlerArg},
}
