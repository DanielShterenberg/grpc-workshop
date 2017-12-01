# gRPC-workshop
Based on the official gRPC workshop - https://codelabs.developers.google.com/codelabs/cloud-grpc/index.html

Uses Go server instead of Node.js server.

## Get the initial files

* Clone this project. The `init-files` folder contains the file `client.go`, a command-line client for interacting with the gRPC service that that will be created in this codelab.
* From `init-files` folder run `go run client.go`.
This command will fail since we don't have a gRPC service running.

## Step 1: List all books

Create a file called books.proto under `/books` directory and add the following:
```proto
syntax = "proto3";

package books;

service BookService {
  rpc List (Empty) returns (Empty) {}
}

message Empty {}
```
This defines a new service called BookService using the proto3 version of the protocol buffers language. This is the latest version of protocol buffers and is recommended for use with gRPC.

From this proto file we will generate Go file that wraps the gRPC connection for us.
The generated files contain structs from all the "messages" defined in the proto files, and getters and setters to all structs.
Also, generated files contain gRPC client and server wrappers for the service.
  

To generate the Go files from the proto file we need to use the following command:

`protoc -I . books/books.proto --go_out=plugins=grpc:.`

* `-I` indicates the path of the project the proto file is in (“.” means current directory, because we run it from the directory “start”).

* `--go_out=plugins=grpc:` indicates the path of the output. “.” means current directory. This is relative to the laction of the proto file. If the proto file is in books directory then the generated file will also be in the same directory if we use “.”.


Now, create a file called server.go and this to the file:
```go
package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	books "github.com/noaleibo1/grpc-workshop/start/books"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	fmt.Println("Server running at http://0.0.0.0:50051")
	service := service{}
	books.RegisterBookServiceServer(grpcServer, &service)
	grpcServer.Serve(lis)
}

type service struct {
}

func (s *service) List(context.Context, *books.Empty) (*books.Empty, error){
	return &books.Empty{}, status.Error(codes.Unimplemented, "The server does not implement this method")
}
```
Run `go run server.go` and from another terminal tab run `go run client.go`. The error we receive now is ``rpc error: code = Unimplemented desc = The server does not implement this method``.
This means we created a gRPC connection :) We just need to fix the List method.

Edit the files as following:

books.proto:
```proto
syntax = "proto3";

package books;

service BookService {
  rpc List (Empty) returns (BookList) {}
}

message Empty {} 

message Book {
  int32 id = 1;
  string title = 2;
  string author = 3;
}

message BookList {
  repeated Book books = 1;
}
```
server.go:
```go
package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	books "github.com/noaleibo1/grpc-workshop/step-1-list-books/books"
	"golang.org/x/net/context"
)

var (
	port = flag.Int("port", 50051, "The server port")
	booksList = []*books.Book{
		{
			Id: 123,
			Title: "A Tale of Two Cities",
			Author: "Charles Dickens",
		},
	}
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	fmt.Println("Server running at http://0.0.0.0:50051")
	service := service{}
	books.RegisterBookServiceServer(grpcServer, &service)
	grpcServer.Serve(lis)
}

type service struct {
}

func (s *service) List(context.Context, *books.Empty) (*books.BookList, error){
	return &books.BookList{Books: booksList}, nil
}
```

Run `go run server.go` and from another terminal tab run `go run client.go`.
You should now see this book listed!
```commandline
Server sent 1 book(s).
{
  "books": [
    {
      "id": 123,
      "title": "A Tale of Two Cities",
      "author": "Charles Dickens"
    }
  ]
}
```

## Step 2: Insert new books

Edit `books.proto`:
```proto
service BookService {
  rpc List (Empty) returns (BookList) {}
  // add the following line
  rpc Insert (Book) returns (Empty) {}
}
```

Now add the function to `server.go` as well:
```go
func (s *service) Insert(ctx context.Context, book *books.Book) (*books.Empty, error){
	booksList = append(booksList, book)
	return &books.Empty{}, nil
}
```
To test this, restart the node server and then run the go gRPC command-line client's insert command, passing id, title, and author as arguments:
```commandline
go run client.go insert 2 "The Three Musketeers" "Alexandre Dumas"
```

