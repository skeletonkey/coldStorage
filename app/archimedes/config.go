// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package archimedes

import "github.com/skeletonkey/lib-core-go/config"

var cfg *archimedes

func getConfig() *archimedes {
	config.LoadConfig("archimedes", &cfg)
	return cfg
}
