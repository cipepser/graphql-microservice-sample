//go:generate gqlgen -schema ../schema.graphql
package graph

import (
	"github.com/cipepser/graphql-microservice-sample/account"
	"github.com/cipepser/graphql-microservice-sample/catalog"
	"github.com/cipepser/graphql-microservice-sample/order"
)

type GraphQLServer struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func NEwGraphQLServer(accountURL, catalogURL, orderURL string) (*GraphQLServer, error) {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return nil, err
	}

	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return nil, err
	}

	orderClient, err := order.NewClient(orderURL)
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return nil, err
	}

	return &GraphQLServer{
		accountClient, catalogClient, orderClient,
	}, nil
}
