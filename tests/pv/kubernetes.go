package pv

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
type getFn func(cli *kubernetes.Clientset, name string, opts metav1.GetOptions) (*corev1.PersistentVolume, error)

// createFn is a typed function that abstracts
// creation of pvc
type createFn func(cli *kubernetes.Clientset, pvc *corev1.PersistentVolume) (*corev1.PersistentVolume, error)

// deleteFn is a typed function that abstracts
// deletion of pvcs
type deleteFn func(cli *kubernetes.Clientset, name string, deleteOpts *metav1.DeleteOptions) error

// Kubeclient enables kubernetes API operations
// on pvc instance
type Kubeclient struct {
	// clientset refers to pvc clientset
	// that will be responsible to
	// make kubernetes API calls
	clientset *kubernetes.Clientset

	// kubeconfig path to get kubernetes clientset
	kubeConfigPath string

	// functions useful during mocking
	getClientset        getClientsetFn
	getClientsetForPath getClientsetForPathFn
	get                 getFn
	create              createFn
	del                 deleteFn
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
func (k *Kubeclient) Get(name string, opts metav1.GetOptions) (*corev1.PersistentVolume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deletes a pvc instance from the
// kubecrnetes cluster
func (k *Kubeclient) Delete(name string, deleteOpts *metav1.DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Create creates a pv in specified namespace in kubernetes cluster
func (k *Kubeclient) Create(pv *corev1.PersistentVolume) (*corev1.PersistentVolume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
