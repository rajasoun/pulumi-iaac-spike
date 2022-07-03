package ephemeral

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func info(ctx *pulumi.Context) error {
	readmeBytes, err := loadFile("/iaac/ephemeral/README.md")
	if err != nil {
		return err
	}

	ctx.Export("readme", pulumi.String(string(readmeBytes)))
	return nil
}
