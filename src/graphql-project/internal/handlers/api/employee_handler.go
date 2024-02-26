package api

import (
	"context"

	"github.com/sheelendar196/go-projects/graphql-project/graph/model"
	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain/factory"
)

type EmployeeProcessor struct {
	employeeInteractor EmployeeService
}

func NewEmployeeProceeor(employeeInteractor EmployeeService) *EmployeeProcessor {
	return &EmployeeProcessor{
		employeeInteractor: employeeInteractor,
	}
}

func (ep *EmployeeProcessor) CreateEmployeeObj(ctx context.Context, name, empID, mobile, email, department, add, managerID string, isActive bool) error {
	empFactory := factory.NewEmployeeFactory(ep.employeeInteractor)
	return empFactory.Create(ctx, name, empID, mobile, email, department, add, managerID, isActive)

}
func (ep *EmployeeProcessor) SaveEmployee(ctx context.Context, employee model.NewEmployee) error {
	return nil
}
func (ep *EmployeeProcessor) GetEmployeeByID(ctx context.Context, empID string) (*model.Employee, error) {
	return nil, nil
}
func (ep *EmployeeProcessor) GetEmployeeList(ctx context.Context) ([]*model.Employee, error) {
	return nil, nil
}
