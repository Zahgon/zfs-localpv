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

package pod

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// getClientsetFn is a typed function that
// abstracts fetching of clientset
type getClientsetFn func() (*clientset.Clientset, error)

// getClientsetFromPathFn is a typed function that
// abstracts fetching of clientset from kubeConfigPath
type getClientsetForPathFn func(kubeConfigPath string) (*clientset.Clientset, error)

// getKubeConfigFn is a typed function that
// abstracts fetching of config
type getKubeConfigFn func() (*rest.Config, error)

// getKubeConfigForPathFn is a typed function that
// abstracts fetching of config from kubeConfigPath
type getKubeConfigForPathFn func(kubeConfigPath string) (*rest.Config, error)

// createFn is a typed function that abstracts
// creation of pod
type createFn func(cli *clientset.Clientset, namespace string, pod *corev1.Pod) (*corev1.Pod, error)

// listFn is a typed function that abstracts
// listing of pods
type listFn func(cli *clientset.Clientset, namespace string, opts metav1.ListOptions) (*corev1.PodList, error)

// deleteFn is a typed function that abstracts
// deleting of pod
type deleteFn func(cli *clientset.Clientset, namespace, name string, opts *metav1.DeleteOptions) error

// deleteFn is a typed function that abstracts
// deletion of pod's collection
type deleteCollectionFn func(cli *clientset.Clientset, namespace string, listOpts metav1.ListOptions, deleteOpts *metav1.DeleteOptions) error

// getFn is a typed function that abstracts
// to get pod
type getFn func(cli *clientset.Clientset, namespace, name string, opts metav1.GetOptions) (*corev1.Pod, error)

// execFn is a typed function that abstracts
// pod exec
type execFn func(cli *clientset.Clientset, config *rest.Config, name, namespace string, opts *corev1.PodExecOptions) (*ExecOutput, error)

// defaultExec is the default implementation of execFn
func defaultExec(
	cli *clientset.Clientset,
	config *rest.Config,
	name string,
	namespace string,
	opts *corev1.PodExecOptions,
) (*ExecOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create exec executor which is an interface
// for transporting shell-style streams

// Stream initiates transport of standard shell streams
// It will transport any non-nil stream to a remote system,
// and return an error if a problem occurs

// KubeClient enables kubernetes API operations
// on pod instance
type KubeClient struct {
	// clientset refers to pod clientset
	// that will be responsible to
	// make kubernetes API calls
	clientset *clientset.Clientset

	// namespace holds the namespace on which
	// KubeClient has to operate
	namespace string

	// kubeConfig represents kubernetes config
	kubeConfig *rest.Config

	// kubeconfig path to get kubernetes clientset
	kubeConfigPath string

	// functions useful during mocking
	getKubeConfig        getKubeConfigFn
	getKubeConfigForPath getKubeConfigForPathFn
	getClientset         getClientsetFn
	getClientsetForPath  getClientsetForPathFn
	create               createFn
	list                 listFn
	del                  deleteFn
	delCollection        deleteCollectionFn
	get                  getFn
	exec                 execFn
}

// ExecOutput struct contains stdout and stderr
type ExecOutput struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

// KubeClientBuildOption defines the abstraction
// to build a KubeClient instance
type KubeClientBuildOption func(*KubeClient)

// withDefaults sets the default options
// of KubeClient instance
func (k *KubeClient) withDefaults() { _ = "STUB: not implemented"; return }

// WithClientSet sets the kubernetes client against
// the KubeClient instance
func WithClientSet(c *clientset.Clientset) KubeClientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeClientBuildOption)
}

// WithKubeConfigPath sets the kubeConfig path
// against client instance
func WithKubeConfigPath(path string) KubeClientBuildOption {
	_ = "STUB: not implemented"
	return *new(KubeClientBuildOption)
}

// NewKubeClient returns a new instance of KubeClient meant for
// zfs volume replica operations
func NewKubeClient(opts ...KubeClientBuildOption) *KubeClient {
	_ = "STUB: not implemented"
	return nil
}

// WithNamespace sets the kubernetes namespace against
// the provided namespace
func (k *KubeClient) WithNamespace(namespace string) *KubeClient {
	_ = "STUB: not implemented"
	return nil
}

// WithKubeConfig sets the kubernetes config against
// the KubeClient instance
func (k *KubeClient) WithKubeConfig(config *rest.Config) *KubeClient {
	_ = "STUB: not implemented"
	return nil
}

func (k *KubeClient) getClientsetForPathOrDirect() (
	*clientset.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getClientsetOrCached returns either a new instance
// of kubernetes client or its cached copy
func (k *KubeClient) getClientsetOrCached() (*clientset.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KubeClient) getKubeConfigForPathOrDirect() (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getKubeConfigOrCached returns either a new instance
// of kubernetes config or its cached copy
func (k *KubeClient) getKubeConfigOrCached() (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List returns a list of pod
// instances present in kubernetes cluster
func (k *KubeClient) List(opts metav1.ListOptions) (*corev1.PodList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deletes a pod instance present in kubernetes cluster
func (k *KubeClient) Delete(name string, opts *metav1.DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Create creates a pod in specified namespace in kubernetes cluster
func (k *KubeClient) Create(pod *corev1.Pod) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get gets a pod object present in kubernetes cluster
func (k *KubeClient) Get(name string,
	opts metav1.GetOptions) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRaw gets pod object for a given name and namespace present
// in kubernetes cluster and returns result in raw byte.
func (k *KubeClient) GetRaw(name string,
	opts metav1.GetOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exec runs a command remotely in a container of a pod
func (k *KubeClient) Exec(name string,
	opts *corev1.PodExecOptions) (*ExecOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExecRaw runs a command remotely in a container of a pod
// and returns raw output
func (k *KubeClient) ExecRaw(name string,
	opts *corev1.PodExecOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteCollection deletes a collection of pod objects.
func (k *KubeClient) DeleteCollection(listOpts metav1.ListOptions, deleteOpts *metav1.DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}
