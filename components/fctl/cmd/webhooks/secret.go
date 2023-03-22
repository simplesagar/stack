package webhooks

import (
	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewChangeSecretCommand() *cobra.Command {
	return fctl.NewCommand("change-secret <config-id> <secret>",
		fctl.WithShortDescription("Change the signing secret of a config. You can bring your own secret. If not passed or empty, a secret is automatically generated. The format is a string of bytes of size 24, base64 encoded. (larger size after encoding)"),
		fctl.WithConfirmFlag(),
		fctl.WithAliases("cs"),
		fctl.WithArgs(cobra.RangeArgs(1, 2)),
		fctl.WithRunE(func(cmd *cobra.Command, args []string) error {
			cfg, err := fctl.GetConfig(cmd)
			if err != nil {
				return errors.Wrap(err, "fctl.GetConfig")
			}

			organizationID, err := fctl.ResolveOrganizationID(cmd, cfg)
			if err != nil {
				return err
			}

			stack, err := fctl.ResolveStack(cmd, cfg, organizationID)
			if err != nil {
				return err
			}

			if !fctl.CheckStackApprobation(cmd, stack, "You are about to change a webhook secret") {
				return fctl.ErrMissingApproval
			}

			client, err := fctl.NewStackClient(cmd, cfg, stack)
			if err != nil {
				return errors.Wrap(err, "creating stack client")
			}

			secret := ""
			if len(args) > 1 {
				secret = args[1]
			}

			request := operations.ChangeConfigSecretRequest{
				ID: args[0],
				ConfigChangeSecret: &shared.ConfigChangeSecret{
					Secret: &secret,
				},
			}
			res, err := client.Webhooks.ChangeConfigSecret(cmd.Context(), request)
			if err != nil {
				return errors.Wrap(err, "changing secret")
			}

			pterm.Success.WithWriter(cmd.OutOrStdout()).Printfln(
				"Config updated successfully with new secret: %s", *res.ConfigResponse.Data)
			// TODO: Need to return only secret
			return nil
		}),
	)
}
