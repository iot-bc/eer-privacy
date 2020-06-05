package main
import (
	"fmt"
	"./privacy"
	"./encryption"
	//"./completeness"
)

func main(){
	letters := [10] string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println(letters[1])
    key := privacy.GenerateRandomKey()
	fmt.Println(key)

	idMap := privacy.AnonymizeDevice("testDeviceID")
	fmt.Println(idMap["testDeviceID"])

	fmt.Println(encryption.MD5Encrypt("abc123"))

	plainText := "testText"
	encryptText := encryption.AES_CBC_Encrypt(plainText, key)
	fmt.Println(encryptText)
	decryptText := encryption.AES_CBC_Decrypt(encryptText, key)
    fmt.Println(decryptText)

    
}