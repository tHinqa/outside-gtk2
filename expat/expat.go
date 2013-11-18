// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package expat provides API definitions for accessing
//libexpat-1.dll.
package expat

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type (
	Enum int

	//Parser *struct{}
	Parser struct{}
	Char   T.Char
	LChar  T.Char
	Bool   T.UnsignedChar
	Size   T.UnsignedLong
	Index  T.Long

	ElementDeclHandler func(
		userData *T.Void, name string, model *Content)

	AttlistDeclHandler func(userData *T.Void,
		elName, attName, attType, dflt string, isRequired int)

	XmlDeclHandler func(userData *T.Void,
		version, encoding string, standalone int)

	StartCdataSectionHandler func(userData *T.Void)

	EndCdataSectionHandler func(userData *T.Void)

	DefaultHandler func(
		userData *T.Void, s *Char, leng int)

	CharacterDataHandler func(
		userData *T.Void, s *Char, leng int)

	ProcessingInstructionHandler func(
		userData *T.Void, target, data string)

	CommentHandler func(
		userData *T.Void, data string)

	StartDoctypeDeclHandler func(
		userData *T.Void,
		doctypeName, sysid, pubid string,
		hasInternalSubset int)

	StartElementHandler func(
		userData *T.Void, name string, atts **Char)

	EndElementHandler func(
		userData *T.Void, name string)

	EndDoctypeDeclHandler func(userData *T.Void)

	EntityDeclHandler func(
		userData *T.Void,
		entityName string,
		isParameterEntity int,
		value string,
		valueLength int,
		base, systemId, publicId, notationName string)

	EndNamespaceDeclHandler func(
		userData *T.Void, prefix string)

	StartNamespaceDeclHandler func(
		userData *T.Void, prefix, uri string)

	NotStandaloneHandler func(userData *T.Void) int

	NotationDeclHandler func(
		userData *T.Void,
		notationName, base, systemId, publicId string)

	ExternalEntityRefHandler func(parser Parser,
		context, base, systemId, publicId string) int

	SkippedEntityHandler func(
		userData *T.Void, entityName string, isParameterEntity int)

	UnparsedEntityDeclHandler func(userData *T.Void,
		entityName, base, systemId, publicId, notationName string)

	UnknownEncodingHandler func(encodingHandlerData *T.Void,
		name string, info *Encoding) int

	MemoryHandlingSuite struct {
		mallocFunc  func(size T.Size_t) *T.Void
		reallocFunc func(ptr *T.Void, size T.Size_t) *T.Void
		freeFunc    func(ptr *T.Void)
	}

	Content struct {
		Type        ContentType
		Quant       ContentQuant
		Name        *Char
		Numchildren uint
		Children    *Content
	}

	ParsingStatus struct {
		Parsing     Parsing
		FinalBuffer Bool
	}

	Encoding struct {
		Map     [256]int
		Data    *T.Void
		Convert func(data *T.Void, s *T.Char) int
		Release func(data *T.Void)
	}

	ExpatVersionS struct { // NOTE(t): Name conflict
		Major, Minor, Micro int
	}

	Feature struct {
		Feature FeatureEnum
		Name    *OVString
		Value   T.Long
	}
)

type Status Enum

const (
	STATUS_ERROR Status = iota
	STATUS_OK
	STATUS_SUSPENDED
)

type ContentType Enum

const (
	CTYPE_EMPTY ContentType = iota + 1
	CTYPE_ANY
	CTYPE_MIXED
	CTYPE_NAME
	CTYPE_CHOICE
	CTYPE_SEQ
)

type ContentQuant Enum

const (
	CQUANT_NONE ContentQuant = iota
	CQUANT_OPT
	CQUANT_REP
	CQUANT_PLUS
)

type Parsing Enum

const (
	INITIALIZED Parsing = iota
	PARSING
	FINISHED
	SUSPENDED
)

type Error Enum

