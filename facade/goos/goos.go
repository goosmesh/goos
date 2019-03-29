package main

import (
	"fmt"
	"github.com/goosmesh/goos/core/env"
	"github.com/goosmesh/goos/facade/goos/lifecycle"
)

//var plugins = &support_plugin.Plugins{}
//
//func checkVersion(plugin *support_plugin.Plugin) {
//	if utils.Empty(&plugin.Version) {
//		plugin.Version = APP_VERSION
//	}
//}


func init() {
	env.InitEnv()
	env.GoosPort = "3321"

	lifecycle.GoosWorker()

	ch := make(chan string)
	select {
	case result := <-ch:
		fmt.Println(result)
	}
}

func main() {

}
