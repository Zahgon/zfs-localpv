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

package pod

import (
	corev1 "k8s.io/api/core/v1"
)

// Pod holds the api's pod objects
type Pod struct {
	object *corev1.Pod
}

// List holds the list of API pod instances
type List struct {
	items []*Pod
}

// PredicateList holds a list of predicate
type predicateList []Predicate

// Predicate defines an abstraction
// to determine conditional checks
// against the provided pod instance
type Predicate func(*Pod) bool

// ToAPIList converts List to API List
func (pl *List) ToAPIList() *corev1.PodList { _ = "STUB: not implemented"; return nil }

type podBuildOption func(*Pod)

// NewForAPIObject returns a new instance of Pod
func NewForAPIObject(obj *corev1.Pod, opts ...podBuildOption) *Pod {
	_ = "STUB: not implemented"
	return nil
}

// Len returns the number of items present in the List
func (pl *List) Len() int { _ = "STUB: not implemented"; return 0 }

// all returns true if all the predicates
// succeed against the provided pod
// instance
func (l predicateList) all(p *Pod) bool { _ = "STUB: not implemented"; return false }

// IsRunning returns true if the pod is in running
// state
func (p *Pod) IsRunning() bool { _ = "STUB: not implemented"; return false }

// IsRunning is a predicate to filter out pods
// which in running state
func IsRunning() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// IsCompleted returns true if the pod is in completed
// state
func (p *Pod) IsCompleted() bool { _ = "STUB: not implemented"; return false }

// IsCompleted is a predicate to filter out pods
// which in completed state
func IsCompleted() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// HasLabels returns true if provided labels
// map[key]value are present in the provided List
// instance
func HasLabels(keyValuePair map[string]string) Predicate {
	_ = "STUB: not implemented"
	return *new(Predicate)
}

//		objKeyValues := p.object.GetLabels()

// HasLabel return true if provided lable
// key and value are present in the the provided List
// instance
func (p *Pod) HasLabel(key, value string) bool { _ = "STUB: not implemented"; return false }

// HasLabel is predicate to filter out labeled
// pod instances
func HasLabel(key, value string) Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// IsNil returns true if the pod instance
// is nil
func (p *Pod) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsNil is predicate to filter out nil pod
// instances
func IsNil() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// GetAPIObject returns a API's Pod
func (p *Pod) GetAPIObject() *corev1.Pod {
	_ = "STUB: not implemented"

	// FromList created a List with provided api List
	return nil
}

func FromList(pods *corev1.PodList) *List { _ = "STUB: not implemented"; return nil }

// GetScheduledNodes returns the nodes on which pods are scheduled
func (pl *List) GetScheduledNodes() map[string]int { _ = "STUB: not implemented"; return nil }

// pin it

// IsMatchNodeAny checks the List is running on the provided nodes
func (pl *List) IsMatchNodeAny(nodes map[string]int) bool { _ = "STUB: not implemented"; return false }

// pin it
