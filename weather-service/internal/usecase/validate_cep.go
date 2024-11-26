package usecase

import (
	"github.com/paulnune/goexpert-weather/orchestrator-api/internal/entity"
)

type ValidateCEPInputDTO struct {
	CEP string
}

type ValidateCEPOutputDTO struct {
	IsValid bool
}

type ValidateCEPUseCase struct {
	CEPRepository entity.CEPRepositoryInterface
}

func NewValidateCEPUseCase(cep_repository entity.CEPRepositoryInterface) *ValidateCEPUseCase {
	return &ValidateCEPUseCase{
		CEPRepository: cep_repository,
	}
}

func (c *ValidateCEPUseCase) Execute(input ValidateCEPInputDTO) bool {
	return c.CEPRepository.IsValid(input.CEP)
}
