package lint

import (
	"fmt"
	"regexp"

	lint "github.com/mafredri/go-lint"
)

var (
	reIDs   = regexp.MustCompile("^(.*)Ids($|[A-Z].*$)")
	reURLs  = regexp.MustCompile("^(.*)Urls($|[A-Z].*$)")
	reIDRef = regexp.MustCompile("^(.*)Idref($|[A-Z].*$)")
)

func init() {
	lint.SetInitialism("DOM", true)
	lint.SetInitialism("GPU", true)
	lint.SetInitialism("SSL", true)
	lint.SetInitialism("MAC", true)
}

// Name returns a different name if it should be different.
func Name(name string) (should string) {
	should = lint.Name(name)

	for _, replace := range []struct {
		re *regexp.Regexp
		to string
	}{
		{re: reIDs, to: "IDs"},
		{re: reURLs, to: "URLs"},
		{re: reIDRef, to: "IDRef"},
	} {
		should = replace.re.ReplaceAllString(should, fmt.Sprintf("${1}%s${2}", replace.to))
	}

	return should
}
