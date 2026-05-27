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
	"golang.org/x/net/context"
)

// identity is the server implementation
// for CSI IdentityServer
type identity struct {
	driver *CSIDriver
	csi.UnimplementedIdentityServer
}

// NewIdentity returns a new instance of CSI
// IdentityServer
func NewIdentity(d *CSIDriver) csi.IdentityServer {
	_ = "STUB: not implemented"
	return *new(csi.IdentityServer)
}

// GetPluginInfo returns the version and name of
// this service
//
// This implements csi.IdentityServer
func (id *identity) GetPluginInfo(
	ctx context.Context,
	req *csi.GetPluginInfoRequest,
) (*csi.GetPluginInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO
// verify which version needs to be used:
// config.version or version.Current()

// TODO
// Need to implement this
//
// # Probe checks if the plugin is running or not
//
// This implements csi.IdentityServer
func (id *identity) Probe(
	ctx context.Context,
	req *csi.ProbeRequest,
) (*csi.ProbeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPluginCapabilities returns supported capabilities
// of this plugin
//
// Currently it reports whether this plugin can serve
// the Controller interface. Controller interface methods
// are called dependant on this
//
// This implements csi.IdentityServer
func (id *identity) GetPluginCapabilities(
	ctx context.Context,
	req *csi.GetPluginCapabilitiesRequest,
) (*csi.GetPluginCapabilitiesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
