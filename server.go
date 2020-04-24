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

	"github.com/syndtr/goleveldb/leveldb"
)

type Server struct {
	serverOptions ServerOptions
	db            *leveldb.DB
}

func NewServer(db *leveldb.DB) *Server {
	serverOptions := ServerOptions{
		PacketDecoderBufferSize: defaultPacketDecoderBufferSize,
	}
	return NewServerWithServerOptions(db, serverOptions)
}

func NewServerWithServerOptions(db *leveldb.DB, serverOptions ServerOptions) *Server {
	server := new(Server)
	server.db = db
	server.serverOptions = serverOptions

	if server.serverOptions.PacketDecoderBufferSize == 0 {
		server.serverOptions.PacketDecoderBufferSize = defaultPacketDecoderBufferSize
	}
	return server
}

func (this *Server) Listen(address string) error {
	if this.db == nil {
		return fmt.Errorf("db is nil")
	}

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

	packetDecoder := NewPacketDecoderSize(conn, this.serverOptions.PacketDecoderBufferSize)
	packetEncoder := NewPacketEncoder(conn)

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

		fmt.Println("request packet:", reqPacket)

		// Process command
		fmt.Println("process command")

		respPacket := &Packet{PacketType: PacketTypeString, Value: []byte("OK")}
		_, err = packetEncoder.Encode(respPacket)
		if err != nil {
			fmt.Printf("ERROR(golds): Encode packet error: %s\n", err)
		}
	}
}
