package factory

import (
	"context"
)

type EmployeeFactory struct {
	employeeInteractor EmployeeService
}

func NewEmployeeFactory(employeeService EmployeeService) *EmployeeFactory {
	return &EmployeeFactory{employeeInteractor: employeeService}
}

func (ef *EmployeeFactory) Create(ctx context.Context, name, empID, mobile, email, department, add, managerID string, isActive bool) error {
	return ef.employeeInteractor.CreateEmployeeObj(ctx, name, empID, mobile, email, department, add, managerID, isActive)
}
