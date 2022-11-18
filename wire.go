//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

type Container struct {
}

func Inject() *Container {
	wire.Build(
		wire.Struct(new(Container), "*"),
	)
	return &Container{}
}
