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

func readAndSend(rw io.ReadWriter) {
	for {
		text, ok := readLine(os.Stdin)
		if ok {
			fmt.Fprintln(rw, text)
			ret, ok := readLine(rw)
			if ok {
				fmt.Println(ret)
			} else {
				return
			}
		} else {
			fmt.Printf("bye\n")
			break
		}
	}
}
func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	readAndSend(conn)
}
