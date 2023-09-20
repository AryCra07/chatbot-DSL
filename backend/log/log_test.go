package log

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("Controller", "This is an information")

}

func TestWarning(t *testing.T) {
	Warning("Service", "This is a warning")
}

func TestError(t *testing.T) {
	Error("Dao", "This is an error")
}
