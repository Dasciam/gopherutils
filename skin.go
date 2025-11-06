package gopherutils

// SkinGeometry returns default skin geometry.
// It usually contains:
// - geometry.cape: cape model
// - geometry.humanoid.custom: wide arms skin model
// - geometry.humanoid.customSlim: slim arms skin model
// geometries.
func SkinGeometry() []byte {
	return skinGeometry
}

//go:linkname skinGeometry github.com/sandertv/gophertunnel/minecraft.skinGeometry
var skinGeometry []byte
