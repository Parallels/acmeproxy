package acmeproxy

import (
	"net"
	"net/http"
	"strconv"

	"github.com/go-acme/lego/v3/challenge"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
)

type Server struct {
	HttpServer     *http.Server
	Provider       challenge.Provider
	HtpasswdFile   string
	AllowedIPs     []string
	AllowedDomains []string
	AccesslogFile  string
	CheckDNS       bool
	CheckResolver  *net.Resolver
}

func NewServer(config *Config) (*Server, error) {
	return &Server{
		HttpServer:     config.HttpServer,
		Provider:       config.Provider,
		HtpasswdFile:   config.HtpasswdFile,
		AllowedIPs:     config.AllowedIPs,
		AllowedDomains: config.AllowedDomains,
		AccesslogFile:  config.AccesslogFile,
		CheckDNS:       config.CheckDNS,
		CheckResolver:  config.CheckResolver,
	}, nil
}

func RunServer(ctx *cli.Context, config *Config) {
	if config.HttpServer.TLSConfig != nil {
		log.WithFields(log.Fields{
			"endpoint": "https://" + ctx.GlobalString("interface") + ":" + strconv.Itoa(ctx.GlobalInt("port")),
			"addr":     config.HttpServer.Addr,
		}).Info("Starting acmeproxy")
		log.Fatal(config.HttpServer.ListenAndServeTLS("", ""))
	} else {
		log.WithFields(log.Fields{
			"endpoint": "http://" + ctx.GlobalString("interface") + ":" + strconv.Itoa(ctx.GlobalInt("port")),
			"addr":     config.HttpServer.Addr,
		}).Info("Starting acmeproxy")
		log.Fatal(config.HttpServer.ListenAndServe())
	}
}
