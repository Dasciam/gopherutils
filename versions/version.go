package versions

import (
	"iter"
	"maps"
)

// Version represents a Minecraft version with its protocol number and version string.
type Version struct {
	// Protocol is the numeric protocol identifier used in Minecraft network communication.
	// This is the official protocol number that clients and servers use.
	// Example: 589 for Minecraft 1.20.0
	Protocol int32
	// Version is the human-readable Minecraft version string in semantic version format.
	// This follows the standard Minecraft version numbering scheme (major.minor.patch).
	// Example: "1.20.0", "1.21.100"
	Version string
}

var (
	// versionMap maps version strings to Version objects for quick lookup by version name.
	versionMap = map[string]Version{}
	// protocolMap maps protocol numbers to Version objects for quick lookup by protocol ID.
	protocolMap = map[int32]Version{}
)

// Register adds a new version to the version registry, making it available for lookup
// by both protocol number and version string.
// Not thread safe, should run during initialisation.
func Register(protocol int32, version string) {
	object := Version{
		Protocol: protocol,
		Version:  version,
	}
	versionMap[version] = object
	protocolMap[protocol] = object
}

// ByVersion returns the Version object for a given version string.
// This is useful for looking up protocol numbers when you have a version string.
func ByVersion(version string) (Version, bool) {
	v, ok := versionMap[version]
	return v, ok
}

// ByProtocol returns the Version object for a given protocol number.
// This is useful for looking up version strings when you have a protocol number.
func ByProtocol(protocol int32) (Version, bool) {
	v, ok := protocolMap[protocol]
	return v, ok
}

// Versions returns a slice of all registered Version objects.
// The versions are returned in an arbitrary order. Use this function when you need
// to iterate through all available versions or display a list of supported versions.
func Versions() iter.Seq[Version] {
	return maps.Values(versionMap)
}

func init() {
	Register(Protocol1_20_0, "1.20.0")
	Register(Protocol1_20_10, "1.20.10")
	Register(Protocol1_20_30, "1.20.30")
	Register(Protocol1_20_40, "1.20.40")
	Register(Protocol1_20_50, "1.20.50")
	Register(Protocol1_20_60, "1.20.60")
	Register(Protocol1_20_70, "1.20.70")
	Register(Protocol1_20_80, "1.20.80")
	Register(Protocol1_21_0, "1.21.0")
	Register(Protocol1_21_2, "1.21.2")
	Register(Protocol1_21_20, "1.21.20")
	Register(Protocol1_21_30, "1.21.30")
	Register(Protocol1_21_40, "1.21.40")
	Register(Protocol1_21_50, "1.21.50")
	Register(Protocol1_21_60, "1.21.60")
	Register(Protocol1_21_70, "1.21.70")
	Register(Protocol1_21_80, "1.21.80")
	Register(Protocol1_21_90, "1.21.90")
	Register(Protocol1_21_93, "1.21.93")
	Register(Protocol1_21_100, "1.21.100")
	Register(Protocol1_21_111, "1.21.111")
	Register(Protocol1_21_120, "1.21.120")
	Register(Protocol1_21_124, "1.21.124")
	Register(Protocol1_21_130, "1.21.130")
}
