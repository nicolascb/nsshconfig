package nsshconfig

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// GetPathConfig return ~/.ssh/config
func GetPathConfig() string {
	return path.Join(CurrentUser().HomeDir, "/.ssh/config")
}

// CurrentUser get current user
func CurrentUser() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return usr
}

func matchStr(rgxp string, compare string) bool {
	r, err := regexp.Compile(rgxp)
	if err != nil {
		fmt.Println("Invalid regexp")
		os.Exit(1)
	}
	return r.MatchString(formatLine(compare))
}

func formatLine(line string) string {
	return strings.TrimLeft(strings.ToLower(line), " ")
}

// Clear entries
func Clear() {
	entries = []*Entry{}
	return
}

// Decode entry
func (e *Entry) Decode() string {

	config := fmt.Sprintf("Host %s\n", e.Host)
	for key, value := range e.Options {
		config += fmt.Sprintf("	%s %s\n", key, value)
	}

	return config
}

func tmpFile(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join("/tmp", prefix+hex.EncodeToString(randBytes)+suffix)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)

	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}
