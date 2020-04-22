package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-15 21:13:09
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-21 22:46:54
 */

import (
	"fmt"
	"io"
	"net"
)

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
			return err
		}

		go this.handleConn(conn)
	}
}

func (this *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	packetEncoder := NewPacketEncoder(conn)
	packetDecoder := NewPacketDecoder(conn)

	for {
		reqPacket, err := packetDecoder.Decode()
		if err == io.EOF {
			fmt.Println("DEBUG(golds): connection closed by peer")
			break
		}
		if err != nil {
			// TODO(monitor1379)
			fmt.Printf("ERROR(golds): Decode packet error: %s\n", err)
			continue
		}

		fmt.Println("request packet:", reqPacket.Array)

		// Process command
		fmt.Println("process command")

		respPacket := &Packet{PacketType: PacketTypeString, Value: []byte("OK")}
		_, err = packetEncoder.Encode(respPacket)
		if err != nil {
			fmt.Printf("ERROR(golds): Encode packet error: %s\n", err)
		}
	}
}
