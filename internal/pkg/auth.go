package pkg

import (
	"crypto/md5"
	"fmt"
	"io"
)

const (
	// TODO: use config
	salt1 = "@#$%"
	salt2 = "^&*()"
)

func GetPasswordHash(name string, password string) string {
	h := md5.New()
	_, _ = io.WriteString(h, password)

	pwMd5 := fmt.Sprintf("%x", h.Sum(nil))

	// salt1 + name + salt2 + MD5
	_, _ = io.WriteString(h, salt1)
	_, _ = io.WriteString(h, name)
	_, _ = io.WriteString(h, salt2)
	_, _ = io.WriteString(h, pwMd5)

	return fmt.Sprintf("%x", h.Sum(nil))
}
