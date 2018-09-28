package main

import (
  "fmt"
)

func main() {
  days := map[string] int { } // also make(map[string] int) can be used
  days["Jan"] = 31 // set 31 as value for key "Jan"
  days["May"] = 31 // set 31 as value for key "May"
  days["Sep"] = 30 // set 30 as value for key "Sep"
  days["Dec"] = 31 // set 31 as value for key "Dec"
  fmt.Println(days)

  numDaysJan := days["Jan"]
  fmt.Println("numDaysJan", numDaysJan)

  numDaysFeb, keyFound := days["Feb"] // we have not set value for key "Feb"
  if keyFound {
    fmt.Println("keyFound == true")
  } else {
    fmt.Println("keyFound == false", "numDaysFeb", numDaysFeb)
    days["Feb"] = 28 // Set the value for key "Feb"
  }

  numDaysFeb, keyFound = days["Feb"]
  if keyFound {
    fmt.Println("keyFound == true", "numDaysFeb", numDaysFeb)
  } else {
    fmt.Println("keyFound == false", "numDaysFeb", numDaysFeb)
    days["Feb"] = 28 // Set the value for key "Feb"
  }
}





















//
