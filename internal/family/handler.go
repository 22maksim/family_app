package family

import (
	"context"

	"github.com/22maksim/family_app/api/proto/family"
	"github.com/22maksim/family_app/internal/family/converter"
	"github.com/22maksim/family_app/internal/family/model"
	"github.com/22maksim/family_app/internal/family/service"
)

type familyServiceGrpcImpl struct {
	family.UnimplementedFamilyServiceServer
	service   service.FamilyService
	converter converter.Mapper
}

func NewFamilyServiceGrpcImpl(srv service.FamilyService, cnvrt converter.Mapper) *familyServiceGrpcImpl {
	return &familyServiceGrpcImpl{
		service:   srv,
		converter: cnvrt,
	}
}

func (f *familyServiceGrpcImpl) CreateFamily(
	ctx context.Context, req *family.CreateFamilyRequest) (*family.CreateFamilyResponse, error) {
	dto := model.CreateFamilyRequestDto{Surname: req.Surname, CreatorId: req.CreatorId}

	createdFamilyId, err := f.service.CreateFamily(ctx, dto)

	if err != nil {
		return nil, err
	}

	return &family.CreateFamilyResponse{FamilyId: createdFamilyId}, nil
}

func (f *familyServiceGrpcImpl) AddRelatives(
	ctx context.Context, req *family.AddRelativesRequest) (*family.AddRelativesResponse, error) {
	dto := f.converter.ToDtoAddRelativesRequest(req)

	err := f.service.AddRelatives(ctx, dto)

	if err != nil {
		return nil, err
	}

	return &family.AddRelativesResponse{}, nil
}

func (f *familyServiceGrpcImpl) RemoveRelative(
	ctx context.Context, req *family.RemoveRelativeRequest) (*family.RemoveRelativeResponse, error) {
	dto := f.converter.ToDtoRemoveRelativeRequest(req)

	err := f.service.RemoveRelative(ctx, dto)

	if err != nil {
		return nil, err
	}

	return &family.RemoveRelativeResponse{}, nil
}

func (f *familyServiceGrpcImpl) GetRelatives(ctx context.Context, req *family.GetRelativesRequest) (*family.GetRelativesResponse, error) {
	dto := f.converter.ToDtoGetRelativesRequest(req)

	respDto, err := f.service.GetRelatives(ctx, dto)

	if err != nil {
		return nil, err
	}

	result := make([]*family.Relative, 0, len(respDto.Relatives))
	for _, v := range respDto.Relatives {
		result = append(result, f.converter.ToRelativeFromApi(&v))
	}

	return &family.GetRelativesResponse{Relatives: result}, nil
}

func (f *familyServiceGrpcImpl) GetRelative(ctx context.Context, req *family.GetRelativeRequest) (*family.GetRelativeResponse, error) {
	relative, err := f.service.GetRelative(ctx, req.GetUserId())

	if err != nil {
		return nil, err
	}

	return &family.GetRelativeResponse{Relative: f.converter.ToRelativeFromApi(&relative)}, nil
}
