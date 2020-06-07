package privacy

import (
	"../encryption"
	"crypto/rand"
	"math/big"
	"fmt"
)

const primeHex = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371"
const generatorHex = "A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5"


func AnonymizeDevice (deviceid string) string {
	priv := &encryption.PrivateKey{
		PublicKey: encryption.PublicKey{
			G: fromHex(generatorHex),
			P: fromHex(primeHex),
		},
		X: fromHex("42"),
	}
	priv.Y = new(big.Int).Exp(priv.G, priv.X, priv.P)
	message := []byte(deviceid)
	c1, c2, err := encryption.ElGamalEncrypt(rand.Reader, &priv.PublicKey, message)
    fmt.Println(err)
	c1_str := c1.String()
	c2_str := c2.String()
	//将c1_str和c2_str拼接而得到的hash串与diviceid对应放入区块链上
    
	return encryption.MD5Encrypt(c1_str + c2_str)
}

func GetRealIDFromBlockChain (c1_c2_str string) string {
	//根据c1_str和c2_str拼接而成的hash在区块链上找到对应的设备真实id
	deviceid := ""
	return deviceid
}

func GetFakeIDFromBlockChain(deviceid string) string {
	return ""
}

func fromHex(hex string) *big.Int {
	n, ok := new(big.Int).SetString(hex, 16)
	if !ok {
		panic("failed to parse hex number")
	}
	return n
}