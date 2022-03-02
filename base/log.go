package base

import (
	"fmt"
	"github.com/mszhangyi/infra/utils"
	"github.com/mszhangyi/work/udpLog"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var formatter *prefixed.TextFormatter
var lfh *utils.LineNumLogrusHook

func init() {
	// 定义日志格式
	formatter = &prefixed.TextFormatter{}
	//开启完整时间戳输出和时间戳格式
	formatter.FullTimestamp = true
	//设置时间格式
	formatter.TimestampFormat = "2006-01-02T15:04:05-0700"
	//设置日志formatter
	log.SetFormatter(formatter)

	//开启调用函数、文件、代码行信息的输出
	log.SetReportCaller(true)

	//设置函数、文件、代码行信息的输出的hook
	SetLineNumLogrusHook()
}

func SetLineNumLogrusHook() {
	lfh = utils.NewLineNumLogrusHook()
	lfh.EnableFileNameLog = true
	lfh.EnableFuncNameLog = true
	log.AddHook(lfh)
}

//初始化log配置，配置logrus日志文件滚动生成和
func InitLog() {
	//设置日志输出级别
	level, err := log.ParseLevel(props.LogLevel)
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
	//配置日志输出目录
	logger, err := udpLog.NewUDPWriter(props.LogDir, props.Name)
	if err != nil {
		fmt.Print(err)
	}
	log.SetOutput(logger)
}
