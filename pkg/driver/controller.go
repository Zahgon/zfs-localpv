package driver

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	analytics "github.com/openebs/google-analytics-4/usage"
	k8sapi "github.com/openebs/lib-csi/pkg/client/k8s"
	"github.com/openebs/lib-csi/pkg/common/errors"
	"golang.org/x/net/context"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"

	clientset "github.com/openebs/zfs-localpv/pkg/generated/clientset/versioned"
	informers "github.com/openebs/zfs-localpv/pkg/generated/informer/externalversions"
	"github.com/openebs/zfs-localpv/pkg/version"
	"github.com/openebs/zfs-localpv/pkg/zfs"
)

// size constants
const (
	MB = 1000 * 1000
	GB = 1000 * 1000 * 1000
	Mi = 1024 * 1024
	Gi = 1024 * 1024 * 1024

	// Ping event is sent periodically
	Ping string = "zfs-ping"
	// Heartbeat message.
	Heartbeat string = "zfs-heartbeat"
	// DefaultCASType Event application name constant for volume event
	DefaultCASType string = "zfs-localpv"

	// LocalPVReplicaCount is the constant used by usage to represent
	// replication factor in LocalPV
	LocalPVReplicaCount string = "1"
)

// controller is the server implementation
// for CSI Controller
type controller struct {
	driver       *CSIDriver
	capabilities []*csi.ControllerServiceCapability

	indexedLabel string

	k8sNodeInformer cache.SharedIndexInformer
	zfsNodeInformer cache.SharedIndexInformer

	volumeLock *volumeLock

	csi.UnimplementedControllerServer
}

// NewController returns a new instance
// of CSI controller
func NewController(d *CSIDriver) csi.ControllerServer {
	_ = "STUB: not implemented"
	return *new(csi.ControllerServer)
}

func (cs *controller) init() error {
	cfg, err := k8sapi.Config().Get()
	if err != nil {
		return errors.Wrapf(err, "failed to build kubeconfig")
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to build k8s clientset")
	}

	openebsClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to build openebs clientset")
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, 0)
	openebsInformerfactory := informers.NewSharedInformerFactoryWithOptions(openebsClient,
		0, informers.WithNamespace(zfs.OpenEBSNamespace))

	// set up signals so we handle the first shutdown signal gracefully
	// TODO: (tech-debt) Setup signal handler more above, several files want to use stopCh and this function is only allowed to be used once (see #647)
	// Affected files: pkg/driver/agent.go pkg/driver/controller.go pkg/driver/grpc.go
	stopCtx := ctrl.SetupSignalHandler()
	stopCh := stopCtx.Done()

	cs.k8sNodeInformer = kubeInformerFactory.Core().V1().Nodes().Informer()
	cs.zfsNodeInformer = openebsInformerfactory.Zfs().V1().ZFSNodes().Informer()

	if err = cs.zfsNodeInformer.AddIndexers(map[string]cache.IndexFunc{
		LabelIndexName(cs.indexedLabel): LabelIndexFunc(cs.indexedLabel),
	}); err != nil {
		return errors.Wrapf(err, "failed to add index on label %v", cs.indexedLabel)
	}

	go cs.k8sNodeInformer.Run(stopCh)
	go cs.zfsNodeInformer.Run(stopCh)

	if zfs.GoogleAnalyticsEnabled == "true" {
		analytics.RegisterVersionGetter(version.GetVersionDetails)
		analytics.New().CommonBuild(DefaultCASType).InstallBuilder(true).Send()
		go analytics.PingCheck(DefaultCASType, Ping, false)
		go analytics.PingCheck(DefaultCASType, Heartbeat, true)
	}

	// wait for all the caches to be populated.
	klog.Info("waiting for k8s & zfs node informer caches to be synced")
	cache.WaitForCacheSync(stopCh,
		cs.k8sNodeInformer.HasSynced,
		cs.zfsNodeInformer.HasSynced)
	klog.Info("synced k8s & zfs node informer caches")

	if zfs.ZFSBackupGCEnabled {
		bgc := &BackupGarbageCollector{}
		err = bgc.Initialize(openebsClient, stopCh)
		if err != nil {
			return errors.Wrap(err, "failed to initialize backup garbage collector")
		}
	}
	return nil
}

