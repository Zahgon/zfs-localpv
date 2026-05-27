// Copyright 2019 The OpenEBS Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pts

import (
	"github.com/openebs/zfs-localpv/tests/container"
	volume "github.com/openebs/zfs-localpv/tests/k8svolume"
	corev1 "k8s.io/api/core/v1"
)

// PodTemplateSpec holds the api's podtemplatespec objects
type PodTemplateSpec struct {
	Object *corev1.PodTemplateSpec
}

// Builder is the builder object for Pod
type Builder struct {
	podtemplatespec *PodTemplateSpec
	errs            []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the Name field of podtemplatespec with provided value.
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithNamespace sets the Namespace field of PodTemplateSpec with provided value.
func (b *Builder) WithNamespace(namespace string) *Builder { _ = "STUB: not implemented"; return nil }

// WithAnnotations merges existing annotations if any
// with the ones that are provided here
func (b *Builder) WithAnnotations(annotations map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithAnnotationsNew resets the annotation field of podtemplatespec
// with provided arguments
func (b *Builder) WithAnnotationsNew(annotations map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original map

// override

// WithLabels merges existing labels if any
// with the ones that are provided here
func (b *Builder) WithLabels(labels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithLabelsNew resets the labels field of podtemplatespec
// with provided arguments
func (b *Builder) WithLabelsNew(labels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original map

// override

// WithNodeSelector merges the nodeselectors if present
// with the provided arguments
func (b *Builder) WithNodeSelector(nodeselectors map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithNodeSelectorNew resets the nodeselector field of podtemplatespec
// with provided arguments
func (b *Builder) WithNodeSelectorNew(nodeselectors map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original map

// override

// WithServiceAccountName sets the ServiceAccountnNme field of podtemplatespec
func (b *Builder) WithServiceAccountName(serviceAccountnNme string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithAffinity sets the affinity field of podtemplatespec
func (b *Builder) WithAffinity(affinity *corev1.Affinity) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original pointer

// WithTolerations merges the existing tolerations
// with the provided arguments
func (b *Builder) WithTolerations(tolerations ...corev1.Toleration) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithTolerationsNew sets the tolerations field of podtemplatespec
func (b *Builder) WithTolerationsNew(tolerations ...corev1.Toleration) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original slice

// WithContainerBuilders builds the list of containerbuilder
// provided and merges it to the containers field of the podtemplatespec
func (b *Builder) WithContainerBuilders(
	containerBuilderList ...*container.Builder,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithTerminationGracePeriodSeconds adds the terminationGracePeriodSeconds
func (b *Builder) WithTerminationGracePeriodSeconds(
	period int64,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithVolumeBuilders builds the list of volumebuilders provided
// and merges it to the volumes field of podtemplatespec.
func (b *Builder) WithVolumeBuilders(
	volumeBuilderList ...*volume.Builder,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithContainerBuildersNew builds the list of containerbuilder
// provided and sets the containers field of the podtemplatespec
func (b *Builder) WithContainerBuildersNew(
	containerBuilderList ...*container.Builder,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithVolumeBuildersNew builds the list of volumebuilders provided
// and sets Volumes field of podtemplatespec.
func (b *Builder) WithVolumeBuildersNew(
	volumeBuilderList ...*volume.Builder,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Build returns a deployment instance
func (b *Builder) Build() (*PodTemplateSpec, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Builder) validate() error { _ = "STUB: not implemented"; return nil }
