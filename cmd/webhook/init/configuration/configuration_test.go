package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Setenv("SERVER_READ_TIMEOUT", "1s")
	t.Setenv("SERVER_WRITE_TIMEOUT", "1s")

	cfg := Init()

	assert.Equal(t, "0.0.0.0", cfg.ServerHost)
	assert.Equal(t, 8888, cfg.ServerPort)
	assert.Equal(t, []string(nil), cfg.DomainFilter)
	assert.Equal(t, []string(nil), cfg.ExcludeDomains)
	assert.Equal(t, "", cfg.RegexDomainFilter)
	assert.Equal(t, "", cfg.RegexDomainExclusion)

	t.Setenv("SERVER_HOST", "testhost")
	t.Setenv("SERVER_PORT", "9999")
	t.Setenv("DOMAIN_FILTER", "test.com,test2.com")
	t.Setenv("EXCLUDE_DOMAIN_FILTER", "exclude.com,exclude2.com")
	t.Setenv("REGEXP_DOMAIN_FILTER", ".*test.*")
	t.Setenv("REGEXP_DOMAIN_FILTER_EXCLUSION", ".*exclude.*")
	t.Setenv("REGEXP_DOMAIN_FILTER", ".*test.*")
	t.Setenv("REGEXP_DOMAIN_FILTER_EXCLUSION", ".*exclude.*")

	cfg = Init()
	assert.Equal(t, "testhost", cfg.ServerHost)
	assert.Equal(t, 9999, cfg.ServerPort)
	assert.Equal(t, []string{"test.com", "test2.com"}, cfg.DomainFilter)
	assert.Equal(t, []string{"exclude.com", "exclude2.com"}, cfg.ExcludeDomains)
	assert.Equal(t, ".*test.*", cfg.RegexDomainFilter)
	assert.Equal(t, ".*exclude.*", cfg.RegexDomainExclusion)
}
