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

package restorebuilder

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// Builder is the builder object for ZFSRestore
type Builder struct {
	rstr *ZFSRestore
	errs []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// BuildFrom returns new instance of Builder
// from the provided api instance
func BuildFrom(rstr *apis.ZFSRestore) *Builder { _ = "STUB: not implemented"; return nil }

// WithNamespace sets the namespace of  ZFSRestore
func (b *Builder) WithNamespace(namespace string) *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the name of ZFSRestore
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithVolume sets the name of ZFSRestore
func (b *Builder) WithVolume(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithVolSpec copies volume spec to ZFSRestore Object
func (b *Builder) WithVolSpec(vspec apis.VolumeInfo) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithNode sets the node id for ZFSRestore
func (b *Builder) WithNode(node string) *Builder { _ = "STUB: not implemented"; return nil }

// WithStatus sets the status for ZFSRestore
func (b *Builder) WithStatus(status apis.ZFSRestoreStatus) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRemote sets the node id for ZFSRestore
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

// Build returns ZFSRestore API object
func (b *Builder) Build() (*apis.ZFSRestore, error) { _ = "STUB: not implemented"; return nil, nil }
