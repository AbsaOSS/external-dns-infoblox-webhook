// package main
package main

/*
Copyright 2024 The external-dns-infoblox-webhook Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/

import (
	"fmt"

	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/configuration"
	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/dnsprovider"
	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/logging"
	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/server"
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

	srv := server.NewServer()

	srv.StartHealth(config)
	srv.Start(config, provider)
}
