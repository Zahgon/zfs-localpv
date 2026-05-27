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
	"time"

	clientset "github.com/openebs/zfs-localpv/pkg/generated/clientset/versioned"
	informers "github.com/openebs/zfs-localpv/pkg/generated/informer/externalversions"
	listers "github.com/openebs/zfs-localpv/pkg/generated/lister/zfs/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

const controllerAgentName = "zfsnode-controller"

// NodeController is the controller implementation for zfs node resources
type NodeController struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface

	// clientset is a openebs custom resource package generated for custom API group.
	clientset clientset.Interface

	NodeLister listers.ZFSNodeLister

	// NodeSynced is used for caches sync to get populated
	NodeSynced cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface

	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder

	// pollInterval controls the polling frequency of syncing up the vg metadata.
	pollInterval time.Duration

	// ownerRef is used to set the owner reference to zfsnode objects.
	ownerRef metav1.OwnerReference
}

// NodeControllerBuilder is the builder object for controller.
type NodeControllerBuilder struct {
	NodeController *NodeController
}

// NewNodeControllerBuilder returns an empty instance of controller builder.
func NewNodeControllerBuilder() *NodeControllerBuilder { _ = "STUB: not implemented"; return nil }

// withKubeClient fills kube client to controller object.
func (cb *NodeControllerBuilder) withKubeClient(ks kubernetes.Interface) *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// withOpenEBSClient fills openebs client to controller object.
func (cb *NodeControllerBuilder) withOpenEBSClient(cs clientset.Interface) *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// withNodeLister fills Node lister to controller object.
func (cb *NodeControllerBuilder) withNodeLister(sl informers.SharedInformerFactory) *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// withNodeSynced adds object sync information in cache to controller object.
func (cb *NodeControllerBuilder) withNodeSynced(sl informers.SharedInformerFactory) *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// withWorkqueue adds workqueue to controller object.
func (cb *NodeControllerBuilder) withWorkqueueRateLimiting() *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// withRecorder adds recorder to controller object.
func (cb *NodeControllerBuilder) withRecorder(ks kubernetes.Interface) *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// withEventHandler adds event handlers controller object.
func (cb *NodeControllerBuilder) withEventHandler(cvcInformerFactory informers.SharedInformerFactory) *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Set up an event handler for when zfs node vg change.
// Note: rather than setting up the resync period at informer level,
// we are controlling the syncing based on pollInternal. See
// NodeController#Run func for more details.

func (cb *NodeControllerBuilder) withPollInterval(interval time.Duration) *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (cb *NodeControllerBuilder) withOwnerReference(ownerRef metav1.OwnerReference) *NodeControllerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Build returns a controller instance.
func (cb *NodeControllerBuilder) Build() (*NodeController, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
