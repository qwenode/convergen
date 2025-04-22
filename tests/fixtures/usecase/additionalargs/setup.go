//go:build convergen
// +build convergen

package converter

import (
	"github.com/qwenode/convergen/tests/fixtures/data/model"
	"github.com/qwenode/convergen/tests/fixtures/data/model/abc222"
)

//go:generate convergen

// convergen
type Convergen interface {
	//:map $2 List
	DomainToModel(*model.Additional,[]abc222.Additional321) *abc222.AdditionalItem123
}
