package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/jankremlacek/go-bazel/proto/servicea"
	"github.com/jankremlacek/go-bazel/shared"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	var (
		err error
		lis net.Listener
		srv *grpc.Server
	)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if len(os.Args) != 2 {
		logger.Fatal("Use: servicea [tcp-addr]")
	}
	if lis, err = net.Listen("tcp", os.Args[1]); err != nil {
		logger.Fatal("Unable to bind gRPC",
			zap.String("addr", os.Args[1]),
			zap.Error(err))
	}

	srv = grpc.NewServer()
	pb.RegisterServiceAServer(srv, new(Handler))

	go func() {
		if err := srv.Serve(lis); err != nil {
			logger.Fatal("Error serving gRPC", zap.Error(err))
		}
	}()

	logger.Info("Listening",
		zap.String("at", os.Args[1]),
		zap.String("version", shared.Version()))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	srv.GracefulStop()
}
