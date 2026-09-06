package gocloudguard

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(cfg)
}
