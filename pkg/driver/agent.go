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
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
	"github.com/openebs/zfs-localpv/pkg/zfs"
	"golang.org/x/net/context"
)

// node is the server implementation
// for CSI NodeServer
type node struct {
	driver *CSIDriver
	csi.UnimplementedNodeServer
}

// NewNode returns a new instance
// of CSI NodeServer
func NewNode(d *CSIDriver) csi.NodeServer { _ = "STUB: not implemented"; return *new(csi.NodeServer) }

// set up signals so we handle the first shutdown signal gracefully
// TODO: (tech-debt) Setup signal handler more above, several files want to use stopCh and this function is only allowed to be used once (see #647)
// Affected files: pkg/driver/agent.go pkg/driver/controller.go pkg/driver/grpc.go

// start the zfsnode resource watcher

// start the zfsvolume watcher

// start the snapshot watcher

// start the backup controller

// start the restore controller

// GetVolAndMountInfo get volume and mount info from node csi volume request
func GetVolAndMountInfo(
	req *csi.NodePublishVolumeRequest,
) (*apis.ZFSVolume, *zfs.MountInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NodePublishVolume publishes (mounts) the volume
// at the corresponding node at a given path
//
// This implements csi.NodeServer
func (ns *node) NodePublishVolume(
	ctx context.Context,
	req *csi.NodePublishVolumeRequest,
) (*csi.NodePublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the access type is block, do nothing for stage

// attempt block mount operation on the requested path

// attempt filesystem mount operation on the requested path

// NodeUnpublishVolume unpublishes (unmounts) the volume
// from the corresponding node from the given path
//
// This implements csi.NodeServer
func (ns *node) NodeUnpublishVolume(
	ctx context.Context,
	req *csi.NodeUnpublishVolumeRequest,
) (*csi.NodeUnpublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NodeGetInfo returns node details
//
// This implements csi.NodeServer
func (ns *node) NodeGetInfo(
	ctx context.Context,
	req *csi.NodeGetInfoRequest,
) (*csi.NodeGetInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
 * The driver will support all the keys and values defined in the node's label.
 * if nodes are labeled with the below keys and values
 * map[beta.kubernetes.io/arch:amd64 beta.kubernetes.io/os:linux kubernetes.io/arch:amd64 kubernetes.io/hostname:pawan-node-1 kubernetes.io/os:linux node-role.kubernetes.io/worker:true openebs.io/zone:zone1 openebs.io/zpool:ssd]
 * The driver will support below key and values
 * {
 *	beta.kubernetes.io/arch:amd64
 *	beta.kubernetes.io/os:linux
 *	kubernetes.io/arch:amd64
 *	kubernetes.io/hostname:pawan-node-1
 *	kubernetes.io/os:linux
 *	node-role.kubernetes.io/worker:true
 *	openebs.io/zone:zone1
 *	openebs.io/zpool:ssd
 * }
 */

// support topologykeys from env ALLOWED_TOPOLOGIES

// add driver's topology key if not labelled already

// add old topology key to support backward compatibility for velero

// NodeGetCapabilities returns capabilities supported
// by this node service
//
// This implements csi.NodeServer
func (ns *node) NodeGetCapabilities(
	ctx context.Context,
	req *csi.NodeGetCapabilitiesRequest,
) (*csi.NodeGetCapabilitiesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO
// This needs to be implemented
//
// NodeStageVolume mounts the volume on the staging
// path
//
// This implements csi.NodeServer
func (ns *node) NodeStageVolume(
	ctx context.Context,
	req *csi.NodeStageVolumeRequest,
) (*csi.NodeStageVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NodeUnstageVolume unmounts the volume from
// the staging path
//
// This implements csi.NodeServer
func (ns *node) NodeUnstageVolume(
	ctx context.Context,
	req *csi.NodeUnstageVolumeRequest,
) (*csi.NodeUnstageVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO
// Verify if this needs to be implemented
//
// # NodeExpandVolume resizes the filesystem if required
//
// If ControllerExpandVolumeResponse returns true in
// node_expansion_required then FileSystemResizePending
// condition will be added to PVC and NodeExpandVolume
// operation will be queued on kubelet
//
// This implements csi.NodeServer
func (ns *node) NodeExpandVolume(
	ctx context.Context,
	req *csi.NodeExpandVolumeRequest,
) (*csi.NodeExpandVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// find if it is block device so that we don't attempt filesystem resize

// doing this dirty check as volume capabilities are not passed for NodeExpandVolume
// CSI 1.2 spec will probably solve this

// it is not a block device, resize the filesystem

// error defined in zfs source code: https://github.com/openzfs/zfs/blob/zfs-2.3.99/lib/libzfs/libzfs_util.c#L551

// NodeGetVolumeStats returns statistics for the
// given volume
func (ns *node) NodeGetVolumeStats(
	ctx context.Context,
	req *csi.NodeGetVolumeStatsRequest,
) (*csi.NodeGetVolumeStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *node) validateNodePublishReq(
	req *csi.NodePublishVolumeRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (ns *node) validateNodeUnpublishReq(
	req *csi.NodeUnpublishVolumeRequest,
) error {
	_ = "STUB: not implemented"
	return nil
}
