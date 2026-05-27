// Copyright © 2019 The OpenEBS Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zfs

import (
	"context"
	"os"

	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
	"k8s.io/klog/v2"
)

const (
	// OpenEBSNamespaceKey is the environment variable to get openebs namespace
	//
	// This environment variable is set via kubernetes downward API
	OpenEBSNamespaceKey string = "OPENEBS_NAMESPACE"
	// GoogleAnalyticsKey This environment variable is set via env
	GoogleAnalyticsKey string = "OPENEBS_IO_ENABLE_ANALYTICS"
	// ZFSBackupGCEnabledKey This environment variable is set via env
	ZFSBackupGCEnabledKey string = "OPENEBS_IO_ENABLE_BACKUP_GC"
	// ZFSFinalizer for the ZfsVolume CR
	ZFSFinalizer string = "zfs.openebs.io/finalizer"
	// ZFSVolKey for the ZfsSnapshot CR to store Persistence Volume name
	ZFSVolKey string = "openebs.io/persistent-volume"
	// ZFSSrcVolKey key for the source Volume name
	ZFSSrcVolKey string = "openebs.io/source-volume"
	// PoolNameKey is key for ZFS pool name
	PoolNameKey string = "openebs.io/poolname"
	// ZFSNodeKey will be used to insert Label in ZfsVolume CR
	ZFSNodeKey string = "kubernetes.io/nodename"
	// ZFSTopologyKey is supported topology key for the zfs driver
	ZFSTopologyKey string = "openebs.io/nodeid"
	// ZFSTopoNodenameKey is supported topology key for the zfs driver
	ZFSTopoNodenameKey string = "openebs.io/nodename"
	// ZFSStatusPending shows object has not handled yet
	ZFSStatusPending string = "Pending"
	// ZFSStatusFailed shows object operation has failed
	ZFSStatusFailed string = "Failed"
	// ZFSStatusReady shows object has been processed
	ZFSStatusReady string = "Ready"
	// OpenEBSCasTypeKey for the cas-type label
	OpenEBSCasTypeKey string = "openebs.io/cas-type"
	// ZFSCasTypeName for the name of the cas-type
	ZFSCasTypeName string = "localpv-zfs"
)

var (
	// OpenEBSNamespace is openebs system namespace
	OpenEBSNamespace string

	// NodeID is the NodeID of the node on which the pod is present
	NodeID string

	// GoogleAnalyticsEnabled should send google analytics or not
	GoogleAnalyticsEnabled string

	// ZFSBackupGCEnabled is the flag to enable backup garbage collection
	ZFSBackupGCEnabled bool
)

func init() {
	var err error

	OpenEBSNamespace = os.Getenv(OpenEBSNamespaceKey)

	if os.Getenv("OPENEBS_NODE_DRIVER") != "" {
		if OpenEBSNamespace == "" {
			klog.Fatalf("OPENEBS_NAMESPACE environment variable not set for daemonset")
		}
		nodename := os.Getenv("OPENEBS_NODE_NAME")
		if nodename == "" {
			klog.Fatalf("OPENEBS_NODE_NAME environment variable not set")
		}
		if NodeID, err = GetNodeID(nodename); err != nil {
			klog.Fatalf("GetNodeID failed for node=%s err: %s", nodename, err.Error())
		}
		klog.Infof("zfs: node(%s) has node affinity %s=%s", nodename, ZFSTopologyKey, NodeID)
	} else if os.Getenv("OPENEBS_CONTROLLER_DRIVER") != "" {
		if OpenEBSNamespace == "" {
			klog.Fatalf("OPENEBS_NAMESPACE environment variable not set for controller")
		}
	}

	GoogleAnalyticsEnabled = os.Getenv(GoogleAnalyticsKey)

	if os.Getenv(ZFSBackupGCEnabledKey) == "true" {
		ZFSBackupGCEnabled = true
	} else {
		ZFSBackupGCEnabled = false
	}
}

