package main

func runServer() {
  // Severe error happened here
  // Complete filesystem became unreadable - not able to read any file
  panic("Complete filesystem became unreadable - not able to read any file")
}

func main() {
  runServer()
}
