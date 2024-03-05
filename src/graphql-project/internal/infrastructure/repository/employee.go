package repository

import (
	"context"
	"errors"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
	"github.com/sheelendar196/go-projects/graphql-project/internal/infrastructure/models/tables"
	"gorm.io/gorm"
)

type Employee struct {
	list []*tables.Employee
}

func NewEmployeeRepo(gormDB *gorm.DB) *Employee {
	return &Employee{}
}

func (e *Employee) SaveEmployee(ctx context.Context, employee *domain.Employee) error {
	e.list = append(e.list, tables.ToTable(employee))
	return nil
}

func (e *Employee) GetEmployeeByID(ctx context.Context, empID string) (*domain.Employee, error) {
	for _, emp := range e.list {
		if emp.EmpID == empID {
			return emp.ToDomain(), nil
		}
	}

	return nil, errors.New("employee not found in db")
}

func (e *Employee) GetEmployeeList(ctx context.Context) ([]*domain.Employee, error) {
	empList := make([]*domain.Employee, len(e.list))
	for i, emp := range e.list {
		empList[i] = emp.ToDomain()
	}
	return empList, nil
}

// UpdateEmployeeDetails update employee into table on once actual table will be there this logic removed.
func (e *Employee) UpdateEmployeeDetails(ctx context.Context, input domain.Employee) (*domain.Employee, error) {
	var employee *tables.Employee
	for i, emp := range e.list {
		if emp.EmpID == input.EmpID {
			tables.UpdateEmployee(emp, input)
			e.list[i] = emp
			employee = emp
			break
		}
	}
	if employee == nil {
		return nil, errors.New("no employee found with this empID")
	}
	return employee.ToDomain(), nil
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
	return emp.ToDomain(), nil
}
