package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

)

// 日志文件
type LogConfig struct {
	Access string `json:"access"`
	Error  string `json:"error"`
}

type HttpConfig struct {
	Listen string `json:"listen"`
	Secret string `json:"secret"`
}

type RpcConfig struct {
	Listen string `json:"listen"`
}

type MysqlConfig struct {
	Addr    string 			  `json:"addr"`
	Idle    int               `json:"idle"`
	Max     int               `json:"max"`
	ShowSQL bool              `json:"sql"`
}

type TimeoutConfig struct {
	Conn  int64 `json:"conn"`
	Read  int64 `json:"read"`
	Write int64 `json:"write"`
}

type RedisConfig struct {
	Addr    string         `json:"addr"`
	Idle    int            `json:"idle"`
	Max     int            `json:"max"`
	Timeout *TimeoutConfig `json:"timeout"`
}

type ConfigMsql struct {
	Addr      string        `json:"addr"`
}

type Limit struct {
	TypeNub   int64 `json:"typenub"`	// 限制用户可以创建的主分类数量
	Mode	  int64 `json:"mode"`		// 限制用户可以创建的模型的数量
}

type UpdateRate struct {
	Status    	bool 	`json:"status"`   // 是否开启限制
	RateTime	int64 	`json:"rate_time"`		// 频率，默认是秒
}

type GlobalConfig struct {
	Debug   	bool                `json:"debug"` //
	Init   		bool                `json:"init"` //是否初始化，
	Version     float64 			`json:"version"`
	UpdateRate  *UpdateRate 		`json:"update_rate"`  // 更新频率，按分钟
	Request     string 				`json:"request"` // 是否使用https
	Salt    	string              `json:"salt"`  // 验证key
	Log     	*LogConfig          `json:"log"`
	LimitNub	*Limit				`json:"limit"`
	Http    	*HttpConfig         `json:"http"`
	Rpc     	*RpcConfig          `json:"rpc"`
	Mysql   	*MysqlConfig        `json:"mysql"`
	ConfMysql   *ConfigMsql         `json:"confmysql"`
	Redis   	*RedisConfig        `json:"redis"`
	Backend 	map[string][]string `json:"backend"`
}


var (
	File string
	G    *GlobalConfig
)

func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("无配置文件，请使用 -c 配置文件 启动")
	}

	if !fileExists(cfg) {
		return fmt.Errorf("配置文件 %s 不存在", cfg)
	}

	File = cfg

	configContent, err := ioutil.ReadFile(cfg)
	if err != nil {
		return fmt.Errorf("read configuration file %s fail %s", cfg, err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal(configContent, &c)
	if err != nil {
		return fmt.Errorf("parse configuration file %s fail %s", cfg, err.Error())
	}

	G = &c

	log.Println("load configuration file", cfg, "successfully")
	return nil
}

func fileExists(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}
