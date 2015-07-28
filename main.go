package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type CacheMap map[string]string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	r := regexp.MustCompile("[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}")

	cache := CacheMap{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(r.ReplaceAllStringFunc(line, func(m string) string {
			parts := r.FindStringSubmatch(m)
			if _, ok := cache[parts[0]]; !ok {
				nextId := len(cache) + 20
				fakeUUID := strings.Repeat(fmt.Sprintf("%x", nextId), 32)[:32]
				fakeUUID = fakeUUID[:8] + "-" + fakeUUID[8:12] + "-" + fakeUUID[12:16] + "-" + fakeUUID[16:20] + "-" + fakeUUID[20:32]
				cache[parts[0]] = fakeUUID
			}
			return cache[parts[0]]
		}))
	}
}
