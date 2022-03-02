package utils

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

// ----------------------------------------------------------    json  操作
var JsonApi = jsoniter.ConfigCompatibleWithStandardLibrary


func DataByJsonByte(params interface{}) []byte {
	by ,err := JsonApi.Marshal(params)
	if err != nil {
		logrus.Error("DataByJsonStr err ", err)
	}
	return by
}

func StrJsonByData(str string, data interface{}) error {
	err := JsonApi.Unmarshal([]byte(str), data)
	if err != nil {
		logrus.Error("StrJsonByData err ", err)
	}
	return err
}

func DataByJsonStr(params interface{}) string {
	by ,err := JsonApi.Marshal(params)
	if err != nil {
		logrus.Error("DataByJsonStr err ", err)
	}
	return string(by)
}

func ByteJsonByData(by []byte,data interface{}) error {
	err := JsonApi.Unmarshal(by, data)
	if err != nil {
		return err
	}
	return nil
}

/* 生成订单号   时间+机器号+用户id*/
func CreateOrderSn(prefix string, uid string)string{
	t := time.Now().Format("20060102150405")
	le := len(uid)
	return prefix + t + uid[le-2:] + uid[le-4:le-2]
}


func RandInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func RandForInt(max int, n int) int {
	rand.Seed(time.Now().UnixNano() + int64(n))
	return rand.Intn(max)
}