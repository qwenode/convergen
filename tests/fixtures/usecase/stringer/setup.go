//go:build convergen
// +build convergen

package stringer

import (
	"github.com/qwenode/convergen/tests/fixtures/data/model"
	"github.com/qwenode/convergen/tests/fixtures/usecase/stringer/local"
)

//go:generate go run github.com/qwenode/convergen
type Convergen interface {
	// :stringer
	// :getter
	LocalToModel(pet *local.Pet) *model.Pet
}
