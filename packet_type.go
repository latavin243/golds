package golds

import (
	"errors"
	"fmt"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 23:29:18
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-16 21:35:14
 */

type PacketType byte

var (
	PacketTypeString     PacketType = '+'
	PacketTypeError      PacketType = '-'
	PacketTypeInt        PacketType = ':'
	PacketTypeBulkString PacketType = '$'
	PacketTypeArray      PacketType = '*'
)

var (
	ErrInvalidPacketType = errors.New("invalid packet type")
)

func (this PacketType) String() string {
	var t string
	switch this {
	case PacketTypeString:
		t = "string"
	case PacketTypeError:
		t = "error"
	case PacketTypeInt:
		t = "int"
	case PacketTypeBulkString:
		t = "bulk string"
	case PacketTypeArray:
		t = "array"
	default:
		t = "unknown type"
	}
	return fmt.Sprintf("{ '%v': (%s) }", string(byte(this)), t)
}