// GetNodeID returns the Node ID of the given K8s nodename.
// It may return an error whilst fetching the node using the k8sapi.
// If the K8s node object does contain the topology label, then the nodename
// itself is returned as the Node ID.
func GetNodeID(nodename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// node is not labelled, use node name as nodeid

func checkVolCreation(ctx context.Context, volname string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ProvisionVolume creates a ZFSVolume(zv) CR,
// watcher for zvc is present in CSI agent
func ProvisionVolume(
	ctx context.Context,
	vol *apis.ZFSVolume,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// update the spec and status

// ResizeVolume resizes the zfs volume
func ResizeVolume(vol *apis.ZFSVolume, newSize int64) error { _ = "STUB: not implemented"; return nil }

// ProvisionSnapshot creates a ZFSSnapshot CR,
// watcher for zvc is present in CSI agent
func ProvisionSnapshot(
	snap *apis.ZFSSnapshot,
) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteSnapshot deletes the corresponding ZFSSnapshot CR
func DeleteSnapshot(snapname string) (err error) { _ = "STUB: not implemented"; return nil }

// DeleteVolume deletes the corresponding ZFSVol CR
func DeleteVolume(volumeID string) (err error) { _ = "STUB: not implemented"; return nil }

// GetVolList fetches the current Published Volume list
func GetVolList(volumeID string) (*apis.ZFSVolumeList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetZFSVolume fetches the given ZFSVolume
func GetZFSVolume(volumeID string) (*apis.ZFSVolume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateZFSVolumeAnnotation updtates the ZFSVolume CR with the given annotation
func UpdateZFSVolumeAnnotation(vol *apis.ZFSVolume) error { _ = "STUB: not implemented"; return nil }

// GetZFSVolumeState returns ZFSVolume OwnerNode and State for
// the given volume. CreateVolume request may call it again and
// again until volume is "Ready".
func GetZFSVolumeState(volID string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// UpdateZvolInfo updates ZFSVolume CR with node id and finalizer
func UpdateZvolInfo(vol *apis.ZFSVolume, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveVolumeFinalizer removes finalizer from ZFSVolume CR
func RemoveVolumeFinalizer(vol *apis.ZFSVolume) error { _ = "STUB: not implemented"; return nil }

// GetZFSSnapshot fetches the given ZFSSnapshot
func GetZFSSnapshot(snapID string) (*apis.ZFSSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetZFSSnapshotStatus returns ZFSSnapshot status
func GetZFSSnapshotStatus(snapID string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetZFSSnapshotCapacity return capacity converted to int64
func GetZFSSnapshotCapacity(snap *apis.ZFSSnapshot) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// UpdateSnapInfo updates ZFSSnapshot CR with node id and finalizer
func UpdateSnapInfo(snap *apis.ZFSSnapshot) error { _ = "STUB: not implemented"; return nil }

// set the status to ready

// RemoveSnapFinalizer removes finalizer from ZFSSnapshot CR
func RemoveSnapFinalizer(snap *apis.ZFSSnapshot) error { _ = "STUB: not implemented"; return nil }

// RemoveBkpFinalizer removes finalizer from ZFSBackup CR
func RemoveBkpFinalizer(bkp *apis.ZFSBackup) error { _ = "STUB: not implemented"; return nil }

// UpdateBkpInfo updates the backup info with the status
func UpdateBkpInfo(bkp *apis.ZFSBackup, status apis.ZFSBackupStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// set the status

// UpdateRestoreInfo updates the rstr info with the status
func UpdateRestoreInfo(rstr *apis.ZFSRestore, status apis.ZFSRestoreStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// set the status

// GetUserFinalizers returns all the finalizers present on the ZFSVolume object
// except the one owned by ZFS node daemonset. We also need to ignore the foregroundDeletion
// finalizer as this will be present because of the foreground cascading deletion
func GetUserFinalizers(finalizers []string) []string { _ = "STUB: not implemented"; return nil }

// IsVolumeReady returns true if volume is Ready
func IsVolumeReady(vol *apis.ZFSVolume) bool {
	_ = "STUB: not implemented"
	// The status was added to ZFSVolume since v0.8.0
	return false
}

// For newer volumes created after v0.8.0, the status is sufficient to determine if the volume is ready
// If we check the finalizer to ensure the volume is Ready while the status is Pending or Failed
// the volume may never become Ready again when the controller provisions the volume again due to a timeout or controller crash

// For older volumes created before v0.8.0, there was no Status field
// so checking the node finalizer to make sure volume is Ready

// GetSnapshotForVolume fetches all the snapshots for the given volume
func GetSnapshotForVolume(volumeID string) (*apis.ZFSSnapshotList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarkForDeletion marks the volume for deletion by adding the annotation
func MarkForDeletion(volumeName string) error { _ = "STUB: not implemented"; return nil }

// IsVolumeEligibleForDeletion checks if the volume can be deleted or not
func IsVolumeEligibleForDeletion(volumeName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
