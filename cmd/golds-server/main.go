/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 21:27:07
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-15 22:07:18
 */
package main

import "github.com/monitor1379/golds"

func main() {
	server := golds.NewServer()
	server.Listen(":3000")

}
