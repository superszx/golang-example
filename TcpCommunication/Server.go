package TcpCommunication

import (
	"fmt"
	"net"
)

type Server struct {
	ip   string
	port string
}

func (server *Server) Init() {
	//create a listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.ip, server.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("server created a listener")

	for {
		fmt.Println("server listening ", server.ip, ":", server.port, " , waiting for connect...")
		// listen to pointed port ,and block to wait client to connect
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleRequest(conn)
	}
}

func HandleRequest(context net.Conn) {
	defer context.Close()
	buf := make([]byte, 1024)
	//read from received content
	n, err := context.Read(buf)
	if err != nil {
		println(err)
		return
	}
	content := string(buf[0:n])
	fmt.Println("server receive : ", content)

	_, err = context.Write([]byte(fmt.Sprintf("has receive you request message %s", content)))
	if err != nil {
		fmt.Println(err)
		return
	}
}
