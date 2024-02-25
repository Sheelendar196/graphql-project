package graph

import "github.com/sheelendar196/go-projects/graphql-project/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
// go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	list []*model.Employee
}

/*

To create graphql dummy templet: go run github.com/99designs/gqlgen init 
do changes in schema.graphqls  file and then run :
go run github.com/99designs/gqlgen generate   

after generate you can run server:
 go run server.go  

 // queries and mutation for employee

mutation createEmployee{
  createEmployee(input:{name:"sheelendar",empID:"em2342",mobile:"948329433",
    email:"she@gmail.com",address:"ballendure",managerID:"em4532",
    isActive:true,department:"IT"}){
		name
    empID
    mobile
    email
    address
    managerID
    isActive
    department
  }
}


query GetEmployeeList{
  employees {
    name
    empID
    mobile
  }
}

query employees{
  getEmployeeList {
    name
    empID
    mobile
  }
}
*/