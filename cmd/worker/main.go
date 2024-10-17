package main

import (
	"worker-go/pkg/config"
	"worker-go/pkg/kafka"
)

func main() {
	config.LoadEnv()

	kafka.ConsumeKafka()

	select {}
}
