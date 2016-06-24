
package main
import(
	"fmt"
)

func insert(p Profile ) (lastId string, err error ){
    fmt.Println("#Insertion of values!")
	pro := Profile{
		Name: p.Name,
		Phone: p.Phone,
		Address: p.Address,
		Email: p.Email,

	}
	err = db.QueryRow("INSERT INTO ADDBOOK(NAME,PHONE,ADDRESS,EMAIL) VALUES($1,$2,$3,$4) returning NAME;",pro.Name,pro.Phone,pro.Address,pro.Email).Scan(&lastId)
	return lastId, err
}
