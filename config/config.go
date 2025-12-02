package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	App		AppConfig
	MySQL	MySQLConfig
	Redis	RedisConfig
	DeepSeek DeepSeekConfig `mapstructure:"deepseek"` // 显式声明
}

type AppConfig struct {
	Name string
	Mode string
	Port int
}

type MySQLConfig struct {
	Host		string
	Port		int
	User 		string
	Password	string
	Database	string
	MaxOpen		int
	MaxIdle 	int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type DeepSeekConfig struct {
	ApiKey  string `mapstructure:"api_key"`
	BaseURL string `mapstructure:"base_url"`
	Model   string `mapstructure:"model"`
}

func InitConfig() {
	v := viper.New()
	// 基础配置
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")
	// 读取 config.yaml
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// 支持环境变量
	v.AutomaticEnv()
	// MySQL 环境变量覆盖
	v.BindEnv("mysql.password", "MYSQL_PASSWORD")
	v.BindEnv("mysql.user", "MYSQL_USER")
	v.BindEnv("mysql.host", "MYSQL_HOST")
	v.BindEnv("mysql.port", "MYSQL_PORT")
	// App 端口环境变量覆盖
	v.BindEnv("app.port", "APP_PORT")
	// Redis 环境变量覆盖
	v.BindEnv("redis.host", "REDIS_HOST")
	v.BindEnv("redis.password", "REDIS_PASSWORD")

	// 解析到全局配置结构体
	Cfg = &Config{}
	err = v.Unmarshal(Cfg)
	if err != nil {
		panic(fmt.Errorf("unmarshal config error: %w", err))
	}
	fmt.Printf("Deepseek Config: %+v\n", Cfg.DeepSeek)
	fmt.Println("Config Loaded Sucessfully")
	fmt.Printf("Current Port: %d\n", Cfg.App.Port)
}