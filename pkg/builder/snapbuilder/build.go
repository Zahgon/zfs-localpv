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

package snapbuilder

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// Builder is the builder object for ZFSSnapshot
type Builder struct {
	snap *ZFSSnapshot
	errs []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// BuildFrom returns new instance of Builder
// from the provided api instance
func BuildFrom(snap *apis.ZFSSnapshot) *Builder { _ = "STUB: not implemented"; return nil }

// WithNamespace sets the namespace of  ZFSSnapshot
func (b *Builder) WithNamespace(namespace string) *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the name of ZFSSnapshot
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithLabels merges existing labels if any
// with the ones that are provided here
func (b *Builder) WithLabels(labels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithFinalizer merge existing finalizers if any
// with the ones that are provided here
func (b *Builder) WithFinalizer(finalizer []string) *Builder { _ = "STUB: not implemented"; return nil }

// Build returns ZFSSnapshot API object
func (b *Builder) Build() (*apis.ZFSSnapshot, error) { _ = "STUB: not implemented"; return nil, nil }
