package order

import (
	"context"
	"encoding/json"
	"fmt"

	// "time"

	"github.com/zhiwei-jian/go-rabbitmq/config"
	"github.com/zhiwei-jian/go-rabbitmq/redis"
	"github.com/zhiwei-jian/go-rabbitmq/user"
	utils "github.com/zhiwei-jian/go-rabbitmq/utils"
)

var (
	Ctx = context.TODO()
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

	dbContext := config.GetDbContext()
	defer dbContext.Db.Close()
	user, _ := user.GetUserById(dbContext, order.Uid)
	if user == nil {
		fmt.Println("User does not exist")
		return nil
	}
	CreateOrder(dbContext, &order)

	redisContext, err := redis.ConnectRedis(config.RedisConfig)
	if err != "" {
		fmt.Println("Failed to process order data")
		return nil
	}

	defer redisContext.RedisClient.Close()

	// count, error := redisContext.RedisClient.Get(Ctx, string(rune(order.Uid))).Int()
	// redisContext.RedisClient.Set(Ctx, string(rune(order.Uid)), count+1, 0)
	// fmt.Println(count)
	// if error != nil {
	// 	fmt.Println("Failed to Get order data from redis")
	// 	return nil
	// }

	signatureID := getUserSignatureID(user)
	count, error := redisContext.RedisClient.HGet(Ctx, "orders", signatureID).Int()
	if error != nil {
		fmt.Println("Failed to Get order data from redis")
	}

	count++
	fmt.Println("User " + string(order.Uid) + " to Get order data from redis")
	redisContext.RedisClient.HSet(Ctx, "orders", signatureID, int64(count))
	return nil
}

func getUserSignatureID(user *user.Userinfo) string {
	if user == nil {
		return ""
	}

	return user.Name + ":" + string(user.Uid)
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
