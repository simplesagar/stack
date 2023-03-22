package balances

import (
	"github.com/formancehq/fctl/cmd/wallets/internal"
	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	return fctl.NewCommand("create <balance-name>",
		fctl.WithShortDescription("Create a new balance"),
		fctl.WithAliases("c", "cr"),
		fctl.WithConfirmFlag(),
		fctl.WithArgs(cobra.ExactArgs(1)),
		internal.WithTargetingWalletByID(),
		internal.WithTargetingWalletByName(),
		fctl.WithRunE(func(cmd *cobra.Command, args []string) error {
			cfg, err := fctl.GetConfig(cmd)
			if err != nil {
				return errors.Wrap(err, "retrieving config")
			}

			organizationID, err := fctl.ResolveOrganizationID(cmd, cfg)
			if err != nil {
				return err
			}

			stack, err := fctl.ResolveStack(cmd, cfg, organizationID)
			if err != nil {
				return err
			}

			client, err := fctl.NewStackClient(cmd, cfg, stack)
			if err != nil {
				return errors.Wrap(err, "creating stack client")
			}

			walletID, err := internal.RequireWalletID(cmd, client)
			if err != nil {
				return err
			}

			request := operations.CreateBalanceRequest{
				ID: walletID,
				CreateBalanceRequest: &shared.CreateBalanceRequest{
					Name: args[0],
				},
			}
			res, err := client.Wallets.CreateBalance(cmd.Context(), request)
			if err != nil {
				return errors.Wrap(err, "Creating wallets")
			}

			pterm.Success.WithWriter(cmd.OutOrStdout()).Printfln(
				"Balance created successfully with name: %s", res.CreateBalanceResponse.Data.Name)
			return nil
		}),
	)
}
