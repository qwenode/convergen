// Code generated by github.com/qwenode/convergen
// DO NOT EDIT.

package literal

import (
	"github.com/qwenode/convergen/tests/fixtures/data/domain"
	"github.com/qwenode/convergen/tests/fixtures/data/model"
)

func DomainToModel(src *domain.Pet) (dst *model.Pet) {
	dst = &model.Pet{}
	// no match: dst.ID
	// no match: dst.Category.CategoryID
	dst.Category.Name = src.Category.Name
	dst.Name = "abc  def"
	// no match: dst.PhotoUrls
	// no match: dst.Status

	return
}

func ModelToDomain(src *model.Pet) (dst *domain.Pet) {
	dst = &domain.Pet{}
	// no match: dst.ID
	// no match: dst.Category.ID
	dst.Category.Name = src.Category.Name
	dst.Name = src.Name
	// no match: dst.PhotoUrls
	// no match: dst.Status

	return
}
