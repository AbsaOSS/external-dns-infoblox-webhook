package dnsprovider

import (
	"testing"

	"github.com/AbsaOSS/external-dns-infoblox-webhook/cmd/webhook/init/configuration"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	cases := []struct {
		name          string
		config        configuration.Config
		env           map[string]string
		expectedError string
	}{
		{
			name:   "minimal config for infoblox provider",
			config: configuration.Config{},
			env: map[string]string{
				"INFOBLOX_WAPI_USER":     "user123",
				"INFOBLOX_WAPI_PASSWORD": "password",
				"INFOBLOX_VERSION":       "2.7.1",
			},
		},
		{
			name: "domain filter config for infoblox provider",
			config: configuration.Config{
				DomainFilter:   []string{"domain.com"},
				ExcludeDomains: []string{"sub.domain.com"},
			},
			env: map[string]string{
				"INFOBLOX_WAPI_USER":     "user123",
				"INFOBLOX_WAPI_PASSWORD": "password",
				"INFOBLOX_VERSION":       "2.7.1",
			},
		},
		{
			name:          "empty configuration",
			config:        configuration.Config{},
			expectedError: "expecting error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.env {
				t.Setenv(k, v)
			}

			dnsProvider, err := Init(tc.config)

			if tc.expectedError != "" {
				assert.Error(t, err, "configuration error, no mandatory Environment variables set")
				return
			}

			assert.NoErrorf(t, err, "error creating provider")
			assert.NotNil(t, dnsProvider)
		})
	}
}
