package main

import (
  "fmt"
)

type Product struct {
  ID    int
  price float32
}

func halfPrice(p *Product) {
  p.price /= 2.0 // divide by 2
}

func tripleThePrice(pr *float32) {
  *pr = *pr + *pr + *pr
}

func main() {
  p := Product{ 7575, 77.46 }
  fmt.Println(p)

  halfPrice(&p) // pass address of the whole struct
  fmt.Println("halfPrice(&p)", p)

  tripleThePrice(&p.price) // pass address of a member
  fmt.Println("tripleThePrice(&p.price)", p)
  fmt.Println()

  // using new to allocate struct Product
  q := new (Product) // q will be of type *Product
  q.ID = 75398
  q.price = 55.75
  fmt.Println(q)

  halfPrice(q) // q is already a pointer pass it directly
  fmt.Println("halfPrice(q)", q)

  tripleThePrice(&q.price) // pass address of a member directly
  fmt.Println("tripleThePrice(&q.price)", q)
  fmt.Println()

  r := &Product { 87632, 22.87 } // this is same as  new (Product) and setting values
  fmt.Println(r)

  halfPrice(r)
  fmt.Println("halfPrice(r)", r)

  tripleThePrice(&r.price)
  fmt.Println("tripleThePrice(&r.price)", r)

}
