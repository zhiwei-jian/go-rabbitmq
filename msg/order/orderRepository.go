package order

import (
	"fmt"
	"log"

	compostgres "github.com/zhiwei-jian/common-go-postgres"
)

func CreateOrder(c *compostgres.AppContext, order *Order) int {
	if order.Uid < 0 || order.GoodId < 0 {
		log.Fatal("order is invalid")
		return 0
	}
	// get insert id
	lastInsertID := 0
	err := c.Db.QueryRow("INSERT INTO orders(uid, good_id, order_time) VALUES($1,$2,$3) RETURNING id", order.Uid, order.GoodId, order.OrderTime).Scan(&lastInsertID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New order id is ", lastInsertID)
	return lastInsertID
}

func getOrderByUid(c *compostgres.AppContext, uid int) []Order {
	if uid < 0 {
		log.Fatal("User name is empty")
		return nil
	}

	stmt, err := c.Db.Prepare("SELECT id, uid, good_id, order_time FROM orders WHERE uid = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)
	defer rows.Close()
	var ordersSlice []Order
	for rows.Next() {
		order := new(Order)
		err := rows.Scan(&order.Oid, &order.Uid, &order.GoodId, &order.OrderTime)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(order.Oid, order.Uid, order.GoodId, order.OrderTime)
		ordersSlice = append(ordersSlice, *order)
	}

	return ordersSlice
}
