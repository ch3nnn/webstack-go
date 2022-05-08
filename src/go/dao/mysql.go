package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// DRIVER 指定驱动
const DRIVER = "mysql"

var SqlSession *gorm.DB

//配置参数映射结构体
type conf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"post"`
}

//获取配置参数数据
func (c *conf) getConf() *conf {
	//读取resources/application.yaml文件
	yamlFile, err := ioutil.ReadFile("resources/application.yaml")
	//若出现错误，打印错误提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

// InitMySql 初始化连接数据库，生成可操作基本增删改查结构的变量
func InitMySql() (err error) {
	var c conf
	//连接数据库
	SqlSession, err = gorm.Open(DRIVER, c.ConnectUrl())
	if err != nil {
		panic(err)
	}
	//验证数据库连接是否成功，若成功，则无异常
	return SqlSession.DB().Ping()
}

// ConnectUrl 数据库连接字符串
func (c conf) ConnectUrl() string {
	//获取yaml配置参数
	conf := c.getConf()
	//将yaml配置参数拼接成连接数据库的url
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
}

// Close 关闭数据库连接
func Close() {
	err := SqlSession.Close()
	if err != nil {
		return
	}
}
