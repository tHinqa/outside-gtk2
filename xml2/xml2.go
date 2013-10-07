package xml2

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-gtk2/types"
)

func init() {
	AddDllApis(dll, false, apiList)
}

type (
	XmlCharEncodingHandlerPtr *XmlCharEncodingHandler
	HtmlDocPtr                XmlDocPtr
	HtmlNodePtr               XmlNodePtr
	HtmlParserCtxt            XmlParserCtxt
	HtmlParserCtxtPtr         XmlParserCtxtPtr
	XmlParserCtxtPtr          *XmlParserCtxt
	XmlAttrPtr                *XmlAttr
	XmlAutomataPtr            *XmlAutomata
	XmlAutomataStatePtr       *XmlAutomataState
	XmlBufPtr                 *XmlBuf
	XmlChar                   Unsigned_char
	XmlDictPtr                *XmlDict
	XmlDocPtr                 *XmlDoc
	XmlElementContentPtr      *XmlElementContent
	XmlEntityPtr              *XmlEntity
	XmlEnumerationPtr         *XmlEnumeration
	XmlErrorPtr               *XmlError
	XmlHashTablePtr           *XmlHashTable
	XmlNodePtr                *XmlNode
	XmlNsPtr                  *XmlNs
	XmlNsType                 XmlElementType
	XmlOutputBufferPtr        *XmlOutputBuffer
	XmlParserInputBufferPtr   *XmlParserInputBuffer
	XmlParserInputPtr         *XmlParserInput
	XmlParserNodeInfoPtr      *XmlParserNodeInfo
	XmlParserNodeInfoSeqPtr   *XmlParserNodeInfoSeq
	XmlSAXLocatorPtr          *XmlSAXLocator
	XmlValidStatePtr          *XmlValidState
	XmlRegexpPtr              *XmlRegexp
	HtmlSAXHandler            XmlSAXHandler
	HtmlSAXHandlerPtr         XmlSAXHandlerPtr
	XmlSAXHandlerPtr          *XmlSAXHandler
	XmlBufferPtr              *XmlBuffer
	XlinkHandlerPtr           *XlinkHandler
	XlinkHRef                 *XmlChar
	XlinkRole                 *XmlChar
	XlinkTitle                *XmlChar
	XmlAttributePtr           *XmlAttribute
	XmlNotationPtr            *XmlNotation
	XmlValidCtxtPtr           *XmlValidCtxt
	XmlDtdPtr                 *XmlDtd
	XmlElementPtr             *XmlElement
	XmlIDPtr                  *XmlID
	XmlRefPtr                 *XmlRef
	XmlNotationTablePtr       *XmlNotationTable
	XmlNotationTable          XmlHashTable
	XmlElementTable           XmlHashTable
	XmlElementTablePtr        *XmlElementTable
	XmlEntitiesTable          XmlHashTable
	XmlEntitiesTablePtr       *XmlEntitiesTable
	XmlAttributeTable         XmlHashTable
	XmlAttributeTablePtr      *XmlAttributeTable
	XmlDOMWrapCtxtPtr         *XmlDOMWrapCtxt
	XmlExpCtxtPtr             *XmlExpCtxt
	XmlExpNodePtr             *XmlExpNode
	XmlIDTable                XmlHashTable
	XmlIDTablePtr             *XmlIDTable
	XmlRefTable               XmlHashTable
	XmlRefTablePtr            *XmlRefTable
	XmlMutexPtr               *XmlMutex
	XmlListPtr                *XmlList
	XmlRMutexPtr              *XmlRMutex
	XmlGlobalStatePtr         *XmlGlobalState
	XmlLinkPtr                *XmlLink
	XmlRegExecCtxtPtr         *XmlRegExecCtxt

	// Temp
	XmlNode           struct{}
	XmlDoc            struct{}
	XmlEntity         struct{}
	XmlEnumeration    struct{}
	XmlElementContent struct{}
	XmlAttr           struct{}
	XmlNs             struct{}

	XmlValidState    struct{}
	XmlBuf           struct{}
	XmlDict          struct{}
	XmlAutomata      struct{}
	XmlAutomataState struct{}
	XmlHashTable     struct{}
	XmlRegexp        struct{}
	XmlDOMWrapCtxt   struct{}
	XmlExpCtxt       struct{}
	XmlExpNode       struct{}
	XmlMutex         struct{}
	XmlRMutex        struct{}
	XmlList          struct{}
	XmlLink          struct{}
	XmlRegExecCtxt   struct{}
	_                struct{}
	_                struct{}
	_                struct{}
	_                struct{}
	_                struct{}
	_                struct{}
	_                struct{}
	_                struct{}
	XmlBuffer        struct{}

	InternalSubsetSAXFunc func(
		ctx *Void, name, ExternalID, SystemID *XmlChar)

	IsStandaloneSAXFunc func(ctx *Void) int

	HasInternalSubsetSAXFunc func(ctx *Void) int

	HasExternalSubsetSAXFunc func(ctx *Void) int

	StartElementNsSAX2Func func(ctx *Void,
		localname, prefix, URI *XmlChar,
		nb_namespaces int, namespaces **XmlChar,
		nb_attributes, nb_defaulted int, attributes **XmlChar)

	EndElementNsSAX2Func func(
		ctx *Void, localname, prefix, URI *XmlChar)

	ResolveEntitySAXFunc func(
		ctx *Void,
		publicId *XmlChar,
		systemId *XmlChar) XmlParserInputPtr

	ExternalSubsetSAXFunc func(
		ctx *Void,
		name *XmlChar,
		ExternalID *XmlChar,
		SystemID *XmlChar)

	GetEntitySAXFunc func(
		ctx *Void, name *XmlChar) XmlEntityPtr

	GetParameterEntitySAXFunc func(
		ctx *Void, name *XmlChar) XmlEntityPtr

	EntityDeclSAXFunc func(ctx *Void, name *XmlChar,
		typ int, publicId, systemId, content *XmlChar)

	NotationDeclSAXFunc func(
		ctx *Void, name, publicId, systemId *XmlChar)

	AttributeDeclSAXFunc func(ctx *Void, elem,
		fullname *XmlChar, typ, def int,
		defaultValue *XmlChar, tree XmlEnumerationPtr)

	ElementDeclSAXFunc func(ctx *Void, name *XmlChar,
		typ int, content XmlElementContentPtr)

	UnparsedEntityDeclSAXFunc func(ctx *Void,
		name, publicId, systemId, notationName *XmlChar)

	SetDocumentLocatorSAXFunc func(
		ctx *Void, loc XmlSAXLocatorPtr)

	StartDocumentSAXFunc func(ctx *Void)

	EndDocumentSAXFunc func(
		ctx *Void)

	StartElementSAXFunc func(
		ctx *Void, name *XmlChar, atts **XmlChar)

	EndElementSAXFunc func(ctx *Void, name *XmlChar)

	AttributeSAXFunc func(ctx *Void, name, value *XmlChar)

	ReferenceSAXFunc func(ctx *Void, name *XmlChar)

	CharactersSAXFunc func(ctx *Void, ch *XmlChar, leng int)

	IgnorableWhitespaceSAXFunc func(
		ctx *Void, ch *XmlChar, leng int)

	ProcessingInstructionSAXFunc func(
		ctx *Void, target, data *XmlChar)

	CommentSAXFunc func(ctx *Void, value *XmlChar)

	CdataBlockSAXFunc func(ctx *Void, value *XmlChar, leng int)

	XmlParserInputDeallocate func(str *XmlChar)

	XmlInputReadCallback func(
		context *Void, buffer *Char, leng int) int

	XmlInputCloseCallback func(context *Void) int

	WarningSAXFunc func(ctx *Void, msg *Char, v ...VArg)

	ErrorSAXFunc func(ctx *Void, msg *Char, v ...VArg)

	FatalErrorSAXFunc func(ctx *Void, msg *Char, v ...VArg)

	XmlDeregisterNodeFunc func(node XmlNodePtr)

	XmlGenericErrorFunc func(ctx *Void, msg *Char, v ...VArg)

	XmlCharEncodingInputFunc func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int) int

	XmlCharEncodingOutputFunc func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int) int

	XmlParserInputBufferCreateFilenameFunc func(
		URI *Char,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlRegisterNodeFunc func(node XmlNodePtr)

	XmlStructuredErrorFunc func(
		userData *Void, Error XmlErrorPtr)

	XmlOutputWriteCallback func(
		context *Void, buffer *Char, leng int) int

	XmlOutputCloseCallback func(context *Void) int

	XmlParserCtxt struct {
		sax               *XmlSAXHandler
		userData          *Void
		myDoc             XmlDocPtr
		wellFormed        int
		replaceEntities   int
		version           *XmlChar
		encoding          *XmlChar
		standalone        int
		html              int
		input             XmlParserInputPtr
		inputNr           int
		inputMax          int
		inputTab          *XmlParserInputPtr
		node              XmlNodePtr
		nodeNr            int
		nodeMax           int
		nodeTab           *XmlNodePtr
		record_info       int
		node_seq          XmlParserNodeInfoSeq
		errNo             int
		hasExternalSubset int
		hasPErefs         int
		external          int
		valid             int
		validate          int
		vctxt             XmlValidCtxt
		instate           XmlParserInputState
		token             int
		directory         *Char
		name              *XmlChar
		nameNr            int
		nameMax           int
		nameTab           **XmlChar
		nbChars           Long
		checkIndex        Long
		keepBlanks        int
		disableSAX        int
		inSubset          int
		intSubName        *XmlChar
		extSubURI         *XmlChar
		extSubSystem      *XmlChar
		space             *int
		spaceNr           int
		spaceMax          int
		spaceTab          *int
		depth             int
		entity            XmlParserInputPtr
		charset           int
		nodelen           int
		nodemem           int
		pedantic          int
		_private          *Void
		loadsubset        int
		linenumbers       int
		catalogs          *Void
		recovery          int
		progressive       int
		dict              XmlDictPtr
		atts              **XmlChar
		maxatts           int
		docdict           int
		str_xml           *XmlChar
		str_XmlNs         *XmlChar
		str_xml_ns        *XmlChar
		sax2              int
		nsNr              int
		nsMax             int
		nsTab             **XmlChar
		attallocs         *int
		pushTab           **Void
		attsDefault       XmlHashTablePtr
		attsSpecial       XmlHashTablePtr
		nsWellFormed      int
		options           int
		dictNames         int
		freeElemsNr       int
		freeElems         XmlNodePtr
		freeAttrsNr       int
		freeAttrs         XmlAttrPtr
		lastError         XmlError
		parseMode         XmlParserMode
		nbentities        Unsigned_long
		sizeentities      Unsigned_long
		nodeInfo          *XmlParserNodeInfo
		nodeInfoNr        int
		nodeInfoMax       int
		nodeInfoTab       *XmlParserNodeInfo
		input_id          int
	}

	XmlValidityErrorFunc func(ctx *Void, msg string, v ...VArg)

	XmlValidityWarningFunc func(ctx *Void, msg string, v ...VArg)

	XmlOutputBufferCreateFilenameFunc func(
		URI *Char,
		encoder XmlCharEncodingHandlerPtr,
		compression int) XmlOutputBufferPtr

	XlinkNodeDetectFunc func(
		ctx *Void,
		node XmlNodePtr)

	XlinkSimpleLinkFunk func(
		ctx *Void,
		node XmlNodePtr,
		href XlinkHRef,
		role XlinkRole,
		title XlinkTitle)

	XlinkExtendedLinkFunk func(
		ctx *Void,
		node XmlNodePtr,
		nbLocators int,
		hrefs *XlinkHRef,
		roles *XlinkRole,
		nbArcs int,
		from *XlinkRole,
		to *XlinkRole,
		show *XlinkShow,
		actuate *XlinkActuate,
		nbTitles int,
		titles *XlinkTitle,
		langs **XmlChar)

	XlinkExtendedLinkSetFunk func(
		ctx *Void,
		node XmlNodePtr,
		nbLocators int,
		hrefs *XlinkHRef,
		roles *XlinkRole,
		nbTitles int,
		titles *XlinkTitle,
		langs **XmlChar)

	XmlFreeFunc func(mem *Void)

	XmlMallocFunc func(size Size_t) *Void

	XmlReallocFunc func(mem *Void, size Size_t) *Void

	XmlStrdupFunc func(str *Char) *Char

	XmlHashDeallocator func(payload *Void, name *XmlChar)

	XmlHashCopier func(payload *Void, name *XmlChar) *Void

	XmlHashScanner func(payload, data *Void, name *XmlChar)

	XmlHashScannerFull func(payload, data *Void,
		name, name2, name3 *XmlChar)

	XmlExternalEntityLoader func(URL, ID *Char,
		context XmlParserCtxtPtr) XmlParserInputPtr

	XmlRegExecCallbacks func(
		exec XmlRegExecCtxtPtr,
		token *XmlChar,
		transdata, inputdata *Void)

	XmlListDeallocator func(lk XmlLinkPtr)

	XmlListDataCompare func(data0, data1 *Void) int

	XmlListWalker func(data, user *Void) int

	XmlInputMatchCallback func(filename string) int

	XmlInputOpenCallback func(filename string) *Void

	XmlOutputMatchCallback func(filename string) int

	XmlOutputOpenCallback func(filename string) *Void

	XmlSAXHandlerV1 struct {
		internalSubset        InternalSubsetSAXFunc
		isStandalone          IsStandaloneSAXFunc
		hasInternalSubset     HasInternalSubsetSAXFunc
		hasExternalSubset     HasExternalSubsetSAXFunc
		resolveEntity         ResolveEntitySAXFunc
		getEntity             GetEntitySAXFunc
		entityDecl            EntityDeclSAXFunc
		notationDecl          NotationDeclSAXFunc
		attributeDecl         AttributeDeclSAXFunc
		elementDecl           ElementDeclSAXFunc
		unparsedEntityDecl    UnparsedEntityDeclSAXFunc
		setDocumentLocator    SetDocumentLocatorSAXFunc
		startDocument         StartDocumentSAXFunc
		endDocument           EndDocumentSAXFunc
		startElement          StartElementSAXFunc
		endElement            EndElementSAXFunc
		reference             ReferenceSAXFunc
		characters            CharactersSAXFunc
		ignorableWhitespace   IgnorableWhitespaceSAXFunc
		processingInstruction ProcessingInstructionSAXFunc
		comment               CommentSAXFunc
		warning               WarningSAXFunc
		error                 ErrorSAXFunc
		fatalError            FatalErrorSAXFunc
		getParameterEntity    GetParameterEntitySAXFunc
		cdataBlock            CdataBlockSAXFunc
		externalSubset        ExternalSubsetSAXFunc
		initialized           Unsigned_int
	}

	XmlParserInput struct {
		buf        XmlParserInputBufferPtr
		filename   *Char
		directory  *Char
		base       *XmlChar
		cur        *XmlChar
		end        *XmlChar
		length     int
		line       int
		col        int
		consumed   Unsigned_long
		free       XmlParserInputDeallocate
		encoding   *XmlChar
		version    *XmlChar
		standalone int
		id         int
	}

	XmlCharEncodingHandler struct {
		name   *Char
		input  XmlCharEncodingInputFunc
		output XmlCharEncodingOutputFunc
	}

	XmlError struct {
		domain  int
		code    int
		message *Char
		level   XmlErrorLevel
		file    *Char
		line    int
		str1    *Char
		str2    *Char
		str3    *Char
		int1    int
		int2    int
		ctxt    *Void
		node    *Void
	}

	XmlOutputBuffer struct {
		context       *Void
		writecallback XmlOutputWriteCallback
		closecallback XmlOutputCloseCallback
		encoder       XmlCharEncodingHandlerPtr
		buffer        XmlBufPtr
		conv          XmlBufPtr
		written       int
		Error         int
	}

	XmlParserInputBuffer struct {
		context       *Void
		readcallback  XmlInputReadCallback
		closecallback XmlInputCloseCallback
		encoder       XmlCharEncodingHandlerPtr
		buffer        XmlBufPtr
		raw           XmlBufPtr
		compressed    int
		error         int
		rawconsumed   Unsigned_long
	}

	XmlSAXLocator struct {
		getPublicId     func(ctx *Void) *XmlChar
		getSystemId     func(ctx *Void) *XmlChar
		getLineNumber   func(ctx *Void) int
		getColumnNumber func(ctx *Void) int
	}

	XmlSAXHandler struct {
		internalSubset        InternalSubsetSAXFunc
		isStandalone          IsStandaloneSAXFunc
		hasInternalSubset     HasInternalSubsetSAXFunc
		hasExternalSubset     HasExternalSubsetSAXFunc
		resolveEntity         ResolveEntitySAXFunc
		getEntity             GetEntitySAXFunc
		entityDecl            EntityDeclSAXFunc
		notationDecl          NotationDeclSAXFunc
		attributeDecl         AttributeDeclSAXFunc
		elementDecl           ElementDeclSAXFunc
		unparsedEntityDecl    UnparsedEntityDeclSAXFunc
		setDocumentLocator    SetDocumentLocatorSAXFunc
		startDocument         StartDocumentSAXFunc
		endDocument           EndDocumentSAXFunc
		startElement          StartElementSAXFunc
		endElement            EndElementSAXFunc
		reference             ReferenceSAXFunc
		characters            CharactersSAXFunc
		ignorableWhitespace   IgnorableWhitespaceSAXFunc
		processingInstruction ProcessingInstructionSAXFunc
		comment               CommentSAXFunc
		warning               WarningSAXFunc
		error                 ErrorSAXFunc
		fatalError            FatalErrorSAXFunc
		getParameterEntity    GetParameterEntitySAXFunc
		cdataBlock            CdataBlockSAXFunc
		externalSubset        ExternalSubsetSAXFunc
		initialized           Unsigned_int
		_                     *Void
		startElementNs        StartElementNsSAX2Func
		endElementNs          EndElementNsSAX2Func
		serror                XmlStructuredErrorFunc
	}

	XmlParserNodeInfoSeq struct {
		maximum Unsigned_long
		length  Unsigned_long
		buffer  *XmlParserNodeInfo
	}

	XmlParserNodeInfo struct {
		node       *XmlNode
		begin_pos  Unsigned_long
		begin_line Unsigned_long
		end_pos    Unsigned_long
		end_line   Unsigned_long
	}

	XmlValidCtxt struct {
		userData  *Void
		Error     XmlValidityErrorFunc
		warning   XmlValidityWarningFunc
		node      XmlNodePtr
		nodeNr    int
		nodeMax   int
		nodeTab   *XmlNodePtr
		finishDtd Unsigned_int
		doc       XmlDocPtr
		valid     int
		vstate    *XmlValidState
		vstateNr  int
		vstateMax int
		vstateTab *XmlValidState
		am        XmlAutomataPtr
		state     XmlAutomataStatePtr
	}

	XmlDtd struct {
		_          *Void
		Type       XmlElementType
		name       *XmlChar
		children   *XmlNode
		last       *XmlNode
		parent     *XmlDoc
		next       *XmlNode
		prev       *XmlNode
		doc        *XmlDoc
		notations  *Void
		elements   *Void
		attributes *Void
		entities   *Void
		ExternalID *XmlChar
		SystemID   *XmlChar
		pentities  *Void
	}

	HtmlElemDesc struct {
		name          *Char
		startTag      Char
		endTag        Char
		saveEndTag    Char
		empty         Char
		depr          Char
		dtd           Char
		isinline      Char
		desc          *Char
		subelts       **Char
		defaultsubelt *Char
		attrs_opt     **Char
		attrs_depr    **Char
		attrs_req     **Char
	}

	HtmlEntityDesc struct {
		value Unsigned_int
		name  *Char
		desc  *Char
	}

	XlinkHandler struct {
		simple   XlinkSimpleLinkFunk
		extended XlinkExtendedLinkFunk
		set      XlinkExtendedLinkSetFunk
	}

	XmlAttribute struct {
		_    *Void
		Type XmlElementType
		name *XmlChar
	}

	XmlNotation struct {
		name     *XmlChar
		PublicID *XmlChar
		SystemID *XmlChar
	}

	XmlElement struct {
		_          *Void
		Type       XmlElementType
		name       *XmlChar
		children   *XmlNode
		last       *XmlNode
		parent     *XmlDtd
		next       *XmlNode
		prev       *XmlNode
		doc        *XmlDoc
		etype      XmlElementTypeVal
		content    XmlElementContentPtr
		attributes XmlAttributePtr
		prefix     *XmlChar
		contModel  XmlRegexpPtr
	}

	XmlID struct {
		next   *XmlID
		value  *XmlChar
		attr   XmlAttrPtr
		name   *XmlChar
		lineno int
		doc    *XmlDoc
	}

	XmlRef struct {
		next   *XmlRef
		value  *XmlChar
		attr   XmlAttrPtr
		name   *XmlChar
		lineno int
	}

	XmlGlobalState struct {
		xmlParserVersion                        *Char
		xmlDefaultSAXLocator                    XmlSAXLocator
		xmlDefaultSAXHandler                    XmlSAXHandlerV1
		docbDefaultSAXHandler                   XmlSAXHandlerV1
		htmlDefaultSAXHandler                   XmlSAXHandlerV1
		xmlFree                                 XmlFreeFunc
		xmlMalloc                               XmlMallocFunc
		xmlMemStrdup                            XmlStrdupFunc
		xmlRealloc                              XmlReallocFunc
		xmlGenericError                         XmlGenericErrorFunc
		xmlStructuredError                      XmlStructuredErrorFunc
		xmlGenericErrorContext                  *Void
		oldXMLWDcompatibility                   int
		XmlBufferAllocScheme                    XmlBufferAllocationScheme
		xmlDefaultBufferSize                    int
		xmlSubstituteEntitiesDefaultValue       int
		xmlDoValidityCheckingDefaultValue       int
		xmlGetWarningsDefaultValue              int
		xmlKeepBlanksDefaultValue               int
		xmlLineNumbersDefaultValue              int
		xmlLoadExtDtdDefaultValue               int
		xmlParserDebugEntities                  int
		xmlPedanticParserDefaultValue           int
		xmlSaveNoEmptyTags                      int
		xmlIndentTreeOutput                     int
		xmlTreeIndentString                     *Char
		xmlRegisterNodeDefaultValue             XmlRegisterNodeFunc
		xmlDeregisterNodeDefaultValue           XmlDeregisterNodeFunc
		xmlMallocAtomic                         XmlMallocFunc
		xmlLastError                            XmlError
		xmlParserInputBufferCreateFilenameValue XmlParserInputBufferCreateFilenameFunc
		xmlOutputBufferCreateFilenameValue      XmlOutputBufferCreateFilenameFunc
		xmlStructuredErrorContext               *Void
	}
)

type XmlEntityType Enum

const (
	XML_INTERNAL_GENERAL_ENTITY XmlEntityType = iota + 1
	XML_EXTERNAL_GENERAL_PARSED_ENTITY
	XML_EXTERNAL_GENERAL_UNPARSED_ENTITY
	XML_INTERNAL_PARAMETER_ENTITY
	XML_EXTERNAL_PARAMETER_ENTITY
	XML_INTERNAL_PREDEFINED_ENTITY
)

type XmlElementType Enum

const (
	XML_ELEMENT_NODE XmlElementType = iota + 1
	XML_ATTRIBUTE_NODE
	XML_TEXT_NODE
	XML_CDATA_SECTION_NODE
	XML_ENTITY_REF_NODE
	XML_ENTITY_NODE
	XML_PI_NODE
	XML_COMMENT_NODE
	XML_DOCUMENT_NODE
	XML_DOCUMENT_TYPE_NODE
	XML_DOCUMENT_FRAG_NODE
	XML_NOTATION_NODE
	XML_HTML_DOCUMENT_NODE
	XML_DTD_NODE
	XML_ELEMENT_DECL
	XML_ATTRIBUTE_DECL
	XML_ENTITY_DECL
	XML_NAMESPACE_DECL
	XML_XINCLUDE_START
	XML_XINCLUDE_END
	XML_DOCB_DOCUMENT_NODE
)

type XmlBufferAllocationScheme Enum

const (
	XML_BUFFER_ALLOC_DOUBLEIT XmlBufferAllocationScheme = iota
	XML_BUFFER_ALLOC_EXACT
	XML_BUFFER_ALLOC_IMMUTABLE
	XML_BUFFER_ALLOC_IO
	XML_BUFFER_ALLOC_HYBRID
)

type XmlElementContentType Enum

const (
	XML_ELEMENT_CONTENT_PCDATA XmlElementContentType = iota + 1
	XML_ELEMENT_CONTENT_ELEMENT
	XML_ELEMENT_CONTENT_SEQ
	XML_ELEMENT_CONTENT_OR
)

type XmlElementContentOccur Enum

const (
	XML_ELEMENT_CONTENT_ONCE XmlElementContentOccur = iota + 1
	XML_ELEMENT_CONTENT_OPT
	XML_ELEMENT_CONTENT_MULT
	XML_ELEMENT_CONTENT_PLUS
)

type XmlCharEncoding Enum

const (
	XML_CHAR_ENCODING_ERROR XmlCharEncoding = iota - 1
	XML_CHAR_ENCODING_NONE
	XML_CHAR_ENCODING_UTF8
	XML_CHAR_ENCODING_UTF16LE
	XML_CHAR_ENCODING_UTF16BE
	XML_CHAR_ENCODING_UCS4LE
	XML_CHAR_ENCODING_UCS4BE
	XML_CHAR_ENCODING_EBCDIC
	XML_CHAR_ENCODING_UCS4_2143
	XML_CHAR_ENCODING_UCS4_3412
	XML_CHAR_ENCODING_UCS2
	XML_CHAR_ENCODING_8859_1
	XML_CHAR_ENCODING_8859_2
	XML_CHAR_ENCODING_8859_3
	XML_CHAR_ENCODING_8859_4
	XML_CHAR_ENCODING_8859_5
	XML_CHAR_ENCODING_8859_6
	XML_CHAR_ENCODING_8859_7
	XML_CHAR_ENCODING_8859_8
	XML_CHAR_ENCODING_8859_9
	XML_CHAR_ENCODING_2022_JP
	XML_CHAR_ENCODING_SHIFT_JIS
	XML_CHAR_ENCODING_EUC_JP
	XML_CHAR_ENCODING_ASCII
)

type XmlAttributeType Enum

const (
	XML_ATTRIBUTE_CDATA XmlAttributeType = iota + 1
	XML_ATTRIBUTE_ID
	XML_ATTRIBUTE_IDREF
	XML_ATTRIBUTE_IDREFS
	XML_ATTRIBUTE_ENTITY
	XML_ATTRIBUTE_ENTITIES
	XML_ATTRIBUTE_NMTOKEN
	XML_ATTRIBUTE_NMTOKENS
	XML_ATTRIBUTE_ENUMERATION
	XML_ATTRIBUTE_NOTATION
)

type XmlErrorLevel Enum

const (
	XML_ERR_NONE XmlErrorLevel = iota
	XML_ERR_WARNING
	XML_ERR_ERROR
	XML_ERR_FATAL
)

type XmlParserInputState Enum

const (
	XML_PARSER_EOF XmlParserInputState = iota - 1
	XML_PARSER_START
	XML_PARSER_MISC
	XML_PARSER_PI
	XML_PARSER_DTD
	XML_PARSER_PROLOG
	XML_PARSER_COMMENT
	XML_PARSER_START_TAG
	XML_PARSER_CONTENT
	XML_PARSER_CDATA_SECTION
	XML_PARSER_END_TAG
	XML_PARSER_ENTITY_DECL
	XML_PARSER_ENTITY_VALUE
	XML_PARSER_ATTRIBUTE_VALUE
	XML_PARSER_SYSTEM_LITERAL
	XML_PARSER_EPILOG
	XML_PARSER_IGNORE
	XML_PARSER_PUBLIC_LITERAL
)

type XmlParserMode Enum

const (
	XML_PARSE_UNKNOWN XmlParserMode = iota
	XML_PARSE_DOM
	XML_PARSE_SAX
	XML_PARSE_PUSH_DOM
	XML_PARSE_PUSH_SAX
	XML_PARSE_READER
)

type HtmlStatus Enum

const (
	HTML_INVALID HtmlStatus = 1 << iota
	HTML_DEPRECATED
	HTML_VALID
	htmlStatusAnon
	HTML_REQUIRED            = HTML_VALID | htmlStatusAnon
	HTML_NA       HtmlStatus = 0
)

type XlinkShow Enum

const (
	XLINK_SHOW_NONE XlinkShow = iota
	XLINK_SHOW_NEW
	XLINK_SHOW_EMBED
	XLINK_SHOW_REPLACE
)

type XlinkActuate Enum

const (
	XLINK_ACTUATE_NONE XlinkActuate = iota
	XLINK_ACTUATE_AUTO
	XLINK_ACTUATE_ONREQUEST
)

type XlinkType Enum

const (
	XLINK_TYPE_NONE XlinkType = iota
	XLINK_TYPE_SIMPLE
	XLINK_TYPE_EXTENDED
	XLINK_TYPE_EXTENDED_SET
)

type XmlElementTypeVal Enum

const (
	XML_ELEMENT_TYPE_UNDEFINED XmlElementTypeVal = iota
	XML_ELEMENT_TYPE_EMPTY
	XML_ELEMENT_TYPE_ANY
	XML_ELEMENT_TYPE_MIXED
	XML_ELEMENT_TYPE_ELEMENT
)

type XmlAttributeDefault Enum

const (
	XML_ATTRIBUTE_NONE XmlAttributeDefault = iota + 1
	XML_ATTRIBUTE_REQUIRED
	XML_ATTRIBUTE_IMPLIED
	XML_ATTRIBUTE_FIXED
)

type XmlFeature Enum

const (
	XML_WITH_THREAD XmlFeature = iota + 1
	XML_WITH_TREE
	XML_WITH_OUTPUT
	XML_WITH_PUSH
	XML_WITH_READER
	XML_WITH_PATTERN
	XML_WITH_WRITER
	XML_WITH_SAX1
	XML_WITH_FTP
	XML_WITH_HTTP
	XML_WITH_VALID
	XML_WITH_HTML
	XML_WITH_LEGACY
	XML_WITH_C14N
	XML_WITH_CATALOG
	XML_WITH_XPATH
	XML_WITH_XPTR
	XML_WITH_XINCLUDE
	XML_WITH_ICONV
	XML_WITH_ISO8859X
	XML_WITH_UNICODE
	XML_WITH_REGEXP
	XML_WITH_AUTOMATA
	XML_WITH_EXPR
	XML_WITH_SCHEMAS
	XML_WITH_SCHEMATRON
	XML_WITH_MODULES
	XML_WITH_DEBUG
	XML_WITH_DEBUG_MEM
	XML_WITH_DEBUG_RUN
	XML_WITH_ZLIB
	XML_WITH_ICU
	XML_WITH_LZMA
	XML_WITH_NONE XmlFeature = 99999
)

type XmlParserErrors Enum

