package golds

import (
	"net"
	"sync"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-21 22:39:22
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-21 23:11:40
 */

const (
	defaultNetwork = "tcp"
)

type Client struct {
	conn          net.Conn
	locker        *sync.Mutex
	packetEncoder *PacketEncoder
	packetDecoder *PacketDecoder
}

func Dial(address string) (*Client, error) {
	conn, err := net.Dial(defaultNetwork, address)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:          conn,
		locker:        &sync.Mutex{},
		packetEncoder: NewPacketEncoder(conn),
		packetDecoder: NewPacketDecoder(conn),
	}
	return client, nil
}

func (this *Client) Set(key, value []byte) error {
	packet := &Packet{
		PacketType: PacketTypeString,
		Value:      []byte("fuck you"),
	}
	_, err := this.packetEncoder.Encode(packet)
	if err != nil {
		return err
	}
	return nil
}

func (this *Client) Get(key []byte) ([]byte, error) {
	packet, err := this.packetDecoder.Decode()
	if err != nil {
		return nil, nil
	}

	// TODO(monitor1379): 解析packet.Value或者key对应的value
	return packet.Value, nil
}
