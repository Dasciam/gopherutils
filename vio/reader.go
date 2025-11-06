package vio

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// Reader is a wrapper around protocol.Reader that adds protocol version awareness.
// It embeds the standard protocol.Reader and extends it with version information
// to enable version-specific packet decoding logic.
type Reader struct {
	*protocol.Reader
	version int32
}

// NewReader creates a new version-aware Reader instance.
func NewReader(r *protocol.Reader, id int32) *Reader {
	return &Reader{Reader: r, version: id}
}

func (r *Reader) Version() int32 {
	return r.version
}

// Reads checks if the provided IO implementation is a version-aware Reader.
func Reads(io protocol.IO) bool {
	_, ok := io.(*Reader)
	return ok
}
