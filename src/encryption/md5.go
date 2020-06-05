package encryption

import (
	"crypto/md5"
	"io"
	"fmt"
)

func MD5Encrypt (data string) string { 
    _md5 := md5.New()
    io.WriteString(_md5, data)
    data_encrypted := fmt.Sprintf("%x", _md5.Sum(nil))
	return data_encrypted
}