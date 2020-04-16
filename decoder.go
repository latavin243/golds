package golds

import (
	"bufio"
	"io"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 22:34:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-16 14:12:28
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
	packet := &Packet{PacketType: PacketType(firstByte)}

	return packet, nil
}
