package main

import (
	"context"
	"log"

	pb "github.com/dharmasastra/blog-grpc/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("Delete blog was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}