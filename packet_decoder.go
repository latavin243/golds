package golds

import (
	"bufio"
	"errors"
	"io"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 22:34:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-21 22:54:26
 */

var (
	ErrBadLF                  = errors.New("end without \\n")
	ErrBulkBytesLengthInvalid = errors.New("invalid bulk bytes length")
	ErrBulkBytesLengthTooLong = errors.New("bulk bytes length is too long")

	ErrArrayLengthInvalid = errors.New("invalid array length")
	ErrArrayLengthTooLong = errors.New("array length is too long")
)

const (
	defaultPacketDecoderReaderSize = 8192
	MaxBulkBytesLength             = 1024 * 1024 * 1 // 1MB
	MaxArrayLength                 = 1024 * 1024 * 1 // 1M
)

type PacketDecoder struct {
	bufioReader *bufio.Reader
}

func NewPacketDecoderSize(reader io.Reader, size int) *PacketDecoder {
	packetDecoder := new(PacketDecoder)
	packetDecoder.bufioReader = bufio.NewReaderSize(reader, size)
	return packetDecoder
}

func NewPacketDecoder(reader io.Reader) *PacketDecoder {
	return NewPacketDecoderSize(reader, defaultPacketDecoderReaderSize)
}

func (this *PacketDecoder) Decode() (*Packet, error) {
	return this.decode()
}

func (this *PacketDecoder) decode() (*Packet, error) {
	firstByte, err := this.bufioReader.ReadByte()
	if err != nil {
		return nil, err
	}
	packet := &Packet{}
	packet.PacketType = PacketType(firstByte)

	switch packet.PacketType {
	case PacketTypeString, PacketTypeError, PacketTypeInt:
		packet.Value, err = this.decodeBytes()
	case PacketTypeBulkString:
		packet.Value, err = this.decodeBulkBytes()
	case PacketTypeArray:
		packet.Array, err = this.decodeArray()
	default:
		return nil, ErrInvalidPacketType
	}

	if err != nil {
		return nil, err
	}
	return packet, nil
}

func (this *PacketDecoder) decodeBytes() ([]byte, error) {
	data, _, err := this.bufioReader.ReadLine()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *PacketDecoder) decodeInt() (int, error) {
	data, _, err := this.bufioReader.ReadLine()
	if err != nil {
		return 0, err
	}

	number, err := Btoi64(data)
	if err != nil {
		return 0, err
	}

	return int(number), nil
}

func (this *PacketDecoder) decodeBulkBytes() ([]byte, error) {
	n, err := this.decodeInt()
	if err != nil {
		return nil, err
	}
	if n < -1 {
		return nil, ErrBulkBytesLengthInvalid
	}
	if n > MaxBulkBytesLength {
		return nil, ErrBulkBytesLengthTooLong
	}
	if n == -1 {
		return nil, nil
	}

	data := make([]byte, n+1)
	_, err = io.ReadFull(this.bufioReader, data)
	if err != nil {
		return nil, err
	}
	if data[n] != '\n' {
		return nil, ErrBadLF
	}
	return data[:n], nil
}

func (this *PacketDecoder) decodeArray() ([]*Packet, error) {
	n, err := this.decodeInt()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrArrayLengthInvalid
	}
	if n > MaxArrayLength {
		return nil, ErrArrayLengthTooLong
	}

	var packets []*Packet
	for i := 0; i < n; i++ {
		packet, err := this.decode()
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)
	}
	return packets, nil
}
