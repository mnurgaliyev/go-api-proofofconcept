package main

import (
	"time"
)

//MES struct type
type MES struct {
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	DateAdded time.Time `json:"dateAdded"`
}

//MESS struct type
type MESS struct {
	MESS []MES `json:"messages"`
}

type US struct {
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Department string `json:"Department"`
}

type USS struct {
	USS []US `json:"user"`
}

type SO struct {
	Id string `json:"Id"`
	SalesOrderNumber   string    `json:"SalesOrderNumber"`
	CustomerNumber      string    `json:"CustomerNumber"`
	OrderDate string `json:"OrderDate"`
	PurchaseOrderNumber string `json:"PurchaseOrderNumber"`
}

type SOO struct {
	SOO []SO `json:"SalesOrders"`
}