
package main

import (
  "fmt"
  "strings"
)

func main() {
    // func Join(a [ ]string, sep string) string
    s1 := [ ]string{"apple", "orange", "peach"}
    fmt.Println("strings.Join(s1, \", \")",strings.Join(s1, ", "))

    // func Split(s, sep string) [ ]string
    s2 := "dog cat lion"
    fmt.Println("strings.Split(s2, \" \")", strings.Split(s2, " "))

    // func Contains(s, substr string) bool
    s3 := "car truck bike van"
    fmt.Println("Contains(s3, \"truck\")", strings.Contains(s3, "truck")) // true
    fmt.Println("Contains(s3, \"ship\")", strings.Contains(s3, "ship")) // false


    // func HasPrefix(s, prefix string) bool
    s4 := "Mc Carthy"
    fmt.Println("strings.HasPrefix(s4, \"Mc\")", strings.HasPrefix(s4, "Mc"))
    fmt.Println("strings.HasPrefix(s4, \"Miss\")", strings.HasPrefix(s4, "Miss"))

    // func HasSuffix(s, suffix string) bool
    s5 := "Akash Kumar"
    fmt.Println("strings.HasSuffix(s5, \"Kumar\")", strings.HasSuffix(s5, "Kumar"))
    fmt.Println("strings.HasSuffix(s5, \"Lee\")", strings.HasSuffix(s5, "Lee"))
}









//
