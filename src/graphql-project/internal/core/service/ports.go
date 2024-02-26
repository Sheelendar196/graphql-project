package service

import (
	"context"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
)

type EmployeeRepo interface {
	CreateEmployeeObj(ctx context.Context, name, empID, mobile, email, department, add, managerID string, isActive bool) error
	SaveEmployee(ctx context.Context, employee *domain.Employee) error
	GetEmployeeByID(ctx context.Context, empID string) (*domain.Employee, error)
	GetEmployeeList(ctx context.Context) ([]*domain.Employee, error)
}
