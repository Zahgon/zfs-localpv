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
	storagev1 "k8s.io/api/storage/v1"
)

// ListBuilder enables building an instance of StorageClassList
type ListBuilder struct {
	list    *StorageClassList
	filters predicateList
	errs    []error
}

// NewListBuilder returns a instance of ListBuilder
func NewListBuilder() *ListBuilder { _ = "STUB: not implemented"; return nil }

// ListBuilderForAPIList builds the ListBuilder object based on SC API list
func ListBuilderForAPIList(scl *storagev1.StorageClassList) *ListBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ListBuilderForObjects returns a instance of ListBuilder from SC instances
func ListBuilderForObjects(scl *StorageClassList) *ListBuilder {
	_ = "STUB: not implemented"
	return nil
}

// List returns the list of StorageClass instances that was built by this builder
func (b *ListBuilder) List() (*StorageClassList, error) { _ = "STUB: not implemented"; return nil, nil }

// Pin it

// WithFilter add filters on which the StorageClass has to be filtered
func (b *ListBuilder) WithFilter(pred ...Predicate) *ListBuilder {
	_ = "STUB: not implemented"
	return nil
}

// APIList builds core API PVC list using listbuilder
func (b *ListBuilder) APIList() (*storagev1.StorageClassList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Len returns the number of items present
// in the List of a builder
func (b *ListBuilder) Len() (int, error) { _ = "STUB: not implemented"; return 0, nil }
