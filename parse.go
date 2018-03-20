package nsshconfig

import (
	"bufio"
	"os"
	"strings"
)

type Entry struct {
	Host    string
	Options map[string]string
}

var configPath *string

var entries []*Entry

func LoadConfig() error {
	var fl = true
	var err error
	var f *os.File
	aux := &Entry{}

	// Clear entries
	Clear()

	aux.Options = make(map[string]string)

	// Open file
	if configPath == nil {
		homeConfig := GetPathConfig()
		configPath = &homeConfig
	}
	f, err = os.Open(*configPath)

	if err != nil {
		return err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	// Loop file
	for scanner.Scan() {

		if MatchString("^host ", scanner.Text()) {
			if !fl {
				entries = append(entries, aux)
				aux = &Entry{}
				aux.Options = make(map[string]string)
			}
			fl = false
			host := strings.Replace(FormatHostLine(scanner.Text()), "host", "", -1)
			aux.Host = strings.TrimSpace(host)
		} else {
			if !MatchString("^#", scanner.Text()) {
				var lineOption []string
				key := ""
				val := ""

				line := strings.Split(strings.TrimSpace(scanner.Text()), "#")[0]

				if strings.Contains(line, "=") {
					lineOption = strings.SplitN(strings.TrimSpace(line), "=", 2)
				} else {
					lineOption = strings.SplitN(strings.TrimSpace(line), " ", 2)
				}

				if len(lineOption) >= 2 {
					key = strings.ToLower(lineOption[0])
					val = strings.TrimSpace(lineOption[1])
					aux.Options[key] = val
				}

				lineOption = nil
			}
		}
	}

	// Last host
	if aux.Host != "" {
		entries = append(entries, aux)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Sort by Host
	By(Prop("Host", true)).Sort(entries)

	return nil
}

func TotalEntries() int {
	return len(entries)
}

func Hosts() []*Entry {
	return entries
}
