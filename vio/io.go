package vio

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// IO is an extended interface that builds upon protocol.IO to provide
// version-aware I/O operations for Minecraft protocol handling.
// It allows access to the specific protocol version being used, which is
// essential for version-specific packet encoding/decoding logic.
type IO interface {
	protocol.IO
	// Version returns the specific protocol version implemented by this IO.
	// This is typically used to determine how to encode/decode packets
	// for different Minecraft versions.
	Version() int32
}

// Version returns the protocol version from an IO implementation if it
// supports version checking. This is a safe utility function that handles
// both version-aware and standard IO implementations gracefully.
//
// If the provided IO implements the versioned IO interface, returns its
// specific version. Otherwise, defaults to the current protocol version
// defined in the protocol package.
func Version(io protocol.IO) int32 {
	p, ok := io.(IO)
	if !ok {
		return protocol.CurrentProtocol
	}
	return p.Version()
}
