package main

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	_"github.com/denisenkom/go-mssqldb"
)

// Database variables
var connString string

func init() {
	connString = "server=ANONYMOUS;user id=sa;password=1234;port=1433;database=cube_db"
}

func dbGetConn() *sql.DB {
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}

	return db
}

func dbGetMessages() (messages []MES, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT subject, body, dateAdded FROM wmTriggerNotification")

	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var Subject, Body string
		var DateAdded time.Time

		err = rows.Scan(&Subject, &Body, &DateAdded)
		if err != nil {
			log.Printf(err.Error())
		}

		messages = append(messages, MES{
			Subject:   Subject,
			Body:      Body,
			DateAdded: DateAdded,
		})

		count++
	}

	return messages, count
}

func dbGetMessageByUser(userID string) (messages []MES, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT subject, body, dateAdded FROM wmTriggerNotification WHERE recipientId = ?1", userID)

	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var Subject, Body string
		var DateAdded time.Time

		err = rows.Scan(&Subject, &Body, &DateAdded)
		if err != nil {
			log.Printf(err.Error())
		}

		messages = append(messages, MES{
			Subject:   Subject,
			Body:      Body,
			DateAdded: DateAdded,
		})

		count++
	}



	return messages, count
}

func dbGetByUser(userID string) (user []US, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT FirstName, LastName, Department FROM aspnet_Profile WHERE UserId=?1 ", userID)

	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var FirstName string
		var LastName string
		var Department string

		err = rows.Scan(&FirstName, &LastName, &Department)
		if err != nil {
			log.Printf(err.Error())
		}

		user = append(user, US{
			FirstName: FirstName,
			LastName: LastName,
			Department: strings.TrimSpace(Department),
		})

		count++
	}


	return user, count
}

func dbGetSalesOrders() (salesOrder []SO, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT sono,custno,ORDERDATE,PONO FROM somain WHERE ORD_TYPE NOT IN ('Cancel','Closed ')")

	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}
	defer rows.Close()

	var Id int

	for rows.Next() {
		var SalesOrderNumber, CustomerNumber, PurchaseOrderNumber string
		var OrderDate string

		err = rows.Scan(&SalesOrderNumber, &CustomerNumber, &OrderDate, &PurchaseOrderNumber)
		if err != nil {
			log.Printf(err.Error())
		}

		Id++

		salesOrder = append(salesOrder, SO{
			Id: strconv.Itoa(Id),
			SalesOrderNumber: SalesOrderNumber,
			CustomerNumber: CustomerNumber,
			OrderDate: strings.Replace(strings.Replace(OrderDate, "T", " ", 1), "Z", "", 1),
			PurchaseOrderNumber: strings.TrimSpace(PurchaseOrderNumber),
		})
		count++
	}
	return salesOrder, count
}