// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package woody

import "github.com/skeletonkey/lib-core-go/config"

var cfg *woody

func getConfig() *woody {
	config.LoadConfig("woody", &cfg)
	return cfg
}
