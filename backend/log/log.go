package log

import (
	"fmt"
	"strings"
	"time"
)

const (
	strWidth     int    = 10
	red          string = "\x1b[31m"
	green        string = "\x1b[32m"
	yellow       string = "\x1b[33m"
	restoreColor string = "\x1b[0m"
)

func logBasic(layer, content string) []string {
	cur := strings.Split(time.Now().Format("2006-01-02 15:04:05"), " ")
	for len(layer) < strWidth {
		layer = " " + layer
	}

	timestamp := "[" + cur[0] + " " + cur[1] + "] "
	content = " --- [" + layer + "] " + content + restoreColor
	return []string{timestamp, content}
}

func Info(layer, info string) {
	basic := logBasic(layer, info)
	infoContent := "\t" + basic[0] + green + "  INFO" + basic[1]
	fmt.Println(infoContent)
}

func Warning(layer, warning string) {
	basic := logBasic(layer, warning)
	warningContent := "\t" + basic[0] + yellow + "  WARN" + basic[1]
	fmt.Println(warningContent)
}

func Error(layer, error string) {
	basic := logBasic(layer, error)
	errorContent := "\t" + basic[0] + red + " ERROR" + basic[1]
	fmt.Println(errorContent)
}
