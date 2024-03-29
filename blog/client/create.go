package main

import (
	"context"
	"log"

	pb "github.com/dharmasastra/blog-grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Creating the blog")

	blog := &pb.Blog{
		AuthorId: "Dharmas",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}