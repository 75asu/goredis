package main

import (
    "fmt"
    "net"
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
        resp := NewResp(conn)
        value, err := resp.Read()
        if err != nil {
                fmt.Println(err)
                return
        }

        _ = value

        writer := NewWriter(conn)
        writer.Write(Value{typ: "string", str: "OK"})

    }


}
