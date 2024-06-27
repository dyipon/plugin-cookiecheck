package plugin_cookiecheck

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Enabled bool
}

func CreateConfig() *Config {
	return &Config{}
}

type Plugin struct {
	next http.Handler
	name string
}

func New(_ context.Context, next http.Handler, cfg *Config, name string) (http.Handler, error) {

	if next == nil {
		return nil, fmt.Errorf("no next handler provided")
	}
	if cfg == nil {
		return nil, fmt.Errorf("no config provided")
	}

	return &Plugin{
		next: next,
		name: name,
	}, nil
}

func (p *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if len(req.Cookies()) == 0 {
		rw.WriteHeader(http.StatusForbidden)
		fmt.Fprint(rw, "terminating")
		log.Println("cookieless request dropped")
	}

	return
}
