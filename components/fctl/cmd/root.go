package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"

	"github.com/formancehq/fctl/cmd/auth"
	"github.com/formancehq/fctl/cmd/cloud"
	"github.com/formancehq/fctl/cmd/ledger"
	"github.com/formancehq/fctl/cmd/profiles"
	"github.com/formancehq/fctl/cmd/search"
	"github.com/formancehq/fctl/cmd/stack"
	fctl "github.com/formancehq/fctl/pkg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	MaxVersionShift = 2
)

func NewRootCommand() *cobra.Command {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cmd := fctl.NewCommand("fctl",
		fctl.WithSilenceError(),
		fctl.WithShortDescription("Formance Control CLI"),
		fctl.WithSilenceUsage(),
		fctl.WithChildCommands(
			NewUICommand(),
			NewVersionCommand(),
			NewLoginCommand(),
			NewPromptCommand(),
			profiles.NewCommand(),
			stack.NewCommand(),
			cloud.NewCommand(),
			auth.NewCommand(),
			ledger.NewCommand(),
			search.NewCommand(),
			//payments.NewCommand(),
			//orchestration.NewCommand(),
			//webhooks.NewCommand(),
			//wallets.NewCommand(),
		),
		fctl.WithPersistentStringPFlag(fctl.ProfileFlag, "p", "", "config profile to use"),
		fctl.WithPersistentStringPFlag(fctl.FileFlag, "c", fmt.Sprintf("%s/.formance/fctl.config", homedir), "Debug mode"),
		fctl.WithPersistentBoolPFlag(fctl.DebugFlag, "d", false, "Debug mode"),
		fctl.WithPersistentBoolFlag(fctl.InsecureTlsFlag, false, "Insecure TLS"),
		fctl.WithPersistentBoolFlag(fctl.TelemetryFlag, false, "Telemetry enabled"),
	)
	return cmd
}

func Execute() {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
		}
	}()

	ctx, _ := signal.NotifyContext(context.TODO(), os.Interrupt)
	err := NewRootCommand().ExecuteContext(ctx)
	if err != nil {
		switch {
		//case errors.Is(err, fctl.ErrMissingApproval):
		default:
			pterm.Error.WithWriter(os.Stderr).Printfln(err.Error())
			pterm.Error.WithWriter(os.Stderr).Printfln("Command aborted as you didn't approve.")
			os.Exit(1)
			//case extractOpenAPIErrorMessage(err) != nil:
			//	pterm.Error.WithWriter(os.Stderr).Printfln(extractOpenAPIErrorMessage(err).Error())
			//	os.Exit(2)
			//default:
			//	pterm.Error.WithWriter(os.Stderr).Printfln(err.Error())
			//	os.Exit(255)
		}
	}
}

//
//func extractOpenAPIErrorMessage(err error) error {
//	if err == nil {
//		return nil
//	}
//
//	if err := unwrapOpenAPIError(err); err != nil {
//		return errors.New(err.GetErrorMessage())
//	}
//
//	return err
//}
//
//func unwrapOpenAPIError(err error) *shared.ErrorResponse {
//	for err != nil {
//		if err, ok := err.(*shared.GenericOpenAPIError); ok {
//			body := err.Body()
//			// Actually, each api redefine errors response
//			// So OpenAPI generator generate an error structure for every service
//			// Manually unmarshal errorResponse allow us to handle only one ErrorResponse
//			// It will be refined once the monorepo fully ready
//			errResponse := api.ErrorResponse{}
//			if err := json.Unmarshal(body, &errResponse); err != nil {
//				return nil
//			}
//
//			if errResponse.ErrorCode != "" {
//				errorCode := formance.ErrorsEnum(errResponse.ErrorCode)
//				return &formance.ErrorResponse{
//					ErrorCode:    &errorCode,
//					ErrorMessage: &errResponse.ErrorMessage,
//					Details:      &errResponse.Details,
//				}
//			}
//		}
//		err = errors.Unwrap(err)
//	}
//
//	return nil
//}
