package nsshconfig

import (
	"errors"
	"strings"
)

func (e *Entry) Save() error {
	options := make(map[string]string)

	for key, value := range e.Options {
		options[strings.ToLower(key)] = value
	}

	// Invalid Host
	if e.Host == "" {
		return errors.New("Invalid host")
	}

	// Hostname key not found
	if e.Host != "*" {
		if _, ok := options["hostname"]; !ok {
			return errors.New("Hostname not found")
		}
	}

	e.Options = options

	for idx, x := range entries {
		if x.Host == e.Host {
			entries[idx] = e
			return nil
		}
	}

	entries = append(entries, e)
	By(Prop("Host", true)).Sort(entries)

	return nil
}

func Delete(host string) error {
	for idx, x := range entries {
		if strings.ToLower(x.Host) == strings.ToLower(host) {
			entries = append(entries[:idx], entries[idx+1:]...)
			return nil
		}
	}

	return errors.New("Host not found")

}

func New(host string, options map[string]string) error {
	n := &Entry{}
	n.Options = make(map[string]string)

	n.Host = strings.TrimSpace(host)
	n.Options = options

	return n.Save()
}

func GetEntryByHost(host string) (*Entry, error) {
	for _, h := range entries {
		if strings.ToLower(h.Host) == strings.ToLower(host) {
			return h, nil
		}
	}

	return nil, errors.New("Host not found")
}

func SetConfigPath(file string) {
	configPath = &file
}
