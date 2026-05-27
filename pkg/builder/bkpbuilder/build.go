/*
Copyright 2020 The OpenEBS Authors

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

package bkpbuilder

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// Builder is the builder object for ZFSBackup
type Builder struct {
	bkp  *ZFSBackup
	errs []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// BuildFrom returns new instance of Builder
// from the provided api instance
func BuildFrom(bkp *apis.ZFSBackup) *Builder { _ = "STUB: not implemented"; return nil }

// WithNamespace sets the namespace of  ZFSBackup
func (b *Builder) WithNamespace(namespace string) *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the name of ZFSBackup
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithPrevSnap sets the previous snapshot for ZFSBackup
func (b *Builder) WithPrevSnap(snap string) *Builder { _ = "STUB: not implemented"; return nil }

// WithSnap sets the snapshot for ZFSBackup
func (b *Builder) WithSnap(snap string) *Builder { _ = "STUB: not implemented"; return nil }

// WithVolume sets the volume name of ZFSBackup
func (b *Builder) WithVolume(volume string) *Builder { _ = "STUB: not implemented"; return nil }

// WithNode sets the owenr node for the ZFSBackup
func (b *Builder) WithNode(node string) *Builder { _ = "STUB: not implemented"; return nil }

// WithStatus sets the status of the Backup progress
func (b *Builder) WithStatus(status apis.ZFSBackupStatus) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRemote sets the remote address for the ZFSBackup
func (b *Builder) WithRemote(server string) *Builder { _ = "STUB: not implemented"; return nil }

// WithLabels merges existing labels if any
// with the ones that are provided here
func (b *Builder) WithLabels(labels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithFinalizer merge existing finalizers if any
// with the ones that are provided here
func (b *Builder) WithFinalizer(finalizer []string) *Builder { _ = "STUB: not implemented"; return nil }

// Build returns ZFSBackup API object
func (b *Builder) Build() (*apis.ZFSBackup, error) { _ = "STUB: not implemented"; return nil, nil }
