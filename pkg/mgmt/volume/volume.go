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

package volume

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// isDeletionCandidate checks if a zfs volume is a deletion candidate.
func (c *ZVController) isDeletionCandidate(zv *apis.ZFSVolume) bool {
	_ = "STUB: not implemented"
	return false
}

// syncHandler compares the actual state with the desired, and attempts to
// converge the two.
func (c *ZVController) syncHandler(key string) error {
	_ = "STUB: not implemented"
	// Convert the namespace/name string into a distinct namespace and name
	return nil
}

// Get the zv resource with this namespace/name

// enqueueZV takes a ZFSVolume resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than ZFSVolume.
func (c *ZVController) enqueueZV(obj interface{}) { _ = "STUB: not implemented"; return }

// synZV is the function which tries to converge to a desired state for the
// ZFSVolume
func (c *ZVController) syncZV(zv *apis.ZFSVolume) error {
	_ = "STUB: not implemented"

	// ZFS Volume should be deleted. Check if deletion timestamp is set
	return nil
}

// destroy only if other finalizers have been removed

// if volume has already been created and its state is Ready
// then this event is for property change only.

// addZV is the add event handler for ZFSVolume
func (c *ZVController) addZV(obj interface{}) { _ = "STUB: not implemented"; return }

// updateZV is the update event handler for ZFSVolume
func (c *ZVController) updateZV(oldObj, newObj interface{}) { _ = "STUB: not implemented"; return }

// deleteZV is the delete event handler for ZFSVolume
func (c *ZVController) deleteZV(obj interface{}) { _ = "STUB: not implemented"; return }

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *ZVController) Run(threadiness int, stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Start the informer factories to begin populating the informer caches

// Wait for the k8s caches to be synced before starting workers

// Launch worker to process ZV resources
// Threadiness will decide the number of workers you want to launch to process work items from queue

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *ZVController) runWorker() { _ = "STUB: not implemented"; return }

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *ZVController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

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
// ZV resource to be synced.

// Put the item back on the workqueue to handle any transient errors.

// Finally, if no error occurs we Forget this item so it does not
// get queued again until another change happens.
