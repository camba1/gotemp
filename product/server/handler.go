package main

import (
	"context"
	"goTemp/product/proto"
)

func (p product) GetProductById(ctx context.Context, id *proto.SearchId, p2 *proto.Product) error {
	panic("implement me")
}

func (p product) GetProducts(ctx context.Context, params *proto.SearchParams, products *proto.Products) error {
	panic("implement me")
}

func (p product) CreateProduct(ctx context.Context, p2 *proto.Product, response *proto.Response) error {
	panic("implement me")
}

func (p product) UpdateProduct(ctx context.Context, p2 *proto.Product, response *proto.Response) error {
	panic("implement me")
}

func (p product) DeleteProduct(ctx context.Context, id *proto.SearchId, response *proto.Response) error {
	panic("implement me")
}