const (
	ERROR_NONE Error = iota
	ERROR_NO_MEMORY
	ERROR_SYNTAX
	ERROR_NO_ELEMENTS
	ERROR_INVALID_TOKEN
	ERROR_UNCLOSED_TOKEN
	ERROR_PARTIAL_CHAR
	ERROR_TAG_MISMATCH
	ERROR_DUPLICATE_ATTRIBUTE
	ERROR_JUNK_AFTER_DOC_ELEMENT
	ERROR_PARAM_ENTITY_REF
	ERROR_UNDEFINED_ENTITY
	ERROR_RECURSIVE_ENTITY_REF
	ERROR_ASYNC_ENTITY
	ERROR_BAD_CHAR_REF
	ERROR_BINARY_ENTITY_REF
	ERROR_ATTRIBUTE_EXTERNAL_ENTITY_REF
	ERROR_MISPLACED_PI
	ERROR_UNKNOWN_ENCODING
	ERROR_INCORRECT_ENCODING
	ERROR_UNCLOSED_CDATA_SECTION
	ERROR_EXTERNAL_ENTITY_HANDLING
	ERROR_NOT_STANDALONE
	ERROR_UNEXPECTED_STATE
	ERROR_ENTITY_DECLARED_IN_PE
	ERROR_FEATURE_REQUIRES_DTD
	ERROR_CANT_CHANGE_FEATURE_ONCE_PARSING
	ERROR_UNBOUND_PREFIX
	ERROR_UNDECLARING_PREFIX
	ERROR_INCOMPLETE_PE
	ERROR_DECL
	ERROR_TEXT_DECL
	ERROR_PUBLICID
	ERROR_SUSPENDED
	ERROR_NOT_SUSPENDED
	ERROR_ABORTED
	ERROR_FINISHED
	ERROR_SUSPEND_PE
	ERROR_RESERVED_PREFIX_XML
	ERROR_RESERVED_PREFIX_XMLNS
	ERROR_RESERVED_NAMESPACE_URI
)

type ParamEntityParsing Enum

const (
	PARAM_ENTITY_PARSING_NEVER ParamEntityParsing = iota
	PARAM_ENTITY_PARSING_UNLESS_STANDALONE
	PARAM_ENTITY_PARSING_ALWAYS
)

type FeatureEnum Enum

const (
	FEATURE_END FeatureEnum = iota
	FEATURE_UNICODE
	FEATURE_UNICODE_WCHAR_T
	FEATURE_DTD
	FEATURE_CONTEXT_BYTES
	FEATURE_MIN_SIZE
	FEATURE_SIZEOF_CHAR
	FEATURE_SIZEOF_LCHAR
	FEATURE_NS
	FEATURE_LARGE_SIZE
)

