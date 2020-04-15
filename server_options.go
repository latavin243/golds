package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 21:13:09
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-15 23:45:13
 */

type ServerOptions struct {
	ReadBufferSize uint64
}

var defaultServerOptions = ServerOptions{
	ReadBufferSize: 1024 * 8,
}