/*
Copyright 2020 The OpenEBS Authors

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

package driver

// scheduling algorithm constants
const (
	// pick the node where less volumes are provisioned for the given pool
	VolumeWeighted = "VolumeWeighted"

	// pick the node where total provisioned volumes have occupied less capacity from the given pool
	// this will be the default scheduler when none provided
	CapacityWeighted = "CapacityWeighted"
)

// getVolumeWeightedMap goes through all the pools on all the nodes
// and creates the node mapping of the volume for all the nodes.
// It returns a map which has nodes as key and volumes present
// on the nodes as corresponding value.
func getVolumeWeightedMap(pool string) (map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create the map of the volume count
// for the given pool

// getCapacityWeightedMap goes through all the pools on all the nodes
// and creates the node mapping of the capacity for all the nodes.
// It returns a map which has nodes as key and capacity provisioned
// on the nodes as corresponding value. The scheduler will use this map
// and picks the node which is less weighted.
func getCapacityWeightedMap(pool string) (map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create the map of the volume capacity
// for the given pool

// getNodeMap returns the node mapping for the given scheduling algorithm
func getNodeMap(schd string, pool string) (map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return CapacityWeighted(default) if not specified
