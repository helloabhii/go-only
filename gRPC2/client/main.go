package main

import (
	"log"

	pb "github.com/helloabhii/grpc2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect : %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	names := &pb.NameList{
		Names: []string{"Abhi", "Binu", "who"},
	}

	// callSayHello(client)  //unary api
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callHelloBidrectionalStream(client, names)
}
