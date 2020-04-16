package golds

import (
	"bufio"
	"io"
)

/*
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 00:43:44
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 00:50:52
 */

type StreamingPacketEncoder struct {
	bufioWriter *bufio.Writer
}

func NewStreamingPacaketEncoderSize(writer io.Writer, size int) *StreamingPacketEncoder {
	streamingPacketEncoder := new(StreamingPacketEncoder)
	streamingPacketEncoder.bufioWriter = bufio.NewWriterSize(writer, size)
	return streamingPacketEncoder
}

func NewStreamingPacaketEncoder(writer io.Writer) *StreamingPacketEncoder {
	return NewStreamingPacaketEncoderSize(writer, defaultPacketEncoderWriterSize)
}

func (this *StreamingPacketEncoder) Encode(packet *Packet) (int, error) {
	return 0, nil
}
