//go:build convergen

package typecast

import (
	"github.com/qwenode/convergen/tests/fixtures/usecase/typecast/domain"
	"github.com/qwenode/convergen/tests/fixtures/usecase/typecast/model"
)

//go:generate go run github.com/qwenode/convergen
type Convergen interface {
	// :typecast
	// DomainToModel converts domain.User to model.User.
	// typecast works:
	// - int64 -> int
	// - enums.Status -> string
	DomainToModel(*domain.User) *model.User

	// :typecast
	// ModelToDomain converts model.User to domain.User.
	// typecast works:
	// - int -> int64
	// - string -> enums.Status
	//   "enums" package will be imported automatically in the generated code!
	ModelToDomain(*model.User) *domain.User
}
