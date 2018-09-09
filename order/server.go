//go:generate protoc ./order.proto --go_out=plugins=grpc:./pb
package order

import (
	"context"
	"fmt"
	"net"

	"github.com/cipepser/graphql-microservice-sample/account"
	"github.com/cipepser/graphql-microservice-sample/catalog"
	"github.com/cipepser/graphql-microservice-sample/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return err
	}

	catalogClient, err := account.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &grpcServer{
		s,
		accountClient,
		catalogClient,
	})
	reflection.Register(serv)

	return serv.Serve(lis)
}

//func (s *grpcServer) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {
//	// check if account exists
//	_, err := s.accountClient.GetAccount(ctx, )
//}
//
//func (s *grpcServer) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
//	panic("implement me")
//}
