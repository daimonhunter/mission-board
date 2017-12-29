package env

import (
	"github.com/larspensjo/config"
	"flag"
	"log"
)

var(
	//https://studygolang.com/articles/686
	//支持命令行输入格式为-configfile=name, 默认为config.ini
	//配置文件一般获取到都是类型
	configFile = flag.String("configfile","./config/config.ini","General configuration file")
	CONFIG = make(map[string]string)
)

//Read configuration file and returns the specific section as a map
func GetSection(s string) map[string]string {
	cfg,err := config.ReadDefault(*configFile)   //读取配置文件，并返回其Config

	if err != nil {
		log.Fatalf("Fail to find %v,%v",*configFile,err)
	}

	if cfg.HasSection(s) {   //判断配置文件中是否有section（一级标签）
		options,err := cfg.SectionOptions(s)    //获取一级标签的所有子标签options（只有标签没有值）
		if err == nil {
			for _,v := range options{
				optionValue,err := cfg.String(s,v)  //根据一级标签section和option获取对应的值
				if err == nil {
					CONFIG[v] =optionValue
				}
			}
		}
	}
	return CONFIG
}

