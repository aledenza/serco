package serco

import (
	"encoding/json"
	"os"
	"path"

	"github.com/creasty/defaults"
	"github.com/gurkankaymak/hocon"
)

const defaultEnv = "default"

func NewConfig[T any](folder string, searchKey string) (configStruct T) {
	configPath := path.Join(folder)
	env, ok := os.LookupEnv("env")
	if !ok {
		env = defaultEnv
	}
	var config *hocon.Config
	baseConfig, err := hocon.ParseResource(path.Join(configPath, defaultEnv+".conf"))
	if err != nil {
		panic(err)
	}
	if env != defaultEnv {
		envConfig, err := hocon.ParseResource(path.Join(configPath, env+".conf"))
		if err != nil {
			panic(err)
		}
		config = envConfig.WithFallback(baseConfig)
	} else {
		config = baseConfig
	}
	var cfg hocon.Value
	if searchKey != "" {
		cfg = config.Get(searchKey)
	} else {
		cfg = config.GetRoot()
	}
	if err := defaults.Set(&configStruct); err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &configStruct); err != nil {
		panic(err)
	}
	return configStruct
}
