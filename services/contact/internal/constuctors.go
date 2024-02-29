package internal

import (
	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/usecase"
)

func NewContactRepository() repository.ContactRepository {
	return nil
}

func NewContactUseCase(repo repository.ContactRepository) usecase.ContactUseCase {
	return usecase.NewContactUseCase(NewContactRepository())
}

func NewContactDelivery(useCase usecase.ContactUseCase) delivery.ContactDelivery {
	return nil
}
