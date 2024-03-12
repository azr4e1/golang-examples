package echo

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func LocalEcho() {
	// create a unix domain socket and listen for incoming connections
	socket, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup the sockfile
	c := make(chan os.Signal, 1)
	// this creates a goroutine that will pass to the channel a signal value of type SIGTERM
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		// as soon as the receive action is unblocked (a SIGTERM has been issued, i.d. application is terminated, this goroutine will clean up the socket and exit the program wil exit code 1)
		<-c
		os.Remove("/tmp/echo.sock")
		os.Exit(1)
	}()

	for {
		// Accept an incoming connection
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a separate goroutine
		go func(conn net.Conn) {
			defer conn.Close()
			// Create a buffer for incoming data
			buf := make([]byte, 4096)

			// Read data from the connection
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			// Echo the data back to the connection
			_, err = conn.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}(conn)
	}
}
