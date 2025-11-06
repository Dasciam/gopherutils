package gopherutils

import _ "unsafe"

// RemoveChemistry removes chemistry by clearing exemptedPacks slice in
// gopher tunnel. Intended to be particle fix.
func RemoveChemistry() {
	exemptedPacks = nil
}

//go:linkname exemptedPacks github.com/sandertv/gophertunnel/minecraft.exemptedPacks
//goland:noinspection GoUnusedGlobalVariable
var exemptedPacks []struct{}
