package user

import (
	"fmt"
	"log"
	_ "time"

	_ "github.com/lib/pq"
	compostgres "github.com/zhiwei-jian/common-go-postgres"
)

type Userinfo struct {
	Uid      int
	Name     string
	NickName string
	Age      int8
	Hobby    string `sql:"type:timestamp"`
}

// Create
func Create(c *compostgres.AppContext, user *Userinfo) {
	// get insert id
	lastInsertId := 0
	// now_str := time.Now().Format("2006-01-02 15:04:05")
	err := c.Db.QueryRow("INSERT INTO users(name,nickname,age,hobby) VALUES($1,$2,$3,$4) RETURNING id", user.Name, user.NickName, user.Age, user.Hobby).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted id is ", lastInsertId)
}

// Read
func Read(c *compostgres.AppContext) {
	rows, err := c.Db.Query("SELECT * FROM userinfo")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := new(Userinfo)
		err := rows.Scan(&p.Uid, &p.Name, &p.NickName, &p.Age, &p.Hobby)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(p.Uid, p.Name, p.NickName, p.Age, p.Hobby)
	}
}

// UPDATE
func Update(c *compostgres.AppContext, user Userinfo) {
	stmt, err := c.Db.Prepare("UPDATE userinfo SET hobby = $1, nickname = $2, age = $3 WHERE uid = $4")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(user.Hobby, user.NickName, user.Age, user.Uid)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update affect rows is ", affectNum)
}

// DELETE
func Delete(c *compostgres.AppContext, uid int) {
	stmt, err := c.Db.Prepare("DELETE FROM userinfo WHERE uid = $1")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(uid)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete affect rows is ", affectNum)
}
