package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	// "io"
	"log"
	"net"
	// "os"
)

func tcpClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// io.Copy(os.Stdout, conn)
	message := make([]byte, 1024)
	n, err := conn.Read(message)
	if err != nil {
		log.Println(err)
	}
	// fmt.Println(string(message[:n]))
	fmt.Println(strings.ToUpper(string(message[:n])))
}

func tcpServer() {
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()
	conn, err := ln.Accept()
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Write([]byte("Hello from Gopher!"))
	if err != nil {
		log.Println(err)
	}
}

func testClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	
	for i:=0; i<3; i++ {
		message := make([]byte, 1024)
		n, err := conn.Read(message)
		if err != nil {
			log.Println(err)
		}
		fmt.Print(strings.ToUpper(string(message[:n])))
	}
}

func testServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("all lower"))
	time.Sleep(1000)
	conn.Write([]byte("MiXeD cAsE"))
	time.Sleep(1000)
	conn.Write([]byte("ALL UPPER"))
	time.Sleep(1000)
}

func SendToServerWithTimeout() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		var source string
		fmt.Print("Input a word: ")
		_, err := fmt.Scan(&source)
		if err != nil {
			fmt.Println("Incorrect input", err)
			continue
		}

		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Response: ")
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))  // large for initial conn
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))  // lessen after conn is set up
		}
	}
}

func exampleHTTPViaNet() {
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: golang.org\n\n"
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		log.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}

func main() {
	// go tcpServer()
	// go testServer()
	// time.Sleep(1 * time.Second)

	// go tcpClient()
	// go testClient()
	// time.Sleep(1 * time.Second)
	exampleHTTPViaNet()
}
