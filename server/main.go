package main

//go:generate bash -c "protoc --experimental_allow_proto3_optional --proto_path=../protobuf --go_out=generated/ --go-grpc_out=generated/ $(find ../protobuf -type f -name *.proto)"

import (
	"context"
	"fmt"
	"log"
	"net"

	gppb "github.com/MettyS/checkers/protobuf/game"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) MakeMoves(ctx context.Context, req *gppb.MoveRequest) *gppb.MoveResponse {

	fmt.Println("MakeMoves Running!")

	res := &gppb.MoveResponse{
		move_success: true,
	}

	return res
}

// rpc MakeMoves(MoveRequest) returns (MoveResponse) {}
// rpc BoardUpdateSubscription(BoardSubscriptionRequest) returns (stream BoardUpdate) {}

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	checkersServer := grpc.NewServer()

	s := gppb.server{}
	gppb.RegisterGameplayServiceServer(checkersServer, &s)

	if err := checkersServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
