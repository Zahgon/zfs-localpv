/*
Copyright 2017 The Kubernetes Authors.

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

package zfs

import (
	"os/exec"

	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// zfs related constants
const (
	ZFSDevPath = "/dev/zvol/"
	FSTypeZFS  = "zfs"
)

// zfs command related constants
const (
	ZFSVolCmd      = "zfs"
	ZFSCreateArg   = "create"
	ZFSCloneArg    = "clone"
	ZFSDestroyArg  = "destroy"
	ZFSSetArg      = "set"
	ZFSGetArg      = "get"
	ZFSListArg     = "list"
	ZFSSnapshotArg = "snapshot"
	ZFSSendArg     = "send"
	ZFSRecvArg     = "recv"
)

// constants to define volume type
const (
	VolTypeDataset = "DATASET"
	VolTypeZVol    = "ZVOL"
)

// runCmd executes cmd, logs the invocation at V(4) and full output at V(5),
// and returns combined stdout+stderr. Use this instead of cmd.CombinedOutput()
// throughout this package so verbosity can be tuned with --v.
func runCmd(cmd *exec.Cmd, dataset string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PropertyChanged return whether volume property is changed
func PropertyChanged(oldVol *apis.ZFSVolume, newVol *apis.ZFSVolume) bool {
	_ = "STUB: not implemented"
	return false
}

// GetVolumeType returns the volume type
// whether it is a zvol or dataset
func GetVolumeType(fstype string) string {
	_ = "STUB: not implemented"
	/*
	 * if fstype is provided as zfs then a zfs dataset will be created
	 * otherwise a zvol will be created
	 */return ""
}

// buildZvolCreateArgs returns zfs create command for zvol along with attributes as a string array
func buildZvolCreateArgs(vol *apis.ZFSVolume) []string { _ = "STUB: not implemented"; return nil }

// buildCloneCreateArgs returns zfs clone commands for zfs volume/dataset along with attributes as a string array
func buildCloneCreateArgs(vol *apis.ZFSVolume) []string { _ = "STUB: not implemented"; return nil }

// Note: Encryption parameters (encryption, keylocation, keyformat) are NOT set when cloning.
// ZFS clones automatically inherit encryption settings from the parent snapshot.
// The encryption property is read-only on clones and attempting to set it will fail with
// "encryption is readonly" error. This is expected ZFS behavior.

// buildZFSSnapCreateArgs returns zfs create command for zfs snapshot
// zfs snapshot <poolname>/<volname>@<snapname>
func buildZFSSnapCreateArgs(snap *apis.ZFSSnapshot) []string { _ = "STUB: not implemented"; return nil }

// buildZFSSnapDestroyArgs returns zfs destroy command for zfs snapshot
// zfs destroy <poolname>/<volname>@<snapname>
func buildZFSSnapDestroyArgs(snap *apis.ZFSSnapshot) []string {
	_ = "STUB: not implemented"
	return nil
}

// buildDatasetCreateArgs returns zfs create command for dataset along with attributes as a string array
func buildDatasetCreateArgs(vol *apis.ZFSVolume) []string { _ = "STUB: not implemented"; return nil }

// set the mount path to none, by default zfs mounts it to the default dataset path

// buildVolumeSetArgs returns volume set command along with attributes as a string array
// TODO(pawan) need to find a way to identify which property has changed
func buildVolumeSetArgs(vol *apis.ZFSVolume) []string { _ = "STUB: not implemented"; return nil }

// buildVolumeResizeArgs returns volume set command for resizing the zfs volume
func buildVolumeResizeArgs(vol *apis.ZFSVolume) []string { _ = "STUB: not implemented"; return nil }

// buildVolumeBackupArgs returns volume send command for sending the zfs volume
func buildVolumeBackupArgs(bkp *apis.ZFSBackup, vol *apis.ZFSVolume) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// do incremental send

// buildVolumeRestoreArgs returns volume recv command for receiving the zfs volume
func buildVolumeRestoreArgs(rstr *apis.ZFSRestore) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildVolumeDestroyArgs returns volume destroy command along with attributes as a string array
func buildVolumeDestroyArgs(vol *apis.ZFSVolume) []string { _ = "STUB: not implemented"; return nil }

func getVolume(volume string) error { _ = "STUB: not implemented"; return nil }

// CreateVolume creates the zvol/dataset as per
// info provided in ZFSVolume object
func CreateVolume(vol *apis.ZFSVolume) error { _ = "STUB: not implemented"; return nil }

// CreateClone creates clone for the zvol/dataset as per
// info provided in ZFSVolume object
func CreateClone(vol *apis.ZFSVolume) error { _ = "STUB: not implemented"; return nil }

// datasource is volume, create the snapshot first

// use volname as snapname

// add src vol name

// SetDatasetMountProp sets mountpoint for the volume
func SetDatasetMountProp(volume string, mountpath string) error {
	_ = "STUB: not implemented"
	return nil
}

