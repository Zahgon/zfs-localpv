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

package sc

import (
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

// Builder enables building an instance of StorageClass
type Builder struct {
	sc   *StorageClass
	errs []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the Name field of storageclass with provided argument.
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithGenerateName appends a random string after the name
func (b *Builder) WithGenerateName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithAnnotations sets the Annotations field of storageclass with provided value.
func (b *Builder) WithAnnotations(annotations map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithParametersNew resets existing parameters if any with
// ones that are provided here
func (b *Builder) WithParametersNew(parameters map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original map

// override

// WithProvisioner sets the Provisioner field of storageclass with provided argument.
func (b *Builder) WithProvisioner(provisioner string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithReclaimPolicy sets the ReclaimPolicy field of storageclass with provided argument.
func (b *Builder) WithReclaimPolicy(reclaimPolicy *corev1.PersistentVolumeReclaimPolicy) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithVolumeExpansion sets the AllowedVolumeExpansion field of storageclass with provided argument.
func (b *Builder) WithVolumeExpansion(expansionAllowed bool) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Build returns the StorageClass API instance
func (b *Builder) Build() (*storagev1.StorageClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithVolumeBindingMode sets the volume binding mode of storageclass with
// provided argument.
func (b *Builder) WithVolumeBindingMode(bindingMode storagev1.VolumeBindingMode) *Builder {
	_ = "STUB: not implemented"
	return nil
}
