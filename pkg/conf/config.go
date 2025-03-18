package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"sync"
)

var (
	cfg  *ini.File
	once sync.Once
	env  string
)

func init() {
	env = os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
}

func Load() error {
	var err error
	once.Do(func() {
		configPath := filepath.Join("config", fmt.Sprintf("config.%s.ini", env))
		cfg, err = ini.Load(configPath)
	})
	return err
}

func GetEnv() string {
	return env
}

func GetString(section, key string) string {
	return cfg.Section(section).Key(key).String()
}

func GetInt(section, key string) int {
	return cfg.Section(section).Key(key).MustInt(0)
}

func GetBool(section, key string) bool {
	return cfg.Section(section).Key(key).MustBool(false)
}

func GetFloat64(section, key string) float64 {
	return cfg.Section(section).Key(key).MustFloat64(0)
}

func GetStringDefault(section, key, defaultValue string) string {
	return cfg.Section(section).Key(key).MustString(defaultValue)
}

func GetIntDefault(section, key string, defaultValue int) int {
	return cfg.Section(section).Key(key).MustInt(defaultValue)
}

func GetBoolDefault(section, key string, defaultValue bool) bool {
	return cfg.Section(section).Key(key).MustBool(defaultValue)
}

func GetFloat64Default(section, key string, defaultValue float64) float64 {
	return cfg.Section(section).Key(key).MustFloat64(defaultValue)
}