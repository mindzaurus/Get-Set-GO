// We will ignore error checking to simplify program look
// Try this only on a PC not on online IDEs

package main

import (
  "fmt"
  "net"
)

func main() {
  fmt.Println("Will be listening on localhost:12001")

  listener, _ := net.Listen("tcp", "localhost:12001")
  for { // serve requests for ever
    conn, _ := listener.Accept()
    go serveRequest(conn)
  }
  listener.Close()
}

func serveRequest(conn net.Conn) {
  buf := make([]byte, 1024)
  conn.Read(buf)
  conn.Write([]byte("Gophers are cool!\n"))
  conn.Close()
}
