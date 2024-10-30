// Code generated by github.com/qwenode/convergen
// DO NOT EDIT.

package style

import (
	"github.com/qwenode/convergen/tests/fixtures/data/model"
)

type Pet struct {
	ID        uint64         `storage:"id"`
	Category  model.Category `storage:"category"`
	Name      string         `storage:"name"`
	PhotoUrls []string       `storage:"photoUrls"`
	Status    string         `storage:"status"`
}

func ArgToArg(dst *model.Pet, pet *Pet) {
	dst.ID = pet.ID
	dst.Category = pet.Category
	dst.Name = pet.Name
	if pet.PhotoUrls != nil {
		dst.PhotoUrls = make([]string, len(pet.PhotoUrls))
		copy(dst.PhotoUrls, pet.PhotoUrls)
	}
	dst.Status = pet.Status
}

func ArgToReturn(pet *Pet) (dst *model.Pet) {
	dst = &model.Pet{}
	dst.ID = pet.ID
	dst.Category = pet.Category
	dst.Name = pet.Name
	if pet.PhotoUrls != nil {
		dst.PhotoUrls = make([]string, len(pet.PhotoUrls))
		copy(dst.PhotoUrls, pet.PhotoUrls)
	}
	dst.Status = pet.Status

	return
}

func (r *Pet) RcvToArg(dst *model.Pet) {
	dst.ID = r.ID
	dst.Category = r.Category
	dst.Name = r.Name
	if r.PhotoUrls != nil {
		dst.PhotoUrls = make([]string, len(r.PhotoUrls))
		copy(dst.PhotoUrls, r.PhotoUrls)
	}
	dst.Status = r.Status
}

func (r *Pet) RcvToReturn(dst *model.Pet) {
	dst.ID = r.ID
	dst.Category = r.Category
	dst.Name = r.Name
	if r.PhotoUrls != nil {
		dst.PhotoUrls = make([]string, len(r.PhotoUrls))
		copy(dst.PhotoUrls, r.PhotoUrls)
	}
	dst.Status = r.Status
}

func (r *Pet) RevRcvFromArgPtr(pet *model.Pet) {
	r.ID = pet.ID
	r.Category = pet.Category
	r.Name = pet.Name
	if pet.PhotoUrls != nil {
		r.PhotoUrls = make([]string, len(pet.PhotoUrls))
		copy(r.PhotoUrls, pet.PhotoUrls)
	}
	r.Status = pet.Status
}

func (r *Pet) RevRcvFromArgVal(src *model.Pet) {
	r.ID = src.ID
	r.Category = src.Category
	r.Name = src.Name
	if src.PhotoUrls != nil {
		r.PhotoUrls = make([]string, len(src.PhotoUrls))
		copy(r.PhotoUrls, src.PhotoUrls)
	}
	r.Status = src.Status
}
