package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	compostgres "github.com/zhiwei-jian/common-go-postgres"
	rabbitmq "github.com/zhiwei-jian/common-go-rabbitmq"
	user "github.com/zhiwei-jian/go-rabbitmq/user"
)

/*
Implement interface Receiver
*/
type RecvPro struct{}

/*
Method of interface Receiver
*/
func (t *RecvPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	content := Base64Decode(string(dataByte))
	fmt.Println(content)

	newUser := UnmarshalJsonStr([]byte(content))
	user.Create(dbContext, &newUser)
	return nil
}

func (t *RecvPro) FailAction(dataByte []byte) error {
	fmt.Println(string(dataByte))
	fmt.Println("Failed to process data, enter db")
	return nil
}

var config = &compostgres.PostgresConfig{
	"10.199.196.93",
	31656,
	"postgres",
	"postgres",
	"k8s",
}

// var config = &compostgres.PostgresConfig{
// 	"172.28.128.5",
// 	5432,
// 	"guest",
// 	"guest",
// 	"uipdb",
// }

var dbContext, err = compostgres.ConnectDB(config)

func main() {
	var t = &RecvPro{}

	rabbitmq.Recv(rabbitmq.QueueExchange{
		"go_test",
		"go_test",
		"hello_go",
		"direct",
		"amqp://guest:guest@10.199.196.93:30285/",
	}, t, 3)
}

func Base64Decode(str string) string {
	var strs []string = strings.Split(str, "\"")
	str = strs[1]
	// reader := strings.NewReader(str)
	// decoder := base64.NewDecoder(base64.RawStdEncoding, reader)

	// buf := make([]byte, 1024)

	// dst := ""
	// for {
	// 	n, err := decoder.Read(buf)
	// 	dst += string(buf[:n])
	// 	if n < 0 || err != nil {
	// 		break
	// 	}
	// }

	rawDecodedText, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(rawDecodedText)
}

func UnmarshalJsonStr(jsonBytes []byte) user.Userinfo {
	var user user.Userinfo
	err := json.Unmarshal(jsonBytes, &user)
	if err != nil {
		fmt.Println("Failed to convert the UserInfo")
	}

	return user
}
