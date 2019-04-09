package nsshconfig

import (
	"bufio"
	"os"
	"strings"
)

// Entry config
type Entry struct {
	Host    string
	Options map[string]string
}

var (
	configPath *string
	entries    []*Entry
)

// LoadConfig read and parse ~/.ssh/config
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

		if matchStr("^host ", scanner.Text()) {
			if !fl {
				entries = append(entries, aux)
				aux = &Entry{}
				aux.Options = make(map[string]string)
			}
			fl = false
			host := scanner.Text()[4:len(scanner.Text())]
			aux.Host = strings.TrimSpace(host)
		} else {
			if !matchStr("^#", scanner.Text()) {
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

// TotalEntries get total entries
func TotalEntries() int {
	return len(entries)
}

// Hosts get hosts
func Hosts() []*Entry {
	return entries
}
