package ephemeral

import (
	"log"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func (aws *AWS) Manage() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := info(ctx)
		if err != nil {
			return err
		}
		err = manageS3Bucket(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func manageS3Bucket(ctx *pulumi.Context) error {
	// Create an AWS resource (S3 Bucket)
	bucket, err := createBucket(ctx, "pulumi-spike-demo-bucket")
	if err != nil {
		log.Printf("bucker creation failed. error = %v", err)
		return err
	}
	// Export the name of the bucket
	ctx.Export("bucketName", bucket.ID())
	return nil
}
