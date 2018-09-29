package conf

import (
	"flag"
	"github.com/cihub/seelog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SmsConf struct {
	// 短信信息
	AccessKeyID     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	TemplateParam   string `yaml:"TemplateParam"`
	SignName        string `yaml:"SignName"`
}

type Conf struct {
	Host          string  `yaml:"Host"`
	Port          string  `yaml:"Port"`
	MysqlUsername string  `yaml:"MysqlUsernam"`
	MysqlPassword string  `yaml:"MysqlPassword"`
	MysqlPort     string  `yaml:"MysqlPort"`
	Redis         string  `yaml:"Redis"`
	RedisPort     string  `yaml:"RedisPort"`
	ALiDaYu       SmsConf `yaml:"ALiDaYu"`
}

var Configuration *Conf

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Conf
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	Configuration = &config
	return err
}

func InitConf() {
	configFilePath := flag.String("C", "conf/conf.yaml", "config file path")
	logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")
	flag.Parse()

	logger, err := seelog.LoggerFromConfigAsFile(*logConfigPath)
	if err != nil {
		seelog.Critical("err parsing seelog config file", err)
		return
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	if err := LoadConfiguration(*configFilePath); err != nil {
		return
	}
}
