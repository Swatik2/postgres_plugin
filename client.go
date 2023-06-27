package main 

import (
    "fmt"
    "log"
    "net/rpc"
)


type DatabaseCredentials struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
    PluginType      string `json:"pluginType"`
    SourceDirectory string `json:"sourceDirectory"`
}


func main() {
    client, err := rpc.Dial("tcp", "localhost:3400")
    if err != nil {
        log.Fatal("Dial error:", err)
    }

    var data DatabaseCredentials
    err = client.Call("MyRPCServer.GetData", struct{}{}, &data)
    if err != nil {
        log.Fatal("RPC error:", err)
    }

    fmt.Println("Received data:", data)
}
