// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package png provides API definitions for accessing
//libpng14-14.dll.
package png

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-gtk2/types"
	Z "github.com/tHinqa/outside-gtk2/zlib"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

//TODO(t):check struct packing/sizes

type (
	JmpBuf     [16]int
	AllocSizeT SizeT
	Doublep    *float64
	FixedPoint Int32
	Int32      int // Anomally: Size?
	SizeT      T.SizeT
	Struct     StructDef
	Textp      *Text
	Uint32     uint // Anomally: Size?

	LongjmpPtr func(JmpBuf, int)

	ErrorPtr func(*Struct, *T.Char)

	UserTransformPtr func(
		*Struct, *RowInfo, *byte)

	FlushPtr func(*Struct)

	RwPtr func(*Struct, *byte, SizeT)

	ReadStatusPtr func(*Struct, Uint32, int)

	WriteStatusPtr func(*Struct, Uint32, int)

	ProgressiveInfoPtr func(*Struct, *Info)

	ProgressiveEndPtr func(*Struct, *Info)

	ProgressiveRowPtr func(
		*Struct, *byte, Uint32, int)

	UserChunkPtr func(*Struct, *UnknownChunk) int

	MallocPtr func(*Struct, AllocSizeT) *T.Void

	FreePtr func(*Struct, *T.Void)

	//UnknownChunkPtr func(*Struct)

	Time struct {
		Year   uint16
		Month  byte
		Day    byte
		Hour   byte
		Minute byte
		Second byte
	}

	Color8 struct {
		Red   byte
		Green byte
		Blue  byte
		Gray  byte
		Alpha byte
	}

	RowInfo struct {
		Width      Uint32
		Rowbytes   SizeT
		ColorType  byte
		BitDepth   byte
		Channels   byte
		PixelDepth byte
	}

	Color struct {
		Red   byte
		Green byte
		Blue  byte
	}

	Color16 struct {
		Index byte
		Red   uint16
		Green uint16
		Blue  uint16
		Gray  uint16
	}

	Text struct {
		Compression int
		Key         *T.Char
		Text        *T.Char
		TextLength  SizeT
		ItxtLength  SizeT
		Lang        *T.Char
		LangKey     *T.Char
	}

	UnknownChunk struct {
		Name     [5]byte
		Data     *byte
		Size     SizeT
		Location byte
	}

	SPLTT struct {
		Name     *T.Char
		Depth    byte
		Entries  *SPLTEntry
		Nentries Int32
	}

	SPLTEntry struct {
		Red       uint16
		Green     uint16
		Blue      uint16
		Alpha     uint16
		Frequency uint16
	}

	StructDef struct { // Deprecated
		Jmpbuf                JmpBuf
		LongjmpFn             LongjmpPtr
		ErrorFn               ErrorPtr
		WarningFn             ErrorPtr
		ErrorPtr              *T.Void
		WriteDataFn           RwPtr
		ReadDataFn            RwPtr
		IoPtr                 *T.Void
		ReadUserTransformFn   UserTransformPtr
		WriteUserTransformFn  UserTransformPtr
		UserTransformPtr      *T.Void
		UserTransformDepth    byte
		UserTransformChannels byte
		Mode                  Uint32
		Flags                 Uint32
		Transformations       Uint32
		Zstream               Z.ZStream
		Zbuf                  *byte
		ZbufSize              SizeT
		ZlibLevel             int
		ZlibMethod            int
		ZlibWindowBits        int
		ZlibMemLevel          int
		ZlibStrategy          int
		Width                 Uint32
		Height                Uint32
		NumRows               Uint32
		UsrWidth              Uint32
		Rowbytes              SizeT
		UserChunkMallocMax    AllocSizeT
		Iwidth                Uint32
		RowNumber             Uint32
		PrevRow               *byte
		RowBuf                *byte
		SubRow                *byte
		UpRow                 *byte
		AvgRow                *byte
		PaethRow              *byte
		RowInfo               RowInfo
		IdatSize              Uint32
		Crc                   Uint32
		Palette               *Color
		NumPalette            uint16
		NumTrans              uint16
		ChunkName             [5]byte
		Compression           byte
		Filter                byte
		Interlaced            byte
		Pass                  byte
		DoFilter              byte
		ColorType             byte
		BitDepth              byte
		UsrBitDepth           byte
		PixelDepth            byte
		Channels              byte
		UsrChannels           byte
		Sigbytes              byte
		Filler                uint16
		BackgroundGammaType   byte
		BackgroundGamma       float32
		Background            Color16
		Background1           Color16
		OutputFlushFn         FlushPtr
		FlushDist             Uint32
		FlushRows             Uint32
		GammaShift            int
		Gamma                 float32
		ScreenGamma           float32
		GammaTable            *byte
		GammaFrom1            *byte
		GammaTo1              *byte
		Gamma16Table          **uint16
		Gamma16From1          **uint16
		Gamma16To1            **uint16
		SigBit                Color8
		Shift                 Color8
		TransAlpha            *byte
		TransColor            Color16
		ReadRowFn             ReadStatusPtr
		WriteRowFn            WriteStatusPtr
		InfoFn                ProgressiveInfoPtr
		RowFn                 ProgressiveRowPtr
		EndFn                 ProgressiveEndPtr
		SaveBufferPtr         *byte
		SaveBuffer            *byte
		CurrentBufferPtr      *byte
		CurrentBuffer         *byte
		PushLength            Uint32
		SkipLength            Uint32
		SaveBufferSize        SizeT
		SaveBufferMax         SizeT
		BufferSize            SizeT
		CurrentBufferSize     SizeT
		ProcessMode           int
		CurPalette            int
		CurrentTextSize       SizeT
		CurrentTextLeft       SizeT
		CurrentText           *T.Char
		CurrentTextPtr        *T.Char
		PaletteLookup         *byte
		QuantizeIndex         *byte
		Hist                  *uint16
		HeuristicMethod       byte
		NumPrevFilters        byte
		PrevFilters           *byte
		FilterWeights         *uint16
		InvFilterWeights      *uint16
		FilterCosts           *uint16
		InvFilterCosts        *uint16
		TimeBuffer            *T.Char
		FreeMe                Uint32
		UserChunkPtr          *T.Void
		ReadUserChunkFn       UserChunkPtr
		NumChunkList          int
		ChunkList             *byte
		RgbToGrayStatus       byte
		RgbToGrayRedCoeff     uint16
		RgbToGrayGreenCoeff   uint16
		RgbToGrayBlueCoeff    uint16
		MngFeaturesPermitted  Uint32
		IntGamma              FixedPoint
		FilterType            byte
		MemPtr                *T.Void
		MallocFn              MallocPtr
		FreeFn                FreePtr
		BigRowBuf             *byte
		QuantizeSort          *byte
		IndexToPalette        *byte
		PaletteToIndex        *byte
		CompressionType       byte
		UserWidthMax          Uint32
		UserHeightMax         Uint32
		UserChunkCacheMax     Uint32
		UnknownChunk          UnknownChunk
		OldBigRowBufSize      Uint32
		OldPrevRowSize        Uint32
		Chunkdata             *T.Char
		IoState               Uint32
	}

	Info struct { // Deprecated
		Width            Uint32
		Height           Uint32
		Valid            Uint32
		Rowbytes         SizeT
		Palette          *Color
		NumPalette       uint16
		NumTrans         uint16
		BitDepth         byte
		ColorType        byte
		CompressionType  byte
		FilterType       byte
		InterlaceType    byte
		Channels         byte
		PixelDepth       byte
		Sparebyte        byte
		Signature        [8]byte
		Gamma            float32
		SrgbIntent       byte
		NumText          int
		MaxText          int
		Text             Textp
		ModTime          Time
		SigBit           Color8
		TransAlpha       *byte
		TransColor       Color16
		Background       Color16
		XOffset          Int32
		YOffset          Int32
		OffsetUnitType   byte
		XPixelsPerUnit   Uint32
		YPixelsPerUnit   Uint32
		PhysUnitType     byte
		Hist             *uint16
		XWhite           float32
		YWhite           float32
		XRed             float32
		YRed             float32
		XGreen           float32
		YGreen           float32
		XBlue            float32
		YBlue            float32
		PcalPurpose      *T.Char
		PcalX0           Int32
		PcalX1           Int32
		PcalUnits        *T.Char
		PcalParams       **T.Char
		PcalType         byte
		PcalNparams      byte
		FreeMe           Uint32
		UnknownChunks    *UnknownChunk
		UnknownChunksNum SizeT
		IccpName         *T.Char
		IccpProfile      *T.Char
		IccpProflen      Uint32
		IccpCompression  byte
		SpltPalettes     *SPLTT
		SpltPalettesNum  Uint32
		ScalUnit         byte
		ScalPixelWidth   float64
		ScalPixelHeight  float64
		ScalSWidth       *T.Char
		ScalSHeight      *T.Char
		RowPointers      **byte
		IntGamma         FixedPoint
		IntXWhite        FixedPoint
		IntYWhite        FixedPoint
		IntXRed          FixedPoint
		IntYRed          FixedPoint
		IntXGreen        FixedPoint
		IntYGreen        FixedPoint
		IntXBlue         FixedPoint
		IntYBlue         FixedPoint
	}
)

