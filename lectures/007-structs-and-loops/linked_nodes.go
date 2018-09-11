package main

import (
  "fmt"
)

type Node struct {
  val   int // value this Node will hold
  next  *Node // pointer to next node
}

func createNodes() {
  firstNode     := new(Node) // create a new node, store its address in firstNode
  firstNode.val = 100

  secondNode      := new(Node) // create a new node, store its address in secondNode
  secondNode.val  = 200
  secondNode.next = nil

  firstNode.next  = secondNode // this links first node's next pointer to second node
  secondNode.next = nil // second node doesn't link to any other node

  ptr := firstNode // ptr points at first node now
  fmt.Printf("Node at ptr %p has value %d and has next node %p \n", ptr, ptr.val, ptr.next)

  ptr = ptr.next // ptr points at second node now
  fmt.Printf("Node at ptr %p has value %d and has next node %p \n", ptr, ptr.val, ptr.next)
}

func printAllNodesInfo(ptr *Node) {
  if ptr == nil {
    return
  }
  fmt.Printf("Node at ptr %p has value %d and has next node %p \n", ptr, ptr.val, ptr.next)
  printAllNodesInfo(ptr.next) // recursive call
}

func createNodes2() {
  var firstNode  *Node  = &Node{ 100, &Node{200, nil} } // Create 2 nodes and link them
  printAllNodesInfo(firstNode)
}


func main() {
  createNodes2()
}
