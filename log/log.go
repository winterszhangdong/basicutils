package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error ]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[32m[info  ]\033[0m ", log.LstdFlags|log.Lshortfile)
	debugLog = log.New(os.Stdout, "\033[34m[debug ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog, debugLog}
	mu       sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
	Debug  = debugLog.Println
	Debugf = debugLog.Printf
)

const (
	DebugLevel = iota
	InfoLevel
	ErrorLevel
	Disabled
)

func SetPath(path string, level int) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 666)

	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	if ErrorLevel >= level {
		errorLog.SetOutput(io.MultiWriter(file, os.Stderr))
	}

	if InfoLevel >= level {
		infoLog.SetOutput(io.MultiWriter(file, os.Stdout))
	}

	if DebugLevel >= level {
		debugLog.SetOutput(io.MultiWriter(file, os.Stdout))
	}
}

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}

	if DebugLevel < level {
		debugLog.SetOutput(ioutil.Discard)
	}
}
