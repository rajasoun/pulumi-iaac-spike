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
		err = manageSecrets(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func manageS3Bucket(ctx *pulumi.Context) error {
	s3Conf := loadS3Config(ctx)
	// Create an AWS resource (S3 Bucket)
	bucket, err := createBucket(ctx, s3Conf.BucketName)
	if err != nil {
		log.Printf("bucket creation failed. error = %v", err)
		return err
	}
	// Export the name of the bucket
	ctx.Export("bucketName", bucket.ID())
	return nil
}

func manageSecrets(ctx *pulumi.Context) error {
	secretsConfig := loadSecretsConfig(ctx)
	// Create an AWS resource (S3 Bucket)
	secret, err := createSecret(ctx, secretsConfig)
	if err != nil {
		log.Printf("Secret creation failed. error = %v", err)
		return err
	}
	// Export the ID (in this case the ARN) of the secret
	ctx.Export("secretContainerName", secret.Name)
	ctx.Export("secretContainer", secret.ID())
	return nil
}
