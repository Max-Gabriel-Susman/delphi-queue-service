package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Max-Gabriel-Susman/delphi-training-service/internal/training"
	pb "github.com/Max-Gabriel-Susman/delphi-training-service/training"
	"google.golang.org/grpc"
)

/*
	TODOs:
		META:
			* start a documentation direcory
			* start a cloud formation directory
			* start implementing testing coverage
			* work more on readme
			* abstract what we can to delphi-go-kit (e.g. logging, tracing, etc.)
			* determine what logging tracing solutions I want to use long term(probably just something within aws honestly)
			* refactor rootlevel protobuf/grpc logic into corresponding
				internal directories
			* refactor main.go to cmd/delphi-x-service/main.go
			* clean up Make targets and keep them up to date
		MESA:
*/

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

var (
	port = flag.Int("port", 50055, "The server port") // actual port dictation
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitCodeErr)
	}
}

func run(ctx context.Context, _ []string) error {
	// Start GRPC Service
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ts := training.NewTrainingServer()
	pb.RegisterGreeterServer(s, &ts.Server)
	log.Printf("server listening at %v", lis.Addr())

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return nil
}
