package log

import (
	"log"
	"os"
)

func PrintNormal(format string, v ...any) {
	log.Printf(format, v...)
}

func PrintGreen(format string, v ...any) {
	log.New(os.Stdout, "\033[32m", log.LstdFlags).Printf(format, v...)
}
