package privacy

import (
	"../encryption"
)

func AESEncryptData(key, data string) string {
	data_encrypted := encryption.AES_CBC_Encrypt(data, key)
	return data_encrypted
}

func AESDecryptData(key, data string) string {
	data_decrypted := encryption.AES_CBC_Decrypt(data, key)
	return data_decrypted
}