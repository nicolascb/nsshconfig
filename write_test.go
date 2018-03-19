package nsshconfig_test

import (
	"testing"

	"github.com/nicolascb/nsshconfig"
)

func TestWrite(t *testing.T) {
	err := nsshconfig.LoadConfig()

	if err != nil {
		t.Error(err.Error())
	}

	err = nsshconfig.WriteConfig()

	if err != nil {
		t.Error(err.Error())
	}
}

func CheckError(err error, t *testing.T) {
	if err != nil {
		t.Error(err.Error())
	}
}
