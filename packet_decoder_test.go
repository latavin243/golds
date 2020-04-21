package golds_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/monitor1379/golds"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-16 13:46:30
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-21 22:54:55
 */

func TestPacketDecoder(t *testing.T) {
	var decoder *golds.PacketDecoder
	var err error
	var packet *golds.Packet

	testCases := []string{
		"+OK\nother",
		"-Error: blah blah blah\n",
		":123\n",
		":-1\n",
		"$5\nhello\n",
		"$11\nhello\nworld\n",
		"*1\n$2\nk1\n",
		"*2\n$2\nk1\n$3\nval\n",
		"*2\n:123\n$3\nval\n",
		"*3\n$3\nSET\n$3\nkey\n$5\nvalue\n",
	}

	for _, testCase := range testCases {
		decoder = golds.NewPacketDecoder(bytes.NewBuffer([]byte(testCase)))
		packet, err = decoder.Decode()
		if err != nil {
			t.Errorf("error: %s", err)
		}
		fmt.Println(packet)
	}
}
