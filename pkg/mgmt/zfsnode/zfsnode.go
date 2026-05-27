/*
 Copyright © 2021 The OpenEBS Authors

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

package zfsnode

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *NodeController) listZFSPool() ([]apis.Pool, error) {
	_ = "STUB: not implemented"
	return nil,

		// syncHandler compares the actual state with the desired, and attempts to
		// converge the two.
		nil
}

func (c *NodeController) syncHandler(key string) error {
	_ = "STUB: not implemented"
	// Convert the namespace/name string into a distinct namespace and name
	return nil
}

// syncNode is the function which tries to converge to a desired state for the
// ZFSNode
func (c *NodeController) syncNode(namespace string, name string) error {
	_ = "STUB: not implemented"
	// Get the node resource with this namespace/name
	return nil
}

// if it doesn't exists, create zfs node object

// zfs node already exists check if we need to update it.

// validate if owner reference updated.

// validate if node pools are upto date.

// addNode is the add event handler for ZFSNode
func (c *NodeController) addNode(obj interface{}) { _ = "STUB: not implemented"; return }

// updateNode is the update event handler for ZFSNode
func (c *NodeController) updateNode(oldObj, newObj interface{}) { _ = "STUB: not implemented"; return }

// deleteNode is the delete event handler for ZFSNode
func (c *NodeController) deleteNode(obj interface{}) { _ = "STUB: not implemented"; return }

// enqueueNode takes a ZFSNode resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than ZFSNode.
func (c *NodeController) enqueueNode(node *apis.ZFSNode) {
	_ = "STUB: not implemented"
	// node must exists in openebs namespace & must equal to the node id.
	return
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *NodeController) Run(threadiness int, stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Start the informer factories to begin populating the informer caches

// Wait for the k8s caches to be synced before starting workers

// Launch worker to process Node resources
// Threadiness will decide the number of workers you want to launch to process work items from queue

// add the item to worker queue.

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *NodeController) runWorker() { _ = "STUB: not implemented"; return }

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *NodeController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

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
// Node resource to be synced.

// Put the item back on the workqueue to handle any transient errors.

// Finally, if no error occurs we Forget this item so it does not
// get queued again until another change happens.

// isOwnerRefUpdateRequired validates if relevant owner references is being
// set for zfs node. If not, it returns the final owner references that needs
// to be set.
func (c *NodeController) isOwnerRefsUpdateRequired(ownerRefs []metav1.OwnerReference) ([]metav1.OwnerReference, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// in case owner reference exists, validate
// if controller field is set correctly or not.
