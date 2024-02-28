//To create graphql dummy templet: 
go run github.com/99designs/gqlgen init 

//do changes in schema.graphqls  file and then run :
go run github.com/99designs/gqlgen generate   

//after generate you can run server:
 go run main.go  

//To generate mock
// Go to interface folder 
    cd internal/core/domain/worker

// run mockery    
mockery --name EmployeeService












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
  getEmployeeList {
    name
    empID
    mobile
    email
  }
} 

query GetEmployeeByID{
  getEmployee(empID:"em1232") {
    name
    empID
    mobile
    empID
    department
    address
    managerID
    email
  }
}

query DeleteEmployeeBy{
  deleteEmployee(empID:"em1232") {
    name
    empID
    mobile
    empID
    department
    address
    managerID
    email
  }
}

mutation updateEmployee{
  updateEmployee(input:{name:"shayam",empID:"em8343"
    ,address:"marathalli"}){
    name
    empID
    address
  }
}
