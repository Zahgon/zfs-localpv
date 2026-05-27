package pv

import (
	corev1 "k8s.io/api/core/v1"
)

// PV is a wrapper over persistentvolume api
type pvBuildOption func(*PV)

// PV is a wrapper over persistentvolume api
// object. It provides build, validations and other common
// logic to be used by various feature specific callers.
type PV struct {
	object *corev1.PersistentVolume
}

// Builder is the builder object for PV
type Builder struct {
	pv   *PV
	errs []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the Name field of PV with provided value.
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithStorageClass sets the StorageClass field of PV with provided arguments
func (b *Builder) WithStorageClass(scName string) *Builder { _ = "STUB: not implemented"; return nil }

// WithCapacity sets the Capacity field in PV with provided arguments
func (b *Builder) WithCapacity(capacity string) *Builder { _ = "STUB: not implemented"; return nil }

// WithAccessModes sets the AccessMode field in PV with provided arguments
func (b *Builder) WithAccessModes(accessMode []corev1.PersistentVolumeAccessMode) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Build returns the PV API instance
func (b *Builder) Build() (*corev1.PersistentVolume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
