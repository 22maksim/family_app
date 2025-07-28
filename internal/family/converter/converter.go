package converter

import (
	api_family "github.com/22maksim/family_app/api/proto/family"
	"github.com/22maksim/family_app/internal/family/model"
)

type Mapper interface {
	ToDtoCreateFamilyRequest(req *api_family.CreateFamilyRequest) model.CreateFamilyRequestDto
	ToDtoAddRelativesRequest(req *api_family.AddRelativesRequest) model.AddRelativesRequestDto
	ToDtoGetRelativesRequest(req *api_family.GetRelativesRequest) model.GetRelativesRequestDto
	ToDtoGetRelativesResponse() model.GetRelativesResponseDto
	ToDtoRemoveRelativeRequest(req *api_family.RemoveRelativeRequest) model.RemoveRelativeRequestDto
	ToRelativeFromApi(rel *model.Relative) *api_family.Relative
}

type MapperImpl struct{}

func NewConverter() Mapper {
	return &MapperImpl{}
}

func (m *MapperImpl) ToDtoCreateFamilyRequest(req *api_family.CreateFamilyRequest) model.CreateFamilyRequestDto {
	return model.CreateFamilyRequestDto{}
}

func (m *MapperImpl) ToDtoAddRelativesRequest(req *api_family.AddRelativesRequest) model.AddRelativesRequestDto {
	return model.AddRelativesRequestDto{}
}

func (m *MapperImpl) ToDtoGetRelativesRequest(req *api_family.GetRelativesRequest) model.GetRelativesRequestDto {
	return model.GetRelativesRequestDto{}
}

func (m *MapperImpl) ToDtoGetRelativesResponse() model.GetRelativesResponseDto {
	return model.GetRelativesResponseDto{}
}

func (m *MapperImpl) ToDtoRemoveRelativeRequest(req *api_family.RemoveRelativeRequest) model.RemoveRelativeRequestDto {
	return model.RemoveRelativeRequestDto{}
}

func (m *MapperImpl) ToRelativeFromApi(rel *model.Relative) *api_family.Relative {
	return &api_family.Relative{}
}
