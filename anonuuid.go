package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	DEFAULT_REGEX_UUID = "[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}"
)

// AnonUUID is the main structure, it contains the cache map and helpers
type AnonUUID struct {
	Cache map[string]string
}

// Sanitize takes a string as input and return sanitized string
func (a *AnonUUID) Sanitize(input string) string {
	r := regexp.MustCompile(DEFAULT_REGEX_UUID)

	return r.ReplaceAllStringFunc(input, func(m string) string {
		parts := r.FindStringSubmatch(m)
		return a.FakeUUID(parts[0])
	})
}

// FakeUUID takes a realUUID and return its corresponding fakeUUID
func (a *AnonUUID) FakeUUID(realUUID string) string {
	if _, ok := a.Cache[realUUID]; !ok {
		nextId := len(a.Cache)
		fakeUUID := strings.Repeat(fmt.Sprintf("%x", nextId), 32)[:32]
		fakeUUID = fakeUUID[:8] + "-" + fakeUUID[8:12] + "-" + fakeUUID[12:16] + "-" + fakeUUID[16:20] + "-" + fakeUUID[20:32]
		a.Cache[realUUID] = fakeUUID
	}
	return a.Cache[realUUID]
}

// New returns a prepared AnonUUID structure
func New() *AnonUUID {
	return &AnonUUID{
		Cache: make(map[string]string),
	}
}

// main is the entrypoint
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	anonuuid := New()

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(anonuuid.Sanitize(line))
	}
}
