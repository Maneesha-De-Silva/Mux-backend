package item

import "backend/internal/types"

func (s *Service) CreateItem(item types.Item) (types.Item, error) {
	if result := s.DB.Save(&item); result.Error != nil {
		return types.Item{}, result.Error
	}
	return item, nil
}

func (s *Service) GetPosts() []types.Item {
	var items []types.Item
	s.DB.Debug().Find(&items)
	return items
}
