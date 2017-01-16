package core

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func GetHash(m Message) string  {
	x := m.Id/2
	x_hash := md5.New()
	x_hash.Write([]byte(strconv.Itoa(x)))

	s := strconv.Itoa(m.Id) + m.Text + hex.EncodeToString(x_hash.Sum(nil))
	s_hash := md5.New()
	s_hash.Write([]byte(s))

	return  hex.EncodeToString(s_hash.Sum(nil))
}
