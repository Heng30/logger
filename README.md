# Go-logger

## 功能：
- 日志输出到文件或标准输出
- 输出到文件达到设置大小重写
- 支持日志级别
- 线程安全

## 例子：
```go
func main() {
    Traceln("hello", " world")
    Tracef("%s - %s", "hello", "world")
    Debugln("hello", " world")
    Debugf("%s - %s", "hello", "world")
    Infoln("hello", " world")
    Infof("%s - %s", "hello", "world")
    Warnln("hello", " world")
    Warnf("%s - %s", "hello", "world")
    Errorln("hello", " world")
    Errorf("%s - %s", "hello", "world")
    Fatalln("hello", " world")
    Fatalf("%s - %s", "hello", "world")
}
```

输出：
```text
[T][2020/04/13/17:12:16][logger_test.go/Test_logger:14]: hello  world
[T][2020/04/13/17:12:16][logger_test.go/Test_logger:15]: hello - world
[D][2020/04/13/17:12:16][logger_test.go/Test_logger:16]: hello  world
[D][2020/04/13/17:12:16][logger_test.go/Test_logger:17]: hello - world
[I][2020/04/13/17:12:16][logger_test.go/Test_logger:18]: hello  world
[I][2020/04/13/17:12:16][logger_test.go/Test_logger:19]: hello - world
[W][2020/04/13/17:12:16][logger_test.go/Test_logger:20]: hello  world
[W][2020/04/13/17:12:16][logger_test.go/Test_logger:21]: hello - world
[E][2020/04/13/17:12:16][logger_test.go/Test_logger:22]: hello  world
[E][2020/04/13/17:12:16][logger_test.go/Test_logger:23]: hello - world
[F][2020/04/13/17:12:16][logger_test.go/Test_logger:24]: hello  world
[F][2020/04/13/17:12:16][logger_test.go/Test_logger:25]: hello - world
```
