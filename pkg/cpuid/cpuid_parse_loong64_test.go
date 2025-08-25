// Copyright 2020 The gVisor Authors.
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

import "gvisor.dev/gvisor/pkg/hostos"

func archSkipFeature(feature Feature, version hostos.Version) bool {
	switch {
	case feature == HWCAP_LOONGARCH_COMPLEX:
		return true
	case feature == LOONG64FeatureCRYPTO:
		return true
	case feature == LOONG64FeatureLBT_X86:
		return true
	case feature == LOONG64FeatureLBT_ARM:
		return true
	case feature == LOONG64FeatureLBT_MIPS:
		return true
	default:
		return false
	}
}
