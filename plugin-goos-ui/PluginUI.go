package plugin_goos_ui

import (
	"github.com/goosmesh/goos/core/support-plugin"
	"github.com/goosmesh/goos/core/support-plugin/manager"
	"github.com/goosmesh/goos/plugin-goos-ui/controller"
)

type PluginUI struct {
}
func (p *PluginUI) Init()  {

	manager.RegisterRouter(&manager.Route{Name: "UIEcho",     Method:"GET",   Pattern:"/uiEcho", HandlerFunc: controller.UIEcho})
	manager.RegisterRouter(&manager.Route{Name: "UI_STATIC_SERVER",     Method:"GET",   Pattern:"/static/**", HandlerFunc: controller.HandleStaticResource})

}
func (p *PluginUI) Start()  {
}

func CreatePlugin() *support_plugin.Plugin {
	return &support_plugin.Plugin{
		PluginMeta: support_plugin.PluginMeta{
			Name: "ui",
			Version: "",
		},
		PluginBoot: new(PluginUI),
	}
}