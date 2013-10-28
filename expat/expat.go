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
		Name    POVString
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

	defaultCurrent                  func(p Parser)
	externalEntityParserCreate      func(p Parser, context, encoding *Char) Parser
	freeContentModel                func(p Parser, model *Content)
	getBase                         func(p Parser) string
	getBuffer                       func(p Parser, leng int) *T.Void
	getCurrentByteCount             func(p Parser) int
	getCurrentByteIndex             func(p Parser) Index
	getCurrentColumnNumber          func(p Parser) Size
	getCurrentLineNumber            func(p Parser) Size
	getErrorCode                    func(p Parser) Error
	getIdAttributeIndex             func(p Parser) int
	getInputContext                 func(p Parser, offset, size *int) string
	getParsingStatus                func(p Parser, status *ParsingStatus)
	getSpecifiedAttributeCount      func(p Parser) int
	memFree                         func(p Parser, ptr *T.Void)
	memMalloc                       func(p Parser, size T.Size_t) *T.Void
	memRealloc                      func(p Parser, ptr *T.Void, size T.Size_t) *T.Void
	parse                           func(p Parser, s *T.Char, leng, isFinal int) Status
	parseBuffer                     func(p Parser, leng, isFinal int) Status
	parserFree                      func(p Parser)
	parserReset                     func(p Parser, encoding *Char) Bool
	resumeParser                    func(p Parser) Status
	setAttlistDeclHandler           func(p Parser, attdecl AttlistDeclHandler)
	setBase                         func(p Parser, base string) Status
	setCdataSectionHandler          func(p Parser, start StartCdataSectionHandler, end EndCdataSectionHandler)
	setCharacterDataHandler         func(p Parser, handler CharacterDataHandler)
	setCommentHandler               func(p Parser, handler CommentHandler)
	setDefaultHandler               func(p Parser, handler DefaultHandler)
	setDefaultHandlerExpand         func(p Parser, handler DefaultHandler)
	setDoctypeDeclHandler           func(p Parser, start StartDoctypeDeclHandler, end EndDoctypeDeclHandler)
	setElementDeclHandler           func(p Parser, eldecl ElementDeclHandler)
	setElementHandler               func(p Parser, start StartElementHandler, end EndElementHandler)
	setEncoding                     func(p Parser, encoding string) Status
	setEndCdataSectionHandler       func(p Parser, end EndCdataSectionHandler)
	setEndDoctypeDeclHandler        func(p Parser, end EndDoctypeDeclHandler)
	setEndElementHandler            func(p Parser, handler EndElementHandler)
	setEndNamespaceDeclHandler      func(p Parser, end EndNamespaceDeclHandler)
	setEntityDeclHandler            func(p Parser, handler EntityDeclHandler)
	setExternalEntityRefHandler     func(p Parser, handler ExternalEntityRefHandler)
	setExternalEntityRefHandlerArg  func(p Parser, arg *T.Void)
	setNamespaceDeclHandler         func(p Parser, start StartNamespaceDeclHandler, end EndNamespaceDeclHandler)
	setNotationDeclHandler          func(p Parser, handler NotationDeclHandler)
	setNotStandaloneHandler         func(p Parser, handler NotStandaloneHandler)
	setParamEntityParsing           func(p Parser, parsing ParamEntityParsing) int
	setProcessingInstructionHandler func(p Parser, handler ProcessingInstructionHandler)
	setReturnNSTriplet              func(p Parser, do_nst int)
	setSkippedEntityHandler         func(p Parser, handler SkippedEntityHandler)
	setStartCdataSectionHandler     func(p Parser, start StartCdataSectionHandler)
	setStartDoctypeDeclHandler      func(p Parser, start StartDoctypeDeclHandler)
	setStartElementHandler          func(p Parser, handler StartElementHandler)
	setStartNamespaceDeclHandler    func(p Parser, start StartNamespaceDeclHandler)
	setUnknownEncodingHandler       func(p Parser, handler UnknownEncodingHandler, encodingHandlerData *T.Void)
	setUnparsedEntityDeclHandler    func(p Parser, handler UnparsedEntityDeclHandler)
	setUserData                     func(p Parser, userData *T.Void)
	setXmlDeclHandler               func(p Parser, xmldecl XmlDeclHandler)
	stopParser                      func(p Parser, resumable Bool) Status
	useForeignDTD                   func(p Parser, useDTD Bool) Error
	useParserAsHandlerArg           func(p Parser)
)

