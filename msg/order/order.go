package order

import (
	"encoding/json"
	"fmt"

	utils "github.com/zhiwei-jian/go-rabbitmq/utils"
)

type Order struct {
	Oid       int // order id
	Uid       int // user id
	GoodId    int // Good id
	OrderTime int64
}

type RecvOrder struct{}

func (r *RecvOrder) Consumer(dataByte []byte) error {
	content := utils.Base64Decode(string(dataByte))
	fmt.Println(content)

	order := UnmarshalJsonStr2Order([]byte(content))
	fmt.Println(order)
	return nil
}

func (t *RecvOrder) FailAction(dataByte []byte) error {
	fmt.Println(string(dataByte))
	fmt.Println("Failed to process order data")
	return nil
}

func UnmarshalJsonStr2Order(jsonBytes []byte) Order {
	var order Order
	err := json.Unmarshal(jsonBytes, &order)
	if err != nil {
		fmt.Println("Failed to convert the UserInfo")
	}

	return order
}
