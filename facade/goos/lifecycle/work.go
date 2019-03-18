package lifecycle

import (
	"github.com/jsen-joker/goos/core/env"
	"github.com/jsen-joker/goos/core/support-plugin"
)

func GoosWorker()  {
	// 初始化环境变量
	env.InitEnv()

	life := GoosLifecycle{
		GoosPlugins: &support_plugin.Plugins{},
		AppVersion: env.GoosVersion,
	}
	life.Init()

	defer life.BeforeDestroy()
}