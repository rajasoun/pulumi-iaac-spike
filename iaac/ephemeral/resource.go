package ephemeral

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func (aws *AWS) Manage() {
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
