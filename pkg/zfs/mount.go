/*
Copyright © 2020 The OpenEBS Authors
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
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// MountInfo contains the volume related info
// for all types of volumes in ZFSVolume
type MountInfo struct {
	// FSType of a volume will specify the
	// format type - ext4(default), xfs of PV
	FSType string `json:"fsType"`

	// AccessMode of a volume will hold the
	// access mode of the volume
	AccessModes []string `json:"accessModes"`

	// MountPath of the volume will hold the
	// path on which the volume is mounted
	// on that node
	MountPath string `json:"mountPath"`

	// MountOptions specifies the options with
	// which mount needs to be attempted
	MountOptions []string `json:"mountOptions"`
}

// FormatAndMountZvol formats and mounts the created volume to the desired mount path
func FormatAndMountZvol(devicePath string, mountInfo *MountInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// UmountVolume unmounts the volume and the corresponding mount path is removed
func UmountVolume(vol *apis.ZFSVolume, targetPath string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// device has already been un-mounted, return successful

// If the mount is broken, unmount it since there's nothing else to do anyway

// ignoring the failure as the volume has already
// been umounted, now the new pod can mount it

func verifyMountRequest(vol *apis.ZFSVolume, mountpath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

/*
 * This check is the famous *Wall Of North*
 * It will not let the volume to be mounted
 * at more than two places. The volume should
 * be unmounted before proceeding to the mount
 * operation.
 */

// if device is already mounted at the mount point, return successful

// if it is not a shared volume, then it should not mounted to more than one path

// MountZvol mounts the disk to the specified path
func MountZvol(vol *apis.ZFSVolume, mount *MountInfo) error { _ = "STUB: not implemented"; return nil }

// MountDataset mounts the zfs dataset to the specified path
func MountDataset(vol *apis.ZFSVolume, mount *MountInfo) error {
	_ = "STUB: not implemented"
	return nil
}

/*
 * We might have created volumes and then upgraded the node agent before
 * getting the mount request for that volume. In this case volume will
 * not be created with mountpoint as legacy. Handling the mount in old way.
 */

// MountFilesystem mounts the disk to the specified path
func MountFilesystem(vol *apis.ZFSVolume, mount *MountInfo) error {
	_ = "STUB: not implemented"
	// creating the directory with 0750 permission so that it can be accessed by other person.
	// if the directory already exist(old k8s), the creator should set the proper permission.
	return nil
}

// the mount may be broken, but present, so proceed to ensure idempotency

// MountBlock mounts the block disk to the specified path
func MountBlock(vol *apis.ZFSVolume, mountinfo *MountInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Create the mount point as a file since bind mount device node requires it to be a file

// do the bind mount of the zvol device at the target path

func makeFile(pathname string) error { _ = "STUB: not implemented"; return nil }

func isBrokenMount(mountPath string) bool { _ = "STUB: not implemented"; return false }

// In case the mount point becomes broken, ex: filesystem shutdown
