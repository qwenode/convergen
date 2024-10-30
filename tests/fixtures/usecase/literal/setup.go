//go:build convergen

package literal

import (
	"github.com/qwenode/convergen/tests/fixtures/data/domain"
	"github.com/qwenode/convergen/tests/fixtures/data/model"
)

//go:generate go run github.com/qwenode/convergen
type Convergen interface {
	// :literal  Name   "abc  def"
	DomainToModel(*domain.Pet) *model.Pet
	ModelToDomain(*model.Pet) *domain.Pet
}
