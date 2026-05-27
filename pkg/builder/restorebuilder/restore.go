// Copyright © 2020 The OpenEBS Authors
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

package restorebuilder

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// ZFSRestore is a wrapper over
// ZFSRestore API instance
type ZFSRestore struct {
	// ZFSSnap object
	Object *apis.ZFSRestore
}

// From returns a new instance of
// zfsrstr rstrume
func From(rstr *apis.ZFSRestore) *ZFSRestore { _ = "STUB: not implemented"; return nil }

// Predicate defines an abstraction
// to determine conditional checks
// against the provided pod instance
type Predicate func(*ZFSRestore) bool

// PredicateList holds a list of predicate
type predicateList []Predicate

// ZFSRestoreList holds the list
// of zfs restore instances
type ZFSRestoreList struct {
	// List contains list of restore
	List apis.ZFSRestoreList
}

// Len returns the number of items present
// in the ZFSRestoreList
func (rstrList *ZFSRestoreList) Len() int { _ = "STUB: not implemented"; return 0 }

// all returns true if all the predicates
// succeed against the provided ZFSRestore
// instance
func (l predicateList) all(rstr *ZFSRestore) bool { _ = "STUB: not implemented"; return false }

// HasLabels returns true if provided labels
// are present in the provided ZFSRestore instance
func HasLabels(keyValuePair map[string]string) Predicate {
	_ = "STUB: not implemented"
	return *new(Predicate)
}

// HasLabel returns true if provided label
// is present in the provided ZFSRestore instance
func (rstr *ZFSRestore) HasLabel(key, value string) bool { _ = "STUB: not implemented"; return false }

// HasLabel returns true if provided label
// is present in the provided ZFSRestore instance
func HasLabel(key, value string) Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// IsNil returns true if the zfsrstr rstrume instance
// is nil
func (rstr *ZFSRestore) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsNil is predicate to filter out nil zfsrstr rstrume
// instances
func IsNil() Predicate { _ = "STUB: not implemented"; return *new(Predicate) }

// GetAPIObject returns zfsrstr rstrume's API instance
func (rstr *ZFSRestore) GetAPIObject() *apis.ZFSRestore { _ = "STUB: not implemented"; return nil }
