//go:build convergen
// +build convergen

package converter

import (
	"github.com/qwenode/convergen/tests/fixtures/usecase/embedded/domain"
	"github.com/qwenode/convergen/tests/fixtures/usecase/embedded/model"
)

//go:generate go run github.com/qwenode/convergen
type Convergen interface {
	// :getter
	// :typecast
	DomainToModel(s *domain.Concrete) (d *model.Concrete)
	// :getter
	// :typecast
	ModelToDomain(*model.Concrete) (*domain.Concrete, error)
}