const (
	XML_ERR_OK XmlParserErrors = iota
	XML_ERR_INTERNAL_ERROR
	XML_ERR_NO_MEMORY
	XML_ERR_DOCUMENT_START
	XML_ERR_DOCUMENT_EMPTY
	XML_ERR_DOCUMENT_END
	XML_ERR_INVALID_HEX_CHARREF
	XML_ERR_INVALID_DEC_CHARREF
	XML_ERR_INVALID_CHARREF
	XML_ERR_INVALID_CHAR
	XML_ERR_CHARREF_AT_EOF
	XML_ERR_CHARREF_IN_PROLOG
	XML_ERR_CHARREF_IN_EPILOG
	XML_ERR_CHARREF_IN_DTD
	XML_ERR_ENTITYREF_AT_EOF
	XML_ERR_ENTITYREF_IN_PROLOG
	XML_ERR_ENTITYREF_IN_EPILOG
	XML_ERR_ENTITYREF_IN_DTD
	XML_ERR_PEREF_AT_EOF
	XML_ERR_PEREF_IN_PROLOG
	XML_ERR_PEREF_IN_EPILOG
	XML_ERR_PEREF_IN_INT_SUBSET
	XML_ERR_ENTITYREF_NO_NAME
	XML_ERR_ENTITYREF_SEMICOL_MISSING
	XML_ERR_PEREF_NO_NAME
	XML_ERR_PEREF_SEMICOL_MISSING
	XML_ERR_UNDECLARED_ENTITY
	XML_WAR_UNDECLARED_ENTITY
	XML_ERR_UNPARSED_ENTITY
	XML_ERR_ENTITY_IS_EXTERNAL
	XML_ERR_ENTITY_IS_PARAMETER
	XML_ERR_UNKNOWN_ENCODING
	XML_ERR_UNSUPPORTED_ENCODING
	XML_ERR_STRING_NOT_STARTED
	XML_ERR_STRING_NOT_CLOSED
	XML_ERR_NS_DECL_ERROR
	XML_ERR_ENTITY_NOT_STARTED
	XML_ERR_ENTITY_NOT_FINISHED
	XML_ERR_LT_IN_ATTRIBUTE
	XML_ERR_ATTRIBUTE_NOT_STARTED
	XML_ERR_ATTRIBUTE_NOT_FINISHED
	XML_ERR_ATTRIBUTE_WITHOUT_VALUE
	XML_ERR_ATTRIBUTE_REDEFINED
	XML_ERR_LITERAL_NOT_STARTED
	XML_ERR_LITERAL_NOT_FINISHED
	XML_ERR_COMMENT_NOT_FINISHED
	XML_ERR_PI_NOT_STARTED
	XML_ERR_PI_NOT_FINISHED
	XML_ERR_NOTATION_NOT_STARTED
	XML_ERR_NOTATION_NOT_FINISHED
	XML_ERR_ATTLIST_NOT_STARTED
	XML_ERR_ATTLIST_NOT_FINISHED
	XML_ERR_MIXED_NOT_STARTED
	XML_ERR_MIXED_NOT_FINISHED
	XML_ERR_ELEMCONTENT_NOT_STARTED
	XML_ERR_ELEMCONTENT_NOT_FINISHED
	XML_ERR_XMLDECL_NOT_STARTED
	XML_ERR_XMLDECL_NOT_FINISHED
	XML_ERR_CONDSEC_NOT_STARTED
	XML_ERR_CONDSEC_NOT_FINISHED
	XML_ERR_EXT_SUBSET_NOT_FINISHED
	XML_ERR_DOCTYPE_NOT_FINISHED
	XML_ERR_MISPLACED_CDATA_END
	XML_ERR_CDATA_NOT_FINISHED
	XML_ERR_RESERVED_XML_NAME
	XML_ERR_SPACE_REQUIRED
	XML_ERR_SEPARATOR_REQUIRED
	XML_ERR_NMTOKEN_REQUIRED
	XML_ERR_NAME_REQUIRED
	XML_ERR_PCDATA_REQUIRED
	XML_ERR_URI_REQUIRED
	XML_ERR_PUBID_REQUIRED
	XML_ERR_LT_REQUIRED
	XML_ERR_GT_REQUIRED
	XML_ERR_LTSLASH_REQUIRED
	XML_ERR_EQUAL_REQUIRED
	XML_ERR_TAG_NAME_MISMATCH
	XML_ERR_TAG_NOT_FINISHED
	XML_ERR_STANDALONE_VALUE
	XML_ERR_ENCODING_NAME
	XML_ERR_HYPHEN_IN_COMMENT
	XML_ERR_INVALID_ENCODING
	XML_ERR_EXT_ENTITY_STANDALONE
	XML_ERR_CONDSEC_INVALID
	XML_ERR_VALUE_REQUIRED
	XML_ERR_NOT_WELL_BALANCED
	XML_ERR_EXTRA_CONTENT
	XML_ERR_ENTITY_CHAR_ERROR
	XML_ERR_ENTITY_PE_INTERNAL
	XML_ERR_ENTITY_LOOP
	XML_ERR_ENTITY_BOUNDARY
	XML_ERR_INVALID_URI
	XML_ERR_URI_FRAGMENT
	XML_WAR_CATALOG_PI
	XML_ERR_NO_DTD
	XML_ERR_CONDSEC_INVALID_KEYWORD
	XML_ERR_VERSION_MISSING
	XML_WAR_UNKNOWN_VERSION
	XML_WAR_LANG_VALUE
	XML_WAR_NS_URI
	XML_WAR_NS_URI_RELATIVE
	XML_ERR_MISSING_ENCODING
	XML_WAR_SPACE_VALUE
	XML_ERR_NOT_STANDALONE
	XML_ERR_ENTITY_PROCESSING
	XML_ERR_NOTATION_PROCESSING
	XML_WAR_NS_COLUMN
	XML_WAR_ENTITY_REDEFINED
	XML_ERR_UNKNOWN_VERSION
	XML_ERR_VERSION_MISMATCH
	XML_ERR_NAME_TOO_LONG
)
const (
	XML_NS_ERR_XML_NAMESPACE XmlParserErrors = iota + 200
	XML_NS_ERR_UNDEFINED_NAMESPACE
	XML_NS_ERR_QNAME
	XML_NS_ERR_ATTRIBUTE_REDEFINED
	XML_NS_ERR_EMPTY
	XML_NS_ERR_COLON
)
const (
	XML_DTD_ATTRIBUTE_DEFAULT XmlParserErrors = iota + 500
	XML_DTD_ATTRIBUTE_REDEFINED
	XML_DTD_ATTRIBUTE_VALUE
	XML_DTD_CONTENT_ERROR
	XML_DTD_CONTENT_MODEL
	XML_DTD_CONTENT_NOT_DETERMINIST
	XML_DTD_DIFFERENT_PREFIX
	XML_DTD_ELEM_DEFAULT_NAMESPACE
	XML_DTD_ELEM_NAMESPACE
	XML_DTD_ELEM_REDEFINED
	XML_DTD_EMPTY_NOTATION
	XML_DTD_ENTITY_TYPE
	XML_DTD_ID_FIXED
	XML_DTD_ID_REDEFINED
	XML_DTD_ID_SUBSET
	XML_DTD_INVALID_CHILD
	XML_DTD_INVALID_DEFAULT
	XML_DTD_LOAD_ERROR
	XML_DTD_MISSING_ATTRIBUTE
	XML_DTD_MIXED_CORRUPT
	XML_DTD_MULTIPLE_ID
	XML_DTD_NO_DOC
	XML_DTD_NO_DTD
	XML_DTD_NO_ELEM_NAME
	XML_DTD_NO_PREFIX
	XML_DTD_NO_ROOT
	XML_DTD_NOTATION_REDEFINED
	XML_DTD_NOTATION_VALUE
	XML_DTD_NOT_EMPTY
	XML_DTD_NOT_PCDATA
	XML_DTD_NOT_STANDALONE
	XML_DTD_ROOT_NAME
	XML_DTD_STANDALONE_WHITE_SPACE
	XML_DTD_UNKNOWN_ATTRIBUTE
	XML_DTD_UNKNOWN_ELEM
	XML_DTD_UNKNOWN_ENTITY
	XML_DTD_UNKNOWN_ID
	XML_DTD_UNKNOWN_NOTATION
	XML_DTD_STANDALONE_DEFAULTED
	XML_DTD_XMLID_VALUE
	XML_DTD_XMLID_TYPE
	XML_DTD_DUP_TOKEN
)
const (
	XML_HTML_STRUCURE_ERROR XmlParserErrors = iota + 800
	XML_HTML_UNKNOWN_TAG
)
const (
	XML_RNGP_ANYNAME_ATTR_ANCESTOR XmlParserErrors = iota + 1000
	XML_RNGP_ATTR_CONFLICT
	XML_RNGP_ATTRIBUTE_CHILDREN
	XML_RNGP_ATTRIBUTE_CONTENT
	XML_RNGP_ATTRIBUTE_EMPTY
	XML_RNGP_ATTRIBUTE_NOOP
	XML_RNGP_CHOICE_CONTENT
	XML_RNGP_CHOICE_EMPTY
	XML_RNGP_CREATE_FAILURE
	XML_RNGP_DATA_CONTENT
	XML_RNGP_DEF_CHOICE_AND_INTERLEAVE
	XML_RNGP_DEFINE_CREATE_FAILED
	XML_RNGP_DEFINE_EMPTY
	XML_RNGP_DEFINE_MISSING
	XML_RNGP_DEFINE_NAME_MISSING
	XML_RNGP_ELEM_CONTENT_EMPTY
	XML_RNGP_ELEM_CONTENT_ERROR
	XML_RNGP_ELEMENT_EMPTY
	XML_RNGP_ELEMENT_CONTENT
	XML_RNGP_ELEMENT_NAME
	XML_RNGP_ELEMENT_NO_CONTENT
	XML_RNGP_ELEM_TEXT_CONFLICT
	XML_RNGP_EMPTY
	XML_RNGP_EMPTY_CONSTRUCT
	XML_RNGP_EMPTY_CONTENT
	XML_RNGP_EMPTY_NOT_EMPTY
	XML_RNGP_ERROR_TYPE_LIB
	XML_RNGP_EXCEPT_EMPTY
	XML_RNGP_EXCEPT_MISSING
	XML_RNGP_EXCEPT_MULTIPLE
	XML_RNGP_EXCEPT_NO_CONTENT
	XML_RNGP_EXTERNALREF_EMTPY
	XML_RNGP_EXTERNAL_REF_FAILURE
	XML_RNGP_EXTERNALREF_RECURSE
	XML_RNGP_FORBIDDEN_ATTRIBUTE
	XML_RNGP_FOREIGN_ELEMENT
	XML_RNGP_GRAMMAR_CONTENT
	XML_RNGP_GRAMMAR_EMPTY
	XML_RNGP_GRAMMAR_MISSING
	XML_RNGP_GRAMMAR_NO_START
	XML_RNGP_GROUP_ATTR_CONFLICT
	XML_RNGP_HREF_ERROR
	XML_RNGP_INCLUDE_EMPTY
	XML_RNGP_INCLUDE_FAILURE
	XML_RNGP_INCLUDE_RECURSE
	XML_RNGP_INTERLEAVE_ADD
	XML_RNGP_INTERLEAVE_CREATE_FAILED
	XML_RNGP_INTERLEAVE_EMPTY
	XML_RNGP_INTERLEAVE_NO_CONTENT
	XML_RNGP_INVALID_DEFINE_NAME
	XML_RNGP_INVALID_URI
	XML_RNGP_INVALID_VALUE
	XML_RNGP_MISSING_HREF
	XML_RNGP_NAME_MISSING
	XML_RNGP_NEED_COMBINE
	XML_RNGP_NOTALLOWED_NOT_EMPTY
	XML_RNGP_NSNAME_ATTR_ANCESTOR
	XML_RNGP_NSNAME_NO_NS
	XML_RNGP_PARAM_FORBIDDEN
	XML_RNGP_PARAM_NAME_MISSING
	XML_RNGP_PARENTREF_CREATE_FAILED
	XML_RNGP_PARENTREF_NAME_INVALID
	XML_RNGP_PARENTREF_NO_NAME
	XML_RNGP_PARENTREF_NO_PARENT
	XML_RNGP_PARENTREF_NOT_EMPTY
	XML_RNGP_PARSE_ERROR
	XML_RNGP_PAT_ANYNAME_EXCEPT_ANYNAME
	XML_RNGP_PAT_ATTR_ATTR
	XML_RNGP_PAT_ATTR_ELEM
	XML_RNGP_PAT_DATA_EXCEPT_ATTR
	XML_RNGP_PAT_DATA_EXCEPT_ELEM
	XML_RNGP_PAT_DATA_EXCEPT_EMPTY
	XML_RNGP_PAT_DATA_EXCEPT_GROUP
	XML_RNGP_PAT_DATA_EXCEPT_INTERLEAVE
	XML_RNGP_PAT_DATA_EXCEPT_LIST
	XML_RNGP_PAT_DATA_EXCEPT_ONEMORE
	XML_RNGP_PAT_DATA_EXCEPT_REF
	XML_RNGP_PAT_DATA_EXCEPT_TEXT
	XML_RNGP_PAT_LIST_ATTR
	XML_RNGP_PAT_LIST_ELEM
	XML_RNGP_PAT_LIST_INTERLEAVE
	XML_RNGP_PAT_LIST_LIST
	XML_RNGP_PAT_LIST_REF
	XML_RNGP_PAT_LIST_TEXT
	XML_RNGP_PAT_NSNAME_EXCEPT_ANYNAME
	XML_RNGP_PAT_NSNAME_EXCEPT_NSNAME
	XML_RNGP_PAT_ONEMORE_GROUP_ATTR
	XML_RNGP_PAT_ONEMORE_INTERLEAVE_ATTR
	XML_RNGP_PAT_START_ATTR
	XML_RNGP_PAT_START_DATA
	XML_RNGP_PAT_START_EMPTY
	XML_RNGP_PAT_START_GROUP
	XML_RNGP_PAT_START_INTERLEAVE
	XML_RNGP_PAT_START_LIST
	XML_RNGP_PAT_START_ONEMORE
	XML_RNGP_PAT_START_TEXT
	XML_RNGP_PAT_START_VALUE
	XML_RNGP_PREFIX_UNDEFINED
	XML_RNGP_REF_CREATE_FAILED
	XML_RNGP_REF_CYCLE
	XML_RNGP_REF_NAME_INVALID
	XML_RNGP_REF_NO_DEF
	XML_RNGP_REF_NO_NAME
	XML_RNGP_REF_NOT_EMPTY
	XML_RNGP_START_CHOICE_AND_INTERLEAVE
	XML_RNGP_START_CONTENT
	XML_RNGP_START_EMPTY
	XML_RNGP_START_MISSING
	XML_RNGP_TEXT_EXPECTED
	XML_RNGP_TEXT_HAS_CHILD
	XML_RNGP_TYPE_MISSING
	XML_RNGP_TYPE_NOT_FOUND
	XML_RNGP_TYPE_VALUE
	XML_RNGP_UNKNOWN_ATTRIBUTE
	XML_RNGP_UNKNOWN_COMBINE
	XML_RNGP_UNKNOWN_CONSTRUCT
	XML_RNGP_UNKNOWN_TYPE_LIB
	XML_RNGP_URI_FRAGMENT
	XML_RNGP_URI_NOT_ABSOLUTE
	XML_RNGP_VALUE_EMPTY
	XML_RNGP_VALUE_NO_CONTENT
	XML_RNGP_XmlNs_NAME
	XML_RNGP_XML_NS
)
const (
	XML_XPATH_EXPRESSION_OK XmlParserErrors = iota + 1200
	XML_XPATH_NUMBER_ERROR
	XML_XPATH_UNFINISHED_LITERAL_ERROR
	XML_XPATH_START_LITERAL_ERROR
	XML_XPATH_VARIABLE_REF_ERROR
	XML_XPATH_UNDEF_VARIABLE_ERROR
	XML_XPATH_INVALID_PREDICATE_ERROR
	XML_XPATH_EXPR_ERROR
	XML_XPATH_UNCLOSED_ERROR
	XML_XPATH_UNKNOWN_FUNC_ERROR
	XML_XPATH_INVALID_OPERAND
	XML_XPATH_INVALID_TYPE
	XML_XPATH_INVALID_ARITY
	XML_XPATH_INVALID_CTXT_SIZE
	XML_XPATH_INVALID_CTXT_POSITION
	XML_XPATH_MEMORY_ERROR
	XML_XPTR_SYNTAX_ERROR
	XML_XPTR_RESOURCE_ERROR
	XML_XPTR_SUB_RESOURCE_ERROR
	XML_XPATH_UNDEF_PREFIX_ERROR
	XML_XPATH_ENCODING_ERROR
	XML_XPATH_INVALID_CHAR_ERROR
)
const (
	XML_TREE_INVALID_HEX XmlParserErrors = iota + 1300
	XML_TREE_INVALID_DEC
	XML_TREE_UNTERMINATED_ENTITY
	XML_TREE_NOT_UTF8
)
const (
	XML_SAVE_NOT_UTF8 XmlParserErrors = iota + 1400
	XML_SAVE_CHAR_INVALID
	XML_SAVE_NO_DOCTYPE
	XML_SAVE_UNKNOWN_ENCODING
	XML_REGEXP_COMPILE_ERROR XmlParserErrors = 1450
)
const (
	XML_IO_UNKNOWN XmlParserErrors = iota + 1500
	XML_IO_EACCES
	XML_IO_EAGAIN
	XML_IO_EBADF
	XML_IO_EBADMSG
	XML_IO_EBUSY
	XML_IO_ECANCELED
	XML_IO_ECHILD
	XML_IO_EDEADLK
	XML_IO_EDOM
	XML_IO_EEXIST
	XML_IO_EFAULT
	XML_IO_EFBIG
	XML_IO_EINPROGRESS
	XML_IO_EINTR
	XML_IO_EINVAL
	XML_IO_EIO
	XML_IO_EISDIR
	XML_IO_EMFILE
	XML_IO_EMLINK
	XML_IO_EMSGSIZE
	XML_IO_ENAMETOOLONG
	XML_IO_ENFILE
	XML_IO_ENODEV
	XML_IO_ENOENT
	XML_IO_ENOEXEC
	XML_IO_ENOLCK
	XML_IO_ENOMEM
	XML_IO_ENOSPC
	XML_IO_ENOSYS
	XML_IO_ENOTDIR
	XML_IO_ENOTEMPTY
	XML_IO_ENOTSUP
	XML_IO_ENOTTY
	XML_IO_ENXIO
	XML_IO_EPERM
	XML_IO_EPIPE
	XML_IO_ERANGE
	XML_IO_EROFS
	XML_IO_ESPIPE
	XML_IO_ESRCH
	XML_IO_ETIMEDOUT
	XML_IO_EXDEV
	XML_IO_NETWORK_ATTEMPT
	XML_IO_ENCODER
	XML_IO_FLUSH
	XML_IO_WRITE
	XML_IO_NO_INPUT
	XML_IO_BUFFER_FULL
	XML_IO_LOAD_ERROR
	XML_IO_ENOTSOCK
	XML_IO_EISCONN
	XML_IO_ECONNREFUSED
	XML_IO_ENETUNREACH
	XML_IO_EADDRINUSE
	XML_IO_EALREADY
	XML_IO_EAFNOSUPPORT
)
const (
	XML_XINCLUDE_RECURSION XmlParserErrors = iota + 1600
	XML_XINCLUDE_PARSE_VALUE
	XML_XINCLUDE_ENTITY_DEF_MISMATCH
	XML_XINCLUDE_NO_HREF
	XML_XINCLUDE_NO_FALLBACK
	XML_XINCLUDE_HREF_URI
	XML_XINCLUDE_TEXT_FRAGMENT
	XML_XINCLUDE_TEXT_DOCUMENT
	XML_XINCLUDE_INVALID_CHAR
	XML_XINCLUDE_BUILD_FAILED
	XML_XINCLUDE_UNKNOWN_ENCODING
	XML_XINCLUDE_MULTIPLE_ROOT
	XML_XINCLUDE_XPTR_FAILED
	XML_XINCLUDE_XPTR_RESULT
	XML_XINCLUDE_INCLUDE_IN_INCLUDE
	XML_XINCLUDE_FALLBACKS_IN_INCLUDE
	XML_XINCLUDE_FALLBACK_NOT_IN_INCLUDE
	XML_XINCLUDE_DEPRECATED_NS
	XML_XINCLUDE_FRAGMENT_ID
)
const (
	XML_CATALOG_MISSING_ATTR XmlParserErrors = iota + 1650
	XML_CATALOG_ENTRY_BROKEN
	XML_CATALOG_PREFER_VALUE
	XML_CATALOG_NOT_CATALOG
	XML_CATALOG_RECURSION
)
const (
	XML_SCHEMAP_PREFIX_UNDEFINED XmlParserErrors = iota + 1700
	XML_SCHEMAP_ATTRFORMDEFAULT_VALUE
	XML_SCHEMAP_ATTRGRP_NONAME_NOREF
	XML_SCHEMAP_ATTR_NONAME_NOREF
	XML_SCHEMAP_COMPLEXTYPE_NONAME_NOREF
	XML_SCHEMAP_ELEMFORMDEFAULT_VALUE
	XML_SCHEMAP_ELEM_NONAME_NOREF
	XML_SCHEMAP_EXTENSION_NO_BASE
	XML_SCHEMAP_FACET_NO_VALUE
	XML_SCHEMAP_FAILED_BUILD_IMPORT
	XML_SCHEMAP_GROUP_NONAME_NOREF
	XML_SCHEMAP_IMPORT_NAMESPACE_NOT_URI
	XML_SCHEMAP_IMPORT_REDEFINE_NSNAME
	XML_SCHEMAP_IMPORT_SCHEMA_NOT_URI
	XML_SCHEMAP_INVALID_BOOLEAN
	XML_SCHEMAP_INVALID_ENUM
	XML_SCHEMAP_INVALID_FACET
	XML_SCHEMAP_INVALID_FACET_VALUE
	XML_SCHEMAP_INVALID_MAXOCCURS
	XML_SCHEMAP_INVALID_MINOCCURS
	XML_SCHEMAP_INVALID_REF_AND_SUBTYPE
	XML_SCHEMAP_INVALID_WHITE_SPACE
	XML_SCHEMAP_NOATTR_NOREF
	XML_SCHEMAP_NOTATION_NO_NAME
	XML_SCHEMAP_NOTYPE_NOREF
	XML_SCHEMAP_REF_AND_SUBTYPE
	XML_SCHEMAP_RESTRICTION_NONAME_NOREF
	XML_SCHEMAP_SIMPLETYPE_NONAME
	XML_SCHEMAP_TYPE_AND_SUBTYPE
	XML_SCHEMAP_UNKNOWN_ALL_CHILD
	XML_SCHEMAP_UNKNOWN_ANYATTRIBUTE_CHILD
	XML_SCHEMAP_UNKNOWN_ATTR_CHILD
	XML_SCHEMAP_UNKNOWN_ATTRGRP_CHILD
	XML_SCHEMAP_UNKNOWN_ATTRIBUTE_GROUP
	XML_SCHEMAP_UNKNOWN_BASE_TYPE
	XML_SCHEMAP_UNKNOWN_CHOICE_CHILD
	XML_SCHEMAP_UNKNOWN_COMPLEXCONTENT_CHILD
	XML_SCHEMAP_UNKNOWN_COMPLEXTYPE_CHILD
	XML_SCHEMAP_UNKNOWN_ELEM_CHILD
	XML_SCHEMAP_UNKNOWN_EXTENSION_CHILD
	XML_SCHEMAP_UNKNOWN_FACET_CHILD
	XML_SCHEMAP_UNKNOWN_FACET_TYPE
	XML_SCHEMAP_UNKNOWN_GROUP_CHILD
	XML_SCHEMAP_UNKNOWN_IMPORT_CHILD
	XML_SCHEMAP_UNKNOWN_LIST_CHILD
	XML_SCHEMAP_UNKNOWN_NOTATION_CHILD
	XML_SCHEMAP_UNKNOWN_PROCESSCONTENT_CHILD
	XML_SCHEMAP_UNKNOWN_REF
	XML_SCHEMAP_UNKNOWN_RESTRICTION_CHILD
	XML_SCHEMAP_UNKNOWN_SCHEMAS_CHILD
	XML_SCHEMAP_UNKNOWN_SEQUENCE_CHILD
	XML_SCHEMAP_UNKNOWN_SIMPLECONTENT_CHILD
	XML_SCHEMAP_UNKNOWN_SIMPLETYPE_CHILD
	XML_SCHEMAP_UNKNOWN_TYPE
	XML_SCHEMAP_UNKNOWN_UNION_CHILD
	XML_SCHEMAP_ELEM_DEFAULT_FIXED
	XML_SCHEMAP_REGEXP_INVALID
	XML_SCHEMAP_FAILED_LOAD
	XML_SCHEMAP_NOTHING_TO_PARSE
	XML_SCHEMAP_NOROOT
	XML_SCHEMAP_REDEFINED_GROUP
	XML_SCHEMAP_REDEFINED_TYPE
	XML_SCHEMAP_REDEFINED_ELEMENT
	XML_SCHEMAP_REDEFINED_ATTRGROUP
	XML_SCHEMAP_REDEFINED_ATTR
	XML_SCHEMAP_REDEFINED_NOTATION
	XML_SCHEMAP_FAILED_PARSE
	XML_SCHEMAP_UNKNOWN_PREFIX
	XML_SCHEMAP_DEF_AND_PREFIX
	XML_SCHEMAP_UNKNOWN_INCLUDE_CHILD
	XML_SCHEMAP_INCLUDE_SCHEMA_NOT_URI
	XML_SCHEMAP_INCLUDE_SCHEMA_NO_URI
	XML_SCHEMAP_NOT_SCHEMA
	XML_SCHEMAP_UNKNOWN_MEMBER_TYPE
	XML_SCHEMAP_INVALID_ATTR_USE
	XML_SCHEMAP_RECURSIVE
	XML_SCHEMAP_SUPERNUMEROUS_LIST_ITEM_TYPE
	XML_SCHEMAP_INVALID_ATTR_COMBINATION
	XML_SCHEMAP_INVALID_ATTR_INLINE_COMBINATION
	XML_SCHEMAP_MISSING_SIMPLETYPE_CHILD
	XML_SCHEMAP_INVALID_ATTR_NAME
	XML_SCHEMAP_REF_AND_CONTENT
	XML_SCHEMAP_CT_PROPS_CORRECT_1
	XML_SCHEMAP_CT_PROPS_CORRECT_2
	XML_SCHEMAP_CT_PROPS_CORRECT_3
	XML_SCHEMAP_CT_PROPS_CORRECT_4
	XML_SCHEMAP_CT_PROPS_CORRECT_5
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_1
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_1
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_2
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_2_2
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_3
	XML_SCHEMAP_WILDCARD_INVALID_NS_MEMBER
	XML_SCHEMAP_INTERSECTION_NOT_EXPRESSIBLE
	XML_SCHEMAP_UNION_NOT_EXPRESSIBLE
	XML_SCHEMAP_SRC_IMPORT_3_1
	XML_SCHEMAP_SRC_IMPORT_3_2
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_4_1
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_4_2
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_4_3
	XML_SCHEMAP_COS_CT_EXTENDS_1_3
)
const (
	XML_SCHEMAV_NOROOT XmlParserErrors = iota + 1801
	XML_SCHEMAV_UNDECLAREDELEM
	XML_SCHEMAV_NOTTOPLEVEL
	XML_SCHEMAV_MISSING
	XML_SCHEMAV_WRONGELEM
	XML_SCHEMAV_NOTYPE
	XML_SCHEMAV_NOROLLBACK
	XML_SCHEMAV_ISABSTRACT
	XML_SCHEMAV_NOTEMPTY
	XML_SCHEMAV_ELEMCONT
	XML_SCHEMAV_HAVEDEFAULT
	XML_SCHEMAV_NOTNILLABLE
	XML_SCHEMAV_EXTRACONTENT
	XML_SCHEMAV_INVALIDATTR
	XML_SCHEMAV_INVALIDELEM
	XML_SCHEMAV_NOTDETERMINIST
	XML_SCHEMAV_CONSTRUCT
	XML_SCHEMAV_INTERNAL
	XML_SCHEMAV_NOTSIMPLE
	XML_SCHEMAV_ATTRUNKNOWN
	XML_SCHEMAV_ATTRINVALID
	XML_SCHEMAV_VALUE
	XML_SCHEMAV_FACET
	XML_SCHEMAV_CVC_DATATYPE_VALID_1_2_1
	XML_SCHEMAV_CVC_DATATYPE_VALID_1_2_2
	XML_SCHEMAV_CVC_DATATYPE_VALID_1_2_3
	XML_SCHEMAV_CVC_TYPE_3_1_1
	XML_SCHEMAV_CVC_TYPE_3_1_2
	XML_SCHEMAV_CVC_FACET_VALID
	XML_SCHEMAV_CVC_LENGTH_VALID
	XML_SCHEMAV_CVC_MINLENGTH_VALID
	XML_SCHEMAV_CVC_MAXLENGTH_VALID
	XML_SCHEMAV_CVC_MININCLUSIVE_VALID
	XML_SCHEMAV_CVC_MAXINCLUSIVE_VALID
	XML_SCHEMAV_CVC_MINEXCLUSIVE_VALID
	XML_SCHEMAV_CVC_MAXEXCLUSIVE_VALID
	XML_SCHEMAV_CVC_TOTALDIGITS_VALID
	XML_SCHEMAV_CVC_FRACTIONDIGITS_VALID
	XML_SCHEMAV_CVC_PATTERN_VALID
	XML_SCHEMAV_CVC_ENUMERATION_VALID
	XML_SCHEMAV_CVC_COMPLEX_TYPE_2_1
	XML_SCHEMAV_CVC_COMPLEX_TYPE_2_2
	XML_SCHEMAV_CVC_COMPLEX_TYPE_2_3
	XML_SCHEMAV_CVC_COMPLEX_TYPE_2_4
	XML_SCHEMAV_CVC_ELT_1
	XML_SCHEMAV_CVC_ELT_2
	XML_SCHEMAV_CVC_ELT_3_1
	XML_SCHEMAV_CVC_ELT_3_2_1
	XML_SCHEMAV_CVC_ELT_3_2_2
	XML_SCHEMAV_CVC_ELT_4_1
	XML_SCHEMAV_CVC_ELT_4_2
	XML_SCHEMAV_CVC_ELT_4_3
	XML_SCHEMAV_CVC_ELT_5_1_1
	XML_SCHEMAV_CVC_ELT_5_1_2
	XML_SCHEMAV_CVC_ELT_5_2_1
	XML_SCHEMAV_CVC_ELT_5_2_2_1
	XML_SCHEMAV_CVC_ELT_5_2_2_2_1
	XML_SCHEMAV_CVC_ELT_5_2_2_2_2
	XML_SCHEMAV_CVC_ELT_6
	XML_SCHEMAV_CVC_ELT_7
	XML_SCHEMAV_CVC_ATTRIBUTE_1
	XML_SCHEMAV_CVC_ATTRIBUTE_2
	XML_SCHEMAV_CVC_ATTRIBUTE_3
	XML_SCHEMAV_CVC_ATTRIBUTE_4
	XML_SCHEMAV_CVC_COMPLEX_TYPE_3_1
	XML_SCHEMAV_CVC_COMPLEX_TYPE_3_2_1
	XML_SCHEMAV_CVC_COMPLEX_TYPE_3_2_2
	XML_SCHEMAV_CVC_COMPLEX_TYPE_4
	XML_SCHEMAV_CVC_COMPLEX_TYPE_5_1
	XML_SCHEMAV_CVC_COMPLEX_TYPE_5_2
	XML_SCHEMAV_ELEMENT_CONTENT
	XML_SCHEMAV_DOCUMENT_ELEMENT_MISSING
	XML_SCHEMAV_CVC_COMPLEX_TYPE_1
	XML_SCHEMAV_CVC_AU
	XML_SCHEMAV_CVC_TYPE_1
	XML_SCHEMAV_CVC_TYPE_2
	XML_SCHEMAV_CVC_IDC
	XML_SCHEMAV_CVC_WILDCARD
	XML_SCHEMAV_MISC
)
const (
	XML_XPTR_UNKNOWN_SCHEME XmlParserErrors = iota + 1900
	XML_XPTR_CHILDSEQ_START
	XML_XPTR_EVAL_FAILED
	XML_XPTR_EXTRA_OBJECTS
)
const (
	XML_C14N_CREATE_CTXT XmlParserErrors = iota + 1950
	XML_C14N_REQUIRES_UTF8
	XML_C14N_CREATE_STACK
	XML_C14N_INVALID_NODE
	XML_C14N_UNKNOW_NODE
	XML_C14N_RELATIVE_NAMESPACE
)
const (
	XML_FTP_PASV_ANSWER XmlParserErrors = iota + 2000
	XML_FTP_EPSV_ANSWER
	XML_FTP_ACCNT
	XML_FTP_URL_SYNTAX
)
const (
	XML_HTTP_URL_SYNTAX XmlParserErrors = iota + 2020
	XML_HTTP_USE_IP
	XML_HTTP_UNKNOWN_HOST
)
const (
	XML_SCHEMAP_SRC_SIMPLE_TYPE_1 XmlParserErrors = iota + 3000
	XML_SCHEMAP_SRC_SIMPLE_TYPE_2
	XML_SCHEMAP_SRC_SIMPLE_TYPE_3
	XML_SCHEMAP_SRC_SIMPLE_TYPE_4
	XML_SCHEMAP_SRC_RESOLVE
	XML_SCHEMAP_SRC_RESTRICTION_BASE_OR_SIMPLETYPE
	XML_SCHEMAP_SRC_LIST_ITEMTYPE_OR_SIMPLETYPE
	XML_SCHEMAP_SRC_UNION_MEMBERTYPES_OR_SIMPLETYPES
	XML_SCHEMAP_ST_PROPS_CORRECT_1
	XML_SCHEMAP_ST_PROPS_CORRECT_2
	XML_SCHEMAP_ST_PROPS_CORRECT_3
	XML_SCHEMAP_COS_ST_RESTRICTS_1_1
	XML_SCHEMAP_COS_ST_RESTRICTS_1_2
	XML_SCHEMAP_COS_ST_RESTRICTS_1_3_1
	XML_SCHEMAP_COS_ST_RESTRICTS_1_3_2
	XML_SCHEMAP_COS_ST_RESTRICTS_2_1
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_1_1
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_1_2
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_1
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_2
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_3
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_4
	XML_SCHEMAP_COS_ST_RESTRICTS_2_3_2_5
	XML_SCHEMAP_COS_ST_RESTRICTS_3_1
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_1
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_1_2
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_2
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_1
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_3
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_4
	XML_SCHEMAP_COS_ST_RESTRICTS_3_3_2_5
	XML_SCHEMAP_COS_ST_DERIVED_OK_2_1
	XML_SCHEMAP_COS_ST_DERIVED_OK_2_2
	XML_SCHEMAP_S4S_ELEM_NOT_ALLOWED
	XML_SCHEMAP_S4S_ELEM_MISSING
	XML_SCHEMAP_S4S_ATTR_NOT_ALLOWED
	XML_SCHEMAP_S4S_ATTR_MISSING
	XML_SCHEMAP_S4S_ATTR_INVALID_VALUE
	XML_SCHEMAP_SRC_ELEMENT_1
	XML_SCHEMAP_SRC_ELEMENT_2_1
	XML_SCHEMAP_SRC_ELEMENT_2_2
	XML_SCHEMAP_SRC_ELEMENT_3
	XML_SCHEMAP_P_PROPS_CORRECT_1
	XML_SCHEMAP_P_PROPS_CORRECT_2_1
	XML_SCHEMAP_P_PROPS_CORRECT_2_2
	XML_SCHEMAP_E_PROPS_CORRECT_2
	XML_SCHEMAP_E_PROPS_CORRECT_3
	XML_SCHEMAP_E_PROPS_CORRECT_4
	XML_SCHEMAP_E_PROPS_CORRECT_5
	XML_SCHEMAP_E_PROPS_CORRECT_6
	XML_SCHEMAP_SRC_INCLUDE
	XML_SCHEMAP_SRC_ATTRIBUTE_1
	XML_SCHEMAP_SRC_ATTRIBUTE_2
	XML_SCHEMAP_SRC_ATTRIBUTE_3_1
	XML_SCHEMAP_SRC_ATTRIBUTE_3_2
	XML_SCHEMAP_SRC_ATTRIBUTE_4
	XML_SCHEMAP_NO_XmlNs
	XML_SCHEMAP_NO_XSI
	XML_SCHEMAP_COS_VALID_DEFAULT_1
	XML_SCHEMAP_COS_VALID_DEFAULT_2_1
	XML_SCHEMAP_COS_VALID_DEFAULT_2_2_1
	XML_SCHEMAP_COS_VALID_DEFAULT_2_2_2
	XML_SCHEMAP_CVC_SIMPLE_TYPE
	XML_SCHEMAP_COS_CT_EXTENDS_1_1
	XML_SCHEMAP_SRC_IMPORT_1_1
	XML_SCHEMAP_SRC_IMPORT_1_2
	XML_SCHEMAP_SRC_IMPORT_2
	XML_SCHEMAP_SRC_IMPORT_2_1
	XML_SCHEMAP_SRC_IMPORT_2_2
	XML_SCHEMAP_INTERNAL
	XML_SCHEMAP_NOT_DETERMINISTIC
	XML_SCHEMAP_SRC_ATTRIBUTE_GROUP_1
	XML_SCHEMAP_SRC_ATTRIBUTE_GROUP_2
	XML_SCHEMAP_SRC_ATTRIBUTE_GROUP_3
	XML_SCHEMAP_MG_PROPS_CORRECT_1
	XML_SCHEMAP_MG_PROPS_CORRECT_2
	XML_SCHEMAP_SRC_CT_1
	XML_SCHEMAP_DERIVATION_OK_RESTRICTION_2_1_3
	XML_SCHEMAP_AU_PROPS_CORRECT_2
	XML_SCHEMAP_A_PROPS_CORRECT_2
	XML_SCHEMAP_C_PROPS_CORRECT
	XML_SCHEMAP_SRC_REDEFINE
	XML_SCHEMAP_SRC_IMPORT
	XML_SCHEMAP_WARN_SKIP_SCHEMA
	XML_SCHEMAP_WARN_UNLOCATED_SCHEMA
	XML_SCHEMAP_WARN_ATTR_REDECL_PROH
	XML_SCHEMAP_WARN_ATTR_POINTLESS_PROH
	XML_SCHEMAP_AG_PROPS_CORRECT
	XML_SCHEMAP_COS_CT_EXTENDS_1_2
	XML_SCHEMAP_AU_PROPS_CORRECT
	XML_SCHEMAP_A_PROPS_CORRECT_3
	XML_SCHEMAP_COS_ALL_LIMITED
)
const (
	XML_SCHEMATRONV_ASSERT XmlParserErrors = iota + 4000
	XML_SCHEMATRONV_REPORT
)
const (
	XML_MODULE_OPEN XmlParserErrors = iota + 4900
	XML_MODULE_CLOSE
)
const (
	XML_CHECK_FOUND_ELEMENT XmlParserErrors = iota + 5000
	XML_CHECK_FOUND_ATTRIBUTE
	XML_CHECK_FOUND_TEXT
	XML_CHECK_FOUND_CDATA
	XML_CHECK_FOUND_ENTITYREF
	XML_CHECK_FOUND_ENTITY
	XML_CHECK_FOUND_PI
	XML_CHECK_FOUND_COMMENT
	XML_CHECK_FOUND_DOCTYPE
	XML_CHECK_FOUND_FRAGMENT
	XML_CHECK_FOUND_NOTATION
	XML_CHECK_UNKNOWN_NODE
	XML_CHECK_ENTITY_TYPE
	XML_CHECK_NO_PARENT
	XML_CHECK_NO_DOC
	XML_CHECK_NO_NAME
	XML_CHECK_NO_ELEM
	XML_CHECK_WRONG_DOC
	XML_CHECK_NO_PREV
	XML_CHECK_WRONG_PREV
	XML_CHECK_NO_NEXT
	XML_CHECK_WRONG_NEXT
	XML_CHECK_NOT_DTD
	XML_CHECK_NOT_ATTR
	XML_CHECK_NOT_ATTR_DECL
	XML_CHECK_NOT_ELEM_DECL
	XML_CHECK_NOT_ENTITY_DECL
	XML_CHECK_NOT_NS_DECL
	XML_CHECK_NO_HREF
	XML_CHECK_WRONG_PARENT
	XML_CHECK_NS_SCOPE
	XML_CHECK_NS_ANCESTOR
	XML_CHECK_NOT_UTF8
	XML_CHECK_NO_DICT
	XML_CHECK_NOT_NCNAME
	XML_CHECK_OUTSIDE_DICT
	XML_CHECK_WRONG_NAME
	XML_CHECK_NAME_NOT_NULL
)
const (
	XML_I18N_NO_NAME XmlParserErrors = iota + 6000
	XML_I18N_NO_HANDLER
	XML_I18N_EXCESS_HANDLER
	XML_I18N_CONV_FAILED
	XML_I18N_NO_OUTPUT
	XML_BUF_OVERFLOW XmlParserErrors = 7000
)

