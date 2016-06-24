package main

import(
	"fmt"
)

func deletecontact(phone string) ( err error ){

    rows, err := db.Query("DELETE FROM ADDBOOK WHERE PHONE=$1", phone)

    for rows.Next() {
    err = rows.Scan(&p.Name,&p.Phone,&p.Address,&p.Email)
    }
    fmt.Println("Contact deleted")
return err
}