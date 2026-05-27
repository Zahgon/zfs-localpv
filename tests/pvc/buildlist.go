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

package pvc

import (
	corev1 "k8s.io/api/core/v1"
)

// ListBuilder enables building an instance of
// List
type ListBuilder struct {
	// template to build a list of pvcs
	template *corev1.PersistentVolumeClaim

	// count determines the number of
	// pvcs to be built using the provided
	// template
	count int

	list    *List
	filters PredicateList
	errs    []error
}

// NewListBuilder returns an instance of ListBuilder
func NewListBuilder() *ListBuilder { _ = "STUB: not implemented"; return nil }

// ListBuilderFromTemplate returns a new instance of
// ListBuilder based on the provided pvc template
func ListBuilderFromTemplate(pvc *corev1.PersistentVolumeClaim) *ListBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ListBuilderForAPIObjects returns a new instance of
// ListBuilder based on provided api pvc list
func ListBuilderForAPIObjects(pvcs *corev1.PersistentVolumeClaimList) *ListBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ListBuilderForObjects returns a new instance of
// ListBuilder based on provided pvc list
func ListBuilderForObjects(pvcs *List) *ListBuilder { _ = "STUB: not implemented"; return nil }

// WithFilter adds filters on which the pvcs
// are filtered
func (b *ListBuilder) WithFilter(pred ...Predicate) *ListBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithCount sets the count that determines
// the number of pvcs to be built
func (b *ListBuilder) WithCount(count int) *ListBuilder { _ = "STUB: not implemented"; return nil }

func (b *ListBuilder) buildFromTemplateIfNilList() { _ = "STUB: not implemented"; return }

// List returns the list of pvc instances
// that was built by this builder
func (b *ListBuilder) List() (*List, error) { _ = "STUB: not implemented"; return nil, nil }

// Len returns the number of items present
// in the List of a builder
func (b *ListBuilder) Len() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// APIList builds core API PVC list using listbuilder
func (b *ListBuilder) APIList() (*corev1.PersistentVolumeClaimList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
