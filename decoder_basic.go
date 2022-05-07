// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE for details

package tl

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/pkg/errors"
)

// A Decoder reads and decodes TL values from an input stream.
type Decoder struct {
	r         bufio.Reader
	registry  *Registry
	endianess binary.ByteOrder

	peekedBytes int
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: *bufio.NewReader(r), registry: defaultRegistry, endianess: binary.LittleEndian}
}

func (d *Decoder) SetRegistry(registry *Registry) { d.registry = registry }

func (d *Decoder) peek(size int) ([]byte, error) {
	bytes, err := d.r.Peek(d.peekedBytes + size)
	if err != nil {
		return nil, err
	}
	seek := d.peekedBytes
	d.peekedBytes += size
	return bytes[seek:], nil
}

func (d *Decoder) success() {
	// discarded will be only exact number or less, BUT with returned error
	_, err := d.r.Discard(d.peekedBytes)
	if err != nil {
		panic(errors.Wrap(err, "something wrong with peeking bytes, Discard must always be ok"))
	}
	d.peekedBytes = 0
}

func (d *Decoder) PopLong() (int64, error) {
	val, err := d.peek(LongLen)
	if err != nil {
		return 0, err
	}

	d.success()
	return int64(binary.LittleEndian.Uint64(val)), nil
}

func (d *Decoder) PopDouble() (float64, error) {
	val, err := d.peek(LongLen)
	if err != nil {
		return 0, err
	}

	d.success()
	return math.Float64frombits(binary.LittleEndian.Uint64(val)), nil
}

func (d *Decoder) PopUint() (uint32, error) {
	val, err := d.peek(WordLen)
	if err != nil {
		return 0, err
	}

	d.success()
	return binary.LittleEndian.Uint32(val), nil
}

func (d *Decoder) PopRawBytes(size int) ([]byte, error) {
	val, err := d.peek(size)
	if err != nil {
		return nil, err
	}

	d.success()
	return val, nil
}

func (d *Decoder) PopBool() (bool, error) {
	crc, err := d.PopUint()
	if err != nil {
		return false, err
	}

	switch crc {
	case CrcTrue:
		return true, nil
	case CrcFalse:
		return false, nil
	default:
		return false, fmt.Errorf("want a 0x%08x (true) or 0x%08x (false); got 0x%08x", CrcTrue, CrcFalse, crc)
	}
}

func (d *Decoder) PopCRC() (uint32, error) {
	return d.PopUint()
}

func (d *Decoder) PopInt() (int32, error) {
	i, err := d.PopUint()
	return int32(i), err
}

func (d *Decoder) PopString() (string, error) {
	v, err := d.PopMessage()
	if err != nil {
		return "", err // consersion nil byte to string will cause a panic
	}
	return string(v), nil
}

func (d *Decoder) PopMessage() ([]byte, error) {
	readLen := 1
	buf, err := d.peek(1)
	if err != nil {
		return nil, err
	}
	firstByte := buf[0]

	// how exact bytes there is a message
	var realSize int

	if firstByte != FuckingMagicNumber { // tiny message, so real size is exact a first byte value
		realSize = int(firstByte)
	} else { // otherwise it's a large message, next three bytes are size of message
		readLen += WordLen - 1
		buf, err := d.peek(WordLen - 1)
		if err != nil {
			return nil, errors.Wrapf(err, "reading last %v bytes of message size", WordLen-1)
		}

		switch d.endianess {
		case binary.LittleEndian:
			buf = append(buf, 0x00)
		case binary.BigEndian:
			buf = append([]byte{0x00}, buf...)
		default:
			panic("wait, what?")
		}

		realSize = int(d.endianess.Uint32(buf))
	}

	// this buf wil be real message
	buf, err = d.peek(realSize)
	if err != nil {
		return nil, errors.Wrapf(err, "reading message data with len of %v", realSize)
	}

	p := pad(readLen, WordLen, realSize)
	if p > 0 {
		_, err = d.peek(p)
		if err != nil {
			return nil, errors.Wrapf(err, "reading %v last void bytes", p)
		}

	}

	d.success()
	return buf, nil
}

func (d *Decoder) GetRestOfMessage() ([]byte, error) {
	d.success()
	return io.ReadAll(&d.r)
}

func (d *Decoder) DumpWithoutRead() ([]byte, error) {
	return io.ReadAll(&d.r)
}
