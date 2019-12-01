// Package dev defines a development backend for comet
package dev

import "github.com/hashicorp/go-plugin"

func main() {

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: backendPlugin.Handshake,
		Plugins: map[string]plugin.Plugin{
			// the key here doesn't seem to matter
			"cursor-sdk": &cursorPlugin.CursorPlugin{Impl: &pipeline},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})

}
