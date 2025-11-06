package out_writer

import (
	"fmt"
	"logger/logger"
	"os"
	"sync"
	"time"
)

// FileSink writes logs to a file
type FileSink struct {
	mu   sync.Mutex
	file *os.File
}

func NewFileSink(path string) (*FileSink, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileSink{file: f}, nil
}

func (f *FileSink) Write(level logger.LogLevel, msg string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	_, err := fmt.Fprintf(f.file, "[%s] [%s] %s\n", time.Now().Format(time.RFC3339), level, msg)
	return err
}

func (f *FileSink) Close() error {
	return f.file.Close()
}
