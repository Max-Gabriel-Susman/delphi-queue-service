package training

import (
	"context"
	"flag"
	"log"

	pb "github.com/Max-Gabriel-Susman/delphi-training-service/training"
)

const (
	defaultName = "world"
)

var (
	// addr = flag.String("addr", "10.96.0.3:50052", "the address to connect to")
	// addr = flag.String("addr", "10.100.0.3:50052", "the address to connect to")
	addr = flag.String("addr", "localhost:50053", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type Server interface {
	SayHello(context.Context, *pb.TrainModelRequest) (*pb.TrainModelReply, error)
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

type TrainingServer struct {
	Server server
}

func NewTrainingServer() *TrainingServer {
	return &TrainingServer{}
}

// TrainModel <behaviors>
func (s *server) TrainModel(ctx context.Context, in *pb.TrainModelRequest) (*pb.TrainModelReply, error) {
	log.Printf("Received: %v", in.GetName())

	// training logic

	return &pb.TrainModelReply{Message: "Hello " + in.GetName()}, nil
	// return &pb.HelloReply{Message: "Hello world"}, nil
}
