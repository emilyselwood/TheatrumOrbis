package schema

import "github.com/emilyselwood/TheatrumOrbis"

type Column struct {
	Name string
	Type TheatrumOrbis.Type
	DefaultValue TheatrumOrbis.Value
}
