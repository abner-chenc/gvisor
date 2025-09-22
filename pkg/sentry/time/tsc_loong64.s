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

TEXT ·Rdtsc(SB),NOSPLIT,$0-8
	RDTIMED	R0, R4
	MOVD	R4, ret+0(FP)
	RET

TEXT ·getCNTFRQ(SB),NOSPLIT,$0-8
	// CC_FREQ: [31:0]
	MOVV	$4, R4
	CPUCFG	R4, R5

	// CC_MUL:[15:0], CC_DIV:[31:16]
	MOVV	$5, R4
	CPUCFG	R4, R6

	// CNTFREQ = CC_FREQ * CC_MUL / CC_DIV
	SRLV	$16, R6, R4
	AND	$0xffff, R4
	AND	$0xffff, R6
	MULVU	R5, R4, R4
	DIVVU	R4, R6, R4
	MOVV	R4, ret+0(FP)
	RET
