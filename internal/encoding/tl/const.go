// Copyright (c) 2024 RoseLoverX

package tl

const (
	WordLen   = 4
	LongLen   = WordLen * 2 // int64
	DoubleLen = WordLen * 2 // float64
	Int128Len = WordLen * 4 // int128
	Int256Len = WordLen * 8 // int256

	MagicNumber = 0xfe // 253

	// https://core.telegram.org/schema/mtproto
	CrcVector = 0x1cb5c415
	CrcFalse  = 0xbc799737
	CrcTrue   = 0x997275b5
	CrcNull   = 0x56730bcc

	bitsInByte = 8 // cause we don't want store magic numbers
)
