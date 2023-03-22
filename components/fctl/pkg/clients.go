package fctl

import (
	"fmt"

	"github.com/formancehq/fctl/membershipclient"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/spf13/cobra"
)

func NewMembershipClient(cmd *cobra.Command, cfg *Config) (*membershipclient.APIClient, error) {
	profile := GetCurrentProfile(cmd, cfg)
	httpClient := GetHttpClient(cmd)
	configuration := membershipclient.NewConfiguration()
	token, err := profile.GetToken(cmd.Context(), httpClient)
	if err != nil {
		return nil, err
	}
	configuration.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	configuration.HTTPClient = httpClient
	configuration.Servers[0].URL = profile.GetMembershipURI()
	return membershipclient.NewAPIClient(configuration), nil
}

func NewStackClient(cmd *cobra.Command, cfg *Config, stack *membershipclient.Stack) (*formance.Formance, error) {
	profile := GetCurrentProfile(cmd, cfg)
	httpClient := GetHttpClient(cmd)

	token, err := profile.GetStackToken(cmd.Context(), httpClient, stack)
	if err != nil {
		return nil, err
	}

	apiConfig := formance.New(
		formance.WithServerURL(stack.Uri),
		formance.WithSecurity(shared.Security{
			Authorization: fmt.Sprintf("Bearer %s", token),
		}),
	)

	return apiConfig, nil
}
