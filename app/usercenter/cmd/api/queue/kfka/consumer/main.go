package main

import (
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	logx.DisableStat()
	var c kq.KqConf
	conf.MustLoad("/home/zzx/go-proj/GOPATH/src/go-zero-demo-single/app/usercenter/cmd/api/queue/kfka/consumer/config.yaml", &c)

	// WithHandle: 具体的处理msg的logic
	// 这也是开发者需要根据自己的业务定制化
	q := kq.MustNewQueue(c, kq.WithHandle(func(k, v string) error {
		fmt.Printf("=> %s\n", v)
		return nil
	}))
	defer q.Stop()
	q.Start()
}
