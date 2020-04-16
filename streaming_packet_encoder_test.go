package golds_test

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/monitor1379/golds"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-17 00:45:45
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 00:51:23
 */

func TestStreamingPacketEncoder(t *testing.T) {
	buf := &bytes.Buffer{}
	writer := bufio.NewWriter(buf)
	streamingPacketEncoder := golds.NewStreamingPacaketEncoder(writer)
	_, err := streamingPacketEncoder.Encode(&golds.Packet{})
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
