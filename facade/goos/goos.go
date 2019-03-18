package main

import "github.com/jsen-joker/goos/facade/goos/lifecycle"

//var plugins = &support_plugin.Plugins{}
//
//func checkVersion(plugin *support_plugin.Plugin) {
//	if utils.Empty(&plugin.Version) {
//		plugin.Version = APP_VERSION
//	}
//}


func init() {
	lifecycle.GoosWorker()
}

func main() {

}
