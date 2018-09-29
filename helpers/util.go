package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/snluu/uuid"
	"math/rand"
)

func Random(min, max int) int {
	bettwen := max - min
	value := rand.Intn(bettwen) + min
	return value
}

func UUID() string {
	return uuid.Rand().Hex()
}

func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}
