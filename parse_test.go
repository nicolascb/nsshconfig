package nsshconfig_test

import (
	"fmt"
	"testing"

	"github.com/nicolascb/nsshconfig"
)

func TestListHosts(t *testing.T) {
	// nsshconfig.SetConfigPath("/home/nicolas/.ssh/config")
	err := nsshconfig.LoadConfig()

	if err != nil {
		t.Error(err)
	}

	if nsshconfig.TotalEntries() != 6 {
		t.Errorf("Total entries invalid: %d\n", nsshconfig.TotalEntries())
	}

	hosts := nsshconfig.Hosts()
	for _, h := range hosts {
		fmt.Println(h.Host)
		fmt.Println(h.Options)
	}

}
