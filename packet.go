package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 23:45:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-16 14:12:36
 */

type Packet struct {
	PacketType PacketType
	Value      []byte
	Array      []*Packet
}
