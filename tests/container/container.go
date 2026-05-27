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

package container

import (
	"github.com/openebs/lib-csi/pkg/common/errors"
	corev1 "k8s.io/api/core/v1"
)

type container struct {
	corev1.Container // kubernetes container type
}

// OptionFunc is a typed function that abstracts anykind of operation
// against the provided container instance
//
// This is the basic building block to create functional operations
// against the container instance
type OptionFunc func(*container)

// Predicate abstracts conditional logic w.r.t the container instance
//
// NOTE:
// Predicate is a functional approach versus traditional approach to mix
// conditions such as *if-else* within blocks of business logic
//
// NOTE:
// Predicate approach enables clear separation of conditionals from
// imperatives i.e. actions that form the business logic
type Predicate func(*container) (nameOrMsg string, ok bool)

// predicateFailedError returns the provided predicate as an error
func predicateFailedError(message string) error { _ = "STUB: not implemented"; return nil }

var (
	errorvalidationFailed = errors.New("container validation failed")
)

// asContainer transforms this container instance into corresponding kubernetes
// container type
func (c *container) asContainer() corev1.Container {
	_ = "STUB: not implemented"
	return *new(corev1.Container)
}

// New returns a new kubernetes container
func New(opts ...OptionFunc) corev1.Container {
	_ = "STUB: not implemented"
	return *new(corev1.Container)
}

// Builder provides utilities required to build a kubernetes container type
type Builder struct {
	con    *container  // container instance
	checks []Predicate // validations to be done while building the container instance
	errors []error     // errors found while building the container instance
}

// NewBuilder returns a new instance of builder
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// validate will run checks against container instance
func (b *Builder) validate() error { _ = "STUB: not implemented"; return nil }

// Build returns the final kubernetes container
func (b *Builder) Build() (corev1.Container, error) {
	_ = "STUB: not implemented"
	return *new(corev1.Container), nil
}

// AddCheck adds the predicate as a condition to be validated against the
// container instance
func (b *Builder) AddCheck(p Predicate) *Builder { _ = "STUB: not implemented"; return nil }

// AddChecks adds the provided predicates as conditions to be validated against
// the container instance
func (b *Builder) AddChecks(p []Predicate) *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the name of the container
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the name of the container
func WithName(name string) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithImage sets the image of the container
func (b *Builder) WithImage(img string) *Builder { _ = "STUB: not implemented"; return nil }

// WithImage sets the image of the container
func WithImage(img string) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithCommandNew sets the command of the container
func (b *Builder) WithCommandNew(cmd []string) *Builder { _ = "STUB: not implemented"; return nil }

// WithArgumentsNew sets the command arguments of the container
func (b *Builder) WithArgumentsNew(args []string) *Builder { _ = "STUB: not implemented"; return nil }

// WithVolumeMountsNew sets the command arguments of the container
func (b *Builder) WithVolumeMountsNew(volumeMounts []corev1.VolumeMount) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithVolumeDevicesNew sets the command arguments of the container
func (b *Builder) WithVolumeDevicesNew(volumeDevices []corev1.VolumeDevice) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithImagePullPolicy sets the image pull policy of the container
func (b *Builder) WithImagePullPolicy(policy corev1.PullPolicy) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithPrivilegedSecurityContext sets securitycontext of the container
func (b *Builder) WithPrivilegedSecurityContext(privileged *bool) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithResources sets resources of the container
func (b *Builder) WithResources(
	resources *corev1.ResourceRequirements,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithResourcesByValue sets resources of the container
func (b *Builder) WithResourcesByValue(resources corev1.ResourceRequirements) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithPortsNew sets ports of the container
func (b *Builder) WithPortsNew(ports []corev1.ContainerPort) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithEnvsNew sets the envs of the container
func (b *Builder) WithEnvsNew(envs []corev1.EnvVar) *Builder { _ = "STUB: not implemented"; return nil }

// WithEnvs sets the envs of the container
func (b *Builder) WithEnvs(envs []corev1.EnvVar) *Builder { _ = "STUB: not implemented"; return nil }

// WithLivenessProbe sets the liveness probe of the container
func (b *Builder) WithLivenessProbe(liveness *corev1.Probe) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithLifeCycle sets the life cycle of the container
func (b *Builder) WithLifeCycle(lc *corev1.Lifecycle) *Builder {
	_ = "STUB: not implemented"
	return nil
}
