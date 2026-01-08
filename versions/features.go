package versions

const (
	// MaceImplemented is the protocol version where the Mace item was introduced (1.21.0).
	MaceImplemented = Protocol1_21_0
	// SpearImplemented is the protocol version where the Spear item was introduced (1.21.130).
	SpearImplemented = Protocol1_21_130
)

// featureMap maps feature string IDs to their implementation protocol version.
var featureMap = map[string]int32{}

// RegisterFeature adds a feature to the registry with a string ID.
// Not thread safe, should run during initialisation.
func RegisterFeature(id string, protocol int32) {
	featureMap[id] = protocol
}

// FeatureByID returns the protocol version where a feature was implemented.
func FeatureByID(id string) (int32, bool) {
	protocol, ok := featureMap[id]
	return protocol, ok
}

// HasFeatureByID returns true if the given protocol version supports the feature with the specified ID.
func HasFeatureByID(protocol int32, id string) bool {
	feature, ok := FeatureByID(id)
	if !ok {
		return false
	}
	return protocol >= feature
}

// HasFeature returns true if the given protocol version supports the specified feature.
func HasFeature(protocol, feature int32) bool {
	return protocol >= feature
}

// HasMace returns true if the given protocol version supports the Mace item.
func HasMace(protocol int32) bool {
	return HasFeature(protocol, MaceImplemented)
}

// HasSpear returns true if the given protocol version supports the Spear item.
func HasSpear(protocol int32) bool {
	return HasFeature(protocol, SpearImplemented)
}

func init() {
	RegisterFeature("mace_implemented", MaceImplemented)
	RegisterFeature("spear_implemented", SpearImplemented)
}
