package regscontract

import "context"

//Request request
type Request struct {
	Username string `json:"username,omitempty"`
	DumpReq  []byte `json:"dump_req,omitempty"`
	DumpRes  []byte `json:"dump_res,omitempty"`
}

//RegsService regservice interface
type RegsService interface {
	Reg(ctx context.Context, req *Request) error
}
