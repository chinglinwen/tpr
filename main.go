package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/bradfitz/tcpproxy"
)

type Config struct {
	Services map[string]Service
}

type Service struct {
	Type string
	Port string
	From string
	To   string
}

var (
	config  Config
	cfgfile = flag.String("c", "config.toml", "toml config file")
)

func main() {
	var p tcpproxy.Proxy
	for _, v := range config.Services {
		serviceAdd(p, v)
	}
	log.Fatal(p.Run())
}

func serviceAdd(p tcpproxy.Proxy, v Service) {
	if v.Type == "https" {
		if v.From == "" && v.To != "" {
			p.AddRoute(v.Port, tcpproxy.To(v.To)) // fallback
			return
		}
		p.AddSNIRoute(v.Port, v.From, tcpproxy.To(v.To))
		return
	}

	if v.From == "" && v.To != "" {
		p.AddRoute(v.Port, tcpproxy.To(v.To)) // fallback
		return
	}
	p.AddHTTPHostRoute(v.Port, v.From, tcpproxy.To(v.To))
}

func init() {
	flag.Parse()
	_, err := toml.DecodeFile(*cfgfile, &config)
	if err != nil {
		log.Fatal(err)
	}
}
