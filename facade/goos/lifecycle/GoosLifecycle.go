package lifecycle

import (
	"github.com/goosmesh/goos/core/support-db"
	"github.com/goosmesh/goos/core/support-plugin"
	"github.com/goosmesh/goos/core/support-plugin/manager"
	"github.com/goosmesh/goos/core/utils"
	"github.com/goosmesh/goos/plugin-config"
	"github.com/goosmesh/goos/plugin-goos-ui"
	"github.com/goosmesh/goos/plugin-security"
	"github.com/goosmesh/goos/plugin-service"
	"log"
	"net/http"
)



type GoosLifecycle struct {
	GoosPlugins *support_plugin.Plugins

	AppVersion string
}

func (g GoosLifecycle) checkVersion(plugin *support_plugin.Plugin) {
	if utils.Empty(&plugin.Version) {
		plugin.Version = g.AppVersion
	}
}

func (g GoosLifecycle) Init()  {
	log.Println("init")
	defer g.BeforeStart()
	// 初始化数据库
	support_db.Init()
	// 插件初始化
	g.GoosPlugins.Init()

}

func (g GoosLifecycle) BeforeStart()  {
	log.Println("before start")
	defer g.Starting()
	var plugin = support_plugin.ReflectCreatePlugin(plugin_security.CreatePlugin)
	g.checkVersion(plugin)
	g.GoosPlugins.Register(plugin.Name, plugin)
	plugin = support_plugin.ReflectCreatePlugin(plugin_config.CreatePlugin)
	g.checkVersion(plugin)
	g.GoosPlugins.Register(plugin.Name, plugin)
	plugin = support_plugin.ReflectCreatePlugin(plugin_goos_ui.CreatePlugin)
	g.checkVersion(plugin)
	g.GoosPlugins.Register(plugin.Name, plugin)
	plugin = support_plugin.ReflectCreatePlugin(plugin_service.CreatePlugin)
	g.checkVersion(plugin)
	g.GoosPlugins.Register(plugin.Name, plugin)
}

func (g GoosLifecycle) Starting()  {
	log.Print("starting")
	defer g.AfterStart()
	// 初始化插件
	for _, p := range g.GoosPlugins.PluginList() {
		p.Init()
	}
	// 创建router
	router := manager.CreateRouter()
	// 启动插件
	for _, p := range g.GoosPlugins.PluginList() {
		p.Start()
	}
	go http.ListenAndServe(":4323", router)
}

func (g GoosLifecycle) AfterStart()  {
	log.Print("after start")
}

func (g GoosLifecycle) BeforeDestroy()  {
	log.Print("before destroy")
	defer support_db.Close()
	defer g.Destroy()
}

func (g GoosLifecycle) Destroy()  {
	defer g.AfterDestroy()
	log.Print("destroy")

}

func (g GoosLifecycle) AfterDestroy()  {
	log.Print("after destroy")
}
