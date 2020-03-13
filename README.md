# seccy
Seccy is a key-value server where the keys and values are stored in an encrypted datastore

Seccy exposes gRPC endpoints for setting values and getting a value for a key.


Example client
```golang
  import (
    "fmt"
    
    "github.com/arnarpall/seccy/pkg/client"
  )
  
  func main() {
    client, err := client.New(":4040", logger)
    if err != nil {
       panic("unable to connect to server")
    }

    err = client.Set("this-is", "awesome")
    if err != nil {
      panic("unable to set key")
    }
    
    val, err = client.Get("this-is")
    if err != nil {
      panic("unable to get key")
    }
    
    fmt.Print(v)
}
```
