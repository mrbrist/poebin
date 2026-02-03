package handlers

import (
	"backend/internal/r2"
	"backend/internal/utils"
)

type APIConfig struct {
	Env *utils.EnvCfg
	R2  *r2.R2
}
