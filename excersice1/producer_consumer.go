package main

import ("fmt")
import ("time")

var communication_channel = make(chan int)
var done = make(chan bool)

func produce_tasks() {
    for i:= 0; i<20; i++ {
        fmt.Println("Producer ", i)
        communication_channel<-i
    }
    fmt.Println(communication_channel)
    done <- true
}

func consume_tasks() {
    for {
      msg := <-communication_channel
      fmt.Println("Consumer ", msg)
      time.Sleep(100 * time.Millisecond)
    }
}

func main() {
  go produce_tasks()
  go consume_tasks()
  <-done

}
