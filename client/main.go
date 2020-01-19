package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net/http"
	"posts_service/proto"
	"posts_service/server/controller"
	"time"
)

var client = proto.NewPostsServiceClient(nil)

func main() {

	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")

	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, err := grpc.Dial("localhost:5555", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client = proto.NewPostsServiceClient(conn)

	r := mux.NewRouter()
	r.HandleFunc("/posts/index", IndexPostsHandler).Methods("GET")
	r.HandleFunc("/posts/create", CreatePostHandler).Methods("POST")
	r.HandleFunc("/posts/update/{id}", UpdatePostHandler).Methods("PUT")
	r.HandleFunc("/posts/delete/{id}", DeletePostHandler).Methods("DELETE")

	http.ListenAndServe(":8787", r)

}

func IndexPostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	ctx, _ := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	params := mux.Vars(r)
	id := params["id"]

	fmt.Println(id)

	req := &proto.ListPostsReq{}

	if response, err := client.ReadPosts(ctx, req); err == nil {
		json.NewEncoder(w).Encode(response.GetPosts())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	ctx, _ := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	var post controller.Post
	_ = json.NewDecoder(r.Body).Decode(&post)

	req := &proto.CreatePostReq{Post:&proto.Post{Body:post.Body}}

	if response, err := client.CreatePost(ctx, req); err == nil {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{ "message": "` + response.GetMessage() + `" }`))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	w.Header().Set("content-type", "application/json")

	ctx, _ := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	post := controller.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		panic(err)
	}

	req := &proto.UpdatePostReq{Id:id,Post:&proto.Post{Body:post.Body}}

	if response, err := client.UpdatePost(ctx, req); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "message": "` + response.GetMessage() + `" }`))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {

	ctx, _ := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	params := mux.Vars(r)
	id := params["id"]

	req := &proto.DeletePostReq{Id:id}

	w.Header().Set("content-type", "application/json")

	if response, err := client.DeletePost(ctx, req); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "message": "`+response.GetMessage()+`" }`))
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "Probelm with deleting!" }`))
	}

}
//https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2
//https://jbrandhorst.com/post/grpc-errors/
//https://stackoverflow.com/questions/52969205/how-to-assert-grpc-error-codes-client-side-in-go
//https://askubuntu.com/questions/767934/mongodb-installation-failed-on-ubuntu-16-04
//https://gobyexample.com/variadic-functions
//https://stackoverflow.com/questions/23669720/meaning-of-interface-dot-dot-dot-interface
//https://yourbasic.org/golang/three-dots-ellipsis/

//show dbs
//use mydb
//show collections
//db.movie.insert({"name":"tutorials point"})
//use movie
//db.movie.find()
