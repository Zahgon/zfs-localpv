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

package driver

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	config "github.com/openebs/zfs-localpv/pkg/config"
)

// CSIDriver defines a common data structure
// for drivers
// TODO check if this can be renamed to Base
type CSIDriver struct {
	// TODO change the field names to make it
	// readable
	config *config.Config
	ids    csi.IdentityServer
	ns     csi.NodeServer
	cs     csi.ControllerServer

	cap []*csi.VolumeCapability_AccessMode
}

// GetVolumeCapabilityAccessModes fetches the access
// modes on which the volume can be exposed
func GetVolumeCapabilityAccessModes() []*csi.VolumeCapability_AccessMode {
	_ = "STUB: not implemented"
	return nil
}

func newVolumeCapabilityAccessMode(mode csi.VolumeCapability_AccessMode_Mode) *csi.VolumeCapability_AccessMode {
	_ = "STUB: not implemented"
	return nil
}

// New returns a new driver instance
func New(config *config.Config) *CSIDriver { _ = "STUB: not implemented"; return nil }

// Start monitor goroutine to monitor the
// ZfsVolume CR. If there is any event
// related to the volume like destroy or
// property change, handle it accordingly.

// Identity server is common to both node and
// controller, it is required to register,
// share capabilities and probe the corresponding
// driver

// Run starts the CSI plugin by communicating
// over the given endpoint
func (d *CSIDriver) Run() error {
	_ = "STUB: not implemented"
	// Initialize and start listening on grpc server
	return nil
}
