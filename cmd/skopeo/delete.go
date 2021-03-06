package main

import (
	"errors"
	"fmt"

	"github.com/containers/image/transports"
	"github.com/urfave/cli"
)

func deleteHandler(context *cli.Context) error {
	if len(context.Args()) != 1 {
		return errors.New("Usage: delete imageReference")
	}

	ref, err := transports.ParseImageName(context.Args()[0])
	if err != nil {
		return fmt.Errorf("Invalid source name %s: %v", context.Args()[0], err)
	}

	ctx, err := contextFromGlobalOptions(context, "")
	if err != nil {
		return err
	}
	if err := ref.DeleteImage(ctx); err != nil {
		return err
	}
	return nil
}

var deleteCmd = cli.Command{
	Name:      "delete",
	Usage:     "Delete image IMAGE-NAME",
	ArgsUsage: "IMAGE-NAME",
	Action:    deleteHandler,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "creds",
			Value: "",
			Usage: "Use `USERNAME[:PASSWORD]` for accessing the registry",
		},
		cli.StringFlag{
			Name:  "cert-dir",
			Value: "",
			Usage: "use certificates at `PATH` (*.crt, *.cert, *.key) to connect to the registry",
		},
		cli.BoolTFlag{
			Name:  "tls-verify",
			Usage: "require HTTPS and verify certificates when talking to docker registries (defaults to true)",
		},
	},
}
