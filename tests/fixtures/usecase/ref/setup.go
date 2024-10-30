//go:build convergen
// +build convergen

package ref

import (
	"github.com/qwenode/convergen/tests/fixtures/data/domain"
	"github.com/qwenode/convergen/tests/fixtures/data/model"
)

//go:generate go run github.com/qwenode/convergen
type Convergen interface {
	// :conv CatDomainToModel Category
	DomainToModel(*domain.Pet) *model.Pet

	// :map ID CategoryID
	// :typecast
	CatDomainToModel(*domain.Category) model.Category
}
