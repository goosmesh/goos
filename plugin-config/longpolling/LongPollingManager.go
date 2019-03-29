package longpolling

import (
	"container/list"
	"context"
	"github.com/goosmesh/goos/plugin-config/longpolling/utils"
)

// 管理长轮询

type ClientListener struct {
	ctx context.Context
	ch chan<- string
	clientMd5Map map[string] string
}

var clients = list.New()

// 增加一个lp
// 1、首先检查是否有数据变化，有变化直接返回
// 2、没有变化，检查是否包含直接返回头，包含则直接返回
// 3、加入到监听池，在配置文件变化时回调
// 4、或超时
func AddLongPolling(ctx context.Context, ch chan<- string, clientMd5Map map[string] string, noHangup bool) (exe func(), clientClose func()) {
	exe = func() {
		changed := utils.FilterChangedConfig(clientMd5Map)
		if len(changed) != 0 {
			// 有数据，直接返回
			ch <- utils.ChangedConfigToResponse(changed)
		} else if noHangup {
			// 立即返回
			ch <- ""
		} else {
			// 加入到监听列表，暂时未实现
			ele := clients.PushBack(ClientListener{
				ctx: ctx,
				ch: ch,
				clientMd5Map: clientMd5Map,
			})

			clientClose = func() {
				// 从监听列表移除
				clients.Remove(ele)
			}
		}
		select {
		case <-ctx.Done():
			close(ch)
		}
	}
	return
}