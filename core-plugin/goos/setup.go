package redis

import (
	"github.com/goosmesh/goos/core/env"
	"github.com/goosmesh/goos/facade/goos/lifecycle"
	"strings"

	"github.com/coredns/coredns/plugin"
	"github.com/mholt/caddy"
)


func init() {
	caddy.RegisterPlugin(Name, caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	err := goosParse(c)
	if err != nil {
		return plugin.Error(Name, err)
	}

	go lifecycle.GoosWorker()

	return nil
}

func goosParse(c *caddy.Controller) error {

	for c.Next() {
		if c.NextBlock() {
			for {
				switch c.Val() {
				case "goos.home":
					if !c.NextArg() {
						return c.ArgErr()
					}
					env.GoosHome = c.Val()
				case "goos.security.ignore.urls":
					if !c.NextArg() {
						return c.ArgErr()
					}
					env.GoosSecurityIgnoreUrls = strings.Split(c.Val(), ",")
				case "goos.database":
					if !c.NextArg() {
						return c.ArgErr()
					}
					env.GoosDatabase = c.Val()
				case "goos.version":
					if !c.NextArg() {
						return c.ArgErr()
					}
					env.GoosVersion = c.Val()
				case "goos.port":
					if !c.NextArg() {
						return c.ArgErr()
					}
					env.GoosPort = c.Val()
				default:
					if c.Val() != "}" {
						return c.Errf("unknown property '%s'", c.Val())
					}
				}

				if !c.Next() {
					break
				}
			}

		}
		return nil
	}
	return nil
}