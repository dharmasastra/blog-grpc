package main

import (
	"context"
	"log"

	pb "github.com/dharmasastra/blog-grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("Update blog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Dharmas",
		Title:    "My first blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
