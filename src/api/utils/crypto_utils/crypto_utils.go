package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(str string) string {
	md5Hash := md5.New()
	defer md5Hash.Reset()

	md5Hash.Write([]byte(str))
	return hex.EncodeToString(md5Hash.Sum(nil))
}
