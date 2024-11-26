package usecase

import (
	"strings"

	"github.com/paulnune/goexpert-weather/orchestrator-api/internal/entity"
)

type CEPInputDTO struct {
	CEP string
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

func (c *GetCEPUseCase) Execute(input CEPInputDTO) (CEPOutputDTO, error) {
	cep := entity.CEP{
		CEP: input.CEP,
	}

	cep_resp, err := c.CEPRepository.Get(cep.CEP)
	if err != nil && strings.Contains(string(cep_resp), "Http 400") {
		return CEPOutputDTO{}, err
	}

	cep_dto, err := c.CEPRepository.Convert(cep_resp)
	if err != nil {
		return CEPOutputDTO{}, err
	}

	dto := CEPOutputDTO{
		CEP:         cep_dto.CEP,
		Logradouro:  cep_dto.Logradouro,
		Complemento: cep_dto.Complemento,
		Bairro:      cep_dto.Bairro,
		Localidade:  cep_dto.Localidade,
		UF:          cep_dto.UF,
		IBGE:        cep_dto.IBGE,
		GIA:         cep_dto.GIA,
		DDD:         cep_dto.DDD,
		SIAFI:       cep_dto.SIAFI,
	}

	return dto, nil
}
