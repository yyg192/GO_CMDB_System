package conf

//配置对象加载
import (
	"github.com/BurntSushi/toml"
)

var (
	global_config *Config
)

// C 全局配置对象
func C() *Config {
	if global_config == nil {
		panic("Load Config first")
	}

	return global_config
}

// LoadConfigFromToml 从toml中添加配置文件, 并初始化全局对象
func LoadConfigFromToml(filePath string) error {
	cfg := newConfig()
	if _, err := toml.Decode(filePath, cfg); err != nil {
		return err
	}
	global_config = cfg
	return nil
}
