package service

import (
	"context"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
)

type EmployeeInteractor struct {
	Repo EmployeeRepo
}

// NewEmployeeInteractor return valid employee interactor.
func NewEmployeeInteractor(repo EmployeeRepo) *EmployeeInteractor {
	if repo == nil {
		return nil
	}
	return &EmployeeInteractor{Repo: repo}
}

func (es EmployeeInteractor) CreateEmployeeObj(ctx context.Context, name, empID, mobile, email, department,
	add, managerID string, isActive bool) error {

	emp := domain.GetEmployee(name, empID, mobile, email, department, add, managerID, isActive)
	return es.SaveEmployee(ctx, emp)
}

func (es EmployeeInteractor) SaveEmployee(ctx context.Context, employee *domain.Employee) error {
	return es.Repo.SaveEmployee(ctx, employee)
}
func (es EmployeeInteractor) GetEmployeeByID(ctx context.Context, empID string) (*domain.Employee, error) {
	return es.GetEmployeeByID(ctx, empID)
}
func (es EmployeeInteractor) GetEmployeeList(ctx context.Context) ([]*domain.Employee, error) {
	return es.Repo.GetEmployeeList(ctx)
}
