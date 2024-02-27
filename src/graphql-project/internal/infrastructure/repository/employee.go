package repository

import (
	"context"
	"errors"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
	"gorm.io/gorm"
)

type Employee struct {
	list []*domain.Employee
}

func NewEmployeeRepo(gormDB *gorm.DB) *Employee {
	return &Employee{}
}

func (e *Employee) SaveEmployee(ctx context.Context, employee *domain.Employee) error {
	e.list = append(e.list, employee)
	return nil
}

func (e *Employee) GetEmployeeByID(ctx context.Context, empID string) (*domain.Employee, error) {
	for _, emp := range e.list {
		if emp.EmpID == empID {
			return emp, nil
		}
	}

	return nil, errors.New("employee not found in db")
}

func (e *Employee) GetEmployeeList(ctx context.Context) ([]*domain.Employee, error) {
	return e.list, nil
}

func (e *Employee) UpdateEmployeeDetails(ctx context.Context, input domain.Employee) (*domain.Employee, error) {
	var employee *domain.Employee
	for i, emp := range e.list {
		if emp.EmpID == input.EmpID {
			updateEmployee(emp, input)
			e.list[i] = emp
			employee = emp
			break
		}
	}
	if employee == nil {
		return nil, errors.New("no employee found with this empID")
	}
	return employee, nil
}

func (es *Employee) DeleteEmployee(ctx context.Context, empID string) (*domain.Employee, error) {
	index := -1
	for i, emp := range es.list {
		if emp.EmpID == empID {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, errors.New("no employee found with this empID")
	}
	emp := es.list[index]
	es.list = append(es.list[:index], es.list[index+1:]...)
	return emp, nil
}

func updateEmployee(emp *domain.Employee, input domain.Employee) {
	if input.Name != "" {
		emp.Name = input.Name
	}
	if input.Address != nil {
		emp.Address = input.Address
	}
	if input.Department != nil {
		emp.Department = input.Department
	}
	if input.Email != nil {
		*emp.Email = input.EmpID
	}
	if input.Mobile != nil {
		emp.Mobile = input.Mobile
	}
	if input.IsActive != nil {
		emp.IsActive = input.IsActive
	}
	if input.ManagerID != nil {
		emp.ManagerID = input.ManagerID
	}
}
