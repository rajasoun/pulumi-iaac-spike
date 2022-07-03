package ephemeral

import (
	"log"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type S3Config struct {
	BucketName string
}

func loadConfig(ctx *pulumi.Context) *S3Config {
	conf := config.New(ctx, "")
	bucketName := conf.Require("bucketName")
	log.Printf("bucketName = %v", bucketName)
	s3Conf := S3Config{
		BucketName: bucketName,
	}
	return &s3Conf
}

func createBucket(ctx *pulumi.Context, bucketName string) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, bucketName, nil)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
