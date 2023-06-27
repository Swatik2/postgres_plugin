package main

import (
	"log"
	"net"
	"net/rpc"
)

type DatabaseCredentials struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	User            string `json:"user"`
	Password        string `json:"password"`
	DBName          string `json:"dbname"`
	PluginType      string `json:"pluginType"`
	SourceDirectory string `json:"sourceDirectory"`
}

type MyRPCServer struct{}

func (s *MyRPCServer) GetData(args struct{}, reply *DatabaseCredentials) error {
	// Create and populate the JSON data to be sent
	data := DatabaseCredentials{
		Host:            "localhost",
		Port:            5431,
		User:            "postgres",
		Password:        "postgres",
		DBName:          "postgres",
		PluginType:      "postgres",
		SourceDirectory: "/home/swati/Documents/shared_folder/cred.json",
	}

	*reply = data // Set the reply value
	return nil
}

func main() {
	server := rpc.NewServer()
	myRPCServer := &MyRPCServer{}
	server.Register(myRPCServer)

	// Start the server to listen for RPC requests
	listener, err := net.Listen("tcp", ":3400")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go server.ServeConn(conn)
	}
}
