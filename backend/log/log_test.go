package log

import (
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	str := strings.Join([]string{"This is an information", "This is an information"}, " <-| ")
	Info("Controller", str)

}

func TestWarning(t *testing.T) {
	Warning("Service", "This is a warning")
}

func TestError(t *testing.T) {
	Error("Dao", "This is an error")
}
