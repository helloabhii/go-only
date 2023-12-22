package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/helloabhii/grpc2/proto"
)

func callHelloBidrectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidrectional Streaming Started")
	stream, err := client.SayHelloBidrectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Couldn't send names: %v", err)
	}
	//use channels
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(3 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidrectional streaming finished")
}
