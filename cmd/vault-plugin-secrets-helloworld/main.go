package main

import (
	"os"

	helloworld "github.com/daveadams/vault-plugin-secrets-helloworld/plugin"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/api"
        "github.com/hashicorp/vault/sdk/plugin"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{})

	defer func() {
		if r := recover(); r != nil {
			logger.Error("plugin panic", "error", r)
			os.Exit(1)
		}
	}()

	meta := &api.PluginAPIClientMeta{}

	flags := meta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := meta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	if err := plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: helloworld.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	}); err != nil {
		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}
