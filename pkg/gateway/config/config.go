package config

import (
	"errors"

	"github.com/kelseyhightower/envconfig"

	"github.com/skygeario/skygear-server/pkg/core/redis"
	"github.com/skygeario/skygear-server/pkg/gateway/model"
)

// Configuration is gateway startup configuration
type Configuration struct {
	Standalone                        bool
	StandaloneTenantConfigurationFile string               `envconfig:"STANDALONE_TENANT_CONFIG_FILE" default:"standalone-tenant-config.yaml"`
	StandaloneHost                    StandaloneHostConfig `envconfig:"STANDALONE_HOST"`
	Host                              string               `envconfig:"SERVER_HOST" default:"localhost:3001"`
	ConnectionStr                     string               `envconfig:"DATABASE_URL"`
	Auth                              GearURLConfig        `envconfig:"AUTH"`
	Asset                             GearURLConfig        `envconfig:"ASSET"`
	Redis                             redis.Configuration  `envconfig:"REDIS"`
	UseInsecureCookie                 bool                 `envconfig:"INSECURE_COOKIE"`
	AuthProxyHeaders                  string               `envconfig:"INSECURE_COOKIE"`
}

// ReadFromEnv reads from environment variable and update the configuration.
func (c *Configuration) ReadFromEnv() error {
	err := envconfig.Process("", c)
	if err != nil {
		return err
	}
	return nil
}

type GearURLConfig struct {
	Live    string `envconfig:"LIVE_URL"`
	Nightly string `envconfig:"NIGHTLY_URL"`
}

// GetGearURL provide router map
func (c *Configuration) GetGearURL(gear model.Gear, version model.GearVersion) (string, error) {
	var g GearURLConfig
	switch gear {
	case model.AuthGear:
		g = c.Auth
	case model.AssetGear:
		g = c.Asset
	default:
		return "", errors.New("invalid gear")
	}

	switch version {
	case model.LiveVersion:
		return g.Live, nil
	case model.NightlyVersion:
		return g.Nightly, nil
	default:
		return "", errors.New("gear is suspended")
	}
}

type StandaloneHostConfig struct {
	AuthHost  string `envconfig:"AUTH_SERVER_HOST"`
	AssetHost string `envconfig:"ASSET_SERVER_HOST"`
	AppHost   string `envconfig:"APP_SERVER_HOST"`
}
