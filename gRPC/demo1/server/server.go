package main

//bidirectional streaming

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"

	proto "grpc/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer //struct can only import in the struct
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer() //engine
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) ServerReply(stream proto.Example_ServerReplyServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.HelloResponse{Reply: "message" + strconv.Itoa(i) + "from server"})
		if err != nil {
			return errors.New("unable to send data from server")
		}
	}
	for {
		req, err := stream.Recv()
		if err == io.EOF { //end of file error
			break
		}
		fmt.Println(req.SomeString)
	}
	return nil //return nil if all goes well
}
