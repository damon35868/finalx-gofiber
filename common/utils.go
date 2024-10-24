package common

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func MD5(s string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x\n", sum)
}

func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

func GenerateOrderNo() string {
	randomNum := rand.Intn(1000)
	date := GetFormatTime(time.Now())
	return fmt.Sprintf("%s%d%03d", date, GetTimestamp(), randomNum)
}
