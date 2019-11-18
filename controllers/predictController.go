package controllers

import (
	"net/http"

	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
	"github.com/joaopandolfi/miia_api/models"
)

type PredictController struct {
}

func (cc PredictController) Predict(w http.ResponseWriter, r *http.Request) {
	var received models.SemusRecPayload
	form, err := handlers.GetForm(r)
	err = handlers.DecodeForm(&received, form)
	if err != nil {
		utils.Error("[PredictController][RESTNewPredict] - Erron on get body", err.Error())
		handlers.RESTResponseError(w, "Invalid body")
		return
	}

	parsed := received.Parse()
	json := parsed.Parse()
	utils.Debug("[PredictController][RestNewPredict] - RECEIVED", received)
	utils.Debug("[PredictController][RestNewPredict] - PARSED", parsed)
	utils.Debug("[PredictController][RestNewPredict] - JSON", json)
	handlers.RESTResponse(w, json)
}
