package server

import (
	"net/http"
	"syscall"

	"github.com/numary/go-libs/sharedlogging"
	"github.com/numary/webhooks/internal/storage/mongo"
	"github.com/numary/webhooks/internal/svix"
	"github.com/numary/webhooks/pkg/mux"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Start(*cobra.Command, []string) error {
	sharedlogging.Infof("env: %+v", syscall.Environ())

	app := fx.New(StartModule(http.DefaultClient))
	app.Run()

	return nil
}

func StartModule(httpClient *http.Client) fx.Option {
	return fx.Module("webhooks server module",
		fx.Provide(
			func() *http.Client { return httpClient },
			mongo.NewConfigStore,
			svix.New,
			newServerHandler,
			mux.NewServer,
		),
		fx.Invoke(register),
	)
}

func register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}