package usecase

import (
	"github.com/paulnune/goexpert-weather/input-api/internal/entity"
)

type CEPInputDTO struct {
	CEP string `json:"cep"`
}

type CEPOutputDTO struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

type GetCEPUseCase struct {
	CEPRepository entity.CEPRepositoryInterface
}

func NewGetCEPUseCase(cep_repository entity.CEPRepositoryInterface) *GetCEPUseCase {
	return &GetCEPUseCase{
		CEPRepository: cep_repository,
	}
}

func (c *GetCEPUseCase) Execute(input CEPInputDTO) error {
	cep := entity.CEP{
		CEP: input.CEP,
	}

	err := c.CEPRepository.Get(cep.CEP)
	if err != nil {
		return err
	}

	return nil
}
