package plugin_config

import (
	"github.com/goosmesh/goos/core/support-plugin"
	"github.com/goosmesh/goos/core/support-plugin/manager"
	"github.com/goosmesh/goos/plugin-config/entity"
	"github.com/goosmesh/goos/plugin-config/handler"
)

type PluginConfig struct {
}
func (p *PluginConfig) Init()  {

	entity.Init()

	manager.RegisterRouter(&manager.Route{Name: "ConfigEcho",     Method:"GET",   Pattern:"/configEcho", HandlerFunc: handler.ConfigEcho})

	manager.RegisterRouter(&manager.Route{Name: "ConfigCreate",     Method:"POST",   Pattern:"/api/config/config", HandlerFunc: handler.Config})
	manager.RegisterRouter(&manager.Route{Name: "ConfigGetList",     Method:"GET",   Pattern:"/api/config/configList", HandlerFunc: handler.GetConfigList})
	manager.RegisterRouter(&manager.Route{Name: "ConfigGet",     Method:"GET",   Pattern:"/api/config/config", HandlerFunc: handler.GetConfig})
	manager.RegisterRouter(&manager.Route{Name: "ConfigDelete",     Method:"DELETE",   Pattern:"/api/config/config", HandlerFunc: handler.DeleteConfig})

	manager.RegisterRouter(&manager.Route{Name: "NamespaceCreate",     Method:"POST",   Pattern:"/api/namespace/namespace", HandlerFunc: handler.Namespace})
	manager.RegisterRouter(&manager.Route{Name: "NamespaceGetList",     Method:"GET",   Pattern:"/api/namespace/namespaceList", HandlerFunc: handler.GetNamespaceList})
	manager.RegisterRouter(&manager.Route{Name: "NamespaceDelete",     Method:"DELETE",   Pattern:"/api/namespace/namespace", HandlerFunc: handler.DeleteNamespace})




	// wait update method to get
	manager.RegisterRouter(&manager.Route{Name: "ClientConfigRefreshGET",     Method:"POST",   Pattern:"/api/pub/config/get/listen", HandlerFunc: handler.RsaGetConfig})
	manager.RegisterRouter(&manager.Route{Name: "ClientConfigGET",     Method:"GET",   Pattern:"/api/pub/config/get", HandlerFunc: handler.GetConfigClient})

}
func (p *PluginConfig) Start()  {
}

func CreatePlugin() *support_plugin.Plugin {
	return &support_plugin.Plugin{
		PluginMeta: support_plugin.PluginMeta{
			Name: "config",
			Version: "",
		},
		PluginBoot: new(PluginConfig),
	}
}