# Bookstore-Grpc
BookStore using GOlang with Grpc

create project using:
go mod init bookstore_grpc

install dependencies
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/lib/pq
go get github.com/golang-jwt/jwt/v4
go get github.com/joho/godotenv
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


generate proto files
protoc --go_out=./bookstore_grpc/proto --go-grpc_out=. proto/bookstore.proto

create DB using below commands
"CREATE DATABASE bookstore;"
\c bookstore;
CREATE TABLE users (id SERIAL PRIMARY KEY, username TEXT, password TEXT);
INSERT INTO users (username, password) VALUES ('admin', 'admin123');
CREATE TABLE users (id SERIAL PRIMARY KEY, username TEXT,password TEXT);

run server
go run main.go
