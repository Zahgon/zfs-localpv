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

package pvc

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

// getClientsetFn is a typed function that
// abstracts fetching of clientset
type getClientsetFn func() (clientset *kubernetes.Clientset, err error)

// getClientsetFromPathFn is a typed function that
// abstracts fetching of clientset from kubeConfigPath
type getClientsetForPathFn func(kubeConfigPath string) (clientset *kubernetes.Clientset, err error)

// getpvcFn is a typed function that
// abstracts fetching of pvc
type getFn func(cli *kubernetes.Clientset, name string, namespace string, opts metav1.GetOptions) (*corev1.PersistentVolumeClaim, error)

// listFn is a typed function that abstracts
// listing of pvcs
type listFn func(cli *kubernetes.Clientset, namespace string, opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error)

// deleteFn is a typed function that abstracts
// deletion of pvcs
type deleteFn func(cli *kubernetes.Clientset, namespace string, name string, deleteOpts *metav1.DeleteOptions) error

// deleteFn is a typed function that abstracts
// deletion of pvc's collection
type deleteCollectionFn func(cli *kubernetes.Clientset, namespace string, listOpts metav1.ListOptions, deleteOpts *metav1.DeleteOptions) error

// createFn is a typed function that abstracts
// creation of pvc
type createFn func(cli *kubernetes.Clientset, namespace string, pvc *corev1.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error)

// updateFn is a typed function that abstracts
// updation of pvc
type updateFn func(cli *kubernetes.Clientset, namespace string, pvc *corev1.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error)

// Kubeclient enables kubernetes API operations
// on pvc instance
type Kubeclient struct {
	// clientset refers to pvc clientset
	// that will be responsible to
	// make kubernetes API calls
	clientset *kubernetes.Clientset

	// namespace holds the namespace on which
	// kubeclient has to operate
	namespace string

	// kubeconfig path to get kubernetes clientset
	kubeConfigPath string

	// functions useful during mocking
	getClientset        getClientsetFn
	getClientsetForPath getClientsetForPathFn
	list                listFn
	get                 getFn
	create              createFn
	update              updateFn
	del                 deleteFn
	delCollection       deleteCollectionFn
}

// KubeclientBuildOption abstracts creating an
// instance of kubeclient
type KubeclientBuildOption func(*Kubeclient)

// withDefaults sets the default options
// of kubeclient instance
func (k *Kubeclient) withDefaults() { _ = "STUB: not implemented"; return }

// WithClientSet sets the kubernetes client against
// the kubeclient instance
func WithClientSet(c *kubernetes.Clientset) KubeclientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeclientBuildOption)
}

// WithKubeConfigPath sets the kubeConfig path
// against client instance
func WithKubeConfigPath(path string) KubeclientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeclientBuildOption)
}

// NewKubeClient returns a new instance of kubeclient meant for
// pvc operations
func NewKubeClient(opts ...KubeclientBuildOption) *Kubeclient {
	_ = "STUB: not implemented"
	return nil
}

// WithNamespace sets the kubernetes client against
// the provided namespace
func (k *Kubeclient) WithNamespace(namespace string) *Kubeclient {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubeclient) getClientsetForPathOrDirect() (*kubernetes.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getClientsetOrCached returns either a new instance
// of kubernetes client or its cached copy
func (k *Kubeclient) getClientsetOrCached() (*kubernetes.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get returns a pvc resource
// instances present in kubernetes cluster
func (k *Kubeclient) Get(name string, opts metav1.GetOptions) (*corev1.PersistentVolumeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List returns a list of pvc
// instances present in kubernetes cluster
func (k *Kubeclient) List(opts metav1.ListOptions) (*corev1.PersistentVolumeClaimList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deletes a pvc instance from the
// kubecrnetes cluster
func (k *Kubeclient) Delete(name string, deleteOpts *metav1.DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Create creates a pvc in specified namespace in kubernetes cluster
func (k *Kubeclient) Create(pvc *corev1.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update updates a pvc in specified namespace in kubernetes cluster
func (k *Kubeclient) Update(pvc *corev1.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateCollection creates a list of pvcs
// in specified namespace in kubernetes cluster
func (k *Kubeclient) CreateCollection(
	list *corev1.PersistentVolumeClaimList,
) (*corev1.PersistentVolumeClaimList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteCollection deletes a collection of pvc objects.
func (k *Kubeclient) DeleteCollection(listOpts metav1.ListOptions, deleteOpts *metav1.DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}
