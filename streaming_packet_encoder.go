package golds

import (
	"bytes"
	"io"
)

/*
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 00:43:44
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 14:43:50
 */

type StreamingPacketEncoder struct {
	writer io.Writer
}

func NewStreamingPacaketEncoderSize(writer io.Writer, size int) *StreamingPacketEncoder {
	streamingPacketEncoder := new(StreamingPacketEncoder)
	streamingPacketEncoder.writer = writer
	return streamingPacketEncoder
}

func NewStreamingPacaketEncoder(writer io.Writer) *StreamingPacketEncoder {
	return NewStreamingPacaketEncoderSize(writer, defaultPacketEncoderWriterSize)
}

func (this *StreamingPacketEncoder) Encode(packet *Packet) (int, error) {
	data, err := this.encode(packet)
	if err != nil {
		return 0, err
	}
	return this.writer.Write(data)
}

func (this *StreamingPacketEncoder) encode(packet *Packet) ([]byte, error) {
	var data []byte
	var err error
	switch packet.PacketType {
	case PacketTypeString:
		data, err = this.encodeString(packet)
	case PacketTypeError:
	case PacketTypeInt:
	case PacketTypeBulkString:
	case PacketTypeArray:
	default:

	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *StreamingPacketEncoder) encodeString(packet *Packet) ([]byte, error) {
	var err error
	buf := &bytes.Buffer{}
	err = buf.WriteByte(byte(packet.PacketType))
	if err != nil {
		return nil, err
	}

	n, err := buf.Write(packet.Value)
	if err != nil {
		return nil, err
	}
	if n != len(packet.Value) {
		return nil, err
	}

	err = buf.WriteByte('\n')
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
