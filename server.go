package golds

import (
	"fmt"
	"net"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 21:13:09
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-16 14:10:56
 */

type Server struct {
	serverOptions ServerOptions
}

func NewServer() *Server {
	return NewServerWithServerOptions(defaultServerOptions)
}

func NewServerWithServerOptions(serverOptions ServerOptions) *Server {
	server := new(Server)
	server.serverOptions = serverOptions
	return server
}

func (this *Server) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			// TODO(monitor1379): what situation will hit here
			fmt.Println("ERROR(Accept):", err)
			return err
		}

		go this.handleConn(conn)
	}
}

func (this *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	// readChan := make(chan []byte)
	// writeChan := make(chan []byte)
	// stopChan := make(chan bool)
	// go this.readConn(conn, readChan, stopChan)
	// go this.writeConn(conn, writeChan, stopChan)

	// var isStopped bool
	// for {
	// 	select {
	// 	case data := <-readChan:
	// 		fmt.Println("READ:", string(data))
	// 	case <-stopChan:
	// 		fmt.Println("STOP")
	// 		isStopped = true
	// 	}
	// 	if isStopped == true {
	// 		break
	// 	}
	// }

}

// func (this *Server) readConn(conn net.Conn, readChan chan []byte, stopChan chan bool) {
// 	decoder := NewStreamingDecoder(this.serverOptions.ReadBufferSize)

// 	for {
// 		data, err := decoder.Decode(conn)
// 		if err != nil {
// 			fmt.Println("ERROR(Read):", err)
// 			break
// 		}
// 		fmt.Println("fuck", data)
// 		readChan <- data.([]byte)
// 	}
// 	stopChan <- true
// }

// func (this *Server) writeConn(conn net.Conn, writeChan chan []byte, stopChan chan bool) {
// 	for {
// 		data := <-writeChan
// 		conn.Write(data)
// 	}
// }
