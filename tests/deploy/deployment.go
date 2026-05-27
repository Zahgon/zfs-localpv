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

package deploy

import (
	templatespec "github.com/openebs/zfs-localpv/tests/pts"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Predicate abstracts conditional logic w.r.t the deployment instance
//
// NOTE:
// predicate is a functional approach versus traditional approach to mix
// conditions such as *if-else* within blocks of business logic
//
// NOTE:
// predicate approach enables clear separation of conditionals from
// imperatives i.e. actions that form the business logic
type Predicate func(*Deploy) bool

// Deploy is the wrapper over k8s deployment Object
type Deploy struct {
	// kubernetes deployment instance
	object *appsv1.Deployment
}

// Builder enables building an instance of
// deployment
type Builder struct {
	deployment *Deploy     // kubernetes deployment instance
	checks     []Predicate // predicate list for deploy
	errors     []error
}

// PredicateName type is wrapper over string.
// It is used to refer predicate and status msg.
type PredicateName string

const (
	// PredicateProgressDeadlineExceeded refer to
	// predicate IsProgressDeadlineExceeded.
	PredicateProgressDeadlineExceeded PredicateName = "ProgressDeadlineExceeded"
	// PredicateNotSpecSynced refer to predicate IsNotSpecSynced
	PredicateNotSpecSynced PredicateName = "NotSpecSynced"
	// PredicateOlderReplicaActive refer to predicate IsOlderReplicaActive
	PredicateOlderReplicaActive PredicateName = "OlderReplicaActive"
	// PredicateTerminationInProgress refer to predicate IsTerminationInProgress
	PredicateTerminationInProgress PredicateName = "TerminationInProgress"
	// PredicateUpdateInProgress refer to predicate IsUpdateInProgress.
	PredicateUpdateInProgress PredicateName = "UpdateInProgress"
)

// String implements the stringer interface
func (d *Deploy) String() string { _ = "STUB: not implemented"; return "" }

// GoString implements the goStringer interface
func (d *Deploy) GoString() string {
	_ = "STUB: not implemented"

	// NewBuilder returns a new instance of builder meant for deployment
	return ""
}

func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// WithName sets the Name field of deployment with provided value.
func (b *Builder) WithName(name string) *Builder { _ = "STUB: not implemented"; return nil }

// WithNamespace sets the Namespace field of deployment with provided value.
func (b *Builder) WithNamespace(namespace string) *Builder { _ = "STUB: not implemented"; return nil }

// WithAnnotations merges existing annotations if any
// with the ones that are provided here
func (b *Builder) WithAnnotations(annotations map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithAnnotationsNew resets existing annotaions if any with
// ones that are provided here
func (b *Builder) WithAnnotationsNew(annotations map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original map

// override

// WithNodeSelector Sets the node selector with the provided argument.
func (b *Builder) WithNodeSelector(selector map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithNodeSelectorNew Sets the node selector with the provided argument.
func (b *Builder) WithNodeSelectorNew(selector map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithOwnerReferenceNew sets ownerreference if any with
// ones that are provided here
func (b *Builder) WithOwnerReferenceNew(ownerRefernce []metav1.OwnerReference) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithLabels merges existing labels if any
// with the ones that are provided here
func (b *Builder) WithLabels(labels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithLabelsNew resets existing labels if any with
// ones that are provided here
func (b *Builder) WithLabelsNew(labels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original map

// override

// WithSelectorMatchLabels merges existing matchlabels if any
// with the ones that are provided here
func (b *Builder) WithSelectorMatchLabels(matchlabels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithSelectorMatchLabelsNew resets existing matchlabels if any with
// ones that are provided here
func (b *Builder) WithSelectorMatchLabelsNew(matchlabels map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// copy of original map

// override

// WithReplicas sets the replica field of deployment
func (b *Builder) WithReplicas(replicas *int32) *Builder { _ = "STUB: not implemented"; return nil }

// WithStrategyType sets the strategy field of the deployment
func (b *Builder) WithStrategyType(
	strategytype appsv1.DeploymentStrategyType,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithPodTemplateSpecBuilder sets the template field of the deployment
func (b *Builder) WithPodTemplateSpecBuilder(
	tmplbuilder *templatespec.Builder,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

type deployBuildOption func(*Deploy)

// NewForAPIObject returns a new instance of builder
// for a given deployment Object
func NewForAPIObject(
	obj *appsv1.Deployment,
	opts ...deployBuildOption,
) *Deploy {
	_ = "STUB: not implemented"
	return nil
}

// Build returns a deployment instance
func (b *Builder) Build() (*appsv1.Deployment, error) {
	_ = "STUB: not implemented"

	// TODO: err in Wrapf is not logged. Fix is required
	return nil, nil
}

func (b *Builder) validate() error { _ = "STUB: not implemented"; return nil }

// IsRollout range over rolloutChecks map and check status of each predicate
// also it generates status message from rolloutStatuses using predicate key
func (d *Deploy) IsRollout() (PredicateName, bool) {
	_ = "STUB: not implemented"
	return *new(PredicateName), false
}

// FailedRollout returns rollout status message for fail condition
func (d *Deploy) FailedRollout(name PredicateName) *RolloutOutput {
	_ = "STUB: not implemented"
	return nil
}

// SuccessRollout returns rollout status message for success condition
func (d *Deploy) SuccessRollout() *RolloutOutput { _ = "STUB: not implemented"; return nil }

// RolloutStatus returns rollout message of deployment instance
func (d *Deploy) RolloutStatus() (op *RolloutOutput, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RolloutStatusRaw returns rollout message of deployment instance
// in byte format
func (d *Deploy) RolloutStatusRaw() (op []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddCheck adds the predicate as a condition to be validated
// against the deployment instance
func (b *Builder) AddCheck(p Predicate) *Builder { _ = "STUB: not implemented"; return nil }

// AddChecks adds the provided predicates as conditions to be
// validated against the deployment instance
func (b *Builder) AddChecks(p []Predicate) *Builder { _ = "STUB: not implemented"; return nil }

// IsProgressDeadlineExceeded is used to check update is timed out or not.
// If `Progressing` condition's reason is `ProgressDeadlineExceeded` then
// it is not rolled out.
func IsProgressDeadlineExceeded() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// IsProgressDeadlineExceeded is used to check update is timed out or not.
// If `Progressing` condition's reason is `ProgressDeadlineExceeded` then
// it is not rolled out.
func (d *Deploy) IsProgressDeadlineExceeded() bool { _ = "STUB: not implemented"; return false }

// IsOlderReplicaActive check if older replica's are still active or not if
// Status.UpdatedReplicas < *Spec.Replicas then some of the replicas are
// updated and some of them are not.
func IsOlderReplicaActive() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// IsOlderReplicaActive check if older replica's are still active or not if
// Status.UpdatedReplicas < *Spec.Replicas then some of the replicas are
// updated and some of them are not.
func (d *Deploy) IsOlderReplicaActive() bool { _ = "STUB: not implemented"; return false }

// IsTerminationInProgress checks for older replicas are waiting to
// terminate or not. If Status.Replicas > Status.UpdatedReplicas then
// some of the older replicas are in running state because newer
// replicas are not in running state. It waits for newer replica to
// come into running state then terminate.
func IsTerminationInProgress() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// IsTerminationInProgress checks for older replicas are waiting to
// terminate or not. If Status.Replicas > Status.UpdatedReplicas then
// some of the older replicas are in running state because newer
// replicas are not in running state. It waits for newer replica to
// come into running state then terminate.
func (d *Deploy) IsTerminationInProgress() bool { _ = "STUB: not implemented"; return false }

// IsUpdateInProgress Checks if all the replicas are updated or not.
// If Status.AvailableReplicas < Status.UpdatedReplicas then all the
// older replicas are not there but there are less number of availableReplicas
func IsUpdateInProgress() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// IsUpdateInProgress Checks if all the replicas are updated or not.
// If Status.AvailableReplicas < Status.UpdatedReplicas then all the
// older replicas are not there but there are less number of availableReplicas
func (d *Deploy) IsUpdateInProgress() bool { _ = "STUB: not implemented"; return false }

// IsNotSyncSpec compare generation in status and spec and check if
// deployment spec is synced or not. If Generation <= Status.ObservedGeneration
// then deployment spec is not updated yet.
func IsNotSyncSpec() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// IsNotSyncSpec compare generation in status and spec and check if
// deployment spec is synced or not. If Generation <= Status.ObservedGeneration
// then deployment spec is not updated yet.
func (d *Deploy) IsNotSyncSpec() bool { _ = "STUB: not implemented"; return false }
