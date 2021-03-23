package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process1(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process1(conn)
	}
}

// 客户端
//func main() {
//	conn, err := net.Dial("tcp", "127.0.0.1:30000")
//	if err != nil {
//		fmt.Println("dial failed, err", err)
//		return
//	}
//	defer conn.Close()
//	for i := 0; i < 20; i++ {
//		msg := `Hello, Hello. How are you?`
//		conn.Write([]byte(msg))
//	}
//}
