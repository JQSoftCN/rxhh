package main

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=224712 dbname=HHSoft sslmode=disable")
	checkErr(err)

	defer db.Close()

	//插入数据
	stmt, err := db.Prepare("INSERT INTO hh_report(rid,rname,datetype,defaultdate,defaultfmt) VALUES($1,$2,$3,$4,$5)")
	checkErr(err)

	defer stmt.Close()

	res, err := stmt.Exec(2, "报表3", 2, "t", "0.00")
	checkErr(err)

	//pg不支持这个函数，因为他没有类似MySQL的自增ID
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update HH_REPORT set RNAME=$1 where RID=$2")
	checkErr(err)

	res, err = stmt.Exec("报表4", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM HH_REPORT")
	checkErr(err)

	for rows.Next() {
		var rid int
		var rname string
		var datetype string
		var defaultDate string
		var defaultTime string

		err = rows.Scan(&rid, &rname, &datetype, &defaultDate, &defaultTime)
		checkErr(err)
		fmt.Println(rid, rname, datetype, defaultDate, defaultTime)
	}

	//删除数据
	stmt, err = db.Prepare("delete from HH_REPORT where RID=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
