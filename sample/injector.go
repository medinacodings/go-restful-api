//go:build wireinject
// +build wireinject

package sample

import (
	"github.com/google/wire"
)

func InitializedService() *SimpleService {
	wire.Build(NewSimpleService, NewSimpleRepository)
	return nil
}
