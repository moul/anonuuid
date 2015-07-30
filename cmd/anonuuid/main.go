package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/moul/anonuuid"
)

// main is the entrypoint
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	anonuuid := anonuuid.New()

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(anonuuid.Sanitize(line))
	}
}
