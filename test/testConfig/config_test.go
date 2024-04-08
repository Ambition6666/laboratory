package testconfig

import (
	"laboratory/config"
	"testing"
)

func TestConfig(t *testing.T) {
	conf := config.GetMySQLDB()
	t.Logf(conf.Host)
}