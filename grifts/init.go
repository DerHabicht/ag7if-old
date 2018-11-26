package grifts

import (
	"github.com/derhabicht/ag7if/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
