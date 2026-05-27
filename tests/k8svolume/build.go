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

package k8svolume

import (
	corev1 "k8s.io/api/core/v1"
)

// Builder is the builder object for Volume
type Builder struct {
	volume *Volume
	errs   []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the Name field of Volume with provided value.
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithHostDirectory sets the VolumeSource field of Volume with provided hostpath
// as type directory.
func (b *Builder) WithHostDirectory(path string) *Builder { _ = "STUB: not implemented"; return nil }

// WithHostPathAndType sets the VolumeSource field of Volume with provided
// hostpath as directory path and type as directory type
func (b *Builder) WithHostPathAndType(
	dirpath string,
	dirtype *corev1.HostPathType,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithPVCSource sets the Volume field of Volume with provided pvc
func (b *Builder) WithPVCSource(pvcName string) *Builder { _ = "STUB: not implemented"; return nil }

// WithEmptyDir sets the EmptyDir field of the Volume with provided dir
func (b *Builder) WithEmptyDir(dir *corev1.EmptyDirVolumeSource) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Build returns the Volume API instance
func (b *Builder) Build() (*corev1.Volume, error) { _ = "STUB: not implemented"; return nil, nil }
