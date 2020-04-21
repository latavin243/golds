package golds

import (
	"bytes"
	"io"
	"strconv"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-16 23:17:05
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-21 22:56:25
 */

const (
	defaultPacketEncoderWriterSize = 8192
)

type PacketEncoder struct {
	writer io.Writer
}

func NewPacketEncoderSize(writer io.Writer, size int) *PacketEncoder {
	packetEncoder := new(PacketEncoder)
	packetEncoder.writer = writer
	return packetEncoder
}

func NewPacketEncoder(writer io.Writer) *PacketEncoder {
	return NewPacketEncoderSize(writer, defaultPacketEncoderWriterSize)
}

func (this *PacketEncoder) Encode(packet *Packet) (int, error) {
	data, err := this.encode(packet)
	if err != nil {
		return 0, err
	}
	return this.writer.Write(data)
}

func (this *PacketEncoder) encode(packet *Packet) ([]byte, error) {
	var data []byte
	var err error
	switch packet.PacketType {
	case PacketTypeString, PacketTypeError, PacketTypeInt:
		data, err = this.encodeBytes(packet)
	case PacketTypeBulkString:
		data, err = this.encodeBulkBytes(packet)
	case PacketTypeArray:
		data, err = this.encodeArray(packet)
	default:

	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *PacketEncoder) encodeBytes(packet *Packet) ([]byte, error) {
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

func (this *PacketEncoder) encodeBulkBytes(packet *Packet) ([]byte, error) {
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

func (this *PacketEncoder) encodeArray(packet *Packet) ([]byte, error) {
	var err error
	buf := &bytes.Buffer{}

	err = buf.WriteByte(byte(packet.PacketType))
	if err != nil {
		return nil, err
	}

	_, err = buf.WriteString(strconv.Itoa(len(packet.Array)))
	if err != nil {
		return nil, err
	}

	err = buf.WriteByte('\n')
	if err != nil {
		return nil, err
	}

	for _, subPacket := range packet.Array {
		data, err := this.encode(subPacket)
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(data)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
