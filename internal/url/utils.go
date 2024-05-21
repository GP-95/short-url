package url

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

const HASH_LENGTH = 11

func createUniqueCode(url string) (string, error) {
	hasher := sha3.NewShake256()
	hasher.Write([]byte(url))

	hash := make([]byte, HASH_LENGTH)

	_, err := hasher.Read(hash)
	if err != nil {
		return "", err
	}

	encodedHash := hex.EncodeToString(hash)[:HASH_LENGTH]

	return encodedHash, nil
}
