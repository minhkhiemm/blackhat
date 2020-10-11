package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatal(err)
	}
	//reader := bufio.NewReader(conn)
	//s, err := reader.ReadString('\n')
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Printf("Read %d bytes: %s", len(s), s)
	//log.Println("Writing data")
	//writer := bufio.NewWriter(conn)
	//if _, err := writer.WriteString(s); err != nil {
	//	log.Fatal(err)
	//}
	//writer.Flush()
	//
	//b := make([]byte, 512)
	//for {
	//	size, err := conn.Read(b[0:])
	//	if err == io.EOF {
	//		log.Println("Client have been closed connection")
	//		break
	//	}
	//	if err != nil {
	//		log.Printf("Unexpected error: %v", err)
	//		break
	//	}
	//	log.Printf("read: %dbytes: %s\n", size, string(b))
	//
	//	if _, err := conn.Write(b[:size]); err != nil {
	//		log.Fatalf("Write error: %v", err)
	//	}
	//}
	// read data from conn via []byte
	// write data to conn
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatal("Unable to bind port")
	}
	log.Println("Listen on port :20080")
	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Connect error: %v", err)
		}

		go echo(conn)
	}
	// create a new listener in port 20000
	// accept all connection
	// run echo concurrently
}
