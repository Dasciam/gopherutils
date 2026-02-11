package versions

const (
	// MaceImplemented is the protocol version where the Mace item was introduced (1.21.0).
	MaceImplemented = Protocol1_21_0
	// SpearImplemented is the protocol version where the Spear item was introduced (1.21.130).
	SpearImplemented = Protocol1_21_130
	// ClientDamageTiltConfiguringImplemented is the protocol version where the client can configure
	// damage tilt strength in client settings (1.21.130).
	ClientDamageTiltConfiguringImplemented = Protocol1_21_130
	// DebugShapeEntityLinkingImplemented is the protocol version where the server can link debug shapes
	// to entities (1.26.0).
	DebugShapeEntityLinkingImplemented = Protocol1_26_0
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

func init() {
	RegisterFeature("mace_implemented", MaceImplemented)
	RegisterFeature("spear_implemented", SpearImplemented)
	RegisterFeature("client_damage_tilt_configuring_implemented", ClientDamageTiltConfiguringImplemented)
	RegisterFeature("debug_shape_entity_linking_implemented", DebugShapeEntityLinkingImplemented)
}
