package api

import (
	"context"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
)

type EmployeeService interface {
	SaveEmployee(ctx context.Context, employee *domain.Employee) error
	GetEmployeeByID(ctx context.Context, empID string) (*domain.Employee, error)
	GetEmployeeList(ctx context.Context) ([]*domain.Employee, error)
}
