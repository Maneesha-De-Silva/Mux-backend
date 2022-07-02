package employee

import "backend/internal/types"

func (s *Service) SaveEmployee(employee types.Employee) (types.Employee, error) {
	if result := s.DB.Save(&employee); result.Error != nil {
		return types.Employee{}, result.Error
	}
	return employee, nil
}

func (s *Service) GetPosts() []types.Employee {
	var employees []types.Employee
	s.DB.Debug().Find(&employees)
	return employees
}

func (s *Service) GetEmployeeByID(id uint64) types.Employee {
	var employee types.Employee
	s.DB.Debug().Where("id = ?", id).First(&employee)
	return employee
}

func (s *Service) DeleteEmployee(id uint64) {
	s.DB.Debug().Delete(&types.Employee{}, id)
}
