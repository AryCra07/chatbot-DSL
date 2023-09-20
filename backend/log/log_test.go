package log

import (
	"backend/consts"
	"os"
	"testing"
)

func TestLogInModeTest(t *testing.T) {
	os.Args = []string{"", consts.ModeTest}

	Info("Model", "In test I Love")

	Info("Dao", "In test I Love")

}

func TestLogInModeRelease(t *testing.T) {
	os.Args = []string{"", consts.ModeRelease}

	Info("Service", "In release can be seen")

	Info("Dao", "can't be seen")
}
