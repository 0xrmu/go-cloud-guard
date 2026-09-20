package main

import (
	"context"
	"fmt"

	"github.com/0xrmu/go-cloud-guard/pkg/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.Background()
	s3client, err := aws.NewS3Client(ctx)
	if err != nil {
		panic(err)
	}

	ec2client, err := aws.NewEC2Client(ctx)
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

	result1, err1 := ec2client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{})
	if err1 != nil {
		panic(err1)
	}

	for _, securityGroup := range result1.SecurityGroups {
		for _, permission := range securityGroup.IpPermissions {
			for _, ipRange := range permission.IpRanges {
				if ipRange.CidrIp != nil && *ipRange.CidrIp == "0.0.0.0/0" {
					fmt.Println("EXPOSED:", *securityGroup.GroupName)
				}
			}
		}
	}
}
