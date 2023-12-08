// package grpc

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"

// 	"github.com/suresh/notification/internal/notify_pb"
// 	"google.golang.org/grpc"
// )

// type server struct {
// 	notify_pb.UnimplementedHelloServiceServer
// }

// func (*server) SayHello(ctx context.Context, req *notify_pb.HelloRequest) (*notify_pb.HelloResponse, error) {
// 	name := req.GetName()
// 	resp := fmt.Sprintf("hello, my name is %s", name)

// 	return &notify_pb.HelloResponse{Greeting: resp}, nil
// }

// func CreateServer() {
// 	opts := []grpc.ServerOption{}
// 	s := grpc.NewServer(opts...)
// 	notify_pb.RegisterHelloServiceServer(s, &server{})
// 	lis, err := net.Listen("tcp", "0.0.0.0:50051")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("protobuf")
// 	s.Serve(lis)
// }

package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/suresh/notification/internal/notify_pb"
	"google.golang.org/grpc"
)

type server struct {
	notify_pb.UnimplementedHelloServiceServer
}

func (*server) SayHello(ctx context.Context, req *notify_pb.HelloRequest) (*notify_pb.HelloResponse, error) {
	name := req.GetName()
	resp := fmt.Sprintf("Hello, %s!", name)
	return &notify_pb.HelloResponse{Greeting: resp}, nil
}

func CreateServer() {
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	notify_pb.RegisterHelloServiceServer(s, &server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gRPC server is running...")
	s.Serve(lis)
}
