package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	proto "grpc/protoc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	//connection to internal grpc server
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(conn)

	//implement rest api

	r := gin.Default()
	r.GET("/sent", clientConnectionServer)
	r.Run(":8000") //by default at 8080 if you didn't give location/path

}

func clientConnectionServer(c *gin.Context) {
	stream, err := client.ServerReply(context.TODO())
	if err != nil {
		fmt.Println("Something Error")
		return
	}
	send, receive := 0, 0
	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.HelloRequest{SomeString: "message" + strconv.Itoa(i) + "from client"})
		if err != nil {
			fmt.Println("unable to send data")
			return
		}
		send++
	}
	if err := stream.CloseSend(); err != nil {
		log.Println(err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("Server Message :- ", message)
		time.Sleep(3 * time.Second)
		receive++
	}
	c.JSON(http.StatusOK, gin.H{
		"message send":     send,
		"message received": receive,
	})
}
