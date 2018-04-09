package nsshconfig

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

func GetPathConfig() string {
	usr := CurrentUser()
	return fmt.Sprintf("%s/.ssh/config", usr.HomeDir)
}

func CurrentUser() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return usr
}

func Contains(sub string, options []string, str ...string) bool {
	sub = strings.ToLower(sub)

	for _, s := range str {
		s = strings.ToLower(s)
		if strings.Contains(s, sub) {
			return true
		}
	}

	for _, o := range options {
		o = strings.ToLower(o)
		if strings.Contains(o, sub) {
			return true
		}
	}

	return false
}

func MatchString(rgxp string, compare string) bool {
	r, err := regexp.Compile(rgxp)
	if err != nil {
		fmt.Println("Invalid regexp")
		os.Exit(1)
	}
	return r.MatchString(FormatHostLine(compare))
}

func FormatHostLine(line string) string {
	return strings.TrimLeft(strings.ToLower(line), " ")
}

func Clear() {
	entries = []*Entry{}
	return
}

func (e *Entry) Decode() string {

	config := fmt.Sprintf("Host %s\n", e.Host)
	for key, value := range e.Options {
		config += fmt.Sprintf("	%s %s\n", key, value)
	}

	return config
}

func TempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join("/tmp", prefix+hex.EncodeToString(randBytes)+suffix)
}

func Copy(src, dst string) error {
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