// MountZFSDataset mounts the dataset to the given mountpoint
func MountZFSDataset(vol *apis.ZFSVolume, mountpath string) error {
	_ = "STUB: not implemented"
	return nil
}

// set the mountpoint to the path where this volume should be mounted

/*
 * see if we should attempt to mount the dataset.
 * Setting the mountpoint is sufficient to mount the zfs dataset,
 * but if dataset has been unmounted, then setting the mountpoint
 * is not sufficient, we have to mount the dataset explicitly
 */

// SetDatasetLegacyMount sets the dataset mountpoint to legacy if not set
func SetDatasetLegacyMount(vol *apis.ZFSVolume) error { _ = "STUB: not implemented"; return nil }

// set the mountpoint to legacy

// GetVolumeProperty gets zfs properties for the volume
func GetVolumeProperty(vol *apis.ZFSVolume, prop string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetVolumeProp sets the volume property
func SetVolumeProp(vol *apis.ZFSVolume) error { _ = "STUB: not implemented"; return nil }

//nothing to set, just return

/* Case: Restart =>
 * In this case we get the add event but here we don't know which
 * property has changed when we were down, so firing the zfs set
 * command with the all property present on the ZFSVolume.

 * Case: Property Change =>
 * TODO(pawan) When we get the update event, we make sure at least
 * one property has changed before adding it to the event queue for
 * handling. At this stage, since we haven't stored the
 * ZFSVolume object as it will be too heavy, we are firing the set
 * command with the all property preset in the ZFSVolume object since
 * it is guaranteed that at least one property has changed.
 */

// DestroyVolume deletes the zfs volume
func DestroyVolume(vol *apis.ZFSVolume) error { _ = "STUB: not implemented"; return nil }

// check if parent dataset is present or not before attempting to delete the volume

// datasource is volume, delete the dependent snapshot

// snapname is same as volname

// add src vol name

// no need to reconcile as volume has already been deleted

// CreateSnapshot creates the zfs volume snapshot
func CreateSnapshot(snap *apis.ZFSSnapshot) error { _ = "STUB: not implemented"; return nil }

// snapshot already there just return

// DestroySnapshot deletes the zfs volume snapshot
func DestroySnapshot(snap *apis.ZFSSnapshot) error { _ = "STUB: not implemented"; return nil }

// check if parent dataset is present or not before attempting to delete the snapshot

// GetVolumeDevPath returns devpath for the given volume
func GetVolumeDevPath(vol *apis.ZFSVolume) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// evaluate the symlink to get the dev path for zvol

// ResizeZFSVolume resize volume
func ResizeZFSVolume(vol *apis.ZFSVolume, mountpath string, resizefs bool) error {
	_ = "STUB: not implemented"
	return nil
}

// resize the filesystem so that applications can use the expanded space

// CreateBackup creates the backup
func CreateBackup(bkp *apis.ZFSBackup) error { _ = "STUB: not implemented"; return nil }

/* create the snapshot for the backup */

// DestoryBackup deletes the snapshot created
func DestoryBackup(bkp *apis.ZFSBackup) error { _ = "STUB: not implemented"; return nil }

// Volume has been deleted, return

/* create the snapshot for the backup */

// getDevice waits for the device to be created and returns the devpath
func getDevice(volume string) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// device should be created within 5 seconds
		nil
}

// CreateRestore creates the restore
func CreateRestore(rstr *apis.ZFSRestore) error { _ = "STUB: not implemented"; return nil }

// for backward compatibility, older version of
// velero will not add spec in the ZFSRestore Object
// query it here and fill that information

/*
 * need to generate a new uuid for zfs and btrfs volumes
 * so that we can mount it.
 */

// ListZFSPool invokes `zfs list` to list all the available
// pools in the node.
func ListZFSPool() ([]apis.Pool, error) { _ = "STUB: not implemented"; return nil, nil }

// The `zfs list` command will list down all the resources including
// pools and volumes and as the pool names cannot have "/" in the name
// the function below filters out the pools. Sample output of command:
// $ zfs list -s name -o name,guid,available -H -p
// zfspv-pool	4734063099997348493	103498467328
// zfspv-pool/pvc-be02d230-3738-4de9-8968-70f5d10d86dd	3380225606535803752	4294942720
func decodeListOutput(raw []byte) ([]apis.Pool, error) { _ = "STUB: not implemented"; return nil, nil }

// reservationProperty returns the reservation property based on the quota type.
func reservationProperty(quotaType, capacity string) string { _ = "STUB: not implemented"; return "" }

// reservationPropertyName returns the reservation property name based on the quota type.
func reservationPropertyName(quotaType string) string { _ = "STUB: not implemented"; return "" }

// Return the mapped property or default to "reservation"

// quotaProperty ensures backwards compatibility for the quota property.
func quotaProperty(quotaType string) string { _ = "STUB: not implemented"; return "" }
