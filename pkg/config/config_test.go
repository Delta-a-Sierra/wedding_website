package config_test

import (
	"testing"

	"github.com/Delta-a-Sierra/wedding_website/pkg/config"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	getHandler    func() (*config.Handler, error)
	expectedError error
	expected      config.Config
}

func TestHandler_LoadConfig(t *testing.T) {
	tests := map[string]testStruct{
		"loading config with no file returns defaults": {
			getHandler: func() (*config.Handler, error) {
				return config.NewHandler()
			},
			expected:      config.GetDefaultConfig(),
			expectedError: nil,
		},
		"loading log config from file returns correct values": {
			getHandler: func() (*config.Handler, error) {
				return config.NewHandler(config.WithConfig(config.HandlerConfig{
					Filepath: "./test_files/config.yml",
				}))
			},
			expected: func() config.Config {
				config := config.GetDefaultConfig()
				config.Log.Level = "debug"
				config.Log.OutputFile = "./log.log"
				return config
			}(),
			expectedError: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			asserts := assert.New(t)
			handler, err := tc.getHandler()
			if err != nil {
				asserts.Equal(tc.expectedError, err)
			}
			config, err := handler.LoadConfig()
			if err != nil {
				asserts.Equal(tc.expectedError, err)
			}
			asserts.Equal(tc.expected, config)
		})
	}
}
