package api

import (
	"context"
	"errors"

	"github.com/sheelendar196/go-projects/graphql-project/graph/model"
	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
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

func (ep *EmployeeProcessor) SaveEmployee(ctx context.Context, employee model.Employee) error {
	empFactory := factory.NewEmployeeFactory(ep.employeeInteractor)
	return empFactory.Create(ctx, mapModelToDomainEmployee(&employee))
}
func (ep *EmployeeProcessor) GetEmployeeByID(ctx context.Context, empID string) (*model.Employee, error) {
	emp, err := ep.employeeInteractor.GetEmployeeByID(ctx, empID)
	if err != nil {
		return nil, err
	}
	return mapDomainEmployeeToModelEmployee(emp), nil

}

func (ep *EmployeeProcessor) GetEmployeeList(ctx context.Context) ([]*model.Employee, error) {
	empList, err := ep.employeeInteractor.GetEmployeeList(ctx)
	if err != nil {
		return nil, errors.New("No Employee found in DB")
	}
	modelList := make([]*model.Employee, len(empList))
	for i, emp := range empList {
		modelList[i] = mapDomainEmployeeToModelEmployee(emp)
	}
	return modelList, nil
}

func mapDomainEmployeeToModelEmployee(emp *domain.Employee) *model.Employee {
	if emp == nil {
		return nil
	}
	return &model.Employee{
		Name:       emp.Name,
		EmpID:      emp.EmpID,
		Mobile:     emp.Mobile,
		Email:      emp.Email,
		Department: emp.Department,
		ManagerID:  emp.ManagerID,
		Address:    emp.Address,
		IsActive:   emp.IsActive,
	}
}

func mapModelToDomainEmployee(emp *model.Employee) *domain.Employee {
	if emp == nil {
		return nil
	}
	return &domain.Employee{
		Name:       emp.Name,
		EmpID:      emp.EmpID,
		Mobile:     emp.Mobile,
		Email:      emp.Email,
		Department: emp.Department,
		ManagerID:  emp.ManagerID,
		Address:    emp.Address,
		IsActive:   emp.IsActive,
	}
}
