package main

import (
    "fmt"
    "io"
    "net"
    "os"
)

func main() {
    fmt.Println("Redis-Server listening on port :6379")

    // create a TCP listener
    l, err := net.Listen("tcp", ":6379")
    if err != nil {
        fmt.Println(err)
        return
    }

    // receive requests
    conn, err := l.Accept()
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()  // close connection once finished

    // infinte loop to receive and response to commands from clients
    for {
        buf := make([]byte, 1024)
        _, err := conn.Read(buf)
        if err != nil {
            if err == io.EOF {
                fmt.Println("error reading from client: ", err.Error())
                os.Exit(1)
            }
        }

        // ignore request and send back a pong
        conn.Write([]byte("+Pong From GoRedis\r\n"))
    }
}
