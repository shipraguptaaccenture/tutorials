package main

import "fmt"
import "github.com/ably/ably-go/ably"

func main() {
    client, err := ably.NewRealtimeClient(ably.NewClientOptions("INSERT-YOUR-ABLY-KEY-HERE"))
    if err != nil {
      panic(err)
    }

	  channel := client.Channels.Get("sport")

	  sub, err := channel.Subscribe()
	  if err != nil {
	      panic(err)
	  }

	  res, err := channel.Publish("team", "Man United")
	  // await confirmation
	  if err = res.Wait(); err != nil {
	      panic(err)
	  }

	  for msg := range sub.MessageChannel() {
		  	fmt.Printf("Received message with name '%v' and data '%v'\n", msg.Name, msg.Data)
		}
}
