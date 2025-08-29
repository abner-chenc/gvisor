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

#include "textflag.h"

// GetCPU makes the getcpu(unsigned *cpu, unsigned *node, NULL) syscall for
// the lack of an optimized way of getting the current CPU number on arm64.

// func GetCPU() uint32
TEXT ·GetCPU(SB), NOSPLIT, $0-4
	MOVW	R0, ret+0(FP)
	MOVV	$ret+0(FP), R4
	MOVV	R0, R5	// unused
	MOVV	R0, R6	// unused
	MOVV	$0xA8, R11	// SYS_GETCPU
	SYSCALL
	RET
