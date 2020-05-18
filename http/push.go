// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"encoding/json"
	"falcon-agent/g"
	"log"
	"net/http"

	"github.com/open-falcon/falcon-plus/common/model"
)

func configPushRoutes() {
	http.HandleFunc("/v1/push", func(w http.ResponseWriter, req *http.Request) {

		if g.Config().Debug {
			log.Println("[debug]/v1/push:", req)
		}

		if req.ContentLength == 0 {
			http.Error(w, "body is blank", http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(req.Body)

		var metrics []*model.MetricValue
		err := decoder.Decode(&metrics)
		if err != nil {
			http.Error(w, "connot decode body", http.StatusBadRequest)
			log.Println("connot decode body:", err)
			return
		}

		if g.Config().Debug {
			log.Println("[debug]/v1/push OK!:", req.Body)
			log.Println("[debug]/v1/push Decoder OK!:", metrics)

		}
		g.SendToTransfer(metrics)
		w.Write([]byte("success"))
	})
}
