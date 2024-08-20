package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRespositoryDb struct {
	client *sql.DB
}

func (d CustomerRespositoryDb) FindAll() ([]Customer, error) {

	findAllsql := "select * from customers"
	rows, err := d.client.Query(findAllsql)
	if err != nil {
		log.Println("Error while querying cutomer table " + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateOfBirth, &c.Status, &c.ZipCode)

		if err != nil {
			log.Println("Error while querying cutomer table " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	fmt.Println(customers)
	return customers, nil
}

func NewCustomerRespositoryDb() CustomerRespositoryDb {
	// client, err := sql.Open("mysql", "root:23602@Game@/banking")
	client, err := sql.Open("mysql", "app_user:Password98#@/banking")

	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRespositoryDb{client}

}
