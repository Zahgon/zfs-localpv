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
	"sync"
)

// Start starts the zfsnode controller.
func Start(controllerMtx *sync.RWMutex, stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"

	// Get in cluster config
	return nil
}

// Building Kubernetes Clientset

// Building OpenEBS Clientset

// setup watch only on node we are interested in.

// as object returned by client go clears all TypeMeta from it.

// Build() fn of all controllers calls AddToScheme to adds all types of this
// clientset into the given scheme.
// If multiple controllers happen to call this AddToScheme same time,
// it causes panic with error saying concurrent map access.
// This lock is used to serialize the AddToScheme call of all controllers.

// blocking call, can't use defer to release the lock

// Threadiness defines the number of workers to be launched in Run function
