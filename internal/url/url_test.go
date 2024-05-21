package url

import (
	"fmt"
	"testing"
)

func TestCreateUrlHash(t *testing.T) {
	const testStr = "http://www.google.com"
	const expectedLength = HASH_LENGTH
	const expectedHash = "1381a8b29f7"

	hash, err := createUniqueCode(testStr)
	if err != nil {
		t.Fatal("Hasher returned an error: ", err)
	}

	fmt.Println(len(hash))
	if len(hash) != expectedLength {
		t.Fatalf("Resulting hash is of length %d instead of the expected %d", len(hash), expectedLength)
	}

	if hash != expectedHash {
		t.Fatalf("Resulting hash differs from the one expected")
	}
}
