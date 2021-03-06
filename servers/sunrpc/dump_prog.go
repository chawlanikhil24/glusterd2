package sunrpc

const (
	dumpProgNum     = 123451501
	dumpProgVersion = 1
)

const (
	_ = iota
	gfDumpDump
	gfDumpPing
)

// GfDump is a type for GlusterFS Dump RPC program
type GfDump genericProgram

func newGfDump() *GfDump {
	// rpc/rpc-lib/src/xdr-common.h
	return &GfDump{
		name:        "GF-DUMP",
		progNum:     dumpProgNum,
		progVersion: dumpProgVersion,
		procedures: []Procedure{
			Procedure{gfDumpDump, "Dump"}, // GF_DUMP_DUMP
			Procedure{gfDumpPing, "Ping"}, // GF_DUMP_PING
		},
	}
}

// Name returns the name of the RPC program
func (p *GfDump) Name() string {
	return p.name
}

// Number returns the RPC Program number
func (p *GfDump) Number() uint32 {
	return p.progNum
}

// Version returns the RPC program version number
func (p *GfDump) Version() uint32 {
	return p.progVersion
}

// Procedures returns a list of procedures provided by the RPC program
func (p *GfDump) Procedures() []Procedure {
	return p.procedures
}

// GfDumpReq is request sent by the client
type GfDumpReq struct {
	GfsID uint64
}

// GfProcDetail contains details for individual RPC program
type GfProcDetail struct {
	ProgName string
	ProgNum  uint64
	ProgVer  uint64
	Next     *GfProcDetail `xdr:"optional"`
}

// GfDumpRsp is response sent by server. It contains a list of GfProcDetail
type GfDumpRsp struct {
	GfsID   uint64
	OpRet   int
	OpErrno int
	Prog    *GfProcDetail `xdr:"optional"`
}

// Dump will return a list of all available RPC programs
func (p *GfDump) Dump(args *GfDumpReq, reply *GfDumpRsp) error {

	// TODO: I don't like doing this in Go. Should abstract it.
	var list *GfProcDetail
	var trav *GfProcDetail

	for _, p := range programsList {
		tmp := &GfProcDetail{
			ProgName: p.Name(),
			ProgNum:  uint64(p.Number()),
			ProgVer:  uint64(p.Version()),
		}
		if list == nil {
			list = tmp
			trav = list
		} else {
			trav.Next = tmp
			trav = trav.Next
		}
	}
	reply.Prog = list

	return nil
}

// Ping is for availability check
func (p *GfDump) Ping(_ *struct{}, reply *GfCommonRsp) error {

	return nil
}
