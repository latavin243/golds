/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 23:37:30
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-15 23:44:45
 */
package golds_test

import (
	"fmt"
	"testing"

	"github.com/monitor1379/golds"
)

func TestPacketType(t *testing.T) {
	fmt.Println(golds.PacketTypeString)
	fmt.Println(golds.PacketTypeError)
	fmt.Println(golds.PacketTypeInt)
	fmt.Println(golds.PacketTypeBulkString)
	fmt.Println(golds.PacketTypeArray)

	var b golds.PacketType
	fmt.Println(b)
}
