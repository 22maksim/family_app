package service

import (
	"context"

	"github.com/22maksim/family_app/internal/family/model"
	"github.com/22maksim/family_app/storage/family"
)

type FamilyService interface {
	CreateFamily(ctx context.Context, dto model.CreateFamilyRequestDto) (string, error)
	AddRelatives(ctx context.Context, dto model.AddRelativesRequestDto) error
	RemoveRelative(ctx context.Context, dto model.RemoveRelativeRequestDto) error
	GetRelatives(ctx context.Context, dto model.GetRelativesRequestDto) (model.GetRelativesResponseDto, error)
	GetRelative(ctx context.Context, user_id string) (model.Relative, error)
}

func NewFamilyService(db family.FamilyRepository) FamilyService {
	return &FamilyServiceImpl{db: db}
}

type FamilyServiceImpl struct {
	db family.FamilyRepository
}

func (fs *FamilyServiceImpl) CreateFamily(ctx context.Context, dto model.CreateFamilyRequestDto) (string, error) {

	return "", nil
}

func (fs *FamilyServiceImpl) AddRelatives(ctx context.Context, dto model.AddRelativesRequestDto) error {
	return nil
}

func (fs *FamilyServiceImpl) RemoveRelative(ctx context.Context, dto model.RemoveRelativeRequestDto) error {
	return nil
}

func (fs *FamilyServiceImpl) GetRelatives(ctx context.Context, dto model.GetRelativesRequestDto) (model.GetRelativesResponseDto, error) {
	return model.GetRelativesResponseDto{}, nil
}

func (fs *FamilyServiceImpl) GetRelative(ctx context.Context, user_id string) (model.Relative, error) {
	return model.Relative{}, nil
}
