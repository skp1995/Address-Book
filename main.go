package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	//"io/ioutil"
)

//defining the variables

type Profile struct {
	Name    string //`json:"name"`
	Phone   string //`json:"phone"`
	Address string //`json:"address"`
	Email   string //`json:"email"`
}

var (
	db       *sql.DB
	profiles []Profile
	p, editp, p1 Profile
	phone1   []byte
)

//initialization process

const (
	db_user     = "postgres"
	db_password = "skp1995"
	db_name     = "skpdb"
)

func init() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", db_user, db_password, db_name)
	var err error

	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
}

//main.go

func main() {
	defer db.Close()
	log.Println("Hello")

	http.HandleFunc("/create", create)
	http.HandleFunc("/showById", showById)

	http.HandleFunc("/showAll", showAll)
	http.HandleFunc("/edit", editContact)

	
	http.HandleFunc("/delete", deleteById)

	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":9000", nil)
	//err := http.ListenAndServe(":9090", http.FileServer(http.Dir("./public"))) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

//creating a new contact

func create(w http.ResponseWriter, r *http.Request) {

	// insert.go
	/*
	t, _ := template.ParseFiles("details.html")
	t.Execute(w, nil)

	r.ParseForm()

	name := r.FormValue("Name")
	phone := r.FormValue("Phone")
	address := r.FormValue("Address")
	email := r.FormValue("Email")
	p := Profile{
		Name:    name,
		Phone:   phone,
		Address: address,
		Email:   email,
	}
	*/
	
	var p Profile
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)


	id, err := insert(p)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(id)

	d, err := json.Marshal(p)
	w.Header().Add("Content-Type", "application/json")
	w.Write(d)
}

func showAll(w http.ResponseWriter, req *http.Request) {
	//showing the details of all the contacts

	profiles, err := querydb()
	if err != nil {
		log.Println(err)
		return
	}
	d, err := json.Marshal(profiles)
	w.Header().Add("Content-Type", "application/json")
	w.Write(d)

}

//showing the details of a contact using the phone number as key

func showById(w http.ResponseWriter, r *http.Request) {
	var p Profile
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	p, err = queryid(p.Phone)

	if err != nil {
		log.Println(err)
		return
	}
	d, err := json.Marshal(p)
	w.Header().Add("Content-Type", "application/json")
	w.Write(d)

}

//deleting a contact using the phone number as key

func deleteById(w http.ResponseWriter, r *http.Request) {

	var p Profile
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	p1, err = queryid(p.Phone)

	err = deletecontact(p.Phone)

	if err != nil {
		log.Println(err)
		return
	}

	d, err := json.Marshal(p1)
	w.Header().Add("Content-Type", "application/json")
	w.Write(d)

}



//editing the details of a contact using the phone number as key

func editContact(w http.ResponseWriter, r *http.Request) {

	var p Profile
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	err = update(p)

	if err != nil {
		log.Println(err)
		return
	}
	d, err := json.Marshal(p)
	w.Header().Add("Content-Type", "application/json")
	w.Write(d)
	}



//end of the program
