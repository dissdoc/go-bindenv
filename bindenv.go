package gobindenv

import (
	"github.com/dissdoc/go-bindenv/internal"
)

var Version string = "0.1.0"

func BindEnv(cfg interface{}) {
	internal.ScanEnv(cfg)
}
