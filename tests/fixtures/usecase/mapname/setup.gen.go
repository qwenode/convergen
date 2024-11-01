// Code generated by github.com/qwenode/convergen
// DO NOT EDIT.

package mapname

import (
	"github.com/qwenode/convergen/tests/fixtures/data/domain"
	"github.com/qwenode/convergen/tests/fixtures/data/model"
)

func DomainToModel(src *domain.Pet) (dst *model.Pet) {
	dst = &model.Pet{}
	dst.ID = uint64(src.ID)
	dst.Category.CategoryID = uint64(src.Category.ID)
	dst.Category.Name = src.Category.Name
	dst.Name = src.Name
	if src.PhotoUrls != nil {
		dst.PhotoUrls = make([]string, len(src.PhotoUrls))
		for i, e := range src.PhotoUrls {
			dst.PhotoUrls[i] = string(e)
		}
	}
	dst.Status = src.Status.String()

	return
}
