package config

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var config *EnvConf

type EnvConf struct {
	MySQL  MySQL  `yaml:"MySQL" mapstructure:"MySQL"`
	Redis  Redis  `yaml:"Redis" mapstructure:"Redis"`
	Web    Web    `yaml:"Web" mapstructure:"Web"`
	Logger Logger `yaml:"Logger" mapstructure:"Logger"`
	Jwt    Jwt    `yaml:"JWT" mapstructure:"JWT"`
}

// MySQL ...
type MySQL struct {
	Source   string `yaml:"source" mapstructure:"source"`
	Replica1 string `yaml:"replica1" mapstructure:"replica1"`
	Replica2 string `yaml:"replica2" mapstructure:"replica2"`
}

// Redis ...
type Redis struct {
	Addr     string `yaml:"addr" mapstructure:"addr"`
	Password string `yaml:"password" mapstructure:"password"`
	DBs      []int  `yaml:"dbs" mapstructure:"dbs"`
}

// Web ...
type Web struct {
	Domain          string `yaml:"domain" mapstructure:"domain"`
	InitialPassword string `yaml:"initial_password" mapstructure:"initial_password"`
}

// Logger ...
type Logger struct {
	Level    string `yaml:"level" mapstructure:"level"`
	Driver   string `yaml:"driver" mapstructure:"driver"`
	FilePath string `yaml:"file_path" mapstructure:"file_path"`
	Format   string `yaml:"format" mapstructure:"format"`
	MaxSize  int    `yaml:"max_size" mapstructure:"max_size"`
}

type Jwt struct {
	Secret    string `yaml:"secret" mapstructure:"secret"`
	ExpiresAt int64  `yaml:"expires_at" mapstructure:"expires_at"`
	Issuer    string `yaml:"issuer" mapstructure:"issuer"`
}

//go:embed env.example.yaml
var f []byte

// LoadConf ...
func LoadConf() {

	r := bytes.NewReader(f)
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		zap.L().Info(err.Error())
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}

	zap.L().Info("Load Config success")
}
func GetConfig() *EnvConf {
	return config
}

func GetMySQLEnv() *MySQL {
	return &config.MySQL
}

func GetRedisEnv() *Redis {
	return &config.Redis
}
func GetWebEnv() *Web {
	return &config.Web
}

func GetLogger() *Logger {
	return &config.Logger
}

// GetJwtEnv 获取 JWT 配置
func GetJwtEnv() *Jwt {
	return &config.Jwt
}

func EnvInfo() *EnvConf {
	return config
}
