package controllers

import (
	"bookstore_grpc/proto"
	"bookstore_grpc/services"
)

func NewBookstoreController() proto.BookstoreServiceServer {
	return &services.BookstoreService{}
}
