package build

import (
	"fmt"
	"runtime/debug"
	"strings"
)

// VersionInformation contains everything we know about our appliation.
//
// Actually, that's a lie, it does not contain everything (yet), but contains
// a lot of information.
type VersionInformation struct {
	AppName  string
	Version  string
	Time     string
	Revision string
	Dirty    bool
}

func (info VersionInformation) String() string {
	parts := []string{}

	if info.AppName != "" {
		parts = append(parts, info.AppName)
	}

	if info.Version != "" {
		parts = append(parts, info.Version)
	} else if info.Revision != "" {
		parts = append(parts, info.Revision)
	}

	if info.Dirty {
		parts = append(parts, "*dirty*")
	}

	if info.Time != "" {
		parts = append(parts, fmt.Sprintf("(%s)", info.Time))
	}

	return strings.Join(parts, " ")
}

// ReadVersion reads build information and tried to create a VersionInformation
// out of it.
func ReadVersion(appName string) (VersionInformation, error) {
	version := VersionInformation{AppName: appName}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return version, BuildInfoError{}
	}

	for _, setting := range info.Settings {
		if setting.Key == "-ldflags" && strings.HasPrefix(setting.Value, "-buildid=") {
			version.Version = strings.TrimPrefix(setting.Value, "-buildid=")
		} else if setting.Key == "vcs.revision" {
			version.Revision = setting.Value
		} else if setting.Key == "vcs.time" {
			version.Time = setting.Value
		} else if setting.Key == "vcs.modified" {
			version.Dirty = true
		}
	}

	return version, nil
}
