package conf

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Urls UrlCfg

type UrlCfg struct {
	URLs []string `yaml:"urls"`
}

func InitUrls() (UrlCfg, error) {
	// 读取YAML文件
	data, err := ioutil.ReadFile("urls.yaml")
	if err != nil {
		log.Printf("Error parsing YAML file: %v", err)
		return UrlCfg{}, err
	}

	// 解析YAML数据
	err = yaml.Unmarshal(data, &Urls)
	if err != nil {
		log.Printf("Error parsing YAML file: %v", err)
		return UrlCfg{}, err
	}

	// 输出解析结果
	fmt.Println("Parsed URLs:")
	for _, url := range Urls.URLs {
		fmt.Println(url)
	}

	return Urls, nil
}

func SelectHandler() (map[string][]string, error) {
	data, err := ioutil.ReadFile("select.yaml")
	if err != nil {
		log.Printf("Error reading YAML file: %v", err)
		return nil, err
	}
	// 定义一个动态 map 来存储解析后的数据
	var vipMap map[string][]string

	// 解析 YAML 数据到 map
	err = yaml.Unmarshal(data, &vipMap)
	if err != nil {
		log.Printf("Error parsing YAML file: %v", err)
		return nil, err
	}
	return vipMap, nil
}

type appcfg struct {
	DbHost   string `yaml:"db_host"`   //数据库地址
	DbPort   int    `yaml:"db_port"`   //数据库端口
	DbUser   string `yaml:"db_user"`   //数据库账号
	DbPasswd string `yaml:"db_passwd"` //数据库密码
	DbName   string `yaml:"db_name"`   //数据库名称
	DbParams string `yaml:"db_params"` //数据库参数
}

const _dsn = "%s:%s@tcp(%s:%d)/%s?%s"

func (app *appcfg) Dsn() string {
	return fmt.Sprintf(_dsn, app.DbUser, app.DbPasswd, app.DbHost, app.DbPort, app.DbName, app.DbParams)
}

var (
	App       *appcfg              //运行配置实体
	defConfig = "./conf/conf.yaml" //配置文件路径，方便测试
)

func Init() {
	var err error
	App, err = initCfg()
	if err != nil {
		log.Printf("config init error : %v", err)
	}

	log.Printf("conf init")
}
func initCfg() (*appcfg, error) {
	app := &appcfg{}

	// 读取 YAML 格式的配置文件内容
	yamlContent, err := ioutil.ReadFile(defConfig)
	if err != nil {
		log.Printf("failed to read config.yml:%v", err)
	}

	// 解析 YAML 格式的配置文件内容到 app 结构体中

	err = yaml.Unmarshal(yamlContent, &app)
	if err != nil {
		log.Printf("failed to unmarshal config:%v", err)
	}
	return app, nil
}
