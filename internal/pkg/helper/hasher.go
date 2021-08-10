package helper

import (
	"crypto/sha256"
	"fmt"
)

type HashHelper struct{}

func (hh *HashHelper) Encode(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