You should see an empty response:
```commandline
Server response:
{}
```

To verify that the book was inserted, run the list command again to see all books:
```commandline
go run client.go list
```

You should now see 2 books listed!
```commandline
Server sent 2 book(s).
{
  "books": [
    {
      "id": 123,
      "title": "A Tale of Two Cities",
      "author": "Charles Dickens"
    },
    {
      "id": 2,
      "title": "The Three Musketeers",
      "author": "Alexandre Dumas"
    }
  ]
}
```

## Step 3: Get and delete books
In this step you will write the code to get and delete Book objects by id via the gRPC service.

### Get a book
To begin, edit books.proto and update BookService with the following:

`books.proto`
```proto
service BookService {
  rpc List (Empty) returns (BookList) {}
  rpc Insert (Book) returns (Empty) {}
  // add the following line
  rpc Get (BookIdRequest) returns (Book) {}
}

// add the message definition below
message BookIdRequest {
  int32 id = 1;
}
```

This defines a new Get rpc call that takes a BookIdRequest as its request and returns a Book as its response.

A BookIdRequest message type is defined for requests containing only a book's id.

To implement the Get method in the server, edit server.go and add the following get handler function:

`server.go`
```go
func (s *service) Get(ctx context.Context, req *books.BookIdRequest) (*books.Book, error){
	for i := 0; i < len(booksList); i++ {
		if booksList[i].Id == req.Id {
			return booksList[i], nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Not found")
}
```

To test this, restart the node server and then run the go gRPC command-line client's get command, passing id as an argument:
```commandline
go run client.go get 123
```

You should see the book response!
```commandline
Server response:
{
  "id": 123,
  "title": "A Tale of Two Cities",
  "author": "Charles Dickens"
}
```

Now try getting a book that doesn't exist:
```commandline
go run client.go get 404
```

You should see the error message returned:
```commandline
Get book (404): rpc error: code = NotFound desc = Not found
```

### Delete a book

Now you will write the code to delete a book by id.

Edit books.proto and add the following Delete rpc method:

`books.proto`
```proto
service BookService {
  // ...
  // add the delete method definition
  rpc Delete (BookIdRequest) returns (Empty) {}
}
```

Now edit `server.go` and add the following delete handler function:

`server.go`
```go
func (s *service) Delete (ctx context.Context, req *books.BookIdRequest) (*books.Empty, error) {
	for i := 0; i < len(booksList); i++ {
		if booksList[i].Id == req.Id {
			booksList = append(booksList[:i], booksList[i+1:]...)
			return &books.Empty{}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Not found")
}
```

If the books array contains a book with the id requested, the book is removed, otherwise a NOT_FOUND error is returned.

To test this, restart the node server and then run the go gRPC command-line client to delete a book:
```commandline
go run client.go list
Server sent 1 book(s).
{
  "books": [
    {
      "id": 123,
      "title": "A Tale of Two Cities",
      "author": "Charles Dickens"
    }
  ]
}

go run client.go delete 123
Server response:
{}

go run client.go list
Server sent 0 book(s).
{}

go run client.go delete 123
Delete book (123): rpc error: code = 5 desc = "Not found"
```

Great!

You implemented a fully functioning gRPC service that can list, insert, get, and delete books!

## Step 4: Stream added books

In this step you will write the code to add a streaming endpoint to the service so the client can establish a stream to the server and listen for added books.
gRPC supports streaming semantics, where either the client or the server (or both) send a stream of messages on a single RPC call. The most general case is Bidirectional Streaming where a single gRPC call establishes a stream where both the client and the server can send a stream of messages to each other.
To begin, edit books.proto and add the following Watch rpc method to BookService:

`books.proto`
```proto
service BookService {
  // ...
  // add the watch method definition
  rpc Watch (Empty) returns (stream Book) {}
}
```

When the client calls the Watch method, it will establish a stream and server will be able to stream Book messages when books are inserted.