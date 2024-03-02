/*
type Listener interface {
	// Accept waits for and returns the next connection to the listener.
	Accept() (c Conn, err error)
	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error
	// Addr returns the listener's network address.
	Addr() Addr
}*/

package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"sync"
)

var mu sync.Mutex

func server() {
	// Listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server listening on port 9999")

	for {
		// Accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection in a goroutine
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// Receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		mu.Lock()
		fmt.Println("Received", msg)
		mu.Unlock()
	}
	c.Close()
}

func client() {
	// Connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the message
	msg := "Hello, world"
	mu.Lock()
	fmt.Println("Sending", msg)
	mu.Unlock()

	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
