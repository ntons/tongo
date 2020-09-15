package env

import (
	"os"
	"strings"

	"github.com/flosch/pongo2"
)

// get value from environ
// eg: {{ 'PATH'|env }}
func filterEnv(
	in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	key, val := in.String(), ""
	for _, s := range os.Environ() {
		pair := strings.SplitN(s, "=", 2)
		if key == pair[0] {
			val = pair[1]
			break
		}
	}
	return pongo2.AsSafeValue(val), nil
}
