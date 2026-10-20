package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-polyglot-lab/go-server/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserProfile(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	userID := req.GetUserId()
	log.Printf("Menerima request untuk User ID: %d", userID)

	if userID == 1 {
		return &pb.UserResponse{
			Name:  "Budi Santoso",
			Email: "budi@universitas.ac.id",
			Phone: "081234567890",
		}, nil
	}

	return &pb.UserResponse{
		Name:  "User Tidak Ditemukan",
		Email: "-",
		Phone: "-",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal membuka port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	fmt.Println("gRPC Server (Go) berjalan di port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
