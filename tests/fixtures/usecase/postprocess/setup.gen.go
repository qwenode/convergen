// Code generated by github.com/qwenode/convergen
// DO NOT EDIT.

package postprocess

import (
	"github.com/qwenode/convergen/tests/fixtures/data/domain"
	"github.com/qwenode/convergen/tests/fixtures/data/model"
	"github.com/qwenode/convergen/tests/fixtures/usecase/postprocess/local"
	_ "github.com/qwenode/convergen/tests/fixtures/usecase/postprocess/local"
)

func DomainToModel(src *domain.Pet) (dst *model.Pet, err error) {
	dst = &model.Pet{}
	PreDomainToModel(dst, *src)
	// no match: dst.ID
	// no match: dst.Category.CategoryID
	dst.Category.Name = src.Category.Name
	dst.Name = src.Name
	// no match: dst.PhotoUrls
	// no match: dst.Status
	err = PostDomainToModel(dst, *src)
	if err != nil {
		return
	}

	return
}

func ModelToDomain(src *model.Pet) (dst *domain.Pet, err error) {
	dst = &domain.Pet{}
	// no match: dst.ID
	// no match: dst.Category.ID
	dst.Category.Name = src.Category.Name
	dst.Name = src.Name
	// no match: dst.PhotoUrls
	// no match: dst.Status
	err = local.PostModelToDomain(dst, src)
	if err != nil {
		return
	}

	return
}

func PreDomainToModel(lhs *model.Pet, rhs domain.Pet) {
}

func PostDomainToModel(lhs *model.Pet, rhs domain.Pet) error {
	return nil
}
