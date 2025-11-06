package translator

import (
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// Translator is an interface that defines a contract for translating Minecraft packets
// between different protocol versions, formats, or for modifying packet contents.
// Implementations of this interface are typically used in proxy servers, cross-version
// compatibility layers, or packet manipulation tools to enable communication between
// clients and servers running different protocol versions.
type Translator interface {
	// Translate takes an incoming packet and connection context, and returns a translated
	// version of the packet. The method may modify the packet contents, convert it to a
	// different protocol version, or even replace it with a different packet entirely.
	Translate(pk packet.Packet, conn *minecraft.Conn) (packet.Packet, error)
}

// Compile-time check.
var _ Translator = Nop{}

// Nop is a no-operation Translator implementation that acts as a pass-through.
type Nop struct{}

func (Nop) Translate(pk packet.Packet, _ *minecraft.Conn) (packet.Packet, error) { return pk, nil }
