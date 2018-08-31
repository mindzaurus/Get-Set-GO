package ezcfg

import "fmt"

var ClientPort int

// We want this init to be run but later we may use
func init() { // Called automatically
  fmt.Println("ezcfg::clientcfg.go setting ClientPort = 23409")
  ClientPort = 23409
}
