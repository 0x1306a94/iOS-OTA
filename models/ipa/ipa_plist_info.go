package ipa

type IpaPlistInfo struct {
	Name string `plist:"CFBundleName"`
	Version string `plist:"CFBundleShortVersionString"`
	Identifier string `plist:"CFBundleIdentifier"`
	Build string `plist:"CFBundleVersion"`
}