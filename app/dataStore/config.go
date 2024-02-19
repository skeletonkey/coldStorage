// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package dataStore

import "github.com/skeletonkey/lib-core-go/config"

var cfg *dataStore

func getConfig() *dataStore {
	config.LoadConfig("dataStore", &cfg)
	return cfg
}
