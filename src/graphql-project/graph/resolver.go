package graph

import (
	"context"

	"github.com/sheelendar196/go-projects/graphql-project/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
// go:generate go run github.com/99designs/gqlgen

type EmployeeProcessor interface {
	SaveEmployee(ctx context.Context, employee model.Employee) error
	GetEmployeeList(ctx context.Context) ([]*model.Employee, error)
	GetEmployee(ctx context.Context, empID string) (*model.Employee, error)
	DeleteEmployee(ctx context.Context, empID string) (*model.Employee, error)
	UpdateEmployee(ctx context.Context, input model.Employee) (*model.Employee, error)
}

type Resolver struct {
	EmployeeProcessor EmployeeProcessor
}