var (
	ErrorString func(code Error) string

	ExpatVersion func() string

	// NOTE: Crashes
	ExpatVersionInfo func() ExpatVersionS

	GetFeatureList func() *[99]Feature //TODO(t):Fix
)
var (
	ParserCreate   func(encoding string) *Parser
	ParserCreateNS func(encoding *Char, namespaceSeparator Char) Parser
	ParserCreateMM func(encoding *Char, memsuite *MemoryHandlingSuite, namespaceSeparator *Char) Parser

	DefaultCurrent                  func(p Parser)
	ExternalEntityParserCreate      func(p Parser, context, encoding *Char) Parser
	FreeContentModel                func(p Parser, model *Content)
	GetBase                         func(p Parser) string
	GetBuffer                       func(p Parser, leng int) *T.Void
	GetCurrentByteCount             func(p Parser) int
	GetCurrentByteIndex             func(p Parser) Index
	GetCurrentColumnNumber          func(p Parser) Size
	GetCurrentLineNumber            func(p Parser) Size
	GetErrorCode                    func(p Parser) Error
	GetIdAttributeIndex             func(p Parser) int
	GetInputContext                 func(p Parser, offset, size *int) string
	GetParsingStatus                func(p Parser, status *ParsingStatus)
	GetSpecifiedAttributeCount      func(p Parser) int
	MemFree                         func(p Parser, ptr *T.Void)
	MemMalloc                       func(p Parser, size T.Size_t) *T.Void
	MemRealloc                      func(p Parser, ptr *T.Void, size T.Size_t) *T.Void
	Parse                           func(p Parser, s *T.Char, leng, isFinal int) Status
	ParseBuffer                     func(p Parser, leng, isFinal int) Status
	ParserFree                      func(p Parser)
	ParserReset                     func(p Parser, encoding *Char) Bool
	ResumeParser                    func(p Parser) Status
	SetAttlistDeclHandler           func(p Parser, attdecl AttlistDeclHandler)
	SetBase                         func(p Parser, base string) Status
	SetCdataSectionHandler          func(p Parser, start StartCdataSectionHandler, end EndCdataSectionHandler)
	SetCharacterDataHandler         func(p Parser, handler CharacterDataHandler)
	SetCommentHandler               func(p Parser, handler CommentHandler)
	SetDefaultHandler               func(p Parser, handler DefaultHandler)
	SetDefaultHandlerExpand         func(p Parser, handler DefaultHandler)
	SetDoctypeDeclHandler           func(p Parser, start StartDoctypeDeclHandler, end EndDoctypeDeclHandler)
	SetElementDeclHandler           func(p Parser, eldecl ElementDeclHandler)
	SetElementHandler               func(p Parser, start StartElementHandler, end EndElementHandler)
	SetEncoding                     func(p Parser, encoding string) Status
	SetEndCdataSectionHandler       func(p Parser, end EndCdataSectionHandler)
	SetEndDoctypeDeclHandler        func(p Parser, end EndDoctypeDeclHandler)
	SetEndElementHandler            func(p Parser, handler EndElementHandler)
	SetEndNamespaceDeclHandler      func(p Parser, end EndNamespaceDeclHandler)
	SetEntityDeclHandler            func(p Parser, handler EntityDeclHandler)
	SetExternalEntityRefHandler     func(p Parser, handler ExternalEntityRefHandler)
	SetExternalEntityRefHandlerArg  func(p Parser, arg *T.Void)
	SetNamespaceDeclHandler         func(p Parser, start StartNamespaceDeclHandler, end EndNamespaceDeclHandler)
	SetNotationDeclHandler          func(p Parser, handler NotationDeclHandler)
	SetNotStandaloneHandler         func(p Parser, handler NotStandaloneHandler)
	SetParamEntityParsing           func(p Parser, parsing ParamEntityParsing) int
	SetProcessingInstructionHandler func(p Parser, handler ProcessingInstructionHandler)
	SetReturnNSTriplet              func(p Parser, do_nst int)
	SetSkippedEntityHandler         func(p Parser, handler SkippedEntityHandler)
	SetStartCdataSectionHandler     func(p Parser, start StartCdataSectionHandler)
	SetStartDoctypeDeclHandler      func(p Parser, start StartDoctypeDeclHandler)
	SetStartElementHandler          func(p Parser, handler StartElementHandler)
	SetStartNamespaceDeclHandler    func(p Parser, start StartNamespaceDeclHandler)
	SetUnknownEncodingHandler       func(p Parser, handler UnknownEncodingHandler, encodingHandlerData *T.Void)
	SetUnparsedEntityDeclHandler    func(p Parser, handler UnparsedEntityDeclHandler)
	SetUserData                     func(p Parser, userData *T.Void)
	SetXmlDeclHandler               func(p Parser, xmldecl XmlDeclHandler)
	StopParser                      func(p Parser, resumable Bool) Status
	UseForeignDTD                   func(p Parser, useDTD Bool) Error
	UseParserAsHandlerArg           func(p Parser)
)

