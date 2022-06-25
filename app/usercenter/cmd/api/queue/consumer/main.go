package main

import (
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	logx.DisableStat()
	consumer := dq.NewConsumer(dq.DqConf{
		Beanstalks: []dq.Beanstalk{
			{
				Endpoint: "192.168.10.146:11300",
				Tube:     "tube",
			},
			{
				Endpoint: "192.168.10.146:11301",
				Tube:     "tube",
			},
		},
		Redis: redis.RedisConf{
			Host: "localhost:6379",
			Type: redis.NodeType,
			Pass: "123456",
		},
	})
	consumer.Consume(func(body []byte) {
		// your consume logic
		logx.Info(string(body))
	})
}
