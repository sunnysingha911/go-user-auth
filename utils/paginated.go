package utils

import "gorm.io/gorm"

// Paginate returns a GORM scope for pagination
func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	if page < 1 {
		page = 1
	}
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}
	offset := (page - 1) * limit

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}
