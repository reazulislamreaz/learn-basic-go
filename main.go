package main

import "fmt"

// struct example
type Person struct {
	name string
	age int
	email string
	password string
	role string
}

func updateInfo(p *Person){
	p.name = "Reazul Islam Reaz"
	p.age = 30
	p.email = "reazul@example.com"
	p.password = "123456"
	p.role = "admin"
}
//  type User struct 
type User struct {
	name string
	score int
	level string
    items []string // slice of string
	address Address // struct as a field
}
type Address struct {
	street string
	city string
	state string
	zip string
}

func updateUser ( u *User){
	u.name = "Reazul Islam Reaz"
	u.score = 80
	u.level = "Platinum"
	u.items = []string{"item1","item2","item3"}
	u.address = Address{
		street:"123 Main St",
		city:"New York",
		state:"NY",
		zip:"10001",
	}
}


func main(){ 



   person:= Person{
	name:"Reazul Islam Reaz",
	age: 20,
	email: "reazul@example.com",
	password: "123456",
	role: "admin",
   }
   updateInfo(&person)
//    fmt.Println(person)
}

	
