// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

const (
	WordLen   = 32 / bitsInByte  // length of word in tl is 32 bits
	LongLen   = 64 / bitsInByte  // int64
	DoubleLen = 64 / bitsInByte  // float64
	Int128Len = 128 / bitsInByte // int128
	Int256Len = 256 / bitsInByte // int256

	// Блядские магические числа
	FuckingMagicNumber = 0xfe // 253 элемента максимум можно закодировать в массиве элементов

	// https://core.telegram.org/schema/mtproto
	CrcVector = 0x1cb5c415
	CrcFalse  = 0xbc799737
	CrcTrue   = 0x997275b5
	CrcNull   = 0x56730bcc

	bitsInByte = 8 // cause we don't want store magic numbers
)

// tag name of struct tags
const (
	tagName = "tl"

	implicitFlag    = "implicit"
	isBitflagFlag   = "bitflag"
	omitemptyPrefix = "omitempty"
)
