package rating

import "backend/internal/types"

func (s *Service) SaveRating(rating types.Rating) (types.Rating, error) {
	if result := s.DB.Save(&rating); result.Error != nil {
		return types.Rating{}, result.Error
	}
	return rating, nil
}

func (s *Service) GetPosts() []types.Rating {
	var ratings []types.Rating
	s.DB.Debug().Find(&ratings)
	return ratings
}

func (s *Service) GetRatingByID(id uint64) types.Rating {
	var rating types.Rating
	s.DB.Debug().Where("id = ?", id).First(&rating)
	return rating
}
