// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package httpServer

import "github.com/skeletonkey/lib-core-go/config"

var cfg *httpServer

func getConfig() *httpServer {
	config.LoadConfig("httpServer", &cfg)
	return cfg
}
