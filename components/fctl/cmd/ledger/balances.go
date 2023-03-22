package ledger

import (
	"fmt"

	"github.com/formancehq/fctl/cmd/ledger/internal"
	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewBalancesCommand() *cobra.Command {
	const (
		afterFlag   = "after"
		addressFlag = "address"
	)
	return fctl.NewCommand("balances",
		fctl.WithAliases("balance", "bal", "b"),
		fctl.WithStringFlag(addressFlag, "", "Filter on specific address"),
		fctl.WithStringFlag(afterFlag, "", "Filter after specific address"),
		fctl.WithShortDescription("Read balances"),
		fctl.WithArgs(cobra.ExactArgs(0)),
		fctl.WithRunE(func(cmd *cobra.Command, args []string) error {
			cfg, err := fctl.GetConfig(cmd)
			if err != nil {
				return err
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
				return err
			}

			request := operations.GetBalancesRequest{
				Ledger:  fctl.GetString(cmd, internal.LedgerFlag),
				After:   formance.String(fctl.GetString(cmd, afterFlag)),
				Address: formance.String(fctl.GetString(cmd, addressFlag)),
			}
			balances, err := client.Ledger.GetBalances(cmd.Context(), request)
			if err != nil {
				return err
			}

			tableData := pterm.TableData{}
			tableData = append(tableData, []string{"Account", "Asset", "Balance"})
			for _, accountBalances := range balances.BalancesCursorResponse.Cursor.Data {
				for account, volumes := range accountBalances {
					for asset, balance := range volumes {
						tableData = append(tableData, []string{
							account, asset, fmt.Sprint(balance),
						})
					}
				}
			}
			if err := pterm.DefaultTable.
				WithHasHeader(true).
				WithWriter(cmd.OutOrStdout()).
				WithData(tableData).
				Render(); err != nil {
				return err
			}

			return nil
		}),
	)
}
