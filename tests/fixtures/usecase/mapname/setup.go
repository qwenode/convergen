//go:build convergen
// +build convergen

package mapname

import (
	"github.com/qwenode/convergen/tests/fixtures/data/domain"
	"github.com/qwenode/convergen/tests/fixtures/data/model"
)

//go:generate go run github.com/qwenode/convergen
type Convergen interface {
	// :map Category.ID Category.CategoryID
	// :map Status.String() Status
	// :typecast
	DomainToModel(*domain.Pet) *model.Pet
}
