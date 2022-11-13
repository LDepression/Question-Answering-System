package settings

import (
	"flag"
	"strings"
	"sync"

	"wenba/internal/global"
	"wenba/internal/pkg/setting"
)

var once sync.Once

var (
	configPaths string // 配置文件路径
	configName  string // 配置文件名
	configType  string // 配置文件类型
)

//StringVar定义了一个有指定名字，默认值，和用法说明的string标签。参数p指向一个存储标签解析值的string变量。
//Args 指定参数名 应用的时候 在命令行输入 -Args xxx
//defaultValue 如果没有指定Args的值，那么Args的内容默认是"defaultValue"
//Usage 用法说明字符串
//flag.StringVar(&args, "Args", "defaultValue", "Usage:xxx")
//解析上面定义的标签

func setupFlag() {
	// 命令行参数绑定
	flag.StringVar(&configName, "name", "app", "配置文件名")
	flag.StringVar(&configType, "type", "yml", "配置文件类型")
	flag.StringVar(&configPaths, "path", global.RootDir+"\\config\\app", "指定要使用的配置文件路径")
	flag.Parse()
}

// 读取配置文件
func init() {
	//只会被执行一次
	once.Do(func() {
		setupFlag()
		// 在调用其他组件的Init时，这个init会首先执行并且把配置文件绑定到全局的结构体上
		newSetting, err := setting.NewSetting(configName, configType, strings.Split(configPaths, ",")...) // 引入配置文件路径
		if err != nil {
			panic(err)
		}
		if err := newSetting.BindAll(&global.Settings); err != nil {
			panic(err)
		}
	})
}
