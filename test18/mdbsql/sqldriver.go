package mydbsql

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"time"
	"gowebtest/models"
	//"strconv"
)

type DbBuilder struct {
	DB *sql.DB
	err error
}

//var db *sql.DB
//var err error

func DbBuilder_Init(){

}

//func main() {
//	Mdb := new(DbBuilder)
//	Mdb.mysqlinit()
//	//Mdb.query()
//	//Mdb.query2()
//	Mdb.insert()
//	Mdb.update()
//	//Mdb.remove()
//}

//查询数据
func (Mdb *DbBuilder) Query() (map[int]map[string]string){
	rows, err := Mdb.DB.Query("SELECT * FROM user")
	Mdb.check(err)

	record := make(map[int]map[string]string)
	rec := make(map[string]string)


    j := 0
	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)

		for i, col := range values {
			if col != nil {
				//fmt.Println(i)
				rec[columns[i]] = string(col.([]byte))


			}
		}
		fmt.Println(rec)

		record[j] = rec
		j++
		rec = map[string]string{}

		fmt.Println(record)
		fmt.Println("query out!")
	}
	rows.Close()
	return record
}

func (Mdb *DbBuilder) Query2()  (map[int]map[string]string){
	rows, err := Mdb.DB.Query("SELECT id, userName, password, nickName, registTime, lastTimeLogin, newLoginTime, bak, online, createTime, creator, updateTime, updator FROM user")
	Mdb.check(err)
	ret := make(map[int]map[string]string)
	rec := make(map[string]string)
	j := 0
	usr := new(models.User)
	fmt.Println("Query2")
	for rows.Next(){



		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&usr.Id, &usr.UserName, &usr.Password, &usr.NickName, &usr.RegistTime, &usr.LastTimeLogin, &usr.NewLoginTime, &usr.Bak, &usr.Online, &usr.CreateTime, &usr.Creator, &usr.UpdateTime, &usr.Updator); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("userName = \"%s\", password = \"%s\", nickName = \"%s\", registTime = \"%s\", lastTimeLogin = \"%s\", newLoginTime = \"%s\", bak = \"%s\", online = \"%s\", createTime = \"%s\", creator = \"%s\", updateTime = \"%s\", updator = \"%s\"\n",usr.UserName ,usr.Password, usr.NickName, usr.RegistTime, usr.LastTimeLogin.String, usr.NewLoginTime.String, usr.Bak.String, usr.Online.String, usr.CreateTime.String, usr.Creator.String,usr.UpdateTime.String, usr.Updator.String)
		rec["Id"] = usr.Id
		rec["UserName"] = usr.UserName
		rec["Password"] = usr.Password
		rec["NickName"] = usr.NickName
		rec["RegistTime"] = usr.RegistTime
		rec["LastTimeLogin"] = usr.LastTimeLogin.String
		rec["NewLoginTime"] = usr.NewLoginTime.String
		rec["Bak"] = usr.Bak.String
		rec["Online"] = usr.Online.String
		rec["CreateTime"] = usr.CreateTime.String
		rec["Creator"] = usr.Creator.String
		rec["UpdateTime"] = usr.UpdateTime.String
		rec["Updator"] = usr.Updator.String


		ret[j] = rec
		j++
		rec = map[string]string{}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()

	return ret
}