var (
	AccessVersionNumber func() Uint32

	SetSigbytes func(
		pngPtr *Struct,
		numbytes int)

	SigCmp func(
		sig *byte,
		start SizeT,
		numToCheck SizeT) int

	CreateReadStruct func(
		userVer *T.Char,
		errorPtr *T.Void,
		errorFn ErrorPtr,
		warnFn ErrorPtr) *Struct

	CreateWriteStruct func(
		userVer *T.Char,
		errorPtr *T.Void,
		errorFn ErrorPtr,
		warnFn ErrorPtr) *Struct

	GetCompressionBufferSize func(
		pngPtr *Struct) SizeT

	SetCompressionBufferSize func(
		pngPtr *Struct,
		size SizeT)

	SetLongjmpFn func(
		pngPtr *Struct,
		longjmpFn LongjmpPtr,
		JmpBufSize T.SizeT) *JmpBuf

	ResetZstream func(
		pngPtr *Struct) int

	CreateReadStruct2 func(
		userVer *T.Char,
		errorPtr *T.Void,
		errorFn ErrorPtr,
		warnFn ErrorPtr,
		memPtr *T.Void,
		mallocFn MallocPtr,
		freeFn FreePtr) *Struct

	CreateWriteStruct2 func(
		userVer *T.Char,
		errorPtr *T.Void,
		errorFn ErrorPtr,
		warnFn ErrorPtr,
		memPtr *T.Void,
		mallocFn MallocPtr,
		freeFn FreePtr) *Struct

	WriteSig func(
		pngPtr *Struct)

	WriteChunk func(
		pngPtr *Struct,
		chunkName *byte,
		data *byte,
		length SizeT)

	WriteChunkStart func(
		pngPtr *Struct,
		chunkName *byte,
		length Uint32)

	WriteChunkData func(
		pngPtr *Struct,
		data *byte,
		length SizeT)

	WriteChunkEnd func(
		pngPtr *Struct)

	CreateInfoStruct func(
		pngPtr *Struct) *Info

	InfoInit3 func(
		infoPtr **Info,
		pngInfoStructSize SizeT)

	WriteInfoBeforePLTE func(
		pngPtr *Struct,
		infoPtr *Info)

	WriteInfo func(
		pngPtr *Struct,
		infoPtr *Info)

	ReadInfo func(
		pngPtr *Struct,
		infoPtr *Info)

	ConvertToRfc1123 func(
		pngPtr *Struct,
		ptime *Time) *T.Char

	ConvertFromStructTm func(
		ptime *Time,
		ttime *T.Tm)

	ConvertFromTimeT func(
		ptime *Time,
		ttime T.TimeT)

	SetExpand func(
		pngPtr *Struct)

	SetExpandGray124To8 func(
		pngPtr *Struct)

	SetPaletteToRgb func(
		pngPtr *Struct)

	SetTRNSToAlpha func(
		pngPtr *Struct)

	SetBgr func(
		pngPtr *Struct)

	SetGrayToRgb func(
		pngPtr *Struct)

	SetRgbToGray func(
		pngPtr *Struct,
		errorAction int,
		red,
		green float64)

	SetRgbToGrayFixed func(
		pngPtr *Struct,
		errorAction int,
		red, green FixedPoint)

	GetRgbToGrayStatus func(
		pngPtr *Struct) byte

	BuildGrayscalePalette func(
		bitDepth int,
		palette *Color)

	SetStripAlpha func(
		pngPtr *Struct)

	SetSwapAlpha func(
		pngPtr *Struct)

	SetInvertAlpha func(
		pngPtr *Struct)

	SetFiller func(
		pngPtr *Struct,
		filler Uint32,
		flags int)

	SetAddAlpha func(
		pngPtr *Struct,
		filler Uint32,
		flags int)

	SetSwap func(
		pngPtr *Struct)

	SetPacking func(
		pngPtr *Struct)

	SetPackswap func(
		pngPtr *Struct)

	SetShift func(
		pngPtr *Struct,
		trueBits *Color8)

	SetInterlaceHandling func(
		pngPtr *Struct) int

	SetInvertMono func(
		pngPtr *Struct)

	SetBackground func(
		pngPtr *Struct,
		backgroundColor *Color16,
		backgroundGammaCode int,
		needExpand int,
		backgroundGamma float64)

	SetStrip16 func(
		pngPtr *Struct)

	SetQuantize func(
		pngPtr *Struct,
		palette *Color,
		numPalette int,
		maximumColors int,
		histogram *uint16,
		fullQuantize int)

	SetGamma func(
		pngPtr *Struct,
		screenGamma float64,
		defaultFileGamma float64)

	SetFlush func(
		pngPtr *Struct,
		nrows int)

	WriteFlush func(
		pngPtr *Struct)

	StartReadImage func(
		pngPtr *Struct)

	ReadUpdateInfo func(
		pngPtr *Struct,
		infoPtr *Info)

	ReadRows func(
		pngPtr *Struct,
		row **byte,
		displayRow **byte,
		numRows Uint32)

	ReadRow func(
		pngPtr *Struct,
		row *byte,
		displayRow *byte)

	ReadImage func(
		pngPtr *Struct,
		image **byte)

	WriteRow func(
		pngPtr *Struct,
		row *byte)

	WriteRows func(
		pngPtr *Struct,
		row **byte,
		numRows Uint32)

	WriteImage func(
		pngPtr *Struct,
		image **byte)

	WriteEnd func(
		pngPtr *Struct,
		infoPtr *Info)

	ReadEnd func(
		pngPtr *Struct,
		infoPtr *Info)

	DestroyInfoStruct func(
		pngPtr *Struct,
		infoPtrPtr **Info)

	DestroyReadStruct func(
		pngPtrPtr **Struct,
		infoPtrPtr **Info,
		endInfoPtrPtr **Info)

	DestroyWriteStruct func(
		pngPtrPtr **Struct,
		infoPtrPtr **Info)

	SetCrcAction func(
		pngPtr *Struct,
		critAction int,
		ancilAction int)

	SetFilter func(
		pngPtr *Struct,
		method int,
		filters int)

	SetFilterHeuristics func(
		pngPtr *Struct,
		heuristicMethod int,
		numWeights int,
		filterWeights Doublep,
		filterCosts Doublep)

	SetCompressionLevel func(
		pngPtr *Struct,
		level int)

	SetCompressionMemLevel func(
		pngPtr *Struct,
		memLevel int)

	SetCompressionStrategy func(
		pngPtr *Struct,
		strategy int)

	SetCompressionWindowBits func(
		pngPtr *Struct,
		windowBits int)

	SetCompressionMethod func(
		pngPtr *Struct,
		method int)

	InitIo func(
		pngPtr *Struct,
		fp *T.FILE)

	SetErrorFn func(
		pngPtr *Struct,
		errorPtr *T.Void,
		errorFn ErrorPtr,
		warningFn ErrorPtr)

	GetErrorPtr func(
		pngPtr *Struct) *T.Void

	SetWriteFn func(
		pngPtr *Struct,
		ioPtr *T.Void,
		writeDataFn RwPtr,
		outputFlushFn FlushPtr)

	SetReadFn func(
		pngPtr *Struct,
		ioPtr *T.Void,
		readDataFn RwPtr)

	GetIoPtr func(
		pngPtr *Struct) *T.Void

	SetReadStatusFn func(
		pngPtr *Struct,
		readRowFn ReadStatusPtr)

	SetWriteStatusFn func(
		pngPtr *Struct,
		writeRowFn WriteStatusPtr)

	SetMemFn func(
		pngPtr *Struct,
		memPtr *T.Void,
		mallocFn MallocPtr,
		freeFn FreePtr)

	GetMemPtr func(
		pngPtr *Struct) *T.Void

	SetReadUserTransformFn func(
		pngPtr *Struct,
		readUserTransformFn UserTransformPtr)

	SetWriteUserTransformFn func(
		pngPtr *Struct,
		writeUserTransformFn UserTransformPtr)

	SetUserTransformInfo func(
		pngPtr *Struct,
		userTransformPtr *T.Void,
		userTransformDepth int,
		userTransformChannels int)

	GetUserTransformPtr func(
		pngPtr *Struct) *T.Void

	SetReadUserChunkFn func(
		pngPtr *Struct,
		userChunkPtr *T.Void,
		readUserChunkFn UserChunkPtr)
	GetUserChunkPtr func(
		pngPtr *Struct) *T.Void

	SetProgressiveReadFn func(
		pngPtr *Struct,
		progressivePtr *T.Void,
		infoFn ProgressiveInfoPtr,
		rowFn ProgressiveRowPtr,
		endFn ProgressiveEndPtr)

	GetProgressivePtr func(
		pngPtr *Struct) *T.Void

	ProcessData func(
		pngPtr *Struct,
		infoPtr *Info,
		buffer *byte,
		bufferSize SizeT)

	ProgressiveCombineRow func(
		pngPtr *Struct,
		oldRow *byte,
		newRow *byte)

	Malloc func(
		pngPtr *Struct,
		size AllocSizeT) *T.Void

	Calloc func(
		pngPtr *Struct,
		size AllocSizeT) *T.Void

	MallocWarn func(
		pngPtr *Struct,
		size AllocSizeT) *T.Void

	Free func(
		pngPtr *Struct,
		ptr *T.Void)

	FreeData func(
		pngPtr *Struct,
		infoPtr *Info,
		freeMe Uint32,
		num int)

	DataFreer func(
		pngPtr *Struct,
		infoPtr *Info,
		freer int,
		mask Uint32)

	MallocDefault func(
		pngPtr *Struct,
		size AllocSizeT) *T.Void

	FreeDefault func(
		pngPtr *Struct,
		ptr *T.Void)

	Error func(
		pngPtr *Struct,
		errorMessage *T.Char)

	ChunkError func(
		pngPtr *Struct,
		errorMessage *T.Char)

	Warning func(
		pngPtr *Struct,
		warningMessage *T.Char)

	ChunkWarning func(
		pngPtr *Struct,
		warningMessage *T.Char)

	GetValid func(
		pngPtr *Struct,
		infoPtr *Info,
		flag Uint32) Uint32

	GetRowbytes func(
		pngPtr *Struct,
		infoPtr *Info) SizeT

	GetRows func(
		pngPtr *Struct,
		infoPtr *Info) **byte

	SetRows func(
		pngPtr *Struct,
		infoPtr *Info,
		rowPointers **byte)

	GetChannels func(
		pngPtr *Struct,
		infoPtr *Info) byte

	GetImageWidth func(
		pngPtr *Struct,
		infoPtr *Info) Uint32

	GetImageHeight func(
		pngPtr *Struct,
		infoPtr *Info) Uint32

	GetBitDepth func(
		pngPtr *Struct,
		infoPtr *Info) byte

	GetColorType func(
		pngPtr *Struct,
		infoPtr *Info) byte

	GetFilterType func(
		pngPtr *Struct,
		infoPtr *Info) byte

	GetInterlaceType func(
		pngPtr *Struct,
		infoPtr *Info) byte

	GetCompressionType func(
		pngPtr *Struct,
		infoPtr *Info) byte

	GetPixelsPerMeter func(
		pngPtr *Struct,
		infoPtr *Info) Uint32

	GetXPixelsPerMeter func(
		pngPtr *Struct,
		infoPtr *Info) Uint32

	GetYPixelsPerMeter func(
		pngPtr *Struct,
		infoPtr *Info) Uint32

	GetPixelAspectRatio func(
		pngPtr *Struct,
		infoPtr *Info) float32

	GetXOffsetPixels func(
		pngPtr *Struct,
		infoPtr *Info) Int32

	GetYOffsetPixels func(
		pngPtr *Struct,
		infoPtr *Info) Int32

	GetXOffsetMicrons func(
		pngPtr *Struct,
		infoPtr *Info) Int32

	GetYOffsetMicrons func(
		pngPtr *Struct,
		infoPtr *Info) Int32

	GetSignature func(
		pngPtr *Struct,
		infoPtr *Info) *byte

	GetBKGD func(
		pngPtr *Struct,
		infoPtr *Info,
		background **Color16) Uint32

	SetBKGD func(
		pngPtr *Struct,
		infoPtr *Info,
		background *Color16)

	GetCHRM func(
		pngPtr *Struct,
		infoPtr *Info,
		whiteX, whiteY,
		redX, redY,
		greenX, greenY,
		blueX, blueY *float64) Uint32

	GetCHRMFixed func(
		pngPtr *Struct,
		infoPtr *Info,
		intWhiteX *FixedPoint,
		intWhiteY *FixedPoint,
		intRedX *FixedPoint,
		intRedY *FixedPoint,
		intGreenX *FixedPoint,
		intGreenY *FixedPoint,
		intBlueX *FixedPoint,
		intBlueY *FixedPoint) Uint32

	SetCHRM func(
		pngPtr *Struct,
		infoPtr *Info,
		whiteX, whiteY,
		redX, redY,
		greenX, greenY,
		blueX, blueY float64)

	SetCHRMFixed func(
		pngPtr *Struct,
		infoPtr *Info,
		intWhiteX FixedPoint,
		intWhiteY FixedPoint,
		intRedX FixedPoint,
		intRedY FixedPoint,
		intGreenX FixedPoint,
		intGreenY FixedPoint,
		intBlueX FixedPoint,
		intBlueY FixedPoint)

	GetGAMA func(
		pngPtr *Struct,
		infoPtr *Info,
		fileGamma *float64) Uint32

	GetGAMAFixed func(
		pngPtr *Struct,
		infoPtr *Info,
		intFileGamma *FixedPoint) Uint32

	SetGAMA func(
		pngPtr *Struct,
		infoPtr *Info,
		fileGamma float64)

	SetGAMAFixed func(
		pngPtr *Struct,
		infoPtr *Info,
		intFileGamma FixedPoint)

	GetHIST func(
		pngPtr *Struct,
		infoPtr *Info,
		hist **uint16) Uint32

	SetHIST func(
		pngPtr *Struct,
		infoPtr *Info,
		hist *uint16)

	GetIHDR func(
		pngPtr *Struct,
		infoPtr *Info,
		width *Uint32,
		height *Uint32,
		bitDepth *int,
		colorType *int,
		interlaceMethod *int,
		compressionMethod *int,
		filterMethod *int) Uint32

	SetIHDR func(
		pngPtr *Struct,
		infoPtr *Info,
		width Uint32,
		height Uint32,
		bitDepth int,
		colorType int,
		interlaceMethod int,
		compressionMethod int,
		filterMethod int)

	GetOFFs func(
		pngPtr *Struct,
		infoPtr *Info,
		offsetX *Int32,
		offsetY *Int32,
		unitType *int) Uint32

	SetOFFs func(
		pngPtr *Struct,
		infoPtr *Info,
		offsetX Int32,
		offsetY Int32,
		unitType int)

	GetPCAL func(
		pngPtr *Struct,
		infoPtr *Info,
		purpose **T.Char,
		X0 *Int32,
		X1 *Int32,
		typ *int,
		nparams *int,
		units **T.Char,
		params ***T.Char) Uint32

	SetPCAL func(
		pngPtr *Struct,
		infoPtr *Info,
		purpose *T.Char,
		X0 Int32,
		X1 Int32,
		typ int,
		nparams int,
		units *T.Char,
		params **T.Char)

	GetPHYs func(
		pngPtr *Struct,
		infoPtr *Info,
		resX *Uint32,
		resY *Uint32,
		unitType *int) Uint32

	SetPHYs func(
		pngPtr *Struct,
		infoPtr *Info,
		resX Uint32,
		resY Uint32,
		unitType int)

	GetPLTE func(
		pngPtr *Struct,
		infoPtr *Info,
		palette **Color,
		numPalette *int) Uint32

	SetPLTE func(
		pngPtr *Struct,
		infoPtr *Info,
		palette *Color,
		numPalette int)

	GetSBIT func(
		pngPtr *Struct,
		infoPtr *Info,
		sigBit **Color8) Uint32

	SetSBIT func(
		pngPtr *Struct,
		infoPtr *Info,
		sigBit *Color8)

	GetSRGB func(
		pngPtr *Struct,
		infoPtr *Info,
		intent *int) Uint32

	SetSRGB func(
		pngPtr *Struct,
		infoPtr *Info,
		intent int)

	SetSRGBGAMAAndCHRM func(
		pngPtr *Struct,
		infoPtr *Info,
		intent int)

	GetICCP func(
		pngPtr *Struct,
		infoPtr *Info,
		name **T.Char,
		compressionType *int,
		profile **T.Char,
		proflen *Uint32) Uint32

	SetICCP func(
		pngPtr *Struct,
		infoPtr *Info,
		name *T.Char,
		compressionType int,
		profile *T.Char,
		proflen Uint32)

	GetSPLT func(
		pngPtr *Struct,
		infoPtr *Info,
		entries **SPLTT) Uint32

	SetSPLT func(
		pngPtr *Struct,
		infoPtr *Info,
		entries *SPLTT,
		nentries int)

	GetText func(
		pngPtr *Struct,
		infoPtr *Info,
		textPtr *Textp,
		numText *int) Uint32

	SetText func(
		pngPtr *Struct,
		infoPtr *Info,
		textPtr Textp,
		numText int)

	GetTIME func(
		pngPtr *Struct,
		infoPtr *Info,
		modTime **Time) Uint32

	SetTIME func(
		pngPtr *Struct,
		infoPtr *Info,
		modTime *Time)

	GetTRNS func(
		pngPtr *Struct,
		infoPtr *Info,
		transAlpha **byte,
		numTrans *int,
		transColor **Color16) Uint32

	SetTRNS func(
		pngPtr *Struct,
		infoPtr *Info,
		transAlpha *byte,
		numTrans int,
		transColor *Color16)

	GetSCAL func(
		pngPtr *Struct,
		infoPtr *Info,
		unit *int,
		width, height *float64) Uint32

	SetSCAL func(
		pngPtr *Struct,
		infoPtr *Info,
		unit int,
		width, height float64)

	SetKeepUnknownChunks func(
		pngPtr *Struct,
		Keep int,
		chunkList *byte,
		numChunks int)

	HandleAsUnknown func(
		pngPtr *Struct,
		chunkName *byte) int

	SetUnknownChunks func(
		pngPtr *Struct,
		infoPtr *Info,
		unknowns *UnknownChunk,
		numUnknowns int)

	SetUnknownChunkLocation func(
		pngPtr *Struct,
		infoPtr *Info,
		chunk int,
		location int)

	GetUnknownChunks func(
		pngPtr *Struct,
		infoPtr *Info,
		entries **UnknownChunk) Uint32

	SetInvalid func(
		pngPtr *Struct,
		infoPtr *Info,
		mask int)

	Read func(
		pngPtr *Struct,
		infoPtr *Info,
		transforms int,
		params *T.Void)

	Write func(
		pngPtr *Struct,
		infoPtr *Info,
		transforms int,
		params *T.Void)

	GetCopyright func(
		pngPtr *Struct) *T.Char

	GetHeaderVer func(
		pngPtr *Struct) *T.Char

	GetHeaderVersion func(
		pngPtr *Struct) *T.Char

	GetLibpngVer func(
		pngPtr *Struct) *T.Char

	PermitMngFeatures func(
		pngPtr *Struct,
		mngFeaturesPermitted Uint32) Uint32

	SetUserLimits func(
		pngPtr *Struct,
		userWidthMax Uint32,
		userHeightMax Uint32)

	GetUserWidthMax func(
		pngPtr *Struct) Uint32

	GetUserHeightMax func(
		pngPtr *Struct) Uint32

	SetChunkCacheMax func(
		pngPtr *Struct,
		userChunkCacheMax Uint32)

	GetChunkCacheMax func(
		pngPtr *Struct) Uint32

	SetChunkMallocMax func(
		pngPtr *Struct,
		userChunkCacheMax AllocSizeT)

	GetChunkMallocMax func(
		pngPtr *Struct) AllocSizeT

	GetIoState func(
		pngPtr *Struct) Uint32

	GetIoChunkName func(
		pngPtr *Struct) *byte

	GetUint31 func(
		pngPtr *Struct,
		buf *byte) Uint32

	SaveUint32 func(
		buf *byte,
		i Uint32)

	SaveInt32 func(
		buf *byte,
		i Int32)

	SaveUint16 func(
		buf *byte,
		i uint)
)

