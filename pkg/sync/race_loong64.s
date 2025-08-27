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

//go:build race && loong64
// +build race,loong64

#include "textflag.h"

// func RaceUncheckedAtomicCompareAndSwapUintptr(ptr *uintptr, old, new uintptr) bool
TEXT ·RaceUncheckedAtomicCompareAndSwapUintptr(SB),NOSPLIT,$0-25
	// refs: src/internal/runtime/atomic/atomic_loong64.s
	// Implemented using the ll-sc instruction pair
	MOVV	ptr+0(FP), R4
	MOVV	old+8(FP), R5
	MOVV	new+16(FP), R6
	DBAR	$0x14
cas64_again:
	MOVV	R6, R7
	LLV	(R4), R8
	BNE	R5, R8, cas64_fail1
	SCV	R7, (R4)
	BEQ	R7, cas64_again
	MOVV	$1, R4
	MOVB	R4, ret+24(FP)
	DBAR	$0x12
	RET
cas64_fail1:
	MOVV	$0, R4
	JMP	-4(PC)
