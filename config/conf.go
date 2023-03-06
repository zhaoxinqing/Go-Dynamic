package config

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/spf13/viper"
)

var Config *EnvConf

type EnvConf struct {
	MySQL    MySQL    `yaml:"MySQL" mapstructure:"MySQL"`
	Redis    Redis    `yaml:"Redis" mapstructure:"Redis"`
	Web      Web      `yaml:"Web" mapstructure:"Web"`
	Contract Contract `yaml:"Contract" mapstructure:"Contract"`
}

type Contract struct {
	ZoneHash            []byte `yaml:"ZoneHash" mapstructure:"ZoneHash"`
	OffererConduitKey   []byte `yaml:"OffererConduitKey" mapstructure:"OffererConduitKey"`
	FulfillerConduitKey []byte `yaml:"FulfillerConduitKey" mapstructure:"FulfillerConduitKey"`
	Signature           []byte `yaml:"Signature" mapstructure:"Signature"`
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
	DBQueue  int    `yaml:"db_queue"  mapstructure:"db_queue"`
}

// Web ...
type Web struct {
	Domain string `yaml:"domain" mapstructure:"domain"`
	AppEnv string `yaml:"app_env" mapstructure:"app_env"`
}

//go:embed env.yaml
var f []byte

// LoadConf ...
func LoadConf() {
	r := bytes.NewReader(f)
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}
}