func (p Parser) DefaultCurrent() { DefaultCurrent(p) }
func (p Parser) ExternalEntityParserCreate(context, encoding *Char) Parser {
	return ExternalEntityParserCreate(p, context, encoding)
}
func (p Parser) FreeContentModel(model *Content)                  { FreeContentModel(p, model) }
func (p Parser) Base() string                                     { return GetBase(p) }
func (p Parser) Buffer(leng int) *T.Void                          { return GetBuffer(p, leng) }
func (p Parser) CurrentByteCount() int                            { return GetCurrentByteCount(p) }
func (p Parser) CurrentByteIndex() Index                          { return GetCurrentByteIndex(p) }
func (p Parser) CurrentColumnNumber() Size                        { return GetCurrentColumnNumber(p) }
func (p Parser) CurrentLineNumber() Size                          { return GetCurrentLineNumber(p) }
func (p Parser) ErrorCode() Error                                 { return GetErrorCode(p) }
func (p Parser) IdAttributeIndex() int                            { return GetIdAttributeIndex(p) }
func (p Parser) InputContext(offset, size *int) string            { return GetInputContext(p, offset, size) }
func (p Parser) MemFree(ptr *T.Void)                              { MemFree(p, ptr) }
func (p Parser) MemMalloc(size T.Size_t) *T.Void                  { return MemMalloc(p, size) }
func (p Parser) MemRealloc(ptr *T.Void, size T.Size_t) *T.Void    { return MemRealloc(p, ptr, size) }
func (p Parser) Parse(s *T.Char, leng, isFinal int) Status        { return Parse(p, s, leng, isFinal) }
func (p Parser) ParseBuffer(leng, isFinal int) Status             { return ParseBuffer(p, leng, isFinal) }
func (p Parser) ParserFree()                                      { ParserFree(p) }
func (p Parser) ParserReset(encoding *Char) Bool                  { return ParserReset(p, encoding) }
func (p Parser) ParsingStatus(status *ParsingStatus)              { GetParsingStatus(p, status) }
func (p Parser) ResumeParser() Status                             { return ResumeParser(p) }
func (p Parser) SetAttlistDeclHandler(attdecl AttlistDeclHandler) { SetAttlistDeclHandler(p, attdecl) }
func (p Parser) SetBase(base string) Status                       { return SetBase(p, base) }
func (p Parser) SetCdataSectionHandler(start StartCdataSectionHandler, end EndCdataSectionHandler) {
	SetCdataSectionHandler(p, start, end)
}
func (p Parser) SetCharacterDataHandler(handler CharacterDataHandler) {
	SetCharacterDataHandler(p, handler)
}
func (p Parser) SetCommentHandler(handler CommentHandler)       { SetCommentHandler(p, handler) }
func (p Parser) SetDefaultHandler(handler DefaultHandler)       { SetDefaultHandler(p, handler) }
func (p Parser) SetDefaultHandlerExpand(handler DefaultHandler) { SetDefaultHandlerExpand(p, handler) }
func (p Parser) SetDoctypeDeclHandler(start StartDoctypeDeclHandler, end EndDoctypeDeclHandler) {
	SetDoctypeDeclHandler(p, start, end)
}
func (p Parser) SetElementDeclHandler(eldecl ElementDeclHandler) { SetElementDeclHandler(p, eldecl) }
func (p Parser) SetElementHandler(start StartElementHandler, end EndElementHandler) {
	SetElementHandler(p, start, end)
}
func (p Parser) SetEncoding(encoding string) Status { return SetEncoding(p, encoding) }
func (p Parser) SetEndCdataSectionHandler(end EndCdataSectionHandler) {
	SetEndCdataSectionHandler(p, end)
}
func (p Parser) SetEndDoctypeDeclHandler(end EndDoctypeDeclHandler) { SetEndDoctypeDeclHandler(p, end) }
func (p Parser) SetEndElementHandler(handler EndElementHandler)     { SetEndElementHandler(p, handler) }
func (p Parser) SetEndNamespaceDeclHandler(end EndNamespaceDeclHandler) {
	SetEndNamespaceDeclHandler(p, end)
}
func (p Parser) SetEntityDeclHandler(handler EntityDeclHandler) { SetEntityDeclHandler(p, handler) }
func (p Parser) SetExternalEntityRefHandler(handler ExternalEntityRefHandler) {
	SetExternalEntityRefHandler(p, handler)
}
func (p Parser) SetExternalEntityRefHandlerArg(arg *T.Void) { SetExternalEntityRefHandlerArg(p, arg) }
func (p Parser) SetNamespaceDeclHandler(start StartNamespaceDeclHandler, end EndNamespaceDeclHandler) {
	SetNamespaceDeclHandler(p, start, end)
}
func (p Parser) SetNotationDeclHandler(handler NotationDeclHandler) {
	SetNotationDeclHandler(p, handler)
}
func (p Parser) SetNotStandaloneHandler(handler NotStandaloneHandler) {
	SetNotStandaloneHandler(p, handler)
}
func (p Parser) SetParamEntityParsing(parsing ParamEntityParsing) int {
	return SetParamEntityParsing(p, parsing)
}
func (p Parser) SetProcessingInstructionHandler(handler ProcessingInstructionHandler) {
	SetProcessingInstructionHandler(p, handler)
}
func (p Parser) SetReturnNSTriplet(do_nst int) { SetReturnNSTriplet(p, do_nst) }
func (p Parser) SetSkippedEntityHandler(handler SkippedEntityHandler) {
	SetSkippedEntityHandler(p, handler)
}
func (p Parser) SetStartCdataSectionHandler(start StartCdataSectionHandler) {
	SetStartCdataSectionHandler(p, start)
}
func (p Parser) SetStartDoctypeDeclHandler(start StartDoctypeDeclHandler) {
	SetStartDoctypeDeclHandler(p, start)
}
func (p Parser) SetStartElementHandler(handler StartElementHandler) {
	SetStartElementHandler(p, handler)
}
func (p Parser) SetStartNamespaceDeclHandler(start StartNamespaceDeclHandler) {
	SetStartNamespaceDeclHandler(p, start)
}
func (p Parser) SetUnknownEncodingHandler(handler UnknownEncodingHandler, encodingHandlerData *T.Void) {
	SetUnknownEncodingHandler(p, handler, encodingHandlerData)
}
func (p Parser) SetUnparsedEntityDeclHandler(handler UnparsedEntityDeclHandler) {
	SetUnparsedEntityDeclHandler(p, handler)
}
func (p Parser) SetUserData(userData *T.Void)             { SetUserData(p, userData) }
func (p Parser) SetXmlDeclHandler(xmldecl XmlDeclHandler) { SetXmlDeclHandler(p, xmldecl) }
func (p Parser) SpecifiedAttributeCount() int             { return GetSpecifiedAttributeCount(p) }
func (p Parser) StopParser(resumable Bool) Status         { return StopParser(p, resumable) }
func (p Parser) UseForeignDTD(useDTD Bool) Error          { return UseForeignDTD(p, useDTD) }
func (p Parser) UseParserAsHandlerArg()                   { UseParserAsHandlerArg(p) }

