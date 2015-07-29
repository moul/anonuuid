package anonuuid

import (
	"fmt"
	"testing"
)

var (
	EXAMPLE_INPUT = `
VOLUMES_0_SERVER_ID=15573749-c89d-41dd-a655-16e79bed52e0
VOLUMES_0_SERVER_NAME=hello
VOLUMES_0_ID=c245c3cb-3336-4567-ada1-70cb1fe4eefe
VOLUMES_0_SIZE=50000000000
ORGANIZATION=fe1e54e8-d69d-4f7c-a9f1-42069e03da31
TEST=15573749-c89d-41dd-a655-16e79bed52e0
`
)

func ExampleNew() {
	anonuuid := New()
	fmt.Println(anonuuid)
	// Output:
	// &{map[]}
}

func ExampleSanitize() {
	anonuuid := New()
	fmt.Println(anonuuid.Sanitize(EXAMPLE_INPUT))
	// Output:
	// VOLUMES_0_SERVER_ID=00000000-0000-0000-0000-000000000000
	// VOLUMES_0_SERVER_NAME=hello
	// VOLUMES_0_ID=11111111-1111-1111-1111-111111111111
	// VOLUMES_0_SIZE=50000000000
	// ORGANIZATION=22222222-2222-2222-2222-222222222222
	// TEST=00000000-0000-0000-0000-000000000000
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkSanitize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anonuuid := New()
		anonuuid.Sanitize(EXAMPLE_INPUT)
	}
}
