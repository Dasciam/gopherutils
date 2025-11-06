package translator

import (
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// Multi is a composite Translator that chains multiple translators together.
// It applies each translator in sequence, passing the output of one translator
// as the input to the next. This allows for building complex translation pipelines
// from simple, single-purpose translators.
type Multi struct {
	// Translators is a slice of Translator implementations that will be applied
	// in sequence. Each translator receives the packet as modified by the previous
	// translator in the slice.
	Translators []Translator
}

// Translate implements the Translator interface by sequentially applying all
// configured translators. Each translator in the Multi.Translators slice is
// called with the packet returned by the previous translator.
func (m Multi) Translate(pk packet.Packet, conn *minecraft.Conn) (packet.Packet, error) {
	var err error

	for _, translator := range m.Translators {
		pk, err = translator.Translate(pk, conn)
		if err != nil {
			break
		}
	}
	return pk, err
}
