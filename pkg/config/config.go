package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func NewConfig(p string) *viper.Viper {
	conf := viper.New()
	conf.AutomaticEnv()

	envConf := conf.GetString("APP_CONF")
	if envConf != "" {
		p = envConf
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		panic(errors.Errorf("config file not found: %s", p))
	}

	conf.SetConfigFile(p)

	if err := conf.ReadInConfig(); err != nil {
		panic(errors.Errorf("failed to read config file: %s", err))
	}

	// 绑定环境变量
	bindEnvVars(conf)

	return conf
}

// bindEnvVars 绑定环境变量
func bindEnvVars(conf *viper.Viper) {
	// 直接设置环境变量值
	if appID := os.Getenv("FEISHU_APP_ID"); appID != "" {
		conf.Set("feishu.app_id", appID)
	}
	if appSecret := os.Getenv("FEISHU_APP_SECRET"); appSecret != "" {
		conf.Set("feishu.app_secret", appSecret)
	}
	if redirectURL := os.Getenv("FEISHU_REDIRECT_URL"); redirectURL != "" {
		conf.Set("feishu.redirect_url", redirectURL)
	}
	if tenantKey := os.Getenv("FEISHU_TENANT_KEY"); tenantKey != "" {
		conf.Set("feishu.tenant_key", tenantKey)
	}
}
