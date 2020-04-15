package golds

import (
	"bufio"
	"io"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 22:34:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-15 23:53:49
 */

type Decoder interface {
	Decode(io.Reader) (interface{}, error)
}

type StreamingDocoder struct {
	bufferSize uint64
}

func NewStreamingDecoder(bufferSize uint64) *StreamingDocoder {
	decoder := new(StreamingDocoder)
	decoder.bufferSize = bufferSize
	return decoder
}

func (this *StreamingDocoder) Decode(reader io.Reader) (interface{}, error) {
	bufioReader := bufio.NewReader(reader)

	return []byte{}, nil
}
