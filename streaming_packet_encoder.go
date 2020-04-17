package golds

import (
	"bytes"
	"io"
	"strconv"
)

/*
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 00:43:44
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 14:48:32
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
	case PacketTypeString, PacketTypeError, PacketTypeInt:
		data, err = this.encodeBytes(packet)
	case PacketTypeBulkString:
		data, err = this.encodeBulkBytes(packet)
	case PacketTypeArray:
	default:

	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *StreamingPacketEncoder) encodeBytes(packet *Packet) ([]byte, error) {
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

func (this *StreamingPacketEncoder) encodeBulkBytes(packet *Packet) ([]byte, error) {
	var err error
	buf := &bytes.Buffer{}

	err = buf.WriteByte(byte(packet.PacketType))
	if err != nil {
		return nil, err
	}

	// 如果value数组为空，说明是"$-1\n"
	if packet.Value == nil {
		_, err = buf.WriteString("-1\n")
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	_, err = buf.WriteString(strconv.Itoa(len(packet.Value)))
	if err != nil {
		return nil, err
	}

	err = buf.WriteByte('\n')
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
