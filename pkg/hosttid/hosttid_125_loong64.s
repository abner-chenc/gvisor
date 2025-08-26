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

//go:build go1.25

#include "textflag.h"

#define M_OFFSET      48 // +checkoffset runtime g.m
#define PROCID_OFFSET 64 // +checkoffset runtime m.procid

TEXT ·Current(SB),NOSPLIT,$0-8
	// procid is in getg().m.procid.
	MOVV	g, R4
	MOVV	M_OFFSET(R4), R4	// gp.m
	MOVV	PROCID_OFFSET(R4), R4	// mp.procid
	MOVV	R4, ret+0(FP)
	RET
