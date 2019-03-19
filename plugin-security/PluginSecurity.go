package plugin_security

import (
	"github.com/goosmesh/goos/core/env"
	"github.com/goosmesh/goos/core/support-plugin"
	"github.com/goosmesh/goos/core/support-plugin/manager"
	"github.com/goosmesh/goos/plugin-security/controller"
	"github.com/goosmesh/goos/plugin-security/entity"
)

type PluginSecurity struct {
}
func (p *PluginSecurity) Init()  {

	entity.Init()

	manager.AddPipline("security", 100, env.GoosSecurityIgnoreUrls, HttpPipeSecurity)

	manager.RegisterRouter(&manager.Route{Name: "SecurityEcho",     Method:"GET",   Pattern:"/securityEcho", HandlerFunc: controller.SecurityEcho})
	manager.RegisterRouter(&manager.Route{Name: "SecurityLogin",     Method:"POST",   Pattern:"/api/security/login", HandlerFunc: controller.Login})
	manager.RegisterRouter(&manager.Route{Name: "SecurityCurrentUser",     Method:"GET",   Pattern:"/api/security/currentUser", HandlerFunc: controller.GetCurrentUser})
	manager.RegisterRouter(&manager.Route{Name: "SecurityChangeSelectNamespace",     Method:"GET",   Pattern:"/api/security/user/changeSelectNamespace", HandlerFunc: controller.ChangeSelectNamespace})


}
func (p *PluginSecurity) Start()  {
}

func CreatePlugin() *support_plugin.Plugin {
	return &support_plugin.Plugin{
		PluginMeta: support_plugin.PluginMeta{
			Name: "security",
			Version: "",
		},
		PluginBoot: new(PluginSecurity),
	}
}