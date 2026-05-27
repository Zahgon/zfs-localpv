/*
Copyright 2020 The OpenEBS Authors

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

package restore

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// isDeletionCandidate checks if a zfs backup is a deletion candidate.
func (c *RstrController) isDeletionCandidate(rstr *apis.ZFSRestore) bool {
	_ = "STUB: not implemented"
	return false
}

// syncHandler compares the actual state with the desired, and attempts to
// converge the two.
func (c *RstrController) syncHandler(key string) error {
	_ = "STUB: not implemented"
	// Convert the namespace/name string into a distinct namespace and name
	return nil
}

// Get the rstr resource with this namespace/name

// enqueueRestore takes a ZFSRestore resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than ZFSRestore.
func (c *RstrController) enqueueRestore(obj interface{}) { _ = "STUB: not implemented"; return }

// synRestore is the function which tries to converge to a desired state for the
// ZFSRestore
func (c *RstrController) syncRestore(rstr *apis.ZFSRestore) error {
	_ = "STUB: not implemented"

	// ZFSRestore should not be deleted. Check if deletion timestamp is set
	return nil
}

// if status is Init, then only do the restore

// addRestore is the add event handler for ZFSRestore
func (c *RstrController) addRestore(obj interface{}) { _ = "STUB: not implemented"; return }

// updateRestore is the update event handler for ZFSRestore
func (c *RstrController) updateRestore(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// deleteRestore is the delete event handler for ZFSRestore
func (c *RstrController) deleteRestore(obj interface{}) { _ = "STUB: not implemented"; return }

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *RstrController) Run(threadiness int, stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Start the informer factories to begin populating the informer caches

// Wait for the k8s caches to be synced before starting workers

// Launch worker to process Restore resources
// Threadiness will decide the number of workers you want to launch to process work items from queue

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *RstrController) runWorker() { _ = "STUB: not implemented"; return }

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *RstrController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// We wrap this block in a func so we can defer c.workqueue.Done.

// We call Done here so the workqueue knows we have finished
// processing this item. We also must remember to call Forget if we
// do not want this work item being re-queued. For example, we do
// not call Forget if a transient error occurs, instead the item is
// put back on the workqueue and attempted again after a back-off
// period.

// We expect strings to come off the workqueue. These are of the
// form namespace/name. We do this as the delayed nature of the
// workqueue means the items in the informer cache may actually be
// more up to date that when the item was initially put onto the
// workqueue.

// As the item in the workqueue is actually invalid, we call
// Forget here else we'd go into a loop of attempting to
// process a work item that is invalid.

// Run the syncHandler, passing it the namespace/name string of the
// Restore resource to be synced.

// Put the item back on the workqueue to handle any transient errors.

// Finally, if no error occurs we Forget this item so it does not
// get queued again until another change happens.
