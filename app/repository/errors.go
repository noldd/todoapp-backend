package repository

import (
	"errors"
	"todoapp-backend/app/ApiErrors"

	"gorm.io/gorm"
)

func wrapError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ApiErrors.WrapError(err, ApiErrors.ErrNotFound)
	default:
		return err
	}
}
