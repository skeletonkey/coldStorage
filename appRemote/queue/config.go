// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package queue

import "github.com/skeletonkey/lib-core-go/config"

var cfg *queue

func getConfig() *queue {
	config.LoadConfig("queue", &cfg)
	return cfg
}
