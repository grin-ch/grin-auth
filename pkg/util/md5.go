package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(src []byte) string {
	w := md5.New()
	io.WriteString(w, string(src))
	return fmt.Sprintf("%x", w.Sum(nil))
}
