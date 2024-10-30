// Code generated by github.com/qwenode/convergen
// DO NOT EDIT.

package nocase

import (
	"github.com/qwenode/convergen/tests/fixtures/usecase/nocase/model"
)

type ModelA struct {
	ID   uint64
	Name string
}

func (a *ModelA) name() string {
	return a.Name
}

type ModelB struct {
	id   uint64
	name string
}

// AtoB demonstrates local to local copy with case-insensitive field matching.
// It shows that a private getter precedence over its (exported) counterpart field.
func AtoB(src *ModelA) (dst *ModelB) {
	dst = &ModelB{}
	dst.id = src.ID
	dst.name = src.name()

	return
}

func BtoA(src *ModelB) (dst *ModelA) {
	dst = &ModelA{}
	dst.ID = src.id
	dst.Name = src.name

	return
}

// BtoUser demonstrates copy an internal to external package type.
// It skips private fields (and getters) in the latter type.
func BtoUser(src *ModelB) (dst *model.User) {
	dst = &model.User{}
	dst.Name = src.name

	return
}

// UserToB demonstrates copy an external package type to internal.
// It skips private fields (and getters) in the former type.
func UserToB(src *model.User) (dst *ModelB) {
	dst = &ModelB{}
	// no match: dst.id
	dst.name = src.Name

	return
}
