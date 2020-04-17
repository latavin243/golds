package golds_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-17 00:45:45
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 14:48:57
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
		golds.Packet{PacketType: golds.PacketTypeError, Value: []byte("Error: error message")},
		golds.Packet{PacketType: golds.PacketTypeInt, Value: []byte("1234")},
		golds.Packet{PacketType: golds.PacketTypeInt, Value: []byte("-1")},
		golds.Packet{PacketType: golds.PacketTypeBulkString, Value: []byte("hello")},
		golds.Packet{PacketType: golds.PacketTypeBulkString, Value: []byte("")},
		golds.Packet{PacketType: golds.PacketTypeBulkString, Value: nil},
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
