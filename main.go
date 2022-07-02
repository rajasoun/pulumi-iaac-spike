package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := createBucket(ctx)
		if err != nil {
			return err
		}
		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}

func createBucket(ctx *pulumi.Context) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, "pulumi-spike-demo-bucket", nil)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
