/*
Copyright 2021 The OpenEBS Authors

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

package nodebuilder

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Builder is the builder object for ZFSNode
type Builder struct {
	node *ZFSNode
	errs []error
}

// ZFSNode is a wrapper over
// ZFSNode API instance
type ZFSNode struct {
	// ZFSVolume object
	Object *apis.ZFSNode
}

// From returns a new instance of
// zfs volume
func From(node *apis.ZFSNode) *ZFSNode { _ = "STUB: not implemented"; return nil }

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// BuildFrom returns new instance of Builder
// from the provided api instance
func BuildFrom(node *apis.ZFSNode) *Builder { _ = "STUB: not implemented"; return nil }

// WithNamespace sets the namespace of ZFSNode
func (b *Builder) WithNamespace(namespace string) *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the name of ZFSNode
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithPools sets the pools of ZFSNode
func (b *Builder) WithPools(pools []apis.Pool) *Builder { _ = "STUB: not implemented"; return nil }

// WithOwnerReferences sets the owner references of ZFSNode
func (b *Builder) WithOwnerReferences(ownerRefs ...metav1.OwnerReference) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Build returns ZFSNode API object
func (b *Builder) Build() (*apis.ZFSNode, error) { _ = "STUB: not implemented"; return nil, nil }
