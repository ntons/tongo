package template

import (
	"github.com/flosch/pongo2"

	_ "github.com/ntons/tongo/template/filters"
)

type M = pongo2.Context

func Render(s string, vals M) (string, error) {
	if tpl, err := pongo2.FromString(s); err != nil {
		return "", err
	} else {
		return tpl.Execute(vals)
	}
}

func RenderBytes(b []byte, vals M) ([]byte, error) {
	if tpl, err := pongo2.FromBytes(b); err != nil {
		return nil, err
	} else {
		return tpl.ExecuteBytes(vals)
	}
}

func RenderFile(fp string, vals M) ([]byte, error) {
	if tpl, err := pongo2.FromFile(fp); err != nil {
		return nil, err
	} else {
		return tpl.ExecuteBytes(vals)
	}
}
