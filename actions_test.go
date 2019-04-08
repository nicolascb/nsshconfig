package nsshconfig_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nicolascb/nsshconfig"
)

func TestDeleteHost(t *testing.T) {
	err := nsshconfig.LoadConfig()
	totalBeforeDelete := nsshconfig.TotalEntries()
	err = nsshconfig.Delete("*")

	if err != nil {
		t.Errorf(err.Error())
	}

	if totalBeforeDelete == nsshconfig.TotalEntries() {
		t.Errorf("Not delete, total is equal before")
	}

	err = nsshconfig.Delete("*")
	if err.Error() != "Host not found" {
		t.Errorf("Delete duplicate accept...")
	}

	nsshconfig.WriteConfig()
}

func TestDeleteHostNotFound(t *testing.T) {
	err := nsshconfig.LoadConfig()
	err = nsshconfig.Delete("*abc")

	if err == nil {
		t.Errorf("Host not found delete accept")
	}
	if err.Error() != "Host not found" {
		t.Errorf("Host not found invalid err")
	}
}

// Edit host
func TestEditHost(t *testing.T) {
	err := nsshconfig.LoadConfig()
	if err != nil {
		t.Errorf(err.Error())
	}
	host := nsshconfig.Hosts()[1]

	host.Options["hostname"] = "hostname2.edited.com.br"
	err = host.Save()
	if err != nil {
		t.Errorf(err.Error())
	}
	nsshconfig.WriteConfig()
}

func TestNewHost(t *testing.T) {
	err := nsshconfig.LoadConfig()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	options := make(map[string]string)
	options["port"] = "5133"
	options["hostname"] = "gremio.net"
	err = nsshconfig.New("novo_pabx", options)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	nsshconfig.WriteConfig()

}

func TestEditGetHost(t *testing.T) {
	// nsshconfig.SetConfigPath("/home/nicolas/.ssh/config")
	err := nsshconfig.LoadConfig()

	if err != nil {
		t.Error(err)
		return
	}

	host, err := nsshconfig.GetEntryByHost("*")
	if err != nil {
		t.Error(err)
		return
	}

	host.Options["port"] = "51222"
	err = host.Save()
	if err != nil {
		t.Error(err)
		return
	}

	err = nsshconfig.WriteConfig()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestInvalidHost(t *testing.T) {
	// nsshconfig.SetConfigPath("/home/nicolas/.ssh/config")
	err := nsshconfig.LoadConfig()

	if err != nil {
		t.Error(err)
		return
	}

	_, err = nsshconfig.GetEntryByHost("*as")
	if err == nil {
		t.Error("Invalid host not handle error")
		return
	}

	if err.Error() != "Host not found" {
		t.Error(errors.New("Invalid err to get invalid host"))
		return
	}

}

func TestGetHost(t *testing.T) {
	// nsshconfig.SetConfigPath("/home/nicolas/.ssh/config")
	err := nsshconfig.LoadConfig()

	if err != nil {
		t.Error(err)
	}

	general, err := nsshconfig.GetEntryByHost("*")
	if err != nil {
		t.Error(err)
		return
	}

	if general.Host != "*" {
		t.Error(errors.New("Invalid general host"))
		return
	}

}
func TestExistHost(t *testing.T) {
	// nsshconfig.SetConfigPath("/home/nicolas/.ssh/config")
	// $ cat ~/.ssh/config
	// Host teste
	//       hostname teste.com
	//       user root

	err := nsshconfig.LoadConfig()

	if err != nil {
		t.Error(err)
	}

	// Check if exist host
	if nsshconfig.ExistHost("abcdfeg") {
		// Host already exist, print error and exit
		t.Error(fmt.Errorf("incorrect existhost"))
		return
	}

	// Check if exist host
	if !nsshconfig.ExistHost("teste") {
		// Host already exist, print error and exit
		t.Error(fmt.Errorf("incorrect existhost - 2"))
		return
	}

}
