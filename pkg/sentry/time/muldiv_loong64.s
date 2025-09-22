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

#include "funcdata.h"
#include "textflag.h"

// Documentation is available in parameters.go.
//
// func muldiv64(value, multiplier, divisor uint64) (uint64, bool)
TEXT ·muldiv64(SB),NOSPLIT,$40-33
	GO_ARGS
	NO_LOCAL_POINTERS
	MOVV	value+0(FP), R4
	MOVV	multiplier+8(FP), R5
	MOVV	divisor+16(FP), R6

	MULHVU	R4, R5, R7
	MULVU	R4, R5, R4
	// if R6 >= R7 then overflow
	BGEU	R6, R7, overflow

	MOVV	R7, 8(R3)
	MOVV	R4, 16(R3)
	MOVV	R6, 24(R3)
	CALL	·divWW(SB)
	MOVV	32(R3), R4
	MOVV	R4, ret+24(FP)
	MOVV	$1, R4
	MOVB	R4, ret1+32(FP)
	RET
overflow:
	MOVV	R0, ret+24(FP)
	MOVB	R0, ret1+32(FP)
	RET
