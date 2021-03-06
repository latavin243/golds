package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-27 21:10:58
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-06 11:48:57
 */

import (
	"fmt"
	"net"
	"sync"

	"github.com/monitor1379/golds/handlers"

	"github.com/monitor1379/golds/goldscore"
)

type Client struct {
	sync.Mutex
	conn          net.Conn
	packetEncoder *goldscore.PacketEncoder
	packetDecoder *goldscore.PacketDecoder
}

func Dial(address string) (*Client, error) {
	conn, err := net.Dial(defaultNetwork, address)
	if err != nil {
		return nil, err
	}
	client := &Client{
		conn:          conn,
		packetEncoder: goldscore.NewPacketEncoder(conn),
		packetDecoder: goldscore.NewPacketDecoder(conn),
	}
	return client, nil
}

func (this *Client) Close() error {
	return this.conn.Close()
}

func (this *Client) do(requestPacket *goldscore.Packet) (*goldscore.Packet, error) {
	err := this.packetEncoder.Encode(requestPacket)
	if err != nil {
		return nil, err
	}

	responsePacket, err := this.packetDecoder.Decode()
	if err != nil {
		return nil, err
	}

	if responsePacket.PacketType == goldscore.PacketTypeError {
		return nil, fmt.Errorf("golds server error: %s", string(responsePacket.Value))
	}

	return responsePacket, nil
}

func (this *Client) Set(key, value []byte) error {
	this.Lock()
	defer this.Unlock()

	requestPacket := goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte(handlers.CommandNameSet))).
		Add(goldscore.NewBulkStringPacket(key)).
		Add(goldscore.NewBulkStringPacket(value))

	_, err := this.do(requestPacket)
	if err != nil {
		return err
	}

	return nil
}

func (this *Client) Get(key []byte) ([]byte, error) {
	this.Lock()
	defer this.Unlock()

	requestPacket := goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte(handlers.CommandNameGet))).
		Add(goldscore.NewBulkStringPacket(key))

	responsePacket, err := this.do(requestPacket)
	if err != nil {
		return nil, err
	}

	return responsePacket.Value, nil
}

func (this *Client) Del(key []byte) error {
	this.Lock()
	defer this.Unlock()

	requestPacket := goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte(handlers.CommandNameDel))).
		Add(goldscore.NewBulkStringPacket(key))

	responsePacket, err := this.do(requestPacket)
	if err != nil {
		return err
	}
	_ = responsePacket

	return nil
}

func (this *Client) Keys() ([][]byte, error) {
	this.Lock()
	defer this.Unlock()

	requestPacket := goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte(handlers.CommandNameKeys)))

	responsePacket, err := this.do(requestPacket)
	if err != nil {
		return nil, err
	}

	keys := [][]byte{}

	for _, subPacket := range responsePacket.Array {
		keys = append(keys, subPacket.Value)
	}

	return keys, nil
}
