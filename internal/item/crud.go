package item

import "backend/internal/types"

func (s *Service) CreateItem(item types.Item) {
	s.DB.Debug().Save(&item)
}

func (s *Service) GetPosts() []types.Item {
	var items []types.Item
	s.DB.Debug().Find(&items)
	return items
}
