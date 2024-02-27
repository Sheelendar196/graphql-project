package repository

import (
	"context"
	"errors"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
	"gorm.io/gorm"
)

type Employee struct {
	list []*domain.Employee
}

func NewEmployeeRepo(gormDB *gorm.DB, nrApp *newrelic.Application) *Employee {
	return &Employee{}
}

func (e *Employee) CreateEmployeeObj(ctx context.Context, name, empID, mobile, email, department, add, managerID string, isActive bool) error {
	return nil
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
