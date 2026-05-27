package pv

import (
	corev1 "k8s.io/api/core/v1"
)

// IsAvailable returns true if the pv is bounded
func (p *PV) IsAvailable() bool { _ = "STUB: not implemented"; return false }

// NewForAPIObject returns a new instance of PV
func NewForAPIObject(obj *corev1.PersistentVolume, opts ...pvBuildOption) *PV {
	_ = "STUB: not implemented"
	return nil
}
