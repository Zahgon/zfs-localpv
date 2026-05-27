/*
Copyright 2020 The Kubernetes Authors.

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

// ResizeExtn can be used to run a resize command on the ext2/3/4 filesystem
// to expand the filesystem to the actual size of the device
func ResizeExtn(devpath string) error { _ = "STUB: not implemented"; return nil }

// ResizeXFS can be used to run a resize command on the xfs filesystem
// to expand the filesystem to the actual size of the device
func ResizeXFS(path string) error { _ = "STUB: not implemented"; return nil }

// handleVolResize resizes the filesystem, it is called after quota
// has been set on the volume. It takes care of expanding the filesystem.
func handleVolResize(vol *apis.ZFSVolume, volumePath string) error {
	_ = "STUB: not implemented"
	return nil
}

// just setting the quota is suffcient
// nothing to handle here
