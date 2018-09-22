package main

import (
  "fmt"
)

func dstSrcSameLen() {
  s1 := []int { 2, 5, 19 }
  s2 := make([]int, 3) // same as []int {0, 0, 0}

  fmt.Println("dstSrcSameLen s1 before copy", s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  fmt.Println("dstSrcSameLen s2 before copy", s2)
  fmt.Println("len(s2)", len(s2), "cap(s2)", cap(s2))

  copy(s2, s1)

  fmt.Println("dstSrcSameLen s1 after copy", s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  fmt.Println("dstSrcSameLen s2 after copy", s2)
  fmt.Println("len(s2)", len(s2), "cap(s2)", cap(s2))
}

func srcLenMore() {
  s1 := []int { 2, 5, 19, 22, 35 }
  s2 := make([]int, 3) // same as []int {0, 0, 0}

  fmt.Println("dstSrcSameLen s1 before copy", s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  fmt.Println("dstSrcSameLen s2 before copy", s2)
  fmt.Println("len(s2)", len(s2), "cap(s2)", cap(s2))

  copy(s2, s1)

  fmt.Println("dstSrcSameLen s1 after copy", s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  fmt.Println("dstSrcSameLen s2 after copy", s2)
  fmt.Println("len(s2)", len(s2), "cap(s2)", cap(s2))
}

func dstLenMore() {
  s1 := []int { 2, 5, 19 }
  s2 := make([]int, 5) // same as []int {0, 0, 0, 0, 0}

  fmt.Println("dstSrcSameLen s1 before copy", s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  fmt.Println("dstSrcSameLen s2 before copy", s2)
  fmt.Println("len(s2)", len(s2), "cap(s2)", cap(s2))

  copy(s2, s1)

  fmt.Println("dstSrcSameLen s1 after copy", s1)
  fmt.Println("len(s1)", len(s1), "cap(s1)", cap(s1))
  fmt.Println("dstSrcSameLen s2 after copy", s2)
  fmt.Println("len(s2)", len(s2), "cap(s2)", cap(s2))
}

func main() {
  dstSrcSameLen()
  fmt.Println()
  srcLenMore()
  fmt.Println()
  dstLenMore()
}
