package main

import (
	"fmt"

	compostgres "github.com/zhiwei-jian/common-go-postgres"
	rabbitmq "github.com/zhiwei-jian/common-go-rabbitmq"
	"github.com/zhiwei-jian/go-rabbitmq/msg/order"
	user "github.com/zhiwei-jian/go-rabbitmq/user"
	utils "github.com/zhiwei-jian/go-rabbitmq/utils"
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
	content := utils.Base64Decode(string(dataByte))
	fmt.Println(content)

	newUser := user.UnmarshalJsonStr2User([]byte(content))
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
	// User
	var userProcessor = &RecvPro{}
	go rabbitmq.Recv(rabbitmq.QueueExchange{
		"amqp.user.go_direct",
		"user.info",
		"direct_go",
		"direct",
		"amqp://guest:guest@10.199.196.93:30285/",
	}, userProcessor, 3)

	// Order
	var orderProcessor = &order.RecvOrder{}
	rabbitmq.Recv(rabbitmq.QueueExchange{
		"amqp.order.go_topic",
		"*.order",
		"topic_go",
		"topic",
		"amqp://guest:guest@10.199.196.93:30285/",
	}, orderProcessor, 3)
}
