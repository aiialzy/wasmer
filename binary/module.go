package binary

const (
	SecCustomID = iota
	SecTypeID
	SecImportID
	SecFuncID
	SecTableID
	SecMemID
	SecGlobalID
	SecExportID
	SecStartID
	SecElemID
	SecCodeID
	SecDataID
)

const (
	ImportTagFunc   = 0
	ImportTagTable  = 1
	ImportTagMem    = 2
	ImportTagGlobal = 3
)
const (
	ExportTagFunc   = 0
	ExportTagTable  = 1
	ExportTagMem    = 2
	ExportTagGlobal = 3
)

const (
	MagicNumber = 0x6D736100
	Version     = 0x00000001
)

type (
	TypeIdx   = uint32
	FuncIdx   = uint32
	TableIdx  = uint32
	MemIdx    = uint32
	GlobalIdx = uint32
	LocalIdx  = uint32
	LabelIdx  = uint32
)

type Module struct {
	Magic      uint32
	Version    uint32
	CustomSecs []CustomSec
	TypeSec    []FuncType
	ImportSec  []Import
	FuncSec    []TypeIdx
	TableSec   []TableType
	MemSec     []MemType
	GlobalSec  []Global
	ExportSec  []Export
	StartSec   *FuncIdx
	ElemSec    []Elem
	CodeSec    []Code
	DataSec    []Data
}

type CustomSec struct {
	Name  string
	Bytes []byte
}

type Import struct {
	Module string
	Name   string
	Desc   ImportDesc
}

type ImportDesc struct {
	Tag      byte
	FuncType TypeIdx
	Table    TableType
	Mem      MemType
	Global   GlobalType
}

type Global struct {
	Type GlobalType
	Init Expr
}

type Export struct {
	Name string
	Desc ExportDesc
}

type ExportDesc struct {
	Tag byte
	Idx uint32
}

type Elem struct {
	Table  TableIdx
	Offset Expr
	Init   []FuncIdx
}

type Code struct {
	Locals []Locals
	Expr   Expr
}

type Locals struct {
	N    uint32
	Type ValType
}

type Data struct {
	Mem    MemIdx
	Offset Expr
	Init   []byte
}

func (code Code) GetLocalCount() uint64 {
	n := uint64(0)
	for _, locals := range code.Locals {
		n += uint64(locals.N)
	}

	return n
}
