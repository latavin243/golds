package golds

import (
	"bufio"
	"io"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 22:34:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-16 14:23:22
 */
var (
	defaultReaderSize = 8192
)

type Decoder interface {
	Decode() (*Packet, error)
}

type StreamingDocoder struct {
	bufioReader *bufio.Reader
}

func NewStreamingDecoderSize(reader io.Reader, size int) *StreamingDocoder {
	streamingDecoder := new(StreamingDocoder)
	streamingDecoder.bufioReader = bufio.NewReaderSize(reader, size)
	return streamingDecoder
}

func NewStreamingDecoder(reader io.Reader) *StreamingDocoder {
	return NewStreamingDecoderSize(reader, defaultReaderSize)
}

func (this *StreamingDocoder) Decode() (*Packet, error) {
	firstByte, err := this.bufioReader.ReadByte()
	if err != nil {
		return nil, err
	}
	packet := &Packet{}
	packet.PacketType = PacketType(firstByte)
	switch packet.PacketType {
	case PacketTypeString:
		packet.Value, err = this.decodeString()
	case PacketTypeError:
		packet.Value, err = this.decodeError()
	case PacketTypeInt:
		packet.Value, err = this.decodeInt()
	case PacketTypeBulkString:
		packet.Value, err = this.decodeBulkString()
	case PacketTypeArray:
		packet.Array, err = this.decodeArray()
	}
	return packet, nil
}

func (this *StreamingDocoder) decodeString() ([]byte, error) {
	return nil, nil
}
func (this *StreamingDocoder) decodeError() ([]byte, error) {
	return nil, nil
}

func (this *StreamingDocoder) decodeInt() ([]byte, error) {
	return nil, nil
}

func (this *StreamingDocoder) decodeBulkString() ([]byte, error) {
	return nil, nil
}

func (this *StreamingDocoder) decodeArray() ([]*Packet, error) {
	return nil, nil
}
