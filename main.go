package main

import (
  "fmt"
  "test/greeting"
  "github.com/IBM/sarama"
)

func main() {
	fmt.Println(greeting.PI)
	greeting.PrintGreeting()
}