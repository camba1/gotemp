package main

import (
	"context"
	"goTemp/product/proto"
)

func (p product) BeforeCreateProduct(ctx context.Context, p2 *proto.Product, err *proto.ValidationErr) error {
	panic("implement me")
}

func (p product) BeforeUpdateProduct(ctx context.Context, p2 *proto.Product, err *proto.ValidationErr) error {
	panic("implement me")
}

func (p product) BeforeDeleteProduct(ctx context.Context, p2 *proto.Product, err *proto.ValidationErr) error {
	panic("implement me")
}

func (p product) AfterCreateProduct(ctx context.Context, p2 *proto.Product, err *proto.AfterFuncErr) error {
	panic("implement me")
}

func (p product) AfterUpdateProduct(ctx context.Context, p2 *proto.Product, err *proto.AfterFuncErr) error {
	panic("implement me")
}

func (p product) AfterDeleteProduct(ctx context.Context, p2 *proto.Product, err *proto.AfterFuncErr) error {
	panic("implement me")
}
