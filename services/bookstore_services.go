package services

import (
	"bookstore_grpc/db"
	"bookstore_grpc/middleware"
	"bookstore_grpc/models"
	"bookstore_grpc/proto"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type BookstoreService struct {
	proto.UnimplementedBookstoreServiceServer
}

func (s *BookstoreService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var user models.User
	err := db.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", req.Username).
		Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, err
	}

	if user.Password != req.Password {
		return nil, fmt.Errorf("invalid credentials")
	}

	token, err := middleware.GenerateJWT(user.Username)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{Token: token}, nil
}

func (s *BookstoreService) AddBook(ctx context.Context, req *proto.AddBookRequest) (*proto.AddBookResponse, error) {
	_, err := db.DB.Exec("INSERT INTO books (title, author) VALUES ($1, $2)", req.Title, req.Author)
	if err != nil {
		return nil, err
	}

	return &proto.AddBookResponse{Message: "Book added successfully"}, nil
}

func (s *BookstoreService) ListBooks(ctx context.Context, req *proto.ListBooksRequest) (*proto.ListBooksResponse, error) {
	rows, err := db.DB.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*proto.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, &proto.Book{
			Id:     int32(b.ID),
			Title:  b.Title,
			Author: b.Author,
		})
	}

	return &proto.ListBooksResponse{Books: books}, nil
}
