package app

import "github.com/Delta-a-Sierra/wedding_website/internal/domain"

type Repository interface {
	domain.WeddingInfoRepo
	domain.RegistryRepo
}
