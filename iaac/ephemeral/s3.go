package ephemeral

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createBucket(ctx *pulumi.Context, bucketName string) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, bucketName, nil)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
