package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/remotes/request"
	"github.com/joaopandolfi/blackwhale/utils"
	"github.com/joaopandolfi/miia_api/config"
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
		handlers.RESTResponseError(w, "Invalid body "+err.Error())
		return
	}

	parsed := received.Parse()
	//js := parsed.Parse()

	body, _ := json.Marshal(parsed.Values)

	var header map[string]string
	header = make(map[string]string)
	header["Content-Type"] = "application/json"
	result, err := request.PostWithHeader(config.Config.FlaskServer, header, body)

	var res []string
	err = json.Unmarshal(result, &res)

	utils.Debug("[PredictController][RestNewPredict] - JSON", res, string(result))
	handlers.RESTResponse(w, res)
}
