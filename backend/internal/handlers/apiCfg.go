package handlers

import (
	"backend/internal/r2"
	"backend/internal/utils"
)

type RecentBuild struct {
	ID    string
	Level uint8
	Class string
}

type APIConfig struct {
	Env          *utils.EnvCfg
	R2           *r2.R2
	RecentBuilds []RecentBuild
}
