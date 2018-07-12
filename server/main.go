package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func readLine(r io.Reader) (string, bool) {
	s := bufio.NewScanner(r)
	if s.Scan() {
		return s.Text(), true
	}
	return "", false
}

func main() {
	session := 0
	fmt.Println("listen start")
	sock, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		c, err := sock.Accept()
		if err != nil {
			panic(err)
		}
		session++
		go func(sessionId int, client io.ReadWriter) {
			fmt.Printf("start session %d\n", session)
			for {
				text, ok := readLine(client)
				fmt.Fprintf(os.Stderr, "receive from session %d [%v]\n", sessionId, text)
				if ok {
					fmt.Fprintf(client, "intput[%s]\n", text)
				} else {
					fmt.Fprintln(client, "bye")
					fmt.Printf("session %d ended\n", sessionId)
					return
				}
			}
		}(session, c)
	}
}
