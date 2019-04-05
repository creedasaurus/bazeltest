package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"google.golang.org/appengine/file"
)

func main() {
	c := context.Background()
	name, err := file.DefaultBucketName(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)

	client, err := storage.NewClient(c, option.WithCredentialsFile("asdfasdf"))
	if err != nil {
		log.Fatal(err)
	}

	dest := client.Bucket("creeds-pubsub").Object("asdf.jpeg")
	destAttrs, err := dest.Attrs(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(destAttrs)
}
