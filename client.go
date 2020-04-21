package golds

import (
	"net"
	"sync"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-21 22:39:22
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-21 22:51:33
 */

const (
	defaultNetwork = "tcp"
)

type Client struct {
	conn   net.Conn
	locker sync.Mutex
}

func Dial(address string) (*Client, error) {
	conn, err := net.Dial(defaultNetwork, address)
	if err != nil {
		return nil, err
	}
	client := &Client{
		conn:   conn,
		locker: sync.Mutex{},
	}
	return client, nil
}

func (this *Client) Set(key, value []byte) error {
	return nil
}

func (this *Client) Get(key []byte) ([]byte, error) {
	return nil, nil
}