var dll = "libpng14-14.dll"

var apiList = outside.Apis{
	{"png_access_version_number", &AccessVersionNumber},
	{"png_build_grayscale_palette", &BuildGrayscalePalette},
	{"png_calloc", &Calloc},
	{"png_chunk_error", &ChunkError},
	{"png_chunk_warning", &ChunkWarning},
	{"png_convert_from_struct_tm", &ConvertFromStructTm},
	{"png_convert_from_time_t", &ConvertFromTimeT},
	{"png_convert_to_rfc1123", &ConvertToRfc1123},
	{"png_create_info_struct", &CreateInfoStruct},
	{"png_create_read_struct", &CreateReadStruct},
	{"png_create_read_struct_2", &CreateReadStruct2},
	{"png_create_write_struct", &CreateWriteStruct},
	{"png_create_write_struct_2", &CreateWriteStruct2},
	{"png_data_freer", &DataFreer},
	{"png_destroy_info_struct", &DestroyInfoStruct},
	{"png_destroy_read_struct", &DestroyReadStruct},
	{"png_destroy_write_struct", &DestroyWriteStruct},
	{"png_error", &Error},
	{"png_free", &Free},
	{"png_free_data", &FreeData},
	{"png_free_default", &FreeDefault},
	{"png_get_IHDR", &GetIHDR},
	{"png_get_PLTE", &GetPLTE},
	{"png_get_bKGD", &GetBKGD},
	{"png_get_bit_depth", &GetBitDepth},
	{"png_get_cHRM", &GetCHRM},
	{"png_get_cHRM_fixed", &GetCHRMFixed},
	{"png_get_channels", &GetChannels},
	{"png_get_chunk_cache_max", &GetChunkCacheMax},
	{"png_get_chunk_malloc_max", &GetChunkMallocMax},
	{"png_get_color_type", &GetColorType},
	{"png_get_compression_buffer_size", &GetCompressionBufferSize},
	{"png_get_compression_type", &GetCompressionType},
	{"png_get_copyright", &GetCopyright},
	{"png_get_error_ptr", &GetErrorPtr},
	{"png_get_filter_type", &GetFilterType},
	{"png_get_gAMA", &GetGAMA},
	{"png_get_gAMA_fixed", &GetGAMAFixed},
	{"png_get_hIST", &GetHIST},
	{"png_get_header_ver", &GetHeaderVer},
	{"png_get_header_version", &GetHeaderVersion},
	{"png_get_iCCP", &GetICCP},
	{"png_get_image_height", &GetImageHeight},
	{"png_get_image_width", &GetImageWidth},
	{"png_get_interlace_type", &GetInterlaceType},
	{"png_get_io_chunk_name", &GetIoChunkName},
	{"png_get_io_ptr", &GetIoPtr},
	{"png_get_io_state", &GetIoState},
	{"png_get_libpng_ver", &GetLibpngVer},
	{"png_get_mem_ptr", &GetMemPtr},
	{"png_get_oFFs", &GetOFFs},
	{"png_get_pCAL", &GetPCAL},
	{"png_get_pHYs", &GetPHYs},
	{"png_get_pixel_aspect_ratio", &GetPixelAspectRatio},
	{"png_get_pixels_per_meter", &GetPixelsPerMeter},
	{"png_get_progressive_ptr", &GetProgressivePtr},
	{"png_get_rgb_to_gray_status", &GetRgbToGrayStatus},
	{"png_get_rowbytes", &GetRowbytes},
	{"png_get_rows", &GetRows},
	{"png_get_sBIT", &GetSBIT},
	{"png_get_sCAL", &GetSCAL},
	{"png_get_sPLT", &GetSPLT},
	{"png_get_sRGB", &GetSRGB},
	{"png_get_signature", &GetSignature},
	{"png_get_tIME", &GetTIME},
	{"png_get_tRNS", &GetTRNS},
	{"png_get_text", &GetText},
	{"png_get_uint_31", &GetUint31},
	{"png_get_unknown_chunks", &GetUnknownChunks},
	{"png_get_user_chunk_ptr", &GetUserChunkPtr},
	{"png_get_user_height_max", &GetUserHeightMax},
	{"png_get_user_transform_ptr", &GetUserTransformPtr},
	{"png_get_user_width_max", &GetUserWidthMax},
	{"png_get_valid", &GetValid},
	{"png_get_x_offset_microns", &GetXOffsetMicrons},
	{"png_get_x_offset_pixels", &GetXOffsetPixels},
	{"png_get_x_pixels_per_meter", &GetXPixelsPerMeter},
	{"png_get_y_offset_microns", &GetYOffsetMicrons},
	{"png_get_y_offset_pixels", &GetYOffsetPixels},
	{"png_get_y_pixels_per_meter", &GetYPixelsPerMeter},
	{"png_handle_as_unknown", &HandleAsUnknown},
	{"png_info_init_3", &InfoInit3},
	{"png_init_io", &InitIo},
	{"png_malloc", &Malloc},
	{"png_malloc_default", &MallocDefault},
	{"png_malloc_warn", &MallocWarn},
	{"png_permit_mng_features", &PermitMngFeatures},
	{"png_process_data", &ProcessData},
	{"png_progressive_combine_row", &ProgressiveCombineRow},
	{"png_read_end", &ReadEnd},
	{"png_read_image", &ReadImage},
	{"png_read_info", &ReadInfo},
	{"png_read_png", &Read},
	{"png_read_row", &ReadRow},
	{"png_read_rows", &ReadRows},
	{"png_read_update_info", &ReadUpdateInfo},
	{"png_reset_zstream", &ResetZstream},
	{"png_save_int_32", &SaveInt32},
	{"png_save_uint_16", &SaveUint16},
	{"png_save_uint_32", &SaveUint32},
	{"png_set_IHDR", &SetIHDR},
	{"png_set_PLTE", &SetPLTE},
	{"png_set_add_alpha", &SetAddAlpha},
	{"png_set_bKGD", &SetBKGD},
	{"png_set_background", &SetBackground},
	{"png_set_bgr", &SetBgr},
	{"png_set_cHRM", &SetCHRM},
	{"png_set_cHRM_fixed", &SetCHRMFixed},
	{"png_set_chunk_cache_max", &SetChunkCacheMax},
	{"png_set_chunk_malloc_max", &SetChunkMallocMax},
	{"png_set_compression_buffer_size", &SetCompressionBufferSize},
	{"png_set_compression_level", &SetCompressionLevel},
	{"png_set_compression_mem_level", &SetCompressionMemLevel},
	{"png_set_compression_method", &SetCompressionMethod},
	{"png_set_compression_strategy", &SetCompressionStrategy},
	{"png_set_compression_window_bits", &SetCompressionWindowBits},
	{"png_set_crc_action", &SetCrcAction},
	{"png_set_error_fn", &SetErrorFn},
	{"png_set_expand", &SetExpand},
	{"png_set_expand_gray_1_2_4_to_8", &SetExpandGray124To8},
	{"png_set_filler", &SetFiller},
	{"png_set_filter", &SetFilter},
	{"png_set_filter_heuristics", &SetFilterHeuristics},
	{"png_set_flush", &SetFlush},
	{"png_set_gAMA", &SetGAMA},
	{"png_set_gAMA_fixed", &SetGAMAFixed},
	{"png_set_gamma", &SetGamma},
	{"png_set_gray_to_rgb", &SetGrayToRgb},
	{"png_set_hIST", &SetHIST},
	{"png_set_iCCP", &SetICCP},
	{"png_set_interlace_handling", &SetInterlaceHandling},
	{"png_set_invalid", &SetInvalid},
	{"png_set_invert_alpha", &SetInvertAlpha},
	{"png_set_invert_mono", &SetInvertMono},
	{"png_set_keep_unknown_chunks", &SetKeepUnknownChunks},
	{"png_set_longjmp_fn", &SetLongjmpFn},
	{"png_set_mem_fn", &SetMemFn},
	{"png_set_oFFs", &SetOFFs},
	{"png_set_pCAL", &SetPCAL},
	{"png_set_pHYs", &SetPHYs},
	{"png_set_packing", &SetPacking},
	{"png_set_packswap", &SetPackswap},
	{"png_set_palette_to_rgb", &SetPaletteToRgb},
	{"png_set_progressive_read_fn", &SetProgressiveReadFn},
	{"png_set_quantize", &SetQuantize},
	{"png_set_read_fn", &SetReadFn},
	{"png_set_read_status_fn", &SetReadStatusFn},
	{"png_set_read_user_chunk_fn", &SetReadUserChunkFn},
	{"png_set_read_user_transform_fn", &SetReadUserTransformFn},
	{"png_set_rgb_to_gray", &SetRgbToGray},
	{"png_set_rgb_to_gray_fixed", &SetRgbToGrayFixed},
	{"png_set_rows", &SetRows},
	{"png_set_sBIT", &SetSBIT},
	{"png_set_sCAL", &SetSCAL},
	{"png_set_sPLT", &SetSPLT},
	{"png_set_sRGB", &SetSRGB},
	{"png_set_sRGB_gAMA_and_cHRM", &SetSRGBGAMAAndCHRM},
	{"png_set_shift", &SetShift},
	{"png_set_sig_bytes", &SetSigbytes},
	{"png_set_strip_16", &SetStrip16},
	{"png_set_strip_alpha", &SetStripAlpha},
	{"png_set_swap", &SetSwap},
	{"png_set_swap_alpha", &SetSwapAlpha},
	{"png_set_tIME", &SetTIME},
	{"png_set_tRNS", &SetTRNS},
	{"png_set_tRNS_to_alpha", &SetTRNSToAlpha},
	{"png_set_text", &SetText},
	{"png_set_unknown_chunk_location", &SetUnknownChunkLocation},
	{"png_set_unknown_chunks", &SetUnknownChunks},
	{"png_set_user_limits", &SetUserLimits},
	{"png_set_user_transform_info", &SetUserTransformInfo},
	{"png_set_write_fn", &SetWriteFn},
	{"png_set_write_status_fn", &SetWriteStatusFn},
	{"png_set_write_user_transform_fn", &SetWriteUserTransformFn},
	{"png_sig_cmp", &SigCmp},
	{"png_start_read_image", &StartReadImage},
	{"png_warning", &Warning},
	{"png_write_chunk", &WriteChunk},
	{"png_write_chunk_data", &WriteChunkData},
	{"png_write_chunk_end", &WriteChunkEnd},
	{"png_write_chunk_start", &WriteChunkStart},
	{"png_write_end", &WriteEnd},
	{"png_write_flush", &WriteFlush},
	{"png_write_image", &WriteImage},
	{"png_write_info", &WriteInfo},
	{"png_write_info_before_PLTE", &WriteInfoBeforePLTE},
	{"png_write_png", &Write},
	{"png_write_row", &WriteRow},
	{"png_write_rows", &WriteRows},
	{"png_write_sig", &WriteSig},
}
