package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 21:13:09
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-16 13:54:58
 */

type ServerOptions struct {
	ReadBufferSize int
}

var defaultServerOptions = ServerOptions{
	ReadBufferSize: 1024 * 8,
}
