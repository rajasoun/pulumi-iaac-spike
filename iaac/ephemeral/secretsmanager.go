package ephemeral

import (
	"log"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type SecertsManager struct {
	SecertsContainerName string
	SecretName           string
	Secret               string
}

func loadSecretsConfig(ctx *pulumi.Context) *SecertsManager {
	conf := config.New(ctx, "")
	secretName := conf.Require("secretName")
	secertsContainerName := conf.Require("secertsContainerName")
	secert := conf.Require("secert")

	secertsManager := SecertsManager{
		SecertsContainerName: secertsContainerName,
		SecretName:           secretName,
		Secret:               secert,
	}
	return &secertsManager
}

func createSecret(ctx *pulumi.Context, sm *SecertsManager) (*secretsmanager.Secret, error) {
	secret, err := secretsmanager.NewSecret(ctx, sm.SecertsContainerName, nil)
	if err != nil {
		log.Printf("secret creation failed err = %v", err)
		return nil, err
	}
	secretVersion, err := secretsmanager.NewSecretVersion(ctx, sm.Secret, &secretsmanager.SecretVersionArgs{
		SecretId:     secret.ID(),
		SecretString: pulumi.String(sm.SecretName),
	})
	if err != nil {
		log.Printf("secret version creation failed err = %v", err)
		return nil, err
	}
	log.Printf("secretName = %v ,version = %v", sm.SecretName, secretVersion)
	return secret, nil
}
