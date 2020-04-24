package golds

import (
	"fmt"
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
	mu            *sync.Mutex
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
		mu:            &sync.Mutex{},
		packetEncoder: NewPacketEncoder(conn),
		packetDecoder: NewPacketDecoder(conn),
	}
	return client, nil
}

func (this *Client) Set(key, value []byte) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	reqPacket := &Packet{
		PacketType: PacketTypeArray,
		Array: []*Packet{
			&Packet{PacketType: PacketTypeBulkString, Value: []byte("SET")},
			&Packet{PacketType: PacketTypeBulkString, Value: key},
			&Packet{PacketType: PacketTypeBulkString, Value: value},
		},
	}

	_, err := this.packetEncoder.Encode(reqPacket)
	if err != nil {
		return err
	}

	respPacket, err := this.packetDecoder.Decode()
	if err != nil {
		return err
	}

	fmt.Println(respPacket)

	return nil
}

func (this *Client) Get(key []byte) ([]byte, error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	packet, err := this.packetDecoder.Decode()
	if err != nil {
		return nil, nil
	}

	// TODO(monitor1379): 解析packet.Value或者key对应的value
	return packet.Value, nil
}
