/*
Copyright © 2019 The OpenEBS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

var (
	// GitCommit that was compiled; filled in by
	// the compiler.
	GitCommit string

	// Version is the version of this repo; filled
	// in by the compiler
	Version string

	// VersionMeta is a pre-release marker for the
	// version. If this is "" (empty string) then
	// it means that it is a final release. Otherwise,
	// this is a pre-release such as "dev" (in
	// development), "beta", "rc1", etc.
	VersionMeta string
)

const (
	versionFile   string = "/src/github.com/openebs/zfs-localpv/VERSION"
	buildMetaFile string = "/src/github.com/openebs/zfs-localpv/BUILDMETA"
)

// Current returns current version of csi driver
func Current() string {
	_ = "STUB: not implemented"

	// Get returns current version from global
	// Version variable. If Version is unset then
	// from VERSION file at the root of this repo.
	return ""
}

func Get() string { _ = "STUB: not implemented"; return "" }

// GetBuildMeta returns build type from
// global VersionMeta variable. If VersionMeta
// is unset then this is fetched from BUILDMETA
// file at the root of this repo.
func GetBuildMeta() string { _ = "STUB: not implemented"; return "" }

// GetGitCommit returns Git commit SHA-1 from
// global GitCommit variable. If GitCommit is
// unset this calls Git directly.
func GetGitCommit() string { _ = "STUB: not implemented"; return "" }

// GetVersionDetails return version info from git commit
func GetVersionDetails() string { _ = "STUB: not implemented"; return "" }

// Verbose returns version details with git
// commit info
func Verbose() string { _ = "STUB: not implemented"; return "" }
