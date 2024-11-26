package entity

type CEPRepositoryInterface interface {
	Get(string) error
	IsValid(string) bool
}
