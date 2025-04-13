package generator_test

import (
	"testing"

	"github.com/qwenode/convergen/pkg/generator"
	"github.com/qwenode/convergen/pkg/generator/model"
	"github.com/stretchr/testify/assert"
)

const pre = `package simple

import (
    "github.com/qwenode/convergen/pkg/tests/fixtures/data/domain"
    "github.com/qwenode/convergen/pkg/tests/fixtures/data/model"
)
`

const header = "// Code generated by github.com/qwenode/convergen\n// DO NOT EDIT.\n\n"

func TestGenerator_ArgRetReceiver(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		fn       *model.Function
		expected string
	}{
		{
			name: "src:ptr/dst:ptr,arg/rhs:simple",
			fn: &model.Function{
				Comments:    []string{"// comment 1", "// comment 2"},
				Name:        "ToModel",
				Receiver:    "",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    false,
				DstVarStyle: model.DstVarArg,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID"},
				},
			},
			expected: header + pre + `
// comment 1
// comment 2
func ToModel(dst *model.Pet, src *domain.Pet) {
    dst.ID = src.ID
}
`,
		},
		{
			name: "src:ptr/dst:ptr,return/rhs:simple",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    false,
				DstVarStyle: model.DstVarReturn,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID"},
				},
			},
			expected: header + pre + `
func ToModel(src *domain.Pet) (dst *model.Pet) {
    dst = &model.Pet{}
    dst.ID = src.ID

    return
}
`,
		},
		{
			name: "src:ptr/dst:copy,return/rhs:simple",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: false},
				RetError:    false,
				DstVarStyle: model.DstVarReturn,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID"},
				},
			},
			expected: header + pre + `
func ToModel(src *domain.Pet) (dst model.Pet) {
    dst.ID = src.ID

    return
}
`,
		},
		{
			name: "src:ptr,receiver/dst:copy,return/rhs:simple",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "src",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: false},
				RetError:    false,
				DstVarStyle: model.DstVarReturn,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID"},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel() (dst model.Pet) {
    dst.ID = src.ID

    return
}
`,
		},
		{
			name: "src:ptr,receiver/dst:ptr,arg/rhs:simple",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "src",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    false,
				DstVarStyle: model.DstVarArg,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID"},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel(dst *model.Pet) {
    dst.ID = src.ID
}
`,
		},
		{
			name: "src:ptr,receiver/dst:ptr,arg/error/rhs:simple",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "src",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    true,
				DstVarStyle: model.DstVarArg,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID()", Error: true},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel(dst *model.Pet) (err error) {
    dst.ID, err = src.ID()
    if err != nil {
        return
    }

    return
}
`,
		},
		{
			name: "src:ptr,receiver/dst:ptr,return/error/rhs:simple",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "src",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    true,
				DstVarStyle: model.DstVarReturn,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID()", Error: true},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel() (dst *model.Pet, err error) {
    dst = &model.Pet{}
    dst.ID, err = src.ID()
    if err != nil {
        return nil, err
    }

    return
}
`,
		},
		{
			name: "src:ptr,receiver/dst:val,return/error/rhs:simple",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "src",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: false},
				RetError:    true,
				DstVarStyle: model.DstVarReturn,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID()", Error: true},
				},
			},
			expected: header + pre + `
func (src *domain.Pet) ToModel() (dst model.Pet, err error) {
    dst.ID, err = src.ID()
    if err != nil {
        return
    }

    return
}
`,
		},
		{
			name: "src:ptr/dst:ptr,arg/error/rhs:skip",
			fn: &model.Function{
				Name:        "ToModel",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    true,
				DstVarStyle: model.DstVarArg,
				Assignments: []model.Assignment{
					model.SkipField{LHS: "dst.ID"},
				},
			},
			expected: header + pre + `
func ToModel(dst *model.Pet, src *domain.Pet) (err error) {
    // skip: dst.ID

    return
}
`,
		},
		{
			name: "src:ptr/dst:ptr,return/error/rhs:nomatch",
			fn: &model.Function{
				Name:        "ToModel",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    true,
				DstVarStyle: model.DstVarReturn,
				Assignments: []model.Assignment{
					model.NoMatchField{LHS: "dst.ID"},
				},
			},
			expected: header + pre + `
func ToModel(src *domain.Pet) (dst *model.Pet, err error) {
    dst = &model.Pet{}
    // no match: dst.ID

    return
}
`,
		},
		{
			name: "preprocess/postprocess",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    false,
				DstVarStyle: model.DstVarArg,
				Assignments: []model.Assignment{
					model.SimpleField{LHS: "dst.ID", RHS: "src.ID"},
				},
				PreProcess: &model.Manipulator{
					Name:     "PreProcess",
					IsSrcPtr: true,
					IsDstPtr: true,
					RetError: false,
				},
				PostProcess: &model.Manipulator{
					Pkg:      "domain",
					Name:     "PostProcess",
					IsSrcPtr: true,
					IsDstPtr: true,
					RetError: false,
				},
			},
			expected: header + pre + `
func ToModel(dst *model.Pet, src *domain.Pet) {
    PreProcess(dst, src)
    dst.ID = src.ID
    domain.PostProcess(dst, src)
}
`,
		},
		{
			name: "postprocess/dst,val/src,val/lhs,ptr/rhs,ptr/error",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: false},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: false},
				RetError:    true,
				DstVarStyle: model.DstVarReturn,
				PostProcess: &model.Manipulator{
					Name:     "PostProcess",
					IsSrcPtr: true,
					IsDstPtr: true,
					RetError: true,
				},
			},
			expected: header + pre + `
func ToModel(src domain.Pet) (dst model.Pet, err error) {
    err = PostProcess(&dst, &src)
    if err != nil {
        return
    }

    return
}
`,
		},
		{
			name: "postprocess/dst,ptr/src,ptr/lhs,val/rhs,val/error",
			fn: &model.Function{
				Name:        "ToModel",
				Receiver:    "",
				Src:         model.Var{Name: "src", Type: "domain.Pet", Pointer: true},
				Dst:         model.Var{Name: "dst", Type: "model.Pet", Pointer: true},
				RetError:    true,
				DstVarStyle: model.DstVarReturn,
				PostProcess: &model.Manipulator{
					Name:     "PostProcess",
					IsSrcPtr: false,
					IsDstPtr: false,
					RetError: true,
				},
			},
			expected: header + pre + `
func ToModel(src *domain.Pet) (dst *model.Pet, err error) {
    dst = &model.Pet{}
    err = PostProcess(*dst, *src)
    if err != nil {
        return
    }

    return
}
`,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				code := model.Code{
					BaseCode: pre + "xxxxx",
					FunctionBlocks: []model.FunctionsBlock{
						{
							Marker:    "xxxxx",
							Functions: []*model.Function{tt.fn},
						},
					},
				}
				g := generator.NewGenerator(code)
				actual, err := g.Generate("temp.gen.go", false, true)
				if assert.Nil(t, err) {
					assert.Equal(t, tt.expected, string(actual))
				}
			},
		)
	}
}
