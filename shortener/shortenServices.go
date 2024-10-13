package shortener

import (
	"crypto/sha1"
	"encoding/hex"
)

func Shorten(url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))

	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:6]

}
