package completeness

import (
	"../encryption"
)

func MD5EncryptData(data string) (data, data_encrypted string) {
	data = data
	data_encrypted = encryption.MD5Encrypt(data)
	return 
}