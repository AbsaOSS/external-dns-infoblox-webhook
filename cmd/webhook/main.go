// package main
package main

import (
	"fmt"

	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/configuration"
	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/dnsprovider"
	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/logging"
	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/server"
	"github.com/AbsaOSS/external-dns-infoblox-webhook/pkg/webhook"
	log "github.com/sirupsen/logrus"
)

const banner = `
  
external-dns-infoblox-webhook
version: %s (%s)
   _____ __________  _________   _____   
  /  _  \\______   \/   _____/  /  _  \  
 /  /_\  \|    |  _/\_____  \  /  /_\  \ 
/    |    \    |   \/        \/    |    \
\____|__  /______  /_______  /\____|__  /
        \/       \/        \/         \/ 
`

var (
	// Version - value can be overridden by ldflags
	Version = "local"
	Gitsha  = "?"
)

func main() {
	fmt.Printf(banner, Version, Gitsha)

	logging.Init()

	config := configuration.Init()
	provider, err := dnsprovider.Init(config)
	if err != nil {
		log.Fatalf("failed to initialize provider: %v", err)
	}

	srv := server.Init(config, webhook.New(provider))
	server.ShutdownGracefully(srv)
}
