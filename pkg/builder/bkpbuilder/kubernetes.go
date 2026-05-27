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

package bkpbuilder

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
	clientset "github.com/openebs/zfs-localpv/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// getClientsetFn is a typed function that
// abstracts fetching of internal clientset
type getClientsetFn func() (clientset *clientset.Clientset, err error)

// getClientsetFromPathFn is a typed function that
// abstracts fetching of clientset from kubeConfigPath
type getClientsetForPathFn func(kubeConfigPath string) (
	clientset *clientset.Clientset,
	err error,
)

// createFn is a typed function that abstracts
// creating zfsbkp bkpume instance
type createFn func(
	cs *clientset.Clientset,
	upgradeResultObj *apis.ZFSBackup,
	namespace string,
) (*apis.ZFSBackup, error)

// getFn is a typed function that abstracts
// fetching a zfsbkp bkpume instance
type getFn func(
	cli *clientset.Clientset,
	name,
	namespace string,
	opts metav1.GetOptions,
) (*apis.ZFSBackup, error)

// listFn is a typed function that abstracts
// listing of zfsbkp bkpume instances
type listFn func(
	cli *clientset.Clientset,
	namespace string,
	opts metav1.ListOptions,
) (*apis.ZFSBackupList, error)

// delFn is a typed function that abstracts
// deleting a zfsbkp bkpume instance
type delFn func(
	cli *clientset.Clientset,
	name,
	namespace string,
	opts *metav1.DeleteOptions,
) error

// updateFn is a typed function that abstracts
// updating zfsbkp bkpume instance
type updateFn func(
	cs *clientset.Clientset,
	bkp *apis.ZFSBackup,
	namespace string,
) (*apis.ZFSBackup, error)

// Kubeclient enables kubernetes API operations
// on zfsbkp bkpume instance
type Kubeclient struct {
	// clientset refers to zfsbkp bkpume's
	// clientset that will be responsible to
	// make kubernetes API calls
	clientset *clientset.Clientset

	kubeConfigPath string

	// namespace holds the namespace on which
	// kubeclient has to operate
	namespace string

	// functions useful during mocking
	getClientset        getClientsetFn
	getClientsetForPath getClientsetForPathFn
	get                 getFn
	list                listFn
	del                 delFn
	create              createFn
	update              updateFn
}

// KubeclientBuildOption defines the abstraction
// to build a kubeclient instance
type KubeclientBuildOption func(*Kubeclient)

// defaultGetClientset is the default implementation to
// get kubernetes clientset instance
func defaultGetClientset() (clients *clientset.Clientset, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaultGetClientsetForPath is the default implementation to
// get kubernetes clientset instance based on the given
// kubeconfig path
func defaultGetClientsetForPath(
	kubeConfigPath string,
) (clients *clientset.Clientset, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaultGet is the default implementation to get
// a zfsbkp bkpume instance in kubernetes cluster
func defaultGet(
	cli *clientset.Clientset,
	name, namespace string,
	opts metav1.GetOptions,
) (*apis.ZFSBackup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaultList is the default implementation to list
// zfsbkp bkpume instances in kubernetes cluster
func defaultList(
	cli *clientset.Clientset,
	namespace string,
	opts metav1.ListOptions,
) (*apis.ZFSBackupList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaultCreate is the default implementation to delete
// a zfsbkp bkpume instance in kubernetes cluster
func defaultDel(
	cli *clientset.Clientset,
	name, namespace string,
	opts *metav1.DeleteOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// defaultCreate is the default implementation to create
// a zfsbkp bkpume instance in kubernetes cluster
func defaultCreate(
	cli *clientset.Clientset,
	bkp *apis.ZFSBackup,
	namespace string,
) (*apis.ZFSBackup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaultUpdate is the default implementation to update
// a zfsbkp bkpume instance in kubernetes cluster
func defaultUpdate(
	cli *clientset.Clientset,
	bkp *apis.ZFSBackup,
	namespace string,
) (*apis.ZFSBackup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// withDefaults sets the default options
// of kubeclient instance
func (k *Kubeclient) withDefaults() { _ = "STUB: not implemented"; return }

// WithClientSet sets the kubernetes client against
// the kubeclient instance
func WithClientSet(c *clientset.Clientset) KubeclientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeclientBuildOption)
}

// WithNamespace sets the kubernetes client against
// the provided namespace
func WithNamespace(namespace string) KubeclientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeclientBuildOption)
}

// WithNamespace sets the provided namespace
// against this Kubeclient instance
func (k *Kubeclient) WithNamespace(namespace string) *Kubeclient {
	_ = "STUB: not implemented"
	return nil
}

// WithKubeConfigPath sets the kubernetes client
// against the provided path
func WithKubeConfigPath(path string) KubeclientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeclientBuildOption)
}

// NewKubeclient returns a new instance of
// kubeclient meant for zfsbkp bkpume operations
func NewKubeclient(opts ...KubeclientBuildOption) *Kubeclient {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kubeclient) getClientsetForPathOrDirect() (
	*clientset.Clientset,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getClientOrCached returns either a new instance
// of kubernetes client or its cached copy
func (k *Kubeclient) getClientOrCached() (*clientset.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a zfsbkp bkpume instance
// in kubernetes cluster
func (k *Kubeclient) Create(bkp *apis.ZFSBackup) (*apis.ZFSBackup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get returns zfsbkp bkpume object for given name
func (k *Kubeclient) Get(
	name string,
	opts metav1.GetOptions,
) (*apis.ZFSBackup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRaw returns zfsbkp bkpume instance
// in bytes
func (k *Kubeclient) GetRaw(
	name string,
	opts metav1.GetOptions,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List returns a list of zfsbkp bkpume
// instances present in kubernetes cluster
func (k *Kubeclient) List(opts metav1.ListOptions) (*apis.ZFSBackupList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deletes the zfsbkp bkpume from
// kubernetes
func (k *Kubeclient) Delete(name string) error { _ = "STUB: not implemented"; return nil }

// Update updates this zfsbkp bkpume instance
// against kubernetes cluster
func (k *Kubeclient) Update(bkp *apis.ZFSBackup) (*apis.ZFSBackup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
