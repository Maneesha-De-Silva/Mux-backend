package payment

import "backend/internal/types"

func (s *Service) SavePayment(payment types.Payment) (types.Payment, error) {
	if result := s.DB.Save(&payment); result.Error != nil {
		return types.Payment{}, result.Error
	}
	return payment, nil
}

func (s *Service) GetPosts() []types.Payment {
	var payments []types.Payment
	s.DB.Debug().Find(&payments)
	return payments
}

func (s *Service) GetPaymentByID(id uint64) types.Payment {
	var payment types.Payment
	s.DB.Debug().Where("id = ?", id).First(&payment)
	return payment
}


