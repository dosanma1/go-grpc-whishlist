package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/dosanma1/go-grpc-wishlist/config"
	"github.com/dosanma1/go-grpc-wishlist/internal/pb"
	"github.com/dosanma1/go-grpc-wishlist/internal/server"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedWishListServiceServer

	config *config.Config
}

func NewGrpcServer(cfg *config.Config) server.Server {
	return &grpcServer{
		config: cfg,
	}
}

func (s *grpcServer) Create(ctx context.Context, req *pb.CreateWishListReq) (*pb.CreateWishListResp, error) {
	fmt.Println("creating the wish list " + req.WishList.Name)
	return &pb.CreateWishListResp{
		WishListId: req.WishList.Id,
	}, nil
}
func (s *grpcServer) Add(ctx context.Context, req *pb.AddItemReq) (*pb.AddItemResp, error) {
	return nil, nil
}
func (s *grpcServer) List(ctx context.Context, req *pb.ListWishListReq) (*pb.ListWishListResp, error) {
	return nil, nil
}

func (s *grpcServer) Run() error {
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	listener, err := net.Listen(s.config.Protocol, addr)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	pb.RegisterWishListServiceServer(srv, &grpcServer{})

	if err := srv.Serve(listener); err != nil {
		return err
	}
	defer srv.GracefulStop()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		done <- true
	}()

	<-done

	return nil
}
