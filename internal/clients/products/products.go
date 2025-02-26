package products

import (
	"context"
	"github.com/codevibe-de/goadv--products/generate/product/pb"
	"google.golang.org/grpc"
)

type ProductClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{
		client: pb.NewProductServiceClient(conn),
	}
}

func (c *ProductClient) GetProduct(ctx context.Context, productID string) (*pb.ProductResponse, error) {
	req := &pb.ProductRequest{ProductId: productID}
	return c.client.GetProduct(ctx, req)
}

func (c *ProductClient) ListProducts(ctx context.Context, productIDs []string) ([]*pb.ProductResponse, error) {
	req := &pb.ProductListRequest{ProductIds: productIDs}
	stream, err := c.client.ListProducts(ctx, req)
	if err != nil {
		return nil, err
	}

	var products []*pb.ProductResponse
	for {
		product, err := stream.Recv()
		if err != nil {
			break
		}
		products = append(products, product)
	}
	return products, nil
}

func (c *ProductClient) CreateProduct(ctx context.Context, productID, name string, price float64) (*pb.ProductResponse, error) {
	req := &pb.ProductCreateRequest{
		ProductId: productID,
		Name:      name,
		Price:     price,
	}
	return c.client.CreateProduct(ctx, req)
}
