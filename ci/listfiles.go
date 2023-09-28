package main

import (
    "context"
    "fmt"
    "log"
    "os"
     "fmt"
     "math/rand"
     "time"

    "dagger.io/dagger"
)

func main() {
    ctx := context.Background()

    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
    if err != nil {
        log.Println(err)
        return
    }
    defer client.Close()

    entries, err := client.Host().Directory(".").Entries(ctx)
    if err != nil {
        log.Println(err)
        return
    }

    fmt.Println(entries)
    rand.Seed(time.Now().UnixNano())

     for i := 1; i <= 10; i++ {
      // Generate random echo
      echo := generateRandomEcho()

      // Print the echo
      fmt.Printf("Loop %d: %s\n", i, echo)

      // Sleep for 1 second
      time.Sleep(1 * time.Second)
     }
}



func generateRandomEcho() string {
     // List of possible echo messages
     echos := []string{
      "Rafał",
      "Szymon",
      "Dzordz",
      "Darek",
      "Michał",
     }
    
     // Generate a random index
     index := rand.Intn(len(echos))
    
     // Return the random echo message
     return echos[index]
}

