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
 * @Last Modified time: 2020-04-16 14:12:54
 */

func TestDecoder(t *testing.T) {
	data := []byte("*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n")
	decoder := golds.NewStreamingDecoder(bytes.NewReader(data))
	packet, err := decoder.Decode()
	if err != nil {
		panic(err)
	}
	fmt.Println(packet)
}
