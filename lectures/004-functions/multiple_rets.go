package main

import "fmt"
                          /* multiple returns */
func twiceThrice (y int)   (   int   ,   int   )     {
  return    y*2, y*3 // return 2 values
}

func onceTwiceThrice (y int)    ( int, int, int )   {
  return  y*1,  y*2,  y*3 // return 3 values
}

func main() {
  r1, r2 := twiceThrice(100)
  fmt.Println("r1 = ", r1, "r2 = ", r2)

  v1, v2, v3 := onceTwiceThrice(2000)
  fmt.Println("v1 = ", v1, "v2 = ", v2, "v3 = ", v3)
}
