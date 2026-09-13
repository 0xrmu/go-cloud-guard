package main

import (
	"context"
	"fmt"

	"github.com/0xrmu/go-cloud-guard/pkg/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.Background()
	s3client, err := aws.NewS3Client(ctx)
	if err != nil {
		panic(err)
	}

	result, err := s3client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	for _, buckets := range result.Buckets {
		fmt.Println(*buckets.Name)
	}

}
