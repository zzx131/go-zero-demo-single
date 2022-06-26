package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
)

type message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"message"`
}

func main() {
	pusher := kq.NewPusher([]string{
		"192.168.10.146:20123",
		"192.168.10.146:20124",
		"192.168.10.146:20125",
	}, "kq")

	ticker := time.NewTicker(time.Millisecond)
	for round := 0; round < 3; round++ {
		<-ticker.C

		count := rand.Intn(100)
		m := message{
			Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
			Value:   fmt.Sprintf("%d,%d", round, count),
			Payload: fmt.Sprintf("%d,%d", round, count),
		}
		body, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
		if err := pusher.Push(string(body)); err != nil {
			log.Fatal(err)
		}
	}

	cmdline.EnterToContinue()
}
