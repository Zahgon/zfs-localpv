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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// getClientsetFn is a typed function that abstracts
// fetching an instance of kubernetes clientset
type getClientsetFn func() (clientset *kubernetes.Clientset, err error)

// getClientsetFromPathFn is a typed function that
// abstracts fetching of clientset from kubeConfigPath
type getClientsetForPathFn func(kubeConfigPath string) (clientset *kubernetes.Clientset, err error)

// listFn is a typed function that abstracts
// listing of storageclasses
type listFn func(cli *kubernetes.Clientset, opts metav1.ListOptions) (*storagev1.StorageClassList, error)

// getFn is a typed function that abstracts
// fetching an instance of storageclass
type getFn func(cli *kubernetes.Clientset, name string, opts metav1.GetOptions) (*storagev1.StorageClass, error)

// createFn is a typed function that abstracts
// to create storage class
type createFn func(cli *kubernetes.Clientset, sc *storagev1.StorageClass) (*storagev1.StorageClass, error)

// deleteFn is a typed function that abstracts
// to delete storageclass
type deleteFn func(cli *kubernetes.Clientset, name string, opts *metav1.DeleteOptions) error

// Kubeclient enables kubernetes API operations on storageclass instance
type Kubeclient struct {
	// clientset refers to storageclass clientset
	// that will be responsible to
	// make kubernetes API calls
	clientset *kubernetes.Clientset

	// kubeconfig path to get kubernetes clientset
	kubeConfigPath string

	// functions useful during mocking
	getClientset        getClientsetFn
	getClientsetForPath getClientsetForPathFn
	list                listFn
	get                 getFn
	create              createFn
	del                 deleteFn
}

// KubeClientBuildOption defines the abstraction
// to build a kubeclient instance
type KubeClientBuildOption func(*Kubeclient)

func (k *Kubeclient) withDefaults() { _ = "STUB: not implemented"; return }

// NewKubeClient returns a new instance of kubeclient meant for storageclass
func NewKubeClient(opts ...KubeClientBuildOption) *Kubeclient {
	_ = "STUB: not implemented"
	return nil
}

// WithClientSet sets the kubernetes client against
// the kubeclient instance
func WithClientSet(c *kubernetes.Clientset) KubeClientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeClientBuildOption)
}

// WithKubeConfigPath sets the kubeConfig path
// against client instance
func WithKubeConfigPath(path string) KubeClientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeClientBuildOption)
}

func (k *Kubeclient) getClientsetForPathOrDirect() (*kubernetes.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getClientsetOrCached returns either a new
// instance of kubernetes clientset or its
// cached copy cached copy
func (k *Kubeclient) getClientsetOrCached() (*kubernetes.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List returns a list of storageclass instances present in kubernetes cluster
func (k *Kubeclient) List(opts metav1.ListOptions) (*storagev1.StorageClassList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get return a storageclass instance present in kubernetes cluster
func (k *Kubeclient) Get(name string, opts metav1.GetOptions) (*storagev1.StorageClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates and returns a storageclass instance
func (k *Kubeclient) Create(sc *storagev1.StorageClass) (*storagev1.StorageClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deletes the storageclass if present in kubernetes cluster
func (k *Kubeclient) Delete(name string, opts *metav1.DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}
