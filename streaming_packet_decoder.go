package golds

import (
	"bufio"
	"io"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-16 23:25:14
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-16 23:56:29
 */

var _ (PacketDecoder) = new(StreamingPacketDecoder)

type StreamingPacketDecoder struct {
	bufioReader *bufio.Reader
}

func NewStreamingPacketDecoderSize(reader io.Reader, size int) *StreamingPacketDecoder {
	streamingPacketDecoder := new(StreamingPacketDecoder)
	streamingPacketDecoder.bufioReader = bufio.NewReaderSize(reader, size)
	return streamingPacketDecoder
}

func NewStreamingPacketDecoder(reader io.Reader) *StreamingPacketDecoder {
	return NewStreamingPacketDecoderSize(reader, defaultPacketDecoderReaderSize)
}

func (this *StreamingPacketDecoder) Decode() (*Packet, error) {
	return this.decode()
}

func (this *StreamingPacketDecoder) decode() (*Packet, error) {
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

func (this *StreamingPacketDecoder) decodeBytes() ([]byte, error) {
	data, _, err := this.bufioReader.ReadLine()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *StreamingPacketDecoder) decodeInt() (int, error) {
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

func (this *StreamingPacketDecoder) decodeBulkBytes() ([]byte, error) {
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

func (this *StreamingPacketDecoder) decodeArray() ([]*Packet, error) {
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
