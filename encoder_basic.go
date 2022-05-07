// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
)

type Encoder struct {
	w        io.Writer
	registry *Registry

	endianess binary.ByteOrder
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w, registry: defaultRegistry, endianess: binary.LittleEndian}
}

func (e *Encoder) SetRegistry(registry *Registry) { e.registry = registry }

func (e *Encoder) write(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	n, err := e.w.Write(b)
	if err != nil {
		return err
	}

	if n != len(b) {
		return &ErrorPartialWrite{Has: n, Want: len(b)}
	}

	return nil
}

func (e *Encoder) PutRawBytes(b []byte) error {
	if len(b)%WordLen != 0 {
		return errors.New("raw bytes does not divide by word size of protocol")
	}

	return e.write(b)
}

// PutBool очень специфичный тип, т.к. есть отдельный конструктор под true и false,
// то можно считать, что это две crc константы
func (e *Encoder) PutBool(v bool) error {
	crc := CrcFalse
	if v {
		crc = CrcTrue
	}

	return e.PutUint(uint32(crc))
}

func (e *Encoder) PutUint(v uint32) error {
	buf := make([]byte, WordLen)
	e.endianess.PutUint32(buf, v)
	return e.write(buf)
}

// PutCRC is an alias for Encoder.PutUint. It uses only for understanding what your code do (like
// self-documented code)
func (e *Encoder) PutCRC(v uint32) error {
	return e.PutUint(v) // я так и не понял, кажется что crc это bigendian, но видимо нет
}

func (e *Encoder) PutInt(v int32) error {
	return e.PutUint(uint32(v))
}

func (e *Encoder) PutLong(v int64) error {
	buf := make([]byte, LongLen)
	e.endianess.PutUint64(buf, uint64(v))
	return e.write(buf)
}

func (e *Encoder) PutDouble(v float64) error {
	buf := make([]byte, DoubleLen)
	e.endianess.PutUint64(buf, math.Float64bits(v))
	return e.write(buf)
}

func (e *Encoder) PutString(msg string) error {
	return e.PutMessage([]byte(msg))
}

func (e *Encoder) PutMessage(msg []byte) error {
	if len(msg) < FuckingMagicNumber {
		return e.putTinyBytes(msg)
	}
	return e.putLargeBytes(msg)
}

func (e *Encoder) putTinyBytes(msg []byte) error {
	if len(msg) >= FuckingMagicNumber {
		// it's panicing, cause, you shouldn' call this func by your
		// hands. panic required for internal purposes
		panic("tiny messages supports maximum 253 elements")
	}

	// how does it works:
	// any object can be putted to byte set ONLY with length, without modula after dividing to word length.
	// e.g. bytes 'Hi!' can be written as:
	//            | 0x03 0x48 0x6A 0x21 |
	// Divides by 32 bits? Yes, so it's good.
	//
	// BUT! bytes 'Hello!' MUST be written as
	//            | 0x06 0x48 0x65 0x6C | 0x6C 0x6F 0x21 0x00 |
	// See? We added extra empty byte to pad message to length of word. That is most important part of putting
	// bytes to buffer.
	//
	// So we must to create a buffer with lenth mod to 32 == 0. To not add extra bytes manually. They could be
	// random, but who needs that, right?

	const byteStreamLengthSize = 1

	for _, elem := range [][]byte{
		{byte(len(msg))},
		msg,
		make([]byte, pad(byteStreamLengthSize, WordLen, len(msg))),
	} {
		err := e.write(elem)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) putLargeBytes(msg []byte) error {
	if len(msg) < FuckingMagicNumber {
		// it's panicing, cause, you shouldn' call this func by your
		// hands. panic required for internal purposes
		panic("can't save binary stream with length less than 253 bytes")
	}

	const maxLen = 1 << ((WordLen - 1) * bitsInByte) // 3 left bytes of word, which is 4 bytes
	if len(msg) > maxLen {
		return fmt.Errorf("message entity too large: expect less than %v, got %v", maxLen, len(msg))
	}

	const byteStreamLengthSize = 4

	for _, elem := range [][]byte{
		{FuckingMagicNumber},
		littleUint24Bytes(len(msg)),
		msg,
		make([]byte, pad(byteStreamLengthSize, WordLen, len(msg))),
	} {
		err := e.write(elem)
		if err != nil {
			return err
		}
	}

	return nil
}
