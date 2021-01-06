package main

import "fmt"
import "github.com/ably/ably-go/ably"

func main() {
    client, err := ably.NewRealtimeClient(ably.NewClientOptions("INSERT-YOUR-ABLY-KEY-HERE"))
    if err != nil {
      panic(err)
    }
}
