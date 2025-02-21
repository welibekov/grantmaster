package config

import (
	"os"
	"strings"
)

func Load() map[string]string {
	config := make(map[string]string)
	for _, env := range os.Environ() {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) == 2 && strings.HasPrefix(kv[0], "GM_") {
			config[kv[0]] = kv[1]
		}
	}
	return config
}
