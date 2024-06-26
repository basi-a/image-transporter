package global

import (
	"image-transporter/utils"
	"os"

	"github.com/spf13/viper"
)

type BaseConfigStruct struct {
	Port   int    `json:"port" mapstructure:"port"`
	DbDir  string `json:"dbdir" mapstructure:"dbdir"`
	LogDir string `json:"logdir" mapstructure:"logdir"`
}

var BaseConfig BaseConfigStruct

func InitBaseConfig() {
	configFile, configDir, isNewFile := configFileStat()
	if isNewFile {
		setDefaultValue(configFile, configDir)
	}
	readConfigAndUnmarshal(configFile)
}

func configFileStat() (configFile, configDir string, isNewFile bool) {
	dir, err := os.UserConfigDir()
	if err != nil {
		utils.PrintfErr("get user config dir faild.", err)
		return "", "", false
	}
	configDir = dir + "/image-transporter"

	if !utils.DirExist(configDir) {
		if err := utils.MkdirAll(configDir); err != nil {
			utils.FatalfErr("mkdir all faild.", err)
		}
	}
	configFile = configDir + "/config.json"
	if !utils.FileStat(configFile) {
		_, err := os.Create(configFile)
		if err != nil {
			utils.PrintfErr("create config file faild.", err)
			return "", "", false
		}
	}
	return configFile, configDir, true
}

func setDefaultValue(configFile, configDir string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json")
	viper.SetDefault("port", 9988)
	viper.SetDefault("dbdir", configDir)
	logdir := configDir + "/log"
	if !utils.DirExist(logdir) {
		utils.MkdirAll(logdir)
	}
	viper.SetDefault("logdir", logdir)
	if err := viper.WriteConfigAs(configFile); err != nil {
		utils.PrintfErr("write and save config file faild.", err)
	}
}

func readConfigAndUnmarshal(configFile string)  {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		utils.FatalfErr("read config file faild.", err)
	}
	if err := viper.Unmarshal(&BaseConfig); err != nil {
		utils.FatalfErr("unmarshal config to struct faild.", err)
	}
}
