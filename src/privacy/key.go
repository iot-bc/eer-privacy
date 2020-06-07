package privacy

import (
	"math/rand"
	"time"
)

func GenerateRandomKey() string {
	letters := [10] string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	numbers := [10] string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	key := ""

	for i := 0; i < 8; i++ {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(9)
		key += letters[random]
	}

	for i := 0; i < 8; i++ {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(9)
		key += numbers[random]
	}

	return key
}

func putKeyToBlockChain(userid, key string) bool {
	//将该用户的id与对应的key放入区块链上
	return false
}

func getKeyFromBlockChain(userid string) string {
	//根据用户的id在区块链上获取用户的id
	key := ""
	return key
}
