// 调试日志库:
//	1.支持trace=1, debug=2, info=4, warn=8, error=16, fatal=32 6种日志级别
//	2.日志级别采用 或 方式，即 LEVEL_TRACE|LEVEL_DEBUG 则开启trace与debug级别

package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

// 日志级别
const (
	LEVEL_TRACE = 1
	LEVEL_DEBUG = 2
	LEVEL_INFO  = 4
	LEVEL_WARN  = 8
	LEVEL_ERROR = 16
	LEVEL_FATAL = 32
	LEVEL_ALL   = 63
)

type logger struct {
	filepath string /* 文件路径 */
	level    int    /* 日志级别 */
	size     int    /* 日志重写大小 */
	mutex    sync.Mutex
}

var g_logger logger = logger{
	filepath: "",
	level:    LEVEL_ALL,
	size:     1024 * 1024, /* 1M */
}


// 设置输出日志文件路径
func SetFilepath(filepath string) {
	g_logger.mutex.Lock()
	defer func() {
		g_logger.mutex.Unlock()
	}()

    g_logger.filepath = filepath
}


// 设置日志级别
func SetLevel(level int) {
	g_logger.mutex.Lock()
	defer func() {
		g_logger.mutex.Unlock()
	}()

    g_logger.level = level;
}

// 设置日志重写大小
func SetSize(size int) {
	g_logger.mutex.Lock()
	defer func() {
		g_logger.mutex.Unlock()
	}()

    g_logger.size = size;
}


// 写入日志文件
func writeLog(prefix, format string, v ...interface{}) (err error) {
	var file *os.File
	var lger *log.Logger
	var flag int
	var perm os.FileMode = 0666
	filename, line, funcname := "???", 0, "???"

	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
		funcname = filepath.Ext(funcname)
		funcname = strings.TrimPrefix(funcname, ".")
		filename = filepath.Base(filename)
	}

	g_logger.mutex.Lock()
	defer func() {
		if len(g_logger.filepath) != 0 {
			file.Close()
		}
		g_logger.mutex.Unlock()
	}()

	if len(g_logger.filepath) == 0 {
		file = os.Stdout
	} else {
		fileInfo, ok := os.Lstat(g_logger.filepath)
		if ok != nil {
			fmt.Println(ok)
			flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
		} else {
			if fileInfo.Size() < int64(g_logger.size) {
				flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
			} else {
				flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
			}
		}

		if file, err = os.OpenFile(g_logger.filepath, flag, perm); err != nil {
			fmt.Println(err)
			return err
		}
	}

	t := time.Now()
	prefix = prefix + fmt.Sprintf("[%d/%02d/%02d/%02d:%02d:%02d]",
		t.Year(), t.Month(), t.Day(), t.Hour(),
		t.Minute(), t.Second()) + "[" +
		filename + "/" + funcname + ":" +
		fmt.Sprint(line) + "]: "

	if lger = log.New(file, prefix, 0); lger == nil {
		return fmt.Errorf("log.New %s error", g_logger.filepath)
	}

	if len(format) > 0 {
		lger.Printf(format+"\n", v...)
	} else {
		lger.Println(v...)
	}
	return
}

//	追踪级别
func Traceln(v ...interface{}) {
	if g_logger.level&LEVEL_TRACE == 0 {
		return
	}

	if err := writeLog("[T]", "", v...); err != nil {
		fmt.Println(err)
	}
}

//	追踪级别
func Tracef(format string, v ...interface{}) {
	if g_logger.level&LEVEL_TRACE == 0 {
		return
	}
	if err := writeLog("[T]", format, v...); err != nil {
		fmt.Println(err)
	}
}

// 调试级别
func Debugln(v ...interface{}) {
	if g_logger.level&LEVEL_DEBUG == 0 {
		return
	}

	if err := writeLog("[D]", "", v...); err != nil {
		fmt.Println(err)
	}
}

// 调试级别
func Debugf(format string, v ...interface{}) {
	if g_logger.level&LEVEL_DEBUG == 0 {
		return
	}
	if err := writeLog("[D]", format, v...); err != nil {
		fmt.Println(err)
	}
}

// 信息级别
func Infoln(v ...interface{}) {
	if g_logger.level&LEVEL_INFO == 0 {
		return
	}

	if err := writeLog("[I]", "", v...); err != nil {
		fmt.Println(err)
	}
}

// 信息级别
func Infof(format string, v ...interface{}) {
	if g_logger.level&LEVEL_INFO == 0 {
		return
	}
	if err := writeLog("[I]", format, v...); err != nil {
		fmt.Println(err)
	}
}

// 警告级别
func Warnln(v ...interface{}) {
	if g_logger.level&LEVEL_WARN == 0 {
		return
	}

	if err := writeLog("[W]", "", v...); err != nil {
		fmt.Println(err)
	}
}

// 警告级别
func Warnf(format string, v ...interface{}) {
	if g_logger.level&LEVEL_WARN == 0 {
		return
	}
	if err := writeLog("[W]", format, v...); err != nil {
		fmt.Println(err)
	}
}

// 错误级别
func Errorln(v ...interface{}) {
	if g_logger.level&LEVEL_ERROR == 0 {
		return
	}

	if err := writeLog("[E]", "", v...); err != nil {
		fmt.Println(err)
	}
}

// 错误级别
func Errorf(format string, v ...interface{}) {
	if g_logger.level&LEVEL_ERROR == 0 {
		return
	}
	if err := writeLog("[E]", format, v...); err != nil {
		fmt.Println(err)
	}
}

// 致命级别
func Fatalln(v ...interface{}) {
	if g_logger.level&LEVEL_FATAL == 0 {
		return
	}

	if err := writeLog("[F]", "", v...); err != nil {
		fmt.Println(err)
	}
}

// 致命级别
func Fatalf(format string, v ...interface{}) {
	if g_logger.level&LEVEL_FATAL == 0 {
		return
	}

	if err := writeLog("[F]", format, v...); err != nil {
		fmt.Println(err)
	}
}

// 打印日志信息
func Dump() {
	s := reflect.ValueOf(&g_logger).Elem()
	t := s.Type()

	Traceln("=====================")
	for i := 0; i < s.NumField(); i++ {
		Traceln(t.Field(i).Name, ":", s.Field(i))
	}
	Traceln("=====================")
}
