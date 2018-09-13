package TcpCommunication

import (
	"fmt"
	"net"
)

type Client struct {
	ip   string
	port string
}

func (client *Client) Init() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", client.ip, client.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("client has connect to server")

	Send(conn)
}

func Send(conn net.Conn) {
	_, err := conn.Write([]byte("hello server"))
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1034)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("client receive reply : ", string(buf[0:n]))
}
