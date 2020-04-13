package logger

import (
	"testing"
)

func Test_logger(t *testing.T) {
	Dump()
	// SetFilepath("/tmp/logger.log");
	SetSize(1024)
	SetLevel(LEVEL_ALL)
	Dump()

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

func BenchmarkTraceln(b *testing.B) {
	SetFilepath("/tmp/logger.log");
	for i := 0; i < b.N; i++ {
		Traceln("hello", " world")
	}
}

func BenchmarkTracef(b *testing.B) {
	SetFilepath("/tmp/logger.log");
	for i := 0; i < b.N; i++ {
		Tracef("%s - %s", "hello", "world")
	}
}
