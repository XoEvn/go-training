package repository

import (
	"eagle-demo/internal/model"
	"github.com/google/wire"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(model.Init())
