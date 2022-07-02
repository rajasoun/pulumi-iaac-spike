package main

import (
	"github.com/rajsoun/aws-manager/iaac/ephemeral"
)

func main() {
	aws := ephemeral.AWS{}
	aws.Manage()
}
