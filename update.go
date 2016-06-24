package main
import(
	"fmt"
	"log"
)

func update(p Profile ) ( err error ){
    fmt.Println("#Editing values!")
	pro := Profile{
		Name : p.Name,
		Phone : p.Phone,
		Address : p.Address,
		Email : p.Email,
	}

	fmt.Println(pro)


	rows, err := db.Query("UPDATE ADDBOOK SET NAME = $1, PHONE = $2, ADDRESS = $3,EMAIL = $4 WHERE PHONE = $5; ",pro.Name,pro.Phone,pro.Address,pro.Email,pro.Phone)
	 
	

	
	log.Println(rows) 
	
	return  err
}
