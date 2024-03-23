package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/structs"
	"github.com/spf13/viper"
)

var defaults = Config{
	Log: Log{
		OutputFile: "",
		Level:      "",
	},
	PresentationMethod: "htmx",
}

func GetDefaultConfig() Config {
	return defaults
}

type Handler struct {
	config HandlerConfig
}

func NewHandler(configs ...Configuration) (*Handler, error) {
	handler := &Handler{}
	for _, config := range configs {
		err := config(handler)
		if err != nil {
			return nil, fmt.Errorf("config: %w", err)
		}
	}
	return handler, nil
}

type Configuration func(h *Handler) error

type HandlerConfig struct {
	Filepath string
}

func WithConfig(config HandlerConfig) Configuration {
	return func(h *Handler) error {
		h.config = config
		return nil
	}
}

func (h *Handler) LoadConfig() (Config, error) {
	if h.config.Filepath == "" {
		config, err := loadConfig()
		if err != nil {
			return config, fmt.Errorf("loadConfig: %w", err)
		}
		return config, err
	}
	config, err := loadConfigFile[Config](h.config.Filepath, defaults)
	if err != nil {
		return config, fmt.Errorf("loadConfigFile: %w", err)
	}
	return config, err
}

// Loads standard/domain config.
func loadConfig() (Config, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetEnvPrefix("WW")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	setDefaults(v, defaults)
	v.AutomaticEnv()
	config := new(Config)
	*config = defaults

	err := v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return *config, fmt.Errorf("v.ReadInConfig: %w", err)
	}

	if err = v.Unmarshal(config); err != nil {
		return *config, fmt.Errorf("v.Unmarshal: %w", err)
	}
	return *config, nil
}

// Loads config file with specifc name (exclude extension).
func loadConfigFile[T any](filename string, defaults T) (T, error) {
	v := viper.New()
	v.SetConfigFile(filename)
	setDefaults(v, defaults)
	config := new(T)
	*config = defaults

	err := v.ReadInConfig()
	if err != nil {
		return *config, fmt.Errorf("v.ReadInConfig: %w", err)
	}

	if err = v.Unmarshal(config); err != nil {
		return *config, fmt.Errorf("v.Unmarshal: %w", err)
	}
	return *config, nil
}

func setDefaults(v *viper.Viper, defaults any) {
	defaultsMap := structs.Map(defaults)
	for key, value := range defaultsMap {
		v.SetDefault(key, value)
	}
}
