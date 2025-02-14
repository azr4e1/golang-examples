package echo

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func UDSEcho() {
	// create a unix domain socket and listen for incoming connections
	uds := "/tmp/echo.sock"
	socket, err := net.Listen("unix", uds)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Opening a unix socket at:", uds)

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
		log.Println("New connection established")
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a separate goroutine
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				// Create a buffer for incoming data
				buf := make([]byte, 4096)
				n, err := conn.Read(buf)
				if n == 0 {
					log.Println("Connection closed")
					break
				}
				log.Println("Incoming message of length:", n)
				if err != nil {
					log.Fatal(err)
				}
				// Echo the data back to the connection
				_, err = conn.Write(buf[:n])
				if err != nil {
					log.Fatal(err)
				}
			}
		}(conn)
	}
}

func TCPEcho() {
	addr := "localhost:9999"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	log.Println("Server is running on:", addr)

	for {
		conn, err := server.Accept()
		log.Println("New connection established")
		if err != nil {
			log.Println("Failed to accept conn.", err)
			continue
		}

		go func(conn net.Conn) {
			defer func() {
				conn.Close()
			}()
			for {
				buf := make([]byte, 4096)
				n, err := conn.Read(buf)
				if n == 0 {
					log.Println("Connection closed")
					break
				}
				log.Println("Incoming message of length:", n)
				if err != nil {
					log.Fatal(err)
				}
				// Echo the data back to the connection
				_, err = conn.Write(buf[:n])
				if err != nil {
					log.Fatal(err)
				}
			}
		}(conn)
	}
}

func STDIOEcho() {
	log.Println("Running on stdio")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		os.Stdout.WriteString(input + "\n")
	}
}
