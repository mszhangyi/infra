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

/*import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/mattn/go-colorable"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/go-utils"
	"github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var formatter *prefixed.TextFormatter
var lfh *utils.LineNumLogrusHook

func init() {
	// 定义日志格式
	formatter = &prefixed.TextFormatter{}
	//设置高亮显示的色彩样式
	formatter.ForceColors = true
	formatter.DisableColors = false
	formatter.ForceFormatting = true
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})
	//开启完整时间戳输出和时间戳格式
	formatter.FullTimestamp = true
	//设置时间格式
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	//设置日志formatter
	log.SetFormatter(formatter)
	log.SetOutput(colorable.NewColorableStdout())
	//日志级别，通过环境变量来设置
	// 后期可以变更到配置中来设置

	if os.Getenv("log.debug") == "true" {
		log.SetLevel(log.DebugLevel)
	}
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

//将滚动日志writer共享给iris glog output
var log_writer io.Writer

//初始化log配置，配置logrus日志文件滚动生成和
func InitLog() {
	//设置日志输出级别
	level, err := log.ParseLevel(props.LogLevel)
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
	if props.LogEnableLineLog {
		lfh.EnableFileNameLog = true
		lfh.EnableFuncNameLog = true
	} else {
		lfh.EnableFileNameLog = false
		lfh.EnableFuncNameLog = false
	}

	//配置日志输出目录
	logPath := props.LogDir //+ "/logs"
	/*logFilePath, _ := filepath.Abs(logPath)
	log.Infof("log dir: %s", logFilePath)*/
	os.MkdirAll(logPath, os.ModePerm)
	baseLogPath := path.Join(logPath, props.Name)
	//设置滚动日志输出writer
	writer, err := rotatelogs.New(
		strings.TrimSuffix(baseLogPath, ".log")+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour * time.Duration(props.LogMaxAge)),             // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(props.LogRotationTime)), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", err)
	}
	//设置日志文件输出的日志格式
	formatter := &log.TextFormatter{}
	formatter.CallerPrettyfier = func(frame *runtime.Frame) (function string, file string) {
		function = frame.Function
		dir, filename := path.Split(frame.File)
		f := path.Base(dir)
		return function, fmt.Sprintf("%s/%s:%d", f, filename, frame.Line)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, formatter)

	log.AddHook(lfHook)
	//
	log_writer = writer
}*/