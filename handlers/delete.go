// Copyright (c) OpenFaaS Author(s) 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type deleteFunctionRequest struct {
	FunctionName string `json:"functionName"`
}

// MakeDeleteHandler delete a function
func MakeDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("delete request")
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)
		request := deleteFunctionRequest{}
		if err := json.Unmarshal(body, &request); err != nil {
			log.Errorf("error de-serializing request body:%s", body)
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(request.FunctionName) == 0 {
			log.Errorln("can not delete a function, request function name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		delete(functions, request.FunctionName)

		log.Infof("delete request %s successful", request.FunctionName)
	}
}
