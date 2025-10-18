package main

import (
	"fmt"
	"net"
)

func main() {
    addr := net.UDPAddr{
        Port: 8800,
        IP:   net.ParseIP("127.0.0.1"),
    }

    conn, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Println("Error starting UDP server:", err)
        return
    }
    defer conn.Close()

    fmt.Println("UDP server listening on 127.0.0.1:8800")

    buf := make([]byte, 1024)
    for {
        n, remoteAddr, err := conn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("Error reading:", err)
            continue
        }

        fmt.Printf("Received %d bytes from %s: %s\n", n, remoteAddr, string(buf[:n]))

        // send a response
        _, err = conn.WriteToUDP([]byte("Hello client!"), remoteAddr)
        if err != nil {
            fmt.Println("Error writing:", err)
        }
    }
}
