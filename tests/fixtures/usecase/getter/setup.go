//go:build convergen
// +build convergen

package getter

import (
	"github.com/qwenode/convergen/tests/fixtures/data/ddd/domain"
	"github.com/qwenode/convergen/tests/fixtures/data/ddd/model"
)

// :getter:off
//
//go:generate go run github.com/qwenode/convergen
type Convergen interface {
	// DomainToModel copies domain.Pet to model.Pet.
	// :skip PhotoUrls
	// :getter
	DomainToModel(pet *domain.Pet) *model.Pet

	// DomainToModelNoGetter copies domain.Pet to model.Pet but not using getters.
	// :getter
	DomainToModelNoGetter(pet *domain.Pet) *model.Pet
}
