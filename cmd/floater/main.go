package main

import (
	"os"

	apiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"

	"cnp.io/clusterlink/cmd/floater/app"
)

func main() {
	ctx := apiserver.SetupSignalContext()
	cmd := app.NewFloaterCommand(ctx)
	code := cli.Run(cmd)
	os.Exit(code)
}
