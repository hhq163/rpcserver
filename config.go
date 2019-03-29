package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var Cfg Config

type Config struct {
	LogLevel string `yaml:"LogLevel"`
	Address  string `yaml:"Address"`

	MySqlURL     string `yaml:"MySqlURL"`
	MySqlThreads int    `yaml:"MySqlThreads"`

	MongoURL     string `yaml:"MongoURL"`
	MongoThreads int    `yaml:"MongoThreads"`

	DebugMode bool `yaml:"DebugMode"`
}

func init() {
	versiontype := flag.String("t", "dev", "发布环境，如test，可选项dev,test,docker,release")
	flag.Parse()

	//ReadFile函数会读取文件的全部内容，并将结果以[]
	filename := GetExecpath() + "/conf-" + *versiontype + ".yaml"
	log.Print(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("读取配置文件错误", err)
	}

	//读取的数据为json格式，需要进行解码
	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		log.Fatal("解析配置文件错误", err)
	}
}

// 获取当前程序运行目录
func GetExecpath() string {
	execpath, _ := os.Executable() // 获得程序路径
	path := filepath.Dir(execpath)
	return strings.Replace(path, "\\", "/", -1)
}
