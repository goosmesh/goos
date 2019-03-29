package lifecycle

import (
	"github.com/goosmesh/goos/core/env"
	"github.com/goosmesh/goos/core/support-plugin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func GoosWorker()  {
	// 初始化环境变量
	log.Println("goos starting")

	life := GoosLifecycle{
		GoosPlugins: &support_plugin.Plugins{},
		AppVersion: env.GoosVersion,
	}
	go signalHandler(life)
	life.Init()

	//defer life.BeforeDestroy()
}

func signalHandler(life GoosLifecycle) {
	ch := make(chan os.Signal)
	//signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			// this ensures a subsequent INT/TERM will trigger standard go behaviour of
			// terminating. 执行标准的go终止行为，程序就结束了
			life.BeforeDestroy()
			signal.Stop(ch)
			os.Exit(0)
			//a.term(wg)
			return
		//case syscall.SIGUSR2: // 这里开始执行优雅重启
			//err := a.preStartProcess()
			//// 这个函数在源代码中没有具体实现功能，只是预留了一个钩子函数，用户可以注册自己的函数，可以在重启之前做些自定义的事情。一般情况下也没有什么可以做的，除非有些特殊的服务环境或是状态保存之类的，至少目前，我们的server还没有遇到
			//if err != nil {
			//	a.errors <- err
			//}
			//// we only return here if there's an error, otherwise the new process
			//// will send us a TERM when it's ready to trigger the actual shutdown.
			//if _, err := a.net.StartProcess(); err != nil { // 这里开始正式所谓的优雅重启
			//	a.errors <- err
			//}
		}
	}
}