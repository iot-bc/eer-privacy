package privacy

import (
	"../encryption"
)

func AESEncryptData(deviceid, data string) string {
	key := GetKeyFromBlockChain(deviceid)
	data_encrypted := encryption.AES_CBC_Encrypt(data, key)
	return data_encrypted
}

func AESDecryptData(deviceid, data string) string {
	key := GetKeyFromBlockChain(deviceid)
	data_decrypted := encryption.AES_CBC_Decrypt(data, key)
	return data_decrypted
}