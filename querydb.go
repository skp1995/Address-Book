//package main

package main

import(
	"fmt"
)



func querydb() ( profiles []Profile ,err error ){

fmt.Println("#Querying")
	rows, err := db.Query("SELECT * FROM ADDBOOK")
	
	for rows.Next() {
		
		var p Profile
		err = rows.Scan(&p.Name,&p.Phone,&p.Address,&p.Email)
		profiles = append(profiles, p)
		//fmt.Printf("%8v | %7v | %30v | %20v \n", p.Name,p.Phone,p.Address,p.Email)
	}
	
	return profiles,err
} 