//find user info by id
func (Mdb *DbBuilder)CheckUserInfo(id string)(map[string]string){
	ret := make(map[string]string)
	usr := new(models.User)
	rows, err := Mdb.DB.Query("SELECT userName, password, nickName, registTime, lastTimeLogin, newLoginTime, bak, online, createTime, creator, updateTime, updator FROM user where id = ?",id)
	Mdb.check(err)
	fmt.Println("-------CheckUserInfo---------")
	for rows.Next(){


		usr.Id = id
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&usr.UserName, &usr.Password, &usr.NickName, &usr.RegistTime, &usr.LastTimeLogin, &usr.NewLoginTime, &usr.Bak, &usr.Online, &usr.CreateTime, &usr.Creator, &usr.UpdateTime, &usr.Updator); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("userName = \"%s\", password = \"%s\", nickName = \"%s\", registTime = \"%s\", lastTimeLogin = \"%s\", newLoginTime = \"%s\", bak = \"%s\", online = \"%s\", createTime = \"%s\", creator = \"%s\", updateTime = \"%s\", updator = \"%s\"\n",usr.UserName ,usr.Password, usr.NickName, usr.RegistTime, usr.LastTimeLogin.String, usr.NewLoginTime.String, usr.Bak.String, usr.Online.String, usr.CreateTime.String, usr.Creator.String,usr.UpdateTime.String, usr.Updator.String)
		ret["Id"] = usr.Id
		ret["UserName"] = usr.UserName
		ret["Password"] = usr.Password
		ret["NickName"] = usr.NickName
		ret["RegistTime"] = usr.RegistTime
		ret["LastTimeLogin"] = usr.LastTimeLogin.String
		ret["NewLoginTime"] = usr.NewLoginTime.String
		ret["Bak"] = usr.Bak.String
		ret["Online"] = usr.Online.String
		ret["CreateTime"] = usr.CreateTime.String
		ret["Creator"] = usr.Creator.String
		ret["UpdateTime"] = usr.UpdateTime.String
		ret["Updator"] = usr.Updator.String
		//ret[""] = usr.UserName
		//ret = map[string]string{}

		fmt.Println(ret)
		fmt.Println("query out!")
	}
	fmt.Println("-------CheckUserInfo--- end------")
	return ret
}
//插入数据
func (Mdb *DbBuilder) Insert(username string,passwd string,nickname string)  {
	fmt.Println("---------------Insert--db=",Mdb.DB)
	stmt, err := Mdb.DB.Prepare("INSERT user (userName, password, nickName) VALUES (?, ?, ?)")
	//Mdb.check(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---------------1--")
	res, err := stmt.Exec(username,passwd,nickname)
	//res, err := stmt.Exec("fuwa_1","123123","佛系青蛙")
	//Mdb.check(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---------------2--")

	id, err := res.LastInsertId()
	//Mdb.check(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---------------31-")

	fmt.Println(id)
	stmt.Close()

}

//修改数据
func (Mdb *DbBuilder) Update(id string) int64{
	stmt, err := Mdb.DB.Prepare("UPDATE user set updateTime=?, updator=? WHERE id=?")
	Mdb.check(err)

	//res, err := stmt.Exec("",time.Now().Format("2006-01-02 15:04:05"),"root","测试更新\r\ngo直连数据库",id)
	res, err := stmt.Exec(time.Now().Format("2006-01-02 15:04:05"),"root",id)

	Mdb.check(err)

	num, err := res.RowsAffected()
	Mdb.check(err)

	fmt.Println(num)
	stmt.Close()
	return num
}

//删除数据
func (Mdb *DbBuilder) Remove(uid string) int{
	stmt, err := Mdb.DB.Prepare("DELETE FROM user WHERE id=?")
	Mdb.check(err)
	//id, err := strconv.Atoi(uid)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("---------Remove-------0---------")
	res, err := stmt.Exec(uid)
	//res, err := stmt.Exec(3)
	Mdb.check(err)
	//fmt.Println("---------Remove-------1---------")
	num, err := res.RowsAffected()
	if err != nil {
		return -1
	}
	fmt.Println(num)
	stmt.Close()
	return 0
}

func (Mdb *DbBuilder) Close() {
	Mdb.DB.Close()
}
func (Mdb *DbBuilder) check(err error) {
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}

func Mysqlinit() DbBuilder {
	var rdb DbBuilder
	rdb.DB, rdb.err = sql.Open("mysql", "root:123456@/mytest?charset=utf8")
	rdb.check(rdb.err)

	rdb.DB.SetMaxOpenConns(2000)
	rdb.DB.SetMaxIdleConns(1000)
	rdb.check(rdb.DB.Ping())
	fmt.Println(rdb.DB)
	return rdb
}

