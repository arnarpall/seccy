![Go](https://github.com/arnarpall/seccy/workflows/Go/badge.svg)

# seccy
Seccy is a key-value server where the keys and values are stored in an encrypted datastore

This little thing was started from the https://gophercises.com/

Seccy exposes gRPC endpoints for setting values and getting a value for a key.

Protobuf client is stored in api/proto/seccy
The protobuf generated client can be imported like so
`import "github.com/arnarpall/seccy/api/proto/seccy"`

# API client
Example client
```go
  package main
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

# Server Development
The project is built using make and targets exist to build the cli and server as
well as the protobuf definition.

## Building with make
`make cli` will build the cli and place it under `cmd/seccy-cli`

`make server` will build the server and place it under `cmd/seccy-service`

`make proto` will build the protobuf definition under `api/proto/seccy`

## Running the server
## Tilt (preferred)
The project has comes with a simple Tiltfile that can deploy the server
for development purposes.
Tilt installation is available [here](https://docs.tilt.dev/install.html)

1. Make sure you have a kubernetes cluster up (kind, minikubet, e.t.c) and the kubeconfig
exported `export KUBECONFIG=/path/to/kubeconfig`

2. Run tilt with `tilt up` this will deploy the helm cart to your local cluster

Tilt will port forward the seccy service to `localhost:4040` locally so the accompanying CLI will
work.

## Local instance with make
A make target exists to run the server in for local development
`make run-server-dev` this will start the server with a dummy
file repository under `/tmp/`

## Running the binary
In order to run the server 2 flags are a requirement, the path to the file store
and the encryption key
```shell script
cmd/seccy-service/seccy-service --encryption-key some-key --store-path /path/to/storefile
```

# CLI
The project has a cli that can be used to interact with the seccy server
run the binary from the `cmd/seccy-cli/` directory

```shell script
   Secrets keeper
   
   Usage:
     seccy [command]
   
   Available Commands:
     get         get a secret
     help        Help about any command
     list        list all keys
     set         set a secret value
   
   Flags:
     -h, --help   help for seccy
   
   Use "seccy [command] --help" for more information about a command.
```

## Setting a value for a key
```shell script
seccy-cli set --key some-key --value "some long value to encrypt"
```

## Getting a value
```shell script
seccy-cli get --key some-key
```

## Listing all keys
```shell script
seccy-cli list
```
