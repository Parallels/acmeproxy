package acmeproxy

import (
	"net"
	"net/http"
	"github.com/go-acme/lego/v3/challenge"
)

type Config struct {
	HttpServer       *http.Server
	Provider         challenge.Provider
	ProviderName     string
	HtpasswdFile     string
	AllowedIPs       []string
	AllowedDomains   []string
	AccesslogFile    string
	CheckDNS         bool
	CheckResolver    *net.Resolver
}

func NewDefaultConfig() *Config {
	return &Config{}
}
