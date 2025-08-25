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

const (
	// LOONG64FeatureCPUCFG indicates support CPUCFG instruction.
	LOONG64FeatureCPUCFG Feature = iota

	// LOONG64FeatureLAM indicates support AM* atomic memory access instruction.
	LOONG64FeatureLAM

	// LOONG64FeatureFP indicates support for basic floating-point instructions.
	LOONG64FeatureFPU

	// LOONG64FeatureLSX indicates support for 128-bit vector extension.
	LOONG64FeatureLSX

	// LOONG64FeatureLSX indicates support for 256-bit vector extension.
	LOONG64FeatureLASX

	// LOONG64FeatureCRC32 indicates support for CRC32 instructions.
	LOONG64FeatureCRC32

	// LOONG64FeatureCOMPLEX indicates support for complex vector operation instructions.
	LOONG64FeatureCOMPLEX

	// LOONG64FeatureCRYPTO indicates support for encryption and decryption vector instructions.
	LOONG64FeatureCRYPTO

	// LOONG64FeatureLVZ indicates support for virtualization expansion.
	LOONG64FeatureLVZ

	// LOONG64FeatureLBT_X86 indicates support for X86 binary translation extension.
	LOONG64FeatureLBT_X86

	// LOONG64FeatureLBT_ARM indicates support for ARM binary translation extension.
	LOONG64FeatureLBT_ARM

	// LOONG64FeatureLBT_MIPS indicates support for MIPS binary translation extension.
	LOONG64FeatureLBT_MIPS

	// LOONG64FeatureLSPW indicates support for the software page table walking instruction.
	LOONG64FeatureLSPW
)

var allFeatures = map[Feature]allFeatureInfo{
	LOONG64FeatureCPUCFG:   {"cpucfg", true},
	LOONG64FeatureLAM:      {"lam", true},
	LOONG64FeatureFPU:      {"fpu", true},
	LOONG64FeatureLSX:      {"lsx", true},
	LOONG64FeatureLASX:     {"lasx", true},
	LOONG64FeatureCRC32:    {"crc32", true},
	LOONG64FeatureCOMPLEX:  {"complex", false},
	LOONG64FeatureCRYPTO:   {"crypto", true},
	LOONG64FeatureLVZ:      {"lvz", true},
	LOONG64FeatureLBT_X86:  {"lbt_x86", true},
	LOONG64FeatureLBT_ARM:  {"lbt_arm", true},
	LOONG64FeatureLBT_MIPS: {"lbt_mips", true},
	LOONG64FeatureLSPW:     {"lspw", true},
}

func archFlagOrder(fn func(Feature)) {
	for i := 0; i < len(allFeatures); i++ {
		fn(Feature(i))
	}
}
