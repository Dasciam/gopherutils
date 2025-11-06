package vio

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// Writer is a wrapper around protocol.Writer that adds protocol version awareness.
// It embeds the standard protocol.Writer and extends it with version information
// to enable version-specific packet encoding logic.
type Writer struct {
	*protocol.Writer
	version int32
}

// NewWriter creates a new version-aware Writer instance.
func NewWriter(r *protocol.Writer, id int32) *Writer {
	return &Writer{Writer: r, version: id}
}

func (w *Writer) Version() int32 {
	return w.version
}

// Writes checks if the provided IO implementation is a version-aware Writer.
func Writes(io protocol.IO) bool {
	_, ok := io.(*Writer)
	return ok
}
