package middleware

import (
	"myapp/data"

	"github.com/genesysflow/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