var dll = "libexpat-1.dll"

var apiList = outside.Apis{
	{"XML_DefaultCurrent", &DefaultCurrent},
	{"XML_ErrorString", &ErrorString},
	{"XML_ExpatVersion", &ExpatVersion},
	{"XML_ExpatVersionInfo", &ExpatVersionInfo},
	{"XML_ExternalEntityParserCreate", &ExternalEntityParserCreate},
	{"XML_FreeContentModel", &FreeContentModel},
	{"XML_GetBase", &GetBase},
	{"XML_GetBuffer", &GetBuffer},
	{"XML_GetCurrentByteCount", &GetCurrentByteCount},
	{"XML_GetCurrentByteIndex", &GetCurrentByteIndex},
	{"XML_GetCurrentColumnNumber", &GetCurrentColumnNumber},
	{"XML_GetCurrentLineNumber", &GetCurrentLineNumber},
	{"XML_GetErrorCode", &GetErrorCode},
	{"XML_GetFeatureList", &GetFeatureList},
	{"XML_GetIdAttributeIndex", &GetIdAttributeIndex},
	{"XML_GetInputContext", &GetInputContext},
	{"XML_GetParsingStatus", &GetParsingStatus},
	{"XML_GetSpecifiedAttributeCount", &GetSpecifiedAttributeCount},
	{"XML_MemFree", &MemFree},
	{"XML_MemMalloc", &MemMalloc},
	{"XML_MemRealloc", &MemRealloc},
	{"XML_Parse", &Parse},
	{"XML_ParseBuffer", &ParseBuffer},
	{"XML_ParserCreate", &ParserCreate},
	{"XML_ParserCreateNS", &ParserCreateNS},
	{"XML_ParserCreate_MM", &ParserCreateMM},
	{"XML_ParserFree", &ParserFree},
	{"XML_ParserReset", &ParserReset},
	{"XML_ResumeParser", &ResumeParser},
	{"XML_SetAttlistDeclHandler", &SetAttlistDeclHandler},
	{"XML_SetBase", &SetBase},
	{"XML_SetCdataSectionHandler", &SetCdataSectionHandler},
	{"XML_SetCharacterDataHandler", &SetCharacterDataHandler},
	{"XML_SetCommentHandler", &SetCommentHandler},
	{"XML_SetDefaultHandler", &SetDefaultHandler},
	{"XML_SetDefaultHandlerExpand", &SetDefaultHandlerExpand},
	{"XML_SetDoctypeDeclHandler", &SetDoctypeDeclHandler},
	{"XML_SetElementDeclHandler", &SetElementDeclHandler},
	{"XML_SetElementHandler", &SetElementHandler},
	{"XML_SetEncoding", &SetEncoding},
	{"XML_SetEndCdataSectionHandler", &SetEndCdataSectionHandler},
	{"XML_SetEndDoctypeDeclHandler", &SetEndDoctypeDeclHandler},
	{"XML_SetEndElementHandler", &SetEndElementHandler},
	{"XML_SetEndNamespaceDeclHandler", &SetEndNamespaceDeclHandler},
	{"XML_SetEntityDeclHandler", &SetEntityDeclHandler},
	{"XML_SetExternalEntityRefHandler", &SetExternalEntityRefHandler},
	{"XML_SetExternalEntityRefHandlerArg", &SetExternalEntityRefHandlerArg},
	{"XML_SetNamespaceDeclHandler", &SetNamespaceDeclHandler},
	{"XML_SetNotStandaloneHandler", &SetNotStandaloneHandler},
	{"XML_SetNotationDeclHandler", &SetNotationDeclHandler},
	{"XML_SetParamEntityParsing", &SetParamEntityParsing},
	{"XML_SetProcessingInstructionHandler", &SetProcessingInstructionHandler},
	{"XML_SetReturnNSTriplet", &SetReturnNSTriplet},
	{"XML_SetSkippedEntityHandler", &SetSkippedEntityHandler},
	{"XML_SetStartCdataSectionHandler", &SetStartCdataSectionHandler},
	{"XML_SetStartDoctypeDeclHandler", &SetStartDoctypeDeclHandler},
	{"XML_SetStartElementHandler", &SetStartElementHandler},
	{"XML_SetStartNamespaceDeclHandler", &SetStartNamespaceDeclHandler},
	{"XML_SetUnknownEncodingHandler", &SetUnknownEncodingHandler},
	{"XML_SetUnparsedEntityDeclHandler", &SetUnparsedEntityDeclHandler},
	{"XML_SetUserData", &SetUserData},
	{"XML_SetXmlDeclHandler", &SetXmlDeclHandler},
	{"XML_StopParser", &StopParser},
	{"XML_UseForeignDTD", &UseForeignDTD},
	{"XML_UseParserAsHandlerArg", &UseParserAsHandlerArg},
}