func (p Parser) DefaultCurrent() { defaultCurrent(p) }
func (p Parser) ExternalEntityParserCreate(context, encoding *Char) Parser {
	return externalEntityParserCreate(p, context, encoding)
}
func (p Parser) FreeContentModel(model *Content)                  { freeContentModel(p, model) }
func (p Parser) Base() string                                     { return getBase(p) }
func (p Parser) Buffer(leng int) *T.Void                          { return getBuffer(p, leng) }
func (p Parser) CurrentByteCount() int                            { return getCurrentByteCount(p) }
func (p Parser) CurrentByteIndex() Index                          { return getCurrentByteIndex(p) }
func (p Parser) CurrentColumnNumber() Size                        { return getCurrentColumnNumber(p) }
func (p Parser) CurrentLineNumber() Size                          { return getCurrentLineNumber(p) }
func (p Parser) ErrorCode() Error                                 { return getErrorCode(p) }
func (p Parser) IdAttributeIndex() int                            { return getIdAttributeIndex(p) }
func (p Parser) InputContext(offset, size *int) string            { return getInputContext(p, offset, size) }
func (p Parser) MemFree(ptr *T.Void)                              { memFree(p, ptr) }
func (p Parser) MemMalloc(size T.Size_t) *T.Void                  { return memMalloc(p, size) }
func (p Parser) MemRealloc(ptr *T.Void, size T.Size_t) *T.Void    { return memRealloc(p, ptr, size) }
func (p Parser) Parse(s *T.Char, leng, isFinal int) Status        { return parse(p, s, leng, isFinal) }
func (p Parser) ParseBuffer(leng, isFinal int) Status             { return parseBuffer(p, leng, isFinal) }
func (p Parser) ParserFree()                                      { parserFree(p) }
func (p Parser) ParserReset(encoding *Char) Bool                  { return parserReset(p, encoding) }
func (p Parser) ParsingStatus(status *ParsingStatus)              { getParsingStatus(p, status) }
func (p Parser) ResumeParser() Status                             { return resumeParser(p) }
func (p Parser) SetAttlistDeclHandler(attdecl AttlistDeclHandler) { setAttlistDeclHandler(p, attdecl) }
func (p Parser) SetBase(base string) Status                       { return setBase(p, base) }
func (p Parser) SetCdataSectionHandler(start StartCdataSectionHandler, end EndCdataSectionHandler) {
	setCdataSectionHandler(p, start, end)
}
func (p Parser) SetCharacterDataHandler(handler CharacterDataHandler) {
	setCharacterDataHandler(p, handler)
}
func (p Parser) SetCommentHandler(handler CommentHandler)       { setCommentHandler(p, handler) }
func (p Parser) SetDefaultHandler(handler DefaultHandler)       { setDefaultHandler(p, handler) }
func (p Parser) SetDefaultHandlerExpand(handler DefaultHandler) { setDefaultHandlerExpand(p, handler) }
func (p Parser) SetDoctypeDeclHandler(start StartDoctypeDeclHandler, end EndDoctypeDeclHandler) {
	setDoctypeDeclHandler(p, start, end)
}
func (p Parser) SetElementDeclHandler(eldecl ElementDeclHandler) { setElementDeclHandler(p, eldecl) }
func (p Parser) SetElementHandler(start StartElementHandler, end EndElementHandler) {
	setElementHandler(p, start, end)
}
func (p Parser) SetEncoding(encoding string) Status { return setEncoding(p, encoding) }
func (p Parser) SetEndCdataSectionHandler(end EndCdataSectionHandler) {
	setEndCdataSectionHandler(p, end)
}
func (p Parser) SetEndDoctypeDeclHandler(end EndDoctypeDeclHandler) { setEndDoctypeDeclHandler(p, end) }
func (p Parser) SetEndElementHandler(handler EndElementHandler)     { setEndElementHandler(p, handler) }
func (p Parser) SetEndNamespaceDeclHandler(end EndNamespaceDeclHandler) {
	setEndNamespaceDeclHandler(p, end)
}
func (p Parser) SetEntityDeclHandler(handler EntityDeclHandler) { setEntityDeclHandler(p, handler) }
func (p Parser) SetExternalEntityRefHandler(handler ExternalEntityRefHandler) {
	setExternalEntityRefHandler(p, handler)
}
func (p Parser) SetExternalEntityRefHandlerArg(arg *T.Void) { setExternalEntityRefHandlerArg(p, arg) }
func (p Parser) SetNamespaceDeclHandler(start StartNamespaceDeclHandler, end EndNamespaceDeclHandler) {
	setNamespaceDeclHandler(p, start, end)
}
func (p Parser) SetNotationDeclHandler(handler NotationDeclHandler) {
	setNotationDeclHandler(p, handler)
}
func (p Parser) SetNotStandaloneHandler(handler NotStandaloneHandler) {
	setNotStandaloneHandler(p, handler)
}
func (p Parser) SetParamEntityParsing(parsing ParamEntityParsing) int {
	return setParamEntityParsing(p, parsing)
}
func (p Parser) SetProcessingInstructionHandler(handler ProcessingInstructionHandler) {
	setProcessingInstructionHandler(p, handler)
}
func (p Parser) SetReturnNSTriplet(do_nst int) { setReturnNSTriplet(p, do_nst) }
func (p Parser) SetSkippedEntityHandler(handler SkippedEntityHandler) {
	setSkippedEntityHandler(p, handler)
}
func (p Parser) SetStartCdataSectionHandler(start StartCdataSectionHandler) {
	setStartCdataSectionHandler(p, start)
}
func (p Parser) SetStartDoctypeDeclHandler(start StartDoctypeDeclHandler) {
	setStartDoctypeDeclHandler(p, start)
}
func (p Parser) SetStartElementHandler(handler StartElementHandler) {
	setStartElementHandler(p, handler)
}
func (p Parser) SetStartNamespaceDeclHandler(start StartNamespaceDeclHandler) {
	setStartNamespaceDeclHandler(p, start)
}
func (p Parser) SetUnknownEncodingHandler(handler UnknownEncodingHandler, encodingHandlerData *T.Void) {
	setUnknownEncodingHandler(p, handler, encodingHandlerData)
}
func (p Parser) SetUnparsedEntityDeclHandler(handler UnparsedEntityDeclHandler) {
	setUnparsedEntityDeclHandler(p, handler)
}
func (p Parser) SetUserData(userData *T.Void)             { setUserData(p, userData) }
func (p Parser) SetXmlDeclHandler(xmldecl XmlDeclHandler) { setXmlDeclHandler(p, xmldecl) }
func (p Parser) SpecifiedAttributeCount() int             { return getSpecifiedAttributeCount(p) }
func (p Parser) StopParser(resumable Bool) Status         { return stopParser(p, resumable) }
func (p Parser) UseForeignDTD(useDTD Bool) Error          { return useForeignDTD(p, useDTD) }
func (p Parser) UseParserAsHandlerArg()                   { useParserAsHandlerArg(p) }

