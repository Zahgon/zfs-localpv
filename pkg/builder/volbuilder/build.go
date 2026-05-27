/*
Copyright 2019 The OpenEBS Authors

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

package volbuilder

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// MarkForDeletionAnnotation is the annotation key
const MarkForDeletionAnnotation string = "openebs.io/marked-for-deletion"

// Builder is the builder object for ZFSVolume
type Builder struct {
	volume *ZFSVolume
	errs   []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// BuildFrom returns new instance of Builder
// from the provided api instance
func BuildFrom(volume *apis.ZFSVolume) *Builder { _ = "STUB: not implemented"; return nil }

// WithNamespace sets the namespace of  ZFSVolume
func (b *Builder) WithNamespace(namespace string) *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the name of ZFSVolume
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithCapacity sets the Capacity of zfs volume by converting string
// capacity into Quantity
func (b *Builder) WithCapacity(capacity string) *Builder { _ = "STUB: not implemented"; return nil }

// WithEncryption sets the encryption on ZFSVolume
func (b *Builder) WithEncryption(encr string) *Builder { _ = "STUB: not implemented"; return nil }

// WithKeyLocation sets the encryption key location on ZFSVolume
func (b *Builder) WithKeyLocation(kl string) *Builder { _ = "STUB: not implemented"; return nil }

// WithKeyFormat sets the encryption key format on ZFSVolume
func (b *Builder) WithKeyFormat(kf string) *Builder { _ = "STUB: not implemented"; return nil }

// WithCompression sets compression of ZFSVolume
func (b *Builder) WithCompression(compression string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithDedup sets dedup property of ZFSVolume
func (b *Builder) WithDedup(dedup string) *Builder { _ = "STUB: not implemented"; return nil }

// WithThinProv sets if ZFSVolume needs to be thin provisioned
func (b *Builder) WithThinProv(thinprov string) *Builder { _ = "STUB: not implemented"; return nil }

// WithOwnerNodeID sets owner nodeid for the ZFSVolume where the volume should be provisioned
func (b *Builder) WithOwnerNodeID(nodeid string) *Builder { _ = "STUB: not implemented"; return nil }

// WithRecordSize sets the recordsize of ZFSVolume
func (b *Builder) WithRecordSize(rs string) *Builder { _ = "STUB: not implemented"; return nil }

// WithVolBlockSize sets the volblocksize of ZFSVolume
func (b *Builder) WithVolBlockSize(bs string) *Builder { _ = "STUB: not implemented"; return nil }

// WithAnnotation sets the annotation of ZFSVolume
func (b *Builder) WithAnnotation() *Builder { _ = "STUB: not implemented"; return nil }

// WithVolumeType sets if ZFSVolume needs to be thin provisioned
func (b *Builder) WithVolumeType(vtype string) *Builder { _ = "STUB: not implemented"; return nil }

// WithVolumeStatus sets ZFSVolume status
func (b *Builder) WithVolumeStatus(status string) *Builder { _ = "STUB: not implemented"; return nil }

// WithFsType sets filesystem for the ZFSVolume
func (b *Builder) WithFsType(fstype string) *Builder { _ = "STUB: not implemented"; return nil }

// WithQuotaType sets quota type for dataset volume
func (b *Builder) WithQuotaType(quotatype string) *Builder { _ = "STUB: not implemented"; return nil }

// WithShared sets where filesystem is shared or not
func (b *Builder) WithShared(shared string) *Builder { _ = "STUB: not implemented"; return nil }

// WithSnapshot sets Snapshot name for creating clone volume
func (b *Builder) WithSnapshot(snap string) *Builder { _ = "STUB: not implemented"; return nil }

// WithPoolName sets Pool name for creating volume
func (b *Builder) WithPoolName(pool string) *Builder { _ = "STUB: not implemented"; return nil }

// WithNodeName sets NodeID for creating the volume
func (b *Builder) WithNodeName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithLabels merges existing labels if any
// with the ones that are provided here
func (b *Builder) WithLabels(labels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithFinalizer sets Finalizer name creating the volume
func (b *Builder) WithFinalizer(finalizer []string) *Builder { _ = "STUB: not implemented"; return nil }

// Build returns ZFSVolume API object
func (b *Builder) Build() (*apis.ZFSVolume, error) { _ = "STUB: not implemented"; return nil, nil }
