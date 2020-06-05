package privacy

import (
	"../encryption"
)

func AnonymizeDevice (deviceid string) map[string]string {
	//首先从链上拿下这个map再往此map里加一条记录？还是直接生成一个新的map放入链上？
	idMap := make(map[string]string)
	idMap[deviceid] = encryption.ElGamalEncrypt(deviceid)
	return idMap
}