/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-21 22:39:07
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-21 22:52:03
 */

package main

import (
	"time"

	"github.com/monitor1379/golds"
)

func main() {
	client, err := golds.Dial("localhost:3000")
	if err != nil {
		panic(err)
	}

	err = client.Set([]byte("key1"), []byte("value"))
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	err = client.Set([]byte("key1"), []byte("value"))
	if err != nil {
		panic(err)
	}

}
