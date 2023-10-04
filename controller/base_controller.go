package controller

import (
	"encoding/json"

	"github.com/dapr/go-sdk/service/common"
	"github.com/rs/zerolog/log"
)

type BaseController struct{}

func (ctrl *BaseController) DefinePublicRoutes() error { return nil }

func (ctrl *BaseController) DaprOut(in interface{}) *common.Content {
	out, err := json.Marshal(in)
	if err != nil {
		log.Fatal().Msg("can't marshal output")
	}
	return &common.Content{
		Data:        out,
		ContentType: "application/json",
	}
}
