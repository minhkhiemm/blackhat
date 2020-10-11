package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	for {
		src, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(src)
	}
}

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "joescatcam.website:80")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	// copy src output to dst
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatal(err)
		}
	}()

	// copy dst output to src
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatal(err)
	}

}
