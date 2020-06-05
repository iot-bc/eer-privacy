package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"bytes"
)

//对明文进行填充
func Padding (plainText []byte, blockSize int) []byte {
	//计算要填充的长度
	n := blockSize - len(plainText)%blockSize
	//对原来的明文填充n个n
	temp := bytes.Repeat([]byte{byte(n)},n)
	plainText = append(plainText,temp...)
	return plainText
}

//对密文删除填充
func UnPadding (cipherText []byte) []byte {
	//取出密文最后一个字节end
	end := cipherText[len(cipherText) - 1]
	//删除填充
	cipherText = cipherText[:len(cipherText) - int(end)]
	return cipherText
}

//AEC加密（CBC模式）
func AES_CBC_Encrypt(plainText, key string) string {
    //现将类型转换成[]byte
    _plainText := []byte(plainText)
    _key := []byte(key)
	//指定加密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(_key)
	if err != nil {
		panic(err)
	}
	//进行填充
	_plainText = Padding(_plainText, block.BlockSize())
	//指定初始向量vi,长度和block的块尺寸一致
	iv := _key
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密连续数据库
	cipherText := make([]byte, len(_plainText))
	blockMode.CryptBlocks(cipherText, _plainText)
	//返回密文
	return string(cipherText[:])
}

//AEC解密（CBC模式）
func AES_CBC_Decrypt(cipherText, key string) string {
	//先将类型转换成[]byte
	_cipherText := []byte(cipherText)
	_key := []byte(key)
	//指定解密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(_key)
	if err != nil {
		panic(err)
	}
	//指定初始化向量IV,和加密的一致
	iv := _key
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//解密
	plainText := make([]byte, len(_cipherText))
	blockMode.CryptBlocks(plainText, _cipherText)
	//删除填充
	plainText = UnPadding(plainText)
	return string(plainText[:])
}