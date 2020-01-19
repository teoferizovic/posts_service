package controller

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var client *mongo.Client
var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
var collection *mongo.Collection

type Post struct {
	ID        primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Body 	  string             	`json:"body,omitempty" bson:"body,omitempty"`
}

func init() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	collection = client.Database("mydb").Collection("posts")

}

func CreatePost(post Post) error {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection.InsertOne(ctx, post)
	return nil

}

func GetPosts() (error,[]Post) {

	var posts []Post

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return err, nil
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return err, nil
	}
	return nil, posts
}

func DeletePost(id primitive.ObjectID,ctx context.Context) error {

	// Delete record
	collection.FindOneAndDelete(ctx, Post{ID: id})
	return nil

}

func UpdatePost(id primitive.ObjectID,post Post) error {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	update := bson.M{
		"$set": bson.M{"body": post.Body},
	}

	//update record
	collection.FindOneAndUpdate(ctx,Post{ID: id},update)

	return nil

}