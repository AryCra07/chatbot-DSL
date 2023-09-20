package log

import (
	"backend/consts"
	"fmt"
	"os"
	"strings"
	"time"
)

const strWidth int = 10

func Info(layer string, info string) {
	cur := strings.Split(time.Now().Format("2006-01-02 15:04:05"), " ")

	for len(layer) < strWidth {
		layer = layer + " "
	}

	logContent := "[" + cur[0] + " " + cur[1] + "] ---" + layer + "\t>" + info

	if os.Args[1] == consts.ModeTest {
		fmt.Println("\t" + logContent + " ayu")
	} else if os.Args[1] == consts.ModeRelease {
		if strings.Contains(layer, consts.Dao) != true {
			fmt.Println("\t" + logContent)
		}
	}
}
