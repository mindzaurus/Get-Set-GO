package main

import (
  "fmt"
)

const (
  GRADE_3 = iota
  GRADE_4
)

// There are 4 students taking 3 subjects in a class
const (
  GEORGE = iota
  LUCY
  KIM
  KUMAR
)

// These students take the toughest classes possible
const (
  LATIN = iota
  MATH
  VIOLIN
)

func studentMarks2D() {
  marks := [4][3]int {} // total marks out of 100

  // good in math bad in latin
  marks[GEORGE][LATIN] = 22 ; marks[GEORGE][MATH] = 98 ; marks[GEORGE][VIOLIN] = 35;

  // good in all subjects
  marks[LUCY][LATIN] = 85; marks[LUCY][MATH] = 95; marks[LUCY][VIOLIN] = 78;

  // good in MATH
  marks[KIM][LATIN] = 76; marks[KIM][MATH] = 100; marks[KIM][VIOLIN] = 46;

  // good in math and latin
  marks[KUMAR][LATIN] = 83; marks[KUMAR][MATH] = 92; marks[KUMAR][VIOLIN] = 48;

  fmt.Println(marks)
}

func studentMarks3D() {
  marks := [2][4][3]int {} // total marks out of 100
  marks[GRADE_3][GEORGE][LATIN] = 100 // George from grade 3 scored 100 in latin
  marks[GRADE_4][KUMAR][VIOLIN] = 100 // Kumar from grade 4 scored 100 in violin
  fmt.Println(marks)
}

func newMarks2D() {
  marks := new([4][3]int)
  marks[GEORGE][LATIN] = 56 ; marks[GEORGE][MATH] = 57 ; marks[GEORGE][VIOLIN] = 58;
  marks[LUCY][LATIN] = 59; marks[LUCY][MATH] = 60; marks[LUCY][VIOLIN] = 59;
  marks[KIM][LATIN] = 58; marks[KIM][MATH] = 57; marks[KIM][VIOLIN] = 56;
  marks[KUMAR][LATIN] = 57; marks[KUMAR][MATH] = 58; marks[KUMAR][VIOLIN] = 59;
  fmt.Println(marks)
}

func newMarks3D() {
  marks := new([2][4][3]int) // total marks out of 100
  marks[GRADE_3][GEORGE][LATIN] = 60 // George from grade 3 scored 60 in latin
  marks[GRADE_4][KUMAR][VIOLIN] = 60 // Kumar from grade 4 scored 60 in violin
  fmt.Println(marks)
}

func main() {
  studentMarks2D()
  fmt.Println("")
  studentMarks3D()
  fmt.Println("")
  newMarks2D()
  fmt.Println("")
  newMarks3D()
}




//
