package golds

import "errors"

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 22:34:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 00:46:49
 */

var (
	ErrBadLF                  = errors.New("end without \\n")
	ErrBulkBytesLengthInvalid = errors.New("invalid bulk bytes length")
	ErrBulkBytesLengthTooLong = errors.New("bulk bytes length is too long")

	ErrArrayLengthInvalid = errors.New("invalid array length")
	ErrArrayLengthTooLong = errors.New("array length is too long")
)

const (
	defaultPacketDecoderReaderSize = 8192
	MaxBulkBytesLength             = 1024 * 1024 * 1 // 1MB
	MaxArrayLength                 = 1024 * 1024 * 1 // 1M
)

type PacketDecoder interface {
	Decode() (*Packet, error)
}
