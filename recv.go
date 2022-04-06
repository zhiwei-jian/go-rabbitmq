package main

import (
	"fmt"
	"time"

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
	time.Sleep(1 * time.Second)
	var user = new(user.Userinfo)
	user.Age = 123
	return nil
}

func (t *RecvPro) FailAction(dataByte []byte) error {
	fmt.Println(string(dataByte))
	fmt.Println("Failed to process data, enter db")
	return nil
}

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
