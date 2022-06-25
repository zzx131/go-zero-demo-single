package main

import (
	"fmt"
	"github.com/zeromicro/go-queue/dq"
	"strconv"
	"time"
)

func main() {
	producer := dq.NewProducer([]dq.Beanstalk{
		{
			Endpoint: "192.168.10.146:11300",
			Tube:     "tube",
		},
		{
			Endpoint: "192.168.10.146:11301",
			Tube:     "tube",
		},
	})

	for i := 1000; i < 1005; i++ {
		_, err := producer.Delay([]byte(strconv.Itoa(i)), time.Minute*2)
		if err != nil {
			fmt.Println(err)
		}
	}
}
