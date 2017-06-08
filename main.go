package main

import (
	"github.com/gorilla/mux"

	"commands"
	"log"
	"net/http"
	"qrcode"

	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	gRPCPort   = ":8086"
	routerPort = ":8085"
)

func gRPCServer() {
	listen, err := net.Listen("tcp", gRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	commands.RegisterGreeterServer(server, &commands.Command{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	fmt.Println("Running gRPC Server")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/qrcode", qrcode.QRCode)

	go gRPCServer()

	fmt.Println("Running Server")
	log.Fatal(http.ListenAndServe(routerPort, router))
}
