package protocol

import "sync"

var (
	RequestPool  sync.Pool
	ResponsePool sync.Pool
)

func init() {
	RequestPool = sync.Pool{New: func() interface{} {
		return &Header{}
	}}
	ResponsePool = sync.Pool{New: func() interface{} {
		return &Header{}
	}}
}

func (h *Header) Reset() {
	h.MagicNumber = 0
	h.Status = 0
	h.CompressType = 0
	h.ServiceMethod = ""
	h.ServicePath = ""
	h.ID = 0
	h.PayloadLen = 0
	h.Checksum = 0
}
