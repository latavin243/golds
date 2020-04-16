package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-16 23:17:05
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 00:47:12
 */

const (
	defaultPacketEncoderWriterSize = 8192
)

type PacketEncoder interface {
	Encode(*Packet) (int, error)
}