var dll = "libexpat-1.dll"

var apiList = outside.Apis{
	{"XML_DefaultCurrent", &defaultCurrent},
	{"XML_ErrorString", &ErrorString},
	{"XML_ExpatVersion", &ExpatVersion},
	{"XML_ExpatVersionInfo", &ExpatVersionInfo},
	{"XML_ExternalEntityParserCreate", &externalEntityParserCreate},
	{"XML_FreeContentModel", &freeContentModel},
	{"XML_GetBase", &getBase},
	{"XML_GetBuffer", &getBuffer},
	{"XML_GetCurrentByteCount", &getCurrentByteCount},
	{"XML_GetCurrentByteIndex", &getCurrentByteIndex},
	{"XML_GetCurrentColumnNumber", &getCurrentColumnNumber},
	{"XML_GetCurrentLineNumber", &getCurrentLineNumber},
	{"XML_GetErrorCode", &getErrorCode},
	{"XML_GetFeatureList", &GetFeatureList},
	{"XML_GetIdAttributeIndex", &getIdAttributeIndex},
	{"XML_GetInputContext", &getInputContext},
	{"XML_GetParsingStatus", &getParsingStatus},
	{"XML_GetSpecifiedAttributeCount", &getSpecifiedAttributeCount},
	{"XML_MemFree", &memFree},
	{"XML_MemMalloc", &memMalloc},
	{"XML_MemRealloc", &memRealloc},
	{"XML_Parse", &parse},
	{"XML_ParseBuffer", &parseBuffer},
	{"XML_ParserCreate", &ParserCreate},
	{"XML_ParserCreateNS", &ParserCreateNS},
	{"XML_ParserCreate_MM", &ParserCreateMM},
	{"XML_ParserFree", &parserFree},
	{"XML_ParserReset", &parserReset},
	{"XML_ResumeParser", &resumeParser},
	{"XML_SetAttlistDeclHandler", &setAttlistDeclHandler},
	{"XML_SetBase", &setBase},
	{"XML_SetCdataSectionHandler", &setCdataSectionHandler},
	{"XML_SetCharacterDataHandler", &setCharacterDataHandler},
	{"XML_SetCommentHandler", &setCommentHandler},
	{"XML_SetDefaultHandler", &setDefaultHandler},
	{"XML_SetDefaultHandlerExpand", &setDefaultHandlerExpand},
	{"XML_SetDoctypeDeclHandler", &setDoctypeDeclHandler},
	{"XML_SetElementDeclHandler", &setElementDeclHandler},
	{"XML_SetElementHandler", &setElementHandler},
	{"XML_SetEncoding", &setEncoding},
	{"XML_SetEndCdataSectionHandler", &setEndCdataSectionHandler},
	{"XML_SetEndDoctypeDeclHandler", &setEndDoctypeDeclHandler},
	{"XML_SetEndElementHandler", &setEndElementHandler},
	{"XML_SetEndNamespaceDeclHandler", &setEndNamespaceDeclHandler},
	{"XML_SetEntityDeclHandler", &setEntityDeclHandler},
	{"XML_SetExternalEntityRefHandler", &setExternalEntityRefHandler},
	{"XML_SetExternalEntityRefHandlerArg", &setExternalEntityRefHandlerArg},
	{"XML_SetNamespaceDeclHandler", &setNamespaceDeclHandler},
	{"XML_SetNotStandaloneHandler", &setNotStandaloneHandler},
	{"XML_SetNotationDeclHandler", &setNotationDeclHandler},
	{"XML_SetParamEntityParsing", &setParamEntityParsing},
	{"XML_SetProcessingInstructionHandler", &setProcessingInstructionHandler},
	{"XML_SetReturnNSTriplet", &setReturnNSTriplet},
	{"XML_SetSkippedEntityHandler", &setSkippedEntityHandler},
	{"XML_SetStartCdataSectionHandler", &setStartCdataSectionHandler},
	{"XML_SetStartDoctypeDeclHandler", &setStartDoctypeDeclHandler},
	{"XML_SetStartElementHandler", &setStartElementHandler},
	{"XML_SetStartNamespaceDeclHandler", &setStartNamespaceDeclHandler},
	{"XML_SetUnknownEncodingHandler", &setUnknownEncodingHandler},
	{"XML_SetUnparsedEntityDeclHandler", &setUnparsedEntityDeclHandler},
	{"XML_SetUserData", &setUserData},
	{"XML_SetXmlDeclHandler", &setXmlDeclHandler},
	{"XML_StopParser", &stopParser},
	{"XML_UseForeignDTD", &useForeignDTD},
	{"XML_UseParserAsHandlerArg", &useParserAsHandlerArg},
}
