/*
Copyright 2011-2013 Frederic Langlet
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
you may obtain a copy of the License at

                http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package transform

import (
	"kanzi"
)

const (
	MAX_VAL32 = 1 << 16
	MIN_VAL32 = -(MAX_VAL32 + 1)
)

var w = []int{
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	90, 90, 88, 85, 82, 78, 73, 67, 61, 54, 46, 38, 31, 22, 13, 4,
	-4, -13, -22, -31, -38, -46, -54, -61, -67, -73, -78, -82, -85, -88, -90, -90,
	90, 87, 80, 70, 57, 43, 25, 9, -9, -25, -43, -57, -70, -80, -87, -90,
	-90, -87, -80, -70, -57, -43, -25, -9, 9, 25, 43, 57, 70, 80, 87, 90,
	90, 82, 67, 46, 22, -4, -31, -54, -73, -85, -90, -88, -78, -61, -38, -13,
	13, 38, 61, 78, 88, 90, 85, 73, 54, 31, 4, -22, -46, -67, -82, -90,
	89, 75, 50, 18, -18, -50, -75, -89, -89, -75, -50, -18, 18, 50, 75, 89,
	89, 75, 50, 18, -18, -50, -75, -89, -89, -75, -50, -18, 18, 50, 75, 89,
	88, 67, 31, -13, -54, -82, -90, -78, -46, -4, 38, 73, 90, 85, 61, 22,
	-22, -61, -85, -90, -73, -38, 4, 46, 78, 90, 82, 54, 13, -31, -67, -88,
	87, 57, 9, -43, -80, -90, -70, -25, 25, 70, 90, 80, 43, -9, -57, -87,
	-87, -57, -9, 43, 80, 90, 70, 25, -25, -70, -90, -80, -43, 9, 57, 87,
	85, 46, -13, -67, -90, -73, -22, 38, 82, 88, 54, -4, -61, -90, -78, -31,
	31, 78, 90, 61, 4, -54, -88, -82, -38, 22, 73, 90, 67, 13, -46, -85,
	83, 36, -36, -83, -83, -36, 36, 83, 83, 36, -36, -83, -83, -36, 36, 83,
	83, 36, -36, -83, -83, -36, 36, 83, 83, 36, -36, -83, -83, -36, 36, 83,
	82, 22, -54, -90, -61, 13, 78, 85, 31, -46, -90, -67, 4, 73, 88, 38,
	-38, -88, -73, -4, 67, 90, 46, -31, -85, -78, -13, 61, 90, 54, -22, -82,
	80, 9, -70, -87, -25, 57, 90, 43, -43, -90, -57, 25, 87, 70, -9, -80,
	-80, -9, 70, 87, 25, -57, -90, -43, 43, 90, 57, -25, -87, -70, 9, 80,
	78, -4, -82, -73, 13, 85, 67, -22, -88, -61, 31, 90, 54, -38, -90, -46,
	46, 90, 38, -54, -90, -31, 61, 88, 22, -67, -85, -13, 73, 82, 4, -78,
	75, -18, -89, -50, 50, 89, 18, -75, -75, 18, 89, 50, -50, -89, -18, 75,
	75, -18, -89, -50, 50, 89, 18, -75, -75, 18, 89, 50, -50, -89, -18, 75,
	73, -31, -90, -22, 78, 67, -38, -90, -13, 82, 61, -46, -88, -4, 85, 54,
	-54, -85, 4, 88, 46, -61, -82, 13, 90, 38, -67, -78, 22, 90, 31, -73,
	70, -43, -87, 9, 90, 25, -80, -57, 57, 80, -25, -90, -9, 87, 43, -70,
	-70, 43, 87, -9, -90, -25, 80, 57, -57, -80, 25, 90, 9, -87, -43, 70,
	67, -54, -78, 38, 85, -22, -90, 4, 90, 13, -88, -31, 82, 46, -73, -61,
	61, 73, -46, -82, 31, 88, -13, -90, -4, 90, 22, -85, -38, 78, 54, -67,
	64, -64, -64, 64, 64, -64, -64, 64, 64, -64, -64, 64, 64, -64, -64, 64,
	64, -64, -64, 64, 64, -64, -64, 64, 64, -64, -64, 64, 64, -64, -64, 64,
	61, -73, -46, 82, 31, -88, -13, 90, -4, -90, 22, 85, -38, -78, 54, 67,
	-67, -54, 78, 38, -85, -22, 90, 4, -90, 13, 88, -31, -82, 46, 73, -61,
	57, -80, -25, 90, -9, -87, 43, 70, -70, -43, 87, 9, -90, 25, 80, -57,
	-57, 80, 25, -90, 9, 87, -43, -70, 70, 43, -87, -9, 90, -25, -80, 57,
	54, -85, -4, 88, -46, -61, 82, 13, -90, 38, 67, -78, -22, 90, -31, -73,
	73, 31, -90, 22, 78, -67, -38, 90, -13, -82, 61, 46, -88, 4, 85, -54,
	50, -89, 18, 75, -75, -18, 89, -50, -50, 89, -18, -75, 75, 18, -89, 50,
	50, -89, 18, 75, -75, -18, 89, -50, -50, 89, -18, -75, 75, 18, -89, 50,
	46, -90, 38, 54, -90, 31, 61, -88, 22, 67, -85, 13, 73, -82, 4, 78,
	-78, -4, 82, -73, -13, 85, -67, -22, 88, -61, -31, 90, -54, -38, 90, -46,
	43, -90, 57, 25, -87, 70, 9, -80, 80, -9, -70, 87, -25, -57, 90, -43,
	-43, 90, -57, -25, 87, -70, -9, 80, -80, 9, 70, -87, 25, 57, -90, 43,
	38, -88, 73, -4, -67, 90, -46, -31, 85, -78, 13, 61, -90, 54, 22, -82,
	82, -22, -54, 90, -61, -13, 78, -85, 31, 46, -90, 67, 4, -73, 88, -38,
	36, -83, 83, -36, -36, 83, -83, 36, 36, -83, 83, -36, -36, 83, -83, 36,
	36, -83, 83, -36, -36, 83, -83, 36, 36, -83, 83, -36, -36, 83, -83, 36,
	31, -78, 90, -61, 4, 54, -88, 82, -38, -22, 73, -90, 67, -13, -46, 85,
	-85, 46, 13, -67, 90, -73, 22, 38, -82, 88, -54, -4, 61, -90, 78, -31,
	25, -70, 90, -80, 43, 9, -57, 87, -87, 57, -9, -43, 80, -90, 70, -25,
	-25, 70, -90, 80, -43, -9, 57, -87, 87, -57, 9, 43, -80, 90, -70, 25,
	22, -61, 85, -90, 73, -38, -4, 46, -78, 90, -82, 54, -13, -31, 67, -88,
	88, -67, 31, 13, -54, 82, -90, 78, -46, 4, 38, -73, 90, -85, 61, -22,
	18, -50, 75, -89, 89, -75, 50, -18, -18, 50, -75, 89, -89, 75, -50, 18,
	18, -50, 75, -89, 89, -75, 50, -18, -18, 50, -75, 89, -89, 75, -50, 18,
	13, -38, 61, -78, 88, -90, 85, -73, 54, -31, 4, 22, -46, 67, -82, 90,
	-90, 82, -67, 46, -22, -4, 31, -54, 73, -85, 90, -88, 78, -61, 38, -13,
	9, -25, 43, -57, 70, -80, 87, -90, 90, -87, 80, -70, 57, -43, 25, -9,
	-9, 25, -43, 57, -70, 80, -87, 90, -90, 87, -80, 70, -57, 43, -25, 9,
	4, -13, 22, -31, 38, -46, 54, -61, 67, -73, 78, -82, 85, -88, 90, -90,
	90, -90, 88, -85, 82, -78, 73, -67, 61, -54, 46, -38, 31, -22, 13, -4,
}

type DCT32 struct {
	fShift uint  // default 14
	iShift uint  // default 20
	data   []int // int[1024]
}

func NewDCT32() (*DCT32, error) {
	this := new(DCT32)
	this.fShift = 14
	this.iShift = 20
	this.data = make([]int, 1024)
	return this, nil
}

func (this *DCT32) Forward(src, dst []int) (uint, uint, error) {
	computeForward32(src, this.data, 7)
	computeForward32(this.data, dst, this.fShift-7)
	return 1024, 1024, nil
}

func computeForward32(input, output []int, shift uint) {
	iIdx := 0
	round := (1 << shift) >> 1

	for i := 0; i < 32; i++ {
		x0 := input[iIdx]
		x1 := input[iIdx+1]
		x2 := input[iIdx+2]
		x3 := input[iIdx+3]
		x4 := input[iIdx+4]
		x5 := input[iIdx+5]
		x6 := input[iIdx+6]
		x7 := input[iIdx+7]
		x8 := input[iIdx+8]
		x9 := input[iIdx+9]
		x10 := input[iIdx+10]
		x11 := input[iIdx+11]
		x12 := input[iIdx+12]
		x13 := input[iIdx+13]
		x14 := input[iIdx+14]
		x15 := input[iIdx+15]
		x16 := input[iIdx+16]
		x17 := input[iIdx+17]
		x18 := input[iIdx+18]
		x19 := input[iIdx+19]
		x20 := input[iIdx+20]
		x21 := input[iIdx+21]
		x22 := input[iIdx+22]
		x23 := input[iIdx+23]
		x24 := input[iIdx+24]
		x25 := input[iIdx+25]
		x26 := input[iIdx+26]
		x27 := input[iIdx+27]
		x28 := input[iIdx+28]
		x29 := input[iIdx+29]
		x30 := input[iIdx+30]
		x31 := input[iIdx+31]

		a0 := x0 + x31
		a1 := x1 + x30
		a2 := x0 - x31
		a3 := x1 - x30
		a4 := x2 + x29
		a5 := x3 + x28
		a6 := x2 - x29
		a7 := x3 - x28
		a8 := x4 + x27
		a9 := x5 + x26
		a10 := x4 - x27
		a11 := x5 - x26
		a12 := x6 + x25
		a13 := x7 + x24
		a14 := x6 - x25
		a15 := x7 - x24
		a16 := x8 + x23
		a17 := x9 + x22
		a18 := x8 - x23
		a19 := x9 - x22
		a20 := x10 + x21
		a21 := x11 + x20
		a22 := x10 - x21
		a23 := x11 - x20
		a24 := x12 + x19
		a25 := x13 + x18
		a26 := x12 - x19
		a27 := x13 - x18
		a28 := x14 + x17
		a29 := x15 + x16
		a30 := x14 - x17
		a31 := x15 - x16

		for n := 32; n < 1024; n += 64 {
			output[i+n] = ((w[n] * a2) + (w[n+1] * a3) + (w[n+2] * a6) + (w[n+3] * a7) +
				(w[n+4] * a10) + (w[n+5] * a11) + (w[n+6] * a14) + (w[n+7] * a15) +
				(w[n+8] * a18) + (w[n+9] * a19) + (w[n+10] * a22) + (w[n+11] * a23) +
				(w[n+12] * a26) + (w[n+13] * a27) + (w[n+14] * a30) + (w[n+15] * a31) + round) >> shift
		}

		b0 := a0 + a29
		b1 := a1 + a28
		b2 := a0 - a29
		b3 := a1 - a28
		b4 := a4 + a25
		b5 := a5 + a24
		b6 := a4 - a25
		b7 := a5 - a24
		b8 := a8 + a21
		b9 := a9 + a20
		b10 := a8 - a21
		b11 := a9 - a20
		b12 := a12 + a17
		b13 := a13 + a16
		b14 := a12 - a17
		b15 := a13 - a16

		output[i+64] = ((w[64] * b2) + (w[65] * b3) + (w[66] * b6) + (w[67] * b7) +
			(w[68] * b10) + (w[69] * b11) + (w[70] * b14) + (w[71] * b15) + round) >> shift
		output[i+192] = ((w[192] * b2) + (w[193] * b3) + (w[194] * b6) + (w[195] * b7) +
			(w[196] * b10) + (w[197] * b11) + (w[198] * b14) + (w[199] * b15) + round) >> shift
		output[i+320] = ((w[320] * b2) + (w[321] * b3) + (w[322] * b6) + (w[323] * b7) +
			(w[324] * b10) + (w[325] * b11) + (w[326] * b14) + (w[327] * b15) + round) >> shift
		output[i+448] = ((w[448] * b2) + (w[449] * b3) + (w[450] * b6) + (w[451] * b7) +
			(w[452] * b10) + (w[453] * b11) + (w[454] * b14) + (w[455] * b15) + round) >> shift
		output[i+576] = ((w[576] * b2) + (w[577] * b3) + (w[578] * b6) + (w[579] * b7) +
			(w[580] * b10) + (w[581] * b11) + (w[582] * b14) + (w[583] * b15) + round) >> shift
		output[i+704] = ((w[704] * b2) + (w[705] * b3) + (w[706] * b6) + (w[707] * b7) +
			(w[708] * b10) + (w[709] * b11) + (w[710] * b14) + (w[711] * b15) + round) >> shift
		output[i+832] = ((w[832] * b2) + (w[833] * b3) + (w[834] * b6) + (w[835] * b7) +
			(w[836] * b10) + (w[837] * b11) + (w[838] * b14) + (w[839] * b15) + round) >> shift
		output[i+960] = ((w[960] * b2) + (w[961] * b3) + (w[962] * b6) + (w[963] * b7) +
			(w[964] * b10) + (w[965] * b11) + (w[966] * b14) + (w[967] * b15) + round) >> shift

		c0 := b0 + b13
		c1 := b1 + b12
		c2 := b0 - b13
		c3 := b1 - b12
		c4 := b4 + b9
		c5 := b5 + b8
		c6 := b4 - b9
		c7 := b5 - b8

		output[i+128] = ((w[128] * c2) + (w[129] * c3) + (w[130] * c6) + (w[131] * c7) + round) >> shift
		output[i+384] = ((w[384] * c2) + (w[385] * c3) + (w[386] * c6) + (w[387] * c7) + round) >> shift
		output[i+640] = ((w[640] * c2) + (w[641] * c3) + (w[642] * c6) + (w[643] * c7) + round) >> shift
		output[i+896] = ((w[896] * c2) + (w[897] * c3) + (w[898] * c6) + (w[899] * c7) + round) >> shift

		d0 := c0 + c5
		d1 := c1 + c4
		d2 := c0 - c5
		d3 := c1 - c4

		output[i] = ((w[0] * d0) + (w[1] * d1) + round) >> shift
		output[i+512] = ((w[512] * d0) + (w[513] * d1) + round) >> shift
		output[i+256] = ((w[256] * d2) + (w[257] * d3) + round) >> shift
		output[i+768] = ((w[768] * d2) + (w[769] * d3) + round) >> shift

		iIdx += 32
	}

}

func (this *DCT32) Inverse(src, dst []int) (uint, uint, error) {
	computeInverse32(src, this.data, 10)
	computeInverse32(this.data, dst, this.iShift-10)
	return 1024, 1024, nil
}

func computeInverse32(input, output []int, shift uint) {
	oIdx := 0
	round := (1 << shift) >> 1

	for i := 0; i < 32; i++ {
		x0 := input[i]
		x1 := input[i+32]
		x2 := input[i+64]
		x3 := input[i+96]
		x4 := input[i+128]
		x5 := input[i+160]
		x6 := input[i+192]
		x7 := input[i+224]
		x8 := input[i+256]
		x9 := input[i+288]
		x10 := input[i+320]
		x11 := input[i+352]
		x12 := input[i+384]
		x13 := input[i+416]
		x14 := input[i+448]
		x15 := input[i+480]
		x16 := input[i+512]
		x17 := input[i+544]
		x18 := input[i+576]
		x19 := input[i+608]
		x20 := input[i+640]
		x21 := input[i+672]
		x22 := input[i+704]
		x23 := input[i+736]
		x24 := input[i+768]
		x25 := input[i+800]
		x26 := input[i+832]
		x27 := input[i+864]
		x28 := input[i+896]
		x29 := input[i+928]
		x30 := input[i+960]
		x31 := input[i+992]

		a0 := (w[32] * x1) + (w[96] * x3) + (w[160] * x5) + (w[224] * x7) +
			(w[288] * x9) + (w[352] * x11) + (w[416] * x13) + (w[480] * x15) +
			(w[544] * x17) + (w[608] * x19) + (w[672] * x21) + (w[736] * x23) +
			(w[800] * x25) + (w[864] * x27) + (w[928] * x29) + (w[992] * x31)
		a1 := (w[33] * x1) + (w[97] * x3) + (w[161] * x5) + (w[225] * x7) +
			(w[289] * x9) + (w[353] * x11) + (w[417] * x13) + (w[481] * x15) +
			(w[545] * x17) + (w[609] * x19) + (w[673] * x21) + (w[737] * x23) +
			(w[801] * x25) + (w[865] * x27) + (w[929] * x29) + (w[993] * x31)
		a2 := (w[34] * x1) + (w[98] * x3) + (w[162] * x5) + (w[226] * x7) +
			(w[290] * x9) + (w[354] * x11) + (w[418] * x13) + (w[482] * x15) +
			(w[546] * x17) + (w[610] * x19) + (w[674] * x21) + (w[738] * x23) +
			(w[802] * x25) + (w[866] * x27) + (w[930] * x29) + (w[994] * x31)
		a3 := (w[35] * x1) + (w[99] * x3) + (w[163] * x5) + (w[227] * x7) +
			(w[291] * x9) + (w[355] * x11) + (w[419] * x13) + (w[483] * x15) +
			(w[547] * x17) + (w[611] * x19) + (w[675] * x21) + (w[739] * x23) +
			(w[803] * x25) + (w[867] * x27) + (w[931] * x29) + (w[995] * x31)
		a4 := (w[36] * x1) + (w[100] * x3) + (w[164] * x5) + (w[228] * x7) +
			(w[292] * x9) + (w[356] * x11) + (w[420] * x13) + (w[484] * x15) +
			(w[548] * x17) + (w[612] * x19) + (w[676] * x21) + (w[740] * x23) +
			(w[804] * x25) + (w[868] * x27) + (w[932] * x29) + (w[996] * x31)
		a5 := (w[37] * x1) + (w[101] * x3) + (w[165] * x5) + (w[229] * x7) +
			(w[293] * x9) + (w[357] * x11) + (w[421] * x13) + (w[485] * x15) +
			(w[549] * x17) + (w[613] * x19) + (w[677] * x21) + (w[741] * x23) +
			(w[805] * x25) + (w[869] * x27) + (w[933] * x29) + (w[997] * x31)
		a6 := (w[38] * x1) + (w[102] * x3) + (w[166] * x5) + (w[230] * x7) +
			(w[294] * x9) + (w[358] * x11) + (w[422] * x13) + (w[486] * x15) +
			(w[550] * x17) + (w[614] * x19) + (w[678] * x21) + (w[742] * x23) +
			(w[806] * x25) + (w[870] * x27) + (w[934] * x29) + (w[998] * x31)
		a7 := (w[39] * x1) + (w[103] * x3) + (w[167] * x5) + (w[231] * x7) +
			(w[295] * x9) + (w[359] * x11) + (w[423] * x13) + (w[487] * x15) +
			(w[551] * x17) + (w[615] * x19) + (w[679] * x21) + (w[743] * x23) +
			(w[807] * x25) + (w[871] * x27) + (w[935] * x29) + (w[999] * x31)
		a8 := (w[40] * x1) + (w[104] * x3) + (w[168] * x5) + (w[232] * x7) +
			(w[296] * x9) + (w[360] * x11) + (w[424] * x13) + (w[488] * x15) +
			(w[552] * x17) + (w[616] * x19) + (w[680] * x21) + (w[744] * x23) +
			(w[808] * x25) + (w[872] * x27) + (w[936] * x29) + (w[1000] * x31)
		a9 := (w[41] * x1) + (w[105] * x3) + (w[169] * x5) + (w[233] * x7) +
			(w[297] * x9) + (w[361] * x11) + (w[425] * x13) + (w[489] * x15) +
			(w[553] * x17) + (w[617] * x19) + (w[681] * x21) + (w[745] * x23) +
			(w[809] * x25) + (w[873] * x27) + (w[937] * x29) + (w[1001] * x31)
		a10 := (w[42] * x1) + (w[106] * x3) + (w[170] * x5) + (w[234] * x7) +
			(w[298] * x9) + (w[362] * x11) + (w[426] * x13) + (w[490] * x15) +
			(w[554] * x17) + (w[618] * x19) + (w[682] * x21) + (w[746] * x23) +
			(w[810] * x25) + (w[874] * x27) + (w[938] * x29) + (w[1002] * x31)
		a11 := (w[43] * x1) + (w[107] * x3) + (w[171] * x5) + (w[235] * x7) +
			(w[299] * x9) + (w[363] * x11) + (w[427] * x13) + (w[491] * x15) +
			(w[555] * x17) + (w[619] * x19) + (w[683] * x21) + (w[747] * x23) +
			(w[811] * x25) + (w[875] * x27) + (w[939] * x29) + (w[1003] * x31)
		a12 := (w[44] * x1) + (w[108] * x3) + (w[172] * x5) + (w[236] * x7) +
			(w[300] * x9) + (w[364] * x11) + (w[428] * x13) + (w[492] * x15) +
			(w[556] * x17) + (w[620] * x19) + (w[684] * x21) + (w[748] * x23) +
			(w[812] * x25) + (w[876] * x27) + (w[940] * x29) + (w[1004] * x31)
		a13 := (w[45] * x1) + (w[109] * x3) + (w[173] * x5) + (w[237] * x7) +
			(w[301] * x9) + (w[365] * x11) + (w[429] * x13) + (w[493] * x15) +
			(w[557] * x17) + (w[621] * x19) + (w[685] * x21) + (w[749] * x23) +
			(w[813] * x25) + (w[877] * x27) + (w[941] * x29) + (w[1005] * x31)
		a14 := (w[46] * x1) + (w[110] * x3) + (w[174] * x5) + (w[238] * x7) +
			(w[302] * x9) + (w[366] * x11) + (w[430] * x13) + (w[494] * x15) +
			(w[558] * x17) + (w[622] * x19) + (w[686] * x21) + (w[750] * x23) +
			(w[814] * x25) + (w[878] * x27) + (w[942] * x29) + (w[1006] * x31)
		a15 := (w[47] * x1) + (w[111] * x3) + (w[175] * x5) + (w[239] * x7) +
			(w[303] * x9) + (w[367] * x11) + (w[431] * x13) + (w[495] * x15) +
			(w[559] * x17) + (w[623] * x19) + (w[687] * x21) + (w[751] * x23) +
			(w[815] * x25) + (w[879] * x27) + (w[943] * x29) + (w[1007] * x31)

		b0 := (w[64] * x2) + (w[192] * x6) + (w[320] * x10) + (w[448] * x14) +
			(w[576] * x18) + (w[704] * x22) + (w[832] * x26) + (w[960] * x30)
		b1 := (w[65] * x2) + (w[193] * x6) + (w[321] * x10) + (w[449] * x14) +
			(w[577] * x18) + (w[705] * x22) + (w[833] * x26) + (w[961] * x30)
		b2 := (w[66] * x2) + (w[194] * x6) + (w[322] * x10) + (w[450] * x14) +
			(w[578] * x18) + (w[706] * x22) + (w[834] * x26) + (w[962] * x30)
		b3 := (w[67] * x2) + (w[195] * x6) + (w[323] * x10) + (w[451] * x14) +
			(w[579] * x18) + (w[707] * x22) + (w[835] * x26) + (w[963] * x30)
		b4 := (w[68] * x2) + (w[196] * x6) + (w[324] * x10) + (w[452] * x14) +
			(w[580] * x18) + (w[708] * x22) + (w[836] * x26) + (w[964] * x30)
		b5 := (w[69] * x2) + (w[197] * x6) + (w[325] * x10) + (w[453] * x14) +
			(w[581] * x18) + (w[709] * x22) + (w[837] * x26) + (w[965] * x30)
		b6 := (w[70] * x2) + (w[198] * x6) + (w[326] * x10) + (w[454] * x14) +
			(w[582] * x18) + (w[710] * x22) + (w[838] * x26) + (w[966] * x30)
		b7 := (w[71] * x2) + (w[199] * x6) + (w[327] * x10) + (w[455] * x14) +
			(w[583] * x18) + (w[711] * x22) + (w[839] * x26) + (w[967] * x30)

		c0 := (w[128] * x4) + (w[384] * x12) + (w[640] * x20) + (w[896] * x28)
		c1 := (w[129] * x4) + (w[385] * x12) + (w[641] * x20) + (w[897] * x28)
		c2 := (w[130] * x4) + (w[386] * x12) + (w[642] * x20) + (w[898] * x28)
		c3 := (w[131] * x4) + (w[387] * x12) + (w[643] * x20) + (w[899] * x28)
		c4 := (w[256] * x8) + (w[768] * x24)
		c5 := (w[257] * x8) + (w[769] * x24)
		c6 := (w[0] * x0) + (w[512] * x16)
		c7 := (w[1] * x0) + (w[513] * x16)
		c8 := c6 + c4
		c9 := c7 + c5
		c10 := c7 - c5
		c11 := c6 - c4

		d0 := c8 + c0
		d1 := c9 + c1
		d2 := c10 + c2
		d3 := c11 + c3
		d4 := c11 - c3
		d5 := c10 - c2
		d6 := c9 - c1
		d7 := c8 - c0

		e0 := d0 + b0
		e1 := d1 + b1
		e2 := d2 + b2
		e3 := d3 + b3
		e4 := d4 + b4
		e5 := d5 + b5
		e6 := d6 + b6
		e7 := d7 + b7
		e8 := d7 - b7
		e9 := d6 - b6
		e10 := d5 - b5
		e11 := d4 - b4
		e12 := d3 - b3
		e13 := d2 - b2
		e14 := d1 - b1
		e15 := d0 - b0

		r0 := (e0 + a0 + round) >> shift
		r1 := (e1 + a1 + round) >> shift
		r2 := (e2 + a2 + round) >> shift
		r3 := (e3 + a3 + round) >> shift
		r4 := (e4 + a4 + round) >> shift
		r5 := (e5 + a5 + round) >> shift
		r6 := (e6 + a6 + round) >> shift
		r7 := (e7 + a7 + round) >> shift
		r8 := (e8 + a8 + round) >> shift
		r9 := (e9 + a9 + round) >> shift
		r10 := (e10 + a10 + round) >> shift
		r11 := (e11 + a11 + round) >> shift
		r12 := (e12 + a12 + round) >> shift
		r13 := (e13 + a13 + round) >> shift
		r14 := (e14 + a14 + round) >> shift
		r15 := (e15 + a15 + round) >> shift
		r16 := (e15 - a15 + round) >> shift
		r17 := (e14 - a14 + round) >> shift
		r18 := (e13 - a13 + round) >> shift
		r19 := (e12 - a12 + round) >> shift
		r20 := (e11 - a11 + round) >> shift
		r21 := (e10 - a10 + round) >> shift
		r22 := (e9 - a9 + round) >> shift
		r23 := (e8 - a8 + round) >> shift
		r24 := (e7 - a7 + round) >> shift
		r25 := (e6 - a6 + round) >> shift
		r26 := (e5 - a5 + round) >> shift
		r27 := (e4 - a4 + round) >> shift
		r28 := (e3 - a3 + round) >> shift
		r29 := (e2 - a2 + round) >> shift
		r30 := (e1 - a1 + round) >> shift
		r31 := (e0 - a0 + round) >> shift

		output[oIdx] = kanzi.Clamp(r0, MIN_VAL32, MAX_VAL32)
		output[oIdx+1] = kanzi.Clamp(r1, MIN_VAL32, MAX_VAL32)
		output[oIdx+2] = kanzi.Clamp(r2, MIN_VAL32, MAX_VAL32)
		output[oIdx+3] = kanzi.Clamp(r3, MIN_VAL32, MAX_VAL32)
		output[oIdx+4] = kanzi.Clamp(r4, MIN_VAL32, MAX_VAL32)
		output[oIdx+5] = kanzi.Clamp(r5, MIN_VAL32, MAX_VAL32)
		output[oIdx+6] = kanzi.Clamp(r6, MIN_VAL32, MAX_VAL32)
		output[oIdx+7] = kanzi.Clamp(r7, MIN_VAL32, MAX_VAL32)
		output[oIdx+8] = kanzi.Clamp(r8, MIN_VAL32, MAX_VAL32)
		output[oIdx+9] = kanzi.Clamp(r9, MIN_VAL32, MAX_VAL32)
		output[oIdx+10] = kanzi.Clamp(r10, MIN_VAL32, MAX_VAL32)
		output[oIdx+11] = kanzi.Clamp(r11, MIN_VAL32, MAX_VAL32)
		output[oIdx+12] = kanzi.Clamp(r12, MIN_VAL32, MAX_VAL32)
		output[oIdx+13] = kanzi.Clamp(r13, MIN_VAL32, MAX_VAL32)
		output[oIdx+14] = kanzi.Clamp(r14, MIN_VAL32, MAX_VAL32)
		output[oIdx+15] = kanzi.Clamp(r15, MIN_VAL32, MAX_VAL32)
		output[oIdx+16] = kanzi.Clamp(r16, MIN_VAL32, MAX_VAL32)
		output[oIdx+17] = kanzi.Clamp(r17, MIN_VAL32, MAX_VAL32)
		output[oIdx+18] = kanzi.Clamp(r18, MIN_VAL32, MAX_VAL32)
		output[oIdx+19] = kanzi.Clamp(r19, MIN_VAL32, MAX_VAL32)
		output[oIdx+20] = kanzi.Clamp(r20, MIN_VAL32, MAX_VAL32)
		output[oIdx+21] = kanzi.Clamp(r21, MIN_VAL32, MAX_VAL32)
		output[oIdx+22] = kanzi.Clamp(r22, MIN_VAL32, MAX_VAL32)
		output[oIdx+23] = kanzi.Clamp(r23, MIN_VAL32, MAX_VAL32)
		output[oIdx+24] = kanzi.Clamp(r24, MIN_VAL32, MAX_VAL32)
		output[oIdx+25] = kanzi.Clamp(r25, MIN_VAL32, MAX_VAL32)
		output[oIdx+26] = kanzi.Clamp(r26, MIN_VAL32, MAX_VAL32)
		output[oIdx+27] = kanzi.Clamp(r27, MIN_VAL32, MAX_VAL32)
		output[oIdx+28] = kanzi.Clamp(r28, MIN_VAL32, MAX_VAL32)
		output[oIdx+29] = kanzi.Clamp(r29, MIN_VAL32, MAX_VAL32)
		output[oIdx+30] = kanzi.Clamp(r30, MIN_VAL32, MAX_VAL32)
		output[oIdx+31] = kanzi.Clamp(r31, MIN_VAL32, MAX_VAL32)

		oIdx += 32
	}
}
