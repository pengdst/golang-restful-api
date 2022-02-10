package repository

import (
	"context"
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/model/domain"
)

type GormCategoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) CategoryRepository {
	return &GormCategoryRepositoryImpl{
		Db: db,
	}
}

func (repository *GormCategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	result := repository.Db.Create(&category)
	helper.PanicIfError(result.Error)

	return category
}

func (repository *GormCategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	result := repository.Db.Save(&category)
	helper.PanicIfError(result.Error)

	return category
}

func (repository *GormCategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	result := repository.Db.Delete(&category, category.Id)
	helper.PanicIfError(result.Error)
}

func (repository *GormCategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	category := domain.Category{}
	result := repository.Db.Find(&category, categoryId)
	helper.PanicIfError(result.Error)

	if result.RowsAffected < 1 {
		return category, errors.New("category not found")
	} else {
		return category, nil
	}
}

func (repository *GormCategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	var categories []domain.Category
	result := repository.Db.Find(&categories)
	helper.PanicIfError(result.Error)

	return categories
}
