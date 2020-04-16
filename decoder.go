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
 * @Last Modified time: 2020-04-16 22:52:09
 */
var (
	defaultReaderSize = 8192
)

var (
	ErrBadLF                  = errors.New("end without \\n")
	ErrBulkBytesLengthInvalid = errors.New("invalid bulk bytes length")
	ErrBulkBytesLengthTooLong = errors.New("bulk bytes length is too long")

	ErrArrayLengthInvalid = errors.New("invalid array length")
	ErrArrayLengthTooLong = errors.New("array length is too long")
)

const (
	MaxBulkBytesLength = 1024 * 1024 * 1 // 1MB
	MaxArrayLength     = 1024 * 1024 * 1 // 1M
)

type Decoder interface {
	Decode() (*Packet, error)
}

type StreamingDecoder struct {
	bufioReader *bufio.Reader
}

func NewStreamingDecoderSize(reader io.Reader, size int) *StreamingDecoder {
	streamingDecoder := new(StreamingDecoder)
	streamingDecoder.bufioReader = bufio.NewReaderSize(reader, size)
	return streamingDecoder
}

func NewStreamingDecoder(reader io.Reader) *StreamingDecoder {
	return NewStreamingDecoderSize(reader, defaultReaderSize)
}

func (this *StreamingDecoder) Decode() (*Packet, error) {
	return this.decode()
}

func (this *StreamingDecoder) decode() (*Packet, error) {
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
	return packet, nil
}

func (this *StreamingDecoder) decodeBytes() ([]byte, error) {
	data, _, err := this.bufioReader.ReadLine()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *StreamingDecoder) decodeInt() (int, error) {
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

func (this *StreamingDecoder) decodeBulkBytes() ([]byte, error) {
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
	if data[n] == '\n' {
		return nil, ErrBadLF
	}

	return data[:n+1], nil
}

func (this *StreamingDecoder) decodeArray() ([]*Packet, error) {
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
