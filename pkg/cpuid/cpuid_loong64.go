// Copyright 2050 The gVisor Authors.
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
	"io"
)

// FeatureSet for Loong64 is defined as a static set of bits.
//
// Loong64 doesn't have a CPUID equivalent, which means it has no architected
// discovery mechanism for hardware features available to userspace code at
// EL0. The kernel exposes the presence of these features to userspace through
// a set of flags(HWCAP) bits, exposed in the auxiliary vector.
//
// +stateify savable
type FeatureSet struct {
	hwCap      hwCap
	cpuFreqMHz float64
	modelName  string
}

// ExtendedStateSize returns the number of bytes needed to save the "extended
// state" for this processor and the boundary it must be aligned to.
// Extended state includes floating point(NEON) registers, and other cpu state that's not
// associated with the normal task context.
func (fs FeatureSet) ExtendedStateSize() (size, align uint) {
	// Ref arch/loongarch64/include/uapi/asm/ptrace.h struct user_lasx_state
	if fs.HasFeature(LOONG64FeatureLASX) {
		return 1024, 32
	}

	// Ref arch/loongarch64/include/uapi/asm/ptrace.h struct user_lsx_state
	if fs.HasFeature(LOONG64FeatureLSX) {
		return 512, 16
	}

	// Ref arch/loongarch64/include/uapi/asm/ptrace.h struct user_fp_state
	return 140, 8
}

// HasFeature checks for the presence of a feature.
func (fs FeatureSet) HasFeature(feature Feature) bool {
	return fs.hwCap.hwCap1&(1<<feature) != 0
}

// WriteCPUInfoTo is to generate a section of one cpu in /proc/cpuinfo. This is
// a minimal /proc/cpuinfo, and the bogomips field is simply made up.
func (fs FeatureSet) WriteCPUInfoTo(cpu, numCPU uint, w io.Writer) {
	fmt.Fprintf(w, "processor\t: %d\n", cpu)
	fmt.Fprintf(w, "Model Name\t: %s\n", fs.modelName)
	fmt.Fprintf(w, "CPU MHz\t\t: %.02f\n", fs.cpuFreqMHz)
	fmt.Fprintf(w, "Features\t\t: %s\n", fs.FlagString())
	fmt.Fprintf(w, "\n") // The /proc/cpuinfo file ends with an extra newline.
}

// archCheckHostCompatible is a noop on arm64.
func (FeatureSet) archCheckHostCompatible(FeatureSet) error {
	return nil
}

// AllowedHWCap1 returns the HWCAP1 bits that the guest is allowed to depend
// on.
func (fs FeatureSet) AllowedHWCap1() uint64 {
	// Pick a set of safe HWCAPS to expose. These do not rely on cpu state
	// that gvisor does not restore after a context switch.
	allowed := HWCAP_LOONGARCH_CPUCFG |
		HWCAP_LOONGARCH_LAM |
		HWCAP_LOONGARCH_UAL |
		HWCAP_LOONGARCH_FPU |
		HWCAP_LOONGARCH_LSX |
		HWCAP_LOONGARCH_LASX |
		HWCAP_LOONGARCH_CRC32 |
		HWCAP_LOONGARCH_COMPLEX |
		HWCAP_LOONGARCH_CRYPTO |
		HWCAP_LOONGARCH_LVZ |
		HWCAP_LOONGARCH_LBT_X86 |
		HWCAP_LOONGARCH_LBT_ARM |
		HWCAP_LOONGARCH_LBT_MIPS |
		HWCAP_LOONGARCH_PTW |
		HWCAP_LOONGARCH_LSPW

	return fs.hwCap.hwCap1 & uint64(allowed)
}

// AllowedHWCap2 returns the HWCAP2 bits that the guest is allowed to depend
// on.
func (fs FeatureSet) AllowedHWCap2() uint64 {
	// HWCAPS are not supported on loong64.
	return 0
}
