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

package tests

import (
	apis "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1"
)

// IsPVCBoundEventually checks if the pvc is bound or not eventually
func IsPVCBoundEventually(pvcName string) bool { _ = "STUB: not implemented"; return false }

// IsPVAvailableEventually checks if the pv is bound or not eventually
func IsPVAvailableEventually(pvName string) bool { _ = "STUB: not implemented"; return false }

// IsPVCResizedEventually checks if the pvc is bound or not eventually
func IsPVCResizedEventually(pvcName string, newCapacity string) bool {
	_ = "STUB: not implemented"
	return false
}

// IsPodRunningEventually return true if the pod comes to running state
func IsPodRunningEventually(namespace, podName string) bool {
	_ = "STUB: not implemented"
	return false
}

// IsPropUpdatedEventually checks if the property is updated or not eventually
func IsPropUpdatedEventually(vol *apis.ZFSVolume, prop string, val string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetVolumeProperty gets zfs properties for the volume
func GetVolumeProperty(vol *apis.ZFSVolume, prop string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsPVCDeletedEventually tries to get the deleted pvc
// and returns true if pvc is not found
// else returns false
func IsPVCDeletedEventually(pvcName string) bool { _ = "STUB: not implemented"; return false }

// IsPVDeletedEventually tries to get the deleted pv
// and returns true if pv is not found
// else returns false
func IsPVDeletedEventually(pvName string) bool { _ = "STUB: not implemented"; return false }

// IsZVDeletedEventually tries to get the deleted zv
// and returns true if zv is not found
// else returns false
func IsZVDeletedEventually(zvName string) bool { _ = "STUB: not implemented"; return false }

// VerifyStorageClassParams verifies the volume properties set at creation time
func VerifyStorageClassParams(property map[string]string) { _ = "STUB: not implemented"; return }

// Check for file system type

// It populates the map for thing provisioning params
// Refer https://github.com/openebs/zfs-localpv/issues/560#issuecomment-2232535073
func generateThinProvisionParams(property map[string]string) { _ = "STUB: not implemented"; return }

func createFstypeStorageClass(addons map[string]string) { _ = "STUB: not implemented"; return }

// Update params with addons

func createStorageClassWithReclaimPolicy() { _ = "STUB: not implemented"; return }

func createStorageClass() { _ = "STUB: not implemented"; return }

func createEncryptedStorageClass() { _ = "STUB: not implemented"; return }

// VerifyZFSVolume verify the properties of a zfs-volume
func VerifyZFSVolume() { _ = "STUB: not implemented"; return }

// it might fail if we are checking finializer before event is processed by node agent

// VerifyZFSVolumePropEdit verifies the volume properties
func VerifyZFSVolumePropEdit() { _ = "STUB: not implemented"; return }

// 4k

// 8k

func deleteStorageClass() { _ = "STUB: not implemented"; return }

func createAndVerifyPVC(pvcName string) { _ = "STUB: not implemented"; return }

func resizeAndVerifyPVC(pvcName string) { _ = "STUB: not implemented"; return }

func createDeployVerifyApp(appName, pvcName string) { _ = "STUB: not implemented"; return }

func createDeployVerifyCloneApp(cloneAppName, clonePvcName string) {
	_ = "STUB: not implemented"
	return
}

func createAndDeployAppPod(appname string, pvcname string) { _ = "STUB: not implemented"; return }

func createAndDeployBlockAppPod(appName, pvcName string) { _ = "STUB: not implemented"; return }

func verifyAppPodRunning(appname string) { _ = "STUB: not implemented"; return }

func deleteAppDeployment(appname string) { _ = "STUB: not implemented"; return }

func deletePVC(pvcname string) { _ = "STUB: not implemented"; return }

func deletePV(pvName string) { _ = "STUB: not implemented"; return }

// DeleteZV deletes the zv
func DeleteZV(zvName string) { _ = "STUB: not implemented"; return }

func getStoragClassParams() []map[string]string { _ = "STUB: not implemented"; return nil }

// IsZVPresentConsistently checks if the zfs volume is present or not consistently
func IsZVPresentConsistently(zvName string) bool { _ = "STUB: not implemented"; return false }

// Check consistency for 60 seconds, polling every 5 seconds

// getZVName return the zv name
func getZVName(pvcName string) string { _ = "STUB: not implemented"; return "" }

func createAndVerifyPVFromRetainedZV(pvName, volumeHandle string) {
	_ = "STUB: not implemented"
	return
}