var (
	XmlCheckVersion func(version int)

	XmlStrdup func(cur *XmlChar) *XmlChar

	XmlStrndup func(cur *XmlChar, leng int) *XmlChar

	XmlCharStrndup func(cur *Char, leng int) *XmlChar

	XmlCharStrdup func(cur *Char) *XmlChar

	XmlStrsub func(str *XmlChar, start int, leng int) *XmlChar

	XmlStrchr func(str *XmlChar, val XmlChar) *XmlChar

	XmlStrstr func(str, val *XmlChar) *XmlChar

	XmlStrcasestr func(str, val *XmlChar) *XmlChar

	XmlStrcmp func(str1, str2 *XmlChar) int

	XmlStrncmp func(str1, str2 *XmlChar, leng int) int

	XmlStrcasecmp func(str1, str2 *XmlChar) int

	XmlStrncasecmp func(str1, str2 *XmlChar, leng int) int

	XmlStrEqual func(str1, str2 *XmlChar) int

	XmlStrQEqual func(pref, name, str *XmlChar) int

	XmlStrlen func(str *XmlChar) int

	XmlStrcat func(cur, add *XmlChar) *XmlChar

	XmlStrncat func(cur, add *XmlChar, leng int) *XmlChar

	XmlStrncatNew func(str1, str2 *XmlChar, leng int) *XmlChar

	XmlStrPrintf func(
		buf *XmlChar, leng int, msg *XmlChar, v ...VArg) int

	XmlStrVPrintf func(
		buf *XmlChar, leng int, msg *XmlChar, ap Va_list) int

	XmlGetUTF8Char func(utf *Unsigned_char, leng *int) int

	XmlCheckUTF8 func(utf *Unsigned_char) int

	XmlUTF8Strsize func(utf *XmlChar, leng int) int

	XmlUTF8Strndup func(utf *XmlChar, leng int) *XmlChar

	XmlUTF8Strpos func(utf *XmlChar, pos int) *XmlChar

	XmlUTF8Strloc func(utf *XmlChar, utfchar *XmlChar) int

	XmlUTF8Strsub func(utf *XmlChar, start, leng int) *XmlChar

	XmlUTF8Strlen func(
		utf *XmlChar) int

	XmlUTF8Size func(
		utf *XmlChar) int

	XmlUTF8Charcmp func(
		utf1 *XmlChar,
		utf2 *XmlChar) int

	XmlBufContent func(
		buf XmlBufPtr) *XmlChar

	XmlBufEnd func(
		buf XmlBufPtr) *XmlChar

	XmlBufUse func(
		buf XmlBufPtr) Size_t

	XmlBufShrink func(
		buf XmlBufPtr,
		leng Size_t) Size_t

	XmlInitializeDict func() int

	XmlDictCreate func() XmlDictPtr

	XmlDictSetLimit func(dict XmlDictPtr, limit Size_t) Size_t

	XmlDictGetUsage func(dict XmlDictPtr) Size_t

	XmlDictCreateSub func(sub XmlDictPtr) XmlDictPtr

	XmlDictReference func(dict XmlDictPtr) int

	XmlDictFree func(dict XmlDictPtr)

	XmlDictLookup func(
		dict XmlDictPtr, name *XmlChar, leng int) *XmlChar

	XmlDictExists func(
		dict XmlDictPtr, name *XmlChar, leng int) *XmlChar

	XmlDictQLookup func(
		dict XmlDictPtr, prefix, name *XmlChar) *XmlChar

	XmlDictOwns func(dict XmlDictPtr, str *XmlChar) int

	XmlDictSize func(dict XmlDictPtr) int

	XmlDictCleanup func()

	XmlRegexpCompile func(regexp *XmlChar) XmlRegexpPtr

	XmlRegFreeRegexp func(
		regexp XmlRegexpPtr)

	XmlRegexpExec func(
		comp XmlRegexpPtr,
		value *XmlChar) int

	XmlRegexpPrint func(
		output *FILE,
		regexp XmlRegexpPtr)

	XmlRegexpIsDeterminist func(
		comp XmlRegexpPtr) int

	XmlRegNewExecCtxt func(
		comp XmlRegexpPtr,
		callback XmlRegExecCallbacks,
		data *Void) XmlRegExecCtxtPtr

	XmlRegFreeExecCtxt func(
		exec XmlRegExecCtxtPtr)

	XmlRegExecPushString func(
		exec XmlRegExecCtxtPtr,
		value *XmlChar,
		data *Void) int

	XmlRegExecPushString2 func(
		exec XmlRegExecCtxtPtr,
		value *XmlChar,
		value2 *XmlChar,
		data *Void) int

	XmlRegExecNextValues func(
		exec XmlRegExecCtxtPtr,
		nbval *int,
		nbneg *int,
		values **XmlChar,
		terminal *int) int

	XmlRegExecErrInfo func(
		exec XmlRegExecCtxtPtr,
		string **XmlChar,
		nbval *int,
		nbneg *int,
		values **XmlChar,
		terminal *int) int

	XmlExpFreeCtxt func(ctxt XmlExpCtxtPtr)

	XmlExpNewCtxt func(
		maxNodes int, dict XmlDictPtr) XmlExpCtxtPtr

	XmlExpCtxtNbNodes func(ctxt XmlExpCtxtPtr) int

	XmlExpCtxtNbCons func(ctxt XmlExpCtxtPtr) int

	XmlExpFree func(ctxt XmlExpCtxtPtr, expr XmlExpNodePtr)

	XmlExpRef func(expr XmlExpNodePtr)

	XmlExpParse func(
		ctxt XmlExpCtxtPtr, expr *Char) XmlExpNodePtr

	XmlExpNewAtom func(ctxt XmlExpCtxtPtr,
		name *XmlChar, leng int) XmlExpNodePtr

	XmlExpNewOr func(
		ctxt XmlExpCtxtPtr,
		left XmlExpNodePtr,
		right XmlExpNodePtr) XmlExpNodePtr

	XmlExpNewSeq func(
		ctxt XmlExpCtxtPtr,
		left XmlExpNodePtr,
		right XmlExpNodePtr) XmlExpNodePtr

	XmlExpNewRange func(
		ctxt XmlExpCtxtPtr,
		subset XmlExpNodePtr,
		min int,
		max int) XmlExpNodePtr

	XmlExpIsNillable func(
		expr XmlExpNodePtr) int

	XmlExpMaxToken func(
		expr XmlExpNodePtr) int

	XmlExpGetLanguage func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		langList **XmlChar,
		leng int) int

	XmlExpGetStart func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		tokList **XmlChar,
		leng int) int

	XmlExpStringDerive func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		str *XmlChar,
		leng int) XmlExpNodePtr

	XmlExpExpDerive func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		sub XmlExpNodePtr) XmlExpNodePtr

	XmlExpSubsume func(
		ctxt XmlExpCtxtPtr,
		expr XmlExpNodePtr,
		sub XmlExpNodePtr) int

	XmlExpDump func(
		buf XmlBufferPtr,
		expr XmlExpNodePtr)

	XmlValidateNCName func(
		value *XmlChar,
		space int) int

	XmlValidateQName func(
		value *XmlChar,
		space int) int

	XmlValidateName func(
		value *XmlChar,
		space int) int

	XmlValidateNMToken func(
		value *XmlChar,
		space int) int

	XmlBuildQName func(
		ncname *XmlChar,
		prefix *XmlChar,
		memory *XmlChar,
		leng int) *XmlChar

	XmlSplitQName2 func(
		name *XmlChar,
		prefix **XmlChar) *XmlChar

	XmlSplitQName3 func(
		name *XmlChar,
		leng *int) *XmlChar

	XmlSetBufferAllocationScheme func(
		scheme XmlBufferAllocationScheme)

	XmlGetBufferAllocationScheme func() XmlBufferAllocationScheme

	XmlBufferCreate func() XmlBufferPtr

	XmlBufferCreateSize func(size Size_t) XmlBufferPtr

	XmlBufferCreateStatic func(
		mem *Void, size Size_t) XmlBufferPtr

	XmlBufferResize func(buf XmlBufferPtr, size Unsigned_int) int

	XmlBufferFree func(buf XmlBufferPtr)

	XmlBufferDump func(
		file *FILE,
		buf XmlBufferPtr) int

	XmlBufferAdd func(
		buf XmlBufferPtr,
		str *XmlChar,
		leng int) int

	XmlBufferAddHead func(
		buf XmlBufferPtr,
		str *XmlChar,
		leng int) int

	XmlBufferCat func(
		buf XmlBufferPtr,
		str *XmlChar) int

	XmlBufferCCat func(
		buf XmlBufferPtr,
		str *Char) int

	XmlBufferShrink func(
		buf XmlBufferPtr,
		leng Unsigned_int) int

	XmlBufferGrow func(
		buf XmlBufferPtr,
		leng Unsigned_int) int

	XmlBufferEmpty func(
		buf XmlBufferPtr)

	XmlBufferContent func(
		buf XmlBufferPtr) *XmlChar

	XmlBufferDetach func(
		buf XmlBufferPtr) *XmlChar

	XmlBufferSetAllocationScheme func(
		buf XmlBufferPtr,
		scheme XmlBufferAllocationScheme)

	XmlBufferLength func(
		buf XmlBufferPtr) int

	XmlCreateIntSubset func(
		doc XmlDocPtr,
		name *XmlChar,
		ExternalID *XmlChar,
		SystemID *XmlChar) XmlDtdPtr

	XmlNewDtd func(
		doc XmlDocPtr,
		name *XmlChar,
		ExternalID *XmlChar,
		SystemID *XmlChar) XmlDtdPtr

	XmlGetIntSubset func(
		doc XmlDocPtr) XmlDtdPtr

	XmlFreeDtd func(
		cur XmlDtdPtr)

	XmlNewGlobalNs func(
		doc XmlDocPtr,
		href *XmlChar,
		prefix *XmlChar) XmlNsPtr

	XmlNewNs func(
		node XmlNodePtr,
		href *XmlChar,
		prefix *XmlChar) XmlNsPtr

	XmlFreeNs func(
		cur XmlNsPtr)

	XmlFreeNsList func(
		cur XmlNsPtr)

	XmlNewDoc func(
		version *XmlChar) XmlDocPtr

	XmlFreeDoc func(
		cur XmlDocPtr)

	XmlNewDocProp func(
		doc XmlDocPtr,
		name *XmlChar,
		value *XmlChar) XmlAttrPtr

	XmlNewProp func(
		node XmlNodePtr,
		name *XmlChar,
		value *XmlChar) XmlAttrPtr

	XmlNewNsProp func(
		node XmlNodePtr,
		ns XmlNsPtr,
		name *XmlChar,
		value *XmlChar) XmlAttrPtr

	XmlNewNsPropEatName func(
		node XmlNodePtr,
		ns XmlNsPtr,
		name *XmlChar,
		value *XmlChar) XmlAttrPtr

	XmlFreePropList func(
		cur XmlAttrPtr)

	XmlFreeProp func(
		cur XmlAttrPtr)

	XmlCopyProp func(
		target XmlNodePtr,
		cur XmlAttrPtr) XmlAttrPtr

	XmlCopyPropList func(
		target XmlNodePtr,
		cur XmlAttrPtr) XmlAttrPtr

	XmlCopyDtd func(
		dtd XmlDtdPtr) XmlDtdPtr

	XmlCopyDoc func(
		doc XmlDocPtr,
		recursive int) XmlDocPtr

	XmlNewDocNode func(
		doc XmlDocPtr,
		ns XmlNsPtr,
		name *XmlChar,
		content *XmlChar) XmlNodePtr

	XmlNewDocNodeEatName func(
		doc XmlDocPtr,
		ns XmlNsPtr,
		name *XmlChar,
		content *XmlChar) XmlNodePtr

	XmlNewNode func(
		ns XmlNsPtr,
		name *XmlChar) XmlNodePtr

	XmlNewNodeEatName func(
		ns XmlNsPtr,
		name *XmlChar) XmlNodePtr

	XmlNewChild func(
		parent XmlNodePtr,
		ns XmlNsPtr,
		name *XmlChar,
		content *XmlChar) XmlNodePtr

	XmlNewDocText func(
		doc XmlDocPtr,
		content *XmlChar) XmlNodePtr

	XmlNewText func(
		content *XmlChar) XmlNodePtr

	XmlNewDocPI func(
		doc XmlDocPtr,
		name *XmlChar,
		content *XmlChar) XmlNodePtr

	XmlNewPI func(
		name *XmlChar,
		content *XmlChar) XmlNodePtr

	XmlNewDocTextLen func(
		doc XmlDocPtr,
		content *XmlChar,
		leng int) XmlNodePtr

	XmlNewTextLen func(
		content *XmlChar,
		leng int) XmlNodePtr

	XmlNewDocComment func(
		doc XmlDocPtr,
		content *XmlChar) XmlNodePtr

	XmlNewComment func(
		content *XmlChar) XmlNodePtr

	XmlNewCDataBlock func(
		doc XmlDocPtr,
		content *XmlChar,
		leng int) XmlNodePtr

	XmlNewCharRef func(
		doc XmlDocPtr,
		name *XmlChar) XmlNodePtr

	XmlNewReference func(
		doc XmlDocPtr,
		name *XmlChar) XmlNodePtr

	XmlCopyNode func(
		node XmlNodePtr,
		recursive int) XmlNodePtr

	XmlDocCopyNode func(
		node XmlNodePtr,
		doc XmlDocPtr,
		recursive int) XmlNodePtr

	XmlDocCopyNodeList func(
		doc XmlDocPtr,
		node XmlNodePtr) XmlNodePtr

	XmlCopyNodeList func(
		node XmlNodePtr) XmlNodePtr

	XmlNewTextChild func(
		parent XmlNodePtr,
		ns XmlNsPtr,
		name *XmlChar,
		content *XmlChar) XmlNodePtr

	XmlNewDocRawNode func(
		doc XmlDocPtr,
		ns XmlNsPtr,
		name *XmlChar,
		content *XmlChar) XmlNodePtr

	XmlNewDocFragment func(
		doc XmlDocPtr) XmlNodePtr

	XmlGetLineNo func(
		node XmlNodePtr) Long

	XmlGetNodePath func(
		node XmlNodePtr) *XmlChar

	XmlDocGetRootElement func(
		doc XmlDocPtr) XmlNodePtr

	XmlGetLastChild func(
		parent XmlNodePtr) XmlNodePtr

	XmlNodeIsText func(
		node XmlNodePtr) int

	XmlIsBlankNode func(
		node XmlNodePtr) int

	XmlDocSetRootElement func(
		doc XmlDocPtr,
		root XmlNodePtr) XmlNodePtr

	XmlNodeSetName func(
		cur XmlNodePtr,
		name *XmlChar)

	XmlAddChild func(
		parent XmlNodePtr,
		cur XmlNodePtr) XmlNodePtr

	XmlAddChildList func(
		parent XmlNodePtr,
		cur XmlNodePtr) XmlNodePtr

	XmlReplaceNode func(
		old XmlNodePtr,
		cur XmlNodePtr) XmlNodePtr

	XmlAddPrevSibling func(
		cur XmlNodePtr,
		elem XmlNodePtr) XmlNodePtr

	XmlAddSibling func(
		cur XmlNodePtr,
		elem XmlNodePtr) XmlNodePtr

	XmlAddNextSibling func(
		cur XmlNodePtr,
		elem XmlNodePtr) XmlNodePtr

	XmlUnlinkNode func(
		cur XmlNodePtr)

	XmlTextMerge func(
		first XmlNodePtr,
		second XmlNodePtr) XmlNodePtr

	XmlTextConcat func(
		node XmlNodePtr,
		content *XmlChar,
		leng int) int

	XmlFreeNodeList func(
		cur XmlNodePtr)

	XmlFreeNode func(
		cur XmlNodePtr)

	XmlSetTreeDoc func(
		tree XmlNodePtr,
		doc XmlDocPtr)

	XmlSetListDoc func(
		list XmlNodePtr,
		doc XmlDocPtr)

	XmlSearchNs func(
		doc XmlDocPtr,
		node XmlNodePtr,
		nameSpace *XmlChar) XmlNsPtr

	XmlSearchNsByHref func(
		doc XmlDocPtr,
		node XmlNodePtr,
		href *XmlChar) XmlNsPtr

	XmlGetNsList func(
		doc XmlDocPtr,
		node XmlNodePtr) *XmlNsPtr

	XmlSetNs func(
		node XmlNodePtr,
		ns XmlNsPtr)

	XmlCopyNamespace func(
		cur XmlNsPtr) XmlNsPtr

	XmlCopyNamespaceList func(
		cur XmlNsPtr) XmlNsPtr

	XmlSetProp func(
		node XmlNodePtr,
		name *XmlChar,
		value *XmlChar) XmlAttrPtr

	XmlSetNsProp func(
		node XmlNodePtr,
		ns XmlNsPtr,
		name *XmlChar,
		value *XmlChar) XmlAttrPtr

	XmlGetNoNsProp func(
		node XmlNodePtr,
		name *XmlChar) *XmlChar

	XmlGetProp func(
		node XmlNodePtr,
		name *XmlChar) *XmlChar

	XmlHasProp func(
		node XmlNodePtr,
		name *XmlChar) XmlAttrPtr

	XmlHasNsProp func(
		node XmlNodePtr,
		name *XmlChar,
		nameSpace *XmlChar) XmlAttrPtr

	XmlGetNsProp func(
		node XmlNodePtr,
		name *XmlChar,
		nameSpace *XmlChar) *XmlChar

	XmlStringGetNodeList func(
		doc XmlDocPtr,
		value *XmlChar) XmlNodePtr

	XmlStringLenGetNodeList func(
		doc XmlDocPtr,
		value *XmlChar,
		leng int) XmlNodePtr

	XmlNodeListGetString func(
		doc XmlDocPtr,
		list XmlNodePtr,
		inLine int) *XmlChar

	XmlNodeListGetRawString func(
		doc XmlDocPtr,
		list XmlNodePtr,
		inLine int) *XmlChar

	XmlNodeSetContent func(
		cur XmlNodePtr,
		content *XmlChar)

	XmlNodeSetContentLen func(
		cur XmlNodePtr,
		content *XmlChar,
		leng int)

	XmlNodeAddContent func(
		cur XmlNodePtr,
		content *XmlChar)

	XmlNodeAddContentLen func(
		cur XmlNodePtr,
		content *XmlChar,
		leng int)

	XmlNodeGetContent func(
		cur XmlNodePtr) *XmlChar

	XmlNodeBufGetContent func(
		buffer XmlBufferPtr,
		cur XmlNodePtr) int

	XmlBufGetNodeContent func(
		buf XmlBufPtr,
		cur XmlNodePtr) int

	XmlNodeGetLang func(
		cur XmlNodePtr) *XmlChar

	XmlNodeGetSpacePreserve func(
		cur XmlNodePtr) int

	XmlNodeSetLang func(
		cur XmlNodePtr,
		lang *XmlChar)

	XmlNodeSetSpacePreserve func(
		cur XmlNodePtr,
		val int)

	XmlNodeGetBase func(
		doc XmlDocPtr,
		cur XmlNodePtr) *XmlChar

	XmlNodeSetBase func(
		cur XmlNodePtr,
		uri *XmlChar)

	XmlRemoveProp func(
		cur XmlAttrPtr) int

	XmlUnsetNsProp func(
		node XmlNodePtr,
		ns XmlNsPtr,
		name *XmlChar) int

	XmlUnsetProp func(
		node XmlNodePtr,
		name *XmlChar) int

	XmlBufferWriteCHAR func(
		buf XmlBufferPtr,
		string *XmlChar)

	XmlBufferWriteChar func(
		buf XmlBufferPtr,
		string *Char)

	XmlBufferWriteQuotedString func(
		buf XmlBufferPtr,
		string *XmlChar)

	XmlAttrSerializeTxtContent func(
		buf XmlBufferPtr,
		doc XmlDocPtr,
		attr XmlAttrPtr,
		string *XmlChar)

	XmlReconciliateNs func(
		doc XmlDocPtr,
		tree XmlNodePtr) int

	XmlDocDumpFormatMemory func(
		cur XmlDocPtr,
		mem **XmlChar,
		size *int,
		format int)

	XmlDocDumpMemory func(
		cur XmlDocPtr,
		mem **XmlChar,
		size *int)

	XmlDocDumpMemoryEnc func(
		out_doc XmlDocPtr,
		doc_txt_ptr **XmlChar,
		doc_txt_len *int,
		txt_encoding *Char)

	XmlDocDumpFormatMemoryEnc func(
		out_doc XmlDocPtr,
		doc_txt_ptr **XmlChar,
		doc_txt_len *int,
		txt_encoding *Char,
		format int)

	XmlDocFormatDump func(
		f *FILE,
		cur XmlDocPtr,
		format int) int

	XmlDocDump func(f *FILE, cur XmlDocPtr) int

	XmlElemDump func(f *FILE, doc XmlDocPtr, cur XmlNodePtr)

	XmlSaveFile func(filename *Char, cur XmlDocPtr) int

	XmlSaveFormatFile func(
		filename *Char, cur XmlDocPtr, format int) int

	XmlBufNodeDump func(buf XmlBufPtr, doc XmlDocPtr,
		cur XmlNodePtr, level, format int) Size_t

	XmlNodeDump func(buf XmlBufferPtr, doc XmlDocPtr,
		cur XmlNodePtr, level, format int) int

	XmlSaveFileTo func(
		buf XmlOutputBufferPtr,
		cur XmlDocPtr,
		encoding *Char) int

	XmlSaveFormatFileTo func(
		buf XmlOutputBufferPtr,
		cur XmlDocPtr,
		encoding *Char,
		format int) int

	XmlNodeDumpOutput func(
		buf XmlOutputBufferPtr,
		doc XmlDocPtr,
		cur XmlNodePtr,
		level int,
		format int,
		encoding *Char)

	XmlSaveFormatFileEnc func(
		filename *Char,
		cur XmlDocPtr,
		encoding *Char,
		format int) int

	XmlSaveFileEnc func(
		filename *Char,
		cur XmlDocPtr,
		encoding *Char) int

	XmlIsXHTML func(
		systemID *XmlChar,
		publicID *XmlChar) int

	XmlGetDocCompressMode func(
		doc XmlDocPtr) int

	XmlSetDocCompressMode func(
		doc XmlDocPtr,
		mode int)

	XmlGetCompressMode func() int

	XmlSetCompressMode func(
		mode int)

	XmlDOMWrapNewCtxt func() XmlDOMWrapCtxtPtr

	XmlDOMWrapFreeCtxt func(
		ctxt XmlDOMWrapCtxtPtr)

	XmlDOMWrapReconcileNamespaces func(
		ctxt XmlDOMWrapCtxtPtr,
		elem XmlNodePtr,
		options int) int

	XmlDOMWrapAdoptNode func(
		ctxt XmlDOMWrapCtxtPtr,
		sourceDoc XmlDocPtr,
		node XmlNodePtr,
		destDoc XmlDocPtr,
		destParent XmlNodePtr,
		options int) int

	XmlDOMWrapRemoveNode func(
		ctxt XmlDOMWrapCtxtPtr,
		doc XmlDocPtr,
		node XmlNodePtr,
		options int) int

	XmlDOMWrapCloneNode func(
		ctxt XmlDOMWrapCtxtPtr,
		sourceDoc XmlDocPtr,
		node XmlNodePtr,
		clonedNode *XmlNodePtr,
		destDoc XmlDocPtr,
		destParent XmlNodePtr,
		deep int,
		options int) int

	XmlChildElementCount func(
		parent XmlNodePtr) Unsigned_long

	XmlNextElementSibling func(
		node XmlNodePtr) XmlNodePtr

	XmlFirstElementChild func(
		parent XmlNodePtr) XmlNodePtr

	XmlLastElementChild func(
		parent XmlNodePtr) XmlNodePtr

	XmlPreviousElementSibling func(
		node XmlNodePtr) XmlNodePtr

	XmlMemSetup func(
		freeFunc XmlFreeFunc,
		mallocFunc XmlMallocFunc,
		reallocFunc XmlReallocFunc,
		strdupFunc XmlStrdupFunc) int

	XmlMemGet func(
		freeFunc *XmlFreeFunc,
		mallocFunc *XmlMallocFunc,
		reallocFunc *XmlReallocFunc,
		strdupFunc *XmlStrdupFunc) int

	XmlGcMemSetup func(
		freeFunc XmlFreeFunc,
		mallocFunc XmlMallocFunc,
		mallocAtomicFunc XmlMallocFunc,
		reallocFunc XmlReallocFunc,
		strdupFunc XmlStrdupFunc) int

	XmlGcMemGet func(
		freeFunc *XmlFreeFunc,
		mallocFunc *XmlMallocFunc,
		mallocAtomicFunc *XmlMallocFunc,
		reallocFunc *XmlReallocFunc,
		strdupFunc *XmlStrdupFunc) int

	XmlInitMemory func() int

	XmlCleanupMemory func()

	XmlMemUsed func() int

	XmlMemBlocks func() int

	XmlMemDisplay func(
		fp *FILE)

	XmlMemDisplayLast func(fp *FILE, nbBytes Long)

	XmlMemShow func(fp *FILE, nr int)

	XmlMemoryDump func()

	XmlMemMalloc func(size Size_t) *Void

	XmlMemRealloc func(ptr *Void, size Size_t) *Void

	XmlMemFree func(ptr *Void)

	XmlMemoryStrdup func(str *Char) *Char

	XmlMallocLoc func(size Size_t, file *Char, line int) *Void

	XmlReallocLoc func(
		ptr *Void, size Size_t, file *Char, line int) *Void

	XmlMallocAtomicLoc func(
		size Size_t,
		file *Char,
		line int) *Void

	XmlMemStrdupLoc func(
		str *Char,
		file *Char,
		line int) *Char

	XmlHashCreate func(
		size int) XmlHashTablePtr

	XmlHashCreateDict func(
		size int,
		dict XmlDictPtr) XmlHashTablePtr

	XmlHashFree func(
		table XmlHashTablePtr,
		f XmlHashDeallocator)

	XmlHashAddEntry func(
		table XmlHashTablePtr,
		name *XmlChar,
		userdata *Void) int

	XmlHashUpdateEntry func(
		table XmlHashTablePtr,
		name *XmlChar,
		userdata *Void,
		f XmlHashDeallocator) int

	XmlHashAddEntry2 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		userdata *Void) int

	XmlHashUpdateEntry2 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		userdata *Void,
		f XmlHashDeallocator) int

	XmlHashAddEntry3 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		name3 *XmlChar,
		userdata *Void) int

	XmlHashUpdateEntry3 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		name3 *XmlChar,
		userdata *Void,
		f XmlHashDeallocator) int

	XmlHashRemoveEntry func(
		table XmlHashTablePtr,
		name *XmlChar,
		f XmlHashDeallocator) int

	XmlHashRemoveEntry2 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		f XmlHashDeallocator) int

	XmlHashRemoveEntry3 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		name3 *XmlChar,
		f XmlHashDeallocator) int

	XmlHashLookup func(
		table XmlHashTablePtr,
		name *XmlChar) *Void

	XmlHashLookup2 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar) *Void

	XmlHashLookup3 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		name3 *XmlChar) *Void

	XmlHashQLookup func(
		table XmlHashTablePtr,
		name *XmlChar,
		prefix *XmlChar) *Void

	XmlHashQLookup2 func(
		table XmlHashTablePtr,
		name *XmlChar,
		prefix *XmlChar,
		name2 *XmlChar,
		prefix2 *XmlChar) *Void

	XmlHashQLookup3 func(
		table XmlHashTablePtr,
		name *XmlChar,
		prefix *XmlChar,
		name2 *XmlChar,
		prefix2 *XmlChar,
		name3 *XmlChar,
		prefix3 *XmlChar) *Void

	XmlHashCopy func(
		table XmlHashTablePtr,
		f XmlHashCopier) XmlHashTablePtr

	XmlHashSize func(
		table XmlHashTablePtr) int

	XmlHashScan func(
		table XmlHashTablePtr,
		f XmlHashScanner,
		data *Void)

	XmlHashScan3 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		name3 *XmlChar,
		f XmlHashScanner,
		data *Void)

	XmlHashScanFull func(
		table XmlHashTablePtr,
		f XmlHashScannerFull,
		data *Void)

	XmlHashScanFull3 func(
		table XmlHashTablePtr,
		name *XmlChar,
		name2 *XmlChar,
		name3 *XmlChar,
		f XmlHashScannerFull,
		data *Void)

	XmlSetGenericErrorFunc func(
		ctx *Void,
		handler XmlGenericErrorFunc)

	InitGenericErrorDefaultFunc func(
		handler *XmlGenericErrorFunc)

	XmlSetStructuredErrorFunc func(
		ctx *Void,
		handler XmlStructuredErrorFunc)

	XmlParserError func(
		ctx *Void,
		msg *Char,
		v ...VArg)

	XmlParserWarning func(
		ctx *Void,
		msg *Char,
		v ...VArg)

	XmlParserValidityError func(
		ctx *Void, msg string, v ...VArg)

	XmlParserValidityWarning func(
		ctx *Void, msg string, v ...VArg)

	XmlParserPrintFileInfo func(input XmlParserInputPtr)

	XmlParserPrintFileContext func(input XmlParserInputPtr)

	XmlGetLastError func() XmlErrorPtr

	XmlResetLastError func()

	XmlCtxtGetLastError func(ctx *Void) XmlErrorPtr

	XmlCtxtResetLastError func(ctx *Void)

	XmlResetError func(err XmlErrorPtr)

	XmlCopyError func(from XmlErrorPtr, to XmlErrorPtr) int

	XmlListCreate func(
		deallocator XmlListDeallocator,
		compare XmlListDataCompare) XmlListPtr

	XmlListDelete func(
		l XmlListPtr)

	XmlListSearch func(
		l XmlListPtr,
		data *Void) *Void

	XmlListReverseSearch func(
		l XmlListPtr,
		data *Void) *Void

	XmlListInsert func(
		l XmlListPtr,
		data *Void) int

	XmlListAppend func(
		l XmlListPtr,
		data *Void) int

	XmlListRemoveFirst func(
		l XmlListPtr,
		data *Void) int

	XmlListRemoveLast func(
		l XmlListPtr,
		data *Void) int

	XmlListRemoveAll func(
		l XmlListPtr,
		data *Void) int

	XmlListClear func(
		l XmlListPtr)

	XmlListEmpty func(
		l XmlListPtr) int

	XmlListFront func(
		l XmlListPtr) XmlLinkPtr

	XmlListEnd func(
		l XmlListPtr) XmlLinkPtr

	XmlListSize func(
		l XmlListPtr) int

	XmlListPopFront func(
		l XmlListPtr)

	XmlListPopBack func(
		l XmlListPtr)

	XmlListPushFront func(
		l XmlListPtr,
		data *Void) int

	XmlListPushBack func(
		l XmlListPtr,
		data *Void) int

	XmlListReverse func(
		l XmlListPtr)

	XmlListSort func(
		l XmlListPtr)

	XmlListWalk func(
		l XmlListPtr,
		walker XmlListWalker,
		user *Void)

	XmlListReverseWalk func(
		l XmlListPtr,
		walker XmlListWalker,
		user *Void)

	XmlListMerge func(
		l1 XmlListPtr,
		l2 XmlListPtr)

	XmlListDup func(
		old XmlListPtr) XmlListPtr

	XmlListCopy func(
		cur XmlListPtr,
		old XmlListPtr) int

	XmlLinkGetData func(
		lk XmlLinkPtr) *Void

	XmlNewAutomata func() XmlAutomataPtr

	XmlFreeAutomata func(
		am XmlAutomataPtr)

	XmlAutomataGetInitState func(
		am XmlAutomataPtr) XmlAutomataStatePtr

	XmlAutomataSetFinalState func(
		am XmlAutomataPtr,
		state XmlAutomataStatePtr) int

	XmlAutomataNewState func(
		am XmlAutomataPtr) XmlAutomataStatePtr

	XmlAutomataNewTransition func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token *XmlChar,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewTransition2 func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token *XmlChar,
		token2 *XmlChar,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewNegTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token *XmlChar,
		token2 *XmlChar,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewCountTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token *XmlChar,
		min int,
		max int,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewCountTrans2 func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token *XmlChar,
		token2 *XmlChar,
		min int,
		max int,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewOnceTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token *XmlChar,
		min int,
		max int,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewOnceTrans2 func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		token *XmlChar,
		token2 *XmlChar,
		min int,
		max int,
		data *Void) XmlAutomataStatePtr

	XmlAutomataNewAllTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		lax int) XmlAutomataStatePtr

	XmlAutomataNewEpsilon func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr) XmlAutomataStatePtr

	XmlAutomataNewCountedTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		counter int) XmlAutomataStatePtr

	XmlAutomataNewCounterTrans func(
		am XmlAutomataPtr,
		from XmlAutomataStatePtr,
		to XmlAutomataStatePtr,
		counter int) XmlAutomataStatePtr

	XmlAutomataNewCounter func(
		am XmlAutomataPtr,
		min int,
		max int) int

	XmlAutomataCompile func(
		am XmlAutomataPtr) XmlRegexpPtr

	XmlAutomataIsDeterminist func(
		am XmlAutomataPtr) int

	XmlAddNotationDecl func(
		ctxt XmlValidCtxtPtr,
		dtd XmlDtdPtr,
		name *XmlChar,
		PublicID *XmlChar,
		SystemID *XmlChar) XmlNotationPtr

	XmlCopyNotationTable func(
		table XmlNotationTablePtr) XmlNotationTablePtr

	XmlFreeNotationTable func(table XmlNotationTablePtr)

	XmlDumpNotationDecl func(
		buf XmlBufferPtr, nota XmlNotationPtr)

	XmlDumpNotationTable func(
		buf XmlBufferPtr, table XmlNotationTablePtr)

	XmlNewElementContent func(name *XmlChar,
		t XmlElementContentType) XmlElementContentPtr

	XmlCopyElementContent func(
		content XmlElementContentPtr) XmlElementContentPtr

	XmlFreeElementContent func(cur XmlElementContentPtr)

	XmlNewDocElementContent func(doc XmlDocPtr,
		name *XmlChar, t XmlElementContentType) XmlElementContentPtr

	XmlCopyDocElementContent func(doc XmlDocPtr,
		content XmlElementContentPtr) XmlElementContentPtr

	XmlFreeDocElementContent func(
		doc XmlDocPtr, cur XmlElementContentPtr)

	XmlSnprintfElementContent func(buf *Char, size int,
		content XmlElementContentPtr, englob int)

	XmlSprintfElementContent func(
		buf *Char, content XmlElementContentPtr, englob int)

	XmlAddElementDecl func(
		ctxt XmlValidCtxtPtr,
		dtd XmlDtdPtr,
		name *XmlChar,
		t XmlElementTypeVal,
		content XmlElementContentPtr) XmlElementPtr

	XmlCopyElementTable func(
		table XmlElementTablePtr) XmlElementTablePtr

	XmlFreeElementTable func(table XmlElementTablePtr)

	XmlDumpElementTable func(
		buf XmlBufferPtr, table XmlElementTablePtr)

	XmlDumpElementDecl func(
		buf XmlBufferPtr, elem XmlElementPtr)

	XmlCreateEnumeration func(name *XmlChar) XmlEnumerationPtr

	XmlFreeEnumeration func(cur XmlEnumerationPtr)

	XmlCopyEnumeration func(
		cur XmlEnumerationPtr) XmlEnumerationPtr

	XmlAddAttributeDecl func(
		ctxt XmlValidCtxtPtr,
		dtd XmlDtdPtr,
		elem, name, ns *XmlChar,
		t XmlAttributeType,
		def XmlAttributeDefault,
		defaultValue *XmlChar,
		tree XmlEnumerationPtr) XmlAttributePtr

	XmlCopyAttributeTable func(
		table XmlAttributeTablePtr) XmlAttributeTablePtr

	XmlFreeAttributeTable func(table XmlAttributeTablePtr)

	XmlDumpAttributeTable func(
		buf XmlBufferPtr, table XmlAttributeTablePtr)

	XmlDumpAttributeDecl func(
		buf XmlBufferPtr, attr XmlAttributePtr)

	XmlAddID func(ctxt XmlValidCtxtPtr, doc XmlDocPtr,
		value *XmlChar, attr XmlAttrPtr) XmlIDPtr

	XmlFreeIDTable func(table XmlIDTablePtr)

	XmlGetID func(doc XmlDocPtr, ID *XmlChar) XmlAttrPtr

	XmlIsID func(
		doc XmlDocPtr, elem XmlNodePtr, attr XmlAttrPtr) int

	XmlRemoveID func(doc XmlDocPtr, attr XmlAttrPtr) int

	XmlAddRef func(ctxt XmlValidCtxtPtr, doc XmlDocPtr,
		value *XmlChar, attr XmlAttrPtr) XmlRefPtr

	XmlFreeRefTable func(table XmlRefTablePtr)

	XmlIsRef func(doc XmlDocPtr, elem XmlNodePtr, attr XmlAttrPtr) int

	XmlRemoveRef func(doc XmlDocPtr, attr XmlAttrPtr) int

	XmlGetRefs func(doc XmlDocPtr, ID *XmlChar) XmlListPtr

	XmlNewValidCtxt func() XmlValidCtxtPtr

	XmlFreeValidCtxt func(XmlValidCtxtPtr)

	XmlValidateRoot func(ctxt XmlValidCtxtPtr, doc XmlDocPtr) int

	XmlValidateElementDecl func(ctxt XmlValidCtxtPtr,
		doc XmlDocPtr, elem XmlElementPtr) int

	XmlValidNormalizeAttributeValue func(doc XmlDocPtr,
		elem XmlNodePtr, name *XmlChar, value *XmlChar) *XmlChar

	XmlValidCtxtNormalizeAttributeValue func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		name *XmlChar,
		value *XmlChar) *XmlChar

	XmlValidateAttributeDecl func(ctxt XmlValidCtxtPtr,
		doc XmlDocPtr, attr XmlAttributePtr) int

	XmlValidateAttributeValue func(
		t XmlAttributeType, value *XmlChar) int

	XmlValidateNotationDecl func(ctxt XmlValidCtxtPtr,
		doc XmlDocPtr, nota XmlNotationPtr) int

	XmlValidateDtd func(ctxt XmlValidCtxtPtr,
		doc XmlDocPtr, dtd XmlDtdPtr) int

	XmlValidateDtdFinal func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr) int

	XmlValidateDocument func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr) int

	XmlValidateElement func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr) int

	XmlValidateOneElement func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr) int

	XmlValidateOneAttribute func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		attr XmlAttrPtr,
		value *XmlChar) int

	XmlValidateOneNamespace func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		prefix *XmlChar,
		ns XmlNsPtr,
		value *XmlChar) int

	XmlValidateDocumentFinal func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr) int

	XmlValidateNotationUse func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		notationName *XmlChar) int

	XmlIsMixedElement func(
		doc XmlDocPtr,
		name *XmlChar) int

	XmlGetDtdAttrDesc func(
		dtd XmlDtdPtr,
		elem *XmlChar,
		name *XmlChar) XmlAttributePtr

	XmlGetDtdQAttrDesc func(
		dtd XmlDtdPtr,
		elem *XmlChar,
		name *XmlChar,
		prefix *XmlChar) XmlAttributePtr

	XmlGetDtdNotationDesc func(
		dtd XmlDtdPtr,
		name *XmlChar) XmlNotationPtr

	XmlGetDtdQElementDesc func(
		dtd XmlDtdPtr,
		name *XmlChar,
		prefix *XmlChar) XmlElementPtr

	XmlGetDtdElementDesc func(
		dtd XmlDtdPtr,
		name *XmlChar) XmlElementPtr

	XmlValidGetPotentialChildren func(
		ctree *XmlElementContent,
		names **XmlChar,
		leng *int,
		max int) int

	XmlValidGetValidElements func(
		prev, next *XmlNode,
		names **XmlChar,
		max int) int

	XmlValidateNameValue func(
		value *XmlChar) int

	XmlValidateNamesValue func(
		value *XmlChar) int

	XmlValidateNmtokenValue func(
		value *XmlChar) int

	XmlValidateNmtokensValue func(
		value *XmlChar) int

	XmlValidBuildContentModel func(
		ctxt XmlValidCtxtPtr,
		elem XmlElementPtr) int

	XmlValidatePushElement func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		qname *XmlChar) int

	XmlValidatePushCData func(
		ctxt XmlValidCtxtPtr,
		data *XmlChar,
		leng int) int

	XmlValidatePopElement func(
		ctxt XmlValidCtxtPtr,
		doc XmlDocPtr,
		elem XmlNodePtr,
		qname *XmlChar) int

	XmlInitializePredefinedEntities func()

	XmlNewEntity func(
		doc XmlDocPtr,
		name *XmlChar,
		typ int,
		ExternalID *XmlChar,
		SystemID *XmlChar,
		content *XmlChar) XmlEntityPtr

	XmlAddDocEntity func(
		doc XmlDocPtr,
		name *XmlChar,
		typ int,
		ExternalID *XmlChar,
		SystemID *XmlChar,
		content *XmlChar) XmlEntityPtr

	XmlAddDtdEntity func(
		doc XmlDocPtr,
		name *XmlChar,
		typ int,
		ExternalID *XmlChar,
		SystemID *XmlChar,
		content *XmlChar) XmlEntityPtr

	XmlGetPredefinedEntity func(
		name *XmlChar) XmlEntityPtr

	XmlGetDocEntity func(
		doc XmlDocPtr,
		name *XmlChar) XmlEntityPtr

	XmlGetDtdEntity func(
		doc XmlDocPtr,
		name *XmlChar) XmlEntityPtr

	XmlGetParameterEntity func(
		doc XmlDocPtr,
		name *XmlChar) XmlEntityPtr

	XmlEncodeEntities func(
		doc XmlDocPtr,
		input *XmlChar) *XmlChar

	XmlEncodeEntitiesReentrant func(
		doc XmlDocPtr,
		input *XmlChar) *XmlChar

	XmlEncodeSpecialChars func(
		doc XmlDocPtr,
		input *XmlChar) *XmlChar

	XmlCreateEntitiesTable func() XmlEntitiesTablePtr

	XmlCopyEntitiesTable func(
		table XmlEntitiesTablePtr) XmlEntitiesTablePtr

	XmlFreeEntitiesTable func(
		table XmlEntitiesTablePtr)

	XmlDumpEntitiesTable func(
		buf XmlBufferPtr,
		table XmlEntitiesTablePtr)

	XmlDumpEntityDecl func(
		buf XmlBufferPtr,
		ent XmlEntityPtr)

	XmlCleanupPredefinedEntities func()

	XmlInitCharEncodingHandlers func()

	XmlCleanupCharEncodingHandlers func()

	XmlRegisterCharEncodingHandler func(
		handler XmlCharEncodingHandlerPtr)

	XmlGetCharEncodingHandler func(
		enc XmlCharEncoding) XmlCharEncodingHandlerPtr

	XmlFindCharEncodingHandler func(
		name *Char) XmlCharEncodingHandlerPtr

	XmlNewCharEncodingHandler func(
		name *Char,
		input XmlCharEncodingInputFunc,
		output XmlCharEncodingOutputFunc) XmlCharEncodingHandlerPtr

	XmlAddEncodingAlias func(
		name *Char,
		alias *Char) int

	XmlDelEncodingAlias func(
		alias *Char) int

	XmlGetEncodingAlias func(
		alias *Char) *Char

	XmlCleanupEncodingAliases func()

	XmlParseCharEncoding func(
		name *Char) XmlCharEncoding

	XmlGetCharEncodingName func(
		enc XmlCharEncoding) *Char

	XmlDetectCharEncoding func(
		in *Unsigned_char,
		leng int) XmlCharEncoding

	XmlCharEncOutFunc func(
		handler *XmlCharEncodingHandler,
		out XmlBufferPtr,
		in XmlBufferPtr) int

	XmlCharEncInFunc func(
		handler *XmlCharEncodingHandler,
		out XmlBufferPtr,
		in XmlBufferPtr) int

	XmlCharEncFirstLine func(
		handler *XmlCharEncodingHandler,
		out XmlBufferPtr,
		in XmlBufferPtr) int

	XmlCharEncCloseFunc func(
		handler *XmlCharEncodingHandler) int

	UTF8Toisolat1 func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int) int

	Isolat1ToUTF8 func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int) int

	XmlCleanupInputCallbacks func()

	XmlPopInputCallbacks func() int

	XmlRegisterDefaultInputCallbacks func()

	XmlAllocParserInputBuffer func(
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateFile func(
		file *FILE,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateFd func(
		fd int,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateMem func(
		mem *Char,
		size int,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateStatic func(
		mem *Char,
		size int,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferCreateIO func(
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlParserInputBufferRead func(
		in XmlParserInputBufferPtr,
		leng int) int

	XmlParserInputBufferGrow func(
		in XmlParserInputBufferPtr,
		leng int) int

	XmlParserInputBufferPush func(
		in XmlParserInputBufferPtr,
		leng int,
		buf *Char) int

	XmlFreeParserInputBuffer func(
		in XmlParserInputBufferPtr)

	XmlParserGetDirectory func(
		filename *Char) *Char

	XmlRegisterInputCallbacks func(
		matchFunc XmlInputMatchCallback,
		openFunc XmlInputOpenCallback,
		readFunc XmlInputReadCallback,
		closeFunc XmlInputCloseCallback) int

	XmlParserInputBufferCreateFilename func(
		URI *Char,
		enc XmlCharEncoding) XmlParserInputBufferPtr

	XmlCleanupOutputCallbacks func()

	XmlRegisterDefaultOutputCallbacks func()

	XmlAllocOutputBuffer func(
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferCreateFilename func(
		URI *Char,
		encoder XmlCharEncodingHandlerPtr,
		compression int) XmlOutputBufferPtr

	XmlOutputBufferCreateFile func(
		file *FILE,
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferCreateBuffer func(
		buffer XmlBufferPtr,
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferCreateFd func(
		fd int,
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferCreateIO func(
		iowrite XmlOutputWriteCallback,
		ioclose XmlOutputCloseCallback,
		ioctx *Void,
		encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr

	XmlOutputBufferGetContent func(
		out XmlOutputBufferPtr) *XmlChar

	XmlOutputBufferGetSize func(
		out XmlOutputBufferPtr) Size_t

	XmlOutputBufferWrite func(
		out XmlOutputBufferPtr,
		leng int,
		buf *Char) int

	XmlOutputBufferWriteString func(
		out XmlOutputBufferPtr,
		str *Char) int

	XmlOutputBufferWriteEscape func(
		out XmlOutputBufferPtr,
		str *XmlChar,
		escaping XmlCharEncodingOutputFunc) int

	XmlOutputBufferFlush func(
		out XmlOutputBufferPtr) int

	XmlOutputBufferClose func(
		out XmlOutputBufferPtr) int

	XmlRegisterOutputCallbacks func(
		matchFunc XmlOutputMatchCallback,
		openFunc XmlOutputOpenCallback,
		writeFunc XmlOutputWriteCallback,
		closeFunc XmlOutputCloseCallback) int

	XmlRegisterHTTPPostCallbacks func()

	XmlCheckHTTPInput func(
		ctxt XmlParserCtxtPtr,
		ret XmlParserInputPtr) XmlParserInputPtr

	XmlNoNetExternalEntityLoader func(
		URL *Char,
		ID *Char,
		ctxt XmlParserCtxtPtr) XmlParserInputPtr

	XmlNormalizeWindowsPath func(
		path *XmlChar) *XmlChar

	XmlCheckFilename func(
		path *Char) int

	XmlFileMatch func(
		filename *Char) int

	XmlFileOpen func(
		filename *Char) *Void

	XmlFileRead func(
		context *Void,
		buffer *Char,
		leng int) int

	XmlFileClose func(context *Void) int

	XmlIOHTTPMatch func(filename *Char) int

	XmlIOHTTPOpen func(filename *Char) *Void

	XmlIOHTTPOpenW func(post_uri *Char, compression int) *Void

	XmlIOHTTPRead func(context *Void, buffer *Char, leng int) int

	XmlIOHTTPClose func(
		context *Void) int

	XmlIOFTPMatch func(
		filename *Char) int

	XmlIOFTPOpen func(
		filename *Char) *Void

	XmlIOFTPRead func(
		context *Void,
		buffer *Char,
		leng int) int

	XmlIOFTPClose func(
		context *Void) int

	XmlInitParser func()

	XmlCleanupParser func()

	XmlParserInputRead func(
		in XmlParserInputPtr,
		leng int) int

	XmlParserInputGrow func(
		in XmlParserInputPtr,
		leng int) int

	XmlParseDoc func(
		cur *XmlChar) XmlDocPtr

	XmlParseFile func(
		filename *Char) XmlDocPtr

	XmlParseMemory func(
		buffer *Char,
		size int) XmlDocPtr

	XmlSubstituteEntitiesDefault func(
		val int) int

	XmlKeepBlanksDefault func(
		val int) int

	XmlStopParser func(
		ctxt XmlParserCtxtPtr)

	XmlPedanticParserDefault func(
		val int) int

	XmlLineNumbersDefault func(
		val int) int

	XmlRecoverDoc func(
		cur *XmlChar) XmlDocPtr

	XmlRecoverMemory func(
		buffer *Char,
		size int) XmlDocPtr

	XmlRecoverFile func(
		filename *Char) XmlDocPtr

	XmlParseDocument func(
		ctxt XmlParserCtxtPtr) int

	XmlParseExtParsedEnt func(
		ctxt XmlParserCtxtPtr) int

	XmlSAXUserParseFile func(
		sax XmlSAXHandlerPtr,
		user_data *Void,
		filename *Char) int

	XmlSAXUserParseMemory func(
		sax XmlSAXHandlerPtr,
		user_data *Void,
		buffer *Char,
		size int) int

	XmlSAXParseDoc func(
		sax XmlSAXHandlerPtr,
		cur *XmlChar,
		recovery int) XmlDocPtr

	XmlSAXParseMemory func(
		sax XmlSAXHandlerPtr,
		buffer *Char,
		size int,
		recovery int) XmlDocPtr

	XmlSAXParseMemoryWithData func(
		sax XmlSAXHandlerPtr,
		buffer *Char,
		size int,
		recovery int,
		data *Void) XmlDocPtr

	XmlSAXParseFile func(
		sax XmlSAXHandlerPtr,
		filename *Char,
		recovery int) XmlDocPtr

	XmlSAXParseFileWithData func(
		sax XmlSAXHandlerPtr,
		filename *Char,
		recovery int,
		data *Void) XmlDocPtr

	XmlSAXParseEntity func(
		sax XmlSAXHandlerPtr,
		filename *Char) XmlDocPtr

	XmlParseEntity func(
		filename *Char) XmlDocPtr

	XmlSAXParseDTD func(
		sax XmlSAXHandlerPtr,
		ExternalID *XmlChar,
		SystemID *XmlChar) XmlDtdPtr

	XmlParseDTD func(
		ExternalID *XmlChar,
		SystemID *XmlChar) XmlDtdPtr

	XmlIOParseDTD func(
		sax XmlSAXHandlerPtr,
		input XmlParserInputBufferPtr,
		enc XmlCharEncoding) XmlDtdPtr

	XmlParseBalancedChunkMemory func(
		doc XmlDocPtr,
		sax XmlSAXHandlerPtr,
		user_data *Void,
		depth int,
		string *XmlChar,
		lst *XmlNodePtr) int

	XmlParseInNodeContext func(
		node XmlNodePtr,
		data *Char,
		datalen int,
		options int,
		lst *XmlNodePtr) XmlParserErrors

	XmlParseBalancedChunkMemoryRecover func(
		doc XmlDocPtr,
		sax XmlSAXHandlerPtr,
		user_data *Void,
		depth int,
		string *XmlChar,
		lst *XmlNodePtr,
		recover int) int

	XmlParseExternalEntity func(
		doc XmlDocPtr,
		sax XmlSAXHandlerPtr,
		user_data *Void,
		depth int,
		URL *XmlChar,
		ID *XmlChar,
		lst *XmlNodePtr) int

	XmlParseCtxtExternalEntity func(
		ctx XmlParserCtxtPtr,
		URL *XmlChar,
		ID *XmlChar,
		lst *XmlNodePtr) int

	XmlNewParserCtxt func() XmlParserCtxtPtr

	XmlInitParserCtxt func(
		ctxt XmlParserCtxtPtr) int

	XmlClearParserCtxt func(
		ctxt XmlParserCtxtPtr)

	XmlFreeParserCtxt func(
		ctxt XmlParserCtxtPtr)

	XmlSetupParserForBuffer func(
		ctxt XmlParserCtxtPtr,
		buffer *XmlChar,
		filename *Char)

	XmlCreateDocParserCtxt func(
		cur *XmlChar) XmlParserCtxtPtr

	XmlGetFeaturesList func(
		leng *int,
		result **Char) int

	XmlGetFeature func(
		ctxt XmlParserCtxtPtr,
		name *Char,
		result *Void) int

	XmlSetFeature func(
		ctxt XmlParserCtxtPtr,
		name *Char,
		value *Void) int

	XmlCreatePushParserCtxt func(
		sax XmlSAXHandlerPtr,
		user_data *Void,
		chunk *Char,
		size int,
		filename *Char) XmlParserCtxtPtr

	XmlParseChunk func(
		ctxt XmlParserCtxtPtr,
		chunk *Char,
		size int,
		terminate int) int

	XmlCreateIOParserCtxt func(
		sax XmlSAXHandlerPtr,
		user_data *Void,
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		enc XmlCharEncoding) XmlParserCtxtPtr

	XmlNewIOInputStream func(
		ctxt XmlParserCtxtPtr,
		input XmlParserInputBufferPtr,
		enc XmlCharEncoding) XmlParserInputPtr

	XmlParserFindNodeInfo func(
		ctxt XmlParserCtxtPtr,
		node XmlNodePtr) *XmlParserNodeInfo

	XmlInitNodeInfoSeq func(
		seq XmlParserNodeInfoSeqPtr)

	XmlClearNodeInfoSeq func(
		seq XmlParserNodeInfoSeqPtr)

	XmlParserFindNodeInfoIndex func(
		seq XmlParserNodeInfoSeqPtr,
		node XmlNodePtr) Unsigned_long

	XmlParserAddNodeInfo func(
		ctxt XmlParserCtxtPtr,
		info XmlParserNodeInfoPtr)

	XmlSetExternalEntityLoader func(
		f XmlExternalEntityLoader)

	XmlGetExternalEntityLoader func() XmlExternalEntityLoader

	XmlLoadExternalEntity func(
		URL *Char,
		ID *Char,
		ctxt XmlParserCtxtPtr) XmlParserInputPtr

	XmlByteConsumed func(
		ctxt XmlParserCtxtPtr) Long

	XmlCtxtReset func(
		ctxt XmlParserCtxtPtr)

	XmlCtxtResetPush func(
		ctxt XmlParserCtxtPtr,
		chunk *Char,
		size int,
		filename *Char,
		encoding *Char) int

	XmlCtxtUseOptions func(
		ctxt XmlParserCtxtPtr,
		options int) int

	XmlReadDoc func(
		cur *XmlChar,
		URL *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlReadFile func(
		URL *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlReadMemory func(
		buffer *Char,
		size int,
		URL *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlReadFd func(
		fd int,
		URL *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlReadIO func(
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlCtxtReadDoc func(
		ctxt XmlParserCtxtPtr,
		cur *XmlChar,
		URL *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlCtxtReadFile func(
		ctxt XmlParserCtxtPtr,
		filename *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlCtxtReadMemory func(
		ctxt XmlParserCtxtPtr,
		buffer *Char,
		size int,
		URL *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlCtxtReadFd func(
		ctxt XmlParserCtxtPtr,
		fd int,
		URL *Char,
		encoding *Char,
		options int) XmlDocPtr

	XmlCtxtReadIO func(
		ctxt XmlParserCtxtPtr,
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL, encoding *Char,
		options int) XmlDocPtr

	XmlHasFeature func(feature XmlFeature) int

	xlinkGetDefaultDetect func() XlinkNodeDetectFunc

	xlinkSetDefaultDetect func(f XlinkNodeDetectFunc)

	xlinkGetDefaultHandler func() XlinkHandlerPtr

	xlinkSetDefaultHandler func(handler XlinkHandlerPtr)

	xlinkIsLink func(doc, node XmlNodePtr) XlinkType

	getPublicId func(ctx *Void) *XmlChar

	getSystemId func(ctx *Void) *XmlChar

	setDocumentLocator func(ctx *Void, loc XmlSAXLocatorPtr)

	getLineNumber func(ctx *Void) int

	getColumnNumber func(ctx *Void) int

	IsStandalone func(ctx *Void) int

	hasInternalSubset func(ctx *Void) int

	hasExternalSubset func(ctx *Void) int

	InternalSubset func(
		ctx *Void, name, ExternalID, SystemID *XmlChar)

	externalSubset func(
		ctx *Void, name, ExternalID, SystemID *XmlChar)

	getEntity func(ctx *Void, name *XmlChar) XmlEntityPtr

	getParameterEntity func(
		ctx *Void, name *XmlChar) XmlEntityPtr

	resolveEntity func(ctx *Void,
		publicId, systemId *XmlChar) XmlParserInputPtr

	entityDecl func(ctx *Void, name *XmlChar, typ int,
		publicId, systemId, content *XmlChar)

	attributeDecl func(ctx *Void, elem, fullname *XmlChar,
		typ, def int, defaultValue *XmlChar,
		tree XmlEnumerationPtr)

	elementDecl func(ctx *Void, name *XmlChar,
		typ int, content XmlElementContentPtr)

	notationDecl func(
		ctx *Void, name, publicId, systemId *XmlChar)

	UnparsedEntityDecl func(ctx *Void,
		name, publicId, systemId, notationName *XmlChar)

	startDocument func(ctx *Void)

	endDocument func(ctx *Void)

	attribute func(ctx *Void, fullname, value *XmlChar)

	startElement func(ctx *Void, fullname, atts **XmlChar)

	endElement func(ctx *Void, name *XmlChar)

	reference func(ctx *Void, name *XmlChar)

	characters func(ctx *Void, ch *XmlChar, leng int)

	IgnorableWhitespace func(ctx *Void, ch *XmlChar, leng int)

	processingInstruction func(ctx *Void, target, data *XmlChar)

	globalNamespace func(ctx *Void, href, prefix *XmlChar)

	setNamespace func(ctx *Void, name *XmlChar)

	getNamespace func(ctx *Void) XmlNsPtr

	checkNamespace func(ctx *Void, nameSpace *XmlChar) int

	namespaceDecl func(ctx *Void, href, prefix *XmlChar)

	comment func(ctx *Void, value *XmlChar)

	cdataBlock func(ctx *Void, value *XmlChar, leng int)

	InitXmlDefaultSAXHandler func(
		hdlr *XmlSAXHandlerV1, warning int)

	InithtmlDefaultSAXHandler func(hdlr *XmlSAXHandlerV1)

	InitdocbDefaultSAXHandler func(hdlr *XmlSAXHandlerV1)

	XmlSAX2GetPublicId func(ctx *Void) *XmlChar

	XmlSAX2GetSystemId func(ctx *Void) *XmlChar

	XmlSAX2SetDocumentLocator func(
		ctx *Void, loc XmlSAXLocatorPtr)

	XmlSAX2GetLineNumber func(ctx *Void) int

	XmlSAX2GetColumnNumber func(ctx *Void) int

	XmlSAX2IsStandalone func(ctx *Void) int

	XmlSAX2HasInternalSubset func(ctx *Void) int

	XmlSAX2HasExternalSubset func(ctx *Void) int

	XmlSAX2InternalSubset func(
		ctx *Void,
		name *XmlChar,
		ExternalID *XmlChar,
		SystemID *XmlChar)

	XmlSAX2ExternalSubset func(
		ctx *Void,
		name *XmlChar,
		ExternalID *XmlChar,
		SystemID *XmlChar)

	XmlSAX2GetEntity func(
		ctx *Void,
		name *XmlChar) XmlEntityPtr

	XmlSAX2GetParameterEntity func(
		ctx *Void,
		name *XmlChar) XmlEntityPtr

	XmlSAX2ResolveEntity func(
		ctx *Void,
		publicId *XmlChar,
		systemId *XmlChar) XmlParserInputPtr

	XmlSAX2EntityDecl func(
		ctx *Void,
		name *XmlChar,
		typ int,
		publicId *XmlChar,
		systemId *XmlChar,
		content *XmlChar)

	XmlSAX2AttributeDecl func(
		ctx *Void,
		elem *XmlChar,
		fullname *XmlChar,
		typ int,
		def int,
		defaultValue *XmlChar,
		tree XmlEnumerationPtr)

	XmlSAX2ElementDecl func(
		ctx *Void,
		name *XmlChar,
		typ int,
		content XmlElementContentPtr)

	XmlSAX2NotationDecl func(
		ctx *Void,
		name *XmlChar,
		publicId *XmlChar,
		systemId *XmlChar)

	XmlSAX2UnparsedEntityDecl func(
		ctx *Void,
		name *XmlChar,
		publicId *XmlChar,
		systemId *XmlChar,
		notationName *XmlChar)

	XmlSAX2StartDocument func(
		ctx *Void)

	XmlSAX2EndDocument func(
		ctx *Void)

	XmlSAX2StartElement func(
		ctx *Void,
		fullname *XmlChar,
		atts **XmlChar)

	XmlSAX2EndElement func(
		ctx *Void,
		name *XmlChar)

	XmlSAX2StartElementNs func(
		ctx *Void,
		localname *XmlChar,
		prefix *XmlChar,
		URI *XmlChar,
		nb_namespaces int,
		namespaces **XmlChar,
		nb_attributes int,
		nb_defaulted int,
		attributes **XmlChar)

	XmlSAX2EndElementNs func(
		ctx *Void,
		localname *XmlChar,
		prefix *XmlChar,
		URI *XmlChar)

	XmlSAX2Reference func(
		ctx *Void,
		name *XmlChar)

	XmlSAX2Characters func(
		ctx *Void,
		ch *XmlChar,
		leng int)

	XmlSAX2IgnorableWhitespace func(
		ctx *Void,
		ch *XmlChar,
		leng int)

	XmlSAX2ProcessingInstruction func(
		ctx *Void,
		target *XmlChar,
		data *XmlChar)

	XmlSAX2Comment func(
		ctx *Void,
		value *XmlChar)

	XmlSAX2CDataBlock func(
		ctx *Void,
		value *XmlChar,
		leng int)

	XmlSAXDefaultVersion func(
		version int) int

	XmlSAXVersion func(
		hdlr *XmlSAXHandler,
		version int) int

	XmlSAX2InitDefaultSAXHandler func(
		hdlr *XmlSAXHandler,
		warning int)

	XmlSAX2InitHtmlDefaultSAXHandler func(
		hdlr *XmlSAXHandler)

	HtmlDefaultSAXHandlerInit func()

	XmlSAX2InitDocbDefaultSAXHandler func(
		hdlr *XmlSAXHandler)
	docbDefaultSAXHandlerInit func()

	XmlDefaultSAXHandlerInit func()

	XmlInitGlobals func()

	XmlCleanupGlobals func()

	XmlInitializeGlobalState func(
		gs XmlGlobalStatePtr)

	XmlThrDefSetGenericErrorFunc func(
		ctx *Void,
		handler XmlGenericErrorFunc)

	XmlThrDefSetStructuredErrorFunc func(
		ctx *Void,
		handler XmlStructuredErrorFunc)

	XmlRegisterNodeDefault func(
		f XmlRegisterNodeFunc) XmlRegisterNodeFunc

	XmlThrDefRegisterNodeDefault func(
		f XmlRegisterNodeFunc) XmlRegisterNodeFunc

	XmlDeregisterNodeDefault func(
		f XmlDeregisterNodeFunc) XmlDeregisterNodeFunc

	XmlThrDefDeregisterNodeDefault func(
		f XmlDeregisterNodeFunc) XmlDeregisterNodeFunc

	XmlThrDefOutputBufferCreateFilenameDefault func(
		f XmlOutputBufferCreateFilenameFunc) XmlOutputBufferCreateFilenameFunc

	XmlThrDefParserInputBufferCreateFilenameDefault func(
		f XmlParserInputBufferCreateFilenameFunc) XmlParserInputBufferCreateFilenameFunc

	DocbDefaultSAXHandler func() *XmlSAXHandlerV1

	HtmlDefaultSAXHandler func() *XmlSAXHandlerV1

	XmlLastError func() *XmlError

	OldXMLWDcompatibility func() *int

	XmlBufferAllocScheme func() *XmlBufferAllocationScheme

	XmlThrDefBufferAllocScheme func(
		v XmlBufferAllocationScheme) XmlBufferAllocationScheme

	XmlDefaultBufferSize func() *int

	XmlThrDefDefaultBufferSize func(v int) int

	XmlDefaultSAXHandler func() *XmlSAXHandlerV1

	XmlDefaultSAXLocator func() *XmlSAXLocator

	XmlDoValidityCheckingDefaultValue func() *int

	XmlThrDefDoValidityCheckingDefaultValue func(v int) int

	XmlGenericError func() *XmlGenericErrorFunc

	XmlStructuredError func() *XmlStructuredErrorFunc

	XmlGenericErrorContext func() **Void

	XmlStructuredErrorContext func() **Void

	XmlGetWarningsDefaultValue func() *int

	XmlThrDefGetWarningsDefaultValue func(v int) int

	XmlIndentTreeOutput func() *int

	XmlThrDefIndentTreeOutput func(v int) int

	XmlTreeIndentString func() **Char

	XmlThrDefTreeIndentString func(v *Char) *Char

	XmlKeepBlanksDefaultValue func() *int

	XmlThrDefKeepBlanksDefaultValue func(v int) int

	XmlLineNumbersDefaultValue func() *int

	XmlThrDefLineNumbersDefaultValue func(v int) int

	XmlLoadExtDtdDefaultValue func() *int

	XmlThrDefLoadExtDtdDefaultValue func(v int) int

	XmlParserDebugEntities func() *int

	XmlThrDefParserDebugEntities func(v int) int

	XmlParserVersion func() **Char

	XmlPedanticParserDefaultValue func() *int

	XmlThrDefPedanticParserDefaultValue func(v int) int

	XmlSaveNoEmptyTags func() *int

	XmlThrDefSaveNoEmptyTags func(v int) int

	XmlSubstituteEntitiesDefaultValue func() *int

	XmlThrDefSubstituteEntitiesDefaultValue func(v int) int

	XmlRegisterNodeDefaultValue func() *XmlRegisterNodeFunc

	XmlDeregisterNodeDefaultValue func() *XmlDeregisterNodeFunc

	XmlParserInputBufferCreateFilenameValue func() *XmlParserInputBufferCreateFilenameFunc

	XmlOutputBufferCreateFilenameValue func() *XmlOutputBufferCreateFilenameFunc

	XmlNewMutex func() XmlMutexPtr

	XmlMutexLock func(
		tok XmlMutexPtr)

	XmlMutexUnlock func(tok XmlMutexPtr)

	XmlFreeMutex func(tok XmlMutexPtr)

	XmlNewRMutex func() XmlRMutexPtr

	XmlRMutexLock func(
		tok XmlRMutexPtr)

	XmlRMutexUnlock func(
		tok XmlRMutexPtr)

	XmlFreeRMutex func(
		tok XmlRMutexPtr)

	XmlInitThreads func()

	XmlLockLibrary func()

	XmlUnlockLibrary func()

	XmlGetThreadId func() int

	XmlIsMainThread func() int

	XmlCleanupThreads func()

	XmlGetGlobalState func() XmlGlobalStatePtr

	HtmlTagLookup func(
		tag *XmlChar) *HtmlElemDesc

	HtmlEntityLookup func(name *XmlChar) *HtmlEntityDesc

	HtmlEntityValueLookup func(
		value Unsigned_int) *HtmlEntityDesc

	HtmlIsAutoClosed func(doc HtmlDocPtr, elem HtmlNodePtr) int

	HtmlAutoCloseTag func(
		doc HtmlDocPtr, name *XmlChar, elem HtmlNodePtr) int

	HtmlParseEntityRef func(
		ctxt HtmlParserCtxtPtr, str **XmlChar) *HtmlEntityDesc

	HtmlParseCharRef func(ctxt HtmlParserCtxtPtr) int

	HtmlParseElement func(ctxt HtmlParserCtxtPtr)

	HtmlNewParserCtxt func() HtmlParserCtxtPtr

	HtmlCreateMemoryParserCtxt func(
		buffer *Char, size int) HtmlParserCtxtPtr

	HtmlParseDocument func(ctxt HtmlParserCtxtPtr) int

	HtmlSAXParseDoc func(
		cur *XmlChar,
		encoding *Char,
		sax HtmlSAXHandlerPtr,
		userData *Void) HtmlDocPtr

	HtmlParseDoc func(
		cur *XmlChar,
		encoding *Char) HtmlDocPtr

	HtmlSAXParseFile func(
		filename *Char,
		encoding *Char,
		sax HtmlSAXHandlerPtr,
		userData *Void) HtmlDocPtr

	HtmlParseFile func(
		filename *Char,
		encoding *Char) HtmlDocPtr
	UTF8ToHtml func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int) int

	HtmlEncodeEntities func(
		out *Unsigned_char,
		outlen *int,
		in *Unsigned_char,
		inlen *int,
		quoteChar int) int

	HtmlIsScriptAttribute func(
		name *XmlChar) int

	HtmlHandleOmittedElem func(
		val int) int

	HtmlCreatePushParserCtxt func(
		sax HtmlSAXHandlerPtr,
		user_data *Void,
		chunk *Char,
		size int,
		filename *Char,
		enc XmlCharEncoding) HtmlParserCtxtPtr

	HtmlParseChunk func(
		ctxt HtmlParserCtxtPtr,
		chunk *Char,
		size int,
		terminate int) int

	HtmlFreeParserCtxt func(
		ctxt HtmlParserCtxtPtr)

	HtmlCtxtReset func(
		ctxt HtmlParserCtxtPtr)

	HtmlCtxtUseOptions func(
		ctxt HtmlParserCtxtPtr,
		options int) int

	HtmlReadDoc func(
		cur *XmlChar,
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlReadFile func(
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlReadMemory func(
		buffer *Char,
		size int,
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlReadFd func(
		fd int,
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlReadIO func(
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlCtxtReadDoc func(
		ctxt XmlParserCtxtPtr,
		cur *XmlChar,
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlCtxtReadFile func(
		ctxt XmlParserCtxtPtr,
		filename *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlCtxtReadMemory func(
		ctxt XmlParserCtxtPtr,
		buffer *Char,
		size int,
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlCtxtReadFd func(
		ctxt XmlParserCtxtPtr,
		fd int,
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlCtxtReadIO func(
		ctxt XmlParserCtxtPtr,
		ioread XmlInputReadCallback,
		ioclose XmlInputCloseCallback,
		ioctx *Void,
		URL *Char,
		encoding *Char,
		options int) HtmlDocPtr

	HtmlAttrAllowed func(
		*HtmlElemDesc,
		*XmlChar,
		int) HtmlStatus

	HtmlElementAllowedHere func(
		*HtmlElemDesc,
		*XmlChar) int

	HtmlElementStatusHere func(
		*HtmlElemDesc,
		*HtmlElemDesc) HtmlStatus

	HtmlNodeStatus func(
		HtmlNodePtr,
		int) HtmlStatus

	HtmlNewDoc func(
		URI *XmlChar,
		ExternalID *XmlChar) HtmlDocPtr

	HtmlNewDocNoDtD func(
		URI *XmlChar,
		ExternalID *XmlChar) HtmlDocPtr

	HtmlGetMetaEncoding func(
		doc HtmlDocPtr) *XmlChar

	HtmlSetMetaEncoding func(
		doc HtmlDocPtr,
		encoding *XmlChar) int

	HtmlDocDumpMemory func(
		cur XmlDocPtr,
		mem **XmlChar,
		size *int)

	HtmlDocDumpMemoryFormat func(
		cur XmlDocPtr,
		mem **XmlChar,
		size *int,
		format int)

	HtmlDocDump func(
		f *FILE,
		cur XmlDocPtr) int

	HtmlSaveFile func(
		filename *Char,
		cur XmlDocPtr) int

	HtmlNodeDump func(
		buf XmlBufferPtr,
		doc XmlDocPtr,
		cur XmlNodePtr) int

	HtmlNodeDumpFile func(
		out *FILE,
		doc XmlDocPtr,
		cur XmlNodePtr)

	HtmlNodeDumpFileFormat func(
		out *FILE,
		doc XmlDocPtr,
		cur XmlNodePtr,
		encoding *Char,
		format int) int

	HtmlSaveFileEnc func(
		filename *Char,
		cur XmlDocPtr,
		encoding *Char) int

	HtmlSaveFileFormat func(
		filename *Char,
		cur XmlDocPtr,
		encoding *Char,
		format int) int

	HtmlNodeDumpFormatOutput func(
		buf XmlOutputBufferPtr,
		doc XmlDocPtr,
		cur XmlNodePtr,
		encoding *Char,
		format int)

	HtmlDocContentDumpOutput func(
		buf XmlOutputBufferPtr,
		cur XmlDocPtr,
		encoding *Char)

	HtmlDocContentDumpFormatOutput func(
		buf XmlOutputBufferPtr,
		cur XmlDocPtr,
		encoding *Char,
		format int)

	HtmlNodeDumpOutput func(
		buf XmlOutputBufferPtr,
		doc XmlDocPtr,
		cur XmlNodePtr,
		encoding *Char)

	HtmlIsBooleanAttr func(
		name *XmlChar) int
)

var dll = "libxml2-2.dll"

var apiList = Apis{
	{"UTF8ToHtml", &UTF8ToHtml},
	{"UTF8Toisolat1", &UTF8Toisolat1},
	{"__docbDefaultSAXHandler", &DocbDefaultSAXHandler},
	{"__htmlDefaultSAXHandler", &HtmlDefaultSAXHandler},
	{"__oldXMLWDcompatibility", &OldXMLWDcompatibility},
	{"__xmlBufferAllocScheme", &XmlBufferAllocScheme},
	{"__xmlDefaultBufferSize", &XmlDefaultBufferSize},
	{"__xmlDefaultSAXHandler", &XmlDefaultSAXHandler},
	{"__xmlDefaultSAXLocator", &XmlDefaultSAXLocator},
	{"__xmlDeregisterNodeDefaultValue", &XmlDeregisterNodeDefaultValue},
	{"__xmlDoValidityCheckingDefaultValue", &XmlDoValidityCheckingDefaultValue},
	// {"__xmlErrEncoding", &XmlErrEncoding},
	{"__xmlGenericError", &XmlGenericError},
	{"__xmlGenericErrorContext", &XmlGenericErrorContext},
	{"__xmlGetWarningsDefaultValue", &XmlGetWarningsDefaultValue},
	{"__xmlIndentTreeOutput", &XmlIndentTreeOutput},
	{"__xmlKeepBlanksDefaultValue", &XmlKeepBlanksDefaultValue},
	{"__xmlLastError", &XmlLastError},
	{"__xmlLineNumbersDefaultValue", &XmlLineNumbersDefaultValue},
	{"__xmlLoadExtDtdDefaultValue", &XmlLoadExtDtdDefaultValue},
	{"__xmlOutputBufferCreateFilenameValue", &XmlOutputBufferCreateFilenameValue},
	{"__xmlParserDebugEntities", &XmlParserDebugEntities},
	{"__xmlParserInputBufferCreateFilenameValue", &XmlParserInputBufferCreateFilenameValue},
	{"__xmlParserVersion", &XmlParserVersion},
	{"__xmlPedanticParserDefaultValue", &XmlPedanticParserDefaultValue},
	// {"__xmlRaiseError", &XmlRaiseError},
	{"__xmlRegisterNodeDefaultValue", &XmlRegisterNodeDefaultValue},
	{"__xmlSaveNoEmptyTags", &XmlSaveNoEmptyTags},
	// {"__xmlSimpleError", &XmlSimpleError},
	{"__xmlStructuredError", &XmlStructuredError},
	{"__xmlStructuredErrorContext", &XmlStructuredErrorContext},
	{"__xmlSubstituteEntitiesDefaultValue", &XmlSubstituteEntitiesDefaultValue},
	{"__xmlTreeIndentString", &XmlTreeIndentString},
	{"attribute", &attribute},
	{"attributeDecl", &attributeDecl},
	{"cdataBlock", &cdataBlock},
	{"characters", &characters},
	{"checkNamespace", &checkNamespace},
	{"comment", &comment},
	// {"docbCreateFileParserCtxt", &DocbCreateFileParserCtxt},
	// {"docbCreatePushParserCtxt", &DocbCreatePushParserCtxt},
	// {"docbDefaultSAXHandlerInit", &DocbDefaultSAXHandlerInit},
	// {"docbEncodeEntities", &DocbEncodeEntities},
	// {"docbFreeParserCtxt", &DocbFreeParserCtxt},
	// {"docbParseChunk", &DocbParseChunk},
	// {"docbParseDoc", &DocbParseDoc},
	// {"docbParseDocument", &DocbParseDocument},
	// {"docbParseFile", &DocbParseFile},
	// {"docbSAXParseDoc", &DocbSAXParseDoc},
	// {"docbSAXParseFile", &DocbSAXParseFile},
	{"elementDecl", &elementDecl},
	// {"emptyExp", &emptyExp},
	{"endDocument", &endDocument},
	{"endElement", &endElement},
	{"entityDecl", &entityDecl},
	{"externalSubset", &externalSubset},
	// {"forbiddenExp", &forbiddenExp},
	{"getColumnNumber", &getColumnNumber},
	{"getEntity", &getEntity},
	{"getLineNumber", &getLineNumber},
	{"getNamespace", &getNamespace},
	{"getParameterEntity", &getParameterEntity},
	{"getPublicId", &getPublicId},
	{"getSystemId", &getSystemId},
	{"globalNamespace", &globalNamespace},
	{"hasExternalSubset", &hasExternalSubset},
	{"hasInternalSubset", &hasInternalSubset},
	{"htmlAttrAllowed", &HtmlAttrAllowed},
	{"htmlAutoCloseTag", &HtmlAutoCloseTag},
	// {"htmlCreateFileParserCtxt", &HtmlCreateFileParserCtxt},
	{"htmlCreateMemoryParserCtxt", &HtmlCreateMemoryParserCtxt},
	{"htmlCreatePushParserCtxt", &HtmlCreatePushParserCtxt},
	{"htmlCtxtReadDoc", &HtmlCtxtReadDoc},
	{"htmlCtxtReadFd", &HtmlCtxtReadFd},
	{"htmlCtxtReadFile", &HtmlCtxtReadFile},
	{"htmlCtxtReadIO", &HtmlCtxtReadIO},
	{"htmlCtxtReadMemory", &HtmlCtxtReadMemory},
	{"htmlCtxtReset", &HtmlCtxtReset},
	{"htmlCtxtUseOptions", &HtmlCtxtUseOptions},
	{"htmlDefaultSAXHandlerInit", &HtmlDefaultSAXHandlerInit},
	{"htmlDocContentDumpFormatOutput", &HtmlDocContentDumpFormatOutput},
	{"htmlDocContentDumpOutput", &HtmlDocContentDumpOutput},
	{"htmlDocDump", &HtmlDocDump},
	{"htmlDocDumpMemory", &HtmlDocDumpMemory},
	{"htmlDocDumpMemoryFormat", &HtmlDocDumpMemoryFormat},
	{"htmlElementAllowedHere", &HtmlElementAllowedHere},
	{"htmlElementStatusHere", &HtmlElementStatusHere},
	{"htmlEncodeEntities", &HtmlEncodeEntities},
	{"htmlEntityLookup", &HtmlEntityLookup},
	{"htmlEntityValueLookup", &HtmlEntityValueLookup},
	{"htmlFreeParserCtxt", &HtmlFreeParserCtxt},
	{"htmlGetMetaEncoding", &HtmlGetMetaEncoding},
	{"htmlHandleOmittedElem", &HtmlHandleOmittedElem},
	// {"htmlInitAutoClose", &HtmlInitAutoClose},
	{"htmlIsAutoClosed", &HtmlIsAutoClosed},
	{"htmlIsBooleanAttr", &HtmlIsBooleanAttr},
	{"htmlIsScriptAttribute", &HtmlIsScriptAttribute},
	{"htmlNewDoc", &HtmlNewDoc},
	{"htmlNewDocNoDtD", &HtmlNewDocNoDtD},
	{"htmlNewParserCtxt", &HtmlNewParserCtxt},
	{"htmlNodeDump", &HtmlNodeDump},
	{"htmlNodeDumpFile", &HtmlNodeDumpFile},
	{"htmlNodeDumpFileFormat", &HtmlNodeDumpFileFormat},
	{"htmlNodeDumpFormatOutput", &HtmlNodeDumpFormatOutput},
	{"htmlNodeDumpOutput", &HtmlNodeDumpOutput},
	{"htmlNodeStatus", &HtmlNodeStatus},
	{"htmlParseCharRef", &HtmlParseCharRef},
	{"htmlParseChunk", &HtmlParseChunk},
	{"htmlParseDoc", &HtmlParseDoc},
	{"htmlParseDocument", &HtmlParseDocument},
	{"htmlParseElement", &HtmlParseElement},
	{"htmlParseEntityRef", &HtmlParseEntityRef},
	{"htmlParseFile", &HtmlParseFile},
	{"htmlReadDoc", &HtmlReadDoc},
	{"htmlReadFd", &HtmlReadFd},
	{"htmlReadFile", &HtmlReadFile},
	{"htmlReadIO", &HtmlReadIO},
	{"htmlReadMemory", &HtmlReadMemory},
	{"htmlSAXParseDoc", &HtmlSAXParseDoc},
	{"htmlSAXParseFile", &HtmlSAXParseFile},
	{"htmlSaveFile", &HtmlSaveFile},
	{"htmlSaveFileEnc", &HtmlSaveFileEnc},
	{"htmlSaveFileFormat", &HtmlSaveFileFormat},
	{"htmlSetMetaEncoding", &HtmlSetMetaEncoding},
	{"htmlTagLookup", &HtmlTagLookup},
	{"ignorableWhitespace", &IgnorableWhitespace},
	{"initGenericErrorDefaultFunc", &InitGenericErrorDefaultFunc},
	{"initdocbDefaultSAXHandler", &InitdocbDefaultSAXHandler},
	{"inithtmlDefaultSAXHandler", &InithtmlDefaultSAXHandler},
	// {"initxmlDefaultSAXHandler", &InitxmlDefaultSAXHandler},
	// {"inputPop", &InputPop},
	// {"inputPush", &InputPush},
	{"internalSubset", &InternalSubset},
	{"isStandalone", &IsStandalone},
	{"isolat1ToUTF8", &Isolat1ToUTF8},
	// {"namePop", &NamePop},
	// {"namePush", &NamePush},
	// {"namespaceDecl", &NamespaceDecl},
	// {"nodePop", &NodePop},
	// {"nodePush", &NodePush},
	// {"notationDecl", &NotationDecl},
	{"processingInstruction", &processingInstruction},
	{"reference", &reference},
	{"resolveEntity", &resolveEntity},
	{"setDocumentLocator", &setDocumentLocator},
	{"setNamespace", &setNamespace},
	{"startDocument", &startDocument},
	{"startElement", &startElement},
	{"unparsedEntityDecl", &UnparsedEntityDecl},
	// {"valuePop", &ValuePop},
	// {"valuePush", &ValuePush},
	{"xlinkGetDefaultDetect", &xlinkGetDefaultDetect},
	{"xlinkGetDefaultHandler", &xlinkGetDefaultHandler},
	{"xlinkIsLink", &xlinkIsLink},
	{"xlinkSetDefaultDetect", &xlinkSetDefaultDetect},
	{"xlinkSetDefaultHandler", &xlinkSetDefaultHandler},
	// {"xmlACatalogAdd", &XmlACatalogAdd},
	// {"xmlACatalogDump", &XmlACatalogDump},
	// {"xmlACatalogRemove", &XmlACatalogRemove},
	// {"xmlACatalogResolve", &XmlACatalogResolve},
	// {"xmlACatalogResolvePublic", &XmlACatalogResolvePublic},
	// {"xmlACatalogResolveSystem", &XmlACatalogResolveSystem},
	// {"xmlACatalogResolveURI", &XmlACatalogResolveURI},
	{"xmlAddAttributeDecl", &XmlAddAttributeDecl},
	{"xmlAddChild", &XmlAddChild},
	{"xmlAddChildList", &XmlAddChildList},
	{"xmlAddDocEntity", &XmlAddDocEntity},
	{"xmlAddDtdEntity", &XmlAddDtdEntity},
	{"xmlAddElementDecl", &XmlAddElementDecl},
	{"xmlAddEncodingAlias", &XmlAddEncodingAlias},
	{"xmlAddID", &XmlAddID},
	{"xmlAddNextSibling", &XmlAddNextSibling},
	{"xmlAddNotationDecl", &XmlAddNotationDecl},
	{"xmlAddPrevSibling", &XmlAddPrevSibling},
	{"xmlAddRef", &XmlAddRef},
	{"xmlAddSibling", &XmlAddSibling},
	{"xmlAllocOutputBuffer", &XmlAllocOutputBuffer},
	{"xmlAllocParserInputBuffer", &XmlAllocParserInputBuffer},
	{"xmlAttrSerializeTxtContent", &XmlAttrSerializeTxtContent},
	{"xmlAutomataCompile", &XmlAutomataCompile},
	{"xmlAutomataGetInitState", &XmlAutomataGetInitState},
	{"xmlAutomataIsDeterminist", &XmlAutomataIsDeterminist},
	{"xmlAutomataNewAllTrans", &XmlAutomataNewAllTrans},
	{"xmlAutomataNewCountTrans", &XmlAutomataNewCountTrans},
	{"xmlAutomataNewCountTrans2", &XmlAutomataNewCountTrans2},
	{"xmlAutomataNewCountedTrans", &XmlAutomataNewCountedTrans},
	{"xmlAutomataNewCounter", &XmlAutomataNewCounter},
	{"xmlAutomataNewCounterTrans", &XmlAutomataNewCounterTrans},
	{"xmlAutomataNewEpsilon", &XmlAutomataNewEpsilon},
	{"xmlAutomataNewNegTrans", &XmlAutomataNewNegTrans},
	{"xmlAutomataNewOnceTrans", &XmlAutomataNewOnceTrans},
	{"xmlAutomataNewOnceTrans2", &XmlAutomataNewOnceTrans2},
	{"xmlAutomataNewState", &XmlAutomataNewState},
	{"xmlAutomataNewTransition", &XmlAutomataNewTransition},
	{"xmlAutomataNewTransition2", &XmlAutomataNewTransition2},
	{"xmlAutomataSetFinalState", &XmlAutomataSetFinalState},
	// {"xmlBoolToText", &XmlBoolToText},
	{"xmlBufferAdd", &XmlBufferAdd},
	{"xmlBufferAddHead", &XmlBufferAddHead},
	{"xmlBufferCCat", &XmlBufferCCat},
	{"xmlBufferCat", &XmlBufferCat},
	{"xmlBufferContent", &XmlBufferContent},
	{"xmlBufferCreate", &XmlBufferCreate},
	{"xmlBufferCreateSize", &XmlBufferCreateSize},
	{"xmlBufferCreateStatic", &XmlBufferCreateStatic},
	{"xmlBufferDump", &XmlBufferDump},
	{"xmlBufferEmpty", &XmlBufferEmpty},
	{"xmlBufferFree", &XmlBufferFree},
	{"xmlBufferGrow", &XmlBufferGrow},
	{"xmlBufferLength", &XmlBufferLength},
	{"xmlBufferResize", &XmlBufferResize},
	{"xmlBufferSetAllocationScheme", &XmlBufferSetAllocationScheme},
	{"xmlBufferShrink", &XmlBufferShrink},
	{"xmlBufferWriteCHAR", &XmlBufferWriteCHAR},
	{"xmlBufferWriteChar", &XmlBufferWriteChar},
	{"xmlBufferWriteQuotedString", &XmlBufferWriteQuotedString},
	{"xmlBuildQName", &XmlBuildQName},
	// {"xmlBuildRelativeURI", &XmlBuildRelativeURI},
	// {"xmlBuildURI", &XmlBuildURI},
	{"xmlByteConsumed", &XmlByteConsumed},
	// {"xmlC14NDocDumpMemory", &XmlC14NDocDumpMemory},
	// {"xmlC14NDocSave", &XmlC14NDocSave},
	// {"xmlC14NDocSaveTo", &XmlC14NDocSaveTo},
	// {"xmlC14NExecute", &XmlC14NExecute},
	// {"xmlCanonicPath", &XmlCanonicPath},
	// {"xmlCatalogAdd", &XmlCatalogAdd},
	// {"xmlCatalogAddLocal", &XmlCatalogAddLocal},
	// {"xmlCatalogCleanup", &XmlCatalogCleanup},
	// {"xmlCatalogConvert", &XmlCatalogConvert},
	// {"xmlCatalogDump", &XmlCatalogDump},
	// {"xmlCatalogFreeLocal", &XmlCatalogFreeLocal},
	// {"xmlCatalogGetDefaults", &XmlCatalogGetDefaults},
	// {"xmlCatalogGetPublic", &XmlCatalogGetPublic},
	// {"xmlCatalogGetSystem", &XmlCatalogGetSystem},
	// {"xmlCatalogIsEmpty", &XmlCatalogIsEmpty},
	// {"xmlCatalogLocalResolve", &XmlCatalogLocalResolve},
	// {"xmlCatalogLocalResolveURI", &XmlCatalogLocalResolveURI},
	// {"xmlCatalogRemove", &XmlCatalogRemove},
	// {"xmlCatalogResolve", &XmlCatalogResolve},
	// {"xmlCatalogResolvePublic", &XmlCatalogResolvePublic},
	// {"xmlCatalogResolveSystem", &XmlCatalogResolveSystem},
	// {"xmlCatalogResolveURI", &XmlCatalogResolveURI},
	// {"xmlCatalogSetDebug", &XmlCatalogSetDebug},
	// {"xmlCatalogSetDefaultPrefer", &XmlCatalogSetDefaultPrefer},
	// {"xmlCatalogSetDefaults", &XmlCatalogSetDefaults},
	{"xmlCharEncCloseFunc", &XmlCharEncCloseFunc},
	{"xmlCharEncFirstLine", &XmlCharEncFirstLine},
	{"xmlCharEncInFunc", &XmlCharEncInFunc},
	{"xmlCharEncOutFunc", &XmlCharEncOutFunc},
	// {"xmlCharInRange", &XmlCharInRange},
	{"xmlCharStrdup", &XmlCharStrdup},
	{"xmlCharStrndup", &XmlCharStrndup},
	{"xmlCheckFilename", &XmlCheckFilename},
	{"xmlCheckHTTPInput", &XmlCheckHTTPInput},
	// {"xmlCheckLanguageID", &XmlCheckLanguageID},
	{"xmlCheckUTF8", &XmlCheckUTF8},
	{"xmlCheckVersion", &XmlCheckVersion},
	{"xmlChildElementCount", &XmlChildElementCount},
	{"xmlCleanupCharEncodingHandlers", &XmlCleanupCharEncodingHandlers},
	{"xmlCleanupEncodingAliases", &XmlCleanupEncodingAliases},
	{"xmlCleanupGlobals", &XmlCleanupGlobals},
	{"xmlCleanupInputCallbacks", &XmlCleanupInputCallbacks},
	{"xmlCleanupMemory", &XmlCleanupMemory},
	{"xmlCleanupOutputCallbacks", &XmlCleanupOutputCallbacks},
	{"xmlCleanupParser", &XmlCleanupParser},
	{"xmlCleanupPredefinedEntities", &XmlCleanupPredefinedEntities},
	{"xmlCleanupThreads", &XmlCleanupThreads},
	{"xmlClearNodeInfoSeq", &XmlClearNodeInfoSeq},
	{"xmlClearParserCtxt", &XmlClearParserCtxt},
	// {"xmlConvertSGMLCatalog", &XmlConvertSGMLCatalog},
	{"xmlCopyAttributeTable", &XmlCopyAttributeTable},
	// {"xmlCopyChar", &XmlCopyChar},
	// {"xmlCopyCharMultiByte", &XmlCopyCharMultiByte},
	{"xmlCopyDoc", &XmlCopyDoc},
	{"xmlCopyDocElementContent", &XmlCopyDocElementContent},
	{"xmlCopyDtd", &XmlCopyDtd},
	{"xmlCopyElementContent", &XmlCopyElementContent},
	{"xmlCopyElementTable", &XmlCopyElementTable},
	{"xmlCopyEntitiesTable", &XmlCopyEntitiesTable},
	{"xmlCopyEnumeration", &XmlCopyEnumeration},
	{"xmlCopyError", &XmlCopyError},
	{"xmlCopyNamespace", &XmlCopyNamespace},
	{"xmlCopyNamespaceList", &XmlCopyNamespaceList},
	{"xmlCopyNode", &XmlCopyNode},
	{"xmlCopyNodeList", &XmlCopyNodeList},
	{"xmlCopyNotationTable", &XmlCopyNotationTable},
	{"xmlCopyProp", &XmlCopyProp},
	{"xmlCopyPropList", &XmlCopyPropList},
	{"xmlCreateDocParserCtxt", &XmlCreateDocParserCtxt},
	{"xmlCreateEntitiesTable", &XmlCreateEntitiesTable},
	// {"xmlCreateEntityParserCtxt", &XmlCreateEntityParserCtxt},
	{"xmlCreateEnumeration", &XmlCreateEnumeration},
	// {"xmlCreateFileParserCtxt", &XmlCreateFileParserCtxt},
	{"xmlCreateIOParserCtxt", &XmlCreateIOParserCtxt},
	{"xmlCreateIntSubset", &XmlCreateIntSubset},
	// {"xmlCreateMemoryParserCtxt", &XmlCreateMemoryParserCtxt},
	{"xmlCreatePushParserCtxt", &XmlCreatePushParserCtxt},
	// {"xmlCreateURI", &XmlCreateURI},
	// {"xmlCreateURLParserCtxt", &XmlCreateURLParserCtxt},
	{"xmlCtxtGetLastError", &XmlCtxtGetLastError},
	{"xmlCtxtReadDoc", &XmlCtxtReadDoc},
	{"xmlCtxtReadFd", &XmlCtxtReadFd},
	{"xmlCtxtReadFile", &XmlCtxtReadFile},
	{"xmlCtxtReadIO", &XmlCtxtReadIO},
	{"xmlCtxtReadMemory", &XmlCtxtReadMemory},
	{"xmlCtxtReset", &XmlCtxtReset},
	{"xmlCtxtResetLastError", &XmlCtxtResetLastError},
	{"xmlCtxtResetPush", &XmlCtxtResetPush},
	{"xmlCtxtUseOptions", &XmlCtxtUseOptions},
	// {"xmlCurrentChar", &XmlCurrentChar},
	{"xmlDOMWrapAdoptNode", &XmlDOMWrapAdoptNode},
	{"xmlDOMWrapCloneNode", &XmlDOMWrapCloneNode},
	{"xmlDOMWrapFreeCtxt", &XmlDOMWrapFreeCtxt},
	{"xmlDOMWrapNewCtxt", &XmlDOMWrapNewCtxt},
	{"xmlDOMWrapReconcileNamespaces", &XmlDOMWrapReconcileNamespaces},
	{"xmlDOMWrapRemoveNode", &XmlDOMWrapRemoveNode},
	// {"xmlDebugCheckDocument", &XmlDebugCheckDocument},
	// {"xmlDebugDumpAttr", &XmlDebugDumpAttr},
	// {"xmlDebugDumpAttrList", &XmlDebugDumpAttrList},
	// {"xmlDebugDumpDTD", &XmlDebugDumpDTD},
	// {"xmlDebugDumpDocument", &XmlDebugDumpDocument},
	// {"xmlDebugDumpDocumentHead", &XmlDebugDumpDocumentHead},
	// {"xmlDebugDumpEntities", &XmlDebugDumpEntities},
	// {"xmlDebugDumpNode", &XmlDebugDumpNode},
	// {"xmlDebugDumpNodeList", &XmlDebugDumpNodeList},
	// {"xmlDebugDumpOneNode", &XmlDebugDumpOneNode},
	// {"xmlDebugDumpString", &XmlDebugDumpString},
	// {"xmlDecodeEntities", &XmlDecodeEntities},
	{"xmlDefaultSAXHandlerInit", &XmlDefaultSAXHandlerInit},
	{"xmlDelEncodingAlias", &XmlDelEncodingAlias},
	{"xmlDeregisterNodeDefault", &XmlDeregisterNodeDefault},
	{"xmlDetectCharEncoding", &XmlDetectCharEncoding},
	{"xmlDictCleanup", &XmlDictCleanup},
	{"xmlDictCreate", &XmlDictCreate},
	{"xmlDictCreateSub", &XmlDictCreateSub},
	{"xmlDictExists", &XmlDictExists},
	{"xmlDictFree", &XmlDictFree},
	{"xmlDictLookup", &XmlDictLookup},
	{"xmlDictOwns", &XmlDictOwns},
	{"xmlDictQLookup", &XmlDictQLookup},
	{"xmlDictReference", &XmlDictReference},
	{"xmlDictSize", &XmlDictSize},
	{"xmlDocCopyNode", &XmlDocCopyNode},
	{"xmlDocCopyNodeList", &XmlDocCopyNodeList},
	{"xmlDocDump", &XmlDocDump},
	{"xmlDocDumpFormatMemory", &XmlDocDumpFormatMemory},
	{"xmlDocDumpFormatMemoryEnc", &XmlDocDumpFormatMemoryEnc},
	{"xmlDocDumpMemory", &XmlDocDumpMemory},
	{"xmlDocDumpMemoryEnc", &XmlDocDumpMemoryEnc},
	{"xmlDocFormatDump", &XmlDocFormatDump},
	{"xmlDocGetRootElement", &XmlDocGetRootElement},
	{"xmlDocSetRootElement", &XmlDocSetRootElement},
	{"xmlDumpAttributeDecl", &XmlDumpAttributeDecl},
	{"xmlDumpAttributeTable", &XmlDumpAttributeTable},
	{"xmlDumpElementDecl", &XmlDumpElementDecl},
	{"xmlDumpElementTable", &XmlDumpElementTable},
	{"xmlDumpEntitiesTable", &XmlDumpEntitiesTable},
	{"xmlDumpEntityDecl", &XmlDumpEntityDecl},
	{"xmlDumpNotationDecl", &XmlDumpNotationDecl},
	{"xmlDumpNotationTable", &XmlDumpNotationTable},
	{"xmlElemDump", &XmlElemDump},
	{"xmlEncodeEntities", &XmlEncodeEntities},
	{"xmlEncodeEntitiesReentrant", &XmlEncodeEntitiesReentrant},
	{"xmlEncodeSpecialChars", &XmlEncodeSpecialChars},
	// {"xmlErrMemory", &XmlErrMemory},
	{"xmlExpCtxtNbCons", &XmlExpCtxtNbCons},
	{"xmlExpCtxtNbNodes", &XmlExpCtxtNbNodes},
	{"xmlExpDump", &XmlExpDump},
	{"xmlExpExpDerive", &XmlExpExpDerive},
	{"xmlExpFree", &XmlExpFree},
	{"xmlExpFreeCtxt", &XmlExpFreeCtxt},
	{"xmlExpGetLanguage", &XmlExpGetLanguage},
	{"xmlExpGetStart", &XmlExpGetStart},
	{"xmlExpIsNillable", &XmlExpIsNillable},
	{"xmlExpMaxToken", &XmlExpMaxToken},
	{"xmlExpNewAtom", &XmlExpNewAtom},
	{"xmlExpNewCtxt", &XmlExpNewCtxt},
	{"xmlExpNewOr", &XmlExpNewOr},
	{"xmlExpNewRange", &XmlExpNewRange},
	{"xmlExpNewSeq", &XmlExpNewSeq},
	{"xmlExpParse", &XmlExpParse},
	{"xmlExpRef", &XmlExpRef},
	{"xmlExpStringDerive", &XmlExpStringDerive},
	{"xmlExpSubsume", &XmlExpSubsume},
	{"xmlFileClose", &XmlFileClose},
	{"xmlFileMatch", &XmlFileMatch},
	{"xmlFileOpen", &XmlFileOpen},
	{"xmlFileRead", &XmlFileRead},
	{"xmlFindCharEncodingHandler", &XmlFindCharEncodingHandler},
	{"xmlFirstElementChild", &XmlFirstElementChild},
	// {"xmlFree", &XmlFree},
	{"xmlFreeAttributeTable", &XmlFreeAttributeTable},
	{"xmlFreeAutomata", &XmlFreeAutomata},
	// {"xmlFreeCatalog", &XmlFreeCatalog},
	{"xmlFreeDoc", &XmlFreeDoc},
	{"xmlFreeDocElementContent", &XmlFreeDocElementContent},
	{"xmlFreeDtd", &XmlFreeDtd},
	{"xmlFreeElementContent", &XmlFreeElementContent},
	{"xmlFreeElementTable", &XmlFreeElementTable},
	{"xmlFreeEntitiesTable", &XmlFreeEntitiesTable},
	{"xmlFreeEnumeration", &XmlFreeEnumeration},
	{"xmlFreeIDTable", &XmlFreeIDTable},
	// {"xmlFreeInputStream", &XmlFreeInputStream},
	{"xmlFreeMutex", &XmlFreeMutex},
	{"xmlFreeNode", &XmlFreeNode},
	{"xmlFreeNodeList", &XmlFreeNodeList},
	{"xmlFreeNotationTable", &XmlFreeNotationTable},
	{"xmlFreeNs", &XmlFreeNs},
	{"xmlFreeNsList", &XmlFreeNsList},
	{"xmlFreeParserCtxt", &XmlFreeParserCtxt},
	{"xmlFreeParserInputBuffer", &XmlFreeParserInputBuffer},
	// {"xmlFreePattern", &XmlFreePattern},
	// {"xmlFreePatternList", &XmlFreePatternList},
	{"xmlFreeProp", &XmlFreeProp},
	{"xmlFreePropList", &XmlFreePropList},
	{"xmlFreeRMutex", &XmlFreeRMutex},
	{"xmlFreeRefTable", &XmlFreeRefTable},
	// {"xmlFreeStreamCtxt", &XmlFreeStreamCtxt},
	// {"xmlFreeTextReader", &XmlFreeTextReader},
	// {"xmlFreeTextWriter", &XmlFreeTextWriter},
	// {"xmlFreeURI", &XmlFreeURI},
	{"xmlFreeValidCtxt", &XmlFreeValidCtxt},
	{"xmlGcMemGet", &XmlGcMemGet},
	{"xmlGcMemSetup", &XmlGcMemSetup},
	{"xmlGetBufferAllocationScheme", &XmlGetBufferAllocationScheme},
	{"xmlGetCharEncodingHandler", &XmlGetCharEncodingHandler},
	{"xmlGetCharEncodingName", &XmlGetCharEncodingName},
	{"xmlGetCompressMode", &XmlGetCompressMode},
	{"xmlGetDocCompressMode", &XmlGetDocCompressMode},
	{"xmlGetDocEntity", &XmlGetDocEntity},
	{"xmlGetDtdAttrDesc", &XmlGetDtdAttrDesc},
	{"xmlGetDtdElementDesc", &XmlGetDtdElementDesc},
	{"xmlGetDtdEntity", &XmlGetDtdEntity},
	{"xmlGetDtdNotationDesc", &XmlGetDtdNotationDesc},
	{"xmlGetDtdQAttrDesc", &XmlGetDtdQAttrDesc},
	{"xmlGetDtdQElementDesc", &XmlGetDtdQElementDesc},
	{"xmlGetEncodingAlias", &XmlGetEncodingAlias},
	{"xmlGetExternalEntityLoader", &XmlGetExternalEntityLoader},
	{"xmlGetFeature", &XmlGetFeature},
	{"xmlGetFeaturesList", &XmlGetFeaturesList},
	{"xmlGetGlobalState", &XmlGetGlobalState},
	{"xmlGetID", &XmlGetID},
	{"xmlGetIntSubset", &XmlGetIntSubset},
	{"xmlGetLastChild", &XmlGetLastChild},
	{"xmlGetLastError", &XmlGetLastError},
	{"xmlGetLineNo", &XmlGetLineNo},
	{"xmlGetNoNsProp", &XmlGetNoNsProp},
	{"xmlGetNodePath", &XmlGetNodePath},
	{"xmlGetNsList", &XmlGetNsList},
	{"xmlGetNsProp", &XmlGetNsProp},
	{"xmlGetParameterEntity", &XmlGetParameterEntity},
	{"xmlGetPredefinedEntity", &XmlGetPredefinedEntity},
	{"xmlGetProp", &XmlGetProp},
	{"xmlGetRefs", &XmlGetRefs},
	{"xmlGetThreadId", &XmlGetThreadId},
	{"xmlGetUTF8Char", &XmlGetUTF8Char},
	// {"xmlHandleEntity", &XmlHandleEntity},
	{"xmlHasFeature", &XmlHasFeature},
	{"xmlHasNsProp", &XmlHasNsProp},
	{"xmlHasProp", &XmlHasProp},
	{"xmlHashAddEntry", &XmlHashAddEntry},
	{"xmlHashAddEntry2", &XmlHashAddEntry2},
	{"xmlHashAddEntry3", &XmlHashAddEntry3},
	{"xmlHashCopy", &XmlHashCopy},
	{"xmlHashCreate", &XmlHashCreate},
	{"xmlHashCreateDict", &XmlHashCreateDict},
	{"xmlHashFree", &XmlHashFree},
	{"xmlHashLookup", &XmlHashLookup},
	{"xmlHashLookup2", &XmlHashLookup2},
	{"xmlHashLookup3", &XmlHashLookup3},
	{"xmlHashQLookup", &XmlHashQLookup},
	{"xmlHashQLookup2", &XmlHashQLookup2},
	{"xmlHashQLookup3", &XmlHashQLookup3},
	{"xmlHashRemoveEntry", &XmlHashRemoveEntry},
	{"xmlHashRemoveEntry2", &XmlHashRemoveEntry2},
	{"xmlHashRemoveEntry3", &XmlHashRemoveEntry3},
	{"xmlHashScan", &XmlHashScan},
	{"xmlHashScan3", &XmlHashScan3},
	{"xmlHashScanFull", &XmlHashScanFull},
	{"xmlHashScanFull3", &XmlHashScanFull3},
	{"xmlHashSize", &XmlHashSize},
	{"xmlHashUpdateEntry", &XmlHashUpdateEntry},
	{"xmlHashUpdateEntry2", &XmlHashUpdateEntry2},
	{"xmlHashUpdateEntry3", &XmlHashUpdateEntry3},
	{"xmlIOFTPClose", &XmlIOFTPClose},
	{"xmlIOFTPMatch", &XmlIOFTPMatch},
	{"xmlIOFTPOpen", &XmlIOFTPOpen},
	{"xmlIOFTPRead", &XmlIOFTPRead},
	{"xmlIOHTTPClose", &XmlIOHTTPClose},
	{"xmlIOHTTPMatch", &XmlIOHTTPMatch},
	{"xmlIOHTTPOpen", &XmlIOHTTPOpen},
	{"xmlIOHTTPOpenW", &XmlIOHTTPOpenW},
	{"xmlIOHTTPRead", &XmlIOHTTPRead},
	{"xmlIOParseDTD", &XmlIOParseDTD},
	{"xmlInitCharEncodingHandlers", &XmlInitCharEncodingHandlers},
	{"xmlInitGlobals", &XmlInitGlobals},
	{"xmlInitMemory", &XmlInitMemory},
	{"xmlInitNodeInfoSeq", &XmlInitNodeInfoSeq},
	{"xmlInitParser", &XmlInitParser},
	{"xmlInitParserCtxt", &XmlInitParserCtxt},
	{"xmlInitThreads", &XmlInitThreads},
	// {"xmlInitializeCatalog", &XmlInitializeCatalog},
	{"xmlInitializeGlobalState", &XmlInitializeGlobalState},
	{"xmlInitializePredefinedEntities", &XmlInitializePredefinedEntities},
	// {"xmlIsBaseChar", &XmlIsBaseChar},
	// {"xmlIsBaseCharGroup", &XmlIsBaseCharGroup},
	// {"xmlIsBlank", &XmlIsBlank},
	{"xmlIsBlankNode", &XmlIsBlankNode},
	// {"xmlIsChar", &XmlIsChar},
	// {"xmlIsCharGroup", &XmlIsCharGroup},
	// {"xmlIsCombining", &XmlIsCombining},
	// {"xmlIsCombiningGroup", &XmlIsCombiningGroup},
	// {"xmlIsDigit", &XmlIsDigit},
	// {"xmlIsDigitGroup", &XmlIsDigitGroup},
	// {"xmlIsExtender", &XmlIsExtender},
	// {"xmlIsExtenderGroup", &XmlIsExtenderGroup},
	{"xmlIsID", &XmlIsID},
	// {"xmlIsIdeographic", &XmlIsIdeographic},
	// {"xmlIsIdeographicGroup", &XmlIsIdeographicGroup},
	// {"xmlIsLetter", &XmlIsLetter},
	{"xmlIsMainThread", &XmlIsMainThread},
	{"xmlIsMixedElement", &XmlIsMixedElement},
	// {"xmlIsPubidChar", &XmlIsPubidChar},
	// {"xmlIsPubidChar_tab", &XmlIsPubidChar_tab},
	{"xmlIsRef", &XmlIsRef},
	{"xmlIsXHTML", &XmlIsXHTML},
	{"xmlKeepBlanksDefault", &XmlKeepBlanksDefault},
	{"xmlLastElementChild", &XmlLastElementChild},
	{"xmlLineNumbersDefault", &XmlLineNumbersDefault},
	{"xmlLinkGetData", &XmlLinkGetData},
	{"xmlListAppend", &XmlListAppend},
	{"xmlListClear", &XmlListClear},
	{"xmlListCopy", &XmlListCopy},
	{"xmlListCreate", &XmlListCreate},
	{"xmlListDelete", &XmlListDelete},
	{"xmlListDup", &XmlListDup},
	{"xmlListEmpty", &XmlListEmpty},
	{"xmlListEnd", &XmlListEnd},
	{"xmlListFront", &XmlListFront},
	{"xmlListInsert", &XmlListInsert},
	{"xmlListMerge", &XmlListMerge},
	{"xmlListPopBack", &XmlListPopBack},
	{"xmlListPopFront", &XmlListPopFront},
	{"xmlListPushBack", &XmlListPushBack},
	{"xmlListPushFront", &XmlListPushFront},
	{"xmlListRemoveAll", &XmlListRemoveAll},
	{"xmlListRemoveFirst", &XmlListRemoveFirst},
	{"xmlListRemoveLast", &XmlListRemoveLast},
	{"xmlListReverse", &XmlListReverse},
	{"xmlListReverseSearch", &XmlListReverseSearch},
	{"xmlListReverseWalk", &XmlListReverseWalk},
	{"xmlListSearch", &XmlListSearch},
	{"xmlListSize", &XmlListSize},
	{"xmlListSort", &XmlListSort},
	{"xmlListWalk", &XmlListWalk},
	// {"xmlLoadACatalog", &XmlLoadACatalog},
	// {"xmlLoadCatalog", &XmlLoadCatalog},
	// {"xmlLoadCatalogs", &XmlLoadCatalogs},
	{"xmlLoadExternalEntity", &XmlLoadExternalEntity},
	// {"xmlLoadSGMLSuperCatalog", &XmlLoadSGMLSuperCatalog},
	{"xmlLockLibrary", &XmlLockLibrary},
	// {"xmlLsCountNode", &XmlLsCountNode},
	// {"xmlLsOneNode", &XmlLsOneNode},
	// {"xmlMalloc", &XmlMalloc},
	// {"xmlMallocAtomic", &XmlMallocAtomic},
	{"xmlMallocAtomicLoc", &XmlMallocAtomicLoc},
	{"xmlMallocLoc", &XmlMallocLoc},
	{"xmlMemBlocks", &XmlMemBlocks},
	{"xmlMemDisplay", &XmlMemDisplay},
	{"xmlMemDisplayLast", &XmlMemDisplayLast},
	{"xmlMemFree", &XmlMemFree},
	{"xmlMemGet", &XmlMemGet},
	{"xmlMemMalloc", &XmlMemMalloc},
	{"xmlMemRealloc", &XmlMemRealloc},
	{"xmlMemSetup", &XmlMemSetup},
	{"xmlMemShow", &XmlMemShow},
	// {"xmlMemStrdup", &XmlMemStrdup},
	{"xmlMemStrdupLoc", &XmlMemStrdupLoc},
	{"xmlMemUsed", &XmlMemUsed},
	{"xmlMemoryDump", &XmlMemoryDump},
	{"xmlMemoryStrdup", &XmlMemoryStrdup},
	// {"xmlModuleClose", &XmlModuleClose},
	// {"xmlModuleFree", &XmlModuleFree},
	// {"xmlModuleOpen", &XmlModuleOpen},
	// {"xmlModuleSymbol", &XmlModuleSymbol},
	{"xmlMutexLock", &XmlMutexLock},
	{"xmlMutexUnlock", &XmlMutexUnlock},
	// {"xmlNamespaceParseNCName", &XmlNamespaceParseNCName},
	// {"xmlNamespaceParseNSDef", &XmlNamespaceParseNSDef},
	// {"xmlNamespaceParseQName", &XmlNamespaceParseQName},
	// {"xmlNanoFTPCheckResponse", &XmlNanoFTPCheckResponse},
	// {"xmlNanoFTPCleanup", &XmlNanoFTPCleanup},
	// {"xmlNanoFTPClose", &XmlNanoFTPClose},
	// {"xmlNanoFTPCloseConnection", &XmlNanoFTPCloseConnection},
	// {"xmlNanoFTPConnect", &XmlNanoFTPConnect},
	// {"xmlNanoFTPConnectTo", &XmlNanoFTPConnectTo},
	// {"xmlNanoFTPCwd", &XmlNanoFTPCwd},
	// {"xmlNanoFTPDele", &XmlNanoFTPDele},
	// {"xmlNanoFTPFreeCtxt", &XmlNanoFTPFreeCtxt},
	// {"xmlNanoFTPGet", &XmlNanoFTPGet},
	// {"xmlNanoFTPGetConnection", &XmlNanoFTPGetConnection},
	// {"xmlNanoFTPGetResponse", &XmlNanoFTPGetResponse},
	// {"xmlNanoFTPGetSocket", &XmlNanoFTPGetSocket},
	// {"xmlNanoFTPInit", &XmlNanoFTPInit},
	// {"xmlNanoFTPList", &XmlNanoFTPList},
	// {"xmlNanoFTPNewCtxt", &XmlNanoFTPNewCtxt},
	// {"xmlNanoFTPOpen", &XmlNanoFTPOpen},
	// {"xmlNanoFTPProxy", &XmlNanoFTPProxy},
	// {"xmlNanoFTPQuit", &XmlNanoFTPQuit},
	// {"xmlNanoFTPRead", &XmlNanoFTPRead},
	// {"xmlNanoFTPScanProxy", &XmlNanoFTPScanProxy},
	// {"xmlNanoFTPUpdateURL", &XmlNanoFTPUpdateURL},
	// {"xmlNanoHTTPAuthHeader", &XmlNanoHTTPAuthHeader},
	// {"xmlNanoHTTPCleanup", &XmlNanoHTTPCleanup},
	// {"xmlNanoHTTPClose", &XmlNanoHTTPClose},
	// {"xmlNanoHTTPContentLength", &XmlNanoHTTPContentLength},
	// {"xmlNanoHTTPEncoding", &XmlNanoHTTPEncoding},
	// {"xmlNanoHTTPFetch", &XmlNanoHTTPFetch},
	// {"xmlNanoHTTPInit", &XmlNanoHTTPInit},
	// {"xmlNanoHTTPMethod", &XmlNanoHTTPMethod},
	// {"xmlNanoHTTPMethodRedir", &XmlNanoHTTPMethodRedir},
	// {"xmlNanoHTTPMimeType", &XmlNanoHTTPMimeType},
	// {"xmlNanoHTTPOpen", &XmlNanoHTTPOpen},
	// {"xmlNanoHTTPOpenRedir", &XmlNanoHTTPOpenRedir},
	// {"xmlNanoHTTPRead", &XmlNanoHTTPRead},
	// {"xmlNanoHTTPRedir", &XmlNanoHTTPRedir},
	// {"xmlNanoHTTPReturnCode", &XmlNanoHTTPReturnCode},
	// {"xmlNanoHTTPSave", &XmlNanoHTTPSave},
	// {"xmlNanoHTTPScanProxy", &XmlNanoHTTPScanProxy},
	{"xmlNewAutomata", &XmlNewAutomata},
	{"xmlNewCDataBlock", &XmlNewCDataBlock},
	// {"xmlNewCatalog", &XmlNewCatalog},
	{"xmlNewCharEncodingHandler", &XmlNewCharEncodingHandler},
	{"xmlNewCharRef", &XmlNewCharRef},
	{"xmlNewChild", &XmlNewChild},
	{"xmlNewComment", &XmlNewComment},
	{"xmlNewDoc", &XmlNewDoc},
	{"xmlNewDocComment", &XmlNewDocComment},
	{"xmlNewDocElementContent", &XmlNewDocElementContent},
	{"xmlNewDocFragment", &XmlNewDocFragment},
	{"xmlNewDocNode", &XmlNewDocNode},
	{"xmlNewDocNodeEatName", &XmlNewDocNodeEatName},
	{"xmlNewDocPI", &XmlNewDocPI},
	{"xmlNewDocProp", &XmlNewDocProp},
	{"xmlNewDocRawNode", &XmlNewDocRawNode},
	{"xmlNewDocText", &XmlNewDocText},
	{"xmlNewDocTextLen", &XmlNewDocTextLen},
	{"xmlNewDtd", &XmlNewDtd},
	{"xmlNewElementContent", &XmlNewElementContent},
	{"xmlNewEntity", &XmlNewEntity},
	// {"xmlNewEntityInputStream", &XmlNewEntityInputStream},
	{"xmlNewGlobalNs", &XmlNewGlobalNs},
	{"xmlNewIOInputStream", &XmlNewIOInputStream},
	// {"xmlNewInputFromFile", &XmlNewInputFromFile},
	// {"xmlNewInputStream", &XmlNewInputStream},
	{"xmlNewMutex", &XmlNewMutex},
	{"xmlNewNode", &XmlNewNode},
	{"xmlNewNodeEatName", &XmlNewNodeEatName},
	{"xmlNewNs", &XmlNewNs},
	{"xmlNewNsProp", &XmlNewNsProp},
	{"xmlNewNsPropEatName", &XmlNewNsPropEatName},
	{"xmlNewPI", &XmlNewPI},
	{"xmlNewParserCtxt", &XmlNewParserCtxt},
	{"xmlNewProp", &XmlNewProp},
	{"xmlNewRMutex", &XmlNewRMutex},
	{"xmlNewReference", &XmlNewReference},
	// {"xmlNewStringInputStream", &XmlNewStringInputStream},
	{"xmlNewText", &XmlNewText},
	{"xmlNewTextChild", &XmlNewTextChild},
	{"xmlNewTextLen", &XmlNewTextLen},
	// {"xmlNewTextReader", &XmlNewTextReader},
	// {"xmlNewTextReaderFilename", &XmlNewTextReaderFilename},
	// {"xmlNewTextWriter", &XmlNewTextWriter},
	// {"xmlNewTextWriterDoc", &XmlNewTextWriterDoc},
	// {"xmlNewTextWriterFilename", &XmlNewTextWriterFilename},
	// {"xmlNewTextWriterMemory", &XmlNewTextWriterMemory},
	// {"xmlNewTextWriterPushParser", &XmlNewTextWriterPushParser},
	// {"xmlNewTextWriterTree", &XmlNewTextWriterTree},
	// {"xmlNewValidCtxt", &XmlNewValidCtxt},
	// {"xmlNextChar", &XmlNextChar},
	{"xmlNextElementSibling", &XmlNextElementSibling},
	{"xmlNoNetExternalEntityLoader", &XmlNoNetExternalEntityLoader},
	{"xmlNodeAddContent", &XmlNodeAddContent},
	{"xmlNodeAddContentLen", &XmlNodeAddContentLen},
	{"xmlNodeBufGetContent", &XmlNodeBufGetContent},
	{"xmlNodeDump", &XmlNodeDump},
	{"xmlNodeDumpOutput", &XmlNodeDumpOutput},
	{"xmlNodeGetBase", &XmlNodeGetBase},
	{"xmlNodeGetContent", &XmlNodeGetContent},
	{"xmlNodeGetLang", &XmlNodeGetLang},
	{"xmlNodeGetSpacePreserve", &XmlNodeGetSpacePreserve},
	{"xmlNodeIsText", &XmlNodeIsText},
	{"xmlNodeListGetRawString", &XmlNodeListGetRawString},
	{"xmlNodeListGetString", &XmlNodeListGetString},
	{"xmlNodeSetBase", &XmlNodeSetBase},
	{"xmlNodeSetContent", &XmlNodeSetContent},
	{"xmlNodeSetContentLen", &XmlNodeSetContentLen},
	{"xmlNodeSetLang", &XmlNodeSetLang},
	{"xmlNodeSetName", &XmlNodeSetName},
	{"xmlNodeSetSpacePreserve", &XmlNodeSetSpacePreserve},
	// {"xmlNormalizeURIPath", &XmlNormalizeURIPath},
	{"xmlNormalizeWindowsPath", &XmlNormalizeWindowsPath},
	{"xmlOutputBufferClose", &XmlOutputBufferClose},
	{"xmlOutputBufferCreateBuffer", &XmlOutputBufferCreateBuffer},
	{"xmlOutputBufferCreateFd", &XmlOutputBufferCreateFd},
	{"xmlOutputBufferCreateFile", &XmlOutputBufferCreateFile},
	{"xmlOutputBufferCreateFilename", &XmlOutputBufferCreateFilename},
	// {"xmlOutputBufferCreateFilenameDefault", &XmlOutputBufferCreateFilenameDefault},
	{"xmlOutputBufferCreateIO", &XmlOutputBufferCreateIO},
	{"xmlOutputBufferFlush", &XmlOutputBufferFlush},
	{"xmlOutputBufferWrite", &XmlOutputBufferWrite},
	{"xmlOutputBufferWriteEscape", &XmlOutputBufferWriteEscape},
	{"xmlOutputBufferWriteString", &XmlOutputBufferWriteString},
	// {"xmlParseAttValue", &XmlParseAttValue},
	// {"xmlParseAttribute", &XmlParseAttribute},
	// {"xmlParseAttributeListDecl", &XmlParseAttributeListDecl},
	// {"xmlParseAttributeType", &XmlParseAttributeType},
	{"xmlParseBalancedChunkMemory", &XmlParseBalancedChunkMemory},
	{"xmlParseBalancedChunkMemoryRecover", &XmlParseBalancedChunkMemoryRecover},
	// {"xmlParseCDSect", &XmlParseCDSect},
	// {"xmlParseCatalogFile", &XmlParseCatalogFile},
	// {"xmlParseCharData", &XmlParseCharData},
	{"xmlParseCharEncoding", &XmlParseCharEncoding},
	// {"xmlParseCharRef", &XmlParseCharRef},
	{"xmlParseChunk", &XmlParseChunk},
	// {"xmlParseComment", &XmlParseComment},
	// {"xmlParseContent", &XmlParseContent},
	{"xmlParseCtxtExternalEntity", &XmlParseCtxtExternalEntity},
	{"xmlParseDTD", &XmlParseDTD},
	// {"xmlParseDefaultDecl", &XmlParseDefaultDecl},
	{"xmlParseDoc", &XmlParseDoc},
	// {"xmlParseDocTypeDecl", &XmlParseDocTypeDecl},
	{"xmlParseDocument", &XmlParseDocument},
	// {"xmlParseElement", &XmlParseElement},
	// {"xmlParseElementChildrenContentDecl", &XmlParseElementChildrenContentDecl},
	// {"xmlParseElementContentDecl", &XmlParseElementContentDecl},
	// {"xmlParseElementDecl", &XmlParseElementDecl},
	// {"xmlParseElementMixedContentDecl", &XmlParseElementMixedContentDecl},
	// {"xmlParseEncName", &XmlParseEncName},
	// {"xmlParseEncodingDecl", &XmlParseEncodingDecl},
	// {"xmlParseEndTag", &XmlParseEndTag},
	{"xmlParseEntity", &XmlParseEntity},
	// {"xmlParseEntityDecl", &XmlParseEntityDecl},
	// {"xmlParseEntityRef", &XmlParseEntityRef},
	// {"xmlParseEntityValue", &XmlParseEntityValue},
	// {"xmlParseEnumeratedType", &XmlParseEnumeratedType},
	// {"xmlParseEnumerationType", &XmlParseEnumerationType},
	{"xmlParseExtParsedEnt", &XmlParseExtParsedEnt},
	{"xmlParseExternalEntity", &XmlParseExternalEntity},
	// {"xmlParseExternalID", &XmlParseExternalID},
	// {"xmlParseExternalSubset", &XmlParseExternalSubset},
	{"xmlParseFile", &XmlParseFile},
	{"xmlParseInNodeContext", &XmlParseInNodeContext},
	// {"xmlParseMarkupDecl", &XmlParseMarkupDecl},
	{"xmlParseMemory", &XmlParseMemory},
	// {"xmlParseMisc", &XmlParseMisc},
	// {"xmlParseName", &XmlParseName},
	// {"xmlParseNamespace", &XmlParseNamespace},
	// {"xmlParseNmtoken", &XmlParseNmtoken},
	// {"xmlParseNotationDecl", &XmlParseNotationDecl},
	// {"xmlParseNotationType", &XmlParseNotationType},
	// {"xmlParsePEReference", &XmlParsePEReference},
	// {"xmlParsePI", &XmlParsePI},
	// {"xmlParsePITarget", &XmlParsePITarget},
	// {"xmlParsePubidLiteral", &XmlParsePubidLiteral},
	// {"xmlParseQuotedString", &XmlParseQuotedString},
	// {"xmlParseReference", &XmlParseReference},
	// {"xmlParseSDDecl", &XmlParseSDDecl},
	// {"xmlParseStartTag", &XmlParseStartTag},
	// {"xmlParseSystemLiteral", &XmlParseSystemLiteral},
	// {"xmlParseTextDecl", &XmlParseTextDecl},
	// {"xmlParseURI", &XmlParseURI},
	// {"xmlParseURIRaw", &XmlParseURIRaw},
	// {"xmlParseURIReference", &XmlParseURIReference},
	// {"xmlParseVersionInfo", &XmlParseVersionInfo},
	// {"xmlParseVersionNum", &XmlParseVersionNum},
	// {"xmlParseXMLDecl", &XmlParseXMLDecl},
	{"xmlParserAddNodeInfo", &XmlParserAddNodeInfo},
	{"xmlParserError", &XmlParserError},
	{"xmlParserFindNodeInfo", &XmlParserFindNodeInfo},
	{"xmlParserFindNodeInfoIndex", &XmlParserFindNodeInfoIndex},
	{"xmlParserGetDirectory", &XmlParserGetDirectory},
	// {"xmlParserHandlePEReference", &XmlParserHandlePEReference},
	// {"xmlParserHandleReference", &XmlParserHandleReference},
	{"xmlParserInputBufferCreateFd", &XmlParserInputBufferCreateFd},
	{"xmlParserInputBufferCreateFile", &XmlParserInputBufferCreateFile},
	{"xmlParserInputBufferCreateFilename", &XmlParserInputBufferCreateFilename},
	// {"xmlParserInputBufferCreateFilenameDefault", &XmlParserInputBufferCreateFilenameDefault},
	{"xmlParserInputBufferCreateIO", &XmlParserInputBufferCreateIO},
	{"xmlParserInputBufferCreateMem", &XmlParserInputBufferCreateMem},
	{"xmlParserInputBufferCreateStatic", &XmlParserInputBufferCreateStatic},
	{"xmlParserInputBufferGrow", &XmlParserInputBufferGrow},
	{"xmlParserInputBufferPush", &XmlParserInputBufferPush},
	{"xmlParserInputBufferRead", &XmlParserInputBufferRead},
	{"xmlParserInputGrow", &XmlParserInputGrow},
	{"xmlParserInputRead", &XmlParserInputRead},
	// {"xmlParserInputShrink", &XmlParserInputShrink},
	// {"xmlParserMaxDepth", &XmlParserMaxDepth},
	{"xmlParserPrintFileContext", &XmlParserPrintFileContext},
	{"xmlParserPrintFileInfo", &XmlParserPrintFileInfo},
	{"xmlParserValidityError", &XmlParserValidityError},
	{"xmlParserValidityWarning", &XmlParserValidityWarning},
	{"xmlParserWarning", &XmlParserWarning},
	// {"xmlPathToURI", &XmlPathToURI},
	// {"xmlPatternFromRoot", &XmlPatternFromRoot},
	// {"xmlPatternGetStreamCtxt", &XmlPatternGetStreamCtxt},
	// {"xmlPatternMatch", &XmlPatternMatch},
	// {"xmlPatternMaxDepth", &XmlPatternMaxDepth},
	// {"xmlPatternMinDepth", &XmlPatternMinDepth},
	// {"xmlPatternStreamable", &XmlPatternStreamable},
	// {"xmlPatterncompile", &XmlPatterncompile},
	{"xmlPedanticParserDefault", &XmlPedanticParserDefault},
	// {"xmlPopInput", &XmlPopInput},
	{"xmlPopInputCallbacks", &XmlPopInputCallbacks},
	{"xmlPreviousElementSibling", &XmlPreviousElementSibling},
	// {"xmlPrintURI", &XmlPrintURI},
	// {"xmlPushInput", &XmlPushInput},
	{"xmlRMutexLock", &XmlRMutexLock},
	{"xmlRMutexUnlock", &XmlRMutexUnlock},
	{"xmlReadDoc", &XmlReadDoc},
	{"xmlReadFd", &XmlReadFd},
	{"xmlReadFile", &XmlReadFile},
	{"xmlReadIO", &XmlReadIO},
	{"xmlReadMemory", &XmlReadMemory},
	// {"xmlReaderForDoc", &XmlReaderForDoc},
	// {"xmlReaderForFd", &XmlReaderForFd},
	// {"xmlReaderForFile", &XmlReaderForFile},
	// {"xmlReaderForIO", &XmlReaderForIO},
	// {"xmlReaderForMemory", &XmlReaderForMemory},
	// {"xmlReaderNewDoc", &XmlReaderNewDoc},
	// {"xmlReaderNewFd", &XmlReaderNewFd},
	// {"xmlReaderNewFile", &XmlReaderNewFile},
	// {"xmlReaderNewIO", &XmlReaderNewIO},
	// {"xmlReaderNewMemory", &XmlReaderNewMemory},
	// {"xmlReaderNewWalker", &XmlReaderNewWalker},
	// {"xmlReaderWalker", &XmlReaderWalker},
	// {"xmlRealloc", &XmlRealloc},
	{"xmlReallocLoc", &XmlReallocLoc},
	{"xmlReconciliateNs", &XmlReconciliateNs},
	{"xmlRecoverDoc", &XmlRecoverDoc},
	{"xmlRecoverFile", &XmlRecoverFile},
	{"xmlRecoverMemory", &XmlRecoverMemory},
	{"xmlRegExecErrInfo", &XmlRegExecErrInfo},
	{"xmlRegExecNextValues", &XmlRegExecNextValues},
	{"xmlRegExecPushString", &XmlRegExecPushString},
	{"xmlRegExecPushString2", &XmlRegExecPushString2},
	{"xmlRegFreeExecCtxt", &XmlRegFreeExecCtxt},
	{"xmlRegFreeRegexp", &XmlRegFreeRegexp},
	{"xmlRegNewExecCtxt", &XmlRegNewExecCtxt},
	{"xmlRegexpCompile", &XmlRegexpCompile},
	{"xmlRegexpExec", &XmlRegexpExec},
	{"xmlRegexpIsDeterminist", &XmlRegexpIsDeterminist},
	{"xmlRegexpPrint", &XmlRegexpPrint},
	{"xmlRegisterCharEncodingHandler", &XmlRegisterCharEncodingHandler},
	{"xmlRegisterDefaultInputCallbacks", &XmlRegisterDefaultInputCallbacks},
	{"xmlRegisterDefaultOutputCallbacks", &XmlRegisterDefaultOutputCallbacks},
	{"xmlRegisterHTTPPostCallbacks", &XmlRegisterHTTPPostCallbacks},
	{"xmlRegisterInputCallbacks", &XmlRegisterInputCallbacks},
	{"xmlRegisterNodeDefault", &XmlRegisterNodeDefault},
	{"xmlRegisterOutputCallbacks", &XmlRegisterOutputCallbacks},
	// {"xmlRelaxNGCleanupTypes", &XmlRelaxNGCleanupTypes},
	// {"xmlRelaxNGDump", &XmlRelaxNGDump},
	// {"xmlRelaxNGDumpTree", &XmlRelaxNGDumpTree},
	// {"xmlRelaxNGFree", &XmlRelaxNGFree},
	// {"xmlRelaxNGFreeParserCtxt", &XmlRelaxNGFreeParserCtxt},
	// {"xmlRelaxNGFreeValidCtxt", &XmlRelaxNGFreeValidCtxt},
	// {"xmlRelaxNGGetParserErrors", &XmlRelaxNGGetParserErrors},
	// {"xmlRelaxNGGetValidErrors", &XmlRelaxNGGetValidErrors},
	// {"xmlRelaxNGInitTypes", &XmlRelaxNGInitTypes},
	// {"xmlRelaxNGNewDocParserCtxt", &XmlRelaxNGNewDocParserCtxt},
	// {"xmlRelaxNGNewMemParserCtxt", &XmlRelaxNGNewMemParserCtxt},
	// {"xmlRelaxNGNewParserCtxt", &XmlRelaxNGNewParserCtxt},
	// {"xmlRelaxNGNewValidCtxt", &XmlRelaxNGNewValidCtxt},
	// {"xmlRelaxNGParse", &XmlRelaxNGParse},
	// {"xmlRelaxNGSetParserErrors", &XmlRelaxNGSetParserErrors},
	// {"xmlRelaxNGSetParserStructuredErrors", &XmlRelaxNGSetParserStructuredErrors},
	// {"xmlRelaxNGSetValidErrors", &XmlRelaxNGSetValidErrors},
	// {"xmlRelaxNGSetValidStructuredErrors", &XmlRelaxNGSetValidStructuredErrors},
	// {"xmlRelaxNGValidateDoc", &XmlRelaxNGValidateDoc},
	// {"xmlRelaxNGValidateFullElement", &XmlRelaxNGValidateFullElement},
	// {"xmlRelaxNGValidatePopElement", &XmlRelaxNGValidatePopElement},
	// {"xmlRelaxNGValidatePushCData", &XmlRelaxNGValidatePushCData},
	// {"xmlRelaxNGValidatePushElement", &XmlRelaxNGValidatePushElement},
	// {"xmlRelaxParserSetFlag", &XmlRelaxParserSetFlag},
	{"xmlRemoveID", &XmlRemoveID},
	{"xmlRemoveProp", &XmlRemoveProp},
	{"xmlRemoveRef", &XmlRemoveRef},
	{"xmlReplaceNode", &XmlReplaceNode},
	{"xmlResetError", &XmlResetError},
	{"xmlResetLastError", &XmlResetLastError},
	{"xmlSAX2AttributeDecl", &XmlSAX2AttributeDecl},
	{"xmlSAX2CDataBlock", &XmlSAX2CDataBlock},
	{"xmlSAX2Characters", &XmlSAX2Characters},
	{"xmlSAX2Comment", &XmlSAX2Comment},
	{"xmlSAX2ElementDecl", &XmlSAX2ElementDecl},
	{"xmlSAX2EndDocument", &XmlSAX2EndDocument},
	{"xmlSAX2EndElement", &XmlSAX2EndElement},
	{"xmlSAX2EndElementNs", &XmlSAX2EndElementNs},
	{"xmlSAX2EntityDecl", &XmlSAX2EntityDecl},
	{"xmlSAX2ExternalSubset", &XmlSAX2ExternalSubset},
	{"xmlSAX2GetColumnNumber", &XmlSAX2GetColumnNumber},
	{"xmlSAX2GetEntity", &XmlSAX2GetEntity},
	{"xmlSAX2GetLineNumber", &XmlSAX2GetLineNumber},
	{"xmlSAX2GetParameterEntity", &XmlSAX2GetParameterEntity},
	{"xmlSAX2GetPublicId", &XmlSAX2GetPublicId},
	{"xmlSAX2GetSystemId", &XmlSAX2GetSystemId},
	{"xmlSAX2HasExternalSubset", &XmlSAX2HasExternalSubset},
	{"xmlSAX2HasInternalSubset", &XmlSAX2HasInternalSubset},
	{"xmlSAX2IgnorableWhitespace", &XmlSAX2IgnorableWhitespace},
	{"xmlSAX2InitDefaultSAXHandler", &XmlSAX2InitDefaultSAXHandler},
	{"xmlSAX2InitDocbDefaultSAXHandler", &XmlSAX2InitDocbDefaultSAXHandler},
	{"xmlSAX2InitHtmlDefaultSAXHandler", &XmlSAX2InitHtmlDefaultSAXHandler},
	{"xmlSAX2InternalSubset", &XmlSAX2InternalSubset},
	{"xmlSAX2IsStandalone", &XmlSAX2IsStandalone},
	{"xmlSAX2NotationDecl", &XmlSAX2NotationDecl},
	{"xmlSAX2ProcessingInstruction", &XmlSAX2ProcessingInstruction},
	{"xmlSAX2Reference", &XmlSAX2Reference},
	{"xmlSAX2ResolveEntity", &XmlSAX2ResolveEntity},
	{"xmlSAX2SetDocumentLocator", &XmlSAX2SetDocumentLocator},
	{"xmlSAX2StartDocument", &XmlSAX2StartDocument},
	{"xmlSAX2StartElement", &XmlSAX2StartElement},
	{"xmlSAX2StartElementNs", &XmlSAX2StartElementNs},
	{"xmlSAX2UnparsedEntityDecl", &XmlSAX2UnparsedEntityDecl},
	{"xmlSAXDefaultVersion", &XmlSAXDefaultVersion},
	{"xmlSAXParseDTD", &XmlSAXParseDTD},
	{"xmlSAXParseDoc", &XmlSAXParseDoc},
	{"xmlSAXParseEntity", &XmlSAXParseEntity},
	{"xmlSAXParseFile", &XmlSAXParseFile},
	{"xmlSAXParseFileWithData", &XmlSAXParseFileWithData},
	{"xmlSAXParseMemory", &XmlSAXParseMemory},
	{"xmlSAXParseMemoryWithData", &XmlSAXParseMemoryWithData},
	{"xmlSAXUserParseFile", &XmlSAXUserParseFile},
	{"xmlSAXUserParseMemory", &XmlSAXUserParseMemory},
	{"xmlSAXVersion", &XmlSAXVersion},
	// {"xmlSaveClose", &XmlSaveClose},
	// {"xmlSaveDoc", &XmlSaveDoc},
	{"xmlSaveFile", &XmlSaveFile},
	{"xmlSaveFileEnc", &XmlSaveFileEnc},
	{"xmlSaveFileTo", &XmlSaveFileTo},
	// {"xmlSaveFlush", &XmlSaveFlush},
	{"xmlSaveFormatFile", &XmlSaveFormatFile},
	{"xmlSaveFormatFileEnc", &XmlSaveFormatFileEnc},
	{"xmlSaveFormatFileTo", &XmlSaveFormatFileTo},
	// {"xmlSaveSetAttrEscape", &XmlSaveSetAttrEscape},
	// {"xmlSaveSetEscape", &XmlSaveSetEscape},
	// {"xmlSaveToBuffer", &XmlSaveToBuffer},
	// {"xmlSaveToFd", &XmlSaveToFd},
	// {"xmlSaveToFilename", &XmlSaveToFilename},
	// {"xmlSaveToIO", &XmlSaveToIO},
	// {"xmlSaveTree", &XmlSaveTree},
	// {"xmlSaveUri", &XmlSaveUri},
	// {"xmlScanName", &XmlScanName},
	// {"xmlSchemaCheckFacet", &XmlSchemaCheckFacet},
	// {"xmlSchemaCleanupTypes", &XmlSchemaCleanupTypes},
	// {"xmlSchemaCollapseString", &XmlSchemaCollapseString},
	// {"xmlSchemaCompareValues", &XmlSchemaCompareValues},
	// {"xmlSchemaCompareValuesWhtsp", &XmlSchemaCompareValuesWhtsp},
	// {"xmlSchemaCopyValue", &XmlSchemaCopyValue},
	// {"xmlSchemaDump", &XmlSchemaDump},
	// {"xmlSchemaFree", &XmlSchemaFree},
	// {"xmlSchemaFreeFacet", &XmlSchemaFreeFacet},
	// {"xmlSchemaFreeParserCtxt", &XmlSchemaFreeParserCtxt},
	// {"xmlSchemaFreeType", &XmlSchemaFreeType},
	// {"xmlSchemaFreeValidCtxt", &XmlSchemaFreeValidCtxt},
	// {"xmlSchemaFreeValue", &XmlSchemaFreeValue},
	// {"xmlSchemaFreeWildcard", &XmlSchemaFreeWildcard},
	// {"xmlSchemaGetBuiltInListSimpleTypeItemType", &XmlSchemaGetBuiltInListSimpleTypeItemType},
	// {"xmlSchemaGetBuiltInType", &XmlSchemaGetBuiltInType},
	// {"xmlSchemaGetCanonValue", &XmlSchemaGetCanonValue},
	// {"xmlSchemaGetCanonValueWhtsp", &XmlSchemaGetCanonValueWhtsp},
	// {"xmlSchemaGetFacetValueAsULong", &XmlSchemaGetFacetValueAsULong},
	// {"xmlSchemaGetParserErrors", &XmlSchemaGetParserErrors},
	// {"xmlSchemaGetPredefinedType", &XmlSchemaGetPredefinedType},
	// {"xmlSchemaGetValType", &XmlSchemaGetValType},
	// {"xmlSchemaGetValidErrors", &XmlSchemaGetValidErrors},
	// {"xmlSchemaInitTypes", &XmlSchemaInitTypes},
	// {"xmlSchemaIsBuiltInTypeFacet", &XmlSchemaIsBuiltInTypeFacet},
	// {"xmlSchemaIsValid", &XmlSchemaIsValid},
	// {"xmlSchemaNewDocParserCtxt", &XmlSchemaNewDocParserCtxt},
	// {"xmlSchemaNewFacet", &XmlSchemaNewFacet},
	// {"xmlSchemaNewMemParserCtxt", &XmlSchemaNewMemParserCtxt},
	// {"xmlSchemaNewNOTATIONValue", &XmlSchemaNewNOTATIONValue},
	// {"xmlSchemaNewParserCtxt", &XmlSchemaNewParserCtxt},
	// {"xmlSchemaNewQNameValue", &XmlSchemaNewQNameValue},
	// {"xmlSchemaNewStringValue", &XmlSchemaNewStringValue},
	// {"xmlSchemaNewValidCtxt", &XmlSchemaNewValidCtxt},
	// {"xmlSchemaParse", &XmlSchemaParse},
	// {"xmlSchemaSAXPlug", &XmlSchemaSAXPlug},
	// {"xmlSchemaSAXUnplug", &XmlSchemaSAXUnplug},
	// {"xmlSchemaSetParserErrors", &XmlSchemaSetParserErrors},
	// {"xmlSchemaSetParserStructuredErrors", &XmlSchemaSetParserStructuredErrors},
	// {"xmlSchemaSetValidErrors", &XmlSchemaSetValidErrors},
	// {"xmlSchemaSetValidOptions", &XmlSchemaSetValidOptions},
	// {"xmlSchemaSetValidStructuredErrors", &XmlSchemaSetValidStructuredErrors},
	// {"xmlSchemaValPredefTypeNode", &XmlSchemaValPredefTypeNode},
	// {"xmlSchemaValPredefTypeNodeNoNorm", &XmlSchemaValPredefTypeNodeNoNorm},
	// {"xmlSchemaValidCtxtGetOptions", &XmlSchemaValidCtxtGetOptions},
	// {"xmlSchemaValidCtxtGetParserCtxt", &XmlSchemaValidCtxtGetParserCtxt},
	// {"xmlSchemaValidateDoc", &XmlSchemaValidateDoc},
	// {"xmlSchemaValidateFacet", &XmlSchemaValidateFacet},
	// {"xmlSchemaValidateFacetWhtsp", &XmlSchemaValidateFacetWhtsp},
	// {"xmlSchemaValidateFile", &XmlSchemaValidateFile},
	// {"xmlSchemaValidateLengthFacet", &XmlSchemaValidateLengthFacet},
	// {"xmlSchemaValidateLengthFacetWhtsp", &XmlSchemaValidateLengthFacetWhtsp},
	// {"xmlSchemaValidateListSimpleTypeFacet", &XmlSchemaValidateListSimpleTypeFacet},
	// {"xmlSchemaValidateOneElement", &XmlSchemaValidateOneElement},
	// {"xmlSchemaValidatePredefinedType", &XmlSchemaValidatePredefinedType},
	// {"xmlSchemaValidateStream", &XmlSchemaValidateStream},
	// {"xmlSchemaValueAppend", &XmlSchemaValueAppend},
	// {"xmlSchemaValueGetAsBoolean", &XmlSchemaValueGetAsBoolean},
	// {"xmlSchemaValueGetAsString", &XmlSchemaValueGetAsString},
	// {"xmlSchemaValueGetNext", &XmlSchemaValueGetNext},
	// {"xmlSchemaWhiteSpaceReplace", &XmlSchemaWhiteSpaceReplace},
	// {"xmlSchematronFree", &XmlSchematronFree},
	// {"xmlSchematronFreeParserCtxt", &XmlSchematronFreeParserCtxt},
	// {"xmlSchematronFreeValidCtxt", &XmlSchematronFreeValidCtxt},
	// {"xmlSchematronNewDocParserCtxt", &XmlSchematronNewDocParserCtxt},
	// {"xmlSchematronNewMemParserCtxt", &XmlSchematronNewMemParserCtxt},
	// {"xmlSchematronNewParserCtxt", &XmlSchematronNewParserCtxt},
	// {"xmlSchematronNewValidCtxt", &XmlSchematronNewValidCtxt},
	// {"xmlSchematronParse", &XmlSchematronParse},
	// {"xmlSchematronSetValidStructuredErrors", &XmlSchematronSetValidStructuredErrors},
	// {"xmlSchematronValidateDoc", &XmlSchematronValidateDoc},
	{"xmlSearchNs", &XmlSearchNs},
	{"xmlSearchNsByHref", &XmlSearchNsByHref},
	{"xmlSetBufferAllocationScheme", &XmlSetBufferAllocationScheme},
	{"xmlSetCompressMode", &XmlSetCompressMode},
	{"xmlSetDocCompressMode", &XmlSetDocCompressMode},
	// {"xmlSetEntityReferenceFunc", &XmlSetEntityReferenceFunc},
	{"xmlSetExternalEntityLoader", &XmlSetExternalEntityLoader},
	{"xmlSetFeature", &XmlSetFeature},
	{"xmlSetGenericErrorFunc", &XmlSetGenericErrorFunc},
	{"xmlSetListDoc", &XmlSetListDoc},
	{"xmlSetNs", &XmlSetNs},
	{"xmlSetNsProp", &XmlSetNsProp},
	{"xmlSetProp", &XmlSetProp},
	{"xmlSetStructuredErrorFunc", &XmlSetStructuredErrorFunc},
	{"xmlSetTreeDoc", &XmlSetTreeDoc},
	{"xmlSetupParserForBuffer", &XmlSetupParserForBuffer},
	// {"xmlShell", &XmlShell},
	// {"xmlShellBase", &XmlShellBase},
	// {"xmlShellCat", &XmlShellCat},
	// {"xmlShellDir", &XmlShellDir},
	// {"xmlShellDu", &XmlShellDu},
	// {"xmlShellList", &XmlShellList},
	// {"xmlShellLoad", &XmlShellLoad},
	// {"xmlShellPrintNode", &XmlShellPrintNode},
	// {"xmlShellPrintXPathError", &XmlShellPrintXPathError},
	// {"xmlShellPrintXPathResult", &XmlShellPrintXPathResult},
	// {"xmlShellPwd", &XmlShellPwd},
	// {"xmlShellSave", &XmlShellSave},
	// {"xmlShellValidate", &XmlShellValidate},
	// {"xmlShellWrite", &XmlShellWrite},
	// {"xmlSkipBlankChars", &XmlSkipBlankChars},
	{"xmlSnprintfElementContent", &XmlSnprintfElementContent},
	// {"xmlSplitQName", &XmlSplitQName},
	{"xmlSplitQName2", &XmlSplitQName2},
	{"xmlSplitQName3", &XmlSplitQName3},
	{"xmlSprintfElementContent", &XmlSprintfElementContent},
	{"xmlStopParser", &XmlStopParser},
	{"xmlStrEqual", &XmlStrEqual},
	{"xmlStrPrintf", &XmlStrPrintf},
	{"xmlStrQEqual", &XmlStrQEqual},
	{"xmlStrVPrintf", &XmlStrVPrintf},
	{"xmlStrcasecmp", &XmlStrcasecmp},
	{"xmlStrcasestr", &XmlStrcasestr},
	{"xmlStrcat", &XmlStrcat},
	{"xmlStrchr", &XmlStrchr},
	{"xmlStrcmp", &XmlStrcmp},
	{"xmlStrdup", &XmlStrdup},
	// {"xmlStreamPop", &XmlStreamPop},
	// {"xmlStreamPush", &XmlStreamPush},
	// {"xmlStreamPushAttr", &XmlStreamPushAttr},
	// {"xmlStreamPushNode", &XmlStreamPushNode},
	// {"xmlStreamWantsAnyNode", &XmlStreamWantsAnyNode},
	// {"xmlStringComment", &XmlStringComment},
	// {"xmlStringCurrentChar", &XmlStringCurrentChar},
	// {"xmlStringDecodeEntities", &XmlStringDecodeEntities},
	{"xmlStringGetNodeList", &XmlStringGetNodeList},
	// {"xmlStringLenDecodeEntities", &XmlStringLenDecodeEntities},
	{"xmlStringLenGetNodeList", &XmlStringLenGetNodeList},
	// {"xmlStringText", &XmlStringText},
	// {"xmlStringTextNoenc", &XmlStringTextNoenc},
	{"xmlStrlen", &XmlStrlen},
	{"xmlStrncasecmp", &XmlStrncasecmp},
	{"xmlStrncat", &XmlStrncat},
	{"xmlStrncatNew", &XmlStrncatNew},
	{"xmlStrncmp", &XmlStrncmp},
	{"xmlStrndup", &XmlStrndup},
	{"xmlStrstr", &XmlStrstr},
	{"xmlStrsub", &XmlStrsub},
	{"xmlSubstituteEntitiesDefault", &XmlSubstituteEntitiesDefault},
	// {"xmlSwitchEncoding", &XmlSwitchEncoding},
	// {"xmlSwitchInputEncoding", &XmlSwitchInputEncoding},
	// {"xmlSwitchToEncoding", &XmlSwitchToEncoding},
	{"xmlTextConcat", &XmlTextConcat},
	{"xmlTextMerge", &XmlTextMerge},
	// {"xmlTextReaderAttributeCount", &XmlTextReaderAttributeCount},
	// {"xmlTextReaderBaseUri", &XmlTextReaderBaseUri},
	// {"xmlTextReaderByteConsumed", &XmlTextReaderByteConsumed},
	// {"xmlTextReaderClose", &XmlTextReaderClose},
	// {"xmlTextReaderConstBaseUri", &XmlTextReaderConstBaseUri},
	// {"xmlTextReaderConstEncoding", &XmlTextReaderConstEncoding},
	// {"xmlTextReaderConstLocalName", &XmlTextReaderConstLocalName},
	// {"xmlTextReaderConstName", &XmlTextReaderConstName},
	// {"xmlTextReaderConstNamespaceUri", &XmlTextReaderConstNamespaceUri},
	// {"xmlTextReaderConstPrefix", &XmlTextReaderConstPrefix},
	// {"xmlTextReaderConstString", &XmlTextReaderConstString},
	// {"xmlTextReaderConstValue", &XmlTextReaderConstValue},
	// {"xmlTextReaderConstXmlLang", &XmlTextReaderConstXmlLang},
	// {"xmlTextReaderConstXmlVersion", &XmlTextReaderConstXmlVersion},
	// {"xmlTextReaderCurrentDoc", &XmlTextReaderCurrentDoc},
	// {"xmlTextReaderCurrentNode", &XmlTextReaderCurrentNode},
	// {"xmlTextReaderDepth", &XmlTextReaderDepth},
	// {"xmlTextReaderExpand", &XmlTextReaderExpand},
	// {"xmlTextReaderGetAttribute", &XmlTextReaderGetAttribute},
	// {"xmlTextReaderGetAttributeNo", &XmlTextReaderGetAttributeNo},
	// {"xmlTextReaderGetAttributeNs", &XmlTextReaderGetAttributeNs},
	// {"xmlTextReaderGetErrorHandler", &XmlTextReaderGetErrorHandler},
	// {"xmlTextReaderGetParserColumnNumber", &XmlTextReaderGetParserColumnNumber},
	// {"xmlTextReaderGetParserLineNumber", &XmlTextReaderGetParserLineNumber},
	// {"xmlTextReaderGetParserProp", &XmlTextReaderGetParserProp},
	// {"xmlTextReaderGetRemainder", &XmlTextReaderGetRemainder},
	// {"xmlTextReaderHasAttributes", &XmlTextReaderHasAttributes},
	// {"xmlTextReaderHasValue", &XmlTextReaderHasValue},
	// {"xmlTextReaderIsDefault", &XmlTextReaderIsDefault},
	// {"xmlTextReaderIsEmptyElement", &XmlTextReaderIsEmptyElement},
	// {"xmlTextReaderIsNamespaceDecl", &XmlTextReaderIsNamespaceDecl},
	// {"xmlTextReaderIsValid", &XmlTextReaderIsValid},
	// {"xmlTextReaderLocalName", &XmlTextReaderLocalName},
	// {"xmlTextReaderLocatorBaseURI", &XmlTextReaderLocatorBaseURI},
	// {"xmlTextReaderLocatorLineNumber", &XmlTextReaderLocatorLineNumber},
	// {"xmlTextReaderLookupNamespace", &XmlTextReaderLookupNamespace},
	// {"xmlTextReaderMoveToAttribute", &XmlTextReaderMoveToAttribute},
	// {"xmlTextReaderMoveToAttributeNo", &XmlTextReaderMoveToAttributeNo},
	// {"xmlTextReaderMoveToAttributeNs", &XmlTextReaderMoveToAttributeNs},
	// {"xmlTextReaderMoveToElement", &XmlTextReaderMoveToElement},
	// {"xmlTextReaderMoveToFirstAttribute", &XmlTextReaderMoveToFirstAttribute},
	// {"xmlTextReaderMoveToNextAttribute", &XmlTextReaderMoveToNextAttribute},
	// {"xmlTextReaderName", &XmlTextReaderName},
	// {"xmlTextReaderNamespaceUri", &XmlTextReaderNamespaceUri},
	// {"xmlTextReaderNext", &XmlTextReaderNext},
	// {"xmlTextReaderNextSibling", &XmlTextReaderNextSibling},
	// {"xmlTextReaderNodeType", &XmlTextReaderNodeType},
	// {"xmlTextReaderNormalization", &XmlTextReaderNormalization},
	// {"xmlTextReaderPrefix", &XmlTextReaderPrefix},
	// {"xmlTextReaderPreserve", &XmlTextReaderPreserve},
	// {"xmlTextReaderPreservePattern", &XmlTextReaderPreservePattern},
	// {"xmlTextReaderQuoteChar", &XmlTextReaderQuoteChar},
	// {"xmlTextReaderRead", &XmlTextReaderRead},
	// {"xmlTextReaderReadAttributeValue", &XmlTextReaderReadAttributeValue},
	// {"xmlTextReaderReadInnerXml", &XmlTextReaderReadInnerXml},
	// {"xmlTextReaderReadOuterXml", &XmlTextReaderReadOuterXml},
	// {"xmlTextReaderReadState", &XmlTextReaderReadState},
	// {"xmlTextReaderReadString", &XmlTextReaderReadString},
	// {"xmlTextReaderRelaxNGSetSchema", &XmlTextReaderRelaxNGSetSchema},
	// {"xmlTextReaderRelaxNGValidate", &XmlTextReaderRelaxNGValidate},
	// {"xmlTextReaderSchemaValidate", &XmlTextReaderSchemaValidate},
	// {"xmlTextReaderSchemaValidateCtxt", &XmlTextReaderSchemaValidateCtxt},
	// {"xmlTextReaderSetErrorHandler", &XmlTextReaderSetErrorHandler},
	// {"xmlTextReaderSetParserProp", &XmlTextReaderSetParserProp},
	// {"xmlTextReaderSetSchema", &XmlTextReaderSetSchema},
	// {"xmlTextReaderSetStructuredErrorHandler", &XmlTextReaderSetStructuredErrorHandler},
	// {"xmlTextReaderSetup", &XmlTextReaderSetup},
	// {"xmlTextReaderStandalone", &XmlTextReaderStandalone},
	// {"xmlTextReaderValue", &XmlTextReaderValue},
	// {"xmlTextReaderXmlLang", &XmlTextReaderXmlLang},
	// {"xmlTextWriterEndAttribute", &XmlTextWriterEndAttribute},
	// {"xmlTextWriterEndCDATA", &XmlTextWriterEndCDATA},
	// {"xmlTextWriterEndComment", &XmlTextWriterEndComment},
	// {"xmlTextWriterEndDTD", &XmlTextWriterEndDTD},
	// {"xmlTextWriterEndDTDAttlist", &XmlTextWriterEndDTDAttlist},
	// {"xmlTextWriterEndDTDElement", &XmlTextWriterEndDTDElement},
	// {"xmlTextWriterEndDTDEntity", &XmlTextWriterEndDTDEntity},
	// {"xmlTextWriterEndDocument", &XmlTextWriterEndDocument},
	// {"xmlTextWriterEndElement", &XmlTextWriterEndElement},
	// {"xmlTextWriterEndPI", &XmlTextWriterEndPI},
	// {"xmlTextWriterFlush", &XmlTextWriterFlush},
	// {"xmlTextWriterFullEndElement", &XmlTextWriterFullEndElement},
	// {"xmlTextWriterSetIndent", &XmlTextWriterSetIndent},
	// {"xmlTextWriterSetIndentString", &XmlTextWriterSetIndentString},
	// {"xmlTextWriterStartAttribute", &XmlTextWriterStartAttribute},
	// {"xmlTextWriterStartAttributeNS", &XmlTextWriterStartAttributeNS},
	// {"xmlTextWriterStartCDATA", &XmlTextWriterStartCDATA},
	// {"xmlTextWriterStartComment", &XmlTextWriterStartComment},
	// {"xmlTextWriterStartDTD", &XmlTextWriterStartDTD},
	// {"xmlTextWriterStartDTDAttlist", &XmlTextWriterStartDTDAttlist},
	// {"xmlTextWriterStartDTDElement", &XmlTextWriterStartDTDElement},
	// {"xmlTextWriterStartDTDEntity", &XmlTextWriterStartDTDEntity},
	// {"xmlTextWriterStartDocument", &XmlTextWriterStartDocument},
	// {"xmlTextWriterStartElement", &XmlTextWriterStartElement},
	// {"xmlTextWriterStartElementNS", &XmlTextWriterStartElementNS},
	// {"xmlTextWriterStartPI", &XmlTextWriterStartPI},
	// {"xmlTextWriterWriteAttribute", &XmlTextWriterWriteAttribute},
	// {"xmlTextWriterWriteAttributeNS", &XmlTextWriterWriteAttributeNS},
	// {"xmlTextWriterWriteBase64", &XmlTextWriterWriteBase64},
	// {"xmlTextWriterWriteBinHex", &XmlTextWriterWriteBinHex},
	// {"xmlTextWriterWriteCDATA", &XmlTextWriterWriteCDATA},
	// {"xmlTextWriterWriteComment", &XmlTextWriterWriteComment},
	// {"xmlTextWriterWriteDTD", &XmlTextWriterWriteDTD},
	// {"xmlTextWriterWriteDTDAttlist", &XmlTextWriterWriteDTDAttlist},
	// {"xmlTextWriterWriteDTDElement", &XmlTextWriterWriteDTDElement},
	// {"xmlTextWriterWriteDTDEntity", &XmlTextWriterWriteDTDEntity},
	// {"xmlTextWriterWriteDTDExternalEntity", &XmlTextWriterWriteDTDExternalEntity},
	// {"xmlTextWriterWriteDTDExternalEntityContents", &XmlTextWriterWriteDTDExternalEntityContents},
	// {"xmlTextWriterWriteDTDInternalEntity", &XmlTextWriterWriteDTDInternalEntity},
	// {"xmlTextWriterWriteDTDNotation", &XmlTextWriterWriteDTDNotation},
	// {"xmlTextWriterWriteElement", &XmlTextWriterWriteElement},
	// {"xmlTextWriterWriteElementNS", &XmlTextWriterWriteElementNS},
	// {"xmlTextWriterWriteFormatAttribute", &XmlTextWriterWriteFormatAttribute},
	// {"xmlTextWriterWriteFormatAttributeNS", &XmlTextWriterWriteFormatAttributeNS},
	// {"xmlTextWriterWriteFormatCDATA", &XmlTextWriterWriteFormatCDATA},
	// {"xmlTextWriterWriteFormatComment", &XmlTextWriterWriteFormatComment},
	// {"xmlTextWriterWriteFormatDTD", &XmlTextWriterWriteFormatDTD},
	// {"xmlTextWriterWriteFormatDTDAttlist", &XmlTextWriterWriteFormatDTDAttlist},
	// {"xmlTextWriterWriteFormatDTDElement", &XmlTextWriterWriteFormatDTDElement},
	// {"xmlTextWriterWriteFormatDTDInternalEntity", &XmlTextWriterWriteFormatDTDInternalEntity},
	// {"xmlTextWriterWriteFormatElement", &XmlTextWriterWriteFormatElement},
	// {"xmlTextWriterWriteFormatElementNS", &XmlTextWriterWriteFormatElementNS},
	// {"xmlTextWriterWriteFormatPI", &XmlTextWriterWriteFormatPI},
	// {"xmlTextWriterWriteFormatRaw", &XmlTextWriterWriteFormatRaw},
	// {"xmlTextWriterWriteFormatString", &XmlTextWriterWriteFormatString},
	// {"xmlTextWriterWritePI", &XmlTextWriterWritePI},
	// {"xmlTextWriterWriteRaw", &XmlTextWriterWriteRaw},
	// {"xmlTextWriterWriteRawLen", &XmlTextWriterWriteRawLen},
	// {"xmlTextWriterWriteString", &XmlTextWriterWriteString},
	// {"xmlTextWriterWriteVFormatAttribute", &XmlTextWriterWriteVFormatAttribute},
	// {"xmlTextWriterWriteVFormatAttributeNS", &XmlTextWriterWriteVFormatAttributeNS},
	// {"xmlTextWriterWriteVFormatCDATA", &XmlTextWriterWriteVFormatCDATA},
	// {"xmlTextWriterWriteVFormatComment", &XmlTextWriterWriteVFormatComment},
	// {"xmlTextWriterWriteVFormatDTD", &XmlTextWriterWriteVFormatDTD},
	// {"xmlTextWriterWriteVFormatDTDAttlist", &XmlTextWriterWriteVFormatDTDAttlist},
	// {"xmlTextWriterWriteVFormatDTDElement", &XmlTextWriterWriteVFormatDTDElement},
	// {"xmlTextWriterWriteVFormatDTDInternalEntity", &XmlTextWriterWriteVFormatDTDInternalEntity},
	// {"xmlTextWriterWriteVFormatElement", &XmlTextWriterWriteVFormatElement},
	// {"xmlTextWriterWriteVFormatElementNS", &XmlTextWriterWriteVFormatElementNS},
	// {"xmlTextWriterWriteVFormatPI", &XmlTextWriterWriteVFormatPI},
	// {"xmlTextWriterWriteVFormatRaw", &XmlTextWriterWriteVFormatRaw},
	// {"xmlTextWriterWriteVFormatString", &XmlTextWriterWriteVFormatString},
	// {"xmlThrDefBufferAllocScheme", &XmlThrDefBufferAllocScheme},
	{"xmlThrDefDefaultBufferSize", &XmlThrDefDefaultBufferSize},
	{"xmlThrDefDeregisterNodeDefault", &XmlThrDefDeregisterNodeDefault},
	{"xmlThrDefDoValidityCheckingDefaultValue", &XmlThrDefDoValidityCheckingDefaultValue},
	{"xmlThrDefGetWarningsDefaultValue", &XmlThrDefGetWarningsDefaultValue},
	{"xmlThrDefIndentTreeOutput", &XmlThrDefIndentTreeOutput},
	{"xmlThrDefKeepBlanksDefaultValue", &XmlThrDefKeepBlanksDefaultValue},
	{"xmlThrDefLineNumbersDefaultValue", &XmlThrDefLineNumbersDefaultValue},
	{"xmlThrDefLoadExtDtdDefaultValue", &XmlThrDefLoadExtDtdDefaultValue},
	{"xmlThrDefOutputBufferCreateFilenameDefault", &XmlThrDefOutputBufferCreateFilenameDefault},
	{"xmlThrDefParserDebugEntities", &XmlThrDefParserDebugEntities},
	{"xmlThrDefParserInputBufferCreateFilenameDefault", &XmlThrDefParserInputBufferCreateFilenameDefault},
	{"xmlThrDefPedanticParserDefaultValue", &XmlThrDefPedanticParserDefaultValue},
	{"xmlThrDefRegisterNodeDefault", &XmlThrDefRegisterNodeDefault},
	{"xmlThrDefSaveNoEmptyTags", &XmlThrDefSaveNoEmptyTags},
	{"xmlThrDefSetGenericErrorFunc", &XmlThrDefSetGenericErrorFunc},
	{"xmlThrDefSetStructuredErrorFunc", &XmlThrDefSetStructuredErrorFunc},
	{"xmlThrDefSubstituteEntitiesDefaultValue", &XmlThrDefSubstituteEntitiesDefaultValue},
	{"xmlThrDefTreeIndentString", &XmlThrDefTreeIndentString},
	// {"xmlUCSIsAegeanNumbers", &XmlUCSIsAegeanNumbers},
	// {"xmlUCSIsAlphabeticPresentationForms", &XmlUCSIsAlphabeticPresentationForms},
	// {"xmlUCSIsArabic", &XmlUCSIsArabic},
	// {"xmlUCSIsArabicPresentationFormsA", &XmlUCSIsArabicPresentationFormsA},
	// {"xmlUCSIsArabicPresentationFormsB", &XmlUCSIsArabicPresentationFormsB},
	// {"xmlUCSIsArmenian", &XmlUCSIsArmenian},
	// {"xmlUCSIsArrows", &XmlUCSIsArrows},
	// {"xmlUCSIsBasicLatin", &XmlUCSIsBasicLatin},
	// {"xmlUCSIsBengali", &XmlUCSIsBengali},
	// {"xmlUCSIsBlock", &XmlUCSIsBlock},
	// {"xmlUCSIsBlockElements", &XmlUCSIsBlockElements},
	// {"xmlUCSIsBopomofo", &XmlUCSIsBopomofo},
	// {"xmlUCSIsBopomofoExtended", &XmlUCSIsBopomofoExtended},
	// {"xmlUCSIsBoxDrawing", &XmlUCSIsBoxDrawing},
	// {"xmlUCSIsBraillePatterns", &XmlUCSIsBraillePatterns},
	// {"xmlUCSIsBuhid", &XmlUCSIsBuhid},
	// {"xmlUCSIsByzantineMusicalSymbols", &XmlUCSIsByzantineMusicalSymbols},
	// {"xmlUCSIsCJKCompatibility", &XmlUCSIsCJKCompatibility},
	// {"xmlUCSIsCJKCompatibilityForms", &XmlUCSIsCJKCompatibilityForms},
	// {"xmlUCSIsCJKCompatibilityIdeographs", &XmlUCSIsCJKCompatibilityIdeographs},
	// {"xmlUCSIsCJKCompatibilityIdeographsSupplement", &XmlUCSIsCJKCompatibilityIdeographsSupplement},
	// {"xmlUCSIsCJKRadicalsSupplement", &XmlUCSIsCJKRadicalsSupplement},
	// {"xmlUCSIsCJKSymbolsandPunctuation", &XmlUCSIsCJKSymbolsandPunctuation},
	// {"xmlUCSIsCJKUnifiedIdeographs", &XmlUCSIsCJKUnifiedIdeographs},
	// {"xmlUCSIsCJKUnifiedIdeographsExtensionA", &XmlUCSIsCJKUnifiedIdeographsExtensionA},
	// {"xmlUCSIsCJKUnifiedIdeographsExtensionB", &XmlUCSIsCJKUnifiedIdeographsExtensionB},
	// {"xmlUCSIsCat", &XmlUCSIsCat},
	// {"xmlUCSIsCatC", &XmlUCSIsCatC},
	// {"xmlUCSIsCatCc", &XmlUCSIsCatCc},
	// {"xmlUCSIsCatCf", &XmlUCSIsCatCf},
	// {"xmlUCSIsCatCo", &XmlUCSIsCatCo},
	// {"xmlUCSIsCatCs", &XmlUCSIsCatCs},
	// {"xmlUCSIsCatL", &XmlUCSIsCatL},
	// {"xmlUCSIsCatLl", &XmlUCSIsCatLl},
	// {"xmlUCSIsCatLm", &XmlUCSIsCatLm},
	// {"xmlUCSIsCatLo", &XmlUCSIsCatLo},
	// {"xmlUCSIsCatLt", &XmlUCSIsCatLt},
	// {"xmlUCSIsCatLu", &XmlUCSIsCatLu},
	// {"xmlUCSIsCatM", &XmlUCSIsCatM},
	// {"xmlUCSIsCatMc", &XmlUCSIsCatMc},
	// {"xmlUCSIsCatMe", &XmlUCSIsCatMe},
	// {"xmlUCSIsCatMn", &XmlUCSIsCatMn},
	// {"xmlUCSIsCatN", &XmlUCSIsCatN},
	// {"xmlUCSIsCatNd", &XmlUCSIsCatNd},
	// {"xmlUCSIsCatNl", &XmlUCSIsCatNl},
	// {"xmlUCSIsCatNo", &XmlUCSIsCatNo},
	// {"xmlUCSIsCatP", &XmlUCSIsCatP},
	// {"xmlUCSIsCatPc", &XmlUCSIsCatPc},
	// {"xmlUCSIsCatPd", &XmlUCSIsCatPd},
	// {"xmlUCSIsCatPe", &XmlUCSIsCatPe},
	// {"xmlUCSIsCatPf", &XmlUCSIsCatPf},
	// {"xmlUCSIsCatPi", &XmlUCSIsCatPi},
	// {"xmlUCSIsCatPo", &XmlUCSIsCatPo},
	// {"xmlUCSIsCatPs", &XmlUCSIsCatPs},
	// {"xmlUCSIsCatS", &XmlUCSIsCatS},
	// {"xmlUCSIsCatSc", &XmlUCSIsCatSc},
	// {"xmlUCSIsCatSk", &XmlUCSIsCatSk},
	// {"xmlUCSIsCatSm", &XmlUCSIsCatSm},
	// {"xmlUCSIsCatSo", &XmlUCSIsCatSo},
	// {"xmlUCSIsCatZ", &XmlUCSIsCatZ},
	// {"xmlUCSIsCatZl", &XmlUCSIsCatZl},
	// {"xmlUCSIsCatZp", &XmlUCSIsCatZp},
	// {"xmlUCSIsCatZs", &XmlUCSIsCatZs},
	// {"xmlUCSIsCherokee", &XmlUCSIsCherokee},
	// {"xmlUCSIsCombiningDiacriticalMarks", &XmlUCSIsCombiningDiacriticalMarks},
	// {"xmlUCSIsCombiningDiacriticalMarksforSymbols", &XmlUCSIsCombiningDiacriticalMarksforSymbols},
	// {"xmlUCSIsCombiningHalfMarks", &XmlUCSIsCombiningHalfMarks},
	// {"xmlUCSIsCombiningMarksforSymbols", &XmlUCSIsCombiningMarksforSymbols},
	// {"xmlUCSIsControlPictures", &XmlUCSIsControlPictures},
	// {"xmlUCSIsCurrencySymbols", &XmlUCSIsCurrencySymbols},
	// {"xmlUCSIsCypriotSyllabary", &XmlUCSIsCypriotSyllabary},
	// {"xmlUCSIsCyrillic", &XmlUCSIsCyrillic},
	// {"xmlUCSIsCyrillicSupplement", &XmlUCSIsCyrillicSupplement},
	// {"xmlUCSIsDeseret", &XmlUCSIsDeseret},
	// {"xmlUCSIsDevanagari", &XmlUCSIsDevanagari},
	// {"xmlUCSIsDingbats", &XmlUCSIsDingbats},
	// {"xmlUCSIsEnclosedAlphanumerics", &XmlUCSIsEnclosedAlphanumerics},
	// {"xmlUCSIsEnclosedCJKLettersandMonths", &XmlUCSIsEnclosedCJKLettersandMonths},
	// {"xmlUCSIsEthiopic", &XmlUCSIsEthiopic},
	// {"xmlUCSIsGeneralPunctuation", &XmlUCSIsGeneralPunctuation},
	// {"xmlUCSIsGeometricShapes", &XmlUCSIsGeometricShapes},
	// {"xmlUCSIsGeorgian", &XmlUCSIsGeorgian},
	// {"xmlUCSIsGothic", &XmlUCSIsGothic},
	// {"xmlUCSIsGreek", &XmlUCSIsGreek},
	// {"xmlUCSIsGreekExtended", &XmlUCSIsGreekExtended},
	// {"xmlUCSIsGreekandCoptic", &XmlUCSIsGreekandCoptic},
	// {"xmlUCSIsGujarati", &XmlUCSIsGujarati},
	// {"xmlUCSIsGurmukhi", &XmlUCSIsGurmukhi},
	// {"xmlUCSIsHalfwidthandFullwidthForms", &XmlUCSIsHalfwidthandFullwidthForms},
	// {"xmlUCSIsHangulCompatibilityJamo", &XmlUCSIsHangulCompatibilityJamo},
	// {"xmlUCSIsHangulJamo", &XmlUCSIsHangulJamo},
	// {"xmlUCSIsHangulSyllables", &XmlUCSIsHangulSyllables},
	// {"xmlUCSIsHanunoo", &XmlUCSIsHanunoo},
	// {"xmlUCSIsHebrew", &XmlUCSIsHebrew},
	// {"xmlUCSIsHighPrivateUseSurrogates", &XmlUCSIsHighPrivateUseSurrogates},
	// {"xmlUCSIsHighSurrogates", &XmlUCSIsHighSurrogates},
	// {"xmlUCSIsHiragana", &XmlUCSIsHiragana},
	// {"xmlUCSIsIPAExtensions", &XmlUCSIsIPAExtensions},
	// {"xmlUCSIsIdeographicDescriptionCharacters", &XmlUCSIsIdeographicDescriptionCharacters},
	// {"xmlUCSIsKanbun", &XmlUCSIsKanbun},
	// {"xmlUCSIsKangxiRadicals", &XmlUCSIsKangxiRadicals},
	// {"xmlUCSIsKannada", &XmlUCSIsKannada},
	// {"xmlUCSIsKatakana", &XmlUCSIsKatakana},
	// {"xmlUCSIsKatakanaPhoneticExtensions", &XmlUCSIsKatakanaPhoneticExtensions},
	// {"xmlUCSIsKhmer", &XmlUCSIsKhmer},
	// {"xmlUCSIsKhmerSymbols", &XmlUCSIsKhmerSymbols},
	// {"xmlUCSIsLao", &XmlUCSIsLao},
	// {"xmlUCSIsLatin1Supplement", &XmlUCSIsLatin1Supplement},
	// {"xmlUCSIsLatinExtendedA", &XmlUCSIsLatinExtendedA},
	// {"xmlUCSIsLatinExtendedAdditional", &XmlUCSIsLatinExtendedAdditional},
	// {"xmlUCSIsLatinExtendedB", &XmlUCSIsLatinExtendedB},
	// {"xmlUCSIsLetterlikeSymbols", &XmlUCSIsLetterlikeSymbols},
	// {"xmlUCSIsLimbu", &XmlUCSIsLimbu},
	// {"xmlUCSIsLinearBIdeograms", &XmlUCSIsLinearBIdeograms},
	// {"xmlUCSIsLinearBSyllabary", &XmlUCSIsLinearBSyllabary},
	// {"xmlUCSIsLowSurrogates", &XmlUCSIsLowSurrogates},
	// {"xmlUCSIsMalayalam", &XmlUCSIsMalayalam},
	// {"xmlUCSIsMathematicalAlphanumericSymbols", &XmlUCSIsMathematicalAlphanumericSymbols},
	// {"xmlUCSIsMathematicalOperators", &XmlUCSIsMathematicalOperators},
	// {"xmlUCSIsMiscellaneousMathematicalSymbolsA", &XmlUCSIsMiscellaneousMathematicalSymbolsA},
	// {"xmlUCSIsMiscellaneousMathematicalSymbolsB", &XmlUCSIsMiscellaneousMathematicalSymbolsB},
	// {"xmlUCSIsMiscellaneousSymbols", &XmlUCSIsMiscellaneousSymbols},
	// {"xmlUCSIsMiscellaneousSymbolsandArrows", &XmlUCSIsMiscellaneousSymbolsandArrows},
	// {"xmlUCSIsMiscellaneousTechnical", &XmlUCSIsMiscellaneousTechnical},
	// {"xmlUCSIsMongolian", &XmlUCSIsMongolian},
	// {"xmlUCSIsMusicalSymbols", &XmlUCSIsMusicalSymbols},
	// {"xmlUCSIsMyanmar", &XmlUCSIsMyanmar},
	// {"xmlUCSIsNumberForms", &XmlUCSIsNumberForms},
	// {"xmlUCSIsOgham", &XmlUCSIsOgham},
	// {"xmlUCSIsOldItalic", &XmlUCSIsOldItalic},
	// {"xmlUCSIsOpticalCharacterRecognition", &XmlUCSIsOpticalCharacterRecognition},
	// {"xmlUCSIsOriya", &XmlUCSIsOriya},
	// {"xmlUCSIsOsmanya", &XmlUCSIsOsmanya},
	// {"xmlUCSIsPhoneticExtensions", &XmlUCSIsPhoneticExtensions},
	// {"xmlUCSIsPrivateUse", &XmlUCSIsPrivateUse},
	// {"xmlUCSIsPrivateUseArea", &XmlUCSIsPrivateUseArea},
	// {"xmlUCSIsRunic", &XmlUCSIsRunic},
	// {"xmlUCSIsShavian", &XmlUCSIsShavian},
	// {"xmlUCSIsSinhala", &XmlUCSIsSinhala},
	// {"xmlUCSIsSmallFormVariants", &XmlUCSIsSmallFormVariants},
	// {"xmlUCSIsSpacingModifierLetters", &XmlUCSIsSpacingModifierLetters},
	// {"xmlUCSIsSpecials", &XmlUCSIsSpecials},
	// {"xmlUCSIsSuperscriptsandSubscripts", &XmlUCSIsSuperscriptsandSubscripts},
	// {"xmlUCSIsSupplementalArrowsA", &XmlUCSIsSupplementalArrowsA},
	// {"xmlUCSIsSupplementalArrowsB", &XmlUCSIsSupplementalArrowsB},
	// {"xmlUCSIsSupplementalMathematicalOperators", &XmlUCSIsSupplementalMathematicalOperators},
	// {"xmlUCSIsSupplementaryPrivateUseAreaA", &XmlUCSIsSupplementaryPrivateUseAreaA},
	// {"xmlUCSIsSupplementaryPrivateUseAreaB", &XmlUCSIsSupplementaryPrivateUseAreaB},
	// {"xmlUCSIsSyriac", &XmlUCSIsSyriac},
	// {"xmlUCSIsTagalog", &XmlUCSIsTagalog},
	// {"xmlUCSIsTagbanwa", &XmlUCSIsTagbanwa},
	// {"xmlUCSIsTags", &XmlUCSIsTags},
	// {"xmlUCSIsTaiLe", &XmlUCSIsTaiLe},
	// {"xmlUCSIsTaiXuanJingSymbols", &XmlUCSIsTaiXuanJingSymbols},
	// {"xmlUCSIsTamil", &XmlUCSIsTamil},
	// {"xmlUCSIsTelugu", &XmlUCSIsTelugu},
	// {"xmlUCSIsThaana", &XmlUCSIsThaana},
	// {"xmlUCSIsThai", &XmlUCSIsThai},
	// {"xmlUCSIsTibetan", &XmlUCSIsTibetan},
	// {"xmlUCSIsUgaritic", &XmlUCSIsUgaritic},
	// {"xmlUCSIsUnifiedCanadianAboriginalSyllabics", &XmlUCSIsUnifiedCanadianAboriginalSyllabics},
	// {"xmlUCSIsVariationSelectors", &XmlUCSIsVariationSelectors},
	// {"xmlUCSIsVariationSelectorsSupplement", &XmlUCSIsVariationSelectorsSupplement},
	// {"xmlUCSIsYiRadicals", &XmlUCSIsYiRadicals},
	// {"xmlUCSIsYiSyllables", &XmlUCSIsYiSyllables},
	// {"xmlUCSIsYijingHexagramSymbols", &XmlUCSIsYijingHexagramSymbols},
	// {"xmlURIEscape", &XmlURIEscape},
	// {"xmlURIEscapeStr", &XmlURIEscapeStr},
	// {"xmlURIUnescapeString", &XmlURIUnescapeString},
	{"xmlUTF8Charcmp", &XmlUTF8Charcmp},
	{"xmlUTF8Size", &XmlUTF8Size},
	{"xmlUTF8Strlen", &XmlUTF8Strlen},
	{"xmlUTF8Strloc", &XmlUTF8Strloc},
	{"xmlUTF8Strndup", &XmlUTF8Strndup},
	{"xmlUTF8Strpos", &XmlUTF8Strpos},
	{"xmlUTF8Strsize", &XmlUTF8Strsize},
	{"xmlUTF8Strsub", &XmlUTF8Strsub},
	{"xmlUnlinkNode", &XmlUnlinkNode},
	{"xmlUnlockLibrary", &XmlUnlockLibrary},
	{"xmlUnsetNsProp", &XmlUnsetNsProp},
	{"xmlUnsetProp", &XmlUnsetProp},
	{"xmlValidBuildContentModel", &XmlValidBuildContentModel},
	{"xmlValidCtxtNormalizeAttributeValue", &XmlValidCtxtNormalizeAttributeValue},
	{"xmlValidGetPotentialChildren", &XmlValidGetPotentialChildren},
	{"xmlValidGetValidElements", &XmlValidGetValidElements},
	{"xmlValidNormalizeAttributeValue", &XmlValidNormalizeAttributeValue},
	{"xmlValidateAttributeDecl", &XmlValidateAttributeDecl},
	{"xmlValidateAttributeValue", &XmlValidateAttributeValue},
	{"xmlValidateDocument", &XmlValidateDocument},
	{"xmlValidateDocumentFinal", &XmlValidateDocumentFinal},
	{"xmlValidateDtd", &XmlValidateDtd},
	{"xmlValidateDtdFinal", &XmlValidateDtdFinal},
	{"xmlValidateElement", &XmlValidateElement},
	{"xmlValidateElementDecl", &XmlValidateElementDecl},
	{"xmlValidateNCName", &XmlValidateNCName},
	{"xmlValidateNMToken", &XmlValidateNMToken},
	{"xmlValidateName", &XmlValidateName},
	{"xmlValidateNameValue", &XmlValidateNameValue},
	{"xmlValidateNamesValue", &XmlValidateNamesValue},
	{"xmlValidateNmtokenValue", &XmlValidateNmtokenValue},
	{"xmlValidateNmtokensValue", &XmlValidateNmtokensValue},
	{"xmlValidateNotationDecl", &XmlValidateNotationDecl},
	{"xmlValidateNotationUse", &XmlValidateNotationUse},
	{"xmlValidateOneAttribute", &XmlValidateOneAttribute},
	{"xmlValidateOneElement", &XmlValidateOneElement},
	{"xmlValidateOneNamespace", &XmlValidateOneNamespace},
	{"xmlValidatePopElement", &XmlValidatePopElement},
	{"xmlValidatePushCData", &XmlValidatePushCData},
	{"xmlValidatePushElement", &XmlValidatePushElement},
	{"xmlValidateQName", &XmlValidateQName},
	{"xmlValidateRoot", &XmlValidateRoot},
	// {"xmlXIncludeFreeContext", &XmlXIncludeFreeContext},
	// {"xmlXIncludeNewContext", &XmlXIncludeNewContext},
	// {"xmlXIncludeProcess", &XmlXIncludeProcess},
	// {"xmlXIncludeProcessFlags", &XmlXIncludeProcessFlags},
	// {"xmlXIncludeProcessFlagsData", &XmlXIncludeProcessFlagsData},
	// {"xmlXIncludeProcessNode", &XmlXIncludeProcessNode},
	// {"xmlXIncludeProcessTree", &XmlXIncludeProcessTree},
	// {"xmlXIncludeProcessTreeFlags", &XmlXIncludeProcessTreeFlags},
	// {"xmlXIncludeProcessTreeFlagsData", &XmlXIncludeProcessTreeFlagsData},
	// {"xmlXIncludeSetFlags", &XmlXIncludeSetFlags},
	// {"xmlXPathAddValues", &XmlXPathAddValues},
	// {"xmlXPathBooleanFunction", &XmlXPathBooleanFunction},
	// {"xmlXPathCastBooleanToNumber", &XmlXPathCastBooleanToNumber},
	// {"xmlXPathCastBooleanToString", &XmlXPathCastBooleanToString},
	// {"xmlXPathCastNodeSetToBoolean", &XmlXPathCastNodeSetToBoolean},
	// {"xmlXPathCastNodeSetToNumber", &XmlXPathCastNodeSetToNumber},
	// {"xmlXPathCastNodeSetToString", &XmlXPathCastNodeSetToString},
	// {"xmlXPathCastNodeToNumber", &XmlXPathCastNodeToNumber},
	// {"xmlXPathCastNodeToString", &XmlXPathCastNodeToString},
	// {"xmlXPathCastNumberToBoolean", &XmlXPathCastNumberToBoolean},
	// {"xmlXPathCastNumberToString", &XmlXPathCastNumberToString},
	// {"xmlXPathCastStringToBoolean", &XmlXPathCastStringToBoolean},
	// {"xmlXPathCastStringToNumber", &XmlXPathCastStringToNumber},
	// {"xmlXPathCastToBoolean", &XmlXPathCastToBoolean},
	// {"xmlXPathCastToNumber", &XmlXPathCastToNumber},
	// {"xmlXPathCastToString", &XmlXPathCastToString},
	// {"xmlXPathCeilingFunction", &XmlXPathCeilingFunction},
	// {"xmlXPathCmpNodes", &XmlXPathCmpNodes},
	// {"xmlXPathCompareValues", &XmlXPathCompareValues},
	// {"xmlXPathCompile", &XmlXPathCompile},
	// {"xmlXPathCompiledEval", &XmlXPathCompiledEval},
	// {"xmlXPathCompiledEvalToBoolean", &XmlXPathCompiledEvalToBoolean},
	// {"xmlXPathConcatFunction", &XmlXPathConcatFunction},
	// {"xmlXPathContainsFunction", &XmlXPathContainsFunction},
	// {"xmlXPathContextSetCache", &XmlXPathContextSetCache},
	// {"xmlXPathConvertBoolean", &XmlXPathConvertBoolean},
	// {"xmlXPathConvertNumber", &XmlXPathConvertNumber},
	// {"xmlXPathConvertString", &XmlXPathConvertString},
	// {"xmlXPathCountFunction", &XmlXPathCountFunction},
	// {"xmlXPathCtxtCompile", &XmlXPathCtxtCompile},
	// {"xmlXPathDebugDumpCompExpr", &XmlXPathDebugDumpCompExpr},
	// {"xmlXPathDebugDumpObject", &XmlXPathDebugDumpObject},
	// {"xmlXPathDifference", &XmlXPathDifference},
	// {"xmlXPathDistinct", &XmlXPathDistinct},
	// {"xmlXPathDistinctSorted", &XmlXPathDistinctSorted},
	// {"xmlXPathDivValues", &XmlXPathDivValues},
	// {"xmlXPathEqualValues", &XmlXPathEqualValues},
	// {"xmlXPathErr", &XmlXPathErr},
	// {"xmlXPathEval", &XmlXPathEval},
	// {"xmlXPathEvalExpr", &XmlXPathEvalExpr},
	// {"xmlXPathEvalExpression", &XmlXPathEvalExpression},
	// {"xmlXPathEvalPredicate", &XmlXPathEvalPredicate},
	// {"xmlXPathEvaluatePredicateResult", &XmlXPathEvaluatePredicateResult},
	// {"xmlXPathFalseFunction", &XmlXPathFalseFunction},
	// {"xmlXPathFloorFunction", &XmlXPathFloorFunction},
	// {"xmlXPathFreeCompExpr", &XmlXPathFreeCompExpr},
	// {"xmlXPathFreeContext", &XmlXPathFreeContext},
	// {"xmlXPathFreeNodeSet", &XmlXPathFreeNodeSet},
	// {"xmlXPathFreeNodeSetList", &XmlXPathFreeNodeSetList},
	// {"xmlXPathFreeObject", &XmlXPathFreeObject},
	// {"xmlXPathFreeParserContext", &XmlXPathFreeParserContext},
	// {"xmlXPathFunctionLookup", &XmlXPathFunctionLookup},
	// {"xmlXPathFunctionLookupNS", &XmlXPathFunctionLookupNS},
	// {"xmlXPathHasSameNodes", &XmlXPathHasSameNodes},
	// {"xmlXPathIdFunction", &XmlXPathIdFunction},
	// {"xmlXPathInit", &XmlXPathInit},
	// {"xmlXPathIntersection", &XmlXPathIntersection},
	// {"xmlXPathIsInf", &XmlXPathIsInf},
	// {"xmlXPathIsNaN", &XmlXPathIsNaN},
	// {"xmlXPathIsNodeType", &XmlXPathIsNodeType},
	// {"xmlXPathLangFunction", &XmlXPathLangFunction},
	// {"xmlXPathLastFunction", &XmlXPathLastFunction},
	// {"xmlXPathLeading", &XmlXPathLeading},
	// {"xmlXPathLeadingSorted", &XmlXPathLeadingSorted},
	// {"xmlXPathLocalNameFunction", &XmlXPathLocalNameFunction},
	// {"xmlXPathModValues", &XmlXPathModValues},
	// {"xmlXPathMultValues", &XmlXPathMultValues},
	// {"xmlXPathNAN", &XmlXPathNAN},
	// {"xmlXPathNINF", &XmlXPathNINF},
	// {"xmlXPathNamespaceURIFunction", &XmlXPathNamespaceURIFunction},
	// {"xmlXPathNewBoolean", &XmlXPathNewBoolean},
	// {"xmlXPathNewCString", &XmlXPathNewCString},
	// {"xmlXPathNewContext", &XmlXPathNewContext},
	// {"xmlXPathNewFloat", &XmlXPathNewFloat},
	// {"xmlXPathNewNodeSet", &XmlXPathNewNodeSet},
	// {"xmlXPathNewNodeSetList", &XmlXPathNewNodeSetList},
	// {"xmlXPathNewParserContext", &XmlXPathNewParserContext},
	// {"xmlXPathNewString", &XmlXPathNewString},
	// {"xmlXPathNewValueTree", &XmlXPathNewValueTree},
	// {"xmlXPathNextAncestor", &XmlXPathNextAncestor},
	// {"xmlXPathNextAncestorOrSelf", &XmlXPathNextAncestorOrSelf},
	// {"xmlXPathNextAttribute", &XmlXPathNextAttribute},
	// {"xmlXPathNextChild", &XmlXPathNextChild},
	// {"xmlXPathNextDescendant", &XmlXPathNextDescendant},
	// {"xmlXPathNextDescendantOrSelf", &XmlXPathNextDescendantOrSelf},
	// {"xmlXPathNextFollowing", &XmlXPathNextFollowing},
	// {"xmlXPathNextFollowingSibling", &XmlXPathNextFollowingSibling},
	// {"xmlXPathNextNamespace", &XmlXPathNextNamespace},
	// {"xmlXPathNextParent", &XmlXPathNextParent},
	// {"xmlXPathNextPreceding", &XmlXPathNextPreceding},
	// {"xmlXPathNextPrecedingSibling", &XmlXPathNextPrecedingSibling},
	// {"xmlXPathNextSelf", &XmlXPathNextSelf},
	// {"xmlXPathNodeLeading", &XmlXPathNodeLeading},
	// {"xmlXPathNodeLeadingSorted", &XmlXPathNodeLeadingSorted},
	// {"xmlXPathNodeSetAdd", &XmlXPathNodeSetAdd},
	// {"xmlXPathNodeSetAddNs", &XmlXPathNodeSetAddNs},
	// {"xmlXPathNodeSetAddUnique", &XmlXPathNodeSetAddUnique},
	// {"xmlXPathNodeSetContains", &XmlXPathNodeSetContains},
	// {"xmlXPathNodeSetCreate", &XmlXPathNodeSetCreate},
	// {"xmlXPathNodeSetDel", &XmlXPathNodeSetDel},
	// {"xmlXPathNodeSetFreeNs", &XmlXPathNodeSetFreeNs},
	// {"xmlXPathNodeSetMerge", &XmlXPathNodeSetMerge},
	// {"xmlXPathNodeSetRemove", &XmlXPathNodeSetRemove},
	// {"xmlXPathNodeSetSort", &XmlXPathNodeSetSort},
	// {"xmlXPathNodeTrailing", &XmlXPathNodeTrailing},
	// {"xmlXPathNodeTrailingSorted", &XmlXPathNodeTrailingSorted},
	// {"xmlXPathNormalizeFunction", &XmlXPathNormalizeFunction},
	// {"xmlXPathNotEqualValues", &XmlXPathNotEqualValues},
	// {"xmlXPathNotFunction", &XmlXPathNotFunction},
	// {"xmlXPathNsLookup", &XmlXPathNsLookup},
	// {"xmlXPathNumberFunction", &XmlXPathNumberFunction},
	// {"xmlXPathObjectCopy", &XmlXPathObjectCopy},
	// {"xmlXPathOrderDocElems", &XmlXPathOrderDocElems},
	// {"xmlXPathPINF", &XmlXPathPINF},
	// {"xmlXPathParseNCName", &XmlXPathParseNCName},
	// {"xmlXPathParseName", &XmlXPathParseName},
	// {"xmlXPathPopBoolean", &XmlXPathPopBoolean},
	// {"xmlXPathPopExternal", &XmlXPathPopExternal},
	// {"xmlXPathPopNodeSet", &XmlXPathPopNodeSet},
	// {"xmlXPathPopNumber", &XmlXPathPopNumber},
	// {"xmlXPathPopString", &XmlXPathPopString},
	// {"xmlXPathPositionFunction", &XmlXPathPositionFunction},
	// {"xmlXPathRegisterAllFunctions", &XmlXPathRegisterAllFunctions},
	// {"xmlXPathRegisterFunc", &XmlXPathRegisterFunc},
	// {"xmlXPathRegisterFuncLookup", &XmlXPathRegisterFuncLookup},
	// {"xmlXPathRegisterFuncNS", &XmlXPathRegisterFuncNS},
	// {"xmlXPathRegisterNs", &XmlXPathRegisterNs},
	// {"xmlXPathRegisterVariable", &XmlXPathRegisterVariable},
	// {"xmlXPathRegisterVariableLookup", &XmlXPathRegisterVariableLookup},
	// {"xmlXPathRegisterVariableNS", &XmlXPathRegisterVariableNS},
	// {"xmlXPathRegisteredFuncsCleanup", &XmlXPathRegisteredFuncsCleanup},
	// {"xmlXPathRegisteredNsCleanup", &XmlXPathRegisteredNsCleanup},
	// {"xmlXPathRegisteredVariablesCleanup", &XmlXPathRegisteredVariablesCleanup},
	// {"xmlXPathRoot", &XmlXPathRoot},
	// {"xmlXPathRoundFunction", &XmlXPathRoundFunction},
	// {"xmlXPathStartsWithFunction", &XmlXPathStartsWithFunction},
	// {"xmlXPathStringEvalNumber", &XmlXPathStringEvalNumber},
	// {"xmlXPathStringFunction", &XmlXPathStringFunction},
	// {"xmlXPathStringLengthFunction", &XmlXPathStringLengthFunction},
	// {"xmlXPathSubValues", &XmlXPathSubValues},
	// {"xmlXPathSubstringAfterFunction", &XmlXPathSubstringAfterFunction},
	// {"xmlXPathSubstringBeforeFunction", &XmlXPathSubstringBeforeFunction},
	// {"xmlXPathSubstringFunction", &XmlXPathSubstringFunction},
	// {"xmlXPathSumFunction", &XmlXPathSumFunction},
	// {"xmlXPathTrailing", &XmlXPathTrailing},
	// {"xmlXPathTrailingSorted", &XmlXPathTrailingSorted},
	// {"xmlXPathTranslateFunction", &XmlXPathTranslateFunction},
	// {"xmlXPathTrueFunction", &XmlXPathTrueFunction},
	// {"xmlXPathValueFlipSign", &XmlXPathValueFlipSign},
	// {"xmlXPathVariableLookup", &XmlXPathVariableLookup},
	// {"xmlXPathVariableLookupNS", &XmlXPathVariableLookupNS},
	// {"xmlXPathWrapCString", &XmlXPathWrapCString},
	// {"xmlXPathWrapExternal", &XmlXPathWrapExternal},
	// {"xmlXPathWrapNodeSet", &XmlXPathWrapNodeSet},
	// {"xmlXPathWrapString", &XmlXPathWrapString},
	// {"xmlXPatherror", &XmlXPatherror},
	// {"xmlXPtrBuildNodeList", &XmlXPtrBuildNodeList},
	// {"xmlXPtrEval", &XmlXPtrEval},
	// {"xmlXPtrEvalRangePredicate", &XmlXPtrEvalRangePredicate},
	// {"xmlXPtrFreeLocationSet", &XmlXPtrFreeLocationSet},
	// {"xmlXPtrLocationSetAdd", &XmlXPtrLocationSetAdd},
	// {"xmlXPtrLocationSetCreate", &XmlXPtrLocationSetCreate},
	// {"xmlXPtrLocationSetDel", &XmlXPtrLocationSetDel},
	// {"xmlXPtrLocationSetMerge", &XmlXPtrLocationSetMerge},
	// {"xmlXPtrLocationSetRemove", &XmlXPtrLocationSetRemove},
	// {"xmlXPtrNewCollapsedRange", &XmlXPtrNewCollapsedRange},
	// {"xmlXPtrNewContext", &XmlXPtrNewContext},
	// {"xmlXPtrNewLocationSetNodeSet", &XmlXPtrNewLocationSetNodeSet},
	// {"xmlXPtrNewLocationSetNodes", &XmlXPtrNewLocationSetNodes},
	// {"xmlXPtrNewRange", &XmlXPtrNewRange},
	// {"xmlXPtrNewRangeNodeObject", &XmlXPtrNewRangeNodeObject},
	// {"xmlXPtrNewRangeNodePoint", &XmlXPtrNewRangeNodePoint},
	// {"xmlXPtrNewRangeNodes", &XmlXPtrNewRangeNodes},
	// {"xmlXPtrNewRangePointNode", &XmlXPtrNewRangePointNode},
	// {"xmlXPtrNewRangePoints", &XmlXPtrNewRangePoints},
	// {"xmlXPtrRangeToFunction", &XmlXPtrRangeToFunction},
	// {"xmlXPtrWrapLocationSet", &XmlXPtrWrapLocationSet},
}
