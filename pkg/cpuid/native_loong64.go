// Copyright 2025 The gVisor Authors.
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

//go:build loong64
// +build loong64

package cpuid

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"gvisor.dev/gvisor/pkg/log"
)

// hostFeatureSet is initialized at startup.
//
// This is copied for HostFeatureSet, below.
var hostFeatureSet FeatureSet

// HostFeatureSet returns a copy of the host FeatureSet.
func HostFeatureSet() FeatureSet {
	return hostFeatureSet
}

// Fixed returns the same feature set.
func (fs FeatureSet) Fixed() FeatureSet {
	return fs
}

// Intersect returns the intersection of features between self and allowedFeatures.
//
// Just return error as there is no ARM64 equivalent to cpuid.Static.Remove().
func (fs FeatureSet) Intersect(allowedFeatures map[Feature]struct{}) (FeatureSet, error) {
	return FeatureSet{}, fmt.Errorf("FeatureSet intersection is not supported on ARM64")
}

// Reads CPU information from host /proc/cpuinfo.
//
// Must run before syscall filter installation. This value is used to create
// the fake /proc/cpuinfo from a FeatureSet.
func initCPUInfo() {
	if runtime.GOOS != "linux" {
		// Don't try to read Linux-specific /proc files or
		// warn about them not existing.
		return
	}
	cpuinfob, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		// Leave everything at 0, nothing can be done.
		log.Warningf("Could not read /proc/cpuinfo: %v", err)
		return
	}
	cpuinfo := string(cpuinfob)

	// We get the value straight from host /proc/cpuinfo.
	for _, line := range strings.Split(cpuinfo, "\n") {
		switch {
		case strings.Contains(line, "CPU MHz"):
			splitMHz := strings.Split(line, ":")
			if len(splitMHz) < 2 {
				log.Warningf("Could not read /proc/cpuinfo: malformed CPU MHz")
				break
			}

			// If there was a problem, leave cpuFreqMHz as 0.
			var err error
			hostFeatureSet.cpuFreqMHz, err = strconv.ParseFloat(strings.TrimSpace(splitMHz[1]), 64)
			if err != nil {
				hostFeatureSet.cpuFreqMHz = 0.0
				log.Warningf("Could not parse CPU MHz value %v: %v", splitMHz[1], err)
			}

		case strings.Contains(line, "Model Name"):
			splitModelName := strings.Split(line, ":")
			if len(splitModelName) < 2 {
				log.Warningf("Could not read /proc/cpuinfo: malformed Model Name")
				break
			}

			hostFeatureSet.modelName = splitModelName[1]
		}
	}
}

// archInitialize initializes hostFeatureSet.
func archInitialize() {
	initCPUInfo()
	initHWCap()
}
