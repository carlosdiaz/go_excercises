package main
//This script will generate elements over a channel and then it will iterate over that channel
import "fmt"
import "time"

// func assignElementsToChannels(inputChannel chan int, value int) {
//     This function will assign some elements into the channel for some
//     reason the for loop doesnt assign the correct value
//     fmt.Println("before assigning value to the channel" , value)
//     //inputChannel<-value
//     for k:=0; k<10;k++ {
//       fmt.Println(k)
//       inputChannel <-value
//     }
// }



func assignElementsToChannels(inputChannel chan int, value int) {
    // This function will assign some elements into the channel
    fmt.Println("before assigning value to the channel" , value)
    inputChannel <-value
    //time.Sleep(100 * time.Millisecond)
}

func readChannelValues(inputChannel chan int) {
    // for channelElement := range inputChannel {
    //     fmt.Println("Reading from the channel ", channelElement)
    //     time.Sleep(10000 * time.Millisecond)
    //
    // }

    for {
      msg := <-inputChannel
      fmt.Println("Consumer ", msg)
      time.Sleep(100 * time.Millisecond)
    }

}


func main() {

    globalChannel := make(chan int)
    for j:=0; j<100; j++ {
      go assignElementsToChannels(globalChannel, j)
      go readChannelValues(globalChannel)
      time.Sleep(100 * time.Millisecond)

    }
}
