package main

import(
	"fmt"
)



func queryid(phone string) ( p Profile ,err error ){

fmt.Println("#Querying")
	rows, err := db.Query("SELECT * FROM ADDBOOK WHERE PHONE=$1", phone)
    for rows.Next() {
    err = rows.Scan(&p.Name,&p.Phone,&p.Address,&p.Email)
    }
return p, err
}
