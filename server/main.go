package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"posts_service/proto"
	"posts_service/server/controller"
)

type server struct{}


func main() {

	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	srv := grpc.NewServer(opts...)

	proto.RegisterPostsServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
	fmt.Println("Server is running....")
}



// errorString is a trivial implementation of error.
/*type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}*/

func (s *server) CreatePost(ctx context.Context, request *proto.CreatePostReq) (*proto.CreatePostRes, error) {

	err := controller.CreatePost(controller.Post{Body:request.GetPost().GetBody()})

	if err != nil {
		panic(err)
	}

	return &proto.CreatePostRes{Message:"Saved post"}, nil
}

func (s *server) ReadPosts(ctx context.Context, request *proto.ListPostsReq) (*proto.ListPostsRes, error) {

	//items := make([]*proto.Post, 0)
	items := []*proto.Post{};

	err, posts := controller.GetPosts()
	if err != nil {
		panic(err)
	}

	for _, p := range posts {
		c := &proto.Post{
			Id : p.ID.Hex(),
			Body: p.Body,
		}
		items = append(items, c)
	}

	return &proto.ListPostsRes{Posts:items}, nil

}

func (s *server) UpdatePost(ctx context.Context, request *proto.UpdatePostReq) (*proto.UpdatePostRes, error) {

	id := request.GetId()

	idd,_ := primitive.ObjectIDFromHex(id)

	err := controller.UpdatePost(idd,controller.Post{Body:request.GetPost().GetBody()})

	if err != nil {
		panic(err)
	}


	return &proto.UpdatePostRes{Message:"Successfully updated"}, nil
}

func (s *server) DeletePost(ctx context.Context, request *proto.DeletePostReq) (*proto.DeletePostRes, error) {
	id := request.GetId()

	idd,_ := primitive.ObjectIDFromHex(id)

	err := controller.DeletePost(idd,ctx)

	if err != nil {
		panic(err)
	}

	return &proto.DeletePostRes{Message: "Successfuly deleted item"}, nil
}

//protoc -I . service2.proto --go_out=plugins=grpc:. *.proto