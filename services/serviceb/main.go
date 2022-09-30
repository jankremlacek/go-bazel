package main

import (
	"context"
	"os"

	pb "github.com/jankremlacek/go-bazel/proto/servicea"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		conn     *grpc.ClientConn
		err      error
		response *pb.SumResponse
	)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if len(os.Args) != 2 {
		logger.Fatal("Use: servicea [tcp-addr]")
	}

	logger.Info("Contacting ServiceA",
		zap.String("at", os.Args[1]))

	if conn, err = grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		logger.Fatal("Error dialing",
			zap.String("to", os.Args[1]),
			zap.Error(err))
	}

	if response, err = pb.NewServiceAClient(conn).Sum(context.TODO(), &pb.SumRequest{
		A: 42,
		B: 21,
	}); err != nil {
		logger.Fatal("Unable to call RPC",
			zap.Error(err))
	}

	logger.Info("Got response",
		zap.Int("sum", int(response.Result)))
}
