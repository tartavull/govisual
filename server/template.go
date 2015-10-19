package server

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"html/template"
)

type visual struct {
	name string
	desc string
	tmpl *template.Template
}

var visuals []visual

// Static files were embed in .go using https://github.com/GeertJohan/go.rice
// by doing this, when go install is called the static files are available to
// executable to be served.
var box *rice.Box

func init() {

	box = rice.MustFindBox("static")

	Register("index")
	Register("chord")
	Register("cluster")
	Register("force")
	Register("force2")
	Register("forcegraph")
	Register("forceimports")
	Register("radial")
	Register("tree")

}

// Register register a new visualisation at the url /name/.
// The URL parts after the /name/ are treated as the package name and
// make available in the template as .package.
func Register(name string) {

	tmpl, err := box.String("templates/" + name + ".html")
	if err != nil {
		fmt.Println(err)
		return
	}

	visuals = append(visuals, visual{
		name: name,
		tmpl: template.Must(template.New(name).Parse(tmpl)),
	})
}
