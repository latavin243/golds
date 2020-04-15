package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 23:45:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-15 23:55:34
 */

type Packet struct {
	Type  PacketType
	Value []byte
	Array []*Packet
}
