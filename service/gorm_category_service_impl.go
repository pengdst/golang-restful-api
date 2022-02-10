package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"pengdst/golang-restful-api/exception"
	"pengdst/golang-restful-api/helper"
	"pengdst/golang-restful-api/model/domain"
	"pengdst/golang-restful-api/model/web"
	"pengdst/golang-restful-api/repository"
)

type GormCategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewGormCategoryService(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &GormCategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

func (service *GormCategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Create(ctx, nil, category)

	return helper.ToCategoryResponse(category)
}

func (service *GormCategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	category, err := service.CategoryRepository.FindById(ctx, nil, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, nil, category)

	return helper.ToCategoryResponse(category)
}

func (service *GormCategoryServiceImpl) Delete(ctx context.Context, categoryId int) {

	category, err := service.CategoryRepository.FindById(ctx, nil, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, nil, category)
}

func (service *GormCategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {

	category, err := service.CategoryRepository.FindById(ctx, nil, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *GormCategoryServiceImpl) GetAll(ctx context.Context) []web.CategoryResponse {
	categories := service.CategoryRepository.GetAll(ctx, nil)

	return helper.ToCategoryResponses(categories)
}
