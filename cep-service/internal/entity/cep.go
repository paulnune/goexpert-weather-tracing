package entity

type CEP struct {
	CEP         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	UF          string
	IBGE        string
	GIA         string
	DDD         string
	SIAFI       string
}

func NewCEP(cep, logradouro, complemento, bairro, localidade, uf, ibge, gia, ddd, siafi string) *CEP {
	new_cep := &CEP{
		CEP:         cep,
		Logradouro:  logradouro,
		Complemento: complemento,
		Bairro:      bairro,
		Localidade:  localidade,
		UF:          uf,
		IBGE:        ibge,
		GIA:         gia,
		DDD:         ddd,
		SIAFI:       siafi,
	}
	return new_cep
}
