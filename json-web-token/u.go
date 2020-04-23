package json_web_token

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Jwt struct {
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func NewHeader() Header {
	return Header{
		Alg: "HS256",
		Typ: "JWT",
	}
}

type Payload struct {
	Iss      string `json:"iss"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
	Username string `json:"username"`
	Uid      int
}

func Create(username string, id int) string {
	header := NewHeader()
	payload := Payload{
		Iss:      "redrock",
		Exp:      strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),//时间，重要
		Iat:      strconv.FormatInt(time.Now().Unix(), 10),//时间，重要
		Username: username,
		Uid:      id,
	}

	h, _ := json.Marshal(header)
	p, _ := json.Marshal(payload)
	headerBase64 := base64.StdEncoding.EncodeToString(h)
	payloadBase64 := base64.StdEncoding.EncodeToString(p)//base64编码
	str1 := strings.Join([]string{headerBase64, payloadBase64}, ".")

	key := "redrock"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str1))
	s := mac.Sum(nil)//hmac sha1散列，签名生成，会有乱码

	signature := base64.StdEncoding.EncodeToString(s)
	token := str1 + "." + signature//再进行一次base64编码
	return token
}

func CheckToken(token string) (uid int, username string, err error) {
	arr := strings.Split(token, ".")//把token根据点分成三部分，token是两个json和一个签名用点连接
	if len(arr) != 3 {//判断长度是否为3，否则token有问题
		err = errors.New("token error")
		return
	}
	_, err = base64.StdEncoding.DecodeString(arr[0])//然后base64解码出来里面的内容
	if err != nil {
		err = errors.New("token error")
		return
	}
	pay, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		err = errors.New("token error")
		return
	}
	sign, err := base64.StdEncoding.DecodeString(arr[2])//签名是第2，也是最后的部分
	if err != nil {//写的时候不能忽略err
		err = errors.New("token error")
		return
	}

	str1 := arr[0] + "." + arr[1]//签名用原始的

	key := []byte("redrock")
	mac := hmac.New(sha256.New, key)//hmac是一种标准，通过密钥生成另一个散列，加盐，盐是密码学
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	//fmt.Println(sign)
	//fmt.Println(s)
	if res := bytes.Compare(sign, s); res != 0 {//得到签名对比是否一样，0是对比相同，比对没问题就是没被篡改
		fmt.Println("test")
		err = errors.New("token error")//token只要知道是否有错，不用知道为什么错
		return
	}

	var payload Payload
	_ = json.Unmarshal(pay, &payload)
	uid=payload.Uid
	username =payload.Username
	return
}