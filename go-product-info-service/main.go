// Execute protoc -I proto proto/product_info.proto  --go_out=plugins=grpc:ecommerce
// Execute go run productinfo_service

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	pb "productinfo/server/ecommerce"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	pb.ProductInfoServer
	productMap map[string]*pb.ProductDetail
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context,
	in *pb.ProductDetail) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}

	if in.Id == "" {
		return nil, status.Errorf(codes.Internal, "product Id is empty")
	}

	// ***** Reading Metadata from Client *****
	md, metadataAvailable := metadata.FromIncomingContext(ctx)
	if !metadataAvailable {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["userid"]; ok {
		fmt.Printf("userid from metadata:\n")
		fmt.Printf("====> Metadata %s\n", t)
	}

	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.ProductDetail)
	}
	s.productMap[in.Id] = in
	log.Printf("Product %v : %v - Added.", in.Id, in.Name)
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.ProductDetail, error) {
	product, exists := s.productMap[in.Value]
	if exists && product != nil {
		log.Printf("Product %v : %v - Retrieved.", product.Id, product.Name)
		return product, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func (s *server) GetRepeatedTypes(ctx context.Context, in *pb.RepeatedTypesMessage) (*pb.RepeatedTypesMessage, error) {
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
