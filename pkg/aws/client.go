// FINDING #2 [SEVERITY: LOW]
// Key fragment 2 of 5 found hardcoded.
// Hint: the line directly above NewEC2Client
//
// NEXT TARGET: cmd/root.go

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const _ = "IV"
func NewS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	cfg.Region = "eu-north-1"

	s3client := s3.NewFromConfig(cfg)
	return s3client, nil
}

func NewEC2Client(ctx context.Context) (*ec2.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	cfg.Region = "eu-north-1"
	ec2client := ec2.NewFromConfig(cfg)
	return ec2client, nil
}

//RY