// SupportedVolumeCapabilityAccessModes contains the list of supported access
// modes for the volume
var SupportedVolumeCapabilityAccessModes = []*csi.VolumeCapability_AccessMode{
	{
		Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
	},
}

// sendEventOrIgnore sends anonymous local-pv provision/delete events
func sendEventOrIgnore(pvcName, pvName, capacity, method string) { _ = "STUB: not implemented"; return }

// getRoundedCapacity rounds the capacity on 1024 base
func getRoundedCapacity(size int64) int64 {
	_ = "STUB: not implemented"

	/*
	 * volblocksize and recordsize must be power of 2 from 512B to 1M
	 * so keeping the size in the form of Gi or Mi should be
	 * sufficient to make volsize multiple of volblocksize/recordsize.
	 */return 0
}

// Keeping minimum allocatable size as 1Mi (1024 * 1024)

func waitForVolDestroy(volname string) error { _ = "STUB: not implemented"; return nil }

func waitForReadySnapshot(ctx context.Context, snapname string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateZFSVolume create new zfs volume from csi volume request
func CreateZFSVolume(ctx context.Context, req *csi.CreateVolumeRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// parameter keys may be mistyped from the CRD specification when declaring
// the storageclass, which kubectl validation will not catch. Because ZFS
// parameter keys (not values!) are all lowercase, keys may safely be forced
// to the lower case.

// (hack): CSI Sanity test does not pass topology information

// run the scheduler

// try volume creation sequentially on all nodes

// if timeout reached, return the error and let csi retry the volume creation

// volume provisioning failed, delete the zfs volume resource
// ignore error

// CreateVolClone creates the clone from a volume
func CreateVolClone(ctx context.Context, req *csi.CreateVolumeRequest, srcVol string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// lower case keys, cf CreateZFSVolume()

// create the clone from the source volume

// use the snapshot name same as new volname

// CreateSnapClone creates the clone from a snapshot
func CreateSnapClone(ctx context.Context, req *csi.CreateVolumeRequest, snapshot string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// lower case keys, cf CreateZFSVolume()

// CreateVolume provisions a volume
func (cs *controller) CreateVolume(
	ctx context.Context,
	req *csi.CreateVolumeRequest,
) (*csi.CreateVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lower case keys, cf CreateZFSVolume()

// DeleteVolume deletes the specified volume
func (cs *controller) DeleteVolume(
	ctx context.Context,
	req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verify if the volume has already been deleted

// if volume is not ready, create volume will delete it

// Fetch the list of snapshot for the given volume

// Delete the corresponding ZV CR only if there are no snapshots present for the volume

// add annotation to the volume to indicate that it is eligible for deletion
// once all the snapshots are deleted and the reclaim policy is not Retain
// this volume will be deleted

func isValidVolumeCapabilities(volCaps []*csi.VolumeCapability) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO Implementation will be taken up later

// ValidateVolumeCapabilities validates the capabilities
// required to create a new volume
// This implements csi.ControllerServer
func (cs *controller) ValidateVolumeCapabilities(
	ctx context.Context,
	req *csi.ValidateVolumeCapabilitiesRequest,
) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ControllerGetCapabilities fetches controller capabilities
//
// This implements csi.ControllerServer
func (cs *controller) ControllerGetCapabilities(
	ctx context.Context,
	req *csi.ControllerGetCapabilitiesRequest,
) (*csi.ControllerGetCapabilitiesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ControllerExpandVolume resizes previously provisioned volume
//
// This implements csi.ControllerServer
func (cs *controller) ControllerExpandVolume(
	ctx context.Context,
	req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* round off the new size */

/*
 * Controller expand volume must be idempotent. If a volume corresponding
 * to the specified volume ID is already larger than or equal to the target
 * capacity of the expansion request, the plugin should reply 0 OK.
 */

func verifySnapshotRequest(req *csi.CreateSnapshotRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateSnapshot creates a snapshot for given volume
//
// This implements csi.ControllerServer
func (cs *controller) CreateSnapshot(
	ctx context.Context,
	req *csi.CreateSnapshotRequest,
) (*csi.CreateSnapshotResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSnapshot deletes given snapshot
//
// This implements csi.ControllerServer
func (cs *controller) DeleteSnapshot(
	ctx context.Context,
	req *csi.DeleteSnapshotRequest,
) (*csi.DeleteSnapshotResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// snapshodID is formed as <volname>@<snapname>
// parsing them here

// should succeed when an invalid snapshot id is used

// verify if the snapshot has already been deleted

// Fetch the list of snapshot for the given volume

// Delete the corresponding ZV CR only if this is the last snapshot
// for the volume and the corresponding pvc is deleted

// ListSnapshots lists all snapshots for the
// given volume
//
// This implements csi.ControllerServer
func (cs *controller) ListSnapshots(
	ctx context.Context,
	req *csi.ListSnapshotsRequest,
) (*csi.ListSnapshotsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ControllerUnpublishVolume removes a previously
// attached volume from the given node
//
// This implements csi.ControllerServer
func (cs *controller) ControllerUnpublishVolume(
	ctx context.Context,
	req *csi.ControllerUnpublishVolumeRequest,
) (*csi.ControllerUnpublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ControllerPublishVolume attaches given volume
// at the specified node
//
// This implements csi.ControllerServer
func (cs *controller) ControllerPublishVolume(
	ctx context.Context,
	req *csi.ControllerPublishVolumeRequest,
) (*csi.ControllerPublishVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCapacity return the capacity of the
// given node topology segment.
//
// This implements csi.ControllerServer
func (cs *controller) GetCapacity(
	ctx context.Context,
	req *csi.GetCapacityRequest,
) (*csi.GetCapacityResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The "poolname" parameter can either be the name of a ZFS pool
// (e.g. "zpool"), or a path to a child dataset (e.g. "zpool/k8s/localpv").
//
// We parse the "poolname" parameter so the name of the ZFS pool and the
// path to the dataset is available separately.
//
// The dataset path is not used now. It could be used later to query the
// capacity of the child dataset, which could be smaller than the capacity
// of the whole pool.
//
// This is necessary because capacity calculation currently only works with
// ZFS pool names. This is why it always returns the capacitry of the whole
// pool, even if the child dataset given as the "poolname" parameter has a
// smaller capacity than the whole pool.

// rather than summing all free capacity, we are calculating maximum
// zv size that gets fit in given pool.
// See https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1472-storage-capacity-tracking#available-capacity-vs-maximum-volume-size &
// https://github.com/container-storage-interface/spec/issues/432 for more details

func (cs *controller) filterNodesByTopology(segments map[string]string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// first see if we need to filter the informer cache by indexed label,
// so that we don't need to iterate over all the nodes for performance
// reasons in large cluster.

// run through all the nodes in case indexer doesn't exists.

// ListVolumes lists all the volumes
//
// This implements csi.ControllerServer
func (cs *controller) ListVolumes(
	ctx context.Context,
	req *csi.ListVolumesRequest,
) (*csi.ListVolumesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *controller) validateDeleteVolumeReq(req *csi.DeleteVolumeRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// IsSupportedVolumeCapabilityAccessMode valides the requested access mode
func IsSupportedVolumeCapabilityAccessMode(
	accessMode csi.VolumeCapability_AccessMode_Mode,
) bool {
	_ = "STUB: not implemented"
	return false
}

// newControllerCapabilities returns a list
// of this controller's capabilities
func newControllerCapabilities() []*csi.ControllerServiceCapability {
	_ = "STUB: not implemented"
	return nil
}

// validateRequest validates if the requested service is
// supported by the driver
func (cs *controller) validateRequest(
	c csi.ControllerServiceCapability_RPC_Type,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *controller) validateVolumeCreateReq(req *csi.CreateVolumeRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// LabelIndexName add prefix for label index.
func LabelIndexName(label string) string { _ = "STUB: not implemented"; return "" }

// LabelIndexFunc defines index values for given label.
func LabelIndexFunc(label string) cache.IndexFunc {
	_ = "STUB: not implemented"
	return *new(cache.IndexFunc)
}

func (cs *controller) ControllerGetVolume(
	ctx context.Context,
	req *csi.ControllerGetVolumeRequest,
) (*csi.ControllerGetVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *controller) ControllerModifyVolume(
	ctx context.Context,
	req *csi.ControllerModifyVolumeRequest,
) (*csi.ControllerModifyVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
