package out_writer

import (
	"fmt"
	"io"
	"logger/logger"
	"os"
	"time"
)

type ConsoleSink struct {
	writer io.Writer
}

func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{writer: os.Stdout}
}

func (c *ConsoleSink) Write(level logger.LogLevel, msg string) error {
	_, err := fmt.Fprintf(c.writer, "[%s] [%s] %s\n", time.Now().Format(time.RFC3339), level, msg)
	return err
}

func (c *ConsoleSink) Close() error { return nil }
