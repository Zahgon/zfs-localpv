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

package pod

import (
	"github.com/openebs/zfs-localpv/tests/container"
	volume "github.com/openebs/zfs-localpv/tests/k8svolume"
	corev1 "k8s.io/api/core/v1"
)

const (
	// k8sNodeLabelKeyHostname is the label key used by Kubernetes
	// to store the hostname on the node resource.
	k8sNodeLabelKeyHostname = "kubernetes.io/hostname"
)

// Builder is the builder object for Pod
type Builder struct {
	pod  *Pod
	errs []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the Name field of Pod with provided value.
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithNamespace sets the Namespace field of Pod with provided value.
func (b *Builder) WithNamespace(namespace string) *Builder { _ = "STUB: not implemented"; return nil }

// WithContainerBuilder adds a container to this pod object.
//
// NOTE:
//
//	container details are present in the provided container
//
// builder object
func (b *Builder) WithContainerBuilder(
	containerBuilder *container.Builder,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithVolumeBuilder sets Volumes field of deployment.
func (b *Builder) WithVolumeBuilder(volumeBuilder *volume.Builder) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRestartPolicy sets the RestartPolicy field in Pod with provided arguments
func (b *Builder) WithRestartPolicy(
	restartPolicy corev1.RestartPolicy,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithNodeName sets the NodeName field of Pod with provided value.
func (b *Builder) WithNodeName(nodeName string) *Builder { _ = "STUB: not implemented"; return nil }

// WithNodeSelectorHostnameNew sets the Pod NodeSelector to the provided hostname value
// This function replaces (resets) the NodeSelector to use only hostname selector
func (b *Builder) WithNodeSelectorHostnameNew(hostname string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithContainers sets the Containers field in Pod with provided arguments
func (b *Builder) WithContainers(containers []corev1.Container) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithContainer sets the Containers field in Pod with provided arguments
func (b *Builder) WithContainer(container corev1.Container) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithVolumes sets the Volumes field in Pod with provided arguments
func (b *Builder) WithVolumes(volumes []corev1.Volume) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithVolume sets the Volumes field in Pod with provided arguments
func (b *Builder) WithVolume(volume corev1.Volume) *Builder { _ = "STUB: not implemented"; return nil }

// Build returns the Pod API instance
func (b *Builder) Build() (*corev1.Pod, error) { _ = "STUB: not implemented"; return nil, nil }
