package main

//go:generate go run directives_generate.go

import (
	"github.com/coredns/coredns/coremain"
	"github.com/goosmesh/goos/facade/goos/lifecycle"

	// Plug in CoreDNS
	_ "github.com/coredns/coredns/core/plugin"
)

func main() {
	go lifecycle.GoosWorker()
	coremain.Run()
}
