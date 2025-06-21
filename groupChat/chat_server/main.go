package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	client    = make(map[net.Conn]string)
	mutex     = sync.Mutex{}
	broadcast = make(chan string)
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Started the chat server on port 8000 ...")
	go broadcastMessage()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)

	}

}

func broadcastMessage() {
	for {
		msg := <- broadcast
		mutex.Lock()
		for conn := range client{
			conn.Write([]byte(msg + "\n"))
		}
		mutex.Unlock()
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Enter your name: "))
	nameInput := bufio.NewReader(conn)
	name, _ := nameInput.ReadString('\n')
	name = strings.TrimSpace(name)

	mutex.Lock()
	client[conn] = name
	mutex.Unlock()

	broadcast <- fmt.Sprintf("%s has joined the chat", name)

	for {
		msg, err := nameInput.ReadString('\n')
		if err != nil {
			break
		}
		trimmed := strings.TrimSpace(msg)
		if trimmed != "" {
			broadcast <- fmt.Sprintf("[%s]: %s", name, trimmed)
		}
	}

	mutex.Lock()
	delete(client, conn)
	mutex.Unlock()
	broadcast <- fmt.Sprintf("%s has left the chat\n", name)
}
