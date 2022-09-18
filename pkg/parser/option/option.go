package option

import (
	"fmt"
	"strings"

	"github.com/reedom/convergen/pkg/model"
)

type FieldMatchSrc string

func (s FieldMatchSrc) String() string {
	return string(s)
}

const (
	FieldMatchField  = FieldMatchSrc("field")
	FieldMatchGetter = FieldMatchSrc("getter")
)

var FieldMatchSrcValues = []FieldMatchSrc{FieldMatchField, FieldMatchGetter}

func NewFieldMatchSrcFromValue(v string) (FieldMatchSrc, bool) {
	for _, src := range FieldMatchSrcValues {
		if src.String() == v {
			return src, true
		}
	}
	return FieldMatchSrc(""), false
}

func FieldMatchOrderFromValue(v string) ([]FieldMatchSrc, bool) {
	order := make([]FieldMatchSrc, 0)
	for _, s := range strings.Split(v, ",") {
		if src, ok := NewFieldMatchSrcFromValue(s); !ok {
			return nil, false
		} else {
			order = append(order, src)
		}
	}
	return order, true
}

type GlobalOption struct {
	Style           model.DstVarStyle
	FieldMatchOrder []FieldMatchSrc
	ExactCase       bool

	Receiver    string
	PostProcess string
	Skip        []*IdentMatcher
	Matchers    []any
	Converters  []any
}

func NewGlobalOption() *GlobalOption {
	return &GlobalOption{
		Style:           model.DstVarReturn,
		FieldMatchOrder: []FieldMatchSrc{FieldMatchField, FieldMatchGetter},
		ExactCase:       true,
		Converters:      make([]any, 0),
	}
}

type MethodOption struct {
	Style           model.DstVarStyle
	FieldMatchOrder []FieldMatchSrc
	ExactCase       bool
	PostProcess     string
	Skip            []IdentMatcher
	Matchers        []any
	Converters      []any
}

func (o *MethodOption) AddMatcher(m any) {
	switch m.(type) {
	case *NameMatcher:
		o.Matchers = append(o.Matchers, m)
	default:
		panic(fmt.Sprintf("unknown matcher: %q", m))
	}
}

func (o *MethodOption) AddConverter(c any) {
	switch c.(type) {
	case *FieldConverter:
		o.Converters = append(o.Converters, c)
	default:
		panic(fmt.Sprintf("unknown converter: %q", c))
	}
}

type FieldOption struct {
}
