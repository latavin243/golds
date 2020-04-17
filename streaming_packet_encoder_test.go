package golds_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-17 00:45:45
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 14:45:31
 */
import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/monitor1379/golds"
)

func TestStreamingPacketEncoder(t *testing.T) {
	packets := []golds.Packet{
		golds.Packet{PacketType: golds.PacketTypeString, Value: []byte("hello world")},
	}

	for _, packet := range packets {
		buf := &bytes.Buffer{}
		streamingPacketEncoder := golds.NewStreamingPacaketEncoder(buf)
		_, err := streamingPacketEncoder.Encode(&packet)
		if err != nil {
			panic(err)
		}
		fmt.Println(strconv.Quote(buf.String()))
	}

}
